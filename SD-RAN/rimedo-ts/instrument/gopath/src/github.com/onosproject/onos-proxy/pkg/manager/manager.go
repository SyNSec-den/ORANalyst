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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
package manager

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:15
)

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	e2v1beta1service "github.com/onosproject/onos-proxy/pkg/e2/v1beta1"
	"github.com/onosproject/onos-proxy/pkg/e2/v1beta1/balancer"
	"github.com/onosproject/onos-proxy/pkg/utils/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

var log = logging.GetLogger("onos", "proxy", "manager")

// Config is a manager configuration
type Config struct {
	CAPath		string
	KeyPath		string
	CertPath	string
	GRPCPort	int
}

// NewManager creates a new manager
func NewManager(config Config) *Manager {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:42
	_go_fuzz_dep_.CoverTab[190588]++
													log.Info("Creating Manager")
													return &Manager{
		Config: config,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:46
	// _ = "end of CoverTab[190588]"
}

// Manager is a manager for the E2T service
type Manager struct {
	Config Config
}

// Run starts the manager and the associated services
func (m *Manager) Run() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:55
	_go_fuzz_dep_.CoverTab[190589]++
													log.Info("Running Manager")
													if err := m.Start(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:57
		_go_fuzz_dep_.CoverTab[190590]++
														log.Fatal("Unable to run Manager", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:58
		// _ = "end of CoverTab[190590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:59
		_go_fuzz_dep_.CoverTab[190591]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:59
		// _ = "end of CoverTab[190591]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:59
	// _ = "end of CoverTab[190589]"
}

// Start starts the manager
func (m *Manager) Start() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:63
	_go_fuzz_dep_.CoverTab[190592]++
													err := m.startNorthboundServer()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:65
		_go_fuzz_dep_.CoverTab[190594]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:66
		// _ = "end of CoverTab[190594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:67
		_go_fuzz_dep_.CoverTab[190595]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:67
		// _ = "end of CoverTab[190595]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:67
	// _ = "end of CoverTab[190592]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:67
	_go_fuzz_dep_.CoverTab[190593]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:68
	// _ = "end of CoverTab[190593]"
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:72
	_go_fuzz_dep_.CoverTab[190596]++
													s := northbound.NewServer(&northbound.ServerConfig{
		CaPath:		&m.Config.CAPath,
		KeyPath:	&m.Config.KeyPath,
		CertPath:	&m.Config.CertPath,
		Port:		int16(m.Config.GRPCPort),
		Insecure:	true,
		SecurityCfg:	&northbound.SecurityConfig{},
	})

	conn, err := m.connect(context.Background())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:83
		_go_fuzz_dep_.CoverTab[190599]++
														log.Errorf("Unable to connect to E2T service")
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:85
		// _ = "end of CoverTab[190599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:86
		_go_fuzz_dep_.CoverTab[190600]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:86
		// _ = "end of CoverTab[190600]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:86
	// _ = "end of CoverTab[190596]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:86
	_go_fuzz_dep_.CoverTab[190597]++

													s.AddService(logging.Service{})
													s.AddService(e2v1beta1service.NewProxyService(conn))

													doneCh := make(chan error)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:91
	_curRoutineNum166_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:91
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum166_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
		_go_fuzz_dep_.CoverTab[190601]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
			_go_fuzz_dep_.CoverTab[190603]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum166_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
			// _ = "end of CoverTab[190603]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:92
		}()
														err := s.Serve(func(started string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:93
			_go_fuzz_dep_.CoverTab[190604]++
															log.Info("Started NBI on ", started)
															close(doneCh)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:95
			// _ = "end of CoverTab[190604]"
		})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:96
		// _ = "end of CoverTab[190601]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:96
		_go_fuzz_dep_.CoverTab[190602]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:97
			_go_fuzz_dep_.CoverTab[190605]++
															doneCh <- err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:98
			// _ = "end of CoverTab[190605]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:99
			_go_fuzz_dep_.CoverTab[190606]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:99
			// _ = "end of CoverTab[190606]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:99
		// _ = "end of CoverTab[190602]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:100
	// _ = "end of CoverTab[190597]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:100
	_go_fuzz_dep_.CoverTab[190598]++
													return <-doneCh
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:101
	// _ = "end of CoverTab[190598]"
}

func (m *Manager) connect(ctx context.Context) (*grpc.ClientConn, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:104
	_go_fuzz_dep_.CoverTab[190607]++
													clientCreds, _ := creds.GetClientCredentials()
													conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:///%s", balancer.ResolverName, "onos-e2t:5150"),
		grpc.WithTransportCredentials(credentials.NewTLS(clientCreds)),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable))),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:110
		_go_fuzz_dep_.CoverTab[190609]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:111
		// _ = "end of CoverTab[190609]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:112
		_go_fuzz_dep_.CoverTab[190610]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:112
		// _ = "end of CoverTab[190610]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:112
	// _ = "end of CoverTab[190607]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:112
	_go_fuzz_dep_.CoverTab[190608]++
													return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:113
	// _ = "end of CoverTab[190608]"
}

// Close kills the connections and manager related objects
func (m *Manager) Close() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:117
	_go_fuzz_dep_.CoverTab[190611]++
													log.Info("Closing Manager")
													if err := m.Stop(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:119
		_go_fuzz_dep_.CoverTab[190612]++
														log.Fatal("Unable to Close Manager", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:120
		// _ = "end of CoverTab[190612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:121
		_go_fuzz_dep_.CoverTab[190613]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:121
		// _ = "end of CoverTab[190613]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:121
	// _ = "end of CoverTab[190611]"
}

// Stop stops the manager
func (m *Manager) Stop() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:125
	_go_fuzz_dep_.CoverTab[190614]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:126
	// _ = "end of CoverTab[190614]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:127
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-proxy@v0.1.0/pkg/manager/manager.go:127
var _ = _go_fuzz_dep_.CoverTab
