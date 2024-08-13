// Copyright © 2022 Steve Francia <spf@spf13.com>.
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

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
package common

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:14
)

import "io/fs"

// FileInfoDirEntry provides an adapter from os.FileInfo to fs.DirEntry
type FileInfoDirEntry struct {
	fs.FileInfo
}

var _ fs.DirEntry = FileInfoDirEntry{}

func (d FileInfoDirEntry) Type() fs.FileMode {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:25
	_go_fuzz_dep_.CoverTab[116793]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:25
	return d.FileInfo.Mode().Type()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:25
	// _ = "end of CoverTab[116793]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:25
}

func (d FileInfoDirEntry) Info() (fs.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
	_go_fuzz_dep_.CoverTab[116794]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
	return d.FileInfo, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
	// _ = "end of CoverTab[116794]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/internal/common/adapters.go:27
var _ = _go_fuzz_dep_.CoverTab
