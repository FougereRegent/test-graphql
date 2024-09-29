BUILD_NAME=bin
FLAGS=

run: generate
	go run .

build: generate
	go build -o $(BUILD_NAME)

generate:
	go generate


test:
	go test
