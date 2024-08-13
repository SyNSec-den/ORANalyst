//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:1
)

type SaslHandshakeRequest struct {
	Mechanism	string
	Version		int16
}

func (r *SaslHandshakeRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:8
	_go_fuzz_dep_.CoverTab[106670]++
													if err := pe.putString(r.Mechanism); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:9
		_go_fuzz_dep_.CoverTab[106672]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:10
		// _ = "end of CoverTab[106672]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:11
		_go_fuzz_dep_.CoverTab[106673]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:11
		// _ = "end of CoverTab[106673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:11
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:11
	// _ = "end of CoverTab[106670]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:11
	_go_fuzz_dep_.CoverTab[106671]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:13
	// _ = "end of CoverTab[106671]"
}

func (r *SaslHandshakeRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:16
	_go_fuzz_dep_.CoverTab[106674]++
													if r.Mechanism, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:17
		_go_fuzz_dep_.CoverTab[106676]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:18
		// _ = "end of CoverTab[106676]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:19
		_go_fuzz_dep_.CoverTab[106677]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:19
		// _ = "end of CoverTab[106677]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:19
	// _ = "end of CoverTab[106674]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:19
	_go_fuzz_dep_.CoverTab[106675]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:21
	// _ = "end of CoverTab[106675]"
}

func (r *SaslHandshakeRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:24
	_go_fuzz_dep_.CoverTab[106678]++
													return 17
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:25
	// _ = "end of CoverTab[106678]"
}

func (r *SaslHandshakeRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:28
	_go_fuzz_dep_.CoverTab[106679]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:29
	// _ = "end of CoverTab[106679]"
}

func (r *SaslHandshakeRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:32
	_go_fuzz_dep_.CoverTab[106680]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:33
	// _ = "end of CoverTab[106680]"
}

func (r *SaslHandshakeRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:36
	_go_fuzz_dep_.CoverTab[106681]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:37
	// _ = "end of CoverTab[106681]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_request.go:38
var _ = _go_fuzz_dep_.CoverTab
