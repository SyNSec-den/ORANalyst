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
	_go_fuzz_dep_.CoverTab[7287]++
					nssConfig.tryUpdate()
					nssConfig.mu.Lock()
					conf := nssConfig.nssConf
					nssConfig.mu.Unlock()
					return conf
//line /usr/local/go/src/net/nss.go:38
	// _ = "end of CoverTab[7287]"
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
	_go_fuzz_dep_.CoverTab[7288]++
					conf.initOnce.Do(conf.init)

//line /usr/local/go/src/net/nss.go:53
	if !conf.tryAcquireSema() {
//line /usr/local/go/src/net/nss.go:53
		_go_fuzz_dep_.CoverTab[7293]++
						return
//line /usr/local/go/src/net/nss.go:54
		// _ = "end of CoverTab[7293]"
	} else {
//line /usr/local/go/src/net/nss.go:55
		_go_fuzz_dep_.CoverTab[7294]++
//line /usr/local/go/src/net/nss.go:55
		// _ = "end of CoverTab[7294]"
//line /usr/local/go/src/net/nss.go:55
	}
//line /usr/local/go/src/net/nss.go:55
	// _ = "end of CoverTab[7288]"
//line /usr/local/go/src/net/nss.go:55
	_go_fuzz_dep_.CoverTab[7289]++
					defer conf.releaseSema()

					now := time.Now()
					if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /usr/local/go/src/net/nss.go:59
		_go_fuzz_dep_.CoverTab[7295]++
						return
//line /usr/local/go/src/net/nss.go:60
		// _ = "end of CoverTab[7295]"
	} else {
//line /usr/local/go/src/net/nss.go:61
		_go_fuzz_dep_.CoverTab[7296]++
//line /usr/local/go/src/net/nss.go:61
		// _ = "end of CoverTab[7296]"
//line /usr/local/go/src/net/nss.go:61
	}
//line /usr/local/go/src/net/nss.go:61
	// _ = "end of CoverTab[7289]"
//line /usr/local/go/src/net/nss.go:61
	_go_fuzz_dep_.CoverTab[7290]++
					conf.lastChecked = now

					var mtime time.Time
					if fi, err := os.Stat(nssConfigPath); err == nil {
//line /usr/local/go/src/net/nss.go:65
		_go_fuzz_dep_.CoverTab[7297]++
						mtime = fi.ModTime()
//line /usr/local/go/src/net/nss.go:66
		// _ = "end of CoverTab[7297]"
	} else {
//line /usr/local/go/src/net/nss.go:67
		_go_fuzz_dep_.CoverTab[7298]++
//line /usr/local/go/src/net/nss.go:67
		// _ = "end of CoverTab[7298]"
//line /usr/local/go/src/net/nss.go:67
	}
//line /usr/local/go/src/net/nss.go:67
	// _ = "end of CoverTab[7290]"
//line /usr/local/go/src/net/nss.go:67
	_go_fuzz_dep_.CoverTab[7291]++
					if mtime.Equal(conf.nssConf.mtime) {
//line /usr/local/go/src/net/nss.go:68
		_go_fuzz_dep_.CoverTab[7299]++
						return
//line /usr/local/go/src/net/nss.go:69
		// _ = "end of CoverTab[7299]"
	} else {
//line /usr/local/go/src/net/nss.go:70
		_go_fuzz_dep_.CoverTab[7300]++
//line /usr/local/go/src/net/nss.go:70
		// _ = "end of CoverTab[7300]"
//line /usr/local/go/src/net/nss.go:70
	}
//line /usr/local/go/src/net/nss.go:70
	// _ = "end of CoverTab[7291]"
//line /usr/local/go/src/net/nss.go:70
	_go_fuzz_dep_.CoverTab[7292]++

					nssConf := parseNSSConfFile(nssConfigPath)
					conf.mu.Lock()
					conf.nssConf = nssConf
					conf.mu.Unlock()
//line /usr/local/go/src/net/nss.go:75
	// _ = "end of CoverTab[7292]"
}

