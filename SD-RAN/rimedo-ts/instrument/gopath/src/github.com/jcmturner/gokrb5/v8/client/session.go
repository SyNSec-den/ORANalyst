//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:1
)

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

// sessions hold TGTs and are keyed on the realm name
type sessions struct {
	Entries	map[string]*session
	mux	sync.RWMutex
}

// destroy erases all sessions
func (s *sessions) destroy() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:24
	_go_fuzz_dep_.CoverTab[88686]++
												s.mux.Lock()
												defer s.mux.Unlock()
												for k, e := range s.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:27
		_go_fuzz_dep_.CoverTab[88687]++
													e.destroy()
													delete(s.Entries, k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:29
		// _ = "end of CoverTab[88687]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:30
	// _ = "end of CoverTab[88686]"
}

// update replaces a session with the one provided or adds it as a new one
func (s *sessions) update(sess *session) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:34
	_go_fuzz_dep_.CoverTab[88688]++
												s.mux.Lock()
												defer s.mux.Unlock()

												if i, ok := s.Entries[sess.realm]; ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:38
		_go_fuzz_dep_.CoverTab[88690]++
													if i != sess {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:39
			_go_fuzz_dep_.CoverTab[88691]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:42
			i.mux.Lock()
														defer i.mux.Unlock()
														i.cancel <- true
														s.Entries[sess.realm] = sess
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:46
			// _ = "end of CoverTab[88691]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:47
			_go_fuzz_dep_.CoverTab[88692]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:47
			// _ = "end of CoverTab[88692]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:47
		// _ = "end of CoverTab[88690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:48
		_go_fuzz_dep_.CoverTab[88693]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:48
		// _ = "end of CoverTab[88693]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:48
	// _ = "end of CoverTab[88688]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:48
	_go_fuzz_dep_.CoverTab[88689]++

												s.Entries[sess.realm] = sess
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:50
	// _ = "end of CoverTab[88689]"
}

// get returns the session for the realm specified
func (s *sessions) get(realm string) (*session, bool) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:54
	_go_fuzz_dep_.CoverTab[88694]++
												s.mux.RLock()
												defer s.mux.RUnlock()
												sess, ok := s.Entries[realm]
												return sess, ok
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:58
	// _ = "end of CoverTab[88694]"
}

// session holds the TGT details for a realm
type session struct {
	realm			string
	authTime		time.Time
	endTime			time.Time
	renewTill		time.Time
	tgt			messages.Ticket
	sessionKey		types.EncryptionKey
	sessionKeyExpiration	time.Time
	cancel			chan bool
	mux			sync.RWMutex
}

// jsonSession is used to enable marshaling some information of a session in a JSON format
type jsonSession struct {
	Realm			string
	AuthTime		time.Time
	EndTime			time.Time
	RenewTill		time.Time
	SessionKeyExpiration	time.Time
}

// AddSession adds a session for a realm with a TGT to the client's session cache.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:83
// A goroutine is started to automatically renew the TGT before expiry.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:85
func (cl *Client) addSession(tgt messages.Ticket, dep messages.EncKDCRepPart) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:85
	_go_fuzz_dep_.CoverTab[88695]++
												if strings.ToLower(tgt.SName.NameString[0]) != "krbtgt" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:86
		_go_fuzz_dep_.CoverTab[88697]++

													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:88
		// _ = "end of CoverTab[88697]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:89
		_go_fuzz_dep_.CoverTab[88698]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:89
		// _ = "end of CoverTab[88698]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:89
	// _ = "end of CoverTab[88695]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:89
	_go_fuzz_dep_.CoverTab[88696]++
												realm := tgt.SName.NameString[len(tgt.SName.NameString)-1]
												s := &session{
		realm:			realm,
		authTime:		dep.AuthTime,
		endTime:		dep.EndTime,
		renewTill:		dep.RenewTill,
		tgt:			tgt,
		sessionKey:		dep.Key,
		sessionKeyExpiration:	dep.KeyExpiration,
	}
												cl.sessions.update(s)
												cl.enableAutoSessionRenewal(s)
												cl.Log("TGT session added for %s (EndTime: %v)", realm, dep.EndTime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:102
	// _ = "end of CoverTab[88696]"
}

