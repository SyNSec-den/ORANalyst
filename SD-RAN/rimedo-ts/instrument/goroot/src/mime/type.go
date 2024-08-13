// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/type.go:5
// Package mime implements parts of the MIME spec.
package mime

//line /usr/local/go/src/mime/type.go:6
import (
//line /usr/local/go/src/mime/type.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/type.go:6
)
//line /usr/local/go/src/mime/type.go:6
import (
//line /usr/local/go/src/mime/type.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/type.go:6
)

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var (
	mimeTypes	sync.Map	// map[string]string; ".Z" => "application/x-compress"
	mimeTypesLower	sync.Map	// map[string]string; ".z" => "application/x-compress"

	// extensions maps from MIME type to list of lowercase file
	// extensions: "image/jpeg" => [".jpg", ".jpeg"]
	extensionsMu	sync.Mutex	// Guards stores (but not loads) on extensions.
	extensions	sync.Map	// map[string][]string; slice values are append-only.
)

func clearSyncMap(m *sync.Map) {
//line /usr/local/go/src/mime/type.go:25
	_go_fuzz_dep_.CoverTab[35810]++
						m.Range(func(k, _ any) bool {
//line /usr/local/go/src/mime/type.go:26
		_go_fuzz_dep_.CoverTab[35811]++
							m.Delete(k)
							return true
//line /usr/local/go/src/mime/type.go:28
		// _ = "end of CoverTab[35811]"
	})
//line /usr/local/go/src/mime/type.go:29
	// _ = "end of CoverTab[35810]"
}

// setMimeTypes is used by initMime's non-test path, and by tests.
func setMimeTypes(lowerExt, mixExt map[string]string) {
//line /usr/local/go/src/mime/type.go:33
	_go_fuzz_dep_.CoverTab[35812]++
						clearSyncMap(&mimeTypes)
						clearSyncMap(&mimeTypesLower)
						clearSyncMap(&extensions)

						for k, v := range lowerExt {
//line /usr/local/go/src/mime/type.go:38
		_go_fuzz_dep_.CoverTab[35815]++
							mimeTypesLower.Store(k, v)
//line /usr/local/go/src/mime/type.go:39
		// _ = "end of CoverTab[35815]"
	}
//line /usr/local/go/src/mime/type.go:40
	// _ = "end of CoverTab[35812]"
//line /usr/local/go/src/mime/type.go:40
	_go_fuzz_dep_.CoverTab[35813]++
						for k, v := range mixExt {
//line /usr/local/go/src/mime/type.go:41
		_go_fuzz_dep_.CoverTab[35816]++
							mimeTypes.Store(k, v)
//line /usr/local/go/src/mime/type.go:42
		// _ = "end of CoverTab[35816]"
	}
//line /usr/local/go/src/mime/type.go:43
	// _ = "end of CoverTab[35813]"
//line /usr/local/go/src/mime/type.go:43
	_go_fuzz_dep_.CoverTab[35814]++

						extensionsMu.Lock()
						defer extensionsMu.Unlock()
						for k, v := range lowerExt {
//line /usr/local/go/src/mime/type.go:47
		_go_fuzz_dep_.CoverTab[35817]++
							justType, _, err := ParseMediaType(v)
							if err != nil {
//line /usr/local/go/src/mime/type.go:49
			_go_fuzz_dep_.CoverTab[35820]++
								panic(err)
//line /usr/local/go/src/mime/type.go:50
			// _ = "end of CoverTab[35820]"
		} else {
//line /usr/local/go/src/mime/type.go:51
			_go_fuzz_dep_.CoverTab[35821]++
//line /usr/local/go/src/mime/type.go:51
			// _ = "end of CoverTab[35821]"
//line /usr/local/go/src/mime/type.go:51
		}
//line /usr/local/go/src/mime/type.go:51
		// _ = "end of CoverTab[35817]"
//line /usr/local/go/src/mime/type.go:51
		_go_fuzz_dep_.CoverTab[35818]++
							var exts []string
							if ei, ok := extensions.Load(justType); ok {
//line /usr/local/go/src/mime/type.go:53
			_go_fuzz_dep_.CoverTab[35822]++
								exts = ei.([]string)
//line /usr/local/go/src/mime/type.go:54
			// _ = "end of CoverTab[35822]"
		} else {
//line /usr/local/go/src/mime/type.go:55
			_go_fuzz_dep_.CoverTab[35823]++
//line /usr/local/go/src/mime/type.go:55
			// _ = "end of CoverTab[35823]"
//line /usr/local/go/src/mime/type.go:55
		}
//line /usr/local/go/src/mime/type.go:55
		// _ = "end of CoverTab[35818]"
//line /usr/local/go/src/mime/type.go:55
		_go_fuzz_dep_.CoverTab[35819]++
							extensions.Store(justType, append(exts, k))
//line /usr/local/go/src/mime/type.go:56
		// _ = "end of CoverTab[35819]"
	}
//line /usr/local/go/src/mime/type.go:57
	// _ = "end of CoverTab[35814]"
}

var builtinTypesLower = map[string]string{
	".avif":	"image/avif",
	".css":		"text/css; charset=utf-8",
	".gif":		"image/gif",
	".htm":		"text/html; charset=utf-8",
	".html":	"text/html; charset=utf-8",
	".jpeg":	"image/jpeg",
	".jpg":		"image/jpeg",
	".js":		"text/javascript; charset=utf-8",
	".json":	"application/json",
	".mjs":		"text/javascript; charset=utf-8",
	".pdf":		"application/pdf",
	".png":		"image/png",
	".svg":		"image/svg+xml",
	".wasm":	"application/wasm",
	".webp":	"image/webp",
	".xml":		"text/xml; charset=utf-8",
}

var once sync.Once	// guards initMime

var testInitMime, osInitMime func()

func initMime() {
//line /usr/local/go/src/mime/type.go:83
	_go_fuzz_dep_.CoverTab[35824]++
						if fn := testInitMime; fn != nil {
//line /usr/local/go/src/mime/type.go:84
		_go_fuzz_dep_.CoverTab[35825]++
							fn()
//line /usr/local/go/src/mime/type.go:85
		// _ = "end of CoverTab[35825]"
	} else {
//line /usr/local/go/src/mime/type.go:86
		_go_fuzz_dep_.CoverTab[35826]++
							setMimeTypes(builtinTypesLower, builtinTypesLower)
							osInitMime()
//line /usr/local/go/src/mime/type.go:88
		// _ = "end of CoverTab[35826]"
	}
//line /usr/local/go/src/mime/type.go:89
	// _ = "end of CoverTab[35824]"
}

// TypeByExtension returns the MIME type associated with the file extension ext.
//line /usr/local/go/src/mime/type.go:92
// The extension ext should begin with a leading dot, as in ".html".
//line /usr/local/go/src/mime/type.go:92
// When ext has no associated type, TypeByExtension returns "".
//line /usr/local/go/src/mime/type.go:92
//
//line /usr/local/go/src/mime/type.go:92
// Extensions are looked up first case-sensitively, then case-insensitively.
//line /usr/local/go/src/mime/type.go:92
//
//line /usr/local/go/src/mime/type.go:92
// The built-in table is small but on unix it is augmented by the local
//line /usr/local/go/src/mime/type.go:92
// system's MIME-info database or mime.types file(s) if available under one or
//line /usr/local/go/src/mime/type.go:92
// more of these names:
//line /usr/local/go/src/mime/type.go:92
//
//line /usr/local/go/src/mime/type.go:92
//	/usr/local/share/mime/globs2
//line /usr/local/go/src/mime/type.go:92
//	/usr/share/mime/globs2
//line /usr/local/go/src/mime/type.go:92
//	/etc/mime.types
//line /usr/local/go/src/mime/type.go:92
//	/etc/apache2/mime.types
//line /usr/local/go/src/mime/type.go:92
//	/etc/apache/mime.types
//line /usr/local/go/src/mime/type.go:92
//
//line /usr/local/go/src/mime/type.go:92
// On Windows, MIME types are extracted from the registry.
//line /usr/local/go/src/mime/type.go:92
//
//line /usr/local/go/src/mime/type.go:92
// Text types have the charset parameter set to "utf-8" by default.
//line /usr/local/go/src/mime/type.go:111
func TypeByExtension(ext string) string {
//line /usr/local/go/src/mime/type.go:111
	_go_fuzz_dep_.CoverTab[35827]++
						once.Do(initMime)

//line /usr/local/go/src/mime/type.go:115
	if v, ok := mimeTypes.Load(ext); ok {
//line /usr/local/go/src/mime/type.go:115
		_go_fuzz_dep_.CoverTab[35830]++
							return v.(string)
//line /usr/local/go/src/mime/type.go:116
		// _ = "end of CoverTab[35830]"
	} else {
//line /usr/local/go/src/mime/type.go:117
		_go_fuzz_dep_.CoverTab[35831]++
//line /usr/local/go/src/mime/type.go:117
		// _ = "end of CoverTab[35831]"
//line /usr/local/go/src/mime/type.go:117
	}
//line /usr/local/go/src/mime/type.go:117
	// _ = "end of CoverTab[35827]"
//line /usr/local/go/src/mime/type.go:117
	_go_fuzz_dep_.CoverTab[35828]++

	// Case-insensitive lookup.
	// Optimistically assume a short ASCII extension and be
	// allocation-free in that case.
	var buf [10]byte
	lower := buf[:0]
	const utf8RuneSelf = 0x80	// from utf8 package, but not importing it.
	for i := 0; i < len(ext); i++ {
//line /usr/local/go/src/mime/type.go:125
		_go_fuzz_dep_.CoverTab[35832]++
							c := ext[i]
							if c >= utf8RuneSelf {
//line /usr/local/go/src/mime/type.go:127
			_go_fuzz_dep_.CoverTab[35834]++

								si, _ := mimeTypesLower.Load(strings.ToLower(ext))
								s, _ := si.(string)
								return s
//line /usr/local/go/src/mime/type.go:131
			// _ = "end of CoverTab[35834]"
		} else {
//line /usr/local/go/src/mime/type.go:132
			_go_fuzz_dep_.CoverTab[35835]++
//line /usr/local/go/src/mime/type.go:132
			// _ = "end of CoverTab[35835]"
//line /usr/local/go/src/mime/type.go:132
		}
//line /usr/local/go/src/mime/type.go:132
		// _ = "end of CoverTab[35832]"
//line /usr/local/go/src/mime/type.go:132
		_go_fuzz_dep_.CoverTab[35833]++
							if 'A' <= c && func() bool {
//line /usr/local/go/src/mime/type.go:133
			_go_fuzz_dep_.CoverTab[35836]++
//line /usr/local/go/src/mime/type.go:133
			return c <= 'Z'
//line /usr/local/go/src/mime/type.go:133
			// _ = "end of CoverTab[35836]"
//line /usr/local/go/src/mime/type.go:133
		}() {
//line /usr/local/go/src/mime/type.go:133
			_go_fuzz_dep_.CoverTab[35837]++
								lower = append(lower, c+('a'-'A'))
//line /usr/local/go/src/mime/type.go:134
			// _ = "end of CoverTab[35837]"
		} else {
//line /usr/local/go/src/mime/type.go:135
			_go_fuzz_dep_.CoverTab[35838]++
								lower = append(lower, c)
//line /usr/local/go/src/mime/type.go:136
			// _ = "end of CoverTab[35838]"
		}
//line /usr/local/go/src/mime/type.go:137
		// _ = "end of CoverTab[35833]"
	}
//line /usr/local/go/src/mime/type.go:138
	// _ = "end of CoverTab[35828]"
//line /usr/local/go/src/mime/type.go:138
	_go_fuzz_dep_.CoverTab[35829]++
						si, _ := mimeTypesLower.Load(string(lower))
						s, _ := si.(string)
						return s
//line /usr/local/go/src/mime/type.go:141
	// _ = "end of CoverTab[35829]"
}

