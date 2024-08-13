//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:1
)

import (
	"fmt"

	"github.com/rcrowley/go-metrics"
)

// Encoder is the interface that wraps the basic Encode method.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:9
// Anything implementing Encoder can be turned into bytes using Kafka's encoding rules.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:11
type encoder interface {
	encode(pe packetEncoder) error
}

type encoderWithHeader interface {
	encoder
	headerVersion() int16
}

// Encode takes an Encoder and turns it into bytes while potentially recording metrics.
func encode(e encoder, metricRegistry metrics.Registry) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:21
	_go_fuzz_dep_.CoverTab[102678]++
												if e == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:22
		_go_fuzz_dep_.CoverTab[102683]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:23
		// _ = "end of CoverTab[102683]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:24
		_go_fuzz_dep_.CoverTab[102684]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:24
		// _ = "end of CoverTab[102684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:24
	// _ = "end of CoverTab[102678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:24
	_go_fuzz_dep_.CoverTab[102679]++

												var prepEnc prepEncoder
												var realEnc realEncoder

												err := e.encode(&prepEnc)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:30
		_go_fuzz_dep_.CoverTab[102685]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:31
		// _ = "end of CoverTab[102685]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:32
		_go_fuzz_dep_.CoverTab[102686]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:32
		// _ = "end of CoverTab[102686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:32
	// _ = "end of CoverTab[102679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:32
	_go_fuzz_dep_.CoverTab[102680]++

												if prepEnc.length < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:34
		_go_fuzz_dep_.CoverTab[102687]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:34
		return prepEnc.length > int(MaxRequestSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:34
		// _ = "end of CoverTab[102687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:34
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:34
		_go_fuzz_dep_.CoverTab[102688]++
													return nil, PacketEncodingError{fmt.Sprintf("invalid request size (%d)", prepEnc.length)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:35
		// _ = "end of CoverTab[102688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:36
		_go_fuzz_dep_.CoverTab[102689]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:36
		// _ = "end of CoverTab[102689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:36
	// _ = "end of CoverTab[102680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:36
	_go_fuzz_dep_.CoverTab[102681]++

												realEnc.raw = make([]byte, prepEnc.length)
												realEnc.registry = metricRegistry
												err = e.encode(&realEnc)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:41
		_go_fuzz_dep_.CoverTab[102690]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:42
		// _ = "end of CoverTab[102690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:43
		_go_fuzz_dep_.CoverTab[102691]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:43
		// _ = "end of CoverTab[102691]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:43
	// _ = "end of CoverTab[102681]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:43
	_go_fuzz_dep_.CoverTab[102682]++

												return realEnc.raw, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:45
	// _ = "end of CoverTab[102682]"
}

// decoder is the interface that wraps the basic Decode method.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:48
// Anything implementing Decoder can be extracted from bytes using Kafka's encoding rules.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:50
type decoder interface {
	decode(pd packetDecoder) error
}

type versionedDecoder interface {
	decode(pd packetDecoder, version int16) error
}

// decode takes bytes and a decoder and fills the fields of the decoder from the bytes,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:58
// interpreted using Kafka's encoding rules.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:60
func decode(buf []byte, in decoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:60
	_go_fuzz_dep_.CoverTab[102692]++
												if buf == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:61
		_go_fuzz_dep_.CoverTab[102696]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:62
		// _ = "end of CoverTab[102696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:63
		_go_fuzz_dep_.CoverTab[102697]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:63
		// _ = "end of CoverTab[102697]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:63
	// _ = "end of CoverTab[102692]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:63
	_go_fuzz_dep_.CoverTab[102693]++

												helper := realDecoder{raw: buf}
												err := in.decode(&helper)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:67
		_go_fuzz_dep_.CoverTab[102698]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:68
		// _ = "end of CoverTab[102698]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:69
		_go_fuzz_dep_.CoverTab[102699]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:69
		// _ = "end of CoverTab[102699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:69
	// _ = "end of CoverTab[102693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:69
	_go_fuzz_dep_.CoverTab[102694]++

												if helper.off != len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:71
		_go_fuzz_dep_.CoverTab[102700]++
													return PacketDecodingError{"invalid length"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:72
		// _ = "end of CoverTab[102700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:73
		_go_fuzz_dep_.CoverTab[102701]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:73
		// _ = "end of CoverTab[102701]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:73
	// _ = "end of CoverTab[102694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:73
	_go_fuzz_dep_.CoverTab[102695]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:75
	// _ = "end of CoverTab[102695]"
}

func versionedDecode(buf []byte, in versionedDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:78
	_go_fuzz_dep_.CoverTab[102702]++
												if buf == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:79
		_go_fuzz_dep_.CoverTab[102706]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:80
		// _ = "end of CoverTab[102706]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:81
		_go_fuzz_dep_.CoverTab[102707]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:81
		// _ = "end of CoverTab[102707]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:81
	// _ = "end of CoverTab[102702]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:81
	_go_fuzz_dep_.CoverTab[102703]++

												helper := realDecoder{raw: buf}
												err := in.decode(&helper, version)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:85
		_go_fuzz_dep_.CoverTab[102708]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:86
		// _ = "end of CoverTab[102708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:87
		_go_fuzz_dep_.CoverTab[102709]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:87
		// _ = "end of CoverTab[102709]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:87
	// _ = "end of CoverTab[102703]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:87
	_go_fuzz_dep_.CoverTab[102704]++

												if helper.off != len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:89
		_go_fuzz_dep_.CoverTab[102710]++
													return PacketDecodingError{
			Info: fmt.Sprintf("invalid length (off=%d, len=%d)", helper.off, len(buf)),
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:92
		// _ = "end of CoverTab[102710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:93
		_go_fuzz_dep_.CoverTab[102711]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:93
		// _ = "end of CoverTab[102711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:93
	// _ = "end of CoverTab[102704]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:93
	_go_fuzz_dep_.CoverTab[102705]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:95
	// _ = "end of CoverTab[102705]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:96
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/encoder_decoder.go:96
var _ = _go_fuzz_dep_.CoverTab
