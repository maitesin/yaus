tools-generate:
	go get -u github.com/matryer/moq

tools-lint: tools-generate
	go get -u github.com/quasilyte/go-ruleguard/dsl # Temporary dependency to avoid go-critic failing
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
