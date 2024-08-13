// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/os/signal/signal.go:5
package signal

//line /usr/local/go/src/os/signal/signal.go:5
import (
//line /usr/local/go/src/os/signal/signal.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/signal/signal.go:5
)
//line /usr/local/go/src/os/signal/signal.go:5
import (
//line /usr/local/go/src/os/signal/signal.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/signal/signal.go:5
)

import (
	"context"
	"os"
	"sync"
)

var handlers struct {
	sync.Mutex
	// Map a channel to the signals that should be sent to it.
	m	map[chan<- os.Signal]*handler
	// Map a signal to the number of channels receiving it.
	ref	[numSig]int64
	// Map channels to signals while the channel is being stopped.
	// Not a map because entries live here only very briefly.
	// We need a separate container because we need m to correspond to ref
	// at all times, and we also need to keep track of the *handler
	// value for a channel being stopped. See the Stop function.
	stopping	[]stopping
}

type stopping struct {
	c	chan<- os.Signal
	h	*handler
}

type handler struct {
	mask [(numSig + 31) / 32]uint32
}

func (h *handler) want(sig int) bool {
//line /usr/local/go/src/os/signal/signal.go:36
	_go_fuzz_dep_.CoverTab[197109]++
							return (h.mask[sig/32]>>uint(sig&31))&1 != 0
//line /usr/local/go/src/os/signal/signal.go:37
	// _ = "end of CoverTab[197109]"
}

func (h *handler) set(sig int) {
//line /usr/local/go/src/os/signal/signal.go:40
	_go_fuzz_dep_.CoverTab[197110]++
							h.mask[sig/32] |= 1 << uint(sig&31)
//line /usr/local/go/src/os/signal/signal.go:41
	// _ = "end of CoverTab[197110]"
}

func (h *handler) clear(sig int) {
//line /usr/local/go/src/os/signal/signal.go:44
	_go_fuzz_dep_.CoverTab[197111]++
							h.mask[sig/32] &^= 1 << uint(sig&31)
//line /usr/local/go/src/os/signal/signal.go:45
	// _ = "end of CoverTab[197111]"
}

// Stop relaying the signals, sigs, to any channels previously registered to
//line /usr/local/go/src/os/signal/signal.go:48
// receive them and either reset the signal handlers to their original values
//line /usr/local/go/src/os/signal/signal.go:48
// (action=disableSignal) or ignore the signals (action=ignoreSignal).
//line /usr/local/go/src/os/signal/signal.go:51
func cancel(sigs []os.Signal, action func(int)) {
//line /usr/local/go/src/os/signal/signal.go:51
	_go_fuzz_dep_.CoverTab[197112]++
							handlers.Lock()
							defer handlers.Unlock()

							remove := func(n int) {
//line /usr/local/go/src/os/signal/signal.go:55
		_go_fuzz_dep_.CoverTab[197114]++
								var zerohandler handler

								for c, h := range handlers.m {
//line /usr/local/go/src/os/signal/signal.go:58
			_go_fuzz_dep_.CoverTab[197116]++
									if h.want(n) {
//line /usr/local/go/src/os/signal/signal.go:59
				_go_fuzz_dep_.CoverTab[197117]++
										handlers.ref[n]--
										h.clear(n)
										if h.mask == zerohandler.mask {
//line /usr/local/go/src/os/signal/signal.go:62
					_go_fuzz_dep_.CoverTab[197118]++
											delete(handlers.m, c)
//line /usr/local/go/src/os/signal/signal.go:63
					// _ = "end of CoverTab[197118]"
				} else {
//line /usr/local/go/src/os/signal/signal.go:64
					_go_fuzz_dep_.CoverTab[197119]++
//line /usr/local/go/src/os/signal/signal.go:64
					// _ = "end of CoverTab[197119]"
//line /usr/local/go/src/os/signal/signal.go:64
				}
//line /usr/local/go/src/os/signal/signal.go:64
				// _ = "end of CoverTab[197117]"
			} else {
//line /usr/local/go/src/os/signal/signal.go:65
				_go_fuzz_dep_.CoverTab[197120]++
//line /usr/local/go/src/os/signal/signal.go:65
				// _ = "end of CoverTab[197120]"
//line /usr/local/go/src/os/signal/signal.go:65
			}
//line /usr/local/go/src/os/signal/signal.go:65
			// _ = "end of CoverTab[197116]"
		}
//line /usr/local/go/src/os/signal/signal.go:66
		// _ = "end of CoverTab[197114]"
//line /usr/local/go/src/os/signal/signal.go:66
		_go_fuzz_dep_.CoverTab[197115]++

								action(n)
//line /usr/local/go/src/os/signal/signal.go:68
		// _ = "end of CoverTab[197115]"
	}
//line /usr/local/go/src/os/signal/signal.go:69
	// _ = "end of CoverTab[197112]"
//line /usr/local/go/src/os/signal/signal.go:69
	_go_fuzz_dep_.CoverTab[197113]++

							if len(sigs) == 0 {
//line /usr/local/go/src/os/signal/signal.go:71
		_go_fuzz_dep_.CoverTab[197121]++
								for n := 0; n < numSig; n++ {
//line /usr/local/go/src/os/signal/signal.go:72
			_go_fuzz_dep_.CoverTab[197122]++
									remove(n)
//line /usr/local/go/src/os/signal/signal.go:73
			// _ = "end of CoverTab[197122]"
		}
//line /usr/local/go/src/os/signal/signal.go:74
		// _ = "end of CoverTab[197121]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:75
		_go_fuzz_dep_.CoverTab[197123]++
								for _, s := range sigs {
//line /usr/local/go/src/os/signal/signal.go:76
			_go_fuzz_dep_.CoverTab[197124]++
									remove(signum(s))
//line /usr/local/go/src/os/signal/signal.go:77
			// _ = "end of CoverTab[197124]"
		}
//line /usr/local/go/src/os/signal/signal.go:78
		// _ = "end of CoverTab[197123]"
	}
//line /usr/local/go/src/os/signal/signal.go:79
	// _ = "end of CoverTab[197113]"
}

