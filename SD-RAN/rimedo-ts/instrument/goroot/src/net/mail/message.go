// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/mail/message.go:5
/*
Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by RFC 5322 and
extended by RFC 6532.
Notable divergences:
  - Obsolete address formats are not parsed, including addresses with
    embedded route information.
  - The full range of spacing (the CFWS syntax element) is not supported,
    such as breaking addresses across lines.
  - No unicode normalization is performed.
  - The special characters ()[]:;@\, are allowed to appear unquoted in names.
  - A leading From line is permitted, as in mbox format (RFC 4155).
*/
package mail

//line /usr/local/go/src/net/mail/message.go:19
import (
//line /usr/local/go/src/net/mail/message.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/mail/message.go:19
)
//line /usr/local/go/src/net/mail/message.go:19
import (
//line /usr/local/go/src/net/mail/message.go:19
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/mail/message.go:19
)

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/textproto"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var debug = debugT(false)

type debugT bool

func (d debugT) Printf(format string, args ...any) {
//line /usr/local/go/src/net/mail/message.go:39
	_go_fuzz_dep_.CoverTab[171054]++
							if d {
//line /usr/local/go/src/net/mail/message.go:40
		_go_fuzz_dep_.CoverTab[171055]++
								log.Printf(format, args...)
//line /usr/local/go/src/net/mail/message.go:41
		// _ = "end of CoverTab[171055]"
	} else {
//line /usr/local/go/src/net/mail/message.go:42
		_go_fuzz_dep_.CoverTab[171056]++
//line /usr/local/go/src/net/mail/message.go:42
		// _ = "end of CoverTab[171056]"
//line /usr/local/go/src/net/mail/message.go:42
	}
//line /usr/local/go/src/net/mail/message.go:42
	// _ = "end of CoverTab[171054]"
}

// A Message represents a parsed mail message.
type Message struct {
	Header	Header
	Body	io.Reader
}

// ReadMessage reads a message from r.
//line /usr/local/go/src/net/mail/message.go:51
// The headers are parsed, and the body of the message will be available
//line /usr/local/go/src/net/mail/message.go:51
// for reading from msg.Body.
//line /usr/local/go/src/net/mail/message.go:54
func ReadMessage(r io.Reader) (msg *Message, err error) {
//line /usr/local/go/src/net/mail/message.go:54
	_go_fuzz_dep_.CoverTab[171057]++
							tp := textproto.NewReader(bufio.NewReader(r))

							hdr, err := readHeader(tp)
							if err != nil {
//line /usr/local/go/src/net/mail/message.go:58
		_go_fuzz_dep_.CoverTab[171059]++
								return nil, err
//line /usr/local/go/src/net/mail/message.go:59
		// _ = "end of CoverTab[171059]"
	} else {
//line /usr/local/go/src/net/mail/message.go:60
		_go_fuzz_dep_.CoverTab[171060]++
//line /usr/local/go/src/net/mail/message.go:60
		// _ = "end of CoverTab[171060]"
//line /usr/local/go/src/net/mail/message.go:60
	}
//line /usr/local/go/src/net/mail/message.go:60
	// _ = "end of CoverTab[171057]"
//line /usr/local/go/src/net/mail/message.go:60
	_go_fuzz_dep_.CoverTab[171058]++

							return &Message{
		Header:	Header(hdr),
		Body:	tp.R,
	}, nil
//line /usr/local/go/src/net/mail/message.go:65
	// _ = "end of CoverTab[171058]"
}

// readHeader reads the message headers from r.
//line /usr/local/go/src/net/mail/message.go:68
// This is like textproto.ReadMIMEHeader, but doesn't validate.
//line /usr/local/go/src/net/mail/message.go:68
// The fix for issue #53188 tightened up net/textproto to enforce
//line /usr/local/go/src/net/mail/message.go:68
// restrictions of RFC 7230.
//line /usr/local/go/src/net/mail/message.go:68
// This package implements RFC 5322, which does not have those restrictions.
//line /usr/local/go/src/net/mail/message.go:68
// This function copies the relevant code from net/textproto,
//line /usr/local/go/src/net/mail/message.go:68
// simplified for RFC 5322.
//line /usr/local/go/src/net/mail/message.go:75
func readHeader(r *textproto.Reader) (map[string][]string, error) {
//line /usr/local/go/src/net/mail/message.go:75
	_go_fuzz_dep_.CoverTab[171061]++
							m := make(map[string][]string)

//line /usr/local/go/src/net/mail/message.go:79
	if buf, err := r.R.Peek(1); err == nil && func() bool {
//line /usr/local/go/src/net/mail/message.go:79
		_go_fuzz_dep_.CoverTab[171063]++
//line /usr/local/go/src/net/mail/message.go:79
		return (buf[0] == ' ' || func() bool {
//line /usr/local/go/src/net/mail/message.go:79
			_go_fuzz_dep_.CoverTab[171064]++
//line /usr/local/go/src/net/mail/message.go:79
			return buf[0] == '\t'
//line /usr/local/go/src/net/mail/message.go:79
			// _ = "end of CoverTab[171064]"
//line /usr/local/go/src/net/mail/message.go:79
		}())
//line /usr/local/go/src/net/mail/message.go:79
		// _ = "end of CoverTab[171063]"
//line /usr/local/go/src/net/mail/message.go:79
	}() {
//line /usr/local/go/src/net/mail/message.go:79
		_go_fuzz_dep_.CoverTab[171065]++
								line, err := r.ReadLine()
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:81
			_go_fuzz_dep_.CoverTab[171067]++
									return m, err
//line /usr/local/go/src/net/mail/message.go:82
			// _ = "end of CoverTab[171067]"
		} else {
//line /usr/local/go/src/net/mail/message.go:83
			_go_fuzz_dep_.CoverTab[171068]++
//line /usr/local/go/src/net/mail/message.go:83
			// _ = "end of CoverTab[171068]"
//line /usr/local/go/src/net/mail/message.go:83
		}
//line /usr/local/go/src/net/mail/message.go:83
		// _ = "end of CoverTab[171065]"
//line /usr/local/go/src/net/mail/message.go:83
		_go_fuzz_dep_.CoverTab[171066]++
								return m, errors.New("malformed initial line: " + line)
//line /usr/local/go/src/net/mail/message.go:84
		// _ = "end of CoverTab[171066]"
	} else {
//line /usr/local/go/src/net/mail/message.go:85
		_go_fuzz_dep_.CoverTab[171069]++
//line /usr/local/go/src/net/mail/message.go:85
		// _ = "end of CoverTab[171069]"
//line /usr/local/go/src/net/mail/message.go:85
	}
//line /usr/local/go/src/net/mail/message.go:85
	// _ = "end of CoverTab[171061]"
//line /usr/local/go/src/net/mail/message.go:85
	_go_fuzz_dep_.CoverTab[171062]++

							for {
//line /usr/local/go/src/net/mail/message.go:87
		_go_fuzz_dep_.CoverTab[171070]++
								kv, err := r.ReadContinuedLine()
								if kv == "" {
//line /usr/local/go/src/net/mail/message.go:89
			_go_fuzz_dep_.CoverTab[171074]++
									return m, err
//line /usr/local/go/src/net/mail/message.go:90
			// _ = "end of CoverTab[171074]"
		} else {
//line /usr/local/go/src/net/mail/message.go:91
			_go_fuzz_dep_.CoverTab[171075]++
//line /usr/local/go/src/net/mail/message.go:91
			// _ = "end of CoverTab[171075]"
//line /usr/local/go/src/net/mail/message.go:91
		}
//line /usr/local/go/src/net/mail/message.go:91
		// _ = "end of CoverTab[171070]"
//line /usr/local/go/src/net/mail/message.go:91
		_go_fuzz_dep_.CoverTab[171071]++

//line /usr/local/go/src/net/mail/message.go:94
		k, v, ok := strings.Cut(kv, ":")
		if !ok {
//line /usr/local/go/src/net/mail/message.go:95
			_go_fuzz_dep_.CoverTab[171076]++
									return m, errors.New("malformed header line: " + kv)
//line /usr/local/go/src/net/mail/message.go:96
			// _ = "end of CoverTab[171076]"
		} else {
//line /usr/local/go/src/net/mail/message.go:97
			_go_fuzz_dep_.CoverTab[171077]++
//line /usr/local/go/src/net/mail/message.go:97
			// _ = "end of CoverTab[171077]"
//line /usr/local/go/src/net/mail/message.go:97
		}
//line /usr/local/go/src/net/mail/message.go:97
		// _ = "end of CoverTab[171071]"
//line /usr/local/go/src/net/mail/message.go:97
		_go_fuzz_dep_.CoverTab[171072]++
								key := textproto.CanonicalMIMEHeaderKey(k)

//line /usr/local/go/src/net/mail/message.go:101
		if key == "" {
//line /usr/local/go/src/net/mail/message.go:101
			_go_fuzz_dep_.CoverTab[171078]++
									continue
//line /usr/local/go/src/net/mail/message.go:102
			// _ = "end of CoverTab[171078]"
		} else {
//line /usr/local/go/src/net/mail/message.go:103
			_go_fuzz_dep_.CoverTab[171079]++
//line /usr/local/go/src/net/mail/message.go:103
			// _ = "end of CoverTab[171079]"
//line /usr/local/go/src/net/mail/message.go:103
		}
//line /usr/local/go/src/net/mail/message.go:103
		// _ = "end of CoverTab[171072]"
//line /usr/local/go/src/net/mail/message.go:103
		_go_fuzz_dep_.CoverTab[171073]++

//line /usr/local/go/src/net/mail/message.go:106
		value := strings.TrimLeft(v, " \t")

		m[key] = append(m[key], value)

		if err != nil {
//line /usr/local/go/src/net/mail/message.go:110
			_go_fuzz_dep_.CoverTab[171080]++
									return m, err
//line /usr/local/go/src/net/mail/message.go:111
			// _ = "end of CoverTab[171080]"
		} else {
//line /usr/local/go/src/net/mail/message.go:112
			_go_fuzz_dep_.CoverTab[171081]++
//line /usr/local/go/src/net/mail/message.go:112
			// _ = "end of CoverTab[171081]"
//line /usr/local/go/src/net/mail/message.go:112
		}
//line /usr/local/go/src/net/mail/message.go:112
		// _ = "end of CoverTab[171073]"
	}
//line /usr/local/go/src/net/mail/message.go:113
	// _ = "end of CoverTab[171062]"
}

// Layouts suitable for passing to time.Parse.
//line /usr/local/go/src/net/mail/message.go:116
// These are tried in order.
//line /usr/local/go/src/net/mail/message.go:118
var (
	dateLayoutsBuildOnce	sync.Once
	dateLayouts		[]string
)

func buildDateLayouts() {
//line /usr/local/go/src/net/mail/message.go:123
	_go_fuzz_dep_.CoverTab[171082]++

//line /usr/local/go/src/net/mail/message.go:126
	dows := [...]string{"", "Mon, "}
	days := [...]string{"2", "02"}
	years := [...]string{"2006", "06"}
	seconds := [...]string{":05", ""}

	zones := [...]string{"-0700", "MST", "UT"}

	for _, dow := range dows {
//line /usr/local/go/src/net/mail/message.go:133
		_go_fuzz_dep_.CoverTab[171083]++
								for _, day := range days {
//line /usr/local/go/src/net/mail/message.go:134
			_go_fuzz_dep_.CoverTab[171084]++
									for _, year := range years {
//line /usr/local/go/src/net/mail/message.go:135
				_go_fuzz_dep_.CoverTab[171085]++
										for _, second := range seconds {
//line /usr/local/go/src/net/mail/message.go:136
					_go_fuzz_dep_.CoverTab[171086]++
											for _, zone := range zones {
//line /usr/local/go/src/net/mail/message.go:137
						_go_fuzz_dep_.CoverTab[171087]++
												s := dow + day + " Jan " + year + " 15:04" + second + " " + zone
												dateLayouts = append(dateLayouts, s)
//line /usr/local/go/src/net/mail/message.go:139
						// _ = "end of CoverTab[171087]"
					}
//line /usr/local/go/src/net/mail/message.go:140
					// _ = "end of CoverTab[171086]"
				}
//line /usr/local/go/src/net/mail/message.go:141
				// _ = "end of CoverTab[171085]"
			}
//line /usr/local/go/src/net/mail/message.go:142
			// _ = "end of CoverTab[171084]"
		}
//line /usr/local/go/src/net/mail/message.go:143
		// _ = "end of CoverTab[171083]"
	}
//line /usr/local/go/src/net/mail/message.go:144
	// _ = "end of CoverTab[171082]"
}

