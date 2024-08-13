//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
package grpcsync

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:19
)

import (
	"sync"
)

// OnceFunc returns a function wrapping f which ensures f is only executed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:25
// once even if the returned function is executed multiple times.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:27
func OnceFunc(f func()) func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:27
	_go_fuzz_dep_.CoverTab[68900]++
													var once sync.Once
													return func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:29
		_go_fuzz_dep_.CoverTab[68901]++
														once.Do(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:30
		// _ = "end of CoverTab[68901]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:31
	// _ = "end of CoverTab[68900]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/oncefunc.go:32
var _ = _go_fuzz_dep_.CoverTab
