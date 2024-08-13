//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:1
)

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:13
const (
	sizeMethod = "Size"
)

// RawBytes interface should be implemented if reading just a number of bytes from the NDR stream
type RawBytes interface {
	Size(interface{}) int
}

func rawBytesSize(parent reflect.Value, v reflect.Value) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:22
	_go_fuzz_dep_.CoverTab[87249]++
												sf := v.MethodByName(sizeMethod)
												if !sf.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:24
		_go_fuzz_dep_.CoverTab[87252]++
													return 0, fmt.Errorf("could not find a method called %s on the implementation of RawBytes", sizeMethod)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:25
		// _ = "end of CoverTab[87252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:26
		_go_fuzz_dep_.CoverTab[87253]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:26
		// _ = "end of CoverTab[87253]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:26
	// _ = "end of CoverTab[87249]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:26
	_go_fuzz_dep_.CoverTab[87250]++
												in := []reflect.Value{parent}
												f := sf.Call(in)
												if f[0].Kind() != reflect.Int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:29
		_go_fuzz_dep_.CoverTab[87254]++
													return 0, errors.New("the RawBytes size function did not return an integer")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:30
		// _ = "end of CoverTab[87254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:31
		_go_fuzz_dep_.CoverTab[87255]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:31
		// _ = "end of CoverTab[87255]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:31
	// _ = "end of CoverTab[87250]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:31
	_go_fuzz_dep_.CoverTab[87251]++
												return int(f[0].Int()), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:32
	// _ = "end of CoverTab[87251]"
}

func addSizeToTag(parent reflect.Value, v reflect.Value, tag reflect.StructTag) (reflect.StructTag, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:35
	_go_fuzz_dep_.CoverTab[87256]++
												size, err := rawBytesSize(parent, v)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:37
		_go_fuzz_dep_.CoverTab[87258]++
													return tag, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:38
		// _ = "end of CoverTab[87258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:39
		_go_fuzz_dep_.CoverTab[87259]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:39
		// _ = "end of CoverTab[87259]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:39
	// _ = "end of CoverTab[87256]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:39
	_go_fuzz_dep_.CoverTab[87257]++
												ndrTag := parseTags(tag)
												ndrTag.Map["size"] = strconv.Itoa(size)
												return ndrTag.StructTag(), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:42
	// _ = "end of CoverTab[87257]"
}

func (dec *Decoder) readRawBytes(v reflect.Value, tag reflect.StructTag) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:45
	_go_fuzz_dep_.CoverTab[87260]++
												ndrTag := parseTags(tag)
												sizeStr, ok := ndrTag.Map["size"]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:48
		_go_fuzz_dep_.CoverTab[87264]++
													return errors.New("size tag not available")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:49
		// _ = "end of CoverTab[87264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:50
		_go_fuzz_dep_.CoverTab[87265]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:50
		// _ = "end of CoverTab[87265]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:50
	// _ = "end of CoverTab[87260]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:50
	_go_fuzz_dep_.CoverTab[87261]++
												size, err := strconv.Atoi(sizeStr)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:52
		_go_fuzz_dep_.CoverTab[87266]++
													return fmt.Errorf("size not valid: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:53
		// _ = "end of CoverTab[87266]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:54
		_go_fuzz_dep_.CoverTab[87267]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:54
		// _ = "end of CoverTab[87267]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:54
	// _ = "end of CoverTab[87261]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:54
	_go_fuzz_dep_.CoverTab[87262]++
												b, err := dec.readBytes(size)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:56
		_go_fuzz_dep_.CoverTab[87268]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:57
		// _ = "end of CoverTab[87268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:58
		_go_fuzz_dep_.CoverTab[87269]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:58
		// _ = "end of CoverTab[87269]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:58
	// _ = "end of CoverTab[87262]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:58
	_go_fuzz_dep_.CoverTab[87263]++
												v.Set(reflect.ValueOf(b).Convert(v.Type()))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:60
	// _ = "end of CoverTab[87263]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/rawbytes.go:61
var _ = _go_fuzz_dep_.CoverTab
