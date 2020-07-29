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

all: clean linux windows

linux:
	GOOS=linux
	go build -o bin/linux/goyoinker cmd/terminal/goyoinker.go

windows:
	GOOS=windows
	go build -o bin/windows/goyoinker.exe cmd/terminal/goyoinker.go

install:
	cp bin/goyoinker /usr/bin

remove: 
	rm /usr/bin/goyoinker