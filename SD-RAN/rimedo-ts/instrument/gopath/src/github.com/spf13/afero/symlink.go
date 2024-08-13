// Copyright © 2018 Steve Francia <spf@spf13.com>.
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:14
)

import (
	"errors"
)

// Symlinker is an optional interface in Afero. It is only implemented by the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
// filesystems saying so.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
// It indicates support for 3 symlink related interfaces that implement the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
// behaviors of the os methods:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
//   - Lstat
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
//   - Symlink, and
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:20
//   - Readlink
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:27
type Symlinker interface {
	Lstater
	Linker
	LinkReader
}

// Linker is an optional interface in Afero. It is only implemented by the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:33
// filesystems saying so.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:33
// It will call Symlink if the filesystem itself is, or it delegates to, the os filesystem,
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:33
// or the filesystem otherwise supports Symlink's.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:37
type Linker interface {
	SymlinkIfPossible(oldname, newname string) error
}

// ErrNoSymlink is the error that will be wrapped in an os.LinkError if a file system
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:41
// does not support Symlink's either directly or through its delegated filesystem.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:41
// As expressed by support for the Linker interface.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:44
var ErrNoSymlink = errors.New("symlink not supported")

// LinkReader is an optional interface in Afero. It is only implemented by the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:46
// filesystems saying so.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:48
type LinkReader interface {
	ReadlinkIfPossible(name string) (string, error)
}

// ErrNoReadlink is the error that will be wrapped in an os.Path if a file system
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:52
// does not support the readlink operation either directly or through its delegated filesystem.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:52
// As expressed by support for the LinkReader interface.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:55
var ErrNoReadlink = errors.New("readlink not supported")
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:55
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/symlink.go:55
var _ = _go_fuzz_dep_.CoverTab
