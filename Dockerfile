FROM quay.io/projectquay/golang:1.20 as builder

ARG TARGETARCH
ARG TARGETOS

WORKDIR /go/src/app
COPY . .
RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH