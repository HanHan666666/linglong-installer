// Copyright 2023 The Regexp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp // modernc.org/regexp

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"modernc.org/fsm"
	"modernc.org/regexp/syntax"
)

const (
	smallClass = 64

	unicodePrivateFirst = 0xe000
	unicodePrivateLast  = 0xf8ff
)

type nfaLongest struct {
	*fsm.NFA
	charClasses register[string]
	err         error
	clsIndex    []*runeSlice
	sink        *fsm.State
	start       *fsm.State

	capturesAsGroups bool
	consumes         bool
	hasOpEndText     bool
}

func compileNFALongest(sre *syntax.Regexp, capturesAsGroups, postProcess bool) (r *nfaLongest, err error) {
	defer func() {
		if r != nil {
			if err == nil {
				err = r.err
			}
			return
		}

		if err == nil {
			err = fmt.Errorf("compile NFA longest failed")
		}
	}()

	r = &nfaLongest{
		NFA:              fsm.NewLimitedNFA(maxFSMStates),
		capturesAsGroups: capturesAsGroups,
	}
	if !r.canCompile(sre) {
		return nil, r.err
	}

	r.start = r.NewState()
	r.sink = r.NewState()
	if r.start == nil || r.sink == nil {
		return nil, r.err
	}

	out := r.add(r.start, sre)
	if out == nil {
		return nil, r.err
	}

	if r.hasOpEndText && !r.consumes {
		return nil, r.err
	}

	s := r.NewState()
	if s == nil {
		return nil, r.err
	}

	s.IsAccepting = true
	out.NewEdge(DFAOpAccept<<DFARuneBits+out.Id(), s)
	r.sink.IsAccepting = true
	for _, state := range r.List() {
		state.NewEdge(dfaOpState<<DFARuneBits+state.Id(), r.sink)
	}
	if postProcess && !r.post() {
		return nil, r.err
	}

	return r, nil
}

func (n *nfaLongest) post() bool {
	for _, state := range n.List() {
		edges := state.Transitions()
		syms := edges.List()
		type info struct {
			rune int
			next *fsm.State
		}
		var infos []info
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				for _, next := range nexts {
					infos = append(infos, info{c, next})
				}
			}
		}
		var charClasses []*runeSlice
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpCharClass:
				charClass := n.clsIndex[c]
				for _, v := range charClasses {
					if !charClass.isDisjoint(*v) {
						return false
					}
				}

				charClasses = append(charClasses, charClass)
				for _, info := range infos {
					c := info.rune
					if charClass.has(rune(c)) {
						for _, next := range nexts {
							state.NewEdge(DFAOpRune<<DFARuneBits+c, next)
						}
					}
				}
			}
		}
	}
	return true
}

func sortedSyms(edges fsm.Transitions) (syms []int) {
	syms = edges.List()
	sort.Slice(syms, func(i, j int) bool {
		a, b := syms[i]>>DFARuneBits, syms[j]>>DFARuneBits
		soa := opSortOrder[a]
		sob := opSortOrder[b]
		if soa < sob {
			return true
		}

		if soa > sob {
			return false
		}

		return syms[i]&DFARuneMask < syms[j]&DFARuneMask
	})
	return syms
}

type runeInfo struct {
	rune
	nexts []*fsm.State
}

type charClassInfo struct {
	id    int
	sym   int
	s     *runeSlice
	nexts []*fsm.State
}

