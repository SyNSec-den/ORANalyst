//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:1
)

import (
	"fmt"
	"strings"
	"time"
)

// -- durationSlice Value
type durationSliceValue struct {
	value	*[]time.Duration
	changed	bool
}

func newDurationSliceValue(val []time.Duration, p *[]time.Duration) *durationSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:15
	_go_fuzz_dep_.CoverTab[119422]++
											dsv := new(durationSliceValue)
											dsv.value = p
											*dsv.value = val
											return dsv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:19
	// _ = "end of CoverTab[119422]"
}

func (s *durationSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:22
	_go_fuzz_dep_.CoverTab[119423]++
											ss := strings.Split(val, ",")
											out := make([]time.Duration, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:25
		_go_fuzz_dep_.CoverTab[119426]++
												var err error
												out[i], err = time.ParseDuration(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:28
			_go_fuzz_dep_.CoverTab[119427]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:29
			// _ = "end of CoverTab[119427]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:30
			_go_fuzz_dep_.CoverTab[119428]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:30
			// _ = "end of CoverTab[119428]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:30
		// _ = "end of CoverTab[119426]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:32
	// _ = "end of CoverTab[119423]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:32
	_go_fuzz_dep_.CoverTab[119424]++
											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:33
		_go_fuzz_dep_.CoverTab[119429]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:34
		// _ = "end of CoverTab[119429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:35
		_go_fuzz_dep_.CoverTab[119430]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:36
		// _ = "end of CoverTab[119430]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:37
	// _ = "end of CoverTab[119424]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:37
	_go_fuzz_dep_.CoverTab[119425]++
											s.changed = true
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:39
	// _ = "end of CoverTab[119425]"
}

func (s *durationSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:42
	_go_fuzz_dep_.CoverTab[119431]++
											return "durationSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:43
	// _ = "end of CoverTab[119431]"
}

func (s *durationSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:46
	_go_fuzz_dep_.CoverTab[119432]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:48
		_go_fuzz_dep_.CoverTab[119434]++
												out[i] = fmt.Sprintf("%s", d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:49
		// _ = "end of CoverTab[119434]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:50
	// _ = "end of CoverTab[119432]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:50
	_go_fuzz_dep_.CoverTab[119433]++
											return "[" + strings.Join(out, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:51
	// _ = "end of CoverTab[119433]"
}

func (s *durationSliceValue) fromString(val string) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:54
	_go_fuzz_dep_.CoverTab[119435]++
											return time.ParseDuration(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:55
	// _ = "end of CoverTab[119435]"
}

func (s *durationSliceValue) toString(val time.Duration) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:58
	_go_fuzz_dep_.CoverTab[119436]++
											return fmt.Sprintf("%s", val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:59
	// _ = "end of CoverTab[119436]"
}

func (s *durationSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:62
	_go_fuzz_dep_.CoverTab[119437]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:64
		_go_fuzz_dep_.CoverTab[119439]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:65
		// _ = "end of CoverTab[119439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:66
		_go_fuzz_dep_.CoverTab[119440]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:66
		// _ = "end of CoverTab[119440]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:66
	// _ = "end of CoverTab[119437]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:66
	_go_fuzz_dep_.CoverTab[119438]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:68
	// _ = "end of CoverTab[119438]"
}

func (s *durationSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:71
	_go_fuzz_dep_.CoverTab[119441]++
											out := make([]time.Duration, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:73
		_go_fuzz_dep_.CoverTab[119443]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:76
			_go_fuzz_dep_.CoverTab[119444]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:77
			// _ = "end of CoverTab[119444]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:78
			_go_fuzz_dep_.CoverTab[119445]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:78
			// _ = "end of CoverTab[119445]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:78
		// _ = "end of CoverTab[119443]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:79
	// _ = "end of CoverTab[119441]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:79
	_go_fuzz_dep_.CoverTab[119442]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:81
	// _ = "end of CoverTab[119442]"
}

func (s *durationSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:84
	_go_fuzz_dep_.CoverTab[119446]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:86
		_go_fuzz_dep_.CoverTab[119448]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:87
		// _ = "end of CoverTab[119448]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:88
	// _ = "end of CoverTab[119446]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:88
	_go_fuzz_dep_.CoverTab[119447]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:89
	// _ = "end of CoverTab[119447]"
}

func durationSliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:92
	_go_fuzz_dep_.CoverTab[119449]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:95
		_go_fuzz_dep_.CoverTab[119452]++
												return []time.Duration{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:96
		// _ = "end of CoverTab[119452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:97
		_go_fuzz_dep_.CoverTab[119453]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:97
		// _ = "end of CoverTab[119453]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:97
	// _ = "end of CoverTab[119449]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:97
	_go_fuzz_dep_.CoverTab[119450]++
											ss := strings.Split(val, ",")
											out := make([]time.Duration, len(ss))
											for i, d := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:100
		_go_fuzz_dep_.CoverTab[119454]++
												var err error
												out[i], err = time.ParseDuration(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:103
			_go_fuzz_dep_.CoverTab[119455]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:104
			// _ = "end of CoverTab[119455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:105
			_go_fuzz_dep_.CoverTab[119456]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:105
			// _ = "end of CoverTab[119456]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:105
		// _ = "end of CoverTab[119454]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:107
	// _ = "end of CoverTab[119450]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:107
	_go_fuzz_dep_.CoverTab[119451]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:108
	// _ = "end of CoverTab[119451]"
}

// GetDurationSlice returns the []time.Duration value of a flag with the given name
func (f *FlagSet) GetDurationSlice(name string) ([]time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:112
	_go_fuzz_dep_.CoverTab[119457]++
											val, err := f.getFlagType(name, "durationSlice", durationSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:114
		_go_fuzz_dep_.CoverTab[119459]++
												return []time.Duration{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:115
		// _ = "end of CoverTab[119459]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:116
		_go_fuzz_dep_.CoverTab[119460]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:116
		// _ = "end of CoverTab[119460]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:116
	// _ = "end of CoverTab[119457]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:116
	_go_fuzz_dep_.CoverTab[119458]++
											return val.([]time.Duration), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:117
	// _ = "end of CoverTab[119458]"
}

// DurationSliceVar defines a durationSlice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:120
// The argument p points to a []time.Duration variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:122
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:122
	_go_fuzz_dep_.CoverTab[119461]++
											f.VarP(newDurationSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:123
	// _ = "end of CoverTab[119461]"
}

// DurationSliceVarP is like DurationSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:127
	_go_fuzz_dep_.CoverTab[119462]++
											f.VarP(newDurationSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:128
	// _ = "end of CoverTab[119462]"
}

// DurationSliceVar defines a duration[] flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:131
// The argument p points to a duration[] variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:133
func DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:133
	_go_fuzz_dep_.CoverTab[119463]++
											CommandLine.VarP(newDurationSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:134
	// _ = "end of CoverTab[119463]"
}

// DurationSliceVarP is like DurationSliceVar, but accepts a shorthand letter that can be used after a single dash.
func DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:138
	_go_fuzz_dep_.CoverTab[119464]++
											CommandLine.VarP(newDurationSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:139
	// _ = "end of CoverTab[119464]"
}

// DurationSlice defines a []time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:142
// The return value is the address of a []time.Duration variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:144
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:144
	_go_fuzz_dep_.CoverTab[119465]++
											p := []time.Duration{}
											f.DurationSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:147
	// _ = "end of CoverTab[119465]"
}

// DurationSliceP is like DurationSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:151
	_go_fuzz_dep_.CoverTab[119466]++
											p := []time.Duration{}
											f.DurationSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:154
	// _ = "end of CoverTab[119466]"
}

// DurationSlice defines a []time.Duration flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:157
// The return value is the address of a []time.Duration variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:159
func DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:159
	_go_fuzz_dep_.CoverTab[119467]++
											return CommandLine.DurationSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:160
	// _ = "end of CoverTab[119467]"
}

// DurationSliceP is like DurationSlice, but accepts a shorthand letter that can be used after a single dash.
func DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:164
	_go_fuzz_dep_.CoverTab[119468]++
											return CommandLine.DurationSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:165
	// _ = "end of CoverTab[119468]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:166
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/duration_slice.go:166
var _ = _go_fuzz_dep_.CoverTab
