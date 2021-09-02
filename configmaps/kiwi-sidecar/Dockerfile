ARG GOLANG_VERSION=1.16.6-alpine3.14
ARG ALPINE_VERSION=3.14

FROM golang:${GOLANG_VERSION} AS builder

ADD . ${GOPATH}/src/

WORKDIR ${GOPATH}/src/
RUN go get && go build -o /go/bin/configmap-watcher main.go

FROM alpine:${ALPINE_VERSION}

ENV FILE_TO_WATCH /config/config.yaml

COPY --from=builder /go/bin/configmap-watcher /go/bin/configmap-watcher

ENTRYPOINT ["/go/bin/configmap-watcher"]