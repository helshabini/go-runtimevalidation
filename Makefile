.PHONY: build test test-verbose cover benchmark install count-loc

build:
	go build ./...

test:
	go test ./...

test-verbose:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

benchmark:
	go test -bench=. -benchmem -count 1 ./...

install:
	go mod download

count-loc:
	find . -type f -name "*.go" -exec wc -l {} \; | awk '{split($$0,a," "); print a[1];}' | awk '{total = total + $$1}END{print total}'