// ParseDate parses an RFC 5322 date string.
func ParseDate(date string) (time.Time, error) {
//line /usr/local/go/src/net/mail/message.go:148
	_go_fuzz_dep_.CoverTab[171088]++
							dateLayoutsBuildOnce.Do(buildDateLayouts)

							date = strings.ReplaceAll(date, "\r\n", "")
							if strings.Contains(date, "\r") {
//line /usr/local/go/src/net/mail/message.go:152
		_go_fuzz_dep_.CoverTab[171093]++
								return time.Time{}, errors.New("mail: header has a CR without LF")
//line /usr/local/go/src/net/mail/message.go:153
		// _ = "end of CoverTab[171093]"
	} else {
//line /usr/local/go/src/net/mail/message.go:154
		_go_fuzz_dep_.CoverTab[171094]++
//line /usr/local/go/src/net/mail/message.go:154
		// _ = "end of CoverTab[171094]"
//line /usr/local/go/src/net/mail/message.go:154
	}
//line /usr/local/go/src/net/mail/message.go:154
	// _ = "end of CoverTab[171088]"
//line /usr/local/go/src/net/mail/message.go:154
	_go_fuzz_dep_.CoverTab[171089]++

							p := addrParser{date, nil}
							p.skipSpace()

//line /usr/local/go/src/net/mail/message.go:161
	if ind := strings.IndexAny(p.s, "+-"); ind != -1 && func() bool {
//line /usr/local/go/src/net/mail/message.go:161
		_go_fuzz_dep_.CoverTab[171095]++
//line /usr/local/go/src/net/mail/message.go:161
		return len(p.s) >= ind+5
//line /usr/local/go/src/net/mail/message.go:161
		// _ = "end of CoverTab[171095]"
//line /usr/local/go/src/net/mail/message.go:161
	}() {
//line /usr/local/go/src/net/mail/message.go:161
		_go_fuzz_dep_.CoverTab[171096]++
								date = p.s[:ind+5]
								p.s = p.s[ind+5:]
//line /usr/local/go/src/net/mail/message.go:163
		// _ = "end of CoverTab[171096]"
	} else {
//line /usr/local/go/src/net/mail/message.go:164
		_go_fuzz_dep_.CoverTab[171097]++
								ind := strings.Index(p.s, "T")
								if ind == 0 {
//line /usr/local/go/src/net/mail/message.go:166
			_go_fuzz_dep_.CoverTab[171099]++

//line /usr/local/go/src/net/mail/message.go:171
			ind = strings.Index(p.s[1:], "T")
			if ind != -1 {
//line /usr/local/go/src/net/mail/message.go:172
				_go_fuzz_dep_.CoverTab[171100]++
										ind++
//line /usr/local/go/src/net/mail/message.go:173
				// _ = "end of CoverTab[171100]"
			} else {
//line /usr/local/go/src/net/mail/message.go:174
				_go_fuzz_dep_.CoverTab[171101]++
//line /usr/local/go/src/net/mail/message.go:174
				// _ = "end of CoverTab[171101]"
//line /usr/local/go/src/net/mail/message.go:174
			}
//line /usr/local/go/src/net/mail/message.go:174
			// _ = "end of CoverTab[171099]"
		} else {
//line /usr/local/go/src/net/mail/message.go:175
			_go_fuzz_dep_.CoverTab[171102]++
//line /usr/local/go/src/net/mail/message.go:175
			// _ = "end of CoverTab[171102]"
//line /usr/local/go/src/net/mail/message.go:175
		}
//line /usr/local/go/src/net/mail/message.go:175
		// _ = "end of CoverTab[171097]"
//line /usr/local/go/src/net/mail/message.go:175
		_go_fuzz_dep_.CoverTab[171098]++

								if ind != -1 && func() bool {
//line /usr/local/go/src/net/mail/message.go:177
			_go_fuzz_dep_.CoverTab[171103]++
//line /usr/local/go/src/net/mail/message.go:177
			return len(p.s) >= ind+5
//line /usr/local/go/src/net/mail/message.go:177
			// _ = "end of CoverTab[171103]"
//line /usr/local/go/src/net/mail/message.go:177
		}() {
//line /usr/local/go/src/net/mail/message.go:177
			_go_fuzz_dep_.CoverTab[171104]++

//line /usr/local/go/src/net/mail/message.go:180
			date = p.s[:ind+1]
									p.s = p.s[ind+1:]
//line /usr/local/go/src/net/mail/message.go:181
			// _ = "end of CoverTab[171104]"
		} else {
//line /usr/local/go/src/net/mail/message.go:182
			_go_fuzz_dep_.CoverTab[171105]++
//line /usr/local/go/src/net/mail/message.go:182
			// _ = "end of CoverTab[171105]"
//line /usr/local/go/src/net/mail/message.go:182
		}
//line /usr/local/go/src/net/mail/message.go:182
		// _ = "end of CoverTab[171098]"
	}
//line /usr/local/go/src/net/mail/message.go:183
	// _ = "end of CoverTab[171089]"
//line /usr/local/go/src/net/mail/message.go:183
	_go_fuzz_dep_.CoverTab[171090]++
							if !p.skipCFWS() {
//line /usr/local/go/src/net/mail/message.go:184
		_go_fuzz_dep_.CoverTab[171106]++
								return time.Time{}, errors.New("mail: misformatted parenthetical comment")
//line /usr/local/go/src/net/mail/message.go:185
		// _ = "end of CoverTab[171106]"
	} else {
//line /usr/local/go/src/net/mail/message.go:186
		_go_fuzz_dep_.CoverTab[171107]++
//line /usr/local/go/src/net/mail/message.go:186
		// _ = "end of CoverTab[171107]"
//line /usr/local/go/src/net/mail/message.go:186
	}
//line /usr/local/go/src/net/mail/message.go:186
	// _ = "end of CoverTab[171090]"
//line /usr/local/go/src/net/mail/message.go:186
	_go_fuzz_dep_.CoverTab[171091]++
							for _, layout := range dateLayouts {
//line /usr/local/go/src/net/mail/message.go:187
		_go_fuzz_dep_.CoverTab[171108]++
								t, err := time.Parse(layout, date)
								if err == nil {
//line /usr/local/go/src/net/mail/message.go:189
			_go_fuzz_dep_.CoverTab[171109]++
									return t, nil
//line /usr/local/go/src/net/mail/message.go:190
			// _ = "end of CoverTab[171109]"
		} else {
//line /usr/local/go/src/net/mail/message.go:191
			_go_fuzz_dep_.CoverTab[171110]++
//line /usr/local/go/src/net/mail/message.go:191
			// _ = "end of CoverTab[171110]"
//line /usr/local/go/src/net/mail/message.go:191
		}
//line /usr/local/go/src/net/mail/message.go:191
		// _ = "end of CoverTab[171108]"
	}
//line /usr/local/go/src/net/mail/message.go:192
	// _ = "end of CoverTab[171091]"
//line /usr/local/go/src/net/mail/message.go:192
	_go_fuzz_dep_.CoverTab[171092]++
							return time.Time{}, errors.New("mail: header could not be parsed")
//line /usr/local/go/src/net/mail/message.go:193
	// _ = "end of CoverTab[171092]"
}

// A Header represents the key-value pairs in a mail message header.
type Header map[string][]string

// Get gets the first value associated with the given key.
//line /usr/local/go/src/net/mail/message.go:199
// It is case insensitive; CanonicalMIMEHeaderKey is used
//line /usr/local/go/src/net/mail/message.go:199
// to canonicalize the provided key.
//line /usr/local/go/src/net/mail/message.go:199
// If there are no values associated with the key, Get returns "".
//line /usr/local/go/src/net/mail/message.go:199
// To access multiple values of a key, or to use non-canonical keys,
//line /usr/local/go/src/net/mail/message.go:199
// access the map directly.
//line /usr/local/go/src/net/mail/message.go:205
func (h Header) Get(key string) string {
//line /usr/local/go/src/net/mail/message.go:205
	_go_fuzz_dep_.CoverTab[171111]++
							return textproto.MIMEHeader(h).Get(key)
//line /usr/local/go/src/net/mail/message.go:206
	// _ = "end of CoverTab[171111]"
}

var ErrHeaderNotPresent = errors.New("mail: header not in message")

// Date parses the Date header field.
func (h Header) Date() (time.Time, error) {
//line /usr/local/go/src/net/mail/message.go:212
	_go_fuzz_dep_.CoverTab[171112]++
							hdr := h.Get("Date")
							if hdr == "" {
//line /usr/local/go/src/net/mail/message.go:214
		_go_fuzz_dep_.CoverTab[171114]++
								return time.Time{}, ErrHeaderNotPresent
//line /usr/local/go/src/net/mail/message.go:215
		// _ = "end of CoverTab[171114]"
	} else {
//line /usr/local/go/src/net/mail/message.go:216
		_go_fuzz_dep_.CoverTab[171115]++
//line /usr/local/go/src/net/mail/message.go:216
		// _ = "end of CoverTab[171115]"
//line /usr/local/go/src/net/mail/message.go:216
	}
//line /usr/local/go/src/net/mail/message.go:216
	// _ = "end of CoverTab[171112]"
//line /usr/local/go/src/net/mail/message.go:216
	_go_fuzz_dep_.CoverTab[171113]++
							return ParseDate(hdr)
//line /usr/local/go/src/net/mail/message.go:217
	// _ = "end of CoverTab[171113]"
}

// AddressList parses the named header field as a list of addresses.
func (h Header) AddressList(key string) ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:221
	_go_fuzz_dep_.CoverTab[171116]++
							hdr := h.Get(key)
							if hdr == "" {
//line /usr/local/go/src/net/mail/message.go:223
		_go_fuzz_dep_.CoverTab[171118]++
								return nil, ErrHeaderNotPresent
//line /usr/local/go/src/net/mail/message.go:224
		// _ = "end of CoverTab[171118]"
	} else {
//line /usr/local/go/src/net/mail/message.go:225
		_go_fuzz_dep_.CoverTab[171119]++
//line /usr/local/go/src/net/mail/message.go:225
		// _ = "end of CoverTab[171119]"
//line /usr/local/go/src/net/mail/message.go:225
	}
//line /usr/local/go/src/net/mail/message.go:225
	// _ = "end of CoverTab[171116]"
//line /usr/local/go/src/net/mail/message.go:225
	_go_fuzz_dep_.CoverTab[171117]++
							return ParseAddressList(hdr)
//line /usr/local/go/src/net/mail/message.go:226
	// _ = "end of CoverTab[171117]"
}

// Address represents a single mail address.
//line /usr/local/go/src/net/mail/message.go:229
// An address such as "Barry Gibbs <bg@example.com>" is represented
//line /usr/local/go/src/net/mail/message.go:229
// as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.
//line /usr/local/go/src/net/mail/message.go:232
type Address struct {
	Name	string	// Proper name; may be empty.
	Address	string	// user@domain
}

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(address string) (*Address, error) {
//line /usr/local/go/src/net/mail/message.go:238
	_go_fuzz_dep_.CoverTab[171120]++
							return (&addrParser{s: address}).parseSingleAddress()
//line /usr/local/go/src/net/mail/message.go:239
	// _ = "end of CoverTab[171120]"
}

// ParseAddressList parses the given string as a list of addresses.
func ParseAddressList(list string) ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:243
	_go_fuzz_dep_.CoverTab[171121]++
							return (&addrParser{s: list}).parseAddressList()
//line /usr/local/go/src/net/mail/message.go:244
	// _ = "end of CoverTab[171121]"
}

// An AddressParser is an RFC 5322 address parser.
type AddressParser struct {
	// WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
	WordDecoder *mime.WordDecoder
}

// Parse parses a single RFC 5322 address of the
//line /usr/local/go/src/net/mail/message.go:253
// form "Gogh Fir <gf@example.com>" or "foo@example.com".
//line /usr/local/go/src/net/mail/message.go:255
func (p *AddressParser) Parse(address string) (*Address, error) {
//line /usr/local/go/src/net/mail/message.go:255
	_go_fuzz_dep_.CoverTab[171122]++
							return (&addrParser{s: address, dec: p.WordDecoder}).parseSingleAddress()
//line /usr/local/go/src/net/mail/message.go:256
	// _ = "end of CoverTab[171122]"
}

// ParseList parses the given string as a list of comma-separated addresses
//line /usr/local/go/src/net/mail/message.go:259
// of the form "Gogh Fir <gf@example.com>" or "foo@example.com".
//line /usr/local/go/src/net/mail/message.go:261
func (p *AddressParser) ParseList(list string) ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:261
	_go_fuzz_dep_.CoverTab[171123]++
							return (&addrParser{s: list, dec: p.WordDecoder}).parseAddressList()
//line /usr/local/go/src/net/mail/message.go:262
	// _ = "end of CoverTab[171123]"
}

