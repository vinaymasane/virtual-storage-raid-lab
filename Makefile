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

stop:
	sudo $(BIN) stop

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

vet:
	go vet ./...

tidy:
	go mod tidy

ci:
	go test ./...

coverage:
	tests/scripts/coverage.sh

integration:
	tests/scripts/run_all.sh

package:
	scripts/release.sh

docker:
	docker build -t raidlab-dev -f Dockerfile.dev .


.PHONY: all bootstrap build image mirror raid launch configure verify collect clean