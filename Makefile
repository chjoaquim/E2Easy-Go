compile:
	if [ -d "bin" ]; then rm -Rf bin; fi
	cd src; go mod download
	cd src; go build -o ../bin/E2Easy-go .
