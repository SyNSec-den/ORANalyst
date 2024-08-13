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
	_go_fuzz_dep_.CoverTab[13331]++
							id = uint16(randInt())
							b := dnsmessage.NewBuilder(make([]byte, 2, 514), dnsmessage.Header{ID: id, RecursionDesired: true, AuthenticData: ad})
							if err := b.StartQuestions(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:58
		_go_fuzz_dep_.CoverTab[13338]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:59
		// _ = "end of CoverTab[13338]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:60
		_go_fuzz_dep_.CoverTab[13339]++
//line /usr/local/go/src/net/dnsclient_unix.go:60
		// _ = "end of CoverTab[13339]"
//line /usr/local/go/src/net/dnsclient_unix.go:60
	}
//line /usr/local/go/src/net/dnsclient_unix.go:60
	// _ = "end of CoverTab[13331]"
//line /usr/local/go/src/net/dnsclient_unix.go:60
	_go_fuzz_dep_.CoverTab[13332]++
							if err := b.Question(q); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:61
		_go_fuzz_dep_.CoverTab[13340]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:62
		// _ = "end of CoverTab[13340]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:63
		_go_fuzz_dep_.CoverTab[13341]++
//line /usr/local/go/src/net/dnsclient_unix.go:63
		// _ = "end of CoverTab[13341]"
//line /usr/local/go/src/net/dnsclient_unix.go:63
	}
//line /usr/local/go/src/net/dnsclient_unix.go:63
	// _ = "end of CoverTab[13332]"
//line /usr/local/go/src/net/dnsclient_unix.go:63
	_go_fuzz_dep_.CoverTab[13333]++

//line /usr/local/go/src/net/dnsclient_unix.go:66
	if err := b.StartAdditionals(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:66
		_go_fuzz_dep_.CoverTab[13342]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:67
		// _ = "end of CoverTab[13342]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:68
		_go_fuzz_dep_.CoverTab[13343]++
//line /usr/local/go/src/net/dnsclient_unix.go:68
		// _ = "end of CoverTab[13343]"
//line /usr/local/go/src/net/dnsclient_unix.go:68
	}
//line /usr/local/go/src/net/dnsclient_unix.go:68
	// _ = "end of CoverTab[13333]"
//line /usr/local/go/src/net/dnsclient_unix.go:68
	_go_fuzz_dep_.CoverTab[13334]++
							var rh dnsmessage.ResourceHeader
							if err := rh.SetEDNS0(maxDNSPacketSize, dnsmessage.RCodeSuccess, false); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:70
		_go_fuzz_dep_.CoverTab[13344]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:71
		// _ = "end of CoverTab[13344]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:72
		_go_fuzz_dep_.CoverTab[13345]++
//line /usr/local/go/src/net/dnsclient_unix.go:72
		// _ = "end of CoverTab[13345]"
//line /usr/local/go/src/net/dnsclient_unix.go:72
	}
//line /usr/local/go/src/net/dnsclient_unix.go:72
	// _ = "end of CoverTab[13334]"
//line /usr/local/go/src/net/dnsclient_unix.go:72
	_go_fuzz_dep_.CoverTab[13335]++
							if err := b.OPTResource(rh, dnsmessage.OPTResource{}); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:73
		_go_fuzz_dep_.CoverTab[13346]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:74
		// _ = "end of CoverTab[13346]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:75
		_go_fuzz_dep_.CoverTab[13347]++
//line /usr/local/go/src/net/dnsclient_unix.go:75
		// _ = "end of CoverTab[13347]"
//line /usr/local/go/src/net/dnsclient_unix.go:75
	}
//line /usr/local/go/src/net/dnsclient_unix.go:75
	// _ = "end of CoverTab[13335]"
//line /usr/local/go/src/net/dnsclient_unix.go:75
	_go_fuzz_dep_.CoverTab[13336]++

							tcpReq, err = b.Finish()
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:78
		_go_fuzz_dep_.CoverTab[13348]++
								return 0, nil, nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:79
		// _ = "end of CoverTab[13348]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:80
		_go_fuzz_dep_.CoverTab[13349]++
//line /usr/local/go/src/net/dnsclient_unix.go:80
		// _ = "end of CoverTab[13349]"
//line /usr/local/go/src/net/dnsclient_unix.go:80
	}
//line /usr/local/go/src/net/dnsclient_unix.go:80
	// _ = "end of CoverTab[13336]"
//line /usr/local/go/src/net/dnsclient_unix.go:80
	_go_fuzz_dep_.CoverTab[13337]++
							udpReq = tcpReq[2:]
							l := len(tcpReq) - 2
							tcpReq[0] = byte(l >> 8)
							tcpReq[1] = byte(l)
							return id, udpReq, tcpReq, nil
//line /usr/local/go/src/net/dnsclient_unix.go:85
	// _ = "end of CoverTab[13337]"
}

func checkResponse(reqID uint16, reqQues dnsmessage.Question, respHdr dnsmessage.Header, respQues dnsmessage.Question) bool {
//line /usr/local/go/src/net/dnsclient_unix.go:88
	_go_fuzz_dep_.CoverTab[13350]++
							if !respHdr.Response {
//line /usr/local/go/src/net/dnsclient_unix.go:89
		_go_fuzz_dep_.CoverTab[13354]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:90
		// _ = "end of CoverTab[13354]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:91
		_go_fuzz_dep_.CoverTab[13355]++
//line /usr/local/go/src/net/dnsclient_unix.go:91
		// _ = "end of CoverTab[13355]"
//line /usr/local/go/src/net/dnsclient_unix.go:91
	}
//line /usr/local/go/src/net/dnsclient_unix.go:91
	// _ = "end of CoverTab[13350]"
//line /usr/local/go/src/net/dnsclient_unix.go:91
	_go_fuzz_dep_.CoverTab[13351]++
							if reqID != respHdr.ID {
//line /usr/local/go/src/net/dnsclient_unix.go:92
		_go_fuzz_dep_.CoverTab[13356]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:93
		// _ = "end of CoverTab[13356]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:94
		_go_fuzz_dep_.CoverTab[13357]++
//line /usr/local/go/src/net/dnsclient_unix.go:94
		// _ = "end of CoverTab[13357]"
//line /usr/local/go/src/net/dnsclient_unix.go:94
	}
//line /usr/local/go/src/net/dnsclient_unix.go:94
	// _ = "end of CoverTab[13351]"
//line /usr/local/go/src/net/dnsclient_unix.go:94
	_go_fuzz_dep_.CoverTab[13352]++
							if reqQues.Type != respQues.Type || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[13358]++
//line /usr/local/go/src/net/dnsclient_unix.go:95
		return reqQues.Class != respQues.Class
//line /usr/local/go/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[13358]"
//line /usr/local/go/src/net/dnsclient_unix.go:95
	}() || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[13359]++
//line /usr/local/go/src/net/dnsclient_unix.go:95
		return !equalASCIIName(reqQues.Name, respQues.Name)
//line /usr/local/go/src/net/dnsclient_unix.go:95
		// _ = "end of CoverTab[13359]"
//line /usr/local/go/src/net/dnsclient_unix.go:95
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:95
		_go_fuzz_dep_.CoverTab[13360]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:96
		// _ = "end of CoverTab[13360]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:97
		_go_fuzz_dep_.CoverTab[13361]++
//line /usr/local/go/src/net/dnsclient_unix.go:97
		// _ = "end of CoverTab[13361]"
//line /usr/local/go/src/net/dnsclient_unix.go:97
	}
//line /usr/local/go/src/net/dnsclient_unix.go:97
	// _ = "end of CoverTab[13352]"
//line /usr/local/go/src/net/dnsclient_unix.go:97
	_go_fuzz_dep_.CoverTab[13353]++
							return true
//line /usr/local/go/src/net/dnsclient_unix.go:98
	// _ = "end of CoverTab[13353]"
}

func dnsPacketRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:101
	_go_fuzz_dep_.CoverTab[13362]++
							if _, err := c.Write(b); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:102
		_go_fuzz_dep_.CoverTab[13364]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:103
		// _ = "end of CoverTab[13364]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:104
		_go_fuzz_dep_.CoverTab[13365]++
//line /usr/local/go/src/net/dnsclient_unix.go:104
		// _ = "end of CoverTab[13365]"
//line /usr/local/go/src/net/dnsclient_unix.go:104
	}
//line /usr/local/go/src/net/dnsclient_unix.go:104
	// _ = "end of CoverTab[13362]"
//line /usr/local/go/src/net/dnsclient_unix.go:104
	_go_fuzz_dep_.CoverTab[13363]++

							b = make([]byte, maxDNSPacketSize)
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:107
		_go_fuzz_dep_.CoverTab[13366]++
								n, err := c.Read(b)
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:109
			_go_fuzz_dep_.CoverTab[13370]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:110
			// _ = "end of CoverTab[13370]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:111
			_go_fuzz_dep_.CoverTab[13371]++
//line /usr/local/go/src/net/dnsclient_unix.go:111
			// _ = "end of CoverTab[13371]"
//line /usr/local/go/src/net/dnsclient_unix.go:111
		}
//line /usr/local/go/src/net/dnsclient_unix.go:111
		// _ = "end of CoverTab[13366]"
//line /usr/local/go/src/net/dnsclient_unix.go:111
		_go_fuzz_dep_.CoverTab[13367]++
								var p dnsmessage.Parser

//line /usr/local/go/src/net/dnsclient_unix.go:116
		h, err := p.Start(b[:n])
		if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:117
			_go_fuzz_dep_.CoverTab[13372]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:118
			// _ = "end of CoverTab[13372]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:119
			_go_fuzz_dep_.CoverTab[13373]++
//line /usr/local/go/src/net/dnsclient_unix.go:119
			// _ = "end of CoverTab[13373]"
//line /usr/local/go/src/net/dnsclient_unix.go:119
		}
//line /usr/local/go/src/net/dnsclient_unix.go:119
		// _ = "end of CoverTab[13367]"
//line /usr/local/go/src/net/dnsclient_unix.go:119
		_go_fuzz_dep_.CoverTab[13368]++
								q, err := p.Question()
								if err != nil || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[13374]++
//line /usr/local/go/src/net/dnsclient_unix.go:121
			return !checkResponse(id, query, h, q)
//line /usr/local/go/src/net/dnsclient_unix.go:121
			// _ = "end of CoverTab[13374]"
//line /usr/local/go/src/net/dnsclient_unix.go:121
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:121
			_go_fuzz_dep_.CoverTab[13375]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:122
			// _ = "end of CoverTab[13375]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:123
			_go_fuzz_dep_.CoverTab[13376]++
//line /usr/local/go/src/net/dnsclient_unix.go:123
			// _ = "end of CoverTab[13376]"
//line /usr/local/go/src/net/dnsclient_unix.go:123
		}
//line /usr/local/go/src/net/dnsclient_unix.go:123
		// _ = "end of CoverTab[13368]"
//line /usr/local/go/src/net/dnsclient_unix.go:123
		_go_fuzz_dep_.CoverTab[13369]++
								return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:124
		// _ = "end of CoverTab[13369]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:125
	// _ = "end of CoverTab[13363]"
}

func dnsStreamRoundTrip(c Conn, id uint16, query dnsmessage.Question, b []byte) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:128
	_go_fuzz_dep_.CoverTab[13377]++
							if _, err := c.Write(b); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:129
		_go_fuzz_dep_.CoverTab[13385]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:130
		// _ = "end of CoverTab[13385]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:131
		_go_fuzz_dep_.CoverTab[13386]++
//line /usr/local/go/src/net/dnsclient_unix.go:131
		// _ = "end of CoverTab[13386]"
//line /usr/local/go/src/net/dnsclient_unix.go:131
	}
//line /usr/local/go/src/net/dnsclient_unix.go:131
	// _ = "end of CoverTab[13377]"
//line /usr/local/go/src/net/dnsclient_unix.go:131
	_go_fuzz_dep_.CoverTab[13378]++

							b = make([]byte, 1280)
							if _, err := io.ReadFull(c, b[:2]); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:134
		_go_fuzz_dep_.CoverTab[13387]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:135
		// _ = "end of CoverTab[13387]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:136
		_go_fuzz_dep_.CoverTab[13388]++
//line /usr/local/go/src/net/dnsclient_unix.go:136
		// _ = "end of CoverTab[13388]"
//line /usr/local/go/src/net/dnsclient_unix.go:136
	}
//line /usr/local/go/src/net/dnsclient_unix.go:136
	// _ = "end of CoverTab[13378]"
//line /usr/local/go/src/net/dnsclient_unix.go:136
	_go_fuzz_dep_.CoverTab[13379]++
							l := int(b[0])<<8 | int(b[1])
							if l > len(b) {
//line /usr/local/go/src/net/dnsclient_unix.go:138
		_go_fuzz_dep_.CoverTab[13389]++
								b = make([]byte, l)
//line /usr/local/go/src/net/dnsclient_unix.go:139
		// _ = "end of CoverTab[13389]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:140
		_go_fuzz_dep_.CoverTab[13390]++
//line /usr/local/go/src/net/dnsclient_unix.go:140
		// _ = "end of CoverTab[13390]"
//line /usr/local/go/src/net/dnsclient_unix.go:140
	}
