//go:build go1.16
// +build go1.16

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:4
)

import (
	"io"
	"io/fs"
	"os"
	"path"
	"sort"
	"time"

	"github.com/spf13/afero/internal/common"
)

// IOFS adopts afero.Fs to stdlib io/fs.FS
type IOFS struct {
	Fs
}

func NewIOFS(fs Fs) IOFS {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:22
	_go_fuzz_dep_.CoverTab[117631]++
										return IOFS{Fs: fs}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:23
	// _ = "end of CoverTab[117631]"
}

var (
	_	fs.FS		= IOFS{}
	_	fs.GlobFS	= IOFS{}
	_	fs.ReadDirFS	= IOFS{}
	_	fs.ReadFileFS	= IOFS{}
	_	fs.StatFS	= IOFS{}
	_	fs.SubFS	= IOFS{}
)

func (iofs IOFS) Open(name string) (fs.File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:35
	_go_fuzz_dep_.CoverTab[117632]++
										const op = "open"

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:39
	if !fs.ValidPath(name) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:39
		_go_fuzz_dep_.CoverTab[117636]++
											return nil, iofs.wrapError(op, name, fs.ErrInvalid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:40
		// _ = "end of CoverTab[117636]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:41
		_go_fuzz_dep_.CoverTab[117637]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:41
		// _ = "end of CoverTab[117637]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:41
	// _ = "end of CoverTab[117632]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:41
	_go_fuzz_dep_.CoverTab[117633]++

										file, err := iofs.Fs.Open(name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:44
		_go_fuzz_dep_.CoverTab[117638]++
											return nil, iofs.wrapError(op, name, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:45
		// _ = "end of CoverTab[117638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:46
		_go_fuzz_dep_.CoverTab[117639]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:46
		// _ = "end of CoverTab[117639]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:46
	// _ = "end of CoverTab[117633]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:46
	_go_fuzz_dep_.CoverTab[117634]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:49
	if _, ok := file.(fs.ReadDirFile); !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:49
		_go_fuzz_dep_.CoverTab[117640]++
											file = readDirFile{file}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:50
		// _ = "end of CoverTab[117640]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:51
		_go_fuzz_dep_.CoverTab[117641]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:51
		// _ = "end of CoverTab[117641]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:51
	// _ = "end of CoverTab[117634]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:51
	_go_fuzz_dep_.CoverTab[117635]++

										return file, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:53
	// _ = "end of CoverTab[117635]"
}

func (iofs IOFS) Glob(pattern string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:56
	_go_fuzz_dep_.CoverTab[117642]++
										const op = "glob"

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:60
	if _, err := path.Match(pattern, ""); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:60
		_go_fuzz_dep_.CoverTab[117645]++
											return nil, iofs.wrapError(op, pattern, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:61
		// _ = "end of CoverTab[117645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:62
		_go_fuzz_dep_.CoverTab[117646]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:62
		// _ = "end of CoverTab[117646]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:62
	// _ = "end of CoverTab[117642]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:62
	_go_fuzz_dep_.CoverTab[117643]++

										items, err := Glob(iofs.Fs, pattern)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:65
		_go_fuzz_dep_.CoverTab[117647]++
											return nil, iofs.wrapError(op, pattern, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:66
		// _ = "end of CoverTab[117647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:67
		_go_fuzz_dep_.CoverTab[117648]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:67
		// _ = "end of CoverTab[117648]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:67
	// _ = "end of CoverTab[117643]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:67
	_go_fuzz_dep_.CoverTab[117644]++

										return items, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:69
	// _ = "end of CoverTab[117644]"
}

func (iofs IOFS) ReadDir(name string) ([]fs.DirEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:72
	_go_fuzz_dep_.CoverTab[117649]++
										f, err := iofs.Fs.Open(name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:74
		_go_fuzz_dep_.CoverTab[117654]++
											return nil, iofs.wrapError("readdir", name, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:75
		// _ = "end of CoverTab[117654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:76
		_go_fuzz_dep_.CoverTab[117655]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:76
		// _ = "end of CoverTab[117655]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:76
	// _ = "end of CoverTab[117649]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:76
	_go_fuzz_dep_.CoverTab[117650]++

										defer f.Close()

										if rdf, ok := f.(fs.ReadDirFile); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:80
		_go_fuzz_dep_.CoverTab[117656]++
											items, err := rdf.ReadDir(-1)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:82
			_go_fuzz_dep_.CoverTab[117659]++
												return nil, iofs.wrapError("readdir", name, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:83
			// _ = "end of CoverTab[117659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:84
			_go_fuzz_dep_.CoverTab[117660]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:84
			// _ = "end of CoverTab[117660]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:84
		// _ = "end of CoverTab[117656]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:84
		_go_fuzz_dep_.CoverTab[117657]++
											sort.Slice(items, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
			_go_fuzz_dep_.CoverTab[117661]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
			return items[i].Name() < items[j].Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
			// _ = "end of CoverTab[117661]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
		})
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
		// _ = "end of CoverTab[117657]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:85
		_go_fuzz_dep_.CoverTab[117658]++
											return items, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:86
		// _ = "end of CoverTab[117658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:87
		_go_fuzz_dep_.CoverTab[117662]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:87
		// _ = "end of CoverTab[117662]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:87
	// _ = "end of CoverTab[117650]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:87
	_go_fuzz_dep_.CoverTab[117651]++

										items, err := f.Readdir(-1)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:90
		_go_fuzz_dep_.CoverTab[117663]++
											return nil, iofs.wrapError("readdir", name, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:91
		// _ = "end of CoverTab[117663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:92
		_go_fuzz_dep_.CoverTab[117664]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:92
		// _ = "end of CoverTab[117664]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:92
	// _ = "end of CoverTab[117651]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:92
	_go_fuzz_dep_.CoverTab[117652]++
										sort.Sort(byName(items))

										ret := make([]fs.DirEntry, len(items))
										for i := range items {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:96
		_go_fuzz_dep_.CoverTab[117665]++
											ret[i] = common.FileInfoDirEntry{FileInfo: items[i]}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:97
		// _ = "end of CoverTab[117665]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:98
	// _ = "end of CoverTab[117652]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:98
	_go_fuzz_dep_.CoverTab[117653]++

										return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:100
	// _ = "end of CoverTab[117653]"
}

func (iofs IOFS) ReadFile(name string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:103
	_go_fuzz_dep_.CoverTab[117666]++
										const op = "readfile"

										if !fs.ValidPath(name) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:106
		_go_fuzz_dep_.CoverTab[117669]++
											return nil, iofs.wrapError(op, name, fs.ErrInvalid)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:107
		// _ = "end of CoverTab[117669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:108
		_go_fuzz_dep_.CoverTab[117670]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:108
		// _ = "end of CoverTab[117670]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:108
	// _ = "end of CoverTab[117666]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:108
	_go_fuzz_dep_.CoverTab[117667]++

										bytes, err := ReadFile(iofs.Fs, name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:111
		_go_fuzz_dep_.CoverTab[117671]++
											return nil, iofs.wrapError(op, name, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:112
		// _ = "end of CoverTab[117671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:113
		_go_fuzz_dep_.CoverTab[117672]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:113
		// _ = "end of CoverTab[117672]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:113
	// _ = "end of CoverTab[117667]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:113
	_go_fuzz_dep_.CoverTab[117668]++

										return bytes, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:115
	// _ = "end of CoverTab[117668]"
}

func (iofs IOFS) Sub(dir string) (fs.FS, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:118
	_go_fuzz_dep_.CoverTab[117673]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:118
	return IOFS{NewBasePathFs(iofs.Fs, dir)}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:118
	// _ = "end of CoverTab[117673]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:118
}

func (IOFS) wrapError(op, path string, err error) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:120
	_go_fuzz_dep_.CoverTab[117674]++
										if _, ok := err.(*fs.PathError); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:121
		_go_fuzz_dep_.CoverTab[117676]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:122
		// _ = "end of CoverTab[117676]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:123
		_go_fuzz_dep_.CoverTab[117677]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:123
		// _ = "end of CoverTab[117677]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:123
	// _ = "end of CoverTab[117674]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:123
	_go_fuzz_dep_.CoverTab[117675]++

										return &fs.PathError{
		Op:	op,
		Path:	path,
		Err:	err,
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:129
	// _ = "end of CoverTab[117675]"
}

// readDirFile provides adapter from afero.File to fs.ReadDirFile needed for correct Open
type readDirFile struct {
	File
}

var _ fs.ReadDirFile = readDirFile{}

func (r readDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:139
	_go_fuzz_dep_.CoverTab[117678]++
										items, err := r.File.Readdir(n)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:141
		_go_fuzz_dep_.CoverTab[117681]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:142
		// _ = "end of CoverTab[117681]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:143
		_go_fuzz_dep_.CoverTab[117682]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:143
		// _ = "end of CoverTab[117682]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:143
	// _ = "end of CoverTab[117678]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:143
	_go_fuzz_dep_.CoverTab[117679]++

										ret := make([]fs.DirEntry, len(items))
										for i := range items {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:146
		_go_fuzz_dep_.CoverTab[117683]++
											ret[i] = common.FileInfoDirEntry{FileInfo: items[i]}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:147
		// _ = "end of CoverTab[117683]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:148
	// _ = "end of CoverTab[117679]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:148
	_go_fuzz_dep_.CoverTab[117680]++

										return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:150
	// _ = "end of CoverTab[117680]"
}

// FromIOFS adopts io/fs.FS to use it as afero.Fs
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:153
// Note that io/fs.FS is read-only so all mutating methods will return fs.PathError with fs.ErrPermission
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:153
// To store modifications you may use afero.CopyOnWriteFs
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:156
type FromIOFS struct {
	fs.FS
}

var _ Fs = FromIOFS{}

func (f FromIOFS) Create(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:162
	_go_fuzz_dep_.CoverTab[117684]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:162
	return nil, notImplemented("create", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:162
	// _ = "end of CoverTab[117684]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:162
}

func (f FromIOFS) Mkdir(name string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:164
	_go_fuzz_dep_.CoverTab[117685]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:164
	return notImplemented("mkdir", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:164
	// _ = "end of CoverTab[117685]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:164
}

func (f FromIOFS) MkdirAll(path string, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:166
	_go_fuzz_dep_.CoverTab[117686]++
										return notImplemented("mkdirall", path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:167
	// _ = "end of CoverTab[117686]"
}

func (f FromIOFS) Open(name string) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:170
	_go_fuzz_dep_.CoverTab[117687]++
										file, err := f.FS.Open(name)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:172
		_go_fuzz_dep_.CoverTab[117689]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:173
		// _ = "end of CoverTab[117689]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:174
		_go_fuzz_dep_.CoverTab[117690]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:174
		// _ = "end of CoverTab[117690]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:174
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:174
	// _ = "end of CoverTab[117687]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:174
	_go_fuzz_dep_.CoverTab[117688]++

										return fromIOFSFile{File: file, name: name}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:176
	// _ = "end of CoverTab[117688]"
}

func (f FromIOFS) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:179
	_go_fuzz_dep_.CoverTab[117691]++
										return f.Open(name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:180
	// _ = "end of CoverTab[117691]"
}

func (f FromIOFS) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:183
	_go_fuzz_dep_.CoverTab[117692]++
										return notImplemented("remove", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:184
	// _ = "end of CoverTab[117692]"
}

func (f FromIOFS) RemoveAll(path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:187
	_go_fuzz_dep_.CoverTab[117693]++
										return notImplemented("removeall", path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:188
	// _ = "end of CoverTab[117693]"
}

func (f FromIOFS) Rename(oldname, newname string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:191
	_go_fuzz_dep_.CoverTab[117694]++
										return notImplemented("rename", oldname)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:192
	// _ = "end of CoverTab[117694]"
}

func (f FromIOFS) Stat(name string) (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:195
	_go_fuzz_dep_.CoverTab[117695]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:195
	return fs.Stat(f.FS, name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:195
	// _ = "end of CoverTab[117695]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:195
}

func (f FromIOFS) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:197
	_go_fuzz_dep_.CoverTab[117696]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:197
	return "fromiofs"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:197
	// _ = "end of CoverTab[117696]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:197
}

func (f FromIOFS) Chmod(name string, mode os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:199
	_go_fuzz_dep_.CoverTab[117697]++
										return notImplemented("chmod", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:200
	// _ = "end of CoverTab[117697]"
}

func (f FromIOFS) Chown(name string, uid, gid int) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:203
	_go_fuzz_dep_.CoverTab[117698]++
										return notImplemented("chown", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:204
	// _ = "end of CoverTab[117698]"
}

func (f FromIOFS) Chtimes(name string, atime time.Time, mtime time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:207
	_go_fuzz_dep_.CoverTab[117699]++
										return notImplemented("chtimes", name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:208
	// _ = "end of CoverTab[117699]"
}

type fromIOFSFile struct {
	fs.File
	name	string
}

func (f fromIOFSFile) ReadAt(p []byte, off int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:216
	_go_fuzz_dep_.CoverTab[117700]++
										readerAt, ok := f.File.(io.ReaderAt)
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:218
		_go_fuzz_dep_.CoverTab[117702]++
											return -1, notImplemented("readat", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:219
		// _ = "end of CoverTab[117702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:220
		_go_fuzz_dep_.CoverTab[117703]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:220
		// _ = "end of CoverTab[117703]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:220
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:220
	// _ = "end of CoverTab[117700]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:220
	_go_fuzz_dep_.CoverTab[117701]++

										return readerAt.ReadAt(p, off)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:222
	// _ = "end of CoverTab[117701]"
}

func (f fromIOFSFile) Seek(offset int64, whence int) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:225
	_go_fuzz_dep_.CoverTab[117704]++
										seeker, ok := f.File.(io.Seeker)
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:227
		_go_fuzz_dep_.CoverTab[117706]++
											return -1, notImplemented("seek", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:228
		// _ = "end of CoverTab[117706]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:229
		_go_fuzz_dep_.CoverTab[117707]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:229
		// _ = "end of CoverTab[117707]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:229
	// _ = "end of CoverTab[117704]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:229
	_go_fuzz_dep_.CoverTab[117705]++

										return seeker.Seek(offset, whence)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:231
	// _ = "end of CoverTab[117705]"
}

func (f fromIOFSFile) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:234
	_go_fuzz_dep_.CoverTab[117708]++
										return -1, notImplemented("write", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:235
	// _ = "end of CoverTab[117708]"
}

func (f fromIOFSFile) WriteAt(p []byte, off int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:238
	_go_fuzz_dep_.CoverTab[117709]++
										return -1, notImplemented("writeat", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:239
	// _ = "end of CoverTab[117709]"
}

func (f fromIOFSFile) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:242
	_go_fuzz_dep_.CoverTab[117710]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:242
	return f.name
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:242
	// _ = "end of CoverTab[117710]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:242
}

func (f fromIOFSFile) Readdir(count int) ([]os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:244
	_go_fuzz_dep_.CoverTab[117711]++
										rdfile, ok := f.File.(fs.ReadDirFile)
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:246
		_go_fuzz_dep_.CoverTab[117715]++
											return nil, notImplemented("readdir", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:247
		// _ = "end of CoverTab[117715]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:248
		_go_fuzz_dep_.CoverTab[117716]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:248
		// _ = "end of CoverTab[117716]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:248
	// _ = "end of CoverTab[117711]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:248
	_go_fuzz_dep_.CoverTab[117712]++

										entries, err := rdfile.ReadDir(count)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:251
		_go_fuzz_dep_.CoverTab[117717]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:252
		// _ = "end of CoverTab[117717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:253
		_go_fuzz_dep_.CoverTab[117718]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:253
		// _ = "end of CoverTab[117718]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:253
	// _ = "end of CoverTab[117712]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:253
	_go_fuzz_dep_.CoverTab[117713]++

										ret := make([]os.FileInfo, len(entries))
										for i := range entries {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:256
		_go_fuzz_dep_.CoverTab[117719]++
											ret[i], err = entries[i].Info()

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:259
			_go_fuzz_dep_.CoverTab[117720]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:260
			// _ = "end of CoverTab[117720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:261
			_go_fuzz_dep_.CoverTab[117721]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:261
			// _ = "end of CoverTab[117721]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:261
		// _ = "end of CoverTab[117719]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:262
	// _ = "end of CoverTab[117713]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:262
	_go_fuzz_dep_.CoverTab[117714]++

										return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:264
	// _ = "end of CoverTab[117714]"
}

func (f fromIOFSFile) Readdirnames(n int) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:267
	_go_fuzz_dep_.CoverTab[117722]++
										rdfile, ok := f.File.(fs.ReadDirFile)
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:269
		_go_fuzz_dep_.CoverTab[117726]++
											return nil, notImplemented("readdir", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:270
		// _ = "end of CoverTab[117726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:271
		_go_fuzz_dep_.CoverTab[117727]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:271
		// _ = "end of CoverTab[117727]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:271
	// _ = "end of CoverTab[117722]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:271
	_go_fuzz_dep_.CoverTab[117723]++

										entries, err := rdfile.ReadDir(n)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:274
		_go_fuzz_dep_.CoverTab[117728]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:275
		// _ = "end of CoverTab[117728]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:276
		_go_fuzz_dep_.CoverTab[117729]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:276
		// _ = "end of CoverTab[117729]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:276
	// _ = "end of CoverTab[117723]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:276
	_go_fuzz_dep_.CoverTab[117724]++

										ret := make([]string, len(entries))
										for i := range entries {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:279
		_go_fuzz_dep_.CoverTab[117730]++
											ret[i] = entries[i].Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:280
		// _ = "end of CoverTab[117730]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:281
	// _ = "end of CoverTab[117724]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:281
	_go_fuzz_dep_.CoverTab[117725]++

										return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:283
	// _ = "end of CoverTab[117725]"
}

func (f fromIOFSFile) Sync() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:286
	_go_fuzz_dep_.CoverTab[117731]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:286
	return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:286
	// _ = "end of CoverTab[117731]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:286
}

func (f fromIOFSFile) Truncate(size int64) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:288
	_go_fuzz_dep_.CoverTab[117732]++
										return notImplemented("truncate", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:289
	// _ = "end of CoverTab[117732]"
}

func (f fromIOFSFile) WriteString(s string) (ret int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:292
	_go_fuzz_dep_.CoverTab[117733]++
										return -1, notImplemented("writestring", f.name)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:293
	// _ = "end of CoverTab[117733]"
}

func notImplemented(op, path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:296
	_go_fuzz_dep_.CoverTab[117734]++
										return &fs.PathError{Op: op, Path: path, Err: fs.ErrPermission}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:297
	// _ = "end of CoverTab[117734]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:298
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/iofs.go:298
var _ = _go_fuzz_dep_.CoverTab
