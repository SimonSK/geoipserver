FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git make bash

COPY . /geoip2-webapi

WORKDIR /geoip2-webapi

RUN make server

RUN echo "nobody:x:65534:65534::/:" > /etc_passwd

FROM alpine:latest

WORKDIR /

COPY --from=builder /geoip2-webapi/bin/geoip2-api-server /usr/local/bin/

COPY --from=builder /etc_passwd /etc/passwd

USER nobody

EXPOSE 8080

VOLUME ["/databases"]

ENTRYPOINT ["geoip2-api-server"]
