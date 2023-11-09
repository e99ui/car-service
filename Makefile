.PHONY: dc
dc:
	docker compose up --remove-orphans --build

.PHONY: run
run:
	go build -o app cmd/app/main.go && HTTP_ADDR=:8080 ./app

.PHONY: test
test: 
	go test -race ./...

.PHONY: lint
lint:
	golangci-lint run
