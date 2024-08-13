// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:5
// Package asn1 contains supporting types for parsing and building ASN.1
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:5
// messages with the cryptobyte package.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
package asn1

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:7
)

// Tag represents an ASN.1 identifier octet, consisting of a tag number
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:9
// (indicating a type) and class (such as context-specific or constructed).
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:9
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:9
// Methods in the cryptobyte package only support the low-tag-number form, i.e.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:9
// a single identifier octet with bits 7-8 encoding the class and bits 1-6
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:9
// encoding the tag number.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:15
type Tag uint8

const (
	classConstructed	= 0x20
	classContextSpecific	= 0x80
)

// Constructed returns t with the constructed class bit set.
func (t Tag) Constructed() Tag {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:23
	_go_fuzz_dep_.CoverTab[8301]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:23
	return t | classConstructed
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:23
	// _ = "end of CoverTab[8301]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:23
}

// ContextSpecific returns t with the context-specific class bit set.
func (t Tag) ContextSpecific() Tag {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:26
	_go_fuzz_dep_.CoverTab[8302]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:26
	return t | classContextSpecific
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:26
	// _ = "end of CoverTab[8302]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:26
}

// The following is a list of standard tag and class combinations.
const (
	BOOLEAN			= Tag(1)
	INTEGER			= Tag(2)
	BIT_STRING		= Tag(3)
	OCTET_STRING		= Tag(4)
	NULL			= Tag(5)
	OBJECT_IDENTIFIER	= Tag(6)
	ENUM			= Tag(10)
	UTF8String		= Tag(12)
	SEQUENCE		= Tag(16 | classConstructed)
	SET			= Tag(17 | classConstructed)
	PrintableString		= Tag(19)
	T61String		= Tag(20)
	IA5String		= Tag(22)
	UTCTime			= Tag(23)
	GeneralizedTime		= Tag(24)
	GeneralString		= Tag(27)
)

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1/asn1.go:46
var _ = _go_fuzz_dep_.CoverTab
