// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:5
)

import (
	"github.com/klauspost/compress/huff0"
)

// history contains the information transferred between blocks.
type history struct {
	b		[]byte
	huffTree	*huff0.Scratch
	recentOffsets	[3]int
	decoders	sequenceDecs
	windowSize	int
	maxSize		int
	error		bool
	dict		*dict
}

// reset will reset the history to initial state of a frame.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:23
// The history must already have been initialized to the desired size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:25
func (h *history) reset() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:25
	_go_fuzz_dep_.CoverTab[94841]++
												h.b = h.b[:0]
												h.error = false
												h.recentOffsets = [3]int{1, 4, 8}
												if f := h.decoders.litLengths.fse; f != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:29
		_go_fuzz_dep_.CoverTab[94846]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:29
		return !f.preDefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:29
		// _ = "end of CoverTab[94846]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:29
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:29
		_go_fuzz_dep_.CoverTab[94847]++
													fseDecoderPool.Put(f)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:30
		// _ = "end of CoverTab[94847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:31
		_go_fuzz_dep_.CoverTab[94848]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:31
		// _ = "end of CoverTab[94848]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:31
	// _ = "end of CoverTab[94841]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:31
	_go_fuzz_dep_.CoverTab[94842]++
												if f := h.decoders.offsets.fse; f != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:32
		_go_fuzz_dep_.CoverTab[94849]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:32
		return !f.preDefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:32
		// _ = "end of CoverTab[94849]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:32
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:32
		_go_fuzz_dep_.CoverTab[94850]++
													fseDecoderPool.Put(f)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:33
		// _ = "end of CoverTab[94850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:34
		_go_fuzz_dep_.CoverTab[94851]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:34
		// _ = "end of CoverTab[94851]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:34
	// _ = "end of CoverTab[94842]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:34
	_go_fuzz_dep_.CoverTab[94843]++
												if f := h.decoders.matchLengths.fse; f != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:35
		_go_fuzz_dep_.CoverTab[94852]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:35
		return !f.preDefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:35
		// _ = "end of CoverTab[94852]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:35
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:35
		_go_fuzz_dep_.CoverTab[94853]++
													fseDecoderPool.Put(f)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:36
		// _ = "end of CoverTab[94853]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:37
		_go_fuzz_dep_.CoverTab[94854]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:37
		// _ = "end of CoverTab[94854]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:37
	// _ = "end of CoverTab[94843]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:37
	_go_fuzz_dep_.CoverTab[94844]++
												h.decoders = sequenceDecs{}
												if h.huffTree != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:39
		_go_fuzz_dep_.CoverTab[94855]++
													if h.dict == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:40
			_go_fuzz_dep_.CoverTab[94856]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:40
			return h.dict.litEnc != h.huffTree
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:40
			// _ = "end of CoverTab[94856]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:40
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:40
			_go_fuzz_dep_.CoverTab[94857]++
														huffDecoderPool.Put(h.huffTree)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:41
			// _ = "end of CoverTab[94857]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:42
			_go_fuzz_dep_.CoverTab[94858]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:42
			// _ = "end of CoverTab[94858]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:42
		// _ = "end of CoverTab[94855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:43
		_go_fuzz_dep_.CoverTab[94859]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:43
		// _ = "end of CoverTab[94859]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:43
	// _ = "end of CoverTab[94844]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:43
	_go_fuzz_dep_.CoverTab[94845]++
												h.huffTree = nil
												h.dict = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:45
	// _ = "end of CoverTab[94845]"

}

func (h *history) setDict(dict *dict) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:49
	_go_fuzz_dep_.CoverTab[94860]++
												if dict == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:50
		_go_fuzz_dep_.CoverTab[94862]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:51
		// _ = "end of CoverTab[94862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:52
		_go_fuzz_dep_.CoverTab[94863]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:52
		// _ = "end of CoverTab[94863]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:52
	// _ = "end of CoverTab[94860]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:52
	_go_fuzz_dep_.CoverTab[94861]++
												h.dict = dict
												h.decoders.litLengths = dict.llDec
												h.decoders.offsets = dict.ofDec
												h.decoders.matchLengths = dict.mlDec
												h.recentOffsets = dict.offsets
												h.huffTree = dict.litEnc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:58
	// _ = "end of CoverTab[94861]"
}

// append bytes to history.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:61
// This function will make sure there is space for it,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:61
// if the buffer has been allocated with enough extra space.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:64
func (h *history) append(b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:64
	_go_fuzz_dep_.CoverTab[94864]++
												if len(b) >= h.windowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:65
		_go_fuzz_dep_.CoverTab[94867]++

													h.b = h.b[:h.windowSize]
													copy(h.b, b[len(b)-h.windowSize:])
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:69
		// _ = "end of CoverTab[94867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:70
		_go_fuzz_dep_.CoverTab[94868]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:70
		// _ = "end of CoverTab[94868]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:70
	// _ = "end of CoverTab[94864]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:70
	_go_fuzz_dep_.CoverTab[94865]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:73
	if len(b) < cap(h.b)-len(h.b) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:73
		_go_fuzz_dep_.CoverTab[94869]++
													h.b = append(h.b, b...)
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:75
		// _ = "end of CoverTab[94869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:76
		_go_fuzz_dep_.CoverTab[94870]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:76
		// _ = "end of CoverTab[94870]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:76
	// _ = "end of CoverTab[94865]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:76
	_go_fuzz_dep_.CoverTab[94866]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:80
	discard := len(b) + len(h.b) - h.windowSize
												copy(h.b, h.b[discard:])
												h.b = h.b[:h.windowSize]
												copy(h.b[h.windowSize-len(b):], b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:83
	// _ = "end of CoverTab[94866]"
}

// append bytes to history without ever discarding anything.
func (h *history) appendKeep(b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:87
	_go_fuzz_dep_.CoverTab[94871]++
												h.b = append(h.b, b...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:88
	// _ = "end of CoverTab[94871]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/history.go:89
var _ = _go_fuzz_dep_.CoverTab
