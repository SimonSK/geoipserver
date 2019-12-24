SHELL = /bin/bash
export GO111MODULE=on

BIN_SERVER = geoip2-api-server
OUTPUT_DIR = bin

.PHONY: clean server

clean:
	@rm -rf ${OUTPUT_DIR}

server: clean
	@mkdir -p ${OUTPUT_DIR}
	go build -v -ldflags="-s -w" -o ${OUTPUT_DIR}/${BIN_SERVER} ./cmd/${BIN_SERVER}/...

docker-image: clean
	docker build -t ${BIN_SERVER} -f ${BIN_SERVER}.Dockerfile .
