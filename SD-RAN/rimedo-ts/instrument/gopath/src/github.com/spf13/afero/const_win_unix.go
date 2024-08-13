// Copyright © 2016 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build !darwin && !openbsd && !freebsd && !dragonfly && !netbsd && !aix
// +build !darwin,!openbsd,!freebsd,!dragonfly,!netbsd,!aix

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:16
)

import (
	"syscall"
)

const BADFD = syscall.EBADFD

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:22
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/const_win_unix.go:22
var _ = _go_fuzz_dep_.CoverTab
