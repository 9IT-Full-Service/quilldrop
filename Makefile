.PHONY: serve build docker-static docker-run test clean docker-up docker-down generate-keys migrate

serve:
	go run . serve

generate:
	go run . generate

build:
	GOOS=linux GOARCH=arm64 go build -o quilldrop

docker-static:
	podman build --platform linux/arm64 -t quilldrop -f ./Dockerfile-static .

docker-run:
	podman run --rm -i -p 8081:80 quilldrop:latest

test:
	go test ./...

clean:
	rm -rf bin/

docker-up:
	docker compose up -d

docker-down:
	docker compose down

generate-keys:
	@mkdir -p keys
	openssl genrsa -out keys/private.pem 2048
	openssl rsa -in keys/private.pem -pubout -out keys/public.pem
	@echo "RSA key pair generated in keys/"

migrate:
	go run ./cmd/migrate