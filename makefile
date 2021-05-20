PREFIX =	/usr
DESTDIR =

.PHONY: all clean build linux windows install deinstall

all: clean build

bin/goyoinker:
	go build -o bin/goyoinker cmd/terminal/goyoinker.go

build: bin/goyoinker

clean:
	rm -rf bin

linux:
	env GOOS=linux go build -o bin/linux/goyoinker cmd/terminal/goyoinker.go

windows:
	env GOOS=windows go build -o bin/windows/goyoinker.exe cmd/terminal/goyoinker.go

install: bin/goyoinker
	mkdir -p ${DESTDIR}${PREFIX}/bin
	install -m755 bin/goyoinker ${DESTDIR}${PREFIX}/bin/goyoinker

uninstall: 
	rm ${DESTDIR}${PREFIX}/bin/goyoinker