//line /usr/local/go/src/net/dnsclient_unix.go:140
	// _ = "end of CoverTab[13379]"
//line /usr/local/go/src/net/dnsclient_unix.go:140
	_go_fuzz_dep_.CoverTab[13380]++
							n, err := io.ReadFull(c, b[:l])
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:142
		_go_fuzz_dep_.CoverTab[13391]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:143
		// _ = "end of CoverTab[13391]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:144
		_go_fuzz_dep_.CoverTab[13392]++
//line /usr/local/go/src/net/dnsclient_unix.go:144
		// _ = "end of CoverTab[13392]"
//line /usr/local/go/src/net/dnsclient_unix.go:144
	}
//line /usr/local/go/src/net/dnsclient_unix.go:144
	// _ = "end of CoverTab[13380]"
//line /usr/local/go/src/net/dnsclient_unix.go:144
	_go_fuzz_dep_.CoverTab[13381]++
							var p dnsmessage.Parser
							h, err := p.Start(b[:n])
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:147
		_go_fuzz_dep_.CoverTab[13393]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:148
		// _ = "end of CoverTab[13393]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:149
		_go_fuzz_dep_.CoverTab[13394]++
//line /usr/local/go/src/net/dnsclient_unix.go:149
		// _ = "end of CoverTab[13394]"
//line /usr/local/go/src/net/dnsclient_unix.go:149
	}
//line /usr/local/go/src/net/dnsclient_unix.go:149
	// _ = "end of CoverTab[13381]"
//line /usr/local/go/src/net/dnsclient_unix.go:149
	_go_fuzz_dep_.CoverTab[13382]++
							q, err := p.Question()
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:151
		_go_fuzz_dep_.CoverTab[13395]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:152
		// _ = "end of CoverTab[13395]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:153
		_go_fuzz_dep_.CoverTab[13396]++
//line /usr/local/go/src/net/dnsclient_unix.go:153
		// _ = "end of CoverTab[13396]"
//line /usr/local/go/src/net/dnsclient_unix.go:153
	}
//line /usr/local/go/src/net/dnsclient_unix.go:153
	// _ = "end of CoverTab[13382]"
//line /usr/local/go/src/net/dnsclient_unix.go:153
	_go_fuzz_dep_.CoverTab[13383]++
							if !checkResponse(id, query, h, q) {
//line /usr/local/go/src/net/dnsclient_unix.go:154
		_go_fuzz_dep_.CoverTab[13397]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /usr/local/go/src/net/dnsclient_unix.go:155
		// _ = "end of CoverTab[13397]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:156
		_go_fuzz_dep_.CoverTab[13398]++
//line /usr/local/go/src/net/dnsclient_unix.go:156
		// _ = "end of CoverTab[13398]"
//line /usr/local/go/src/net/dnsclient_unix.go:156
	}
//line /usr/local/go/src/net/dnsclient_unix.go:156
	// _ = "end of CoverTab[13383]"
//line /usr/local/go/src/net/dnsclient_unix.go:156
	_go_fuzz_dep_.CoverTab[13384]++
							return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:157
	// _ = "end of CoverTab[13384]"
}

// exchange sends a query on the connection and hopes for a response.
func (r *Resolver) exchange(ctx context.Context, server string, q dnsmessage.Question, timeout time.Duration, useTCP, ad bool) (dnsmessage.Parser, dnsmessage.Header, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:161
	_go_fuzz_dep_.CoverTab[13399]++
							q.Class = dnsmessage.ClassINET
							id, udpReq, tcpReq, err := newRequest(q, ad)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:164
		_go_fuzz_dep_.CoverTab[13403]++
								return dnsmessage.Parser{}, dnsmessage.Header{}, errCannotMarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:165
		// _ = "end of CoverTab[13403]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:166
		_go_fuzz_dep_.CoverTab[13404]++
//line /usr/local/go/src/net/dnsclient_unix.go:166
		// _ = "end of CoverTab[13404]"
//line /usr/local/go/src/net/dnsclient_unix.go:166
	}
//line /usr/local/go/src/net/dnsclient_unix.go:166
	// _ = "end of CoverTab[13399]"
//line /usr/local/go/src/net/dnsclient_unix.go:166
	_go_fuzz_dep_.CoverTab[13400]++
							var networks []string
							if useTCP {
//line /usr/local/go/src/net/dnsclient_unix.go:168
		_go_fuzz_dep_.CoverTab[13405]++
								networks = []string{"tcp"}
//line /usr/local/go/src/net/dnsclient_unix.go:169
		// _ = "end of CoverTab[13405]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:170
		_go_fuzz_dep_.CoverTab[13406]++
								networks = []string{"udp", "tcp"}
//line /usr/local/go/src/net/dnsclient_unix.go:171
		// _ = "end of CoverTab[13406]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:172
	// _ = "end of CoverTab[13400]"
//line /usr/local/go/src/net/dnsclient_unix.go:172
	_go_fuzz_dep_.CoverTab[13401]++
							for _, network := range networks {
//line /usr/local/go/src/net/dnsclient_unix.go:173
		_go_fuzz_dep_.CoverTab[13407]++
								ctx, cancel := context.WithDeadline(ctx, time.Now().Add(timeout))
								defer cancel()

								c, err := r.dial(ctx, network, server)
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:178
			_go_fuzz_dep_.CoverTab[13414]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:179
			// _ = "end of CoverTab[13414]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:180
			_go_fuzz_dep_.CoverTab[13415]++
//line /usr/local/go/src/net/dnsclient_unix.go:180
			// _ = "end of CoverTab[13415]"
//line /usr/local/go/src/net/dnsclient_unix.go:180
		}
//line /usr/local/go/src/net/dnsclient_unix.go:180
		// _ = "end of CoverTab[13407]"
//line /usr/local/go/src/net/dnsclient_unix.go:180
		_go_fuzz_dep_.CoverTab[13408]++
								if d, ok := ctx.Deadline(); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[13416]++
//line /usr/local/go/src/net/dnsclient_unix.go:181
			return !d.IsZero()
//line /usr/local/go/src/net/dnsclient_unix.go:181
			// _ = "end of CoverTab[13416]"
//line /usr/local/go/src/net/dnsclient_unix.go:181
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:181
			_go_fuzz_dep_.CoverTab[13417]++
									c.SetDeadline(d)
//line /usr/local/go/src/net/dnsclient_unix.go:182
			// _ = "end of CoverTab[13417]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:183
			_go_fuzz_dep_.CoverTab[13418]++
//line /usr/local/go/src/net/dnsclient_unix.go:183
			// _ = "end of CoverTab[13418]"
//line /usr/local/go/src/net/dnsclient_unix.go:183
		}
//line /usr/local/go/src/net/dnsclient_unix.go:183
		// _ = "end of CoverTab[13408]"
//line /usr/local/go/src/net/dnsclient_unix.go:183
		_go_fuzz_dep_.CoverTab[13409]++
								var p dnsmessage.Parser
								var h dnsmessage.Header
								if _, ok := c.(PacketConn); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:186
			_go_fuzz_dep_.CoverTab[13419]++
									p, h, err = dnsPacketRoundTrip(c, id, q, udpReq)
//line /usr/local/go/src/net/dnsclient_unix.go:187
			// _ = "end of CoverTab[13419]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:188
			_go_fuzz_dep_.CoverTab[13420]++
									p, h, err = dnsStreamRoundTrip(c, id, q, tcpReq)
//line /usr/local/go/src/net/dnsclient_unix.go:189
			// _ = "end of CoverTab[13420]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:190
		// _ = "end of CoverTab[13409]"
//line /usr/local/go/src/net/dnsclient_unix.go:190
		_go_fuzz_dep_.CoverTab[13410]++
								c.Close()
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:192
			_go_fuzz_dep_.CoverTab[13421]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, mapErr(err)
//line /usr/local/go/src/net/dnsclient_unix.go:193
			// _ = "end of CoverTab[13421]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:194
			_go_fuzz_dep_.CoverTab[13422]++
//line /usr/local/go/src/net/dnsclient_unix.go:194
			// _ = "end of CoverTab[13422]"
//line /usr/local/go/src/net/dnsclient_unix.go:194
		}
//line /usr/local/go/src/net/dnsclient_unix.go:194
		// _ = "end of CoverTab[13410]"
//line /usr/local/go/src/net/dnsclient_unix.go:194
		_go_fuzz_dep_.CoverTab[13411]++
								if err := p.SkipQuestion(); err != dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:195
			_go_fuzz_dep_.CoverTab[13423]++
									return dnsmessage.Parser{}, dnsmessage.Header{}, errInvalidDNSResponse
//line /usr/local/go/src/net/dnsclient_unix.go:196
			// _ = "end of CoverTab[13423]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:197
			_go_fuzz_dep_.CoverTab[13424]++
//line /usr/local/go/src/net/dnsclient_unix.go:197
			// _ = "end of CoverTab[13424]"
//line /usr/local/go/src/net/dnsclient_unix.go:197
		}
//line /usr/local/go/src/net/dnsclient_unix.go:197
		// _ = "end of CoverTab[13411]"
//line /usr/local/go/src/net/dnsclient_unix.go:197
		_go_fuzz_dep_.CoverTab[13412]++
								if h.Truncated {
//line /usr/local/go/src/net/dnsclient_unix.go:198
			_go_fuzz_dep_.CoverTab[13425]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:199
			// _ = "end of CoverTab[13425]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:200
			_go_fuzz_dep_.CoverTab[13426]++
//line /usr/local/go/src/net/dnsclient_unix.go:200
			// _ = "end of CoverTab[13426]"
//line /usr/local/go/src/net/dnsclient_unix.go:200
		}
//line /usr/local/go/src/net/dnsclient_unix.go:200
		// _ = "end of CoverTab[13412]"
//line /usr/local/go/src/net/dnsclient_unix.go:200
		_go_fuzz_dep_.CoverTab[13413]++
								return p, h, nil
//line /usr/local/go/src/net/dnsclient_unix.go:201
		// _ = "end of CoverTab[13413]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:202
	// _ = "end of CoverTab[13401]"
//line /usr/local/go/src/net/dnsclient_unix.go:202
	_go_fuzz_dep_.CoverTab[13402]++
							return dnsmessage.Parser{}, dnsmessage.Header{}, errNoAnswerFromDNSServer
//line /usr/local/go/src/net/dnsclient_unix.go:203
	// _ = "end of CoverTab[13402]"
}

// checkHeader performs basic sanity checks on the header.
func checkHeader(p *dnsmessage.Parser, h dnsmessage.Header) error {
//line /usr/local/go/src/net/dnsclient_unix.go:207
	_go_fuzz_dep_.CoverTab[13427]++
							if h.RCode == dnsmessage.RCodeNameError {
//line /usr/local/go/src/net/dnsclient_unix.go:208
		_go_fuzz_dep_.CoverTab[13432]++
								return errNoSuchHost
//line /usr/local/go/src/net/dnsclient_unix.go:209
		// _ = "end of CoverTab[13432]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:210
		_go_fuzz_dep_.CoverTab[13433]++
//line /usr/local/go/src/net/dnsclient_unix.go:210
		// _ = "end of CoverTab[13433]"
//line /usr/local/go/src/net/dnsclient_unix.go:210
	}
//line /usr/local/go/src/net/dnsclient_unix.go:210
	// _ = "end of CoverTab[13427]"
//line /usr/local/go/src/net/dnsclient_unix.go:210
	_go_fuzz_dep_.CoverTab[13428]++

							_, err := p.AnswerHeader()
							if err != nil && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[13434]++
//line /usr/local/go/src/net/dnsclient_unix.go:213
		return err != dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:213
		// _ = "end of CoverTab[13434]"
//line /usr/local/go/src/net/dnsclient_unix.go:213
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:213
		_go_fuzz_dep_.CoverTab[13435]++
								return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:214
		// _ = "end of CoverTab[13435]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:215
		_go_fuzz_dep_.CoverTab[13436]++
//line /usr/local/go/src/net/dnsclient_unix.go:215
		// _ = "end of CoverTab[13436]"
//line /usr/local/go/src/net/dnsclient_unix.go:215
	}
//line /usr/local/go/src/net/dnsclient_unix.go:215
	// _ = "end of CoverTab[13428]"
//line /usr/local/go/src/net/dnsclient_unix.go:215
	_go_fuzz_dep_.CoverTab[13429]++

//line /usr/local/go/src/net/dnsclient_unix.go:219
	if h.RCode == dnsmessage.RCodeSuccess && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[13437]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return !h.Authoritative
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[13437]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[13438]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return !h.RecursionAvailable
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[13438]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[13439]++
//line /usr/local/go/src/net/dnsclient_unix.go:219
		return err == dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:219
		// _ = "end of CoverTab[13439]"
//line /usr/local/go/src/net/dnsclient_unix.go:219
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:219
		_go_fuzz_dep_.CoverTab[13440]++
								return errLameReferral
//line /usr/local/go/src/net/dnsclient_unix.go:220
		// _ = "end of CoverTab[13440]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:221
		_go_fuzz_dep_.CoverTab[13441]++
//line /usr/local/go/src/net/dnsclient_unix.go:221
		// _ = "end of CoverTab[13441]"