// Ignore causes the provided signals to be ignored. If they are received by
//line /usr/local/go/src/os/signal/signal.go:82
// the program, nothing will happen. Ignore undoes the effect of any prior
//line /usr/local/go/src/os/signal/signal.go:82
// calls to Notify for the provided signals.
//line /usr/local/go/src/os/signal/signal.go:82
// If no signals are provided, all incoming signals will be ignored.
//line /usr/local/go/src/os/signal/signal.go:86
func Ignore(sig ...os.Signal) {
//line /usr/local/go/src/os/signal/signal.go:86
	_go_fuzz_dep_.CoverTab[197125]++
							cancel(sig, ignoreSignal)
//line /usr/local/go/src/os/signal/signal.go:87
	// _ = "end of CoverTab[197125]"
}

// Ignored reports whether sig is currently ignored.
func Ignored(sig os.Signal) bool {
//line /usr/local/go/src/os/signal/signal.go:91
	_go_fuzz_dep_.CoverTab[197126]++
							sn := signum(sig)
							return sn >= 0 && func() bool {
//line /usr/local/go/src/os/signal/signal.go:93
		_go_fuzz_dep_.CoverTab[197127]++
//line /usr/local/go/src/os/signal/signal.go:93
		return signalIgnored(sn)
//line /usr/local/go/src/os/signal/signal.go:93
		// _ = "end of CoverTab[197127]"
//line /usr/local/go/src/os/signal/signal.go:93
	}()
//line /usr/local/go/src/os/signal/signal.go:93
	// _ = "end of CoverTab[197126]"
}

var (
	// watchSignalLoopOnce guards calling the conditionally
	// initialized watchSignalLoop. If watchSignalLoop is non-nil,
	// it will be run in a goroutine lazily once Notify is invoked.
	// See Issue 21576.
	watchSignalLoopOnce	sync.Once
	watchSignalLoop		func()
)

