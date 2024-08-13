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

//line /usr/local/go/src/net/dnsclient_unix.go:15
package net

//line /usr/local/go/src/net/dnsclient_unix.go:15
import (
//line /usr/local/go/src/net/dnsclient_unix.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/dnsclient_unix.go:15
)
//line /usr/local/go/src/net/dnsclient_unix.go:15
import (
//line /usr/local/go/src/net/dnsclient_unix.go:15
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/dnsclient_unix.go:15
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
//line /usr/local/go/src/net/dnsclient_unix.go:55
	_go_fuzz_dep_.CoverTab[4941]++
							id = uint16(randInt())
							b := dnsmessage.NewBuilder(make([]byte, 2, 514), dnsmessage.Header{ID: id, RecursionDesired: true, AuthenticData: ad})
							if err := b.StartQuestions(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:58
		_go_fuzz_dep_.CoverTab[4948]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:59
		// _ = "end of CoverTab[4948]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:60
		_go_fuzz_dep_.CoverTab[4949]++
//line /usr/local/go/src/net/dnsclient_unix.go:60
		// _ = "end of CoverTab[4949]"
//line /usr/local/go/src/net/dnsclient_unix.go:60
	}
//line /usr/local/go/src/net/dnsclient_unix.go:60
	// _ = "end of CoverTab[4941]"
//line /usr/local/go/src/net/dnsclient_unix.go:60
	_go_fuzz_dep_.CoverTab[4942]++
							if err := b.Question(q); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:61
		_go_fuzz_dep_.CoverTab[4950]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:62
		// _ = "end of CoverTab[4950]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:63
		_go_fuzz_dep_.CoverTab[4951]++
//line /usr/local/go/src/net/dnsclient_unix.go:63
		// _ = "end of CoverTab[4951]"
//line /usr/local/go/src/net/dnsclient_unix.go:63
	}
//line /usr/local/go/src/net/dnsclient_unix.go:63
	// _ = "end of CoverTab[4942]"
//line /usr/local/go/src/net/dnsclient_unix.go:63
	_go_fuzz_dep_.CoverTab[4943]++

//line /usr/local/go/src/net/dnsclient_unix.go:66
	if err := b.StartAdditionals(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:66
		_go_fuzz_dep_.CoverTab[4952]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:67
		// _ = "end of CoverTab[4952]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:68
		_go_fuzz_dep_.CoverTab[4953]++
//line /usr/local/go/src/net/dnsclient_unix.go:68
		// _ = "end of CoverTab[4953]"
//line /usr/local/go/src/net/dnsclient_unix.go:68
	}
//line /usr/local/go/src/net/dnsclient_unix.go:68
	// _ = "end of CoverTab[4943]"
//line /usr/local/go/src/net/dnsclient_unix.go:68
	_go_fuzz_dep_.CoverTab[4944]++
							var rh dnsmessage.ResourceHeader
							if err := rh.SetEDNS0(maxDNSPacketSize, dnsmessage.RCodeSuccess, false); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:70
		_go_fuzz_dep_.CoverTab[4954]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:71
		// _ = "end of CoverTab[4954]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:72
		_go_fuzz_dep_.CoverTab[4955]++
//line /usr/local/go/src/net/dnsclient_unix.go:72
		// _ = "end of CoverTab[4955]"
//line /usr/local/go/src/net/dnsclient_unix.go:72
	}
//line /usr/local/go/src/net/dnsclient_unix.go:72
	// _ = "end of CoverTab[4944]"
//line /usr/local/go/src/net/dnsclient_unix.go:72
	_go_fuzz_dep_.CoverTab[4945]++
							if err := b.OPTResource(rh, dnsmessage.OPTResource{}); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:73
		_go_fuzz_dep_.CoverTab[4956]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:74
		// _ = "end of CoverTab[4956]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:75
		_go_fuzz_dep_.CoverTab[4957]++
//line /usr/local/go/src/net/dnsclient_unix.go:75
		// _ = "end of CoverTab[4957]"
//line /usr/local/go/src/net/dnsclient_unix.go:75
	}
//line /usr/local/go/src/net/dnsclient_unix.go:75
	// _ = "end of CoverTab[4945]"
//line /usr/local/go/src/net/dnsclient_unix.go:75
	_go_fuzz_dep_.CoverTab[4946]++

							tcpReq, err = b.Finish()
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:78
		_go_fuzz_dep_.CoverTab[4958]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:79
		// _ = "end of CoverTab[4958]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:80
		_go_fuzz_dep_.CoverTab[4959]++
//line /usr/local/go/src/net/dnsclient_unix.go:80
		// _ = "end of CoverTab[4959]"
//line /usr/local/go/src/net/dnsclient_unix.go:80
	}
//line /usr/local/go/src/net/dnsclient_unix.go:80
	// _ = "end of CoverTab[4946]"
//line /usr/local/go/src/net/dnsclient_unix.go:80
	_go_fuzz_dep_.CoverTab[4947]++
							udpReq = tcpReq[2:]
							l := len(tcpReq) - 2
							tcpReq[0] = byte(l >> 8)
							tcpReq[1] = byte(l)
							return id, udpReq, tcpReq, nil
//line /usr/local/go/src/net/dnsclient_unix.go:85
	// _ = "end of CoverTab[4947]"
}

func checkResponse(reqID uint16, reqQues dnsmessage.Question, respHdr dnsmessage.Header, respQues dnsmessage.Question) bool {
//line /usr/local/go/src/net/dnsclient_unix.go:88
	_go_fuzz_dep_.CoverTab[4960]++
							if !respHdr.Response {
//line /usr/local/go/src/net/dnsclient_unix.go:89
		_go_fuzz_dep_.CoverTab[4964]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:90
		// _ = "end of CoverTab[4964]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:91
		_go_fuzz_dep_.CoverTab[4965]++
//line /usr/local/go/src/net/dnsclient_unix.go:91
		// _ = "end of CoverTab[4965]"
//line /usr/local/go/src/net/dnsclient_unix.go:91
	}
//line /usr/local/go/src/net/dnsclient_unix.go:91
	// _ = "end of CoverTab[4960]"
//line /usr/local/go/src/net/dnsclient_unix.go:91
	_go_fuzz_dep_.CoverTab[4961]++
							if reqID != respHdr.ID {
//line /usr/local/go/src/net/dnsclient_unix.go:92
		_go_fuzz_dep_.CoverTab[4966]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:93
		// _ = "end of CoverTab[4966]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:94
		_go_fuzz_dep_.CoverTab[4967]++
//line /usr/local/go/src/net/dnsclient_unix.go:94
		// _ = "end of CoverTab[4967]"
//line /usr/local/go/src/net/dnsclient_unix.go:94
	}
//line /usr/local/go/src/net/dnsclient_unix.go:94
	// _ = "end of CoverTab[4961]"
//line /usr/local/go/src/net/dnsclient_unix.go:94
	_go_fuzz_dep_.CoverTab[4962]++
							if reqQues.Type != respQues.Type || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[4968]++
//line /usr/local/go/src/net/dnsclient_unix.go:95
		return reqQues.Class != respQues.Class
//line /usr/local/go/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[4968]"
//line /usr/local/go/src/net/dnsclient_unix.go:95
	}() || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[4969]++
//line /usr/local/go/src/net/dnsclient_unix.go:95
		return !equalASCIIName(reqQues.Name, respQues.Name)
//line /usr/local/go/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[4969]"
//line /usr/local/go/src/net/dnsclient_unix.go:95
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[4970]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:96
		// _ = "end of CoverTab[4970]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:97
		_go_fuzz_dep_.CoverTab[4971]++
//line /usr/local/go/src/net/dnsclient_unix.go:97
		// _ = "end of CoverTab[4971]"
//line /usr/local/go/src/net/dnsclient_unix.go:97
	}
//line /usr/local/go/src/net/dnsclient_unix.go:97
	// _ = "end of CoverTab[4962]"
//line /usr/local/go/src/net/dnsclient_unix.go:97
	_go_fuzz_dep_.CoverTab[4963]++
							return true
//line /usr/local/go/src/net/dnsclient_unix.go:98
	// _ = "end of CoverTab[4963]"
}

func dnsPacketRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:101
	_go_fuzz_dep_.CoverTab[4972]++
							if _, err := c.Write(b); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:102
		_go_fuzz_dep_.CoverTab[4974]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:103
		// _ = "end of CoverTab[4974]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:104
		_go_fuzz_dep_.CoverTab[4975]++
//line /usr/local/go/src/net/dnsclient_unix.go:104
		// _ = "end of CoverTab[4975]"
//line /usr/local/go/src/net/dnsclient_unix.go:104
	}
//line /usr/local/go/src/net/dnsclient_unix.go:104
	// _ = "end of CoverTab[4972]"
//line /usr/local/go/src/net/dnsclient_unix.go:104
	_go_fuzz_dep_.CoverTab[4973]++

							b = make([]byte, maxDNSPacketSize)
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:107
		_go_fuzz_dep_.CoverTab[4976]++
								n, err := c.Read(b)
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:109
			_go_fuzz_dep_.CoverTab[4980]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:110
			// _ = "end of CoverTab[4980]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:111
			_go_fuzz_dep_.CoverTab[4981]++
//line /usr/local/go/src/net/dnsclient_unix.go:111
			// _ = "end of CoverTab[4981]"
//line /usr/local/go/src/net/dnsclient_unix.go:111
		}
//line /usr/local/go/src/net/dnsclient_unix.go:111
		// _ = "end of CoverTab[4976]"
//line /usr/local/go/src/net/dnsclient_unix.go:111
		_go_fuzz_dep_.CoverTab[4977]++
								var p dnsmessage.Parser

//line /usr/local/go/src/net/dnsclient_unix.go:116
		h, err := p.Start(b[:n])
		if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:117
			_go_fuzz_dep_.CoverTab[4982]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:118
			// _ = "end of CoverTab[4982]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:119
			_go_fuzz_dep_.CoverTab[4983]++
//line /usr/local/go/src/net/dnsclient_unix.go:119
			// _ = "end of CoverTab[4983]"
//line /usr/local/go/src/net/dnsclient_unix.go:119
		}
//line /usr/local/go/src/net/dnsclient_unix.go:119
		// _ = "end of CoverTab[4977]"
//line /usr/local/go/src/net/dnsclient_unix.go:119
		_go_fuzz_dep_.CoverTab[4978]++
								q, err := p.Question()
								if err != nil || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[4984]++
//line /usr/local/go/src/net/dnsclient_unix.go:121
			return !checkResponse(id, query, h, q)
//line /usr/local/go/src/net/dnsclient_unix.go:121
			// _ = "end of CoverTab[4984]"
//line /usr/local/go/src/net/dnsclient_unix.go:121
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[4985]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:122
			// _ = "end of CoverTab[4985]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:123
			_go_fuzz_dep_.CoverTab[4986]++
//line /usr/local/go/src/net/dnsclient_unix.go:123
			// _ = "end of CoverTab[4986]"
//line /usr/local/go/src/net/dnsclient_unix.go:123
		}
//line /usr/local/go/src/net/dnsclient_unix.go:123
		// _ = "end of CoverTab[4978]"
//line /usr/local/go/src/net/dnsclient_unix.go:123
		_go_fuzz_dep_.CoverTab[4979]++
								return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:124
		// _ = "end of CoverTab[4979]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:125
	// _ = "end of CoverTab[4973]"
}

func dnsStreamRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:128
	_go_fuzz_dep_.CoverTab[4987]++
							if _, err := c.Write(b); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:129
		_go_fuzz_dep_.CoverTab[4995]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:130
		// _ = "end of CoverTab[4995]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:131
		_go_fuzz_dep_.CoverTab[4996]++
//line /usr/local/go/src/net/dnsclient_unix.go:131
		// _ = "end of CoverTab[4996]"
//line /usr/local/go/src/net/dnsclient_unix.go:131
	}
//line /usr/local/go/src/net/dnsclient_unix.go:131
	// _ = "end of CoverTab[4987]"
//line /usr/local/go/src/net/dnsclient_unix.go:131
	_go_fuzz_dep_.CoverTab[4988]++

							b = make([]byte, 1280)
							if _, err := io.ReadFull(c, b[:2]); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:134
		_go_fuzz_dep_.CoverTab[4997]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:135
		// _ = "end of CoverTab[4997]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:136
		_go_fuzz_dep_.CoverTab[4998]++
//line /usr/local/go/src/net/dnsclient_unix.go:136
		// _ = "end of CoverTab[4998]"
//line /usr/local/go/src/net/dnsclient_unix.go:136
	}
//line /usr/local/go/src/net/dnsclient_unix.go:136
	// _ = "end of CoverTab[4988]"
