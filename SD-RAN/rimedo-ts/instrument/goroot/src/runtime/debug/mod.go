// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/runtime/debug/mod.go:5
package debug

//line /usr/local/go/src/runtime/debug/mod.go:5
import (
//line /usr/local/go/src/runtime/debug/mod.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/runtime/debug/mod.go:5
)
//line /usr/local/go/src/runtime/debug/mod.go:5
import (
//line /usr/local/go/src/runtime/debug/mod.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/runtime/debug/mod.go:5
)

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// exported from runtime.
func modinfo() string

// ReadBuildInfo returns the build information embedded
//line /usr/local/go/src/runtime/debug/mod.go:17
// in the running binary. The information is available only
//line /usr/local/go/src/runtime/debug/mod.go:17
// in binaries built with module support.
//line /usr/local/go/src/runtime/debug/mod.go:20
func ReadBuildInfo() (info *BuildInfo, ok bool) {
//line /usr/local/go/src/runtime/debug/mod.go:20
	_go_fuzz_dep_.CoverTab[90734]++
							data := modinfo()
							if len(data) < 32 {
//line /usr/local/go/src/runtime/debug/mod.go:22
		_go_fuzz_dep_.CoverTab[90737]++
								return nil, false
//line /usr/local/go/src/runtime/debug/mod.go:23
		// _ = "end of CoverTab[90737]"
	} else {
//line /usr/local/go/src/runtime/debug/mod.go:24
		_go_fuzz_dep_.CoverTab[90738]++
//line /usr/local/go/src/runtime/debug/mod.go:24
		// _ = "end of CoverTab[90738]"
//line /usr/local/go/src/runtime/debug/mod.go:24
	}
//line /usr/local/go/src/runtime/debug/mod.go:24
	// _ = "end of CoverTab[90734]"
//line /usr/local/go/src/runtime/debug/mod.go:24
	_go_fuzz_dep_.CoverTab[90735]++
							data = data[16 : len(data)-16]
							bi, err := ParseBuildInfo(data)
							if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:27
		_go_fuzz_dep_.CoverTab[90739]++
								return nil, false
//line /usr/local/go/src/runtime/debug/mod.go:28
		// _ = "end of CoverTab[90739]"
	} else {
//line /usr/local/go/src/runtime/debug/mod.go:29
		_go_fuzz_dep_.CoverTab[90740]++
//line /usr/local/go/src/runtime/debug/mod.go:29
		// _ = "end of CoverTab[90740]"
//line /usr/local/go/src/runtime/debug/mod.go:29
	}
//line /usr/local/go/src/runtime/debug/mod.go:29
	// _ = "end of CoverTab[90735]"
//line /usr/local/go/src/runtime/debug/mod.go:29
	_go_fuzz_dep_.CoverTab[90736]++

//line /usr/local/go/src/runtime/debug/mod.go:35
	bi.GoVersion = runtime.Version()

							return bi, true
//line /usr/local/go/src/runtime/debug/mod.go:37
	// _ = "end of CoverTab[90736]"
}

// BuildInfo represents the build information read from a Go binary.
type BuildInfo struct {
	// GoVersion is the version of the Go toolchain that built the binary
	// (for example, "go1.19.2").
	GoVersion	string

	// Path is the package path of the main package for the binary
	// (for example, "golang.org/x/tools/cmd/stringer").
	Path	string

	// Main describes the module that contains the main package for the binary.
	Main	Module

	// Deps describes all the dependency modules, both direct and indirect,
	// that contributed packages to the build of this binary.
	Deps	[]*Module

	// Settings describes the build settings used to build the binary.
	Settings	[]BuildSetting
}

// A Module describes a single module included in a build.
type Module struct {
	Path	string	// module path
	Version	string	// module version
	Sum	string	// checksum
	Replace	*Module	// replaced by this module
}

