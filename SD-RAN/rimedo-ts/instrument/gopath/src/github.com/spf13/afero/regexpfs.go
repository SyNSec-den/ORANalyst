//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:1
)

import (
	"os"
	"regexp"
	"syscall"
	"time"
)

// The RegexpFs filters files (not directories) by regular expression. Only
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:10
// files matching the given regexp will be allowed, all others get a ENOENT error (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:10
// "No such file or directory").
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:14
type RegexpFs struct {
	re	*regexp.Regexp
	source	Fs
}

func NewRegexpFs(source Fs, re *regexp.Regexp) Fs {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:19
	_go_fuzz_dep_.CoverTab[118086]++
										return &RegexpFs{source: source, re: re}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:20
	// _ = "end of CoverTab[118086]"
}

type RegexpFile struct {
	f	File
	re	*regexp.Regexp
}

func (r *RegexpFs) matchesName(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:28
	_go_fuzz_dep_.CoverTab[118087]++
										if r.re == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:29
		_go_fuzz_dep_.CoverTab[118090]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:30
		// _ = "end of CoverTab[118090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:31
		_go_fuzz_dep_.CoverTab[118091]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:31
		// _ = "end of CoverTab[118091]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:31
	// _ = "end of CoverTab[118087]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:31
	_go_fuzz_dep_.CoverTab[118088]++
										if r.re.MatchString(name) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:32
		_go_fuzz_dep_.CoverTab[118092]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:33
		// _ = "end of CoverTab[118092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:34
		_go_fuzz_dep_.CoverTab[118093]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:34
		// _ = "end of CoverTab[118093]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:34
	// _ = "end of CoverTab[118088]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:34
	_go_fuzz_dep_.CoverTab[118089]++
										return syscall.ENOENT
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:35
	// _ = "end of CoverTab[118089]"
}

func (r *RegexpFs) dirOrMatches(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:38
	_go_fuzz_dep_.CoverTab[118094]++
										dir, err := IsDir(r.source, name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:40
		_go_fuzz_dep_.CoverTab[118097]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:41
		// _ = "end of CoverTab[118097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:42
		_go_fuzz_dep_.CoverTab[118098]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:42
		// _ = "end of CoverTab[118098]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:42
	// _ = "end of CoverTab[118094]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:42
	_go_fuzz_dep_.CoverTab[118095]++
										if dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:43
		_go_fuzz_dep_.CoverTab[118099]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:44
		// _ = "end of CoverTab[118099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:45
		_go_fuzz_dep_.CoverTab[118100]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:45
		// _ = "end of CoverTab[118100]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:45
	// _ = "end of CoverTab[118095]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:45
	_go_fuzz_dep_.CoverTab[118096]++
										return r.matchesName(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:46
	// _ = "end of CoverTab[118096]"
}

func (r *RegexpFs) Chtimes(name string, a, m time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:49
	_go_fuzz_dep_.CoverTab[118101]++
										if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:50
		_go_fuzz_dep_.CoverTab[118103]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:51
		// _ = "end of CoverTab[118103]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:52
		_go_fuzz_dep_.CoverTab[118104]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:52
		// _ = "end of CoverTab[118104]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:52
	// _ = "end of CoverTab[118101]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:52
	_go_fuzz_dep_.CoverTab[118102]++
										return r.source.Chtimes(name, a, m)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:53
	// _ = "end of CoverTab[118102]"
}

func (r *RegexpFs) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:56
	_go_fuzz_dep_.CoverTab[118105]++
										if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:57
		_go_fuzz_dep_.CoverTab[118107]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:58
		// _ = "end of CoverTab[118107]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:59
		_go_fuzz_dep_.CoverTab[118108]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:59
		// _ = "end of CoverTab[118108]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:59
	// _ = "end of CoverTab[118105]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:59
	_go_fuzz_dep_.CoverTab[118106]++
										return r.source.Chmod(name, mode)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:60
	// _ = "end of CoverTab[118106]"
}

func (r *RegexpFs) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:63
	_go_fuzz_dep_.CoverTab[118109]++
										if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:64
		_go_fuzz_dep_.CoverTab[118111]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:65
		// _ = "end of CoverTab[118111]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:66
		_go_fuzz_dep_.CoverTab[118112]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:66
		// _ = "end of CoverTab[118112]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:66
	// _ = "end of CoverTab[118109]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:66
	_go_fuzz_dep_.CoverTab[118110]++
										return r.source.Chown(name, uid, gid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:67
	// _ = "end of CoverTab[118110]"
}

func (r *RegexpFs) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:70
	_go_fuzz_dep_.CoverTab[118113]++
										return "RegexpFs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:71
	// _ = "end of CoverTab[118113]"
}

