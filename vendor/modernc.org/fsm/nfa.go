// Copyright 2014 The fsm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package fsm provides some utilities for dealing with finite state machines.

# Dead DFA State

"A state is a dead state if it is not an accepting state and has no out-going
transitions except to itself."[7]

Passing withDeadState == true to some methods in this package makes the
produced DFAs "complete". For many practical purposes the dead state is not
needed and all the additional edges to it are only a waste of memory.

Note: Negative symbol values are reserved for internal purposes.

# References

Linked from elsewhere:

	[1]: http://en.wikipedia.org/wiki/Finite-state_machine
	[2]: http://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
	[3]: http://en.wikipedia.org/wiki/Powerset_construction
	[4]: http://en.wikipedia.org/wiki/Nondeterministic_finite_automaton_with_%CE%B5-moves
	[5]: http://en.wikipedia.org/wiki/DFA_minimization
	[6]: http://en.wikipedia.org/wiki/Janusz_Brzozowski_%28computer_scientist%29
	[7]: http://www.cs.odu.edu/~toida/nerzic/390teched/regular/fa/min-fa.html
	[8]: http://godoc.org/modernc.org/golex
	[9]: http://en.wikipedia.org/wiki/Powerset_construction#Complexity
*/
package fsm // import "modernc.org/fsm"

import (
	"bytes"
	"fmt"
	"sort"

	"modernc.org/mathutil"
	"modernc.org/strutil"
)

// Epsilon is a symbol value representing an ε edge.
const Epsilon = -1

// -------------------------------------------------------------------- Closure

// Closure is a set of states.
type Closure struct{ closure }

// NewClosure returns a newly created Closure.
func NewClosure() Closure { return Closure{closure{}} }

type closure map[*State]struct{}

func (c closure) id() string {
	a := make([]int, 0, len(c))
	for s := range c {
		a = append(a, s.Id())
	}
	sort.Ints(a)
	return fmt.Sprint(a)
}

// Exclude removes state s from the closure.
func (c closure) Exclude(s *State) { delete(c, s) }

// Has returns whether s is in the closure.
func (c closure) Has(s *State) (ok bool) { _, ok = c[s]; return ok }

// Include adds s to the closure.
func (c closure) Include(s *State) { c[s] = struct{}{} }

// Len returns the number of states in the closure.
func (c closure) Len() int { return len(c) }

// List returns a sorted slice of all states in c.
func (c closure) List() (r []*State) {
	r = make([]*State, len(c))
	i := 0
	for state := range c {
		r[i] = state
		i++
	}
	sort.Slice(r, func(i, j int) bool { return r[i].Id() < r[j].Id() })
	return
}

// ------------------------------------------------------------------------ NFA

// NFA is a nondeterministic finite automaton [2].
type NFA struct {
	max    int
	budget int
	s2i    map[*State]int
	i2s    map[int]*State
	start  *State
}

// NewNFA returns a new, empty NFA.
func NewNFA() *NFA { return NewLimitedNFA(-1) }

// NewLimitedNFA returns a new, empty NFA with a limit 'max' on number of
// states. The limit is also applied to objects derived via PowerSet,
// MinimalDFA etc. Passing a negative value for 'max' makes the NFA limited
// only by process resources.
func NewLimitedNFA(max int) *NFA {
	return &NFA{
		budget: max,
		i2s:    map[int]*State{},
		max:    max,
		s2i:    map[*State]int{},
	}
}

func (n *NFA) id(s *State) int {
	if id, ok := n.s2i[s]; ok {
		return id
	}

	i := n.Len()
	n.s2i[s] = i
	n.i2s[i] = s
	return i
}

// Identical reports whether n and m represent the same graph, having the same
// number of equally numbered states, each state has the same set of equally
// labeled edges going to equally numbered states and that n.Start().Id() ==
// m.Start().Id() and the accepting states of n and m are the same.
func (n *NFA) Identical(m *NFA) bool {
	if len(n.s2i) != len(m.s2i) || n.Start().Id() != m.Start().Id() {
		return false
	}

	for id, s := range n.i2s {
		s2 := m.i2s[id]
		if s2 == nil || s.IsAccepting != s2.IsAccepting {
			return false
		}
		t := s.Transitions()
		t2 := s2.Transitions()
		edges := t.List()
		edges2 := t2.List()
		if len(edges) != len(edges2) {
			return false
		}

		for i, v := range edges {
			if edges2[i] != v {
				return false
			}

			nexts := t.Get(v).List()
			nexts2 := t2.Get(v).List()
			if len(nexts) != len(nexts2) {
				return false
			}

			for i, next := range nexts {
				if next.Id() != nexts2[i].Id() {
					return false
				}
			}
		}
	}
	return true
}

// Equals returns wheter m accepts the same language as n.
func (n *NFA) Equals(m *NFA) bool { return n.MinimalDFA(false).equals(m.MinimalDFA(false)) }

