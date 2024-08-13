//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
package afero

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:1
)

import (
	"io"
	"os"
	"path/filepath"
	"syscall"
)

// The UnionFile implements the afero.File interface and will be returned
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// when reading a directory present at least in the overlay or opening a file
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// for writing.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// The calls to
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// Readdir() and Readdirnames() merge the file os.FileInfo / names from the
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// base and the overlay - for files present in both layers, only those
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// from the overlay will be used.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// When opening files for writing (Create() / OpenFile() with the right flags)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// the operations will be done in both layers, starting with the overlay. A
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// successful read in the overlay will move the cursor position in the base layer
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:10
// by the number of bytes read.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:23
type UnionFile struct {
	Base	File
	Layer	File
	Merger	DirsMerger
	off	int
	files	[]os.FileInfo
}

func (f *UnionFile) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:31
	_go_fuzz_dep_.CoverTab[118192]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:35
	if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:35
		_go_fuzz_dep_.CoverTab[118195]++
												f.Base.Close()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:36
		// _ = "end of CoverTab[118195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:37
		_go_fuzz_dep_.CoverTab[118196]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:37
		// _ = "end of CoverTab[118196]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:37
	// _ = "end of CoverTab[118192]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:37
	_go_fuzz_dep_.CoverTab[118193]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:38
		_go_fuzz_dep_.CoverTab[118197]++
												return f.Layer.Close()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:39
		// _ = "end of CoverTab[118197]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:40
		_go_fuzz_dep_.CoverTab[118198]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:40
		// _ = "end of CoverTab[118198]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:40
	// _ = "end of CoverTab[118193]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:40
	_go_fuzz_dep_.CoverTab[118194]++
											return BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:41
	// _ = "end of CoverTab[118194]"
}

