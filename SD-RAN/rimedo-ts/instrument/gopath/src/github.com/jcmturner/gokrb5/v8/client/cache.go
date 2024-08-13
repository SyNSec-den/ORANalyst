//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:1
)

import (
	"encoding/json"
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Cache for service tickets held by the client.
type Cache struct {
	Entries	map[string]CacheEntry
	mux	sync.RWMutex
}

// CacheEntry holds details for a cache entry.
type CacheEntry struct {
	SPN		string
	Ticket		messages.Ticket	`json:"-"`
	AuthTime	time.Time
	StartTime	time.Time
	EndTime		time.Time
	RenewTill	time.Time
	SessionKey	types.EncryptionKey	`json:"-"`
}

// NewCache creates a new client ticket cache instance.
func NewCache() *Cache {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:32
	_go_fuzz_dep_.CoverTab[88377]++
												return &Cache{
		Entries: map[string]CacheEntry{},
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:35
	// _ = "end of CoverTab[88377]"
}

// getEntry returns a cache entry that matches the SPN.
func (c *Cache) getEntry(spn string) (CacheEntry, bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:39
	_go_fuzz_dep_.CoverTab[88378]++
												c.mux.RLock()
												defer c.mux.RUnlock()
												e, ok := (*c).Entries[spn]
												return e, ok
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:43
	// _ = "end of CoverTab[88378]"
}

// JSON returns information about the cached service tickets in a JSON format.
func (c *Cache) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:47
	_go_fuzz_dep_.CoverTab[88379]++
												c.mux.RLock()
												defer c.mux.RUnlock()
												var es []CacheEntry
												keys := make([]string, 0, len(c.Entries))
												for k := range c.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:52
		_go_fuzz_dep_.CoverTab[88383]++
													keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:53
		// _ = "end of CoverTab[88383]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:54
	// _ = "end of CoverTab[88379]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:54
	_go_fuzz_dep_.CoverTab[88380]++
												sort.Strings(keys)
												for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:56
		_go_fuzz_dep_.CoverTab[88384]++
													es = append(es, c.Entries[k])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:57
		// _ = "end of CoverTab[88384]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:58
	// _ = "end of CoverTab[88380]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:58
	_go_fuzz_dep_.CoverTab[88381]++
												b, err := json.MarshalIndent(&es, "", "  ")
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:60
		_go_fuzz_dep_.CoverTab[88385]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:61
		// _ = "end of CoverTab[88385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:62
		_go_fuzz_dep_.CoverTab[88386]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:62
		// _ = "end of CoverTab[88386]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:62
	// _ = "end of CoverTab[88381]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:62
	_go_fuzz_dep_.CoverTab[88382]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:63
	// _ = "end of CoverTab[88382]"
}

// addEntry adds a ticket to the cache.
func (c *Cache) addEntry(tkt messages.Ticket, authTime, startTime, endTime, renewTill time.Time, sessionKey types.EncryptionKey) CacheEntry {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:67
	_go_fuzz_dep_.CoverTab[88387]++
												spn := tkt.SName.PrincipalNameString()
												c.mux.Lock()
												defer c.mux.Unlock()
												(*c).Entries[spn] = CacheEntry{
		SPN:		spn,
		Ticket:		tkt,
		AuthTime:	authTime,
		StartTime:	startTime,
		EndTime:	endTime,
		RenewTill:	renewTill,
		SessionKey:	sessionKey,
	}
												return c.Entries[spn]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:80
	// _ = "end of CoverTab[88387]"
}

// clear deletes all the cache entries
func (c *Cache) clear() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:84
	_go_fuzz_dep_.CoverTab[88388]++
												c.mux.Lock()
												defer c.mux.Unlock()
												for k := range c.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:87
		_go_fuzz_dep_.CoverTab[88389]++
													delete(c.Entries, k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:88
		// _ = "end of CoverTab[88389]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:89
	// _ = "end of CoverTab[88388]"
}

// RemoveEntry removes the cache entry for the defined SPN.
func (c *Cache) RemoveEntry(spn string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:93
	_go_fuzz_dep_.CoverTab[88390]++
												c.mux.Lock()
												defer c.mux.Unlock()
												delete(c.Entries, spn)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:96
	// _ = "end of CoverTab[88390]"
}

// GetCachedTicket returns a ticket from the cache for the SPN.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:99
// Only a ticket that is currently valid will be returned.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:101
func (cl *Client) GetCachedTicket(spn string) (messages.Ticket, types.EncryptionKey, bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:101
	_go_fuzz_dep_.CoverTab[88391]++
												if e, ok := cl.cache.getEntry(spn); ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:102
		_go_fuzz_dep_.CoverTab[88393]++

													if time.Now().UTC().After(e.StartTime) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:104
			_go_fuzz_dep_.CoverTab[88394]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:104
			return time.Now().UTC().Before(e.EndTime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:104
			// _ = "end of CoverTab[88394]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:104
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:104
			_go_fuzz_dep_.CoverTab[88395]++
														cl.Log("ticket received from cache for %s", spn)
														return e.Ticket, e.SessionKey, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:106
			// _ = "end of CoverTab[88395]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:107
			_go_fuzz_dep_.CoverTab[88396]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:107
			if time.Now().UTC().Before(e.RenewTill) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:107
				_go_fuzz_dep_.CoverTab[88397]++
															e, err := cl.renewTicket(e)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:109
					_go_fuzz_dep_.CoverTab[88399]++
																return e.Ticket, e.SessionKey, false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:110
					// _ = "end of CoverTab[88399]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:111
					_go_fuzz_dep_.CoverTab[88400]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:111
					// _ = "end of CoverTab[88400]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:111
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:111
				// _ = "end of CoverTab[88397]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:111
				_go_fuzz_dep_.CoverTab[88398]++
															return e.Ticket, e.SessionKey, true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:112
				// _ = "end of CoverTab[88398]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
				_go_fuzz_dep_.CoverTab[88401]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
				// _ = "end of CoverTab[88401]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
			// _ = "end of CoverTab[88396]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:113
		// _ = "end of CoverTab[88393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:114
		_go_fuzz_dep_.CoverTab[88402]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:114
		// _ = "end of CoverTab[88402]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:114
	// _ = "end of CoverTab[88391]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:114
	_go_fuzz_dep_.CoverTab[88392]++
												var tkt messages.Ticket
												var key types.EncryptionKey
												return tkt, key, false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:117
	// _ = "end of CoverTab[88392]"
}

// renewTicket renews a cache entry ticket.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:120
// To renew from outside the client package use GetCachedTicket
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:122
func (cl *Client) renewTicket(e CacheEntry) (CacheEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:122
	_go_fuzz_dep_.CoverTab[88403]++
												spn := e.Ticket.SName
												_, _, err := cl.TGSREQGenerateAndExchange(spn, e.Ticket.Realm, e.Ticket, e.SessionKey, true)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:125
		_go_fuzz_dep_.CoverTab[88406]++
													return e, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:126
		// _ = "end of CoverTab[88406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:127
		_go_fuzz_dep_.CoverTab[88407]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:127
		// _ = "end of CoverTab[88407]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:127
	// _ = "end of CoverTab[88403]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:127
	_go_fuzz_dep_.CoverTab[88404]++
												e, ok := cl.cache.getEntry(e.Ticket.SName.PrincipalNameString())
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:129
		_go_fuzz_dep_.CoverTab[88408]++
													return e, errors.New("ticket was not added to cache")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:130
		// _ = "end of CoverTab[88408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:131
		_go_fuzz_dep_.CoverTab[88409]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:131
		// _ = "end of CoverTab[88409]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:131
	// _ = "end of CoverTab[88404]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:131
	_go_fuzz_dep_.CoverTab[88405]++
												cl.Log("ticket renewed for %s (EndTime: %v)", spn.PrincipalNameString(), e.EndTime)
												return e, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:133
	// _ = "end of CoverTab[88405]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/cache.go:134
var _ = _go_fuzz_dep_.CoverTab