// A BuildSetting is a key-value pair describing one setting that influenced a build.
//line /usr/local/go/src/runtime/debug/mod.go:69
//
//line /usr/local/go/src/runtime/debug/mod.go:69
// Defined keys include:
//line /usr/local/go/src/runtime/debug/mod.go:69
//
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - -buildmode: the buildmode flag used (typically "exe")
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - -compiler: the compiler toolchain flag used (typically "gc")
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - CGO_ENABLED: the effective CGO_ENABLED environment variable
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - CGO_CFLAGS: the effective CGO_CFLAGS environment variable
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - CGO_CXXFLAGS:  the effective CGO_CPPFLAGS environment variable
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - CGO_LDFLAGS: the effective CGO_CPPFLAGS environment variable
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - GOARCH: the architecture target
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - GOAMD64/GOARM64/GO386/etc: the architecture feature level for GOARCH
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - GOOS: the operating system target
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - vcs: the version control system for the source tree where the build ran
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - vcs.revision: the revision identifier for the current commit or checkout
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - vcs.time: the modification time associated with vcs.revision, in RFC3339 format
//line /usr/local/go/src/runtime/debug/mod.go:69
//   - vcs.modified: true or false indicating whether the source tree had local modifications
//line /usr/local/go/src/runtime/debug/mod.go:87
type BuildSetting struct {
	// Key and Value describe the build setting.
	// Key must not contain an equals sign, space, tab, or newline.
	// Value must not contain newlines ('\n').
	Key, Value string
}

// quoteKey reports whether key is required to be quoted.
func quoteKey(key string) bool {
//line /usr/local/go/src/runtime/debug/mod.go:95
	_go_fuzz_dep_.CoverTab[90741]++
							return len(key) == 0 || func() bool {
//line /usr/local/go/src/runtime/debug/mod.go:96
		_go_fuzz_dep_.CoverTab[90742]++
//line /usr/local/go/src/runtime/debug/mod.go:96
		return strings.ContainsAny(key, "= \t\r\n\"`")
//line /usr/local/go/src/runtime/debug/mod.go:96
		// _ = "end of CoverTab[90742]"
//line /usr/local/go/src/runtime/debug/mod.go:96
	}()
//line /usr/local/go/src/runtime/debug/mod.go:96
	// _ = "end of CoverTab[90741]"
}

// quoteValue reports whether value is required to be quoted.
func quoteValue(value string) bool {
//line /usr/local/go/src/runtime/debug/mod.go:100
	_go_fuzz_dep_.CoverTab[90743]++
							return strings.ContainsAny(value, " \t\r\n\"`")
//line /usr/local/go/src/runtime/debug/mod.go:101
	// _ = "end of CoverTab[90743]"
}