//line /usr/local/go/src/net/dnsclient_unix.go:221
	}
//line /usr/local/go/src/net/dnsclient_unix.go:221
	// _ = "end of CoverTab[13429]"
//line /usr/local/go/src/net/dnsclient_unix.go:221
	_go_fuzz_dep_.CoverTab[13430]++

							if h.RCode != dnsmessage.RCodeSuccess && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[13442]++
//line /usr/local/go/src/net/dnsclient_unix.go:223
		return h.RCode != dnsmessage.RCodeNameError
//line /usr/local/go/src/net/dnsclient_unix.go:223
		// _ = "end of CoverTab[13442]"
//line /usr/local/go/src/net/dnsclient_unix.go:223
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:223
		_go_fuzz_dep_.CoverTab[13443]++

//line /usr/local/go/src/net/dnsclient_unix.go:229
		if h.RCode == dnsmessage.RCodeServerFailure {
//line /usr/local/go/src/net/dnsclient_unix.go:229
			_go_fuzz_dep_.CoverTab[13445]++
									return errServerTemporarilyMisbehaving
//line /usr/local/go/src/net/dnsclient_unix.go:230
			// _ = "end of CoverTab[13445]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:231
			_go_fuzz_dep_.CoverTab[13446]++
//line /usr/local/go/src/net/dnsclient_unix.go:231
			// _ = "end of CoverTab[13446]"
//line /usr/local/go/src/net/dnsclient_unix.go:231
		}
//line /usr/local/go/src/net/dnsclient_unix.go:231
		// _ = "end of CoverTab[13443]"
//line /usr/local/go/src/net/dnsclient_unix.go:231
		_go_fuzz_dep_.CoverTab[13444]++
								return errServerMisbehaving
//line /usr/local/go/src/net/dnsclient_unix.go:232
		// _ = "end of CoverTab[13444]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:233
		_go_fuzz_dep_.CoverTab[13447]++
//line /usr/local/go/src/net/dnsclient_unix.go:233
		// _ = "end of CoverTab[13447]"
//line /usr/local/go/src/net/dnsclient_unix.go:233
	}
//line /usr/local/go/src/net/dnsclient_unix.go:233
	// _ = "end of CoverTab[13430]"
//line /usr/local/go/src/net/dnsclient_unix.go:233
	_go_fuzz_dep_.CoverTab[13431]++

							return nil
//line /usr/local/go/src/net/dnsclient_unix.go:235
	// _ = "end of CoverTab[13431]"
}

func skipToAnswer(p *dnsmessage.Parser, qtype dnsmessage.Type) error {
//line /usr/local/go/src/net/dnsclient_unix.go:238
	_go_fuzz_dep_.CoverTab[13448]++
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:239
		_go_fuzz_dep_.CoverTab[13449]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:241
			_go_fuzz_dep_.CoverTab[13453]++
									return errNoSuchHost
//line /usr/local/go/src/net/dnsclient_unix.go:242
			// _ = "end of CoverTab[13453]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:243
			_go_fuzz_dep_.CoverTab[13454]++
//line /usr/local/go/src/net/dnsclient_unix.go:243
			// _ = "end of CoverTab[13454]"
//line /usr/local/go/src/net/dnsclient_unix.go:243
		}
//line /usr/local/go/src/net/dnsclient_unix.go:243
		// _ = "end of CoverTab[13449]"
//line /usr/local/go/src/net/dnsclient_unix.go:243
		_go_fuzz_dep_.CoverTab[13450]++
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:244
			_go_fuzz_dep_.CoverTab[13455]++
									return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:245
			// _ = "end of CoverTab[13455]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:246
			_go_fuzz_dep_.CoverTab[13456]++
//line /usr/local/go/src/net/dnsclient_unix.go:246
			// _ = "end of CoverTab[13456]"
//line /usr/local/go/src/net/dnsclient_unix.go:246
		}
//line /usr/local/go/src/net/dnsclient_unix.go:246
		// _ = "end of CoverTab[13450]"
//line /usr/local/go/src/net/dnsclient_unix.go:246
		_go_fuzz_dep_.CoverTab[13451]++
								if h.Type == qtype {
//line /usr/local/go/src/net/dnsclient_unix.go:247
			_go_fuzz_dep_.CoverTab[13457]++
									return nil
//line /usr/local/go/src/net/dnsclient_unix.go:248
			// _ = "end of CoverTab[13457]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:249
			_go_fuzz_dep_.CoverTab[13458]++
//line /usr/local/go/src/net/dnsclient_unix.go:249
			// _ = "end of CoverTab[13458]"
//line /usr/local/go/src/net/dnsclient_unix.go:249
		}
//line /usr/local/go/src/net/dnsclient_unix.go:249
		// _ = "end of CoverTab[13451]"
//line /usr/local/go/src/net/dnsclient_unix.go:249
		_go_fuzz_dep_.CoverTab[13452]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:250
			_go_fuzz_dep_.CoverTab[13459]++
									return errCannotUnmarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:251
			// _ = "end of CoverTab[13459]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:252
			_go_fuzz_dep_.CoverTab[13460]++
//line /usr/local/go/src/net/dnsclient_unix.go:252
			// _ = "end of CoverTab[13460]"
//line /usr/local/go/src/net/dnsclient_unix.go:252
		}
//line /usr/local/go/src/net/dnsclient_unix.go:252
		// _ = "end of CoverTab[13452]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:253
	// _ = "end of CoverTab[13448]"
}

// Do a lookup for a single name, which must be rooted
//line /usr/local/go/src/net/dnsclient_unix.go:256
// (otherwise answer will not find the answers).
//line /usr/local/go/src/net/dnsclient_unix.go:258
func (r *Resolver) tryOneName(ctx context.Context, cfg *dnsConfig, name string, qtype dnsmessage.Type) (dnsmessage.Parser, string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:258
	_go_fuzz_dep_.CoverTab[13461]++
							var lastErr error
							serverOffset := cfg.serverOffset()
							sLen := uint32(len(cfg.servers))

							n, err := dnsmessage.NewName(name)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:264
		_go_fuzz_dep_.CoverTab[13464]++
								return dnsmessage.Parser{}, "", errCannotMarshalDNSMessage
//line /usr/local/go/src/net/dnsclient_unix.go:265
		// _ = "end of CoverTab[13464]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:266
		_go_fuzz_dep_.CoverTab[13465]++
//line /usr/local/go/src/net/dnsclient_unix.go:266
		// _ = "end of CoverTab[13465]"
//line /usr/local/go/src/net/dnsclient_unix.go:266
	}
//line /usr/local/go/src/net/dnsclient_unix.go:266
	// _ = "end of CoverTab[13461]"
//line /usr/local/go/src/net/dnsclient_unix.go:266
	_go_fuzz_dep_.CoverTab[13462]++
							q := dnsmessage.Question{
		Name:	n,
		Type:	qtype,
		Class:	dnsmessage.ClassINET,
	}

	for i := 0; i < cfg.attempts; i++ {
//line /usr/local/go/src/net/dnsclient_unix.go:273
		_go_fuzz_dep_.CoverTab[13466]++
								for j := uint32(0); j < sLen; j++ {
//line /usr/local/go/src/net/dnsclient_unix.go:274
			_go_fuzz_dep_.CoverTab[13467]++
									server := cfg.servers[(serverOffset+j)%sLen]

									p, h, err := r.exchange(ctx, server, q, cfg.timeout, cfg.useTCP, cfg.trustAD)
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:278
				_go_fuzz_dep_.CoverTab[13471]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if nerr, ok := err.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[13474]++
//line /usr/local/go/src/net/dnsclient_unix.go:284
					return nerr.Timeout()
//line /usr/local/go/src/net/dnsclient_unix.go:284
					// _ = "end of CoverTab[13474]"
//line /usr/local/go/src/net/dnsclient_unix.go:284
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:284
					_go_fuzz_dep_.CoverTab[13475]++
											dnsErr.IsTimeout = true
//line /usr/local/go/src/net/dnsclient_unix.go:285
					// _ = "end of CoverTab[13475]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:286
					_go_fuzz_dep_.CoverTab[13476]++
//line /usr/local/go/src/net/dnsclient_unix.go:286
					// _ = "end of CoverTab[13476]"
//line /usr/local/go/src/net/dnsclient_unix.go:286
				}
//line /usr/local/go/src/net/dnsclient_unix.go:286
				// _ = "end of CoverTab[13471]"
//line /usr/local/go/src/net/dnsclient_unix.go:286
				_go_fuzz_dep_.CoverTab[13472]++

//line /usr/local/go/src/net/dnsclient_unix.go:289
				if _, ok := err.(*OpError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:289
					_go_fuzz_dep_.CoverTab[13477]++
											dnsErr.IsTemporary = true
//line /usr/local/go/src/net/dnsclient_unix.go:290
					// _ = "end of CoverTab[13477]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:291
					_go_fuzz_dep_.CoverTab[13478]++
//line /usr/local/go/src/net/dnsclient_unix.go:291
					// _ = "end of CoverTab[13478]"
//line /usr/local/go/src/net/dnsclient_unix.go:291
				}
//line /usr/local/go/src/net/dnsclient_unix.go:291
				// _ = "end of CoverTab[13472]"
//line /usr/local/go/src/net/dnsclient_unix.go:291
				_go_fuzz_dep_.CoverTab[13473]++
										lastErr = dnsErr
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:293
				// _ = "end of CoverTab[13473]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:294
				_go_fuzz_dep_.CoverTab[13479]++
//line /usr/local/go/src/net/dnsclient_unix.go:294
				// _ = "end of CoverTab[13479]"
//line /usr/local/go/src/net/dnsclient_unix.go:294
			}
//line /usr/local/go/src/net/dnsclient_unix.go:294
			// _ = "end of CoverTab[13467]"
//line /usr/local/go/src/net/dnsclient_unix.go:294
			_go_fuzz_dep_.CoverTab[13468]++

									if err := checkHeader(&p, h); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:296
				_go_fuzz_dep_.CoverTab[13480]++
										dnsErr := &DNSError{
					Err:	err.Error(),
					Name:	name,
					Server:	server,
				}
				if err == errServerTemporarilyMisbehaving {
//line /usr/local/go/src/net/dnsclient_unix.go:302
					_go_fuzz_dep_.CoverTab[13483]++
											dnsErr.IsTemporary = true
//line /usr/local/go/src/net/dnsclient_unix.go:303
					// _ = "end of CoverTab[13483]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:304
					_go_fuzz_dep_.CoverTab[13484]++
//line /usr/local/go/src/net/dnsclient_unix.go:304
					// _ = "end of CoverTab[13484]"
//line /usr/local/go/src/net/dnsclient_unix.go:304
				}
//line /usr/local/go/src/net/dnsclient_unix.go:304
				// _ = "end of CoverTab[13480]"
//line /usr/local/go/src/net/dnsclient_unix.go:304
				_go_fuzz_dep_.CoverTab[13481]++
										if err == errNoSuchHost {
//line /usr/local/go/src/net/dnsclient_unix.go:305
					_go_fuzz_dep_.CoverTab[13485]++

//line /usr/local/go/src/net/dnsclient_unix.go:309
					dnsErr.IsNotFound = true
											return p, server, dnsErr
//line /usr/local/go/src/net/dnsclient_unix.go:310
					// _ = "end of CoverTab[13485]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:311
					_go_fuzz_dep_.CoverTab[13486]++
//line /usr/local/go/src/net/dnsclient_unix.go:311
					// _ = "end of CoverTab[13486]"
//line /usr/local/go/src/net/dnsclient_unix.go:311
				}
//line /usr/local/go/src/net/dnsclient_unix.go:311
				// _ = "end of CoverTab[13481]"
//line /usr/local/go/src/net/dnsclient_unix.go:311
				_go_fuzz_dep_.CoverTab[13482]++
										lastErr = dnsErr
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:313
				// _ = "end of CoverTab[13482]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:314
				_go_fuzz_dep_.CoverTab[13487]++
//line /usr/local/go/src/net/dnsclient_unix.go:314
				// _ = "end of CoverTab[13487]"
//line /usr/local/go/src/net/dnsclient_unix.go:314
			}
//line /usr/local/go/src/net/dnsclient_unix.go:314
			// _ = "end of CoverTab[13468]"
//line /usr/local/go/src/net/dnsclient_unix.go:314
			_go_fuzz_dep_.CoverTab[13469]++

									err = skipToAnswer(&p, qtype)
									if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:317
				_go_fuzz_dep_.CoverTab[13488]++
										return p, server, nil
//line /usr/local/go/src/net/dnsclient_unix.go:318
				// _ = "end of CoverTab[13488]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:319
				_go_fuzz_dep_.CoverTab[13489]++
//line /usr/local/go/src/net/dnsclient_unix.go:319
				// _ = "end of CoverTab[13489]"
//line /usr/local/go/src/net/dnsclient_unix.go:319
			}
//line /usr/local/go/src/net/dnsclient_unix.go:319
			// _ = "end of CoverTab[13469]"
