.PHONY: setup lint

setup:
		go install github.com/golang/mock/mockgen

lint:
		golangci-lint run --issues-exit-code 0
