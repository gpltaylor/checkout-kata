build:
	@go build -o bin/$(APP_NAME) .

run: build
	@go run .

test:
	@go test -v ./...


