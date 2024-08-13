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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:14
)

import (
	"os"
)

// Lstater is an optional interface in Afero. It is only implemented by the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:20
// filesystems saying so.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:20
// It will call Lstat if the filesystem iself is, or it delegates to, the os filesystem.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:20
// Else it will call Stat.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:20
// In addtion to the FileInfo, it will return a boolean telling whether Lstat was called or not.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:25
type Lstater interface {
	LstatIfPossible(name string) (os.FileInfo, bool, error)
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/lstater.go:27
var _ = _go_fuzz_dep_.CoverTab
