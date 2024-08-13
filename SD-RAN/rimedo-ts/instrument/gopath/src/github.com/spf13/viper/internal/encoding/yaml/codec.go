//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
package yaml

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:1
)

import "gopkg.in/yaml.v2"

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for YAML encoding.
type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:8
	_go_fuzz_dep_.CoverTab[128144]++
													return yaml.Marshal(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:9
	// _ = "end of CoverTab[128144]"
}

func (Codec) Decode(b []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:12
	_go_fuzz_dep_.CoverTab[128145]++
													return yaml.Unmarshal(b, v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:13
	// _ = "end of CoverTab[128145]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:14
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/yaml/codec.go:14
var _ = _go_fuzz_dep_.CoverTab
