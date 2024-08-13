//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:1
)

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

// ScramFormatter implementation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:10
// @see: https://github.com/apache/kafka/blob/99b9b3e84f4e98c3f07714e1de6a139a004cbc5b/clients/src/main/java/org/apache/kafka/common/security/scram/internals/ScramFormatter.java#L93
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:12
type scramFormatter struct {
	mechanism ScramMechanismType
}

func (s scramFormatter) mac(key []byte) (hash.Hash, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:16
	_go_fuzz_dep_.CoverTab[106694]++
												var m hash.Hash

												switch s.mechanism {
	case SCRAM_MECHANISM_SHA_256:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:20
		_go_fuzz_dep_.CoverTab[106696]++
													m = hmac.New(sha256.New, key)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:21
		// _ = "end of CoverTab[106696]"

	case SCRAM_MECHANISM_SHA_512:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:23
		_go_fuzz_dep_.CoverTab[106697]++
													m = hmac.New(sha512.New, key)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:24
		// _ = "end of CoverTab[106697]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:25
		_go_fuzz_dep_.CoverTab[106698]++
													return nil, ErrUnknownScramMechanism
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:26
		// _ = "end of CoverTab[106698]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:27
	// _ = "end of CoverTab[106694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:27
	_go_fuzz_dep_.CoverTab[106695]++

												return m, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:29
	// _ = "end of CoverTab[106695]"
}

func (s scramFormatter) hmac(key []byte, extra []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:32
	_go_fuzz_dep_.CoverTab[106699]++
												mac, err := s.mac(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:34
		_go_fuzz_dep_.CoverTab[106702]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:35
		// _ = "end of CoverTab[106702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:36
		_go_fuzz_dep_.CoverTab[106703]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:36
		// _ = "end of CoverTab[106703]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:36
	// _ = "end of CoverTab[106699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:36
	_go_fuzz_dep_.CoverTab[106700]++

												if _, err := mac.Write(extra); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:38
		_go_fuzz_dep_.CoverTab[106704]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:39
		// _ = "end of CoverTab[106704]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:40
		_go_fuzz_dep_.CoverTab[106705]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:40
		// _ = "end of CoverTab[106705]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:40
	// _ = "end of CoverTab[106700]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:40
	_go_fuzz_dep_.CoverTab[106701]++
												return mac.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:41
	// _ = "end of CoverTab[106701]"
}

func (s scramFormatter) xor(result []byte, second []byte) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:44
	_go_fuzz_dep_.CoverTab[106706]++
												for i := 0; i < len(result); i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:45
		_go_fuzz_dep_.CoverTab[106707]++
													result[i] = result[i] ^ second[i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:46
		// _ = "end of CoverTab[106707]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:47
	// _ = "end of CoverTab[106706]"
}

func (s scramFormatter) saltedPassword(password []byte, salt []byte, iterations int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:50
	_go_fuzz_dep_.CoverTab[106708]++
												mac, err := s.mac(password)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:52
		_go_fuzz_dep_.CoverTab[106713]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:53
		// _ = "end of CoverTab[106713]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:54
		_go_fuzz_dep_.CoverTab[106714]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:54
		// _ = "end of CoverTab[106714]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:54
	// _ = "end of CoverTab[106708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:54
	_go_fuzz_dep_.CoverTab[106709]++

												if _, err := mac.Write(salt); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:56
		_go_fuzz_dep_.CoverTab[106715]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:57
		// _ = "end of CoverTab[106715]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:58
		_go_fuzz_dep_.CoverTab[106716]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:58
		// _ = "end of CoverTab[106716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:58
	// _ = "end of CoverTab[106709]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:58
	_go_fuzz_dep_.CoverTab[106710]++
												if _, err := mac.Write([]byte{0, 0, 0, 1}); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:59
		_go_fuzz_dep_.CoverTab[106717]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:60
		// _ = "end of CoverTab[106717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:61
		_go_fuzz_dep_.CoverTab[106718]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:61
		// _ = "end of CoverTab[106718]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:61
	// _ = "end of CoverTab[106710]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:61
	_go_fuzz_dep_.CoverTab[106711]++

												u1 := mac.Sum(nil)
												prev := u1
												result := u1

												for i := 2; i <= iterations; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:67
		_go_fuzz_dep_.CoverTab[106719]++
													ui, err := s.hmac(password, prev)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:69
			_go_fuzz_dep_.CoverTab[106721]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:70
			// _ = "end of CoverTab[106721]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:71
			_go_fuzz_dep_.CoverTab[106722]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:71
			// _ = "end of CoverTab[106722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:71
		// _ = "end of CoverTab[106719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:71
		_go_fuzz_dep_.CoverTab[106720]++

													s.xor(result, ui)
													prev = ui
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:74
		// _ = "end of CoverTab[106720]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:75
	// _ = "end of CoverTab[106711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:75
	_go_fuzz_dep_.CoverTab[106712]++

												return result, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:77
	// _ = "end of CoverTab[106712]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:78
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/scram_formatter.go:78
var _ = _go_fuzz_dep_.CoverTab
