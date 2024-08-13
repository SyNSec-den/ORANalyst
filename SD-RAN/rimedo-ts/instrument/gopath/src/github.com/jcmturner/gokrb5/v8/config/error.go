//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
package config

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:1
)

import "fmt"

// UnsupportedDirective error.
type UnsupportedDirective struct {
	text string
}

// Error implements the error interface for unsupported directives.
func (e UnsupportedDirective) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:11
	_go_fuzz_dep_.CoverTab[83256]++
												return e.text
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:12
	// _ = "end of CoverTab[83256]"
}

// Invalid config error.
type Invalid struct {
	text string
}

// Error implements the error interface for invalid config error.
func (e Invalid) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:21
	_go_fuzz_dep_.CoverTab[83257]++
												return e.text
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:22
	// _ = "end of CoverTab[83257]"
}

// InvalidErrorf creates a new Invalid error.
func InvalidErrorf(format string, a ...interface{}) Invalid {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:26
	_go_fuzz_dep_.CoverTab[83258]++
												return Invalid{
		text: fmt.Sprintf("invalid krb5 config "+format, a...),
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:29
	// _ = "end of CoverTab[83258]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/config/error.go:30
var _ = _go_fuzz_dep_.CoverTab
