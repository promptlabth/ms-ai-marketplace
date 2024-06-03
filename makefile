image:
	docker build -t goapi:latest -f Dockerfile .
up:
	docker run --init -p:8081:8081 --env-file ./.env --name myapp goapi:latest
run:
	go run ./cmd/*