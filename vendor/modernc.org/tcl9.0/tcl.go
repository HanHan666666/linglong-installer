// Copyright 2024 The tcl9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tcl9.0 is an idiomatic Go wrapper for [libtcl9.0].
//
// [libtcl9.0]: https://pkg.go.dev/modernc.org/libtcl9.0
package tcl9_0 // import "modernc.org/tcl9.0"

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/evilsocket/islazy/zip"
	"modernc.org/libc"
	lib "modernc.org/libtcl9.0"
	tcllib "modernc.org/libtcl9.0/library"
)

var (
	token atomic.Uintptr

	objectMu sync.Mutex
	objects  = map[uintptr]interface{}{}
)

func addObject(o interface{}) uintptr {
	t := token.Add(1)
	objectMu.Lock()
	objects[t] = o
	objectMu.Unlock()
	return t
}

func getObject(t uintptr) interface{} {
	objectMu.Lock()
	o := objects[t]
	if o == nil {
		panic(todo("%#x", t))
	}

	objectMu.Unlock()
	return o
}

func removeObject(t uintptr) {
	objectMu.Lock()
	if _, ok := objects[t]; !ok {
		panic(todo("%#x"))
	}

	delete(objects, t)
	objectMu.Unlock()
}

// EvalFlag is a type used to parameterize Interp.Eval().
type EvalFlag int32

const (
	// If this flag bit is set, the script is not compiled to bytecodes; instead it
	// is executed directly as is done by Tcl_EvalEx. The TCL_EVAL_DIRECT flag is
	// useful in situations where the contents of a value are going to change
	// immediately, so the bytecodes will not be reused in a future execution. In
	// this case, it is faster to execute the script directly.
	EvalDirect = lib.TCL_EVAL_DIRECT
	// If this flag is set, the script is evaluated in the global namespace instead
	// of the current namespace and its variable context consists of global
	// variables only (it ignores any Tcl procedures that are active).
	EvalGlobal = lib.TCL_EVAL_GLOBAL
)

var (
	onceStdlib    sync.Once
	onceStdlibErr error
	stdlib        string
)

// Stdlib returns the path to the Tcl standard library or an error, if any. It
// once creates a temporary directory where the standard library is written.
// Subsequent calls to Stdlib share the same temporary directory.
//
// Stdlib is safe for concurrent access by multiple goroutines.
func Stdlib() (string, error) {
	onceStdlib.Do(func() {
		dir, err := os.MkdirTemp("", "tcl-library-")
		defer func() { onceStdlibErr = err }()
		if err != nil {
			return
		}

		fn := filepath.Join(dir, "library.zip")
		if err = os.WriteFile(fn, []byte(tcllib.Zip), 0600); err != nil {
			return
		}

		if _, err = zip.Unzip(fn, dir); err != nil {
			return
		}

		stdlib = filepath.Join(dir, "library")
	})
	return stdlib, onceStdlibErr
}

// MustStdlib is like Stdlib but panics on error.
func MustStdlib() (r string) {
	r, err := Stdlib()
	if err != nil {
		panic(err)
	}

	return r
}

// Interp represents a Tcl interpreter instance.
type Interp struct {
	tls    *libc.TLS
	interp uintptr
}

// NewInterp returns a newly created Interp or an error, if any.
//
// The 'tclvars' argument is used to set any interpreter variables before
// initializing it.
func NewInterp(tclvars map[string]string) (r *Interp, err error) {
	tls := libc.NewTLS()
	interp := lib.XTcl_CreateInterp(tls)
	if interp == 0 {
		tls.Close()
		return nil, fmt.Errorf("failed to create a Tcl interpreter")
	}

	r = &Interp{tls, interp}
	for k, v := range tclvars {
		cmd := fmt.Sprintf(`set {%s} {%s};`, k, v)
		if _, err := r.Eval(fmt.Sprintf(cmd), EvalGlobal); err != nil {
			r.Close()
			return nil, fmt.Errorf("eval failed: %s", cmd)
		}
	}

	if rc := lib.XTcl_Init(tls, interp); rc != lib.TCL_OK {
		r.Close()
		return nil, fmt.Errorf("failed to initialize the Tcl interpreter")
	}

	return r, nil
}

