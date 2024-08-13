// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/nss.go:5
package net

//line /snap/go/10455/src/net/nss.go:5
import (
//line /snap/go/10455/src/net/nss.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/nss.go:5
)
//line /snap/go/10455/src/net/nss.go:5
import (
//line /snap/go/10455/src/net/nss.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/nss.go:5
)

import (
	"errors"
	"internal/bytealg"
	"os"
	"sync"
	"time"
)

const (
	nssConfigPath = "/etc/nsswitch.conf"
)

var nssConfig nsswitchConfig

type nsswitchConfig struct {
	initOnce	sync.Once	// guards init of nsswitchConfig

	// ch is used as a semaphore that only allows one lookup at a
	// time to recheck nsswitch.conf
	ch		chan struct{}	// guards lastChecked and modTime
	lastChecked	time.Time	// last time nsswitch.conf was checked

	mu	sync.Mutex	// protects nssConf
	nssConf	*nssConf
}

func getSystemNSS() *nssConf {
//line /snap/go/10455/src/net/nss.go:33
	_go_fuzz_dep_.CoverTab[7567]++
						nssConfig.tryUpdate()
						nssConfig.mu.Lock()
						conf := nssConfig.nssConf
						nssConfig.mu.Unlock()
						return conf
//line /snap/go/10455/src/net/nss.go:38
	// _ = "end of CoverTab[7567]"
}

// init initializes conf and is only called via conf.initOnce.
func (conf *nsswitchConfig) init() {
	conf.nssConf = parseNSSConfFile("/etc/nsswitch.conf")
	conf.lastChecked = time.Now()
	conf.ch = make(chan struct{}, 1)
}

// tryUpdate tries to update conf.
func (conf *nsswitchConfig) tryUpdate() {
//line /snap/go/10455/src/net/nss.go:49
	_go_fuzz_dep_.CoverTab[7568]++
						conf.initOnce.Do(conf.init)

//line /snap/go/10455/src/net/nss.go:53
	if !conf.tryAcquireSema() {
//line /snap/go/10455/src/net/nss.go:53
		_go_fuzz_dep_.CoverTab[529325]++
//line /snap/go/10455/src/net/nss.go:53
		_go_fuzz_dep_.CoverTab[7573]++
							return
//line /snap/go/10455/src/net/nss.go:54
		// _ = "end of CoverTab[7573]"
	} else {
//line /snap/go/10455/src/net/nss.go:55
		_go_fuzz_dep_.CoverTab[529326]++
//line /snap/go/10455/src/net/nss.go:55
		_go_fuzz_dep_.CoverTab[7574]++
//line /snap/go/10455/src/net/nss.go:55
		// _ = "end of CoverTab[7574]"
//line /snap/go/10455/src/net/nss.go:55
	}
//line /snap/go/10455/src/net/nss.go:55
	// _ = "end of CoverTab[7568]"
//line /snap/go/10455/src/net/nss.go:55
	_go_fuzz_dep_.CoverTab[7569]++
						defer conf.releaseSema()

						now := time.Now()
						if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /snap/go/10455/src/net/nss.go:59
		_go_fuzz_dep_.CoverTab[529327]++
//line /snap/go/10455/src/net/nss.go:59
		_go_fuzz_dep_.CoverTab[7575]++
							return
//line /snap/go/10455/src/net/nss.go:60
		// _ = "end of CoverTab[7575]"
	} else {
//line /snap/go/10455/src/net/nss.go:61
		_go_fuzz_dep_.CoverTab[529328]++
//line /snap/go/10455/src/net/nss.go:61
		_go_fuzz_dep_.CoverTab[7576]++
//line /snap/go/10455/src/net/nss.go:61
		// _ = "end of CoverTab[7576]"
//line /snap/go/10455/src/net/nss.go:61
	}
//line /snap/go/10455/src/net/nss.go:61
	// _ = "end of CoverTab[7569]"
//line /snap/go/10455/src/net/nss.go:61
	_go_fuzz_dep_.CoverTab[7570]++
						conf.lastChecked = now

						var mtime time.Time
						if fi, err := os.Stat(nssConfigPath); err == nil {
//line /snap/go/10455/src/net/nss.go:65
		_go_fuzz_dep_.CoverTab[529329]++
//line /snap/go/10455/src/net/nss.go:65
		_go_fuzz_dep_.CoverTab[7577]++
							mtime = fi.ModTime()
//line /snap/go/10455/src/net/nss.go:66
		// _ = "end of CoverTab[7577]"
	} else {
//line /snap/go/10455/src/net/nss.go:67
		_go_fuzz_dep_.CoverTab[529330]++
//line /snap/go/10455/src/net/nss.go:67
		_go_fuzz_dep_.CoverTab[7578]++
//line /snap/go/10455/src/net/nss.go:67
		// _ = "end of CoverTab[7578]"
//line /snap/go/10455/src/net/nss.go:67
	}
