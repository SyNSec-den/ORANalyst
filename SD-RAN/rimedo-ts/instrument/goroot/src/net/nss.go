// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/nss.go:5
package net

//line /usr/local/go/src/net/nss.go:5
import (
//line /usr/local/go/src/net/nss.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/nss.go:5
)
//line /usr/local/go/src/net/nss.go:5
import (
//line /usr/local/go/src/net/nss.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/nss.go:5
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
//line /usr/local/go/src/net/nss.go:33
	_go_fuzz_dep_.CoverTab[15677]++
					nssConfig.tryUpdate()
					nssConfig.mu.Lock()
					conf := nssConfig.nssConf
					nssConfig.mu.Unlock()
					return conf
//line /usr/local/go/src/net/nss.go:38
	// _ = "end of CoverTab[15677]"
}

// init initializes conf and is only called via conf.initOnce.
func (conf *nsswitchConfig) init() {
	conf.nssConf = parseNSSConfFile("/etc/nsswitch.conf")
	conf.lastChecked = time.Now()
	conf.ch = make(chan struct{}, 1)
}

// tryUpdate tries to update conf.
func (conf *nsswitchConfig) tryUpdate() {
//line /usr/local/go/src/net/nss.go:49
	_go_fuzz_dep_.CoverTab[15678]++
					conf.initOnce.Do(conf.init)

//line /usr/local/go/src/net/nss.go:53
	if !conf.tryAcquireSema() {
//line /usr/local/go/src/net/nss.go:53
		_go_fuzz_dep_.CoverTab[15683]++
						return
//line /usr/local/go/src/net/nss.go:54
		// _ = "end of CoverTab[15683]"
	} else {
//line /usr/local/go/src/net/nss.go:55
		_go_fuzz_dep_.CoverTab[15684]++
//line /usr/local/go/src/net/nss.go:55
		// _ = "end of CoverTab[15684]"
//line /usr/local/go/src/net/nss.go:55
	}
//line /usr/local/go/src/net/nss.go:55
	// _ = "end of CoverTab[15678]"
//line /usr/local/go/src/net/nss.go:55
	_go_fuzz_dep_.CoverTab[15679]++
					defer conf.releaseSema()

					now := time.Now()
					if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /usr/local/go/src/net/nss.go:59
		_go_fuzz_dep_.CoverTab[15685]++
						return
//line /usr/local/go/src/net/nss.go:60
		// _ = "end of CoverTab[15685]"
	} else {
//line /usr/local/go/src/net/nss.go:61
		_go_fuzz_dep_.CoverTab[15686]++
//line /usr/local/go/src/net/nss.go:61
		// _ = "end of CoverTab[15686]"
//line /usr/local/go/src/net/nss.go:61
	}
//line /usr/local/go/src/net/nss.go:61
	// _ = "end of CoverTab[15679]"
//line /usr/local/go/src/net/nss.go:61
	_go_fuzz_dep_.CoverTab[15680]++
					conf.lastChecked = now

					var mtime time.Time
					if fi, err := os.Stat(nssConfigPath); err == nil {
//line /usr/local/go/src/net/nss.go:65
		_go_fuzz_dep_.CoverTab[15687]++
						mtime = fi.ModTime()
//line /usr/local/go/src/net/nss.go:66
		// _ = "end of CoverTab[15687]"
	} else {
//line /usr/local/go/src/net/nss.go:67
		_go_fuzz_dep_.CoverTab[15688]++
//line /usr/local/go/src/net/nss.go:67
		// _ = "end of CoverTab[15688]"
//line /usr/local/go/src/net/nss.go:67
	}
//line /usr/local/go/src/net/nss.go:67
	// _ = "end of CoverTab[15680]"
//line /usr/local/go/src/net/nss.go:67
	_go_fuzz_dep_.CoverTab[15681]++
					if mtime.Equal(conf.nssConf.mtime) {
//line /usr/local/go/src/net/nss.go:68
		_go_fuzz_dep_.CoverTab[15689]++
						return
//line /usr/local/go/src/net/nss.go:69
		// _ = "end of CoverTab[15689]"
	} else {
//line /usr/local/go/src/net/nss.go:70
		_go_fuzz_dep_.CoverTab[15690]++
//line /usr/local/go/src/net/nss.go:70
		// _ = "end of CoverTab[15690]"
//line /usr/local/go/src/net/nss.go:70
	}
