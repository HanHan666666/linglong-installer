// Copyright 2023 The Rec Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rec // import "modernc.org/rec/lib"

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"modernc.org/regexp"
	"modernc.org/regexp/syntax"
)

func (c *cfg) match(fn string, args []string, utf8Input, strInput bool) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 non-flag argument: a Go-flavored regular expression")
	}

	expr := args[0]
	dfa, err := regexp.NewMatchDFA(expr)
	if err != nil {
		return err
	}

	if *c.oDbg {
		if _, err := fmt.Fprintf(c.stderr, "%s\nprefix %q\n", dfa, dfa.Prefix()); err != nil {
			return err
		}
	}

	c.prolog(false)
	startCond := dfa.Cond()
	c.w("// %s matches %s, start condition %v\n", fn, pretty(expr), startCond)
	rx := c.rxString()
	switch {
	case strInput:
		c.w("func %s%s(s string) bool {\n", rx, fn)
	default:
		c.w("func %s%s(s []byte) bool {\n", rx, fn)
	}

	defer c.w("}\n")

	if expr == "" {
		c.w("\treturn true\n")
		return nil
	}

	if startCond == ^syntax.EmptyOp(0) {
		c.w("\treturn false\n")
		return nil
	}

	c.w("\tconst endOfText = %#0x\n", unicode.MaxRune+1)
	c.w("\tconst minInputLen = %d\n", dfa.MinInputLen())
	if dfa.Prefix() != "" {
		switch pr := dfa.PrefixRune(); {
		case utf8Input:
			pb := rune([]byte(string(pr))[0])
			c.w("\tconst prefixByte = %s\n", strconv.QuoteRuneToASCII(pb))
		default:
			pb := rune(byte(pr))
			c.w("\tconst prefixByte = %s\n", strconv.QuoteRuneToASCII(pb))
		}
		switch {
		case strInput:
			c.w("\tprefix := %q\n", dfa.Prefix())
		default:
			c.w("\tprefix := []byte(%q)\n", dfa.Prefix())
		}
	}
	c.w("\tvar pos, pos0, width, width0, width1 int\n")
	if dfa.Prefix() != "" {
		c.w("\tvar advance int\n")
	}
	c.w("\tvar r, r1 rune\n")
	c.w("\t_ = r\n")
	c.w("\t_ = r1\n")
	c.w("\t_ = width1\n")
	c.w("\tlastRestartPos := -1\n")
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
	c.w("restart:\n")
	if *c.oTrc {
		c.trc("==== start/restart: pos %d\n", "pos")
	}
	c.w("\tr, r1 = endOfText, endOfText\n")
	c.w("\twidth, width1 = 0, 0\n")

	stepf := func(prefix, r, width, pos string) string {
		return fmt.Sprintf("%s\t%s, %s = step(%s);", prefix, r, width, pos)
	}

	step := func(prefix, r, width, pos string) { c.w("%s", stepf(prefix, r, width, pos)) }

	step("", "r", "width", "pos")
	c.w("\tif r != endOfText {\n")
	step("\t", "r1", "width1", "pos+width")
	c.w(" }\n")
	c.w("\tif pos == lastRestartPos { return false }; lastRestartPos = pos\n")
	pc := dfa.StartPC()
	switch {
	case startCond&syntax.EmptyBeginText != 0:
		c.w("\tif pos != 0 { return false }\n")
		if dfa.Prefix() != "" {
			c.w("\tif r != %s { return false }\n", strconv.QuoteRuneToASCII(dfa.PrefixRune()))
		}
	case dfa.Prefix() != "":
		pc = dfa.PrefixPC()
		switch {
		case strInput:
			c.w("\tfor {\n")
			c.w("\t\tif len(s)-pos < minInputLen { return false }\n")
			c.w("\t\tif advance = strings.IndexByte(s[pos:], prefixByte); advance < 0 { return false }\n")
			c.w("\t\tpos += advance\n")
			c.w("\t\tif len(s)-pos < minInputLen { return false }\n")
			c.w("\t\tif s[pos:pos+len(prefix)] == prefix { break }\n")
			c.w("\t\tpos++\n")
			c.w("\t}\n")
		default:
			c.w("\tfor {\n")
			c.w("\t\tif len(s)-pos < minInputLen { return false }\n")
			c.w("\t\tif advance = bytes.IndexByte(s[pos:], prefixByte); advance < 0 { return false }\n")
			c.w("\t\tpos += advance\n")
			c.w("\t\tif len(s)-pos < minInputLen { return false }\n")
			c.w("\t\tif bytes.Equal(s[pos:pos+len(prefix)], prefix) { break }\n")
			c.w("\t\tpos++\n")
			c.w("\t}\n")
		}
	default:
		c.w("\t\tif len(s)-pos < minInputLen { return false }\n")
	}
	c.w("\tpos0 = pos; width0 = width\n")
	if dfa.Prefix() != "" {
		c.w("\tpos += len(prefix)\n")
		step("", "r", "width", "pos")
		step("", "r1", "width1", "pos+width")
		c.w("\twidth0 = width\n")
		c.w("\n")
	}
	prog := dfa.Prog()

	gotof := func(pc uint32) string {
		switch prog[pc] >> regexp.DFARuneBits {
		case regexp.DFAOpAccept:
			return "return true"
		default:
			return fmt.Sprintf("goto l%d", pc)
		}
	}

	consumef := func(pc uint32) string {
		s := gotof(pc)
		if strings.Contains(s, "return") {
			return s
		}

		return fmt.Sprintf("pos += width; if r, width = r1, width1; r != endOfText { %s }; %s;", stepf("", "r1", "width1", "pos+width"), s)
	}

	c.w("goto l%d\n", pc)
	pc = dfa.StartPC()
	for pc < len(prog) {
		c.w("goto l%d; l%[1]d:\n", pc)
		if *c.oTrc {
			c.trc("---- new state\n", "")
		}
		sym := prog[pc]
		pc0 := pc
		pc++
		if *c.oTrc {
			c.trc("%04d: %#U, pos %d, r %#U, %d, r1 %#U, %d, pos0 %d, width0 %d\n", fmt.Sprintf("%d, %d, pos, r, width, r1, width, pos0, width0", pc0, sym))
		}
		switch op, ch := sym>>regexp.DFARuneBits, sym&regexp.DFARuneMask; op {
		case regexp.DFAOpCharClass:
			next := prog[pc]
			pc++
			cls := dfa.CharClass(int(ch))
			out := false
			for i := 0; i < len(cls); i += 2 {
				switch lo, hi := cls[i], cls[i+1]; {
				case lo == hi:
					c.w("\tif r == %s { %s }\n", strconv.QuoteRuneToASCII(lo), consumef(next))
				default:
					out = true
					c.w("\tif r < %s { goto l%dout }\n", strconv.QuoteRuneToASCII(lo), pc0)
					c.w("\tif r <= %s { %s }\n", strconv.QuoteRuneToASCII(hi), consumef(next))
				}
			}
			if out {
				c.w("l%dout:\n", pc0)
			}
		case regexp.DFAOpStop:
			c.w("\tpos = pos0 + width0; goto restart\n")
		case regexp.DFAOpRune:
			next := prog[pc]
			pc++
			c.w("\tif r == %s { %s }\n", strconv.QuoteRuneToASCII(rune(ch)), consumef(next))
		case regexp.DFAOpAccept:
			c.w("\treturn true\n")
		case regexp.DFAOpBeginText:
			next := prog[pc]
			pc++
			c.w("\tif pos != pos0 { return false }\n")
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
