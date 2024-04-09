build: 
	go build -o ./bin/os-file-management

run: build
	./bin/os-file-management

test: 
	go test -v ./...

install: build
	go install
