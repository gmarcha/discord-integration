NAME := discord-integration

build:
	go build -o bin/$(NAME) cmd/v1/main.go

run: build
	./bin/$(NAME)
