//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
package cacheobject

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:18
)

// Repersents a potential Reason to not cache an object.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:20
//
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:20
// Applications may wish to ignore specific reasons, which will make them non-RFC
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:20
// compliant, but this type gives them specific cases they can choose to ignore,
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:20
// making them compliant in as many cases as they can.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:25
type Reason int

const (

	// The request method was POST and an Expiration header was not supplied.
	ReasonRequestMethodPOST	Reason	= iota

	// The request method was PUT and PUTs are not cachable.
	ReasonRequestMethodPUT

	// The request method was DELETE and DELETEs are not cachable.
	ReasonRequestMethodDELETE

	// The request method was CONNECT and CONNECTs are not cachable.
	ReasonRequestMethodCONNECT

	// The request method was OPTIONS and OPTIONS are not cachable.
	ReasonRequestMethodOPTIONS

	// The request method was TRACE and TRACEs are not cachable.
	ReasonRequestMethodTRACE

	// The request method was not recognized by cachecontrol, and should not be cached.
	ReasonRequestMethodUnkown

	// The request included an Cache-Control: no-store header
	ReasonRequestNoStore

	// The request included an Authorization header without an explicit Public or Expiration time: http://tools.ietf.org/html/rfc7234#section-3.2
	ReasonRequestAuthorizationHeader

	// The response included an Cache-Control: no-store header
	ReasonResponseNoStore

	// The response included an Cache-Control: private header and this is not a Private cache
	ReasonResponsePrivate

	// The response failed to meet at least one of the conditions specified in RFC 7234 section 3: http://tools.ietf.org/html/rfc7234#section-3
	ReasonResponseUncachableByDefault
)

func (r Reason) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:66
	_go_fuzz_dep_.CoverTab[183995]++
																switch r {
	case ReasonRequestMethodPOST:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:68
		_go_fuzz_dep_.CoverTab[183997]++
																	return "ReasonRequestMethodPOST"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:69
		// _ = "end of CoverTab[183997]"
	case ReasonRequestMethodPUT:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:70
		_go_fuzz_dep_.CoverTab[183998]++
																	return "ReasonRequestMethodPUT"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:71
		// _ = "end of CoverTab[183998]"
	case ReasonRequestMethodDELETE:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:72
		_go_fuzz_dep_.CoverTab[183999]++
																	return "ReasonRequestMethodDELETE"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:73
		// _ = "end of CoverTab[183999]"
	case ReasonRequestMethodCONNECT:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:74
		_go_fuzz_dep_.CoverTab[184000]++
																	return "ReasonRequestMethodCONNECT"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:75
		// _ = "end of CoverTab[184000]"
	case ReasonRequestMethodOPTIONS:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:76
		_go_fuzz_dep_.CoverTab[184001]++
																	return "ReasonRequestMethodOPTIONS"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:77
		// _ = "end of CoverTab[184001]"
	case ReasonRequestMethodTRACE:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:78
		_go_fuzz_dep_.CoverTab[184002]++
																	return "ReasonRequestMethodTRACE"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:79
		// _ = "end of CoverTab[184002]"
	case ReasonRequestMethodUnkown:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:80
		_go_fuzz_dep_.CoverTab[184003]++
																	return "ReasonRequestMethodUnkown"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:81
		// _ = "end of CoverTab[184003]"
	case ReasonRequestNoStore:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:82
		_go_fuzz_dep_.CoverTab[184004]++
																	return "ReasonRequestNoStore"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:83
		// _ = "end of CoverTab[184004]"
	case ReasonRequestAuthorizationHeader:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:84
		_go_fuzz_dep_.CoverTab[184005]++
																	return "ReasonRequestAuthorizationHeader"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:85
		// _ = "end of CoverTab[184005]"
	case ReasonResponseNoStore:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:86
		_go_fuzz_dep_.CoverTab[184006]++
																	return "ReasonResponseNoStore"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:87
		// _ = "end of CoverTab[184006]"
	case ReasonResponsePrivate:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:88
		_go_fuzz_dep_.CoverTab[184007]++
																	return "ReasonResponsePrivate"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:89
		// _ = "end of CoverTab[184007]"
	case ReasonResponseUncachableByDefault:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:90
		_go_fuzz_dep_.CoverTab[184008]++
																	return "ReasonResponseUncachableByDefault"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:91
		// _ = "end of CoverTab[184008]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:91
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:91
		_go_fuzz_dep_.CoverTab[184009]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:91
		// _ = "end of CoverTab[184009]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:92
	// _ = "end of CoverTab[183995]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:92
	_go_fuzz_dep_.CoverTab[183996]++

																panic(r)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:94
	// _ = "end of CoverTab[183996]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/reasons.go:95
var _ = _go_fuzz_dep_.CoverTab
