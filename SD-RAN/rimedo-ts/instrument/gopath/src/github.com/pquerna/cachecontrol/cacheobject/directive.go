//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
package cacheobject

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:18
)

import (
	"errors"
	"math"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
)

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:31
var (
	ErrQuoteMismatch		= errors.New("Missing closing quote")
	ErrMaxAgeDeltaSeconds		= errors.New("Failed to parse delta-seconds in `max-age`")
	ErrSMaxAgeDeltaSeconds		= errors.New("Failed to parse delta-seconds in `s-maxage`")
	ErrMaxStaleDeltaSeconds		= errors.New("Failed to parse delta-seconds in `min-fresh`")
	ErrMinFreshDeltaSeconds		= errors.New("Failed to parse delta-seconds in `min-fresh`")
	ErrNoCacheNoArgs		= errors.New("Unexpected argument to `no-cache`")
	ErrNoStoreNoArgs		= errors.New("Unexpected argument to `no-store`")
	ErrNoTransformNoArgs		= errors.New("Unexpected argument to `no-transform`")
	ErrOnlyIfCachedNoArgs		= errors.New("Unexpected argument to `only-if-cached`")
	ErrMustRevalidateNoArgs		= errors.New("Unexpected argument to `must-revalidate`")
	ErrPublicNoArgs			= errors.New("Unexpected argument to `public`")
	ErrProxyRevalidateNoArgs	= errors.New("Unexpected argument to `proxy-revalidate`")
	// Experimental
	ErrImmutableNoArgs			= errors.New("Unexpected argument to `immutable`")
	ErrStaleIfErrorDeltaSeconds		= errors.New("Failed to parse delta-seconds in `stale-if-error`")
	ErrStaleWhileRevalidateDeltaSeconds	= errors.New("Failed to parse delta-seconds in `stale-while-revalidate`")
)

func whitespace(b byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:50
	_go_fuzz_dep_.CoverTab[183705]++
																	if b == '\t' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:51
		_go_fuzz_dep_.CoverTab[183707]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:51
		return b == ' '
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:51
		// _ = "end of CoverTab[183707]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:51
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:51
		_go_fuzz_dep_.CoverTab[183708]++
																		return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:52
		// _ = "end of CoverTab[183708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:53
		_go_fuzz_dep_.CoverTab[183709]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:53
		// _ = "end of CoverTab[183709]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:53
	// _ = "end of CoverTab[183705]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:53
	_go_fuzz_dep_.CoverTab[183706]++
																	return false
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:54
	// _ = "end of CoverTab[183706]"
}

