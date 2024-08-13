// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:5
)

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Encoding specifies encoding of the input data.
type Encoding uint

const (
	// utf8Default is a private placeholder for the zero value of Encoding to
	// ensure that it has the correct meaning. UTF8 is the default encoding but
	// was assigned a non-zero value which cannot be changed without breaking
	// existing code. Clients should continue to use the public constants.
	utf8Default	Encoding	= iota

	// UTF8 interprets the input data as UTF-8.
	UTF8

	// ISO_8859_1 interprets the input data as ISO-8859-1.
	ISO_8859_1
)

type Loader struct {
	// Encoding determines how the data from files and byte buffers
	// is interpreted. For URLs the Content-Type header is used
	// to determine the encoding of the data.
	Encoding	Encoding

	// DisableExpansion configures the property expansion of the
	// returned property object. When set to true, the property values
	// will not be expanded and the Property object will not be checked
	// for invalid expansion expressions.
	DisableExpansion	bool

	// IgnoreMissing configures whether missing files or URLs which return
	// 404 are reported as errors. When set to true, missing files and 404
	// status codes are not reported as errors.
	IgnoreMissing	bool
}

// Load reads a buffer into a Properties struct.
func (l *Loader) LoadBytes(buf []byte) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:51
	_go_fuzz_dep_.CoverTab[115742]++
											return l.loadBytes(buf, l.Encoding)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:52
	// _ = "end of CoverTab[115742]"
}

