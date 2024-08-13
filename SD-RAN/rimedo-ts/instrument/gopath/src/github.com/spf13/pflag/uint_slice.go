//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- uintSlice Value
type uintSliceValue struct {
	value	*[]uint
	changed	bool
}

func newUintSliceValue(val []uint, p *[]uint) *uintSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:15
	_go_fuzz_dep_.CoverTab[120730]++
											uisv := new(uintSliceValue)
											uisv.value = p
											*uisv.value = val
											return uisv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:19
	// _ = "end of CoverTab[120730]"
}

func (s *uintSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:22
	_go_fuzz_dep_.CoverTab[120731]++
											ss := strings.Split(val, ",")
											out := make([]uint, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:25
		_go_fuzz_dep_.CoverTab[120734]++
												u, err := strconv.ParseUint(d, 10, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:27
			_go_fuzz_dep_.CoverTab[120736]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:28
			// _ = "end of CoverTab[120736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:29
			_go_fuzz_dep_.CoverTab[120737]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:29
			// _ = "end of CoverTab[120737]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:29
		// _ = "end of CoverTab[120734]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:29
		_go_fuzz_dep_.CoverTab[120735]++
												out[i] = uint(u)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:30
		// _ = "end of CoverTab[120735]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:31
	// _ = "end of CoverTab[120731]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:31
	_go_fuzz_dep_.CoverTab[120732]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:32
		_go_fuzz_dep_.CoverTab[120738]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:33
		// _ = "end of CoverTab[120738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:34
		_go_fuzz_dep_.CoverTab[120739]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:35
		// _ = "end of CoverTab[120739]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:36
	// _ = "end of CoverTab[120732]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:36
	_go_fuzz_dep_.CoverTab[120733]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:38
	// _ = "end of CoverTab[120733]"
}

func (s *uintSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:41
	_go_fuzz_dep_.CoverTab[120740]++
											return "uintSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:42
	// _ = "end of CoverTab[120740]"
}

func (s *uintSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:45
	_go_fuzz_dep_.CoverTab[120741]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:47
		_go_fuzz_dep_.CoverTab[120743]++
												out[i] = fmt.Sprintf("%d", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:48
		// _ = "end of CoverTab[120743]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:49
	// _ = "end of CoverTab[120741]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:49
	_go_fuzz_dep_.CoverTab[120742]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:50
	// _ = "end of CoverTab[120742]"
}

func (s *uintSliceValue) fromString(val string) (uint, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:53
	_go_fuzz_dep_.CoverTab[120744]++
											t, err := strconv.ParseUint(val, 10, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:55
		_go_fuzz_dep_.CoverTab[120746]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:56
		// _ = "end of CoverTab[120746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:57
		_go_fuzz_dep_.CoverTab[120747]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:57
		// _ = "end of CoverTab[120747]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:57
	// _ = "end of CoverTab[120744]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:57
	_go_fuzz_dep_.CoverTab[120745]++
											return uint(t), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:58
	// _ = "end of CoverTab[120745]"
}

func (s *uintSliceValue) toString(val uint) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:61
	_go_fuzz_dep_.CoverTab[120748]++
											return fmt.Sprintf("%d", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:62
	// _ = "end of CoverTab[120748]"
}

func (s *uintSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:65
	_go_fuzz_dep_.CoverTab[120749]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:67
		_go_fuzz_dep_.CoverTab[120751]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:68
		// _ = "end of CoverTab[120751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:69
		_go_fuzz_dep_.CoverTab[120752]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:69
		// _ = "end of CoverTab[120752]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:69
	// _ = "end of CoverTab[120749]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:69
	_go_fuzz_dep_.CoverTab[120750]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:71
	// _ = "end of CoverTab[120750]"
}

func (s *uintSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:74
	_go_fuzz_dep_.CoverTab[120753]++
											out := make([]uint, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:76
		_go_fuzz_dep_.CoverTab[120755]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:79
			_go_fuzz_dep_.CoverTab[120756]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:80
			// _ = "end of CoverTab[120756]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:81
			_go_fuzz_dep_.CoverTab[120757]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:81
			// _ = "end of CoverTab[120757]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:81
		// _ = "end of CoverTab[120755]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:82
	// _ = "end of CoverTab[120753]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:82
	_go_fuzz_dep_.CoverTab[120754]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:84
	// _ = "end of CoverTab[120754]"
}

func (s *uintSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:87
	_go_fuzz_dep_.CoverTab[120758]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:89
		_go_fuzz_dep_.CoverTab[120760]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:90
		// _ = "end of CoverTab[120760]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:91
	// _ = "end of CoverTab[120758]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:91
	_go_fuzz_dep_.CoverTab[120759]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:92
	// _ = "end of CoverTab[120759]"
}

func uintSliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:95
	_go_fuzz_dep_.CoverTab[120761]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:98
		_go_fuzz_dep_.CoverTab[120764]++
												return []uint{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:99
		// _ = "end of CoverTab[120764]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:100
		_go_fuzz_dep_.CoverTab[120765]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:100
		// _ = "end of CoverTab[120765]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:100
	// _ = "end of CoverTab[120761]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:100
	_go_fuzz_dep_.CoverTab[120762]++
											ss := strings.Split(val, ",")
											out := make([]uint, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:103
		_go_fuzz_dep_.CoverTab[120766]++
												u, err := strconv.ParseUint(d, 10, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:105
			_go_fuzz_dep_.CoverTab[120768]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:106
			// _ = "end of CoverTab[120768]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:107
			_go_fuzz_dep_.CoverTab[120769]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:107
			// _ = "end of CoverTab[120769]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:107
		// _ = "end of CoverTab[120766]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:107
		_go_fuzz_dep_.CoverTab[120767]++
												out[i] = uint(u)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:108
		// _ = "end of CoverTab[120767]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:109
	// _ = "end of CoverTab[120762]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:109
	_go_fuzz_dep_.CoverTab[120763]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:110
	// _ = "end of CoverTab[120763]"
}

// GetUintSlice returns the []uint value of a flag with the given name.
func (f *FlagSet) GetUintSlice(name string) ([]uint, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:114
	_go_fuzz_dep_.CoverTab[120770]++
											val, err := f.getFlagType(name, "uintSlice", uintSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:116
		_go_fuzz_dep_.CoverTab[120772]++
												return []uint{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:117
		// _ = "end of CoverTab[120772]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:118
		_go_fuzz_dep_.CoverTab[120773]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:118
		// _ = "end of CoverTab[120773]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:118
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:118
	// _ = "end of CoverTab[120770]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:118
	_go_fuzz_dep_.CoverTab[120771]++
											return val.([]uint), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:119
	// _ = "end of CoverTab[120771]"
}

// UintSliceVar defines a uintSlice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:122
// The argument p points to a []uint variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:124
func (f *FlagSet) UintSliceVar(p *[]uint, name string, value []uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:124
	_go_fuzz_dep_.CoverTab[120774]++
											f.VarP(newUintSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:125
	// _ = "end of CoverTab[120774]"
}

// UintSliceVarP is like UintSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:129
	_go_fuzz_dep_.CoverTab[120775]++
											f.VarP(newUintSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:130
	// _ = "end of CoverTab[120775]"
}

// UintSliceVar defines a uint[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:133
// The argument p points to a uint[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:135
func UintSliceVar(p *[]uint, name string, value []uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:135
	_go_fuzz_dep_.CoverTab[120776]++
											CommandLine.VarP(newUintSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:136
	// _ = "end of CoverTab[120776]"
}

// UintSliceVarP is like the UintSliceVar, but accepts a shorthand letter that can be used after a single dash.
func UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:140
	_go_fuzz_dep_.CoverTab[120777]++
											CommandLine.VarP(newUintSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:141
	// _ = "end of CoverTab[120777]"
}

// UintSlice defines a []uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:144
// The return value is the address of a []uint variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:146
func (f *FlagSet) UintSlice(name string, value []uint, usage string) *[]uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:146
	_go_fuzz_dep_.CoverTab[120778]++
											p := []uint{}
											f.UintSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:149
	// _ = "end of CoverTab[120778]"
}

// UintSliceP is like UintSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintSliceP(name, shorthand string, value []uint, usage string) *[]uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:153
	_go_fuzz_dep_.CoverTab[120779]++
											p := []uint{}
											f.UintSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:156
	// _ = "end of CoverTab[120779]"
}

// UintSlice defines a []uint flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:159
// The return value is the address of a []uint variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:161
func UintSlice(name string, value []uint, usage string) *[]uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:161
	_go_fuzz_dep_.CoverTab[120780]++
											return CommandLine.UintSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:162
	// _ = "end of CoverTab[120780]"
}

// UintSliceP is like UintSlice, but accepts a shorthand letter that can be used after a single dash.
func UintSliceP(name, shorthand string, value []uint, usage string) *[]uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:166
	_go_fuzz_dep_.CoverTab[120781]++
											return CommandLine.UintSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:167
	// _ = "end of CoverTab[120781]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/uint_slice.go:168
var _ = _go_fuzz_dep_.CoverTab
