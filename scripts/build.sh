#/bin/bash

env GOOS=linux GOARCH=amd64 GOARM=5 go build -o ./build/linux/goyoinker ./cmd/terminal/goyoinker.go
env GOOS=windows GOARCH=amd64 go build -o ./build/windows/goyoinker ./cmd/terminal/goyoinker.go 