//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
package binarylog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:19
)

import (
	"context"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	binlogpb "google.golang.org/grpc/binarylog/grpc_binarylog_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type callIDGenerator struct {
	id uint64
}

func (g *callIDGenerator) next() uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:39
	_go_fuzz_dep_.CoverTab[68722]++
														id := atomic.AddUint64(&g.id, 1)
														return id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:41
	// _ = "end of CoverTab[68722]"
}

// reset is for testing only, and doesn't need to be thread safe.
func (g *callIDGenerator) reset() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:45
	_go_fuzz_dep_.CoverTab[68723]++
														g.id = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:46
	// _ = "end of CoverTab[68723]"
}

var idGen callIDGenerator

// MethodLogger is the sub-logger for each method.
type MethodLogger interface {
	Log(context.Context, LogEntryConfig)
}

// TruncatingMethodLogger is a method logger that truncates headers and messages
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:56
// based on configured fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:58
type TruncatingMethodLogger struct {
	headerMaxLen, messageMaxLen	uint64

	callID		uint64
	idWithinCallGen	*callIDGenerator

	sink	Sink	// TODO(blog): make this plugable.
}

// NewTruncatingMethodLogger returns a new truncating method logger.
func NewTruncatingMethodLogger(h, m uint64) *TruncatingMethodLogger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:68
	_go_fuzz_dep_.CoverTab[68724]++
														return &TruncatingMethodLogger{
		headerMaxLen:	h,
		messageMaxLen:	m,

		callID:			idGen.next(),
		idWithinCallGen:	&callIDGenerator{},

		sink:	DefaultSink,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:77
	// _ = "end of CoverTab[68724]"
}

// Build is an internal only method for building the proto message out of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:80
// input event. It's made public to enable other library to reuse as much logic
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:80
// in TruncatingMethodLogger as possible.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:83
func (ml *TruncatingMethodLogger) Build(c LogEntryConfig) *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:83
	_go_fuzz_dep_.CoverTab[68725]++
														m := c.toProto()
														timestamp, _ := ptypes.TimestampProto(time.Now())
														m.Timestamp = timestamp
														m.CallId = ml.callID
														m.SequenceIdWithinCall = ml.idWithinCallGen.next()

														switch pay := m.Payload.(type) {
	case *binlogpb.GrpcLogEntry_ClientHeader:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:91
		_go_fuzz_dep_.CoverTab[68727]++
															m.PayloadTruncated = ml.truncateMetadata(pay.ClientHeader.GetMetadata())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:92
		// _ = "end of CoverTab[68727]"
	case *binlogpb.GrpcLogEntry_ServerHeader:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:93
		_go_fuzz_dep_.CoverTab[68728]++
															m.PayloadTruncated = ml.truncateMetadata(pay.ServerHeader.GetMetadata())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:94
		// _ = "end of CoverTab[68728]"
	case *binlogpb.GrpcLogEntry_Message:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:95
		_go_fuzz_dep_.CoverTab[68729]++
															m.PayloadTruncated = ml.truncateMessage(pay.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:96
		// _ = "end of CoverTab[68729]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:97
	// _ = "end of CoverTab[68725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:97
	_go_fuzz_dep_.CoverTab[68726]++
														return m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:98
	// _ = "end of CoverTab[68726]"
}

// Log creates a proto binary log entry, and logs it to the sink.
func (ml *TruncatingMethodLogger) Log(ctx context.Context, c LogEntryConfig) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:102
	_go_fuzz_dep_.CoverTab[68730]++
														ml.sink.Write(ml.Build(c))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:103
	// _ = "end of CoverTab[68730]"
}

