//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
package oidc

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:1
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pquerna/cachecontrol"
	"golang.org/x/net/context"
	jose "gopkg.in/square/go-jose.v1"
)

// No matter what insist on caching keys. This is so our request code can be
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:19
// asynchronous from matching keys. If the request code retrieved keys that
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:19
// expired immediately, the goroutine to match a JWT to a key would always see
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:19
// expired keys.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:19
//
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:19
// TODO(ericchiang): Review this logic.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:25
var minCache = 2 * time.Minute

type cachedKeys struct {
	keys	map[string]jose.JsonWebKey	// immutable
	expiry	time.Time
}

type remoteKeySet struct {
	client	*http.Client

	// "jwks_uri" from discovery.
	keysURL	string

	// The value is always of type *cachedKeys.
	//
	// To ensure consistency always call keyCache.Store when holding cond.L.
	keyCache	atomic.Value

	// cond.L guards all following fields. sync.Cond is used in place of a mutex
	// so multiple processes can wait on a single request to update keys.
	cond	sync.Cond
	// Is there an existing request to get the remote keys?
	inflight	bool
	// If the last attempt to refresh keys failed, the error will be saved here.
	//
	// TODO(ericchiang): If a routine sets this before calling cond.Broadcast(),
	// there's no guarentee that a routine calling cond.Wait() will actual see
	// the error called by the previous routine. Since Broadcast() unlocks
	// cond.L and Wait() must reacquire the lock, other routines waiting on the
	// lock might acquire it first. Maybe just log the error?
	lastErr	error
}

func newRemoteKeySet(ctx context.Context, jwksURL string) *remoteKeySet {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:58
	_go_fuzz_dep_.CoverTab[186771]++
														r := &remoteKeySet{
		client:		contextClient(ctx),
		keysURL:	jwksURL,
		cond:		sync.Cond{L: new(sync.Mutex)},
	}
														return r
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:64
	// _ = "end of CoverTab[186771]"
}