// Notify causes package signal to relay incoming signals to c.
//line /usr/local/go/src/os/signal/signal.go:105
// If no signals are provided, all incoming signals will be relayed to c.
//line /usr/local/go/src/os/signal/signal.go:105
// Otherwise, just the provided signals will.
//line /usr/local/go/src/os/signal/signal.go:105
//
//line /usr/local/go/src/os/signal/signal.go:105
// Package signal will not block sending to c: the caller must ensure
//line /usr/local/go/src/os/signal/signal.go:105
// that c has sufficient buffer space to keep up with the expected
//line /usr/local/go/src/os/signal/signal.go:105
// signal rate. For a channel used for notification of just one signal value,
//line /usr/local/go/src/os/signal/signal.go:105
// a buffer of size 1 is sufficient.
//line /usr/local/go/src/os/signal/signal.go:105
//
//line /usr/local/go/src/os/signal/signal.go:105
// It is allowed to call Notify multiple times with the same channel:
//line /usr/local/go/src/os/signal/signal.go:105
// each call expands the set of signals sent to that channel.
//line /usr/local/go/src/os/signal/signal.go:105
// The only way to remove signals from the set is to call Stop.
//line /usr/local/go/src/os/signal/signal.go:105
//
//line /usr/local/go/src/os/signal/signal.go:105
// It is allowed to call Notify multiple times with different channels
//line /usr/local/go/src/os/signal/signal.go:105
// and the same signals: each channel receives copies of incoming
//line /usr/local/go/src/os/signal/signal.go:105
// signals independently.
//line /usr/local/go/src/os/signal/signal.go:121
func Notify(c chan<- os.Signal, sig ...os.Signal) {
//line /usr/local/go/src/os/signal/signal.go:121
	_go_fuzz_dep_.CoverTab[197128]++
							if c == nil {
//line /usr/local/go/src/os/signal/signal.go:122
		_go_fuzz_dep_.CoverTab[197132]++
								panic("os/signal: Notify using nil channel")
//line /usr/local/go/src/os/signal/signal.go:123
		// _ = "end of CoverTab[197132]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:124
		_go_fuzz_dep_.CoverTab[197133]++
//line /usr/local/go/src/os/signal/signal.go:124
		// _ = "end of CoverTab[197133]"
//line /usr/local/go/src/os/signal/signal.go:124
	}
//line /usr/local/go/src/os/signal/signal.go:124
	// _ = "end of CoverTab[197128]"
//line /usr/local/go/src/os/signal/signal.go:124
	_go_fuzz_dep_.CoverTab[197129]++

							handlers.Lock()
							defer handlers.Unlock()

							h := handlers.m[c]
							if h == nil {
//line /usr/local/go/src/os/signal/signal.go:130
		_go_fuzz_dep_.CoverTab[197134]++
								if handlers.m == nil {
//line /usr/local/go/src/os/signal/signal.go:131
			_go_fuzz_dep_.CoverTab[197136]++
									handlers.m = make(map[chan<- os.Signal]*handler)
//line /usr/local/go/src/os/signal/signal.go:132
			// _ = "end of CoverTab[197136]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:133
			_go_fuzz_dep_.CoverTab[197137]++
//line /usr/local/go/src/os/signal/signal.go:133
			// _ = "end of CoverTab[197137]"
//line /usr/local/go/src/os/signal/signal.go:133
		}
//line /usr/local/go/src/os/signal/signal.go:133
		// _ = "end of CoverTab[197134]"
//line /usr/local/go/src/os/signal/signal.go:133
		_go_fuzz_dep_.CoverTab[197135]++
								h = new(handler)
								handlers.m[c] = h
//line /usr/local/go/src/os/signal/signal.go:135
		// _ = "end of CoverTab[197135]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:136
		_go_fuzz_dep_.CoverTab[197138]++
//line /usr/local/go/src/os/signal/signal.go:136
		// _ = "end of CoverTab[197138]"
//line /usr/local/go/src/os/signal/signal.go:136
	}
//line /usr/local/go/src/os/signal/signal.go:136
	// _ = "end of CoverTab[197129]"
//line /usr/local/go/src/os/signal/signal.go:136
	_go_fuzz_dep_.CoverTab[197130]++

							add := func(n int) {
//line /usr/local/go/src/os/signal/signal.go:138
		_go_fuzz_dep_.CoverTab[197139]++
								if n < 0 {
//line /usr/local/go/src/os/signal/signal.go:139
			_go_fuzz_dep_.CoverTab[197141]++
									return
//line /usr/local/go/src/os/signal/signal.go:140
			// _ = "end of CoverTab[197141]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:141
			_go_fuzz_dep_.CoverTab[197142]++
//line /usr/local/go/src/os/signal/signal.go:141
			// _ = "end of CoverTab[197142]"
//line /usr/local/go/src/os/signal/signal.go:141
		}
//line /usr/local/go/src/os/signal/signal.go:141
		// _ = "end of CoverTab[197139]"
//line /usr/local/go/src/os/signal/signal.go:141
		_go_fuzz_dep_.CoverTab[197140]++
								if !h.want(n) {
//line /usr/local/go/src/os/signal/signal.go:142
			_go_fuzz_dep_.CoverTab[197143]++
									h.set(n)
									if handlers.ref[n] == 0 {
//line /usr/local/go/src/os/signal/signal.go:144
				_go_fuzz_dep_.CoverTab[197145]++
										enableSignal(n)

//line /usr/local/go/src/os/signal/signal.go:149
				watchSignalLoopOnce.Do(func() {
//line /usr/local/go/src/os/signal/signal.go:149
					_go_fuzz_dep_.CoverTab[197146]++
											if watchSignalLoop != nil {
//line /usr/local/go/src/os/signal/signal.go:150
						_go_fuzz_dep_.CoverTab[197147]++
//line /usr/local/go/src/os/signal/signal.go:150
						_curRoutineNum199_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/os/signal/signal.go:150
						_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum199_)
												go func() {
//line /usr/local/go/src/os/signal/signal.go:151
							_go_fuzz_dep_.CoverTab[197148]++
//line /usr/local/go/src/os/signal/signal.go:151
							defer func() {
//line /usr/local/go/src/os/signal/signal.go:151
								_go_fuzz_dep_.CoverTab[197149]++
//line /usr/local/go/src/os/signal/signal.go:151
								_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum199_)
//line /usr/local/go/src/os/signal/signal.go:151
								// _ = "end of CoverTab[197149]"
//line /usr/local/go/src/os/signal/signal.go:151
							}()
//line /usr/local/go/src/os/signal/signal.go:151
							watchSignalLoop()
//line /usr/local/go/src/os/signal/signal.go:151
							// _ = "end of CoverTab[197148]"
//line /usr/local/go/src/os/signal/signal.go:151
						}()
//line /usr/local/go/src/os/signal/signal.go:151
						// _ = "end of CoverTab[197147]"
					} else {
//line /usr/local/go/src/os/signal/signal.go:152
						_go_fuzz_dep_.CoverTab[197150]++
//line /usr/local/go/src/os/signal/signal.go:152
						// _ = "end of CoverTab[197150]"
//line /usr/local/go/src/os/signal/signal.go:152
					}
//line /usr/local/go/src/os/signal/signal.go:152
					// _ = "end of CoverTab[197146]"
				})
//line /usr/local/go/src/os/signal/signal.go:153
				// _ = "end of CoverTab[197145]"
			} else {
//line /usr/local/go/src/os/signal/signal.go:154
				_go_fuzz_dep_.CoverTab[197151]++
//line /usr/local/go/src/os/signal/signal.go:154
				// _ = "end of CoverTab[197151]"
//line /usr/local/go/src/os/signal/signal.go:154
			}
//line /usr/local/go/src/os/signal/signal.go:154
			// _ = "end of CoverTab[197143]"
//line /usr/local/go/src/os/signal/signal.go:154
			_go_fuzz_dep_.CoverTab[197144]++
									handlers.ref[n]++
//line /usr/local/go/src/os/signal/signal.go:155
			// _ = "end of CoverTab[197144]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:156
			_go_fuzz_dep_.CoverTab[197152]++
//line /usr/local/go/src/os/signal/signal.go:156
			// _ = "end of CoverTab[197152]"
//line /usr/local/go/src/os/signal/signal.go:156
		}