//line /snap/go/10455/src/net/nss.go:67
	// _ = "end of CoverTab[7570]"
//line /snap/go/10455/src/net/nss.go:67
	_go_fuzz_dep_.CoverTab[7571]++
						if mtime.Equal(conf.nssConf.mtime) {
//line /snap/go/10455/src/net/nss.go:68
		_go_fuzz_dep_.CoverTab[529331]++
//line /snap/go/10455/src/net/nss.go:68
		_go_fuzz_dep_.CoverTab[7579]++
							return
//line /snap/go/10455/src/net/nss.go:69
		// _ = "end of CoverTab[7579]"
	} else {
//line /snap/go/10455/src/net/nss.go:70
		_go_fuzz_dep_.CoverTab[529332]++
//line /snap/go/10455/src/net/nss.go:70
		_go_fuzz_dep_.CoverTab[7580]++
//line /snap/go/10455/src/net/nss.go:70
		// _ = "end of CoverTab[7580]"
//line /snap/go/10455/src/net/nss.go:70
	}
//line /snap/go/10455/src/net/nss.go:70
	// _ = "end of CoverTab[7571]"
//line /snap/go/10455/src/net/nss.go:70
	_go_fuzz_dep_.CoverTab[7572]++

						nssConf := parseNSSConfFile(nssConfigPath)
						conf.mu.Lock()
						conf.nssConf = nssConf
						conf.mu.Unlock()
//line /snap/go/10455/src/net/nss.go:75
	// _ = "end of CoverTab[7572]"
}

func (conf *nsswitchConfig) acquireSema() {
//line /snap/go/10455/src/net/nss.go:78
	_go_fuzz_dep_.CoverTab[7581]++
						conf.ch <- struct{}{}
//line /snap/go/10455/src/net/nss.go:79
	// _ = "end of CoverTab[7581]"
}

func (conf *nsswitchConfig) tryAcquireSema() bool {
//line /snap/go/10455/src/net/nss.go:82
	_go_fuzz_dep_.CoverTab[7582]++
						select {
	case conf.ch <- struct{}{}:
//line /snap/go/10455/src/net/nss.go:84
		_go_fuzz_dep_.CoverTab[7583]++
							return true
//line /snap/go/10455/src/net/nss.go:85
		// _ = "end of CoverTab[7583]"
	default:
//line /snap/go/10455/src/net/nss.go:86
		_go_fuzz_dep_.CoverTab[7584]++
							return false
//line /snap/go/10455/src/net/nss.go:87
		// _ = "end of CoverTab[7584]"
	}
//line /snap/go/10455/src/net/nss.go:88
	// _ = "end of CoverTab[7582]"
}

func (conf *nsswitchConfig) releaseSema() {
//line /snap/go/10455/src/net/nss.go:91
	_go_fuzz_dep_.CoverTab[7585]++
						<-conf.ch
//line /snap/go/10455/src/net/nss.go:92
	// _ = "end of CoverTab[7585]"
}

// nssConf represents the state of the machine's /etc/nsswitch.conf file.
type nssConf struct {
	mtime	time.Time		// time of nsswitch.conf modification
	err	error			// any error encountered opening or parsing the file
	sources	map[string][]nssSource	// keyed by database (e.g. "hosts")
}

type nssSource struct {
	source		string	// e.g. "compat", "files", "mdns4_minimal"
	criteria	[]nssCriterion
}

