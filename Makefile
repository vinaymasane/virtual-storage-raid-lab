# ============================================================================
# virtual-storage-raid-lab
# Makefile
# ============================================================================

PROJECT := raidlab
BIN  := bin/$(PROJECT)

GO := go

.PHONY: help bootstrap build image

help:
	@echo ""
	@echo "Available Targets"
	@echo "-----------------"
	@echo "make bootstrap"
	@echo "make build"
	@echo "make image"
	@echo "make mirror"
	@echo "make raid"
	@echo "make launch"
	@echo "make configure"
	@echo "make test"
	@echo "make collect"
	@echo "make clean"

bootstrap:
	@chmod +x bootstrap/*.sh
	@echo "[INFO ] Running pre-installation steps for the host machine..."
	@
	sudo ./bootstrap/preinstall_host.sh
	@echo "[INFO ] Pre-installation steps completed successfully."
	@echo "[INFO ] Running bootstrap process..."
	@$(BIN) -config configs/config.yaml bootstrap
	@echo "[INFO ] Bootstrap process completed successfully."

build:
	@echo "[INFO ] Building $(PROJECT) binary..."
	@mkdir -p bin
	$(GO) mod tidy
	$(GO) build -o $(BIN) ./cmd/$(PROJECT)
	@echo "[INFO ] Build completed successfully."

image: build
	$(BIN) image

# ============================================================================
# Storage
# ============================================================================

.PHONY: mirror raid

mirror:
	$(BIN) mirror

raid:
	sudo $(BIN) raid

# ============================================================================
# Virtual Machine
# ============================================================================

.PHONY: launch stop reboot status

launch:
	sudo $(BIN) launch

stop:
	sudo $(BIN) stop

reboot:
	sudo $(BIN) reboot

status:
	$(BIN) status

# ============================================================================
# Configuration
# ============================================================================

.PHONY: configure verify

configure:
	sudo $(BIN) configure

verify:
	$(BIN) verify

# ============================================================================
# Testing : Unit , Integration Tests
# ============================================================================

.PHONY: test

test:
	$(BIN) test

# ============================================================================
# Unit / Integration Tests
# ============================================================================

.PHONY: unit integration

unit:
	$(GO) test ./tests/unit/... -v

integration:
	$(GO) test ./tests/integration/... -v

# ============================================================================
# Code Quality
# ============================================================================

.PHONY: fmt vet lint coverage

fmt:
	@echo "[INFO ] Running go fmt..."
	$(GO) fmt ./...

vet:
	@echo "[INFO ] Running go vet..."
	$(GO) vet ./...

lint:
	@echo "[INFO ] Running golangci-lint..."
	@golangci-lint run ./...

tidy:
	@echo "[INFO ] Running go mod tidy..."
	$(GO) mod tidy

coverage:
	@echo "[INFO ] Running coverage tests..."
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out -o coverage.html

# ============================================================================
# Artifacts
# ============================================================================

.PHONY: collect

collect:
	@echo "[INFO ] Collecting artifacts..."
	$(BIN) collect

# ============================================================================
# Cleanup
# ============================================================================

.PHONY: clean

clean:
	sudo ./bootstrap/cleanup_host.sh
	@rm -rf bin output artifacts coverage*

## Check target runs all the checks: formatting, vetting, linting, and unit tests
check:
	make fmt
	make vet
	make lint
	make unit

## Continuous Integration target runs all checks and integration tests
ci:
	make check
	make integration

## Package the project for release
package:
	scripts/release.sh

## Build the Docker image for development
docker:
	docker build -t raidlab-dev -f Dockerfile.dev .


.PHONY: all bootstrap build image mirror raid launch configure verify collect clean