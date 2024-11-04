build:
	@go build -o ./bin/store

run: build
	@./bin/store

test:
	@go test ./... -v