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

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
package ini

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:15
)

func inSlice(str string, s []string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:17
	_go_fuzz_dep_.CoverTab[128496]++
									for _, v := range s {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:18
		_go_fuzz_dep_.CoverTab[128498]++
										if str == v {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:19
			_go_fuzz_dep_.CoverTab[128499]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:20
			// _ = "end of CoverTab[128499]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:21
			_go_fuzz_dep_.CoverTab[128500]++
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:21
			// _ = "end of CoverTab[128500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:21
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:21
		// _ = "end of CoverTab[128498]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:22
	// _ = "end of CoverTab[128496]"
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:22
	_go_fuzz_dep_.CoverTab[128497]++
									return false
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:23
	// _ = "end of CoverTab[128497]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/ini.v1@v1.63.2/helper.go:24
var _ = _go_fuzz_dep_.CoverTab