//line /usr/local/go/src/net/dnsclient_unix.go:136
	_go_fuzz_dep_.CoverTab[4989]++
							l := int(b[0])<<8 | int(b[1])
							if l > len(b) {
//line /usr/local/go/src/net/dnsclient_unix.go:138
		_go_fuzz_dep_.CoverTab[4999]++
								b = make([]byte, l)
//line /usr/local/go/src/net/dnsclient_unix.go:139
		// _ = "end of CoverTab[4999]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:140
		_go_fuzz_dep_.CoverTab[5000]++
//line /usr/local/go/src/net/dnsclient_unix.go:140
		// _ = "end of CoverTab[5000]"
//line /usr/local/go/src/net/dnsclient_unix.go:140
	}
//line /usr/local/go/src/net/dnsclient_unix.go:140
	// _ = "end of CoverTab[4989]"
//line /usr/local/go/src/net/dnsclient_unix.go:140
	_go_fuzz_dep_.CoverTab[4990]++
							n, err := io.ReadFull(c, b[:l])
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:142
		_go_fuzz_dep_.CoverTab[5001]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:143
		// _ = "end of CoverTab[5001]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:144
		_go_fuzz_dep_.CoverTab[5002]++
//line /usr/local/go/src/net/dnsclient_unix.go:144
		// _ = "end of CoverTab[5002]"
//line /usr/local/go/src/net/dnsclient_unix.go:144
	}
//line /usr/local/go/src/net/dnsclient_unix.go:144
	// _ = "end of CoverTab[4990]"
//line /usr/local/go/src/net/dnsclient_unix.go:144
	_go_fuzz_dep_.CoverTab[4991]++
							var p dnsmessage.Parser
							h, err := p.Start(b[:n])
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:147
		_go_fuzz_dep_.CoverTab[5003]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:148
		// _ = "end of CoverTab[5003]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:149
		_go_fuzz_dep_.CoverTab[5004]++
//line /usr/local/go/src/net/dnsclient_unix.go:149
		// _ = "end of CoverTab[5004]"
//line /usr/local/go/src/net/dnsclient_unix.go:149
	}
//line /usr/local/go/src/net/dnsclient_unix.go:149
	// _ = "end of CoverTab[4991]"
//line /usr/local/go/src/net/dnsclient_unix.go:149
	_go_fuzz_dep_.CoverTab[4992]++
							q, err := p.Question()
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:151
		_go_fuzz_dep_.CoverTab[5005]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:152
		// _ = "end of CoverTab[5005]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:153
		_go_fuzz_dep_.CoverTab[5006]++
//line /usr/local/go/src/net/dnsclient_unix.go:153
		// _ = "end of CoverTab[5006]"
//line /usr/local/go/src/net/dnsclient_unix.go:153
	}
//line /usr/local/go/src/net/dnsclient_unix.go:153
	// _ = "end of CoverTab[4992]"
//line /usr/local/go/src/net/dnsclient_unix.go:153
	_go_fuzz_dep_.CoverTab[4993]++
							if !checkResponse(id, query, h, q) {
//line /usr/local/go/src/net/dnsclient_unix.go:154
		_go_fuzz_dep_.CoverTab[5007]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /usr/local/go/src/net/dnsclient_unix.go:155
		// _ = "end of CoverTab[5007]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:156
		_go_fuzz_dep_.CoverTab[5008]++
//line /usr/local/go/src/net/dnsclient_unix.go:156
		// _ = "end of CoverTab[5008]"
//line /usr/local/go/src/net/dnsclient_unix.go:156
	}
//line /usr/local/go/src/net/dnsclient_unix.go:156
	// _ = "end of CoverTab[4993]"
//line /usr/local/go/src/net/dnsclient_unix.go:156
	_go_fuzz_dep_.CoverTab[4994]++
							return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:157
	// _ = "end of CoverTab[4994]"
}

// exchange sends a query on the connection and hopes for a response.
func (r *Resolver) exchange(ctx context.Context, server string, q dnsmessage.Question, timeout time.Duration, useTCP, ad bool) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:161
	_go_fuzz_dep_.CoverTab[5009]++
							q.Class = dnsmessage.ClassINET
							id, udpReq, tcpReq, err := newRequest(q, ad)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:164
		_go_fuzz_dep_.CoverTab[5013]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotMarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:165
		// _ = "end of CoverTab[5013]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:166
		_go_fuzz_dep_.CoverTab[5014]++
//line /usr/local/go/src/net/dnsclient_unix.go:166
		// _ = "end of CoverTab[5014]"
//line /usr/local/go/src/net/dnsclient_unix.go:166
	}
//line /usr/local/go/src/net/dnsclient_unix.go:166
	// _ = "end of CoverTab[5009]"
//line /usr/local/go/src/net/dnsclient_unix.go:166
	_go_fuzz_dep_.CoverTab[5010]++
							var networks []string
							if useTCP {
//line /usr/local/go/src/net/dnsclient_unix.go:168
		_go_fuzz_dep_.CoverTab[5015]++
								networks = []string{"tcp"}
//line /usr/local/go/src/net/dnsclient_unix.go:169
		// _ = "end of CoverTab[5015]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:170
		_go_fuzz_dep_.CoverTab[5016]++
								networks = []string{"udp", "tcp"}
//line /usr/local/go/src/net/dnsclient_unix.go:171
		// _ = "end of CoverTab[5016]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:172
	// _ = "end of CoverTab[5010]"
//line /usr/local/go/src/net/dnsclient_unix.go:172
	_go_fuzz_dep_.CoverTab[5011]++
							for _, network := range networks {
//line /usr/local/go/src/net/dnsclient_unix.go:173
		_go_fuzz_dep_.CoverTab[5017]++
								ctx, cancel := context.WithDeadline(ctx, time.Now().Add(timeout))
								defer cancel()

								c, err := r.dial(ctx, network, server)
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:178
			_go_fuzz_dep_.CoverTab[5024]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:179
			// _ = "end of CoverTab[5024]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:180
			_go_fuzz_dep_.CoverTab[5025]++
//line /usr/local/go/src/net/dnsclient_unix.go:180
			// _ = "end of CoverTab[5025]"
//line /usr/local/go/src/net/dnsclient_unix.go:180
		}
//line /usr/local/go/src/net/dnsclient_unix.go:180
		// _ = "end of CoverTab[5017]"
//line /usr/local/go/src/net/dnsclient_unix.go:180
		_go_fuzz_dep_.CoverTab[5018]++
								if d, ok := ctx.Deadline(); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[5026]++
//line /usr/local/go/src/net/dnsclient_unix.go:181
			return !d.IsZero()
//line /usr/local/go/src/net/dnsclient_unix.go:181
			// _ = "end of CoverTab[5026]"
//line /usr/local/go/src/net/dnsclient_unix.go:181
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[5027]++
									c.SetDeadline(d)
//line /usr/local/go/src/net/dnsclient_unix.go:182
			// _ = "end of CoverTab[5027]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:183
			_go_fuzz_dep_.CoverTab[5028]++
//line /usr/local/go/src/net/dnsclient_unix.go:183
			// _ = "end of CoverTab[5028]"
//line /usr/local/go/src/net/dnsclient_unix.go:183
		}
//line /usr/local/go/src/net/dnsclient_unix.go:183
		// _ = "end of CoverTab[5018]"
//line /usr/local/go/src/net/dnsclient_unix.go:183
		_go_fuzz_dep_.CoverTab[5019]++
								var p dnsmessage.Parser
								var h dnsmessage.Header
								if _, ok := c.(PacketConn); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:186
			_go_fuzz_dep_.CoverTab[5029]++
									p, h, err = dnsPacketRoundTrip(c, id, q, udpReq)
//line /usr/local/go/src/net/dnsclient_unix.go:187
			// _ = "end of CoverTab[5029]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:188
			_go_fuzz_dep_.CoverTab[5030]++
									p, h, err = dnsStreamRoundTrip(c, id, q, tcpReq)
//line /usr/local/go/src/net/dnsclient_unix.go:189
			// _ = "end of CoverTab[5030]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:190
		// _ = "end of CoverTab[5019]"
//line /usr/local/go/src/net/dnsclient_unix.go:190
		_go_fuzz_dep_.CoverTab[5020]++
								c.Close()
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:192
			_go_fuzz_dep_.CoverTab[5031]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, mapErr(err)
//line /usr/local/go/src/net/dnsclient_unix.go:193
			// _ = "end of CoverTab[5031]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:194
			_go_fuzz_dep_.CoverTab[5032]++
//line /usr/local/go/src/net/dnsclient_unix.go:194
			// _ = "end of CoverTab[5032]"
//line /usr/local/go/src/net/dnsclient_unix.go:194
		}
//line /usr/local/go/src/net/dnsclient_unix.go:194
		// _ = "end of CoverTab[5020]"
//line /usr/local/go/src/net/dnsclient_unix.go:194
		_go_fuzz_dep_.CoverTab[5021]++
								if err := p.SkipQuestion(); err != dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:195
			_go_fuzz_dep_.CoverTab[5033]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /usr/local/go/src/net/dnsclient_unix.go:196
			// _ = "end of CoverTab[5033]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:197
			_go_fuzz_dep_.CoverTab[5034]++
//line /usr/local/go/src/net/dnsclient_unix.go:197
			// _ = "end of CoverTab[5034]"
//line /usr/local/go/src/net/dnsclient_unix.go:197
		}
//line /usr/local/go/src/net/dnsclient_unix.go:197
		// _ = "end of CoverTab[5021]"
//line /usr/local/go/src/net/dnsclient_unix.go:197
		_go_fuzz_dep_.CoverTab[5022]++
								if h.Truncated {
//line /usr/local/go/src/net/dnsclient_unix.go:198
			_go_fuzz_dep_.CoverTab[5035]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:199
			// _ = "end of CoverTab[5035]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:200
			_go_fuzz_dep_.CoverTab[5036]++
//line /usr/local/go/src/net/dnsclient_unix.go:200
			// _ = "end of CoverTab[5036]"
//line /usr/local/go/src/net/dnsclient_unix.go:200
		}
//line /usr/local/go/src/net/dnsclient_unix.go:200
		// _ = "end of CoverTab[5022]"
//line /usr/local/go/src/net/dnsclient_unix.go:200
		_go_fuzz_dep_.CoverTab[5023]++
								return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:201
		// _ = "end of CoverTab[5023]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:202
	// _ = "end of CoverTab[5011]"
//line /usr/local/go/src/net/dnsclient_unix.go:202
	_go_fuzz_dep_.CoverTab[5012]++
							return dnsmessage.Parser{}, dnsmessage.Header{}, errNoAnswerFromDNSServer
//line /usr/local/go/src/net/dnsclient_unix.go:203
	// _ = "end of CoverTab[5012]"
}

// checkHeader performs basic sanity checks on the header.
func checkHeader(p *dnsmessage.Parser, h dnsmessage.Header) error {
//line /usr/local/go/src/net/dnsclient_unix.go:207
	_go_fuzz_dep_.CoverTab[5037]++
							if h.RCode == dnsmessage.RCodeNameError {
//line /usr/local/go/src/net/dnsclient_unix.go:208
		_go_fuzz_dep_.CoverTab[5042]++
								return errNoSuchHost
//line /usr/local/go/src/net/dnsclient_unix.go:209
		// _ = "end of CoverTab[5042]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:210
		_go_fuzz_dep_.CoverTab[5043]++
//line /usr/local/go/src/net/dnsclient_unix.go:210
		// _ = "end of CoverTab[5043]"
//line /usr/local/go/src/net/dnsclient_unix.go:210
	}
//line /usr/local/go/src/net/dnsclient_unix.go:210
	// _ = "end of CoverTab[5037]"
//line /usr/local/go/src/net/dnsclient_unix.go:210
	_go_fuzz_dep_.CoverTab[5038]++

							_, err := p.AnswerHeader()
							if err != nil && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[5044]++
//line /usr/local/go/src/net/dnsclient_unix.go:213
		return err != dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:213
		// _ = "end of CoverTab[5044]"
//line /usr/local/go/src/net/dnsclient_unix.go:213
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[5045]++
								return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:214
		// _ = "end of CoverTab[5045]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:215
		_go_fuzz_dep_.CoverTab[5046]++
//line /usr/local/go/src/net/dnsclient_unix.go:215
		// _ = "end of CoverTab[5046]"
//line /usr/local/go/src/net/dnsclient_unix.go:215
	}
//line /usr/local/go/src/net/dnsclient_unix.go:215
	// _ = "end of CoverTab[5038]"
//line /usr/local/go/src/net/dnsclient_unix.go:215
	_go_fuzz_dep_.CoverTab[5039]++

//line /usr/local/go/src/net/dnsclient_unix.go:219
	if h.RCode == dnsmessage.RCodeSuccess && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5047]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return !h.Authoritative
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5047]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5048]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return !h.RecursionAvailable
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5048]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5049]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return err == dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[5049]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[5050]++
								return errLameReferral
//line /usr/local/go/src/net/dnsclient_unix.go:220
		// _ = "end of CoverTab[5050]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:221
		_go_fuzz_dep_.CoverTab[5051]++
