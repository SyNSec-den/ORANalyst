//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:19
// Package metadata define the structure of the metadata supported by gRPC library.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:19
// Please refer to https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:19
// for more information about custom-metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
package metadata

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:22
)

import (
	"context"
	"fmt"
	"strings"
)

// DecodeKeyValue returns k, v, nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:30
// Deprecated: use k and v directly instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:33
func DecodeKeyValue(k, v string) (string, string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:33
	_go_fuzz_dep_.CoverTab[67331]++
												return k, v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:34
	// _ = "end of CoverTab[67331]"
}

// MD is a mapping from metadata keys to values. Users should use the following
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:37
// two convenience functions New and Pairs to generate MD.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:39
type MD map[string][]string

// New creates an MD from a given key-value map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
// Only the following ASCII characters are allowed in keys:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//   - digits: 0-9
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//   - uppercase letters: A-Z (normalized to lower)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//   - lowercase letters: a-z
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//   - special characters: -_.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
// Uppercase letters are automatically converted to lowercase.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
// Keys beginning with "grpc-" are reserved for grpc-internal use only and may
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:41
// result in errors if set in metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:53
func New(m map[string]string) MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:53
	_go_fuzz_dep_.CoverTab[67332]++
												md := make(MD, len(m))
												for k, val := range m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:55
		_go_fuzz_dep_.CoverTab[67334]++
													key := strings.ToLower(k)
													md[key] = append(md[key], val)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:57
		// _ = "end of CoverTab[67334]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:58
	// _ = "end of CoverTab[67332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:58
	_go_fuzz_dep_.CoverTab[67333]++
												return md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:59
	// _ = "end of CoverTab[67333]"
}

