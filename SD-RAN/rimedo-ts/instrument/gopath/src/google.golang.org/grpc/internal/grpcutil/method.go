//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
package grpcutil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:19
)

import (
	"errors"
	"strings"
)

// ParseMethod splits service and method from the input. It expects format
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:26
// "/service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:28
func ParseMethod(methodName string) (service, method string, _ error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:28
	_go_fuzz_dep_.CoverTab[67630]++
													if !strings.HasPrefix(methodName, "/") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:29
		_go_fuzz_dep_.CoverTab[67633]++
														return "", "", errors.New("invalid method name: should start with /")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:30
		// _ = "end of CoverTab[67633]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:31
		_go_fuzz_dep_.CoverTab[67634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:31
		// _ = "end of CoverTab[67634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:31
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:31
	// _ = "end of CoverTab[67630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:31
	_go_fuzz_dep_.CoverTab[67631]++
													methodName = methodName[1:]

													pos := strings.LastIndex(methodName, "/")
													if pos < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:35
		_go_fuzz_dep_.CoverTab[67635]++
														return "", "", errors.New("invalid method name: suffix /method is missing")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:36
		// _ = "end of CoverTab[67635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:37
		_go_fuzz_dep_.CoverTab[67636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:37
		// _ = "end of CoverTab[67636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:37
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:37
	// _ = "end of CoverTab[67631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:37
	_go_fuzz_dep_.CoverTab[67632]++
													return methodName[:pos], methodName[pos+1:], nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:38
	// _ = "end of CoverTab[67632]"
}

// baseContentType is the base content-type for gRPC.  This is a valid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:41
// content-type on it's own, but can also include a content-subtype such as
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:41
// "proto" as a suffix after "+" or ";".  See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:41
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:41
// for more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:46
const baseContentType = "application/grpc"

// ContentSubtype returns the content-subtype for the given content-type.  The
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// given content-type must be a valid content-type that starts with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// "application/grpc". A content-subtype will follow "application/grpc" after a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// "+" or ";". See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// If contentType is not a valid content-type for gRPC, the boolean
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// will be false, otherwise true. If content-type == "application/grpc",
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// "application/grpc+", or "application/grpc;", the boolean will be true,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// but no content-subtype will be returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:48
// contentType is assumed to be lowercase already.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:61
func ContentSubtype(contentType string) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:61
	_go_fuzz_dep_.CoverTab[67637]++
													if contentType == baseContentType {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:62
		_go_fuzz_dep_.CoverTab[67640]++
														return "", true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:63
		// _ = "end of CoverTab[67640]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:64
		_go_fuzz_dep_.CoverTab[67641]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:64
		// _ = "end of CoverTab[67641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:64
	// _ = "end of CoverTab[67637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:64
	_go_fuzz_dep_.CoverTab[67638]++
													if !strings.HasPrefix(contentType, baseContentType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:65
		_go_fuzz_dep_.CoverTab[67642]++
														return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:66
		// _ = "end of CoverTab[67642]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:67
		_go_fuzz_dep_.CoverTab[67643]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:67
		// _ = "end of CoverTab[67643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:67
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:67
	// _ = "end of CoverTab[67638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:67
	_go_fuzz_dep_.CoverTab[67639]++

													switch contentType[len(baseContentType)] {
	case '+', ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:70
		_go_fuzz_dep_.CoverTab[67644]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:74
		return contentType[len(baseContentType)+1:], true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:74
		// _ = "end of CoverTab[67644]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:75
		_go_fuzz_dep_.CoverTab[67645]++
														return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:76
		// _ = "end of CoverTab[67645]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:77
	// _ = "end of CoverTab[67639]"
}

// ContentType builds full content type with the given sub-type.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:80
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:80
// contentSubtype is assumed to be lowercase
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:83
func ContentType(contentSubtype string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:83
	_go_fuzz_dep_.CoverTab[67646]++
													if contentSubtype == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:84
		_go_fuzz_dep_.CoverTab[67648]++
														return baseContentType
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:85
		// _ = "end of CoverTab[67648]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:86
		_go_fuzz_dep_.CoverTab[67649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:86
		// _ = "end of CoverTab[67649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:86
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:86
	// _ = "end of CoverTab[67646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:86
	_go_fuzz_dep_.CoverTab[67647]++
													return baseContentType + "+" + contentSubtype
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:87
	// _ = "end of CoverTab[67647]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/method.go:88
var _ = _go_fuzz_dep_.CoverTab
