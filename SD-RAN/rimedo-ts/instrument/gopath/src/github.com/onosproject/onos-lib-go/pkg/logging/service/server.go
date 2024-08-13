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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
package service

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:15
)

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/onosproject/onos-lib-go/pkg/certs"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

// Service provides service-specific registration for grpc services.
type Service interface {
	Register(s *grpc.Server)
}

// Server provides NB gNMI server for onos-lib-go.
type Server struct {
	cfg		*ServerConfig
	services	[]Service
}

// ServerConfig comprises a set of server configuration options.
type ServerConfig struct {
	CaPath		*string
	KeyPath		*string
	CertPath	*string
	Port		int16
	Insecure	bool
}

// NewServer initializes server using the supplied configuration.
func NewServer(cfg *ServerConfig) *Server {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:52
	_go_fuzz_dep_.CoverTab[114724]++
														return &Server{
		services:	[]Service{},
		cfg:		cfg,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:56
	// _ = "end of CoverTab[114724]"
}

// NewServerConfig creates a server config created with the specified end-point security details.
func NewServerConfig(caPath string, keyPath string, certPath string) *ServerConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:60
	_go_fuzz_dep_.CoverTab[114725]++
														return &ServerConfig{
		Port:		5150,
		Insecure:	true,
		CaPath:		&caPath,
		KeyPath:	&keyPath,
		CertPath:	&certPath,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:67
	// _ = "end of CoverTab[114725]"
}

// AddService adds a Service to the server to be registered on Serve.
func (s *Server) AddService(r Service) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:71
	_go_fuzz_dep_.CoverTab[114726]++
														s.services = append(s.services, r)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:72
	// _ = "end of CoverTab[114726]"
}

// Serve starts the NB server.
func (s *Server) Serve(started func(string)) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:76
	_go_fuzz_dep_.CoverTab[114727]++
														lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:78
		_go_fuzz_dep_.CoverTab[114733]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:79
		// _ = "end of CoverTab[114733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:80
		_go_fuzz_dep_.CoverTab[114734]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:80
		// _ = "end of CoverTab[114734]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:80
	// _ = "end of CoverTab[114727]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:80
	_go_fuzz_dep_.CoverTab[114728]++

														tlsCfg := &tls.Config{}

														if *s.cfg.CertPath == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:84
		_go_fuzz_dep_.CoverTab[114735]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:84
		return *s.cfg.KeyPath == ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:84
		// _ = "end of CoverTab[114735]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:84
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:84
		_go_fuzz_dep_.CoverTab[114736]++

															clientCerts, err := tls.X509KeyPair([]byte(certs.DefaultLocalhostCrt), []byte(certs.DefaultLocalhostKey))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:87
			_go_fuzz_dep_.CoverTab[114738]++
																fmt.Println("Error loading default certs")
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:89
			// _ = "end of CoverTab[114738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:90
			_go_fuzz_dep_.CoverTab[114739]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:90
			// _ = "end of CoverTab[114739]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:90
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:90
		// _ = "end of CoverTab[114736]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:90
		_go_fuzz_dep_.CoverTab[114737]++
															tlsCfg.Certificates = []tls.Certificate{clientCerts}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:91
		// _ = "end of CoverTab[114737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:92
		_go_fuzz_dep_.CoverTab[114740]++

															clientCerts, err := tls.LoadX509KeyPair(*s.cfg.CertPath, *s.cfg.KeyPath)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:95
			_go_fuzz_dep_.CoverTab[114742]++
																fmt.Println("Error loading default certs")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:96
			// _ = "end of CoverTab[114742]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:97
			_go_fuzz_dep_.CoverTab[114743]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:97
			// _ = "end of CoverTab[114743]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:97
		// _ = "end of CoverTab[114740]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:97
		_go_fuzz_dep_.CoverTab[114741]++
															tlsCfg.Certificates = []tls.Certificate{clientCerts}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:98
		// _ = "end of CoverTab[114741]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:99
	// _ = "end of CoverTab[114728]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:99
	_go_fuzz_dep_.CoverTab[114729]++

														if s.cfg.Insecure {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:101
		_go_fuzz_dep_.CoverTab[114744]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:105
		tlsCfg.ClientAuth = tls.RequestClientCert
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:105
		// _ = "end of CoverTab[114744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:106
		_go_fuzz_dep_.CoverTab[114745]++
															tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:107
		// _ = "end of CoverTab[114745]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:108
	// _ = "end of CoverTab[114729]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:108
	_go_fuzz_dep_.CoverTab[114730]++

														if *s.cfg.CaPath == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:110
		_go_fuzz_dep_.CoverTab[114746]++

															tlsCfg.ClientCAs = getCertPoolDefault()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:112
		// _ = "end of CoverTab[114746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:113
		_go_fuzz_dep_.CoverTab[114747]++
															tlsCfg.ClientCAs = getCertPool(*s.cfg.CaPath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:114
		// _ = "end of CoverTab[114747]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:115
	// _ = "end of CoverTab[114730]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:115
	_go_fuzz_dep_.CoverTab[114731]++

														opts := []grpc.ServerOption{grpc.Creds(credentials.NewTLS(tlsCfg))}
														server := grpc.NewServer(opts...)
														for i := range s.services {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:119
		_go_fuzz_dep_.CoverTab[114748]++
															s.services[i].Register(server)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:120
		// _ = "end of CoverTab[114748]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:121
	// _ = "end of CoverTab[114731]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:121
	_go_fuzz_dep_.CoverTab[114732]++
														started(lis.Addr().String())

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:125
	return server.Serve(lis)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:125
	// _ = "end of CoverTab[114732]"
}

func getCertPoolDefault() *x509.CertPool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:128
	_go_fuzz_dep_.CoverTab[114749]++
														certPool := x509.NewCertPool()
														if ok := certPool.AppendCertsFromPEM([]byte(certs.OnfCaCrt)); !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:130
		_go_fuzz_dep_.CoverTab[114751]++
															fmt.Println("failed to append CA certificates")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:131
		// _ = "end of CoverTab[114751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:132
		_go_fuzz_dep_.CoverTab[114752]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:132
		// _ = "end of CoverTab[114752]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:132
	// _ = "end of CoverTab[114749]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:132
	_go_fuzz_dep_.CoverTab[114750]++
														return certPool
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:133
	// _ = "end of CoverTab[114750]"
}

func getCertPool(CaPath string) *x509.CertPool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:136
	_go_fuzz_dep_.CoverTab[114753]++
														certPool := x509.NewCertPool()
														ca, err := ioutil.ReadFile(CaPath)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:139
		_go_fuzz_dep_.CoverTab[114756]++
															fmt.Println("could not read ", CaPath, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:140
		// _ = "end of CoverTab[114756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:141
		_go_fuzz_dep_.CoverTab[114757]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:141
		// _ = "end of CoverTab[114757]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:141
	// _ = "end of CoverTab[114753]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:141
	_go_fuzz_dep_.CoverTab[114754]++
														if ok := certPool.AppendCertsFromPEM(ca); !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:142
		_go_fuzz_dep_.CoverTab[114758]++
															fmt.Println("failed to append CA certificates")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:143
		// _ = "end of CoverTab[114758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:144
		_go_fuzz_dep_.CoverTab[114759]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:144
		// _ = "end of CoverTab[114759]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:144
	// _ = "end of CoverTab[114754]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:144
	_go_fuzz_dep_.CoverTab[114755]++
														return certPool
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:145
	// _ = "end of CoverTab[114755]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:146
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/service/server.go:146
var _ = _go_fuzz_dep_.CoverTab
