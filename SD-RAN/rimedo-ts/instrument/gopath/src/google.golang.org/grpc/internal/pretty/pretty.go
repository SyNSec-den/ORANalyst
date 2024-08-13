//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:19
// Package pretty defines helper functions to pretty-print structs for logging.
package pretty

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:20
)

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	protov1 "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	protov2 "google.golang.org/protobuf/proto"
)

const jsonIndent = "  "

// ToJSON marshals the input into a json string.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:35
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:35
// If marshal fails, it falls back to fmt.Sprintf("%+v").
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:38
func ToJSON(e interface{}) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:38
	_go_fuzz_dep_.CoverTab[67260]++
												switch ee := e.(type) {
	case protov1.Message:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:40
		_go_fuzz_dep_.CoverTab[67261]++
													mm := jsonpb.Marshaler{Indent: jsonIndent}
													ret, err := mm.MarshalToString(ee)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:43
			_go_fuzz_dep_.CoverTab[67267]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:47
			return fmt.Sprintf("%+v", ee)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:47
			// _ = "end of CoverTab[67267]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:48
			_go_fuzz_dep_.CoverTab[67268]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:48
			// _ = "end of CoverTab[67268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:48
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:48
		// _ = "end of CoverTab[67261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:48
		_go_fuzz_dep_.CoverTab[67262]++
													return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:49
		// _ = "end of CoverTab[67262]"
	case protov2.Message:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:50
		_go_fuzz_dep_.CoverTab[67263]++
													mm := protojson.MarshalOptions{
			Multiline:	true,
			Indent:		jsonIndent,
		}
		ret, err := mm.Marshal(ee)
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:56
			_go_fuzz_dep_.CoverTab[67269]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:60
			return fmt.Sprintf("%+v", ee)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:60
			// _ = "end of CoverTab[67269]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:61
			_go_fuzz_dep_.CoverTab[67270]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:61
			// _ = "end of CoverTab[67270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:61
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:61
		// _ = "end of CoverTab[67263]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:61
		_go_fuzz_dep_.CoverTab[67264]++
													return string(ret)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:62
		// _ = "end of CoverTab[67264]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:63
		_go_fuzz_dep_.CoverTab[67265]++
													ret, err := json.MarshalIndent(ee, "", jsonIndent)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:65
			_go_fuzz_dep_.CoverTab[67271]++
														return fmt.Sprintf("%+v", ee)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:66
			// _ = "end of CoverTab[67271]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:67
			_go_fuzz_dep_.CoverTab[67272]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:67
			// _ = "end of CoverTab[67272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:67
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:67
		// _ = "end of CoverTab[67265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:67
		_go_fuzz_dep_.CoverTab[67266]++
													return string(ret)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:68
		// _ = "end of CoverTab[67266]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:69
	// _ = "end of CoverTab[67260]"
}

// FormatJSON formats the input json bytes with indentation.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:72
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:72
// If Indent fails, it returns the unchanged input as string.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:75
func FormatJSON(b []byte) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:75
	_go_fuzz_dep_.CoverTab[67273]++
												var out bytes.Buffer
												err := json.Indent(&out, b, "", jsonIndent)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:78
		_go_fuzz_dep_.CoverTab[67275]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:79
		// _ = "end of CoverTab[67275]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:80
		_go_fuzz_dep_.CoverTab[67276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:80
		// _ = "end of CoverTab[67276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:80
	// _ = "end of CoverTab[67273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:80
	_go_fuzz_dep_.CoverTab[67274]++
												return out.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:81
	// _ = "end of CoverTab[67274]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/pretty/pretty.go:82
var _ = _go_fuzz_dep_.CoverTab
