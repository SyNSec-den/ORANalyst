// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/sniff.go:5
package http

//line /usr/local/go/src/net/http/sniff.go:5
import (
//line /usr/local/go/src/net/http/sniff.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/sniff.go:5
)
//line /usr/local/go/src/net/http/sniff.go:5
import (
//line /usr/local/go/src/net/http/sniff.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/sniff.go:5
)

import (
	"bytes"
	"encoding/binary"
)

// The algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 512

// DetectContentType implements the algorithm described
//line /usr/local/go/src/net/http/sniff.go:15
// at https://mimesniff.spec.whatwg.org/ to determine the
//line /usr/local/go/src/net/http/sniff.go:15
// Content-Type of the given data. It considers at most the
//line /usr/local/go/src/net/http/sniff.go:15
// first 512 bytes of data. DetectContentType always returns
//line /usr/local/go/src/net/http/sniff.go:15
// a valid MIME type: if it cannot determine a more specific one, it
//line /usr/local/go/src/net/http/sniff.go:15
// returns "application/octet-stream".
//line /usr/local/go/src/net/http/sniff.go:21
func DetectContentType(data []byte) string {
//line /usr/local/go/src/net/http/sniff.go:21
	_go_fuzz_dep_.CoverTab[43206]++
						if len(data) > sniffLen {
//line /usr/local/go/src/net/http/sniff.go:22
		_go_fuzz_dep_.CoverTab[43210]++
							data = data[:sniffLen]
//line /usr/local/go/src/net/http/sniff.go:23
		// _ = "end of CoverTab[43210]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:24
		_go_fuzz_dep_.CoverTab[43211]++
//line /usr/local/go/src/net/http/sniff.go:24
		// _ = "end of CoverTab[43211]"
//line /usr/local/go/src/net/http/sniff.go:24
	}
//line /usr/local/go/src/net/http/sniff.go:24
	// _ = "end of CoverTab[43206]"
//line /usr/local/go/src/net/http/sniff.go:24
	_go_fuzz_dep_.CoverTab[43207]++

//line /usr/local/go/src/net/http/sniff.go:27
	firstNonWS := 0
	for ; firstNonWS < len(data) && func() bool {
//line /usr/local/go/src/net/http/sniff.go:28
		_go_fuzz_dep_.CoverTab[43212]++
//line /usr/local/go/src/net/http/sniff.go:28
		return isWS(data[firstNonWS])
//line /usr/local/go/src/net/http/sniff.go:28
		// _ = "end of CoverTab[43212]"
//line /usr/local/go/src/net/http/sniff.go:28
	}(); firstNonWS++ {
//line /usr/local/go/src/net/http/sniff.go:28
		_go_fuzz_dep_.CoverTab[43213]++
//line /usr/local/go/src/net/http/sniff.go:28
		// _ = "end of CoverTab[43213]"
	}
//line /usr/local/go/src/net/http/sniff.go:29
	// _ = "end of CoverTab[43207]"
//line /usr/local/go/src/net/http/sniff.go:29
	_go_fuzz_dep_.CoverTab[43208]++

						for _, sig := range sniffSignatures {
//line /usr/local/go/src/net/http/sniff.go:31
		_go_fuzz_dep_.CoverTab[43214]++
							if ct := sig.match(data, firstNonWS); ct != "" {
//line /usr/local/go/src/net/http/sniff.go:32
			_go_fuzz_dep_.CoverTab[43215]++
								return ct
//line /usr/local/go/src/net/http/sniff.go:33
			// _ = "end of CoverTab[43215]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:34
			_go_fuzz_dep_.CoverTab[43216]++
//line /usr/local/go/src/net/http/sniff.go:34
			// _ = "end of CoverTab[43216]"
//line /usr/local/go/src/net/http/sniff.go:34
		}
//line /usr/local/go/src/net/http/sniff.go:34
		// _ = "end of CoverTab[43214]"
	}
//line /usr/local/go/src/net/http/sniff.go:35
	// _ = "end of CoverTab[43208]"
//line /usr/local/go/src/net/http/sniff.go:35
	_go_fuzz_dep_.CoverTab[43209]++

						return "application/octet-stream"
//line /usr/local/go/src/net/http/sniff.go:37
	// _ = "end of CoverTab[43209]"
}

