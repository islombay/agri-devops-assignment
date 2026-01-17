run:
	go run cmd/app/main.go

build:
	go build -o app cmd/app/main.go

test:
	go test ./...

dbuild:
	docker build -t islombay/agriculture-app:latest .

drun:
	docker run --rm -p 8080:8080 islombay/agriculture-app:latest

dpush:
	docker push islombay/agriculture-app:latest