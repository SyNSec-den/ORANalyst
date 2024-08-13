# SPDX-FileCopyrightText: 2022-present Intel Corporation
# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.19.0

RUN apt-get update && \
    apt-get install -y unzip python3 python3-pip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip && \
    unzip -o protoc-3.19.4-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-3.19.4-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-3.19.4-linux-x86_64.zip

RUN export GO111MODULE=on && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0

RUN export GO111MODULE=off && \
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -u github.com/gogo/protobuf/proto && \
    go get -u github.com/gogo/protobuf/gogoproto && \
    go get -u github.com/gogo/protobuf/protoc-gen-gofast && \
    go get -u github.com/gogo/protobuf/protoc-gen-gogo && \
    go get -u github.com/gogo/protobuf/protoc-gen-gogofast && \
    go get -u github.com/gogo/protobuf/protoc-gen-gogofaster && \
    go get -u github.com/gogo/protobuf/protoc-gen-gogoslick && \
    go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/envoyproxy/protoc-gen-validate && \
    go get -u github.com/favadi/protoc-go-inject-tag

RUN mkdir -p /go/src/github.com/google /go/src/github.com/p4lang && \
    git clone --branch main https://github.com/google/protobuf /go/src/github.com/google/protobuf && \
    git clone --branch master https://github.com/openconfig/gnmi /go/src/github.com/openconfig/gnmi && \
    git clone --branch main https://github.com/p4lang/p4runtime /go/src/github.com/p4lang/p4runtime && \
    git clone --branch main https://github.com/googleapis/api-common-protos /go/src/github.com/googleapis/api-common-protos && \
    mkdir -p /go/src/github.com/ &&\
    wget "https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64" --quiet && \
    mv protoc-gen-grpc-web-1.2.1-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web && \
    chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN python3 -m pip install -U build markupsafe==2.0.1 betterproto[compiler]==2.0.0b2

WORKDIR "/go/src/github.com/"

ENTRYPOINT ["/bin/bash"]