func (bi *BuildInfo) String() string {
//line /usr/local/go/src/runtime/debug/mod.go:104
	_go_fuzz_dep_.CoverTab[90744]++
							buf := new(strings.Builder)
							if bi.GoVersion != "" {
//line /usr/local/go/src/runtime/debug/mod.go:106
		_go_fuzz_dep_.CoverTab[90751]++
								fmt.Fprintf(buf, "go\t%s\n", bi.GoVersion)
//line /usr/local/go/src/runtime/debug/mod.go:107
		// _ = "end of CoverTab[90751]"
	} else {
//line /usr/local/go/src/runtime/debug/mod.go:108
		_go_fuzz_dep_.CoverTab[90752]++
//line /usr/local/go/src/runtime/debug/mod.go:108
		// _ = "end of CoverTab[90752]"
//line /usr/local/go/src/runtime/debug/mod.go:108
	}
//line /usr/local/go/src/runtime/debug/mod.go:108
	// _ = "end of CoverTab[90744]"
//line /usr/local/go/src/runtime/debug/mod.go:108
	_go_fuzz_dep_.CoverTab[90745]++
							if bi.Path != "" {
//line /usr/local/go/src/runtime/debug/mod.go:109
		_go_fuzz_dep_.CoverTab[90753]++
								fmt.Fprintf(buf, "path\t%s\n", bi.Path)
//line /usr/local/go/src/runtime/debug/mod.go:110
		// _ = "end of CoverTab[90753]"
	} else {
//line /usr/local/go/src/runtime/debug/mod.go:111
		_go_fuzz_dep_.CoverTab[90754]++
//line /usr/local/go/src/runtime/debug/mod.go:111
		// _ = "end of CoverTab[90754]"
//line /usr/local/go/src/runtime/debug/mod.go:111
	}
//line /usr/local/go/src/runtime/debug/mod.go:111
	// _ = "end of CoverTab[90745]"
//line /usr/local/go/src/runtime/debug/mod.go:111
	_go_fuzz_dep_.CoverTab[90746]++
							var formatMod func(string, Module)
							formatMod = func(word string, m Module) {
//line /usr/local/go/src/runtime/debug/mod.go:113
		_go_fuzz_dep_.CoverTab[90755]++
								buf.WriteString(word)
								buf.WriteByte('\t')
								buf.WriteString(m.Path)
								buf.WriteByte('\t')
								buf.WriteString(m.Version)
								if m.Replace == nil {
//line /usr/local/go/src/runtime/debug/mod.go:119
			_go_fuzz_dep_.CoverTab[90757]++
									buf.WriteByte('\t')
									buf.WriteString(m.Sum)
//line /usr/local/go/src/runtime/debug/mod.go:121
			// _ = "end of CoverTab[90757]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:122
			_go_fuzz_dep_.CoverTab[90758]++
									buf.WriteByte('\n')
									formatMod("=>", *m.Replace)
//line /usr/local/go/src/runtime/debug/mod.go:124
			// _ = "end of CoverTab[90758]"
		}
//line /usr/local/go/src/runtime/debug/mod.go:125
		// _ = "end of CoverTab[90755]"
//line /usr/local/go/src/runtime/debug/mod.go:125
		_go_fuzz_dep_.CoverTab[90756]++
								buf.WriteByte('\n')
//line /usr/local/go/src/runtime/debug/mod.go:126
		// _ = "end of CoverTab[90756]"
	}
//line /usr/local/go/src/runtime/debug/mod.go:127
	// _ = "end of CoverTab[90746]"
//line /usr/local/go/src/runtime/debug/mod.go:127
	_go_fuzz_dep_.CoverTab[90747]++
							if bi.Main != (Module{}) {
//line /usr/local/go/src/runtime/debug/mod.go:128
		_go_fuzz_dep_.CoverTab[90759]++
								formatMod("mod", bi.Main)
//line /usr/local/go/src/runtime/debug/mod.go:129
		// _ = "end of CoverTab[90759]"
	} else {
//line /usr/local/go/src/runtime/debug/mod.go:130
		_go_fuzz_dep_.CoverTab[90760]++
//line /usr/local/go/src/runtime/debug/mod.go:130
		// _ = "end of CoverTab[90760]"
//line /usr/local/go/src/runtime/debug/mod.go:130
	}
//line /usr/local/go/src/runtime/debug/mod.go:130
	// _ = "end of CoverTab[90747]"
//line /usr/local/go/src/runtime/debug/mod.go:130
	_go_fuzz_dep_.CoverTab[90748]++
							for _, dep := range bi.Deps {
//line /usr/local/go/src/runtime/debug/mod.go:131
		_go_fuzz_dep_.CoverTab[90761]++
								formatMod("dep", *dep)
//line /usr/local/go/src/runtime/debug/mod.go:132
		// _ = "end of CoverTab[90761]"
	}
//line /usr/local/go/src/runtime/debug/mod.go:133
	// _ = "end of CoverTab[90748]"
//line /usr/local/go/src/runtime/debug/mod.go:133
	_go_fuzz_dep_.CoverTab[90749]++
							for _, s := range bi.Settings {
//line /usr/local/go/src/runtime/debug/mod.go:134
		_go_fuzz_dep_.CoverTab[90762]++
								key := s.Key
								if quoteKey(key) {
//line /usr/local/go/src/runtime/debug/mod.go:136
			_go_fuzz_dep_.CoverTab[90765]++
									key = strconv.Quote(key)
//line /usr/local/go/src/runtime/debug/mod.go:137
			// _ = "end of CoverTab[90765]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:138
			_go_fuzz_dep_.CoverTab[90766]++
//line /usr/local/go/src/runtime/debug/mod.go:138
			// _ = "end of CoverTab[90766]"
//line /usr/local/go/src/runtime/debug/mod.go:138
		}
//line /usr/local/go/src/runtime/debug/mod.go:138
		// _ = "end of CoverTab[90762]"
//line /usr/local/go/src/runtime/debug/mod.go:138
		_go_fuzz_dep_.CoverTab[90763]++
								value := s.Value
								if quoteValue(value) {
//line /usr/local/go/src/runtime/debug/mod.go:140
			_go_fuzz_dep_.CoverTab[90767]++
									value = strconv.Quote(value)
//line /usr/local/go/src/runtime/debug/mod.go:141
			// _ = "end of CoverTab[90767]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:142
			_go_fuzz_dep_.CoverTab[90768]++
//line /usr/local/go/src/runtime/debug/mod.go:142
			// _ = "end of CoverTab[90768]"
//line /usr/local/go/src/runtime/debug/mod.go:142
		}
//line /usr/local/go/src/runtime/debug/mod.go:142
		// _ = "end of CoverTab[90763]"
//line /usr/local/go/src/runtime/debug/mod.go:142
		_go_fuzz_dep_.CoverTab[90764]++
								fmt.Fprintf(buf, "build\t%s=%s\n", key, value)
//line /usr/local/go/src/runtime/debug/mod.go:143
		// _ = "end of CoverTab[90764]"
	}
//line /usr/local/go/src/runtime/debug/mod.go:144
	// _ = "end of CoverTab[90749]"
//line /usr/local/go/src/runtime/debug/mod.go:144
	_go_fuzz_dep_.CoverTab[90750]++

							return buf.String()
//line /usr/local/go/src/runtime/debug/mod.go:146
	// _ = "end of CoverTab[90750]"
}

