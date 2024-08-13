//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
package hcl

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:1
)

import (
	"bytes"
	"encoding/json"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
)

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for HCL encoding.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:11
// TODO: add printer config to the codec?
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:13
type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:15
	_go_fuzz_dep_.CoverTab[122483]++
													b, err := json.Marshal(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:17
		_go_fuzz_dep_.CoverTab[122487]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:18
		// _ = "end of CoverTab[122487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:19
		_go_fuzz_dep_.CoverTab[122488]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:19
		// _ = "end of CoverTab[122488]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:19
	// _ = "end of CoverTab[122483]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:19
	_go_fuzz_dep_.CoverTab[122484]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:23
	ast, err := hcl.Parse(string(b))
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:24
		_go_fuzz_dep_.CoverTab[122489]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:25
		// _ = "end of CoverTab[122489]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:26
		_go_fuzz_dep_.CoverTab[122490]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:26
		// _ = "end of CoverTab[122490]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:26
	// _ = "end of CoverTab[122484]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:26
	_go_fuzz_dep_.CoverTab[122485]++

													var buf bytes.Buffer

													err = printer.Fprint(&buf, ast.Node)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:31
		_go_fuzz_dep_.CoverTab[122491]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:32
		// _ = "end of CoverTab[122491]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:33
		_go_fuzz_dep_.CoverTab[122492]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:33
		// _ = "end of CoverTab[122492]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:33
	// _ = "end of CoverTab[122485]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:33
	_go_fuzz_dep_.CoverTab[122486]++

													return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:35
	// _ = "end of CoverTab[122486]"
}

func (Codec) Decode(b []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:38
	_go_fuzz_dep_.CoverTab[122493]++
													return hcl.Unmarshal(b, v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:39
	// _ = "end of CoverTab[122493]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/hcl/codec.go:40
var _ = _go_fuzz_dep_.CoverTab
