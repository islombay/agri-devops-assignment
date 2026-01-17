run:
	go run cmd/app/main.go

build:
	go build -o app cmd/app/main.go

test:
	go test ./...
