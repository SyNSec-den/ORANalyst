//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:1
)

import (
	"fmt"
	"time"
)

type Timestamp struct {
	*time.Time
}

func (t Timestamp) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:12
	_go_fuzz_dep_.CoverTab[106894]++
											timestamp := int64(-1)

											if !t.Before(time.Unix(0, 0)) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:15
		_go_fuzz_dep_.CoverTab[106896]++
												timestamp = t.UnixNano() / int64(time.Millisecond)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:16
		// _ = "end of CoverTab[106896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:17
		_go_fuzz_dep_.CoverTab[106897]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:17
		if !t.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:17
			_go_fuzz_dep_.CoverTab[106898]++
													return PacketEncodingError{fmt.Sprintf("invalid timestamp (%v)", t)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:18
			// _ = "end of CoverTab[106898]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
			_go_fuzz_dep_.CoverTab[106899]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
			// _ = "end of CoverTab[106899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
		// _ = "end of CoverTab[106897]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
	// _ = "end of CoverTab[106894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:19
	_go_fuzz_dep_.CoverTab[106895]++

											pe.putInt64(timestamp)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:22
	// _ = "end of CoverTab[106895]"
}

func (t Timestamp) decode(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:25
	_go_fuzz_dep_.CoverTab[106900]++
											millis, err := pd.getInt64()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:27
		_go_fuzz_dep_.CoverTab[106903]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:28
		// _ = "end of CoverTab[106903]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:29
		_go_fuzz_dep_.CoverTab[106904]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:29
		// _ = "end of CoverTab[106904]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:29
	// _ = "end of CoverTab[106900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:29
	_go_fuzz_dep_.CoverTab[106901]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:33
	timestamp := time.Time{}
	if millis >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:34
		_go_fuzz_dep_.CoverTab[106905]++
												timestamp = time.Unix(millis/1000, (millis%1000)*int64(time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:35
		// _ = "end of CoverTab[106905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:36
		_go_fuzz_dep_.CoverTab[106906]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:36
		// _ = "end of CoverTab[106906]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:36
	// _ = "end of CoverTab[106901]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:36
	_go_fuzz_dep_.CoverTab[106902]++

											*t.Time = timestamp
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:39
	// _ = "end of CoverTab[106902]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/timestamp.go:40
var _ = _go_fuzz_dep_.CoverTab
