run:
	cd app/src/main && go run .

build:
	cd app/src/main && go build .
test:
	go test -race ./...

lint:
	golangci-lint run
