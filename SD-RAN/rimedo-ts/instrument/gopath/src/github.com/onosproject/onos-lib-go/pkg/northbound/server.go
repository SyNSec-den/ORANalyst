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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:15
// Package northbound houses implementations of various application-oriented interfaces
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:15
// for the ONOS configuration subsystem.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
package northbound

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:17
)

import (
	"crypto/tls"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/grpc/auth"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound")

// Service provides service-specific registration for grpc services.
type Service interface {
	Register(s *grpc.Server)
}

// Server provides NB gNMI server for onos-config.
type Server struct {
	cfg		*ServerConfig
	services	[]Service
	server		*grpc.Server
}

// SecurityConfig security configuration
type SecurityConfig struct {
	AuthenticationEnabled	bool
	AuthorizationEnabled	bool
}

// ServerConfig comprises a set of server configuration options.
type ServerConfig struct {
	CaPath		*string
	KeyPath		*string
	CertPath	*string
	Port		int16
	Insecure	bool
	SecurityCfg	*SecurityConfig
}

// NewServer initializes gNMI server using the supplied configuration.
func NewServer(cfg *ServerConfig) *Server {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:66
	_go_fuzz_dep_.CoverTab[190455]++
														return &Server{
		services:	[]Service{},
		cfg:		cfg,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:70
	// _ = "end of CoverTab[190455]"
}

// NewServerConfig creates a server config created with the specified end-point security details.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:73
// Deprecated: Use NewServerCfg instead
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:75
func NewServerConfig(caPath string, keyPath string, certPath string, port int16, insecure bool) *ServerConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:75
	_go_fuzz_dep_.CoverTab[190456]++
														return &ServerConfig{
		Port:		port,
		Insecure:	insecure,
		CaPath:		&caPath,
		KeyPath:	&keyPath,
		CertPath:	&certPath,
		SecurityCfg:	&SecurityConfig{},
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:83
	// _ = "end of CoverTab[190456]"
}

// NewServerCfg creates a server config created with the specified end-point security details.
func NewServerCfg(caPath string, keyPath string, certPath string, port int16, insecure bool, secCfg SecurityConfig) *ServerConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:87
	_go_fuzz_dep_.CoverTab[190457]++
														return &ServerConfig{
		Port:		port,
		Insecure:	insecure,
		CaPath:		&caPath,
		KeyPath:	&keyPath,
		CertPath:	&certPath,
		SecurityCfg:	&secCfg,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:95
	// _ = "end of CoverTab[190457]"
}

// AddService adds a Service to the server to be registered on Serve.
func (s *Server) AddService(r Service) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:99
	_go_fuzz_dep_.CoverTab[190458]++
														s.services = append(s.services, r)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:100
	// _ = "end of CoverTab[190458]"
}

