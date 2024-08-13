package main

// #cgo CFLAGS: -I/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm
// #cgo LDFLAGS: -L/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm/build -lkpm -lm
// #include "gen.h"
import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

func main() {
	mutate("e2sm_message_cucp_ue.bin")
}

func mutate(filename string) {

	// Read the original data
	data, err := readFile(filename)
	if err != nil {
		panic(err)
	}

	// Decode and print the original message
	decodeAndPrint(data)

	// Mutate the message
	mutateMessage(data)

}

// Call C function to mutate the message
func mutateMessage(data []byte) {
	var newBuffer unsafe.Pointer

	// Call mutateMessage C function
	result := C.mutateMessage(unsafe.Pointer(&data[0]), C.long(len(data)), &newBuffer)
	if result != 0 {
		return
	}

	C.decode_msg_print(newBuffer, C.size_t(len(data)))
}

// Read file and return its contents
func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Call C function to decode and print the message
func decodeAndPrint(data []byte) {
	C.decode_msg_print(unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

func gen() {
	C.setRandSeed()
	i := 0
	for {
		if i >= 10 {
			break
		}
		buf := make([]byte, 100000)
		var msg *C.E2SM_KPM_IndicationMessage_t
		ret := C.gen_msg(&msg)
		if ret != 0 {
			fmt.Printf("gen_msg failed: %v\n", ret)
			continue
		}

		len := C.encode_msg(msg, unsafe.Pointer(&buf[0]))

		if len < 0 {
			fmt.Printf("encode_msg failed: %v\n", len)
			continue
		}

		file, err := os.Create(fmt.Sprintf("input_%v", i))
		if err != nil {
			fmt.Printf("create file failed: %v\n", err)
			continue
		}
		defer file.Close()

		_, err = file.Write(buf[:len])
		if err != nil {
			fmt.Printf("write file failed: %v\n", err)
			continue
		}

	}
}