//line /usr/local/go/src/net/dnsclient_unix.go:221
		// _ = "end of CoverTab[5051]"
//line /usr/local/go/src/net/dnsclient_unix.go:221
	}
//line /usr/local/go/src/net/dnsclient_unix.go:221
	// _ = "end of CoverTab[5039]"
//line /usr/local/go/src/net/dnsclient_unix.go:221
	_go_fuzz_dep_.CoverTab[5040]++

							if h.RCode != dnsmessage.RCodeSuccess && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[5052]++
//line /usr/local/go/src/net/dnsclient_unix.go:223
		return h.RCode != dnsmessage.RCodeNameError
//line /usr/local/go/src/net/dnsclient_unix.go:223
		// _ = "end of CoverTab[5052]"
//line /usr/local/go/src/net/dnsclient_unix.go:223
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[5053]++

//line /usr/local/go/src/net/dnsclient_unix.go:229
		if h.RCode == dnsmessage.RCodeServerFailure {
//line /usr/local/go/src/net/dnsclient_unix.go:229
			_go_fuzz_dep_.CoverTab[5055]++
									return errServerTemporarilyMisbehaving
//line /usr/local/go/src/net/dnsclient_unix.go:230
			// _ = "end of CoverTab[5055]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:231
			_go_fuzz_dep_.CoverTab[5056]++
//line /usr/local/go/src/net/dnsclient_unix.go:231
			// _ = "end of CoverTab[5056]"
//line /usr/local/go/src/net/dnsclient_unix.go:231
		}
//line /usr/local/go/src/net/dnsclient_unix.go:231
		// _ = "end of CoverTab[5053]"
//line /usr/local/go/src/net/dnsclient_unix.go:231
		_go_fuzz_dep_.CoverTab[5054]++
								return errServerMisbehaving
//line /usr/local/go/src/net/dnsclient_unix.go:232
		// _ = "end of CoverTab[5054]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:233
		_go_fuzz_dep_.CoverTab[5057]++
//line /usr/local/go/src/net/dnsclient_unix.go:233
		// _ = "end of CoverTab[5057]"
//line /usr/local/go/src/net/dnsclient_unix.go:233
	}
//line /usr/local/go/src/net/dnsclient_unix.go:233
	// _ = "end of CoverTab[5040]"
//line /usr/local/go/src/net/dnsclient_unix.go:233
	_go_fuzz_dep_.CoverTab[5041]++

							return nil
//line /usr/local/go/src/net/dnsclient_unix.go:235
	// _ = "end of CoverTab[5041]"
}

func skipToAnswer(p *dnsmessage.Parser, qtype dnsmessage.Type) error {
//line /usr/local/go/src/net/dnsclient_unix.go:238
	_go_fuzz_dep_.CoverTab[5058]++
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:239
		_go_fuzz_dep_.CoverTab[5059]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:241
			_go_fuzz_dep_.CoverTab[5063]++
									return errNoSuchHost
//line /usr/local/go/src/net/dnsclient_unix.go:242
			// _ = "end of CoverTab[5063]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:243
			_go_fuzz_dep_.CoverTab[5064]++
//line /usr/local/go/src/net/dnsclient_unix.go:243
			// _ = "end of CoverTab[5064]"
//line /usr/local/go/src/net/dnsclient_unix.go:243
		}
//line /usr/local/go/src/net/dnsclient_unix.go:243
		// _ = "end of CoverTab[5059]"
//line /usr/local/go/src/net/dnsclient_unix.go:243
		_go_fuzz_dep_.CoverTab[5060]++
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:244
			_go_fuzz_dep_.CoverTab[5065]++
									return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:245
			// _ = "end of CoverTab[5065]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:246
			_go_fuzz_dep_.CoverTab[5066]++
//line /usr/local/go/src/net/dnsclient_unix.go:246
			// _ = "end of CoverTab[5066]"
//line /usr/local/go/src/net/dnsclient_unix.go:246
		}
//line /usr/local/go/src/net/dnsclient_unix.go:246
		// _ = "end of CoverTab[5060]"
//line /usr/local/go/src/net/dnsclient_unix.go:246
		_go_fuzz_dep_.CoverTab[5061]++
								if h.Type == qtype {
//line /usr/local/go/src/net/dnsclient_unix.go:247
			_go_fuzz_dep_.CoverTab[5067]++
									return nil
//line /usr/local/go/src/net/dnsclient_unix.go:248
			// _ = "end of CoverTab[5067]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:249
			_go_fuzz_dep_.CoverTab[5068]++
//line /usr/local/go/src/net/dnsclient_unix.go:249
			// _ = "end of CoverTab[5068]"
//line /usr/local/go/src/net/dnsclient_unix.go:249
		}
//line /usr/local/go/src/net/dnsclient_unix.go:249
		// _ = "end of CoverTab[5061]"
//line /usr/local/go/src/net/dnsclient_unix.go:249
		_go_fuzz_dep_.CoverTab[5062]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:250
			_go_fuzz_dep_.CoverTab[5069]++
									return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:251
			// _ = "end of CoverTab[5069]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:252
			_go_fuzz_dep_.CoverTab[5070]++
//line /usr/local/go/src/net/dnsclient_unix.go:252
			// _ = "end of CoverTab[5070]"
//line /usr/local/go/src/net/dnsclient_unix.go:252
		}
//line /usr/local/go/src/net/dnsclient_unix.go:252
		// _ = "end of CoverTab[5062]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:253
	// _ = "end of CoverTab[5058]"
}

// Do a lookup for a single name, which must be rooted
//line /usr/local/go/src/net/dnsclient_unix.go:256
// (otherwise answer will not find the answers).
//line /usr/local/go/src/net/dnsclient_unix.go:258
func (r *Resolver) tryOneName(ctx context.Context, cfg *dnsConfig, name string, qtype dnsmessage.Type) (dnsmessage.Parser, string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:258
	_go_fuzz_dep_.CoverTab[5071]++
							var lastErr error
							serverOffset := cfg.serverOffset()
							sLen := uint32(len(cfg.servers))

							n, err := dnsmessage.NewName(name)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:264
		_go_fuzz_dep_.CoverTab[5074]++
								return dnsmessage.Parser{}, "", errCannotMarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:265
		// _ = "end of CoverTab[5074]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:266
		_go_fuzz_dep_.CoverTab[5075]++
//line /usr/local/go/src/net/dnsclient_unix.go:266
		// _ = "end of CoverTab[5075]"
//line /usr/local/go/src/net/dnsclient_unix.go:266
	}
//line /usr/local/go/src/net/dnsclient_unix.go:266
	// _ = "end of CoverTab[5071]"
//line /usr/local/go/src/net/dnsclient_unix.go:266
	_go_fuzz_dep_.CoverTab[5072]++
							q := dnsmessage.Question{
		Name:	n,
		Type:	qtype,
		Class:	dnsmessage.ClassINET,
	}

	for i := 0; i < cfg.attempts; i++ {
//line /usr/local/go/src/net/dnsclient_unix.go:273
		_go_fuzz_dep_.CoverTab[5076]++
								for j := uint32(0); j < sLen; j++ {
//line /usr/local/go/src/net/dnsclient_unix.go:274
			_go_fuzz_dep_.CoverTab[5077]++
									server := cfg.servers[(serverOffset+j)%sLen]

									p, h, err := r.exchange(ctx, server, q, cfg.timeout, cfg.useTCP, cfg.trustAD)
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:278
				_go_fuzz_dep_.CoverTab[5081]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if nerr, ok := err.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[5084]++
//line /usr/local/go/src/net/dnsclient_unix.go:284
					return nerr.Timeout()
//line /usr/local/go/src/net/dnsclient_unix.go:284
					// _ = "end of CoverTab[5084]"
//line /usr/local/go/src/net/dnsclient_unix.go:284
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[5085]++
											dnsErr.IsTimeout = true
//line /usr/local/go/src/net/dnsclient_unix.go:285
					// _ = "end of CoverTab[5085]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:286
					_go_fuzz_dep_.CoverTab[5086]++
//line /usr/local/go/src/net/dnsclient_unix.go:286
					// _ = "end of CoverTab[5086]"
//line /usr/local/go/src/net/dnsclient_unix.go:286
				}
//line /usr/local/go/src/net/dnsclient_unix.go:286
				// _ = "end of CoverTab[5081]"
//line /usr/local/go/src/net/dnsclient_unix.go:286
				_go_fuzz_dep_.CoverTab[5082]++

//line /usr/local/go/src/net/dnsclient_unix.go:289
				if _, ok := err.(*OpError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:289
					_go_fuzz_dep_.CoverTab[5087]++
											dnsErr.IsTemporary = true
//line /usr/local/go/src/net/dnsclient_unix.go:290
					// _ = "end of CoverTab[5087]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:291
					_go_fuzz_dep_.CoverTab[5088]++
//line /usr/local/go/src/net/dnsclient_unix.go:291
					// _ = "end of CoverTab[5088]"
//line /usr/local/go/src/net/dnsclient_unix.go:291
				}
//line /usr/local/go/src/net/dnsclient_unix.go:291
				// _ = "end of CoverTab[5082]"
//line /usr/local/go/src/net/dnsclient_unix.go:291
				_go_fuzz_dep_.CoverTab[5083]++
										lastErr = dnsErr
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:293
				// _ = "end of CoverTab[5083]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:294
				_go_fuzz_dep_.CoverTab[5089]++
//line /usr/local/go/src/net/dnsclient_unix.go:294
				// _ = "end of CoverTab[5089]"
//line /usr/local/go/src/net/dnsclient_unix.go:294
			}
//line /usr/local/go/src/net/dnsclient_unix.go:294
			// _ = "end of CoverTab[5077]"
//line /usr/local/go/src/net/dnsclient_unix.go:294
			_go_fuzz_dep_.CoverTab[5078]++

									if err := checkHeader(&p, h); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:296
				_go_fuzz_dep_.CoverTab[5090]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if err == errServerTemporarilyMisbehaving {
//line /usr/local/go/src/net/dnsclient_unix.go:302
					_go_fuzz_dep_.CoverTab[5093]++
											dnsErr.IsTemporary = true
//line /usr/local/go/src/net/dnsclient_unix.go:303
					// _ = "end of CoverTab[5093]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:304
					_go_fuzz_dep_.CoverTab[5094]++
//line /usr/local/go/src/net/dnsclient_unix.go:304
					// _ = "end of CoverTab[5094]"
//line /usr/local/go/src/net/dnsclient_unix.go:304
				}
//line /usr/local/go/src/net/dnsclient_unix.go:304
				// _ = "end of CoverTab[5090]"
//line /usr/local/go/src/net/dnsclient_unix.go:304
				_go_fuzz_dep_.CoverTab[5091]++
										if err == errNoSuchHost {
//line /usr/local/go/src/net/dnsclient_unix.go:305
					_go_fuzz_dep_.CoverTab[5095]++

//line /usr/local/go/src/net/dnsclient_unix.go:309
					dnsErr.IsNotFound = true
											return p, server, dnsErr
//line /usr/local/go/src/net/dnsclient_unix.go:310
					// _ = "end of CoverTab[5095]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:311
					_go_fuzz_dep_.CoverTab[5096]++
//line /usr/local/go/src/net/dnsclient_unix.go:311
					// _ = "end of CoverTab[5096]"
//line /usr/local/go/src/net/dnsclient_unix.go:311
				}
//line /usr/local/go/src/net/dnsclient_unix.go:311
				// _ = "end of CoverTab[5091]"
//line /usr/local/go/src/net/dnsclient_unix.go:311
				_go_fuzz_dep_.CoverTab[5092]++
										lastErr = dnsErr
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:313
				// _ = "end of CoverTab[5092]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:314
				_go_fuzz_dep_.CoverTab[5097]++
//line /usr/local/go/src/net/dnsclient_unix.go:314
				// _ = "end of CoverTab[5097]"
//line /usr/local/go/src/net/dnsclient_unix.go:314
			}
//line /usr/local/go/src/net/dnsclient_unix.go:314
			// _ = "end of CoverTab[5078]"
//line /usr/local/go/src/net/dnsclient_unix.go:314
			_go_fuzz_dep_.CoverTab[5079]++

									err = skipToAnswer(&p, qtype)
									if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:317
				_go_fuzz_dep_.CoverTab[5098]++
										return p, server, nil
//line /usr/local/go/src/net/dnsclient_unix.go:318
				// _ = "end of CoverTab[5098]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:319
				_go_fuzz_dep_.CoverTab[5099]++
//line /usr/local/go/src/net/dnsclient_unix.go:319
				// _ = "end of CoverTab[5099]"
//line /usr/local/go/src/net/dnsclient_unix.go:319
			}
//line /usr/local/go/src/net/dnsclient_unix.go:319
			// _ = "end of CoverTab[5079]"