// String formats the address as a valid RFC 5322 address.
//line /usr/local/go/src/net/mail/message.go:265
// If the address's name contains non-ASCII characters
//line /usr/local/go/src/net/mail/message.go:265
// the name will be rendered according to RFC 2047.
//line /usr/local/go/src/net/mail/message.go:268
func (a *Address) String() string {
//line /usr/local/go/src/net/mail/message.go:268
	_go_fuzz_dep_.CoverTab[171124]++

							at := strings.LastIndex(a.Address, "@")
							var local, domain string
							if at < 0 {
//line /usr/local/go/src/net/mail/message.go:272
		_go_fuzz_dep_.CoverTab[171132]++

//line /usr/local/go/src/net/mail/message.go:275
		local = a.Address
//line /usr/local/go/src/net/mail/message.go:275
		// _ = "end of CoverTab[171132]"
	} else {
//line /usr/local/go/src/net/mail/message.go:276
		_go_fuzz_dep_.CoverTab[171133]++
								local, domain = a.Address[:at], a.Address[at+1:]
//line /usr/local/go/src/net/mail/message.go:277
		// _ = "end of CoverTab[171133]"
	}
//line /usr/local/go/src/net/mail/message.go:278
	// _ = "end of CoverTab[171124]"
//line /usr/local/go/src/net/mail/message.go:278
	_go_fuzz_dep_.CoverTab[171125]++

//line /usr/local/go/src/net/mail/message.go:281
	quoteLocal := false
	for i, r := range local {
//line /usr/local/go/src/net/mail/message.go:282
		_go_fuzz_dep_.CoverTab[171134]++
								if isAtext(r, false, false) {
//line /usr/local/go/src/net/mail/message.go:283
			_go_fuzz_dep_.CoverTab[171137]++
									continue
//line /usr/local/go/src/net/mail/message.go:284
			// _ = "end of CoverTab[171137]"
		} else {
//line /usr/local/go/src/net/mail/message.go:285
			_go_fuzz_dep_.CoverTab[171138]++
//line /usr/local/go/src/net/mail/message.go:285
			// _ = "end of CoverTab[171138]"
//line /usr/local/go/src/net/mail/message.go:285
		}
//line /usr/local/go/src/net/mail/message.go:285
		// _ = "end of CoverTab[171134]"
//line /usr/local/go/src/net/mail/message.go:285
		_go_fuzz_dep_.CoverTab[171135]++
								if r == '.' {
//line /usr/local/go/src/net/mail/message.go:286
			_go_fuzz_dep_.CoverTab[171139]++

//line /usr/local/go/src/net/mail/message.go:290
			if i > 0 && func() bool {
//line /usr/local/go/src/net/mail/message.go:290
				_go_fuzz_dep_.CoverTab[171140]++
//line /usr/local/go/src/net/mail/message.go:290
				return local[i-1] != '.'
//line /usr/local/go/src/net/mail/message.go:290
				// _ = "end of CoverTab[171140]"
//line /usr/local/go/src/net/mail/message.go:290
			}() && func() bool {
//line /usr/local/go/src/net/mail/message.go:290
				_go_fuzz_dep_.CoverTab[171141]++
//line /usr/local/go/src/net/mail/message.go:290
				return i < len(local)-1
//line /usr/local/go/src/net/mail/message.go:290
				// _ = "end of CoverTab[171141]"
//line /usr/local/go/src/net/mail/message.go:290
			}() {
//line /usr/local/go/src/net/mail/message.go:290
				_go_fuzz_dep_.CoverTab[171142]++
										continue
//line /usr/local/go/src/net/mail/message.go:291
				// _ = "end of CoverTab[171142]"
			} else {
//line /usr/local/go/src/net/mail/message.go:292
				_go_fuzz_dep_.CoverTab[171143]++
//line /usr/local/go/src/net/mail/message.go:292
				// _ = "end of CoverTab[171143]"
//line /usr/local/go/src/net/mail/message.go:292
			}
//line /usr/local/go/src/net/mail/message.go:292
			// _ = "end of CoverTab[171139]"
		} else {
//line /usr/local/go/src/net/mail/message.go:293
			_go_fuzz_dep_.CoverTab[171144]++
//line /usr/local/go/src/net/mail/message.go:293
			// _ = "end of CoverTab[171144]"
//line /usr/local/go/src/net/mail/message.go:293
		}
//line /usr/local/go/src/net/mail/message.go:293
		// _ = "end of CoverTab[171135]"
//line /usr/local/go/src/net/mail/message.go:293
		_go_fuzz_dep_.CoverTab[171136]++
								quoteLocal = true
								break
//line /usr/local/go/src/net/mail/message.go:295
		// _ = "end of CoverTab[171136]"
	}
//line /usr/local/go/src/net/mail/message.go:296
	// _ = "end of CoverTab[171125]"
//line /usr/local/go/src/net/mail/message.go:296
	_go_fuzz_dep_.CoverTab[171126]++
							if quoteLocal {
//line /usr/local/go/src/net/mail/message.go:297
		_go_fuzz_dep_.CoverTab[171145]++
								local = quoteString(local)
//line /usr/local/go/src/net/mail/message.go:298
		// _ = "end of CoverTab[171145]"

	} else {
//line /usr/local/go/src/net/mail/message.go:300
		_go_fuzz_dep_.CoverTab[171146]++
//line /usr/local/go/src/net/mail/message.go:300
		// _ = "end of CoverTab[171146]"
//line /usr/local/go/src/net/mail/message.go:300
	}
//line /usr/local/go/src/net/mail/message.go:300
	// _ = "end of CoverTab[171126]"
//line /usr/local/go/src/net/mail/message.go:300
	_go_fuzz_dep_.CoverTab[171127]++

							s := "<" + local + "@" + domain + ">"

							if a.Name == "" {
//line /usr/local/go/src/net/mail/message.go:304
		_go_fuzz_dep_.CoverTab[171147]++
								return s
//line /usr/local/go/src/net/mail/message.go:305
		// _ = "end of CoverTab[171147]"
	} else {
//line /usr/local/go/src/net/mail/message.go:306
		_go_fuzz_dep_.CoverTab[171148]++
//line /usr/local/go/src/net/mail/message.go:306
		// _ = "end of CoverTab[171148]"
//line /usr/local/go/src/net/mail/message.go:306
	}
//line /usr/local/go/src/net/mail/message.go:306
	// _ = "end of CoverTab[171127]"
//line /usr/local/go/src/net/mail/message.go:306
	_go_fuzz_dep_.CoverTab[171128]++

//line /usr/local/go/src/net/mail/message.go:309
	allPrintable := true
	for _, r := range a.Name {
//line /usr/local/go/src/net/mail/message.go:310
		_go_fuzz_dep_.CoverTab[171149]++

//line /usr/local/go/src/net/mail/message.go:313
		if !isVchar(r) && func() bool {
//line /usr/local/go/src/net/mail/message.go:313
			_go_fuzz_dep_.CoverTab[171150]++
//line /usr/local/go/src/net/mail/message.go:313
			return !isWSP(r)
//line /usr/local/go/src/net/mail/message.go:313
			// _ = "end of CoverTab[171150]"
//line /usr/local/go/src/net/mail/message.go:313
		}() || func() bool {
//line /usr/local/go/src/net/mail/message.go:313
			_go_fuzz_dep_.CoverTab[171151]++
//line /usr/local/go/src/net/mail/message.go:313
			return isMultibyte(r)
//line /usr/local/go/src/net/mail/message.go:313
			// _ = "end of CoverTab[171151]"
//line /usr/local/go/src/net/mail/message.go:313
		}() {
//line /usr/local/go/src/net/mail/message.go:313
			_go_fuzz_dep_.CoverTab[171152]++
									allPrintable = false
									break
//line /usr/local/go/src/net/mail/message.go:315
			// _ = "end of CoverTab[171152]"
		} else {
//line /usr/local/go/src/net/mail/message.go:316
			_go_fuzz_dep_.CoverTab[171153]++
//line /usr/local/go/src/net/mail/message.go:316
			// _ = "end of CoverTab[171153]"
//line /usr/local/go/src/net/mail/message.go:316
		}
//line /usr/local/go/src/net/mail/message.go:316
		// _ = "end of CoverTab[171149]"
	}
//line /usr/local/go/src/net/mail/message.go:317
	// _ = "end of CoverTab[171128]"
//line /usr/local/go/src/net/mail/message.go:317
	_go_fuzz_dep_.CoverTab[171129]++
							if allPrintable {
//line /usr/local/go/src/net/mail/message.go:318
		_go_fuzz_dep_.CoverTab[171154]++
								return quoteString(a.Name) + " " + s
//line /usr/local/go/src/net/mail/message.go:319
		// _ = "end of CoverTab[171154]"
	} else {
//line /usr/local/go/src/net/mail/message.go:320
		_go_fuzz_dep_.CoverTab[171155]++
//line /usr/local/go/src/net/mail/message.go:320
		// _ = "end of CoverTab[171155]"
//line /usr/local/go/src/net/mail/message.go:320
	}
//line /usr/local/go/src/net/mail/message.go:320
	// _ = "end of CoverTab[171129]"
//line /usr/local/go/src/net/mail/message.go:320
	_go_fuzz_dep_.CoverTab[171130]++

//line /usr/local/go/src/net/mail/message.go:325
	if strings.ContainsAny(a.Name, "\"#$%&'(),.:;<>@[]^`{|}~") {
//line /usr/local/go/src/net/mail/message.go:325
		_go_fuzz_dep_.CoverTab[171156]++
								return mime.BEncoding.Encode("utf-8", a.Name) + " " + s
//line /usr/local/go/src/net/mail/message.go:326
		// _ = "end of CoverTab[171156]"
	} else {
//line /usr/local/go/src/net/mail/message.go:327
		_go_fuzz_dep_.CoverTab[171157]++
//line /usr/local/go/src/net/mail/message.go:327
		// _ = "end of CoverTab[171157]"
//line /usr/local/go/src/net/mail/message.go:327
	}
//line /usr/local/go/src/net/mail/message.go:327
	// _ = "end of CoverTab[171130]"
//line /usr/local/go/src/net/mail/message.go:327
	_go_fuzz_dep_.CoverTab[171131]++
							return mime.QEncoding.Encode("utf-8", a.Name) + " " + s
//line /usr/local/go/src/net/mail/message.go:328
	// _ = "end of CoverTab[171131]"
}

type addrParser struct {
	s	string
	dec	*mime.WordDecoder	// may be nil
}

func (p *addrParser) parseAddressList() ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:336
	_go_fuzz_dep_.CoverTab[171158]++
							var list []*Address
							for {
//line /usr/local/go/src/net/mail/message.go:338
		_go_fuzz_dep_.CoverTab[171160]++
								p.skipSpace()

//line /usr/local/go/src/net/mail/message.go:342
		if p.consume(',') {
//line /usr/local/go/src/net/mail/message.go:342
			_go_fuzz_dep_.CoverTab[171167]++
									continue
//line /usr/local/go/src/net/mail/message.go:343
			// _ = "end of CoverTab[171167]"
		} else {
//line /usr/local/go/src/net/mail/message.go:344
			_go_fuzz_dep_.CoverTab[171168]++
//line /usr/local/go/src/net/mail/message.go:344
			// _ = "end of CoverTab[171168]"
//line /usr/local/go/src/net/mail/message.go:344
		}
//line /usr/local/go/src/net/mail/message.go:344
		// _ = "end of CoverTab[171160]"
//line /usr/local/go/src/net/mail/message.go:344
		_go_fuzz_dep_.CoverTab[171161]++

								addrs, err := p.parseAddress(true)
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:347
			_go_fuzz_dep_.CoverTab[171169]++
									return nil, err
//line /usr/local/go/src/net/mail/message.go:348
			// _ = "end of CoverTab[171169]"
		} else {
//line /usr/local/go/src/net/mail/message.go:349
			_go_fuzz_dep_.CoverTab[171170]++
//line /usr/local/go/src/net/mail/message.go:349
			// _ = "end of CoverTab[171170]"
//line /usr/local/go/src/net/mail/message.go:349
		}
//line /usr/local/go/src/net/mail/message.go:349
		// _ = "end of CoverTab[171161]"
//line /usr/local/go/src/net/mail/message.go:349
		_go_fuzz_dep_.CoverTab[171162]++
								list = append(list, addrs...)

								if !p.skipCFWS() {
//line /usr/local/go/src/net/mail/message.go:352
			_go_fuzz_dep_.CoverTab[171171]++
									return nil, errors.New("mail: misformatted parenthetical comment")
//line /usr/local/go/src/net/mail/message.go:353
			// _ = "end of CoverTab[171171]"
		} else {
//line /usr/local/go/src/net/mail/message.go:354
			_go_fuzz_dep_.CoverTab[171172]++
//line /usr/local/go/src/net/mail/message.go:354
			// _ = "end of CoverTab[171172]"
//line /usr/local/go/src/net/mail/message.go:354
		}
//line /usr/local/go/src/net/mail/message.go:354
		// _ = "end of CoverTab[171162]"
//line /usr/local/go/src/net/mail/message.go:354
		_go_fuzz_dep_.CoverTab[171163]++
								if p.empty() {
//line /usr/local/go/src/net/mail/message.go:355
			_go_fuzz_dep_.CoverTab[171173]++
									break
//line /usr/local/go/src/net/mail/message.go:356
			// _ = "end of CoverTab[171173]"
		} else {
//line /usr/local/go/src/net/mail/message.go:357
			_go_fuzz_dep_.CoverTab[171174]++
//line /usr/local/go/src/net/mail/message.go:357
			// _ = "end of CoverTab[171174]"
//line /usr/local/go/src/net/mail/message.go:357
		}
//line /usr/local/go/src/net/mail/message.go:357
		// _ = "end of CoverTab[171163]"
//line /usr/local/go/src/net/mail/message.go:357
		_go_fuzz_dep_.CoverTab[171164]++
								if p.peek() != ',' {
//line /usr/local/go/src/net/mail/message.go:358
			_go_fuzz_dep_.CoverTab[171175]++
									return nil, errors.New("mail: expected comma")
//line /usr/local/go/src/net/mail/message.go:359
			// _ = "end of CoverTab[171175]"
		} else {
//line /usr/local/go/src/net/mail/message.go:360
			_go_fuzz_dep_.CoverTab[171176]++
//line /usr/local/go/src/net/mail/message.go:360
			// _ = "end of CoverTab[171176]"
//line /usr/local/go/src/net/mail/message.go:360
		}
//line /usr/local/go/src/net/mail/message.go:360
		// _ = "end of CoverTab[171164]"
//line /usr/local/go/src/net/mail/message.go:360
		_go_fuzz_dep_.CoverTab[171165]++

//line /usr/local/go/src/net/mail/message.go:363
		for p.consume(',') {
//line /usr/local/go/src/net/mail/message.go:363
			_go_fuzz_dep_.CoverTab[171177]++
									p.skipSpace()
//line /usr/local/go/src/net/mail/message.go:364
			// _ = "end of CoverTab[171177]"
		}
//line /usr/local/go/src/net/mail/message.go:365
		// _ = "end of CoverTab[171165]"
//line /usr/local/go/src/net/mail/message.go:365
		_go_fuzz_dep_.CoverTab[171166]++
								if p.empty() {
//line /usr/local/go/src/net/mail/message.go:366
			_go_fuzz_dep_.CoverTab[171178]++
									break
//line /usr/local/go/src/net/mail/message.go:367
			// _ = "end of CoverTab[171178]"
		} else {
//line /usr/local/go/src/net/mail/message.go:368
			_go_fuzz_dep_.CoverTab[171179]++
//line /usr/local/go/src/net/mail/message.go:368
			// _ = "end of CoverTab[171179]"
//line /usr/local/go/src/net/mail/message.go:368
		}
//line /usr/local/go/src/net/mail/message.go:368
		// _ = "end of CoverTab[171166]"
	}
//line /usr/local/go/src/net/mail/message.go:369
	// _ = "end of CoverTab[171158]"
//line /usr/local/go/src/net/mail/message.go:369
	_go_fuzz_dep_.CoverTab[171159]++
							return list, nil
//line /usr/local/go/src/net/mail/message.go:370
	// _ = "end of CoverTab[171159]"
}

func (p *addrParser) parseSingleAddress() (*Address, error) {
//line /usr/local/go/src/net/mail/message.go:373
	_go_fuzz_dep_.CoverTab[171180]++
							addrs, err := p.parseAddress(true)
							if err != nil {
//line /usr/local/go/src/net/mail/message.go:375
		_go_fuzz_dep_.CoverTab[171186]++
								return nil, err
//line /usr/local/go/src/net/mail/message.go:376
		// _ = "end of CoverTab[171186]"
	} else {
//line /usr/local/go/src/net/mail/message.go:377
		_go_fuzz_dep_.CoverTab[171187]++
//line /usr/local/go/src/net/mail/message.go:377
		// _ = "end of CoverTab[171187]"
//line /usr/local/go/src/net/mail/message.go:377
	}
//line /usr/local/go/src/net/mail/message.go:377
	// _ = "end of CoverTab[171180]"
//line /usr/local/go/src/net/mail/message.go:377
	_go_fuzz_dep_.CoverTab[171181]++
							if !p.skipCFWS() {
//line /usr/local/go/src/net/mail/message.go:378
		_go_fuzz_dep_.CoverTab[171188]++
								return nil, errors.New("mail: misformatted parenthetical comment")
//line /usr/local/go/src/net/mail/message.go:379
		// _ = "end of CoverTab[171188]"
	} else {
//line /usr/local/go/src/net/mail/message.go:380
		_go_fuzz_dep_.CoverTab[171189]++
//line /usr/local/go/src/net/mail/message.go:380
		// _ = "end of CoverTab[171189]"
//line /usr/local/go/src/net/mail/message.go:380
	}
//line /usr/local/go/src/net/mail/message.go:380
	// _ = "end of CoverTab[171181]"
//line /usr/local/go/src/net/mail/message.go:380
	_go_fuzz_dep_.CoverTab[171182]++
							if !p.empty() {
//line /usr/local/go/src/net/mail/message.go:381
		_go_fuzz_dep_.CoverTab[171190]++
								return nil, fmt.Errorf("mail: expected single address, got %q", p.s)
//line /usr/local/go/src/net/mail/message.go:382
		// _ = "end of CoverTab[171190]"
	} else {
//line /usr/local/go/src/net/mail/message.go:383
		_go_fuzz_dep_.CoverTab[171191]++
//line /usr/local/go/src/net/mail/message.go:383
		// _ = "end of CoverTab[171191]"
//line /usr/local/go/src/net/mail/message.go:383
	}
//line /usr/local/go/src/net/mail/message.go:383
	// _ = "end of CoverTab[171182]"
//line /usr/local/go/src/net/mail/message.go:383
	_go_fuzz_dep_.CoverTab[171183]++
							if len(addrs) == 0 {
//line /usr/local/go/src/net/mail/message.go:384
		_go_fuzz_dep_.CoverTab[171192]++
								return nil, errors.New("mail: empty group")
//line /usr/local/go/src/net/mail/message.go:385
		// _ = "end of CoverTab[171192]"
	} else {
//line /usr/local/go/src/net/mail/message.go:386
		_go_fuzz_dep_.CoverTab[171193]++
//line /usr/local/go/src/net/mail/message.go:386
		// _ = "end of CoverTab[171193]"
//line /usr/local/go/src/net/mail/message.go:386
	}
//line /usr/local/go/src/net/mail/message.go:386
	// _ = "end of CoverTab[171183]"
//line /usr/local/go/src/net/mail/message.go:386
	_go_fuzz_dep_.CoverTab[171184]++
							if len(addrs) > 1 {
//line /usr/local/go/src/net/mail/message.go:387
		_go_fuzz_dep_.CoverTab[171194]++
								return nil, errors.New("mail: group with multiple addresses")
//line /usr/local/go/src/net/mail/message.go:388
		// _ = "end of CoverTab[171194]"
	} else {
//line /usr/local/go/src/net/mail/message.go:389
		_go_fuzz_dep_.CoverTab[171195]++
//line /usr/local/go/src/net/mail/message.go:389
		// _ = "end of CoverTab[171195]"
//line /usr/local/go/src/net/mail/message.go:389
	}
//line /usr/local/go/src/net/mail/message.go:389
	// _ = "end of CoverTab[171184]"
//line /usr/local/go/src/net/mail/message.go:389
	_go_fuzz_dep_.CoverTab[171185]++
							return addrs[0], nil
//line /usr/local/go/src/net/mail/message.go:390
	// _ = "end of CoverTab[171185]"
}

