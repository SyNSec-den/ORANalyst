//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:17
)

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/binary"
	"io"
	"math/big"
	"regexp"
	"strings"

	"gopkg.in/square/go-jose.v1/json"
)

var stripWhitespaceRegex = regexp.MustCompile("\\s")

// Url-safe base64 encode that strips padding
func base64URLEncode(data []byte) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:35
	_go_fuzz_dep_.CoverTab[186138]++
											var result = base64.URLEncoding.EncodeToString(data)
											return strings.TrimRight(result, "=")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:37
	// _ = "end of CoverTab[186138]"
}

// Url-safe base64 decoder that adds padding
func base64URLDecode(data string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:41
	_go_fuzz_dep_.CoverTab[186139]++
											var missing = (4 - len(data)%4) % 4
											data += strings.Repeat("=", missing)
											return base64.URLEncoding.DecodeString(data)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:44
	// _ = "end of CoverTab[186139]"
}

// Helper function to serialize known-good objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:47
// Precondition: value is not a nil pointer.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:49
func mustSerializeJSON(value interface{}) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:49
	_go_fuzz_dep_.CoverTab[186140]++
											out, err := json.Marshal(value)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:51
		_go_fuzz_dep_.CoverTab[186143]++
												panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:52
		// _ = "end of CoverTab[186143]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:53
		_go_fuzz_dep_.CoverTab[186144]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:53
		// _ = "end of CoverTab[186144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:53
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:53
	// _ = "end of CoverTab[186140]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:53
	_go_fuzz_dep_.CoverTab[186141]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:65
	if string(out) == "null" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:65
		_go_fuzz_dep_.CoverTab[186145]++
												panic("Tried to serialize a nil pointer.")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:66
		// _ = "end of CoverTab[186145]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:67
		_go_fuzz_dep_.CoverTab[186146]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:67
		// _ = "end of CoverTab[186146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:67
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:67
	// _ = "end of CoverTab[186141]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:67
	_go_fuzz_dep_.CoverTab[186142]++
											return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:68
	// _ = "end of CoverTab[186142]"
}

// Strip all newlines and whitespace
func stripWhitespace(data string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:72
	_go_fuzz_dep_.CoverTab[186147]++
											return stripWhitespaceRegex.ReplaceAllString(data, "")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:73
	// _ = "end of CoverTab[186147]"
}