func ParseBuildInfo(data string) (bi *BuildInfo, err error) {
//line /usr/local/go/src/runtime/debug/mod.go:149
	_go_fuzz_dep_.CoverTab[90769]++
							lineNum := 1
							defer func() {
//line /usr/local/go/src/runtime/debug/mod.go:151
		_go_fuzz_dep_.CoverTab[90773]++
								if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:152
			_go_fuzz_dep_.CoverTab[90774]++
									err = fmt.Errorf("could not parse Go build info: line %d: %w", lineNum, err)
//line /usr/local/go/src/runtime/debug/mod.go:153
			// _ = "end of CoverTab[90774]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:154
			_go_fuzz_dep_.CoverTab[90775]++
//line /usr/local/go/src/runtime/debug/mod.go:154
			// _ = "end of CoverTab[90775]"
//line /usr/local/go/src/runtime/debug/mod.go:154
		}
//line /usr/local/go/src/runtime/debug/mod.go:154
		// _ = "end of CoverTab[90773]"
	}()
//line /usr/local/go/src/runtime/debug/mod.go:155
	// _ = "end of CoverTab[90769]"
//line /usr/local/go/src/runtime/debug/mod.go:155
	_go_fuzz_dep_.CoverTab[90770]++

							var (
		pathLine	= "path\t"
		modLine		= "mod\t"
		depLine		= "dep\t"
		repLine		= "=>\t"
		buildLine	= "build\t"
		newline		= "\n"
		tab		= "\t"
	)

	readModuleLine := func(elem []string) (Module, error) {
//line /usr/local/go/src/runtime/debug/mod.go:167
		_go_fuzz_dep_.CoverTab[90776]++
								if len(elem) != 2 && func() bool {
//line /usr/local/go/src/runtime/debug/mod.go:168
			_go_fuzz_dep_.CoverTab[90779]++
//line /usr/local/go/src/runtime/debug/mod.go:168
			return len(elem) != 3
//line /usr/local/go/src/runtime/debug/mod.go:168
			// _ = "end of CoverTab[90779]"
//line /usr/local/go/src/runtime/debug/mod.go:168
		}() {
//line /usr/local/go/src/runtime/debug/mod.go:168
			_go_fuzz_dep_.CoverTab[90780]++
									return Module{}, fmt.Errorf("expected 2 or 3 columns; got %d", len(elem))
//line /usr/local/go/src/runtime/debug/mod.go:169
			// _ = "end of CoverTab[90780]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:170
			_go_fuzz_dep_.CoverTab[90781]++
//line /usr/local/go/src/runtime/debug/mod.go:170
			// _ = "end of CoverTab[90781]"
//line /usr/local/go/src/runtime/debug/mod.go:170
		}
//line /usr/local/go/src/runtime/debug/mod.go:170
		// _ = "end of CoverTab[90776]"
//line /usr/local/go/src/runtime/debug/mod.go:170
		_go_fuzz_dep_.CoverTab[90777]++
								version := elem[1]
								sum := ""
								if len(elem) == 3 {
//line /usr/local/go/src/runtime/debug/mod.go:173
			_go_fuzz_dep_.CoverTab[90782]++
									sum = elem[2]
//line /usr/local/go/src/runtime/debug/mod.go:174
			// _ = "end of CoverTab[90782]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:175
			_go_fuzz_dep_.CoverTab[90783]++
//line /usr/local/go/src/runtime/debug/mod.go:175
			// _ = "end of CoverTab[90783]"
//line /usr/local/go/src/runtime/debug/mod.go:175
		}
//line /usr/local/go/src/runtime/debug/mod.go:175
		// _ = "end of CoverTab[90777]"
//line /usr/local/go/src/runtime/debug/mod.go:175
		_go_fuzz_dep_.CoverTab[90778]++
								return Module{
			Path:		elem[0],
			Version:	version,
			Sum:		sum,
		}, nil
//line /usr/local/go/src/runtime/debug/mod.go:180
		// _ = "end of CoverTab[90778]"
	}
