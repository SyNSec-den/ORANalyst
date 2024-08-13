//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:17
)

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/binary"
	"io"
	"math/big"
	"strings"
	"unicode"

	"gopkg.in/square/go-jose.v2/json"
)

// Helper function to serialize known-good objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:32
// Precondition: value is not a nil pointer.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:34
func mustSerializeJSON(value interface{}) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:34
	_go_fuzz_dep_.CoverTab[189291]++
											out, err := json.Marshal(value)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:36
		_go_fuzz_dep_.CoverTab[189294]++
												panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:37
		// _ = "end of CoverTab[189294]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:38
		_go_fuzz_dep_.CoverTab[189295]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:38
		// _ = "end of CoverTab[189295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:38
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:38
	// _ = "end of CoverTab[189291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:38
	_go_fuzz_dep_.CoverTab[189292]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:50
	if string(out) == "null" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:50
		_go_fuzz_dep_.CoverTab[189296]++
												panic("Tried to serialize a nil pointer.")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:51
		// _ = "end of CoverTab[189296]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:52
		_go_fuzz_dep_.CoverTab[189297]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:52
		// _ = "end of CoverTab[189297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:52
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:52
	// _ = "end of CoverTab[189292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:52
	_go_fuzz_dep_.CoverTab[189293]++
											return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:53
	// _ = "end of CoverTab[189293]"
}