//line /usr/local/go/src/net/dnsclient_unix.go:319
			_go_fuzz_dep_.CoverTab[5080]++
									lastErr = &DNSError{
				Err:	err.Error(),
				Name:	name,
				Server:	server,
			}
			if err == errNoSuchHost {
//line /usr/local/go/src/net/dnsclient_unix.go:325
				_go_fuzz_dep_.CoverTab[5100]++

//line /usr/local/go/src/net/dnsclient_unix.go:329
				lastErr.(*DNSError).IsNotFound = true
										return p, server, lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:330
				// _ = "end of CoverTab[5100]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:331
				_go_fuzz_dep_.CoverTab[5101]++
//line /usr/local/go/src/net/dnsclient_unix.go:331
				// _ = "end of CoverTab[5101]"
//line /usr/local/go/src/net/dnsclient_unix.go:331
			}
//line /usr/local/go/src/net/dnsclient_unix.go:331
			// _ = "end of CoverTab[5080]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:332
		// _ = "end of CoverTab[5076]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:333
	// _ = "end of CoverTab[5072]"
//line /usr/local/go/src/net/dnsclient_unix.go:333
	_go_fuzz_dep_.CoverTab[5073]++
							return dnsmessage.Parser{}, "", lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:334
	// _ = "end of CoverTab[5073]"
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
//line /usr/local/go/src/net/dnsclient_unix.go:351
	_go_fuzz_dep_.CoverTab[5102]++
							resolvConf.tryUpdate("/etc/resolv.conf")
							return resolvConf.dnsConfig.Load()
//line /usr/local/go/src/net/dnsclient_unix.go:353
	// _ = "end of CoverTab[5102]"
}

// init initializes conf and is only called via conf.initOnce.
func (conf *resolverConfig) init() {

//line /usr/local/go/src/net/dnsclient_unix.go:360
	conf.dnsConfig.Store(dnsReadConfig("/etc/resolv.conf"))
							conf.lastChecked = time.Now()

//line /usr/local/go/src/net/dnsclient_unix.go:365
	conf.ch = make(chan struct{}, 1)
}

// tryUpdate tries to update conf with the named resolv.conf file.
//line /usr/local/go/src/net/dnsclient_unix.go:368
// The name variable only exists for testing. It is otherwise always
//line /usr/local/go/src/net/dnsclient_unix.go:368
// "/etc/resolv.conf".
//line /usr/local/go/src/net/dnsclient_unix.go:371
func (conf *resolverConfig) tryUpdate(name string) {
//line /usr/local/go/src/net/dnsclient_unix.go:371
	_go_fuzz_dep_.CoverTab[5103]++
							conf.initOnce.Do(conf.init)

							if conf.dnsConfig.Load().noReload {
//line /usr/local/go/src/net/dnsclient_unix.go:374
		_go_fuzz_dep_.CoverTab[5108]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:375
		// _ = "end of CoverTab[5108]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:376
		_go_fuzz_dep_.CoverTab[5109]++
//line /usr/local/go/src/net/dnsclient_unix.go:376
		// _ = "end of CoverTab[5109]"
//line /usr/local/go/src/net/dnsclient_unix.go:376
	}
//line /usr/local/go/src/net/dnsclient_unix.go:376
	// _ = "end of CoverTab[5103]"
//line /usr/local/go/src/net/dnsclient_unix.go:376
	_go_fuzz_dep_.CoverTab[5104]++

//line /usr/local/go/src/net/dnsclient_unix.go:379
	if !conf.tryAcquireSema() {
//line /usr/local/go/src/net/dnsclient_unix.go:379
		_go_fuzz_dep_.CoverTab[5110]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:380
		// _ = "end of CoverTab[5110]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:381
		_go_fuzz_dep_.CoverTab[5111]++
//line /usr/local/go/src/net/dnsclient_unix.go:381
		// _ = "end of CoverTab[5111]"
//line /usr/local/go/src/net/dnsclient_unix.go:381
	}
//line /usr/local/go/src/net/dnsclient_unix.go:381
	// _ = "end of CoverTab[5104]"
//line /usr/local/go/src/net/dnsclient_unix.go:381
	_go_fuzz_dep_.CoverTab[5105]++
							defer conf.releaseSema()

							now := time.Now()
							if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /usr/local/go/src/net/dnsclient_unix.go:385
		_go_fuzz_dep_.CoverTab[5112]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:386
		// _ = "end of CoverTab[5112]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:387
		_go_fuzz_dep_.CoverTab[5113]++
//line /usr/local/go/src/net/dnsclient_unix.go:387
		// _ = "end of CoverTab[5113]"
//line /usr/local/go/src/net/dnsclient_unix.go:387
	}
//line /usr/local/go/src/net/dnsclient_unix.go:387
	// _ = "end of CoverTab[5105]"
//line /usr/local/go/src/net/dnsclient_unix.go:387
	_go_fuzz_dep_.CoverTab[5106]++
							conf.lastChecked = now

							switch runtime.GOOS {
	case "windows":
//line /usr/local/go/src/net/dnsclient_unix.go:391
		_go_fuzz_dep_.CoverTab[5114]++
//line /usr/local/go/src/net/dnsclient_unix.go:391
		// _ = "end of CoverTab[5114]"

//line /usr/local/go/src/net/dnsclient_unix.go:397
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:397
		_go_fuzz_dep_.CoverTab[5115]++
								var mtime time.Time
								if fi, err := os.Stat(name); err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:399
			_go_fuzz_dep_.CoverTab[5117]++
									mtime = fi.ModTime()
//line /usr/local/go/src/net/dnsclient_unix.go:400
			// _ = "end of CoverTab[5117]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:401
			_go_fuzz_dep_.CoverTab[5118]++
//line /usr/local/go/src/net/dnsclient_unix.go:401
			// _ = "end of CoverTab[5118]"
//line /usr/local/go/src/net/dnsclient_unix.go:401
		}
//line /usr/local/go/src/net/dnsclient_unix.go:401
		// _ = "end of CoverTab[5115]"
//line /usr/local/go/src/net/dnsclient_unix.go:401
		_go_fuzz_dep_.CoverTab[5116]++
								if mtime.Equal(conf.dnsConfig.Load().mtime) {
//line /usr/local/go/src/net/dnsclient_unix.go:402
			_go_fuzz_dep_.CoverTab[5119]++
									return
//line /usr/local/go/src/net/dnsclient_unix.go:403
			// _ = "end of CoverTab[5119]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:404
			_go_fuzz_dep_.CoverTab[5120]++
//line /usr/local/go/src/net/dnsclient_unix.go:404
			// _ = "end of CoverTab[5120]"
//line /usr/local/go/src/net/dnsclient_unix.go:404
		}
//line /usr/local/go/src/net/dnsclient_unix.go:404
		// _ = "end of CoverTab[5116]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:405
	// _ = "end of CoverTab[5106]"
//line /usr/local/go/src/net/dnsclient_unix.go:405
	_go_fuzz_dep_.CoverTab[5107]++

							dnsConf := dnsReadConfig(name)
							conf.dnsConfig.Store(dnsConf)
//line /usr/local/go/src/net/dnsclient_unix.go:408
	// _ = "end of CoverTab[5107]"
}

func (conf *resolverConfig) tryAcquireSema() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:411
	_go_fuzz_dep_.CoverTab[5121]++
							select {
	case conf.ch <- struct{}{}:
//line /usr/local/go/src/net/dnsclient_unix.go:413
		_go_fuzz_dep_.CoverTab[5122]++
								return true
//line /usr/local/go/src/net/dnsclient_unix.go:414
		// _ = "end of CoverTab[5122]"
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:415
		_go_fuzz_dep_.CoverTab[5123]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:416
		// _ = "end of CoverTab[5123]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:417
	// _ = "end of CoverTab[5121]"
}

func (conf *resolverConfig) releaseSema() {
//line /usr/local/go/src/net/dnsclient_unix.go:420
	_go_fuzz_dep_.CoverTab[5124]++
							<-conf.ch
//line /usr/local/go/src/net/dnsclient_unix.go:421
	// _ = "end of CoverTab[5124]"
}

func (r *Resolver) lookup(ctx context.Context, name string, qtype dnsmessage.Type, conf *dnsConfig) (dnsmessage.Parser, string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:424
	_go_fuzz_dep_.CoverTab[5125]++
							if !isDomainName(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:425
		_go_fuzz_dep_.CoverTab[5131]++

//line /usr/local/go/src/net/dnsclient_unix.go:431
		return dnsmessage.Parser{}, "", &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:431
		// _ = "end of CoverTab[5131]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:432
		_go_fuzz_dep_.CoverTab[5132]++
//line /usr/local/go/src/net/dnsclient_unix.go:432
		// _ = "end of CoverTab[5132]"
//line /usr/local/go/src/net/dnsclient_unix.go:432
	}
//line /usr/local/go/src/net/dnsclient_unix.go:432
	// _ = "end of CoverTab[5125]"
//line /usr/local/go/src/net/dnsclient_unix.go:432
	_go_fuzz_dep_.CoverTab[5126]++

							if conf == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:434
		_go_fuzz_dep_.CoverTab[5133]++
								conf = getSystemDNSConfig()
//line /usr/local/go/src/net/dnsclient_unix.go:435
		// _ = "end of CoverTab[5133]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:436
		_go_fuzz_dep_.CoverTab[5134]++
//line /usr/local/go/src/net/dnsclient_unix.go:436
		// _ = "end of CoverTab[5134]"
//line /usr/local/go/src/net/dnsclient_unix.go:436
	}
//line /usr/local/go/src/net/dnsclient_unix.go:436
	// _ = "end of CoverTab[5126]"
//line /usr/local/go/src/net/dnsclient_unix.go:436
	_go_fuzz_dep_.CoverTab[5127]++

							var (
		p	dnsmessage.Parser
		server	string
		err	error
	)
	for _, fqdn := range conf.nameList(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:443
		_go_fuzz_dep_.CoverTab[5135]++
								p, server, err = r.tryOneName(ctx, conf, fqdn, qtype)
								if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:445
			_go_fuzz_dep_.CoverTab[5137]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:446
			// _ = "end of CoverTab[5137]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:447
			_go_fuzz_dep_.CoverTab[5138]++
//line /usr/local/go/src/net/dnsclient_unix.go:447
			// _ = "end of CoverTab[5138]"
//line /usr/local/go/src/net/dnsclient_unix.go:447
		}
//line /usr/local/go/src/net/dnsclient_unix.go:447
		// _ = "end of CoverTab[5135]"
//line /usr/local/go/src/net/dnsclient_unix.go:447
		_go_fuzz_dep_.CoverTab[5136]++
								if nerr, ok := err.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5139]++
//line /usr/local/go/src/net/dnsclient_unix.go:448
			return nerr.Temporary()
//line /usr/local/go/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[5139]"
//line /usr/local/go/src/net/dnsclient_unix.go:448
		}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5140]++
//line /usr/local/go/src/net/dnsclient_unix.go:448
			return r.strictErrors()
//line /usr/local/go/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[5140]"
//line /usr/local/go/src/net/dnsclient_unix.go:448
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[5141]++

//line /usr/local/go/src/net/dnsclient_unix.go:451
			break
//line /usr/local/go/src/net/dnsclient_unix.go:451
			// _ = "end of CoverTab[5141]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:452
			_go_fuzz_dep_.CoverTab[5142]++
//line /usr/local/go/src/net/dnsclient_unix.go:452
			// _ = "end of CoverTab[5142]"
//line /usr/local/go/src/net/dnsclient_unix.go:452
		}
//line /usr/local/go/src/net/dnsclient_unix.go:452
		// _ = "end of CoverTab[5136]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:453
	// _ = "end of CoverTab[5127]"
//line /usr/local/go/src/net/dnsclient_unix.go:453
	_go_fuzz_dep_.CoverTab[5128]++
							if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:454
		_go_fuzz_dep_.CoverTab[5143]++
								return p, server, nil
//line /usr/local/go/src/net/dnsclient_unix.go:455
		// _ = "end of CoverTab[5143]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:456
		_go_fuzz_dep_.CoverTab[5144]++
//line /usr/local/go/src/net/dnsclient_unix.go:456
		// _ = "end of CoverTab[5144]"
//line /usr/local/go/src/net/dnsclient_unix.go:456
	}
//line /usr/local/go/src/net/dnsclient_unix.go:456
	// _ = "end of CoverTab[5128]"
//line /usr/local/go/src/net/dnsclient_unix.go:456
	_go_fuzz_dep_.CoverTab[5129]++
							if err, ok := err.(*DNSError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:457
		_go_fuzz_dep_.CoverTab[5145]++

//line /usr/local/go/src/net/dnsclient_unix.go:461
		err.Name = name
//line /usr/local/go/src/net/dnsclient_unix.go:461
		// _ = "end of CoverTab[5145]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:462
		_go_fuzz_dep_.CoverTab[5146]++
//line /usr/local/go/src/net/dnsclient_unix.go:462
		// _ = "end of CoverTab[5146]"
//line /usr/local/go/src/net/dnsclient_unix.go:462
	}