// ExtensionsByType returns the extensions known to be associated with the MIME
//line /usr/local/go/src/mime/type.go:144
// type typ. The returned extensions will each begin with a leading dot, as in
//line /usr/local/go/src/mime/type.go:144
// ".html". When typ has no associated extensions, ExtensionsByType returns an
//line /usr/local/go/src/mime/type.go:144
// nil slice.
//line /usr/local/go/src/mime/type.go:148
func ExtensionsByType(typ string) ([]string, error) {
//line /usr/local/go/src/mime/type.go:148
	_go_fuzz_dep_.CoverTab[35839]++
						justType, _, err := ParseMediaType(typ)
						if err != nil {
//line /usr/local/go/src/mime/type.go:150
		_go_fuzz_dep_.CoverTab[35842]++
							return nil, err
//line /usr/local/go/src/mime/type.go:151
		// _ = "end of CoverTab[35842]"
	} else {
//line /usr/local/go/src/mime/type.go:152
		_go_fuzz_dep_.CoverTab[35843]++
//line /usr/local/go/src/mime/type.go:152
		// _ = "end of CoverTab[35843]"
//line /usr/local/go/src/mime/type.go:152
	}
//line /usr/local/go/src/mime/type.go:152
	// _ = "end of CoverTab[35839]"
//line /usr/local/go/src/mime/type.go:152
	_go_fuzz_dep_.CoverTab[35840]++

						once.Do(initMime)
						s, ok := extensions.Load(justType)
						if !ok {
//line /usr/local/go/src/mime/type.go:156
		_go_fuzz_dep_.CoverTab[35844]++
							return nil, nil
//line /usr/local/go/src/mime/type.go:157
		// _ = "end of CoverTab[35844]"
	} else {
//line /usr/local/go/src/mime/type.go:158
		_go_fuzz_dep_.CoverTab[35845]++
//line /usr/local/go/src/mime/type.go:158
		// _ = "end of CoverTab[35845]"
//line /usr/local/go/src/mime/type.go:158
	}
//line /usr/local/go/src/mime/type.go:158
	// _ = "end of CoverTab[35840]"
//line /usr/local/go/src/mime/type.go:158
	_go_fuzz_dep_.CoverTab[35841]++
						ret := append([]string(nil), s.([]string)...)
						sort.Strings(ret)
						return ret, nil
//line /usr/local/go/src/mime/type.go:161
	// _ = "end of CoverTab[35841]"
}

// AddExtensionType sets the MIME type associated with
//line /usr/local/go/src/mime/type.go:164
// the extension ext to typ. The extension should begin with
//line /usr/local/go/src/mime/type.go:164
// a leading dot, as in ".html".
//line /usr/local/go/src/mime/type.go:167
func AddExtensionType(ext, typ string) error {
//line /usr/local/go/src/mime/type.go:167
	_go_fuzz_dep_.CoverTab[35846]++
						if !strings.HasPrefix(ext, ".") {
//line /usr/local/go/src/mime/type.go:168
		_go_fuzz_dep_.CoverTab[35848]++
							return fmt.Errorf("mime: extension %q missing leading dot", ext)
//line /usr/local/go/src/mime/type.go:169
		// _ = "end of CoverTab[35848]"
	} else {
//line /usr/local/go/src/mime/type.go:170
		_go_fuzz_dep_.CoverTab[35849]++
//line /usr/local/go/src/mime/type.go:170
		// _ = "end of CoverTab[35849]"
//line /usr/local/go/src/mime/type.go:170
	}
//line /usr/local/go/src/mime/type.go:170
	// _ = "end of CoverTab[35846]"
//line /usr/local/go/src/mime/type.go:170
	_go_fuzz_dep_.CoverTab[35847]++
						once.Do(initMime)
						return setExtensionType(ext, typ)
//line /usr/local/go/src/mime/type.go:172
	// _ = "end of CoverTab[35847]"
}

