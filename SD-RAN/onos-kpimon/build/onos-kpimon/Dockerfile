# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

ARG ONOS_BUILD_VERSION=undefined

FROM golang:1.20.7 AS build-env

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY go-fuzz/ /go/src/github.com/dvyukov/go-fuzz/
RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz-build && go build
COPY Makefile go.mod go.sum /go/src/github.com/onosproject/onos-kpimon/
COPY onos-kpimon.go /go/src/github.com/onosproject/onos-kpimon/
COPY pkg/ /go/src/github.com/onosproject/onos-kpimon/pkg/
COPY api/ /go/src/github.com/onosproject/onos-kpimon/api
COPY build/build-tools /go/src/github.com/onosproject/onos-kpimon/build/build-tools

RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz-build && cp go-fuzz-build /go/src/github.com/onosproject/onos-kpimon
RUN cd /go/src/github.com/onosproject/onos-kpimon && ./go-fuzz-build --coverdir=/go/src/github.com/onosproject/onos-kpimon/instrument

RUN cp -r /go/src/github.com/dvyukov/go-fuzz/go-fuzz-run/ /go/src/github.com/onosproject/onos-kpimon/go-fuzz-run/
RUN cd /go/src/github.com/onosproject/onos-kpimon/go-fuzz-run && go build
WORKDIR /go/src/github.com/onosproject/onos-kpimon/go-fuzz-run

RUN cp /go/src/github.com/onosproject/onos-kpimon/go-fuzz-run/go-fuzz-run /usr/local/bin/go-fuzz-run


EXPOSE 19999
EXPOSE 5150

ENTRYPOINT ["go-fuzz-run", "--file=/go/src/github.com/onosproject/onos-kpimon/gofuzzdep-fuzz.zip"]