// standardCriteria reports all specified criteria have the default
//line /snap/go/10455/src/net/nss.go:107
// status actions.
//line /snap/go/10455/src/net/nss.go:109
func (s nssSource) standardCriteria() bool {
//line /snap/go/10455/src/net/nss.go:109
	_go_fuzz_dep_.CoverTab[7586]++
//line /snap/go/10455/src/net/nss.go:109
	_go_fuzz_dep_.CoverTab[786723] = 0
						for i, crit := range s.criteria {
//line /snap/go/10455/src/net/nss.go:110
		if _go_fuzz_dep_.CoverTab[786723] == 0 {
//line /snap/go/10455/src/net/nss.go:110
			_go_fuzz_dep_.CoverTab[529370]++
//line /snap/go/10455/src/net/nss.go:110
		} else {
//line /snap/go/10455/src/net/nss.go:110
			_go_fuzz_dep_.CoverTab[529371]++
//line /snap/go/10455/src/net/nss.go:110
		}
//line /snap/go/10455/src/net/nss.go:110
		_go_fuzz_dep_.CoverTab[786723] = 1
//line /snap/go/10455/src/net/nss.go:110
		_go_fuzz_dep_.CoverTab[7588]++
							if !crit.standardStatusAction(i == len(s.criteria)-1) {
//line /snap/go/10455/src/net/nss.go:111
			_go_fuzz_dep_.CoverTab[529333]++
//line /snap/go/10455/src/net/nss.go:111
			_go_fuzz_dep_.CoverTab[7589]++
								return false
//line /snap/go/10455/src/net/nss.go:112
			// _ = "end of CoverTab[7589]"
		} else {
//line /snap/go/10455/src/net/nss.go:113
			_go_fuzz_dep_.CoverTab[529334]++
//line /snap/go/10455/src/net/nss.go:113
			_go_fuzz_dep_.CoverTab[7590]++
//line /snap/go/10455/src/net/nss.go:113
			// _ = "end of CoverTab[7590]"
//line /snap/go/10455/src/net/nss.go:113
		}
//line /snap/go/10455/src/net/nss.go:113
		// _ = "end of CoverTab[7588]"
	}
//line /snap/go/10455/src/net/nss.go:114
	if _go_fuzz_dep_.CoverTab[786723] == 0 {
//line /snap/go/10455/src/net/nss.go:114
		_go_fuzz_dep_.CoverTab[529372]++
//line /snap/go/10455/src/net/nss.go:114
	} else {
//line /snap/go/10455/src/net/nss.go:114
		_go_fuzz_dep_.CoverTab[529373]++
//line /snap/go/10455/src/net/nss.go:114
	}
//line /snap/go/10455/src/net/nss.go:114
	// _ = "end of CoverTab[7586]"
//line /snap/go/10455/src/net/nss.go:114
	_go_fuzz_dep_.CoverTab[7587]++
						return true
//line /snap/go/10455/src/net/nss.go:115
	// _ = "end of CoverTab[7587]"
}

// nssCriterion is the parsed structure of one of the criteria in brackets
//line /snap/go/10455/src/net/nss.go:118
// after an NSS source name.
//line /snap/go/10455/src/net/nss.go:120
type nssCriterion struct {
	negate	bool	// if "!" was present
	status	string	// e.g. "success", "unavail" (lowercase)
	action	string	// e.g. "return", "continue" (lowercase)
}

