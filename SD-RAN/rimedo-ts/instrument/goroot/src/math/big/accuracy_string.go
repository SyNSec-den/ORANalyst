// Code generated by "stringer -type=Accuracy"; DO NOT EDIT.

//line /usr/local/go/src/math/big/accuracy_string.go:3
package big

//line /usr/local/go/src/math/big/accuracy_string.go:3
import (
//line /usr/local/go/src/math/big/accuracy_string.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/accuracy_string.go:3
)
//line /usr/local/go/src/math/big/accuracy_string.go:3
import (
//line /usr/local/go/src/math/big/accuracy_string.go:3
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/accuracy_string.go:3
)

import "strconv"

const _Accuracy_name = "BelowExactAbove"

var _Accuracy_index = [...]uint8{0, 5, 10, 15}

func (i Accuracy) String() string {
//line /usr/local/go/src/math/big/accuracy_string.go:11
	_go_fuzz_dep_.CoverTab[4040]++
								i -= -1
								if i < 0 || func() bool {
//line /usr/local/go/src/math/big/accuracy_string.go:13
		_go_fuzz_dep_.CoverTab[4042]++
//line /usr/local/go/src/math/big/accuracy_string.go:13
		return i >= Accuracy(len(_Accuracy_index)-1)
//line /usr/local/go/src/math/big/accuracy_string.go:13
		// _ = "end of CoverTab[4042]"
//line /usr/local/go/src/math/big/accuracy_string.go:13
	}() {
//line /usr/local/go/src/math/big/accuracy_string.go:13
		_go_fuzz_dep_.CoverTab[4043]++
									return "Accuracy(" + strconv.FormatInt(int64(i+-1), 10) + ")"
//line /usr/local/go/src/math/big/accuracy_string.go:14
		// _ = "end of CoverTab[4043]"
	} else {
//line /usr/local/go/src/math/big/accuracy_string.go:15
		_go_fuzz_dep_.CoverTab[4044]++
//line /usr/local/go/src/math/big/accuracy_string.go:15
		// _ = "end of CoverTab[4044]"
//line /usr/local/go/src/math/big/accuracy_string.go:15
	}
//line /usr/local/go/src/math/big/accuracy_string.go:15
	// _ = "end of CoverTab[4040]"
//line /usr/local/go/src/math/big/accuracy_string.go:15
	_go_fuzz_dep_.CoverTab[4041]++
								return _Accuracy_name[_Accuracy_index[i]:_Accuracy_index[i+1]]
//line /usr/local/go/src/math/big/accuracy_string.go:16
	// _ = "end of CoverTab[4041]"
}

//line /usr/local/go/src/math/big/accuracy_string.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/accuracy_string.go:17
var _ = _go_fuzz_dep_.CoverTab