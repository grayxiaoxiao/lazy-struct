
build:
	go build -o lazy-struct main.go

install:
	mv lazy-struct ${GOPATH}/bin