//lint:ignore U1000 debug helper
func nfaString(n *fsm.NFA, clsIndex []*runeSlice) string {
	a := strings.Split(n.String(), "\n")
	w := 0
	for _, v := range a {
		v0 := v
		var c uint64
		if x := strings.Index(v, "(U+"); x >= 0 {
			s := v[x+3:]
			s = s[:strings.Index(s, ")")]
			c, _ = strconv.ParseUint(s, 16, 32)
			c = c & DFARuneMask
		}
		switch {
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*dfaOpState)):
			// skip
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpAccept)):
			a[w] = fmt.Sprintf("%s (accept)", v0)
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpCharClass)):
			a[w] = fmt.Sprintf("%s (class[%d]=%s)", v0, c, clsIndex[c])
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpEndText)):
			a[w] = fmt.Sprintf("%s (EOT)", v0)
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpBeginText)):
			a[w] = fmt.Sprintf("%s (begin text)", v0)
			w++
		default:
			a[w] = v0
			w++
		}
	}
	a = a[:w]
	for i, v := range clsIndex {
		a = append(a, fmt.Sprintf("char class %d: %v", i, v))
	}
	a = append(a, "")
	return strings.Join(a, "\n")
}

func optimizeDFA(n *fsm.NFA, clsReg *register[string], clsIndex *[]*runeSlice) (r *fsm.NFA) {
	const classify = 7
	for _, state := range n.List() {
		edges := state.Transitions()
		syms := edges.List()
		m := map[*fsm.State][]int{} // state: sym
		for _, sym := range syms {
			if sym == fsm.Epsilon {
				panic("internal error, not a DFA")
			}

			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				targets := edges.Get(c).List()
				if len(targets) != 1 {
					return nil
				}

				target := targets[0]
				m[target] = append(m[target], sym)
			}
		}
		for next, syms := range m {
			if len(syms) < classify {
				continue
			}

			var rs runeSlice
			for _, r := range syms {
				c := rune(r & DFARuneMask)
				rs = append(rs, c, c)
			}
			rs = rs.normalize()
			id := registerCharClass(clsReg, clsIndex, &rs)
			for _, sym := range syms {
				edges.Delete(sym)
			}
			state.NewEdge(DFAOpCharClass<<DFARuneBits+id, next)
		}
	}
	return n
}

func resolveCharClasses2(n *fsm.NFA, clsReg *register[string], clsIndex *[]*runeSlice) (r *fsm.NFA, modified bool) {
	for _, state := range n.List() {
	restart:
		edges := state.Transitions()
		syms := edges.List()
		var runeInfos []runeInfo
		for _, sym := range syms {
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				if c >= unicodePrivateFirst && c <= unicodePrivateLast {
					break
				}

				runeInfos = append(runeInfos, runeInfo{rune(c), edges.Get(c).List()})
			}
		}
		var charClassInfos []charClassInfo
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			switch op, sID := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpCharClass:
				cls := (*clsIndex)[sID]
				if cls.cardinality() <= smallClass {
					nexts := edges.Get(sym).List()
					state.Transitions().Delete(sym)
					runes := cls.expand()
					for _, next := range nexts {
						for _, v := range runes {
							state.NewEdge(DFAOpRune<<DFARuneBits+int(v), next)
						}
					}
					modified = true
					goto restart
				}

				si := charClassInfo{sID, sym, cls, nexts}
				if len(charClassInfos) == 0 {
					charClassInfos = append(charClassInfos, si)
					break
				}

				// resolve any splits
				for _, ti := range charClassInfos {
					s0 := *si.s
					t0 := *ti.s
					switch s, b, t := s0.split(t0); {
					case len(b) == 0: // s0 and t0 are disjoint
						charClassInfos = append(charClassInfos, si)
					default:
						// s0 ['a', 'b']	['a']		['a', 'b']
						// t0 ['b', 'c']	['a', 'b']	['b']
						//
						// s  ['a']		[]		['a']
						// b  ['b']		['a']		['b']
						// t  ['c']		['b']		[]
						bID := registerCharClass(clsReg, clsIndex, &b)
						sNexts := edges.Get(si.sym).List()
						tNexts := edges.Get(ti.sym).List()
						state.Transitions().Delete(si.sym)
						state.Transitions().Delete(ti.sym)
						if len(s) != 0 {
							sID := registerCharClass(clsReg, clsIndex, &s)
							for _, next := range sNexts {
								state.NewEdge(DFAOpCharClass<<DFARuneBits+sID, next)
							}
						}
						for _, next := range sNexts {
							state.NewEdge(DFAOpCharClass<<DFARuneBits+bID, next)
						}
						for _, next := range tNexts {
							state.NewEdge(DFAOpCharClass<<DFARuneBits+bID, next)
						}
						if len(t) != 0 {
							tID := registerCharClass(clsReg, clsIndex, &t)
							for _, next := range tNexts {
								state.NewEdge(DFAOpCharClass<<DFARuneBits+tID, next)
							}

						}
						// trc("---- state %d", state.Id())
						// trc("s0=%s", s0)
						// trc("t0=%s", t0)
						// trc("s=%s", s)
						// trc("b=%s", b)
						// trc("t=%s", t)
						modified = true
						goto restart
					}
				}
			}
		}
		for _, runeInfo := range runeInfos {
			for _, charClassInfo := range charClassInfos {
				cls := *charClassInfo.s
				r := runeInfo.rune
				if cls.has(r) {
					sym := DFAOpRune<<DFARuneBits + int(r)
					for _, next := range charClassInfo.nexts {
						closure := edges.Get(sym)
						if !closure.Has(next) {
							modified = true
							state.NewEdge(sym, next)
						}
					}
				}
			}
		}
	}
	return n, modified
}

