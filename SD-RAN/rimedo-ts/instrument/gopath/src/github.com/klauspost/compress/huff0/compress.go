//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:1
)

import (
	"fmt"
	"runtime"
	"sync"
)

// Compress1X will compress the input.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:9
// The output can be decoded using Decompress1X.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:9
// Supply a Scratch object. The scratch object contains state about re-use,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:9
// So when sharing across independent encodes, be sure to set the re-use policy.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:13
func Compress1X(in []byte, s *Scratch) (out []byte, reUsed bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:13
	_go_fuzz_dep_.CoverTab[89535]++
												s, err = s.prepare(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:15
		_go_fuzz_dep_.CoverTab[89537]++
													return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:16
		// _ = "end of CoverTab[89537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:17
		_go_fuzz_dep_.CoverTab[89538]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:17
		// _ = "end of CoverTab[89538]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:17
	// _ = "end of CoverTab[89535]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:17
	_go_fuzz_dep_.CoverTab[89536]++
												return compress(in, s, s.compress1X)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:18
	// _ = "end of CoverTab[89536]"
}

// Compress4X will compress the input. The input is split into 4 independent blocks
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:21
// and compressed similar to Compress1X.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:21
// The output can be decoded using Decompress4X.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:21
// Supply a Scratch object. The scratch object contains state about re-use,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:21
// So when sharing across independent encodes, be sure to set the re-use policy.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:26
func Compress4X(in []byte, s *Scratch) (out []byte, reUsed bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:26
	_go_fuzz_dep_.CoverTab[89539]++
												s, err = s.prepare(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:28
		_go_fuzz_dep_.CoverTab[89542]++
													return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:29
		// _ = "end of CoverTab[89542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:30
		_go_fuzz_dep_.CoverTab[89543]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:30
		// _ = "end of CoverTab[89543]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:30
	// _ = "end of CoverTab[89539]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:30
	_go_fuzz_dep_.CoverTab[89540]++
												if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:31
		_go_fuzz_dep_.CoverTab[89544]++
		// TODO: compress4Xp only slightly faster.
		const parallelThreshold = 8 << 10
		if len(in) < parallelThreshold || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:34
			_go_fuzz_dep_.CoverTab[89546]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:34
			return runtime.GOMAXPROCS(0) == 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:34
			// _ = "end of CoverTab[89546]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:34
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:34
			_go_fuzz_dep_.CoverTab[89547]++
														return compress(in, s, s.compress4X)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:35
			// _ = "end of CoverTab[89547]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:36
			_go_fuzz_dep_.CoverTab[89548]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:36
			// _ = "end of CoverTab[89548]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:36
		// _ = "end of CoverTab[89544]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:36
		_go_fuzz_dep_.CoverTab[89545]++
													return compress(in, s, s.compress4Xp)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:37
		// _ = "end of CoverTab[89545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:38
		_go_fuzz_dep_.CoverTab[89549]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:38
		// _ = "end of CoverTab[89549]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:38
	// _ = "end of CoverTab[89540]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:38
	_go_fuzz_dep_.CoverTab[89541]++
												return compress(in, s, s.compress4X)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:39
	// _ = "end of CoverTab[89541]"
}

func compress(in []byte, s *Scratch, compressor func(src []byte) ([]byte, error)) (out []byte, reUsed bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:42
	_go_fuzz_dep_.CoverTab[89550]++

												if s.Reuse == ReusePolicyNone {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:44
		_go_fuzz_dep_.CoverTab[89564]++
													s.prevTable = s.prevTable[:0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:45
		// _ = "end of CoverTab[89564]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:46
		_go_fuzz_dep_.CoverTab[89565]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:46
		// _ = "end of CoverTab[89565]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:46
	// _ = "end of CoverTab[89550]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:46
	_go_fuzz_dep_.CoverTab[89551]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:49
	maxCount := s.maxCount
	var canReuse = false
	if maxCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:51
		_go_fuzz_dep_.CoverTab[89566]++
													maxCount, canReuse = s.countSimple(in)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:52
		// _ = "end of CoverTab[89566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:53
		_go_fuzz_dep_.CoverTab[89567]++
													canReuse = s.canUseTable(s.prevTable)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:54
		// _ = "end of CoverTab[89567]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:55
	// _ = "end of CoverTab[89551]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:55
	_go_fuzz_dep_.CoverTab[89552]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:58
	wantSize := len(in)
	if s.WantLogLess > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:59
		_go_fuzz_dep_.CoverTab[89568]++
													wantSize -= wantSize >> s.WantLogLess
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:60
		// _ = "end of CoverTab[89568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:61
		_go_fuzz_dep_.CoverTab[89569]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:61
		// _ = "end of CoverTab[89569]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:61
	// _ = "end of CoverTab[89552]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:61
	_go_fuzz_dep_.CoverTab[89553]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:64
	s.clearCount = true
	s.maxCount = 0
	if maxCount >= len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:66
		_go_fuzz_dep_.CoverTab[89570]++
													if maxCount > len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:67
			_go_fuzz_dep_.CoverTab[89573]++
														return nil, false, fmt.Errorf("maxCount (%d) > length (%d)", maxCount, len(in))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:68
			// _ = "end of CoverTab[89573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:69
			_go_fuzz_dep_.CoverTab[89574]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:69
			// _ = "end of CoverTab[89574]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:69
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:69
		// _ = "end of CoverTab[89570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:69
		_go_fuzz_dep_.CoverTab[89571]++
													if len(in) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:70
			_go_fuzz_dep_.CoverTab[89575]++
														return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:71
			// _ = "end of CoverTab[89575]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:72
			_go_fuzz_dep_.CoverTab[89576]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:72
			// _ = "end of CoverTab[89576]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:72
		// _ = "end of CoverTab[89571]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:72
		_go_fuzz_dep_.CoverTab[89572]++

													return nil, false, ErrUseRLE
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:74
		// _ = "end of CoverTab[89572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:75
		_go_fuzz_dep_.CoverTab[89577]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:75
		// _ = "end of CoverTab[89577]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:75
	// _ = "end of CoverTab[89553]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:75
	_go_fuzz_dep_.CoverTab[89554]++
												if maxCount == 1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:76
		_go_fuzz_dep_.CoverTab[89578]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:76
		return maxCount < (len(in) >> 7)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:76
		// _ = "end of CoverTab[89578]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:76
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:76
		_go_fuzz_dep_.CoverTab[89579]++

													return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:78
		// _ = "end of CoverTab[89579]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:79
		_go_fuzz_dep_.CoverTab[89580]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:79
		// _ = "end of CoverTab[89580]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:79
	// _ = "end of CoverTab[89554]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:79
	_go_fuzz_dep_.CoverTab[89555]++
												if s.Reuse == ReusePolicyMust && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:80
		_go_fuzz_dep_.CoverTab[89581]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:80
		return !canReuse
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:80
		// _ = "end of CoverTab[89581]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:80
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:80
		_go_fuzz_dep_.CoverTab[89582]++

													return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:82
		// _ = "end of CoverTab[89582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:83
		_go_fuzz_dep_.CoverTab[89583]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:83
		// _ = "end of CoverTab[89583]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:83
	// _ = "end of CoverTab[89555]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:83
	_go_fuzz_dep_.CoverTab[89556]++
												if (s.Reuse == ReusePolicyPrefer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		_go_fuzz_dep_.CoverTab[89584]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		return s.Reuse == ReusePolicyMust
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		// _ = "end of CoverTab[89584]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		_go_fuzz_dep_.CoverTab[89585]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		return canReuse
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		// _ = "end of CoverTab[89585]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:84
		_go_fuzz_dep_.CoverTab[89586]++
													keepTable := s.cTable
													keepTL := s.actualTableLog
													s.cTable = s.prevTable
													s.actualTableLog = s.prevTableLog
													s.Out, err = compressor(in)
													s.cTable = keepTable
													s.actualTableLog = keepTL
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:92
			_go_fuzz_dep_.CoverTab[89589]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:92
			return len(s.Out) < wantSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:92
			// _ = "end of CoverTab[89589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:92
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:92
			_go_fuzz_dep_.CoverTab[89590]++
														s.OutData = s.Out
														return s.Out, true, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:94
			// _ = "end of CoverTab[89590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:95
			_go_fuzz_dep_.CoverTab[89591]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:95
			// _ = "end of CoverTab[89591]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:95
		// _ = "end of CoverTab[89586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:95
		_go_fuzz_dep_.CoverTab[89587]++
													if s.Reuse == ReusePolicyMust {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:96
			_go_fuzz_dep_.CoverTab[89592]++
														return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:97
			// _ = "end of CoverTab[89592]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:98
			_go_fuzz_dep_.CoverTab[89593]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:98
			// _ = "end of CoverTab[89593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:98
		// _ = "end of CoverTab[89587]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:98
		_go_fuzz_dep_.CoverTab[89588]++

													s.prevTable = s.prevTable[:0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:100
		// _ = "end of CoverTab[89588]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:101
		_go_fuzz_dep_.CoverTab[89594]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:101
		// _ = "end of CoverTab[89594]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:101
	// _ = "end of CoverTab[89556]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:101
	_go_fuzz_dep_.CoverTab[89557]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:104
	err = s.buildCTable()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:105
		_go_fuzz_dep_.CoverTab[89595]++
													return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:106
		// _ = "end of CoverTab[89595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:107
		_go_fuzz_dep_.CoverTab[89596]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:107
		// _ = "end of CoverTab[89596]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:107
	// _ = "end of CoverTab[89557]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:107
	_go_fuzz_dep_.CoverTab[89558]++

												if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:109
		_go_fuzz_dep_.CoverTab[89597]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:109
		return !s.canUseTable(s.cTable)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:109
		// _ = "end of CoverTab[89597]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:109
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:109
		_go_fuzz_dep_.CoverTab[89598]++
													panic("invalid table generated")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:110
		// _ = "end of CoverTab[89598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:111
		_go_fuzz_dep_.CoverTab[89599]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:111
		// _ = "end of CoverTab[89599]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:111
	// _ = "end of CoverTab[89558]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:111
	_go_fuzz_dep_.CoverTab[89559]++

												if s.Reuse == ReusePolicyAllow && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:113
		_go_fuzz_dep_.CoverTab[89600]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:113
		return canReuse
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:113
		// _ = "end of CoverTab[89600]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:113
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:113
		_go_fuzz_dep_.CoverTab[89601]++
													hSize := len(s.Out)
													oldSize := s.prevTable.estimateSize(s.count[:s.symbolLen])
													newSize := s.cTable.estimateSize(s.count[:s.symbolLen])
													if oldSize <= hSize+newSize || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:117
			_go_fuzz_dep_.CoverTab[89602]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:117
			return hSize+12 >= wantSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:117
			// _ = "end of CoverTab[89602]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:117
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:117
			_go_fuzz_dep_.CoverTab[89603]++

														keepTable := s.cTable
														keepTL := s.actualTableLog

														s.cTable = s.prevTable
														s.actualTableLog = s.prevTableLog
														s.Out, err = compressor(in)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:127
			s.cTable = keepTable
			s.actualTableLog = keepTL
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:129
				_go_fuzz_dep_.CoverTab[89606]++
															return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:130
				// _ = "end of CoverTab[89606]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:131
				_go_fuzz_dep_.CoverTab[89607]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:131
				// _ = "end of CoverTab[89607]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:131
			// _ = "end of CoverTab[89603]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:131
			_go_fuzz_dep_.CoverTab[89604]++
														if len(s.Out) >= wantSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:132
				_go_fuzz_dep_.CoverTab[89608]++
															return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:133
				// _ = "end of CoverTab[89608]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:134
				_go_fuzz_dep_.CoverTab[89609]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:134
				// _ = "end of CoverTab[89609]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:134
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:134
			// _ = "end of CoverTab[89604]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:134
			_go_fuzz_dep_.CoverTab[89605]++
														s.OutData = s.Out
														return s.Out, true, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:136
			// _ = "end of CoverTab[89605]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:137
			_go_fuzz_dep_.CoverTab[89610]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:137
			// _ = "end of CoverTab[89610]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:137
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:137
		// _ = "end of CoverTab[89601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:138
		_go_fuzz_dep_.CoverTab[89611]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:138
		// _ = "end of CoverTab[89611]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:138
	// _ = "end of CoverTab[89559]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:138
	_go_fuzz_dep_.CoverTab[89560]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:141
	err = s.cTable.write(s)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:142
		_go_fuzz_dep_.CoverTab[89612]++
													s.OutTable = nil
													return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:144
		// _ = "end of CoverTab[89612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:145
		_go_fuzz_dep_.CoverTab[89613]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:145
		// _ = "end of CoverTab[89613]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:145
	// _ = "end of CoverTab[89560]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:145
	_go_fuzz_dep_.CoverTab[89561]++
												s.OutTable = s.Out

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:149
	s.Out, err = compressor(in)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:150
		_go_fuzz_dep_.CoverTab[89614]++
													s.OutTable = nil
													return nil, false, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:152
		// _ = "end of CoverTab[89614]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:153
		_go_fuzz_dep_.CoverTab[89615]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:153
		// _ = "end of CoverTab[89615]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:153
	// _ = "end of CoverTab[89561]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:153
	_go_fuzz_dep_.CoverTab[89562]++
												if len(s.Out) >= wantSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:154
		_go_fuzz_dep_.CoverTab[89616]++
													s.OutTable = nil
													return nil, false, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:156
		// _ = "end of CoverTab[89616]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:157
		_go_fuzz_dep_.CoverTab[89617]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:157
		// _ = "end of CoverTab[89617]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:157
	// _ = "end of CoverTab[89562]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:157
	_go_fuzz_dep_.CoverTab[89563]++

												s.prevTable, s.prevTableLog, s.cTable = s.cTable, s.actualTableLog, s.prevTable[:0]
												s.OutData = s.Out[len(s.OutTable):]
												return s.Out, false, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:161
	// _ = "end of CoverTab[89563]"
}

// EstimateSizes will estimate the data sizes
func EstimateSizes(in []byte, s *Scratch) (tableSz, dataSz, reuseSz int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:165
	_go_fuzz_dep_.CoverTab[89618]++
												s, err = s.prepare(in)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:167
		_go_fuzz_dep_.CoverTab[89628]++
													return 0, 0, 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:168
		// _ = "end of CoverTab[89628]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:169
		_go_fuzz_dep_.CoverTab[89629]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:169
		// _ = "end of CoverTab[89629]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:169
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:169
	// _ = "end of CoverTab[89618]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:169
	_go_fuzz_dep_.CoverTab[89619]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:172
	tableSz, dataSz, reuseSz = -1, -1, -1
	maxCount := s.maxCount
	var canReuse = false
	if maxCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:175
		_go_fuzz_dep_.CoverTab[89630]++
													maxCount, canReuse = s.countSimple(in)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:176
		// _ = "end of CoverTab[89630]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:177
		_go_fuzz_dep_.CoverTab[89631]++
													canReuse = s.canUseTable(s.prevTable)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:178
		// _ = "end of CoverTab[89631]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:179
	// _ = "end of CoverTab[89619]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:179
	_go_fuzz_dep_.CoverTab[89620]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:182
	wantSize := len(in)
	if s.WantLogLess > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:183
		_go_fuzz_dep_.CoverTab[89632]++
													wantSize -= wantSize >> s.WantLogLess
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:184
		// _ = "end of CoverTab[89632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:185
		_go_fuzz_dep_.CoverTab[89633]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:185
		// _ = "end of CoverTab[89633]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:185
	// _ = "end of CoverTab[89620]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:185
	_go_fuzz_dep_.CoverTab[89621]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:188
	s.clearCount = true
	s.maxCount = 0
	if maxCount >= len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:190
		_go_fuzz_dep_.CoverTab[89634]++
													if maxCount > len(in) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:191
			_go_fuzz_dep_.CoverTab[89637]++
														return 0, 0, 0, fmt.Errorf("maxCount (%d) > length (%d)", maxCount, len(in))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:192
			// _ = "end of CoverTab[89637]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:193
			_go_fuzz_dep_.CoverTab[89638]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:193
			// _ = "end of CoverTab[89638]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:193
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:193
		// _ = "end of CoverTab[89634]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:193
		_go_fuzz_dep_.CoverTab[89635]++
													if len(in) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:194
			_go_fuzz_dep_.CoverTab[89639]++
														return 0, 0, 0, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:195
			// _ = "end of CoverTab[89639]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:196
			_go_fuzz_dep_.CoverTab[89640]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:196
			// _ = "end of CoverTab[89640]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:196
		// _ = "end of CoverTab[89635]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:196
		_go_fuzz_dep_.CoverTab[89636]++

													return 0, 0, 0, ErrUseRLE
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:198
		// _ = "end of CoverTab[89636]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:199
		_go_fuzz_dep_.CoverTab[89641]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:199
		// _ = "end of CoverTab[89641]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:199
	// _ = "end of CoverTab[89621]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:199
	_go_fuzz_dep_.CoverTab[89622]++
												if maxCount == 1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:200
		_go_fuzz_dep_.CoverTab[89642]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:200
		return maxCount < (len(in) >> 7)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:200
		// _ = "end of CoverTab[89642]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:200
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:200
		_go_fuzz_dep_.CoverTab[89643]++

													return 0, 0, 0, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:202
		// _ = "end of CoverTab[89643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:203
		_go_fuzz_dep_.CoverTab[89644]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:203
		// _ = "end of CoverTab[89644]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:203
	// _ = "end of CoverTab[89622]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:203
	_go_fuzz_dep_.CoverTab[89623]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:206
	err = s.buildCTable()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:207
		_go_fuzz_dep_.CoverTab[89645]++
													return 0, 0, 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:208
		// _ = "end of CoverTab[89645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:209
		_go_fuzz_dep_.CoverTab[89646]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:209
		// _ = "end of CoverTab[89646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:209
	// _ = "end of CoverTab[89623]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:209
	_go_fuzz_dep_.CoverTab[89624]++

												if false && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:211
		_go_fuzz_dep_.CoverTab[89647]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:211
		return !s.canUseTable(s.cTable)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:211
		// _ = "end of CoverTab[89647]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:211
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:211
		_go_fuzz_dep_.CoverTab[89648]++
													panic("invalid table generated")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:212
		// _ = "end of CoverTab[89648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:213
		_go_fuzz_dep_.CoverTab[89649]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:213
		// _ = "end of CoverTab[89649]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:213
	// _ = "end of CoverTab[89624]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:213
	_go_fuzz_dep_.CoverTab[89625]++

												tableSz, err = s.cTable.estTableSize(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:216
		_go_fuzz_dep_.CoverTab[89650]++
													return 0, 0, 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:217
		// _ = "end of CoverTab[89650]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:218
		_go_fuzz_dep_.CoverTab[89651]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:218
		// _ = "end of CoverTab[89651]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:218
	// _ = "end of CoverTab[89625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:218
	_go_fuzz_dep_.CoverTab[89626]++
												if canReuse {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:219
		_go_fuzz_dep_.CoverTab[89652]++
													reuseSz = s.prevTable.estimateSize(s.count[:s.symbolLen])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:220
		// _ = "end of CoverTab[89652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:221
		_go_fuzz_dep_.CoverTab[89653]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:221
		// _ = "end of CoverTab[89653]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:221
	// _ = "end of CoverTab[89626]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:221
	_go_fuzz_dep_.CoverTab[89627]++
												dataSz = s.cTable.estimateSize(s.count[:s.symbolLen])

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:225
	return tableSz, dataSz, reuseSz, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:225
	// _ = "end of CoverTab[89627]"
}

func (s *Scratch) compress1X(src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:228
	_go_fuzz_dep_.CoverTab[89654]++
												return s.compress1xDo(s.Out, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:229
	// _ = "end of CoverTab[89654]"
}

func (s *Scratch) compress1xDo(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:232
	_go_fuzz_dep_.CoverTab[89655]++
												var bw = bitWriter{out: dst}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:236
	n := len(src)
												n -= n & 3
												cTable := s.cTable[:256]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:241
	for i := len(src) & 3; i > 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:241
		_go_fuzz_dep_.CoverTab[89658]++
													bw.encSymbol(cTable, src[n+i-1])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:242
		// _ = "end of CoverTab[89658]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:243
	// _ = "end of CoverTab[89655]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:243
	_go_fuzz_dep_.CoverTab[89656]++
												n -= 4
												if s.actualTableLog <= 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:245
		_go_fuzz_dep_.CoverTab[89659]++
													for ; n >= 0; n -= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:246
			_go_fuzz_dep_.CoverTab[89660]++
														tmp := src[n : n+4]

														bw.flush32()
														bw.encTwoSymbols(cTable, tmp[3], tmp[2])
														bw.encTwoSymbols(cTable, tmp[1], tmp[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:251
			// _ = "end of CoverTab[89660]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:252
		// _ = "end of CoverTab[89659]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:253
		_go_fuzz_dep_.CoverTab[89661]++
													for ; n >= 0; n -= 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:254
			_go_fuzz_dep_.CoverTab[89662]++
														tmp := src[n : n+4]

														bw.flush32()
														bw.encTwoSymbols(cTable, tmp[3], tmp[2])
														bw.flush32()
														bw.encTwoSymbols(cTable, tmp[1], tmp[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:260
			// _ = "end of CoverTab[89662]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:261
		// _ = "end of CoverTab[89661]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:262
	// _ = "end of CoverTab[89656]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:262
	_go_fuzz_dep_.CoverTab[89657]++
												err := bw.close()
												return bw.out, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:264
	// _ = "end of CoverTab[89657]"
}

var sixZeros [6]byte

func (s *Scratch) compress4X(src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:269
	_go_fuzz_dep_.CoverTab[89663]++
												if len(src) < 12 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:270
		_go_fuzz_dep_.CoverTab[89666]++
													return nil, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:271
		// _ = "end of CoverTab[89666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:272
		_go_fuzz_dep_.CoverTab[89667]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:272
		// _ = "end of CoverTab[89667]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:272
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:272
	// _ = "end of CoverTab[89663]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:272
	_go_fuzz_dep_.CoverTab[89664]++
												segmentSize := (len(src) + 3) / 4

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:276
	offsetIdx := len(s.Out)
	s.Out = append(s.Out, sixZeros[:]...)

	for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:279
		_go_fuzz_dep_.CoverTab[89668]++
													toDo := src
													if len(toDo) > segmentSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:281
			_go_fuzz_dep_.CoverTab[89671]++
														toDo = toDo[:segmentSize]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:282
			// _ = "end of CoverTab[89671]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:283
			_go_fuzz_dep_.CoverTab[89672]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:283
			// _ = "end of CoverTab[89672]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:283
		// _ = "end of CoverTab[89668]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:283
		_go_fuzz_dep_.CoverTab[89669]++
													src = src[len(toDo):]

													var err error
													idx := len(s.Out)
													s.Out, err = s.compress1xDo(s.Out, toDo)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:289
			_go_fuzz_dep_.CoverTab[89673]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:290
			// _ = "end of CoverTab[89673]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:291
			_go_fuzz_dep_.CoverTab[89674]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:291
			// _ = "end of CoverTab[89674]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:291
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:291
		// _ = "end of CoverTab[89669]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:291
		_go_fuzz_dep_.CoverTab[89670]++

													if i < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:293
			_go_fuzz_dep_.CoverTab[89675]++

														length := len(s.Out) - idx
														s.Out[i*2+offsetIdx] = byte(length)
														s.Out[i*2+offsetIdx+1] = byte(length >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:297
			// _ = "end of CoverTab[89675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:298
			_go_fuzz_dep_.CoverTab[89676]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:298
			// _ = "end of CoverTab[89676]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:298
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:298
		// _ = "end of CoverTab[89670]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:299
	// _ = "end of CoverTab[89664]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:299
	_go_fuzz_dep_.CoverTab[89665]++

												return s.Out, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:301
	// _ = "end of CoverTab[89665]"
}

// compress4Xp will compress 4 streams using separate goroutines.
func (s *Scratch) compress4Xp(src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:305
	_go_fuzz_dep_.CoverTab[89677]++
												if len(src) < 12 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:306
		_go_fuzz_dep_.CoverTab[89681]++
													return nil, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:307
		// _ = "end of CoverTab[89681]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:308
		_go_fuzz_dep_.CoverTab[89682]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:308
		// _ = "end of CoverTab[89682]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:308
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:308
	// _ = "end of CoverTab[89677]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:308
	_go_fuzz_dep_.CoverTab[89678]++

												s.Out = s.Out[:6]

												segmentSize := (len(src) + 3) / 4
												var wg sync.WaitGroup
												var errs [4]error
												wg.Add(4)
												for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:316
		_go_fuzz_dep_.CoverTab[89683]++
													toDo := src
													if len(toDo) > segmentSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:318
			_go_fuzz_dep_.CoverTab[89685]++
														toDo = toDo[:segmentSize]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:319
			// _ = "end of CoverTab[89685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:320
			_go_fuzz_dep_.CoverTab[89686]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:320
			// _ = "end of CoverTab[89686]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:320
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:320
		// _ = "end of CoverTab[89683]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:320
		_go_fuzz_dep_.CoverTab[89684]++
													src = src[len(toDo):]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:321
		_curRoutineNum102_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:321
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum102_)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
		go func(i int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
			_go_fuzz_dep_.CoverTab[89687]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
				_go_fuzz_dep_.CoverTab[89688]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum102_)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
				// _ = "end of CoverTab[89688]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:324
			}()
														s.tmpOut[i], errs[i] = s.compress1xDo(s.tmpOut[i][:0], toDo)
														wg.Done()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:326
			// _ = "end of CoverTab[89687]"
		}(i)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:327
		// _ = "end of CoverTab[89684]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:328
	// _ = "end of CoverTab[89678]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:328
	_go_fuzz_dep_.CoverTab[89679]++
												wg.Wait()
												for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:330
		_go_fuzz_dep_.CoverTab[89689]++
													if errs[i] != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:331
			_go_fuzz_dep_.CoverTab[89692]++
														return nil, errs[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:332
			// _ = "end of CoverTab[89692]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:333
			_go_fuzz_dep_.CoverTab[89693]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:333
			// _ = "end of CoverTab[89693]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:333
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:333
		// _ = "end of CoverTab[89689]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:333
		_go_fuzz_dep_.CoverTab[89690]++
													o := s.tmpOut[i]

													if i < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:336
			_go_fuzz_dep_.CoverTab[89694]++

														s.Out[i*2] = byte(len(o))
														s.Out[i*2+1] = byte(len(o) >> 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:339
			// _ = "end of CoverTab[89694]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:340
			_go_fuzz_dep_.CoverTab[89695]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:340
			// _ = "end of CoverTab[89695]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:340
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:340
		// _ = "end of CoverTab[89690]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:340
		_go_fuzz_dep_.CoverTab[89691]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:343
		s.Out = append(s.Out, o...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:343
		// _ = "end of CoverTab[89691]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:344
	// _ = "end of CoverTab[89679]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:344
	_go_fuzz_dep_.CoverTab[89680]++
												return s.Out, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:345
	// _ = "end of CoverTab[89680]"
}

// countSimple will create a simple histogram in s.count.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:348
// Returns the biggest count.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:348
// Does not update s.clearCount.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:351
func (s *Scratch) countSimple(in []byte) (max int, reuse bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:351
	_go_fuzz_dep_.CoverTab[89696]++
												reuse = true
												for _, v := range in {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:353
		_go_fuzz_dep_.CoverTab[89700]++
													s.count[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:354
		// _ = "end of CoverTab[89700]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:355
	// _ = "end of CoverTab[89696]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:355
	_go_fuzz_dep_.CoverTab[89697]++
												m := uint32(0)
												if len(s.prevTable) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:357
		_go_fuzz_dep_.CoverTab[89701]++
													for i, v := range s.count[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:358
			_go_fuzz_dep_.CoverTab[89703]++
														if v > m {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:359
				_go_fuzz_dep_.CoverTab[89705]++
															m = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:360
				// _ = "end of CoverTab[89705]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:361
				_go_fuzz_dep_.CoverTab[89706]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:361
				// _ = "end of CoverTab[89706]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:361
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:361
			// _ = "end of CoverTab[89703]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:361
			_go_fuzz_dep_.CoverTab[89704]++
														if v > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:362
				_go_fuzz_dep_.CoverTab[89707]++
															s.symbolLen = uint16(i) + 1
															if i >= len(s.prevTable) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:364
					_go_fuzz_dep_.CoverTab[89708]++
																reuse = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:365
					// _ = "end of CoverTab[89708]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:366
					_go_fuzz_dep_.CoverTab[89709]++
																if s.prevTable[i].nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:367
						_go_fuzz_dep_.CoverTab[89710]++
																	reuse = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:368
						// _ = "end of CoverTab[89710]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:369
						_go_fuzz_dep_.CoverTab[89711]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:369
						// _ = "end of CoverTab[89711]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:369
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:369
					// _ = "end of CoverTab[89709]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:370
				// _ = "end of CoverTab[89707]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:371
				_go_fuzz_dep_.CoverTab[89712]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:371
				// _ = "end of CoverTab[89712]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:371
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:371
			// _ = "end of CoverTab[89704]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:372
		// _ = "end of CoverTab[89701]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:372
		_go_fuzz_dep_.CoverTab[89702]++
													return int(m), reuse
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:373
		// _ = "end of CoverTab[89702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:374
		_go_fuzz_dep_.CoverTab[89713]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:374
		// _ = "end of CoverTab[89713]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:374
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:374
	// _ = "end of CoverTab[89697]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:374
	_go_fuzz_dep_.CoverTab[89698]++
												for i, v := range s.count[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:375
		_go_fuzz_dep_.CoverTab[89714]++
													if v > m {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:376
			_go_fuzz_dep_.CoverTab[89716]++
														m = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:377
			// _ = "end of CoverTab[89716]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:378
			_go_fuzz_dep_.CoverTab[89717]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:378
			// _ = "end of CoverTab[89717]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:378
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:378
		// _ = "end of CoverTab[89714]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:378
		_go_fuzz_dep_.CoverTab[89715]++
													if v > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:379
			_go_fuzz_dep_.CoverTab[89718]++
														s.symbolLen = uint16(i) + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:380
			// _ = "end of CoverTab[89718]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:381
			_go_fuzz_dep_.CoverTab[89719]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:381
			// _ = "end of CoverTab[89719]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:381
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:381
		// _ = "end of CoverTab[89715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:382
	// _ = "end of CoverTab[89698]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:382
	_go_fuzz_dep_.CoverTab[89699]++
												return int(m), false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:383
	// _ = "end of CoverTab[89699]"
}

func (s *Scratch) canUseTable(c cTable) bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:386
	_go_fuzz_dep_.CoverTab[89720]++
												if len(c) < int(s.symbolLen) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:387
		_go_fuzz_dep_.CoverTab[89723]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:388
		// _ = "end of CoverTab[89723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:389
		_go_fuzz_dep_.CoverTab[89724]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:389
		// _ = "end of CoverTab[89724]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:389
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:389
	// _ = "end of CoverTab[89720]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:389
	_go_fuzz_dep_.CoverTab[89721]++
												for i, v := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:390
		_go_fuzz_dep_.CoverTab[89725]++
													if v != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:391
			_go_fuzz_dep_.CoverTab[89726]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:391
			return c[i].nBits == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:391
			// _ = "end of CoverTab[89726]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:391
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:391
			_go_fuzz_dep_.CoverTab[89727]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:392
			// _ = "end of CoverTab[89727]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:393
			_go_fuzz_dep_.CoverTab[89728]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:393
			// _ = "end of CoverTab[89728]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:393
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:393
		// _ = "end of CoverTab[89725]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:394
	// _ = "end of CoverTab[89721]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:394
	_go_fuzz_dep_.CoverTab[89722]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:395
	// _ = "end of CoverTab[89722]"
}

func (s *Scratch) validateTable(c cTable) bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:398
	_go_fuzz_dep_.CoverTab[89729]++
												if len(c) < int(s.symbolLen) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:399
		_go_fuzz_dep_.CoverTab[89732]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:400
		// _ = "end of CoverTab[89732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:401
		_go_fuzz_dep_.CoverTab[89733]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:401
		// _ = "end of CoverTab[89733]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:401
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:401
	// _ = "end of CoverTab[89729]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:401
	_go_fuzz_dep_.CoverTab[89730]++
												for i, v := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:402
		_go_fuzz_dep_.CoverTab[89734]++
													if v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:403
			_go_fuzz_dep_.CoverTab[89735]++
														if c[i].nBits == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:404
				_go_fuzz_dep_.CoverTab[89737]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:405
				// _ = "end of CoverTab[89737]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:406
				_go_fuzz_dep_.CoverTab[89738]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:406
				// _ = "end of CoverTab[89738]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:406
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:406
			// _ = "end of CoverTab[89735]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:406
			_go_fuzz_dep_.CoverTab[89736]++
														if c[i].nBits > s.actualTableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:407
				_go_fuzz_dep_.CoverTab[89739]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:408
				// _ = "end of CoverTab[89739]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:409
				_go_fuzz_dep_.CoverTab[89740]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:409
				// _ = "end of CoverTab[89740]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:409
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:409
			// _ = "end of CoverTab[89736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:410
			_go_fuzz_dep_.CoverTab[89741]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:410
			// _ = "end of CoverTab[89741]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:410
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:410
		// _ = "end of CoverTab[89734]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:411
	// _ = "end of CoverTab[89730]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:411
	_go_fuzz_dep_.CoverTab[89731]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:412
	// _ = "end of CoverTab[89731]"
}

// minTableLog provides the minimum logSize to safely represent a distribution.
func (s *Scratch) minTableLog() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:416
	_go_fuzz_dep_.CoverTab[89742]++
												minBitsSrc := highBit32(uint32(s.br.remain())) + 1
												minBitsSymbols := highBit32(uint32(s.symbolLen-1)) + 2
												if minBitsSrc < minBitsSymbols {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:419
		_go_fuzz_dep_.CoverTab[89744]++
													return uint8(minBitsSrc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:420
		// _ = "end of CoverTab[89744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:421
		_go_fuzz_dep_.CoverTab[89745]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:421
		// _ = "end of CoverTab[89745]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:421
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:421
	// _ = "end of CoverTab[89742]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:421
	_go_fuzz_dep_.CoverTab[89743]++
												return uint8(minBitsSymbols)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:422
	// _ = "end of CoverTab[89743]"
}

// optimalTableLog calculates and sets the optimal tableLog in s.actualTableLog
func (s *Scratch) optimalTableLog() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:426
	_go_fuzz_dep_.CoverTab[89746]++
												tableLog := s.TableLog
												minBits := s.minTableLog()
												maxBitsSrc := uint8(highBit32(uint32(s.br.remain()-1))) - 1
												if maxBitsSrc < tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:430
		_go_fuzz_dep_.CoverTab[89751]++

													tableLog = maxBitsSrc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:432
		// _ = "end of CoverTab[89751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:433
		_go_fuzz_dep_.CoverTab[89752]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:433
		// _ = "end of CoverTab[89752]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:433
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:433
	// _ = "end of CoverTab[89746]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:433
	_go_fuzz_dep_.CoverTab[89747]++
												if minBits > tableLog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:434
		_go_fuzz_dep_.CoverTab[89753]++
													tableLog = minBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:435
		// _ = "end of CoverTab[89753]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:436
		_go_fuzz_dep_.CoverTab[89754]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:436
		// _ = "end of CoverTab[89754]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:436
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:436
	// _ = "end of CoverTab[89747]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:436
	_go_fuzz_dep_.CoverTab[89748]++

												if tableLog < minTablelog {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:438
		_go_fuzz_dep_.CoverTab[89755]++
													tableLog = minTablelog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:439
		// _ = "end of CoverTab[89755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:440
		_go_fuzz_dep_.CoverTab[89756]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:440
		// _ = "end of CoverTab[89756]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:440
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:440
	// _ = "end of CoverTab[89748]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:440
	_go_fuzz_dep_.CoverTab[89749]++
												if tableLog > tableLogMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:441
		_go_fuzz_dep_.CoverTab[89757]++
													tableLog = tableLogMax
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:442
		// _ = "end of CoverTab[89757]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:443
		_go_fuzz_dep_.CoverTab[89758]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:443
		// _ = "end of CoverTab[89758]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:443
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:443
	// _ = "end of CoverTab[89749]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:443
	_go_fuzz_dep_.CoverTab[89750]++
												s.actualTableLog = tableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:444
	// _ = "end of CoverTab[89750]"
}

type cTableEntry struct {
	val	uint16
	nBits	uint8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:451
}

const huffNodesMask = huffNodesLen - 1

func (s *Scratch) buildCTable() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:455
	_go_fuzz_dep_.CoverTab[89759]++
												s.optimalTableLog()
												s.huffSort()
												if cap(s.cTable) < maxSymbolValue+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:458
		_go_fuzz_dep_.CoverTab[89770]++
													s.cTable = make([]cTableEntry, s.symbolLen, maxSymbolValue+1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:459
		// _ = "end of CoverTab[89770]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:460
		_go_fuzz_dep_.CoverTab[89771]++
													s.cTable = s.cTable[:s.symbolLen]
													for i := range s.cTable {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:462
			_go_fuzz_dep_.CoverTab[89772]++
														s.cTable[i] = cTableEntry{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:463
			// _ = "end of CoverTab[89772]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:464
		// _ = "end of CoverTab[89771]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:465
	// _ = "end of CoverTab[89759]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:465
	_go_fuzz_dep_.CoverTab[89760]++

												var startNode = int16(s.symbolLen)
												nonNullRank := s.symbolLen - 1

												nodeNb := startNode
												huffNode := s.nodes[1 : huffNodesLen+1]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:475
	huffNode0 := s.nodes[0 : huffNodesLen+1]

	for huffNode[nonNullRank].count == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:477
		_go_fuzz_dep_.CoverTab[89773]++
													nonNullRank--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:478
		// _ = "end of CoverTab[89773]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:479
	// _ = "end of CoverTab[89760]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:479
	_go_fuzz_dep_.CoverTab[89761]++

												lowS := int16(nonNullRank)
												nodeRoot := nodeNb + lowS - 1
												lowN := nodeNb
												huffNode[nodeNb].count = huffNode[lowS].count + huffNode[lowS-1].count
												huffNode[lowS].parent, huffNode[lowS-1].parent = uint16(nodeNb), uint16(nodeNb)
												nodeNb++
												lowS -= 2
												for n := nodeNb; n <= nodeRoot; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:488
		_go_fuzz_dep_.CoverTab[89774]++
													huffNode[n].count = 1 << 30
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:489
		// _ = "end of CoverTab[89774]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:490
	// _ = "end of CoverTab[89761]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:490
	_go_fuzz_dep_.CoverTab[89762]++

												huffNode0[0].count = 1 << 31

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:495
	for nodeNb <= nodeRoot {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:495
		_go_fuzz_dep_.CoverTab[89775]++
													var n1, n2 int16
													if huffNode0[lowS+1].count < huffNode0[lowN+1].count {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:497
			_go_fuzz_dep_.CoverTab[89778]++
														n1 = lowS
														lowS--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:499
			// _ = "end of CoverTab[89778]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:500
			_go_fuzz_dep_.CoverTab[89779]++
														n1 = lowN
														lowN++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:502
			// _ = "end of CoverTab[89779]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:503
		// _ = "end of CoverTab[89775]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:503
		_go_fuzz_dep_.CoverTab[89776]++
													if huffNode0[lowS+1].count < huffNode0[lowN+1].count {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:504
			_go_fuzz_dep_.CoverTab[89780]++
														n2 = lowS
														lowS--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:506
			// _ = "end of CoverTab[89780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:507
			_go_fuzz_dep_.CoverTab[89781]++
														n2 = lowN
														lowN++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:509
			// _ = "end of CoverTab[89781]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:510
		// _ = "end of CoverTab[89776]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:510
		_go_fuzz_dep_.CoverTab[89777]++

													huffNode[nodeNb].count = huffNode0[n1+1].count + huffNode0[n2+1].count
													huffNode0[n1+1].parent, huffNode0[n2+1].parent = uint16(nodeNb), uint16(nodeNb)
													nodeNb++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:514
		// _ = "end of CoverTab[89777]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:515
	// _ = "end of CoverTab[89762]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:515
	_go_fuzz_dep_.CoverTab[89763]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:518
	huffNode[nodeRoot].nbBits = 0
	for n := nodeRoot - 1; n >= startNode; n-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:519
		_go_fuzz_dep_.CoverTab[89782]++
													huffNode[n].nbBits = huffNode[huffNode[n].parent].nbBits + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:520
		// _ = "end of CoverTab[89782]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:521
	// _ = "end of CoverTab[89763]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:521
	_go_fuzz_dep_.CoverTab[89764]++
												for n := uint16(0); n <= nonNullRank; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:522
		_go_fuzz_dep_.CoverTab[89783]++
													huffNode[n].nbBits = huffNode[huffNode[n].parent].nbBits + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:523
		// _ = "end of CoverTab[89783]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:524
	// _ = "end of CoverTab[89764]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:524
	_go_fuzz_dep_.CoverTab[89765]++
												s.actualTableLog = s.setMaxHeight(int(nonNullRank))
												maxNbBits := s.actualTableLog

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:529
	if maxNbBits > tableLogMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:529
		_go_fuzz_dep_.CoverTab[89784]++
													return fmt.Errorf("internal error: maxNbBits (%d) > tableLogMax (%d)", maxNbBits, tableLogMax)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:530
		// _ = "end of CoverTab[89784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:531
		_go_fuzz_dep_.CoverTab[89785]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:531
		// _ = "end of CoverTab[89785]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:531
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:531
	// _ = "end of CoverTab[89765]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:531
	_go_fuzz_dep_.CoverTab[89766]++
												var nbPerRank [tableLogMax + 1]uint16
												var valPerRank [16]uint16
												for _, v := range huffNode[:nonNullRank+1] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:534
		_go_fuzz_dep_.CoverTab[89786]++
													nbPerRank[v.nbBits]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:535
		// _ = "end of CoverTab[89786]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:536
	// _ = "end of CoverTab[89766]"

												{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:538
		_go_fuzz_dep_.CoverTab[89787]++
													min := uint16(0)
													for n := maxNbBits; n > 0; n-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:540
			_go_fuzz_dep_.CoverTab[89788]++

														valPerRank[n] = min
														min += nbPerRank[n]
														min >>= 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:544
			// _ = "end of CoverTab[89788]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:545
		// _ = "end of CoverTab[89787]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:546
	_go_fuzz_dep_.CoverTab[89767]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:549
	for _, v := range huffNode[:nonNullRank+1] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:549
		_go_fuzz_dep_.CoverTab[89789]++
													s.cTable[v.symbol].nBits = v.nbBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:550
		// _ = "end of CoverTab[89789]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:551
	// _ = "end of CoverTab[89767]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:551
	_go_fuzz_dep_.CoverTab[89768]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:554
	t := s.cTable[:s.symbolLen]
	for n, val := range t {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:555
		_go_fuzz_dep_.CoverTab[89790]++
													nbits := val.nBits & 15
													v := valPerRank[nbits]
													t[n].val = v
													valPerRank[nbits] = v + 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:559
		// _ = "end of CoverTab[89790]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:560
	// _ = "end of CoverTab[89768]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:560
	_go_fuzz_dep_.CoverTab[89769]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:562
	// _ = "end of CoverTab[89769]"
}

// huffSort will sort symbols, decreasing order.
func (s *Scratch) huffSort() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:566
	_go_fuzz_dep_.CoverTab[89791]++
												type rankPos struct {
		base	uint32
		current	uint32
	}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:573
	nodes := s.nodes[:huffNodesLen+1]
	s.nodes = nodes
	nodes = nodes[1 : huffNodesLen+1]

	// Sort into buckets based on length of symbol count.
	var rank [32]rankPos
	for _, v := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:579
		_go_fuzz_dep_.CoverTab[89795]++
													r := highBit32(v+1) & 31
													rank[r].base++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:581
		// _ = "end of CoverTab[89795]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:582
	// _ = "end of CoverTab[89791]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:582
	_go_fuzz_dep_.CoverTab[89792]++
	// maxBitLength is log2(BlockSizeMax) + 1
	const maxBitLength = 18 + 1
	for n := maxBitLength; n > 0; n-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:585
		_go_fuzz_dep_.CoverTab[89796]++
													rank[n-1].base += rank[n].base
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:586
		// _ = "end of CoverTab[89796]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:587
	// _ = "end of CoverTab[89792]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:587
	_go_fuzz_dep_.CoverTab[89793]++
												for n := range rank[:maxBitLength] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:588
		_go_fuzz_dep_.CoverTab[89797]++
													rank[n].current = rank[n].base
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:589
		// _ = "end of CoverTab[89797]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:590
	// _ = "end of CoverTab[89793]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:590
	_go_fuzz_dep_.CoverTab[89794]++
												for n, c := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:591
		_go_fuzz_dep_.CoverTab[89798]++
													r := (highBit32(c+1) + 1) & 31
													pos := rank[r].current
													rank[r].current++
													prev := nodes[(pos-1)&huffNodesMask]
													for pos > rank[r].base && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:596
			_go_fuzz_dep_.CoverTab[89800]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:596
			return c > prev.count
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:596
			// _ = "end of CoverTab[89800]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:596
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:596
			_go_fuzz_dep_.CoverTab[89801]++
														nodes[pos&huffNodesMask] = prev
														pos--
														prev = nodes[(pos-1)&huffNodesMask]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:599
			// _ = "end of CoverTab[89801]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:600
		// _ = "end of CoverTab[89798]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:600
		_go_fuzz_dep_.CoverTab[89799]++
													nodes[pos&huffNodesMask] = nodeElt{count: c, symbol: byte(n)}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:601
		// _ = "end of CoverTab[89799]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:602
	// _ = "end of CoverTab[89794]"
}

func (s *Scratch) setMaxHeight(lastNonNull int) uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:605
	_go_fuzz_dep_.CoverTab[89802]++
												maxNbBits := s.actualTableLog
												huffNode := s.nodes[1 : huffNodesLen+1]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:610
	largestBits := huffNode[lastNonNull].nbBits

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:613
	if largestBits <= maxNbBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:613
		_go_fuzz_dep_.CoverTab[89807]++
													return largestBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:614
		// _ = "end of CoverTab[89807]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:615
		_go_fuzz_dep_.CoverTab[89808]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:615
		// _ = "end of CoverTab[89808]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:615
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:615
	// _ = "end of CoverTab[89802]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:615
	_go_fuzz_dep_.CoverTab[89803]++
												totalCost := int(0)
												baseCost := int(1) << (largestBits - maxNbBits)
												n := uint32(lastNonNull)

												for huffNode[n].nbBits > maxNbBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:620
		_go_fuzz_dep_.CoverTab[89809]++
													totalCost += baseCost - (1 << (largestBits - huffNode[n].nbBits))
													huffNode[n].nbBits = maxNbBits
													n--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:623
		// _ = "end of CoverTab[89809]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:624
	// _ = "end of CoverTab[89803]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:624
	_go_fuzz_dep_.CoverTab[89804]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:627
	for huffNode[n].nbBits == maxNbBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:627
		_go_fuzz_dep_.CoverTab[89810]++
													n--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:628
		// _ = "end of CoverTab[89810]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:629
	// _ = "end of CoverTab[89804]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:629
	_go_fuzz_dep_.CoverTab[89805]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:633
	totalCost >>= largestBits - maxNbBits

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:636
	{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:636
		_go_fuzz_dep_.CoverTab[89811]++
													const noSymbol = 0xF0F0F0F0
													var rankLast [tableLogMax + 2]uint32

													for i := range rankLast[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:640
			_go_fuzz_dep_.CoverTab[89814]++
														rankLast[i] = noSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:641
			// _ = "end of CoverTab[89814]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:642
		// _ = "end of CoverTab[89811]"

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:645
		{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:645
			_go_fuzz_dep_.CoverTab[89815]++
														currentNbBits := maxNbBits
														for pos := int(n); pos >= 0; pos-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:647
				_go_fuzz_dep_.CoverTab[89816]++
															if huffNode[pos].nbBits >= currentNbBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:648
					_go_fuzz_dep_.CoverTab[89818]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:649
					// _ = "end of CoverTab[89818]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:650
					_go_fuzz_dep_.CoverTab[89819]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:650
					// _ = "end of CoverTab[89819]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:650
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:650
				// _ = "end of CoverTab[89816]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:650
				_go_fuzz_dep_.CoverTab[89817]++
															currentNbBits = huffNode[pos].nbBits
															rankLast[maxNbBits-currentNbBits] = uint32(pos)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:652
				// _ = "end of CoverTab[89817]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:653
			// _ = "end of CoverTab[89815]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:654
		_go_fuzz_dep_.CoverTab[89812]++

													for totalCost > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:656
			_go_fuzz_dep_.CoverTab[89820]++
														nBitsToDecrease := uint8(highBit32(uint32(totalCost))) + 1

														for ; nBitsToDecrease > 1; nBitsToDecrease-- {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:659
				_go_fuzz_dep_.CoverTab[89824]++
															highPos := rankLast[nBitsToDecrease]
															lowPos := rankLast[nBitsToDecrease-1]
															if highPos == noSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:662
					_go_fuzz_dep_.CoverTab[89827]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:663
					// _ = "end of CoverTab[89827]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:664
					_go_fuzz_dep_.CoverTab[89828]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:664
					// _ = "end of CoverTab[89828]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:664
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:664
				// _ = "end of CoverTab[89824]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:664
				_go_fuzz_dep_.CoverTab[89825]++
															if lowPos == noSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:665
					_go_fuzz_dep_.CoverTab[89829]++
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:666
					// _ = "end of CoverTab[89829]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:667
					_go_fuzz_dep_.CoverTab[89830]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:667
					// _ = "end of CoverTab[89830]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:667
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:667
				// _ = "end of CoverTab[89825]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:667
				_go_fuzz_dep_.CoverTab[89826]++
															highTotal := huffNode[highPos].count
															lowTotal := 2 * huffNode[lowPos].count
															if highTotal <= lowTotal {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:670
					_go_fuzz_dep_.CoverTab[89831]++
																break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:671
					// _ = "end of CoverTab[89831]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:672
					_go_fuzz_dep_.CoverTab[89832]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:672
					// _ = "end of CoverTab[89832]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:672
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:672
				// _ = "end of CoverTab[89826]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:673
			// _ = "end of CoverTab[89820]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:673
			_go_fuzz_dep_.CoverTab[89821]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
			for (nBitsToDecrease <= tableLogMax) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
				_go_fuzz_dep_.CoverTab[89833]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
				return (rankLast[nBitsToDecrease] == noSymbol)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
				// _ = "end of CoverTab[89833]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:677
				_go_fuzz_dep_.CoverTab[89834]++
															nBitsToDecrease++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:678
				// _ = "end of CoverTab[89834]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:679
			// _ = "end of CoverTab[89821]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:679
			_go_fuzz_dep_.CoverTab[89822]++
														totalCost -= 1 << (nBitsToDecrease - 1)
														if rankLast[nBitsToDecrease-1] == noSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:681
				_go_fuzz_dep_.CoverTab[89835]++

															rankLast[nBitsToDecrease-1] = rankLast[nBitsToDecrease]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:683
				// _ = "end of CoverTab[89835]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:684
				_go_fuzz_dep_.CoverTab[89836]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:684
				// _ = "end of CoverTab[89836]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:684
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:684
			// _ = "end of CoverTab[89822]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:684
			_go_fuzz_dep_.CoverTab[89823]++
														huffNode[rankLast[nBitsToDecrease]].nbBits++
														if rankLast[nBitsToDecrease] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:686
				_go_fuzz_dep_.CoverTab[89837]++

															rankLast[nBitsToDecrease] = noSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:688
				// _ = "end of CoverTab[89837]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:689
				_go_fuzz_dep_.CoverTab[89838]++
															rankLast[nBitsToDecrease]--
															if huffNode[rankLast[nBitsToDecrease]].nbBits != maxNbBits-nBitsToDecrease {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:691
					_go_fuzz_dep_.CoverTab[89839]++
																rankLast[nBitsToDecrease] = noSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:692
					// _ = "end of CoverTab[89839]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:693
					_go_fuzz_dep_.CoverTab[89840]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:693
					// _ = "end of CoverTab[89840]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:693
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:693
				// _ = "end of CoverTab[89838]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:694
			// _ = "end of CoverTab[89823]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:695
		// _ = "end of CoverTab[89812]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:695
		_go_fuzz_dep_.CoverTab[89813]++

													for totalCost < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:697
			_go_fuzz_dep_.CoverTab[89841]++
														if rankLast[1] == noSymbol {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:698
				_go_fuzz_dep_.CoverTab[89843]++
															for huffNode[n].nbBits == maxNbBits {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:699
					_go_fuzz_dep_.CoverTab[89845]++
																n--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:700
					// _ = "end of CoverTab[89845]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:701
				// _ = "end of CoverTab[89843]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:701
				_go_fuzz_dep_.CoverTab[89844]++
															huffNode[n+1].nbBits--
															rankLast[1] = n + 1
															totalCost++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:705
				// _ = "end of CoverTab[89844]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:706
				_go_fuzz_dep_.CoverTab[89846]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:706
				// _ = "end of CoverTab[89846]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:706
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:706
			// _ = "end of CoverTab[89841]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:706
			_go_fuzz_dep_.CoverTab[89842]++
														huffNode[rankLast[1]+1].nbBits--
														rankLast[1]++
														totalCost++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:709
			// _ = "end of CoverTab[89842]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:710
		// _ = "end of CoverTab[89813]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:711
	// _ = "end of CoverTab[89805]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:711
	_go_fuzz_dep_.CoverTab[89806]++
												return maxNbBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:712
	// _ = "end of CoverTab[89806]"
}

type nodeElt struct {
	count	uint32
	parent	uint16
	symbol	byte
	nbBits	uint8
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:720
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/compress.go:720
var _ = _go_fuzz_dep_.CoverTab