// isWS reports whether the provided byte is a whitespace byte (0xWS)
//line /usr/local/go/src/net/http/sniff.go:40
// as defined in https://mimesniff.spec.whatwg.org/#terminology.
//line /usr/local/go/src/net/http/sniff.go:42
func isWS(b byte) bool {
//line /usr/local/go/src/net/http/sniff.go:42
	_go_fuzz_dep_.CoverTab[43217]++
						switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
//line /usr/local/go/src/net/http/sniff.go:44
		_go_fuzz_dep_.CoverTab[43219]++
							return true
//line /usr/local/go/src/net/http/sniff.go:45
		// _ = "end of CoverTab[43219]"
//line /usr/local/go/src/net/http/sniff.go:45
	default:
//line /usr/local/go/src/net/http/sniff.go:45
		_go_fuzz_dep_.CoverTab[43220]++
//line /usr/local/go/src/net/http/sniff.go:45
		// _ = "end of CoverTab[43220]"
	}
//line /usr/local/go/src/net/http/sniff.go:46
	// _ = "end of CoverTab[43217]"
//line /usr/local/go/src/net/http/sniff.go:46
	_go_fuzz_dep_.CoverTab[43218]++
						return false
//line /usr/local/go/src/net/http/sniff.go:47
	// _ = "end of CoverTab[43218]"
}

// isTT reports whether the provided byte is a tag-terminating byte (0xTT)
//line /usr/local/go/src/net/http/sniff.go:50
// as defined in https://mimesniff.spec.whatwg.org/#terminology.
//line /usr/local/go/src/net/http/sniff.go:52
func isTT(b byte) bool {
//line /usr/local/go/src/net/http/sniff.go:52
	_go_fuzz_dep_.CoverTab[43221]++
						switch b {
	case ' ', '>':
//line /usr/local/go/src/net/http/sniff.go:54
		_go_fuzz_dep_.CoverTab[43223]++
							return true
//line /usr/local/go/src/net/http/sniff.go:55
		// _ = "end of CoverTab[43223]"
//line /usr/local/go/src/net/http/sniff.go:55
	default:
//line /usr/local/go/src/net/http/sniff.go:55
		_go_fuzz_dep_.CoverTab[43224]++
//line /usr/local/go/src/net/http/sniff.go:55
		// _ = "end of CoverTab[43224]"
	}
//line /usr/local/go/src/net/http/sniff.go:56
	// _ = "end of CoverTab[43221]"
//line /usr/local/go/src/net/http/sniff.go:56
	_go_fuzz_dep_.CoverTab[43222]++
						return false
//line /usr/local/go/src/net/http/sniff.go:57
	// _ = "end of CoverTab[43222]"
}

type sniffSig interface {
	// match returns the MIME type of the data, or "" if unknown.
	match(data []byte, firstNonWS int) string
}

// Data matching the table in section 6.
var sniffSignatures = []sniffSig{
	htmlSig("<!DOCTYPE HTML"),
	htmlSig("<HTML"),
	htmlSig("<HEAD"),
	htmlSig("<SCRIPT"),
	htmlSig("<IFRAME"),
	htmlSig("<H1"),
	htmlSig("<DIV"),
	htmlSig("<FONT"),
	htmlSig("<TABLE"),
	htmlSig("<A"),
	htmlSig("<STYLE"),
	htmlSig("<TITLE"),
	htmlSig("<B"),
	htmlSig("<BODY"),
	htmlSig("<BR"),
	htmlSig("<P"),
	htmlSig("<!--"),
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\xFF"),
		pat:	[]byte("<?xml"),
		skipWS:	true,
		ct:	"text/xml; charset=utf-8"},
	&exactSig{[]byte("%PDF-"), "application/pdf"},
	&exactSig{[]byte("%!PS-Adobe-"), "application/postscript"},