func setExtensionType(extension, mimeType string) error {
//line /usr/local/go/src/mime/type.go:175
	_go_fuzz_dep_.CoverTab[35850]++
						justType, param, err := ParseMediaType(mimeType)
						if err != nil {
//line /usr/local/go/src/mime/type.go:177
		_go_fuzz_dep_.CoverTab[35855]++
							return err
//line /usr/local/go/src/mime/type.go:178
		// _ = "end of CoverTab[35855]"
	} else {
//line /usr/local/go/src/mime/type.go:179
		_go_fuzz_dep_.CoverTab[35856]++
//line /usr/local/go/src/mime/type.go:179
		// _ = "end of CoverTab[35856]"
//line /usr/local/go/src/mime/type.go:179
	}
//line /usr/local/go/src/mime/type.go:179
	// _ = "end of CoverTab[35850]"
//line /usr/local/go/src/mime/type.go:179
	_go_fuzz_dep_.CoverTab[35851]++
						if strings.HasPrefix(mimeType, "text/") && func() bool {
//line /usr/local/go/src/mime/type.go:180
		_go_fuzz_dep_.CoverTab[35857]++
//line /usr/local/go/src/mime/type.go:180
		return param["charset"] == ""
//line /usr/local/go/src/mime/type.go:180
		// _ = "end of CoverTab[35857]"
//line /usr/local/go/src/mime/type.go:180
	}() {
//line /usr/local/go/src/mime/type.go:180
		_go_fuzz_dep_.CoverTab[35858]++
							param["charset"] = "utf-8"
							mimeType = FormatMediaType(mimeType, param)
//line /usr/local/go/src/mime/type.go:182
		// _ = "end of CoverTab[35858]"
	} else {
//line /usr/local/go/src/mime/type.go:183
		_go_fuzz_dep_.CoverTab[35859]++
//line /usr/local/go/src/mime/type.go:183
		// _ = "end of CoverTab[35859]"
//line /usr/local/go/src/mime/type.go:183
	}
//line /usr/local/go/src/mime/type.go:183
	// _ = "end of CoverTab[35851]"
//line /usr/local/go/src/mime/type.go:183
	_go_fuzz_dep_.CoverTab[35852]++
						extLower := strings.ToLower(extension)

						mimeTypes.Store(extension, mimeType)
						mimeTypesLower.Store(extLower, mimeType)

						extensionsMu.Lock()
						defer extensionsMu.Unlock()
						var exts []string
						if ei, ok := extensions.Load(justType); ok {
//line /usr/local/go/src/mime/type.go:192
		_go_fuzz_dep_.CoverTab[35860]++
							exts = ei.([]string)
//line /usr/local/go/src/mime/type.go:193
		// _ = "end of CoverTab[35860]"
	} else {
//line /usr/local/go/src/mime/type.go:194
		_go_fuzz_dep_.CoverTab[35861]++
//line /usr/local/go/src/mime/type.go:194
		// _ = "end of CoverTab[35861]"
//line /usr/local/go/src/mime/type.go:194
	}
//line /usr/local/go/src/mime/type.go:194
	// _ = "end of CoverTab[35852]"
//line /usr/local/go/src/mime/type.go:194
	_go_fuzz_dep_.CoverTab[35853]++
						for _, v := range exts {
//line /usr/local/go/src/mime/type.go:195
		_go_fuzz_dep_.CoverTab[35862]++
							if v == extLower {
//line /usr/local/go/src/mime/type.go:196
			_go_fuzz_dep_.CoverTab[35863]++
								return nil
//line /usr/local/go/src/mime/type.go:197
			// _ = "end of CoverTab[35863]"
		} else {
//line /usr/local/go/src/mime/type.go:198
			_go_fuzz_dep_.CoverTab[35864]++
//line /usr/local/go/src/mime/type.go:198
			// _ = "end of CoverTab[35864]"
//line /usr/local/go/src/mime/type.go:198
		}
//line /usr/local/go/src/mime/type.go:198
		// _ = "end of CoverTab[35862]"
	}
//line /usr/local/go/src/mime/type.go:199
	// _ = "end of CoverTab[35853]"
//line /usr/local/go/src/mime/type.go:199
	_go_fuzz_dep_.CoverTab[35854]++
						extensions.Store(justType, append(exts, extLower))
						return nil
//line /usr/local/go/src/mime/type.go:201
	// _ = "end of CoverTab[35854]"
}

//line /usr/local/go/src/mime/type.go:202
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/type.go:202
var _ = _go_fuzz_dep_.CoverTab
