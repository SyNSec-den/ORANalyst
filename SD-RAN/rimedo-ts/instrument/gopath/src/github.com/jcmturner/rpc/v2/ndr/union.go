//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:1
)

import (
	"errors"
	"fmt"
	"reflect"
)

// Union interface must be implemented by structs that will be unmarshaled into from the NDR byte stream union representation.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:9
// The union's discriminating tag will be passed to the SwitchFunc method.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:9
// The discriminating tag field must have the struct tag: `ndr:"unionTag"`
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:9
// If the union is encapsulated the discriminating tag field must have the struct tag: `ndr:"encapsulated"`
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:9
// The possible value fields that can be selected from must have the struct tag: `ndr:"unionField"`
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:14
type Union interface {
	SwitchFunc(t interface{}) string
}

// Union related constants such as struct tag values
const (
	unionSelectionFuncName	= "SwitchFunc"
	TagEncapsulated		= "encapsulated"
	TagUnionTag		= "unionTag"
	TagUnionField		= "unionField"
)

func (dec *Decoder) isUnion(field reflect.Value, tag reflect.StructTag) (r reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:26
	_go_fuzz_dep_.CoverTab[87316]++
											ndrTag := parseTags(tag)
											if !ndrTag.HasValue(TagUnionTag) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:28
		_go_fuzz_dep_.CoverTab[87319]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:29
		// _ = "end of CoverTab[87319]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:30
		_go_fuzz_dep_.CoverTab[87320]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:30
		// _ = "end of CoverTab[87320]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:30
	// _ = "end of CoverTab[87316]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:30
	_go_fuzz_dep_.CoverTab[87317]++
											r = field

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:35
	if !ndrTag.HasValue(TagEncapsulated) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:35
		_go_fuzz_dep_.CoverTab[87321]++
												dec.r.Discard(int(r.Type().Size()))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:36
		// _ = "end of CoverTab[87321]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:37
		_go_fuzz_dep_.CoverTab[87322]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:37
		// _ = "end of CoverTab[87322]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:37
	// _ = "end of CoverTab[87317]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:37
	_go_fuzz_dep_.CoverTab[87318]++
											return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:38
	// _ = "end of CoverTab[87318]"
}

// unionSelectedField returns the field name of which of the union values to fill
func unionSelectedField(union, discriminant reflect.Value) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:42
	_go_fuzz_dep_.CoverTab[87323]++
											if !union.Type().Implements(reflect.TypeOf(new(Union)).Elem()) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:43
		_go_fuzz_dep_.CoverTab[87327]++
												return "", errors.New("struct does not implement union interface")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:44
		// _ = "end of CoverTab[87327]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:45
		_go_fuzz_dep_.CoverTab[87328]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:45
		// _ = "end of CoverTab[87328]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:45
	// _ = "end of CoverTab[87323]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:45
	_go_fuzz_dep_.CoverTab[87324]++
											args := []reflect.Value{discriminant}

											sf := union.MethodByName(unionSelectionFuncName)
											if !sf.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:49
		_go_fuzz_dep_.CoverTab[87329]++
												return "", fmt.Errorf("could not find a selection function called %s in the unions struct representation", unionSelectionFuncName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:50
		// _ = "end of CoverTab[87329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:51
		_go_fuzz_dep_.CoverTab[87330]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:51
		// _ = "end of CoverTab[87330]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:51
	// _ = "end of CoverTab[87324]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:51
	_go_fuzz_dep_.CoverTab[87325]++
											f := sf.Call(args)
											if f[0].Kind() != reflect.String || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:53
		_go_fuzz_dep_.CoverTab[87331]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:53
		return f[0].String() == ""
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:53
		// _ = "end of CoverTab[87331]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:53
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:53
		_go_fuzz_dep_.CoverTab[87332]++
												return "", fmt.Errorf("the union select function did not return a string for the name of the field to fill")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:54
		// _ = "end of CoverTab[87332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:55
		_go_fuzz_dep_.CoverTab[87333]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:55
		// _ = "end of CoverTab[87333]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:55
	// _ = "end of CoverTab[87325]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:55
	_go_fuzz_dep_.CoverTab[87326]++
											return f[0].String(), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:56
	// _ = "end of CoverTab[87326]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/union.go:57
var _ = _go_fuzz_dep_.CoverTab
