test:
	go test -v -count 1 -race -timeout 1m ./...

run-dev:
	go run -race ./cmd/pikchr/pikchr.go