// standardStatusAction reports whether c is equivalent to not
//line /snap/go/10455/src/net/nss.go:126
// specifying the criterion at all. last is whether this criteria is the
//line /snap/go/10455/src/net/nss.go:126
// last in the list.
//line /snap/go/10455/src/net/nss.go:129
func (c nssCriterion) standardStatusAction(last bool) bool {
//line /snap/go/10455/src/net/nss.go:129
	_go_fuzz_dep_.CoverTab[7591]++
						if c.negate {
//line /snap/go/10455/src/net/nss.go:130
		_go_fuzz_dep_.CoverTab[529335]++
//line /snap/go/10455/src/net/nss.go:130
		_go_fuzz_dep_.CoverTab[7595]++
							return false
//line /snap/go/10455/src/net/nss.go:131
		// _ = "end of CoverTab[7595]"
	} else {
//line /snap/go/10455/src/net/nss.go:132
		_go_fuzz_dep_.CoverTab[529336]++
//line /snap/go/10455/src/net/nss.go:132
		_go_fuzz_dep_.CoverTab[7596]++
//line /snap/go/10455/src/net/nss.go:132
		// _ = "end of CoverTab[7596]"
//line /snap/go/10455/src/net/nss.go:132
	}
//line /snap/go/10455/src/net/nss.go:132
	// _ = "end of CoverTab[7591]"
//line /snap/go/10455/src/net/nss.go:132
	_go_fuzz_dep_.CoverTab[7592]++
						var def string
						switch c.status {
	case "success":
//line /snap/go/10455/src/net/nss.go:135
		_go_fuzz_dep_.CoverTab[529337]++
//line /snap/go/10455/src/net/nss.go:135
		_go_fuzz_dep_.CoverTab[7597]++
							def = "return"
//line /snap/go/10455/src/net/nss.go:136
		// _ = "end of CoverTab[7597]"
	case "notfound", "unavail", "tryagain":
//line /snap/go/10455/src/net/nss.go:137
		_go_fuzz_dep_.CoverTab[529338]++
//line /snap/go/10455/src/net/nss.go:137
		_go_fuzz_dep_.CoverTab[7598]++
							def = "continue"
//line /snap/go/10455/src/net/nss.go:138
		// _ = "end of CoverTab[7598]"
	default:
//line /snap/go/10455/src/net/nss.go:139
		_go_fuzz_dep_.CoverTab[529339]++
//line /snap/go/10455/src/net/nss.go:139
		_go_fuzz_dep_.CoverTab[7599]++

							return false
//line /snap/go/10455/src/net/nss.go:141
		// _ = "end of CoverTab[7599]"
	}
//line /snap/go/10455/src/net/nss.go:142
	// _ = "end of CoverTab[7592]"
//line /snap/go/10455/src/net/nss.go:142
	_go_fuzz_dep_.CoverTab[7593]++
						if last && func() bool {
//line /snap/go/10455/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[7600]++
//line /snap/go/10455/src/net/nss.go:143
		return c.action == "return"
//line /snap/go/10455/src/net/nss.go:143
		// _ = "end of CoverTab[7600]"
//line /snap/go/10455/src/net/nss.go:143
	}() {
//line /snap/go/10455/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[529340]++
//line /snap/go/10455/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[7601]++
							return true
//line /snap/go/10455/src/net/nss.go:144
		// _ = "end of CoverTab[7601]"
	} else {
//line /snap/go/10455/src/net/nss.go:145
		_go_fuzz_dep_.CoverTab[529341]++
//line /snap/go/10455/src/net/nss.go:145
		_go_fuzz_dep_.CoverTab[7602]++
//line /snap/go/10455/src/net/nss.go:145
		// _ = "end of CoverTab[7602]"
//line /snap/go/10455/src/net/nss.go:145
	}
//line /snap/go/10455/src/net/nss.go:145
	// _ = "end of CoverTab[7593]"
//line /snap/go/10455/src/net/nss.go:145
	_go_fuzz_dep_.CoverTab[7594]++
						return c.action == def
//line /snap/go/10455/src/net/nss.go:146
	// _ = "end of CoverTab[7594]"
}

func parseNSSConfFile(file string) *nssConf {
//line /snap/go/10455/src/net/nss.go:149
	_go_fuzz_dep_.CoverTab[7603]++
						f, err := open(file)
						if err != nil {
//line /snap/go/10455/src/net/nss.go:151
		_go_fuzz_dep_.CoverTab[529342]++
//line /snap/go/10455/src/net/nss.go:151
		_go_fuzz_dep_.CoverTab[7606]++
							return &nssConf{err: err}
//line /snap/go/10455/src/net/nss.go:152
		// _ = "end of CoverTab[7606]"
	} else {
//line /snap/go/10455/src/net/nss.go:153
		_go_fuzz_dep_.CoverTab[529343]++
//line /snap/go/10455/src/net/nss.go:153
		_go_fuzz_dep_.CoverTab[7607]++
//line /snap/go/10455/src/net/nss.go:153
		// _ = "end of CoverTab[7607]"
//line /snap/go/10455/src/net/nss.go:153
	}
//line /snap/go/10455/src/net/nss.go:153
	// _ = "end of CoverTab[7603]"
//line /snap/go/10455/src/net/nss.go:153
	_go_fuzz_dep_.CoverTab[7604]++
						defer f.close()
						mtime, _, err := f.stat()
						if err != nil {
//line /snap/go/10455/src/net/nss.go:156
		_go_fuzz_dep_.CoverTab[529344]++
//line /snap/go/10455/src/net/nss.go:156
		_go_fuzz_dep_.CoverTab[7608]++
							return &nssConf{err: err}
//line /snap/go/10455/src/net/nss.go:157
		// _ = "end of CoverTab[7608]"
	} else {
//line /snap/go/10455/src/net/nss.go:158
		_go_fuzz_dep_.CoverTab[529345]++
//line /snap/go/10455/src/net/nss.go:158
		_go_fuzz_dep_.CoverTab[7609]++
//line /snap/go/10455/src/net/nss.go:158
		// _ = "end of CoverTab[7609]"
//line /snap/go/10455/src/net/nss.go:158
	}
//line /snap/go/10455/src/net/nss.go:158
	// _ = "end of CoverTab[7604]"
//line /snap/go/10455/src/net/nss.go:158
	_go_fuzz_dep_.CoverTab[7605]++

						conf := parseNSSConf(f)
						conf.mtime = mtime
						return conf
//line /snap/go/10455/src/net/nss.go:162
	// _ = "end of CoverTab[7605]"
}

