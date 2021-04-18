tools-generate:
	go get -u github.com/matryer/moq

tools-lint: tools-generate
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

generate:
	go generate ./...

test: generate
	go test -v ./...

lint:
	golangci-lint run