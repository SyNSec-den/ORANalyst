//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:19
)

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const (
	// http2MaxFrameLen specifies the max length of a HTTP2 frame.
	http2MaxFrameLen	= 16384	// 16KB frame
	// https://httpwg.org/specs/rfc7540.html#SettingValues
	http2InitHeaderTableSize	= 4096
)

var (
	clientPreface	= []byte(http2.ClientPreface)
	http2ErrConvTab	= map[http2.ErrCode]codes.Code{
		http2.ErrCodeNo:			codes.Internal,
		http2.ErrCodeProtocol:			codes.Internal,
		http2.ErrCodeInternal:			codes.Internal,
		http2.ErrCodeFlowControl:		codes.ResourceExhausted,
		http2.ErrCodeSettingsTimeout:		codes.Internal,
		http2.ErrCodeStreamClosed:		codes.Internal,
		http2.ErrCodeFrameSize:			codes.Internal,
		http2.ErrCodeRefusedStream:		codes.Unavailable,
		http2.ErrCodeCancel:			codes.Canceled,
		http2.ErrCodeCompression:		codes.Internal,
		http2.ErrCodeConnect:			codes.Internal,
		http2.ErrCodeEnhanceYourCalm:		codes.ResourceExhausted,
		http2.ErrCodeInadequateSecurity:	codes.PermissionDenied,
		http2.ErrCodeHTTP11Required:		codes.Internal,
	}
	// HTTPStatusConvTab is the HTTP status code to gRPC error code conversion table.
	HTTPStatusConvTab	= map[int]codes.Code{

		http.StatusBadRequest:	codes.Internal,

		http.StatusUnauthorized:	codes.Unauthenticated,

		http.StatusForbidden:	codes.PermissionDenied,

		http.StatusNotFound:	codes.Unimplemented,

		http.StatusTooManyRequests:	codes.Unavailable,

		http.StatusBadGateway:	codes.Unavailable,

		http.StatusServiceUnavailable:	codes.Unavailable,

		http.StatusGatewayTimeout:	codes.Unavailable,
	}
	logger	= grpclog.Component("transport")
)

