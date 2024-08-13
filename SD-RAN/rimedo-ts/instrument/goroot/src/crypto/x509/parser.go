// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/parser.go:5
package x509

//line /usr/local/go/src/crypto/x509/parser.go:5
import (
//line /usr/local/go/src/crypto/x509/parser.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/parser.go:5
)
//line /usr/local/go/src/crypto/x509/parser.go:5
import (
//line /usr/local/go/src/crypto/x509/parser.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/parser.go:5
)

import (
	"bytes"
	"crypto/dsa"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/crypto/cryptobyte"
	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

// isPrintable reports whether the given b is in the ASN.1 PrintableString set.
//line /usr/local/go/src/crypto/x509/parser.go:32
// This is a simplified version of encoding/asn1.isPrintable.
//line /usr/local/go/src/crypto/x509/parser.go:34
func isPrintable(b byte) bool {
//line /usr/local/go/src/crypto/x509/parser.go:34
	_go_fuzz_dep_.CoverTab[18397]++
							return 'a' <= b && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:35
		_go_fuzz_dep_.CoverTab[18398]++
//line /usr/local/go/src/crypto/x509/parser.go:35
		return b <= 'z'
//line /usr/local/go/src/crypto/x509/parser.go:35
		// _ = "end of CoverTab[18398]"
//line /usr/local/go/src/crypto/x509/parser.go:35
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:35
		_go_fuzz_dep_.CoverTab[18399]++
//line /usr/local/go/src/crypto/x509/parser.go:35
		return 'A' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[18400]++
//line /usr/local/go/src/crypto/x509/parser.go:36
			return b <= 'Z'
//line /usr/local/go/src/crypto/x509/parser.go:36
			// _ = "end of CoverTab[18400]"
//line /usr/local/go/src/crypto/x509/parser.go:36
		}()
//line /usr/local/go/src/crypto/x509/parser.go:36
		// _ = "end of CoverTab[18399]"
//line /usr/local/go/src/crypto/x509/parser.go:36
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:36
		_go_fuzz_dep_.CoverTab[18401]++
//line /usr/local/go/src/crypto/x509/parser.go:36
		return '0' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[18402]++
//line /usr/local/go/src/crypto/x509/parser.go:37
			return b <= '9'
//line /usr/local/go/src/crypto/x509/parser.go:37
			// _ = "end of CoverTab[18402]"
//line /usr/local/go/src/crypto/x509/parser.go:37
		}()
//line /usr/local/go/src/crypto/x509/parser.go:37
		// _ = "end of CoverTab[18401]"
//line /usr/local/go/src/crypto/x509/parser.go:37
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:37
		_go_fuzz_dep_.CoverTab[18403]++
//line /usr/local/go/src/crypto/x509/parser.go:37
		return '\'' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[18404]++
//line /usr/local/go/src/crypto/x509/parser.go:38
			return b <= ')'
//line /usr/local/go/src/crypto/x509/parser.go:38
			// _ = "end of CoverTab[18404]"
//line /usr/local/go/src/crypto/x509/parser.go:38
		}()
//line /usr/local/go/src/crypto/x509/parser.go:38
		// _ = "end of CoverTab[18403]"
//line /usr/local/go/src/crypto/x509/parser.go:38
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:38
		_go_fuzz_dep_.CoverTab[18405]++
//line /usr/local/go/src/crypto/x509/parser.go:38
		return '+' <= b && func() bool {
									_go_fuzz_dep_.CoverTab[18406]++
//line /usr/local/go/src/crypto/x509/parser.go:39
			return b <= '/'
//line /usr/local/go/src/crypto/x509/parser.go:39
			// _ = "end of CoverTab[18406]"
//line /usr/local/go/src/crypto/x509/parser.go:39
		}()
//line /usr/local/go/src/crypto/x509/parser.go:39
		// _ = "end of CoverTab[18405]"
//line /usr/local/go/src/crypto/x509/parser.go:39
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:39
		_go_fuzz_dep_.CoverTab[18407]++
//line /usr/local/go/src/crypto/x509/parser.go:39
		return b == ' '
								// _ = "end of CoverTab[18407]"
//line /usr/local/go/src/crypto/x509/parser.go:40
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:40
		_go_fuzz_dep_.CoverTab[18408]++
//line /usr/local/go/src/crypto/x509/parser.go:40
		return b == ':'
								// _ = "end of CoverTab[18408]"
//line /usr/local/go/src/crypto/x509/parser.go:41
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:41
		_go_fuzz_dep_.CoverTab[18409]++
//line /usr/local/go/src/crypto/x509/parser.go:41
		return b == '='
								// _ = "end of CoverTab[18409]"
//line /usr/local/go/src/crypto/x509/parser.go:42
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:42
		_go_fuzz_dep_.CoverTab[18410]++
//line /usr/local/go/src/crypto/x509/parser.go:42
		return b == '?'
								// _ = "end of CoverTab[18410]"
//line /usr/local/go/src/crypto/x509/parser.go:43
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:43
		_go_fuzz_dep_.CoverTab[18411]++
//line /usr/local/go/src/crypto/x509/parser.go:43
		return b == '*'
//line /usr/local/go/src/crypto/x509/parser.go:47
		// _ = "end of CoverTab[18411]"
//line /usr/local/go/src/crypto/x509/parser.go:47
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:47
		_go_fuzz_dep_.CoverTab[18412]++
//line /usr/local/go/src/crypto/x509/parser.go:47
		return b == '&'
//line /usr/local/go/src/crypto/x509/parser.go:52
		// _ = "end of CoverTab[18412]"
//line /usr/local/go/src/crypto/x509/parser.go:52
	}()
//line /usr/local/go/src/crypto/x509/parser.go:52
	// _ = "end of CoverTab[18397]"
}

// parseASN1String parses the ASN.1 string types T61String, PrintableString,
//line /usr/local/go/src/crypto/x509/parser.go:55
// UTF8String, BMPString, IA5String, and NumericString. This is mostly copied
//line /usr/local/go/src/crypto/x509/parser.go:55
// from the respective encoding/asn1.parse... methods, rather than just
//line /usr/local/go/src/crypto/x509/parser.go:55
// increasing the API surface of that package.
//line /usr/local/go/src/crypto/x509/parser.go:59
func parseASN1String(tag cryptobyte_asn1.Tag, value []byte) (string, error) {
//line /usr/local/go/src/crypto/x509/parser.go:59
	_go_fuzz_dep_.CoverTab[18413]++
							switch tag {
	case cryptobyte_asn1.T61String:
//line /usr/local/go/src/crypto/x509/parser.go:61
		_go_fuzz_dep_.CoverTab[18415]++
								return string(value), nil
//line /usr/local/go/src/crypto/x509/parser.go:62
		// _ = "end of CoverTab[18415]"
	case cryptobyte_asn1.PrintableString:
//line /usr/local/go/src/crypto/x509/parser.go:63
		_go_fuzz_dep_.CoverTab[18416]++
								for _, b := range value {
//line /usr/local/go/src/crypto/x509/parser.go:64
			_go_fuzz_dep_.CoverTab[18429]++
									if !isPrintable(b) {
//line /usr/local/go/src/crypto/x509/parser.go:65
				_go_fuzz_dep_.CoverTab[18430]++
										return "", errors.New("invalid PrintableString")
//line /usr/local/go/src/crypto/x509/parser.go:66
				// _ = "end of CoverTab[18430]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:67
				_go_fuzz_dep_.CoverTab[18431]++
//line /usr/local/go/src/crypto/x509/parser.go:67
				// _ = "end of CoverTab[18431]"
//line /usr/local/go/src/crypto/x509/parser.go:67
			}
//line /usr/local/go/src/crypto/x509/parser.go:67
			// _ = "end of CoverTab[18429]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:68
		// _ = "end of CoverTab[18416]"
//line /usr/local/go/src/crypto/x509/parser.go:68
		_go_fuzz_dep_.CoverTab[18417]++
								return string(value), nil
//line /usr/local/go/src/crypto/x509/parser.go:69
		// _ = "end of CoverTab[18417]"
	case cryptobyte_asn1.UTF8String:
//line /usr/local/go/src/crypto/x509/parser.go:70
		_go_fuzz_dep_.CoverTab[18418]++
								if !utf8.Valid(value) {
//line /usr/local/go/src/crypto/x509/parser.go:71
			_go_fuzz_dep_.CoverTab[18432]++
									return "", errors.New("invalid UTF-8 string")
//line /usr/local/go/src/crypto/x509/parser.go:72
			// _ = "end of CoverTab[18432]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:73
			_go_fuzz_dep_.CoverTab[18433]++
//line /usr/local/go/src/crypto/x509/parser.go:73
			// _ = "end of CoverTab[18433]"
//line /usr/local/go/src/crypto/x509/parser.go:73
		}
//line /usr/local/go/src/crypto/x509/parser.go:73
		// _ = "end of CoverTab[18418]"
//line /usr/local/go/src/crypto/x509/parser.go:73
		_go_fuzz_dep_.CoverTab[18419]++
								return string(value), nil
//line /usr/local/go/src/crypto/x509/parser.go:74
		// _ = "end of CoverTab[18419]"
	case cryptobyte_asn1.Tag(asn1.TagBMPString):
//line /usr/local/go/src/crypto/x509/parser.go:75
		_go_fuzz_dep_.CoverTab[18420]++
								if len(value)%2 != 0 {
//line /usr/local/go/src/crypto/x509/parser.go:76
			_go_fuzz_dep_.CoverTab[18434]++
									return "", errors.New("invalid BMPString")
//line /usr/local/go/src/crypto/x509/parser.go:77
			// _ = "end of CoverTab[18434]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:78
			_go_fuzz_dep_.CoverTab[18435]++
//line /usr/local/go/src/crypto/x509/parser.go:78
			// _ = "end of CoverTab[18435]"
//line /usr/local/go/src/crypto/x509/parser.go:78
		}
//line /usr/local/go/src/crypto/x509/parser.go:78
		// _ = "end of CoverTab[18420]"
//line /usr/local/go/src/crypto/x509/parser.go:78
		_go_fuzz_dep_.CoverTab[18421]++

//line /usr/local/go/src/crypto/x509/parser.go:81
		if l := len(value); l >= 2 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:81
			_go_fuzz_dep_.CoverTab[18436]++
//line /usr/local/go/src/crypto/x509/parser.go:81
			return value[l-1] == 0
//line /usr/local/go/src/crypto/x509/parser.go:81
			// _ = "end of CoverTab[18436]"
//line /usr/local/go/src/crypto/x509/parser.go:81
		}() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:81
			_go_fuzz_dep_.CoverTab[18437]++
//line /usr/local/go/src/crypto/x509/parser.go:81
			return value[l-2] == 0
//line /usr/local/go/src/crypto/x509/parser.go:81
			// _ = "end of CoverTab[18437]"
//line /usr/local/go/src/crypto/x509/parser.go:81
		}() {
//line /usr/local/go/src/crypto/x509/parser.go:81
			_go_fuzz_dep_.CoverTab[18438]++
									value = value[:l-2]
//line /usr/local/go/src/crypto/x509/parser.go:82
			// _ = "end of CoverTab[18438]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:83
			_go_fuzz_dep_.CoverTab[18439]++
//line /usr/local/go/src/crypto/x509/parser.go:83
			// _ = "end of CoverTab[18439]"
//line /usr/local/go/src/crypto/x509/parser.go:83
		}
//line /usr/local/go/src/crypto/x509/parser.go:83
		// _ = "end of CoverTab[18421]"
//line /usr/local/go/src/crypto/x509/parser.go:83
		_go_fuzz_dep_.CoverTab[18422]++

								s := make([]uint16, 0, len(value)/2)
								for len(value) > 0 {
//line /usr/local/go/src/crypto/x509/parser.go:86
			_go_fuzz_dep_.CoverTab[18440]++
									s = append(s, uint16(value[0])<<8+uint16(value[1]))
									value = value[2:]
//line /usr/local/go/src/crypto/x509/parser.go:88
			// _ = "end of CoverTab[18440]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:89
		// _ = "end of CoverTab[18422]"
//line /usr/local/go/src/crypto/x509/parser.go:89
		_go_fuzz_dep_.CoverTab[18423]++

								return string(utf16.Decode(s)), nil
//line /usr/local/go/src/crypto/x509/parser.go:91
		// _ = "end of CoverTab[18423]"
	case cryptobyte_asn1.IA5String:
//line /usr/local/go/src/crypto/x509/parser.go:92
		_go_fuzz_dep_.CoverTab[18424]++
								s := string(value)
								if isIA5String(s) != nil {
//line /usr/local/go/src/crypto/x509/parser.go:94
			_go_fuzz_dep_.CoverTab[18441]++
									return "", errors.New("invalid IA5String")
//line /usr/local/go/src/crypto/x509/parser.go:95
			// _ = "end of CoverTab[18441]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:96
			_go_fuzz_dep_.CoverTab[18442]++
//line /usr/local/go/src/crypto/x509/parser.go:96
			// _ = "end of CoverTab[18442]"
//line /usr/local/go/src/crypto/x509/parser.go:96
		}
//line /usr/local/go/src/crypto/x509/parser.go:96
		// _ = "end of CoverTab[18424]"
//line /usr/local/go/src/crypto/x509/parser.go:96
		_go_fuzz_dep_.CoverTab[18425]++
								return s, nil
//line /usr/local/go/src/crypto/x509/parser.go:97
		// _ = "end of CoverTab[18425]"
	case cryptobyte_asn1.Tag(asn1.TagNumericString):
//line /usr/local/go/src/crypto/x509/parser.go:98
		_go_fuzz_dep_.CoverTab[18426]++
								for _, b := range value {
//line /usr/local/go/src/crypto/x509/parser.go:99
			_go_fuzz_dep_.CoverTab[18443]++
									if !('0' <= b && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:100
				_go_fuzz_dep_.CoverTab[18444]++
//line /usr/local/go/src/crypto/x509/parser.go:100
				return b <= '9'
//line /usr/local/go/src/crypto/x509/parser.go:100
				// _ = "end of CoverTab[18444]"
//line /usr/local/go/src/crypto/x509/parser.go:100
			}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:100
				_go_fuzz_dep_.CoverTab[18445]++
//line /usr/local/go/src/crypto/x509/parser.go:100
				return b == ' '
//line /usr/local/go/src/crypto/x509/parser.go:100
				// _ = "end of CoverTab[18445]"
//line /usr/local/go/src/crypto/x509/parser.go:100
			}()) {
//line /usr/local/go/src/crypto/x509/parser.go:100
				_go_fuzz_dep_.CoverTab[18446]++
										return "", errors.New("invalid NumericString")
//line /usr/local/go/src/crypto/x509/parser.go:101
				// _ = "end of CoverTab[18446]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:102
				_go_fuzz_dep_.CoverTab[18447]++
//line /usr/local/go/src/crypto/x509/parser.go:102
				// _ = "end of CoverTab[18447]"
//line /usr/local/go/src/crypto/x509/parser.go:102
			}
//line /usr/local/go/src/crypto/x509/parser.go:102
			// _ = "end of CoverTab[18443]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:103
		// _ = "end of CoverTab[18426]"
//line /usr/local/go/src/crypto/x509/parser.go:103
		_go_fuzz_dep_.CoverTab[18427]++
								return string(value), nil
//line /usr/local/go/src/crypto/x509/parser.go:104
		// _ = "end of CoverTab[18427]"
//line /usr/local/go/src/crypto/x509/parser.go:104
	default:
//line /usr/local/go/src/crypto/x509/parser.go:104
		_go_fuzz_dep_.CoverTab[18428]++
//line /usr/local/go/src/crypto/x509/parser.go:104
		// _ = "end of CoverTab[18428]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:105
	// _ = "end of CoverTab[18413]"
//line /usr/local/go/src/crypto/x509/parser.go:105
	_go_fuzz_dep_.CoverTab[18414]++
							return "", fmt.Errorf("unsupported string type: %v", tag)
//line /usr/local/go/src/crypto/x509/parser.go:106
	// _ = "end of CoverTab[18414]"
}

// parseName parses a DER encoded Name as defined in RFC 5280. We may
//line /usr/local/go/src/crypto/x509/parser.go:109
// want to export this function in the future for use in crypto/tls.
//line /usr/local/go/src/crypto/x509/parser.go:111
func parseName(raw cryptobyte.String) (*pkix.RDNSequence, error) {
//line /usr/local/go/src/crypto/x509/parser.go:111
	_go_fuzz_dep_.CoverTab[18448]++
							if !raw.ReadASN1(&raw, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:112
		_go_fuzz_dep_.CoverTab[18451]++
								return nil, errors.New("x509: invalid RDNSequence")
//line /usr/local/go/src/crypto/x509/parser.go:113
		// _ = "end of CoverTab[18451]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:114
		_go_fuzz_dep_.CoverTab[18452]++
//line /usr/local/go/src/crypto/x509/parser.go:114
		// _ = "end of CoverTab[18452]"
//line /usr/local/go/src/crypto/x509/parser.go:114
	}
//line /usr/local/go/src/crypto/x509/parser.go:114
	// _ = "end of CoverTab[18448]"
//line /usr/local/go/src/crypto/x509/parser.go:114
	_go_fuzz_dep_.CoverTab[18449]++

							var rdnSeq pkix.RDNSequence
							for !raw.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:117
		_go_fuzz_dep_.CoverTab[18453]++
								var rdnSet pkix.RelativeDistinguishedNameSET
								var set cryptobyte.String
								if !raw.ReadASN1(&set, cryptobyte_asn1.SET) {
//line /usr/local/go/src/crypto/x509/parser.go:120
			_go_fuzz_dep_.CoverTab[18456]++
									return nil, errors.New("x509: invalid RDNSequence")
//line /usr/local/go/src/crypto/x509/parser.go:121
			// _ = "end of CoverTab[18456]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:122
			_go_fuzz_dep_.CoverTab[18457]++
//line /usr/local/go/src/crypto/x509/parser.go:122
			// _ = "end of CoverTab[18457]"
//line /usr/local/go/src/crypto/x509/parser.go:122
		}
//line /usr/local/go/src/crypto/x509/parser.go:122
		// _ = "end of CoverTab[18453]"
//line /usr/local/go/src/crypto/x509/parser.go:122
		_go_fuzz_dep_.CoverTab[18454]++
								for !set.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:123
			_go_fuzz_dep_.CoverTab[18458]++
									var atav cryptobyte.String
									if !set.ReadASN1(&atav, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:125
				_go_fuzz_dep_.CoverTab[18463]++
										return nil, errors.New("x509: invalid RDNSequence: invalid attribute")
//line /usr/local/go/src/crypto/x509/parser.go:126
				// _ = "end of CoverTab[18463]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:127
				_go_fuzz_dep_.CoverTab[18464]++
//line /usr/local/go/src/crypto/x509/parser.go:127
				// _ = "end of CoverTab[18464]"
//line /usr/local/go/src/crypto/x509/parser.go:127
			}
//line /usr/local/go/src/crypto/x509/parser.go:127
			// _ = "end of CoverTab[18458]"
//line /usr/local/go/src/crypto/x509/parser.go:127
			_go_fuzz_dep_.CoverTab[18459]++
									var attr pkix.AttributeTypeAndValue
									if !atav.ReadASN1ObjectIdentifier(&attr.Type) {
//line /usr/local/go/src/crypto/x509/parser.go:129
				_go_fuzz_dep_.CoverTab[18465]++
										return nil, errors.New("x509: invalid RDNSequence: invalid attribute type")
//line /usr/local/go/src/crypto/x509/parser.go:130
				// _ = "end of CoverTab[18465]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:131
				_go_fuzz_dep_.CoverTab[18466]++
//line /usr/local/go/src/crypto/x509/parser.go:131
				// _ = "end of CoverTab[18466]"
//line /usr/local/go/src/crypto/x509/parser.go:131
			}
//line /usr/local/go/src/crypto/x509/parser.go:131
			// _ = "end of CoverTab[18459]"
//line /usr/local/go/src/crypto/x509/parser.go:131
			_go_fuzz_dep_.CoverTab[18460]++
									var rawValue cryptobyte.String
									var valueTag cryptobyte_asn1.Tag
									if !atav.ReadAnyASN1(&rawValue, &valueTag) {
//line /usr/local/go/src/crypto/x509/parser.go:134
				_go_fuzz_dep_.CoverTab[18467]++
										return nil, errors.New("x509: invalid RDNSequence: invalid attribute value")
//line /usr/local/go/src/crypto/x509/parser.go:135
				// _ = "end of CoverTab[18467]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:136
				_go_fuzz_dep_.CoverTab[18468]++
//line /usr/local/go/src/crypto/x509/parser.go:136
				// _ = "end of CoverTab[18468]"
//line /usr/local/go/src/crypto/x509/parser.go:136
			}
//line /usr/local/go/src/crypto/x509/parser.go:136
			// _ = "end of CoverTab[18460]"
//line /usr/local/go/src/crypto/x509/parser.go:136
			_go_fuzz_dep_.CoverTab[18461]++
									var err error
									attr.Value, err = parseASN1String(valueTag, rawValue)
									if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:139
				_go_fuzz_dep_.CoverTab[18469]++
										return nil, fmt.Errorf("x509: invalid RDNSequence: invalid attribute value: %s", err)
//line /usr/local/go/src/crypto/x509/parser.go:140
				// _ = "end of CoverTab[18469]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:141
				_go_fuzz_dep_.CoverTab[18470]++
//line /usr/local/go/src/crypto/x509/parser.go:141
				// _ = "end of CoverTab[18470]"
//line /usr/local/go/src/crypto/x509/parser.go:141
			}
//line /usr/local/go/src/crypto/x509/parser.go:141
			// _ = "end of CoverTab[18461]"
//line /usr/local/go/src/crypto/x509/parser.go:141
			_go_fuzz_dep_.CoverTab[18462]++
									rdnSet = append(rdnSet, attr)
//line /usr/local/go/src/crypto/x509/parser.go:142
			// _ = "end of CoverTab[18462]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:143
		// _ = "end of CoverTab[18454]"
//line /usr/local/go/src/crypto/x509/parser.go:143
		_go_fuzz_dep_.CoverTab[18455]++

								rdnSeq = append(rdnSeq, rdnSet)
//line /usr/local/go/src/crypto/x509/parser.go:145
		// _ = "end of CoverTab[18455]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:146
	// _ = "end of CoverTab[18449]"
//line /usr/local/go/src/crypto/x509/parser.go:146
	_go_fuzz_dep_.CoverTab[18450]++

							return &rdnSeq, nil
//line /usr/local/go/src/crypto/x509/parser.go:148
	// _ = "end of CoverTab[18450]"
}

func parseAI(der cryptobyte.String) (pkix.AlgorithmIdentifier, error) {
//line /usr/local/go/src/crypto/x509/parser.go:151
	_go_fuzz_dep_.CoverTab[18471]++
							ai := pkix.AlgorithmIdentifier{}
							if !der.ReadASN1ObjectIdentifier(&ai.Algorithm) {
//line /usr/local/go/src/crypto/x509/parser.go:153
		_go_fuzz_dep_.CoverTab[18475]++
								return ai, errors.New("x509: malformed OID")
//line /usr/local/go/src/crypto/x509/parser.go:154
		// _ = "end of CoverTab[18475]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:155
		_go_fuzz_dep_.CoverTab[18476]++
//line /usr/local/go/src/crypto/x509/parser.go:155
		// _ = "end of CoverTab[18476]"
//line /usr/local/go/src/crypto/x509/parser.go:155
	}
//line /usr/local/go/src/crypto/x509/parser.go:155
	// _ = "end of CoverTab[18471]"
//line /usr/local/go/src/crypto/x509/parser.go:155
	_go_fuzz_dep_.CoverTab[18472]++
							if der.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:156
		_go_fuzz_dep_.CoverTab[18477]++
								return ai, nil
//line /usr/local/go/src/crypto/x509/parser.go:157
		// _ = "end of CoverTab[18477]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:158
		_go_fuzz_dep_.CoverTab[18478]++
//line /usr/local/go/src/crypto/x509/parser.go:158
		// _ = "end of CoverTab[18478]"
//line /usr/local/go/src/crypto/x509/parser.go:158
	}
