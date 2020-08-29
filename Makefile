run:
	cd src/main && go run .

build:
	cd src/main && go build .
test:
	go test -race ./...
