//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:1
)

import (
	"os"
	"syscall"
	"time"
)

// If the cache duration is 0, cache time will be unlimited, i.e. once
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// a file is in the layer, the base will never be read again for this file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// For cache times greater than 0, the modification time of a file is
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// checked. Note that a lot of file system implementations only allow a
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// resolution of a second for timestamps... or as the godoc for os.Chtimes()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// states: "The underlying filesystem may truncate or round the values to a
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// less precise time unit."
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// This caching union will forward all write calls also to the base file
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// system first. To prevent writing to the base Fs, wrap it in a read-only
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// filter - Note: this will also make the overlay read-only, for writing files
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:9
// in the overlay, use the overlay Fs directly, not via the union Fs.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:22
type CacheOnReadFs struct {
	base		Fs
	layer		Fs
	cacheTime	time.Duration
}

func NewCacheOnReadFs(base Fs, layer Fs, cacheTime time.Duration) Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:28
	_go_fuzz_dep_.CoverTab[117248]++
											return &CacheOnReadFs{base: base, layer: layer, cacheTime: cacheTime}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:29
	// _ = "end of CoverTab[117248]"
}

type cacheState int

const (
	// not present in the overlay, unknown if it exists in the base:
	cacheMiss	cacheState	= iota
	// present in the overlay and in base, base file is newer:
	cacheStale
	// present in the overlay - with cache time == 0 it may exist in the base,
	// with cacheTime > 0 it exists in the base and is same age or newer in the
	// overlay
	cacheHit
	// happens if someone writes directly to the overlay without
	// going through this union
	cacheLocal
)

