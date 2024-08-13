// Copyright 2019 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:15
)

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	_	dataSource	= (*sourceFile)(nil)
	_	dataSource	= (*sourceData)(nil)
	_	dataSource	= (*sourceReadCloser)(nil)
)

// dataSource is an interface that returns object which can be read and closed.
type dataSource interface {
	ReadCloser() (io.ReadCloser, error)
}

// sourceFile represents an object that contains content on the local file system.
type sourceFile struct {
	name string
}

func (s sourceFile) ReadCloser() (_ io.ReadCloser, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:41
	_go_fuzz_dep_.CoverTab[128236]++
										return os.Open(s.name)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:42
	// _ = "end of CoverTab[128236]"
}

// sourceData represents an object that contains content in memory.
type sourceData struct {
	data []byte
}

func (s *sourceData) ReadCloser() (io.ReadCloser, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:50
	_go_fuzz_dep_.CoverTab[128237]++
										return ioutil.NopCloser(bytes.NewReader(s.data)), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:51
	// _ = "end of CoverTab[128237]"
}

// sourceReadCloser represents an input stream with Close method.
type sourceReadCloser struct {
	reader io.ReadCloser
}

func (s *sourceReadCloser) ReadCloser() (io.ReadCloser, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:59
	_go_fuzz_dep_.CoverTab[128238]++
										return s.reader, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:60
	// _ = "end of CoverTab[128238]"
}

func parseDataSource(source interface{}) (dataSource, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:63
	_go_fuzz_dep_.CoverTab[128239]++
										switch s := source.(type) {
	case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:65
		_go_fuzz_dep_.CoverTab[128240]++
											return sourceFile{s}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:66
		// _ = "end of CoverTab[128240]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:67
		_go_fuzz_dep_.CoverTab[128241]++
											return &sourceData{s}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:68
		// _ = "end of CoverTab[128241]"
	case io.ReadCloser:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:69
		_go_fuzz_dep_.CoverTab[128242]++
											return &sourceReadCloser{s}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:70
		// _ = "end of CoverTab[128242]"
	case io.Reader:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:71
		_go_fuzz_dep_.CoverTab[128243]++
											return &sourceReadCloser{ioutil.NopCloser(s)}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:72
		// _ = "end of CoverTab[128243]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:73
		_go_fuzz_dep_.CoverTab[128244]++
											return nil, fmt.Errorf("error parsing data source: unknown type %q", s)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:74
		// _ = "end of CoverTab[128244]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:75
	// _ = "end of CoverTab[128239]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/data_source.go:76
var _ = _go_fuzz_dep_.CoverTab
