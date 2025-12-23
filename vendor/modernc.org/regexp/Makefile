# Copyright 2023 The Regexp Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all benchcmp benchstat clean devbenchstat edit editor game game_input test

DIR := ${CURDIR}

all:

benchcmp:
	sh -c 'cd $$(go env GOROOT)/src/regexp && \
		go test -timeout 24h -count 1 -run @ -bench . -benchmem | tee $(DIR)/log-benchcmp-stdlib'
	go test -timeout 24h -count 1 -run @ -bench . -benchmem | tee log-benchcmp-modernc
	benchcmp -mag -changed log-benchcmp-stdlib log-benchcmp-modernc | tee benchcmp-results

benchstat:
	echo "This will run for about 90 minutes on a modernc machine"
	sh -c 'cd $$(go env GOROOT)/src/regexp && \
		go test -timeout 24h -count 20 -run @ -bench . -benchmem | tee $(DIR)/log-benchstat-stdlib'
	go test -timeout 24h -count 20 -run @ -bench . -benchmem | tee log-benchstat-modernc
	sed -i 's/modernc.org\/regexp/regexp/' log-benchstat-modernc
	benchstat log-benchstat-stdlib log-benchstat-modernc | tee benchstat-results

devbenchstat:
	go test -timeout 24h -count 20 -run @ -bench . -benchmem | tee log-benchstat-modernc
	sed -i 's/modernc.org\/regexp/regexp/' log-benchstat-modernc
	benchstat log-benchstat-stdlib log-benchstat-modernc | tee benchstat-results

clean:
	rm -f *.text *.out *.test
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json all_test.go dfa.go rec.go & fi

editor:
	date > log-editor
	gofmt -l -s -w *.go
	# go test -tags=probes 2>&1 | tee -a log-editor
	go test -c -o /dev/null | tee -a log-test
	date >> log-editor
	staticcheck

game:
	date > log-game
	go test -v -run @ -bench Game -benchmem -- -game 2>&1 | tee -a log-game
	date >> log-game

game_input: clean
	make -C testdata

test:
	date > log-test
	go test -v -failfast 2>&1 | tee -a log-test
	go build -v -o /dev/null | tee -a log-test
	date >> log-test
