//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:24
)

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/http2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/grpcutil"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

// NewServerHandlerTransport returns a ServerTransport handling gRPC from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:49
// inside an http.Handler, or writes an HTTP error to w and returns an error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:49
// It requires that the http Server supports HTTP/2.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:52
func NewServerHandlerTransport(w http.ResponseWriter, r *http.Request, stats []stats.Handler) (ServerTransport, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:52
	_go_fuzz_dep_.CoverTab[76997]++
														if r.ProtoMajor != 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:53
		_go_fuzz_dep_.CoverTab[77005]++
															msg := "gRPC requires HTTP/2"
															http.Error(w, msg, http.StatusBadRequest)
															return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:56
		// _ = "end of CoverTab[77005]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:57
		_go_fuzz_dep_.CoverTab[77006]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:57
		// _ = "end of CoverTab[77006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:57
	// _ = "end of CoverTab[76997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:57
	_go_fuzz_dep_.CoverTab[76998]++
														if r.Method != "POST" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:58
		_go_fuzz_dep_.CoverTab[77007]++
															msg := fmt.Sprintf("invalid gRPC request method %q", r.Method)
															http.Error(w, msg, http.StatusBadRequest)
															return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:61
		// _ = "end of CoverTab[77007]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:62
		_go_fuzz_dep_.CoverTab[77008]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:62
		// _ = "end of CoverTab[77008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:62
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:62
	// _ = "end of CoverTab[76998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:62
	_go_fuzz_dep_.CoverTab[76999]++
														contentType := r.Header.Get("Content-Type")

														contentSubtype, validContentType := grpcutil.ContentSubtype(contentType)
														if !validContentType {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:66
		_go_fuzz_dep_.CoverTab[77009]++
															msg := fmt.Sprintf("invalid gRPC request content-type %q", contentType)
															http.Error(w, msg, http.StatusUnsupportedMediaType)
															return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:69
		// _ = "end of CoverTab[77009]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:70
		_go_fuzz_dep_.CoverTab[77010]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:70
		// _ = "end of CoverTab[77010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:70
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:70
	// _ = "end of CoverTab[76999]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:70
	_go_fuzz_dep_.CoverTab[77000]++
														if _, ok := w.(http.Flusher); !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:71
		_go_fuzz_dep_.CoverTab[77011]++
															msg := "gRPC requires a ResponseWriter supporting http.Flusher"
															http.Error(w, msg, http.StatusInternalServerError)
															return nil, errors.New(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:74
		// _ = "end of CoverTab[77011]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:75
		_go_fuzz_dep_.CoverTab[77012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:75
		// _ = "end of CoverTab[77012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:75
	// _ = "end of CoverTab[77000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:75
	_go_fuzz_dep_.CoverTab[77001]++

														st := &serverHandlerTransport{
		rw:		w,
		req:		r,
		closedCh:	make(chan struct{}),
		writes:		make(chan func()),
		contentType:	contentType,
		contentSubtype:	contentSubtype,
		stats:		stats,
	}

	if v := r.Header.Get("grpc-timeout"); v != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:87
		_go_fuzz_dep_.CoverTab[77013]++
															to, err := decodeTimeout(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:89
			_go_fuzz_dep_.CoverTab[77015]++
																msg := fmt.Sprintf("malformed grpc-timeout: %v", err)
																http.Error(w, msg, http.StatusBadRequest)
																return nil, status.Error(codes.Internal, msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:92
			// _ = "end of CoverTab[77015]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:93
			_go_fuzz_dep_.CoverTab[77016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:93
			// _ = "end of CoverTab[77016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:93
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:93
		// _ = "end of CoverTab[77013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:93
		_go_fuzz_dep_.CoverTab[77014]++
															st.timeoutSet = true
															st.timeout = to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:95
		// _ = "end of CoverTab[77014]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:96
		_go_fuzz_dep_.CoverTab[77017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:96
		// _ = "end of CoverTab[77017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:96
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:96
	// _ = "end of CoverTab[77001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:96
	_go_fuzz_dep_.CoverTab[77002]++

														metakv := []string{"content-type", contentType}
														if r.Host != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:99
		_go_fuzz_dep_.CoverTab[77018]++
															metakv = append(metakv, ":authority", r.Host)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:100
		// _ = "end of CoverTab[77018]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:101
		_go_fuzz_dep_.CoverTab[77019]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:101
		// _ = "end of CoverTab[77019]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:101
	// _ = "end of CoverTab[77002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:101
	_go_fuzz_dep_.CoverTab[77003]++
														for k, vv := range r.Header {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:102
		_go_fuzz_dep_.CoverTab[77020]++
															k = strings.ToLower(k)
															if isReservedHeader(k) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:104
			_go_fuzz_dep_.CoverTab[77022]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:104
			return !isWhitelistedHeader(k)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:104
			// _ = "end of CoverTab[77022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:104
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:104
			_go_fuzz_dep_.CoverTab[77023]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:105
			// _ = "end of CoverTab[77023]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:106
			_go_fuzz_dep_.CoverTab[77024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:106
			// _ = "end of CoverTab[77024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:106
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:106
		// _ = "end of CoverTab[77020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:106
		_go_fuzz_dep_.CoverTab[77021]++
															for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:107
			_go_fuzz_dep_.CoverTab[77025]++
																v, err := decodeMetadataHeader(k, v)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:109
				_go_fuzz_dep_.CoverTab[77027]++
																	msg := fmt.Sprintf("malformed binary metadata %q in header %q: %v", v, k, err)
																	http.Error(w, msg, http.StatusBadRequest)
																	return nil, status.Error(codes.Internal, msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:112
				// _ = "end of CoverTab[77027]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:113
				_go_fuzz_dep_.CoverTab[77028]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:113
				// _ = "end of CoverTab[77028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:113
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:113
			// _ = "end of CoverTab[77025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:113
			_go_fuzz_dep_.CoverTab[77026]++
																metakv = append(metakv, k, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:114
			// _ = "end of CoverTab[77026]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:115
		// _ = "end of CoverTab[77021]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:116
	// _ = "end of CoverTab[77003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:116
	_go_fuzz_dep_.CoverTab[77004]++
														st.headerMD = metadata.Pairs(metakv...)

														return st, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:119
	// _ = "end of CoverTab[77004]"
}

// serverHandlerTransport is an implementation of ServerTransport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:122
// which replies to exactly one gRPC request (exactly one HTTP request),
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:122
// using the net/http.Handler interface. This http.Handler is guaranteed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:122
// at this point to be speaking over HTTP/2, so it's able to speak valid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:122
// gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:127
type serverHandlerTransport struct {
	rw		http.ResponseWriter
	req		*http.Request
	timeoutSet	bool
	timeout		time.Duration

	headerMD	metadata.MD

	closeOnce	sync.Once
	closedCh	chan struct{}	// closed on Close

	// writes is a channel of code to run serialized in the
	// ServeHTTP (HandleStreams) goroutine. The channel is closed
	// when WriteStatus is called.
	writes	chan func()

	// block concurrent WriteStatus calls
	// e.g. grpc/(*serverStream).SendMsg/RecvMsg
	writeStatusMu	sync.Mutex

	// we just mirror the request content-type
	contentType	string
	// we store both contentType and contentSubtype so we don't keep recreating them
	// TODO make sure this is consistent across handler_server and http2_server
	contentSubtype	string

	stats	[]stats.Handler
}

func (ht *serverHandlerTransport) Close(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:156
	_go_fuzz_dep_.CoverTab[77029]++
														ht.closeOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:157
		_go_fuzz_dep_.CoverTab[77030]++
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:158
			_go_fuzz_dep_.CoverTab[77032]++
																logger.Infof("Closing serverHandlerTransport: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:159
			// _ = "end of CoverTab[77032]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:160
			_go_fuzz_dep_.CoverTab[77033]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:160
			// _ = "end of CoverTab[77033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:160
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:160
		// _ = "end of CoverTab[77030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:160
		_go_fuzz_dep_.CoverTab[77031]++
															close(ht.closedCh)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:161
		// _ = "end of CoverTab[77031]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:162
	// _ = "end of CoverTab[77029]"
}

func (ht *serverHandlerTransport) RemoteAddr() net.Addr {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:165
	_go_fuzz_dep_.CoverTab[77034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:165
	return strAddr(ht.req.RemoteAddr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:165
	// _ = "end of CoverTab[77034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:165
}

// strAddr is a net.Addr backed by either a TCP "ip:port" string, or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:167
// the empty string if unknown.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:169
type strAddr string

func (a strAddr) Network() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:171
	_go_fuzz_dep_.CoverTab[77035]++
														if a != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:172
		_go_fuzz_dep_.CoverTab[77037]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:182
		return "tcp"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:182
		// _ = "end of CoverTab[77037]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:183
		_go_fuzz_dep_.CoverTab[77038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:183
		// _ = "end of CoverTab[77038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:183
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:183
	// _ = "end of CoverTab[77035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:183
	_go_fuzz_dep_.CoverTab[77036]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:184
	// _ = "end of CoverTab[77036]"
}

func (a strAddr) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:187
	_go_fuzz_dep_.CoverTab[77039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:187
	return string(a)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:187
	// _ = "end of CoverTab[77039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:187
}

// do runs fn in the ServeHTTP goroutine.
func (ht *serverHandlerTransport) do(fn func()) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:190
	_go_fuzz_dep_.CoverTab[77040]++
														select {
	case <-ht.closedCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:192
		_go_fuzz_dep_.CoverTab[77041]++
															return ErrConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:193
		// _ = "end of CoverTab[77041]"
	case ht.writes <- fn:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:194
		_go_fuzz_dep_.CoverTab[77042]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:195
		// _ = "end of CoverTab[77042]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:196
	// _ = "end of CoverTab[77040]"
}

func (ht *serverHandlerTransport) WriteStatus(s *Stream, st *status.Status) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:199
	_go_fuzz_dep_.CoverTab[77043]++
														ht.writeStatusMu.Lock()
														defer ht.writeStatusMu.Unlock()

														headersWritten := s.updateHeaderSent()
														err := ht.do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:204
		_go_fuzz_dep_.CoverTab[77046]++
															if !headersWritten {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:205
			_go_fuzz_dep_.CoverTab[77050]++
																ht.writePendingHeaders(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:206
			// _ = "end of CoverTab[77050]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:207
			_go_fuzz_dep_.CoverTab[77051]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:207
			// _ = "end of CoverTab[77051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:207
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:207
		// _ = "end of CoverTab[77046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:207
		_go_fuzz_dep_.CoverTab[77047]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:212
		ht.rw.(http.Flusher).Flush()

		h := ht.rw.Header()
		h.Set("Grpc-Status", fmt.Sprintf("%d", st.Code()))
		if m := st.Message(); m != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:216
			_go_fuzz_dep_.CoverTab[77052]++
																h.Set("Grpc-Message", encodeGrpcMessage(m))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:217
			// _ = "end of CoverTab[77052]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:218
			_go_fuzz_dep_.CoverTab[77053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:218
			// _ = "end of CoverTab[77053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:218
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:218
		// _ = "end of CoverTab[77047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:218
		_go_fuzz_dep_.CoverTab[77048]++

															if p := st.Proto(); p != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:220
			_go_fuzz_dep_.CoverTab[77054]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:220
			return len(p.Details) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:220
			// _ = "end of CoverTab[77054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:220
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:220
			_go_fuzz_dep_.CoverTab[77055]++
																stBytes, err := proto.Marshal(p)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:222
				_go_fuzz_dep_.CoverTab[77057]++

																	panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:224
				// _ = "end of CoverTab[77057]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:225
				_go_fuzz_dep_.CoverTab[77058]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:225
				// _ = "end of CoverTab[77058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:225
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:225
			// _ = "end of CoverTab[77055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:225
			_go_fuzz_dep_.CoverTab[77056]++

																h.Set("Grpc-Status-Details-Bin", encodeBinHeader(stBytes))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:227
			// _ = "end of CoverTab[77056]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:228
			_go_fuzz_dep_.CoverTab[77059]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:228
			// _ = "end of CoverTab[77059]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:228
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:228
		// _ = "end of CoverTab[77048]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:228
		_go_fuzz_dep_.CoverTab[77049]++

															if md := s.Trailer(); len(md) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:230
			_go_fuzz_dep_.CoverTab[77060]++
																for k, vv := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:231
				_go_fuzz_dep_.CoverTab[77061]++

																	if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:233
					_go_fuzz_dep_.CoverTab[77063]++
																		continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:234
					// _ = "end of CoverTab[77063]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:235
					_go_fuzz_dep_.CoverTab[77064]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:235
					// _ = "end of CoverTab[77064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:235
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:235
				// _ = "end of CoverTab[77061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:235
				_go_fuzz_dep_.CoverTab[77062]++
																	for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:236
					_go_fuzz_dep_.CoverTab[77065]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:239
					h.Add(http2.TrailerPrefix+k, encodeMetadataHeader(k, v))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:239
					// _ = "end of CoverTab[77065]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:240
				// _ = "end of CoverTab[77062]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:241
			// _ = "end of CoverTab[77060]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:242
			_go_fuzz_dep_.CoverTab[77066]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:242
			// _ = "end of CoverTab[77066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:242
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:242
		// _ = "end of CoverTab[77049]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:243
	// _ = "end of CoverTab[77043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:243
	_go_fuzz_dep_.CoverTab[77044]++

														if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:245
		_go_fuzz_dep_.CoverTab[77067]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:248
		for _, sh := range ht.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:248
			_go_fuzz_dep_.CoverTab[77068]++
																sh.HandleRPC(s.Context(), &stats.OutTrailer{
				Trailer: s.trailer.Copy(),
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:251
			// _ = "end of CoverTab[77068]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:252
		// _ = "end of CoverTab[77067]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:253
		_go_fuzz_dep_.CoverTab[77069]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:253
		// _ = "end of CoverTab[77069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:253
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:253
	// _ = "end of CoverTab[77044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:253
	_go_fuzz_dep_.CoverTab[77045]++
														ht.Close(errors.New("finished writing status"))
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:255
	// _ = "end of CoverTab[77045]"
}

// writePendingHeaders sets common and custom headers on the first
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:258
// write call (Write, WriteHeader, or WriteStatus)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:260
func (ht *serverHandlerTransport) writePendingHeaders(s *Stream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:260
	_go_fuzz_dep_.CoverTab[77070]++
														ht.writeCommonHeaders(s)
														ht.writeCustomHeaders(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:262
	// _ = "end of CoverTab[77070]"
}

// writeCommonHeaders sets common headers on the first write
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:265
// call (Write, WriteHeader, or WriteStatus).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:267
func (ht *serverHandlerTransport) writeCommonHeaders(s *Stream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:267
	_go_fuzz_dep_.CoverTab[77071]++
														h := ht.rw.Header()
														h["Date"] = nil
														h.Set("Content-Type", ht.contentType)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:277
	h.Add("Trailer", "Grpc-Status")
	h.Add("Trailer", "Grpc-Message")
	h.Add("Trailer", "Grpc-Status-Details-Bin")

	if s.sendCompress != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:281
		_go_fuzz_dep_.CoverTab[77072]++
															h.Set("Grpc-Encoding", s.sendCompress)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:282
		// _ = "end of CoverTab[77072]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:283
		_go_fuzz_dep_.CoverTab[77073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:283
		// _ = "end of CoverTab[77073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:283
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:283
	// _ = "end of CoverTab[77071]"
}

// writeCustomHeaders sets custom headers set on the stream via SetHeader
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:286
// on the first write call (Write, WriteHeader, or WriteStatus).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:288
func (ht *serverHandlerTransport) writeCustomHeaders(s *Stream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:288
	_go_fuzz_dep_.CoverTab[77074]++
														h := ht.rw.Header()

														s.hdrMu.Lock()
														for k, vv := range s.header {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:292
		_go_fuzz_dep_.CoverTab[77076]++
															if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:293
			_go_fuzz_dep_.CoverTab[77078]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:294
			// _ = "end of CoverTab[77078]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:295
			_go_fuzz_dep_.CoverTab[77079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:295
			// _ = "end of CoverTab[77079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:295
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:295
		// _ = "end of CoverTab[77076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:295
		_go_fuzz_dep_.CoverTab[77077]++
															for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:296
			_go_fuzz_dep_.CoverTab[77080]++
																h.Add(k, encodeMetadataHeader(k, v))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:297
			// _ = "end of CoverTab[77080]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:298
		// _ = "end of CoverTab[77077]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:299
	// _ = "end of CoverTab[77074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:299
	_go_fuzz_dep_.CoverTab[77075]++

														s.hdrMu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:301
	// _ = "end of CoverTab[77075]"
}

func (ht *serverHandlerTransport) Write(s *Stream, hdr []byte, data []byte, opts *Options) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:304
	_go_fuzz_dep_.CoverTab[77081]++
														headersWritten := s.updateHeaderSent()
														return ht.do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:306
		_go_fuzz_dep_.CoverTab[77082]++
															if !headersWritten {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:307
			_go_fuzz_dep_.CoverTab[77084]++
																ht.writePendingHeaders(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:308
			// _ = "end of CoverTab[77084]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:309
			_go_fuzz_dep_.CoverTab[77085]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:309
			// _ = "end of CoverTab[77085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:309
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:309
		// _ = "end of CoverTab[77082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:309
		_go_fuzz_dep_.CoverTab[77083]++
															ht.rw.Write(hdr)
															ht.rw.Write(data)
															ht.rw.(http.Flusher).Flush()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:312
		// _ = "end of CoverTab[77083]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:313
	// _ = "end of CoverTab[77081]"
}

func (ht *serverHandlerTransport) WriteHeader(s *Stream, md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:316
	_go_fuzz_dep_.CoverTab[77086]++
														if err := s.SetHeader(md); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:317
		_go_fuzz_dep_.CoverTab[77090]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:318
		// _ = "end of CoverTab[77090]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:319
		_go_fuzz_dep_.CoverTab[77091]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:319
		// _ = "end of CoverTab[77091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:319
	// _ = "end of CoverTab[77086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:319
	_go_fuzz_dep_.CoverTab[77087]++

														headersWritten := s.updateHeaderSent()
														err := ht.do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:322
		_go_fuzz_dep_.CoverTab[77092]++
															if !headersWritten {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:323
			_go_fuzz_dep_.CoverTab[77094]++
																ht.writePendingHeaders(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:324
			// _ = "end of CoverTab[77094]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:325
			_go_fuzz_dep_.CoverTab[77095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:325
			// _ = "end of CoverTab[77095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:325
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:325
		// _ = "end of CoverTab[77092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:325
		_go_fuzz_dep_.CoverTab[77093]++

															ht.rw.WriteHeader(200)
															ht.rw.(http.Flusher).Flush()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:328
		// _ = "end of CoverTab[77093]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:329
	// _ = "end of CoverTab[77087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:329
	_go_fuzz_dep_.CoverTab[77088]++

														if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:331
		_go_fuzz_dep_.CoverTab[77096]++
															for _, sh := range ht.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:332
			_go_fuzz_dep_.CoverTab[77097]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:335
			sh.HandleRPC(s.Context(), &stats.OutHeader{
				Header:		md.Copy(),
				Compression:	s.sendCompress,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:338
			// _ = "end of CoverTab[77097]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:339
		// _ = "end of CoverTab[77096]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:340
		_go_fuzz_dep_.CoverTab[77098]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:340
		// _ = "end of CoverTab[77098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:340
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:340
	// _ = "end of CoverTab[77088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:340
	_go_fuzz_dep_.CoverTab[77089]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:341
	// _ = "end of CoverTab[77089]"
}

func (ht *serverHandlerTransport) HandleStreams(startStream func(*Stream), traceCtx func(context.Context, string) context.Context) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:344
	_go_fuzz_dep_.CoverTab[77099]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:347
	ctx := ht.req.Context()
	var cancel context.CancelFunc
	if ht.timeoutSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:349
		_go_fuzz_dep_.CoverTab[77107]++
															ctx, cancel = context.WithTimeout(ctx, ht.timeout)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:350
		// _ = "end of CoverTab[77107]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:351
		_go_fuzz_dep_.CoverTab[77108]++
															ctx, cancel = context.WithCancel(ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:352
		// _ = "end of CoverTab[77108]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:353
	// _ = "end of CoverTab[77099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:353
	_go_fuzz_dep_.CoverTab[77100]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:356
	requestOver := make(chan struct{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:356
	_curRoutineNum76_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:356
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum76_)
														go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
		_go_fuzz_dep_.CoverTab[77109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
			_go_fuzz_dep_.CoverTab[77111]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum76_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
			// _ = "end of CoverTab[77111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:357
		}()
															select {
		case <-requestOver:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:359
			_go_fuzz_dep_.CoverTab[77112]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:359
			// _ = "end of CoverTab[77112]"
		case <-ht.closedCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:360
			_go_fuzz_dep_.CoverTab[77113]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:360
			// _ = "end of CoverTab[77113]"
		case <-ht.req.Context().Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:361
			_go_fuzz_dep_.CoverTab[77114]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:361
			// _ = "end of CoverTab[77114]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:362
		// _ = "end of CoverTab[77109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:362
		_go_fuzz_dep_.CoverTab[77110]++
															cancel()
															ht.Close(errors.New("request is done processing"))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:364
		// _ = "end of CoverTab[77110]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:365
	// _ = "end of CoverTab[77100]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:365
	_go_fuzz_dep_.CoverTab[77101]++

														req := ht.req

														s := &Stream{
		id:		0,
		requestRead:	func(int) { _go_fuzz_dep_.CoverTab[77115]++; // _ = "end of CoverTab[77115]" },
		cancel:		cancel,
		buf:		newRecvBuffer(),
		st:		ht,
		method:		req.URL.Path,
		recvCompress:	req.Header.Get("grpc-encoding"),
		contentSubtype:	ht.contentSubtype,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:378
	// _ = "end of CoverTab[77101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:378
	_go_fuzz_dep_.CoverTab[77102]++
														pr := &peer.Peer{
		Addr: ht.RemoteAddr(),
	}
	if req.TLS != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:382
		_go_fuzz_dep_.CoverTab[77116]++
															pr.AuthInfo = credentials.TLSInfo{State: *req.TLS, CommonAuthInfo: credentials.CommonAuthInfo{SecurityLevel: credentials.PrivacyAndIntegrity}}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:383
		// _ = "end of CoverTab[77116]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:384
		_go_fuzz_dep_.CoverTab[77117]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:384
		// _ = "end of CoverTab[77117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:384
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:384
	// _ = "end of CoverTab[77102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:384
	_go_fuzz_dep_.CoverTab[77103]++
														ctx = metadata.NewIncomingContext(ctx, ht.headerMD)
														s.ctx = peer.NewContext(ctx, pr)
														for _, sh := range ht.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:387
		_go_fuzz_dep_.CoverTab[77118]++
															s.ctx = sh.TagRPC(s.ctx, &stats.RPCTagInfo{FullMethodName: s.method})
															inHeader := &stats.InHeader{
			FullMethod:	s.method,
			RemoteAddr:	ht.RemoteAddr(),
			Compression:	s.recvCompress,
		}
															sh.HandleRPC(s.ctx, inHeader)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:394
		// _ = "end of CoverTab[77118]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:395
	// _ = "end of CoverTab[77103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:395
	_go_fuzz_dep_.CoverTab[77104]++
														s.trReader = &transportReader{
		reader:		&recvBufferReader{ctx: s.ctx, ctxDone: s.ctx.Done(), recv: s.buf, freeBuffer: func(*bytes.Buffer) { _go_fuzz_dep_.CoverTab[77119]++; // _ = "end of CoverTab[77119]" }},
		windowHandler:	func(int) { _go_fuzz_dep_.CoverTab[77120]++; // _ = "end of CoverTab[77120]" },
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:399
	// _ = "end of CoverTab[77104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:399
	_go_fuzz_dep_.CoverTab[77105]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:402
	readerDone := make(chan struct{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:402
	_curRoutineNum77_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:402
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum77_)
														go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
		_go_fuzz_dep_.CoverTab[77121]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
			_go_fuzz_dep_.CoverTab[77122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum77_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
			// _ = "end of CoverTab[77122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:403
		}()
															defer close(readerDone)

		// TODO: minimize garbage, optimize recvBuffer code/ownership
		const readSize = 8196
		for buf := make([]byte, readSize); ; {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:408
			_go_fuzz_dep_.CoverTab[77123]++
																n, err := req.Body.Read(buf)
																if n > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:410
				_go_fuzz_dep_.CoverTab[77126]++
																	s.buf.put(recvMsg{buffer: bytes.NewBuffer(buf[:n:n])})
																	buf = buf[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:412
				// _ = "end of CoverTab[77126]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:413
				_go_fuzz_dep_.CoverTab[77127]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:413
				// _ = "end of CoverTab[77127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:413
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:413
			// _ = "end of CoverTab[77123]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:413
			_go_fuzz_dep_.CoverTab[77124]++
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:414
				_go_fuzz_dep_.CoverTab[77128]++
																	s.buf.put(recvMsg{err: mapRecvMsgError(err)})
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:416
				// _ = "end of CoverTab[77128]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:417
				_go_fuzz_dep_.CoverTab[77129]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:417
				// _ = "end of CoverTab[77129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:417
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:417
			// _ = "end of CoverTab[77124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:417
			_go_fuzz_dep_.CoverTab[77125]++
																if len(buf) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:418
				_go_fuzz_dep_.CoverTab[77130]++
																	buf = make([]byte, readSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:419
				// _ = "end of CoverTab[77130]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:420
				_go_fuzz_dep_.CoverTab[77131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:420
				// _ = "end of CoverTab[77131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:420
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:420
			// _ = "end of CoverTab[77125]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:421
		// _ = "end of CoverTab[77121]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:422
	// _ = "end of CoverTab[77105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:422
	_go_fuzz_dep_.CoverTab[77106]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:428
	startStream(s)

														ht.runStream()
														close(requestOver)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:434
	req.Body.Close()
														<-readerDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:435
	// _ = "end of CoverTab[77106]"
}

func (ht *serverHandlerTransport) runStream() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:438
	_go_fuzz_dep_.CoverTab[77132]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:439
		_go_fuzz_dep_.CoverTab[77133]++
															select {
		case fn := <-ht.writes:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:441
			_go_fuzz_dep_.CoverTab[77134]++
																fn()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:442
			// _ = "end of CoverTab[77134]"
		case <-ht.closedCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:443
			_go_fuzz_dep_.CoverTab[77135]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:444
			// _ = "end of CoverTab[77135]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:445
		// _ = "end of CoverTab[77133]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:446
	// _ = "end of CoverTab[77132]"
}

func (ht *serverHandlerTransport) IncrMsgSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:449
	_go_fuzz_dep_.CoverTab[77136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:449
	// _ = "end of CoverTab[77136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:449
}

func (ht *serverHandlerTransport) IncrMsgRecv() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:451
	_go_fuzz_dep_.CoverTab[77137]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:451
	// _ = "end of CoverTab[77137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:451
}

func (ht *serverHandlerTransport) Drain() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:453
	_go_fuzz_dep_.CoverTab[77138]++
														panic("Drain() is not implemented")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:454
	// _ = "end of CoverTab[77138]"
}

// mapRecvMsgError returns the non-nil err into the appropriate
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
// error value as expected by callers of *grpc.parser.recvMsg.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
// In particular, in can only be:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
//   - io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
//   - io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
//   - of type transport.ConnectionError
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:457
//   - an error from the status package
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:464
func mapRecvMsgError(err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:464
	_go_fuzz_dep_.CoverTab[77139]++
														if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:465
		_go_fuzz_dep_.CoverTab[77143]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:465
		return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:465
		// _ = "end of CoverTab[77143]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:465
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:465
		_go_fuzz_dep_.CoverTab[77144]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:466
		// _ = "end of CoverTab[77144]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:467
		_go_fuzz_dep_.CoverTab[77145]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:467
		// _ = "end of CoverTab[77145]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:467
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:467
	// _ = "end of CoverTab[77139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:467
	_go_fuzz_dep_.CoverTab[77140]++
														if se, ok := err.(http2.StreamError); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:468
		_go_fuzz_dep_.CoverTab[77146]++
															if code, ok := http2ErrConvTab[se.Code]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:469
			_go_fuzz_dep_.CoverTab[77147]++
																return status.Error(code, se.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:470
			// _ = "end of CoverTab[77147]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:471
			_go_fuzz_dep_.CoverTab[77148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:471
			// _ = "end of CoverTab[77148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:471
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:471
		// _ = "end of CoverTab[77146]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:472
		_go_fuzz_dep_.CoverTab[77149]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:472
		// _ = "end of CoverTab[77149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:472
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:472
	// _ = "end of CoverTab[77140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:472
	_go_fuzz_dep_.CoverTab[77141]++
														if strings.Contains(err.Error(), "body closed by handler") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:473
		_go_fuzz_dep_.CoverTab[77150]++
															return status.Error(codes.Canceled, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:474
		// _ = "end of CoverTab[77150]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:475
		_go_fuzz_dep_.CoverTab[77151]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:475
		// _ = "end of CoverTab[77151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:475
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:475
	// _ = "end of CoverTab[77141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:475
	_go_fuzz_dep_.CoverTab[77142]++
														return connectionErrorf(true, err, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:476
	// _ = "end of CoverTab[77142]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:477
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/handler_server.go:477
var _ = _go_fuzz_dep_.CoverTab