// Strip all newlines and whitespace
func stripWhitespace(data string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:57
	_go_fuzz_dep_.CoverTab[189298]++
											buf := strings.Builder{}
											buf.Grow(len(data))
											for _, r := range data {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:60
		_go_fuzz_dep_.CoverTab[189300]++
												if !unicode.IsSpace(r) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:61
			_go_fuzz_dep_.CoverTab[189301]++
													buf.WriteRune(r)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:62
			// _ = "end of CoverTab[189301]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:63
			_go_fuzz_dep_.CoverTab[189302]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:63
			// _ = "end of CoverTab[189302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:63
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:63
		// _ = "end of CoverTab[189300]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:64
	// _ = "end of CoverTab[189298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:64
	_go_fuzz_dep_.CoverTab[189299]++
											return buf.String()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:65
	// _ = "end of CoverTab[189299]"
}

// Perform compression based on algorithm
func compress(algorithm CompressionAlgorithm, input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:69
	_go_fuzz_dep_.CoverTab[189303]++
											switch algorithm {
	case DEFLATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:71
		_go_fuzz_dep_.CoverTab[189304]++
												return deflate(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:72
		// _ = "end of CoverTab[189304]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:73
		_go_fuzz_dep_.CoverTab[189305]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:74
		// _ = "end of CoverTab[189305]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:75
	// _ = "end of CoverTab[189303]"
}

// Perform decompression based on algorithm
func decompress(algorithm CompressionAlgorithm, input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:79
	_go_fuzz_dep_.CoverTab[189306]++
											switch algorithm {
	case DEFLATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:81
		_go_fuzz_dep_.CoverTab[189307]++
												return inflate(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:82
		// _ = "end of CoverTab[189307]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:83
		_go_fuzz_dep_.CoverTab[189308]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:84
		// _ = "end of CoverTab[189308]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:85
	// _ = "end of CoverTab[189306]"
}

// Compress with DEFLATE
func deflate(input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:89
	_go_fuzz_dep_.CoverTab[189309]++
											output := new(bytes.Buffer)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:93
	writer, _ := flate.NewWriter(output, 1)
											_, _ = io.Copy(writer, bytes.NewBuffer(input))

											err := writer.Close()
											return output.Bytes(), err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:97
	// _ = "end of CoverTab[189309]"
}

// Decompress with DEFLATE
func inflate(input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:101
	_go_fuzz_dep_.CoverTab[189310]++
											output := new(bytes.Buffer)
											reader := flate.NewReader(bytes.NewBuffer(input))

											_, err := io.Copy(output, reader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:106
		_go_fuzz_dep_.CoverTab[189312]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:107
		// _ = "end of CoverTab[189312]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:108
		_go_fuzz_dep_.CoverTab[189313]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:108
		// _ = "end of CoverTab[189313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:108
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:108
	// _ = "end of CoverTab[189310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:108
	_go_fuzz_dep_.CoverTab[189311]++

											err = reader.Close()
											return output.Bytes(), err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:111
	// _ = "end of CoverTab[189311]"
}

// byteBuffer represents a slice of bytes that can be serialized to url-safe base64.
type byteBuffer struct {
	data []byte
}

func newBuffer(data []byte) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:119
	_go_fuzz_dep_.CoverTab[189314]++
											if data == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:120
		_go_fuzz_dep_.CoverTab[189316]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:121
		// _ = "end of CoverTab[189316]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:122
		_go_fuzz_dep_.CoverTab[189317]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:122
		// _ = "end of CoverTab[189317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:122
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:122
	// _ = "end of CoverTab[189314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:122
	_go_fuzz_dep_.CoverTab[189315]++
											return &byteBuffer{
		data: data,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:125
	// _ = "end of CoverTab[189315]"
}

func newFixedSizeBuffer(data []byte, length int) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:128
	_go_fuzz_dep_.CoverTab[189318]++
											if len(data) > length {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:129
		_go_fuzz_dep_.CoverTab[189320]++
												panic("square/go-jose: invalid call to newFixedSizeBuffer (len(data) > length)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:130
		// _ = "end of CoverTab[189320]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:131
		_go_fuzz_dep_.CoverTab[189321]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:131
		// _ = "end of CoverTab[189321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:131
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:131
	// _ = "end of CoverTab[189318]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:131
	_go_fuzz_dep_.CoverTab[189319]++
											pad := make([]byte, length-len(data))
											return newBuffer(append(pad, data...))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:133
	// _ = "end of CoverTab[189319]"
}

func newBufferFromInt(num uint64) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:136
	_go_fuzz_dep_.CoverTab[189322]++
											data := make([]byte, 8)
											binary.BigEndian.PutUint64(data, num)
											return newBuffer(bytes.TrimLeft(data, "\x00"))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:139
	// _ = "end of CoverTab[189322]"
}

func (b *byteBuffer) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:142
	_go_fuzz_dep_.CoverTab[189323]++
											return json.Marshal(b.base64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:143
	// _ = "end of CoverTab[189323]"
}

func (b *byteBuffer) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:146
	_go_fuzz_dep_.CoverTab[189324]++
											var encoded string
											err := json.Unmarshal(data, &encoded)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:149
		_go_fuzz_dep_.CoverTab[189328]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:150
		// _ = "end of CoverTab[189328]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:151
		_go_fuzz_dep_.CoverTab[189329]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:151
		// _ = "end of CoverTab[189329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:151
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:151
	// _ = "end of CoverTab[189324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:151
	_go_fuzz_dep_.CoverTab[189325]++

											if encoded == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:153
		_go_fuzz_dep_.CoverTab[189330]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:154
		// _ = "end of CoverTab[189330]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:155
		_go_fuzz_dep_.CoverTab[189331]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:155
		// _ = "end of CoverTab[189331]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:155
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:155
	// _ = "end of CoverTab[189325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:155
	_go_fuzz_dep_.CoverTab[189326]++

											decoded, err := base64.RawURLEncoding.DecodeString(encoded)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:158
		_go_fuzz_dep_.CoverTab[189332]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:159
		// _ = "end of CoverTab[189332]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:160
		_go_fuzz_dep_.CoverTab[189333]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:160
		// _ = "end of CoverTab[189333]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:160
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:160
	// _ = "end of CoverTab[189326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:160
	_go_fuzz_dep_.CoverTab[189327]++

											*b = *newBuffer(decoded)

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:164
	// _ = "end of CoverTab[189327]"
}

func (b *byteBuffer) base64() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:167
	_go_fuzz_dep_.CoverTab[189334]++
											return base64.RawURLEncoding.EncodeToString(b.data)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:168
	// _ = "end of CoverTab[189334]"
}

func (b *byteBuffer) bytes() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:171
	_go_fuzz_dep_.CoverTab[189335]++

											if b == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:173
		_go_fuzz_dep_.CoverTab[189337]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:174
		// _ = "end of CoverTab[189337]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:175
		_go_fuzz_dep_.CoverTab[189338]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:175
		// _ = "end of CoverTab[189338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:175
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:175
	// _ = "end of CoverTab[189335]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:175
	_go_fuzz_dep_.CoverTab[189336]++
											return b.data
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:176
	// _ = "end of CoverTab[189336]"
}

func (b byteBuffer) bigInt() *big.Int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:179
	_go_fuzz_dep_.CoverTab[189339]++
											return new(big.Int).SetBytes(b.data)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:180
	// _ = "end of CoverTab[189339]"
}

func (b byteBuffer) toInt() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:183
	_go_fuzz_dep_.CoverTab[189340]++
											return int(b.bigInt().Int64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:184
	// _ = "end of CoverTab[189340]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:185
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/encoding.go:185
var _ = _go_fuzz_dep_.CoverTab
