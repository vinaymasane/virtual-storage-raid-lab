APP=raidlab

BIN=bin/$(APP)

GO=go

all: build

bootstrap:
	sudo ./bootstrap/preinstall_host.sh

build:
	mkdir -p bin
	$(GO) mod tidy
	$(GO) build -o $(BIN) ./cmd/raidlab

image:
	$(BIN) build

mirror:
	sudo $(BIN) mirror

raid:
	sudo $(BIN) raid

launch:
	sudo $(BIN) launch

configure:
	sudo $(BIN) configure

verify:
	sudo $(BIN) verify

collect:
	sudo ./bootstrap/collect_host.sh

clean:
	sudo ./bootstrap/cleanup_host.sh
	rm -rf bin output artifacts

fmt:
	go fmt ./...

test:
	go test ./... -v

lint:
	go vet ./...

.PHONY: all bootstrap build image mirror raid launch configure verify collect clean