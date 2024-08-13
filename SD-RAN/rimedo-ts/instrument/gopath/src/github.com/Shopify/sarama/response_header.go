//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:1
)

import "fmt"

const (
	responseLengthSize	= 4
	correlationIDSize	= 4
)

type responseHeader struct {
	length		int32
	correlationID	int32
}

func (r *responseHeader) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:15
	_go_fuzz_dep_.CoverTab[106633]++
												r.length, err = pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:17
		_go_fuzz_dep_.CoverTab[106637]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:18
		// _ = "end of CoverTab[106637]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:19
		_go_fuzz_dep_.CoverTab[106638]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:19
		// _ = "end of CoverTab[106638]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:19
	// _ = "end of CoverTab[106633]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:19
	_go_fuzz_dep_.CoverTab[106634]++
												if r.length <= 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:20
		_go_fuzz_dep_.CoverTab[106639]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:20
		return r.length > MaxResponseSize
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:20
		// _ = "end of CoverTab[106639]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:20
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:20
		_go_fuzz_dep_.CoverTab[106640]++
													return PacketDecodingError{fmt.Sprintf("message of length %d too large or too small", r.length)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:21
		// _ = "end of CoverTab[106640]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:22
		_go_fuzz_dep_.CoverTab[106641]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:22
		// _ = "end of CoverTab[106641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:22
	// _ = "end of CoverTab[106634]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:22
	_go_fuzz_dep_.CoverTab[106635]++

												r.correlationID, err = pd.getInt32()

												if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:26
		_go_fuzz_dep_.CoverTab[106642]++
													if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:27
			_go_fuzz_dep_.CoverTab[106643]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:28
			// _ = "end of CoverTab[106643]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:29
			_go_fuzz_dep_.CoverTab[106644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:29
			// _ = "end of CoverTab[106644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:29
		// _ = "end of CoverTab[106642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:30
		_go_fuzz_dep_.CoverTab[106645]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:30
		// _ = "end of CoverTab[106645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:30
	// _ = "end of CoverTab[106635]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:30
	_go_fuzz_dep_.CoverTab[106636]++

												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:32
	// _ = "end of CoverTab[106636]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:33
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/response_header.go:33
var _ = _go_fuzz_dep_.CoverTab
