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

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description		Different strategies to load JSON files.
// 					Includes References (file and HTTP), JSON strings and Go types.
//
// created          01-02-2015

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:27
)

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xeipuuv/gojsonreference"
)

var osFS = osFileSystem(os.Open)

// JSONLoader defines the JSON loader interface
type JSONLoader interface {
	JsonSource() interface{}
	LoadJSON() (interface{}, error)
	JsonReference() (gojsonreference.JsonReference, error)
	LoaderFactory() JSONLoaderFactory
}

// JSONLoaderFactory defines the JSON loader factory interface
type JSONLoaderFactory interface {
	// New creates a new JSON loader for the given source
	New(source string) JSONLoader
}

// DefaultJSONLoaderFactory is the default JSON loader factory
type DefaultJSONLoaderFactory struct {
}

// FileSystemJSONLoaderFactory is a JSON loader factory that uses http.FileSystem
type FileSystemJSONLoaderFactory struct {
	fs http.FileSystem
}

// New creates a new JSON loader for the given source
func (d DefaultJSONLoaderFactory) New(source string) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:71
	_go_fuzz_dep_.CoverTab[194924]++
												return &jsonReferenceLoader{
		fs:	osFS,
		source:	source,
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:75
	// _ = "end of CoverTab[194924]"
}

// New creates a new JSON loader for the given source
func (f FileSystemJSONLoaderFactory) New(source string) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:79
	_go_fuzz_dep_.CoverTab[194925]++
												return &jsonReferenceLoader{
		fs:	f.fs,
		source:	source,
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:83
	// _ = "end of CoverTab[194925]"
}

// osFileSystem is a functional wrapper for os.Open that implements http.FileSystem.
type osFileSystem func(string) (*os.File, error)

