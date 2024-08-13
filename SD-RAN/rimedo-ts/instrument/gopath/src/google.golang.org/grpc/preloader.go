//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:19
)

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PreparedMsg is responsible for creating a Marshalled and Compressed object.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:26
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:26
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:26
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:32
type PreparedMsg struct {
	// Struct for preparing msg before sending them
	encodedData	[]byte
	hdr		[]byte
	payload		[]byte
}

// Encode marshalls and compresses the message using the codec and compressor for the stream.
func (p *PreparedMsg) Encode(s Stream, msg interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:40
	_go_fuzz_dep_.CoverTab[79594]++
											ctx := s.Context()
											rpcInfo, ok := rpcInfoFromContext(ctx)
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:43
		_go_fuzz_dep_.CoverTab[79600]++
												return status.Errorf(codes.Internal, "grpc: unable to get rpcInfo")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:44
		// _ = "end of CoverTab[79600]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:45
		_go_fuzz_dep_.CoverTab[79601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:45
		// _ = "end of CoverTab[79601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:45
	// _ = "end of CoverTab[79594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:45
	_go_fuzz_dep_.CoverTab[79595]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:48
	if rpcInfo.preloaderInfo == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:48
		_go_fuzz_dep_.CoverTab[79602]++
												return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo is nil")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:49
		// _ = "end of CoverTab[79602]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:50
		_go_fuzz_dep_.CoverTab[79603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:50
		// _ = "end of CoverTab[79603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:50
	// _ = "end of CoverTab[79595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:50
	_go_fuzz_dep_.CoverTab[79596]++
											if rpcInfo.preloaderInfo.codec == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:51
		_go_fuzz_dep_.CoverTab[79604]++
												return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo.codec is nil")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:52
		// _ = "end of CoverTab[79604]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:53
		_go_fuzz_dep_.CoverTab[79605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:53
		// _ = "end of CoverTab[79605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:53
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:53
	// _ = "end of CoverTab[79596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:53
	_go_fuzz_dep_.CoverTab[79597]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:56
	data, err := encode(rpcInfo.preloaderInfo.codec, msg)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:57
		_go_fuzz_dep_.CoverTab[79606]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:58
		// _ = "end of CoverTab[79606]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:59
		_go_fuzz_dep_.CoverTab[79607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:59
		// _ = "end of CoverTab[79607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:59
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:59
	// _ = "end of CoverTab[79597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:59
	_go_fuzz_dep_.CoverTab[79598]++
											p.encodedData = data
											compData, err := compress(data, rpcInfo.preloaderInfo.cp, rpcInfo.preloaderInfo.comp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:62
		_go_fuzz_dep_.CoverTab[79608]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:63
		// _ = "end of CoverTab[79608]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:64
		_go_fuzz_dep_.CoverTab[79609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:64
		// _ = "end of CoverTab[79609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:64
	// _ = "end of CoverTab[79598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:64
	_go_fuzz_dep_.CoverTab[79599]++
											p.hdr, p.payload = msgHeader(data, compData)
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:66
	// _ = "end of CoverTab[79599]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:67
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/preloader.go:67
var _ = _go_fuzz_dep_.CoverTab