//line /usr/local/go/src/net/dnsclient_unix.go:462
	// _ = "end of CoverTab[5129]"
//line /usr/local/go/src/net/dnsclient_unix.go:462
	_go_fuzz_dep_.CoverTab[5130]++
							return dnsmessage.Parser{}, "", err
//line /usr/local/go/src/net/dnsclient_unix.go:463
	// _ = "end of CoverTab[5130]"
}

// avoidDNS reports whether this is a hostname for which we should not
//line /usr/local/go/src/net/dnsclient_unix.go:466
// use DNS. Currently this includes only .onion, per RFC 7686. See
//line /usr/local/go/src/net/dnsclient_unix.go:466
// golang.org/issue/13705. Does not cover .local names (RFC 6762),
//line /usr/local/go/src/net/dnsclient_unix.go:466
// see golang.org/issue/16739.
//line /usr/local/go/src/net/dnsclient_unix.go:470
func avoidDNS(name string) bool {
//line /usr/local/go/src/net/dnsclient_unix.go:470
	_go_fuzz_dep_.CoverTab[5147]++
							if name == "" {
//line /usr/local/go/src/net/dnsclient_unix.go:471
		_go_fuzz_dep_.CoverTab[5150]++
								return true
//line /usr/local/go/src/net/dnsclient_unix.go:472
		// _ = "end of CoverTab[5150]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:473
		_go_fuzz_dep_.CoverTab[5151]++
//line /usr/local/go/src/net/dnsclient_unix.go:473
		// _ = "end of CoverTab[5151]"
//line /usr/local/go/src/net/dnsclient_unix.go:473
	}
//line /usr/local/go/src/net/dnsclient_unix.go:473
	// _ = "end of CoverTab[5147]"
//line /usr/local/go/src/net/dnsclient_unix.go:473
	_go_fuzz_dep_.CoverTab[5148]++
							if name[len(name)-1] == '.' {
//line /usr/local/go/src/net/dnsclient_unix.go:474
		_go_fuzz_dep_.CoverTab[5152]++
								name = name[:len(name)-1]
//line /usr/local/go/src/net/dnsclient_unix.go:475
		// _ = "end of CoverTab[5152]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:476
		_go_fuzz_dep_.CoverTab[5153]++
//line /usr/local/go/src/net/dnsclient_unix.go:476
		// _ = "end of CoverTab[5153]"
//line /usr/local/go/src/net/dnsclient_unix.go:476
	}
//line /usr/local/go/src/net/dnsclient_unix.go:476
	// _ = "end of CoverTab[5148]"
//line /usr/local/go/src/net/dnsclient_unix.go:476
	_go_fuzz_dep_.CoverTab[5149]++
							return stringsHasSuffixFold(name, ".onion")
//line /usr/local/go/src/net/dnsclient_unix.go:477
	// _ = "end of CoverTab[5149]"
}

// nameList returns a list of names for sequential DNS queries.
func (conf *dnsConfig) nameList(name string) []string {
//line /usr/local/go/src/net/dnsclient_unix.go:481
	_go_fuzz_dep_.CoverTab[5154]++
							if avoidDNS(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:482
		_go_fuzz_dep_.CoverTab[5161]++
								return nil
//line /usr/local/go/src/net/dnsclient_unix.go:483
		// _ = "end of CoverTab[5161]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:484
		_go_fuzz_dep_.CoverTab[5162]++
//line /usr/local/go/src/net/dnsclient_unix.go:484
		// _ = "end of CoverTab[5162]"
//line /usr/local/go/src/net/dnsclient_unix.go:484
	}
//line /usr/local/go/src/net/dnsclient_unix.go:484
	// _ = "end of CoverTab[5154]"
//line /usr/local/go/src/net/dnsclient_unix.go:484
	_go_fuzz_dep_.CoverTab[5155]++

//line /usr/local/go/src/net/dnsclient_unix.go:487
	l := len(name)
	rooted := l > 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:488
		_go_fuzz_dep_.CoverTab[5163]++
//line /usr/local/go/src/net/dnsclient_unix.go:488
		return name[l-1] == '.'
//line /usr/local/go/src/net/dnsclient_unix.go:488
		// _ = "end of CoverTab[5163]"
//line /usr/local/go/src/net/dnsclient_unix.go:488
	}()
							if l > 254 || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[5164]++
//line /usr/local/go/src/net/dnsclient_unix.go:489
		return l == 254 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:489
			_go_fuzz_dep_.CoverTab[5165]++
//line /usr/local/go/src/net/dnsclient_unix.go:489
			return !rooted
//line /usr/local/go/src/net/dnsclient_unix.go:489
			// _ = "end of CoverTab[5165]"
//line /usr/local/go/src/net/dnsclient_unix.go:489
		}()
//line /usr/local/go/src/net/dnsclient_unix.go:489
		// _ = "end of CoverTab[5164]"
//line /usr/local/go/src/net/dnsclient_unix.go:489
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[5166]++
								return nil
//line /usr/local/go/src/net/dnsclient_unix.go:490
		// _ = "end of CoverTab[5166]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:491
		_go_fuzz_dep_.CoverTab[5167]++
//line /usr/local/go/src/net/dnsclient_unix.go:491
		// _ = "end of CoverTab[5167]"
//line /usr/local/go/src/net/dnsclient_unix.go:491
	}
//line /usr/local/go/src/net/dnsclient_unix.go:491
	// _ = "end of CoverTab[5155]"
//line /usr/local/go/src/net/dnsclient_unix.go:491
	_go_fuzz_dep_.CoverTab[5156]++

//line /usr/local/go/src/net/dnsclient_unix.go:494
	if rooted {
//line /usr/local/go/src/net/dnsclient_unix.go:494
		_go_fuzz_dep_.CoverTab[5168]++
								return []string{name}
//line /usr/local/go/src/net/dnsclient_unix.go:495
		// _ = "end of CoverTab[5168]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:496
		_go_fuzz_dep_.CoverTab[5169]++
//line /usr/local/go/src/net/dnsclient_unix.go:496
		// _ = "end of CoverTab[5169]"
//line /usr/local/go/src/net/dnsclient_unix.go:496
	}
//line /usr/local/go/src/net/dnsclient_unix.go:496
	// _ = "end of CoverTab[5156]"
//line /usr/local/go/src/net/dnsclient_unix.go:496
	_go_fuzz_dep_.CoverTab[5157]++

							hasNdots := count(name, '.') >= conf.ndots
							name += "."
							l++

//line /usr/local/go/src/net/dnsclient_unix.go:503
	names := make([]string, 0, 1+len(conf.search))

	if hasNdots {
//line /usr/local/go/src/net/dnsclient_unix.go:505
		_go_fuzz_dep_.CoverTab[5170]++
								names = append(names, name)
//line /usr/local/go/src/net/dnsclient_unix.go:506
		// _ = "end of CoverTab[5170]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:507
		_go_fuzz_dep_.CoverTab[5171]++
//line /usr/local/go/src/net/dnsclient_unix.go:507
		// _ = "end of CoverTab[5171]"
//line /usr/local/go/src/net/dnsclient_unix.go:507
	}
//line /usr/local/go/src/net/dnsclient_unix.go:507
	// _ = "end of CoverTab[5157]"
//line /usr/local/go/src/net/dnsclient_unix.go:507
	_go_fuzz_dep_.CoverTab[5158]++

							for _, suffix := range conf.search {
//line /usr/local/go/src/net/dnsclient_unix.go:509
		_go_fuzz_dep_.CoverTab[5172]++
								if l+len(suffix) <= 254 {
//line /usr/local/go/src/net/dnsclient_unix.go:510
			_go_fuzz_dep_.CoverTab[5173]++
									names = append(names, name+suffix)
//line /usr/local/go/src/net/dnsclient_unix.go:511
			// _ = "end of CoverTab[5173]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:512
			_go_fuzz_dep_.CoverTab[5174]++
//line /usr/local/go/src/net/dnsclient_unix.go:512
			// _ = "end of CoverTab[5174]"
//line /usr/local/go/src/net/dnsclient_unix.go:512
		}
//line /usr/local/go/src/net/dnsclient_unix.go:512
		// _ = "end of CoverTab[5172]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:513
	// _ = "end of CoverTab[5158]"
//line /usr/local/go/src/net/dnsclient_unix.go:513
	_go_fuzz_dep_.CoverTab[5159]++

							if !hasNdots {
//line /usr/local/go/src/net/dnsclient_unix.go:515
		_go_fuzz_dep_.CoverTab[5175]++
								names = append(names, name)
//line /usr/local/go/src/net/dnsclient_unix.go:516
		// _ = "end of CoverTab[5175]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:517
		_go_fuzz_dep_.CoverTab[5176]++
//line /usr/local/go/src/net/dnsclient_unix.go:517
		// _ = "end of CoverTab[5176]"
//line /usr/local/go/src/net/dnsclient_unix.go:517
	}
//line /usr/local/go/src/net/dnsclient_unix.go:517
	// _ = "end of CoverTab[5159]"
//line /usr/local/go/src/net/dnsclient_unix.go:517
	_go_fuzz_dep_.CoverTab[5160]++
							return names
//line /usr/local/go/src/net/dnsclient_unix.go:518
	// _ = "end of CoverTab[5160]"
}

// hostLookupOrder specifies the order of LookupHost lookup strategies.
//line /usr/local/go/src/net/dnsclient_unix.go:521
// It is basically a simplified representation of nsswitch.conf.
//line /usr/local/go/src/net/dnsclient_unix.go:521
// "files" means /etc/hosts.
//line /usr/local/go/src/net/dnsclient_unix.go:524
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
//line /usr/local/go/src/net/dnsclient_unix.go:543
	_go_fuzz_dep_.CoverTab[5177]++
							if s, ok := lookupOrderName[o]; ok {
//line /usr/local/go/src/net/dnsclient_unix.go:544
		_go_fuzz_dep_.CoverTab[5179]++
								return s
//line /usr/local/go/src/net/dnsclient_unix.go:545
		// _ = "end of CoverTab[5179]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:546
		_go_fuzz_dep_.CoverTab[5180]++
//line /usr/local/go/src/net/dnsclient_unix.go:546
		// _ = "end of CoverTab[5180]"
//line /usr/local/go/src/net/dnsclient_unix.go:546
	}
//line /usr/local/go/src/net/dnsclient_unix.go:546
	// _ = "end of CoverTab[5177]"
//line /usr/local/go/src/net/dnsclient_unix.go:546
	_go_fuzz_dep_.CoverTab[5178]++
							return "hostLookupOrder=" + itoa.Itoa(int(o)) + "??"
//line /usr/local/go/src/net/dnsclient_unix.go:547
	// _ = "end of CoverTab[5178]"
}

func (r *Resolver) goLookupHostOrder(ctx context.Context, name string, order hostLookupOrder, conf *dnsConfig) (addrs []string, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:550
	_go_fuzz_dep_.CoverTab[5181]++
							if order == hostLookupFilesDNS || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[5185]++
//line /usr/local/go/src/net/dnsclient_unix.go:551
		return order == hostLookupFiles
//line /usr/local/go/src/net/dnsclient_unix.go:551
		// _ = "end of CoverTab[5185]"
//line /usr/local/go/src/net/dnsclient_unix.go:551
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[5186]++

								addrs, _ = lookupStaticHost(name)
								if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:554
			_go_fuzz_dep_.CoverTab[5188]++
									return
//line /usr/local/go/src/net/dnsclient_unix.go:555
			// _ = "end of CoverTab[5188]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:556
			_go_fuzz_dep_.CoverTab[5189]++
//line /usr/local/go/src/net/dnsclient_unix.go:556
			// _ = "end of CoverTab[5189]"
//line /usr/local/go/src/net/dnsclient_unix.go:556
		}
//line /usr/local/go/src/net/dnsclient_unix.go:556
		// _ = "end of CoverTab[5186]"
//line /usr/local/go/src/net/dnsclient_unix.go:556
		_go_fuzz_dep_.CoverTab[5187]++

								if order == hostLookupFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:558
			_go_fuzz_dep_.CoverTab[5190]++
									return nil, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:559
			// _ = "end of CoverTab[5190]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:560
			_go_fuzz_dep_.CoverTab[5191]++
//line /usr/local/go/src/net/dnsclient_unix.go:560
			// _ = "end of CoverTab[5191]"
//line /usr/local/go/src/net/dnsclient_unix.go:560
		}
//line /usr/local/go/src/net/dnsclient_unix.go:560
		// _ = "end of CoverTab[5187]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:561
		_go_fuzz_dep_.CoverTab[5192]++
//line /usr/local/go/src/net/dnsclient_unix.go:561
		// _ = "end of CoverTab[5192]"
//line /usr/local/go/src/net/dnsclient_unix.go:561
	}