// update overwrites the session details with those from the TGT and decrypted encPart
func (s *session) update(tgt messages.Ticket, dep messages.EncKDCRepPart) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:106
	_go_fuzz_dep_.CoverTab[88699]++
												s.mux.Lock()
												defer s.mux.Unlock()
												s.authTime = dep.AuthTime
												s.endTime = dep.EndTime
												s.renewTill = dep.RenewTill
												s.tgt = tgt
												s.sessionKey = dep.Key
												s.sessionKeyExpiration = dep.KeyExpiration
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:114
	// _ = "end of CoverTab[88699]"
}

// destroy will cancel any auto renewal of the session and set the expiration times to the current time
func (s *session) destroy() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:118
	_go_fuzz_dep_.CoverTab[88700]++
												s.mux.Lock()
												defer s.mux.Unlock()
												if s.cancel != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:121
		_go_fuzz_dep_.CoverTab[88702]++
													s.cancel <- true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:122
		// _ = "end of CoverTab[88702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:123
		_go_fuzz_dep_.CoverTab[88703]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:123
		// _ = "end of CoverTab[88703]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:123
	// _ = "end of CoverTab[88700]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:123
	_go_fuzz_dep_.CoverTab[88701]++
												s.endTime = time.Now().UTC()
												s.renewTill = s.endTime
												s.sessionKeyExpiration = s.endTime
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:126
	// _ = "end of CoverTab[88701]"
}

// valid informs if the TGT is still within the valid time window
func (s *session) valid() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:130
	_go_fuzz_dep_.CoverTab[88704]++
												s.mux.RLock()
												defer s.mux.RUnlock()
												t := time.Now().UTC()
												if t.Before(s.endTime) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:134
		_go_fuzz_dep_.CoverTab[88706]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:134
		return s.authTime.Before(t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:134
		// _ = "end of CoverTab[88706]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:134
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:134
		_go_fuzz_dep_.CoverTab[88707]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:135
		// _ = "end of CoverTab[88707]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:136
		_go_fuzz_dep_.CoverTab[88708]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:136
		// _ = "end of CoverTab[88708]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:136
	// _ = "end of CoverTab[88704]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:136
	_go_fuzz_dep_.CoverTab[88705]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:137
	// _ = "end of CoverTab[88705]"
}

// tgtDetails is a thread safe way to get the session's realm, TGT and session key values
func (s *session) tgtDetails() (string, messages.Ticket, types.EncryptionKey) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:141
	_go_fuzz_dep_.CoverTab[88709]++
												s.mux.RLock()
												defer s.mux.RUnlock()
												return s.realm, s.tgt, s.sessionKey
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:144
	// _ = "end of CoverTab[88709]"
}

// timeDetails is a thread safe way to get the session's validity time values
func (s *session) timeDetails() (string, time.Time, time.Time, time.Time, time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:148
	_go_fuzz_dep_.CoverTab[88710]++
												s.mux.RLock()
												defer s.mux.RUnlock()
												return s.realm, s.authTime, s.endTime, s.renewTill, s.sessionKeyExpiration
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:151
	// _ = "end of CoverTab[88710]"
}

// JSON return information about the held sessions in a JSON format.
func (s *sessions) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:155
	_go_fuzz_dep_.CoverTab[88711]++
												s.mux.RLock()
												defer s.mux.RUnlock()
												var js []jsonSession
												keys := make([]string, 0, len(s.Entries))
												for k := range s.Entries {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:160
		_go_fuzz_dep_.CoverTab[88715]++
													keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:161
		// _ = "end of CoverTab[88715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:162
	// _ = "end of CoverTab[88711]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:162
	_go_fuzz_dep_.CoverTab[88712]++
												sort.Strings(keys)
												for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:164
		_go_fuzz_dep_.CoverTab[88716]++
													r, at, et, rt, kt := s.Entries[k].timeDetails()
													j := jsonSession{
			Realm:			r,
			AuthTime:		at,
			EndTime:		et,
			RenewTill:		rt,
			SessionKeyExpiration:	kt,
		}
													js = append(js, j)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:173
		// _ = "end of CoverTab[88716]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:174
	// _ = "end of CoverTab[88712]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:174
	_go_fuzz_dep_.CoverTab[88713]++
												b, err := json.MarshalIndent(js, "", "  ")
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:176
		_go_fuzz_dep_.CoverTab[88717]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:177
		// _ = "end of CoverTab[88717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:178
		_go_fuzz_dep_.CoverTab[88718]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:178
		// _ = "end of CoverTab[88718]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:178
	// _ = "end of CoverTab[88713]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:178
	_go_fuzz_dep_.CoverTab[88714]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:179
	// _ = "end of CoverTab[88714]"
}