//line /usr/local/go/src/runtime/debug/mod.go:181
	// _ = "end of CoverTab[90770]"
//line /usr/local/go/src/runtime/debug/mod.go:181
	_go_fuzz_dep_.CoverTab[90771]++

							bi = new(BuildInfo)
							var (
		last	*Module
		line	string
		ok	bool
	)

	for len(data) > 0 {
//line /usr/local/go/src/runtime/debug/mod.go:190
		_go_fuzz_dep_.CoverTab[90784]++
								line, data, ok = strings.Cut(data, newline)
								if !ok {
//line /usr/local/go/src/runtime/debug/mod.go:192
			_go_fuzz_dep_.CoverTab[90787]++
									break
//line /usr/local/go/src/runtime/debug/mod.go:193
			// _ = "end of CoverTab[90787]"
		} else {
//line /usr/local/go/src/runtime/debug/mod.go:194
			_go_fuzz_dep_.CoverTab[90788]++
//line /usr/local/go/src/runtime/debug/mod.go:194
			// _ = "end of CoverTab[90788]"
//line /usr/local/go/src/runtime/debug/mod.go:194
		}
//line /usr/local/go/src/runtime/debug/mod.go:194
		// _ = "end of CoverTab[90784]"
//line /usr/local/go/src/runtime/debug/mod.go:194
		_go_fuzz_dep_.CoverTab[90785]++
								switch {
		case strings.HasPrefix(line, pathLine):
//line /usr/local/go/src/runtime/debug/mod.go:196
			_go_fuzz_dep_.CoverTab[90789]++
									elem := line[len(pathLine):]
									bi.Path = string(elem)
//line /usr/local/go/src/runtime/debug/mod.go:198
			// _ = "end of CoverTab[90789]"
		case strings.HasPrefix(line, modLine):
//line /usr/local/go/src/runtime/debug/mod.go:199
			_go_fuzz_dep_.CoverTab[90790]++
									elem := strings.Split(line[len(modLine):], tab)
									last = &bi.Main
									*last, err = readModuleLine(elem)
									if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:203
				_go_fuzz_dep_.CoverTab[90800]++
										return nil, err
//line /usr/local/go/src/runtime/debug/mod.go:204
				// _ = "end of CoverTab[90800]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:205
				_go_fuzz_dep_.CoverTab[90801]++
//line /usr/local/go/src/runtime/debug/mod.go:205
				// _ = "end of CoverTab[90801]"
//line /usr/local/go/src/runtime/debug/mod.go:205
			}
//line /usr/local/go/src/runtime/debug/mod.go:205
			// _ = "end of CoverTab[90790]"
		case strings.HasPrefix(line, depLine):
//line /usr/local/go/src/runtime/debug/mod.go:206
			_go_fuzz_dep_.CoverTab[90791]++
									elem := strings.Split(line[len(depLine):], tab)
									last = new(Module)
									bi.Deps = append(bi.Deps, last)
									*last, err = readModuleLine(elem)
									if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:211
				_go_fuzz_dep_.CoverTab[90802]++
										return nil, err
//line /usr/local/go/src/runtime/debug/mod.go:212
				// _ = "end of CoverTab[90802]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:213
				_go_fuzz_dep_.CoverTab[90803]++
//line /usr/local/go/src/runtime/debug/mod.go:213
				// _ = "end of CoverTab[90803]"
//line /usr/local/go/src/runtime/debug/mod.go:213
			}
//line /usr/local/go/src/runtime/debug/mod.go:213
			// _ = "end of CoverTab[90791]"
		case strings.HasPrefix(line, repLine):
