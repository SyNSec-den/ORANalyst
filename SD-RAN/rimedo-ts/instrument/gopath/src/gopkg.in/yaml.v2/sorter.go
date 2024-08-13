//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:1
)

import (
	"reflect"
	"unicode"
)

type keyList []reflect.Value

func (l keyList) Len() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:10
	_go_fuzz_dep_.CoverTab[127763]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:10
	return len(l)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:10
	// _ = "end of CoverTab[127763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:10
}
func (l keyList) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:11
	_go_fuzz_dep_.CoverTab[127764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:11
	l[i], l[j] = l[j], l[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:11
	// _ = "end of CoverTab[127764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:11
}
func (l keyList) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:12
	_go_fuzz_dep_.CoverTab[127765]++
									a := l[i]
									b := l[j]
									ak := a.Kind()
									bk := b.Kind()
									for (ak == reflect.Interface || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		_go_fuzz_dep_.CoverTab[127771]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		return ak == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		// _ = "end of CoverTab[127771]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		_go_fuzz_dep_.CoverTab[127772]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		return !a.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		// _ = "end of CoverTab[127772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:17
		_go_fuzz_dep_.CoverTab[127773]++
										a = a.Elem()
										ak = a.Kind()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:19
		// _ = "end of CoverTab[127773]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:20
	// _ = "end of CoverTab[127765]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:20
	_go_fuzz_dep_.CoverTab[127766]++
									for (bk == reflect.Interface || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		_go_fuzz_dep_.CoverTab[127774]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		return bk == reflect.Ptr
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		// _ = "end of CoverTab[127774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		_go_fuzz_dep_.CoverTab[127775]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		return !b.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		// _ = "end of CoverTab[127775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:21
		_go_fuzz_dep_.CoverTab[127776]++
										b = b.Elem()
										bk = b.Kind()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:23
		// _ = "end of CoverTab[127776]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:24
	// _ = "end of CoverTab[127766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:24
	_go_fuzz_dep_.CoverTab[127767]++
									af, aok := keyFloat(a)
									bf, bok := keyFloat(b)
									if aok && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:27
		_go_fuzz_dep_.CoverTab[127777]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:27
		return bok
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:27
		// _ = "end of CoverTab[127777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:27
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:27
		_go_fuzz_dep_.CoverTab[127778]++
										if af != bf {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:28
			_go_fuzz_dep_.CoverTab[127781]++
											return af < bf
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:29
			// _ = "end of CoverTab[127781]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:30
			_go_fuzz_dep_.CoverTab[127782]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:30
			// _ = "end of CoverTab[127782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:30
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:30
		// _ = "end of CoverTab[127778]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:30
		_go_fuzz_dep_.CoverTab[127779]++
										if ak != bk {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:31
			_go_fuzz_dep_.CoverTab[127783]++
											return ak < bk
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:32
			// _ = "end of CoverTab[127783]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:33
			_go_fuzz_dep_.CoverTab[127784]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:33
			// _ = "end of CoverTab[127784]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:33
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:33
		// _ = "end of CoverTab[127779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:33
		_go_fuzz_dep_.CoverTab[127780]++
										return numLess(a, b)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:34
		// _ = "end of CoverTab[127780]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:35
		_go_fuzz_dep_.CoverTab[127785]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:35
		// _ = "end of CoverTab[127785]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:35
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:35
	// _ = "end of CoverTab[127767]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:35
	_go_fuzz_dep_.CoverTab[127768]++
									if ak != reflect.String || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:36
		_go_fuzz_dep_.CoverTab[127786]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:36
		return bk != reflect.String
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:36
		// _ = "end of CoverTab[127786]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:36
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:36
		_go_fuzz_dep_.CoverTab[127787]++
										return ak < bk
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:37
		// _ = "end of CoverTab[127787]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:38
		_go_fuzz_dep_.CoverTab[127788]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:38
		// _ = "end of CoverTab[127788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:38
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:38
	// _ = "end of CoverTab[127768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:38
	_go_fuzz_dep_.CoverTab[127769]++
									ar, br := []rune(a.String()), []rune(b.String())
									for i := 0; i < len(ar) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:40
		_go_fuzz_dep_.CoverTab[127789]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:40
		return i < len(br)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:40
		// _ = "end of CoverTab[127789]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:40
	}(); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:40
		_go_fuzz_dep_.CoverTab[127790]++
										if ar[i] == br[i] {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:41
			_go_fuzz_dep_.CoverTab[127799]++
											continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:42
			// _ = "end of CoverTab[127799]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:43
			_go_fuzz_dep_.CoverTab[127800]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:43
			// _ = "end of CoverTab[127800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:43
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:43
		// _ = "end of CoverTab[127790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:43
		_go_fuzz_dep_.CoverTab[127791]++
										al := unicode.IsLetter(ar[i])
										bl := unicode.IsLetter(br[i])
										if al && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:46
			_go_fuzz_dep_.CoverTab[127801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:46
			return bl
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:46
			// _ = "end of CoverTab[127801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:46
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:46
			_go_fuzz_dep_.CoverTab[127802]++
											return ar[i] < br[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:47
			// _ = "end of CoverTab[127802]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:48
			_go_fuzz_dep_.CoverTab[127803]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:48
			// _ = "end of CoverTab[127803]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:48
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:48
		// _ = "end of CoverTab[127791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:48
		_go_fuzz_dep_.CoverTab[127792]++
										if al || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:49
			_go_fuzz_dep_.CoverTab[127804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:49
			return bl
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:49
			// _ = "end of CoverTab[127804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:49
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:49
			_go_fuzz_dep_.CoverTab[127805]++
											return bl
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:50
			// _ = "end of CoverTab[127805]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:51
			_go_fuzz_dep_.CoverTab[127806]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:51
			// _ = "end of CoverTab[127806]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:51
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:51
		// _ = "end of CoverTab[127792]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:51
		_go_fuzz_dep_.CoverTab[127793]++
										var ai, bi int
										var an, bn int64
										if ar[i] == '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:54
			_go_fuzz_dep_.CoverTab[127807]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:54
			return br[i] == '0'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:54
			// _ = "end of CoverTab[127807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:54
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:54
			_go_fuzz_dep_.CoverTab[127808]++
											for j := i - 1; j >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:55
				_go_fuzz_dep_.CoverTab[127809]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:55
				return unicode.IsDigit(ar[j])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:55
				// _ = "end of CoverTab[127809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:55
			}(); j-- {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:55
				_go_fuzz_dep_.CoverTab[127810]++
												if ar[j] != '0' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:56
					_go_fuzz_dep_.CoverTab[127811]++
													an = 1
													bn = 1
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:59
					// _ = "end of CoverTab[127811]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:60
					_go_fuzz_dep_.CoverTab[127812]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:60
					// _ = "end of CoverTab[127812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:60
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:60
				// _ = "end of CoverTab[127810]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:61
			// _ = "end of CoverTab[127808]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:62
			_go_fuzz_dep_.CoverTab[127813]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:62
			// _ = "end of CoverTab[127813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:62
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:62
		// _ = "end of CoverTab[127793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:62
		_go_fuzz_dep_.CoverTab[127794]++
										for ai = i; ai < len(ar) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:63
			_go_fuzz_dep_.CoverTab[127814]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:63
			return unicode.IsDigit(ar[ai])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:63
			// _ = "end of CoverTab[127814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:63
		}(); ai++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:63
			_go_fuzz_dep_.CoverTab[127815]++
											an = an*10 + int64(ar[ai]-'0')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:64
			// _ = "end of CoverTab[127815]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:65
		// _ = "end of CoverTab[127794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:65
		_go_fuzz_dep_.CoverTab[127795]++
										for bi = i; bi < len(br) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:66
			_go_fuzz_dep_.CoverTab[127816]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:66
			return unicode.IsDigit(br[bi])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:66
			// _ = "end of CoverTab[127816]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:66
		}(); bi++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:66
			_go_fuzz_dep_.CoverTab[127817]++
											bn = bn*10 + int64(br[bi]-'0')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:67
			// _ = "end of CoverTab[127817]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:68
		// _ = "end of CoverTab[127795]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:68
		_go_fuzz_dep_.CoverTab[127796]++
										if an != bn {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:69
			_go_fuzz_dep_.CoverTab[127818]++
											return an < bn
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:70
			// _ = "end of CoverTab[127818]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:71
			_go_fuzz_dep_.CoverTab[127819]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:71
			// _ = "end of CoverTab[127819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:71
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:71
		// _ = "end of CoverTab[127796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:71
		_go_fuzz_dep_.CoverTab[127797]++
										if ai != bi {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:72
			_go_fuzz_dep_.CoverTab[127820]++
											return ai < bi
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:73
			// _ = "end of CoverTab[127820]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:74
			_go_fuzz_dep_.CoverTab[127821]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:74
			// _ = "end of CoverTab[127821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:74
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:74
		// _ = "end of CoverTab[127797]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:74
		_go_fuzz_dep_.CoverTab[127798]++
										return ar[i] < br[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:75
		// _ = "end of CoverTab[127798]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:76
	// _ = "end of CoverTab[127769]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:76
	_go_fuzz_dep_.CoverTab[127770]++
									return len(ar) < len(br)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:77
	// _ = "end of CoverTab[127770]"
}

// keyFloat returns a float value for v if it is a number/bool
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:80
// and whether it is a number/bool or not.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:82
func keyFloat(v reflect.Value) (f float64, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:82
	_go_fuzz_dep_.CoverTab[127822]++
									switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:84
		_go_fuzz_dep_.CoverTab[127824]++
										return float64(v.Int()), true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:85
		// _ = "end of CoverTab[127824]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:86
		_go_fuzz_dep_.CoverTab[127825]++
										return v.Float(), true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:87
		// _ = "end of CoverTab[127825]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:88
		_go_fuzz_dep_.CoverTab[127826]++
										return float64(v.Uint()), true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:89
		// _ = "end of CoverTab[127826]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:90
		_go_fuzz_dep_.CoverTab[127827]++
										if v.Bool() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:91
			_go_fuzz_dep_.CoverTab[127830]++
											return 1, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:92
			// _ = "end of CoverTab[127830]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:93
			_go_fuzz_dep_.CoverTab[127831]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:93
			// _ = "end of CoverTab[127831]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:93
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:93
		// _ = "end of CoverTab[127827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:93
		_go_fuzz_dep_.CoverTab[127828]++
										return 0, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:94
		// _ = "end of CoverTab[127828]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:94
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:94
		_go_fuzz_dep_.CoverTab[127829]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:94
		// _ = "end of CoverTab[127829]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:95
	// _ = "end of CoverTab[127822]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:95
	_go_fuzz_dep_.CoverTab[127823]++
									return 0, false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:96
	// _ = "end of CoverTab[127823]"
}

// numLess returns whether a < b.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:99
// a and b must necessarily have the same kind.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:101
func numLess(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:101
	_go_fuzz_dep_.CoverTab[127832]++
										switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:103
		_go_fuzz_dep_.CoverTab[127834]++
											return a.Int() < b.Int()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:104
		// _ = "end of CoverTab[127834]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:105
		_go_fuzz_dep_.CoverTab[127835]++
											return a.Float() < b.Float()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:106
		// _ = "end of CoverTab[127835]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:107
		_go_fuzz_dep_.CoverTab[127836]++
											return a.Uint() < b.Uint()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:108
		// _ = "end of CoverTab[127836]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:109
		_go_fuzz_dep_.CoverTab[127837]++
											return !a.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
			_go_fuzz_dep_.CoverTab[127839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
			return b.Bool()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
			// _ = "end of CoverTab[127839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
		// _ = "end of CoverTab[127837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
		_go_fuzz_dep_.CoverTab[127838]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:110
		// _ = "end of CoverTab[127838]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:111
	// _ = "end of CoverTab[127832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:111
	_go_fuzz_dep_.CoverTab[127833]++
										panic("not a number")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:112
	// _ = "end of CoverTab[127833]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/sorter.go:113
var _ = _go_fuzz_dep_.CoverTab