func resolveCharClasses(n *fsm.NFA, clsReg *register[string], clsIndex *[]*runeSlice) (r *fsm.NFA) {
	for i := 0; i < 100; i++ {
		s := n.String()
		m, modified := resolveCharClasses2(n, clsReg, clsIndex)
		if !modified {
			return n
		}

		n2 := m.MinimalDFA(false)
		if n2 == nil {
			return nil
		}

		if s == n2.String() {
			return n2
		}

		n = n2
	}
	return nil
}

func (n *nfaLongest) canCompile(sre *syntax.Regexp) bool {
	switch sre.Op {
	case syntax.OpNoMatch:
		return false
	case syntax.OpEmptyMatch:
		return true
	case syntax.OpLiteral:
		n.consumes = true
		return true
	case syntax.OpCharClass:
		if len(sre.Rune) != 0 {
			n.consumes = true
		}
		return true
	case syntax.OpAnyCharNotNL:
		n.consumes = true
		return true
	case syntax.OpAnyChar:
		n.consumes = true
		return true
	case syntax.OpBeginLine:
		return false
	case syntax.OpEndLine:
		return false
	case syntax.OpBeginText:
		return true
	case syntax.OpEndText:
		n.hasOpEndText = true
		return true
	case syntax.OpWordBoundary:
		return false
	case syntax.OpNoWordBoundary:
		return false
	case syntax.OpCapture:
		if !n.capturesAsGroups {
			return false
		}

		for _, v := range sre.Sub {
			if !n.canCompile(v) {
				return false
			}
		}

		return true
	case syntax.OpStar:
		return n.canCompile(sre.Sub[0])
	case syntax.OpPlus:
		return n.canCompile(sre.Sub[0])
	case syntax.OpQuest:
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
			if !n.canCompile(v) {
				return false
			}
		}

		return true
	default:
		return false
	}
}