// parseAddress parses a single RFC 5322 address at the start of p.
func (p *addrParser) parseAddress(handleGroup bool) ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:394
	_go_fuzz_dep_.CoverTab[171196]++
							debug.Printf("parseAddress: %q", p.s)
							p.skipSpace()
							if p.empty() {
//line /usr/local/go/src/net/mail/message.go:397
		_go_fuzz_dep_.CoverTab[171204]++
								return nil, errors.New("mail: no address")
//line /usr/local/go/src/net/mail/message.go:398
		// _ = "end of CoverTab[171204]"
	} else {
//line /usr/local/go/src/net/mail/message.go:399
		_go_fuzz_dep_.CoverTab[171205]++
//line /usr/local/go/src/net/mail/message.go:399
		// _ = "end of CoverTab[171205]"
//line /usr/local/go/src/net/mail/message.go:399
	}
//line /usr/local/go/src/net/mail/message.go:399
	// _ = "end of CoverTab[171196]"
//line /usr/local/go/src/net/mail/message.go:399
	_go_fuzz_dep_.CoverTab[171197]++

//line /usr/local/go/src/net/mail/message.go:408
	spec, err := p.consumeAddrSpec()
	if err == nil {
//line /usr/local/go/src/net/mail/message.go:409
		_go_fuzz_dep_.CoverTab[171206]++
								var displayName string
								p.skipSpace()
								if !p.empty() && func() bool {
//line /usr/local/go/src/net/mail/message.go:412
			_go_fuzz_dep_.CoverTab[171208]++
//line /usr/local/go/src/net/mail/message.go:412
			return p.peek() == '('
//line /usr/local/go/src/net/mail/message.go:412
			// _ = "end of CoverTab[171208]"
//line /usr/local/go/src/net/mail/message.go:412
		}() {
//line /usr/local/go/src/net/mail/message.go:412
			_go_fuzz_dep_.CoverTab[171209]++
									displayName, err = p.consumeDisplayNameComment()
									if err != nil {
//line /usr/local/go/src/net/mail/message.go:414
				_go_fuzz_dep_.CoverTab[171210]++
										return nil, err
//line /usr/local/go/src/net/mail/message.go:415
				// _ = "end of CoverTab[171210]"
			} else {
//line /usr/local/go/src/net/mail/message.go:416
				_go_fuzz_dep_.CoverTab[171211]++
//line /usr/local/go/src/net/mail/message.go:416
				// _ = "end of CoverTab[171211]"
//line /usr/local/go/src/net/mail/message.go:416
			}
//line /usr/local/go/src/net/mail/message.go:416
			// _ = "end of CoverTab[171209]"
		} else {
//line /usr/local/go/src/net/mail/message.go:417
			_go_fuzz_dep_.CoverTab[171212]++
//line /usr/local/go/src/net/mail/message.go:417
			// _ = "end of CoverTab[171212]"
//line /usr/local/go/src/net/mail/message.go:417
		}
//line /usr/local/go/src/net/mail/message.go:417
		// _ = "end of CoverTab[171206]"
//line /usr/local/go/src/net/mail/message.go:417
		_go_fuzz_dep_.CoverTab[171207]++

								return []*Address{{
			Name:		displayName,
			Address:	spec,
		}}, err
//line /usr/local/go/src/net/mail/message.go:422
		// _ = "end of CoverTab[171207]"
	} else {
//line /usr/local/go/src/net/mail/message.go:423
		_go_fuzz_dep_.CoverTab[171213]++
//line /usr/local/go/src/net/mail/message.go:423
		// _ = "end of CoverTab[171213]"
//line /usr/local/go/src/net/mail/message.go:423
	}
//line /usr/local/go/src/net/mail/message.go:423
	// _ = "end of CoverTab[171197]"
//line /usr/local/go/src/net/mail/message.go:423
	_go_fuzz_dep_.CoverTab[171198]++
							debug.Printf("parseAddress: not an addr-spec: %v", err)
							debug.Printf("parseAddress: state is now %q", p.s)

	// display-name
	var displayName string
	if p.peek() != '<' {
//line /usr/local/go/src/net/mail/message.go:429
		_go_fuzz_dep_.CoverTab[171214]++
								displayName, err = p.consumePhrase()
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:431
			_go_fuzz_dep_.CoverTab[171215]++
									return nil, err
//line /usr/local/go/src/net/mail/message.go:432
			// _ = "end of CoverTab[171215]"
		} else {
//line /usr/local/go/src/net/mail/message.go:433
			_go_fuzz_dep_.CoverTab[171216]++
//line /usr/local/go/src/net/mail/message.go:433
			// _ = "end of CoverTab[171216]"
//line /usr/local/go/src/net/mail/message.go:433
		}
//line /usr/local/go/src/net/mail/message.go:433
		// _ = "end of CoverTab[171214]"
	} else {
//line /usr/local/go/src/net/mail/message.go:434
		_go_fuzz_dep_.CoverTab[171217]++
//line /usr/local/go/src/net/mail/message.go:434
		// _ = "end of CoverTab[171217]"
//line /usr/local/go/src/net/mail/message.go:434
	}
//line /usr/local/go/src/net/mail/message.go:434
	// _ = "end of CoverTab[171198]"
//line /usr/local/go/src/net/mail/message.go:434
	_go_fuzz_dep_.CoverTab[171199]++
							debug.Printf("parseAddress: displayName=%q", displayName)

							p.skipSpace()
							if handleGroup {
//line /usr/local/go/src/net/mail/message.go:438
		_go_fuzz_dep_.CoverTab[171218]++
								if p.consume(':') {
//line /usr/local/go/src/net/mail/message.go:439
			_go_fuzz_dep_.CoverTab[171219]++
									return p.consumeGroupList()
//line /usr/local/go/src/net/mail/message.go:440
			// _ = "end of CoverTab[171219]"
		} else {
//line /usr/local/go/src/net/mail/message.go:441
			_go_fuzz_dep_.CoverTab[171220]++
//line /usr/local/go/src/net/mail/message.go:441
			// _ = "end of CoverTab[171220]"
//line /usr/local/go/src/net/mail/message.go:441
		}
//line /usr/local/go/src/net/mail/message.go:441
		// _ = "end of CoverTab[171218]"
	} else {
//line /usr/local/go/src/net/mail/message.go:442
		_go_fuzz_dep_.CoverTab[171221]++
//line /usr/local/go/src/net/mail/message.go:442
		// _ = "end of CoverTab[171221]"
//line /usr/local/go/src/net/mail/message.go:442
	}
//line /usr/local/go/src/net/mail/message.go:442
	// _ = "end of CoverTab[171199]"
//line /usr/local/go/src/net/mail/message.go:442
	_go_fuzz_dep_.CoverTab[171200]++

							if !p.consume('<') {
//line /usr/local/go/src/net/mail/message.go:444
		_go_fuzz_dep_.CoverTab[171222]++
								atext := true
								for _, r := range displayName {
//line /usr/local/go/src/net/mail/message.go:446
			_go_fuzz_dep_.CoverTab[171225]++
									if !isAtext(r, true, false) {
//line /usr/local/go/src/net/mail/message.go:447
				_go_fuzz_dep_.CoverTab[171226]++
										atext = false
										break
//line /usr/local/go/src/net/mail/message.go:449
				// _ = "end of CoverTab[171226]"
			} else {
//line /usr/local/go/src/net/mail/message.go:450
				_go_fuzz_dep_.CoverTab[171227]++
//line /usr/local/go/src/net/mail/message.go:450
				// _ = "end of CoverTab[171227]"
//line /usr/local/go/src/net/mail/message.go:450
			}
//line /usr/local/go/src/net/mail/message.go:450
			// _ = "end of CoverTab[171225]"
		}
//line /usr/local/go/src/net/mail/message.go:451
		// _ = "end of CoverTab[171222]"
//line /usr/local/go/src/net/mail/message.go:451
		_go_fuzz_dep_.CoverTab[171223]++
								if atext {
//line /usr/local/go/src/net/mail/message.go:452
			_go_fuzz_dep_.CoverTab[171228]++

//line /usr/local/go/src/net/mail/message.go:455
			return nil, errors.New("mail: missing '@' or angle-addr")
//line /usr/local/go/src/net/mail/message.go:455
			// _ = "end of CoverTab[171228]"
		} else {
//line /usr/local/go/src/net/mail/message.go:456
			_go_fuzz_dep_.CoverTab[171229]++
//line /usr/local/go/src/net/mail/message.go:456
			// _ = "end of CoverTab[171229]"
//line /usr/local/go/src/net/mail/message.go:456
		}
//line /usr/local/go/src/net/mail/message.go:456
		// _ = "end of CoverTab[171223]"
//line /usr/local/go/src/net/mail/message.go:456
		_go_fuzz_dep_.CoverTab[171224]++

//line /usr/local/go/src/net/mail/message.go:460
		return nil, errors.New("mail: no angle-addr")
//line /usr/local/go/src/net/mail/message.go:460
		// _ = "end of CoverTab[171224]"
	} else {
//line /usr/local/go/src/net/mail/message.go:461
		_go_fuzz_dep_.CoverTab[171230]++
//line /usr/local/go/src/net/mail/message.go:461
		// _ = "end of CoverTab[171230]"
//line /usr/local/go/src/net/mail/message.go:461
	}
//line /usr/local/go/src/net/mail/message.go:461
	// _ = "end of CoverTab[171200]"
//line /usr/local/go/src/net/mail/message.go:461
	_go_fuzz_dep_.CoverTab[171201]++
							spec, err = p.consumeAddrSpec()
							if err != nil {
//line /usr/local/go/src/net/mail/message.go:463
		_go_fuzz_dep_.CoverTab[171231]++
								return nil, err
//line /usr/local/go/src/net/mail/message.go:464
		// _ = "end of CoverTab[171231]"
	} else {
//line /usr/local/go/src/net/mail/message.go:465
		_go_fuzz_dep_.CoverTab[171232]++
//line /usr/local/go/src/net/mail/message.go:465
		// _ = "end of CoverTab[171232]"
//line /usr/local/go/src/net/mail/message.go:465
	}
//line /usr/local/go/src/net/mail/message.go:465
	// _ = "end of CoverTab[171201]"
//line /usr/local/go/src/net/mail/message.go:465
	_go_fuzz_dep_.CoverTab[171202]++
							if !p.consume('>') {
//line /usr/local/go/src/net/mail/message.go:466
		_go_fuzz_dep_.CoverTab[171233]++
								return nil, errors.New("mail: unclosed angle-addr")
//line /usr/local/go/src/net/mail/message.go:467
		// _ = "end of CoverTab[171233]"
	} else {
//line /usr/local/go/src/net/mail/message.go:468
		_go_fuzz_dep_.CoverTab[171234]++
//line /usr/local/go/src/net/mail/message.go:468
		// _ = "end of CoverTab[171234]"
//line /usr/local/go/src/net/mail/message.go:468
	}
//line /usr/local/go/src/net/mail/message.go:468
	// _ = "end of CoverTab[171202]"
//line /usr/local/go/src/net/mail/message.go:468
	_go_fuzz_dep_.CoverTab[171203]++
							debug.Printf("parseAddress: spec=%q", spec)

							return []*Address{{
		Name:		displayName,
		Address:	spec,
	}}, nil
//line /usr/local/go/src/net/mail/message.go:474
	// _ = "end of CoverTab[171203]"
}