func (r *remoteKeySet) verifyJWT(jwt string) (payload []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:67
	_go_fuzz_dep_.CoverTab[186772]++
														jws, err := jose.ParseSigned(jwt)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:69
		_go_fuzz_dep_.CoverTab[186776]++
															return nil, fmt.Errorf("parsing jwt: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:70
		// _ = "end of CoverTab[186776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:71
		_go_fuzz_dep_.CoverTab[186777]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:71
		// _ = "end of CoverTab[186777]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:71
	// _ = "end of CoverTab[186772]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:71
	_go_fuzz_dep_.CoverTab[186773]++
														keyIDs := make([]string, len(jws.Signatures))
														for i, signature := range jws.Signatures {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:73
		_go_fuzz_dep_.CoverTab[186778]++
															keyIDs[i] = signature.Header.KeyID
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:74
		// _ = "end of CoverTab[186778]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:75
	// _ = "end of CoverTab[186773]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:75
	_go_fuzz_dep_.CoverTab[186774]++
														key, err := r.getKey(keyIDs)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:77
		_go_fuzz_dep_.CoverTab[186779]++
															return nil, fmt.Errorf("oidc: %s", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:78
		// _ = "end of CoverTab[186779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:79
		_go_fuzz_dep_.CoverTab[186780]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:79
		// _ = "end of CoverTab[186780]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:79
	// _ = "end of CoverTab[186774]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:79
	_go_fuzz_dep_.CoverTab[186775]++
														return jws.Verify(key)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:80
	// _ = "end of CoverTab[186775]"
}

func (r *remoteKeySet) getKeyFromCache(keyIDs []string) (*jose.JsonWebKey, bool) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:83
	_go_fuzz_dep_.CoverTab[186781]++
														cachedKeys, ok := r.keyCache.Load().(*cachedKeys)
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:85
		_go_fuzz_dep_.CoverTab[186785]++
															return nil, false
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:86
		// _ = "end of CoverTab[186785]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:87
		_go_fuzz_dep_.CoverTab[186786]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:87
		// _ = "end of CoverTab[186786]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:87
	// _ = "end of CoverTab[186781]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:87
	_go_fuzz_dep_.CoverTab[186782]++
														if time.Now().After(cachedKeys.expiry) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:88
		_go_fuzz_dep_.CoverTab[186787]++
															return nil, false
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:89
		// _ = "end of CoverTab[186787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:90
		_go_fuzz_dep_.CoverTab[186788]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:90
		// _ = "end of CoverTab[186788]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:90
	// _ = "end of CoverTab[186782]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:90
	_go_fuzz_dep_.CoverTab[186783]++
														for _, keyID := range keyIDs {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:91
		_go_fuzz_dep_.CoverTab[186789]++
															if key, ok := cachedKeys.keys[keyID]; ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:92
			_go_fuzz_dep_.CoverTab[186790]++
																return &key, true
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:93
			// _ = "end of CoverTab[186790]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:94
			_go_fuzz_dep_.CoverTab[186791]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:94
			// _ = "end of CoverTab[186791]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:94
		// _ = "end of CoverTab[186789]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:95
	// _ = "end of CoverTab[186783]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:95
	_go_fuzz_dep_.CoverTab[186784]++
														return nil, false
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:96
	// _ = "end of CoverTab[186784]"
}

func (r *remoteKeySet) getKey(keyIDs []string) (*jose.JsonWebKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:99
	_go_fuzz_dep_.CoverTab[186792]++

														if key, ok := r.getKeyFromCache(keyIDs); ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:101
		_go_fuzz_dep_.CoverTab[186798]++
															return key, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:102
		// _ = "end of CoverTab[186798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:103
		_go_fuzz_dep_.CoverTab[186799]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:103
		// _ = "end of CoverTab[186799]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:103
	// _ = "end of CoverTab[186792]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:103
	_go_fuzz_dep_.CoverTab[186793]++

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:106
	r.cond.L.Lock()
														defer r.cond.L.Unlock()

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:110
	if key, ok := r.getKeyFromCache(keyIDs); ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:110
		_go_fuzz_dep_.CoverTab[186800]++
															return key, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:111
		// _ = "end of CoverTab[186800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:112
		_go_fuzz_dep_.CoverTab[186801]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:112
		// _ = "end of CoverTab[186801]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:112
	// _ = "end of CoverTab[186793]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:112
	_go_fuzz_dep_.CoverTab[186794]++

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:116
	if !r.inflight {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:116
		_go_fuzz_dep_.CoverTab[186802]++

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:119
		r.inflight = true
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:119
		_curRoutineNum164_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:119
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum164_)
															go func() {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
			_go_fuzz_dep_.CoverTab[186803]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
				_go_fuzz_dep_.CoverTab[186805]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum164_)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
				// _ = "end of CoverTab[186805]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:120
			}()
																newKeys, newExpiry, err := requestKeys(r.client, r.keysURL)

																r.cond.L.Lock()
																defer r.cond.L.Unlock()

																r.inflight = false
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:127
				_go_fuzz_dep_.CoverTab[186806]++
																	r.lastErr = err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:128
				// _ = "end of CoverTab[186806]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:129
				_go_fuzz_dep_.CoverTab[186807]++
																	r.keyCache.Store(&cachedKeys{newKeys, newExpiry})
																	r.lastErr = nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:131
				// _ = "end of CoverTab[186807]"
			}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:132
			// _ = "end of CoverTab[186803]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:132
			_go_fuzz_dep_.CoverTab[186804]++

																r.cond.Broadcast()
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:134
			// _ = "end of CoverTab[186804]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:135
		// _ = "end of CoverTab[186802]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:136
		_go_fuzz_dep_.CoverTab[186808]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:136
		// _ = "end of CoverTab[186808]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:136
	// _ = "end of CoverTab[186794]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:136
	_go_fuzz_dep_.CoverTab[186795]++

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:140
	r.cond.Wait()

	if key, ok := r.getKeyFromCache(keyIDs); ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:142
		_go_fuzz_dep_.CoverTab[186809]++
															return key, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:143
		// _ = "end of CoverTab[186809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:144
		_go_fuzz_dep_.CoverTab[186810]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:144
		// _ = "end of CoverTab[186810]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:144
	// _ = "end of CoverTab[186795]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:144
	_go_fuzz_dep_.CoverTab[186796]++
														if r.lastErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:145
		_go_fuzz_dep_.CoverTab[186811]++
															return nil, r.lastErr
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:146
		// _ = "end of CoverTab[186811]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:147
		_go_fuzz_dep_.CoverTab[186812]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:147
		// _ = "end of CoverTab[186812]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:147
	// _ = "end of CoverTab[186796]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:147
	_go_fuzz_dep_.CoverTab[186797]++
														return nil, errors.New("no signing keys can validate the signature")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:148
	// _ = "end of CoverTab[186797]"
}

func requestKeys(client *http.Client, keysURL string) (map[string]jose.JsonWebKey, time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:151
	_go_fuzz_dep_.CoverTab[186813]++
														req, err := http.NewRequest("GET", keysURL, nil)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:153
		_go_fuzz_dep_.CoverTab[186821]++
															return nil, time.Time{}, fmt.Errorf("can't create request: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:154
		// _ = "end of CoverTab[186821]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:155
		_go_fuzz_dep_.CoverTab[186822]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:155
		// _ = "end of CoverTab[186822]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:155
	// _ = "end of CoverTab[186813]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:155
	_go_fuzz_dep_.CoverTab[186814]++
														resp, err := client.Do(req)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:157
		_go_fuzz_dep_.CoverTab[186823]++
															return nil, time.Time{}, fmt.Errorf("can't GET new keys %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:158
		// _ = "end of CoverTab[186823]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:159
		_go_fuzz_dep_.CoverTab[186824]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:159
		// _ = "end of CoverTab[186824]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:159
	// _ = "end of CoverTab[186814]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:159
	_go_fuzz_dep_.CoverTab[186815]++
														defer resp.Body.Close()

														body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:163
		_go_fuzz_dep_.CoverTab[186825]++
															return nil, time.Time{}, fmt.Errorf("can't fetch new keys: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:164
		// _ = "end of CoverTab[186825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:165
		_go_fuzz_dep_.CoverTab[186826]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:165
		// _ = "end of CoverTab[186826]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:165
	// _ = "end of CoverTab[186815]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:165
	_go_fuzz_dep_.CoverTab[186816]++
														if resp.StatusCode != http.StatusOK {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:166
		_go_fuzz_dep_.CoverTab[186827]++
															return nil, time.Time{}, fmt.Errorf("can't fetch new keys: %s %s", resp.Status, body)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:167
		// _ = "end of CoverTab[186827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:168
		_go_fuzz_dep_.CoverTab[186828]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:168
		// _ = "end of CoverTab[186828]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:168
	// _ = "end of CoverTab[186816]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:168
	_go_fuzz_dep_.CoverTab[186817]++

														var keySet jose.JsonWebKeySet
														if err := json.Unmarshal(body, &keySet); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:171
		_go_fuzz_dep_.CoverTab[186829]++
															return nil, time.Time{}, fmt.Errorf("can't decode keys: %v %s", err, body)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:172
		// _ = "end of CoverTab[186829]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:173
		_go_fuzz_dep_.CoverTab[186830]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:173
		// _ = "end of CoverTab[186830]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:173
	// _ = "end of CoverTab[186817]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:173
	_go_fuzz_dep_.CoverTab[186818]++

														keys := make(map[string]jose.JsonWebKey, len(keySet.Keys))
														for _, key := range keySet.Keys {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:176
		_go_fuzz_dep_.CoverTab[186831]++
															keys[key.KeyID] = key
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:177
		// _ = "end of CoverTab[186831]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:178
	// _ = "end of CoverTab[186818]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:178
	_go_fuzz_dep_.CoverTab[186819]++

														minExpiry := time.Now().Add(minCache)

														if _, expiry, err := cachecontrol.CachableResponse(req, resp, cachecontrol.Options{}); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:182
		_go_fuzz_dep_.CoverTab[186832]++
															if minExpiry.Before(expiry) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:183
			_go_fuzz_dep_.CoverTab[186833]++
																return keys, expiry, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:184
			// _ = "end of CoverTab[186833]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:185
			_go_fuzz_dep_.CoverTab[186834]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:185
			// _ = "end of CoverTab[186834]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:185
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:185
		// _ = "end of CoverTab[186832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:186
		_go_fuzz_dep_.CoverTab[186835]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:186
		// _ = "end of CoverTab[186835]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:186
	// _ = "end of CoverTab[186819]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:186
	_go_fuzz_dep_.CoverTab[186820]++
														return keys, minExpiry, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:187
	// _ = "end of CoverTab[186820]"
}

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:188
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/jwks.go:188
var _ = _go_fuzz_dep_.CoverTab