// LoadAll reads the content of multiple URLs or files in the given order into
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:55
// a Properties struct. If IgnoreMissing is true then a 404 status code or
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:55
// missing file will not be reported as error. Encoding sets the encoding for
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:55
// files. For the URLs see LoadURL for the Content-Type header and the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:55
// encoding.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:60
func (l *Loader) LoadAll(names []string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:60
	_go_fuzz_dep_.CoverTab[115743]++
											all := NewProperties()
											for _, name := range names {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:62
		_go_fuzz_dep_.CoverTab[115746]++
												n, err := expandName(name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:64
			_go_fuzz_dep_.CoverTab[115750]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:65
			// _ = "end of CoverTab[115750]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:66
			_go_fuzz_dep_.CoverTab[115751]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:66
			// _ = "end of CoverTab[115751]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:66
		// _ = "end of CoverTab[115746]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:66
		_go_fuzz_dep_.CoverTab[115747]++

												var p *Properties
												switch {
		case strings.HasPrefix(n, "http://"):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:70
			_go_fuzz_dep_.CoverTab[115752]++
													p, err = l.LoadURL(n)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:71
			// _ = "end of CoverTab[115752]"
		case strings.HasPrefix(n, "https://"):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:72
			_go_fuzz_dep_.CoverTab[115753]++
													p, err = l.LoadURL(n)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:73
			// _ = "end of CoverTab[115753]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:74
			_go_fuzz_dep_.CoverTab[115754]++
													p, err = l.LoadFile(n)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:75
			// _ = "end of CoverTab[115754]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:76
		// _ = "end of CoverTab[115747]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:76
		_go_fuzz_dep_.CoverTab[115748]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:77
			_go_fuzz_dep_.CoverTab[115755]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:78
			// _ = "end of CoverTab[115755]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:79
			_go_fuzz_dep_.CoverTab[115756]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:79
			// _ = "end of CoverTab[115756]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:79
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:79
		// _ = "end of CoverTab[115748]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:79
		_go_fuzz_dep_.CoverTab[115749]++
												all.Merge(p)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:80
		// _ = "end of CoverTab[115749]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:81
	// _ = "end of CoverTab[115743]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:81
	_go_fuzz_dep_.CoverTab[115744]++

											all.DisableExpansion = l.DisableExpansion
											if all.DisableExpansion {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:84
		_go_fuzz_dep_.CoverTab[115757]++
												return all, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:85
		// _ = "end of CoverTab[115757]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:86
		_go_fuzz_dep_.CoverTab[115758]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:86
		// _ = "end of CoverTab[115758]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:86
	// _ = "end of CoverTab[115744]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:86
	_go_fuzz_dep_.CoverTab[115745]++
											return all, all.check()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:87
	// _ = "end of CoverTab[115745]"
}

// LoadFile reads a file into a Properties struct.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:90
// If IgnoreMissing is true then a missing file will not be
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:90
// reported as error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:93
func (l *Loader) LoadFile(filename string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:93
	_go_fuzz_dep_.CoverTab[115759]++
											data, err := ioutil.ReadFile(filename)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:95
		_go_fuzz_dep_.CoverTab[115761]++
												if l.IgnoreMissing && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:96
			_go_fuzz_dep_.CoverTab[115763]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:96
			return os.IsNotExist(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:96
			// _ = "end of CoverTab[115763]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:96
		}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:96
			_go_fuzz_dep_.CoverTab[115764]++
													LogPrintf("properties: %s not found. skipping", filename)
													return NewProperties(), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:98
			// _ = "end of CoverTab[115764]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:99
			_go_fuzz_dep_.CoverTab[115765]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:99
			// _ = "end of CoverTab[115765]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:99
		// _ = "end of CoverTab[115761]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:99
		_go_fuzz_dep_.CoverTab[115762]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:100
		// _ = "end of CoverTab[115762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:101
		_go_fuzz_dep_.CoverTab[115766]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:101
		// _ = "end of CoverTab[115766]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:101
	// _ = "end of CoverTab[115759]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:101
	_go_fuzz_dep_.CoverTab[115760]++
											return l.loadBytes(data, l.Encoding)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:102
	// _ = "end of CoverTab[115760]"
}

// LoadURL reads the content of the URL into a Properties struct.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// The encoding is determined via the Content-Type header which
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// should be set to 'text/plain'. If the 'charset' parameter is
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// missing, 'iso-8859-1' or 'latin1' the encoding is set to
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// ISO-8859-1. If the 'charset' parameter is set to 'utf-8' the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// encoding is set to UTF-8. A missing content type header is
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:105
// interpreted as 'text/plain; charset=utf-8'.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:113
func (l *Loader) LoadURL(url string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:113
	_go_fuzz_dep_.CoverTab[115767]++
											resp, err := http.Get(url)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:115
		_go_fuzz_dep_.CoverTab[115773]++
												return nil, fmt.Errorf("properties: error fetching %q. %s", url, err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:116
		// _ = "end of CoverTab[115773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:117
		_go_fuzz_dep_.CoverTab[115774]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:117
		// _ = "end of CoverTab[115774]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:117
	// _ = "end of CoverTab[115767]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:117
	_go_fuzz_dep_.CoverTab[115768]++
											defer resp.Body.Close()

											if resp.StatusCode == 404 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:120
		_go_fuzz_dep_.CoverTab[115775]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:120
		return l.IgnoreMissing
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:120
		// _ = "end of CoverTab[115775]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:120
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:120
		_go_fuzz_dep_.CoverTab[115776]++
												LogPrintf("properties: %s returned %d. skipping", url, resp.StatusCode)
												return NewProperties(), nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:122
		// _ = "end of CoverTab[115776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:123
		_go_fuzz_dep_.CoverTab[115777]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:123
		// _ = "end of CoverTab[115777]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:123
	// _ = "end of CoverTab[115768]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:123
	_go_fuzz_dep_.CoverTab[115769]++

											if resp.StatusCode != 200 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:125
		_go_fuzz_dep_.CoverTab[115778]++
												return nil, fmt.Errorf("properties: %s returned %d", url, resp.StatusCode)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:126
		// _ = "end of CoverTab[115778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:127
		_go_fuzz_dep_.CoverTab[115779]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:127
		// _ = "end of CoverTab[115779]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:127
	// _ = "end of CoverTab[115769]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:127
	_go_fuzz_dep_.CoverTab[115770]++

											body, err := ioutil.ReadAll(resp.Body)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:130
		_go_fuzz_dep_.CoverTab[115780]++
												return nil, fmt.Errorf("properties: %s error reading response. %s", url, err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:131
		// _ = "end of CoverTab[115780]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:132
		_go_fuzz_dep_.CoverTab[115781]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:132
		// _ = "end of CoverTab[115781]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:132
	// _ = "end of CoverTab[115770]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:132
	_go_fuzz_dep_.CoverTab[115771]++

											ct := resp.Header.Get("Content-Type")
											ct = strings.Join(strings.Fields(ct), "")
											var enc Encoding
											switch strings.ToLower(ct) {
	case "text/plain", "text/plain;charset=iso-8859-1", "text/plain;charset=latin1":
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:138
		_go_fuzz_dep_.CoverTab[115782]++
												enc = ISO_8859_1
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:139
		// _ = "end of CoverTab[115782]"
	case "", "text/plain;charset=utf-8":
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:140
		_go_fuzz_dep_.CoverTab[115783]++
												enc = UTF8
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:141
		// _ = "end of CoverTab[115783]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:142
		_go_fuzz_dep_.CoverTab[115784]++
												return nil, fmt.Errorf("properties: invalid content type %s", ct)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:143
		// _ = "end of CoverTab[115784]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:144
	// _ = "end of CoverTab[115771]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:144
	_go_fuzz_dep_.CoverTab[115772]++

											return l.loadBytes(body, enc)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:146
	// _ = "end of CoverTab[115772]"
}

func (l *Loader) loadBytes(buf []byte, enc Encoding) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:149
	_go_fuzz_dep_.CoverTab[115785]++
											p, err := parse(convert(buf, enc))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:151
		_go_fuzz_dep_.CoverTab[115788]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:152
		// _ = "end of CoverTab[115788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:153
		_go_fuzz_dep_.CoverTab[115789]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:153
		// _ = "end of CoverTab[115789]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:153
	// _ = "end of CoverTab[115785]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:153
	_go_fuzz_dep_.CoverTab[115786]++
											p.DisableExpansion = l.DisableExpansion
											if p.DisableExpansion {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:155
		_go_fuzz_dep_.CoverTab[115790]++
												return p, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:156
		// _ = "end of CoverTab[115790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:157
		_go_fuzz_dep_.CoverTab[115791]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:157
		// _ = "end of CoverTab[115791]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:157
	// _ = "end of CoverTab[115786]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:157
	_go_fuzz_dep_.CoverTab[115787]++
											return p, p.check()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:158
	// _ = "end of CoverTab[115787]"
}

// Load reads a buffer into a Properties struct.
func Load(buf []byte, enc Encoding) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:162
	_go_fuzz_dep_.CoverTab[115792]++
											l := &Loader{Encoding: enc}
											return l.LoadBytes(buf)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:164
	// _ = "end of CoverTab[115792]"
}

// LoadString reads an UTF8 string into a properties struct.
func LoadString(s string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:168
	_go_fuzz_dep_.CoverTab[115793]++
											l := &Loader{Encoding: UTF8}
											return l.LoadBytes([]byte(s))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:170
	// _ = "end of CoverTab[115793]"
}

// LoadMap creates a new Properties struct from a string map.
func LoadMap(m map[string]string) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:174
	_go_fuzz_dep_.CoverTab[115794]++
											p := NewProperties()
											for k, v := range m {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:176
		_go_fuzz_dep_.CoverTab[115796]++
												p.Set(k, v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:177
		// _ = "end of CoverTab[115796]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:178
	// _ = "end of CoverTab[115794]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:178
	_go_fuzz_dep_.CoverTab[115795]++
											return p
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:179
	// _ = "end of CoverTab[115795]"
}

// LoadFile reads a file into a Properties struct.
func LoadFile(filename string, enc Encoding) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:183
	_go_fuzz_dep_.CoverTab[115797]++
											l := &Loader{Encoding: enc}
											return l.LoadAll([]string{filename})
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:185
	// _ = "end of CoverTab[115797]"
}

// LoadFiles reads multiple files in the given order into
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:188
// a Properties struct. If 'ignoreMissing' is true then
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:188
// non-existent files will not be reported as error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:191
func LoadFiles(filenames []string, enc Encoding, ignoreMissing bool) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:191
	_go_fuzz_dep_.CoverTab[115798]++
											l := &Loader{Encoding: enc, IgnoreMissing: ignoreMissing}
											return l.LoadAll(filenames)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:193
	// _ = "end of CoverTab[115798]"
}

// LoadURL reads the content of the URL into a Properties struct.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:196
// See Loader#LoadURL for details.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:198
func LoadURL(url string) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:198
	_go_fuzz_dep_.CoverTab[115799]++
											l := &Loader{Encoding: UTF8}
											return l.LoadAll([]string{url})
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:200
	// _ = "end of CoverTab[115799]"
}

// LoadURLs reads the content of multiple URLs in the given order into a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:203
// Properties struct. If IgnoreMissing is true then a 404 status code will
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:203
// not be reported as error. See Loader#LoadURL for the Content-Type header
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:203
// and the encoding.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:207
func LoadURLs(urls []string, ignoreMissing bool) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:207
	_go_fuzz_dep_.CoverTab[115800]++
											l := &Loader{Encoding: UTF8, IgnoreMissing: ignoreMissing}
											return l.LoadAll(urls)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:209
	// _ = "end of CoverTab[115800]"
}

// LoadAll reads the content of multiple URLs or files in the given order into a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:212
// Properties struct. If 'ignoreMissing' is true then a 404 status code or missing file will
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:212
// not be reported as error. Encoding sets the encoding for files. For the URLs please see
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:212
// LoadURL for the Content-Type header and the encoding.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:216
func LoadAll(names []string, enc Encoding, ignoreMissing bool) (*Properties, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:216
	_go_fuzz_dep_.CoverTab[115801]++
											l := &Loader{Encoding: enc, IgnoreMissing: ignoreMissing}
											return l.LoadAll(names)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:218
	// _ = "end of CoverTab[115801]"
}

// MustLoadString reads an UTF8 string into a Properties struct and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:221
// panics on error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:223
func MustLoadString(s string) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:223
	_go_fuzz_dep_.CoverTab[115802]++
											return must(LoadString(s))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:224
	// _ = "end of CoverTab[115802]"
}

// MustLoadFile reads a file into a Properties struct and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:227
// panics on error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:229
func MustLoadFile(filename string, enc Encoding) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:229
	_go_fuzz_dep_.CoverTab[115803]++
											return must(LoadFile(filename, enc))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:230
	// _ = "end of CoverTab[115803]"
}

// MustLoadFiles reads multiple files in the given order into
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:233
// a Properties struct and panics on error. If 'ignoreMissing'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:233
// is true then non-existent files will not be reported as error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:236
func MustLoadFiles(filenames []string, enc Encoding, ignoreMissing bool) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:236
	_go_fuzz_dep_.CoverTab[115804]++
											return must(LoadFiles(filenames, enc, ignoreMissing))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:237
	// _ = "end of CoverTab[115804]"
}

