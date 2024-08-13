//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- int32Slice Value
type int32SliceValue struct {
	value	*[]int32
	changed	bool
}

func newInt32SliceValue(val []int32, p *[]int32) *int32SliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:15
	_go_fuzz_dep_.CoverTab[120093]++
											isv := new(int32SliceValue)
											isv.value = p
											*isv.value = val
											return isv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:19
	// _ = "end of CoverTab[120093]"
}

func (s *int32SliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:22
	_go_fuzz_dep_.CoverTab[120094]++
											ss := strings.Split(val, ",")
											out := make([]int32, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:25
		_go_fuzz_dep_.CoverTab[120097]++
												var err error
												var temp64 int64
												temp64, err = strconv.ParseInt(d, 0, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:29
			_go_fuzz_dep_.CoverTab[120099]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:30
			// _ = "end of CoverTab[120099]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:31
			_go_fuzz_dep_.CoverTab[120100]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:31
			// _ = "end of CoverTab[120100]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:31
		// _ = "end of CoverTab[120097]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:31
		_go_fuzz_dep_.CoverTab[120098]++
												out[i] = int32(temp64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:32
		// _ = "end of CoverTab[120098]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:34
	// _ = "end of CoverTab[120094]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:34
	_go_fuzz_dep_.CoverTab[120095]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:35
		_go_fuzz_dep_.CoverTab[120101]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:36
		// _ = "end of CoverTab[120101]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:37
		_go_fuzz_dep_.CoverTab[120102]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:38
		// _ = "end of CoverTab[120102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:39
	// _ = "end of CoverTab[120095]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:39
	_go_fuzz_dep_.CoverTab[120096]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:41
	// _ = "end of CoverTab[120096]"
}

func (s *int32SliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:44
	_go_fuzz_dep_.CoverTab[120103]++
											return "int32Slice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:45
	// _ = "end of CoverTab[120103]"
}

func (s *int32SliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:48
	_go_fuzz_dep_.CoverTab[120104]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:50
		_go_fuzz_dep_.CoverTab[120106]++
												out[i] = fmt.Sprintf("%d", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:51
		// _ = "end of CoverTab[120106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:52
	// _ = "end of CoverTab[120104]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:52
	_go_fuzz_dep_.CoverTab[120105]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:53
	// _ = "end of CoverTab[120105]"
}

func (s *int32SliceValue) fromString(val string) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:56
	_go_fuzz_dep_.CoverTab[120107]++
											t64, err := strconv.ParseInt(val, 0, 32)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:58
		_go_fuzz_dep_.CoverTab[120109]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:59
		// _ = "end of CoverTab[120109]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:60
		_go_fuzz_dep_.CoverTab[120110]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:60
		// _ = "end of CoverTab[120110]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:60
	// _ = "end of CoverTab[120107]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:60
	_go_fuzz_dep_.CoverTab[120108]++
											return int32(t64), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:61
	// _ = "end of CoverTab[120108]"
}

func (s *int32SliceValue) toString(val int32) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:64
	_go_fuzz_dep_.CoverTab[120111]++
											return fmt.Sprintf("%d", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:65
	// _ = "end of CoverTab[120111]"
}

func (s *int32SliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:68
	_go_fuzz_dep_.CoverTab[120112]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:70
		_go_fuzz_dep_.CoverTab[120114]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:71
		// _ = "end of CoverTab[120114]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:72
		_go_fuzz_dep_.CoverTab[120115]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:72
		// _ = "end of CoverTab[120115]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:72
	// _ = "end of CoverTab[120112]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:72
	_go_fuzz_dep_.CoverTab[120113]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:74
	// _ = "end of CoverTab[120113]"
}

func (s *int32SliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:77
	_go_fuzz_dep_.CoverTab[120116]++
											out := make([]int32, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:79
		_go_fuzz_dep_.CoverTab[120118]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:82
			_go_fuzz_dep_.CoverTab[120119]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:83
			// _ = "end of CoverTab[120119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:84
			_go_fuzz_dep_.CoverTab[120120]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:84
			// _ = "end of CoverTab[120120]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:84
		// _ = "end of CoverTab[120118]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:85
	// _ = "end of CoverTab[120116]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:85
	_go_fuzz_dep_.CoverTab[120117]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:87
	// _ = "end of CoverTab[120117]"
}

func (s *int32SliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:90
	_go_fuzz_dep_.CoverTab[120121]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:92
		_go_fuzz_dep_.CoverTab[120123]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:93
		// _ = "end of CoverTab[120123]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:94
	// _ = "end of CoverTab[120121]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:94
	_go_fuzz_dep_.CoverTab[120122]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:95
	// _ = "end of CoverTab[120122]"
}

func int32SliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:98
	_go_fuzz_dep_.CoverTab[120124]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:101
		_go_fuzz_dep_.CoverTab[120127]++
												return []int32{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:102
		// _ = "end of CoverTab[120127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:103
		_go_fuzz_dep_.CoverTab[120128]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:103
		// _ = "end of CoverTab[120128]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:103
	// _ = "end of CoverTab[120124]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:103
	_go_fuzz_dep_.CoverTab[120125]++
											ss := strings.Split(val, ",")
											out := make([]int32, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:106
		_go_fuzz_dep_.CoverTab[120129]++
												var err error
												var temp64 int64
												temp64, err = strconv.ParseInt(d, 0, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:110
			_go_fuzz_dep_.CoverTab[120131]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:111
			// _ = "end of CoverTab[120131]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:112
			_go_fuzz_dep_.CoverTab[120132]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:112
			// _ = "end of CoverTab[120132]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:112
		// _ = "end of CoverTab[120129]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:112
		_go_fuzz_dep_.CoverTab[120130]++
												out[i] = int32(temp64)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:113
		// _ = "end of CoverTab[120130]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:115
	// _ = "end of CoverTab[120125]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:115
	_go_fuzz_dep_.CoverTab[120126]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:116
	// _ = "end of CoverTab[120126]"
}

// GetInt32Slice return the []int32 value of a flag with the given name
func (f *FlagSet) GetInt32Slice(name string) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:120
	_go_fuzz_dep_.CoverTab[120133]++
											val, err := f.getFlagType(name, "int32Slice", int32SliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:122
		_go_fuzz_dep_.CoverTab[120135]++
												return []int32{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:123
		// _ = "end of CoverTab[120135]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:124
		_go_fuzz_dep_.CoverTab[120136]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:124
		// _ = "end of CoverTab[120136]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:124
	// _ = "end of CoverTab[120133]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:124
	_go_fuzz_dep_.CoverTab[120134]++
											return val.([]int32), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:125
	// _ = "end of CoverTab[120134]"
}

// Int32SliceVar defines a int32Slice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:128
// The argument p points to a []int32 variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:130
func (f *FlagSet) Int32SliceVar(p *[]int32, name string, value []int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:130
	_go_fuzz_dep_.CoverTab[120137]++
											f.VarP(newInt32SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:131
	// _ = "end of CoverTab[120137]"
}

// Int32SliceVarP is like Int32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:135
	_go_fuzz_dep_.CoverTab[120138]++
											f.VarP(newInt32SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:136
	// _ = "end of CoverTab[120138]"
}

// Int32SliceVar defines a int32[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:139
// The argument p points to a int32[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:141
func Int32SliceVar(p *[]int32, name string, value []int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:141
	_go_fuzz_dep_.CoverTab[120139]++
											CommandLine.VarP(newInt32SliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:142
	// _ = "end of CoverTab[120139]"
}

// Int32SliceVarP is like Int32SliceVar, but accepts a shorthand letter that can be used after a single dash.
func Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:146
	_go_fuzz_dep_.CoverTab[120140]++
											CommandLine.VarP(newInt32SliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:147
	// _ = "end of CoverTab[120140]"
}

// Int32Slice defines a []int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:150
// The return value is the address of a []int32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:152
func (f *FlagSet) Int32Slice(name string, value []int32, usage string) *[]int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:152
	_go_fuzz_dep_.CoverTab[120141]++
											p := []int32{}
											f.Int32SliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:155
	// _ = "end of CoverTab[120141]"
}

// Int32SliceP is like Int32Slice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:159
	_go_fuzz_dep_.CoverTab[120142]++
											p := []int32{}
											f.Int32SliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:162
	// _ = "end of CoverTab[120142]"
}

// Int32Slice defines a []int32 flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:165
// The return value is the address of a []int32 variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:167
func Int32Slice(name string, value []int32, usage string) *[]int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:167
	_go_fuzz_dep_.CoverTab[120143]++
											return CommandLine.Int32SliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:168
	// _ = "end of CoverTab[120143]"
}

// Int32SliceP is like Int32Slice, but accepts a shorthand letter that can be used after a single dash.
func Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:172
	_go_fuzz_dep_.CoverTab[120144]++
											return CommandLine.Int32SliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:173
	// _ = "end of CoverTab[120144]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:174
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int32_slice.go:174
var _ = _go_fuzz_dep_.CoverTab
