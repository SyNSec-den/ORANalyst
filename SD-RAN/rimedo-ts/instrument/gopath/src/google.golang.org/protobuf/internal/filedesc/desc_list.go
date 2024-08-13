// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:5
)

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"google.golang.org/protobuf/internal/genid"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/descfmt"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FileImports []protoreflect.FileImport

func (p *FileImports) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:24
	_go_fuzz_dep_.CoverTab[53088]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:24
	return len(*p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:24
	// _ = "end of CoverTab[53088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:24
}
func (p *FileImports) Get(i int) protoreflect.FileImport {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:25
	_go_fuzz_dep_.CoverTab[53089]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:25
	return (*p)[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:25
	// _ = "end of CoverTab[53089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:25
}
func (p *FileImports) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:26
	_go_fuzz_dep_.CoverTab[53090]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:26
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:26
	// _ = "end of CoverTab[53090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:26
}
func (p *FileImports) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:27
	_go_fuzz_dep_.CoverTab[53091]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:27
	// _ = "end of CoverTab[53091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:27
}

type Names struct {
	List	[]protoreflect.Name
	once	sync.Once
	has	map[protoreflect.Name]int	// protected by once
}

func (p *Names) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:35
	_go_fuzz_dep_.CoverTab[53092]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:35
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:35
	// _ = "end of CoverTab[53092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:35
}
func (p *Names) Get(i int) protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:36
	_go_fuzz_dep_.CoverTab[53093]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:36
	return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:36
	// _ = "end of CoverTab[53093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:36
}
func (p *Names) Has(s protoreflect.Name) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:37
	_go_fuzz_dep_.CoverTab[53094]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:37
	return p.lazyInit().has[s] > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:37
	// _ = "end of CoverTab[53094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:37
}
func (p *Names) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:38
	_go_fuzz_dep_.CoverTab[53095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:38
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:38
	// _ = "end of CoverTab[53095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:38
}
func (p *Names) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:39
	_go_fuzz_dep_.CoverTab[53096]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:39
	// _ = "end of CoverTab[53096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:39
}
func (p *Names) lazyInit() *Names {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:40
	_go_fuzz_dep_.CoverTab[53097]++
													p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:41
		_go_fuzz_dep_.CoverTab[53099]++
														if len(p.List) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:42
			_go_fuzz_dep_.CoverTab[53100]++
															p.has = make(map[protoreflect.Name]int, len(p.List))
															for _, s := range p.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:44
				_go_fuzz_dep_.CoverTab[53101]++
																p.has[s] = p.has[s] + 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:45
				// _ = "end of CoverTab[53101]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:46
			// _ = "end of CoverTab[53100]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:47
			_go_fuzz_dep_.CoverTab[53102]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:47
			// _ = "end of CoverTab[53102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:47
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:47
		// _ = "end of CoverTab[53099]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:48
	// _ = "end of CoverTab[53097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:48
	_go_fuzz_dep_.CoverTab[53098]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:49
	// _ = "end of CoverTab[53098]"
}

