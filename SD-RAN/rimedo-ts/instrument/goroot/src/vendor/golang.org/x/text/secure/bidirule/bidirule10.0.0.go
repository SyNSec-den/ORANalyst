// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.10
// +build go1.10

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
package bidirule

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:8
)

func (t *Transformer) isFinal() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:10
	_go_fuzz_dep_.CoverTab[32929]++
											return t.state == ruleLTRFinal || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		_go_fuzz_dep_.CoverTab[32930]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		return t.state == ruleRTLFinal
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		// _ = "end of CoverTab[32930]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		_go_fuzz_dep_.CoverTab[32931]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		return t.state == ruleInitial
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
		// _ = "end of CoverTab[32931]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:11
	// _ = "end of CoverTab[32929]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:12
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go:12
var _ = _go_fuzz_dep_.CoverTab
