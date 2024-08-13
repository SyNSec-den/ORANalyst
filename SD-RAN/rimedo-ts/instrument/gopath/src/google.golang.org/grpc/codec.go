//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:19
)

import (
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"	// to register the Codec for "proto"
)

// baseCodec contains the functionality of both Codec and encoding.Codec, but
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:26
// omits the name/string, which vary between the two and are not needed for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:26
// anything besides the registry in the encoding package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:29
type baseCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var _ baseCodec = Codec(nil)
var _ baseCodec = encoding.Codec(nil)

// Codec defines the interface gRPC uses to encode and decode messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:37
// Note that implementations of this interface must be thread safe;
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:37
// a Codec's methods can be called from concurrent goroutines.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:37
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:37
// Deprecated: use encoding.Codec instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:42
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error
	// String returns the name of the Codec implementation.  This is unused by
	// gRPC.
	String() string
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:50
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codec.go:50
var _ = _go_fuzz_dep_.CoverTab
