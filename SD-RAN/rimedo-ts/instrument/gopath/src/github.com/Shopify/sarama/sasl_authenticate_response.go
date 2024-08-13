//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:1
)

type SaslAuthenticateResponse struct {
	Err		KError
	ErrorMessage	*string
	SaslAuthBytes	[]byte
}

func (r *SaslAuthenticateResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:9
	_go_fuzz_dep_.CoverTab[106655]++
													pe.putInt16(int16(r.Err))
													if err := pe.putNullableString(r.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:11
		_go_fuzz_dep_.CoverTab[106657]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:12
		// _ = "end of CoverTab[106657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:13
		_go_fuzz_dep_.CoverTab[106658]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:13
		// _ = "end of CoverTab[106658]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:13
	// _ = "end of CoverTab[106655]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:13
	_go_fuzz_dep_.CoverTab[106656]++
													return pe.putBytes(r.SaslAuthBytes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:14
	// _ = "end of CoverTab[106656]"
}

func (r *SaslAuthenticateResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:17
	_go_fuzz_dep_.CoverTab[106659]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:19
		_go_fuzz_dep_.CoverTab[106662]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:20
		// _ = "end of CoverTab[106662]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:21
		_go_fuzz_dep_.CoverTab[106663]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:21
		// _ = "end of CoverTab[106663]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:21
	// _ = "end of CoverTab[106659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:21
	_go_fuzz_dep_.CoverTab[106660]++

													r.Err = KError(kerr)

													if r.ErrorMessage, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:25
		_go_fuzz_dep_.CoverTab[106664]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:26
		// _ = "end of CoverTab[106664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:27
		_go_fuzz_dep_.CoverTab[106665]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:27
		// _ = "end of CoverTab[106665]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:27
	// _ = "end of CoverTab[106660]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:27
	_go_fuzz_dep_.CoverTab[106661]++

													r.SaslAuthBytes, err = pd.getBytes()

													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:31
	// _ = "end of CoverTab[106661]"
}

func (r *SaslAuthenticateResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:34
	_go_fuzz_dep_.CoverTab[106666]++
													return APIKeySASLAuth
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:35
	// _ = "end of CoverTab[106666]"
}

func (r *SaslAuthenticateResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:38
	_go_fuzz_dep_.CoverTab[106667]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:39
	// _ = "end of CoverTab[106667]"
}

func (r *SaslAuthenticateResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:42
	_go_fuzz_dep_.CoverTab[106668]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:43
	// _ = "end of CoverTab[106668]"
}

func (r *SaslAuthenticateResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:46
	_go_fuzz_dep_.CoverTab[106669]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:47
	// _ = "end of CoverTab[106669]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sasl_authenticate_response.go:48
var _ = _go_fuzz_dep_.CoverTab
