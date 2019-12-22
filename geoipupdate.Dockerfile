FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git make bash

COPY . /geoip2-webapi

WORKDIR /geoip2-webapi

RUN make updater

FROM alpine:latest

WORKDIR /

COPY --from=builder /geoip2-webapi/bin/geoipupdate /usr/local/bin/

ADD ./conf/GeoIP.conf.default /

VOLUME ["/databases"]

ENTRYPOINT ["geoipupdate"]