//line /usr/local/go/src/os/signal/signal.go:156
		// _ = "end of CoverTab[197140]"
	}
//line /usr/local/go/src/os/signal/signal.go:157
	// _ = "end of CoverTab[197130]"
//line /usr/local/go/src/os/signal/signal.go:157
	_go_fuzz_dep_.CoverTab[197131]++

							if len(sig) == 0 {
//line /usr/local/go/src/os/signal/signal.go:159
		_go_fuzz_dep_.CoverTab[197153]++
								for n := 0; n < numSig; n++ {
//line /usr/local/go/src/os/signal/signal.go:160
			_go_fuzz_dep_.CoverTab[197154]++
									add(n)
//line /usr/local/go/src/os/signal/signal.go:161
			// _ = "end of CoverTab[197154]"
		}
//line /usr/local/go/src/os/signal/signal.go:162
		// _ = "end of CoverTab[197153]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:163
		_go_fuzz_dep_.CoverTab[197155]++
								for _, s := range sig {
//line /usr/local/go/src/os/signal/signal.go:164
			_go_fuzz_dep_.CoverTab[197156]++
									add(signum(s))
//line /usr/local/go/src/os/signal/signal.go:165
			// _ = "end of CoverTab[197156]"
		}
//line /usr/local/go/src/os/signal/signal.go:166
		// _ = "end of CoverTab[197155]"
	}
