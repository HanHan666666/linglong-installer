// Copyright 2024 The libtcl9_0-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package libtcl9_0 // import "modernc.org/libtcl9_0"

import (
	"fmt"
	goexec "os/exec"
	"syscall"
	"unsafe"

	"modernc.org/libc"
)

// tclBasic.c:85:1:
func _TclGetCStackPtr(tls *libc.TLS) (r uintptr) {
	p := tls.Alloc(8)
	defer tls.Free(8)
	return p
}

func _TclWinCPUID(tls *libc.TLS, index uint32, regsPtr uintptr) (r int32) {
	return TCL_ERROR
}

// int pthread_once(pthread_once_t *once_control, void (*init_routine)(void));
func _pthread_once(tls *libc.TLS, p, f uintptr) int32 {
	panic(todo(""))
}

// int pthread_kill(pthread_t thread, int sig);
func _pthread_kill(tls *libc.TLS, thread uintptr, sig int32) int32 {
	panic(todo(""))
}

/*
 *---------------------------------------------------------------------------
 *
 * TclpCreateProcess --
 *
 *	Create a child process that has the specified files as its standard
 *	input, output, and error. The child process runs asynchronously and
 *	runs with the same environment variables as the creating process.
 *
 *	The path is searched to find the specified executable.
 *
 * Results:
 *	The return value is TCL_ERROR and an error message is left in the
 *	interp's result if there was a problem creating the child process.
 *	Otherwise, the return value is TCL_OK and *pidPtr is filled with the
 *	process id of the child process.
 *
 * Side effects:
 *	A process is created.
 *
 *---------------------------------------------------------------------------
 */

// C documentation
func _TclpCreateProcess(tls *libc.TLS, interp uintptr, argc Tsize_t, argv uintptr, inputFile uintptr, outputFile uintptr, errorFile uintptr, pidPtr uintptr) (r int32) {
	bp := tls.Alloc(2 * 8)
	defer tls.Free(2 * 8)
	var args []string
	for i := 0; i < int(argc); i++ {
		p := *(*uintptr)(unsafe.Pointer(argv + unsafe.Sizeof(uintptr(0))*uintptr(i)))
		args = append(args, libc.GoString(p))
	}
	if len(args) == 0 {
		panic(todo(""))
	}

	args0, err := goexec.LookPath(args[0])
	if err != nil {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = libc.ENOENT
		s, err := libc.CString(fmt.Sprintf("couldn't execute \"%.150s\"", args[0]))
		if err != nil {
			panic(todo(""))
		}

		defer libc.Xfree(tls, s)

		XTcl_SetObjResult(tls, interp, XTcl_ObjPrintf(tls, uintptr(unsafe.Pointer(&createProcessMsg[0])), libc.VaList(bp, s, XTcl_PosixError(tls, interp))))
		return TCL_ERROR
	}

	args[0] = args0
	env := libc.GetEnviron()
	attr := &syscall.ProcAttr{
		Env:   env,
		Files: []uintptr{^uintptr(0), ^uintptr(0), ^uintptr(0)},
	}
	if inputFile != 0 {
		attr.Files[syscall.Stdin] = inputFile - 1
	}
	if outputFile != 0 {
		attr.Files[syscall.Stdout] = outputFile - 1
	}
	if errorFile != 0 {
		attr.Files[syscall.Stderr] = errorFile - 1
	}
	pid, err := syscall.ForkExec(args0, args, attr)
	if err != nil {
		trc("TclpCreateProcess(%#x, %d, %q, %v, %v, %v, %#x): %v", interp, len(args), args, inputFile, outputFile, errorFile, pidPtr, err)
		panic(todo(""))
	}

	*(*uintptr)(unsafe.Pointer(pidPtr)) = uintptr(pid)
	return TCL_OK
}

var createProcessMsg = [...]byte{'%', 's', ':', ' ', '%', 's', 0}