// Perform compression based on algorithm
func compress(algorithm CompressionAlgorithm, input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:77
	_go_fuzz_dep_.CoverTab[186148]++
											switch algorithm {
	case DEFLATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:79
		_go_fuzz_dep_.CoverTab[186149]++
												return deflate(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:80
		// _ = "end of CoverTab[186149]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:81
		_go_fuzz_dep_.CoverTab[186150]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:82
		// _ = "end of CoverTab[186150]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:83
	// _ = "end of CoverTab[186148]"
}

// Perform decompression based on algorithm
func decompress(algorithm CompressionAlgorithm, input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:87
	_go_fuzz_dep_.CoverTab[186151]++
											switch algorithm {
	case DEFLATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:89
		_go_fuzz_dep_.CoverTab[186152]++
												return inflate(input)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:90
		// _ = "end of CoverTab[186152]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:91
		_go_fuzz_dep_.CoverTab[186153]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:92
		// _ = "end of CoverTab[186153]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:93
	// _ = "end of CoverTab[186151]"
}

// Compress with DEFLATE
func deflate(input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:97
	_go_fuzz_dep_.CoverTab[186154]++
											output := new(bytes.Buffer)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:101
	writer, _ := flate.NewWriter(output, 1)
											_, _ = io.Copy(writer, bytes.NewBuffer(input))

											err := writer.Close()
											return output.Bytes(), err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:105
	// _ = "end of CoverTab[186154]"
}

// Decompress with DEFLATE
func inflate(input []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:109
	_go_fuzz_dep_.CoverTab[186155]++
											output := new(bytes.Buffer)
											reader := flate.NewReader(bytes.NewBuffer(input))

											_, err := io.Copy(output, reader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:114
		_go_fuzz_dep_.CoverTab[186157]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:115
		// _ = "end of CoverTab[186157]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:116
		_go_fuzz_dep_.CoverTab[186158]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:116
		// _ = "end of CoverTab[186158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:116
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:116
	// _ = "end of CoverTab[186155]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:116
	_go_fuzz_dep_.CoverTab[186156]++

											err = reader.Close()
											return output.Bytes(), err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:119
	// _ = "end of CoverTab[186156]"
}

// byteBuffer represents a slice of bytes that can be serialized to url-safe base64.
type byteBuffer struct {
	data []byte
}

func newBuffer(data []byte) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:127
	_go_fuzz_dep_.CoverTab[186159]++
											if data == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:128
		_go_fuzz_dep_.CoverTab[186161]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:129
		// _ = "end of CoverTab[186161]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:130
		_go_fuzz_dep_.CoverTab[186162]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:130
		// _ = "end of CoverTab[186162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:130
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:130
	// _ = "end of CoverTab[186159]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:130
	_go_fuzz_dep_.CoverTab[186160]++
											return &byteBuffer{
		data: data,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:133
	// _ = "end of CoverTab[186160]"
}

func newFixedSizeBuffer(data []byte, length int) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:136
	_go_fuzz_dep_.CoverTab[186163]++
											if len(data) > length {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:137
		_go_fuzz_dep_.CoverTab[186165]++
												panic("square/go-jose: invalid call to newFixedSizeBuffer (len(data) > length)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:138
		// _ = "end of CoverTab[186165]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:139
		_go_fuzz_dep_.CoverTab[186166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:139
		// _ = "end of CoverTab[186166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:139
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:139
	// _ = "end of CoverTab[186163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:139
	_go_fuzz_dep_.CoverTab[186164]++
											pad := make([]byte, length-len(data))
											return newBuffer(append(pad, data...))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:141
	// _ = "end of CoverTab[186164]"
}

func newBufferFromInt(num uint64) *byteBuffer {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:144
	_go_fuzz_dep_.CoverTab[186167]++
											data := make([]byte, 8)
											binary.BigEndian.PutUint64(data, num)
											return newBuffer(bytes.TrimLeft(data, "\x00"))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:147
	// _ = "end of CoverTab[186167]"
}

func (b *byteBuffer) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:150
	_go_fuzz_dep_.CoverTab[186168]++
											return json.Marshal(b.base64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:151
	// _ = "end of CoverTab[186168]"
}

func (b *byteBuffer) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:154
	_go_fuzz_dep_.CoverTab[186169]++
											var encoded string
											err := json.Unmarshal(data, &encoded)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:157
		_go_fuzz_dep_.CoverTab[186173]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:158
		// _ = "end of CoverTab[186173]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:159
		_go_fuzz_dep_.CoverTab[186174]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:159
		// _ = "end of CoverTab[186174]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:159
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:159
	// _ = "end of CoverTab[186169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:159
	_go_fuzz_dep_.CoverTab[186170]++

											if encoded == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:161
		_go_fuzz_dep_.CoverTab[186175]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:162
		// _ = "end of CoverTab[186175]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:163
		_go_fuzz_dep_.CoverTab[186176]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:163
		// _ = "end of CoverTab[186176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:163
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:163
	// _ = "end of CoverTab[186170]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:163
	_go_fuzz_dep_.CoverTab[186171]++

											decoded, err := base64URLDecode(encoded)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:166
		_go_fuzz_dep_.CoverTab[186177]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:167
		// _ = "end of CoverTab[186177]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:168
		_go_fuzz_dep_.CoverTab[186178]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:168
		// _ = "end of CoverTab[186178]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:168
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:168
	// _ = "end of CoverTab[186171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:168
	_go_fuzz_dep_.CoverTab[186172]++

											*b = *newBuffer(decoded)

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:172
	// _ = "end of CoverTab[186172]"
}

func (b *byteBuffer) base64() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:175
	_go_fuzz_dep_.CoverTab[186179]++
											return base64URLEncode(b.data)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:176
	// _ = "end of CoverTab[186179]"
}

func (b *byteBuffer) bytes() []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:179
	_go_fuzz_dep_.CoverTab[186180]++

											if b == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:181
		_go_fuzz_dep_.CoverTab[186182]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:182
		// _ = "end of CoverTab[186182]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:183
		_go_fuzz_dep_.CoverTab[186183]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:183
		// _ = "end of CoverTab[186183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:183
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:183
	// _ = "end of CoverTab[186180]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:183
	_go_fuzz_dep_.CoverTab[186181]++
											return b.data
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:184
	// _ = "end of CoverTab[186181]"
}

func (b byteBuffer) bigInt() *big.Int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:187
	_go_fuzz_dep_.CoverTab[186184]++
											return new(big.Int).SetBytes(b.data)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:188
	// _ = "end of CoverTab[186184]"
}

func (b byteBuffer) toInt() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:191
	_go_fuzz_dep_.CoverTab[186185]++
											return int(b.bigInt().Int64())
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:192
	// _ = "end of CoverTab[186185]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:193
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/encoding.go:193
var _ = _go_fuzz_dep_.CoverTab
