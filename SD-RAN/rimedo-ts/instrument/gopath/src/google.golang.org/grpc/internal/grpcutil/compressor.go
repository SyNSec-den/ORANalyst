//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
package grpcutil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:19
)

import (
	"strings"

	"google.golang.org/grpc/internal/envconfig"
)

// RegisteredCompressorNames holds names of the registered compressors.
var RegisteredCompressorNames []string

// IsCompressorNameRegistered returns true when name is available in registry.
func IsCompressorNameRegistered(name string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:31
	_go_fuzz_dep_.CoverTab[67596]++
													for _, compressor := range RegisteredCompressorNames {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:32
		_go_fuzz_dep_.CoverTab[67598]++
														if compressor == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:33
			_go_fuzz_dep_.CoverTab[67599]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:34
			// _ = "end of CoverTab[67599]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:35
			_go_fuzz_dep_.CoverTab[67600]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:35
			// _ = "end of CoverTab[67600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:35
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:35
		// _ = "end of CoverTab[67598]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:36
	// _ = "end of CoverTab[67596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:36
	_go_fuzz_dep_.CoverTab[67597]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:37
	// _ = "end of CoverTab[67597]"
}

// RegisteredCompressors returns a string of registered compressor names
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:40
// separated by comma.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:42
func RegisteredCompressors() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:42
	_go_fuzz_dep_.CoverTab[67601]++
													if !envconfig.AdvertiseCompressors {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:43
		_go_fuzz_dep_.CoverTab[67603]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:44
		// _ = "end of CoverTab[67603]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:45
		_go_fuzz_dep_.CoverTab[67604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:45
		// _ = "end of CoverTab[67604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:45
	// _ = "end of CoverTab[67601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:45
	_go_fuzz_dep_.CoverTab[67602]++
													return strings.Join(RegisteredCompressorNames, ",")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:46
	// _ = "end of CoverTab[67602]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:47
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/compressor.go:47
var _ = _go_fuzz_dep_.CoverTab
