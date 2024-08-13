//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:19
// Package grpcrand implements math/rand functions in a concurrent-safe way
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:19
// with a global random source, independent of math/rand's global source.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
package grpcrand

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:21
)

import (
	"math/rand"
	"sync"
	"time"
)

var (
	r	= rand.New(rand.NewSource(time.Now().UnixNano()))
	mu	sync.Mutex
)

// Int implements rand.Int on the grpcrand global source.
func Int() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:35
	_go_fuzz_dep_.CoverTab[67503]++
													mu.Lock()
													defer mu.Unlock()
													return r.Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:38
	// _ = "end of CoverTab[67503]"
}

// Int63n implements rand.Int63n on the grpcrand global source.
func Int63n(n int64) int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:42
	_go_fuzz_dep_.CoverTab[67504]++
													mu.Lock()
													defer mu.Unlock()
													return r.Int63n(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:45
	// _ = "end of CoverTab[67504]"
}

// Intn implements rand.Intn on the grpcrand global source.
func Intn(n int) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:49
	_go_fuzz_dep_.CoverTab[67505]++
													mu.Lock()
													defer mu.Unlock()
													return r.Intn(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:52
	// _ = "end of CoverTab[67505]"
}

// Int31n implements rand.Int31n on the grpcrand global source.
func Int31n(n int32) int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:56
	_go_fuzz_dep_.CoverTab[67506]++
													mu.Lock()
													defer mu.Unlock()
													return r.Int31n(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:59
	// _ = "end of CoverTab[67506]"
}

// Float64 implements rand.Float64 on the grpcrand global source.
func Float64() float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:63
	_go_fuzz_dep_.CoverTab[67507]++
													mu.Lock()
													defer mu.Unlock()
													return r.Float64()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:66
	// _ = "end of CoverTab[67507]"
}

// Uint64 implements rand.Uint64 on the grpcrand global source.
func Uint64() uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:70
	_go_fuzz_dep_.CoverTab[67508]++
													mu.Lock()
													defer mu.Unlock()
													return r.Uint64()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:73
	// _ = "end of CoverTab[67508]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcrand/grpcrand.go:74
var _ = _go_fuzz_dep_.CoverTab