//line /usr/local/go/src/net/http/sniff.go:93
	&maskedSig{
		mask:	[]byte("\xFF\xFF\x00\x00"),
		pat:	[]byte("\xFE\xFF\x00\x00"),
		ct:	"text/plain; charset=utf-16be",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\x00\x00"),
		pat:	[]byte("\xFF\xFE\x00\x00"),
		ct:	"text/plain; charset=utf-16le",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\x00"),
		pat:	[]byte("\xEF\xBB\xBF\x00"),
		ct:	"text/plain; charset=utf-8",
	},

//line /usr/local/go/src/net/http/sniff.go:115
	&exactSig{[]byte("\x00\x00\x01\x00"), "image/x-icon"},
						&exactSig{[]byte("\x00\x00\x02\x00"), "image/x-icon"},
						&exactSig{[]byte("BM"), "image/bmp"},
						&exactSig{[]byte("GIF87a"), "image/gif"},
						&exactSig{[]byte("GIF89a"), "image/gif"},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:	[]byte("RIFF\x00\x00\x00\x00WEBPVP"),
		ct:	"image/webp",
	},
	&exactSig{[]byte("\x89PNG\x0D\x0A\x1A\x0A"), "image/png"},
	&exactSig{[]byte("\xFF\xD8\xFF"), "image/jpeg"},

//line /usr/local/go/src/net/http/sniff.go:131
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:	[]byte("FORM\x00\x00\x00\x00AIFF"),
		ct:	"audio/aiff",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF"),
		pat:	[]byte("ID3"),
		ct:	"audio/mpeg",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\xFF"),
		pat:	[]byte("OggS\x00"),
		ct:	"application/ogg",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:	[]byte("MThd\x00\x00\x00\x06"),
		ct:	"audio/midi",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:	[]byte("RIFF\x00\x00\x00\x00AVI "),
		ct:	"video/avi",
	},
	&maskedSig{
		mask:	[]byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:	[]byte("RIFF\x00\x00\x00\x00WAVE"),
		ct:	"audio/wave",
	},

	mp4Sig{},

	&exactSig{[]byte("\x1A\x45\xDF\xA3"), "video/webm"},

//line /usr/local/go/src/net/http/sniff.go:167
	&maskedSig{

		pat:	[]byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00LP"),

		mask:	[]byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF"),
		ct:	"application/vnd.ms-fontobject",
	},
						&exactSig{[]byte("\x00\x01\x00\x00"), "font/ttf"},
						&exactSig{[]byte("OTTO"), "font/otf"},
						&exactSig{[]byte("ttcf"), "font/collection"},
						&exactSig{[]byte("wOFF"), "font/woff"},
						&exactSig{[]byte("wOF2"), "font/woff2"},

//line /usr/local/go/src/net/http/sniff.go:181
	&exactSig{[]byte("\x1F\x8B\x08"), "application/x-gzip"},
						&exactSig{[]byte("PK\x03\x04"), "application/zip"},

//line /usr/local/go/src/net/http/sniff.go:189
	&exactSig{[]byte("Rar!\x1A\x07\x00"), "application/x-rar-compressed"},
						&exactSig{[]byte("Rar!\x1A\x07\x01\x00"), "application/x-rar-compressed"},

						&exactSig{[]byte("\x00\x61\x73\x6D"), "application/wasm"},

	textSig{},
}

type exactSig struct {
	sig	[]byte
	ct	string
}