// enableAutoSessionRenewal turns on the automatic renewal for the client's TGT session.
func (cl *Client) enableAutoSessionRenewal(s *session) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:183
	_go_fuzz_dep_.CoverTab[88719]++
												var timer *time.Timer
												s.mux.Lock()
												s.cancel = make(chan bool, 1)
												s.mux.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:187
	_curRoutineNum101_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:187
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum101_)
												go func(s *session) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
		_go_fuzz_dep_.CoverTab[88720]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
			_go_fuzz_dep_.CoverTab[88721]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum101_)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
			// _ = "end of CoverTab[88721]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:188
		}()
													for {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:189
			_go_fuzz_dep_.CoverTab[88722]++
														s.mux.RLock()
														w := (s.endTime.Sub(time.Now().UTC()) * 5) / 6
														s.mux.RUnlock()
														if w < 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:193
				_go_fuzz_dep_.CoverTab[88724]++
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:194
				// _ = "end of CoverTab[88724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:195
				_go_fuzz_dep_.CoverTab[88725]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:195
				// _ = "end of CoverTab[88725]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:195
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:195
			// _ = "end of CoverTab[88722]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:195
			_go_fuzz_dep_.CoverTab[88723]++
														timer = time.NewTimer(w)
														select {
			case <-timer.C:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:198
				_go_fuzz_dep_.CoverTab[88726]++
															renewal, err := cl.refreshSession(s)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:200
					_go_fuzz_dep_.CoverTab[88729]++
																cl.Log("error refreshing session: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:201
					// _ = "end of CoverTab[88729]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:202
					_go_fuzz_dep_.CoverTab[88730]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:202
					// _ = "end of CoverTab[88730]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:202
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:202
				// _ = "end of CoverTab[88726]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:202
				_go_fuzz_dep_.CoverTab[88727]++
															if !renewal && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:203
					_go_fuzz_dep_.CoverTab[88731]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:203
					return err == nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:203
					// _ = "end of CoverTab[88731]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:203
				}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:203
					_go_fuzz_dep_.CoverTab[88732]++

																return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:205
					// _ = "end of CoverTab[88732]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:206
					_go_fuzz_dep_.CoverTab[88733]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:206
					// _ = "end of CoverTab[88733]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:206
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:206
				// _ = "end of CoverTab[88727]"
			case <-s.cancel:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:207
				_go_fuzz_dep_.CoverTab[88728]++

															timer.Stop()
															return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:210
				// _ = "end of CoverTab[88728]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:211
			// _ = "end of CoverTab[88723]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:212
		// _ = "end of CoverTab[88720]"
	}(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:213
	// _ = "end of CoverTab[88719]"
}

// renewTGT renews the client's TGT session.
func (cl *Client) renewTGT(s *session) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:217
	_go_fuzz_dep_.CoverTab[88734]++
												realm, tgt, skey := s.tgtDetails()
												spn := types.PrincipalName{
		NameType:	nametype.KRB_NT_SRV_INST,
		NameString:	[]string{"krbtgt", realm},
	}
	_, tgsRep, err := cl.TGSREQGenerateAndExchange(spn, cl.Credentials.Domain(), tgt, skey, true)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:224
		_go_fuzz_dep_.CoverTab[88736]++
													return krberror.Errorf(err, krberror.KRBMsgError, "error renewing TGT for %s", realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:225
		// _ = "end of CoverTab[88736]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:226
		_go_fuzz_dep_.CoverTab[88737]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:226
		// _ = "end of CoverTab[88737]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:226
	// _ = "end of CoverTab[88734]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:226
	_go_fuzz_dep_.CoverTab[88735]++
												s.update(tgsRep.Ticket, tgsRep.DecryptedEncPart)
												cl.sessions.update(s)
												cl.Log("TGT session renewed for %s (EndTime: %v)", realm, tgsRep.DecryptedEncPart.EndTime)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:230
	// _ = "end of CoverTab[88735]"
}

// refreshSession updates either through renewal or creating a new login.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:233
// The boolean indicates if the update was a renewal.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:235
func (cl *Client) refreshSession(s *session) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:235
	_go_fuzz_dep_.CoverTab[88738]++
												s.mux.RLock()
												realm := s.realm
												renewTill := s.renewTill
												s.mux.RUnlock()
												cl.Log("refreshing TGT session for %s", realm)
												if time.Now().UTC().Before(renewTill) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:241
		_go_fuzz_dep_.CoverTab[88740]++
													err := cl.renewTGT(s)
													return true, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:243
		// _ = "end of CoverTab[88740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:244
		_go_fuzz_dep_.CoverTab[88741]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:244
		// _ = "end of CoverTab[88741]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:244
	// _ = "end of CoverTab[88738]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:244
	_go_fuzz_dep_.CoverTab[88739]++
												err := cl.realmLogin(realm)
												return false, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:246
	// _ = "end of CoverTab[88739]"
}