//line /usr/local/go/src/os/signal/signal.go:167
	// _ = "end of CoverTab[197131]"
}

// Reset undoes the effect of any prior calls to Notify for the provided
//line /usr/local/go/src/os/signal/signal.go:170
// signals.
//line /usr/local/go/src/os/signal/signal.go:170
// If no signals are provided, all signal handlers will be reset.
//line /usr/local/go/src/os/signal/signal.go:173
func Reset(sig ...os.Signal) {
//line /usr/local/go/src/os/signal/signal.go:173
	_go_fuzz_dep_.CoverTab[197157]++
							cancel(sig, disableSignal)
//line /usr/local/go/src/os/signal/signal.go:174
	// _ = "end of CoverTab[197157]"
}

// Stop causes package signal to stop relaying incoming signals to c.
//line /usr/local/go/src/os/signal/signal.go:177
// It undoes the effect of all prior calls to Notify using c.
//line /usr/local/go/src/os/signal/signal.go:177
// When Stop returns, it is guaranteed that c will receive no more signals.
//line /usr/local/go/src/os/signal/signal.go:180
func Stop(c chan<- os.Signal) {
//line /usr/local/go/src/os/signal/signal.go:180
	_go_fuzz_dep_.CoverTab[197158]++
							handlers.Lock()

							h := handlers.m[c]
							if h == nil {
//line /usr/local/go/src/os/signal/signal.go:184
		_go_fuzz_dep_.CoverTab[197162]++
								handlers.Unlock()
								return
//line /usr/local/go/src/os/signal/signal.go:186
		// _ = "end of CoverTab[197162]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:187
		_go_fuzz_dep_.CoverTab[197163]++
//line /usr/local/go/src/os/signal/signal.go:187
		// _ = "end of CoverTab[197163]"
//line /usr/local/go/src/os/signal/signal.go:187
	}
//line /usr/local/go/src/os/signal/signal.go:187
	// _ = "end of CoverTab[197158]"
//line /usr/local/go/src/os/signal/signal.go:187
	_go_fuzz_dep_.CoverTab[197159]++
							delete(handlers.m, c)

							for n := 0; n < numSig; n++ {
//line /usr/local/go/src/os/signal/signal.go:190
		_go_fuzz_dep_.CoverTab[197164]++
								if h.want(n) {
//line /usr/local/go/src/os/signal/signal.go:191
			_go_fuzz_dep_.CoverTab[197165]++
									handlers.ref[n]--
									if handlers.ref[n] == 0 {
//line /usr/local/go/src/os/signal/signal.go:193
				_go_fuzz_dep_.CoverTab[197166]++
										disableSignal(n)
//line /usr/local/go/src/os/signal/signal.go:194
				// _ = "end of CoverTab[197166]"
			} else {
//line /usr/local/go/src/os/signal/signal.go:195
				_go_fuzz_dep_.CoverTab[197167]++
//line /usr/local/go/src/os/signal/signal.go:195
				// _ = "end of CoverTab[197167]"
//line /usr/local/go/src/os/signal/signal.go:195
			}
//line /usr/local/go/src/os/signal/signal.go:195
			// _ = "end of CoverTab[197165]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:196
			_go_fuzz_dep_.CoverTab[197168]++
//line /usr/local/go/src/os/signal/signal.go:196
			// _ = "end of CoverTab[197168]"
//line /usr/local/go/src/os/signal/signal.go:196
		}
//line /usr/local/go/src/os/signal/signal.go:196
		// _ = "end of CoverTab[197164]"
	}
//line /usr/local/go/src/os/signal/signal.go:197
	// _ = "end of CoverTab[197159]"
//line /usr/local/go/src/os/signal/signal.go:197
	_go_fuzz_dep_.CoverTab[197160]++

//line /usr/local/go/src/os/signal/signal.go:210
	handlers.stopping = append(handlers.stopping, stopping{c, h})

	handlers.Unlock()

	signalWaitUntilIdle()

	handlers.Lock()

	for i, s := range handlers.stopping {
//line /usr/local/go/src/os/signal/signal.go:218
		_go_fuzz_dep_.CoverTab[197169]++
								if s.c == c {
//line /usr/local/go/src/os/signal/signal.go:219
			_go_fuzz_dep_.CoverTab[197170]++
									handlers.stopping = append(handlers.stopping[:i], handlers.stopping[i+1:]...)
									break
//line /usr/local/go/src/os/signal/signal.go:221
			// _ = "end of CoverTab[197170]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:222
			_go_fuzz_dep_.CoverTab[197171]++
//line /usr/local/go/src/os/signal/signal.go:222
			// _ = "end of CoverTab[197171]"
//line /usr/local/go/src/os/signal/signal.go:222
		}
//line /usr/local/go/src/os/signal/signal.go:222
		// _ = "end of CoverTab[197169]"
	}
//line /usr/local/go/src/os/signal/signal.go:223
	// _ = "end of CoverTab[197160]"
//line /usr/local/go/src/os/signal/signal.go:223
	_go_fuzz_dep_.CoverTab[197161]++

							handlers.Unlock()
//line /usr/local/go/src/os/signal/signal.go:225
	// _ = "end of CoverTab[197161]"
}

