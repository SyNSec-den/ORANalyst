//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
package mapstructure

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:1
)

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Error implements the error interface and can represents multiple
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:10
// errors that occur in the course of a single decode.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:12
type Error struct {
	Errors []string
}

func (e *Error) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:16
	_go_fuzz_dep_.CoverTab[116269]++
											points := make([]string, len(e.Errors))
											for i, err := range e.Errors {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:18
		_go_fuzz_dep_.CoverTab[116271]++
												points[i] = fmt.Sprintf("* %s", err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:19
		// _ = "end of CoverTab[116271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:20
	// _ = "end of CoverTab[116269]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:20
	_go_fuzz_dep_.CoverTab[116270]++

											sort.Strings(points)
											return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(e.Errors), strings.Join(points, "\n"))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:25
	// _ = "end of CoverTab[116270]"
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:28
// return value more useful with the errwrap and go-multierror libraries.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:30
func (e *Error) WrappedErrors() []error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:30
	_go_fuzz_dep_.CoverTab[116272]++
											if e == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:31
		_go_fuzz_dep_.CoverTab[116275]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:32
		// _ = "end of CoverTab[116275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:33
		_go_fuzz_dep_.CoverTab[116276]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:33
		// _ = "end of CoverTab[116276]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:33
	// _ = "end of CoverTab[116272]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:33
	_go_fuzz_dep_.CoverTab[116273]++

											result := make([]error, len(e.Errors))
											for i, e := range e.Errors {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:36
		_go_fuzz_dep_.CoverTab[116277]++
												result[i] = errors.New(e)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:37
		// _ = "end of CoverTab[116277]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:38
	// _ = "end of CoverTab[116273]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:38
	_go_fuzz_dep_.CoverTab[116274]++

											return result
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:40
	// _ = "end of CoverTab[116274]"
}

func appendErrors(errors []string, err error) []string {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:43
	_go_fuzz_dep_.CoverTab[116278]++
											switch e := err.(type) {
	case *Error:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:45
		_go_fuzz_dep_.CoverTab[116279]++
												return append(errors, e.Errors...)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:46
		// _ = "end of CoverTab[116279]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:47
		_go_fuzz_dep_.CoverTab[116280]++
												return append(errors, e.Error())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:48
		// _ = "end of CoverTab[116280]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:49
	// _ = "end of CoverTab[116278]"
}

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:50
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/error.go:50
var _ = _go_fuzz_dep_.CoverTab