//line /usr/local/go/src/net/dnsclient_unix.go:319
			_go_fuzz_dep_.CoverTab[13470]++
									lastErr = &DNSError{
				Err:	err.Error(),
				Name:	name,
				Server:	server,
			}
			if err == errNoSuchHost {
//line /usr/local/go/src/net/dnsclient_unix.go:325
				_go_fuzz_dep_.CoverTab[13490]++

//line /usr/local/go/src/net/dnsclient_unix.go:329
				lastErr.(*DNSError).IsNotFound = true
										return p, server, lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:330
				// _ = "end of CoverTab[13490]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:331
				_go_fuzz_dep_.CoverTab[13491]++
//line /usr/local/go/src/net/dnsclient_unix.go:331
				// _ = "end of CoverTab[13491]"
//line /usr/local/go/src/net/dnsclient_unix.go:331
			}
//line /usr/local/go/src/net/dnsclient_unix.go:331
			// _ = "end of CoverTab[13470]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:332
		// _ = "end of CoverTab[13466]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:333
	// _ = "end of CoverTab[13462]"
//line /usr/local/go/src/net/dnsclient_unix.go:333
	_go_fuzz_dep_.CoverTab[13463]++
							return dnsmessage.Parser{}, "", lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:334
	// _ = "end of CoverTab[13463]"
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
	_go_fuzz_dep_.CoverTab[13492]++
							resolvConf.tryUpdate("/etc/resolv.conf")
							return resolvConf.dnsConfig.Load()
//line /usr/local/go/src/net/dnsclient_unix.go:353
	// _ = "end of CoverTab[13492]"
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
	_go_fuzz_dep_.CoverTab[13493]++
							conf.initOnce.Do(conf.init)

							if conf.dnsConfig.Load().noReload {
//line /usr/local/go/src/net/dnsclient_unix.go:374
		_go_fuzz_dep_.CoverTab[13498]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:375
		// _ = "end of CoverTab[13498]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:376
		_go_fuzz_dep_.CoverTab[13499]++
//line /usr/local/go/src/net/dnsclient_unix.go:376
		// _ = "end of CoverTab[13499]"
//line /usr/local/go/src/net/dnsclient_unix.go:376
	}
//line /usr/local/go/src/net/dnsclient_unix.go:376
	// _ = "end of CoverTab[13493]"
//line /usr/local/go/src/net/dnsclient_unix.go:376
	_go_fuzz_dep_.CoverTab[13494]++

//line /usr/local/go/src/net/dnsclient_unix.go:379
	if !conf.tryAcquireSema() {
//line /usr/local/go/src/net/dnsclient_unix.go:379
		_go_fuzz_dep_.CoverTab[13500]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:380
		// _ = "end of CoverTab[13500]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:381
		_go_fuzz_dep_.CoverTab[13501]++
//line /usr/local/go/src/net/dnsclient_unix.go:381
		// _ = "end of CoverTab[13501]"
//line /usr/local/go/src/net/dnsclient_unix.go:381
	}
//line /usr/local/go/src/net/dnsclient_unix.go:381
	// _ = "end of CoverTab[13494]"
//line /usr/local/go/src/net/dnsclient_unix.go:381
	_go_fuzz_dep_.CoverTab[13495]++
							defer conf.releaseSema()

							now := time.Now()
							if conf.lastChecked.After(now.Add(-5 * time.Second)) {
//line /usr/local/go/src/net/dnsclient_unix.go:385
		_go_fuzz_dep_.CoverTab[13502]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:386
		// _ = "end of CoverTab[13502]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:387
		_go_fuzz_dep_.CoverTab[13503]++
//line /usr/local/go/src/net/dnsclient_unix.go:387
		// _ = "end of CoverTab[13503]"
//line /usr/local/go/src/net/dnsclient_unix.go:387
	}
//line /usr/local/go/src/net/dnsclient_unix.go:387
	// _ = "end of CoverTab[13495]"
//line /usr/local/go/src/net/dnsclient_unix.go:387
	_go_fuzz_dep_.CoverTab[13496]++
							conf.lastChecked = now

							switch runtime.GOOS {
	case "windows":
//line /usr/local/go/src/net/dnsclient_unix.go:391
		_go_fuzz_dep_.CoverTab[13504]++
//line /usr/local/go/src/net/dnsclient_unix.go:391
		// _ = "end of CoverTab[13504]"

//line /usr/local/go/src/net/dnsclient_unix.go:397
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:397
		_go_fuzz_dep_.CoverTab[13505]++
								var mtime time.Time
								if fi, err := os.Stat(name); err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:399
			_go_fuzz_dep_.CoverTab[13507]++
									mtime = fi.ModTime()
//line /usr/local/go/src/net/dnsclient_unix.go:400
			// _ = "end of CoverTab[13507]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:401
			_go_fuzz_dep_.CoverTab[13508]++
//line /usr/local/go/src/net/dnsclient_unix.go:401
			// _ = "end of CoverTab[13508]"
//line /usr/local/go/src/net/dnsclient_unix.go:401
		}
//line /usr/local/go/src/net/dnsclient_unix.go:401
		// _ = "end of CoverTab[13505]"
//line /usr/local/go/src/net/dnsclient_unix.go:401
		_go_fuzz_dep_.CoverTab[13506]++
								if mtime.Equal(conf.dnsConfig.Load().mtime) {
//line /usr/local/go/src/net/dnsclient_unix.go:402
			_go_fuzz_dep_.CoverTab[13509]++
									return
//line /usr/local/go/src/net/dnsclient_unix.go:403
			// _ = "end of CoverTab[13509]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:404
			_go_fuzz_dep_.CoverTab[13510]++
//line /usr/local/go/src/net/dnsclient_unix.go:404
			// _ = "end of CoverTab[13510]"
//line /usr/local/go/src/net/dnsclient_unix.go:404
		}
//line /usr/local/go/src/net/dnsclient_unix.go:404
		// _ = "end of CoverTab[13506]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:405
	// _ = "end of CoverTab[13496]"
//line /usr/local/go/src/net/dnsclient_unix.go:405
	_go_fuzz_dep_.CoverTab[13497]++

							dnsConf := dnsReadConfig(name)
							conf.dnsConfig.Store(dnsConf)
//line /usr/local/go/src/net/dnsclient_unix.go:408
	// _ = "end of CoverTab[13497]"
}

func (conf *resolverConfig) tryAcquireSema() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:411
	_go_fuzz_dep_.CoverTab[13511]++
							select {
	case conf.ch <- struct{}{}:
//line /usr/local/go/src/net/dnsclient_unix.go:413
		_go_fuzz_dep_.CoverTab[13512]++
								return true
//line /usr/local/go/src/net/dnsclient_unix.go:414
		// _ = "end of CoverTab[13512]"
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:415
		_go_fuzz_dep_.CoverTab[13513]++
								return false
//line /usr/local/go/src/net/dnsclient_unix.go:416
		// _ = "end of CoverTab[13513]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:417
	// _ = "end of CoverTab[13511]"
}

func (conf *resolverConfig) releaseSema() {
//line /usr/local/go/src/net/dnsclient_unix.go:420
	_go_fuzz_dep_.CoverTab[13514]++
							<-conf.ch
//line /usr/local/go/src/net/dnsclient_unix.go:421
	// _ = "end of CoverTab[13514]"
}

func (r *Resolver) lookup(ctx context.Context, name string, qtype dnsmessage.Type, conf *dnsConfig) (dnsmessage.Parser, string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:424
	_go_fuzz_dep_.CoverTab[13515]++
							if !isDomainName(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:425
		_go_fuzz_dep_.CoverTab[13521]++

//line /usr/local/go/src/net/dnsclient_unix.go:431
		return dnsmessage.Parser{}, "", &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:431
		// _ = "end of CoverTab[13521]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:432
		_go_fuzz_dep_.CoverTab[13522]++
//line /usr/local/go/src/net/dnsclient_unix.go:432
		// _ = "end of CoverTab[13522]"
//line /usr/local/go/src/net/dnsclient_unix.go:432
	}
//line /usr/local/go/src/net/dnsclient_unix.go:432
	// _ = "end of CoverTab[13515]"
//line /usr/local/go/src/net/dnsclient_unix.go:432
	_go_fuzz_dep_.CoverTab[13516]++

							if conf == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:434
		_go_fuzz_dep_.CoverTab[13523]++
								conf = getSystemDNSConfig()
//line /usr/local/go/src/net/dnsclient_unix.go:435
		// _ = "end of CoverTab[13523]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:436
		_go_fuzz_dep_.CoverTab[13524]++
//line /usr/local/go/src/net/dnsclient_unix.go:436
		// _ = "end of CoverTab[13524]"
//line /usr/local/go/src/net/dnsclient_unix.go:436
	}
//line /usr/local/go/src/net/dnsclient_unix.go:436
	// _ = "end of CoverTab[13516]"
//line /usr/local/go/src/net/dnsclient_unix.go:436
	_go_fuzz_dep_.CoverTab[13517]++

							var (
		p	dnsmessage.Parser
		server	string
		err	error
	)
	for _, fqdn := range conf.nameList(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:443
		_go_fuzz_dep_.CoverTab[13525]++
								p, server, err = r.tryOneName(ctx, conf, fqdn, qtype)
								if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:445
			_go_fuzz_dep_.CoverTab[13527]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:446
			// _ = "end of CoverTab[13527]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:447
			_go_fuzz_dep_.CoverTab[13528]++
//line /usr/local/go/src/net/dnsclient_unix.go:447
			// _ = "end of CoverTab[13528]"
//line /usr/local/go/src/net/dnsclient_unix.go:447
		}
//line /usr/local/go/src/net/dnsclient_unix.go:447
		// _ = "end of CoverTab[13525]"
//line /usr/local/go/src/net/dnsclient_unix.go:447
		_go_fuzz_dep_.CoverTab[13526]++
								if nerr, ok := err.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[13529]++
//line /usr/local/go/src/net/dnsclient_unix.go:448
			return nerr.Temporary()
//line /usr/local/go/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[13529]"
//line /usr/local/go/src/net/dnsclient_unix.go:448
		}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[13530]++
//line /usr/local/go/src/net/dnsclient_unix.go:448
			return r.strictErrors()
//line /usr/local/go/src/net/dnsclient_unix.go:448
			// _ = "end of CoverTab[13530]"
//line /usr/local/go/src/net/dnsclient_unix.go:448
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:448
			_go_fuzz_dep_.CoverTab[13531]++

//line /usr/local/go/src/net/dnsclient_unix.go:451
			break
//line /usr/local/go/src/net/dnsclient_unix.go:451
			// _ = "end of CoverTab[13531]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:452
			_go_fuzz_dep_.CoverTab[13532]++
//line /usr/local/go/src/net/dnsclient_unix.go:452
			// _ = "end of CoverTab[13532]"
//line /usr/local/go/src/net/dnsclient_unix.go:452
		}
//line /usr/local/go/src/net/dnsclient_unix.go:452
		// _ = "end of CoverTab[13526]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:453
	// _ = "end of CoverTab[13517]"
//line /usr/local/go/src/net/dnsclient_unix.go:453
	_go_fuzz_dep_.CoverTab[13518]++
							if err == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:454
		_go_fuzz_dep_.CoverTab[13533]++
								return p, server, nil
//line /usr/local/go/src/net/dnsclient_unix.go:455
		// _ = "end of CoverTab[13533]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:456
		_go_fuzz_dep_.CoverTab[13534]++
//line /usr/local/go/src/net/dnsclient_unix.go:456
		// _ = "end of CoverTab[13534]"
//line /usr/local/go/src/net/dnsclient_unix.go:456
	}
//line /usr/local/go/src/net/dnsclient_unix.go:456
	// _ = "end of CoverTab[13518]"
//line /usr/local/go/src/net/dnsclient_unix.go:456
	_go_fuzz_dep_.CoverTab[13519]++
							if err, ok := err.(*DNSError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:457
		_go_fuzz_dep_.CoverTab[13535]++

//line /usr/local/go/src/net/dnsclient_unix.go:461
		err.Name = name
//line /usr/local/go/src/net/dnsclient_unix.go:461
		// _ = "end of CoverTab[13535]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:462
		_go_fuzz_dep_.CoverTab[13536]++
//line /usr/local/go/src/net/dnsclient_unix.go:462
		// _ = "end of CoverTab[13536]"
//line /usr/local/go/src/net/dnsclient_unix.go:462
	}
//line /usr/local/go/src/net/dnsclient_unix.go:462
	// _ = "end of CoverTab[13519]"
//line /usr/local/go/src/net/dnsclient_unix.go:462
	_go_fuzz_dep_.CoverTab[13520]++
							return dnsmessage.Parser{}, "", err
//line /usr/local/go/src/net/dnsclient_unix.go:463
	// _ = "end of CoverTab[13520]"
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
	_go_fuzz_dep_.CoverTab[13537]++
							if name == "" {
//line /usr/local/go/src/net/dnsclient_unix.go:471
		_go_fuzz_dep_.CoverTab[13540]++
								return true
//line /usr/local/go/src/net/dnsclient_unix.go:472
		// _ = "end of CoverTab[13540]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:473
		_go_fuzz_dep_.CoverTab[13541]++
//line /usr/local/go/src/net/dnsclient_unix.go:473
		// _ = "end of CoverTab[13541]"
//line /usr/local/go/src/net/dnsclient_unix.go:473
	}
