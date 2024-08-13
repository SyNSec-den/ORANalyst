// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author  			xeipuuv
// author-github 	https://github.com/xeipuuv
// author-mail		xeipuuv@gmail.com
//
// repository-name	gojsonreference
// repository-desc	An implementation of JSON Reference - Go language
//
// description		Main and unique file.
//
// created      	26-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
package gojsonreference

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:26
)

import (
	"errors"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xeipuuv/gojsonpointer"
)

const (
	const_fragment_char = `#`
)

func NewJsonReference(jsonReferenceString string) (JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:42
	_go_fuzz_dep_.CoverTab[194702]++

																var r JsonReference
																err := r.parse(jsonReferenceString)
																return r, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:46
	// _ = "end of CoverTab[194702]"

}

type JsonReference struct {
	referenceUrl		*url.URL
	referencePointer	gojsonpointer.JsonPointer

	HasFullUrl	bool
	HasUrlPathOnly	bool
	HasFragmentOnly	bool
	HasFileScheme	bool
	HasFullFilePath	bool
}

func (r *JsonReference) GetUrl() *url.URL {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:61
	_go_fuzz_dep_.CoverTab[194703]++
																return r.referenceUrl
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:62
	// _ = "end of CoverTab[194703]"
}

func (r *JsonReference) GetPointer() *gojsonpointer.JsonPointer {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:65
	_go_fuzz_dep_.CoverTab[194704]++
																return &r.referencePointer
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:66
	// _ = "end of CoverTab[194704]"
}

func (r *JsonReference) String() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:69
	_go_fuzz_dep_.CoverTab[194705]++

																if r.referenceUrl != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:71
		_go_fuzz_dep_.CoverTab[194708]++
																	return r.referenceUrl.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:72
		// _ = "end of CoverTab[194708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:73
		_go_fuzz_dep_.CoverTab[194709]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:73
		// _ = "end of CoverTab[194709]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:73
	// _ = "end of CoverTab[194705]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:73
	_go_fuzz_dep_.CoverTab[194706]++

																if r.HasFragmentOnly {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:75
		_go_fuzz_dep_.CoverTab[194710]++
																	return const_fragment_char + r.referencePointer.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:76
		// _ = "end of CoverTab[194710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:77
		_go_fuzz_dep_.CoverTab[194711]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:77
		// _ = "end of CoverTab[194711]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:77
	// _ = "end of CoverTab[194706]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:77
	_go_fuzz_dep_.CoverTab[194707]++

																return r.referencePointer.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:79
	// _ = "end of CoverTab[194707]"
}

func (r *JsonReference) IsCanonical() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:82
	_go_fuzz_dep_.CoverTab[194712]++
																return (r.HasFileScheme && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		_go_fuzz_dep_.CoverTab[194713]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		return r.HasFullFilePath
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		// _ = "end of CoverTab[194713]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		_go_fuzz_dep_.CoverTab[194714]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		return (!r.HasFileScheme && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
			_go_fuzz_dep_.CoverTab[194715]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
			return r.HasFullUrl
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
			// _ = "end of CoverTab[194715]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		}())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
		// _ = "end of CoverTab[194714]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:83
	// _ = "end of CoverTab[194712]"
}

