//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:1
)

import (
	"errors"
	"fmt"
	"os"
	rdebug "runtime/debug"
)

var (
	// ErrInvalidSourceShortBuffer is returned by UncompressBlock or CompressBLock when a compressed
	// block is corrupted or the destination buffer is not large enough for the uncompressed data.
	ErrInvalidSourceShortBuffer	= errors.New("lz4: invalid source or destination buffer too short")
	// ErrInvalid is returned when reading an invalid LZ4 archive.
	ErrInvalid	= errors.New("lz4: bad magic number")
	// ErrBlockDependency is returned when attempting to decompress an archive created with block dependency.
	ErrBlockDependency	= errors.New("lz4: block dependency not supported")
	// ErrUnsupportedSeek is returned when attempting to Seek any way but forward from the current position.
	ErrUnsupportedSeek	= errors.New("lz4: can only seek forward from io.SeekCurrent")
)

func recoverBlock(e *error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:22
	_go_fuzz_dep_.CoverTab[95471]++
												if r := recover(); r != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:23
		_go_fuzz_dep_.CoverTab[95472]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:23
		return *e == nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:23
		// _ = "end of CoverTab[95472]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:23
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:23
		_go_fuzz_dep_.CoverTab[95473]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:24
			_go_fuzz_dep_.CoverTab[95475]++
														fmt.Fprintln(os.Stderr, r)
														rdebug.PrintStack()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:26
			// _ = "end of CoverTab[95475]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:27
			_go_fuzz_dep_.CoverTab[95476]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:27
			// _ = "end of CoverTab[95476]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:27
		// _ = "end of CoverTab[95473]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:27
		_go_fuzz_dep_.CoverTab[95474]++
													*e = ErrInvalidSourceShortBuffer
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:28
		// _ = "end of CoverTab[95474]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:29
		_go_fuzz_dep_.CoverTab[95477]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:29
		// _ = "end of CoverTab[95477]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:29
	// _ = "end of CoverTab[95471]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/errors.go:30
var _ = _go_fuzz_dep_.CoverTab