//line /usr/local/go/src/runtime/debug/mod.go:214
			_go_fuzz_dep_.CoverTab[90792]++
									elem := strings.Split(line[len(repLine):], tab)
									if len(elem) != 3 {
//line /usr/local/go/src/runtime/debug/mod.go:216
				_go_fuzz_dep_.CoverTab[90804]++
										return nil, fmt.Errorf("expected 3 columns for replacement; got %d", len(elem))
//line /usr/local/go/src/runtime/debug/mod.go:217
				// _ = "end of CoverTab[90804]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:218
				_go_fuzz_dep_.CoverTab[90805]++
//line /usr/local/go/src/runtime/debug/mod.go:218
				// _ = "end of CoverTab[90805]"
//line /usr/local/go/src/runtime/debug/mod.go:218
			}
//line /usr/local/go/src/runtime/debug/mod.go:218
			// _ = "end of CoverTab[90792]"
//line /usr/local/go/src/runtime/debug/mod.go:218
			_go_fuzz_dep_.CoverTab[90793]++
									if last == nil {
//line /usr/local/go/src/runtime/debug/mod.go:219
				_go_fuzz_dep_.CoverTab[90806]++
										return nil, fmt.Errorf("replacement with no module on previous line")
//line /usr/local/go/src/runtime/debug/mod.go:220
				// _ = "end of CoverTab[90806]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:221
				_go_fuzz_dep_.CoverTab[90807]++
//line /usr/local/go/src/runtime/debug/mod.go:221
				// _ = "end of CoverTab[90807]"
//line /usr/local/go/src/runtime/debug/mod.go:221
			}
//line /usr/local/go/src/runtime/debug/mod.go:221
			// _ = "end of CoverTab[90793]"
//line /usr/local/go/src/runtime/debug/mod.go:221
			_go_fuzz_dep_.CoverTab[90794]++
									last.Replace = &Module{
				Path:		string(elem[0]),
				Version:	string(elem[1]),
				Sum:		string(elem[2]),
			}
									last = nil
//line /usr/local/go/src/runtime/debug/mod.go:227
			// _ = "end of CoverTab[90794]"
		case strings.HasPrefix(line, buildLine):
//line /usr/local/go/src/runtime/debug/mod.go:228
			_go_fuzz_dep_.CoverTab[90795]++
									kv := line[len(buildLine):]
									if len(kv) < 1 {
//line /usr/local/go/src/runtime/debug/mod.go:230
				_go_fuzz_dep_.CoverTab[90808]++
										return nil, fmt.Errorf("build line missing '='")
//line /usr/local/go/src/runtime/debug/mod.go:231
				// _ = "end of CoverTab[90808]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:232
				_go_fuzz_dep_.CoverTab[90809]++
//line /usr/local/go/src/runtime/debug/mod.go:232
				// _ = "end of CoverTab[90809]"
//line /usr/local/go/src/runtime/debug/mod.go:232
			}
//line /usr/local/go/src/runtime/debug/mod.go:232
			// _ = "end of CoverTab[90795]"
//line /usr/local/go/src/runtime/debug/mod.go:232
			_go_fuzz_dep_.CoverTab[90796]++

									var key, rawValue string
									switch kv[0] {
			case '=':
//line /usr/local/go/src/runtime/debug/mod.go:236
				_go_fuzz_dep_.CoverTab[90810]++
										return nil, fmt.Errorf("build line with missing key")
//line /usr/local/go/src/runtime/debug/mod.go:237
				// _ = "end of CoverTab[90810]"

			case '`', '"':
//line /usr/local/go/src/runtime/debug/mod.go:239
				_go_fuzz_dep_.CoverTab[90811]++
										rawKey, err := strconv.QuotedPrefix(kv)
										if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:241
					_go_fuzz_dep_.CoverTab[90817]++
											return nil, fmt.Errorf("invalid quoted key in build line")
//line /usr/local/go/src/runtime/debug/mod.go:242
					// _ = "end of CoverTab[90817]"
				} else {
//line /usr/local/go/src/runtime/debug/mod.go:243
					_go_fuzz_dep_.CoverTab[90818]++
//line /usr/local/go/src/runtime/debug/mod.go:243
					// _ = "end of CoverTab[90818]"
//line /usr/local/go/src/runtime/debug/mod.go:243
				}
