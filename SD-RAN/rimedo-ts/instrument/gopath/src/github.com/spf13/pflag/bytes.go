//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:1
)

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// BytesHex adapts []byte for use as a flag. Value of flag is HEX encoded
type bytesHexValue []byte

// String implements pflag.Value.String.
func (bytesHex bytesHexValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:14
	_go_fuzz_dep_.CoverTab[119336]++
										return fmt.Sprintf("%X", []byte(bytesHex))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:15
	// _ = "end of CoverTab[119336]"
}

// Set implements pflag.Value.Set.
func (bytesHex *bytesHexValue) Set(value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:19
	_go_fuzz_dep_.CoverTab[119337]++
										bin, err := hex.DecodeString(strings.TrimSpace(value))

										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:22
		_go_fuzz_dep_.CoverTab[119339]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:23
		// _ = "end of CoverTab[119339]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:24
		_go_fuzz_dep_.CoverTab[119340]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:24
		// _ = "end of CoverTab[119340]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:24
	// _ = "end of CoverTab[119337]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:24
	_go_fuzz_dep_.CoverTab[119338]++

										*bytesHex = bin

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:28
	// _ = "end of CoverTab[119338]"
}

// Type implements pflag.Value.Type.
func (*bytesHexValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:32
	_go_fuzz_dep_.CoverTab[119341]++
										return "bytesHex"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:33
	// _ = "end of CoverTab[119341]"
}

func newBytesHexValue(val []byte, p *[]byte) *bytesHexValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:36
	_go_fuzz_dep_.CoverTab[119342]++
										*p = val
										return (*bytesHexValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:38
	// _ = "end of CoverTab[119342]"
}

func bytesHexConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:41
	_go_fuzz_dep_.CoverTab[119343]++

										bin, err := hex.DecodeString(sval)

										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:45
		_go_fuzz_dep_.CoverTab[119345]++
											return bin, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:46
		// _ = "end of CoverTab[119345]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:47
		_go_fuzz_dep_.CoverTab[119346]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:47
		// _ = "end of CoverTab[119346]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:47
	// _ = "end of CoverTab[119343]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:47
	_go_fuzz_dep_.CoverTab[119344]++

										return nil, fmt.Errorf("invalid string being converted to Bytes: %s %s", sval, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:49
	// _ = "end of CoverTab[119344]"
}

// GetBytesHex return the []byte value of a flag with the given name
func (f *FlagSet) GetBytesHex(name string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:53
	_go_fuzz_dep_.CoverTab[119347]++
										val, err := f.getFlagType(name, "bytesHex", bytesHexConv)

										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:56
		_go_fuzz_dep_.CoverTab[119349]++
											return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:57
		// _ = "end of CoverTab[119349]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:58
		_go_fuzz_dep_.CoverTab[119350]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:58
		// _ = "end of CoverTab[119350]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:58
	// _ = "end of CoverTab[119347]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:58
	_go_fuzz_dep_.CoverTab[119348]++

										return val.([]byte), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:60
	// _ = "end of CoverTab[119348]"
}

// BytesHexVar defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:63
// The argument p points to an []byte variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:65
func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:65
	_go_fuzz_dep_.CoverTab[119351]++
										f.VarP(newBytesHexValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:66
	// _ = "end of CoverTab[119351]"
}

// BytesHexVarP is like BytesHexVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:70
	_go_fuzz_dep_.CoverTab[119352]++
										f.VarP(newBytesHexValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:71
	// _ = "end of CoverTab[119352]"
}

// BytesHexVar defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:74
// The argument p points to an []byte variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:76
func BytesHexVar(p *[]byte, name string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:76
	_go_fuzz_dep_.CoverTab[119353]++
										CommandLine.VarP(newBytesHexValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:77
	// _ = "end of CoverTab[119353]"
}

// BytesHexVarP is like BytesHexVar, but accepts a shorthand letter that can be used after a single dash.
func BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:81
	_go_fuzz_dep_.CoverTab[119354]++
										CommandLine.VarP(newBytesHexValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:82
	// _ = "end of CoverTab[119354]"
}

// BytesHex defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:85
// The return value is the address of an []byte variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:87
func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:87
	_go_fuzz_dep_.CoverTab[119355]++
										p := new([]byte)
										f.BytesHexVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:90
	// _ = "end of CoverTab[119355]"
}

// BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:94
	_go_fuzz_dep_.CoverTab[119356]++
										p := new([]byte)
										f.BytesHexVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:97
	// _ = "end of CoverTab[119356]"
}

// BytesHex defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:100
// The return value is the address of an []byte variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:102
func BytesHex(name string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:102
	_go_fuzz_dep_.CoverTab[119357]++
										return CommandLine.BytesHexP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:103
	// _ = "end of CoverTab[119357]"
}

// BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.
func BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:107
	_go_fuzz_dep_.CoverTab[119358]++
										return CommandLine.BytesHexP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:108
	// _ = "end of CoverTab[119358]"
}

// BytesBase64 adapts []byte for use as a flag. Value of flag is Base64 encoded
type bytesBase64Value []byte

