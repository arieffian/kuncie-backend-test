.PHONY: all test clean build docker

run:
	go run cmd/main.go

build:
	test
	go build -ldflags "-s -w" -o build/kuncie-backend.app cmd/main.go
	
test:
	go test ./... -cover -vet -all -v

docker-build:
	test
	docker build -t kuncie-backend -f Dockerfile .

docker-run:
	docker run -d -p 8080:8080 kuncie-backend:latest