//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
package codes

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:19
)

import (
	"strconv"

	"google.golang.org/grpc/internal"
)

func init() {
	internal.CanonicalString = canonicalString
}

func (c Code) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:31
	_go_fuzz_dep_.CoverTab[67517]++
												switch c {
	case OK:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:33
		_go_fuzz_dep_.CoverTab[67518]++
													return "OK"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:34
		// _ = "end of CoverTab[67518]"
	case Canceled:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:35
		_go_fuzz_dep_.CoverTab[67519]++
													return "Canceled"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:36
		// _ = "end of CoverTab[67519]"
	case Unknown:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:37
		_go_fuzz_dep_.CoverTab[67520]++
													return "Unknown"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:38
		// _ = "end of CoverTab[67520]"
	case InvalidArgument:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:39
		_go_fuzz_dep_.CoverTab[67521]++
													return "InvalidArgument"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:40
		// _ = "end of CoverTab[67521]"
	case DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:41
		_go_fuzz_dep_.CoverTab[67522]++
													return "DeadlineExceeded"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:42
		// _ = "end of CoverTab[67522]"
	case NotFound:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:43
		_go_fuzz_dep_.CoverTab[67523]++
													return "NotFound"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:44
		// _ = "end of CoverTab[67523]"
	case AlreadyExists:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:45
		_go_fuzz_dep_.CoverTab[67524]++
													return "AlreadyExists"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:46
		// _ = "end of CoverTab[67524]"
	case PermissionDenied:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:47
		_go_fuzz_dep_.CoverTab[67525]++
													return "PermissionDenied"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:48
		// _ = "end of CoverTab[67525]"
	case ResourceExhausted:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:49
		_go_fuzz_dep_.CoverTab[67526]++
													return "ResourceExhausted"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:50
		// _ = "end of CoverTab[67526]"
	case FailedPrecondition:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:51
		_go_fuzz_dep_.CoverTab[67527]++
													return "FailedPrecondition"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:52
		// _ = "end of CoverTab[67527]"
	case Aborted:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:53
		_go_fuzz_dep_.CoverTab[67528]++
													return "Aborted"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:54
		// _ = "end of CoverTab[67528]"
	case OutOfRange:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:55
		_go_fuzz_dep_.CoverTab[67529]++
													return "OutOfRange"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:56
		// _ = "end of CoverTab[67529]"
	case Unimplemented:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:57
		_go_fuzz_dep_.CoverTab[67530]++
													return "Unimplemented"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:58
		// _ = "end of CoverTab[67530]"
	case Internal:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:59
		_go_fuzz_dep_.CoverTab[67531]++
													return "Internal"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:60
		// _ = "end of CoverTab[67531]"
	case Unavailable:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:61
		_go_fuzz_dep_.CoverTab[67532]++
													return "Unavailable"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:62
		// _ = "end of CoverTab[67532]"
	case DataLoss:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:63
		_go_fuzz_dep_.CoverTab[67533]++
													return "DataLoss"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:64
		// _ = "end of CoverTab[67533]"
	case Unauthenticated:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:65
		_go_fuzz_dep_.CoverTab[67534]++
													return "Unauthenticated"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:66
		// _ = "end of CoverTab[67534]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:67
		_go_fuzz_dep_.CoverTab[67535]++
													return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:68
		// _ = "end of CoverTab[67535]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:69
	// _ = "end of CoverTab[67517]"
}

func canonicalString(c Code) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:72
	_go_fuzz_dep_.CoverTab[67536]++
												switch c {
	case OK:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:74
		_go_fuzz_dep_.CoverTab[67537]++
													return "OK"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:75
		// _ = "end of CoverTab[67537]"
	case Canceled:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:76
		_go_fuzz_dep_.CoverTab[67538]++
													return "CANCELLED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:77
		// _ = "end of CoverTab[67538]"
	case Unknown:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:78
		_go_fuzz_dep_.CoverTab[67539]++
													return "UNKNOWN"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:79
		// _ = "end of CoverTab[67539]"
	case InvalidArgument:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:80
		_go_fuzz_dep_.CoverTab[67540]++
													return "INVALID_ARGUMENT"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:81
		// _ = "end of CoverTab[67540]"
	case DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:82
		_go_fuzz_dep_.CoverTab[67541]++
													return "DEADLINE_EXCEEDED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:83
		// _ = "end of CoverTab[67541]"
	case NotFound:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:84
		_go_fuzz_dep_.CoverTab[67542]++
													return "NOT_FOUND"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:85
		// _ = "end of CoverTab[67542]"
	case AlreadyExists:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:86
		_go_fuzz_dep_.CoverTab[67543]++
													return "ALREADY_EXISTS"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:87
		// _ = "end of CoverTab[67543]"
	case PermissionDenied:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:88
		_go_fuzz_dep_.CoverTab[67544]++
													return "PERMISSION_DENIED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:89
		// _ = "end of CoverTab[67544]"
	case ResourceExhausted:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:90
		_go_fuzz_dep_.CoverTab[67545]++
													return "RESOURCE_EXHAUSTED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:91
		// _ = "end of CoverTab[67545]"
	case FailedPrecondition:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:92
		_go_fuzz_dep_.CoverTab[67546]++
													return "FAILED_PRECONDITION"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:93
		// _ = "end of CoverTab[67546]"
	case Aborted:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:94
		_go_fuzz_dep_.CoverTab[67547]++
													return "ABORTED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:95
		// _ = "end of CoverTab[67547]"
	case OutOfRange:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:96
		_go_fuzz_dep_.CoverTab[67548]++
													return "OUT_OF_RANGE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:97
		// _ = "end of CoverTab[67548]"
	case Unimplemented:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:98
		_go_fuzz_dep_.CoverTab[67549]++
													return "UNIMPLEMENTED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:99
		// _ = "end of CoverTab[67549]"
	case Internal:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:100
		_go_fuzz_dep_.CoverTab[67550]++
													return "INTERNAL"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:101
		// _ = "end of CoverTab[67550]"
	case Unavailable:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:102
		_go_fuzz_dep_.CoverTab[67551]++
													return "UNAVAILABLE"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:103
		// _ = "end of CoverTab[67551]"
	case DataLoss:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:104
		_go_fuzz_dep_.CoverTab[67552]++
													return "DATA_LOSS"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:105
		// _ = "end of CoverTab[67552]"
	case Unauthenticated:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:106
		_go_fuzz_dep_.CoverTab[67553]++
													return "UNAUTHENTICATED"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:107
		// _ = "end of CoverTab[67553]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:108
		_go_fuzz_dep_.CoverTab[67554]++
													return "CODE(" + strconv.FormatInt(int64(c), 10) + ")"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:109
		// _ = "end of CoverTab[67554]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:110
	// _ = "end of CoverTab[67536]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:111
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/codes/code_string.go:111
var _ = _go_fuzz_dep_.CoverTab
