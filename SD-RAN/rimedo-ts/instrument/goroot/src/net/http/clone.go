// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/clone.go:5
package http

//line /usr/local/go/src/net/http/clone.go:5
import (
//line /usr/local/go/src/net/http/clone.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/clone.go:5
)
//line /usr/local/go/src/net/http/clone.go:5
import (
//line /usr/local/go/src/net/http/clone.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/clone.go:5
)

import (
	"mime/multipart"
	"net/textproto"
	"net/url"
)

func cloneURLValues(v url.Values) url.Values {
//line /usr/local/go/src/net/http/clone.go:13
	_go_fuzz_dep_.CoverTab[36882]++
						if v == nil {
//line /usr/local/go/src/net/http/clone.go:14
		_go_fuzz_dep_.CoverTab[36884]++
							return nil
//line /usr/local/go/src/net/http/clone.go:15
		// _ = "end of CoverTab[36884]"
	} else {
//line /usr/local/go/src/net/http/clone.go:16
		_go_fuzz_dep_.CoverTab[36885]++
//line /usr/local/go/src/net/http/clone.go:16
		// _ = "end of CoverTab[36885]"
//line /usr/local/go/src/net/http/clone.go:16
	}
//line /usr/local/go/src/net/http/clone.go:16
	// _ = "end of CoverTab[36882]"
//line /usr/local/go/src/net/http/clone.go:16
	_go_fuzz_dep_.CoverTab[36883]++

//line /usr/local/go/src/net/http/clone.go:19
	return url.Values(Header(v).Clone())
//line /usr/local/go/src/net/http/clone.go:19
	// _ = "end of CoverTab[36883]"
}

func cloneURL(u *url.URL) *url.URL {
//line /usr/local/go/src/net/http/clone.go:22
	_go_fuzz_dep_.CoverTab[36886]++
						if u == nil {
//line /usr/local/go/src/net/http/clone.go:23
		_go_fuzz_dep_.CoverTab[36889]++
							return nil
//line /usr/local/go/src/net/http/clone.go:24
		// _ = "end of CoverTab[36889]"
	} else {
//line /usr/local/go/src/net/http/clone.go:25
		_go_fuzz_dep_.CoverTab[36890]++
//line /usr/local/go/src/net/http/clone.go:25
		// _ = "end of CoverTab[36890]"
//line /usr/local/go/src/net/http/clone.go:25
	}
//line /usr/local/go/src/net/http/clone.go:25
	// _ = "end of CoverTab[36886]"
//line /usr/local/go/src/net/http/clone.go:25
	_go_fuzz_dep_.CoverTab[36887]++
						u2 := new(url.URL)
						*u2 = *u
						if u.User != nil {
//line /usr/local/go/src/net/http/clone.go:28
		_go_fuzz_dep_.CoverTab[36891]++
							u2.User = new(url.Userinfo)
							*u2.User = *u.User
//line /usr/local/go/src/net/http/clone.go:30
		// _ = "end of CoverTab[36891]"
	} else {
//line /usr/local/go/src/net/http/clone.go:31
		_go_fuzz_dep_.CoverTab[36892]++
//line /usr/local/go/src/net/http/clone.go:31
		// _ = "end of CoverTab[36892]"
//line /usr/local/go/src/net/http/clone.go:31
	}
//line /usr/local/go/src/net/http/clone.go:31
	// _ = "end of CoverTab[36887]"
//line /usr/local/go/src/net/http/clone.go:31
	_go_fuzz_dep_.CoverTab[36888]++
						return u2
//line /usr/local/go/src/net/http/clone.go:32
	// _ = "end of CoverTab[36888]"
}

