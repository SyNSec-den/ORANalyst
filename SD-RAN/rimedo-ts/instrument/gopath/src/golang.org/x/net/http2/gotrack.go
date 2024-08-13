// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Defensive debug-only utility to track that functions run on the
// goroutine that they're supposed to.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:8
)

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var DebugGoroutines = os.Getenv("DEBUG_HTTP2_GOROUTINES") == "1"

type goroutineLock uint64

func newGoroutineLock() goroutineLock {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:24
	_go_fuzz_dep_.CoverTab[72980]++
										if !DebugGoroutines {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:25
		_go_fuzz_dep_.CoverTab[72982]++
											return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:26
		// _ = "end of CoverTab[72982]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:27
		_go_fuzz_dep_.CoverTab[72983]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:27
		// _ = "end of CoverTab[72983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:27
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:27
	// _ = "end of CoverTab[72980]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:27
	_go_fuzz_dep_.CoverTab[72981]++
										return goroutineLock(curGoroutineID())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:28
	// _ = "end of CoverTab[72981]"
}

func (g goroutineLock) check() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:31
	_go_fuzz_dep_.CoverTab[72984]++
										if !DebugGoroutines {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:32
		_go_fuzz_dep_.CoverTab[72986]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:33
		// _ = "end of CoverTab[72986]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:34
		_go_fuzz_dep_.CoverTab[72987]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:34
		// _ = "end of CoverTab[72987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:34
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:34
	// _ = "end of CoverTab[72984]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:34
	_go_fuzz_dep_.CoverTab[72985]++
										if curGoroutineID() != uint64(g) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:35
		_go_fuzz_dep_.CoverTab[72988]++
											panic("running on the wrong goroutine")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:36
		// _ = "end of CoverTab[72988]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:37
		_go_fuzz_dep_.CoverTab[72989]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:37
		// _ = "end of CoverTab[72989]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:37
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:37
	// _ = "end of CoverTab[72985]"
}

func (g goroutineLock) checkNotOn() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:40
	_go_fuzz_dep_.CoverTab[72990]++
										if !DebugGoroutines {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:41
		_go_fuzz_dep_.CoverTab[72992]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:42
		// _ = "end of CoverTab[72992]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:43
		_go_fuzz_dep_.CoverTab[72993]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:43
		// _ = "end of CoverTab[72993]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:43
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:43
	// _ = "end of CoverTab[72990]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:43
	_go_fuzz_dep_.CoverTab[72991]++
										if curGoroutineID() == uint64(g) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:44
		_go_fuzz_dep_.CoverTab[72994]++
											panic("running on the wrong goroutine")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:45
		// _ = "end of CoverTab[72994]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:46
		_go_fuzz_dep_.CoverTab[72995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:46
		// _ = "end of CoverTab[72995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:46
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:46
	// _ = "end of CoverTab[72991]"
}

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:51
	_go_fuzz_dep_.CoverTab[72996]++
										bp := littleBuf.Get().(*[]byte)
										defer littleBuf.Put(bp)
										b := *bp
										b = b[:runtime.Stack(b, false)]

										b = bytes.TrimPrefix(b, goroutineSpace)
										i := bytes.IndexByte(b, ' ')
										if i < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:59
		_go_fuzz_dep_.CoverTab[72999]++
											panic(fmt.Sprintf("No space found in %q", b))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:60
		// _ = "end of CoverTab[72999]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:61
		_go_fuzz_dep_.CoverTab[73000]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:61
		// _ = "end of CoverTab[73000]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:61
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:61
	// _ = "end of CoverTab[72996]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:61
	_go_fuzz_dep_.CoverTab[72997]++
										b = b[:i]
										n, err := parseUintBytes(b, 10, 64)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:64
		_go_fuzz_dep_.CoverTab[73001]++
											panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:65
		// _ = "end of CoverTab[73001]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:66
		_go_fuzz_dep_.CoverTab[73002]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:66
		// _ = "end of CoverTab[73002]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:66
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:66
	// _ = "end of CoverTab[72997]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:66
	_go_fuzz_dep_.CoverTab[72998]++
										return n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:67
	// _ = "end of CoverTab[72998]"
}

var littleBuf = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:71
		_go_fuzz_dep_.CoverTab[73003]++
											buf := make([]byte, 64)
											return &buf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:73
		// _ = "end of CoverTab[73003]"
	},
}

