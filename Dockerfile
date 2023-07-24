# Build the manager binary
FROM golang:1.19-buster as builder
ARG TARGETOS
ARG TARGETARCH

RUN echo "deb http://mirrors.aliyun.com/debian/ buster main non-free contrib" > /etc/apt/sources.list && \
    echo "deb http://mirrors.aliyun.com/debian-security buster/updates main" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib" >> /etc/apt/sources.list

RUN apt-get -y update && apt-get -y install upx

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
# COPY main.go main.go
# COPY common/ common/
# COPY static/ static/
COPY . .

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o app main.go

RUN upx /workspace/app

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot

# Use alpine as minimal base image to package the manager binary
FROM alpine:3.9.2
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shangha

WORKDIR /
COPY --from=builder /workspace/app .
COPY ./config.yaml .
USER 65532:65532

ENTRYPOINT ["/app"]
