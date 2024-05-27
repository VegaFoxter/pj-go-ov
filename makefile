tools-local: tool-golangci-lint

tool-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum

docker-up:
	docker compose -f .docker/docker-compose.yml up -d --build

docker-down:
	docker compose -f .docker/docker-compose.yml down
	