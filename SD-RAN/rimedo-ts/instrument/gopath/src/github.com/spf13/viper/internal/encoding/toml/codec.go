//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:1
)

import (
	"github.com/pelletier/go-toml"
)

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for TOML encoding.
type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:10
	_go_fuzz_dep_.CoverTab[124422]++
													if m, ok := v.(map[string]interface{}); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:11
		_go_fuzz_dep_.CoverTab[124424]++
														t, err := toml.TreeFromMap(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:13
			_go_fuzz_dep_.CoverTab[124427]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:14
			// _ = "end of CoverTab[124427]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:15
			_go_fuzz_dep_.CoverTab[124428]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:15
			// _ = "end of CoverTab[124428]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:15
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:15
		// _ = "end of CoverTab[124424]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:15
		_go_fuzz_dep_.CoverTab[124425]++

														s, err := t.ToTomlString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:18
			_go_fuzz_dep_.CoverTab[124429]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:19
			// _ = "end of CoverTab[124429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:20
			_go_fuzz_dep_.CoverTab[124430]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:20
			// _ = "end of CoverTab[124430]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:20
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:20
		// _ = "end of CoverTab[124425]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:20
		_go_fuzz_dep_.CoverTab[124426]++

														return []byte(s), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:22
		// _ = "end of CoverTab[124426]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:23
		_go_fuzz_dep_.CoverTab[124431]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:23
		// _ = "end of CoverTab[124431]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:23
	// _ = "end of CoverTab[124422]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:23
	_go_fuzz_dep_.CoverTab[124423]++

													return toml.Marshal(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:25
	// _ = "end of CoverTab[124423]"
}

func (Codec) Decode(b []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:28
	_go_fuzz_dep_.CoverTab[124432]++
													tree, err := toml.LoadBytes(b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:30
		_go_fuzz_dep_.CoverTab[124435]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:31
		// _ = "end of CoverTab[124435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:32
		_go_fuzz_dep_.CoverTab[124436]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:32
		// _ = "end of CoverTab[124436]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:32
	// _ = "end of CoverTab[124432]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:32
	_go_fuzz_dep_.CoverTab[124433]++

													if m, ok := v.(*map[string]interface{}); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:34
		_go_fuzz_dep_.CoverTab[124437]++
														vmap := *m
														tmap := tree.ToMap()
														for k, v := range tmap {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:37
			_go_fuzz_dep_.CoverTab[124439]++
															vmap[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:38
			// _ = "end of CoverTab[124439]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:39
		// _ = "end of CoverTab[124437]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:39
		_go_fuzz_dep_.CoverTab[124438]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:41
		// _ = "end of CoverTab[124438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:42
		_go_fuzz_dep_.CoverTab[124440]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:42
		// _ = "end of CoverTab[124440]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:42
	// _ = "end of CoverTab[124433]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:42
	_go_fuzz_dep_.CoverTab[124434]++

													return tree.Unmarshal(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:44
	// _ = "end of CoverTab[124434]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:45
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/toml/codec.go:45
var _ = _go_fuzz_dep_.CoverTab
