//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:1
// Package krberror provides error type and functions for gokrb5.
package krberror

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:2
)

import (
	"fmt"
	"strings"
)

// Error type descriptions.
const (
	separator	= " < "
	EncodingError	= "Encoding_Error"
	NetworkingError	= "Networking_Error"
	DecryptingError	= "Decrypting_Error"
	EncryptingError	= "Encrypting_Error"
	ChksumError	= "Checksum_Error"
	KRBMsgError	= "KRBMessage_Handling_Error"
	ConfigError	= "Configuration_Error"
	KDCError	= "KDC_Error"
)

// Krberror is an error type for gokrb5
type Krberror struct {
	RootCause	string
	EText		[]string
}

// Error function to implement the error interface.
func (e Krberror) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:29
	_go_fuzz_dep_.CoverTab[86737]++
												return fmt.Sprintf("[Root cause: %s] ", e.RootCause) + strings.Join(e.EText, separator)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:30
	// _ = "end of CoverTab[86737]"
}

// Add another error statement to the error.
func (e *Krberror) Add(et string, s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:34
	_go_fuzz_dep_.CoverTab[86738]++
												e.EText = append([]string{fmt.Sprintf("%s: %s", et, s)}, e.EText...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:35
	// _ = "end of CoverTab[86738]"
}

// New creates a new instance of Krberror.
func New(et, s string) Krberror {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:39
	_go_fuzz_dep_.CoverTab[86739]++
												return Krberror{
		RootCause:	et,
		EText:		[]string{s},
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:43
	// _ = "end of CoverTab[86739]"
}

// Errorf appends to or creates a new Krberror.
func Errorf(err error, et, format string, a ...interface{}) Krberror {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:47
	_go_fuzz_dep_.CoverTab[86740]++
												if e, ok := err.(Krberror); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:48
		_go_fuzz_dep_.CoverTab[86742]++
													e.Add(et, fmt.Sprintf(format, a...))
													return e
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:50
		// _ = "end of CoverTab[86742]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:51
		_go_fuzz_dep_.CoverTab[86743]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:51
		// _ = "end of CoverTab[86743]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:51
	// _ = "end of CoverTab[86740]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:51
	_go_fuzz_dep_.CoverTab[86741]++
												return NewErrorf(et, format+": %s", append(a, err)...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:52
	// _ = "end of CoverTab[86741]"
}

// NewErrorf creates a new Krberror from a formatted string.
func NewErrorf(et, format string, a ...interface{}) Krberror {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:56
	_go_fuzz_dep_.CoverTab[86744]++
												var s string
												if len(a) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:58
		_go_fuzz_dep_.CoverTab[86746]++
													s = fmt.Sprintf("%s: %s", et, fmt.Sprintf(format, a...))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:59
		// _ = "end of CoverTab[86746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:60
		_go_fuzz_dep_.CoverTab[86747]++
													s = fmt.Sprintf("%s: %s", et, format)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:61
		// _ = "end of CoverTab[86747]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:62
	// _ = "end of CoverTab[86744]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:62
	_go_fuzz_dep_.CoverTab[86745]++
												return Krberror{
		RootCause:	et,
		EText:		[]string{s},
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:66
	// _ = "end of CoverTab[86745]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:67
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/krberror/error.go:67
var _ = _go_fuzz_dep_.CoverTab