//line /usr/local/go/src/net/dnsclient_unix.go:473
	// _ = "end of CoverTab[13537]"
//line /usr/local/go/src/net/dnsclient_unix.go:473
	_go_fuzz_dep_.CoverTab[13538]++
							if name[len(name)-1] == '.' {
//line /usr/local/go/src/net/dnsclient_unix.go:474
		_go_fuzz_dep_.CoverTab[13542]++
								name = name[:len(name)-1]
//line /usr/local/go/src/net/dnsclient_unix.go:475
		// _ = "end of CoverTab[13542]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:476
		_go_fuzz_dep_.CoverTab[13543]++
//line /usr/local/go/src/net/dnsclient_unix.go:476
		// _ = "end of CoverTab[13543]"
//line /usr/local/go/src/net/dnsclient_unix.go:476
	}
//line /usr/local/go/src/net/dnsclient_unix.go:476
	// _ = "end of CoverTab[13538]"
//line /usr/local/go/src/net/dnsclient_unix.go:476
	_go_fuzz_dep_.CoverTab[13539]++
							return stringsHasSuffixFold(name, ".onion")
//line /usr/local/go/src/net/dnsclient_unix.go:477
	// _ = "end of CoverTab[13539]"
}

// nameList returns a list of names for sequential DNS queries.
func (conf *dnsConfig) nameList(name string) []string {
//line /usr/local/go/src/net/dnsclient_unix.go:481
	_go_fuzz_dep_.CoverTab[13544]++
							if avoidDNS(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:482
		_go_fuzz_dep_.CoverTab[13551]++
								return nil
//line /usr/local/go/src/net/dnsclient_unix.go:483
		// _ = "end of CoverTab[13551]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:484
		_go_fuzz_dep_.CoverTab[13552]++
//line /usr/local/go/src/net/dnsclient_unix.go:484
		// _ = "end of CoverTab[13552]"
//line /usr/local/go/src/net/dnsclient_unix.go:484
	}
//line /usr/local/go/src/net/dnsclient_unix.go:484
	// _ = "end of CoverTab[13544]"
//line /usr/local/go/src/net/dnsclient_unix.go:484
	_go_fuzz_dep_.CoverTab[13545]++

//line /usr/local/go/src/net/dnsclient_unix.go:487
	l := len(name)
	rooted := l > 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:488
		_go_fuzz_dep_.CoverTab[13553]++
//line /usr/local/go/src/net/dnsclient_unix.go:488
		return name[l-1] == '.'
//line /usr/local/go/src/net/dnsclient_unix.go:488
		// _ = "end of CoverTab[13553]"
//line /usr/local/go/src/net/dnsclient_unix.go:488
	}()
							if l > 254 || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[13554]++
//line /usr/local/go/src/net/dnsclient_unix.go:489
		return l == 254 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:489
			_go_fuzz_dep_.CoverTab[13555]++
//line /usr/local/go/src/net/dnsclient_unix.go:489
			return !rooted
//line /usr/local/go/src/net/dnsclient_unix.go:489
			// _ = "end of CoverTab[13555]"
//line /usr/local/go/src/net/dnsclient_unix.go:489
		}()
//line /usr/local/go/src/net/dnsclient_unix.go:489
		// _ = "end of CoverTab[13554]"
//line /usr/local/go/src/net/dnsclient_unix.go:489
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:489
		_go_fuzz_dep_.CoverTab[13556]++
								return nil
//line /usr/local/go/src/net/dnsclient_unix.go:490
		// _ = "end of CoverTab[13556]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:491
		_go_fuzz_dep_.CoverTab[13557]++
//line /usr/local/go/src/net/dnsclient_unix.go:491
		// _ = "end of CoverTab[13557]"
//line /usr/local/go/src/net/dnsclient_unix.go:491
	}
//line /usr/local/go/src/net/dnsclient_unix.go:491
	// _ = "end of CoverTab[13545]"
//line /usr/local/go/src/net/dnsclient_unix.go:491
	_go_fuzz_dep_.CoverTab[13546]++

//line /usr/local/go/src/net/dnsclient_unix.go:494
	if rooted {
//line /usr/local/go/src/net/dnsclient_unix.go:494
		_go_fuzz_dep_.CoverTab[13558]++
								return []string{name}
//line /usr/local/go/src/net/dnsclient_unix.go:495
		// _ = "end of CoverTab[13558]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:496
		_go_fuzz_dep_.CoverTab[13559]++
//line /usr/local/go/src/net/dnsclient_unix.go:496
		// _ = "end of CoverTab[13559]"
//line /usr/local/go/src/net/dnsclient_unix.go:496
	}
//line /usr/local/go/src/net/dnsclient_unix.go:496
	// _ = "end of CoverTab[13546]"
//line /usr/local/go/src/net/dnsclient_unix.go:496
	_go_fuzz_dep_.CoverTab[13547]++

							hasNdots := count(name, '.') >= conf.ndots
							name += "."
							l++

//line /usr/local/go/src/net/dnsclient_unix.go:503
	names := make([]string, 0, 1+len(conf.search))

	if hasNdots {
//line /usr/local/go/src/net/dnsclient_unix.go:505
		_go_fuzz_dep_.CoverTab[13560]++
								names = append(names, name)
//line /usr/local/go/src/net/dnsclient_unix.go:506
		// _ = "end of CoverTab[13560]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:507
		_go_fuzz_dep_.CoverTab[13561]++
//line /usr/local/go/src/net/dnsclient_unix.go:507
		// _ = "end of CoverTab[13561]"
//line /usr/local/go/src/net/dnsclient_unix.go:507
	}
//line /usr/local/go/src/net/dnsclient_unix.go:507
	// _ = "end of CoverTab[13547]"
//line /usr/local/go/src/net/dnsclient_unix.go:507
	_go_fuzz_dep_.CoverTab[13548]++

							for _, suffix := range conf.search {
//line /usr/local/go/src/net/dnsclient_unix.go:509
		_go_fuzz_dep_.CoverTab[13562]++
								if l+len(suffix) <= 254 {
//line /usr/local/go/src/net/dnsclient_unix.go:510
			_go_fuzz_dep_.CoverTab[13563]++
									names = append(names, name+suffix)
//line /usr/local/go/src/net/dnsclient_unix.go:511
			// _ = "end of CoverTab[13563]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:512
			_go_fuzz_dep_.CoverTab[13564]++
//line /usr/local/go/src/net/dnsclient_unix.go:512
			// _ = "end of CoverTab[13564]"
//line /usr/local/go/src/net/dnsclient_unix.go:512
		}
//line /usr/local/go/src/net/dnsclient_unix.go:512
		// _ = "end of CoverTab[13562]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:513
	// _ = "end of CoverTab[13548]"
//line /usr/local/go/src/net/dnsclient_unix.go:513
	_go_fuzz_dep_.CoverTab[13549]++

							if !hasNdots {
//line /usr/local/go/src/net/dnsclient_unix.go:515
		_go_fuzz_dep_.CoverTab[13565]++
								names = append(names, name)
//line /usr/local/go/src/net/dnsclient_unix.go:516
		// _ = "end of CoverTab[13565]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:517
		_go_fuzz_dep_.CoverTab[13566]++
//line /usr/local/go/src/net/dnsclient_unix.go:517
		// _ = "end of CoverTab[13566]"
//line /usr/local/go/src/net/dnsclient_unix.go:517
	}
//line /usr/local/go/src/net/dnsclient_unix.go:517
	// _ = "end of CoverTab[13549]"
//line /usr/local/go/src/net/dnsclient_unix.go:517
	_go_fuzz_dep_.CoverTab[13550]++
							return names
//line /usr/local/go/src/net/dnsclient_unix.go:518
	// _ = "end of CoverTab[13550]"
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
	_go_fuzz_dep_.CoverTab[13567]++
							if s, ok := lookupOrderName[o]; ok {
//line /usr/local/go/src/net/dnsclient_unix.go:544
		_go_fuzz_dep_.CoverTab[13569]++
								return s
//line /usr/local/go/src/net/dnsclient_unix.go:545
		// _ = "end of CoverTab[13569]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:546
		_go_fuzz_dep_.CoverTab[13570]++
//line /usr/local/go/src/net/dnsclient_unix.go:546
		// _ = "end of CoverTab[13570]"
//line /usr/local/go/src/net/dnsclient_unix.go:546
	}
//line /usr/local/go/src/net/dnsclient_unix.go:546
	// _ = "end of CoverTab[13567]"
//line /usr/local/go/src/net/dnsclient_unix.go:546
	_go_fuzz_dep_.CoverTab[13568]++
							return "hostLookupOrder=" + itoa.Itoa(int(o)) + "??"
//line /usr/local/go/src/net/dnsclient_unix.go:547
	// _ = "end of CoverTab[13568]"
}

func (r *Resolver) goLookupHostOrder(ctx context.Context, name string, order hostLookupOrder, conf *dnsConfig) (addrs []string, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:550
	_go_fuzz_dep_.CoverTab[13571]++
							if order == hostLookupFilesDNS || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[13575]++
//line /usr/local/go/src/net/dnsclient_unix.go:551
		return order == hostLookupFiles
//line /usr/local/go/src/net/dnsclient_unix.go:551
		// _ = "end of CoverTab[13575]"
//line /usr/local/go/src/net/dnsclient_unix.go:551
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:551
		_go_fuzz_dep_.CoverTab[13576]++

								addrs, _ = lookupStaticHost(name)
								if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:554
			_go_fuzz_dep_.CoverTab[13578]++
									return
//line /usr/local/go/src/net/dnsclient_unix.go:555
			// _ = "end of CoverTab[13578]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:556
			_go_fuzz_dep_.CoverTab[13579]++
//line /usr/local/go/src/net/dnsclient_unix.go:556
			// _ = "end of CoverTab[13579]"
//line /usr/local/go/src/net/dnsclient_unix.go:556
		}
//line /usr/local/go/src/net/dnsclient_unix.go:556
		// _ = "end of CoverTab[13576]"
//line /usr/local/go/src/net/dnsclient_unix.go:556
		_go_fuzz_dep_.CoverTab[13577]++

								if order == hostLookupFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:558
			_go_fuzz_dep_.CoverTab[13580]++
									return nil, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:559
			// _ = "end of CoverTab[13580]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:560
			_go_fuzz_dep_.CoverTab[13581]++
//line /usr/local/go/src/net/dnsclient_unix.go:560
			// _ = "end of CoverTab[13581]"
//line /usr/local/go/src/net/dnsclient_unix.go:560
		}
//line /usr/local/go/src/net/dnsclient_unix.go:560
		// _ = "end of CoverTab[13577]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:561
		_go_fuzz_dep_.CoverTab[13582]++
//line /usr/local/go/src/net/dnsclient_unix.go:561
		// _ = "end of CoverTab[13582]"
//line /usr/local/go/src/net/dnsclient_unix.go:561
	}
//line /usr/local/go/src/net/dnsclient_unix.go:561
	// _ = "end of CoverTab[13571]"
//line /usr/local/go/src/net/dnsclient_unix.go:561
	_go_fuzz_dep_.CoverTab[13572]++
							ips, _, err := r.goLookupIPCNAMEOrder(ctx, "ip", name, order, conf)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:563
		_go_fuzz_dep_.CoverTab[13583]++
								return
//line /usr/local/go/src/net/dnsclient_unix.go:564
		// _ = "end of CoverTab[13583]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:565
		_go_fuzz_dep_.CoverTab[13584]++
//line /usr/local/go/src/net/dnsclient_unix.go:565
		// _ = "end of CoverTab[13584]"
//line /usr/local/go/src/net/dnsclient_unix.go:565
	}
//line /usr/local/go/src/net/dnsclient_unix.go:565
	// _ = "end of CoverTab[13572]"
//line /usr/local/go/src/net/dnsclient_unix.go:565
	_go_fuzz_dep_.CoverTab[13573]++
							addrs = make([]string, 0, len(ips))
							for _, ip := range ips {
//line /usr/local/go/src/net/dnsclient_unix.go:567
		_go_fuzz_dep_.CoverTab[13585]++
								addrs = append(addrs, ip.String())
//line /usr/local/go/src/net/dnsclient_unix.go:568
		// _ = "end of CoverTab[13585]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:569
	// _ = "end of CoverTab[13573]"
//line /usr/local/go/src/net/dnsclient_unix.go:569
	_go_fuzz_dep_.CoverTab[13574]++
							return
//line /usr/local/go/src/net/dnsclient_unix.go:570
	// _ = "end of CoverTab[13574]"
}

