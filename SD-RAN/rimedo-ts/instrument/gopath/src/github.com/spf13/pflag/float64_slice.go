//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- float64Slice Value
type float64SliceValue struct {
	value	*[]float64
	changed	bool
}

func newFloat64SliceValue(val []float64, p *[]float64) *float64SliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:15
	_go_fuzz_dep_.CoverTab[119957]++
											isv := new(float64SliceValue)
											isv.value = p
											*isv.value = val
											return isv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:19
	// _ = "end of CoverTab[119957]"
}

func (s *float64SliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:22
	_go_fuzz_dep_.CoverTab[119958]++
											ss := strings.Split(val, ",")
											out := make([]float64, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:25
		_go_fuzz_dep_.CoverTab[119961]++
												var err error
												out[i], err = strconv.ParseFloat(d, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:28
			_go_fuzz_dep_.CoverTab[119962]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:29
			// _ = "end of CoverTab[119962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:30
			_go_fuzz_dep_.CoverTab[119963]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:30
			// _ = "end of CoverTab[119963]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:30
		// _ = "end of CoverTab[119961]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:32
	// _ = "end of CoverTab[119958]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:32
	_go_fuzz_dep_.CoverTab[119959]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:33
		_go_fuzz_dep_.CoverTab[119964]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:34
		// _ = "end of CoverTab[119964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:35
		_go_fuzz_dep_.CoverTab[119965]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:36
		// _ = "end of CoverTab[119965]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:37
	// _ = "end of CoverTab[119959]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:37
	_go_fuzz_dep_.CoverTab[119960]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:39
	// _ = "end of CoverTab[119960]"
}

func (s *float64SliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:42
	_go_fuzz_dep_.CoverTab[119966]++
											return "float64Slice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:43
	// _ = "end of CoverTab[119966]"
}

func (s *float64SliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:46
	_go_fuzz_dep_.CoverTab[119967]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:48
		_go_fuzz_dep_.CoverTab[119969]++
												out[i] = fmt.Sprintf("%f", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:49
		// _ = "end of CoverTab[119969]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:50
	// _ = "end of CoverTab[119967]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:50
	_go_fuzz_dep_.CoverTab[119968]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:51
	// _ = "end of CoverTab[119968]"
}

func (s *float64SliceValue) fromString(val string) (float64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:54
	_go_fuzz_dep_.CoverTab[119970]++
											return strconv.ParseFloat(val, 64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:55
	// _ = "end of CoverTab[119970]"
}

func (s *float64SliceValue) toString(val float64) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:58
	_go_fuzz_dep_.CoverTab[119971]++
											return fmt.Sprintf("%f", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:59
	// _ = "end of CoverTab[119971]"
}

func (s *float64SliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:62
	_go_fuzz_dep_.CoverTab[119972]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:64
		_go_fuzz_dep_.CoverTab[119974]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:65
		// _ = "end of CoverTab[119974]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:66
		_go_fuzz_dep_.CoverTab[119975]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:66
		// _ = "end of CoverTab[119975]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:66
	// _ = "end of CoverTab[119972]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:66
	_go_fuzz_dep_.CoverTab[119973]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:68
	// _ = "end of CoverTab[119973]"
}

func (s *float64SliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:71
	_go_fuzz_dep_.CoverTab[119976]++
											out := make([]float64, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:73
		_go_fuzz_dep_.CoverTab[119978]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:76
			_go_fuzz_dep_.CoverTab[119979]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:77
			// _ = "end of CoverTab[119979]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:78
			_go_fuzz_dep_.CoverTab[119980]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:78
			// _ = "end of CoverTab[119980]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:78
		// _ = "end of CoverTab[119978]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:79
	// _ = "end of CoverTab[119976]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:79
	_go_fuzz_dep_.CoverTab[119977]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:81
	// _ = "end of CoverTab[119977]"
}

func (s *float64SliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:84
	_go_fuzz_dep_.CoverTab[119981]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:86
		_go_fuzz_dep_.CoverTab[119983]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:87
		// _ = "end of CoverTab[119983]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:88
	// _ = "end of CoverTab[119981]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:88
	_go_fuzz_dep_.CoverTab[119982]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:89
	// _ = "end of CoverTab[119982]"
}

func float64SliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:92
	_go_fuzz_dep_.CoverTab[119984]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:95
		_go_fuzz_dep_.CoverTab[119987]++
												return []float64{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:96
		// _ = "end of CoverTab[119987]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:97
		_go_fuzz_dep_.CoverTab[119988]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:97
		// _ = "end of CoverTab[119988]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:97
	// _ = "end of CoverTab[119984]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:97
	_go_fuzz_dep_.CoverTab[119985]++
											ss := strings.Split(val, ",")
											out := make([]float64, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:100
		_go_fuzz_dep_.CoverTab[119989]++
												var err error
												out[i], err = strconv.ParseFloat(d, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:103
			_go_fuzz_dep_.CoverTab[119990]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:104
			// _ = "end of CoverTab[119990]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:105
			_go_fuzz_dep_.CoverTab[119991]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:105
			// _ = "end of CoverTab[119991]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:105
		// _ = "end of CoverTab[119989]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:107
	// _ = "end of CoverTab[119985]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:107
	_go_fuzz_dep_.CoverTab[119986]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:108
	// _ = "end of CoverTab[119986]"
}

// GetFloat64Slice return the []float64 value of a flag with the given name
func (f *FlagSet) GetFloat64Slice(name string) ([]float64, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:112
	_go_fuzz_dep_.CoverTab[119992]++
											val, err := f.getFlagType(name, "float64Slice", float64SliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:114
		_go_fuzz_dep_.CoverTab[119994]++
												return []float64{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:115
		// _ = "end of CoverTab[119994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:116
		_go_fuzz_dep_.CoverTab[119995]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:116
		// _ = "end of CoverTab[119995]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:116
	// _ = "end of CoverTab[119992]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:116
	_go_fuzz_dep_.CoverTab[119993]++
											return val.([]float64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:117
	// _ = "end of CoverTab[119993]"
}

// Float64SliceVar defines a float64Slice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:120
// The argument p points to a []float64 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:122
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:122
	_go_fuzz_dep_.CoverTab[119996]++
											f.VarP(newFloat64SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:123
	// _ = "end of CoverTab[119996]"
}

// Float64SliceVarP is like Float64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:127
	_go_fuzz_dep_.CoverTab[119997]++
											f.VarP(newFloat64SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:128
	// _ = "end of CoverTab[119997]"
}

// Float64SliceVar defines a float64[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:131
// The argument p points to a float64[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:133
func Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:133
	_go_fuzz_dep_.CoverTab[119998]++
											CommandLine.VarP(newFloat64SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:134
	// _ = "end of CoverTab[119998]"
}

// Float64SliceVarP is like Float64SliceVar, but accepts a shorthand letter that can be used after a single dash.
func Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:138
	_go_fuzz_dep_.CoverTab[119999]++
											CommandLine.VarP(newFloat64SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:139
	// _ = "end of CoverTab[119999]"
}

// Float64Slice defines a []float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:142
// The return value is the address of a []float64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:144
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:144
	_go_fuzz_dep_.CoverTab[120000]++
											p := []float64{}
											f.Float64SliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:147
	// _ = "end of CoverTab[120000]"
}

// Float64SliceP is like Float64Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:151
	_go_fuzz_dep_.CoverTab[120001]++
											p := []float64{}
											f.Float64SliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:154
	// _ = "end of CoverTab[120001]"
}

// Float64Slice defines a []float64 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:157
// The return value is the address of a []float64 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:159
func Float64Slice(name string, value []float64, usage string) *[]float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:159
	_go_fuzz_dep_.CoverTab[120002]++
											return CommandLine.Float64SliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:160
	// _ = "end of CoverTab[120002]"
}

// Float64SliceP is like Float64Slice, but accepts a shorthand letter that can be used after a single dash.
func Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:164
	_go_fuzz_dep_.CoverTab[120003]++
											return CommandLine.Float64SliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:165
	// _ = "end of CoverTab[120003]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:166
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/float64_slice.go:166
var _ = _go_fuzz_dep_.CoverTab
