// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
package grpc_auth

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:4
)

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	headerAuthorize = "authorization"
)

// AuthFromMD is a helper function for extracting the :authorization header from the gRPC metadata of the request.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:19
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:19
// It expects the `:authorization` header to be of a certain scheme (e.g. `basic`, `bearer`), in a
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:19
// case-insensitive format (see rfc2617, sec 1.2). If no such authorization is found, or the token
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:19
// is of wrong scheme, an error with gRPC status `Unauthenticated` is returned.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:24
func AuthFromMD(ctx context.Context, expectedScheme string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:24
	_go_fuzz_dep_.CoverTab[183695]++
														val := metautils.ExtractIncoming(ctx).Get(headerAuthorize)
														if val == "" {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:26
		_go_fuzz_dep_.CoverTab[183699]++
															return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:27
		// _ = "end of CoverTab[183699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:28
		_go_fuzz_dep_.CoverTab[183700]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:28
		// _ = "end of CoverTab[183700]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:28
	// _ = "end of CoverTab[183695]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:28
	_go_fuzz_dep_.CoverTab[183696]++
														splits := strings.SplitN(val, " ", 2)
														if len(splits) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:30
		_go_fuzz_dep_.CoverTab[183701]++
															return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:31
		// _ = "end of CoverTab[183701]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:32
		_go_fuzz_dep_.CoverTab[183702]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:32
		// _ = "end of CoverTab[183702]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:32
	// _ = "end of CoverTab[183696]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:32
	_go_fuzz_dep_.CoverTab[183697]++
														if !strings.EqualFold(splits[0], expectedScheme) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:33
		_go_fuzz_dep_.CoverTab[183703]++
															return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:34
		// _ = "end of CoverTab[183703]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:35
		_go_fuzz_dep_.CoverTab[183704]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:35
		// _ = "end of CoverTab[183704]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:35
	// _ = "end of CoverTab[183697]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:35
	_go_fuzz_dep_.CoverTab[183698]++
														return splits[1], nil
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:36
	// _ = "end of CoverTab[183698]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/auth/metadata.go:37
var _ = _go_fuzz_dep_.CoverTab
