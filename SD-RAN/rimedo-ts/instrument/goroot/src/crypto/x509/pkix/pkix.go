// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/pkix/pkix.go:5
// Package pkix contains shared, low level structures used for ASN.1 parsing
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:5
// and serialization of X.509 certificates, CRL and OCSP.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
package pkix

//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
import (
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
import (
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:7
)

import (
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

// AlgorithmIdentifier represents the ASN.1 structure of the same name. See RFC
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:17
// 5280, section 4.1.1.2.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:19
type AlgorithmIdentifier struct {
	Algorithm	asn1.ObjectIdentifier
	Parameters	asn1.RawValue	`asn1:"optional"`
}

type RDNSequence []RelativeDistinguishedNameSET

var attributeTypeNames = map[string]string{
	"2.5.4.6":	"C",
	"2.5.4.10":	"O",
	"2.5.4.11":	"OU",
	"2.5.4.3":	"CN",
	"2.5.4.5":	"SERIALNUMBER",
	"2.5.4.7":	"L",
	"2.5.4.8":	"ST",
	"2.5.4.9":	"STREET",
	"2.5.4.17":	"POSTALCODE",
}

// String returns a string representation of the sequence r,
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:38
// roughly following the RFC 2253 Distinguished Names syntax.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:40
func (r RDNSequence) String() string {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:40
	_go_fuzz_dep_.CoverTab[10436]++
							s := ""
							for i := 0; i < len(r); i++ {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:42
		_go_fuzz_dep_.CoverTab[10438]++
								rdn := r[len(r)-1-i]
								if i > 0 {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:44
			_go_fuzz_dep_.CoverTab[10440]++
									s += ","
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:45
			// _ = "end of CoverTab[10440]"
		} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:46
			_go_fuzz_dep_.CoverTab[10441]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:46
			// _ = "end of CoverTab[10441]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:46
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:46
		// _ = "end of CoverTab[10438]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:46
		_go_fuzz_dep_.CoverTab[10439]++
								for j, tv := range rdn {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:47
			_go_fuzz_dep_.CoverTab[10442]++
									if j > 0 {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:48
				_go_fuzz_dep_.CoverTab[10446]++
										s += "+"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:49
				// _ = "end of CoverTab[10446]"
			} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:50
				_go_fuzz_dep_.CoverTab[10447]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:50
				// _ = "end of CoverTab[10447]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:50
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:50
			// _ = "end of CoverTab[10442]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:50
			_go_fuzz_dep_.CoverTab[10443]++

									oidString := tv.Type.String()
									typeName, ok := attributeTypeNames[oidString]
									if !ok {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:54
				_go_fuzz_dep_.CoverTab[10448]++
										derBytes, err := asn1.Marshal(tv.Value)
										if err == nil {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:56
					_go_fuzz_dep_.CoverTab[10450]++
											s += oidString + "=#" + hex.EncodeToString(derBytes)
											continue
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:58
					// _ = "end of CoverTab[10450]"
				} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:59
					_go_fuzz_dep_.CoverTab[10451]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:59
					// _ = "end of CoverTab[10451]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:59
				}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:59
				// _ = "end of CoverTab[10448]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:59
				_go_fuzz_dep_.CoverTab[10449]++

										typeName = oidString
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:61
				// _ = "end of CoverTab[10449]"
			} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:62
				_go_fuzz_dep_.CoverTab[10452]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:62
				// _ = "end of CoverTab[10452]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:62
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:62
			// _ = "end of CoverTab[10443]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:62
			_go_fuzz_dep_.CoverTab[10444]++

									valueString := fmt.Sprint(tv.Value)
									escaped := make([]rune, 0, len(valueString))

									for k, c := range valueString {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:67
				_go_fuzz_dep_.CoverTab[10453]++
										escape := false

										switch c {
				case ',', '+', '"', '\\', '<', '>', ';':
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:71
					_go_fuzz_dep_.CoverTab[10455]++
											escape = true
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:72
					// _ = "end of CoverTab[10455]"

				case ' ':
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:74
					_go_fuzz_dep_.CoverTab[10456]++
											escape = k == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:75
						_go_fuzz_dep_.CoverTab[10459]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:75
						return k == len(valueString)-1
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:75
						// _ = "end of CoverTab[10459]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:75
					}()
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:75
					// _ = "end of CoverTab[10456]"

				case '#':
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:77
					_go_fuzz_dep_.CoverTab[10457]++
											escape = k == 0
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:78
					// _ = "end of CoverTab[10457]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:78
				default:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:78
					_go_fuzz_dep_.CoverTab[10458]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:78
					// _ = "end of CoverTab[10458]"
				}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:79
				// _ = "end of CoverTab[10453]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:79
				_go_fuzz_dep_.CoverTab[10454]++

										if escape {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:81
					_go_fuzz_dep_.CoverTab[10460]++
											escaped = append(escaped, '\\', c)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:82
					// _ = "end of CoverTab[10460]"
				} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:83
					_go_fuzz_dep_.CoverTab[10461]++
											escaped = append(escaped, c)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:84
					// _ = "end of CoverTab[10461]"
				}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:85
				// _ = "end of CoverTab[10454]"
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:86
			// _ = "end of CoverTab[10444]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:86
			_go_fuzz_dep_.CoverTab[10445]++

									s += typeName + "=" + string(escaped)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:88
			// _ = "end of CoverTab[10445]"
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:89
		// _ = "end of CoverTab[10439]"
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:90
	// _ = "end of CoverTab[10436]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:90
	_go_fuzz_dep_.CoverTab[10437]++

							return s
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:92
	// _ = "end of CoverTab[10437]"
}

type RelativeDistinguishedNameSET []AttributeTypeAndValue

// AttributeTypeAndValue mirrors the ASN.1 structure of the same name in
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:97
// RFC 5280, Section 4.1.2.4.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:99
type AttributeTypeAndValue struct {
	Type	asn1.ObjectIdentifier
	Value	any
}

// AttributeTypeAndValueSET represents a set of ASN.1 sequences of
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:104
// AttributeTypeAndValue sequences from RFC 2986 (PKCS #10).
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:106
type AttributeTypeAndValueSET struct {
	Type	asn1.ObjectIdentifier
	Value	[][]AttributeTypeAndValue	`asn1:"set"`
}

// Extension represents the ASN.1 structure of the same name. See RFC
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:111
// 5280, section 4.2.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:113
type Extension struct {
	Id		asn1.ObjectIdentifier
	Critical	bool	`asn1:"optional"`
	Value		[]byte
}

// Name represents an X.509 distinguished name. This only includes the common
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:119
// elements of a DN. Note that Name is only an approximation of the X.509
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:119
// structure. If an accurate representation is needed, asn1.Unmarshal the raw
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:119
// subject or issuer as an RDNSequence.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:123
type Name struct {
	Country, Organization, OrganizationalUnit	[]string
	Locality, Province				[]string
	StreetAddress, PostalCode			[]string
	SerialNumber, CommonName			string

	// Names contains all parsed attributes. When parsing distinguished names,
	// this can be used to extract non-standard attributes that are not parsed
	// by this package. When marshaling to RDNSequences, the Names field is
	// ignored, see ExtraNames.
	Names	[]AttributeTypeAndValue

	// ExtraNames contains attributes to be copied, raw, into any marshaled
	// distinguished names. Values override any attributes with the same OID.
	// The ExtraNames field is not populated when parsing, see Names.
	ExtraNames	[]AttributeTypeAndValue
}

// FillFromRDNSequence populates n from the provided RDNSequence.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:141
// Multi-entry RDNs are flattened, all entries are added to the
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:141
// relevant n fields, and the grouping is not preserved.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:144
func (n *Name) FillFromRDNSequence(rdns *RDNSequence) {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:144
	_go_fuzz_dep_.CoverTab[10462]++
							for _, rdn := range *rdns {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:145
		_go_fuzz_dep_.CoverTab[10463]++
								if len(rdn) == 0 {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:146
			_go_fuzz_dep_.CoverTab[10465]++
									continue
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:147
			// _ = "end of CoverTab[10465]"
		} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:148
			_go_fuzz_dep_.CoverTab[10466]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:148
			// _ = "end of CoverTab[10466]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:148
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:148
		// _ = "end of CoverTab[10463]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:148
		_go_fuzz_dep_.CoverTab[10464]++

								for _, atv := range rdn {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:150
			_go_fuzz_dep_.CoverTab[10467]++
									n.Names = append(n.Names, atv)
									value, ok := atv.Value.(string)
									if !ok {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:153
				_go_fuzz_dep_.CoverTab[10469]++
										continue
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:154
				// _ = "end of CoverTab[10469]"
			} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:155
				_go_fuzz_dep_.CoverTab[10470]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:155
				// _ = "end of CoverTab[10470]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:155
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:155
			// _ = "end of CoverTab[10467]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:155
			_go_fuzz_dep_.CoverTab[10468]++

									t := atv.Type
									if len(t) == 4 && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				_go_fuzz_dep_.CoverTab[10471]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				return t[0] == 2
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				// _ = "end of CoverTab[10471]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
			}() && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				_go_fuzz_dep_.CoverTab[10472]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				return t[1] == 5
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				// _ = "end of CoverTab[10472]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
			}() && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				_go_fuzz_dep_.CoverTab[10473]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				return t[2] == 4
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				// _ = "end of CoverTab[10473]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
			}() {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:158
				_go_fuzz_dep_.CoverTab[10474]++
										switch t[3] {
				case 3:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:160
					_go_fuzz_dep_.CoverTab[10475]++
											n.CommonName = value
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:161
					// _ = "end of CoverTab[10475]"
				case 5:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:162
					_go_fuzz_dep_.CoverTab[10476]++
											n.SerialNumber = value
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:163
					// _ = "end of CoverTab[10476]"
				case 6:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:164
					_go_fuzz_dep_.CoverTab[10477]++
											n.Country = append(n.Country, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:165
					// _ = "end of CoverTab[10477]"
				case 7:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:166
					_go_fuzz_dep_.CoverTab[10478]++
											n.Locality = append(n.Locality, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:167
					// _ = "end of CoverTab[10478]"
				case 8:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:168
					_go_fuzz_dep_.CoverTab[10479]++
											n.Province = append(n.Province, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:169
					// _ = "end of CoverTab[10479]"
				case 9:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:170
					_go_fuzz_dep_.CoverTab[10480]++
											n.StreetAddress = append(n.StreetAddress, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:171
					// _ = "end of CoverTab[10480]"
				case 10:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:172
					_go_fuzz_dep_.CoverTab[10481]++
											n.Organization = append(n.Organization, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:173
					// _ = "end of CoverTab[10481]"
				case 11:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:174
					_go_fuzz_dep_.CoverTab[10482]++
											n.OrganizationalUnit = append(n.OrganizationalUnit, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:175
					// _ = "end of CoverTab[10482]"
				case 17:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:176
					_go_fuzz_dep_.CoverTab[10483]++
											n.PostalCode = append(n.PostalCode, value)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:177
					// _ = "end of CoverTab[10483]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:177
				default:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:177
					_go_fuzz_dep_.CoverTab[10484]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:177
					// _ = "end of CoverTab[10484]"
				}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:178
				// _ = "end of CoverTab[10474]"
			} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:179
				_go_fuzz_dep_.CoverTab[10485]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:179
				// _ = "end of CoverTab[10485]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:179
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:179
			// _ = "end of CoverTab[10468]"
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:180
		// _ = "end of CoverTab[10464]"
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:181
	// _ = "end of CoverTab[10462]"
}

var (
	oidCountry		= []int{2, 5, 4, 6}
	oidOrganization		= []int{2, 5, 4, 10}
	oidOrganizationalUnit	= []int{2, 5, 4, 11}
	oidCommonName		= []int{2, 5, 4, 3}
	oidSerialNumber		= []int{2, 5, 4, 5}
	oidLocality		= []int{2, 5, 4, 7}
	oidProvince		= []int{2, 5, 4, 8}
	oidStreetAddress	= []int{2, 5, 4, 9}
	oidPostalCode		= []int{2, 5, 4, 17}
)

// appendRDNs appends a relativeDistinguishedNameSET to the given RDNSequence
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:196
// and returns the new value. The relativeDistinguishedNameSET contains an
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:196
// attributeTypeAndValue for each of the given values. See RFC 5280, A.1, and
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:196
// search for AttributeTypeAndValue.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:200
func (n Name) appendRDNs(in RDNSequence, values []string, oid asn1.ObjectIdentifier) RDNSequence {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:200
	_go_fuzz_dep_.CoverTab[10486]++
							if len(values) == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:201
		_go_fuzz_dep_.CoverTab[10489]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:201
		return oidInAttributeTypeAndValue(oid, n.ExtraNames)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:201
		// _ = "end of CoverTab[10489]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:201
	}() {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:201
		_go_fuzz_dep_.CoverTab[10490]++
								return in
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:202
		// _ = "end of CoverTab[10490]"
	} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:203
		_go_fuzz_dep_.CoverTab[10491]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:203
		// _ = "end of CoverTab[10491]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:203
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:203
	// _ = "end of CoverTab[10486]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:203
	_go_fuzz_dep_.CoverTab[10487]++

							s := make([]AttributeTypeAndValue, len(values))
							for i, value := range values {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:206
		_go_fuzz_dep_.CoverTab[10492]++
								s[i].Type = oid
								s[i].Value = value
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:208
		// _ = "end of CoverTab[10492]"
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:209
	// _ = "end of CoverTab[10487]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:209
	_go_fuzz_dep_.CoverTab[10488]++

							return append(in, s)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:211
	// _ = "end of CoverTab[10488]"
}

// ToRDNSequence converts n into a single RDNSequence. The following
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
// attributes are encoded as multi-value RDNs:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - Country
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - Organization
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - OrganizationalUnit
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - Locality
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - Province
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - StreetAddress
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//   - PostalCode
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
//
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:214
// Each ExtraNames entry is encoded as an individual RDN.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:226
func (n Name) ToRDNSequence() (ret RDNSequence) {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:226
	_go_fuzz_dep_.CoverTab[10493]++
							ret = n.appendRDNs(ret, n.Country, oidCountry)
							ret = n.appendRDNs(ret, n.Province, oidProvince)
							ret = n.appendRDNs(ret, n.Locality, oidLocality)
							ret = n.appendRDNs(ret, n.StreetAddress, oidStreetAddress)
							ret = n.appendRDNs(ret, n.PostalCode, oidPostalCode)
							ret = n.appendRDNs(ret, n.Organization, oidOrganization)
							ret = n.appendRDNs(ret, n.OrganizationalUnit, oidOrganizationalUnit)
							if len(n.CommonName) > 0 {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:234
		_go_fuzz_dep_.CoverTab[10497]++
								ret = n.appendRDNs(ret, []string{n.CommonName}, oidCommonName)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:235
		// _ = "end of CoverTab[10497]"
	} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:236
		_go_fuzz_dep_.CoverTab[10498]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:236
		// _ = "end of CoverTab[10498]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:236
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:236
	// _ = "end of CoverTab[10493]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:236
	_go_fuzz_dep_.CoverTab[10494]++
							if len(n.SerialNumber) > 0 {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:237
		_go_fuzz_dep_.CoverTab[10499]++
								ret = n.appendRDNs(ret, []string{n.SerialNumber}, oidSerialNumber)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:238
		// _ = "end of CoverTab[10499]"
	} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:239
		_go_fuzz_dep_.CoverTab[10500]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:239
		// _ = "end of CoverTab[10500]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:239
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:239
	// _ = "end of CoverTab[10494]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:239
	_go_fuzz_dep_.CoverTab[10495]++
							for _, atv := range n.ExtraNames {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:240
		_go_fuzz_dep_.CoverTab[10501]++
								ret = append(ret, []AttributeTypeAndValue{atv})
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:241
		// _ = "end of CoverTab[10501]"
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:242
	// _ = "end of CoverTab[10495]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:242
	_go_fuzz_dep_.CoverTab[10496]++

							return ret
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:244
	// _ = "end of CoverTab[10496]"
}

// String returns the string form of n, roughly following
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:247
// the RFC 2253 Distinguished Names syntax.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:249
func (n Name) String() string {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:249
	_go_fuzz_dep_.CoverTab[10502]++
							var rdns RDNSequence

//line /usr/local/go/src/crypto/x509/pkix/pkix.go:253
	if n.ExtraNames == nil {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:253
		_go_fuzz_dep_.CoverTab[10504]++
								for _, atv := range n.Names {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:254
			_go_fuzz_dep_.CoverTab[10505]++
									t := atv.Type
									if len(t) == 4 && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				_go_fuzz_dep_.CoverTab[10507]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				return t[0] == 2
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				// _ = "end of CoverTab[10507]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
			}() && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				_go_fuzz_dep_.CoverTab[10508]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				return t[1] == 5
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				// _ = "end of CoverTab[10508]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
			}() && func() bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				_go_fuzz_dep_.CoverTab[10509]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				return t[2] == 4
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				// _ = "end of CoverTab[10509]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
			}() {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:256
				_go_fuzz_dep_.CoverTab[10510]++
										switch t[3] {
				case 3, 5, 6, 7, 8, 9, 10, 11, 17:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:258
					_go_fuzz_dep_.CoverTab[10511]++

											continue
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:260
					// _ = "end of CoverTab[10511]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:260
				default:
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:260
					_go_fuzz_dep_.CoverTab[10512]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:260
					// _ = "end of CoverTab[10512]"
				}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:261
				// _ = "end of CoverTab[10510]"
			} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:262
				_go_fuzz_dep_.CoverTab[10513]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:262
				// _ = "end of CoverTab[10513]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:262
			}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:262
			// _ = "end of CoverTab[10505]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:262
			_go_fuzz_dep_.CoverTab[10506]++

//line /usr/local/go/src/crypto/x509/pkix/pkix.go:265
			rdns = append(rdns, []AttributeTypeAndValue{atv})
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:265
			// _ = "end of CoverTab[10506]"
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:266
		// _ = "end of CoverTab[10504]"
	} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:267
		_go_fuzz_dep_.CoverTab[10514]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:267
		// _ = "end of CoverTab[10514]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:267
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:267
	// _ = "end of CoverTab[10502]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:267
	_go_fuzz_dep_.CoverTab[10503]++
							rdns = append(rdns, n.ToRDNSequence()...)
							return rdns.String()
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:269
	// _ = "end of CoverTab[10503]"
}

// oidInAttributeTypeAndValue reports whether a type with the given OID exists
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:272
// in atv.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:274
func oidInAttributeTypeAndValue(oid asn1.ObjectIdentifier, atv []AttributeTypeAndValue) bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:274
	_go_fuzz_dep_.CoverTab[10515]++
							for _, a := range atv {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:275
		_go_fuzz_dep_.CoverTab[10517]++
								if a.Type.Equal(oid) {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:276
			_go_fuzz_dep_.CoverTab[10518]++
									return true
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:277
			// _ = "end of CoverTab[10518]"
		} else {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:278
			_go_fuzz_dep_.CoverTab[10519]++
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:278
			// _ = "end of CoverTab[10519]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:278
		}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:278
		// _ = "end of CoverTab[10517]"
	}
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:279
	// _ = "end of CoverTab[10515]"
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:279
	_go_fuzz_dep_.CoverTab[10516]++
							return false
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:280
	// _ = "end of CoverTab[10516]"
}

// CertificateList represents the ASN.1 structure of the same name. See RFC
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:283
// 5280, section 5.1. Use Certificate.CheckCRLSignature to verify the
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:283
// signature.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:283
//
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:283
// Deprecated: x509.RevocationList should be used instead.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:288
type CertificateList struct {
	TBSCertList		TBSCertificateList
	SignatureAlgorithm	AlgorithmIdentifier
	SignatureValue		asn1.BitString
}

// HasExpired reports whether certList should have been updated by now.
func (certList *CertificateList) HasExpired(now time.Time) bool {
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:295
	_go_fuzz_dep_.CoverTab[10520]++
							return !now.Before(certList.TBSCertList.NextUpdate)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:296
	// _ = "end of CoverTab[10520]"
}

// TBSCertificateList represents the ASN.1 structure of the same name. See RFC
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:299
// 5280, section 5.1.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:299
//
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:299
// Deprecated: x509.RevocationList should be used instead.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:303
type TBSCertificateList struct {
	Raw			asn1.RawContent
	Version			int	`asn1:"optional,default:0"`
	Signature		AlgorithmIdentifier
	Issuer			RDNSequence
	ThisUpdate		time.Time
	NextUpdate		time.Time		`asn1:"optional"`
	RevokedCertificates	[]RevokedCertificate	`asn1:"optional"`
	Extensions		[]Extension		`asn1:"tag:0,optional,explicit"`
}

// RevokedCertificate represents the ASN.1 structure of the same name. See RFC
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:314
// 5280, section 5.1.
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:316
type RevokedCertificate struct {
	SerialNumber	*big.Int
	RevocationTime	time.Time
	Extensions	[]Extension	`asn1:"optional"`
}

//line /usr/local/go/src/crypto/x509/pkix/pkix.go:320
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/pkix/pkix.go:320
var _ = _go_fuzz_dep_.CoverTab
