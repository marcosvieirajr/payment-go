watch:
	@echo starting reflex to watch '.go' files changes and re-run the service
	@PATH=$(PATH):$(HOME)/go/bin reflex -r '.go' -s -- go run ./cmd/restapp/main.go

setup:
	@echo installing reflex developer tool
	go install github.com/cespare/reflex@latest
	go mod tidy && go mod vendor

run:
	@echo building application docker image and starting stack \(PostgreSQL\) from docker compose
	docker-compose up --build

test:
	@echo run all tests
	go test ./...