//line /usr/local/go/src/runtime/debug/mod.go:243
				// _ = "end of CoverTab[90811]"
//line /usr/local/go/src/runtime/debug/mod.go:243
				_go_fuzz_dep_.CoverTab[90812]++
										if len(kv) == len(rawKey) {
//line /usr/local/go/src/runtime/debug/mod.go:244
					_go_fuzz_dep_.CoverTab[90819]++
											return nil, fmt.Errorf("build line missing '=' after quoted key")
//line /usr/local/go/src/runtime/debug/mod.go:245
					// _ = "end of CoverTab[90819]"
				} else {
//line /usr/local/go/src/runtime/debug/mod.go:246
					_go_fuzz_dep_.CoverTab[90820]++
//line /usr/local/go/src/runtime/debug/mod.go:246
					// _ = "end of CoverTab[90820]"
//line /usr/local/go/src/runtime/debug/mod.go:246
				}
//line /usr/local/go/src/runtime/debug/mod.go:246
				// _ = "end of CoverTab[90812]"
//line /usr/local/go/src/runtime/debug/mod.go:246
				_go_fuzz_dep_.CoverTab[90813]++
										if c := kv[len(rawKey)]; c != '=' {
//line /usr/local/go/src/runtime/debug/mod.go:247
					_go_fuzz_dep_.CoverTab[90821]++
											return nil, fmt.Errorf("unexpected character after quoted key: %q", c)
//line /usr/local/go/src/runtime/debug/mod.go:248
					// _ = "end of CoverTab[90821]"
				} else {
//line /usr/local/go/src/runtime/debug/mod.go:249
					_go_fuzz_dep_.CoverTab[90822]++
//line /usr/local/go/src/runtime/debug/mod.go:249
					// _ = "end of CoverTab[90822]"
//line /usr/local/go/src/runtime/debug/mod.go:249
				}
//line /usr/local/go/src/runtime/debug/mod.go:249
				// _ = "end of CoverTab[90813]"
//line /usr/local/go/src/runtime/debug/mod.go:249
				_go_fuzz_dep_.CoverTab[90814]++
										key, _ = strconv.Unquote(rawKey)
										rawValue = kv[len(rawKey)+1:]
//line /usr/local/go/src/runtime/debug/mod.go:251
				// _ = "end of CoverTab[90814]"

			default:
//line /usr/local/go/src/runtime/debug/mod.go:253
				_go_fuzz_dep_.CoverTab[90815]++
										var ok bool
										key, rawValue, ok = strings.Cut(kv, "=")
										if !ok {
//line /usr/local/go/src/runtime/debug/mod.go:256
					_go_fuzz_dep_.CoverTab[90823]++
											return nil, fmt.Errorf("build line missing '=' after key")
//line /usr/local/go/src/runtime/debug/mod.go:257
					// _ = "end of CoverTab[90823]"
				} else {
//line /usr/local/go/src/runtime/debug/mod.go:258
					_go_fuzz_dep_.CoverTab[90824]++
//line /usr/local/go/src/runtime/debug/mod.go:258
					// _ = "end of CoverTab[90824]"
//line /usr/local/go/src/runtime/debug/mod.go:258
				}
//line /usr/local/go/src/runtime/debug/mod.go:258
				// _ = "end of CoverTab[90815]"
//line /usr/local/go/src/runtime/debug/mod.go:258
				_go_fuzz_dep_.CoverTab[90816]++
										if quoteKey(key) {
//line /usr/local/go/src/runtime/debug/mod.go:259
					_go_fuzz_dep_.CoverTab[90825]++
											return nil, fmt.Errorf("unquoted key %q must be quoted", key)
//line /usr/local/go/src/runtime/debug/mod.go:260
					// _ = "end of CoverTab[90825]"
				} else {
//line /usr/local/go/src/runtime/debug/mod.go:261
					_go_fuzz_dep_.CoverTab[90826]++
//line /usr/local/go/src/runtime/debug/mod.go:261
					// _ = "end of CoverTab[90826]"
//line /usr/local/go/src/runtime/debug/mod.go:261
				}
//line /usr/local/go/src/runtime/debug/mod.go:261
				// _ = "end of CoverTab[90816]"
			}
//line /usr/local/go/src/runtime/debug/mod.go:262
			// _ = "end of CoverTab[90796]"