// Serve starts the NB gNMI server.
func (s *Server) Serve(started func(string)) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:104
	_go_fuzz_dep_.CoverTab[190459]++
														lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:106
		_go_fuzz_dep_.CoverTab[190467]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:107
		// _ = "end of CoverTab[190467]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:108
		_go_fuzz_dep_.CoverTab[190468]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:108
		// _ = "end of CoverTab[190468]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:108
	// _ = "end of CoverTab[190459]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:108
	_go_fuzz_dep_.CoverTab[190460]++
														tlsCfg := &tls.Config{}

														if *s.cfg.CertPath == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:111
		_go_fuzz_dep_.CoverTab[190469]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:111
		return *s.cfg.KeyPath == ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:111
		// _ = "end of CoverTab[190469]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:111
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:111
		_go_fuzz_dep_.CoverTab[190470]++

															clientCerts, err := tls.X509KeyPair([]byte(certs.DefaultLocalhostCrt), []byte(certs.DefaultLocalhostKey))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:114
			_go_fuzz_dep_.CoverTab[190472]++
																log.Error("Error loading default certs")
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:116
			// _ = "end of CoverTab[190472]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:117
			_go_fuzz_dep_.CoverTab[190473]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:117
			// _ = "end of CoverTab[190473]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:117
		// _ = "end of CoverTab[190470]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:117
		_go_fuzz_dep_.CoverTab[190471]++
															tlsCfg.Certificates = []tls.Certificate{clientCerts}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:118
		// _ = "end of CoverTab[190471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:119
		_go_fuzz_dep_.CoverTab[190474]++
															log.Infof("Loading certs: %s %s", *s.cfg.CertPath, *s.cfg.KeyPath)
															clientCerts, err := tls.LoadX509KeyPair(*s.cfg.CertPath, *s.cfg.KeyPath)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:122
			_go_fuzz_dep_.CoverTab[190476]++
																log.Info("Error loading default certs")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:123
			// _ = "end of CoverTab[190476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:124
			_go_fuzz_dep_.CoverTab[190477]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:124
			// _ = "end of CoverTab[190477]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:124
		// _ = "end of CoverTab[190474]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:124
		_go_fuzz_dep_.CoverTab[190475]++
															tlsCfg.Certificates = []tls.Certificate{clientCerts}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:125
		// _ = "end of CoverTab[190475]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:126
	// _ = "end of CoverTab[190460]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:126
	_go_fuzz_dep_.CoverTab[190461]++

														if s.cfg.Insecure {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:128
		_go_fuzz_dep_.CoverTab[190478]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:132
		tlsCfg.ClientAuth = tls.RequestClientCert
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:132
		// _ = "end of CoverTab[190478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:133
		_go_fuzz_dep_.CoverTab[190479]++
															tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:134
		// _ = "end of CoverTab[190479]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:135
	// _ = "end of CoverTab[190461]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:135
	_go_fuzz_dep_.CoverTab[190462]++

														if *s.cfg.CaPath == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:137
		_go_fuzz_dep_.CoverTab[190480]++
															log.Info("Loading default CA onfca")
															tlsCfg.ClientCAs, err = certs.GetCertPoolDefault()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:139
		// _ = "end of CoverTab[190480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:140
		_go_fuzz_dep_.CoverTab[190481]++
															tlsCfg.ClientCAs, err = certs.GetCertPool(*s.cfg.CaPath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:141
		// _ = "end of CoverTab[190481]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:142
	// _ = "end of CoverTab[190462]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:142
	_go_fuzz_dep_.CoverTab[190463]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:143
		_go_fuzz_dep_.CoverTab[190482]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:144
		// _ = "end of CoverTab[190482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:145
		_go_fuzz_dep_.CoverTab[190483]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:145
		// _ = "end of CoverTab[190483]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:145
	// _ = "end of CoverTab[190463]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:145
	_go_fuzz_dep_.CoverTab[190464]++
														opts := []grpc.ServerOption{grpc.Creds(credentials.NewTLS(tlsCfg))}
														if s.cfg.SecurityCfg.AuthenticationEnabled {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:147
		_go_fuzz_dep_.CoverTab[190484]++
															log.Info("Authentication Enabled")
															opts = append(opts, grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_auth.UnaryServerInterceptor(auth.AuthenticationInterceptor),
			)))
		opts = append(opts, grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_auth.StreamServerInterceptor(auth.AuthenticationInterceptor))))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:155
		// _ = "end of CoverTab[190484]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:157
		_go_fuzz_dep_.CoverTab[190485]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:157
		// _ = "end of CoverTab[190485]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:157
	// _ = "end of CoverTab[190464]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:157
	_go_fuzz_dep_.CoverTab[190465]++

														s.server = grpc.NewServer(opts...)
														for i := range s.services {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:160
		_go_fuzz_dep_.CoverTab[190486]++
															s.services[i].Register(s.server)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:161
		// _ = "end of CoverTab[190486]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:162
	// _ = "end of CoverTab[190465]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:162
	_go_fuzz_dep_.CoverTab[190466]++
														started(lis.Addr().String())

														log.Infof("Starting RPC server on address: %s", lis.Addr().String())
														return s.server.Serve(lis)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:166
	// _ = "end of CoverTab[190466]"
}

// Stop stops the server.
func (s *Server) Stop() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:170
	_go_fuzz_dep_.CoverTab[190487]++
														s.server.Stop()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:171
	// _ = "end of CoverTab[190487]"
}

// GracefulStop stops the server gracefully.
func (s *Server) GracefulStop() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:175
	_go_fuzz_dep_.CoverTab[190488]++
														s.server.GracefulStop()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:176
	// _ = "end of CoverTab[190488]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:177
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/northbound/server.go:177
var _ = _go_fuzz_dep_.CoverTab