func (p *addrParser) consumeGroupList() ([]*Address, error) {
//line /usr/local/go/src/net/mail/message.go:477
	_go_fuzz_dep_.CoverTab[171235]++
							var group []*Address

							p.skipSpace()
							if p.consume(';') {
//line /usr/local/go/src/net/mail/message.go:481
		_go_fuzz_dep_.CoverTab[171238]++
								p.skipCFWS()
								return group, nil
//line /usr/local/go/src/net/mail/message.go:483
		// _ = "end of CoverTab[171238]"
	} else {
//line /usr/local/go/src/net/mail/message.go:484
		_go_fuzz_dep_.CoverTab[171239]++
//line /usr/local/go/src/net/mail/message.go:484
		// _ = "end of CoverTab[171239]"
//line /usr/local/go/src/net/mail/message.go:484
	}
//line /usr/local/go/src/net/mail/message.go:484
	// _ = "end of CoverTab[171235]"
//line /usr/local/go/src/net/mail/message.go:484
	_go_fuzz_dep_.CoverTab[171236]++

							for {
//line /usr/local/go/src/net/mail/message.go:486
		_go_fuzz_dep_.CoverTab[171240]++
								p.skipSpace()

								addrs, err := p.parseAddress(false)
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:490
			_go_fuzz_dep_.CoverTab[171244]++
									return nil, err
//line /usr/local/go/src/net/mail/message.go:491
			// _ = "end of CoverTab[171244]"
		} else {
//line /usr/local/go/src/net/mail/message.go:492
			_go_fuzz_dep_.CoverTab[171245]++
//line /usr/local/go/src/net/mail/message.go:492
			// _ = "end of CoverTab[171245]"
//line /usr/local/go/src/net/mail/message.go:492
		}
//line /usr/local/go/src/net/mail/message.go:492
		// _ = "end of CoverTab[171240]"
//line /usr/local/go/src/net/mail/message.go:492
		_go_fuzz_dep_.CoverTab[171241]++
								group = append(group, addrs...)

								if !p.skipCFWS() {
//line /usr/local/go/src/net/mail/message.go:495
			_go_fuzz_dep_.CoverTab[171246]++
									return nil, errors.New("mail: misformatted parenthetical comment")
//line /usr/local/go/src/net/mail/message.go:496
			// _ = "end of CoverTab[171246]"
		} else {
//line /usr/local/go/src/net/mail/message.go:497
			_go_fuzz_dep_.CoverTab[171247]++
//line /usr/local/go/src/net/mail/message.go:497
			// _ = "end of CoverTab[171247]"
//line /usr/local/go/src/net/mail/message.go:497
		}
//line /usr/local/go/src/net/mail/message.go:497
		// _ = "end of CoverTab[171241]"
//line /usr/local/go/src/net/mail/message.go:497
		_go_fuzz_dep_.CoverTab[171242]++
								if p.consume(';') {
//line /usr/local/go/src/net/mail/message.go:498
			_go_fuzz_dep_.CoverTab[171248]++
									p.skipCFWS()
									break
//line /usr/local/go/src/net/mail/message.go:500
			// _ = "end of CoverTab[171248]"
		} else {
//line /usr/local/go/src/net/mail/message.go:501
			_go_fuzz_dep_.CoverTab[171249]++
//line /usr/local/go/src/net/mail/message.go:501
			// _ = "end of CoverTab[171249]"
//line /usr/local/go/src/net/mail/message.go:501
		}
//line /usr/local/go/src/net/mail/message.go:501
		// _ = "end of CoverTab[171242]"
//line /usr/local/go/src/net/mail/message.go:501
		_go_fuzz_dep_.CoverTab[171243]++
								if !p.consume(',') {
//line /usr/local/go/src/net/mail/message.go:502
			_go_fuzz_dep_.CoverTab[171250]++
									return nil, errors.New("mail: expected comma")
//line /usr/local/go/src/net/mail/message.go:503
			// _ = "end of CoverTab[171250]"
		} else {
//line /usr/local/go/src/net/mail/message.go:504
			_go_fuzz_dep_.CoverTab[171251]++
//line /usr/local/go/src/net/mail/message.go:504
			// _ = "end of CoverTab[171251]"
//line /usr/local/go/src/net/mail/message.go:504
		}
//line /usr/local/go/src/net/mail/message.go:504
		// _ = "end of CoverTab[171243]"
	}
//line /usr/local/go/src/net/mail/message.go:505
	// _ = "end of CoverTab[171236]"
//line /usr/local/go/src/net/mail/message.go:505
	_go_fuzz_dep_.CoverTab[171237]++
							return group, nil
//line /usr/local/go/src/net/mail/message.go:506
	// _ = "end of CoverTab[171237]"
}

// consumeAddrSpec parses a single RFC 5322 addr-spec at the start of p.
func (p *addrParser) consumeAddrSpec() (spec string, err error) {
//line /usr/local/go/src/net/mail/message.go:510
	_go_fuzz_dep_.CoverTab[171252]++
							debug.Printf("consumeAddrSpec: %q", p.s)

							orig := *p
							defer func() {
//line /usr/local/go/src/net/mail/message.go:514
		_go_fuzz_dep_.CoverTab[171260]++
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:515
			_go_fuzz_dep_.CoverTab[171261]++
									*p = orig
//line /usr/local/go/src/net/mail/message.go:516
			// _ = "end of CoverTab[171261]"
		} else {
//line /usr/local/go/src/net/mail/message.go:517
			_go_fuzz_dep_.CoverTab[171262]++
//line /usr/local/go/src/net/mail/message.go:517
			// _ = "end of CoverTab[171262]"
//line /usr/local/go/src/net/mail/message.go:517
		}
//line /usr/local/go/src/net/mail/message.go:517
		// _ = "end of CoverTab[171260]"
	}()
//line /usr/local/go/src/net/mail/message.go:518
	// _ = "end of CoverTab[171252]"
//line /usr/local/go/src/net/mail/message.go:518
	_go_fuzz_dep_.CoverTab[171253]++

	// local-part = dot-atom / quoted-string
	var localPart string
	p.skipSpace()
	if p.empty() {
//line /usr/local/go/src/net/mail/message.go:523
		_go_fuzz_dep_.CoverTab[171263]++
								return "", errors.New("mail: no addr-spec")
//line /usr/local/go/src/net/mail/message.go:524
		// _ = "end of CoverTab[171263]"
	} else {
//line /usr/local/go/src/net/mail/message.go:525
		_go_fuzz_dep_.CoverTab[171264]++
//line /usr/local/go/src/net/mail/message.go:525
		// _ = "end of CoverTab[171264]"
//line /usr/local/go/src/net/mail/message.go:525
	}
//line /usr/local/go/src/net/mail/message.go:525
	// _ = "end of CoverTab[171253]"
//line /usr/local/go/src/net/mail/message.go:525
	_go_fuzz_dep_.CoverTab[171254]++
							if p.peek() == '"' {
//line /usr/local/go/src/net/mail/message.go:526
		_go_fuzz_dep_.CoverTab[171265]++

								debug.Printf("consumeAddrSpec: parsing quoted-string")
								localPart, err = p.consumeQuotedString()
								if localPart == "" {
//line /usr/local/go/src/net/mail/message.go:530
			_go_fuzz_dep_.CoverTab[171266]++
									err = errors.New("mail: empty quoted string in addr-spec")
//line /usr/local/go/src/net/mail/message.go:531
			// _ = "end of CoverTab[171266]"
		} else {
//line /usr/local/go/src/net/mail/message.go:532
			_go_fuzz_dep_.CoverTab[171267]++
//line /usr/local/go/src/net/mail/message.go:532
			// _ = "end of CoverTab[171267]"
//line /usr/local/go/src/net/mail/message.go:532
		}
//line /usr/local/go/src/net/mail/message.go:532
		// _ = "end of CoverTab[171265]"
	} else {
//line /usr/local/go/src/net/mail/message.go:533
		_go_fuzz_dep_.CoverTab[171268]++

								debug.Printf("consumeAddrSpec: parsing dot-atom")
								localPart, err = p.consumeAtom(true, false)
//line /usr/local/go/src/net/mail/message.go:536
		// _ = "end of CoverTab[171268]"
	}
//line /usr/local/go/src/net/mail/message.go:537
	// _ = "end of CoverTab[171254]"
//line /usr/local/go/src/net/mail/message.go:537
	_go_fuzz_dep_.CoverTab[171255]++
							if err != nil {
//line /usr/local/go/src/net/mail/message.go:538
		_go_fuzz_dep_.CoverTab[171269]++
								debug.Printf("consumeAddrSpec: failed: %v", err)
								return "", err
//line /usr/local/go/src/net/mail/message.go:540
		// _ = "end of CoverTab[171269]"
	} else {
//line /usr/local/go/src/net/mail/message.go:541
		_go_fuzz_dep_.CoverTab[171270]++
//line /usr/local/go/src/net/mail/message.go:541
		// _ = "end of CoverTab[171270]"
//line /usr/local/go/src/net/mail/message.go:541
	}
//line /usr/local/go/src/net/mail/message.go:541
	// _ = "end of CoverTab[171255]"
//line /usr/local/go/src/net/mail/message.go:541
	_go_fuzz_dep_.CoverTab[171256]++

							if !p.consume('@') {
//line /usr/local/go/src/net/mail/message.go:543
		_go_fuzz_dep_.CoverTab[171271]++
								return "", errors.New("mail: missing @ in addr-spec")
//line /usr/local/go/src/net/mail/message.go:544
		// _ = "end of CoverTab[171271]"
	} else {
//line /usr/local/go/src/net/mail/message.go:545
		_go_fuzz_dep_.CoverTab[171272]++
//line /usr/local/go/src/net/mail/message.go:545
		// _ = "end of CoverTab[171272]"
//line /usr/local/go/src/net/mail/message.go:545
	}
//line /usr/local/go/src/net/mail/message.go:545
	// _ = "end of CoverTab[171256]"
//line /usr/local/go/src/net/mail/message.go:545
	_go_fuzz_dep_.CoverTab[171257]++

	// domain = dot-atom / domain-literal
	var domain string
	p.skipSpace()
	if p.empty() {
//line /usr/local/go/src/net/mail/message.go:550
		_go_fuzz_dep_.CoverTab[171273]++
								return "", errors.New("mail: no domain in addr-spec")
//line /usr/local/go/src/net/mail/message.go:551
		// _ = "end of CoverTab[171273]"
	} else {
//line /usr/local/go/src/net/mail/message.go:552
		_go_fuzz_dep_.CoverTab[171274]++
//line /usr/local/go/src/net/mail/message.go:552
		// _ = "end of CoverTab[171274]"
//line /usr/local/go/src/net/mail/message.go:552
	}
//line /usr/local/go/src/net/mail/message.go:552
	// _ = "end of CoverTab[171257]"
//line /usr/local/go/src/net/mail/message.go:552
	_go_fuzz_dep_.CoverTab[171258]++

							domain, err = p.consumeAtom(true, false)
							if err != nil {
//line /usr/local/go/src/net/mail/message.go:555
		_go_fuzz_dep_.CoverTab[171275]++
								return "", err
//line /usr/local/go/src/net/mail/message.go:556
		// _ = "end of CoverTab[171275]"
	} else {
//line /usr/local/go/src/net/mail/message.go:557
		_go_fuzz_dep_.CoverTab[171276]++
//line /usr/local/go/src/net/mail/message.go:557
		// _ = "end of CoverTab[171276]"
//line /usr/local/go/src/net/mail/message.go:557
	}
//line /usr/local/go/src/net/mail/message.go:557
	// _ = "end of CoverTab[171258]"
//line /usr/local/go/src/net/mail/message.go:557
	_go_fuzz_dep_.CoverTab[171259]++

							return localPart + "@" + domain, nil
//line /usr/local/go/src/net/mail/message.go:559
	// _ = "end of CoverTab[171259]"
}

// consumePhrase parses the RFC 5322 phrase at the start of p.
func (p *addrParser) consumePhrase() (phrase string, err error) {
//line /usr/local/go/src/net/mail/message.go:563
	_go_fuzz_dep_.CoverTab[171277]++
							debug.Printf("consumePhrase: [%s]", p.s)
	// phrase = 1*word
	var words []string
	var isPrevEncoded bool
	for {
//line /usr/local/go/src/net/mail/message.go:568
		_go_fuzz_dep_.CoverTab[171280]++
		// word = atom / quoted-string
		var word string
		p.skipSpace()
		if p.empty() {
//line /usr/local/go/src/net/mail/message.go:572
			_go_fuzz_dep_.CoverTab[171285]++
									break
//line /usr/local/go/src/net/mail/message.go:573
			// _ = "end of CoverTab[171285]"
		} else {
//line /usr/local/go/src/net/mail/message.go:574
			_go_fuzz_dep_.CoverTab[171286]++
//line /usr/local/go/src/net/mail/message.go:574
			// _ = "end of CoverTab[171286]"
//line /usr/local/go/src/net/mail/message.go:574
		}
//line /usr/local/go/src/net/mail/message.go:574
		// _ = "end of CoverTab[171280]"
//line /usr/local/go/src/net/mail/message.go:574
		_go_fuzz_dep_.CoverTab[171281]++
								isEncoded := false
								if p.peek() == '"' {
//line /usr/local/go/src/net/mail/message.go:576
			_go_fuzz_dep_.CoverTab[171287]++

									word, err = p.consumeQuotedString()
//line /usr/local/go/src/net/mail/message.go:578
			// _ = "end of CoverTab[171287]"
		} else {
//line /usr/local/go/src/net/mail/message.go:579
			_go_fuzz_dep_.CoverTab[171288]++

//line /usr/local/go/src/net/mail/message.go:583
			word, err = p.consumeAtom(true, true)
			if err == nil {
//line /usr/local/go/src/net/mail/message.go:584
				_go_fuzz_dep_.CoverTab[171289]++
										word, isEncoded, err = p.decodeRFC2047Word(word)
//line /usr/local/go/src/net/mail/message.go:585
				// _ = "end of CoverTab[171289]"
			} else {
//line /usr/local/go/src/net/mail/message.go:586
				_go_fuzz_dep_.CoverTab[171290]++
//line /usr/local/go/src/net/mail/message.go:586
				// _ = "end of CoverTab[171290]"
//line /usr/local/go/src/net/mail/message.go:586
			}
//line /usr/local/go/src/net/mail/message.go:586
			// _ = "end of CoverTab[171288]"
		}
//line /usr/local/go/src/net/mail/message.go:587
		// _ = "end of CoverTab[171281]"
//line /usr/local/go/src/net/mail/message.go:587
		_go_fuzz_dep_.CoverTab[171282]++

								if err != nil {
//line /usr/local/go/src/net/mail/message.go:589
			_go_fuzz_dep_.CoverTab[171291]++
									break
//line /usr/local/go/src/net/mail/message.go:590
			// _ = "end of CoverTab[171291]"
		} else {
//line /usr/local/go/src/net/mail/message.go:591
			_go_fuzz_dep_.CoverTab[171292]++
//line /usr/local/go/src/net/mail/message.go:591
			// _ = "end of CoverTab[171292]"
//line /usr/local/go/src/net/mail/message.go:591
		}
//line /usr/local/go/src/net/mail/message.go:591
		// _ = "end of CoverTab[171282]"
//line /usr/local/go/src/net/mail/message.go:591
		_go_fuzz_dep_.CoverTab[171283]++
								debug.Printf("consumePhrase: consumed %q", word)
								if isPrevEncoded && func() bool {
//line /usr/local/go/src/net/mail/message.go:593
			_go_fuzz_dep_.CoverTab[171293]++
//line /usr/local/go/src/net/mail/message.go:593
			return isEncoded
//line /usr/local/go/src/net/mail/message.go:593
			// _ = "end of CoverTab[171293]"
//line /usr/local/go/src/net/mail/message.go:593
		}() {
//line /usr/local/go/src/net/mail/message.go:593
			_go_fuzz_dep_.CoverTab[171294]++
									words[len(words)-1] += word
//line /usr/local/go/src/net/mail/message.go:594
			// _ = "end of CoverTab[171294]"
		} else {
//line /usr/local/go/src/net/mail/message.go:595
			_go_fuzz_dep_.CoverTab[171295]++
									words = append(words, word)
//line /usr/local/go/src/net/mail/message.go:596
			// _ = "end of CoverTab[171295]"
		}
//line /usr/local/go/src/net/mail/message.go:597
		// _ = "end of CoverTab[171283]"
//line /usr/local/go/src/net/mail/message.go:597
		_go_fuzz_dep_.CoverTab[171284]++
								isPrevEncoded = isEncoded
//line /usr/local/go/src/net/mail/message.go:598
		// _ = "end of CoverTab[171284]"
	}
//line /usr/local/go/src/net/mail/message.go:599
	// _ = "end of CoverTab[171277]"
//line /usr/local/go/src/net/mail/message.go:599
	_go_fuzz_dep_.CoverTab[171278]++

							if err != nil && func() bool {
//line /usr/local/go/src/net/mail/message.go:601
		_go_fuzz_dep_.CoverTab[171296]++
//line /usr/local/go/src/net/mail/message.go:601
		return len(words) == 0
//line /usr/local/go/src/net/mail/message.go:601
		// _ = "end of CoverTab[171296]"
//line /usr/local/go/src/net/mail/message.go:601
	}() {
//line /usr/local/go/src/net/mail/message.go:601
		_go_fuzz_dep_.CoverTab[171297]++
								debug.Printf("consumePhrase: hit err: %v", err)
								return "", fmt.Errorf("mail: missing word in phrase: %v", err)
//line /usr/local/go/src/net/mail/message.go:603
		// _ = "end of CoverTab[171297]"
	} else {
//line /usr/local/go/src/net/mail/message.go:604
		_go_fuzz_dep_.CoverTab[171298]++
//line /usr/local/go/src/net/mail/message.go:604
		// _ = "end of CoverTab[171298]"
//line /usr/local/go/src/net/mail/message.go:604
	}
//line /usr/local/go/src/net/mail/message.go:604
	// _ = "end of CoverTab[171278]"
//line /usr/local/go/src/net/mail/message.go:604
	_go_fuzz_dep_.CoverTab[171279]++
							phrase = strings.Join(words, " ")
							return phrase, nil
//line /usr/local/go/src/net/mail/message.go:606
	// _ = "end of CoverTab[171279]"
}

