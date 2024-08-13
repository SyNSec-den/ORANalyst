//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:1
)

import (
	"fmt"
	"strconv"
	"strings"
)

// -- intSlice Value
type intSliceValue struct {
	value	*[]int
	changed	bool
}

func newIntSliceValue(val []int, p *[]int) *intSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:15
	_go_fuzz_dep_.CoverTab[120229]++
											isv := new(intSliceValue)
											isv.value = p
											*isv.value = val
											return isv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:19
	// _ = "end of CoverTab[120229]"
}

func (s *intSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:22
	_go_fuzz_dep_.CoverTab[120230]++
											ss := strings.Split(val, ",")
											out := make([]int, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:25
		_go_fuzz_dep_.CoverTab[120233]++
												var err error
												out[i], err = strconv.Atoi(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:28
			_go_fuzz_dep_.CoverTab[120234]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:29
			// _ = "end of CoverTab[120234]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:30
			_go_fuzz_dep_.CoverTab[120235]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:30
			// _ = "end of CoverTab[120235]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:30
		// _ = "end of CoverTab[120233]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:32
	// _ = "end of CoverTab[120230]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:32
	_go_fuzz_dep_.CoverTab[120231]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:33
		_go_fuzz_dep_.CoverTab[120236]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:34
		// _ = "end of CoverTab[120236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:35
		_go_fuzz_dep_.CoverTab[120237]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:36
		// _ = "end of CoverTab[120237]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:37
	// _ = "end of CoverTab[120231]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:37
	_go_fuzz_dep_.CoverTab[120232]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:39
	// _ = "end of CoverTab[120232]"
}

func (s *intSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:42
	_go_fuzz_dep_.CoverTab[120238]++
											return "intSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:43
	// _ = "end of CoverTab[120238]"
}

func (s *intSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:46
	_go_fuzz_dep_.CoverTab[120239]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:48
		_go_fuzz_dep_.CoverTab[120241]++
												out[i] = fmt.Sprintf("%d", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:49
		// _ = "end of CoverTab[120241]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:50
	// _ = "end of CoverTab[120239]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:50
	_go_fuzz_dep_.CoverTab[120240]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:51
	// _ = "end of CoverTab[120240]"
}

func (s *intSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:54
	_go_fuzz_dep_.CoverTab[120242]++
											i, err := strconv.Atoi(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:56
		_go_fuzz_dep_.CoverTab[120244]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:57
		// _ = "end of CoverTab[120244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:58
		_go_fuzz_dep_.CoverTab[120245]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:58
		// _ = "end of CoverTab[120245]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:58
	// _ = "end of CoverTab[120242]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:58
	_go_fuzz_dep_.CoverTab[120243]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:60
	// _ = "end of CoverTab[120243]"
}

func (s *intSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:63
	_go_fuzz_dep_.CoverTab[120246]++
											out := make([]int, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:65
		_go_fuzz_dep_.CoverTab[120248]++
												var err error
												out[i], err = strconv.Atoi(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:68
			_go_fuzz_dep_.CoverTab[120249]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:69
			// _ = "end of CoverTab[120249]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:70
			_go_fuzz_dep_.CoverTab[120250]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:70
			// _ = "end of CoverTab[120250]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:70
		// _ = "end of CoverTab[120248]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:71
	// _ = "end of CoverTab[120246]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:71
	_go_fuzz_dep_.CoverTab[120247]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:73
	// _ = "end of CoverTab[120247]"
}

func (s *intSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:76
	_go_fuzz_dep_.CoverTab[120251]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:78
		_go_fuzz_dep_.CoverTab[120253]++
												out[i] = strconv.Itoa(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:79
		// _ = "end of CoverTab[120253]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:80
	// _ = "end of CoverTab[120251]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:80
	_go_fuzz_dep_.CoverTab[120252]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:81
	// _ = "end of CoverTab[120252]"
}

func intSliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:84
	_go_fuzz_dep_.CoverTab[120254]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:87
		_go_fuzz_dep_.CoverTab[120257]++
												return []int{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:88
		// _ = "end of CoverTab[120257]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:89
		_go_fuzz_dep_.CoverTab[120258]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:89
		// _ = "end of CoverTab[120258]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:89
	// _ = "end of CoverTab[120254]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:89
	_go_fuzz_dep_.CoverTab[120255]++
											ss := strings.Split(val, ",")
											out := make([]int, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:92
		_go_fuzz_dep_.CoverTab[120259]++
												var err error
												out[i], err = strconv.Atoi(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:95
			_go_fuzz_dep_.CoverTab[120260]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:96
			// _ = "end of CoverTab[120260]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:97
			_go_fuzz_dep_.CoverTab[120261]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:97
			// _ = "end of CoverTab[120261]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:97
		// _ = "end of CoverTab[120259]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:99
	// _ = "end of CoverTab[120255]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:99
	_go_fuzz_dep_.CoverTab[120256]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:100
	// _ = "end of CoverTab[120256]"
}

// GetIntSlice return the []int value of a flag with the given name
func (f *FlagSet) GetIntSlice(name string) ([]int, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:104
	_go_fuzz_dep_.CoverTab[120262]++
											val, err := f.getFlagType(name, "intSlice", intSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:106
		_go_fuzz_dep_.CoverTab[120264]++
												return []int{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:107
		// _ = "end of CoverTab[120264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:108
		_go_fuzz_dep_.CoverTab[120265]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:108
		// _ = "end of CoverTab[120265]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:108
	// _ = "end of CoverTab[120262]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:108
	_go_fuzz_dep_.CoverTab[120263]++
											return val.([]int), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:109
	// _ = "end of CoverTab[120263]"
}

// IntSliceVar defines a intSlice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:112
// The argument p points to a []int variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:114
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:114
	_go_fuzz_dep_.CoverTab[120266]++
											f.VarP(newIntSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:115
	// _ = "end of CoverTab[120266]"
}

// IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:119
	_go_fuzz_dep_.CoverTab[120267]++
											f.VarP(newIntSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:120
	// _ = "end of CoverTab[120267]"
}

// IntSliceVar defines a int[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:123
// The argument p points to a int[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:125
func IntSliceVar(p *[]int, name string, value []int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:125
	_go_fuzz_dep_.CoverTab[120268]++
											CommandLine.VarP(newIntSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:126
	// _ = "end of CoverTab[120268]"
}

// IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.
func IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:130
	_go_fuzz_dep_.CoverTab[120269]++
											CommandLine.VarP(newIntSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:131
	// _ = "end of CoverTab[120269]"
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:134
// The return value is the address of a []int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:136
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:136
	_go_fuzz_dep_.CoverTab[120270]++
											p := []int{}
											f.IntSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:139
	// _ = "end of CoverTab[120270]"
}

// IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:143
	_go_fuzz_dep_.CoverTab[120271]++
											p := []int{}
											f.IntSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:146
	// _ = "end of CoverTab[120271]"
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:149
// The return value is the address of a []int variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:151
func IntSlice(name string, value []int, usage string) *[]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:151
	_go_fuzz_dep_.CoverTab[120272]++
											return CommandLine.IntSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:152
	// _ = "end of CoverTab[120272]"
}

// IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.
func IntSliceP(name, shorthand string, value []int, usage string) *[]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:156
	_go_fuzz_dep_.CoverTab[120273]++
											return CommandLine.IntSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:157
	// _ = "end of CoverTab[120273]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:158
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/int_slice.go:158
var _ = _go_fuzz_dep_.CoverTab