// lookup entries from /etc/hosts
func goLookupIPFiles(name string) (addrs []IPAddr, canonical string) {
//line /usr/local/go/src/net/dnsclient_unix.go:574
	_go_fuzz_dep_.CoverTab[13586]++
							addr, canonical := lookupStaticHost(name)
							for _, haddr := range addr {
//line /usr/local/go/src/net/dnsclient_unix.go:576
		_go_fuzz_dep_.CoverTab[13588]++
								haddr, zone := splitHostZone(haddr)
								if ip := ParseIP(haddr); ip != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:578
			_go_fuzz_dep_.CoverTab[13589]++
									addr := IPAddr{IP: ip, Zone: zone}
									addrs = append(addrs, addr)
//line /usr/local/go/src/net/dnsclient_unix.go:580
			// _ = "end of CoverTab[13589]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:581
			_go_fuzz_dep_.CoverTab[13590]++
//line /usr/local/go/src/net/dnsclient_unix.go:581
			// _ = "end of CoverTab[13590]"
//line /usr/local/go/src/net/dnsclient_unix.go:581
		}
//line /usr/local/go/src/net/dnsclient_unix.go:581
		// _ = "end of CoverTab[13588]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:582
	// _ = "end of CoverTab[13586]"
//line /usr/local/go/src/net/dnsclient_unix.go:582
	_go_fuzz_dep_.CoverTab[13587]++
							sortByRFC6724(addrs)
							return addrs, canonical
//line /usr/local/go/src/net/dnsclient_unix.go:584
	// _ = "end of CoverTab[13587]"
}

// goLookupIP is the native Go implementation of LookupIP.
//line /usr/local/go/src/net/dnsclient_unix.go:587
// The libc versions are in cgo_*.go.
//line /usr/local/go/src/net/dnsclient_unix.go:589
func (r *Resolver) goLookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:589
	_go_fuzz_dep_.CoverTab[13591]++
							order, conf := systemConf().hostLookupOrder(r, host)
							addrs, _, err = r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
							return
//line /usr/local/go/src/net/dnsclient_unix.go:592
	// _ = "end of CoverTab[13591]"
}

func (r *Resolver) goLookupIPCNAMEOrder(ctx context.Context, network, name string, order hostLookupOrder, conf *dnsConfig) (addrs []IPAddr, cname dnsmessage.Name, err error) {
//line /usr/local/go/src/net/dnsclient_unix.go:595
	_go_fuzz_dep_.CoverTab[13592]++
							if order == hostLookupFilesDNS || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[13602]++
//line /usr/local/go/src/net/dnsclient_unix.go:596
		return order == hostLookupFiles
//line /usr/local/go/src/net/dnsclient_unix.go:596
		// _ = "end of CoverTab[13602]"
//line /usr/local/go/src/net/dnsclient_unix.go:596
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:596
		_go_fuzz_dep_.CoverTab[13603]++
								var canonical string
								addrs, canonical = goLookupIPFiles(name)

								if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:600
			_go_fuzz_dep_.CoverTab[13605]++
									var err error
									cname, err = dnsmessage.NewName(canonical)
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:603
				_go_fuzz_dep_.CoverTab[13607]++
										return nil, dnsmessage.Name{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:604
				// _ = "end of CoverTab[13607]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:605
				_go_fuzz_dep_.CoverTab[13608]++
//line /usr/local/go/src/net/dnsclient_unix.go:605
				// _ = "end of CoverTab[13608]"
//line /usr/local/go/src/net/dnsclient_unix.go:605
			}
//line /usr/local/go/src/net/dnsclient_unix.go:605
			// _ = "end of CoverTab[13605]"
//line /usr/local/go/src/net/dnsclient_unix.go:605
			_go_fuzz_dep_.CoverTab[13606]++
									return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:606
			// _ = "end of CoverTab[13606]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:607
			_go_fuzz_dep_.CoverTab[13609]++
//line /usr/local/go/src/net/dnsclient_unix.go:607
			// _ = "end of CoverTab[13609]"
//line /usr/local/go/src/net/dnsclient_unix.go:607
		}
//line /usr/local/go/src/net/dnsclient_unix.go:607
		// _ = "end of CoverTab[13603]"
//line /usr/local/go/src/net/dnsclient_unix.go:607
		_go_fuzz_dep_.CoverTab[13604]++

								if order == hostLookupFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:609
			_go_fuzz_dep_.CoverTab[13610]++
									return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:610
			// _ = "end of CoverTab[13610]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:611
			_go_fuzz_dep_.CoverTab[13611]++
//line /usr/local/go/src/net/dnsclient_unix.go:611
			// _ = "end of CoverTab[13611]"
//line /usr/local/go/src/net/dnsclient_unix.go:611
		}
//line /usr/local/go/src/net/dnsclient_unix.go:611
		// _ = "end of CoverTab[13604]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:612
		_go_fuzz_dep_.CoverTab[13612]++
//line /usr/local/go/src/net/dnsclient_unix.go:612
		// _ = "end of CoverTab[13612]"
//line /usr/local/go/src/net/dnsclient_unix.go:612
	}
//line /usr/local/go/src/net/dnsclient_unix.go:612
	// _ = "end of CoverTab[13592]"
//line /usr/local/go/src/net/dnsclient_unix.go:612
	_go_fuzz_dep_.CoverTab[13593]++

							if !isDomainName(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:614
		_go_fuzz_dep_.CoverTab[13613]++

								return nil, dnsmessage.Name{}, &DNSError{Err: errNoSuchHost.Error(), Name: name, IsNotFound: true}
//line /usr/local/go/src/net/dnsclient_unix.go:616
		// _ = "end of CoverTab[13613]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:617
		_go_fuzz_dep_.CoverTab[13614]++
//line /usr/local/go/src/net/dnsclient_unix.go:617
		// _ = "end of CoverTab[13614]"
//line /usr/local/go/src/net/dnsclient_unix.go:617
	}
//line /usr/local/go/src/net/dnsclient_unix.go:617
	// _ = "end of CoverTab[13593]"
//line /usr/local/go/src/net/dnsclient_unix.go:617
	_go_fuzz_dep_.CoverTab[13594]++
							type result struct {
		p	dnsmessage.Parser
		server	string
		error
	}

	if conf == nil {
//line /usr/local/go/src/net/dnsclient_unix.go:624
		_go_fuzz_dep_.CoverTab[13615]++
								conf = getSystemDNSConfig()
//line /usr/local/go/src/net/dnsclient_unix.go:625
		// _ = "end of CoverTab[13615]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:626
		_go_fuzz_dep_.CoverTab[13616]++
//line /usr/local/go/src/net/dnsclient_unix.go:626
		// _ = "end of CoverTab[13616]"
//line /usr/local/go/src/net/dnsclient_unix.go:626
	}
//line /usr/local/go/src/net/dnsclient_unix.go:626
	// _ = "end of CoverTab[13594]"
//line /usr/local/go/src/net/dnsclient_unix.go:626
	_go_fuzz_dep_.CoverTab[13595]++

							lane := make(chan result, 1)
							qtypes := []dnsmessage.Type{dnsmessage.TypeA, dnsmessage.TypeAAAA}
							if network == "CNAME" {
//line /usr/local/go/src/net/dnsclient_unix.go:630
		_go_fuzz_dep_.CoverTab[13617]++
								qtypes = append(qtypes, dnsmessage.TypeCNAME)
//line /usr/local/go/src/net/dnsclient_unix.go:631
		// _ = "end of CoverTab[13617]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:632
		_go_fuzz_dep_.CoverTab[13618]++
//line /usr/local/go/src/net/dnsclient_unix.go:632
		// _ = "end of CoverTab[13618]"
//line /usr/local/go/src/net/dnsclient_unix.go:632
	}
//line /usr/local/go/src/net/dnsclient_unix.go:632
	// _ = "end of CoverTab[13595]"
//line /usr/local/go/src/net/dnsclient_unix.go:632
	_go_fuzz_dep_.CoverTab[13596]++
							switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/dnsclient_unix.go:634
		_go_fuzz_dep_.CoverTab[13619]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeA}
//line /usr/local/go/src/net/dnsclient_unix.go:635
		// _ = "end of CoverTab[13619]"
	case '6':
//line /usr/local/go/src/net/dnsclient_unix.go:636
		_go_fuzz_dep_.CoverTab[13620]++
								qtypes = []dnsmessage.Type{dnsmessage.TypeAAAA}
//line /usr/local/go/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[13620]"
//line /usr/local/go/src/net/dnsclient_unix.go:637
	default:
//line /usr/local/go/src/net/dnsclient_unix.go:637
		_go_fuzz_dep_.CoverTab[13621]++
//line /usr/local/go/src/net/dnsclient_unix.go:637
		// _ = "end of CoverTab[13621]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:638
	// _ = "end of CoverTab[13596]"
//line /usr/local/go/src/net/dnsclient_unix.go:638
	_go_fuzz_dep_.CoverTab[13597]++
							var queryFn func(fqdn string, qtype dnsmessage.Type)
							var responseFn func(fqdn string, qtype dnsmessage.Type) result
							if conf.singleRequest {
//line /usr/local/go/src/net/dnsclient_unix.go:641
		_go_fuzz_dep_.CoverTab[13622]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:642
			_go_fuzz_dep_.CoverTab[13624]++
//line /usr/local/go/src/net/dnsclient_unix.go:642
			// _ = "end of CoverTab[13624]"
//line /usr/local/go/src/net/dnsclient_unix.go:642
		}
//line /usr/local/go/src/net/dnsclient_unix.go:642
		// _ = "end of CoverTab[13622]"
//line /usr/local/go/src/net/dnsclient_unix.go:642
		_go_fuzz_dep_.CoverTab[13623]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /usr/local/go/src/net/dnsclient_unix.go:643
			_go_fuzz_dep_.CoverTab[13625]++
									dnsWaitGroup.Add(1)
									defer dnsWaitGroup.Done()
									p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
									return result{p, server, err}
//line /usr/local/go/src/net/dnsclient_unix.go:647
			// _ = "end of CoverTab[13625]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:648
		// _ = "end of CoverTab[13623]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:649
		_go_fuzz_dep_.CoverTab[13626]++
								queryFn = func(fqdn string, qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:650
			_go_fuzz_dep_.CoverTab[13628]++
									dnsWaitGroup.Add(1)
//line /usr/local/go/src/net/dnsclient_unix.go:651
			_curRoutineNum8_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/dnsclient_unix.go:651
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum8_)
									go func(qtype dnsmessage.Type) {
//line /usr/local/go/src/net/dnsclient_unix.go:652
				_go_fuzz_dep_.CoverTab[13629]++
//line /usr/local/go/src/net/dnsclient_unix.go:652
				defer func() {
//line /usr/local/go/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.CoverTab[13630]++
//line /usr/local/go/src/net/dnsclient_unix.go:652
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum8_)
//line /usr/local/go/src/net/dnsclient_unix.go:652
					// _ = "end of CoverTab[13630]"
//line /usr/local/go/src/net/dnsclient_unix.go:652
				}()
										p, server, err := r.tryOneName(ctx, conf, fqdn, qtype)
										lane <- result{p, server, err}
										dnsWaitGroup.Done()
//line /usr/local/go/src/net/dnsclient_unix.go:655
				// _ = "end of CoverTab[13629]"
			}(qtype)
//line /usr/local/go/src/net/dnsclient_unix.go:656
			// _ = "end of CoverTab[13628]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:657
		// _ = "end of CoverTab[13626]"
//line /usr/local/go/src/net/dnsclient_unix.go:657
		_go_fuzz_dep_.CoverTab[13627]++
								responseFn = func(fqdn string, qtype dnsmessage.Type) result {
//line /usr/local/go/src/net/dnsclient_unix.go:658
			_go_fuzz_dep_.CoverTab[13631]++
									return <-lane
//line /usr/local/go/src/net/dnsclient_unix.go:659
			// _ = "end of CoverTab[13631]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:660
		// _ = "end of CoverTab[13627]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:661
	// _ = "end of CoverTab[13597]"
//line /usr/local/go/src/net/dnsclient_unix.go:661
	_go_fuzz_dep_.CoverTab[13598]++
							var lastErr error
							for _, fqdn := range conf.nameList(name) {
//line /usr/local/go/src/net/dnsclient_unix.go:663
		_go_fuzz_dep_.CoverTab[13632]++
								for _, qtype := range qtypes {
//line /usr/local/go/src/net/dnsclient_unix.go:664
			_go_fuzz_dep_.CoverTab[13636]++
									queryFn(fqdn, qtype)
//line /usr/local/go/src/net/dnsclient_unix.go:665
			// _ = "end of CoverTab[13636]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:666
		// _ = "end of CoverTab[13632]"
//line /usr/local/go/src/net/dnsclient_unix.go:666
		_go_fuzz_dep_.CoverTab[13633]++
								hitStrictError := false
								for _, qtype := range qtypes {
//line /usr/local/go/src/net/dnsclient_unix.go:668
			_go_fuzz_dep_.CoverTab[13637]++
									result := responseFn(fqdn, qtype)
									if result.error != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:670
				_go_fuzz_dep_.CoverTab[13639]++
										if nerr, ok := result.error.(Error); ok && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[13641]++
//line /usr/local/go/src/net/dnsclient_unix.go:671
					return nerr.Temporary()
//line /usr/local/go/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[13641]"
//line /usr/local/go/src/net/dnsclient_unix.go:671
				}() && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[13642]++
//line /usr/local/go/src/net/dnsclient_unix.go:671
					return r.strictErrors()
//line /usr/local/go/src/net/dnsclient_unix.go:671
					// _ = "end of CoverTab[13642]"
//line /usr/local/go/src/net/dnsclient_unix.go:671
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:671
					_go_fuzz_dep_.CoverTab[13643]++

											hitStrictError = true
											lastErr = result.error
//line /usr/local/go/src/net/dnsclient_unix.go:674
					// _ = "end of CoverTab[13643]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:675
					_go_fuzz_dep_.CoverTab[13644]++
//line /usr/local/go/src/net/dnsclient_unix.go:675
					if lastErr == nil || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[13645]++
//line /usr/local/go/src/net/dnsclient_unix.go:675
						return fqdn == name+"."
//line /usr/local/go/src/net/dnsclient_unix.go:675
						// _ = "end of CoverTab[13645]"
//line /usr/local/go/src/net/dnsclient_unix.go:675
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:675
						_go_fuzz_dep_.CoverTab[13646]++

												lastErr = result.error
//line /usr/local/go/src/net/dnsclient_unix.go:677
						// _ = "end of CoverTab[13646]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:678
						_go_fuzz_dep_.CoverTab[13647]++
//line /usr/local/go/src/net/dnsclient_unix.go:678
						// _ = "end of CoverTab[13647]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
					}
//line /usr/local/go/src/net/dnsclient_unix.go:678
					// _ = "end of CoverTab[13644]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
				}
//line /usr/local/go/src/net/dnsclient_unix.go:678
				// _ = "end of CoverTab[13639]"
//line /usr/local/go/src/net/dnsclient_unix.go:678
				_go_fuzz_dep_.CoverTab[13640]++
										continue
//line /usr/local/go/src/net/dnsclient_unix.go:679
				// _ = "end of CoverTab[13640]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:680
				_go_fuzz_dep_.CoverTab[13648]++
//line /usr/local/go/src/net/dnsclient_unix.go:680
				// _ = "end of CoverTab[13648]"
//line /usr/local/go/src/net/dnsclient_unix.go:680
			}