// consumeQuotedString parses the quoted string at the start of p.
func (p *addrParser) consumeQuotedString() (qs string, err error) {
//line /usr/local/go/src/net/mail/message.go:610
	_go_fuzz_dep_.CoverTab[171299]++

							i := 1
							qsb := make([]rune, 0, 10)

							escaped := false

Loop:
	for {
//line /usr/local/go/src/net/mail/message.go:618
		_go_fuzz_dep_.CoverTab[171301]++
								r, size := utf8.DecodeRuneInString(p.s[i:])

								switch {
		case size == 0:
//line /usr/local/go/src/net/mail/message.go:622
			_go_fuzz_dep_.CoverTab[171303]++
									return "", errors.New("mail: unclosed quoted-string")
//line /usr/local/go/src/net/mail/message.go:623
			// _ = "end of CoverTab[171303]"

		case size == 1 && func() bool {
//line /usr/local/go/src/net/mail/message.go:625
			_go_fuzz_dep_.CoverTab[171311]++
//line /usr/local/go/src/net/mail/message.go:625
			return r == utf8.RuneError
//line /usr/local/go/src/net/mail/message.go:625
			// _ = "end of CoverTab[171311]"
//line /usr/local/go/src/net/mail/message.go:625
		}():
//line /usr/local/go/src/net/mail/message.go:625
			_go_fuzz_dep_.CoverTab[171304]++
									return "", fmt.Errorf("mail: invalid utf-8 in quoted-string: %q", p.s)
//line /usr/local/go/src/net/mail/message.go:626
			// _ = "end of CoverTab[171304]"

		case escaped:
//line /usr/local/go/src/net/mail/message.go:628
			_go_fuzz_dep_.CoverTab[171305]++

//line /usr/local/go/src/net/mail/message.go:631
			if !isVchar(r) && func() bool {
//line /usr/local/go/src/net/mail/message.go:631
				_go_fuzz_dep_.CoverTab[171312]++
//line /usr/local/go/src/net/mail/message.go:631
				return !isWSP(r)
//line /usr/local/go/src/net/mail/message.go:631
				// _ = "end of CoverTab[171312]"
//line /usr/local/go/src/net/mail/message.go:631
			}() {
//line /usr/local/go/src/net/mail/message.go:631
				_go_fuzz_dep_.CoverTab[171313]++
										return "", fmt.Errorf("mail: bad character in quoted-string: %q", r)
//line /usr/local/go/src/net/mail/message.go:632
				// _ = "end of CoverTab[171313]"
			} else {
//line /usr/local/go/src/net/mail/message.go:633
				_go_fuzz_dep_.CoverTab[171314]++
//line /usr/local/go/src/net/mail/message.go:633
				// _ = "end of CoverTab[171314]"
//line /usr/local/go/src/net/mail/message.go:633
			}
//line /usr/local/go/src/net/mail/message.go:633
			// _ = "end of CoverTab[171305]"
//line /usr/local/go/src/net/mail/message.go:633
			_go_fuzz_dep_.CoverTab[171306]++

									qsb = append(qsb, r)
									escaped = false
//line /usr/local/go/src/net/mail/message.go:636
			// _ = "end of CoverTab[171306]"

		case isQtext(r) || func() bool {
//line /usr/local/go/src/net/mail/message.go:638
			_go_fuzz_dep_.CoverTab[171315]++
//line /usr/local/go/src/net/mail/message.go:638
			return isWSP(r)
//line /usr/local/go/src/net/mail/message.go:638
			// _ = "end of CoverTab[171315]"
//line /usr/local/go/src/net/mail/message.go:638
		}():
//line /usr/local/go/src/net/mail/message.go:638
			_go_fuzz_dep_.CoverTab[171307]++

//line /usr/local/go/src/net/mail/message.go:641
			qsb = append(qsb, r)
//line /usr/local/go/src/net/mail/message.go:641
			// _ = "end of CoverTab[171307]"

		case r == '"':
//line /usr/local/go/src/net/mail/message.go:643
			_go_fuzz_dep_.CoverTab[171308]++
									break Loop
//line /usr/local/go/src/net/mail/message.go:644
			// _ = "end of CoverTab[171308]"

		case r == '\\':
//line /usr/local/go/src/net/mail/message.go:646
			_go_fuzz_dep_.CoverTab[171309]++
									escaped = true
//line /usr/local/go/src/net/mail/message.go:647
			// _ = "end of CoverTab[171309]"

		default:
//line /usr/local/go/src/net/mail/message.go:649
			_go_fuzz_dep_.CoverTab[171310]++
									return "", fmt.Errorf("mail: bad character in quoted-string: %q", r)
//line /usr/local/go/src/net/mail/message.go:650
			// _ = "end of CoverTab[171310]"

		}
//line /usr/local/go/src/net/mail/message.go:652
		// _ = "end of CoverTab[171301]"
//line /usr/local/go/src/net/mail/message.go:652
		_go_fuzz_dep_.CoverTab[171302]++

								i += size
//line /usr/local/go/src/net/mail/message.go:654
		// _ = "end of CoverTab[171302]"
	}
//line /usr/local/go/src/net/mail/message.go:655
	// _ = "end of CoverTab[171299]"
//line /usr/local/go/src/net/mail/message.go:655
	_go_fuzz_dep_.CoverTab[171300]++
							p.s = p.s[i+1:]
							return string(qsb), nil
//line /usr/local/go/src/net/mail/message.go:657
	// _ = "end of CoverTab[171300]"
}

// consumeAtom parses an RFC 5322 atom at the start of p.
//line /usr/local/go/src/net/mail/message.go:660
// If dot is true, consumeAtom parses an RFC 5322 dot-atom instead.
//line /usr/local/go/src/net/mail/message.go:660
// If permissive is true, consumeAtom will not fail on:
//line /usr/local/go/src/net/mail/message.go:660
// - leading/trailing/double dots in the atom (see golang.org/issue/4938)
//line /usr/local/go/src/net/mail/message.go:660
// - special characters (RFC 5322 3.2.3) except '<', '>', ':' and '"' (see golang.org/issue/21018)
//line /usr/local/go/src/net/mail/message.go:665
func (p *addrParser) consumeAtom(dot bool, permissive bool) (atom string, err error) {
//line /usr/local/go/src/net/mail/message.go:665
	_go_fuzz_dep_.CoverTab[171316]++
							i := 0

Loop:
	for {
//line /usr/local/go/src/net/mail/message.go:669
		_go_fuzz_dep_.CoverTab[171320]++
								r, size := utf8.DecodeRuneInString(p.s[i:])
								switch {
		case size == 1 && func() bool {
//line /usr/local/go/src/net/mail/message.go:672
			_go_fuzz_dep_.CoverTab[171324]++
//line /usr/local/go/src/net/mail/message.go:672
			return r == utf8.RuneError
//line /usr/local/go/src/net/mail/message.go:672
			// _ = "end of CoverTab[171324]"
//line /usr/local/go/src/net/mail/message.go:672
		}():
//line /usr/local/go/src/net/mail/message.go:672
			_go_fuzz_dep_.CoverTab[171321]++
									return "", fmt.Errorf("mail: invalid utf-8 in address: %q", p.s)
//line /usr/local/go/src/net/mail/message.go:673
			// _ = "end of CoverTab[171321]"

		case size == 0 || func() bool {
//line /usr/local/go/src/net/mail/message.go:675
			_go_fuzz_dep_.CoverTab[171325]++
//line /usr/local/go/src/net/mail/message.go:675
			return !isAtext(r, dot, permissive)
//line /usr/local/go/src/net/mail/message.go:675
			// _ = "end of CoverTab[171325]"
//line /usr/local/go/src/net/mail/message.go:675
		}():
//line /usr/local/go/src/net/mail/message.go:675
			_go_fuzz_dep_.CoverTab[171322]++
									break Loop
//line /usr/local/go/src/net/mail/message.go:676
			// _ = "end of CoverTab[171322]"

		default:
//line /usr/local/go/src/net/mail/message.go:678
			_go_fuzz_dep_.CoverTab[171323]++
									i += size
//line /usr/local/go/src/net/mail/message.go:679
			// _ = "end of CoverTab[171323]"

		}
//line /usr/local/go/src/net/mail/message.go:681
		// _ = "end of CoverTab[171320]"
	}
//line /usr/local/go/src/net/mail/message.go:682
	// _ = "end of CoverTab[171316]"
//line /usr/local/go/src/net/mail/message.go:682
	_go_fuzz_dep_.CoverTab[171317]++

							if i == 0 {
//line /usr/local/go/src/net/mail/message.go:684
		_go_fuzz_dep_.CoverTab[171326]++
								return "", errors.New("mail: invalid string")
//line /usr/local/go/src/net/mail/message.go:685
		// _ = "end of CoverTab[171326]"
	} else {
//line /usr/local/go/src/net/mail/message.go:686
		_go_fuzz_dep_.CoverTab[171327]++
//line /usr/local/go/src/net/mail/message.go:686
		// _ = "end of CoverTab[171327]"
//line /usr/local/go/src/net/mail/message.go:686
	}
//line /usr/local/go/src/net/mail/message.go:686
	// _ = "end of CoverTab[171317]"
//line /usr/local/go/src/net/mail/message.go:686
	_go_fuzz_dep_.CoverTab[171318]++
							atom, p.s = p.s[:i], p.s[i:]
							if !permissive {
//line /usr/local/go/src/net/mail/message.go:688
		_go_fuzz_dep_.CoverTab[171328]++
								if strings.HasPrefix(atom, ".") {
//line /usr/local/go/src/net/mail/message.go:689
			_go_fuzz_dep_.CoverTab[171331]++
									return "", errors.New("mail: leading dot in atom")
//line /usr/local/go/src/net/mail/message.go:690
			// _ = "end of CoverTab[171331]"
		} else {
//line /usr/local/go/src/net/mail/message.go:691
			_go_fuzz_dep_.CoverTab[171332]++
//line /usr/local/go/src/net/mail/message.go:691
			// _ = "end of CoverTab[171332]"
//line /usr/local/go/src/net/mail/message.go:691
		}
//line /usr/local/go/src/net/mail/message.go:691
		// _ = "end of CoverTab[171328]"
//line /usr/local/go/src/net/mail/message.go:691
		_go_fuzz_dep_.CoverTab[171329]++
								if strings.Contains(atom, "..") {
//line /usr/local/go/src/net/mail/message.go:692
			_go_fuzz_dep_.CoverTab[171333]++
									return "", errors.New("mail: double dot in atom")
//line /usr/local/go/src/net/mail/message.go:693
			// _ = "end of CoverTab[171333]"
		} else {
//line /usr/local/go/src/net/mail/message.go:694
			_go_fuzz_dep_.CoverTab[171334]++
//line /usr/local/go/src/net/mail/message.go:694
			// _ = "end of CoverTab[171334]"
//line /usr/local/go/src/net/mail/message.go:694
		}
//line /usr/local/go/src/net/mail/message.go:694
		// _ = "end of CoverTab[171329]"
//line /usr/local/go/src/net/mail/message.go:694
		_go_fuzz_dep_.CoverTab[171330]++
								if strings.HasSuffix(atom, ".") {
//line /usr/local/go/src/net/mail/message.go:695
			_go_fuzz_dep_.CoverTab[171335]++
									return "", errors.New("mail: trailing dot in atom")
//line /usr/local/go/src/net/mail/message.go:696
			// _ = "end of CoverTab[171335]"
		} else {
//line /usr/local/go/src/net/mail/message.go:697
			_go_fuzz_dep_.CoverTab[171336]++
//line /usr/local/go/src/net/mail/message.go:697
			// _ = "end of CoverTab[171336]"
//line /usr/local/go/src/net/mail/message.go:697
		}
//line /usr/local/go/src/net/mail/message.go:697
		// _ = "end of CoverTab[171330]"
	} else {
//line /usr/local/go/src/net/mail/message.go:698
		_go_fuzz_dep_.CoverTab[171337]++
//line /usr/local/go/src/net/mail/message.go:698
		// _ = "end of CoverTab[171337]"
//line /usr/local/go/src/net/mail/message.go:698
	}
//line /usr/local/go/src/net/mail/message.go:698
	// _ = "end of CoverTab[171318]"
//line /usr/local/go/src/net/mail/message.go:698
	_go_fuzz_dep_.CoverTab[171319]++
							return atom, nil
//line /usr/local/go/src/net/mail/message.go:699
	// _ = "end of CoverTab[171319]"
}

