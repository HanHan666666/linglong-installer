# Copyright 2024 The libtk9_0-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev dwonload edit editor generate test work

TAR = tk9.0.1-src.tar.gz
URL = http://prdownloads.sourceforge.net/tcl/$(TAR)

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

clean-dev:
	rm -rf include/linux/amd64/
	rm -f ccgo_linux_amd64.go internal/tcltest/ccgo_linux_amd64.go

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi

edit:
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile all_test.go go.mod builder.json generator.go libtk.go & fi

editor:
	gofmt -l -s -w .
	GOOS=linux GOARCH=386 go test -c -o /dev/null
	GOOS=linux GOARCH=386 go build -v  -o /dev/null ./...
	GOOS=linux GOARCH=386 go build -o /dev/null generator*.go

generate: download
	go run generator*.go
	go build -v ./...
	git status

dev: download
	GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	go build -v ./...
	git status

windows: download
	GO_GENERATE_WIN=1 go run generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	#TODO GOOS=windows GOARCH=arm64 go build -v ./...
	#TODO GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status

windows-dev: download
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN=1 GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	#TODO GOOS=windows GOARCH=arm64 go build -v ./...
	#TODO GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status

test:
	go test -vet=off -v -timeout 24h -count=1

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../libc
	go work use ../libtcl9.0
	go work use ../libz