//line /usr/local/go/src/net/dnsclient_unix.go:680
			// _ = "end of CoverTab[13637]"
//line /usr/local/go/src/net/dnsclient_unix.go:680
			_go_fuzz_dep_.CoverTab[13638]++

//line /usr/local/go/src/net/dnsclient_unix.go:697
		loop:
			for {
//line /usr/local/go/src/net/dnsclient_unix.go:698
				_go_fuzz_dep_.CoverTab[13649]++
										h, err := result.p.AnswerHeader()
										if err != nil && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[13652]++
//line /usr/local/go/src/net/dnsclient_unix.go:700
					return err != dnsmessage.ErrSectionDone
//line /usr/local/go/src/net/dnsclient_unix.go:700
					// _ = "end of CoverTab[13652]"
//line /usr/local/go/src/net/dnsclient_unix.go:700
				}() {
//line /usr/local/go/src/net/dnsclient_unix.go:700
					_go_fuzz_dep_.CoverTab[13653]++
											lastErr = &DNSError{
						Err:	"cannot marshal DNS message",
						Name:	name,
						Server:	result.server,
					}
//line /usr/local/go/src/net/dnsclient_unix.go:705
					// _ = "end of CoverTab[13653]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:706
					_go_fuzz_dep_.CoverTab[13654]++
//line /usr/local/go/src/net/dnsclient_unix.go:706
					// _ = "end of CoverTab[13654]"
//line /usr/local/go/src/net/dnsclient_unix.go:706
				}
//line /usr/local/go/src/net/dnsclient_unix.go:706
				// _ = "end of CoverTab[13649]"
//line /usr/local/go/src/net/dnsclient_unix.go:706
				_go_fuzz_dep_.CoverTab[13650]++
										if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:707
					_go_fuzz_dep_.CoverTab[13655]++
											break
//line /usr/local/go/src/net/dnsclient_unix.go:708
					// _ = "end of CoverTab[13655]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:709
					_go_fuzz_dep_.CoverTab[13656]++
//line /usr/local/go/src/net/dnsclient_unix.go:709
					// _ = "end of CoverTab[13656]"
//line /usr/local/go/src/net/dnsclient_unix.go:709
				}
//line /usr/local/go/src/net/dnsclient_unix.go:709
				// _ = "end of CoverTab[13650]"
//line /usr/local/go/src/net/dnsclient_unix.go:709
				_go_fuzz_dep_.CoverTab[13651]++
										switch h.Type {
				case dnsmessage.TypeA:
//line /usr/local/go/src/net/dnsclient_unix.go:711
					_go_fuzz_dep_.CoverTab[13657]++
											a, err := result.p.AResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:713
						_go_fuzz_dep_.CoverTab[13665]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:719
						// _ = "end of CoverTab[13665]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:720
						_go_fuzz_dep_.CoverTab[13666]++
//line /usr/local/go/src/net/dnsclient_unix.go:720
						// _ = "end of CoverTab[13666]"
//line /usr/local/go/src/net/dnsclient_unix.go:720
					}
//line /usr/local/go/src/net/dnsclient_unix.go:720
					// _ = "end of CoverTab[13657]"
//line /usr/local/go/src/net/dnsclient_unix.go:720
					_go_fuzz_dep_.CoverTab[13658]++
											addrs = append(addrs, IPAddr{IP: IP(a.A[:])})
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[13667]++
//line /usr/local/go/src/net/dnsclient_unix.go:722
						return h.Name.Length != 0
//line /usr/local/go/src/net/dnsclient_unix.go:722
						// _ = "end of CoverTab[13667]"
//line /usr/local/go/src/net/dnsclient_unix.go:722
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:722
						_go_fuzz_dep_.CoverTab[13668]++
												cname = h.Name
//line /usr/local/go/src/net/dnsclient_unix.go:723
						// _ = "end of CoverTab[13668]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:724
						_go_fuzz_dep_.CoverTab[13669]++
//line /usr/local/go/src/net/dnsclient_unix.go:724
						// _ = "end of CoverTab[13669]"
//line /usr/local/go/src/net/dnsclient_unix.go:724
					}
//line /usr/local/go/src/net/dnsclient_unix.go:724
					// _ = "end of CoverTab[13658]"

				case dnsmessage.TypeAAAA:
//line /usr/local/go/src/net/dnsclient_unix.go:726
					_go_fuzz_dep_.CoverTab[13659]++
											aaaa, err := result.p.AAAAResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:728
						_go_fuzz_dep_.CoverTab[13670]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:734
						// _ = "end of CoverTab[13670]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:735
						_go_fuzz_dep_.CoverTab[13671]++
//line /usr/local/go/src/net/dnsclient_unix.go:735
						// _ = "end of CoverTab[13671]"
//line /usr/local/go/src/net/dnsclient_unix.go:735
					}
//line /usr/local/go/src/net/dnsclient_unix.go:735
					// _ = "end of CoverTab[13659]"
//line /usr/local/go/src/net/dnsclient_unix.go:735
					_go_fuzz_dep_.CoverTab[13660]++
											addrs = append(addrs, IPAddr{IP: IP(aaaa.AAAA[:])})
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[13672]++
//line /usr/local/go/src/net/dnsclient_unix.go:737
						return h.Name.Length != 0
//line /usr/local/go/src/net/dnsclient_unix.go:737
						// _ = "end of CoverTab[13672]"
//line /usr/local/go/src/net/dnsclient_unix.go:737
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:737
						_go_fuzz_dep_.CoverTab[13673]++
												cname = h.Name
//line /usr/local/go/src/net/dnsclient_unix.go:738
						// _ = "end of CoverTab[13673]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:739
						_go_fuzz_dep_.CoverTab[13674]++
//line /usr/local/go/src/net/dnsclient_unix.go:739
						// _ = "end of CoverTab[13674]"
//line /usr/local/go/src/net/dnsclient_unix.go:739
					}
//line /usr/local/go/src/net/dnsclient_unix.go:739
					// _ = "end of CoverTab[13660]"

				case dnsmessage.TypeCNAME:
//line /usr/local/go/src/net/dnsclient_unix.go:741
					_go_fuzz_dep_.CoverTab[13661]++
											c, err := result.p.CNAMEResource()
											if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:743
						_go_fuzz_dep_.CoverTab[13675]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:749
						// _ = "end of CoverTab[13675]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:750
						_go_fuzz_dep_.CoverTab[13676]++
//line /usr/local/go/src/net/dnsclient_unix.go:750
						// _ = "end of CoverTab[13676]"
//line /usr/local/go/src/net/dnsclient_unix.go:750
					}
//line /usr/local/go/src/net/dnsclient_unix.go:750
					// _ = "end of CoverTab[13661]"
//line /usr/local/go/src/net/dnsclient_unix.go:750
					_go_fuzz_dep_.CoverTab[13662]++
											if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[13677]++
//line /usr/local/go/src/net/dnsclient_unix.go:751
						return c.CNAME.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:751
						// _ = "end of CoverTab[13677]"
//line /usr/local/go/src/net/dnsclient_unix.go:751
					}() {
//line /usr/local/go/src/net/dnsclient_unix.go:751
						_go_fuzz_dep_.CoverTab[13678]++
												cname = c.CNAME
//line /usr/local/go/src/net/dnsclient_unix.go:752
						// _ = "end of CoverTab[13678]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:753
						_go_fuzz_dep_.CoverTab[13679]++
//line /usr/local/go/src/net/dnsclient_unix.go:753
						// _ = "end of CoverTab[13679]"
//line /usr/local/go/src/net/dnsclient_unix.go:753
					}
//line /usr/local/go/src/net/dnsclient_unix.go:753
					// _ = "end of CoverTab[13662]"

				default:
//line /usr/local/go/src/net/dnsclient_unix.go:755
					_go_fuzz_dep_.CoverTab[13663]++
											if err := result.p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:756
						_go_fuzz_dep_.CoverTab[13680]++
												lastErr = &DNSError{
							Err:	"cannot marshal DNS message",
							Name:	name,
							Server:	result.server,
						}
												break loop
//line /usr/local/go/src/net/dnsclient_unix.go:762
						// _ = "end of CoverTab[13680]"
					} else {
//line /usr/local/go/src/net/dnsclient_unix.go:763
						_go_fuzz_dep_.CoverTab[13681]++
//line /usr/local/go/src/net/dnsclient_unix.go:763
						// _ = "end of CoverTab[13681]"
//line /usr/local/go/src/net/dnsclient_unix.go:763
					}
//line /usr/local/go/src/net/dnsclient_unix.go:763
					// _ = "end of CoverTab[13663]"
//line /usr/local/go/src/net/dnsclient_unix.go:763
					_go_fuzz_dep_.CoverTab[13664]++
											continue
//line /usr/local/go/src/net/dnsclient_unix.go:764
					// _ = "end of CoverTab[13664]"
				}
//line /usr/local/go/src/net/dnsclient_unix.go:765
				// _ = "end of CoverTab[13651]"
			}
//line /usr/local/go/src/net/dnsclient_unix.go:766
			// _ = "end of CoverTab[13638]"
		}
//line /usr/local/go/src/net/dnsclient_unix.go:767
		// _ = "end of CoverTab[13633]"
//line /usr/local/go/src/net/dnsclient_unix.go:767
		_go_fuzz_dep_.CoverTab[13634]++
								if hitStrictError {
//line /usr/local/go/src/net/dnsclient_unix.go:768
			_go_fuzz_dep_.CoverTab[13682]++

//line /usr/local/go/src/net/dnsclient_unix.go:772
			addrs = nil
									break
//line /usr/local/go/src/net/dnsclient_unix.go:773
			// _ = "end of CoverTab[13682]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:774
			_go_fuzz_dep_.CoverTab[13683]++
//line /usr/local/go/src/net/dnsclient_unix.go:774
			// _ = "end of CoverTab[13683]"
//line /usr/local/go/src/net/dnsclient_unix.go:774
		}
//line /usr/local/go/src/net/dnsclient_unix.go:774
		// _ = "end of CoverTab[13634]"
//line /usr/local/go/src/net/dnsclient_unix.go:774
		_go_fuzz_dep_.CoverTab[13635]++
								if len(addrs) > 0 || func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[13684]++
//line /usr/local/go/src/net/dnsclient_unix.go:775
			return network == "CNAME" && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:775
				_go_fuzz_dep_.CoverTab[13685]++
//line /usr/local/go/src/net/dnsclient_unix.go:775
				return cname.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:775
				// _ = "end of CoverTab[13685]"
//line /usr/local/go/src/net/dnsclient_unix.go:775
			}()
//line /usr/local/go/src/net/dnsclient_unix.go:775
			// _ = "end of CoverTab[13684]"
//line /usr/local/go/src/net/dnsclient_unix.go:775
		}() {
//line /usr/local/go/src/net/dnsclient_unix.go:775
			_go_fuzz_dep_.CoverTab[13686]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:776
			// _ = "end of CoverTab[13686]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:777
			_go_fuzz_dep_.CoverTab[13687]++
//line /usr/local/go/src/net/dnsclient_unix.go:777
			// _ = "end of CoverTab[13687]"
//line /usr/local/go/src/net/dnsclient_unix.go:777
		}
//line /usr/local/go/src/net/dnsclient_unix.go:777
		// _ = "end of CoverTab[13635]"
	}
//line /usr/local/go/src/net/dnsclient_unix.go:778
	// _ = "end of CoverTab[13598]"
