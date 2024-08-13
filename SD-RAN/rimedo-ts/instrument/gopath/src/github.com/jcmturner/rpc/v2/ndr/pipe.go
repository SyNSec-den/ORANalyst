//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:1
)

import (
	"fmt"
	"reflect"
)

func (dec *Decoder) fillPipe(v reflect.Value, tag reflect.StructTag) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:8
	_go_fuzz_dep_.CoverTab[87170]++
											s, err := dec.readUint32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:10
		_go_fuzz_dep_.CoverTab[87173]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:11
		// _ = "end of CoverTab[87173]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:12
		_go_fuzz_dep_.CoverTab[87174]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:12
		// _ = "end of CoverTab[87174]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:12
	// _ = "end of CoverTab[87170]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:12
	_go_fuzz_dep_.CoverTab[87171]++
											a := reflect.MakeSlice(v.Type(), 0, 0)
											c := reflect.MakeSlice(v.Type(), int(s), int(s))
											for s != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:15
		_go_fuzz_dep_.CoverTab[87175]++
												for i := 0; i < int(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:16
			_go_fuzz_dep_.CoverTab[87178]++
													err := dec.fill(c.Index(i), tag, &[]deferedPtr{})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:18
				_go_fuzz_dep_.CoverTab[87179]++
														return fmt.Errorf("could not fill element %d of pipe: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:19
				// _ = "end of CoverTab[87179]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:20
				_go_fuzz_dep_.CoverTab[87180]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:20
				// _ = "end of CoverTab[87180]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:20
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:20
			// _ = "end of CoverTab[87178]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:21
		// _ = "end of CoverTab[87175]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:21
		_go_fuzz_dep_.CoverTab[87176]++
												s, err = dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:23
			_go_fuzz_dep_.CoverTab[87181]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:24
			// _ = "end of CoverTab[87181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:25
			_go_fuzz_dep_.CoverTab[87182]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:25
			// _ = "end of CoverTab[87182]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:25
		// _ = "end of CoverTab[87176]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:25
		_go_fuzz_dep_.CoverTab[87177]++
												a = reflect.AppendSlice(a, c)
												c = reflect.MakeSlice(v.Type(), int(s), int(s))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:27
		// _ = "end of CoverTab[87177]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:28
	// _ = "end of CoverTab[87171]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:28
	_go_fuzz_dep_.CoverTab[87172]++
											v.Set(a)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:30
	// _ = "end of CoverTab[87172]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/pipe.go:31
var _ = _go_fuzz_dep_.CoverTab