// TLS returns the thread local storage of the underlying interpreter. It is
// used when calling libtcl directly
func (in *Interp) TLS() *libc.TLS { return in.tls }

// Handle returns the handle of the underlying interpreter. It is used when
// calling libtcl directly:
func (in *Interp) Handle() uintptr { return in.interp }

// Close invalidates the interpreter and releases all its associated resources.
func (in *Interp) Close() error {
	lib.XTcl_DeleteInterp(in.tls, in.interp)
	in.tls.Close()
	*in = Interp{}
	return nil
}

// Eval evaluates script and returns the interpreter's result and an error, if any.
func (in *Interp) Eval(script string, flag EvalFlag) (r string, err error) {
	s, err := libc.CString(script)
	if err != nil {
		return "", err
	}

	defer libc.Xfree(in.tls, s)

	lib.XTcl_Preserve(in.tls, in.interp)

	defer lib.XTcl_Release(in.tls, in.interp)

	if rc := lib.XTcl_EvalEx(in.tls, in.interp, s, lib.TTcl_Size(len(script)), int32(flag)); rc != lib.TCL_OK {
		err = fmt.Errorf("return code: %d", rc)
	}
	return in.StringResult(), err
}

// StringResult returns the interpreter result as a string.
func (in *Interp) StringResult() (r string) {
	return libc.GoString(lib.XTcl_GetString(in.tls, lib.XTcl_GetObjResult(in.tls, in.interp)))
}

// SetResult sets the result of the interpreter.
func (in *Interp) SetResult(s string) error {
	cs, err := libc.CString(s)
	if err != nil {
		return err
	}

	defer libc.Xfree(in.tls, cs)

	lib.XTcl_SetObjResult(in.tls, in.interp, lib.XTcl_NewStringObj(in.tls, cs, lib.TTcl_Size(len(s))))
	return nil
}

// CmdProc is a Tcl command implemented in Go.
type CmdProc func(clientData interface{}, in *Interp, args []string) int

// DeleteProc is a function called when CmdProc is deleted.
type DeleteProc func(clientData interface{})

type cmdProc struct {
	clientData interface{}
	del        DeleteProc
	f          CmdProc
	in         *Interp
}

func fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var (
	runCmdP = fp(runCmd)
	delCmdP = fp(delCmd)
)

func runCmd(tls *libc.TLS, clientData, in uintptr, argc int32, argv uintptr) int32 {
	cmd := getObject(clientData).(*cmdProc)
	var a []string
	for i := int32(0); i < argc; i++ {
		p := *(*uintptr)(unsafe.Pointer(argv))
		argv += unsafe.Sizeof(argv)
		a = append(a, libc.GoString(p))
	}
	return int32(cmd.f(cmd.clientData, cmd.in, a))
}

func delCmd(tls *libc.TLS, clientData uintptr) {
	cmd := getObject(clientData).(*cmdProc)
	if cmd.del != nil {
		cmd.del(cmd.clientData)
	}
	removeObject(clientData)
}

// RegisterCommand creates a new Tcl command.
func (in *Interp) RegisterCommand(name string, proc CmdProc, clientData interface{}, del DeleteProc) (err error) {
	nm, err := libc.CString(name)
	if err != nil {
		return err
	}

	lib.XTcl_Preserve(in.tls, in.interp)

	defer func() {
		libc.Xfree(in.tls, nm)
		lib.XTcl_Release(in.tls, in.interp)
	}()

	p := &cmdProc{f: proc, clientData: clientData, del: del, in: in}
	h := addObject(p)
	cmd := lib.XTcl_CreateCommand(in.tls, in.interp, nm, runCmdP, h, delCmdP)
	if cmd == 0 {
		return fmt.Errorf("failed to create command: %s", name)
	}

	return nil
}