//line /usr/local/go/src/crypto/x509/parser.go:158
	// _ = "end of CoverTab[18472]"
//line /usr/local/go/src/crypto/x509/parser.go:158
	_go_fuzz_dep_.CoverTab[18473]++
							var params cryptobyte.String
							var tag cryptobyte_asn1.Tag
							if !der.ReadAnyASN1Element(&params, &tag) {
//line /usr/local/go/src/crypto/x509/parser.go:161
		_go_fuzz_dep_.CoverTab[18479]++
								return ai, errors.New("x509: malformed parameters")
//line /usr/local/go/src/crypto/x509/parser.go:162
		// _ = "end of CoverTab[18479]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:163
		_go_fuzz_dep_.CoverTab[18480]++
//line /usr/local/go/src/crypto/x509/parser.go:163
		// _ = "end of CoverTab[18480]"
//line /usr/local/go/src/crypto/x509/parser.go:163
	}
//line /usr/local/go/src/crypto/x509/parser.go:163
	// _ = "end of CoverTab[18473]"
//line /usr/local/go/src/crypto/x509/parser.go:163
	_go_fuzz_dep_.CoverTab[18474]++
							ai.Parameters.Tag = int(tag)
							ai.Parameters.FullBytes = params
							return ai, nil
//line /usr/local/go/src/crypto/x509/parser.go:166
	// _ = "end of CoverTab[18474]"
}

func parseTime(der *cryptobyte.String) (time.Time, error) {
//line /usr/local/go/src/crypto/x509/parser.go:169
	_go_fuzz_dep_.CoverTab[18481]++
							var t time.Time
							switch {
	case der.PeekASN1Tag(cryptobyte_asn1.UTCTime):
//line /usr/local/go/src/crypto/x509/parser.go:172
		_go_fuzz_dep_.CoverTab[18483]++
								if !der.ReadASN1UTCTime(&t) {
//line /usr/local/go/src/crypto/x509/parser.go:173
			_go_fuzz_dep_.CoverTab[18486]++
									return t, errors.New("x509: malformed UTCTime")
//line /usr/local/go/src/crypto/x509/parser.go:174
			// _ = "end of CoverTab[18486]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:175
			_go_fuzz_dep_.CoverTab[18487]++
//line /usr/local/go/src/crypto/x509/parser.go:175
			// _ = "end of CoverTab[18487]"
//line /usr/local/go/src/crypto/x509/parser.go:175
		}
//line /usr/local/go/src/crypto/x509/parser.go:175
		// _ = "end of CoverTab[18483]"
	case der.PeekASN1Tag(cryptobyte_asn1.GeneralizedTime):
//line /usr/local/go/src/crypto/x509/parser.go:176
		_go_fuzz_dep_.CoverTab[18484]++
								if !der.ReadASN1GeneralizedTime(&t) {
//line /usr/local/go/src/crypto/x509/parser.go:177
			_go_fuzz_dep_.CoverTab[18488]++
									return t, errors.New("x509: malformed GeneralizedTime")
//line /usr/local/go/src/crypto/x509/parser.go:178
			// _ = "end of CoverTab[18488]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:179
			_go_fuzz_dep_.CoverTab[18489]++
//line /usr/local/go/src/crypto/x509/parser.go:179
			// _ = "end of CoverTab[18489]"
//line /usr/local/go/src/crypto/x509/parser.go:179
		}
//line /usr/local/go/src/crypto/x509/parser.go:179
		// _ = "end of CoverTab[18484]"
	default:
//line /usr/local/go/src/crypto/x509/parser.go:180
		_go_fuzz_dep_.CoverTab[18485]++
								return t, errors.New("x509: unsupported time format")
//line /usr/local/go/src/crypto/x509/parser.go:181
		// _ = "end of CoverTab[18485]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:182
	// _ = "end of CoverTab[18481]"
//line /usr/local/go/src/crypto/x509/parser.go:182
	_go_fuzz_dep_.CoverTab[18482]++
							return t, nil
//line /usr/local/go/src/crypto/x509/parser.go:183
	// _ = "end of CoverTab[18482]"
}

func parseValidity(der cryptobyte.String) (time.Time, time.Time, error) {
//line /usr/local/go/src/crypto/x509/parser.go:186
	_go_fuzz_dep_.CoverTab[18490]++
							notBefore, err := parseTime(&der)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:188
		_go_fuzz_dep_.CoverTab[18493]++
								return time.Time{}, time.Time{}, err
//line /usr/local/go/src/crypto/x509/parser.go:189
		// _ = "end of CoverTab[18493]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:190
		_go_fuzz_dep_.CoverTab[18494]++
//line /usr/local/go/src/crypto/x509/parser.go:190
		// _ = "end of CoverTab[18494]"
//line /usr/local/go/src/crypto/x509/parser.go:190
	}
//line /usr/local/go/src/crypto/x509/parser.go:190
	// _ = "end of CoverTab[18490]"
//line /usr/local/go/src/crypto/x509/parser.go:190
	_go_fuzz_dep_.CoverTab[18491]++
							notAfter, err := parseTime(&der)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:192
		_go_fuzz_dep_.CoverTab[18495]++
								return time.Time{}, time.Time{}, err
//line /usr/local/go/src/crypto/x509/parser.go:193
		// _ = "end of CoverTab[18495]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:194
		_go_fuzz_dep_.CoverTab[18496]++
//line /usr/local/go/src/crypto/x509/parser.go:194
		// _ = "end of CoverTab[18496]"
//line /usr/local/go/src/crypto/x509/parser.go:194
	}
//line /usr/local/go/src/crypto/x509/parser.go:194
	// _ = "end of CoverTab[18491]"
//line /usr/local/go/src/crypto/x509/parser.go:194
	_go_fuzz_dep_.CoverTab[18492]++

							return notBefore, notAfter, nil
//line /usr/local/go/src/crypto/x509/parser.go:196
	// _ = "end of CoverTab[18492]"
}

func parseExtension(der cryptobyte.String) (pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/parser.go:199
	_go_fuzz_dep_.CoverTab[18497]++
							var ext pkix.Extension
							if !der.ReadASN1ObjectIdentifier(&ext.Id) {
//line /usr/local/go/src/crypto/x509/parser.go:201
		_go_fuzz_dep_.CoverTab[18501]++
								return ext, errors.New("x509: malformed extension OID field")
//line /usr/local/go/src/crypto/x509/parser.go:202
		// _ = "end of CoverTab[18501]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:203
		_go_fuzz_dep_.CoverTab[18502]++
//line /usr/local/go/src/crypto/x509/parser.go:203
		// _ = "end of CoverTab[18502]"
//line /usr/local/go/src/crypto/x509/parser.go:203
	}
//line /usr/local/go/src/crypto/x509/parser.go:203
	// _ = "end of CoverTab[18497]"
//line /usr/local/go/src/crypto/x509/parser.go:203
	_go_fuzz_dep_.CoverTab[18498]++
							if der.PeekASN1Tag(cryptobyte_asn1.BOOLEAN) {
//line /usr/local/go/src/crypto/x509/parser.go:204
		_go_fuzz_dep_.CoverTab[18503]++
								if !der.ReadASN1Boolean(&ext.Critical) {
//line /usr/local/go/src/crypto/x509/parser.go:205
			_go_fuzz_dep_.CoverTab[18504]++
									return ext, errors.New("x509: malformed extension critical field")
//line /usr/local/go/src/crypto/x509/parser.go:206
			// _ = "end of CoverTab[18504]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:207
			_go_fuzz_dep_.CoverTab[18505]++
//line /usr/local/go/src/crypto/x509/parser.go:207
			// _ = "end of CoverTab[18505]"
//line /usr/local/go/src/crypto/x509/parser.go:207
		}
//line /usr/local/go/src/crypto/x509/parser.go:207
		// _ = "end of CoverTab[18503]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:208
		_go_fuzz_dep_.CoverTab[18506]++
//line /usr/local/go/src/crypto/x509/parser.go:208
		// _ = "end of CoverTab[18506]"
//line /usr/local/go/src/crypto/x509/parser.go:208
	}
//line /usr/local/go/src/crypto/x509/parser.go:208
	// _ = "end of CoverTab[18498]"
//line /usr/local/go/src/crypto/x509/parser.go:208
	_go_fuzz_dep_.CoverTab[18499]++
							var val cryptobyte.String
							if !der.ReadASN1(&val, cryptobyte_asn1.OCTET_STRING) {
//line /usr/local/go/src/crypto/x509/parser.go:210
		_go_fuzz_dep_.CoverTab[18507]++
								return ext, errors.New("x509: malformed extension value field")
//line /usr/local/go/src/crypto/x509/parser.go:211
		// _ = "end of CoverTab[18507]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:212
		_go_fuzz_dep_.CoverTab[18508]++
//line /usr/local/go/src/crypto/x509/parser.go:212
		// _ = "end of CoverTab[18508]"
//line /usr/local/go/src/crypto/x509/parser.go:212
	}
//line /usr/local/go/src/crypto/x509/parser.go:212
	// _ = "end of CoverTab[18499]"
//line /usr/local/go/src/crypto/x509/parser.go:212
	_go_fuzz_dep_.CoverTab[18500]++
							ext.Value = val
							return ext, nil
//line /usr/local/go/src/crypto/x509/parser.go:214
	// _ = "end of CoverTab[18500]"
}

func parsePublicKey(keyData *publicKeyInfo) (any, error) {
//line /usr/local/go/src/crypto/x509/parser.go:217
	_go_fuzz_dep_.CoverTab[18509]++
							oid := keyData.Algorithm.Algorithm
							params := keyData.Algorithm.Parameters
							der := cryptobyte.String(keyData.PublicKey.RightAlign())
							switch {
	case oid.Equal(oidPublicKeyRSA):
//line /usr/local/go/src/crypto/x509/parser.go:222
		_go_fuzz_dep_.CoverTab[18510]++

//line /usr/local/go/src/crypto/x509/parser.go:225
		if !bytes.Equal(params.FullBytes, asn1.NullBytes) {
//line /usr/local/go/src/crypto/x509/parser.go:225
			_go_fuzz_dep_.CoverTab[18531]++
									return nil, errors.New("x509: RSA key missing NULL parameters")
//line /usr/local/go/src/crypto/x509/parser.go:226
			// _ = "end of CoverTab[18531]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:227
			_go_fuzz_dep_.CoverTab[18532]++
//line /usr/local/go/src/crypto/x509/parser.go:227
			// _ = "end of CoverTab[18532]"
//line /usr/local/go/src/crypto/x509/parser.go:227
		}
//line /usr/local/go/src/crypto/x509/parser.go:227
		// _ = "end of CoverTab[18510]"
//line /usr/local/go/src/crypto/x509/parser.go:227
		_go_fuzz_dep_.CoverTab[18511]++

								p := &pkcs1PublicKey{N: new(big.Int)}
								if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:230
			_go_fuzz_dep_.CoverTab[18533]++
									return nil, errors.New("x509: invalid RSA public key")
//line /usr/local/go/src/crypto/x509/parser.go:231
			// _ = "end of CoverTab[18533]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:232
			_go_fuzz_dep_.CoverTab[18534]++
//line /usr/local/go/src/crypto/x509/parser.go:232
			// _ = "end of CoverTab[18534]"
//line /usr/local/go/src/crypto/x509/parser.go:232
		}
//line /usr/local/go/src/crypto/x509/parser.go:232
		// _ = "end of CoverTab[18511]"
//line /usr/local/go/src/crypto/x509/parser.go:232
		_go_fuzz_dep_.CoverTab[18512]++
								if !der.ReadASN1Integer(p.N) {
//line /usr/local/go/src/crypto/x509/parser.go:233
			_go_fuzz_dep_.CoverTab[18535]++
									return nil, errors.New("x509: invalid RSA modulus")
//line /usr/local/go/src/crypto/x509/parser.go:234
			// _ = "end of CoverTab[18535]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:235
			_go_fuzz_dep_.CoverTab[18536]++
//line /usr/local/go/src/crypto/x509/parser.go:235
			// _ = "end of CoverTab[18536]"
//line /usr/local/go/src/crypto/x509/parser.go:235
		}
//line /usr/local/go/src/crypto/x509/parser.go:235
		// _ = "end of CoverTab[18512]"
//line /usr/local/go/src/crypto/x509/parser.go:235
		_go_fuzz_dep_.CoverTab[18513]++
								if !der.ReadASN1Integer(&p.E) {
//line /usr/local/go/src/crypto/x509/parser.go:236
			_go_fuzz_dep_.CoverTab[18537]++
									return nil, errors.New("x509: invalid RSA public exponent")
//line /usr/local/go/src/crypto/x509/parser.go:237
			// _ = "end of CoverTab[18537]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:238
			_go_fuzz_dep_.CoverTab[18538]++
//line /usr/local/go/src/crypto/x509/parser.go:238
			// _ = "end of CoverTab[18538]"
//line /usr/local/go/src/crypto/x509/parser.go:238
		}
//line /usr/local/go/src/crypto/x509/parser.go:238
		// _ = "end of CoverTab[18513]"
//line /usr/local/go/src/crypto/x509/parser.go:238
		_go_fuzz_dep_.CoverTab[18514]++

								if p.N.Sign() <= 0 {
//line /usr/local/go/src/crypto/x509/parser.go:240
			_go_fuzz_dep_.CoverTab[18539]++
									return nil, errors.New("x509: RSA modulus is not a positive number")
//line /usr/local/go/src/crypto/x509/parser.go:241
			// _ = "end of CoverTab[18539]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:242
			_go_fuzz_dep_.CoverTab[18540]++
//line /usr/local/go/src/crypto/x509/parser.go:242
			// _ = "end of CoverTab[18540]"
//line /usr/local/go/src/crypto/x509/parser.go:242
		}
//line /usr/local/go/src/crypto/x509/parser.go:242
		// _ = "end of CoverTab[18514]"
//line /usr/local/go/src/crypto/x509/parser.go:242
		_go_fuzz_dep_.CoverTab[18515]++
								if p.E <= 0 {
//line /usr/local/go/src/crypto/x509/parser.go:243
			_go_fuzz_dep_.CoverTab[18541]++
									return nil, errors.New("x509: RSA public exponent is not a positive number")
//line /usr/local/go/src/crypto/x509/parser.go:244
			// _ = "end of CoverTab[18541]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:245
			_go_fuzz_dep_.CoverTab[18542]++
//line /usr/local/go/src/crypto/x509/parser.go:245
			// _ = "end of CoverTab[18542]"
//line /usr/local/go/src/crypto/x509/parser.go:245
		}
//line /usr/local/go/src/crypto/x509/parser.go:245
		// _ = "end of CoverTab[18515]"
//line /usr/local/go/src/crypto/x509/parser.go:245
		_go_fuzz_dep_.CoverTab[18516]++

								pub := &rsa.PublicKey{
			E:	p.E,
			N:	p.N,
		}
								return pub, nil
//line /usr/local/go/src/crypto/x509/parser.go:251
		// _ = "end of CoverTab[18516]"
	case oid.Equal(oidPublicKeyECDSA):
//line /usr/local/go/src/crypto/x509/parser.go:252
		_go_fuzz_dep_.CoverTab[18517]++
								paramsDer := cryptobyte.String(params.FullBytes)
								namedCurveOID := new(asn1.ObjectIdentifier)
								if !paramsDer.ReadASN1ObjectIdentifier(namedCurveOID) {
//line /usr/local/go/src/crypto/x509/parser.go:255
			_go_fuzz_dep_.CoverTab[18543]++
									return nil, errors.New("x509: invalid ECDSA parameters")
//line /usr/local/go/src/crypto/x509/parser.go:256
			// _ = "end of CoverTab[18543]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:257
			_go_fuzz_dep_.CoverTab[18544]++
//line /usr/local/go/src/crypto/x509/parser.go:257
			// _ = "end of CoverTab[18544]"
//line /usr/local/go/src/crypto/x509/parser.go:257
		}
//line /usr/local/go/src/crypto/x509/parser.go:257
		// _ = "end of CoverTab[18517]"
//line /usr/local/go/src/crypto/x509/parser.go:257
		_go_fuzz_dep_.CoverTab[18518]++
								namedCurve := namedCurveFromOID(*namedCurveOID)
								if namedCurve == nil {
//line /usr/local/go/src/crypto/x509/parser.go:259
			_go_fuzz_dep_.CoverTab[18545]++
									return nil, errors.New("x509: unsupported elliptic curve")
//line /usr/local/go/src/crypto/x509/parser.go:260
			// _ = "end of CoverTab[18545]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:261
			_go_fuzz_dep_.CoverTab[18546]++
//line /usr/local/go/src/crypto/x509/parser.go:261
			// _ = "end of CoverTab[18546]"
//line /usr/local/go/src/crypto/x509/parser.go:261
		}
//line /usr/local/go/src/crypto/x509/parser.go:261
		// _ = "end of CoverTab[18518]"
//line /usr/local/go/src/crypto/x509/parser.go:261
		_go_fuzz_dep_.CoverTab[18519]++
								x, y := elliptic.Unmarshal(namedCurve, der)
								if x == nil {
//line /usr/local/go/src/crypto/x509/parser.go:263
			_go_fuzz_dep_.CoverTab[18547]++
									return nil, errors.New("x509: failed to unmarshal elliptic curve point")
//line /usr/local/go/src/crypto/x509/parser.go:264
			// _ = "end of CoverTab[18547]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:265
			_go_fuzz_dep_.CoverTab[18548]++
//line /usr/local/go/src/crypto/x509/parser.go:265
			// _ = "end of CoverTab[18548]"
//line /usr/local/go/src/crypto/x509/parser.go:265
		}
//line /usr/local/go/src/crypto/x509/parser.go:265
		// _ = "end of CoverTab[18519]"
//line /usr/local/go/src/crypto/x509/parser.go:265
		_go_fuzz_dep_.CoverTab[18520]++
								pub := &ecdsa.PublicKey{
			Curve:	namedCurve,
			X:	x,
			Y:	y,
		}
								return pub, nil
//line /usr/local/go/src/crypto/x509/parser.go:271
		// _ = "end of CoverTab[18520]"
	case oid.Equal(oidPublicKeyEd25519):
//line /usr/local/go/src/crypto/x509/parser.go:272
		_go_fuzz_dep_.CoverTab[18521]++

//line /usr/local/go/src/crypto/x509/parser.go:275
		if len(params.FullBytes) != 0 {
//line /usr/local/go/src/crypto/x509/parser.go:275
			_go_fuzz_dep_.CoverTab[18549]++
									return nil, errors.New("x509: Ed25519 key encoded with illegal parameters")
//line /usr/local/go/src/crypto/x509/parser.go:276
			// _ = "end of CoverTab[18549]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:277
			_go_fuzz_dep_.CoverTab[18550]++
//line /usr/local/go/src/crypto/x509/parser.go:277
			// _ = "end of CoverTab[18550]"
//line /usr/local/go/src/crypto/x509/parser.go:277
		}
//line /usr/local/go/src/crypto/x509/parser.go:277
		// _ = "end of CoverTab[18521]"
//line /usr/local/go/src/crypto/x509/parser.go:277
		_go_fuzz_dep_.CoverTab[18522]++
								if len(der) != ed25519.PublicKeySize {
//line /usr/local/go/src/crypto/x509/parser.go:278
			_go_fuzz_dep_.CoverTab[18551]++
									return nil, errors.New("x509: wrong Ed25519 public key size")
//line /usr/local/go/src/crypto/x509/parser.go:279
			// _ = "end of CoverTab[18551]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:280
			_go_fuzz_dep_.CoverTab[18552]++
//line /usr/local/go/src/crypto/x509/parser.go:280
			// _ = "end of CoverTab[18552]"
//line /usr/local/go/src/crypto/x509/parser.go:280
		}
//line /usr/local/go/src/crypto/x509/parser.go:280
		// _ = "end of CoverTab[18522]"
//line /usr/local/go/src/crypto/x509/parser.go:280
		_go_fuzz_dep_.CoverTab[18523]++
								return ed25519.PublicKey(der), nil
//line /usr/local/go/src/crypto/x509/parser.go:281
		// _ = "end of CoverTab[18523]"
	case oid.Equal(oidPublicKeyX25519):
//line /usr/local/go/src/crypto/x509/parser.go:282
		_go_fuzz_dep_.CoverTab[18524]++

//line /usr/local/go/src/crypto/x509/parser.go:285
		if len(params.FullBytes) != 0 {
//line /usr/local/go/src/crypto/x509/parser.go:285
			_go_fuzz_dep_.CoverTab[18553]++
									return nil, errors.New("x509: X25519 key encoded with illegal parameters")
//line /usr/local/go/src/crypto/x509/parser.go:286
			// _ = "end of CoverTab[18553]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:287
			_go_fuzz_dep_.CoverTab[18554]++
//line /usr/local/go/src/crypto/x509/parser.go:287
			// _ = "end of CoverTab[18554]"
//line /usr/local/go/src/crypto/x509/parser.go:287
		}
//line /usr/local/go/src/crypto/x509/parser.go:287
		// _ = "end of CoverTab[18524]"
//line /usr/local/go/src/crypto/x509/parser.go:287
		_go_fuzz_dep_.CoverTab[18525]++
								return ecdh.X25519().NewPublicKey(der)
//line /usr/local/go/src/crypto/x509/parser.go:288
		// _ = "end of CoverTab[18525]"
	case oid.Equal(oidPublicKeyDSA):
//line /usr/local/go/src/crypto/x509/parser.go:289
		_go_fuzz_dep_.CoverTab[18526]++
								y := new(big.Int)
								if !der.ReadASN1Integer(y) {
//line /usr/local/go/src/crypto/x509/parser.go:291
			_go_fuzz_dep_.CoverTab[18555]++
									return nil, errors.New("x509: invalid DSA public key")
//line /usr/local/go/src/crypto/x509/parser.go:292
			// _ = "end of CoverTab[18555]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:293
			_go_fuzz_dep_.CoverTab[18556]++
//line /usr/local/go/src/crypto/x509/parser.go:293
			// _ = "end of CoverTab[18556]"
//line /usr/local/go/src/crypto/x509/parser.go:293
		}
//line /usr/local/go/src/crypto/x509/parser.go:293
		// _ = "end of CoverTab[18526]"
//line /usr/local/go/src/crypto/x509/parser.go:293
		_go_fuzz_dep_.CoverTab[18527]++
								pub := &dsa.PublicKey{
			Y:	y,
			Parameters: dsa.Parameters{
				P:	new(big.Int),
				Q:	new(big.Int),
				G:	new(big.Int),
			},
		}
		paramsDer := cryptobyte.String(params.FullBytes)
		if !paramsDer.ReadASN1(&paramsDer, cryptobyte_asn1.SEQUENCE) || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:303
			_go_fuzz_dep_.CoverTab[18557]++
//line /usr/local/go/src/crypto/x509/parser.go:303
			return !paramsDer.ReadASN1Integer(pub.Parameters.P)
									// _ = "end of CoverTab[18557]"
//line /usr/local/go/src/crypto/x509/parser.go:304
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:304
			_go_fuzz_dep_.CoverTab[18558]++
//line /usr/local/go/src/crypto/x509/parser.go:304
			return !paramsDer.ReadASN1Integer(pub.Parameters.Q)
									// _ = "end of CoverTab[18558]"
//line /usr/local/go/src/crypto/x509/parser.go:305
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:305
			_go_fuzz_dep_.CoverTab[18559]++
//line /usr/local/go/src/crypto/x509/parser.go:305
			return !paramsDer.ReadASN1Integer(pub.Parameters.G)
									// _ = "end of CoverTab[18559]"
//line /usr/local/go/src/crypto/x509/parser.go:306
		}() {
//line /usr/local/go/src/crypto/x509/parser.go:306
			_go_fuzz_dep_.CoverTab[18560]++
									return nil, errors.New("x509: invalid DSA parameters")
//line /usr/local/go/src/crypto/x509/parser.go:307
			// _ = "end of CoverTab[18560]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:308
			_go_fuzz_dep_.CoverTab[18561]++
//line /usr/local/go/src/crypto/x509/parser.go:308
			// _ = "end of CoverTab[18561]"
//line /usr/local/go/src/crypto/x509/parser.go:308
		}