//line /usr/local/go/src/runtime/debug/mod.go:262
			_go_fuzz_dep_.CoverTab[90797]++

									var value string
									if len(rawValue) > 0 {
//line /usr/local/go/src/runtime/debug/mod.go:265
				_go_fuzz_dep_.CoverTab[90827]++
										switch rawValue[0] {
				case '`', '"':
//line /usr/local/go/src/runtime/debug/mod.go:267
					_go_fuzz_dep_.CoverTab[90828]++
											var err error
											value, err = strconv.Unquote(rawValue)
											if err != nil {
//line /usr/local/go/src/runtime/debug/mod.go:270
						_go_fuzz_dep_.CoverTab[90830]++
												return nil, fmt.Errorf("invalid quoted value in build line")
//line /usr/local/go/src/runtime/debug/mod.go:271
						// _ = "end of CoverTab[90830]"
					} else {
//line /usr/local/go/src/runtime/debug/mod.go:272
						_go_fuzz_dep_.CoverTab[90831]++
//line /usr/local/go/src/runtime/debug/mod.go:272
						// _ = "end of CoverTab[90831]"
//line /usr/local/go/src/runtime/debug/mod.go:272
					}
//line /usr/local/go/src/runtime/debug/mod.go:272
					// _ = "end of CoverTab[90828]"

				default:
//line /usr/local/go/src/runtime/debug/mod.go:274
					_go_fuzz_dep_.CoverTab[90829]++
											value = rawValue
											if quoteValue(value) {
//line /usr/local/go/src/runtime/debug/mod.go:276
						_go_fuzz_dep_.CoverTab[90832]++
												return nil, fmt.Errorf("unquoted value %q must be quoted", value)
//line /usr/local/go/src/runtime/debug/mod.go:277
						// _ = "end of CoverTab[90832]"
					} else {
//line /usr/local/go/src/runtime/debug/mod.go:278
						_go_fuzz_dep_.CoverTab[90833]++
//line /usr/local/go/src/runtime/debug/mod.go:278
						// _ = "end of CoverTab[90833]"
//line /usr/local/go/src/runtime/debug/mod.go:278
					}
//line /usr/local/go/src/runtime/debug/mod.go:278
					// _ = "end of CoverTab[90829]"
				}
//line /usr/local/go/src/runtime/debug/mod.go:279
				// _ = "end of CoverTab[90827]"
			} else {
//line /usr/local/go/src/runtime/debug/mod.go:280
				_go_fuzz_dep_.CoverTab[90834]++
//line /usr/local/go/src/runtime/debug/mod.go:280
				// _ = "end of CoverTab[90834]"
//line /usr/local/go/src/runtime/debug/mod.go:280
			}
//line /usr/local/go/src/runtime/debug/mod.go:280
			// _ = "end of CoverTab[90797]"
//line /usr/local/go/src/runtime/debug/mod.go:280
			_go_fuzz_dep_.CoverTab[90798]++

									bi.Settings = append(bi.Settings, BuildSetting{Key: key, Value: value})
//line /usr/local/go/src/runtime/debug/mod.go:282
			// _ = "end of CoverTab[90798]"
//line /usr/local/go/src/runtime/debug/mod.go:282
		default:
//line /usr/local/go/src/runtime/debug/mod.go:282
			_go_fuzz_dep_.CoverTab[90799]++
//line /usr/local/go/src/runtime/debug/mod.go:282
			// _ = "end of CoverTab[90799]"
		}
//line /usr/local/go/src/runtime/debug/mod.go:283
		// _ = "end of CoverTab[90785]"
//line /usr/local/go/src/runtime/debug/mod.go:283
		_go_fuzz_dep_.CoverTab[90786]++
								lineNum++
//line /usr/local/go/src/runtime/debug/mod.go:284
		// _ = "end of CoverTab[90786]"
	}
//line /usr/local/go/src/runtime/debug/mod.go:285
	// _ = "end of CoverTab[90771]"
//line /usr/local/go/src/runtime/debug/mod.go:285
	_go_fuzz_dep_.CoverTab[90772]++
							return bi, nil
//line /usr/local/go/src/runtime/debug/mod.go:286
	// _ = "end of CoverTab[90772]"
}

//line /usr/local/go/src/runtime/debug/mod.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/runtime/debug/mod.go:287
var _ = _go_fuzz_dep_.CoverTab
