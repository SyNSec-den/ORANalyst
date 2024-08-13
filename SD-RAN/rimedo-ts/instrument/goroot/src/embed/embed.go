// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/embed/embed.go:5
// Package embed provides access to files embedded in the running Go program.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// Go source files that import "embed" can use the //go:embed directive
//line /usr/local/go/src/embed/embed.go:5
// to initialize a variable of type string, []byte, or FS with the contents of
//line /usr/local/go/src/embed/embed.go:5
// files read from the package directory or subdirectories at compile time.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// For example, here are three ways to embed a file named hello.txt
//line /usr/local/go/src/embed/embed.go:5
// and then print its contents at run time.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// Embedding one file into a string:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	import _ "embed"
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed hello.txt
//line /usr/local/go/src/embed/embed.go:5
//	var s string
//line /usr/local/go/src/embed/embed.go:5
//	print(s)
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// Embedding one file into a slice of bytes:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	import _ "embed"
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed hello.txt
//line /usr/local/go/src/embed/embed.go:5
//	var b []byte
//line /usr/local/go/src/embed/embed.go:5
//	print(string(b))
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// Embedded one or more files into a file system:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	import "embed"
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed hello.txt
//line /usr/local/go/src/embed/embed.go:5
//	var f embed.FS
//line /usr/local/go/src/embed/embed.go:5
//	data, _ := f.ReadFile("hello.txt")
//line /usr/local/go/src/embed/embed.go:5
//	print(string(data))
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// # Directives
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// A //go:embed directive above a variable declaration specifies which files to embed,
//line /usr/local/go/src/embed/embed.go:5
// using one or more path.Match patterns.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The directive must immediately precede a line containing the declaration of a single variable.
//line /usr/local/go/src/embed/embed.go:5
// Only blank lines and ‘//’ line comments are permitted between the directive and the declaration.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The type of the variable must be a string type, or a slice of a byte type,
//line /usr/local/go/src/embed/embed.go:5
// or FS (or an alias of FS).
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// For example:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	package server
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	import "embed"
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	// content holds our static web server content.
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed image/* template/*
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed html/index.html
//line /usr/local/go/src/embed/embed.go:5
//	var content embed.FS
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The Go build system will recognize the directives and arrange for the declared variable
//line /usr/local/go/src/embed/embed.go:5
// (in the example above, content) to be populated with the matching files from the file system.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The //go:embed directive accepts multiple space-separated patterns for
//line /usr/local/go/src/embed/embed.go:5
// brevity, but it can also be repeated, to avoid very long lines when there are
//line /usr/local/go/src/embed/embed.go:5
// many patterns. The patterns are interpreted relative to the package directory
//line /usr/local/go/src/embed/embed.go:5
// containing the source file. The path separator is a forward slash, even on
//line /usr/local/go/src/embed/embed.go:5
// Windows systems. Patterns may not contain ‘.’ or ‘..’ or empty path elements,
//line /usr/local/go/src/embed/embed.go:5
// nor may they begin or end with a slash. To match everything in the current
//line /usr/local/go/src/embed/embed.go:5
// directory, use ‘*’ instead of ‘.’. To allow for naming files with spaces in
//line /usr/local/go/src/embed/embed.go:5
// their names, patterns can be written as Go double-quoted or back-quoted
//line /usr/local/go/src/embed/embed.go:5
// string literals.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// If a pattern names a directory, all files in the subtree rooted at that directory are
//line /usr/local/go/src/embed/embed.go:5
// embedded (recursively), except that files with names beginning with ‘.’ or ‘_’
//line /usr/local/go/src/embed/embed.go:5
// are excluded. So the variable in the above example is almost equivalent to:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	// content is our static web server content.
//line /usr/local/go/src/embed/embed.go:5
//	//go:embed image template html/index.html
//line /usr/local/go/src/embed/embed.go:5
//	var content embed.FS
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The difference is that ‘image/*’ embeds ‘image/.tempfile’ while ‘image’ does not.
//line /usr/local/go/src/embed/embed.go:5
// Neither embeds ‘image/dir/.tempfile’.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// If a pattern begins with the prefix ‘all:’, then the rule for walking directories is changed
//line /usr/local/go/src/embed/embed.go:5
// to include those files beginning with ‘.’ or ‘_’. For example, ‘all:image’ embeds
//line /usr/local/go/src/embed/embed.go:5
// both ‘image/.tempfile’ and ‘image/dir/.tempfile’.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The //go:embed directive can be used with both exported and unexported variables,
//line /usr/local/go/src/embed/embed.go:5
// depending on whether the package wants to make the data available to other packages.
//line /usr/local/go/src/embed/embed.go:5
// It can only be used with variables at package scope, not with local variables.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// Patterns must not match files outside the package's module, such as ‘.git/*’ or symbolic links.
//line /usr/local/go/src/embed/embed.go:5
// Patterns must not match files whose names include the special punctuation characters  " * < > ? ` ' | / \ and :.
//line /usr/local/go/src/embed/embed.go:5
// Matches for empty directories are ignored. After that, each pattern in a //go:embed line
//line /usr/local/go/src/embed/embed.go:5
// must match at least one file or non-empty directory.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// If any patterns are invalid or have invalid matches, the build will fail.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// # Strings and Bytes
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The //go:embed line for a variable of type string or []byte can have only a single pattern,
//line /usr/local/go/src/embed/embed.go:5
// and that pattern can match only a single file. The string or []byte is initialized with
//line /usr/local/go/src/embed/embed.go:5
// the contents of that file.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// The //go:embed directive requires importing "embed", even when using a string or []byte.
//line /usr/local/go/src/embed/embed.go:5
// In source files that don't refer to embed.FS, use a blank import (import _ "embed").
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// # File Systems
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// For embedding a single file, a variable of type string or []byte is often best.
//line /usr/local/go/src/embed/embed.go:5
// The FS type enables embedding a tree of files, such as a directory of static
//line /usr/local/go/src/embed/embed.go:5
// web server content, as in the example above.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// FS implements the io/fs package's FS interface, so it can be used with any package that
//line /usr/local/go/src/embed/embed.go:5
// understands file systems, including net/http, text/template, and html/template.
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// For example, given the content variable in the example above, we can write:
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
//	template.ParseFS(content, "*.tmpl")
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// # Tools
//line /usr/local/go/src/embed/embed.go:5
//
//line /usr/local/go/src/embed/embed.go:5
// To support tools that analyze Go packages, the patterns found in //go:embed lines
//line /usr/local/go/src/embed/embed.go:5
// are available in “go list” output. See the EmbedPatterns, TestEmbedPatterns,
//line /usr/local/go/src/embed/embed.go:5
// and XTestEmbedPatterns fields in the “go help list” output.
//line /usr/local/go/src/embed/embed.go:129
package embed

//line /usr/local/go/src/embed/embed.go:129
import (
//line /usr/local/go/src/embed/embed.go:129
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/embed/embed.go:129
)
//line /usr/local/go/src/embed/embed.go:129
import (
//line /usr/local/go/src/embed/embed.go:129
	_atomic_ "sync/atomic"
//line /usr/local/go/src/embed/embed.go:129
)

import (
	"errors"
	"io"
	"io/fs"
	"time"
)

// An FS is a read-only collection of files, usually initialized with a //go:embed directive.
//line /usr/local/go/src/embed/embed.go:138
// When declared without a //go:embed directive, an FS is an empty file system.
//line /usr/local/go/src/embed/embed.go:138
//
//line /usr/local/go/src/embed/embed.go:138
// An FS is a read-only value, so it is safe to use from multiple goroutines
//line /usr/local/go/src/embed/embed.go:138
// simultaneously and also safe to assign values of type FS to each other.
//line /usr/local/go/src/embed/embed.go:138
//
//line /usr/local/go/src/embed/embed.go:138
// FS implements fs.FS, so it can be used with any package that understands
//line /usr/local/go/src/embed/embed.go:138
// file system interfaces, including net/http, text/template, and html/template.
//line /usr/local/go/src/embed/embed.go:138
//
//line /usr/local/go/src/embed/embed.go:138
// See the package documentation for more details about initializing an FS.
//line /usr/local/go/src/embed/embed.go:148
type FS struct {
	// The compiler knows the layout of this struct.
	// See cmd/compile/internal/staticdata's WriteEmbed.
	//
	// The files list is sorted by name but not by simple string comparison.
	// Instead, each file's name takes the form "dir/elem" or "dir/elem/".
	// The optional trailing slash indicates that the file is itself a directory.
	// The files list is sorted first by dir (if dir is missing, it is taken to be ".")
	// and then by base, so this list of files:
	//
	//	p
	//	q/
	//	q/r
	//	q/s/
	//	q/s/t
	//	q/s/u
	//	q/v
	//	w
	//
	// is actually sorted as:
	//
	//	p       # dir=.    elem=p
	//	q/      # dir=.    elem=q
	//	w/      # dir=.    elem=w
	//	q/r     # dir=q    elem=r
	//	q/s/    # dir=q    elem=s
	//	q/v     # dir=q    elem=v
	//	q/s/t   # dir=q/s  elem=t
	//	q/s/u   # dir=q/s  elem=u
	//
	// This order brings directory contents together in contiguous sections
	// of the list, allowing a directory read to use binary search to find
	// the relevant sequence of entries.
	files *[]file
}

// split splits the name into dir and elem as described in the
//line /usr/local/go/src/embed/embed.go:184
// comment in the FS struct above. isDir reports whether the
//line /usr/local/go/src/embed/embed.go:184
// final trailing slash was present, indicating that name is a directory.
//line /usr/local/go/src/embed/embed.go:187
func split(name string) (dir, elem string, isDir bool) {
//line /usr/local/go/src/embed/embed.go:187
	_go_fuzz_dep_.CoverTab[2273]++
						if name[len(name)-1] == '/' {
//line /usr/local/go/src/embed/embed.go:188
		_go_fuzz_dep_.CoverTab[2277]++
							isDir = true
							name = name[:len(name)-1]
//line /usr/local/go/src/embed/embed.go:190
		// _ = "end of CoverTab[2277]"
	} else {
//line /usr/local/go/src/embed/embed.go:191
		_go_fuzz_dep_.CoverTab[2278]++
//line /usr/local/go/src/embed/embed.go:191
		// _ = "end of CoverTab[2278]"
//line /usr/local/go/src/embed/embed.go:191
	}
//line /usr/local/go/src/embed/embed.go:191
	// _ = "end of CoverTab[2273]"
//line /usr/local/go/src/embed/embed.go:191
	_go_fuzz_dep_.CoverTab[2274]++
						i := len(name) - 1
						for i >= 0 && func() bool {
//line /usr/local/go/src/embed/embed.go:193
		_go_fuzz_dep_.CoverTab[2279]++
//line /usr/local/go/src/embed/embed.go:193
		return name[i] != '/'
//line /usr/local/go/src/embed/embed.go:193
		// _ = "end of CoverTab[2279]"
//line /usr/local/go/src/embed/embed.go:193
	}() {
//line /usr/local/go/src/embed/embed.go:193
		_go_fuzz_dep_.CoverTab[2280]++
							i--
//line /usr/local/go/src/embed/embed.go:194
		// _ = "end of CoverTab[2280]"
	}
//line /usr/local/go/src/embed/embed.go:195
	// _ = "end of CoverTab[2274]"
//line /usr/local/go/src/embed/embed.go:195
	_go_fuzz_dep_.CoverTab[2275]++
						if i < 0 {
//line /usr/local/go/src/embed/embed.go:196
		_go_fuzz_dep_.CoverTab[2281]++
							return ".", name, isDir
//line /usr/local/go/src/embed/embed.go:197
		// _ = "end of CoverTab[2281]"
	} else {
//line /usr/local/go/src/embed/embed.go:198
		_go_fuzz_dep_.CoverTab[2282]++
//line /usr/local/go/src/embed/embed.go:198
		// _ = "end of CoverTab[2282]"
//line /usr/local/go/src/embed/embed.go:198
	}
//line /usr/local/go/src/embed/embed.go:198
	// _ = "end of CoverTab[2275]"
//line /usr/local/go/src/embed/embed.go:198
	_go_fuzz_dep_.CoverTab[2276]++
						return name[:i], name[i+1:], isDir
//line /usr/local/go/src/embed/embed.go:199
	// _ = "end of CoverTab[2276]"
}

// trimSlash trims a trailing slash from name, if present,
//line /usr/local/go/src/embed/embed.go:202
// returning the possibly shortened name.
//line /usr/local/go/src/embed/embed.go:204
func trimSlash(name string) string {
//line /usr/local/go/src/embed/embed.go:204
	_go_fuzz_dep_.CoverTab[2283]++
						if len(name) > 0 && func() bool {
//line /usr/local/go/src/embed/embed.go:205
		_go_fuzz_dep_.CoverTab[2285]++
//line /usr/local/go/src/embed/embed.go:205
		return name[len(name)-1] == '/'
//line /usr/local/go/src/embed/embed.go:205
		// _ = "end of CoverTab[2285]"
//line /usr/local/go/src/embed/embed.go:205
	}() {
//line /usr/local/go/src/embed/embed.go:205
		_go_fuzz_dep_.CoverTab[2286]++
							return name[:len(name)-1]
//line /usr/local/go/src/embed/embed.go:206
		// _ = "end of CoverTab[2286]"
	} else {
//line /usr/local/go/src/embed/embed.go:207
		_go_fuzz_dep_.CoverTab[2287]++
//line /usr/local/go/src/embed/embed.go:207
		// _ = "end of CoverTab[2287]"
//line /usr/local/go/src/embed/embed.go:207
	}
//line /usr/local/go/src/embed/embed.go:207
	// _ = "end of CoverTab[2283]"
//line /usr/local/go/src/embed/embed.go:207
	_go_fuzz_dep_.CoverTab[2284]++
						return name
//line /usr/local/go/src/embed/embed.go:208
	// _ = "end of CoverTab[2284]"
}

var (
	_	fs.ReadDirFS	= FS{}
	_	fs.ReadFileFS	= FS{}
)

// A file is a single file in the FS.
//line /usr/local/go/src/embed/embed.go:216
// It implements fs.FileInfo and fs.DirEntry.
//line /usr/local/go/src/embed/embed.go:218
type file struct {
	// The compiler knows the layout of this struct.
	// See cmd/compile/internal/staticdata's WriteEmbed.
	name	string
	data	string
	hash	[16]byte	// truncated SHA256 hash
}

var (
	_	fs.FileInfo	= (*file)(nil)
	_	fs.DirEntry	= (*file)(nil)
)

func (f *file) Name() string {
//line /usr/local/go/src/embed/embed.go:231
	_go_fuzz_dep_.CoverTab[2288]++
//line /usr/local/go/src/embed/embed.go:231
	_, elem, _ := split(f.name)
//line /usr/local/go/src/embed/embed.go:231
	return elem
//line /usr/local/go/src/embed/embed.go:231
	// _ = "end of CoverTab[2288]"
//line /usr/local/go/src/embed/embed.go:231
}
func (f *file) Size() int64 {
//line /usr/local/go/src/embed/embed.go:232
	_go_fuzz_dep_.CoverTab[2289]++
//line /usr/local/go/src/embed/embed.go:232
	return int64(len(f.data))
//line /usr/local/go/src/embed/embed.go:232
	// _ = "end of CoverTab[2289]"
//line /usr/local/go/src/embed/embed.go:232
}
func (f *file) ModTime() time.Time {
//line /usr/local/go/src/embed/embed.go:233
	_go_fuzz_dep_.CoverTab[2290]++
//line /usr/local/go/src/embed/embed.go:233
	return time.Time{}
//line /usr/local/go/src/embed/embed.go:233
	// _ = "end of CoverTab[2290]"
//line /usr/local/go/src/embed/embed.go:233
}
func (f *file) IsDir() bool {
//line /usr/local/go/src/embed/embed.go:234
	_go_fuzz_dep_.CoverTab[2291]++
//line /usr/local/go/src/embed/embed.go:234
	_, _, isDir := split(f.name)
//line /usr/local/go/src/embed/embed.go:234
	return isDir
//line /usr/local/go/src/embed/embed.go:234
	// _ = "end of CoverTab[2291]"
//line /usr/local/go/src/embed/embed.go:234
}
func (f *file) Sys() any	{ _go_fuzz_dep_.CoverTab[2292]++; return nil; // _ = "end of CoverTab[2292]" }
func (f *file) Type() fs.FileMode {
//line /usr/local/go/src/embed/embed.go:236
	_go_fuzz_dep_.CoverTab[2293]++
//line /usr/local/go/src/embed/embed.go:236
	return f.Mode().Type()
//line /usr/local/go/src/embed/embed.go:236
	// _ = "end of CoverTab[2293]"
//line /usr/local/go/src/embed/embed.go:236
}
func (f *file) Info() (fs.FileInfo, error) {
//line /usr/local/go/src/embed/embed.go:237
	_go_fuzz_dep_.CoverTab[2294]++
//line /usr/local/go/src/embed/embed.go:237
	return f, nil
//line /usr/local/go/src/embed/embed.go:237
	// _ = "end of CoverTab[2294]"
//line /usr/local/go/src/embed/embed.go:237
}

func (f *file) Mode() fs.FileMode {
//line /usr/local/go/src/embed/embed.go:239
	_go_fuzz_dep_.CoverTab[2295]++
						if f.IsDir() {
//line /usr/local/go/src/embed/embed.go:240
		_go_fuzz_dep_.CoverTab[2297]++
							return fs.ModeDir | 0555
//line /usr/local/go/src/embed/embed.go:241
		// _ = "end of CoverTab[2297]"
	} else {
//line /usr/local/go/src/embed/embed.go:242
		_go_fuzz_dep_.CoverTab[2298]++
//line /usr/local/go/src/embed/embed.go:242
		// _ = "end of CoverTab[2298]"
//line /usr/local/go/src/embed/embed.go:242
	}
//line /usr/local/go/src/embed/embed.go:242
	// _ = "end of CoverTab[2295]"
//line /usr/local/go/src/embed/embed.go:242
	_go_fuzz_dep_.CoverTab[2296]++
						return 0444
//line /usr/local/go/src/embed/embed.go:243
	// _ = "end of CoverTab[2296]"
}

// dotFile is a file for the root directory,
//line /usr/local/go/src/embed/embed.go:246
// which is omitted from the files list in a FS.
//line /usr/local/go/src/embed/embed.go:248
var dotFile = &file{name: "./"}

// lookup returns the named file, or nil if it is not present.
func (f FS) lookup(name string) *file {
//line /usr/local/go/src/embed/embed.go:251
	_go_fuzz_dep_.CoverTab[2299]++
						if !fs.ValidPath(name) {
//line /usr/local/go/src/embed/embed.go:252
		_go_fuzz_dep_.CoverTab[2305]++

//line /usr/local/go/src/embed/embed.go:256
		return nil
//line /usr/local/go/src/embed/embed.go:256
		// _ = "end of CoverTab[2305]"
	} else {
//line /usr/local/go/src/embed/embed.go:257
		_go_fuzz_dep_.CoverTab[2306]++
//line /usr/local/go/src/embed/embed.go:257
		// _ = "end of CoverTab[2306]"
//line /usr/local/go/src/embed/embed.go:257
	}
//line /usr/local/go/src/embed/embed.go:257
	// _ = "end of CoverTab[2299]"
//line /usr/local/go/src/embed/embed.go:257
	_go_fuzz_dep_.CoverTab[2300]++
						if name == "." {
//line /usr/local/go/src/embed/embed.go:258
		_go_fuzz_dep_.CoverTab[2307]++
							return dotFile
//line /usr/local/go/src/embed/embed.go:259
		// _ = "end of CoverTab[2307]"
	} else {
//line /usr/local/go/src/embed/embed.go:260
		_go_fuzz_dep_.CoverTab[2308]++
//line /usr/local/go/src/embed/embed.go:260
		// _ = "end of CoverTab[2308]"
//line /usr/local/go/src/embed/embed.go:260
	}
//line /usr/local/go/src/embed/embed.go:260
	// _ = "end of CoverTab[2300]"
//line /usr/local/go/src/embed/embed.go:260
	_go_fuzz_dep_.CoverTab[2301]++
						if f.files == nil {
//line /usr/local/go/src/embed/embed.go:261
		_go_fuzz_dep_.CoverTab[2309]++
							return nil
//line /usr/local/go/src/embed/embed.go:262
		// _ = "end of CoverTab[2309]"
	} else {
//line /usr/local/go/src/embed/embed.go:263
		_go_fuzz_dep_.CoverTab[2310]++
//line /usr/local/go/src/embed/embed.go:263
		// _ = "end of CoverTab[2310]"
//line /usr/local/go/src/embed/embed.go:263
	}
//line /usr/local/go/src/embed/embed.go:263
	// _ = "end of CoverTab[2301]"
//line /usr/local/go/src/embed/embed.go:263
	_go_fuzz_dep_.CoverTab[2302]++

//line /usr/local/go/src/embed/embed.go:267
	dir, elem, _ := split(name)
	files := *f.files
	i := sortSearch(len(files), func(i int) bool {
//line /usr/local/go/src/embed/embed.go:269
		_go_fuzz_dep_.CoverTab[2311]++
							idir, ielem, _ := split(files[i].name)
							return idir > dir || func() bool {
//line /usr/local/go/src/embed/embed.go:271
			_go_fuzz_dep_.CoverTab[2312]++
//line /usr/local/go/src/embed/embed.go:271
			return idir == dir && func() bool {
//line /usr/local/go/src/embed/embed.go:271
				_go_fuzz_dep_.CoverTab[2313]++
//line /usr/local/go/src/embed/embed.go:271
				return ielem >= elem
//line /usr/local/go/src/embed/embed.go:271
				// _ = "end of CoverTab[2313]"
//line /usr/local/go/src/embed/embed.go:271
			}()
//line /usr/local/go/src/embed/embed.go:271
			// _ = "end of CoverTab[2312]"
//line /usr/local/go/src/embed/embed.go:271
		}()
//line /usr/local/go/src/embed/embed.go:271
		// _ = "end of CoverTab[2311]"
	})
//line /usr/local/go/src/embed/embed.go:272
	// _ = "end of CoverTab[2302]"
//line /usr/local/go/src/embed/embed.go:272
	_go_fuzz_dep_.CoverTab[2303]++
						if i < len(files) && func() bool {
//line /usr/local/go/src/embed/embed.go:273
		_go_fuzz_dep_.CoverTab[2314]++
//line /usr/local/go/src/embed/embed.go:273
		return trimSlash(files[i].name) == name
//line /usr/local/go/src/embed/embed.go:273
		// _ = "end of CoverTab[2314]"
//line /usr/local/go/src/embed/embed.go:273
	}() {
//line /usr/local/go/src/embed/embed.go:273
		_go_fuzz_dep_.CoverTab[2315]++
							return &files[i]
//line /usr/local/go/src/embed/embed.go:274
		// _ = "end of CoverTab[2315]"
	} else {
//line /usr/local/go/src/embed/embed.go:275
		_go_fuzz_dep_.CoverTab[2316]++
//line /usr/local/go/src/embed/embed.go:275
		// _ = "end of CoverTab[2316]"
//line /usr/local/go/src/embed/embed.go:275
	}
//line /usr/local/go/src/embed/embed.go:275
	// _ = "end of CoverTab[2303]"
//line /usr/local/go/src/embed/embed.go:275
	_go_fuzz_dep_.CoverTab[2304]++
						return nil
//line /usr/local/go/src/embed/embed.go:276
	// _ = "end of CoverTab[2304]"
}

// readDir returns the list of files corresponding to the directory dir.
func (f FS) readDir(dir string) []file {
//line /usr/local/go/src/embed/embed.go:280
	_go_fuzz_dep_.CoverTab[2317]++
						if f.files == nil {
//line /usr/local/go/src/embed/embed.go:281
		_go_fuzz_dep_.CoverTab[2321]++
							return nil
//line /usr/local/go/src/embed/embed.go:282
		// _ = "end of CoverTab[2321]"
	} else {
//line /usr/local/go/src/embed/embed.go:283
		_go_fuzz_dep_.CoverTab[2322]++
//line /usr/local/go/src/embed/embed.go:283
		// _ = "end of CoverTab[2322]"
//line /usr/local/go/src/embed/embed.go:283
	}
//line /usr/local/go/src/embed/embed.go:283
	// _ = "end of CoverTab[2317]"
//line /usr/local/go/src/embed/embed.go:283
	_go_fuzz_dep_.CoverTab[2318]++

//line /usr/local/go/src/embed/embed.go:286
	files := *f.files
	i := sortSearch(len(files), func(i int) bool {
//line /usr/local/go/src/embed/embed.go:287
		_go_fuzz_dep_.CoverTab[2323]++
							idir, _, _ := split(files[i].name)
							return idir >= dir
//line /usr/local/go/src/embed/embed.go:289
		// _ = "end of CoverTab[2323]"
	})
//line /usr/local/go/src/embed/embed.go:290
	// _ = "end of CoverTab[2318]"
//line /usr/local/go/src/embed/embed.go:290
	_go_fuzz_dep_.CoverTab[2319]++
						j := sortSearch(len(files), func(j int) bool {
//line /usr/local/go/src/embed/embed.go:291
		_go_fuzz_dep_.CoverTab[2324]++
							jdir, _, _ := split(files[j].name)
							return jdir > dir
//line /usr/local/go/src/embed/embed.go:293
		// _ = "end of CoverTab[2324]"
	})
//line /usr/local/go/src/embed/embed.go:294
	// _ = "end of CoverTab[2319]"
//line /usr/local/go/src/embed/embed.go:294
	_go_fuzz_dep_.CoverTab[2320]++
						return files[i:j]
//line /usr/local/go/src/embed/embed.go:295
	// _ = "end of CoverTab[2320]"
}

// Open opens the named file for reading and returns it as an fs.File.
//line /usr/local/go/src/embed/embed.go:298
//
//line /usr/local/go/src/embed/embed.go:298
// The returned file implements io.Seeker when the file is not a directory.
//line /usr/local/go/src/embed/embed.go:301
func (f FS) Open(name string) (fs.File, error) {
//line /usr/local/go/src/embed/embed.go:301
	_go_fuzz_dep_.CoverTab[2325]++
						file := f.lookup(name)
						if file == nil {
//line /usr/local/go/src/embed/embed.go:303
		_go_fuzz_dep_.CoverTab[2328]++
							return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
//line /usr/local/go/src/embed/embed.go:304
		// _ = "end of CoverTab[2328]"
	} else {
//line /usr/local/go/src/embed/embed.go:305
		_go_fuzz_dep_.CoverTab[2329]++
//line /usr/local/go/src/embed/embed.go:305
		// _ = "end of CoverTab[2329]"
//line /usr/local/go/src/embed/embed.go:305
	}
//line /usr/local/go/src/embed/embed.go:305
	// _ = "end of CoverTab[2325]"
//line /usr/local/go/src/embed/embed.go:305
	_go_fuzz_dep_.CoverTab[2326]++
						if file.IsDir() {
//line /usr/local/go/src/embed/embed.go:306
		_go_fuzz_dep_.CoverTab[2330]++
							return &openDir{file, f.readDir(name), 0}, nil
//line /usr/local/go/src/embed/embed.go:307
		// _ = "end of CoverTab[2330]"
	} else {
//line /usr/local/go/src/embed/embed.go:308
		_go_fuzz_dep_.CoverTab[2331]++
//line /usr/local/go/src/embed/embed.go:308
		// _ = "end of CoverTab[2331]"
//line /usr/local/go/src/embed/embed.go:308
	}
//line /usr/local/go/src/embed/embed.go:308
	// _ = "end of CoverTab[2326]"
//line /usr/local/go/src/embed/embed.go:308
	_go_fuzz_dep_.CoverTab[2327]++
						return &openFile{file, 0}, nil
//line /usr/local/go/src/embed/embed.go:309
	// _ = "end of CoverTab[2327]"
}

// ReadDir reads and returns the entire named directory.
func (f FS) ReadDir(name string) ([]fs.DirEntry, error) {
//line /usr/local/go/src/embed/embed.go:313
	_go_fuzz_dep_.CoverTab[2332]++
						file, err := f.Open(name)
						if err != nil {
//line /usr/local/go/src/embed/embed.go:315
		_go_fuzz_dep_.CoverTab[2336]++
							return nil, err
//line /usr/local/go/src/embed/embed.go:316
		// _ = "end of CoverTab[2336]"
	} else {
//line /usr/local/go/src/embed/embed.go:317
		_go_fuzz_dep_.CoverTab[2337]++
//line /usr/local/go/src/embed/embed.go:317
		// _ = "end of CoverTab[2337]"
//line /usr/local/go/src/embed/embed.go:317
	}
//line /usr/local/go/src/embed/embed.go:317
	// _ = "end of CoverTab[2332]"
//line /usr/local/go/src/embed/embed.go:317
	_go_fuzz_dep_.CoverTab[2333]++
						dir, ok := file.(*openDir)
						if !ok {
//line /usr/local/go/src/embed/embed.go:319
		_go_fuzz_dep_.CoverTab[2338]++
							return nil, &fs.PathError{Op: "read", Path: name, Err: errors.New("not a directory")}
//line /usr/local/go/src/embed/embed.go:320
		// _ = "end of CoverTab[2338]"
	} else {
//line /usr/local/go/src/embed/embed.go:321
		_go_fuzz_dep_.CoverTab[2339]++
//line /usr/local/go/src/embed/embed.go:321
		// _ = "end of CoverTab[2339]"
//line /usr/local/go/src/embed/embed.go:321
	}
//line /usr/local/go/src/embed/embed.go:321
	// _ = "end of CoverTab[2333]"
//line /usr/local/go/src/embed/embed.go:321
	_go_fuzz_dep_.CoverTab[2334]++
						list := make([]fs.DirEntry, len(dir.files))
						for i := range list {
//line /usr/local/go/src/embed/embed.go:323
		_go_fuzz_dep_.CoverTab[2340]++
							list[i] = &dir.files[i]
//line /usr/local/go/src/embed/embed.go:324
		// _ = "end of CoverTab[2340]"
	}
//line /usr/local/go/src/embed/embed.go:325
	// _ = "end of CoverTab[2334]"
//line /usr/local/go/src/embed/embed.go:325
	_go_fuzz_dep_.CoverTab[2335]++
						return list, nil
//line /usr/local/go/src/embed/embed.go:326
	// _ = "end of CoverTab[2335]"
}

// ReadFile reads and returns the content of the named file.
func (f FS) ReadFile(name string) ([]byte, error) {
//line /usr/local/go/src/embed/embed.go:330
	_go_fuzz_dep_.CoverTab[2341]++
						file, err := f.Open(name)
						if err != nil {
//line /usr/local/go/src/embed/embed.go:332
		_go_fuzz_dep_.CoverTab[2344]++
							return nil, err
//line /usr/local/go/src/embed/embed.go:333
		// _ = "end of CoverTab[2344]"
	} else {
//line /usr/local/go/src/embed/embed.go:334
		_go_fuzz_dep_.CoverTab[2345]++
//line /usr/local/go/src/embed/embed.go:334
		// _ = "end of CoverTab[2345]"
//line /usr/local/go/src/embed/embed.go:334
	}
//line /usr/local/go/src/embed/embed.go:334
	// _ = "end of CoverTab[2341]"
//line /usr/local/go/src/embed/embed.go:334
	_go_fuzz_dep_.CoverTab[2342]++
						ofile, ok := file.(*openFile)
						if !ok {
//line /usr/local/go/src/embed/embed.go:336
		_go_fuzz_dep_.CoverTab[2346]++
							return nil, &fs.PathError{Op: "read", Path: name, Err: errors.New("is a directory")}
//line /usr/local/go/src/embed/embed.go:337
		// _ = "end of CoverTab[2346]"
	} else {
//line /usr/local/go/src/embed/embed.go:338
		_go_fuzz_dep_.CoverTab[2347]++
//line /usr/local/go/src/embed/embed.go:338
		// _ = "end of CoverTab[2347]"
//line /usr/local/go/src/embed/embed.go:338
	}
//line /usr/local/go/src/embed/embed.go:338
	// _ = "end of CoverTab[2342]"
//line /usr/local/go/src/embed/embed.go:338
	_go_fuzz_dep_.CoverTab[2343]++
						return []byte(ofile.f.data), nil
//line /usr/local/go/src/embed/embed.go:339
	// _ = "end of CoverTab[2343]"
}

// An openFile is a regular file open for reading.
type openFile struct {
	f	*file	// the file itself
	offset	int64	// current read offset
}

var (
	_ io.Seeker = (*openFile)(nil)
)

func (f *openFile) Close() error {
//line /usr/local/go/src/embed/embed.go:352
	_go_fuzz_dep_.CoverTab[2348]++
//line /usr/local/go/src/embed/embed.go:352
	return nil
//line /usr/local/go/src/embed/embed.go:352
	// _ = "end of CoverTab[2348]"
//line /usr/local/go/src/embed/embed.go:352
}
func (f *openFile) Stat() (fs.FileInfo, error) {
//line /usr/local/go/src/embed/embed.go:353
	_go_fuzz_dep_.CoverTab[2349]++
//line /usr/local/go/src/embed/embed.go:353
	return f.f, nil
//line /usr/local/go/src/embed/embed.go:353
	// _ = "end of CoverTab[2349]"
//line /usr/local/go/src/embed/embed.go:353
}

func (f *openFile) Read(b []byte) (int, error) {
//line /usr/local/go/src/embed/embed.go:355
	_go_fuzz_dep_.CoverTab[2350]++
						if f.offset >= int64(len(f.f.data)) {
//line /usr/local/go/src/embed/embed.go:356
		_go_fuzz_dep_.CoverTab[2353]++
							return 0, io.EOF
//line /usr/local/go/src/embed/embed.go:357
		// _ = "end of CoverTab[2353]"
	} else {
//line /usr/local/go/src/embed/embed.go:358
		_go_fuzz_dep_.CoverTab[2354]++
//line /usr/local/go/src/embed/embed.go:358
		// _ = "end of CoverTab[2354]"
//line /usr/local/go/src/embed/embed.go:358
	}
//line /usr/local/go/src/embed/embed.go:358
	// _ = "end of CoverTab[2350]"
//line /usr/local/go/src/embed/embed.go:358
	_go_fuzz_dep_.CoverTab[2351]++
						if f.offset < 0 {
//line /usr/local/go/src/embed/embed.go:359
		_go_fuzz_dep_.CoverTab[2355]++
							return 0, &fs.PathError{Op: "read", Path: f.f.name, Err: fs.ErrInvalid}
//line /usr/local/go/src/embed/embed.go:360
		// _ = "end of CoverTab[2355]"
	} else {
//line /usr/local/go/src/embed/embed.go:361
		_go_fuzz_dep_.CoverTab[2356]++
//line /usr/local/go/src/embed/embed.go:361
		// _ = "end of CoverTab[2356]"
//line /usr/local/go/src/embed/embed.go:361
	}
//line /usr/local/go/src/embed/embed.go:361
	// _ = "end of CoverTab[2351]"
//line /usr/local/go/src/embed/embed.go:361
	_go_fuzz_dep_.CoverTab[2352]++
						n := copy(b, f.f.data[f.offset:])
						f.offset += int64(n)
						return n, nil
//line /usr/local/go/src/embed/embed.go:364
	// _ = "end of CoverTab[2352]"
}

func (f *openFile) Seek(offset int64, whence int) (int64, error) {
//line /usr/local/go/src/embed/embed.go:367
	_go_fuzz_dep_.CoverTab[2357]++
						switch whence {
	case 0:
//line /usr/local/go/src/embed/embed.go:369
		_go_fuzz_dep_.CoverTab[2360]++
//line /usr/local/go/src/embed/embed.go:369
		// _ = "end of CoverTab[2360]"

	case 1:
//line /usr/local/go/src/embed/embed.go:371
		_go_fuzz_dep_.CoverTab[2361]++
							offset += f.offset
//line /usr/local/go/src/embed/embed.go:372
		// _ = "end of CoverTab[2361]"
	case 2:
//line /usr/local/go/src/embed/embed.go:373
		_go_fuzz_dep_.CoverTab[2362]++
							offset += int64(len(f.f.data))
//line /usr/local/go/src/embed/embed.go:374
		// _ = "end of CoverTab[2362]"
//line /usr/local/go/src/embed/embed.go:374
	default:
//line /usr/local/go/src/embed/embed.go:374
		_go_fuzz_dep_.CoverTab[2363]++
//line /usr/local/go/src/embed/embed.go:374
		// _ = "end of CoverTab[2363]"
	}
//line /usr/local/go/src/embed/embed.go:375
	// _ = "end of CoverTab[2357]"
//line /usr/local/go/src/embed/embed.go:375
	_go_fuzz_dep_.CoverTab[2358]++
						if offset < 0 || func() bool {
//line /usr/local/go/src/embed/embed.go:376
		_go_fuzz_dep_.CoverTab[2364]++
//line /usr/local/go/src/embed/embed.go:376
		return offset > int64(len(f.f.data))
//line /usr/local/go/src/embed/embed.go:376
		// _ = "end of CoverTab[2364]"
//line /usr/local/go/src/embed/embed.go:376
	}() {
//line /usr/local/go/src/embed/embed.go:376
		_go_fuzz_dep_.CoverTab[2365]++
							return 0, &fs.PathError{Op: "seek", Path: f.f.name, Err: fs.ErrInvalid}
//line /usr/local/go/src/embed/embed.go:377
		// _ = "end of CoverTab[2365]"
	} else {
//line /usr/local/go/src/embed/embed.go:378
		_go_fuzz_dep_.CoverTab[2366]++
//line /usr/local/go/src/embed/embed.go:378
		// _ = "end of CoverTab[2366]"
//line /usr/local/go/src/embed/embed.go:378
	}
//line /usr/local/go/src/embed/embed.go:378
	// _ = "end of CoverTab[2358]"
//line /usr/local/go/src/embed/embed.go:378
	_go_fuzz_dep_.CoverTab[2359]++
						f.offset = offset
						return offset, nil
//line /usr/local/go/src/embed/embed.go:380
	// _ = "end of CoverTab[2359]"
}

// An openDir is a directory open for reading.
type openDir struct {
	f	*file	// the directory file itself
	files	[]file	// the directory contents
	offset	int	// the read offset, an index into the files slice
}

func (d *openDir) Close() error {
//line /usr/local/go/src/embed/embed.go:390
	_go_fuzz_dep_.CoverTab[2367]++
//line /usr/local/go/src/embed/embed.go:390
	return nil
//line /usr/local/go/src/embed/embed.go:390
	// _ = "end of CoverTab[2367]"
//line /usr/local/go/src/embed/embed.go:390
}
func (d *openDir) Stat() (fs.FileInfo, error) {
//line /usr/local/go/src/embed/embed.go:391
	_go_fuzz_dep_.CoverTab[2368]++
//line /usr/local/go/src/embed/embed.go:391
	return d.f, nil
//line /usr/local/go/src/embed/embed.go:391
	// _ = "end of CoverTab[2368]"
//line /usr/local/go/src/embed/embed.go:391
}

func (d *openDir) Read([]byte) (int, error) {
//line /usr/local/go/src/embed/embed.go:393
	_go_fuzz_dep_.CoverTab[2369]++
						return 0, &fs.PathError{Op: "read", Path: d.f.name, Err: errors.New("is a directory")}
//line /usr/local/go/src/embed/embed.go:394
	// _ = "end of CoverTab[2369]"
}

func (d *openDir) ReadDir(count int) ([]fs.DirEntry, error) {
//line /usr/local/go/src/embed/embed.go:397
	_go_fuzz_dep_.CoverTab[2370]++
						n := len(d.files) - d.offset
						if n == 0 {
//line /usr/local/go/src/embed/embed.go:399
		_go_fuzz_dep_.CoverTab[2374]++
							if count <= 0 {
//line /usr/local/go/src/embed/embed.go:400
			_go_fuzz_dep_.CoverTab[2376]++
								return nil, nil
//line /usr/local/go/src/embed/embed.go:401
			// _ = "end of CoverTab[2376]"
		} else {
//line /usr/local/go/src/embed/embed.go:402
			_go_fuzz_dep_.CoverTab[2377]++
//line /usr/local/go/src/embed/embed.go:402
			// _ = "end of CoverTab[2377]"
//line /usr/local/go/src/embed/embed.go:402
		}
//line /usr/local/go/src/embed/embed.go:402
		// _ = "end of CoverTab[2374]"
//line /usr/local/go/src/embed/embed.go:402
		_go_fuzz_dep_.CoverTab[2375]++
							return nil, io.EOF
//line /usr/local/go/src/embed/embed.go:403
		// _ = "end of CoverTab[2375]"
	} else {
//line /usr/local/go/src/embed/embed.go:404
		_go_fuzz_dep_.CoverTab[2378]++
//line /usr/local/go/src/embed/embed.go:404
		// _ = "end of CoverTab[2378]"
//line /usr/local/go/src/embed/embed.go:404
	}
//line /usr/local/go/src/embed/embed.go:404
	// _ = "end of CoverTab[2370]"
//line /usr/local/go/src/embed/embed.go:404
	_go_fuzz_dep_.CoverTab[2371]++
						if count > 0 && func() bool {
//line /usr/local/go/src/embed/embed.go:405
		_go_fuzz_dep_.CoverTab[2379]++
//line /usr/local/go/src/embed/embed.go:405
		return n > count
//line /usr/local/go/src/embed/embed.go:405
		// _ = "end of CoverTab[2379]"
//line /usr/local/go/src/embed/embed.go:405
	}() {
//line /usr/local/go/src/embed/embed.go:405
		_go_fuzz_dep_.CoverTab[2380]++
							n = count
//line /usr/local/go/src/embed/embed.go:406
		// _ = "end of CoverTab[2380]"
	} else {
//line /usr/local/go/src/embed/embed.go:407
		_go_fuzz_dep_.CoverTab[2381]++
//line /usr/local/go/src/embed/embed.go:407
		// _ = "end of CoverTab[2381]"
//line /usr/local/go/src/embed/embed.go:407
	}
//line /usr/local/go/src/embed/embed.go:407
	// _ = "end of CoverTab[2371]"
//line /usr/local/go/src/embed/embed.go:407
	_go_fuzz_dep_.CoverTab[2372]++
						list := make([]fs.DirEntry, n)
						for i := range list {
//line /usr/local/go/src/embed/embed.go:409
		_go_fuzz_dep_.CoverTab[2382]++
							list[i] = &d.files[d.offset+i]
//line /usr/local/go/src/embed/embed.go:410
		// _ = "end of CoverTab[2382]"
	}
//line /usr/local/go/src/embed/embed.go:411
	// _ = "end of CoverTab[2372]"
//line /usr/local/go/src/embed/embed.go:411
	_go_fuzz_dep_.CoverTab[2373]++
						d.offset += n
						return list, nil
//line /usr/local/go/src/embed/embed.go:413
	// _ = "end of CoverTab[2373]"
}

// sortSearch is like sort.Search, avoiding an import.
func sortSearch(n int, f func(int) bool) int {
//line /usr/local/go/src/embed/embed.go:417
	_go_fuzz_dep_.CoverTab[2383]++

//line /usr/local/go/src/embed/embed.go:420
	i, j := 0, n
	for i < j {
//line /usr/local/go/src/embed/embed.go:421
		_go_fuzz_dep_.CoverTab[2385]++
							h := int(uint(i+j) >> 1)

							if !f(h) {
//line /usr/local/go/src/embed/embed.go:424
			_go_fuzz_dep_.CoverTab[2386]++
								i = h + 1
//line /usr/local/go/src/embed/embed.go:425
			// _ = "end of CoverTab[2386]"
		} else {
//line /usr/local/go/src/embed/embed.go:426
			_go_fuzz_dep_.CoverTab[2387]++
								j = h
//line /usr/local/go/src/embed/embed.go:427
			// _ = "end of CoverTab[2387]"
		}
//line /usr/local/go/src/embed/embed.go:428
		// _ = "end of CoverTab[2385]"
	}
//line /usr/local/go/src/embed/embed.go:429
	// _ = "end of CoverTab[2383]"
//line /usr/local/go/src/embed/embed.go:429
	_go_fuzz_dep_.CoverTab[2384]++

						return i
//line /usr/local/go/src/embed/embed.go:431
	// _ = "end of CoverTab[2384]"
}

//line /usr/local/go/src/embed/embed.go:432
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/embed/embed.go:432
var _ = _go_fuzz_dep_.CoverTab
