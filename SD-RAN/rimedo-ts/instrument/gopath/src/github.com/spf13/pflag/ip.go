//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
package pflag

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:1
)

import (
	"fmt"
	"net"
	"strings"
)

// -- net.IP value
type ipValue net.IP

func newIPValue(val net.IP, p *net.IP) *ipValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:12
	_go_fuzz_dep_.CoverTab[120274]++
										*p = val
										return (*ipValue)(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:14
	// _ = "end of CoverTab[120274]"
}

func (i *ipValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:17
	_go_fuzz_dep_.CoverTab[120275]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:17
	return net.IP(*i).String()
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:17
	// _ = "end of CoverTab[120275]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:17
}
func (i *ipValue) Set(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:18
	_go_fuzz_dep_.CoverTab[120276]++
										ip := net.ParseIP(strings.TrimSpace(s))
										if ip == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:20
		_go_fuzz_dep_.CoverTab[120278]++
											return fmt.Errorf("failed to parse IP: %q", s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:21
		// _ = "end of CoverTab[120278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:22
		_go_fuzz_dep_.CoverTab[120279]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:22
		// _ = "end of CoverTab[120279]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:22
	// _ = "end of CoverTab[120276]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:22
	_go_fuzz_dep_.CoverTab[120277]++
										*i = ipValue(ip)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:24
	// _ = "end of CoverTab[120277]"
}

func (i *ipValue) Type() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:27
	_go_fuzz_dep_.CoverTab[120280]++
										return "ip"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:28
	// _ = "end of CoverTab[120280]"
}

func ipConv(sval string) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:31
	_go_fuzz_dep_.CoverTab[120281]++
										ip := net.ParseIP(sval)
										if ip != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:33
		_go_fuzz_dep_.CoverTab[120283]++
											return ip, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:34
		// _ = "end of CoverTab[120283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:35
		_go_fuzz_dep_.CoverTab[120284]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:35
		// _ = "end of CoverTab[120284]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:35
	// _ = "end of CoverTab[120281]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:35
	_go_fuzz_dep_.CoverTab[120282]++
										return nil, fmt.Errorf("invalid string being converted to IP address: %s", sval)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:36
	// _ = "end of CoverTab[120282]"
}

// GetIP return the net.IP value of a flag with the given name
func (f *FlagSet) GetIP(name string) (net.IP, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:40
	_go_fuzz_dep_.CoverTab[120285]++
										val, err := f.getFlagType(name, "ip", ipConv)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:42
		_go_fuzz_dep_.CoverTab[120287]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:43
		// _ = "end of CoverTab[120287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:44
		_go_fuzz_dep_.CoverTab[120288]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:44
		// _ = "end of CoverTab[120288]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:44
	// _ = "end of CoverTab[120285]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:44
	_go_fuzz_dep_.CoverTab[120286]++
										return val.(net.IP), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:45
	// _ = "end of CoverTab[120286]"
}

// IPVar defines an net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:48
// The argument p points to an net.IP variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:50
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:50
	_go_fuzz_dep_.CoverTab[120289]++
										f.VarP(newIPValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:51
	// _ = "end of CoverTab[120289]"
}

// IPVarP is like IPVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:55
	_go_fuzz_dep_.CoverTab[120290]++
										f.VarP(newIPValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:56
	// _ = "end of CoverTab[120290]"
}

// IPVar defines an net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:59
// The argument p points to an net.IP variable in which to store the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:61
func IPVar(p *net.IP, name string, value net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:61
	_go_fuzz_dep_.CoverTab[120291]++
										CommandLine.VarP(newIPValue(value, p), name, "", usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:62
	// _ = "end of CoverTab[120291]"
}

// IPVarP is like IPVar, but accepts a shorthand letter that can be used after a single dash.
func IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:66
	_go_fuzz_dep_.CoverTab[120292]++
										CommandLine.VarP(newIPValue(value, p), name, shorthand, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:67
	// _ = "end of CoverTab[120292]"
}

// IP defines an net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:70
// The return value is the address of an net.IP variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:72
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:72
	_go_fuzz_dep_.CoverTab[120293]++
										p := new(net.IP)
										f.IPVarP(p, name, "", value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:75
	// _ = "end of CoverTab[120293]"
}

// IPP is like IP, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPP(name, shorthand string, value net.IP, usage string) *net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:79
	_go_fuzz_dep_.CoverTab[120294]++
										p := new(net.IP)
										f.IPVarP(p, name, shorthand, value, usage)
										return p
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:82
	// _ = "end of CoverTab[120294]"
}

// IP defines an net.IP flag with specified name, default value, and usage string.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:85
// The return value is the address of an net.IP variable that stores the value of the flag.
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:87
func IP(name string, value net.IP, usage string) *net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:87
	_go_fuzz_dep_.CoverTab[120295]++
										return CommandLine.IPP(name, "", value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:88
	// _ = "end of CoverTab[120295]"
}

// IPP is like IP, but accepts a shorthand letter that can be used after a single dash.
func IPP(name, shorthand string, value net.IP, usage string) *net.IP {
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:92
	_go_fuzz_dep_.CoverTab[120296]++
										return CommandLine.IPP(name, shorthand, value, usage)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:93
	// _ = "end of CoverTab[120296]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/pflag@v1.0.5/ip.go:94
var _ = _go_fuzz_dep_.CoverTab