// parseUintBytes is like strconv.ParseUint, but using a []byte.
func parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:78
	_go_fuzz_dep_.CoverTab[73004]++
										var cutoff, maxVal uint64

										if bitSize == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:81
		_go_fuzz_dep_.CoverTab[73008]++
											bitSize = int(strconv.IntSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:82
		// _ = "end of CoverTab[73008]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:83
		_go_fuzz_dep_.CoverTab[73009]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:83
		// _ = "end of CoverTab[73009]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:83
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:83
	// _ = "end of CoverTab[73004]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:83
	_go_fuzz_dep_.CoverTab[73005]++

										s0 := s
										switch {
	case len(s) < 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:87
		_go_fuzz_dep_.CoverTab[73010]++
											err = strconv.ErrSyntax
											goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:89
		// _ = "end of CoverTab[73010]"

	case 2 <= base && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
		_go_fuzz_dep_.CoverTab[73014]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
		return base <= 36
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
		// _ = "end of CoverTab[73014]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
	}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
		_go_fuzz_dep_.CoverTab[73011]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:91
		// _ = "end of CoverTab[73011]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:94
	case base == 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:94
		_go_fuzz_dep_.CoverTab[73012]++

											switch {
		case s[0] == '0' && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			_go_fuzz_dep_.CoverTab[73018]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			return len(s) > 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			// _ = "end of CoverTab[73018]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			_go_fuzz_dep_.CoverTab[73019]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			return (s[1] == 'x' || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
				_go_fuzz_dep_.CoverTab[73020]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
				return s[1] == 'X'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
				// _ = "end of CoverTab[73020]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			// _ = "end of CoverTab[73019]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
		}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:97
			_go_fuzz_dep_.CoverTab[73015]++
												base = 16
												s = s[2:]
												if len(s) < 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:100
				_go_fuzz_dep_.CoverTab[73021]++
														err = strconv.ErrSyntax
														goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:102
				// _ = "end of CoverTab[73021]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:103
				_go_fuzz_dep_.CoverTab[73022]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:103
				// _ = "end of CoverTab[73022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:103
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:103
			// _ = "end of CoverTab[73015]"
		case s[0] == '0':
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:104
			_go_fuzz_dep_.CoverTab[73016]++
													base = 8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:105
			// _ = "end of CoverTab[73016]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:106
			_go_fuzz_dep_.CoverTab[73017]++
													base = 10
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:107
			// _ = "end of CoverTab[73017]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:108
		// _ = "end of CoverTab[73012]"

	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:110
		_go_fuzz_dep_.CoverTab[73013]++
												err = errors.New("invalid base " + strconv.Itoa(base))
												goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:112
		// _ = "end of CoverTab[73013]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:113
	// _ = "end of CoverTab[73005]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:113
	_go_fuzz_dep_.CoverTab[73006]++

											n = 0
											cutoff = cutoff64(base)
											maxVal = 1<<uint(bitSize) - 1

											for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:119
		_go_fuzz_dep_.CoverTab[73023]++
												var v byte
												d := s[i]
												switch {
		case '0' <= d && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:123
			_go_fuzz_dep_.CoverTab[73032]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:123
			return d <= '9'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:123
			// _ = "end of CoverTab[73032]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:123
		}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:123
			_go_fuzz_dep_.CoverTab[73028]++
													v = d - '0'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:124
			// _ = "end of CoverTab[73028]"
		case 'a' <= d && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:125
			_go_fuzz_dep_.CoverTab[73033]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:125
			return d <= 'z'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:125
			// _ = "end of CoverTab[73033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:125
		}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:125
			_go_fuzz_dep_.CoverTab[73029]++
													v = d - 'a' + 10
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:126
			// _ = "end of CoverTab[73029]"
		case 'A' <= d && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:127
			_go_fuzz_dep_.CoverTab[73034]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:127
			return d <= 'Z'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:127
			// _ = "end of CoverTab[73034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:127
		}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:127
			_go_fuzz_dep_.CoverTab[73030]++
													v = d - 'A' + 10
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:128
			// _ = "end of CoverTab[73030]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:129
			_go_fuzz_dep_.CoverTab[73031]++
													n = 0
													err = strconv.ErrSyntax
													goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:132
			// _ = "end of CoverTab[73031]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:133
		// _ = "end of CoverTab[73023]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:133
		_go_fuzz_dep_.CoverTab[73024]++
												if int(v) >= base {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:134
			_go_fuzz_dep_.CoverTab[73035]++
													n = 0
													err = strconv.ErrSyntax
													goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:137
			// _ = "end of CoverTab[73035]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:138
			_go_fuzz_dep_.CoverTab[73036]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:138
			// _ = "end of CoverTab[73036]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:138
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:138
		// _ = "end of CoverTab[73024]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:138
		_go_fuzz_dep_.CoverTab[73025]++

												if n >= cutoff {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:140
			_go_fuzz_dep_.CoverTab[73037]++

													n = 1<<64 - 1
													err = strconv.ErrRange
													goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:144
			// _ = "end of CoverTab[73037]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:145
			_go_fuzz_dep_.CoverTab[73038]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:145
			// _ = "end of CoverTab[73038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:145
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:145
		// _ = "end of CoverTab[73025]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:145
		_go_fuzz_dep_.CoverTab[73026]++
												n *= uint64(base)

												n1 := n + uint64(v)
												if n1 < n || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:149
			_go_fuzz_dep_.CoverTab[73039]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:149
			return n1 > maxVal
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:149
			// _ = "end of CoverTab[73039]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:149
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:149
			_go_fuzz_dep_.CoverTab[73040]++

													n = 1<<64 - 1
													err = strconv.ErrRange
													goto Error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:153
			// _ = "end of CoverTab[73040]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:154
			_go_fuzz_dep_.CoverTab[73041]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:154
			// _ = "end of CoverTab[73041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:154
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:154
		// _ = "end of CoverTab[73026]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:154
		_go_fuzz_dep_.CoverTab[73027]++
												n = n1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:155
		// _ = "end of CoverTab[73027]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:156
	// _ = "end of CoverTab[73006]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:156
	_go_fuzz_dep_.CoverTab[73007]++

											return n, nil

Error:
											return n, &strconv.NumError{Func: "ParseUint", Num: string(s0), Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:161
	// _ = "end of CoverTab[73007]"
}

// Return the first number n such that n*base >= 1<<64.
func cutoff64(base int) uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:165
	_go_fuzz_dep_.CoverTab[73042]++
											if base < 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:166
		_go_fuzz_dep_.CoverTab[73044]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:167
		// _ = "end of CoverTab[73044]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:168
		_go_fuzz_dep_.CoverTab[73045]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:168
		// _ = "end of CoverTab[73045]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:168
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:168
	// _ = "end of CoverTab[73042]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:168
	_go_fuzz_dep_.CoverTab[73043]++
											return (1<<64-1)/uint64(base) + 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:169
	// _ = "end of CoverTab[73043]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:170
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/gotrack.go:170
var _ = _go_fuzz_dep_.CoverTab