func (conf *nsswitchConfig) acquireSema() {
//line /usr/local/go/src/net/nss.go:78
	_go_fuzz_dep_.CoverTab[7301]++
					conf.ch <- struct{}{}
//line /usr/local/go/src/net/nss.go:79
	// _ = "end of CoverTab[7301]"
}

func (conf *nsswitchConfig) tryAcquireSema() bool {
//line /usr/local/go/src/net/nss.go:82
	_go_fuzz_dep_.CoverTab[7302]++
					select {
	case conf.ch <- struct{}{}:
//line /usr/local/go/src/net/nss.go:84
		_go_fuzz_dep_.CoverTab[7303]++
						return true
//line /usr/local/go/src/net/nss.go:85
		// _ = "end of CoverTab[7303]"
	default:
//line /usr/local/go/src/net/nss.go:86
		_go_fuzz_dep_.CoverTab[7304]++
						return false
//line /usr/local/go/src/net/nss.go:87
		// _ = "end of CoverTab[7304]"
	}
//line /usr/local/go/src/net/nss.go:88
	// _ = "end of CoverTab[7302]"
}

func (conf *nsswitchConfig) releaseSema() {
//line /usr/local/go/src/net/nss.go:91
	_go_fuzz_dep_.CoverTab[7305]++
					<-conf.ch
//line /usr/local/go/src/net/nss.go:92
	// _ = "end of CoverTab[7305]"
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
	_go_fuzz_dep_.CoverTab[7306]++
						for i, crit := range s.criteria {
//line /usr/local/go/src/net/nss.go:110
		_go_fuzz_dep_.CoverTab[7308]++
							if !crit.standardStatusAction(i == len(s.criteria)-1) {
//line /usr/local/go/src/net/nss.go:111
			_go_fuzz_dep_.CoverTab[7309]++
								return false
//line /usr/local/go/src/net/nss.go:112
			// _ = "end of CoverTab[7309]"
		} else {
//line /usr/local/go/src/net/nss.go:113
			_go_fuzz_dep_.CoverTab[7310]++
//line /usr/local/go/src/net/nss.go:113
			// _ = "end of CoverTab[7310]"
//line /usr/local/go/src/net/nss.go:113
		}
//line /usr/local/go/src/net/nss.go:113
		// _ = "end of CoverTab[7308]"
	}
//line /usr/local/go/src/net/nss.go:114
	// _ = "end of CoverTab[7306]"
//line /usr/local/go/src/net/nss.go:114
	_go_fuzz_dep_.CoverTab[7307]++
						return true
//line /usr/local/go/src/net/nss.go:115
	// _ = "end of CoverTab[7307]"
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
	_go_fuzz_dep_.CoverTab[7311]++
						if c.negate {
//line /usr/local/go/src/net/nss.go:130
		_go_fuzz_dep_.CoverTab[7315]++
							return false
//line /usr/local/go/src/net/nss.go:131
		// _ = "end of CoverTab[7315]"
	} else {
//line /usr/local/go/src/net/nss.go:132
		_go_fuzz_dep_.CoverTab[7316]++
//line /usr/local/go/src/net/nss.go:132
		// _ = "end of CoverTab[7316]"
//line /usr/local/go/src/net/nss.go:132
	}
//line /usr/local/go/src/net/nss.go:132
	// _ = "end of CoverTab[7311]"
//line /usr/local/go/src/net/nss.go:132
	_go_fuzz_dep_.CoverTab[7312]++
						var def string
						switch c.status {
	case "success":
//line /usr/local/go/src/net/nss.go:135
		_go_fuzz_dep_.CoverTab[7317]++
							def = "return"
//line /usr/local/go/src/net/nss.go:136
		// _ = "end of CoverTab[7317]"
	case "notfound", "unavail", "tryagain":
//line /usr/local/go/src/net/nss.go:137
		_go_fuzz_dep_.CoverTab[7318]++
							def = "continue"
//line /usr/local/go/src/net/nss.go:138
		// _ = "end of CoverTab[7318]"
	default:
//line /usr/local/go/src/net/nss.go:139
		_go_fuzz_dep_.CoverTab[7319]++

							return false
//line /usr/local/go/src/net/nss.go:141
		// _ = "end of CoverTab[7319]"
	}
//line /usr/local/go/src/net/nss.go:142
	// _ = "end of CoverTab[7312]"
//line /usr/local/go/src/net/nss.go:142
	_go_fuzz_dep_.CoverTab[7313]++
						if last && func() bool {
//line /usr/local/go/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[7320]++
//line /usr/local/go/src/net/nss.go:143
		return c.action == "return"
//line /usr/local/go/src/net/nss.go:143
		// _ = "end of CoverTab[7320]"
//line /usr/local/go/src/net/nss.go:143
	}() {
//line /usr/local/go/src/net/nss.go:143
		_go_fuzz_dep_.CoverTab[7321]++
							return true
//line /usr/local/go/src/net/nss.go:144
		// _ = "end of CoverTab[7321]"
	} else {
//line /usr/local/go/src/net/nss.go:145
		_go_fuzz_dep_.CoverTab[7322]++
//line /usr/local/go/src/net/nss.go:145
		// _ = "end of CoverTab[7322]"
//line /usr/local/go/src/net/nss.go:145
	}
//line /usr/local/go/src/net/nss.go:145
	// _ = "end of CoverTab[7313]"
//line /usr/local/go/src/net/nss.go:145
	_go_fuzz_dep_.CoverTab[7314]++
						return c.action == def
//line /usr/local/go/src/net/nss.go:146
	// _ = "end of CoverTab[7314]"
}

