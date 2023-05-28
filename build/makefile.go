package build

import "strings"

func MakeFile_tem(svr string) string {
	t := `
server:
	mkdir -p bin
	go build -o 'bin/server' -mod=vendor ./cmd/server
run:
	go run cmd/server.go

build:
	GOOS=linux GOARCH=amd64 go build -o 'bin/server' ./cmd/server.go

test:
	sh ./test.sh

prod:
	sh ./build.sh sample 10003 1.0


.PHONY: server
.PHONY: proto
.PHONY: test

`

	t = strings.ReplaceAll(t, "sample", svr)
	return t

}