func (p *addrParser) consumeDisplayNameComment() (string, error) {
//line /usr/local/go/src/net/mail/message.go:702
	_go_fuzz_dep_.CoverTab[171338]++
							if !p.consume('(') {
//line /usr/local/go/src/net/mail/message.go:703
		_go_fuzz_dep_.CoverTab[171343]++
								return "", errors.New("mail: comment does not start with (")
//line /usr/local/go/src/net/mail/message.go:704
		// _ = "end of CoverTab[171343]"
	} else {
//line /usr/local/go/src/net/mail/message.go:705
		_go_fuzz_dep_.CoverTab[171344]++
//line /usr/local/go/src/net/mail/message.go:705
		// _ = "end of CoverTab[171344]"
//line /usr/local/go/src/net/mail/message.go:705
	}
//line /usr/local/go/src/net/mail/message.go:705
	// _ = "end of CoverTab[171338]"
//line /usr/local/go/src/net/mail/message.go:705
	_go_fuzz_dep_.CoverTab[171339]++
							comment, ok := p.consumeComment()
							if !ok {
//line /usr/local/go/src/net/mail/message.go:707
		_go_fuzz_dep_.CoverTab[171345]++
								return "", errors.New("mail: misformatted parenthetical comment")
//line /usr/local/go/src/net/mail/message.go:708
		// _ = "end of CoverTab[171345]"
	} else {
//line /usr/local/go/src/net/mail/message.go:709
		_go_fuzz_dep_.CoverTab[171346]++
//line /usr/local/go/src/net/mail/message.go:709
		// _ = "end of CoverTab[171346]"
//line /usr/local/go/src/net/mail/message.go:709
	}
//line /usr/local/go/src/net/mail/message.go:709
	// _ = "end of CoverTab[171339]"
//line /usr/local/go/src/net/mail/message.go:709
	_go_fuzz_dep_.CoverTab[171340]++

//line /usr/local/go/src/net/mail/message.go:712
	words := strings.FieldsFunc(comment, func(r rune) bool {
//line /usr/local/go/src/net/mail/message.go:712
		_go_fuzz_dep_.CoverTab[171347]++
//line /usr/local/go/src/net/mail/message.go:712
		return r == ' ' || func() bool {
//line /usr/local/go/src/net/mail/message.go:712
			_go_fuzz_dep_.CoverTab[171348]++
//line /usr/local/go/src/net/mail/message.go:712
			return r == '\t'
//line /usr/local/go/src/net/mail/message.go:712
			// _ = "end of CoverTab[171348]"
//line /usr/local/go/src/net/mail/message.go:712
		}()
//line /usr/local/go/src/net/mail/message.go:712
		// _ = "end of CoverTab[171347]"
//line /usr/local/go/src/net/mail/message.go:712
	})
//line /usr/local/go/src/net/mail/message.go:712
	// _ = "end of CoverTab[171340]"
//line /usr/local/go/src/net/mail/message.go:712
	_go_fuzz_dep_.CoverTab[171341]++
							for idx, word := range words {
//line /usr/local/go/src/net/mail/message.go:713
		_go_fuzz_dep_.CoverTab[171349]++
								decoded, isEncoded, err := p.decodeRFC2047Word(word)
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:715
			_go_fuzz_dep_.CoverTab[171351]++
									return "", err
//line /usr/local/go/src/net/mail/message.go:716
			// _ = "end of CoverTab[171351]"
		} else {
//line /usr/local/go/src/net/mail/message.go:717
			_go_fuzz_dep_.CoverTab[171352]++
//line /usr/local/go/src/net/mail/message.go:717
			// _ = "end of CoverTab[171352]"
//line /usr/local/go/src/net/mail/message.go:717
		}
//line /usr/local/go/src/net/mail/message.go:717
		// _ = "end of CoverTab[171349]"
//line /usr/local/go/src/net/mail/message.go:717
		_go_fuzz_dep_.CoverTab[171350]++
								if isEncoded {
//line /usr/local/go/src/net/mail/message.go:718
			_go_fuzz_dep_.CoverTab[171353]++
									words[idx] = decoded
//line /usr/local/go/src/net/mail/message.go:719
			// _ = "end of CoverTab[171353]"
		} else {
//line /usr/local/go/src/net/mail/message.go:720
			_go_fuzz_dep_.CoverTab[171354]++
//line /usr/local/go/src/net/mail/message.go:720
			// _ = "end of CoverTab[171354]"
//line /usr/local/go/src/net/mail/message.go:720
		}
//line /usr/local/go/src/net/mail/message.go:720
		// _ = "end of CoverTab[171350]"
	}
//line /usr/local/go/src/net/mail/message.go:721
	// _ = "end of CoverTab[171341]"
//line /usr/local/go/src/net/mail/message.go:721
	_go_fuzz_dep_.CoverTab[171342]++

							return strings.Join(words, " "), nil
//line /usr/local/go/src/net/mail/message.go:723
	// _ = "end of CoverTab[171342]"
}

func (p *addrParser) consume(c byte) bool {
//line /usr/local/go/src/net/mail/message.go:726
	_go_fuzz_dep_.CoverTab[171355]++
							if p.empty() || func() bool {
//line /usr/local/go/src/net/mail/message.go:727
		_go_fuzz_dep_.CoverTab[171357]++
//line /usr/local/go/src/net/mail/message.go:727
		return p.peek() != c
//line /usr/local/go/src/net/mail/message.go:727
		// _ = "end of CoverTab[171357]"
//line /usr/local/go/src/net/mail/message.go:727
	}() {
//line /usr/local/go/src/net/mail/message.go:727
		_go_fuzz_dep_.CoverTab[171358]++
								return false
//line /usr/local/go/src/net/mail/message.go:728
		// _ = "end of CoverTab[171358]"
	} else {
//line /usr/local/go/src/net/mail/message.go:729
		_go_fuzz_dep_.CoverTab[171359]++
//line /usr/local/go/src/net/mail/message.go:729
		// _ = "end of CoverTab[171359]"
//line /usr/local/go/src/net/mail/message.go:729
	}
//line /usr/local/go/src/net/mail/message.go:729
	// _ = "end of CoverTab[171355]"
//line /usr/local/go/src/net/mail/message.go:729
	_go_fuzz_dep_.CoverTab[171356]++
							p.s = p.s[1:]
							return true
//line /usr/local/go/src/net/mail/message.go:731
	// _ = "end of CoverTab[171356]"
}

// skipSpace skips the leading space and tab characters.
func (p *addrParser) skipSpace() {
//line /usr/local/go/src/net/mail/message.go:735
	_go_fuzz_dep_.CoverTab[171360]++
							p.s = strings.TrimLeft(p.s, " \t")
//line /usr/local/go/src/net/mail/message.go:736
	// _ = "end of CoverTab[171360]"
}

func (p *addrParser) peek() byte {
//line /usr/local/go/src/net/mail/message.go:739
	_go_fuzz_dep_.CoverTab[171361]++
							return p.s[0]
//line /usr/local/go/src/net/mail/message.go:740
	// _ = "end of CoverTab[171361]"
}

func (p *addrParser) empty() bool {
//line /usr/local/go/src/net/mail/message.go:743
	_go_fuzz_dep_.CoverTab[171362]++
							return p.len() == 0
//line /usr/local/go/src/net/mail/message.go:744
	// _ = "end of CoverTab[171362]"
}

func (p *addrParser) len() int {
//line /usr/local/go/src/net/mail/message.go:747
	_go_fuzz_dep_.CoverTab[171363]++
							return len(p.s)
//line /usr/local/go/src/net/mail/message.go:748
	// _ = "end of CoverTab[171363]"
}

// skipCFWS skips CFWS as defined in RFC5322.
func (p *addrParser) skipCFWS() bool {
//line /usr/local/go/src/net/mail/message.go:752
	_go_fuzz_dep_.CoverTab[171364]++
							p.skipSpace()

							for {
//line /usr/local/go/src/net/mail/message.go:755
		_go_fuzz_dep_.CoverTab[171366]++
								if !p.consume('(') {
//line /usr/local/go/src/net/mail/message.go:756
			_go_fuzz_dep_.CoverTab[171369]++
									break
//line /usr/local/go/src/net/mail/message.go:757
			// _ = "end of CoverTab[171369]"
		} else {
//line /usr/local/go/src/net/mail/message.go:758
			_go_fuzz_dep_.CoverTab[171370]++
//line /usr/local/go/src/net/mail/message.go:758
			// _ = "end of CoverTab[171370]"
//line /usr/local/go/src/net/mail/message.go:758
		}
//line /usr/local/go/src/net/mail/message.go:758
		// _ = "end of CoverTab[171366]"
//line /usr/local/go/src/net/mail/message.go:758
		_go_fuzz_dep_.CoverTab[171367]++

								if _, ok := p.consumeComment(); !ok {
//line /usr/local/go/src/net/mail/message.go:760
			_go_fuzz_dep_.CoverTab[171371]++
									return false
//line /usr/local/go/src/net/mail/message.go:761
			// _ = "end of CoverTab[171371]"
		} else {
//line /usr/local/go/src/net/mail/message.go:762
			_go_fuzz_dep_.CoverTab[171372]++
//line /usr/local/go/src/net/mail/message.go:762
			// _ = "end of CoverTab[171372]"
//line /usr/local/go/src/net/mail/message.go:762
		}
//line /usr/local/go/src/net/mail/message.go:762
		// _ = "end of CoverTab[171367]"
//line /usr/local/go/src/net/mail/message.go:762
		_go_fuzz_dep_.CoverTab[171368]++

								p.skipSpace()
//line /usr/local/go/src/net/mail/message.go:764
		// _ = "end of CoverTab[171368]"
	}
//line /usr/local/go/src/net/mail/message.go:765
	// _ = "end of CoverTab[171364]"
//line /usr/local/go/src/net/mail/message.go:765
	_go_fuzz_dep_.CoverTab[171365]++

							return true
//line /usr/local/go/src/net/mail/message.go:767
	// _ = "end of CoverTab[171365]"
}

func (p *addrParser) consumeComment() (string, bool) {
//line /usr/local/go/src/net/mail/message.go:770
	_go_fuzz_dep_.CoverTab[171373]++

							depth := 1

							var comment string
							for {
//line /usr/local/go/src/net/mail/message.go:775
		_go_fuzz_dep_.CoverTab[171375]++
								if p.empty() || func() bool {
//line /usr/local/go/src/net/mail/message.go:776
			_go_fuzz_dep_.CoverTab[171379]++
//line /usr/local/go/src/net/mail/message.go:776
			return depth == 0
//line /usr/local/go/src/net/mail/message.go:776
			// _ = "end of CoverTab[171379]"
//line /usr/local/go/src/net/mail/message.go:776
		}() {
//line /usr/local/go/src/net/mail/message.go:776
			_go_fuzz_dep_.CoverTab[171380]++
									break
//line /usr/local/go/src/net/mail/message.go:777
			// _ = "end of CoverTab[171380]"
		} else {
//line /usr/local/go/src/net/mail/message.go:778
			_go_fuzz_dep_.CoverTab[171381]++
//line /usr/local/go/src/net/mail/message.go:778
			// _ = "end of CoverTab[171381]"
//line /usr/local/go/src/net/mail/message.go:778
		}
//line /usr/local/go/src/net/mail/message.go:778
		// _ = "end of CoverTab[171375]"
//line /usr/local/go/src/net/mail/message.go:778
		_go_fuzz_dep_.CoverTab[171376]++

								if p.peek() == '\\' && func() bool {
//line /usr/local/go/src/net/mail/message.go:780
			_go_fuzz_dep_.CoverTab[171382]++
//line /usr/local/go/src/net/mail/message.go:780
			return p.len() > 1
//line /usr/local/go/src/net/mail/message.go:780
			// _ = "end of CoverTab[171382]"
//line /usr/local/go/src/net/mail/message.go:780
		}() {
//line /usr/local/go/src/net/mail/message.go:780
			_go_fuzz_dep_.CoverTab[171383]++
									p.s = p.s[1:]
//line /usr/local/go/src/net/mail/message.go:781
			// _ = "end of CoverTab[171383]"
		} else {
//line /usr/local/go/src/net/mail/message.go:782
			_go_fuzz_dep_.CoverTab[171384]++
//line /usr/local/go/src/net/mail/message.go:782
			if p.peek() == '(' {
//line /usr/local/go/src/net/mail/message.go:782
				_go_fuzz_dep_.CoverTab[171385]++
										depth++
//line /usr/local/go/src/net/mail/message.go:783
				// _ = "end of CoverTab[171385]"
			} else {
//line /usr/local/go/src/net/mail/message.go:784
				_go_fuzz_dep_.CoverTab[171386]++
//line /usr/local/go/src/net/mail/message.go:784
				if p.peek() == ')' {
//line /usr/local/go/src/net/mail/message.go:784
					_go_fuzz_dep_.CoverTab[171387]++
											depth--
//line /usr/local/go/src/net/mail/message.go:785
					// _ = "end of CoverTab[171387]"
				} else {
//line /usr/local/go/src/net/mail/message.go:786
					_go_fuzz_dep_.CoverTab[171388]++
//line /usr/local/go/src/net/mail/message.go:786
					// _ = "end of CoverTab[171388]"
//line /usr/local/go/src/net/mail/message.go:786
				}
//line /usr/local/go/src/net/mail/message.go:786
				// _ = "end of CoverTab[171386]"
//line /usr/local/go/src/net/mail/message.go:786
			}
//line /usr/local/go/src/net/mail/message.go:786
			// _ = "end of CoverTab[171384]"
//line /usr/local/go/src/net/mail/message.go:786
		}
//line /usr/local/go/src/net/mail/message.go:786
		// _ = "end of CoverTab[171376]"
//line /usr/local/go/src/net/mail/message.go:786
		_go_fuzz_dep_.CoverTab[171377]++
								if depth > 0 {
//line /usr/local/go/src/net/mail/message.go:787
			_go_fuzz_dep_.CoverTab[171389]++
									comment += p.s[:1]
//line /usr/local/go/src/net/mail/message.go:788
			// _ = "end of CoverTab[171389]"
		} else {
//line /usr/local/go/src/net/mail/message.go:789
			_go_fuzz_dep_.CoverTab[171390]++
//line /usr/local/go/src/net/mail/message.go:789
			// _ = "end of CoverTab[171390]"
//line /usr/local/go/src/net/mail/message.go:789
		}
//line /usr/local/go/src/net/mail/message.go:789
		// _ = "end of CoverTab[171377]"
//line /usr/local/go/src/net/mail/message.go:789
		_go_fuzz_dep_.CoverTab[171378]++
								p.s = p.s[1:]
//line /usr/local/go/src/net/mail/message.go:790
		// _ = "end of CoverTab[171378]"
	}
//line /usr/local/go/src/net/mail/message.go:791
	// _ = "end of CoverTab[171373]"
//line /usr/local/go/src/net/mail/message.go:791
	_go_fuzz_dep_.CoverTab[171374]++

							return comment, depth == 0
//line /usr/local/go/src/net/mail/message.go:793
	// _ = "end of CoverTab[171374]"
}