func (n *nfaLongest) add(in *fsm.State, sre *syntax.Regexp) (out *fsm.State) {
	switch sre.Op {
	case syntax.OpLiteral:
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
		return out
	case syntax.OpConcat:
		for _, v := range sre.Sub {
			if out = n.add(in, v); out == nil {
				return nil
			}

			in = out
		}
		return out
	case syntax.OpCapture:
		if !n.capturesAsGroups {
			return nil
		}

		for _, v := range sre.Sub {
			if out = n.add(in, v); out == nil {
				return nil
			}

			in = out
		}
		return out
	case syntax.OpAlternate:
		for _, v := range sre.Sub {
			o := n.add(in, v)
			if o == nil {
				return nil
			}

			switch {
			case out == nil:
				out = o
			default:
				o.NewEdge(fsm.Epsilon, out)
			}
		}
		return out
	case syntax.OpQuest:
		if out = n.add(in, sre.Sub[0]); out == nil {
			return nil
		}

		in.NewEdge(fsm.Epsilon, out)
		return out
	case syntax.OpPlus:
		if n.hasEOT(sre.Sub[0]) {
			return nil
		}

		s := n.NewState()
		if s == nil {
			return nil
		}

		in.NewEdge(fsm.Epsilon, s)
		if out = n.add(s, sre.Sub[0]); out == nil {
			return nil
		}

		out.NewEdge(fsm.Epsilon, s)
		if s = n.NewState(); s == nil {
			return nil
		}

		out.NewEdge(fsm.Epsilon, s)
		return s
	case syntax.OpStar:
		if n.hasEOT(sre.Sub[0]) {
			return nil
		}

		s := n.NewState()
		if s == nil {
			return nil
		}

		in.NewEdge(fsm.Epsilon, s)
		if out = n.add(s, sre.Sub[0]); out == nil {
			return nil
		}

		out.NewEdge(fsm.Epsilon, s)
		s.NewEdge(fsm.Epsilon, out)
		if s = n.NewState(); s == nil {
			return nil
		}

		out.NewEdge(fsm.Epsilon, s)
		return s
	case syntax.OpBeginText:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(DFAOpBeginText<<DFARuneBits, out)
		return out
	case syntax.OpCharClass:
		return n.charClass(in, (*runeSlice)(&sre.Rune))
	case syntax.OpAnyCharNotNL:
		return n.charClass(in, (*runeSlice)(&anyRuneNotNL))
	case syntax.OpAnyChar:
		return n.charClass(in, (*runeSlice)(&anyRune))
	case syntax.OpEndText:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(DFAOpEndText<<DFARuneBits, out)
		return out
	case syntax.OpEmptyMatch:
		if out = n.NewState(); out == nil {
			return nil
		}

		in.NewEdge(fsm.Epsilon, out)
		return out
	default:
		return nil
	}
	panic("unreachable")
}

func (n *nfaLongest) hasEOT(sre *syntax.Regexp) bool {
	switch sre.Op {
	case syntax.OpEndText:
		return true
	case syntax.OpConcat:
		for _, v := range sre.Sub {
			if n.hasEOT(v) {
				return true
			}
		}

		return false
	case syntax.OpCapture:
		for _, v := range sre.Sub {
			if n.hasEOT(v) {
				return true
			}
		}

		return false
	case syntax.OpAlternate:
		for _, v := range sre.Sub {
			if n.hasEOT(v) {
				return true
			}
		}

		return false
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest:
		return n.hasEOT(sre.Sub[0])
	default:
		return false
	}
}

func (n *nfaLongest) charClass(in *fsm.State, p *runeSlice) (out *fsm.State) {
	if out = n.NewState(); out == nil {
		return nil
	}

	id := registerCharClass(&n.charClasses, &n.clsIndex, p)
	in.NewEdge(DFAOpCharClass<<DFARuneBits+id, out)
	return out
}

func registerCharClass(reg *register[string], index *[]*runeSlice, charClass *runeSlice) (id int) {
	if id = reg.id(charClass.hashString()); id == len(*index) {
		*index = append(*index, charClass)
	}
	return id
}

// DFA represents a regexp compiled to a DFA. Only a subset of regexps is
// supported.
type DFA struct {
	cond            syntax.EmptyOp
	dfaPC2nfaStates map[int][]int
	err             error
	id2charClass    []*runeSlice
	minInputLen     int
	nfa             *nfaLongest
	prefix          string
	prefixRune      rune
	prog            []uint32

	startPC  int
	prefixPC int
}

// CharClasses returns the number of distinc character classes used by d.
func (d *DFA) CharClasses() int { return len(d.id2charClass) }

