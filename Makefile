NAME := discord-integration
VERSION := 2

build:
	go build -o bin/$(NAME) cmd/v$(VERSION)/main.go

run: build
	./bin/$(NAME)