// CheckValid reports any errors with the set of names with an error message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:52
// that completes the sentence: "ranges is invalid because it has ..."
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:54
func (p *Names) CheckValid() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:54
	_go_fuzz_dep_.CoverTab[53103]++
													for s, n := range p.lazyInit().has {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:55
		_go_fuzz_dep_.CoverTab[53105]++
														switch {
		case n > 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:57
			_go_fuzz_dep_.CoverTab[53106]++
															return errors.New("duplicate name: %q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:58
			// _ = "end of CoverTab[53106]"
		case false && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:59
			_go_fuzz_dep_.CoverTab[53109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:59
			return !s.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:59
			// _ = "end of CoverTab[53109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:59
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:59
			_go_fuzz_dep_.CoverTab[53107]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:62
			return errors.New("invalid name: %q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:62
			// _ = "end of CoverTab[53107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:62
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:62
			_go_fuzz_dep_.CoverTab[53108]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:62
			// _ = "end of CoverTab[53108]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:63
		// _ = "end of CoverTab[53105]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:64
	// _ = "end of CoverTab[53103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:64
	_go_fuzz_dep_.CoverTab[53104]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:65
	// _ = "end of CoverTab[53104]"
}

type EnumRanges struct {
	List	[][2]protoreflect.EnumNumber	// start inclusive; end inclusive
	once	sync.Once
	sorted	[][2]protoreflect.EnumNumber	// protected by once
}

func (p *EnumRanges) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:74
	_go_fuzz_dep_.CoverTab[53110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:74
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:74
	// _ = "end of CoverTab[53110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:74
}
func (p *EnumRanges) Get(i int) [2]protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:75
	_go_fuzz_dep_.CoverTab[53111]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:75
	return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:75
	// _ = "end of CoverTab[53111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:75
}
func (p *EnumRanges) Has(n protoreflect.EnumNumber) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:76
	_go_fuzz_dep_.CoverTab[53112]++
													for ls := p.lazyInit().sorted; len(ls) > 0; {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:77
		_go_fuzz_dep_.CoverTab[53114]++
														i := len(ls) / 2
														switch r := enumRange(ls[i]); {
		case n < r.Start():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:80
			_go_fuzz_dep_.CoverTab[53115]++
															ls = ls[:i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:81
			// _ = "end of CoverTab[53115]"
		case n > r.End():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:82
			_go_fuzz_dep_.CoverTab[53116]++
															ls = ls[i+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:83
			// _ = "end of CoverTab[53116]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:84
			_go_fuzz_dep_.CoverTab[53117]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:85
			// _ = "end of CoverTab[53117]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:86
		// _ = "end of CoverTab[53114]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:87
	// _ = "end of CoverTab[53112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:87
	_go_fuzz_dep_.CoverTab[53113]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:88
	// _ = "end of CoverTab[53113]"
}
func (p *EnumRanges) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:90
	_go_fuzz_dep_.CoverTab[53118]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:90
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:90
	// _ = "end of CoverTab[53118]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:90
}
func (p *EnumRanges) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:91
	_go_fuzz_dep_.CoverTab[53119]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:91
	// _ = "end of CoverTab[53119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:91
}
func (p *EnumRanges) lazyInit() *EnumRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:92
	_go_fuzz_dep_.CoverTab[53120]++
													p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:93
		_go_fuzz_dep_.CoverTab[53122]++
														p.sorted = append(p.sorted, p.List...)
														sort.Slice(p.sorted, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:95
			_go_fuzz_dep_.CoverTab[53123]++
															return p.sorted[i][0] < p.sorted[j][0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:96
			// _ = "end of CoverTab[53123]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:97
		// _ = "end of CoverTab[53122]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:98
	// _ = "end of CoverTab[53120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:98
	_go_fuzz_dep_.CoverTab[53121]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:99
	// _ = "end of CoverTab[53121]"
}

// CheckValid reports any errors with the set of names with an error message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:102
// that completes the sentence: "ranges is invalid because it has ..."
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:104
func (p *EnumRanges) CheckValid() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:104
	_go_fuzz_dep_.CoverTab[53124]++
														var rp enumRange
														for i, r := range p.lazyInit().sorted {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:106
		_go_fuzz_dep_.CoverTab[53126]++
															r := enumRange(r)
															switch {
		case !(r.Start() <= r.End()):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:109
			_go_fuzz_dep_.CoverTab[53128]++
																return errors.New("invalid range: %v", r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:110
			// _ = "end of CoverTab[53128]"
		case !(rp.End() < r.Start()) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:111
			_go_fuzz_dep_.CoverTab[53131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:111
			return i > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:111
			// _ = "end of CoverTab[53131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:111
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:111
			_go_fuzz_dep_.CoverTab[53129]++
																return errors.New("overlapping ranges: %v with %v", rp, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:112
			// _ = "end of CoverTab[53129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:112
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:112
			_go_fuzz_dep_.CoverTab[53130]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:112
			// _ = "end of CoverTab[53130]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:113
		// _ = "end of CoverTab[53126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:113
		_go_fuzz_dep_.CoverTab[53127]++
															rp = r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:114
		// _ = "end of CoverTab[53127]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:115
	// _ = "end of CoverTab[53124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:115
	_go_fuzz_dep_.CoverTab[53125]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:116
	// _ = "end of CoverTab[53125]"
}

type enumRange [2]protoreflect.EnumNumber

func (r enumRange) Start() protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:121
	_go_fuzz_dep_.CoverTab[53132]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:121
	return r[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:121
	// _ = "end of CoverTab[53132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:121
}
func (r enumRange) End() protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:122
	_go_fuzz_dep_.CoverTab[53133]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:122
	return r[1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:122
	// _ = "end of CoverTab[53133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:122
}
func (r enumRange) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:123
	_go_fuzz_dep_.CoverTab[53134]++
														if r.Start() == r.End() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:124
		_go_fuzz_dep_.CoverTab[53136]++
															return fmt.Sprintf("%d", r.Start())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:125
		// _ = "end of CoverTab[53136]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:126
		_go_fuzz_dep_.CoverTab[53137]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:126
		// _ = "end of CoverTab[53137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:126
	// _ = "end of CoverTab[53134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:126
	_go_fuzz_dep_.CoverTab[53135]++
														return fmt.Sprintf("%d to %d", r.Start(), r.End())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:127
	// _ = "end of CoverTab[53135]"
}

type FieldRanges struct {
	List	[][2]protoreflect.FieldNumber	// start inclusive; end exclusive
	once	sync.Once
	sorted	[][2]protoreflect.FieldNumber	// protected by once
}

func (p *FieldRanges) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:136
	_go_fuzz_dep_.CoverTab[53138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:136
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:136
	// _ = "end of CoverTab[53138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:136
}
func (p *FieldRanges) Get(i int) [2]protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:137
	_go_fuzz_dep_.CoverTab[53139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:137
	return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:137
	// _ = "end of CoverTab[53139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:137
}
func (p *FieldRanges) Has(n protoreflect.FieldNumber) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:138
	_go_fuzz_dep_.CoverTab[53140]++
														for ls := p.lazyInit().sorted; len(ls) > 0; {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:139
		_go_fuzz_dep_.CoverTab[53142]++
															i := len(ls) / 2
															switch r := fieldRange(ls[i]); {
		case n < r.Start():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:142
			_go_fuzz_dep_.CoverTab[53143]++
																ls = ls[:i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:143
			// _ = "end of CoverTab[53143]"
		case n > r.End():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:144
			_go_fuzz_dep_.CoverTab[53144]++
																ls = ls[i+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:145
			// _ = "end of CoverTab[53144]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:146
			_go_fuzz_dep_.CoverTab[53145]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:147
			// _ = "end of CoverTab[53145]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:148
		// _ = "end of CoverTab[53142]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:149
	// _ = "end of CoverTab[53140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:149
	_go_fuzz_dep_.CoverTab[53141]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:150
	// _ = "end of CoverTab[53141]"
}
func (p *FieldRanges) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:152
	_go_fuzz_dep_.CoverTab[53146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:152
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:152
	// _ = "end of CoverTab[53146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:152
}
func (p *FieldRanges) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:153
	_go_fuzz_dep_.CoverTab[53147]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:153
	// _ = "end of CoverTab[53147]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:153
}
func (p *FieldRanges) lazyInit() *FieldRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:154
	_go_fuzz_dep_.CoverTab[53148]++
														p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:155
		_go_fuzz_dep_.CoverTab[53150]++
															p.sorted = append(p.sorted, p.List...)
															sort.Slice(p.sorted, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:157
			_go_fuzz_dep_.CoverTab[53151]++
																return p.sorted[i][0] < p.sorted[j][0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:158
			// _ = "end of CoverTab[53151]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:159
		// _ = "end of CoverTab[53150]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:160
	// _ = "end of CoverTab[53148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:160
	_go_fuzz_dep_.CoverTab[53149]++
														return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:161
	// _ = "end of CoverTab[53149]"
}

// CheckValid reports any errors with the set of ranges with an error message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:164
// that completes the sentence: "ranges is invalid because it has ..."
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:166
func (p *FieldRanges) CheckValid(isMessageSet bool) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:166
	_go_fuzz_dep_.CoverTab[53152]++
														var rp fieldRange
														for i, r := range p.lazyInit().sorted {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:168
		_go_fuzz_dep_.CoverTab[53154]++
															r := fieldRange(r)
															switch {
		case !isValidFieldNumber(r.Start(), isMessageSet):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:171
			_go_fuzz_dep_.CoverTab[53156]++
																return errors.New("invalid field number: %d", r.Start())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:172
			// _ = "end of CoverTab[53156]"
		case !isValidFieldNumber(r.End(), isMessageSet):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:173
			_go_fuzz_dep_.CoverTab[53157]++
																return errors.New("invalid field number: %d", r.End())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:174
			// _ = "end of CoverTab[53157]"
		case !(r.Start() <= r.End()):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:175
			_go_fuzz_dep_.CoverTab[53158]++
																return errors.New("invalid range: %v", r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:176
			// _ = "end of CoverTab[53158]"
		case !(rp.End() < r.Start()) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:177
			_go_fuzz_dep_.CoverTab[53161]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:177
			return i > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:177
			// _ = "end of CoverTab[53161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:177
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:177
			_go_fuzz_dep_.CoverTab[53159]++
																return errors.New("overlapping ranges: %v with %v", rp, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:178
			// _ = "end of CoverTab[53159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:178
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:178
			_go_fuzz_dep_.CoverTab[53160]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:178
			// _ = "end of CoverTab[53160]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:179
		// _ = "end of CoverTab[53154]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:179
		_go_fuzz_dep_.CoverTab[53155]++
															rp = r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:180
		// _ = "end of CoverTab[53155]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:181
	// _ = "end of CoverTab[53152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:181
	_go_fuzz_dep_.CoverTab[53153]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:182
	// _ = "end of CoverTab[53153]"
}

// isValidFieldNumber reports whether the field number is valid.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:185
// Unlike the FieldNumber.IsValid method, it allows ranges that cover the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:185
// reserved number range.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:188
func isValidFieldNumber(n protoreflect.FieldNumber, isMessageSet bool) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:188
	_go_fuzz_dep_.CoverTab[53162]++
														return protowire.MinValidNumber <= n && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
		_go_fuzz_dep_.CoverTab[53163]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
		return (n <= protowire.MaxValidNumber || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
			_go_fuzz_dep_.CoverTab[53164]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
			return isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
			// _ = "end of CoverTab[53164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
		// _ = "end of CoverTab[53163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:189
	// _ = "end of CoverTab[53162]"
}

// CheckOverlap reports an error if p and q overlap.
func (p *FieldRanges) CheckOverlap(q *FieldRanges) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:193
	_go_fuzz_dep_.CoverTab[53165]++
														rps := p.lazyInit().sorted
														rqs := q.lazyInit().sorted
														for pi, qi := 0, 0; pi < len(rps) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:196
		_go_fuzz_dep_.CoverTab[53167]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:196
		return qi < len(rqs)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:196
		// _ = "end of CoverTab[53167]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:196
	}(); {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:196
		_go_fuzz_dep_.CoverTab[53168]++
															rp := fieldRange(rps[pi])
															rq := fieldRange(rqs[qi])
															if !(rp.End() < rq.Start() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:199
			_go_fuzz_dep_.CoverTab[53170]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:199
			return rq.End() < rp.Start()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:199
			// _ = "end of CoverTab[53170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:199
		}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:199
			_go_fuzz_dep_.CoverTab[53171]++
																return errors.New("overlapping ranges: %v with %v", rp, rq)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:200
			// _ = "end of CoverTab[53171]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:201
			_go_fuzz_dep_.CoverTab[53172]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:201
			// _ = "end of CoverTab[53172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:201
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:201
		// _ = "end of CoverTab[53168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:201
		_go_fuzz_dep_.CoverTab[53169]++
															if rp.Start() < rq.Start() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:202
			_go_fuzz_dep_.CoverTab[53173]++
																pi++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:203
			// _ = "end of CoverTab[53173]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:204
			_go_fuzz_dep_.CoverTab[53174]++
																qi++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:205
			// _ = "end of CoverTab[53174]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:206
		// _ = "end of CoverTab[53169]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:207
	// _ = "end of CoverTab[53165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:207
	_go_fuzz_dep_.CoverTab[53166]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:208
	// _ = "end of CoverTab[53166]"
}

type fieldRange [2]protoreflect.FieldNumber

func (r fieldRange) Start() protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:213
	_go_fuzz_dep_.CoverTab[53175]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:213
	return r[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:213
	// _ = "end of CoverTab[53175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:213
}
func (r fieldRange) End() protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:214
	_go_fuzz_dep_.CoverTab[53176]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:214
	return r[1] - 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:214
	// _ = "end of CoverTab[53176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:214
}
func (r fieldRange) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:215
	_go_fuzz_dep_.CoverTab[53177]++
														if r.Start() == r.End() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:216
		_go_fuzz_dep_.CoverTab[53179]++
															return fmt.Sprintf("%d", r.Start())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:217
		// _ = "end of CoverTab[53179]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:218
		_go_fuzz_dep_.CoverTab[53180]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:218
		// _ = "end of CoverTab[53180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:218
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:218
	// _ = "end of CoverTab[53177]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:218
	_go_fuzz_dep_.CoverTab[53178]++
														return fmt.Sprintf("%d to %d", r.Start(), r.End())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:219
	// _ = "end of CoverTab[53178]"
}

type FieldNumbers struct {
	List	[]protoreflect.FieldNumber
	once	sync.Once
	has	map[protoreflect.FieldNumber]struct{}	// protected by once
}

func (p *FieldNumbers) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:228
	_go_fuzz_dep_.CoverTab[53181]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:228
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:228
	// _ = "end of CoverTab[53181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:228
}
func (p *FieldNumbers) Get(i int) protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:229
	_go_fuzz_dep_.CoverTab[53182]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:229
	return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:229
	// _ = "end of CoverTab[53182]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:229
}
func (p *FieldNumbers) Has(n protoreflect.FieldNumber) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:230
	_go_fuzz_dep_.CoverTab[53183]++
														p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:231
		_go_fuzz_dep_.CoverTab[53185]++
															if len(p.List) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:232
			_go_fuzz_dep_.CoverTab[53186]++
																p.has = make(map[protoreflect.FieldNumber]struct{}, len(p.List))
																for _, n := range p.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:234
				_go_fuzz_dep_.CoverTab[53187]++
																	p.has[n] = struct{}{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:235
				// _ = "end of CoverTab[53187]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:236
			// _ = "end of CoverTab[53186]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:237
			_go_fuzz_dep_.CoverTab[53188]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:237
			// _ = "end of CoverTab[53188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:237
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:237
		// _ = "end of CoverTab[53185]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:238
	// _ = "end of CoverTab[53183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:238
	_go_fuzz_dep_.CoverTab[53184]++
														_, ok := p.has[n]
														return ok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:240
	// _ = "end of CoverTab[53184]"
}
func (p *FieldNumbers) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:242
	_go_fuzz_dep_.CoverTab[53189]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:242
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:242
	// _ = "end of CoverTab[53189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:242
}
func (p *FieldNumbers) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:243
	_go_fuzz_dep_.CoverTab[53190]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:243
	// _ = "end of CoverTab[53190]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:243
}

type OneofFields struct {
	List	[]protoreflect.FieldDescriptor
	once	sync.Once
	byName	map[protoreflect.Name]protoreflect.FieldDescriptor		// protected by once
	byJSON	map[string]protoreflect.FieldDescriptor				// protected by once
	byText	map[string]protoreflect.FieldDescriptor				// protected by once
	byNum	map[protoreflect.FieldNumber]protoreflect.FieldDescriptor	// protected by once
}

func (p *OneofFields) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:254
	_go_fuzz_dep_.CoverTab[53191]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:254
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:254
	// _ = "end of CoverTab[53191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:254
}
func (p *OneofFields) Get(i int) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:255
	_go_fuzz_dep_.CoverTab[53192]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:255
	return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:255
	// _ = "end of CoverTab[53192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:255
}
func (p *OneofFields) ByName(s protoreflect.Name) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:256
	_go_fuzz_dep_.CoverTab[53193]++
														return p.lazyInit().byName[s]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:257
	// _ = "end of CoverTab[53193]"
}
func (p *OneofFields) ByJSONName(s string) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:259
	_go_fuzz_dep_.CoverTab[53194]++
														return p.lazyInit().byJSON[s]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:260
	// _ = "end of CoverTab[53194]"
}
func (p *OneofFields) ByTextName(s string) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:262
	_go_fuzz_dep_.CoverTab[53195]++
														return p.lazyInit().byText[s]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:263
	// _ = "end of CoverTab[53195]"
}
func (p *OneofFields) ByNumber(n protoreflect.FieldNumber) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:265
	_go_fuzz_dep_.CoverTab[53196]++
														return p.lazyInit().byNum[n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:266
	// _ = "end of CoverTab[53196]"
}
func (p *OneofFields) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:268
	_go_fuzz_dep_.CoverTab[53197]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:268
	descfmt.FormatList(s, r, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:268
	// _ = "end of CoverTab[53197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:268
}
func (p *OneofFields) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:269
	_go_fuzz_dep_.CoverTab[53198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:269
	// _ = "end of CoverTab[53198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:269
}

func (p *OneofFields) lazyInit() *OneofFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:271
	_go_fuzz_dep_.CoverTab[53199]++
														p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:272
		_go_fuzz_dep_.CoverTab[53201]++
															if len(p.List) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:273
			_go_fuzz_dep_.CoverTab[53202]++
																p.byName = make(map[protoreflect.Name]protoreflect.FieldDescriptor, len(p.List))
																p.byJSON = make(map[string]protoreflect.FieldDescriptor, len(p.List))
																p.byText = make(map[string]protoreflect.FieldDescriptor, len(p.List))
																p.byNum = make(map[protoreflect.FieldNumber]protoreflect.FieldDescriptor, len(p.List))
																for _, f := range p.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:278
				_go_fuzz_dep_.CoverTab[53203]++

																	p.byName[f.Name()] = f
																	p.byJSON[f.JSONName()] = f
																	p.byText[f.TextName()] = f
																	p.byNum[f.Number()] = f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:283
				// _ = "end of CoverTab[53203]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:284
			// _ = "end of CoverTab[53202]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:285
			_go_fuzz_dep_.CoverTab[53204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:285
			// _ = "end of CoverTab[53204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:285
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:285
		// _ = "end of CoverTab[53201]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:286
	// _ = "end of CoverTab[53199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:286
	_go_fuzz_dep_.CoverTab[53200]++
														return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:287
	// _ = "end of CoverTab[53200]"
}

type SourceLocations struct {
	// List is a list of SourceLocations.
	// The SourceLocation.Next field does not need to be populated
	// as it will be lazily populated upon first need.
	List	[]protoreflect.SourceLocation

	// File is the parent file descriptor that these locations are relative to.
	// If non-nil, ByDescriptor verifies that the provided descriptor
	// is a child of this file descriptor.
	File	protoreflect.FileDescriptor

	once	sync.Once
	byPath	map[pathKey]int
}

func (p *SourceLocations) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:305
	_go_fuzz_dep_.CoverTab[53205]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:305
	return len(p.List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:305
	// _ = "end of CoverTab[53205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:305
}
func (p *SourceLocations) Get(i int) protoreflect.SourceLocation {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:306
	_go_fuzz_dep_.CoverTab[53206]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:306
	return p.lazyInit().List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:306
	// _ = "end of CoverTab[53206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:306
}
func (p *SourceLocations) byKey(k pathKey) protoreflect.SourceLocation {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:307
	_go_fuzz_dep_.CoverTab[53207]++
														if i, ok := p.lazyInit().byPath[k]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:308
		_go_fuzz_dep_.CoverTab[53209]++
															return p.List[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:309
		// _ = "end of CoverTab[53209]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:310
		_go_fuzz_dep_.CoverTab[53210]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:310
		// _ = "end of CoverTab[53210]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:310
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:310
	// _ = "end of CoverTab[53207]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:310
	_go_fuzz_dep_.CoverTab[53208]++
														return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:311
	// _ = "end of CoverTab[53208]"
}
func (p *SourceLocations) ByPath(path protoreflect.SourcePath) protoreflect.SourceLocation {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:313
	_go_fuzz_dep_.CoverTab[53211]++
														return p.byKey(newPathKey(path))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:314
	// _ = "end of CoverTab[53211]"
}
func (p *SourceLocations) ByDescriptor(desc protoreflect.Descriptor) protoreflect.SourceLocation {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:316
	_go_fuzz_dep_.CoverTab[53212]++
														if p.File != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		_go_fuzz_dep_.CoverTab[53214]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		return desc != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		// _ = "end of CoverTab[53214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		_go_fuzz_dep_.CoverTab[53215]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		return p.File != desc.ParentFile()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		// _ = "end of CoverTab[53215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:317
		_go_fuzz_dep_.CoverTab[53216]++
															return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:318
		// _ = "end of CoverTab[53216]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:319
		_go_fuzz_dep_.CoverTab[53217]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:319
		// _ = "end of CoverTab[53217]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:319
	// _ = "end of CoverTab[53212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:319
	_go_fuzz_dep_.CoverTab[53213]++
														var pathArr [16]int32
														path := pathArr[:0]
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:322
		_go_fuzz_dep_.CoverTab[53218]++
															switch desc.(type) {
		case protoreflect.FileDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:324
			_go_fuzz_dep_.CoverTab[53219]++

																for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:326
				_go_fuzz_dep_.CoverTab[53229]++
																	path[i], path[j] = path[j], path[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:327
				// _ = "end of CoverTab[53229]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:328
			// _ = "end of CoverTab[53219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:328
			_go_fuzz_dep_.CoverTab[53220]++
																return p.byKey(newPathKey(path))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:329
			// _ = "end of CoverTab[53220]"
		case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:330
			_go_fuzz_dep_.CoverTab[53221]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.FileDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:334
				_go_fuzz_dep_.CoverTab[53230]++
																	path = append(path, int32(genid.FileDescriptorProto_MessageType_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:335
				// _ = "end of CoverTab[53230]"
			case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:336
				_go_fuzz_dep_.CoverTab[53231]++
																	path = append(path, int32(genid.DescriptorProto_NestedType_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:337
				// _ = "end of CoverTab[53231]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:338
				_go_fuzz_dep_.CoverTab[53232]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:339
				// _ = "end of CoverTab[53232]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:340
			// _ = "end of CoverTab[53221]"
		case protoreflect.FieldDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:341
			_go_fuzz_dep_.CoverTab[53222]++
																isExtension := desc.(protoreflect.FieldDescriptor).IsExtension()
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																if isExtension {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:345
				_go_fuzz_dep_.CoverTab[53233]++
																	switch desc.(type) {
				case protoreflect.FileDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:347
					_go_fuzz_dep_.CoverTab[53234]++
																		path = append(path, int32(genid.FileDescriptorProto_Extension_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:348
					// _ = "end of CoverTab[53234]"
				case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:349
					_go_fuzz_dep_.CoverTab[53235]++
																		path = append(path, int32(genid.DescriptorProto_Extension_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:350
					// _ = "end of CoverTab[53235]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:351
					_go_fuzz_dep_.CoverTab[53236]++
																		return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:352
					// _ = "end of CoverTab[53236]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:353
				// _ = "end of CoverTab[53233]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:354
				_go_fuzz_dep_.CoverTab[53237]++
																	switch desc.(type) {
				case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:356
					_go_fuzz_dep_.CoverTab[53238]++
																		path = append(path, int32(genid.DescriptorProto_Field_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:357
					// _ = "end of CoverTab[53238]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:358
					_go_fuzz_dep_.CoverTab[53239]++
																		return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:359
					// _ = "end of CoverTab[53239]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:360
				// _ = "end of CoverTab[53237]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:361
			// _ = "end of CoverTab[53222]"
		case protoreflect.OneofDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:362
			_go_fuzz_dep_.CoverTab[53223]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:366
				_go_fuzz_dep_.CoverTab[53240]++
																	path = append(path, int32(genid.DescriptorProto_OneofDecl_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:367
				// _ = "end of CoverTab[53240]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:368
				_go_fuzz_dep_.CoverTab[53241]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:369
				// _ = "end of CoverTab[53241]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:370
			// _ = "end of CoverTab[53223]"
		case protoreflect.EnumDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:371
			_go_fuzz_dep_.CoverTab[53224]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.FileDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:375
				_go_fuzz_dep_.CoverTab[53242]++
																	path = append(path, int32(genid.FileDescriptorProto_EnumType_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:376
				// _ = "end of CoverTab[53242]"
			case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:377
				_go_fuzz_dep_.CoverTab[53243]++
																	path = append(path, int32(genid.DescriptorProto_EnumType_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:378
				// _ = "end of CoverTab[53243]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:379
				_go_fuzz_dep_.CoverTab[53244]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:380
				// _ = "end of CoverTab[53244]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:381
			// _ = "end of CoverTab[53224]"
		case protoreflect.EnumValueDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:382
			_go_fuzz_dep_.CoverTab[53225]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.EnumDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:386
				_go_fuzz_dep_.CoverTab[53245]++
																	path = append(path, int32(genid.EnumDescriptorProto_Value_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:387
				// _ = "end of CoverTab[53245]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:388
				_go_fuzz_dep_.CoverTab[53246]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:389
				// _ = "end of CoverTab[53246]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:390
			// _ = "end of CoverTab[53225]"
		case protoreflect.ServiceDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:391
			_go_fuzz_dep_.CoverTab[53226]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.FileDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:395
				_go_fuzz_dep_.CoverTab[53247]++
																	path = append(path, int32(genid.FileDescriptorProto_Service_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:396
				// _ = "end of CoverTab[53247]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:397
				_go_fuzz_dep_.CoverTab[53248]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:398
				// _ = "end of CoverTab[53248]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:399
			// _ = "end of CoverTab[53226]"
		case protoreflect.MethodDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:400
			_go_fuzz_dep_.CoverTab[53227]++
																path = append(path, int32(desc.Index()))
																desc = desc.Parent()
																switch desc.(type) {
			case protoreflect.ServiceDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:404
				_go_fuzz_dep_.CoverTab[53249]++
																	path = append(path, int32(genid.ServiceDescriptorProto_Method_field_number))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:405
				// _ = "end of CoverTab[53249]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:406
				_go_fuzz_dep_.CoverTab[53250]++
																	return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:407
				// _ = "end of CoverTab[53250]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:408
			// _ = "end of CoverTab[53227]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:409
			_go_fuzz_dep_.CoverTab[53228]++
																return protoreflect.SourceLocation{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:410
			// _ = "end of CoverTab[53228]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:411
		// _ = "end of CoverTab[53218]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:412
	// _ = "end of CoverTab[53213]"
}
func (p *SourceLocations) lazyInit() *SourceLocations {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:414
	_go_fuzz_dep_.CoverTab[53251]++
														p.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:415
		_go_fuzz_dep_.CoverTab[53253]++
															if len(p.List) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:416
			_go_fuzz_dep_.CoverTab[53254]++

																pathIdxs := make(map[pathKey][]int, len(p.List))
																for i, l := range p.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:419
				_go_fuzz_dep_.CoverTab[53256]++
																	k := newPathKey(l.Path)
																	pathIdxs[k] = append(pathIdxs[k], i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:421
				// _ = "end of CoverTab[53256]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:422
			// _ = "end of CoverTab[53254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:422
			_go_fuzz_dep_.CoverTab[53255]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:425
			p.byPath = make(map[pathKey]int, len(p.List))
			for k, idxs := range pathIdxs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:426
				_go_fuzz_dep_.CoverTab[53257]++
																	for i := 0; i < len(idxs)-1; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:427
					_go_fuzz_dep_.CoverTab[53259]++
																		p.List[idxs[i]].Next = idxs[i+1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:428
					// _ = "end of CoverTab[53259]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:429
				// _ = "end of CoverTab[53257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:429
				_go_fuzz_dep_.CoverTab[53258]++
																	p.List[idxs[len(idxs)-1]].Next = 0
																	p.byPath[k] = idxs[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:431
				// _ = "end of CoverTab[53258]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:432
			// _ = "end of CoverTab[53255]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:433
			_go_fuzz_dep_.CoverTab[53260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:433
			// _ = "end of CoverTab[53260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:433
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:433
		// _ = "end of CoverTab[53253]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:434
	// _ = "end of CoverTab[53251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:434
	_go_fuzz_dep_.CoverTab[53252]++
														return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:435
	// _ = "end of CoverTab[53252]"
}
func (p *SourceLocations) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:437
	_go_fuzz_dep_.CoverTab[53261]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:437
	// _ = "end of CoverTab[53261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:437
}

// pathKey is a comparable representation of protoreflect.SourcePath.
type pathKey struct {
	arr	[16]uint8	// first n-1 path segments; last element is the length
	str	string		// used if the path does not fit in arr
}

func newPathKey(p protoreflect.SourcePath) (k pathKey) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:445
	_go_fuzz_dep_.CoverTab[53262]++
														if len(p) < len(k.arr) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:446
		_go_fuzz_dep_.CoverTab[53264]++
															for i, ps := range p {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:447
			_go_fuzz_dep_.CoverTab[53266]++
																if ps < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:448
				_go_fuzz_dep_.CoverTab[53268]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:448
				return math.MaxUint8 <= ps
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:448
				// _ = "end of CoverTab[53268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:448
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:448
				_go_fuzz_dep_.CoverTab[53269]++
																	return pathKey{str: p.String()}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:449
				// _ = "end of CoverTab[53269]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:450
				_go_fuzz_dep_.CoverTab[53270]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:450
				// _ = "end of CoverTab[53270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:450
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:450
			// _ = "end of CoverTab[53266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:450
			_go_fuzz_dep_.CoverTab[53267]++
																k.arr[i] = uint8(ps)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:451
			// _ = "end of CoverTab[53267]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:452
		// _ = "end of CoverTab[53264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:452
		_go_fuzz_dep_.CoverTab[53265]++
															k.arr[len(k.arr)-1] = uint8(len(p))
															return k
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:454
		// _ = "end of CoverTab[53265]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:455
		_go_fuzz_dep_.CoverTab[53271]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:455
		// _ = "end of CoverTab[53271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:455
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:455
	// _ = "end of CoverTab[53262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:455
	_go_fuzz_dep_.CoverTab[53263]++
														return pathKey{str: p.String()}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:456
	// _ = "end of CoverTab[53263]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:457
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_list.go:457
var _ = _go_fuzz_dep_.CoverTab