//line /usr/local/go/src/crypto/x509/parser.go:308
		// _ = "end of CoverTab[18527]"
//line /usr/local/go/src/crypto/x509/parser.go:308
		_go_fuzz_dep_.CoverTab[18528]++
								if pub.Y.Sign() <= 0 || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:309
			_go_fuzz_dep_.CoverTab[18562]++
//line /usr/local/go/src/crypto/x509/parser.go:309
			return pub.Parameters.P.Sign() <= 0
//line /usr/local/go/src/crypto/x509/parser.go:309
			// _ = "end of CoverTab[18562]"
//line /usr/local/go/src/crypto/x509/parser.go:309
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:309
			_go_fuzz_dep_.CoverTab[18563]++
//line /usr/local/go/src/crypto/x509/parser.go:309
			return pub.Parameters.Q.Sign() <= 0
									// _ = "end of CoverTab[18563]"
//line /usr/local/go/src/crypto/x509/parser.go:310
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:310
			_go_fuzz_dep_.CoverTab[18564]++
//line /usr/local/go/src/crypto/x509/parser.go:310
			return pub.Parameters.G.Sign() <= 0
//line /usr/local/go/src/crypto/x509/parser.go:310
			// _ = "end of CoverTab[18564]"
//line /usr/local/go/src/crypto/x509/parser.go:310
		}() {
//line /usr/local/go/src/crypto/x509/parser.go:310
			_go_fuzz_dep_.CoverTab[18565]++
									return nil, errors.New("x509: zero or negative DSA parameter")
//line /usr/local/go/src/crypto/x509/parser.go:311
			// _ = "end of CoverTab[18565]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:312
			_go_fuzz_dep_.CoverTab[18566]++
//line /usr/local/go/src/crypto/x509/parser.go:312
			// _ = "end of CoverTab[18566]"
//line /usr/local/go/src/crypto/x509/parser.go:312
		}
//line /usr/local/go/src/crypto/x509/parser.go:312
		// _ = "end of CoverTab[18528]"
//line /usr/local/go/src/crypto/x509/parser.go:312
		_go_fuzz_dep_.CoverTab[18529]++
								return pub, nil
//line /usr/local/go/src/crypto/x509/parser.go:313
		// _ = "end of CoverTab[18529]"
	default:
//line /usr/local/go/src/crypto/x509/parser.go:314
		_go_fuzz_dep_.CoverTab[18530]++
								return nil, errors.New("x509: unknown public key algorithm")
//line /usr/local/go/src/crypto/x509/parser.go:315
		// _ = "end of CoverTab[18530]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:316
	// _ = "end of CoverTab[18509]"
}

func parseKeyUsageExtension(der cryptobyte.String) (KeyUsage, error) {
//line /usr/local/go/src/crypto/x509/parser.go:319
	_go_fuzz_dep_.CoverTab[18567]++
							var usageBits asn1.BitString
							if !der.ReadASN1BitString(&usageBits) {
//line /usr/local/go/src/crypto/x509/parser.go:321
		_go_fuzz_dep_.CoverTab[18570]++
								return 0, errors.New("x509: invalid key usage")
//line /usr/local/go/src/crypto/x509/parser.go:322
		// _ = "end of CoverTab[18570]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:323
		_go_fuzz_dep_.CoverTab[18571]++
//line /usr/local/go/src/crypto/x509/parser.go:323
		// _ = "end of CoverTab[18571]"
//line /usr/local/go/src/crypto/x509/parser.go:323
	}
//line /usr/local/go/src/crypto/x509/parser.go:323
	// _ = "end of CoverTab[18567]"
//line /usr/local/go/src/crypto/x509/parser.go:323
	_go_fuzz_dep_.CoverTab[18568]++

							var usage int
							for i := 0; i < 9; i++ {
//line /usr/local/go/src/crypto/x509/parser.go:326
		_go_fuzz_dep_.CoverTab[18572]++
								if usageBits.At(i) != 0 {
//line /usr/local/go/src/crypto/x509/parser.go:327
			_go_fuzz_dep_.CoverTab[18573]++
									usage |= 1 << uint(i)
//line /usr/local/go/src/crypto/x509/parser.go:328
			// _ = "end of CoverTab[18573]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:329
			_go_fuzz_dep_.CoverTab[18574]++
//line /usr/local/go/src/crypto/x509/parser.go:329
			// _ = "end of CoverTab[18574]"
//line /usr/local/go/src/crypto/x509/parser.go:329
		}
//line /usr/local/go/src/crypto/x509/parser.go:329
		// _ = "end of CoverTab[18572]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:330
	// _ = "end of CoverTab[18568]"
//line /usr/local/go/src/crypto/x509/parser.go:330
	_go_fuzz_dep_.CoverTab[18569]++
							return KeyUsage(usage), nil
//line /usr/local/go/src/crypto/x509/parser.go:331
	// _ = "end of CoverTab[18569]"
}

func parseBasicConstraintsExtension(der cryptobyte.String) (bool, int, error) {
//line /usr/local/go/src/crypto/x509/parser.go:334
	_go_fuzz_dep_.CoverTab[18575]++
							var isCA bool
							if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:336
		_go_fuzz_dep_.CoverTab[18579]++
								return false, 0, errors.New("x509: invalid basic constraints a")
//line /usr/local/go/src/crypto/x509/parser.go:337
		// _ = "end of CoverTab[18579]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:338
		_go_fuzz_dep_.CoverTab[18580]++
//line /usr/local/go/src/crypto/x509/parser.go:338
		// _ = "end of CoverTab[18580]"
//line /usr/local/go/src/crypto/x509/parser.go:338
	}
//line /usr/local/go/src/crypto/x509/parser.go:338
	// _ = "end of CoverTab[18575]"
//line /usr/local/go/src/crypto/x509/parser.go:338
	_go_fuzz_dep_.CoverTab[18576]++
							if der.PeekASN1Tag(cryptobyte_asn1.BOOLEAN) {
//line /usr/local/go/src/crypto/x509/parser.go:339
		_go_fuzz_dep_.CoverTab[18581]++
								if !der.ReadASN1Boolean(&isCA) {
//line /usr/local/go/src/crypto/x509/parser.go:340
			_go_fuzz_dep_.CoverTab[18582]++
									return false, 0, errors.New("x509: invalid basic constraints b")
//line /usr/local/go/src/crypto/x509/parser.go:341
			// _ = "end of CoverTab[18582]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:342
			_go_fuzz_dep_.CoverTab[18583]++
//line /usr/local/go/src/crypto/x509/parser.go:342
			// _ = "end of CoverTab[18583]"
//line /usr/local/go/src/crypto/x509/parser.go:342
		}
//line /usr/local/go/src/crypto/x509/parser.go:342
		// _ = "end of CoverTab[18581]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:343
		_go_fuzz_dep_.CoverTab[18584]++
//line /usr/local/go/src/crypto/x509/parser.go:343
		// _ = "end of CoverTab[18584]"
//line /usr/local/go/src/crypto/x509/parser.go:343
	}
//line /usr/local/go/src/crypto/x509/parser.go:343
	// _ = "end of CoverTab[18576]"
//line /usr/local/go/src/crypto/x509/parser.go:343
	_go_fuzz_dep_.CoverTab[18577]++
							maxPathLen := -1
							if !der.Empty() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:345
		_go_fuzz_dep_.CoverTab[18585]++
//line /usr/local/go/src/crypto/x509/parser.go:345
		return der.PeekASN1Tag(cryptobyte_asn1.INTEGER)
//line /usr/local/go/src/crypto/x509/parser.go:345
		// _ = "end of CoverTab[18585]"
//line /usr/local/go/src/crypto/x509/parser.go:345
	}() {
//line /usr/local/go/src/crypto/x509/parser.go:345
		_go_fuzz_dep_.CoverTab[18586]++
								if !der.ReadASN1Integer(&maxPathLen) {
//line /usr/local/go/src/crypto/x509/parser.go:346
			_go_fuzz_dep_.CoverTab[18587]++
									return false, 0, errors.New("x509: invalid basic constraints c")
//line /usr/local/go/src/crypto/x509/parser.go:347
			// _ = "end of CoverTab[18587]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:348
			_go_fuzz_dep_.CoverTab[18588]++
//line /usr/local/go/src/crypto/x509/parser.go:348
			// _ = "end of CoverTab[18588]"
//line /usr/local/go/src/crypto/x509/parser.go:348
		}
//line /usr/local/go/src/crypto/x509/parser.go:348
		// _ = "end of CoverTab[18586]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:349
		_go_fuzz_dep_.CoverTab[18589]++
//line /usr/local/go/src/crypto/x509/parser.go:349
		// _ = "end of CoverTab[18589]"
//line /usr/local/go/src/crypto/x509/parser.go:349
	}
//line /usr/local/go/src/crypto/x509/parser.go:349
	// _ = "end of CoverTab[18577]"
//line /usr/local/go/src/crypto/x509/parser.go:349
	_go_fuzz_dep_.CoverTab[18578]++

//line /usr/local/go/src/crypto/x509/parser.go:352
	return isCA, maxPathLen, nil
//line /usr/local/go/src/crypto/x509/parser.go:352
	// _ = "end of CoverTab[18578]"
}

func forEachSAN(der cryptobyte.String, callback func(tag int, data []byte) error) error {
//line /usr/local/go/src/crypto/x509/parser.go:355
	_go_fuzz_dep_.CoverTab[18590]++
							if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:356
		_go_fuzz_dep_.CoverTab[18593]++
								return errors.New("x509: invalid subject alternative names")
//line /usr/local/go/src/crypto/x509/parser.go:357
		// _ = "end of CoverTab[18593]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:358
		_go_fuzz_dep_.CoverTab[18594]++
//line /usr/local/go/src/crypto/x509/parser.go:358
		// _ = "end of CoverTab[18594]"
//line /usr/local/go/src/crypto/x509/parser.go:358
	}
//line /usr/local/go/src/crypto/x509/parser.go:358
	// _ = "end of CoverTab[18590]"
//line /usr/local/go/src/crypto/x509/parser.go:358
	_go_fuzz_dep_.CoverTab[18591]++
							for !der.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:359
		_go_fuzz_dep_.CoverTab[18595]++
								var san cryptobyte.String
								var tag cryptobyte_asn1.Tag
								if !der.ReadAnyASN1(&san, &tag) {
//line /usr/local/go/src/crypto/x509/parser.go:362
			_go_fuzz_dep_.CoverTab[18597]++
									return errors.New("x509: invalid subject alternative name")
//line /usr/local/go/src/crypto/x509/parser.go:363
			// _ = "end of CoverTab[18597]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:364
			_go_fuzz_dep_.CoverTab[18598]++
//line /usr/local/go/src/crypto/x509/parser.go:364
			// _ = "end of CoverTab[18598]"
//line /usr/local/go/src/crypto/x509/parser.go:364
		}
//line /usr/local/go/src/crypto/x509/parser.go:364
		// _ = "end of CoverTab[18595]"
//line /usr/local/go/src/crypto/x509/parser.go:364
		_go_fuzz_dep_.CoverTab[18596]++
								if err := callback(int(tag^0x80), san); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:365
			_go_fuzz_dep_.CoverTab[18599]++
									return err
//line /usr/local/go/src/crypto/x509/parser.go:366
			// _ = "end of CoverTab[18599]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:367
			_go_fuzz_dep_.CoverTab[18600]++
//line /usr/local/go/src/crypto/x509/parser.go:367
			// _ = "end of CoverTab[18600]"
//line /usr/local/go/src/crypto/x509/parser.go:367
		}
//line /usr/local/go/src/crypto/x509/parser.go:367
		// _ = "end of CoverTab[18596]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:368
	// _ = "end of CoverTab[18591]"
//line /usr/local/go/src/crypto/x509/parser.go:368
	_go_fuzz_dep_.CoverTab[18592]++

							return nil
//line /usr/local/go/src/crypto/x509/parser.go:370
	// _ = "end of CoverTab[18592]"
}

func parseSANExtension(der cryptobyte.String) (dnsNames, emailAddresses []string, ipAddresses []net.IP, uris []*url.URL, err error) {
//line /usr/local/go/src/crypto/x509/parser.go:373
	_go_fuzz_dep_.CoverTab[18601]++
							err = forEachSAN(der, func(tag int, data []byte) error {
//line /usr/local/go/src/crypto/x509/parser.go:374
		_go_fuzz_dep_.CoverTab[18603]++
								switch tag {
		case nameTypeEmail:
//line /usr/local/go/src/crypto/x509/parser.go:376
			_go_fuzz_dep_.CoverTab[18605]++
									email := string(data)
									if err := isIA5String(email); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:378
				_go_fuzz_dep_.CoverTab[18615]++
										return errors.New("x509: SAN rfc822Name is malformed")
//line /usr/local/go/src/crypto/x509/parser.go:379
				// _ = "end of CoverTab[18615]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:380
				_go_fuzz_dep_.CoverTab[18616]++
//line /usr/local/go/src/crypto/x509/parser.go:380
				// _ = "end of CoverTab[18616]"
//line /usr/local/go/src/crypto/x509/parser.go:380
			}
//line /usr/local/go/src/crypto/x509/parser.go:380
			// _ = "end of CoverTab[18605]"
//line /usr/local/go/src/crypto/x509/parser.go:380
			_go_fuzz_dep_.CoverTab[18606]++
									emailAddresses = append(emailAddresses, email)
//line /usr/local/go/src/crypto/x509/parser.go:381
			// _ = "end of CoverTab[18606]"
		case nameTypeDNS:
//line /usr/local/go/src/crypto/x509/parser.go:382
			_go_fuzz_dep_.CoverTab[18607]++
									name := string(data)
									if err := isIA5String(name); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:384
				_go_fuzz_dep_.CoverTab[18617]++
										return errors.New("x509: SAN dNSName is malformed")
//line /usr/local/go/src/crypto/x509/parser.go:385
				// _ = "end of CoverTab[18617]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:386
				_go_fuzz_dep_.CoverTab[18618]++
//line /usr/local/go/src/crypto/x509/parser.go:386
				// _ = "end of CoverTab[18618]"
//line /usr/local/go/src/crypto/x509/parser.go:386
			}
//line /usr/local/go/src/crypto/x509/parser.go:386
			// _ = "end of CoverTab[18607]"
//line /usr/local/go/src/crypto/x509/parser.go:386
			_go_fuzz_dep_.CoverTab[18608]++
									dnsNames = append(dnsNames, string(name))
//line /usr/local/go/src/crypto/x509/parser.go:387
			// _ = "end of CoverTab[18608]"
		case nameTypeURI:
//line /usr/local/go/src/crypto/x509/parser.go:388
			_go_fuzz_dep_.CoverTab[18609]++
									uriStr := string(data)
									if err := isIA5String(uriStr); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:390
				_go_fuzz_dep_.CoverTab[18619]++
										return errors.New("x509: SAN uniformResourceIdentifier is malformed")
//line /usr/local/go/src/crypto/x509/parser.go:391
				// _ = "end of CoverTab[18619]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:392
				_go_fuzz_dep_.CoverTab[18620]++
//line /usr/local/go/src/crypto/x509/parser.go:392
				// _ = "end of CoverTab[18620]"
//line /usr/local/go/src/crypto/x509/parser.go:392
			}
//line /usr/local/go/src/crypto/x509/parser.go:392
			// _ = "end of CoverTab[18609]"
//line /usr/local/go/src/crypto/x509/parser.go:392
			_go_fuzz_dep_.CoverTab[18610]++
									uri, err := url.Parse(uriStr)
									if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:394
				_go_fuzz_dep_.CoverTab[18621]++
										return fmt.Errorf("x509: cannot parse URI %q: %s", uriStr, err)
//line /usr/local/go/src/crypto/x509/parser.go:395
				// _ = "end of CoverTab[18621]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:396
				_go_fuzz_dep_.CoverTab[18622]++
//line /usr/local/go/src/crypto/x509/parser.go:396
				// _ = "end of CoverTab[18622]"
//line /usr/local/go/src/crypto/x509/parser.go:396
			}
//line /usr/local/go/src/crypto/x509/parser.go:396
			// _ = "end of CoverTab[18610]"
//line /usr/local/go/src/crypto/x509/parser.go:396
			_go_fuzz_dep_.CoverTab[18611]++
									if len(uri.Host) > 0 {
//line /usr/local/go/src/crypto/x509/parser.go:397
				_go_fuzz_dep_.CoverTab[18623]++
										if _, ok := domainToReverseLabels(uri.Host); !ok {
//line /usr/local/go/src/crypto/x509/parser.go:398
					_go_fuzz_dep_.CoverTab[18624]++
											return fmt.Errorf("x509: cannot parse URI %q: invalid domain", uriStr)
//line /usr/local/go/src/crypto/x509/parser.go:399
					// _ = "end of CoverTab[18624]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:400
					_go_fuzz_dep_.CoverTab[18625]++
//line /usr/local/go/src/crypto/x509/parser.go:400
					// _ = "end of CoverTab[18625]"
//line /usr/local/go/src/crypto/x509/parser.go:400
				}
//line /usr/local/go/src/crypto/x509/parser.go:400
				// _ = "end of CoverTab[18623]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:401
				_go_fuzz_dep_.CoverTab[18626]++
//line /usr/local/go/src/crypto/x509/parser.go:401
				// _ = "end of CoverTab[18626]"
//line /usr/local/go/src/crypto/x509/parser.go:401
			}
//line /usr/local/go/src/crypto/x509/parser.go:401
			// _ = "end of CoverTab[18611]"
//line /usr/local/go/src/crypto/x509/parser.go:401
			_go_fuzz_dep_.CoverTab[18612]++
									uris = append(uris, uri)
//line /usr/local/go/src/crypto/x509/parser.go:402
			// _ = "end of CoverTab[18612]"
		case nameTypeIP:
//line /usr/local/go/src/crypto/x509/parser.go:403
			_go_fuzz_dep_.CoverTab[18613]++
									switch len(data) {
			case net.IPv4len, net.IPv6len:
//line /usr/local/go/src/crypto/x509/parser.go:405
				_go_fuzz_dep_.CoverTab[18627]++
										ipAddresses = append(ipAddresses, data)
//line /usr/local/go/src/crypto/x509/parser.go:406
				// _ = "end of CoverTab[18627]"
			default:
//line /usr/local/go/src/crypto/x509/parser.go:407
				_go_fuzz_dep_.CoverTab[18628]++
										return errors.New("x509: cannot parse IP address of length " + strconv.Itoa(len(data)))
//line /usr/local/go/src/crypto/x509/parser.go:408
				// _ = "end of CoverTab[18628]"
			}
//line /usr/local/go/src/crypto/x509/parser.go:409
			// _ = "end of CoverTab[18613]"
//line /usr/local/go/src/crypto/x509/parser.go:409
		default:
//line /usr/local/go/src/crypto/x509/parser.go:409
			_go_fuzz_dep_.CoverTab[18614]++
//line /usr/local/go/src/crypto/x509/parser.go:409
			// _ = "end of CoverTab[18614]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:410
		// _ = "end of CoverTab[18603]"
//line /usr/local/go/src/crypto/x509/parser.go:410
		_go_fuzz_dep_.CoverTab[18604]++

								return nil
//line /usr/local/go/src/crypto/x509/parser.go:412
		// _ = "end of CoverTab[18604]"
	})
//line /usr/local/go/src/crypto/x509/parser.go:413
	// _ = "end of CoverTab[18601]"
//line /usr/local/go/src/crypto/x509/parser.go:413
	_go_fuzz_dep_.CoverTab[18602]++

							return
//line /usr/local/go/src/crypto/x509/parser.go:415
	// _ = "end of CoverTab[18602]"
}

func parseExtKeyUsageExtension(der cryptobyte.String) ([]ExtKeyUsage, []asn1.ObjectIdentifier, error) {
//line /usr/local/go/src/crypto/x509/parser.go:418
	_go_fuzz_dep_.CoverTab[18629]++
							var extKeyUsages []ExtKeyUsage
							var unknownUsages []asn1.ObjectIdentifier
							if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:421
		_go_fuzz_dep_.CoverTab[18632]++
								return nil, nil, errors.New("x509: invalid extended key usages")
//line /usr/local/go/src/crypto/x509/parser.go:422
		// _ = "end of CoverTab[18632]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:423
		_go_fuzz_dep_.CoverTab[18633]++
//line /usr/local/go/src/crypto/x509/parser.go:423
		// _ = "end of CoverTab[18633]"
//line /usr/local/go/src/crypto/x509/parser.go:423
	}
//line /usr/local/go/src/crypto/x509/parser.go:423
	// _ = "end of CoverTab[18629]"
//line /usr/local/go/src/crypto/x509/parser.go:423
	_go_fuzz_dep_.CoverTab[18630]++
							for !der.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:424
		_go_fuzz_dep_.CoverTab[18634]++
								var eku asn1.ObjectIdentifier
								if !der.ReadASN1ObjectIdentifier(&eku) {
//line /usr/local/go/src/crypto/x509/parser.go:426
			_go_fuzz_dep_.CoverTab[18636]++
									return nil, nil, errors.New("x509: invalid extended key usages")
//line /usr/local/go/src/crypto/x509/parser.go:427
			// _ = "end of CoverTab[18636]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:428
			_go_fuzz_dep_.CoverTab[18637]++
//line /usr/local/go/src/crypto/x509/parser.go:428
			// _ = "end of CoverTab[18637]"
//line /usr/local/go/src/crypto/x509/parser.go:428
		}
//line /usr/local/go/src/crypto/x509/parser.go:428
		// _ = "end of CoverTab[18634]"
//line /usr/local/go/src/crypto/x509/parser.go:428
		_go_fuzz_dep_.CoverTab[18635]++
								if extKeyUsage, ok := extKeyUsageFromOID(eku); ok {
//line /usr/local/go/src/crypto/x509/parser.go:429
			_go_fuzz_dep_.CoverTab[18638]++
									extKeyUsages = append(extKeyUsages, extKeyUsage)
//line /usr/local/go/src/crypto/x509/parser.go:430
			// _ = "end of CoverTab[18638]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:431
			_go_fuzz_dep_.CoverTab[18639]++
									unknownUsages = append(unknownUsages, eku)
//line /usr/local/go/src/crypto/x509/parser.go:432
			// _ = "end of CoverTab[18639]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:433
		// _ = "end of CoverTab[18635]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:434
	// _ = "end of CoverTab[18630]"
//line /usr/local/go/src/crypto/x509/parser.go:434
	_go_fuzz_dep_.CoverTab[18631]++
							return extKeyUsages, unknownUsages, nil
//line /usr/local/go/src/crypto/x509/parser.go:435
	// _ = "end of CoverTab[18631]"
}

