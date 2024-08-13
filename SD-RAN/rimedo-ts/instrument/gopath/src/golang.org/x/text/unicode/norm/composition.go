// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:5
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:20
// It is returned by streamSafe.next.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:22
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:36
// it is known p represents the first rune in a segment.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:38
func (ss *streamSafe) first(p Properties) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:38
	_go_fuzz_dep_.CoverTab[70213]++
												*ss = streamSafe(p.nTrailingNonStarters())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:39
	// _ = "end of CoverTab[70213]"
}

// insert returns a ssState value to indicate whether a rune represented by p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:42
// can be inserted.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:44
func (ss *streamSafe) next(p Properties) ssState {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:44
	_go_fuzz_dep_.CoverTab[70214]++
												if *ss > maxNonStarters {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:45
		_go_fuzz_dep_.CoverTab[70218]++
													panic("streamSafe was not reset")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:46
		// _ = "end of CoverTab[70218]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:47
		_go_fuzz_dep_.CoverTab[70219]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:47
		// _ = "end of CoverTab[70219]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:47
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:47
	// _ = "end of CoverTab[70214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:47
	_go_fuzz_dep_.CoverTab[70215]++
												n := p.nLeadingNonStarters()
												if *ss += streamSafe(n); *ss > maxNonStarters {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:49
		_go_fuzz_dep_.CoverTab[70220]++
													*ss = 0
													return ssOverflow
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:51
		// _ = "end of CoverTab[70220]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:52
		_go_fuzz_dep_.CoverTab[70221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:52
		// _ = "end of CoverTab[70221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:52
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:52
	// _ = "end of CoverTab[70215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:52
	_go_fuzz_dep_.CoverTab[70216]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:60
	if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:60
		_go_fuzz_dep_.CoverTab[70222]++
													*ss = streamSafe(p.nTrailingNonStarters())
													return ssStarter
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:62
		// _ = "end of CoverTab[70222]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:63
		_go_fuzz_dep_.CoverTab[70223]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:63
		// _ = "end of CoverTab[70223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:63
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:63
	// _ = "end of CoverTab[70216]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:63
	_go_fuzz_dep_.CoverTab[70217]++
												return ssSuccess
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:64
	// _ = "end of CoverTab[70217]"
}

// backwards is used for checking for overflow and segment starts
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:67
// when traversing a string backwards. Users do not need to call first
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:67
// for the first rune. The state of the streamSafe retains the count of
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:67
// the non-starters loaded.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:71
func (ss *streamSafe) backwards(p Properties) ssState {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:71
	_go_fuzz_dep_.CoverTab[70224]++
												if *ss > maxNonStarters {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:72
		_go_fuzz_dep_.CoverTab[70228]++
													panic("streamSafe was not reset")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:73
		// _ = "end of CoverTab[70228]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:74
		_go_fuzz_dep_.CoverTab[70229]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:74
		// _ = "end of CoverTab[70229]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:74
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:74
	// _ = "end of CoverTab[70224]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:74
	_go_fuzz_dep_.CoverTab[70225]++
												c := *ss + streamSafe(p.nTrailingNonStarters())
												if c > maxNonStarters {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:76
		_go_fuzz_dep_.CoverTab[70230]++
													return ssOverflow
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:77
		// _ = "end of CoverTab[70230]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:78
		_go_fuzz_dep_.CoverTab[70231]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:78
		// _ = "end of CoverTab[70231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:78
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:78
	// _ = "end of CoverTab[70225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:78
	_go_fuzz_dep_.CoverTab[70226]++
												*ss = c
												if p.nLeadingNonStarters() == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:80
		_go_fuzz_dep_.CoverTab[70232]++
													return ssStarter
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:81
		// _ = "end of CoverTab[70232]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:82
		_go_fuzz_dep_.CoverTab[70233]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:82
		// _ = "end of CoverTab[70233]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:82
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:82
	// _ = "end of CoverTab[70226]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:82
	_go_fuzz_dep_.CoverTab[70227]++
												return ssSuccess
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:83
	// _ = "end of CoverTab[70227]"
}

func (ss streamSafe) isMax() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:86
	_go_fuzz_dep_.CoverTab[70234]++
												return ss == maxNonStarters
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:87
	// _ = "end of CoverTab[70234]"
}

// GraphemeJoiner is inserted after maxNonStarters non-starter runes.
const GraphemeJoiner = "\u034F"

// reorderBuffer is used to normalize a single segment.  Characters inserted with
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:93
// insert are decomposed and reordered based on CCC. The compose method can
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:93
// be used to recombine characters.  Note that the byte buffer does not hold
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:93
// the UTF-8 characters in order.  Only the rune array is maintained in sorted
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:93
// order. flush writes the resulting segment to a byte array.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:98
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:121
	_go_fuzz_dep_.CoverTab[70235]++
												rb.f = *formTable[f]
												rb.src.setString(src)
												rb.nsrc = len(src)
												rb.ss = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:125
	// _ = "end of CoverTab[70235]"
}

func (rb *reorderBuffer) setFlusher(out []byte, f func(*reorderBuffer) bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:128
	_go_fuzz_dep_.CoverTab[70236]++
												rb.out = out
												rb.flushF = f
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:130
	// _ = "end of CoverTab[70236]"
}

// reset discards all characters from the buffer.
func (rb *reorderBuffer) reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:134
	_go_fuzz_dep_.CoverTab[70237]++
												rb.nrune = 0
												rb.nbyte = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:136
	// _ = "end of CoverTab[70237]"
}

func (rb *reorderBuffer) doFlush() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:139
	_go_fuzz_dep_.CoverTab[70238]++
												if rb.f.composing {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:140
		_go_fuzz_dep_.CoverTab[70240]++
													rb.compose()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:141
		// _ = "end of CoverTab[70240]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:142
		_go_fuzz_dep_.CoverTab[70241]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:142
		// _ = "end of CoverTab[70241]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:142
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:142
	// _ = "end of CoverTab[70238]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:142
	_go_fuzz_dep_.CoverTab[70239]++
												res := rb.flushF(rb)
												rb.reset()
												return res
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:145
	// _ = "end of CoverTab[70239]"
}

// appendFlush appends the normalized segment to rb.out.
func appendFlush(rb *reorderBuffer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:149
	_go_fuzz_dep_.CoverTab[70242]++
												for i := 0; i < rb.nrune; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:150
		_go_fuzz_dep_.CoverTab[70244]++
													start := rb.rune[i].pos
													end := start + rb.rune[i].size
													rb.out = append(rb.out, rb.byte[start:end]...)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:153
		// _ = "end of CoverTab[70244]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:154
	// _ = "end of CoverTab[70242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:154
	_go_fuzz_dep_.CoverTab[70243]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:155
	// _ = "end of CoverTab[70243]"
}

// flush appends the normalized segment to out and resets rb.
func (rb *reorderBuffer) flush(out []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:159
	_go_fuzz_dep_.CoverTab[70245]++
												for i := 0; i < rb.nrune; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:160
		_go_fuzz_dep_.CoverTab[70247]++
													start := rb.rune[i].pos
													end := start + rb.rune[i].size
													out = append(out, rb.byte[start:end]...)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:163
		// _ = "end of CoverTab[70247]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:164
	// _ = "end of CoverTab[70245]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:164
	_go_fuzz_dep_.CoverTab[70246]++
												rb.reset()
												return out
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:166
	// _ = "end of CoverTab[70246]"
}

// flushCopy copies the normalized segment to buf and resets rb.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:169
// It returns the number of bytes written to buf.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:171
func (rb *reorderBuffer) flushCopy(buf []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:171
	_go_fuzz_dep_.CoverTab[70248]++
												p := 0
												for i := 0; i < rb.nrune; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:173
		_go_fuzz_dep_.CoverTab[70250]++
													runep := rb.rune[i]
													p += copy(buf[p:], rb.byte[runep.pos:runep.pos+runep.size])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:175
		// _ = "end of CoverTab[70250]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:176
	// _ = "end of CoverTab[70248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:176
	_go_fuzz_dep_.CoverTab[70249]++
												rb.reset()
												return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:178
	// _ = "end of CoverTab[70249]"
}

// insertOrdered inserts a rune in the buffer, ordered by Canonical Combining Class.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:181
// It returns false if the buffer is not large enough to hold the rune.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:181
// It is used internally by insert and insertString only.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:184
func (rb *reorderBuffer) insertOrdered(info Properties) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:184
	_go_fuzz_dep_.CoverTab[70251]++
												n := rb.nrune
												b := rb.rune[:]
												cc := info.ccc
												if cc > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:188
		_go_fuzz_dep_.CoverTab[70253]++

													for ; n > 0; n-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:190
			_go_fuzz_dep_.CoverTab[70254]++
														if b[n-1].ccc <= cc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:191
				_go_fuzz_dep_.CoverTab[70256]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:192
				// _ = "end of CoverTab[70256]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:193
				_go_fuzz_dep_.CoverTab[70257]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:193
				// _ = "end of CoverTab[70257]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:193
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:193
			// _ = "end of CoverTab[70254]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:193
			_go_fuzz_dep_.CoverTab[70255]++
														b[n] = b[n-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:194
			// _ = "end of CoverTab[70255]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:195
		// _ = "end of CoverTab[70253]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:196
		_go_fuzz_dep_.CoverTab[70258]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:196
		// _ = "end of CoverTab[70258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:196
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:196
	// _ = "end of CoverTab[70251]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:196
	_go_fuzz_dep_.CoverTab[70252]++
												rb.nrune += 1
												pos := uint8(rb.nbyte)
												rb.nbyte += utf8.UTFMax
												info.pos = pos
												b[n] = info
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:201
	// _ = "end of CoverTab[70252]"
}

// insertErr is an error code returned by insert. Using this type instead
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:204
// of error improves performance up to 20% for many of the benchmarks.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:206
type insertErr int

const (
	iSuccess	insertErr	= -iota
	iShortDst
	iShortSrc
)

// insertFlush inserts the given rune in the buffer ordered by CCC.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:214
// If a decomposition with multiple segments are encountered, they leading
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:214
// ones are flushed.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:214
// It returns a non-zero error code if the rune was not inserted.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:218
func (rb *reorderBuffer) insertFlush(src input, i int, info Properties) insertErr {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:218
	_go_fuzz_dep_.CoverTab[70259]++
												if rune := src.hangul(i); rune != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:219
		_go_fuzz_dep_.CoverTab[70262]++
													rb.decomposeHangul(rune)
													return iSuccess
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:221
		// _ = "end of CoverTab[70262]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:222
		_go_fuzz_dep_.CoverTab[70263]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:222
		// _ = "end of CoverTab[70263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:222
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:222
	// _ = "end of CoverTab[70259]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:222
	_go_fuzz_dep_.CoverTab[70260]++
												if info.hasDecomposition() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:223
		_go_fuzz_dep_.CoverTab[70264]++
													return rb.insertDecomposed(info.Decomposition())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:224
		// _ = "end of CoverTab[70264]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:225
		_go_fuzz_dep_.CoverTab[70265]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:225
		// _ = "end of CoverTab[70265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:225
	// _ = "end of CoverTab[70260]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:225
	_go_fuzz_dep_.CoverTab[70261]++
												rb.insertSingle(src, i, info)
												return iSuccess
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:227
	// _ = "end of CoverTab[70261]"
}

// insertUnsafe inserts the given rune in the buffer ordered by CCC.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:230
// It is assumed there is sufficient space to hold the runes. It is the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:230
// responsibility of the caller to ensure this. This can be done by checking
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:230
// the state returned by the streamSafe type.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:234
func (rb *reorderBuffer) insertUnsafe(src input, i int, info Properties) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:234
	_go_fuzz_dep_.CoverTab[70266]++
												if rune := src.hangul(i); rune != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:235
		_go_fuzz_dep_.CoverTab[70268]++
													rb.decomposeHangul(rune)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:236
		// _ = "end of CoverTab[70268]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:237
		_go_fuzz_dep_.CoverTab[70269]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:237
		// _ = "end of CoverTab[70269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:237
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:237
	// _ = "end of CoverTab[70266]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:237
	_go_fuzz_dep_.CoverTab[70267]++
												if info.hasDecomposition() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:238
		_go_fuzz_dep_.CoverTab[70270]++

													rb.insertDecomposed(info.Decomposition())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:240
		// _ = "end of CoverTab[70270]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:241
		_go_fuzz_dep_.CoverTab[70271]++
													rb.insertSingle(src, i, info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:242
		// _ = "end of CoverTab[70271]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:243
	// _ = "end of CoverTab[70267]"
}

// insertDecomposed inserts an entry in to the reorderBuffer for each rune
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:246
// in dcomp. dcomp must be a sequence of decomposed UTF-8-encoded runes.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:246
// It flushes the buffer on each new segment start.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:249
func (rb *reorderBuffer) insertDecomposed(dcomp []byte) insertErr {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:249
	_go_fuzz_dep_.CoverTab[70272]++
												rb.tmpBytes.setBytes(dcomp)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:254
	for i := 0; i < len(dcomp); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:254
		_go_fuzz_dep_.CoverTab[70274]++
													info := rb.f.info(rb.tmpBytes, i)
													if info.BoundaryBefore() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[70276]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			return rb.nrune > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			// _ = "end of CoverTab[70276]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[70277]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			return !rb.doFlush()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			// _ = "end of CoverTab[70277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:256
			_go_fuzz_dep_.CoverTab[70278]++
														return iShortDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:257
			// _ = "end of CoverTab[70278]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:258
			_go_fuzz_dep_.CoverTab[70279]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:258
			// _ = "end of CoverTab[70279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:258
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:258
		// _ = "end of CoverTab[70274]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:258
		_go_fuzz_dep_.CoverTab[70275]++
													i += copy(rb.byte[rb.nbyte:], dcomp[i:i+int(info.size)])
													rb.insertOrdered(info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:260
		// _ = "end of CoverTab[70275]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:261
	// _ = "end of CoverTab[70272]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:261
	_go_fuzz_dep_.CoverTab[70273]++
												return iSuccess
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:262
	// _ = "end of CoverTab[70273]"
}

// insertSingle inserts an entry in the reorderBuffer for the rune at
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:265
// position i. info is the runeInfo for the rune at position i.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:267
func (rb *reorderBuffer) insertSingle(src input, i int, info Properties) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:267
	_go_fuzz_dep_.CoverTab[70280]++
												src.copySlice(rb.byte[rb.nbyte:], i, i+int(info.size))
												rb.insertOrdered(info)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:269
	// _ = "end of CoverTab[70280]"
}

// insertCGJ inserts a Combining Grapheme Joiner (0x034f) into rb.
func (rb *reorderBuffer) insertCGJ() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:273
	_go_fuzz_dep_.CoverTab[70281]++
												rb.insertSingle(input{str: GraphemeJoiner}, 0, Properties{size: uint8(len(GraphemeJoiner))})
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:274
	// _ = "end of CoverTab[70281]"
}

// appendRune inserts a rune at the end of the buffer. It is used for Hangul.
func (rb *reorderBuffer) appendRune(r rune) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:278
	_go_fuzz_dep_.CoverTab[70282]++
												bn := rb.nbyte
												sz := utf8.EncodeRune(rb.byte[bn:], rune(r))
												rb.nbyte += utf8.UTFMax
												rb.rune[rb.nrune] = Properties{pos: bn, size: uint8(sz)}
												rb.nrune++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:283
	// _ = "end of CoverTab[70282]"
}

