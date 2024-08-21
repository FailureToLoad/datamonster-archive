update: 
	go generate .

build: update
	GOOS=linux GOARCH=amd64 go build -o ./bin/apiserver ./cmd/apiserver/main.go

dev: build
	APP_ENV=dev go run ./cmd/apiserver/main.go

schema: build
	APP_ENV=schema go run ./cmd/apiserver/main.go

preview: build
	APP_ENV=prod go run ./cmd/apiserver/main.go