func cloneMultipartForm(f *multipart.Form) *multipart.Form {
//line /usr/local/go/src/net/http/clone.go:35
	_go_fuzz_dep_.CoverTab[36893]++
						if f == nil {
//line /usr/local/go/src/net/http/clone.go:36
		_go_fuzz_dep_.CoverTab[36896]++
							return nil
//line /usr/local/go/src/net/http/clone.go:37
		// _ = "end of CoverTab[36896]"
	} else {
//line /usr/local/go/src/net/http/clone.go:38
		_go_fuzz_dep_.CoverTab[36897]++
//line /usr/local/go/src/net/http/clone.go:38
		// _ = "end of CoverTab[36897]"
//line /usr/local/go/src/net/http/clone.go:38
	}
//line /usr/local/go/src/net/http/clone.go:38
	// _ = "end of CoverTab[36893]"
//line /usr/local/go/src/net/http/clone.go:38
	_go_fuzz_dep_.CoverTab[36894]++
						f2 := &multipart.Form{
		Value: (map[string][]string)(Header(f.Value).Clone()),
	}
	if f.File != nil {
//line /usr/local/go/src/net/http/clone.go:42
		_go_fuzz_dep_.CoverTab[36898]++
							m := make(map[string][]*multipart.FileHeader)
							for k, vv := range f.File {
//line /usr/local/go/src/net/http/clone.go:44
			_go_fuzz_dep_.CoverTab[36900]++
								vv2 := make([]*multipart.FileHeader, len(vv))
								for i, v := range vv {
//line /usr/local/go/src/net/http/clone.go:46
				_go_fuzz_dep_.CoverTab[36902]++
									vv2[i] = cloneMultipartFileHeader(v)
//line /usr/local/go/src/net/http/clone.go:47
				// _ = "end of CoverTab[36902]"
			}
//line /usr/local/go/src/net/http/clone.go:48
			// _ = "end of CoverTab[36900]"
//line /usr/local/go/src/net/http/clone.go:48
			_go_fuzz_dep_.CoverTab[36901]++
								m[k] = vv2
//line /usr/local/go/src/net/http/clone.go:49
			// _ = "end of CoverTab[36901]"
		}
//line /usr/local/go/src/net/http/clone.go:50
		// _ = "end of CoverTab[36898]"
//line /usr/local/go/src/net/http/clone.go:50
		_go_fuzz_dep_.CoverTab[36899]++
							f2.File = m
//line /usr/local/go/src/net/http/clone.go:51
		// _ = "end of CoverTab[36899]"
	} else {
//line /usr/local/go/src/net/http/clone.go:52
		_go_fuzz_dep_.CoverTab[36903]++
//line /usr/local/go/src/net/http/clone.go:52
		// _ = "end of CoverTab[36903]"
//line /usr/local/go/src/net/http/clone.go:52
	}
//line /usr/local/go/src/net/http/clone.go:52
	// _ = "end of CoverTab[36894]"
//line /usr/local/go/src/net/http/clone.go:52
	_go_fuzz_dep_.CoverTab[36895]++
						return f2
//line /usr/local/go/src/net/http/clone.go:53
	// _ = "end of CoverTab[36895]"
}

func cloneMultipartFileHeader(fh *multipart.FileHeader) *multipart.FileHeader {
//line /usr/local/go/src/net/http/clone.go:56
	_go_fuzz_dep_.CoverTab[36904]++
						if fh == nil {
//line /usr/local/go/src/net/http/clone.go:57
		_go_fuzz_dep_.CoverTab[36906]++
							return nil
//line /usr/local/go/src/net/http/clone.go:58
		// _ = "end of CoverTab[36906]"
	} else {
//line /usr/local/go/src/net/http/clone.go:59
		_go_fuzz_dep_.CoverTab[36907]++
//line /usr/local/go/src/net/http/clone.go:59
		// _ = "end of CoverTab[36907]"
//line /usr/local/go/src/net/http/clone.go:59
	}
//line /usr/local/go/src/net/http/clone.go:59
	// _ = "end of CoverTab[36904]"
//line /usr/local/go/src/net/http/clone.go:59
	_go_fuzz_dep_.CoverTab[36905]++
						fh2 := new(multipart.FileHeader)
						*fh2 = *fh
						fh2.Header = textproto.MIMEHeader(Header(fh.Header).Clone())
						return fh2
//line /usr/local/go/src/net/http/clone.go:63
	// _ = "end of CoverTab[36905]"
}

// cloneOrMakeHeader invokes Header.Clone but if the
//line /usr/local/go/src/net/http/clone.go:66
// result is nil, it'll instead make and return a non-nil Header.
//line /usr/local/go/src/net/http/clone.go:68
func cloneOrMakeHeader(hdr Header) Header {
//line /usr/local/go/src/net/http/clone.go:68
	_go_fuzz_dep_.CoverTab[36908]++
						clone := hdr.Clone()
						if clone == nil {
//line /usr/local/go/src/net/http/clone.go:70
		_go_fuzz_dep_.CoverTab[36910]++
							clone = make(Header)
//line /usr/local/go/src/net/http/clone.go:71
		// _ = "end of CoverTab[36910]"
	} else {
//line /usr/local/go/src/net/http/clone.go:72
		_go_fuzz_dep_.CoverTab[36911]++
//line /usr/local/go/src/net/http/clone.go:72
		// _ = "end of CoverTab[36911]"
//line /usr/local/go/src/net/http/clone.go:72
	}
//line /usr/local/go/src/net/http/clone.go:72
	// _ = "end of CoverTab[36908]"
//line /usr/local/go/src/net/http/clone.go:72
	_go_fuzz_dep_.CoverTab[36909]++
						return clone
//line /usr/local/go/src/net/http/clone.go:73
	// _ = "end of CoverTab[36909]"
}

//line /usr/local/go/src/net/http/clone.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/clone.go:74
var _ = _go_fuzz_dep_.CoverTab