func parseNSSConf(f *file) *nssConf {
//line /snap/go/10455/src/net/nss.go:165
	_go_fuzz_dep_.CoverTab[7610]++
						conf := new(nssConf)
//line /snap/go/10455/src/net/nss.go:166
	_go_fuzz_dep_.CoverTab[786724] = 0
						for line, ok := f.readLine(); ok; line, ok = f.readLine() {
//line /snap/go/10455/src/net/nss.go:167
		if _go_fuzz_dep_.CoverTab[786724] == 0 {
//line /snap/go/10455/src/net/nss.go:167
			_go_fuzz_dep_.CoverTab[529374]++
//line /snap/go/10455/src/net/nss.go:167
		} else {
//line /snap/go/10455/src/net/nss.go:167
			_go_fuzz_dep_.CoverTab[529375]++
//line /snap/go/10455/src/net/nss.go:167
		}
//line /snap/go/10455/src/net/nss.go:167
		_go_fuzz_dep_.CoverTab[786724] = 1
//line /snap/go/10455/src/net/nss.go:167
		_go_fuzz_dep_.CoverTab[7612]++
							line = trimSpace(removeComment(line))
							if len(line) == 0 {
//line /snap/go/10455/src/net/nss.go:169
			_go_fuzz_dep_.CoverTab[529346]++
//line /snap/go/10455/src/net/nss.go:169
			_go_fuzz_dep_.CoverTab[7615]++
								continue
//line /snap/go/10455/src/net/nss.go:170
			// _ = "end of CoverTab[7615]"
		} else {
//line /snap/go/10455/src/net/nss.go:171
			_go_fuzz_dep_.CoverTab[529347]++
//line /snap/go/10455/src/net/nss.go:171
			_go_fuzz_dep_.CoverTab[7616]++
//line /snap/go/10455/src/net/nss.go:171
			// _ = "end of CoverTab[7616]"
//line /snap/go/10455/src/net/nss.go:171
		}
//line /snap/go/10455/src/net/nss.go:171
		// _ = "end of CoverTab[7612]"
//line /snap/go/10455/src/net/nss.go:171
		_go_fuzz_dep_.CoverTab[7613]++
							colon := bytealg.IndexByteString(line, ':')
							if colon == -1 {
//line /snap/go/10455/src/net/nss.go:173
			_go_fuzz_dep_.CoverTab[529348]++
//line /snap/go/10455/src/net/nss.go:173
			_go_fuzz_dep_.CoverTab[7617]++
								conf.err = errors.New("no colon on line")
								return conf
//line /snap/go/10455/src/net/nss.go:175
			// _ = "end of CoverTab[7617]"
		} else {
//line /snap/go/10455/src/net/nss.go:176
			_go_fuzz_dep_.CoverTab[529349]++
//line /snap/go/10455/src/net/nss.go:176
			_go_fuzz_dep_.CoverTab[7618]++
//line /snap/go/10455/src/net/nss.go:176
			// _ = "end of CoverTab[7618]"
//line /snap/go/10455/src/net/nss.go:176
		}
//line /snap/go/10455/src/net/nss.go:176
		// _ = "end of CoverTab[7613]"
//line /snap/go/10455/src/net/nss.go:176
		_go_fuzz_dep_.CoverTab[7614]++
							db := trimSpace(line[:colon])
							srcs := line[colon+1:]
//line /snap/go/10455/src/net/nss.go:178
		_go_fuzz_dep_.CoverTab[786725] = 0
							for {
//line /snap/go/10455/src/net/nss.go:179
			if _go_fuzz_dep_.CoverTab[786725] == 0 {
//line /snap/go/10455/src/net/nss.go:179
				_go_fuzz_dep_.CoverTab[529378]++
//line /snap/go/10455/src/net/nss.go:179
			} else {
//line /snap/go/10455/src/net/nss.go:179
				_go_fuzz_dep_.CoverTab[529379]++
//line /snap/go/10455/src/net/nss.go:179
			}
//line /snap/go/10455/src/net/nss.go:179
			_go_fuzz_dep_.CoverTab[786725] = 1
//line /snap/go/10455/src/net/nss.go:179
			_go_fuzz_dep_.CoverTab[7619]++
								srcs = trimSpace(srcs)
								if len(srcs) == 0 {
//line /snap/go/10455/src/net/nss.go:181
				_go_fuzz_dep_.CoverTab[529350]++
//line /snap/go/10455/src/net/nss.go:181
				_go_fuzz_dep_.CoverTab[7624]++
									break
//line /snap/go/10455/src/net/nss.go:182
				// _ = "end of CoverTab[7624]"
			} else {
//line /snap/go/10455/src/net/nss.go:183
				_go_fuzz_dep_.CoverTab[529351]++
//line /snap/go/10455/src/net/nss.go:183
				_go_fuzz_dep_.CoverTab[7625]++
//line /snap/go/10455/src/net/nss.go:183
				// _ = "end of CoverTab[7625]"
//line /snap/go/10455/src/net/nss.go:183
			}
//line /snap/go/10455/src/net/nss.go:183
			// _ = "end of CoverTab[7619]"
//line /snap/go/10455/src/net/nss.go:183
			_go_fuzz_dep_.CoverTab[7620]++
								sp := bytealg.IndexByteString(srcs, ' ')
								var src string
								if sp == -1 {
//line /snap/go/10455/src/net/nss.go:186
				_go_fuzz_dep_.CoverTab[529352]++
//line /snap/go/10455/src/net/nss.go:186
				_go_fuzz_dep_.CoverTab[7626]++
									src = srcs
									srcs = ""
//line /snap/go/10455/src/net/nss.go:188
				// _ = "end of CoverTab[7626]"
			} else {
//line /snap/go/10455/src/net/nss.go:189
				_go_fuzz_dep_.CoverTab[529353]++
//line /snap/go/10455/src/net/nss.go:189
				_go_fuzz_dep_.CoverTab[7627]++
									src = srcs[:sp]
									srcs = trimSpace(srcs[sp+1:])
//line /snap/go/10455/src/net/nss.go:191
				// _ = "end of CoverTab[7627]"
			}
//line /snap/go/10455/src/net/nss.go:192
			// _ = "end of CoverTab[7620]"
//line /snap/go/10455/src/net/nss.go:192
			_go_fuzz_dep_.CoverTab[7621]++
								var criteria []nssCriterion

								if len(srcs) > 0 && func() bool {
//line /snap/go/10455/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[7628]++
//line /snap/go/10455/src/net/nss.go:195
				return srcs[0] == '['
//line /snap/go/10455/src/net/nss.go:195
				// _ = "end of CoverTab[7628]"
//line /snap/go/10455/src/net/nss.go:195
			}() {
//line /snap/go/10455/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[529354]++
//line /snap/go/10455/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[7629]++
									bclose := bytealg.IndexByteString(srcs, ']')
									if bclose == -1 {
//line /snap/go/10455/src/net/nss.go:197
					_go_fuzz_dep_.CoverTab[529356]++
//line /snap/go/10455/src/net/nss.go:197
					_go_fuzz_dep_.CoverTab[7632]++
										conf.err = errors.New("unclosed criterion bracket")
										return conf
//line /snap/go/10455/src/net/nss.go:199
					// _ = "end of CoverTab[7632]"
				} else {
//line /snap/go/10455/src/net/nss.go:200
					_go_fuzz_dep_.CoverTab[529357]++
//line /snap/go/10455/src/net/nss.go:200
					_go_fuzz_dep_.CoverTab[7633]++
//line /snap/go/10455/src/net/nss.go:200
					// _ = "end of CoverTab[7633]"
//line /snap/go/10455/src/net/nss.go:200
				}
//line /snap/go/10455/src/net/nss.go:200
				// _ = "end of CoverTab[7629]"
//line /snap/go/10455/src/net/nss.go:200
				_go_fuzz_dep_.CoverTab[7630]++
									var err error
									criteria, err = parseCriteria(srcs[1:bclose])
									if err != nil {
//line /snap/go/10455/src/net/nss.go:203
					_go_fuzz_dep_.CoverTab[529358]++
//line /snap/go/10455/src/net/nss.go:203
					_go_fuzz_dep_.CoverTab[7634]++
										conf.err = errors.New("invalid criteria: " + srcs[1:bclose])
										return conf
//line /snap/go/10455/src/net/nss.go:205
					// _ = "end of CoverTab[7634]"
				} else {
//line /snap/go/10455/src/net/nss.go:206
					_go_fuzz_dep_.CoverTab[529359]++
//line /snap/go/10455/src/net/nss.go:206
					_go_fuzz_dep_.CoverTab[7635]++
//line /snap/go/10455/src/net/nss.go:206
					// _ = "end of CoverTab[7635]"
//line /snap/go/10455/src/net/nss.go:206
				}
//line /snap/go/10455/src/net/nss.go:206
				// _ = "end of CoverTab[7630]"
//line /snap/go/10455/src/net/nss.go:206
				_go_fuzz_dep_.CoverTab[7631]++
									srcs = srcs[bclose+1:]
//line /snap/go/10455/src/net/nss.go:207
				// _ = "end of CoverTab[7631]"
			} else {
//line /snap/go/10455/src/net/nss.go:208
				_go_fuzz_dep_.CoverTab[529355]++
//line /snap/go/10455/src/net/nss.go:208
				_go_fuzz_dep_.CoverTab[7636]++
//line /snap/go/10455/src/net/nss.go:208
				// _ = "end of CoverTab[7636]"
//line /snap/go/10455/src/net/nss.go:208
			}
//line /snap/go/10455/src/net/nss.go:208
			// _ = "end of CoverTab[7621]"
//line /snap/go/10455/src/net/nss.go:208
			_go_fuzz_dep_.CoverTab[7622]++
								if conf.sources == nil {
//line /snap/go/10455/src/net/nss.go:209
				_go_fuzz_dep_.CoverTab[529360]++
//line /snap/go/10455/src/net/nss.go:209
				_go_fuzz_dep_.CoverTab[7637]++
									conf.sources = make(map[string][]nssSource)
//line /snap/go/10455/src/net/nss.go:210
				// _ = "end of CoverTab[7637]"
			} else {
//line /snap/go/10455/src/net/nss.go:211
				_go_fuzz_dep_.CoverTab[529361]++
//line /snap/go/10455/src/net/nss.go:211
				_go_fuzz_dep_.CoverTab[7638]++
//line /snap/go/10455/src/net/nss.go:211
				// _ = "end of CoverTab[7638]"
//line /snap/go/10455/src/net/nss.go:211
			}
//line /snap/go/10455/src/net/nss.go:211
			// _ = "end of CoverTab[7622]"
//line /snap/go/10455/src/net/nss.go:211
			_go_fuzz_dep_.CoverTab[7623]++
								conf.sources[db] = append(conf.sources[db], nssSource{
				source:		src,
				criteria:	criteria,
			})
//line /snap/go/10455/src/net/nss.go:215
			// _ = "end of CoverTab[7623]"
		}
//line /snap/go/10455/src/net/nss.go:216
		// _ = "end of CoverTab[7614]"
	}
//line /snap/go/10455/src/net/nss.go:217
	if _go_fuzz_dep_.CoverTab[786724] == 0 {
//line /snap/go/10455/src/net/nss.go:217
		_go_fuzz_dep_.CoverTab[529376]++
//line /snap/go/10455/src/net/nss.go:217
	} else {
//line /snap/go/10455/src/net/nss.go:217
		_go_fuzz_dep_.CoverTab[529377]++
//line /snap/go/10455/src/net/nss.go:217
	}
//line /snap/go/10455/src/net/nss.go:217
	// _ = "end of CoverTab[7610]"
//line /snap/go/10455/src/net/nss.go:217
	_go_fuzz_dep_.CoverTab[7611]++
						return conf
//line /snap/go/10455/src/net/nss.go:218
	// _ = "end of CoverTab[7611]"
}

