// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:5
)

import (
	"fmt"
	"math"
	"sort"
)

// RFC 7540, Section 5.3.5: the default weight is 16.
const priorityDefaultWeight = 15	// 16 = 15 + 1

// PriorityWriteSchedulerConfig configures a priorityWriteScheduler.
type PriorityWriteSchedulerConfig struct {
	// MaxClosedNodesInTree controls the maximum number of closed streams to
	// retain in the priority tree. Setting this to zero saves a small amount
	// of memory at the cost of performance.
	//
	// See RFC 7540, Section 5.3.4:
	//   "It is possible for a stream to become closed while prioritization
	//   information ... is in transit. ... This potentially creates suboptimal
	//   prioritization, since the stream could be given a priority that is
	//   different from what is intended. To avoid these problems, an endpoint
	//   SHOULD retain stream prioritization state for a period after streams
	//   become closed. The longer state is retained, the lower the chance that
	//   streams are assigned incorrect or default priority values."
	MaxClosedNodesInTree	int

	// MaxIdleNodesInTree controls the maximum number of idle streams to
	// retain in the priority tree. Setting this to zero saves a small amount
	// of memory at the cost of performance.
	//
	// See RFC 7540, Section 5.3.4:
	//   Similarly, streams that are in the "idle" state can be assigned
	//   priority or become a parent of other streams. This allows for the
	//   creation of a grouping node in the dependency tree, which enables
	//   more flexible expressions of priority. Idle streams begin with a
	//   default priority (Section 5.3.5).
	MaxIdleNodesInTree	int

	// ThrottleOutOfOrderWrites enables write throttling to help ensure that
	// data is delivered in priority order. This works around a race where
	// stream B depends on stream A and both streams are about to call Write
	// to queue DATA frames. If B wins the race, a naive scheduler would eagerly
	// write as much data from B as possible, but this is suboptimal because A
	// is a higher-priority stream. With throttling enabled, we write a small
	// amount of data from B to minimize the amount of bandwidth that B can
	// steal from A.
	ThrottleOutOfOrderWrites	bool
}

