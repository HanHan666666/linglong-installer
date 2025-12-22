// Copyright 2023 The Regexp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp // modernc.org/regexp

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"modernc.org/fsm"
	"modernc.org/regexp/syntax"
	"modernc.org/sortutil"
)

const (
	// DFARuneBits is the number of lower bits that are used for encoding
	// the instruction operand, if any.  The remaining upper bits of the
	// uint32 DFA instruction carry the op code.
	DFARuneBits = 21
	// DFARuneMask isolates the encoded rune/character class ID argument
	// from the uint32 DFA instruction.
	DFARuneMask = 1<<DFARuneBits - 1

	halfRuneBits = 10
	halfRuneMask = 1<<halfRuneBits - 1
	maxFSMStates = 1 << 16
	maxSplits    = 64
)

// DFA instruction op codes. The particular value may change, always refer to
// the value by name.
const (
	// DFAOpRune jumps to the prog address in the following prog word if
	// the argument of the instruction matches the input.
	//
	//	pc:	DFAOpRune<<DFARuneBits | rune
	//	pc+1:	<next pc>
	DFAOpRune = iota
	// DFAOpCharClass jumps to the prog address in the following prog word if
	// the character class of the instruction matches the input.
	//
	//	pc:	DFAOpCharClass<<DFARuneBits | character class ID
	//	pc+1:	<next pc>
	DFAOpCharClass
	// DFAOpAccept puts the DFA in matched state.
	//
	//	pc:	DFAOpAccept<<DFARuneBits | lex ID
	//
	// If the DFA was created using NewLexDFA then the ID determines the zero based
	// regexp number accepted in this state.  Otherwise the ID is undefined.
	DFAOpAccept
	// DFAOpStop stops DFA execution.
	//
	//	pc:	DFAOpStop<<DFARuneBits
	DFAOpStop
	// DFAOpEndText matches the end of input.
	//
	//	pc:	DFAOpEndText<<DFARuneBits
	DFAOpEndText
	// DFAOpBeginText matches the start of input.
	//
	//	pc:	DFAOpBeginText<<DFARuneBits
	DFAOpBeginText

	dfaOpState
	opAcceptThreads
	opCapture
	opKillThreads
	opThreads
)

var (
	opSortOrder = [...]int{
		dfaOpState:      1,
		DFAOpAccept:     2,
		opCapture:       3,
		DFAOpBeginText:  4,
		DFAOpRune:       5,
		DFAOpCharClass:  6,
		DFAOpEndText:    7,
		DFAOpStop:       8,
		opThreads:       9,
		opAcceptThreads: 10,
		opKillThreads:   11,
	}
)

type mask struct {
	a [(maxSplits + 63) >> 6]uint64
}

func (m *mask) isNonZero() bool {
	for _, v := range m.a {
		if v != 0 {
			return true
		}
	}
	return false
}

func (m *mask) set(bit int) {
	m.a[bit>>6] |= uint64(1) << (bit & 63)
}

func (m *mask) and(n mask) {
	for i := range m.a {
		m.a[i] &= n.a[i]
	}
}

func (m *mask) andNot(n mask) {
	for i := range m.a {
		m.a[i] &^= n.a[i]
	}
}

func (m *mask) or(n mask) {
	for i := range m.a {
		m.a[i] |= n.a[i]
	}
}

type threadMask struct {
	l, r mask
}

func (m *threadMask) isNonZero() bool {
	for i, v := range m.l.a {
		if v != 0 || m.r.a[i] != 0 {
			return true
		}
	}
	return false
}

func (m *threadMask) andNot(n threadMask) {
	m.l.andNot(n.l)
	m.r.andNot(n.r)
}

func (m *threadMask) and(n threadMask) {
	m.l.and(n.l)
	m.r.and(n.r)
}

func (m *threadMask) or(n threadMask) {
	m.l.or(n.l)
	m.r.or(n.r)
}

type register[T comparable] struct {
	toID map[T]int
	reg  []T
}

func (r *register[T]) byID(id int) T {
	if id < len(r.reg) {
		return r.reg[id]
	}

	var zero T
	return zero
}

func (r *register[T]) id(x T) int {
	if id, ok := r.toID[x]; ok {
		return id
	}

	if r.toID == nil {
		r.toID = map[T]int{}
	}
	id := len(r.toID)
	r.toID[x] = id
	r.reg = append(r.reg, x)
	return id
}

type split struct {
	splitID     int
	outL, outR  *fsm.State
	left, right *split
	killThreads threadMask
}

type nfa struct {
	*fsm.NFA
	caps               map[int]struct{}
	charClasses        register[string]
	id2captureKills    [][]int
	id2charClass       []*runeSlice
	id2split           []*split
	re                 *Regexp
	rootSplit          *split
	sink               *fsm.State
	start              *fsm.State
	state2ThreadMaskID map[*fsm.State]int
	state2split        map[*fsm.State]*split
	threadMasks        register[threadMask]

	hasCapture     bool
	hasEOT         bool
	hasRuneOrClass bool
}

