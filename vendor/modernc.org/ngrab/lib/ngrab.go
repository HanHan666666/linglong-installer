// Copyright 2024 The ngrab Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package implements the ngrab command.
package lib // import "modernc.org/ngrab/lib"

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Node represents a documentation node
//
// Every input .n file is represented by a separate node tree. The root node
// has Type == "document" and Text == file path.
type Node struct {
	ID     int
	Parent int `json:",omitempty"`
	Type   string
	Text   string `json:",omitempty"`
	p      *Node
}

func (n *Node) indent() string {
	i := 0
	for n.p != nil {
		i++
		n = n.p
	}
	return strings.Repeat("· ", i)
}

// Task describes a job.
type Task struct {
	// The nodes used for producing the JSON
	Nodes []*Node

	files []string
	fn    string
	lines []string
	w     io.Writer

	id     int
	lineNo int
}

// NewTask returns a newly created Task or an error, if any.
func NewTask(w io.Writer, files []string) (r *Task, err error) {
	return &Task{
		w:     w,
		files: files,
	}, nil
}

// Main executes 't'.
func (t *Task) Main() (err error) {
	for _, v := range t.files {
		if err := t.file(v); err != nil {
			return err
		}
	}
	b, err := json.MarshalIndent(t.Nodes, "", "\t")
	if err != nil {
		return err
	}

	_, err = t.w.Write(b)
	return err
}

func (t *Task) newNode(p *Node, typ, text string) (r *Node) {
	var pid int
	if p != nil {
		pid = p.ID
	}
	t.id++
	r = &Node{ID: t.id, Parent: pid, Type: typ, Text: text, p: p}
	t.Nodes = append(t.Nodes, r)
	return r
}

type parseError string

func (t *Task) file(fn string) (err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// ok
		case parseError:
			if err == nil {
				err = fmt.Errorf("%v:%v: %s", t.fn, t.lineNo, x)
			}
		default:
			panic(x)
		}
	}()

	b, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	t.lines = strings.Split(string(b), "\n")
	t.fn = fn
	t.parseDoc(fn)
	return nil
}

func (t *Task) peek() (r string) {
	if len(t.lines) != 0 {
		r = t.lines[0]
		if r == "" {
			r = "."
		}
	}
	return r
}

func (t *Task) consume() (r string) {
	if len(t.lines) != 0 {
		r = t.lines[0]
		if r == "" {
			r = "."
		}
		t.lines = t.lines[1:]
		t.lineNo++
	}
	return r
}

func (t *Task) parseDoc(fn string) {
	p := t.newNode(nil, "document", fn)
	t.lineNo = 1
	t.parse(p, 0)
}

func (t *Task) parse(p *Node, nest int) {
	for line := t.peek(); line != ""; line = t.peek() {
		switch {
		case
			strings.HasPrefix(line, "'"),
			strings.HasPrefix(line, `.\"`):

			t.newNode(p, "comment", t.consume()[1:])
		case
			!strings.HasPrefix(line, "."),
			line == ".",
			strings.HasPrefix(line, ". "):

			t.newNode(p, "text", t.consume())
		case
			strings.HasPrefix(line, ".BS"),
			strings.HasPrefix(line, ".CS"),
			strings.HasPrefix(line, ".DS"),
			strings.HasPrefix(line, ".RS"),
			strings.HasPrefix(line, ".SO"):

			tag := t.peek()[1:3]
			t.parse(t.newNode(p, tag, t.consume()[3:]), nest+1)
		case
			strings.HasPrefix(line, ".BE"),
			strings.HasPrefix(line, ".CE"),
			strings.HasPrefix(line, ".DE"),
			strings.HasPrefix(line, ".RE"),
			strings.HasPrefix(line, ".SE"):

			tag := t.peek()[1:3]
			t.newNode(p, tag, t.consume()[3:])
			if nest != 0 {
				return
			}
		default:
			tag := line[1:3] // .TH etc
			t.newNode(p, tag, t.consume()[3:])
		}
	}
}
