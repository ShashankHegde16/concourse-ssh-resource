FROM golang:alpine as builder
COPY . /go/src/github.com/henry40408/concourse-ssh-resource
ENV CGO_ENABLED 0

RUN apk --no-cache add make && \
    cd /go/src/github.com/henry40408/concourse-ssh-resource && \
    make build-linux && \
    rm -r /go && \
    apk --no-cache del make

WORKDIR /opt/resource

################ thats our production image

FROM alpine:edge AS resource
RUN apk --no-cache add \
  bash \
  curl \
  gzip \
  jq \
  tar \
  openssl
COPY --from=builder /opt/resource /opt/resource
