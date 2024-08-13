.PHONY: build go-fuzz-build

build:
	cd .. && go build && cp go-fuzz-build build-test

go-fuzz-build: build
	chmod +x go-fuzz-build
	./go-fuzz-build  --coverdir=instrument