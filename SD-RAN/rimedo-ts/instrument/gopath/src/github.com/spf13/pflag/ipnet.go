//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:1
)

import (
	"fmt"
	"net"
	"strings"
)

// IPNet adapts net.IPNet for use as a flag.
type ipNetValue net.IPNet

func (ipnet ipNetValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:12
	_go_fuzz_dep_.CoverTab[120387]++
										n := net.IPNet(ipnet)
										return n.String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:14
	// _ = "end of CoverTab[120387]"
}

func (ipnet *ipNetValue) Set(value string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:17
	_go_fuzz_dep_.CoverTab[120388]++
										_, n, err := net.ParseCIDR(strings.TrimSpace(value))
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:19
		_go_fuzz_dep_.CoverTab[120390]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:20
		// _ = "end of CoverTab[120390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:21
		_go_fuzz_dep_.CoverTab[120391]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:21
		// _ = "end of CoverTab[120391]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:21
	// _ = "end of CoverTab[120388]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:21
	_go_fuzz_dep_.CoverTab[120389]++
										*ipnet = ipNetValue(*n)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:23
	// _ = "end of CoverTab[120389]"
}

func (*ipNetValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:26
	_go_fuzz_dep_.CoverTab[120392]++
										return "ipNet"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:27
	// _ = "end of CoverTab[120392]"
}

func newIPNetValue(val net.IPNet, p *net.IPNet) *ipNetValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:30
	_go_fuzz_dep_.CoverTab[120393]++
										*p = val
										return (*ipNetValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:32
	// _ = "end of CoverTab[120393]"
}

func ipNetConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:35
	_go_fuzz_dep_.CoverTab[120394]++
										_, n, err := net.ParseCIDR(strings.TrimSpace(sval))
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:37
		_go_fuzz_dep_.CoverTab[120396]++
											return *n, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:38
		// _ = "end of CoverTab[120396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:39
		_go_fuzz_dep_.CoverTab[120397]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:39
		// _ = "end of CoverTab[120397]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:39
	// _ = "end of CoverTab[120394]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:39
	_go_fuzz_dep_.CoverTab[120395]++
										return nil, fmt.Errorf("invalid string being converted to IPNet: %s", sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:40
	// _ = "end of CoverTab[120395]"
}

// GetIPNet return the net.IPNet value of a flag with the given name
func (f *FlagSet) GetIPNet(name string) (net.IPNet, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:44
	_go_fuzz_dep_.CoverTab[120398]++
										val, err := f.getFlagType(name, "ipNet", ipNetConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:46
		_go_fuzz_dep_.CoverTab[120400]++
											return net.IPNet{}, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:47
		// _ = "end of CoverTab[120400]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:48
		_go_fuzz_dep_.CoverTab[120401]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:48
		// _ = "end of CoverTab[120401]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:48
	// _ = "end of CoverTab[120398]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:48
	_go_fuzz_dep_.CoverTab[120399]++
										return val.(net.IPNet), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:49
	// _ = "end of CoverTab[120399]"
}

// IPNetVar defines an net.IPNet flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:52
// The argument p points to an net.IPNet variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:54
func (f *FlagSet) IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:54
	_go_fuzz_dep_.CoverTab[120402]++
										f.VarP(newIPNetValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:55
	// _ = "end of CoverTab[120402]"
}

// IPNetVarP is like IPNetVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:59
	_go_fuzz_dep_.CoverTab[120403]++
										f.VarP(newIPNetValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:60
	// _ = "end of CoverTab[120403]"
}

// IPNetVar defines an net.IPNet flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:63
// The argument p points to an net.IPNet variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:65
func IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:65
	_go_fuzz_dep_.CoverTab[120404]++
										CommandLine.VarP(newIPNetValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:66
	// _ = "end of CoverTab[120404]"
}

// IPNetVarP is like IPNetVar, but accepts a shorthand letter that can be used after a single dash.
func IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:70
	_go_fuzz_dep_.CoverTab[120405]++
										CommandLine.VarP(newIPNetValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:71
	// _ = "end of CoverTab[120405]"
}

// IPNet defines an net.IPNet flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:74
// The return value is the address of an net.IPNet variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:76
func (f *FlagSet) IPNet(name string, value net.IPNet, usage string) *net.IPNet {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:76
	_go_fuzz_dep_.CoverTab[120406]++
										p := new(net.IPNet)
										f.IPNetVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:79
	// _ = "end of CoverTab[120406]"
}

// IPNetP is like IPNet, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:83
	_go_fuzz_dep_.CoverTab[120407]++
										p := new(net.IPNet)
										f.IPNetVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:86
	// _ = "end of CoverTab[120407]"
}

// IPNet defines an net.IPNet flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:89
// The return value is the address of an net.IPNet variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:91
func IPNet(name string, value net.IPNet, usage string) *net.IPNet {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:91
	_go_fuzz_dep_.CoverTab[120408]++
										return CommandLine.IPNetP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:92
	// _ = "end of CoverTab[120408]"
}

// IPNetP is like IPNet, but accepts a shorthand letter that can be used after a single dash.
func IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:96
	_go_fuzz_dep_.CoverTab[120409]++
										return CommandLine.IPNetP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:97
	// _ = "end of CoverTab[120409]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ipnet.go:98
var _ = _go_fuzz_dep_.CoverTab
