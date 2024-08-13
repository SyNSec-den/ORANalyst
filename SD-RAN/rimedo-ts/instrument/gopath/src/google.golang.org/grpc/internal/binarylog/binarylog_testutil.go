//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
package binarylog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:30
)

var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger	= NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.
	MdToMetadataProto	= mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto	= addrToProto
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog_testutil.go:42
var _ = _go_fuzz_dep_.CoverTab