// Wait until there are no more signals waiting to be delivered.
//line /usr/local/go/src/os/signal/signal.go:228
// Defined by the runtime package.
//line /usr/local/go/src/os/signal/signal.go:230
func signalWaitUntilIdle()

func process(sig os.Signal) {
//line /usr/local/go/src/os/signal/signal.go:232
	_go_fuzz_dep_.CoverTab[197172]++
							n := signum(sig)
							if n < 0 {
//line /usr/local/go/src/os/signal/signal.go:234
		_go_fuzz_dep_.CoverTab[197175]++
								return
//line /usr/local/go/src/os/signal/signal.go:235
		// _ = "end of CoverTab[197175]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:236
		_go_fuzz_dep_.CoverTab[197176]++
//line /usr/local/go/src/os/signal/signal.go:236
		// _ = "end of CoverTab[197176]"
//line /usr/local/go/src/os/signal/signal.go:236
	}
//line /usr/local/go/src/os/signal/signal.go:236
	// _ = "end of CoverTab[197172]"
//line /usr/local/go/src/os/signal/signal.go:236
	_go_fuzz_dep_.CoverTab[197173]++

							handlers.Lock()
							defer handlers.Unlock()

							for c, h := range handlers.m {
//line /usr/local/go/src/os/signal/signal.go:241
		_go_fuzz_dep_.CoverTab[197177]++
								if h.want(n) {
//line /usr/local/go/src/os/signal/signal.go:242
			_go_fuzz_dep_.CoverTab[197178]++

									select {
			case c <- sig:
//line /usr/local/go/src/os/signal/signal.go:245
				_go_fuzz_dep_.CoverTab[197179]++
//line /usr/local/go/src/os/signal/signal.go:245
				// _ = "end of CoverTab[197179]"
			default:
//line /usr/local/go/src/os/signal/signal.go:246
				_go_fuzz_dep_.CoverTab[197180]++
//line /usr/local/go/src/os/signal/signal.go:246
				// _ = "end of CoverTab[197180]"
			}
//line /usr/local/go/src/os/signal/signal.go:247
			// _ = "end of CoverTab[197178]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:248
			_go_fuzz_dep_.CoverTab[197181]++
//line /usr/local/go/src/os/signal/signal.go:248
			// _ = "end of CoverTab[197181]"
//line /usr/local/go/src/os/signal/signal.go:248
		}
//line /usr/local/go/src/os/signal/signal.go:248
		// _ = "end of CoverTab[197177]"
	}
//line /usr/local/go/src/os/signal/signal.go:249
	// _ = "end of CoverTab[197173]"
//line /usr/local/go/src/os/signal/signal.go:249
	_go_fuzz_dep_.CoverTab[197174]++

//line /usr/local/go/src/os/signal/signal.go:252
	for _, d := range handlers.stopping {
//line /usr/local/go/src/os/signal/signal.go:252
		_go_fuzz_dep_.CoverTab[197182]++
								if d.h.want(n) {
//line /usr/local/go/src/os/signal/signal.go:253
			_go_fuzz_dep_.CoverTab[197183]++
									select {
			case d.c <- sig:
//line /usr/local/go/src/os/signal/signal.go:255
				_go_fuzz_dep_.CoverTab[197184]++
//line /usr/local/go/src/os/signal/signal.go:255
				// _ = "end of CoverTab[197184]"
			default:
//line /usr/local/go/src/os/signal/signal.go:256
				_go_fuzz_dep_.CoverTab[197185]++
//line /usr/local/go/src/os/signal/signal.go:256
				// _ = "end of CoverTab[197185]"
			}
//line /usr/local/go/src/os/signal/signal.go:257
			// _ = "end of CoverTab[197183]"
		} else {
//line /usr/local/go/src/os/signal/signal.go:258
			_go_fuzz_dep_.CoverTab[197186]++
//line /usr/local/go/src/os/signal/signal.go:258
			// _ = "end of CoverTab[197186]"
//line /usr/local/go/src/os/signal/signal.go:258
		}
//line /usr/local/go/src/os/signal/signal.go:258
		// _ = "end of CoverTab[197182]"
	}
//line /usr/local/go/src/os/signal/signal.go:259
	// _ = "end of CoverTab[197174]"
}

