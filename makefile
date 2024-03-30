image:
	docker build -t goapi:latest -f Dockerfile .
up:
	docker run --init -p:8080:8080 --env-file ./.env --name myapp goapi:latest
run:
	go run ./cmd/*