// "Constructor", parses the given string JSON reference
func (r *JsonReference) parse(jsonReferenceString string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:87
	_go_fuzz_dep_.CoverTab[194716]++

																r.referenceUrl, err = url.Parse(jsonReferenceString)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:90
		_go_fuzz_dep_.CoverTab[194720]++
																	return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:91
		// _ = "end of CoverTab[194720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:92
		_go_fuzz_dep_.CoverTab[194721]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:92
		// _ = "end of CoverTab[194721]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:92
	// _ = "end of CoverTab[194716]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:92
	_go_fuzz_dep_.CoverTab[194717]++
																refUrl := r.referenceUrl

																if refUrl.Scheme != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:95
		_go_fuzz_dep_.CoverTab[194722]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:95
		return refUrl.Host != ""
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:95
		// _ = "end of CoverTab[194722]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:95
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:95
		_go_fuzz_dep_.CoverTab[194723]++
																	r.HasFullUrl = true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:96
		// _ = "end of CoverTab[194723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:97
		_go_fuzz_dep_.CoverTab[194724]++
																	if refUrl.Path != "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:98
			_go_fuzz_dep_.CoverTab[194725]++
																		r.HasUrlPathOnly = true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:99
			// _ = "end of CoverTab[194725]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
			_go_fuzz_dep_.CoverTab[194726]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
			if refUrl.RawQuery == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
				_go_fuzz_dep_.CoverTab[194727]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
				return refUrl.Fragment != ""
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
				// _ = "end of CoverTab[194727]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
			}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:100
				_go_fuzz_dep_.CoverTab[194728]++
																			r.HasFragmentOnly = true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:101
				// _ = "end of CoverTab[194728]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
				_go_fuzz_dep_.CoverTab[194729]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
				// _ = "end of CoverTab[194729]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
			// _ = "end of CoverTab[194726]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:102
		// _ = "end of CoverTab[194724]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:103
	// _ = "end of CoverTab[194717]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:103
	_go_fuzz_dep_.CoverTab[194718]++

																r.HasFileScheme = refUrl.Scheme == "file"
																if runtime.GOOS == "windows" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:106
		_go_fuzz_dep_.CoverTab[194730]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
		if refUrl.Host == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
			_go_fuzz_dep_.CoverTab[194731]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
			return strings.HasPrefix(refUrl.Path, "/")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
			// _ = "end of CoverTab[194731]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
		}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:110
			_go_fuzz_dep_.CoverTab[194732]++
																		r.HasFullFilePath = filepath.IsAbs(refUrl.Path[1:])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:111
			// _ = "end of CoverTab[194732]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:112
			_go_fuzz_dep_.CoverTab[194733]++
																		r.HasFullFilePath = filepath.IsAbs(refUrl.Host + refUrl.Path)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:113
			// _ = "end of CoverTab[194733]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:114
		// _ = "end of CoverTab[194730]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:115
		_go_fuzz_dep_.CoverTab[194734]++
																	r.HasFullFilePath = filepath.IsAbs(refUrl.Path)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:116
		// _ = "end of CoverTab[194734]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:117
	// _ = "end of CoverTab[194718]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:117
	_go_fuzz_dep_.CoverTab[194719]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:120
	r.referencePointer, _ = gojsonpointer.NewJsonPointer(refUrl.Fragment)

																return
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:122
	// _ = "end of CoverTab[194719]"
}

// Creates a new reference from a parent and a child
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:125
// If the child cannot inherit from the parent, an error is returned
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:127
func (r *JsonReference) Inherits(child JsonReference) (*JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:127
	_go_fuzz_dep_.CoverTab[194735]++
																if child.GetUrl() == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:128
		_go_fuzz_dep_.CoverTab[194739]++
																	return nil, errors.New("childUrl is nil!")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:129
		// _ = "end of CoverTab[194739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:130
		_go_fuzz_dep_.CoverTab[194740]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:130
		// _ = "end of CoverTab[194740]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:130
	// _ = "end of CoverTab[194735]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:130
	_go_fuzz_dep_.CoverTab[194736]++

																if r.GetUrl() == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:132
		_go_fuzz_dep_.CoverTab[194741]++
																	return nil, errors.New("parentUrl is nil!")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:133
		// _ = "end of CoverTab[194741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:134
		_go_fuzz_dep_.CoverTab[194742]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:134
		// _ = "end of CoverTab[194742]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:134
	// _ = "end of CoverTab[194736]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:134
	_go_fuzz_dep_.CoverTab[194737]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:139
	parentUrl := *r.GetUrl()
	parentUrl.Fragment = ""

	ref, err := NewJsonReference(parentUrl.ResolveReference(child.GetUrl()).String())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:143
		_go_fuzz_dep_.CoverTab[194743]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:144
		// _ = "end of CoverTab[194743]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:145
		_go_fuzz_dep_.CoverTab[194744]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:145
		// _ = "end of CoverTab[194744]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:145
	// _ = "end of CoverTab[194737]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:145
	_go_fuzz_dep_.CoverTab[194738]++
																return &ref, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:146
	// _ = "end of CoverTab[194738]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:147
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonreference@v0.0.0-20180127040603-bd5ef7bd5415/reference.go:147
var _ = _go_fuzz_dep_.CoverTab