// Pairs returns an MD formed by the mapping of key, value ...
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
// Pairs panics if len(kv) is odd.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
// Only the following ASCII characters are allowed in keys:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//   - digits: 0-9
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//   - uppercase letters: A-Z (normalized to lower)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//   - lowercase letters: a-z
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//   - special characters: -_.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
// Uppercase letters are automatically converted to lowercase.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
// Keys beginning with "grpc-" are reserved for grpc-internal use only and may
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:62
// result in errors if set in metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:75
func Pairs(kv ...string) MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:75
	_go_fuzz_dep_.CoverTab[67335]++
												if len(kv)%2 == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:76
		_go_fuzz_dep_.CoverTab[67338]++
													panic(fmt.Sprintf("metadata: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:77
		// _ = "end of CoverTab[67338]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:78
		_go_fuzz_dep_.CoverTab[67339]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:78
		// _ = "end of CoverTab[67339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:78
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:78
	// _ = "end of CoverTab[67335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:78
	_go_fuzz_dep_.CoverTab[67336]++
												md := make(MD, len(kv)/2)
												for i := 0; i < len(kv); i += 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:80
		_go_fuzz_dep_.CoverTab[67340]++
													key := strings.ToLower(kv[i])
													md[key] = append(md[key], kv[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:82
		// _ = "end of CoverTab[67340]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:83
	// _ = "end of CoverTab[67336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:83
	_go_fuzz_dep_.CoverTab[67337]++
												return md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:84
	// _ = "end of CoverTab[67337]"
}

// Len returns the number of items in md.
func (md MD) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:88
	_go_fuzz_dep_.CoverTab[67341]++
												return len(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:89
	// _ = "end of CoverTab[67341]"
}

// Copy returns a copy of md.
func (md MD) Copy() MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:93
	_go_fuzz_dep_.CoverTab[67342]++
												out := make(MD, len(md))
												for k, v := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:95
		_go_fuzz_dep_.CoverTab[67344]++
													out[k] = copyOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:96
		// _ = "end of CoverTab[67344]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:97
	// _ = "end of CoverTab[67342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:97
	_go_fuzz_dep_.CoverTab[67343]++
												return out
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:98
	// _ = "end of CoverTab[67343]"
}

// Get obtains the values for a given key.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:101
// k is converted to lowercase before searching in md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:104
func (md MD) Get(k string) []string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:104
	_go_fuzz_dep_.CoverTab[67345]++
												k = strings.ToLower(k)
												return md[k]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:106
	// _ = "end of CoverTab[67345]"
}

// Set sets the value of a given key with a slice of values.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:109
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:109
// k is converted to lowercase before storing in md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:112
func (md MD) Set(k string, vals ...string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:112
	_go_fuzz_dep_.CoverTab[67346]++
												if len(vals) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:113
		_go_fuzz_dep_.CoverTab[67348]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:114
		// _ = "end of CoverTab[67348]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:115
		_go_fuzz_dep_.CoverTab[67349]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:115
		// _ = "end of CoverTab[67349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:115
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:115
	// _ = "end of CoverTab[67346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:115
	_go_fuzz_dep_.CoverTab[67347]++
												k = strings.ToLower(k)
												md[k] = vals
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:117
	// _ = "end of CoverTab[67347]"
}

// Append adds the values to key k, not overwriting what was already stored at
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:120
// that key.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:120
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:120
// k is converted to lowercase before storing in md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:124
func (md MD) Append(k string, vals ...string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:124
	_go_fuzz_dep_.CoverTab[67350]++
												if len(vals) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:125
		_go_fuzz_dep_.CoverTab[67352]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:126
		// _ = "end of CoverTab[67352]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:127
		_go_fuzz_dep_.CoverTab[67353]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:127
		// _ = "end of CoverTab[67353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:127
	// _ = "end of CoverTab[67350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:127
	_go_fuzz_dep_.CoverTab[67351]++
												k = strings.ToLower(k)
												md[k] = append(md[k], vals...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:129
	// _ = "end of CoverTab[67351]"
}

// Delete removes the values for a given key k which is converted to lowercase
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:132
// before removing it from md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:134
func (md MD) Delete(k string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:134
	_go_fuzz_dep_.CoverTab[67354]++
												k = strings.ToLower(k)
												delete(md, k)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:136
	// _ = "end of CoverTab[67354]"
}

// Join joins any number of mds into a single MD.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:139
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:139
// The order of values for each key is determined by the order in which the mds
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:139
// containing those values are presented to Join.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:143
func Join(mds ...MD) MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:143
	_go_fuzz_dep_.CoverTab[67355]++
												out := MD{}
												for _, md := range mds {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:145
		_go_fuzz_dep_.CoverTab[67357]++
													for k, v := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:146
			_go_fuzz_dep_.CoverTab[67358]++
														out[k] = append(out[k], v...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:147
			// _ = "end of CoverTab[67358]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:148
		// _ = "end of CoverTab[67357]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:149
	// _ = "end of CoverTab[67355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:149
	_go_fuzz_dep_.CoverTab[67356]++
												return out
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:150
	// _ = "end of CoverTab[67356]"
}

type mdIncomingKey struct{}
type mdOutgoingKey struct{}

// NewIncomingContext creates a new context with incoming md attached.
func NewIncomingContext(ctx context.Context, md MD) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:157
	_go_fuzz_dep_.CoverTab[67359]++
												return context.WithValue(ctx, mdIncomingKey{}, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:158
	// _ = "end of CoverTab[67359]"
}

// NewOutgoingContext creates a new context with outgoing md attached. If used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:161
// in conjunction with AppendToOutgoingContext, NewOutgoingContext will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:161
// overwrite any previously-appended metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:164
func NewOutgoingContext(ctx context.Context, md MD) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:164
	_go_fuzz_dep_.CoverTab[67360]++
												return context.WithValue(ctx, mdOutgoingKey{}, rawMD{md: md})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:165
	// _ = "end of CoverTab[67360]"
}

// AppendToOutgoingContext returns a new context with the provided kv merged
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:168
// with any existing metadata in the context. Please refer to the documentation
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:168
// of Pairs for a description of kv.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:171
func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:171
	_go_fuzz_dep_.CoverTab[67361]++
												if len(kv)%2 == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:172
		_go_fuzz_dep_.CoverTab[67364]++
													panic(fmt.Sprintf("metadata: AppendToOutgoingContext got an odd number of input pairs for metadata: %d", len(kv)))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:173
		// _ = "end of CoverTab[67364]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:174
		_go_fuzz_dep_.CoverTab[67365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:174
		// _ = "end of CoverTab[67365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:174
	// _ = "end of CoverTab[67361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:174
	_go_fuzz_dep_.CoverTab[67362]++
												md, _ := ctx.Value(mdOutgoingKey{}).(rawMD)
												added := make([][]string, len(md.added)+1)
												copy(added, md.added)
												kvCopy := make([]string, 0, len(kv))
												for i := 0; i < len(kv); i += 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:179
		_go_fuzz_dep_.CoverTab[67366]++
													kvCopy = append(kvCopy, strings.ToLower(kv[i]), kv[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:180
		// _ = "end of CoverTab[67366]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:181
	// _ = "end of CoverTab[67362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:181
	_go_fuzz_dep_.CoverTab[67363]++
												added[len(added)-1] = kvCopy
												return context.WithValue(ctx, mdOutgoingKey{}, rawMD{md: md.md, added: added})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:183
	// _ = "end of CoverTab[67363]"
}

// FromIncomingContext returns the incoming metadata in ctx if it exists.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:186
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:186
// All keys in the returned MD are lowercase.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:189
func FromIncomingContext(ctx context.Context) (MD, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:189
	_go_fuzz_dep_.CoverTab[67367]++
												md, ok := ctx.Value(mdIncomingKey{}).(MD)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:191
		_go_fuzz_dep_.CoverTab[67370]++
													return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:192
		// _ = "end of CoverTab[67370]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:193
		_go_fuzz_dep_.CoverTab[67371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:193
		// _ = "end of CoverTab[67371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:193
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:193
	// _ = "end of CoverTab[67367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:193
	_go_fuzz_dep_.CoverTab[67368]++
												out := make(MD, len(md))
												for k, v := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:195
		_go_fuzz_dep_.CoverTab[67372]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:199
		key := strings.ToLower(k)
													out[key] = copyOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:200
		// _ = "end of CoverTab[67372]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:201
	// _ = "end of CoverTab[67368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:201
	_go_fuzz_dep_.CoverTab[67369]++
												return out, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:202
	// _ = "end of CoverTab[67369]"
}

// ValueFromIncomingContext returns the metadata value corresponding to the metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
// key from the incoming metadata if it exists. Key must be lower-case.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:205
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:212
func ValueFromIncomingContext(ctx context.Context, key string) []string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:212
	_go_fuzz_dep_.CoverTab[67373]++
												md, ok := ctx.Value(mdIncomingKey{}).(MD)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:214
		_go_fuzz_dep_.CoverTab[67377]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:215
		// _ = "end of CoverTab[67377]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:216
		_go_fuzz_dep_.CoverTab[67378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:216
		// _ = "end of CoverTab[67378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:216
	// _ = "end of CoverTab[67373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:216
	_go_fuzz_dep_.CoverTab[67374]++

												if v, ok := md[key]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:218
		_go_fuzz_dep_.CoverTab[67379]++
													return copyOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:219
		// _ = "end of CoverTab[67379]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:220
		_go_fuzz_dep_.CoverTab[67380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:220
		// _ = "end of CoverTab[67380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:220
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:220
	// _ = "end of CoverTab[67374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:220
	_go_fuzz_dep_.CoverTab[67375]++
												for k, v := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:221
		_go_fuzz_dep_.CoverTab[67381]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:225
		if strings.ToLower(k) == key {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:225
			_go_fuzz_dep_.CoverTab[67382]++
														return copyOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:226
			// _ = "end of CoverTab[67382]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:227
			_go_fuzz_dep_.CoverTab[67383]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:227
			// _ = "end of CoverTab[67383]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:227
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:227
		// _ = "end of CoverTab[67381]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:228
	// _ = "end of CoverTab[67375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:228
	_go_fuzz_dep_.CoverTab[67376]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:229
	// _ = "end of CoverTab[67376]"
}

// the returned slice must not be modified in place
func copyOf(v []string) []string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:233
	_go_fuzz_dep_.CoverTab[67384]++
												vals := make([]string, len(v))
												copy(vals, v)
												return vals
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:236
	// _ = "end of CoverTab[67384]"
}

// FromOutgoingContextRaw returns the un-merged, intermediary contents of rawMD.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// Remember to perform strings.ToLower on the keys, for both the returned MD (MD
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// is a map, there's no guarantee it's created using our helper functions) and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// the extra kv pairs (AppendToOutgoingContext doesn't turn them into
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// lowercase).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// This is intended for gRPC-internal use ONLY. Users should use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:239
// FromOutgoingContext instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:248
func FromOutgoingContextRaw(ctx context.Context) (MD, [][]string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:248
	_go_fuzz_dep_.CoverTab[67385]++
												raw, ok := ctx.Value(mdOutgoingKey{}).(rawMD)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:250
		_go_fuzz_dep_.CoverTab[67387]++
													return nil, nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:251
		// _ = "end of CoverTab[67387]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:252
		_go_fuzz_dep_.CoverTab[67388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:252
		// _ = "end of CoverTab[67388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:252
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:252
	// _ = "end of CoverTab[67385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:252
	_go_fuzz_dep_.CoverTab[67386]++

												return raw.md, raw.added, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:254
	// _ = "end of CoverTab[67386]"
}

// FromOutgoingContext returns the outgoing metadata in ctx if it exists.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:257
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:257
// All keys in the returned MD are lowercase.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:260
func FromOutgoingContext(ctx context.Context) (MD, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:260
	_go_fuzz_dep_.CoverTab[67389]++
												raw, ok := ctx.Value(mdOutgoingKey{}).(rawMD)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:262
		_go_fuzz_dep_.CoverTab[67394]++
													return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:263
		// _ = "end of CoverTab[67394]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:264
		_go_fuzz_dep_.CoverTab[67395]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:264
		// _ = "end of CoverTab[67395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:264
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:264
	// _ = "end of CoverTab[67389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:264
	_go_fuzz_dep_.CoverTab[67390]++

												mdSize := len(raw.md)
												for i := range raw.added {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:267
		_go_fuzz_dep_.CoverTab[67396]++
													mdSize += len(raw.added[i]) / 2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:268
		// _ = "end of CoverTab[67396]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:269
	// _ = "end of CoverTab[67390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:269
	_go_fuzz_dep_.CoverTab[67391]++

												out := make(MD, mdSize)
												for k, v := range raw.md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:272
		_go_fuzz_dep_.CoverTab[67397]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:276
		key := strings.ToLower(k)
													out[key] = copyOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:277
		// _ = "end of CoverTab[67397]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:278
	// _ = "end of CoverTab[67391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:278
	_go_fuzz_dep_.CoverTab[67392]++
												for _, added := range raw.added {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:279
		_go_fuzz_dep_.CoverTab[67398]++
													if len(added)%2 == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:280
			_go_fuzz_dep_.CoverTab[67400]++
														panic(fmt.Sprintf("metadata: FromOutgoingContext got an odd number of input pairs for metadata: %d", len(added)))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:281
			// _ = "end of CoverTab[67400]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:282
			_go_fuzz_dep_.CoverTab[67401]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:282
			// _ = "end of CoverTab[67401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:282
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:282
		// _ = "end of CoverTab[67398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:282
		_go_fuzz_dep_.CoverTab[67399]++

													for i := 0; i < len(added); i += 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:284
			_go_fuzz_dep_.CoverTab[67402]++
														key := strings.ToLower(added[i])
														out[key] = append(out[key], added[i+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:286
			// _ = "end of CoverTab[67402]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:287
		// _ = "end of CoverTab[67399]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:288
	// _ = "end of CoverTab[67392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:288
	_go_fuzz_dep_.CoverTab[67393]++
												return out, ok
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:289
	// _ = "end of CoverTab[67393]"
}

type rawMD struct {
	md	MD
	added	[][]string
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:295
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/metadata/metadata.go:295
var _ = _go_fuzz_dep_.CoverTab
