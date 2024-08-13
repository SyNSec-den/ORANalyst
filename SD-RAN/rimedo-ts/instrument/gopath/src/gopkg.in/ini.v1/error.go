// Copyright 2016 Unknwon
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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:15
)

import (
	"fmt"
)

// ErrDelimiterNotFound indicates the error type of no delimiter is found which there should be one.
type ErrDelimiterNotFound struct {
	Line string
}

// IsErrDelimiterNotFound returns true if the given error is an instance of ErrDelimiterNotFound.
func IsErrDelimiterNotFound(err error) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:27
	_go_fuzz_dep_.CoverTab[128245]++
									_, ok := err.(ErrDelimiterNotFound)
									return ok
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:29
	// _ = "end of CoverTab[128245]"
}

func (err ErrDelimiterNotFound) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:32
	_go_fuzz_dep_.CoverTab[128246]++
									return fmt.Sprintf("key-value delimiter not found: %s", err.Line)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:33
	// _ = "end of CoverTab[128246]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/error.go:34
var _ = _go_fuzz_dep_.CoverTab
