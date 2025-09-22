.PHONY: main help test

main: help

help:
	@echo "Pagescript Makefile commands"
	@echo "usage: \`make <command>\`"
	@echo "  - help: display this help text"
	@echo "  - test: run all module tests"

test:
	go test ./...