// NewPriorityWriteScheduler constructs a WriteScheduler that schedules
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:55
// frames by following HTTP/2 priorities as described in RFC 7540 Section 5.3.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:55
// If cfg is nil, default options are used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:58
func NewPriorityWriteScheduler(cfg *PriorityWriteSchedulerConfig) WriteScheduler {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:58
	_go_fuzz_dep_.CoverTab[75806]++
												if cfg == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:59
		_go_fuzz_dep_.CoverTab[75809]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:62
		cfg = &PriorityWriteSchedulerConfig{
			MaxClosedNodesInTree:		10,
			MaxIdleNodesInTree:		10,
			ThrottleOutOfOrderWrites:	false,
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:66
		// _ = "end of CoverTab[75809]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:67
		_go_fuzz_dep_.CoverTab[75810]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:67
		// _ = "end of CoverTab[75810]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:67
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:67
	// _ = "end of CoverTab[75806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:67
	_go_fuzz_dep_.CoverTab[75807]++

												ws := &priorityWriteScheduler{
		nodes:			make(map[uint32]*priorityNode),
		maxClosedNodesInTree:	cfg.MaxClosedNodesInTree,
		maxIdleNodesInTree:	cfg.MaxIdleNodesInTree,
		enableWriteThrottle:	cfg.ThrottleOutOfOrderWrites,
	}
	ws.nodes[0] = &ws.root
	if cfg.ThrottleOutOfOrderWrites {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:76
		_go_fuzz_dep_.CoverTab[75811]++
													ws.writeThrottleLimit = 1024
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:77
		// _ = "end of CoverTab[75811]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:78
		_go_fuzz_dep_.CoverTab[75812]++
													ws.writeThrottleLimit = math.MaxInt32
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:79
		// _ = "end of CoverTab[75812]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:80
	// _ = "end of CoverTab[75807]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:80
	_go_fuzz_dep_.CoverTab[75808]++
												return ws
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:81
	// _ = "end of CoverTab[75808]"
}

type priorityNodeState int

const (
	priorityNodeOpen	priorityNodeState	= iota
	priorityNodeClosed
	priorityNodeIdle
)

// priorityNode is a node in an HTTP/2 priority tree.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:92
// Each node is associated with a single stream ID.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:92
// See RFC 7540, Section 5.3.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:95
type priorityNode struct {
	q		writeQueue		// queue of pending frames to write
	id		uint32			// id of the stream, or 0 for the root of the tree
	weight		uint8			// the actual weight is weight+1, so the value is in [1,256]
	state		priorityNodeState	// open | closed | idle
	bytes		int64			// number of bytes written by this node, or 0 if closed
	subtreeBytes	int64			// sum(node.bytes) of all nodes in this subtree

	// These links form the priority tree.
	parent		*priorityNode
	kids		*priorityNode	// start of the kids list
	prev, next	*priorityNode	// doubly-linked list of siblings
}

func (n *priorityNode) setParent(parent *priorityNode) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:109
	_go_fuzz_dep_.CoverTab[75813]++
												if n == parent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:110
		_go_fuzz_dep_.CoverTab[75817]++
													panic("setParent to self")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:111
		// _ = "end of CoverTab[75817]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:112
		_go_fuzz_dep_.CoverTab[75818]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:112
		// _ = "end of CoverTab[75818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:112
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:112
	// _ = "end of CoverTab[75813]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:112
	_go_fuzz_dep_.CoverTab[75814]++
												if n.parent == parent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:113
		_go_fuzz_dep_.CoverTab[75819]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:114
		// _ = "end of CoverTab[75819]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:115
		_go_fuzz_dep_.CoverTab[75820]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:115
		// _ = "end of CoverTab[75820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:115
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:115
	// _ = "end of CoverTab[75814]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:115
	_go_fuzz_dep_.CoverTab[75815]++

												if parent := n.parent; parent != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:117
		_go_fuzz_dep_.CoverTab[75821]++
													if n.prev == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:118
			_go_fuzz_dep_.CoverTab[75823]++
														parent.kids = n.next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:119
			// _ = "end of CoverTab[75823]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:120
			_go_fuzz_dep_.CoverTab[75824]++
														n.prev.next = n.next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:121
			// _ = "end of CoverTab[75824]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:122
		// _ = "end of CoverTab[75821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:122
		_go_fuzz_dep_.CoverTab[75822]++
													if n.next != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:123
			_go_fuzz_dep_.CoverTab[75825]++
														n.next.prev = n.prev
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:124
			// _ = "end of CoverTab[75825]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:125
			_go_fuzz_dep_.CoverTab[75826]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:125
			// _ = "end of CoverTab[75826]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:125
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:125
		// _ = "end of CoverTab[75822]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:126
		_go_fuzz_dep_.CoverTab[75827]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:126
		// _ = "end of CoverTab[75827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:126
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:126
	// _ = "end of CoverTab[75815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:126
	_go_fuzz_dep_.CoverTab[75816]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:130
	n.parent = parent
	if parent == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:131
		_go_fuzz_dep_.CoverTab[75828]++
													n.next = nil
													n.prev = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:133
		// _ = "end of CoverTab[75828]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:134
		_go_fuzz_dep_.CoverTab[75829]++
													n.next = parent.kids
													n.prev = nil
													if n.next != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:137
			_go_fuzz_dep_.CoverTab[75831]++
														n.next.prev = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:138
			// _ = "end of CoverTab[75831]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:139
			_go_fuzz_dep_.CoverTab[75832]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:139
			// _ = "end of CoverTab[75832]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:139
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:139
		// _ = "end of CoverTab[75829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:139
		_go_fuzz_dep_.CoverTab[75830]++
													parent.kids = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:140
		// _ = "end of CoverTab[75830]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:141
	// _ = "end of CoverTab[75816]"
}

func (n *priorityNode) addBytes(b int64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:144
	_go_fuzz_dep_.CoverTab[75833]++
												n.bytes += b
												for ; n != nil; n = n.parent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:146
		_go_fuzz_dep_.CoverTab[75834]++
													n.subtreeBytes += b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:147
		// _ = "end of CoverTab[75834]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:148
	// _ = "end of CoverTab[75833]"
}

// walkReadyInOrder iterates over the tree in priority order, calling f for each node
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:151
// with a non-empty write queue. When f returns true, this function returns true and the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:151
// walk halts. tmp is used as scratch space for sorting.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:151
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:151
// f(n, openParent) takes two arguments: the node to visit, n, and a bool that is true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:151
// if any ancestor p of n is still open (ignoring the root node).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:157
func (n *priorityNode) walkReadyInOrder(openParent bool, tmp *[]*priorityNode, f func(*priorityNode, bool) bool) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:157
	_go_fuzz_dep_.CoverTab[75835]++
												if !n.q.empty() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:158
		_go_fuzz_dep_.CoverTab[75844]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:158
		return f(n, openParent)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:158
		// _ = "end of CoverTab[75844]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:158
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:158
		_go_fuzz_dep_.CoverTab[75845]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:159
		// _ = "end of CoverTab[75845]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:160
		_go_fuzz_dep_.CoverTab[75846]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:160
		// _ = "end of CoverTab[75846]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:160
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:160
	// _ = "end of CoverTab[75835]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:160
	_go_fuzz_dep_.CoverTab[75836]++
												if n.kids == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:161
		_go_fuzz_dep_.CoverTab[75847]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:162
		// _ = "end of CoverTab[75847]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:163
		_go_fuzz_dep_.CoverTab[75848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:163
		// _ = "end of CoverTab[75848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:163
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:163
	// _ = "end of CoverTab[75836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:163
	_go_fuzz_dep_.CoverTab[75837]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:167
	if n.id != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:167
		_go_fuzz_dep_.CoverTab[75849]++
													openParent = openParent || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:168
			_go_fuzz_dep_.CoverTab[75850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:168
			return (n.state == priorityNodeOpen)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:168
			// _ = "end of CoverTab[75850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:168
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:168
		// _ = "end of CoverTab[75849]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:169
		_go_fuzz_dep_.CoverTab[75851]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:169
		// _ = "end of CoverTab[75851]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:169
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:169
	// _ = "end of CoverTab[75837]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:169
	_go_fuzz_dep_.CoverTab[75838]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:174
	w := n.kids.weight
	needSort := false
	for k := n.kids.next; k != nil; k = k.next {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:176
		_go_fuzz_dep_.CoverTab[75852]++
													if k.weight != w {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:177
			_go_fuzz_dep_.CoverTab[75853]++
														needSort = true
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:179
			// _ = "end of CoverTab[75853]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:180
			_go_fuzz_dep_.CoverTab[75854]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:180
			// _ = "end of CoverTab[75854]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:180
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:180
		// _ = "end of CoverTab[75852]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:181
	// _ = "end of CoverTab[75838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:181
	_go_fuzz_dep_.CoverTab[75839]++
												if !needSort {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:182
		_go_fuzz_dep_.CoverTab[75855]++
													for k := n.kids; k != nil; k = k.next {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:183
			_go_fuzz_dep_.CoverTab[75857]++
														if k.walkReadyInOrder(openParent, tmp, f) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:184
				_go_fuzz_dep_.CoverTab[75858]++
															return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:185
				// _ = "end of CoverTab[75858]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:186
				_go_fuzz_dep_.CoverTab[75859]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:186
				// _ = "end of CoverTab[75859]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:186
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:186
			// _ = "end of CoverTab[75857]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:187
		// _ = "end of CoverTab[75855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:187
		_go_fuzz_dep_.CoverTab[75856]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:188
		// _ = "end of CoverTab[75856]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:189
		_go_fuzz_dep_.CoverTab[75860]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:189
		// _ = "end of CoverTab[75860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:189
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:189
	// _ = "end of CoverTab[75839]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:189
	_go_fuzz_dep_.CoverTab[75840]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:193
	*tmp = (*tmp)[:0]
	for n.kids != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:194
		_go_fuzz_dep_.CoverTab[75861]++
													*tmp = append(*tmp, n.kids)
													n.kids.setParent(nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:196
		// _ = "end of CoverTab[75861]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:197
	// _ = "end of CoverTab[75840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:197
	_go_fuzz_dep_.CoverTab[75841]++
												sort.Sort(sortPriorityNodeSiblings(*tmp))
												for i := len(*tmp) - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:199
		_go_fuzz_dep_.CoverTab[75862]++
													(*tmp)[i].setParent(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:200
		// _ = "end of CoverTab[75862]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:201
	// _ = "end of CoverTab[75841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:201
	_go_fuzz_dep_.CoverTab[75842]++
												for k := n.kids; k != nil; k = k.next {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:202
		_go_fuzz_dep_.CoverTab[75863]++
													if k.walkReadyInOrder(openParent, tmp, f) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:203
			_go_fuzz_dep_.CoverTab[75864]++
														return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:204
			// _ = "end of CoverTab[75864]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:205
			_go_fuzz_dep_.CoverTab[75865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:205
			// _ = "end of CoverTab[75865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:205
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:205
		// _ = "end of CoverTab[75863]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:206
	// _ = "end of CoverTab[75842]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:206
	_go_fuzz_dep_.CoverTab[75843]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:207
	// _ = "end of CoverTab[75843]"
}

type sortPriorityNodeSiblings []*priorityNode

func (z sortPriorityNodeSiblings) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:212
	_go_fuzz_dep_.CoverTab[75866]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:212
	return len(z)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:212
	// _ = "end of CoverTab[75866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:212
}
func (z sortPriorityNodeSiblings) Swap(i, k int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:213
	_go_fuzz_dep_.CoverTab[75867]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:213
	z[i], z[k] = z[k], z[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:213
	// _ = "end of CoverTab[75867]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:213
}
func (z sortPriorityNodeSiblings) Less(i, k int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:214
	_go_fuzz_dep_.CoverTab[75868]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:217
	wi, bi := float64(z[i].weight+1), float64(z[i].subtreeBytes)
	wk, bk := float64(z[k].weight+1), float64(z[k].subtreeBytes)
	if bi == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:219
		_go_fuzz_dep_.CoverTab[75871]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:219
		return bk == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:219
		// _ = "end of CoverTab[75871]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:219
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:219
		_go_fuzz_dep_.CoverTab[75872]++
													return wi >= wk
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:220
		// _ = "end of CoverTab[75872]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:221
		_go_fuzz_dep_.CoverTab[75873]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:221
		// _ = "end of CoverTab[75873]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:221
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:221
	// _ = "end of CoverTab[75868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:221
	_go_fuzz_dep_.CoverTab[75869]++
												if bk == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:222
		_go_fuzz_dep_.CoverTab[75874]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:223
		// _ = "end of CoverTab[75874]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:224
		_go_fuzz_dep_.CoverTab[75875]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:224
		// _ = "end of CoverTab[75875]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:224
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:224
	// _ = "end of CoverTab[75869]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:224
	_go_fuzz_dep_.CoverTab[75870]++
												return bi/bk <= wi/wk
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:225
	// _ = "end of CoverTab[75870]"
}

type priorityWriteScheduler struct {
	// root is the root of the priority tree, where root.id = 0.
	// The root queues control frames that are not associated with any stream.
	root	priorityNode

	// nodes maps stream ids to priority tree nodes.
	nodes	map[uint32]*priorityNode

	// maxID is the maximum stream id in nodes.
	maxID	uint32

	// lists of nodes that have been closed or are idle, but are kept in
	// the tree for improved prioritization. When the lengths exceed either
	// maxClosedNodesInTree or maxIdleNodesInTree, old nodes are discarded.
	closedNodes, idleNodes	[]*priorityNode

	// From the config.
	maxClosedNodesInTree	int
	maxIdleNodesInTree	int
	writeThrottleLimit	int32
	enableWriteThrottle	bool

	// tmp is scratch space for priorityNode.walkReadyInOrder to reduce allocations.
	tmp	[]*priorityNode

	// pool of empty queues for reuse.
	queuePool	writeQueuePool
}

func (ws *priorityWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:257
	_go_fuzz_dep_.CoverTab[75876]++

												if curr := ws.nodes[streamID]; curr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:259
		_go_fuzz_dep_.CoverTab[75879]++
													if curr.state != priorityNodeIdle {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:260
			_go_fuzz_dep_.CoverTab[75881]++
														panic(fmt.Sprintf("stream %d already opened", streamID))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:261
			// _ = "end of CoverTab[75881]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:262
			_go_fuzz_dep_.CoverTab[75882]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:262
			// _ = "end of CoverTab[75882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:262
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:262
		// _ = "end of CoverTab[75879]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:262
		_go_fuzz_dep_.CoverTab[75880]++
													curr.state = priorityNodeOpen
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:264
		// _ = "end of CoverTab[75880]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:265
		_go_fuzz_dep_.CoverTab[75883]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:265
		// _ = "end of CoverTab[75883]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:265
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:265
	// _ = "end of CoverTab[75876]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:265
	_go_fuzz_dep_.CoverTab[75877]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:271
	parent := ws.nodes[options.PusherID]
	if parent == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:272
		_go_fuzz_dep_.CoverTab[75884]++
													parent = &ws.root
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:273
		// _ = "end of CoverTab[75884]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:274
		_go_fuzz_dep_.CoverTab[75885]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:274
		// _ = "end of CoverTab[75885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:274
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:274
	// _ = "end of CoverTab[75877]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:274
	_go_fuzz_dep_.CoverTab[75878]++
												n := &priorityNode{
		q:	*ws.queuePool.get(),
		id:	streamID,
		weight:	priorityDefaultWeight,
		state:	priorityNodeOpen,
	}
	n.setParent(parent)
	ws.nodes[streamID] = n
	if streamID > ws.maxID {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:283
		_go_fuzz_dep_.CoverTab[75886]++
													ws.maxID = streamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:284
		// _ = "end of CoverTab[75886]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:285
		_go_fuzz_dep_.CoverTab[75887]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:285
		// _ = "end of CoverTab[75887]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:285
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:285
	// _ = "end of CoverTab[75878]"
}

func (ws *priorityWriteScheduler) CloseStream(streamID uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:288
	_go_fuzz_dep_.CoverTab[75888]++
												if streamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:289
		_go_fuzz_dep_.CoverTab[75892]++
													panic("violation of WriteScheduler interface: cannot close stream 0")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:290
		// _ = "end of CoverTab[75892]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:291
		_go_fuzz_dep_.CoverTab[75893]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:291
		// _ = "end of CoverTab[75893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:291
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:291
	// _ = "end of CoverTab[75888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:291
	_go_fuzz_dep_.CoverTab[75889]++
												if ws.nodes[streamID] == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:292
		_go_fuzz_dep_.CoverTab[75894]++
													panic(fmt.Sprintf("violation of WriteScheduler interface: unknown stream %d", streamID))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:293
		// _ = "end of CoverTab[75894]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:294
		_go_fuzz_dep_.CoverTab[75895]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:294
		// _ = "end of CoverTab[75895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:294
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:294
	// _ = "end of CoverTab[75889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:294
	_go_fuzz_dep_.CoverTab[75890]++
												if ws.nodes[streamID].state != priorityNodeOpen {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:295
		_go_fuzz_dep_.CoverTab[75896]++
													panic(fmt.Sprintf("violation of WriteScheduler interface: stream %d already closed", streamID))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:296
		// _ = "end of CoverTab[75896]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:297
		_go_fuzz_dep_.CoverTab[75897]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:297
		// _ = "end of CoverTab[75897]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:297
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:297
	// _ = "end of CoverTab[75890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:297
	_go_fuzz_dep_.CoverTab[75891]++

												n := ws.nodes[streamID]
												n.state = priorityNodeClosed
												n.addBytes(-n.bytes)

												q := n.q
												ws.queuePool.put(&q)
												n.q.s = nil
												if ws.maxClosedNodesInTree > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:306
		_go_fuzz_dep_.CoverTab[75898]++
													ws.addClosedOrIdleNode(&ws.closedNodes, ws.maxClosedNodesInTree, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:307
		// _ = "end of CoverTab[75898]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:308
		_go_fuzz_dep_.CoverTab[75899]++
													ws.removeNode(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:309
		// _ = "end of CoverTab[75899]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:310
	// _ = "end of CoverTab[75891]"
}

func (ws *priorityWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:313
	_go_fuzz_dep_.CoverTab[75900]++
												if streamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:314
		_go_fuzz_dep_.CoverTab[75907]++
													panic("adjustPriority on root")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:315
		// _ = "end of CoverTab[75907]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:316
		_go_fuzz_dep_.CoverTab[75908]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:316
		// _ = "end of CoverTab[75908]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:316
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:316
	// _ = "end of CoverTab[75900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:316
	_go_fuzz_dep_.CoverTab[75901]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:321
	n := ws.nodes[streamID]
	if n == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:322
		_go_fuzz_dep_.CoverTab[75909]++
													if streamID <= ws.maxID || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:323
			_go_fuzz_dep_.CoverTab[75911]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:323
			return ws.maxIdleNodesInTree == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:323
			// _ = "end of CoverTab[75911]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:323
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:323
			_go_fuzz_dep_.CoverTab[75912]++
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:324
			// _ = "end of CoverTab[75912]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:325
			_go_fuzz_dep_.CoverTab[75913]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:325
			// _ = "end of CoverTab[75913]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:325
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:325
		// _ = "end of CoverTab[75909]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:325
		_go_fuzz_dep_.CoverTab[75910]++
													ws.maxID = streamID
													n = &priorityNode{
			q:	*ws.queuePool.get(),
			id:	streamID,
			weight:	priorityDefaultWeight,
			state:	priorityNodeIdle,
		}
													n.setParent(&ws.root)
													ws.nodes[streamID] = n
													ws.addClosedOrIdleNode(&ws.idleNodes, ws.maxIdleNodesInTree, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:335
		// _ = "end of CoverTab[75910]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:336
		_go_fuzz_dep_.CoverTab[75914]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:336
		// _ = "end of CoverTab[75914]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:336
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:336
	// _ = "end of CoverTab[75901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:336
	_go_fuzz_dep_.CoverTab[75902]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:340
	parent := ws.nodes[priority.StreamDep]
	if parent == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:341
		_go_fuzz_dep_.CoverTab[75915]++
													n.setParent(&ws.root)
													n.weight = priorityDefaultWeight
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:344
		// _ = "end of CoverTab[75915]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:345
		_go_fuzz_dep_.CoverTab[75916]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:345
		// _ = "end of CoverTab[75916]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:345
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:345
	// _ = "end of CoverTab[75902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:345
	_go_fuzz_dep_.CoverTab[75903]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:348
	if n == parent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:348
		_go_fuzz_dep_.CoverTab[75917]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:349
		// _ = "end of CoverTab[75917]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:350
		_go_fuzz_dep_.CoverTab[75918]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:350
		// _ = "end of CoverTab[75918]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:350
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:350
	// _ = "end of CoverTab[75903]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:350
	_go_fuzz_dep_.CoverTab[75904]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:359
	for x := parent.parent; x != nil; x = x.parent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:359
		_go_fuzz_dep_.CoverTab[75919]++
													if x == n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:360
			_go_fuzz_dep_.CoverTab[75920]++
														parent.setParent(n.parent)
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:362
			// _ = "end of CoverTab[75920]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:363
			_go_fuzz_dep_.CoverTab[75921]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:363
			// _ = "end of CoverTab[75921]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:363
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:363
		// _ = "end of CoverTab[75919]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:364
	// _ = "end of CoverTab[75904]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:364
	_go_fuzz_dep_.CoverTab[75905]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:369
	if priority.Exclusive {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:369
		_go_fuzz_dep_.CoverTab[75922]++
													k := parent.kids
													for k != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:371
			_go_fuzz_dep_.CoverTab[75923]++
														next := k.next
														if k != n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:373
				_go_fuzz_dep_.CoverTab[75925]++
															k.setParent(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:374
				// _ = "end of CoverTab[75925]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:375
				_go_fuzz_dep_.CoverTab[75926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:375
				// _ = "end of CoverTab[75926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:375
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:375
			// _ = "end of CoverTab[75923]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:375
			_go_fuzz_dep_.CoverTab[75924]++
														k = next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:376
			// _ = "end of CoverTab[75924]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:377
		// _ = "end of CoverTab[75922]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:378
		_go_fuzz_dep_.CoverTab[75927]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:378
		// _ = "end of CoverTab[75927]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:378
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:378
	// _ = "end of CoverTab[75905]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:378
	_go_fuzz_dep_.CoverTab[75906]++

												n.setParent(parent)
												n.weight = priority.Weight
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:381
	// _ = "end of CoverTab[75906]"
}

func (ws *priorityWriteScheduler) Push(wr FrameWriteRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:384
	_go_fuzz_dep_.CoverTab[75928]++
												var n *priorityNode
												if wr.isControl() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:386
		_go_fuzz_dep_.CoverTab[75930]++
													n = &ws.root
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:387
		// _ = "end of CoverTab[75930]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:388
		_go_fuzz_dep_.CoverTab[75931]++
													id := wr.StreamID()
													n = ws.nodes[id]
													if n == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:391
			_go_fuzz_dep_.CoverTab[75932]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:395
			if wr.DataSize() > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:395
				_go_fuzz_dep_.CoverTab[75934]++
															panic("add DATA on non-open stream")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:396
				// _ = "end of CoverTab[75934]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:397
				_go_fuzz_dep_.CoverTab[75935]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:397
				// _ = "end of CoverTab[75935]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:397
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:397
			// _ = "end of CoverTab[75932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:397
			_go_fuzz_dep_.CoverTab[75933]++
														n = &ws.root
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:398
			// _ = "end of CoverTab[75933]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:399
			_go_fuzz_dep_.CoverTab[75936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:399
			// _ = "end of CoverTab[75936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:399
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:399
		// _ = "end of CoverTab[75931]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:400
	// _ = "end of CoverTab[75928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:400
	_go_fuzz_dep_.CoverTab[75929]++
												n.q.push(wr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:401
	// _ = "end of CoverTab[75929]"
}

func (ws *priorityWriteScheduler) Pop() (wr FrameWriteRequest, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:404
	_go_fuzz_dep_.CoverTab[75937]++
												ws.root.walkReadyInOrder(false, &ws.tmp, func(n *priorityNode, openParent bool) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:405
		_go_fuzz_dep_.CoverTab[75939]++
													limit := int32(math.MaxInt32)
													if openParent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:407
			_go_fuzz_dep_.CoverTab[75943]++
														limit = ws.writeThrottleLimit
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:408
			// _ = "end of CoverTab[75943]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:409
			_go_fuzz_dep_.CoverTab[75944]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:409
			// _ = "end of CoverTab[75944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:409
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:409
		// _ = "end of CoverTab[75939]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:409
		_go_fuzz_dep_.CoverTab[75940]++
													wr, ok = n.q.consume(limit)
													if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:411
			_go_fuzz_dep_.CoverTab[75945]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:412
			// _ = "end of CoverTab[75945]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:413
			_go_fuzz_dep_.CoverTab[75946]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:413
			// _ = "end of CoverTab[75946]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:413
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:413
		// _ = "end of CoverTab[75940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:413
		_go_fuzz_dep_.CoverTab[75941]++
													n.addBytes(int64(wr.DataSize()))

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:418
		if openParent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:418
			_go_fuzz_dep_.CoverTab[75947]++
														ws.writeThrottleLimit += 1024
														if ws.writeThrottleLimit < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:420
				_go_fuzz_dep_.CoverTab[75948]++
															ws.writeThrottleLimit = math.MaxInt32
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:421
				// _ = "end of CoverTab[75948]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:422
				_go_fuzz_dep_.CoverTab[75949]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:422
				// _ = "end of CoverTab[75949]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:422
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:422
			// _ = "end of CoverTab[75947]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:423
			_go_fuzz_dep_.CoverTab[75950]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:423
			if ws.enableWriteThrottle {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:423
				_go_fuzz_dep_.CoverTab[75951]++
															ws.writeThrottleLimit = 1024
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:424
				// _ = "end of CoverTab[75951]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
				_go_fuzz_dep_.CoverTab[75952]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
				// _ = "end of CoverTab[75952]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
			// _ = "end of CoverTab[75950]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
		// _ = "end of CoverTab[75941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:425
		_go_fuzz_dep_.CoverTab[75942]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:426
		// _ = "end of CoverTab[75942]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:427
	// _ = "end of CoverTab[75937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:427
	_go_fuzz_dep_.CoverTab[75938]++
												return wr, ok
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:428
	// _ = "end of CoverTab[75938]"
}

func (ws *priorityWriteScheduler) addClosedOrIdleNode(list *[]*priorityNode, maxSize int, n *priorityNode) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:431
	_go_fuzz_dep_.CoverTab[75953]++
												if maxSize == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:432
		_go_fuzz_dep_.CoverTab[75956]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:433
		// _ = "end of CoverTab[75956]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:434
		_go_fuzz_dep_.CoverTab[75957]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:434
		// _ = "end of CoverTab[75957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:434
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:434
	// _ = "end of CoverTab[75953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:434
	_go_fuzz_dep_.CoverTab[75954]++
												if len(*list) == maxSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:435
		_go_fuzz_dep_.CoverTab[75958]++

													ws.removeNode((*list)[0])
													x := (*list)[1:]
													copy(*list, x)
													*list = (*list)[:len(x)]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:440
		// _ = "end of CoverTab[75958]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:441
		_go_fuzz_dep_.CoverTab[75959]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:441
		// _ = "end of CoverTab[75959]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:441
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:441
	// _ = "end of CoverTab[75954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:441
	_go_fuzz_dep_.CoverTab[75955]++
												*list = append(*list, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:442
	// _ = "end of CoverTab[75955]"
}

func (ws *priorityWriteScheduler) removeNode(n *priorityNode) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:445
	_go_fuzz_dep_.CoverTab[75960]++
												for k := n.kids; k != nil; k = k.next {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:446
		_go_fuzz_dep_.CoverTab[75962]++
													k.setParent(n.parent)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:447
		// _ = "end of CoverTab[75962]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:448
	// _ = "end of CoverTab[75960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:448
	_go_fuzz_dep_.CoverTab[75961]++
												n.setParent(nil)
												delete(ws.nodes, n.id)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:450
	// _ = "end of CoverTab[75961]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:451
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_priority.go:451
var _ = _go_fuzz_dep_.CoverTab
