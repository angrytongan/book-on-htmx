.PHONY: help web lint lint-fix tailwind

include .env
export

help: # me
	@grep -e "^[a-z0-9_-]*:.* #" Makefile | sed -e 's/\(^[a-z0-9_-]*\):.* # \(.*\)/\1: \2/'

install: # Install things.
	npm install

web: # Run the web server.
	air

lint: # Run the linter.
	golangci-lint run ./cmd/web
	golangci-lint run ./internal/*

lint-fix: # Run the linter and fix things.
	golangci-lint run --fix ./cmd/web
	golangci-lint run --fix ./internal/*

tailwind:
	npx @tailwindcss/cli --input ./assets/css/input.css --output ./assets/css/style.css --watch