// CharClass returns a character class for id which must come from the
// DFAOpCharClass instruction argument of d.Prog. The returned slice consist of
// pairs describing ranges of members of the set.
func (d *DFA) CharClass(id int) (r []rune) {
	return append(r, []rune(*d.id2charClass[id])...)
}

// MinInputLen reports the minimum length of the input in bytes for a match.
func (d *DFA) MinInputLen() int { return d.minInputLen }

// Prefix reports the required prefix in unanchored matches, if any.
func (d *DFA) Prefix() string { return d.prefix }

// PrefixRune reports the first rune in prefix.
func (d *DFA) PrefixRune() rune { return d.prefixRune }

// Cond reports the empty-width conditions required at start of match.
func (d *DFA) Cond() syntax.EmptyOp { return d.cond }

// Prog returns the DFA program.
func (d *DFA) Prog() []uint32 { return d.prog }

// StartPC reports the PC to start the DFA.
func (d *DFA) StartPC() int { return d.startPC }

// PrefixPC reports the PC to start the DFA after the prefix is consumed.
func (d *DFA) PrefixPC() int { return d.prefixPC }

// NewLexDFA returns a DFA suitable for a lexer or an error, if any.
func NewLexDFA(expr []string) (r *DFA, err error) {
	var b strings.Builder
	for i, v := range expr {
		if i != 0 {
			b.WriteByte('|')
		}
		fmt.Fprintf(&b, "(%s)\\x{%04x}", v, unicodePrivateFirst+i)
	}
	sre, err := syntax.Parse(b.String(), syntax.Perl|syntax.DisableOptimizations)
	if err != nil {
		return nil, err
	}

	// sre = sre.Simplify()
	prog, err := syntax.Compile(sre)
	if err != nil {
		return nil, err
	}

	nfa, err := compileNFALongest(sre, true, false)
	if err != nil {
		return nil, err
	}

	if nfa.NFA, _ = resolveCharClasses2(nfa.NFA, &nfa.charClasses, &nfa.clsIndex); nfa.NFA == nil {
		return nil, fmt.Errorf("compile NFA longest failed")
	}

	defer func() {
		if r != nil {
			if err == nil {
				err = r.err
			}
			return
		}

		if err == nil {
			err = fmt.Errorf("compile DFA longest failed")
		}
	}()

	dfa := nfa.MinimalDFA(false)
	if dfa == nil {
		return nil, nil
	}

	if dfa = resolveCharClasses(dfa, &nfa.charClasses, &nfa.clsIndex); dfa == nil {
		return nil, nil
	}

	prefix, _ := prog.Prefix()
	d := &DFA{
		cond:         prog.StartCond(),
		id2charClass: nfa.clsIndex,
		minInputLen:  minInputLen(sre),
		nfa:          nfa,
		prefix:       prefix,
		prefixPC:     -1,
	}
	if prefix != "" {
		d.prefixRune, _ = utf8.DecodeRuneInString(prefix)
	}
	for _, state := range dfa.List() {
		edges := state.Transitions()
		syms := sortedSyms(edges)
		for _, sym := range syms {
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				if c >= unicodePrivateFirst && c <= unicodePrivateLast {
					edges.Delete(sym)
					state.NewEdge(DFAOpAccept<<DFARuneBits|(c-unicodePrivateFirst), state)
				}
			}
		}
	}
	if dfa = dfa.MinimalDFA(false); dfa == nil {
		return nil, nil
	}

	if dfa = optimizeDFA(dfa, &nfa.charClasses, &nfa.clsIndex); dfa == nil {
		return nil, nil
	}

	d.id2charClass = nfa.clsIndex
	state2pc := []int{}
	states := dfa.List()
	linker := make([][]int, len(states))
	dfaPC2nfaStates := map[int][]int{}
	d.dfaPC2nfaStates = dfaPC2nfaStates
	for _, state := range states {
		state2pc = append(state2pc, len(d.prog))
		if state.IsAccepting {
			continue
		}

		edges := state.Transitions()
		syms := sortedSyms(edges)
		for _, sym := range syms {
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case dfaOpState:
				dfaPC2nfaStates[len(d.prog)] = append(dfaPC2nfaStates[len(d.prog)], c)
			}
		}
		instrCnt := 0
		hasOpBeginText := false
		var charClasses []*runeSlice
		accepted := false
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			if len(nexts) > 1 {
				return
			}

			next := nexts[0].Id()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				d.prog = append(d.prog, uint32(DFAOpRune<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpEndText:
				d.prog = append(d.prog, uint32(DFAOpEndText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpCharClass:
				charClass := d.id2charClass[c]
				for _, v := range charClasses {
					if !charClass.isDisjoint(*v) {
						return
					}
				}

				charClasses = append(charClasses, charClass)
				d.prog = append(d.prog, uint32(DFAOpCharClass<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpBeginText:
				if state != dfa.Start() {
					return
				}

				d.prog = append(d.prog, uint32(DFAOpBeginText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
				hasOpBeginText = true
			case DFAOpAccept:
				if !accepted {
					d.prog = append(d.prog, uint32(sym))
					accepted = true
				}
			case dfaOpState:
				// nop
			default:
				return
			}
		}
		if hasOpBeginText && instrCnt != 1 {
			return
		}

		d.prog = append(d.prog, uint32(DFAOpStop<<DFARuneBits))
	}
	for _, v := range dfaPC2nfaStates {
		sort.Ints(v)
	}
	for state, links := range linker {
		pc := state2pc[state]
		for _, ref := range links {
			d.prog[ref] = uint32(pc)
		}
	}
	return d, nil
}

// NewMatchDFA returns a DFA suitable for the Match operation or an error, if
// any.
func NewMatchDFA(expr string) (r *DFA, err error) {
	sre, err := syntax.Parse(expr, syntax.Perl)
	if err != nil {
		return nil, err
	}

	sre = sre.Simplify()
	prog, err := syntax.Compile(sre)
	if err != nil {
		return nil, err
	}

	nfa, err := compileNFALongest(sre, true, true)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r != nil {
			if err == nil {
				err = r.err
			}
			return
		}

		if err == nil {
			err = fmt.Errorf("compile DFA longest failed")
		}
	}()

	dfa := nfa.MinimalDFA(false)
	if dfa == nil {
		return nil, nil
	}

	prefix, _ := prog.Prefix()
	d := &DFA{
		cond:         prog.StartCond(),
		id2charClass: nfa.clsIndex,
		minInputLen:  minInputLen(sre),
		nfa:          nfa,
		prefix:       prefix,
		prefixPC:     -1,
	}
	if prefix != "" {
		d.prefixRune, _ = utf8.DecodeRuneInString(prefix)
	}
	states := dfa.List()
	rebuild := false
	for _, state := range states {
		edges := state.Transitions()
		syms := sortedSyms(edges)
		type info struct {
			rune int
			next *fsm.State
		}
		var infos []info
		hasCharClass := false
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				for _, next := range nexts {
					infos = append(infos, info{c, next})
				}
			case DFAOpCharClass:
				if hasCharClass {
					return nil, nil
				}

				hasCharClass = true
				charClass := d.id2charClass[c]
				for _, info := range infos {
					c := info.rune
					if charClass.has(rune(c)) {
						rebuild = true
						for _, next := range nexts {
							state.NewEdge(DFAOpRune<<DFARuneBits+c, next)
						}
					}
				}
			}
		}
	}
	if rebuild {
		if dfa = dfa.MinimalDFA(false); dfa == nil {
			return nil, nil
		}
	}

	state2pc := []int{}
	states = dfa.List()
	linker := make([][]int, len(states))
	dfaPC2nfaStates := map[int][]int{}
	d.dfaPC2nfaStates = dfaPC2nfaStates
	for _, state := range states {
		state2pc = append(state2pc, len(d.prog))
		if state.IsAccepting {
			continue
		}

		edges := state.Transitions()
		syms := sortedSyms(edges)
		hasCharClass := false
		isAccepting := false
		for _, sym := range syms {
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case dfaOpState:
				dfaPC2nfaStates[len(d.prog)] = append(dfaPC2nfaStates[len(d.prog)], c)
			case DFAOpAccept:
				isAccepting = true
			}
		}
		if isAccepting {
			d.prog = append(d.prog, uint32(DFAOpAccept<<DFARuneBits))
		}
		instrCnt := 0
		hasOpBeginText := false
		for _, sym := range syms {
			nexts := edges.Get(sym).List()
			if len(nexts) > 1 {
				return
			}

			next := nexts[0].Id()
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				d.prog = append(d.prog, uint32(DFAOpRune<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpEndText:
				d.prog = append(d.prog, uint32(DFAOpEndText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpCharClass:
				if hasCharClass {
					return
				}

				hasCharClass = true
				d.prog = append(d.prog, uint32(DFAOpCharClass<<DFARuneBits+c), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
			case DFAOpBeginText:
				if state != dfa.Start() {
					return
				}

				d.prog = append(d.prog, uint32(DFAOpBeginText<<DFARuneBits), 0)
				linker[next] = append(linker[next], len(d.prog)-1)
				instrCnt++
				hasOpBeginText = true
			case DFAOpAccept, dfaOpState:
				// nop
			default:
				return
			}
		}
		if hasOpBeginText && instrCnt != 1 {
			return
		}

		d.prog = append(d.prog, uint32(DFAOpStop<<DFARuneBits))
	}
	for _, v := range dfaPC2nfaStates {
		sort.Ints(v)
	}
	for state, links := range linker {
		pc := state2pc[state]
		for _, ref := range links {
			d.prog[ref] = uint32(pc)
		}
	}
	if !d.post() {
		return nil, nil
	}

	return d, nil
}

func (d *DFA) post() bool {
	s := d.prefix
	if s == "" {
		return true
	}

	pc := d.startPC
state:
	for len(s) != 0 {
		r, n := utf8.DecodeRuneInString(s)
		if n == 0 {
			return false
		}

		s = s[n:]
	instr:
		for {
			pc0 := pc
			sym := d.prog[pc]
			pc++
			switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
			case DFAOpRune:
				next := d.prog[pc]
				pc++
				if r == rune(c) {
					pc = int(next)
					continue state
				}
			case DFAOpBeginText:
				if pc0 != d.startPC {
					return false
				}

				pc = int(d.prog[pc])
				continue instr
			default:
				return false
			}
		}
	}
	d.prefixPC = pc
	return true
}

func (n *nfaLongest) String() string {
	a := strings.Split(n.NFA.String(), "\n")
	w := 0
	for _, v := range a {
		v0 := v
		var c uint64
		if x := strings.Index(v, "(U+"); x >= 0 {
			s := v[x+3:]
			s = s[:strings.Index(s, ")")]
			c, _ = strconv.ParseUint(s, 16, 32)
			c = c & DFARuneMask
		}
		switch {
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*dfaOpState)):
			// skip
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpAccept)):
			a[w] = fmt.Sprintf("%s (accept)", v0)
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpCharClass)):
			a[w] = fmt.Sprintf("%s (%s)", v0, n.clsIndex[c])
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpEndText)):
			a[w] = fmt.Sprintf("%s (EOT)", v0)
			w++
		case strings.Contains(v0, fmt.Sprintf("(U+%X", 2*DFAOpBeginText)):
			a[w] = fmt.Sprintf("%s (begin text)", v0)
			w++
		default:
			a[w] = v0
			w++
		}
	}
	a = a[:w]
	return strings.Join(a, "\n")
}

