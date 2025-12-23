# Copyright 2024 The ngrab Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor test

all: editor
	golint 2>&1
	staticcheck 2>&1

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w . 2>&1 | tee log-editor
	# go test -c -o /dev/null 2>&1 | tee -a log-editor
	go build -v  -o /dev/null ./... 2>&1 | tee -a log-editor
	go test  2>&1 | tee -a log-editor

test:
	go test -v -timeout 24h -count=1 2>&1 | tee log-test
	grep -a 'TRC\|TODO\|ERRORF\|FAIL' log-test || true 2>&1 | tee -a log-test
