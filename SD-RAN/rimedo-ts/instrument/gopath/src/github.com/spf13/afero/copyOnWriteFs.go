//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:1
)

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

var _ Lstater = (*CopyOnWriteFs)(nil)

// The CopyOnWriteFs is a union filesystem: a read only base file system with
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
// a possibly writeable layer on top. Changes to the file system will only
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
// be made in the overlay: Changing an existing file in the base layer which
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
// is not present in the overlay will copy the file to the overlay ("changing"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
// includes also calls to e.g. Chtimes(), Chmod() and Chown()).
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:13
// Reading directories is currently only supported via Open(), not OpenFile().
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:20
type CopyOnWriteFs struct {
	base	Fs
	layer	Fs
}

func NewCopyOnWriteFs(base Fs, layer Fs) Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:25
	_go_fuzz_dep_.CoverTab[117422]++
											return &CopyOnWriteFs{base: base, layer: layer}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:26
	// _ = "end of CoverTab[117422]"
}

// Returns true if the file is not in the overlay
func (u *CopyOnWriteFs) isBaseFile(name string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:30
	_go_fuzz_dep_.CoverTab[117423]++
											if _, err := u.layer.Stat(name); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:31
		_go_fuzz_dep_.CoverTab[117426]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:32
		// _ = "end of CoverTab[117426]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:33
		_go_fuzz_dep_.CoverTab[117427]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:33
		// _ = "end of CoverTab[117427]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:33
	// _ = "end of CoverTab[117423]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:33
	_go_fuzz_dep_.CoverTab[117424]++
											_, err := u.base.Stat(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:35
		_go_fuzz_dep_.CoverTab[117428]++
												if oerr, ok := err.(*os.PathError); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:36
			_go_fuzz_dep_.CoverTab[117430]++
													if oerr.Err == os.ErrNotExist || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				_go_fuzz_dep_.CoverTab[117431]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				return oerr.Err == syscall.ENOENT
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				// _ = "end of CoverTab[117431]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				_go_fuzz_dep_.CoverTab[117432]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				return oerr.Err == syscall.ENOTDIR
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				// _ = "end of CoverTab[117432]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
			}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:37
				_go_fuzz_dep_.CoverTab[117433]++
														return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:38
				// _ = "end of CoverTab[117433]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:39
				_go_fuzz_dep_.CoverTab[117434]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:39
				// _ = "end of CoverTab[117434]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:39
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:39
			// _ = "end of CoverTab[117430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:40
			_go_fuzz_dep_.CoverTab[117435]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:40
			// _ = "end of CoverTab[117435]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:40
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:40
		// _ = "end of CoverTab[117428]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:40
		_go_fuzz_dep_.CoverTab[117429]++
												if err == syscall.ENOENT {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:41
			_go_fuzz_dep_.CoverTab[117436]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:42
			// _ = "end of CoverTab[117436]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:43
			_go_fuzz_dep_.CoverTab[117437]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:43
			// _ = "end of CoverTab[117437]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:43
		// _ = "end of CoverTab[117429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:44
		_go_fuzz_dep_.CoverTab[117438]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:44
		// _ = "end of CoverTab[117438]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:44
	// _ = "end of CoverTab[117424]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:44
	_go_fuzz_dep_.CoverTab[117425]++
											return true, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:45
	// _ = "end of CoverTab[117425]"
}

func (u *CopyOnWriteFs) copyToLayer(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:48
	_go_fuzz_dep_.CoverTab[117439]++
											return copyToLayer(u.base, u.layer, name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:49
	// _ = "end of CoverTab[117439]"
}

func (u *CopyOnWriteFs) Chtimes(name string, atime, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:52
	_go_fuzz_dep_.CoverTab[117440]++
											b, err := u.isBaseFile(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:54
		_go_fuzz_dep_.CoverTab[117443]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:55
		// _ = "end of CoverTab[117443]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:56
		_go_fuzz_dep_.CoverTab[117444]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:56
		// _ = "end of CoverTab[117444]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:56
	// _ = "end of CoverTab[117440]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:56
	_go_fuzz_dep_.CoverTab[117441]++
											if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:57
		_go_fuzz_dep_.CoverTab[117445]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:58
			_go_fuzz_dep_.CoverTab[117446]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:59
			// _ = "end of CoverTab[117446]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:60
			_go_fuzz_dep_.CoverTab[117447]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:60
			// _ = "end of CoverTab[117447]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:60
		// _ = "end of CoverTab[117445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:61
		_go_fuzz_dep_.CoverTab[117448]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:61
		// _ = "end of CoverTab[117448]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:61
	// _ = "end of CoverTab[117441]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:61
	_go_fuzz_dep_.CoverTab[117442]++
											return u.layer.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:62
	// _ = "end of CoverTab[117442]"
}

func (u *CopyOnWriteFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:65
	_go_fuzz_dep_.CoverTab[117449]++
											b, err := u.isBaseFile(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:67
		_go_fuzz_dep_.CoverTab[117452]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:68
		// _ = "end of CoverTab[117452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:69
		_go_fuzz_dep_.CoverTab[117453]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:69
		// _ = "end of CoverTab[117453]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:69
	// _ = "end of CoverTab[117449]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:69
	_go_fuzz_dep_.CoverTab[117450]++
											if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:70
		_go_fuzz_dep_.CoverTab[117454]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:71
			_go_fuzz_dep_.CoverTab[117455]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:72
			// _ = "end of CoverTab[117455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:73
			_go_fuzz_dep_.CoverTab[117456]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:73
			// _ = "end of CoverTab[117456]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:73
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:73
		// _ = "end of CoverTab[117454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:74
		_go_fuzz_dep_.CoverTab[117457]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:74
		// _ = "end of CoverTab[117457]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:74
	// _ = "end of CoverTab[117450]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:74
	_go_fuzz_dep_.CoverTab[117451]++
											return u.layer.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:75
	// _ = "end of CoverTab[117451]"
}

func (u *CopyOnWriteFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:78
	_go_fuzz_dep_.CoverTab[117458]++
											b, err := u.isBaseFile(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:80
		_go_fuzz_dep_.CoverTab[117461]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:81
		// _ = "end of CoverTab[117461]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:82
		_go_fuzz_dep_.CoverTab[117462]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:82
		// _ = "end of CoverTab[117462]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:82
	// _ = "end of CoverTab[117458]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:82
	_go_fuzz_dep_.CoverTab[117459]++
											if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:83
		_go_fuzz_dep_.CoverTab[117463]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:84
			_go_fuzz_dep_.CoverTab[117464]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:85
			// _ = "end of CoverTab[117464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:86
			_go_fuzz_dep_.CoverTab[117465]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:86
			// _ = "end of CoverTab[117465]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:86
		// _ = "end of CoverTab[117463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:87
		_go_fuzz_dep_.CoverTab[117466]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:87
		// _ = "end of CoverTab[117466]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:87
	// _ = "end of CoverTab[117459]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:87
	_go_fuzz_dep_.CoverTab[117460]++
											return u.layer.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:88
	// _ = "end of CoverTab[117460]"
}

func (u *CopyOnWriteFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:91
	_go_fuzz_dep_.CoverTab[117467]++
											fi, err := u.layer.Stat(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:93
		_go_fuzz_dep_.CoverTab[117469]++
												isNotExist := u.isNotExist(err)
												if isNotExist {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:95
			_go_fuzz_dep_.CoverTab[117471]++
													return u.base.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:96
			// _ = "end of CoverTab[117471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:97
			_go_fuzz_dep_.CoverTab[117472]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:97
			// _ = "end of CoverTab[117472]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:97
		// _ = "end of CoverTab[117469]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:97
		_go_fuzz_dep_.CoverTab[117470]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:98
		// _ = "end of CoverTab[117470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:99
		_go_fuzz_dep_.CoverTab[117473]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:99
		// _ = "end of CoverTab[117473]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:99
	// _ = "end of CoverTab[117467]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:99
	_go_fuzz_dep_.CoverTab[117468]++
											return fi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:100
	// _ = "end of CoverTab[117468]"
}

func (u *CopyOnWriteFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:103
	_go_fuzz_dep_.CoverTab[117474]++
											llayer, ok1 := u.layer.(Lstater)
											lbase, ok2 := u.base.(Lstater)

											if ok1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:107
		_go_fuzz_dep_.CoverTab[117477]++
												fi, b, err := llayer.LstatIfPossible(name)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:109
			_go_fuzz_dep_.CoverTab[117479]++
													return fi, b, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:110
			// _ = "end of CoverTab[117479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:111
			_go_fuzz_dep_.CoverTab[117480]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:111
			// _ = "end of CoverTab[117480]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:111
		// _ = "end of CoverTab[117477]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:111
		_go_fuzz_dep_.CoverTab[117478]++

												if !u.isNotExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:113
			_go_fuzz_dep_.CoverTab[117481]++
													return nil, b, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:114
			// _ = "end of CoverTab[117481]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:115
			_go_fuzz_dep_.CoverTab[117482]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:115
			// _ = "end of CoverTab[117482]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:115
		// _ = "end of CoverTab[117478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:116
		_go_fuzz_dep_.CoverTab[117483]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:116
		// _ = "end of CoverTab[117483]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:116
	// _ = "end of CoverTab[117474]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:116
	_go_fuzz_dep_.CoverTab[117475]++

											if ok2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:118
		_go_fuzz_dep_.CoverTab[117484]++
												fi, b, err := lbase.LstatIfPossible(name)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:120
			_go_fuzz_dep_.CoverTab[117486]++
													return fi, b, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:121
			// _ = "end of CoverTab[117486]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:122
			_go_fuzz_dep_.CoverTab[117487]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:122
			// _ = "end of CoverTab[117487]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:122
		// _ = "end of CoverTab[117484]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:122
		_go_fuzz_dep_.CoverTab[117485]++
												if !u.isNotExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:123
			_go_fuzz_dep_.CoverTab[117488]++
													return nil, b, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:124
			// _ = "end of CoverTab[117488]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:125
			_go_fuzz_dep_.CoverTab[117489]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:125
			// _ = "end of CoverTab[117489]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:125
		// _ = "end of CoverTab[117485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:126
		_go_fuzz_dep_.CoverTab[117490]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:126
		// _ = "end of CoverTab[117490]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:126
	// _ = "end of CoverTab[117475]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:126
	_go_fuzz_dep_.CoverTab[117476]++

											fi, err := u.Stat(name)

											return fi, false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:130
	// _ = "end of CoverTab[117476]"
}

func (u *CopyOnWriteFs) SymlinkIfPossible(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:133
	_go_fuzz_dep_.CoverTab[117491]++
											if slayer, ok := u.layer.(Linker); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:134
		_go_fuzz_dep_.CoverTab[117493]++
												return slayer.SymlinkIfPossible(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:135
		// _ = "end of CoverTab[117493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:136
		_go_fuzz_dep_.CoverTab[117494]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:136
		// _ = "end of CoverTab[117494]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:136
	// _ = "end of CoverTab[117491]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:136
	_go_fuzz_dep_.CoverTab[117492]++

											return &os.LinkError{Op: "symlink", Old: oldname, New: newname, Err: ErrNoSymlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:138
	// _ = "end of CoverTab[117492]"
}

func (u *CopyOnWriteFs) ReadlinkIfPossible(name string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:141
	_go_fuzz_dep_.CoverTab[117495]++
											if rlayer, ok := u.layer.(LinkReader); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:142
		_go_fuzz_dep_.CoverTab[117498]++
												return rlayer.ReadlinkIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:143
		// _ = "end of CoverTab[117498]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:144
		_go_fuzz_dep_.CoverTab[117499]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:144
		// _ = "end of CoverTab[117499]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:144
	// _ = "end of CoverTab[117495]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:144
	_go_fuzz_dep_.CoverTab[117496]++

											if rbase, ok := u.base.(LinkReader); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:146
		_go_fuzz_dep_.CoverTab[117500]++
												return rbase.ReadlinkIfPossible(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:147
		// _ = "end of CoverTab[117500]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:148
		_go_fuzz_dep_.CoverTab[117501]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:148
		// _ = "end of CoverTab[117501]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:148
	// _ = "end of CoverTab[117496]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:148
	_go_fuzz_dep_.CoverTab[117497]++

											return "", &os.PathError{Op: "readlink", Path: name, Err: ErrNoReadlink}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:150
	// _ = "end of CoverTab[117497]"
}

func (u *CopyOnWriteFs) isNotExist(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:153
	_go_fuzz_dep_.CoverTab[117502]++
											if e, ok := err.(*os.PathError); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:154
		_go_fuzz_dep_.CoverTab[117505]++
												err = e.Err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:155
		// _ = "end of CoverTab[117505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:156
		_go_fuzz_dep_.CoverTab[117506]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:156
		// _ = "end of CoverTab[117506]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:156
	// _ = "end of CoverTab[117502]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:156
	_go_fuzz_dep_.CoverTab[117503]++
											if err == os.ErrNotExist || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		_go_fuzz_dep_.CoverTab[117507]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		return err == syscall.ENOENT
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		// _ = "end of CoverTab[117507]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		_go_fuzz_dep_.CoverTab[117508]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		return err == syscall.ENOTDIR
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		// _ = "end of CoverTab[117508]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:157
		_go_fuzz_dep_.CoverTab[117509]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:158
		// _ = "end of CoverTab[117509]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:159
		_go_fuzz_dep_.CoverTab[117510]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:159
		// _ = "end of CoverTab[117510]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:159
	// _ = "end of CoverTab[117503]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:159
	_go_fuzz_dep_.CoverTab[117504]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:160
	// _ = "end of CoverTab[117504]"
}

// Renaming files present only in the base layer is not permitted
func (u *CopyOnWriteFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:164
	_go_fuzz_dep_.CoverTab[117511]++
											b, err := u.isBaseFile(oldname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:166
		_go_fuzz_dep_.CoverTab[117514]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:167
		// _ = "end of CoverTab[117514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:168
		_go_fuzz_dep_.CoverTab[117515]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:168
		// _ = "end of CoverTab[117515]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:168
	// _ = "end of CoverTab[117511]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:168
	_go_fuzz_dep_.CoverTab[117512]++
											if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:169
		_go_fuzz_dep_.CoverTab[117516]++
												return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:170
		// _ = "end of CoverTab[117516]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:171
		_go_fuzz_dep_.CoverTab[117517]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:171
		// _ = "end of CoverTab[117517]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:171
	// _ = "end of CoverTab[117512]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:171
	_go_fuzz_dep_.CoverTab[117513]++
											return u.layer.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:172
	// _ = "end of CoverTab[117513]"
}

// Removing files present only in the base layer is not permitted. If
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:175
// a file is present in the base layer and the overlay, only the overlay
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:175
// will be removed.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:178
func (u *CopyOnWriteFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:178
	_go_fuzz_dep_.CoverTab[117518]++
											err := u.layer.Remove(name)
											switch err {
	case syscall.ENOENT:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:181
		_go_fuzz_dep_.CoverTab[117519]++
												_, err = u.base.Stat(name)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:183
			_go_fuzz_dep_.CoverTab[117522]++
													return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:184
			// _ = "end of CoverTab[117522]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:185
			_go_fuzz_dep_.CoverTab[117523]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:185
			// _ = "end of CoverTab[117523]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:185
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:185
		// _ = "end of CoverTab[117519]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:185
		_go_fuzz_dep_.CoverTab[117520]++
												return syscall.ENOENT
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:186
		// _ = "end of CoverTab[117520]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:187
		_go_fuzz_dep_.CoverTab[117521]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:188
		// _ = "end of CoverTab[117521]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:189
	// _ = "end of CoverTab[117518]"
}

func (u *CopyOnWriteFs) RemoveAll(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:192
	_go_fuzz_dep_.CoverTab[117524]++
											err := u.layer.RemoveAll(name)
											switch err {
	case syscall.ENOENT:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:195
		_go_fuzz_dep_.CoverTab[117525]++
												_, err = u.base.Stat(name)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:197
			_go_fuzz_dep_.CoverTab[117528]++
													return syscall.EPERM
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:198
			// _ = "end of CoverTab[117528]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:199
			_go_fuzz_dep_.CoverTab[117529]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:199
			// _ = "end of CoverTab[117529]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:199
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:199
		// _ = "end of CoverTab[117525]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:199
		_go_fuzz_dep_.CoverTab[117526]++
												return syscall.ENOENT
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:200
		// _ = "end of CoverTab[117526]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:201
		_go_fuzz_dep_.CoverTab[117527]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:202
		// _ = "end of CoverTab[117527]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:203
	// _ = "end of CoverTab[117524]"
}

func (u *CopyOnWriteFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:206
	_go_fuzz_dep_.CoverTab[117530]++
											b, err := u.isBaseFile(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:208
		_go_fuzz_dep_.CoverTab[117534]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:209
		// _ = "end of CoverTab[117534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:210
		_go_fuzz_dep_.CoverTab[117535]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:210
		// _ = "end of CoverTab[117535]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:210
	// _ = "end of CoverTab[117530]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:210
	_go_fuzz_dep_.CoverTab[117531]++

											if flag&(os.O_WRONLY|os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:212
		_go_fuzz_dep_.CoverTab[117536]++
												if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:213
			_go_fuzz_dep_.CoverTab[117542]++
													if err = u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:214
				_go_fuzz_dep_.CoverTab[117544]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:215
				// _ = "end of CoverTab[117544]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:216
				_go_fuzz_dep_.CoverTab[117545]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:216
				// _ = "end of CoverTab[117545]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:216
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:216
			// _ = "end of CoverTab[117542]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:216
			_go_fuzz_dep_.CoverTab[117543]++
													return u.layer.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:217
			// _ = "end of CoverTab[117543]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:218
			_go_fuzz_dep_.CoverTab[117546]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:218
			// _ = "end of CoverTab[117546]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:218
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:218
		// _ = "end of CoverTab[117536]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:218
		_go_fuzz_dep_.CoverTab[117537]++

												dir := filepath.Dir(name)
												isaDir, err := IsDir(u.base, dir)
												if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:222
			_go_fuzz_dep_.CoverTab[117547]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:222
			return !os.IsNotExist(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:222
			// _ = "end of CoverTab[117547]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:222
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:222
			_go_fuzz_dep_.CoverTab[117548]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:223
			// _ = "end of CoverTab[117548]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:224
			_go_fuzz_dep_.CoverTab[117549]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:224
			// _ = "end of CoverTab[117549]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:224
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:224
		// _ = "end of CoverTab[117537]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:224
		_go_fuzz_dep_.CoverTab[117538]++
												if isaDir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:225
			_go_fuzz_dep_.CoverTab[117550]++
													if err = u.layer.MkdirAll(dir, 0777); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:226
				_go_fuzz_dep_.CoverTab[117552]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:227
				// _ = "end of CoverTab[117552]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:228
				_go_fuzz_dep_.CoverTab[117553]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:228
				// _ = "end of CoverTab[117553]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:228
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:228
			// _ = "end of CoverTab[117550]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:228
			_go_fuzz_dep_.CoverTab[117551]++
													return u.layer.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:229
			// _ = "end of CoverTab[117551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:230
			_go_fuzz_dep_.CoverTab[117554]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:230
			// _ = "end of CoverTab[117554]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:230
		// _ = "end of CoverTab[117538]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:230
		_go_fuzz_dep_.CoverTab[117539]++

												isaDir, err = IsDir(u.layer, dir)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:233
			_go_fuzz_dep_.CoverTab[117555]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:234
			// _ = "end of CoverTab[117555]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:235
			_go_fuzz_dep_.CoverTab[117556]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:235
			// _ = "end of CoverTab[117556]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:235
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:235
		// _ = "end of CoverTab[117539]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:235
		_go_fuzz_dep_.CoverTab[117540]++
												if isaDir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:236
			_go_fuzz_dep_.CoverTab[117557]++
													return u.layer.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:237
			// _ = "end of CoverTab[117557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:238
			_go_fuzz_dep_.CoverTab[117558]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:238
			// _ = "end of CoverTab[117558]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:238
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:238
		// _ = "end of CoverTab[117540]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:238
		_go_fuzz_dep_.CoverTab[117541]++

												return nil, &os.PathError{Op: "open", Path: name, Err: syscall.ENOTDIR}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:240
		// _ = "end of CoverTab[117541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:241
		_go_fuzz_dep_.CoverTab[117559]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:241
		// _ = "end of CoverTab[117559]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:241
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:241
	// _ = "end of CoverTab[117531]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:241
	_go_fuzz_dep_.CoverTab[117532]++
											if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:242
		_go_fuzz_dep_.CoverTab[117560]++
												return u.base.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:243
		// _ = "end of CoverTab[117560]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:244
		_go_fuzz_dep_.CoverTab[117561]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:244
		// _ = "end of CoverTab[117561]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:244
	// _ = "end of CoverTab[117532]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:244
	_go_fuzz_dep_.CoverTab[117533]++
											return u.layer.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:245
	// _ = "end of CoverTab[117533]"
}

// This function handles the 9 different possibilities caused
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:248
// by the union which are the intersection of the following...
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:248
//	layer: doesn't exist, exists as a file, and exists as a directory
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:248
//	base:  doesn't exist, exists as a file, and exists as a directory
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:252
func (u *CopyOnWriteFs) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:252
	_go_fuzz_dep_.CoverTab[117562]++

											b, err := u.isBaseFile(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:255
		_go_fuzz_dep_.CoverTab[117569]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:256
		// _ = "end of CoverTab[117569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:257
		_go_fuzz_dep_.CoverTab[117570]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:257
		// _ = "end of CoverTab[117570]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:257
	// _ = "end of CoverTab[117562]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:257
	_go_fuzz_dep_.CoverTab[117563]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:260
	if b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:260
		_go_fuzz_dep_.CoverTab[117571]++
												return u.base.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:261
		// _ = "end of CoverTab[117571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:262
		_go_fuzz_dep_.CoverTab[117572]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:262
		// _ = "end of CoverTab[117572]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:262
	// _ = "end of CoverTab[117563]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:262
	_go_fuzz_dep_.CoverTab[117564]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:265
	dir, err := IsDir(u.layer, name)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:266
		_go_fuzz_dep_.CoverTab[117573]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:267
		// _ = "end of CoverTab[117573]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:268
		_go_fuzz_dep_.CoverTab[117574]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:268
		// _ = "end of CoverTab[117574]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:268
	// _ = "end of CoverTab[117564]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:268
	_go_fuzz_dep_.CoverTab[117565]++
											if !dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:269
		_go_fuzz_dep_.CoverTab[117575]++
												return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:270
		// _ = "end of CoverTab[117575]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:271
		_go_fuzz_dep_.CoverTab[117576]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:271
		// _ = "end of CoverTab[117576]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:271
	// _ = "end of CoverTab[117565]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:271
	_go_fuzz_dep_.CoverTab[117566]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:279
	dir, err = IsDir(u.base, name)
	if !dir || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:280
		_go_fuzz_dep_.CoverTab[117577]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:280
		return err != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:280
		// _ = "end of CoverTab[117577]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:280
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:280
		_go_fuzz_dep_.CoverTab[117578]++
												return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:281
		// _ = "end of CoverTab[117578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:282
		_go_fuzz_dep_.CoverTab[117579]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:282
		// _ = "end of CoverTab[117579]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:282
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:282
	// _ = "end of CoverTab[117566]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:282
	_go_fuzz_dep_.CoverTab[117567]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:286
	bfile, bErr := u.base.Open(name)
											lfile, lErr := u.layer.Open(name)

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
	if bErr != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
		_go_fuzz_dep_.CoverTab[117580]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
		return lErr != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
		// _ = "end of CoverTab[117580]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:290
		_go_fuzz_dep_.CoverTab[117581]++
												return nil, fmt.Errorf("BaseErr: %v\nOverlayErr: %v", bErr, lErr)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:291
		// _ = "end of CoverTab[117581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:292
		_go_fuzz_dep_.CoverTab[117582]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:292
		// _ = "end of CoverTab[117582]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:292
	// _ = "end of CoverTab[117567]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:292
	_go_fuzz_dep_.CoverTab[117568]++

											return &UnionFile{Base: bfile, Layer: lfile}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:294
	// _ = "end of CoverTab[117568]"
}

func (u *CopyOnWriteFs) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:297
	_go_fuzz_dep_.CoverTab[117583]++
											dir, err := IsDir(u.base, name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:299
		_go_fuzz_dep_.CoverTab[117586]++
												return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:300
		// _ = "end of CoverTab[117586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:301
		_go_fuzz_dep_.CoverTab[117587]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:301
		// _ = "end of CoverTab[117587]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:301
	// _ = "end of CoverTab[117583]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:301
	_go_fuzz_dep_.CoverTab[117584]++
											if dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:302
		_go_fuzz_dep_.CoverTab[117588]++
												return ErrFileExists
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:303
		// _ = "end of CoverTab[117588]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:304
		_go_fuzz_dep_.CoverTab[117589]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:304
		// _ = "end of CoverTab[117589]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:304
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:304
	// _ = "end of CoverTab[117584]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:304
	_go_fuzz_dep_.CoverTab[117585]++
											return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:305
	// _ = "end of CoverTab[117585]"
}

func (u *CopyOnWriteFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:308
	_go_fuzz_dep_.CoverTab[117590]++
											return "CopyOnWriteFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:309
	// _ = "end of CoverTab[117590]"
}

func (u *CopyOnWriteFs) MkdirAll(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:312
	_go_fuzz_dep_.CoverTab[117591]++
											dir, err := IsDir(u.base, name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:314
		_go_fuzz_dep_.CoverTab[117594]++
												return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:315
		// _ = "end of CoverTab[117594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:316
		_go_fuzz_dep_.CoverTab[117595]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:316
		// _ = "end of CoverTab[117595]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:316
	// _ = "end of CoverTab[117591]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:316
	_go_fuzz_dep_.CoverTab[117592]++
											if dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:317
		_go_fuzz_dep_.CoverTab[117596]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:319
		// _ = "end of CoverTab[117596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:320
		_go_fuzz_dep_.CoverTab[117597]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:320
		// _ = "end of CoverTab[117597]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:320
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:320
	// _ = "end of CoverTab[117592]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:320
	_go_fuzz_dep_.CoverTab[117593]++
											return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:321
	// _ = "end of CoverTab[117593]"
}

func (u *CopyOnWriteFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:324
	_go_fuzz_dep_.CoverTab[117598]++
											return u.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:325
	// _ = "end of CoverTab[117598]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:326
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/copyOnWriteFs.go:326
var _ = _go_fuzz_dep_.CoverTab
