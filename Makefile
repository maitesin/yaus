generate:
	go generate ./...

test: generate
	go test -v ./...

lint: generate
	golangci-lint run