// ensureValidSession makes sure there is a valid session for the realm
func (cl *Client) ensureValidSession(realm string) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:250
	_go_fuzz_dep_.CoverTab[88742]++
												s, ok := cl.sessions.get(realm)
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:252
		_go_fuzz_dep_.CoverTab[88744]++
													s.mux.RLock()
													d := s.endTime.Sub(s.authTime) / 6
													if s.endTime.Sub(time.Now().UTC()) > d {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:255
			_go_fuzz_dep_.CoverTab[88746]++
														s.mux.RUnlock()
														return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:257
			// _ = "end of CoverTab[88746]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:258
			_go_fuzz_dep_.CoverTab[88747]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:258
			// _ = "end of CoverTab[88747]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:258
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:258
		// _ = "end of CoverTab[88744]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:258
		_go_fuzz_dep_.CoverTab[88745]++
													s.mux.RUnlock()
													_, err := cl.refreshSession(s)
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:261
		// _ = "end of CoverTab[88745]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:262
		_go_fuzz_dep_.CoverTab[88748]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:262
		// _ = "end of CoverTab[88748]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:262
	// _ = "end of CoverTab[88742]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:262
	_go_fuzz_dep_.CoverTab[88743]++
												return cl.realmLogin(realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:263
	// _ = "end of CoverTab[88743]"
}

// sessionTGTDetails is a thread safe way to get the TGT and session key values for a realm
func (cl *Client) sessionTGT(realm string) (tgt messages.Ticket, sessionKey types.EncryptionKey, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:267
	_go_fuzz_dep_.CoverTab[88749]++
												err = cl.ensureValidSession(realm)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:269
		_go_fuzz_dep_.CoverTab[88752]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:270
		// _ = "end of CoverTab[88752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:271
		_go_fuzz_dep_.CoverTab[88753]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:271
		// _ = "end of CoverTab[88753]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:271
	// _ = "end of CoverTab[88749]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:271
	_go_fuzz_dep_.CoverTab[88750]++
												s, ok := cl.sessions.get(realm)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:273
		_go_fuzz_dep_.CoverTab[88754]++
													err = fmt.Errorf("could not find TGT session for %s", realm)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:275
		// _ = "end of CoverTab[88754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:276
		_go_fuzz_dep_.CoverTab[88755]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:276
		// _ = "end of CoverTab[88755]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:276
	// _ = "end of CoverTab[88750]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:276
	_go_fuzz_dep_.CoverTab[88751]++
												_, tgt, sessionKey = s.tgtDetails()
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:278
	// _ = "end of CoverTab[88751]"
}

// sessionTimes provides the timing information with regards to a session for the realm specified.
func (cl *Client) sessionTimes(realm string) (authTime, endTime, renewTime, sessionExp time.Time, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:282
	_go_fuzz_dep_.CoverTab[88756]++
												s, ok := cl.sessions.get(realm)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:284
		_go_fuzz_dep_.CoverTab[88758]++
													err = fmt.Errorf("could not find TGT session for %s", realm)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:286
		// _ = "end of CoverTab[88758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:287
		_go_fuzz_dep_.CoverTab[88759]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:287
		// _ = "end of CoverTab[88759]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:287
	// _ = "end of CoverTab[88756]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:287
	_go_fuzz_dep_.CoverTab[88757]++
												_, authTime, endTime, renewTime, sessionExp = s.timeDetails()
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:289
	// _ = "end of CoverTab[88757]"
}

// spnRealm resolves the realm name of a service principal name
func (cl *Client) spnRealm(spn types.PrincipalName) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:293
	_go_fuzz_dep_.CoverTab[88760]++
												return cl.Config.ResolveRealm(spn.NameString[len(spn.NameString)-1])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:294
	// _ = "end of CoverTab[88760]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:295
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/session.go:295
var _ = _go_fuzz_dep_.CoverTab
