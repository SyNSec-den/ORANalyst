//+build go1.10

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:3
)

import (
	"fmt"
	"strings"
)

func (h Header) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:10
	_go_fuzz_dep_.CoverTab[95489]++
												var s strings.Builder

												s.WriteString(fmt.Sprintf("%T{", h))
												if h.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:14
		_go_fuzz_dep_.CoverTab[95494]++
													s.WriteString("BlockChecksum: true ")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:15
		// _ = "end of CoverTab[95494]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:16
		_go_fuzz_dep_.CoverTab[95495]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:16
		// _ = "end of CoverTab[95495]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:16
	// _ = "end of CoverTab[95489]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:16
	_go_fuzz_dep_.CoverTab[95490]++
												if h.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:17
		_go_fuzz_dep_.CoverTab[95496]++
													s.WriteString("NoChecksum: true ")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:18
		// _ = "end of CoverTab[95496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:19
		_go_fuzz_dep_.CoverTab[95497]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:19
		// _ = "end of CoverTab[95497]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:19
	// _ = "end of CoverTab[95490]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:19
	_go_fuzz_dep_.CoverTab[95491]++
												if bs := h.BlockMaxSize; bs != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:20
		_go_fuzz_dep_.CoverTab[95498]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:20
		return bs != 4<<20
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:20
		// _ = "end of CoverTab[95498]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:20
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:20
		_go_fuzz_dep_.CoverTab[95499]++
													s.WriteString(fmt.Sprintf("BlockMaxSize: %d ", bs))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:21
		// _ = "end of CoverTab[95499]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:22
		_go_fuzz_dep_.CoverTab[95500]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:22
		// _ = "end of CoverTab[95500]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:22
	// _ = "end of CoverTab[95491]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:22
	_go_fuzz_dep_.CoverTab[95492]++
												if l := h.CompressionLevel; l != 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:23
		_go_fuzz_dep_.CoverTab[95501]++
													s.WriteString(fmt.Sprintf("CompressionLevel: %d ", l))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:24
		// _ = "end of CoverTab[95501]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:25
		_go_fuzz_dep_.CoverTab[95502]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:25
		// _ = "end of CoverTab[95502]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:25
	// _ = "end of CoverTab[95492]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:25
	_go_fuzz_dep_.CoverTab[95493]++
												s.WriteByte('}')

												return s.String()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:28
	// _ = "end of CoverTab[95493]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4_go1.10.go:29
var _ = _go_fuzz_dep_.CoverTab
