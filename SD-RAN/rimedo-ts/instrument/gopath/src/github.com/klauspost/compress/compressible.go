//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
package compress

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:1
)

import "math"

// Estimate returns a normalized compressibility estimate of block b.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:5
// Values close to zero are likely uncompressible.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:5
// Values above 0.1 are likely to be compressible.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:5
// Values above 0.5 are very compressible.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:5
// Very small lengths will return 0.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:10
func Estimate(b []byte) float64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:10
	_go_fuzz_dep_.CoverTab[88930]++
												if len(b) < 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:11
		_go_fuzz_dep_.CoverTab[88935]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:12
		// _ = "end of CoverTab[88935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:13
		_go_fuzz_dep_.CoverTab[88936]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:13
		// _ = "end of CoverTab[88936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:13
	// _ = "end of CoverTab[88930]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:13
	_go_fuzz_dep_.CoverTab[88931]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:16
	hits := 0
	lastMatch := false
	var o1 [256]byte
	var hist [256]int
	c1 := byte(0)
	for _, c := range b {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:21
		_go_fuzz_dep_.CoverTab[88937]++
													if c == o1[c1] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:22
			_go_fuzz_dep_.CoverTab[88939]++

														if lastMatch {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:24
				_go_fuzz_dep_.CoverTab[88941]++
															hits++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:25
				// _ = "end of CoverTab[88941]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:26
				_go_fuzz_dep_.CoverTab[88942]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:26
				// _ = "end of CoverTab[88942]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:26
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:26
			// _ = "end of CoverTab[88939]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:26
			_go_fuzz_dep_.CoverTab[88940]++
														lastMatch = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:27
			// _ = "end of CoverTab[88940]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:28
			_go_fuzz_dep_.CoverTab[88943]++
														lastMatch = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:29
			// _ = "end of CoverTab[88943]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:30
		// _ = "end of CoverTab[88937]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:30
		_go_fuzz_dep_.CoverTab[88938]++
													o1[c1] = c
													c1 = c
													hist[c]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:33
		// _ = "end of CoverTab[88938]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:34
	// _ = "end of CoverTab[88931]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:34
	_go_fuzz_dep_.CoverTab[88932]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:37
	prediction := math.Pow(float64(hits)/float64(len(b)), 0.6)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:40
	variance := float64(0)
	avg := float64(len(b)) / 256

	for _, v := range hist {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:43
		_go_fuzz_dep_.CoverTab[88944]++
													Δ := float64(v) - avg
													variance += Δ * Δ
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:45
		// _ = "end of CoverTab[88944]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:46
	// _ = "end of CoverTab[88932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:46
	_go_fuzz_dep_.CoverTab[88933]++

												stddev := math.Sqrt(float64(variance)) / float64(len(b))
												exp := math.Sqrt(1 / float64(len(b)))

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:52
	stddev -= exp
	if stddev < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:53
		_go_fuzz_dep_.CoverTab[88945]++
													stddev = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:54
		// _ = "end of CoverTab[88945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:55
		_go_fuzz_dep_.CoverTab[88946]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:55
		// _ = "end of CoverTab[88946]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:55
	// _ = "end of CoverTab[88933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:55
	_go_fuzz_dep_.CoverTab[88934]++
												stddev *= 1 + exp

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:59
	entropy := math.Pow(stddev, 0.4)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:62
	return math.Pow((prediction+entropy)/2, 0.9)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:62
	// _ = "end of CoverTab[88934]"
}

// ShannonEntropyBits returns the number of bits minimum required to represent
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:65
// an entropy encoding of the input bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:65
// https://en.wiktionary.org/wiki/Shannon_entropy
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:68
func ShannonEntropyBits(b []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:68
	_go_fuzz_dep_.CoverTab[88947]++
												if len(b) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:69
		_go_fuzz_dep_.CoverTab[88951]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:70
		// _ = "end of CoverTab[88951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:71
		_go_fuzz_dep_.CoverTab[88952]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:71
		// _ = "end of CoverTab[88952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:71
	// _ = "end of CoverTab[88947]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:71
	_go_fuzz_dep_.CoverTab[88948]++
												var hist [256]int
												for _, c := range b {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:73
		_go_fuzz_dep_.CoverTab[88953]++
													hist[c]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:74
		// _ = "end of CoverTab[88953]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:75
	// _ = "end of CoverTab[88948]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:75
	_go_fuzz_dep_.CoverTab[88949]++
												shannon := float64(0)
												invTotal := 1.0 / float64(len(b))
												for _, v := range hist[:] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:78
		_go_fuzz_dep_.CoverTab[88954]++
													if v > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:79
			_go_fuzz_dep_.CoverTab[88955]++
														n := float64(v)
														shannon += math.Ceil(-math.Log2(n*invTotal) * n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:81
			// _ = "end of CoverTab[88955]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:82
			_go_fuzz_dep_.CoverTab[88956]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:82
			// _ = "end of CoverTab[88956]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:82
		// _ = "end of CoverTab[88954]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:83
	// _ = "end of CoverTab[88949]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:83
	_go_fuzz_dep_.CoverTab[88950]++
												return int(math.Ceil(shannon))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:84
	// _ = "end of CoverTab[88950]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/compressible.go:85
var _ = _go_fuzz_dep_.CoverTab