//line /usr/local/go/src/net/dnsclient_unix.go:561
	// _ = "end of CoverTab[5181]"
//line /usr/local/go/src/net/dnsclient_unix.go:561
	_go_fuzz_dep_.CoverTab[5182]++
							ips, _, err := r.goLookupIPCNAMEOrder(ctx, "ip", name, order, conf)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:563
		_go_fuzz_dep_.CoverTab[5193]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:564
		// _ = "end of CoverTab[5193]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:565
		_go_fuzz_dep_.CoverTab[5194]++
//line /usr/local/go/src/net/dnsclient_unix.go:565
		// _ = "end of CoverTab[5194]"
//line /usr/local/go/src/net/dnsclient_unix.go:565
	}
//line /usr/local/go/src/net/dnsclient_unix.go:565
	// _ = "end of CoverTab[5182]"
//line /usr/local/go/src/net/dnsclient_unix.go:565
	_go_fuzz_dep_.CoverTab[5183]++
							addrs = make([]string, 0, len(ips))
							for _, ip := range ips {
//line /usr/local/go/src/net/dnsclient_unix.go:567
		_go_fuzz_dep_.CoverTab[5195]++
								addrs = append(addrs, ip.String())
//line /usr/local/go/src/net/dnsclient_unix.go:568
		// _ = "end of CoverTab[5195]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:569
	// _ = "end of CoverTab[5183]"
//line /usr/local/go/src/net/dnsclient_unix.go:569
	_go_fuzz_dep_.CoverTab[5184]++
							return
//line /usr/local/go/src/net/dnsclient_unix.go:570
	// _ = "end of CoverTab[5184]"
}

// lookup entries from /etc/hosts
func goLookupIPFiles(name string) (addrs []IPAddr, canonical string) {
//line /usr/local/go/src/net/dnsclient_unix.go:574
	_go_fuzz_dep_.CoverTab[5196]++
							addr, canonical := lookupStaticHost(name)
							for _, haddr := range addr {
//line /usr/local/go/src/net/dnsclient_unix.go:576
		_go_fuzz_dep_.CoverTab[5198]++
								haddr, zone := splitHostZone(haddr)
								if ip := ParseIP(haddr); ip != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:578
			_go_fuzz_dep_.CoverTab[5199]++
									addr := IPAddr{IP: ip, Zone: zone}
									addrs = append(addrs, addr)
//line /usr/local/go/src/net/dnsclient_unix.go:580
			// _ = "end of CoverTab[5199]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:581
			_go_fuzz_dep_.CoverTab[5200]++
//line /usr/local/go/src/net/dnsclient_unix.go:581
			// _ = "end of CoverTab[5200]"
//line /usr/local/go/src/net/dnsclient_unix.go:581
		}
//line /usr/local/go/src/net/dnsclient_unix.go:581
		// _ = "end of CoverTab[5198]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:582
	// _ = "end of CoverTab[5196]"
//line /usr/local/go/src/net/dnsclient_unix.go:582
	_go_fuzz_dep_.CoverTab[5197]++
							sortByRFC6724(addrs)
							return addrs, canonical
//line /usr/local/go/src/net/dnsclient_unix.go:584
	// _ = "end of CoverTab[5197]"
}

// goLookupIP is the native Go implementation of LookupIP.
//line /usr/local/go/src/net/dnsclient_unix.go:587
// The libc versions are in cgo_*.go.
//line /usr/local/go/src/net/dnsclient_unix.go:589
func (r *Resolver) goLookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:589
	_go_fuzz_dep_.CoverTab[5201]++
							order, conf := systemConf().hostLookupOrder(r, host)
							addrs, _, err = r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
							return
//line /usr/local/go/src/net/dnsclient_unix.go:592
	// _ = "end of CoverTab[5201]"
}

func (r *Resolver) goLookupIPCNAMEOrder(ctx context.Context, network, name string, order hostLookupOrder, conf *dnsConfig) (addrs []IPAddr, cname dnsmessage.Name, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:595
	_go_fuzz_dep_.CoverTab[5202]++
							if order == hostLookupFilesDNS || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[5212]++
//line /usr/local/go/src/net/dnsclient_unix.go:596
		return order == hostLookupFiles
//line /usr/local/go/src/net/dnsclient_unix.go:596
		// _ = "end of CoverTab[5212]"
//line /usr/local/go/src/net/dnsclient_unix.go:596
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[5213]++
								var canonical string
								addrs, canonical = goLookupIPFiles(name)

								if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:600
			_go_fuzz_dep_.CoverTab[5215]++
									var err error
									cname, err = dnsmessage.NewName(canonical)
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:603
				_go_fuzz_dep_.CoverTab[5217]++
										return nil, dnsmessage.Name{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:604
				// _ = "end of CoverTab[5217]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:605
				_go_fuzz_dep_.CoverTab[5218]++
//line /usr/local/go/src/net/dnsclient_unix.go:605
				// _ = "end of CoverTab[5218]"
//line /usr/local/go/src/net/dnsclient_unix.go:605
			}
//line /usr/local/go/src/net/dnsclient_unix.go:605
			// _ = "end of CoverTab[5215]"
//line /usr/local/go/src/net/dnsclient_unix.go:605
			_go_fuzz_dep_.CoverTab[5216]++
									return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:606
			// _ = "end of CoverTab[5216]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:607
			_go_fuzz_dep_.CoverTab[5219]++
//line /usr/local/go/src/net/dnsclient_unix.go:607
			// _ = "end of CoverTab[5219]"
//line /usr/local/go/src/net/dnsclient_unix.go:607
		}
//line /usr/local/go/src/net/dnsclient_unix.go:607
		// _ = "end of CoverTab[5213]"
//line /usr/local/go/src/net/dnsclient_unix.go:607
		_go_fuzz_dep_.CoverTab[5214]++

								if order == hostLookupFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:609
			_go_fuzz_dep_.CoverTab[5220]++
									return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:610
			// _ = "end of CoverTab[5220]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:611
			_go_fuzz_dep_.CoverTab[5221]++
//line /usr/local/go/src/net/dnsclient_unix.go:611
			// _ = "end of CoverTab[5221]"
//line /usr/local/go/src/net/dnsclient_unix.go:611
		}
//line /usr/local/go/src/net/dnsclient_unix.go:611
		// _ = "end of CoverTab[5214]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:612
		_go_fuzz_dep_.CoverTab[5222]++
//line /usr/local/go/src/net/dnsclient_unix.go:612
		// _ = "end of CoverTab[5222]"
//line /usr/local/go/src/net/dnsclient_unix.go:612
	}
//line /usr/local/go/src/net/dnsclient_unix.go:612
	// _ = "end of CoverTab[5202]"
//line /usr/local/go/src/net/dnsclient_unix.go:612
	_go_fuzz_dep_.CoverTab[5203]++

							if !isDomainName(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:614
		_go_fuzz_dep_.CoverTab[5223]++

								return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:616
		// _ = "end of CoverTab[5223]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:617
		_go_fuzz_dep_.CoverTab[5224]++
//line /usr/local/go/src/net/dnsclient_unix.go:617
		// _ = "end of CoverTab[5224]"
//line /usr/local/go/src/net/dnsclient_unix.go:617
	}
//line /usr/local/go/src/net/dnsclient_unix.go:617
	// _ = "end of CoverTab[5203]"
//line /usr/local/go/src/net/dnsclient_unix.go:617
	_go_fuzz_dep_.CoverTab[5204]++
							type result struct {
		p	dnsmessage.Parser
		server	string
		error
	}

	if conf == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:624
		_go_fuzz_dep_.CoverTab[5225]++
								conf = getSystemDNSConfig()
//line /usr/local/go/src/net/dnsclient_unix.go:625
		// _ = "end of CoverTab[5225]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:626
		_go_fuzz_dep_.CoverTab[5226]++
//line /usr/local/go/src/net/dnsclient_unix.go:626
		// _ = "end of CoverTab[5226]"
//line /usr/local/go/src/net/dnsclient_unix.go:626
	}
//line /usr/local/go/src/net/dnsclient_unix.go:626
	// _ = "end of CoverTab[5204]"
//line /usr/local/go/src/net/dnsclient_unix.go:626
	_go_fuzz_dep_.CoverTab[5205]++

							lane := make(chan result, 1)
							qtypes := []dnsmessage.Type{dnsmessage.TypeA, dnsmessage.TypeAAAA}
							if network == "CNAME" {
//line /usr/local/go/src/net/dnsclient_unix.go:630
		_go_fuzz_dep_.CoverTab[5227]++
								qtypes = append(qtypes, dnsmessage.TypeCNAME)
//line /usr/local/go/src/net/dnsclient_unix.go:631
		// _ = "end of CoverTab[5227]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:632
		_go_fuzz_dep_.CoverTab[5228]++
//line /usr/local/go/src/net/dnsclient_unix.go:632
		// _ = "end of CoverTab[5228]"
//line /usr/local/go/src/net/dnsclient_unix.go:632
	}
//line /usr/local/go/src/net/dnsclient_unix.go:632
	// _ = "end of CoverTab[5205]"
//line /usr/local/go/src/net/dnsclient_unix.go:632
	_go_fuzz_dep_.CoverTab[5206]++
							switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/dnsclient_unix.go:634
		_go_fuzz_dep_.CoverTab[5229]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeA}
//line /usr/local/go/src/net/dnsclient_unix.go:635
		// _ = "end of CoverTab[5229]"
	case '6':
//line /usr/local/go/src/net/dnsclient_unix.go:636
		_go_fuzz_dep_.CoverTab[5230]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeAAAA}
//line /usr/local/go/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[5230]"
//line /usr/local/go/src/net/dnsclient_unix.go:637
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:637
		_go_fuzz_dep_.CoverTab[5231]++
//line /usr/local/go/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[5231]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:638
	// _ = "end of CoverTab[5206]"
//line /usr/local/go/src/net/dnsclient_unix.go:638
	_go_fuzz_dep_.CoverTab[5207]++
							var queryFn func(fqdn string, qtype dnsmessage.Type)
							var responseFn func(fqdn string, qtype dnsmessage.Type) result
							if conf.singleRequest {
//line /usr/local/go/src/net/dnsclient_unix.go:641
		_go_fuzz_dep_.CoverTab[5232]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:642
			_go_fuzz_dep_.CoverTab[5234]++
//line /usr/local/go/src/net/dnsclient_unix.go:642
			// _ = "end of CoverTab[5234]"
//line /usr/local/go/src/net/dnsclient_unix.go:642
		}
//line /usr/local/go/src/net/dnsclient_unix.go:642
		// _ = "end of CoverTab[5232]"
//line /usr/local/go/src/net/dnsclient_unix.go:642
		_go_fuzz_dep_.CoverTab[5233]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /usr/local/go/src/net/dnsclient_unix.go:643
			_go_fuzz_dep_.CoverTab[5235]++
									dnsWaitGroup.Add(1)
									defer dnsWaitGroup.Done()
									p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
									return result{p, server, err}
//line /usr/local/go/src/net/dnsclient_unix.go:647
			// _ = "end of CoverTab[5235]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:648
		// _ = "end of CoverTab[5233]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:649
		_go_fuzz_dep_.CoverTab[5236]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:650
			_go_fuzz_dep_.CoverTab[5238]++
									dnsWaitGroup.Add(1)
//line /usr/local/go/src/net/dnsclient_unix.go:651
			_curRoutineNum8_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/dnsclient_unix.go:651
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum8_)
									go func(qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:652
				_go_fuzz_dep_.CoverTab[5239]++
//line /usr/local/go/src/net/dnsclient_unix.go:652
				defer func() {
//line /usr/local/go/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.CoverTab[5240]++
//line /usr/local/go/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum8_)
//line /usr/local/go/src/net/dnsclient_unix.go:652
					// _ = "end of CoverTab[5240]"
//line /usr/local/go/src/net/dnsclient_unix.go:652
				}()
										p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
										lane <- result{p, server, err}
										dnsWaitGroup.Done()
//line /usr/local/go/src/net/dnsclient_unix.go:655
				// _ = "end of CoverTab[5239]"
			}(qtype)
//line /usr/local/go/src/net/dnsclient_unix.go:656
			// _ = "end of CoverTab[5238]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:657
		// _ = "end of CoverTab[5236]"
//line /usr/local/go/src/net/dnsclient_unix.go:657
		_go_fuzz_dep_.CoverTab[5237]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /usr/local/go/src/net/dnsclient_unix.go:658
			_go_fuzz_dep_.CoverTab[5241]++
									return <-lane
//line /usr/local/go/src/net/dnsclient_unix.go:659
			// _ = "end of CoverTab[5241]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:660
		// _ = "end of CoverTab[5237]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:661
	// _ = "end of CoverTab[5207]"
