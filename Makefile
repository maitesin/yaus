tools-generate:
	go get -u github.com/matryer/moq

tools-lint: tools-generate
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

generate:
	go generate ./...

test: generate
	go test -cover -v ./...

lint: generate
	golangci-lint run

run:
	docker-compose up -d --build app

build:
	cd cmd/yaus && go build .