// NotifyContext returns a copy of the parent context that is marked done
//line /usr/local/go/src/os/signal/signal.go:262
// (its Done channel is closed) when one of the listed signals arrives,
//line /usr/local/go/src/os/signal/signal.go:262
// when the returned stop function is called, or when the parent context's
//line /usr/local/go/src/os/signal/signal.go:262
// Done channel is closed, whichever happens first.
//line /usr/local/go/src/os/signal/signal.go:262
//
//line /usr/local/go/src/os/signal/signal.go:262
// The stop function unregisters the signal behavior, which, like signal.Reset,
//line /usr/local/go/src/os/signal/signal.go:262
// may restore the default behavior for a given signal. For example, the default
//line /usr/local/go/src/os/signal/signal.go:262
// behavior of a Go program receiving os.Interrupt is to exit. Calling
//line /usr/local/go/src/os/signal/signal.go:262
// NotifyContext(parent, os.Interrupt) will change the behavior to cancel
//line /usr/local/go/src/os/signal/signal.go:262
// the returned context. Future interrupts received will not trigger the default
//line /usr/local/go/src/os/signal/signal.go:262
// (exit) behavior until the returned stop function is called.
//line /usr/local/go/src/os/signal/signal.go:262
//
//line /usr/local/go/src/os/signal/signal.go:262
// The stop function releases resources associated with it, so code should
//line /usr/local/go/src/os/signal/signal.go:262
// call stop as soon as the operations running in this Context complete and
//line /usr/local/go/src/os/signal/signal.go:262
// signals no longer need to be diverted to the context.
//line /usr/local/go/src/os/signal/signal.go:277
func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc) {
//line /usr/local/go/src/os/signal/signal.go:277
	_go_fuzz_dep_.CoverTab[197187]++
							ctx, cancel := context.WithCancel(parent)
							c := &signalCtx{
		Context:	ctx,
		cancel:		cancel,
		signals:	signals,
	}
	c.ch = make(chan os.Signal, 1)
	Notify(c.ch, c.signals...)
	if ctx.Err() == nil {
//line /usr/local/go/src/os/signal/signal.go:286
		_go_fuzz_dep_.CoverTab[197189]++
//line /usr/local/go/src/os/signal/signal.go:286
		_curRoutineNum200_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/os/signal/signal.go:286
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum200_)
								go func() {
//line /usr/local/go/src/os/signal/signal.go:287
			_go_fuzz_dep_.CoverTab[197190]++
//line /usr/local/go/src/os/signal/signal.go:287
			defer func() {
//line /usr/local/go/src/os/signal/signal.go:287
				_go_fuzz_dep_.CoverTab[197191]++
//line /usr/local/go/src/os/signal/signal.go:287
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum200_)
//line /usr/local/go/src/os/signal/signal.go:287
				// _ = "end of CoverTab[197191]"
//line /usr/local/go/src/os/signal/signal.go:287
			}()
									select {
			case <-c.ch:
//line /usr/local/go/src/os/signal/signal.go:289
				_go_fuzz_dep_.CoverTab[197192]++
										c.cancel()
//line /usr/local/go/src/os/signal/signal.go:290
				// _ = "end of CoverTab[197192]"
			case <-c.Done():
//line /usr/local/go/src/os/signal/signal.go:291
				_go_fuzz_dep_.CoverTab[197193]++
//line /usr/local/go/src/os/signal/signal.go:291
				// _ = "end of CoverTab[197193]"
			}
//line /usr/local/go/src/os/signal/signal.go:292
			// _ = "end of CoverTab[197190]"
		}()
