// +build gofuzz

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
package snappy

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:3
)

func Fuzz(data []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:5
	_go_fuzz_dep_.CoverTab[82221]++
															decode, err := Decode(data)
															if decode == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:7
		_go_fuzz_dep_.CoverTab[82224]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:7
		return err == nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:7
		// _ = "end of CoverTab[82224]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:7
	}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:7
		_go_fuzz_dep_.CoverTab[82225]++
																panic("nil error with nil result")
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:8
		// _ = "end of CoverTab[82225]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:9
		_go_fuzz_dep_.CoverTab[82226]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:9
		// _ = "end of CoverTab[82226]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:9
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:9
	// _ = "end of CoverTab[82221]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:9
	_go_fuzz_dep_.CoverTab[82222]++

															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:11
		_go_fuzz_dep_.CoverTab[82227]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:12
		// _ = "end of CoverTab[82227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:13
		_go_fuzz_dep_.CoverTab[82228]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:13
		// _ = "end of CoverTab[82228]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:13
	// _ = "end of CoverTab[82222]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:13
	_go_fuzz_dep_.CoverTab[82223]++

															return 1
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:15
	// _ = "end of CoverTab[82223]"
}

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/fuzz.go:16
var _ = _go_fuzz_dep_.CoverTab
