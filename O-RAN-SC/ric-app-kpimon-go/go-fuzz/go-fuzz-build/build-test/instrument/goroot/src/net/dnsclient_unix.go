// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

// DNS client: see RFC 1035.
// Has to be linked into package net for Dial.

// TODO(rsc):
//	Could potentially handle many outstanding lookups faster.
//	Random UDP source port (net.Dial should do that for us).
//	Random request IDs.

//line /snap/go/10455/src/net/dnsclient_unix.go:15
package net

//line /snap/go/10455/src/net/dnsclient_unix.go:15
import (
//line /snap/go/10455/src/net/dnsclient_unix.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/dnsclient_unix.go:15
)
//line /snap/go/10455/src/net/dnsclient_unix.go:15
import (
//line /snap/go/10455/src/net/dnsclient_unix.go:15
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/dnsclient_unix.go:15
)

import (
	"context"
	"errors"
	"internal/itoa"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/dns/dnsmessage"
)

const (
	// to be used as a useTCP parameter to exchange
	useTCPOnly	= true
	useUDPOrTCP	= false

	// Maximum DNS packet size.
	// Value taken from https://dnsflagday.net/2020/.
	maxDNSPacketSize	= 1232
)

var (
	errLameReferral			= errors.New("lame referral")
	errCannotUnmarshalDNSMessage	= errors.New("cannot unmarshal DNS message")
	errCannotMarshalDNSMessage	= errors.New("cannot marshal DNS message")
	errServerMisbehaving		= errors.New("server misbehaving")
	errInvalidDNSResponse		= errors.New("invalid DNS response")
	errNoAnswerFromDNSServer	= errors.New("no answer from DNS server")

	// errServerTemporarilyMisbehaving is like errServerMisbehaving, except
	// that when it gets translated to a DNSError, the IsTemporary field
	// gets set to true.
	errServerTemporarilyMisbehaving	= errors.New("server misbehaving")
)

func newRequest(q dnsmessage.Question, ad bool) (id uint16, udpReq, tcpReq []byte, err error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:55
	_go_fuzz_dep_.CoverTab[5305]++
							id = uint16(randInt())
							b := dnsmessage.NewBuilder(make([]byte, 2, 514), dnsmessage.Header{ID: id, RecursionDesired: true, AuthenticData: ad})
							if err := b.StartQuestions(); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:58
		_go_fuzz_dep_.CoverTab[527900]++
//line /snap/go/10455/src/net/dnsclient_unix.go:58
		_go_fuzz_dep_.CoverTab[5312]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:59
		// _ = "end of CoverTab[5312]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:60
		_go_fuzz_dep_.CoverTab[527901]++
//line /snap/go/10455/src/net/dnsclient_unix.go:60
		_go_fuzz_dep_.CoverTab[5313]++
//line /snap/go/10455/src/net/dnsclient_unix.go:60
		// _ = "end of CoverTab[5313]"
//line /snap/go/10455/src/net/dnsclient_unix.go:60
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:60
	// _ = "end of CoverTab[5305]"
//line /snap/go/10455/src/net/dnsclient_unix.go:60
	_go_fuzz_dep_.CoverTab[5306]++
							if err := b.Question(q); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:61
		_go_fuzz_dep_.CoverTab[527902]++
//line /snap/go/10455/src/net/dnsclient_unix.go:61
		_go_fuzz_dep_.CoverTab[5314]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:62
		// _ = "end of CoverTab[5314]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:63
		_go_fuzz_dep_.CoverTab[527903]++
//line /snap/go/10455/src/net/dnsclient_unix.go:63
		_go_fuzz_dep_.CoverTab[5315]++
//line /snap/go/10455/src/net/dnsclient_unix.go:63
		// _ = "end of CoverTab[5315]"
//line /snap/go/10455/src/net/dnsclient_unix.go:63
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:63
	// _ = "end of CoverTab[5306]"
//line /snap/go/10455/src/net/dnsclient_unix.go:63
	_go_fuzz_dep_.CoverTab[5307]++

//line /snap/go/10455/src/net/dnsclient_unix.go:66
	if err := b.StartAdditionals(); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:66
		_go_fuzz_dep_.CoverTab[527904]++
//line /snap/go/10455/src/net/dnsclient_unix.go:66
		_go_fuzz_dep_.CoverTab[5316]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:67
		// _ = "end of CoverTab[5316]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:68
		_go_fuzz_dep_.CoverTab[527905]++
//line /snap/go/10455/src/net/dnsclient_unix.go:68
		_go_fuzz_dep_.CoverTab[5317]++
//line /snap/go/10455/src/net/dnsclient_unix.go:68
		// _ = "end of CoverTab[5317]"
//line /snap/go/10455/src/net/dnsclient_unix.go:68
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:68
	// _ = "end of CoverTab[5307]"
//line /snap/go/10455/src/net/dnsclient_unix.go:68
	_go_fuzz_dep_.CoverTab[5308]++
							var rh dnsmessage.ResourceHeader
							if err := rh.SetEDNS0(maxDNSPacketSize, dnsmessage.RCodeSuccess, false); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:70
		_go_fuzz_dep_.CoverTab[527906]++
//line /snap/go/10455/src/net/dnsclient_unix.go:70
		_go_fuzz_dep_.CoverTab[5318]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:71
		// _ = "end of CoverTab[5318]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:72
		_go_fuzz_dep_.CoverTab[527907]++
//line /snap/go/10455/src/net/dnsclient_unix.go:72
		_go_fuzz_dep_.CoverTab[5319]++
//line /snap/go/10455/src/net/dnsclient_unix.go:72
		// _ = "end of CoverTab[5319]"
//line /snap/go/10455/src/net/dnsclient_unix.go:72
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:72
	// _ = "end of CoverTab[5308]"
//line /snap/go/10455/src/net/dnsclient_unix.go:72
	_go_fuzz_dep_.CoverTab[5309]++
							if err := b.OPTResource(rh, dnsmessage.OPTResource{}); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:73
		_go_fuzz_dep_.CoverTab[527908]++
//line /snap/go/10455/src/net/dnsclient_unix.go:73
		_go_fuzz_dep_.CoverTab[5320]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:74
		// _ = "end of CoverTab[5320]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:75
		_go_fuzz_dep_.CoverTab[527909]++
//line /snap/go/10455/src/net/dnsclient_unix.go:75
		_go_fuzz_dep_.CoverTab[5321]++
//line /snap/go/10455/src/net/dnsclient_unix.go:75
		// _ = "end of CoverTab[5321]"
//line /snap/go/10455/src/net/dnsclient_unix.go:75
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:75
	// _ = "end of CoverTab[5309]"
//line /snap/go/10455/src/net/dnsclient_unix.go:75
	_go_fuzz_dep_.CoverTab[5310]++

							tcpReq, err = b.Finish()
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:78
		_go_fuzz_dep_.CoverTab[527910]++
//line /snap/go/10455/src/net/dnsclient_unix.go:78
		_go_fuzz_dep_.CoverTab[5322]++
								return 0, nil, nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:79
		// _ = "end of CoverTab[5322]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:80
		_go_fuzz_dep_.CoverTab[527911]++
//line /snap/go/10455/src/net/dnsclient_unix.go:80
		_go_fuzz_dep_.CoverTab[5323]++
//line /snap/go/10455/src/net/dnsclient_unix.go:80
		// _ = "end of CoverTab[5323]"
//line /snap/go/10455/src/net/dnsclient_unix.go:80
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:80
	// _ = "end of CoverTab[5310]"
//line /snap/go/10455/src/net/dnsclient_unix.go:80
	_go_fuzz_dep_.CoverTab[5311]++
							udpReq = tcpReq[2:]
							l := len(tcpReq) - 2
							tcpReq[0] = byte(l >> 8)
							tcpReq[1] = byte(l)
							return id, udpReq, tcpReq, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:85
	// _ = "end of CoverTab[5311]"
}

func checkResponse(reqID uint16, reqQues dnsmessage.Question, respHdr dnsmessage.Header, respQues dnsmessage.Question) bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:88
	_go_fuzz_dep_.CoverTab[5324]++
							if !respHdr.Response {
//line /snap/go/10455/src/net/dnsclient_unix.go:89
		_go_fuzz_dep_.CoverTab[527912]++
//line /snap/go/10455/src/net/dnsclient_unix.go:89
		_go_fuzz_dep_.CoverTab[5328]++
								return false
//line /snap/go/10455/src/net/dnsclient_unix.go:90
		// _ = "end of CoverTab[5328]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:91
		_go_fuzz_dep_.CoverTab[527913]++
//line /snap/go/10455/src/net/dnsclient_unix.go:91
		_go_fuzz_dep_.CoverTab[5329]++
//line /snap/go/10455/src/net/dnsclient_unix.go:91
		// _ = "end of CoverTab[5329]"
//line /snap/go/10455/src/net/dnsclient_unix.go:91
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:91
	// _ = "end of CoverTab[5324]"
//line /snap/go/10455/src/net/dnsclient_unix.go:91
	_go_fuzz_dep_.CoverTab[5325]++
							if reqID != respHdr.ID {
//line /snap/go/10455/src/net/dnsclient_unix.go:92
		_go_fuzz_dep_.CoverTab[527914]++
//line /snap/go/10455/src/net/dnsclient_unix.go:92
		_go_fuzz_dep_.CoverTab[5330]++
								return false
//line /snap/go/10455/src/net/dnsclient_unix.go:93
		// _ = "end of CoverTab[5330]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:94
		_go_fuzz_dep_.CoverTab[527915]++
//line /snap/go/10455/src/net/dnsclient_unix.go:94
		_go_fuzz_dep_.CoverTab[5331]++
//line /snap/go/10455/src/net/dnsclient_unix.go:94
		// _ = "end of CoverTab[5331]"
//line /snap/go/10455/src/net/dnsclient_unix.go:94
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:94
	// _ = "end of CoverTab[5325]"
//line /snap/go/10455/src/net/dnsclient_unix.go:94
	_go_fuzz_dep_.CoverTab[5326]++
							if reqQues.Type != respQues.Type || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[5332]++
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		return reqQues.Class != respQues.Class
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[5332]"
//line /snap/go/10455/src/net/dnsclient_unix.go:95
	}() || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[5333]++
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		return !equalASCIIName(reqQues.Name, respQues.Name)
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[5333]"
//line /snap/go/10455/src/net/dnsclient_unix.go:95
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[527916]++
//line /snap/go/10455/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[5334]++
								return false
//line /snap/go/10455/src/net/dnsclient_unix.go:96
		// _ = "end of CoverTab[5334]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:97
		_go_fuzz_dep_.CoverTab[527917]++
//line /snap/go/10455/src/net/dnsclient_unix.go:97
		_go_fuzz_dep_.CoverTab[5335]++
//line /snap/go/10455/src/net/dnsclient_unix.go:97
		// _ = "end of CoverTab[5335]"
//line /snap/go/10455/src/net/dnsclient_unix.go:97
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:97
	// _ = "end of CoverTab[5326]"
//line /snap/go/10455/src/net/dnsclient_unix.go:97
	_go_fuzz_dep_.CoverTab[5327]++
							return true
//line /snap/go/10455/src/net/dnsclient_unix.go:98
	// _ = "end of CoverTab[5327]"
}

func dnsPacketRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:101
	_go_fuzz_dep_.CoverTab[5336]++
							if _, err := c.Write(b); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:102
		_go_fuzz_dep_.CoverTab[527918]++
//line /snap/go/10455/src/net/dnsclient_unix.go:102
		_go_fuzz_dep_.CoverTab[5338]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:103
		// _ = "end of CoverTab[5338]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:104
		_go_fuzz_dep_.CoverTab[527919]++
//line /snap/go/10455/src/net/dnsclient_unix.go:104
		_go_fuzz_dep_.CoverTab[5339]++
//line /snap/go/10455/src/net/dnsclient_unix.go:104
		// _ = "end of CoverTab[5339]"
//line /snap/go/10455/src/net/dnsclient_unix.go:104
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:104
	// _ = "end of CoverTab[5336]"
//line /snap/go/10455/src/net/dnsclient_unix.go:104
	_go_fuzz_dep_.CoverTab[5337]++

							b = make([]byte, maxDNSPacketSize)
//line /snap/go/10455/src/net/dnsclient_unix.go:106
	_go_fuzz_dep_.CoverTab[786660] = 0
							for {
//line /snap/go/10455/src/net/dnsclient_unix.go:107
		if _go_fuzz_dep_.CoverTab[786660] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:107
			_go_fuzz_dep_.CoverTab[528133]++
//line /snap/go/10455/src/net/dnsclient_unix.go:107
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:107
			_go_fuzz_dep_.CoverTab[528134]++
//line /snap/go/10455/src/net/dnsclient_unix.go:107
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:107
		_go_fuzz_dep_.CoverTab[786660] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:107
		_go_fuzz_dep_.CoverTab[5340]++
								n, err := c.Read(b)
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:109
			_go_fuzz_dep_.CoverTab[527920]++
//line /snap/go/10455/src/net/dnsclient_unix.go:109
			_go_fuzz_dep_.CoverTab[5344]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:110
			// _ = "end of CoverTab[5344]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:111
			_go_fuzz_dep_.CoverTab[527921]++
//line /snap/go/10455/src/net/dnsclient_unix.go:111
			_go_fuzz_dep_.CoverTab[5345]++
//line /snap/go/10455/src/net/dnsclient_unix.go:111
			// _ = "end of CoverTab[5345]"
//line /snap/go/10455/src/net/dnsclient_unix.go:111
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:111
		// _ = "end of CoverTab[5340]"
//line /snap/go/10455/src/net/dnsclient_unix.go:111
		_go_fuzz_dep_.CoverTab[5341]++
								var p dnsmessage.Parser

//line /snap/go/10455/src/net/dnsclient_unix.go:116
		h, err := p.Start(b[:n])
		if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:117
			_go_fuzz_dep_.CoverTab[527922]++
//line /snap/go/10455/src/net/dnsclient_unix.go:117
			_go_fuzz_dep_.CoverTab[5346]++
									continue
//line /snap/go/10455/src/net/dnsclient_unix.go:118
			// _ = "end of CoverTab[5346]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:119
			_go_fuzz_dep_.CoverTab[527923]++
//line /snap/go/10455/src/net/dnsclient_unix.go:119
			_go_fuzz_dep_.CoverTab[5347]++
//line /snap/go/10455/src/net/dnsclient_unix.go:119
			// _ = "end of CoverTab[5347]"
//line /snap/go/10455/src/net/dnsclient_unix.go:119
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:119
		// _ = "end of CoverTab[5341]"
//line /snap/go/10455/src/net/dnsclient_unix.go:119
		_go_fuzz_dep_.CoverTab[5342]++
								q, err := p.Question()
								if err != nil || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[5348]++
//line /snap/go/10455/src/net/dnsclient_unix.go:121
			return !checkResponse(id, query, h, q)
//line /snap/go/10455/src/net/dnsclient_unix.go:121
			// _ = "end of CoverTab[5348]"
//line /snap/go/10455/src/net/dnsclient_unix.go:121
		}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[527924]++
//line /snap/go/10455/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[5349]++
									continue
//line /snap/go/10455/src/net/dnsclient_unix.go:122
			// _ = "end of CoverTab[5349]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:123
			_go_fuzz_dep_.CoverTab[527925]++
//line /snap/go/10455/src/net/dnsclient_unix.go:123
			_go_fuzz_dep_.CoverTab[5350]++
//line /snap/go/10455/src/net/dnsclient_unix.go:123
			// _ = "end of CoverTab[5350]"
//line /snap/go/10455/src/net/dnsclient_unix.go:123
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:123
		// _ = "end of CoverTab[5342]"
//line /snap/go/10455/src/net/dnsclient_unix.go:123
		_go_fuzz_dep_.CoverTab[5343]++
								return p, h, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:124
		// _ = "end of CoverTab[5343]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:125
	// _ = "end of CoverTab[5337]"
}

func dnsStreamRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:128
	_go_fuzz_dep_.CoverTab[5351]++
							if _, err := c.Write(b); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:129
		_go_fuzz_dep_.CoverTab[527926]++
//line /snap/go/10455/src/net/dnsclient_unix.go:129
		_go_fuzz_dep_.CoverTab[5359]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:130
		// _ = "end of CoverTab[5359]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:131
		_go_fuzz_dep_.CoverTab[527927]++
//line /snap/go/10455/src/net/dnsclient_unix.go:131
		_go_fuzz_dep_.CoverTab[5360]++
//line /snap/go/10455/src/net/dnsclient_unix.go:131
		// _ = "end of CoverTab[5360]"
//line /snap/go/10455/src/net/dnsclient_unix.go:131
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:131
	// _ = "end of CoverTab[5351]"
//line /snap/go/10455/src/net/dnsclient_unix.go:131
	_go_fuzz_dep_.CoverTab[5352]++

							b = make([]byte, 1280)
							if _, err := io.ReadFull(c, b[:2]); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:134
		_go_fuzz_dep_.CoverTab[527928]++
//line /snap/go/10455/src/net/dnsclient_unix.go:134
		_go_fuzz_dep_.CoverTab[5361]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:135
		// _ = "end of CoverTab[5361]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:136
		_go_fuzz_dep_.CoverTab[527929]++
//line /snap/go/10455/src/net/dnsclient_unix.go:136
		_go_fuzz_dep_.CoverTab[5362]++
//line /snap/go/10455/src/net/dnsclient_unix.go:136
		// _ = "end of CoverTab[5362]"
//line /snap/go/10455/src/net/dnsclient_unix.go:136
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:136
	// _ = "end of CoverTab[5352]"
