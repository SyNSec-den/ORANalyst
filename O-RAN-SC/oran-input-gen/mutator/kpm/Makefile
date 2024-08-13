all:
	CGO_CFLAGS="-Wno-error -DASN_DISABLE_OER_SUPPORT=1" CGO_CXXFLAGS="-Wno-error -DASN_DISABLE_OER_SUPPORT=1" go build

wrapper:
	gcc -Wno-error -DASN_DISABLE_OER_SUPPORT=1 wrapper.c -I/home/tianchang/Desktop/e2_bindings/E2SM-KPM -L/home/tianchang/Desktop/asn1c-v771/kpm/build_mutate -lKPM