//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
package cacheobject

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:18
)

import (
	"net/http"
	"time"
)

// LOW LEVEL API: Repersents a potentially cachable HTTP object.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:25
// This struct is designed to be serialized efficiently, so in a high
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:25
// performance caching server, things like Date-Strings don't need to be
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:25
// parsed for every use of a cached object.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:30
type Object struct {
	CacheIsPrivate	bool

	RespDirectives		*ResponseCacheDirectives
	RespHeaders		http.Header
	RespStatusCode		int
	RespExpiresHeader	time.Time
	RespDateHeader		time.Time
	RespLastModifiedHeader	time.Time

	ReqDirectives	*RequestCacheDirectives
	ReqHeaders	http.Header
	ReqMethod	string

	NowUTC	time.Time
}

// LOW LEVEL API: Repersents the results of examinig an Object with
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:47
// CachableObject and ExpirationObject.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:47
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:47
// TODO(pquerna): decide if this is a good idea or bad
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:51
type ObjectResults struct {
	OutReasons		[]Reason
	OutWarnings		[]Warning
	OutExpirationTime	time.Time
	OutErr			error
}

// LOW LEVEL API: Check if a object is cachable.
func CachableObject(obj *Object, rv *ObjectResults) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:59
	_go_fuzz_dep_.CoverTab[183879]++
																rv.OutReasons = nil
																rv.OutWarnings = nil
																rv.OutErr = nil

																switch obj.ReqMethod {
	case "GET":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:65
		_go_fuzz_dep_.CoverTab[183885]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:66
		// _ = "end of CoverTab[183885]"
	case "HEAD":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:67
		_go_fuzz_dep_.CoverTab[183886]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:68
		// _ = "end of CoverTab[183886]"
	case "POST":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:69
		_go_fuzz_dep_.CoverTab[183887]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:82
		if !hasFreshness(obj.ReqDirectives, obj.RespDirectives, obj.RespHeaders, obj.RespExpiresHeader, obj.CacheIsPrivate) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:82
			_go_fuzz_dep_.CoverTab[183894]++
																		rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodPOST)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:83
			// _ = "end of CoverTab[183894]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:84
			_go_fuzz_dep_.CoverTab[183895]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:84
			// _ = "end of CoverTab[183895]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:84
		// _ = "end of CoverTab[183887]"

	case "PUT":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:86
		_go_fuzz_dep_.CoverTab[183888]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodPUT)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:87
		// _ = "end of CoverTab[183888]"

	case "DELETE":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:89
		_go_fuzz_dep_.CoverTab[183889]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodDELETE)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:90
		// _ = "end of CoverTab[183889]"

	case "CONNECT":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:92
		_go_fuzz_dep_.CoverTab[183890]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodCONNECT)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:93
		// _ = "end of CoverTab[183890]"

	case "OPTIONS":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:95
		_go_fuzz_dep_.CoverTab[183891]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodOPTIONS)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:96
		// _ = "end of CoverTab[183891]"

	case "TRACE":
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:98
		_go_fuzz_dep_.CoverTab[183892]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodTRACE)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:99
		// _ = "end of CoverTab[183892]"

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:105
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:105
		_go_fuzz_dep_.CoverTab[183893]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestMethodUnkown)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:106
		// _ = "end of CoverTab[183893]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:107
	// _ = "end of CoverTab[183879]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:107
	_go_fuzz_dep_.CoverTab[183880]++

																if obj.ReqDirectives.NoStore {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:109
		_go_fuzz_dep_.CoverTab[183896]++
																	rv.OutReasons = append(rv.OutReasons, ReasonRequestNoStore)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:110
		// _ = "end of CoverTab[183896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:111
		_go_fuzz_dep_.CoverTab[183897]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:111
		// _ = "end of CoverTab[183897]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:111
	// _ = "end of CoverTab[183880]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:111
	_go_fuzz_dep_.CoverTab[183881]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:114
	authz := obj.ReqHeaders.Get("Authorization")
	if authz != "" {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:115
		_go_fuzz_dep_.CoverTab[183898]++
																	if obj.RespDirectives.MustRevalidate || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:116
			_go_fuzz_dep_.CoverTab[183899]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:116
			return obj.RespDirectives.Public
																		// _ = "end of CoverTab[183899]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:117
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:117
			_go_fuzz_dep_.CoverTab[183900]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:117
			return obj.RespDirectives.SMaxAge != -1
																		// _ = "end of CoverTab[183900]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:118
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:118
			_go_fuzz_dep_.CoverTab[183901]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:118
			// _ = "end of CoverTab[183901]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:120
			_go_fuzz_dep_.CoverTab[183902]++
																		rv.OutReasons = append(rv.OutReasons, ReasonRequestAuthorizationHeader)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:121
			// _ = "end of CoverTab[183902]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:122
		// _ = "end of CoverTab[183898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:123
		_go_fuzz_dep_.CoverTab[183903]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:123
		// _ = "end of CoverTab[183903]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:123
	// _ = "end of CoverTab[183881]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:123
	_go_fuzz_dep_.CoverTab[183882]++

																if obj.RespDirectives.PrivatePresent && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:125
		_go_fuzz_dep_.CoverTab[183904]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:125
		return !obj.CacheIsPrivate
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:125
		// _ = "end of CoverTab[183904]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:125
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:125
		_go_fuzz_dep_.CoverTab[183905]++
																	rv.OutReasons = append(rv.OutReasons, ReasonResponsePrivate)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:126
		// _ = "end of CoverTab[183905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:127
		_go_fuzz_dep_.CoverTab[183906]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:127
		// _ = "end of CoverTab[183906]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:127
	// _ = "end of CoverTab[183882]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:127
	_go_fuzz_dep_.CoverTab[183883]++

																if obj.RespDirectives.NoStore {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:129
		_go_fuzz_dep_.CoverTab[183907]++
																	rv.OutReasons = append(rv.OutReasons, ReasonResponseNoStore)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:130
		// _ = "end of CoverTab[183907]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:131
		_go_fuzz_dep_.CoverTab[183908]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:131
		// _ = "end of CoverTab[183908]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:131
	// _ = "end of CoverTab[183883]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:131
	_go_fuzz_dep_.CoverTab[183884]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:152
	expires := obj.RespHeaders.Get("Expires") != ""
	statusCachable := cachableStatusCode(obj.RespStatusCode)

	if expires || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:155
		_go_fuzz_dep_.CoverTab[183909]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:155
		return obj.RespDirectives.MaxAge != -1
																	// _ = "end of CoverTab[183909]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:156
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:156
		_go_fuzz_dep_.CoverTab[183910]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:156
		return (obj.RespDirectives.SMaxAge != -1 && func() bool {
																		_go_fuzz_dep_.CoverTab[183911]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
			return !obj.CacheIsPrivate
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
			// _ = "end of CoverTab[183911]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
		}())
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
		// _ = "end of CoverTab[183910]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
		_go_fuzz_dep_.CoverTab[183912]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:157
		return statusCachable
																	// _ = "end of CoverTab[183912]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:158
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:158
		_go_fuzz_dep_.CoverTab[183913]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:158
		return obj.RespDirectives.Public
																	// _ = "end of CoverTab[183913]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:159
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:159
		_go_fuzz_dep_.CoverTab[183914]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:159
		// _ = "end of CoverTab[183914]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:161
		_go_fuzz_dep_.CoverTab[183915]++
																	rv.OutReasons = append(rv.OutReasons, ReasonResponseUncachableByDefault)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:162
		// _ = "end of CoverTab[183915]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:163
	// _ = "end of CoverTab[183884]"
}

var twentyFourHours = time.Duration(24 * time.Hour)

const debug = false

// LOW LEVEL API: Update an objects expiration time.
func ExpirationObject(obj *Object, rv *ObjectResults) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:171
	_go_fuzz_dep_.CoverTab[183916]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:192
	var expiresTime time.Time

	if obj.RespDirectives.SMaxAge != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:194
		_go_fuzz_dep_.CoverTab[183918]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:194
		return !obj.CacheIsPrivate
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:194
		// _ = "end of CoverTab[183918]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:194
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:194
		_go_fuzz_dep_.CoverTab[183919]++
																	expiresTime = obj.NowUTC.Add(time.Second * time.Duration(obj.RespDirectives.SMaxAge))
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:195
		// _ = "end of CoverTab[183919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:196
		_go_fuzz_dep_.CoverTab[183920]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:196
		if obj.RespDirectives.MaxAge != -1 {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:196
			_go_fuzz_dep_.CoverTab[183921]++
																		expiresTime = obj.NowUTC.UTC().Add(time.Second * time.Duration(obj.RespDirectives.MaxAge))
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:197
			// _ = "end of CoverTab[183921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:198
			_go_fuzz_dep_.CoverTab[183922]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:198
			if !obj.RespExpiresHeader.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:198
				_go_fuzz_dep_.CoverTab[183923]++
																			serverDate := obj.RespDateHeader
																			if serverDate.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:200
					_go_fuzz_dep_.CoverTab[183925]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:203
					serverDate = obj.NowUTC
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:203
					// _ = "end of CoverTab[183925]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:204
					_go_fuzz_dep_.CoverTab[183926]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:204
					// _ = "end of CoverTab[183926]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:204
				}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:204
				// _ = "end of CoverTab[183923]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:204
				_go_fuzz_dep_.CoverTab[183924]++
																			expiresTime = obj.NowUTC.Add(obj.RespExpiresHeader.Sub(serverDate))
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:205
				// _ = "end of CoverTab[183924]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:206
				_go_fuzz_dep_.CoverTab[183927]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:206
				if !obj.RespLastModifiedHeader.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:206
					_go_fuzz_dep_.CoverTab[183928]++

																				rv.OutWarnings = append(rv.OutWarnings, WarningHeuristicExpiration)

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:218
					since := obj.RespLastModifiedHeader.Sub(obj.NowUTC)
					since = time.Duration(float64(since) * -0.1)

					if since > twentyFourHours {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:221
						_go_fuzz_dep_.CoverTab[183930]++
																					expiresTime = obj.NowUTC.Add(twentyFourHours)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:222
						// _ = "end of CoverTab[183930]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:223
						_go_fuzz_dep_.CoverTab[183931]++
																					expiresTime = obj.NowUTC.Add(since)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:224
						// _ = "end of CoverTab[183931]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:225
					// _ = "end of CoverTab[183928]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:225
					_go_fuzz_dep_.CoverTab[183929]++

																				if debug {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:227
						_go_fuzz_dep_.CoverTab[183932]++
																					println("Now UTC: ", obj.NowUTC.String())
																					println("Last-Modified: ", obj.RespLastModifiedHeader.String())
																					println("Since: ", since.String())
																					println("TwentyFourHours: ", twentyFourHours.String())
																					println("Expiration: ", expiresTime.String())
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:232
						// _ = "end of CoverTab[183932]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:233
						_go_fuzz_dep_.CoverTab[183933]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:233
						// _ = "end of CoverTab[183933]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:233
					}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:233
					// _ = "end of CoverTab[183929]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:234
					_go_fuzz_dep_.CoverTab[183934]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:234
					// _ = "end of CoverTab[183934]"

				}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
				// _ = "end of CoverTab[183927]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
			}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
			// _ = "end of CoverTab[183922]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
		// _ = "end of CoverTab[183920]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
	// _ = "end of CoverTab[183916]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:236
	_go_fuzz_dep_.CoverTab[183917]++

																rv.OutExpirationTime = expiresTime
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:238
	// _ = "end of CoverTab[183917]"
}

