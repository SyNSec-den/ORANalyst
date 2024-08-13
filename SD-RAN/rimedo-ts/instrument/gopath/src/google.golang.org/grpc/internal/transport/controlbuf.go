//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:19
)

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
	"google.golang.org/grpc/internal/grpcutil"
	"google.golang.org/grpc/status"
)

var updateHeaderTblSize = func(e *hpack.Encoder, v uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:37
	_go_fuzz_dep_.CoverTab[76586]++
													e.SetMaxDynamicTableSizeLimit(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:38
	// _ = "end of CoverTab[76586]"
}

type itemNode struct {
	it	interface{}
	next	*itemNode
}

type itemList struct {
	head	*itemNode
	tail	*itemNode
}

func (il *itemList) enqueue(i interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:51
	_go_fuzz_dep_.CoverTab[76587]++
													n := &itemNode{it: i}
													if il.tail == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:53
		_go_fuzz_dep_.CoverTab[76589]++
														il.head, il.tail = n, n
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:55
		// _ = "end of CoverTab[76589]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:56
		_go_fuzz_dep_.CoverTab[76590]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:56
		// _ = "end of CoverTab[76590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:56
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:56
	// _ = "end of CoverTab[76587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:56
	_go_fuzz_dep_.CoverTab[76588]++
													il.tail.next = n
													il.tail = n
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:58
	// _ = "end of CoverTab[76588]"
}

// peek returns the first item in the list without removing it from the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:61
// list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:63
func (il *itemList) peek() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:63
	_go_fuzz_dep_.CoverTab[76591]++
													return il.head.it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:64
	// _ = "end of CoverTab[76591]"
}

func (il *itemList) dequeue() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:67
	_go_fuzz_dep_.CoverTab[76592]++
													if il.head == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:68
		_go_fuzz_dep_.CoverTab[76595]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:69
		// _ = "end of CoverTab[76595]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:70
		_go_fuzz_dep_.CoverTab[76596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:70
		// _ = "end of CoverTab[76596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:70
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:70
	// _ = "end of CoverTab[76592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:70
	_go_fuzz_dep_.CoverTab[76593]++
													i := il.head.it
													il.head = il.head.next
													if il.head == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:73
		_go_fuzz_dep_.CoverTab[76597]++
														il.tail = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:74
		// _ = "end of CoverTab[76597]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:75
		_go_fuzz_dep_.CoverTab[76598]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:75
		// _ = "end of CoverTab[76598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:75
	// _ = "end of CoverTab[76593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:75
	_go_fuzz_dep_.CoverTab[76594]++
													return i
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:76
	// _ = "end of CoverTab[76594]"
}

func (il *itemList) dequeueAll() *itemNode {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:79
	_go_fuzz_dep_.CoverTab[76599]++
													h := il.head
													il.head, il.tail = nil, nil
													return h
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:82
	// _ = "end of CoverTab[76599]"
}

func (il *itemList) isEmpty() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:85
	_go_fuzz_dep_.CoverTab[76600]++
													return il.head == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:86
	// _ = "end of CoverTab[76600]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:93
// maxQueuedTransportResponseFrames is the most queued "transport response"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:93
// frames we will buffer before preventing new reads from occurring on the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:93
// transport.  These are control frames sent in response to client requests,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:93
// such as RST_STREAM due to bad headers or settings acks.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:97
const maxQueuedTransportResponseFrames = 50

type cbItem interface {
	isTransportResponseFrame() bool
}

// registerStream is used to register an incoming stream with loopy writer.
type registerStream struct {
	streamID	uint32
	wq		*writeQuota
}

func (*registerStream) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:109
	_go_fuzz_dep_.CoverTab[76601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:109
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:109
	// _ = "end of CoverTab[76601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:109
}

// headerFrame is also used to register stream on the client-side.
type headerFrame struct {
	streamID	uint32
	hf		[]hpack.HeaderField
	endStream	bool			// Valid on server side.
	initStream	func(uint32) error	// Used only on the client side.
	onWrite		func()
	wq		*writeQuota	// write quota for the stream created.
	cleanup		*cleanupStream	// Valid on the server side.
	onOrphaned	func(error)	// Valid on client-side
}

func (h *headerFrame) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:123
	_go_fuzz_dep_.CoverTab[76602]++
													return h.cleanup != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:124
		_go_fuzz_dep_.CoverTab[76603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:124
		return h.cleanup.rst
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:124
		// _ = "end of CoverTab[76603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:124
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:124
	// _ = "end of CoverTab[76602]"
}

type cleanupStream struct {
	streamID	uint32
	rst		bool
	rstCode		http2.ErrCode
	onWrite		func()
}

func (c *cleanupStream) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:134
	_go_fuzz_dep_.CoverTab[76604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:134
	return c.rst
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:134
	// _ = "end of CoverTab[76604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:134
}

type earlyAbortStream struct {
	httpStatus	uint32
	streamID	uint32
	contentSubtype	string
	status		*status.Status
	rst		bool
}

func (*earlyAbortStream) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:144
	_go_fuzz_dep_.CoverTab[76605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:144
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:144
	// _ = "end of CoverTab[76605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:144
}

type dataFrame struct {
	streamID	uint32
	endStream	bool
	h		[]byte
	d		[]byte
	// onEachWrite is called every time
	// a part of d is written out.
	onEachWrite	func()
}

func (*dataFrame) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:156
	_go_fuzz_dep_.CoverTab[76606]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:156
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:156
	// _ = "end of CoverTab[76606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:156
}

type incomingWindowUpdate struct {
	streamID	uint32
	increment	uint32
}

func (*incomingWindowUpdate) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:163
	_go_fuzz_dep_.CoverTab[76607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:163
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:163
	// _ = "end of CoverTab[76607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:163
}

type outgoingWindowUpdate struct {
	streamID	uint32
	increment	uint32
}

func (*outgoingWindowUpdate) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:170
	_go_fuzz_dep_.CoverTab[76608]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:171
	// _ = "end of CoverTab[76608]"
}

type incomingSettings struct {
	ss []http2.Setting
}

func (*incomingSettings) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:178
	_go_fuzz_dep_.CoverTab[76609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:178
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:178
	// _ = "end of CoverTab[76609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:178
}

type outgoingSettings struct {
	ss []http2.Setting
}

func (*outgoingSettings) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:184
	_go_fuzz_dep_.CoverTab[76610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:184
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:184
	// _ = "end of CoverTab[76610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:184
}

type incomingGoAway struct {
}

func (*incomingGoAway) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:189
	_go_fuzz_dep_.CoverTab[76611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:189
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:189
	// _ = "end of CoverTab[76611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:189
}

type goAway struct {
	code		http2.ErrCode
	debugData	[]byte
	headsUp		bool
	closeConn	error	// if set, loopyWriter will exit, resulting in conn closure
}

func (*goAway) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:198
	_go_fuzz_dep_.CoverTab[76612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:198
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:198
	// _ = "end of CoverTab[76612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:198
}

type ping struct {
	ack	bool
	data	[8]byte
}

func (*ping) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:205
	_go_fuzz_dep_.CoverTab[76613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:205
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:205
	// _ = "end of CoverTab[76613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:205
}

type outFlowControlSizeRequest struct {
	resp chan uint32
}

func (*outFlowControlSizeRequest) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:211
	_go_fuzz_dep_.CoverTab[76614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:211
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:211
	// _ = "end of CoverTab[76614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:211
}

// closeConnection is an instruction to tell the loopy writer to flush the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:213
// framer and exit, which will cause the transport's connection to be closed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:213
// (by the client or server).  The transport itself will close after the reader
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:213
// encounters the EOF caused by the connection closure.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:217
type closeConnection struct{}

func (closeConnection) isTransportResponseFrame() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:219
	_go_fuzz_dep_.CoverTab[76615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:219
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:219
	// _ = "end of CoverTab[76615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:219
}

type outStreamState int

const (
	active	outStreamState	= iota
	empty
	waitingOnStreamQuota
)

type outStream struct {
	id			uint32
	state			outStreamState
	itl			*itemList
	bytesOutStanding	int
	wq			*writeQuota

	next	*outStream
	prev	*outStream
}

