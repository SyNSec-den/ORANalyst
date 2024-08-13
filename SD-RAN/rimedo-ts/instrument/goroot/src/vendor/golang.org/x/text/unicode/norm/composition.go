// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:5
)

import "unicode/utf8"

const (
	maxNonStarters	= 30
	// The maximum number of characters needed for a buffer is
	// maxNonStarters + 1 for the starter + 1 for the GCJ
	maxBufferSize		= maxNonStarters + 2
	maxNFCExpansion		= 3	// NFC(0x1D160)
	maxNFKCExpansion	= 18	// NFKC(0xFDFA)

	maxByteBufferSize	= utf8.UTFMax * maxBufferSize	// 128
)

// ssState is used for reporting the segment state after inserting a rune.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:20
// It is returned by streamSafe.next.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:22
type ssState int

const (
	// Indicates a rune was successfully added to the segment.
	ssSuccess	ssState	= iota
	// Indicates a rune starts a new segment and should not be added.
	ssStarter
	// Indicates a rune caused a segment overflow and a CGJ should be inserted.
	ssOverflow
)

// streamSafe implements the policy of when a CGJ should be inserted.
type streamSafe uint8

// first inserts the first rune of a segment. It is a faster version of next if
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:36
// it is known p represents the first rune in a segment.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:38
func (ss *streamSafe) first(p Properties) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:38
	_go_fuzz_dep_.CoverTab[32932]++
											*ss = streamSafe(p.nTrailingNonStarters())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:39
	// _ = "end of CoverTab[32932]"
}