// assignRune sets a rune at position pos. It is used for Hangul and recomposition.
func (rb *reorderBuffer) assignRune(pos int, r rune) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:287
	_go_fuzz_dep_.CoverTab[70283]++
												bn := rb.rune[pos].pos
												sz := utf8.EncodeRune(rb.byte[bn:], rune(r))
												rb.rune[pos] = Properties{pos: bn, size: uint8(sz)}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:290
	// _ = "end of CoverTab[70283]"
}

// runeAt returns the rune at position n. It is used for Hangul and recomposition.
func (rb *reorderBuffer) runeAt(n int) rune {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:294
	_go_fuzz_dep_.CoverTab[70284]++
												inf := rb.rune[n]
												r, _ := utf8.DecodeRune(rb.byte[inf.pos : inf.pos+inf.size])
												return r
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:297
	// _ = "end of CoverTab[70284]"
}

// bytesAt returns the UTF-8 encoding of the rune at position n.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:300
// It is used for Hangul and recomposition.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:302
func (rb *reorderBuffer) bytesAt(n int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:302
	_go_fuzz_dep_.CoverTab[70285]++
												inf := rb.rune[n]
												return rb.byte[inf.pos : int(inf.pos)+int(inf.size)]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:304
	// _ = "end of CoverTab[70285]"
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:336
	_go_fuzz_dep_.CoverTab[70286]++
												if len(b) < hangulUTF8Size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:337
		_go_fuzz_dep_.CoverTab[70290]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:338
		// _ = "end of CoverTab[70290]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:339
		_go_fuzz_dep_.CoverTab[70291]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:339
		// _ = "end of CoverTab[70291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:339
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:339
	// _ = "end of CoverTab[70286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:339
	_go_fuzz_dep_.CoverTab[70287]++
												b0 := b[0]
												if b0 < hangulBase0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:341
		_go_fuzz_dep_.CoverTab[70292]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:342
		// _ = "end of CoverTab[70292]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:343
		_go_fuzz_dep_.CoverTab[70293]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:343
		// _ = "end of CoverTab[70293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:343
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:343
	// _ = "end of CoverTab[70287]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:343
	_go_fuzz_dep_.CoverTab[70288]++
												b1 := b[1]
												switch {
	case b0 == hangulBase0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:346
		_go_fuzz_dep_.CoverTab[70294]++
													return b1 >= hangulBase1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:347
		// _ = "end of CoverTab[70294]"
	case b0 < hangulEnd0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:348
		_go_fuzz_dep_.CoverTab[70295]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:349
		// _ = "end of CoverTab[70295]"
	case b0 > hangulEnd0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:350
		_go_fuzz_dep_.CoverTab[70296]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:351
		// _ = "end of CoverTab[70296]"
	case b1 < hangulEnd1:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:352
		_go_fuzz_dep_.CoverTab[70297]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:353
		// _ = "end of CoverTab[70297]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:353
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:353
		_go_fuzz_dep_.CoverTab[70298]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:353
		// _ = "end of CoverTab[70298]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:354
	// _ = "end of CoverTab[70288]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:354
	_go_fuzz_dep_.CoverTab[70289]++
												return b1 == hangulEnd1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:355
		_go_fuzz_dep_.CoverTab[70299]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:355
		return b[2] < hangulEnd2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:355
		// _ = "end of CoverTab[70299]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:355
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:355
	// _ = "end of CoverTab[70289]"
}

func isHangulString(b string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:358
	_go_fuzz_dep_.CoverTab[70300]++
												if len(b) < hangulUTF8Size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:359
		_go_fuzz_dep_.CoverTab[70304]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:360
		// _ = "end of CoverTab[70304]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:361
		_go_fuzz_dep_.CoverTab[70305]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:361
		// _ = "end of CoverTab[70305]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:361
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:361
	// _ = "end of CoverTab[70300]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:361
	_go_fuzz_dep_.CoverTab[70301]++
												b0 := b[0]
												if b0 < hangulBase0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:363
		_go_fuzz_dep_.CoverTab[70306]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:364
		// _ = "end of CoverTab[70306]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:365
		_go_fuzz_dep_.CoverTab[70307]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:365
		// _ = "end of CoverTab[70307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:365
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:365
	// _ = "end of CoverTab[70301]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:365
	_go_fuzz_dep_.CoverTab[70302]++
												b1 := b[1]
												switch {
	case b0 == hangulBase0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:368
		_go_fuzz_dep_.CoverTab[70308]++
													return b1 >= hangulBase1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:369
		// _ = "end of CoverTab[70308]"
	case b0 < hangulEnd0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:370
		_go_fuzz_dep_.CoverTab[70309]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:371
		// _ = "end of CoverTab[70309]"
	case b0 > hangulEnd0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:372
		_go_fuzz_dep_.CoverTab[70310]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:373
		// _ = "end of CoverTab[70310]"
	case b1 < hangulEnd1:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:374
		_go_fuzz_dep_.CoverTab[70311]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:375
		// _ = "end of CoverTab[70311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:375
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:375
		_go_fuzz_dep_.CoverTab[70312]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:375
		// _ = "end of CoverTab[70312]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:376
	// _ = "end of CoverTab[70302]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:376
	_go_fuzz_dep_.CoverTab[70303]++
												return b1 == hangulEnd1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:377
		_go_fuzz_dep_.CoverTab[70313]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:377
		return b[2] < hangulEnd2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:377
		// _ = "end of CoverTab[70313]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:377
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:377
	// _ = "end of CoverTab[70303]"
}

// Caller must ensure len(b) >= 2.
func isJamoVT(b []byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:381
	_go_fuzz_dep_.CoverTab[70314]++

												return b[0] == jamoLBase0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:383
		_go_fuzz_dep_.CoverTab[70315]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:383
		return (b[1] & 0xFC) == jamoLBase1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:383
		// _ = "end of CoverTab[70315]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:383
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:383
	// _ = "end of CoverTab[70314]"
}

func isHangulWithoutJamoT(b []byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:386
	_go_fuzz_dep_.CoverTab[70316]++
												c, _ := utf8.DecodeRune(b)
												c -= hangulBase
												return c < jamoLVTCount && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:389
		_go_fuzz_dep_.CoverTab[70317]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:389
		return c%jamoTCount == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:389
		// _ = "end of CoverTab[70317]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:389
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:389
	// _ = "end of CoverTab[70316]"
}

// decomposeHangul writes the decomposed Hangul to buf and returns the number
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:392
// of bytes written.  len(buf) should be at least 9.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:394
func decomposeHangul(buf []byte, r rune) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:394
	_go_fuzz_dep_.CoverTab[70318]++
												const JamoUTF8Len = 3
												r -= hangulBase
												x := r % jamoTCount
												r /= jamoTCount
												utf8.EncodeRune(buf, jamoLBase+r/jamoVCount)
												utf8.EncodeRune(buf[JamoUTF8Len:], jamoVBase+r%jamoVCount)
												if x != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:401
		_go_fuzz_dep_.CoverTab[70320]++
													utf8.EncodeRune(buf[2*JamoUTF8Len:], jamoTBase+x)
													return 3 * JamoUTF8Len
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:403
		// _ = "end of CoverTab[70320]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:404
		_go_fuzz_dep_.CoverTab[70321]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:404
		// _ = "end of CoverTab[70321]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:404
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:404
	// _ = "end of CoverTab[70318]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:404
	_go_fuzz_dep_.CoverTab[70319]++
												return 2 * JamoUTF8Len
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:405
	// _ = "end of CoverTab[70319]"
}

// decomposeHangul algorithmically decomposes a Hangul rune into
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:408
// its Jamo components.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:408
// See https://unicode.org/reports/tr15/#Hangul for details on decomposing Hangul.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:411
func (rb *reorderBuffer) decomposeHangul(r rune) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:411
	_go_fuzz_dep_.CoverTab[70322]++
												r -= hangulBase
												x := r % jamoTCount
												r /= jamoTCount
												rb.appendRune(jamoLBase + r/jamoVCount)
												rb.appendRune(jamoVBase + r%jamoVCount)
												if x != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:417
		_go_fuzz_dep_.CoverTab[70323]++
													rb.appendRune(jamoTBase + x)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:418
		// _ = "end of CoverTab[70323]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:419
		_go_fuzz_dep_.CoverTab[70324]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:419
		// _ = "end of CoverTab[70324]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:419
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:419
	// _ = "end of CoverTab[70322]"
}