//line /usr/local/go/src/net/dnsclient_unix.go:661
	_go_fuzz_dep_.CoverTab[5208]++
							var lastErr error
							for _, fqdn := range conf.nameList(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:663
		_go_fuzz_dep_.CoverTab[5242]++
								for _, qtype := range qtypes {
//line /usr/local/go/src/net/dnsclient_unix.go:664
			_go_fuzz_dep_.CoverTab[5246]++
									queryFn(fqdn, qtype)
//line /usr/local/go/src/net/dnsclient_unix.go:665
			// _ = "end of CoverTab[5246]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:666
		// _ = "end of CoverTab[5242]"
//line /usr/local/go/src/net/dnsclient_unix.go:666
		_go_fuzz_dep_.CoverTab[5243]++
								hitStrictError := false
								for _, qtype := range qtypes {
//line /usr/local/go/src/net/dnsclient_unix.go:668
			_go_fuzz_dep_.CoverTab[5247]++
									result := responseFn(fqdn, qtype)
									if result.error != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:670
				_go_fuzz_dep_.CoverTab[5249]++
										if nerr, ok := result.error.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5251]++
//line /usr/local/go/src/net/dnsclient_unix.go:671
					return nerr.Temporary()
//line /usr/local/go/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[5251]"
//line /usr/local/go/src/net/dnsclient_unix.go:671
				}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5252]++
//line /usr/local/go/src/net/dnsclient_unix.go:671
					return r.strictErrors()
//line /usr/local/go/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[5252]"
//line /usr/local/go/src/net/dnsclient_unix.go:671
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[5253]++

											hitStrictError = true
											lastErr = result.error
//line /usr/local/go/src/net/dnsclient_unix.go:674
					// _ = "end of CoverTab[5253]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:675
					_go_fuzz_dep_.CoverTab[5254]++
//line /usr/local/go/src/net/dnsclient_unix.go:675
					if lastErr == nil || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[5255]++
//line /usr/local/go/src/net/dnsclient_unix.go:675
						return fqdn == name+"."
//line /usr/local/go/src/net/dnsclient_unix.go:675
						// _ = "end of CoverTab[5255]"
//line /usr/local/go/src/net/dnsclient_unix.go:675
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[5256]++

												lastErr = result.error
//line /usr/local/go/src/net/dnsclient_unix.go:677
						// _ = "end of CoverTab[5256]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:678
						_go_fuzz_dep_.CoverTab[5257]++
//line /usr/local/go/src/net/dnsclient_unix.go:678
						// _ = "end of CoverTab[5257]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
					}
//line /usr/local/go/src/net/dnsclient_unix.go:678
					// _ = "end of CoverTab[5254]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
				}
//line /usr/local/go/src/net/dnsclient_unix.go:678
				// _ = "end of CoverTab[5249]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
				_go_fuzz_dep_.CoverTab[5250]++
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:679
				// _ = "end of CoverTab[5250]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:680
				_go_fuzz_dep_.CoverTab[5258]++
//line /usr/local/go/src/net/dnsclient_unix.go:680
				// _ = "end of CoverTab[5258]"
//line /usr/local/go/src/net/dnsclient_unix.go:680
			}
//line /usr/local/go/src/net/dnsclient_unix.go:680
			// _ = "end of CoverTab[5247]"
//line /usr/local/go/src/net/dnsclient_unix.go:680
			_go_fuzz_dep_.CoverTab[5248]++

//line /usr/local/go/src/net/dnsclient_unix.go:697
		loop:
			for {
//line /usr/local/go/src/net/dnsclient_unix.go:698
				_go_fuzz_dep_.CoverTab[5259]++
										h, err := result.p.AnswerHeader()
										if err != nil && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[5262]++
//line /usr/local/go/src/net/dnsclient_unix.go:700
					return err != dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:700
					// _ = "end of CoverTab[5262]"
//line /usr/local/go/src/net/dnsclient_unix.go:700
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[5263]++
											lastErr = &DNSError{
						Err:	"cannot marshal DNS message",
						Name:	name,
						Server:	result.server,
					}
//line /usr/local/go/src/net/dnsclient_unix.go:705
					// _ = "end of CoverTab[5263]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:706
					_go_fuzz_dep_.CoverTab[5264]++
//line /usr/local/go/src/net/dnsclient_unix.go:706
					// _ = "end of CoverTab[5264]"
//line /usr/local/go/src/net/dnsclient_unix.go:706
				}
//line /usr/local/go/src/net/dnsclient_unix.go:706
				// _ = "end of CoverTab[5259]"
//line /usr/local/go/src/net/dnsclient_unix.go:706
				_go_fuzz_dep_.CoverTab[5260]++
										if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:707
					_go_fuzz_dep_.CoverTab[5265]++
											break
//line /usr/local/go/src/net/dnsclient_unix.go:708
					// _ = "end of CoverTab[5265]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:709
					_go_fuzz_dep_.CoverTab[5266]++
//line /usr/local/go/src/net/dnsclient_unix.go:709
					// _ = "end of CoverTab[5266]"
//line /usr/local/go/src/net/dnsclient_unix.go:709
				}
//line /usr/local/go/src/net/dnsclient_unix.go:709
				// _ = "end of CoverTab[5260]"
//line /usr/local/go/src/net/dnsclient_unix.go:709
				_go_fuzz_dep_.CoverTab[5261]++
										switch h.Type {
				case dnsmessage.TypeA:
//line /usr/local/go/src/net/dnsclient_unix.go:711
					_go_fuzz_dep_.CoverTab[5267]++
											a, err := result.p.AResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:713
						_go_fuzz_dep_.CoverTab[5275]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:719
						// _ = "end of CoverTab[5275]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:720
						_go_fuzz_dep_.CoverTab[5276]++
//line /usr/local/go/src/net/dnsclient_unix.go:720
						// _ = "end of CoverTab[5276]"
//line /usr/local/go/src/net/dnsclient_unix.go:720
					}
//line /usr/local/go/src/net/dnsclient_unix.go:720
					// _ = "end of CoverTab[5267]"
//line /usr/local/go/src/net/dnsclient_unix.go:720
					_go_fuzz_dep_.CoverTab[5268]++
											addrs = append(addrs, IPAddr{IP: IP(a.A[:])})
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[5277]++
//line /usr/local/go/src/net/dnsclient_unix.go:722
						return h.Name.Length != 0
//line /usr/local/go/src/net/dnsclient_unix.go:722
						// _ = "end of CoverTab[5277]"
//line /usr/local/go/src/net/dnsclient_unix.go:722
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[5278]++
												cname = h.Name
//line /usr/local/go/src/net/dnsclient_unix.go:723
						// _ = "end of CoverTab[5278]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:724
						_go_fuzz_dep_.CoverTab[5279]++
//line /usr/local/go/src/net/dnsclient_unix.go:724
						// _ = "end of CoverTab[5279]"
//line /usr/local/go/src/net/dnsclient_unix.go:724
					}
//line /usr/local/go/src/net/dnsclient_unix.go:724
					// _ = "end of CoverTab[5268]"

				case dnsmessage.TypeAAAA:
//line /usr/local/go/src/net/dnsclient_unix.go:726
					_go_fuzz_dep_.CoverTab[5269]++
											aaaa, err := result.p.AAAAResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:728
						_go_fuzz_dep_.CoverTab[5280]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:734
						// _ = "end of CoverTab[5280]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:735
						_go_fuzz_dep_.CoverTab[5281]++
//line /usr/local/go/src/net/dnsclient_unix.go:735
						// _ = "end of CoverTab[5281]"
//line /usr/local/go/src/net/dnsclient_unix.go:735
					}
//line /usr/local/go/src/net/dnsclient_unix.go:735
					// _ = "end of CoverTab[5269]"
//line /usr/local/go/src/net/dnsclient_unix.go:735
					_go_fuzz_dep_.CoverTab[5270]++
											addrs = append(addrs, IPAddr{IP: IP(aaaa.AAAA[:])})
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[5282]++
//line /usr/local/go/src/net/dnsclient_unix.go:737
						return h.Name.Length != 0
//line /usr/local/go/src/net/dnsclient_unix.go:737
						// _ = "end of CoverTab[5282]"
//line /usr/local/go/src/net/dnsclient_unix.go:737
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[5283]++
												cname = h.Name
//line /usr/local/go/src/net/dnsclient_unix.go:738
						// _ = "end of CoverTab[5283]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:739
						_go_fuzz_dep_.CoverTab[5284]++
//line /usr/local/go/src/net/dnsclient_unix.go:739
						// _ = "end of CoverTab[5284]"
//line /usr/local/go/src/net/dnsclient_unix.go:739
					}
//line /usr/local/go/src/net/dnsclient_unix.go:739
					// _ = "end of CoverTab[5270]"

				case dnsmessage.TypeCNAME:
//line /usr/local/go/src/net/dnsclient_unix.go:741
					_go_fuzz_dep_.CoverTab[5271]++
											c, err := result.p.CNAMEResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:743
						_go_fuzz_dep_.CoverTab[5285]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:749
						// _ = "end of CoverTab[5285]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:750
						_go_fuzz_dep_.CoverTab[5286]++
//line /usr/local/go/src/net/dnsclient_unix.go:750
						// _ = "end of CoverTab[5286]"
//line /usr/local/go/src/net/dnsclient_unix.go:750
					}
//line /usr/local/go/src/net/dnsclient_unix.go:750
					// _ = "end of CoverTab[5271]"
//line /usr/local/go/src/net/dnsclient_unix.go:750
					_go_fuzz_dep_.CoverTab[5272]++
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[5287]++
//line /usr/local/go/src/net/dnsclient_unix.go:751
						return c.CNAME.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:751
						// _ = "end of CoverTab[5287]"
//line /usr/local/go/src/net/dnsclient_unix.go:751
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[5288]++
												cname = c.CNAME
//line /usr/local/go/src/net/dnsclient_unix.go:752
						// _ = "end of CoverTab[5288]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:753
						_go_fuzz_dep_.CoverTab[5289]++
//line /usr/local/go/src/net/dnsclient_unix.go:753
						// _ = "end of CoverTab[5289]"
//line /usr/local/go/src/net/dnsclient_unix.go:753
					}
//line /usr/local/go/src/net/dnsclient_unix.go:753
					// _ = "end of CoverTab[5272]"

				default:
//line /usr/local/go/src/net/dnsclient_unix.go:755
					_go_fuzz_dep_.CoverTab[5273]++
											if err := result.p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:756
						_go_fuzz_dep_.CoverTab[5290]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:762
						// _ = "end of CoverTab[5290]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:763
						_go_fuzz_dep_.CoverTab[5291]++
//line /usr/local/go/src/net/dnsclient_unix.go:763
						// _ = "end of CoverTab[5291]"
//line /usr/local/go/src/net/dnsclient_unix.go:763
					}
//line /usr/local/go/src/net/dnsclient_unix.go:763
					// _ = "end of CoverTab[5273]"
//line /usr/local/go/src/net/dnsclient_unix.go:763
					_go_fuzz_dep_.CoverTab[5274]++
											continue
//line /usr/local/go/src/net/dnsclient_unix.go:764
					// _ = "end of CoverTab[5274]"
				}
//line /usr/local/go/src/net/dnsclient_unix.go:765
				// _ = "end of CoverTab[5261]"
			}
//line /usr/local/go/src/net/dnsclient_unix.go:766
			// _ = "end of CoverTab[5248]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:767
		// _ = "end of CoverTab[5243]"
//line /usr/local/go/src/net/dnsclient_unix.go:767
		_go_fuzz_dep_.CoverTab[5244]++
								if hitStrictError {
//line /usr/local/go/src/net/dnsclient_unix.go:768
			_go_fuzz_dep_.CoverTab[5292]++

//line /usr/local/go/src/net/dnsclient_unix.go:772
			addrs = nil
									break
//line /usr/local/go/src/net/dnsclient_unix.go:773
			// _ = "end of CoverTab[5292]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:774
			_go_fuzz_dep_.CoverTab[5293]++
//line /usr/local/go/src/net/dnsclient_unix.go:774
			// _ = "end of CoverTab[5293]"
//line /usr/local/go/src/net/dnsclient_unix.go:774
		}
//line /usr/local/go/src/net/dnsclient_unix.go:774
		// _ = "end of CoverTab[5244]"
//line /usr/local/go/src/net/dnsclient_unix.go:774
		_go_fuzz_dep_.CoverTab[5245]++
								if len(addrs) > 0 || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[5294]++
//line /usr/local/go/src/net/dnsclient_unix.go:775
			return network == "CNAME" && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:775
				_go_fuzz_dep_.CoverTab[5295]++
//line /usr/local/go/src/net/dnsclient_unix.go:775
				return cname.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:775
				// _ = "end of CoverTab[5295]"
//line /usr/local/go/src/net/dnsclient_unix.go:775
			}()
//line /usr/local/go/src/net/dnsclient_unix.go:775
			// _ = "end of CoverTab[5294]"