func (s *outStream) deleteSelf() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:240
	_go_fuzz_dep_.CoverTab[76616]++
													if s.prev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:241
		_go_fuzz_dep_.CoverTab[76619]++
														s.prev.next = s.next
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:242
		// _ = "end of CoverTab[76619]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:243
		_go_fuzz_dep_.CoverTab[76620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:243
		// _ = "end of CoverTab[76620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:243
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:243
	// _ = "end of CoverTab[76616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:243
	_go_fuzz_dep_.CoverTab[76617]++
													if s.next != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:244
		_go_fuzz_dep_.CoverTab[76621]++
														s.next.prev = s.prev
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:245
		// _ = "end of CoverTab[76621]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:246
		_go_fuzz_dep_.CoverTab[76622]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:246
		// _ = "end of CoverTab[76622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:246
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:246
	// _ = "end of CoverTab[76617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:246
	_go_fuzz_dep_.CoverTab[76618]++
													s.next, s.prev = nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:247
	// _ = "end of CoverTab[76618]"
}

type outStreamList struct {
	// Following are sentinel objects that mark the
	// beginning and end of the list. They do not
	// contain any item lists. All valid objects are
	// inserted in between them.
	// This is needed so that an outStream object can
	// deleteSelf() in O(1) time without knowing which
	// list it belongs to.
	head	*outStream
	tail	*outStream
}