func (r *RegexpFs) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:74
	_go_fuzz_dep_.CoverTab[118114]++
										if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:75
		_go_fuzz_dep_.CoverTab[118116]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:76
		// _ = "end of CoverTab[118116]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:77
		_go_fuzz_dep_.CoverTab[118117]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:77
		// _ = "end of CoverTab[118117]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:77
	// _ = "end of CoverTab[118114]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:77
	_go_fuzz_dep_.CoverTab[118115]++
										return r.source.Stat(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:78
	// _ = "end of CoverTab[118115]"
}

func (r *RegexpFs) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:81
	_go_fuzz_dep_.CoverTab[118118]++
										dir, err := IsDir(r.source, oldname)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:83
		_go_fuzz_dep_.CoverTab[118123]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:84
		// _ = "end of CoverTab[118123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:85
		_go_fuzz_dep_.CoverTab[118124]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:85
		// _ = "end of CoverTab[118124]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:85
	// _ = "end of CoverTab[118118]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:85
	_go_fuzz_dep_.CoverTab[118119]++
										if dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:86
		_go_fuzz_dep_.CoverTab[118125]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:87
		// _ = "end of CoverTab[118125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:88
		_go_fuzz_dep_.CoverTab[118126]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:88
		// _ = "end of CoverTab[118126]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:88
	// _ = "end of CoverTab[118119]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:88
	_go_fuzz_dep_.CoverTab[118120]++
										if err := r.matchesName(oldname); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:89
		_go_fuzz_dep_.CoverTab[118127]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:90
		// _ = "end of CoverTab[118127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:91
		_go_fuzz_dep_.CoverTab[118128]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:91
		// _ = "end of CoverTab[118128]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:91
	// _ = "end of CoverTab[118120]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:91
	_go_fuzz_dep_.CoverTab[118121]++
										if err := r.matchesName(newname); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:92
		_go_fuzz_dep_.CoverTab[118129]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:93
		// _ = "end of CoverTab[118129]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:94
		_go_fuzz_dep_.CoverTab[118130]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:94
		// _ = "end of CoverTab[118130]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:94
	// _ = "end of CoverTab[118121]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:94
	_go_fuzz_dep_.CoverTab[118122]++
										return r.source.Rename(oldname, newname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:95
	// _ = "end of CoverTab[118122]"
}

