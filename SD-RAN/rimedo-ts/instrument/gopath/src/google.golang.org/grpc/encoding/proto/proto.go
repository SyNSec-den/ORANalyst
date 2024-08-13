//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:19
// Package proto defines the protobuf codec. Importing this package will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:19
// register the codec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:21
)

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
)

// Name is the name registered for the proto compressor.
const Name = "proto"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with protobuf. It is the default codec for gRPC.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:40
	_go_fuzz_dep_.CoverTab[67666]++
												vv, ok := v.(proto.Message)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:42
		_go_fuzz_dep_.CoverTab[67668]++
													return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:43
		// _ = "end of CoverTab[67668]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:44
		_go_fuzz_dep_.CoverTab[67669]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:44
		// _ = "end of CoverTab[67669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:44
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:44
	// _ = "end of CoverTab[67666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:44
	_go_fuzz_dep_.CoverTab[67667]++
												return proto.Marshal(vv)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:45
	// _ = "end of CoverTab[67667]"
}

func (codec) Unmarshal(data []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:48
	_go_fuzz_dep_.CoverTab[67670]++
												vv, ok := v.(proto.Message)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:50
		_go_fuzz_dep_.CoverTab[67672]++
													return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:51
		// _ = "end of CoverTab[67672]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:52
		_go_fuzz_dep_.CoverTab[67673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:52
		// _ = "end of CoverTab[67673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:52
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:52
	// _ = "end of CoverTab[67670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:52
	_go_fuzz_dep_.CoverTab[67671]++
												return proto.Unmarshal(data, vv)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:53
	// _ = "end of CoverTab[67671]"
}

func (codec) Name() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:56
	_go_fuzz_dep_.CoverTab[67674]++
												return Name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:57
	// _ = "end of CoverTab[67674]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:58
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/encoding/proto/proto.go:58
var _ = _go_fuzz_dep_.CoverTab