func parse(value string, cd cacheDirective) error {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:57
	_go_fuzz_dep_.CoverTab[183710]++
																	var err error = nil
																	i := 0

																	for i < len(value) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:61
		_go_fuzz_dep_.CoverTab[183712]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:61
		return err == nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:61
		// _ = "end of CoverTab[183712]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:61
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:61
		_go_fuzz_dep_.CoverTab[183713]++

																		if whitespace(value[i]) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:63
			_go_fuzz_dep_.CoverTab[183716]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:63
			return value[i] == ','
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:63
			// _ = "end of CoverTab[183716]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:63
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:63
			_go_fuzz_dep_.CoverTab[183717]++
																			i++
																			continue
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:65
			// _ = "end of CoverTab[183717]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:66
			_go_fuzz_dep_.CoverTab[183718]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:66
			// _ = "end of CoverTab[183718]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:66
		// _ = "end of CoverTab[183713]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:66
		_go_fuzz_dep_.CoverTab[183714]++

																		j := i + 1

																		for j < len(value) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:70
			_go_fuzz_dep_.CoverTab[183719]++
																			if !isToken(value[j]) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:71
				_go_fuzz_dep_.CoverTab[183721]++
																				break
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:72
				// _ = "end of CoverTab[183721]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:73
				_go_fuzz_dep_.CoverTab[183722]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:73
				// _ = "end of CoverTab[183722]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:73
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:73
			// _ = "end of CoverTab[183719]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:73
			_go_fuzz_dep_.CoverTab[183720]++
																			j++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:74
			// _ = "end of CoverTab[183720]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:75
		// _ = "end of CoverTab[183714]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:75
		_go_fuzz_dep_.CoverTab[183715]++

																		token := strings.ToLower(value[i:j])
																		tokenHasFields := hasFieldNames(token)

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
		if j+1 < len(value) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
			_go_fuzz_dep_.CoverTab[183723]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
			return value[j] == '='
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
			// _ = "end of CoverTab[183723]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:86
			_go_fuzz_dep_.CoverTab[183724]++
																			k := j + 1

																			if k < len(value) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:89
				_go_fuzz_dep_.CoverTab[183725]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:89
				return value[k] == '"'
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:89
				// _ = "end of CoverTab[183725]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:89
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:89
				_go_fuzz_dep_.CoverTab[183726]++
																				eaten, result := httpUnquote(value[k:])
																				if eaten == -1 {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:91
					_go_fuzz_dep_.CoverTab[183728]++
																					return ErrQuoteMismatch
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:92
					// _ = "end of CoverTab[183728]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:93
					_go_fuzz_dep_.CoverTab[183729]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:93
					// _ = "end of CoverTab[183729]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:93
				}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:93
				// _ = "end of CoverTab[183726]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:93
				_go_fuzz_dep_.CoverTab[183727]++
																				i = k + eaten

																				err = cd.addPair(token, result)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:96
				// _ = "end of CoverTab[183727]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:97
				_go_fuzz_dep_.CoverTab[183730]++
																				z := k
																				for z < len(value) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:99
					_go_fuzz_dep_.CoverTab[183733]++
																					if tokenHasFields {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:100
						_go_fuzz_dep_.CoverTab[183735]++
																						if whitespace(value[z]) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:101
							_go_fuzz_dep_.CoverTab[183736]++
																							break
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:102
							// _ = "end of CoverTab[183736]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:103
							_go_fuzz_dep_.CoverTab[183737]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:103
							// _ = "end of CoverTab[183737]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:103
						}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:103
						// _ = "end of CoverTab[183735]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:104
						_go_fuzz_dep_.CoverTab[183738]++
																						if whitespace(value[z]) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:105
							_go_fuzz_dep_.CoverTab[183739]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:105
							return value[z] == ','
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:105
							// _ = "end of CoverTab[183739]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:105
						}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:105
							_go_fuzz_dep_.CoverTab[183740]++
																							break
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:106
							// _ = "end of CoverTab[183740]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:107
							_go_fuzz_dep_.CoverTab[183741]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:107
							// _ = "end of CoverTab[183741]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:107
						}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:107
						// _ = "end of CoverTab[183738]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:108
					// _ = "end of CoverTab[183733]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:108
					_go_fuzz_dep_.CoverTab[183734]++
																					z++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:109
					// _ = "end of CoverTab[183734]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:110
				// _ = "end of CoverTab[183730]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:110
				_go_fuzz_dep_.CoverTab[183731]++
																				i = z

																				result := value[k:z]
																				if result != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:114
					_go_fuzz_dep_.CoverTab[183742]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:114
					return result[len(result)-1] == ','
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:114
					// _ = "end of CoverTab[183742]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:114
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:114
					_go_fuzz_dep_.CoverTab[183743]++
																					result = result[:len(result)-1]
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:115
					// _ = "end of CoverTab[183743]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:116
					_go_fuzz_dep_.CoverTab[183744]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:116
					// _ = "end of CoverTab[183744]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:116
				}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:116
				// _ = "end of CoverTab[183731]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:116
				_go_fuzz_dep_.CoverTab[183732]++

																				err = cd.addPair(token, result)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:118
				// _ = "end of CoverTab[183732]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:119
			// _ = "end of CoverTab[183724]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:120
			_go_fuzz_dep_.CoverTab[183745]++
																			if token != "," {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:121
				_go_fuzz_dep_.CoverTab[183747]++
																				err = cd.addToken(token)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:122
				// _ = "end of CoverTab[183747]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:123
				_go_fuzz_dep_.CoverTab[183748]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:123
				// _ = "end of CoverTab[183748]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:123
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:123
			// _ = "end of CoverTab[183745]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:123
			_go_fuzz_dep_.CoverTab[183746]++
																			i = j
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:124
			// _ = "end of CoverTab[183746]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:125
		// _ = "end of CoverTab[183715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:126
	// _ = "end of CoverTab[183710]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:126
	_go_fuzz_dep_.CoverTab[183711]++

																	return err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:128
	// _ = "end of CoverTab[183711]"
}