// insert returns a ssState value to indicate whether a rune represented by p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:42
// can be inserted.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:44
func (ss *streamSafe) next(p Properties) ssState {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:44
	_go_fuzz_dep_.CoverTab[32933]++
											if *ss > maxNonStarters {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:45
		_go_fuzz_dep_.CoverTab[32937]++
												panic("streamSafe was not reset")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:46
		// _ = "end of CoverTab[32937]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:47
		_go_fuzz_dep_.CoverTab[32938]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:47
		// _ = "end of CoverTab[32938]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:47
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:47
	// _ = "end of CoverTab[32933]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:47
	_go_fuzz_dep_.CoverTab[32934]++
											n := p.nLeadingNonStarters()
											if *ss += streamSafe(n); *ss > maxNonStarters {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:49
		_go_fuzz_dep_.CoverTab[32939]++
												*ss = 0
												return ssOverflow
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:51
		// _ = "end of CoverTab[32939]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:52
		_go_fuzz_dep_.CoverTab[32940]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:52
		// _ = "end of CoverTab[32940]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:52
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:52
	// _ = "end of CoverTab[32934]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:52
	_go_fuzz_dep_.CoverTab[32935]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:60
	if n == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:60
		_go_fuzz_dep_.CoverTab[32941]++
												*ss = streamSafe(p.nTrailingNonStarters())
												return ssStarter
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:62
		// _ = "end of CoverTab[32941]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:63
		_go_fuzz_dep_.CoverTab[32942]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:63
		// _ = "end of CoverTab[32942]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:63
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:63
	// _ = "end of CoverTab[32935]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:63
	_go_fuzz_dep_.CoverTab[32936]++
											return ssSuccess
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:64
	// _ = "end of CoverTab[32936]"
}

// backwards is used for checking for overflow and segment starts
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:67
// when traversing a string backwards. Users do not need to call first
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:67
// for the first rune. The state of the streamSafe retains the count of
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:67
// the non-starters loaded.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:71
func (ss *streamSafe) backwards(p Properties) ssState {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:71
	_go_fuzz_dep_.CoverTab[32943]++
											if *ss > maxNonStarters {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:72
		_go_fuzz_dep_.CoverTab[32947]++
												panic("streamSafe was not reset")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:73
		// _ = "end of CoverTab[32947]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:74
		_go_fuzz_dep_.CoverTab[32948]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:74
		// _ = "end of CoverTab[32948]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:74
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:74
	// _ = "end of CoverTab[32943]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:74
	_go_fuzz_dep_.CoverTab[32944]++
											c := *ss + streamSafe(p.nTrailingNonStarters())
											if c > maxNonStarters {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:76
		_go_fuzz_dep_.CoverTab[32949]++
												return ssOverflow
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:77
		// _ = "end of CoverTab[32949]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:78
		_go_fuzz_dep_.CoverTab[32950]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:78
		// _ = "end of CoverTab[32950]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:78
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:78
	// _ = "end of CoverTab[32944]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:78
	_go_fuzz_dep_.CoverTab[32945]++
											*ss = c
											if p.nLeadingNonStarters() == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:80
		_go_fuzz_dep_.CoverTab[32951]++
												return ssStarter
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:81
		// _ = "end of CoverTab[32951]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:82
		_go_fuzz_dep_.CoverTab[32952]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:82
		// _ = "end of CoverTab[32952]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:82
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:82
	// _ = "end of CoverTab[32945]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:82
	_go_fuzz_dep_.CoverTab[32946]++
											return ssSuccess
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:83
	// _ = "end of CoverTab[32946]"
}

func (ss streamSafe) isMax() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:86
	_go_fuzz_dep_.CoverTab[32953]++
											return ss == maxNonStarters
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:87
	// _ = "end of CoverTab[32953]"
}

// GraphemeJoiner is inserted after maxNonStarters non-starter runes.
const GraphemeJoiner = "\u034F"

// reorderBuffer is used to normalize a single segment.  Characters inserted with
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:93
// insert are decomposed and reordered based on CCC. The compose method can
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:93
// be used to recombine characters.  Note that the byte buffer does not hold
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:93
// the UTF-8 characters in order.  Only the rune array is maintained in sorted
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:93
// order. flush writes the resulting segment to a byte array.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:98
type reorderBuffer struct {
	rune	[maxBufferSize]Properties	// Per character info.
	byte	[maxByteBufferSize]byte		// UTF-8 buffer. Referenced by runeInfo.pos.
	nbyte	uint8				// Number or bytes.
	ss	streamSafe			// For limiting length of non-starter sequence.
	nrune	int				// Number of runeInfos.
	f	formInfo

	src		input
	nsrc		int
	tmpBytes	input

	out	[]byte
	flushF	func(*reorderBuffer) bool
}

func (rb *reorderBuffer) init(f Form, src []byte) {
	rb.f = *formTable[f]
	rb.src.setBytes(src)
	rb.nsrc = len(src)
	rb.ss = 0
}

func (rb *reorderBuffer) initString(f Form, src string) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:121
	_go_fuzz_dep_.CoverTab[32954]++
											rb.f = *formTable[f]
											rb.src.setString(src)
											rb.nsrc = len(src)
											rb.ss = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:125
	// _ = "end of CoverTab[32954]"
}

func (rb *reorderBuffer) setFlusher(out []byte, f func(*reorderBuffer) bool) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:128
	_go_fuzz_dep_.CoverTab[32955]++
											rb.out = out
											rb.flushF = f
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:130
	// _ = "end of CoverTab[32955]"
}

// reset discards all characters from the buffer.
func (rb *reorderBuffer) reset() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:134
	_go_fuzz_dep_.CoverTab[32956]++
											rb.nrune = 0
											rb.nbyte = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:136
	// _ = "end of CoverTab[32956]"
}

func (rb *reorderBuffer) doFlush() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:139
	_go_fuzz_dep_.CoverTab[32957]++
											if rb.f.composing {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:140
		_go_fuzz_dep_.CoverTab[32959]++
												rb.compose()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:141
		// _ = "end of CoverTab[32959]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:142
		_go_fuzz_dep_.CoverTab[32960]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:142
		// _ = "end of CoverTab[32960]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:142
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:142
	// _ = "end of CoverTab[32957]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:142
	_go_fuzz_dep_.CoverTab[32958]++
											res := rb.flushF(rb)
											rb.reset()
											return res
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:145
	// _ = "end of CoverTab[32958]"
}

// appendFlush appends the normalized segment to rb.out.
func appendFlush(rb *reorderBuffer) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:149
	_go_fuzz_dep_.CoverTab[32961]++
											for i := 0; i < rb.nrune; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:150
		_go_fuzz_dep_.CoverTab[32963]++
												start := rb.rune[i].pos
												end := start + rb.rune[i].size
												rb.out = append(rb.out, rb.byte[start:end]...)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:153
		// _ = "end of CoverTab[32963]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:154
	// _ = "end of CoverTab[32961]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:154
	_go_fuzz_dep_.CoverTab[32962]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:155
	// _ = "end of CoverTab[32962]"
}

