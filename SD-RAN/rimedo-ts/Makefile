# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
# SPDX-FileCopyrightText: 2019-present Rimedo Labs
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build images
#GO111MODULE=on 

XAPPNAME=rimedo-ts
RIMEDO_TS_VERSION := latest

build:
	GOPRIVATE="github.com/onosproject/*" go build -o build/_output/$(XAPPNAME) ./cmd/$(XAPPNAME)

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

docker:
	@go mod vendor
	sudo docker build --network host -f build/Dockerfile -t onosproject/$(XAPPNAME):$(RIMEDO_TS_VERSION) .
	@rm -rf vendor

images: 
	docker build -f build/Dockerfile -t localhost:5001/$(XAPPNAME):$(RIMEDO_TS_VERSION) .
	docker run -d --name $(XAPPNAME) localhost:5001/$(XAPPNAME):$(RIMEDO_TS_VERSION)
	docker cp $(XAPPNAME):/go/src/github.com/rimedo-labs/rimedo-ts/gofuzzdep-fuzz/metadata docker-out/metadata
	docker stop $(XAPPNAME)
	docker rm $(XAPPNAME)

kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/$(XAPPNAME):${RIMEDO_TS_VERSION}

helmit-ts: integration-test-namespace # @HELP run PCI tests locally
	helmit test -n test ./cmd/rimedo-ts-test --timeout 30m --no-teardown \
			--suite ts

integration-tests: helmit-ts

test: build license
jenkins-test: build license

docker-login:
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif

docker-push-latest:
	docker push localhost:5001/$(XAPPNAME):latest

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION} onosproject/$(XAPPNAME)

jenkins-publish: jenkins-tools images docker-push-latest # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit

install-xapp:
	helm install -n riab $(XAPPNAME) ../sdran-helm-charts/$(XAPPNAME) --values ../sdran-helm-charts/$(XAPPNAME)/values.yaml


delete-xapp:
	-helm uninstall -n riab $(XAPPNAME)

deploy-xapp: images docker-push-latest install-xapp

docker-run:
	docker run -it --rm --name $(XAPPNAME) localhost:5001/$(XAPPNAME):${RIMEDO_TS_VERSION}


