// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
package auth

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:15
)

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/onosproject/onos-lib-go/pkg/auth"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

const (
	// ContextMetadataTokenKey metadata token key
	ContextMetadataTokenKey = "bearer"
)

// AuthenticationInterceptor an interceptor for authentication
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:33
	_go_fuzz_dep_.CoverTab[190411]++

													tokenString, err := grpc_auth.AuthFromMD(ctx, ContextMetadataTokenKey)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:36
		_go_fuzz_dep_.CoverTab[190425]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:37
		// _ = "end of CoverTab[190425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:38
		_go_fuzz_dep_.CoverTab[190426]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:38
		// _ = "end of CoverTab[190426]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:38
	// _ = "end of CoverTab[190411]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:38
	_go_fuzz_dep_.CoverTab[190412]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:41
	jwtAuth := new(auth.JwtAuthenticator)
	authClaims, err := jwtAuth.ParseAndValidate(tokenString)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:43
		_go_fuzz_dep_.CoverTab[190427]++
														return ctx, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:44
		// _ = "end of CoverTab[190427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:45
		_go_fuzz_dep_.CoverTab[190428]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:45
		// _ = "end of CoverTab[190428]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:45
	// _ = "end of CoverTab[190412]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:45
	_go_fuzz_dep_.CoverTab[190413]++

													niceMd := metautils.ExtractIncoming(ctx)
													niceMd.Del("authorization")
													if name, ok := authClaims["name"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:49
		_go_fuzz_dep_.CoverTab[190429]++
														niceMd.Set("name", name.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:50
		// _ = "end of CoverTab[190429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:51
		_go_fuzz_dep_.CoverTab[190430]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:51
		// _ = "end of CoverTab[190430]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:51
	// _ = "end of CoverTab[190413]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:51
	_go_fuzz_dep_.CoverTab[190414]++
													if email, ok := authClaims["email"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:52
		_go_fuzz_dep_.CoverTab[190431]++
														niceMd.Set("email", email.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:53
		// _ = "end of CoverTab[190431]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:54
		_go_fuzz_dep_.CoverTab[190432]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:54
		// _ = "end of CoverTab[190432]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:54
	// _ = "end of CoverTab[190414]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:54
	_go_fuzz_dep_.CoverTab[190415]++
													if aud, ok := authClaims["aud"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:55
		_go_fuzz_dep_.CoverTab[190433]++
														niceMd.Set("aud", aud.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:56
		// _ = "end of CoverTab[190433]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:57
		_go_fuzz_dep_.CoverTab[190434]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:57
		// _ = "end of CoverTab[190434]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:57
	// _ = "end of CoverTab[190415]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:57
	_go_fuzz_dep_.CoverTab[190416]++
													if exp, ok := authClaims["exp"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:58
		_go_fuzz_dep_.CoverTab[190435]++
														niceMd.Set("exp", fmt.Sprintf("%s", exp))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:59
		// _ = "end of CoverTab[190435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:60
		_go_fuzz_dep_.CoverTab[190436]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:60
		// _ = "end of CoverTab[190436]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:60
	// _ = "end of CoverTab[190416]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:60
	_go_fuzz_dep_.CoverTab[190417]++
													if iat, ok := authClaims["iat"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:61
		_go_fuzz_dep_.CoverTab[190437]++
														niceMd.Set("iat", fmt.Sprintf("%s", iat))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:62
		// _ = "end of CoverTab[190437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:63
		_go_fuzz_dep_.CoverTab[190438]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:63
		// _ = "end of CoverTab[190438]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:63
	// _ = "end of CoverTab[190417]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:63
	_go_fuzz_dep_.CoverTab[190418]++
													if iss, ok := authClaims["iss"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:64
		_go_fuzz_dep_.CoverTab[190439]++
														niceMd.Set("iss", iss.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:65
		// _ = "end of CoverTab[190439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:66
		_go_fuzz_dep_.CoverTab[190440]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:66
		// _ = "end of CoverTab[190440]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:66
	// _ = "end of CoverTab[190418]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:66
	_go_fuzz_dep_.CoverTab[190419]++
													if sub, ok := authClaims["sub"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:67
		_go_fuzz_dep_.CoverTab[190441]++
														niceMd.Set("sub", sub.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:68
		// _ = "end of CoverTab[190441]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:69
		_go_fuzz_dep_.CoverTab[190442]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:69
		// _ = "end of CoverTab[190442]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:69
	// _ = "end of CoverTab[190419]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:69
	_go_fuzz_dep_.CoverTab[190420]++
													if atHash, ok := authClaims["at_hash"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:70
		_go_fuzz_dep_.CoverTab[190443]++
														niceMd.Set("at_hash", atHash.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:71
		// _ = "end of CoverTab[190443]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:72
		_go_fuzz_dep_.CoverTab[190444]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:72
		// _ = "end of CoverTab[190444]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:72
	// _ = "end of CoverTab[190420]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:72
	_go_fuzz_dep_.CoverTab[190421]++
													if preferred, ok := authClaims["preferred_username"]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:73
		_go_fuzz_dep_.CoverTab[190445]++
														niceMd.Set("preferred_username", preferred.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:74
		// _ = "end of CoverTab[190445]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:75
		_go_fuzz_dep_.CoverTab[190446]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:75
		// _ = "end of CoverTab[190446]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:75
	// _ = "end of CoverTab[190421]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:75
	_go_fuzz_dep_.CoverTab[190422]++

													groupsIf, ok := authClaims["groups"].([]interface{})
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:78
		_go_fuzz_dep_.CoverTab[190447]++
														groups := make([]string, 0)
														for _, g := range groupsIf {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:80
			_go_fuzz_dep_.CoverTab[190449]++
															groups = append(groups, g.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:81
			// _ = "end of CoverTab[190449]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:82
		// _ = "end of CoverTab[190447]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:82
		_go_fuzz_dep_.CoverTab[190448]++
														niceMd.Set("groups", strings.Join(groups, ";"))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:83
		// _ = "end of CoverTab[190448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:84
		_go_fuzz_dep_.CoverTab[190450]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:84
		// _ = "end of CoverTab[190450]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:84
	// _ = "end of CoverTab[190422]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:84
	_go_fuzz_dep_.CoverTab[190423]++
													rolesIf, ok := authClaims["roles"].([]interface{})
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:86
		_go_fuzz_dep_.CoverTab[190451]++
														roles := make([]string, 0)
														for _, r := range rolesIf {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:88
			_go_fuzz_dep_.CoverTab[190453]++
															roles = append(roles, r.(string))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:89
			// _ = "end of CoverTab[190453]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:90
		// _ = "end of CoverTab[190451]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:90
		_go_fuzz_dep_.CoverTab[190452]++
														niceMd.Set("roles", strings.Join(roles, ";"))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:91
		// _ = "end of CoverTab[190452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:92
		_go_fuzz_dep_.CoverTab[190454]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:92
		// _ = "end of CoverTab[190454]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:92
	// _ = "end of CoverTab[190423]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:92
	_go_fuzz_dep_.CoverTab[190424]++
													return niceMd.ToIncoming(ctx), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:93
	// _ = "end of CoverTab[190424]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/auth/auth.go:94
var _ = _go_fuzz_dep_.CoverTab
