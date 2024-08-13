package asn1_mutator

// #cgo CFLAGS: -I/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm
// #cgo LDFLAGS: -L/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm/build -lkpm -lm
// #include "../../../mutator/e2ap/mutator.h"
import "C"
import "fmt"

var (
	lb      = 0
	ub      = 10240
	timeout = 5
)

func MutateStructWrapper(data []byte) ([]byte, error) {
	i := 0
	for i < timeout {
		i++
		data, err := mutateMessageFork(data)
		if err != nil {
			return make([]byte, 0), err
		}
		if len(data) > lb && len(data) < ub {
			return data, nil
		}
	}
	return make([]byte, 0), fmt.Errorf("timeout reached")
}
