//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:19
// Package envconfig contains grpc settings configured by environment variables.
package envconfig

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:20
)

import (
	"os"
	"strconv"
	"strings"
)

var (
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore	= boolFromEnv("GRPC_GO_IGNORE_TXT_ERRORS", true)
	// AdvertiseCompressors is set if registered compressor should be advertised
	// ("GRPC_GO_ADVERTISE_COMPRESSORS" is not "false").
	AdvertiseCompressors	= boolFromEnv("GRPC_GO_ADVERTISE_COMPRESSORS", true)
	// RingHashCap indicates the maximum ring size which defaults to 4096
	// entries but may be overridden by setting the environment variable
	// "GRPC_RING_HASH_CAP".  This does not override the default bounds
	// checking which NACKs configs specifying ring sizes > 8*1024*1024 (~8M).
	RingHashCap	= uint64FromEnv("GRPC_RING_HASH_CAP", 4096, 1, 8*1024*1024)
)

func boolFromEnv(envVar string, def bool) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:41
	_go_fuzz_dep_.CoverTab[67582]++
													if def {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:42
		_go_fuzz_dep_.CoverTab[67584]++

														return !strings.EqualFold(os.Getenv(envVar), "false")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:44
		// _ = "end of CoverTab[67584]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:45
		_go_fuzz_dep_.CoverTab[67585]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:45
		// _ = "end of CoverTab[67585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:45
	// _ = "end of CoverTab[67582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:45
	_go_fuzz_dep_.CoverTab[67583]++

													return strings.EqualFold(os.Getenv(envVar), "true")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:47
	// _ = "end of CoverTab[67583]"
}

func uint64FromEnv(envVar string, def, min, max uint64) uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:50
	_go_fuzz_dep_.CoverTab[67586]++
													v, err := strconv.ParseUint(os.Getenv(envVar), 10, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:52
		_go_fuzz_dep_.CoverTab[67590]++
														return def
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:53
		// _ = "end of CoverTab[67590]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:54
		_go_fuzz_dep_.CoverTab[67591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:54
		// _ = "end of CoverTab[67591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:54
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:54
	// _ = "end of CoverTab[67586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:54
	_go_fuzz_dep_.CoverTab[67587]++
													if v < min {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:55
		_go_fuzz_dep_.CoverTab[67592]++
														return min
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:56
		// _ = "end of CoverTab[67592]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:57
		_go_fuzz_dep_.CoverTab[67593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:57
		// _ = "end of CoverTab[67593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:57
	// _ = "end of CoverTab[67587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:57
	_go_fuzz_dep_.CoverTab[67588]++
													if v > max {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:58
		_go_fuzz_dep_.CoverTab[67594]++
														return max
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:59
		// _ = "end of CoverTab[67594]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:60
		_go_fuzz_dep_.CoverTab[67595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:60
		// _ = "end of CoverTab[67595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:60
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:60
	// _ = "end of CoverTab[67588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:60
	_go_fuzz_dep_.CoverTab[67589]++
													return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:61
	// _ = "end of CoverTab[67589]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/envconfig.go:62
var _ = _go_fuzz_dep_.CoverTab
