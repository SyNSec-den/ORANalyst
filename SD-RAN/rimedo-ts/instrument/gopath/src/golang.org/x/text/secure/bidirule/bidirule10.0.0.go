// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.10
// +build go1.10

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
package bidirule

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:8
)

func (t *Transformer) isFinal() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:10
	_go_fuzz_dep_.CoverTab[70210]++
													return t.state == ruleLTRFinal || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		_go_fuzz_dep_.CoverTab[70211]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		return t.state == ruleRTLFinal
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		// _ = "end of CoverTab[70211]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		_go_fuzz_dep_.CoverTab[70212]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		return t.state == ruleInitial
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
		// _ = "end of CoverTab[70212]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:11
	// _ = "end of CoverTab[70210]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:12
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule10.0.0.go:12
var _ = _go_fuzz_dep_.CoverTab
