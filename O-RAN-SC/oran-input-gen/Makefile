all:
	CGO_CFLAGS="-Wno-error -DASN_DISABLE_OER_SUPPORT=1" CGO_CXXFLAGS="-Wno-error -DASN_DISABLE_OER_SUPPORT=1" go build -o cgo_mutator

wrapper:
	gcc -Wno-error -DASN_DISABLE_OER_SUPPORT=1 wrapper.c -I/home/tianchang/Desktop/e2_bindings/E2SM-KPM -L/home/tianchang/Desktop/asn1c-v771/kpm/build_mutate -lKPM

fuzzer:
	cd mutator/e2ap && go build
	cd mutator/kpm && go build
	cd go-fuzz/go-fuzz && make
	cp go-fuzz/go-fuzz/go-fuzz ./go-fuzz-bin
	make
	mv cgo_mutator sonar.exe
	zip fuzz.zip sonar.exe metadata

run-fuzzer:
	./go-fuzz-bin --procs=1 --bin=fuzz.zip --timeout=10 -v=5
	# ./go-fuzz-bin --procs=1 --bin=fuzz.zip --timeout=10

log-fuzzer:
	./go-fuzz-bin --procs=1 --bin=fuzz.zip --timeout=10 -v=5 2>&1 | tee fuzz.log