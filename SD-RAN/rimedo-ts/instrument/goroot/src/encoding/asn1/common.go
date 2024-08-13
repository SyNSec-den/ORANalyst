// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/asn1/common.go:5
package asn1

//line /usr/local/go/src/encoding/asn1/common.go:5
import (
//line /usr/local/go/src/encoding/asn1/common.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/asn1/common.go:5
)
//line /usr/local/go/src/encoding/asn1/common.go:5
import (
//line /usr/local/go/src/encoding/asn1/common.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/asn1/common.go:5
)

import (
	"reflect"
	"strconv"
	"strings"
)

//line /usr/local/go/src/encoding/asn1/common.go:21
// ASN.1 tags represent the type of the following object.
const (
	TagBoolean		= 1
	TagInteger		= 2
	TagBitString		= 3
	TagOctetString		= 4
	TagNull			= 5
	TagOID			= 6
	TagEnum			= 10
	TagUTF8String		= 12
	TagSequence		= 16
	TagSet			= 17
	TagNumericString	= 18
	TagPrintableString	= 19
	TagT61String		= 20
	TagIA5String		= 22
	TagUTCTime		= 23
	TagGeneralizedTime	= 24
	TagGeneralString	= 27
	TagBMPString		= 30
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

//line /usr/local/go/src/encoding/asn1/common.go:74
// fieldParameters is the parsed representation of tag string from a structure field.
type fieldParameters struct {
	optional	bool	// true iff the field is OPTIONAL
	explicit	bool	// true iff an EXPLICIT tag is in use.
	application	bool	// true iff an APPLICATION tag is in use.
	private		bool	// true iff a PRIVATE tag is in use.
	defaultValue	*int64	// a default value for INTEGER typed fields (maybe nil).
	tag		*int	// the EXPLICIT or IMPLICIT tag (maybe nil).
	stringType	int	// the string tag to use when marshaling.
	timeType	int	// the time tag to use when marshaling.
	set		bool	// true iff this should be encoded as a SET
	omitEmpty	bool	// true iff this should be omitted if empty when marshaling.

//line /usr/local/go/src/encoding/asn1/common.go:89
}

// Given a tag string with the format specified in the package comment,
//line /usr/local/go/src/encoding/asn1/common.go:91
// parseFieldParameters will parse it into a fieldParameters structure,
//line /usr/local/go/src/encoding/asn1/common.go:91
// ignoring unknown parts of the string.
//line /usr/local/go/src/encoding/asn1/common.go:94
func parseFieldParameters(str string) (ret fieldParameters) {
//line /usr/local/go/src/encoding/asn1/common.go:94
	_go_fuzz_dep_.CoverTab[7953]++
							var part string
							for len(str) > 0 {
//line /usr/local/go/src/encoding/asn1/common.go:96
		_go_fuzz_dep_.CoverTab[7955]++
								part, str, _ = strings.Cut(str, ",")
								switch {
		case part == "optional":
//line /usr/local/go/src/encoding/asn1/common.go:99
			_go_fuzz_dep_.CoverTab[7956]++
									ret.optional = true
//line /usr/local/go/src/encoding/asn1/common.go:100
			// _ = "end of CoverTab[7956]"
		case part == "explicit":
//line /usr/local/go/src/encoding/asn1/common.go:101
			_go_fuzz_dep_.CoverTab[7957]++
									ret.explicit = true
									if ret.tag == nil {
//line /usr/local/go/src/encoding/asn1/common.go:103
				_go_fuzz_dep_.CoverTab[7971]++
										ret.tag = new(int)
//line /usr/local/go/src/encoding/asn1/common.go:104
				// _ = "end of CoverTab[7971]"
			} else {
//line /usr/local/go/src/encoding/asn1/common.go:105
				_go_fuzz_dep_.CoverTab[7972]++
//line /usr/local/go/src/encoding/asn1/common.go:105
				// _ = "end of CoverTab[7972]"
//line /usr/local/go/src/encoding/asn1/common.go:105
			}
//line /usr/local/go/src/encoding/asn1/common.go:105
			// _ = "end of CoverTab[7957]"
		case part == "generalized":
//line /usr/local/go/src/encoding/asn1/common.go:106
			_go_fuzz_dep_.CoverTab[7958]++
									ret.timeType = TagGeneralizedTime
//line /usr/local/go/src/encoding/asn1/common.go:107
			// _ = "end of CoverTab[7958]"
		case part == "utc":
//line /usr/local/go/src/encoding/asn1/common.go:108
			_go_fuzz_dep_.CoverTab[7959]++
									ret.timeType = TagUTCTime
//line /usr/local/go/src/encoding/asn1/common.go:109
			// _ = "end of CoverTab[7959]"
		case part == "ia5":
//line /usr/local/go/src/encoding/asn1/common.go:110
			_go_fuzz_dep_.CoverTab[7960]++
									ret.stringType = TagIA5String
//line /usr/local/go/src/encoding/asn1/common.go:111
			// _ = "end of CoverTab[7960]"
		case part == "printable":
//line /usr/local/go/src/encoding/asn1/common.go:112
			_go_fuzz_dep_.CoverTab[7961]++
									ret.stringType = TagPrintableString
//line /usr/local/go/src/encoding/asn1/common.go:113
			// _ = "end of CoverTab[7961]"
		case part == "numeric":
//line /usr/local/go/src/encoding/asn1/common.go:114
			_go_fuzz_dep_.CoverTab[7962]++
									ret.stringType = TagNumericString
//line /usr/local/go/src/encoding/asn1/common.go:115
			// _ = "end of CoverTab[7962]"
		case part == "utf8":
//line /usr/local/go/src/encoding/asn1/common.go:116
			_go_fuzz_dep_.CoverTab[7963]++
									ret.stringType = TagUTF8String
//line /usr/local/go/src/encoding/asn1/common.go:117
			// _ = "end of CoverTab[7963]"
		case strings.HasPrefix(part, "default:"):
//line /usr/local/go/src/encoding/asn1/common.go:118
			_go_fuzz_dep_.CoverTab[7964]++
									i, err := strconv.ParseInt(part[8:], 10, 64)
									if err == nil {
//line /usr/local/go/src/encoding/asn1/common.go:120
				_go_fuzz_dep_.CoverTab[7973]++
										ret.defaultValue = new(int64)
										*ret.defaultValue = i
//line /usr/local/go/src/encoding/asn1/common.go:122
				// _ = "end of CoverTab[7973]"
			} else {
//line /usr/local/go/src/encoding/asn1/common.go:123
				_go_fuzz_dep_.CoverTab[7974]++
//line /usr/local/go/src/encoding/asn1/common.go:123
				// _ = "end of CoverTab[7974]"
//line /usr/local/go/src/encoding/asn1/common.go:123
			}
//line /usr/local/go/src/encoding/asn1/common.go:123
			// _ = "end of CoverTab[7964]"
		case strings.HasPrefix(part, "tag:"):
//line /usr/local/go/src/encoding/asn1/common.go:124
			_go_fuzz_dep_.CoverTab[7965]++
									i, err := strconv.Atoi(part[4:])
									if err == nil {
//line /usr/local/go/src/encoding/asn1/common.go:126
				_go_fuzz_dep_.CoverTab[7975]++
										ret.tag = new(int)
										*ret.tag = i
//line /usr/local/go/src/encoding/asn1/common.go:128
				// _ = "end of CoverTab[7975]"
			} else {
//line /usr/local/go/src/encoding/asn1/common.go:129
				_go_fuzz_dep_.CoverTab[7976]++
//line /usr/local/go/src/encoding/asn1/common.go:129
				// _ = "end of CoverTab[7976]"
//line /usr/local/go/src/encoding/asn1/common.go:129
			}
//line /usr/local/go/src/encoding/asn1/common.go:129
			// _ = "end of CoverTab[7965]"
		case part == "set":
//line /usr/local/go/src/encoding/asn1/common.go:130
			_go_fuzz_dep_.CoverTab[7966]++
									ret.set = true
//line /usr/local/go/src/encoding/asn1/common.go:131
			// _ = "end of CoverTab[7966]"
		case part == "application":
//line /usr/local/go/src/encoding/asn1/common.go:132
			_go_fuzz_dep_.CoverTab[7967]++
									ret.application = true
									if ret.tag == nil {
//line /usr/local/go/src/encoding/asn1/common.go:134
				_go_fuzz_dep_.CoverTab[7977]++
										ret.tag = new(int)
//line /usr/local/go/src/encoding/asn1/common.go:135
				// _ = "end of CoverTab[7977]"
			} else {
//line /usr/local/go/src/encoding/asn1/common.go:136
				_go_fuzz_dep_.CoverTab[7978]++
//line /usr/local/go/src/encoding/asn1/common.go:136
				// _ = "end of CoverTab[7978]"
//line /usr/local/go/src/encoding/asn1/common.go:136
			}
//line /usr/local/go/src/encoding/asn1/common.go:136
			// _ = "end of CoverTab[7967]"
		case part == "private":
//line /usr/local/go/src/encoding/asn1/common.go:137
			_go_fuzz_dep_.CoverTab[7968]++
									ret.private = true
									if ret.tag == nil {
//line /usr/local/go/src/encoding/asn1/common.go:139
				_go_fuzz_dep_.CoverTab[7979]++
										ret.tag = new(int)
//line /usr/local/go/src/encoding/asn1/common.go:140
				// _ = "end of CoverTab[7979]"
			} else {
//line /usr/local/go/src/encoding/asn1/common.go:141
				_go_fuzz_dep_.CoverTab[7980]++
//line /usr/local/go/src/encoding/asn1/common.go:141
				// _ = "end of CoverTab[7980]"
//line /usr/local/go/src/encoding/asn1/common.go:141
			}
//line /usr/local/go/src/encoding/asn1/common.go:141
			// _ = "end of CoverTab[7968]"
		case part == "omitempty":
//line /usr/local/go/src/encoding/asn1/common.go:142
			_go_fuzz_dep_.CoverTab[7969]++
									ret.omitEmpty = true
//line /usr/local/go/src/encoding/asn1/common.go:143
			// _ = "end of CoverTab[7969]"
//line /usr/local/go/src/encoding/asn1/common.go:143
		default:
//line /usr/local/go/src/encoding/asn1/common.go:143
			_go_fuzz_dep_.CoverTab[7970]++
//line /usr/local/go/src/encoding/asn1/common.go:143
			// _ = "end of CoverTab[7970]"
		}
//line /usr/local/go/src/encoding/asn1/common.go:144
		// _ = "end of CoverTab[7955]"
	}
//line /usr/local/go/src/encoding/asn1/common.go:145
	// _ = "end of CoverTab[7953]"
//line /usr/local/go/src/encoding/asn1/common.go:145
	_go_fuzz_dep_.CoverTab[7954]++
							return
//line /usr/local/go/src/encoding/asn1/common.go:146
	// _ = "end of CoverTab[7954]"
}

// Given a reflected Go type, getUniversalType returns the default tag number
//line /usr/local/go/src/encoding/asn1/common.go:149
// and expected compound flag.
//line /usr/local/go/src/encoding/asn1/common.go:151
func getUniversalType(t reflect.Type) (matchAny bool, tagNumber int, isCompound, ok bool) {
//line /usr/local/go/src/encoding/asn1/common.go:151
	_go_fuzz_dep_.CoverTab[7981]++
							switch t {
	case rawValueType:
//line /usr/local/go/src/encoding/asn1/common.go:153
		_go_fuzz_dep_.CoverTab[7984]++
								return true, -1, false, true
//line /usr/local/go/src/encoding/asn1/common.go:154
		// _ = "end of CoverTab[7984]"
	case objectIdentifierType:
//line /usr/local/go/src/encoding/asn1/common.go:155
		_go_fuzz_dep_.CoverTab[7985]++
								return false, TagOID, false, true
//line /usr/local/go/src/encoding/asn1/common.go:156
		// _ = "end of CoverTab[7985]"
	case bitStringType:
//line /usr/local/go/src/encoding/asn1/common.go:157
		_go_fuzz_dep_.CoverTab[7986]++
								return false, TagBitString, false, true
//line /usr/local/go/src/encoding/asn1/common.go:158
		// _ = "end of CoverTab[7986]"
	case timeType:
//line /usr/local/go/src/encoding/asn1/common.go:159
		_go_fuzz_dep_.CoverTab[7987]++
								return false, TagUTCTime, false, true
//line /usr/local/go/src/encoding/asn1/common.go:160
		// _ = "end of CoverTab[7987]"
	case enumeratedType:
//line /usr/local/go/src/encoding/asn1/common.go:161
		_go_fuzz_dep_.CoverTab[7988]++
								return false, TagEnum, false, true
//line /usr/local/go/src/encoding/asn1/common.go:162
		// _ = "end of CoverTab[7988]"
	case bigIntType:
//line /usr/local/go/src/encoding/asn1/common.go:163
		_go_fuzz_dep_.CoverTab[7989]++
								return false, TagInteger, false, true
//line /usr/local/go/src/encoding/asn1/common.go:164
		// _ = "end of CoverTab[7989]"
//line /usr/local/go/src/encoding/asn1/common.go:164
	default:
//line /usr/local/go/src/encoding/asn1/common.go:164
		_go_fuzz_dep_.CoverTab[7990]++
//line /usr/local/go/src/encoding/asn1/common.go:164
		// _ = "end of CoverTab[7990]"
	}
//line /usr/local/go/src/encoding/asn1/common.go:165
	// _ = "end of CoverTab[7981]"
//line /usr/local/go/src/encoding/asn1/common.go:165
	_go_fuzz_dep_.CoverTab[7982]++
							switch t.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/encoding/asn1/common.go:167
		_go_fuzz_dep_.CoverTab[7991]++
								return false, TagBoolean, false, true
//line /usr/local/go/src/encoding/asn1/common.go:168
		// _ = "end of CoverTab[7991]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/asn1/common.go:169
		_go_fuzz_dep_.CoverTab[7992]++
								return false, TagInteger, false, true
//line /usr/local/go/src/encoding/asn1/common.go:170
		// _ = "end of CoverTab[7992]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/asn1/common.go:171
		_go_fuzz_dep_.CoverTab[7993]++
								return false, TagSequence, true, true
//line /usr/local/go/src/encoding/asn1/common.go:172
		// _ = "end of CoverTab[7993]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/asn1/common.go:173
		_go_fuzz_dep_.CoverTab[7994]++
								if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/asn1/common.go:174
			_go_fuzz_dep_.CoverTab[7999]++
									return false, TagOctetString, false, true
//line /usr/local/go/src/encoding/asn1/common.go:175
			// _ = "end of CoverTab[7999]"
		} else {
//line /usr/local/go/src/encoding/asn1/common.go:176
			_go_fuzz_dep_.CoverTab[8000]++
//line /usr/local/go/src/encoding/asn1/common.go:176
			// _ = "end of CoverTab[8000]"
//line /usr/local/go/src/encoding/asn1/common.go:176
		}
//line /usr/local/go/src/encoding/asn1/common.go:176
		// _ = "end of CoverTab[7994]"
//line /usr/local/go/src/encoding/asn1/common.go:176
		_go_fuzz_dep_.CoverTab[7995]++
								if strings.HasSuffix(t.Name(), "SET") {
//line /usr/local/go/src/encoding/asn1/common.go:177
			_go_fuzz_dep_.CoverTab[8001]++
									return false, TagSet, true, true
//line /usr/local/go/src/encoding/asn1/common.go:178
			// _ = "end of CoverTab[8001]"
		} else {
//line /usr/local/go/src/encoding/asn1/common.go:179
			_go_fuzz_dep_.CoverTab[8002]++
//line /usr/local/go/src/encoding/asn1/common.go:179
			// _ = "end of CoverTab[8002]"
//line /usr/local/go/src/encoding/asn1/common.go:179
		}
//line /usr/local/go/src/encoding/asn1/common.go:179
		// _ = "end of CoverTab[7995]"
//line /usr/local/go/src/encoding/asn1/common.go:179
		_go_fuzz_dep_.CoverTab[7996]++
								return false, TagSequence, true, true
//line /usr/local/go/src/encoding/asn1/common.go:180
		// _ = "end of CoverTab[7996]"
	case reflect.String:
//line /usr/local/go/src/encoding/asn1/common.go:181
		_go_fuzz_dep_.CoverTab[7997]++
								return false, TagPrintableString, false, true
//line /usr/local/go/src/encoding/asn1/common.go:182
		// _ = "end of CoverTab[7997]"
//line /usr/local/go/src/encoding/asn1/common.go:182
	default:
//line /usr/local/go/src/encoding/asn1/common.go:182
		_go_fuzz_dep_.CoverTab[7998]++
//line /usr/local/go/src/encoding/asn1/common.go:182
		// _ = "end of CoverTab[7998]"
	}
//line /usr/local/go/src/encoding/asn1/common.go:183
	// _ = "end of CoverTab[7982]"
//line /usr/local/go/src/encoding/asn1/common.go:183
	_go_fuzz_dep_.CoverTab[7983]++
							return false, 0, false, false
//line /usr/local/go/src/encoding/asn1/common.go:184
	// _ = "end of CoverTab[7983]"
}

//line /usr/local/go/src/encoding/asn1/common.go:185
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/asn1/common.go:185
var _ = _go_fuzz_dep_.CoverTab
