// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
package connection

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:5
)

import (
	"sync"

	"google.golang.org/grpc"
)

// NewManager creates a new connection manager
func NewManager() *Manager {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:14
	_go_fuzz_dep_.CoverTab[182950]++
															return &Manager{
		conns: make(map[string]*grpc.ClientConn),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:17
	// _ = "end of CoverTab[182950]"
}

// Manager is a connection manager
type Manager struct {
	conns	map[string]*grpc.ClientConn
	mu	sync.RWMutex
}

// Connect connects to the given address
func (m *Manager) Connect(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:27
	_go_fuzz_dep_.CoverTab[182951]++
															m.mu.RLock()
															conn, ok := m.conns[address]
															m.mu.RUnlock()
															if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:31
		_go_fuzz_dep_.CoverTab[182955]++
																return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:32
		// _ = "end of CoverTab[182955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:33
		_go_fuzz_dep_.CoverTab[182956]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:33
		// _ = "end of CoverTab[182956]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:33
	// _ = "end of CoverTab[182951]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:33
	_go_fuzz_dep_.CoverTab[182952]++

															m.mu.Lock()
															defer m.mu.Unlock()

															conn, ok = m.conns[address]
															if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:39
		_go_fuzz_dep_.CoverTab[182957]++
																return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:40
		// _ = "end of CoverTab[182957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:41
		_go_fuzz_dep_.CoverTab[182958]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:41
		// _ = "end of CoverTab[182958]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:41
	// _ = "end of CoverTab[182952]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:41
	_go_fuzz_dep_.CoverTab[182953]++

															conn, err := grpc.Dial(address, opts...)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:44
		_go_fuzz_dep_.CoverTab[182959]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:45
		// _ = "end of CoverTab[182959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:46
		_go_fuzz_dep_.CoverTab[182960]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:46
		// _ = "end of CoverTab[182960]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:46
	// _ = "end of CoverTab[182953]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:46
	_go_fuzz_dep_.CoverTab[182954]++
															m.conns[address] = conn
															return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:48
	// _ = "end of CoverTab[182954]"
}

// Close closes the connection manager
func (m *Manager) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:52
	_go_fuzz_dep_.CoverTab[182961]++
															m.mu.Lock()
															defer m.mu.Unlock()
															var err error
															for _, conn := range m.conns {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:56
		_go_fuzz_dep_.CoverTab[182963]++
																if e := conn.Close(); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:57
			_go_fuzz_dep_.CoverTab[182964]++
																	err = e
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:58
			// _ = "end of CoverTab[182964]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:59
			_go_fuzz_dep_.CoverTab[182965]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:59
			// _ = "end of CoverTab[182965]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:59
		// _ = "end of CoverTab[182963]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:60
	// _ = "end of CoverTab[182961]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:60
	_go_fuzz_dep_.CoverTab[182962]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:61
	// _ = "end of CoverTab[182962]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/connection/manager.go:62
var _ = _go_fuzz_dep_.CoverTab
