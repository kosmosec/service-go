all: run-db build-app startup-db start-app
run-db:
	docker-compose up -d

build-app:
	go build -o db-service

startup-db:
	sudo ./db-service migrate
	sudo ./db-service seed
start-app:
	./db-service

