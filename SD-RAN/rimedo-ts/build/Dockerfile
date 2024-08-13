# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
# SPDX-FileCopyrightText: 2019-present Rimedo Labs
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.20.7 AS build-env

ENV GO111MODULE=on
ENV XAPPNAME=rimedo-ts

RUN apt update && apt install -y unzip
COPY go-fuzz/ /go/src/github.com/dvyukov/go-fuzz/
RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz-build && go build
COPY Makefile go.mod go.sum /go/src/github.com/rimedo-labs/rimedo-ts/
# COPY cmd/ /go/src/github.com/rimedo-labs/rimedo-ts/cmd/
COPY rimedo-ts.go /go/src/github.com/rimedo-labs/rimedo-ts/
# COPY vendor/ /go/src/github.com/rimedo-labs/rimedo-ts/vendor/
COPY pkg/ /go/src/github.com/rimedo-labs/rimedo-ts/pkg/
COPY build/build-tools /go/src/github.com/rimedo-labs/rimedo-ts/build/build-tools/
# COPY policytypes /go/src/github.com/rimedo-labs/rimedo-ts/policytypes/
# COPY schemas /go/src/github.com/rimedo-labs/rimedo-ts/schemas/

RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz-build && cp go-fuzz-build /go/src/github.com/rimedo-labs/rimedo-ts
RUN cd /go/src/github.com/rimedo-labs/rimedo-ts && ./go-fuzz-build --coverdir=/go/src/github.com/rimedo-labs/rimedo-ts/instrument
RUN unzip /go/src/github.com/rimedo-labs/rimedo-ts/gofuzzdep-fuzz.zip -d /go/src/github.com/rimedo-labs/rimedo-ts/gofuzzdep-fuzz

# FROM alpine:3.11
# RUN apk add libc6-compat

# USER nobody

RUN cp -r /go/src/github.com/dvyukov/go-fuzz/go-fuzz-run/ /go/src/github.com/rimedo-labs/rimedo-ts/go-fuzz-run/
RUN cd /go/src/github.com/rimedo-labs/rimedo-ts/go-fuzz-run && go build
WORKDIR go/src/github.com/rimedo-labs/rimedo-ts/go-fuzz-run

# COPY --from=build /go/src/github.com/rimedo-labs/rimedo-ts/build/_output/rimedo-ts /usr/local/bin/rimedo-ts
# COPY --from=build /go/src/github.com/rimedo-labs/rimedo-ts/schemas /data/schemas
# COPY --from=build /go/src/github.com/rimedo-labs/rimedo-ts/policytypes /data/policytypes

RUN echo $PATH
RUN cp /go/src/github.com/rimedo-labs/rimedo-ts/go-fuzz-run/go-fuzz-run /usr/local/bin/go-fuzz-run
# RUN cd /usr/local/bin/ && ls -l

EXPOSE 19999
EXPOSE 5150

# ENTRYPOINT ["go-fuzz-run", "--file=/go/src/github.com/rimedo-labs/rimedo-ts/gofuzzdep-fuzz.zip", \
#     "--trackdir=", "--wkdir=/go/src/github.com/rimedo-labs/rimedo-ts", "--coverfreq=10"]

ENTRYPOINT ["go-fuzz-run", "--file=/go/src/github.com/rimedo-labs/rimedo-ts/gofuzzdep-fuzz.zip", \
    "--trackdir=gopath/src/github.com/onosproject/rimedo-ts", "--wkdir=/go/src/github.com/rimedo-labs/rimedo-ts", "--coverfreq=10"]