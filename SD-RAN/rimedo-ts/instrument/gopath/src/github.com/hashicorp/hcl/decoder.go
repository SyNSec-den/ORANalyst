//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
package hcl

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:1
)

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/parser"
	"github.com/hashicorp/hcl/hcl/token"
)

// This is the tag to use with structures to have settings for HCL
const tagName = "hcl"

var (
	// nodeType holds a reference to the type of ast.Node
	nodeType reflect.Type = findNodeType()
)

// Unmarshal accepts a byte slice as input and writes the
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:24
// data to the value pointed to by v.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:26
func Unmarshal(bs []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:26
	_go_fuzz_dep_.CoverTab[121884]++
											root, err := parse(bs)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:28
		_go_fuzz_dep_.CoverTab[121886]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:29
		// _ = "end of CoverTab[121886]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:30
		_go_fuzz_dep_.CoverTab[121887]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:30
		// _ = "end of CoverTab[121887]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:30
	// _ = "end of CoverTab[121884]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:30
	_go_fuzz_dep_.CoverTab[121885]++

											return DecodeObject(v, root)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:32
	// _ = "end of CoverTab[121885]"
}

// Decode reads the given input and decodes it into the structure
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:35
// given by `out`.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:37
func Decode(out interface{}, in string) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:37
	_go_fuzz_dep_.CoverTab[121888]++
											obj, err := Parse(in)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:39
		_go_fuzz_dep_.CoverTab[121890]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:40
		// _ = "end of CoverTab[121890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:41
		_go_fuzz_dep_.CoverTab[121891]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:41
		// _ = "end of CoverTab[121891]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:41
	// _ = "end of CoverTab[121888]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:41
	_go_fuzz_dep_.CoverTab[121889]++

											return DecodeObject(out, obj)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:43
	// _ = "end of CoverTab[121889]"
}

// DecodeObject is a lower-level version of Decode. It decodes a
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:46
// raw Object into the given output.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:48
func DecodeObject(out interface{}, n ast.Node) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:48
	_go_fuzz_dep_.CoverTab[121892]++
											val := reflect.ValueOf(out)
											if val.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:50
		_go_fuzz_dep_.CoverTab[121895]++
												return errors.New("result must be a pointer")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:51
		// _ = "end of CoverTab[121895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:52
		_go_fuzz_dep_.CoverTab[121896]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:52
		// _ = "end of CoverTab[121896]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:52
	// _ = "end of CoverTab[121892]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:52
	_go_fuzz_dep_.CoverTab[121893]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:55
	if f, ok := n.(*ast.File); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:55
		_go_fuzz_dep_.CoverTab[121897]++
												n = f.Node
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:56
		// _ = "end of CoverTab[121897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:57
		_go_fuzz_dep_.CoverTab[121898]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:57
		// _ = "end of CoverTab[121898]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:57
	// _ = "end of CoverTab[121893]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:57
	_go_fuzz_dep_.CoverTab[121894]++

											var d decoder
											return d.decode("root", n, val.Elem())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:60
	// _ = "end of CoverTab[121894]"
}

type decoder struct {
	stack []reflect.Kind
}

