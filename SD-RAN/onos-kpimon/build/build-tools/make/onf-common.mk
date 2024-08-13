DOCKER_REPOSITORY ?= onosproject/
KIND_CLUSTER_NAME ?= kind

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 15m

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go install github.com/jstemmer/go-junit-report@v1.0.0 && go install github.com/t-yuki/gocover-cobertura@master

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version | grep 1.52.2 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b `go env GOPATH`/bin v1.52.2

reuse-tool: # @HELP install reuse if not present
	command -v reuse || python3 -m pip install --user reuse

twine: # @HELP install twine if not present
	python3 -m twine --version || ( python3 -m pip install --upgrade pip && pip3 install twine )

license: reuse-tool # @HELP run license checks
	reuse lint

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git.
	./build/build-tools/bump-onos-deps ${VERSION}

integration-test-namespace: # @HELP create the kubernetes namespace used by the integration tests
	(kubectl delete ns test || exit 0) && kubectl create ns test

docker-login: # @HELP log in to docker hub
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif

clean:: # @HELP cleans the downloaded build tools directory
	rm -rf ./build/build-tools