func (re *Regexp) compileNFA(sre *syntax.Regexp) (r *nfa) {
	r = &nfa{
		NFA: fsm.NewLimitedNFA(maxFSMStates),
		re:  re,
	}
	if !r.canCompile(sre) {
		return nil
	}

	r.start = r.NewState()
	r.sink = r.NewState()
	if r.start == nil || r.sink == nil {
		return nil
	}

	r.sink.IsAccepting = true
	outs := r.add(r.start, sre, 0, nil)
	if len(outs) == 0 {
		return nil
	}

	if r.hasEOT && !r.hasRuneOrClass {
		return nil
	}

	for _, out := range outs {
		s := r.NewState()
		if s == nil {
			return nil
		}

		s.IsAccepting = true
		out.NewEdge(DFAOpAccept<<DFARuneBits+out.Id(), s)
	}
	r.computeMasks()
	for _, state := range r.List() {
		state.NewEdge(dfaOpState<<DFARuneBits+state.Id(), r.sink)
	}
	if r.rootSplit != nil {
		id2captureThreads := map[int]*threadMask{}
		for _, state := range r.List() {
			edges := state.Transitions()
			for _, sym := range edges.List() {
				switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
				case opCapture:
					ix := c & halfRuneMask
					tm := r.threadMasks.byID(r.state2ThreadMaskID[state])
					m := id2captureThreads[ix]
					if m == nil {
						id2captureThreads[ix] = &tm
						continue
					}

					if *m != tm {
						return nil
					}
				}
			}
		}
	}
	for i, v := range r.id2captureKills {
		n := sortutil.Dedupe(sort.IntSlice(v))
		var w []int
		if n != 0 {
			w = make([]int, n)
			copy(w, v)
			r.id2captureKills[i] = w
		}
	}

	return r
}

func (n *nfa) canCompile(sre *syntax.Regexp) bool {
	switch sre.Op {
	case syntax.OpNoMatch:
		return false
	case syntax.OpEmptyMatch:
		return true
	case syntax.OpLiteral:
		return true
	case syntax.OpCharClass:
		n.hasRuneOrClass = true
		return true
	case syntax.OpAnyCharNotNL:
		n.hasRuneOrClass = true
		return true
	case syntax.OpAnyChar:
		n.hasRuneOrClass = true
		return true
	case syntax.OpBeginLine:
		return false
	case syntax.OpEndLine:
		return false
	case syntax.OpBeginText:
		return true
	case syntax.OpEndText:
		n.hasEOT = true
		return true
	case syntax.OpWordBoundary:
		return false
	case syntax.OpNoWordBoundary:
		return false
	case syntax.OpCapture:
		switch n.last(sre.Sub[0]) {
		case syntax.OpStar:
			return false
		case syntax.OpPlus:
			return false
		case syntax.OpAlternate:
			return false
		}

		if n.caps == nil {
			n.caps = map[int]struct{}{}
		}

		if _, ok := n.caps[sre.Cap]; ok {
			return false
		}

		n.caps[sre.Cap] = struct{}{}
		n.hasCapture = true
		return n.canCompile(sre.Sub[0])
	case syntax.OpStar:
		switch n.last(sre.Sub[0]) {
		case syntax.OpStar:
			return false
		case syntax.OpQuest:
			return false
		case syntax.OpCapture:
			return false
		}

		if n.hasEmptyMatch(sre.Sub[0]) {
			return false
		}

		return n.canCompile(sre.Sub[0])
	case syntax.OpPlus:
		switch n.last(sre.Sub[0]) {
		case syntax.OpStar:
			return false
		case syntax.OpQuest:
			return false
		case syntax.OpCapture:
			return false
		}

		if n.hasEmptyMatch(sre.Sub[0]) {
			return false
		}

		return n.canCompile(sre.Sub[0])
	case syntax.OpQuest:
		switch n.last(sre.Sub[0]) {
		case syntax.OpCapture:
			return false
		}

		if n.hasEmptyMatch(sre.Sub[0]) {
			return false
		}

		return n.canCompile(sre.Sub[0])
	case syntax.OpRepeat:
		return false
	case syntax.OpConcat:
		for _, v := range sre.Sub {
			if !n.canCompile(v) {
				return false
			}
		}

		return true
	case syntax.OpAlternate:
		for _, v := range sre.Sub {
			if v.Op == syntax.OpEmptyMatch || !n.canCompile(v) {
				return false
			}
		}

		return true
	default:
		return false
	}
}