//line /usr/local/go/src/net/dnsclient_unix.go:775
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[5296]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:776
			// _ = "end of CoverTab[5296]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:777
			_go_fuzz_dep_.CoverTab[5297]++
//line /usr/local/go/src/net/dnsclient_unix.go:777
			// _ = "end of CoverTab[5297]"
//line /usr/local/go/src/net/dnsclient_unix.go:777
		}
//line /usr/local/go/src/net/dnsclient_unix.go:777
		// _ = "end of CoverTab[5245]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:778
	// _ = "end of CoverTab[5208]"
//line /usr/local/go/src/net/dnsclient_unix.go:778
	_go_fuzz_dep_.CoverTab[5209]++
							if lastErr, ok := lastErr.(*DNSError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:779
		_go_fuzz_dep_.CoverTab[5298]++

//line /usr/local/go/src/net/dnsclient_unix.go:783
		lastErr.Name = name
//line /usr/local/go/src/net/dnsclient_unix.go:783
		// _ = "end of CoverTab[5298]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:784
		_go_fuzz_dep_.CoverTab[5299]++
//line /usr/local/go/src/net/dnsclient_unix.go:784
		// _ = "end of CoverTab[5299]"
//line /usr/local/go/src/net/dnsclient_unix.go:784
	}
//line /usr/local/go/src/net/dnsclient_unix.go:784
	// _ = "end of CoverTab[5209]"
//line /usr/local/go/src/net/dnsclient_unix.go:784
	_go_fuzz_dep_.CoverTab[5210]++
							sortByRFC6724(addrs)
							if len(addrs) == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[5300]++
//line /usr/local/go/src/net/dnsclient_unix.go:786
		return !(network == "CNAME" && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:786
			_go_fuzz_dep_.CoverTab[5301]++
//line /usr/local/go/src/net/dnsclient_unix.go:786
			return cname.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:786
			// _ = "end of CoverTab[5301]"
//line /usr/local/go/src/net/dnsclient_unix.go:786
		}())
//line /usr/local/go/src/net/dnsclient_unix.go:786
		// _ = "end of CoverTab[5300]"
//line /usr/local/go/src/net/dnsclient_unix.go:786
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[5302]++
								if order == hostLookupDNSFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:787
			_go_fuzz_dep_.CoverTab[5304]++
									var canonical string
									addrs, canonical = goLookupIPFiles(name)
									if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:790
				_go_fuzz_dep_.CoverTab[5305]++
										var err error
										cname, err = dnsmessage.NewName(canonical)
										if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:793
					_go_fuzz_dep_.CoverTab[5307]++
											return nil, dnsmessage.Name{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:794
					// _ = "end of CoverTab[5307]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:795
					_go_fuzz_dep_.CoverTab[5308]++
//line /usr/local/go/src/net/dnsclient_unix.go:795
					// _ = "end of CoverTab[5308]"
//line /usr/local/go/src/net/dnsclient_unix.go:795
				}
//line /usr/local/go/src/net/dnsclient_unix.go:795
				// _ = "end of CoverTab[5305]"
//line /usr/local/go/src/net/dnsclient_unix.go:795
				_go_fuzz_dep_.CoverTab[5306]++
										return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:796
				// _ = "end of CoverTab[5306]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:797
				_go_fuzz_dep_.CoverTab[5309]++
//line /usr/local/go/src/net/dnsclient_unix.go:797
				// _ = "end of CoverTab[5309]"
//line /usr/local/go/src/net/dnsclient_unix.go:797
			}
//line /usr/local/go/src/net/dnsclient_unix.go:797
			// _ = "end of CoverTab[5304]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:798
			_go_fuzz_dep_.CoverTab[5310]++
//line /usr/local/go/src/net/dnsclient_unix.go:798
			// _ = "end of CoverTab[5310]"
//line /usr/local/go/src/net/dnsclient_unix.go:798
		}
//line /usr/local/go/src/net/dnsclient_unix.go:798
		// _ = "end of CoverTab[5302]"
//line /usr/local/go/src/net/dnsclient_unix.go:798
		_go_fuzz_dep_.CoverTab[5303]++
								if lastErr != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:799
			_go_fuzz_dep_.CoverTab[5311]++
									return nil, dnsmessage.Name{}, lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:800
			// _ = "end of CoverTab[5311]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:801
			_go_fuzz_dep_.CoverTab[5312]++
//line /usr/local/go/src/net/dnsclient_unix.go:801
			// _ = "end of CoverTab[5312]"
//line /usr/local/go/src/net/dnsclient_unix.go:801
		}
//line /usr/local/go/src/net/dnsclient_unix.go:801
		// _ = "end of CoverTab[5303]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:802
		_go_fuzz_dep_.CoverTab[5313]++
//line /usr/local/go/src/net/dnsclient_unix.go:802
		// _ = "end of CoverTab[5313]"
//line /usr/local/go/src/net/dnsclient_unix.go:802
	}
//line /usr/local/go/src/net/dnsclient_unix.go:802
	// _ = "end of CoverTab[5210]"
//line /usr/local/go/src/net/dnsclient_unix.go:802
	_go_fuzz_dep_.CoverTab[5211]++
							return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:803
	// _ = "end of CoverTab[5211]"
}

// goLookupCNAME is the native Go (non-cgo) implementation of LookupCNAME.
func (r *Resolver) goLookupCNAME(ctx context.Context, host string, order hostLookupOrder, conf *dnsConfig) (string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:807
	_go_fuzz_dep_.CoverTab[5314]++
							_, cname, err := r.goLookupIPCNAMEOrder(ctx, "CNAME", host, order, conf)
							return cname.String(), err
//line /usr/local/go/src/net/dnsclient_unix.go:809
	// _ = "end of CoverTab[5314]"
}

// goLookupPTR is the native Go implementation of LookupAddr.
//line /usr/local/go/src/net/dnsclient_unix.go:812
// Used only if cgoLookupPTR refuses to handle the request (that is,
//line /usr/local/go/src/net/dnsclient_unix.go:812
// only if cgoLookupPTR is the stub in cgo_stub.go).
//line /usr/local/go/src/net/dnsclient_unix.go:812
// Normally we let cgo use the C library resolver instead of depending
//line /usr/local/go/src/net/dnsclient_unix.go:812
// on our lookup code, so that Go and C get the same answers.
//line /usr/local/go/src/net/dnsclient_unix.go:817
func (r *Resolver) goLookupPTR(ctx context.Context, addr string, conf *dnsConfig) ([]string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:817
	_go_fuzz_dep_.CoverTab[5315]++
							names := lookupStaticAddr(addr)
							if len(names) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:819
		_go_fuzz_dep_.CoverTab[5320]++
								return names, nil
//line /usr/local/go/src/net/dnsclient_unix.go:820
		// _ = "end of CoverTab[5320]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:821
		_go_fuzz_dep_.CoverTab[5321]++
//line /usr/local/go/src/net/dnsclient_unix.go:821
		// _ = "end of CoverTab[5321]"
//line /usr/local/go/src/net/dnsclient_unix.go:821
	}
//line /usr/local/go/src/net/dnsclient_unix.go:821
	// _ = "end of CoverTab[5315]"
//line /usr/local/go/src/net/dnsclient_unix.go:821
	_go_fuzz_dep_.CoverTab[5316]++
							arpa, err := reverseaddr(addr)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:823
		_go_fuzz_dep_.CoverTab[5322]++
								return nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:824
		// _ = "end of CoverTab[5322]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:825
		_go_fuzz_dep_.CoverTab[5323]++
//line /usr/local/go/src/net/dnsclient_unix.go:825
		// _ = "end of CoverTab[5323]"
//line /usr/local/go/src/net/dnsclient_unix.go:825
	}
//line /usr/local/go/src/net/dnsclient_unix.go:825
	// _ = "end of CoverTab[5316]"
//line /usr/local/go/src/net/dnsclient_unix.go:825
	_go_fuzz_dep_.CoverTab[5317]++
							p, server, err := r.lookup(ctx, arpa, dnsmessage.TypePTR, conf)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:827
		_go_fuzz_dep_.CoverTab[5324]++
								return nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:828
		// _ = "end of CoverTab[5324]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:829
		_go_fuzz_dep_.CoverTab[5325]++
//line /usr/local/go/src/net/dnsclient_unix.go:829
		// _ = "end of CoverTab[5325]"
//line /usr/local/go/src/net/dnsclient_unix.go:829
	}
//line /usr/local/go/src/net/dnsclient_unix.go:829
	// _ = "end of CoverTab[5317]"
//line /usr/local/go/src/net/dnsclient_unix.go:829
	_go_fuzz_dep_.CoverTab[5318]++
							var ptrs []string
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:831
		_go_fuzz_dep_.CoverTab[5326]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:833
			_go_fuzz_dep_.CoverTab[5331]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:834
			// _ = "end of CoverTab[5331]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:835
			_go_fuzz_dep_.CoverTab[5332]++
//line /usr/local/go/src/net/dnsclient_unix.go:835
			// _ = "end of CoverTab[5332]"
//line /usr/local/go/src/net/dnsclient_unix.go:835
		}
//line /usr/local/go/src/net/dnsclient_unix.go:835
		// _ = "end of CoverTab[5326]"
//line /usr/local/go/src/net/dnsclient_unix.go:835
		_go_fuzz_dep_.CoverTab[5327]++
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:836
			_go_fuzz_dep_.CoverTab[5333]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /usr/local/go/src/net/dnsclient_unix.go:841
			// _ = "end of CoverTab[5333]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:842
			_go_fuzz_dep_.CoverTab[5334]++
//line /usr/local/go/src/net/dnsclient_unix.go:842
			// _ = "end of CoverTab[5334]"
//line /usr/local/go/src/net/dnsclient_unix.go:842
		}
//line /usr/local/go/src/net/dnsclient_unix.go:842
		// _ = "end of CoverTab[5327]"
//line /usr/local/go/src/net/dnsclient_unix.go:842
		_go_fuzz_dep_.CoverTab[5328]++
								if h.Type != dnsmessage.TypePTR {
//line /usr/local/go/src/net/dnsclient_unix.go:843
			_go_fuzz_dep_.CoverTab[5335]++
									err := p.SkipAnswer()
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:845
				_go_fuzz_dep_.CoverTab[5337]++
										return nil, &DNSError{
					Err:	"cannot marshal DNS message",
					Name:	addr,
					Server:	server,
				}
//line /usr/local/go/src/net/dnsclient_unix.go:850
				// _ = "end of CoverTab[5337]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:851
				_go_fuzz_dep_.CoverTab[5338]++
//line /usr/local/go/src/net/dnsclient_unix.go:851
				// _ = "end of CoverTab[5338]"
//line /usr/local/go/src/net/dnsclient_unix.go:851
			}
//line /usr/local/go/src/net/dnsclient_unix.go:851
			// _ = "end of CoverTab[5335]"
//line /usr/local/go/src/net/dnsclient_unix.go:851
			_go_fuzz_dep_.CoverTab[5336]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:852
			// _ = "end of CoverTab[5336]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:853
			_go_fuzz_dep_.CoverTab[5339]++
//line /usr/local/go/src/net/dnsclient_unix.go:853
			// _ = "end of CoverTab[5339]"
//line /usr/local/go/src/net/dnsclient_unix.go:853
		}
//line /usr/local/go/src/net/dnsclient_unix.go:853
		// _ = "end of CoverTab[5328]"
//line /usr/local/go/src/net/dnsclient_unix.go:853
		_go_fuzz_dep_.CoverTab[5329]++
								ptr, err := p.PTRResource()
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:855
			_go_fuzz_dep_.CoverTab[5340]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /usr/local/go/src/net/dnsclient_unix.go:860
			// _ = "end of CoverTab[5340]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:861
			_go_fuzz_dep_.CoverTab[5341]++
//line /usr/local/go/src/net/dnsclient_unix.go:861
			// _ = "end of CoverTab[5341]"
//line /usr/local/go/src/net/dnsclient_unix.go:861
		}
//line /usr/local/go/src/net/dnsclient_unix.go:861
		// _ = "end of CoverTab[5329]"
//line /usr/local/go/src/net/dnsclient_unix.go:861
		_go_fuzz_dep_.CoverTab[5330]++
								ptrs = append(ptrs, ptr.PTR.String())
//line /usr/local/go/src/net/dnsclient_unix.go:862
		// _ = "end of CoverTab[5330]"

	}
//line /usr/local/go/src/net/dnsclient_unix.go:864
	// _ = "end of CoverTab[5318]"
//line /usr/local/go/src/net/dnsclient_unix.go:864
	_go_fuzz_dep_.CoverTab[5319]++
							return ptrs, nil
//line /usr/local/go/src/net/dnsclient_unix.go:865
	// _ = "end of CoverTab[5319]"
}

//line /usr/local/go/src/net/dnsclient_unix.go:866
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dnsclient_unix.go:866
var _ = _go_fuzz_dep_.CoverTab
