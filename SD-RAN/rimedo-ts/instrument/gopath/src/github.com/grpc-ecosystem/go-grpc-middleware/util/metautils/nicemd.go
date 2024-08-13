// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
package metautils

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:4
)

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// NiceMD is a convenience wrapper definiting extra functions on the metadata.
type NiceMD metadata.MD

// ExtractIncoming extracts an inbound metadata from the server-side context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:16
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:16
// This function always returns a NiceMD wrapper of the metadata.MD, in case the context doesn't have metadata it returns
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:16
// a new empty NiceMD.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:20
func ExtractIncoming(ctx context.Context) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:20
	_go_fuzz_dep_.CoverTab[183650]++
															md, ok := metadata.FromIncomingContext(ctx)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:22
		_go_fuzz_dep_.CoverTab[183652]++
																return NiceMD(metadata.Pairs())
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:23
		// _ = "end of CoverTab[183652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:24
		_go_fuzz_dep_.CoverTab[183653]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:24
		// _ = "end of CoverTab[183653]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:24
	// _ = "end of CoverTab[183650]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:24
	_go_fuzz_dep_.CoverTab[183651]++
															return NiceMD(md)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:25
	// _ = "end of CoverTab[183651]"
}

// ExtractOutgoing extracts an outbound metadata from the client-side context.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:28
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:28
// This function always returns a NiceMD wrapper of the metadata.MD, in case the context doesn't have metadata it returns
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:28
// a new empty NiceMD.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:32
func ExtractOutgoing(ctx context.Context) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:32
	_go_fuzz_dep_.CoverTab[183654]++
															md, ok := metadata.FromOutgoingContext(ctx)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:34
		_go_fuzz_dep_.CoverTab[183656]++
																return NiceMD(metadata.Pairs())
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:35
		// _ = "end of CoverTab[183656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:36
		_go_fuzz_dep_.CoverTab[183657]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:36
		// _ = "end of CoverTab[183657]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:36
	// _ = "end of CoverTab[183654]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:36
	_go_fuzz_dep_.CoverTab[183655]++
															return NiceMD(md)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:37
	// _ = "end of CoverTab[183655]"
}

// Clone performs a *deep* copy of the metadata.MD.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:40
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:40
// You can specify the lower-case copiedKeys to only copy certain whitelisted keys. If no keys are explicitly whitelisted
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:40
// all keys get copied.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:44
func (m NiceMD) Clone(copiedKeys ...string) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:44
	_go_fuzz_dep_.CoverTab[183658]++
															newMd := NiceMD(metadata.Pairs())
															for k, vv := range m {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:46
		_go_fuzz_dep_.CoverTab[183660]++
																found := false
																if len(copiedKeys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:48
			_go_fuzz_dep_.CoverTab[183663]++
																	found = true
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:49
			// _ = "end of CoverTab[183663]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:50
			_go_fuzz_dep_.CoverTab[183664]++
																	for _, allowedKey := range copiedKeys {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:51
				_go_fuzz_dep_.CoverTab[183665]++
																		if strings.EqualFold(allowedKey, k) {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:52
					_go_fuzz_dep_.CoverTab[183666]++
																			found = true
																			break
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:54
					// _ = "end of CoverTab[183666]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:55
					_go_fuzz_dep_.CoverTab[183667]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:55
					// _ = "end of CoverTab[183667]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:55
				}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:55
				// _ = "end of CoverTab[183665]"
			}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:56
			// _ = "end of CoverTab[183664]"
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:57
		// _ = "end of CoverTab[183660]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:57
		_go_fuzz_dep_.CoverTab[183661]++
																if !found {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:58
			_go_fuzz_dep_.CoverTab[183668]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:59
			// _ = "end of CoverTab[183668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:60
			_go_fuzz_dep_.CoverTab[183669]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:60
			// _ = "end of CoverTab[183669]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:60
		// _ = "end of CoverTab[183661]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:60
		_go_fuzz_dep_.CoverTab[183662]++
																newMd[k] = make([]string, len(vv))
																copy(newMd[k], vv)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:62
		// _ = "end of CoverTab[183662]"
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:63
	// _ = "end of CoverTab[183658]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:63
	_go_fuzz_dep_.CoverTab[183659]++
															return NiceMD(newMd)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:64
	// _ = "end of CoverTab[183659]"
}

// ToOutgoing sets the given NiceMD as a client-side context for dispatching.
func (m NiceMD) ToOutgoing(ctx context.Context) context.Context {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:68
	_go_fuzz_dep_.CoverTab[183670]++
															return metadata.NewOutgoingContext(ctx, metadata.MD(m))
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:69
	// _ = "end of CoverTab[183670]"
}

// ToIncoming sets the given NiceMD as a server-side context for dispatching.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:72
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:72
// This is mostly useful in ServerInterceptors..
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:75
func (m NiceMD) ToIncoming(ctx context.Context) context.Context {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:75
	_go_fuzz_dep_.CoverTab[183671]++
															return metadata.NewIncomingContext(ctx, metadata.MD(m))
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:76
	// _ = "end of CoverTab[183671]"
}

// Get retrieves a single value from the metadata.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:79
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:79
// It works analogously to http.Header.Get, returning the first value if there are many set. If the value is not set,
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:79
// an empty string is returned.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:79
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:79
// The function is binary-key safe.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:85
func (m NiceMD) Get(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:85
	_go_fuzz_dep_.CoverTab[183672]++
															k := strings.ToLower(key)
															vv, ok := m[k]
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:88
		_go_fuzz_dep_.CoverTab[183674]++
																return ""
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:89
		// _ = "end of CoverTab[183674]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:90
		_go_fuzz_dep_.CoverTab[183675]++
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:90
		// _ = "end of CoverTab[183675]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:90
	// _ = "end of CoverTab[183672]"
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:90
	_go_fuzz_dep_.CoverTab[183673]++
															return vv[0]
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:91
	// _ = "end of CoverTab[183673]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:100
func (m NiceMD) Del(key string) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:100
	_go_fuzz_dep_.CoverTab[183676]++
															k := strings.ToLower(key)
															delete(m, k)
															return m
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:103
	// _ = "end of CoverTab[183676]"
}

// Set sets the given value in a metadata.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:106
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:106
// It works analogously to http.Header.Set, overwriting all previous metadata values.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:106
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:106
// The function is binary-key safe.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:111
func (m NiceMD) Set(key string, value string) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:111
	_go_fuzz_dep_.CoverTab[183677]++
															k := strings.ToLower(key)
															m[k] = []string{value}
															return m
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:114
	// _ = "end of CoverTab[183677]"
}

// Add retrieves a single value from the metadata.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:117
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:117
// It works analogously to http.Header.Add, as it appends to any existing values associated with key.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:117
//
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:117
// The function is binary-key safe.
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:122
func (m NiceMD) Add(key string, value string) NiceMD {
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:122
	_go_fuzz_dep_.CoverTab[183678]++
															k := strings.ToLower(key)
															m[k] = append(m[k], value)
															return m
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:125
	// _ = "end of CoverTab[183678]"
}

//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:126
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0/util/metautils/nicemd.go:126
var _ = _go_fuzz_dep_.CoverTab
