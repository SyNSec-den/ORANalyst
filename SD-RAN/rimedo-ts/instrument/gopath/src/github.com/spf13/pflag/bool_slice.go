//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:1
)

import (
	"io"
	"strconv"
	"strings"
)

// -- boolSlice Value
type boolSliceValue struct {
	value	*[]bool
	changed	bool
}

func newBoolSliceValue(val []bool, p *[]bool) *boolSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:15
	_go_fuzz_dep_.CoverTab[119284]++
											bsv := new(boolSliceValue)
											bsv.value = p
											*bsv.value = val
											return bsv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:19
	// _ = "end of CoverTab[119284]"
}

// Set converts, and assigns, the comma-separated boolean argument string representation as the []bool value of this flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:22
// If Set is called on a flag that already has a []bool assigned, the newly converted values will be appended.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:24
func (s *boolSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:24
	_go_fuzz_dep_.CoverTab[119285]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:27
	rmQuote := strings.NewReplacer(`"`, "", `'`, "", "`", "")

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:30
	boolStrSlice, err := readAsCSV(rmQuote.Replace(val))
	if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:31
		_go_fuzz_dep_.CoverTab[119289]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:31
		return err != io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:31
		// _ = "end of CoverTab[119289]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:31
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:31
		_go_fuzz_dep_.CoverTab[119290]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:32
		// _ = "end of CoverTab[119290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:33
		_go_fuzz_dep_.CoverTab[119291]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:33
		// _ = "end of CoverTab[119291]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:33
	// _ = "end of CoverTab[119285]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:33
	_go_fuzz_dep_.CoverTab[119286]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:36
	out := make([]bool, 0, len(boolStrSlice))
	for _, boolStr := range boolStrSlice {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:37
		_go_fuzz_dep_.CoverTab[119292]++
												b, err := strconv.ParseBool(strings.TrimSpace(boolStr))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:39
			_go_fuzz_dep_.CoverTab[119294]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:40
			// _ = "end of CoverTab[119294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:41
			_go_fuzz_dep_.CoverTab[119295]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:41
			// _ = "end of CoverTab[119295]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:41
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:41
		// _ = "end of CoverTab[119292]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:41
		_go_fuzz_dep_.CoverTab[119293]++
												out = append(out, b)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:42
		// _ = "end of CoverTab[119293]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:43
	// _ = "end of CoverTab[119286]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:43
	_go_fuzz_dep_.CoverTab[119287]++

											if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:45
		_go_fuzz_dep_.CoverTab[119296]++
												*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:46
		// _ = "end of CoverTab[119296]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:47
		_go_fuzz_dep_.CoverTab[119297]++
												*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:48
		// _ = "end of CoverTab[119297]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:49
	// _ = "end of CoverTab[119287]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:49
	_go_fuzz_dep_.CoverTab[119288]++

											s.changed = true

											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:53
	// _ = "end of CoverTab[119288]"
}

// Type returns a string that uniquely represents this flag's type.
func (s *boolSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:57
	_go_fuzz_dep_.CoverTab[119298]++
											return "boolSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:58
	// _ = "end of CoverTab[119298]"
}

// String defines a "native" format for this boolean slice flag value.
func (s *boolSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:62
	_go_fuzz_dep_.CoverTab[119299]++

											boolStrSlice := make([]string, len(*s.value))
											for i, b := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:65
		_go_fuzz_dep_.CoverTab[119301]++
												boolStrSlice[i] = strconv.FormatBool(b)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:66
		// _ = "end of CoverTab[119301]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:67
	// _ = "end of CoverTab[119299]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:67
	_go_fuzz_dep_.CoverTab[119300]++

											out, _ := writeAsCSV(boolStrSlice)

											return "[" + out + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:71
	// _ = "end of CoverTab[119300]"
}

func (s *boolSliceValue) fromString(val string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:74
	_go_fuzz_dep_.CoverTab[119302]++
											return strconv.ParseBool(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:75
	// _ = "end of CoverTab[119302]"
}

func (s *boolSliceValue) toString(val bool) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:78
	_go_fuzz_dep_.CoverTab[119303]++
											return strconv.FormatBool(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:79
	// _ = "end of CoverTab[119303]"
}

func (s *boolSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:82
	_go_fuzz_dep_.CoverTab[119304]++
											i, err := s.fromString(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:84
		_go_fuzz_dep_.CoverTab[119306]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:85
		// _ = "end of CoverTab[119306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:86
		_go_fuzz_dep_.CoverTab[119307]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:86
		// _ = "end of CoverTab[119307]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:86
	// _ = "end of CoverTab[119304]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:86
	_go_fuzz_dep_.CoverTab[119305]++
											*s.value = append(*s.value, i)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:88
	// _ = "end of CoverTab[119305]"
}

func (s *boolSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:91
	_go_fuzz_dep_.CoverTab[119308]++
											out := make([]bool, len(val))
											for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:93
		_go_fuzz_dep_.CoverTab[119310]++
												var err error
												out[i], err = s.fromString(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:96
			_go_fuzz_dep_.CoverTab[119311]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:97
			// _ = "end of CoverTab[119311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:98
			_go_fuzz_dep_.CoverTab[119312]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:98
			// _ = "end of CoverTab[119312]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:98
		// _ = "end of CoverTab[119310]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:99
	// _ = "end of CoverTab[119308]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:99
	_go_fuzz_dep_.CoverTab[119309]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:101
	// _ = "end of CoverTab[119309]"
}

func (s *boolSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:104
	_go_fuzz_dep_.CoverTab[119313]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:106
		_go_fuzz_dep_.CoverTab[119315]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:107
		// _ = "end of CoverTab[119315]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:108
	// _ = "end of CoverTab[119313]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:108
	_go_fuzz_dep_.CoverTab[119314]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:109
	// _ = "end of CoverTab[119314]"
}

func boolSliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:112
	_go_fuzz_dep_.CoverTab[119316]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:115
		_go_fuzz_dep_.CoverTab[119319]++
												return []bool{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:116
		// _ = "end of CoverTab[119319]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:117
		_go_fuzz_dep_.CoverTab[119320]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:117
		// _ = "end of CoverTab[119320]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:117
	// _ = "end of CoverTab[119316]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:117
	_go_fuzz_dep_.CoverTab[119317]++
											ss := strings.Split(val, ",")
											out := make([]bool, len(ss))
											for i, t := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:120
		_go_fuzz_dep_.CoverTab[119321]++
												var err error
												out[i], err = strconv.ParseBool(t)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:123
			_go_fuzz_dep_.CoverTab[119322]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:124
			// _ = "end of CoverTab[119322]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:125
			_go_fuzz_dep_.CoverTab[119323]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:125
			// _ = "end of CoverTab[119323]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:125
		// _ = "end of CoverTab[119321]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:126
	// _ = "end of CoverTab[119317]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:126
	_go_fuzz_dep_.CoverTab[119318]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:127
	// _ = "end of CoverTab[119318]"
}

// GetBoolSlice returns the []bool value of a flag with the given name.
func (f *FlagSet) GetBoolSlice(name string) ([]bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:131
	_go_fuzz_dep_.CoverTab[119324]++
											val, err := f.getFlagType(name, "boolSlice", boolSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:133
		_go_fuzz_dep_.CoverTab[119326]++
												return []bool{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:134
		// _ = "end of CoverTab[119326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:135
		_go_fuzz_dep_.CoverTab[119327]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:135
		// _ = "end of CoverTab[119327]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:135
	// _ = "end of CoverTab[119324]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:135
	_go_fuzz_dep_.CoverTab[119325]++
											return val.([]bool), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:136
	// _ = "end of CoverTab[119325]"
}

// BoolSliceVar defines a boolSlice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:139
// The argument p points to a []bool variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:141
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:141
	_go_fuzz_dep_.CoverTab[119328]++
											f.VarP(newBoolSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:142
	// _ = "end of CoverTab[119328]"
}

// BoolSliceVarP is like BoolSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:146
	_go_fuzz_dep_.CoverTab[119329]++
											f.VarP(newBoolSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:147
	// _ = "end of CoverTab[119329]"
}

// BoolSliceVar defines a []bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:150
// The argument p points to a []bool variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:152
func BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:152
	_go_fuzz_dep_.CoverTab[119330]++
											CommandLine.VarP(newBoolSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:153
	// _ = "end of CoverTab[119330]"
}

// BoolSliceVarP is like BoolSliceVar, but accepts a shorthand letter that can be used after a single dash.
func BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:157
	_go_fuzz_dep_.CoverTab[119331]++
											CommandLine.VarP(newBoolSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:158
	// _ = "end of CoverTab[119331]"
}

// BoolSlice defines a []bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:161
// The return value is the address of a []bool variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:163
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:163
	_go_fuzz_dep_.CoverTab[119332]++
											p := []bool{}
											f.BoolSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:166
	// _ = "end of CoverTab[119332]"
}

// BoolSliceP is like BoolSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:170
	_go_fuzz_dep_.CoverTab[119333]++
											p := []bool{}
											f.BoolSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:173
	// _ = "end of CoverTab[119333]"
}

// BoolSlice defines a []bool flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:176
// The return value is the address of a []bool variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:178
func BoolSlice(name string, value []bool, usage string) *[]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:178
	_go_fuzz_dep_.CoverTab[119334]++
											return CommandLine.BoolSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:179
	// _ = "end of CoverTab[119334]"
}

// BoolSliceP is like BoolSlice, but accepts a shorthand letter that can be used after a single dash.
func BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:183
	_go_fuzz_dep_.CoverTab[119335]++
											return CommandLine.BoolSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:184
	// _ = "end of CoverTab[119335]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:185
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/bool_slice.go:185
var _ = _go_fuzz_dep_.CoverTab
