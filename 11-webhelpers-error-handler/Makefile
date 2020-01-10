all: run-db build-sales-admin build-sales-api startup-db start-app
run-db:
	docker-compose up -d

build-sales-admin:
	go build -o sales-admin ./cmd/sales-admin/main.go

build-sales-api:
	go build -o sales-api ./cmd/sales-api/main.go

startup-db:
	./sales-admin migrate
	./sales-admin seed
start-app:
	./sales-api --web-address=localhost:8080

