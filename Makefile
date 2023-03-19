.PHONY: all
all: build run

.PHONY: build
build:
	@go build -o ag

.PHONY: run
run:
	@./ag