// String implements pflag.Value.String.
func (bytesBase64 bytesBase64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:115
	_go_fuzz_dep_.CoverTab[119359]++
										return base64.StdEncoding.EncodeToString([]byte(bytesBase64))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:116
	// _ = "end of CoverTab[119359]"
}

// Set implements pflag.Value.Set.
func (bytesBase64 *bytesBase64Value) Set(value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:120
	_go_fuzz_dep_.CoverTab[119360]++
										bin, err := base64.StdEncoding.DecodeString(strings.TrimSpace(value))

										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:123
		_go_fuzz_dep_.CoverTab[119362]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:124
		// _ = "end of CoverTab[119362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:125
		_go_fuzz_dep_.CoverTab[119363]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:125
		// _ = "end of CoverTab[119363]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:125
	// _ = "end of CoverTab[119360]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:125
	_go_fuzz_dep_.CoverTab[119361]++

										*bytesBase64 = bin

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:129
	// _ = "end of CoverTab[119361]"
}

// Type implements pflag.Value.Type.
func (*bytesBase64Value) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:133
	_go_fuzz_dep_.CoverTab[119364]++
										return "bytesBase64"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:134
	// _ = "end of CoverTab[119364]"
}

func newBytesBase64Value(val []byte, p *[]byte) *bytesBase64Value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:137
	_go_fuzz_dep_.CoverTab[119365]++
										*p = val
										return (*bytesBase64Value)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:139
	// _ = "end of CoverTab[119365]"
}

func bytesBase64ValueConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:142
	_go_fuzz_dep_.CoverTab[119366]++

										bin, err := base64.StdEncoding.DecodeString(sval)
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:145
		_go_fuzz_dep_.CoverTab[119368]++
											return bin, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:146
		// _ = "end of CoverTab[119368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:147
		_go_fuzz_dep_.CoverTab[119369]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:147
		// _ = "end of CoverTab[119369]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:147
	// _ = "end of CoverTab[119366]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:147
	_go_fuzz_dep_.CoverTab[119367]++

										return nil, fmt.Errorf("invalid string being converted to Bytes: %s %s", sval, err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:149
	// _ = "end of CoverTab[119367]"
}

// GetBytesBase64 return the []byte value of a flag with the given name
func (f *FlagSet) GetBytesBase64(name string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:153
	_go_fuzz_dep_.CoverTab[119370]++
										val, err := f.getFlagType(name, "bytesBase64", bytesBase64ValueConv)

										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:156
		_go_fuzz_dep_.CoverTab[119372]++
											return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:157
		// _ = "end of CoverTab[119372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:158
		_go_fuzz_dep_.CoverTab[119373]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:158
		// _ = "end of CoverTab[119373]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:158
	// _ = "end of CoverTab[119370]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:158
	_go_fuzz_dep_.CoverTab[119371]++

										return val.([]byte), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:160
	// _ = "end of CoverTab[119371]"
}

// BytesBase64Var defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:163
// The argument p points to an []byte variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:165
func (f *FlagSet) BytesBase64Var(p *[]byte, name string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:165
	_go_fuzz_dep_.CoverTab[119374]++
										f.VarP(newBytesBase64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:166
	// _ = "end of CoverTab[119374]"
}

// BytesBase64VarP is like BytesBase64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:170
	_go_fuzz_dep_.CoverTab[119375]++
										f.VarP(newBytesBase64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:171
	// _ = "end of CoverTab[119375]"
}

// BytesBase64Var defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:174
// The argument p points to an []byte variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:176
func BytesBase64Var(p *[]byte, name string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:176
	_go_fuzz_dep_.CoverTab[119376]++
										CommandLine.VarP(newBytesBase64Value(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:177
	// _ = "end of CoverTab[119376]"
}

// BytesBase64VarP is like BytesBase64Var, but accepts a shorthand letter that can be used after a single dash.
func BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:181
	_go_fuzz_dep_.CoverTab[119377]++
										CommandLine.VarP(newBytesBase64Value(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:182
	// _ = "end of CoverTab[119377]"
}

// BytesBase64 defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:185
// The return value is the address of an []byte variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:187
func (f *FlagSet) BytesBase64(name string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:187
	_go_fuzz_dep_.CoverTab[119378]++
										p := new([]byte)
										f.BytesBase64VarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:190
	// _ = "end of CoverTab[119378]"
}

// BytesBase64P is like BytesBase64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:194
	_go_fuzz_dep_.CoverTab[119379]++
										p := new([]byte)
										f.BytesBase64VarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:197
	// _ = "end of CoverTab[119379]"
}

// BytesBase64 defines an []byte flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:200
// The return value is the address of an []byte variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:202
func BytesBase64(name string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:202
	_go_fuzz_dep_.CoverTab[119380]++
										return CommandLine.BytesBase64P(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:203
	// _ = "end of CoverTab[119380]"
}

// BytesBase64P is like BytesBase64, but accepts a shorthand letter that can be used after a single dash.
func BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:207
	_go_fuzz_dep_.CoverTab[119381]++
										return CommandLine.BytesBase64P(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:208
	// _ = "end of CoverTab[119381]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:209
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bytes.go:209
var _ = _go_fuzz_dep_.CoverTab
