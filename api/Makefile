build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/apiserver ./cmd/apiserver/main.go

run: build
	go run ./cmd/apiserver/main.go