func parseCertificatePoliciesExtension(der cryptobyte.String) ([]asn1.ObjectIdentifier, error) {
//line /usr/local/go/src/crypto/x509/parser.go:438
	_go_fuzz_dep_.CoverTab[18640]++
							var oids []asn1.ObjectIdentifier
							if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:440
		_go_fuzz_dep_.CoverTab[18643]++
								return nil, errors.New("x509: invalid certificate policies")
//line /usr/local/go/src/crypto/x509/parser.go:441
		// _ = "end of CoverTab[18643]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:442
		_go_fuzz_dep_.CoverTab[18644]++
//line /usr/local/go/src/crypto/x509/parser.go:442
		// _ = "end of CoverTab[18644]"
//line /usr/local/go/src/crypto/x509/parser.go:442
	}
//line /usr/local/go/src/crypto/x509/parser.go:442
	// _ = "end of CoverTab[18640]"
//line /usr/local/go/src/crypto/x509/parser.go:442
	_go_fuzz_dep_.CoverTab[18641]++
							for !der.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:443
		_go_fuzz_dep_.CoverTab[18645]++
								var cp cryptobyte.String
								if !der.ReadASN1(&cp, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:445
			_go_fuzz_dep_.CoverTab[18648]++
									return nil, errors.New("x509: invalid certificate policies")
//line /usr/local/go/src/crypto/x509/parser.go:446
			// _ = "end of CoverTab[18648]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:447
			_go_fuzz_dep_.CoverTab[18649]++
//line /usr/local/go/src/crypto/x509/parser.go:447
			// _ = "end of CoverTab[18649]"
//line /usr/local/go/src/crypto/x509/parser.go:447
		}
//line /usr/local/go/src/crypto/x509/parser.go:447
		// _ = "end of CoverTab[18645]"
//line /usr/local/go/src/crypto/x509/parser.go:447
		_go_fuzz_dep_.CoverTab[18646]++
								var oid asn1.ObjectIdentifier
								if !cp.ReadASN1ObjectIdentifier(&oid) {
//line /usr/local/go/src/crypto/x509/parser.go:449
			_go_fuzz_dep_.CoverTab[18650]++
									return nil, errors.New("x509: invalid certificate policies")
//line /usr/local/go/src/crypto/x509/parser.go:450
			// _ = "end of CoverTab[18650]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:451
			_go_fuzz_dep_.CoverTab[18651]++
//line /usr/local/go/src/crypto/x509/parser.go:451
			// _ = "end of CoverTab[18651]"
//line /usr/local/go/src/crypto/x509/parser.go:451
		}
//line /usr/local/go/src/crypto/x509/parser.go:451
		// _ = "end of CoverTab[18646]"
//line /usr/local/go/src/crypto/x509/parser.go:451
		_go_fuzz_dep_.CoverTab[18647]++
								oids = append(oids, oid)
//line /usr/local/go/src/crypto/x509/parser.go:452
		// _ = "end of CoverTab[18647]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:453
	// _ = "end of CoverTab[18641]"
//line /usr/local/go/src/crypto/x509/parser.go:453
	_go_fuzz_dep_.CoverTab[18642]++

							return oids, nil
//line /usr/local/go/src/crypto/x509/parser.go:455
	// _ = "end of CoverTab[18642]"
}

// isValidIPMask reports whether mask consists of zero or more 1 bits, followed by zero bits.
func isValidIPMask(mask []byte) bool {
//line /usr/local/go/src/crypto/x509/parser.go:459
	_go_fuzz_dep_.CoverTab[18652]++
							seenZero := false

							for _, b := range mask {
//line /usr/local/go/src/crypto/x509/parser.go:462
		_go_fuzz_dep_.CoverTab[18654]++
								if seenZero {
//line /usr/local/go/src/crypto/x509/parser.go:463
			_go_fuzz_dep_.CoverTab[18656]++
									if b != 0 {
//line /usr/local/go/src/crypto/x509/parser.go:464
				_go_fuzz_dep_.CoverTab[18658]++
										return false
//line /usr/local/go/src/crypto/x509/parser.go:465
				// _ = "end of CoverTab[18658]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:466
				_go_fuzz_dep_.CoverTab[18659]++
//line /usr/local/go/src/crypto/x509/parser.go:466
				// _ = "end of CoverTab[18659]"
//line /usr/local/go/src/crypto/x509/parser.go:466
			}
//line /usr/local/go/src/crypto/x509/parser.go:466
			// _ = "end of CoverTab[18656]"
//line /usr/local/go/src/crypto/x509/parser.go:466
			_go_fuzz_dep_.CoverTab[18657]++

									continue
//line /usr/local/go/src/crypto/x509/parser.go:468
			// _ = "end of CoverTab[18657]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:469
			_go_fuzz_dep_.CoverTab[18660]++
//line /usr/local/go/src/crypto/x509/parser.go:469
			// _ = "end of CoverTab[18660]"
//line /usr/local/go/src/crypto/x509/parser.go:469
		}
//line /usr/local/go/src/crypto/x509/parser.go:469
		// _ = "end of CoverTab[18654]"
//line /usr/local/go/src/crypto/x509/parser.go:469
		_go_fuzz_dep_.CoverTab[18655]++

								switch b {
		case 0x00, 0x80, 0xc0, 0xe0, 0xf0, 0xf8, 0xfc, 0xfe:
//line /usr/local/go/src/crypto/x509/parser.go:472
			_go_fuzz_dep_.CoverTab[18661]++
									seenZero = true
//line /usr/local/go/src/crypto/x509/parser.go:473
			// _ = "end of CoverTab[18661]"
		case 0xff:
//line /usr/local/go/src/crypto/x509/parser.go:474
			_go_fuzz_dep_.CoverTab[18662]++
//line /usr/local/go/src/crypto/x509/parser.go:474
			// _ = "end of CoverTab[18662]"
		default:
//line /usr/local/go/src/crypto/x509/parser.go:475
			_go_fuzz_dep_.CoverTab[18663]++
									return false
//line /usr/local/go/src/crypto/x509/parser.go:476
			// _ = "end of CoverTab[18663]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:477
		// _ = "end of CoverTab[18655]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:478
	// _ = "end of CoverTab[18652]"
//line /usr/local/go/src/crypto/x509/parser.go:478
	_go_fuzz_dep_.CoverTab[18653]++

							return true
//line /usr/local/go/src/crypto/x509/parser.go:480
	// _ = "end of CoverTab[18653]"
}

func parseNameConstraintsExtension(out *Certificate, e pkix.Extension) (unhandled bool, err error) {
//line /usr/local/go/src/crypto/x509/parser.go:483
	_go_fuzz_dep_.CoverTab[18664]++

//line /usr/local/go/src/crypto/x509/parser.go:499
	outer := cryptobyte.String(e.Value)
	var toplevel, permitted, excluded cryptobyte.String
	var havePermitted, haveExcluded bool
	if !outer.ReadASN1(&toplevel, cryptobyte_asn1.SEQUENCE) || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:502
		_go_fuzz_dep_.CoverTab[18670]++
//line /usr/local/go/src/crypto/x509/parser.go:502
		return !outer.Empty()
								// _ = "end of CoverTab[18670]"
//line /usr/local/go/src/crypto/x509/parser.go:503
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:503
		_go_fuzz_dep_.CoverTab[18671]++
//line /usr/local/go/src/crypto/x509/parser.go:503
		return !toplevel.ReadOptionalASN1(&permitted, &havePermitted, cryptobyte_asn1.Tag(0).ContextSpecific().Constructed())
								// _ = "end of CoverTab[18671]"
//line /usr/local/go/src/crypto/x509/parser.go:504
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:504
		_go_fuzz_dep_.CoverTab[18672]++
//line /usr/local/go/src/crypto/x509/parser.go:504
		return !toplevel.ReadOptionalASN1(&excluded, &haveExcluded, cryptobyte_asn1.Tag(1).ContextSpecific().Constructed())
								// _ = "end of CoverTab[18672]"
//line /usr/local/go/src/crypto/x509/parser.go:505
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:505
		_go_fuzz_dep_.CoverTab[18673]++
//line /usr/local/go/src/crypto/x509/parser.go:505
		return !toplevel.Empty()
								// _ = "end of CoverTab[18673]"
//line /usr/local/go/src/crypto/x509/parser.go:506
	}() {
//line /usr/local/go/src/crypto/x509/parser.go:506
		_go_fuzz_dep_.CoverTab[18674]++
								return false, errors.New("x509: invalid NameConstraints extension")
//line /usr/local/go/src/crypto/x509/parser.go:507
		// _ = "end of CoverTab[18674]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:508
		_go_fuzz_dep_.CoverTab[18675]++
//line /usr/local/go/src/crypto/x509/parser.go:508
		// _ = "end of CoverTab[18675]"
//line /usr/local/go/src/crypto/x509/parser.go:508
	}
//line /usr/local/go/src/crypto/x509/parser.go:508
	// _ = "end of CoverTab[18664]"
//line /usr/local/go/src/crypto/x509/parser.go:508
	_go_fuzz_dep_.CoverTab[18665]++

							if !havePermitted && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:510
		_go_fuzz_dep_.CoverTab[18676]++
//line /usr/local/go/src/crypto/x509/parser.go:510
		return !haveExcluded
//line /usr/local/go/src/crypto/x509/parser.go:510
		// _ = "end of CoverTab[18676]"
//line /usr/local/go/src/crypto/x509/parser.go:510
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:510
		_go_fuzz_dep_.CoverTab[18677]++
//line /usr/local/go/src/crypto/x509/parser.go:510
		return len(permitted) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:510
			_go_fuzz_dep_.CoverTab[18678]++
//line /usr/local/go/src/crypto/x509/parser.go:510
			return len(excluded) == 0
//line /usr/local/go/src/crypto/x509/parser.go:510
			// _ = "end of CoverTab[18678]"
//line /usr/local/go/src/crypto/x509/parser.go:510
		}()
//line /usr/local/go/src/crypto/x509/parser.go:510
		// _ = "end of CoverTab[18677]"
//line /usr/local/go/src/crypto/x509/parser.go:510
	}() {
//line /usr/local/go/src/crypto/x509/parser.go:510
		_go_fuzz_dep_.CoverTab[18679]++

//line /usr/local/go/src/crypto/x509/parser.go:515
		return false, errors.New("x509: empty name constraints extension")
//line /usr/local/go/src/crypto/x509/parser.go:515
		// _ = "end of CoverTab[18679]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:516
		_go_fuzz_dep_.CoverTab[18680]++
//line /usr/local/go/src/crypto/x509/parser.go:516
		// _ = "end of CoverTab[18680]"
//line /usr/local/go/src/crypto/x509/parser.go:516
	}
//line /usr/local/go/src/crypto/x509/parser.go:516
	// _ = "end of CoverTab[18665]"
//line /usr/local/go/src/crypto/x509/parser.go:516
	_go_fuzz_dep_.CoverTab[18666]++

							getValues := func(subtrees cryptobyte.String) (dnsNames []string, ips []*net.IPNet, emails, uriDomains []string, err error) {
//line /usr/local/go/src/crypto/x509/parser.go:518
		_go_fuzz_dep_.CoverTab[18681]++
								for !subtrees.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:519
			_go_fuzz_dep_.CoverTab[18683]++
									var seq, value cryptobyte.String
									var tag cryptobyte_asn1.Tag
									if !subtrees.ReadASN1(&seq, cryptobyte_asn1.SEQUENCE) || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:522
				_go_fuzz_dep_.CoverTab[18685]++
//line /usr/local/go/src/crypto/x509/parser.go:522
				return !seq.ReadAnyASN1(&value, &tag)
										// _ = "end of CoverTab[18685]"
//line /usr/local/go/src/crypto/x509/parser.go:523
			}() {
//line /usr/local/go/src/crypto/x509/parser.go:523
				_go_fuzz_dep_.CoverTab[18686]++
										return nil, nil, nil, nil, fmt.Errorf("x509: invalid NameConstraints extension")
//line /usr/local/go/src/crypto/x509/parser.go:524
				// _ = "end of CoverTab[18686]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:525
				_go_fuzz_dep_.CoverTab[18687]++
//line /usr/local/go/src/crypto/x509/parser.go:525
				// _ = "end of CoverTab[18687]"
//line /usr/local/go/src/crypto/x509/parser.go:525
			}
//line /usr/local/go/src/crypto/x509/parser.go:525
			// _ = "end of CoverTab[18683]"
//line /usr/local/go/src/crypto/x509/parser.go:525
			_go_fuzz_dep_.CoverTab[18684]++

									var (
				dnsTag		= cryptobyte_asn1.Tag(2).ContextSpecific()
				emailTag	= cryptobyte_asn1.Tag(1).ContextSpecific()
				ipTag		= cryptobyte_asn1.Tag(7).ContextSpecific()
				uriTag		= cryptobyte_asn1.Tag(6).ContextSpecific()
			)

			switch tag {
			case dnsTag:
//line /usr/local/go/src/crypto/x509/parser.go:535
				_go_fuzz_dep_.CoverTab[18688]++
										domain := string(value)
										if err := isIA5String(domain); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:537
					_go_fuzz_dep_.CoverTab[18704]++
											return nil, nil, nil, nil, errors.New("x509: invalid constraint value: " + err.Error())
//line /usr/local/go/src/crypto/x509/parser.go:538
					// _ = "end of CoverTab[18704]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:539
					_go_fuzz_dep_.CoverTab[18705]++
//line /usr/local/go/src/crypto/x509/parser.go:539
					// _ = "end of CoverTab[18705]"
//line /usr/local/go/src/crypto/x509/parser.go:539
				}
//line /usr/local/go/src/crypto/x509/parser.go:539
				// _ = "end of CoverTab[18688]"
//line /usr/local/go/src/crypto/x509/parser.go:539
				_go_fuzz_dep_.CoverTab[18689]++

										trimmedDomain := domain
										if len(trimmedDomain) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:542
					_go_fuzz_dep_.CoverTab[18706]++
//line /usr/local/go/src/crypto/x509/parser.go:542
					return trimmedDomain[0] == '.'
//line /usr/local/go/src/crypto/x509/parser.go:542
					// _ = "end of CoverTab[18706]"
//line /usr/local/go/src/crypto/x509/parser.go:542
				}() {
//line /usr/local/go/src/crypto/x509/parser.go:542
					_go_fuzz_dep_.CoverTab[18707]++

//line /usr/local/go/src/crypto/x509/parser.go:547
					trimmedDomain = trimmedDomain[1:]
//line /usr/local/go/src/crypto/x509/parser.go:547
					// _ = "end of CoverTab[18707]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:548
					_go_fuzz_dep_.CoverTab[18708]++
//line /usr/local/go/src/crypto/x509/parser.go:548
					// _ = "end of CoverTab[18708]"
//line /usr/local/go/src/crypto/x509/parser.go:548
				}
//line /usr/local/go/src/crypto/x509/parser.go:548
				// _ = "end of CoverTab[18689]"
//line /usr/local/go/src/crypto/x509/parser.go:548
				_go_fuzz_dep_.CoverTab[18690]++
										if _, ok := domainToReverseLabels(trimmedDomain); !ok {
//line /usr/local/go/src/crypto/x509/parser.go:549
					_go_fuzz_dep_.CoverTab[18709]++
											return nil, nil, nil, nil, fmt.Errorf("x509: failed to parse dnsName constraint %q", domain)
//line /usr/local/go/src/crypto/x509/parser.go:550
					// _ = "end of CoverTab[18709]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:551
					_go_fuzz_dep_.CoverTab[18710]++
//line /usr/local/go/src/crypto/x509/parser.go:551
					// _ = "end of CoverTab[18710]"
//line /usr/local/go/src/crypto/x509/parser.go:551
				}
//line /usr/local/go/src/crypto/x509/parser.go:551
				// _ = "end of CoverTab[18690]"
//line /usr/local/go/src/crypto/x509/parser.go:551
				_go_fuzz_dep_.CoverTab[18691]++
										dnsNames = append(dnsNames, domain)
//line /usr/local/go/src/crypto/x509/parser.go:552
				// _ = "end of CoverTab[18691]"

			case ipTag:
//line /usr/local/go/src/crypto/x509/parser.go:554
				_go_fuzz_dep_.CoverTab[18692]++
										l := len(value)
										var ip, mask []byte

										switch l {
				case 8:
//line /usr/local/go/src/crypto/x509/parser.go:559
					_go_fuzz_dep_.CoverTab[18711]++
											ip = value[:4]
											mask = value[4:]
//line /usr/local/go/src/crypto/x509/parser.go:561
					// _ = "end of CoverTab[18711]"

				case 32:
//line /usr/local/go/src/crypto/x509/parser.go:563
					_go_fuzz_dep_.CoverTab[18712]++
											ip = value[:16]
											mask = value[16:]
//line /usr/local/go/src/crypto/x509/parser.go:565
					// _ = "end of CoverTab[18712]"

				default:
//line /usr/local/go/src/crypto/x509/parser.go:567
					_go_fuzz_dep_.CoverTab[18713]++
											return nil, nil, nil, nil, fmt.Errorf("x509: IP constraint contained value of length %d", l)
//line /usr/local/go/src/crypto/x509/parser.go:568
					// _ = "end of CoverTab[18713]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:569
				// _ = "end of CoverTab[18692]"
//line /usr/local/go/src/crypto/x509/parser.go:569
				_go_fuzz_dep_.CoverTab[18693]++

										if !isValidIPMask(mask) {
//line /usr/local/go/src/crypto/x509/parser.go:571
					_go_fuzz_dep_.CoverTab[18714]++
											return nil, nil, nil, nil, fmt.Errorf("x509: IP constraint contained invalid mask %x", mask)
//line /usr/local/go/src/crypto/x509/parser.go:572
					// _ = "end of CoverTab[18714]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:573
					_go_fuzz_dep_.CoverTab[18715]++
//line /usr/local/go/src/crypto/x509/parser.go:573
					// _ = "end of CoverTab[18715]"
//line /usr/local/go/src/crypto/x509/parser.go:573
				}
//line /usr/local/go/src/crypto/x509/parser.go:573
				// _ = "end of CoverTab[18693]"
//line /usr/local/go/src/crypto/x509/parser.go:573
				_go_fuzz_dep_.CoverTab[18694]++

										ips = append(ips, &net.IPNet{IP: net.IP(ip), Mask: net.IPMask(mask)})
//line /usr/local/go/src/crypto/x509/parser.go:575
				// _ = "end of CoverTab[18694]"

			case emailTag:
//line /usr/local/go/src/crypto/x509/parser.go:577
				_go_fuzz_dep_.CoverTab[18695]++
										constraint := string(value)
										if err := isIA5String(constraint); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:579
					_go_fuzz_dep_.CoverTab[18716]++
											return nil, nil, nil, nil, errors.New("x509: invalid constraint value: " + err.Error())
//line /usr/local/go/src/crypto/x509/parser.go:580
					// _ = "end of CoverTab[18716]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:581
					_go_fuzz_dep_.CoverTab[18717]++
//line /usr/local/go/src/crypto/x509/parser.go:581
					// _ = "end of CoverTab[18717]"
//line /usr/local/go/src/crypto/x509/parser.go:581
				}
//line /usr/local/go/src/crypto/x509/parser.go:581
				// _ = "end of CoverTab[18695]"
//line /usr/local/go/src/crypto/x509/parser.go:581
				_go_fuzz_dep_.CoverTab[18696]++

//line /usr/local/go/src/crypto/x509/parser.go:585
				if strings.Contains(constraint, "@") {
//line /usr/local/go/src/crypto/x509/parser.go:585
					_go_fuzz_dep_.CoverTab[18718]++
											if _, ok := parseRFC2821Mailbox(constraint); !ok {
//line /usr/local/go/src/crypto/x509/parser.go:586
						_go_fuzz_dep_.CoverTab[18719]++
												return nil, nil, nil, nil, fmt.Errorf("x509: failed to parse rfc822Name constraint %q", constraint)
//line /usr/local/go/src/crypto/x509/parser.go:587
						// _ = "end of CoverTab[18719]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:588
						_go_fuzz_dep_.CoverTab[18720]++
//line /usr/local/go/src/crypto/x509/parser.go:588
						// _ = "end of CoverTab[18720]"
//line /usr/local/go/src/crypto/x509/parser.go:588
					}
//line /usr/local/go/src/crypto/x509/parser.go:588
					// _ = "end of CoverTab[18718]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:589
					_go_fuzz_dep_.CoverTab[18721]++

											domain := constraint
											if len(domain) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:592
						_go_fuzz_dep_.CoverTab[18723]++
//line /usr/local/go/src/crypto/x509/parser.go:592
						return domain[0] == '.'
//line /usr/local/go/src/crypto/x509/parser.go:592
						// _ = "end of CoverTab[18723]"
//line /usr/local/go/src/crypto/x509/parser.go:592
					}() {
//line /usr/local/go/src/crypto/x509/parser.go:592
						_go_fuzz_dep_.CoverTab[18724]++
												domain = domain[1:]
//line /usr/local/go/src/crypto/x509/parser.go:593
						// _ = "end of CoverTab[18724]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:594
						_go_fuzz_dep_.CoverTab[18725]++
//line /usr/local/go/src/crypto/x509/parser.go:594
						// _ = "end of CoverTab[18725]"
//line /usr/local/go/src/crypto/x509/parser.go:594
					}
//line /usr/local/go/src/crypto/x509/parser.go:594
					// _ = "end of CoverTab[18721]"
//line /usr/local/go/src/crypto/x509/parser.go:594
					_go_fuzz_dep_.CoverTab[18722]++
											if _, ok := domainToReverseLabels(domain); !ok {
//line /usr/local/go/src/crypto/x509/parser.go:595
						_go_fuzz_dep_.CoverTab[18726]++
												return nil, nil, nil, nil, fmt.Errorf("x509: failed to parse rfc822Name constraint %q", constraint)
//line /usr/local/go/src/crypto/x509/parser.go:596
						// _ = "end of CoverTab[18726]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:597
						_go_fuzz_dep_.CoverTab[18727]++
//line /usr/local/go/src/crypto/x509/parser.go:597
						// _ = "end of CoverTab[18727]"
//line /usr/local/go/src/crypto/x509/parser.go:597
					}
//line /usr/local/go/src/crypto/x509/parser.go:597
					// _ = "end of CoverTab[18722]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:598
				// _ = "end of CoverTab[18696]"
//line /usr/local/go/src/crypto/x509/parser.go:598
				_go_fuzz_dep_.CoverTab[18697]++
										emails = append(emails, constraint)
//line /usr/local/go/src/crypto/x509/parser.go:599
				// _ = "end of CoverTab[18697]"

			case uriTag:
//line /usr/local/go/src/crypto/x509/parser.go:601
				_go_fuzz_dep_.CoverTab[18698]++
										domain := string(value)
										if err := isIA5String(domain); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:603
					_go_fuzz_dep_.CoverTab[18728]++
											return nil, nil, nil, nil, errors.New("x509: invalid constraint value: " + err.Error())
//line /usr/local/go/src/crypto/x509/parser.go:604
					// _ = "end of CoverTab[18728]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:605
					_go_fuzz_dep_.CoverTab[18729]++
//line /usr/local/go/src/crypto/x509/parser.go:605
					// _ = "end of CoverTab[18729]"
//line /usr/local/go/src/crypto/x509/parser.go:605
				}
//line /usr/local/go/src/crypto/x509/parser.go:605
				// _ = "end of CoverTab[18698]"
