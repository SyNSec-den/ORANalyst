// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
package a1endpoint

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:5
)

// NewServer creates a new server struct
func NewServer(caPath string, keyPath string, certPath string, grpcPort int) Server {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:8
	_go_fuzz_dep_.CoverTab[182662]++
														return Server{
		CAPath:		caPath,
		KeyPath:	keyPath,
		CertPath:	certPath,
		GRPCPort:	grpcPort,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:14
	// _ = "end of CoverTab[182662]"
}

// Server is a A1 server
type Server struct {
	CAPath		string
	KeyPath		string
	CertPath	string
	GRPCPort	int
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:23
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/a1/endpoint/server.go:23
var _ = _go_fuzz_dep_.CoverTab