func (e *exactSig) match(data []byte, firstNonWS int) string {
//line /usr/local/go/src/net/http/sniff.go:202
	_go_fuzz_dep_.CoverTab[43225]++
						if bytes.HasPrefix(data, e.sig) {
//line /usr/local/go/src/net/http/sniff.go:203
		_go_fuzz_dep_.CoverTab[43227]++
							return e.ct
//line /usr/local/go/src/net/http/sniff.go:204
		// _ = "end of CoverTab[43227]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:205
		_go_fuzz_dep_.CoverTab[43228]++
//line /usr/local/go/src/net/http/sniff.go:205
		// _ = "end of CoverTab[43228]"
//line /usr/local/go/src/net/http/sniff.go:205
	}
//line /usr/local/go/src/net/http/sniff.go:205
	// _ = "end of CoverTab[43225]"
//line /usr/local/go/src/net/http/sniff.go:205
	_go_fuzz_dep_.CoverTab[43226]++
						return ""
//line /usr/local/go/src/net/http/sniff.go:206
	// _ = "end of CoverTab[43226]"
}

type maskedSig struct {
	mask, pat	[]byte
	skipWS		bool
	ct		string
}

func (m *maskedSig) match(data []byte, firstNonWS int) string {
//line /usr/local/go/src/net/http/sniff.go:215
	_go_fuzz_dep_.CoverTab[43229]++

//line /usr/local/go/src/net/http/sniff.go:219
	if m.skipWS {
//line /usr/local/go/src/net/http/sniff.go:219
		_go_fuzz_dep_.CoverTab[43234]++
							data = data[firstNonWS:]
//line /usr/local/go/src/net/http/sniff.go:220
		// _ = "end of CoverTab[43234]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:221
		_go_fuzz_dep_.CoverTab[43235]++
//line /usr/local/go/src/net/http/sniff.go:221
		// _ = "end of CoverTab[43235]"
//line /usr/local/go/src/net/http/sniff.go:221
	}
//line /usr/local/go/src/net/http/sniff.go:221
	// _ = "end of CoverTab[43229]"
//line /usr/local/go/src/net/http/sniff.go:221
	_go_fuzz_dep_.CoverTab[43230]++
						if len(m.pat) != len(m.mask) {
//line /usr/local/go/src/net/http/sniff.go:222
		_go_fuzz_dep_.CoverTab[43236]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:223
		// _ = "end of CoverTab[43236]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:224
		_go_fuzz_dep_.CoverTab[43237]++
//line /usr/local/go/src/net/http/sniff.go:224
		// _ = "end of CoverTab[43237]"
//line /usr/local/go/src/net/http/sniff.go:224
	}
//line /usr/local/go/src/net/http/sniff.go:224
	// _ = "end of CoverTab[43230]"
//line /usr/local/go/src/net/http/sniff.go:224
	_go_fuzz_dep_.CoverTab[43231]++
						if len(data) < len(m.pat) {
//line /usr/local/go/src/net/http/sniff.go:225
		_go_fuzz_dep_.CoverTab[43238]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:226
		// _ = "end of CoverTab[43238]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:227
		_go_fuzz_dep_.CoverTab[43239]++
//line /usr/local/go/src/net/http/sniff.go:227
		// _ = "end of CoverTab[43239]"
//line /usr/local/go/src/net/http/sniff.go:227
	}
//line /usr/local/go/src/net/http/sniff.go:227
	// _ = "end of CoverTab[43231]"
//line /usr/local/go/src/net/http/sniff.go:227
	_go_fuzz_dep_.CoverTab[43232]++
						for i, pb := range m.pat {
//line /usr/local/go/src/net/http/sniff.go:228
		_go_fuzz_dep_.CoverTab[43240]++
							maskedData := data[i] & m.mask[i]
							if maskedData != pb {
//line /usr/local/go/src/net/http/sniff.go:230
			_go_fuzz_dep_.CoverTab[43241]++
								return ""
//line /usr/local/go/src/net/http/sniff.go:231
			// _ = "end of CoverTab[43241]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:232
			_go_fuzz_dep_.CoverTab[43242]++
//line /usr/local/go/src/net/http/sniff.go:232
			// _ = "end of CoverTab[43242]"
//line /usr/local/go/src/net/http/sniff.go:232
		}
//line /usr/local/go/src/net/http/sniff.go:232
		// _ = "end of CoverTab[43240]"
	}
//line /usr/local/go/src/net/http/sniff.go:233
	// _ = "end of CoverTab[43232]"
//line /usr/local/go/src/net/http/sniff.go:233
	_go_fuzz_dep_.CoverTab[43233]++
						return m.ct
//line /usr/local/go/src/net/http/sniff.go:234
	// _ = "end of CoverTab[43233]"
}

