//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
// Package encoding defines the interface for the compressor and codec, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
// functions to register and retrieve compressors and codecs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:19
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
package encoding

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:26
)

import (
	"io"
	"strings"

	"google.golang.org/grpc/internal/grpcutil"
)

// Identity specifies the optional encoding for uncompressed streams.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:35
// It is intended for grpc internal use only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:37
const Identity = "identity"

// Compressor is used for compressing and decompressing when sending or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:39
// receiving messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:41
type Compressor interface {
	// Compress writes the data written to wc to w after compressing it.  If an
	// error occurs while initializing the compressor, that error is returned
	// instead.
	Compress(w io.Writer) (io.WriteCloser, error)
	// Decompress reads data from r, decompresses it, and provides the
	// uncompressed data via the returned io.Reader.  If an error occurs while
	// initializing the decompressor, that error is returned instead.
	Decompress(r io.Reader) (io.Reader, error)
	// Name is the name of the compression codec and is used to set the content
	// coding header.  The result must be static; the result cannot change
	// between calls.
	Name() string
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:63
}

var registeredCompressor = make(map[string]Compressor)

// RegisterCompressor registers the compressor with gRPC by its name.  It can
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// be activated when sending an RPC via grpc.UseCompressor().  It will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// automatically accessed when receiving a message based on the content coding
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// header.  Servers also use it to send a response with the same encoding as
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// the request.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// NOTE: this function must only be called during initialization time (i.e. in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// an init() function), and is not thread-safe.  If multiple Compressors are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:67
// registered with the same name, the one registered last will take effect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:76
func RegisterCompressor(c Compressor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:76
	_go_fuzz_dep_.CoverTab[67654]++
												registeredCompressor[c.Name()] = c
												if !grpcutil.IsCompressorNameRegistered(c.Name()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:78
		_go_fuzz_dep_.CoverTab[67655]++
													grpcutil.RegisteredCompressorNames = append(grpcutil.RegisteredCompressorNames, c.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:79
		// _ = "end of CoverTab[67655]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:80
		_go_fuzz_dep_.CoverTab[67656]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:80
		// _ = "end of CoverTab[67656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:80
	// _ = "end of CoverTab[67654]"
}

// GetCompressor returns Compressor for the given compressor name.
func GetCompressor(name string) Compressor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:84
	_go_fuzz_dep_.CoverTab[67657]++
												return registeredCompressor[name]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:85
	// _ = "end of CoverTab[67657]"
}

// Codec defines the interface gRPC uses to encode and decode messages.  Note
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:88
// that implementations of this interface must be thread safe; a Codec's
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:88
// methods can be called from concurrent goroutines.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:91
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
}

var registeredCodecs = make(map[string]Codec)

// RegisterCodec registers the provided Codec for use with all gRPC clients and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// servers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// The Codec will be stored and looked up by result of its Name() method, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// should match the content-subtype of the encoding handled by the Codec.  This
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// is case-insensitive, and is stored and looked up as lowercase.  If the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// result of calling Name() is an empty string, RegisterCodec will panic. See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// Content-Type on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// NOTE: this function must only be called during initialization time (i.e. in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// an init() function), and is not thread-safe.  If multiple Codecs are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:104
// registered with the same name, the one registered last will take effect.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:118
func RegisterCodec(codec Codec) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:118
	_go_fuzz_dep_.CoverTab[67658]++
												if codec == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:119
		_go_fuzz_dep_.CoverTab[67661]++
													panic("cannot register a nil Codec")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:120
		// _ = "end of CoverTab[67661]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:121
		_go_fuzz_dep_.CoverTab[67662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:121
		// _ = "end of CoverTab[67662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:121
	// _ = "end of CoverTab[67658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:121
	_go_fuzz_dep_.CoverTab[67659]++
												if codec.Name() == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:122
		_go_fuzz_dep_.CoverTab[67663]++
													panic("cannot register Codec with empty string result for Name()")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:123
		// _ = "end of CoverTab[67663]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:124
		_go_fuzz_dep_.CoverTab[67664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:124
		// _ = "end of CoverTab[67664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:124
	// _ = "end of CoverTab[67659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:124
	_go_fuzz_dep_.CoverTab[67660]++
												contentSubtype := strings.ToLower(codec.Name())
												registeredCodecs[contentSubtype] = codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:126
	// _ = "end of CoverTab[67660]"
}

// GetCodec gets a registered Codec by content-subtype, or nil if no Codec is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:129
// registered for the content-subtype.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:129
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:129
// The content-subtype is expected to be lowercase.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:133
func GetCodec(contentSubtype string) Codec {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:133
	_go_fuzz_dep_.CoverTab[67665]++
												return registeredCodecs[contentSubtype]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:134
	// _ = "end of CoverTab[67665]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:135
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/encoding.go:135
var _ = _go_fuzz_dep_.CoverTab