// Equals returns wheter a accepts the same language as b. Both a and b must be
// minimal DFAs and both must have been created using the same value of 'v' in
// MinimalDFA(v).
func (n *NFA) equals(b *NFA) bool {
	a := n
	nstates := a.Len()
	if b.Len() != nstates { // must have same # of states
		return false
	}

	x := make(map[int]int, nstates) // a.id -> b.id
	visited := make(map[int]bool, nstates)
	var f func(*State, *State) bool

	f = func(sa, sb *State) bool {
		ida := sa.Id()
		if visited[ida] {
			return true
		}

		visited[ida] = true
		ta, tb := sa.Transitions(), sb.Transitions()
		if ta.Len() != tb.Len() { // must have same # of edges
			return false
		}

		pairs := []struct{ a, b *State }{}
		for _, sym := range ta.List() {
			ca := ta.Get(sym)
			cb := tb.Get(sym)
			targa, targb := ca.List(), cb.List()
			nexta, nextb := targa[0], targb[0]
			nida, nidb := nexta.Id(), nextb.Id()
			if v, ok := x[nida]; ok {
				if v != nidb {
					return false
				}

				continue
			}

			x[nida] = nidb
			pairs = append(pairs, struct{ a, b *State }{nexta, nextb})
		}

		for _, pair := range pairs {
			if !f(pair.a, pair.b) {
				return false
			}
		}

		return true
	}

	return f(a.Start(), b.Start())
}

// Len returns the number of NFA's states.
func (n *NFA) Len() int { return len(n.s2i) }

// List returns a sorted slice of n's states.
func (n *NFA) List() (r []*State) {
	r = make([]*State, n.Len())
	for i, state := range n.i2s {
		r[i] = state
	}
	sort.Slice(r, func(i, j int) bool { return r[i].Id() < r[j].Id() })
	return
}

// MinimalDFA returns the NFA converted to a minimal DFA[5]. Dead state is
// possibly constructed if withDeadState == true.
//
// Note: Algorithm used is Brzozowski[6].
func (n *NFA) MinimalDFA(withDeadState bool) *NFA {
	r := n.Reverse()
	if r == nil {
		return nil
	}

	p := r.Powerset(withDeadState)
	if p == nil {
		return nil
	}

	r = p.Reverse()
	if r == nil {
		return nil
	}

	return r.Powerset(withDeadState)
}

// NewState returns a new state added to the NFA. If the NFA was empty, the new
// state becomes the start state. NewState returns nil if the new state would
// be over n's limit.
func (n *NFA) NewState() *State {
	if n.budget != 0 {
		n.budget--
		s := &State{nfa: n}
		if n.Len() == 0 {
			n.start = s
		}
		s.Id()
		return s
	}

	return nil
}

// Powerset converts[3] the NFA into a NFA without ε edges, ie. into a DFA.
// Dead state is possibly constructed if withDeadState == true.
func (n *NFA) Powerset(withDeadState bool) (out *NFA) {
	alphabetSize := 0
	out = NewLimitedNFA(n.max)
	closures := map[string]*State{}
	var f func(closure) *State

	f = func(c closure) (result *State) {
		cid := c.id()
		if s, ok := closures[cid]; ok {
			return s
		}

		result = out.NewState()
		if result == nil {
			return nil
		}

		closures[cid] = result
		transitions := transitions{}
		for _, cset := range c.List() {
			result.IsAccepting = result.IsAccepting || cset.IsAccepting
			transitions2 := cset.transitions()
			for _, sym := range transitions2.List() {
				if sym < 0 {
					continue
				}

				nextStates := transitions2[sym]
				alphabetSize = mathutil.Max(alphabetSize, sym+1)
				for _, nextState := range nextStates.List() {
					for _, nextState := range nextState.closure().List() {
						transitions.newEdge(sym, true, nextState)
					}
				}
			}
		}
		for _, sym := range transitions.List() {
			result.NewEdge(sym, f(transitions[sym]))
		}

		return
	}

	out.start = f(n.Start().closure())
	var dead *State
	if withDeadState {
		for state := range out.s2i {
			edges := state.transitions()
			for sym := 0; sym < alphabetSize; sym++ {
				if _, ok := edges[sym]; !ok {
					if dead == nil {
						dead = out.NewState()
						if dead == nil {
							return nil
						}
					}
					state.NewEdge(sym, dead)
				}
			}
		}
		if dead != nil {
			for sym := 0; sym < alphabetSize; sym++ {
				dead.NewEdge(sym, dead)
			}
		}
	}
	return
}

// Reverse returns a NFA for the reverse language accepted by n.
func (n *NFA) Reverse() (out *NFA) {
	out = NewLimitedNFA(n.max)
	a := make([]*State, n.Len())
	for i := range a {
		a[i] = out.NewState()
		if a[i] == nil {
			return nil
		}
	}

	var acceptingIds []int
	for idFrom := 0; idFrom < n.Len(); idFrom++ {
		state := n.State(idFrom)
		if state.IsAccepting {
			acceptingIds = append(acceptingIds, idFrom)
		}
		for sym, tos := range n.State(idFrom).edges {
			for to := range tos {
				a[to.Id()].NewEdge(sym, a[idFrom])
			}
		}
	}

	a[n.start.Id()].IsAccepting = true
	switch len(acceptingIds) {
	case 1:
		out.start = a[acceptingIds[0]]
	default:
		out.start = out.NewState()
		if out.start == nil {
			return nil
		}

		for _, id := range acceptingIds {
			out.start.NewEdge(Epsilon, a[id])
		}
	}
	return
}

