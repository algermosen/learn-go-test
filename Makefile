test:
	@go test ./...

vtest:
	@go test -v ./...

bench:
	@go test ./... -bench="."