//line /snap/go/10455/src/net/dnsclient_unix.go:136
	_go_fuzz_dep_.CoverTab[5353]++
							l := int(b[0])<<8 | int(b[1])
							if l > len(b) {
//line /snap/go/10455/src/net/dnsclient_unix.go:138
		_go_fuzz_dep_.CoverTab[527930]++
//line /snap/go/10455/src/net/dnsclient_unix.go:138
		_go_fuzz_dep_.CoverTab[5363]++
								b = make([]byte, l)
//line /snap/go/10455/src/net/dnsclient_unix.go:139
		// _ = "end of CoverTab[5363]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:140
		_go_fuzz_dep_.CoverTab[527931]++
//line /snap/go/10455/src/net/dnsclient_unix.go:140
		_go_fuzz_dep_.CoverTab[5364]++
//line /snap/go/10455/src/net/dnsclient_unix.go:140
		// _ = "end of CoverTab[5364]"
//line /snap/go/10455/src/net/dnsclient_unix.go:140
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:140
	// _ = "end of CoverTab[5353]"
//line /snap/go/10455/src/net/dnsclient_unix.go:140
	_go_fuzz_dep_.CoverTab[5354]++
							n, err := io.ReadFull(c, b[:l])
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:142
		_go_fuzz_dep_.CoverTab[527932]++
//line /snap/go/10455/src/net/dnsclient_unix.go:142
		_go_fuzz_dep_.CoverTab[5365]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:143
		// _ = "end of CoverTab[5365]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:144
		_go_fuzz_dep_.CoverTab[527933]++
//line /snap/go/10455/src/net/dnsclient_unix.go:144
		_go_fuzz_dep_.CoverTab[5366]++
//line /snap/go/10455/src/net/dnsclient_unix.go:144
		// _ = "end of CoverTab[5366]"
//line /snap/go/10455/src/net/dnsclient_unix.go:144
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:144
	// _ = "end of CoverTab[5354]"
//line /snap/go/10455/src/net/dnsclient_unix.go:144
	_go_fuzz_dep_.CoverTab[5355]++
							var p dnsmessage.Parser
							h, err := p.Start(b[:n])
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:147
		_go_fuzz_dep_.CoverTab[527934]++
//line /snap/go/10455/src/net/dnsclient_unix.go:147
		_go_fuzz_dep_.CoverTab[5367]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:148
		// _ = "end of CoverTab[5367]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:149
		_go_fuzz_dep_.CoverTab[527935]++
//line /snap/go/10455/src/net/dnsclient_unix.go:149
		_go_fuzz_dep_.CoverTab[5368]++
//line /snap/go/10455/src/net/dnsclient_unix.go:149
		// _ = "end of CoverTab[5368]"
//line /snap/go/10455/src/net/dnsclient_unix.go:149
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:149
	// _ = "end of CoverTab[5355]"
//line /snap/go/10455/src/net/dnsclient_unix.go:149
	_go_fuzz_dep_.CoverTab[5356]++
							q, err := p.Question()
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:151
		_go_fuzz_dep_.CoverTab[527936]++
//line /snap/go/10455/src/net/dnsclient_unix.go:151
		_go_fuzz_dep_.CoverTab[5369]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:152
		// _ = "end of CoverTab[5369]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:153
		_go_fuzz_dep_.CoverTab[527937]++
//line /snap/go/10455/src/net/dnsclient_unix.go:153
		_go_fuzz_dep_.CoverTab[5370]++
//line /snap/go/10455/src/net/dnsclient_unix.go:153
		// _ = "end of CoverTab[5370]"
//line /snap/go/10455/src/net/dnsclient_unix.go:153
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:153
	// _ = "end of CoverTab[5356]"
//line /snap/go/10455/src/net/dnsclient_unix.go:153
	_go_fuzz_dep_.CoverTab[5357]++
							if !checkResponse(id, query, h, q) {
//line /snap/go/10455/src/net/dnsclient_unix.go:154
		_go_fuzz_dep_.CoverTab[527938]++
//line /snap/go/10455/src/net/dnsclient_unix.go:154
		_go_fuzz_dep_.CoverTab[5371]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /snap/go/10455/src/net/dnsclient_unix.go:155
		// _ = "end of CoverTab[5371]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:156
		_go_fuzz_dep_.CoverTab[527939]++
//line /snap/go/10455/src/net/dnsclient_unix.go:156
		_go_fuzz_dep_.CoverTab[5372]++
//line /snap/go/10455/src/net/dnsclient_unix.go:156
		// _ = "end of CoverTab[5372]"
//line /snap/go/10455/src/net/dnsclient_unix.go:156
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:156
	// _ = "end of CoverTab[5357]"
//line /snap/go/10455/src/net/dnsclient_unix.go:156
	_go_fuzz_dep_.CoverTab[5358]++
							return p, h, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:157
	// _ = "end of CoverTab[5358]"
}

// exchange sends a query on the connection and hopes for a response.
func (r *Resolver) exchange(ctx context.Context, server string, q dnsmessage.Question, timeout time.Duration, useTCP, ad bool) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:161
	_go_fuzz_dep_.CoverTab[5373]++
							q.Class = dnsmessage.ClassINET
							id, udpReq, tcpReq, err := newRequest(q, ad)
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:164
		_go_fuzz_dep_.CoverTab[527940]++
//line /snap/go/10455/src/net/dnsclient_unix.go:164
		_go_fuzz_dep_.CoverTab[5377]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotMarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:165
		// _ = "end of CoverTab[5377]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:166
		_go_fuzz_dep_.CoverTab[527941]++
//line /snap/go/10455/src/net/dnsclient_unix.go:166
		_go_fuzz_dep_.CoverTab[5378]++
//line /snap/go/10455/src/net/dnsclient_unix.go:166
		// _ = "end of CoverTab[5378]"
//line /snap/go/10455/src/net/dnsclient_unix.go:166
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:166
	// _ = "end of CoverTab[5373]"
//line /snap/go/10455/src/net/dnsclient_unix.go:166
	_go_fuzz_dep_.CoverTab[5374]++
							var networks []string
							if useTCP {
//line /snap/go/10455/src/net/dnsclient_unix.go:168
		_go_fuzz_dep_.CoverTab[527942]++
//line /snap/go/10455/src/net/dnsclient_unix.go:168
		_go_fuzz_dep_.CoverTab[5379]++
								networks = []string{"tcp"}
//line /snap/go/10455/src/net/dnsclient_unix.go:169
		// _ = "end of CoverTab[5379]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:170
		_go_fuzz_dep_.CoverTab[527943]++
//line /snap/go/10455/src/net/dnsclient_unix.go:170
		_go_fuzz_dep_.CoverTab[5380]++
								networks = []string{"udp", "tcp"}
//line /snap/go/10455/src/net/dnsclient_unix.go:171
		// _ = "end of CoverTab[5380]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:172
	// _ = "end of CoverTab[5374]"
//line /snap/go/10455/src/net/dnsclient_unix.go:172
	_go_fuzz_dep_.CoverTab[5375]++
//line /snap/go/10455/src/net/dnsclient_unix.go:172
	_go_fuzz_dep_.CoverTab[786661] = 0
							for _, network := range networks {
//line /snap/go/10455/src/net/dnsclient_unix.go:173
		if _go_fuzz_dep_.CoverTab[786661] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:173
			_go_fuzz_dep_.CoverTab[528137]++
//line /snap/go/10455/src/net/dnsclient_unix.go:173
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:173
			_go_fuzz_dep_.CoverTab[528138]++
//line /snap/go/10455/src/net/dnsclient_unix.go:173
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:173
		_go_fuzz_dep_.CoverTab[786661] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:173
		_go_fuzz_dep_.CoverTab[5381]++
								ctx, cancel := context.WithDeadline(ctx, time.Now().Add(timeout))
								defer cancel()

								c, err := r.dial(ctx, network, server)
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:178
			_go_fuzz_dep_.CoverTab[527944]++
//line /snap/go/10455/src/net/dnsclient_unix.go:178
			_go_fuzz_dep_.CoverTab[5388]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:179
			// _ = "end of CoverTab[5388]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:180
			_go_fuzz_dep_.CoverTab[527945]++
//line /snap/go/10455/src/net/dnsclient_unix.go:180
			_go_fuzz_dep_.CoverTab[5389]++
//line /snap/go/10455/src/net/dnsclient_unix.go:180
			// _ = "end of CoverTab[5389]"
//line /snap/go/10455/src/net/dnsclient_unix.go:180
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:180
		// _ = "end of CoverTab[5381]"
//line /snap/go/10455/src/net/dnsclient_unix.go:180
		_go_fuzz_dep_.CoverTab[5382]++
								if d, ok := ctx.Deadline(); ok && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[5390]++
//line /snap/go/10455/src/net/dnsclient_unix.go:181
			return !d.IsZero()
//line /snap/go/10455/src/net/dnsclient_unix.go:181
			// _ = "end of CoverTab[5390]"
//line /snap/go/10455/src/net/dnsclient_unix.go:181
		}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[527946]++
//line /snap/go/10455/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[5391]++
									c.SetDeadline(d)
//line /snap/go/10455/src/net/dnsclient_unix.go:182
			// _ = "end of CoverTab[5391]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:183
			_go_fuzz_dep_.CoverTab[527947]++
//line /snap/go/10455/src/net/dnsclient_unix.go:183
			_go_fuzz_dep_.CoverTab[5392]++
//line /snap/go/10455/src/net/dnsclient_unix.go:183
			// _ = "end of CoverTab[5392]"
//line /snap/go/10455/src/net/dnsclient_unix.go:183
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:183
		// _ = "end of CoverTab[5382]"
//line /snap/go/10455/src/net/dnsclient_unix.go:183
		_go_fuzz_dep_.CoverTab[5383]++
								var p dnsmessage.Parser
								var h dnsmessage.Header
								if _, ok := c.(PacketConn); ok {
//line /snap/go/10455/src/net/dnsclient_unix.go:186
			_go_fuzz_dep_.CoverTab[527948]++
//line /snap/go/10455/src/net/dnsclient_unix.go:186
			_go_fuzz_dep_.CoverTab[5393]++
									p, h, err = dnsPacketRoundTrip(c, id, q, udpReq)
//line /snap/go/10455/src/net/dnsclient_unix.go:187
			// _ = "end of CoverTab[5393]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:188
			_go_fuzz_dep_.CoverTab[527949]++
//line /snap/go/10455/src/net/dnsclient_unix.go:188
			_go_fuzz_dep_.CoverTab[5394]++
									p, h, err = dnsStreamRoundTrip(c, id, q, tcpReq)
//line /snap/go/10455/src/net/dnsclient_unix.go:189
			// _ = "end of CoverTab[5394]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:190
		// _ = "end of CoverTab[5383]"
//line /snap/go/10455/src/net/dnsclient_unix.go:190
		_go_fuzz_dep_.CoverTab[5384]++
								c.Close()
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:192
			_go_fuzz_dep_.CoverTab[527950]++
//line /snap/go/10455/src/net/dnsclient_unix.go:192
			_go_fuzz_dep_.CoverTab[5395]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, mapErr(err)
//line /snap/go/10455/src/net/dnsclient_unix.go:193
			// _ = "end of CoverTab[5395]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:194
			_go_fuzz_dep_.CoverTab[527951]++
//line /snap/go/10455/src/net/dnsclient_unix.go:194
			_go_fuzz_dep_.CoverTab[5396]++
//line /snap/go/10455/src/net/dnsclient_unix.go:194
			// _ = "end of CoverTab[5396]"
//line /snap/go/10455/src/net/dnsclient_unix.go:194
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:194
		// _ = "end of CoverTab[5384]"
//line /snap/go/10455/src/net/dnsclient_unix.go:194
		_go_fuzz_dep_.CoverTab[5385]++
								if err := p.SkipQuestion(); err != dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/dnsclient_unix.go:195
			_go_fuzz_dep_.CoverTab[527952]++
//line /snap/go/10455/src/net/dnsclient_unix.go:195
			_go_fuzz_dep_.CoverTab[5397]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /snap/go/10455/src/net/dnsclient_unix.go:196
			// _ = "end of CoverTab[5397]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:197
			_go_fuzz_dep_.CoverTab[527953]++
//line /snap/go/10455/src/net/dnsclient_unix.go:197
			_go_fuzz_dep_.CoverTab[5398]++
//line /snap/go/10455/src/net/dnsclient_unix.go:197
			// _ = "end of CoverTab[5398]"
//line /snap/go/10455/src/net/dnsclient_unix.go:197
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:197
		// _ = "end of CoverTab[5385]"
//line /snap/go/10455/src/net/dnsclient_unix.go:197
		_go_fuzz_dep_.CoverTab[5386]++
								if h.Truncated {
//line /snap/go/10455/src/net/dnsclient_unix.go:198
			_go_fuzz_dep_.CoverTab[527954]++
//line /snap/go/10455/src/net/dnsclient_unix.go:198
			_go_fuzz_dep_.CoverTab[5399]++
									continue
//line /snap/go/10455/src/net/dnsclient_unix.go:199
			// _ = "end of CoverTab[5399]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:200
			_go_fuzz_dep_.CoverTab[527955]++
//line /snap/go/10455/src/net/dnsclient_unix.go:200
			_go_fuzz_dep_.CoverTab[5400]++
//line /snap/go/10455/src/net/dnsclient_unix.go:200
			// _ = "end of CoverTab[5400]"
//line /snap/go/10455/src/net/dnsclient_unix.go:200
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:200
		// _ = "end of CoverTab[5386]"
//line /snap/go/10455/src/net/dnsclient_unix.go:200
		_go_fuzz_dep_.CoverTab[5387]++
								return p, h, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:201
		// _ = "end of CoverTab[5387]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:202
	if _go_fuzz_dep_.CoverTab[786661] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:202
		_go_fuzz_dep_.CoverTab[528139]++
//line /snap/go/10455/src/net/dnsclient_unix.go:202
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:202
		_go_fuzz_dep_.CoverTab[528140]++
//line /snap/go/10455/src/net/dnsclient_unix.go:202
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:202
	// _ = "end of CoverTab[5375]"
//line /snap/go/10455/src/net/dnsclient_unix.go:202
	_go_fuzz_dep_.CoverTab[5376]++
							return dnsmessage.Parser{}, dnsmessage.Header{}, errNoAnswerFromDNSServer
//line /snap/go/10455/src/net/dnsclient_unix.go:203
	// _ = "end of CoverTab[5376]"
}

// checkHeader performs basic sanity checks on the header.
func checkHeader(p *dnsmessage.Parser, h dnsmessage.Header) error {
//line /snap/go/10455/src/net/dnsclient_unix.go:207
	_go_fuzz_dep_.CoverTab[5401]++
							if h.RCode == dnsmessage.RCodeNameError {
//line /snap/go/10455/src/net/dnsclient_unix.go:208
		_go_fuzz_dep_.CoverTab[527956]++
//line /snap/go/10455/src/net/dnsclient_unix.go:208
		_go_fuzz_dep_.CoverTab[5406]++
								return errNoSuchHost
//line /snap/go/10455/src/net/dnsclient_unix.go:209
		// _ = "end of CoverTab[5406]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:210
		_go_fuzz_dep_.CoverTab[527957]++
//line /snap/go/10455/src/net/dnsclient_unix.go:210
		_go_fuzz_dep_.CoverTab[5407]++
//line /snap/go/10455/src/net/dnsclient_unix.go:210
		// _ = "end of CoverTab[5407]"
//line /snap/go/10455/src/net/dnsclient_unix.go:210
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:210
	// _ = "end of CoverTab[5401]"
//line /snap/go/10455/src/net/dnsclient_unix.go:210
	_go_fuzz_dep_.CoverTab[5402]++

							_, err := p.AnswerHeader()
							if err != nil && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[5408]++
//line /snap/go/10455/src/net/dnsclient_unix.go:213
		return err != dnsmessage.ErrSectionDone
//line /snap/go/10455/src/net/dnsclient_unix.go:213
		// _ = "end of CoverTab[5408]"
//line /snap/go/10455/src/net/dnsclient_unix.go:213
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[527958]++
//line /snap/go/10455/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[5409]++
								return errCannotUnmarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:214
		// _ = "end of CoverTab[5409]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:215
		_go_fuzz_dep_.CoverTab[527959]++
//line /snap/go/10455/src/net/dnsclient_unix.go:215
		_go_fuzz_dep_.CoverTab[5410]++
//line /snap/go/10455/src/net/dnsclient_unix.go:215
		// _ = "end of CoverTab[5410]"
//line /snap/go/10455/src/net/dnsclient_unix.go:215
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:215
	// _ = "end of CoverTab[5402]"
//line /snap/go/10455/src/net/dnsclient_unix.go:215
	_go_fuzz_dep_.CoverTab[5403]++

//line /snap/go/10455/src/net/dnsclient_unix.go:219
	if h.RCode == dnsmessage.RCodeSuccess && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5411]++
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		return !h.Authoritative
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5411]"
//line /snap/go/10455/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5412]++
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		return !h.RecursionAvailable
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5412]"
//line /snap/go/10455/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5413]++
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		return err == dnsmessage.ErrSectionDone
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5413]"
//line /snap/go/10455/src/net/dnsclient_unix.go:219
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[527960]++
//line /snap/go/10455/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5414]++
								return errLameReferral
//line /snap/go/10455/src/net/dnsclient_unix.go:220
		// _ = "end of CoverTab[5414]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:221
		_go_fuzz_dep_.CoverTab[527961]++
//line /snap/go/10455/src/net/dnsclient_unix.go:221
		_go_fuzz_dep_.CoverTab[5415]++
//line /snap/go/10455/src/net/dnsclient_unix.go:221
		// _ = "end of CoverTab[5415]"
//line /snap/go/10455/src/net/dnsclient_unix.go:221
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:221
	// _ = "end of CoverTab[5403]"
//line /snap/go/10455/src/net/dnsclient_unix.go:221
	_go_fuzz_dep_.CoverTab[5404]++

							if h.RCode != dnsmessage.RCodeSuccess && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[5416]++
//line /snap/go/10455/src/net/dnsclient_unix.go:223
		return h.RCode != dnsmessage.RCodeNameError
//line /snap/go/10455/src/net/dnsclient_unix.go:223
		// _ = "end of CoverTab[5416]"
//line /snap/go/10455/src/net/dnsclient_unix.go:223
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[527962]++
//line /snap/go/10455/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[5417]++

//line /snap/go/10455/src/net/dnsclient_unix.go:229
		if h.RCode == dnsmessage.RCodeServerFailure {
//line /snap/go/10455/src/net/dnsclient_unix.go:229
			_go_fuzz_dep_.CoverTab[527964]++
//line /snap/go/10455/src/net/dnsclient_unix.go:229
			_go_fuzz_dep_.CoverTab[5419]++
									return errServerTemporarilyMisbehaving
//line /snap/go/10455/src/net/dnsclient_unix.go:230
			// _ = "end of CoverTab[5419]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:231
			_go_fuzz_dep_.CoverTab[527965]++
//line /snap/go/10455/src/net/dnsclient_unix.go:231
			_go_fuzz_dep_.CoverTab[5420]++
//line /snap/go/10455/src/net/dnsclient_unix.go:231
			// _ = "end of CoverTab[5420]"
//line /snap/go/10455/src/net/dnsclient_unix.go:231
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:231
		// _ = "end of CoverTab[5417]"
//line /snap/go/10455/src/net/dnsclient_unix.go:231
		_go_fuzz_dep_.CoverTab[5418]++
								return errServerMisbehaving
//line /snap/go/10455/src/net/dnsclient_unix.go:232
		// _ = "end of CoverTab[5418]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:233
		_go_fuzz_dep_.CoverTab[527963]++
//line /snap/go/10455/src/net/dnsclient_unix.go:233
		_go_fuzz_dep_.CoverTab[5421]++
//line /snap/go/10455/src/net/dnsclient_unix.go:233
		// _ = "end of CoverTab[5421]"
//line /snap/go/10455/src/net/dnsclient_unix.go:233
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:233
	// _ = "end of CoverTab[5404]"
//line /snap/go/10455/src/net/dnsclient_unix.go:233
	_go_fuzz_dep_.CoverTab[5405]++

							return nil
//line /snap/go/10455/src/net/dnsclient_unix.go:235
	// _ = "end of CoverTab[5405]"
}

func skipToAnswer(p *dnsmessage.Parser, qtype dnsmessage.Type) error {
//line /snap/go/10455/src/net/dnsclient_unix.go:238
	_go_fuzz_dep_.CoverTab[5422]++
//line /snap/go/10455/src/net/dnsclient_unix.go:238
	_go_fuzz_dep_.CoverTab[786662] = 0
							for {
//line /snap/go/10455/src/net/dnsclient_unix.go:239
		if _go_fuzz_dep_.CoverTab[786662] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:239
			_go_fuzz_dep_.CoverTab[528141]++
//line /snap/go/10455/src/net/dnsclient_unix.go:239
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:239
			_go_fuzz_dep_.CoverTab[528142]++
//line /snap/go/10455/src/net/dnsclient_unix.go:239
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:239
		_go_fuzz_dep_.CoverTab[786662] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:239
		_go_fuzz_dep_.CoverTab[5423]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/dnsclient_unix.go:241
			_go_fuzz_dep_.CoverTab[527966]++
//line /snap/go/10455/src/net/dnsclient_unix.go:241
			_go_fuzz_dep_.CoverTab[5427]++
									return errNoSuchHost
//line /snap/go/10455/src/net/dnsclient_unix.go:242
			// _ = "end of CoverTab[5427]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:243
			_go_fuzz_dep_.CoverTab[527967]++
//line /snap/go/10455/src/net/dnsclient_unix.go:243
			_go_fuzz_dep_.CoverTab[5428]++
//line /snap/go/10455/src/net/dnsclient_unix.go:243
			// _ = "end of CoverTab[5428]"
//line /snap/go/10455/src/net/dnsclient_unix.go:243
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:243
		// _ = "end of CoverTab[5423]"
//line /snap/go/10455/src/net/dnsclient_unix.go:243
		_go_fuzz_dep_.CoverTab[5424]++
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:244
			_go_fuzz_dep_.CoverTab[527968]++
//line /snap/go/10455/src/net/dnsclient_unix.go:244
			_go_fuzz_dep_.CoverTab[5429]++
									return errCannotUnmarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:245
			// _ = "end of CoverTab[5429]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:246
			_go_fuzz_dep_.CoverTab[527969]++
//line /snap/go/10455/src/net/dnsclient_unix.go:246
			_go_fuzz_dep_.CoverTab[5430]++
//line /snap/go/10455/src/net/dnsclient_unix.go:246
			// _ = "end of CoverTab[5430]"
//line /snap/go/10455/src/net/dnsclient_unix.go:246
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:246
		// _ = "end of CoverTab[5424]"
//line /snap/go/10455/src/net/dnsclient_unix.go:246
		_go_fuzz_dep_.CoverTab[5425]++
								if h.Type == qtype {
//line /snap/go/10455/src/net/dnsclient_unix.go:247
			_go_fuzz_dep_.CoverTab[527970]++
//line /snap/go/10455/src/net/dnsclient_unix.go:247
			_go_fuzz_dep_.CoverTab[5431]++
									return nil
//line /snap/go/10455/src/net/dnsclient_unix.go:248
			// _ = "end of CoverTab[5431]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:249
			_go_fuzz_dep_.CoverTab[527971]++
//line /snap/go/10455/src/net/dnsclient_unix.go:249
			_go_fuzz_dep_.CoverTab[5432]++
//line /snap/go/10455/src/net/dnsclient_unix.go:249
			// _ = "end of CoverTab[5432]"
//line /snap/go/10455/src/net/dnsclient_unix.go:249
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:249
		// _ = "end of CoverTab[5425]"
//line /snap/go/10455/src/net/dnsclient_unix.go:249
		_go_fuzz_dep_.CoverTab[5426]++
								if err := p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:250
			_go_fuzz_dep_.CoverTab[527972]++
//line /snap/go/10455/src/net/dnsclient_unix.go:250
			_go_fuzz_dep_.CoverTab[5433]++
									return errCannotUnmarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:251
			// _ = "end of CoverTab[5433]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:252
			_go_fuzz_dep_.CoverTab[527973]++
//line /snap/go/10455/src/net/dnsclient_unix.go:252
			_go_fuzz_dep_.CoverTab[5434]++
//line /snap/go/10455/src/net/dnsclient_unix.go:252
			// _ = "end of CoverTab[5434]"
//line /snap/go/10455/src/net/dnsclient_unix.go:252
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:252
		// _ = "end of CoverTab[5426]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:253
	// _ = "end of CoverTab[5422]"
}

// Do a lookup for a single name, which must be rooted
//line /snap/go/10455/src/net/dnsclient_unix.go:256
// (otherwise answer will not find the answers).
//line /snap/go/10455/src/net/dnsclient_unix.go:258
func (r *Resolver) tryOneName(ctx context.Context, cfg *dnsConfig, name string, qtype dnsmessage.Type) (dnsmessage.Parser, string, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:258
	_go_fuzz_dep_.CoverTab[5435]++
							var lastErr error
							serverOffset := cfg.serverOffset()
							sLen := uint32(len(cfg.servers))

							n, err := dnsmessage.NewName(name)
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:264
		_go_fuzz_dep_.CoverTab[527974]++
//line /snap/go/10455/src/net/dnsclient_unix.go:264
		_go_fuzz_dep_.CoverTab[5438]++
								return dnsmessage.Parser{}, "", errCannotMarshalDNSMessage
//line /snap/go/10455/src/net/dnsclient_unix.go:265
		// _ = "end of CoverTab[5438]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:266
		_go_fuzz_dep_.CoverTab[527975]++
//line /snap/go/10455/src/net/dnsclient_unix.go:266
		_go_fuzz_dep_.CoverTab[5439]++
//line /snap/go/10455/src/net/dnsclient_unix.go:266
		// _ = "end of CoverTab[5439]"
//line /snap/go/10455/src/net/dnsclient_unix.go:266
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:266
	// _ = "end of CoverTab[5435]"
//line /snap/go/10455/src/net/dnsclient_unix.go:266
	_go_fuzz_dep_.CoverTab[5436]++
							q := dnsmessage.Question{
		Name:	n,
		Type:	qtype,
		Class:	dnsmessage.ClassINET,
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:271
	_go_fuzz_dep_.CoverTab[786663] = 0

							for i := 0; i < cfg.attempts; i++ {
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		if _go_fuzz_dep_.CoverTab[786663] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:273
			_go_fuzz_dep_.CoverTab[528145]++
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:273
			_go_fuzz_dep_.CoverTab[528146]++
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		_go_fuzz_dep_.CoverTab[786663] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		_go_fuzz_dep_.CoverTab[5440]++
//line /snap/go/10455/src/net/dnsclient_unix.go:273
		_go_fuzz_dep_.CoverTab[786664] = 0
								for j := uint32(0); j < sLen; j++ {
//line /snap/go/10455/src/net/dnsclient_unix.go:274
			if _go_fuzz_dep_.CoverTab[786664] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:274
				_go_fuzz_dep_.CoverTab[528149]++
//line /snap/go/10455/src/net/dnsclient_unix.go:274
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:274
				_go_fuzz_dep_.CoverTab[528150]++
//line /snap/go/10455/src/net/dnsclient_unix.go:274
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:274
			_go_fuzz_dep_.CoverTab[786664] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:274
			_go_fuzz_dep_.CoverTab[5441]++
									server := cfg.servers[(serverOffset+j)%sLen]

									p, h, err := r.exchange(ctx, server, q, cfg.timeout, cfg.useTCP, cfg.trustAD)
									if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:278
				_go_fuzz_dep_.CoverTab[527976]++
//line /snap/go/10455/src/net/dnsclient_unix.go:278
				_go_fuzz_dep_.CoverTab[5445]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if nerr, ok := err.(Error); ok && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[5448]++
//line /snap/go/10455/src/net/dnsclient_unix.go:284
					return nerr.Timeout()
//line /snap/go/10455/src/net/dnsclient_unix.go:284
					// _ = "end of CoverTab[5448]"
//line /snap/go/10455/src/net/dnsclient_unix.go:284
				}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[527978]++
//line /snap/go/10455/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[5449]++
											dnsErr.IsTimeout = true
//line /snap/go/10455/src/net/dnsclient_unix.go:285
					// _ = "end of CoverTab[5449]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:286
					_go_fuzz_dep_.CoverTab[527979]++
//line /snap/go/10455/src/net/dnsclient_unix.go:286
					_go_fuzz_dep_.CoverTab[5450]++
//line /snap/go/10455/src/net/dnsclient_unix.go:286
					// _ = "end of CoverTab[5450]"
//line /snap/go/10455/src/net/dnsclient_unix.go:286
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:286
				// _ = "end of CoverTab[5445]"
//line /snap/go/10455/src/net/dnsclient_unix.go:286
				_go_fuzz_dep_.CoverTab[5446]++

//line /snap/go/10455/src/net/dnsclient_unix.go:289
				if _, ok := err.(*OpError); ok {
//line /snap/go/10455/src/net/dnsclient_unix.go:289
					_go_fuzz_dep_.CoverTab[527980]++
//line /snap/go/10455/src/net/dnsclient_unix.go:289
					_go_fuzz_dep_.CoverTab[5451]++
											dnsErr.IsTemporary = true
//line /snap/go/10455/src/net/dnsclient_unix.go:290
					// _ = "end of CoverTab[5451]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:291
					_go_fuzz_dep_.CoverTab[527981]++
//line /snap/go/10455/src/net/dnsclient_unix.go:291
					_go_fuzz_dep_.CoverTab[5452]++
//line /snap/go/10455/src/net/dnsclient_unix.go:291
					// _ = "end of CoverTab[5452]"
//line /snap/go/10455/src/net/dnsclient_unix.go:291
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:291
				// _ = "end of CoverTab[5446]"
//line /snap/go/10455/src/net/dnsclient_unix.go:291
				_go_fuzz_dep_.CoverTab[5447]++
										lastErr = dnsErr
										continue
//line /snap/go/10455/src/net/dnsclient_unix.go:293
				// _ = "end of CoverTab[5447]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:294
				_go_fuzz_dep_.CoverTab[527977]++
//line /snap/go/10455/src/net/dnsclient_unix.go:294
				_go_fuzz_dep_.CoverTab[5453]++
//line /snap/go/10455/src/net/dnsclient_unix.go:294
				// _ = "end of CoverTab[5453]"
//line /snap/go/10455/src/net/dnsclient_unix.go:294
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:294
			// _ = "end of CoverTab[5441]"
//line /snap/go/10455/src/net/dnsclient_unix.go:294
			_go_fuzz_dep_.CoverTab[5442]++

									if err := checkHeader(&p, h); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:296
				_go_fuzz_dep_.CoverTab[527982]++
//line /snap/go/10455/src/net/dnsclient_unix.go:296
				_go_fuzz_dep_.CoverTab[5454]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if err == errServerTemporarilyMisbehaving {
//line /snap/go/10455/src/net/dnsclient_unix.go:302
					_go_fuzz_dep_.CoverTab[527984]++
//line /snap/go/10455/src/net/dnsclient_unix.go:302
					_go_fuzz_dep_.CoverTab[5457]++
											dnsErr.IsTemporary = true
//line /snap/go/10455/src/net/dnsclient_unix.go:303
					// _ = "end of CoverTab[5457]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:304
					_go_fuzz_dep_.CoverTab[527985]++
//line /snap/go/10455/src/net/dnsclient_unix.go:304
					_go_fuzz_dep_.CoverTab[5458]++
//line /snap/go/10455/src/net/dnsclient_unix.go:304
					// _ = "end of CoverTab[5458]"
//line /snap/go/10455/src/net/dnsclient_unix.go:304
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:304
				// _ = "end of CoverTab[5454]"
//line /snap/go/10455/src/net/dnsclient_unix.go:304
				_go_fuzz_dep_.CoverTab[5455]++
										if err == errNoSuchHost {
//line /snap/go/10455/src/net/dnsclient_unix.go:305
					_go_fuzz_dep_.CoverTab[527986]++
//line /snap/go/10455/src/net/dnsclient_unix.go:305
					_go_fuzz_dep_.CoverTab[5459]++

//line /snap/go/10455/src/net/dnsclient_unix.go:309
					dnsErr.IsNotFound = true
											return p, server, dnsErr
//line /snap/go/10455/src/net/dnsclient_unix.go:310
					// _ = "end of CoverTab[5459]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:311
					_go_fuzz_dep_.CoverTab[527987]++
//line /snap/go/10455/src/net/dnsclient_unix.go:311
					_go_fuzz_dep_.CoverTab[5460]++
//line /snap/go/10455/src/net/dnsclient_unix.go:311
					// _ = "end of CoverTab[5460]"
//line /snap/go/10455/src/net/dnsclient_unix.go:311
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:311
				// _ = "end of CoverTab[5455]"
//line /snap/go/10455/src/net/dnsclient_unix.go:311
				_go_fuzz_dep_.CoverTab[5456]++
										lastErr = dnsErr
										continue
//line /snap/go/10455/src/net/dnsclient_unix.go:313
				// _ = "end of CoverTab[5456]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:314
				_go_fuzz_dep_.CoverTab[527983]++
//line /snap/go/10455/src/net/dnsclient_unix.go:314
				_go_fuzz_dep_.CoverTab[5461]++
//line /snap/go/10455/src/net/dnsclient_unix.go:314
				// _ = "end of CoverTab[5461]"
//line /snap/go/10455/src/net/dnsclient_unix.go:314
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:314
			// _ = "end of CoverTab[5442]"
//line /snap/go/10455/src/net/dnsclient_unix.go:314
			_go_fuzz_dep_.CoverTab[5443]++

									err = skipToAnswer(&p, qtype)
									if err == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:317
				_go_fuzz_dep_.CoverTab[527988]++
//line /snap/go/10455/src/net/dnsclient_unix.go:317
				_go_fuzz_dep_.CoverTab[5462]++
										return p, server, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:318
				// _ = "end of CoverTab[5462]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:319
				_go_fuzz_dep_.CoverTab[527989]++
//line /snap/go/10455/src/net/dnsclient_unix.go:319
				_go_fuzz_dep_.CoverTab[5463]++
//line /snap/go/10455/src/net/dnsclient_unix.go:319
				// _ = "end of CoverTab[5463]"
//line /snap/go/10455/src/net/dnsclient_unix.go:319
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:319
			// _ = "end of CoverTab[5443]"
//line /snap/go/10455/src/net/dnsclient_unix.go:319
			_go_fuzz_dep_.CoverTab[5444]++
									lastErr = &DNSError{
				Err:	err.Error(),
				Name:	name,
				Server:	server,
			}
			if err == errNoSuchHost {
//line /snap/go/10455/src/net/dnsclient_unix.go:325
				_go_fuzz_dep_.CoverTab[527990]++
//line /snap/go/10455/src/net/dnsclient_unix.go:325
				_go_fuzz_dep_.CoverTab[5464]++

//line /snap/go/10455/src/net/dnsclient_unix.go:329
				lastErr.(*DNSError).IsNotFound = true
										return p, server, lastErr
//line /snap/go/10455/src/net/dnsclient_unix.go:330
				// _ = "end of CoverTab[5464]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:331
				_go_fuzz_dep_.CoverTab[527991]++
//line /snap/go/10455/src/net/dnsclient_unix.go:331
				_go_fuzz_dep_.CoverTab[5465]++
//line /snap/go/10455/src/net/dnsclient_unix.go:331
				// _ = "end of CoverTab[5465]"
//line /snap/go/10455/src/net/dnsclient_unix.go:331
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:331
			// _ = "end of CoverTab[5444]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:332
		if _go_fuzz_dep_.CoverTab[786664] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:332
			_go_fuzz_dep_.CoverTab[528151]++
//line /snap/go/10455/src/net/dnsclient_unix.go:332
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:332
			_go_fuzz_dep_.CoverTab[528152]++
//line /snap/go/10455/src/net/dnsclient_unix.go:332
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:332
		// _ = "end of CoverTab[5440]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:333
	if _go_fuzz_dep_.CoverTab[786663] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:333
		_go_fuzz_dep_.CoverTab[528147]++
//line /snap/go/10455/src/net/dnsclient_unix.go:333
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:333
		_go_fuzz_dep_.CoverTab[528148]++
//line /snap/go/10455/src/net/dnsclient_unix.go:333
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:333
	// _ = "end of CoverTab[5436]"
//line /snap/go/10455/src/net/dnsclient_unix.go:333
	_go_fuzz_dep_.CoverTab[5437]++
							return dnsmessage.Parser{}, "", lastErr
//line /snap/go/10455/src/net/dnsclient_unix.go:334
	// _ = "end of CoverTab[5437]"
}

// A resolverConfig represents a DNS stub resolver configuration.
type resolverConfig struct {
	initOnce	sync.Once	// guards init of resolverConfig

	// ch is used as a semaphore that only allows one lookup at a
	// time to recheck resolv.conf.
	ch		chan struct{}	// guards lastChecked and modTime
	lastChecked	time.Time	// last time resolv.conf was checked

	dnsConfig	atomic.Pointer[dnsConfig]	// parsed resolv.conf structure used in lookups
}

var resolvConf resolverConfig

func getSystemDNSConfig() *dnsConfig {
//line /snap/go/10455/src/net/dnsclient_unix.go:351
	_go_fuzz_dep_.CoverTab[5466]++
							resolvConf.tryUpdate("/etc/resolv.conf")
							return resolvConf.dnsConfig.Load()
//line /snap/go/10455/src/net/dnsclient_unix.go:353
	// _ = "end of CoverTab[5466]"
}

// init initializes conf and is only called via conf.initOnce.
func (conf *resolverConfig) init() {

//line /snap/go/10455/src/net/dnsclient_unix.go:360
	conf.dnsConfig.Store(dnsReadConfig("/etc/resolv.conf"))
							conf.lastChecked = time.Now()

//line /snap/go/10455/src/net/dnsclient_unix.go:365
	conf.ch = make(chan struct{}, 1)
}

// tryUpdate tries to update conf with the named resolv.conf file.
//line /snap/go/10455/src/net/dnsclient_unix.go:368
// The name variable only exists for testing. It is otherwise always
//line /snap/go/10455/src/net/dnsclient_unix.go:368
// "/etc/resolv.conf".
//line /snap/go/10455/src/net/dnsclient_unix.go:371
func (conf *resolverConfig) tryUpdate(name string) {
//line /snap/go/10455/src/net/dnsclient_unix.go:371
	_go_fuzz_dep_.CoverTab[5467]++
							conf.initOnce.Do(conf.init)

							if conf.dnsConfig.Load().noReload {
//line /snap/go/10455/src/net/dnsclient_unix.go:374
		_go_fuzz_dep_.CoverTab[527992]++
//line /snap/go/10455/src/net/dnsclient_unix.go:374
		_go_fuzz_dep_.CoverTab[5472]++
								return
//line /snap/go/10455/src/net/dnsclient_unix.go:375
		// _ = "end of CoverTab[5472]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:376
		_go_fuzz_dep_.CoverTab[527993]++
//line /snap/go/10455/src/net/dnsclient_unix.go:376
		_go_fuzz_dep_.CoverTab[5473]++
//line /snap/go/10455/src/net/dnsclient_unix.go:376
		// _ = "end of CoverTab[5473]"
//line /snap/go/10455/src/net/dnsclient_unix.go:376
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:376
	// _ = "end of CoverTab[5467]"
//line /snap/go/10455/src/net/dnsclient_unix.go:376
	_go_fuzz_dep_.CoverTab[5468]++

//line /snap/go/10455/src/net/dnsclient_unix.go:379
	if !conf.tryAcquireSema() {
//line /snap/go/10455/src/net/dnsclient_unix.go:379
		_go_fuzz_dep_.CoverTab[527994]++
//line /snap/go/10455/src/net/dnsclient_unix.go:379
		_go_fuzz_dep_.CoverTab[5474]++
								return
//line /snap/go/10455/src/net/dnsclient_unix.go:380
		// _ = "end of CoverTab[5474]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:381
		_go_fuzz_dep_.CoverTab[527995]++
//line /snap/go/10455/src/net/dnsclient_unix.go:381
		_go_fuzz_dep_.CoverTab[5475]++
//line /snap/go/10455/src/net/dnsclient_unix.go:381
		// _ = "end of CoverTab[5475]"
//line /snap/go/10455/src/net/dnsclient_unix.go:381
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:381
	// _ = "end of CoverTab[5468]"
//line /snap/go/10455/src/net/dnsclient_unix.go:381
	_go_fuzz_dep_.CoverTab[5469]++
							defer conf.releaseSema()

							now := time.Now()
							if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /snap/go/10455/src/net/dnsclient_unix.go:385
		_go_fuzz_dep_.CoverTab[527996]++
//line /snap/go/10455/src/net/dnsclient_unix.go:385
		_go_fuzz_dep_.CoverTab[5476]++
								return
//line /snap/go/10455/src/net/dnsclient_unix.go:386
		// _ = "end of CoverTab[5476]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:387
		_go_fuzz_dep_.CoverTab[527997]++
//line /snap/go/10455/src/net/dnsclient_unix.go:387
		_go_fuzz_dep_.CoverTab[5477]++
//line /snap/go/10455/src/net/dnsclient_unix.go:387
		// _ = "end of CoverTab[5477]"
//line /snap/go/10455/src/net/dnsclient_unix.go:387
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:387
	// _ = "end of CoverTab[5469]"
//line /snap/go/10455/src/net/dnsclient_unix.go:387
	_go_fuzz_dep_.CoverTab[5470]++
							conf.lastChecked = now

							switch runtime.GOOS {
	case "windows":
//line /snap/go/10455/src/net/dnsclient_unix.go:391
		_go_fuzz_dep_.CoverTab[527998]++
//line /snap/go/10455/src/net/dnsclient_unix.go:391
		_go_fuzz_dep_.CoverTab[5478]++
//line /snap/go/10455/src/net/dnsclient_unix.go:391
		// _ = "end of CoverTab[5478]"

//line /snap/go/10455/src/net/dnsclient_unix.go:397
	default:
//line /snap/go/10455/src/net/dnsclient_unix.go:397
		_go_fuzz_dep_.CoverTab[527999]++
//line /snap/go/10455/src/net/dnsclient_unix.go:397
		_go_fuzz_dep_.CoverTab[5479]++
								var mtime time.Time
								if fi, err := os.Stat(name); err == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:399
			_go_fuzz_dep_.CoverTab[528000]++
//line /snap/go/10455/src/net/dnsclient_unix.go:399
			_go_fuzz_dep_.CoverTab[5481]++
									mtime = fi.ModTime()
//line /snap/go/10455/src/net/dnsclient_unix.go:400
			// _ = "end of CoverTab[5481]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:401
			_go_fuzz_dep_.CoverTab[528001]++
//line /snap/go/10455/src/net/dnsclient_unix.go:401
			_go_fuzz_dep_.CoverTab[5482]++
//line /snap/go/10455/src/net/dnsclient_unix.go:401
			// _ = "end of CoverTab[5482]"
//line /snap/go/10455/src/net/dnsclient_unix.go:401
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:401
		// _ = "end of CoverTab[5479]"
//line /snap/go/10455/src/net/dnsclient_unix.go:401
		_go_fuzz_dep_.CoverTab[5480]++
								if mtime.Equal(conf.dnsConfig.Load().mtime) {
//line /snap/go/10455/src/net/dnsclient_unix.go:402
			_go_fuzz_dep_.CoverTab[528002]++
//line /snap/go/10455/src/net/dnsclient_unix.go:402
			_go_fuzz_dep_.CoverTab[5483]++
									return
//line /snap/go/10455/src/net/dnsclient_unix.go:403
			// _ = "end of CoverTab[5483]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:404
			_go_fuzz_dep_.CoverTab[528003]++
//line /snap/go/10455/src/net/dnsclient_unix.go:404
			_go_fuzz_dep_.CoverTab[5484]++
//line /snap/go/10455/src/net/dnsclient_unix.go:404
			// _ = "end of CoverTab[5484]"
//line /snap/go/10455/src/net/dnsclient_unix.go:404
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:404
		// _ = "end of CoverTab[5480]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:405
	// _ = "end of CoverTab[5470]"
//line /snap/go/10455/src/net/dnsclient_unix.go:405
	_go_fuzz_dep_.CoverTab[5471]++

							dnsConf := dnsReadConfig(name)
							conf.dnsConfig.Store(dnsConf)
//line /snap/go/10455/src/net/dnsclient_unix.go:408
	// _ = "end of CoverTab[5471]"
}

func (conf *resolverConfig) tryAcquireSema() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:411
	_go_fuzz_dep_.CoverTab[5485]++
							select {
	case conf.ch <- struct{}{}:
//line /snap/go/10455/src/net/dnsclient_unix.go:413
		_go_fuzz_dep_.CoverTab[5486]++
								return true
//line /snap/go/10455/src/net/dnsclient_unix.go:414
		// _ = "end of CoverTab[5486]"
	default:
//line /snap/go/10455/src/net/dnsclient_unix.go:415
		_go_fuzz_dep_.CoverTab[5487]++
								return false
//line /snap/go/10455/src/net/dnsclient_unix.go:416
		// _ = "end of CoverTab[5487]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:417
	// _ = "end of CoverTab[5485]"
}

func (conf *resolverConfig) releaseSema() {
//line /snap/go/10455/src/net/dnsclient_unix.go:420
	_go_fuzz_dep_.CoverTab[5488]++
							<-conf.ch
//line /snap/go/10455/src/net/dnsclient_unix.go:421
	// _ = "end of CoverTab[5488]"
}

func (r *Resolver) lookup(ctx context.Context, name string, qtype dnsmessage.Type, conf *dnsConfig) (dnsmessage.Parser, string, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:424
	_go_fuzz_dep_.CoverTab[5489]++
							if !isDomainName(name) {
//line /snap/go/10455/src/net/dnsclient_unix.go:425
		_go_fuzz_dep_.CoverTab[528004]++
//line /snap/go/10455/src/net/dnsclient_unix.go:425
		_go_fuzz_dep_.CoverTab[5495]++

//line /snap/go/10455/src/net/dnsclient_unix.go:431
		return dnsmessage.Parser{}, "", &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /snap/go/10455/src/net/dnsclient_unix.go:431
		// _ = "end of CoverTab[5495]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:432
		_go_fuzz_dep_.CoverTab[528005]++
//line /snap/go/10455/src/net/dnsclient_unix.go:432
		_go_fuzz_dep_.CoverTab[5496]++
//line /snap/go/10455/src/net/dnsclient_unix.go:432
		// _ = "end of CoverTab[5496]"
//line /snap/go/10455/src/net/dnsclient_unix.go:432
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:432
	// _ = "end of CoverTab[5489]"
//line /snap/go/10455/src/net/dnsclient_unix.go:432
	_go_fuzz_dep_.CoverTab[5490]++

							if conf == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:434
		_go_fuzz_dep_.CoverTab[528006]++
//line /snap/go/10455/src/net/dnsclient_unix.go:434
		_go_fuzz_dep_.CoverTab[5497]++
								conf = getSystemDNSConfig()
//line /snap/go/10455/src/net/dnsclient_unix.go:435
		// _ = "end of CoverTab[5497]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:436
		_go_fuzz_dep_.CoverTab[528007]++
//line /snap/go/10455/src/net/dnsclient_unix.go:436
		_go_fuzz_dep_.CoverTab[5498]++
//line /snap/go/10455/src/net/dnsclient_unix.go:436
		// _ = "end of CoverTab[5498]"
//line /snap/go/10455/src/net/dnsclient_unix.go:436
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:436
	// _ = "end of CoverTab[5490]"
//line /snap/go/10455/src/net/dnsclient_unix.go:436
	_go_fuzz_dep_.CoverTab[5491]++

							var (
		p	dnsmessage.Parser
		server	string
		err	error
	)
//line /snap/go/10455/src/net/dnsclient_unix.go:442
	_go_fuzz_dep_.CoverTab[786665] = 0
							for _, fqdn := range conf.nameList(name) {
//line /snap/go/10455/src/net/dnsclient_unix.go:443
		if _go_fuzz_dep_.CoverTab[786665] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:443
			_go_fuzz_dep_.CoverTab[528153]++
//line /snap/go/10455/src/net/dnsclient_unix.go:443
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:443
			_go_fuzz_dep_.CoverTab[528154]++
//line /snap/go/10455/src/net/dnsclient_unix.go:443
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:443
		_go_fuzz_dep_.CoverTab[786665] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:443
		_go_fuzz_dep_.CoverTab[5499]++
								p, server, err = r.tryOneName(ctx, conf, fqdn, qtype)
								if err == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:445
			_go_fuzz_dep_.CoverTab[528008]++
//line /snap/go/10455/src/net/dnsclient_unix.go:445
			_go_fuzz_dep_.CoverTab[5501]++
									break
//line /snap/go/10455/src/net/dnsclient_unix.go:446
			// _ = "end of CoverTab[5501]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:447
			_go_fuzz_dep_.CoverTab[528009]++
//line /snap/go/10455/src/net/dnsclient_unix.go:447
			_go_fuzz_dep_.CoverTab[5502]++
//line /snap/go/10455/src/net/dnsclient_unix.go:447
			// _ = "end of CoverTab[5502]"
//line /snap/go/10455/src/net/dnsclient_unix.go:447
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:447
		// _ = "end of CoverTab[5499]"
//line /snap/go/10455/src/net/dnsclient_unix.go:447
		_go_fuzz_dep_.CoverTab[5500]++
								if nerr, ok := err.(Error); ok && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5503]++
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			return nerr.Temporary()
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[5503]"
//line /snap/go/10455/src/net/dnsclient_unix.go:448
		}() && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5504]++
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			return r.strictErrors()
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[5504]"
//line /snap/go/10455/src/net/dnsclient_unix.go:448
		}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[528010]++
//line /snap/go/10455/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5505]++

//line /snap/go/10455/src/net/dnsclient_unix.go:451
			break
//line /snap/go/10455/src/net/dnsclient_unix.go:451
			// _ = "end of CoverTab[5505]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:452
			_go_fuzz_dep_.CoverTab[528011]++
//line /snap/go/10455/src/net/dnsclient_unix.go:452
			_go_fuzz_dep_.CoverTab[5506]++
//line /snap/go/10455/src/net/dnsclient_unix.go:452
			// _ = "end of CoverTab[5506]"
//line /snap/go/10455/src/net/dnsclient_unix.go:452
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:452
		// _ = "end of CoverTab[5500]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:453
	if _go_fuzz_dep_.CoverTab[786665] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:453
		_go_fuzz_dep_.CoverTab[528155]++
//line /snap/go/10455/src/net/dnsclient_unix.go:453
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:453
		_go_fuzz_dep_.CoverTab[528156]++
//line /snap/go/10455/src/net/dnsclient_unix.go:453
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:453
	// _ = "end of CoverTab[5491]"
//line /snap/go/10455/src/net/dnsclient_unix.go:453
	_go_fuzz_dep_.CoverTab[5492]++
							if err == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:454
		_go_fuzz_dep_.CoverTab[528012]++
//line /snap/go/10455/src/net/dnsclient_unix.go:454
		_go_fuzz_dep_.CoverTab[5507]++
								return p, server, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:455
		// _ = "end of CoverTab[5507]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:456
		_go_fuzz_dep_.CoverTab[528013]++
//line /snap/go/10455/src/net/dnsclient_unix.go:456
		_go_fuzz_dep_.CoverTab[5508]++
//line /snap/go/10455/src/net/dnsclient_unix.go:456
		// _ = "end of CoverTab[5508]"
//line /snap/go/10455/src/net/dnsclient_unix.go:456
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:456
	// _ = "end of CoverTab[5492]"
//line /snap/go/10455/src/net/dnsclient_unix.go:456
	_go_fuzz_dep_.CoverTab[5493]++
							if err, ok := err.(*DNSError); ok {
//line /snap/go/10455/src/net/dnsclient_unix.go:457
		_go_fuzz_dep_.CoverTab[528014]++
//line /snap/go/10455/src/net/dnsclient_unix.go:457
		_go_fuzz_dep_.CoverTab[5509]++

//line /snap/go/10455/src/net/dnsclient_unix.go:461
		err.Name = name
//line /snap/go/10455/src/net/dnsclient_unix.go:461
		// _ = "end of CoverTab[5509]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:462
		_go_fuzz_dep_.CoverTab[528015]++
//line /snap/go/10455/src/net/dnsclient_unix.go:462
		_go_fuzz_dep_.CoverTab[5510]++
//line /snap/go/10455/src/net/dnsclient_unix.go:462
		// _ = "end of CoverTab[5510]"
//line /snap/go/10455/src/net/dnsclient_unix.go:462
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:462
	// _ = "end of CoverTab[5493]"
//line /snap/go/10455/src/net/dnsclient_unix.go:462
	_go_fuzz_dep_.CoverTab[5494]++
							return dnsmessage.Parser{}, "", err
//line /snap/go/10455/src/net/dnsclient_unix.go:463
	// _ = "end of CoverTab[5494]"
}

// avoidDNS reports whether this is a hostname for which we should not
//line /snap/go/10455/src/net/dnsclient_unix.go:466
// use DNS. Currently this includes only .onion, per RFC 7686. See
//line /snap/go/10455/src/net/dnsclient_unix.go:466
// golang.org/issue/13705. Does not cover .local names (RFC 6762),
//line /snap/go/10455/src/net/dnsclient_unix.go:466
// see golang.org/issue/16739.
//line /snap/go/10455/src/net/dnsclient_unix.go:470
func avoidDNS(name string) bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:470
	_go_fuzz_dep_.CoverTab[5511]++
							if name == "" {
//line /snap/go/10455/src/net/dnsclient_unix.go:471
		_go_fuzz_dep_.CoverTab[528016]++
//line /snap/go/10455/src/net/dnsclient_unix.go:471
		_go_fuzz_dep_.CoverTab[5514]++
								return true
//line /snap/go/10455/src/net/dnsclient_unix.go:472
		// _ = "end of CoverTab[5514]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:473
		_go_fuzz_dep_.CoverTab[528017]++
//line /snap/go/10455/src/net/dnsclient_unix.go:473
		_go_fuzz_dep_.CoverTab[5515]++
//line /snap/go/10455/src/net/dnsclient_unix.go:473
		// _ = "end of CoverTab[5515]"
//line /snap/go/10455/src/net/dnsclient_unix.go:473
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:473
	// _ = "end of CoverTab[5511]"
//line /snap/go/10455/src/net/dnsclient_unix.go:473
	_go_fuzz_dep_.CoverTab[5512]++
							if name[len(name)-1] == '.' {
//line /snap/go/10455/src/net/dnsclient_unix.go:474
		_go_fuzz_dep_.CoverTab[528018]++
//line /snap/go/10455/src/net/dnsclient_unix.go:474
		_go_fuzz_dep_.CoverTab[5516]++
								name = name[:len(name)-1]
//line /snap/go/10455/src/net/dnsclient_unix.go:475
		// _ = "end of CoverTab[5516]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:476
		_go_fuzz_dep_.CoverTab[528019]++
//line /snap/go/10455/src/net/dnsclient_unix.go:476
		_go_fuzz_dep_.CoverTab[5517]++
//line /snap/go/10455/src/net/dnsclient_unix.go:476
		// _ = "end of CoverTab[5517]"
//line /snap/go/10455/src/net/dnsclient_unix.go:476
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:476
	// _ = "end of CoverTab[5512]"
//line /snap/go/10455/src/net/dnsclient_unix.go:476
	_go_fuzz_dep_.CoverTab[5513]++
							return stringsHasSuffixFold(name, ".onion")
//line /snap/go/10455/src/net/dnsclient_unix.go:477
	// _ = "end of CoverTab[5513]"
}

// nameList returns a list of names for sequential DNS queries.
func (conf *dnsConfig) nameList(name string) []string {
//line /snap/go/10455/src/net/dnsclient_unix.go:481
	_go_fuzz_dep_.CoverTab[5518]++
							if avoidDNS(name) {
//line /snap/go/10455/src/net/dnsclient_unix.go:482
		_go_fuzz_dep_.CoverTab[528020]++
//line /snap/go/10455/src/net/dnsclient_unix.go:482
		_go_fuzz_dep_.CoverTab[5525]++
								return nil
//line /snap/go/10455/src/net/dnsclient_unix.go:483
		// _ = "end of CoverTab[5525]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:484
		_go_fuzz_dep_.CoverTab[528021]++
//line /snap/go/10455/src/net/dnsclient_unix.go:484
		_go_fuzz_dep_.CoverTab[5526]++
//line /snap/go/10455/src/net/dnsclient_unix.go:484
		// _ = "end of CoverTab[5526]"
//line /snap/go/10455/src/net/dnsclient_unix.go:484
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:484
	// _ = "end of CoverTab[5518]"
//line /snap/go/10455/src/net/dnsclient_unix.go:484
	_go_fuzz_dep_.CoverTab[5519]++

//line /snap/go/10455/src/net/dnsclient_unix.go:487
	l := len(name)
	rooted := l > 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:488
		_go_fuzz_dep_.CoverTab[5527]++
//line /snap/go/10455/src/net/dnsclient_unix.go:488
		return name[l-1] == '.'
//line /snap/go/10455/src/net/dnsclient_unix.go:488
		// _ = "end of CoverTab[5527]"
//line /snap/go/10455/src/net/dnsclient_unix.go:488
	}()
							if l > 254 || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[5528]++
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		return l == 254 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:489
			_go_fuzz_dep_.CoverTab[5529]++
//line /snap/go/10455/src/net/dnsclient_unix.go:489
			return !rooted
//line /snap/go/10455/src/net/dnsclient_unix.go:489
			// _ = "end of CoverTab[5529]"
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		}()
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		// _ = "end of CoverTab[5528]"
//line /snap/go/10455/src/net/dnsclient_unix.go:489
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[528022]++
//line /snap/go/10455/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[5530]++
								return nil
//line /snap/go/10455/src/net/dnsclient_unix.go:490
		// _ = "end of CoverTab[5530]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:491
		_go_fuzz_dep_.CoverTab[528023]++
//line /snap/go/10455/src/net/dnsclient_unix.go:491
		_go_fuzz_dep_.CoverTab[5531]++
//line /snap/go/10455/src/net/dnsclient_unix.go:491
		// _ = "end of CoverTab[5531]"
//line /snap/go/10455/src/net/dnsclient_unix.go:491
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:491
	// _ = "end of CoverTab[5519]"
//line /snap/go/10455/src/net/dnsclient_unix.go:491
	_go_fuzz_dep_.CoverTab[5520]++

//line /snap/go/10455/src/net/dnsclient_unix.go:494
	if rooted {
//line /snap/go/10455/src/net/dnsclient_unix.go:494
		_go_fuzz_dep_.CoverTab[528024]++
//line /snap/go/10455/src/net/dnsclient_unix.go:494
		_go_fuzz_dep_.CoverTab[5532]++
								return []string{name}
//line /snap/go/10455/src/net/dnsclient_unix.go:495
		// _ = "end of CoverTab[5532]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:496
		_go_fuzz_dep_.CoverTab[528025]++
//line /snap/go/10455/src/net/dnsclient_unix.go:496
		_go_fuzz_dep_.CoverTab[5533]++
//line /snap/go/10455/src/net/dnsclient_unix.go:496
		// _ = "end of CoverTab[5533]"
//line /snap/go/10455/src/net/dnsclient_unix.go:496
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:496
	// _ = "end of CoverTab[5520]"
//line /snap/go/10455/src/net/dnsclient_unix.go:496
	_go_fuzz_dep_.CoverTab[5521]++

							hasNdots := count(name, '.') >= conf.ndots
							name += "."
							l++

//line /snap/go/10455/src/net/dnsclient_unix.go:503
	names := make([]string, 0, 1+len(conf.search))

	if hasNdots {
//line /snap/go/10455/src/net/dnsclient_unix.go:505
		_go_fuzz_dep_.CoverTab[528026]++
//line /snap/go/10455/src/net/dnsclient_unix.go:505
		_go_fuzz_dep_.CoverTab[5534]++
								names = append(names, name)
//line /snap/go/10455/src/net/dnsclient_unix.go:506
		// _ = "end of CoverTab[5534]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:507
		_go_fuzz_dep_.CoverTab[528027]++
//line /snap/go/10455/src/net/dnsclient_unix.go:507
		_go_fuzz_dep_.CoverTab[5535]++
//line /snap/go/10455/src/net/dnsclient_unix.go:507
		// _ = "end of CoverTab[5535]"
//line /snap/go/10455/src/net/dnsclient_unix.go:507
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:507
	// _ = "end of CoverTab[5521]"
//line /snap/go/10455/src/net/dnsclient_unix.go:507
	_go_fuzz_dep_.CoverTab[5522]++
//line /snap/go/10455/src/net/dnsclient_unix.go:507
	_go_fuzz_dep_.CoverTab[786666] = 0

							for _, suffix := range conf.search {
//line /snap/go/10455/src/net/dnsclient_unix.go:509
		if _go_fuzz_dep_.CoverTab[786666] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:509
			_go_fuzz_dep_.CoverTab[528157]++
//line /snap/go/10455/src/net/dnsclient_unix.go:509
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:509
			_go_fuzz_dep_.CoverTab[528158]++
//line /snap/go/10455/src/net/dnsclient_unix.go:509
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:509
		_go_fuzz_dep_.CoverTab[786666] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:509
		_go_fuzz_dep_.CoverTab[5536]++
								if l+len(suffix) <= 254 {
//line /snap/go/10455/src/net/dnsclient_unix.go:510
			_go_fuzz_dep_.CoverTab[528028]++
//line /snap/go/10455/src/net/dnsclient_unix.go:510
			_go_fuzz_dep_.CoverTab[5537]++
									names = append(names, name+suffix)
//line /snap/go/10455/src/net/dnsclient_unix.go:511
			// _ = "end of CoverTab[5537]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:512
			_go_fuzz_dep_.CoverTab[528029]++
//line /snap/go/10455/src/net/dnsclient_unix.go:512
			_go_fuzz_dep_.CoverTab[5538]++
//line /snap/go/10455/src/net/dnsclient_unix.go:512
			// _ = "end of CoverTab[5538]"
//line /snap/go/10455/src/net/dnsclient_unix.go:512
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:512
		// _ = "end of CoverTab[5536]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:513
	if _go_fuzz_dep_.CoverTab[786666] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:513
		_go_fuzz_dep_.CoverTab[528159]++
//line /snap/go/10455/src/net/dnsclient_unix.go:513
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:513
		_go_fuzz_dep_.CoverTab[528160]++
//line /snap/go/10455/src/net/dnsclient_unix.go:513
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:513
	// _ = "end of CoverTab[5522]"
//line /snap/go/10455/src/net/dnsclient_unix.go:513
	_go_fuzz_dep_.CoverTab[5523]++

							if !hasNdots {
//line /snap/go/10455/src/net/dnsclient_unix.go:515
		_go_fuzz_dep_.CoverTab[528030]++
//line /snap/go/10455/src/net/dnsclient_unix.go:515
		_go_fuzz_dep_.CoverTab[5539]++
								names = append(names, name)
//line /snap/go/10455/src/net/dnsclient_unix.go:516
		// _ = "end of CoverTab[5539]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:517
		_go_fuzz_dep_.CoverTab[528031]++
//line /snap/go/10455/src/net/dnsclient_unix.go:517
		_go_fuzz_dep_.CoverTab[5540]++
//line /snap/go/10455/src/net/dnsclient_unix.go:517
		// _ = "end of CoverTab[5540]"
//line /snap/go/10455/src/net/dnsclient_unix.go:517
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:517
	// _ = "end of CoverTab[5523]"
//line /snap/go/10455/src/net/dnsclient_unix.go:517
	_go_fuzz_dep_.CoverTab[5524]++
							return names
//line /snap/go/10455/src/net/dnsclient_unix.go:518
	// _ = "end of CoverTab[5524]"
}

// hostLookupOrder specifies the order of LookupHost lookup strategies.
//line /snap/go/10455/src/net/dnsclient_unix.go:521
// It is basically a simplified representation of nsswitch.conf.
//line /snap/go/10455/src/net/dnsclient_unix.go:521
// "files" means /etc/hosts.
//line /snap/go/10455/src/net/dnsclient_unix.go:524
type hostLookupOrder int

const (
	// hostLookupCgo means defer to cgo.
	hostLookupCgo		hostLookupOrder	= iota
	hostLookupFilesDNS			// files first
	hostLookupDNSFiles			// dns first
	hostLookupFiles				// only files
	hostLookupDNS				// only DNS
)

var lookupOrderName = map[hostLookupOrder]string{
	hostLookupCgo:		"cgo",
	hostLookupFilesDNS:	"files,dns",
	hostLookupDNSFiles:	"dns,files",
	hostLookupFiles:	"files",
	hostLookupDNS:		"dns",
}

func (o hostLookupOrder) String() string {
//line /snap/go/10455/src/net/dnsclient_unix.go:543
	_go_fuzz_dep_.CoverTab[5541]++
							if s, ok := lookupOrderName[o]; ok {
//line /snap/go/10455/src/net/dnsclient_unix.go:544
		_go_fuzz_dep_.CoverTab[528032]++
//line /snap/go/10455/src/net/dnsclient_unix.go:544
		_go_fuzz_dep_.CoverTab[5543]++
								return s
//line /snap/go/10455/src/net/dnsclient_unix.go:545
		// _ = "end of CoverTab[5543]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:546
		_go_fuzz_dep_.CoverTab[528033]++
//line /snap/go/10455/src/net/dnsclient_unix.go:546
		_go_fuzz_dep_.CoverTab[5544]++
//line /snap/go/10455/src/net/dnsclient_unix.go:546
		// _ = "end of CoverTab[5544]"
//line /snap/go/10455/src/net/dnsclient_unix.go:546
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:546
	// _ = "end of CoverTab[5541]"
//line /snap/go/10455/src/net/dnsclient_unix.go:546
	_go_fuzz_dep_.CoverTab[5542]++
							return "hostLookupOrder=" + itoa.Itoa(int(o)) + "??"
//line /snap/go/10455/src/net/dnsclient_unix.go:547
	// _ = "end of CoverTab[5542]"
}

func (r *Resolver) goLookupHostOrder(ctx context.Context, name string, order hostLookupOrder, conf *dnsConfig) (addrs []string, err error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:550
	_go_fuzz_dep_.CoverTab[5545]++
							if order == hostLookupFilesDNS || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[5549]++
//line /snap/go/10455/src/net/dnsclient_unix.go:551
		return order == hostLookupFiles
//line /snap/go/10455/src/net/dnsclient_unix.go:551
		// _ = "end of CoverTab[5549]"
//line /snap/go/10455/src/net/dnsclient_unix.go:551
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[528034]++
//line /snap/go/10455/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[5550]++

								addrs, _ = lookupStaticHost(name)
								if len(addrs) > 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:554
			_go_fuzz_dep_.CoverTab[528036]++
//line /snap/go/10455/src/net/dnsclient_unix.go:554
			_go_fuzz_dep_.CoverTab[5552]++
									return
//line /snap/go/10455/src/net/dnsclient_unix.go:555
			// _ = "end of CoverTab[5552]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:556
			_go_fuzz_dep_.CoverTab[528037]++
//line /snap/go/10455/src/net/dnsclient_unix.go:556
			_go_fuzz_dep_.CoverTab[5553]++
//line /snap/go/10455/src/net/dnsclient_unix.go:556
			// _ = "end of CoverTab[5553]"
//line /snap/go/10455/src/net/dnsclient_unix.go:556
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:556
		// _ = "end of CoverTab[5550]"
//line /snap/go/10455/src/net/dnsclient_unix.go:556
		_go_fuzz_dep_.CoverTab[5551]++

								if order == hostLookupFiles {
//line /snap/go/10455/src/net/dnsclient_unix.go:558
			_go_fuzz_dep_.CoverTab[528038]++
//line /snap/go/10455/src/net/dnsclient_unix.go:558
			_go_fuzz_dep_.CoverTab[5554]++
									return nil, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /snap/go/10455/src/net/dnsclient_unix.go:559
			// _ = "end of CoverTab[5554]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:560
			_go_fuzz_dep_.CoverTab[528039]++
//line /snap/go/10455/src/net/dnsclient_unix.go:560
			_go_fuzz_dep_.CoverTab[5555]++
//line /snap/go/10455/src/net/dnsclient_unix.go:560
			// _ = "end of CoverTab[5555]"
//line /snap/go/10455/src/net/dnsclient_unix.go:560
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:560
		// _ = "end of CoverTab[5551]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:561
		_go_fuzz_dep_.CoverTab[528035]++
//line /snap/go/10455/src/net/dnsclient_unix.go:561
		_go_fuzz_dep_.CoverTab[5556]++
//line /snap/go/10455/src/net/dnsclient_unix.go:561
		// _ = "end of CoverTab[5556]"
//line /snap/go/10455/src/net/dnsclient_unix.go:561
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:561
	// _ = "end of CoverTab[5545]"
//line /snap/go/10455/src/net/dnsclient_unix.go:561
	_go_fuzz_dep_.CoverTab[5546]++
							ips, _, err := r.goLookupIPCNAMEOrder(ctx, "ip", name, order, conf)
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:563
		_go_fuzz_dep_.CoverTab[528040]++
//line /snap/go/10455/src/net/dnsclient_unix.go:563
		_go_fuzz_dep_.CoverTab[5557]++
								return
//line /snap/go/10455/src/net/dnsclient_unix.go:564
		// _ = "end of CoverTab[5557]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:565
		_go_fuzz_dep_.CoverTab[528041]++
//line /snap/go/10455/src/net/dnsclient_unix.go:565
		_go_fuzz_dep_.CoverTab[5558]++
//line /snap/go/10455/src/net/dnsclient_unix.go:565
		// _ = "end of CoverTab[5558]"
//line /snap/go/10455/src/net/dnsclient_unix.go:565
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:565
	// _ = "end of CoverTab[5546]"
//line /snap/go/10455/src/net/dnsclient_unix.go:565
	_go_fuzz_dep_.CoverTab[5547]++
							addrs = make([]string, 0, len(ips))
//line /snap/go/10455/src/net/dnsclient_unix.go:566
	_go_fuzz_dep_.CoverTab[786667] = 0
							for _, ip := range ips {
//line /snap/go/10455/src/net/dnsclient_unix.go:567
		if _go_fuzz_dep_.CoverTab[786667] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:567
			_go_fuzz_dep_.CoverTab[528161]++
//line /snap/go/10455/src/net/dnsclient_unix.go:567
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:567
			_go_fuzz_dep_.CoverTab[528162]++
//line /snap/go/10455/src/net/dnsclient_unix.go:567
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:567
		_go_fuzz_dep_.CoverTab[786667] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:567
		_go_fuzz_dep_.CoverTab[5559]++
								addrs = append(addrs, ip.String())
//line /snap/go/10455/src/net/dnsclient_unix.go:568
		// _ = "end of CoverTab[5559]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:569
	if _go_fuzz_dep_.CoverTab[786667] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:569
		_go_fuzz_dep_.CoverTab[528163]++
//line /snap/go/10455/src/net/dnsclient_unix.go:569
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:569
		_go_fuzz_dep_.CoverTab[528164]++
//line /snap/go/10455/src/net/dnsclient_unix.go:569
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:569
	// _ = "end of CoverTab[5547]"
//line /snap/go/10455/src/net/dnsclient_unix.go:569
	_go_fuzz_dep_.CoverTab[5548]++
							return
//line /snap/go/10455/src/net/dnsclient_unix.go:570
	// _ = "end of CoverTab[5548]"
}

// lookup entries from /etc/hosts
func goLookupIPFiles(name string) (addrs []IPAddr, canonical string) {
//line /snap/go/10455/src/net/dnsclient_unix.go:574
	_go_fuzz_dep_.CoverTab[5560]++
							addr, canonical := lookupStaticHost(name)
//line /snap/go/10455/src/net/dnsclient_unix.go:575
	_go_fuzz_dep_.CoverTab[786668] = 0
							for _, haddr := range addr {
//line /snap/go/10455/src/net/dnsclient_unix.go:576
		if _go_fuzz_dep_.CoverTab[786668] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:576
			_go_fuzz_dep_.CoverTab[528165]++
//line /snap/go/10455/src/net/dnsclient_unix.go:576
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:576
			_go_fuzz_dep_.CoverTab[528166]++
//line /snap/go/10455/src/net/dnsclient_unix.go:576
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:576
		_go_fuzz_dep_.CoverTab[786668] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:576
		_go_fuzz_dep_.CoverTab[5562]++
								haddr, zone := splitHostZone(haddr)
								if ip := ParseIP(haddr); ip != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:578
			_go_fuzz_dep_.CoverTab[528042]++
//line /snap/go/10455/src/net/dnsclient_unix.go:578
			_go_fuzz_dep_.CoverTab[5563]++
									addr := IPAddr{IP: ip, Zone: zone}
									addrs = append(addrs, addr)
//line /snap/go/10455/src/net/dnsclient_unix.go:580
			// _ = "end of CoverTab[5563]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:581
			_go_fuzz_dep_.CoverTab[528043]++
//line /snap/go/10455/src/net/dnsclient_unix.go:581
			_go_fuzz_dep_.CoverTab[5564]++
//line /snap/go/10455/src/net/dnsclient_unix.go:581
			// _ = "end of CoverTab[5564]"
//line /snap/go/10455/src/net/dnsclient_unix.go:581
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:581
		// _ = "end of CoverTab[5562]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:582
	if _go_fuzz_dep_.CoverTab[786668] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:582
		_go_fuzz_dep_.CoverTab[528167]++
//line /snap/go/10455/src/net/dnsclient_unix.go:582
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:582
		_go_fuzz_dep_.CoverTab[528168]++
//line /snap/go/10455/src/net/dnsclient_unix.go:582
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:582
	// _ = "end of CoverTab[5560]"
//line /snap/go/10455/src/net/dnsclient_unix.go:582
	_go_fuzz_dep_.CoverTab[5561]++
							sortByRFC6724(addrs)
							return addrs, canonical
//line /snap/go/10455/src/net/dnsclient_unix.go:584
	// _ = "end of CoverTab[5561]"
}

// goLookupIP is the native Go implementation of LookupIP.
//line /snap/go/10455/src/net/dnsclient_unix.go:587
// The libc versions are in cgo_*.go.
//line /snap/go/10455/src/net/dnsclient_unix.go:589
func (r *Resolver) goLookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:589
	_go_fuzz_dep_.CoverTab[5565]++
							order, conf := systemConf().hostLookupOrder(r, host)
							addrs, _, err = r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
							return
//line /snap/go/10455/src/net/dnsclient_unix.go:592
	// _ = "end of CoverTab[5565]"
}

func (r *Resolver) goLookupIPCNAMEOrder(ctx context.Context, network, name string, order hostLookupOrder, conf *dnsConfig) (addrs []IPAddr, cname dnsmessage.Name, err error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:595
	_go_fuzz_dep_.CoverTab[5566]++
							if order == hostLookupFilesDNS || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[5576]++
//line /snap/go/10455/src/net/dnsclient_unix.go:596
		return order == hostLookupFiles
//line /snap/go/10455/src/net/dnsclient_unix.go:596
		// _ = "end of CoverTab[5576]"
//line /snap/go/10455/src/net/dnsclient_unix.go:596
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[528044]++
//line /snap/go/10455/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[5577]++
								var canonical string
								addrs, canonical = goLookupIPFiles(name)

								if len(addrs) > 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:600
			_go_fuzz_dep_.CoverTab[528046]++
//line /snap/go/10455/src/net/dnsclient_unix.go:600
			_go_fuzz_dep_.CoverTab[5579]++
									var err error
									cname, err = dnsmessage.NewName(canonical)
									if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:603
				_go_fuzz_dep_.CoverTab[528048]++
//line /snap/go/10455/src/net/dnsclient_unix.go:603
				_go_fuzz_dep_.CoverTab[5581]++
										return nil, dnsmessage.Name{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:604
				// _ = "end of CoverTab[5581]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:605
				_go_fuzz_dep_.CoverTab[528049]++
//line /snap/go/10455/src/net/dnsclient_unix.go:605
				_go_fuzz_dep_.CoverTab[5582]++
//line /snap/go/10455/src/net/dnsclient_unix.go:605
				// _ = "end of CoverTab[5582]"
//line /snap/go/10455/src/net/dnsclient_unix.go:605
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:605
			// _ = "end of CoverTab[5579]"
//line /snap/go/10455/src/net/dnsclient_unix.go:605
			_go_fuzz_dep_.CoverTab[5580]++
									return addrs, cname, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:606
			// _ = "end of CoverTab[5580]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:607
			_go_fuzz_dep_.CoverTab[528047]++
//line /snap/go/10455/src/net/dnsclient_unix.go:607
			_go_fuzz_dep_.CoverTab[5583]++
//line /snap/go/10455/src/net/dnsclient_unix.go:607
			// _ = "end of CoverTab[5583]"
//line /snap/go/10455/src/net/dnsclient_unix.go:607
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:607
		// _ = "end of CoverTab[5577]"
//line /snap/go/10455/src/net/dnsclient_unix.go:607
		_go_fuzz_dep_.CoverTab[5578]++

								if order == hostLookupFiles {
//line /snap/go/10455/src/net/dnsclient_unix.go:609
			_go_fuzz_dep_.CoverTab[528050]++
//line /snap/go/10455/src/net/dnsclient_unix.go:609
			_go_fuzz_dep_.CoverTab[5584]++
									return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /snap/go/10455/src/net/dnsclient_unix.go:610
			// _ = "end of CoverTab[5584]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:611
			_go_fuzz_dep_.CoverTab[528051]++
//line /snap/go/10455/src/net/dnsclient_unix.go:611
			_go_fuzz_dep_.CoverTab[5585]++
//line /snap/go/10455/src/net/dnsclient_unix.go:611
			// _ = "end of CoverTab[5585]"
//line /snap/go/10455/src/net/dnsclient_unix.go:611
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:611
		// _ = "end of CoverTab[5578]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:612
		_go_fuzz_dep_.CoverTab[528045]++
//line /snap/go/10455/src/net/dnsclient_unix.go:612
		_go_fuzz_dep_.CoverTab[5586]++
//line /snap/go/10455/src/net/dnsclient_unix.go:612
		// _ = "end of CoverTab[5586]"
//line /snap/go/10455/src/net/dnsclient_unix.go:612
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:612
	// _ = "end of CoverTab[5566]"
//line /snap/go/10455/src/net/dnsclient_unix.go:612
	_go_fuzz_dep_.CoverTab[5567]++

							if !isDomainName(name) {
//line /snap/go/10455/src/net/dnsclient_unix.go:614
		_go_fuzz_dep_.CoverTab[528052]++
//line /snap/go/10455/src/net/dnsclient_unix.go:614
		_go_fuzz_dep_.CoverTab[5587]++

								return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /snap/go/10455/src/net/dnsclient_unix.go:616
		// _ = "end of CoverTab[5587]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:617
		_go_fuzz_dep_.CoverTab[528053]++
//line /snap/go/10455/src/net/dnsclient_unix.go:617
		_go_fuzz_dep_.CoverTab[5588]++
//line /snap/go/10455/src/net/dnsclient_unix.go:617
		// _ = "end of CoverTab[5588]"
//line /snap/go/10455/src/net/dnsclient_unix.go:617
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:617
	// _ = "end of CoverTab[5567]"
//line /snap/go/10455/src/net/dnsclient_unix.go:617
	_go_fuzz_dep_.CoverTab[5568]++
							type result struct {
		p	dnsmessage.Parser
		server	string
		error
	}

	if conf == nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:624
		_go_fuzz_dep_.CoverTab[528054]++
//line /snap/go/10455/src/net/dnsclient_unix.go:624
		_go_fuzz_dep_.CoverTab[5589]++
								conf = getSystemDNSConfig()
//line /snap/go/10455/src/net/dnsclient_unix.go:625
		// _ = "end of CoverTab[5589]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:626
		_go_fuzz_dep_.CoverTab[528055]++
//line /snap/go/10455/src/net/dnsclient_unix.go:626
		_go_fuzz_dep_.CoverTab[5590]++
//line /snap/go/10455/src/net/dnsclient_unix.go:626
		// _ = "end of CoverTab[5590]"
//line /snap/go/10455/src/net/dnsclient_unix.go:626
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:626
	// _ = "end of CoverTab[5568]"
//line /snap/go/10455/src/net/dnsclient_unix.go:626
	_go_fuzz_dep_.CoverTab[5569]++

							lane := make(chan result, 1)
							qtypes := []dnsmessage.Type{dnsmessage.TypeA, dnsmessage.TypeAAAA}
							if network == "CNAME" {
//line /snap/go/10455/src/net/dnsclient_unix.go:630
		_go_fuzz_dep_.CoverTab[528056]++
//line /snap/go/10455/src/net/dnsclient_unix.go:630
		_go_fuzz_dep_.CoverTab[5591]++
								qtypes = append(qtypes, dnsmessage.TypeCNAME)
//line /snap/go/10455/src/net/dnsclient_unix.go:631
		// _ = "end of CoverTab[5591]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:632
		_go_fuzz_dep_.CoverTab[528057]++
//line /snap/go/10455/src/net/dnsclient_unix.go:632
		_go_fuzz_dep_.CoverTab[5592]++
//line /snap/go/10455/src/net/dnsclient_unix.go:632
		// _ = "end of CoverTab[5592]"
//line /snap/go/10455/src/net/dnsclient_unix.go:632
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:632
	// _ = "end of CoverTab[5569]"
//line /snap/go/10455/src/net/dnsclient_unix.go:632
	_go_fuzz_dep_.CoverTab[5570]++
							switch ipVersion(network) {
	case '4':
//line /snap/go/10455/src/net/dnsclient_unix.go:634
		_go_fuzz_dep_.CoverTab[528058]++
//line /snap/go/10455/src/net/dnsclient_unix.go:634
		_go_fuzz_dep_.CoverTab[5593]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeA}
//line /snap/go/10455/src/net/dnsclient_unix.go:635
		// _ = "end of CoverTab[5593]"
	case '6':
//line /snap/go/10455/src/net/dnsclient_unix.go:636
		_go_fuzz_dep_.CoverTab[528059]++
//line /snap/go/10455/src/net/dnsclient_unix.go:636
		_go_fuzz_dep_.CoverTab[5594]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeAAAA}
//line /snap/go/10455/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[5594]"
//line /snap/go/10455/src/net/dnsclient_unix.go:637
	default:
//line /snap/go/10455/src/net/dnsclient_unix.go:637
		_go_fuzz_dep_.CoverTab[528060]++
//line /snap/go/10455/src/net/dnsclient_unix.go:637
		_go_fuzz_dep_.CoverTab[5595]++
//line /snap/go/10455/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[5595]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:638
	// _ = "end of CoverTab[5570]"
//line /snap/go/10455/src/net/dnsclient_unix.go:638
	_go_fuzz_dep_.CoverTab[5571]++
							var queryFn func(fqdn string, qtype dnsmessage.Type)
							var responseFn func(fqdn string, qtype dnsmessage.Type) result
							if conf.singleRequest {
//line /snap/go/10455/src/net/dnsclient_unix.go:641
		_go_fuzz_dep_.CoverTab[528061]++
//line /snap/go/10455/src/net/dnsclient_unix.go:641
		_go_fuzz_dep_.CoverTab[5596]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /snap/go/10455/src/net/dnsclient_unix.go:642
			_go_fuzz_dep_.CoverTab[5598]++
//line /snap/go/10455/src/net/dnsclient_unix.go:642
			// _ = "end of CoverTab[5598]"
//line /snap/go/10455/src/net/dnsclient_unix.go:642
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:642
		// _ = "end of CoverTab[5596]"
//line /snap/go/10455/src/net/dnsclient_unix.go:642
		_go_fuzz_dep_.CoverTab[5597]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /snap/go/10455/src/net/dnsclient_unix.go:643
			_go_fuzz_dep_.CoverTab[5599]++
									dnsWaitGroup.Add(1)
									defer dnsWaitGroup.Done()
									p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
									return result{p, server, err}
//line /snap/go/10455/src/net/dnsclient_unix.go:647
			// _ = "end of CoverTab[5599]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:648
		// _ = "end of CoverTab[5597]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:649
		_go_fuzz_dep_.CoverTab[528062]++
//line /snap/go/10455/src/net/dnsclient_unix.go:649
		_go_fuzz_dep_.CoverTab[5600]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /snap/go/10455/src/net/dnsclient_unix.go:650
			_go_fuzz_dep_.CoverTab[5602]++
									dnsWaitGroup.Add(1)
//line /snap/go/10455/src/net/dnsclient_unix.go:651
			_curRoutineNum6_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/dnsclient_unix.go:651
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum6_)
									go func(qtype dnsmessage.Type) {
//line /snap/go/10455/src/net/dnsclient_unix.go:652
				_go_fuzz_dep_.CoverTab[5603]++
//line /snap/go/10455/src/net/dnsclient_unix.go:652
				defer func() {
//line /snap/go/10455/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.CoverTab[5604]++
//line /snap/go/10455/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum6_)
//line /snap/go/10455/src/net/dnsclient_unix.go:652
					// _ = "end of CoverTab[5604]"
//line /snap/go/10455/src/net/dnsclient_unix.go:652
				}()
										p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
										lane <- result{p, server, err}
										dnsWaitGroup.Done()
//line /snap/go/10455/src/net/dnsclient_unix.go:655
				// _ = "end of CoverTab[5603]"
			}(qtype)
//line /snap/go/10455/src/net/dnsclient_unix.go:656
			// _ = "end of CoverTab[5602]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:657
		// _ = "end of CoverTab[5600]"
//line /snap/go/10455/src/net/dnsclient_unix.go:657
		_go_fuzz_dep_.CoverTab[5601]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /snap/go/10455/src/net/dnsclient_unix.go:658
			_go_fuzz_dep_.CoverTab[5605]++
									return <-lane
//line /snap/go/10455/src/net/dnsclient_unix.go:659
			// _ = "end of CoverTab[5605]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:660
		// _ = "end of CoverTab[5601]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:661
	// _ = "end of CoverTab[5571]"
//line /snap/go/10455/src/net/dnsclient_unix.go:661
	_go_fuzz_dep_.CoverTab[5572]++
							var lastErr error
//line /snap/go/10455/src/net/dnsclient_unix.go:662
	_go_fuzz_dep_.CoverTab[786669] = 0
							for _, fqdn := range conf.nameList(name) {
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		if _go_fuzz_dep_.CoverTab[786669] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:663
			_go_fuzz_dep_.CoverTab[528169]++
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:663
			_go_fuzz_dep_.CoverTab[528170]++
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		_go_fuzz_dep_.CoverTab[786669] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		_go_fuzz_dep_.CoverTab[5606]++
//line /snap/go/10455/src/net/dnsclient_unix.go:663
		_go_fuzz_dep_.CoverTab[786670] = 0
								for _, qtype := range qtypes {
//line /snap/go/10455/src/net/dnsclient_unix.go:664
			if _go_fuzz_dep_.CoverTab[786670] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:664
				_go_fuzz_dep_.CoverTab[528173]++
//line /snap/go/10455/src/net/dnsclient_unix.go:664
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:664
				_go_fuzz_dep_.CoverTab[528174]++
//line /snap/go/10455/src/net/dnsclient_unix.go:664
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:664
			_go_fuzz_dep_.CoverTab[786670] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:664
			_go_fuzz_dep_.CoverTab[5610]++
									queryFn(fqdn, qtype)
//line /snap/go/10455/src/net/dnsclient_unix.go:665
			// _ = "end of CoverTab[5610]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:666
		if _go_fuzz_dep_.CoverTab[786670] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:666
			_go_fuzz_dep_.CoverTab[528175]++
//line /snap/go/10455/src/net/dnsclient_unix.go:666
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:666
			_go_fuzz_dep_.CoverTab[528176]++
//line /snap/go/10455/src/net/dnsclient_unix.go:666
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:666
		// _ = "end of CoverTab[5606]"
//line /snap/go/10455/src/net/dnsclient_unix.go:666
		_go_fuzz_dep_.CoverTab[5607]++
								hitStrictError := false
//line /snap/go/10455/src/net/dnsclient_unix.go:667
		_go_fuzz_dep_.CoverTab[786671] = 0
								for _, qtype := range qtypes {
//line /snap/go/10455/src/net/dnsclient_unix.go:668
			if _go_fuzz_dep_.CoverTab[786671] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:668
				_go_fuzz_dep_.CoverTab[528177]++
//line /snap/go/10455/src/net/dnsclient_unix.go:668
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:668
				_go_fuzz_dep_.CoverTab[528178]++
//line /snap/go/10455/src/net/dnsclient_unix.go:668
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:668
			_go_fuzz_dep_.CoverTab[786671] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:668
			_go_fuzz_dep_.CoverTab[5611]++
									result := responseFn(fqdn, qtype)
									if result.error != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:670
				_go_fuzz_dep_.CoverTab[528063]++
//line /snap/go/10455/src/net/dnsclient_unix.go:670
				_go_fuzz_dep_.CoverTab[5613]++
										if nerr, ok := result.error.(Error); ok && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5615]++
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					return nerr.Temporary()
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[5615]"
//line /snap/go/10455/src/net/dnsclient_unix.go:671
				}() && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5616]++
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					return r.strictErrors()
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[5616]"
//line /snap/go/10455/src/net/dnsclient_unix.go:671
				}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[528065]++
//line /snap/go/10455/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5617]++

											hitStrictError = true
											lastErr = result.error
//line /snap/go/10455/src/net/dnsclient_unix.go:674
					// _ = "end of CoverTab[5617]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:675
					_go_fuzz_dep_.CoverTab[528066]++
//line /snap/go/10455/src/net/dnsclient_unix.go:675
					_go_fuzz_dep_.CoverTab[5618]++
//line /snap/go/10455/src/net/dnsclient_unix.go:675
					if lastErr == nil || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[5619]++
//line /snap/go/10455/src/net/dnsclient_unix.go:675
						return fqdn == name+"."
//line /snap/go/10455/src/net/dnsclient_unix.go:675
						// _ = "end of CoverTab[5619]"
//line /snap/go/10455/src/net/dnsclient_unix.go:675
					}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[528067]++
//line /snap/go/10455/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[5620]++

												lastErr = result.error
//line /snap/go/10455/src/net/dnsclient_unix.go:677
						// _ = "end of CoverTab[5620]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:678
						_go_fuzz_dep_.CoverTab[528068]++
//line /snap/go/10455/src/net/dnsclient_unix.go:678
						_go_fuzz_dep_.CoverTab[5621]++
//line /snap/go/10455/src/net/dnsclient_unix.go:678
						// _ = "end of CoverTab[5621]"
//line /snap/go/10455/src/net/dnsclient_unix.go:678
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:678
					// _ = "end of CoverTab[5618]"
//line /snap/go/10455/src/net/dnsclient_unix.go:678
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:678
				// _ = "end of CoverTab[5613]"
//line /snap/go/10455/src/net/dnsclient_unix.go:678
				_go_fuzz_dep_.CoverTab[5614]++
										continue
//line /snap/go/10455/src/net/dnsclient_unix.go:679
				// _ = "end of CoverTab[5614]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:680
				_go_fuzz_dep_.CoverTab[528064]++
//line /snap/go/10455/src/net/dnsclient_unix.go:680
				_go_fuzz_dep_.CoverTab[5622]++
//line /snap/go/10455/src/net/dnsclient_unix.go:680
				// _ = "end of CoverTab[5622]"
//line /snap/go/10455/src/net/dnsclient_unix.go:680
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:680
			// _ = "end of CoverTab[5611]"
//line /snap/go/10455/src/net/dnsclient_unix.go:680
			_go_fuzz_dep_.CoverTab[5612]++

//line /snap/go/10455/src/net/dnsclient_unix.go:697
		loop:
			for {
//line /snap/go/10455/src/net/dnsclient_unix.go:698
				_go_fuzz_dep_.CoverTab[5623]++
										h, err := result.p.AnswerHeader()
										if err != nil && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[5626]++
//line /snap/go/10455/src/net/dnsclient_unix.go:700
					return err != dnsmessage.ErrSectionDone
//line /snap/go/10455/src/net/dnsclient_unix.go:700
					// _ = "end of CoverTab[5626]"
//line /snap/go/10455/src/net/dnsclient_unix.go:700
				}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[528069]++
//line /snap/go/10455/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[5627]++
											lastErr = &DNSError{
						Err:	"cannot marshal DNS message",
						Name:	name,
						Server:	result.server,
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:705
					// _ = "end of CoverTab[5627]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:706
					_go_fuzz_dep_.CoverTab[528070]++
//line /snap/go/10455/src/net/dnsclient_unix.go:706
					_go_fuzz_dep_.CoverTab[5628]++
//line /snap/go/10455/src/net/dnsclient_unix.go:706
					// _ = "end of CoverTab[5628]"
//line /snap/go/10455/src/net/dnsclient_unix.go:706
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:706
				// _ = "end of CoverTab[5623]"
//line /snap/go/10455/src/net/dnsclient_unix.go:706
				_go_fuzz_dep_.CoverTab[5624]++
										if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:707
					_go_fuzz_dep_.CoverTab[528071]++
//line /snap/go/10455/src/net/dnsclient_unix.go:707
					_go_fuzz_dep_.CoverTab[5629]++
											break
//line /snap/go/10455/src/net/dnsclient_unix.go:708
					// _ = "end of CoverTab[5629]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:709
					_go_fuzz_dep_.CoverTab[528072]++
//line /snap/go/10455/src/net/dnsclient_unix.go:709
					_go_fuzz_dep_.CoverTab[5630]++
//line /snap/go/10455/src/net/dnsclient_unix.go:709
					// _ = "end of CoverTab[5630]"
//line /snap/go/10455/src/net/dnsclient_unix.go:709
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:709
				// _ = "end of CoverTab[5624]"
//line /snap/go/10455/src/net/dnsclient_unix.go:709
				_go_fuzz_dep_.CoverTab[5625]++
										switch h.Type {
				case dnsmessage.TypeA:
//line /snap/go/10455/src/net/dnsclient_unix.go:711
					_go_fuzz_dep_.CoverTab[528073]++
//line /snap/go/10455/src/net/dnsclient_unix.go:711
					_go_fuzz_dep_.CoverTab[5631]++
											a, err := result.p.AResource()
											if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:713
						_go_fuzz_dep_.CoverTab[528077]++
//line /snap/go/10455/src/net/dnsclient_unix.go:713
						_go_fuzz_dep_.CoverTab[5639]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /snap/go/10455/src/net/dnsclient_unix.go:719
						// _ = "end of CoverTab[5639]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:720
						_go_fuzz_dep_.CoverTab[528078]++
//line /snap/go/10455/src/net/dnsclient_unix.go:720
						_go_fuzz_dep_.CoverTab[5640]++
//line /snap/go/10455/src/net/dnsclient_unix.go:720
						// _ = "end of CoverTab[5640]"
//line /snap/go/10455/src/net/dnsclient_unix.go:720
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:720
					// _ = "end of CoverTab[5631]"
//line /snap/go/10455/src/net/dnsclient_unix.go:720
					_go_fuzz_dep_.CoverTab[5632]++
											addrs = append(addrs, IPAddr{IP: IP(a.A[:])})
											if cname.Length == 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[5641]++
//line /snap/go/10455/src/net/dnsclient_unix.go:722
						return h.Name.Length != 0
//line /snap/go/10455/src/net/dnsclient_unix.go:722
						// _ = "end of CoverTab[5641]"
//line /snap/go/10455/src/net/dnsclient_unix.go:722
					}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[528079]++
//line /snap/go/10455/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[5642]++
												cname = h.Name
//line /snap/go/10455/src/net/dnsclient_unix.go:723
						// _ = "end of CoverTab[5642]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:724
						_go_fuzz_dep_.CoverTab[528080]++
//line /snap/go/10455/src/net/dnsclient_unix.go:724
						_go_fuzz_dep_.CoverTab[5643]++
//line /snap/go/10455/src/net/dnsclient_unix.go:724
						// _ = "end of CoverTab[5643]"
//line /snap/go/10455/src/net/dnsclient_unix.go:724
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:724
					// _ = "end of CoverTab[5632]"

				case dnsmessage.TypeAAAA:
//line /snap/go/10455/src/net/dnsclient_unix.go:726
					_go_fuzz_dep_.CoverTab[528074]++
//line /snap/go/10455/src/net/dnsclient_unix.go:726
					_go_fuzz_dep_.CoverTab[5633]++
											aaaa, err := result.p.AAAAResource()
											if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:728
						_go_fuzz_dep_.CoverTab[528081]++
//line /snap/go/10455/src/net/dnsclient_unix.go:728
						_go_fuzz_dep_.CoverTab[5644]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /snap/go/10455/src/net/dnsclient_unix.go:734
						// _ = "end of CoverTab[5644]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:735
						_go_fuzz_dep_.CoverTab[528082]++
//line /snap/go/10455/src/net/dnsclient_unix.go:735
						_go_fuzz_dep_.CoverTab[5645]++
//line /snap/go/10455/src/net/dnsclient_unix.go:735
						// _ = "end of CoverTab[5645]"
//line /snap/go/10455/src/net/dnsclient_unix.go:735
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:735
					// _ = "end of CoverTab[5633]"
//line /snap/go/10455/src/net/dnsclient_unix.go:735
					_go_fuzz_dep_.CoverTab[5634]++
											addrs = append(addrs, IPAddr{IP: IP(aaaa.AAAA[:])})
											if cname.Length == 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[5646]++
//line /snap/go/10455/src/net/dnsclient_unix.go:737
						return h.Name.Length != 0
//line /snap/go/10455/src/net/dnsclient_unix.go:737
						// _ = "end of CoverTab[5646]"
//line /snap/go/10455/src/net/dnsclient_unix.go:737
					}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[528083]++
//line /snap/go/10455/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[5647]++
												cname = h.Name
//line /snap/go/10455/src/net/dnsclient_unix.go:738
						// _ = "end of CoverTab[5647]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:739
						_go_fuzz_dep_.CoverTab[528084]++
//line /snap/go/10455/src/net/dnsclient_unix.go:739
						_go_fuzz_dep_.CoverTab[5648]++
//line /snap/go/10455/src/net/dnsclient_unix.go:739
						// _ = "end of CoverTab[5648]"
//line /snap/go/10455/src/net/dnsclient_unix.go:739
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:739
					// _ = "end of CoverTab[5634]"

				case dnsmessage.TypeCNAME:
//line /snap/go/10455/src/net/dnsclient_unix.go:741
					_go_fuzz_dep_.CoverTab[528075]++
//line /snap/go/10455/src/net/dnsclient_unix.go:741
					_go_fuzz_dep_.CoverTab[5635]++
											c, err := result.p.CNAMEResource()
											if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:743
						_go_fuzz_dep_.CoverTab[528085]++
//line /snap/go/10455/src/net/dnsclient_unix.go:743
						_go_fuzz_dep_.CoverTab[5649]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /snap/go/10455/src/net/dnsclient_unix.go:749
						// _ = "end of CoverTab[5649]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:750
						_go_fuzz_dep_.CoverTab[528086]++
//line /snap/go/10455/src/net/dnsclient_unix.go:750
						_go_fuzz_dep_.CoverTab[5650]++
//line /snap/go/10455/src/net/dnsclient_unix.go:750
						// _ = "end of CoverTab[5650]"
//line /snap/go/10455/src/net/dnsclient_unix.go:750
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:750
					// _ = "end of CoverTab[5635]"
//line /snap/go/10455/src/net/dnsclient_unix.go:750
					_go_fuzz_dep_.CoverTab[5636]++
											if cname.Length == 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[5651]++
//line /snap/go/10455/src/net/dnsclient_unix.go:751
						return c.CNAME.Length > 0
//line /snap/go/10455/src/net/dnsclient_unix.go:751
						// _ = "end of CoverTab[5651]"
//line /snap/go/10455/src/net/dnsclient_unix.go:751
					}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[528087]++
//line /snap/go/10455/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[5652]++
												cname = c.CNAME
//line /snap/go/10455/src/net/dnsclient_unix.go:752
						// _ = "end of CoverTab[5652]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:753
						_go_fuzz_dep_.CoverTab[528088]++
//line /snap/go/10455/src/net/dnsclient_unix.go:753
						_go_fuzz_dep_.CoverTab[5653]++
//line /snap/go/10455/src/net/dnsclient_unix.go:753
						// _ = "end of CoverTab[5653]"
//line /snap/go/10455/src/net/dnsclient_unix.go:753
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:753
					// _ = "end of CoverTab[5636]"

				default:
//line /snap/go/10455/src/net/dnsclient_unix.go:755
					_go_fuzz_dep_.CoverTab[528076]++
//line /snap/go/10455/src/net/dnsclient_unix.go:755
					_go_fuzz_dep_.CoverTab[5637]++
											if err := result.p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:756
						_go_fuzz_dep_.CoverTab[528089]++
//line /snap/go/10455/src/net/dnsclient_unix.go:756
						_go_fuzz_dep_.CoverTab[5654]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /snap/go/10455/src/net/dnsclient_unix.go:762
						// _ = "end of CoverTab[5654]"
					} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:763
						_go_fuzz_dep_.CoverTab[528090]++
//line /snap/go/10455/src/net/dnsclient_unix.go:763
						_go_fuzz_dep_.CoverTab[5655]++
//line /snap/go/10455/src/net/dnsclient_unix.go:763
						// _ = "end of CoverTab[5655]"
//line /snap/go/10455/src/net/dnsclient_unix.go:763
					}
//line /snap/go/10455/src/net/dnsclient_unix.go:763
					// _ = "end of CoverTab[5637]"
//line /snap/go/10455/src/net/dnsclient_unix.go:763
					_go_fuzz_dep_.CoverTab[5638]++
											continue
//line /snap/go/10455/src/net/dnsclient_unix.go:764
					// _ = "end of CoverTab[5638]"
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:765
				// _ = "end of CoverTab[5625]"
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:766
			// _ = "end of CoverTab[5612]"
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:767
		if _go_fuzz_dep_.CoverTab[786671] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:767
			_go_fuzz_dep_.CoverTab[528179]++
//line /snap/go/10455/src/net/dnsclient_unix.go:767
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:767
			_go_fuzz_dep_.CoverTab[528180]++
//line /snap/go/10455/src/net/dnsclient_unix.go:767
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:767
		// _ = "end of CoverTab[5607]"
//line /snap/go/10455/src/net/dnsclient_unix.go:767
		_go_fuzz_dep_.CoverTab[5608]++
								if hitStrictError {
//line /snap/go/10455/src/net/dnsclient_unix.go:768
			_go_fuzz_dep_.CoverTab[528091]++
//line /snap/go/10455/src/net/dnsclient_unix.go:768
			_go_fuzz_dep_.CoverTab[5656]++

//line /snap/go/10455/src/net/dnsclient_unix.go:772
			addrs = nil
									break
//line /snap/go/10455/src/net/dnsclient_unix.go:773
			// _ = "end of CoverTab[5656]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:774
			_go_fuzz_dep_.CoverTab[528092]++
//line /snap/go/10455/src/net/dnsclient_unix.go:774
			_go_fuzz_dep_.CoverTab[5657]++
//line /snap/go/10455/src/net/dnsclient_unix.go:774
			// _ = "end of CoverTab[5657]"
//line /snap/go/10455/src/net/dnsclient_unix.go:774
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:774
		// _ = "end of CoverTab[5608]"
//line /snap/go/10455/src/net/dnsclient_unix.go:774
		_go_fuzz_dep_.CoverTab[5609]++
								if len(addrs) > 0 || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[5658]++
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			return network == "CNAME" && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:775
				_go_fuzz_dep_.CoverTab[5659]++
//line /snap/go/10455/src/net/dnsclient_unix.go:775
				return cname.Length > 0
//line /snap/go/10455/src/net/dnsclient_unix.go:775
				// _ = "end of CoverTab[5659]"
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			}()
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			// _ = "end of CoverTab[5658]"
//line /snap/go/10455/src/net/dnsclient_unix.go:775
		}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[528093]++
//line /snap/go/10455/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[5660]++
									break
//line /snap/go/10455/src/net/dnsclient_unix.go:776
			// _ = "end of CoverTab[5660]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:777
			_go_fuzz_dep_.CoverTab[528094]++
//line /snap/go/10455/src/net/dnsclient_unix.go:777
			_go_fuzz_dep_.CoverTab[5661]++
//line /snap/go/10455/src/net/dnsclient_unix.go:777
			// _ = "end of CoverTab[5661]"
//line /snap/go/10455/src/net/dnsclient_unix.go:777
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:777
		// _ = "end of CoverTab[5609]"
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:778
	if _go_fuzz_dep_.CoverTab[786669] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:778
		_go_fuzz_dep_.CoverTab[528171]++
//line /snap/go/10455/src/net/dnsclient_unix.go:778
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:778
		_go_fuzz_dep_.CoverTab[528172]++
//line /snap/go/10455/src/net/dnsclient_unix.go:778
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:778
	// _ = "end of CoverTab[5572]"
//line /snap/go/10455/src/net/dnsclient_unix.go:778
	_go_fuzz_dep_.CoverTab[5573]++
							if lastErr, ok := lastErr.(*DNSError); ok {
//line /snap/go/10455/src/net/dnsclient_unix.go:779
		_go_fuzz_dep_.CoverTab[528095]++
//line /snap/go/10455/src/net/dnsclient_unix.go:779
		_go_fuzz_dep_.CoverTab[5662]++

//line /snap/go/10455/src/net/dnsclient_unix.go:783
		lastErr.Name = name
//line /snap/go/10455/src/net/dnsclient_unix.go:783
		// _ = "end of CoverTab[5662]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:784
		_go_fuzz_dep_.CoverTab[528096]++
//line /snap/go/10455/src/net/dnsclient_unix.go:784
		_go_fuzz_dep_.CoverTab[5663]++
//line /snap/go/10455/src/net/dnsclient_unix.go:784
		// _ = "end of CoverTab[5663]"
//line /snap/go/10455/src/net/dnsclient_unix.go:784
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:784
	// _ = "end of CoverTab[5573]"
//line /snap/go/10455/src/net/dnsclient_unix.go:784
	_go_fuzz_dep_.CoverTab[5574]++
							sortByRFC6724(addrs)
							if len(addrs) == 0 && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[5664]++
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		return !(network == "CNAME" && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:786
			_go_fuzz_dep_.CoverTab[5665]++
//line /snap/go/10455/src/net/dnsclient_unix.go:786
			return cname.Length > 0
//line /snap/go/10455/src/net/dnsclient_unix.go:786
			// _ = "end of CoverTab[5665]"
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		}())
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		// _ = "end of CoverTab[5664]"
//line /snap/go/10455/src/net/dnsclient_unix.go:786
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[528097]++
//line /snap/go/10455/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[5666]++
								if order == hostLookupDNSFiles {
//line /snap/go/10455/src/net/dnsclient_unix.go:787
			_go_fuzz_dep_.CoverTab[528099]++
//line /snap/go/10455/src/net/dnsclient_unix.go:787
			_go_fuzz_dep_.CoverTab[5668]++
									var canonical string
									addrs, canonical = goLookupIPFiles(name)
									if len(addrs) > 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:790
				_go_fuzz_dep_.CoverTab[528101]++
//line /snap/go/10455/src/net/dnsclient_unix.go:790
				_go_fuzz_dep_.CoverTab[5669]++
										var err error
										cname, err = dnsmessage.NewName(canonical)
										if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:793
					_go_fuzz_dep_.CoverTab[528103]++
//line /snap/go/10455/src/net/dnsclient_unix.go:793
					_go_fuzz_dep_.CoverTab[5671]++
											return nil, dnsmessage.Name{}, err
//line /snap/go/10455/src/net/dnsclient_unix.go:794
					// _ = "end of CoverTab[5671]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:795
					_go_fuzz_dep_.CoverTab[528104]++
//line /snap/go/10455/src/net/dnsclient_unix.go:795
					_go_fuzz_dep_.CoverTab[5672]++
//line /snap/go/10455/src/net/dnsclient_unix.go:795
					// _ = "end of CoverTab[5672]"
//line /snap/go/10455/src/net/dnsclient_unix.go:795
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:795
				// _ = "end of CoverTab[5669]"
//line /snap/go/10455/src/net/dnsclient_unix.go:795
				_go_fuzz_dep_.CoverTab[5670]++
										return addrs, cname, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:796
				// _ = "end of CoverTab[5670]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:797
				_go_fuzz_dep_.CoverTab[528102]++
//line /snap/go/10455/src/net/dnsclient_unix.go:797
				_go_fuzz_dep_.CoverTab[5673]++
//line /snap/go/10455/src/net/dnsclient_unix.go:797
				// _ = "end of CoverTab[5673]"
//line /snap/go/10455/src/net/dnsclient_unix.go:797
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:797
			// _ = "end of CoverTab[5668]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:798
			_go_fuzz_dep_.CoverTab[528100]++
//line /snap/go/10455/src/net/dnsclient_unix.go:798
			_go_fuzz_dep_.CoverTab[5674]++
//line /snap/go/10455/src/net/dnsclient_unix.go:798
			// _ = "end of CoverTab[5674]"
//line /snap/go/10455/src/net/dnsclient_unix.go:798
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:798
		// _ = "end of CoverTab[5666]"
//line /snap/go/10455/src/net/dnsclient_unix.go:798
		_go_fuzz_dep_.CoverTab[5667]++
								if lastErr != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:799
			_go_fuzz_dep_.CoverTab[528105]++
//line /snap/go/10455/src/net/dnsclient_unix.go:799
			_go_fuzz_dep_.CoverTab[5675]++
									return nil, dnsmessage.Name{}, lastErr
//line /snap/go/10455/src/net/dnsclient_unix.go:800
			// _ = "end of CoverTab[5675]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:801
			_go_fuzz_dep_.CoverTab[528106]++
//line /snap/go/10455/src/net/dnsclient_unix.go:801
			_go_fuzz_dep_.CoverTab[5676]++
//line /snap/go/10455/src/net/dnsclient_unix.go:801
			// _ = "end of CoverTab[5676]"
//line /snap/go/10455/src/net/dnsclient_unix.go:801
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:801
		// _ = "end of CoverTab[5667]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:802
		_go_fuzz_dep_.CoverTab[528098]++
//line /snap/go/10455/src/net/dnsclient_unix.go:802
		_go_fuzz_dep_.CoverTab[5677]++
//line /snap/go/10455/src/net/dnsclient_unix.go:802
		// _ = "end of CoverTab[5677]"
//line /snap/go/10455/src/net/dnsclient_unix.go:802
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:802
	// _ = "end of CoverTab[5574]"
//line /snap/go/10455/src/net/dnsclient_unix.go:802
	_go_fuzz_dep_.CoverTab[5575]++
							return addrs, cname, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:803
	// _ = "end of CoverTab[5575]"
}

// goLookupCNAME is the native Go (non-cgo) implementation of LookupCNAME.
func (r *Resolver) goLookupCNAME(ctx context.Context, host string, order hostLookupOrder, conf *dnsConfig) (string, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:807
	_go_fuzz_dep_.CoverTab[5678]++
							_, cname, err := r.goLookupIPCNAMEOrder(ctx, "CNAME", host, order, conf)
							return cname.String(), err
//line /snap/go/10455/src/net/dnsclient_unix.go:809
	// _ = "end of CoverTab[5678]"
}

// goLookupPTR is the native Go implementation of LookupAddr.
func (r *Resolver) goLookupPTR(ctx context.Context, addr string, order hostLookupOrder, conf *dnsConfig) ([]string, error) {
//line /snap/go/10455/src/net/dnsclient_unix.go:813
	_go_fuzz_dep_.CoverTab[5679]++
							if order == hostLookupFiles || func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:814
		_go_fuzz_dep_.CoverTab[5684]++
//line /snap/go/10455/src/net/dnsclient_unix.go:814
		return order == hostLookupFilesDNS
//line /snap/go/10455/src/net/dnsclient_unix.go:814
		// _ = "end of CoverTab[5684]"
//line /snap/go/10455/src/net/dnsclient_unix.go:814
	}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:814
		_go_fuzz_dep_.CoverTab[528107]++
//line /snap/go/10455/src/net/dnsclient_unix.go:814
		_go_fuzz_dep_.CoverTab[5685]++
								names := lookupStaticAddr(addr)
								if len(names) > 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:816
			_go_fuzz_dep_.CoverTab[528109]++
//line /snap/go/10455/src/net/dnsclient_unix.go:816
			_go_fuzz_dep_.CoverTab[5687]++
									return names, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:817
			// _ = "end of CoverTab[5687]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:818
			_go_fuzz_dep_.CoverTab[528110]++
//line /snap/go/10455/src/net/dnsclient_unix.go:818
			_go_fuzz_dep_.CoverTab[5688]++
//line /snap/go/10455/src/net/dnsclient_unix.go:818
			// _ = "end of CoverTab[5688]"
//line /snap/go/10455/src/net/dnsclient_unix.go:818
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:818
		// _ = "end of CoverTab[5685]"
//line /snap/go/10455/src/net/dnsclient_unix.go:818
		_go_fuzz_dep_.CoverTab[5686]++

								if order == hostLookupFiles {
//line /snap/go/10455/src/net/dnsclient_unix.go:820
			_go_fuzz_dep_.CoverTab[528111]++
//line /snap/go/10455/src/net/dnsclient_unix.go:820
			_go_fuzz_dep_.CoverTab[5689]++
									return nil, &DNSError{Err: errNoSuchHost.Error(), Name: addr, IsNotFound: true}
//line /snap/go/10455/src/net/dnsclient_unix.go:821
			// _ = "end of CoverTab[5689]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:822
			_go_fuzz_dep_.CoverTab[528112]++
//line /snap/go/10455/src/net/dnsclient_unix.go:822
			_go_fuzz_dep_.CoverTab[5690]++
//line /snap/go/10455/src/net/dnsclient_unix.go:822
			// _ = "end of CoverTab[5690]"
//line /snap/go/10455/src/net/dnsclient_unix.go:822
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:822
		// _ = "end of CoverTab[5686]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:823
		_go_fuzz_dep_.CoverTab[528108]++
//line /snap/go/10455/src/net/dnsclient_unix.go:823
		_go_fuzz_dep_.CoverTab[5691]++
//line /snap/go/10455/src/net/dnsclient_unix.go:823
		// _ = "end of CoverTab[5691]"
//line /snap/go/10455/src/net/dnsclient_unix.go:823
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:823
	// _ = "end of CoverTab[5679]"
//line /snap/go/10455/src/net/dnsclient_unix.go:823
	_go_fuzz_dep_.CoverTab[5680]++

							arpa, err := reverseaddr(addr)
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:826
		_go_fuzz_dep_.CoverTab[528113]++
//line /snap/go/10455/src/net/dnsclient_unix.go:826
		_go_fuzz_dep_.CoverTab[5692]++
								return nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:827
		// _ = "end of CoverTab[5692]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:828
		_go_fuzz_dep_.CoverTab[528114]++
//line /snap/go/10455/src/net/dnsclient_unix.go:828
		_go_fuzz_dep_.CoverTab[5693]++
//line /snap/go/10455/src/net/dnsclient_unix.go:828
		// _ = "end of CoverTab[5693]"
//line /snap/go/10455/src/net/dnsclient_unix.go:828
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:828
	// _ = "end of CoverTab[5680]"
//line /snap/go/10455/src/net/dnsclient_unix.go:828
	_go_fuzz_dep_.CoverTab[5681]++
							p, server, err := r.lookup(ctx, arpa, dnsmessage.TypePTR, conf)
							if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:830
		_go_fuzz_dep_.CoverTab[528115]++
//line /snap/go/10455/src/net/dnsclient_unix.go:830
		_go_fuzz_dep_.CoverTab[5694]++
								var dnsErr *DNSError
								if errors.As(err, &dnsErr) && func() bool {
//line /snap/go/10455/src/net/dnsclient_unix.go:832
			_go_fuzz_dep_.CoverTab[5696]++
//line /snap/go/10455/src/net/dnsclient_unix.go:832
			return dnsErr.IsNotFound
//line /snap/go/10455/src/net/dnsclient_unix.go:832
			// _ = "end of CoverTab[5696]"
//line /snap/go/10455/src/net/dnsclient_unix.go:832
		}() {
//line /snap/go/10455/src/net/dnsclient_unix.go:832
			_go_fuzz_dep_.CoverTab[528117]++
//line /snap/go/10455/src/net/dnsclient_unix.go:832
			_go_fuzz_dep_.CoverTab[5697]++
									if order == hostLookupDNSFiles {
//line /snap/go/10455/src/net/dnsclient_unix.go:833
				_go_fuzz_dep_.CoverTab[528119]++
//line /snap/go/10455/src/net/dnsclient_unix.go:833
				_go_fuzz_dep_.CoverTab[5698]++
										names := lookupStaticAddr(addr)
										if len(names) > 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:835
					_go_fuzz_dep_.CoverTab[528121]++
//line /snap/go/10455/src/net/dnsclient_unix.go:835
					_go_fuzz_dep_.CoverTab[5699]++
											return names, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:836
					// _ = "end of CoverTab[5699]"
				} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:837
					_go_fuzz_dep_.CoverTab[528122]++
//line /snap/go/10455/src/net/dnsclient_unix.go:837
					_go_fuzz_dep_.CoverTab[5700]++
//line /snap/go/10455/src/net/dnsclient_unix.go:837
					// _ = "end of CoverTab[5700]"
//line /snap/go/10455/src/net/dnsclient_unix.go:837
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:837
				// _ = "end of CoverTab[5698]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:838
				_go_fuzz_dep_.CoverTab[528120]++
//line /snap/go/10455/src/net/dnsclient_unix.go:838
				_go_fuzz_dep_.CoverTab[5701]++
//line /snap/go/10455/src/net/dnsclient_unix.go:838
				// _ = "end of CoverTab[5701]"
//line /snap/go/10455/src/net/dnsclient_unix.go:838
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:838
			// _ = "end of CoverTab[5697]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:839
			_go_fuzz_dep_.CoverTab[528118]++
//line /snap/go/10455/src/net/dnsclient_unix.go:839
			_go_fuzz_dep_.CoverTab[5702]++
//line /snap/go/10455/src/net/dnsclient_unix.go:839
			// _ = "end of CoverTab[5702]"
//line /snap/go/10455/src/net/dnsclient_unix.go:839
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:839
		// _ = "end of CoverTab[5694]"
//line /snap/go/10455/src/net/dnsclient_unix.go:839
		_go_fuzz_dep_.CoverTab[5695]++
								return nil, err
//line /snap/go/10455/src/net/dnsclient_unix.go:840
		// _ = "end of CoverTab[5695]"
	} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:841
		_go_fuzz_dep_.CoverTab[528116]++
//line /snap/go/10455/src/net/dnsclient_unix.go:841
		_go_fuzz_dep_.CoverTab[5703]++
//line /snap/go/10455/src/net/dnsclient_unix.go:841
		// _ = "end of CoverTab[5703]"
//line /snap/go/10455/src/net/dnsclient_unix.go:841
	}
//line /snap/go/10455/src/net/dnsclient_unix.go:841
	// _ = "end of CoverTab[5681]"
//line /snap/go/10455/src/net/dnsclient_unix.go:841
	_go_fuzz_dep_.CoverTab[5682]++
							var ptrs []string
//line /snap/go/10455/src/net/dnsclient_unix.go:842
	_go_fuzz_dep_.CoverTab[786672] = 0
							for {
//line /snap/go/10455/src/net/dnsclient_unix.go:843
		if _go_fuzz_dep_.CoverTab[786672] == 0 {
//line /snap/go/10455/src/net/dnsclient_unix.go:843
			_go_fuzz_dep_.CoverTab[528181]++
//line /snap/go/10455/src/net/dnsclient_unix.go:843
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:843
			_go_fuzz_dep_.CoverTab[528182]++
//line /snap/go/10455/src/net/dnsclient_unix.go:843
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:843
		_go_fuzz_dep_.CoverTab[786672] = 1
//line /snap/go/10455/src/net/dnsclient_unix.go:843
		_go_fuzz_dep_.CoverTab[5704]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/dnsclient_unix.go:845
			_go_fuzz_dep_.CoverTab[528123]++
//line /snap/go/10455/src/net/dnsclient_unix.go:845
			_go_fuzz_dep_.CoverTab[5709]++
									break
//line /snap/go/10455/src/net/dnsclient_unix.go:846
			// _ = "end of CoverTab[5709]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:847
			_go_fuzz_dep_.CoverTab[528124]++
//line /snap/go/10455/src/net/dnsclient_unix.go:847
			_go_fuzz_dep_.CoverTab[5710]++
//line /snap/go/10455/src/net/dnsclient_unix.go:847
			// _ = "end of CoverTab[5710]"
//line /snap/go/10455/src/net/dnsclient_unix.go:847
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:847
		// _ = "end of CoverTab[5704]"
//line /snap/go/10455/src/net/dnsclient_unix.go:847
		_go_fuzz_dep_.CoverTab[5705]++
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:848
			_go_fuzz_dep_.CoverTab[528125]++
//line /snap/go/10455/src/net/dnsclient_unix.go:848
			_go_fuzz_dep_.CoverTab[5711]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:853
			// _ = "end of CoverTab[5711]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:854
			_go_fuzz_dep_.CoverTab[528126]++
//line /snap/go/10455/src/net/dnsclient_unix.go:854
			_go_fuzz_dep_.CoverTab[5712]++
//line /snap/go/10455/src/net/dnsclient_unix.go:854
			// _ = "end of CoverTab[5712]"
//line /snap/go/10455/src/net/dnsclient_unix.go:854
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:854
		// _ = "end of CoverTab[5705]"
//line /snap/go/10455/src/net/dnsclient_unix.go:854
		_go_fuzz_dep_.CoverTab[5706]++
								if h.Type != dnsmessage.TypePTR {
//line /snap/go/10455/src/net/dnsclient_unix.go:855
			_go_fuzz_dep_.CoverTab[528127]++
//line /snap/go/10455/src/net/dnsclient_unix.go:855
			_go_fuzz_dep_.CoverTab[5713]++
									err := p.SkipAnswer()
									if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:857
				_go_fuzz_dep_.CoverTab[528129]++
//line /snap/go/10455/src/net/dnsclient_unix.go:857
				_go_fuzz_dep_.CoverTab[5715]++
										return nil, &DNSError{
					Err:	"cannot marshal DNS message",
					Name:	addr,
					Server:	server,
				}
//line /snap/go/10455/src/net/dnsclient_unix.go:862
				// _ = "end of CoverTab[5715]"
			} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:863
				_go_fuzz_dep_.CoverTab[528130]++
//line /snap/go/10455/src/net/dnsclient_unix.go:863
				_go_fuzz_dep_.CoverTab[5716]++
//line /snap/go/10455/src/net/dnsclient_unix.go:863
				// _ = "end of CoverTab[5716]"
//line /snap/go/10455/src/net/dnsclient_unix.go:863
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:863
			// _ = "end of CoverTab[5713]"
//line /snap/go/10455/src/net/dnsclient_unix.go:863
			_go_fuzz_dep_.CoverTab[5714]++
									continue
//line /snap/go/10455/src/net/dnsclient_unix.go:864
			// _ = "end of CoverTab[5714]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:865
			_go_fuzz_dep_.CoverTab[528128]++
//line /snap/go/10455/src/net/dnsclient_unix.go:865
			_go_fuzz_dep_.CoverTab[5717]++
//line /snap/go/10455/src/net/dnsclient_unix.go:865
			// _ = "end of CoverTab[5717]"
//line /snap/go/10455/src/net/dnsclient_unix.go:865
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:865
		// _ = "end of CoverTab[5706]"
//line /snap/go/10455/src/net/dnsclient_unix.go:865
		_go_fuzz_dep_.CoverTab[5707]++
								ptr, err := p.PTRResource()
								if err != nil {
//line /snap/go/10455/src/net/dnsclient_unix.go:867
			_go_fuzz_dep_.CoverTab[528131]++
//line /snap/go/10455/src/net/dnsclient_unix.go:867
			_go_fuzz_dep_.CoverTab[5718]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /snap/go/10455/src/net/dnsclient_unix.go:872
			// _ = "end of CoverTab[5718]"
		} else {
//line /snap/go/10455/src/net/dnsclient_unix.go:873
			_go_fuzz_dep_.CoverTab[528132]++
//line /snap/go/10455/src/net/dnsclient_unix.go:873
			_go_fuzz_dep_.CoverTab[5719]++
//line /snap/go/10455/src/net/dnsclient_unix.go:873
			// _ = "end of CoverTab[5719]"
//line /snap/go/10455/src/net/dnsclient_unix.go:873
		}
//line /snap/go/10455/src/net/dnsclient_unix.go:873
		// _ = "end of CoverTab[5707]"
//line /snap/go/10455/src/net/dnsclient_unix.go:873
		_go_fuzz_dep_.CoverTab[5708]++
								ptrs = append(ptrs, ptr.PTR.String())
//line /snap/go/10455/src/net/dnsclient_unix.go:874
		// _ = "end of CoverTab[5708]"

	}
//line /snap/go/10455/src/net/dnsclient_unix.go:876
	// _ = "end of CoverTab[5682]"
//line /snap/go/10455/src/net/dnsclient_unix.go:876
	_go_fuzz_dep_.CoverTab[5683]++

							return ptrs, nil
//line /snap/go/10455/src/net/dnsclient_unix.go:878
	// _ = "end of CoverTab[5683]"
}

//line /snap/go/10455/src/net/dnsclient_unix.go:879
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/dnsclient_unix.go:879
var _ = _go_fuzz_dep_.CoverTab
