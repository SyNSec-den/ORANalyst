//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
package resolver

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:19
)

type addressMapEntry struct {
	addr	Address
	value	interface{}
}

// AddressMap is a map of addresses to arbitrary values taking into account
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:26
// Attributes.  BalancerAttributes are ignored, as are Metadata and Type.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:26
// Multiple accesses may not be performed concurrently.  Must be created via
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:26
// NewAddressMap; do not construct directly.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:30
type AddressMap struct {
	// The underlying map is keyed by an Address with fields that we don't care
	// about being set to their zero values. The only fields that we care about
	// are `Addr`, `ServerName` and `Attributes`. Since we need to be able to
	// distinguish between addresses with same `Addr` and `ServerName`, but
	// different `Attributes`, we cannot store the `Attributes` in the map key.
	//
	// The comparison operation for structs work as follows:
	//  Struct values are comparable if all their fields are comparable. Two
	//  struct values are equal if their corresponding non-blank fields are equal.
	//
	// The value type of the map contains a slice of addresses which match the key
	// in their `Addr` and `ServerName` fields and contain the corresponding value
	// associated with them.
	m map[Address]addressMapEntryList
}

func toMapKey(addr *Address) Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:47
	_go_fuzz_dep_.CoverTab[67277]++
											return Address{Addr: addr.Addr, ServerName: addr.ServerName}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:48
	// _ = "end of CoverTab[67277]"
}

type addressMapEntryList []*addressMapEntry

// NewAddressMap creates a new AddressMap.
func NewAddressMap() *AddressMap {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:54
	_go_fuzz_dep_.CoverTab[67278]++
											return &AddressMap{m: make(map[Address]addressMapEntryList)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:55
	// _ = "end of CoverTab[67278]"
}

// find returns the index of addr in the addressMapEntry slice, or -1 if not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:58
// present.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:60
func (l addressMapEntryList) find(addr Address) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:60
	_go_fuzz_dep_.CoverTab[67279]++
											for i, entry := range l {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:61
		_go_fuzz_dep_.CoverTab[67281]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:64
		if entry.addr.Attributes.Equal(addr.Attributes) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:64
			_go_fuzz_dep_.CoverTab[67282]++
													return i
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:65
			// _ = "end of CoverTab[67282]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:66
			_go_fuzz_dep_.CoverTab[67283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:66
			// _ = "end of CoverTab[67283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:66
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:66
		// _ = "end of CoverTab[67281]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:67
	// _ = "end of CoverTab[67279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:67
	_go_fuzz_dep_.CoverTab[67280]++
											return -1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:68
	// _ = "end of CoverTab[67280]"
}

// Get returns the value for the address in the map, if present.
func (a *AddressMap) Get(addr Address) (value interface{}, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:72
	_go_fuzz_dep_.CoverTab[67284]++
											addrKey := toMapKey(&addr)
											entryList := a.m[addrKey]
											if entry := entryList.find(addr); entry != -1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:75
		_go_fuzz_dep_.CoverTab[67286]++
												return entryList[entry].value, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:76
		// _ = "end of CoverTab[67286]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:77
		_go_fuzz_dep_.CoverTab[67287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:77
		// _ = "end of CoverTab[67287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:77
	// _ = "end of CoverTab[67284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:77
	_go_fuzz_dep_.CoverTab[67285]++
											return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:78
	// _ = "end of CoverTab[67285]"
}

// Set updates or adds the value to the address in the map.
func (a *AddressMap) Set(addr Address, value interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:82
	_go_fuzz_dep_.CoverTab[67288]++
											addrKey := toMapKey(&addr)
											entryList := a.m[addrKey]
											if entry := entryList.find(addr); entry != -1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:85
		_go_fuzz_dep_.CoverTab[67290]++
												entryList[entry].value = value
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:87
		// _ = "end of CoverTab[67290]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:88
		_go_fuzz_dep_.CoverTab[67291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:88
		// _ = "end of CoverTab[67291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:88
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:88
	// _ = "end of CoverTab[67288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:88
	_go_fuzz_dep_.CoverTab[67289]++
											a.m[addrKey] = append(entryList, &addressMapEntry{addr: addr, value: value})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:89
	// _ = "end of CoverTab[67289]"
}

// Delete removes addr from the map.
func (a *AddressMap) Delete(addr Address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:93
	_go_fuzz_dep_.CoverTab[67292]++
											addrKey := toMapKey(&addr)
											entryList := a.m[addrKey]
											entry := entryList.find(addr)
											if entry == -1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:97
		_go_fuzz_dep_.CoverTab[67295]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:98
		// _ = "end of CoverTab[67295]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:99
		_go_fuzz_dep_.CoverTab[67296]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:99
		// _ = "end of CoverTab[67296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:99
	// _ = "end of CoverTab[67292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:99
	_go_fuzz_dep_.CoverTab[67293]++
											if len(entryList) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:100
		_go_fuzz_dep_.CoverTab[67297]++
												entryList = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:101
		// _ = "end of CoverTab[67297]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:102
		_go_fuzz_dep_.CoverTab[67298]++
												copy(entryList[entry:], entryList[entry+1:])
												entryList = entryList[:len(entryList)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:104
		// _ = "end of CoverTab[67298]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:105
	// _ = "end of CoverTab[67293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:105
	_go_fuzz_dep_.CoverTab[67294]++
											a.m[addrKey] = entryList
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:106
	// _ = "end of CoverTab[67294]"
}

// Len returns the number of entries in the map.
func (a *AddressMap) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:110
	_go_fuzz_dep_.CoverTab[67299]++
											ret := 0
											for _, entryList := range a.m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:112
		_go_fuzz_dep_.CoverTab[67301]++
												ret += len(entryList)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:113
		// _ = "end of CoverTab[67301]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:114
	// _ = "end of CoverTab[67299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:114
	_go_fuzz_dep_.CoverTab[67300]++
											return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:115
	// _ = "end of CoverTab[67300]"
}

// Keys returns a slice of all current map keys.
func (a *AddressMap) Keys() []Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:119
	_go_fuzz_dep_.CoverTab[67302]++
											ret := make([]Address, 0, a.Len())
											for _, entryList := range a.m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:121
		_go_fuzz_dep_.CoverTab[67304]++
												for _, entry := range entryList {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:122
			_go_fuzz_dep_.CoverTab[67305]++
													ret = append(ret, entry.addr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:123
			// _ = "end of CoverTab[67305]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:124
		// _ = "end of CoverTab[67304]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:125
	// _ = "end of CoverTab[67302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:125
	_go_fuzz_dep_.CoverTab[67303]++
											return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:126
	// _ = "end of CoverTab[67303]"
}

// Values returns a slice of all current map values.
func (a *AddressMap) Values() []interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:130
	_go_fuzz_dep_.CoverTab[67306]++
											ret := make([]interface{}, 0, a.Len())
											for _, entryList := range a.m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:132
		_go_fuzz_dep_.CoverTab[67308]++
												for _, entry := range entryList {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:133
			_go_fuzz_dep_.CoverTab[67309]++
													ret = append(ret, entry.value)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:134
			// _ = "end of CoverTab[67309]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:135
		// _ = "end of CoverTab[67308]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:136
	// _ = "end of CoverTab[67306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:136
	_go_fuzz_dep_.CoverTab[67307]++
											return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:137
	// _ = "end of CoverTab[67307]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:138
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/resolver/map.go:138
var _ = _go_fuzz_dep_.CoverTab
