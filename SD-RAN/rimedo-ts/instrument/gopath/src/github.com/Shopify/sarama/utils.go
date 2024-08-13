//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:1
)

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

type none struct{}

// make []int32 sortable so we can sort partition numbers
type int32Slice []int32

func (slice int32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:15
	_go_fuzz_dep_.CoverTab[107011]++
											return len(slice)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:16
	// _ = "end of CoverTab[107011]"
}

func (slice int32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:19
	_go_fuzz_dep_.CoverTab[107012]++
											return slice[i] < slice[j]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:20
	// _ = "end of CoverTab[107012]"
}

func (slice int32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:23
	_go_fuzz_dep_.CoverTab[107013]++
											slice[i], slice[j] = slice[j], slice[i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:24
	// _ = "end of CoverTab[107013]"
}

func dupInt32Slice(input []int32) []int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:27
	_go_fuzz_dep_.CoverTab[107014]++
											ret := make([]int32, 0, len(input))
											ret = append(ret, input...)
											return ret
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:30
	// _ = "end of CoverTab[107014]"
}

func withRecover(fn func()) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:33
	_go_fuzz_dep_.CoverTab[107015]++
											defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:34
		_go_fuzz_dep_.CoverTab[107017]++
												handler := PanicHandler
												if handler != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:36
			_go_fuzz_dep_.CoverTab[107018]++
													if err := recover(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:37
				_go_fuzz_dep_.CoverTab[107019]++
														handler(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:38
				// _ = "end of CoverTab[107019]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:39
				_go_fuzz_dep_.CoverTab[107020]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:39
				// _ = "end of CoverTab[107020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:39
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:39
			// _ = "end of CoverTab[107018]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:40
			_go_fuzz_dep_.CoverTab[107021]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:40
			// _ = "end of CoverTab[107021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:40
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:40
		// _ = "end of CoverTab[107017]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:41
	// _ = "end of CoverTab[107015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:41
	_go_fuzz_dep_.CoverTab[107016]++

											fn()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:43
	// _ = "end of CoverTab[107016]"
}

func safeAsyncClose(b *Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:46
	_go_fuzz_dep_.CoverTab[107022]++
											tmp := b
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:47
	_curRoutineNum149_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:47
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum149_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
		_go_fuzz_dep_.CoverTab[107023]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
			_go_fuzz_dep_.CoverTab[107024]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum149_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
			// _ = "end of CoverTab[107024]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
		withRecover(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:48
			_go_fuzz_dep_.CoverTab[107025]++
													if connected, _ := tmp.Connected(); connected {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:49
				_go_fuzz_dep_.CoverTab[107026]++
														if err := tmp.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:50
					_go_fuzz_dep_.CoverTab[107027]++
															Logger.Println("Error closing broker", tmp.ID(), ":", err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:51
					// _ = "end of CoverTab[107027]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:52
					_go_fuzz_dep_.CoverTab[107028]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:52
					// _ = "end of CoverTab[107028]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:52
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:52
				// _ = "end of CoverTab[107026]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:53
				_go_fuzz_dep_.CoverTab[107029]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:53
				// _ = "end of CoverTab[107029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:53
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:53
			// _ = "end of CoverTab[107025]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:54
		// _ = "end of CoverTab[107023]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:54
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:54
	// _ = "end of CoverTab[107022]"
}

// Encoder is a simple interface for any type that can be encoded as an array of bytes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:57
// in order to be sent as the key or value of a Kafka message. Length() is provided as an
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:57
// optimization, and must return the same as len() on the result of Encode().
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:60
type Encoder interface {
	Encode() ([]byte, error)
	Length() int
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:68
// StringEncoder implements the Encoder interface for Go strings so that they can be used
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:68
// as the Key or Value in a ProducerMessage.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:70
type StringEncoder string

func (s StringEncoder) Encode() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:72
	_go_fuzz_dep_.CoverTab[107030]++
											return []byte(s), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:73
	// _ = "end of CoverTab[107030]"
}

func (s StringEncoder) Length() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:76
	_go_fuzz_dep_.CoverTab[107031]++
											return len(s)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:77
	// _ = "end of CoverTab[107031]"
}

// ByteEncoder implements the Encoder interface for Go byte slices so that they can be used
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:80
// as the Key or Value in a ProducerMessage.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:82
type ByteEncoder []byte

func (b ByteEncoder) Encode() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:84
	_go_fuzz_dep_.CoverTab[107032]++
											return b, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:85
	// _ = "end of CoverTab[107032]"
}

func (b ByteEncoder) Length() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:88
	_go_fuzz_dep_.CoverTab[107033]++
											return len(b)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:89
	// _ = "end of CoverTab[107033]"
}

// bufConn wraps a net.Conn with a buffer for reads to reduce the number of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:92
// reads that trigger syscalls.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:94
type bufConn struct {
	net.Conn
	buf	*bufio.Reader
}

func newBufConn(conn net.Conn) *bufConn {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:99
	_go_fuzz_dep_.CoverTab[107034]++
											return &bufConn{
		Conn:	conn,
		buf:	bufio.NewReader(conn),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:103
	// _ = "end of CoverTab[107034]"
}

func (bc *bufConn) Read(b []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:106
	_go_fuzz_dep_.CoverTab[107035]++
											return bc.buf.Read(b)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:107
	// _ = "end of CoverTab[107035]"
}

// KafkaVersion instances represent versions of the upstream Kafka broker.
type KafkaVersion struct {
	// it's a struct rather than just typing the array directly to make it opaque and stop people
	// generating their own arbitrary versions
	version [4]uint
}

func newKafkaVersion(major, minor, veryMinor, patch uint) KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:117
	_go_fuzz_dep_.CoverTab[107036]++
											return KafkaVersion{
		version: [4]uint{major, minor, veryMinor, patch},
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:120
	// _ = "end of CoverTab[107036]"
}

// IsAtLeast return true if and only if the version it is called on is
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:123
// greater than or equal to the version passed in:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:123
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:123
//	V1.IsAtLeast(V2) // false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:123
//	V2.IsAtLeast(V1) // true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:127
func (v KafkaVersion) IsAtLeast(other KafkaVersion) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:127
	_go_fuzz_dep_.CoverTab[107037]++
											for i := range v.version {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:128
		_go_fuzz_dep_.CoverTab[107039]++
												if v.version[i] > other.version[i] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:129
			_go_fuzz_dep_.CoverTab[107040]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:130
			// _ = "end of CoverTab[107040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:131
			_go_fuzz_dep_.CoverTab[107041]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:131
			if v.version[i] < other.version[i] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:131
				_go_fuzz_dep_.CoverTab[107042]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:132
				// _ = "end of CoverTab[107042]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
				_go_fuzz_dep_.CoverTab[107043]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
				// _ = "end of CoverTab[107043]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
			// _ = "end of CoverTab[107041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:133
		// _ = "end of CoverTab[107039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:134
	// _ = "end of CoverTab[107037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:134
	_go_fuzz_dep_.CoverTab[107038]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:135
	// _ = "end of CoverTab[107038]"
}

// Effective constants defining the supported kafka versions.
var (
	V0_8_2_0	= newKafkaVersion(0, 8, 2, 0)
	V0_8_2_1	= newKafkaVersion(0, 8, 2, 1)
	V0_8_2_2	= newKafkaVersion(0, 8, 2, 2)
	V0_9_0_0	= newKafkaVersion(0, 9, 0, 0)
	V0_9_0_1	= newKafkaVersion(0, 9, 0, 1)
	V0_10_0_0	= newKafkaVersion(0, 10, 0, 0)
	V0_10_0_1	= newKafkaVersion(0, 10, 0, 1)
	V0_10_1_0	= newKafkaVersion(0, 10, 1, 0)
	V0_10_1_1	= newKafkaVersion(0, 10, 1, 1)
	V0_10_2_0	= newKafkaVersion(0, 10, 2, 0)
	V0_10_2_1	= newKafkaVersion(0, 10, 2, 1)
	V0_10_2_2	= newKafkaVersion(0, 10, 2, 2)
	V0_11_0_0	= newKafkaVersion(0, 11, 0, 0)
	V0_11_0_1	= newKafkaVersion(0, 11, 0, 1)
	V0_11_0_2	= newKafkaVersion(0, 11, 0, 2)
	V1_0_0_0	= newKafkaVersion(1, 0, 0, 0)
	V1_0_1_0	= newKafkaVersion(1, 0, 1, 0)
	V1_0_2_0	= newKafkaVersion(1, 0, 2, 0)
	V1_1_0_0	= newKafkaVersion(1, 1, 0, 0)
	V1_1_1_0	= newKafkaVersion(1, 1, 1, 0)
	V2_0_0_0	= newKafkaVersion(2, 0, 0, 0)
	V2_0_1_0	= newKafkaVersion(2, 0, 1, 0)
	V2_1_0_0	= newKafkaVersion(2, 1, 0, 0)
	V2_1_1_0	= newKafkaVersion(2, 1, 1, 0)
	V2_2_0_0	= newKafkaVersion(2, 2, 0, 0)
	V2_2_1_0	= newKafkaVersion(2, 2, 1, 0)
	V2_2_2_0	= newKafkaVersion(2, 2, 2, 0)
	V2_3_0_0	= newKafkaVersion(2, 3, 0, 0)
	V2_3_1_0	= newKafkaVersion(2, 3, 1, 0)
	V2_4_0_0	= newKafkaVersion(2, 4, 0, 0)
	V2_4_1_0	= newKafkaVersion(2, 4, 1, 0)
	V2_5_0_0	= newKafkaVersion(2, 5, 0, 0)
	V2_5_1_0	= newKafkaVersion(2, 5, 1, 0)
	V2_6_0_0	= newKafkaVersion(2, 6, 0, 0)
	V2_6_1_0	= newKafkaVersion(2, 6, 1, 0)
	V2_6_2_0	= newKafkaVersion(2, 6, 2, 0)
	V2_6_3_0	= newKafkaVersion(2, 6, 3, 0)
	V2_7_0_0	= newKafkaVersion(2, 7, 0, 0)
	V2_7_1_0	= newKafkaVersion(2, 7, 1, 0)
	V2_7_2_0	= newKafkaVersion(2, 7, 2, 0)
	V2_8_0_0	= newKafkaVersion(2, 8, 0, 0)
	V2_8_1_0	= newKafkaVersion(2, 8, 1, 0)
	V3_0_0_0	= newKafkaVersion(3, 0, 0, 0)
	V3_1_0_0	= newKafkaVersion(3, 1, 0, 0)

	SupportedVersions	= []KafkaVersion{
		V0_8_2_0,
		V0_8_2_1,
		V0_8_2_2,
		V0_9_0_0,
		V0_9_0_1,
		V0_10_0_0,
		V0_10_0_1,
		V0_10_1_0,
		V0_10_1_1,
		V0_10_2_0,
		V0_10_2_1,
		V0_10_2_2,
		V0_11_0_0,
		V0_11_0_1,
		V0_11_0_2,
		V1_0_0_0,
		V1_0_1_0,
		V1_0_2_0,
		V1_1_0_0,
		V1_1_1_0,
		V2_0_0_0,
		V2_0_1_0,
		V2_1_0_0,
		V2_1_1_0,
		V2_2_0_0,
		V2_2_1_0,
		V2_2_2_0,
		V2_3_0_0,
		V2_3_1_0,
		V2_4_0_0,
		V2_4_1_0,
		V2_5_0_0,
		V2_5_1_0,
		V2_6_0_0,
		V2_6_1_0,
		V2_6_2_0,
		V2_7_0_0,
		V2_7_1_0,
		V2_8_0_0,
		V2_8_1_0,
		V3_0_0_0,
		V3_1_0_0,
	}
	MinVersion	= V0_8_2_0
	MaxVersion	= V3_1_0_0
	DefaultVersion	= V1_0_0_0
)

// ParseKafkaVersion parses and returns kafka version or error from a string
func ParseKafkaVersion(s string) (KafkaVersion, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:235
	_go_fuzz_dep_.CoverTab[107044]++
											if len(s) < 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:236
		_go_fuzz_dep_.CoverTab[107048]++
												return DefaultVersion, fmt.Errorf("invalid version `%s`", s)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:237
		// _ = "end of CoverTab[107048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:238
		_go_fuzz_dep_.CoverTab[107049]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:238
		// _ = "end of CoverTab[107049]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:238
	// _ = "end of CoverTab[107044]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:238
	_go_fuzz_dep_.CoverTab[107045]++
											var major, minor, veryMinor, patch uint
											var err error
											if s[0] == '0' {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:241
		_go_fuzz_dep_.CoverTab[107050]++
												err = scanKafkaVersion(s, `^0\.\d+\.\d+\.\d+$`, "0.%d.%d.%d", [3]*uint{&minor, &veryMinor, &patch})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:242
		// _ = "end of CoverTab[107050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:243
		_go_fuzz_dep_.CoverTab[107051]++
												err = scanKafkaVersion(s, `^\d+\.\d+\.\d+$`, "%d.%d.%d", [3]*uint{&major, &minor, &veryMinor})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:244
		// _ = "end of CoverTab[107051]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:245
	// _ = "end of CoverTab[107045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:245
	_go_fuzz_dep_.CoverTab[107046]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:246
		_go_fuzz_dep_.CoverTab[107052]++
												return DefaultVersion, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:247
		// _ = "end of CoverTab[107052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:248
		_go_fuzz_dep_.CoverTab[107053]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:248
		// _ = "end of CoverTab[107053]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:248
	// _ = "end of CoverTab[107046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:248
	_go_fuzz_dep_.CoverTab[107047]++
											return newKafkaVersion(major, minor, veryMinor, patch), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:249
	// _ = "end of CoverTab[107047]"
}

func scanKafkaVersion(s string, pattern string, format string, v [3]*uint) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:252
	_go_fuzz_dep_.CoverTab[107054]++
											if !regexp.MustCompile(pattern).MatchString(s) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:253
		_go_fuzz_dep_.CoverTab[107056]++
												return fmt.Errorf("invalid version `%s`", s)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:254
		// _ = "end of CoverTab[107056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:255
		_go_fuzz_dep_.CoverTab[107057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:255
		// _ = "end of CoverTab[107057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:255
	// _ = "end of CoverTab[107054]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:255
	_go_fuzz_dep_.CoverTab[107055]++
											_, err := fmt.Sscanf(s, format, v[0], v[1], v[2])
											return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:257
	// _ = "end of CoverTab[107055]"
}

func (v KafkaVersion) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:260
	_go_fuzz_dep_.CoverTab[107058]++
											if v.version[0] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:261
		_go_fuzz_dep_.CoverTab[107060]++
												return fmt.Sprintf("0.%d.%d.%d", v.version[1], v.version[2], v.version[3])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:262
		// _ = "end of CoverTab[107060]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:263
		_go_fuzz_dep_.CoverTab[107061]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:263
		// _ = "end of CoverTab[107061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:263
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:263
	// _ = "end of CoverTab[107058]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:263
	_go_fuzz_dep_.CoverTab[107059]++

											return fmt.Sprintf("%d.%d.%d", v.version[0], v.version[1], v.version[2])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:265
	// _ = "end of CoverTab[107059]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:266
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/utils.go:266
var _ = _go_fuzz_dep_.CoverTab