// Opens a file with the given name
func (o osFileSystem) Open(name string) (http.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:90
	_go_fuzz_dep_.CoverTab[194926]++
												return o(name)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:91
	// _ = "end of CoverTab[194926]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:97
type jsonReferenceLoader struct {
	fs	http.FileSystem
	source	string
}

func (l *jsonReferenceLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:102
	_go_fuzz_dep_.CoverTab[194927]++
												return l.source
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:103
	// _ = "end of CoverTab[194927]"
}

func (l *jsonReferenceLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:106
	_go_fuzz_dep_.CoverTab[194928]++
												return gojsonreference.NewJsonReference(l.JsonSource().(string))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:107
	// _ = "end of CoverTab[194928]"
}

func (l *jsonReferenceLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:110
	_go_fuzz_dep_.CoverTab[194929]++
												return &FileSystemJSONLoaderFactory{
		fs: l.fs,
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:113
	// _ = "end of CoverTab[194929]"
}

// NewReferenceLoader returns a JSON reference loader using the given source and the local OS file system.
func NewReferenceLoader(source string) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:117
	_go_fuzz_dep_.CoverTab[194930]++
												return &jsonReferenceLoader{
		fs:	osFS,
		source:	source,
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:121
	// _ = "end of CoverTab[194930]"
}

// NewReferenceLoaderFileSystem returns a JSON reference loader using the given source and file system.
func NewReferenceLoaderFileSystem(source string, fs http.FileSystem) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:125
	_go_fuzz_dep_.CoverTab[194931]++
												return &jsonReferenceLoader{
		fs:	fs,
		source:	source,
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:129
	// _ = "end of CoverTab[194931]"
}

func (l *jsonReferenceLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:132
	_go_fuzz_dep_.CoverTab[194932]++

												var err error

												reference, err := gojsonreference.NewJsonReference(l.JsonSource().(string))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:137
		_go_fuzz_dep_.CoverTab[194935]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:138
		// _ = "end of CoverTab[194935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:139
		_go_fuzz_dep_.CoverTab[194936]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:139
		// _ = "end of CoverTab[194936]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:139
	// _ = "end of CoverTab[194932]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:139
	_go_fuzz_dep_.CoverTab[194933]++

												refToURL := reference
												refToURL.GetUrl().Fragment = ""

												var document interface{}

												if reference.HasFileScheme {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:146
		_go_fuzz_dep_.CoverTab[194937]++

													filename := strings.TrimPrefix(refToURL.String(), "file://")
													filename, err = url.QueryUnescape(filename)

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:151
			_go_fuzz_dep_.CoverTab[194940]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:152
			// _ = "end of CoverTab[194940]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:153
			_go_fuzz_dep_.CoverTab[194941]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:153
			// _ = "end of CoverTab[194941]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:153
		// _ = "end of CoverTab[194937]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:153
		_go_fuzz_dep_.CoverTab[194938]++

													if runtime.GOOS == "windows" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:155
			_go_fuzz_dep_.CoverTab[194942]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:158
			filename = strings.TrimPrefix(filename, "/")
														filename = filepath.FromSlash(filename)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:159
			// _ = "end of CoverTab[194942]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:160
			_go_fuzz_dep_.CoverTab[194943]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:160
			// _ = "end of CoverTab[194943]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:160
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:160
		// _ = "end of CoverTab[194938]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:160
		_go_fuzz_dep_.CoverTab[194939]++

													document, err = l.loadFromFile(filename)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:163
			_go_fuzz_dep_.CoverTab[194944]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:164
			// _ = "end of CoverTab[194944]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:165
			_go_fuzz_dep_.CoverTab[194945]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:165
			// _ = "end of CoverTab[194945]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:165
		// _ = "end of CoverTab[194939]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:167
		_go_fuzz_dep_.CoverTab[194946]++

													document, err = l.loadFromHTTP(refToURL.String())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:170
			_go_fuzz_dep_.CoverTab[194947]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:171
			// _ = "end of CoverTab[194947]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:172
			_go_fuzz_dep_.CoverTab[194948]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:172
			// _ = "end of CoverTab[194948]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:172
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:172
		// _ = "end of CoverTab[194946]"

	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:174
	// _ = "end of CoverTab[194933]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:174
	_go_fuzz_dep_.CoverTab[194934]++

												return document, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:176
	// _ = "end of CoverTab[194934]"

}

func (l *jsonReferenceLoader) loadFromHTTP(address string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:180
	_go_fuzz_dep_.CoverTab[194949]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:184
	if metaSchema := drafts.GetMetaSchema(address); metaSchema != "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:184
		_go_fuzz_dep_.CoverTab[194954]++
													return decodeJSONUsingNumber(strings.NewReader(metaSchema))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:185
		// _ = "end of CoverTab[194954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:186
		_go_fuzz_dep_.CoverTab[194955]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:186
		// _ = "end of CoverTab[194955]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:186
	// _ = "end of CoverTab[194949]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:186
	_go_fuzz_dep_.CoverTab[194950]++

												resp, err := http.Get(address)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:189
		_go_fuzz_dep_.CoverTab[194956]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:190
		// _ = "end of CoverTab[194956]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:191
		_go_fuzz_dep_.CoverTab[194957]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:191
		// _ = "end of CoverTab[194957]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:191
	// _ = "end of CoverTab[194950]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:191
	_go_fuzz_dep_.CoverTab[194951]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:194
	if resp.StatusCode != http.StatusOK {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:194
		_go_fuzz_dep_.CoverTab[194958]++
													return nil, errors.New(formatErrorDescription(Locale.HttpBadStatus(), ErrorDetails{"status": resp.Status}))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:195
		// _ = "end of CoverTab[194958]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:196
		_go_fuzz_dep_.CoverTab[194959]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:196
		// _ = "end of CoverTab[194959]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:196
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:196
	// _ = "end of CoverTab[194951]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:196
	_go_fuzz_dep_.CoverTab[194952]++

												bodyBuff, err := ioutil.ReadAll(resp.Body)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:199
		_go_fuzz_dep_.CoverTab[194960]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:200
		// _ = "end of CoverTab[194960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:201
		_go_fuzz_dep_.CoverTab[194961]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:201
		// _ = "end of CoverTab[194961]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:201
	// _ = "end of CoverTab[194952]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:201
	_go_fuzz_dep_.CoverTab[194953]++

												return decodeJSONUsingNumber(bytes.NewReader(bodyBuff))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:203
	// _ = "end of CoverTab[194953]"
}

func (l *jsonReferenceLoader) loadFromFile(path string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:206
	_go_fuzz_dep_.CoverTab[194962]++
												f, err := l.fs.Open(path)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:208
		_go_fuzz_dep_.CoverTab[194965]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:209
		// _ = "end of CoverTab[194965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:210
		_go_fuzz_dep_.CoverTab[194966]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:210
		// _ = "end of CoverTab[194966]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:210
	// _ = "end of CoverTab[194962]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:210
	_go_fuzz_dep_.CoverTab[194963]++
												defer f.Close()

												bodyBuff, err := ioutil.ReadAll(f)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:214
		_go_fuzz_dep_.CoverTab[194967]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:215
		// _ = "end of CoverTab[194967]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:216
		_go_fuzz_dep_.CoverTab[194968]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:216
		// _ = "end of CoverTab[194968]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:216
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:216
	// _ = "end of CoverTab[194963]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:216
	_go_fuzz_dep_.CoverTab[194964]++

												return decodeJSONUsingNumber(bytes.NewReader(bodyBuff))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:218
	// _ = "end of CoverTab[194964]"

}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:224
type jsonStringLoader struct {
	source string
}

func (l *jsonStringLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:228
	_go_fuzz_dep_.CoverTab[194969]++
												return l.source
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:229
	// _ = "end of CoverTab[194969]"
}

func (l *jsonStringLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:232
	_go_fuzz_dep_.CoverTab[194970]++
												return gojsonreference.NewJsonReference("#")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:233
	// _ = "end of CoverTab[194970]"
}

func (l *jsonStringLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:236
	_go_fuzz_dep_.CoverTab[194971]++
												return &DefaultJSONLoaderFactory{}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:237
	// _ = "end of CoverTab[194971]"
}

// NewStringLoader creates a new JSONLoader, taking a string as source
func NewStringLoader(source string) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:241
	_go_fuzz_dep_.CoverTab[194972]++
												return &jsonStringLoader{source: source}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:242
	// _ = "end of CoverTab[194972]"
}

func (l *jsonStringLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:245
	_go_fuzz_dep_.CoverTab[194973]++

												return decodeJSONUsingNumber(strings.NewReader(l.JsonSource().(string)))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:247
	// _ = "end of CoverTab[194973]"

}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:253
type jsonBytesLoader struct {
	source []byte
}

func (l *jsonBytesLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:257
	_go_fuzz_dep_.CoverTab[194974]++
												return l.source
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:258
	// _ = "end of CoverTab[194974]"
}

func (l *jsonBytesLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:261
	_go_fuzz_dep_.CoverTab[194975]++
												return gojsonreference.NewJsonReference("#")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:262
	// _ = "end of CoverTab[194975]"
}

func (l *jsonBytesLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:265
	_go_fuzz_dep_.CoverTab[194976]++
												return &DefaultJSONLoaderFactory{}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:266
	// _ = "end of CoverTab[194976]"
}

// NewBytesLoader creates a new JSONLoader, taking a `[]byte` as source
func NewBytesLoader(source []byte) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:270
	_go_fuzz_dep_.CoverTab[194977]++
												return &jsonBytesLoader{source: source}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:271
	// _ = "end of CoverTab[194977]"
}

func (l *jsonBytesLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:274
	_go_fuzz_dep_.CoverTab[194978]++
												return decodeJSONUsingNumber(bytes.NewReader(l.JsonSource().([]byte)))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:275
	// _ = "end of CoverTab[194978]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:281
type jsonGoLoader struct {
	source interface{}
}

func (l *jsonGoLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:285
	_go_fuzz_dep_.CoverTab[194979]++
												return l.source
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:286
	// _ = "end of CoverTab[194979]"
}

func (l *jsonGoLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:289
	_go_fuzz_dep_.CoverTab[194980]++
												return gojsonreference.NewJsonReference("#")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:290
	// _ = "end of CoverTab[194980]"
}

func (l *jsonGoLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:293
	_go_fuzz_dep_.CoverTab[194981]++
												return &DefaultJSONLoaderFactory{}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:294
	// _ = "end of CoverTab[194981]"
}

// NewGoLoader creates a new JSONLoader from a given Go struct
func NewGoLoader(source interface{}) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:298
	_go_fuzz_dep_.CoverTab[194982]++
												return &jsonGoLoader{source: source}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:299
	// _ = "end of CoverTab[194982]"
}

func (l *jsonGoLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:302
	_go_fuzz_dep_.CoverTab[194983]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:306
	jsonBytes, err := json.Marshal(l.JsonSource())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:307
		_go_fuzz_dep_.CoverTab[194985]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:308
		// _ = "end of CoverTab[194985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:309
		_go_fuzz_dep_.CoverTab[194986]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:309
		// _ = "end of CoverTab[194986]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:309
	// _ = "end of CoverTab[194983]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:309
	_go_fuzz_dep_.CoverTab[194984]++

												return decodeJSONUsingNumber(bytes.NewReader(jsonBytes))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:311
	// _ = "end of CoverTab[194984]"

}

type jsonIOLoader struct {
	buf *bytes.Buffer
}

// NewReaderLoader creates a new JSON loader using the provided io.Reader
func NewReaderLoader(source io.Reader) (JSONLoader, io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:320
	_go_fuzz_dep_.CoverTab[194987]++
												buf := &bytes.Buffer{}
												return &jsonIOLoader{buf: buf}, io.TeeReader(source, buf)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:322
	// _ = "end of CoverTab[194987]"
}

// NewWriterLoader creates a new JSON loader using the provided io.Writer
func NewWriterLoader(source io.Writer) (JSONLoader, io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:326
	_go_fuzz_dep_.CoverTab[194988]++
												buf := &bytes.Buffer{}
												return &jsonIOLoader{buf: buf}, io.MultiWriter(source, buf)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:328
	// _ = "end of CoverTab[194988]"
}

func (l *jsonIOLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:331
	_go_fuzz_dep_.CoverTab[194989]++
												return l.buf.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:332
	// _ = "end of CoverTab[194989]"
}

func (l *jsonIOLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:335
	_go_fuzz_dep_.CoverTab[194990]++
												return decodeJSONUsingNumber(l.buf)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:336
	// _ = "end of CoverTab[194990]"
}

func (l *jsonIOLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:339
	_go_fuzz_dep_.CoverTab[194991]++
												return gojsonreference.NewJsonReference("#")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:340
	// _ = "end of CoverTab[194991]"
}

func (l *jsonIOLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:343
	_go_fuzz_dep_.CoverTab[194992]++
												return &DefaultJSONLoaderFactory{}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:344
	// _ = "end of CoverTab[194992]"
}

// JSON raw loader
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:347
// In case the JSON is already marshalled to interface{} use this loader
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:347
// This is used for testing as otherwise there is no guarantee the JSON is marshalled
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:347
// "properly" by using https://golang.org/pkg/encoding/json/#Decoder.UseNumber
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:351
type jsonRawLoader struct {
	source interface{}
}

// NewRawLoader creates a new JSON raw loader for the given source
func NewRawLoader(source interface{}) JSONLoader {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:356
	_go_fuzz_dep_.CoverTab[194993]++
												return &jsonRawLoader{source: source}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:357
	// _ = "end of CoverTab[194993]"
}
func (l *jsonRawLoader) JsonSource() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:359
	_go_fuzz_dep_.CoverTab[194994]++
												return l.source
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:360
	// _ = "end of CoverTab[194994]"
}
func (l *jsonRawLoader) LoadJSON() (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:362
	_go_fuzz_dep_.CoverTab[194995]++
												return l.source, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:363
	// _ = "end of CoverTab[194995]"
}
func (l *jsonRawLoader) JsonReference() (gojsonreference.JsonReference, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:365
	_go_fuzz_dep_.CoverTab[194996]++
												return gojsonreference.NewJsonReference("#")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:366
	// _ = "end of CoverTab[194996]"
}
func (l *jsonRawLoader) LoaderFactory() JSONLoaderFactory {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:368
	_go_fuzz_dep_.CoverTab[194997]++
												return &DefaultJSONLoaderFactory{}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:369
	// _ = "end of CoverTab[194997]"
}

func decodeJSONUsingNumber(r io.Reader) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:372
	_go_fuzz_dep_.CoverTab[194998]++

												var document interface{}

												decoder := json.NewDecoder(r)
												decoder.UseNumber()

												err := decoder.Decode(&document)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:380
		_go_fuzz_dep_.CoverTab[195000]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:381
		// _ = "end of CoverTab[195000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:382
		_go_fuzz_dep_.CoverTab[195001]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:382
		// _ = "end of CoverTab[195001]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:382
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:382
	// _ = "end of CoverTab[194998]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:382
	_go_fuzz_dep_.CoverTab[194999]++

												return document, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:384
	// _ = "end of CoverTab[194999]"

}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:386
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonLoader.go:386
var _ = _go_fuzz_dep_.CoverTab
