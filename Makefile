.PHONY: cli-build build install clean

MAKE_SCRIPT="./build/make.sh"

default: all

all: getdeps fmt build

fmt:
	${MAKE_SCRIPT} fmt

build:
	${MAKE_SCRIPT} compile

getdeps:
	${MAKE_SCRIPT} getdeps

clean:
	${MAKE_SCRIPT} clean

