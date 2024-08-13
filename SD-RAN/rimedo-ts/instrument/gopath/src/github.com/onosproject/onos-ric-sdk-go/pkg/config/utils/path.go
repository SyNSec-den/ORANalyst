// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
package utils

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:5
)

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// GnmiFullPath builds the full path from the prefix and path.
func GnmiFullPath(prefix, path *pb.Path) *pb.Path {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:17
	_go_fuzz_dep_.CoverTab[193616]++
														fullPath := &pb.Path{Origin: path.Origin}
														if path.GetElem() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:19
		_go_fuzz_dep_.CoverTab[193618]++
															fullPath.Elem = append(prefix.GetElem(), path.GetElem()...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:20
		// _ = "end of CoverTab[193618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:21
		_go_fuzz_dep_.CoverTab[193619]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:21
		// _ = "end of CoverTab[193619]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:21
	// _ = "end of CoverTab[193616]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:21
	_go_fuzz_dep_.CoverTab[193617]++
														return fullPath
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:22
	// _ = "end of CoverTab[193617]"
}

func writeSafeString(b *strings.Builder, s string, esc rune) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:25
	_go_fuzz_dep_.CoverTab[193620]++
														for _, c := range s {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:26
		_go_fuzz_dep_.CoverTab[193621]++
															if c == esc || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:27
			_go_fuzz_dep_.CoverTab[193623]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:27
			return c == '\\'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:27
			// _ = "end of CoverTab[193623]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:27
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:27
			_go_fuzz_dep_.CoverTab[193624]++
																b.WriteRune('\\')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:28
			// _ = "end of CoverTab[193624]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:29
			_go_fuzz_dep_.CoverTab[193625]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:29
			// _ = "end of CoverTab[193625]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:29
		// _ = "end of CoverTab[193621]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:29
		_go_fuzz_dep_.CoverTab[193622]++
															b.WriteRune(c)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:30
		// _ = "end of CoverTab[193622]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:31
	// _ = "end of CoverTab[193620]"
}

// ToXPath builds a human-readable form of a gnmi path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:34
// e.g. /a/b/c[e=f]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:36
func ToXPath(path *pb.Path) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:36
	_go_fuzz_dep_.CoverTab[193626]++
														if path == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:37
		_go_fuzz_dep_.CoverTab[193628]++
															return "/"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:38
		// _ = "end of CoverTab[193628]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:39
		_go_fuzz_dep_.CoverTab[193629]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:39
		if len(path.Elem) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:39
			_go_fuzz_dep_.CoverTab[193630]++
																return toXPathV04(path)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:40
			// _ = "end of CoverTab[193630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
			_go_fuzz_dep_.CoverTab[193631]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
			// _ = "end of CoverTab[193631]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
		// _ = "end of CoverTab[193629]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
	// _ = "end of CoverTab[193626]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:41
	_go_fuzz_dep_.CoverTab[193627]++
														return "/"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:42
	// _ = "end of CoverTab[193627]"
}

// StrPathElem builds a human-readable form of a list of path elements.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:45
// e.g. /a/b/c[e=f]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:47
func StrPathElem(pathElem []*pb.PathElem) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:47
	_go_fuzz_dep_.CoverTab[193632]++
														b := &strings.Builder{}
														for _, elm := range pathElem {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:49
		_go_fuzz_dep_.CoverTab[193634]++
															b.WriteRune('/')
															writeSafeString(b, elm.Name, '/')
															if len(elm.Key) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:52
			_go_fuzz_dep_.CoverTab[193635]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:56
			keys := make([]string, 0, len(elm.Key))
			for k := range elm.Key {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:57
				_go_fuzz_dep_.CoverTab[193637]++
																	keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:58
				// _ = "end of CoverTab[193637]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:59
			// _ = "end of CoverTab[193635]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:59
			_go_fuzz_dep_.CoverTab[193636]++
																sort.Strings(keys)
																for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:61
				_go_fuzz_dep_.CoverTab[193638]++
																	b.WriteRune('[')
																	b.WriteString(k)
																	b.WriteRune('=')
																	writeSafeString(b, elm.Key[k], ']')
																	b.WriteRune(']')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:66
				// _ = "end of CoverTab[193638]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:67
			// _ = "end of CoverTab[193636]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:68
			_go_fuzz_dep_.CoverTab[193639]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:68
			// _ = "end of CoverTab[193639]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:68
		// _ = "end of CoverTab[193634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:69
	// _ = "end of CoverTab[193632]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:69
	_go_fuzz_dep_.CoverTab[193633]++
														return b.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:70
	// _ = "end of CoverTab[193633]"
}

// strPathV04 handles the v0.4 gnmi and later path.Elem member.
func toXPathV04(path *pb.Path) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:74
	_go_fuzz_dep_.CoverTab[193640]++
														pathElem := path.Elem
														return StrPathElem(pathElem)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:76
	// _ = "end of CoverTab[193640]"
}