type htmlSig []byte

func (h htmlSig) match(data []byte, firstNonWS int) string {
//line /usr/local/go/src/net/http/sniff.go:239
	_go_fuzz_dep_.CoverTab[43243]++
						data = data[firstNonWS:]
						if len(data) < len(h)+1 {
//line /usr/local/go/src/net/http/sniff.go:241
		_go_fuzz_dep_.CoverTab[43247]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:242
		// _ = "end of CoverTab[43247]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:243
		_go_fuzz_dep_.CoverTab[43248]++
//line /usr/local/go/src/net/http/sniff.go:243
		// _ = "end of CoverTab[43248]"
//line /usr/local/go/src/net/http/sniff.go:243
	}
//line /usr/local/go/src/net/http/sniff.go:243
	// _ = "end of CoverTab[43243]"
//line /usr/local/go/src/net/http/sniff.go:243
	_go_fuzz_dep_.CoverTab[43244]++
						for i, b := range h {
//line /usr/local/go/src/net/http/sniff.go:244
		_go_fuzz_dep_.CoverTab[43249]++
							db := data[i]
							if 'A' <= b && func() bool {
//line /usr/local/go/src/net/http/sniff.go:246
			_go_fuzz_dep_.CoverTab[43251]++
//line /usr/local/go/src/net/http/sniff.go:246
			return b <= 'Z'
//line /usr/local/go/src/net/http/sniff.go:246
			// _ = "end of CoverTab[43251]"
//line /usr/local/go/src/net/http/sniff.go:246
		}() {
//line /usr/local/go/src/net/http/sniff.go:246
			_go_fuzz_dep_.CoverTab[43252]++
								db &= 0xDF
//line /usr/local/go/src/net/http/sniff.go:247
			// _ = "end of CoverTab[43252]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:248
			_go_fuzz_dep_.CoverTab[43253]++
//line /usr/local/go/src/net/http/sniff.go:248
			// _ = "end of CoverTab[43253]"
//line /usr/local/go/src/net/http/sniff.go:248
		}
//line /usr/local/go/src/net/http/sniff.go:248
		// _ = "end of CoverTab[43249]"
//line /usr/local/go/src/net/http/sniff.go:248
		_go_fuzz_dep_.CoverTab[43250]++
							if b != db {
//line /usr/local/go/src/net/http/sniff.go:249
			_go_fuzz_dep_.CoverTab[43254]++
								return ""
//line /usr/local/go/src/net/http/sniff.go:250
			// _ = "end of CoverTab[43254]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:251
			_go_fuzz_dep_.CoverTab[43255]++
//line /usr/local/go/src/net/http/sniff.go:251
			// _ = "end of CoverTab[43255]"
//line /usr/local/go/src/net/http/sniff.go:251
		}
//line /usr/local/go/src/net/http/sniff.go:251
		// _ = "end of CoverTab[43250]"
	}
//line /usr/local/go/src/net/http/sniff.go:252
	// _ = "end of CoverTab[43244]"
//line /usr/local/go/src/net/http/sniff.go:252
	_go_fuzz_dep_.CoverTab[43245]++

						if !isTT(data[len(h)]) {
//line /usr/local/go/src/net/http/sniff.go:254
		_go_fuzz_dep_.CoverTab[43256]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:255
		// _ = "end of CoverTab[43256]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:256
		_go_fuzz_dep_.CoverTab[43257]++
//line /usr/local/go/src/net/http/sniff.go:256
		// _ = "end of CoverTab[43257]"
//line /usr/local/go/src/net/http/sniff.go:256
	}
//line /usr/local/go/src/net/http/sniff.go:256
	// _ = "end of CoverTab[43245]"
//line /usr/local/go/src/net/http/sniff.go:256
	_go_fuzz_dep_.CoverTab[43246]++
						return "text/html; charset=utf-8"
//line /usr/local/go/src/net/http/sniff.go:257
	// _ = "end of CoverTab[43246]"
}