// DeltaSeconds specifies a non-negative integer, representing
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:131
// time in seconds: http://tools.ietf.org/html/rfc7234#section-1.2.1
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:131
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:131
// When set to -1, this means unset.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:136
type DeltaSeconds int32

// Parser for delta-seconds, a uint31, more or less:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:138
// http://tools.ietf.org/html/rfc7234#section-1.2.1
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:140
func parseDeltaSeconds(v string) (DeltaSeconds, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:140
	_go_fuzz_dep_.CoverTab[183749]++
																	n, err := strconv.ParseUint(v, 10, 32)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:142
		_go_fuzz_dep_.CoverTab[183750]++
																		if numError, ok := err.(*strconv.NumError); ok {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:143
			_go_fuzz_dep_.CoverTab[183752]++
																			if numError.Err == strconv.ErrRange {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:144
				_go_fuzz_dep_.CoverTab[183753]++
																				return DeltaSeconds(math.MaxInt32), nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:145
				// _ = "end of CoverTab[183753]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:146
				_go_fuzz_dep_.CoverTab[183754]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:146
				// _ = "end of CoverTab[183754]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:146
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:146
			// _ = "end of CoverTab[183752]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:147
			_go_fuzz_dep_.CoverTab[183755]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:147
			// _ = "end of CoverTab[183755]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:147
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:147
		// _ = "end of CoverTab[183750]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:147
		_go_fuzz_dep_.CoverTab[183751]++
																		return DeltaSeconds(-1), err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:148
		// _ = "end of CoverTab[183751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:149
		_go_fuzz_dep_.CoverTab[183756]++
																		if n > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:150
			_go_fuzz_dep_.CoverTab[183757]++
																			return DeltaSeconds(math.MaxInt32), nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:151
			// _ = "end of CoverTab[183757]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:152
			_go_fuzz_dep_.CoverTab[183758]++
																			return DeltaSeconds(n), nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:153
			// _ = "end of CoverTab[183758]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:154
		// _ = "end of CoverTab[183756]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:155
	// _ = "end of CoverTab[183749]"
}

// Fields present in a header.
type FieldNames map[string]bool

// internal interface for shared methods of RequestCacheDirectives and ResponseCacheDirectives
type cacheDirective interface {
	addToken(s string) error
	addPair(s string, v string) error
}

// LOW LEVEL API: Repersentation of possible request directives in a `Cache-Control` header: http://tools.ietf.org/html/rfc7234#section-5.2.1
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:167
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:167
// Note: Many fields will be `nil` in practice.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:171
type RequestCacheDirectives struct {

	// max-age(delta seconds): http://tools.ietf.org/html/rfc7234#section-5.2.1.1
	//
	// The "max-age" request directive indicates that the client is
	// unwilling to accept a response whose age is greater than the
	// specified number of seconds.  Unless the max-stale request directive
	// is also present, the client is not willing to accept a stale
	// response.
	MaxAge	DeltaSeconds

	// max-stale(delta seconds): http://tools.ietf.org/html/rfc7234#section-5.2.1.2
	//
	// The "max-stale" request directive indicates that the client is
	// willing to accept a response that has exceeded its freshness
	// lifetime.  If max-stale is assigned a value, then the client is
	// willing to accept a response that has exceeded its freshness lifetime
	// by no more than the specified number of seconds.  If no value is
	// assigned to max-stale, then the client is willing to accept a stale
	// response of any age.
	MaxStale	DeltaSeconds

	// min-fresh(delta seconds): http://tools.ietf.org/html/rfc7234#section-5.2.1.3
	//
	// The "min-fresh" request directive indicates that the client is
	// willing to accept a response whose freshness lifetime is no less than
	// its current age plus the specified time in seconds.  That is, the
	// client wants a response that will still be fresh for at least the
	// specified number of seconds.
	MinFresh	DeltaSeconds

	// no-cache(bool): http://tools.ietf.org/html/rfc7234#section-5.2.1.4
	//
	// The "no-cache" request directive indicates that a cache MUST NOT use
	// a stored response to satisfy the request without successful
	// validation on the origin server.
	NoCache	bool

	// no-store(bool): http://tools.ietf.org/html/rfc7234#section-5.2.1.5
	//
	// The "no-store" request directive indicates that a cache MUST NOT
	// store any part of either this request or any response to it.  This
	// directive applies to both private and shared caches.
	NoStore	bool

	// no-transform(bool): http://tools.ietf.org/html/rfc7234#section-5.2.1.6
	//
	// The "no-transform" request directive indicates that an intermediary
	// (whether or not it implements a cache) MUST NOT transform the
	// payload, as defined in Section 5.7.2 of RFC7230.
	NoTransform	bool

	// only-if-cached(bool): http://tools.ietf.org/html/rfc7234#section-5.2.1.7
	//
	// The "only-if-cached" request directive indicates that the client only
	// wishes to obtain a stored response.
	OnlyIfCached	bool

	// Extensions: http://tools.ietf.org/html/rfc7234#section-5.2.3
	//
	// The Cache-Control header field can be extended through the use of one
	// or more cache-extension tokens, each with an optional value.  A cache
	// MUST ignore unrecognized cache directives.
	Extensions	[]string
}

func (cd *RequestCacheDirectives) addToken(token string) error {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:237
	_go_fuzz_dep_.CoverTab[183759]++
																	var err error = nil

																	switch token {
	case "max-age":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:241
		_go_fuzz_dep_.CoverTab[183761]++
																		err = ErrMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:242
		// _ = "end of CoverTab[183761]"
	case "max-stale":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:243
		_go_fuzz_dep_.CoverTab[183762]++
																		err = ErrMaxStaleDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:244
		// _ = "end of CoverTab[183762]"
	case "min-fresh":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:245
		_go_fuzz_dep_.CoverTab[183763]++
																		err = ErrMinFreshDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:246
		// _ = "end of CoverTab[183763]"
	case "no-cache":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:247
		_go_fuzz_dep_.CoverTab[183764]++
																		cd.NoCache = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:248
		// _ = "end of CoverTab[183764]"
	case "no-store":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:249
		_go_fuzz_dep_.CoverTab[183765]++
																		cd.NoStore = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:250
		// _ = "end of CoverTab[183765]"
	case "no-transform":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:251
		_go_fuzz_dep_.CoverTab[183766]++
																		cd.NoTransform = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:252
		// _ = "end of CoverTab[183766]"
	case "only-if-cached":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:253
		_go_fuzz_dep_.CoverTab[183767]++
																		cd.OnlyIfCached = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:254
		// _ = "end of CoverTab[183767]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:255
		_go_fuzz_dep_.CoverTab[183768]++
																		cd.Extensions = append(cd.Extensions, token)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:256
		// _ = "end of CoverTab[183768]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:257
	// _ = "end of CoverTab[183759]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:257
	_go_fuzz_dep_.CoverTab[183760]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:258
	// _ = "end of CoverTab[183760]"
}

func (cd *RequestCacheDirectives) addPair(token string, v string) error {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:261
	_go_fuzz_dep_.CoverTab[183769]++
																	var err error = nil

																	switch token {
	case "max-age":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:265
		_go_fuzz_dep_.CoverTab[183771]++
																		cd.MaxAge, err = parseDeltaSeconds(v)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:267
			_go_fuzz_dep_.CoverTab[183779]++
																			err = ErrMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:268
			// _ = "end of CoverTab[183779]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:269
			_go_fuzz_dep_.CoverTab[183780]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:269
			// _ = "end of CoverTab[183780]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:269
		// _ = "end of CoverTab[183771]"
	case "max-stale":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:270
		_go_fuzz_dep_.CoverTab[183772]++
																		cd.MaxStale, err = parseDeltaSeconds(v)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:272
			_go_fuzz_dep_.CoverTab[183781]++
																			err = ErrMaxStaleDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:273
			// _ = "end of CoverTab[183781]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:274
			_go_fuzz_dep_.CoverTab[183782]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:274
			// _ = "end of CoverTab[183782]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:274
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:274
		// _ = "end of CoverTab[183772]"
	case "min-fresh":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:275
		_go_fuzz_dep_.CoverTab[183773]++
																		cd.MinFresh, err = parseDeltaSeconds(v)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:277
			_go_fuzz_dep_.CoverTab[183783]++
																			err = ErrMinFreshDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:278
			// _ = "end of CoverTab[183783]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:279
			_go_fuzz_dep_.CoverTab[183784]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:279
			// _ = "end of CoverTab[183784]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:279
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:279
		// _ = "end of CoverTab[183773]"
	case "no-cache":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:280
		_go_fuzz_dep_.CoverTab[183774]++
																		err = ErrNoCacheNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:281
		// _ = "end of CoverTab[183774]"
	case "no-store":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:282
		_go_fuzz_dep_.CoverTab[183775]++
																		err = ErrNoStoreNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:283
		// _ = "end of CoverTab[183775]"
	case "no-transform":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:284
		_go_fuzz_dep_.CoverTab[183776]++
																		err = ErrNoTransformNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:285
		// _ = "end of CoverTab[183776]"
	case "only-if-cached":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:286
		_go_fuzz_dep_.CoverTab[183777]++
																		err = ErrOnlyIfCachedNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:287
		// _ = "end of CoverTab[183777]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:288
		_go_fuzz_dep_.CoverTab[183778]++

																		cd.Extensions = append(cd.Extensions, token+"="+v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:290
		// _ = "end of CoverTab[183778]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:291
	// _ = "end of CoverTab[183769]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:291
	_go_fuzz_dep_.CoverTab[183770]++

																	return err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:293
	// _ = "end of CoverTab[183770]"
}

// LOW LEVEL API: Parses a Cache Control Header from a Request into a set of directives.
func ParseRequestCacheControl(value string) (*RequestCacheDirectives, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:297
	_go_fuzz_dep_.CoverTab[183785]++
																	cd := &RequestCacheDirectives{
		MaxAge:		-1,
		MaxStale:	-1,
		MinFresh:	-1,
	}

	err := parse(value, cd)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:305
		_go_fuzz_dep_.CoverTab[183787]++
																		return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:306
		// _ = "end of CoverTab[183787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:307
		_go_fuzz_dep_.CoverTab[183788]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:307
		// _ = "end of CoverTab[183788]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:307
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:307
	// _ = "end of CoverTab[183785]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:307
	_go_fuzz_dep_.CoverTab[183786]++
																	return cd, nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:308
	// _ = "end of CoverTab[183786]"
}

// LOW LEVEL API: Repersentation of possible response directives in a `Cache-Control` header: http://tools.ietf.org/html/rfc7234#section-5.2.2
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:311
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:311
// Note: Many fields will be `nil` in practice.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:315
type ResponseCacheDirectives struct {

	// must-revalidate(bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.1
	//
	// The "must-revalidate" response directive indicates that once it has
	// become stale, a cache MUST NOT use the response to satisfy subsequent
	// requests without successful validation on the origin server.
	MustRevalidate	bool

	// no-cache(FieldName): http://tools.ietf.org/html/rfc7234#section-5.2.2.2
	//
	// The "no-cache" response directive indicates that the response MUST
	// NOT be used to satisfy a subsequent request without successful
	// validation on the origin server.
	//
	// If the no-cache response directive specifies one or more field-names,
	// then a cache MAY use the response to satisfy a subsequent request,
	// subject to any other restrictions on caching.  However, any header
	// fields in the response that have the field-name(s) listed MUST NOT be
	// sent in the response to a subsequent request without successful
	// revalidation with the origin server.
	NoCache	FieldNames

	// no-cache(cast-to-bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.2
	//
	// While the RFC defines optional field-names on a no-cache directive,
	// many applications only want to know if any no-cache directives were
	// present at all.
	NoCachePresent	bool

	// no-store(bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.3
	//
	// The "no-store" request directive indicates that a cache MUST NOT
	// store any part of either this request or any response to it.  This
	// directive applies to both private and shared caches.
	NoStore	bool

	// no-transform(bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.4
	//
	// The "no-transform" response directive indicates that an intermediary
	// (regardless of whether it implements a cache) MUST NOT transform the
	// payload, as defined in Section 5.7.2 of RFC7230.
	NoTransform	bool

	// public(bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.5
	//
	// The "public" response directive indicates that any cache MAY store
	// the response, even if the response would normally be non-cacheable or
	// cacheable only within a private cache.
	Public	bool

	// private(FieldName): http://tools.ietf.org/html/rfc7234#section-5.2.2.6
	//
	// The "private" response directive indicates that the response message
	// is intended for a single user and MUST NOT be stored by a shared
	// cache.  A private cache MAY store the response and reuse it for later
	// requests, even if the response would normally be non-cacheable.
	//
	// If the private response directive specifies one or more field-names,
	// this requirement is limited to the field-values associated with the
	// listed response header fields.  That is, a shared cache MUST NOT
	// store the specified field-names(s), whereas it MAY store the
	// remainder of the response message.
	Private	FieldNames

	// private(cast-to-bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.6
	//
	// While the RFC defines optional field-names on a private directive,
	// many applications only want to know if any private directives were
	// present at all.
	PrivatePresent	bool

	// proxy-revalidate(bool): http://tools.ietf.org/html/rfc7234#section-5.2.2.7
	//
	// The "proxy-revalidate" response directive has the same meaning as the
	// must-revalidate response directive, except that it does not apply to
	// private caches.
	ProxyRevalidate	bool

	// max-age(delta seconds): http://tools.ietf.org/html/rfc7234#section-5.2.2.8
	//
	// The "max-age" response directive indicates that the response is to be
	// considered stale after its age is greater than the specified number
	// of seconds.
	MaxAge	DeltaSeconds

	// s-maxage(delta seconds): http://tools.ietf.org/html/rfc7234#section-5.2.2.9
	//
	// The "s-maxage" response directive indicates that, in shared caches,
	// the maximum age specified by this directive overrides the maximum age
	// specified by either the max-age directive or the Expires header
	// field.  The s-maxage directive also implies the semantics of the
	// proxy-revalidate response directive.
	SMaxAge	DeltaSeconds

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:416
	// immutable(cast-to-bool): experimental feature
																	Immutable	bool

	// stale-if-error(delta seconds): experimental feature
	StaleIfError	DeltaSeconds

	// stale-while-revalidate(delta seconds): experimental feature
	StaleWhileRevalidate	DeltaSeconds

	// Extensions: http://tools.ietf.org/html/rfc7234#section-5.2.3
	//
	// The Cache-Control header field can be extended through the use of one
	// or more cache-extension tokens, each with an optional value.  A cache
	// MUST ignore unrecognized cache directives.
	Extensions	[]string
}

// LOW LEVEL API: Parses a Cache Control Header from a Response into a set of directives.
func ParseResponseCacheControl(value string) (*ResponseCacheDirectives, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:434
	_go_fuzz_dep_.CoverTab[183789]++
																	cd := &ResponseCacheDirectives{
		MaxAge:		-1,
		SMaxAge:	-1,

		StaleIfError:		-1,
		StaleWhileRevalidate:	-1,
	}

	err := parse(value, cd)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:444
		_go_fuzz_dep_.CoverTab[183791]++
																		return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:445
		// _ = "end of CoverTab[183791]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:446
		_go_fuzz_dep_.CoverTab[183792]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:446
		// _ = "end of CoverTab[183792]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:446
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:446
	// _ = "end of CoverTab[183789]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:446
	_go_fuzz_dep_.CoverTab[183790]++
																	return cd, nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:447
	// _ = "end of CoverTab[183790]"
}

func (cd *ResponseCacheDirectives) addToken(token string) error {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:450
	_go_fuzz_dep_.CoverTab[183793]++
																	var err error = nil
																	switch token {
	case "must-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:453
		_go_fuzz_dep_.CoverTab[183795]++
																		cd.MustRevalidate = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:454
		// _ = "end of CoverTab[183795]"
	case "no-cache":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:455
		_go_fuzz_dep_.CoverTab[183796]++
																		cd.NoCachePresent = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:456
		// _ = "end of CoverTab[183796]"
	case "no-store":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:457
		_go_fuzz_dep_.CoverTab[183797]++
																		cd.NoStore = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:458
		// _ = "end of CoverTab[183797]"
	case "no-transform":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:459
		_go_fuzz_dep_.CoverTab[183798]++
																		cd.NoTransform = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:460
		// _ = "end of CoverTab[183798]"
	case "public":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:461
		_go_fuzz_dep_.CoverTab[183799]++
																		cd.Public = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:462
		// _ = "end of CoverTab[183799]"
	case "private":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:463
		_go_fuzz_dep_.CoverTab[183800]++
																		cd.PrivatePresent = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:464
		// _ = "end of CoverTab[183800]"
	case "proxy-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:465
		_go_fuzz_dep_.CoverTab[183801]++
																		cd.ProxyRevalidate = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:466
		// _ = "end of CoverTab[183801]"
	case "max-age":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:467
		_go_fuzz_dep_.CoverTab[183802]++
																		err = ErrMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:468
		// _ = "end of CoverTab[183802]"
	case "s-maxage":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:469
		_go_fuzz_dep_.CoverTab[183803]++
																		err = ErrSMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:470
		// _ = "end of CoverTab[183803]"

	case "immutable":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:472
		_go_fuzz_dep_.CoverTab[183804]++
																		cd.Immutable = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:473
		// _ = "end of CoverTab[183804]"
	case "stale-if-error":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:474
		_go_fuzz_dep_.CoverTab[183805]++
																		err = ErrMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:475
		// _ = "end of CoverTab[183805]"
	case "stale-while-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:476
		_go_fuzz_dep_.CoverTab[183806]++
																		err = ErrMaxAgeDeltaSeconds
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:477
		// _ = "end of CoverTab[183806]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:478
		_go_fuzz_dep_.CoverTab[183807]++
																		cd.Extensions = append(cd.Extensions, token)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:479
		// _ = "end of CoverTab[183807]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:480
	// _ = "end of CoverTab[183793]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:480
	_go_fuzz_dep_.CoverTab[183794]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:481
	// _ = "end of CoverTab[183794]"
}

func hasFieldNames(token string) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:484
	_go_fuzz_dep_.CoverTab[183808]++
																	switch token {
	case "no-cache":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:486
		_go_fuzz_dep_.CoverTab[183810]++
																		return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:487
		// _ = "end of CoverTab[183810]"
	case "private":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:488
		_go_fuzz_dep_.CoverTab[183811]++
																		return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:489
		// _ = "end of CoverTab[183811]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:489
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:489
		_go_fuzz_dep_.CoverTab[183812]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:489
		// _ = "end of CoverTab[183812]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:490
	// _ = "end of CoverTab[183808]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:490
	_go_fuzz_dep_.CoverTab[183809]++
																	return false
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:491
	// _ = "end of CoverTab[183809]"
}

func (cd *ResponseCacheDirectives) addPair(token string, v string) error {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:494
	_go_fuzz_dep_.CoverTab[183813]++
																	var err error = nil

																	switch token {
	case "must-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:498
		_go_fuzz_dep_.CoverTab[183815]++
																		err = ErrMustRevalidateNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:499
		// _ = "end of CoverTab[183815]"
	case "no-cache":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:500
		_go_fuzz_dep_.CoverTab[183816]++
																		cd.NoCachePresent = true
																		tokens := strings.Split(v, ",")
																		if cd.NoCache == nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:503
			_go_fuzz_dep_.CoverTab[183830]++
																			cd.NoCache = make(FieldNames)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:504
			// _ = "end of CoverTab[183830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:505
			_go_fuzz_dep_.CoverTab[183831]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:505
			// _ = "end of CoverTab[183831]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:505
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:505
		// _ = "end of CoverTab[183816]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:505
		_go_fuzz_dep_.CoverTab[183817]++
																		for _, t := range tokens {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:506
			_go_fuzz_dep_.CoverTab[183832]++
																			k := http.CanonicalHeaderKey(textproto.TrimString(t))
																			cd.NoCache[k] = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:508
			// _ = "end of CoverTab[183832]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:509
		// _ = "end of CoverTab[183817]"
	case "no-store":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:510
		_go_fuzz_dep_.CoverTab[183818]++
																		err = ErrNoStoreNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:511
		// _ = "end of CoverTab[183818]"
	case "no-transform":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:512
		_go_fuzz_dep_.CoverTab[183819]++
																		err = ErrNoTransformNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:513
		// _ = "end of CoverTab[183819]"
	case "public":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:514
		_go_fuzz_dep_.CoverTab[183820]++
																		err = ErrPublicNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:515
		// _ = "end of CoverTab[183820]"
	case "private":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:516
		_go_fuzz_dep_.CoverTab[183821]++
																		cd.PrivatePresent = true
																		tokens := strings.Split(v, ",")
																		if cd.Private == nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:519
			_go_fuzz_dep_.CoverTab[183833]++
																			cd.Private = make(FieldNames)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:520
			// _ = "end of CoverTab[183833]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:521
			_go_fuzz_dep_.CoverTab[183834]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:521
			// _ = "end of CoverTab[183834]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:521
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:521
		// _ = "end of CoverTab[183821]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:521
		_go_fuzz_dep_.CoverTab[183822]++
																		for _, t := range tokens {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:522
			_go_fuzz_dep_.CoverTab[183835]++
																			k := http.CanonicalHeaderKey(textproto.TrimString(t))
																			cd.Private[k] = true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:524
			// _ = "end of CoverTab[183835]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:525
		// _ = "end of CoverTab[183822]"
	case "proxy-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:526
		_go_fuzz_dep_.CoverTab[183823]++
																		err = ErrProxyRevalidateNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:527
		// _ = "end of CoverTab[183823]"
	case "max-age":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:528
		_go_fuzz_dep_.CoverTab[183824]++
																		cd.MaxAge, err = parseDeltaSeconds(v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:529
		// _ = "end of CoverTab[183824]"
	case "s-maxage":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:530
		_go_fuzz_dep_.CoverTab[183825]++
																		cd.SMaxAge, err = parseDeltaSeconds(v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:531
		// _ = "end of CoverTab[183825]"

	case "immutable":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:533
		_go_fuzz_dep_.CoverTab[183826]++
																		err = ErrImmutableNoArgs
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:534
		// _ = "end of CoverTab[183826]"
	case "stale-if-error":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:535
		_go_fuzz_dep_.CoverTab[183827]++
																		cd.StaleIfError, err = parseDeltaSeconds(v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:536
		// _ = "end of CoverTab[183827]"
	case "stale-while-revalidate":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:537
		_go_fuzz_dep_.CoverTab[183828]++
																		cd.StaleWhileRevalidate, err = parseDeltaSeconds(v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:538
		// _ = "end of CoverTab[183828]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:539
		_go_fuzz_dep_.CoverTab[183829]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:542
		cd.Extensions = append(cd.Extensions, token+"="+v)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:542
		// _ = "end of CoverTab[183829]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:543
	// _ = "end of CoverTab[183813]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:543
	_go_fuzz_dep_.CoverTab[183814]++

																	return err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:545
	// _ = "end of CoverTab[183814]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:546
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/directive.go:546
var _ = _go_fuzz_dep_.CoverTab