//line /usr/local/go/src/os/signal/signal.go:293
		// _ = "end of CoverTab[197189]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:294
		_go_fuzz_dep_.CoverTab[197194]++
//line /usr/local/go/src/os/signal/signal.go:294
		// _ = "end of CoverTab[197194]"
//line /usr/local/go/src/os/signal/signal.go:294
	}
//line /usr/local/go/src/os/signal/signal.go:294
	// _ = "end of CoverTab[197187]"
//line /usr/local/go/src/os/signal/signal.go:294
	_go_fuzz_dep_.CoverTab[197188]++
							return c, c.stop
//line /usr/local/go/src/os/signal/signal.go:295
	// _ = "end of CoverTab[197188]"
}

type signalCtx struct {
	context.Context

	cancel	context.CancelFunc
	signals	[]os.Signal
	ch	chan os.Signal
}

func (c *signalCtx) stop() {
//line /usr/local/go/src/os/signal/signal.go:306
	_go_fuzz_dep_.CoverTab[197195]++
							c.cancel()
							Stop(c.ch)
//line /usr/local/go/src/os/signal/signal.go:308
	// _ = "end of CoverTab[197195]"
}

type stringer interface {
	String() string
}

func (c *signalCtx) String() string {
//line /usr/local/go/src/os/signal/signal.go:315
	_go_fuzz_dep_.CoverTab[197196]++
							var buf []byte

//line /usr/local/go/src/os/signal/signal.go:319
	name := c.Context.(stringer).String()
	name = name[:len(name)-len(".WithCancel")]
	buf = append(buf, "signal.NotifyContext("+name...)
	if len(c.signals) != 0 {
//line /usr/local/go/src/os/signal/signal.go:322
		_go_fuzz_dep_.CoverTab[197198]++
								buf = append(buf, ", ["...)
								for i, s := range c.signals {
//line /usr/local/go/src/os/signal/signal.go:324
			_go_fuzz_dep_.CoverTab[197200]++
									buf = append(buf, s.String()...)
									if i != len(c.signals)-1 {
//line /usr/local/go/src/os/signal/signal.go:326
				_go_fuzz_dep_.CoverTab[197201]++
										buf = append(buf, ' ')
//line /usr/local/go/src/os/signal/signal.go:327
				// _ = "end of CoverTab[197201]"
			} else {
//line /usr/local/go/src/os/signal/signal.go:328
				_go_fuzz_dep_.CoverTab[197202]++
//line /usr/local/go/src/os/signal/signal.go:328
				// _ = "end of CoverTab[197202]"
//line /usr/local/go/src/os/signal/signal.go:328
			}
//line /usr/local/go/src/os/signal/signal.go:328
			// _ = "end of CoverTab[197200]"
		}
//line /usr/local/go/src/os/signal/signal.go:329
		// _ = "end of CoverTab[197198]"
//line /usr/local/go/src/os/signal/signal.go:329
		_go_fuzz_dep_.CoverTab[197199]++
								buf = append(buf, ']')
//line /usr/local/go/src/os/signal/signal.go:330
		// _ = "end of CoverTab[197199]"
	} else {
//line /usr/local/go/src/os/signal/signal.go:331
		_go_fuzz_dep_.CoverTab[197203]++
//line /usr/local/go/src/os/signal/signal.go:331
		// _ = "end of CoverTab[197203]"
//line /usr/local/go/src/os/signal/signal.go:331
	}
//line /usr/local/go/src/os/signal/signal.go:331
	// _ = "end of CoverTab[197196]"
//line /usr/local/go/src/os/signal/signal.go:331
	_go_fuzz_dep_.CoverTab[197197]++
							buf = append(buf, ')')
							return string(buf)
//line /usr/local/go/src/os/signal/signal.go:333
	// _ = "end of CoverTab[197197]"
}

//line /usr/local/go/src/os/signal/signal.go:334
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/signal/signal.go:334
var _ = _go_fuzz_dep_.CoverTab
