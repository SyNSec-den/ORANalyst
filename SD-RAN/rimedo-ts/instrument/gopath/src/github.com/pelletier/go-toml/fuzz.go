// +build gofuzz

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:3
)

func Fuzz(data []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:5
	_go_fuzz_dep_.CoverTab[122496]++
											tree, err := LoadBytes(data)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:7
		_go_fuzz_dep_.CoverTab[122500]++
												if tree != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:8
			_go_fuzz_dep_.CoverTab[122502]++
													panic("tree must be nil if there is an error")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:9
			// _ = "end of CoverTab[122502]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:10
			_go_fuzz_dep_.CoverTab[122503]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:10
			// _ = "end of CoverTab[122503]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:10
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:10
		// _ = "end of CoverTab[122500]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:10
		_go_fuzz_dep_.CoverTab[122501]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:11
		// _ = "end of CoverTab[122501]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:12
		_go_fuzz_dep_.CoverTab[122504]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:12
		// _ = "end of CoverTab[122504]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:12
	// _ = "end of CoverTab[122496]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:12
	_go_fuzz_dep_.CoverTab[122497]++

											str, err := tree.ToTomlString()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:15
		_go_fuzz_dep_.CoverTab[122505]++
												if str != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:16
			_go_fuzz_dep_.CoverTab[122507]++
													panic(`str must be "" if there is an error`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:17
			// _ = "end of CoverTab[122507]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:18
			_go_fuzz_dep_.CoverTab[122508]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:18
			// _ = "end of CoverTab[122508]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:18
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:18
		// _ = "end of CoverTab[122505]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:18
		_go_fuzz_dep_.CoverTab[122506]++
												panic(err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:19
		// _ = "end of CoverTab[122506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:20
		_go_fuzz_dep_.CoverTab[122509]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:20
		// _ = "end of CoverTab[122509]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:20
	// _ = "end of CoverTab[122497]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:20
	_go_fuzz_dep_.CoverTab[122498]++

											tree, err = Load(str)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:23
		_go_fuzz_dep_.CoverTab[122510]++
												if tree != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:24
			_go_fuzz_dep_.CoverTab[122512]++
													panic("tree must be nil if there is an error")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:25
			// _ = "end of CoverTab[122512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:26
			_go_fuzz_dep_.CoverTab[122513]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:26
			// _ = "end of CoverTab[122513]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:26
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:26
		// _ = "end of CoverTab[122510]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:26
		_go_fuzz_dep_.CoverTab[122511]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:27
		// _ = "end of CoverTab[122511]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:28
		_go_fuzz_dep_.CoverTab[122514]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:28
		// _ = "end of CoverTab[122514]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:28
	// _ = "end of CoverTab[122498]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:28
	_go_fuzz_dep_.CoverTab[122499]++

											return 1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:30
	// _ = "end of CoverTab[122499]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/fuzz.go:31
var _ = _go_fuzz_dep_.CoverTab