// isReservedHeader checks whether hdr belongs to HTTP2 headers
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:92
// reserved by gRPC protocol. Any other headers are classified as the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:92
// user-specified metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:95
func isReservedHeader(hdr string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:95
	_go_fuzz_dep_.CoverTab[78410]++
													if hdr != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:96
		_go_fuzz_dep_.CoverTab[78412]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:96
		return hdr[0] == ':'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:96
		// _ = "end of CoverTab[78412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:96
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:96
		_go_fuzz_dep_.CoverTab[78413]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:97
		// _ = "end of CoverTab[78413]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:98
		_go_fuzz_dep_.CoverTab[78414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:98
		// _ = "end of CoverTab[78414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:98
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:98
	// _ = "end of CoverTab[78410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:98
	_go_fuzz_dep_.CoverTab[78411]++
													switch hdr {
	case "content-type",
														"user-agent",
														"grpc-message-type",
														"grpc-encoding",
														"grpc-message",
														"grpc-status",
														"grpc-timeout",
														"grpc-status-details-bin",

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:111
		"te":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:111
		_go_fuzz_dep_.CoverTab[78415]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:112
		// _ = "end of CoverTab[78415]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:113
		_go_fuzz_dep_.CoverTab[78416]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:114
		// _ = "end of CoverTab[78416]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:115
	// _ = "end of CoverTab[78411]"
}

// isWhitelistedHeader checks whether hdr should be propagated into metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:118
// visible to users, even though it is classified as "reserved", above.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:120
func isWhitelistedHeader(hdr string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:120
	_go_fuzz_dep_.CoverTab[78417]++
													switch hdr {
	case ":authority", "user-agent":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:122
		_go_fuzz_dep_.CoverTab[78418]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:123
		// _ = "end of CoverTab[78418]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:124
		_go_fuzz_dep_.CoverTab[78419]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:125
		// _ = "end of CoverTab[78419]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:126
	// _ = "end of CoverTab[78417]"
}

const binHdrSuffix = "-bin"

func encodeBinHeader(v []byte) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:131
	_go_fuzz_dep_.CoverTab[78420]++
													return base64.RawStdEncoding.EncodeToString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:132
	// _ = "end of CoverTab[78420]"
}

func decodeBinHeader(v string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:135
	_go_fuzz_dep_.CoverTab[78421]++
													if len(v)%4 == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:136
		_go_fuzz_dep_.CoverTab[78423]++

														return base64.StdEncoding.DecodeString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:138
		// _ = "end of CoverTab[78423]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:139
		_go_fuzz_dep_.CoverTab[78424]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:139
		// _ = "end of CoverTab[78424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:139
	// _ = "end of CoverTab[78421]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:139
	_go_fuzz_dep_.CoverTab[78422]++
													return base64.RawStdEncoding.DecodeString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:140
	// _ = "end of CoverTab[78422]"
}

func encodeMetadataHeader(k, v string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:143
	_go_fuzz_dep_.CoverTab[78425]++
													if strings.HasSuffix(k, binHdrSuffix) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:144
		_go_fuzz_dep_.CoverTab[78427]++
														return encodeBinHeader(([]byte)(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:145
		// _ = "end of CoverTab[78427]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:146
		_go_fuzz_dep_.CoverTab[78428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:146
		// _ = "end of CoverTab[78428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:146
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:146
	// _ = "end of CoverTab[78425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:146
	_go_fuzz_dep_.CoverTab[78426]++
													return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:147
	// _ = "end of CoverTab[78426]"
}

func decodeMetadataHeader(k, v string) (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:150
	_go_fuzz_dep_.CoverTab[78429]++
													if strings.HasSuffix(k, binHdrSuffix) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:151
		_go_fuzz_dep_.CoverTab[78431]++
														b, err := decodeBinHeader(v)
														return string(b), err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:153
		// _ = "end of CoverTab[78431]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:154
		_go_fuzz_dep_.CoverTab[78432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:154
		// _ = "end of CoverTab[78432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:154
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:154
	// _ = "end of CoverTab[78429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:154
	_go_fuzz_dep_.CoverTab[78430]++
													return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:155
	// _ = "end of CoverTab[78430]"
}

func decodeGRPCStatusDetails(rawDetails string) (*status.Status, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:158
	_go_fuzz_dep_.CoverTab[78433]++
													v, err := decodeBinHeader(rawDetails)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:160
		_go_fuzz_dep_.CoverTab[78436]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:161
		// _ = "end of CoverTab[78436]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:162
		_go_fuzz_dep_.CoverTab[78437]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:162
		// _ = "end of CoverTab[78437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:162
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:162
	// _ = "end of CoverTab[78433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:162
	_go_fuzz_dep_.CoverTab[78434]++
													st := &spb.Status{}
													if err = proto.Unmarshal(v, st); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:164
		_go_fuzz_dep_.CoverTab[78438]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:165
		// _ = "end of CoverTab[78438]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:166
		_go_fuzz_dep_.CoverTab[78439]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:166
		// _ = "end of CoverTab[78439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:166
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:166
	// _ = "end of CoverTab[78434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:166
	_go_fuzz_dep_.CoverTab[78435]++
													return status.FromProto(st), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:167
	// _ = "end of CoverTab[78435]"
}

type timeoutUnit uint8

const (
	hour		timeoutUnit	= 'H'
	minute		timeoutUnit	= 'M'
	second		timeoutUnit	= 'S'
	millisecond	timeoutUnit	= 'm'
	microsecond	timeoutUnit	= 'u'
	nanosecond	timeoutUnit	= 'n'
)

func timeoutUnitToDuration(u timeoutUnit) (d time.Duration, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:181
	_go_fuzz_dep_.CoverTab[78440]++
													switch u {
	case hour:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:183
		_go_fuzz_dep_.CoverTab[78442]++
														return time.Hour, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:184
		// _ = "end of CoverTab[78442]"
	case minute:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:185
		_go_fuzz_dep_.CoverTab[78443]++
														return time.Minute, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:186
		// _ = "end of CoverTab[78443]"
	case second:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:187
		_go_fuzz_dep_.CoverTab[78444]++
														return time.Second, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:188
		// _ = "end of CoverTab[78444]"
	case millisecond:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:189
		_go_fuzz_dep_.CoverTab[78445]++
														return time.Millisecond, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:190
		// _ = "end of CoverTab[78445]"
	case microsecond:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:191
		_go_fuzz_dep_.CoverTab[78446]++
														return time.Microsecond, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:192
		// _ = "end of CoverTab[78446]"
	case nanosecond:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:193
		_go_fuzz_dep_.CoverTab[78447]++
														return time.Nanosecond, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:194
		// _ = "end of CoverTab[78447]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:195
		_go_fuzz_dep_.CoverTab[78448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:195
		// _ = "end of CoverTab[78448]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:196
	// _ = "end of CoverTab[78440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:196
	_go_fuzz_dep_.CoverTab[78441]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:197
	// _ = "end of CoverTab[78441]"
}

func decodeTimeout(s string) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:200
	_go_fuzz_dep_.CoverTab[78449]++
													size := len(s)
													if size < 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:202
		_go_fuzz_dep_.CoverTab[78455]++
														return 0, fmt.Errorf("transport: timeout string is too short: %q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:203
		// _ = "end of CoverTab[78455]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:204
		_go_fuzz_dep_.CoverTab[78456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:204
		// _ = "end of CoverTab[78456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:204
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:204
	// _ = "end of CoverTab[78449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:204
	_go_fuzz_dep_.CoverTab[78450]++
													if size > 9 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:205
		_go_fuzz_dep_.CoverTab[78457]++

														return 0, fmt.Errorf("transport: timeout string is too long: %q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:207
		// _ = "end of CoverTab[78457]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:208
		_go_fuzz_dep_.CoverTab[78458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:208
		// _ = "end of CoverTab[78458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:208
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:208
	// _ = "end of CoverTab[78450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:208
	_go_fuzz_dep_.CoverTab[78451]++
													unit := timeoutUnit(s[size-1])
													d, ok := timeoutUnitToDuration(unit)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:211
		_go_fuzz_dep_.CoverTab[78459]++
														return 0, fmt.Errorf("transport: timeout unit is not recognized: %q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:212
		// _ = "end of CoverTab[78459]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:213
		_go_fuzz_dep_.CoverTab[78460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:213
		// _ = "end of CoverTab[78460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:213
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:213
	// _ = "end of CoverTab[78451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:213
	_go_fuzz_dep_.CoverTab[78452]++
													t, err := strconv.ParseInt(s[:size-1], 10, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:215
		_go_fuzz_dep_.CoverTab[78461]++
														return 0, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:216
		// _ = "end of CoverTab[78461]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:217
		_go_fuzz_dep_.CoverTab[78462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:217
		// _ = "end of CoverTab[78462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:217
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:217
	// _ = "end of CoverTab[78452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:217
	_go_fuzz_dep_.CoverTab[78453]++
													const maxHours = math.MaxInt64 / int64(time.Hour)
													if d == time.Hour && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:219
		_go_fuzz_dep_.CoverTab[78463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:219
		return t > maxHours
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:219
		// _ = "end of CoverTab[78463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:219
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:219
		_go_fuzz_dep_.CoverTab[78464]++

														return time.Duration(math.MaxInt64), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:221
		// _ = "end of CoverTab[78464]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:222
		_go_fuzz_dep_.CoverTab[78465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:222
		// _ = "end of CoverTab[78465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:222
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:222
	// _ = "end of CoverTab[78453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:222
	_go_fuzz_dep_.CoverTab[78454]++
													return d * time.Duration(t), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:223
	// _ = "end of CoverTab[78454]"
}

const (
	spaceByte	= ' '
	tildeByte	= '~'
	percentByte	= '%'
)

// encodeGrpcMessage is used to encode status code in header field
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
// "grpc-message". It does percent encoding and also replaces invalid utf-8
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
// characters with Unicode replacement character.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
// It checks to see if each individual byte in msg is an allowable byte, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
// then either percent encoding or passing it through. When percent encoding,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:232
// the byte is converted into hexadecimal notation with a '%' prepended.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:239
func encodeGrpcMessage(msg string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:239
	_go_fuzz_dep_.CoverTab[78466]++
													if msg == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:240
		_go_fuzz_dep_.CoverTab[78469]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:241
		// _ = "end of CoverTab[78469]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:242
		_go_fuzz_dep_.CoverTab[78470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:242
		// _ = "end of CoverTab[78470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:242
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:242
	// _ = "end of CoverTab[78466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:242
	_go_fuzz_dep_.CoverTab[78467]++
													lenMsg := len(msg)
													for i := 0; i < lenMsg; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:244
		_go_fuzz_dep_.CoverTab[78471]++
														c := msg[i]
														if !(c >= spaceByte && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			_go_fuzz_dep_.CoverTab[78472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			return c <= tildeByte
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			// _ = "end of CoverTab[78472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			_go_fuzz_dep_.CoverTab[78473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			return c != percentByte
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			// _ = "end of CoverTab[78473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
		}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:246
			_go_fuzz_dep_.CoverTab[78474]++
															return encodeGrpcMessageUnchecked(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:247
			// _ = "end of CoverTab[78474]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:248
			_go_fuzz_dep_.CoverTab[78475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:248
			// _ = "end of CoverTab[78475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:248
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:248
		// _ = "end of CoverTab[78471]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:249
	// _ = "end of CoverTab[78467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:249
	_go_fuzz_dep_.CoverTab[78468]++
													return msg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:250
	// _ = "end of CoverTab[78468]"
}

func encodeGrpcMessageUnchecked(msg string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:253
	_go_fuzz_dep_.CoverTab[78476]++
													var sb strings.Builder
													for len(msg) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:255
		_go_fuzz_dep_.CoverTab[78478]++
														r, size := utf8.DecodeRuneInString(msg)
														for _, b := range []byte(string(r)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:257
			_go_fuzz_dep_.CoverTab[78480]++
															if size > 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:258
				_go_fuzz_dep_.CoverTab[78482]++

																fmt.Fprintf(&sb, "%%%02X", b)
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:261
				// _ = "end of CoverTab[78482]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:262
				_go_fuzz_dep_.CoverTab[78483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:262
				// _ = "end of CoverTab[78483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:262
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:262
			// _ = "end of CoverTab[78480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:262
			_go_fuzz_dep_.CoverTab[78481]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
			if b >= spaceByte && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				_go_fuzz_dep_.CoverTab[78484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				return b <= tildeByte
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				// _ = "end of CoverTab[78484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				_go_fuzz_dep_.CoverTab[78485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				return b != percentByte
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				// _ = "end of CoverTab[78485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:268
				_go_fuzz_dep_.CoverTab[78486]++
																sb.WriteByte(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:269
				// _ = "end of CoverTab[78486]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:270
				_go_fuzz_dep_.CoverTab[78487]++
																fmt.Fprintf(&sb, "%%%02X", b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:271
				// _ = "end of CoverTab[78487]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:272
			// _ = "end of CoverTab[78481]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:273
		// _ = "end of CoverTab[78478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:273
		_go_fuzz_dep_.CoverTab[78479]++
														msg = msg[size:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:274
		// _ = "end of CoverTab[78479]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:275
	// _ = "end of CoverTab[78476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:275
	_go_fuzz_dep_.CoverTab[78477]++
													return sb.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:276
	// _ = "end of CoverTab[78477]"
}

// decodeGrpcMessage decodes the msg encoded by encodeGrpcMessage.
func decodeGrpcMessage(msg string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:280
	_go_fuzz_dep_.CoverTab[78488]++
													if msg == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:281
		_go_fuzz_dep_.CoverTab[78491]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:282
		// _ = "end of CoverTab[78491]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:283
		_go_fuzz_dep_.CoverTab[78492]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:283
		// _ = "end of CoverTab[78492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:283
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:283
	// _ = "end of CoverTab[78488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:283
	_go_fuzz_dep_.CoverTab[78489]++
													lenMsg := len(msg)
													for i := 0; i < lenMsg; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:285
		_go_fuzz_dep_.CoverTab[78493]++
														if msg[i] == percentByte && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:286
			_go_fuzz_dep_.CoverTab[78494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:286
			return i+2 < lenMsg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:286
			// _ = "end of CoverTab[78494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:286
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:286
			_go_fuzz_dep_.CoverTab[78495]++
															return decodeGrpcMessageUnchecked(msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:287
			// _ = "end of CoverTab[78495]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:288
			_go_fuzz_dep_.CoverTab[78496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:288
			// _ = "end of CoverTab[78496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:288
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:288
		// _ = "end of CoverTab[78493]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:289
	// _ = "end of CoverTab[78489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:289
	_go_fuzz_dep_.CoverTab[78490]++
													return msg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:290
	// _ = "end of CoverTab[78490]"
}

func decodeGrpcMessageUnchecked(msg string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:293
	_go_fuzz_dep_.CoverTab[78497]++
													var sb strings.Builder
													lenMsg := len(msg)
													for i := 0; i < lenMsg; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:296
		_go_fuzz_dep_.CoverTab[78499]++
														c := msg[i]
														if c == percentByte && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:298
			_go_fuzz_dep_.CoverTab[78500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:298
			return i+2 < lenMsg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:298
			// _ = "end of CoverTab[78500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:298
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:298
			_go_fuzz_dep_.CoverTab[78501]++
															parsed, err := strconv.ParseUint(msg[i+1:i+3], 16, 8)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:300
				_go_fuzz_dep_.CoverTab[78502]++
																sb.WriteByte(c)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:301
				// _ = "end of CoverTab[78502]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:302
				_go_fuzz_dep_.CoverTab[78503]++
																sb.WriteByte(byte(parsed))
																i += 2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:304
				// _ = "end of CoverTab[78503]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:305
			// _ = "end of CoverTab[78501]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:306
			_go_fuzz_dep_.CoverTab[78504]++
															sb.WriteByte(c)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:307
			// _ = "end of CoverTab[78504]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:308
		// _ = "end of CoverTab[78499]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:309
	// _ = "end of CoverTab[78497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:309
	_go_fuzz_dep_.CoverTab[78498]++
													return sb.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:310
	// _ = "end of CoverTab[78498]"
}

type bufWriter struct {
	buf		[]byte
	offset		int
	batchSize	int
	conn		net.Conn
	err		error
}

func newBufWriter(conn net.Conn, batchSize int) *bufWriter {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:321
	_go_fuzz_dep_.CoverTab[78505]++
													return &bufWriter{
		buf:		make([]byte, batchSize*2),
		batchSize:	batchSize,
		conn:		conn,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:326
	// _ = "end of CoverTab[78505]"
}

func (w *bufWriter) Write(b []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:329
	_go_fuzz_dep_.CoverTab[78506]++
													if w.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:330
		_go_fuzz_dep_.CoverTab[78510]++
														return 0, w.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:331
		// _ = "end of CoverTab[78510]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:332
		_go_fuzz_dep_.CoverTab[78511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:332
		// _ = "end of CoverTab[78511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:332
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:332
	// _ = "end of CoverTab[78506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:332
	_go_fuzz_dep_.CoverTab[78507]++
													if w.batchSize == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:333
		_go_fuzz_dep_.CoverTab[78512]++
														n, err = w.conn.Write(b)
														return n, toIOError(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:335
		// _ = "end of CoverTab[78512]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:336
		_go_fuzz_dep_.CoverTab[78513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:336
		// _ = "end of CoverTab[78513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:336
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:336
	// _ = "end of CoverTab[78507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:336
	_go_fuzz_dep_.CoverTab[78508]++
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:337
		_go_fuzz_dep_.CoverTab[78514]++
														nn := copy(w.buf[w.offset:], b)
														b = b[nn:]
														w.offset += nn
														n += nn
														if w.offset >= w.batchSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:342
			_go_fuzz_dep_.CoverTab[78515]++
															err = w.Flush()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:343
			// _ = "end of CoverTab[78515]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:344
			_go_fuzz_dep_.CoverTab[78516]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:344
			// _ = "end of CoverTab[78516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:344
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:344
		// _ = "end of CoverTab[78514]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:345
	// _ = "end of CoverTab[78508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:345
	_go_fuzz_dep_.CoverTab[78509]++
													return n, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:346
	// _ = "end of CoverTab[78509]"
}

func (w *bufWriter) Flush() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:349
	_go_fuzz_dep_.CoverTab[78517]++
													if w.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:350
		_go_fuzz_dep_.CoverTab[78520]++
														return w.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:351
		// _ = "end of CoverTab[78520]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:352
		_go_fuzz_dep_.CoverTab[78521]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:352
		// _ = "end of CoverTab[78521]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:352
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:352
	// _ = "end of CoverTab[78517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:352
	_go_fuzz_dep_.CoverTab[78518]++
													if w.offset == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:353
		_go_fuzz_dep_.CoverTab[78522]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:354
		// _ = "end of CoverTab[78522]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:355
		_go_fuzz_dep_.CoverTab[78523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:355
		// _ = "end of CoverTab[78523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:355
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:355
	// _ = "end of CoverTab[78518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:355
	_go_fuzz_dep_.CoverTab[78519]++
													_, w.err = w.conn.Write(w.buf[:w.offset])
													w.err = toIOError(w.err)
													w.offset = 0
													return w.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:359
	// _ = "end of CoverTab[78519]"
}

type ioError struct {
	error
}

func (i ioError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:366
	_go_fuzz_dep_.CoverTab[78524]++
													return i.error
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:367
	// _ = "end of CoverTab[78524]"
}

func isIOError(err error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:370
	_go_fuzz_dep_.CoverTab[78525]++
													return errors.As(err, &ioError{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:371
	// _ = "end of CoverTab[78525]"
}

func toIOError(err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:374
	_go_fuzz_dep_.CoverTab[78526]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:375
		_go_fuzz_dep_.CoverTab[78528]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:376
		// _ = "end of CoverTab[78528]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:377
		_go_fuzz_dep_.CoverTab[78529]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:377
		// _ = "end of CoverTab[78529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:377
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:377
	// _ = "end of CoverTab[78526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:377
	_go_fuzz_dep_.CoverTab[78527]++
													return ioError{error: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:378
	// _ = "end of CoverTab[78527]"
}

type framer struct {
	writer	*bufWriter
	fr	*http2.Framer
}

func newFramer(conn net.Conn, writeBufferSize, readBufferSize int, maxHeaderListSize uint32) *framer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:386
	_go_fuzz_dep_.CoverTab[78530]++
													if writeBufferSize < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:387
		_go_fuzz_dep_.CoverTab[78533]++
														writeBufferSize = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:388
		// _ = "end of CoverTab[78533]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:389
		_go_fuzz_dep_.CoverTab[78534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:389
		// _ = "end of CoverTab[78534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:389
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:389
	// _ = "end of CoverTab[78530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:389
	_go_fuzz_dep_.CoverTab[78531]++
													var r io.Reader = conn
													if readBufferSize > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:391
		_go_fuzz_dep_.CoverTab[78535]++
														r = bufio.NewReaderSize(r, readBufferSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:392
		// _ = "end of CoverTab[78535]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:393
		_go_fuzz_dep_.CoverTab[78536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:393
		// _ = "end of CoverTab[78536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:393
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:393
	// _ = "end of CoverTab[78531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:393
	_go_fuzz_dep_.CoverTab[78532]++
													w := newBufWriter(conn, writeBufferSize)
													f := &framer{
		writer:	w,
		fr:	http2.NewFramer(w, r),
	}
													f.fr.SetMaxReadFrameSize(http2MaxFrameLen)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:402
	f.fr.SetReuseFrames()
													f.fr.MaxHeaderListSize = maxHeaderListSize
													f.fr.ReadMetaHeaders = hpack.NewDecoder(http2InitHeaderTableSize, nil)
													return f
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:405
	// _ = "end of CoverTab[78532]"
}

// parseDialTarget returns the network and address to pass to dialer.
func parseDialTarget(target string) (string, string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:409
	_go_fuzz_dep_.CoverTab[78537]++
													net := "tcp"
													m1 := strings.Index(target, ":")
													m2 := strings.Index(target, ":/")

													if m1 >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:414
		_go_fuzz_dep_.CoverTab[78540]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:414
		return m2 < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:414
		// _ = "end of CoverTab[78540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:414
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:414
		_go_fuzz_dep_.CoverTab[78541]++
														if n := target[0:m1]; n == "unix" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:415
			_go_fuzz_dep_.CoverTab[78542]++
															return n, target[m1+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:416
			// _ = "end of CoverTab[78542]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:417
			_go_fuzz_dep_.CoverTab[78543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:417
			// _ = "end of CoverTab[78543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:417
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:417
		// _ = "end of CoverTab[78541]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:418
		_go_fuzz_dep_.CoverTab[78544]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:418
		// _ = "end of CoverTab[78544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:418
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:418
	// _ = "end of CoverTab[78537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:418
	_go_fuzz_dep_.CoverTab[78538]++
													if m2 >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:419
		_go_fuzz_dep_.CoverTab[78545]++
														t, err := url.Parse(target)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:421
			_go_fuzz_dep_.CoverTab[78547]++
															return net, target
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:422
			// _ = "end of CoverTab[78547]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:423
			_go_fuzz_dep_.CoverTab[78548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:423
			// _ = "end of CoverTab[78548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:423
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:423
		// _ = "end of CoverTab[78545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:423
		_go_fuzz_dep_.CoverTab[78546]++
														scheme := t.Scheme
														addr := t.Path
														if scheme == "unix" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:426
			_go_fuzz_dep_.CoverTab[78549]++
															if addr == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:427
				_go_fuzz_dep_.CoverTab[78551]++
																addr = t.Host
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:428
				// _ = "end of CoverTab[78551]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:429
				_go_fuzz_dep_.CoverTab[78552]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:429
				// _ = "end of CoverTab[78552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:429
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:429
			// _ = "end of CoverTab[78549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:429
			_go_fuzz_dep_.CoverTab[78550]++
															return scheme, addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:430
			// _ = "end of CoverTab[78550]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:431
			_go_fuzz_dep_.CoverTab[78553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:431
			// _ = "end of CoverTab[78553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:431
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:431
		// _ = "end of CoverTab[78546]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:432
		_go_fuzz_dep_.CoverTab[78554]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:432
		// _ = "end of CoverTab[78554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:432
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:432
	// _ = "end of CoverTab[78538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:432
	_go_fuzz_dep_.CoverTab[78539]++
													return net, target
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:433
	// _ = "end of CoverTab[78539]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:434
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http_util.go:434
var _ = _go_fuzz_dep_.CoverTab
