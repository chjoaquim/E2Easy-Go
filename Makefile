GOOS=linux
export SHELL:=/bin/bash
export SHELLOPTS:=$(if $(SHELLOPTS),$(SHELLOPTS):)pipefail:errexit

compile-linux:
	if [ -d "bin" ]; then rm -Rf bin; fi
	cd src; go mod download
	cd src; GOOS=$(GOOS) go build -o ../bin/E2Easy-go .
	cd bin; chmod +x E2Easy-go

compile:
	if [ -d "bin" ]; then rm -Rf bin; fi
	cd src; go mod download
	cd src; go build -o ../bin/E2Easy-go .
	cd bin; chmod +x E2Easy-go