//line /usr/local/go/src/net/nss.go:70
	// _ = "end of CoverTab[15681]"
//line /usr/local/go/src/net/nss.go:70
	_go_fuzz_dep_.CoverTab[15682]++

					nssConf := parseNSSConfFile(nssConfigPath)
					conf.mu.Lock()
					conf.nssConf = nssConf
					conf.mu.Unlock()
//line /usr/local/go/src/net/nss.go:75
	// _ = "end of CoverTab[15682]"
}

func (conf *nsswitchConfig) acquireSema() {
//line /usr/local/go/src/net/nss.go:78
	_go_fuzz_dep_.CoverTab[15691]++
					conf.ch <- struct{}{}
//line /usr/local/go/src/net/nss.go:79
	// _ = "end of CoverTab[15691]"
}

func (conf *nsswitchConfig) tryAcquireSema() bool {
//line /usr/local/go/src/net/nss.go:82
	_go_fuzz_dep_.CoverTab[15692]++
					select {
	case conf.ch <- struct{}{}:
//line /usr/local/go/src/net/nss.go:84
		_go_fuzz_dep_.CoverTab[15693]++
						return true
//line /usr/local/go/src/net/nss.go:85
		// _ = "end of CoverTab[15693]"
	default:
//line /usr/local/go/src/net/nss.go:86
		_go_fuzz_dep_.CoverTab[15694]++
						return false
//line /usr/local/go/src/net/nss.go:87
		// _ = "end of CoverTab[15694]"
	}
//line /usr/local/go/src/net/nss.go:88
	// _ = "end of CoverTab[15692]"
}