// combineHangul algorithmically combines Jamo character components into Hangul.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:422
// See https://unicode.org/reports/tr15/#Hangul for details on combining Hangul.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:424
func (rb *reorderBuffer) combineHangul(s, i, k int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:424
	_go_fuzz_dep_.CoverTab[70325]++
												b := rb.rune[:]
												bn := rb.nrune
												for ; i < bn; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:427
		_go_fuzz_dep_.CoverTab[70327]++
													cccB := b[k-1].ccc
													cccC := b[i].ccc
													if cccB == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:430
			_go_fuzz_dep_.CoverTab[70329]++
														s = k - 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:431
			// _ = "end of CoverTab[70329]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:432
			_go_fuzz_dep_.CoverTab[70330]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:432
			// _ = "end of CoverTab[70330]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:432
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:432
		// _ = "end of CoverTab[70327]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:432
		_go_fuzz_dep_.CoverTab[70328]++
													if s != k-1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:433
			_go_fuzz_dep_.CoverTab[70331]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:433
			return cccB >= cccC
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:433
			// _ = "end of CoverTab[70331]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:433
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:433
			_go_fuzz_dep_.CoverTab[70332]++

														b[k] = b[i]
														k++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:436
			// _ = "end of CoverTab[70332]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:437
			_go_fuzz_dep_.CoverTab[70333]++
														l := rb.runeAt(s)
														v := rb.runeAt(i)
														switch {
			case jamoLBase <= l && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
				_go_fuzz_dep_.CoverTab[70337]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
				return l < jamoLEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
				// _ = "end of CoverTab[70337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
				_go_fuzz_dep_.CoverTab[70338]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:441
				return jamoVBase <= v
															// _ = "end of CoverTab[70338]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
				_go_fuzz_dep_.CoverTab[70339]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
				return v < jamoVEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
				// _ = "end of CoverTab[70339]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
			}():
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:442
				_go_fuzz_dep_.CoverTab[70334]++

															rb.assignRune(s, hangulBase+
					(l-jamoLBase)*jamoVTCount+(v-jamoVBase)*jamoTCount)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:445
				// _ = "end of CoverTab[70334]"
			case hangulBase <= l && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
				_go_fuzz_dep_.CoverTab[70340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
				return l < hangulEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
				// _ = "end of CoverTab[70340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
				_go_fuzz_dep_.CoverTab[70341]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:446
				return jamoTBase < v
															// _ = "end of CoverTab[70341]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
				_go_fuzz_dep_.CoverTab[70342]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
				return v < jamoTEnd
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
				// _ = "end of CoverTab[70342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
				_go_fuzz_dep_.CoverTab[70343]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:447
				return ((l - hangulBase) % jamoTCount) == 0
															// _ = "end of CoverTab[70343]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:448
			}():
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:448
				_go_fuzz_dep_.CoverTab[70335]++

															rb.assignRune(s, l+v-jamoTBase)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:450
				// _ = "end of CoverTab[70335]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:451
				_go_fuzz_dep_.CoverTab[70336]++
															b[k] = b[i]
															k++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:453
				// _ = "end of CoverTab[70336]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:454
			// _ = "end of CoverTab[70333]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:455
		// _ = "end of CoverTab[70328]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:456
	// _ = "end of CoverTab[70325]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:456
	_go_fuzz_dep_.CoverTab[70326]++
												rb.nrune = k
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:457
	// _ = "end of CoverTab[70326]"
}

