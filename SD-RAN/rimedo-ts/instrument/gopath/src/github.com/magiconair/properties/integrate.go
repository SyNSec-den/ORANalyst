// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:5
)

import "flag"

// MustFlag sets flags that are skipped by dst.Parse when p contains
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
// the respective key for flag.Flag.Name.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
// It's use is recommended with command line arguments as in:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
//	flag.Parse()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:9
//	p.MustFlag(flag.CommandLine)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:15
func (p *Properties) MustFlag(dst *flag.FlagSet) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:15
	_go_fuzz_dep_.CoverTab[115632]++
												m := make(map[string]*flag.Flag)
												dst.VisitAll(func(f *flag.Flag) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:17
		_go_fuzz_dep_.CoverTab[115635]++
													m[f.Name] = f
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:18
		// _ = "end of CoverTab[115635]"
	})
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:19
	// _ = "end of CoverTab[115632]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:19
	_go_fuzz_dep_.CoverTab[115633]++
												dst.Visit(func(f *flag.Flag) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:20
		_go_fuzz_dep_.CoverTab[115636]++
													delete(m, f.Name)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:21
		// _ = "end of CoverTab[115636]"
	})
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:22
	// _ = "end of CoverTab[115633]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:22
	_go_fuzz_dep_.CoverTab[115634]++

												for name, f := range m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:24
		_go_fuzz_dep_.CoverTab[115637]++
													v, ok := p.Get(name)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:26
			_go_fuzz_dep_.CoverTab[115639]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:27
			// _ = "end of CoverTab[115639]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:28
			_go_fuzz_dep_.CoverTab[115640]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:28
			// _ = "end of CoverTab[115640]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:28
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:28
		// _ = "end of CoverTab[115637]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:28
		_go_fuzz_dep_.CoverTab[115638]++

													if err := f.Value.Set(v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:30
			_go_fuzz_dep_.CoverTab[115641]++
														ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:31
			// _ = "end of CoverTab[115641]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:32
			_go_fuzz_dep_.CoverTab[115642]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:32
			// _ = "end of CoverTab[115642]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:32
		// _ = "end of CoverTab[115638]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:33
	// _ = "end of CoverTab[115634]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/integrate.go:34
var _ = _go_fuzz_dep_.CoverTab