//line /usr/local/go/src/net/dnsclient_unix.go:778
	_go_fuzz_dep_.CoverTab[13599]++
							if lastErr, ok := lastErr.(*DNSError); ok {
//line /usr/local/go/src/net/dnsclient_unix.go:779
		_go_fuzz_dep_.CoverTab[13688]++

//line /usr/local/go/src/net/dnsclient_unix.go:783
		lastErr.Name = name
//line /usr/local/go/src/net/dnsclient_unix.go:783
		// _ = "end of CoverTab[13688]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:784
		_go_fuzz_dep_.CoverTab[13689]++
//line /usr/local/go/src/net/dnsclient_unix.go:784
		// _ = "end of CoverTab[13689]"
//line /usr/local/go/src/net/dnsclient_unix.go:784
	}
//line /usr/local/go/src/net/dnsclient_unix.go:784
	// _ = "end of CoverTab[13599]"
//line /usr/local/go/src/net/dnsclient_unix.go:784
	_go_fuzz_dep_.CoverTab[13600]++
							sortByRFC6724(addrs)
							if len(addrs) == 0 && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[13690]++
//line /usr/local/go/src/net/dnsclient_unix.go:786
		return !(network == "CNAME" && func() bool {
//line /usr/local/go/src/net/dnsclient_unix.go:786
			_go_fuzz_dep_.CoverTab[13691]++
//line /usr/local/go/src/net/dnsclient_unix.go:786
			return cname.Length > 0
//line /usr/local/go/src/net/dnsclient_unix.go:786
			// _ = "end of CoverTab[13691]"
//line /usr/local/go/src/net/dnsclient_unix.go:786
		}())
//line /usr/local/go/src/net/dnsclient_unix.go:786
		// _ = "end of CoverTab[13690]"
//line /usr/local/go/src/net/dnsclient_unix.go:786
	}() {
//line /usr/local/go/src/net/dnsclient_unix.go:786
		_go_fuzz_dep_.CoverTab[13692]++
								if order == hostLookupDNSFiles {
//line /usr/local/go/src/net/dnsclient_unix.go:787
			_go_fuzz_dep_.CoverTab[13694]++
									var canonical string
									addrs, canonical = goLookupIPFiles(name)
									if len(addrs) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:790
				_go_fuzz_dep_.CoverTab[13695]++
										var err error
										cname, err = dnsmessage.NewName(canonical)
										if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:793
					_go_fuzz_dep_.CoverTab[13697]++
											return nil, dnsmessage.Name{}, err
//line /usr/local/go/src/net/dnsclient_unix.go:794
					// _ = "end of CoverTab[13697]"
				} else {
//line /usr/local/go/src/net/dnsclient_unix.go:795
					_go_fuzz_dep_.CoverTab[13698]++
//line /usr/local/go/src/net/dnsclient_unix.go:795
					// _ = "end of CoverTab[13698]"
//line /usr/local/go/src/net/dnsclient_unix.go:795
				}
//line /usr/local/go/src/net/dnsclient_unix.go:795
				// _ = "end of CoverTab[13695]"
//line /usr/local/go/src/net/dnsclient_unix.go:795
				_go_fuzz_dep_.CoverTab[13696]++
										return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:796
				// _ = "end of CoverTab[13696]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:797
				_go_fuzz_dep_.CoverTab[13699]++
//line /usr/local/go/src/net/dnsclient_unix.go:797
				// _ = "end of CoverTab[13699]"
//line /usr/local/go/src/net/dnsclient_unix.go:797
			}
//line /usr/local/go/src/net/dnsclient_unix.go:797
			// _ = "end of CoverTab[13694]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:798
			_go_fuzz_dep_.CoverTab[13700]++
//line /usr/local/go/src/net/dnsclient_unix.go:798
			// _ = "end of CoverTab[13700]"
//line /usr/local/go/src/net/dnsclient_unix.go:798
		}
//line /usr/local/go/src/net/dnsclient_unix.go:798
		// _ = "end of CoverTab[13692]"
//line /usr/local/go/src/net/dnsclient_unix.go:798
		_go_fuzz_dep_.CoverTab[13693]++
								if lastErr != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:799
			_go_fuzz_dep_.CoverTab[13701]++
									return nil, dnsmessage.Name{}, lastErr
//line /usr/local/go/src/net/dnsclient_unix.go:800
			// _ = "end of CoverTab[13701]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:801
			_go_fuzz_dep_.CoverTab[13702]++
//line /usr/local/go/src/net/dnsclient_unix.go:801
			// _ = "end of CoverTab[13702]"
//line /usr/local/go/src/net/dnsclient_unix.go:801
		}
//line /usr/local/go/src/net/dnsclient_unix.go:801
		// _ = "end of CoverTab[13693]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:802
		_go_fuzz_dep_.CoverTab[13703]++
//line /usr/local/go/src/net/dnsclient_unix.go:802
		// _ = "end of CoverTab[13703]"
//line /usr/local/go/src/net/dnsclient_unix.go:802
	}
//line /usr/local/go/src/net/dnsclient_unix.go:802
	// _ = "end of CoverTab[13600]"
//line /usr/local/go/src/net/dnsclient_unix.go:802
	_go_fuzz_dep_.CoverTab[13601]++
							return addrs, cname, nil
//line /usr/local/go/src/net/dnsclient_unix.go:803
	// _ = "end of CoverTab[13601]"
}

// goLookupCNAME is the native Go (non-cgo) implementation of LookupCNAME.
func (r *Resolver) goLookupCNAME(ctx context.Context, host string, order hostLookupOrder, conf *dnsConfig) (string, error) {
//line /usr/local/go/src/net/dnsclient_unix.go:807
	_go_fuzz_dep_.CoverTab[13704]++
							_, cname, err := r.goLookupIPCNAMEOrder(ctx, "CNAME", host, order, conf)
							return cname.String(), err
//line /usr/local/go/src/net/dnsclient_unix.go:809
	// _ = "end of CoverTab[13704]"
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
	_go_fuzz_dep_.CoverTab[13705]++
							names := lookupStaticAddr(addr)
							if len(names) > 0 {
//line /usr/local/go/src/net/dnsclient_unix.go:819
		_go_fuzz_dep_.CoverTab[13710]++
								return names, nil
//line /usr/local/go/src/net/dnsclient_unix.go:820
		// _ = "end of CoverTab[13710]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:821
		_go_fuzz_dep_.CoverTab[13711]++
//line /usr/local/go/src/net/dnsclient_unix.go:821
		// _ = "end of CoverTab[13711]"
//line /usr/local/go/src/net/dnsclient_unix.go:821
	}
//line /usr/local/go/src/net/dnsclient_unix.go:821
	// _ = "end of CoverTab[13705]"
//line /usr/local/go/src/net/dnsclient_unix.go:821
	_go_fuzz_dep_.CoverTab[13706]++
							arpa, err := reverseaddr(addr)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:823
		_go_fuzz_dep_.CoverTab[13712]++
								return nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:824
		// _ = "end of CoverTab[13712]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:825
		_go_fuzz_dep_.CoverTab[13713]++
//line /usr/local/go/src/net/dnsclient_unix.go:825
		// _ = "end of CoverTab[13713]"
//line /usr/local/go/src/net/dnsclient_unix.go:825
	}
//line /usr/local/go/src/net/dnsclient_unix.go:825
	// _ = "end of CoverTab[13706]"
//line /usr/local/go/src/net/dnsclient_unix.go:825
	_go_fuzz_dep_.CoverTab[13707]++
							p, server, err := r.lookup(ctx, arpa, dnsmessage.TypePTR, conf)
							if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:827
		_go_fuzz_dep_.CoverTab[13714]++
								return nil, err
//line /usr/local/go/src/net/dnsclient_unix.go:828
		// _ = "end of CoverTab[13714]"
	} else {
//line /usr/local/go/src/net/dnsclient_unix.go:829
		_go_fuzz_dep_.CoverTab[13715]++
//line /usr/local/go/src/net/dnsclient_unix.go:829
		// _ = "end of CoverTab[13715]"
//line /usr/local/go/src/net/dnsclient_unix.go:829
	}
//line /usr/local/go/src/net/dnsclient_unix.go:829
	// _ = "end of CoverTab[13707]"
//line /usr/local/go/src/net/dnsclient_unix.go:829
	_go_fuzz_dep_.CoverTab[13708]++
							var ptrs []string
							for {
//line /usr/local/go/src/net/dnsclient_unix.go:831
		_go_fuzz_dep_.CoverTab[13716]++
								h, err := p.AnswerHeader()
								if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/dnsclient_unix.go:833
			_go_fuzz_dep_.CoverTab[13721]++
									break
//line /usr/local/go/src/net/dnsclient_unix.go:834
			// _ = "end of CoverTab[13721]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:835
			_go_fuzz_dep_.CoverTab[13722]++
//line /usr/local/go/src/net/dnsclient_unix.go:835
			// _ = "end of CoverTab[13722]"
//line /usr/local/go/src/net/dnsclient_unix.go:835
		}
//line /usr/local/go/src/net/dnsclient_unix.go:835
		// _ = "end of CoverTab[13716]"
//line /usr/local/go/src/net/dnsclient_unix.go:835
		_go_fuzz_dep_.CoverTab[13717]++
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:836
			_go_fuzz_dep_.CoverTab[13723]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /usr/local/go/src/net/dnsclient_unix.go:841
			// _ = "end of CoverTab[13723]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:842
			_go_fuzz_dep_.CoverTab[13724]++
//line /usr/local/go/src/net/dnsclient_unix.go:842
			// _ = "end of CoverTab[13724]"
//line /usr/local/go/src/net/dnsclient_unix.go:842
		}
//line /usr/local/go/src/net/dnsclient_unix.go:842
		// _ = "end of CoverTab[13717]"
//line /usr/local/go/src/net/dnsclient_unix.go:842
		_go_fuzz_dep_.CoverTab[13718]++
								if h.Type != dnsmessage.TypePTR {
//line /usr/local/go/src/net/dnsclient_unix.go:843
			_go_fuzz_dep_.CoverTab[13725]++
									err := p.SkipAnswer()
									if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:845
				_go_fuzz_dep_.CoverTab[13727]++
										return nil, &DNSError{
					Err:	"cannot marshal DNS message",
					Name:	addr,
					Server:	server,
				}
//line /usr/local/go/src/net/dnsclient_unix.go:850
				// _ = "end of CoverTab[13727]"
			} else {
//line /usr/local/go/src/net/dnsclient_unix.go:851
				_go_fuzz_dep_.CoverTab[13728]++
//line /usr/local/go/src/net/dnsclient_unix.go:851
				// _ = "end of CoverTab[13728]"
//line /usr/local/go/src/net/dnsclient_unix.go:851
			}
//line /usr/local/go/src/net/dnsclient_unix.go:851
			// _ = "end of CoverTab[13725]"
//line /usr/local/go/src/net/dnsclient_unix.go:851
			_go_fuzz_dep_.CoverTab[13726]++
									continue
//line /usr/local/go/src/net/dnsclient_unix.go:852
			// _ = "end of CoverTab[13726]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:853
			_go_fuzz_dep_.CoverTab[13729]++
//line /usr/local/go/src/net/dnsclient_unix.go:853
			// _ = "end of CoverTab[13729]"
//line /usr/local/go/src/net/dnsclient_unix.go:853
		}
//line /usr/local/go/src/net/dnsclient_unix.go:853
		// _ = "end of CoverTab[13718]"
//line /usr/local/go/src/net/dnsclient_unix.go:853
		_go_fuzz_dep_.CoverTab[13719]++
								ptr, err := p.PTRResource()
								if err != nil {
//line /usr/local/go/src/net/dnsclient_unix.go:855
			_go_fuzz_dep_.CoverTab[13730]++
									return nil, &DNSError{
				Err:	"cannot marshal DNS message",
				Name:	addr,
				Server:	server,
			}
//line /usr/local/go/src/net/dnsclient_unix.go:860
			// _ = "end of CoverTab[13730]"
		} else {
//line /usr/local/go/src/net/dnsclient_unix.go:861
			_go_fuzz_dep_.CoverTab[13731]++
//line /usr/local/go/src/net/dnsclient_unix.go:861
			// _ = "end of CoverTab[13731]"
//line /usr/local/go/src/net/dnsclient_unix.go:861
		}
//line /usr/local/go/src/net/dnsclient_unix.go:861
		// _ = "end of CoverTab[13719]"
//line /usr/local/go/src/net/dnsclient_unix.go:861
		_go_fuzz_dep_.CoverTab[13720]++
								ptrs = append(ptrs, ptr.PTR.String())
//line /usr/local/go/src/net/dnsclient_unix.go:862
		// _ = "end of CoverTab[13720]"

	}
//line /usr/local/go/src/net/dnsclient_unix.go:864
	// _ = "end of CoverTab[13708]"
//line /usr/local/go/src/net/dnsclient_unix.go:864
	_go_fuzz_dep_.CoverTab[13709]++
							return ptrs, nil
//line /usr/local/go/src/net/dnsclient_unix.go:865
	// _ = "end of CoverTab[13709]"
}

//line /usr/local/go/src/net/dnsclient_unix.go:866
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dnsclient_unix.go:866
var _ = _go_fuzz_dep_.CoverTab
