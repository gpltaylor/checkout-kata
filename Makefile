build:
	@go build -o bin/$(APP_NAME) .

run: build
	@go run .

test:
	@go test -v ./...

bench:
	@$(eval COMMIT_HASH=$(shell git rev-parse HEAD))
	@go test -benchmem -bench BenchmarkCheckout > benchmark-$(COMMIT_HASH)