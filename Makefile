SHELL = /bin/bash
export GO111MODULE=on

BIN_NAME = geoip2-api-server
OUTPUT_DIR = bin

.PHONY: build clean

clean:
	@rm -rf ${OUTPUT_DIR}

build: clean
	@mkdir -p ${OUTPUT_DIR}
	go build -v -ldflags="-s -w" -o ${OUTPUT_DIR}/${BIN_NAME} ./main/...
