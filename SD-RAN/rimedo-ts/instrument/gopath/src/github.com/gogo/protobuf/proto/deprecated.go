// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2018 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:32
)

import "errors"

// Deprecated: do not use.
type Stats struct{ Emalloc, Dmalloc, Encode, Decode, Chit, Cmiss, Size uint64 }

// Deprecated: do not use.
func GetStats() Stats {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:40
	_go_fuzz_dep_.CoverTab[107797]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:40
	return Stats{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:40
	// _ = "end of CoverTab[107797]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:40
}

// Deprecated: do not use.
func MarshalMessageSet(interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:43
	_go_fuzz_dep_.CoverTab[107798]++
												return nil, errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:44
	// _ = "end of CoverTab[107798]"
}

// Deprecated: do not use.
func UnmarshalMessageSet([]byte, interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:48
	_go_fuzz_dep_.CoverTab[107799]++
												return errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:49
	// _ = "end of CoverTab[107799]"
}

// Deprecated: do not use.
func MarshalMessageSetJSON(interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:53
	_go_fuzz_dep_.CoverTab[107800]++
												return nil, errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:54
	// _ = "end of CoverTab[107800]"
}

// Deprecated: do not use.
func UnmarshalMessageSetJSON([]byte, interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:58
	_go_fuzz_dep_.CoverTab[107801]++
												return errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:59
	// _ = "end of CoverTab[107801]"
}

// Deprecated: do not use.
func RegisterMessageSetType(Message, int32, string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:63
	_go_fuzz_dep_.CoverTab[107802]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:63
	// _ = "end of CoverTab[107802]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:63
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/deprecated.go:63
var _ = _go_fuzz_dep_.CoverTab
