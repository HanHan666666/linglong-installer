# Copyright 2023 The libXau-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev download edit editor generate work test

DIR = /tmp/libXau
TAR = libXau-1.0.9.tar.gz # linux/amd64
URL = https://www.x.org/releases/individual/lib/$(TAR)

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile go.mod builder.json generator.go & fi

editor:
	gofmt -l -s -w . 2>&1 | tee log-editor
	go test -c -o /dev/null 2>&1 | tee -a log-editor
	go install -v  ./... 2>&1 | tee -a log-editor
	go build -o /dev/null generator*.go

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi


generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_DIR=$(DIR) go run generator*.go 2> log-generate-errors | tee log-generate
	cat log-generate-errors
	go build -v ./...
	go build -v ./...  | tee -a log-generate
	go test -c -o /dev/null | tee -a log-generate
	git status
	grep 'PASS\|TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'PASS\|TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true

dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go 2>&1 | tee log-generate
	go build -v ./...  | tee -a log-generate
	go test -c -o /dev/null | tee -a log-generate
	git status
	grep 'PASS\|TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'PASS\|TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true
	grep 'PASS\|TRC\|TODO\|ERRORF\|FAIL' /tmp/ccgo.log || true

test:
	go test -v -timeout 24h -count=1 2>&1 | tee log-test
	grep -a 'TRC\|TODO\|ERRORF\|FAIL' log-test || true 2>&1 | tee -a log-test

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../ccgo/v4
	go work use ../ccgo/v3
	go work use ../cc/v4
	go work use ../libc
