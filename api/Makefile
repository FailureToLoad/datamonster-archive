build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/apiserver ./cmd/apiserver/main.go

run: build
	MODE=cluster go run ./cmd/apiserver/main.go