func (d *DFA) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "NFA\n%s\n", d.nfa)
	prog := d.prog
	fmt.Fprintf(&b, "DFA prog: %4d, start %05d", len(prog), d.startPC)
	if d.prefixPC > 0 {
		fmt.Fprintf(&b, ", prefix start %05d", d.prefixPC)
	}
	fmt.Fprintln(&b)
	state := 0
	newState := true
	for pc := 0; pc < len(prog); {
		if newState {
			fmt.Fprintf(&b, "[%d] nfa states %v\n", state, d.dfaPC2nfaStates[pc])
			state++
			newState = false
		}
		sym := prog[pc]
		fmt.Fprintf(&b, "\t%05d:\t%#U\t", pc, sym)
		pc++
		switch op, c := sym>>DFARuneBits, sym&DFARuneMask; op {
		case DFAOpRune:
			next := prog[pc]
			pc++
			fmt.Fprintf(&b, "%s -> %05d\n", strconv.QuoteRuneToASCII(rune(c)), next)
		case DFAOpEndText:
			next := prog[pc]
			pc++
			fmt.Fprintf(&b, "EOT -> %05d\n", next)
		case DFAOpCharClass:
			next := prog[pc]
			pc++
			cls := d.id2charClass[c]
			fmt.Fprintf(&b, "%s -> %05d\n", cls, next)
		case DFAOpAccept:
			fmt.Fprintf(&b, "accept\n")
		case DFAOpStop:
			fmt.Fprintf(&b, "stop\n")
			newState = true
		case DFAOpBeginText:
			next := prog[pc]
			pc++
			fmt.Fprintf(&b, "begin text -> %05d\n", next)
		default:
			fmt.Fprintf(&b, "???\n")
		}
	}
	return b.String()
}