func (n *nfa) hasEmptyMatch(sre *syntax.Regexp) bool {
	switch sre.Op {
	case syntax.OpEmptyMatch:
		return true
	case syntax.OpConcat:
		for _, v := range sre.Sub {
			if n.hasEmptyMatch(v) {
				return true
			}
		}

		return false
	case syntax.OpCapture:
		for _, v := range sre.Sub {
			if n.hasEmptyMatch(v) {
				return true
			}
		}

		return false
	case syntax.OpAlternate:
		for _, v := range sre.Sub {
			if n.hasEmptyMatch(v) {
				return true
			}
		}

		return false
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest:
		return n.hasEmptyMatch(sre.Sub[0])
	default:
		return false
	}
}

func (n *nfa) computeMasks() {
	if n.rootSplit == nil {
		return
	}

	type key struct {
		state        *fsm.State
		threadMaskID int
	}

	var f func(state *fsm.State, threadMaskID int, m map[key]struct{}, link **split)

	f = func(state *fsm.State, threadMaskID int, m map[key]struct{}, link **split) {
		k := key{state, threadMaskID}
		if _, ok := m[k]; ok {
			return
		}

		m[k] = struct{}{}
		mask := n.threadMasks.byID(n.state2ThreadMaskID[state])
		mask.or(n.threadMasks.byID(threadMaskID))
		threadMaskID = n.threadMasks.id(mask)
		n.state2ThreadMaskID[state] = threadMaskID
		var outL, outR *fsm.State
		var splitID int
		split := n.state2split[state]
		if split != nil {
			splitID = split.splitID
			outL = split.outL
			outR = split.outR
			if link != nil {
				*link = split
			}
		}
		x := state.Transitions()
		for _, sym := range x.List() {
			for _, next := range x.Get(sym).List() {
				switch {
				case next == outL:
					maskL := mask
					maskL.l.set(splitID)
					f(next, n.threadMasks.id(maskL), m, &split.left)
				case next == outR:
					maskR := mask
					maskR.r.set(splitID)
					f(next, n.threadMasks.id(maskR), m, &split.right)
				default:
					if sym != fsm.Epsilon {
						f(next, threadMaskID, m, link)
					}
				}
			}
		}
	}

	n.state2ThreadMaskID = map[*fsm.State]int{}
	m := map[key]struct{}{}
	f(n.start, n.threadMasks.id(threadMask{}), m, nil)
	n.killThreads(n.rootSplit)
}

func (n *nfa) killThreads(s *split) (r mask) {
	s.killThreads.l.set(s.splitID)
	if s.left != nil {
		s.killThreads.l.or(n.killThreads(s.left))
	}
	if s.right != nil {
		s.killThreads.r.or(n.killThreads(s.right))
	}
	r = s.killThreads.l
	r.or(s.killThreads.r)
	return r
}

func (n *nfa) add(in *fsm.State, sre *syntax.Regexp, ctx syntax.Op, caps map[int]struct{}) (outs []*fsm.State) {
	var out *fsm.State
	switch sre.Op {
	case syntax.OpLiteral:
		n.hasRuneOrClass = true
		for _, v := range sre.Rune {
			if out = n.NewState(); out == nil {
				return nil
			}

			in.NewEdge(DFAOpRune<<DFARuneBits+int(v), out)
			if sre.Flags&syntax.FoldCase != 0 {
				if v2 := unicode.SimpleFold(v); v2 != v {
					in.NewEdge(DFAOpRune<<DFARuneBits+int(v2), out)
				}
			}
			in = out
		}
		return append(outs, out)
	case syntax.OpCapture:
		if l := len(n.id2captureKills); l < 2*sre.Cap+2 {
			n.id2captureKills = append(n.id2captureKills, make([][]int, 2*sre.Cap+2-l)...)
		}
		in.NewEdge(opCapture<<DFARuneBits+sre.Cap*2<<halfRuneBits+in.Id(), n.sink)
		if caps != nil {
			caps[sre.Cap*2] = struct{}{}
			caps[sre.Cap*2+1] = struct{}{}
		}
		if outs = n.add(in, sre.Sub[0], ctx, caps); len(outs) == 0 {
			return nil
		}

		for _, v := range outs {
			v.NewEdge(opCapture<<DFARuneBits+(sre.Cap*2+1)<<halfRuneBits+v.Id(), n.sink) // op:cap:sid
		}
		return outs
	case syntax.OpCharClass:
		return n.charClass(in, (*runeSlice)(&sre.Rune))
	case syntax.OpAnyCharNotNL:
		return n.charClass(in, (*runeSlice)(&anyRuneNotNL))
	case syntax.OpAnyChar:
		return n.charClass(in, (*runeSlice)(&anyRune))
	case syntax.OpBeginText:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(DFAOpBeginText<<DFARuneBits, out)
		return append(outs, out)
	case syntax.OpEndText:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(DFAOpEndText<<DFARuneBits, out)
		return append(outs, out)
	case syntax.OpStar:
		outL, outR := n.split(in)
		if outL == nil || outR == nil {
			return nil
		}

		if sre.Flags&syntax.NonGreedy != 0 {
			outL, outR = outR, outL
		}

		aa := n.add(outL, sre.Sub[0], syntax.OpStar, caps)
		if len(aa) == 0 {
			return nil
		}

		for _, a := range aa {
			a.NewEdge(fsm.Epsilon, outL)
		}
		return append(aa, outR)
	case syntax.OpPlus:
		aa := n.add(in, sre.Sub[0], syntax.OpPlus, caps)
		for _, a := range aa {
			outL, outR := n.split(a)
			if outL == nil || outR == nil {
				return nil
			}

			if sre.Flags&syntax.NonGreedy != 0 {
				outL, outR = outR, outL
			}
			bb := n.add(outL, sre.Sub[0], syntax.OpPlus, caps)
			if len(bb) == 0 {
				return nil
			}

			for _, b := range bb {
				b.NewEdge(fsm.Epsilon, outL)
			}
			outs = append(outs, bb...)
			outs = append(outs, outR)
		}
		return outs
	case syntax.OpQuest:
		outL, outR := n.split(in)
		if outL == nil || outR == nil {
			return nil
		}

		if sre.Flags&syntax.NonGreedy != 0 {
			outL, outR = outR, outL
		}
		outs := n.add(outL, sre.Sub[0], syntax.OpQuest, caps)
		if len(outs) == 0 {
			return nil
		}

		return append(outs, outR)
	case syntax.OpConcat:
		ins := []*fsm.State{in}
		for _, v := range sre.Sub {
			if v.Op == syntax.OpEmptyMatch {
				continue
			}

			outs = nil
			for _, in := range ins {
				o := n.add(in, v, ctx, caps)
				if len(o) == 0 {
					return nil
				}

				outs = append(outs, o...)
			}
			ins = outs
		}
		return outs
	case syntax.OpAlternate:
		return n.alt(in, sre.Sub)
	case syntax.OpEmptyMatch:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(fsm.Epsilon, out)
		return append(outs, out)
	default:
		return nil
	}
	panic("unreachable")
}

