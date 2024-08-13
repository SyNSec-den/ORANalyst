# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build license

ONOS_BUILD_VERSION := latest

all:
	cat README.md

include ./make/onf-common.mk

golang-build-docker: # @HELP build golang-build Docker image
	docker build -t onosproject/golang-build:${ONOS_BUILD_VERSION} build/golang-build

protoc-go-docker: # @HELP build protoc-go Docker image
	docker build -t onosproject/protoc-go:${ONOS_BUILD_VERSION} build/protoc-go

publish: # @HELP publish version on github and dockerhub
	./publish-version ${VERSION} onosproject/protoc-go onosproject/golang-build

images: # @HELP create docker images
images: protoc-go-docker golang-build-docker

clean:: # @HELP remove all the build artifacts
	rm -rf ./web/onos-gui/dist

jenkins-test: # @HELP jenkins verify target
jenkins-test: jenkins-tools test
	TEST_PACKAGES="NONE" ./build/jenkins/make-unit

test: # @HELP testing target
test: images linters license

jenkins-publish: # @HELP jenkins publishing target
jenkins-publish: # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	./release-merge-commit