func (f *UnionFile) Read(s []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:44
	_go_fuzz_dep_.CoverTab[118199]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:45
		_go_fuzz_dep_.CoverTab[118202]++
												n, err := f.Layer.Read(s)
												if (err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			_go_fuzz_dep_.CoverTab[118204]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			return err == io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			// _ = "end of CoverTab[118204]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			_go_fuzz_dep_.CoverTab[118205]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			// _ = "end of CoverTab[118205]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:47
			_go_fuzz_dep_.CoverTab[118206]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:50
			if _, seekErr := f.Base.Seek(int64(n), os.SEEK_CUR); seekErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:50
				_go_fuzz_dep_.CoverTab[118207]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:53
				err = seekErr
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:53
				// _ = "end of CoverTab[118207]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:54
				_go_fuzz_dep_.CoverTab[118208]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:54
				// _ = "end of CoverTab[118208]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:54
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:54
			// _ = "end of CoverTab[118206]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:55
			_go_fuzz_dep_.CoverTab[118209]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:55
			// _ = "end of CoverTab[118209]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:55
		// _ = "end of CoverTab[118202]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:55
		_go_fuzz_dep_.CoverTab[118203]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:56
		// _ = "end of CoverTab[118203]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:57
		_go_fuzz_dep_.CoverTab[118210]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:57
		// _ = "end of CoverTab[118210]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:57
	// _ = "end of CoverTab[118199]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:57
	_go_fuzz_dep_.CoverTab[118200]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:58
		_go_fuzz_dep_.CoverTab[118211]++
												return f.Base.Read(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:59
		// _ = "end of CoverTab[118211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:60
		_go_fuzz_dep_.CoverTab[118212]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:60
		// _ = "end of CoverTab[118212]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:60
	// _ = "end of CoverTab[118200]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:60
	_go_fuzz_dep_.CoverTab[118201]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:61
	// _ = "end of CoverTab[118201]"
}

func (f *UnionFile) ReadAt(s []byte, o int64) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:64
	_go_fuzz_dep_.CoverTab[118213]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:65
		_go_fuzz_dep_.CoverTab[118216]++
												n, err := f.Layer.ReadAt(s, o)
												if (err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			_go_fuzz_dep_.CoverTab[118218]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			return err == io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			// _ = "end of CoverTab[118218]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			_go_fuzz_dep_.CoverTab[118219]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			// _ = "end of CoverTab[118219]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:67
			_go_fuzz_dep_.CoverTab[118220]++
													_, err = f.Base.Seek(o+int64(n), io.SeekStart)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:68
			// _ = "end of CoverTab[118220]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:69
			_go_fuzz_dep_.CoverTab[118221]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:69
			// _ = "end of CoverTab[118221]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:69
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:69
		// _ = "end of CoverTab[118216]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:69
		_go_fuzz_dep_.CoverTab[118217]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:70
		// _ = "end of CoverTab[118217]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:71
		_go_fuzz_dep_.CoverTab[118222]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:71
		// _ = "end of CoverTab[118222]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:71
	// _ = "end of CoverTab[118213]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:71
	_go_fuzz_dep_.CoverTab[118214]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:72
		_go_fuzz_dep_.CoverTab[118223]++
												return f.Base.ReadAt(s, o)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:73
		// _ = "end of CoverTab[118223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:74
		_go_fuzz_dep_.CoverTab[118224]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:74
		// _ = "end of CoverTab[118224]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:74
	// _ = "end of CoverTab[118214]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:74
	_go_fuzz_dep_.CoverTab[118215]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:75
	// _ = "end of CoverTab[118215]"
}

func (f *UnionFile) Seek(o int64, w int) (pos int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:78
	_go_fuzz_dep_.CoverTab[118225]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:79
		_go_fuzz_dep_.CoverTab[118228]++
												pos, err = f.Layer.Seek(o, w)
												if (err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			_go_fuzz_dep_.CoverTab[118230]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			return err == io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			// _ = "end of CoverTab[118230]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			_go_fuzz_dep_.CoverTab[118231]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			// _ = "end of CoverTab[118231]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:81
			_go_fuzz_dep_.CoverTab[118232]++
													_, err = f.Base.Seek(o, w)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:82
			// _ = "end of CoverTab[118232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:83
			_go_fuzz_dep_.CoverTab[118233]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:83
			// _ = "end of CoverTab[118233]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:83
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:83
		// _ = "end of CoverTab[118228]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:83
		_go_fuzz_dep_.CoverTab[118229]++
												return pos, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:84
		// _ = "end of CoverTab[118229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:85
		_go_fuzz_dep_.CoverTab[118234]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:85
		// _ = "end of CoverTab[118234]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:85
	// _ = "end of CoverTab[118225]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:85
	_go_fuzz_dep_.CoverTab[118226]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:86
		_go_fuzz_dep_.CoverTab[118235]++
												return f.Base.Seek(o, w)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:87
		// _ = "end of CoverTab[118235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:88
		_go_fuzz_dep_.CoverTab[118236]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:88
		// _ = "end of CoverTab[118236]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:88
	// _ = "end of CoverTab[118226]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:88
	_go_fuzz_dep_.CoverTab[118227]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:89
	// _ = "end of CoverTab[118227]"
}

func (f *UnionFile) Write(s []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:92
	_go_fuzz_dep_.CoverTab[118237]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:93
		_go_fuzz_dep_.CoverTab[118240]++
												n, err = f.Layer.Write(s)
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:95
			_go_fuzz_dep_.CoverTab[118242]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:95
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:95
			// _ = "end of CoverTab[118242]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:95
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:95
			_go_fuzz_dep_.CoverTab[118243]++
													_, err = f.Base.Write(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:96
			// _ = "end of CoverTab[118243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:97
			_go_fuzz_dep_.CoverTab[118244]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:97
			// _ = "end of CoverTab[118244]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:97
		// _ = "end of CoverTab[118240]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:97
		_go_fuzz_dep_.CoverTab[118241]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:98
		// _ = "end of CoverTab[118241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:99
		_go_fuzz_dep_.CoverTab[118245]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:99
		// _ = "end of CoverTab[118245]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:99
	// _ = "end of CoverTab[118237]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:99
	_go_fuzz_dep_.CoverTab[118238]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:100
		_go_fuzz_dep_.CoverTab[118246]++
												return f.Base.Write(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:101
		// _ = "end of CoverTab[118246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:102
		_go_fuzz_dep_.CoverTab[118247]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:102
		// _ = "end of CoverTab[118247]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:102
	// _ = "end of CoverTab[118238]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:102
	_go_fuzz_dep_.CoverTab[118239]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:103
	// _ = "end of CoverTab[118239]"
}

func (f *UnionFile) WriteAt(s []byte, o int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:106
	_go_fuzz_dep_.CoverTab[118248]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:107
		_go_fuzz_dep_.CoverTab[118251]++
												n, err = f.Layer.WriteAt(s, o)
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:109
			_go_fuzz_dep_.CoverTab[118253]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:109
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:109
			// _ = "end of CoverTab[118253]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:109
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:109
			_go_fuzz_dep_.CoverTab[118254]++
													_, err = f.Base.WriteAt(s, o)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:110
			// _ = "end of CoverTab[118254]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:111
			_go_fuzz_dep_.CoverTab[118255]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:111
			// _ = "end of CoverTab[118255]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:111
		// _ = "end of CoverTab[118251]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:111
		_go_fuzz_dep_.CoverTab[118252]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:112
		// _ = "end of CoverTab[118252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:113
		_go_fuzz_dep_.CoverTab[118256]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:113
		// _ = "end of CoverTab[118256]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:113
	// _ = "end of CoverTab[118248]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:113
	_go_fuzz_dep_.CoverTab[118249]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:114
		_go_fuzz_dep_.CoverTab[118257]++
												return f.Base.WriteAt(s, o)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:115
		// _ = "end of CoverTab[118257]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:116
		_go_fuzz_dep_.CoverTab[118258]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:116
		// _ = "end of CoverTab[118258]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:116
	// _ = "end of CoverTab[118249]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:116
	_go_fuzz_dep_.CoverTab[118250]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:117
	// _ = "end of CoverTab[118250]"
}

func (f *UnionFile) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:120
	_go_fuzz_dep_.CoverTab[118259]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:121
		_go_fuzz_dep_.CoverTab[118261]++
												return f.Layer.Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:122
		// _ = "end of CoverTab[118261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:123
		_go_fuzz_dep_.CoverTab[118262]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:123
		// _ = "end of CoverTab[118262]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:123
	// _ = "end of CoverTab[118259]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:123
	_go_fuzz_dep_.CoverTab[118260]++
											return f.Base.Name()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:124
	// _ = "end of CoverTab[118260]"
}

// DirsMerger is how UnionFile weaves two directories together.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:127
// It takes the FileInfo slices from the layer and the base and returns a
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:127
// single view.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:130
type DirsMerger func(lofi, bofi []os.FileInfo) ([]os.FileInfo, error)

var defaultUnionMergeDirsFn = func(lofi, bofi []os.FileInfo) ([]os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:132
	_go_fuzz_dep_.CoverTab[118263]++
											var files = make(map[string]os.FileInfo)

											for _, fi := range lofi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:135
		_go_fuzz_dep_.CoverTab[118267]++
												files[fi.Name()] = fi
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:136
		// _ = "end of CoverTab[118267]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:137
	// _ = "end of CoverTab[118263]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:137
	_go_fuzz_dep_.CoverTab[118264]++

											for _, fi := range bofi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:139
		_go_fuzz_dep_.CoverTab[118268]++
												if _, exists := files[fi.Name()]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:140
			_go_fuzz_dep_.CoverTab[118269]++
													files[fi.Name()] = fi
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:141
			// _ = "end of CoverTab[118269]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:142
			_go_fuzz_dep_.CoverTab[118270]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:142
			// _ = "end of CoverTab[118270]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:142
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:142
		// _ = "end of CoverTab[118268]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:143
	// _ = "end of CoverTab[118264]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:143
	_go_fuzz_dep_.CoverTab[118265]++

											rfi := make([]os.FileInfo, len(files))

											i := 0
											for _, fi := range files {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:148
		_go_fuzz_dep_.CoverTab[118271]++
												rfi[i] = fi
												i++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:150
		// _ = "end of CoverTab[118271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:151
	// _ = "end of CoverTab[118265]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:151
	_go_fuzz_dep_.CoverTab[118266]++

											return rfi, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:153
	// _ = "end of CoverTab[118266]"

}

// Readdir will weave the two directories together and
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:157
// return a single view of the overlayed directories.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:157
// At the end of the directory view, the error is io.EOF if c > 0.
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:160
func (f *UnionFile) Readdir(c int) (ofi []os.FileInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:160
	_go_fuzz_dep_.CoverTab[118272]++
											var merge DirsMerger = f.Merger
											if merge == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:162
		_go_fuzz_dep_.CoverTab[118279]++
												merge = defaultUnionMergeDirsFn
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:163
		// _ = "end of CoverTab[118279]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:164
		_go_fuzz_dep_.CoverTab[118280]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:164
		// _ = "end of CoverTab[118280]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:164
	// _ = "end of CoverTab[118272]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:164
	_go_fuzz_dep_.CoverTab[118273]++

											if f.off == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:166
		_go_fuzz_dep_.CoverTab[118281]++
												var lfi []os.FileInfo
												if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:168
			_go_fuzz_dep_.CoverTab[118285]++
													lfi, err = f.Layer.Readdir(-1)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:170
				_go_fuzz_dep_.CoverTab[118286]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:171
				// _ = "end of CoverTab[118286]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:172
				_go_fuzz_dep_.CoverTab[118287]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:172
				// _ = "end of CoverTab[118287]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:172
			// _ = "end of CoverTab[118285]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:173
			_go_fuzz_dep_.CoverTab[118288]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:173
			// _ = "end of CoverTab[118288]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:173
		// _ = "end of CoverTab[118281]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:173
		_go_fuzz_dep_.CoverTab[118282]++

												var bfi []os.FileInfo
												if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:176
			_go_fuzz_dep_.CoverTab[118289]++
													bfi, err = f.Base.Readdir(-1)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:178
				_go_fuzz_dep_.CoverTab[118290]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:179
				// _ = "end of CoverTab[118290]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:180
				_go_fuzz_dep_.CoverTab[118291]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:180
				// _ = "end of CoverTab[118291]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:180
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:180
			// _ = "end of CoverTab[118289]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:182
			_go_fuzz_dep_.CoverTab[118292]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:182
			// _ = "end of CoverTab[118292]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:182
		// _ = "end of CoverTab[118282]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:182
		_go_fuzz_dep_.CoverTab[118283]++
												merged, err := merge(lfi, bfi)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:184
			_go_fuzz_dep_.CoverTab[118293]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:185
			// _ = "end of CoverTab[118293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:186
			_go_fuzz_dep_.CoverTab[118294]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:186
			// _ = "end of CoverTab[118294]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:186
		// _ = "end of CoverTab[118283]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:186
		_go_fuzz_dep_.CoverTab[118284]++
												f.files = append(f.files, merged...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:187
		// _ = "end of CoverTab[118284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:188
		_go_fuzz_dep_.CoverTab[118295]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:188
		// _ = "end of CoverTab[118295]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:188
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:188
	// _ = "end of CoverTab[118273]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:188
	_go_fuzz_dep_.CoverTab[118274]++
											files := f.files[f.off:]

											if c <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:191
		_go_fuzz_dep_.CoverTab[118296]++
												return files, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:192
		// _ = "end of CoverTab[118296]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:193
		_go_fuzz_dep_.CoverTab[118297]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:193
		// _ = "end of CoverTab[118297]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:193
	// _ = "end of CoverTab[118274]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:193
	_go_fuzz_dep_.CoverTab[118275]++

											if len(files) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:195
		_go_fuzz_dep_.CoverTab[118298]++
												return nil, io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:196
		// _ = "end of CoverTab[118298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:197
		_go_fuzz_dep_.CoverTab[118299]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:197
		// _ = "end of CoverTab[118299]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:197
	// _ = "end of CoverTab[118275]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:197
	_go_fuzz_dep_.CoverTab[118276]++

											if c > len(files) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:199
		_go_fuzz_dep_.CoverTab[118300]++
												c = len(files)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:200
		// _ = "end of CoverTab[118300]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:201
		_go_fuzz_dep_.CoverTab[118301]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:201
		// _ = "end of CoverTab[118301]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:201
	// _ = "end of CoverTab[118276]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:201
	_go_fuzz_dep_.CoverTab[118277]++

											defer func() { _go_fuzz_dep_.CoverTab[118302]++; f.off += c; // _ = "end of CoverTab[118302]" }()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:203
	// _ = "end of CoverTab[118277]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:203
	_go_fuzz_dep_.CoverTab[118278]++
											return files[:c], nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:204
	// _ = "end of CoverTab[118278]"
}

func (f *UnionFile) Readdirnames(c int) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:207
	_go_fuzz_dep_.CoverTab[118303]++
											rfi, err := f.Readdir(c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:209
		_go_fuzz_dep_.CoverTab[118306]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:210
		// _ = "end of CoverTab[118306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:211
		_go_fuzz_dep_.CoverTab[118307]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:211
		// _ = "end of CoverTab[118307]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:211
	// _ = "end of CoverTab[118303]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:211
	_go_fuzz_dep_.CoverTab[118304]++
											var names []string
											for _, fi := range rfi {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:213
		_go_fuzz_dep_.CoverTab[118308]++
												names = append(names, fi.Name())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:214
		// _ = "end of CoverTab[118308]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:215
	// _ = "end of CoverTab[118304]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:215
	_go_fuzz_dep_.CoverTab[118305]++
											return names, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:216
	// _ = "end of CoverTab[118305]"
}

func (f *UnionFile) Stat() (os.FileInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:219
	_go_fuzz_dep_.CoverTab[118309]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:220
		_go_fuzz_dep_.CoverTab[118312]++
												return f.Layer.Stat()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:221
		// _ = "end of CoverTab[118312]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:222
		_go_fuzz_dep_.CoverTab[118313]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:222
		// _ = "end of CoverTab[118313]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:222
	// _ = "end of CoverTab[118309]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:222
	_go_fuzz_dep_.CoverTab[118310]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:223
		_go_fuzz_dep_.CoverTab[118314]++
												return f.Base.Stat()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:224
		// _ = "end of CoverTab[118314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:225
		_go_fuzz_dep_.CoverTab[118315]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:225
		// _ = "end of CoverTab[118315]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:225
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:225
	// _ = "end of CoverTab[118310]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:225
	_go_fuzz_dep_.CoverTab[118311]++
											return nil, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:226
	// _ = "end of CoverTab[118311]"
}

func (f *UnionFile) Sync() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:229
	_go_fuzz_dep_.CoverTab[118316]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:230
		_go_fuzz_dep_.CoverTab[118319]++
												err = f.Layer.Sync()
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:232
			_go_fuzz_dep_.CoverTab[118321]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:232
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:232
			// _ = "end of CoverTab[118321]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:232
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:232
			_go_fuzz_dep_.CoverTab[118322]++
													err = f.Base.Sync()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:233
			// _ = "end of CoverTab[118322]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:234
			_go_fuzz_dep_.CoverTab[118323]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:234
			// _ = "end of CoverTab[118323]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:234
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:234
		// _ = "end of CoverTab[118319]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:234
		_go_fuzz_dep_.CoverTab[118320]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:235
		// _ = "end of CoverTab[118320]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:236
		_go_fuzz_dep_.CoverTab[118324]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:236
		// _ = "end of CoverTab[118324]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:236
	// _ = "end of CoverTab[118316]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:236
	_go_fuzz_dep_.CoverTab[118317]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:237
		_go_fuzz_dep_.CoverTab[118325]++
												return f.Base.Sync()
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:238
		// _ = "end of CoverTab[118325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:239
		_go_fuzz_dep_.CoverTab[118326]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:239
		// _ = "end of CoverTab[118326]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:239
	// _ = "end of CoverTab[118317]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:239
	_go_fuzz_dep_.CoverTab[118318]++
											return BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:240
	// _ = "end of CoverTab[118318]"
}

func (f *UnionFile) Truncate(s int64) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:243
	_go_fuzz_dep_.CoverTab[118327]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:244
		_go_fuzz_dep_.CoverTab[118330]++
												err = f.Layer.Truncate(s)
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:246
			_go_fuzz_dep_.CoverTab[118332]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:246
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:246
			// _ = "end of CoverTab[118332]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:246
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:246
			_go_fuzz_dep_.CoverTab[118333]++
													err = f.Base.Truncate(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:247
			// _ = "end of CoverTab[118333]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:248
			_go_fuzz_dep_.CoverTab[118334]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:248
			// _ = "end of CoverTab[118334]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:248
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:248
		// _ = "end of CoverTab[118330]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:248
		_go_fuzz_dep_.CoverTab[118331]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:249
		// _ = "end of CoverTab[118331]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:250
		_go_fuzz_dep_.CoverTab[118335]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:250
		// _ = "end of CoverTab[118335]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:250
	// _ = "end of CoverTab[118327]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:250
	_go_fuzz_dep_.CoverTab[118328]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:251
		_go_fuzz_dep_.CoverTab[118336]++
												return f.Base.Truncate(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:252
		// _ = "end of CoverTab[118336]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:253
		_go_fuzz_dep_.CoverTab[118337]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:253
		// _ = "end of CoverTab[118337]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:253
	// _ = "end of CoverTab[118328]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:253
	_go_fuzz_dep_.CoverTab[118329]++
											return BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:254
	// _ = "end of CoverTab[118329]"
}

func (f *UnionFile) WriteString(s string) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:257
	_go_fuzz_dep_.CoverTab[118338]++
											if f.Layer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:258
		_go_fuzz_dep_.CoverTab[118341]++
												n, err = f.Layer.WriteString(s)
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:260
			_go_fuzz_dep_.CoverTab[118343]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:260
			return f.Base != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:260
			// _ = "end of CoverTab[118343]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:260
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:260
			_go_fuzz_dep_.CoverTab[118344]++
													_, err = f.Base.WriteString(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:261
			// _ = "end of CoverTab[118344]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:262
			_go_fuzz_dep_.CoverTab[118345]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:262
			// _ = "end of CoverTab[118345]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:262
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:262
		// _ = "end of CoverTab[118341]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:262
		_go_fuzz_dep_.CoverTab[118342]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:263
		// _ = "end of CoverTab[118342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:264
		_go_fuzz_dep_.CoverTab[118346]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:264
		// _ = "end of CoverTab[118346]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:264
	// _ = "end of CoverTab[118338]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:264
	_go_fuzz_dep_.CoverTab[118339]++
											if f.Base != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:265
		_go_fuzz_dep_.CoverTab[118347]++
												return f.Base.WriteString(s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:266
		// _ = "end of CoverTab[118347]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:267
		_go_fuzz_dep_.CoverTab[118348]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:267
		// _ = "end of CoverTab[118348]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:267
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:267
	// _ = "end of CoverTab[118339]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:267
	_go_fuzz_dep_.CoverTab[118340]++
											return 0, BADFD
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:268
	// _ = "end of CoverTab[118340]"
}

func copyFile(base Fs, layer Fs, name string, bfh File) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:271
	_go_fuzz_dep_.CoverTab[118349]++

											exists, err := Exists(layer, filepath.Dir(name))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:274
		_go_fuzz_dep_.CoverTab[118356]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:275
		// _ = "end of CoverTab[118356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:276
		_go_fuzz_dep_.CoverTab[118357]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:276
		// _ = "end of CoverTab[118357]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:276
	// _ = "end of CoverTab[118349]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:276
	_go_fuzz_dep_.CoverTab[118350]++
											if !exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:277
		_go_fuzz_dep_.CoverTab[118358]++
												err = layer.MkdirAll(filepath.Dir(name), 0777)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:279
			_go_fuzz_dep_.CoverTab[118359]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:280
			// _ = "end of CoverTab[118359]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:281
			_go_fuzz_dep_.CoverTab[118360]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:281
			// _ = "end of CoverTab[118360]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:281
		// _ = "end of CoverTab[118358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:282
		_go_fuzz_dep_.CoverTab[118361]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:282
		// _ = "end of CoverTab[118361]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:282
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:282
	// _ = "end of CoverTab[118350]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:282
	_go_fuzz_dep_.CoverTab[118351]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:285
	lfh, err := layer.Create(name)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:286
		_go_fuzz_dep_.CoverTab[118362]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:287
		// _ = "end of CoverTab[118362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:288
		_go_fuzz_dep_.CoverTab[118363]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:288
		// _ = "end of CoverTab[118363]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:288
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:288
	// _ = "end of CoverTab[118351]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:288
	_go_fuzz_dep_.CoverTab[118352]++
											n, err := io.Copy(lfh, bfh)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:290
		_go_fuzz_dep_.CoverTab[118364]++

												layer.Remove(name)
												lfh.Close()
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:294
		// _ = "end of CoverTab[118364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:295
		_go_fuzz_dep_.CoverTab[118365]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:295
		// _ = "end of CoverTab[118365]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:295
	// _ = "end of CoverTab[118352]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:295
	_go_fuzz_dep_.CoverTab[118353]++

											bfi, err := bfh.Stat()
											if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:298
		_go_fuzz_dep_.CoverTab[118366]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:298
		return bfi.Size() != n
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:298
		// _ = "end of CoverTab[118366]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:298
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:298
		_go_fuzz_dep_.CoverTab[118367]++
												layer.Remove(name)
												lfh.Close()
												return syscall.EIO
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:301
		// _ = "end of CoverTab[118367]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:302
		_go_fuzz_dep_.CoverTab[118368]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:302
		// _ = "end of CoverTab[118368]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:302
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:302
	// _ = "end of CoverTab[118353]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:302
	_go_fuzz_dep_.CoverTab[118354]++

											err = lfh.Close()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:305
		_go_fuzz_dep_.CoverTab[118369]++
												layer.Remove(name)
												lfh.Close()
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:308
		// _ = "end of CoverTab[118369]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:309
		_go_fuzz_dep_.CoverTab[118370]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:309
		// _ = "end of CoverTab[118370]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:309
	// _ = "end of CoverTab[118354]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:309
	_go_fuzz_dep_.CoverTab[118355]++
											return layer.Chtimes(name, bfi.ModTime(), bfi.ModTime())
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:310
	// _ = "end of CoverTab[118355]"
}

func copyToLayer(base Fs, layer Fs, name string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:313
	_go_fuzz_dep_.CoverTab[118371]++
											bfh, err := base.Open(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:315
		_go_fuzz_dep_.CoverTab[118373]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:316
		// _ = "end of CoverTab[118373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:317
		_go_fuzz_dep_.CoverTab[118374]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:317
		// _ = "end of CoverTab[118374]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:317
	// _ = "end of CoverTab[118371]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:317
	_go_fuzz_dep_.CoverTab[118372]++
											defer bfh.Close()

											return copyFile(base, layer, name, bfh)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:320
	// _ = "end of CoverTab[118372]"
}

func copyFileToLayer(base Fs, layer Fs, name string, flag int, perm os.FileMode) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:323
	_go_fuzz_dep_.CoverTab[118375]++
											bfh, err := base.OpenFile(name, flag, perm)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:325
		_go_fuzz_dep_.CoverTab[118377]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:326
		// _ = "end of CoverTab[118377]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:327
		_go_fuzz_dep_.CoverTab[118378]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:327
		// _ = "end of CoverTab[118378]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:327
	// _ = "end of CoverTab[118375]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:327
	_go_fuzz_dep_.CoverTab[118376]++
											defer bfh.Close()

											return copyFile(base, layer, name, bfh)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:330
	// _ = "end of CoverTab[118376]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:331
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/afero@v1.9.2/unionFile.go:331
var _ = _go_fuzz_dep_.CoverTab