// flush appends the normalized segment to out and resets rb.
func (rb *reorderBuffer) flush(out []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:159
	_go_fuzz_dep_.CoverTab[32964]++
											for i := 0; i < rb.nrune; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:160
		_go_fuzz_dep_.CoverTab[32966]++
												start := rb.rune[i].pos
												end := start + rb.rune[i].size
												out = append(out, rb.byte[start:end]...)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:163
		// _ = "end of CoverTab[32966]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:164
	// _ = "end of CoverTab[32964]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:164
	_go_fuzz_dep_.CoverTab[32965]++
											rb.reset()
											return out
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:166
	// _ = "end of CoverTab[32965]"
}

// flushCopy copies the normalized segment to buf and resets rb.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:169
// It returns the number of bytes written to buf.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:171
func (rb *reorderBuffer) flushCopy(buf []byte) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:171
	_go_fuzz_dep_.CoverTab[32967]++
											p := 0
											for i := 0; i < rb.nrune; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:173
		_go_fuzz_dep_.CoverTab[32969]++
												runep := rb.rune[i]
												p += copy(buf[p:], rb.byte[runep.pos:runep.pos+runep.size])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:175
		// _ = "end of CoverTab[32969]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:176
	// _ = "end of CoverTab[32967]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:176
	_go_fuzz_dep_.CoverTab[32968]++
											rb.reset()
											return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:178
	// _ = "end of CoverTab[32968]"
}

