// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:5
// Package version records versioning information about this module.
package version

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:6
)

import (
	"fmt"
	"strings"
)

// These constants determine the current version of this module.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
// For our release process, we enforce the following rules:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//   - Tagged releases use a tag that is identical to String.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//   - Tagged releases never reference a commit where the String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     contains "devel".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//   - The set of all commits in this repository where String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     does not contain "devel" must have a unique String.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
// Steps for tagging a new release:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  1. Create a new CL.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  2. Update Minor, Patch, and/or PreRelease as necessary.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     PreRelease must not contain the string "devel".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  3. Since the last released minor version, have there been any changes to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     generator that relies on new functionality in the runtime?
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     If yes, then increment RequiredGenerated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  4. Since the last released minor version, have there been any changes to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     the runtime that removes support for old .pb.go source code?
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     If yes, then increment SupportMinimum.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  5. Send out the CL for review and submit it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     Note that the next CL in step 8 must be submitted after this CL
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     without any other CLs in-between.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  6. Tag a new version, where the tag is is the current String.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  7. Write release notes for all notable changes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     between this release and the last release.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  8. Create a new CL.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  9. Update PreRelease to include the string "devel".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//     For example: "" -> "devel" or "rc.1" -> "rc.1.devel"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:13
//  10. Send out the CL for review and submit it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:52
const (
	Major		= 1
	Minor		= 28
	Patch		= 1
	PreRelease	= ""
)

// String formats the version string for this module in semver format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:59
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:59
// Examples:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:59
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:59
//	v1.20.1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:59
//	v1.21.0-rc.1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:65
func String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:65
	_go_fuzz_dep_.CoverTab[59089]++
													v := fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
													if PreRelease != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:67
		_go_fuzz_dep_.CoverTab[59091]++
														v += "-" + PreRelease

		// TODO: Add metadata about the commit or build hash.
		// See https://golang.org/issue/29814
		// See https://golang.org/issue/33533
		var metadata string
		if strings.Contains(PreRelease, "devel") && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:74
			_go_fuzz_dep_.CoverTab[59092]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:74
			return metadata != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:74
			// _ = "end of CoverTab[59092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:74
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:74
			_go_fuzz_dep_.CoverTab[59093]++
															v += "+" + metadata
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:75
			// _ = "end of CoverTab[59093]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:76
			_go_fuzz_dep_.CoverTab[59094]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:76
			// _ = "end of CoverTab[59094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:76
		// _ = "end of CoverTab[59091]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:77
		_go_fuzz_dep_.CoverTab[59095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:77
		// _ = "end of CoverTab[59095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:77
	// _ = "end of CoverTab[59089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:77
	_go_fuzz_dep_.CoverTab[59090]++
													return v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:78
	// _ = "end of CoverTab[59090]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:79
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/version/version.go:79
var _ = _go_fuzz_dep_.CoverTab