func newOutStreamList() *outStreamList {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:262
	_go_fuzz_dep_.CoverTab[76623]++
													head, tail := new(outStream), new(outStream)
													head.next = tail
													tail.prev = head
													return &outStreamList{
		head:	head,
		tail:	tail,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:269
	// _ = "end of CoverTab[76623]"
}

func (l *outStreamList) enqueue(s *outStream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:272
	_go_fuzz_dep_.CoverTab[76624]++
													e := l.tail.prev
													e.next = s
													s.prev = e
													s.next = l.tail
													l.tail.prev = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:277
	// _ = "end of CoverTab[76624]"
}

// remove from the beginning of the list.
func (l *outStreamList) dequeue() *outStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:281
	_go_fuzz_dep_.CoverTab[76625]++
													b := l.head.next
													if b == l.tail {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:283
		_go_fuzz_dep_.CoverTab[76627]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:284
		// _ = "end of CoverTab[76627]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:285
		_go_fuzz_dep_.CoverTab[76628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:285
		// _ = "end of CoverTab[76628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:285
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:285
	// _ = "end of CoverTab[76625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:285
	_go_fuzz_dep_.CoverTab[76626]++
													b.deleteSelf()
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:287
	// _ = "end of CoverTab[76626]"
}

// controlBuffer is a way to pass information to loopy.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:290
// Information is passed as specific struct types called control frames.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:290
// A control frame not only represents data, messages or headers to be sent out
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:290
// but can also be used to instruct loopy to update its internal state.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:290
// It shouldn't be confused with an HTTP2 frame, although some of the control frames
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:290
// like dataFrame and headerFrame do go out on wire as HTTP2 frames.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:296
type controlBuffer struct {
	ch		chan struct{}
	done		<-chan struct{}
	mu		sync.Mutex
	consumerWaiting	bool
	list		*itemList
	err		error

	// transportResponseFrames counts the number of queued items that represent
	// the response of an action initiated by the peer.  trfChan is created
	// when transportResponseFrames >= maxQueuedTransportResponseFrames and is
	// closed and nilled when transportResponseFrames drops below the
	// threshold.  Both fields are protected by mu.
	transportResponseFrames	int
	trfChan			atomic.Value	// chan struct{}
}

func newControlBuffer(done <-chan struct{}) *controlBuffer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:313
	_go_fuzz_dep_.CoverTab[76629]++
													return &controlBuffer{
		ch:	make(chan struct{}, 1),
		list:	&itemList{},
		done:	done,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:318
	// _ = "end of CoverTab[76629]"
}

// throttle blocks if there are too many incomingSettings/cleanupStreams in the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:321
// controlbuf.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:323
func (c *controlBuffer) throttle() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:323
	_go_fuzz_dep_.CoverTab[76630]++
													ch, _ := c.trfChan.Load().(chan struct{})
													if ch != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:325
		_go_fuzz_dep_.CoverTab[76631]++
														select {
		case <-ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:327
			_go_fuzz_dep_.CoverTab[76632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:327
			// _ = "end of CoverTab[76632]"
		case <-c.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:328
			_go_fuzz_dep_.CoverTab[76633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:328
			// _ = "end of CoverTab[76633]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:329
		// _ = "end of CoverTab[76631]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:330
		_go_fuzz_dep_.CoverTab[76634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:330
		// _ = "end of CoverTab[76634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:330
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:330
	// _ = "end of CoverTab[76630]"
}

func (c *controlBuffer) put(it cbItem) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:333
	_go_fuzz_dep_.CoverTab[76635]++
													_, err := c.executeAndPut(nil, it)
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:335
	// _ = "end of CoverTab[76635]"
}

func (c *controlBuffer) executeAndPut(f func(it interface{}) bool, it cbItem) (bool, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:338
	_go_fuzz_dep_.CoverTab[76636]++
													var wakeUp bool
													c.mu.Lock()
													if c.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:341
		_go_fuzz_dep_.CoverTab[76642]++
														c.mu.Unlock()
														return false, c.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:343
		// _ = "end of CoverTab[76642]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:344
		_go_fuzz_dep_.CoverTab[76643]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:344
		// _ = "end of CoverTab[76643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:344
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:344
	// _ = "end of CoverTab[76636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:344
	_go_fuzz_dep_.CoverTab[76637]++
													if f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:345
		_go_fuzz_dep_.CoverTab[76644]++
														if !f(it) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:346
			_go_fuzz_dep_.CoverTab[76645]++
															c.mu.Unlock()
															return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:348
			// _ = "end of CoverTab[76645]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:349
			_go_fuzz_dep_.CoverTab[76646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:349
			// _ = "end of CoverTab[76646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:349
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:349
		// _ = "end of CoverTab[76644]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:350
		_go_fuzz_dep_.CoverTab[76647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:350
		// _ = "end of CoverTab[76647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:350
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:350
	// _ = "end of CoverTab[76637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:350
	_go_fuzz_dep_.CoverTab[76638]++
													if c.consumerWaiting {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:351
		_go_fuzz_dep_.CoverTab[76648]++
														wakeUp = true
														c.consumerWaiting = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:353
		// _ = "end of CoverTab[76648]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:354
		_go_fuzz_dep_.CoverTab[76649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:354
		// _ = "end of CoverTab[76649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:354
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:354
	// _ = "end of CoverTab[76638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:354
	_go_fuzz_dep_.CoverTab[76639]++
													c.list.enqueue(it)
													if it.isTransportResponseFrame() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:356
		_go_fuzz_dep_.CoverTab[76650]++
														c.transportResponseFrames++
														if c.transportResponseFrames == maxQueuedTransportResponseFrames {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:358
			_go_fuzz_dep_.CoverTab[76651]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:361
			c.trfChan.Store(make(chan struct{}))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:361
			// _ = "end of CoverTab[76651]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:362
			_go_fuzz_dep_.CoverTab[76652]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:362
			// _ = "end of CoverTab[76652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:362
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:362
		// _ = "end of CoverTab[76650]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:363
		_go_fuzz_dep_.CoverTab[76653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:363
		// _ = "end of CoverTab[76653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:363
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:363
	// _ = "end of CoverTab[76639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:363
	_go_fuzz_dep_.CoverTab[76640]++
													c.mu.Unlock()
													if wakeUp {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:365
		_go_fuzz_dep_.CoverTab[76654]++
														select {
		case c.ch <- struct{}{}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:367
			_go_fuzz_dep_.CoverTab[76655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:367
			// _ = "end of CoverTab[76655]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:368
			_go_fuzz_dep_.CoverTab[76656]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:368
			// _ = "end of CoverTab[76656]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:369
		// _ = "end of CoverTab[76654]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:370
		_go_fuzz_dep_.CoverTab[76657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:370
		// _ = "end of CoverTab[76657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:370
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:370
	// _ = "end of CoverTab[76640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:370
	_go_fuzz_dep_.CoverTab[76641]++
													return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:371
	// _ = "end of CoverTab[76641]"
}

// Note argument f should never be nil.
func (c *controlBuffer) execute(f func(it interface{}) bool, it interface{}) (bool, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:375
	_go_fuzz_dep_.CoverTab[76658]++
													c.mu.Lock()
													if c.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:377
		_go_fuzz_dep_.CoverTab[76661]++
														c.mu.Unlock()
														return false, c.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:379
		// _ = "end of CoverTab[76661]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:380
		_go_fuzz_dep_.CoverTab[76662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:380
		// _ = "end of CoverTab[76662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:380
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:380
	// _ = "end of CoverTab[76658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:380
	_go_fuzz_dep_.CoverTab[76659]++
													if !f(it) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:381
		_go_fuzz_dep_.CoverTab[76663]++
														c.mu.Unlock()
														return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:383
		// _ = "end of CoverTab[76663]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:384
		_go_fuzz_dep_.CoverTab[76664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:384
		// _ = "end of CoverTab[76664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:384
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:384
	// _ = "end of CoverTab[76659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:384
	_go_fuzz_dep_.CoverTab[76660]++
													c.mu.Unlock()
													return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:386
	// _ = "end of CoverTab[76660]"
}

func (c *controlBuffer) get(block bool) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:389
	_go_fuzz_dep_.CoverTab[76665]++
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:390
		_go_fuzz_dep_.CoverTab[76666]++
														c.mu.Lock()
														if c.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:392
			_go_fuzz_dep_.CoverTab[76670]++
															c.mu.Unlock()
															return nil, c.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:394
			// _ = "end of CoverTab[76670]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:395
			_go_fuzz_dep_.CoverTab[76671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:395
			// _ = "end of CoverTab[76671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:395
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:395
		// _ = "end of CoverTab[76666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:395
		_go_fuzz_dep_.CoverTab[76667]++
														if !c.list.isEmpty() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:396
			_go_fuzz_dep_.CoverTab[76672]++
															h := c.list.dequeue().(cbItem)
															if h.isTransportResponseFrame() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:398
				_go_fuzz_dep_.CoverTab[76674]++
																if c.transportResponseFrames == maxQueuedTransportResponseFrames {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:399
					_go_fuzz_dep_.CoverTab[76676]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:402
					ch := c.trfChan.Load().(chan struct{})
																	close(ch)
																	c.trfChan.Store((chan struct{})(nil))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:404
					// _ = "end of CoverTab[76676]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:405
					_go_fuzz_dep_.CoverTab[76677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:405
					// _ = "end of CoverTab[76677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:405
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:405
				// _ = "end of CoverTab[76674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:405
				_go_fuzz_dep_.CoverTab[76675]++
																c.transportResponseFrames--
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:406
				// _ = "end of CoverTab[76675]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:407
				_go_fuzz_dep_.CoverTab[76678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:407
				// _ = "end of CoverTab[76678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:407
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:407
			// _ = "end of CoverTab[76672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:407
			_go_fuzz_dep_.CoverTab[76673]++
															c.mu.Unlock()
															return h, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:409
			// _ = "end of CoverTab[76673]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:410
			_go_fuzz_dep_.CoverTab[76679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:410
			// _ = "end of CoverTab[76679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:410
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:410
		// _ = "end of CoverTab[76667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:410
		_go_fuzz_dep_.CoverTab[76668]++
														if !block {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:411
			_go_fuzz_dep_.CoverTab[76680]++
															c.mu.Unlock()
															return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:413
			// _ = "end of CoverTab[76680]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:414
			_go_fuzz_dep_.CoverTab[76681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:414
			// _ = "end of CoverTab[76681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:414
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:414
		// _ = "end of CoverTab[76668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:414
		_go_fuzz_dep_.CoverTab[76669]++
														c.consumerWaiting = true
														c.mu.Unlock()
														select {
		case <-c.ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:418
			_go_fuzz_dep_.CoverTab[76682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:418
			// _ = "end of CoverTab[76682]"
		case <-c.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:419
			_go_fuzz_dep_.CoverTab[76683]++
															return nil, errors.New("transport closed by client")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:420
			// _ = "end of CoverTab[76683]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:421
		// _ = "end of CoverTab[76669]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:422
	// _ = "end of CoverTab[76665]"
}

func (c *controlBuffer) finish() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:425
	_go_fuzz_dep_.CoverTab[76684]++
													c.mu.Lock()
													if c.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:427
		_go_fuzz_dep_.CoverTab[76688]++
														c.mu.Unlock()
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:429
		// _ = "end of CoverTab[76688]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:430
		_go_fuzz_dep_.CoverTab[76689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:430
		// _ = "end of CoverTab[76689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:430
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:430
	// _ = "end of CoverTab[76684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:430
	_go_fuzz_dep_.CoverTab[76685]++
													c.err = ErrConnClosing

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:435
	for head := c.list.dequeueAll(); head != nil; head = head.next {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:435
		_go_fuzz_dep_.CoverTab[76690]++
														hdr, ok := head.it.(*headerFrame)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:437
			_go_fuzz_dep_.CoverTab[76692]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:438
			// _ = "end of CoverTab[76692]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:439
			_go_fuzz_dep_.CoverTab[76693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:439
			// _ = "end of CoverTab[76693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:439
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:439
		// _ = "end of CoverTab[76690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:439
		_go_fuzz_dep_.CoverTab[76691]++
														if hdr.onOrphaned != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:440
			_go_fuzz_dep_.CoverTab[76694]++
															hdr.onOrphaned(ErrConnClosing)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:441
			// _ = "end of CoverTab[76694]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:442
			_go_fuzz_dep_.CoverTab[76695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:442
			// _ = "end of CoverTab[76695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:442
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:442
		// _ = "end of CoverTab[76691]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:443
	// _ = "end of CoverTab[76685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:443
	_go_fuzz_dep_.CoverTab[76686]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:447
	ch, _ := c.trfChan.Load().(chan struct{})
	if ch != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:448
		_go_fuzz_dep_.CoverTab[76696]++
														close(ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:449
		// _ = "end of CoverTab[76696]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:450
		_go_fuzz_dep_.CoverTab[76697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:450
		// _ = "end of CoverTab[76697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:450
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:450
	// _ = "end of CoverTab[76686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:450
	_go_fuzz_dep_.CoverTab[76687]++
													c.trfChan.Store((chan struct{})(nil))
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:452
	// _ = "end of CoverTab[76687]"
}

type side int

const (
	clientSide	side	= iota
	serverSide
)

// Loopy receives frames from the control buffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// Each frame is handled individually; most of the work done by loopy goes
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// into handling data frames. Loopy maintains a queue of active streams, and each
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// stream maintains a queue of data frames; as loopy receives data frames
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// it gets added to the queue of the relevant stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// Loopy goes over this list of active streams by processing one node every iteration,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// thereby closely resemebling to a round-robin scheduling over all streams. While
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// processing a stream, loopy writes out data bytes from this stream capped by the min
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:462
// of http2MaxFrameLen, connection-level flow control and stream-level flow control.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:471
type loopyWriter struct {
	side		side
	cbuf		*controlBuffer
	sendQuota	uint32
	oiws		uint32	// outbound initial window size.
	// estdStreams is map of all established streams that are not cleaned-up yet.
	// On client-side, this is all streams whose headers were sent out.
	// On server-side, this is all streams whose headers were received.
	estdStreams	map[uint32]*outStream	// Established streams.
	// activeStreams is a linked-list of all streams that have data to send and some
	// stream-level flow control quota.
	// Each of these streams internally have a list of data items(and perhaps trailers
	// on the server-side) to be sent out.
	activeStreams	*outStreamList
	framer		*framer
	hBuf		*bytes.Buffer	// The buffer for HPACK encoding.
	hEnc		*hpack.Encoder	// HPACK encoder.
	bdpEst		*bdpEstimator
	draining	bool
	conn		net.Conn

	// Side-specific handlers
	ssGoAwayHandler	func(*goAway) (bool, error)
}

func newLoopyWriter(s side, fr *framer, cbuf *controlBuffer, bdpEst *bdpEstimator, conn net.Conn) *loopyWriter {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:496
	_go_fuzz_dep_.CoverTab[76698]++
													var buf bytes.Buffer
													l := &loopyWriter{
		side:		s,
		cbuf:		cbuf,
		sendQuota:	defaultWindowSize,
		oiws:		defaultWindowSize,
		estdStreams:	make(map[uint32]*outStream),
		activeStreams:	newOutStreamList(),
		framer:		fr,
		hBuf:		&buf,
		hEnc:		hpack.NewEncoder(&buf),
		bdpEst:		bdpEst,
		conn:		conn,
	}
													return l
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:511
	// _ = "end of CoverTab[76698]"
}

const minBatchSize = 1000

// run should be run in a separate goroutine.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// It reads control frames from controlBuf and processes them by:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// 1. Updating loopy's internal state, or/and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// 2. Writing out HTTP2 frames on the wire.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// Loopy keeps all active streams with data to send in a linked-list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// All streams in the activeStreams linked-list must have both:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// 1. Data to send, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// 2. Stream level flow control quota available.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// In each iteration of run loop, other than processing the incoming control
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// frame, loopy calls processData, which processes one node from the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// activeStreams linked-list.  This results in writing of HTTP2 frames into an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// underlying write buffer.  When there's no more control frames to read from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// controlBuf, loopy flushes the write buffer.  As an optimization, to increase
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// the batch size for each flush, loopy yields the processor, once if the batch
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// size is too low to give stream goroutines a chance to fill it up.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// Upon exiting, if the error causing the exit is not an I/O error, run()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// flushes and closes the underlying connection.  Otherwise, the connection is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:516
// left open to allow the I/O error to be encountered by the reader instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:537
func (l *loopyWriter) run() (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:537
	_go_fuzz_dep_.CoverTab[76699]++
													defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:538
		_go_fuzz_dep_.CoverTab[76701]++
														if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:539
			_go_fuzz_dep_.CoverTab[76704]++
															logger.Infof("transport: loopyWriter exiting with error: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:540
			// _ = "end of CoverTab[76704]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:541
			_go_fuzz_dep_.CoverTab[76705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:541
			// _ = "end of CoverTab[76705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:541
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:541
		// _ = "end of CoverTab[76701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:541
		_go_fuzz_dep_.CoverTab[76702]++
														if !isIOError(err) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:542
			_go_fuzz_dep_.CoverTab[76706]++
															l.framer.writer.Flush()
															l.conn.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:544
			// _ = "end of CoverTab[76706]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:545
			_go_fuzz_dep_.CoverTab[76707]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:545
			// _ = "end of CoverTab[76707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:545
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:545
		// _ = "end of CoverTab[76702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:545
		_go_fuzz_dep_.CoverTab[76703]++
														l.cbuf.finish()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:546
		// _ = "end of CoverTab[76703]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:547
	// _ = "end of CoverTab[76699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:547
	_go_fuzz_dep_.CoverTab[76700]++
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:548
		_go_fuzz_dep_.CoverTab[76708]++
														it, err := l.cbuf.get(true)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:550
			_go_fuzz_dep_.CoverTab[76712]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:551
			// _ = "end of CoverTab[76712]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:552
			_go_fuzz_dep_.CoverTab[76713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:552
			// _ = "end of CoverTab[76713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:552
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:552
		// _ = "end of CoverTab[76708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:552
		_go_fuzz_dep_.CoverTab[76709]++
														if err = l.handle(it); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:553
			_go_fuzz_dep_.CoverTab[76714]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:554
			// _ = "end of CoverTab[76714]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:555
			_go_fuzz_dep_.CoverTab[76715]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:555
			// _ = "end of CoverTab[76715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:555
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:555
		// _ = "end of CoverTab[76709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:555
		_go_fuzz_dep_.CoverTab[76710]++
														if _, err = l.processData(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:556
			_go_fuzz_dep_.CoverTab[76716]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:557
			// _ = "end of CoverTab[76716]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:558
			_go_fuzz_dep_.CoverTab[76717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:558
			// _ = "end of CoverTab[76717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:558
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:558
		// _ = "end of CoverTab[76710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:558
		_go_fuzz_dep_.CoverTab[76711]++
														gosched := true
	hasdata:
		for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:561
			_go_fuzz_dep_.CoverTab[76718]++
															it, err := l.cbuf.get(false)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:563
				_go_fuzz_dep_.CoverTab[76724]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:564
				// _ = "end of CoverTab[76724]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:565
				_go_fuzz_dep_.CoverTab[76725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:565
				// _ = "end of CoverTab[76725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:565
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:565
			// _ = "end of CoverTab[76718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:565
			_go_fuzz_dep_.CoverTab[76719]++
															if it != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:566
				_go_fuzz_dep_.CoverTab[76726]++
																if err = l.handle(it); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:567
					_go_fuzz_dep_.CoverTab[76729]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:568
					// _ = "end of CoverTab[76729]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:569
					_go_fuzz_dep_.CoverTab[76730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:569
					// _ = "end of CoverTab[76730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:569
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:569
				// _ = "end of CoverTab[76726]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:569
				_go_fuzz_dep_.CoverTab[76727]++
																if _, err = l.processData(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:570
					_go_fuzz_dep_.CoverTab[76731]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:571
					// _ = "end of CoverTab[76731]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:572
					_go_fuzz_dep_.CoverTab[76732]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:572
					// _ = "end of CoverTab[76732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:572
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:572
				// _ = "end of CoverTab[76727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:572
				_go_fuzz_dep_.CoverTab[76728]++
																continue hasdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:573
				// _ = "end of CoverTab[76728]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:574
				_go_fuzz_dep_.CoverTab[76733]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:574
				// _ = "end of CoverTab[76733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:574
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:574
			// _ = "end of CoverTab[76719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:574
			_go_fuzz_dep_.CoverTab[76720]++
															isEmpty, err := l.processData()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:576
				_go_fuzz_dep_.CoverTab[76734]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:577
				// _ = "end of CoverTab[76734]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:578
				_go_fuzz_dep_.CoverTab[76735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:578
				// _ = "end of CoverTab[76735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:578
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:578
			// _ = "end of CoverTab[76720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:578
			_go_fuzz_dep_.CoverTab[76721]++
															if !isEmpty {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:579
				_go_fuzz_dep_.CoverTab[76736]++
																continue hasdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:580
				// _ = "end of CoverTab[76736]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:581
				_go_fuzz_dep_.CoverTab[76737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:581
				// _ = "end of CoverTab[76737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:581
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:581
			// _ = "end of CoverTab[76721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:581
			_go_fuzz_dep_.CoverTab[76722]++
															if gosched {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:582
				_go_fuzz_dep_.CoverTab[76738]++
																gosched = false
																if l.framer.writer.offset < minBatchSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:584
					_go_fuzz_dep_.CoverTab[76739]++
																	runtime.Gosched()
																	continue hasdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:586
					// _ = "end of CoverTab[76739]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:587
					_go_fuzz_dep_.CoverTab[76740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:587
					// _ = "end of CoverTab[76740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:587
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:587
				// _ = "end of CoverTab[76738]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:588
				_go_fuzz_dep_.CoverTab[76741]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:588
				// _ = "end of CoverTab[76741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:588
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:588
			// _ = "end of CoverTab[76722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:588
			_go_fuzz_dep_.CoverTab[76723]++
															l.framer.writer.Flush()
															break hasdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:590
			// _ = "end of CoverTab[76723]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:591
		// _ = "end of CoverTab[76711]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:592
	// _ = "end of CoverTab[76700]"
}

func (l *loopyWriter) outgoingWindowUpdateHandler(w *outgoingWindowUpdate) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:595
	_go_fuzz_dep_.CoverTab[76742]++
													return l.framer.fr.WriteWindowUpdate(w.streamID, w.increment)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:596
	// _ = "end of CoverTab[76742]"
}

func (l *loopyWriter) incomingWindowUpdateHandler(w *incomingWindowUpdate) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:599
	_go_fuzz_dep_.CoverTab[76743]++

													if w.streamID == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:601
		_go_fuzz_dep_.CoverTab[76745]++
														l.sendQuota += w.increment
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:603
		// _ = "end of CoverTab[76745]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:604
		_go_fuzz_dep_.CoverTab[76746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:604
		// _ = "end of CoverTab[76746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:604
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:604
	// _ = "end of CoverTab[76743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:604
	_go_fuzz_dep_.CoverTab[76744]++

													if str, ok := l.estdStreams[w.streamID]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:606
		_go_fuzz_dep_.CoverTab[76747]++
														str.bytesOutStanding -= int(w.increment)
														if strQuota := int(l.oiws) - str.bytesOutStanding; strQuota > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:608
			_go_fuzz_dep_.CoverTab[76748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:608
			return str.state == waitingOnStreamQuota
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:608
			// _ = "end of CoverTab[76748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:608
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:608
			_go_fuzz_dep_.CoverTab[76749]++
															str.state = active
															l.activeStreams.enqueue(str)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:611
			// _ = "end of CoverTab[76749]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:612
			_go_fuzz_dep_.CoverTab[76750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:612
			// _ = "end of CoverTab[76750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:612
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:612
		// _ = "end of CoverTab[76747]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:613
		_go_fuzz_dep_.CoverTab[76751]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:613
		// _ = "end of CoverTab[76751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:613
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:613
	// _ = "end of CoverTab[76744]"
}

func (l *loopyWriter) outgoingSettingsHandler(s *outgoingSettings) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:616
	_go_fuzz_dep_.CoverTab[76752]++
													return l.framer.fr.WriteSettings(s.ss...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:617
	// _ = "end of CoverTab[76752]"
}

func (l *loopyWriter) incomingSettingsHandler(s *incomingSettings) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:620
	_go_fuzz_dep_.CoverTab[76753]++
													l.applySettings(s.ss)
													return l.framer.fr.WriteSettingsAck()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:622
	// _ = "end of CoverTab[76753]"
}

func (l *loopyWriter) registerStreamHandler(h *registerStream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:625
	_go_fuzz_dep_.CoverTab[76754]++
													str := &outStream{
		id:	h.streamID,
		state:	empty,
		itl:	&itemList{},
		wq:	h.wq,
	}
													l.estdStreams[h.streamID] = str
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:632
	// _ = "end of CoverTab[76754]"
}

func (l *loopyWriter) headerHandler(h *headerFrame) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:635
	_go_fuzz_dep_.CoverTab[76755]++
													if l.side == serverSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:636
		_go_fuzz_dep_.CoverTab[76757]++
														str, ok := l.estdStreams[h.streamID]
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:638
			_go_fuzz_dep_.CoverTab[76762]++
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:639
				_go_fuzz_dep_.CoverTab[76764]++
																logger.Warningf("transport: loopy doesn't recognize the stream: %d", h.streamID)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:640
				// _ = "end of CoverTab[76764]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:641
				_go_fuzz_dep_.CoverTab[76765]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:641
				// _ = "end of CoverTab[76765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:641
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:641
			// _ = "end of CoverTab[76762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:641
			_go_fuzz_dep_.CoverTab[76763]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:642
			// _ = "end of CoverTab[76763]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:643
			_go_fuzz_dep_.CoverTab[76766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:643
			// _ = "end of CoverTab[76766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:643
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:643
		// _ = "end of CoverTab[76757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:643
		_go_fuzz_dep_.CoverTab[76758]++

														if !h.endStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:645
			_go_fuzz_dep_.CoverTab[76767]++
															return l.writeHeader(h.streamID, h.endStream, h.hf, h.onWrite)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:646
			// _ = "end of CoverTab[76767]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:647
			_go_fuzz_dep_.CoverTab[76768]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:647
			// _ = "end of CoverTab[76768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:647
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:647
		// _ = "end of CoverTab[76758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:647
		_go_fuzz_dep_.CoverTab[76759]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:650
		if str.state != empty {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:650
			_go_fuzz_dep_.CoverTab[76769]++

															str.itl.enqueue(h)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:653
			// _ = "end of CoverTab[76769]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:654
			_go_fuzz_dep_.CoverTab[76770]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:654
			// _ = "end of CoverTab[76770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:654
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:654
		// _ = "end of CoverTab[76759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:654
		_go_fuzz_dep_.CoverTab[76760]++
														if err := l.writeHeader(h.streamID, h.endStream, h.hf, h.onWrite); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:655
			_go_fuzz_dep_.CoverTab[76771]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:656
			// _ = "end of CoverTab[76771]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:657
			_go_fuzz_dep_.CoverTab[76772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:657
			// _ = "end of CoverTab[76772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:657
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:657
		// _ = "end of CoverTab[76760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:657
		_go_fuzz_dep_.CoverTab[76761]++
														return l.cleanupStreamHandler(h.cleanup)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:658
		// _ = "end of CoverTab[76761]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:659
		_go_fuzz_dep_.CoverTab[76773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:659
		// _ = "end of CoverTab[76773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:659
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:659
	// _ = "end of CoverTab[76755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:659
	_go_fuzz_dep_.CoverTab[76756]++

													str := &outStream{
		id:	h.streamID,
		state:	empty,
		itl:	&itemList{},
		wq:	h.wq,
	}
													return l.originateStream(str, h)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:667
	// _ = "end of CoverTab[76756]"
}

func (l *loopyWriter) originateStream(str *outStream, hdr *headerFrame) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:670
	_go_fuzz_dep_.CoverTab[76774]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:673
	if l.draining {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:673
		_go_fuzz_dep_.CoverTab[76778]++

														hdr.onOrphaned(errStreamDrain)
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:676
		// _ = "end of CoverTab[76778]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:677
		_go_fuzz_dep_.CoverTab[76779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:677
		// _ = "end of CoverTab[76779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:677
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:677
	// _ = "end of CoverTab[76774]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:677
	_go_fuzz_dep_.CoverTab[76775]++
													if err := hdr.initStream(str.id); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:678
		_go_fuzz_dep_.CoverTab[76780]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:679
		// _ = "end of CoverTab[76780]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:680
		_go_fuzz_dep_.CoverTab[76781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:680
		// _ = "end of CoverTab[76781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:680
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:680
	// _ = "end of CoverTab[76775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:680
	_go_fuzz_dep_.CoverTab[76776]++
													if err := l.writeHeader(str.id, hdr.endStream, hdr.hf, hdr.onWrite); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:681
		_go_fuzz_dep_.CoverTab[76782]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:682
		// _ = "end of CoverTab[76782]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:683
		_go_fuzz_dep_.CoverTab[76783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:683
		// _ = "end of CoverTab[76783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:683
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:683
	// _ = "end of CoverTab[76776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:683
	_go_fuzz_dep_.CoverTab[76777]++
													l.estdStreams[str.id] = str
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:685
	// _ = "end of CoverTab[76777]"
}

func (l *loopyWriter) writeHeader(streamID uint32, endStream bool, hf []hpack.HeaderField, onWrite func()) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:688
	_go_fuzz_dep_.CoverTab[76784]++
													if onWrite != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:689
		_go_fuzz_dep_.CoverTab[76788]++
														onWrite()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:690
		// _ = "end of CoverTab[76788]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:691
		_go_fuzz_dep_.CoverTab[76789]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:691
		// _ = "end of CoverTab[76789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:691
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:691
	// _ = "end of CoverTab[76784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:691
	_go_fuzz_dep_.CoverTab[76785]++
													l.hBuf.Reset()
													for _, f := range hf {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:693
		_go_fuzz_dep_.CoverTab[76790]++
														if err := l.hEnc.WriteField(f); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:694
			_go_fuzz_dep_.CoverTab[76791]++
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:695
				_go_fuzz_dep_.CoverTab[76792]++
																logger.Warningf("transport: loopyWriter.writeHeader encountered error while encoding headers: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:696
				// _ = "end of CoverTab[76792]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:697
				_go_fuzz_dep_.CoverTab[76793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:697
				// _ = "end of CoverTab[76793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:697
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:697
			// _ = "end of CoverTab[76791]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:698
			_go_fuzz_dep_.CoverTab[76794]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:698
			// _ = "end of CoverTab[76794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:698
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:698
		// _ = "end of CoverTab[76790]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:699
	// _ = "end of CoverTab[76785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:699
	_go_fuzz_dep_.CoverTab[76786]++
													var (
		err			error
		endHeaders, first	bool
	)
	first = true
	for !endHeaders {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:705
		_go_fuzz_dep_.CoverTab[76795]++
														size := l.hBuf.Len()
														if size > http2MaxFrameLen {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:707
			_go_fuzz_dep_.CoverTab[76798]++
															size = http2MaxFrameLen
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:708
			// _ = "end of CoverTab[76798]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:709
			_go_fuzz_dep_.CoverTab[76799]++
															endHeaders = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:710
			// _ = "end of CoverTab[76799]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:711
		// _ = "end of CoverTab[76795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:711
		_go_fuzz_dep_.CoverTab[76796]++
														if first {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:712
			_go_fuzz_dep_.CoverTab[76800]++
															first = false
															err = l.framer.fr.WriteHeaders(http2.HeadersFrameParam{
				StreamID:	streamID,
				BlockFragment:	l.hBuf.Next(size),
				EndStream:	endStream,
				EndHeaders:	endHeaders,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:719
			// _ = "end of CoverTab[76800]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:720
			_go_fuzz_dep_.CoverTab[76801]++
															err = l.framer.fr.WriteContinuation(
				streamID,
				endHeaders,
				l.hBuf.Next(size),
			)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:725
			// _ = "end of CoverTab[76801]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:726
		// _ = "end of CoverTab[76796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:726
		_go_fuzz_dep_.CoverTab[76797]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:727
			_go_fuzz_dep_.CoverTab[76802]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:728
			// _ = "end of CoverTab[76802]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:729
			_go_fuzz_dep_.CoverTab[76803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:729
			// _ = "end of CoverTab[76803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:729
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:729
		// _ = "end of CoverTab[76797]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:730
	// _ = "end of CoverTab[76786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:730
	_go_fuzz_dep_.CoverTab[76787]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:731
	// _ = "end of CoverTab[76787]"
}

func (l *loopyWriter) preprocessData(df *dataFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:734
	_go_fuzz_dep_.CoverTab[76804]++
													str, ok := l.estdStreams[df.streamID]
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:736
		_go_fuzz_dep_.CoverTab[76806]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:737
		// _ = "end of CoverTab[76806]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:738
		_go_fuzz_dep_.CoverTab[76807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:738
		// _ = "end of CoverTab[76807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:738
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:738
	// _ = "end of CoverTab[76804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:738
	_go_fuzz_dep_.CoverTab[76805]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:741
	str.itl.enqueue(df)
	if str.state == empty {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:742
		_go_fuzz_dep_.CoverTab[76808]++
														str.state = active
														l.activeStreams.enqueue(str)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:744
		// _ = "end of CoverTab[76808]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:745
		_go_fuzz_dep_.CoverTab[76809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:745
		// _ = "end of CoverTab[76809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:745
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:745
	// _ = "end of CoverTab[76805]"
}

func (l *loopyWriter) pingHandler(p *ping) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:748
	_go_fuzz_dep_.CoverTab[76810]++
													if !p.ack {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:749
		_go_fuzz_dep_.CoverTab[76812]++
														l.bdpEst.timesnap(p.data)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:750
		// _ = "end of CoverTab[76812]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:751
		_go_fuzz_dep_.CoverTab[76813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:751
		// _ = "end of CoverTab[76813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:751
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:751
	// _ = "end of CoverTab[76810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:751
	_go_fuzz_dep_.CoverTab[76811]++
													return l.framer.fr.WritePing(p.ack, p.data)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:752
	// _ = "end of CoverTab[76811]"

}

func (l *loopyWriter) outFlowControlSizeRequestHandler(o *outFlowControlSizeRequest) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:756
	_go_fuzz_dep_.CoverTab[76814]++
													o.resp <- l.sendQuota
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:757
	// _ = "end of CoverTab[76814]"
}

func (l *loopyWriter) cleanupStreamHandler(c *cleanupStream) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:760
	_go_fuzz_dep_.CoverTab[76815]++
													c.onWrite()
													if str, ok := l.estdStreams[c.streamID]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:762
		_go_fuzz_dep_.CoverTab[76819]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:766
		delete(l.estdStreams, c.streamID)
														str.deleteSelf()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:767
		// _ = "end of CoverTab[76819]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:768
		_go_fuzz_dep_.CoverTab[76820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:768
		// _ = "end of CoverTab[76820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:768
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:768
	// _ = "end of CoverTab[76815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:768
	_go_fuzz_dep_.CoverTab[76816]++
													if c.rst {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:769
		_go_fuzz_dep_.CoverTab[76821]++
														if err := l.framer.fr.WriteRSTStream(c.streamID, c.rstCode); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:770
			_go_fuzz_dep_.CoverTab[76822]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:771
			// _ = "end of CoverTab[76822]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:772
			_go_fuzz_dep_.CoverTab[76823]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:772
			// _ = "end of CoverTab[76823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:772
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:772
		// _ = "end of CoverTab[76821]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:773
		_go_fuzz_dep_.CoverTab[76824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:773
		// _ = "end of CoverTab[76824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:773
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:773
	// _ = "end of CoverTab[76816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:773
	_go_fuzz_dep_.CoverTab[76817]++
													if l.draining && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:774
		_go_fuzz_dep_.CoverTab[76825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:774
		return len(l.estdStreams) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:774
		// _ = "end of CoverTab[76825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:774
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:774
		_go_fuzz_dep_.CoverTab[76826]++

														return errors.New("finished processing active streams while in draining mode")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:776
		// _ = "end of CoverTab[76826]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:777
		_go_fuzz_dep_.CoverTab[76827]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:777
		// _ = "end of CoverTab[76827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:777
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:777
	// _ = "end of CoverTab[76817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:777
	_go_fuzz_dep_.CoverTab[76818]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:778
	// _ = "end of CoverTab[76818]"
}

func (l *loopyWriter) earlyAbortStreamHandler(eas *earlyAbortStream) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:781
	_go_fuzz_dep_.CoverTab[76828]++
													if l.side == clientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:782
		_go_fuzz_dep_.CoverTab[76833]++
														return errors.New("earlyAbortStream not handled on client")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:783
		// _ = "end of CoverTab[76833]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:784
		_go_fuzz_dep_.CoverTab[76834]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:784
		// _ = "end of CoverTab[76834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:784
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:784
	// _ = "end of CoverTab[76828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:784
	_go_fuzz_dep_.CoverTab[76829]++

													if eas.httpStatus == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:786
		_go_fuzz_dep_.CoverTab[76835]++
														eas.httpStatus = 200
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:787
		// _ = "end of CoverTab[76835]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:788
		_go_fuzz_dep_.CoverTab[76836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:788
		// _ = "end of CoverTab[76836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:788
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:788
	// _ = "end of CoverTab[76829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:788
	_go_fuzz_dep_.CoverTab[76830]++
													headerFields := []hpack.HeaderField{
		{Name: ":status", Value: strconv.Itoa(int(eas.httpStatus))},
		{Name: "content-type", Value: grpcutil.ContentType(eas.contentSubtype)},
		{Name: "grpc-status", Value: strconv.Itoa(int(eas.status.Code()))},
		{Name: "grpc-message", Value: encodeGrpcMessage(eas.status.Message())},
	}

	if err := l.writeHeader(eas.streamID, true, headerFields, nil); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:796
		_go_fuzz_dep_.CoverTab[76837]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:797
		// _ = "end of CoverTab[76837]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:798
		_go_fuzz_dep_.CoverTab[76838]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:798
		// _ = "end of CoverTab[76838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:798
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:798
	// _ = "end of CoverTab[76830]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:798
	_go_fuzz_dep_.CoverTab[76831]++
													if eas.rst {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:799
		_go_fuzz_dep_.CoverTab[76839]++
														if err := l.framer.fr.WriteRSTStream(eas.streamID, http2.ErrCodeNo); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:800
			_go_fuzz_dep_.CoverTab[76840]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:801
			// _ = "end of CoverTab[76840]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:802
			_go_fuzz_dep_.CoverTab[76841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:802
			// _ = "end of CoverTab[76841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:802
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:802
		// _ = "end of CoverTab[76839]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:803
		_go_fuzz_dep_.CoverTab[76842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:803
		// _ = "end of CoverTab[76842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:803
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:803
	// _ = "end of CoverTab[76831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:803
	_go_fuzz_dep_.CoverTab[76832]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:804
	// _ = "end of CoverTab[76832]"
}

func (l *loopyWriter) incomingGoAwayHandler(*incomingGoAway) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:807
	_go_fuzz_dep_.CoverTab[76843]++
													if l.side == clientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:808
		_go_fuzz_dep_.CoverTab[76845]++
														l.draining = true
														if len(l.estdStreams) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:810
			_go_fuzz_dep_.CoverTab[76846]++

															return errors.New("received GOAWAY with no active streams")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:812
			// _ = "end of CoverTab[76846]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:813
			_go_fuzz_dep_.CoverTab[76847]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:813
			// _ = "end of CoverTab[76847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:813
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:813
		// _ = "end of CoverTab[76845]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:814
		_go_fuzz_dep_.CoverTab[76848]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:814
		// _ = "end of CoverTab[76848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:814
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:814
	// _ = "end of CoverTab[76843]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:814
	_go_fuzz_dep_.CoverTab[76844]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:815
	// _ = "end of CoverTab[76844]"
}

func (l *loopyWriter) goAwayHandler(g *goAway) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:818
	_go_fuzz_dep_.CoverTab[76849]++

													if l.ssGoAwayHandler != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:820
		_go_fuzz_dep_.CoverTab[76851]++
														draining, err := l.ssGoAwayHandler(g)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:822
			_go_fuzz_dep_.CoverTab[76853]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:823
			// _ = "end of CoverTab[76853]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:824
			_go_fuzz_dep_.CoverTab[76854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:824
			// _ = "end of CoverTab[76854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:824
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:824
		// _ = "end of CoverTab[76851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:824
		_go_fuzz_dep_.CoverTab[76852]++
														l.draining = draining
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:825
		// _ = "end of CoverTab[76852]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:826
		_go_fuzz_dep_.CoverTab[76855]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:826
		// _ = "end of CoverTab[76855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:826
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:826
	// _ = "end of CoverTab[76849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:826
	_go_fuzz_dep_.CoverTab[76850]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:827
	// _ = "end of CoverTab[76850]"
}

func (l *loopyWriter) handle(i interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:830
	_go_fuzz_dep_.CoverTab[76856]++
													switch i := i.(type) {
	case *incomingWindowUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:832
		_go_fuzz_dep_.CoverTab[76858]++
														l.incomingWindowUpdateHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:833
		// _ = "end of CoverTab[76858]"
	case *outgoingWindowUpdate:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:834
		_go_fuzz_dep_.CoverTab[76859]++
														return l.outgoingWindowUpdateHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:835
		// _ = "end of CoverTab[76859]"
	case *incomingSettings:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:836
		_go_fuzz_dep_.CoverTab[76860]++
														return l.incomingSettingsHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:837
		// _ = "end of CoverTab[76860]"
	case *outgoingSettings:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:838
		_go_fuzz_dep_.CoverTab[76861]++
														return l.outgoingSettingsHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:839
		// _ = "end of CoverTab[76861]"
	case *headerFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:840
		_go_fuzz_dep_.CoverTab[76862]++
														return l.headerHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:841
		// _ = "end of CoverTab[76862]"
	case *registerStream:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:842
		_go_fuzz_dep_.CoverTab[76863]++
														l.registerStreamHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:843
		// _ = "end of CoverTab[76863]"
	case *cleanupStream:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:844
		_go_fuzz_dep_.CoverTab[76864]++
														return l.cleanupStreamHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:845
		// _ = "end of CoverTab[76864]"
	case *earlyAbortStream:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:846
		_go_fuzz_dep_.CoverTab[76865]++
														return l.earlyAbortStreamHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:847
		// _ = "end of CoverTab[76865]"
	case *incomingGoAway:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:848
		_go_fuzz_dep_.CoverTab[76866]++
														return l.incomingGoAwayHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:849
		// _ = "end of CoverTab[76866]"
	case *dataFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:850
		_go_fuzz_dep_.CoverTab[76867]++
														l.preprocessData(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:851
		// _ = "end of CoverTab[76867]"
	case *ping:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:852
		_go_fuzz_dep_.CoverTab[76868]++
														return l.pingHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:853
		// _ = "end of CoverTab[76868]"
	case *goAway:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:854
		_go_fuzz_dep_.CoverTab[76869]++
														return l.goAwayHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:855
		// _ = "end of CoverTab[76869]"
	case *outFlowControlSizeRequest:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:856
		_go_fuzz_dep_.CoverTab[76870]++
														l.outFlowControlSizeRequestHandler(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:857
		// _ = "end of CoverTab[76870]"
	case closeConnection:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:858
		_go_fuzz_dep_.CoverTab[76871]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:861
		return ErrConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:861
		// _ = "end of CoverTab[76871]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:862
		_go_fuzz_dep_.CoverTab[76872]++
														return fmt.Errorf("transport: unknown control message type %T", i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:863
		// _ = "end of CoverTab[76872]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:864
	// _ = "end of CoverTab[76856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:864
	_go_fuzz_dep_.CoverTab[76857]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:865
	// _ = "end of CoverTab[76857]"
}

func (l *loopyWriter) applySettings(ss []http2.Setting) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:868
	_go_fuzz_dep_.CoverTab[76873]++
													for _, s := range ss {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:869
		_go_fuzz_dep_.CoverTab[76874]++
														switch s.ID {
		case http2.SettingInitialWindowSize:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:871
			_go_fuzz_dep_.CoverTab[76875]++
															o := l.oiws
															l.oiws = s.Val
															if o < l.oiws {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:874
				_go_fuzz_dep_.CoverTab[76878]++

																for _, stream := range l.estdStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:876
					_go_fuzz_dep_.CoverTab[76879]++
																	if stream.state == waitingOnStreamQuota {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:877
						_go_fuzz_dep_.CoverTab[76880]++
																		stream.state = active
																		l.activeStreams.enqueue(stream)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:879
						// _ = "end of CoverTab[76880]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:880
						_go_fuzz_dep_.CoverTab[76881]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:880
						// _ = "end of CoverTab[76881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:880
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:880
					// _ = "end of CoverTab[76879]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:881
				// _ = "end of CoverTab[76878]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:882
				_go_fuzz_dep_.CoverTab[76882]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:882
				// _ = "end of CoverTab[76882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:882
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:882
			// _ = "end of CoverTab[76875]"
		case http2.SettingHeaderTableSize:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:883
			_go_fuzz_dep_.CoverTab[76876]++
															updateHeaderTblSize(l.hEnc, s.Val)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:884
			// _ = "end of CoverTab[76876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:884
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:884
			_go_fuzz_dep_.CoverTab[76877]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:884
			// _ = "end of CoverTab[76877]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:885
		// _ = "end of CoverTab[76874]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:886
	// _ = "end of CoverTab[76873]"
}

// processData removes the first stream from active streams, writes out at most 16KB
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:889
// of its data and then puts it at the end of activeStreams if there's still more data
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:889
// to be sent and stream has some stream-level flow control.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:892
func (l *loopyWriter) processData() (bool, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:892
	_go_fuzz_dep_.CoverTab[76883]++
													if l.sendQuota == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:893
		_go_fuzz_dep_.CoverTab[76895]++
														return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:894
		// _ = "end of CoverTab[76895]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:895
		_go_fuzz_dep_.CoverTab[76896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:895
		// _ = "end of CoverTab[76896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:895
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:895
	// _ = "end of CoverTab[76883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:895
	_go_fuzz_dep_.CoverTab[76884]++
													str := l.activeStreams.dequeue()
													if str == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:897
		_go_fuzz_dep_.CoverTab[76897]++
														return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:898
		// _ = "end of CoverTab[76897]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:899
		_go_fuzz_dep_.CoverTab[76898]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:899
		// _ = "end of CoverTab[76898]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:899
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:899
	// _ = "end of CoverTab[76884]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:899
	_go_fuzz_dep_.CoverTab[76885]++
													dataItem := str.itl.peek().(*dataFrame)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
	if len(dataItem.h) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
		_go_fuzz_dep_.CoverTab[76899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
		return len(dataItem.d) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
		// _ = "end of CoverTab[76899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:907
		_go_fuzz_dep_.CoverTab[76900]++

														if err := l.framer.fr.WriteData(dataItem.streamID, dataItem.endStream, nil); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:909
			_go_fuzz_dep_.CoverTab[76903]++
															return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:910
			// _ = "end of CoverTab[76903]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:911
			_go_fuzz_dep_.CoverTab[76904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:911
			// _ = "end of CoverTab[76904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:911
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:911
		// _ = "end of CoverTab[76900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:911
		_go_fuzz_dep_.CoverTab[76901]++
														str.itl.dequeue()
														if str.itl.isEmpty() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:913
			_go_fuzz_dep_.CoverTab[76905]++
															str.state = empty
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:914
			// _ = "end of CoverTab[76905]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:915
			_go_fuzz_dep_.CoverTab[76906]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:915
			if trailer, ok := str.itl.peek().(*headerFrame); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:915
				_go_fuzz_dep_.CoverTab[76907]++
																if err := l.writeHeader(trailer.streamID, trailer.endStream, trailer.hf, trailer.onWrite); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:916
					_go_fuzz_dep_.CoverTab[76909]++
																	return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:917
					// _ = "end of CoverTab[76909]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:918
					_go_fuzz_dep_.CoverTab[76910]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:918
					// _ = "end of CoverTab[76910]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:918
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:918
				// _ = "end of CoverTab[76907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:918
				_go_fuzz_dep_.CoverTab[76908]++
																if err := l.cleanupStreamHandler(trailer.cleanup); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:919
					_go_fuzz_dep_.CoverTab[76911]++
																	return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:920
					// _ = "end of CoverTab[76911]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:921
					_go_fuzz_dep_.CoverTab[76912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:921
					// _ = "end of CoverTab[76912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:921
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:921
				// _ = "end of CoverTab[76908]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:922
				_go_fuzz_dep_.CoverTab[76913]++
																l.activeStreams.enqueue(str)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:923
				// _ = "end of CoverTab[76913]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:924
			// _ = "end of CoverTab[76906]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:924
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:924
		// _ = "end of CoverTab[76901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:924
		_go_fuzz_dep_.CoverTab[76902]++
														return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:925
		// _ = "end of CoverTab[76902]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:926
		_go_fuzz_dep_.CoverTab[76914]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:926
		// _ = "end of CoverTab[76914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:926
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:926
	// _ = "end of CoverTab[76885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:926
	_go_fuzz_dep_.CoverTab[76886]++
													var (
		buf []byte
	)

	maxSize := http2MaxFrameLen
	if strQuota := int(l.oiws) - str.bytesOutStanding; strQuota <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:932
		_go_fuzz_dep_.CoverTab[76915]++
														str.state = waitingOnStreamQuota
														return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:934
		// _ = "end of CoverTab[76915]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:935
		_go_fuzz_dep_.CoverTab[76916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:935
		if maxSize > strQuota {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:935
			_go_fuzz_dep_.CoverTab[76917]++
															maxSize = strQuota
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:936
			// _ = "end of CoverTab[76917]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
			_go_fuzz_dep_.CoverTab[76918]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
			// _ = "end of CoverTab[76918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
		// _ = "end of CoverTab[76916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
	// _ = "end of CoverTab[76886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:937
	_go_fuzz_dep_.CoverTab[76887]++
													if maxSize > int(l.sendQuota) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:938
		_go_fuzz_dep_.CoverTab[76919]++
														maxSize = int(l.sendQuota)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:939
		// _ = "end of CoverTab[76919]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:940
		_go_fuzz_dep_.CoverTab[76920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:940
		// _ = "end of CoverTab[76920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:940
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:940
	// _ = "end of CoverTab[76887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:940
	_go_fuzz_dep_.CoverTab[76888]++

													hSize := min(maxSize, len(dataItem.h))
													dSize := min(maxSize-hSize, len(dataItem.d))
													if hSize != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:944
		_go_fuzz_dep_.CoverTab[76921]++
														if dSize == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:945
			_go_fuzz_dep_.CoverTab[76922]++
															buf = dataItem.h
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:946
			// _ = "end of CoverTab[76922]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:947
			_go_fuzz_dep_.CoverTab[76923]++
			// We can add some data to grpc message header to distribute bytes more equally across frames.
															// Copy on the stack to avoid generating garbage
															var localBuf [http2MaxFrameLen]byte
															copy(localBuf[:hSize], dataItem.h)
															copy(localBuf[hSize:], dataItem.d[:dSize])
															buf = localBuf[:hSize+dSize]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:953
			// _ = "end of CoverTab[76923]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:954
		// _ = "end of CoverTab[76921]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:955
		_go_fuzz_dep_.CoverTab[76924]++
														buf = dataItem.d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:956
		// _ = "end of CoverTab[76924]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:957
	// _ = "end of CoverTab[76888]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:957
	_go_fuzz_dep_.CoverTab[76889]++

													size := hSize + dSize

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:962
	str.wq.replenish(size)
	var endStream bool

	if dataItem.endStream && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:965
		_go_fuzz_dep_.CoverTab[76925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:965
		return len(dataItem.h)+len(dataItem.d) <= size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:965
		// _ = "end of CoverTab[76925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:965
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:965
		_go_fuzz_dep_.CoverTab[76926]++
														endStream = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:966
		// _ = "end of CoverTab[76926]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:967
		_go_fuzz_dep_.CoverTab[76927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:967
		// _ = "end of CoverTab[76927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:967
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:967
	// _ = "end of CoverTab[76889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:967
	_go_fuzz_dep_.CoverTab[76890]++
													if dataItem.onEachWrite != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:968
		_go_fuzz_dep_.CoverTab[76928]++
														dataItem.onEachWrite()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:969
		// _ = "end of CoverTab[76928]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:970
		_go_fuzz_dep_.CoverTab[76929]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:970
		// _ = "end of CoverTab[76929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:970
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:970
	// _ = "end of CoverTab[76890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:970
	_go_fuzz_dep_.CoverTab[76891]++
													if err := l.framer.fr.WriteData(dataItem.streamID, endStream, buf[:size]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:971
		_go_fuzz_dep_.CoverTab[76930]++
														return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:972
		// _ = "end of CoverTab[76930]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:973
		_go_fuzz_dep_.CoverTab[76931]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:973
		// _ = "end of CoverTab[76931]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:973
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:973
	// _ = "end of CoverTab[76891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:973
	_go_fuzz_dep_.CoverTab[76892]++
													str.bytesOutStanding += size
													l.sendQuota -= uint32(size)
													dataItem.h = dataItem.h[hSize:]
													dataItem.d = dataItem.d[dSize:]

													if len(dataItem.h) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:979
		_go_fuzz_dep_.CoverTab[76932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:979
		return len(dataItem.d) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:979
		// _ = "end of CoverTab[76932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:979
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:979
		_go_fuzz_dep_.CoverTab[76933]++
														str.itl.dequeue()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:980
		// _ = "end of CoverTab[76933]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:981
		_go_fuzz_dep_.CoverTab[76934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:981
		// _ = "end of CoverTab[76934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:981
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:981
	// _ = "end of CoverTab[76892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:981
	_go_fuzz_dep_.CoverTab[76893]++
													if str.itl.isEmpty() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:982
		_go_fuzz_dep_.CoverTab[76935]++
														str.state = empty
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:983
		// _ = "end of CoverTab[76935]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:984
		_go_fuzz_dep_.CoverTab[76936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:984
		if trailer, ok := str.itl.peek().(*headerFrame); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:984
			_go_fuzz_dep_.CoverTab[76937]++
															if err := l.writeHeader(trailer.streamID, trailer.endStream, trailer.hf, trailer.onWrite); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:985
				_go_fuzz_dep_.CoverTab[76939]++
																return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:986
				// _ = "end of CoverTab[76939]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:987
				_go_fuzz_dep_.CoverTab[76940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:987
				// _ = "end of CoverTab[76940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:987
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:987
			// _ = "end of CoverTab[76937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:987
			_go_fuzz_dep_.CoverTab[76938]++
															if err := l.cleanupStreamHandler(trailer.cleanup); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:988
				_go_fuzz_dep_.CoverTab[76941]++
																return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:989
				// _ = "end of CoverTab[76941]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:990
				_go_fuzz_dep_.CoverTab[76942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:990
				// _ = "end of CoverTab[76942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:990
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:990
			// _ = "end of CoverTab[76938]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:991
			_go_fuzz_dep_.CoverTab[76943]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:991
			if int(l.oiws)-str.bytesOutStanding <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:991
				_go_fuzz_dep_.CoverTab[76944]++
																str.state = waitingOnStreamQuota
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:992
				// _ = "end of CoverTab[76944]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:993
				_go_fuzz_dep_.CoverTab[76945]++
																l.activeStreams.enqueue(str)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:994
				// _ = "end of CoverTab[76945]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
			// _ = "end of CoverTab[76943]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
		// _ = "end of CoverTab[76936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
	// _ = "end of CoverTab[76893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:995
	_go_fuzz_dep_.CoverTab[76894]++
													return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:996
	// _ = "end of CoverTab[76894]"
}

func min(a, b int) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:999
	_go_fuzz_dep_.CoverTab[76946]++
													if a < b {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1000
		_go_fuzz_dep_.CoverTab[76948]++
														return a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1001
		// _ = "end of CoverTab[76948]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1002
		_go_fuzz_dep_.CoverTab[76949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1002
		// _ = "end of CoverTab[76949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1002
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1002
	// _ = "end of CoverTab[76946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1002
	_go_fuzz_dep_.CoverTab[76947]++
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1003
	// _ = "end of CoverTab[76947]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1004
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/controlbuf.go:1004
var _ = _go_fuzz_dep_.CoverTab
