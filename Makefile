compile:
	if [ -d "bin" ]; then rm -Rf bin; fi
	go mod download
	go build -o bin/E2Easy-go .