func (n *nfa) last(sre *syntax.Regexp) (r syntax.Op) {
	for {
		switch r = sre.Op; r {
		case syntax.OpConcat:
			sre = sre.Sub[len(sre.Sub)-1]
		default:
			return r
		}
	}
}

func (n *nfa) charClass(in *fsm.State, p *runeSlice) (outs []*fsm.State) {
	out := n.NewState()
	if out == nil {
		return nil
	}

	id := n.charClasses.id(p.hashString())
	if id == len(n.id2charClass) {
		n.id2charClass = append(n.id2charClass, p)
	}
	in.NewEdge(DFAOpCharClass<<DFARuneBits+id, out)
	return append(outs, out)
}

func (n *nfa) alt(in *fsm.State, subs []*syntax.Regexp) (outs []*fsm.State) {
	outL, outR := n.split(in)
	if outL == nil || outR == nil {
		return nil
	}

	if len(subs) < 2 {
		panic("internal error")
	}

	capsL := map[int]struct{}{}
	capsR := map[int]struct{}{}

	defer func() {
		for k := range capsL {
			if k&1 == 0 {
				continue
			}

			for l := range capsR {
				n.id2captureKills[k] = append(n.id2captureKills[k], l)
			}
		}
		for k := range capsR {
			if k&1 == 0 {
				continue
			}

			for l := range capsL {
				n.id2captureKills[k] = append(n.id2captureKills[k], l)
			}
		}
	}()

	outsL := n.add(outL, subs[0], syntax.OpAlternate, capsL)
	if len(outsL) == 0 {
		return nil
	}

	switch len(subs) {
	case 2:
		outsR := n.add(outR, subs[1], syntax.OpAlternate, capsR)
		if len(outsR) == 0 {
			return nil
		}

		return append(outsL, outsR...)
	default:
		if len(capsL) != 0 {
			return nil
		}

		outsR := n.alt(outR, subs[1:])
		if len(outsR) == 0 {
			return nil
		}

		return append(outsL, outsR...)
	}
}

func (n *nfa) split(in *fsm.State) (outL, outR *fsm.State) {
	s := &split{
		splitID: len(n.state2split),
		outL:    n.NewState(),
		outR:    n.NewState(),
	}
	if s.splitID >= maxSplits || s.outL == nil || s.outR == nil {
		return nil, nil
	}

	n.id2split = append(n.id2split, s)
	in.NewEdge(fsm.Epsilon, s.outL)
	in.NewEdge(fsm.Epsilon, s.outR)
	if n.state2split == nil {
		n.state2split = map[*fsm.State]*split{}
		n.rootSplit = s
	}
	n.state2split[in] = s
	return s.outL, s.outR
}

type dfaProg struct {
	id2captureKills [][]int
	id2charClass    []*runeSlice
	id2threadMask   []threadMask
	prog            []uint32
	re              *Regexp

	startPC int

	hasCapture bool
}