func parseNSSConfFile(file string) *nssConf {
//line /usr/local/go/src/net/nss.go:149
	_go_fuzz_dep_.CoverTab[7323]++
						f, err := open(file)
						if err != nil {
//line /usr/local/go/src/net/nss.go:151
		_go_fuzz_dep_.CoverTab[7326]++
							return &nssConf{err: err}
//line /usr/local/go/src/net/nss.go:152
		// _ = "end of CoverTab[7326]"
	} else {
//line /usr/local/go/src/net/nss.go:153
		_go_fuzz_dep_.CoverTab[7327]++
//line /usr/local/go/src/net/nss.go:153
		// _ = "end of CoverTab[7327]"
//line /usr/local/go/src/net/nss.go:153
	}
//line /usr/local/go/src/net/nss.go:153
	// _ = "end of CoverTab[7323]"
//line /usr/local/go/src/net/nss.go:153
	_go_fuzz_dep_.CoverTab[7324]++
						defer f.close()
						mtime, _, err := f.stat()
						if err != nil {
//line /usr/local/go/src/net/nss.go:156
		_go_fuzz_dep_.CoverTab[7328]++
							return &nssConf{err: err}
//line /usr/local/go/src/net/nss.go:157
		// _ = "end of CoverTab[7328]"
	} else {
//line /usr/local/go/src/net/nss.go:158
		_go_fuzz_dep_.CoverTab[7329]++
//line /usr/local/go/src/net/nss.go:158
		// _ = "end of CoverTab[7329]"
//line /usr/local/go/src/net/nss.go:158
	}
//line /usr/local/go/src/net/nss.go:158
	// _ = "end of CoverTab[7324]"
//line /usr/local/go/src/net/nss.go:158
	_go_fuzz_dep_.CoverTab[7325]++

						conf := parseNSSConf(f)
						conf.mtime = mtime
						return conf
//line /usr/local/go/src/net/nss.go:162
	// _ = "end of CoverTab[7325]"
}

