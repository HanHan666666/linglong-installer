# Copyright 2023 The Rec Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all bench clean edit editor generate test

all:
	staticcheck

bench:
	date > log-bench
	go test -bench . 2>&1 | tee -a log-bench

clean:
	rm -f *.text *.out *.test
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile ../go.mod *.go & fi

editor:
	date > log-editor
	gofmt -l -s -w *.go
	go install -v ..
	go test -o /dev/null -c 2>&1 | tee -a log-editor
	date >> log-editor
	golint 2>&1 | tee -a log-editor

test:
	date > log-test
	go test -v 2>&1 | tee -a log-test