//line /usr/local/go/src/crypto/x509/parser.go:605
				_go_fuzz_dep_.CoverTab[18699]++

										if net.ParseIP(domain) != nil {
//line /usr/local/go/src/crypto/x509/parser.go:607
					_go_fuzz_dep_.CoverTab[18730]++
											return nil, nil, nil, nil, fmt.Errorf("x509: failed to parse URI constraint %q: cannot be IP address", domain)
//line /usr/local/go/src/crypto/x509/parser.go:608
					// _ = "end of CoverTab[18730]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:609
					_go_fuzz_dep_.CoverTab[18731]++
//line /usr/local/go/src/crypto/x509/parser.go:609
					// _ = "end of CoverTab[18731]"
//line /usr/local/go/src/crypto/x509/parser.go:609
				}
//line /usr/local/go/src/crypto/x509/parser.go:609
				// _ = "end of CoverTab[18699]"
//line /usr/local/go/src/crypto/x509/parser.go:609
				_go_fuzz_dep_.CoverTab[18700]++

										trimmedDomain := domain
										if len(trimmedDomain) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:612
					_go_fuzz_dep_.CoverTab[18732]++
//line /usr/local/go/src/crypto/x509/parser.go:612
					return trimmedDomain[0] == '.'
//line /usr/local/go/src/crypto/x509/parser.go:612
					// _ = "end of CoverTab[18732]"
//line /usr/local/go/src/crypto/x509/parser.go:612
				}() {
//line /usr/local/go/src/crypto/x509/parser.go:612
					_go_fuzz_dep_.CoverTab[18733]++

//line /usr/local/go/src/crypto/x509/parser.go:617
					trimmedDomain = trimmedDomain[1:]
//line /usr/local/go/src/crypto/x509/parser.go:617
					// _ = "end of CoverTab[18733]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:618
					_go_fuzz_dep_.CoverTab[18734]++
//line /usr/local/go/src/crypto/x509/parser.go:618
					// _ = "end of CoverTab[18734]"
//line /usr/local/go/src/crypto/x509/parser.go:618
				}
//line /usr/local/go/src/crypto/x509/parser.go:618
				// _ = "end of CoverTab[18700]"
//line /usr/local/go/src/crypto/x509/parser.go:618
				_go_fuzz_dep_.CoverTab[18701]++
										if _, ok := domainToReverseLabels(trimmedDomain); !ok {
//line /usr/local/go/src/crypto/x509/parser.go:619
					_go_fuzz_dep_.CoverTab[18735]++
											return nil, nil, nil, nil, fmt.Errorf("x509: failed to parse URI constraint %q", domain)
//line /usr/local/go/src/crypto/x509/parser.go:620
					// _ = "end of CoverTab[18735]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:621
					_go_fuzz_dep_.CoverTab[18736]++
//line /usr/local/go/src/crypto/x509/parser.go:621
					// _ = "end of CoverTab[18736]"
//line /usr/local/go/src/crypto/x509/parser.go:621
				}
//line /usr/local/go/src/crypto/x509/parser.go:621
				// _ = "end of CoverTab[18701]"
//line /usr/local/go/src/crypto/x509/parser.go:621
				_go_fuzz_dep_.CoverTab[18702]++
										uriDomains = append(uriDomains, domain)
//line /usr/local/go/src/crypto/x509/parser.go:622
				// _ = "end of CoverTab[18702]"

			default:
//line /usr/local/go/src/crypto/x509/parser.go:624
				_go_fuzz_dep_.CoverTab[18703]++
										unhandled = true
//line /usr/local/go/src/crypto/x509/parser.go:625
				// _ = "end of CoverTab[18703]"
			}
//line /usr/local/go/src/crypto/x509/parser.go:626
			// _ = "end of CoverTab[18684]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:627
		// _ = "end of CoverTab[18681]"
//line /usr/local/go/src/crypto/x509/parser.go:627
		_go_fuzz_dep_.CoverTab[18682]++

								return dnsNames, ips, emails, uriDomains, nil
//line /usr/local/go/src/crypto/x509/parser.go:629
		// _ = "end of CoverTab[18682]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:630
	// _ = "end of CoverTab[18666]"
//line /usr/local/go/src/crypto/x509/parser.go:630
	_go_fuzz_dep_.CoverTab[18667]++

							if out.PermittedDNSDomains, out.PermittedIPRanges, out.PermittedEmailAddresses, out.PermittedURIDomains, err = getValues(permitted); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:632
		_go_fuzz_dep_.CoverTab[18737]++
								return false, err
//line /usr/local/go/src/crypto/x509/parser.go:633
		// _ = "end of CoverTab[18737]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:634
		_go_fuzz_dep_.CoverTab[18738]++
//line /usr/local/go/src/crypto/x509/parser.go:634
		// _ = "end of CoverTab[18738]"
//line /usr/local/go/src/crypto/x509/parser.go:634
	}
//line /usr/local/go/src/crypto/x509/parser.go:634
	// _ = "end of CoverTab[18667]"
//line /usr/local/go/src/crypto/x509/parser.go:634
	_go_fuzz_dep_.CoverTab[18668]++
							if out.ExcludedDNSDomains, out.ExcludedIPRanges, out.ExcludedEmailAddresses, out.ExcludedURIDomains, err = getValues(excluded); err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:635
		_go_fuzz_dep_.CoverTab[18739]++
								return false, err
//line /usr/local/go/src/crypto/x509/parser.go:636
		// _ = "end of CoverTab[18739]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:637
		_go_fuzz_dep_.CoverTab[18740]++
//line /usr/local/go/src/crypto/x509/parser.go:637
		// _ = "end of CoverTab[18740]"
//line /usr/local/go/src/crypto/x509/parser.go:637
	}
//line /usr/local/go/src/crypto/x509/parser.go:637
	// _ = "end of CoverTab[18668]"
//line /usr/local/go/src/crypto/x509/parser.go:637
	_go_fuzz_dep_.CoverTab[18669]++
							out.PermittedDNSDomainsCritical = e.Critical

							return unhandled, nil
//line /usr/local/go/src/crypto/x509/parser.go:640
	// _ = "end of CoverTab[18669]"
}

func processExtensions(out *Certificate) error {
//line /usr/local/go/src/crypto/x509/parser.go:643
	_go_fuzz_dep_.CoverTab[18741]++
							var err error
							for _, e := range out.Extensions {
//line /usr/local/go/src/crypto/x509/parser.go:645
		_go_fuzz_dep_.CoverTab[18743]++
								unhandled := false

								if len(e.Id) == 4 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:648
			_go_fuzz_dep_.CoverTab[18745]++
//line /usr/local/go/src/crypto/x509/parser.go:648
			return e.Id[0] == 2
//line /usr/local/go/src/crypto/x509/parser.go:648
			// _ = "end of CoverTab[18745]"
//line /usr/local/go/src/crypto/x509/parser.go:648
		}() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:648
			_go_fuzz_dep_.CoverTab[18746]++
//line /usr/local/go/src/crypto/x509/parser.go:648
			return e.Id[1] == 5
//line /usr/local/go/src/crypto/x509/parser.go:648
			// _ = "end of CoverTab[18746]"
//line /usr/local/go/src/crypto/x509/parser.go:648
		}() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:648
			_go_fuzz_dep_.CoverTab[18747]++
//line /usr/local/go/src/crypto/x509/parser.go:648
			return e.Id[2] == 29
//line /usr/local/go/src/crypto/x509/parser.go:648
			// _ = "end of CoverTab[18747]"
//line /usr/local/go/src/crypto/x509/parser.go:648
		}() {
//line /usr/local/go/src/crypto/x509/parser.go:648
			_go_fuzz_dep_.CoverTab[18748]++
									switch e.Id[3] {
			case 15:
//line /usr/local/go/src/crypto/x509/parser.go:650
				_go_fuzz_dep_.CoverTab[18749]++
										out.KeyUsage, err = parseKeyUsageExtension(e.Value)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:652
					_go_fuzz_dep_.CoverTab[18764]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:653
					// _ = "end of CoverTab[18764]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:654
					_go_fuzz_dep_.CoverTab[18765]++
//line /usr/local/go/src/crypto/x509/parser.go:654
					// _ = "end of CoverTab[18765]"
//line /usr/local/go/src/crypto/x509/parser.go:654
				}
//line /usr/local/go/src/crypto/x509/parser.go:654
				// _ = "end of CoverTab[18749]"
			case 19:
//line /usr/local/go/src/crypto/x509/parser.go:655
				_go_fuzz_dep_.CoverTab[18750]++
										out.IsCA, out.MaxPathLen, err = parseBasicConstraintsExtension(e.Value)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:657
					_go_fuzz_dep_.CoverTab[18766]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:658
					// _ = "end of CoverTab[18766]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:659
					_go_fuzz_dep_.CoverTab[18767]++
//line /usr/local/go/src/crypto/x509/parser.go:659
					// _ = "end of CoverTab[18767]"
//line /usr/local/go/src/crypto/x509/parser.go:659
				}
//line /usr/local/go/src/crypto/x509/parser.go:659
				// _ = "end of CoverTab[18750]"
//line /usr/local/go/src/crypto/x509/parser.go:659
				_go_fuzz_dep_.CoverTab[18751]++
										out.BasicConstraintsValid = true
										out.MaxPathLenZero = out.MaxPathLen == 0
//line /usr/local/go/src/crypto/x509/parser.go:661
				// _ = "end of CoverTab[18751]"
			case 17:
//line /usr/local/go/src/crypto/x509/parser.go:662
				_go_fuzz_dep_.CoverTab[18752]++
										out.DNSNames, out.EmailAddresses, out.IPAddresses, out.URIs, err = parseSANExtension(e.Value)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:664
					_go_fuzz_dep_.CoverTab[18768]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:665
					// _ = "end of CoverTab[18768]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:666
					_go_fuzz_dep_.CoverTab[18769]++
//line /usr/local/go/src/crypto/x509/parser.go:666
					// _ = "end of CoverTab[18769]"
//line /usr/local/go/src/crypto/x509/parser.go:666
				}
//line /usr/local/go/src/crypto/x509/parser.go:666
				// _ = "end of CoverTab[18752]"
//line /usr/local/go/src/crypto/x509/parser.go:666
				_go_fuzz_dep_.CoverTab[18753]++

										if len(out.DNSNames) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:668
					_go_fuzz_dep_.CoverTab[18770]++
//line /usr/local/go/src/crypto/x509/parser.go:668
					return len(out.EmailAddresses) == 0
//line /usr/local/go/src/crypto/x509/parser.go:668
					// _ = "end of CoverTab[18770]"
//line /usr/local/go/src/crypto/x509/parser.go:668
				}() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:668
					_go_fuzz_dep_.CoverTab[18771]++
//line /usr/local/go/src/crypto/x509/parser.go:668
					return len(out.IPAddresses) == 0
//line /usr/local/go/src/crypto/x509/parser.go:668
					// _ = "end of CoverTab[18771]"
//line /usr/local/go/src/crypto/x509/parser.go:668
				}() && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:668
					_go_fuzz_dep_.CoverTab[18772]++
//line /usr/local/go/src/crypto/x509/parser.go:668
					return len(out.URIs) == 0
//line /usr/local/go/src/crypto/x509/parser.go:668
					// _ = "end of CoverTab[18772]"
//line /usr/local/go/src/crypto/x509/parser.go:668
				}() {
//line /usr/local/go/src/crypto/x509/parser.go:668
					_go_fuzz_dep_.CoverTab[18773]++

											unhandled = true
//line /usr/local/go/src/crypto/x509/parser.go:670
					// _ = "end of CoverTab[18773]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:671
					_go_fuzz_dep_.CoverTab[18774]++
//line /usr/local/go/src/crypto/x509/parser.go:671
					// _ = "end of CoverTab[18774]"
//line /usr/local/go/src/crypto/x509/parser.go:671
				}
//line /usr/local/go/src/crypto/x509/parser.go:671
				// _ = "end of CoverTab[18753]"

			case 30:
//line /usr/local/go/src/crypto/x509/parser.go:673
				_go_fuzz_dep_.CoverTab[18754]++
										unhandled, err = parseNameConstraintsExtension(out, e)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:675
					_go_fuzz_dep_.CoverTab[18775]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:676
					// _ = "end of CoverTab[18775]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:677
					_go_fuzz_dep_.CoverTab[18776]++
//line /usr/local/go/src/crypto/x509/parser.go:677
					// _ = "end of CoverTab[18776]"
//line /usr/local/go/src/crypto/x509/parser.go:677
				}
//line /usr/local/go/src/crypto/x509/parser.go:677
				// _ = "end of CoverTab[18754]"

			case 31:
//line /usr/local/go/src/crypto/x509/parser.go:679
				_go_fuzz_dep_.CoverTab[18755]++

//line /usr/local/go/src/crypto/x509/parser.go:692
				val := cryptobyte.String(e.Value)
				if !val.ReadASN1(&val, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:693
					_go_fuzz_dep_.CoverTab[18777]++
											return errors.New("x509: invalid CRL distribution points")
//line /usr/local/go/src/crypto/x509/parser.go:694
					// _ = "end of CoverTab[18777]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:695
					_go_fuzz_dep_.CoverTab[18778]++
//line /usr/local/go/src/crypto/x509/parser.go:695
					// _ = "end of CoverTab[18778]"
//line /usr/local/go/src/crypto/x509/parser.go:695
				}
//line /usr/local/go/src/crypto/x509/parser.go:695
				// _ = "end of CoverTab[18755]"
//line /usr/local/go/src/crypto/x509/parser.go:695
				_go_fuzz_dep_.CoverTab[18756]++
										for !val.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:696
					_go_fuzz_dep_.CoverTab[18779]++
											var dpDER cryptobyte.String
											if !val.ReadASN1(&dpDER, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:698
						_go_fuzz_dep_.CoverTab[18784]++
												return errors.New("x509: invalid CRL distribution point")
//line /usr/local/go/src/crypto/x509/parser.go:699
						// _ = "end of CoverTab[18784]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:700
						_go_fuzz_dep_.CoverTab[18785]++
//line /usr/local/go/src/crypto/x509/parser.go:700
						// _ = "end of CoverTab[18785]"
//line /usr/local/go/src/crypto/x509/parser.go:700
					}
//line /usr/local/go/src/crypto/x509/parser.go:700
					// _ = "end of CoverTab[18779]"
//line /usr/local/go/src/crypto/x509/parser.go:700
					_go_fuzz_dep_.CoverTab[18780]++
											var dpNameDER cryptobyte.String
											var dpNamePresent bool
											if !dpDER.ReadOptionalASN1(&dpNameDER, &dpNamePresent, cryptobyte_asn1.Tag(0).Constructed().ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:703
						_go_fuzz_dep_.CoverTab[18786]++
												return errors.New("x509: invalid CRL distribution point")
//line /usr/local/go/src/crypto/x509/parser.go:704
						// _ = "end of CoverTab[18786]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:705
						_go_fuzz_dep_.CoverTab[18787]++
//line /usr/local/go/src/crypto/x509/parser.go:705
						// _ = "end of CoverTab[18787]"
//line /usr/local/go/src/crypto/x509/parser.go:705
					}
//line /usr/local/go/src/crypto/x509/parser.go:705
					// _ = "end of CoverTab[18780]"
//line /usr/local/go/src/crypto/x509/parser.go:705
					_go_fuzz_dep_.CoverTab[18781]++
											if !dpNamePresent {
//line /usr/local/go/src/crypto/x509/parser.go:706
						_go_fuzz_dep_.CoverTab[18788]++
												continue
//line /usr/local/go/src/crypto/x509/parser.go:707
						// _ = "end of CoverTab[18788]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:708
						_go_fuzz_dep_.CoverTab[18789]++
//line /usr/local/go/src/crypto/x509/parser.go:708
						// _ = "end of CoverTab[18789]"
//line /usr/local/go/src/crypto/x509/parser.go:708
					}
//line /usr/local/go/src/crypto/x509/parser.go:708
					// _ = "end of CoverTab[18781]"
//line /usr/local/go/src/crypto/x509/parser.go:708
					_go_fuzz_dep_.CoverTab[18782]++
											if !dpNameDER.ReadASN1(&dpNameDER, cryptobyte_asn1.Tag(0).Constructed().ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:709
						_go_fuzz_dep_.CoverTab[18790]++
												return errors.New("x509: invalid CRL distribution point")
//line /usr/local/go/src/crypto/x509/parser.go:710
						// _ = "end of CoverTab[18790]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:711
						_go_fuzz_dep_.CoverTab[18791]++
//line /usr/local/go/src/crypto/x509/parser.go:711
						// _ = "end of CoverTab[18791]"
//line /usr/local/go/src/crypto/x509/parser.go:711
					}
//line /usr/local/go/src/crypto/x509/parser.go:711
					// _ = "end of CoverTab[18782]"
//line /usr/local/go/src/crypto/x509/parser.go:711
					_go_fuzz_dep_.CoverTab[18783]++
											for !dpNameDER.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:712
						_go_fuzz_dep_.CoverTab[18792]++
												if !dpNameDER.PeekASN1Tag(cryptobyte_asn1.Tag(6).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:713
							_go_fuzz_dep_.CoverTab[18795]++
													break
//line /usr/local/go/src/crypto/x509/parser.go:714
							// _ = "end of CoverTab[18795]"
						} else {
//line /usr/local/go/src/crypto/x509/parser.go:715
							_go_fuzz_dep_.CoverTab[18796]++
//line /usr/local/go/src/crypto/x509/parser.go:715
							// _ = "end of CoverTab[18796]"
//line /usr/local/go/src/crypto/x509/parser.go:715
						}
//line /usr/local/go/src/crypto/x509/parser.go:715
						// _ = "end of CoverTab[18792]"
//line /usr/local/go/src/crypto/x509/parser.go:715
						_go_fuzz_dep_.CoverTab[18793]++
												var uri cryptobyte.String
												if !dpNameDER.ReadASN1(&uri, cryptobyte_asn1.Tag(6).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:717
							_go_fuzz_dep_.CoverTab[18797]++
													return errors.New("x509: invalid CRL distribution point")
//line /usr/local/go/src/crypto/x509/parser.go:718
							// _ = "end of CoverTab[18797]"
						} else {
//line /usr/local/go/src/crypto/x509/parser.go:719
							_go_fuzz_dep_.CoverTab[18798]++
//line /usr/local/go/src/crypto/x509/parser.go:719
							// _ = "end of CoverTab[18798]"
//line /usr/local/go/src/crypto/x509/parser.go:719
						}
//line /usr/local/go/src/crypto/x509/parser.go:719
						// _ = "end of CoverTab[18793]"
//line /usr/local/go/src/crypto/x509/parser.go:719
						_go_fuzz_dep_.CoverTab[18794]++
												out.CRLDistributionPoints = append(out.CRLDistributionPoints, string(uri))
//line /usr/local/go/src/crypto/x509/parser.go:720
						// _ = "end of CoverTab[18794]"
					}
//line /usr/local/go/src/crypto/x509/parser.go:721
					// _ = "end of CoverTab[18783]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:722
				// _ = "end of CoverTab[18756]"

			case 35:
//line /usr/local/go/src/crypto/x509/parser.go:724
				_go_fuzz_dep_.CoverTab[18757]++

										val := cryptobyte.String(e.Value)
										var akid cryptobyte.String
										if !val.ReadASN1(&akid, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:728
					_go_fuzz_dep_.CoverTab[18799]++
											return errors.New("x509: invalid authority key identifier")
//line /usr/local/go/src/crypto/x509/parser.go:729
					// _ = "end of CoverTab[18799]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:730
					_go_fuzz_dep_.CoverTab[18800]++
//line /usr/local/go/src/crypto/x509/parser.go:730
					// _ = "end of CoverTab[18800]"
//line /usr/local/go/src/crypto/x509/parser.go:730
				}
//line /usr/local/go/src/crypto/x509/parser.go:730
				// _ = "end of CoverTab[18757]"
//line /usr/local/go/src/crypto/x509/parser.go:730
				_go_fuzz_dep_.CoverTab[18758]++
										if akid.PeekASN1Tag(cryptobyte_asn1.Tag(0).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:731
					_go_fuzz_dep_.CoverTab[18801]++
											if !akid.ReadASN1(&akid, cryptobyte_asn1.Tag(0).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:732
						_go_fuzz_dep_.CoverTab[18803]++
												return errors.New("x509: invalid authority key identifier")
//line /usr/local/go/src/crypto/x509/parser.go:733
						// _ = "end of CoverTab[18803]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:734
						_go_fuzz_dep_.CoverTab[18804]++
//line /usr/local/go/src/crypto/x509/parser.go:734
						// _ = "end of CoverTab[18804]"
//line /usr/local/go/src/crypto/x509/parser.go:734
					}
//line /usr/local/go/src/crypto/x509/parser.go:734
					// _ = "end of CoverTab[18801]"
//line /usr/local/go/src/crypto/x509/parser.go:734
					_go_fuzz_dep_.CoverTab[18802]++
											out.AuthorityKeyId = akid
//line /usr/local/go/src/crypto/x509/parser.go:735
					// _ = "end of CoverTab[18802]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:736
					_go_fuzz_dep_.CoverTab[18805]++
//line /usr/local/go/src/crypto/x509/parser.go:736
					// _ = "end of CoverTab[18805]"
//line /usr/local/go/src/crypto/x509/parser.go:736
				}
//line /usr/local/go/src/crypto/x509/parser.go:736
				// _ = "end of CoverTab[18758]"
			case 37:
//line /usr/local/go/src/crypto/x509/parser.go:737
				_go_fuzz_dep_.CoverTab[18759]++
										out.ExtKeyUsage, out.UnknownExtKeyUsage, err = parseExtKeyUsageExtension(e.Value)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:739
					_go_fuzz_dep_.CoverTab[18806]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:740
					// _ = "end of CoverTab[18806]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:741
					_go_fuzz_dep_.CoverTab[18807]++
//line /usr/local/go/src/crypto/x509/parser.go:741
					// _ = "end of CoverTab[18807]"
//line /usr/local/go/src/crypto/x509/parser.go:741
				}
//line /usr/local/go/src/crypto/x509/parser.go:741
				// _ = "end of CoverTab[18759]"
			case 14:
//line /usr/local/go/src/crypto/x509/parser.go:742
				_go_fuzz_dep_.CoverTab[18760]++

										val := cryptobyte.String(e.Value)
										var skid cryptobyte.String
										if !val.ReadASN1(&skid, cryptobyte_asn1.OCTET_STRING) {
//line /usr/local/go/src/crypto/x509/parser.go:746
					_go_fuzz_dep_.CoverTab[18808]++
											return errors.New("x509: invalid subject key identifier")
//line /usr/local/go/src/crypto/x509/parser.go:747
					// _ = "end of CoverTab[18808]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:748
					_go_fuzz_dep_.CoverTab[18809]++
//line /usr/local/go/src/crypto/x509/parser.go:748
					// _ = "end of CoverTab[18809]"
//line /usr/local/go/src/crypto/x509/parser.go:748
				}
//line /usr/local/go/src/crypto/x509/parser.go:748
				// _ = "end of CoverTab[18760]"
//line /usr/local/go/src/crypto/x509/parser.go:748
				_go_fuzz_dep_.CoverTab[18761]++
										out.SubjectKeyId = skid
//line /usr/local/go/src/crypto/x509/parser.go:749
				// _ = "end of CoverTab[18761]"
			case 32:
//line /usr/local/go/src/crypto/x509/parser.go:750
				_go_fuzz_dep_.CoverTab[18762]++
										out.PolicyIdentifiers, err = parseCertificatePoliciesExtension(e.Value)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:752
					_go_fuzz_dep_.CoverTab[18810]++
											return err
