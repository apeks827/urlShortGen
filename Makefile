run:
	docker-compose -f docker-compose.yml up --build

postgres:
	docker-compose run -e MEMORY_MODE=postgres -p 8080:8080 urlshortgen

test:
	go test --short -v ./...