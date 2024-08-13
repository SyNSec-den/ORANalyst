//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:19
)

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

// Compressor defines the interface gRPC uses to compress a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:44
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:44
// Deprecated: use package encoding.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:47
type Compressor interface {
	// Do compresses p into w.
	Do(w io.Writer, p []byte) error
	// Type returns the compression algorithm the Compressor uses.
	Type() string
}

type gzipCompressor struct {
	pool sync.Pool
}

// NewGZIPCompressor creates a Compressor based on GZIP.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:58
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:58
// Deprecated: use package encoding/gzip.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:61
func NewGZIPCompressor() Compressor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:61
	_go_fuzz_dep_.CoverTab[79668]++
											c, _ := NewGZIPCompressorWithLevel(gzip.DefaultCompression)
											return c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:63
	// _ = "end of CoverTab[79668]"
}

// NewGZIPCompressorWithLevel is like NewGZIPCompressor but specifies the gzip compression level instead
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:66
// of assuming DefaultCompression.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:66
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:66
// The error returned will be nil if the level is valid.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:66
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:66
// Deprecated: use package encoding/gzip.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:72
func NewGZIPCompressorWithLevel(level int) (Compressor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:72
	_go_fuzz_dep_.CoverTab[79669]++
											if level < gzip.DefaultCompression || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:73
		_go_fuzz_dep_.CoverTab[79671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:73
		return level > gzip.BestCompression
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:73
		// _ = "end of CoverTab[79671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:73
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:73
		_go_fuzz_dep_.CoverTab[79672]++
												return nil, fmt.Errorf("grpc: invalid compression level: %d", level)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:74
		// _ = "end of CoverTab[79672]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:75
		_go_fuzz_dep_.CoverTab[79673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:75
		// _ = "end of CoverTab[79673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:75
	// _ = "end of CoverTab[79669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:75
	_go_fuzz_dep_.CoverTab[79670]++
											return &gzipCompressor{
		pool: sync.Pool{
			New: func() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:78
				_go_fuzz_dep_.CoverTab[79674]++
														w, err := gzip.NewWriterLevel(io.Discard, level)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:80
					_go_fuzz_dep_.CoverTab[79676]++
															panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:81
					// _ = "end of CoverTab[79676]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:82
					_go_fuzz_dep_.CoverTab[79677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:82
					// _ = "end of CoverTab[79677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:82
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:82
				// _ = "end of CoverTab[79674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:82
				_go_fuzz_dep_.CoverTab[79675]++
														return w
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:83
				// _ = "end of CoverTab[79675]"
			},
		},
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:86
	// _ = "end of CoverTab[79670]"
}

func (c *gzipCompressor) Do(w io.Writer, p []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:89
	_go_fuzz_dep_.CoverTab[79678]++
											z := c.pool.Get().(*gzip.Writer)
											defer c.pool.Put(z)
											z.Reset(w)
											if _, err := z.Write(p); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:93
		_go_fuzz_dep_.CoverTab[79680]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:94
		// _ = "end of CoverTab[79680]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:95
		_go_fuzz_dep_.CoverTab[79681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:95
		// _ = "end of CoverTab[79681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:95
	// _ = "end of CoverTab[79678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:95
	_go_fuzz_dep_.CoverTab[79679]++
											return z.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:96
	// _ = "end of CoverTab[79679]"
}

func (c *gzipCompressor) Type() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:99
	_go_fuzz_dep_.CoverTab[79682]++
											return "gzip"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:100
	// _ = "end of CoverTab[79682]"
}

// Decompressor defines the interface gRPC uses to decompress a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:103
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:103
// Deprecated: use package encoding.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:106
type Decompressor interface {
	// Do reads the data from r and uncompress them.
	Do(r io.Reader) ([]byte, error)
	// Type returns the compression algorithm the Decompressor uses.
	Type() string
}

type gzipDecompressor struct {
	pool sync.Pool
}

// NewGZIPDecompressor creates a Decompressor based on GZIP.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:117
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:117
// Deprecated: use package encoding/gzip.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:120
func NewGZIPDecompressor() Decompressor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:120
	_go_fuzz_dep_.CoverTab[79683]++
											return &gzipDecompressor{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:121
	// _ = "end of CoverTab[79683]"
}

func (d *gzipDecompressor) Do(r io.Reader) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:124
	_go_fuzz_dep_.CoverTab[79684]++
											var z *gzip.Reader
											switch maybeZ := d.pool.Get().(type) {
	case nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:127
		_go_fuzz_dep_.CoverTab[79687]++
												newZ, err := gzip.NewReader(r)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:129
			_go_fuzz_dep_.CoverTab[79690]++
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:130
			// _ = "end of CoverTab[79690]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:131
			_go_fuzz_dep_.CoverTab[79691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:131
			// _ = "end of CoverTab[79691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:131
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:131
		// _ = "end of CoverTab[79687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:131
		_go_fuzz_dep_.CoverTab[79688]++
												z = newZ
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:132
		// _ = "end of CoverTab[79688]"
	case *gzip.Reader:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:133
		_go_fuzz_dep_.CoverTab[79689]++
												z = maybeZ
												if err := z.Reset(r); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:135
			_go_fuzz_dep_.CoverTab[79692]++
													d.pool.Put(z)
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:137
			// _ = "end of CoverTab[79692]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:138
			_go_fuzz_dep_.CoverTab[79693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:138
			// _ = "end of CoverTab[79693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:138
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:138
		// _ = "end of CoverTab[79689]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:139
	// _ = "end of CoverTab[79684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:139
	_go_fuzz_dep_.CoverTab[79685]++

											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:141
		_go_fuzz_dep_.CoverTab[79694]++
												z.Close()
												d.pool.Put(z)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:143
		// _ = "end of CoverTab[79694]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:144
	// _ = "end of CoverTab[79685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:144
	_go_fuzz_dep_.CoverTab[79686]++
											return io.ReadAll(z)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:145
	// _ = "end of CoverTab[79686]"
}

func (d *gzipDecompressor) Type() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:148
	_go_fuzz_dep_.CoverTab[79695]++
											return "gzip"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:149
	// _ = "end of CoverTab[79695]"
}

// callInfo contains all related configuration and information about an RPC.
type callInfo struct {
	compressorType		string
	failFast		bool
	maxReceiveMessageSize	*int
	maxSendMessageSize	*int
	creds			credentials.PerRPCCredentials
	contentSubtype		string
	codec			baseCodec
	maxRetryRPCBufferSize	int
	onFinish		[]func(err error)
}

func defaultCallInfo() *callInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:165
	_go_fuzz_dep_.CoverTab[79696]++
											return &callInfo{
		failFast:		true,
		maxRetryRPCBufferSize:	256 * 1024,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:169
	// _ = "end of CoverTab[79696]"
}

// CallOption configures a Call before it starts or extracts information from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:172
// a Call after it completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:174
type CallOption interface {
	// before is called before the call is sent to any server.  If before
	// returns a non-nil error, the RPC fails with that error.
	before(*callInfo) error

	// after is called after the call has completed.  after cannot return an
	// error, so any failures should be reported via output parameters.
	after(*callInfo, *csAttempt)
}

// EmptyCallOption does not alter the Call configuration.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:184
// It can be embedded in another structure to carry satellite data for use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:184
// by interceptors.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:187
type EmptyCallOption struct{}

func (EmptyCallOption) before(*callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:189
	_go_fuzz_dep_.CoverTab[79697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:189
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:189
	// _ = "end of CoverTab[79697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:189
}
func (EmptyCallOption) after(*callInfo, *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:190
	_go_fuzz_dep_.CoverTab[79698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:190
	// _ = "end of CoverTab[79698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:190
}

// Header returns a CallOptions that retrieves the header metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:192
// for a unary RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:194
func Header(md *metadata.MD) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:194
	_go_fuzz_dep_.CoverTab[79699]++
											return HeaderCallOption{HeaderAddr: md}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:195
	// _ = "end of CoverTab[79699]"
}

// HeaderCallOption is a CallOption for collecting response header metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
// The metadata field will be populated *after* the RPC completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:198
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:205
type HeaderCallOption struct {
	HeaderAddr *metadata.MD
}

func (o HeaderCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:209
	_go_fuzz_dep_.CoverTab[79700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:209
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:209
	// _ = "end of CoverTab[79700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:209
}
func (o HeaderCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:210
	_go_fuzz_dep_.CoverTab[79701]++
											*o.HeaderAddr, _ = attempt.s.Header()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:211
	// _ = "end of CoverTab[79701]"
}

// Trailer returns a CallOptions that retrieves the trailer metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:214
// for a unary RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:216
func Trailer(md *metadata.MD) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:216
	_go_fuzz_dep_.CoverTab[79702]++
											return TrailerCallOption{TrailerAddr: md}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:217
	// _ = "end of CoverTab[79702]"
}

// TrailerCallOption is a CallOption for collecting response trailer metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
// The metadata field will be populated *after* the RPC completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:220
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:227
type TrailerCallOption struct {
	TrailerAddr *metadata.MD
}

func (o TrailerCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:231
	_go_fuzz_dep_.CoverTab[79703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:231
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:231
	// _ = "end of CoverTab[79703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:231
}
func (o TrailerCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:232
	_go_fuzz_dep_.CoverTab[79704]++
											*o.TrailerAddr = attempt.s.Trailer()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:233
	// _ = "end of CoverTab[79704]"
}

// Peer returns a CallOption that retrieves peer information for a unary RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:236
// The peer field will be populated *after* the RPC completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:238
func Peer(p *peer.Peer) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:238
	_go_fuzz_dep_.CoverTab[79705]++
											return PeerCallOption{PeerAddr: p}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:239
	// _ = "end of CoverTab[79705]"
}

// PeerCallOption is a CallOption for collecting the identity of the remote
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
// peer. The peer field will be populated *after* the RPC completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:242
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:249
type PeerCallOption struct {
	PeerAddr *peer.Peer
}

func (o PeerCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:253
	_go_fuzz_dep_.CoverTab[79706]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:253
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:253
	// _ = "end of CoverTab[79706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:253
}
func (o PeerCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:254
	_go_fuzz_dep_.CoverTab[79707]++
											if x, ok := peer.FromContext(attempt.s.Context()); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:255
		_go_fuzz_dep_.CoverTab[79708]++
												*o.PeerAddr = *x
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:256
		// _ = "end of CoverTab[79708]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:257
		_go_fuzz_dep_.CoverTab[79709]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:257
		// _ = "end of CoverTab[79709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:257
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:257
	// _ = "end of CoverTab[79707]"
}

// WaitForReady configures the action to take when an RPC is attempted on broken
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// connections or unreachable servers. If waitForReady is false and the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// connection is in the TRANSIENT_FAILURE state, the RPC will fail
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// immediately. Otherwise, the RPC client will block the call until a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// connection is available (or the call is canceled or times out) and will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// retry the call if it fails due to a transient error.  gRPC will not retry if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// data was written to the wire unless the server indicates it did not process
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// the data.  Please refer to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:260
// By default, RPCs don't "wait for ready".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:271
func WaitForReady(waitForReady bool) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:271
	_go_fuzz_dep_.CoverTab[79710]++
											return FailFastCallOption{FailFast: !waitForReady}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:272
	// _ = "end of CoverTab[79710]"
}

// FailFast is the opposite of WaitForReady.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:275
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:275
// Deprecated: use WaitForReady.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:278
func FailFast(failFast bool) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:278
	_go_fuzz_dep_.CoverTab[79711]++
											return FailFastCallOption{FailFast: failFast}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:279
	// _ = "end of CoverTab[79711]"
}

// FailFastCallOption is a CallOption for indicating whether an RPC should fail
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
// fast or not.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:282
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:289
type FailFastCallOption struct {
	FailFast bool
}

func (o FailFastCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:293
	_go_fuzz_dep_.CoverTab[79712]++
											c.failFast = o.FailFast
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:295
	// _ = "end of CoverTab[79712]"
}
func (o FailFastCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:297
	_go_fuzz_dep_.CoverTab[79713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:297
	// _ = "end of CoverTab[79713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:297
}

// OnFinish returns a CallOption that configures a callback to be called when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// the call completes. The error passed to the callback is the status of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// RPC, and may be nil. The onFinish callback provided will only be called once
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// by gRPC. This is mainly used to be used by streaming interceptors, to be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// notified when the RPC completes along with information about the status of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// the RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:299
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:310
func OnFinish(onFinish func(err error)) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:310
	_go_fuzz_dep_.CoverTab[79714]++
											return OnFinishCallOption{
		OnFinish: onFinish,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:313
	// _ = "end of CoverTab[79714]"
}

// OnFinishCallOption is CallOption that indicates a callback to be called when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
// the call completes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:316
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:323
type OnFinishCallOption struct {
	OnFinish func(error)
}

func (o OnFinishCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:327
	_go_fuzz_dep_.CoverTab[79715]++
											c.onFinish = append(c.onFinish, o.OnFinish)
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:329
	// _ = "end of CoverTab[79715]"
}

func (o OnFinishCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:332
	_go_fuzz_dep_.CoverTab[79716]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:332
	// _ = "end of CoverTab[79716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:332
}

// MaxCallRecvMsgSize returns a CallOption which sets the maximum message size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:334
// in bytes the client can receive. If this is not set, gRPC uses the default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:334
// 4MB.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:337
func MaxCallRecvMsgSize(bytes int) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:337
	_go_fuzz_dep_.CoverTab[79717]++
											return MaxRecvMsgSizeCallOption{MaxRecvMsgSize: bytes}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:338
	// _ = "end of CoverTab[79717]"
}

// MaxRecvMsgSizeCallOption is a CallOption that indicates the maximum message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
// size in bytes the client can receive.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:341
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:348
type MaxRecvMsgSizeCallOption struct {
	MaxRecvMsgSize int
}

func (o MaxRecvMsgSizeCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:352
	_go_fuzz_dep_.CoverTab[79718]++
											c.maxReceiveMessageSize = &o.MaxRecvMsgSize
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:354
	// _ = "end of CoverTab[79718]"
}
func (o MaxRecvMsgSizeCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:356
	_go_fuzz_dep_.CoverTab[79719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:356
	// _ = "end of CoverTab[79719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:356
}

// MaxCallSendMsgSize returns a CallOption which sets the maximum message size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:358
// in bytes the client can send. If this is not set, gRPC uses the default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:358
// `math.MaxInt32`.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:361
func MaxCallSendMsgSize(bytes int) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:361
	_go_fuzz_dep_.CoverTab[79720]++
											return MaxSendMsgSizeCallOption{MaxSendMsgSize: bytes}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:362
	// _ = "end of CoverTab[79720]"
}

// MaxSendMsgSizeCallOption is a CallOption that indicates the maximum message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
// size in bytes the client can send.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:365
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:372
type MaxSendMsgSizeCallOption struct {
	MaxSendMsgSize int
}

func (o MaxSendMsgSizeCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:376
	_go_fuzz_dep_.CoverTab[79721]++
											c.maxSendMessageSize = &o.MaxSendMsgSize
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:378
	// _ = "end of CoverTab[79721]"
}
func (o MaxSendMsgSizeCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:380
	_go_fuzz_dep_.CoverTab[79722]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:380
	// _ = "end of CoverTab[79722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:380
}

// PerRPCCredentials returns a CallOption that sets credentials.PerRPCCredentials
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:382
// for a call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:384
func PerRPCCredentials(creds credentials.PerRPCCredentials) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:384
	_go_fuzz_dep_.CoverTab[79723]++
											return PerRPCCredsCallOption{Creds: creds}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:385
	// _ = "end of CoverTab[79723]"
}

// PerRPCCredsCallOption is a CallOption that indicates the per-RPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
// credentials to use for the call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:388
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:395
type PerRPCCredsCallOption struct {
	Creds credentials.PerRPCCredentials
}

func (o PerRPCCredsCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:399
	_go_fuzz_dep_.CoverTab[79724]++
											c.creds = o.Creds
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:401
	// _ = "end of CoverTab[79724]"
}
func (o PerRPCCredsCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:403
	_go_fuzz_dep_.CoverTab[79725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:403
	// _ = "end of CoverTab[79725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:403
}

// UseCompressor returns a CallOption which sets the compressor used when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
// sending the request.  If WithCompressor is also set, UseCompressor has
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
// higher priority.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:405
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:413
func UseCompressor(name string) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:413
	_go_fuzz_dep_.CoverTab[79726]++
											return CompressorCallOption{CompressorType: name}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:414
	// _ = "end of CoverTab[79726]"
}

// CompressorCallOption is a CallOption that indicates the compressor to use.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:417
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:417
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:417
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:417
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:417
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:423
type CompressorCallOption struct {
	CompressorType string
}

func (o CompressorCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:427
	_go_fuzz_dep_.CoverTab[79727]++
											c.compressorType = o.CompressorType
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:429
	// _ = "end of CoverTab[79727]"
}
func (o CompressorCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:431
	_go_fuzz_dep_.CoverTab[79728]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:431
	// _ = "end of CoverTab[79728]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:431
}

// CallContentSubtype returns a CallOption that will set the content-subtype
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// for a call. For example, if content-subtype is "json", the Content-Type over
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// the wire will be "application/grpc+json". The content-subtype is converted
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// to lowercase before being included in Content-Type. See Content-Type on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// If ForceCodec is not also used, the content-subtype will be used to look up
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// the Codec to use in the registry controlled by RegisterCodec. See the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// documentation on RegisterCodec for details on registration. The lookup of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// content-subtype is case-insensitive. If no such Codec is found, the call
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// will result in an error with code codes.Internal.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// If ForceCodec is also used, that Codec will be used for all request and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// response messages, with the content-subtype set to the given contentSubtype
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:433
// here for requests.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:449
func CallContentSubtype(contentSubtype string) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:449
	_go_fuzz_dep_.CoverTab[79729]++
											return ContentSubtypeCallOption{ContentSubtype: strings.ToLower(contentSubtype)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:450
	// _ = "end of CoverTab[79729]"
}

// ContentSubtypeCallOption is a CallOption that indicates the content-subtype
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
// used for marshaling messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:453
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:460
type ContentSubtypeCallOption struct {
	ContentSubtype string
}

func (o ContentSubtypeCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:464
	_go_fuzz_dep_.CoverTab[79730]++
											c.contentSubtype = o.ContentSubtype
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:466
	// _ = "end of CoverTab[79730]"
}
func (o ContentSubtypeCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:468
	_go_fuzz_dep_.CoverTab[79731]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:468
	// _ = "end of CoverTab[79731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:468
}

// ForceCodec returns a CallOption that will set codec to be used for all
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// request and response messages for a call. The result of calling Name() will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// be used as the content-subtype after converting to lowercase, unless
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// CallContentSubtype is also used.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// See Content-Type on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// more details. Also see the documentation on RegisterCodec and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// CallContentSubtype for more details on the interaction between Codec and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// content-subtype.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// This function is provided for advanced users; prefer to use only
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// CallContentSubtype to select a registered codec instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:470
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:488
func ForceCodec(codec encoding.Codec) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:488
	_go_fuzz_dep_.CoverTab[79732]++
											return ForceCodecCallOption{Codec: codec}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:489
	// _ = "end of CoverTab[79732]"
}

// ForceCodecCallOption is a CallOption that indicates the codec used for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
// marshaling messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:492
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:499
type ForceCodecCallOption struct {
	Codec encoding.Codec
}

func (o ForceCodecCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:503
	_go_fuzz_dep_.CoverTab[79733]++
											c.codec = o.Codec
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:505
	// _ = "end of CoverTab[79733]"
}
func (o ForceCodecCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:507
	_go_fuzz_dep_.CoverTab[79734]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:507
	// _ = "end of CoverTab[79734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:507
}

// CallCustomCodec behaves like ForceCodec, but accepts a grpc.Codec instead of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:509
// an encoding.Codec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:509
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:509
// Deprecated: use ForceCodec instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:513
func CallCustomCodec(codec Codec) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:513
	_go_fuzz_dep_.CoverTab[79735]++
											return CustomCodecCallOption{Codec: codec}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:514
	// _ = "end of CoverTab[79735]"
}