//line /usr/local/go/src/crypto/x509/parser.go:753
					// _ = "end of CoverTab[18810]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:754
					_go_fuzz_dep_.CoverTab[18811]++
//line /usr/local/go/src/crypto/x509/parser.go:754
					// _ = "end of CoverTab[18811]"
//line /usr/local/go/src/crypto/x509/parser.go:754
				}
//line /usr/local/go/src/crypto/x509/parser.go:754
				// _ = "end of CoverTab[18762]"
			default:
//line /usr/local/go/src/crypto/x509/parser.go:755
				_go_fuzz_dep_.CoverTab[18763]++

										unhandled = true
//line /usr/local/go/src/crypto/x509/parser.go:757
				// _ = "end of CoverTab[18763]"
			}
//line /usr/local/go/src/crypto/x509/parser.go:758
			// _ = "end of CoverTab[18748]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:759
			_go_fuzz_dep_.CoverTab[18812]++
//line /usr/local/go/src/crypto/x509/parser.go:759
			if e.Id.Equal(oidExtensionAuthorityInfoAccess) {
//line /usr/local/go/src/crypto/x509/parser.go:759
				_go_fuzz_dep_.CoverTab[18813]++

										val := cryptobyte.String(e.Value)
										if !val.ReadASN1(&val, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:762
					_go_fuzz_dep_.CoverTab[18815]++
											return errors.New("x509: invalid authority info access")
//line /usr/local/go/src/crypto/x509/parser.go:763
					// _ = "end of CoverTab[18815]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:764
					_go_fuzz_dep_.CoverTab[18816]++
//line /usr/local/go/src/crypto/x509/parser.go:764
					// _ = "end of CoverTab[18816]"
//line /usr/local/go/src/crypto/x509/parser.go:764
				}
//line /usr/local/go/src/crypto/x509/parser.go:764
				// _ = "end of CoverTab[18813]"
//line /usr/local/go/src/crypto/x509/parser.go:764
				_go_fuzz_dep_.CoverTab[18814]++
										for !val.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:765
					_go_fuzz_dep_.CoverTab[18817]++
											var aiaDER cryptobyte.String
											if !val.ReadASN1(&aiaDER, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:767
						_go_fuzz_dep_.CoverTab[18822]++
												return errors.New("x509: invalid authority info access")
//line /usr/local/go/src/crypto/x509/parser.go:768
						// _ = "end of CoverTab[18822]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:769
						_go_fuzz_dep_.CoverTab[18823]++
//line /usr/local/go/src/crypto/x509/parser.go:769
						// _ = "end of CoverTab[18823]"
//line /usr/local/go/src/crypto/x509/parser.go:769
					}
//line /usr/local/go/src/crypto/x509/parser.go:769
					// _ = "end of CoverTab[18817]"
//line /usr/local/go/src/crypto/x509/parser.go:769
					_go_fuzz_dep_.CoverTab[18818]++
											var method asn1.ObjectIdentifier
											if !aiaDER.ReadASN1ObjectIdentifier(&method) {
//line /usr/local/go/src/crypto/x509/parser.go:771
						_go_fuzz_dep_.CoverTab[18824]++
												return errors.New("x509: invalid authority info access")
//line /usr/local/go/src/crypto/x509/parser.go:772
						// _ = "end of CoverTab[18824]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:773
						_go_fuzz_dep_.CoverTab[18825]++
//line /usr/local/go/src/crypto/x509/parser.go:773
						// _ = "end of CoverTab[18825]"
//line /usr/local/go/src/crypto/x509/parser.go:773
					}
//line /usr/local/go/src/crypto/x509/parser.go:773
					// _ = "end of CoverTab[18818]"
//line /usr/local/go/src/crypto/x509/parser.go:773
					_go_fuzz_dep_.CoverTab[18819]++
											if !aiaDER.PeekASN1Tag(cryptobyte_asn1.Tag(6).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:774
						_go_fuzz_dep_.CoverTab[18826]++
												continue
//line /usr/local/go/src/crypto/x509/parser.go:775
						// _ = "end of CoverTab[18826]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:776
						_go_fuzz_dep_.CoverTab[18827]++
//line /usr/local/go/src/crypto/x509/parser.go:776
						// _ = "end of CoverTab[18827]"
//line /usr/local/go/src/crypto/x509/parser.go:776
					}
//line /usr/local/go/src/crypto/x509/parser.go:776
					// _ = "end of CoverTab[18819]"
//line /usr/local/go/src/crypto/x509/parser.go:776
					_go_fuzz_dep_.CoverTab[18820]++
											if !aiaDER.ReadASN1(&aiaDER, cryptobyte_asn1.Tag(6).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:777
						_go_fuzz_dep_.CoverTab[18828]++
												return errors.New("x509: invalid authority info access")
//line /usr/local/go/src/crypto/x509/parser.go:778
						// _ = "end of CoverTab[18828]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:779
						_go_fuzz_dep_.CoverTab[18829]++
//line /usr/local/go/src/crypto/x509/parser.go:779
						// _ = "end of CoverTab[18829]"
//line /usr/local/go/src/crypto/x509/parser.go:779
					}
//line /usr/local/go/src/crypto/x509/parser.go:779
					// _ = "end of CoverTab[18820]"
//line /usr/local/go/src/crypto/x509/parser.go:779
					_go_fuzz_dep_.CoverTab[18821]++
											switch {
					case method.Equal(oidAuthorityInfoAccessOcsp):
//line /usr/local/go/src/crypto/x509/parser.go:781
						_go_fuzz_dep_.CoverTab[18830]++
												out.OCSPServer = append(out.OCSPServer, string(aiaDER))
//line /usr/local/go/src/crypto/x509/parser.go:782
						// _ = "end of CoverTab[18830]"
					case method.Equal(oidAuthorityInfoAccessIssuers):
//line /usr/local/go/src/crypto/x509/parser.go:783
						_go_fuzz_dep_.CoverTab[18831]++
												out.IssuingCertificateURL = append(out.IssuingCertificateURL, string(aiaDER))
//line /usr/local/go/src/crypto/x509/parser.go:784
						// _ = "end of CoverTab[18831]"
//line /usr/local/go/src/crypto/x509/parser.go:784
					default:
//line /usr/local/go/src/crypto/x509/parser.go:784
						_go_fuzz_dep_.CoverTab[18832]++
//line /usr/local/go/src/crypto/x509/parser.go:784
						// _ = "end of CoverTab[18832]"
					}
//line /usr/local/go/src/crypto/x509/parser.go:785
					// _ = "end of CoverTab[18821]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:786
				// _ = "end of CoverTab[18814]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:787
				_go_fuzz_dep_.CoverTab[18833]++

										unhandled = true
//line /usr/local/go/src/crypto/x509/parser.go:789
				// _ = "end of CoverTab[18833]"
			}
//line /usr/local/go/src/crypto/x509/parser.go:790
			// _ = "end of CoverTab[18812]"
//line /usr/local/go/src/crypto/x509/parser.go:790
		}
//line /usr/local/go/src/crypto/x509/parser.go:790
		// _ = "end of CoverTab[18743]"
//line /usr/local/go/src/crypto/x509/parser.go:790
		_go_fuzz_dep_.CoverTab[18744]++

								if e.Critical && func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:792
			_go_fuzz_dep_.CoverTab[18834]++
//line /usr/local/go/src/crypto/x509/parser.go:792
			return unhandled
//line /usr/local/go/src/crypto/x509/parser.go:792
			// _ = "end of CoverTab[18834]"
//line /usr/local/go/src/crypto/x509/parser.go:792
		}() {
//line /usr/local/go/src/crypto/x509/parser.go:792
			_go_fuzz_dep_.CoverTab[18835]++
									out.UnhandledCriticalExtensions = append(out.UnhandledCriticalExtensions, e.Id)
//line /usr/local/go/src/crypto/x509/parser.go:793
			// _ = "end of CoverTab[18835]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:794
			_go_fuzz_dep_.CoverTab[18836]++
//line /usr/local/go/src/crypto/x509/parser.go:794
			// _ = "end of CoverTab[18836]"
//line /usr/local/go/src/crypto/x509/parser.go:794
		}
//line /usr/local/go/src/crypto/x509/parser.go:794
		// _ = "end of CoverTab[18744]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:795
	// _ = "end of CoverTab[18741]"
//line /usr/local/go/src/crypto/x509/parser.go:795
	_go_fuzz_dep_.CoverTab[18742]++

							return nil
//line /usr/local/go/src/crypto/x509/parser.go:797
	// _ = "end of CoverTab[18742]"
}

func parseCertificate(der []byte) (*Certificate, error) {
//line /usr/local/go/src/crypto/x509/parser.go:800
	_go_fuzz_dep_.CoverTab[18837]++
							cert := &Certificate{}

							input := cryptobyte.String(der)

//line /usr/local/go/src/crypto/x509/parser.go:807
	if !input.ReadASN1Element(&input, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:807
		_go_fuzz_dep_.CoverTab[18864]++
								return nil, errors.New("x509: malformed certificate")
//line /usr/local/go/src/crypto/x509/parser.go:808
		// _ = "end of CoverTab[18864]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:809
		_go_fuzz_dep_.CoverTab[18865]++
//line /usr/local/go/src/crypto/x509/parser.go:809
		// _ = "end of CoverTab[18865]"
//line /usr/local/go/src/crypto/x509/parser.go:809
	}
//line /usr/local/go/src/crypto/x509/parser.go:809
	// _ = "end of CoverTab[18837]"
//line /usr/local/go/src/crypto/x509/parser.go:809
	_go_fuzz_dep_.CoverTab[18838]++
							cert.Raw = input
							if !input.ReadASN1(&input, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:811
		_go_fuzz_dep_.CoverTab[18866]++
								return nil, errors.New("x509: malformed certificate")
//line /usr/local/go/src/crypto/x509/parser.go:812
		// _ = "end of CoverTab[18866]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:813
		_go_fuzz_dep_.CoverTab[18867]++
//line /usr/local/go/src/crypto/x509/parser.go:813
		// _ = "end of CoverTab[18867]"
//line /usr/local/go/src/crypto/x509/parser.go:813
	}
//line /usr/local/go/src/crypto/x509/parser.go:813
	// _ = "end of CoverTab[18838]"
//line /usr/local/go/src/crypto/x509/parser.go:813
	_go_fuzz_dep_.CoverTab[18839]++

							var tbs cryptobyte.String

//line /usr/local/go/src/crypto/x509/parser.go:818
	if !input.ReadASN1Element(&tbs, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:818
		_go_fuzz_dep_.CoverTab[18868]++
								return nil, errors.New("x509: malformed tbs certificate")
//line /usr/local/go/src/crypto/x509/parser.go:819
		// _ = "end of CoverTab[18868]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:820
		_go_fuzz_dep_.CoverTab[18869]++
//line /usr/local/go/src/crypto/x509/parser.go:820
		// _ = "end of CoverTab[18869]"
//line /usr/local/go/src/crypto/x509/parser.go:820
	}
//line /usr/local/go/src/crypto/x509/parser.go:820
	// _ = "end of CoverTab[18839]"
//line /usr/local/go/src/crypto/x509/parser.go:820
	_go_fuzz_dep_.CoverTab[18840]++
							cert.RawTBSCertificate = tbs
							if !tbs.ReadASN1(&tbs, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:822
		_go_fuzz_dep_.CoverTab[18870]++
								return nil, errors.New("x509: malformed tbs certificate")
//line /usr/local/go/src/crypto/x509/parser.go:823
		// _ = "end of CoverTab[18870]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:824
		_go_fuzz_dep_.CoverTab[18871]++
//line /usr/local/go/src/crypto/x509/parser.go:824
		// _ = "end of CoverTab[18871]"
//line /usr/local/go/src/crypto/x509/parser.go:824
	}
//line /usr/local/go/src/crypto/x509/parser.go:824
	// _ = "end of CoverTab[18840]"
//line /usr/local/go/src/crypto/x509/parser.go:824
	_go_fuzz_dep_.CoverTab[18841]++

							if !tbs.ReadOptionalASN1Integer(&cert.Version, cryptobyte_asn1.Tag(0).Constructed().ContextSpecific(), 0) {
//line /usr/local/go/src/crypto/x509/parser.go:826
		_go_fuzz_dep_.CoverTab[18872]++
								return nil, errors.New("x509: malformed version")
//line /usr/local/go/src/crypto/x509/parser.go:827
		// _ = "end of CoverTab[18872]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:828
		_go_fuzz_dep_.CoverTab[18873]++
//line /usr/local/go/src/crypto/x509/parser.go:828
		// _ = "end of CoverTab[18873]"
//line /usr/local/go/src/crypto/x509/parser.go:828
	}
//line /usr/local/go/src/crypto/x509/parser.go:828
	// _ = "end of CoverTab[18841]"
//line /usr/local/go/src/crypto/x509/parser.go:828
	_go_fuzz_dep_.CoverTab[18842]++
							if cert.Version < 0 {
//line /usr/local/go/src/crypto/x509/parser.go:829
		_go_fuzz_dep_.CoverTab[18874]++
								return nil, errors.New("x509: malformed version")
//line /usr/local/go/src/crypto/x509/parser.go:830
		// _ = "end of CoverTab[18874]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:831
		_go_fuzz_dep_.CoverTab[18875]++
//line /usr/local/go/src/crypto/x509/parser.go:831
		// _ = "end of CoverTab[18875]"
//line /usr/local/go/src/crypto/x509/parser.go:831
	}
//line /usr/local/go/src/crypto/x509/parser.go:831
	// _ = "end of CoverTab[18842]"
//line /usr/local/go/src/crypto/x509/parser.go:831
	_go_fuzz_dep_.CoverTab[18843]++

//line /usr/local/go/src/crypto/x509/parser.go:834
	cert.Version++
	if cert.Version > 3 {
//line /usr/local/go/src/crypto/x509/parser.go:835
		_go_fuzz_dep_.CoverTab[18876]++
								return nil, errors.New("x509: invalid version")
//line /usr/local/go/src/crypto/x509/parser.go:836
		// _ = "end of CoverTab[18876]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:837
		_go_fuzz_dep_.CoverTab[18877]++
//line /usr/local/go/src/crypto/x509/parser.go:837
		// _ = "end of CoverTab[18877]"
//line /usr/local/go/src/crypto/x509/parser.go:837
	}
//line /usr/local/go/src/crypto/x509/parser.go:837
	// _ = "end of CoverTab[18843]"
//line /usr/local/go/src/crypto/x509/parser.go:837
	_go_fuzz_dep_.CoverTab[18844]++

							serial := new(big.Int)
							if !tbs.ReadASN1Integer(serial) {
//line /usr/local/go/src/crypto/x509/parser.go:840
		_go_fuzz_dep_.CoverTab[18878]++
								return nil, errors.New("x509: malformed serial number")
//line /usr/local/go/src/crypto/x509/parser.go:841
		// _ = "end of CoverTab[18878]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:842
		_go_fuzz_dep_.CoverTab[18879]++
//line /usr/local/go/src/crypto/x509/parser.go:842
		// _ = "end of CoverTab[18879]"
//line /usr/local/go/src/crypto/x509/parser.go:842
	}
//line /usr/local/go/src/crypto/x509/parser.go:842
	// _ = "end of CoverTab[18844]"
//line /usr/local/go/src/crypto/x509/parser.go:842
	_go_fuzz_dep_.CoverTab[18845]++

//line /usr/local/go/src/crypto/x509/parser.go:848
	cert.SerialNumber = serial

	var sigAISeq cryptobyte.String
	if !tbs.ReadASN1(&sigAISeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:851
		_go_fuzz_dep_.CoverTab[18880]++
								return nil, errors.New("x509: malformed signature algorithm identifier")
//line /usr/local/go/src/crypto/x509/parser.go:852
		// _ = "end of CoverTab[18880]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:853
		_go_fuzz_dep_.CoverTab[18881]++
//line /usr/local/go/src/crypto/x509/parser.go:853
		// _ = "end of CoverTab[18881]"
//line /usr/local/go/src/crypto/x509/parser.go:853
	}
//line /usr/local/go/src/crypto/x509/parser.go:853
	// _ = "end of CoverTab[18845]"
//line /usr/local/go/src/crypto/x509/parser.go:853
	_go_fuzz_dep_.CoverTab[18846]++
	// Before parsing the inner algorithm identifier, extract
	// the outer algorithm identifier and make sure that they
	// match.
	var outerSigAISeq cryptobyte.String
	if !input.ReadASN1(&outerSigAISeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:858
		_go_fuzz_dep_.CoverTab[18882]++
								return nil, errors.New("x509: malformed algorithm identifier")
//line /usr/local/go/src/crypto/x509/parser.go:859
		// _ = "end of CoverTab[18882]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:860
		_go_fuzz_dep_.CoverTab[18883]++
//line /usr/local/go/src/crypto/x509/parser.go:860
		// _ = "end of CoverTab[18883]"
//line /usr/local/go/src/crypto/x509/parser.go:860
	}
//line /usr/local/go/src/crypto/x509/parser.go:860
	// _ = "end of CoverTab[18846]"
//line /usr/local/go/src/crypto/x509/parser.go:860
	_go_fuzz_dep_.CoverTab[18847]++
							if !bytes.Equal(outerSigAISeq, sigAISeq) {
//line /usr/local/go/src/crypto/x509/parser.go:861
		_go_fuzz_dep_.CoverTab[18884]++
								return nil, errors.New("x509: inner and outer signature algorithm identifiers don't match")
//line /usr/local/go/src/crypto/x509/parser.go:862
		// _ = "end of CoverTab[18884]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:863
		_go_fuzz_dep_.CoverTab[18885]++
//line /usr/local/go/src/crypto/x509/parser.go:863
		// _ = "end of CoverTab[18885]"
//line /usr/local/go/src/crypto/x509/parser.go:863
	}
//line /usr/local/go/src/crypto/x509/parser.go:863
	// _ = "end of CoverTab[18847]"
//line /usr/local/go/src/crypto/x509/parser.go:863
	_go_fuzz_dep_.CoverTab[18848]++
							sigAI, err := parseAI(sigAISeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:865
		_go_fuzz_dep_.CoverTab[18886]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:866
		// _ = "end of CoverTab[18886]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:867
		_go_fuzz_dep_.CoverTab[18887]++
//line /usr/local/go/src/crypto/x509/parser.go:867
		// _ = "end of CoverTab[18887]"
//line /usr/local/go/src/crypto/x509/parser.go:867
	}
//line /usr/local/go/src/crypto/x509/parser.go:867
	// _ = "end of CoverTab[18848]"
//line /usr/local/go/src/crypto/x509/parser.go:867
	_go_fuzz_dep_.CoverTab[18849]++
							cert.SignatureAlgorithm = getSignatureAlgorithmFromAI(sigAI)

							var issuerSeq cryptobyte.String
							if !tbs.ReadASN1Element(&issuerSeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:871
		_go_fuzz_dep_.CoverTab[18888]++
								return nil, errors.New("x509: malformed issuer")
//line /usr/local/go/src/crypto/x509/parser.go:872
		// _ = "end of CoverTab[18888]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:873
		_go_fuzz_dep_.CoverTab[18889]++
//line /usr/local/go/src/crypto/x509/parser.go:873
		// _ = "end of CoverTab[18889]"
//line /usr/local/go/src/crypto/x509/parser.go:873
	}
//line /usr/local/go/src/crypto/x509/parser.go:873
	// _ = "end of CoverTab[18849]"
//line /usr/local/go/src/crypto/x509/parser.go:873
	_go_fuzz_dep_.CoverTab[18850]++
							cert.RawIssuer = issuerSeq
							issuerRDNs, err := parseName(issuerSeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:876
		_go_fuzz_dep_.CoverTab[18890]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:877
		// _ = "end of CoverTab[18890]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:878
		_go_fuzz_dep_.CoverTab[18891]++
//line /usr/local/go/src/crypto/x509/parser.go:878
		// _ = "end of CoverTab[18891]"
//line /usr/local/go/src/crypto/x509/parser.go:878
	}
//line /usr/local/go/src/crypto/x509/parser.go:878
	// _ = "end of CoverTab[18850]"
//line /usr/local/go/src/crypto/x509/parser.go:878
	_go_fuzz_dep_.CoverTab[18851]++
							cert.Issuer.FillFromRDNSequence(issuerRDNs)

							var validity cryptobyte.String
							if !tbs.ReadASN1(&validity, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:882
		_go_fuzz_dep_.CoverTab[18892]++
								return nil, errors.New("x509: malformed validity")
//line /usr/local/go/src/crypto/x509/parser.go:883
		// _ = "end of CoverTab[18892]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:884
		_go_fuzz_dep_.CoverTab[18893]++
//line /usr/local/go/src/crypto/x509/parser.go:884
		// _ = "end of CoverTab[18893]"
//line /usr/local/go/src/crypto/x509/parser.go:884
	}
//line /usr/local/go/src/crypto/x509/parser.go:884
	// _ = "end of CoverTab[18851]"
//line /usr/local/go/src/crypto/x509/parser.go:884
	_go_fuzz_dep_.CoverTab[18852]++
							cert.NotBefore, cert.NotAfter, err = parseValidity(validity)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:886
		_go_fuzz_dep_.CoverTab[18894]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:887
		// _ = "end of CoverTab[18894]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:888
		_go_fuzz_dep_.CoverTab[18895]++
//line /usr/local/go/src/crypto/x509/parser.go:888
		// _ = "end of CoverTab[18895]"
//line /usr/local/go/src/crypto/x509/parser.go:888
	}
//line /usr/local/go/src/crypto/x509/parser.go:888
	// _ = "end of CoverTab[18852]"
//line /usr/local/go/src/crypto/x509/parser.go:888
	_go_fuzz_dep_.CoverTab[18853]++

							var subjectSeq cryptobyte.String
							if !tbs.ReadASN1Element(&subjectSeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:891
		_go_fuzz_dep_.CoverTab[18896]++
								return nil, errors.New("x509: malformed issuer")
//line /usr/local/go/src/crypto/x509/parser.go:892
		// _ = "end of CoverTab[18896]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:893
		_go_fuzz_dep_.CoverTab[18897]++
//line /usr/local/go/src/crypto/x509/parser.go:893
		// _ = "end of CoverTab[18897]"
//line /usr/local/go/src/crypto/x509/parser.go:893
	}
//line /usr/local/go/src/crypto/x509/parser.go:893
	// _ = "end of CoverTab[18853]"
//line /usr/local/go/src/crypto/x509/parser.go:893
	_go_fuzz_dep_.CoverTab[18854]++
							cert.RawSubject = subjectSeq
							subjectRDNs, err := parseName(subjectSeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:896
		_go_fuzz_dep_.CoverTab[18898]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:897
		// _ = "end of CoverTab[18898]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:898
		_go_fuzz_dep_.CoverTab[18899]++
//line /usr/local/go/src/crypto/x509/parser.go:898
		// _ = "end of CoverTab[18899]"
//line /usr/local/go/src/crypto/x509/parser.go:898
	}
//line /usr/local/go/src/crypto/x509/parser.go:898
	// _ = "end of CoverTab[18854]"
//line /usr/local/go/src/crypto/x509/parser.go:898
	_go_fuzz_dep_.CoverTab[18855]++
							cert.Subject.FillFromRDNSequence(subjectRDNs)

							var spki cryptobyte.String
							if !tbs.ReadASN1Element(&spki, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:902
		_go_fuzz_dep_.CoverTab[18900]++
								return nil, errors.New("x509: malformed spki")
//line /usr/local/go/src/crypto/x509/parser.go:903
		// _ = "end of CoverTab[18900]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:904
		_go_fuzz_dep_.CoverTab[18901]++
//line /usr/local/go/src/crypto/x509/parser.go:904
		// _ = "end of CoverTab[18901]"
//line /usr/local/go/src/crypto/x509/parser.go:904
	}
