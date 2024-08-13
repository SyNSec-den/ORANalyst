//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
package grpcutil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:19
)

import "regexp"

// FullMatchWithRegex returns whether the full text matches the regex provided.
func FullMatchWithRegex(re *regexp.Regexp, text string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:24
	_go_fuzz_dep_.CoverTab[67650]++
												if len(text) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:25
		_go_fuzz_dep_.CoverTab[67652]++
													return re.MatchString(text)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:26
		// _ = "end of CoverTab[67652]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:27
		_go_fuzz_dep_.CoverTab[67653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:27
		// _ = "end of CoverTab[67653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:27
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:27
	// _ = "end of CoverTab[67650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:27
	_go_fuzz_dep_.CoverTab[67651]++
												re.Longest()
												rem := re.FindString(text)
												return len(rem) == len(text)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:30
	// _ = "end of CoverTab[67651]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/regex.go:31
var _ = _go_fuzz_dep_.CoverTab
