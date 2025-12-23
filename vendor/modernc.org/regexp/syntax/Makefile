# Copyright 2023 The Regexp Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor

DIR := ${CURDIR}

all:

clean:
	rm -f *.text *.out *.test
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	date > log-editor
	gofmt -l -s -w *.go
	go test -tags=probes 2>&1 | tee -a log-editor
	date >> log-editor
	golint