//line /usr/local/go/src/crypto/x509/parser.go:904
	// _ = "end of CoverTab[18855]"
//line /usr/local/go/src/crypto/x509/parser.go:904
	_go_fuzz_dep_.CoverTab[18856]++
							cert.RawSubjectPublicKeyInfo = spki
							if !spki.ReadASN1(&spki, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:906
		_go_fuzz_dep_.CoverTab[18902]++
								return nil, errors.New("x509: malformed spki")
//line /usr/local/go/src/crypto/x509/parser.go:907
		// _ = "end of CoverTab[18902]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:908
		_go_fuzz_dep_.CoverTab[18903]++
//line /usr/local/go/src/crypto/x509/parser.go:908
		// _ = "end of CoverTab[18903]"
//line /usr/local/go/src/crypto/x509/parser.go:908
	}
//line /usr/local/go/src/crypto/x509/parser.go:908
	// _ = "end of CoverTab[18856]"
//line /usr/local/go/src/crypto/x509/parser.go:908
	_go_fuzz_dep_.CoverTab[18857]++
							var pkAISeq cryptobyte.String
							if !spki.ReadASN1(&pkAISeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:910
		_go_fuzz_dep_.CoverTab[18904]++
								return nil, errors.New("x509: malformed public key algorithm identifier")
//line /usr/local/go/src/crypto/x509/parser.go:911
		// _ = "end of CoverTab[18904]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:912
		_go_fuzz_dep_.CoverTab[18905]++
//line /usr/local/go/src/crypto/x509/parser.go:912
		// _ = "end of CoverTab[18905]"
//line /usr/local/go/src/crypto/x509/parser.go:912
	}
//line /usr/local/go/src/crypto/x509/parser.go:912
	// _ = "end of CoverTab[18857]"
//line /usr/local/go/src/crypto/x509/parser.go:912
	_go_fuzz_dep_.CoverTab[18858]++
							pkAI, err := parseAI(pkAISeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:914
		_go_fuzz_dep_.CoverTab[18906]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:915
		// _ = "end of CoverTab[18906]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:916
		_go_fuzz_dep_.CoverTab[18907]++
//line /usr/local/go/src/crypto/x509/parser.go:916
		// _ = "end of CoverTab[18907]"
//line /usr/local/go/src/crypto/x509/parser.go:916
	}
//line /usr/local/go/src/crypto/x509/parser.go:916
	// _ = "end of CoverTab[18858]"
//line /usr/local/go/src/crypto/x509/parser.go:916
	_go_fuzz_dep_.CoverTab[18859]++
							cert.PublicKeyAlgorithm = getPublicKeyAlgorithmFromOID(pkAI.Algorithm)
							var spk asn1.BitString
							if !spki.ReadASN1BitString(&spk) {
//line /usr/local/go/src/crypto/x509/parser.go:919
		_go_fuzz_dep_.CoverTab[18908]++
								return nil, errors.New("x509: malformed subjectPublicKey")
//line /usr/local/go/src/crypto/x509/parser.go:920
		// _ = "end of CoverTab[18908]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:921
		_go_fuzz_dep_.CoverTab[18909]++
//line /usr/local/go/src/crypto/x509/parser.go:921
		// _ = "end of CoverTab[18909]"
//line /usr/local/go/src/crypto/x509/parser.go:921
	}
//line /usr/local/go/src/crypto/x509/parser.go:921
	// _ = "end of CoverTab[18859]"
//line /usr/local/go/src/crypto/x509/parser.go:921
	_go_fuzz_dep_.CoverTab[18860]++
							if cert.PublicKeyAlgorithm != UnknownPublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/parser.go:922
		_go_fuzz_dep_.CoverTab[18910]++
								cert.PublicKey, err = parsePublicKey(&publicKeyInfo{
			Algorithm:	pkAI,
			PublicKey:	spk,
		})
		if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:927
			_go_fuzz_dep_.CoverTab[18911]++
									return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:928
			// _ = "end of CoverTab[18911]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:929
			_go_fuzz_dep_.CoverTab[18912]++
//line /usr/local/go/src/crypto/x509/parser.go:929
			// _ = "end of CoverTab[18912]"
//line /usr/local/go/src/crypto/x509/parser.go:929
		}
//line /usr/local/go/src/crypto/x509/parser.go:929
		// _ = "end of CoverTab[18910]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:930
		_go_fuzz_dep_.CoverTab[18913]++
//line /usr/local/go/src/crypto/x509/parser.go:930
		// _ = "end of CoverTab[18913]"
//line /usr/local/go/src/crypto/x509/parser.go:930
	}
//line /usr/local/go/src/crypto/x509/parser.go:930
	// _ = "end of CoverTab[18860]"
//line /usr/local/go/src/crypto/x509/parser.go:930
	_go_fuzz_dep_.CoverTab[18861]++

							if cert.Version > 1 {
//line /usr/local/go/src/crypto/x509/parser.go:932
		_go_fuzz_dep_.CoverTab[18914]++
								if !tbs.SkipOptionalASN1(cryptobyte_asn1.Tag(1).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:933
			_go_fuzz_dep_.CoverTab[18917]++
									return nil, errors.New("x509: malformed issuerUniqueID")
//line /usr/local/go/src/crypto/x509/parser.go:934
			// _ = "end of CoverTab[18917]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:935
			_go_fuzz_dep_.CoverTab[18918]++
//line /usr/local/go/src/crypto/x509/parser.go:935
			// _ = "end of CoverTab[18918]"
//line /usr/local/go/src/crypto/x509/parser.go:935
		}
//line /usr/local/go/src/crypto/x509/parser.go:935
		// _ = "end of CoverTab[18914]"
//line /usr/local/go/src/crypto/x509/parser.go:935
		_go_fuzz_dep_.CoverTab[18915]++
								if !tbs.SkipOptionalASN1(cryptobyte_asn1.Tag(2).ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:936
			_go_fuzz_dep_.CoverTab[18919]++
									return nil, errors.New("x509: malformed subjectUniqueID")
//line /usr/local/go/src/crypto/x509/parser.go:937
			// _ = "end of CoverTab[18919]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:938
			_go_fuzz_dep_.CoverTab[18920]++
//line /usr/local/go/src/crypto/x509/parser.go:938
			// _ = "end of CoverTab[18920]"
//line /usr/local/go/src/crypto/x509/parser.go:938
		}
//line /usr/local/go/src/crypto/x509/parser.go:938
		// _ = "end of CoverTab[18915]"
//line /usr/local/go/src/crypto/x509/parser.go:938
		_go_fuzz_dep_.CoverTab[18916]++
								if cert.Version == 3 {
//line /usr/local/go/src/crypto/x509/parser.go:939
			_go_fuzz_dep_.CoverTab[18921]++
									var extensions cryptobyte.String
									var present bool
									if !tbs.ReadOptionalASN1(&extensions, &present, cryptobyte_asn1.Tag(3).Constructed().ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:942
				_go_fuzz_dep_.CoverTab[18923]++
										return nil, errors.New("x509: malformed extensions")
//line /usr/local/go/src/crypto/x509/parser.go:943
				// _ = "end of CoverTab[18923]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:944
				_go_fuzz_dep_.CoverTab[18924]++
//line /usr/local/go/src/crypto/x509/parser.go:944
				// _ = "end of CoverTab[18924]"
//line /usr/local/go/src/crypto/x509/parser.go:944
			}
//line /usr/local/go/src/crypto/x509/parser.go:944
			// _ = "end of CoverTab[18921]"
//line /usr/local/go/src/crypto/x509/parser.go:944
			_go_fuzz_dep_.CoverTab[18922]++
									if present {
//line /usr/local/go/src/crypto/x509/parser.go:945
				_go_fuzz_dep_.CoverTab[18925]++
										seenExts := make(map[string]bool)
										if !extensions.ReadASN1(&extensions, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:947
					_go_fuzz_dep_.CoverTab[18928]++
											return nil, errors.New("x509: malformed extensions")
//line /usr/local/go/src/crypto/x509/parser.go:948
					// _ = "end of CoverTab[18928]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:949
					_go_fuzz_dep_.CoverTab[18929]++
//line /usr/local/go/src/crypto/x509/parser.go:949
					// _ = "end of CoverTab[18929]"
//line /usr/local/go/src/crypto/x509/parser.go:949
				}
//line /usr/local/go/src/crypto/x509/parser.go:949
				// _ = "end of CoverTab[18925]"
//line /usr/local/go/src/crypto/x509/parser.go:949
				_go_fuzz_dep_.CoverTab[18926]++
										for !extensions.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:950
					_go_fuzz_dep_.CoverTab[18930]++
											var extension cryptobyte.String
											if !extensions.ReadASN1(&extension, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:952
						_go_fuzz_dep_.CoverTab[18934]++
												return nil, errors.New("x509: malformed extension")
//line /usr/local/go/src/crypto/x509/parser.go:953
						// _ = "end of CoverTab[18934]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:954
						_go_fuzz_dep_.CoverTab[18935]++
//line /usr/local/go/src/crypto/x509/parser.go:954
						// _ = "end of CoverTab[18935]"
//line /usr/local/go/src/crypto/x509/parser.go:954
					}
//line /usr/local/go/src/crypto/x509/parser.go:954
					// _ = "end of CoverTab[18930]"
//line /usr/local/go/src/crypto/x509/parser.go:954
					_go_fuzz_dep_.CoverTab[18931]++
											ext, err := parseExtension(extension)
											if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:956
						_go_fuzz_dep_.CoverTab[18936]++
												return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:957
						// _ = "end of CoverTab[18936]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:958
						_go_fuzz_dep_.CoverTab[18937]++
//line /usr/local/go/src/crypto/x509/parser.go:958
						// _ = "end of CoverTab[18937]"
//line /usr/local/go/src/crypto/x509/parser.go:958
					}
//line /usr/local/go/src/crypto/x509/parser.go:958
					// _ = "end of CoverTab[18931]"
//line /usr/local/go/src/crypto/x509/parser.go:958
					_go_fuzz_dep_.CoverTab[18932]++
											oidStr := ext.Id.String()
											if seenExts[oidStr] {
//line /usr/local/go/src/crypto/x509/parser.go:960
						_go_fuzz_dep_.CoverTab[18938]++
												return nil, errors.New("x509: certificate contains duplicate extensions")
//line /usr/local/go/src/crypto/x509/parser.go:961
						// _ = "end of CoverTab[18938]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:962
						_go_fuzz_dep_.CoverTab[18939]++
//line /usr/local/go/src/crypto/x509/parser.go:962
						// _ = "end of CoverTab[18939]"
//line /usr/local/go/src/crypto/x509/parser.go:962
					}
//line /usr/local/go/src/crypto/x509/parser.go:962
					// _ = "end of CoverTab[18932]"
//line /usr/local/go/src/crypto/x509/parser.go:962
					_go_fuzz_dep_.CoverTab[18933]++
											seenExts[oidStr] = true
											cert.Extensions = append(cert.Extensions, ext)
//line /usr/local/go/src/crypto/x509/parser.go:964
					// _ = "end of CoverTab[18933]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:965
				// _ = "end of CoverTab[18926]"
//line /usr/local/go/src/crypto/x509/parser.go:965
				_go_fuzz_dep_.CoverTab[18927]++
										err = processExtensions(cert)
										if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:967
					_go_fuzz_dep_.CoverTab[18940]++
											return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:968
					// _ = "end of CoverTab[18940]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:969
					_go_fuzz_dep_.CoverTab[18941]++
//line /usr/local/go/src/crypto/x509/parser.go:969
					// _ = "end of CoverTab[18941]"
//line /usr/local/go/src/crypto/x509/parser.go:969
				}
//line /usr/local/go/src/crypto/x509/parser.go:969
				// _ = "end of CoverTab[18927]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:970
				_go_fuzz_dep_.CoverTab[18942]++
//line /usr/local/go/src/crypto/x509/parser.go:970
				// _ = "end of CoverTab[18942]"
//line /usr/local/go/src/crypto/x509/parser.go:970
			}
//line /usr/local/go/src/crypto/x509/parser.go:970
			// _ = "end of CoverTab[18922]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:971
			_go_fuzz_dep_.CoverTab[18943]++
//line /usr/local/go/src/crypto/x509/parser.go:971
			// _ = "end of CoverTab[18943]"
//line /usr/local/go/src/crypto/x509/parser.go:971
		}
//line /usr/local/go/src/crypto/x509/parser.go:971
		// _ = "end of CoverTab[18916]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:972
		_go_fuzz_dep_.CoverTab[18944]++
//line /usr/local/go/src/crypto/x509/parser.go:972
		// _ = "end of CoverTab[18944]"
//line /usr/local/go/src/crypto/x509/parser.go:972
	}
//line /usr/local/go/src/crypto/x509/parser.go:972
	// _ = "end of CoverTab[18861]"
//line /usr/local/go/src/crypto/x509/parser.go:972
	_go_fuzz_dep_.CoverTab[18862]++

							var signature asn1.BitString
							if !input.ReadASN1BitString(&signature) {
//line /usr/local/go/src/crypto/x509/parser.go:975
		_go_fuzz_dep_.CoverTab[18945]++
								return nil, errors.New("x509: malformed signature")
//line /usr/local/go/src/crypto/x509/parser.go:976
		// _ = "end of CoverTab[18945]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:977
		_go_fuzz_dep_.CoverTab[18946]++
//line /usr/local/go/src/crypto/x509/parser.go:977
		// _ = "end of CoverTab[18946]"
//line /usr/local/go/src/crypto/x509/parser.go:977
	}
//line /usr/local/go/src/crypto/x509/parser.go:977
	// _ = "end of CoverTab[18862]"
//line /usr/local/go/src/crypto/x509/parser.go:977
	_go_fuzz_dep_.CoverTab[18863]++
							cert.Signature = signature.RightAlign()

							return cert, nil
//line /usr/local/go/src/crypto/x509/parser.go:980
	// _ = "end of CoverTab[18863]"
}

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
func ParseCertificate(der []byte) (*Certificate, error) {
//line /usr/local/go/src/crypto/x509/parser.go:984
	_go_fuzz_dep_.CoverTab[18947]++
							cert, err := parseCertificate(der)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:986
		_go_fuzz_dep_.CoverTab[18950]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:987
		// _ = "end of CoverTab[18950]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:988
		_go_fuzz_dep_.CoverTab[18951]++
//line /usr/local/go/src/crypto/x509/parser.go:988
		// _ = "end of CoverTab[18951]"
//line /usr/local/go/src/crypto/x509/parser.go:988
	}
//line /usr/local/go/src/crypto/x509/parser.go:988
	// _ = "end of CoverTab[18947]"
//line /usr/local/go/src/crypto/x509/parser.go:988
	_go_fuzz_dep_.CoverTab[18948]++
							if len(der) != len(cert.Raw) {
//line /usr/local/go/src/crypto/x509/parser.go:989
		_go_fuzz_dep_.CoverTab[18952]++
								return nil, errors.New("x509: trailing data")
//line /usr/local/go/src/crypto/x509/parser.go:990
		// _ = "end of CoverTab[18952]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:991
		_go_fuzz_dep_.CoverTab[18953]++
//line /usr/local/go/src/crypto/x509/parser.go:991
		// _ = "end of CoverTab[18953]"
//line /usr/local/go/src/crypto/x509/parser.go:991
	}
//line /usr/local/go/src/crypto/x509/parser.go:991
	// _ = "end of CoverTab[18948]"
//line /usr/local/go/src/crypto/x509/parser.go:991
	_go_fuzz_dep_.CoverTab[18949]++
							return cert, err
//line /usr/local/go/src/crypto/x509/parser.go:992
	// _ = "end of CoverTab[18949]"
}

// ParseCertificates parses one or more certificates from the given ASN.1 DER
//line /usr/local/go/src/crypto/x509/parser.go:995
// data. The certificates must be concatenated with no intermediate padding.
//line /usr/local/go/src/crypto/x509/parser.go:997
func ParseCertificates(der []byte) ([]*Certificate, error) {
//line /usr/local/go/src/crypto/x509/parser.go:997
	_go_fuzz_dep_.CoverTab[18954]++
							var certs []*Certificate
							for len(der) > 0 {
//line /usr/local/go/src/crypto/x509/parser.go:999
		_go_fuzz_dep_.CoverTab[18956]++
								cert, err := parseCertificate(der)
								if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1001
			_go_fuzz_dep_.CoverTab[18958]++
									return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1002
			// _ = "end of CoverTab[18958]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:1003
			_go_fuzz_dep_.CoverTab[18959]++
//line /usr/local/go/src/crypto/x509/parser.go:1003
			// _ = "end of CoverTab[18959]"
//line /usr/local/go/src/crypto/x509/parser.go:1003
		}
//line /usr/local/go/src/crypto/x509/parser.go:1003
		// _ = "end of CoverTab[18956]"
//line /usr/local/go/src/crypto/x509/parser.go:1003
		_go_fuzz_dep_.CoverTab[18957]++
								certs = append(certs, cert)
								der = der[len(cert.Raw):]
//line /usr/local/go/src/crypto/x509/parser.go:1005
		// _ = "end of CoverTab[18957]"
	}
//line /usr/local/go/src/crypto/x509/parser.go:1006
	// _ = "end of CoverTab[18954]"
//line /usr/local/go/src/crypto/x509/parser.go:1006
	_go_fuzz_dep_.CoverTab[18955]++
							return certs, nil
//line /usr/local/go/src/crypto/x509/parser.go:1007
	// _ = "end of CoverTab[18955]"
}

// The X.509 standards confusingly 1-indexed the version names, but 0-indexed
//line /usr/local/go/src/crypto/x509/parser.go:1010
// the actual encoded version, so the version for X.509v2 is 1.
//line /usr/local/go/src/crypto/x509/parser.go:1012
const x509v2Version = 1

