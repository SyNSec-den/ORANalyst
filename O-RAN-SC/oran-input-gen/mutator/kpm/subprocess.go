package main

// #cgo CFLAGS: -I/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm
// #cgo LDFLAGS: -L/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm/build -lkpm -lm
// #include "mutator.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/edsrzf/mmap-go"
)

func main() {
	C.setRandSeed()
	sharedFileName := os.Args[1]

	// The subprocess received file descriptors 3, 4, and 5 for the shared file, input pipe, and output pipe, respectively.
	sharedFile := os.NewFile(3, sharedFileName)
	pipeIn := os.NewFile(4, "pipeIn")
	pipeOut := os.NewFile(5, "pipeOut")

	// Create a memory-mapped area from the file.
	mmap, err := mmap.Map(sharedFile, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer mmap.Unmap()

	// Read the length of the data from the input pipe.
	var length int
	fmt.Fscanf(pipeIn, "%d", &length) // This reads from the pipe.

	// Close the input pipe file descriptor once we're done.
	pipeIn.Close()

	// Perform the data processing.
	data := make([]byte, length)
	copy(data, mmap[:length])
	processedData := MutateMessage(data) // Just an example operation.
	// fmt.Printf("Processed data: %v\n", processedData)

	// Write processed data back to shared memory.
	copy(mmap, processedData)

	// Send back the length of the processed data through the output pipe.
	fmt.Fprintln(pipeOut, len(processedData))

	// Close the output pipe file descriptor once we're done.
	pipeOut.Close()
}

func MutateMessage(data []byte) []byte {
	bufPtr := C.malloc(C.get_sizeof_void_ptr())
	defer C.free(bufPtr)
	res := C.mutateMessage(
		unsafe.Pointer(&data[0]),
		C.long(len(data)),
		(*unsafe.Pointer)(bufPtr),
	)

	if res < 0 {
		return nil
	}

	actualBufPtr := *(**C.char)(bufPtr)
	buf := C.GoBytes(unsafe.Pointer(actualBufPtr), res)
	C.free(unsafe.Pointer(actualBufPtr))

	return buf
}
