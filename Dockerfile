# build stage

FROM dcil-docker-release.art.tools.deloitteinnovation.us/golang:1.8.1-alpine as builder

COPY . /go/src/concourse-ssh-resource

RUN apk --no-cache add make=4.2.1-r0 && \
  cd /go/src/concourse-ssh-resource && \
  make build-linux

WORKDIR /opt/resource

# release stage

FROM dcil-docker-release.art.tools.deloitteinnovation.us/alpine:edge

COPY --from=builder /opt/resource /opt/resource