func parseNSSConf(f *file) *nssConf {
//line /usr/local/go/src/net/nss.go:165
	_go_fuzz_dep_.CoverTab[7330]++
						conf := new(nssConf)
						for line, ok := f.readLine(); ok; line, ok = f.readLine() {
//line /usr/local/go/src/net/nss.go:167
		_go_fuzz_dep_.CoverTab[7332]++
							line = trimSpace(removeComment(line))
							if len(line) == 0 {
//line /usr/local/go/src/net/nss.go:169
			_go_fuzz_dep_.CoverTab[7335]++
								continue
//line /usr/local/go/src/net/nss.go:170
			// _ = "end of CoverTab[7335]"
		} else {
//line /usr/local/go/src/net/nss.go:171
			_go_fuzz_dep_.CoverTab[7336]++
//line /usr/local/go/src/net/nss.go:171
			// _ = "end of CoverTab[7336]"
//line /usr/local/go/src/net/nss.go:171
		}
//line /usr/local/go/src/net/nss.go:171
		// _ = "end of CoverTab[7332]"
//line /usr/local/go/src/net/nss.go:171
		_go_fuzz_dep_.CoverTab[7333]++
							colon := bytealg.IndexByteString(line, ':')
							if colon == -1 {
//line /usr/local/go/src/net/nss.go:173
			_go_fuzz_dep_.CoverTab[7337]++
								conf.err = errors.New("no colon on line")
								return conf
//line /usr/local/go/src/net/nss.go:175
			// _ = "end of CoverTab[7337]"
		} else {
//line /usr/local/go/src/net/nss.go:176
			_go_fuzz_dep_.CoverTab[7338]++
//line /usr/local/go/src/net/nss.go:176
			// _ = "end of CoverTab[7338]"
//line /usr/local/go/src/net/nss.go:176
		}
//line /usr/local/go/src/net/nss.go:176
		// _ = "end of CoverTab[7333]"
//line /usr/local/go/src/net/nss.go:176
		_go_fuzz_dep_.CoverTab[7334]++
							db := trimSpace(line[:colon])
							srcs := line[colon+1:]
							for {
//line /usr/local/go/src/net/nss.go:179
			_go_fuzz_dep_.CoverTab[7339]++
								srcs = trimSpace(srcs)
								if len(srcs) == 0 {
//line /usr/local/go/src/net/nss.go:181
				_go_fuzz_dep_.CoverTab[7344]++
									break
//line /usr/local/go/src/net/nss.go:182
				// _ = "end of CoverTab[7344]"
			} else {
//line /usr/local/go/src/net/nss.go:183
				_go_fuzz_dep_.CoverTab[7345]++
//line /usr/local/go/src/net/nss.go:183
				// _ = "end of CoverTab[7345]"
//line /usr/local/go/src/net/nss.go:183
			}
//line /usr/local/go/src/net/nss.go:183
			// _ = "end of CoverTab[7339]"
//line /usr/local/go/src/net/nss.go:183
			_go_fuzz_dep_.CoverTab[7340]++
								sp := bytealg.IndexByteString(srcs, ' ')
								var src string
								if sp == -1 {
//line /usr/local/go/src/net/nss.go:186
				_go_fuzz_dep_.CoverTab[7346]++
									src = srcs
									srcs = ""
//line /usr/local/go/src/net/nss.go:188
				// _ = "end of CoverTab[7346]"
			} else {
//line /usr/local/go/src/net/nss.go:189
				_go_fuzz_dep_.CoverTab[7347]++
									src = srcs[:sp]
									srcs = trimSpace(srcs[sp+1:])
//line /usr/local/go/src/net/nss.go:191
				// _ = "end of CoverTab[7347]"
			}
//line /usr/local/go/src/net/nss.go:192
			// _ = "end of CoverTab[7340]"
//line /usr/local/go/src/net/nss.go:192
			_go_fuzz_dep_.CoverTab[7341]++
								var criteria []nssCriterion

								if len(srcs) > 0 && func() bool {
//line /usr/local/go/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[7348]++
//line /usr/local/go/src/net/nss.go:195
				return srcs[0] == '['
//line /usr/local/go/src/net/nss.go:195
				// _ = "end of CoverTab[7348]"
//line /usr/local/go/src/net/nss.go:195
			}() {
//line /usr/local/go/src/net/nss.go:195
				_go_fuzz_dep_.CoverTab[7349]++
									bclose := bytealg.IndexByteString(srcs, ']')
									if bclose == -1 {
//line /usr/local/go/src/net/nss.go:197
					_go_fuzz_dep_.CoverTab[7352]++
										conf.err = errors.New("unclosed criterion bracket")
										return conf
//line /usr/local/go/src/net/nss.go:199
					// _ = "end of CoverTab[7352]"
				} else {
//line /usr/local/go/src/net/nss.go:200
					_go_fuzz_dep_.CoverTab[7353]++
//line /usr/local/go/src/net/nss.go:200
					// _ = "end of CoverTab[7353]"
//line /usr/local/go/src/net/nss.go:200
				}
//line /usr/local/go/src/net/nss.go:200
				// _ = "end of CoverTab[7349]"
//line /usr/local/go/src/net/nss.go:200
				_go_fuzz_dep_.CoverTab[7350]++
									var err error
									criteria, err = parseCriteria(srcs[1:bclose])
									if err != nil {
//line /usr/local/go/src/net/nss.go:203
					_go_fuzz_dep_.CoverTab[7354]++
										conf.err = errors.New("invalid criteria: " + srcs[1:bclose])
										return conf
//line /usr/local/go/src/net/nss.go:205
					// _ = "end of CoverTab[7354]"
				} else {
//line /usr/local/go/src/net/nss.go:206
					_go_fuzz_dep_.CoverTab[7355]++
//line /usr/local/go/src/net/nss.go:206
					// _ = "end of CoverTab[7355]"
//line /usr/local/go/src/net/nss.go:206
				}
//line /usr/local/go/src/net/nss.go:206
				// _ = "end of CoverTab[7350]"
//line /usr/local/go/src/net/nss.go:206
				_go_fuzz_dep_.CoverTab[7351]++
									srcs = srcs[bclose+1:]
//line /usr/local/go/src/net/nss.go:207
				// _ = "end of CoverTab[7351]"
			} else {
//line /usr/local/go/src/net/nss.go:208
				_go_fuzz_dep_.CoverTab[7356]++
//line /usr/local/go/src/net/nss.go:208
				// _ = "end of CoverTab[7356]"
//line /usr/local/go/src/net/nss.go:208
			}
//line /usr/local/go/src/net/nss.go:208
			// _ = "end of CoverTab[7341]"
//line /usr/local/go/src/net/nss.go:208
			_go_fuzz_dep_.CoverTab[7342]++
								if conf.sources == nil {
//line /usr/local/go/src/net/nss.go:209
				_go_fuzz_dep_.CoverTab[7357]++
									conf.sources = make(map[string][]nssSource)
//line /usr/local/go/src/net/nss.go:210
				// _ = "end of CoverTab[7357]"
			} else {
//line /usr/local/go/src/net/nss.go:211
				_go_fuzz_dep_.CoverTab[7358]++
//line /usr/local/go/src/net/nss.go:211
				// _ = "end of CoverTab[7358]"
//line /usr/local/go/src/net/nss.go:211
			}
//line /usr/local/go/src/net/nss.go:211
			// _ = "end of CoverTab[7342]"
//line /usr/local/go/src/net/nss.go:211
			_go_fuzz_dep_.CoverTab[7343]++
								conf.sources[db] = append(conf.sources[db], nssSource{
				source:		src,
				criteria:	criteria,
			})
//line /usr/local/go/src/net/nss.go:215
			// _ = "end of CoverTab[7343]"
		}
//line /usr/local/go/src/net/nss.go:216
		// _ = "end of CoverTab[7334]"
	}
//line /usr/local/go/src/net/nss.go:217
	// _ = "end of CoverTab[7330]"
//line /usr/local/go/src/net/nss.go:217
	_go_fuzz_dep_.CoverTab[7331]++
						return conf
//line /usr/local/go/src/net/nss.go:218
	// _ = "end of CoverTab[7331]"
}

