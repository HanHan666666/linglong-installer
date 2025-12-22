# Copyright 2024 The libXdmcp-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev download edit editor generate work test 

DIR = /tmp/libXdmcp
TAR = libXdmcp-1.1.2.tar.gz # linux/amd64
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
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile go.mod builder.json generator.go & fi

editor:
	gofmt -l -s -w .
	go test -c -o /dev/null
	go install -v  ./...
	go build -o /dev/null generator*.go

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi


generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	GO_GENERATE_DIR=$(DIR) go run generator*.go
	go build -v ./...
	go test -c -o /dev/null
	git status

dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	go build -v ./...
	go test -c -o /dev/null
	git status

test:
	go test -v -timeout 24h -count=1

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v4
	go work use ../libbsd
	go work use ../libc