// ParseRevocationList parses a X509 v2 Certificate Revocation List from the given
//line /usr/local/go/src/crypto/x509/parser.go:1014
// ASN.1 DER data.
//line /usr/local/go/src/crypto/x509/parser.go:1016
func ParseRevocationList(der []byte) (*RevocationList, error) {
//line /usr/local/go/src/crypto/x509/parser.go:1016
	_go_fuzz_dep_.CoverTab[18960]++
							rl := &RevocationList{}

							input := cryptobyte.String(der)

//line /usr/local/go/src/crypto/x509/parser.go:1023
	if !input.ReadASN1Element(&input, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1023
		_go_fuzz_dep_.CoverTab[18980]++
								return nil, errors.New("x509: malformed crl")
//line /usr/local/go/src/crypto/x509/parser.go:1024
		// _ = "end of CoverTab[18980]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1025
		_go_fuzz_dep_.CoverTab[18981]++
//line /usr/local/go/src/crypto/x509/parser.go:1025
		// _ = "end of CoverTab[18981]"
//line /usr/local/go/src/crypto/x509/parser.go:1025
	}
//line /usr/local/go/src/crypto/x509/parser.go:1025
	// _ = "end of CoverTab[18960]"
//line /usr/local/go/src/crypto/x509/parser.go:1025
	_go_fuzz_dep_.CoverTab[18961]++
							rl.Raw = input
							if !input.ReadASN1(&input, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1027
		_go_fuzz_dep_.CoverTab[18982]++
								return nil, errors.New("x509: malformed crl")
//line /usr/local/go/src/crypto/x509/parser.go:1028
		// _ = "end of CoverTab[18982]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1029
		_go_fuzz_dep_.CoverTab[18983]++
//line /usr/local/go/src/crypto/x509/parser.go:1029
		// _ = "end of CoverTab[18983]"
//line /usr/local/go/src/crypto/x509/parser.go:1029
	}
//line /usr/local/go/src/crypto/x509/parser.go:1029
	// _ = "end of CoverTab[18961]"
//line /usr/local/go/src/crypto/x509/parser.go:1029
	_go_fuzz_dep_.CoverTab[18962]++

							var tbs cryptobyte.String

//line /usr/local/go/src/crypto/x509/parser.go:1034
	if !input.ReadASN1Element(&tbs, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1034
		_go_fuzz_dep_.CoverTab[18984]++
								return nil, errors.New("x509: malformed tbs crl")
//line /usr/local/go/src/crypto/x509/parser.go:1035
		// _ = "end of CoverTab[18984]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1036
		_go_fuzz_dep_.CoverTab[18985]++
//line /usr/local/go/src/crypto/x509/parser.go:1036
		// _ = "end of CoverTab[18985]"
//line /usr/local/go/src/crypto/x509/parser.go:1036
	}
//line /usr/local/go/src/crypto/x509/parser.go:1036
	// _ = "end of CoverTab[18962]"
//line /usr/local/go/src/crypto/x509/parser.go:1036
	_go_fuzz_dep_.CoverTab[18963]++
							rl.RawTBSRevocationList = tbs
							if !tbs.ReadASN1(&tbs, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1038
		_go_fuzz_dep_.CoverTab[18986]++
								return nil, errors.New("x509: malformed tbs crl")
//line /usr/local/go/src/crypto/x509/parser.go:1039
		// _ = "end of CoverTab[18986]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1040
		_go_fuzz_dep_.CoverTab[18987]++
//line /usr/local/go/src/crypto/x509/parser.go:1040
		// _ = "end of CoverTab[18987]"
//line /usr/local/go/src/crypto/x509/parser.go:1040
	}
//line /usr/local/go/src/crypto/x509/parser.go:1040
	// _ = "end of CoverTab[18963]"
//line /usr/local/go/src/crypto/x509/parser.go:1040
	_go_fuzz_dep_.CoverTab[18964]++

							var version int
							if !tbs.PeekASN1Tag(cryptobyte_asn1.INTEGER) {
//line /usr/local/go/src/crypto/x509/parser.go:1043
		_go_fuzz_dep_.CoverTab[18988]++
								return nil, errors.New("x509: unsupported crl version")
//line /usr/local/go/src/crypto/x509/parser.go:1044
		// _ = "end of CoverTab[18988]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1045
		_go_fuzz_dep_.CoverTab[18989]++
//line /usr/local/go/src/crypto/x509/parser.go:1045
		// _ = "end of CoverTab[18989]"
//line /usr/local/go/src/crypto/x509/parser.go:1045
	}
//line /usr/local/go/src/crypto/x509/parser.go:1045
	// _ = "end of CoverTab[18964]"
//line /usr/local/go/src/crypto/x509/parser.go:1045
	_go_fuzz_dep_.CoverTab[18965]++
							if !tbs.ReadASN1Integer(&version) {
//line /usr/local/go/src/crypto/x509/parser.go:1046
		_go_fuzz_dep_.CoverTab[18990]++
								return nil, errors.New("x509: malformed crl")
//line /usr/local/go/src/crypto/x509/parser.go:1047
		// _ = "end of CoverTab[18990]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1048
		_go_fuzz_dep_.CoverTab[18991]++
//line /usr/local/go/src/crypto/x509/parser.go:1048
		// _ = "end of CoverTab[18991]"
//line /usr/local/go/src/crypto/x509/parser.go:1048
	}
//line /usr/local/go/src/crypto/x509/parser.go:1048
	// _ = "end of CoverTab[18965]"
//line /usr/local/go/src/crypto/x509/parser.go:1048
	_go_fuzz_dep_.CoverTab[18966]++
							if version != x509v2Version {
//line /usr/local/go/src/crypto/x509/parser.go:1049
		_go_fuzz_dep_.CoverTab[18992]++
								return nil, fmt.Errorf("x509: unsupported crl version: %d", version)
//line /usr/local/go/src/crypto/x509/parser.go:1050
		// _ = "end of CoverTab[18992]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1051
		_go_fuzz_dep_.CoverTab[18993]++
//line /usr/local/go/src/crypto/x509/parser.go:1051
		// _ = "end of CoverTab[18993]"
//line /usr/local/go/src/crypto/x509/parser.go:1051
	}
//line /usr/local/go/src/crypto/x509/parser.go:1051
	// _ = "end of CoverTab[18966]"
//line /usr/local/go/src/crypto/x509/parser.go:1051
	_go_fuzz_dep_.CoverTab[18967]++

							var sigAISeq cryptobyte.String
							if !tbs.ReadASN1(&sigAISeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1054
		_go_fuzz_dep_.CoverTab[18994]++
								return nil, errors.New("x509: malformed signature algorithm identifier")
//line /usr/local/go/src/crypto/x509/parser.go:1055
		// _ = "end of CoverTab[18994]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1056
		_go_fuzz_dep_.CoverTab[18995]++
//line /usr/local/go/src/crypto/x509/parser.go:1056
		// _ = "end of CoverTab[18995]"
//line /usr/local/go/src/crypto/x509/parser.go:1056
	}
//line /usr/local/go/src/crypto/x509/parser.go:1056
	// _ = "end of CoverTab[18967]"
//line /usr/local/go/src/crypto/x509/parser.go:1056
	_go_fuzz_dep_.CoverTab[18968]++
	// Before parsing the inner algorithm identifier, extract
	// the outer algorithm identifier and make sure that they
	// match.
	var outerSigAISeq cryptobyte.String
	if !input.ReadASN1(&outerSigAISeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1061
		_go_fuzz_dep_.CoverTab[18996]++
								return nil, errors.New("x509: malformed algorithm identifier")
//line /usr/local/go/src/crypto/x509/parser.go:1062
		// _ = "end of CoverTab[18996]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1063
		_go_fuzz_dep_.CoverTab[18997]++
//line /usr/local/go/src/crypto/x509/parser.go:1063
		// _ = "end of CoverTab[18997]"
//line /usr/local/go/src/crypto/x509/parser.go:1063
	}
//line /usr/local/go/src/crypto/x509/parser.go:1063
	// _ = "end of CoverTab[18968]"
//line /usr/local/go/src/crypto/x509/parser.go:1063
	_go_fuzz_dep_.CoverTab[18969]++
							if !bytes.Equal(outerSigAISeq, sigAISeq) {
//line /usr/local/go/src/crypto/x509/parser.go:1064
		_go_fuzz_dep_.CoverTab[18998]++
								return nil, errors.New("x509: inner and outer signature algorithm identifiers don't match")
//line /usr/local/go/src/crypto/x509/parser.go:1065
		// _ = "end of CoverTab[18998]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1066
		_go_fuzz_dep_.CoverTab[18999]++
//line /usr/local/go/src/crypto/x509/parser.go:1066
		// _ = "end of CoverTab[18999]"
//line /usr/local/go/src/crypto/x509/parser.go:1066
	}
//line /usr/local/go/src/crypto/x509/parser.go:1066
	// _ = "end of CoverTab[18969]"
//line /usr/local/go/src/crypto/x509/parser.go:1066
	_go_fuzz_dep_.CoverTab[18970]++
							sigAI, err := parseAI(sigAISeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1068
		_go_fuzz_dep_.CoverTab[19000]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1069
		// _ = "end of CoverTab[19000]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1070
		_go_fuzz_dep_.CoverTab[19001]++
//line /usr/local/go/src/crypto/x509/parser.go:1070
		// _ = "end of CoverTab[19001]"
//line /usr/local/go/src/crypto/x509/parser.go:1070
	}
//line /usr/local/go/src/crypto/x509/parser.go:1070
	// _ = "end of CoverTab[18970]"
//line /usr/local/go/src/crypto/x509/parser.go:1070
	_go_fuzz_dep_.CoverTab[18971]++
							rl.SignatureAlgorithm = getSignatureAlgorithmFromAI(sigAI)

							var signature asn1.BitString
							if !input.ReadASN1BitString(&signature) {
//line /usr/local/go/src/crypto/x509/parser.go:1074
		_go_fuzz_dep_.CoverTab[19002]++
								return nil, errors.New("x509: malformed signature")
//line /usr/local/go/src/crypto/x509/parser.go:1075
		// _ = "end of CoverTab[19002]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1076
		_go_fuzz_dep_.CoverTab[19003]++
//line /usr/local/go/src/crypto/x509/parser.go:1076
		// _ = "end of CoverTab[19003]"
//line /usr/local/go/src/crypto/x509/parser.go:1076
	}
//line /usr/local/go/src/crypto/x509/parser.go:1076
	// _ = "end of CoverTab[18971]"
//line /usr/local/go/src/crypto/x509/parser.go:1076
	_go_fuzz_dep_.CoverTab[18972]++
							rl.Signature = signature.RightAlign()

							var issuerSeq cryptobyte.String
							if !tbs.ReadASN1Element(&issuerSeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1080
		_go_fuzz_dep_.CoverTab[19004]++
								return nil, errors.New("x509: malformed issuer")
//line /usr/local/go/src/crypto/x509/parser.go:1081
		// _ = "end of CoverTab[19004]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1082
		_go_fuzz_dep_.CoverTab[19005]++
//line /usr/local/go/src/crypto/x509/parser.go:1082
		// _ = "end of CoverTab[19005]"
//line /usr/local/go/src/crypto/x509/parser.go:1082
	}
//line /usr/local/go/src/crypto/x509/parser.go:1082
	// _ = "end of CoverTab[18972]"
//line /usr/local/go/src/crypto/x509/parser.go:1082
	_go_fuzz_dep_.CoverTab[18973]++
							rl.RawIssuer = issuerSeq
							issuerRDNs, err := parseName(issuerSeq)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1085
		_go_fuzz_dep_.CoverTab[19006]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1086
		// _ = "end of CoverTab[19006]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1087
		_go_fuzz_dep_.CoverTab[19007]++
//line /usr/local/go/src/crypto/x509/parser.go:1087
		// _ = "end of CoverTab[19007]"
//line /usr/local/go/src/crypto/x509/parser.go:1087
	}
//line /usr/local/go/src/crypto/x509/parser.go:1087
	// _ = "end of CoverTab[18973]"
//line /usr/local/go/src/crypto/x509/parser.go:1087
	_go_fuzz_dep_.CoverTab[18974]++
							rl.Issuer.FillFromRDNSequence(issuerRDNs)

							rl.ThisUpdate, err = parseTime(&tbs)
							if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1091
		_go_fuzz_dep_.CoverTab[19008]++
								return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1092
		// _ = "end of CoverTab[19008]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1093
		_go_fuzz_dep_.CoverTab[19009]++
//line /usr/local/go/src/crypto/x509/parser.go:1093
		// _ = "end of CoverTab[19009]"
//line /usr/local/go/src/crypto/x509/parser.go:1093
	}
//line /usr/local/go/src/crypto/x509/parser.go:1093
	// _ = "end of CoverTab[18974]"
//line /usr/local/go/src/crypto/x509/parser.go:1093
	_go_fuzz_dep_.CoverTab[18975]++
							if tbs.PeekASN1Tag(cryptobyte_asn1.GeneralizedTime) || func() bool {
//line /usr/local/go/src/crypto/x509/parser.go:1094
		_go_fuzz_dep_.CoverTab[19010]++
//line /usr/local/go/src/crypto/x509/parser.go:1094
		return tbs.PeekASN1Tag(cryptobyte_asn1.UTCTime)
//line /usr/local/go/src/crypto/x509/parser.go:1094
		// _ = "end of CoverTab[19010]"
//line /usr/local/go/src/crypto/x509/parser.go:1094
	}() {
//line /usr/local/go/src/crypto/x509/parser.go:1094
		_go_fuzz_dep_.CoverTab[19011]++
								rl.NextUpdate, err = parseTime(&tbs)
								if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1096
			_go_fuzz_dep_.CoverTab[19012]++
									return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1097
			// _ = "end of CoverTab[19012]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:1098
			_go_fuzz_dep_.CoverTab[19013]++
//line /usr/local/go/src/crypto/x509/parser.go:1098
			// _ = "end of CoverTab[19013]"
//line /usr/local/go/src/crypto/x509/parser.go:1098
		}
//line /usr/local/go/src/crypto/x509/parser.go:1098
		// _ = "end of CoverTab[19011]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1099
		_go_fuzz_dep_.CoverTab[19014]++
//line /usr/local/go/src/crypto/x509/parser.go:1099
		// _ = "end of CoverTab[19014]"
//line /usr/local/go/src/crypto/x509/parser.go:1099
	}
//line /usr/local/go/src/crypto/x509/parser.go:1099
	// _ = "end of CoverTab[18975]"
//line /usr/local/go/src/crypto/x509/parser.go:1099
	_go_fuzz_dep_.CoverTab[18976]++

							if tbs.PeekASN1Tag(cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1101
		_go_fuzz_dep_.CoverTab[19015]++
								var revokedSeq cryptobyte.String
								if !tbs.ReadASN1(&revokedSeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1103
			_go_fuzz_dep_.CoverTab[19017]++
									return nil, errors.New("x509: malformed crl")
//line /usr/local/go/src/crypto/x509/parser.go:1104
			// _ = "end of CoverTab[19017]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:1105
			_go_fuzz_dep_.CoverTab[19018]++
//line /usr/local/go/src/crypto/x509/parser.go:1105
			// _ = "end of CoverTab[19018]"
//line /usr/local/go/src/crypto/x509/parser.go:1105
		}
//line /usr/local/go/src/crypto/x509/parser.go:1105
		// _ = "end of CoverTab[19015]"
//line /usr/local/go/src/crypto/x509/parser.go:1105
		_go_fuzz_dep_.CoverTab[19016]++
								for !revokedSeq.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:1106
			_go_fuzz_dep_.CoverTab[19019]++
									var certSeq cryptobyte.String
									if !revokedSeq.ReadASN1(&certSeq, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1108
				_go_fuzz_dep_.CoverTab[19025]++
										return nil, errors.New("x509: malformed crl")
//line /usr/local/go/src/crypto/x509/parser.go:1109
				// _ = "end of CoverTab[19025]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1110
				_go_fuzz_dep_.CoverTab[19026]++
//line /usr/local/go/src/crypto/x509/parser.go:1110
				// _ = "end of CoverTab[19026]"
//line /usr/local/go/src/crypto/x509/parser.go:1110
			}
//line /usr/local/go/src/crypto/x509/parser.go:1110
			// _ = "end of CoverTab[19019]"
//line /usr/local/go/src/crypto/x509/parser.go:1110
			_go_fuzz_dep_.CoverTab[19020]++
									rc := pkix.RevokedCertificate{}
									rc.SerialNumber = new(big.Int)
									if !certSeq.ReadASN1Integer(rc.SerialNumber) {
//line /usr/local/go/src/crypto/x509/parser.go:1113
				_go_fuzz_dep_.CoverTab[19027]++
										return nil, errors.New("x509: malformed serial number")
//line /usr/local/go/src/crypto/x509/parser.go:1114
				// _ = "end of CoverTab[19027]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1115
				_go_fuzz_dep_.CoverTab[19028]++
//line /usr/local/go/src/crypto/x509/parser.go:1115
				// _ = "end of CoverTab[19028]"
//line /usr/local/go/src/crypto/x509/parser.go:1115
			}
//line /usr/local/go/src/crypto/x509/parser.go:1115
			// _ = "end of CoverTab[19020]"
//line /usr/local/go/src/crypto/x509/parser.go:1115
			_go_fuzz_dep_.CoverTab[19021]++
									rc.RevocationTime, err = parseTime(&certSeq)
									if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1117
				_go_fuzz_dep_.CoverTab[19029]++
										return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1118
				// _ = "end of CoverTab[19029]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1119
				_go_fuzz_dep_.CoverTab[19030]++
//line /usr/local/go/src/crypto/x509/parser.go:1119
				// _ = "end of CoverTab[19030]"
//line /usr/local/go/src/crypto/x509/parser.go:1119
			}
//line /usr/local/go/src/crypto/x509/parser.go:1119
			// _ = "end of CoverTab[19021]"
//line /usr/local/go/src/crypto/x509/parser.go:1119
			_go_fuzz_dep_.CoverTab[19022]++
									var extensions cryptobyte.String
									var present bool
									if !certSeq.ReadOptionalASN1(&extensions, &present, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1122
				_go_fuzz_dep_.CoverTab[19031]++
										return nil, errors.New("x509: malformed extensions")
//line /usr/local/go/src/crypto/x509/parser.go:1123
				// _ = "end of CoverTab[19031]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1124
				_go_fuzz_dep_.CoverTab[19032]++
//line /usr/local/go/src/crypto/x509/parser.go:1124
				// _ = "end of CoverTab[19032]"
//line /usr/local/go/src/crypto/x509/parser.go:1124
			}
//line /usr/local/go/src/crypto/x509/parser.go:1124
			// _ = "end of CoverTab[19022]"
//line /usr/local/go/src/crypto/x509/parser.go:1124
			_go_fuzz_dep_.CoverTab[19023]++
									if present {
//line /usr/local/go/src/crypto/x509/parser.go:1125
				_go_fuzz_dep_.CoverTab[19033]++
										for !extensions.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:1126
					_go_fuzz_dep_.CoverTab[19034]++
											var extension cryptobyte.String
											if !extensions.ReadASN1(&extension, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1128
						_go_fuzz_dep_.CoverTab[19037]++
												return nil, errors.New("x509: malformed extension")
//line /usr/local/go/src/crypto/x509/parser.go:1129
						// _ = "end of CoverTab[19037]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:1130
						_go_fuzz_dep_.CoverTab[19038]++
//line /usr/local/go/src/crypto/x509/parser.go:1130
						// _ = "end of CoverTab[19038]"
//line /usr/local/go/src/crypto/x509/parser.go:1130
					}
//line /usr/local/go/src/crypto/x509/parser.go:1130
					// _ = "end of CoverTab[19034]"
//line /usr/local/go/src/crypto/x509/parser.go:1130
					_go_fuzz_dep_.CoverTab[19035]++
											ext, err := parseExtension(extension)
											if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1132
						_go_fuzz_dep_.CoverTab[19039]++
												return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1133
						// _ = "end of CoverTab[19039]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:1134
						_go_fuzz_dep_.CoverTab[19040]++
//line /usr/local/go/src/crypto/x509/parser.go:1134
						// _ = "end of CoverTab[19040]"
//line /usr/local/go/src/crypto/x509/parser.go:1134
					}
//line /usr/local/go/src/crypto/x509/parser.go:1134
					// _ = "end of CoverTab[19035]"
//line /usr/local/go/src/crypto/x509/parser.go:1134
					_go_fuzz_dep_.CoverTab[19036]++
											rc.Extensions = append(rc.Extensions, ext)
//line /usr/local/go/src/crypto/x509/parser.go:1135
					// _ = "end of CoverTab[19036]"
				}
//line /usr/local/go/src/crypto/x509/parser.go:1136
				// _ = "end of CoverTab[19033]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1137
				_go_fuzz_dep_.CoverTab[19041]++
//line /usr/local/go/src/crypto/x509/parser.go:1137
				// _ = "end of CoverTab[19041]"
//line /usr/local/go/src/crypto/x509/parser.go:1137
			}
//line /usr/local/go/src/crypto/x509/parser.go:1137
			// _ = "end of CoverTab[19023]"
//line /usr/local/go/src/crypto/x509/parser.go:1137
			_go_fuzz_dep_.CoverTab[19024]++

									rl.RevokedCertificates = append(rl.RevokedCertificates, rc)
//line /usr/local/go/src/crypto/x509/parser.go:1139
			// _ = "end of CoverTab[19024]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:1140
		// _ = "end of CoverTab[19016]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1141
		_go_fuzz_dep_.CoverTab[19042]++
//line /usr/local/go/src/crypto/x509/parser.go:1141
		// _ = "end of CoverTab[19042]"
//line /usr/local/go/src/crypto/x509/parser.go:1141
	}
//line /usr/local/go/src/crypto/x509/parser.go:1141
	// _ = "end of CoverTab[18976]"
//line /usr/local/go/src/crypto/x509/parser.go:1141
	_go_fuzz_dep_.CoverTab[18977]++

							var extensions cryptobyte.String
							var present bool
							if !tbs.ReadOptionalASN1(&extensions, &present, cryptobyte_asn1.Tag(0).Constructed().ContextSpecific()) {
//line /usr/local/go/src/crypto/x509/parser.go:1145
		_go_fuzz_dep_.CoverTab[19043]++
								return nil, errors.New("x509: malformed extensions")
//line /usr/local/go/src/crypto/x509/parser.go:1146
		// _ = "end of CoverTab[19043]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1147
		_go_fuzz_dep_.CoverTab[19044]++
//line /usr/local/go/src/crypto/x509/parser.go:1147
		// _ = "end of CoverTab[19044]"
//line /usr/local/go/src/crypto/x509/parser.go:1147
	}
//line /usr/local/go/src/crypto/x509/parser.go:1147
	// _ = "end of CoverTab[18977]"
//line /usr/local/go/src/crypto/x509/parser.go:1147
	_go_fuzz_dep_.CoverTab[18978]++
							if present {
//line /usr/local/go/src/crypto/x509/parser.go:1148
		_go_fuzz_dep_.CoverTab[19045]++
								if !extensions.ReadASN1(&extensions, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1149
			_go_fuzz_dep_.CoverTab[19047]++
									return nil, errors.New("x509: malformed extensions")
//line /usr/local/go/src/crypto/x509/parser.go:1150
			// _ = "end of CoverTab[19047]"
		} else {
//line /usr/local/go/src/crypto/x509/parser.go:1151
			_go_fuzz_dep_.CoverTab[19048]++
//line /usr/local/go/src/crypto/x509/parser.go:1151
			// _ = "end of CoverTab[19048]"
//line /usr/local/go/src/crypto/x509/parser.go:1151
		}
//line /usr/local/go/src/crypto/x509/parser.go:1151
		// _ = "end of CoverTab[19045]"
//line /usr/local/go/src/crypto/x509/parser.go:1151
		_go_fuzz_dep_.CoverTab[19046]++
								for !extensions.Empty() {
//line /usr/local/go/src/crypto/x509/parser.go:1152
			_go_fuzz_dep_.CoverTab[19049]++
									var extension cryptobyte.String
									if !extensions.ReadASN1(&extension, cryptobyte_asn1.SEQUENCE) {
//line /usr/local/go/src/crypto/x509/parser.go:1154
				_go_fuzz_dep_.CoverTab[19053]++
										return nil, errors.New("x509: malformed extension")
//line /usr/local/go/src/crypto/x509/parser.go:1155
				// _ = "end of CoverTab[19053]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1156
				_go_fuzz_dep_.CoverTab[19054]++
//line /usr/local/go/src/crypto/x509/parser.go:1156
				// _ = "end of CoverTab[19054]"
//line /usr/local/go/src/crypto/x509/parser.go:1156
			}
//line /usr/local/go/src/crypto/x509/parser.go:1156
			// _ = "end of CoverTab[19049]"
//line /usr/local/go/src/crypto/x509/parser.go:1156
			_go_fuzz_dep_.CoverTab[19050]++
									ext, err := parseExtension(extension)
									if err != nil {
//line /usr/local/go/src/crypto/x509/parser.go:1158
				_go_fuzz_dep_.CoverTab[19055]++
										return nil, err
//line /usr/local/go/src/crypto/x509/parser.go:1159
				// _ = "end of CoverTab[19055]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1160
				_go_fuzz_dep_.CoverTab[19056]++
//line /usr/local/go/src/crypto/x509/parser.go:1160
				// _ = "end of CoverTab[19056]"
//line /usr/local/go/src/crypto/x509/parser.go:1160
			}
//line /usr/local/go/src/crypto/x509/parser.go:1160
			// _ = "end of CoverTab[19050]"
//line /usr/local/go/src/crypto/x509/parser.go:1160
			_go_fuzz_dep_.CoverTab[19051]++
									if ext.Id.Equal(oidExtensionAuthorityKeyId) {
//line /usr/local/go/src/crypto/x509/parser.go:1161
				_go_fuzz_dep_.CoverTab[19057]++
										rl.AuthorityKeyId = ext.Value
//line /usr/local/go/src/crypto/x509/parser.go:1162
				// _ = "end of CoverTab[19057]"
			} else {
//line /usr/local/go/src/crypto/x509/parser.go:1163
				_go_fuzz_dep_.CoverTab[19058]++
//line /usr/local/go/src/crypto/x509/parser.go:1163
				if ext.Id.Equal(oidExtensionCRLNumber) {
//line /usr/local/go/src/crypto/x509/parser.go:1163
					_go_fuzz_dep_.CoverTab[19059]++
											value := cryptobyte.String(ext.Value)
											rl.Number = new(big.Int)
											if !value.ReadASN1Integer(rl.Number) {
//line /usr/local/go/src/crypto/x509/parser.go:1166
						_go_fuzz_dep_.CoverTab[19060]++
												return nil, errors.New("x509: malformed crl number")
//line /usr/local/go/src/crypto/x509/parser.go:1167
						// _ = "end of CoverTab[19060]"
					} else {
//line /usr/local/go/src/crypto/x509/parser.go:1168
						_go_fuzz_dep_.CoverTab[19061]++
//line /usr/local/go/src/crypto/x509/parser.go:1168
						// _ = "end of CoverTab[19061]"
//line /usr/local/go/src/crypto/x509/parser.go:1168
					}
//line /usr/local/go/src/crypto/x509/parser.go:1168
					// _ = "end of CoverTab[19059]"
				} else {
//line /usr/local/go/src/crypto/x509/parser.go:1169
					_go_fuzz_dep_.CoverTab[19062]++
//line /usr/local/go/src/crypto/x509/parser.go:1169
					// _ = "end of CoverTab[19062]"
//line /usr/local/go/src/crypto/x509/parser.go:1169
				}
//line /usr/local/go/src/crypto/x509/parser.go:1169
				// _ = "end of CoverTab[19058]"
//line /usr/local/go/src/crypto/x509/parser.go:1169
			}
//line /usr/local/go/src/crypto/x509/parser.go:1169
			// _ = "end of CoverTab[19051]"
//line /usr/local/go/src/crypto/x509/parser.go:1169
			_go_fuzz_dep_.CoverTab[19052]++
									rl.Extensions = append(rl.Extensions, ext)
//line /usr/local/go/src/crypto/x509/parser.go:1170
			// _ = "end of CoverTab[19052]"
		}
//line /usr/local/go/src/crypto/x509/parser.go:1171
		// _ = "end of CoverTab[19046]"
	} else {
//line /usr/local/go/src/crypto/x509/parser.go:1172
		_go_fuzz_dep_.CoverTab[19063]++
//line /usr/local/go/src/crypto/x509/parser.go:1172
		// _ = "end of CoverTab[19063]"
//line /usr/local/go/src/crypto/x509/parser.go:1172
	}
//line /usr/local/go/src/crypto/x509/parser.go:1172
	// _ = "end of CoverTab[18978]"
//line /usr/local/go/src/crypto/x509/parser.go:1172
	_go_fuzz_dep_.CoverTab[18979]++

							return rl, nil
//line /usr/local/go/src/crypto/x509/parser.go:1174
	// _ = "end of CoverTab[18979]"
}

//line /usr/local/go/src/crypto/x509/parser.go:1175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/parser.go:1175
var _ = _go_fuzz_dep_.CoverTab
