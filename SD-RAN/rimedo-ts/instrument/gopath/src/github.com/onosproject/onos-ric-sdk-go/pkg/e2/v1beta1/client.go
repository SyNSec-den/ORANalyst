// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
package e2

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:5
)

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"sync"
)

var log = logging.GetLogger("e2", "v1beta1")

// Client is an E2 client
type Client interface {
	// Node returns a Node with the given NodeID
	Node(nodeID NodeID) Node
}

// NewClient creates a new E2 client
func NewClient(opts ...Option) Client {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:21
	_go_fuzz_dep_.CoverTab[196514]++
														return &e2Client{
		opts:	opts,
		nodes:	make(map[NodeID]Node),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:25
	// _ = "end of CoverTab[196514]"
}

// e2Client is the default E2 client implementation
type e2Client struct {
	opts	[]Option
	nodes	map[NodeID]Node
	mu	sync.RWMutex
}

func (c *e2Client) Node(nodeID NodeID) Node {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:35
	_go_fuzz_dep_.CoverTab[196515]++
														c.mu.RLock()
														node, ok := c.nodes[nodeID]
														c.mu.RUnlock()
														if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:39
		_go_fuzz_dep_.CoverTab[196519]++
															return node
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:40
		// _ = "end of CoverTab[196519]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:41
		_go_fuzz_dep_.CoverTab[196520]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:41
		// _ = "end of CoverTab[196520]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:41
	// _ = "end of CoverTab[196515]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:41
	_go_fuzz_dep_.CoverTab[196516]++

														c.mu.Lock()
														defer c.mu.Unlock()
														node, ok = c.nodes[nodeID]
														if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:46
		_go_fuzz_dep_.CoverTab[196521]++
															return node
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:47
		// _ = "end of CoverTab[196521]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:48
		_go_fuzz_dep_.CoverTab[196522]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:48
		// _ = "end of CoverTab[196522]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:48
	// _ = "end of CoverTab[196516]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:48
	_go_fuzz_dep_.CoverTab[196517]++

														node = NewNode(nodeID, c.opts...)
														c.nodes[nodeID] = node
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:51
	_curRoutineNum185_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:51
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum185_)
														go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
		_go_fuzz_dep_.CoverTab[196523]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
			_go_fuzz_dep_.CoverTab[196524]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum185_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
			// _ = "end of CoverTab[196524]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:52
		}()
															<-node.Context().Done()
															c.mu.Lock()
															delete(c.nodes, nodeID)
															c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:56
		// _ = "end of CoverTab[196523]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:57
	// _ = "end of CoverTab[196517]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:57
	_go_fuzz_dep_.CoverTab[196518]++
														return node
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:58
	// _ = "end of CoverTab[196518]"
}

var _ Client = &e2Client{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/client.go:61
var _ = _go_fuzz_dep_.CoverTab