// MustLoadURL reads the content of a URL into a Properties struct and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:240
// panics on error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:242
func MustLoadURL(url string) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:242
	_go_fuzz_dep_.CoverTab[115805]++
											return must(LoadURL(url))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:243
	// _ = "end of CoverTab[115805]"
}

// MustLoadURLs reads the content of multiple URLs in the given order into a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:246
// Properties struct and panics on error. If 'ignoreMissing' is true then a 404
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:246
// status code will not be reported as error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:249
func MustLoadURLs(urls []string, ignoreMissing bool) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:249
	_go_fuzz_dep_.CoverTab[115806]++
											return must(LoadURLs(urls, ignoreMissing))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:250
	// _ = "end of CoverTab[115806]"
}

// MustLoadAll reads the content of multiple URLs or files in the given order into a
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:253
// Properties struct. If 'ignoreMissing' is true then a 404 status code or missing file will
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:253
// not be reported as error. Encoding sets the encoding for files. For the URLs please see
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:253
// LoadURL for the Content-Type header and the encoding. It panics on error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:257
func MustLoadAll(names []string, enc Encoding, ignoreMissing bool) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:257
	_go_fuzz_dep_.CoverTab[115807]++
											return must(LoadAll(names, enc, ignoreMissing))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:258
	// _ = "end of CoverTab[115807]"
}

