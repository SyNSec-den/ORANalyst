//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
package json

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:1
)

import (
	"encoding/json"
)

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for JSON encoding.
type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:10
	_go_fuzz_dep_.CoverTab[122494]++

													return json.MarshalIndent(v, "", "  ")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:12
	// _ = "end of CoverTab[122494]"
}

func (Codec) Decode(b []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:15
	_go_fuzz_dep_.CoverTab[122495]++
													return json.Unmarshal(b, v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:16
	// _ = "end of CoverTab[122495]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/json/codec.go:17
var _ = _go_fuzz_dep_.CoverTab
