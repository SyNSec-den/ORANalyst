// Code generated by "stringer -type=RoundingMode"; DO NOT EDIT.

//line /usr/local/go/src/math/big/roundingmode_string.go:3
package big

//line /usr/local/go/src/math/big/roundingmode_string.go:3
import (
//line /usr/local/go/src/math/big/roundingmode_string.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/roundingmode_string.go:3
)
//line /usr/local/go/src/math/big/roundingmode_string.go:3
import (
//line /usr/local/go/src/math/big/roundingmode_string.go:3
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/roundingmode_string.go:3
)

import "strconv"

const _RoundingMode_name = "ToNearestEvenToNearestAwayToZeroAwayFromZeroToNegativeInfToPositiveInf"

var _RoundingMode_index = [...]uint8{0, 13, 26, 32, 44, 57, 70}

func (i RoundingMode) String() string {
//line /usr/local/go/src/math/big/roundingmode_string.go:11
	_go_fuzz_dep_.CoverTab[6908]++
								if i >= RoundingMode(len(_RoundingMode_index)-1) {
//line /usr/local/go/src/math/big/roundingmode_string.go:12
		_go_fuzz_dep_.CoverTab[6910]++
									return "RoundingMode(" + strconv.FormatInt(int64(i), 10) + ")"
//line /usr/local/go/src/math/big/roundingmode_string.go:13
		// _ = "end of CoverTab[6910]"
	} else {
//line /usr/local/go/src/math/big/roundingmode_string.go:14
		_go_fuzz_dep_.CoverTab[6911]++
//line /usr/local/go/src/math/big/roundingmode_string.go:14
		// _ = "end of CoverTab[6911]"
//line /usr/local/go/src/math/big/roundingmode_string.go:14
	}
//line /usr/local/go/src/math/big/roundingmode_string.go:14
	// _ = "end of CoverTab[6908]"
//line /usr/local/go/src/math/big/roundingmode_string.go:14
	_go_fuzz_dep_.CoverTab[6909]++
								return _RoundingMode_name[_RoundingMode_index[i]:_RoundingMode_index[i+1]]
//line /usr/local/go/src/math/big/roundingmode_string.go:15
	// _ = "end of CoverTab[6909]"
}

//line /usr/local/go/src/math/big/roundingmode_string.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/roundingmode_string.go:16
var _ = _go_fuzz_dep_.CoverTab