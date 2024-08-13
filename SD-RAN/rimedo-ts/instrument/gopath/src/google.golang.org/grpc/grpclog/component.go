//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:19
)

import (
	"fmt"

	"google.golang.org/grpc/internal/grpclog"
)

// componentData records the settings for a component.
type componentData struct {
	name string
}

var cache = map[string]*componentData{}

func (c *componentData) InfoDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:34
	_go_fuzz_dep_.CoverTab[48082]++
												args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
												grpclog.InfoDepth(depth+1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:36
	// _ = "end of CoverTab[48082]"
}

func (c *componentData) WarningDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:39
	_go_fuzz_dep_.CoverTab[48083]++
												args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
												grpclog.WarningDepth(depth+1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:41
	// _ = "end of CoverTab[48083]"
}

func (c *componentData) ErrorDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:44
	_go_fuzz_dep_.CoverTab[48084]++
												args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
												grpclog.ErrorDepth(depth+1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:46
	// _ = "end of CoverTab[48084]"
}

func (c *componentData) FatalDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:49
	_go_fuzz_dep_.CoverTab[48085]++
												args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
												grpclog.FatalDepth(depth+1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:51
	// _ = "end of CoverTab[48085]"
}

func (c *componentData) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:54
	_go_fuzz_dep_.CoverTab[48086]++
												c.InfoDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:55
	// _ = "end of CoverTab[48086]"
}

func (c *componentData) Warning(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:58
	_go_fuzz_dep_.CoverTab[48087]++
												c.WarningDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:59
	// _ = "end of CoverTab[48087]"
}

func (c *componentData) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:62
	_go_fuzz_dep_.CoverTab[48088]++
												c.ErrorDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:63
	// _ = "end of CoverTab[48088]"
}

func (c *componentData) Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:66
	_go_fuzz_dep_.CoverTab[48089]++
												c.FatalDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:67
	// _ = "end of CoverTab[48089]"
}

func (c *componentData) Infof(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:70
	_go_fuzz_dep_.CoverTab[48090]++
												c.InfoDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:71
	// _ = "end of CoverTab[48090]"
}

func (c *componentData) Warningf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:74
	_go_fuzz_dep_.CoverTab[48091]++
												c.WarningDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:75
	// _ = "end of CoverTab[48091]"
}

func (c *componentData) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:78
	_go_fuzz_dep_.CoverTab[48092]++
												c.ErrorDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:79
	// _ = "end of CoverTab[48092]"
}

func (c *componentData) Fatalf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:82
	_go_fuzz_dep_.CoverTab[48093]++
												c.FatalDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:83
	// _ = "end of CoverTab[48093]"
}

func (c *componentData) Infoln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:86
	_go_fuzz_dep_.CoverTab[48094]++
												c.InfoDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:87
	// _ = "end of CoverTab[48094]"
}

func (c *componentData) Warningln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:90
	_go_fuzz_dep_.CoverTab[48095]++
												c.WarningDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:91
	// _ = "end of CoverTab[48095]"
}

func (c *componentData) Errorln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:94
	_go_fuzz_dep_.CoverTab[48096]++
												c.ErrorDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:95
	// _ = "end of CoverTab[48096]"
}

func (c *componentData) Fatalln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:98
	_go_fuzz_dep_.CoverTab[48097]++
												c.FatalDepth(1, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:99
	// _ = "end of CoverTab[48097]"
}

func (c *componentData) V(l int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:102
	_go_fuzz_dep_.CoverTab[48098]++
												return V(l)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:103
	// _ = "end of CoverTab[48098]"
}

// Component creates a new component and returns it for logging. If a component
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:106
// with the name already exists, nothing will be created and it will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:106
// returned. SetLoggerV2 will panic if it is called with a logger created by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:106
// Component.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:110
func Component(componentName string) DepthLoggerV2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:110
	_go_fuzz_dep_.CoverTab[48099]++
												if cData, ok := cache[componentName]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:111
		_go_fuzz_dep_.CoverTab[48101]++
													return cData
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:112
		// _ = "end of CoverTab[48101]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:113
		_go_fuzz_dep_.CoverTab[48102]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:113
		// _ = "end of CoverTab[48102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:113
	// _ = "end of CoverTab[48099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:113
	_go_fuzz_dep_.CoverTab[48100]++
												c := &componentData{componentName}
												cache[componentName] = c
												return c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:116
	// _ = "end of CoverTab[48100]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:117
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/component.go:117
var _ = _go_fuzz_dep_.CoverTab
