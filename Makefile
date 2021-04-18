tools-lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

generate:
	go generate ./...

test: generate
	go test -v ./...

lint: generate
	golangci-lint run