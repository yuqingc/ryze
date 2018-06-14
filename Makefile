# Makefile of the project

.PHONY: all

all: cmd/ryze/main.go vendor/%
	cd cmd/ryze/ && go build -o ryze.out
vendor/%:
	dep ensure