// Evaluate cachability based on an HTTP request, and parts of the response.
func UsingRequestResponse(req *http.Request,
	statusCode int,
	respHeaders http.Header,
	privateCache bool) ([]Reason, time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:245
	_go_fuzz_dep_.CoverTab[183935]++
																reasons, time, _, _, err := UsingRequestResponseWithObject(req, statusCode, respHeaders, privateCache)
																return reasons, time, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:247
	// _ = "end of CoverTab[183935]"
}

// Evaluate cachability based on an HTTP request, and parts of the response.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:250
// Returns the parsed Object as well.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:252
func UsingRequestResponseWithObject(req *http.Request,
	statusCode int,
	respHeaders http.Header,
	privateCache bool) ([]Reason, time.Time, []Warning, *Object, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:255
	_go_fuzz_dep_.CoverTab[183936]++
																var reqHeaders http.Header
																var reqMethod string

																var reqDir *RequestCacheDirectives = nil
																respDir, err := ParseResponseCacheControl(respHeaders.Get("Cache-Control"))
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:261
		_go_fuzz_dep_.CoverTab[183944]++
																	return nil, time.Time{}, nil, nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:262
		// _ = "end of CoverTab[183944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:263
		_go_fuzz_dep_.CoverTab[183945]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:263
		// _ = "end of CoverTab[183945]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:263
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:263
	// _ = "end of CoverTab[183936]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:263
	_go_fuzz_dep_.CoverTab[183937]++

																if req != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:265
		_go_fuzz_dep_.CoverTab[183946]++
																	reqDir, err = ParseRequestCacheControl(req.Header.Get("Cache-Control"))
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:267
			_go_fuzz_dep_.CoverTab[183948]++
																		return nil, time.Time{}, nil, nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:268
			// _ = "end of CoverTab[183948]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:269
			_go_fuzz_dep_.CoverTab[183949]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:269
			// _ = "end of CoverTab[183949]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:269
		// _ = "end of CoverTab[183946]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:269
		_go_fuzz_dep_.CoverTab[183947]++
																	reqHeaders = req.Header
																	reqMethod = req.Method
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:271
		// _ = "end of CoverTab[183947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:272
		_go_fuzz_dep_.CoverTab[183950]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:272
		// _ = "end of CoverTab[183950]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:272
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:272
	// _ = "end of CoverTab[183937]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:272
	_go_fuzz_dep_.CoverTab[183938]++

																var expiresHeader time.Time
																var dateHeader time.Time
																var lastModifiedHeader time.Time

																if respHeaders.Get("Expires") != "" {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:278
		_go_fuzz_dep_.CoverTab[183951]++
																	expiresHeader, err = http.ParseTime(respHeaders.Get("Expires"))
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:280
			_go_fuzz_dep_.CoverTab[183953]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:283
			expiresHeader = time.Time{}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:283
			// _ = "end of CoverTab[183953]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:284
			_go_fuzz_dep_.CoverTab[183954]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:284
			// _ = "end of CoverTab[183954]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:284
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:284
		// _ = "end of CoverTab[183951]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:284
		_go_fuzz_dep_.CoverTab[183952]++
																	expiresHeader = expiresHeader.UTC()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:285
		// _ = "end of CoverTab[183952]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:286
		_go_fuzz_dep_.CoverTab[183955]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:286
		// _ = "end of CoverTab[183955]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:286
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:286
	// _ = "end of CoverTab[183938]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:286
	_go_fuzz_dep_.CoverTab[183939]++

																if respHeaders.Get("Date") != "" {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:288
		_go_fuzz_dep_.CoverTab[183956]++
																	dateHeader, err = http.ParseTime(respHeaders.Get("Date"))
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:290
			_go_fuzz_dep_.CoverTab[183958]++
																		return nil, time.Time{}, nil, nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:291
			// _ = "end of CoverTab[183958]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:292
			_go_fuzz_dep_.CoverTab[183959]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:292
			// _ = "end of CoverTab[183959]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:292
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:292
		// _ = "end of CoverTab[183956]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:292
		_go_fuzz_dep_.CoverTab[183957]++
																	dateHeader = dateHeader.UTC()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:293
		// _ = "end of CoverTab[183957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:294
		_go_fuzz_dep_.CoverTab[183960]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:294
		// _ = "end of CoverTab[183960]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:294
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:294
	// _ = "end of CoverTab[183939]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:294
	_go_fuzz_dep_.CoverTab[183940]++

																if respHeaders.Get("Last-Modified") != "" {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:296
		_go_fuzz_dep_.CoverTab[183961]++
																	lastModifiedHeader, err = http.ParseTime(respHeaders.Get("Last-Modified"))
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:298
			_go_fuzz_dep_.CoverTab[183963]++
																		return nil, time.Time{}, nil, nil, err
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:299
			// _ = "end of CoverTab[183963]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:300
			_go_fuzz_dep_.CoverTab[183964]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:300
			// _ = "end of CoverTab[183964]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:300
		// _ = "end of CoverTab[183961]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:300
		_go_fuzz_dep_.CoverTab[183962]++
																	lastModifiedHeader = lastModifiedHeader.UTC()
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:301
		// _ = "end of CoverTab[183962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:302
		_go_fuzz_dep_.CoverTab[183965]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:302
		// _ = "end of CoverTab[183965]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:302
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:302
	// _ = "end of CoverTab[183940]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:302
	_go_fuzz_dep_.CoverTab[183941]++

																obj := Object{
		CacheIsPrivate:	privateCache,

		RespDirectives:		respDir,
		RespHeaders:		respHeaders,
		RespStatusCode:		statusCode,
		RespExpiresHeader:	expiresHeader,
		RespDateHeader:		dateHeader,
		RespLastModifiedHeader:	lastModifiedHeader,

		ReqDirectives:	reqDir,
		ReqHeaders:	reqHeaders,
		ReqMethod:	reqMethod,

		NowUTC:	time.Now().UTC(),
	}
	rv := ObjectResults{}

	CachableObject(&obj, &rv)
	if rv.OutErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:323
		_go_fuzz_dep_.CoverTab[183966]++
																	return nil, time.Time{}, nil, nil, rv.OutErr
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:324
		// _ = "end of CoverTab[183966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:325
		_go_fuzz_dep_.CoverTab[183967]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:325
		// _ = "end of CoverTab[183967]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:325
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:325
	// _ = "end of CoverTab[183941]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:325
	_go_fuzz_dep_.CoverTab[183942]++

																ExpirationObject(&obj, &rv)
																if rv.OutErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:328
		_go_fuzz_dep_.CoverTab[183968]++
																	return nil, time.Time{}, nil, nil, rv.OutErr
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:329
		// _ = "end of CoverTab[183968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:330
		_go_fuzz_dep_.CoverTab[183969]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:330
		// _ = "end of CoverTab[183969]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:330
	// _ = "end of CoverTab[183942]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:330
	_go_fuzz_dep_.CoverTab[183943]++

																return rv.OutReasons, rv.OutExpirationTime, rv.OutWarnings, &obj, nil
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:332
	// _ = "end of CoverTab[183943]"
}

// calculate if a freshness directive is present: http://tools.ietf.org/html/rfc7234#section-4.2.1
func hasFreshness(reqDir *RequestCacheDirectives, respDir *ResponseCacheDirectives, respHeaders http.Header, respExpires time.Time, privateCache bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:336
	_go_fuzz_dep_.CoverTab[183970]++
																if !privateCache && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:337
		_go_fuzz_dep_.CoverTab[183974]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:337
		return respDir.SMaxAge != -1
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:337
		// _ = "end of CoverTab[183974]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:337
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:337
		_go_fuzz_dep_.CoverTab[183975]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:338
		// _ = "end of CoverTab[183975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:339
		_go_fuzz_dep_.CoverTab[183976]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:339
		// _ = "end of CoverTab[183976]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:339
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:339
	// _ = "end of CoverTab[183970]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:339
	_go_fuzz_dep_.CoverTab[183971]++

																if respDir.MaxAge != -1 {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:341
		_go_fuzz_dep_.CoverTab[183977]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:342
		// _ = "end of CoverTab[183977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:343
		_go_fuzz_dep_.CoverTab[183978]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:343
		// _ = "end of CoverTab[183978]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:343
	// _ = "end of CoverTab[183971]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:343
	_go_fuzz_dep_.CoverTab[183972]++

																if !respExpires.IsZero() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:345
		_go_fuzz_dep_.CoverTab[183979]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:345
		return respHeaders.Get("Expires") != ""
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:345
		// _ = "end of CoverTab[183979]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:345
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:345
		_go_fuzz_dep_.CoverTab[183980]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:346
		// _ = "end of CoverTab[183980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:347
		_go_fuzz_dep_.CoverTab[183981]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:347
		// _ = "end of CoverTab[183981]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:347
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:347
	// _ = "end of CoverTab[183972]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:347
	_go_fuzz_dep_.CoverTab[183973]++

																return false
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:349
	// _ = "end of CoverTab[183973]"
}

func cachableStatusCode(statusCode int) bool {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:352
	_go_fuzz_dep_.CoverTab[183982]++

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:361
	switch statusCode {
	case 200:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:362
		_go_fuzz_dep_.CoverTab[183983]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:363
		// _ = "end of CoverTab[183983]"
	case 203:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:364
		_go_fuzz_dep_.CoverTab[183984]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:365
		// _ = "end of CoverTab[183984]"
	case 204:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:366
		_go_fuzz_dep_.CoverTab[183985]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:367
		// _ = "end of CoverTab[183985]"
	case 206:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:368
		_go_fuzz_dep_.CoverTab[183986]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:369
		// _ = "end of CoverTab[183986]"
	case 300:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:370
		_go_fuzz_dep_.CoverTab[183987]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:371
		// _ = "end of CoverTab[183987]"
	case 301:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:372
		_go_fuzz_dep_.CoverTab[183988]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:373
		// _ = "end of CoverTab[183988]"
	case 404:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:374
		_go_fuzz_dep_.CoverTab[183989]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:375
		// _ = "end of CoverTab[183989]"
	case 405:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:376
		_go_fuzz_dep_.CoverTab[183990]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:377
		// _ = "end of CoverTab[183990]"
	case 410:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:378
		_go_fuzz_dep_.CoverTab[183991]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:379
		// _ = "end of CoverTab[183991]"
	case 414:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:380
		_go_fuzz_dep_.CoverTab[183992]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:381
		// _ = "end of CoverTab[183992]"
	case 501:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:382
		_go_fuzz_dep_.CoverTab[183993]++
																	return true
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:383
		// _ = "end of CoverTab[183993]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:384
		_go_fuzz_dep_.CoverTab[183994]++
																	return false
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:385
		// _ = "end of CoverTab[183994]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:386
	// _ = "end of CoverTab[183982]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:387
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/object.go:387
var _ = _go_fuzz_dep_.CoverTab
