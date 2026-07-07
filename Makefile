build:
	go build -o bin/raidlab cmd/raidlab/main.go

bootstrap:
	sudo bash bootstrap/preinstall_host.sh
	bash bootstrap/validate_host.sh

image:
	./bin/raidlab build

mirror:
	./bin/raidlab mirror

raid:
	sudo ./bin/raidlab raid

launch:
	sudo ./bin/raidlab launch

configure:
	sudo ./bin/raidlab configure

test:
	go test ./test/... -v

full:
	make bootstrap
	make build
	make image
	make mirror
	make raid
	make launch
	make configure
	make test