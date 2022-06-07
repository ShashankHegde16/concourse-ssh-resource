# build stage

FROM golang:1.8.1-alpine as builder

COPY . /go/src/stash.tools.deloitteinnovation.us/projects/CRE/repos/concourse-ssh-resource

RUN apk --no-cache add make=4.2.1-r0 && \
  cd /go/src/stash.tools.deloitteinnovation.us/projects/CRE/repos/concourse-ssh-resource && \
  make build-linux

WORKDIR /opt/resource

# release stage

FROM alpine:edge

COPY --from=builder /opt/resource /opt/resource
