update: 
	go generate .

build: update
	GOOS=linux GOARCH=amd64 go build -o ./bin/apiserver ./cmd/apiserver/main.go

cluster: build
	MODE=cluster go run ./cmd/apiserver/main.go

dev: build
	MODE=dev go run ./cmd/apiserver/main.go

schema: build
	MODE=schema go run ./cmd/apiserver/main.go

