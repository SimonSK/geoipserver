SHELL = /bin/bash
export GO111MODULE=on

BIN_SERVER = geoip2-api-server
BIN_UPDATER = geoipupdate
OUTPUT_DIR = bin

.PHONY: clean updater server all

clean:
	@rm -rf ${OUTPUT_DIR}

clean-server:
	@rm -f ${OUTPUT_DIR}/${BIN_SERVER}

clean-updater:
	@rm -f ${OUTPUT_DIR}/${BIN_UPDATER}

server: clean-server
	@mkdir -p ${OUTPUT_DIR}
	go build -v -ldflags="-s -w" -o ${OUTPUT_DIR}/${BIN_SERVER} ./cmd/${BIN_SERVER}/...

updater: clean-updater
	@mkdir -p ${OUTPUT_DIR}
	go build -v -ldflags="-s -w" -o ${OUTPUT_DIR}/${BIN_UPDATER} ./cmd/${BIN_UPDATER}/...

all: clean server updater