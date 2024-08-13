//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:1
)

import "runtime/debug"

var v string

func version() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:7
	_go_fuzz_dep_.CoverTab[107062]++
											if v == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:8
		_go_fuzz_dep_.CoverTab[107064]++
												bi, ok := debug.ReadBuildInfo()
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:10
			_go_fuzz_dep_.CoverTab[107065]++
													v = bi.Main.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:11
			// _ = "end of CoverTab[107065]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:12
			_go_fuzz_dep_.CoverTab[107066]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:16
			v = "dev"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:16
			// _ = "end of CoverTab[107066]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:17
		// _ = "end of CoverTab[107064]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:18
		_go_fuzz_dep_.CoverTab[107067]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:18
		// _ = "end of CoverTab[107067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:18
	// _ = "end of CoverTab[107062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:18
	_go_fuzz_dep_.CoverTab[107063]++
											return v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:19
	// _ = "end of CoverTab[107063]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/version.go:20
var _ = _go_fuzz_dep_.CoverTab
