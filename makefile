build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go clean -testcache && go test -v ./internal

bench:
	go test -bench=... ./...