// parses "foo=bar !foo=bar"
func parseCriteria(x string) (c []nssCriterion, err error) {
//line /snap/go/10455/src/net/nss.go:222
	_go_fuzz_dep_.CoverTab[7639]++
						err = foreachField(x, func(f string) error {
//line /snap/go/10455/src/net/nss.go:223
		_go_fuzz_dep_.CoverTab[7641]++
							not := false
							if len(f) > 0 && func() bool {
//line /snap/go/10455/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[7646]++
//line /snap/go/10455/src/net/nss.go:225
			return f[0] == '!'
//line /snap/go/10455/src/net/nss.go:225
			// _ = "end of CoverTab[7646]"
//line /snap/go/10455/src/net/nss.go:225
		}() {
//line /snap/go/10455/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[529362]++
//line /snap/go/10455/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[7647]++
								not = true
								f = f[1:]
//line /snap/go/10455/src/net/nss.go:227
			// _ = "end of CoverTab[7647]"
		} else {
//line /snap/go/10455/src/net/nss.go:228
			_go_fuzz_dep_.CoverTab[529363]++
//line /snap/go/10455/src/net/nss.go:228
			_go_fuzz_dep_.CoverTab[7648]++
//line /snap/go/10455/src/net/nss.go:228
			// _ = "end of CoverTab[7648]"
//line /snap/go/10455/src/net/nss.go:228
		}
//line /snap/go/10455/src/net/nss.go:228
		// _ = "end of CoverTab[7641]"
//line /snap/go/10455/src/net/nss.go:228
		_go_fuzz_dep_.CoverTab[7642]++
							if len(f) < 3 {
//line /snap/go/10455/src/net/nss.go:229
			_go_fuzz_dep_.CoverTab[529364]++
//line /snap/go/10455/src/net/nss.go:229
			_go_fuzz_dep_.CoverTab[7649]++
								return errors.New("criterion too short")
//line /snap/go/10455/src/net/nss.go:230
			// _ = "end of CoverTab[7649]"
		} else {
//line /snap/go/10455/src/net/nss.go:231
			_go_fuzz_dep_.CoverTab[529365]++
//line /snap/go/10455/src/net/nss.go:231
			_go_fuzz_dep_.CoverTab[7650]++
//line /snap/go/10455/src/net/nss.go:231
			// _ = "end of CoverTab[7650]"
//line /snap/go/10455/src/net/nss.go:231
		}
//line /snap/go/10455/src/net/nss.go:231
		// _ = "end of CoverTab[7642]"
//line /snap/go/10455/src/net/nss.go:231
		_go_fuzz_dep_.CoverTab[7643]++
							eq := bytealg.IndexByteString(f, '=')
							if eq == -1 {
//line /snap/go/10455/src/net/nss.go:233
			_go_fuzz_dep_.CoverTab[529366]++
//line /snap/go/10455/src/net/nss.go:233
			_go_fuzz_dep_.CoverTab[7651]++
								return errors.New("criterion lacks equal sign")
//line /snap/go/10455/src/net/nss.go:234
			// _ = "end of CoverTab[7651]"
		} else {
//line /snap/go/10455/src/net/nss.go:235
			_go_fuzz_dep_.CoverTab[529367]++
//line /snap/go/10455/src/net/nss.go:235
			_go_fuzz_dep_.CoverTab[7652]++
//line /snap/go/10455/src/net/nss.go:235
			// _ = "end of CoverTab[7652]"
//line /snap/go/10455/src/net/nss.go:235
		}
//line /snap/go/10455/src/net/nss.go:235
		// _ = "end of CoverTab[7643]"
//line /snap/go/10455/src/net/nss.go:235
		_go_fuzz_dep_.CoverTab[7644]++
							if hasUpperCase(f) {
//line /snap/go/10455/src/net/nss.go:236
			_go_fuzz_dep_.CoverTab[529368]++
//line /snap/go/10455/src/net/nss.go:236
			_go_fuzz_dep_.CoverTab[7653]++
								lower := []byte(f)
								lowerASCIIBytes(lower)
								f = string(lower)
//line /snap/go/10455/src/net/nss.go:239
			// _ = "end of CoverTab[7653]"
		} else {
//line /snap/go/10455/src/net/nss.go:240
			_go_fuzz_dep_.CoverTab[529369]++
//line /snap/go/10455/src/net/nss.go:240
			_go_fuzz_dep_.CoverTab[7654]++
//line /snap/go/10455/src/net/nss.go:240
			// _ = "end of CoverTab[7654]"
//line /snap/go/10455/src/net/nss.go:240
		}
//line /snap/go/10455/src/net/nss.go:240
		// _ = "end of CoverTab[7644]"
//line /snap/go/10455/src/net/nss.go:240
		_go_fuzz_dep_.CoverTab[7645]++
							c = append(c, nssCriterion{
			negate:	not,
			status:	f[:eq],
			action:	f[eq+1:],
		})
							return nil
//line /snap/go/10455/src/net/nss.go:246
		// _ = "end of CoverTab[7645]"
	})
//line /snap/go/10455/src/net/nss.go:247
	// _ = "end of CoverTab[7639]"
//line /snap/go/10455/src/net/nss.go:247
	_go_fuzz_dep_.CoverTab[7640]++
						return
//line /snap/go/10455/src/net/nss.go:248
	// _ = "end of CoverTab[7640]"
}

//line /snap/go/10455/src/net/nss.go:249
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/nss.go:249
var _ = _go_fuzz_dep_.CoverTab
