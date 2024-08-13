//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:1
)

import (
	"os"
	"syscall"
	"time"
)

var _ Lstater = (*ReadOnlyFs)(nil)

type ReadOnlyFs struct {
	source Fs
}

func NewReadOnlyFs(source Fs) Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:15
	_go_fuzz_dep_.CoverTab[118059]++
											return &ReadOnlyFs{source: source}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:16
	// _ = "end of CoverTab[118059]"
}

func (r *ReadOnlyFs) ReadDir(name string) ([]os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:19
	_go_fuzz_dep_.CoverTab[118060]++
											return ReadDir(r.source, name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:20
	// _ = "end of CoverTab[118060]"
}

func (r *ReadOnlyFs) Chtimes(n string, a, m time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:23
	_go_fuzz_dep_.CoverTab[118061]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:24
	// _ = "end of CoverTab[118061]"
}

func (r *ReadOnlyFs) Chmod(n string, m os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:27
	_go_fuzz_dep_.CoverTab[118062]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:28
	// _ = "end of CoverTab[118062]"
}

func (r *ReadOnlyFs) Chown(n string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:31
	_go_fuzz_dep_.CoverTab[118063]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:32
	// _ = "end of CoverTab[118063]"
}

func (r *ReadOnlyFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:35
	_go_fuzz_dep_.CoverTab[118064]++
											return "ReadOnlyFilter"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:36
	// _ = "end of CoverTab[118064]"
}

func (r *ReadOnlyFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:39
	_go_fuzz_dep_.CoverTab[118065]++
											return r.source.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:40
	// _ = "end of CoverTab[118065]"
}

func (r *ReadOnlyFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:43
	_go_fuzz_dep_.CoverTab[118066]++
											if lsf, ok := r.source.(Lstater); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:44
		_go_fuzz_dep_.CoverTab[118068]++
												return lsf.LstatIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:45
		// _ = "end of CoverTab[118068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:46
		_go_fuzz_dep_.CoverTab[118069]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:46
		// _ = "end of CoverTab[118069]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:46
	// _ = "end of CoverTab[118066]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:46
	_go_fuzz_dep_.CoverTab[118067]++
											fi, err := r.Stat(name)
											return fi, false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:48
	// _ = "end of CoverTab[118067]"
}

func (r *ReadOnlyFs) SymlinkIfPossible(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:51
	_go_fuzz_dep_.CoverTab[118070]++
											return &os.LinkError{Op: "symlink", Old: oldname, New: newname, Err: ErrNoSymlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:52
	// _ = "end of CoverTab[118070]"
}

func (r *ReadOnlyFs) ReadlinkIfPossible(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:55
	_go_fuzz_dep_.CoverTab[118071]++
											if srdr, ok := r.source.(LinkReader); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:56
		_go_fuzz_dep_.CoverTab[118073]++
												return srdr.ReadlinkIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:57
		// _ = "end of CoverTab[118073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:58
		_go_fuzz_dep_.CoverTab[118074]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:58
		// _ = "end of CoverTab[118074]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:58
	// _ = "end of CoverTab[118071]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:58
	_go_fuzz_dep_.CoverTab[118072]++

											return "", &os.PathError{Op: "readlink", Path: name, Err: ErrNoReadlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:60
	// _ = "end of CoverTab[118072]"
}

func (r *ReadOnlyFs) Rename(o, n string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:63
	_go_fuzz_dep_.CoverTab[118075]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:64
	// _ = "end of CoverTab[118075]"
}

func (r *ReadOnlyFs) RemoveAll(p string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:67
	_go_fuzz_dep_.CoverTab[118076]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:68
	// _ = "end of CoverTab[118076]"
}

func (r *ReadOnlyFs) Remove(n string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:71
	_go_fuzz_dep_.CoverTab[118077]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:72
	// _ = "end of CoverTab[118077]"
}

func (r *ReadOnlyFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:75
	_go_fuzz_dep_.CoverTab[118078]++
											if flag&(os.O_WRONLY|syscall.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:76
		_go_fuzz_dep_.CoverTab[118080]++
												return nil, syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:77
		// _ = "end of CoverTab[118080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:78
		_go_fuzz_dep_.CoverTab[118081]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:78
		// _ = "end of CoverTab[118081]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:78
	// _ = "end of CoverTab[118078]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:78
	_go_fuzz_dep_.CoverTab[118079]++
											return r.source.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:79
	// _ = "end of CoverTab[118079]"
}

func (r *ReadOnlyFs) Open(n string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:82
	_go_fuzz_dep_.CoverTab[118082]++
											return r.source.Open(n)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:83
	// _ = "end of CoverTab[118082]"
}

func (r *ReadOnlyFs) Mkdir(n string, p os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:86
	_go_fuzz_dep_.CoverTab[118083]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:87
	// _ = "end of CoverTab[118083]"
}

func (r *ReadOnlyFs) MkdirAll(n string, p os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:90
	_go_fuzz_dep_.CoverTab[118084]++
											return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:91
	// _ = "end of CoverTab[118084]"
}

func (r *ReadOnlyFs) Create(n string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:94
	_go_fuzz_dep_.CoverTab[118085]++
											return nil, syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:95
	// _ = "end of CoverTab[118085]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:96
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/readonlyfs.go:96
var _ = _go_fuzz_dep_.CoverTab
