#build all binaries

build:
	go build -o bin/go-server internal/main.go
clean: 
	rm -rf bin/*

