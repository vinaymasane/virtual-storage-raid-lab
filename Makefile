APP=raidlab

BIN=bin/$(APP)

GO=go

## Targets
all: build

## Bootstrap the host environment
bootstrap:
	sudo ./bootstrap/preinstall_host.sh

## Build the raidlab binary
build:
	mkdir -p bin
	$(GO) mod tidy
	$(GO) build -o $(BIN) ./cmd/raidlab

## Build the image using packer
image:
	$(BIN) build

## Build Mirror raw image
mirror:
	sudo $(BIN) mirror

## Build RAID lab environment
raid:
	sudo $(BIN) raid

## Launch the RAID lab environment
launch:
	sudo $(BIN) launch
## Stop the RAID lab environment
stop:
	sudo $(BIN) stop

## Configure the RAID lab environment
configure:
	sudo $(BIN) configure

## Verify the RAID lab environment
verify:
	sudo $(BIN) verify

## Collect artifacts from the RAID lab environment
collect:
	sudo ./bootstrap/collect_host.sh

## Clean up the RAID lab environment
clean:
	sudo ./bootstrap/cleanup_host.sh
	rm -rf bin output artifacts

### Development targets
## Format the Go code
fmt:
	go fmt ./...

## Run Go vet to check for potential issues
vet:
	go vet ./...

## Run Go lint to analyze the code for style and potential errors
lint:
	golangci-lint run

## Run Go mod tidy to clean up the go.mod and go.sum files
tidy:
	go mod tidy

## Run all tests in the project
ci:
	go test ./...

## Run tests with coverage analysis
coverage:
	tests/scripts/coverage.sh

## Run integration tests
integration:
	tests/scripts/run_all.sh

## Package the project for release
package:
	scripts/release.sh

## Build the Docker image for development
docker:
	docker build -t raidlab-dev -f Dockerfile.dev .


.PHONY: all bootstrap build image mirror raid launch configure verify collect clean