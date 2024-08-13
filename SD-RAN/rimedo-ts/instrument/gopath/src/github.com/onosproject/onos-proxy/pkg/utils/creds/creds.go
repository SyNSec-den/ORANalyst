// Copyright 2021-present Open Networking Foundation.
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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
package creds

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:15
)

import (
	"crypto/tls"
	"github.com/onosproject/onos-lib-go/pkg/certs"
)

// GetClientCredentials :
func GetClientCredentials() (*tls.Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:23
	_go_fuzz_dep_.CoverTab[190584]++
													cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:25
		_go_fuzz_dep_.CoverTab[190586]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:26
		// _ = "end of CoverTab[190586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:27
		_go_fuzz_dep_.CoverTab[190587]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:27
		// _ = "end of CoverTab[190587]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:27
	// _ = "end of CoverTab[190584]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:27
	_go_fuzz_dep_.CoverTab[190585]++
													return &tls.Config{
		Certificates:		[]tls.Certificate{cert},
		InsecureSkipVerify:	true,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:31
	// _ = "end of CoverTab[190585]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/utils/creds/creds.go:32
var _ = _go_fuzz_dep_.CoverTab
