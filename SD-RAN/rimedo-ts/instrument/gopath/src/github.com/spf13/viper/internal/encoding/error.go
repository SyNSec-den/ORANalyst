//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
package encoding

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:1
)

type encodingError string

func (e encodingError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:5
	_go_fuzz_dep_.CoverTab[120800]++
												return string(e)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:6
	// _ = "end of CoverTab[120800]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:7
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/error.go:7
var _ = _go_fuzz_dep_.CoverTab