// parses "foo=bar !foo=bar"
func parseCriteria(x string) (c []nssCriterion, err error) {
//line /usr/local/go/src/net/nss.go:222
	_go_fuzz_dep_.CoverTab[7359]++
						err = foreachField(x, func(f string) error {
//line /usr/local/go/src/net/nss.go:223
		_go_fuzz_dep_.CoverTab[7361]++
							not := false
							if len(f) > 0 && func() bool {
//line /usr/local/go/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[7366]++
//line /usr/local/go/src/net/nss.go:225
			return f[0] == '!'
//line /usr/local/go/src/net/nss.go:225
			// _ = "end of CoverTab[7366]"
//line /usr/local/go/src/net/nss.go:225
		}() {
//line /usr/local/go/src/net/nss.go:225
			_go_fuzz_dep_.CoverTab[7367]++
								not = true
								f = f[1:]
//line /usr/local/go/src/net/nss.go:227
			// _ = "end of CoverTab[7367]"
		} else {
//line /usr/local/go/src/net/nss.go:228
			_go_fuzz_dep_.CoverTab[7368]++
//line /usr/local/go/src/net/nss.go:228
			// _ = "end of CoverTab[7368]"
//line /usr/local/go/src/net/nss.go:228
		}
//line /usr/local/go/src/net/nss.go:228
		// _ = "end of CoverTab[7361]"
//line /usr/local/go/src/net/nss.go:228
		_go_fuzz_dep_.CoverTab[7362]++
							if len(f) < 3 {
//line /usr/local/go/src/net/nss.go:229
			_go_fuzz_dep_.CoverTab[7369]++
								return errors.New("criterion too short")
//line /usr/local/go/src/net/nss.go:230
			// _ = "end of CoverTab[7369]"
		} else {
//line /usr/local/go/src/net/nss.go:231
			_go_fuzz_dep_.CoverTab[7370]++
//line /usr/local/go/src/net/nss.go:231
			// _ = "end of CoverTab[7370]"
//line /usr/local/go/src/net/nss.go:231
		}
//line /usr/local/go/src/net/nss.go:231
		// _ = "end of CoverTab[7362]"
//line /usr/local/go/src/net/nss.go:231
		_go_fuzz_dep_.CoverTab[7363]++
							eq := bytealg.IndexByteString(f, '=')
							if eq == -1 {
//line /usr/local/go/src/net/nss.go:233
			_go_fuzz_dep_.CoverTab[7371]++
								return errors.New("criterion lacks equal sign")
//line /usr/local/go/src/net/nss.go:234
			// _ = "end of CoverTab[7371]"
		} else {
//line /usr/local/go/src/net/nss.go:235
			_go_fuzz_dep_.CoverTab[7372]++
//line /usr/local/go/src/net/nss.go:235
			// _ = "end of CoverTab[7372]"
//line /usr/local/go/src/net/nss.go:235
		}
//line /usr/local/go/src/net/nss.go:235
		// _ = "end of CoverTab[7363]"
//line /usr/local/go/src/net/nss.go:235
		_go_fuzz_dep_.CoverTab[7364]++
							if hasUpperCase(f) {
//line /usr/local/go/src/net/nss.go:236
			_go_fuzz_dep_.CoverTab[7373]++
								lower := []byte(f)
								lowerASCIIBytes(lower)
								f = string(lower)
//line /usr/local/go/src/net/nss.go:239
			// _ = "end of CoverTab[7373]"
		} else {
//line /usr/local/go/src/net/nss.go:240
			_go_fuzz_dep_.CoverTab[7374]++
//line /usr/local/go/src/net/nss.go:240
			// _ = "end of CoverTab[7374]"
//line /usr/local/go/src/net/nss.go:240
		}
//line /usr/local/go/src/net/nss.go:240
		// _ = "end of CoverTab[7364]"
//line /usr/local/go/src/net/nss.go:240
		_go_fuzz_dep_.CoverTab[7365]++
							c = append(c, nssCriterion{
			negate:	not,
			status:	f[:eq],
			action:	f[eq+1:],
		})
							return nil
//line /usr/local/go/src/net/nss.go:246
		// _ = "end of CoverTab[7365]"
	})
//line /usr/local/go/src/net/nss.go:247
	// _ = "end of CoverTab[7359]"
//line /usr/local/go/src/net/nss.go:247
	_go_fuzz_dep_.CoverTab[7360]++
						return
//line /usr/local/go/src/net/nss.go:248
	// _ = "end of CoverTab[7360]"
}

//line /usr/local/go/src/net/nss.go:249
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/nss.go:249
var _ = _go_fuzz_dep_.CoverTab
