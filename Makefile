.PHONY: serve serve-dev generate build build-native build-macos build-linux \
	build-raspberry build-windows checksums new-theme clean \
	docker-images docker-static docker-run

GIT := $(shell git pull --quiet 2>/dev/null || true)
LATEST_TAG := $(shell git tag --sort=-version:refname | head -n 1)

tag:
	@echo "Latest tag: $(LATEST_TAG)"

serve:
	go run . serve

# Server mit Template-Live-Reload (Theme-Entwicklung), optional: make serve-dev THEME=newdesign
serve-dev:
	go run . serve -dev $(if $(THEME),-theme $(THEME))

generate:
	go run . generate

# ---------------------------------------------------------------------------
# Builds
#
# Alle Binaries landen in $(DIST)/ und heissen quilldrop-<os>-<arch>[v<arm>].
# Einzelne Ziele: make build-linux, build-macos, build-windows, build-raspberry
# ---------------------------------------------------------------------------

BINARY_NAME := quilldrop
DIST        := dist
GOFLAGS_REL := -trimpath -ldflags=-s\ -w

# $(call go_build,GOOS,GOARCH[,GOARM][,Kommentar])
define go_build
	@mkdir -p $(DIST)
	@echo "  -> $(1)/$(2)$(if $(3),v$(3))$(if $(4),   [$(4)])"
	@GOOS=$(1) GOARCH=$(2) GOARM=$(3) CGO_ENABLED=0 \
		go build $(GOFLAGS_REL) \
		-o $(DIST)/$(BINARY_NAME)-$(1)-$(2)$(if $(3),v$(3))$(if $(filter windows,$(1)),.exe) .
endef

# Alle Plattformen
build: build-macos build-linux build-windows checksums
	@echo ""
	@ls -1 $(DIST)

# Nur fuer die aktuelle Maschine -> ./quilldrop
build-native:
	CGO_ENABLED=0 go build $(GOFLAGS_REL) -o $(BINARY_NAME) .

build-macos:
	@echo "macOS:"
	$(call go_build,darwin,amd64,,Intel)
	$(call go_build,darwin,arm64,,Apple Silicon M1-M4)

build-linux:
	@echo "Linux:"
	$(call go_build,linux,amd64,,x86_64)
	$(call go_build,linux,386,,x86 32-bit)
	$(call go_build,linux,arm64,,ARM 64-bit)
	$(call go_build,linux,arm,7,ARMv7 32-bit)
	$(call go_build,linux,arm,6,ARMv6 32-bit)

# Raspberry Pi (Teilmenge von build-linux, zum separaten Bauen)
#   armv6  : Pi 1, Pi Zero / Zero W, CM1
#   armv7  : Pi 2 und Pi 3/4/5 mit 32-bit OS
#   arm64  : Pi 3/4/5, Pi Zero 2 W, CM3/CM4/CM5 mit 64-bit OS (Raspberry Pi OS 64-bit)
build-raspberry:
	@echo "Raspberry Pi:"
	$(call go_build,linux,arm,6,Pi 1 / Zero / Zero W / CM1)
	$(call go_build,linux,arm,7,Pi 2 / Pi 3-5 mit 32-bit OS)
	$(call go_build,linux,arm64,,Pi 3-5 / Zero 2 W / CM3-CM5 mit 64-bit OS)

build-windows:
	@echo "Windows:"
	$(call go_build,windows,amd64,,x86_64)
	$(call go_build,windows,386,,x86 32-bit)
	$(call go_build,windows,arm64,,ARM 64-bit)

# SHA256-Summen ueber alle gebauten Binaries
checksums:
	@cd $(DIST) && \
		if command -v sha256sum >/dev/null 2>&1; then \
			sha256sum $(BINARY_NAME)-* > checksums.txt; \
		else \
			shasum -a 256 $(BINARY_NAME)-* > checksums.txt; \
		fi
	@echo "Checksums: $(DIST)/checksums.txt"

commit-build:
	@echo "Commit build: $(GIT)"
	@git add dist
	@git commit -m "Commit binaries for" || true
	@git push 

# Copy the default theme: make new-theme NAME=newdesign
new-theme:
	@test -n "$(NAME)" || { echo "Usage: make new-theme NAME=<theme>"; exit 1; }
	@test ! -d themes/$(NAME) || { echo "themes/$(NAME) already exists"; exit 1; }
	cp -r themes/default themes/$(NAME)
	@echo "Created themes/$(NAME) — activate with 'theme: $(NAME)' in config.yaml or 'go run . serve -theme $(NAME)'"

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
	rm -rf bin/ $(DIST)/ $(BINARY_NAME)

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