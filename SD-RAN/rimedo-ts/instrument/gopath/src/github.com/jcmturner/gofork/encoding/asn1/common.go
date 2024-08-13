// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
package asn1

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:5
)

import (
	"reflect"
	"strconv"
	"strings"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:21
// ASN.1 tags represent the type of the following object.
const (
	TagBoolean		= 1
	TagInteger		= 2
	TagBitString		= 3
	TagOctetString		= 4
	TagOID			= 6
	TagEnum			= 10
	TagUTF8String		= 12
	TagSequence		= 16
	TagSet			= 17
	TagPrintableString	= 19
	TagT61String		= 20
	TagIA5String		= 22
	TagUTCTime		= 23
	TagGeneralizedTime	= 24
	TagGeneralString	= 27
)

// ASN.1 class types represent the namespace of the tag.
const (
	ClassUniversal		= 0
	ClassApplication	= 1
	ClassContextSpecific	= 2
	ClassPrivate		= 3
)

type tagAndLength struct {
	class, tag, length	int
	isCompound		bool
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:71
// fieldParameters is the parsed representation of tag string from a structure field.
type fieldParameters struct {
	optional	bool	// true iff the field is OPTIONAL
	explicit	bool	// true iff an EXPLICIT tag is in use.
	application	bool	// true iff an APPLICATION tag is in use.
	defaultValue	*int64	// a default value for INTEGER typed fields (maybe nil).
	tag		*int	// the EXPLICIT or IMPLICIT tag (maybe nil).
	stringType	int	// the string tag to use when marshaling.
	timeType	int	// the time tag to use when marshaling.
	set		bool	// true iff this should be encoded as a SET
	omitEmpty	bool	// true iff this should be omitted if empty when marshaling.

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:85
}

// Given a tag string with the format specified in the package comment,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:87
// parseFieldParameters will parse it into a fieldParameters structure,
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:87
// ignoring unknown parts of the string.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:90
func parseFieldParameters(str string) (ret fieldParameters) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:90
	_go_fuzz_dep_.CoverTab[82721]++
													for _, part := range strings.Split(str, ",") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:91
		_go_fuzz_dep_.CoverTab[82723]++
														switch {
		case part == "optional":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:93
			_go_fuzz_dep_.CoverTab[82724]++
															ret.optional = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:94
			// _ = "end of CoverTab[82724]"
		case part == "explicit":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:95
			_go_fuzz_dep_.CoverTab[82725]++
															ret.explicit = true
															if ret.tag == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:97
				_go_fuzz_dep_.CoverTab[82738]++
																ret.tag = new(int)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:98
				// _ = "end of CoverTab[82738]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:99
				_go_fuzz_dep_.CoverTab[82739]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:99
				// _ = "end of CoverTab[82739]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:99
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:99
			// _ = "end of CoverTab[82725]"
		case part == "generalized":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:100
			_go_fuzz_dep_.CoverTab[82726]++
															ret.timeType = TagGeneralizedTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:101
			// _ = "end of CoverTab[82726]"
		case part == "utc":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:102
			_go_fuzz_dep_.CoverTab[82727]++
															ret.timeType = TagUTCTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:103
			// _ = "end of CoverTab[82727]"
		case part == "ia5":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:104
			_go_fuzz_dep_.CoverTab[82728]++
															ret.stringType = TagIA5String
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:105
			// _ = "end of CoverTab[82728]"

		case part == "generalstring":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:107
			_go_fuzz_dep_.CoverTab[82729]++
															ret.stringType = TagGeneralString
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:108
			// _ = "end of CoverTab[82729]"
		case part == "printable":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:109
			_go_fuzz_dep_.CoverTab[82730]++
															ret.stringType = TagPrintableString
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:110
			// _ = "end of CoverTab[82730]"
		case part == "utf8":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:111
			_go_fuzz_dep_.CoverTab[82731]++
															ret.stringType = TagUTF8String
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:112
			// _ = "end of CoverTab[82731]"
		case strings.HasPrefix(part, "default:"):
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:113
			_go_fuzz_dep_.CoverTab[82732]++
															i, err := strconv.ParseInt(part[8:], 10, 64)
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:115
				_go_fuzz_dep_.CoverTab[82740]++
																ret.defaultValue = new(int64)
																*ret.defaultValue = i
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:117
				// _ = "end of CoverTab[82740]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:118
				_go_fuzz_dep_.CoverTab[82741]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:118
				// _ = "end of CoverTab[82741]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:118
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:118
			// _ = "end of CoverTab[82732]"
		case strings.HasPrefix(part, "tag:"):
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:119
			_go_fuzz_dep_.CoverTab[82733]++
															i, err := strconv.Atoi(part[4:])
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:121
				_go_fuzz_dep_.CoverTab[82742]++
																ret.tag = new(int)
																*ret.tag = i
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:123
				// _ = "end of CoverTab[82742]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:124
				_go_fuzz_dep_.CoverTab[82743]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:124
				// _ = "end of CoverTab[82743]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:124
			// _ = "end of CoverTab[82733]"
		case part == "set":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:125
			_go_fuzz_dep_.CoverTab[82734]++
															ret.set = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:126
			// _ = "end of CoverTab[82734]"
		case part == "application":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:127
			_go_fuzz_dep_.CoverTab[82735]++
															ret.application = true
															if ret.tag == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:129
				_go_fuzz_dep_.CoverTab[82744]++
																ret.tag = new(int)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:130
				// _ = "end of CoverTab[82744]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:131
				_go_fuzz_dep_.CoverTab[82745]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:131
				// _ = "end of CoverTab[82745]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:131
			// _ = "end of CoverTab[82735]"
		case part == "omitempty":
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:132
			_go_fuzz_dep_.CoverTab[82736]++
															ret.omitEmpty = true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:133
			// _ = "end of CoverTab[82736]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:133
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:133
			_go_fuzz_dep_.CoverTab[82737]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:133
			// _ = "end of CoverTab[82737]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:134
		// _ = "end of CoverTab[82723]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:135
	// _ = "end of CoverTab[82721]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:135
	_go_fuzz_dep_.CoverTab[82722]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:136
	// _ = "end of CoverTab[82722]"
}

// Given a reflected Go type, getUniversalType returns the default tag number
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:139
// and expected compound flag.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:141
func getUniversalType(t reflect.Type) (tagNumber int, isCompound, ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:141
	_go_fuzz_dep_.CoverTab[82746]++
													switch t {
	case objectIdentifierType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:143
		_go_fuzz_dep_.CoverTab[82749]++
														return TagOID, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:144
		// _ = "end of CoverTab[82749]"
	case bitStringType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:145
		_go_fuzz_dep_.CoverTab[82750]++
														return TagBitString, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:146
		// _ = "end of CoverTab[82750]"
	case timeType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:147
		_go_fuzz_dep_.CoverTab[82751]++
														return TagUTCTime, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:148
		// _ = "end of CoverTab[82751]"
	case enumeratedType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:149
		_go_fuzz_dep_.CoverTab[82752]++
														return TagEnum, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:150
		// _ = "end of CoverTab[82752]"
	case bigIntType:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:151
		_go_fuzz_dep_.CoverTab[82753]++
														return TagInteger, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:152
		// _ = "end of CoverTab[82753]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:152
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:152
		_go_fuzz_dep_.CoverTab[82754]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:152
		// _ = "end of CoverTab[82754]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:153
	// _ = "end of CoverTab[82746]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:153
	_go_fuzz_dep_.CoverTab[82747]++
													switch t.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:155
		_go_fuzz_dep_.CoverTab[82755]++
														return TagBoolean, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:156
		// _ = "end of CoverTab[82755]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:157
		_go_fuzz_dep_.CoverTab[82756]++
														return TagInteger, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:158
		// _ = "end of CoverTab[82756]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:159
		_go_fuzz_dep_.CoverTab[82757]++
														return TagSequence, true, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:160
		// _ = "end of CoverTab[82757]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:161
		_go_fuzz_dep_.CoverTab[82758]++
														if t.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:162
			_go_fuzz_dep_.CoverTab[82763]++
															return TagOctetString, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:163
			// _ = "end of CoverTab[82763]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:164
			_go_fuzz_dep_.CoverTab[82764]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:164
			// _ = "end of CoverTab[82764]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:164
		// _ = "end of CoverTab[82758]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:164
		_go_fuzz_dep_.CoverTab[82759]++
														if strings.HasSuffix(t.Name(), "SET") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:165
			_go_fuzz_dep_.CoverTab[82765]++
															return TagSet, true, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:166
			// _ = "end of CoverTab[82765]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:167
			_go_fuzz_dep_.CoverTab[82766]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:167
			// _ = "end of CoverTab[82766]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:167
		// _ = "end of CoverTab[82759]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:167
		_go_fuzz_dep_.CoverTab[82760]++
														return TagSequence, true, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:168
		// _ = "end of CoverTab[82760]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:169
		_go_fuzz_dep_.CoverTab[82761]++
														return TagPrintableString, false, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:170
		// _ = "end of CoverTab[82761]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:170
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:170
		_go_fuzz_dep_.CoverTab[82762]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:170
		// _ = "end of CoverTab[82762]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:171
	// _ = "end of CoverTab[82747]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:171
	_go_fuzz_dep_.CoverTab[82748]++
													return 0, false, false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:172
	// _ = "end of CoverTab[82748]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:173
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/encoding/asn1/common.go:173
var _ = _go_fuzz_dep_.CoverTab
