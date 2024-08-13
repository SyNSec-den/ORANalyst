//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:1
)

import (
	"fmt"
	"net"
	"strconv"
)

// -- net.IPMask value
type ipMaskValue net.IPMask

func newIPMaskValue(val net.IPMask, p *net.IPMask) *ipMaskValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:12
	_go_fuzz_dep_.CoverTab[120350]++
										*p = val
										return (*ipMaskValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:14
	// _ = "end of CoverTab[120350]"
}

func (i *ipMaskValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:17
	_go_fuzz_dep_.CoverTab[120351]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:17
	return net.IPMask(*i).String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:17
	// _ = "end of CoverTab[120351]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:17
}
func (i *ipMaskValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:18
	_go_fuzz_dep_.CoverTab[120352]++
										ip := ParseIPv4Mask(s)
										if ip == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:20
		_go_fuzz_dep_.CoverTab[120354]++
											return fmt.Errorf("failed to parse IP mask: %q", s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:21
		// _ = "end of CoverTab[120354]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:22
		_go_fuzz_dep_.CoverTab[120355]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:22
		// _ = "end of CoverTab[120355]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:22
	// _ = "end of CoverTab[120352]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:22
	_go_fuzz_dep_.CoverTab[120353]++
										*i = ipMaskValue(ip)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:24
	// _ = "end of CoverTab[120353]"
}

func (i *ipMaskValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:27
	_go_fuzz_dep_.CoverTab[120356]++
										return "ipMask"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:28
	// _ = "end of CoverTab[120356]"
}

// ParseIPv4Mask written in IP form (e.g. 255.255.255.0).
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:31
// This function should really belong to the net package.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:33
func ParseIPv4Mask(s string) net.IPMask {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:33
	_go_fuzz_dep_.CoverTab[120357]++
										mask := net.ParseIP(s)
										if mask == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:35
		_go_fuzz_dep_.CoverTab[120359]++
											if len(s) != 8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:36
			_go_fuzz_dep_.CoverTab[120362]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:37
			// _ = "end of CoverTab[120362]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:38
			_go_fuzz_dep_.CoverTab[120363]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:38
			// _ = "end of CoverTab[120363]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:38
		// _ = "end of CoverTab[120359]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:38
		_go_fuzz_dep_.CoverTab[120360]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:41
		m := []int{}
		for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:42
			_go_fuzz_dep_.CoverTab[120364]++
												b := "0x" + s[2*i:2*i+2]
												d, err := strconv.ParseInt(b, 0, 0)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:45
				_go_fuzz_dep_.CoverTab[120366]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:46
				// _ = "end of CoverTab[120366]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:47
				_go_fuzz_dep_.CoverTab[120367]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:47
				// _ = "end of CoverTab[120367]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:47
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:47
			// _ = "end of CoverTab[120364]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:47
			_go_fuzz_dep_.CoverTab[120365]++
												m = append(m, int(d))
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:48
			// _ = "end of CoverTab[120365]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:49
		// _ = "end of CoverTab[120360]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:49
		_go_fuzz_dep_.CoverTab[120361]++
											s := fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
											mask = net.ParseIP(s)
											if mask == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:52
			_go_fuzz_dep_.CoverTab[120368]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:53
			// _ = "end of CoverTab[120368]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:54
			_go_fuzz_dep_.CoverTab[120369]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:54
			// _ = "end of CoverTab[120369]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:54
		// _ = "end of CoverTab[120361]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:55
		_go_fuzz_dep_.CoverTab[120370]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:55
		// _ = "end of CoverTab[120370]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:55
	// _ = "end of CoverTab[120357]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:55
	_go_fuzz_dep_.CoverTab[120358]++
										return net.IPv4Mask(mask[12], mask[13], mask[14], mask[15])
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:56
	// _ = "end of CoverTab[120358]"
}

func parseIPv4Mask(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:59
	_go_fuzz_dep_.CoverTab[120371]++
										mask := ParseIPv4Mask(sval)
										if mask == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:61
		_go_fuzz_dep_.CoverTab[120373]++
											return nil, fmt.Errorf("unable to parse %s as net.IPMask", sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:62
		// _ = "end of CoverTab[120373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:63
		_go_fuzz_dep_.CoverTab[120374]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:63
		// _ = "end of CoverTab[120374]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:63
	// _ = "end of CoverTab[120371]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:63
	_go_fuzz_dep_.CoverTab[120372]++
										return mask, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:64
	// _ = "end of CoverTab[120372]"
}

// GetIPv4Mask return the net.IPv4Mask value of a flag with the given name
func (f *FlagSet) GetIPv4Mask(name string) (net.IPMask, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:68
	_go_fuzz_dep_.CoverTab[120375]++
										val, err := f.getFlagType(name, "ipMask", parseIPv4Mask)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:70
		_go_fuzz_dep_.CoverTab[120377]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:71
		// _ = "end of CoverTab[120377]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:72
		_go_fuzz_dep_.CoverTab[120378]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:72
		// _ = "end of CoverTab[120378]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:72
	// _ = "end of CoverTab[120375]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:72
	_go_fuzz_dep_.CoverTab[120376]++
										return val.(net.IPMask), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:73
	// _ = "end of CoverTab[120376]"
}

// IPMaskVar defines an net.IPMask flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:76
// The argument p points to an net.IPMask variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:78
func (f *FlagSet) IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:78
	_go_fuzz_dep_.CoverTab[120379]++
										f.VarP(newIPMaskValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:79
	// _ = "end of CoverTab[120379]"
}

// IPMaskVarP is like IPMaskVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:83
	_go_fuzz_dep_.CoverTab[120380]++
										f.VarP(newIPMaskValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:84
	// _ = "end of CoverTab[120380]"
}

// IPMaskVar defines an net.IPMask flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:87
// The argument p points to an net.IPMask variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:89
func IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:89
	_go_fuzz_dep_.CoverTab[120381]++
										CommandLine.VarP(newIPMaskValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:90
	// _ = "end of CoverTab[120381]"
}

// IPMaskVarP is like IPMaskVar, but accepts a shorthand letter that can be used after a single dash.
func IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:94
	_go_fuzz_dep_.CoverTab[120382]++
										CommandLine.VarP(newIPMaskValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:95
	// _ = "end of CoverTab[120382]"
}

// IPMask defines an net.IPMask flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:98
// The return value is the address of an net.IPMask variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:100
func (f *FlagSet) IPMask(name string, value net.IPMask, usage string) *net.IPMask {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:100
	_go_fuzz_dep_.CoverTab[120383]++
										p := new(net.IPMask)
										f.IPMaskVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:103
	// _ = "end of CoverTab[120383]"
}

// IPMaskP is like IPMask, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:107
	_go_fuzz_dep_.CoverTab[120384]++
										p := new(net.IPMask)
										f.IPMaskVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:110
	// _ = "end of CoverTab[120384]"
}

// IPMask defines an net.IPMask flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:113
// The return value is the address of an net.IPMask variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:115
func IPMask(name string, value net.IPMask, usage string) *net.IPMask {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:115
	_go_fuzz_dep_.CoverTab[120385]++
										return CommandLine.IPMaskP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:116
	// _ = "end of CoverTab[120385]"
}

// IPMaskP is like IP, but accepts a shorthand letter that can be used after a single dash.
func IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:120
	_go_fuzz_dep_.CoverTab[120386]++
										return CommandLine.IPMaskP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:121
	// _ = "end of CoverTab[120386]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipmask.go:122
var _ = _go_fuzz_dep_.CoverTab
