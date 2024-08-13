//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
package cacheobject

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:18
)

import (
	"fmt"
	"net/http"
	"time"
)

// Repersents an HTTP Warning: http://tools.ietf.org/html/rfc7234#section-5.5
type Warning int

const (
	// Response is Stale
	// A cache SHOULD generate this whenever the sent response is stale.
	WarningResponseIsStale	Warning	= 110

	// Revalidation Failed
	// A cache SHOULD generate this when sending a stale
	// response because an attempt to validate the response failed, due to an
	// inability to reach the server.
	WarningRevalidationFailed	Warning	= 111

	// Disconnected Operation
	// A cache SHOULD generate this if it is intentionally disconnected from
	// the rest of the network for a period of time.
	WarningDisconnectedOperation	Warning	= 112

	// Heuristic Expiration
	//
	// A cache SHOULD generate this if it heuristically chose a freshness
	// lifetime greater than 24 hours and the response's age is greater than
	// 24 hours.
	WarningHeuristicExpiration	Warning	= 113

	// Miscellaneous Warning
	//
	// The warning text can include arbitrary information to be presented to
	// a human user or logged.  A system receiving this warning MUST NOT
	// take any automated action, besides presenting the warning to the
	// user.
	WarningMiscellaneousWarning	Warning	= 199

	// Transformation Applied
	//
	// This Warning code MUST be added by a proxy if it applies any
	// transformation to the representation, such as changing the
	// content-coding, media-type, or modifying the representation data,
	// unless this Warning code already appears in the response.
	WarningTransformationApplied	Warning	= 214

	// Miscellaneous Persistent Warning
	//
	// The warning text can include arbitrary information to be presented to
	// a human user or logged.  A system receiving this warning MUST NOT
	// take any automated action.
	WarningMiscellaneousPersistentWarning	Warning	= 299
)

func (w Warning) HeaderString(agent string, date time.Time) string {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:76
	_go_fuzz_dep_.CoverTab[184010]++
																if agent == "" {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:77
		_go_fuzz_dep_.CoverTab[184012]++
																	agent = "-"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:78
		// _ = "end of CoverTab[184012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:79
		_go_fuzz_dep_.CoverTab[184013]++

																	agent = `"` + agent + `"`
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:81
		// _ = "end of CoverTab[184013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:82
	// _ = "end of CoverTab[184010]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:82
	_go_fuzz_dep_.CoverTab[184011]++
																return fmt.Sprintf(`%d %s "%s" %s`, w, agent, w.String(), date.Format(http.TimeFormat))
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:83
	// _ = "end of CoverTab[184011]"
}

func (w Warning) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:86
	_go_fuzz_dep_.CoverTab[184014]++
																switch w {
	case WarningResponseIsStale:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:88
		_go_fuzz_dep_.CoverTab[184016]++
																	return "Response is Stale"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:89
		// _ = "end of CoverTab[184016]"
	case WarningRevalidationFailed:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:90
		_go_fuzz_dep_.CoverTab[184017]++
																	return "Revalidation Failed"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:91
		// _ = "end of CoverTab[184017]"
	case WarningDisconnectedOperation:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:92
		_go_fuzz_dep_.CoverTab[184018]++
																	return "Disconnected Operation"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:93
		// _ = "end of CoverTab[184018]"
	case WarningHeuristicExpiration:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:94
		_go_fuzz_dep_.CoverTab[184019]++
																	return "Heuristic Expiration"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:95
		// _ = "end of CoverTab[184019]"
	case WarningMiscellaneousWarning:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:96
		_go_fuzz_dep_.CoverTab[184020]++

																	return "Miscellaneous Warning"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:98
		// _ = "end of CoverTab[184020]"
	case WarningTransformationApplied:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:99
			_go_fuzz_dep_.CoverTab[184021]++
																		return "Transformation Applied"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:100
		// _ = "end of CoverTab[184021]"
	case WarningMiscellaneousPersistentWarning:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:101
		_go_fuzz_dep_.CoverTab[184022]++

																		return "Miscellaneous Persistent Warning"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:103
		// _ = "end of CoverTab[184022]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:103
	default:
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:103
		_go_fuzz_dep_.CoverTab[184023]++
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:103
		// _ = "end of CoverTab[184023]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:104
	// _ = "end of CoverTab[184014]"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:104
	_go_fuzz_dep_.CoverTab[184015]++

																	panic(w)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:106
	// _ = "end of CoverTab[184015]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:107
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/cacheobject/warning.go:107
var _ = _go_fuzz_dep_.CoverTab
