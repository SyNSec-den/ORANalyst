//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:1
// Package ndr provides the ability to unmarshal NDR encoded byte steams into Go data structures
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:2
)

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Struct tag values
const (
	TagConformant	= "conformant"
	TagVarying	= "varying"
	TagPointer	= "pointer"
	TagPipe		= "pipe"
)

// Decoder unmarshals NDR byte stream data into a Go struct representation
type Decoder struct {
	r		*bufio.Reader	// source of the data
	size		int		// initial size of bytes in buffer
	ch		CommonHeader	// NDR common header
	ph		PrivateHeader	// NDR private header
	conformantMax	[]uint32	// conformant max values that were moved to the beginning of the structure
	s		interface{}	// pointer to the structure being populated
	current		[]string	// keeps track of the current field being populated
}

type deferedPtr struct {
	v	reflect.Value
	tag	reflect.StructTag
}

// NewDecoder creates a new instance of a NDR Decoder.
func NewDecoder(r io.Reader) *Decoder {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:37
	_go_fuzz_dep_.CoverTab[86926]++
											dec := new(Decoder)
											dec.r = bufio.NewReader(r)
											dec.r.Peek(int(commonHeaderBytes))
											dec.size = dec.r.Buffered()
											return dec
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:42
	// _ = "end of CoverTab[86926]"
}

