GOARCH=amd64
GOARM=5

ifeq ($(OS),Windows_NT)
    uname_S := Windows
	GOOS=windows
endif
ifeq ($(OS),Linux)
    uname_S := $(shell uname -s)
	GOOS=linux
endif

build: clean
	go build -o bin/goyoinker cmd/terminal/goyoinker.go
clean:
	rm -rf bin

install:
	cp bin/goyoinker /usr/bin

remove:
	rm /usr/bin/goyoinker