// insertOrdered inserts a rune in the buffer, ordered by Canonical Combining Class.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:181
// It returns false if the buffer is not large enough to hold the rune.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:181
// It is used internally by insert and insertString only.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:184
func (rb *reorderBuffer) insertOrdered(info Properties) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:184
	_go_fuzz_dep_.CoverTab[32970]++
											n := rb.nrune
											b := rb.rune[:]
											cc := info.ccc
											if cc > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:188
		_go_fuzz_dep_.CoverTab[32972]++

												for ; n > 0; n-- {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:190
			_go_fuzz_dep_.CoverTab[32973]++
													if b[n-1].ccc <= cc {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:191
				_go_fuzz_dep_.CoverTab[32975]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:192
				// _ = "end of CoverTab[32975]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:193
				_go_fuzz_dep_.CoverTab[32976]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:193
				// _ = "end of CoverTab[32976]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:193
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:193
			// _ = "end of CoverTab[32973]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:193
			_go_fuzz_dep_.CoverTab[32974]++
													b[n] = b[n-1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:194
			// _ = "end of CoverTab[32974]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:195
		// _ = "end of CoverTab[32972]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:196
		_go_fuzz_dep_.CoverTab[32977]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:196
		// _ = "end of CoverTab[32977]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:196
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:196
	// _ = "end of CoverTab[32970]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:196
	_go_fuzz_dep_.CoverTab[32971]++
											rb.nrune += 1
											pos := uint8(rb.nbyte)
											rb.nbyte += utf8.UTFMax
											info.pos = pos
											b[n] = info
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:201
	// _ = "end of CoverTab[32971]"
}

// insertErr is an error code returned by insert. Using this type instead
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:204
// of error improves performance up to 20% for many of the benchmarks.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:206
type insertErr int

const (
	iSuccess	insertErr	= -iota
	iShortDst
	iShortSrc
)

// insertFlush inserts the given rune in the buffer ordered by CCC.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:214
// If a decomposition with multiple segments are encountered, they leading
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:214
// ones are flushed.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:214
// It returns a non-zero error code if the rune was not inserted.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:218
func (rb *reorderBuffer) insertFlush(src input, i int, info Properties) insertErr {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:218
	_go_fuzz_dep_.CoverTab[32978]++
											if rune := src.hangul(i); rune != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:219
		_go_fuzz_dep_.CoverTab[32981]++
												rb.decomposeHangul(rune)
												return iSuccess
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:221
		// _ = "end of CoverTab[32981]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:222
		_go_fuzz_dep_.CoverTab[32982]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:222
		// _ = "end of CoverTab[32982]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:222
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:222
	// _ = "end of CoverTab[32978]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:222
	_go_fuzz_dep_.CoverTab[32979]++
											if info.hasDecomposition() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:223
		_go_fuzz_dep_.CoverTab[32983]++
												return rb.insertDecomposed(info.Decomposition())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:224
		// _ = "end of CoverTab[32983]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:225
		_go_fuzz_dep_.CoverTab[32984]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:225
		// _ = "end of CoverTab[32984]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:225
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:225
	// _ = "end of CoverTab[32979]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:225
	_go_fuzz_dep_.CoverTab[32980]++
											rb.insertSingle(src, i, info)
											return iSuccess
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:227
	// _ = "end of CoverTab[32980]"
}

// insertUnsafe inserts the given rune in the buffer ordered by CCC.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:230
// It is assumed there is sufficient space to hold the runes. It is the
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:230
// responsibility of the caller to ensure this. This can be done by checking
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:230
// the state returned by the streamSafe type.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:234
func (rb *reorderBuffer) insertUnsafe(src input, i int, info Properties) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:234
	_go_fuzz_dep_.CoverTab[32985]++
											if rune := src.hangul(i); rune != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:235
		_go_fuzz_dep_.CoverTab[32987]++
												rb.decomposeHangul(rune)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:236
		// _ = "end of CoverTab[32987]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:237
		_go_fuzz_dep_.CoverTab[32988]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:237
		// _ = "end of CoverTab[32988]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:237
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:237
	// _ = "end of CoverTab[32985]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:237
	_go_fuzz_dep_.CoverTab[32986]++
											if info.hasDecomposition() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:238
		_go_fuzz_dep_.CoverTab[32989]++

												rb.insertDecomposed(info.Decomposition())
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:240
		// _ = "end of CoverTab[32989]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:241
		_go_fuzz_dep_.CoverTab[32990]++
												rb.insertSingle(src, i, info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:242
		// _ = "end of CoverTab[32990]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:243
	// _ = "end of CoverTab[32986]"
}

// insertDecomposed inserts an entry in to the reorderBuffer for each rune
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:246
// in dcomp. dcomp must be a sequence of decomposed UTF-8-encoded runes.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:246
// It flushes the buffer on each new segment start.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:249
func (rb *reorderBuffer) insertDecomposed(dcomp []byte) insertErr {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:249
	_go_fuzz_dep_.CoverTab[32991]++
											rb.tmpBytes.setBytes(dcomp)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:254
	for i := 0; i < len(dcomp); {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:254
		_go_fuzz_dep_.CoverTab[32993]++
												info := rb.f.info(rb.tmpBytes, i)
												if info.BoundaryBefore() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[32995]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			return rb.nrune > 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			// _ = "end of CoverTab[32995]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
		}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[32996]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			return !rb.doFlush()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			// _ = "end of CoverTab[32996]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[32997]++
													return iShortDst
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:257
			// _ = "end of CoverTab[32997]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:258
			_go_fuzz_dep_.CoverTab[32998]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:258
			// _ = "end of CoverTab[32998]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:258
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:258
		// _ = "end of CoverTab[32993]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:258
		_go_fuzz_dep_.CoverTab[32994]++
												i += copy(rb.byte[rb.nbyte:], dcomp[i:i+int(info.size)])
												rb.insertOrdered(info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:260
		// _ = "end of CoverTab[32994]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:261
	// _ = "end of CoverTab[32991]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:261
	_go_fuzz_dep_.CoverTab[32992]++
											return iSuccess
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:262
	// _ = "end of CoverTab[32992]"
}

// insertSingle inserts an entry in the reorderBuffer for the rune at
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:265
// position i. info is the runeInfo for the rune at position i.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:267
func (rb *reorderBuffer) insertSingle(src input, i int, info Properties) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:267
	_go_fuzz_dep_.CoverTab[32999]++
											src.copySlice(rb.byte[rb.nbyte:], i, i+int(info.size))
											rb.insertOrdered(info)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:269
	// _ = "end of CoverTab[32999]"
}

// insertCGJ inserts a Combining Grapheme Joiner (0x034f) into rb.
func (rb *reorderBuffer) insertCGJ() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:273
	_go_fuzz_dep_.CoverTab[33000]++
											rb.insertSingle(input{str: GraphemeJoiner}, 0, Properties{size: uint8(len(GraphemeJoiner))})
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:274
	// _ = "end of CoverTab[33000]"
}

// appendRune inserts a rune at the end of the buffer. It is used for Hangul.
func (rb *reorderBuffer) appendRune(r rune) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:278
	_go_fuzz_dep_.CoverTab[33001]++
											bn := rb.nbyte
											sz := utf8.EncodeRune(rb.byte[bn:], rune(r))
											rb.nbyte += utf8.UTFMax
											rb.rune[rb.nrune] = Properties{pos: bn, size: uint8(sz)}
											rb.nrune++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:283
	// _ = "end of CoverTab[33001]"
}

