// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
package creds

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:5
)

import (
	"crypto/tls"

	"github.com/onosproject/onos-lib-go/pkg/certs"
)

// GetClientCredentials :
func GetClientCredentials() (*tls.Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:14
	_go_fuzz_dep_.CoverTab[182966]++
														cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:16
		_go_fuzz_dep_.CoverTab[182968]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:17
		// _ = "end of CoverTab[182968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:18
		_go_fuzz_dep_.CoverTab[182969]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:18
		// _ = "end of CoverTab[182969]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:18
	// _ = "end of CoverTab[182966]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:18
	_go_fuzz_dep_.CoverTab[182967]++
														return &tls.Config{
		Certificates:		[]tls.Certificate{cert},
		InsecureSkipVerify:	true,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:22
	// _ = "end of CoverTab[182967]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/creds/creds.go:23
var _ = _go_fuzz_dep_.CoverTab
