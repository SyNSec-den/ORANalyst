# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

VERSION := latest

build: # @HELP build the Go binaries and run all validations (default)
build:
	go build

docker-build: # @HELP build Docker image
	docker build -t localhost:5001/input-generator:${VERSION} .

images: # @HELP build all Docker images
images: docker-build

docker-push:
	docker push localhost:5001/input-generator:${VERSION}

docker-run:
	docker run -it --rm --name input-generator localhost:5001/input-generator:${VERSION}


install-ransim:
	helm install -n riab ran-simulator ../sdran-helm-charts/ran-simulator --values ../sdran-helm-charts/ran-simulator/values.yaml

delete-ransim:
	-helm uninstall -n riab ran-simulator

sender-image:
	docker build -t localhost:5001/input-generator:${VERSION} -f sender/Dockerfile .