// Decode unmarshals the NDR encoded bytes into the pointer of a struct provided.
func (dec *Decoder) Decode(s interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:46
	_go_fuzz_dep_.CoverTab[86927]++
											dec.s = s
											err := dec.readCommonHeader()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:49
		_go_fuzz_dep_.CoverTab[86931]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:50
		// _ = "end of CoverTab[86931]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:51
		_go_fuzz_dep_.CoverTab[86932]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:51
		// _ = "end of CoverTab[86932]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:51
	// _ = "end of CoverTab[86927]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:51
	_go_fuzz_dep_.CoverTab[86928]++
											err = dec.readPrivateHeader()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:53
		_go_fuzz_dep_.CoverTab[86933]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:54
		// _ = "end of CoverTab[86933]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:55
		_go_fuzz_dep_.CoverTab[86934]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:55
		// _ = "end of CoverTab[86934]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:55
	// _ = "end of CoverTab[86928]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:55
	_go_fuzz_dep_.CoverTab[86929]++
											_, err = dec.r.Discard(4)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:57
		_go_fuzz_dep_.CoverTab[86935]++
												return Errorf("unable to process byte stream: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:58
		// _ = "end of CoverTab[86935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:59
		_go_fuzz_dep_.CoverTab[86936]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:59
		// _ = "end of CoverTab[86936]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:59
	// _ = "end of CoverTab[86929]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:59
	_go_fuzz_dep_.CoverTab[86930]++

											return dec.process(s, reflect.StructTag(""))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:61
	// _ = "end of CoverTab[86930]"
}

func (dec *Decoder) process(s interface{}, tag reflect.StructTag) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:64
	_go_fuzz_dep_.CoverTab[86937]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:67
	err := dec.scanConformantArrays(s, tag)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:68
		_go_fuzz_dep_.CoverTab[86941]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:69
		// _ = "end of CoverTab[86941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:70
		_go_fuzz_dep_.CoverTab[86942]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:70
		// _ = "end of CoverTab[86942]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:70
	// _ = "end of CoverTab[86937]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:70
	_go_fuzz_dep_.CoverTab[86938]++
	// Recursively fill the struct fields
	var localDef []deferedPtr
	err = dec.fill(s, tag, &localDef)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:74
		_go_fuzz_dep_.CoverTab[86943]++
												return Errorf("could not decode: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:75
		// _ = "end of CoverTab[86943]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:76
		_go_fuzz_dep_.CoverTab[86944]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:76
		// _ = "end of CoverTab[86944]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:76
	// _ = "end of CoverTab[86938]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:76
	_go_fuzz_dep_.CoverTab[86939]++

											for _, p := range localDef {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:78
		_go_fuzz_dep_.CoverTab[86945]++
												err = dec.process(p.v, p.tag)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:80
			_go_fuzz_dep_.CoverTab[86946]++
													return fmt.Errorf("could not decode deferred referent: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:81
			// _ = "end of CoverTab[86946]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:82
			_go_fuzz_dep_.CoverTab[86947]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:82
			// _ = "end of CoverTab[86947]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:82
		// _ = "end of CoverTab[86945]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:83
	// _ = "end of CoverTab[86939]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:83
	_go_fuzz_dep_.CoverTab[86940]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:84
	// _ = "end of CoverTab[86940]"
}

// scanConformantArrays scans the structure for embedded conformant fields and captures the maximum element counts for
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:87
// dimensions of the array that are moved to the beginning of the structure.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:89
func (dec *Decoder) scanConformantArrays(s interface{}, tag reflect.StructTag) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:89
	_go_fuzz_dep_.CoverTab[86948]++
											err := dec.conformantScan(s, tag)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:91
		_go_fuzz_dep_.CoverTab[86951]++
												return fmt.Errorf("failed to scan for embedded conformant arrays: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:92
		// _ = "end of CoverTab[86951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:93
		_go_fuzz_dep_.CoverTab[86952]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:93
		// _ = "end of CoverTab[86952]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:93
	// _ = "end of CoverTab[86948]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:93
	_go_fuzz_dep_.CoverTab[86949]++
											for i := range dec.conformantMax {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:94
		_go_fuzz_dep_.CoverTab[86953]++
												dec.conformantMax[i], err = dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:96
			_go_fuzz_dep_.CoverTab[86954]++
													return fmt.Errorf("could not read preceding conformant max count index %d: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:97
			// _ = "end of CoverTab[86954]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:98
			_go_fuzz_dep_.CoverTab[86955]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:98
			// _ = "end of CoverTab[86955]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:98
		// _ = "end of CoverTab[86953]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:99
		// _ = "end of CoverTab[86949]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:99
		_go_fuzz_dep_.CoverTab[86950]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:100
	// _ = "end of CoverTab[86950]"
}

// conformantScan inspects the structure's fields for whether they are conformant.
func (dec *Decoder) conformantScan(s interface{}, tag reflect.StructTag) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:104
	_go_fuzz_dep_.CoverTab[86956]++
												ndrTag := parseTags(tag)
												if ndrTag.HasValue(TagPointer) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:106
		_go_fuzz_dep_.CoverTab[86959]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:107
		// _ = "end of CoverTab[86959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:108
		_go_fuzz_dep_.CoverTab[86960]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:108
		// _ = "end of CoverTab[86960]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:108
	// _ = "end of CoverTab[86956]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:108
	_go_fuzz_dep_.CoverTab[86957]++
												v := getReflectValue(s)
												switch v.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:111
		_go_fuzz_dep_.CoverTab[86961]++
													for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:112
			_go_fuzz_dep_.CoverTab[86968]++
														err := dec.conformantScan(v.Field(i), v.Type().Field(i).Tag)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:114
				_go_fuzz_dep_.CoverTab[86969]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:115
				// _ = "end of CoverTab[86969]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:116
				_go_fuzz_dep_.CoverTab[86970]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:116
				// _ = "end of CoverTab[86970]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:116
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:116
			// _ = "end of CoverTab[86968]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:117
		// _ = "end of CoverTab[86961]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:118
		_go_fuzz_dep_.CoverTab[86962]++
													if !ndrTag.HasValue(TagConformant) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:119
			_go_fuzz_dep_.CoverTab[86971]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:120
			// _ = "end of CoverTab[86971]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:121
			_go_fuzz_dep_.CoverTab[86972]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:121
			// _ = "end of CoverTab[86972]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:121
		// _ = "end of CoverTab[86962]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:121
		_go_fuzz_dep_.CoverTab[86963]++
													dec.conformantMax = append(dec.conformantMax, uint32(0))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:122
		// _ = "end of CoverTab[86963]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:123
		_go_fuzz_dep_.CoverTab[86964]++
													if !ndrTag.HasValue(TagConformant) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:124
			_go_fuzz_dep_.CoverTab[86973]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:125
			// _ = "end of CoverTab[86973]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:126
			_go_fuzz_dep_.CoverTab[86974]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:126
			// _ = "end of CoverTab[86974]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:126
		// _ = "end of CoverTab[86964]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:126
		_go_fuzz_dep_.CoverTab[86965]++
													d, t := sliceDimensions(v.Type())
													for i := 0; i < d; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:128
			_go_fuzz_dep_.CoverTab[86975]++
														dec.conformantMax = append(dec.conformantMax, uint32(0))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:129
			// _ = "end of CoverTab[86975]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:130
		// _ = "end of CoverTab[86965]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:130
		_go_fuzz_dep_.CoverTab[86966]++

													if t.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:132
			_go_fuzz_dep_.CoverTab[86976]++
														dec.conformantMax = append(dec.conformantMax, uint32(0))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:133
			// _ = "end of CoverTab[86976]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
			_go_fuzz_dep_.CoverTab[86977]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
			// _ = "end of CoverTab[86977]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
		// _ = "end of CoverTab[86966]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
		_go_fuzz_dep_.CoverTab[86967]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:134
		// _ = "end of CoverTab[86967]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:135
	// _ = "end of CoverTab[86957]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:135
	_go_fuzz_dep_.CoverTab[86958]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:136
	// _ = "end of CoverTab[86958]"
}

func (dec *Decoder) isPointer(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:139
	_go_fuzz_dep_.CoverTab[86978]++

												ndrTag := parseTags(tag)
												if ndrTag.HasValue(TagPointer) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:142
		_go_fuzz_dep_.CoverTab[86980]++
													p, err := dec.readUint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:144
			_go_fuzz_dep_.CoverTab[86983]++
														return true, fmt.Errorf("could not read pointer: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:145
			// _ = "end of CoverTab[86983]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:146
			_go_fuzz_dep_.CoverTab[86984]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:146
			// _ = "end of CoverTab[86984]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:146
		// _ = "end of CoverTab[86980]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:146
		_go_fuzz_dep_.CoverTab[86981]++
													ndrTag.delete(TagPointer)
													if p != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:148
			_go_fuzz_dep_.CoverTab[86985]++

														*def = append(*def, deferedPtr{v, ndrTag.StructTag()})
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:150
			// _ = "end of CoverTab[86985]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:151
			_go_fuzz_dep_.CoverTab[86986]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:151
			// _ = "end of CoverTab[86986]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:151
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:151
		// _ = "end of CoverTab[86981]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:151
		_go_fuzz_dep_.CoverTab[86982]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:152
		// _ = "end of CoverTab[86982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:153
		_go_fuzz_dep_.CoverTab[86987]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:153
		// _ = "end of CoverTab[86987]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:153
	// _ = "end of CoverTab[86978]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:153
	_go_fuzz_dep_.CoverTab[86979]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:154
	// _ = "end of CoverTab[86979]"
}

func getReflectValue(s interface{}) (v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:157
	_go_fuzz_dep_.CoverTab[86988]++
												if r, ok := s.(reflect.Value); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:158
		_go_fuzz_dep_.CoverTab[86990]++
													v = r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:159
		// _ = "end of CoverTab[86990]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:160
		_go_fuzz_dep_.CoverTab[86991]++
													if reflect.ValueOf(s).Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:161
			_go_fuzz_dep_.CoverTab[86992]++
														v = reflect.ValueOf(s).Elem()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:162
			// _ = "end of CoverTab[86992]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:163
			_go_fuzz_dep_.CoverTab[86993]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:163
			// _ = "end of CoverTab[86993]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:163
		// _ = "end of CoverTab[86991]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:164
	// _ = "end of CoverTab[86988]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:164
	_go_fuzz_dep_.CoverTab[86989]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:165
	// _ = "end of CoverTab[86989]"
}

// fill populates fields with values from the NDR byte stream.
func (dec *Decoder) fill(s interface{}, tag reflect.StructTag, localDef *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:169
	_go_fuzz_dep_.CoverTab[86994]++
												v := getReflectValue(s)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:173
	ptr, err := dec.isPointer(v, tag, localDef)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:174
		_go_fuzz_dep_.CoverTab[86998]++
													return fmt.Errorf("could not process struct field(%s): %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:175
		// _ = "end of CoverTab[86998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:176
		_go_fuzz_dep_.CoverTab[86999]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:176
		// _ = "end of CoverTab[86999]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:176
	// _ = "end of CoverTab[86994]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:176
	_go_fuzz_dep_.CoverTab[86995]++
												if ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:177
		_go_fuzz_dep_.CoverTab[87000]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:178
		// _ = "end of CoverTab[87000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:179
		_go_fuzz_dep_.CoverTab[87001]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:179
		// _ = "end of CoverTab[87001]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:179
	// _ = "end of CoverTab[86995]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:179
	_go_fuzz_dep_.CoverTab[86996]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:182
	switch v.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:183
		_go_fuzz_dep_.CoverTab[87002]++
													dec.current = append(dec.current, v.Type().Name())
		// in case struct is a union, track this and the selected union field for efficiency
		var unionTag reflect.Value
		var unionField string	// field to fill if struct is a union

		for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:189
			_go_fuzz_dep_.CoverTab[87034]++
														fieldName := v.Type().Field(i).Name
														dec.current = append(dec.current, fieldName)

														structTag := v.Type().Field(i).Tag
														ndrTag := parseTags(structTag)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:197
			if !unionTag.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:197
				_go_fuzz_dep_.CoverTab[87037]++

															unionTag = dec.isUnion(v.Field(i), structTag)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:199
				// _ = "end of CoverTab[87037]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:200
				_go_fuzz_dep_.CoverTab[87038]++

															if unionField == "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:202
					_go_fuzz_dep_.CoverTab[87040]++
																unionField, err = unionSelectedField(v, unionTag)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:204
						_go_fuzz_dep_.CoverTab[87041]++
																	return fmt.Errorf("could not determine selected union value field for %s with discriminat"+
							" tag %s: %v", v.Type().Name(), unionTag, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:206
						// _ = "end of CoverTab[87041]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:207
						_go_fuzz_dep_.CoverTab[87042]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:207
						// _ = "end of CoverTab[87042]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:207
					}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:207
					// _ = "end of CoverTab[87040]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:208
					_go_fuzz_dep_.CoverTab[87043]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:208
					// _ = "end of CoverTab[87043]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:208
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:208
				// _ = "end of CoverTab[87038]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:208
				_go_fuzz_dep_.CoverTab[87039]++
															if ndrTag.HasValue(TagUnionField) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:209
					_go_fuzz_dep_.CoverTab[87044]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:209
					return fieldName != unionField
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:209
					// _ = "end of CoverTab[87044]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:209
				}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:209
					_go_fuzz_dep_.CoverTab[87045]++

																dec.current = dec.current[:len(dec.current)-1]
																continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:212
					// _ = "end of CoverTab[87045]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:213
					_go_fuzz_dep_.CoverTab[87046]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:213
					// _ = "end of CoverTab[87046]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:213
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:213
				// _ = "end of CoverTab[87039]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:214
			// _ = "end of CoverTab[87034]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:214
			_go_fuzz_dep_.CoverTab[87035]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:217
			if v.Field(i).Type().Implements(reflect.TypeOf(new(RawBytes)).Elem()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:217
				_go_fuzz_dep_.CoverTab[87047]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:217
				return v.Field(i).Type().Kind() == reflect.Slice
															// _ = "end of CoverTab[87047]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
				_go_fuzz_dep_.CoverTab[87048]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
				return v.Field(i).Type().Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
				// _ = "end of CoverTab[87048]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:218
				_go_fuzz_dep_.CoverTab[87049]++

															structTag, err = addSizeToTag(v, v.Field(i), structTag)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:221
					_go_fuzz_dep_.CoverTab[87052]++
																return fmt.Errorf("could not get rawbytes field(%s) size: %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:222
					// _ = "end of CoverTab[87052]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:223
					_go_fuzz_dep_.CoverTab[87053]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:223
					// _ = "end of CoverTab[87053]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:223
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:223
				// _ = "end of CoverTab[87049]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:223
				_go_fuzz_dep_.CoverTab[87050]++
															ptr, err := dec.isPointer(v.Field(i), structTag, localDef)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:225
					_go_fuzz_dep_.CoverTab[87054]++
																return fmt.Errorf("could not process struct field(%s): %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:226
					// _ = "end of CoverTab[87054]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:227
					_go_fuzz_dep_.CoverTab[87055]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:227
					// _ = "end of CoverTab[87055]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:227
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:227
				// _ = "end of CoverTab[87050]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:227
				_go_fuzz_dep_.CoverTab[87051]++
															if !ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:228
					_go_fuzz_dep_.CoverTab[87056]++
																err := dec.readRawBytes(v.Field(i), structTag)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:230
						_go_fuzz_dep_.CoverTab[87057]++
																	return fmt.Errorf("could not fill raw bytes struct field(%s): %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:231
						// _ = "end of CoverTab[87057]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:232
						_go_fuzz_dep_.CoverTab[87058]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:232
						// _ = "end of CoverTab[87058]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:232
					}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:232
					// _ = "end of CoverTab[87056]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:233
					_go_fuzz_dep_.CoverTab[87059]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:233
					// _ = "end of CoverTab[87059]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:233
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:233
				// _ = "end of CoverTab[87051]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:234
				_go_fuzz_dep_.CoverTab[87060]++
															err := dec.fill(v.Field(i), structTag, localDef)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:236
					_go_fuzz_dep_.CoverTab[87061]++
																return fmt.Errorf("could not fill struct field(%s): %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:237
					// _ = "end of CoverTab[87061]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:238
					_go_fuzz_dep_.CoverTab[87062]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:238
					// _ = "end of CoverTab[87062]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:238
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:238
				// _ = "end of CoverTab[87060]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:239
			// _ = "end of CoverTab[87035]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:239
			_go_fuzz_dep_.CoverTab[87036]++
														dec.current = dec.current[:len(dec.current)-1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:240
			// _ = "end of CoverTab[87036]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:241
		// _ = "end of CoverTab[87002]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:241
		_go_fuzz_dep_.CoverTab[87003]++
													dec.current = dec.current[:len(dec.current)-1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:242
		// _ = "end of CoverTab[87003]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:243
		_go_fuzz_dep_.CoverTab[87004]++
													i, err := dec.readBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:245
			_go_fuzz_dep_.CoverTab[87063]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:246
			// _ = "end of CoverTab[87063]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:247
			_go_fuzz_dep_.CoverTab[87064]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:247
			// _ = "end of CoverTab[87064]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:247
		// _ = "end of CoverTab[87004]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:247
		_go_fuzz_dep_.CoverTab[87005]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:248
		// _ = "end of CoverTab[87005]"
	case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:249
		_go_fuzz_dep_.CoverTab[87006]++
													i, err := dec.readUint8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:251
			_go_fuzz_dep_.CoverTab[87065]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:252
			// _ = "end of CoverTab[87065]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:253
			_go_fuzz_dep_.CoverTab[87066]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:253
			// _ = "end of CoverTab[87066]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:253
		// _ = "end of CoverTab[87006]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:253
		_go_fuzz_dep_.CoverTab[87007]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:254
		// _ = "end of CoverTab[87007]"
	case reflect.Uint16:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:255
		_go_fuzz_dep_.CoverTab[87008]++
													i, err := dec.readUint16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:257
			_go_fuzz_dep_.CoverTab[87067]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:258
			// _ = "end of CoverTab[87067]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:259
			_go_fuzz_dep_.CoverTab[87068]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:259
			// _ = "end of CoverTab[87068]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:259
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:259
		// _ = "end of CoverTab[87008]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:259
		_go_fuzz_dep_.CoverTab[87009]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:260
		// _ = "end of CoverTab[87009]"
	case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:261
		_go_fuzz_dep_.CoverTab[87010]++
													i, err := dec.readUint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:263
			_go_fuzz_dep_.CoverTab[87069]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:264
			// _ = "end of CoverTab[87069]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:265
			_go_fuzz_dep_.CoverTab[87070]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:265
			// _ = "end of CoverTab[87070]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:265
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:265
		// _ = "end of CoverTab[87010]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:265
		_go_fuzz_dep_.CoverTab[87011]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:266
		// _ = "end of CoverTab[87011]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:267
		_go_fuzz_dep_.CoverTab[87012]++
													i, err := dec.readUint64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:269
			_go_fuzz_dep_.CoverTab[87071]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:270
			// _ = "end of CoverTab[87071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:271
			_go_fuzz_dep_.CoverTab[87072]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:271
			// _ = "end of CoverTab[87072]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:271
		// _ = "end of CoverTab[87012]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:271
		_go_fuzz_dep_.CoverTab[87013]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:272
		// _ = "end of CoverTab[87013]"
	case reflect.Int8:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:273
		_go_fuzz_dep_.CoverTab[87014]++
													i, err := dec.readInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:275
			_go_fuzz_dep_.CoverTab[87073]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:276
			// _ = "end of CoverTab[87073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:277
			_go_fuzz_dep_.CoverTab[87074]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:277
			// _ = "end of CoverTab[87074]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:277
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:277
		// _ = "end of CoverTab[87014]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:277
		_go_fuzz_dep_.CoverTab[87015]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:278
		// _ = "end of CoverTab[87015]"
	case reflect.Int16:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:279
		_go_fuzz_dep_.CoverTab[87016]++
													i, err := dec.readInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:281
			_go_fuzz_dep_.CoverTab[87075]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:282
			// _ = "end of CoverTab[87075]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:283
			_go_fuzz_dep_.CoverTab[87076]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:283
			// _ = "end of CoverTab[87076]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:283
		// _ = "end of CoverTab[87016]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:283
		_go_fuzz_dep_.CoverTab[87017]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:284
		// _ = "end of CoverTab[87017]"
	case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:285
		_go_fuzz_dep_.CoverTab[87018]++
													i, err := dec.readInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:287
			_go_fuzz_dep_.CoverTab[87077]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:288
			// _ = "end of CoverTab[87077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:289
			_go_fuzz_dep_.CoverTab[87078]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:289
			// _ = "end of CoverTab[87078]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:289
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:289
		// _ = "end of CoverTab[87018]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:289
		_go_fuzz_dep_.CoverTab[87019]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:290
		// _ = "end of CoverTab[87019]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:291
		_go_fuzz_dep_.CoverTab[87020]++
													i, err := dec.readInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:293
			_go_fuzz_dep_.CoverTab[87079]++
														return fmt.Errorf("could not fill %s: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:294
			// _ = "end of CoverTab[87079]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:295
			_go_fuzz_dep_.CoverTab[87080]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:295
			// _ = "end of CoverTab[87080]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:295
		// _ = "end of CoverTab[87020]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:295
		_go_fuzz_dep_.CoverTab[87021]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:296
		// _ = "end of CoverTab[87021]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:297
		_go_fuzz_dep_.CoverTab[87022]++
													ndrTag := parseTags(tag)
													conformant := ndrTag.HasValue(TagConformant)
		// strings are always varying so this is assumed without an explicit tag
		var s string
		var err error
		if conformant {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:303
			_go_fuzz_dep_.CoverTab[87081]++
														s, err = dec.readConformantVaryingString(localDef)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:305
				_go_fuzz_dep_.CoverTab[87082]++
															return fmt.Errorf("could not fill with conformant varying string: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:306
				// _ = "end of CoverTab[87082]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:307
				_go_fuzz_dep_.CoverTab[87083]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:307
				// _ = "end of CoverTab[87083]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:307
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:307
			// _ = "end of CoverTab[87081]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:308
			_go_fuzz_dep_.CoverTab[87084]++
														s, err = dec.readVaryingString(localDef)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:310
				_go_fuzz_dep_.CoverTab[87085]++
															return fmt.Errorf("could not fill with varying string: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:311
				// _ = "end of CoverTab[87085]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:312
				_go_fuzz_dep_.CoverTab[87086]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:312
				// _ = "end of CoverTab[87086]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:312
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:312
			// _ = "end of CoverTab[87084]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:313
		// _ = "end of CoverTab[87022]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:313
		_go_fuzz_dep_.CoverTab[87023]++
													v.Set(reflect.ValueOf(s))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:314
		// _ = "end of CoverTab[87023]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:315
		_go_fuzz_dep_.CoverTab[87024]++
													i, err := dec.readFloat32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:317
			_go_fuzz_dep_.CoverTab[87087]++
														return fmt.Errorf("could not fill %v: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:318
			// _ = "end of CoverTab[87087]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:319
			_go_fuzz_dep_.CoverTab[87088]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:319
			// _ = "end of CoverTab[87088]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:319
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:319
		// _ = "end of CoverTab[87024]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:319
		_go_fuzz_dep_.CoverTab[87025]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:320
		// _ = "end of CoverTab[87025]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:321
		_go_fuzz_dep_.CoverTab[87026]++
													i, err := dec.readFloat64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:323
			_go_fuzz_dep_.CoverTab[87089]++
														return fmt.Errorf("could not fill %v: %v", v.Type().Name(), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:324
			// _ = "end of CoverTab[87089]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:325
			_go_fuzz_dep_.CoverTab[87090]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:325
			// _ = "end of CoverTab[87090]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:325
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:325
		// _ = "end of CoverTab[87026]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:325
		_go_fuzz_dep_.CoverTab[87027]++
													v.Set(reflect.ValueOf(i))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:326
		// _ = "end of CoverTab[87027]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:327
		_go_fuzz_dep_.CoverTab[87028]++
													err := dec.fillFixedArray(v, tag, localDef)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:329
			_go_fuzz_dep_.CoverTab[87091]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:330
			// _ = "end of CoverTab[87091]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:331
			_go_fuzz_dep_.CoverTab[87092]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:331
			// _ = "end of CoverTab[87092]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:331
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:331
		// _ = "end of CoverTab[87028]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:332
		_go_fuzz_dep_.CoverTab[87029]++
													if v.Type().Implements(reflect.TypeOf(new(RawBytes)).Elem()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:333
			_go_fuzz_dep_.CoverTab[87093]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:333
			return v.Type().Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:333
			// _ = "end of CoverTab[87093]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:333
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:333
			_go_fuzz_dep_.CoverTab[87094]++

														err := dec.readRawBytes(v, tag)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:336
				_go_fuzz_dep_.CoverTab[87096]++
															return fmt.Errorf("could not fill raw bytes struct field(%s): %v", strings.Join(dec.current, "/"), err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:337
				// _ = "end of CoverTab[87096]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:338
				_go_fuzz_dep_.CoverTab[87097]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:338
				// _ = "end of CoverTab[87097]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:338
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:338
			// _ = "end of CoverTab[87094]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:338
			_go_fuzz_dep_.CoverTab[87095]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:339
			// _ = "end of CoverTab[87095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:340
			_go_fuzz_dep_.CoverTab[87098]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:340
			// _ = "end of CoverTab[87098]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:340
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:340
		// _ = "end of CoverTab[87029]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:340
		_go_fuzz_dep_.CoverTab[87030]++
													ndrTag := parseTags(tag)
													conformant := ndrTag.HasValue(TagConformant)
													varying := ndrTag.HasValue(TagVarying)
													if ndrTag.HasValue(TagPipe) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:344
			_go_fuzz_dep_.CoverTab[87099]++
														err := dec.fillPipe(v, tag)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:346
				_go_fuzz_dep_.CoverTab[87101]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:347
				// _ = "end of CoverTab[87101]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:348
				_go_fuzz_dep_.CoverTab[87102]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:348
				// _ = "end of CoverTab[87102]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:348
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:348
			// _ = "end of CoverTab[87099]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:348
			_go_fuzz_dep_.CoverTab[87100]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:349
			// _ = "end of CoverTab[87100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:350
			_go_fuzz_dep_.CoverTab[87103]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:350
			// _ = "end of CoverTab[87103]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:350
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:350
		// _ = "end of CoverTab[87030]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:350
		_go_fuzz_dep_.CoverTab[87031]++
													_, t := sliceDimensions(v.Type())
													if t.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:352
			_go_fuzz_dep_.CoverTab[87104]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:352
			return !ndrTag.HasValue(subStringArrayValue)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:352
			// _ = "end of CoverTab[87104]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:352
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:352
			_go_fuzz_dep_.CoverTab[87105]++

														err := dec.readStringsArray(v, tag, localDef)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:355
				_go_fuzz_dep_.CoverTab[87107]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:356
				// _ = "end of CoverTab[87107]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:357
				_go_fuzz_dep_.CoverTab[87108]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:357
				// _ = "end of CoverTab[87108]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:357
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:357
			// _ = "end of CoverTab[87105]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:357
			_go_fuzz_dep_.CoverTab[87106]++
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:358
			// _ = "end of CoverTab[87106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:359
			_go_fuzz_dep_.CoverTab[87109]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:359
			// _ = "end of CoverTab[87109]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:359
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:359
		// _ = "end of CoverTab[87031]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:359
		_go_fuzz_dep_.CoverTab[87032]++

													if conformant && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:361
			_go_fuzz_dep_.CoverTab[87110]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:361
			return varying
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:361
			// _ = "end of CoverTab[87110]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:361
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:361
			_go_fuzz_dep_.CoverTab[87111]++
														err := dec.fillConformantVaryingArray(v, tag, localDef)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:363
				_go_fuzz_dep_.CoverTab[87112]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:364
				// _ = "end of CoverTab[87112]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:365
				_go_fuzz_dep_.CoverTab[87113]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:365
				// _ = "end of CoverTab[87113]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:365
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:365
			// _ = "end of CoverTab[87111]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
			_go_fuzz_dep_.CoverTab[87114]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
			if !conformant && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
				_go_fuzz_dep_.CoverTab[87115]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
				return varying
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
				// _ = "end of CoverTab[87115]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:366
				_go_fuzz_dep_.CoverTab[87116]++
															err := dec.fillVaryingArray(v, tag, localDef)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:368
					_go_fuzz_dep_.CoverTab[87117]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:369
					// _ = "end of CoverTab[87117]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:370
					_go_fuzz_dep_.CoverTab[87118]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:370
					// _ = "end of CoverTab[87118]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:370
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:370
				// _ = "end of CoverTab[87116]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:371
				_go_fuzz_dep_.CoverTab[87119]++

															err := dec.fillConformantArray(v, tag, localDef)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:374
					_go_fuzz_dep_.CoverTab[87120]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:375
					// _ = "end of CoverTab[87120]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:376
					_go_fuzz_dep_.CoverTab[87121]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:376
					// _ = "end of CoverTab[87121]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:376
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:376
				// _ = "end of CoverTab[87119]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:377
			// _ = "end of CoverTab[87114]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:377
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:377
		// _ = "end of CoverTab[87032]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:378
		_go_fuzz_dep_.CoverTab[87033]++
													return fmt.Errorf("unsupported type")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:379
		// _ = "end of CoverTab[87033]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:380
	// _ = "end of CoverTab[86996]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:380
	_go_fuzz_dep_.CoverTab[86997]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:381
	// _ = "end of CoverTab[86997]"
}

// readBytes returns a number of bytes from the NDR byte stream.
func (dec *Decoder) readBytes(n int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:385
	_go_fuzz_dep_.CoverTab[87122]++

												b := make([]byte, n, n)
												m, err := dec.r.Read(b)
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:389
		_go_fuzz_dep_.CoverTab[87124]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:389
		return m != n
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:389
		// _ = "end of CoverTab[87124]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:389
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:389
		_go_fuzz_dep_.CoverTab[87125]++
													return b, fmt.Errorf("error reading bytes from stream: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:390
		// _ = "end of CoverTab[87125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:391
		_go_fuzz_dep_.CoverTab[87126]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:391
		// _ = "end of CoverTab[87126]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:391
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:391
	// _ = "end of CoverTab[87122]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:391
	_go_fuzz_dep_.CoverTab[87123]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:392
	// _ = "end of CoverTab[87123]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:393
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/decoder.go:393
var _ = _go_fuzz_dep_.CoverTab