func (re *Regexp) compileDFA(sre *syntax.Regexp) {
	nfa := re.compileNFA(sre)
	if nfa == nil {
		return
	}

	dfa := nfa.MinimalDFA(false)
	if dfa == nil {
		return
	}

	d := &dfaProg{
		hasCapture:      nfa.hasCapture,
		id2captureKills: nfa.id2captureKills,
		id2charClass:    nfa.id2charClass,
		re:              re,
	}
	states := dfa.List()
	rebuild := false
	for _, state := range states {
		edges := state.Transitions()
		syms := edges.List()
		sort.Slice(syms, func(i, j int) bool {
			a, b := syms[i]>>DFARuneBits, syms[j]>>DFARuneBits
			return opSortOrder[a] < opSortOrder[b]
		})
		type info struct {
			rune rune
			next *fsm.State
		}
		var infos []info
		hasCharClass := false
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			if len(nexts) != 1 {
				return
			}

			next := nexts[0]
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				infos = append(infos, info{rune(c), next})
			case DFAOpCharClass:
				if hasCharClass {
					return
				}

				hasCharClass = true
				charClass := nfa.id2charClass[c]
				del := false
				for _, info := range infos {
					c := info.rune
					if charClass.has(c) {
						switch {
						case charClass.cardinality() <= 10:
							rebuild = true
							hasCharClass = false
							cls := *charClass
							if !del {
								for i := 0; i < len(cls); i += 2 {
									for lo, hi := cls[i], cls[i+1]; lo <= hi; lo++ {
										state.NewEdge(int(DFAOpRune<<DFARuneBits+lo), next)
									}
								}
							}
							del = true
							state.NewEdge(int(DFAOpRune<<DFARuneBits+c), info.next)
						default:
							return
						}
					}
				}
				if del {
					edges.Delete(sym)
				}
			}
		}
	}
	if rebuild {
		if dfa = dfa.MinimalDFA(false); dfa == nil {
			return
		}
	}

	state2pc := []int{}
	states = dfa.List()
	linker := make([][]int, len(states))
	var threadMasks register[threadMask]
	for _, state := range states {
		state2pc = append(state2pc, len(d.prog))
		if state.IsAccepting {
			continue
		}

		edges := state.Transitions()
		syms := edges.List()
		sort.Slice(syms, func(i, j int) bool {
			a, b := syms[i]>>DFARuneBits, syms[j]>>DFARuneBits
			if opSortOrder[a] < opSortOrder[b] {
				return true
			}
			if opSortOrder[a] > opSortOrder[b] {
				return false
			}

			if a == opCapture {
				capi, capj := syms[i]>>halfRuneBits, syms[j]>>halfRuneBits
				return capi < capj
			}

			return syms[i]&DFARuneMask < syms[j]&DFARuneMask
		})
		hasCharClass := false
		var threads, accepts, killThreads threadMask
		isAccepting := false
		var caps map[int][]int
		for _, sym := range syms {
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case dfaOpState:
				if threadMaskID, ok := nfa.state2ThreadMaskID[nfa.State(c)]; ok {
					threads.or(nfa.threadMasks.byID(threadMaskID))
				}
			case DFAOpAccept:
				isAccepting = true
				if threadMaskID, ok := nfa.state2ThreadMaskID[nfa.State(c)]; ok {
					threadMask := nfa.threadMasks.byID(threadMaskID)
					accepts.or(threadMask)
					for y, v := range threadMask.l.a {
						for splitID := 64 * y; v != 0; splitID, v = splitID+1, v>>1 {
							if v&1 != 0 {
								killThreads.or(nfa.id2split[splitID].killThreads)
							}
						}
					}
				}
			case opCapture: // op:cap:sid
				if caps == nil {
					caps = map[int][]int{}
				}
				ix := c >> halfRuneBits
				stateID := c & halfRuneMask
				threadMaskID := nfa.state2ThreadMaskID[nfa.State(stateID)]
				caps[ix] = append(caps[ix], threadMasks.id(nfa.threadMasks.byID(threadMaskID)))

			}
		}
		switch {
		case nfa.rootSplit != nil && !isAccepting:
			d.prog = append(d.prog, uint32(opThreads<<DFARuneBits+threadMasks.id(threads)))
		case nfa.rootSplit != nil && isAccepting:
			d.prog = append(d.prog, uint32(opAcceptThreads<<DFARuneBits+threadMasks.id(threads)))
			d.prog = append(d.prog, uint32(DFAOpAccept<<DFARuneBits+threadMasks.id(accepts)))
			d.prog = append(d.prog, uint32(opKillThreads<<DFARuneBits+threadMasks.id(killThreads)))
		case isAccepting:
			d.prog = append(d.prog, uint32(DFAOpAccept<<DFARuneBits))
		}
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			if len(nexts) != 1 {
				// panic(todo("%d: %#U\n%s", state.Id(), sym, dfa))
				return
			}

			next := nexts[0].Id()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case opCapture: // op:cap:sid
				cap := c >> halfRuneBits
				sid := c & halfRuneMask
				tid0 := nfa.state2ThreadMaskID[nfa.State(sid)]
				tid := threadMasks.id(nfa.threadMasks.byID(tid0))
				d.prog = append(d.prog, uint32(opCapture<<DFARuneBits+cap<<halfRuneBits+tid))
			case DFAOpRune:
				d.prog = append(d.prog, uint32(DFAOpRune<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
			case DFAOpEndText:
				d.prog = append(d.prog, uint32(DFAOpEndText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
			case DFAOpCharClass:
				if hasCharClass {
					return
				}

				hasCharClass = true
				d.prog = append(d.prog, uint32(DFAOpCharClass<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
			case DFAOpBeginText:
				if state != dfa.Start() || threads.isNonZero() {
					return
				}

				d.prog = append(d.prog, uint32(DFAOpBeginText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
			case DFAOpAccept, dfaOpState:
				// nop
			default:
				return
			}
		}
		d.prog = append(d.prog, uint32(DFAOpStop<<DFARuneBits))
	}
	d.id2threadMask = threadMasks.reg
	for state, links := range linker {
		pc := state2pc[state]
		for _, ref := range links {
			d.prog[ref] = uint32(pc)
		}
	}
	re.dfa = d
}

func (s runeSlice) isDisjoint(t runeSlice) (r bool) {
	var a, b int
	for a < len(s) && b < len(t) {
		aFirst, aLast := s[a], s[a+1]
		bFirst, bLast := t[b], t[b+1]
		switch {
		case aLast < bFirst:
			// a
			//  b
			a += 2
		case aFirst > bLast:
			//  a
			// b
			b += 2
		default: // aFirst <= bLast && aLast >= bFirst
			// a
			// b
			return false
		}
	}
	return true
}

// split compares 's' and 't' and returns sets of items appearing only in 's',
// in both 's' and 't' and only in 't'. Any of the results can be nil if it's an
// empty set. None of the results share memory with any other slice, namely 's'
// or 't'.
func (s runeSlice) split(t runeSlice) (inS, inBoth, inT runeSlice) {
	s = append([]rune(nil), s...)
	t = append([]rune(nil), t...)
	if len(s) == 0 {
		return nil, nil, t
	}
	if len(t) == 0 {
		return s, nil, nil
	}

	var si, ti int
	for si < len(s) && ti < len(t) {
		sFirst, sLast := s[si], s[si+1]
		tFirst, tLast := t[ti], t[ti+1]
		switch {
		case sFirst < tFirst:
			switch {
			case sLast < tFirst:
				//  s:	<--->
				//  t:	          <--->
				inS = append(inS, sFirst, sLast)
				si += 2
			case sLast == tFirst:
				//  s:	<--->
				//  t:	    <--->
				inS = append(inS, sFirst, sLast-1)
				inBoth = append(inBoth, sLast, sLast)
				si += 2
				if tFirst == tLast {
					ti += 2
					break
				}

				t[ti] = tFirst + 1
			case sLast < tLast:
				//  s:	<----->
				//  t:	    <--->
				inS = append(inS, sFirst, tFirst-1)
				inBoth = append(inBoth, tFirst, sLast)
				si += 2
				t[ti] = sLast + 1
			case sLast == tLast:
				//  s:	<------->
				//  t:	    <--->
				inS = append(inS, sFirst, tFirst-1)
				inBoth = append(inBoth, tFirst, tLast)
				si += 2
				ti += 2
			case sLast > tLast:
				//  s:	<---------->
				//  t:	    <--->
				inS = append(inS, sFirst, tFirst-1)
				inBoth = append(inBoth, tFirst, tLast)
				s[si] = tLast + 1
				ti += 2
			default:
				panic("internal error")
			}
		case sFirst == tFirst:
			switch {
			case sLast == tFirst:
				//  s:	<>
				//  t:  <--->
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				if tFirst == tLast {
					ti += 2
					break
				}

				t[ti] = tFirst + 1
			case sLast < tLast:
				//  s:	<--->
				//  t:  <------>
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				t[ti] = sLast + 1
			case sLast == tLast:
				//  s:	<--->
				//  t:  <--->
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				ti += 2
			case sLast > tLast:
				//  s:	<------>
				//  t:  <--->
				inBoth = append(inBoth, sFirst, tLast)
				s[si] = tLast + 1
				ti += 2
			default:
				panic("internal error")
			}
		case sFirst < tLast:
			switch {
			case sLast < tLast:
				//  s:	  <--->
				//  t:  <-------->
				inT = append(inT, tFirst, sFirst-1)
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				t[ti] = sLast + 1
			case sLast == tLast:
				//  s:	  <--->
				//  t:  <----->
				inT = append(inT, tFirst, sFirst-1)
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				ti += 2
			case sLast > tLast:
				//  s:	  <------>
				//  t:  <----->
				inT = append(inT, tFirst, sFirst-1)
				inBoth = append(inBoth, sFirst, tLast)
				s[si] = tLast + 1
				ti += 2
			default:
				panic("internal error")
			}
		case sFirst == tLast:
			switch {
			case sLast == tLast:
				//  s:	    <>
				//  t:  <--->
				inBoth = append(inBoth, sFirst, sLast)
				si += 2
				if tFirst != tLast {
					inT = append(inT, tFirst, tLast-1)
				}
				ti += 2
			case sLast > tLast:
				//  s:	    <--->
				//  t:  <--->
				inT = append(inT, tFirst, tLast-1)
				inBoth = append(inBoth, sFirst, sFirst)
				s[si] = tLast + 1
				ti += 2
			default:
				panic("internal error")
			}
		case sFirst > tLast:
			switch {
			case sLast > tLast:
				//  s:	        <--->
				//  t:  <--->
				inT = append(inT, tFirst, tLast)
				ti += 2
			default:
				panic("internal error")
			}
		default:
			panic("internal error")
		}
	}
	if si < len(s) {
		inS = append(inS, s[si:]...)
	}
	if ti < len(t) {
		inT = append(inT, t[ti:]...)
	}
	return inS.normalize(), inBoth.normalize(), inT.normalize()
}

func (s runeSlice) normalize() (r runeSlice) {
	for i := 0; i < len(s)-2; i += 2 {
		if l, f := s[i+1], s[i+2]; l == f || l+1 == f {
			r = append(r, s[:i]...)
			s = s[i:]
			for len(s) >= 4 {
				switch l, f := s[1], s[2]; {
				case l == f, l+1 == f:
					s[2] = s[0]
				default:
					r = append(r, s[0], s[1])
				}
				s = s[2:]
			}
			return append(r, s...)
		}
	}
	return s
}

func (s *runeSlice) hashString() string {
	var b strings.Builder
	for _, v := range *s {
		for i := 0; i < 4; i++ {
			b.WriteByte(byte(v))
			v >>= 8
		}
	}
	return b.String()
}

func (s runeSlice) String() string {
	switch c := s.cardinality(); {
	case c <= 40:
		var a []string
		for i := 0; i < len(s); i += 2 {
			switch x, y := s[i], s[i+1]; {
			case x == y:
				a = append(a, fmt.Sprintf("%#U", x))
			default:
				a = append(a, fmt.Sprintf("%#U-%#U", x, y))
			}
		}
		return fmt.Sprintf("%v.%d", a, c)
	case unicode.MaxRune-c+1 <= 40:
		return "^" + s.neg().String()
	default:
		return fmt.Sprintf("[%d runes]", c)
	}
}

func (s runeSlice) neg() (r runeSlice) {
	if s[0] != 0 {
		r = append(r, 0, s[0]-1)
	}
	for i := 1; i < len(s)-1; i += 2 {
		r = append(r, s[i]+1, s[i+1]-1)
	}
	n := len(s) - 1
	if m := s[n]; m != unicode.MaxRune {
		r = append(r, m+1, unicode.MaxRune)
	}
	return r
}

func (s runeSlice) cardinality() (r int) {
	for i := 0; i < len(s); i += 2 {
		r += int(s[i+1]) - int(s[i]) + 1
	}
	return r
}

func (s runeSlice) expand() (r []rune) {
	for i := 0; i < len(s); i += 2 {
		for j := s[i]; j <= s[i+1]; j++ {
			r = append(r, j)
		}
	}
	return r
}

func (s runeSlice) has(c rune) (r bool) {
	//TODO binary search for len(s) > threshold
	for i := 0; i < len(s); i += 2 {
		if c >= s[i] && c <= s[i+1] {
			return true
		}
	}
	return false
}

func (re *Regexp) doDFA(ib []byte, is string, pos, ncap int, dstCap []int) []int {
	startCond := re.cond
	if startCond == ^syntax.EmptyOp(0) { // impossible
		return nil
	}

	if startCond&syntax.EmptyBeginText != 0 && pos != 0 {
		// Anchored match, past beginning of text.
		return nil
	}

	m := newOnePassMachine()

	defer freeOnePassMachine(m)

	if cap(m.matchcap) < ncap {
		m.matchcap = make([]int, ncap)
	} else {
		m.matchcap = m.matchcap[:ncap]
	}
	var matched bool
	for i := range m.matchcap {
		m.matchcap[i] = -1
	}
	i, _ := m.inputs.init(nil, ib, is)
	d := re.dfa
	prog := d.prog
	var pc, sym, pos0, width0 int
	var matchedThreads, deadThreads, runningThreads threadMask
	lastStatePC := -1
	lastStatePos := -1
	lastRestartPos := -1
restart:
	r, r1 := endOfText, endOfText
	width, width1 := 0, 0
	r, width = i.step(pos)
	if r != endOfText {
		r1, width1 = i.step(pos + width)
	}
	if pos == lastRestartPos {
		goto stop
	}

	lastRestartPos = pos
	matchedThreads = threadMask{}
	deadThreads = threadMask{}
	matched = false
	for i := 2; i < ncap; i++ {
		m.matchcap[i] = -1
	}
	pc = d.startPC
	if startCond&syntax.EmptyBeginText != 0 && pos != 0 {
		// Anchored match, past beginning of text.
		return nil
	}

	if len(re.prefix) > 0 && r1 != re.prefixRune && i.canCheckPrefix() {
		// Match requires literal prefix; fast search for it.
		advance := i.index(re, pos)
		if advance < 0 {
			goto stop
		}

		pos += advance
		r, width = i.step(pos)
		r1, width1 = i.step(pos + width)
	}
	pos0 = pos
	width0 = width
	if ncap > 0 {
		m.matchcap[0] = pos
	}
state:
	if pc == lastStatePC && pos == lastStatePos {
		goto stop
	}

	lastStatePC = pc
	lastStatePos = pos
	runningThreads = threadMask{}
instr:
	sym = int(prog[pc])
	pc++
	switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
	case DFAOpRune:
		next := int(prog[pc])
		pc++
		if r == rune(c) {
			pos += width
			r, width = r1, width1
			if r != endOfText {
				r1, width1 = i.step(pos + width)
			}
			pc = next
			goto state
		}

		goto instr
	case DFAOpEndText:
		next := int(prog[pc])
		pc++
		if r == endOfText {
			pc = next
			goto state
		}

		goto instr
	case DFAOpCharClass:
		next := int(prog[pc])
		pc++
		set := d.id2charClass[c]
		if set.has(r) {
			pos += width
			r, width = r1, width1
			if r != endOfText {
				r1, width1 = i.step(pos + width)
			}
			pc = next
			goto state
		}

		goto instr
	case opThreads:
		if !re.longest && ncap != 0 {
			runningThreads = d.id2threadMask[c]
			if matchedThreads.l.isNonZero() {
				g := matchedThreads.l
				g.and(runningThreads.l)
				if g != matchedThreads.l {
					goto stop
				}
			}

			runningThreads.andNot(deadThreads)
		}

		goto instr
	case opAcceptThreads:
		if !re.longest && ncap != 0 {
			runningThreads = d.id2threadMask[c]
			acceptMask := &d.id2threadMask[prog[pc]&DFARuneMask]
			pc++
			newKills := &d.id2threadMask[prog[pc]&DFARuneMask]
			pc++
			if matchedThreads.l.isNonZero() {
				g := matchedThreads.l
				g.and(runningThreads.l)
				g.andNot(newKills.r)
				e := matchedThreads.l
				e.andNot(newKills.r)
				if g != e {
					goto stop
				}
			}

			acceptMask.l.andNot(deadThreads.l)
			acceptMask.r.andNot(acceptMask.l)
			if matchedThreads.l.isNonZero() {
				g := matchedThreads.l
				g.and(acceptMask.l)
				g.andNot(newKills.r)
				e := matchedThreads.l
				e.andNot(newKills.r)
				if g != e {
					goto instr
				}
			}

			deadThreads.l.or(newKills.r)
			matchedThreads.l.andNot(deadThreads.l)
			matchedThreads.l.or(acceptMask.l)
			if acceptMask.isNonZero() {
				matched = true
				if ncap == 0 {
					goto stop
				}

				m.matchcap[1] = pos
				matchedThreads.l.or(acceptMask.l)
			}
			goto instr
		}

		pc += 2
		fallthrough
	case DFAOpAccept:
		matched = true
		if ncap == 0 {
			goto stop
		}

		m.matchcap[1] = pos
		goto instr
	case DFAOpBeginText:
		next := int(prog[pc])
		pc++
		if pos != pos0 {
			goto stop
		}

		pc = next
		goto state
	case DFAOpStop:
		if matched {
			goto stop
		}

		pos = pos0 + width0
		goto restart
	case opCapture:
		cap := c >> halfRuneBits
		if cap < ncap {
			switch {
			case re.longest || len(d.id2threadMask) == 0:
				m.matchcap[cap] = pos
			default:
				mask := d.id2threadMask[c&halfRuneMask]
				switch {
				case mask.isNonZero():
					mask.and(runningThreads)
					if mask.isNonZero() {
						m.matchcap[cap] = pos
					}
				default:
					m.matchcap[cap] = pos
				}
			}
		}
		if cap&1 != 0 {
			for _, v := range d.id2captureKills[cap] {
				if v < ncap {
					m.matchcap[v] = -1
				}
			}
		}
		goto instr
	default:
		panic("internal error")
	}

stop:
	if !matched {
		return nil
	}

	if ncap == 0 {
		return arrayNoInts[:0:0]
	}

	for i := 3; i < ncap; i += 2 {
		if m.matchcap[i] < 0 {
			m.matchcap[i-1] = -1
		}
	}
	return append(dstCap, m.matchcap...)
}