func (conf *nsswitchConfig) releaseSema() {
//line /usr/local/go/src/net/nss.go:91
	_go_fuzz_dep_.CoverTab[15695]++
					<-conf.ch
//line /usr/local/go/src/net/nss.go:92
	// _ = "end of CoverTab[15695]"
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
//line /usr/local/go/src/net/nss.go:107
// status actions.
//line /usr/local/go/src/net/nss.go:109
func (s nssSource) standardCriteria() bool {
//line /usr/local/go/src/net/nss.go:109
	_go_fuzz_dep_.CoverTab[15696]++
						for i, crit := range s.criteria {
//line /usr/local/go/src/net/nss.go:110
		_go_fuzz_dep_.CoverTab[15698]++
							if !crit.standardStatusAction(i == len(s.criteria)-1) {
//line /usr/local/go/src/net/nss.go:111
			_go_fuzz_dep_.CoverTab[15699]++
								return false
//line /usr/local/go/src/net/nss.go:112
			// _ = "end of CoverTab[15699]"
		} else {
//line /usr/local/go/src/net/nss.go:113
			_go_fuzz_dep_.CoverTab[15700]++
//line /usr/local/go/src/net/nss.go:113
			// _ = "end of CoverTab[15700]"
//line /usr/local/go/src/net/nss.go:113
		}
//line /usr/local/go/src/net/nss.go:113
		// _ = "end of CoverTab[15698]"
	}
//line /usr/local/go/src/net/nss.go:114
	// _ = "end of CoverTab[15696]"
//line /usr/local/go/src/net/nss.go:114
	_go_fuzz_dep_.CoverTab[15697]++
						return true
//line /usr/local/go/src/net/nss.go:115
	// _ = "end of CoverTab[15697]"
}

// nssCriterion is the parsed structure of one of the criteria in brackets
//line /usr/local/go/src/net/nss.go:118
// after an NSS source name.
//line /usr/local/go/src/net/nss.go:120
type nssCriterion struct {
	negate	bool	// if "!" was present
	status	string	// e.g. "success", "unavail" (lowercase)
	action	string	// e.g. "return", "continue" (lowercase)
}

// standardStatusAction reports whether c is equivalent to not
//line /usr/local/go/src/net/nss.go:126
// specifying the criterion at all. last is whether this criteria is the
//line /usr/local/go/src/net/nss.go:126
// last in the list.
//line /usr/local/go/src/net/nss.go:129
func (c nssCriterion) standardStatusAction(last bool) bool {
//line /usr/local/go/src/net/nss.go:129
	_go_fuzz_dep_.CoverTab[15701]++
						if c.negate {
//line /usr/local/go/src/net/nss.go:130
		_go_fuzz_dep_.CoverTab[15705]++
							return false
//line /usr/local/go/src/net/nss.go:131
		// _ = "end of CoverTab[15705]"
	} else {
//line /usr/local/go/src/net/nss.go:132
		_go_fuzz_dep_.CoverTab[15706]++
//line /usr/local/go/src/net/nss.go:132
		// _ = "end of CoverTab[15706]"
//line /usr/local/go/src/net/nss.go:132
	}
//line /usr/local/go/src/net/nss.go:132
	// _ = "end of CoverTab[15701]"
//line /usr/local/go/src/net/nss.go:132
	_go_fuzz_dep_.CoverTab[15702]++
						var def string
						switch c.status {
	case "success":
//line /usr/local/go/src/net/nss.go:135
		_go_fuzz_dep_.CoverTab[15707]++
							def = "return"
//line /usr/local/go/src/net/nss.go:136
		// _ = "end of CoverTab[15707]"
	case "notfound", "unavail", "tryagain":
//line /usr/local/go/src/net/nss.go:137
		_go_fuzz_dep_.CoverTab[15708]++
							def = "continue"
//line /usr/local/go/src/net/nss.go:138
		// _ = "end of CoverTab[15708]"
	default:
//line /usr/local/go/src/net/nss.go:139
		_go_fuzz_dep_.CoverTab[15709]++

							return false
//line /usr/local/go/src/net/nss.go:141
		// _ = "end of CoverTab[15709]"
	}
//line /usr/local/go/src/net/nss.go:142
	// _ = "end of CoverTab[15702]"
//line /usr/local/go/src/net/nss.go:142
	_go_fuzz_dep_.CoverTab[15703]++
						if last && func() bool {
//line /usr/local/go/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[15710]++
//line /usr/local/go/src/net/nss.go:143
		return c.action == "return"
//line /usr/local/go/src/net/nss.go:143
		// _ = "end of CoverTab[15710]"
//line /usr/local/go/src/net/nss.go:143
	}() {
//line /usr/local/go/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[15711]++
							return true
//line /usr/local/go/src/net/nss.go:144
		// _ = "end of CoverTab[15711]"
	} else {
//line /usr/local/go/src/net/nss.go:145
		_go_fuzz_dep_.CoverTab[15712]++
//line /usr/local/go/src/net/nss.go:145
		// _ = "end of CoverTab[15712]"
//line /usr/local/go/src/net/nss.go:145
	}
//line /usr/local/go/src/net/nss.go:145
	// _ = "end of CoverTab[15703]"
//line /usr/local/go/src/net/nss.go:145
	_go_fuzz_dep_.CoverTab[15704]++
						return c.action == def
//line /usr/local/go/src/net/nss.go:146
	// _ = "end of CoverTab[15704]"
}

func parseNSSConfFile(file string) *nssConf {
//line /usr/local/go/src/net/nss.go:149
	_go_fuzz_dep_.CoverTab[15713]++
						f, err := open(file)
						if err != nil {
//line /usr/local/go/src/net/nss.go:151
		_go_fuzz_dep_.CoverTab[15716]++
							return &nssConf{err: err}
//line /usr/local/go/src/net/nss.go:152
		// _ = "end of CoverTab[15716]"
	} else {
//line /usr/local/go/src/net/nss.go:153
		_go_fuzz_dep_.CoverTab[15717]++
//line /usr/local/go/src/net/nss.go:153
		// _ = "end of CoverTab[15717]"
//line /usr/local/go/src/net/nss.go:153
	}
//line /usr/local/go/src/net/nss.go:153
	// _ = "end of CoverTab[15713]"
//line /usr/local/go/src/net/nss.go:153
	_go_fuzz_dep_.CoverTab[15714]++
						defer f.close()
						mtime, _, err := f.stat()
						if err != nil {
//line /usr/local/go/src/net/nss.go:156
		_go_fuzz_dep_.CoverTab[15718]++
							return &nssConf{err: err}
//line /usr/local/go/src/net/nss.go:157
		// _ = "end of CoverTab[15718]"
	} else {
//line /usr/local/go/src/net/nss.go:158
		_go_fuzz_dep_.CoverTab[15719]++
//line /usr/local/go/src/net/nss.go:158
		// _ = "end of CoverTab[15719]"
//line /usr/local/go/src/net/nss.go:158
	}
//line /usr/local/go/src/net/nss.go:158
	// _ = "end of CoverTab[15714]"
//line /usr/local/go/src/net/nss.go:158
	_go_fuzz_dep_.CoverTab[15715]++

						conf := parseNSSConf(f)
						conf.mtime = mtime
						return conf
//line /usr/local/go/src/net/nss.go:162
	// _ = "end of CoverTab[15715]"
}

func parseNSSConf(f *file) *nssConf {
//line /usr/local/go/src/net/nss.go:165
	_go_fuzz_dep_.CoverTab[15720]++
						conf := new(nssConf)
						for line, ok := f.readLine(); ok; line, ok = f.readLine() {
//line /usr/local/go/src/net/nss.go:167
		_go_fuzz_dep_.CoverTab[15722]++
							line = trimSpace(removeComment(line))
							if len(line) == 0 {
//line /usr/local/go/src/net/nss.go:169
			_go_fuzz_dep_.CoverTab[15725]++
								continue
//line /usr/local/go/src/net/nss.go:170
			// _ = "end of CoverTab[15725]"
		} else {
//line /usr/local/go/src/net/nss.go:171
			_go_fuzz_dep_.CoverTab[15726]++
//line /usr/local/go/src/net/nss.go:171
			// _ = "end of CoverTab[15726]"
//line /usr/local/go/src/net/nss.go:171
		}
//line /usr/local/go/src/net/nss.go:171
		// _ = "end of CoverTab[15722]"
//line /usr/local/go/src/net/nss.go:171
		_go_fuzz_dep_.CoverTab[15723]++
							colon := bytealg.IndexByteString(line, ':')
							if colon == -1 {
//line /usr/local/go/src/net/nss.go:173
			_go_fuzz_dep_.CoverTab[15727]++
								conf.err = errors.New("no colon on line")
								return conf
//line /usr/local/go/src/net/nss.go:175
			// _ = "end of CoverTab[15727]"
		} else {
//line /usr/local/go/src/net/nss.go:176
			_go_fuzz_dep_.CoverTab[15728]++
//line /usr/local/go/src/net/nss.go:176
			// _ = "end of CoverTab[15728]"
//line /usr/local/go/src/net/nss.go:176
		}
//line /usr/local/go/src/net/nss.go:176
		// _ = "end of CoverTab[15723]"
//line /usr/local/go/src/net/nss.go:176
		_go_fuzz_dep_.CoverTab[15724]++
							db := trimSpace(line[:colon])
							srcs := line[colon+1:]
							for {
//line /usr/local/go/src/net/nss.go:179
			_go_fuzz_dep_.CoverTab[15729]++
								srcs = trimSpace(srcs)
								if len(srcs) == 0 {
//line /usr/local/go/src/net/nss.go:181
				_go_fuzz_dep_.CoverTab[15734]++
									break
//line /usr/local/go/src/net/nss.go:182
				// _ = "end of CoverTab[15734]"
			} else {
//line /usr/local/go/src/net/nss.go:183
				_go_fuzz_dep_.CoverTab[15735]++
//line /usr/local/go/src/net/nss.go:183
				// _ = "end of CoverTab[15735]"
//line /usr/local/go/src/net/nss.go:183
			}
//line /usr/local/go/src/net/nss.go:183
			// _ = "end of CoverTab[15729]"
//line /usr/local/go/src/net/nss.go:183
			_go_fuzz_dep_.CoverTab[15730]++
								sp := bytealg.IndexByteString(srcs, ' ')
								var src string
								if sp == -1 {
//line /usr/local/go/src/net/nss.go:186
				_go_fuzz_dep_.CoverTab[15736]++
									src = srcs
									srcs = ""
//line /usr/local/go/src/net/nss.go:188
				// _ = "end of CoverTab[15736]"
			} else {
//line /usr/local/go/src/net/nss.go:189
				_go_fuzz_dep_.CoverTab[15737]++
									src = srcs[:sp]
									srcs = trimSpace(srcs[sp+1:])
//line /usr/local/go/src/net/nss.go:191
				// _ = "end of CoverTab[15737]"
			}
//line /usr/local/go/src/net/nss.go:192
			// _ = "end of CoverTab[15730]"
//line /usr/local/go/src/net/nss.go:192
			_go_fuzz_dep_.CoverTab[15731]++
								var criteria []nssCriterion

								if len(srcs) > 0 && func() bool {
//line /usr/local/go/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[15738]++
//line /usr/local/go/src/net/nss.go:195
				return srcs[0] == '['
//line /usr/local/go/src/net/nss.go:195
				// _ = "end of CoverTab[15738]"
//line /usr/local/go/src/net/nss.go:195
			}() {
//line /usr/local/go/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[15739]++
									bclose := bytealg.IndexByteString(srcs, ']')
									if bclose == -1 {
//line /usr/local/go/src/net/nss.go:197
					_go_fuzz_dep_.CoverTab[15742]++
										conf.err = errors.New("unclosed criterion bracket")
										return conf
//line /usr/local/go/src/net/nss.go:199
					// _ = "end of CoverTab[15742]"
				} else {
//line /usr/local/go/src/net/nss.go:200
					_go_fuzz_dep_.CoverTab[15743]++
//line /usr/local/go/src/net/nss.go:200
					// _ = "end of CoverTab[15743]"
//line /usr/local/go/src/net/nss.go:200
				}
//line /usr/local/go/src/net/nss.go:200
				// _ = "end of CoverTab[15739]"
//line /usr/local/go/src/net/nss.go:200
				_go_fuzz_dep_.CoverTab[15740]++
									var err error
									criteria, err = parseCriteria(srcs[1:bclose])
									if err != nil {
//line /usr/local/go/src/net/nss.go:203
					_go_fuzz_dep_.CoverTab[15744]++
										conf.err = errors.New("invalid criteria: " + srcs[1:bclose])
										return conf
//line /usr/local/go/src/net/nss.go:205
					// _ = "end of CoverTab[15744]"
				} else {
//line /usr/local/go/src/net/nss.go:206
					_go_fuzz_dep_.CoverTab[15745]++
//line /usr/local/go/src/net/nss.go:206
					// _ = "end of CoverTab[15745]"
//line /usr/local/go/src/net/nss.go:206
				}
//line /usr/local/go/src/net/nss.go:206
				// _ = "end of CoverTab[15740]"
//line /usr/local/go/src/net/nss.go:206
				_go_fuzz_dep_.CoverTab[15741]++
									srcs = srcs[bclose+1:]
//line /usr/local/go/src/net/nss.go:207
				// _ = "end of CoverTab[15741]"
			} else {
//line /usr/local/go/src/net/nss.go:208
				_go_fuzz_dep_.CoverTab[15746]++
//line /usr/local/go/src/net/nss.go:208
				// _ = "end of CoverTab[15746]"
//line /usr/local/go/src/net/nss.go:208
			}
//line /usr/local/go/src/net/nss.go:208
			// _ = "end of CoverTab[15731]"
//line /usr/local/go/src/net/nss.go:208
			_go_fuzz_dep_.CoverTab[15732]++
								if conf.sources == nil {
//line /usr/local/go/src/net/nss.go:209
				_go_fuzz_dep_.CoverTab[15747]++
									conf.sources = make(map[string][]nssSource)
//line /usr/local/go/src/net/nss.go:210
				// _ = "end of CoverTab[15747]"
			} else {
//line /usr/local/go/src/net/nss.go:211
				_go_fuzz_dep_.CoverTab[15748]++
//line /usr/local/go/src/net/nss.go:211
				// _ = "end of CoverTab[15748]"
//line /usr/local/go/src/net/nss.go:211
			}
//line /usr/local/go/src/net/nss.go:211
			// _ = "end of CoverTab[15732]"
//line /usr/local/go/src/net/nss.go:211
			_go_fuzz_dep_.CoverTab[15733]++
								conf.sources[db] = append(conf.sources[db], nssSource{
				source:		src,
				criteria:	criteria,
			})
//line /usr/local/go/src/net/nss.go:215
			// _ = "end of CoverTab[15733]"
		}
//line /usr/local/go/src/net/nss.go:216
		// _ = "end of CoverTab[15724]"
	}
//line /usr/local/go/src/net/nss.go:217
	// _ = "end of CoverTab[15720]"
//line /usr/local/go/src/net/nss.go:217
	_go_fuzz_dep_.CoverTab[15721]++
						return conf
//line /usr/local/go/src/net/nss.go:218
	// _ = "end of CoverTab[15721]"
}

// parses "foo=bar !foo=bar"
func parseCriteria(x string) (c []nssCriterion, err error) {
//line /usr/local/go/src/net/nss.go:222
	_go_fuzz_dep_.CoverTab[15749]++
						err = foreachField(x, func(f string) error {
//line /usr/local/go/src/net/nss.go:223
		_go_fuzz_dep_.CoverTab[15751]++
							not := false
							if len(f) > 0 && func() bool {
//line /usr/local/go/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[15756]++
//line /usr/local/go/src/net/nss.go:225
			return f[0] == '!'
//line /usr/local/go/src/net/nss.go:225
			// _ = "end of CoverTab[15756]"
//line /usr/local/go/src/net/nss.go:225
		}() {
//line /usr/local/go/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[15757]++
								not = true
								f = f[1:]
//line /usr/local/go/src/net/nss.go:227
			// _ = "end of CoverTab[15757]"
		} else {
//line /usr/local/go/src/net/nss.go:228
			_go_fuzz_dep_.CoverTab[15758]++
//line /usr/local/go/src/net/nss.go:228
			// _ = "end of CoverTab[15758]"
//line /usr/local/go/src/net/nss.go:228
		}
//line /usr/local/go/src/net/nss.go:228
		// _ = "end of CoverTab[15751]"
//line /usr/local/go/src/net/nss.go:228
		_go_fuzz_dep_.CoverTab[15752]++
							if len(f) < 3 {
//line /usr/local/go/src/net/nss.go:229
			_go_fuzz_dep_.CoverTab[15759]++
								return errors.New("criterion too short")
//line /usr/local/go/src/net/nss.go:230
			// _ = "end of CoverTab[15759]"
		} else {
//line /usr/local/go/src/net/nss.go:231
			_go_fuzz_dep_.CoverTab[15760]++
//line /usr/local/go/src/net/nss.go:231
			// _ = "end of CoverTab[15760]"
//line /usr/local/go/src/net/nss.go:231
		}
//line /usr/local/go/src/net/nss.go:231
		// _ = "end of CoverTab[15752]"
//line /usr/local/go/src/net/nss.go:231
		_go_fuzz_dep_.CoverTab[15753]++
							eq := bytealg.IndexByteString(f, '=')
							if eq == -1 {
//line /usr/local/go/src/net/nss.go:233
			_go_fuzz_dep_.CoverTab[15761]++
								return errors.New("criterion lacks equal sign")
//line /usr/local/go/src/net/nss.go:234
			// _ = "end of CoverTab[15761]"
		} else {
//line /usr/local/go/src/net/nss.go:235
			_go_fuzz_dep_.CoverTab[15762]++
//line /usr/local/go/src/net/nss.go:235
			// _ = "end of CoverTab[15762]"
//line /usr/local/go/src/net/nss.go:235
		}
//line /usr/local/go/src/net/nss.go:235
		// _ = "end of CoverTab[15753]"
//line /usr/local/go/src/net/nss.go:235
		_go_fuzz_dep_.CoverTab[15754]++
							if hasUpperCase(f) {
//line /usr/local/go/src/net/nss.go:236
			_go_fuzz_dep_.CoverTab[15763]++
								lower := []byte(f)
								lowerASCIIBytes(lower)
								f = string(lower)
//line /usr/local/go/src/net/nss.go:239
			// _ = "end of CoverTab[15763]"
		} else {
//line /usr/local/go/src/net/nss.go:240
			_go_fuzz_dep_.CoverTab[15764]++
//line /usr/local/go/src/net/nss.go:240
			// _ = "end of CoverTab[15764]"
//line /usr/local/go/src/net/nss.go:240
		}
//line /usr/local/go/src/net/nss.go:240
		// _ = "end of CoverTab[15754]"
//line /usr/local/go/src/net/nss.go:240
		_go_fuzz_dep_.CoverTab[15755]++
							c = append(c, nssCriterion{
			negate:	not,
			status:	f[:eq],
			action:	f[eq+1:],
		})
							return nil
//line /usr/local/go/src/net/nss.go:246
		// _ = "end of CoverTab[15755]"
	})
//line /usr/local/go/src/net/nss.go:247
	// _ = "end of CoverTab[15749]"
//line /usr/local/go/src/net/nss.go:247
	_go_fuzz_dep_.CoverTab[15750]++
						return
//line /usr/local/go/src/net/nss.go:248
	// _ = "end of CoverTab[15750]"
}

//line /usr/local/go/src/net/nss.go:249
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/nss.go:249
var _ = _go_fuzz_dep_.CoverTab