// assignRune sets a rune at position pos. It is used for Hangul and recomposition.
func (rb *reorderBuffer) assignRune(pos int, r rune) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:287
	_go_fuzz_dep_.CoverTab[33002]++
											bn := rb.rune[pos].pos
											sz := utf8.EncodeRune(rb.byte[bn:], rune(r))
											rb.rune[pos] = Properties{pos: bn, size: uint8(sz)}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:290
	// _ = "end of CoverTab[33002]"
}

// runeAt returns the rune at position n. It is used for Hangul and recomposition.
func (rb *reorderBuffer) runeAt(n int) rune {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:294
	_go_fuzz_dep_.CoverTab[33003]++
											inf := rb.rune[n]
											r, _ := utf8.DecodeRune(rb.byte[inf.pos : inf.pos+inf.size])
											return r
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:297
	// _ = "end of CoverTab[33003]"
}

// bytesAt returns the UTF-8 encoding of the rune at position n.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:300
// It is used for Hangul and recomposition.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:302
func (rb *reorderBuffer) bytesAt(n int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:302
	_go_fuzz_dep_.CoverTab[33004]++
											inf := rb.rune[n]
											return rb.byte[inf.pos : int(inf.pos)+int(inf.size)]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:304
	// _ = "end of CoverTab[33004]"
}

// For Hangul we combine algorithmically, instead of using tables.
const (
	hangulBase	= 0xAC00	// UTF-8(hangulBase) -> EA B0 80
	hangulBase0	= 0xEA
	hangulBase1	= 0xB0
	hangulBase2	= 0x80

	hangulEnd	= hangulBase + jamoLVTCount	// UTF-8(0xD7A4) -> ED 9E A4
	hangulEnd0	= 0xED
	hangulEnd1	= 0x9E
	hangulEnd2	= 0xA4

	jamoLBase	= 0x1100	// UTF-8(jamoLBase) -> E1 84 00
	jamoLBase0	= 0xE1
	jamoLBase1	= 0x84
	jamoLEnd	= 0x1113
	jamoVBase	= 0x1161
	jamoVEnd	= 0x1176
	jamoTBase	= 0x11A7
	jamoTEnd	= 0x11C3

	jamoTCount	= 28
	jamoVCount	= 21
	jamoVTCount	= 21 * 28
	jamoLVTCount	= 19 * 21 * 28
)

const hangulUTF8Size = 3

