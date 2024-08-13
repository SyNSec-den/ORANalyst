// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/status.go:5
package http

//line /usr/local/go/src/net/http/status.go:5
import (
//line /usr/local/go/src/net/http/status.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/status.go:5
)
//line /usr/local/go/src/net/http/status.go:5
import (
//line /usr/local/go/src/net/http/status.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/status.go:5
)

// HTTP status codes as registered with IANA.
//line /usr/local/go/src/net/http/status.go:7
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
//line /usr/local/go/src/net/http/status.go:9
const (
	StatusContinue			= 100	// RFC 9110, 15.2.1
	StatusSwitchingProtocols	= 101	// RFC 9110, 15.2.2
	StatusProcessing		= 102	// RFC 2518, 10.1
	StatusEarlyHints		= 103	// RFC 8297

	StatusOK			= 200	// RFC 9110, 15.3.1
	StatusCreated			= 201	// RFC 9110, 15.3.2
	StatusAccepted			= 202	// RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo	= 203	// RFC 9110, 15.3.4
	StatusNoContent			= 204	// RFC 9110, 15.3.5
	StatusResetContent		= 205	// RFC 9110, 15.3.6
	StatusPartialContent		= 206	// RFC 9110, 15.3.7
	StatusMultiStatus		= 207	// RFC 4918, 11.1
	StatusAlreadyReported		= 208	// RFC 5842, 7.1
	StatusIMUsed			= 226	// RFC 3229, 10.4.1

	StatusMultipleChoices	= 300	// RFC 9110, 15.4.1
	StatusMovedPermanently	= 301	// RFC 9110, 15.4.2
	StatusFound		= 302	// RFC 9110, 15.4.3
	StatusSeeOther		= 303	// RFC 9110, 15.4.4
	StatusNotModified	= 304	// RFC 9110, 15.4.5
	StatusUseProxy		= 305	// RFC 9110, 15.4.6
	_			= 306	// RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect	= 307	// RFC 9110, 15.4.8
	StatusPermanentRedirect	= 308	// RFC 9110, 15.4.9

	StatusBadRequest			= 400	// RFC 9110, 15.5.1
	StatusUnauthorized			= 401	// RFC 9110, 15.5.2
	StatusPaymentRequired			= 402	// RFC 9110, 15.5.3
	StatusForbidden				= 403	// RFC 9110, 15.5.4
	StatusNotFound				= 404	// RFC 9110, 15.5.5
	StatusMethodNotAllowed			= 405	// RFC 9110, 15.5.6
	StatusNotAcceptable			= 406	// RFC 9110, 15.5.7
	StatusProxyAuthRequired			= 407	// RFC 9110, 15.5.8
	StatusRequestTimeout			= 408	// RFC 9110, 15.5.9
	StatusConflict				= 409	// RFC 9110, 15.5.10
	StatusGone				= 410	// RFC 9110, 15.5.11
	StatusLengthRequired			= 411	// RFC 9110, 15.5.12
	StatusPreconditionFailed		= 412	// RFC 9110, 15.5.13
	StatusRequestEntityTooLarge		= 413	// RFC 9110, 15.5.14
	StatusRequestURITooLong			= 414	// RFC 9110, 15.5.15
	StatusUnsupportedMediaType		= 415	// RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable	= 416	// RFC 9110, 15.5.17
	StatusExpectationFailed			= 417	// RFC 9110, 15.5.18
	StatusTeapot				= 418	// RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest		= 421	// RFC 9110, 15.5.20
	StatusUnprocessableEntity		= 422	// RFC 9110, 15.5.21
	StatusLocked				= 423	// RFC 4918, 11.3
	StatusFailedDependency			= 424	// RFC 4918, 11.4
	StatusTooEarly				= 425	// RFC 8470, 5.2.
	StatusUpgradeRequired			= 426	// RFC 9110, 15.5.22
	StatusPreconditionRequired		= 428	// RFC 6585, 3
	StatusTooManyRequests			= 429	// RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge	= 431	// RFC 6585, 5
	StatusUnavailableForLegalReasons	= 451	// RFC 7725, 3

	StatusInternalServerError		= 500	// RFC 9110, 15.6.1
	StatusNotImplemented			= 501	// RFC 9110, 15.6.2
	StatusBadGateway			= 502	// RFC 9110, 15.6.3
	StatusServiceUnavailable		= 503	// RFC 9110, 15.6.4
	StatusGatewayTimeout			= 504	// RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported		= 505	// RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates		= 506	// RFC 2295, 8.1
	StatusInsufficientStorage		= 507	// RFC 4918, 11.5
	StatusLoopDetected			= 508	// RFC 5842, 7.2
	StatusNotExtended			= 510	// RFC 2774, 7
	StatusNetworkAuthenticationRequired	= 511	// RFC 6585, 6
)