// OnePassProg represents a one pass program.
//
// "One-pass" regexp execution.
//
// Some regexps can be analyzed to determine that they never need backtracking:
// they are guaranteed to run in one pass over the string without bothering to
// save all the usual NFA state.
type OnePassProg struct {
	Prog        *syntax.Prog
	numSubexp   int
	prefix      string
	subexpNames []string

	prefixComplete bool // prefix is the entire regexp
}

func (p *OnePassProg) String() string { return p.Prog.String() }

// NumSubexp returns the number of parenthesized subexpressions in this
// OnePassProg.
func (p *OnePassProg) NumSubexp() int {
	return p.numSubexp
}

// SubexpNames returns the names of the parenthesized subexpressions in this
// OnePassProg. The name for the first sub-expression is names[1], so that if m
// is a match slice, the name for m[i] is SubexpNames()[i].  Since the
// OnePassProg as a whole cannot be named, names[0] is always the empty string.
func (p *OnePassProg) SubexpNames() []string {
	return append([]string(nil), p.subexpNames...)
}

// LiteralPrefix returns a literal string that must begin any match
// of the regular expression p. It returns the boolean true if the
// literal string comprises the entire regular expression.
func (p *OnePassProg) LiteralPrefix() (prefix string, complete bool) {
	return p.prefix, p.prefixComplete
}

// OnePassProg returns a new OnePassProg or an error, if any.
func NewOnePassProg(expr string, mode syntax.Flags) (*OnePassProg, error) {
	re, err := syntax.Parse(expr, mode)
	if err != nil {
		return nil, err
	}

	maxCap := re.MaxCap()
	capNames := re.CapNames()

	re = re.Simplify()
	prog, err := syntax.Compile(re)
	if err != nil {
		return nil, err
	}

	onepass := compileOnePass(prog)
	if onepass == nil {
		return nil, fmt.Errorf("cannot handle `%s`", expr)
	}

	prefix, prefixComplete, _ := onePassPrefix(prog)
	return &OnePassProg{
		Prog:           prog,
		numSubexp:      maxCap,
		subexpNames:    capNames,
		prefix:         prefix,
		prefixComplete: prefixComplete,
	}, nil
}