func (d *decoder) decode(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:67
	_go_fuzz_dep_.CoverTab[121899]++
											k := result

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:72
	if result.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:72
		_go_fuzz_dep_.CoverTab[121902]++
												elem := result.Elem()
												if elem.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:74
			_go_fuzz_dep_.CoverTab[121903]++
													k = elem
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:75
			// _ = "end of CoverTab[121903]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:76
			_go_fuzz_dep_.CoverTab[121904]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:76
			// _ = "end of CoverTab[121904]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:76
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:76
		// _ = "end of CoverTab[121902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:77
		_go_fuzz_dep_.CoverTab[121905]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:77
		// _ = "end of CoverTab[121905]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:77
	// _ = "end of CoverTab[121899]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:77
	_go_fuzz_dep_.CoverTab[121900]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:80
	if k.Kind() != reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:80
		_go_fuzz_dep_.CoverTab[121906]++
												d.stack = append(d.stack, k.Kind())

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:84
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:84
			_go_fuzz_dep_.CoverTab[121907]++
													d.stack = d.stack[:len(d.stack)-1]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:85
			// _ = "end of CoverTab[121907]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:86
		// _ = "end of CoverTab[121906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:87
		_go_fuzz_dep_.CoverTab[121908]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:87
		// _ = "end of CoverTab[121908]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:87
	// _ = "end of CoverTab[121900]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:87
	_go_fuzz_dep_.CoverTab[121901]++

											switch k.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:90
		_go_fuzz_dep_.CoverTab[121909]++
												return d.decodeBool(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:91
		// _ = "end of CoverTab[121909]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:92
		_go_fuzz_dep_.CoverTab[121910]++
												return d.decodeFloat(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:93
		// _ = "end of CoverTab[121910]"
	case reflect.Int, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:94
		_go_fuzz_dep_.CoverTab[121911]++
												return d.decodeInt(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:95
		// _ = "end of CoverTab[121911]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:96
		_go_fuzz_dep_.CoverTab[121912]++

												return d.decodeInterface(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:98
		// _ = "end of CoverTab[121912]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:99
		_go_fuzz_dep_.CoverTab[121913]++
												return d.decodeMap(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:100
		// _ = "end of CoverTab[121913]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:101
		_go_fuzz_dep_.CoverTab[121914]++
												return d.decodePtr(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:102
		// _ = "end of CoverTab[121914]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:103
		_go_fuzz_dep_.CoverTab[121915]++
												return d.decodeSlice(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:104
		// _ = "end of CoverTab[121915]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:105
		_go_fuzz_dep_.CoverTab[121916]++
												return d.decodeString(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:106
		// _ = "end of CoverTab[121916]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:107
		_go_fuzz_dep_.CoverTab[121917]++
												return d.decodeStruct(name, node, result)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:108
		// _ = "end of CoverTab[121917]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:109
		_go_fuzz_dep_.CoverTab[121918]++
												return &parser.PosError{
			Pos:	node.Pos(),
			Err:	fmt.Errorf("%s: unknown kind to decode into: %s", name, k.Kind()),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:113
		// _ = "end of CoverTab[121918]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:114
	// _ = "end of CoverTab[121901]"
}

func (d *decoder) decodeBool(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:117
	_go_fuzz_dep_.CoverTab[121919]++
											switch n := node.(type) {
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:119
		_go_fuzz_dep_.CoverTab[121921]++
												if n.Token.Type == token.BOOL {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:120
			_go_fuzz_dep_.CoverTab[121922]++
													v, err := strconv.ParseBool(n.Token.Text)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:122
				_go_fuzz_dep_.CoverTab[121924]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:123
				// _ = "end of CoverTab[121924]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:124
				_go_fuzz_dep_.CoverTab[121925]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:124
				// _ = "end of CoverTab[121925]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:124
			// _ = "end of CoverTab[121922]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:124
			_go_fuzz_dep_.CoverTab[121923]++

													result.Set(reflect.ValueOf(v))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:127
			// _ = "end of CoverTab[121923]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:128
			_go_fuzz_dep_.CoverTab[121926]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:128
			// _ = "end of CoverTab[121926]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:128
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:128
		// _ = "end of CoverTab[121921]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:129
	// _ = "end of CoverTab[121919]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:129
	_go_fuzz_dep_.CoverTab[121920]++

											return &parser.PosError{
		Pos:	node.Pos(),
		Err:	fmt.Errorf("%s: unknown type %T", name, node),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:134
	// _ = "end of CoverTab[121920]"
}

func (d *decoder) decodeFloat(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:137
	_go_fuzz_dep_.CoverTab[121927]++
											switch n := node.(type) {
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:139
		_go_fuzz_dep_.CoverTab[121929]++
												if n.Token.Type == token.FLOAT || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:140
			_go_fuzz_dep_.CoverTab[121930]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:140
			return n.Token.Type == token.NUMBER
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:140
			// _ = "end of CoverTab[121930]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:140
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:140
			_go_fuzz_dep_.CoverTab[121931]++
													v, err := strconv.ParseFloat(n.Token.Text, 64)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:142
				_go_fuzz_dep_.CoverTab[121933]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:143
				// _ = "end of CoverTab[121933]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:144
				_go_fuzz_dep_.CoverTab[121934]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:144
				// _ = "end of CoverTab[121934]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:144
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:144
			// _ = "end of CoverTab[121931]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:144
			_go_fuzz_dep_.CoverTab[121932]++

													result.Set(reflect.ValueOf(v).Convert(result.Type()))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:147
			// _ = "end of CoverTab[121932]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:148
			_go_fuzz_dep_.CoverTab[121935]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:148
			// _ = "end of CoverTab[121935]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:148
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:148
		// _ = "end of CoverTab[121929]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:149
	// _ = "end of CoverTab[121927]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:149
	_go_fuzz_dep_.CoverTab[121928]++

											return &parser.PosError{
		Pos:	node.Pos(),
		Err:	fmt.Errorf("%s: unknown type %T", name, node),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:154
	// _ = "end of CoverTab[121928]"
}

func (d *decoder) decodeInt(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:157
	_go_fuzz_dep_.CoverTab[121936]++
											switch n := node.(type) {
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:159
		_go_fuzz_dep_.CoverTab[121938]++
												switch n.Token.Type {
		case token.NUMBER:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:161
			_go_fuzz_dep_.CoverTab[121939]++
													v, err := strconv.ParseInt(n.Token.Text, 0, 0)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:163
				_go_fuzz_dep_.CoverTab[121946]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:164
				// _ = "end of CoverTab[121946]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:165
				_go_fuzz_dep_.CoverTab[121947]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:165
				// _ = "end of CoverTab[121947]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:165
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:165
			// _ = "end of CoverTab[121939]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:165
			_go_fuzz_dep_.CoverTab[121940]++

													if result.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:167
				_go_fuzz_dep_.CoverTab[121948]++
														result.Set(reflect.ValueOf(int(v)))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:168
				// _ = "end of CoverTab[121948]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:169
				_go_fuzz_dep_.CoverTab[121949]++
														result.SetInt(v)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:170
				// _ = "end of CoverTab[121949]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:171
			// _ = "end of CoverTab[121940]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:171
			_go_fuzz_dep_.CoverTab[121941]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:172
			// _ = "end of CoverTab[121941]"
		case token.STRING:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:173
			_go_fuzz_dep_.CoverTab[121942]++
													v, err := strconv.ParseInt(n.Token.Value().(string), 0, 0)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:175
				_go_fuzz_dep_.CoverTab[121950]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:176
				// _ = "end of CoverTab[121950]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:177
				_go_fuzz_dep_.CoverTab[121951]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:177
				// _ = "end of CoverTab[121951]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:177
			// _ = "end of CoverTab[121942]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:177
			_go_fuzz_dep_.CoverTab[121943]++

													if result.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:179
				_go_fuzz_dep_.CoverTab[121952]++
														result.Set(reflect.ValueOf(int(v)))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:180
				// _ = "end of CoverTab[121952]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:181
				_go_fuzz_dep_.CoverTab[121953]++
														result.SetInt(v)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:182
				// _ = "end of CoverTab[121953]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:183
			// _ = "end of CoverTab[121943]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:183
			_go_fuzz_dep_.CoverTab[121944]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:184
			// _ = "end of CoverTab[121944]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:184
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:184
			_go_fuzz_dep_.CoverTab[121945]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:184
			// _ = "end of CoverTab[121945]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:185
		// _ = "end of CoverTab[121938]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:186
	// _ = "end of CoverTab[121936]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:186
	_go_fuzz_dep_.CoverTab[121937]++

											return &parser.PosError{
		Pos:	node.Pos(),
		Err:	fmt.Errorf("%s: unknown type %T", name, node),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:191
	// _ = "end of CoverTab[121937]"
}

func (d *decoder) decodeInterface(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:194
	_go_fuzz_dep_.CoverTab[121954]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
	if result.Type() == nodeType && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
		_go_fuzz_dep_.CoverTab[121959]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
		return result.CanSet()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
		// _ = "end of CoverTab[121959]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:198
		_go_fuzz_dep_.CoverTab[121960]++
												result.Set(reflect.ValueOf(node))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:200
		// _ = "end of CoverTab[121960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:201
		_go_fuzz_dep_.CoverTab[121961]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:201
		// _ = "end of CoverTab[121961]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:201
	// _ = "end of CoverTab[121954]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:201
	_go_fuzz_dep_.CoverTab[121955]++

											var set reflect.Value
											redecode := true

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:208
	testNode := node
	if ot, ok := node.(*ast.ObjectType); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:209
		_go_fuzz_dep_.CoverTab[121962]++
												testNode = ot.List
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:210
		// _ = "end of CoverTab[121962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:211
		_go_fuzz_dep_.CoverTab[121963]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:211
		// _ = "end of CoverTab[121963]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:211
	// _ = "end of CoverTab[121955]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:211
	_go_fuzz_dep_.CoverTab[121956]++

											switch n := testNode.(type) {
	case *ast.ObjectList:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:214
		_go_fuzz_dep_.CoverTab[121964]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
		if len(d.stack) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
			_go_fuzz_dep_.CoverTab[121969]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
			return d.stack[len(d.stack)-1] == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
			// _ = "end of CoverTab[121969]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:218
			_go_fuzz_dep_.CoverTab[121970]++
													var temp map[string]interface{}
													tempVal := reflect.ValueOf(temp)
													result := reflect.MakeMap(
				reflect.MapOf(
					reflect.TypeOf(""),
					tempVal.Type().Elem()))

													set = result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:226
			// _ = "end of CoverTab[121970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:227
			_go_fuzz_dep_.CoverTab[121971]++
													var temp []map[string]interface{}
													tempVal := reflect.ValueOf(temp)
													result := reflect.MakeSlice(
				reflect.SliceOf(tempVal.Type().Elem()), 0, len(n.Items))
													set = result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:232
			// _ = "end of CoverTab[121971]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:233
		// _ = "end of CoverTab[121964]"
	case *ast.ObjectType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:234
		_go_fuzz_dep_.CoverTab[121965]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
		if len(d.stack) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
			_go_fuzz_dep_.CoverTab[121972]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
			return d.stack[len(d.stack)-1] == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
			// _ = "end of CoverTab[121972]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:238
			_go_fuzz_dep_.CoverTab[121973]++
													var temp map[string]interface{}
													tempVal := reflect.ValueOf(temp)
													result := reflect.MakeMap(
				reflect.MapOf(
					reflect.TypeOf(""),
					tempVal.Type().Elem()))

													set = result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:246
			// _ = "end of CoverTab[121973]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:247
			_go_fuzz_dep_.CoverTab[121974]++
													var temp []map[string]interface{}
													tempVal := reflect.ValueOf(temp)
													result := reflect.MakeSlice(
				reflect.SliceOf(tempVal.Type().Elem()), 0, 1)
													set = result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:252
			// _ = "end of CoverTab[121974]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:253
		// _ = "end of CoverTab[121965]"
	case *ast.ListType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:254
		_go_fuzz_dep_.CoverTab[121966]++
												var temp []interface{}
												tempVal := reflect.ValueOf(temp)
												result := reflect.MakeSlice(
			reflect.SliceOf(tempVal.Type().Elem()), 0, 0)
												set = result
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:259
		// _ = "end of CoverTab[121966]"
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:260
		_go_fuzz_dep_.CoverTab[121967]++
												switch n.Token.Type {
		case token.BOOL:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:262
			_go_fuzz_dep_.CoverTab[121975]++
													var result bool
													set = reflect.Indirect(reflect.New(reflect.TypeOf(result)))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:264
			// _ = "end of CoverTab[121975]"
		case token.FLOAT:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:265
			_go_fuzz_dep_.CoverTab[121976]++
													var result float64
													set = reflect.Indirect(reflect.New(reflect.TypeOf(result)))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:267
			// _ = "end of CoverTab[121976]"
		case token.NUMBER:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:268
			_go_fuzz_dep_.CoverTab[121977]++
													var result int
													set = reflect.Indirect(reflect.New(reflect.TypeOf(result)))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:270
			// _ = "end of CoverTab[121977]"
		case token.STRING, token.HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:271
			_go_fuzz_dep_.CoverTab[121978]++
													set = reflect.Indirect(reflect.New(reflect.TypeOf("")))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:272
			// _ = "end of CoverTab[121978]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:273
			_go_fuzz_dep_.CoverTab[121979]++
													return &parser.PosError{
				Pos:	node.Pos(),
				Err:	fmt.Errorf("%s: cannot decode into interface: %T", name, node),
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:277
			// _ = "end of CoverTab[121979]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:278
		// _ = "end of CoverTab[121967]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:279
		_go_fuzz_dep_.CoverTab[121968]++
												return fmt.Errorf(
			"%s: cannot decode into interface: %T",
			name, node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:282
		// _ = "end of CoverTab[121968]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:283
	// _ = "end of CoverTab[121956]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:283
	_go_fuzz_dep_.CoverTab[121957]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:287
	result.Set(set)

	if redecode {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:289
		_go_fuzz_dep_.CoverTab[121980]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:292
		if err := d.decode(name, node, result); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:292
			_go_fuzz_dep_.CoverTab[121981]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:293
			// _ = "end of CoverTab[121981]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:294
			_go_fuzz_dep_.CoverTab[121982]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:294
			// _ = "end of CoverTab[121982]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:294
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:294
		// _ = "end of CoverTab[121980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:295
		_go_fuzz_dep_.CoverTab[121983]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:295
		// _ = "end of CoverTab[121983]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:295
	// _ = "end of CoverTab[121957]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:295
	_go_fuzz_dep_.CoverTab[121958]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:297
	// _ = "end of CoverTab[121958]"
}

func (d *decoder) decodeMap(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:300
	_go_fuzz_dep_.CoverTab[121984]++
											if item, ok := node.(*ast.ObjectItem); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:301
		_go_fuzz_dep_.CoverTab[121992]++
												node = &ast.ObjectList{Items: []*ast.ObjectItem{item}}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:302
		// _ = "end of CoverTab[121992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:303
		_go_fuzz_dep_.CoverTab[121993]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:303
		// _ = "end of CoverTab[121993]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:303
	// _ = "end of CoverTab[121984]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:303
	_go_fuzz_dep_.CoverTab[121985]++

											if ot, ok := node.(*ast.ObjectType); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:305
		_go_fuzz_dep_.CoverTab[121994]++
												node = ot.List
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:306
		// _ = "end of CoverTab[121994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:307
		_go_fuzz_dep_.CoverTab[121995]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:307
		// _ = "end of CoverTab[121995]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:307
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:307
	// _ = "end of CoverTab[121985]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:307
	_go_fuzz_dep_.CoverTab[121986]++

											n, ok := node.(*ast.ObjectList)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:310
		_go_fuzz_dep_.CoverTab[121996]++
												return &parser.PosError{
			Pos:	node.Pos(),
			Err:	fmt.Errorf("%s: not an object type for map (%T)", name, node),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:314
		// _ = "end of CoverTab[121996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:315
		_go_fuzz_dep_.CoverTab[121997]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:315
		// _ = "end of CoverTab[121997]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:315
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:315
	// _ = "end of CoverTab[121986]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:315
	_go_fuzz_dep_.CoverTab[121987]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:319
	set := result
	if result.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:320
		_go_fuzz_dep_.CoverTab[121998]++
												result = result.Elem()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:321
		// _ = "end of CoverTab[121998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:322
		_go_fuzz_dep_.CoverTab[121999]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:322
		// _ = "end of CoverTab[121999]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:322
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:322
	// _ = "end of CoverTab[121987]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:322
	_go_fuzz_dep_.CoverTab[121988]++

											resultType := result.Type()
											resultElemType := resultType.Elem()
											resultKeyType := resultType.Key()
											if resultKeyType.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:327
		_go_fuzz_dep_.CoverTab[122000]++
												return &parser.PosError{
			Pos:	node.Pos(),
			Err:	fmt.Errorf("%s: map must have string keys", name),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:331
		// _ = "end of CoverTab[122000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:332
		_go_fuzz_dep_.CoverTab[122001]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:332
		// _ = "end of CoverTab[122001]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:332
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:332
	// _ = "end of CoverTab[121988]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:332
	_go_fuzz_dep_.CoverTab[121989]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:335
	resultMap := result
	if result.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:336
		_go_fuzz_dep_.CoverTab[122002]++
												resultMap = reflect.MakeMap(
			reflect.MapOf(resultKeyType, resultElemType))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:338
		// _ = "end of CoverTab[122002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:339
		_go_fuzz_dep_.CoverTab[122003]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:339
		// _ = "end of CoverTab[122003]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:339
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:339
	// _ = "end of CoverTab[121989]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:339
	_go_fuzz_dep_.CoverTab[121990]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:342
	done := make(map[string]struct{})
	for _, item := range n.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:343
		_go_fuzz_dep_.CoverTab[122004]++
												if item.Val == nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:344
			_go_fuzz_dep_.CoverTab[122011]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:345
			// _ = "end of CoverTab[122011]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:346
			_go_fuzz_dep_.CoverTab[122012]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:346
			// _ = "end of CoverTab[122012]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:346
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:346
		// _ = "end of CoverTab[122004]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:346
		_go_fuzz_dep_.CoverTab[122005]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:349
		if len(item.Keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:349
			_go_fuzz_dep_.CoverTab[122013]++
													return &parser.PosError{
				Pos:	node.Pos(),
				Err:	fmt.Errorf("%s: map must have string keys", name),
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:353
			// _ = "end of CoverTab[122013]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:354
			_go_fuzz_dep_.CoverTab[122014]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:354
			// _ = "end of CoverTab[122014]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:354
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:354
		// _ = "end of CoverTab[122005]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:354
		_go_fuzz_dep_.CoverTab[122006]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:357
		keyStr := item.Keys[0].Token.Value().(string)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:360
		if _, ok := done[keyStr]; ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:360
			_go_fuzz_dep_.CoverTab[122015]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:361
			// _ = "end of CoverTab[122015]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:362
			_go_fuzz_dep_.CoverTab[122016]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:362
			// _ = "end of CoverTab[122016]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:362
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:362
		// _ = "end of CoverTab[122006]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:362
		_go_fuzz_dep_.CoverTab[122007]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:366
		itemVal := item.Val
		if len(item.Keys) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:367
			_go_fuzz_dep_.CoverTab[122017]++
													itemVal = n.Filter(keyStr)
													done[keyStr] = struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:369
			// _ = "end of CoverTab[122017]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:370
			_go_fuzz_dep_.CoverTab[122018]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:370
			// _ = "end of CoverTab[122018]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:370
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:370
		// _ = "end of CoverTab[122007]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:370
		_go_fuzz_dep_.CoverTab[122008]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:373
		fieldName := fmt.Sprintf("%s.%s", name, keyStr)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:376
		key := reflect.ValueOf(keyStr)
												val := reflect.Indirect(reflect.New(resultElemType))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:380
		oldVal := resultMap.MapIndex(key)
		if oldVal.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:381
			_go_fuzz_dep_.CoverTab[122019]++
													val.Set(oldVal)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:382
			// _ = "end of CoverTab[122019]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:383
			_go_fuzz_dep_.CoverTab[122020]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:383
			// _ = "end of CoverTab[122020]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:383
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:383
		// _ = "end of CoverTab[122008]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:383
		_go_fuzz_dep_.CoverTab[122009]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:386
		if err := d.decode(fieldName, itemVal, val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:386
			_go_fuzz_dep_.CoverTab[122021]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:387
			// _ = "end of CoverTab[122021]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:388
			_go_fuzz_dep_.CoverTab[122022]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:388
			// _ = "end of CoverTab[122022]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:388
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:388
		// _ = "end of CoverTab[122009]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:388
		_go_fuzz_dep_.CoverTab[122010]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:391
		resultMap.SetMapIndex(key, val)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:391
		// _ = "end of CoverTab[122010]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:392
	// _ = "end of CoverTab[121990]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:392
	_go_fuzz_dep_.CoverTab[121991]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:395
	set.Set(resultMap)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:396
	// _ = "end of CoverTab[121991]"
}

func (d *decoder) decodePtr(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:399
	_go_fuzz_dep_.CoverTab[122023]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:402
	resultType := result.Type()
	resultElemType := resultType.Elem()
	val := reflect.New(resultElemType)
	if err := d.decode(name, node, reflect.Indirect(val)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:405
		_go_fuzz_dep_.CoverTab[122025]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:406
		// _ = "end of CoverTab[122025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:407
		_go_fuzz_dep_.CoverTab[122026]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:407
		// _ = "end of CoverTab[122026]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:407
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:407
	// _ = "end of CoverTab[122023]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:407
	_go_fuzz_dep_.CoverTab[122024]++

											result.Set(val)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:410
	// _ = "end of CoverTab[122024]"
}

func (d *decoder) decodeSlice(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:413
	_go_fuzz_dep_.CoverTab[122027]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:416
	set := result
	if result.Kind() == reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:417
		_go_fuzz_dep_.CoverTab[122032]++
												result = result.Elem()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:418
		// _ = "end of CoverTab[122032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:419
		_go_fuzz_dep_.CoverTab[122033]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:419
		// _ = "end of CoverTab[122033]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:419
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:419
	// _ = "end of CoverTab[122027]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:419
	_go_fuzz_dep_.CoverTab[122028]++

											resultType := result.Type()
											resultElemType := resultType.Elem()
											if result.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:423
		_go_fuzz_dep_.CoverTab[122034]++
												resultSliceType := reflect.SliceOf(resultElemType)
												result = reflect.MakeSlice(
			resultSliceType, 0, 0)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:426
		// _ = "end of CoverTab[122034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:427
		_go_fuzz_dep_.CoverTab[122035]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:427
		// _ = "end of CoverTab[122035]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:427
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:427
	// _ = "end of CoverTab[122028]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:427
	_go_fuzz_dep_.CoverTab[122029]++

	// Figure out the items we'll be copying into the slice
	var items []ast.Node
	switch n := node.(type) {
	case *ast.ObjectList:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:432
		_go_fuzz_dep_.CoverTab[122036]++
												items = make([]ast.Node, len(n.Items))
												for i, item := range n.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:434
			_go_fuzz_dep_.CoverTab[122040]++
													items[i] = item
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:435
			// _ = "end of CoverTab[122040]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:436
		// _ = "end of CoverTab[122036]"
	case *ast.ObjectType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:437
		_go_fuzz_dep_.CoverTab[122037]++
												items = []ast.Node{n}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:438
		// _ = "end of CoverTab[122037]"
	case *ast.ListType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:439
		_go_fuzz_dep_.CoverTab[122038]++
												items = n.List
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:440
		// _ = "end of CoverTab[122038]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:441
		_go_fuzz_dep_.CoverTab[122039]++
												return &parser.PosError{
			Pos:	node.Pos(),
			Err:	fmt.Errorf("unknown slice type: %T", node),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:445
		// _ = "end of CoverTab[122039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:446
	// _ = "end of CoverTab[122029]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:446
	_go_fuzz_dep_.CoverTab[122030]++

											for i, item := range items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:448
		_go_fuzz_dep_.CoverTab[122041]++
												fieldName := fmt.Sprintf("%s[%d]", name, i)

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:452
		val := reflect.Indirect(reflect.New(resultElemType))

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:457
		item := expandObject(item, val)

		if err := d.decode(fieldName, item, val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:459
			_go_fuzz_dep_.CoverTab[122043]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:460
			// _ = "end of CoverTab[122043]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:461
			_go_fuzz_dep_.CoverTab[122044]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:461
			// _ = "end of CoverTab[122044]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:461
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:461
		// _ = "end of CoverTab[122041]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:461
		_go_fuzz_dep_.CoverTab[122042]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:464
		result = reflect.Append(result, val)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:464
		// _ = "end of CoverTab[122042]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:465
	// _ = "end of CoverTab[122030]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:465
	_go_fuzz_dep_.CoverTab[122031]++

											set.Set(result)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:468
	// _ = "end of CoverTab[122031]"
}

// expandObject detects if an ambiguous JSON object was flattened to a List which
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:471
// should be decoded into a struct, and expands the ast to properly deocode.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:473
func expandObject(node ast.Node, result reflect.Value) ast.Node {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:473
	_go_fuzz_dep_.CoverTab[122045]++
											item, ok := node.(*ast.ObjectItem)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:475
		_go_fuzz_dep_.CoverTab[122049]++
												return node
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:476
		// _ = "end of CoverTab[122049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:477
		_go_fuzz_dep_.CoverTab[122050]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:477
		// _ = "end of CoverTab[122050]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:477
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:477
	// _ = "end of CoverTab[122045]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:477
	_go_fuzz_dep_.CoverTab[122046]++

											elemType := result.Type()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:482
	switch elemType.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:483
		_go_fuzz_dep_.CoverTab[122051]++
												switch elemType.Elem().Kind() {
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:485
			_go_fuzz_dep_.CoverTab[122054]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:485
			// _ = "end of CoverTab[122054]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:487
			_go_fuzz_dep_.CoverTab[122055]++
													return node
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:488
			// _ = "end of CoverTab[122055]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:489
		// _ = "end of CoverTab[122051]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:490
		_go_fuzz_dep_.CoverTab[122052]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:490
		// _ = "end of CoverTab[122052]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:492
		_go_fuzz_dep_.CoverTab[122053]++
												return node
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:493
		// _ = "end of CoverTab[122053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:494
	// _ = "end of CoverTab[122046]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:494
	_go_fuzz_dep_.CoverTab[122047]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:498
	if len(item.Keys) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:498
		_go_fuzz_dep_.CoverTab[122056]++
												return node
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:499
		// _ = "end of CoverTab[122056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:500
		_go_fuzz_dep_.CoverTab[122057]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:500
		// _ = "end of CoverTab[122057]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:500
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:500
	// _ = "end of CoverTab[122047]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:500
	_go_fuzz_dep_.CoverTab[122048]++

											keyToken := item.Keys[0].Token
											item.Keys = item.Keys[1:]

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:506
	newNode := &ast.ObjectItem{
		Keys: []*ast.ObjectKey{
			&ast.ObjectKey{
				Token: keyToken,
			},
		},
		Val: &ast.ObjectType{
			List: &ast.ObjectList{
				Items: []*ast.ObjectItem{item},
			},
		},
	}

											return newNode
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:519
	// _ = "end of CoverTab[122048]"
}

func (d *decoder) decodeString(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:522
	_go_fuzz_dep_.CoverTab[122058]++
											switch n := node.(type) {
	case *ast.LiteralType:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:524
		_go_fuzz_dep_.CoverTab[122060]++
												switch n.Token.Type {
		case token.NUMBER:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:526
			_go_fuzz_dep_.CoverTab[122061]++
													result.Set(reflect.ValueOf(n.Token.Text).Convert(result.Type()))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:528
			// _ = "end of CoverTab[122061]"
		case token.STRING, token.HEREDOC:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:529
			_go_fuzz_dep_.CoverTab[122062]++
													result.Set(reflect.ValueOf(n.Token.Value()).Convert(result.Type()))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:531
			// _ = "end of CoverTab[122062]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:531
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:531
			_go_fuzz_dep_.CoverTab[122063]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:531
			// _ = "end of CoverTab[122063]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:532
		// _ = "end of CoverTab[122060]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:533
	// _ = "end of CoverTab[122058]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:533
	_go_fuzz_dep_.CoverTab[122059]++

											return &parser.PosError{
		Pos:	node.Pos(),
		Err:	fmt.Errorf("%s: unknown type for string %T", name, node),
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:538
	// _ = "end of CoverTab[122059]"
}

func (d *decoder) decodeStruct(name string, node ast.Node, result reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:541
	_go_fuzz_dep_.CoverTab[122064]++
											var item *ast.ObjectItem
											if it, ok := node.(*ast.ObjectItem); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:543
		_go_fuzz_dep_.CoverTab[122072]++
												item = it
												node = it.Val
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:545
		// _ = "end of CoverTab[122072]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:546
		_go_fuzz_dep_.CoverTab[122073]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:546
		// _ = "end of CoverTab[122073]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:546
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:546
	// _ = "end of CoverTab[122064]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:546
	_go_fuzz_dep_.CoverTab[122065]++

											if ot, ok := node.(*ast.ObjectType); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:548
		_go_fuzz_dep_.CoverTab[122074]++
												node = ot.List
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:549
		// _ = "end of CoverTab[122074]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:550
		_go_fuzz_dep_.CoverTab[122075]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:550
		// _ = "end of CoverTab[122075]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:550
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:550
	// _ = "end of CoverTab[122065]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:550
	_go_fuzz_dep_.CoverTab[122066]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
	if _, ok := node.(*ast.LiteralType); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
		_go_fuzz_dep_.CoverTab[122076]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
		return item != nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
		// _ = "end of CoverTab[122076]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:556
		_go_fuzz_dep_.CoverTab[122077]++
												node = &ast.ObjectList{Items: []*ast.ObjectItem{item}}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:557
		// _ = "end of CoverTab[122077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:558
		_go_fuzz_dep_.CoverTab[122078]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:558
		// _ = "end of CoverTab[122078]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:558
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:558
	// _ = "end of CoverTab[122066]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:558
	_go_fuzz_dep_.CoverTab[122067]++

											list, ok := node.(*ast.ObjectList)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:561
		_go_fuzz_dep_.CoverTab[122079]++
												return &parser.PosError{
			Pos:	node.Pos(),
			Err:	fmt.Errorf("%s: not an object type for struct (%T)", name, node),
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:565
		// _ = "end of CoverTab[122079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:566
		_go_fuzz_dep_.CoverTab[122080]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:566
		// _ = "end of CoverTab[122080]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:566
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:566
	// _ = "end of CoverTab[122067]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:566
	_go_fuzz_dep_.CoverTab[122068]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:571
	structs := make([]reflect.Value, 1, 5)
	structs[0] = result

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type field struct {
		field	reflect.StructField
		val	reflect.Value
	}
	fields := []field{}
	for len(structs) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:581
		_go_fuzz_dep_.CoverTab[122081]++
												structVal := structs[0]
												structs = structs[1:]

												structType := structVal.Type()
												for i := 0; i < structType.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:586
			_go_fuzz_dep_.CoverTab[122082]++
													fieldType := structType.Field(i)
													tagParts := strings.Split(fieldType.Tag.Get(tagName), ",")

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:591
			if tagParts[0] == "-" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:591
				_go_fuzz_dep_.CoverTab[122085]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:592
				// _ = "end of CoverTab[122085]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:593
				_go_fuzz_dep_.CoverTab[122086]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:593
				// _ = "end of CoverTab[122086]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:593
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:593
			// _ = "end of CoverTab[122082]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:593
			_go_fuzz_dep_.CoverTab[122083]++

													if fieldType.Anonymous {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:595
				_go_fuzz_dep_.CoverTab[122087]++
														fieldKind := fieldType.Type.Kind()
														if fieldKind != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:597
					_go_fuzz_dep_.CoverTab[122090]++
															return &parser.PosError{
						Pos:	node.Pos(),
						Err: fmt.Errorf("%s: unsupported type to struct: %s",
							fieldType.Name, fieldKind),
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:602
					// _ = "end of CoverTab[122090]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:603
					_go_fuzz_dep_.CoverTab[122091]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:603
					// _ = "end of CoverTab[122091]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:603
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:603
				// _ = "end of CoverTab[122087]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:603
				_go_fuzz_dep_.CoverTab[122088]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:607
				squash := false
				for _, tag := range tagParts[1:] {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:608
					_go_fuzz_dep_.CoverTab[122092]++
															if tag == "squash" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:609
						_go_fuzz_dep_.CoverTab[122093]++
																squash = true
																break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:611
						// _ = "end of CoverTab[122093]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:612
						_go_fuzz_dep_.CoverTab[122094]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:612
						// _ = "end of CoverTab[122094]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:612
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:612
					// _ = "end of CoverTab[122092]"
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:613
				// _ = "end of CoverTab[122088]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:613
				_go_fuzz_dep_.CoverTab[122089]++

														if squash {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:615
					_go_fuzz_dep_.CoverTab[122095]++
															structs = append(
						structs, result.FieldByName(fieldType.Name))
															continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:618
					// _ = "end of CoverTab[122095]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:619
					_go_fuzz_dep_.CoverTab[122096]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:619
					// _ = "end of CoverTab[122096]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:619
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:619
				// _ = "end of CoverTab[122089]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:620
				_go_fuzz_dep_.CoverTab[122097]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:620
				// _ = "end of CoverTab[122097]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:620
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:620
			// _ = "end of CoverTab[122083]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:620
			_go_fuzz_dep_.CoverTab[122084]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:623
			fields = append(fields, field{fieldType, structVal.Field(i)})
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:623
			// _ = "end of CoverTab[122084]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:624
		// _ = "end of CoverTab[122081]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:625
	// _ = "end of CoverTab[122068]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:625
	_go_fuzz_dep_.CoverTab[122069]++

											usedKeys := make(map[string]struct{})
											decodedFields := make([]string, 0, len(fields))
											decodedFieldsVal := make([]reflect.Value, 0)
											unusedKeysVal := make([]reflect.Value, 0)
											for _, f := range fields {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:631
		_go_fuzz_dep_.CoverTab[122098]++
												field, fieldValue := f.field, f.val
												if !fieldValue.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:633
			_go_fuzz_dep_.CoverTab[122106]++

													panic("field is not valid")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:635
			// _ = "end of CoverTab[122106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:636
			_go_fuzz_dep_.CoverTab[122107]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:636
			// _ = "end of CoverTab[122107]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:636
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:636
		// _ = "end of CoverTab[122098]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:636
		_go_fuzz_dep_.CoverTab[122099]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:640
		if !fieldValue.CanSet() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:640
			_go_fuzz_dep_.CoverTab[122108]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:641
			// _ = "end of CoverTab[122108]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:642
			_go_fuzz_dep_.CoverTab[122109]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:642
			// _ = "end of CoverTab[122109]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:642
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:642
		// _ = "end of CoverTab[122099]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:642
		_go_fuzz_dep_.CoverTab[122100]++

												fieldName := field.Name

												tagValue := field.Tag.Get(tagName)
												tagParts := strings.SplitN(tagValue, ",", 2)
												if len(tagParts) >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:648
			_go_fuzz_dep_.CoverTab[122110]++
													switch tagParts[1] {
			case "decodedFields":
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:650
				_go_fuzz_dep_.CoverTab[122111]++
														decodedFieldsVal = append(decodedFieldsVal, fieldValue)
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:652
				// _ = "end of CoverTab[122111]"
			case "key":
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:653
				_go_fuzz_dep_.CoverTab[122112]++
														if item == nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:654
					_go_fuzz_dep_.CoverTab[122116]++
															return &parser.PosError{
						Pos:	node.Pos(),
						Err: fmt.Errorf("%s: %s asked for 'key', impossible",
							name, fieldName),
					}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:659
					// _ = "end of CoverTab[122116]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:660
					_go_fuzz_dep_.CoverTab[122117]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:660
					// _ = "end of CoverTab[122117]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:660
				}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:660
				// _ = "end of CoverTab[122112]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:660
				_go_fuzz_dep_.CoverTab[122113]++

														fieldValue.SetString(item.Keys[0].Token.Value().(string))
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:663
				// _ = "end of CoverTab[122113]"
			case "unusedKeys":
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:664
				_go_fuzz_dep_.CoverTab[122114]++
														unusedKeysVal = append(unusedKeysVal, fieldValue)
														continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:666
				// _ = "end of CoverTab[122114]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:666
			default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:666
				_go_fuzz_dep_.CoverTab[122115]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:666
				// _ = "end of CoverTab[122115]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:667
			// _ = "end of CoverTab[122110]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:668
			_go_fuzz_dep_.CoverTab[122118]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:668
			// _ = "end of CoverTab[122118]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:668
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:668
		// _ = "end of CoverTab[122100]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:668
		_go_fuzz_dep_.CoverTab[122101]++

												if tagParts[0] != "" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:670
			_go_fuzz_dep_.CoverTab[122119]++
													fieldName = tagParts[0]
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:671
			// _ = "end of CoverTab[122119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:672
			_go_fuzz_dep_.CoverTab[122120]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:672
			// _ = "end of CoverTab[122120]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:672
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:672
		// _ = "end of CoverTab[122101]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:672
		_go_fuzz_dep_.CoverTab[122102]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:677
		filter := list.Filter(fieldName)

		prefixMatches := filter.Children()
		matches := filter.Elem()
		if len(matches.Items) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:681
			_go_fuzz_dep_.CoverTab[122121]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:681
			return len(prefixMatches.Items) == 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:681
			// _ = "end of CoverTab[122121]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:681
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:681
			_go_fuzz_dep_.CoverTab[122122]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:682
			// _ = "end of CoverTab[122122]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:683
			_go_fuzz_dep_.CoverTab[122123]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:683
			// _ = "end of CoverTab[122123]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:683
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:683
		// _ = "end of CoverTab[122102]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:683
		_go_fuzz_dep_.CoverTab[122103]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:686
		usedKeys[fieldName] = struct{}{}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:690
		fieldName = fmt.Sprintf("%s.%s", name, fieldName)
		if len(prefixMatches.Items) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:691
			_go_fuzz_dep_.CoverTab[122124]++
													if err := d.decode(fieldName, prefixMatches, fieldValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:692
				_go_fuzz_dep_.CoverTab[122125]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:693
				// _ = "end of CoverTab[122125]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:694
				_go_fuzz_dep_.CoverTab[122126]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:694
				// _ = "end of CoverTab[122126]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:694
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:694
			// _ = "end of CoverTab[122124]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:695
			_go_fuzz_dep_.CoverTab[122127]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:695
			// _ = "end of CoverTab[122127]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:695
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:695
		// _ = "end of CoverTab[122103]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:695
		_go_fuzz_dep_.CoverTab[122104]++
												for _, match := range matches.Items {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:696
			_go_fuzz_dep_.CoverTab[122128]++
													var decodeNode ast.Node = match.Val
													if ot, ok := decodeNode.(*ast.ObjectType); ok {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:698
				_go_fuzz_dep_.CoverTab[122130]++
														decodeNode = &ast.ObjectList{Items: ot.List.Items}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:699
				// _ = "end of CoverTab[122130]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:700
				_go_fuzz_dep_.CoverTab[122131]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:700
				// _ = "end of CoverTab[122131]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:700
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:700
			// _ = "end of CoverTab[122128]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:700
			_go_fuzz_dep_.CoverTab[122129]++

													if err := d.decode(fieldName, decodeNode, fieldValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:702
				_go_fuzz_dep_.CoverTab[122132]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:703
				// _ = "end of CoverTab[122132]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:704
				_go_fuzz_dep_.CoverTab[122133]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:704
				// _ = "end of CoverTab[122133]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:704
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:704
			// _ = "end of CoverTab[122129]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:705
		// _ = "end of CoverTab[122104]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:705
		_go_fuzz_dep_.CoverTab[122105]++

												decodedFields = append(decodedFields, field.Name)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:707
		// _ = "end of CoverTab[122105]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:708
	// _ = "end of CoverTab[122069]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:708
	_go_fuzz_dep_.CoverTab[122070]++

											if len(decodedFieldsVal) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:710
		_go_fuzz_dep_.CoverTab[122134]++

												sort.Strings(decodedFields)

												for _, v := range decodedFieldsVal {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:714
			_go_fuzz_dep_.CoverTab[122135]++
													v.Set(reflect.ValueOf(decodedFields))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:715
			// _ = "end of CoverTab[122135]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:716
		// _ = "end of CoverTab[122134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:717
		_go_fuzz_dep_.CoverTab[122136]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:717
		// _ = "end of CoverTab[122136]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:717
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:717
	// _ = "end of CoverTab[122070]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:717
	_go_fuzz_dep_.CoverTab[122071]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:719
	// _ = "end of CoverTab[122071]"
}

// findNodeType returns the type of ast.Node
func findNodeType() reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:723
	_go_fuzz_dep_.CoverTab[122137]++
											var nodeContainer struct {
		Node ast.Node
	}
											value := reflect.ValueOf(nodeContainer).FieldByName("Node")
											return value.Type()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:728
	// _ = "end of CoverTab[122137]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:729
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/decoder.go:729
var _ = _go_fuzz_dep_.CoverTab
