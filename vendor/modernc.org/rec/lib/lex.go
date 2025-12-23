// Copyright 2023 The Rec Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rec // import "modernc.org/rec/lib"

import (
	"fmt"
	"strconv"
	"unicode"

	"modernc.org/regexp"
)

func (c *cfg) lex(fn string, args []string, utf8Input, strInput bool) error {
	if len(args) == 0 {
		return fmt.Errorf("expected 1 or more non-flag argument: a Go-flavored regular expression")
	}

	dfa, err := regexp.NewLexDFA(args)
	if err != nil {
		return err
	}

	if *c.oDbg {
		if _, err := fmt.Fprintf(c.stderr, "%s", dfa); err != nil {
			return err
		}
	}

	m := map[int]string{}
	for id := 0; id < dfa.CharClasses(); id++ {
		cls := dfa.CharClass(id)
		if s := tables.find(cls); s != "" {
			m[id] = s
		}
	}
	c.prolog(true)
	kind := "ASCII"
	if utf8Input {
		kind = "UTF-8"
	}
	c.w(`// %s recognizes longest %s lexemes. Lower IDs take precedence on same length.
//
`, fn, kind)
	for i, v := range args {
		c.w("//\tid %3d: %s\n", i, pretty(v))
	}
	c.w("//\n// ID == -1 is returned when no lexeme was recognized.\n")
	rx := c.rxString()
	switch {
	case strInput:
		c.w("func %s%s(s string) (id, length int) {\n", rx, fn)
	default:
		c.w("func %s%s(s []byte) (id, length int) {\n", rx, fn)
	}

	defer c.w("}\n")

	c.w("\tconst endOfText = %#0x\n", unicode.MaxRune+1)
	c.w("\tvar pos, pos0, width, width1 int\n")
	c.w("\tid = -1\n")
	c.w("\tvar r, r1 rune\n")
	c.w("\t_ = pos0\n")
	c.w("\t_ = r\n")
	c.w("\t_ = r1\n")
	c.w("\t_ = width1\n")
	switch {
	case strInput:
		switch {
		case utf8Input:
			c.w("\tstep := func(pos int) (r rune, n int) { %sif pos < len(s) { c := s[pos]; if c < utf8.RuneSelf { return rune(c), 1 }; return utf8.DecodeRuneInString(s[pos:]) }; return endOfText, 0 }\n", c.lcase)
		default:
			c.w("\tstep := func(pos int) (r rune, n int) { %sif pos < len(s) { return rune(s[pos]), 1}; return endOfText, 0 }\n", c.lcase)
		}
	default:
		switch {
		case utf8Input:
			c.w("\tstep := func(pos int) (r rune, n int) { %sif pos < len(s) { c := s[pos]; if c < utf8.RuneSelf { return rune(c), 1 }; return utf8.DecodeRune(s[pos:]) }; return endOfText, 0 }\n", c.lcase)
		default:
			c.w("\tstep := func(pos int) (r rune, n int) { %sif pos < len(s) { return rune(s[pos]), 1}; return endOfText, 0 }\n", c.lcase)
		}
	}

	stepf := func(prefix, r, width, pos string) string {
		return fmt.Sprintf("%s\t%s, %s = step(%s);", prefix, r, width, pos)
	}

	c.w("\tmove := func() { pos += width; if r, width = r1, width1; r != endOfText { %s }; }\n", stepf("", "r1", "width1", "pos+width"))
	c.w("\taccept := func(x rune) bool { if r == x { move(); return true }; return false }\n_ = accept\n")
	c.w("\taccept2 := func(x rune) bool { if r <= x { move(); return true }; return false }\n_ = accept2\n")
	if *c.oTrc {
		c.trc("==== start: pos %d\n", "pos")
	}
	c.w("\tr, r1 = endOfText, endOfText\n")
	c.w("\twidth, width1 = 0, 0\n")

	step := func(prefix, r, width, pos string) { c.w("%s", stepf(prefix, r, width, pos)) }

	step("", "r", "width", "pos")
	c.w("\tif r != endOfText {\n")
	step("\t", "r1", "width1", "pos+width")
	c.w(" }\n")
	prog := dfa.Prog()

	gotof := func(pc uint32) string { return fmt.Sprintf("goto l%d", pc) }
	consumef := func(pc uint32) string { return fmt.Sprintf("move(); %s;", gotof(pc)) }

	for pc := dfa.StartPC(); pc < len(prog); {
		if pc != dfa.StartPC() {
			c.w("goto l%d\nl%[1]d:\n", pc)
		}
		if *c.oTrc {
			c.trc("---- new state\n", "")
		}
		sym := prog[pc]
		pc0 := pc
		pc++
		if *c.oTrc {
			c.trc("%04d: %#U, pos %d, r %#U, %d, r1 %#U, %d, pos0 %d\n", fmt.Sprintf("%d, %d, pos, r, width, r1, width, pos0", pc0, sym))
		}
		switch op, ch := sym>>regexp.DFARuneBits, sym&regexp.DFARuneMask; op {
		case regexp.DFAOpCharClass:
			next := prog[pc]
			pc++
			cls := dfa.CharClass(int(ch))
			switch nm := m[int(ch)]; {
			case nm != "":
				c.w("\tif unicode.Is(unicode.%s, r) { %s }\n", nm, consumef(next))
			default:
				for i := 0; i < len(cls); i += 2 {
					switch lo, hi := cls[i], cls[i+1]; {
					case lo == hi:
						c.w("\tif accept(%s) { %s }\n", strconv.QuoteRuneToASCII(lo), gotof(next))
					default:
						if lo > 0 {
							c.w("\tif r < %s { goto l%dout }\n", strconv.QuoteRuneToASCII(lo), pc0)
						}
						c.w("\tif accept2(%s) { %s }\n", strconv.QuoteRuneToASCII(hi), gotof(next))
					}
				}
				c.w("l%dout:\n", pc0)
			}
		case regexp.DFAOpStop:
			c.w("\treturn id, length\n")
		case regexp.DFAOpRune:
			next := prog[pc]
			pc++
			c.w("\tif accept(%s) { %s }\n", strconv.QuoteRuneToASCII(rune(ch)), gotof(next))
		case regexp.DFAOpAccept:
			if *c.oTrc {
				c.trc("accept %d\n", fmt.Sprint(ch))
			}
			c.w("\tid, length = %[1]d, pos\n", ch)
			if *c.oTrc {
				c.trc("id %d, length %d\n", "id, length")
			}
		case regexp.DFAOpBeginText:
			next := prog[pc]
			pc++
			c.w("\tif pos != pos0 { return -1, 0 }\n")
			c.w("\t%s\n", gotof(next))
		case regexp.DFAOpEndText:
			next := prog[pc]
			pc++
			c.w("\tif r == endOfText { %s }\n", gotof(next))
		default:
			c.w("\tpanic(%q)\n", fmt.Sprintf("%04d: %#U", pc0, sym))
		}
	}
	return nil
}
