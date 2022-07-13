test:
	go test -v ./...

lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.46.2 golangci-lint run -v

run:
	docker-compose up --build

.PHONY: test lint