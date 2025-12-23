# Copyright 2014 The fsm Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor

all:

clean:
	rm -f cpu.test mem.test *.out
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json *.go & fi

editor:
	gofmt -l -s -w *.go
	go test 2>&1 | tee log-editor
	golint