// compose recombines the runes in the buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:460
// It should only be used to recompose a single segment, as it will not
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:460
// handle alternations between Hangul and non-Hangul characters correctly.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:463
func (rb *reorderBuffer) compose() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:463
	_go_fuzz_dep_.CoverTab[70344]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:466
	recompMapOnce.Do(buildRecompMap)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:473
	bn := rb.nrune
	if bn == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:474
		_go_fuzz_dep_.CoverTab[70347]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:475
		// _ = "end of CoverTab[70347]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:476
		_go_fuzz_dep_.CoverTab[70348]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:476
		// _ = "end of CoverTab[70348]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:476
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:476
	// _ = "end of CoverTab[70344]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:476
	_go_fuzz_dep_.CoverTab[70345]++
												k := 1
												b := rb.rune[:]
												for s, i := 0, 1; i < bn; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:479
		_go_fuzz_dep_.CoverTab[70349]++
													if isJamoVT(rb.bytesAt(i)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:480
			_go_fuzz_dep_.CoverTab[70352]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:483
			rb.combineHangul(s, i, k)
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:484
			// _ = "end of CoverTab[70352]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:485
			_go_fuzz_dep_.CoverTab[70353]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:485
			// _ = "end of CoverTab[70353]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:485
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:485
		// _ = "end of CoverTab[70349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:485
		_go_fuzz_dep_.CoverTab[70350]++
													ii := b[i]

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:491
		if ii.combinesBackward() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:491
			_go_fuzz_dep_.CoverTab[70354]++
														cccB := b[k-1].ccc
														cccC := ii.ccc
														blocked := false
														if cccB == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:495
				_go_fuzz_dep_.CoverTab[70356]++
															s = k - 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:496
				// _ = "end of CoverTab[70356]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:497
				_go_fuzz_dep_.CoverTab[70357]++
															blocked = s != k-1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:498
					_go_fuzz_dep_.CoverTab[70358]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:498
					return cccB >= cccC
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:498
					// _ = "end of CoverTab[70358]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:498
				}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:498
				// _ = "end of CoverTab[70357]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:499
			// _ = "end of CoverTab[70354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:499
			_go_fuzz_dep_.CoverTab[70355]++
														if !blocked {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:500
				_go_fuzz_dep_.CoverTab[70359]++
															combined := combine(rb.runeAt(s), rb.runeAt(i))
															if combined != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:502
					_go_fuzz_dep_.CoverTab[70360]++
																rb.assignRune(s, combined)
																continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:504
					// _ = "end of CoverTab[70360]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:505
					_go_fuzz_dep_.CoverTab[70361]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:505
					// _ = "end of CoverTab[70361]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:505
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:505
				// _ = "end of CoverTab[70359]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:506
				_go_fuzz_dep_.CoverTab[70362]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:506
				// _ = "end of CoverTab[70362]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:506
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:506
			// _ = "end of CoverTab[70355]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:507
			_go_fuzz_dep_.CoverTab[70363]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:507
			// _ = "end of CoverTab[70363]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:507
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:507
		// _ = "end of CoverTab[70350]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:507
		_go_fuzz_dep_.CoverTab[70351]++
													b[k] = b[i]
													k++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:509
		// _ = "end of CoverTab[70351]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:510
	// _ = "end of CoverTab[70345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:510
	_go_fuzz_dep_.CoverTab[70346]++
												rb.nrune = k
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:511
	// _ = "end of CoverTab[70346]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:512
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/composition.go:512
var _ = _go_fuzz_dep_.CoverTab