// CustomCodecCallOption is a CallOption that indicates the codec used for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
// marshaling messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:517
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:524
type CustomCodecCallOption struct {
	Codec Codec
}

func (o CustomCodecCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:528
	_go_fuzz_dep_.CoverTab[79736]++
											c.codec = o.Codec
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:530
	// _ = "end of CoverTab[79736]"
}
func (o CustomCodecCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:532
	_go_fuzz_dep_.CoverTab[79737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:532
	// _ = "end of CoverTab[79737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:532
}

// MaxRetryRPCBufferSize returns a CallOption that limits the amount of memory
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
// used for buffering this RPC's requests for retry purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:534
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:541
func MaxRetryRPCBufferSize(bytes int) CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:541
	_go_fuzz_dep_.CoverTab[79738]++
											return MaxRetryRPCBufferSizeCallOption{bytes}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:542
	// _ = "end of CoverTab[79738]"
}

// MaxRetryRPCBufferSizeCallOption is a CallOption indicating the amount of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
// memory to be used for caching this RPC for retry purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:545
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:552
type MaxRetryRPCBufferSizeCallOption struct {
	MaxRetryRPCBufferSize int
}

func (o MaxRetryRPCBufferSizeCallOption) before(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:556
	_go_fuzz_dep_.CoverTab[79739]++
											c.maxRetryRPCBufferSize = o.MaxRetryRPCBufferSize
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:558
	// _ = "end of CoverTab[79739]"
}
func (o MaxRetryRPCBufferSizeCallOption) after(c *callInfo, attempt *csAttempt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:560
	_go_fuzz_dep_.CoverTab[79740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:560
	// _ = "end of CoverTab[79740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:560
}

// The format of the payload: compressed or not?
type payloadFormat uint8

const (
	compressionNone	payloadFormat	= 0	// no compression
	compressionMade	payloadFormat	= 1	// compressed
)

// parser reads complete gRPC messages from the underlying reader.
type parser struct {
	// r is the underlying reader.
	// See the comment on recvMsg for the permissible
	// error types.
	r	io.Reader

	// The header of a gRPC message. Find more detail at
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
	header	[5]byte
}

// recvMsg reads a complete gRPC message from the stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// It returns the message and its payload (compression/encoding)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// format. The caller owns the returned msg memory.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// If there is an error, possible values are:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//   - io.EOF, when no messages remain
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//   - io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//   - of type transport.ConnectionError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//   - an error from the status package
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// No other error values or types must be returned, which also means
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// that the underlying io.Reader must not return an incompatible
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:582
// error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:596
func (p *parser) recvMsg(maxReceiveMessageSize int) (pf payloadFormat, msg []byte, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:596
	_go_fuzz_dep_.CoverTab[79741]++
											if _, err := p.r.Read(p.header[:]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:597
		_go_fuzz_dep_.CoverTab[79747]++
												return 0, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:598
		// _ = "end of CoverTab[79747]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:599
		_go_fuzz_dep_.CoverTab[79748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:599
		// _ = "end of CoverTab[79748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:599
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:599
	// _ = "end of CoverTab[79741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:599
	_go_fuzz_dep_.CoverTab[79742]++

											pf = payloadFormat(p.header[0])
											length := binary.BigEndian.Uint32(p.header[1:])

											if length == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:604
		_go_fuzz_dep_.CoverTab[79749]++
												return pf, nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:605
		// _ = "end of CoverTab[79749]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:606
		_go_fuzz_dep_.CoverTab[79750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:606
		// _ = "end of CoverTab[79750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:606
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:606
	// _ = "end of CoverTab[79742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:606
	_go_fuzz_dep_.CoverTab[79743]++
											if int64(length) > int64(maxInt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:607
		_go_fuzz_dep_.CoverTab[79751]++
												return 0, nil, status.Errorf(codes.ResourceExhausted, "grpc: received message larger than max length allowed on current machine (%d vs. %d)", length, maxInt)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:608
		// _ = "end of CoverTab[79751]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:609
		_go_fuzz_dep_.CoverTab[79752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:609
		// _ = "end of CoverTab[79752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:609
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:609
	// _ = "end of CoverTab[79743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:609
	_go_fuzz_dep_.CoverTab[79744]++
											if int(length) > maxReceiveMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:610
		_go_fuzz_dep_.CoverTab[79753]++
												return 0, nil, status.Errorf(codes.ResourceExhausted, "grpc: received message larger than max (%d vs. %d)", length, maxReceiveMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:611
		// _ = "end of CoverTab[79753]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:612
		_go_fuzz_dep_.CoverTab[79754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:612
		// _ = "end of CoverTab[79754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:612
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:612
	// _ = "end of CoverTab[79744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:612
	_go_fuzz_dep_.CoverTab[79745]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:615
	msg = make([]byte, int(length))
	if _, err := p.r.Read(msg); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:616
		_go_fuzz_dep_.CoverTab[79755]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:617
			_go_fuzz_dep_.CoverTab[79757]++
													err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:618
			// _ = "end of CoverTab[79757]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:619
			_go_fuzz_dep_.CoverTab[79758]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:619
			// _ = "end of CoverTab[79758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:619
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:619
		// _ = "end of CoverTab[79755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:619
		_go_fuzz_dep_.CoverTab[79756]++
												return 0, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:620
		// _ = "end of CoverTab[79756]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:621
		_go_fuzz_dep_.CoverTab[79759]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:621
		// _ = "end of CoverTab[79759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:621
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:621
	// _ = "end of CoverTab[79745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:621
	_go_fuzz_dep_.CoverTab[79746]++
											return pf, msg, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:622
	// _ = "end of CoverTab[79746]"
}

// encode serializes msg and returns a buffer containing the message, or an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:625
// error if it is too large to be transmitted by grpc.  If msg is nil, it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:625
// generates an empty message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:628
func encode(c baseCodec, msg interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:628
	_go_fuzz_dep_.CoverTab[79760]++
											if msg == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:629
		_go_fuzz_dep_.CoverTab[79764]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:630
		// _ = "end of CoverTab[79764]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:631
		_go_fuzz_dep_.CoverTab[79765]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:631
		// _ = "end of CoverTab[79765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:631
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:631
	// _ = "end of CoverTab[79760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:631
	_go_fuzz_dep_.CoverTab[79761]++
											b, err := c.Marshal(msg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:633
		_go_fuzz_dep_.CoverTab[79766]++
												return nil, status.Errorf(codes.Internal, "grpc: error while marshaling: %v", err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:634
		// _ = "end of CoverTab[79766]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:635
		_go_fuzz_dep_.CoverTab[79767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:635
		// _ = "end of CoverTab[79767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:635
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:635
	// _ = "end of CoverTab[79761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:635
	_go_fuzz_dep_.CoverTab[79762]++
											if uint(len(b)) > math.MaxUint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:636
		_go_fuzz_dep_.CoverTab[79768]++
												return nil, status.Errorf(codes.ResourceExhausted, "grpc: message too large (%d bytes)", len(b))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:637
		// _ = "end of CoverTab[79768]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:638
		_go_fuzz_dep_.CoverTab[79769]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:638
		// _ = "end of CoverTab[79769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:638
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:638
	// _ = "end of CoverTab[79762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:638
	_go_fuzz_dep_.CoverTab[79763]++
											return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:639
	// _ = "end of CoverTab[79763]"
}

// compress returns the input bytes compressed by compressor or cp.  If both
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:642
// compressors are nil, returns nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:642
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:642
// TODO(dfawley): eliminate cp parameter by wrapping Compressor in an encoding.Compressor.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:646
func compress(in []byte, cp Compressor, compressor encoding.Compressor) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:646
	_go_fuzz_dep_.CoverTab[79770]++
											if compressor == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:647
		_go_fuzz_dep_.CoverTab[79774]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:647
		return cp == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:647
		// _ = "end of CoverTab[79774]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:647
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:647
		_go_fuzz_dep_.CoverTab[79775]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:648
		// _ = "end of CoverTab[79775]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:649
		_go_fuzz_dep_.CoverTab[79776]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:649
		// _ = "end of CoverTab[79776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:649
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:649
	// _ = "end of CoverTab[79770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:649
	_go_fuzz_dep_.CoverTab[79771]++
											wrapErr := func(err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:650
		_go_fuzz_dep_.CoverTab[79777]++
												return status.Errorf(codes.Internal, "grpc: error while compressing: %v", err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:651
		// _ = "end of CoverTab[79777]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:652
	// _ = "end of CoverTab[79771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:652
	_go_fuzz_dep_.CoverTab[79772]++
											cbuf := &bytes.Buffer{}
											if compressor != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:654
		_go_fuzz_dep_.CoverTab[79778]++
												z, err := compressor.Compress(cbuf)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:656
			_go_fuzz_dep_.CoverTab[79781]++
													return nil, wrapErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:657
			// _ = "end of CoverTab[79781]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:658
			_go_fuzz_dep_.CoverTab[79782]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:658
			// _ = "end of CoverTab[79782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:658
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:658
		// _ = "end of CoverTab[79778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:658
		_go_fuzz_dep_.CoverTab[79779]++
												if _, err := z.Write(in); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:659
			_go_fuzz_dep_.CoverTab[79783]++
													return nil, wrapErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:660
			// _ = "end of CoverTab[79783]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:661
			_go_fuzz_dep_.CoverTab[79784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:661
			// _ = "end of CoverTab[79784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:661
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:661
		// _ = "end of CoverTab[79779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:661
		_go_fuzz_dep_.CoverTab[79780]++
												if err := z.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:662
			_go_fuzz_dep_.CoverTab[79785]++
													return nil, wrapErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:663
			// _ = "end of CoverTab[79785]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:664
			_go_fuzz_dep_.CoverTab[79786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:664
			// _ = "end of CoverTab[79786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:664
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:664
		// _ = "end of CoverTab[79780]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:665
		_go_fuzz_dep_.CoverTab[79787]++
												if err := cp.Do(cbuf, in); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:666
			_go_fuzz_dep_.CoverTab[79788]++
													return nil, wrapErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:667
			// _ = "end of CoverTab[79788]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:668
			_go_fuzz_dep_.CoverTab[79789]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:668
			// _ = "end of CoverTab[79789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:668
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:668
		// _ = "end of CoverTab[79787]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:669
	// _ = "end of CoverTab[79772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:669
	_go_fuzz_dep_.CoverTab[79773]++
											return cbuf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:670
	// _ = "end of CoverTab[79773]"
}

const (
	payloadLen	= 1
	sizeLen		= 4
	headerLen	= payloadLen + sizeLen
)

// msgHeader returns a 5-byte header for the message being transmitted and the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:679
// payload, which is compData if non-nil or data otherwise.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:681
func msgHeader(data, compData []byte) (hdr []byte, payload []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:681
	_go_fuzz_dep_.CoverTab[79790]++
											hdr = make([]byte, headerLen)
											if compData != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:683
		_go_fuzz_dep_.CoverTab[79792]++
												hdr[0] = byte(compressionMade)
												data = compData
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:685
		// _ = "end of CoverTab[79792]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:686
		_go_fuzz_dep_.CoverTab[79793]++
												hdr[0] = byte(compressionNone)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:687
		// _ = "end of CoverTab[79793]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:688
	// _ = "end of CoverTab[79790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:688
	_go_fuzz_dep_.CoverTab[79791]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:691
	binary.BigEndian.PutUint32(hdr[payloadLen:], uint32(len(data)))
											return hdr, data
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:692
	// _ = "end of CoverTab[79791]"
}

func outPayload(client bool, msg interface{}, data, payload []byte, t time.Time) *stats.OutPayload {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:695
	_go_fuzz_dep_.CoverTab[79794]++
											return &stats.OutPayload{
		Client:			client,
		Payload:		msg,
		Data:			data,
		Length:			len(data),
		WireLength:		len(payload) + headerLen,
		CompressedLength:	len(payload),
		SentTime:		t,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:704
	// _ = "end of CoverTab[79794]"
}

func checkRecvPayload(pf payloadFormat, recvCompress string, haveCompressor bool) *status.Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:707
	_go_fuzz_dep_.CoverTab[79795]++
											switch pf {
	case compressionNone:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:709
		_go_fuzz_dep_.CoverTab[79797]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:709
		// _ = "end of CoverTab[79797]"
	case compressionMade:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:710
		_go_fuzz_dep_.CoverTab[79798]++
												if recvCompress == "" || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:711
			_go_fuzz_dep_.CoverTab[79801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:711
			return recvCompress == encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:711
			// _ = "end of CoverTab[79801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:711
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:711
			_go_fuzz_dep_.CoverTab[79802]++
													return status.New(codes.Internal, "grpc: compressed flag set with identity or empty encoding")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:712
			// _ = "end of CoverTab[79802]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:713
			_go_fuzz_dep_.CoverTab[79803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:713
			// _ = "end of CoverTab[79803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:713
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:713
		// _ = "end of CoverTab[79798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:713
		_go_fuzz_dep_.CoverTab[79799]++
												if !haveCompressor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:714
			_go_fuzz_dep_.CoverTab[79804]++
													return status.Newf(codes.Unimplemented, "grpc: Decompressor is not installed for grpc-encoding %q", recvCompress)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:715
			// _ = "end of CoverTab[79804]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:716
			_go_fuzz_dep_.CoverTab[79805]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:716
			// _ = "end of CoverTab[79805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:716
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:716
		// _ = "end of CoverTab[79799]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:717
		_go_fuzz_dep_.CoverTab[79800]++
												return status.Newf(codes.Internal, "grpc: received unexpected payload format %d", pf)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:718
		// _ = "end of CoverTab[79800]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:719
	// _ = "end of CoverTab[79795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:719
	_go_fuzz_dep_.CoverTab[79796]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:720
	// _ = "end of CoverTab[79796]"
}

type payloadInfo struct {
	compressedLength	int	// The compressed length got from wire.
	uncompressedBytes	[]byte
}

func recvAndDecompress(p *parser, s *transport.Stream, dc Decompressor, maxReceiveMessageSize int, payInfo *payloadInfo, compressor encoding.Compressor) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:728
	_go_fuzz_dep_.CoverTab[79806]++
											pf, d, err := p.recvMsg(maxReceiveMessageSize)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:730
		_go_fuzz_dep_.CoverTab[79811]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:731
		// _ = "end of CoverTab[79811]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:732
		_go_fuzz_dep_.CoverTab[79812]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:732
		// _ = "end of CoverTab[79812]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:732
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:732
	// _ = "end of CoverTab[79806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:732
	_go_fuzz_dep_.CoverTab[79807]++
											if payInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:733
		_go_fuzz_dep_.CoverTab[79813]++
												payInfo.compressedLength = len(d)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:734
		// _ = "end of CoverTab[79813]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:735
		_go_fuzz_dep_.CoverTab[79814]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:735
		// _ = "end of CoverTab[79814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:735
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:735
	// _ = "end of CoverTab[79807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:735
	_go_fuzz_dep_.CoverTab[79808]++

											if st := checkRecvPayload(pf, s.RecvCompress(), compressor != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:737
		_go_fuzz_dep_.CoverTab[79815]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:737
		return dc != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:737
		// _ = "end of CoverTab[79815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:737
	}()); st != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:737
		_go_fuzz_dep_.CoverTab[79816]++
												return nil, st.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:738
		// _ = "end of CoverTab[79816]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:739
		_go_fuzz_dep_.CoverTab[79817]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:739
		// _ = "end of CoverTab[79817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:739
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:739
	// _ = "end of CoverTab[79808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:739
	_go_fuzz_dep_.CoverTab[79809]++

											var size int
											if pf == compressionMade {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:742
		_go_fuzz_dep_.CoverTab[79818]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:745
		if dc != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:745
			_go_fuzz_dep_.CoverTab[79821]++
													d, err = dc.Do(bytes.NewReader(d))
													size = len(d)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:747
			// _ = "end of CoverTab[79821]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:748
			_go_fuzz_dep_.CoverTab[79822]++
													d, size, err = decompress(compressor, d, maxReceiveMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:749
			// _ = "end of CoverTab[79822]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:750
		// _ = "end of CoverTab[79818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:750
		_go_fuzz_dep_.CoverTab[79819]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:751
			_go_fuzz_dep_.CoverTab[79823]++
													return nil, status.Errorf(codes.Internal, "grpc: failed to decompress the received message: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:752
			// _ = "end of CoverTab[79823]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:753
			_go_fuzz_dep_.CoverTab[79824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:753
			// _ = "end of CoverTab[79824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:753
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:753
		// _ = "end of CoverTab[79819]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:753
		_go_fuzz_dep_.CoverTab[79820]++
												if size > maxReceiveMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:754
			_go_fuzz_dep_.CoverTab[79825]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:757
			return nil, status.Errorf(codes.ResourceExhausted, "grpc: received message after decompression larger than max (%d vs. %d)", size, maxReceiveMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:757
			// _ = "end of CoverTab[79825]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:758
			_go_fuzz_dep_.CoverTab[79826]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:758
			// _ = "end of CoverTab[79826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:758
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:758
		// _ = "end of CoverTab[79820]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:759
		_go_fuzz_dep_.CoverTab[79827]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:759
		// _ = "end of CoverTab[79827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:759
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:759
	// _ = "end of CoverTab[79809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:759
	_go_fuzz_dep_.CoverTab[79810]++
											return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:760
	// _ = "end of CoverTab[79810]"
}

// Using compressor, decompress d, returning data and size.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:763
// Optionally, if data will be over maxReceiveMessageSize, just return the size.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:765
func decompress(compressor encoding.Compressor, d []byte, maxReceiveMessageSize int) ([]byte, int, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:765
	_go_fuzz_dep_.CoverTab[79828]++
											dcReader, err := compressor.Decompress(bytes.NewReader(d))
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:767
		_go_fuzz_dep_.CoverTab[79831]++
												return nil, 0, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:768
		// _ = "end of CoverTab[79831]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:769
		_go_fuzz_dep_.CoverTab[79832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:769
		// _ = "end of CoverTab[79832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:769
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:769
	// _ = "end of CoverTab[79828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:769
	_go_fuzz_dep_.CoverTab[79829]++
											if sizer, ok := compressor.(interface {
		DecompressedSize(compressedBytes []byte) int
	}); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:772
		_go_fuzz_dep_.CoverTab[79833]++
												if size := sizer.DecompressedSize(d); size >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:773
			_go_fuzz_dep_.CoverTab[79834]++
													if size > maxReceiveMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:774
				_go_fuzz_dep_.CoverTab[79836]++
														return nil, size, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:775
				// _ = "end of CoverTab[79836]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:776
				_go_fuzz_dep_.CoverTab[79837]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:776
				// _ = "end of CoverTab[79837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:776
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:776
			// _ = "end of CoverTab[79834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:776
			_go_fuzz_dep_.CoverTab[79835]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:780
			buf := bytes.NewBuffer(make([]byte, 0, size+bytes.MinRead))
													bytesRead, err := buf.ReadFrom(io.LimitReader(dcReader, int64(maxReceiveMessageSize)+1))
													return buf.Bytes(), int(bytesRead), err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:782
			// _ = "end of CoverTab[79835]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:783
			_go_fuzz_dep_.CoverTab[79838]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:783
			// _ = "end of CoverTab[79838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:783
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:783
		// _ = "end of CoverTab[79833]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:784
		_go_fuzz_dep_.CoverTab[79839]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:784
		// _ = "end of CoverTab[79839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:784
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:784
	// _ = "end of CoverTab[79829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:784
	_go_fuzz_dep_.CoverTab[79830]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:787
	d, err = io.ReadAll(io.LimitReader(dcReader, int64(maxReceiveMessageSize)+1))
											return d, len(d), err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:788
	// _ = "end of CoverTab[79830]"
}

// For the two compressor parameters, both should not be set, but if they are,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:791
// dc takes precedence over compressor.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:791
// TODO(dfawley): wrap the old compressor/decompressor using the new API?
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:794
func recv(p *parser, c baseCodec, s *transport.Stream, dc Decompressor, m interface{}, maxReceiveMessageSize int, payInfo *payloadInfo, compressor encoding.Compressor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:794
	_go_fuzz_dep_.CoverTab[79840]++
											d, err := recvAndDecompress(p, s, dc, maxReceiveMessageSize, payInfo, compressor)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:796
		_go_fuzz_dep_.CoverTab[79844]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:797
		// _ = "end of CoverTab[79844]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:798
		_go_fuzz_dep_.CoverTab[79845]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:798
		// _ = "end of CoverTab[79845]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:798
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:798
	// _ = "end of CoverTab[79840]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:798
	_go_fuzz_dep_.CoverTab[79841]++
											if err := c.Unmarshal(d, m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:799
		_go_fuzz_dep_.CoverTab[79846]++
												return status.Errorf(codes.Internal, "grpc: failed to unmarshal the received message: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:800
		// _ = "end of CoverTab[79846]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:801
		_go_fuzz_dep_.CoverTab[79847]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:801
		// _ = "end of CoverTab[79847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:801
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:801
	// _ = "end of CoverTab[79841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:801
	_go_fuzz_dep_.CoverTab[79842]++
											if payInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:802
		_go_fuzz_dep_.CoverTab[79848]++
												payInfo.uncompressedBytes = d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:803
		// _ = "end of CoverTab[79848]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:804
		_go_fuzz_dep_.CoverTab[79849]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:804
		// _ = "end of CoverTab[79849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:804
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:804
	// _ = "end of CoverTab[79842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:804
	_go_fuzz_dep_.CoverTab[79843]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:805
	// _ = "end of CoverTab[79843]"
}

// Information about RPC
type rpcInfo struct {
	failfast	bool
	preloaderInfo	*compressorInfo
}

// Information about Preloader
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:814
// Responsible for storing codec, and compressors
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:814
// If stream (s) has  context s.Context which stores rpcInfo that has non nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:814
// pointers to codec, and compressors, then we can use preparedMsg for Async message prep
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:814
// and reuse marshalled bytes
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:819
type compressorInfo struct {
	codec	baseCodec
	cp	Compressor
	comp	encoding.Compressor
}

type rpcInfoContextKey struct{}

func newContextWithRPCInfo(ctx context.Context, failfast bool, codec baseCodec, cp Compressor, comp encoding.Compressor) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:827
	_go_fuzz_dep_.CoverTab[79850]++
											return context.WithValue(ctx, rpcInfoContextKey{}, &rpcInfo{
		failfast:	failfast,
		preloaderInfo: &compressorInfo{
			codec:	codec,
			cp:	cp,
			comp:	comp,
		},
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:835
	// _ = "end of CoverTab[79850]"
}

func rpcInfoFromContext(ctx context.Context) (s *rpcInfo, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:838
	_go_fuzz_dep_.CoverTab[79851]++
											s, ok = ctx.Value(rpcInfoContextKey{}).(*rpcInfo)
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:840
	// _ = "end of CoverTab[79851]"
}

// Code returns the error code for err if it was produced by the rpc system.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:843
// Otherwise, it returns codes.Unknown.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:843
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:843
// Deprecated: use status.Code instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:847
func Code(err error) codes.Code {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:847
	_go_fuzz_dep_.CoverTab[79852]++
											return status.Code(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:848
	// _ = "end of CoverTab[79852]"
}

// ErrorDesc returns the error description of err if it was produced by the rpc system.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:851
// Otherwise, it returns err.Error() or empty string when err is nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:851
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:851
// Deprecated: use status.Convert and Message method instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:855
func ErrorDesc(err error) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:855
	_go_fuzz_dep_.CoverTab[79853]++
											return status.Convert(err).Message()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:856
	// _ = "end of CoverTab[79853]"
}

// Errorf returns an error containing an error code and a description;
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:859
// Errorf returns nil if c is OK.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:859
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:859
// Deprecated: use status.Errorf instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:863
func Errorf(c codes.Code, format string, a ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:863
	_go_fuzz_dep_.CoverTab[79854]++
											return status.Errorf(c, format, a...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:864
	// _ = "end of CoverTab[79854]"
}

// toRPCErr converts an error into an error from the status package.
func toRPCErr(err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:868
	_go_fuzz_dep_.CoverTab[79855]++
											switch err {
	case nil, io.EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:870
		_go_fuzz_dep_.CoverTab[79859]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:871
		// _ = "end of CoverTab[79859]"
	case context.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:872
		_go_fuzz_dep_.CoverTab[79860]++
												return status.Error(codes.DeadlineExceeded, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:873
		// _ = "end of CoverTab[79860]"
	case context.Canceled:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:874
		_go_fuzz_dep_.CoverTab[79861]++
												return status.Error(codes.Canceled, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:875
		// _ = "end of CoverTab[79861]"
	case io.ErrUnexpectedEOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:876
		_go_fuzz_dep_.CoverTab[79862]++
												return status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:877
		// _ = "end of CoverTab[79862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:877
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:877
		_go_fuzz_dep_.CoverTab[79863]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:877
		// _ = "end of CoverTab[79863]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:878
	// _ = "end of CoverTab[79855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:878
	_go_fuzz_dep_.CoverTab[79856]++

											switch e := err.(type) {
	case transport.ConnectionError:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:881
		_go_fuzz_dep_.CoverTab[79864]++
												return status.Error(codes.Unavailable, e.Desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:882
		// _ = "end of CoverTab[79864]"
	case *transport.NewStreamError:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:883
		_go_fuzz_dep_.CoverTab[79865]++
												return toRPCErr(e.Err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:884
		// _ = "end of CoverTab[79865]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:885
	// _ = "end of CoverTab[79856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:885
	_go_fuzz_dep_.CoverTab[79857]++

											if _, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:887
		_go_fuzz_dep_.CoverTab[79866]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:888
		// _ = "end of CoverTab[79866]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:889
		_go_fuzz_dep_.CoverTab[79867]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:889
		// _ = "end of CoverTab[79867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:889
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:889
	// _ = "end of CoverTab[79857]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:889
	_go_fuzz_dep_.CoverTab[79858]++

											return status.Error(codes.Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:891
	// _ = "end of CoverTab[79858]"
}

// setCallInfoCodec should only be called after CallOptions have been applied.
func setCallInfoCodec(c *callInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:895
	_go_fuzz_dep_.CoverTab[79868]++
											if c.codec != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:896
		_go_fuzz_dep_.CoverTab[79872]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:899
		if c.contentSubtype == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:899
			_go_fuzz_dep_.CoverTab[79874]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:904
			if ec, ok := c.codec.(encoding.Codec); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:904
				_go_fuzz_dep_.CoverTab[79875]++
														c.contentSubtype = strings.ToLower(ec.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:905
				// _ = "end of CoverTab[79875]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:906
				_go_fuzz_dep_.CoverTab[79876]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:906
				// _ = "end of CoverTab[79876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:906
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:906
			// _ = "end of CoverTab[79874]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:907
			_go_fuzz_dep_.CoverTab[79877]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:907
			// _ = "end of CoverTab[79877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:907
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:907
		// _ = "end of CoverTab[79872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:907
		_go_fuzz_dep_.CoverTab[79873]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:908
		// _ = "end of CoverTab[79873]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:909
		_go_fuzz_dep_.CoverTab[79878]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:909
		// _ = "end of CoverTab[79878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:909
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:909
	// _ = "end of CoverTab[79868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:909
	_go_fuzz_dep_.CoverTab[79869]++

											if c.contentSubtype == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:911
		_go_fuzz_dep_.CoverTab[79879]++

												c.codec = encoding.GetCodec(proto.Name)
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:914
		// _ = "end of CoverTab[79879]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:915
		_go_fuzz_dep_.CoverTab[79880]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:915
		// _ = "end of CoverTab[79880]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:915
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:915
	// _ = "end of CoverTab[79869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:915
	_go_fuzz_dep_.CoverTab[79870]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:918
	c.codec = encoding.GetCodec(c.contentSubtype)
	if c.codec == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:919
		_go_fuzz_dep_.CoverTab[79881]++
												return status.Errorf(codes.Internal, "no codec registered for content-subtype %s", c.contentSubtype)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:920
		// _ = "end of CoverTab[79881]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:921
		_go_fuzz_dep_.CoverTab[79882]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:921
		// _ = "end of CoverTab[79882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:921
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:921
	// _ = "end of CoverTab[79870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:921
	_go_fuzz_dep_.CoverTab[79871]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:922
	// _ = "end of CoverTab[79871]"
}

// channelzData is used to store channelz related data for ClientConn, addrConn and Server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:925
// These fields cannot be embedded in the original structs (e.g. ClientConn), since to do atomic
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:925
// operation on int64 variable on 32-bit machine, user is responsible to enforce memory alignment.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:925
// Here, by grouping those int64 fields inside a struct, we are enforcing the alignment.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:929
type channelzData struct {
	callsStarted	int64
	callsFailed	int64
	callsSucceeded	int64
	// lastCallStartedTime stores the timestamp that last call starts. It is of int64 type instead of
	// time.Time since it's more costly to atomically update time.Time variable than int64 variable.
	lastCallStartedTime	int64
}

// The SupportPackageIsVersion variables are referenced from generated protocol
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
// buffer files to ensure compatibility with the gRPC version used.  The latest
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
// support package version is 7.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
// Older versions are kept for compatibility.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:938
// These constants should not be referenced from any other code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:945
const (
	SupportPackageIsVersion3	= true
	SupportPackageIsVersion4	= true
	SupportPackageIsVersion5	= true
	SupportPackageIsVersion6	= true
	SupportPackageIsVersion7	= true
)

const grpcUA = "grpc-go/" + Version

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:953
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/rpc_util.go:953
var _ = _go_fuzz_dep_.CoverTab