// StatusText returns a text for the HTTP status code. It returns the empty
//line /usr/local/go/src/net/http/status.go:79
// string if the code is unknown.
//line /usr/local/go/src/net/http/status.go:81
func StatusText(code int) string {
//line /usr/local/go/src/net/http/status.go:81
	_go_fuzz_dep_.CoverTab[43488]++
						switch code {
	case StatusContinue:
//line /usr/local/go/src/net/http/status.go:83
		_go_fuzz_dep_.CoverTab[43489]++
							return "Continue"
//line /usr/local/go/src/net/http/status.go:84
		// _ = "end of CoverTab[43489]"
	case StatusSwitchingProtocols:
//line /usr/local/go/src/net/http/status.go:85
		_go_fuzz_dep_.CoverTab[43490]++
							return "Switching Protocols"
//line /usr/local/go/src/net/http/status.go:86
		// _ = "end of CoverTab[43490]"
	case StatusProcessing:
//line /usr/local/go/src/net/http/status.go:87
		_go_fuzz_dep_.CoverTab[43491]++
							return "Processing"
//line /usr/local/go/src/net/http/status.go:88
		// _ = "end of CoverTab[43491]"
	case StatusEarlyHints:
//line /usr/local/go/src/net/http/status.go:89
		_go_fuzz_dep_.CoverTab[43492]++
							return "Early Hints"
//line /usr/local/go/src/net/http/status.go:90
		// _ = "end of CoverTab[43492]"
	case StatusOK:
//line /usr/local/go/src/net/http/status.go:91
		_go_fuzz_dep_.CoverTab[43493]++
							return "OK"
//line /usr/local/go/src/net/http/status.go:92
		// _ = "end of CoverTab[43493]"
	case StatusCreated:
//line /usr/local/go/src/net/http/status.go:93
		_go_fuzz_dep_.CoverTab[43494]++
							return "Created"
//line /usr/local/go/src/net/http/status.go:94
		// _ = "end of CoverTab[43494]"
	case StatusAccepted:
//line /usr/local/go/src/net/http/status.go:95
		_go_fuzz_dep_.CoverTab[43495]++
							return "Accepted"
//line /usr/local/go/src/net/http/status.go:96
		// _ = "end of CoverTab[43495]"
	case StatusNonAuthoritativeInfo:
//line /usr/local/go/src/net/http/status.go:97
		_go_fuzz_dep_.CoverTab[43496]++
							return "Non-Authoritative Information"
//line /usr/local/go/src/net/http/status.go:98
		// _ = "end of CoverTab[43496]"
	case StatusNoContent:
//line /usr/local/go/src/net/http/status.go:99
			_go_fuzz_dep_.CoverTab[43497]++
								return "No Content"
//line /usr/local/go/src/net/http/status.go:100
		// _ = "end of CoverTab[43497]"
	case StatusResetContent:
//line /usr/local/go/src/net/http/status.go:101
		_go_fuzz_dep_.CoverTab[43498]++
								return "Reset Content"
//line /usr/local/go/src/net/http/status.go:102
		// _ = "end of CoverTab[43498]"
	case StatusPartialContent:
//line /usr/local/go/src/net/http/status.go:103
		_go_fuzz_dep_.CoverTab[43499]++
								return "Partial Content"
//line /usr/local/go/src/net/http/status.go:104
		// _ = "end of CoverTab[43499]"
	case StatusMultiStatus:
//line /usr/local/go/src/net/http/status.go:105
		_go_fuzz_dep_.CoverTab[43500]++
								return "Multi-Status"
//line /usr/local/go/src/net/http/status.go:106
		// _ = "end of CoverTab[43500]"
	case StatusAlreadyReported:
//line /usr/local/go/src/net/http/status.go:107
		_go_fuzz_dep_.CoverTab[43501]++
								return "Already Reported"
//line /usr/local/go/src/net/http/status.go:108
		// _ = "end of CoverTab[43501]"
	case StatusIMUsed:
//line /usr/local/go/src/net/http/status.go:109
		_go_fuzz_dep_.CoverTab[43502]++
								return "IM Used"
//line /usr/local/go/src/net/http/status.go:110
		// _ = "end of CoverTab[43502]"
	case StatusMultipleChoices:
//line /usr/local/go/src/net/http/status.go:111
		_go_fuzz_dep_.CoverTab[43503]++
								return "Multiple Choices"
//line /usr/local/go/src/net/http/status.go:112
		// _ = "end of CoverTab[43503]"
	case StatusMovedPermanently:
//line /usr/local/go/src/net/http/status.go:113
		_go_fuzz_dep_.CoverTab[43504]++
								return "Moved Permanently"
//line /usr/local/go/src/net/http/status.go:114
		// _ = "end of CoverTab[43504]"
	case StatusFound:
//line /usr/local/go/src/net/http/status.go:115
		_go_fuzz_dep_.CoverTab[43505]++
								return "Found"
//line /usr/local/go/src/net/http/status.go:116
		// _ = "end of CoverTab[43505]"
	case StatusSeeOther:
//line /usr/local/go/src/net/http/status.go:117
		_go_fuzz_dep_.CoverTab[43506]++
								return "See Other"
//line /usr/local/go/src/net/http/status.go:118
		// _ = "end of CoverTab[43506]"
	case StatusNotModified:
//line /usr/local/go/src/net/http/status.go:119
		_go_fuzz_dep_.CoverTab[43507]++
								return "Not Modified"
//line /usr/local/go/src/net/http/status.go:120
		// _ = "end of CoverTab[43507]"
	case StatusUseProxy:
//line /usr/local/go/src/net/http/status.go:121
		_go_fuzz_dep_.CoverTab[43508]++
								return "Use Proxy"
//line /usr/local/go/src/net/http/status.go:122
		// _ = "end of CoverTab[43508]"
	case StatusTemporaryRedirect:
//line /usr/local/go/src/net/http/status.go:123
		_go_fuzz_dep_.CoverTab[43509]++
								return "Temporary Redirect"
//line /usr/local/go/src/net/http/status.go:124
		// _ = "end of CoverTab[43509]"
	case StatusPermanentRedirect:
//line /usr/local/go/src/net/http/status.go:125
		_go_fuzz_dep_.CoverTab[43510]++
								return "Permanent Redirect"
//line /usr/local/go/src/net/http/status.go:126
		// _ = "end of CoverTab[43510]"
	case StatusBadRequest:
//line /usr/local/go/src/net/http/status.go:127
		_go_fuzz_dep_.CoverTab[43511]++
								return "Bad Request"
//line /usr/local/go/src/net/http/status.go:128
		// _ = "end of CoverTab[43511]"
	case StatusUnauthorized:
//line /usr/local/go/src/net/http/status.go:129
		_go_fuzz_dep_.CoverTab[43512]++
								return "Unauthorized"
//line /usr/local/go/src/net/http/status.go:130
		// _ = "end of CoverTab[43512]"
	case StatusPaymentRequired:
//line /usr/local/go/src/net/http/status.go:131
		_go_fuzz_dep_.CoverTab[43513]++
								return "Payment Required"
//line /usr/local/go/src/net/http/status.go:132
		// _ = "end of CoverTab[43513]"
	case StatusForbidden:
//line /usr/local/go/src/net/http/status.go:133
		_go_fuzz_dep_.CoverTab[43514]++
								return "Forbidden"
//line /usr/local/go/src/net/http/status.go:134
		// _ = "end of CoverTab[43514]"
	case StatusNotFound:
//line /usr/local/go/src/net/http/status.go:135
		_go_fuzz_dep_.CoverTab[43515]++
								return "Not Found"
//line /usr/local/go/src/net/http/status.go:136
		// _ = "end of CoverTab[43515]"
	case StatusMethodNotAllowed:
//line /usr/local/go/src/net/http/status.go:137
		_go_fuzz_dep_.CoverTab[43516]++
								return "Method Not Allowed"
//line /usr/local/go/src/net/http/status.go:138
		// _ = "end of CoverTab[43516]"
	case StatusNotAcceptable:
//line /usr/local/go/src/net/http/status.go:139
		_go_fuzz_dep_.CoverTab[43517]++
								return "Not Acceptable"
//line /usr/local/go/src/net/http/status.go:140
		// _ = "end of CoverTab[43517]"
	case StatusProxyAuthRequired:
//line /usr/local/go/src/net/http/status.go:141
		_go_fuzz_dep_.CoverTab[43518]++
								return "Proxy Authentication Required"
//line /usr/local/go/src/net/http/status.go:142
		// _ = "end of CoverTab[43518]"
	case StatusRequestTimeout:
//line /usr/local/go/src/net/http/status.go:143
		_go_fuzz_dep_.CoverTab[43519]++
								return "Request Timeout"
//line /usr/local/go/src/net/http/status.go:144
		// _ = "end of CoverTab[43519]"
	case StatusConflict:
//line /usr/local/go/src/net/http/status.go:145
		_go_fuzz_dep_.CoverTab[43520]++
								return "Conflict"
//line /usr/local/go/src/net/http/status.go:146
		// _ = "end of CoverTab[43520]"
	case StatusGone:
//line /usr/local/go/src/net/http/status.go:147
		_go_fuzz_dep_.CoverTab[43521]++
								return "Gone"
//line /usr/local/go/src/net/http/status.go:148
		// _ = "end of CoverTab[43521]"
	case StatusLengthRequired:
//line /usr/local/go/src/net/http/status.go:149
		_go_fuzz_dep_.CoverTab[43522]++
								return "Length Required"
//line /usr/local/go/src/net/http/status.go:150
		// _ = "end of CoverTab[43522]"
	case StatusPreconditionFailed:
//line /usr/local/go/src/net/http/status.go:151
		_go_fuzz_dep_.CoverTab[43523]++
								return "Precondition Failed"
//line /usr/local/go/src/net/http/status.go:152
		// _ = "end of CoverTab[43523]"
	case StatusRequestEntityTooLarge:
//line /usr/local/go/src/net/http/status.go:153
		_go_fuzz_dep_.CoverTab[43524]++
								return "Request Entity Too Large"
//line /usr/local/go/src/net/http/status.go:154
		// _ = "end of CoverTab[43524]"
	case StatusRequestURITooLong:
//line /usr/local/go/src/net/http/status.go:155
		_go_fuzz_dep_.CoverTab[43525]++
								return "Request URI Too Long"
//line /usr/local/go/src/net/http/status.go:156
		// _ = "end of CoverTab[43525]"
	case StatusUnsupportedMediaType:
//line /usr/local/go/src/net/http/status.go:157
		_go_fuzz_dep_.CoverTab[43526]++
								return "Unsupported Media Type"
//line /usr/local/go/src/net/http/status.go:158
		// _ = "end of CoverTab[43526]"
	case StatusRequestedRangeNotSatisfiable:
//line /usr/local/go/src/net/http/status.go:159
		_go_fuzz_dep_.CoverTab[43527]++
								return "Requested Range Not Satisfiable"
//line /usr/local/go/src/net/http/status.go:160
		// _ = "end of CoverTab[43527]"
	case StatusExpectationFailed:
//line /usr/local/go/src/net/http/status.go:161
		_go_fuzz_dep_.CoverTab[43528]++
								return "Expectation Failed"
//line /usr/local/go/src/net/http/status.go:162
		// _ = "end of CoverTab[43528]"
	case StatusTeapot:
//line /usr/local/go/src/net/http/status.go:163
		_go_fuzz_dep_.CoverTab[43529]++
								return "I'm a teapot"
//line /usr/local/go/src/net/http/status.go:164
		// _ = "end of CoverTab[43529]"
	case StatusMisdirectedRequest:
//line /usr/local/go/src/net/http/status.go:165
		_go_fuzz_dep_.CoverTab[43530]++
								return "Misdirected Request"
//line /usr/local/go/src/net/http/status.go:166
		// _ = "end of CoverTab[43530]"
	case StatusUnprocessableEntity:
//line /usr/local/go/src/net/http/status.go:167
		_go_fuzz_dep_.CoverTab[43531]++
								return "Unprocessable Entity"
//line /usr/local/go/src/net/http/status.go:168
		// _ = "end of CoverTab[43531]"
	case StatusLocked:
//line /usr/local/go/src/net/http/status.go:169
		_go_fuzz_dep_.CoverTab[43532]++
								return "Locked"
//line /usr/local/go/src/net/http/status.go:170
		// _ = "end of CoverTab[43532]"
	case StatusFailedDependency:
//line /usr/local/go/src/net/http/status.go:171
		_go_fuzz_dep_.CoverTab[43533]++
								return "Failed Dependency"
//line /usr/local/go/src/net/http/status.go:172
		// _ = "end of CoverTab[43533]"
	case StatusTooEarly:
//line /usr/local/go/src/net/http/status.go:173
		_go_fuzz_dep_.CoverTab[43534]++
								return "Too Early"
//line /usr/local/go/src/net/http/status.go:174
		// _ = "end of CoverTab[43534]"
	case StatusUpgradeRequired:
//line /usr/local/go/src/net/http/status.go:175
		_go_fuzz_dep_.CoverTab[43535]++
								return "Upgrade Required"
//line /usr/local/go/src/net/http/status.go:176
		// _ = "end of CoverTab[43535]"
	case StatusPreconditionRequired:
//line /usr/local/go/src/net/http/status.go:177
		_go_fuzz_dep_.CoverTab[43536]++
								return "Precondition Required"
//line /usr/local/go/src/net/http/status.go:178
		// _ = "end of CoverTab[43536]"
	case StatusTooManyRequests:
//line /usr/local/go/src/net/http/status.go:179
		_go_fuzz_dep_.CoverTab[43537]++
								return "Too Many Requests"
//line /usr/local/go/src/net/http/status.go:180
		// _ = "end of CoverTab[43537]"
	case StatusRequestHeaderFieldsTooLarge:
//line /usr/local/go/src/net/http/status.go:181
		_go_fuzz_dep_.CoverTab[43538]++
								return "Request Header Fields Too Large"
//line /usr/local/go/src/net/http/status.go:182
		// _ = "end of CoverTab[43538]"
	case StatusUnavailableForLegalReasons:
//line /usr/local/go/src/net/http/status.go:183
		_go_fuzz_dep_.CoverTab[43539]++
								return "Unavailable For Legal Reasons"
//line /usr/local/go/src/net/http/status.go:184
		// _ = "end of CoverTab[43539]"
	case StatusInternalServerError:
//line /usr/local/go/src/net/http/status.go:185
		_go_fuzz_dep_.CoverTab[43540]++
								return "Internal Server Error"
//line /usr/local/go/src/net/http/status.go:186
		// _ = "end of CoverTab[43540]"
	case StatusNotImplemented:
//line /usr/local/go/src/net/http/status.go:187
		_go_fuzz_dep_.CoverTab[43541]++
								return "Not Implemented"
//line /usr/local/go/src/net/http/status.go:188
		// _ = "end of CoverTab[43541]"
	case StatusBadGateway:
//line /usr/local/go/src/net/http/status.go:189
		_go_fuzz_dep_.CoverTab[43542]++
								return "Bad Gateway"
//line /usr/local/go/src/net/http/status.go:190
		// _ = "end of CoverTab[43542]"
	case StatusServiceUnavailable:
//line /usr/local/go/src/net/http/status.go:191
		_go_fuzz_dep_.CoverTab[43543]++
								return "Service Unavailable"
//line /usr/local/go/src/net/http/status.go:192
		// _ = "end of CoverTab[43543]"
	case StatusGatewayTimeout:
//line /usr/local/go/src/net/http/status.go:193
		_go_fuzz_dep_.CoverTab[43544]++
								return "Gateway Timeout"
//line /usr/local/go/src/net/http/status.go:194
		// _ = "end of CoverTab[43544]"
	case StatusHTTPVersionNotSupported:
//line /usr/local/go/src/net/http/status.go:195
		_go_fuzz_dep_.CoverTab[43545]++
								return "HTTP Version Not Supported"
//line /usr/local/go/src/net/http/status.go:196
		// _ = "end of CoverTab[43545]"
	case StatusVariantAlsoNegotiates:
//line /usr/local/go/src/net/http/status.go:197
		_go_fuzz_dep_.CoverTab[43546]++
								return "Variant Also Negotiates"
//line /usr/local/go/src/net/http/status.go:198
		// _ = "end of CoverTab[43546]"
	case StatusInsufficientStorage:
//line /usr/local/go/src/net/http/status.go:199
		_go_fuzz_dep_.CoverTab[43547]++
								return "Insufficient Storage"
//line /usr/local/go/src/net/http/status.go:200
		// _ = "end of CoverTab[43547]"
	case StatusLoopDetected:
//line /usr/local/go/src/net/http/status.go:201
		_go_fuzz_dep_.CoverTab[43548]++
								return "Loop Detected"
//line /usr/local/go/src/net/http/status.go:202
		// _ = "end of CoverTab[43548]"
	case StatusNotExtended:
//line /usr/local/go/src/net/http/status.go:203
		_go_fuzz_dep_.CoverTab[43549]++
								return "Not Extended"
//line /usr/local/go/src/net/http/status.go:204
		// _ = "end of CoverTab[43549]"
	case StatusNetworkAuthenticationRequired:
//line /usr/local/go/src/net/http/status.go:205
		_go_fuzz_dep_.CoverTab[43550]++
								return "Network Authentication Required"
//line /usr/local/go/src/net/http/status.go:206
		// _ = "end of CoverTab[43550]"
	default:
//line /usr/local/go/src/net/http/status.go:207
		_go_fuzz_dep_.CoverTab[43551]++
								return ""
//line /usr/local/go/src/net/http/status.go:208
		// _ = "end of CoverTab[43551]"
	}
//line /usr/local/go/src/net/http/status.go:209
	// _ = "end of CoverTab[43488]"
}

//line /usr/local/go/src/net/http/status.go:210
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/status.go:210
var _ = _go_fuzz_dep_.CoverTab