func (p *addrParser) decodeRFC2047Word(s string) (word string, isEncoded bool, err error) {
//line /usr/local/go/src/net/mail/message.go:796
	_go_fuzz_dep_.CoverTab[171391]++
							dec := p.dec
							if dec == nil {
//line /usr/local/go/src/net/mail/message.go:798
		_go_fuzz_dep_.CoverTab[171396]++
								dec = &rfc2047Decoder
//line /usr/local/go/src/net/mail/message.go:799
		// _ = "end of CoverTab[171396]"
	} else {
//line /usr/local/go/src/net/mail/message.go:800
		_go_fuzz_dep_.CoverTab[171397]++
//line /usr/local/go/src/net/mail/message.go:800
		// _ = "end of CoverTab[171397]"
//line /usr/local/go/src/net/mail/message.go:800
	}
//line /usr/local/go/src/net/mail/message.go:800
	// _ = "end of CoverTab[171391]"
//line /usr/local/go/src/net/mail/message.go:800
	_go_fuzz_dep_.CoverTab[171392]++

//line /usr/local/go/src/net/mail/message.go:808
	adec := *dec
	charsetReaderError := false
	adec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
//line /usr/local/go/src/net/mail/message.go:810
		_go_fuzz_dep_.CoverTab[171398]++
								if dec.CharsetReader == nil {
//line /usr/local/go/src/net/mail/message.go:811
			_go_fuzz_dep_.CoverTab[171401]++
									charsetReaderError = true
									return nil, charsetError(charset)
//line /usr/local/go/src/net/mail/message.go:813
			// _ = "end of CoverTab[171401]"
		} else {
//line /usr/local/go/src/net/mail/message.go:814
			_go_fuzz_dep_.CoverTab[171402]++
//line /usr/local/go/src/net/mail/message.go:814
			// _ = "end of CoverTab[171402]"
//line /usr/local/go/src/net/mail/message.go:814
		}
//line /usr/local/go/src/net/mail/message.go:814
		// _ = "end of CoverTab[171398]"
//line /usr/local/go/src/net/mail/message.go:814
		_go_fuzz_dep_.CoverTab[171399]++
								r, err := dec.CharsetReader(charset, input)
								if err != nil {
//line /usr/local/go/src/net/mail/message.go:816
			_go_fuzz_dep_.CoverTab[171403]++
									charsetReaderError = true
//line /usr/local/go/src/net/mail/message.go:817
			// _ = "end of CoverTab[171403]"
		} else {
//line /usr/local/go/src/net/mail/message.go:818
			_go_fuzz_dep_.CoverTab[171404]++
//line /usr/local/go/src/net/mail/message.go:818
			// _ = "end of CoverTab[171404]"
//line /usr/local/go/src/net/mail/message.go:818
		}
//line /usr/local/go/src/net/mail/message.go:818
		// _ = "end of CoverTab[171399]"
//line /usr/local/go/src/net/mail/message.go:818
		_go_fuzz_dep_.CoverTab[171400]++
								return r, err
//line /usr/local/go/src/net/mail/message.go:819
		// _ = "end of CoverTab[171400]"
	}
//line /usr/local/go/src/net/mail/message.go:820
	// _ = "end of CoverTab[171392]"
//line /usr/local/go/src/net/mail/message.go:820
	_go_fuzz_dep_.CoverTab[171393]++
							word, err = adec.Decode(s)
							if err == nil {
//line /usr/local/go/src/net/mail/message.go:822
		_go_fuzz_dep_.CoverTab[171405]++
								return word, true, nil
//line /usr/local/go/src/net/mail/message.go:823
		// _ = "end of CoverTab[171405]"
	} else {
//line /usr/local/go/src/net/mail/message.go:824
		_go_fuzz_dep_.CoverTab[171406]++
//line /usr/local/go/src/net/mail/message.go:824
		// _ = "end of CoverTab[171406]"
//line /usr/local/go/src/net/mail/message.go:824
	}
//line /usr/local/go/src/net/mail/message.go:824
	// _ = "end of CoverTab[171393]"
//line /usr/local/go/src/net/mail/message.go:824
	_go_fuzz_dep_.CoverTab[171394]++

//line /usr/local/go/src/net/mail/message.go:831
	if charsetReaderError {
//line /usr/local/go/src/net/mail/message.go:831
		_go_fuzz_dep_.CoverTab[171407]++
								return s, true, err
//line /usr/local/go/src/net/mail/message.go:832
		// _ = "end of CoverTab[171407]"
	} else {
//line /usr/local/go/src/net/mail/message.go:833
		_go_fuzz_dep_.CoverTab[171408]++
//line /usr/local/go/src/net/mail/message.go:833
		// _ = "end of CoverTab[171408]"
//line /usr/local/go/src/net/mail/message.go:833
	}
//line /usr/local/go/src/net/mail/message.go:833
	// _ = "end of CoverTab[171394]"
//line /usr/local/go/src/net/mail/message.go:833
	_go_fuzz_dep_.CoverTab[171395]++

//line /usr/local/go/src/net/mail/message.go:836
	return s, false, nil
//line /usr/local/go/src/net/mail/message.go:836
	// _ = "end of CoverTab[171395]"
}

var rfc2047Decoder = mime.WordDecoder{
	CharsetReader: func(charset string, input io.Reader) (io.Reader, error) {
//line /usr/local/go/src/net/mail/message.go:840
		_go_fuzz_dep_.CoverTab[171409]++
								return nil, charsetError(charset)
//line /usr/local/go/src/net/mail/message.go:841
		// _ = "end of CoverTab[171409]"
	},
}

type charsetError string

func (e charsetError) Error() string {
//line /usr/local/go/src/net/mail/message.go:847
	_go_fuzz_dep_.CoverTab[171410]++
							return fmt.Sprintf("charset not supported: %q", string(e))
//line /usr/local/go/src/net/mail/message.go:848
	// _ = "end of CoverTab[171410]"
}

// isAtext reports whether r is an RFC 5322 atext character.
//line /usr/local/go/src/net/mail/message.go:851
// If dot is true, period is included.
//line /usr/local/go/src/net/mail/message.go:851
// If permissive is true, RFC 5322 3.2.3 specials is included,
//line /usr/local/go/src/net/mail/message.go:851
// except '<', '>', ':' and '"'.
//line /usr/local/go/src/net/mail/message.go:855
func isAtext(r rune, dot, permissive bool) bool {
//line /usr/local/go/src/net/mail/message.go:855
	_go_fuzz_dep_.CoverTab[171411]++
							switch r {
	case '.':
//line /usr/local/go/src/net/mail/message.go:857
		_go_fuzz_dep_.CoverTab[171413]++
								return dot
//line /usr/local/go/src/net/mail/message.go:858
		// _ = "end of CoverTab[171413]"

//line /usr/local/go/src/net/mail/message.go:861
	case '(', ')', '[', ']', ';', '@', '\\', ',':
//line /usr/local/go/src/net/mail/message.go:861
		_go_fuzz_dep_.CoverTab[171414]++
								return permissive
//line /usr/local/go/src/net/mail/message.go:862
		// _ = "end of CoverTab[171414]"

	case '<', '>', '"', ':':
//line /usr/local/go/src/net/mail/message.go:864
		_go_fuzz_dep_.CoverTab[171415]++
								return false
//line /usr/local/go/src/net/mail/message.go:865
		// _ = "end of CoverTab[171415]"
//line /usr/local/go/src/net/mail/message.go:865
	default:
//line /usr/local/go/src/net/mail/message.go:865
		_go_fuzz_dep_.CoverTab[171416]++
//line /usr/local/go/src/net/mail/message.go:865
		// _ = "end of CoverTab[171416]"
	}
//line /usr/local/go/src/net/mail/message.go:866
	// _ = "end of CoverTab[171411]"
//line /usr/local/go/src/net/mail/message.go:866
	_go_fuzz_dep_.CoverTab[171412]++
							return isVchar(r)
//line /usr/local/go/src/net/mail/message.go:867
	// _ = "end of CoverTab[171412]"
}

// isQtext reports whether r is an RFC 5322 qtext character.
func isQtext(r rune) bool {
//line /usr/local/go/src/net/mail/message.go:871
	_go_fuzz_dep_.CoverTab[171417]++

							if r == '\\' || func() bool {
//line /usr/local/go/src/net/mail/message.go:873
		_go_fuzz_dep_.CoverTab[171419]++
//line /usr/local/go/src/net/mail/message.go:873
		return r == '"'
//line /usr/local/go/src/net/mail/message.go:873
		// _ = "end of CoverTab[171419]"
//line /usr/local/go/src/net/mail/message.go:873
	}() {
//line /usr/local/go/src/net/mail/message.go:873
		_go_fuzz_dep_.CoverTab[171420]++
								return false
//line /usr/local/go/src/net/mail/message.go:874
		// _ = "end of CoverTab[171420]"
	} else {
//line /usr/local/go/src/net/mail/message.go:875
		_go_fuzz_dep_.CoverTab[171421]++
//line /usr/local/go/src/net/mail/message.go:875
		// _ = "end of CoverTab[171421]"
//line /usr/local/go/src/net/mail/message.go:875
	}
//line /usr/local/go/src/net/mail/message.go:875
	// _ = "end of CoverTab[171417]"
//line /usr/local/go/src/net/mail/message.go:875
	_go_fuzz_dep_.CoverTab[171418]++
							return isVchar(r)
//line /usr/local/go/src/net/mail/message.go:876
	// _ = "end of CoverTab[171418]"
}

// quoteString renders a string as an RFC 5322 quoted-string.
func quoteString(s string) string {
//line /usr/local/go/src/net/mail/message.go:880
	_go_fuzz_dep_.CoverTab[171422]++
							var b strings.Builder
							b.WriteByte('"')
							for _, r := range s {
//line /usr/local/go/src/net/mail/message.go:883
		_go_fuzz_dep_.CoverTab[171424]++
								if isQtext(r) || func() bool {
//line /usr/local/go/src/net/mail/message.go:884
			_go_fuzz_dep_.CoverTab[171425]++
//line /usr/local/go/src/net/mail/message.go:884
			return isWSP(r)
//line /usr/local/go/src/net/mail/message.go:884
			// _ = "end of CoverTab[171425]"
//line /usr/local/go/src/net/mail/message.go:884
		}() {
//line /usr/local/go/src/net/mail/message.go:884
			_go_fuzz_dep_.CoverTab[171426]++
									b.WriteRune(r)
//line /usr/local/go/src/net/mail/message.go:885
			// _ = "end of CoverTab[171426]"
		} else {
//line /usr/local/go/src/net/mail/message.go:886
			_go_fuzz_dep_.CoverTab[171427]++
//line /usr/local/go/src/net/mail/message.go:886
			if isVchar(r) {
//line /usr/local/go/src/net/mail/message.go:886
				_go_fuzz_dep_.CoverTab[171428]++
										b.WriteByte('\\')
										b.WriteRune(r)
//line /usr/local/go/src/net/mail/message.go:888
				// _ = "end of CoverTab[171428]"
			} else {
//line /usr/local/go/src/net/mail/message.go:889
				_go_fuzz_dep_.CoverTab[171429]++
//line /usr/local/go/src/net/mail/message.go:889
				// _ = "end of CoverTab[171429]"
//line /usr/local/go/src/net/mail/message.go:889
			}
//line /usr/local/go/src/net/mail/message.go:889
			// _ = "end of CoverTab[171427]"
//line /usr/local/go/src/net/mail/message.go:889
		}
//line /usr/local/go/src/net/mail/message.go:889
		// _ = "end of CoverTab[171424]"
	}
//line /usr/local/go/src/net/mail/message.go:890
	// _ = "end of CoverTab[171422]"
//line /usr/local/go/src/net/mail/message.go:890
	_go_fuzz_dep_.CoverTab[171423]++
							b.WriteByte('"')
							return b.String()
//line /usr/local/go/src/net/mail/message.go:892
	// _ = "end of CoverTab[171423]"
}

// isVchar reports whether r is an RFC 5322 VCHAR character.
func isVchar(r rune) bool {
//line /usr/local/go/src/net/mail/message.go:896
	_go_fuzz_dep_.CoverTab[171430]++

							return '!' <= r && func() bool {
//line /usr/local/go/src/net/mail/message.go:898
		_go_fuzz_dep_.CoverTab[171431]++
//line /usr/local/go/src/net/mail/message.go:898
		return r <= '~'
//line /usr/local/go/src/net/mail/message.go:898
		// _ = "end of CoverTab[171431]"
//line /usr/local/go/src/net/mail/message.go:898
	}() || func() bool {
//line /usr/local/go/src/net/mail/message.go:898
		_go_fuzz_dep_.CoverTab[171432]++
//line /usr/local/go/src/net/mail/message.go:898
		return isMultibyte(r)
//line /usr/local/go/src/net/mail/message.go:898
		// _ = "end of CoverTab[171432]"
//line /usr/local/go/src/net/mail/message.go:898
	}()
//line /usr/local/go/src/net/mail/message.go:898
	// _ = "end of CoverTab[171430]"
}

// isMultibyte reports whether r is a multi-byte UTF-8 character
//line /usr/local/go/src/net/mail/message.go:901
// as supported by RFC 6532.
//line /usr/local/go/src/net/mail/message.go:903
func isMultibyte(r rune) bool {
//line /usr/local/go/src/net/mail/message.go:903
	_go_fuzz_dep_.CoverTab[171433]++
							return r >= utf8.RuneSelf
//line /usr/local/go/src/net/mail/message.go:904
	// _ = "end of CoverTab[171433]"
}

// isWSP reports whether r is a WSP (white space).
//line /usr/local/go/src/net/mail/message.go:907
// WSP is a space or horizontal tab (RFC 5234 Appendix B).
//line /usr/local/go/src/net/mail/message.go:909
func isWSP(r rune) bool {
//line /usr/local/go/src/net/mail/message.go:909
	_go_fuzz_dep_.CoverTab[171434]++
							return r == ' ' || func() bool {
//line /usr/local/go/src/net/mail/message.go:910
		_go_fuzz_dep_.CoverTab[171435]++
//line /usr/local/go/src/net/mail/message.go:910
		return r == '\t'
//line /usr/local/go/src/net/mail/message.go:910
		// _ = "end of CoverTab[171435]"
//line /usr/local/go/src/net/mail/message.go:910
	}()
//line /usr/local/go/src/net/mail/message.go:910
	// _ = "end of CoverTab[171434]"
}

//line /usr/local/go/src/net/mail/message.go:911
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/mail/message.go:911
var _ = _go_fuzz_dep_.CoverTab