var mp4ftype = []byte("ftyp")
var mp4 = []byte("mp4")

type mp4Sig struct{}

func (mp4Sig) match(data []byte, firstNonWS int) string {
//line /usr/local/go/src/net/http/sniff.go:265
	_go_fuzz_dep_.CoverTab[43258]++

//line /usr/local/go/src/net/http/sniff.go:268
	if len(data) < 12 {
//line /usr/local/go/src/net/http/sniff.go:268
		_go_fuzz_dep_.CoverTab[43263]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:269
		// _ = "end of CoverTab[43263]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:270
		_go_fuzz_dep_.CoverTab[43264]++
//line /usr/local/go/src/net/http/sniff.go:270
		// _ = "end of CoverTab[43264]"
//line /usr/local/go/src/net/http/sniff.go:270
	}
//line /usr/local/go/src/net/http/sniff.go:270
	// _ = "end of CoverTab[43258]"
//line /usr/local/go/src/net/http/sniff.go:270
	_go_fuzz_dep_.CoverTab[43259]++
						boxSize := int(binary.BigEndian.Uint32(data[:4]))
						if len(data) < boxSize || func() bool {
//line /usr/local/go/src/net/http/sniff.go:272
		_go_fuzz_dep_.CoverTab[43265]++
//line /usr/local/go/src/net/http/sniff.go:272
		return boxSize%4 != 0
//line /usr/local/go/src/net/http/sniff.go:272
		// _ = "end of CoverTab[43265]"
//line /usr/local/go/src/net/http/sniff.go:272
	}() {
//line /usr/local/go/src/net/http/sniff.go:272
		_go_fuzz_dep_.CoverTab[43266]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:273
		// _ = "end of CoverTab[43266]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:274
		_go_fuzz_dep_.CoverTab[43267]++
//line /usr/local/go/src/net/http/sniff.go:274
		// _ = "end of CoverTab[43267]"
//line /usr/local/go/src/net/http/sniff.go:274
	}
//line /usr/local/go/src/net/http/sniff.go:274
	// _ = "end of CoverTab[43259]"
//line /usr/local/go/src/net/http/sniff.go:274
	_go_fuzz_dep_.CoverTab[43260]++
						if !bytes.Equal(data[4:8], mp4ftype) {
//line /usr/local/go/src/net/http/sniff.go:275
		_go_fuzz_dep_.CoverTab[43268]++
							return ""
//line /usr/local/go/src/net/http/sniff.go:276
		// _ = "end of CoverTab[43268]"
	} else {
//line /usr/local/go/src/net/http/sniff.go:277
		_go_fuzz_dep_.CoverTab[43269]++
//line /usr/local/go/src/net/http/sniff.go:277
		// _ = "end of CoverTab[43269]"
//line /usr/local/go/src/net/http/sniff.go:277
	}
//line /usr/local/go/src/net/http/sniff.go:277
	// _ = "end of CoverTab[43260]"
//line /usr/local/go/src/net/http/sniff.go:277
	_go_fuzz_dep_.CoverTab[43261]++
						for st := 8; st < boxSize; st += 4 {
//line /usr/local/go/src/net/http/sniff.go:278
		_go_fuzz_dep_.CoverTab[43270]++
							if st == 12 {
//line /usr/local/go/src/net/http/sniff.go:279
			_go_fuzz_dep_.CoverTab[43272]++

								continue
//line /usr/local/go/src/net/http/sniff.go:281
			// _ = "end of CoverTab[43272]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:282
			_go_fuzz_dep_.CoverTab[43273]++
//line /usr/local/go/src/net/http/sniff.go:282
			// _ = "end of CoverTab[43273]"
//line /usr/local/go/src/net/http/sniff.go:282
		}
//line /usr/local/go/src/net/http/sniff.go:282
		// _ = "end of CoverTab[43270]"
//line /usr/local/go/src/net/http/sniff.go:282
		_go_fuzz_dep_.CoverTab[43271]++
							if bytes.Equal(data[st:st+3], mp4) {
//line /usr/local/go/src/net/http/sniff.go:283
			_go_fuzz_dep_.CoverTab[43274]++
								return "video/mp4"
//line /usr/local/go/src/net/http/sniff.go:284
			// _ = "end of CoverTab[43274]"
		} else {
//line /usr/local/go/src/net/http/sniff.go:285
			_go_fuzz_dep_.CoverTab[43275]++
//line /usr/local/go/src/net/http/sniff.go:285
			// _ = "end of CoverTab[43275]"
//line /usr/local/go/src/net/http/sniff.go:285
		}
//line /usr/local/go/src/net/http/sniff.go:285
		// _ = "end of CoverTab[43271]"
	}
//line /usr/local/go/src/net/http/sniff.go:286
	// _ = "end of CoverTab[43261]"
//line /usr/local/go/src/net/http/sniff.go:286
	_go_fuzz_dep_.CoverTab[43262]++
						return ""
//line /usr/local/go/src/net/http/sniff.go:287
	// _ = "end of CoverTab[43262]"
}