func (ml *TruncatingMethodLogger) truncateMetadata(mdPb *binlogpb.Metadata) (truncated bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:106
	_go_fuzz_dep_.CoverTab[68731]++
														if ml.headerMaxLen == maxUInt {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:107
		_go_fuzz_dep_.CoverTab[68734]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:108
		// _ = "end of CoverTab[68734]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:109
		_go_fuzz_dep_.CoverTab[68735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:109
		// _ = "end of CoverTab[68735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:109
	// _ = "end of CoverTab[68731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:109
	_go_fuzz_dep_.CoverTab[68732]++
														var (
		bytesLimit	= ml.headerMaxLen
		index		int
	)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:118
	for ; index < len(mdPb.Entry); index++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:118
		_go_fuzz_dep_.CoverTab[68736]++
															entry := mdPb.Entry[index]
															if entry.Key == "grpc-trace-bin" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:120
			_go_fuzz_dep_.CoverTab[68739]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:123
			continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:123
			// _ = "end of CoverTab[68739]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:124
			_go_fuzz_dep_.CoverTab[68740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:124
			// _ = "end of CoverTab[68740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:124
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:124
		// _ = "end of CoverTab[68736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:124
		_go_fuzz_dep_.CoverTab[68737]++
															currentEntryLen := uint64(len(entry.GetKey())) + uint64(len(entry.GetValue()))
															if currentEntryLen > bytesLimit {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:126
			_go_fuzz_dep_.CoverTab[68741]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:127
			// _ = "end of CoverTab[68741]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:128
			_go_fuzz_dep_.CoverTab[68742]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:128
			// _ = "end of CoverTab[68742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:128
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:128
		// _ = "end of CoverTab[68737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:128
		_go_fuzz_dep_.CoverTab[68738]++
															bytesLimit -= currentEntryLen
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:129
		// _ = "end of CoverTab[68738]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:130
	// _ = "end of CoverTab[68732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:130
	_go_fuzz_dep_.CoverTab[68733]++
														truncated = index < len(mdPb.Entry)
														mdPb.Entry = mdPb.Entry[:index]
														return truncated
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:133
	// _ = "end of CoverTab[68733]"
}

func (ml *TruncatingMethodLogger) truncateMessage(msgPb *binlogpb.Message) (truncated bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:136
	_go_fuzz_dep_.CoverTab[68743]++
														if ml.messageMaxLen == maxUInt {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:137
		_go_fuzz_dep_.CoverTab[68746]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:138
		// _ = "end of CoverTab[68746]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:139
		_go_fuzz_dep_.CoverTab[68747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:139
		// _ = "end of CoverTab[68747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:139
	// _ = "end of CoverTab[68743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:139
	_go_fuzz_dep_.CoverTab[68744]++
														if ml.messageMaxLen >= uint64(len(msgPb.Data)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:140
		_go_fuzz_dep_.CoverTab[68748]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:141
		// _ = "end of CoverTab[68748]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:142
		_go_fuzz_dep_.CoverTab[68749]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:142
		// _ = "end of CoverTab[68749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:142
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:142
	// _ = "end of CoverTab[68744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:142
	_go_fuzz_dep_.CoverTab[68745]++
														msgPb.Data = msgPb.Data[:ml.messageMaxLen]
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:144
	// _ = "end of CoverTab[68745]"
}

// LogEntryConfig represents the configuration for binary log entry.
type LogEntryConfig interface {
	toProto() *binlogpb.GrpcLogEntry
}

// ClientHeader configs the binary log entry to be a ClientHeader entry.
type ClientHeader struct {
	OnClientSide	bool
	Header		metadata.MD
	MethodName	string
	Authority	string
	Timeout		time.Duration
	// PeerAddr is required only when it's on server side.
	PeerAddr	net.Addr
}

func (c *ClientHeader) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:163
	_go_fuzz_dep_.CoverTab[68750]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:166
	clientHeader := &binlogpb.ClientHeader{
		Metadata:	mdToMetadataProto(c.Header),
		MethodName:	c.MethodName,
		Authority:	c.Authority,
	}
	if c.Timeout > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:171
		_go_fuzz_dep_.CoverTab[68754]++
															clientHeader.Timeout = ptypes.DurationProto(c.Timeout)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:172
		// _ = "end of CoverTab[68754]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:173
		_go_fuzz_dep_.CoverTab[68755]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:173
		// _ = "end of CoverTab[68755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:173
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:173
	// _ = "end of CoverTab[68750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:173
	_go_fuzz_dep_.CoverTab[68751]++
														ret := &binlogpb.GrpcLogEntry{
		Type:	binlogpb.GrpcLogEntry_EVENT_TYPE_CLIENT_HEADER,
		Payload: &binlogpb.GrpcLogEntry_ClientHeader{
			ClientHeader: clientHeader,
		},
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:180
		_go_fuzz_dep_.CoverTab[68756]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:181
		// _ = "end of CoverTab[68756]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:182
		_go_fuzz_dep_.CoverTab[68757]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:183
		// _ = "end of CoverTab[68757]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:184
	// _ = "end of CoverTab[68751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:184
	_go_fuzz_dep_.CoverTab[68752]++
														if c.PeerAddr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:185
		_go_fuzz_dep_.CoverTab[68758]++
															ret.Peer = addrToProto(c.PeerAddr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:186
		// _ = "end of CoverTab[68758]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:187
		_go_fuzz_dep_.CoverTab[68759]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:187
		// _ = "end of CoverTab[68759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:187
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:187
	// _ = "end of CoverTab[68752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:187
	_go_fuzz_dep_.CoverTab[68753]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:188
	// _ = "end of CoverTab[68753]"
}

// ServerHeader configs the binary log entry to be a ServerHeader entry.
type ServerHeader struct {
	OnClientSide	bool
	Header		metadata.MD
	// PeerAddr is required only when it's on client side.
	PeerAddr	net.Addr
}

func (c *ServerHeader) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:199
	_go_fuzz_dep_.CoverTab[68760]++
														ret := &binlogpb.GrpcLogEntry{
		Type:	binlogpb.GrpcLogEntry_EVENT_TYPE_SERVER_HEADER,
		Payload: &binlogpb.GrpcLogEntry_ServerHeader{
			ServerHeader: &binlogpb.ServerHeader{
				Metadata: mdToMetadataProto(c.Header),
			},
		},
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:208
		_go_fuzz_dep_.CoverTab[68763]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:209
		// _ = "end of CoverTab[68763]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:210
		_go_fuzz_dep_.CoverTab[68764]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:211
		// _ = "end of CoverTab[68764]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:212
	// _ = "end of CoverTab[68760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:212
	_go_fuzz_dep_.CoverTab[68761]++
														if c.PeerAddr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:213
		_go_fuzz_dep_.CoverTab[68765]++
															ret.Peer = addrToProto(c.PeerAddr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:214
		// _ = "end of CoverTab[68765]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:215
		_go_fuzz_dep_.CoverTab[68766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:215
		// _ = "end of CoverTab[68766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:215
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:215
	// _ = "end of CoverTab[68761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:215
	_go_fuzz_dep_.CoverTab[68762]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:216
	// _ = "end of CoverTab[68762]"
}

// ClientMessage configs the binary log entry to be a ClientMessage entry.
type ClientMessage struct {
	OnClientSide	bool
	// Message can be a proto.Message or []byte. Other messages formats are not
	// supported.
	Message	interface{}
}

func (c *ClientMessage) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:227
	_go_fuzz_dep_.CoverTab[68767]++
														var (
		data	[]byte
		err	error
	)
	if m, ok := c.Message.(proto.Message); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:232
		_go_fuzz_dep_.CoverTab[68770]++
															data, err = proto.Marshal(m)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:234
			_go_fuzz_dep_.CoverTab[68771]++
																grpclogLogger.Infof("binarylogging: failed to marshal proto message: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:235
			// _ = "end of CoverTab[68771]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:236
			_go_fuzz_dep_.CoverTab[68772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:236
			// _ = "end of CoverTab[68772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:236
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:236
		// _ = "end of CoverTab[68770]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:237
		_go_fuzz_dep_.CoverTab[68773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:237
		if b, ok := c.Message.([]byte); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:237
			_go_fuzz_dep_.CoverTab[68774]++
																data = b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:238
			// _ = "end of CoverTab[68774]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:239
			_go_fuzz_dep_.CoverTab[68775]++
																grpclogLogger.Infof("binarylogging: message to log is neither proto.message nor []byte")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:240
			// _ = "end of CoverTab[68775]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:241
		// _ = "end of CoverTab[68773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:241
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:241
	// _ = "end of CoverTab[68767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:241
	_go_fuzz_dep_.CoverTab[68768]++
														ret := &binlogpb.GrpcLogEntry{
		Type:	binlogpb.GrpcLogEntry_EVENT_TYPE_CLIENT_MESSAGE,
		Payload: &binlogpb.GrpcLogEntry_Message{
			Message: &binlogpb.Message{
				Length:	uint32(len(data)),
				Data:	data,
			},
		},
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:251
		_go_fuzz_dep_.CoverTab[68776]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:252
		// _ = "end of CoverTab[68776]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:253
		_go_fuzz_dep_.CoverTab[68777]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:254
		// _ = "end of CoverTab[68777]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:255
	// _ = "end of CoverTab[68768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:255
	_go_fuzz_dep_.CoverTab[68769]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:256
	// _ = "end of CoverTab[68769]"
}

// ServerMessage configs the binary log entry to be a ServerMessage entry.
type ServerMessage struct {
	OnClientSide	bool
	// Message can be a proto.Message or []byte. Other messages formats are not
	// supported.
	Message	interface{}
}

func (c *ServerMessage) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:267
	_go_fuzz_dep_.CoverTab[68778]++
														var (
		data	[]byte
		err	error
	)
	if m, ok := c.Message.(proto.Message); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:272
		_go_fuzz_dep_.CoverTab[68781]++
															data, err = proto.Marshal(m)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:274
			_go_fuzz_dep_.CoverTab[68782]++
																grpclogLogger.Infof("binarylogging: failed to marshal proto message: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:275
			// _ = "end of CoverTab[68782]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:276
			_go_fuzz_dep_.CoverTab[68783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:276
			// _ = "end of CoverTab[68783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:276
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:276
		// _ = "end of CoverTab[68781]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:277
		_go_fuzz_dep_.CoverTab[68784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:277
		if b, ok := c.Message.([]byte); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:277
			_go_fuzz_dep_.CoverTab[68785]++
																data = b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:278
			// _ = "end of CoverTab[68785]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:279
			_go_fuzz_dep_.CoverTab[68786]++
																grpclogLogger.Infof("binarylogging: message to log is neither proto.message nor []byte")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:280
			// _ = "end of CoverTab[68786]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:281
		// _ = "end of CoverTab[68784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:281
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:281
	// _ = "end of CoverTab[68778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:281
	_go_fuzz_dep_.CoverTab[68779]++
														ret := &binlogpb.GrpcLogEntry{
		Type:	binlogpb.GrpcLogEntry_EVENT_TYPE_SERVER_MESSAGE,
		Payload: &binlogpb.GrpcLogEntry_Message{
			Message: &binlogpb.Message{
				Length:	uint32(len(data)),
				Data:	data,
			},
		},
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:291
		_go_fuzz_dep_.CoverTab[68787]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:292
		// _ = "end of CoverTab[68787]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:293
		_go_fuzz_dep_.CoverTab[68788]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:294
		// _ = "end of CoverTab[68788]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:295
	// _ = "end of CoverTab[68779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:295
	_go_fuzz_dep_.CoverTab[68780]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:296
	// _ = "end of CoverTab[68780]"
}

// ClientHalfClose configs the binary log entry to be a ClientHalfClose entry.
type ClientHalfClose struct {
	OnClientSide bool
}

func (c *ClientHalfClose) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:304
	_go_fuzz_dep_.CoverTab[68789]++
														ret := &binlogpb.GrpcLogEntry{
		Type:		binlogpb.GrpcLogEntry_EVENT_TYPE_CLIENT_HALF_CLOSE,
		Payload:	nil,
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:309
		_go_fuzz_dep_.CoverTab[68791]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:310
		// _ = "end of CoverTab[68791]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:311
		_go_fuzz_dep_.CoverTab[68792]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:312
		// _ = "end of CoverTab[68792]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:313
	// _ = "end of CoverTab[68789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:313
	_go_fuzz_dep_.CoverTab[68790]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:314
	// _ = "end of CoverTab[68790]"
}

// ServerTrailer configs the binary log entry to be a ServerTrailer entry.
type ServerTrailer struct {
	OnClientSide	bool
	Trailer		metadata.MD
	// Err is the status error.
	Err	error
	// PeerAddr is required only when it's on client side and the RPC is trailer
	// only.
	PeerAddr	net.Addr
}

func (c *ServerTrailer) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:328
	_go_fuzz_dep_.CoverTab[68793]++
														st, ok := status.FromError(c.Err)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:330
		_go_fuzz_dep_.CoverTab[68798]++
															grpclogLogger.Info("binarylogging: error in trailer is not a status error")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:331
		// _ = "end of CoverTab[68798]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:332
		_go_fuzz_dep_.CoverTab[68799]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:332
		// _ = "end of CoverTab[68799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:332
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:332
	// _ = "end of CoverTab[68793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:332
	_go_fuzz_dep_.CoverTab[68794]++
														var (
		detailsBytes	[]byte
		err		error
	)
	stProto := st.Proto()
	if stProto != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:338
		_go_fuzz_dep_.CoverTab[68800]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:338
		return len(stProto.Details) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:338
		// _ = "end of CoverTab[68800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:338
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:338
		_go_fuzz_dep_.CoverTab[68801]++
															detailsBytes, err = proto.Marshal(stProto)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:340
			_go_fuzz_dep_.CoverTab[68802]++
																grpclogLogger.Infof("binarylogging: failed to marshal status proto: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:341
			// _ = "end of CoverTab[68802]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:342
			_go_fuzz_dep_.CoverTab[68803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:342
			// _ = "end of CoverTab[68803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:342
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:342
		// _ = "end of CoverTab[68801]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:343
		_go_fuzz_dep_.CoverTab[68804]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:343
		// _ = "end of CoverTab[68804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:343
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:343
	// _ = "end of CoverTab[68794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:343
	_go_fuzz_dep_.CoverTab[68795]++
														ret := &binlogpb.GrpcLogEntry{
		Type:	binlogpb.GrpcLogEntry_EVENT_TYPE_SERVER_TRAILER,
		Payload: &binlogpb.GrpcLogEntry_Trailer{
			Trailer: &binlogpb.Trailer{
				Metadata:	mdToMetadataProto(c.Trailer),
				StatusCode:	uint32(st.Code()),
				StatusMessage:	st.Message(),
				StatusDetails:	detailsBytes,
			},
		},
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:355
		_go_fuzz_dep_.CoverTab[68805]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:356
		// _ = "end of CoverTab[68805]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:357
		_go_fuzz_dep_.CoverTab[68806]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:358
		// _ = "end of CoverTab[68806]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:359
	// _ = "end of CoverTab[68795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:359
	_go_fuzz_dep_.CoverTab[68796]++
														if c.PeerAddr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:360
		_go_fuzz_dep_.CoverTab[68807]++
															ret.Peer = addrToProto(c.PeerAddr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:361
		// _ = "end of CoverTab[68807]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:362
		_go_fuzz_dep_.CoverTab[68808]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:362
		// _ = "end of CoverTab[68808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:362
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:362
	// _ = "end of CoverTab[68796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:362
	_go_fuzz_dep_.CoverTab[68797]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:363
	// _ = "end of CoverTab[68797]"
}

// Cancel configs the binary log entry to be a Cancel entry.
type Cancel struct {
	OnClientSide bool
}

func (c *Cancel) toProto() *binlogpb.GrpcLogEntry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:371
	_go_fuzz_dep_.CoverTab[68809]++
														ret := &binlogpb.GrpcLogEntry{
		Type:		binlogpb.GrpcLogEntry_EVENT_TYPE_CANCEL,
		Payload:	nil,
	}
	if c.OnClientSide {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:376
		_go_fuzz_dep_.CoverTab[68811]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_CLIENT
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:377
		// _ = "end of CoverTab[68811]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:378
		_go_fuzz_dep_.CoverTab[68812]++
															ret.Logger = binlogpb.GrpcLogEntry_LOGGER_SERVER
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:379
		// _ = "end of CoverTab[68812]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:380
	// _ = "end of CoverTab[68809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:380
	_go_fuzz_dep_.CoverTab[68810]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:381
	// _ = "end of CoverTab[68810]"
}

// metadataKeyOmit returns whether the metadata entry with this key should be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:384
// omitted.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:386
func metadataKeyOmit(key string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:386
	_go_fuzz_dep_.CoverTab[68813]++
														switch key {
	case "lb-token", ":path", ":authority", "content-encoding", "content-type", "user-agent", "te":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:388
		_go_fuzz_dep_.CoverTab[68815]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:389
		// _ = "end of CoverTab[68815]"
	case "grpc-trace-bin":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:390
		_go_fuzz_dep_.CoverTab[68816]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:391
		// _ = "end of CoverTab[68816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:391
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:391
		_go_fuzz_dep_.CoverTab[68817]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:391
		// _ = "end of CoverTab[68817]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:392
	// _ = "end of CoverTab[68813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:392
	_go_fuzz_dep_.CoverTab[68814]++
														return strings.HasPrefix(key, "grpc-")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:393
	// _ = "end of CoverTab[68814]"
}

func mdToMetadataProto(md metadata.MD) *binlogpb.Metadata {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:396
	_go_fuzz_dep_.CoverTab[68818]++
														ret := &binlogpb.Metadata{}
														for k, vv := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:398
		_go_fuzz_dep_.CoverTab[68820]++
															if metadataKeyOmit(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:399
			_go_fuzz_dep_.CoverTab[68822]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:400
			// _ = "end of CoverTab[68822]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:401
			_go_fuzz_dep_.CoverTab[68823]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:401
			// _ = "end of CoverTab[68823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:401
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:401
		// _ = "end of CoverTab[68820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:401
		_go_fuzz_dep_.CoverTab[68821]++
															for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:402
			_go_fuzz_dep_.CoverTab[68824]++
																ret.Entry = append(ret.Entry,
				&binlogpb.MetadataEntry{
					Key:	k,
					Value:	[]byte(v),
				},
			)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:408
			// _ = "end of CoverTab[68824]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:409
		// _ = "end of CoverTab[68821]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:410
	// _ = "end of CoverTab[68818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:410
	_go_fuzz_dep_.CoverTab[68819]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:411
	// _ = "end of CoverTab[68819]"
}

func addrToProto(addr net.Addr) *binlogpb.Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:414
	_go_fuzz_dep_.CoverTab[68825]++
														ret := &binlogpb.Address{}
														switch a := addr.(type) {
	case *net.TCPAddr:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:417
		_go_fuzz_dep_.CoverTab[68827]++
															if a.IP.To4() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:418
			_go_fuzz_dep_.CoverTab[68831]++
																ret.Type = binlogpb.Address_TYPE_IPV4
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:419
			// _ = "end of CoverTab[68831]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:420
			_go_fuzz_dep_.CoverTab[68832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:420
			if a.IP.To16() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:420
				_go_fuzz_dep_.CoverTab[68833]++
																	ret.Type = binlogpb.Address_TYPE_IPV6
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:421
				// _ = "end of CoverTab[68833]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:422
				_go_fuzz_dep_.CoverTab[68834]++
																	ret.Type = binlogpb.Address_TYPE_UNKNOWN

																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:425
				// _ = "end of CoverTab[68834]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:426
			// _ = "end of CoverTab[68832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:426
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:426
		// _ = "end of CoverTab[68827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:426
		_go_fuzz_dep_.CoverTab[68828]++
															ret.Address = a.IP.String()
															ret.IpPort = uint32(a.Port)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:428
		// _ = "end of CoverTab[68828]"
	case *net.UnixAddr:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:429
		_go_fuzz_dep_.CoverTab[68829]++
															ret.Type = binlogpb.Address_TYPE_UNIX
															ret.Address = a.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:431
		// _ = "end of CoverTab[68829]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:432
		_go_fuzz_dep_.CoverTab[68830]++
															ret.Type = binlogpb.Address_TYPE_UNKNOWN
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:433
		// _ = "end of CoverTab[68830]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:434
	// _ = "end of CoverTab[68825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:434
	_go_fuzz_dep_.CoverTab[68826]++
														return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:435
	// _ = "end of CoverTab[68826]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:436
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/method_logger.go:436
var _ = _go_fuzz_dep_.CoverTab
