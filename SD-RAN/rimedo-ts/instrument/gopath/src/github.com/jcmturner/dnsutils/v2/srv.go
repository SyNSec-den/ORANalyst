//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
package dnsutils

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:1
)

import (
	"math/rand"
	"net"
	"sort"
)

// OrderedSRV returns a count of the results and a map keyed on the order they should be used.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
// This based on the records' priority and randomised selection based on their relative weighting.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
// The function's inputs are the same as those for net.LookupSRV
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
// To use in the correct order:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
// count, orderedSRV, err := OrderedSRV(service, proto, name)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
// i := 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//	for  i <= count {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//	  srv := orderedSRV[i]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//	  // Do something such as dial this SRV. If fails move on the the next or break if it succeeds.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//	  i += 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:9
//	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:21
func OrderedSRV(service, proto, name string) (int, map[int]*net.SRV, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:21
	_go_fuzz_dep_.CoverTab[83111]++
											_, addrs, err := net.LookupSRV(service, proto, name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:23
		_go_fuzz_dep_.CoverTab[83113]++
												return 0, make(map[int]*net.SRV), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:24
		// _ = "end of CoverTab[83113]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:25
		_go_fuzz_dep_.CoverTab[83114]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:25
		// _ = "end of CoverTab[83114]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:25
	// _ = "end of CoverTab[83111]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:25
	_go_fuzz_dep_.CoverTab[83112]++
											index, osrv := orderSRV(addrs)
											return index, osrv, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:27
	// _ = "end of CoverTab[83112]"
}

func orderSRV(addrs []*net.SRV) (int, map[int]*net.SRV) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:30
	_go_fuzz_dep_.CoverTab[83115]++
	// Initialise the ordered map
	var o int
	osrv := make(map[int]*net.SRV)

	prioMap := make(map[int][]*net.SRV, 0)
	for _, srv := range addrs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:36
		_go_fuzz_dep_.CoverTab[83119]++
												prioMap[int(srv.Priority)] = append(prioMap[int(srv.Priority)], srv)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:37
		// _ = "end of CoverTab[83119]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:38
	// _ = "end of CoverTab[83115]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:38
	_go_fuzz_dep_.CoverTab[83116]++

											priorities := make([]int, 0)
											for p := range prioMap {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:41
		_go_fuzz_dep_.CoverTab[83120]++
												priorities = append(priorities, p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:42
		// _ = "end of CoverTab[83120]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:43
	// _ = "end of CoverTab[83116]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:43
	_go_fuzz_dep_.CoverTab[83117]++

											var count int
											sort.Ints(priorities)
											for _, p := range priorities {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:47
		_go_fuzz_dep_.CoverTab[83121]++
												tos := weightedOrder(prioMap[p])
												for i, s := range tos {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:49
			_go_fuzz_dep_.CoverTab[83123]++
													count += 1
													osrv[o+i] = s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:51
			// _ = "end of CoverTab[83123]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:52
		// _ = "end of CoverTab[83121]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:52
		_go_fuzz_dep_.CoverTab[83122]++
												o += len(tos)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:53
		// _ = "end of CoverTab[83122]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:54
	// _ = "end of CoverTab[83117]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:54
	_go_fuzz_dep_.CoverTab[83118]++
											return count, osrv
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:55
	// _ = "end of CoverTab[83118]"
}

func weightedOrder(srvs []*net.SRV) map[int]*net.SRV {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:58
	_go_fuzz_dep_.CoverTab[83124]++
	// Get the total weight
	var tw int
	for _, s := range srvs {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:61
		_go_fuzz_dep_.CoverTab[83127]++
												tw += int(s.Weight)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:62
		// _ = "end of CoverTab[83127]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:63
	// _ = "end of CoverTab[83124]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:63
	_go_fuzz_dep_.CoverTab[83125]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:66
	o := 1
											osrv := make(map[int]*net.SRV)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:70
	l := len(srvs)
	for l > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:71
		_go_fuzz_dep_.CoverTab[83128]++
												i := rand.Intn(l)
												s := srvs[i]
												var rw int
												if tw > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:75
			_go_fuzz_dep_.CoverTab[83130]++

													rw = rand.Intn(tw) - int(s.Weight)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:77
			// _ = "end of CoverTab[83130]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:78
			_go_fuzz_dep_.CoverTab[83131]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:78
			// _ = "end of CoverTab[83131]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:78
		// _ = "end of CoverTab[83128]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:78
		_go_fuzz_dep_.CoverTab[83129]++
												if rw <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:79
			_go_fuzz_dep_.CoverTab[83132]++

													osrv[o] = s
													if len(srvs) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:82
				_go_fuzz_dep_.CoverTab[83134]++

														srvs[len(srvs)-1], srvs[i] = srvs[i], srvs[len(srvs)-1]
														srvs = srvs[:len(srvs)-1]
														l = len(srvs)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:86
				// _ = "end of CoverTab[83134]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:87
				_go_fuzz_dep_.CoverTab[83135]++
														l = 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:88
				// _ = "end of CoverTab[83135]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:89
			// _ = "end of CoverTab[83132]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:89
			_go_fuzz_dep_.CoverTab[83133]++
													o += 1
													tw = tw - int(s.Weight)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:91
			// _ = "end of CoverTab[83133]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:92
			_go_fuzz_dep_.CoverTab[83136]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:92
			// _ = "end of CoverTab[83136]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:92
		// _ = "end of CoverTab[83129]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:93
	// _ = "end of CoverTab[83125]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:93
	_go_fuzz_dep_.CoverTab[83126]++
											return osrv
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:94
	// _ = "end of CoverTab[83126]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/dnsutils/v2@v2.0.0/srv.go:95
var _ = _go_fuzz_dep_.CoverTab