type textSig struct{}

func (textSig) match(data []byte, firstNonWS int) string {
//line /usr/local/go/src/net/http/sniff.go:292
	_go_fuzz_dep_.CoverTab[43276]++

						for _, b := range data[firstNonWS:] {
//line /usr/local/go/src/net/http/sniff.go:294
		_go_fuzz_dep_.CoverTab[43278]++
							switch {
		case b <= 0x08,
			b == 0x0B,
			0x0E <= b && func() bool {
//line /usr/local/go/src/net/http/sniff.go:298
				_go_fuzz_dep_.CoverTab[43281]++
//line /usr/local/go/src/net/http/sniff.go:298
				return b <= 0x1A
//line /usr/local/go/src/net/http/sniff.go:298
				// _ = "end of CoverTab[43281]"
//line /usr/local/go/src/net/http/sniff.go:298
			}(),
			0x1C <= b && func() bool {
//line /usr/local/go/src/net/http/sniff.go:299
				_go_fuzz_dep_.CoverTab[43282]++
//line /usr/local/go/src/net/http/sniff.go:299
				return b <= 0x1F
//line /usr/local/go/src/net/http/sniff.go:299
				// _ = "end of CoverTab[43282]"
//line /usr/local/go/src/net/http/sniff.go:299
			}():
//line /usr/local/go/src/net/http/sniff.go:299
			_go_fuzz_dep_.CoverTab[43279]++
								return ""
//line /usr/local/go/src/net/http/sniff.go:300
			// _ = "end of CoverTab[43279]"
//line /usr/local/go/src/net/http/sniff.go:300
		default:
//line /usr/local/go/src/net/http/sniff.go:300
			_go_fuzz_dep_.CoverTab[43280]++
//line /usr/local/go/src/net/http/sniff.go:300
			// _ = "end of CoverTab[43280]"
		}
//line /usr/local/go/src/net/http/sniff.go:301
		// _ = "end of CoverTab[43278]"
	}
//line /usr/local/go/src/net/http/sniff.go:302
	// _ = "end of CoverTab[43276]"
//line /usr/local/go/src/net/http/sniff.go:302
	_go_fuzz_dep_.CoverTab[43277]++
						return "text/plain; charset=utf-8"
//line /usr/local/go/src/net/http/sniff.go:303
	// _ = "end of CoverTab[43277]"
}

//line /usr/local/go/src/net/http/sniff.go:304
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/sniff.go:304
var _ = _go_fuzz_dep_.CoverTab
