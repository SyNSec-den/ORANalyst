// Copyright 2019-present Open Networking Foundation.
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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
package certs

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:15
)

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

// HandleCertPaths is a common function for clients and servers like admin/net-changes for
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:26
// handling certificate args if given, or else loading defaults
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:28
func HandleCertPaths(caPath string, keyPath string, certPath string, insecure bool) ([]grpc.DialOption, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:28
	_go_fuzz_dep_.CoverTab[81469]++
													var opts = []grpc.DialOption{}
													var cert tls.Certificate
													var err error
													if keyPath != Client1Key && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
		_go_fuzz_dep_.CoverTab[81473]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
		return keyPath != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
		// _ = "end of CoverTab[81473]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
		_go_fuzz_dep_.CoverTab[81474]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:32
		return certPath != Client1Crt
														// _ = "end of CoverTab[81474]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
		_go_fuzz_dep_.CoverTab[81475]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
		return certPath != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
		// _ = "end of CoverTab[81475]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:33
		_go_fuzz_dep_.CoverTab[81476]++
														cert, err = tls.LoadX509KeyPair(certPath, keyPath)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:35
			_go_fuzz_dep_.CoverTab[81477]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:36
			// _ = "end of CoverTab[81477]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:37
			_go_fuzz_dep_.CoverTab[81478]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:37
			// _ = "end of CoverTab[81478]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:37
		// _ = "end of CoverTab[81476]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:39
		_go_fuzz_dep_.CoverTab[81479]++

														cert, err = tls.X509KeyPair([]byte(DefaultClientCrt), []byte(DefaultClientKey))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:42
			_go_fuzz_dep_.CoverTab[81480]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:43
			// _ = "end of CoverTab[81480]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:44
			_go_fuzz_dep_.CoverTab[81481]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:44
			// _ = "end of CoverTab[81481]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:44
		// _ = "end of CoverTab[81479]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:45
	// _ = "end of CoverTab[81469]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:45
	_go_fuzz_dep_.CoverTab[81470]++
													var clientCAs *x509.CertPool

													if caPath == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:48
		_go_fuzz_dep_.CoverTab[81482]++
														clientCAs, err = GetCertPoolDefault()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:49
		// _ = "end of CoverTab[81482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:50
		_go_fuzz_dep_.CoverTab[81483]++
														clientCAs, err = GetCertPool(caPath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:51
		// _ = "end of CoverTab[81483]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:52
	// _ = "end of CoverTab[81470]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:52
	_go_fuzz_dep_.CoverTab[81471]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:53
		_go_fuzz_dep_.CoverTab[81484]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:54
		// _ = "end of CoverTab[81484]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:55
		_go_fuzz_dep_.CoverTab[81485]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:55
		// _ = "end of CoverTab[81485]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:55
	// _ = "end of CoverTab[81471]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:55
	_go_fuzz_dep_.CoverTab[81472]++

													tlsConfig := &tls.Config{
		Certificates:		[]tls.Certificate{cert},
		ClientCAs:		clientCAs,
		InsecureSkipVerify:	insecure,
	}
													opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))

													return opts, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:64
	// _ = "end of CoverTab[81472]"
}

// GetCertPoolDefault load the default ONF Cert Authority
func GetCertPoolDefault() (*x509.CertPool, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:68
	_go_fuzz_dep_.CoverTab[81486]++
													certPool := x509.NewCertPool()
													if ok := certPool.AppendCertsFromPEM([]byte(OnfCaCrt)); !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:70
		_go_fuzz_dep_.CoverTab[81488]++
														return nil, fmt.Errorf("failed to append default ONF CA certificate")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:71
		// _ = "end of CoverTab[81488]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:72
		_go_fuzz_dep_.CoverTab[81489]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:72
		// _ = "end of CoverTab[81489]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:72
	// _ = "end of CoverTab[81486]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:72
	_go_fuzz_dep_.CoverTab[81487]++
													return certPool, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:73
	// _ = "end of CoverTab[81487]"
}

// GetCertPool loads the Certificate Authority from the given path
func GetCertPool(CaPath string) (*x509.CertPool, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:77
	_go_fuzz_dep_.CoverTab[81490]++
													certPool := x509.NewCertPool()
													ca, err := ioutil.ReadFile(CaPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:80
		_go_fuzz_dep_.CoverTab[81493]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:81
		// _ = "end of CoverTab[81493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:82
		_go_fuzz_dep_.CoverTab[81494]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:82
		// _ = "end of CoverTab[81494]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:82
	// _ = "end of CoverTab[81490]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:82
	_go_fuzz_dep_.CoverTab[81491]++
													if ok := certPool.AppendCertsFromPEM(ca); !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:83
		_go_fuzz_dep_.CoverTab[81495]++
														return nil, fmt.Errorf("failed to append CA certificate from %s", CaPath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:84
		// _ = "end of CoverTab[81495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:85
		_go_fuzz_dep_.CoverTab[81496]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:85
		// _ = "end of CoverTab[81496]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:85
	// _ = "end of CoverTab[81491]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:85
	_go_fuzz_dep_.CoverTab[81492]++
													return certPool, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:86
	// _ = "end of CoverTab[81492]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/certs/certs.go:87
var _ = _go_fuzz_dep_.CoverTab
