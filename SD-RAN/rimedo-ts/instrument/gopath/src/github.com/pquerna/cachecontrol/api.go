//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
package cachecontrol

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
import (
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:18
)

import (
	"github.com/pquerna/cachecontrol/cacheobject"

	"net/http"
	"time"
)

type Options struct {
	// Set to True for a prviate cache, which is not shared amoung users (eg, in a browser)
	// Set to False for a "shared" cache, which is more common in a server context.
	PrivateCache bool
}

// Given an HTTP Request, the future Status Code, and an ResponseWriter,
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:33
// determine the possible reasons a response SHOULD NOT be cached.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:35
func CachableResponseWriter(req *http.Request,
	statusCode int,
	resp http.ResponseWriter,
	opts Options) ([]cacheobject.Reason, time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:38
	_go_fuzz_dep_.CoverTab[184024]++
														return cacheobject.UsingRequestResponse(req, statusCode, resp.Header(), opts.PrivateCache)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:39
	// _ = "end of CoverTab[184024]"
}

// Given an HTTP Request and Response, determine the possible reasons a response SHOULD NOT
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:42
// be cached.
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:44
func CachableResponse(req *http.Request,
	resp *http.Response,
	opts Options) ([]cacheobject.Reason, time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:46
	_go_fuzz_dep_.CoverTab[184025]++
														return cacheobject.UsingRequestResponse(req, resp.StatusCode, resp.Header, opts.PrivateCache)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:47
	// _ = "end of CoverTab[184025]"
}

//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pquerna/cachecontrol@v0.0.0-20180517163645-1555304b9b35/api.go:48
var _ = _go_fuzz_dep_.CoverTab