// SetStart sets the NFA's start state. Passing a state from a different NFA
// will panic.
func (n *NFA) SetStart(s *State) {
	if s.nfa != n {
		panic(s)
	}

	n.start = s
}

// Start returns the NFA's start state.
func (n *NFA) Start() *State { return n.start }

// State returns the NFA's state with Id() == id or nil if no such state exists.
func (n *NFA) State(id int) *State { return n.i2s[id] }

// String implements fmt.Stringer for debugging, etc.
func (n *NFA) String() string {
	return n.Str(nil)
}

// Str is like String but accepts a function transforming a non-ε edge value to
// string.
func (n *NFA) Str(fn func(int) string) string {
	var b bytes.Buffer
	for i := 0; i < n.Len(); i++ {
		b.WriteString(n.i2s[i].Str(fn))
	}
	return b.String()
}

// ---------------------------------------------------------------------- State

// State is one of the NFA states.
type State struct {
	nfa         *NFA
	IsAccepting bool // Whether this state is an accepting one.
	edges       transitions
}

// Closure returns a state set consisting of s and all states reachable from s
// through ε edges, transitively.
func (s *State) Closure() (c Closure) { return Closure{s.closure()} }

func (s *State) closure() (c closure) {
	c = closure{}
	var f func(*State)
	f = func(s *State) {
		if _, ok := c[s]; ok {
			return
		}

		c[s] = struct{}{}
		for s := range s.ε() {
			f(s)
		}
		return
	}
	f(s)
	return
}

func (s *State) edge(sym int) closure { return s.transitions().edge(sym, false) }

// Transitions returns the symbol -> closure projection of state s.
func (s *State) Transitions() Transitions {
	return Transitions{s.transitions()}
}

func (s *State) transitions() transitions {
	if s.edges == nil {
		s.edges = transitions{}
	}
	return s.edges
}

func (s *State) ε() closure { return s.edge(Epsilon) }

// Id returns the state's zero based index.
func (s *State) Id() int { return s.nfa.id(s) }

// NewEdge connects state s and state next by a new edge, labeled by sym. By
// convention, passing sym == Epsilon is reserved to indicate adding of an ε
// edge.
func (s *State) NewEdge(sym int, next *State) { s.transitions().newEdge(sym, true, next) }

var (
	isAcceptingL = map[bool]string{true: "["}
	isAcceptingR = map[bool]string{true: "]"}
	isStart      = map[bool]string{true: "->"}
	isSep        = map[bool]string{true: " "}
)

// String implements fmt.Stringer for debugging, etc.
func (s *State) String() string {
	return s.Str(nil)
}

// Str is like String but accepts a function transforming a non-ε edge value to
// string.
func (s *State) Str(fn func(int) string) string {
	var b bytes.Buffer
	f := strutil.IndentFormatter(&b, "\t")
	f.Format("%s%s[%d]%s\n%i",
		isStart[s == s.nfa.start],
		isAcceptingL[s.IsAccepting],
		s.Id(),
		isAcceptingR[s.IsAccepting],
	)
	transitions := s.transitions()
	for _, edge := range transitions.List() {
		nextSet := transitions[edge]
		switch {
		case edge == Epsilon:
			f.Format("ε -> ")
		default:
			switch {
			case fn != nil:
				f.Format("%s -> ", fn(edge))
			default:
				f.Format("%d (%#[1]U) -> ", edge)
			}
		}
		isFirst := true
		for _, next := range nextSet.List() {
			f.Format("%s[%d]", isSep[!isFirst], next.Id())
			isFirst = false
		}
		f.Format("\n")
	}
	return b.String()
}

// ----------------------------------------------------------------- Transitions

// Transitions maps symbols to their associated closures.
type Transitions struct{ transitions }

// NewTransitions returns a newly created Transitions.
func NewTransitions() Transitions { return Transitions{transitions{}} }

type transitions map[int]closure

func (t transitions) edge(sym int, canCreate bool) (c closure) {
	c = t[sym]
	if c == nil {
		c = closure{}
		if canCreate {
			t[sym] = c
		}
	}
	return c
}

func (t transitions) newEdge(sym int, canCreate bool, next *State) (c closure) {
	c = t.edge(sym, canCreate)
	c[next] = struct{}{}
	return c
}

// Delete removes the closure associated with sym.
func (t transitions) Delete(sym int) { delete(t, sym) }

// Get returns the closure associated with sym.
func (t transitions) Get(sym int) (c Closure) { return Closure{t[sym]} }

// Len returns the number of edges in transitions.
func (t transitions) Len() int { return len(t) }

// Set sets c as the closure associated with sym.
func (t transitions) Set(sym int, c Closure) { t[sym] = c.closure }

// List returns a sorted slice of all symbols appearing in t.
func (t transitions) List() (r []int) {
	r = make([]int, len(t))
	i := 0
	for sym := range t {
		r[i] = sym
		i++
	}
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return
}
