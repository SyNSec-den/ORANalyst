//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:1
)

import (
	"fmt"
	"reflect"
)

const (
	subStringArrayTag	= `ndr:"varying,X-subStringArray"`
	subStringArrayValue	= "X-subStringArray"
)

func uint16SliceToString(a []uint16) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:13
	_go_fuzz_dep_.CoverTab[87270]++
											s := make([]rune, len(a), len(a))
											for i := range s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:15
		_go_fuzz_dep_.CoverTab[87273]++
												s[i] = rune(a[i])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:16
		// _ = "end of CoverTab[87273]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:17
	// _ = "end of CoverTab[87270]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:17
	_go_fuzz_dep_.CoverTab[87271]++
											if len(s) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:18
		_go_fuzz_dep_.CoverTab[87274]++

												if s[len(s)-1] == rune(0) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:20
			_go_fuzz_dep_.CoverTab[87275]++
													s = s[:len(s)-1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:21
			// _ = "end of CoverTab[87275]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:22
			_go_fuzz_dep_.CoverTab[87276]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:22
			// _ = "end of CoverTab[87276]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:22
		// _ = "end of CoverTab[87274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:23
		_go_fuzz_dep_.CoverTab[87277]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:23
		// _ = "end of CoverTab[87277]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:23
	// _ = "end of CoverTab[87271]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:23
	_go_fuzz_dep_.CoverTab[87272]++
											return string(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:24
	// _ = "end of CoverTab[87272]"
}

func (dec *Decoder) readVaryingString(def *[]deferedPtr) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:27
	_go_fuzz_dep_.CoverTab[87278]++
											a := new([]uint16)
											v := reflect.ValueOf(a)
											var t reflect.StructTag
											err := dec.fillUniDimensionalVaryingArray(v.Elem(), t, def)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:32
		_go_fuzz_dep_.CoverTab[87280]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:33
		// _ = "end of CoverTab[87280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:34
		_go_fuzz_dep_.CoverTab[87281]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:34
		// _ = "end of CoverTab[87281]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:34
	// _ = "end of CoverTab[87278]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:34
	_go_fuzz_dep_.CoverTab[87279]++
											s := uint16SliceToString(*a)
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:36
	// _ = "end of CoverTab[87279]"
}

func (dec *Decoder) readConformantVaryingString(def *[]deferedPtr) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:39
	_go_fuzz_dep_.CoverTab[87282]++
											a := new([]uint16)
											v := reflect.ValueOf(a)
											var t reflect.StructTag
											err := dec.fillUniDimensionalConformantVaryingArray(v.Elem(), t, def)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:44
		_go_fuzz_dep_.CoverTab[87284]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:45
		// _ = "end of CoverTab[87284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:46
		_go_fuzz_dep_.CoverTab[87285]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:46
		// _ = "end of CoverTab[87285]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:46
	// _ = "end of CoverTab[87282]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:46
	_go_fuzz_dep_.CoverTab[87283]++
											s := uint16SliceToString(*a)
											return s, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:48
	// _ = "end of CoverTab[87283]"
}

func (dec *Decoder) readStringsArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:51
	_go_fuzz_dep_.CoverTab[87286]++
											d, _ := sliceDimensions(v.Type())
											ndrTag := parseTags(tag)
											var m []int

											if ndrTag.HasValue(TagConformant) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:56
		_go_fuzz_dep_.CoverTab[87289]++
												for i := 0; i < d; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:57
			_go_fuzz_dep_.CoverTab[87291]++
													m = append(m, int(dec.precedingMax()))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:58
			// _ = "end of CoverTab[87291]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:59
		// _ = "end of CoverTab[87289]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:59
		_go_fuzz_dep_.CoverTab[87290]++

												_ = dec.precedingMax()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:61
		// _ = "end of CoverTab[87290]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:63
		_go_fuzz_dep_.CoverTab[87292]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:63
		// _ = "end of CoverTab[87292]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:63
	// _ = "end of CoverTab[87286]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:63
	_go_fuzz_dep_.CoverTab[87287]++
											tag = reflect.StructTag(subStringArrayTag)
											err := dec.fillVaryingArray(v, tag, def)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:66
		_go_fuzz_dep_.CoverTab[87293]++
												return fmt.Errorf("could not read string array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:67
		// _ = "end of CoverTab[87293]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:68
		_go_fuzz_dep_.CoverTab[87294]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:68
		// _ = "end of CoverTab[87294]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:68
	// _ = "end of CoverTab[87287]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:68
	_go_fuzz_dep_.CoverTab[87288]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:69
	// _ = "end of CoverTab[87288]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/strings.go:70
var _ = _go_fuzz_dep_.CoverTab
