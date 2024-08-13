//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:1
)

type SaslAuthenticateRequest struct {
	SaslAuthBytes []byte
}

// APIKeySASLAuth is the API key for the SaslAuthenticate Kafka API
const APIKeySASLAuth = 36

func (r *SaslAuthenticateRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:10
	_go_fuzz_dep_.CoverTab[106649]++
													return pe.putBytes(r.SaslAuthBytes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:11
	// _ = "end of CoverTab[106649]"
}

func (r *SaslAuthenticateRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:14
	_go_fuzz_dep_.CoverTab[106650]++
													r.SaslAuthBytes, err = pd.getBytes()
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:16
	// _ = "end of CoverTab[106650]"
}

func (r *SaslAuthenticateRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:19
	_go_fuzz_dep_.CoverTab[106651]++
													return APIKeySASLAuth
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:20
	// _ = "end of CoverTab[106651]"
}

func (r *SaslAuthenticateRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:23
	_go_fuzz_dep_.CoverTab[106652]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:24
	// _ = "end of CoverTab[106652]"
}

func (r *SaslAuthenticateRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:27
	_go_fuzz_dep_.CoverTab[106653]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:28
	// _ = "end of CoverTab[106653]"
}

func (r *SaslAuthenticateRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:31
	_go_fuzz_dep_.CoverTab[106654]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:32
	// _ = "end of CoverTab[106654]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:33
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_request.go:33
var _ = _go_fuzz_dep_.CoverTab
