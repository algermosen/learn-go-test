test:
	@go test ./...

vtest:
	@go test -v ./...

btest:
	@go test ./... -bench="."

ctest:
	@go test ./... -cover