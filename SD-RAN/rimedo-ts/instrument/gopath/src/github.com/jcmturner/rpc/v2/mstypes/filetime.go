//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:1
// Package mstypes implements representations of Microsoft types
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:2
)

import (
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:18
const unixEpochDiff = 116444736000000000

// FileTime implements the Microsoft FILETIME type https://msdn.microsoft.com/en-us/library/cc230324.aspx
type FileTime struct {
	LowDateTime	uint32
	HighDateTime	uint32
}

// Time return a golang Time type from the FileTime
func (ft FileTime) Time() time.Time {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:27
	_go_fuzz_dep_.CoverTab[87355]++
												ns := (ft.MSEpoch() - unixEpochDiff) * 100
												return time.Unix(0, int64(ns)).UTC()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:29
	// _ = "end of CoverTab[87355]"
}

// MSEpoch returns the FileTime as a Microsoft epoch, the number of 100 nano second periods elapsed from January 1, 1601 UTC.
func (ft FileTime) MSEpoch() int64 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:33
	_go_fuzz_dep_.CoverTab[87356]++
												return (int64(ft.HighDateTime) << 32) + int64(ft.LowDateTime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:34
	// _ = "end of CoverTab[87356]"
}

// Unix returns the FileTime as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.
func (ft FileTime) Unix() int64 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:38
	_go_fuzz_dep_.CoverTab[87357]++
												return (ft.MSEpoch() - unixEpochDiff) / 10000000
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:39
	// _ = "end of CoverTab[87357]"
}

// GetFileTime returns a FileTime type from the provided Golang Time type.
func GetFileTime(t time.Time) FileTime {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:43
	_go_fuzz_dep_.CoverTab[87358]++
												ns := t.UnixNano()
												fp := (ns / 100) + unixEpochDiff
												hd := fp >> 32
												ld := fp - (hd << 32)
												return FileTime{
		LowDateTime:	uint32(ld),
		HighDateTime:	uint32(hd),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:51
	// _ = "end of CoverTab[87358]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/filetime.go:52
var _ = _go_fuzz_dep_.CoverTab