func must(p *Properties, err error) *Properties {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:261
	_go_fuzz_dep_.CoverTab[115808]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:262
		_go_fuzz_dep_.CoverTab[115810]++
												ErrorHandler(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:263
		// _ = "end of CoverTab[115810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:264
		_go_fuzz_dep_.CoverTab[115811]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:264
		// _ = "end of CoverTab[115811]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:264
	// _ = "end of CoverTab[115808]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:264
	_go_fuzz_dep_.CoverTab[115809]++
											return p
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:265
	// _ = "end of CoverTab[115809]"
}

// expandName expands ${ENV_VAR} expressions in a name.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:268
// If the environment variable does not exist then it will be replaced
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:268
// with an empty string. Malformed expressions like "${ENV_VAR" will
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:268
// be reported as error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:272
func expandName(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:272
	_go_fuzz_dep_.CoverTab[115812]++
											return expand(name, []string{}, "${", "}", make(map[string]string))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:273
	// _ = "end of CoverTab[115812]"
}

// Interprets a byte buffer either as an ISO-8859-1 or UTF-8 encoded string.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:276
// For ISO-8859-1 we can convert each byte straight into a rune since the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:276
// first 256 unicode code points cover ISO-8859-1.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:279
func convert(buf []byte, enc Encoding) string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:279
	_go_fuzz_dep_.CoverTab[115813]++
											switch enc {
	case utf8Default, UTF8:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:281
		_go_fuzz_dep_.CoverTab[115815]++
												return string(buf)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:282
		// _ = "end of CoverTab[115815]"
	case ISO_8859_1:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:283
		_go_fuzz_dep_.CoverTab[115816]++
												runes := make([]rune, len(buf))
												for i, b := range buf {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:285
			_go_fuzz_dep_.CoverTab[115819]++
													runes[i] = rune(b)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:286
			// _ = "end of CoverTab[115819]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:287
		// _ = "end of CoverTab[115816]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:287
		_go_fuzz_dep_.CoverTab[115817]++
												return string(runes)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:288
		// _ = "end of CoverTab[115817]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:289
		_go_fuzz_dep_.CoverTab[115818]++
												ErrorHandler(fmt.Errorf("unsupported encoding %v", enc))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:290
		// _ = "end of CoverTab[115818]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:291
	// _ = "end of CoverTab[115813]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:291
	_go_fuzz_dep_.CoverTab[115814]++
											panic("ErrorHandler should exit")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:292
	// _ = "end of CoverTab[115814]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:293
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/load.go:293
var _ = _go_fuzz_dep_.CoverTab