func (u *CacheOnReadFs) cacheStatus(name string) (state cacheState, fi os.FileInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:48
	_go_fuzz_dep_.CoverTab[117249]++
											var lfi, bfi os.FileInfo
											lfi, err = u.layer.Stat(name)
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:51
		_go_fuzz_dep_.CoverTab[117252]++
												if u.cacheTime == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:52
			_go_fuzz_dep_.CoverTab[117255]++
													return cacheHit, lfi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:53
			// _ = "end of CoverTab[117255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:54
			_go_fuzz_dep_.CoverTab[117256]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:54
			// _ = "end of CoverTab[117256]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:54
		// _ = "end of CoverTab[117252]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:54
		_go_fuzz_dep_.CoverTab[117253]++
												if lfi.ModTime().Add(u.cacheTime).Before(time.Now()) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:55
			_go_fuzz_dep_.CoverTab[117257]++
													bfi, err = u.base.Stat(name)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:57
				_go_fuzz_dep_.CoverTab[117259]++
														return cacheLocal, lfi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:58
				// _ = "end of CoverTab[117259]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:59
				_go_fuzz_dep_.CoverTab[117260]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:59
				// _ = "end of CoverTab[117260]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:59
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:59
			// _ = "end of CoverTab[117257]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:59
			_go_fuzz_dep_.CoverTab[117258]++
													if bfi.ModTime().After(lfi.ModTime()) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:60
				_go_fuzz_dep_.CoverTab[117261]++
														return cacheStale, bfi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:61
				// _ = "end of CoverTab[117261]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:62
				_go_fuzz_dep_.CoverTab[117262]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:62
				// _ = "end of CoverTab[117262]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:62
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:62
			// _ = "end of CoverTab[117258]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:63
			_go_fuzz_dep_.CoverTab[117263]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:63
			// _ = "end of CoverTab[117263]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:63
		// _ = "end of CoverTab[117253]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:63
		_go_fuzz_dep_.CoverTab[117254]++
												return cacheHit, lfi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:64
		// _ = "end of CoverTab[117254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:65
		_go_fuzz_dep_.CoverTab[117264]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:65
		// _ = "end of CoverTab[117264]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:65
	// _ = "end of CoverTab[117249]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:65
	_go_fuzz_dep_.CoverTab[117250]++

											if err == syscall.ENOENT || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:67
		_go_fuzz_dep_.CoverTab[117265]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:67
		return os.IsNotExist(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:67
		// _ = "end of CoverTab[117265]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:67
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:67
		_go_fuzz_dep_.CoverTab[117266]++
												return cacheMiss, nil, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:68
		// _ = "end of CoverTab[117266]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:69
		_go_fuzz_dep_.CoverTab[117267]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:69
		// _ = "end of CoverTab[117267]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:69
	// _ = "end of CoverTab[117250]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:69
	_go_fuzz_dep_.CoverTab[117251]++

											return cacheMiss, nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:71
	// _ = "end of CoverTab[117251]"
}

func (u *CacheOnReadFs) copyToLayer(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:74
	_go_fuzz_dep_.CoverTab[117268]++
											return copyToLayer(u.base, u.layer, name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:75
	// _ = "end of CoverTab[117268]"
}

func (u *CacheOnReadFs) copyFileToLayer(name string, flag int, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:78
	_go_fuzz_dep_.CoverTab[117269]++
											return copyFileToLayer(u.base, u.layer, name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:79
	// _ = "end of CoverTab[117269]"
}

func (u *CacheOnReadFs) Chtimes(name string, atime, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:82
	_go_fuzz_dep_.CoverTab[117270]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:84
		_go_fuzz_dep_.CoverTab[117274]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:85
		// _ = "end of CoverTab[117274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:86
		_go_fuzz_dep_.CoverTab[117275]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:86
		// _ = "end of CoverTab[117275]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:86
	// _ = "end of CoverTab[117270]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:86
	_go_fuzz_dep_.CoverTab[117271]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:88
		_go_fuzz_dep_.CoverTab[117276]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:88
		// _ = "end of CoverTab[117276]"
	case cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:89
		_go_fuzz_dep_.CoverTab[117277]++
												err = u.base.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:90
		// _ = "end of CoverTab[117277]"
	case cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:91
		_go_fuzz_dep_.CoverTab[117278]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:92
			_go_fuzz_dep_.CoverTab[117281]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:93
			// _ = "end of CoverTab[117281]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:94
			_go_fuzz_dep_.CoverTab[117282]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:94
			// _ = "end of CoverTab[117282]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:94
		// _ = "end of CoverTab[117278]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:94
		_go_fuzz_dep_.CoverTab[117279]++
												err = u.base.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:95
		// _ = "end of CoverTab[117279]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:95
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:95
		_go_fuzz_dep_.CoverTab[117280]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:95
		// _ = "end of CoverTab[117280]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:96
	// _ = "end of CoverTab[117271]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:96
	_go_fuzz_dep_.CoverTab[117272]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:97
		_go_fuzz_dep_.CoverTab[117283]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:98
		// _ = "end of CoverTab[117283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:99
		_go_fuzz_dep_.CoverTab[117284]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:99
		// _ = "end of CoverTab[117284]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:99
	// _ = "end of CoverTab[117272]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:99
	_go_fuzz_dep_.CoverTab[117273]++
											return u.layer.Chtimes(name, atime, mtime)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:100
	// _ = "end of CoverTab[117273]"
}

func (u *CacheOnReadFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:103
	_go_fuzz_dep_.CoverTab[117285]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:105
		_go_fuzz_dep_.CoverTab[117289]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:106
		// _ = "end of CoverTab[117289]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:107
		_go_fuzz_dep_.CoverTab[117290]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:107
		// _ = "end of CoverTab[117290]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:107
	// _ = "end of CoverTab[117285]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:107
	_go_fuzz_dep_.CoverTab[117286]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:109
		_go_fuzz_dep_.CoverTab[117291]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:109
		// _ = "end of CoverTab[117291]"
	case cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:110
		_go_fuzz_dep_.CoverTab[117292]++
												err = u.base.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:111
		// _ = "end of CoverTab[117292]"
	case cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:112
		_go_fuzz_dep_.CoverTab[117293]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:113
			_go_fuzz_dep_.CoverTab[117296]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:114
			// _ = "end of CoverTab[117296]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:115
			_go_fuzz_dep_.CoverTab[117297]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:115
			// _ = "end of CoverTab[117297]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:115
		// _ = "end of CoverTab[117293]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:115
		_go_fuzz_dep_.CoverTab[117294]++
												err = u.base.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:116
		// _ = "end of CoverTab[117294]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:116
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:116
		_go_fuzz_dep_.CoverTab[117295]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:116
		// _ = "end of CoverTab[117295]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:117
	// _ = "end of CoverTab[117286]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:117
	_go_fuzz_dep_.CoverTab[117287]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:118
		_go_fuzz_dep_.CoverTab[117298]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:119
		// _ = "end of CoverTab[117298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:120
		_go_fuzz_dep_.CoverTab[117299]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:120
		// _ = "end of CoverTab[117299]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:120
	// _ = "end of CoverTab[117287]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:120
	_go_fuzz_dep_.CoverTab[117288]++
											return u.layer.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:121
	// _ = "end of CoverTab[117288]"
}

func (u *CacheOnReadFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:124
	_go_fuzz_dep_.CoverTab[117300]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:126
		_go_fuzz_dep_.CoverTab[117304]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:127
		// _ = "end of CoverTab[117304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:128
		_go_fuzz_dep_.CoverTab[117305]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:128
		// _ = "end of CoverTab[117305]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:128
	// _ = "end of CoverTab[117300]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:128
	_go_fuzz_dep_.CoverTab[117301]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:130
		_go_fuzz_dep_.CoverTab[117306]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:130
		// _ = "end of CoverTab[117306]"
	case cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:131
		_go_fuzz_dep_.CoverTab[117307]++
												err = u.base.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:132
		// _ = "end of CoverTab[117307]"
	case cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:133
		_go_fuzz_dep_.CoverTab[117308]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:134
			_go_fuzz_dep_.CoverTab[117311]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:135
			// _ = "end of CoverTab[117311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:136
			_go_fuzz_dep_.CoverTab[117312]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:136
			// _ = "end of CoverTab[117312]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:136
		// _ = "end of CoverTab[117308]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:136
		_go_fuzz_dep_.CoverTab[117309]++
												err = u.base.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:137
		// _ = "end of CoverTab[117309]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:137
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:137
		_go_fuzz_dep_.CoverTab[117310]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:137
		// _ = "end of CoverTab[117310]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:138
	// _ = "end of CoverTab[117301]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:138
	_go_fuzz_dep_.CoverTab[117302]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:139
		_go_fuzz_dep_.CoverTab[117313]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:140
		// _ = "end of CoverTab[117313]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:141
		_go_fuzz_dep_.CoverTab[117314]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:141
		// _ = "end of CoverTab[117314]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:141
	// _ = "end of CoverTab[117302]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:141
	_go_fuzz_dep_.CoverTab[117303]++
											return u.layer.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:142
	// _ = "end of CoverTab[117303]"
}

func (u *CacheOnReadFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:145
	_go_fuzz_dep_.CoverTab[117315]++
											st, fi, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:147
		_go_fuzz_dep_.CoverTab[117317]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:148
		// _ = "end of CoverTab[117317]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:149
		_go_fuzz_dep_.CoverTab[117318]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:149
		// _ = "end of CoverTab[117318]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:149
	// _ = "end of CoverTab[117315]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:149
	_go_fuzz_dep_.CoverTab[117316]++
											switch st {
	case cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:151
		_go_fuzz_dep_.CoverTab[117319]++
												return u.base.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:152
		// _ = "end of CoverTab[117319]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:153
		_go_fuzz_dep_.CoverTab[117320]++
												return fi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:154
		// _ = "end of CoverTab[117320]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:155
	// _ = "end of CoverTab[117316]"
}

func (u *CacheOnReadFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:158
	_go_fuzz_dep_.CoverTab[117321]++
											st, _, err := u.cacheStatus(oldname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:160
		_go_fuzz_dep_.CoverTab[117325]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:161
		// _ = "end of CoverTab[117325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:162
		_go_fuzz_dep_.CoverTab[117326]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:162
		// _ = "end of CoverTab[117326]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:162
	// _ = "end of CoverTab[117321]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:162
	_go_fuzz_dep_.CoverTab[117322]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:164
		_go_fuzz_dep_.CoverTab[117327]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:164
		// _ = "end of CoverTab[117327]"
	case cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:165
		_go_fuzz_dep_.CoverTab[117328]++
												err = u.base.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:166
		// _ = "end of CoverTab[117328]"
	case cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:167
		_go_fuzz_dep_.CoverTab[117329]++
												if err := u.copyToLayer(oldname); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:168
			_go_fuzz_dep_.CoverTab[117332]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:169
			// _ = "end of CoverTab[117332]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:170
			_go_fuzz_dep_.CoverTab[117333]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:170
			// _ = "end of CoverTab[117333]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:170
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:170
		// _ = "end of CoverTab[117329]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:170
		_go_fuzz_dep_.CoverTab[117330]++
												err = u.base.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:171
		// _ = "end of CoverTab[117330]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:171
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:171
		_go_fuzz_dep_.CoverTab[117331]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:171
		// _ = "end of CoverTab[117331]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:172
	// _ = "end of CoverTab[117322]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:172
	_go_fuzz_dep_.CoverTab[117323]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:173
		_go_fuzz_dep_.CoverTab[117334]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:174
		// _ = "end of CoverTab[117334]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:175
		_go_fuzz_dep_.CoverTab[117335]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:175
		// _ = "end of CoverTab[117335]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:175
	// _ = "end of CoverTab[117323]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:175
	_go_fuzz_dep_.CoverTab[117324]++
											return u.layer.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:176
	// _ = "end of CoverTab[117324]"
}

func (u *CacheOnReadFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:179
	_go_fuzz_dep_.CoverTab[117336]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:181
		_go_fuzz_dep_.CoverTab[117340]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:182
		// _ = "end of CoverTab[117340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:183
		_go_fuzz_dep_.CoverTab[117341]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:183
		// _ = "end of CoverTab[117341]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:183
	// _ = "end of CoverTab[117336]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:183
	_go_fuzz_dep_.CoverTab[117337]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:185
		_go_fuzz_dep_.CoverTab[117342]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:185
		// _ = "end of CoverTab[117342]"
	case cacheHit, cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:186
		_go_fuzz_dep_.CoverTab[117343]++
												err = u.base.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:187
		// _ = "end of CoverTab[117343]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:187
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:187
		_go_fuzz_dep_.CoverTab[117344]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:187
		// _ = "end of CoverTab[117344]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:188
	// _ = "end of CoverTab[117337]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:188
	_go_fuzz_dep_.CoverTab[117338]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:189
		_go_fuzz_dep_.CoverTab[117345]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:190
		// _ = "end of CoverTab[117345]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:191
		_go_fuzz_dep_.CoverTab[117346]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:191
		// _ = "end of CoverTab[117346]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:191
	// _ = "end of CoverTab[117338]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:191
	_go_fuzz_dep_.CoverTab[117339]++
											return u.layer.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:192
	// _ = "end of CoverTab[117339]"
}

func (u *CacheOnReadFs) RemoveAll(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:195
	_go_fuzz_dep_.CoverTab[117347]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:197
		_go_fuzz_dep_.CoverTab[117351]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:198
		// _ = "end of CoverTab[117351]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:199
		_go_fuzz_dep_.CoverTab[117352]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:199
		// _ = "end of CoverTab[117352]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:199
	// _ = "end of CoverTab[117347]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:199
	_go_fuzz_dep_.CoverTab[117348]++
											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:201
		_go_fuzz_dep_.CoverTab[117353]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:201
		// _ = "end of CoverTab[117353]"
	case cacheHit, cacheStale, cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:202
		_go_fuzz_dep_.CoverTab[117354]++
												err = u.base.RemoveAll(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:203
		// _ = "end of CoverTab[117354]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:203
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:203
		_go_fuzz_dep_.CoverTab[117355]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:203
		// _ = "end of CoverTab[117355]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:204
	// _ = "end of CoverTab[117348]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:204
	_go_fuzz_dep_.CoverTab[117349]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:205
		_go_fuzz_dep_.CoverTab[117356]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:206
		// _ = "end of CoverTab[117356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:207
		_go_fuzz_dep_.CoverTab[117357]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:207
		// _ = "end of CoverTab[117357]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:207
	// _ = "end of CoverTab[117349]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:207
	_go_fuzz_dep_.CoverTab[117350]++
											return u.layer.RemoveAll(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:208
	// _ = "end of CoverTab[117350]"
}

func (u *CacheOnReadFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:211
	_go_fuzz_dep_.CoverTab[117358]++
											st, _, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:213
		_go_fuzz_dep_.CoverTab[117362]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:214
		// _ = "end of CoverTab[117362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:215
		_go_fuzz_dep_.CoverTab[117363]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:215
		// _ = "end of CoverTab[117363]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:215
	// _ = "end of CoverTab[117358]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:215
	_go_fuzz_dep_.CoverTab[117359]++
											switch st {
	case cacheLocal, cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:217
		_go_fuzz_dep_.CoverTab[117364]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:217
		// _ = "end of CoverTab[117364]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:218
		_go_fuzz_dep_.CoverTab[117365]++
												if err := u.copyFileToLayer(name, flag, perm); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:219
			_go_fuzz_dep_.CoverTab[117366]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:220
			// _ = "end of CoverTab[117366]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:221
			_go_fuzz_dep_.CoverTab[117367]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:221
			// _ = "end of CoverTab[117367]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:221
		// _ = "end of CoverTab[117365]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:222
	// _ = "end of CoverTab[117359]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:222
	_go_fuzz_dep_.CoverTab[117360]++
											if flag&(os.O_WRONLY|syscall.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:223
		_go_fuzz_dep_.CoverTab[117368]++
												bfi, err := u.base.OpenFile(name, flag, perm)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:225
			_go_fuzz_dep_.CoverTab[117371]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:226
			// _ = "end of CoverTab[117371]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:227
			_go_fuzz_dep_.CoverTab[117372]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:227
			// _ = "end of CoverTab[117372]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:227
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:227
		// _ = "end of CoverTab[117368]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:227
		_go_fuzz_dep_.CoverTab[117369]++
												lfi, err := u.layer.OpenFile(name, flag, perm)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:229
			_go_fuzz_dep_.CoverTab[117373]++
													bfi.Close()
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:231
			// _ = "end of CoverTab[117373]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:232
			_go_fuzz_dep_.CoverTab[117374]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:232
			// _ = "end of CoverTab[117374]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:232
		// _ = "end of CoverTab[117369]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:232
		_go_fuzz_dep_.CoverTab[117370]++
												return &UnionFile{Base: bfi, Layer: lfi}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:233
		// _ = "end of CoverTab[117370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:234
		_go_fuzz_dep_.CoverTab[117375]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:234
		// _ = "end of CoverTab[117375]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:234
	// _ = "end of CoverTab[117360]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:234
	_go_fuzz_dep_.CoverTab[117361]++
											return u.layer.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:235
	// _ = "end of CoverTab[117361]"
}

func (u *CacheOnReadFs) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:238
	_go_fuzz_dep_.CoverTab[117376]++
											st, fi, err := u.cacheStatus(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:240
		_go_fuzz_dep_.CoverTab[117380]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:241
		// _ = "end of CoverTab[117380]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:242
		_go_fuzz_dep_.CoverTab[117381]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:242
		// _ = "end of CoverTab[117381]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:242
	// _ = "end of CoverTab[117376]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:242
	_go_fuzz_dep_.CoverTab[117377]++

											switch st {
	case cacheLocal:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:245
		_go_fuzz_dep_.CoverTab[117382]++
												return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:246
		// _ = "end of CoverTab[117382]"

	case cacheMiss:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:248
		_go_fuzz_dep_.CoverTab[117383]++
												bfi, err := u.base.Stat(name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:250
			_go_fuzz_dep_.CoverTab[117390]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:251
			// _ = "end of CoverTab[117390]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:252
			_go_fuzz_dep_.CoverTab[117391]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:252
			// _ = "end of CoverTab[117391]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:252
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:252
		// _ = "end of CoverTab[117383]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:252
		_go_fuzz_dep_.CoverTab[117384]++
												if bfi.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:253
			_go_fuzz_dep_.CoverTab[117392]++
													return u.base.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:254
			// _ = "end of CoverTab[117392]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:255
			_go_fuzz_dep_.CoverTab[117393]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:255
			// _ = "end of CoverTab[117393]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:255
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:255
		// _ = "end of CoverTab[117384]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:255
		_go_fuzz_dep_.CoverTab[117385]++
												if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:256
			_go_fuzz_dep_.CoverTab[117394]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:257
			// _ = "end of CoverTab[117394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:258
			_go_fuzz_dep_.CoverTab[117395]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:258
			// _ = "end of CoverTab[117395]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:258
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:258
		// _ = "end of CoverTab[117385]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:258
		_go_fuzz_dep_.CoverTab[117386]++
												return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:259
		// _ = "end of CoverTab[117386]"

	case cacheStale:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:261
		_go_fuzz_dep_.CoverTab[117387]++
												if !fi.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:262
			_go_fuzz_dep_.CoverTab[117396]++
													if err := u.copyToLayer(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:263
				_go_fuzz_dep_.CoverTab[117398]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:264
				// _ = "end of CoverTab[117398]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:265
				_go_fuzz_dep_.CoverTab[117399]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:265
				// _ = "end of CoverTab[117399]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:265
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:265
			// _ = "end of CoverTab[117396]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:265
			_go_fuzz_dep_.CoverTab[117397]++
													return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:266
			// _ = "end of CoverTab[117397]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:267
			_go_fuzz_dep_.CoverTab[117400]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:267
			// _ = "end of CoverTab[117400]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:267
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:267
		// _ = "end of CoverTab[117387]"
	case cacheHit:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:268
		_go_fuzz_dep_.CoverTab[117388]++
												if !fi.IsDir() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:269
			_go_fuzz_dep_.CoverTab[117401]++
													return u.layer.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:270
			// _ = "end of CoverTab[117401]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
			_go_fuzz_dep_.CoverTab[117402]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
			// _ = "end of CoverTab[117402]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
		// _ = "end of CoverTab[117388]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
		_go_fuzz_dep_.CoverTab[117389]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:271
		// _ = "end of CoverTab[117389]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:272
	// _ = "end of CoverTab[117377]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:272
	_go_fuzz_dep_.CoverTab[117378]++

											bfile, _ := u.base.Open(name)
											lfile, err := u.layer.Open(name)
											if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:276
		_go_fuzz_dep_.CoverTab[117403]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:276
		return bfile == nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:276
		// _ = "end of CoverTab[117403]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:276
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:276
		_go_fuzz_dep_.CoverTab[117404]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:277
		// _ = "end of CoverTab[117404]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:278
		_go_fuzz_dep_.CoverTab[117405]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:278
		// _ = "end of CoverTab[117405]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:278
	// _ = "end of CoverTab[117378]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:278
	_go_fuzz_dep_.CoverTab[117379]++
											return &UnionFile{Base: bfile, Layer: lfile}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:279
	// _ = "end of CoverTab[117379]"
}

func (u *CacheOnReadFs) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:282
	_go_fuzz_dep_.CoverTab[117406]++
											err := u.base.Mkdir(name, perm)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:284
		_go_fuzz_dep_.CoverTab[117408]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:285
		// _ = "end of CoverTab[117408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:286
		_go_fuzz_dep_.CoverTab[117409]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:286
		// _ = "end of CoverTab[117409]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:286
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:286
	// _ = "end of CoverTab[117406]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:286
	_go_fuzz_dep_.CoverTab[117407]++
											return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:287
	// _ = "end of CoverTab[117407]"
}

func (u *CacheOnReadFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:290
	_go_fuzz_dep_.CoverTab[117410]++
											return "CacheOnReadFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:291
	// _ = "end of CoverTab[117410]"
}

func (u *CacheOnReadFs) MkdirAll(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:294
	_go_fuzz_dep_.CoverTab[117411]++
											err := u.base.MkdirAll(name, perm)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:296
		_go_fuzz_dep_.CoverTab[117413]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:297
		// _ = "end of CoverTab[117413]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:298
		_go_fuzz_dep_.CoverTab[117414]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:298
		// _ = "end of CoverTab[117414]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:298
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:298
	// _ = "end of CoverTab[117411]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:298
	_go_fuzz_dep_.CoverTab[117412]++
											return u.layer.MkdirAll(name, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:299
	// _ = "end of CoverTab[117412]"
}

func (u *CacheOnReadFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:302
	_go_fuzz_dep_.CoverTab[117415]++
											bfh, err := u.base.Create(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:304
		_go_fuzz_dep_.CoverTab[117418]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:305
		// _ = "end of CoverTab[117418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:306
		_go_fuzz_dep_.CoverTab[117419]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:306
		// _ = "end of CoverTab[117419]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:306
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:306
	// _ = "end of CoverTab[117415]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:306
	_go_fuzz_dep_.CoverTab[117416]++
											lfh, err := u.layer.Create(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:308
		_go_fuzz_dep_.CoverTab[117420]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:311
		bfh.Close()
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:312
		// _ = "end of CoverTab[117420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:313
		_go_fuzz_dep_.CoverTab[117421]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:313
		// _ = "end of CoverTab[117421]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:313
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:313
	// _ = "end of CoverTab[117416]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:313
	_go_fuzz_dep_.CoverTab[117417]++
											return &UnionFile{Base: bfh, Layer: lfh}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:314
	// _ = "end of CoverTab[117417]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/cacheOnReadFs.go:315
var _ = _go_fuzz_dep_.CoverTab
