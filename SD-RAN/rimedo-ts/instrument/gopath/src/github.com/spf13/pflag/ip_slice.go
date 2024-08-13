//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:1
)

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// -- ipSlice Value
type ipSliceValue struct {
	value	*[]net.IP
	changed	bool
}

func newIPSliceValue(val []net.IP, p *[]net.IP) *ipSliceValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:16
	_go_fuzz_dep_.CoverTab[120297]++
										ipsv := new(ipSliceValue)
										ipsv.value = p
										*ipsv.value = val
										return ipsv
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:20
	// _ = "end of CoverTab[120297]"
}

// Set converts, and assigns, the comma-separated IP argument string representation as the []net.IP value of this flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:23
// If Set is called on a flag that already has a []net.IP assigned, the newly converted values will be appended.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:25
func (s *ipSliceValue) Set(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:25
	_go_fuzz_dep_.CoverTab[120298]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:28
	rmQuote := strings.NewReplacer(`"`, "", `'`, "", "`", "")

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:31
	ipStrSlice, err := readAsCSV(rmQuote.Replace(val))
	if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:32
		_go_fuzz_dep_.CoverTab[120302]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:32
		return err != io.EOF
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:32
		// _ = "end of CoverTab[120302]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:32
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:32
		_go_fuzz_dep_.CoverTab[120303]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:33
		// _ = "end of CoverTab[120303]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:34
		_go_fuzz_dep_.CoverTab[120304]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:34
		// _ = "end of CoverTab[120304]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:34
	// _ = "end of CoverTab[120298]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:34
	_go_fuzz_dep_.CoverTab[120299]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:37
	out := make([]net.IP, 0, len(ipStrSlice))
	for _, ipStr := range ipStrSlice {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:38
		_go_fuzz_dep_.CoverTab[120305]++
											ip := net.ParseIP(strings.TrimSpace(ipStr))
											if ip == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:40
			_go_fuzz_dep_.CoverTab[120307]++
												return fmt.Errorf("invalid string being converted to IP address: %s", ipStr)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:41
			// _ = "end of CoverTab[120307]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:42
			_go_fuzz_dep_.CoverTab[120308]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:42
			// _ = "end of CoverTab[120308]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:42
		// _ = "end of CoverTab[120305]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:42
		_go_fuzz_dep_.CoverTab[120306]++
											out = append(out, ip)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:43
		// _ = "end of CoverTab[120306]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:44
	// _ = "end of CoverTab[120299]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:44
	_go_fuzz_dep_.CoverTab[120300]++

										if !s.changed {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:46
		_go_fuzz_dep_.CoverTab[120309]++
											*s.value = out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:47
		// _ = "end of CoverTab[120309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:48
		_go_fuzz_dep_.CoverTab[120310]++
											*s.value = append(*s.value, out...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:49
		// _ = "end of CoverTab[120310]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:50
	// _ = "end of CoverTab[120300]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:50
	_go_fuzz_dep_.CoverTab[120301]++

										s.changed = true

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:54
	// _ = "end of CoverTab[120301]"
}

// Type returns a string that uniquely represents this flag's type.
func (s *ipSliceValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:58
	_go_fuzz_dep_.CoverTab[120311]++
										return "ipSlice"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:59
	// _ = "end of CoverTab[120311]"
}

// String defines a "native" format for this net.IP slice flag value.
func (s *ipSliceValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:63
	_go_fuzz_dep_.CoverTab[120312]++

										ipStrSlice := make([]string, len(*s.value))
										for i, ip := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:66
		_go_fuzz_dep_.CoverTab[120314]++
											ipStrSlice[i] = ip.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:67
		// _ = "end of CoverTab[120314]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:68
	// _ = "end of CoverTab[120312]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:68
	_go_fuzz_dep_.CoverTab[120313]++

										out, _ := writeAsCSV(ipStrSlice)

										return "[" + out + "]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:72
	// _ = "end of CoverTab[120313]"
}

func (s *ipSliceValue) fromString(val string) (net.IP, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:75
	_go_fuzz_dep_.CoverTab[120315]++
										return net.ParseIP(strings.TrimSpace(val)), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:76
	// _ = "end of CoverTab[120315]"
}

func (s *ipSliceValue) toString(val net.IP) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:79
	_go_fuzz_dep_.CoverTab[120316]++
										return val.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:80
	// _ = "end of CoverTab[120316]"
}

func (s *ipSliceValue) Append(val string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:83
	_go_fuzz_dep_.CoverTab[120317]++
										i, err := s.fromString(val)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:85
		_go_fuzz_dep_.CoverTab[120319]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:86
		// _ = "end of CoverTab[120319]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:87
		_go_fuzz_dep_.CoverTab[120320]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:87
		// _ = "end of CoverTab[120320]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:87
	// _ = "end of CoverTab[120317]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:87
	_go_fuzz_dep_.CoverTab[120318]++
										*s.value = append(*s.value, i)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:89
	// _ = "end of CoverTab[120318]"
}

func (s *ipSliceValue) Replace(val []string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:92
	_go_fuzz_dep_.CoverTab[120321]++
										out := make([]net.IP, len(val))
										for i, d := range val {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:94
		_go_fuzz_dep_.CoverTab[120323]++
											var err error
											out[i], err = s.fromString(d)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:97
			_go_fuzz_dep_.CoverTab[120324]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:98
			// _ = "end of CoverTab[120324]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:99
			_go_fuzz_dep_.CoverTab[120325]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:99
			// _ = "end of CoverTab[120325]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:99
		// _ = "end of CoverTab[120323]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:100
	// _ = "end of CoverTab[120321]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:100
	_go_fuzz_dep_.CoverTab[120322]++
											*s.value = out
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:102
	// _ = "end of CoverTab[120322]"
}

func (s *ipSliceValue) GetSlice() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:105
	_go_fuzz_dep_.CoverTab[120326]++
											out := make([]string, len(*s.value))
											for i, d := range *s.value {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:107
		_go_fuzz_dep_.CoverTab[120328]++
												out[i] = s.toString(d)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:108
		// _ = "end of CoverTab[120328]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:109
	// _ = "end of CoverTab[120326]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:109
	_go_fuzz_dep_.CoverTab[120327]++
											return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:110
	// _ = "end of CoverTab[120327]"
}

func ipSliceConv(val string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:113
	_go_fuzz_dep_.CoverTab[120329]++
											val = strings.Trim(val, "[]")

											if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:116
		_go_fuzz_dep_.CoverTab[120332]++
												return []net.IP{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:117
		// _ = "end of CoverTab[120332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:118
		_go_fuzz_dep_.CoverTab[120333]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:118
		// _ = "end of CoverTab[120333]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:118
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:118
	// _ = "end of CoverTab[120329]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:118
	_go_fuzz_dep_.CoverTab[120330]++
											ss := strings.Split(val, ",")
											out := make([]net.IP, len(ss))
											for i, sval := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:121
		_go_fuzz_dep_.CoverTab[120334]++
												ip := net.ParseIP(strings.TrimSpace(sval))
												if ip == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:123
			_go_fuzz_dep_.CoverTab[120336]++
													return nil, fmt.Errorf("invalid string being converted to IP address: %s", sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:124
			// _ = "end of CoverTab[120336]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:125
			_go_fuzz_dep_.CoverTab[120337]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:125
			// _ = "end of CoverTab[120337]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:125
		// _ = "end of CoverTab[120334]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:125
		_go_fuzz_dep_.CoverTab[120335]++
												out[i] = ip
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:126
		// _ = "end of CoverTab[120335]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:127
	// _ = "end of CoverTab[120330]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:127
	_go_fuzz_dep_.CoverTab[120331]++
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:128
	// _ = "end of CoverTab[120331]"
}

// GetIPSlice returns the []net.IP value of a flag with the given name
func (f *FlagSet) GetIPSlice(name string) ([]net.IP, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:132
	_go_fuzz_dep_.CoverTab[120338]++
											val, err := f.getFlagType(name, "ipSlice", ipSliceConv)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:134
		_go_fuzz_dep_.CoverTab[120340]++
												return []net.IP{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:135
		// _ = "end of CoverTab[120340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:136
		_go_fuzz_dep_.CoverTab[120341]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:136
		// _ = "end of CoverTab[120341]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:136
	// _ = "end of CoverTab[120338]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:136
	_go_fuzz_dep_.CoverTab[120339]++
											return val.([]net.IP), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:137
	// _ = "end of CoverTab[120339]"
}

// IPSliceVar defines a ipSlice flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:140
// The argument p points to a []net.IP variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:142
func (f *FlagSet) IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:142
	_go_fuzz_dep_.CoverTab[120342]++
											f.VarP(newIPSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:143
	// _ = "end of CoverTab[120342]"
}

// IPSliceVarP is like IPSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:147
	_go_fuzz_dep_.CoverTab[120343]++
											f.VarP(newIPSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:148
	// _ = "end of CoverTab[120343]"
}

// IPSliceVar defines a []net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:151
// The argument p points to a []net.IP variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:153
func IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:153
	_go_fuzz_dep_.CoverTab[120344]++
											CommandLine.VarP(newIPSliceValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:154
	// _ = "end of CoverTab[120344]"
}

// IPSliceVarP is like IPSliceVar, but accepts a shorthand letter that can be used after a single dash.
func IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:158
	_go_fuzz_dep_.CoverTab[120345]++
											CommandLine.VarP(newIPSliceValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:159
	// _ = "end of CoverTab[120345]"
}

// IPSlice defines a []net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:162
// The return value is the address of a []net.IP variable that stores the value of that flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:164
func (f *FlagSet) IPSlice(name string, value []net.IP, usage string) *[]net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:164
	_go_fuzz_dep_.CoverTab[120346]++
											p := []net.IP{}
											f.IPSliceVarP(&p, name, "", value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:167
	// _ = "end of CoverTab[120346]"
}

// IPSliceP is like IPSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:171
	_go_fuzz_dep_.CoverTab[120347]++
											p := []net.IP{}
											f.IPSliceVarP(&p, name, shorthand, value, usage)
											return &p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:174
	// _ = "end of CoverTab[120347]"
}

// IPSlice defines a []net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:177
// The return value is the address of a []net.IP variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:179
func IPSlice(name string, value []net.IP, usage string) *[]net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:179
	_go_fuzz_dep_.CoverTab[120348]++
											return CommandLine.IPSliceP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:180
	// _ = "end of CoverTab[120348]"
}

// IPSliceP is like IPSlice, but accepts a shorthand letter that can be used after a single dash.
func IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:184
	_go_fuzz_dep_.CoverTab[120349]++
											return CommandLine.IPSliceP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:185
	// _ = "end of CoverTab[120349]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:186
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip_slice.go:186
var _ = _go_fuzz_dep_.CoverTab