func (r *RegexpFs) RemoveAll(p string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:98
	_go_fuzz_dep_.CoverTab[118131]++
										dir, err := IsDir(r.source, p)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:100
		_go_fuzz_dep_.CoverTab[118134]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:101
		// _ = "end of CoverTab[118134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:102
		_go_fuzz_dep_.CoverTab[118135]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:102
		// _ = "end of CoverTab[118135]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:102
	// _ = "end of CoverTab[118131]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:102
	_go_fuzz_dep_.CoverTab[118132]++
											if !dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:103
		_go_fuzz_dep_.CoverTab[118136]++
												if err := r.matchesName(p); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:104
			_go_fuzz_dep_.CoverTab[118137]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:105
			// _ = "end of CoverTab[118137]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:106
			_go_fuzz_dep_.CoverTab[118138]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:106
			// _ = "end of CoverTab[118138]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:106
		// _ = "end of CoverTab[118136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:107
		_go_fuzz_dep_.CoverTab[118139]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:107
		// _ = "end of CoverTab[118139]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:107
	// _ = "end of CoverTab[118132]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:107
	_go_fuzz_dep_.CoverTab[118133]++
											return r.source.RemoveAll(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:108
	// _ = "end of CoverTab[118133]"
}

func (r *RegexpFs) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:111
	_go_fuzz_dep_.CoverTab[118140]++
											if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:112
		_go_fuzz_dep_.CoverTab[118142]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:113
		// _ = "end of CoverTab[118142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:114
		_go_fuzz_dep_.CoverTab[118143]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:114
		// _ = "end of CoverTab[118143]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:114
	// _ = "end of CoverTab[118140]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:114
	_go_fuzz_dep_.CoverTab[118141]++
											return r.source.Remove(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:115
	// _ = "end of CoverTab[118141]"
}

func (r *RegexpFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:118
	_go_fuzz_dep_.CoverTab[118144]++
											if err := r.dirOrMatches(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:119
		_go_fuzz_dep_.CoverTab[118146]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:120
		// _ = "end of CoverTab[118146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:121
		_go_fuzz_dep_.CoverTab[118147]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:121
		// _ = "end of CoverTab[118147]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:121
	// _ = "end of CoverTab[118144]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:121
	_go_fuzz_dep_.CoverTab[118145]++
											return r.source.OpenFile(name, flag, perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:122
	// _ = "end of CoverTab[118145]"
}

func (r *RegexpFs) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:125
	_go_fuzz_dep_.CoverTab[118148]++
											dir, err := IsDir(r.source, name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:127
		_go_fuzz_dep_.CoverTab[118152]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:128
		// _ = "end of CoverTab[118152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:129
		_go_fuzz_dep_.CoverTab[118153]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:129
		// _ = "end of CoverTab[118153]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:129
	// _ = "end of CoverTab[118148]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:129
	_go_fuzz_dep_.CoverTab[118149]++
											if !dir {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:130
		_go_fuzz_dep_.CoverTab[118154]++
												if err := r.matchesName(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:131
			_go_fuzz_dep_.CoverTab[118155]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:132
			// _ = "end of CoverTab[118155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:133
			_go_fuzz_dep_.CoverTab[118156]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:133
			// _ = "end of CoverTab[118156]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:133
		// _ = "end of CoverTab[118154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:134
		_go_fuzz_dep_.CoverTab[118157]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:134
		// _ = "end of CoverTab[118157]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:134
	// _ = "end of CoverTab[118149]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:134
	_go_fuzz_dep_.CoverTab[118150]++
											f, err := r.source.Open(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:136
		_go_fuzz_dep_.CoverTab[118158]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:137
		// _ = "end of CoverTab[118158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:138
		_go_fuzz_dep_.CoverTab[118159]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:138
		// _ = "end of CoverTab[118159]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:138
	// _ = "end of CoverTab[118150]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:138
	_go_fuzz_dep_.CoverTab[118151]++
											return &RegexpFile{f: f, re: r.re}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:139
	// _ = "end of CoverTab[118151]"
}

func (r *RegexpFs) Mkdir(n string, p os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:142
	_go_fuzz_dep_.CoverTab[118160]++
											return r.source.Mkdir(n, p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:143
	// _ = "end of CoverTab[118160]"
}

func (r *RegexpFs) MkdirAll(n string, p os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:146
	_go_fuzz_dep_.CoverTab[118161]++
											return r.source.MkdirAll(n, p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:147
	// _ = "end of CoverTab[118161]"
}

func (r *RegexpFs) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:150
	_go_fuzz_dep_.CoverTab[118162]++
											if err := r.matchesName(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:151
		_go_fuzz_dep_.CoverTab[118164]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:152
		// _ = "end of CoverTab[118164]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:153
		_go_fuzz_dep_.CoverTab[118165]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:153
		// _ = "end of CoverTab[118165]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:153
	// _ = "end of CoverTab[118162]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:153
	_go_fuzz_dep_.CoverTab[118163]++
											return r.source.Create(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:154
	// _ = "end of CoverTab[118163]"
}

func (f *RegexpFile) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:157
	_go_fuzz_dep_.CoverTab[118166]++
											return f.f.Close()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:158
	// _ = "end of CoverTab[118166]"
}

func (f *RegexpFile) Read(s []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:161
	_go_fuzz_dep_.CoverTab[118167]++
											return f.f.Read(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:162
	// _ = "end of CoverTab[118167]"
}

func (f *RegexpFile) ReadAt(s []byte, o int64) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:165
	_go_fuzz_dep_.CoverTab[118168]++
											return f.f.ReadAt(s, o)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:166
	// _ = "end of CoverTab[118168]"
}

func (f *RegexpFile) Seek(o int64, w int) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:169
	_go_fuzz_dep_.CoverTab[118169]++
											return f.f.Seek(o, w)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:170
	// _ = "end of CoverTab[118169]"
}

func (f *RegexpFile) Write(s []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:173
	_go_fuzz_dep_.CoverTab[118170]++
											return f.f.Write(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:174
	// _ = "end of CoverTab[118170]"
}

func (f *RegexpFile) WriteAt(s []byte, o int64) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:177
	_go_fuzz_dep_.CoverTab[118171]++
											return f.f.WriteAt(s, o)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:178
	// _ = "end of CoverTab[118171]"
}

func (f *RegexpFile) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:181
	_go_fuzz_dep_.CoverTab[118172]++
											return f.f.Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:182
	// _ = "end of CoverTab[118172]"
}

func (f *RegexpFile) Readdir(c int) (fi []os.FileInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:185
	_go_fuzz_dep_.CoverTab[118173]++
											var rfi []os.FileInfo
											rfi, err = f.f.Readdir(c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:188
		_go_fuzz_dep_.CoverTab[118176]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:189
		// _ = "end of CoverTab[118176]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:190
		_go_fuzz_dep_.CoverTab[118177]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:190
		// _ = "end of CoverTab[118177]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:190
	// _ = "end of CoverTab[118173]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:190
	_go_fuzz_dep_.CoverTab[118174]++
											for _, i := range rfi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:191
		_go_fuzz_dep_.CoverTab[118178]++
												if i.IsDir() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:192
			_go_fuzz_dep_.CoverTab[118179]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:192
			return f.re.MatchString(i.Name())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:192
			// _ = "end of CoverTab[118179]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:192
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:192
			_go_fuzz_dep_.CoverTab[118180]++
													fi = append(fi, i)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:193
			// _ = "end of CoverTab[118180]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:194
			_go_fuzz_dep_.CoverTab[118181]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:194
			// _ = "end of CoverTab[118181]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:194
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:194
		// _ = "end of CoverTab[118178]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:195
	// _ = "end of CoverTab[118174]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:195
	_go_fuzz_dep_.CoverTab[118175]++
											return fi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:196
	// _ = "end of CoverTab[118175]"
}

func (f *RegexpFile) Readdirnames(c int) (n []string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:199
	_go_fuzz_dep_.CoverTab[118182]++
											fi, err := f.Readdir(c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:201
		_go_fuzz_dep_.CoverTab[118185]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:202
		// _ = "end of CoverTab[118185]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:203
		_go_fuzz_dep_.CoverTab[118186]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:203
		// _ = "end of CoverTab[118186]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:203
	// _ = "end of CoverTab[118182]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:203
	_go_fuzz_dep_.CoverTab[118183]++
											for _, s := range fi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:204
		_go_fuzz_dep_.CoverTab[118187]++
												n = append(n, s.Name())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:205
		// _ = "end of CoverTab[118187]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:206
	// _ = "end of CoverTab[118183]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:206
	_go_fuzz_dep_.CoverTab[118184]++
											return n, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:207
	// _ = "end of CoverTab[118184]"
}

func (f *RegexpFile) Stat() (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:210
	_go_fuzz_dep_.CoverTab[118188]++
											return f.f.Stat()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:211
	// _ = "end of CoverTab[118188]"
}

func (f *RegexpFile) Sync() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:214
	_go_fuzz_dep_.CoverTab[118189]++
											return f.f.Sync()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:215
	// _ = "end of CoverTab[118189]"
}

func (f *RegexpFile) Truncate(s int64) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:218
	_go_fuzz_dep_.CoverTab[118190]++
											return f.f.Truncate(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:219
	// _ = "end of CoverTab[118190]"
}

func (f *RegexpFile) WriteString(s string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:222
	_go_fuzz_dep_.CoverTab[118191]++
											return f.f.WriteString(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:223
	// _ = "end of CoverTab[118191]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:224
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/regexpfs.go:224
var _ = _go_fuzz_dep_.CoverTab
