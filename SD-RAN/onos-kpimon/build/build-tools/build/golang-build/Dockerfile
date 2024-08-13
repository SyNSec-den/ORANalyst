# SPDX-FileCopyrightText: 2022-present Intel Corporation
# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.19.0

COPY dep-*.sha256 .

RUN apt-get update && apt-get install -y protobuf-compiler
RUN mkdir release && \
    wget -q -O release/dep-$(go env GOOS)-$(go env GOARCH) https://github.com/golang/dep/releases/download/v0.5.4/dep-$(go env GOOS)-$(go env GOARCH) && \
    sha256sum -c dep-$(go env GOOS)-$(go env GOARCH).sha256 && \
    mv release/dep-$(go env GOOS)-$(go env GOARCH) /go/bin/dep && \
    chmod +x /go/bin/dep && \
    rm dep-*.sha256 && \
    mkdir -p /go/src/github.com/protocolbuffers && \
    git clone --branch main https://github.com/protocolbuffers/protobuf /go/src/github.com/protocolbuffers/protobuf

RUN export GO111MODULE=on && \
    go install github.com/golang/protobuf/protoc-gen-go@latest && \
    go install github.com/openconfig/ygot/generator@latest

RUN export GO111MODULE=off && \
    go get -u google.golang.org/protobuf/proto && \
    go get -u golang.org/x/lint/golint && \
    go get -u github.com/fzipp/gocyclo && \
    go get -u github.com/client9/misspell/cmd/misspell && \
    go get -u github.com/gordonklaus/ineffassign
WORKDIR "/go/src/github.com"

ENTRYPOINT ["make"]
