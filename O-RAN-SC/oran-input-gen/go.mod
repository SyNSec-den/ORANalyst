module cgo_mutator

go 1.21.1

replace github.com/dvyukov/go-fuzz => /home/tianchang/Desktop/proj/oaic-input-gen/go-fuzz

require (
	github.com/dvyukov/go-fuzz v0.0.0-20231019021653-5581da83c52f
	github.com/edsrzf/mmap-go v1.1.0
	github.com/google/uuid v1.3.1
	google.golang.org/protobuf v1.31.0
)

require golang.org/x/sys v0.8.0 // indirect
