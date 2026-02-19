.PHONY: serve build docker-static docker-run test clean docker-up docker-down generate-keys migrate

LATEST_TAG := $(shell git tag --sort=-version:refname | head -n 1)

serve:
	go run . serve

generate:
	go run . generate

build:
	GOOS=linux GOARCH=arm64 go build -o quilldrop

docker-static:
	podman build --platform linux/arm64 -t quilldrop -f ./Dockerfile-static .

docker-images:
	@echo "Building for tag: $(LATEST_TAG)"
	@echo "Building AMD64..."
	podman build --platform linux/amd64 \
		-t ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images-amd64 \
		-f ./Dockerfile-images .
	@echo "Building ARM64..."
	podman build --platform linux/arm64 \
		-t ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images-arm64 \
		-f ./Dockerfile-images .
	@echo "Creating manifest..."
	podman manifest rm ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images 2>/dev/null || true
	podman manifest create ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images
	@echo "Adding AMD64 to manifest..."
	podman manifest add ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images \
		ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images-amd64
	@echo "Adding ARM64 to manifest..."
	podman manifest add ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images \
		ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images-arm64
	@echo "Inspecting manifest:"
	podman manifest inspect ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images
	@echo "Pushing manifest:"
	podman manifest push --all ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images
# docker-images:
# 	$(eval LATEST_TAG := $(shell git tag --sort=-version:refname | head -n 1))
# 	podman build --platform linux/amd64,linux/arm64 \
# 		--manifest ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images \
# 		-f ./Dockerfile-images .
# 	podman manifest push --all ghcr.io/9it-full-service/quilldrop:$(LATEST_TAG)-images

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