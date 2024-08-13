// Copyright ©2015 The Go Authors
// Copyright ©2015 Steve Francia <spf@spf13.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:16
)

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// byName implements sort.Interface.
type byName []os.FileInfo

func (f byName) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:33
	_go_fuzz_dep_.CoverTab[117735]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:33
	return len(f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:33
	// _ = "end of CoverTab[117735]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:33
}
func (f byName) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:34
	_go_fuzz_dep_.CoverTab[117736]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:34
	return f[i].Name() < f[j].Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:34
	// _ = "end of CoverTab[117736]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:34
}
func (f byName) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:35
	_go_fuzz_dep_.CoverTab[117737]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:35
	f[i], f[j] = f[j], f[i]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:35
	// _ = "end of CoverTab[117737]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:35
}

// ReadDir reads the directory named by dirname and returns
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:37
// a list of sorted directory entries.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:39
func (a Afero) ReadDir(dirname string) ([]os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:39
	_go_fuzz_dep_.CoverTab[117738]++
										return ReadDir(a.Fs, dirname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:40
	// _ = "end of CoverTab[117738]"
}

func ReadDir(fs Fs, dirname string) ([]os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:43
	_go_fuzz_dep_.CoverTab[117739]++
										f, err := fs.Open(dirname)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:45
		_go_fuzz_dep_.CoverTab[117742]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:46
		// _ = "end of CoverTab[117742]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:47
		_go_fuzz_dep_.CoverTab[117743]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:47
		// _ = "end of CoverTab[117743]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:47
	// _ = "end of CoverTab[117739]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:47
	_go_fuzz_dep_.CoverTab[117740]++
										list, err := f.Readdir(-1)
										f.Close()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:50
		_go_fuzz_dep_.CoverTab[117744]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:51
		// _ = "end of CoverTab[117744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:52
		_go_fuzz_dep_.CoverTab[117745]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:52
		// _ = "end of CoverTab[117745]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:52
	// _ = "end of CoverTab[117740]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:52
	_go_fuzz_dep_.CoverTab[117741]++
										sort.Sort(byName(list))
										return list, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:54
	// _ = "end of CoverTab[117741]"
}

// ReadFile reads the file named by filename and returns the contents.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:57
// A successful call returns err == nil, not err == EOF. Because ReadFile
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:57
// reads the whole file, it does not treat an EOF from Read as an error
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:57
// to be reported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:61
func (a Afero) ReadFile(filename string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:61
	_go_fuzz_dep_.CoverTab[117746]++
										return ReadFile(a.Fs, filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:62
	// _ = "end of CoverTab[117746]"
}

func ReadFile(fs Fs, filename string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:65
	_go_fuzz_dep_.CoverTab[117747]++
										f, err := fs.Open(filename)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:67
		_go_fuzz_dep_.CoverTab[117750]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:68
		// _ = "end of CoverTab[117750]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:69
		_go_fuzz_dep_.CoverTab[117751]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:69
		// _ = "end of CoverTab[117751]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:69
	// _ = "end of CoverTab[117747]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:69
	_go_fuzz_dep_.CoverTab[117748]++
										defer f.Close()
	// It's a good but not certain bet that FileInfo will tell us exactly how much to
	// read, so let's try it but be prepared for the answer to be wrong.
	var n int64

	if fi, err := f.Stat(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:75
		_go_fuzz_dep_.CoverTab[117752]++

											if size := fi.Size(); size < 1e9 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:77
			_go_fuzz_dep_.CoverTab[117753]++
												n = size
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:78
			// _ = "end of CoverTab[117753]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:79
			_go_fuzz_dep_.CoverTab[117754]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:79
			// _ = "end of CoverTab[117754]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:79
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:79
		// _ = "end of CoverTab[117752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:80
		_go_fuzz_dep_.CoverTab[117755]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:80
		// _ = "end of CoverTab[117755]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:80
	// _ = "end of CoverTab[117748]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:80
	_go_fuzz_dep_.CoverTab[117749]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:86
	return readAll(f, n+bytes.MinRead)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:86
	// _ = "end of CoverTab[117749]"
}

// readAll reads from r until an error or EOF and returns the data it read
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:89
// from the internal buffer allocated with a specified capacity.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:91
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:91
	_go_fuzz_dep_.CoverTab[117756]++
										buf := bytes.NewBuffer(make([]byte, 0, capacity))

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:95
	defer func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:95
		_go_fuzz_dep_.CoverTab[117758]++
											e := recover()
											if e == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:97
			_go_fuzz_dep_.CoverTab[117760]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:98
			// _ = "end of CoverTab[117760]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:99
			_go_fuzz_dep_.CoverTab[117761]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:99
			// _ = "end of CoverTab[117761]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:99
		// _ = "end of CoverTab[117758]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:99
		_go_fuzz_dep_.CoverTab[117759]++
											if panicErr, ok := e.(error); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:100
			_go_fuzz_dep_.CoverTab[117762]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:100
			return panicErr == bytes.ErrTooLarge
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:100
			// _ = "end of CoverTab[117762]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:100
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:100
			_go_fuzz_dep_.CoverTab[117763]++
												err = panicErr
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:101
			// _ = "end of CoverTab[117763]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:102
			_go_fuzz_dep_.CoverTab[117764]++
												panic(e)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:103
			// _ = "end of CoverTab[117764]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:104
		// _ = "end of CoverTab[117759]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:105
	// _ = "end of CoverTab[117756]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:105
	_go_fuzz_dep_.CoverTab[117757]++
										_, err = buf.ReadFrom(r)
										return buf.Bytes(), err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:107
	// _ = "end of CoverTab[117757]"
}

// ReadAll reads from r until an error or EOF and returns the data it read.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:110
// A successful call returns err == nil, not err == EOF. Because ReadAll is
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:110
// defined to read from src until EOF, it does not treat an EOF from Read
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:110
// as an error to be reported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:114
func ReadAll(r io.Reader) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:114
	_go_fuzz_dep_.CoverTab[117765]++
										return readAll(r, bytes.MinRead)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:115
	// _ = "end of CoverTab[117765]"
}

// WriteFile writes data to a file named by filename.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:118
// If the file does not exist, WriteFile creates it with permissions perm;
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:118
// otherwise WriteFile truncates it before writing.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:121
func (a Afero) WriteFile(filename string, data []byte, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:121
	_go_fuzz_dep_.CoverTab[117766]++
										return WriteFile(a.Fs, filename, data, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:122
	// _ = "end of CoverTab[117766]"
}

func WriteFile(fs Fs, filename string, data []byte, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:125
	_go_fuzz_dep_.CoverTab[117767]++
										f, err := fs.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:127
		_go_fuzz_dep_.CoverTab[117771]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:128
		// _ = "end of CoverTab[117771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:129
		_go_fuzz_dep_.CoverTab[117772]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:129
		// _ = "end of CoverTab[117772]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:129
	// _ = "end of CoverTab[117767]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:129
	_go_fuzz_dep_.CoverTab[117768]++
										n, err := f.Write(data)
										if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:131
		_go_fuzz_dep_.CoverTab[117773]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:131
		return n < len(data)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:131
		// _ = "end of CoverTab[117773]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:131
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:131
		_go_fuzz_dep_.CoverTab[117774]++
											err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:132
		// _ = "end of CoverTab[117774]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:133
		_go_fuzz_dep_.CoverTab[117775]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:133
		// _ = "end of CoverTab[117775]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:133
	// _ = "end of CoverTab[117768]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:133
	_go_fuzz_dep_.CoverTab[117769]++
										if err1 := f.Close(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:134
		_go_fuzz_dep_.CoverTab[117776]++
											err = err1
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:135
		// _ = "end of CoverTab[117776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:136
		_go_fuzz_dep_.CoverTab[117777]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:136
		// _ = "end of CoverTab[117777]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:136
	// _ = "end of CoverTab[117769]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:136
	_go_fuzz_dep_.CoverTab[117770]++
										return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:137
	// _ = "end of CoverTab[117770]"
}

// Random number state.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:140
// We generate random temporary file names so that there's a good
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:140
// chance the file doesn't exist yet - keeps the number of tries in
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:140
// TempFile to a minimum.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:144
var randNum uint32
var randmu sync.Mutex

func reseed() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:147
	_go_fuzz_dep_.CoverTab[117778]++
										return uint32(time.Now().UnixNano() + int64(os.Getpid()))
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:148
	// _ = "end of CoverTab[117778]"
}

func nextRandom() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:151
	_go_fuzz_dep_.CoverTab[117779]++
										randmu.Lock()
										r := randNum
										if r == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:154
		_go_fuzz_dep_.CoverTab[117781]++
											r = reseed()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:155
		// _ = "end of CoverTab[117781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:156
		_go_fuzz_dep_.CoverTab[117782]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:156
		// _ = "end of CoverTab[117782]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:156
	// _ = "end of CoverTab[117779]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:156
	_go_fuzz_dep_.CoverTab[117780]++
										r = r*1664525 + 1013904223
										randNum = r
										randmu.Unlock()
										return strconv.Itoa(int(1e9 + r%1e9))[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:160
	// _ = "end of CoverTab[117780]"
}

// TempFile creates a new temporary file in the directory dir,
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// opens the file for reading and writing, and returns the resulting *os.File.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// The filename is generated by taking pattern and adding a random
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// string to the end. If pattern includes a "*", the random string
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// replaces the last "*".
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// If dir is the empty string, TempFile uses the default directory
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// for temporary files (see os.TempDir).
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// Multiple programs calling TempFile simultaneously
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// will not choose the same file. The caller can use f.Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// to find the pathname of the file. It is the caller's responsibility
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:163
// to remove the file when no longer needed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:174
func (a Afero) TempFile(dir, pattern string) (f File, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:174
	_go_fuzz_dep_.CoverTab[117783]++
										return TempFile(a.Fs, dir, pattern)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:175
	// _ = "end of CoverTab[117783]"
}

func TempFile(fs Fs, dir, pattern string) (f File, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:178
	_go_fuzz_dep_.CoverTab[117784]++
										if dir == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:179
		_go_fuzz_dep_.CoverTab[117788]++
											dir = os.TempDir()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:180
		// _ = "end of CoverTab[117788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:181
		_go_fuzz_dep_.CoverTab[117789]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:181
		// _ = "end of CoverTab[117789]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:181
	// _ = "end of CoverTab[117784]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:181
	_go_fuzz_dep_.CoverTab[117785]++

										var prefix, suffix string
										if pos := strings.LastIndex(pattern, "*"); pos != -1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:184
		_go_fuzz_dep_.CoverTab[117790]++
											prefix, suffix = pattern[:pos], pattern[pos+1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:185
		// _ = "end of CoverTab[117790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:186
		_go_fuzz_dep_.CoverTab[117791]++
											prefix = pattern
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:187
		// _ = "end of CoverTab[117791]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:188
	// _ = "end of CoverTab[117785]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:188
	_go_fuzz_dep_.CoverTab[117786]++

										nconflict := 0
										for i := 0; i < 10000; i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:191
		_go_fuzz_dep_.CoverTab[117792]++
											name := filepath.Join(dir, prefix+nextRandom()+suffix)
											f, err = fs.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
											if os.IsExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:194
			_go_fuzz_dep_.CoverTab[117794]++
												if nconflict++; nconflict > 10 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:195
				_go_fuzz_dep_.CoverTab[117796]++
													randmu.Lock()
													randNum = reseed()
													randmu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:198
				// _ = "end of CoverTab[117796]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:199
				_go_fuzz_dep_.CoverTab[117797]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:199
				// _ = "end of CoverTab[117797]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:199
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:199
			// _ = "end of CoverTab[117794]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:199
			_go_fuzz_dep_.CoverTab[117795]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:200
			// _ = "end of CoverTab[117795]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:201
			_go_fuzz_dep_.CoverTab[117798]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:201
			// _ = "end of CoverTab[117798]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:201
		// _ = "end of CoverTab[117792]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:201
		_go_fuzz_dep_.CoverTab[117793]++
											break
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:202
		// _ = "end of CoverTab[117793]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:203
	// _ = "end of CoverTab[117786]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:203
	_go_fuzz_dep_.CoverTab[117787]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:204
	// _ = "end of CoverTab[117787]"
}

// TempDir creates a new temporary directory in the directory dir
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// with a name beginning with prefix and returns the path of the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// new directory.  If dir is the empty string, TempDir uses the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// default directory for temporary files (see os.TempDir).
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// Multiple programs calling TempDir simultaneously
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// will not choose the same directory.  It is the caller's responsibility
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:207
// to remove the directory when no longer needed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:214
func (a Afero) TempDir(dir, prefix string) (name string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:214
	_go_fuzz_dep_.CoverTab[117799]++
										return TempDir(a.Fs, dir, prefix)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:215
	// _ = "end of CoverTab[117799]"
}
func TempDir(fs Fs, dir, prefix string) (name string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:217
	_go_fuzz_dep_.CoverTab[117800]++
										if dir == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:218
		_go_fuzz_dep_.CoverTab[117803]++
											dir = os.TempDir()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:219
		// _ = "end of CoverTab[117803]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:220
		_go_fuzz_dep_.CoverTab[117804]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:220
		// _ = "end of CoverTab[117804]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:220
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:220
	// _ = "end of CoverTab[117800]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:220
	_go_fuzz_dep_.CoverTab[117801]++

										nconflict := 0
										for i := 0; i < 10000; i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:223
		_go_fuzz_dep_.CoverTab[117805]++
											try := filepath.Join(dir, prefix+nextRandom())
											err = fs.Mkdir(try, 0700)
											if os.IsExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:226
			_go_fuzz_dep_.CoverTab[117808]++
												if nconflict++; nconflict > 10 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:227
				_go_fuzz_dep_.CoverTab[117810]++
													randmu.Lock()
													randNum = reseed()
													randmu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:230
				// _ = "end of CoverTab[117810]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:231
				_go_fuzz_dep_.CoverTab[117811]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:231
				// _ = "end of CoverTab[117811]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:231
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:231
			// _ = "end of CoverTab[117808]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:231
			_go_fuzz_dep_.CoverTab[117809]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:232
			// _ = "end of CoverTab[117809]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:233
			_go_fuzz_dep_.CoverTab[117812]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:233
			// _ = "end of CoverTab[117812]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:233
		// _ = "end of CoverTab[117805]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:233
		_go_fuzz_dep_.CoverTab[117806]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:234
			_go_fuzz_dep_.CoverTab[117813]++
												name = try
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:235
			// _ = "end of CoverTab[117813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:236
			_go_fuzz_dep_.CoverTab[117814]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:236
			// _ = "end of CoverTab[117814]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:236
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:236
		// _ = "end of CoverTab[117806]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:236
		_go_fuzz_dep_.CoverTab[117807]++
											break
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:237
		// _ = "end of CoverTab[117807]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:238
	// _ = "end of CoverTab[117801]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:238
	_go_fuzz_dep_.CoverTab[117802]++
										return
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:239
	// _ = "end of CoverTab[117802]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:240
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/ioutil.go:240
var _ = _go_fuzz_dep_.CoverTab
