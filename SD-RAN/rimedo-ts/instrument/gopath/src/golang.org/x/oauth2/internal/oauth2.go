// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
package internal

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:5
)

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// ParseKey converts the binary contents of a private key file
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:15
// to an *rsa.PrivateKey. It detects whether the private key is in a
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:15
// PEM container or not. If so, it extracts the the private key
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:15
// from PEM container before conversion. It only supports PEM
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:15
// containers with no passphrase.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:20
func ParseKey(key []byte) (*rsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:20
	_go_fuzz_dep_.CoverTab[184048]++
											block, _ := pem.Decode(key)
											if block != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:22
		_go_fuzz_dep_.CoverTab[184052]++
												key = block.Bytes
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:23
		// _ = "end of CoverTab[184052]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:24
		_go_fuzz_dep_.CoverTab[184053]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:24
		// _ = "end of CoverTab[184053]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:24
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:24
	// _ = "end of CoverTab[184048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:24
	_go_fuzz_dep_.CoverTab[184049]++
											parsedKey, err := x509.ParsePKCS8PrivateKey(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:26
		_go_fuzz_dep_.CoverTab[184054]++
												parsedKey, err = x509.ParsePKCS1PrivateKey(key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:28
			_go_fuzz_dep_.CoverTab[184055]++
													return nil, fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:29
			// _ = "end of CoverTab[184055]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:30
			_go_fuzz_dep_.CoverTab[184056]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:30
			// _ = "end of CoverTab[184056]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:30
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:30
		// _ = "end of CoverTab[184054]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:31
		_go_fuzz_dep_.CoverTab[184057]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:31
		// _ = "end of CoverTab[184057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:31
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:31
	// _ = "end of CoverTab[184049]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:31
	_go_fuzz_dep_.CoverTab[184050]++
											parsed, ok := parsedKey.(*rsa.PrivateKey)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:33
		_go_fuzz_dep_.CoverTab[184058]++
												return nil, errors.New("private key is invalid")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:34
		// _ = "end of CoverTab[184058]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:35
		_go_fuzz_dep_.CoverTab[184059]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:35
		// _ = "end of CoverTab[184059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:35
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:35
	// _ = "end of CoverTab[184050]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:35
	_go_fuzz_dep_.CoverTab[184051]++
											return parsed, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:36
	// _ = "end of CoverTab[184051]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/oauth2.go:37
var _ = _go_fuzz_dep_.CoverTab
