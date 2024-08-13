//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:1
)

type SaslHandshakeResponse struct {
	Err			KError
	EnabledMechanisms	[]string
}

func (r *SaslHandshakeResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:8
	_go_fuzz_dep_.CoverTab[106682]++
													pe.putInt16(int16(r.Err))
													return pe.putStringArray(r.EnabledMechanisms)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:10
	// _ = "end of CoverTab[106682]"
}

func (r *SaslHandshakeResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:13
	_go_fuzz_dep_.CoverTab[106683]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:15
		_go_fuzz_dep_.CoverTab[106686]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:16
		// _ = "end of CoverTab[106686]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:17
		_go_fuzz_dep_.CoverTab[106687]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:17
		// _ = "end of CoverTab[106687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:17
	// _ = "end of CoverTab[106683]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:17
	_go_fuzz_dep_.CoverTab[106684]++

													r.Err = KError(kerr)

													if r.EnabledMechanisms, err = pd.getStringArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:21
		_go_fuzz_dep_.CoverTab[106688]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:22
		// _ = "end of CoverTab[106688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:23
		_go_fuzz_dep_.CoverTab[106689]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:23
		// _ = "end of CoverTab[106689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:23
	// _ = "end of CoverTab[106684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:23
	_go_fuzz_dep_.CoverTab[106685]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:25
	// _ = "end of CoverTab[106685]"
}

func (r *SaslHandshakeResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:28
	_go_fuzz_dep_.CoverTab[106690]++
													return 17
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:29
	// _ = "end of CoverTab[106690]"
}

func (r *SaslHandshakeResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:32
	_go_fuzz_dep_.CoverTab[106691]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:33
	// _ = "end of CoverTab[106691]"
}

func (r *SaslHandshakeResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:36
	_go_fuzz_dep_.CoverTab[106692]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:37
	// _ = "end of CoverTab[106692]"
}

func (r *SaslHandshakeResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:40
	_go_fuzz_dep_.CoverTab[106693]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:41
	// _ = "end of CoverTab[106693]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_handshake_response.go:42
var _ = _go_fuzz_dep_.CoverTab
