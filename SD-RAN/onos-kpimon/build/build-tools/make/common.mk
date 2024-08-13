MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))
BIN     = $(CURDIR)/bin
BUILD 	= $(CURDIR)/build
OUT 	= $(BUILD)/_output
MK_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CUR_DIR := $(notdir $(patsubst %/,%,$(dir $(MK_PATH))))

GO      = go
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# Export environment variables if a .env file is present.
ifeq ($(ENV_EXPORTED),) # ENV vars not yet exported
ifneq ("$(wildcard .env)","")
sinclude .env
export $(shell [ -f .env ] && sed 's/=.*//' .env)
export ENV_EXPORTED=true
$(info Note — An .env file exists. Its contents have been exported as environment variables.)
endif
endif

# defaults if not set in the .env file

GO111MODULE?=on
CGO_ENABLED?=0
PRJ_NAME?=$(CUR_DIR)
PRJ_VERSION?=latest
DOCKERFILE?=.$(BUILD)/$(PRJ_NAME)/Dockerfile

.DEFAULT_GOAL := all

# Tools
$(OUT):
	@mkdir -p $@
$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)...)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLANGCI_LINT = $(BIN)/golangci-lint
$(BIN)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint

# Local coverage
GOCOV = $(BIN)/gocov
$(BIN)/gocov: PACKAGE=github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: PACKAGE=github.com/AlekSi/gocov-xml

# Travis coverage
OVERALLS = $(BIN)/overalls
$(BIN)/overalls: PACKAGE=github.com/go-playground/overalls

.PHONY: overalls
overalls: | $(OVERALLS) ; $(info $(M) running overalls) @ ## Common run overalls
	$Q $(OVERALLS) -project=$(CURDIR) -concurrency 2 -covermode=count -ignore=".git,vendor,models,tools"

GOVERALLS = $(BIN)/goveralls
$(BIN)/goveralls: PACKAGE=github.com/mattn/goveralls

.PHONY: build
build: $(BIN) ; $(info $(M) building executable...) @ ## Common build program binary
	$Q CGO_ENABLED=$(CGO_ENABLED) $(GO) build \
		-tags release \
		-ldflags '-X $(MODULE)/cmd.version=$(PRJ_VERSION) -X $(MODULE)/cmd.commit=$(VERSION) -X $(MODULE)/cmd.date=$(DATE)' \
		-o $(BIN)/$(basename $(MODULE)) main.go

.PHONY: all
all: license fmt lint test build ; $(info $(M) building all...) @ ## Common checks, format and build binary

.PHONY: docker-$(PRJ_NAME)
docker-$(PRJ_NAME): ; $(info $(M) building docker image...) @ ## Common build docker image
	docker build . -f $(DOCKERFILE) \
		--build-arg o=./bin/$(basename $(MODULE)) \
		-t onosproject/$(PRJ_NAME):$(PRJ_VERSION)

.PHONY: docker-login
docker-login: ; $(info $(M) docker login...) @ ## Common login to docker.io
	$Q @echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USER)" --password-stdin


# Tests

TEST_TARGETS := test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test
test-short:   ARGS=-short        ## Common run only short tests
test-verbose: ARGS=-v            ## Common run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Common run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
test: ; $(info $(M) running $(NAME:%=% )tests...) @ ## Common run tests
	$Q $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML     = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML    = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage
test-coverage: COVERAGE_DIR := $(OUT)/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: | $(GOCOV) $(GOCOVXML) fmt lint ; $(info $(M) running coverage tests...) @ ## Common run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)
	$Q $(GO) test \
		-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $(TESTPKGS) | \
					grep '^$(MODULE)/' | \
					tr '\n' ',' | sed 's/,$$//') \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile="$(COVERAGE_PROFILE)" $(TESTPKGS)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: | $(GOLANGCI_LINT) ; $(info $(M) running golangcli-lint...) @ ## Common run golangci-lint
	$Q $(GOLANGCI_LINT) run

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Common run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

.PHONY: tidy
tidy: test ; $(info $(M) modules tidy...) @ ## Common run test before and after go mod tidy
	$Q $(GO) mod tidy
	make test

# Misc

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
		
.PHONY: version
version: ; $(info $(M) getting version) @ ## Common get version info
	@echo $(VERSION)

.PHONY: vars
vars: 
	$(foreach V,$(sort $(.VARIABLES)),$(if $(filter-out environment% default automatic, $(origin $V)), $(warning $V=$($V))))


# aggregate targets, can create project specific same target using make ::

.PHONY: clean
clean:: ; $(info $(M) cleaning...)	@ ## Common clean
	@rm -rf $(BIN)/$(basename $(MODULE))
	@find . -name "coverage.*" | xargs rm -f
	@find . -name "profile.coverprofile" | xargs rm -f


.PHONY: clean-tools
clean-tools: ; $(info $(M) cleaning tools used by build...)	@ ## Common clean build tools
	@rm -rf $(BIN)
	@rm -rf $(BUILD)/build-tools