func isHangul(b []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:336
	_go_fuzz_dep_.CoverTab[33005]++
											if len(b) < hangulUTF8Size {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:337
		_go_fuzz_dep_.CoverTab[33009]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:338
		// _ = "end of CoverTab[33009]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:339
		_go_fuzz_dep_.CoverTab[33010]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:339
		// _ = "end of CoverTab[33010]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:339
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:339
	// _ = "end of CoverTab[33005]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:339
	_go_fuzz_dep_.CoverTab[33006]++
											b0 := b[0]
											if b0 < hangulBase0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:341
		_go_fuzz_dep_.CoverTab[33011]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:342
		// _ = "end of CoverTab[33011]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:343
		_go_fuzz_dep_.CoverTab[33012]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:343
		// _ = "end of CoverTab[33012]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:343
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:343
	// _ = "end of CoverTab[33006]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:343
	_go_fuzz_dep_.CoverTab[33007]++
											b1 := b[1]
											switch {
	case b0 == hangulBase0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:346
		_go_fuzz_dep_.CoverTab[33013]++
												return b1 >= hangulBase1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:347
		// _ = "end of CoverTab[33013]"
	case b0 < hangulEnd0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:348
		_go_fuzz_dep_.CoverTab[33014]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:349
		// _ = "end of CoverTab[33014]"
	case b0 > hangulEnd0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:350
		_go_fuzz_dep_.CoverTab[33015]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:351
		// _ = "end of CoverTab[33015]"
	case b1 < hangulEnd1:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:352
		_go_fuzz_dep_.CoverTab[33016]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:353
		// _ = "end of CoverTab[33016]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:353
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:353
		_go_fuzz_dep_.CoverTab[33017]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:353
		// _ = "end of CoverTab[33017]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:354
	// _ = "end of CoverTab[33007]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:354
	_go_fuzz_dep_.CoverTab[33008]++
											return b1 == hangulEnd1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:355
		_go_fuzz_dep_.CoverTab[33018]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:355
		return b[2] < hangulEnd2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:355
		// _ = "end of CoverTab[33018]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:355
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:355
	// _ = "end of CoverTab[33008]"
}

func isHangulString(b string) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:358
	_go_fuzz_dep_.CoverTab[33019]++
											if len(b) < hangulUTF8Size {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:359
		_go_fuzz_dep_.CoverTab[33023]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:360
		// _ = "end of CoverTab[33023]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:361
		_go_fuzz_dep_.CoverTab[33024]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:361
		// _ = "end of CoverTab[33024]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:361
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:361
	// _ = "end of CoverTab[33019]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:361
	_go_fuzz_dep_.CoverTab[33020]++
											b0 := b[0]
											if b0 < hangulBase0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:363
		_go_fuzz_dep_.CoverTab[33025]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:364
		// _ = "end of CoverTab[33025]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:365
		_go_fuzz_dep_.CoverTab[33026]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:365
		// _ = "end of CoverTab[33026]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:365
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:365
	// _ = "end of CoverTab[33020]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:365
	_go_fuzz_dep_.CoverTab[33021]++
											b1 := b[1]
											switch {
	case b0 == hangulBase0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:368
		_go_fuzz_dep_.CoverTab[33027]++
												return b1 >= hangulBase1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:369
		// _ = "end of CoverTab[33027]"
	case b0 < hangulEnd0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:370
		_go_fuzz_dep_.CoverTab[33028]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:371
		// _ = "end of CoverTab[33028]"
	case b0 > hangulEnd0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:372
		_go_fuzz_dep_.CoverTab[33029]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:373
		// _ = "end of CoverTab[33029]"
	case b1 < hangulEnd1:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:374
		_go_fuzz_dep_.CoverTab[33030]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:375
		// _ = "end of CoverTab[33030]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:375
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:375
		_go_fuzz_dep_.CoverTab[33031]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:375
		// _ = "end of CoverTab[33031]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:376
	// _ = "end of CoverTab[33021]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:376
	_go_fuzz_dep_.CoverTab[33022]++
											return b1 == hangulEnd1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:377
		_go_fuzz_dep_.CoverTab[33032]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:377
		return b[2] < hangulEnd2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:377
		// _ = "end of CoverTab[33032]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:377
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:377
	// _ = "end of CoverTab[33022]"
}

// Caller must ensure len(b) >= 2.
func isJamoVT(b []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:381
	_go_fuzz_dep_.CoverTab[33033]++

											return b[0] == jamoLBase0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:383
		_go_fuzz_dep_.CoverTab[33034]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:383
		return (b[1] & 0xFC) == jamoLBase1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:383
		// _ = "end of CoverTab[33034]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:383
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:383
	// _ = "end of CoverTab[33033]"
}

func isHangulWithoutJamoT(b []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:386
	_go_fuzz_dep_.CoverTab[33035]++
											c, _ := utf8.DecodeRune(b)
											c -= hangulBase
											return c < jamoLVTCount && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:389
		_go_fuzz_dep_.CoverTab[33036]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:389
		return c%jamoTCount == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:389
		// _ = "end of CoverTab[33036]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:389
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:389
	// _ = "end of CoverTab[33035]"
}

// decomposeHangul writes the decomposed Hangul to buf and returns the number
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:392
// of bytes written.  len(buf) should be at least 9.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:394
func decomposeHangul(buf []byte, r rune) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:394
	_go_fuzz_dep_.CoverTab[33037]++
											const JamoUTF8Len = 3
											r -= hangulBase
											x := r % jamoTCount
											r /= jamoTCount
											utf8.EncodeRune(buf, jamoLBase+r/jamoVCount)
											utf8.EncodeRune(buf[JamoUTF8Len:], jamoVBase+r%jamoVCount)
											if x != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:401
		_go_fuzz_dep_.CoverTab[33039]++
												utf8.EncodeRune(buf[2*JamoUTF8Len:], jamoTBase+x)
												return 3 * JamoUTF8Len
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:403
		// _ = "end of CoverTab[33039]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:404
		_go_fuzz_dep_.CoverTab[33040]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:404
		// _ = "end of CoverTab[33040]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:404
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:404
	// _ = "end of CoverTab[33037]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:404
	_go_fuzz_dep_.CoverTab[33038]++
											return 2 * JamoUTF8Len
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:405
	// _ = "end of CoverTab[33038]"
}

// decomposeHangul algorithmically decomposes a Hangul rune into
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:408
// its Jamo components.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:408
// See https://unicode.org/reports/tr15/#Hangul for details on decomposing Hangul.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:411
func (rb *reorderBuffer) decomposeHangul(r rune) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:411
	_go_fuzz_dep_.CoverTab[33041]++
											r -= hangulBase
											x := r % jamoTCount
											r /= jamoTCount
											rb.appendRune(jamoLBase + r/jamoVCount)
											rb.appendRune(jamoVBase + r%jamoVCount)
											if x != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:417
		_go_fuzz_dep_.CoverTab[33042]++
												rb.appendRune(jamoTBase + x)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:418
		// _ = "end of CoverTab[33042]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:419
		_go_fuzz_dep_.CoverTab[33043]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:419
		// _ = "end of CoverTab[33043]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:419
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:419
	// _ = "end of CoverTab[33041]"
}

// combineHangul algorithmically combines Jamo character components into Hangul.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:422
// See https://unicode.org/reports/tr15/#Hangul for details on combining Hangul.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:424
func (rb *reorderBuffer) combineHangul(s, i, k int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:424
	_go_fuzz_dep_.CoverTab[33044]++
											b := rb.rune[:]
											bn := rb.nrune
											for ; i < bn; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:427
		_go_fuzz_dep_.CoverTab[33046]++
												cccB := b[k-1].ccc
												cccC := b[i].ccc
												if cccB == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:430
			_go_fuzz_dep_.CoverTab[33048]++
													s = k - 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:431
			// _ = "end of CoverTab[33048]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:432
			_go_fuzz_dep_.CoverTab[33049]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:432
			// _ = "end of CoverTab[33049]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:432
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:432
		// _ = "end of CoverTab[33046]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:432
		_go_fuzz_dep_.CoverTab[33047]++
												if s != k-1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:433
			_go_fuzz_dep_.CoverTab[33050]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:433
			return cccB >= cccC
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:433
			// _ = "end of CoverTab[33050]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:433
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:433
			_go_fuzz_dep_.CoverTab[33051]++

													b[k] = b[i]
													k++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:436
			// _ = "end of CoverTab[33051]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:437
			_go_fuzz_dep_.CoverTab[33052]++
													l := rb.runeAt(s)
													v := rb.runeAt(i)
													switch {
			case jamoLBase <= l && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
				_go_fuzz_dep_.CoverTab[33056]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
				return l < jamoLEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
				// _ = "end of CoverTab[33056]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
				_go_fuzz_dep_.CoverTab[33057]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:441
				return jamoVBase <= v
														// _ = "end of CoverTab[33057]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
				_go_fuzz_dep_.CoverTab[33058]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
				return v < jamoVEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
				// _ = "end of CoverTab[33058]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
			}():
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:442
				_go_fuzz_dep_.CoverTab[33053]++

														rb.assignRune(s, hangulBase+
					(l-jamoLBase)*jamoVTCount+(v-jamoVBase)*jamoTCount)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:445
				// _ = "end of CoverTab[33053]"
			case hangulBase <= l && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
				_go_fuzz_dep_.CoverTab[33059]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
				return l < hangulEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
				// _ = "end of CoverTab[33059]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
				_go_fuzz_dep_.CoverTab[33060]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:446
				return jamoTBase < v
														// _ = "end of CoverTab[33060]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
				_go_fuzz_dep_.CoverTab[33061]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
				return v < jamoTEnd
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
				// _ = "end of CoverTab[33061]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
				_go_fuzz_dep_.CoverTab[33062]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:447
				return ((l - hangulBase) % jamoTCount) == 0
														// _ = "end of CoverTab[33062]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:448
			}():
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:448
				_go_fuzz_dep_.CoverTab[33054]++

														rb.assignRune(s, l+v-jamoTBase)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:450
				// _ = "end of CoverTab[33054]"
			default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:451
				_go_fuzz_dep_.CoverTab[33055]++
														b[k] = b[i]
														k++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:453
				// _ = "end of CoverTab[33055]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:454
			// _ = "end of CoverTab[33052]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:455
		// _ = "end of CoverTab[33047]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:456
	// _ = "end of CoverTab[33044]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:456
	_go_fuzz_dep_.CoverTab[33045]++
											rb.nrune = k
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:457
	// _ = "end of CoverTab[33045]"
}

