//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- float32Slice Value
type float32SliceValue struct {
	value	*[]float32
	changed	bool
}

func newFloat32SliceValue(val []float32, p *[]float32) *float32SliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:15
	_go_fuzz_dep_.CoverTab[119888]++
											isv := new(float32SliceValue)
											isv.value = p
											*isv.value = val
											return isv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:19
	// _ = "end of CoverTab[119888]"
}

func (s *float32SliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:22
	_go_fuzz_dep_.CoverTab[119889]++
											ss := strings.Split(val, ",")
											out := make([]float32, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:25
		_go_fuzz_dep_.CoverTab[119892]++
												var err error
												var temp64 float64
												temp64, err = strconv.ParseFloat(d, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:29
			_go_fuzz_dep_.CoverTab[119894]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:30
			// _ = "end of CoverTab[119894]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:31
			_go_fuzz_dep_.CoverTab[119895]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:31
			// _ = "end of CoverTab[119895]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:31
		// _ = "end of CoverTab[119892]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:31
		_go_fuzz_dep_.CoverTab[119893]++
												out[i] = float32(temp64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:32
		// _ = "end of CoverTab[119893]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:34
	// _ = "end of CoverTab[119889]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:34
	_go_fuzz_dep_.CoverTab[119890]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:35
		_go_fuzz_dep_.CoverTab[119896]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:36
		// _ = "end of CoverTab[119896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:37
		_go_fuzz_dep_.CoverTab[119897]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:38
		// _ = "end of CoverTab[119897]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:39
	// _ = "end of CoverTab[119890]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:39
	_go_fuzz_dep_.CoverTab[119891]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:41
	// _ = "end of CoverTab[119891]"
}

func (s *float32SliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:44
	_go_fuzz_dep_.CoverTab[119898]++
											return "float32Slice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:45
	// _ = "end of CoverTab[119898]"
}

func (s *float32SliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:48
	_go_fuzz_dep_.CoverTab[119899]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:50
		_go_fuzz_dep_.CoverTab[119901]++
												out[i] = fmt.Sprintf("%f", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:51
		// _ = "end of CoverTab[119901]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:52
	// _ = "end of CoverTab[119899]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:52
	_go_fuzz_dep_.CoverTab[119900]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:53
	// _ = "end of CoverTab[119900]"
}

func (s *float32SliceValue) fromString(val string) (float32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:56
	_go_fuzz_dep_.CoverTab[119902]++
											t64, err := strconv.ParseFloat(val, 32)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:58
		_go_fuzz_dep_.CoverTab[119904]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:59
		// _ = "end of CoverTab[119904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:60
		_go_fuzz_dep_.CoverTab[119905]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:60
		// _ = "end of CoverTab[119905]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:60
	// _ = "end of CoverTab[119902]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:60
	_go_fuzz_dep_.CoverTab[119903]++
											return float32(t64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:61
	// _ = "end of CoverTab[119903]"
}

func (s *float32SliceValue) toString(val float32) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:64
	_go_fuzz_dep_.CoverTab[119906]++
											return fmt.Sprintf("%f", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:65
	// _ = "end of CoverTab[119906]"
}

func (s *float32SliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:68
	_go_fuzz_dep_.CoverTab[119907]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:70
		_go_fuzz_dep_.CoverTab[119909]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:71
		// _ = "end of CoverTab[119909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:72
		_go_fuzz_dep_.CoverTab[119910]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:72
		// _ = "end of CoverTab[119910]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:72
	// _ = "end of CoverTab[119907]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:72
	_go_fuzz_dep_.CoverTab[119908]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:74
	// _ = "end of CoverTab[119908]"
}

func (s *float32SliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:77
	_go_fuzz_dep_.CoverTab[119911]++
											out := make([]float32, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:79
		_go_fuzz_dep_.CoverTab[119913]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:82
			_go_fuzz_dep_.CoverTab[119914]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:83
			// _ = "end of CoverTab[119914]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:84
			_go_fuzz_dep_.CoverTab[119915]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:84
			// _ = "end of CoverTab[119915]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:84
		// _ = "end of CoverTab[119913]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:85
	// _ = "end of CoverTab[119911]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:85
	_go_fuzz_dep_.CoverTab[119912]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:87
	// _ = "end of CoverTab[119912]"
}

func (s *float32SliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:90
	_go_fuzz_dep_.CoverTab[119916]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:92
		_go_fuzz_dep_.CoverTab[119918]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:93
		// _ = "end of CoverTab[119918]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:94
	// _ = "end of CoverTab[119916]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:94
	_go_fuzz_dep_.CoverTab[119917]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:95
	// _ = "end of CoverTab[119917]"
}

func float32SliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:98
	_go_fuzz_dep_.CoverTab[119919]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:101
		_go_fuzz_dep_.CoverTab[119922]++
												return []float32{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:102
		// _ = "end of CoverTab[119922]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:103
		_go_fuzz_dep_.CoverTab[119923]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:103
		// _ = "end of CoverTab[119923]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:103
	// _ = "end of CoverTab[119919]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:103
	_go_fuzz_dep_.CoverTab[119920]++
											ss := strings.Split(val, ",")
											out := make([]float32, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:106
		_go_fuzz_dep_.CoverTab[119924]++
												var err error
												var temp64 float64
												temp64, err = strconv.ParseFloat(d, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:110
			_go_fuzz_dep_.CoverTab[119926]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:111
			// _ = "end of CoverTab[119926]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:112
			_go_fuzz_dep_.CoverTab[119927]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:112
			// _ = "end of CoverTab[119927]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:112
		// _ = "end of CoverTab[119924]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:112
		_go_fuzz_dep_.CoverTab[119925]++
												out[i] = float32(temp64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:113
		// _ = "end of CoverTab[119925]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:115
	// _ = "end of CoverTab[119920]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:115
	_go_fuzz_dep_.CoverTab[119921]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:116
	// _ = "end of CoverTab[119921]"
}

// GetFloat32Slice return the []float32 value of a flag with the given name
func (f *FlagSet) GetFloat32Slice(name string) ([]float32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:120
	_go_fuzz_dep_.CoverTab[119928]++
											val, err := f.getFlagType(name, "float32Slice", float32SliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:122
		_go_fuzz_dep_.CoverTab[119930]++
												return []float32{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:123
		// _ = "end of CoverTab[119930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:124
		_go_fuzz_dep_.CoverTab[119931]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:124
		// _ = "end of CoverTab[119931]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:124
	// _ = "end of CoverTab[119928]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:124
	_go_fuzz_dep_.CoverTab[119929]++
											return val.([]float32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:125
	// _ = "end of CoverTab[119929]"
}

// Float32SliceVar defines a float32Slice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:128
// The argument p points to a []float32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:130
func (f *FlagSet) Float32SliceVar(p *[]float32, name string, value []float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:130
	_go_fuzz_dep_.CoverTab[119932]++
											f.VarP(newFloat32SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:131
	// _ = "end of CoverTab[119932]"
}

// Float32SliceVarP is like Float32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:135
	_go_fuzz_dep_.CoverTab[119933]++
											f.VarP(newFloat32SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:136
	// _ = "end of CoverTab[119933]"
}

// Float32SliceVar defines a float32[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:139
// The argument p points to a float32[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:141
func Float32SliceVar(p *[]float32, name string, value []float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:141
	_go_fuzz_dep_.CoverTab[119934]++
											CommandLine.VarP(newFloat32SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:142
	// _ = "end of CoverTab[119934]"
}

// Float32SliceVarP is like Float32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:146
	_go_fuzz_dep_.CoverTab[119935]++
											CommandLine.VarP(newFloat32SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:147
	// _ = "end of CoverTab[119935]"
}

// Float32Slice defines a []float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:150
// The return value is the address of a []float32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:152
func (f *FlagSet) Float32Slice(name string, value []float32, usage string) *[]float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:152
	_go_fuzz_dep_.CoverTab[119936]++
											p := []float32{}
											f.Float32SliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:155
	// _ = "end of CoverTab[119936]"
}

// Float32SliceP is like Float32Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:159
	_go_fuzz_dep_.CoverTab[119937]++
											p := []float32{}
											f.Float32SliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:162
	// _ = "end of CoverTab[119937]"
}

// Float32Slice defines a []float32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:165
// The return value is the address of a []float32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:167
func Float32Slice(name string, value []float32, usage string) *[]float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:167
	_go_fuzz_dep_.CoverTab[119938]++
											return CommandLine.Float32SliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:168
	// _ = "end of CoverTab[119938]"
}

// Float32SliceP is like Float32Slice, but accepts a shorthand letter that can be used after a single dash.
func Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:172
	_go_fuzz_dep_.CoverTab[119939]++
											return CommandLine.Float32SliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:173
	// _ = "end of CoverTab[119939]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:174
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float32_slice.go:174
var _ = _go_fuzz_dep_.CoverTab
