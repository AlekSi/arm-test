export GOPATH := $(PWD)

all:
	env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -v
	./armtest