// compose recombines the runes in the buffer.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:460
// It should only be used to recompose a single segment, as it will not
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:460
// handle alternations between Hangul and non-Hangul characters correctly.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:463
func (rb *reorderBuffer) compose() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:463
	_go_fuzz_dep_.CoverTab[33063]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:466
	recompMapOnce.Do(buildRecompMap)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:473
	bn := rb.nrune
	if bn == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:474
		_go_fuzz_dep_.CoverTab[33066]++
												return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:475
		// _ = "end of CoverTab[33066]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:476
		_go_fuzz_dep_.CoverTab[33067]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:476
		// _ = "end of CoverTab[33067]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:476
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:476
	// _ = "end of CoverTab[33063]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:476
	_go_fuzz_dep_.CoverTab[33064]++
											k := 1
											b := rb.rune[:]
											for s, i := 0, 1; i < bn; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:479
		_go_fuzz_dep_.CoverTab[33068]++
												if isJamoVT(rb.bytesAt(i)) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:480
			_go_fuzz_dep_.CoverTab[33071]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:483
			rb.combineHangul(s, i, k)
													return
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:484
			// _ = "end of CoverTab[33071]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:485
			_go_fuzz_dep_.CoverTab[33072]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:485
			// _ = "end of CoverTab[33072]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:485
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:485
		// _ = "end of CoverTab[33068]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:485
		_go_fuzz_dep_.CoverTab[33069]++
												ii := b[i]

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:491
		if ii.combinesBackward() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:491
			_go_fuzz_dep_.CoverTab[33073]++
													cccB := b[k-1].ccc
													cccC := ii.ccc
													blocked := false
													if cccB == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:495
				_go_fuzz_dep_.CoverTab[33075]++
														s = k - 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:496
				// _ = "end of CoverTab[33075]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:497
				_go_fuzz_dep_.CoverTab[33076]++
														blocked = s != k-1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:498
					_go_fuzz_dep_.CoverTab[33077]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:498
					return cccB >= cccC
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:498
					// _ = "end of CoverTab[33077]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:498
				}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:498
				// _ = "end of CoverTab[33076]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:499
			// _ = "end of CoverTab[33073]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:499
			_go_fuzz_dep_.CoverTab[33074]++
													if !blocked {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:500
				_go_fuzz_dep_.CoverTab[33078]++
														combined := combine(rb.runeAt(s), rb.runeAt(i))
														if combined != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:502
					_go_fuzz_dep_.CoverTab[33079]++
															rb.assignRune(s, combined)
															continue
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:504
					// _ = "end of CoverTab[33079]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:505
					_go_fuzz_dep_.CoverTab[33080]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:505
					// _ = "end of CoverTab[33080]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:505
				}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:505
				// _ = "end of CoverTab[33078]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:506
				_go_fuzz_dep_.CoverTab[33081]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:506
				// _ = "end of CoverTab[33081]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:506
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:506
			// _ = "end of CoverTab[33074]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:507
			_go_fuzz_dep_.CoverTab[33082]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:507
			// _ = "end of CoverTab[33082]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:507
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:507
		// _ = "end of CoverTab[33069]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:507
		_go_fuzz_dep_.CoverTab[33070]++
												b[k] = b[i]
												k++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:509
		// _ = "end of CoverTab[33070]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:510
	// _ = "end of CoverTab[33064]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:510
	_go_fuzz_dep_.CoverTab[33065]++
											rb.nrune = k
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:511
	// _ = "end of CoverTab[33065]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:512
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/composition.go:512
var _ = _go_fuzz_dep_.CoverTab
