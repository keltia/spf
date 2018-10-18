# Main Makefile for spf
#
# Copyright 2018 © by Ollivier Robert
#

GO=		go
GOBIN=  ${GOPATH}/bin

SRCS= domain.go resolve.go result.go utils.go
SRCT= domain_test.go resolve_test.go result_test.go utils_test.go

OPTS=	-ldflags="-s -w" -v

all: build

build: ${SRCS}
	${GO} build ${OPTS} ./cmd/...

test: ${SRCS} ${SRCT}
	${GO} test -v .

install: ${BIN}
	${GO} install ${OPTS} .

clean:
	${GO} clean -v

push:
	git push --all
	git push --tags