var (
	idPattern	= `[a-zA-Z_][a-zA-Z\d\_\-\.]*`
	// YANG identifiers must follow RFC 6020:
	// https://tools.ietf.org/html/rfc6020#section-6.2.
	idRe	= regexp.MustCompile(`^` + idPattern + `$`)
	// The sting representation of List key value pairs must follow the
	// following pattern: [key=value], where key is the List key leaf name,
	// and value is the string representation of key leaf value.
	kvRe	= regexp.MustCompile(`^\[` +

		idPattern + `=` +

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:92
		`((?s).+)` +
		`\]$`)
)

// splitPath splits a string representation of path into []string. Path
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// elements are separated by '/'. String splitting scans from left to right. A
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// '[' marks the beginning of a List key value pair substring. A List key value
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// pair string ends at the first ']' encountered. Neither an escaped '[', i.e.,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// `\[`, nor an escaped ']', i.e., `\]`, serves as the boundary of a List key
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// value pair string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// Within a List key value string, '/', '[' and ']' are treated differently:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//  1. A '/' does not act as a separator, and is allowed to be part of a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//     List key leaf value.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//  2. A '[' is allowed within a List key value. '[' and `\[` are
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//     equivalent within a List key value.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//  3. If a ']' needs to be part of a List key value, it must be escaped as
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//     '\]'. The first unescaped ']' terminates a List key value string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// Outside of any List key value pair string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//  1. A ']' without a matching '[' does not generate any error in this
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//     API. This error is caught later by another API.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//  2. A '[' without an closing ']' is treated as an error, because it
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//     indicates an incomplete List key leaf value string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// For example, "/a/b/c" is split into []string{"a", "b", "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// "/a/b[k=eth1/1]/c" is split into []string{"a", "b[k=eth1/1]", "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// `/a/b/[k=v\]]/c` is split into []string{"a", "b", `[k=v\]]`, "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// "a/b][k=v]/c" is split into []string{"a", "b][k=v]", "c"}. The invalid List
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// name "b]" error will be caught later by another API. "/a/b[k=v/c" generates
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:96
// an error because of incomplete List key value pair string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:128
func splitPath(str string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:128
	_go_fuzz_dep_.CoverTab[193641]++
														var path []string
														str += "/"

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:133
	insideBrackets := false

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:136
	begin := 0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:139
	end := 0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:142
	for end < len(str) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:142
		_go_fuzz_dep_.CoverTab[193644]++
															switch str[end] {
		case '/':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:144
			_go_fuzz_dep_.CoverTab[193645]++
																if !insideBrackets {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:145
				_go_fuzz_dep_.CoverTab[193651]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:148
				if end > begin {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:148
					_go_fuzz_dep_.CoverTab[193653]++
																		path = append(path, str[begin:end])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:149
					// _ = "end of CoverTab[193653]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:150
					_go_fuzz_dep_.CoverTab[193654]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:150
					// _ = "end of CoverTab[193654]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:150
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:150
				// _ = "end of CoverTab[193651]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:150
				_go_fuzz_dep_.CoverTab[193652]++
																	end++
																	begin = end
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:152
				// _ = "end of CoverTab[193652]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:153
				_go_fuzz_dep_.CoverTab[193655]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:156
				end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:156
				// _ = "end of CoverTab[193655]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:157
			// _ = "end of CoverTab[193645]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:158
			_go_fuzz_dep_.CoverTab[193646]++
																if (end == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				_go_fuzz_dep_.CoverTab[193656]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				return str[end-1] != '\\'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				// _ = "end of CoverTab[193656]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				_go_fuzz_dep_.CoverTab[193657]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				return !insideBrackets
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				// _ = "end of CoverTab[193657]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:159
				_go_fuzz_dep_.CoverTab[193658]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:162
				insideBrackets = true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:162
				// _ = "end of CoverTab[193658]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:163
				_go_fuzz_dep_.CoverTab[193659]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:163
				// _ = "end of CoverTab[193659]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:163
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:163
			// _ = "end of CoverTab[193646]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:163
			_go_fuzz_dep_.CoverTab[193647]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:164
			// _ = "end of CoverTab[193647]"
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:165
			_go_fuzz_dep_.CoverTab[193648]++
																if (end == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				_go_fuzz_dep_.CoverTab[193660]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				return str[end-1] != '\\'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				// _ = "end of CoverTab[193660]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				_go_fuzz_dep_.CoverTab[193661]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				return insideBrackets
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				// _ = "end of CoverTab[193661]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:166
				_go_fuzz_dep_.CoverTab[193662]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:169
				insideBrackets = false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:169
				// _ = "end of CoverTab[193662]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:170
				_go_fuzz_dep_.CoverTab[193663]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:170
				// _ = "end of CoverTab[193663]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:170
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:170
			// _ = "end of CoverTab[193648]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:170
			_go_fuzz_dep_.CoverTab[193649]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:171
			// _ = "end of CoverTab[193649]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:172
			_go_fuzz_dep_.CoverTab[193650]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:173
			// _ = "end of CoverTab[193650]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:174
		// _ = "end of CoverTab[193644]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:175
	// _ = "end of CoverTab[193641]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:175
	_go_fuzz_dep_.CoverTab[193642]++

														if insideBrackets {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:177
		_go_fuzz_dep_.CoverTab[193664]++
															return nil, fmt.Errorf("missing ] in path string: %s", str)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:178
		// _ = "end of CoverTab[193664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:179
		_go_fuzz_dep_.CoverTab[193665]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:179
		// _ = "end of CoverTab[193665]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:179
	// _ = "end of CoverTab[193642]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:179
	_go_fuzz_dep_.CoverTab[193643]++
														return path, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:180
	// _ = "end of CoverTab[193643]"
}

// parseKeyValueString parses a List key-value pair, and returns a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// map[string]string whose key is the List key leaf name and whose value is the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// string representation of List key leaf value. The input path-valur pairs are
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// encoded using the following pattern: [k1=v1][k2=v2]..., where k1 and k2 must be
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// valid YANG identifiers, v1 and v2 can be any non-empty strings where any ']'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// must be escapced by an '\'. Any malformed key-value pair generates an error.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// For example, given
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
//	"[k1=v1][k2=v2]",
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
// this API returns
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:183
//	map[string]string{"k1": "v1", "k2": "v2"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:193
func parseKeyValueString(str string) (map[string]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:193
	_go_fuzz_dep_.CoverTab[193666]++
														keyValuePairs := make(map[string]string)

														begin := 0

														end := 0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:201
	insideBrackets := false

	for end < len(str) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:203
		_go_fuzz_dep_.CoverTab[193669]++
															switch str[end] {
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:205
			_go_fuzz_dep_.CoverTab[193670]++
																if (end == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				_go_fuzz_dep_.CoverTab[193675]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				return str[end-1] != '\\'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				// _ = "end of CoverTab[193675]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				_go_fuzz_dep_.CoverTab[193676]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				return !insideBrackets
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				// _ = "end of CoverTab[193676]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:206
				_go_fuzz_dep_.CoverTab[193677]++
																	insideBrackets = true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:207
				// _ = "end of CoverTab[193677]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:208
				_go_fuzz_dep_.CoverTab[193678]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:208
				// _ = "end of CoverTab[193678]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:208
			// _ = "end of CoverTab[193670]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:208
			_go_fuzz_dep_.CoverTab[193671]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:209
			// _ = "end of CoverTab[193671]"
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:210
			_go_fuzz_dep_.CoverTab[193672]++
																if (end == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				_go_fuzz_dep_.CoverTab[193679]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				return str[end-1] != '\\'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				// _ = "end of CoverTab[193679]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				_go_fuzz_dep_.CoverTab[193680]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				return insideBrackets
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				// _ = "end of CoverTab[193680]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:211
				_go_fuzz_dep_.CoverTab[193681]++
																	insideBrackets = false
																	keyValue := str[begin : end+1]

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:218
				if !kvRe.MatchString(keyValue) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:218
					_go_fuzz_dep_.CoverTab[193683]++
																		return nil, fmt.Errorf("malformed List key-value pair string: %s, in: %s", keyValue, str)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:219
					// _ = "end of CoverTab[193683]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:220
					_go_fuzz_dep_.CoverTab[193684]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:220
					// _ = "end of CoverTab[193684]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:220
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:220
				// _ = "end of CoverTab[193681]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:220
				_go_fuzz_dep_.CoverTab[193682]++
																	keyValue = keyValue[1 : len(keyValue)-1]
																	i := strings.Index(keyValue, "=")
																	key, val := keyValue[:i], keyValue[i+1:]

																	val = strings.Replace(val, `\]`, `]`, -1)
																	val = strings.Replace(val, `\[`, `[`, -1)
																	keyValuePairs[key] = val
																	begin = end + 1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:228
				// _ = "end of CoverTab[193682]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:229
				_go_fuzz_dep_.CoverTab[193685]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:229
				// _ = "end of CoverTab[193685]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:229
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:229
			// _ = "end of CoverTab[193672]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:229
			_go_fuzz_dep_.CoverTab[193673]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:230
			// _ = "end of CoverTab[193673]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:231
			_go_fuzz_dep_.CoverTab[193674]++
																end++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:232
			// _ = "end of CoverTab[193674]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:233
		// _ = "end of CoverTab[193669]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:234
	// _ = "end of CoverTab[193666]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:234
	_go_fuzz_dep_.CoverTab[193667]++

														if begin < end {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:236
		_go_fuzz_dep_.CoverTab[193686]++
															return nil, fmt.Errorf("malformed List key-value pair string: %s", str)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:237
		// _ = "end of CoverTab[193686]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:238
		_go_fuzz_dep_.CoverTab[193687]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:238
		// _ = "end of CoverTab[193687]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:238
	// _ = "end of CoverTab[193667]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:238
	_go_fuzz_dep_.CoverTab[193668]++

														return keyValuePairs, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:240
	// _ = "end of CoverTab[193668]"
}

// parseElement parses a split path element, and returns the parsed elements.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// Two types of path elements are supported:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// 1. Non-List schema node names which must be valid YANG identifiers. A valid
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// schema node name is returned as it is. For example, given "abc", this API
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// returns []interface{"abc"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// 2. List elements following this pattern: list-name[k1=v1], where list-name
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// is the substring from the beginning of the input string to the first '[', k1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// is the substring from the letter after '[' to the first '=', and v1 is the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// substring from the letter after '=' to the first unescaped ']'. list-name
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// and k1 must be valid YANG identifier, and v1 can be any non-empty string
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// where ']' is escaped by '\'. A List element is parsed into two parts: List
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// name and List key value pair(s). List key value pairs are saved in a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// map[string]string whose key is List key leaf name and whose value is the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// string representation of List key leaf value. For example, given
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//	"list-name[k1=v1]",
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// this API returns
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//	[]interface{}{"list-name", map[string]string{"k1": "v1"}}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
// Multi-key List elements follow a similar pattern:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:243
//	list-name[k1=v1]...[kN=vN].
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:264
func parseElement(elem string) ([]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:264
	_go_fuzz_dep_.CoverTab[193688]++
														i := strings.Index(elem, "[")
														if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:266
		_go_fuzz_dep_.CoverTab[193692]++
															if !idRe.MatchString(elem) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:267
			_go_fuzz_dep_.CoverTab[193694]++
																return nil, fmt.Errorf("invalid node name: %q", elem)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:268
			// _ = "end of CoverTab[193694]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:269
			_go_fuzz_dep_.CoverTab[193695]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:269
			// _ = "end of CoverTab[193695]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:269
		// _ = "end of CoverTab[193692]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:269
		_go_fuzz_dep_.CoverTab[193693]++
															return []interface{}{elem}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:270
		// _ = "end of CoverTab[193693]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:271
		_go_fuzz_dep_.CoverTab[193696]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:271
		// _ = "end of CoverTab[193696]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:271
	// _ = "end of CoverTab[193688]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:271
	_go_fuzz_dep_.CoverTab[193689]++

														listName := elem[:i]
														if !idRe.MatchString(listName) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:274
		_go_fuzz_dep_.CoverTab[193697]++
															return nil, fmt.Errorf("invalid List name: %q, in: %s", listName, elem)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:275
		// _ = "end of CoverTab[193697]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:276
		_go_fuzz_dep_.CoverTab[193698]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:276
		// _ = "end of CoverTab[193698]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:276
	// _ = "end of CoverTab[193689]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:276
	_go_fuzz_dep_.CoverTab[193690]++
														keyValuePairs, err := parseKeyValueString(elem[i:])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:278
		_go_fuzz_dep_.CoverTab[193699]++
															return nil, fmt.Errorf("invalid path element %s: %v", elem, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:279
		// _ = "end of CoverTab[193699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:280
		_go_fuzz_dep_.CoverTab[193700]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:280
		// _ = "end of CoverTab[193700]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:280
	// _ = "end of CoverTab[193690]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:280
	_go_fuzz_dep_.CoverTab[193691]++
														return []interface{}{listName, keyValuePairs}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:281
	// _ = "end of CoverTab[193691]"
}

// ParseStringPath parses a string path and produces a []interface{} of parsed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// path elements. Path elements in a string path are separated by '/'. Each
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// path element can either be a schema node name or a List path element. Schema
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// node names must be valid YANG identifiers. A List path element is encoded
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// using the following pattern: list-name[key1=value1]...[keyN=valueN]. Each
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// List path element generates two parsed path elements: List name and a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// map[string]string containing List key-value pairs with value(s) in string
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// representation. A '/' within a List key value pair string, i.e., between a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// pair of '[' and ']', does not serve as a path separator, and is allowed to be
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// part of a List key leaf value. For example, given a string path:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//	"/a/list-name[k=v/v]/c",
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// this API returns:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//	[]interface{}{"a", "list-name", map[string]string{"k": "v/v"}, "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// String path parsing consists of two passes. In the first pass, the input
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// string is split into []string using valid separator '/'. An incomplete List
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// key value string, i.e, a '[' which starts a List key value string without a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// closing ']', in input string generates an error. In the above example, this
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// pass produces:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//	[]string{"a", "list-name[k=v/v]", "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// In the second pass, each element in split []string is parsed checking syntax
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// and pattern correctness. Errors are generated for invalid YANG identifiers,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// malformed List key-value string, etc.. In the above example, the second pass
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
// produces:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:284
//	[]interface{}{"a", "list-name", map[string]string{"k", "v/v"}, "c"}.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:309
func ParseStringPath(stringPath string) ([]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:309
	_go_fuzz_dep_.CoverTab[193701]++
														elems, err := splitPath(stringPath)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:311
		_go_fuzz_dep_.CoverTab[193704]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:312
		// _ = "end of CoverTab[193704]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:313
		_go_fuzz_dep_.CoverTab[193705]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:313
		// _ = "end of CoverTab[193705]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:313
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:313
	// _ = "end of CoverTab[193701]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:313
	_go_fuzz_dep_.CoverTab[193702]++

														var path []interface{}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:318
	for _, elem := range elems {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:318
		_go_fuzz_dep_.CoverTab[193706]++
															parts, err := parseElement(elem)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:320
			_go_fuzz_dep_.CoverTab[193708]++
																return nil, fmt.Errorf("invalid string path %s: %v", stringPath, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:321
			// _ = "end of CoverTab[193708]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:322
			_go_fuzz_dep_.CoverTab[193709]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:322
			// _ = "end of CoverTab[193709]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:322
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:322
		// _ = "end of CoverTab[193706]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:322
		_go_fuzz_dep_.CoverTab[193707]++
															path = append(path, parts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:323
		// _ = "end of CoverTab[193707]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:324
	// _ = "end of CoverTab[193702]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:324
	_go_fuzz_dep_.CoverTab[193703]++

														return path, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:326
	// _ = "end of CoverTab[193703]"
}

// ToGNMIPath parses an xpath string into a gnmi Path struct defined in gnmi
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
// proto. Path convention can be found in
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
// https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-path-conventions.md
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
// For example, xpath /interfaces/interface[name=Ethernet1/2/3]/state/counters
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
// will be parsed to:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	elem: <name: "interfaces" >
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	elem: <
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	    name: "interface"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	    key: <
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	        key: "name"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	        value: "Ethernet1/2/3"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	    >
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	>
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	elem: <name: "state" >
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:329
//	elem: <name: "counters" >
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:346
func ToGNMIPath(xpath string) (*pb.Path, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:346
	_go_fuzz_dep_.CoverTab[193710]++
														xpathElements, err := ParseStringPath(xpath)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:348
		_go_fuzz_dep_.CoverTab[193713]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:349
		// _ = "end of CoverTab[193713]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:350
		_go_fuzz_dep_.CoverTab[193714]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:350
		// _ = "end of CoverTab[193714]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:350
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:350
	// _ = "end of CoverTab[193710]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:350
	_go_fuzz_dep_.CoverTab[193711]++
														var pbPathElements []*pb.PathElem
														for _, elem := range xpathElements {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:352
		_go_fuzz_dep_.CoverTab[193715]++
															switch v := elem.(type) {
		case string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:354
			_go_fuzz_dep_.CoverTab[193716]++
																pbPathElements = append(pbPathElements, &pb.PathElem{Name: v})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:355
			// _ = "end of CoverTab[193716]"
		case map[string]string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:356
			_go_fuzz_dep_.CoverTab[193717]++
																n := len(pbPathElements)
																if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:358
				_go_fuzz_dep_.CoverTab[193721]++
																	return nil, fmt.Errorf("missing name before key-value list")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:359
				// _ = "end of CoverTab[193721]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:360
				_go_fuzz_dep_.CoverTab[193722]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:360
				// _ = "end of CoverTab[193722]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:360
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:360
			// _ = "end of CoverTab[193717]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:360
			_go_fuzz_dep_.CoverTab[193718]++
																if pbPathElements[n-1].Key != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:361
				_go_fuzz_dep_.CoverTab[193723]++
																	return nil, fmt.Errorf("two subsequent key-value lists")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:362
				// _ = "end of CoverTab[193723]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:363
				_go_fuzz_dep_.CoverTab[193724]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:363
				// _ = "end of CoverTab[193724]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:363
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:363
			// _ = "end of CoverTab[193718]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:363
			_go_fuzz_dep_.CoverTab[193719]++
																pbPathElements[n-1].Key = v
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:364
			// _ = "end of CoverTab[193719]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:365
			_go_fuzz_dep_.CoverTab[193720]++
																return nil, fmt.Errorf("wrong data type: %T", v)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:366
			// _ = "end of CoverTab[193720]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:367
		// _ = "end of CoverTab[193715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:368
	// _ = "end of CoverTab[193711]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:368
	_go_fuzz_dep_.CoverTab[193712]++
														return &pb.Path{Elem: pbPathElements}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:369
	// _ = "end of CoverTab[193712]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/utils/path.go:370
var _ = _go_fuzz_dep_.CoverTab
