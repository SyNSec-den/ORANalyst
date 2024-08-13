// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP file system request handler

//line /usr/local/go/src/net/http/fs.go:7
package http

//line /usr/local/go/src/net/http/fs.go:7
import (
//line /usr/local/go/src/net/http/fs.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/fs.go:7
)
//line /usr/local/go/src/net/http/fs.go:7
import (
//line /usr/local/go/src/net/http/fs.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/fs.go:7
)

import (
	"errors"
	"fmt"
	"internal/safefilepath"
	"io"
	"io/fs"
	"mime"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// A Dir implements FileSystem using the native file system restricted to a
//line /usr/local/go/src/net/http/fs.go:28
// specific directory tree.
//line /usr/local/go/src/net/http/fs.go:28
//
//line /usr/local/go/src/net/http/fs.go:28
// While the FileSystem.Open method takes '/'-separated paths, a Dir's string
//line /usr/local/go/src/net/http/fs.go:28
// value is a filename on the native file system, not a URL, so it is separated
//line /usr/local/go/src/net/http/fs.go:28
// by filepath.Separator, which isn't necessarily '/'.
//line /usr/local/go/src/net/http/fs.go:28
//
//line /usr/local/go/src/net/http/fs.go:28
// Note that Dir could expose sensitive files and directories. Dir will follow
//line /usr/local/go/src/net/http/fs.go:28
// symlinks pointing out of the directory tree, which can be especially dangerous
//line /usr/local/go/src/net/http/fs.go:28
// if serving from a directory in which users are able to create arbitrary symlinks.
//line /usr/local/go/src/net/http/fs.go:28
// Dir will also allow access to files and directories starting with a period,
//line /usr/local/go/src/net/http/fs.go:28
// which could expose sensitive directories like .git or sensitive files like
//line /usr/local/go/src/net/http/fs.go:28
// .htpasswd. To exclude files with a leading period, remove the files/directories
//line /usr/local/go/src/net/http/fs.go:28
// from the server or create a custom FileSystem implementation.
//line /usr/local/go/src/net/http/fs.go:28
//
//line /usr/local/go/src/net/http/fs.go:28
// An empty Dir is treated as ".".
//line /usr/local/go/src/net/http/fs.go:44
type Dir string

// mapOpenError maps the provided non-nil error from opening name
//line /usr/local/go/src/net/http/fs.go:46
// to a possibly better non-nil error. In particular, it turns OS-specific errors
//line /usr/local/go/src/net/http/fs.go:46
// about opening files in non-directories into fs.ErrNotExist. See Issues 18984 and 49552.
//line /usr/local/go/src/net/http/fs.go:49
func mapOpenError(originalErr error, name string, sep rune, stat func(string) (fs.FileInfo, error)) error {
//line /usr/local/go/src/net/http/fs.go:49
	_go_fuzz_dep_.CoverTab[37180]++
						if errors.Is(originalErr, fs.ErrNotExist) || func() bool {
//line /usr/local/go/src/net/http/fs.go:50
		_go_fuzz_dep_.CoverTab[37183]++
//line /usr/local/go/src/net/http/fs.go:50
		return errors.Is(originalErr, fs.ErrPermission)
//line /usr/local/go/src/net/http/fs.go:50
		// _ = "end of CoverTab[37183]"
//line /usr/local/go/src/net/http/fs.go:50
	}() {
//line /usr/local/go/src/net/http/fs.go:50
		_go_fuzz_dep_.CoverTab[37184]++
							return originalErr
//line /usr/local/go/src/net/http/fs.go:51
		// _ = "end of CoverTab[37184]"
	} else {
//line /usr/local/go/src/net/http/fs.go:52
		_go_fuzz_dep_.CoverTab[37185]++
//line /usr/local/go/src/net/http/fs.go:52
		// _ = "end of CoverTab[37185]"
//line /usr/local/go/src/net/http/fs.go:52
	}
//line /usr/local/go/src/net/http/fs.go:52
	// _ = "end of CoverTab[37180]"
//line /usr/local/go/src/net/http/fs.go:52
	_go_fuzz_dep_.CoverTab[37181]++

						parts := strings.Split(name, string(sep))
						for i := range parts {
//line /usr/local/go/src/net/http/fs.go:55
		_go_fuzz_dep_.CoverTab[37186]++
							if parts[i] == "" {
//line /usr/local/go/src/net/http/fs.go:56
			_go_fuzz_dep_.CoverTab[37189]++
								continue
//line /usr/local/go/src/net/http/fs.go:57
			// _ = "end of CoverTab[37189]"
		} else {
//line /usr/local/go/src/net/http/fs.go:58
			_go_fuzz_dep_.CoverTab[37190]++
//line /usr/local/go/src/net/http/fs.go:58
			// _ = "end of CoverTab[37190]"
//line /usr/local/go/src/net/http/fs.go:58
		}
//line /usr/local/go/src/net/http/fs.go:58
		// _ = "end of CoverTab[37186]"
//line /usr/local/go/src/net/http/fs.go:58
		_go_fuzz_dep_.CoverTab[37187]++
							fi, err := stat(strings.Join(parts[:i+1], string(sep)))
							if err != nil {
//line /usr/local/go/src/net/http/fs.go:60
			_go_fuzz_dep_.CoverTab[37191]++
								return originalErr
//line /usr/local/go/src/net/http/fs.go:61
			// _ = "end of CoverTab[37191]"
		} else {
//line /usr/local/go/src/net/http/fs.go:62
			_go_fuzz_dep_.CoverTab[37192]++
//line /usr/local/go/src/net/http/fs.go:62
			// _ = "end of CoverTab[37192]"
//line /usr/local/go/src/net/http/fs.go:62
		}
//line /usr/local/go/src/net/http/fs.go:62
		// _ = "end of CoverTab[37187]"
//line /usr/local/go/src/net/http/fs.go:62
		_go_fuzz_dep_.CoverTab[37188]++
							if !fi.IsDir() {
//line /usr/local/go/src/net/http/fs.go:63
			_go_fuzz_dep_.CoverTab[37193]++
								return fs.ErrNotExist
//line /usr/local/go/src/net/http/fs.go:64
			// _ = "end of CoverTab[37193]"
		} else {
//line /usr/local/go/src/net/http/fs.go:65
			_go_fuzz_dep_.CoverTab[37194]++
//line /usr/local/go/src/net/http/fs.go:65
			// _ = "end of CoverTab[37194]"
//line /usr/local/go/src/net/http/fs.go:65
		}
//line /usr/local/go/src/net/http/fs.go:65
		// _ = "end of CoverTab[37188]"
	}
//line /usr/local/go/src/net/http/fs.go:66
	// _ = "end of CoverTab[37181]"
//line /usr/local/go/src/net/http/fs.go:66
	_go_fuzz_dep_.CoverTab[37182]++
						return originalErr
//line /usr/local/go/src/net/http/fs.go:67
	// _ = "end of CoverTab[37182]"
}

// Open implements FileSystem using os.Open, opening files for reading rooted
//line /usr/local/go/src/net/http/fs.go:70
// and relative to the directory d.
//line /usr/local/go/src/net/http/fs.go:72
func (d Dir) Open(name string) (File, error) {
//line /usr/local/go/src/net/http/fs.go:72
	_go_fuzz_dep_.CoverTab[37195]++
						path, err := safefilepath.FromFS(path.Clean("/" + name))
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:74
		_go_fuzz_dep_.CoverTab[37199]++
							return nil, errors.New("http: invalid or unsafe file path")
//line /usr/local/go/src/net/http/fs.go:75
		// _ = "end of CoverTab[37199]"
	} else {
//line /usr/local/go/src/net/http/fs.go:76
		_go_fuzz_dep_.CoverTab[37200]++
//line /usr/local/go/src/net/http/fs.go:76
		// _ = "end of CoverTab[37200]"
//line /usr/local/go/src/net/http/fs.go:76
	}
//line /usr/local/go/src/net/http/fs.go:76
	// _ = "end of CoverTab[37195]"
//line /usr/local/go/src/net/http/fs.go:76
	_go_fuzz_dep_.CoverTab[37196]++
						dir := string(d)
						if dir == "" {
//line /usr/local/go/src/net/http/fs.go:78
		_go_fuzz_dep_.CoverTab[37201]++
							dir = "."
//line /usr/local/go/src/net/http/fs.go:79
		// _ = "end of CoverTab[37201]"
	} else {
//line /usr/local/go/src/net/http/fs.go:80
		_go_fuzz_dep_.CoverTab[37202]++
//line /usr/local/go/src/net/http/fs.go:80
		// _ = "end of CoverTab[37202]"
//line /usr/local/go/src/net/http/fs.go:80
	}
//line /usr/local/go/src/net/http/fs.go:80
	// _ = "end of CoverTab[37196]"
//line /usr/local/go/src/net/http/fs.go:80
	_go_fuzz_dep_.CoverTab[37197]++
						fullName := filepath.Join(dir, path)
						f, err := os.Open(fullName)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:83
		_go_fuzz_dep_.CoverTab[37203]++
							return nil, mapOpenError(err, fullName, filepath.Separator, os.Stat)
//line /usr/local/go/src/net/http/fs.go:84
		// _ = "end of CoverTab[37203]"
	} else {
//line /usr/local/go/src/net/http/fs.go:85
		_go_fuzz_dep_.CoverTab[37204]++
//line /usr/local/go/src/net/http/fs.go:85
		// _ = "end of CoverTab[37204]"
//line /usr/local/go/src/net/http/fs.go:85
	}
//line /usr/local/go/src/net/http/fs.go:85
	// _ = "end of CoverTab[37197]"
//line /usr/local/go/src/net/http/fs.go:85
	_go_fuzz_dep_.CoverTab[37198]++
						return f, nil
//line /usr/local/go/src/net/http/fs.go:86
	// _ = "end of CoverTab[37198]"
}

// A FileSystem implements access to a collection of named files.
//line /usr/local/go/src/net/http/fs.go:89
// The elements in a file path are separated by slash ('/', U+002F)
//line /usr/local/go/src/net/http/fs.go:89
// characters, regardless of host operating system convention.
//line /usr/local/go/src/net/http/fs.go:89
// See the FileServer function to convert a FileSystem to a Handler.
//line /usr/local/go/src/net/http/fs.go:89
//
//line /usr/local/go/src/net/http/fs.go:89
// This interface predates the fs.FS interface, which can be used instead:
//line /usr/local/go/src/net/http/fs.go:89
// the FS adapter function converts an fs.FS to a FileSystem.
//line /usr/local/go/src/net/http/fs.go:96
type FileSystem interface {
	Open(name string) (File, error)
}

// A File is returned by a FileSystem's Open method and can be
//line /usr/local/go/src/net/http/fs.go:100
// served by the FileServer implementation.
//line /usr/local/go/src/net/http/fs.go:100
//
//line /usr/local/go/src/net/http/fs.go:100
// The methods should behave the same as those on an *os.File.
//line /usr/local/go/src/net/http/fs.go:104
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
}

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int {
//line /usr/local/go/src/net/http/fs.go:120
	_go_fuzz_dep_.CoverTab[37205]++
//line /usr/local/go/src/net/http/fs.go:120
	return len(d)
//line /usr/local/go/src/net/http/fs.go:120
	// _ = "end of CoverTab[37205]"
//line /usr/local/go/src/net/http/fs.go:120
}
func (d fileInfoDirs) isDir(i int) bool {
//line /usr/local/go/src/net/http/fs.go:121
	_go_fuzz_dep_.CoverTab[37206]++
//line /usr/local/go/src/net/http/fs.go:121
	return d[i].IsDir()
//line /usr/local/go/src/net/http/fs.go:121
	// _ = "end of CoverTab[37206]"
//line /usr/local/go/src/net/http/fs.go:121
}
func (d fileInfoDirs) name(i int) string {
//line /usr/local/go/src/net/http/fs.go:122
	_go_fuzz_dep_.CoverTab[37207]++
//line /usr/local/go/src/net/http/fs.go:122
	return d[i].Name()
//line /usr/local/go/src/net/http/fs.go:122
	// _ = "end of CoverTab[37207]"
//line /usr/local/go/src/net/http/fs.go:122
}

type dirEntryDirs []fs.DirEntry

func (d dirEntryDirs) len() int {
//line /usr/local/go/src/net/http/fs.go:126
	_go_fuzz_dep_.CoverTab[37208]++
//line /usr/local/go/src/net/http/fs.go:126
	return len(d)
//line /usr/local/go/src/net/http/fs.go:126
	// _ = "end of CoverTab[37208]"
//line /usr/local/go/src/net/http/fs.go:126
}
func (d dirEntryDirs) isDir(i int) bool {
//line /usr/local/go/src/net/http/fs.go:127
	_go_fuzz_dep_.CoverTab[37209]++
//line /usr/local/go/src/net/http/fs.go:127
	return d[i].IsDir()
//line /usr/local/go/src/net/http/fs.go:127
	// _ = "end of CoverTab[37209]"
//line /usr/local/go/src/net/http/fs.go:127
}
func (d dirEntryDirs) name(i int) string {
//line /usr/local/go/src/net/http/fs.go:128
	_go_fuzz_dep_.CoverTab[37210]++
//line /usr/local/go/src/net/http/fs.go:128
	return d[i].Name()
//line /usr/local/go/src/net/http/fs.go:128
	// _ = "end of CoverTab[37210]"
//line /usr/local/go/src/net/http/fs.go:128
}

func dirList(w ResponseWriter, r *Request, f File) {
//line /usr/local/go/src/net/http/fs.go:130
	_go_fuzz_dep_.CoverTab[37211]++
	// Prefer to use ReadDir instead of Readdir,
	// because the former doesn't require calling
	// Stat on every entry of a directory on Unix.
	var dirs anyDirs
	var err error
	if d, ok := f.(fs.ReadDirFile); ok {
//line /usr/local/go/src/net/http/fs.go:136
		_go_fuzz_dep_.CoverTab[37216]++
							var list dirEntryDirs
							list, err = d.ReadDir(-1)
							dirs = list
//line /usr/local/go/src/net/http/fs.go:139
		// _ = "end of CoverTab[37216]"
	} else {
//line /usr/local/go/src/net/http/fs.go:140
		_go_fuzz_dep_.CoverTab[37217]++
							var list fileInfoDirs
							list, err = f.Readdir(-1)
							dirs = list
//line /usr/local/go/src/net/http/fs.go:143
		// _ = "end of CoverTab[37217]"
	}
//line /usr/local/go/src/net/http/fs.go:144
	// _ = "end of CoverTab[37211]"
//line /usr/local/go/src/net/http/fs.go:144
	_go_fuzz_dep_.CoverTab[37212]++

						if err != nil {
//line /usr/local/go/src/net/http/fs.go:146
		_go_fuzz_dep_.CoverTab[37218]++
							logf(r, "http: error reading directory: %v", err)
							Error(w, "Error reading directory", StatusInternalServerError)
							return
//line /usr/local/go/src/net/http/fs.go:149
		// _ = "end of CoverTab[37218]"
	} else {
//line /usr/local/go/src/net/http/fs.go:150
		_go_fuzz_dep_.CoverTab[37219]++
//line /usr/local/go/src/net/http/fs.go:150
		// _ = "end of CoverTab[37219]"
//line /usr/local/go/src/net/http/fs.go:150
	}
//line /usr/local/go/src/net/http/fs.go:150
	// _ = "end of CoverTab[37212]"
//line /usr/local/go/src/net/http/fs.go:150
	_go_fuzz_dep_.CoverTab[37213]++
						sort.Slice(dirs, func(i, j int) bool {
//line /usr/local/go/src/net/http/fs.go:151
		_go_fuzz_dep_.CoverTab[37220]++
//line /usr/local/go/src/net/http/fs.go:151
		return dirs.name(i) < dirs.name(j)
//line /usr/local/go/src/net/http/fs.go:151
		// _ = "end of CoverTab[37220]"
//line /usr/local/go/src/net/http/fs.go:151
	})
//line /usr/local/go/src/net/http/fs.go:151
	// _ = "end of CoverTab[37213]"
//line /usr/local/go/src/net/http/fs.go:151
	_go_fuzz_dep_.CoverTab[37214]++

						w.Header().Set("Content-Type", "text/html; charset=utf-8")
						fmt.Fprintf(w, "<pre>\n")
						for i, n := 0, dirs.len(); i < n; i++ {
//line /usr/local/go/src/net/http/fs.go:155
		_go_fuzz_dep_.CoverTab[37221]++
							name := dirs.name(i)
							if dirs.isDir(i) {
//line /usr/local/go/src/net/http/fs.go:157
			_go_fuzz_dep_.CoverTab[37223]++
								name += "/"
//line /usr/local/go/src/net/http/fs.go:158
			// _ = "end of CoverTab[37223]"
		} else {
//line /usr/local/go/src/net/http/fs.go:159
			_go_fuzz_dep_.CoverTab[37224]++
//line /usr/local/go/src/net/http/fs.go:159
			// _ = "end of CoverTab[37224]"
//line /usr/local/go/src/net/http/fs.go:159
		}
//line /usr/local/go/src/net/http/fs.go:159
		// _ = "end of CoverTab[37221]"
//line /usr/local/go/src/net/http/fs.go:159
		_go_fuzz_dep_.CoverTab[37222]++

//line /usr/local/go/src/net/http/fs.go:163
		url := url.URL{Path: name}
							fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
//line /usr/local/go/src/net/http/fs.go:164
		// _ = "end of CoverTab[37222]"
	}
//line /usr/local/go/src/net/http/fs.go:165
	// _ = "end of CoverTab[37214]"
//line /usr/local/go/src/net/http/fs.go:165
	_go_fuzz_dep_.CoverTab[37215]++
						fmt.Fprintf(w, "</pre>\n")
//line /usr/local/go/src/net/http/fs.go:166
	// _ = "end of CoverTab[37215]"
}

// ServeContent replies to the request using the content in the
//line /usr/local/go/src/net/http/fs.go:169
// provided ReadSeeker. The main benefit of ServeContent over io.Copy
//line /usr/local/go/src/net/http/fs.go:169
// is that it handles Range requests properly, sets the MIME type, and
//line /usr/local/go/src/net/http/fs.go:169
// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
//line /usr/local/go/src/net/http/fs.go:169
// and If-Range requests.
//line /usr/local/go/src/net/http/fs.go:169
//
//line /usr/local/go/src/net/http/fs.go:169
// If the response's Content-Type header is not set, ServeContent
//line /usr/local/go/src/net/http/fs.go:169
// first tries to deduce the type from name's file extension and,
//line /usr/local/go/src/net/http/fs.go:169
// if that fails, falls back to reading the first block of the content
//line /usr/local/go/src/net/http/fs.go:169
// and passing it to DetectContentType.
//line /usr/local/go/src/net/http/fs.go:169
// The name is otherwise unused; in particular it can be empty and is
//line /usr/local/go/src/net/http/fs.go:169
// never sent in the response.
//line /usr/local/go/src/net/http/fs.go:169
//
//line /usr/local/go/src/net/http/fs.go:169
// If modtime is not the zero time or Unix epoch, ServeContent
//line /usr/local/go/src/net/http/fs.go:169
// includes it in a Last-Modified header in the response. If the
//line /usr/local/go/src/net/http/fs.go:169
// request includes an If-Modified-Since header, ServeContent uses
//line /usr/local/go/src/net/http/fs.go:169
// modtime to decide whether the content needs to be sent at all.
//line /usr/local/go/src/net/http/fs.go:169
//
//line /usr/local/go/src/net/http/fs.go:169
// The content's Seek method must work: ServeContent uses
//line /usr/local/go/src/net/http/fs.go:169
// a seek to the end of the content to determine its size.
//line /usr/local/go/src/net/http/fs.go:169
//
//line /usr/local/go/src/net/http/fs.go:169
// If the caller has set w's ETag header formatted per RFC 7232, section 2.3,
//line /usr/local/go/src/net/http/fs.go:169
// ServeContent uses it to handle requests using If-Match, If-None-Match, or If-Range.
//line /usr/local/go/src/net/http/fs.go:169
//
//line /usr/local/go/src/net/http/fs.go:169
// Note that *os.File implements the io.ReadSeeker interface.
//line /usr/local/go/src/net/http/fs.go:194
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker) {
//line /usr/local/go/src/net/http/fs.go:194
	_go_fuzz_dep_.CoverTab[37225]++
						sizeFunc := func() (int64, error) {
//line /usr/local/go/src/net/http/fs.go:195
		_go_fuzz_dep_.CoverTab[37227]++
							size, err := content.Seek(0, io.SeekEnd)
							if err != nil {
//line /usr/local/go/src/net/http/fs.go:197
			_go_fuzz_dep_.CoverTab[37230]++
								return 0, errSeeker
//line /usr/local/go/src/net/http/fs.go:198
			// _ = "end of CoverTab[37230]"
		} else {
//line /usr/local/go/src/net/http/fs.go:199
			_go_fuzz_dep_.CoverTab[37231]++
//line /usr/local/go/src/net/http/fs.go:199
			// _ = "end of CoverTab[37231]"
//line /usr/local/go/src/net/http/fs.go:199
		}
//line /usr/local/go/src/net/http/fs.go:199
		// _ = "end of CoverTab[37227]"
//line /usr/local/go/src/net/http/fs.go:199
		_go_fuzz_dep_.CoverTab[37228]++
							_, err = content.Seek(0, io.SeekStart)
							if err != nil {
//line /usr/local/go/src/net/http/fs.go:201
			_go_fuzz_dep_.CoverTab[37232]++
								return 0, errSeeker
//line /usr/local/go/src/net/http/fs.go:202
			// _ = "end of CoverTab[37232]"
		} else {
//line /usr/local/go/src/net/http/fs.go:203
			_go_fuzz_dep_.CoverTab[37233]++
//line /usr/local/go/src/net/http/fs.go:203
			// _ = "end of CoverTab[37233]"
//line /usr/local/go/src/net/http/fs.go:203
		}
//line /usr/local/go/src/net/http/fs.go:203
		// _ = "end of CoverTab[37228]"
//line /usr/local/go/src/net/http/fs.go:203
		_go_fuzz_dep_.CoverTab[37229]++
							return size, nil
//line /usr/local/go/src/net/http/fs.go:204
		// _ = "end of CoverTab[37229]"
	}
//line /usr/local/go/src/net/http/fs.go:205
	// _ = "end of CoverTab[37225]"
//line /usr/local/go/src/net/http/fs.go:205
	_go_fuzz_dep_.CoverTab[37226]++
						serveContent(w, req, name, modtime, sizeFunc, content)
//line /usr/local/go/src/net/http/fs.go:206
	// _ = "end of CoverTab[37226]"
}

// errSeeker is returned by ServeContent's sizeFunc when the content
//line /usr/local/go/src/net/http/fs.go:209
// doesn't seek properly. The underlying Seeker's error text isn't
//line /usr/local/go/src/net/http/fs.go:209
// included in the sizeFunc reply so it's not sent over HTTP to end
//line /usr/local/go/src/net/http/fs.go:209
// users.
//line /usr/local/go/src/net/http/fs.go:213
var errSeeker = errors.New("seeker can't seek")

// errNoOverlap is returned by serveContent's parseRange if first-byte-pos of
//line /usr/local/go/src/net/http/fs.go:215
// all of the byte-range-spec values is greater than the content size.
//line /usr/local/go/src/net/http/fs.go:217
var errNoOverlap = errors.New("invalid range: failed to overlap")

// if name is empty, filename is unknown. (used for mime type, before sniffing)
//line /usr/local/go/src/net/http/fs.go:219
// if modtime.IsZero(), modtime is unknown.
//line /usr/local/go/src/net/http/fs.go:219
// content must be seeked to the beginning of the file.
//line /usr/local/go/src/net/http/fs.go:219
// The sizeFunc is called at most once. Its error, if any, is sent in the HTTP response.
//line /usr/local/go/src/net/http/fs.go:223
func serveContent(w ResponseWriter, r *Request, name string, modtime time.Time, sizeFunc func() (int64, error), content io.ReadSeeker) {
//line /usr/local/go/src/net/http/fs.go:223
	_go_fuzz_dep_.CoverTab[37234]++
						setLastModified(w, modtime)
						done, rangeReq := checkPreconditions(w, r, modtime)
						if done {
//line /usr/local/go/src/net/http/fs.go:226
		_go_fuzz_dep_.CoverTab[37243]++
							return
//line /usr/local/go/src/net/http/fs.go:227
		// _ = "end of CoverTab[37243]"
	} else {
//line /usr/local/go/src/net/http/fs.go:228
		_go_fuzz_dep_.CoverTab[37244]++
//line /usr/local/go/src/net/http/fs.go:228
		// _ = "end of CoverTab[37244]"
//line /usr/local/go/src/net/http/fs.go:228
	}
//line /usr/local/go/src/net/http/fs.go:228
	// _ = "end of CoverTab[37234]"
//line /usr/local/go/src/net/http/fs.go:228
	_go_fuzz_dep_.CoverTab[37235]++

						code := StatusOK

//line /usr/local/go/src/net/http/fs.go:234
	ctypes, haveType := w.Header()["Content-Type"]
	var ctype string
	if !haveType {
//line /usr/local/go/src/net/http/fs.go:236
		_go_fuzz_dep_.CoverTab[37245]++
							ctype = mime.TypeByExtension(filepath.Ext(name))
							if ctype == "" {
//line /usr/local/go/src/net/http/fs.go:238
			_go_fuzz_dep_.CoverTab[37247]++
			// read a chunk to decide between utf-8 text and binary
			var buf [sniffLen]byte
			n, _ := io.ReadFull(content, buf[:])
			ctype = DetectContentType(buf[:n])
			_, err := content.Seek(0, io.SeekStart)
			if err != nil {
//line /usr/local/go/src/net/http/fs.go:244
				_go_fuzz_dep_.CoverTab[37248]++
									Error(w, "seeker can't seek", StatusInternalServerError)
									return
//line /usr/local/go/src/net/http/fs.go:246
				// _ = "end of CoverTab[37248]"
			} else {
//line /usr/local/go/src/net/http/fs.go:247
				_go_fuzz_dep_.CoverTab[37249]++
//line /usr/local/go/src/net/http/fs.go:247
				// _ = "end of CoverTab[37249]"
//line /usr/local/go/src/net/http/fs.go:247
			}
//line /usr/local/go/src/net/http/fs.go:247
			// _ = "end of CoverTab[37247]"
		} else {
//line /usr/local/go/src/net/http/fs.go:248
			_go_fuzz_dep_.CoverTab[37250]++
//line /usr/local/go/src/net/http/fs.go:248
			// _ = "end of CoverTab[37250]"
//line /usr/local/go/src/net/http/fs.go:248
		}
//line /usr/local/go/src/net/http/fs.go:248
		// _ = "end of CoverTab[37245]"
//line /usr/local/go/src/net/http/fs.go:248
		_go_fuzz_dep_.CoverTab[37246]++
							w.Header().Set("Content-Type", ctype)
//line /usr/local/go/src/net/http/fs.go:249
		// _ = "end of CoverTab[37246]"
	} else {
//line /usr/local/go/src/net/http/fs.go:250
		_go_fuzz_dep_.CoverTab[37251]++
//line /usr/local/go/src/net/http/fs.go:250
		if len(ctypes) > 0 {
//line /usr/local/go/src/net/http/fs.go:250
			_go_fuzz_dep_.CoverTab[37252]++
								ctype = ctypes[0]
//line /usr/local/go/src/net/http/fs.go:251
			// _ = "end of CoverTab[37252]"
		} else {
//line /usr/local/go/src/net/http/fs.go:252
			_go_fuzz_dep_.CoverTab[37253]++
//line /usr/local/go/src/net/http/fs.go:252
			// _ = "end of CoverTab[37253]"
//line /usr/local/go/src/net/http/fs.go:252
		}
//line /usr/local/go/src/net/http/fs.go:252
		// _ = "end of CoverTab[37251]"
//line /usr/local/go/src/net/http/fs.go:252
	}
//line /usr/local/go/src/net/http/fs.go:252
	// _ = "end of CoverTab[37235]"
//line /usr/local/go/src/net/http/fs.go:252
	_go_fuzz_dep_.CoverTab[37236]++

						size, err := sizeFunc()
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:255
		_go_fuzz_dep_.CoverTab[37254]++
							Error(w, err.Error(), StatusInternalServerError)
							return
//line /usr/local/go/src/net/http/fs.go:257
		// _ = "end of CoverTab[37254]"
	} else {
//line /usr/local/go/src/net/http/fs.go:258
		_go_fuzz_dep_.CoverTab[37255]++
//line /usr/local/go/src/net/http/fs.go:258
		// _ = "end of CoverTab[37255]"
//line /usr/local/go/src/net/http/fs.go:258
	}
//line /usr/local/go/src/net/http/fs.go:258
	// _ = "end of CoverTab[37236]"
//line /usr/local/go/src/net/http/fs.go:258
	_go_fuzz_dep_.CoverTab[37237]++
						if size < 0 {
//line /usr/local/go/src/net/http/fs.go:259
		_go_fuzz_dep_.CoverTab[37256]++

							Error(w, "negative content size computed", StatusInternalServerError)
							return
//line /usr/local/go/src/net/http/fs.go:262
		// _ = "end of CoverTab[37256]"
	} else {
//line /usr/local/go/src/net/http/fs.go:263
		_go_fuzz_dep_.CoverTab[37257]++
//line /usr/local/go/src/net/http/fs.go:263
		// _ = "end of CoverTab[37257]"
//line /usr/local/go/src/net/http/fs.go:263
	}
//line /usr/local/go/src/net/http/fs.go:263
	// _ = "end of CoverTab[37237]"
//line /usr/local/go/src/net/http/fs.go:263
	_go_fuzz_dep_.CoverTab[37238]++

//line /usr/local/go/src/net/http/fs.go:266
	sendSize := size
	var sendContent io.Reader = content
	ranges, err := parseRange(rangeReq, size)
	switch err {
	case nil:
//line /usr/local/go/src/net/http/fs.go:270
		_go_fuzz_dep_.CoverTab[37258]++
//line /usr/local/go/src/net/http/fs.go:270
		// _ = "end of CoverTab[37258]"
	case errNoOverlap:
//line /usr/local/go/src/net/http/fs.go:271
		_go_fuzz_dep_.CoverTab[37259]++
							if size == 0 {
//line /usr/local/go/src/net/http/fs.go:272
			_go_fuzz_dep_.CoverTab[37262]++

//line /usr/local/go/src/net/http/fs.go:277
			ranges = nil
								break
//line /usr/local/go/src/net/http/fs.go:278
			// _ = "end of CoverTab[37262]"
		} else {
//line /usr/local/go/src/net/http/fs.go:279
			_go_fuzz_dep_.CoverTab[37263]++
//line /usr/local/go/src/net/http/fs.go:279
			// _ = "end of CoverTab[37263]"
//line /usr/local/go/src/net/http/fs.go:279
		}
//line /usr/local/go/src/net/http/fs.go:279
		// _ = "end of CoverTab[37259]"
//line /usr/local/go/src/net/http/fs.go:279
		_go_fuzz_dep_.CoverTab[37260]++
							w.Header().Set("Content-Range", fmt.Sprintf("bytes */%d", size))
							fallthrough
//line /usr/local/go/src/net/http/fs.go:281
		// _ = "end of CoverTab[37260]"
	default:
//line /usr/local/go/src/net/http/fs.go:282
		_go_fuzz_dep_.CoverTab[37261]++
							Error(w, err.Error(), StatusRequestedRangeNotSatisfiable)
							return
//line /usr/local/go/src/net/http/fs.go:284
		// _ = "end of CoverTab[37261]"
	}
//line /usr/local/go/src/net/http/fs.go:285
	// _ = "end of CoverTab[37238]"
//line /usr/local/go/src/net/http/fs.go:285
	_go_fuzz_dep_.CoverTab[37239]++

						if sumRangesSize(ranges) > size {
//line /usr/local/go/src/net/http/fs.go:287
		_go_fuzz_dep_.CoverTab[37264]++

//line /usr/local/go/src/net/http/fs.go:292
		ranges = nil
//line /usr/local/go/src/net/http/fs.go:292
		// _ = "end of CoverTab[37264]"
	} else {
//line /usr/local/go/src/net/http/fs.go:293
		_go_fuzz_dep_.CoverTab[37265]++
//line /usr/local/go/src/net/http/fs.go:293
		// _ = "end of CoverTab[37265]"
//line /usr/local/go/src/net/http/fs.go:293
	}
//line /usr/local/go/src/net/http/fs.go:293
	// _ = "end of CoverTab[37239]"
//line /usr/local/go/src/net/http/fs.go:293
	_go_fuzz_dep_.CoverTab[37240]++
						switch {
	case len(ranges) == 1:
//line /usr/local/go/src/net/http/fs.go:295
		_go_fuzz_dep_.CoverTab[37266]++

//line /usr/local/go/src/net/http/fs.go:307
		ra := ranges[0]
		if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
//line /usr/local/go/src/net/http/fs.go:308
			_go_fuzz_dep_.CoverTab[37270]++
								Error(w, err.Error(), StatusRequestedRangeNotSatisfiable)
								return
//line /usr/local/go/src/net/http/fs.go:310
			// _ = "end of CoverTab[37270]"
		} else {
//line /usr/local/go/src/net/http/fs.go:311
			_go_fuzz_dep_.CoverTab[37271]++
//line /usr/local/go/src/net/http/fs.go:311
			// _ = "end of CoverTab[37271]"
//line /usr/local/go/src/net/http/fs.go:311
		}
//line /usr/local/go/src/net/http/fs.go:311
		// _ = "end of CoverTab[37266]"
//line /usr/local/go/src/net/http/fs.go:311
		_go_fuzz_dep_.CoverTab[37267]++
							sendSize = ra.length
							code = StatusPartialContent
							w.Header().Set("Content-Range", ra.contentRange(size))
//line /usr/local/go/src/net/http/fs.go:314
		// _ = "end of CoverTab[37267]"
	case len(ranges) > 1:
//line /usr/local/go/src/net/http/fs.go:315
		_go_fuzz_dep_.CoverTab[37268]++
							sendSize = rangesMIMESize(ranges, ctype, size)
							code = StatusPartialContent

							pr, pw := io.Pipe()
							mw := multipart.NewWriter(pw)
							w.Header().Set("Content-Type", "multipart/byteranges; boundary="+mw.Boundary())
							sendContent = pr
							defer pr.Close()
							go func() {
//line /usr/local/go/src/net/http/fs.go:324
			_go_fuzz_dep_.CoverTab[37272]++
								for _, ra := range ranges {
//line /usr/local/go/src/net/http/fs.go:325
				_go_fuzz_dep_.CoverTab[37274]++
									part, err := mw.CreatePart(ra.mimeHeader(ctype, size))
									if err != nil {
//line /usr/local/go/src/net/http/fs.go:327
					_go_fuzz_dep_.CoverTab[37277]++
										pw.CloseWithError(err)
										return
//line /usr/local/go/src/net/http/fs.go:329
					// _ = "end of CoverTab[37277]"
				} else {
//line /usr/local/go/src/net/http/fs.go:330
					_go_fuzz_dep_.CoverTab[37278]++
//line /usr/local/go/src/net/http/fs.go:330
					// _ = "end of CoverTab[37278]"
//line /usr/local/go/src/net/http/fs.go:330
				}
//line /usr/local/go/src/net/http/fs.go:330
				// _ = "end of CoverTab[37274]"
//line /usr/local/go/src/net/http/fs.go:330
				_go_fuzz_dep_.CoverTab[37275]++
									if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
//line /usr/local/go/src/net/http/fs.go:331
					_go_fuzz_dep_.CoverTab[37279]++
										pw.CloseWithError(err)
										return
//line /usr/local/go/src/net/http/fs.go:333
					// _ = "end of CoverTab[37279]"
				} else {
//line /usr/local/go/src/net/http/fs.go:334
					_go_fuzz_dep_.CoverTab[37280]++
//line /usr/local/go/src/net/http/fs.go:334
					// _ = "end of CoverTab[37280]"
//line /usr/local/go/src/net/http/fs.go:334
				}
//line /usr/local/go/src/net/http/fs.go:334
				// _ = "end of CoverTab[37275]"
//line /usr/local/go/src/net/http/fs.go:334
				_go_fuzz_dep_.CoverTab[37276]++
									if _, err := io.CopyN(part, content, ra.length); err != nil {
//line /usr/local/go/src/net/http/fs.go:335
					_go_fuzz_dep_.CoverTab[37281]++
										pw.CloseWithError(err)
										return
//line /usr/local/go/src/net/http/fs.go:337
					// _ = "end of CoverTab[37281]"
				} else {
//line /usr/local/go/src/net/http/fs.go:338
					_go_fuzz_dep_.CoverTab[37282]++
//line /usr/local/go/src/net/http/fs.go:338
					// _ = "end of CoverTab[37282]"
//line /usr/local/go/src/net/http/fs.go:338
				}
//line /usr/local/go/src/net/http/fs.go:338
				// _ = "end of CoverTab[37276]"
			}
//line /usr/local/go/src/net/http/fs.go:339
			// _ = "end of CoverTab[37272]"
//line /usr/local/go/src/net/http/fs.go:339
			_go_fuzz_dep_.CoverTab[37273]++
								mw.Close()
								pw.Close()
//line /usr/local/go/src/net/http/fs.go:341
			// _ = "end of CoverTab[37273]"
		}()
//line /usr/local/go/src/net/http/fs.go:342
		// _ = "end of CoverTab[37268]"
//line /usr/local/go/src/net/http/fs.go:342
	default:
//line /usr/local/go/src/net/http/fs.go:342
		_go_fuzz_dep_.CoverTab[37269]++
//line /usr/local/go/src/net/http/fs.go:342
		// _ = "end of CoverTab[37269]"
	}
//line /usr/local/go/src/net/http/fs.go:343
	// _ = "end of CoverTab[37240]"
//line /usr/local/go/src/net/http/fs.go:343
	_go_fuzz_dep_.CoverTab[37241]++

						w.Header().Set("Accept-Ranges", "bytes")
						if w.Header().Get("Content-Encoding") == "" {
//line /usr/local/go/src/net/http/fs.go:346
		_go_fuzz_dep_.CoverTab[37283]++
							w.Header().Set("Content-Length", strconv.FormatInt(sendSize, 10))
//line /usr/local/go/src/net/http/fs.go:347
		// _ = "end of CoverTab[37283]"
	} else {
//line /usr/local/go/src/net/http/fs.go:348
		_go_fuzz_dep_.CoverTab[37284]++
//line /usr/local/go/src/net/http/fs.go:348
		// _ = "end of CoverTab[37284]"
//line /usr/local/go/src/net/http/fs.go:348
	}
//line /usr/local/go/src/net/http/fs.go:348
	// _ = "end of CoverTab[37241]"
//line /usr/local/go/src/net/http/fs.go:348
	_go_fuzz_dep_.CoverTab[37242]++

						w.WriteHeader(code)

						if r.Method != "HEAD" {
//line /usr/local/go/src/net/http/fs.go:352
		_go_fuzz_dep_.CoverTab[37285]++
							io.CopyN(w, sendContent, sendSize)
//line /usr/local/go/src/net/http/fs.go:353
		// _ = "end of CoverTab[37285]"
	} else {
//line /usr/local/go/src/net/http/fs.go:354
		_go_fuzz_dep_.CoverTab[37286]++
//line /usr/local/go/src/net/http/fs.go:354
		// _ = "end of CoverTab[37286]"
//line /usr/local/go/src/net/http/fs.go:354
	}
//line /usr/local/go/src/net/http/fs.go:354
	// _ = "end of CoverTab[37242]"
}

// scanETag determines if a syntactically valid ETag is present at s. If so,
//line /usr/local/go/src/net/http/fs.go:357
// the ETag and remaining text after consuming ETag is returned. Otherwise,
//line /usr/local/go/src/net/http/fs.go:357
// it returns "", "".
//line /usr/local/go/src/net/http/fs.go:360
func scanETag(s string) (etag string, remain string) {
//line /usr/local/go/src/net/http/fs.go:360
	_go_fuzz_dep_.CoverTab[37287]++
						s = textproto.TrimString(s)
						start := 0
						if strings.HasPrefix(s, "W/") {
//line /usr/local/go/src/net/http/fs.go:363
		_go_fuzz_dep_.CoverTab[37291]++
							start = 2
//line /usr/local/go/src/net/http/fs.go:364
		// _ = "end of CoverTab[37291]"
	} else {
//line /usr/local/go/src/net/http/fs.go:365
		_go_fuzz_dep_.CoverTab[37292]++
//line /usr/local/go/src/net/http/fs.go:365
		// _ = "end of CoverTab[37292]"
//line /usr/local/go/src/net/http/fs.go:365
	}
//line /usr/local/go/src/net/http/fs.go:365
	// _ = "end of CoverTab[37287]"
//line /usr/local/go/src/net/http/fs.go:365
	_go_fuzz_dep_.CoverTab[37288]++
						if len(s[start:]) < 2 || func() bool {
//line /usr/local/go/src/net/http/fs.go:366
		_go_fuzz_dep_.CoverTab[37293]++
//line /usr/local/go/src/net/http/fs.go:366
		return s[start] != '"'
//line /usr/local/go/src/net/http/fs.go:366
		// _ = "end of CoverTab[37293]"
//line /usr/local/go/src/net/http/fs.go:366
	}() {
//line /usr/local/go/src/net/http/fs.go:366
		_go_fuzz_dep_.CoverTab[37294]++
							return "", ""
//line /usr/local/go/src/net/http/fs.go:367
		// _ = "end of CoverTab[37294]"
	} else {
//line /usr/local/go/src/net/http/fs.go:368
		_go_fuzz_dep_.CoverTab[37295]++
//line /usr/local/go/src/net/http/fs.go:368
		// _ = "end of CoverTab[37295]"
//line /usr/local/go/src/net/http/fs.go:368
	}
//line /usr/local/go/src/net/http/fs.go:368
	// _ = "end of CoverTab[37288]"
//line /usr/local/go/src/net/http/fs.go:368
	_go_fuzz_dep_.CoverTab[37289]++

//line /usr/local/go/src/net/http/fs.go:371
	for i := start + 1; i < len(s); i++ {
//line /usr/local/go/src/net/http/fs.go:371
		_go_fuzz_dep_.CoverTab[37296]++
							c := s[i]
							switch {

		case c == 0x21 || func() bool {
//line /usr/local/go/src/net/http/fs.go:375
			_go_fuzz_dep_.CoverTab[37300]++
//line /usr/local/go/src/net/http/fs.go:375
			return c >= 0x23 && func() bool {
//line /usr/local/go/src/net/http/fs.go:375
				_go_fuzz_dep_.CoverTab[37301]++
//line /usr/local/go/src/net/http/fs.go:375
				return c <= 0x7E
//line /usr/local/go/src/net/http/fs.go:375
				// _ = "end of CoverTab[37301]"
//line /usr/local/go/src/net/http/fs.go:375
			}()
//line /usr/local/go/src/net/http/fs.go:375
			// _ = "end of CoverTab[37300]"
//line /usr/local/go/src/net/http/fs.go:375
		}() || func() bool {
//line /usr/local/go/src/net/http/fs.go:375
			_go_fuzz_dep_.CoverTab[37302]++
//line /usr/local/go/src/net/http/fs.go:375
			return c >= 0x80
//line /usr/local/go/src/net/http/fs.go:375
			// _ = "end of CoverTab[37302]"
//line /usr/local/go/src/net/http/fs.go:375
		}():
//line /usr/local/go/src/net/http/fs.go:375
			_go_fuzz_dep_.CoverTab[37297]++
//line /usr/local/go/src/net/http/fs.go:375
			// _ = "end of CoverTab[37297]"
		case c == '"':
//line /usr/local/go/src/net/http/fs.go:376
			_go_fuzz_dep_.CoverTab[37298]++
								return s[:i+1], s[i+1:]
//line /usr/local/go/src/net/http/fs.go:377
			// _ = "end of CoverTab[37298]"
		default:
//line /usr/local/go/src/net/http/fs.go:378
			_go_fuzz_dep_.CoverTab[37299]++
								return "", ""
//line /usr/local/go/src/net/http/fs.go:379
			// _ = "end of CoverTab[37299]"
		}
//line /usr/local/go/src/net/http/fs.go:380
		// _ = "end of CoverTab[37296]"
	}
//line /usr/local/go/src/net/http/fs.go:381
	// _ = "end of CoverTab[37289]"
//line /usr/local/go/src/net/http/fs.go:381
	_go_fuzz_dep_.CoverTab[37290]++
						return "", ""
//line /usr/local/go/src/net/http/fs.go:382
	// _ = "end of CoverTab[37290]"
}

// etagStrongMatch reports whether a and b match using strong ETag comparison.
//line /usr/local/go/src/net/http/fs.go:385
// Assumes a and b are valid ETags.
//line /usr/local/go/src/net/http/fs.go:387
func etagStrongMatch(a, b string) bool {
//line /usr/local/go/src/net/http/fs.go:387
	_go_fuzz_dep_.CoverTab[37303]++
						return a == b && func() bool {
//line /usr/local/go/src/net/http/fs.go:388
		_go_fuzz_dep_.CoverTab[37304]++
//line /usr/local/go/src/net/http/fs.go:388
		return a != ""
//line /usr/local/go/src/net/http/fs.go:388
		// _ = "end of CoverTab[37304]"
//line /usr/local/go/src/net/http/fs.go:388
	}() && func() bool {
//line /usr/local/go/src/net/http/fs.go:388
		_go_fuzz_dep_.CoverTab[37305]++
//line /usr/local/go/src/net/http/fs.go:388
		return a[0] == '"'
//line /usr/local/go/src/net/http/fs.go:388
		// _ = "end of CoverTab[37305]"
//line /usr/local/go/src/net/http/fs.go:388
	}()
//line /usr/local/go/src/net/http/fs.go:388
	// _ = "end of CoverTab[37303]"
}

// etagWeakMatch reports whether a and b match using weak ETag comparison.
//line /usr/local/go/src/net/http/fs.go:391
// Assumes a and b are valid ETags.
//line /usr/local/go/src/net/http/fs.go:393
func etagWeakMatch(a, b string) bool {
//line /usr/local/go/src/net/http/fs.go:393
	_go_fuzz_dep_.CoverTab[37306]++
						return strings.TrimPrefix(a, "W/") == strings.TrimPrefix(b, "W/")
//line /usr/local/go/src/net/http/fs.go:394
	// _ = "end of CoverTab[37306]"
}

// condResult is the result of an HTTP request precondition check.
//line /usr/local/go/src/net/http/fs.go:397
// See https://tools.ietf.org/html/rfc7232 section 3.
//line /usr/local/go/src/net/http/fs.go:399
type condResult int

const (
	condNone	condResult	= iota
	condTrue
	condFalse
)

func checkIfMatch(w ResponseWriter, r *Request) condResult {
//line /usr/local/go/src/net/http/fs.go:407
	_go_fuzz_dep_.CoverTab[37307]++
						im := r.Header.Get("If-Match")
						if im == "" {
//line /usr/local/go/src/net/http/fs.go:409
		_go_fuzz_dep_.CoverTab[37310]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:410
		// _ = "end of CoverTab[37310]"
	} else {
//line /usr/local/go/src/net/http/fs.go:411
		_go_fuzz_dep_.CoverTab[37311]++
//line /usr/local/go/src/net/http/fs.go:411
		// _ = "end of CoverTab[37311]"
//line /usr/local/go/src/net/http/fs.go:411
	}
//line /usr/local/go/src/net/http/fs.go:411
	// _ = "end of CoverTab[37307]"
//line /usr/local/go/src/net/http/fs.go:411
	_go_fuzz_dep_.CoverTab[37308]++
						for {
//line /usr/local/go/src/net/http/fs.go:412
		_go_fuzz_dep_.CoverTab[37312]++
							im = textproto.TrimString(im)
							if len(im) == 0 {
//line /usr/local/go/src/net/http/fs.go:414
			_go_fuzz_dep_.CoverTab[37318]++
								break
//line /usr/local/go/src/net/http/fs.go:415
			// _ = "end of CoverTab[37318]"
		} else {
//line /usr/local/go/src/net/http/fs.go:416
			_go_fuzz_dep_.CoverTab[37319]++
//line /usr/local/go/src/net/http/fs.go:416
			// _ = "end of CoverTab[37319]"
//line /usr/local/go/src/net/http/fs.go:416
		}
//line /usr/local/go/src/net/http/fs.go:416
		// _ = "end of CoverTab[37312]"
//line /usr/local/go/src/net/http/fs.go:416
		_go_fuzz_dep_.CoverTab[37313]++
							if im[0] == ',' {
//line /usr/local/go/src/net/http/fs.go:417
			_go_fuzz_dep_.CoverTab[37320]++
								im = im[1:]
								continue
//line /usr/local/go/src/net/http/fs.go:419
			// _ = "end of CoverTab[37320]"
		} else {
//line /usr/local/go/src/net/http/fs.go:420
			_go_fuzz_dep_.CoverTab[37321]++
//line /usr/local/go/src/net/http/fs.go:420
			// _ = "end of CoverTab[37321]"
//line /usr/local/go/src/net/http/fs.go:420
		}
//line /usr/local/go/src/net/http/fs.go:420
		// _ = "end of CoverTab[37313]"
//line /usr/local/go/src/net/http/fs.go:420
		_go_fuzz_dep_.CoverTab[37314]++
							if im[0] == '*' {
//line /usr/local/go/src/net/http/fs.go:421
			_go_fuzz_dep_.CoverTab[37322]++
								return condTrue
//line /usr/local/go/src/net/http/fs.go:422
			// _ = "end of CoverTab[37322]"
		} else {
//line /usr/local/go/src/net/http/fs.go:423
			_go_fuzz_dep_.CoverTab[37323]++
//line /usr/local/go/src/net/http/fs.go:423
			// _ = "end of CoverTab[37323]"
//line /usr/local/go/src/net/http/fs.go:423
		}
//line /usr/local/go/src/net/http/fs.go:423
		// _ = "end of CoverTab[37314]"
//line /usr/local/go/src/net/http/fs.go:423
		_go_fuzz_dep_.CoverTab[37315]++
							etag, remain := scanETag(im)
							if etag == "" {
//line /usr/local/go/src/net/http/fs.go:425
			_go_fuzz_dep_.CoverTab[37324]++
								break
//line /usr/local/go/src/net/http/fs.go:426
			// _ = "end of CoverTab[37324]"
		} else {
//line /usr/local/go/src/net/http/fs.go:427
			_go_fuzz_dep_.CoverTab[37325]++
//line /usr/local/go/src/net/http/fs.go:427
			// _ = "end of CoverTab[37325]"
//line /usr/local/go/src/net/http/fs.go:427
		}
//line /usr/local/go/src/net/http/fs.go:427
		// _ = "end of CoverTab[37315]"
//line /usr/local/go/src/net/http/fs.go:427
		_go_fuzz_dep_.CoverTab[37316]++
							if etagStrongMatch(etag, w.Header().get("Etag")) {
//line /usr/local/go/src/net/http/fs.go:428
			_go_fuzz_dep_.CoverTab[37326]++
								return condTrue
//line /usr/local/go/src/net/http/fs.go:429
			// _ = "end of CoverTab[37326]"
		} else {
//line /usr/local/go/src/net/http/fs.go:430
			_go_fuzz_dep_.CoverTab[37327]++
//line /usr/local/go/src/net/http/fs.go:430
			// _ = "end of CoverTab[37327]"
//line /usr/local/go/src/net/http/fs.go:430
		}
//line /usr/local/go/src/net/http/fs.go:430
		// _ = "end of CoverTab[37316]"
//line /usr/local/go/src/net/http/fs.go:430
		_go_fuzz_dep_.CoverTab[37317]++
							im = remain
//line /usr/local/go/src/net/http/fs.go:431
		// _ = "end of CoverTab[37317]"
	}
//line /usr/local/go/src/net/http/fs.go:432
	// _ = "end of CoverTab[37308]"
//line /usr/local/go/src/net/http/fs.go:432
	_go_fuzz_dep_.CoverTab[37309]++

						return condFalse
//line /usr/local/go/src/net/http/fs.go:434
	// _ = "end of CoverTab[37309]"
}

func checkIfUnmodifiedSince(r *Request, modtime time.Time) condResult {
//line /usr/local/go/src/net/http/fs.go:437
	_go_fuzz_dep_.CoverTab[37328]++
						ius := r.Header.Get("If-Unmodified-Since")
						if ius == "" || func() bool {
//line /usr/local/go/src/net/http/fs.go:439
		_go_fuzz_dep_.CoverTab[37332]++
//line /usr/local/go/src/net/http/fs.go:439
		return isZeroTime(modtime)
//line /usr/local/go/src/net/http/fs.go:439
		// _ = "end of CoverTab[37332]"
//line /usr/local/go/src/net/http/fs.go:439
	}() {
//line /usr/local/go/src/net/http/fs.go:439
		_go_fuzz_dep_.CoverTab[37333]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:440
		// _ = "end of CoverTab[37333]"
	} else {
//line /usr/local/go/src/net/http/fs.go:441
		_go_fuzz_dep_.CoverTab[37334]++
//line /usr/local/go/src/net/http/fs.go:441
		// _ = "end of CoverTab[37334]"
//line /usr/local/go/src/net/http/fs.go:441
	}
//line /usr/local/go/src/net/http/fs.go:441
	// _ = "end of CoverTab[37328]"
//line /usr/local/go/src/net/http/fs.go:441
	_go_fuzz_dep_.CoverTab[37329]++
						t, err := ParseTime(ius)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:443
		_go_fuzz_dep_.CoverTab[37335]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:444
		// _ = "end of CoverTab[37335]"
	} else {
//line /usr/local/go/src/net/http/fs.go:445
		_go_fuzz_dep_.CoverTab[37336]++
//line /usr/local/go/src/net/http/fs.go:445
		// _ = "end of CoverTab[37336]"
//line /usr/local/go/src/net/http/fs.go:445
	}
//line /usr/local/go/src/net/http/fs.go:445
	// _ = "end of CoverTab[37329]"
//line /usr/local/go/src/net/http/fs.go:445
	_go_fuzz_dep_.CoverTab[37330]++

//line /usr/local/go/src/net/http/fs.go:449
	modtime = modtime.Truncate(time.Second)
	if ret := modtime.Compare(t); ret <= 0 {
//line /usr/local/go/src/net/http/fs.go:450
		_go_fuzz_dep_.CoverTab[37337]++
							return condTrue
//line /usr/local/go/src/net/http/fs.go:451
		// _ = "end of CoverTab[37337]"
	} else {
//line /usr/local/go/src/net/http/fs.go:452
		_go_fuzz_dep_.CoverTab[37338]++
//line /usr/local/go/src/net/http/fs.go:452
		// _ = "end of CoverTab[37338]"
//line /usr/local/go/src/net/http/fs.go:452
	}
//line /usr/local/go/src/net/http/fs.go:452
	// _ = "end of CoverTab[37330]"
//line /usr/local/go/src/net/http/fs.go:452
	_go_fuzz_dep_.CoverTab[37331]++
						return condFalse
//line /usr/local/go/src/net/http/fs.go:453
	// _ = "end of CoverTab[37331]"
}

func checkIfNoneMatch(w ResponseWriter, r *Request) condResult {
//line /usr/local/go/src/net/http/fs.go:456
	_go_fuzz_dep_.CoverTab[37339]++
						inm := r.Header.get("If-None-Match")
						if inm == "" {
//line /usr/local/go/src/net/http/fs.go:458
		_go_fuzz_dep_.CoverTab[37342]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:459
		// _ = "end of CoverTab[37342]"
	} else {
//line /usr/local/go/src/net/http/fs.go:460
		_go_fuzz_dep_.CoverTab[37343]++
//line /usr/local/go/src/net/http/fs.go:460
		// _ = "end of CoverTab[37343]"
//line /usr/local/go/src/net/http/fs.go:460
	}
//line /usr/local/go/src/net/http/fs.go:460
	// _ = "end of CoverTab[37339]"
//line /usr/local/go/src/net/http/fs.go:460
	_go_fuzz_dep_.CoverTab[37340]++
						buf := inm
						for {
//line /usr/local/go/src/net/http/fs.go:462
		_go_fuzz_dep_.CoverTab[37344]++
							buf = textproto.TrimString(buf)
							if len(buf) == 0 {
//line /usr/local/go/src/net/http/fs.go:464
			_go_fuzz_dep_.CoverTab[37350]++
								break
//line /usr/local/go/src/net/http/fs.go:465
			// _ = "end of CoverTab[37350]"
		} else {
//line /usr/local/go/src/net/http/fs.go:466
			_go_fuzz_dep_.CoverTab[37351]++
//line /usr/local/go/src/net/http/fs.go:466
			// _ = "end of CoverTab[37351]"
//line /usr/local/go/src/net/http/fs.go:466
		}
//line /usr/local/go/src/net/http/fs.go:466
		// _ = "end of CoverTab[37344]"
//line /usr/local/go/src/net/http/fs.go:466
		_go_fuzz_dep_.CoverTab[37345]++
							if buf[0] == ',' {
//line /usr/local/go/src/net/http/fs.go:467
			_go_fuzz_dep_.CoverTab[37352]++
								buf = buf[1:]
								continue
//line /usr/local/go/src/net/http/fs.go:469
			// _ = "end of CoverTab[37352]"
		} else {
//line /usr/local/go/src/net/http/fs.go:470
			_go_fuzz_dep_.CoverTab[37353]++
//line /usr/local/go/src/net/http/fs.go:470
			// _ = "end of CoverTab[37353]"
//line /usr/local/go/src/net/http/fs.go:470
		}
//line /usr/local/go/src/net/http/fs.go:470
		// _ = "end of CoverTab[37345]"
//line /usr/local/go/src/net/http/fs.go:470
		_go_fuzz_dep_.CoverTab[37346]++
							if buf[0] == '*' {
//line /usr/local/go/src/net/http/fs.go:471
			_go_fuzz_dep_.CoverTab[37354]++
								return condFalse
//line /usr/local/go/src/net/http/fs.go:472
			// _ = "end of CoverTab[37354]"
		} else {
//line /usr/local/go/src/net/http/fs.go:473
			_go_fuzz_dep_.CoverTab[37355]++
//line /usr/local/go/src/net/http/fs.go:473
			// _ = "end of CoverTab[37355]"
//line /usr/local/go/src/net/http/fs.go:473
		}
//line /usr/local/go/src/net/http/fs.go:473
		// _ = "end of CoverTab[37346]"
//line /usr/local/go/src/net/http/fs.go:473
		_go_fuzz_dep_.CoverTab[37347]++
							etag, remain := scanETag(buf)
							if etag == "" {
//line /usr/local/go/src/net/http/fs.go:475
			_go_fuzz_dep_.CoverTab[37356]++
								break
//line /usr/local/go/src/net/http/fs.go:476
			// _ = "end of CoverTab[37356]"
		} else {
//line /usr/local/go/src/net/http/fs.go:477
			_go_fuzz_dep_.CoverTab[37357]++
//line /usr/local/go/src/net/http/fs.go:477
			// _ = "end of CoverTab[37357]"
//line /usr/local/go/src/net/http/fs.go:477
		}
//line /usr/local/go/src/net/http/fs.go:477
		// _ = "end of CoverTab[37347]"
//line /usr/local/go/src/net/http/fs.go:477
		_go_fuzz_dep_.CoverTab[37348]++
							if etagWeakMatch(etag, w.Header().get("Etag")) {
//line /usr/local/go/src/net/http/fs.go:478
			_go_fuzz_dep_.CoverTab[37358]++
								return condFalse
//line /usr/local/go/src/net/http/fs.go:479
			// _ = "end of CoverTab[37358]"
		} else {
//line /usr/local/go/src/net/http/fs.go:480
			_go_fuzz_dep_.CoverTab[37359]++
//line /usr/local/go/src/net/http/fs.go:480
			// _ = "end of CoverTab[37359]"
//line /usr/local/go/src/net/http/fs.go:480
		}
//line /usr/local/go/src/net/http/fs.go:480
		// _ = "end of CoverTab[37348]"
//line /usr/local/go/src/net/http/fs.go:480
		_go_fuzz_dep_.CoverTab[37349]++
							buf = remain
//line /usr/local/go/src/net/http/fs.go:481
		// _ = "end of CoverTab[37349]"
	}
//line /usr/local/go/src/net/http/fs.go:482
	// _ = "end of CoverTab[37340]"
//line /usr/local/go/src/net/http/fs.go:482
	_go_fuzz_dep_.CoverTab[37341]++
						return condTrue
//line /usr/local/go/src/net/http/fs.go:483
	// _ = "end of CoverTab[37341]"
}

func checkIfModifiedSince(r *Request, modtime time.Time) condResult {
//line /usr/local/go/src/net/http/fs.go:486
	_go_fuzz_dep_.CoverTab[37360]++
						if r.Method != "GET" && func() bool {
//line /usr/local/go/src/net/http/fs.go:487
		_go_fuzz_dep_.CoverTab[37365]++
//line /usr/local/go/src/net/http/fs.go:487
		return r.Method != "HEAD"
//line /usr/local/go/src/net/http/fs.go:487
		// _ = "end of CoverTab[37365]"
//line /usr/local/go/src/net/http/fs.go:487
	}() {
//line /usr/local/go/src/net/http/fs.go:487
		_go_fuzz_dep_.CoverTab[37366]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:488
		// _ = "end of CoverTab[37366]"
	} else {
//line /usr/local/go/src/net/http/fs.go:489
		_go_fuzz_dep_.CoverTab[37367]++
//line /usr/local/go/src/net/http/fs.go:489
		// _ = "end of CoverTab[37367]"
//line /usr/local/go/src/net/http/fs.go:489
	}
//line /usr/local/go/src/net/http/fs.go:489
	// _ = "end of CoverTab[37360]"
//line /usr/local/go/src/net/http/fs.go:489
	_go_fuzz_dep_.CoverTab[37361]++
						ims := r.Header.Get("If-Modified-Since")
						if ims == "" || func() bool {
//line /usr/local/go/src/net/http/fs.go:491
		_go_fuzz_dep_.CoverTab[37368]++
//line /usr/local/go/src/net/http/fs.go:491
		return isZeroTime(modtime)
//line /usr/local/go/src/net/http/fs.go:491
		// _ = "end of CoverTab[37368]"
//line /usr/local/go/src/net/http/fs.go:491
	}() {
//line /usr/local/go/src/net/http/fs.go:491
		_go_fuzz_dep_.CoverTab[37369]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:492
		// _ = "end of CoverTab[37369]"
	} else {
//line /usr/local/go/src/net/http/fs.go:493
		_go_fuzz_dep_.CoverTab[37370]++
//line /usr/local/go/src/net/http/fs.go:493
		// _ = "end of CoverTab[37370]"
//line /usr/local/go/src/net/http/fs.go:493
	}
//line /usr/local/go/src/net/http/fs.go:493
	// _ = "end of CoverTab[37361]"
//line /usr/local/go/src/net/http/fs.go:493
	_go_fuzz_dep_.CoverTab[37362]++
						t, err := ParseTime(ims)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:495
		_go_fuzz_dep_.CoverTab[37371]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:496
		// _ = "end of CoverTab[37371]"
	} else {
//line /usr/local/go/src/net/http/fs.go:497
		_go_fuzz_dep_.CoverTab[37372]++
//line /usr/local/go/src/net/http/fs.go:497
		// _ = "end of CoverTab[37372]"
//line /usr/local/go/src/net/http/fs.go:497
	}
//line /usr/local/go/src/net/http/fs.go:497
	// _ = "end of CoverTab[37362]"
//line /usr/local/go/src/net/http/fs.go:497
	_go_fuzz_dep_.CoverTab[37363]++

//line /usr/local/go/src/net/http/fs.go:500
	modtime = modtime.Truncate(time.Second)
	if ret := modtime.Compare(t); ret <= 0 {
//line /usr/local/go/src/net/http/fs.go:501
		_go_fuzz_dep_.CoverTab[37373]++
							return condFalse
//line /usr/local/go/src/net/http/fs.go:502
		// _ = "end of CoverTab[37373]"
	} else {
//line /usr/local/go/src/net/http/fs.go:503
		_go_fuzz_dep_.CoverTab[37374]++
//line /usr/local/go/src/net/http/fs.go:503
		// _ = "end of CoverTab[37374]"
//line /usr/local/go/src/net/http/fs.go:503
	}
//line /usr/local/go/src/net/http/fs.go:503
	// _ = "end of CoverTab[37363]"
//line /usr/local/go/src/net/http/fs.go:503
	_go_fuzz_dep_.CoverTab[37364]++
						return condTrue
//line /usr/local/go/src/net/http/fs.go:504
	// _ = "end of CoverTab[37364]"
}

func checkIfRange(w ResponseWriter, r *Request, modtime time.Time) condResult {
//line /usr/local/go/src/net/http/fs.go:507
	_go_fuzz_dep_.CoverTab[37375]++
						if r.Method != "GET" && func() bool {
//line /usr/local/go/src/net/http/fs.go:508
		_go_fuzz_dep_.CoverTab[37382]++
//line /usr/local/go/src/net/http/fs.go:508
		return r.Method != "HEAD"
//line /usr/local/go/src/net/http/fs.go:508
		// _ = "end of CoverTab[37382]"
//line /usr/local/go/src/net/http/fs.go:508
	}() {
//line /usr/local/go/src/net/http/fs.go:508
		_go_fuzz_dep_.CoverTab[37383]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:509
		// _ = "end of CoverTab[37383]"
	} else {
//line /usr/local/go/src/net/http/fs.go:510
		_go_fuzz_dep_.CoverTab[37384]++
//line /usr/local/go/src/net/http/fs.go:510
		// _ = "end of CoverTab[37384]"
//line /usr/local/go/src/net/http/fs.go:510
	}
//line /usr/local/go/src/net/http/fs.go:510
	// _ = "end of CoverTab[37375]"
//line /usr/local/go/src/net/http/fs.go:510
	_go_fuzz_dep_.CoverTab[37376]++
						ir := r.Header.get("If-Range")
						if ir == "" {
//line /usr/local/go/src/net/http/fs.go:512
		_go_fuzz_dep_.CoverTab[37385]++
							return condNone
//line /usr/local/go/src/net/http/fs.go:513
		// _ = "end of CoverTab[37385]"
	} else {
//line /usr/local/go/src/net/http/fs.go:514
		_go_fuzz_dep_.CoverTab[37386]++
//line /usr/local/go/src/net/http/fs.go:514
		// _ = "end of CoverTab[37386]"
//line /usr/local/go/src/net/http/fs.go:514
	}
//line /usr/local/go/src/net/http/fs.go:514
	// _ = "end of CoverTab[37376]"
//line /usr/local/go/src/net/http/fs.go:514
	_go_fuzz_dep_.CoverTab[37377]++
						etag, _ := scanETag(ir)
						if etag != "" {
//line /usr/local/go/src/net/http/fs.go:516
		_go_fuzz_dep_.CoverTab[37387]++
							if etagStrongMatch(etag, w.Header().Get("Etag")) {
//line /usr/local/go/src/net/http/fs.go:517
			_go_fuzz_dep_.CoverTab[37388]++
								return condTrue
//line /usr/local/go/src/net/http/fs.go:518
			// _ = "end of CoverTab[37388]"
		} else {
//line /usr/local/go/src/net/http/fs.go:519
			_go_fuzz_dep_.CoverTab[37389]++
								return condFalse
//line /usr/local/go/src/net/http/fs.go:520
			// _ = "end of CoverTab[37389]"
		}
//line /usr/local/go/src/net/http/fs.go:521
		// _ = "end of CoverTab[37387]"
	} else {
//line /usr/local/go/src/net/http/fs.go:522
		_go_fuzz_dep_.CoverTab[37390]++
//line /usr/local/go/src/net/http/fs.go:522
		// _ = "end of CoverTab[37390]"
//line /usr/local/go/src/net/http/fs.go:522
	}
//line /usr/local/go/src/net/http/fs.go:522
	// _ = "end of CoverTab[37377]"
//line /usr/local/go/src/net/http/fs.go:522
	_go_fuzz_dep_.CoverTab[37378]++

//line /usr/local/go/src/net/http/fs.go:525
	if modtime.IsZero() {
//line /usr/local/go/src/net/http/fs.go:525
		_go_fuzz_dep_.CoverTab[37391]++
							return condFalse
//line /usr/local/go/src/net/http/fs.go:526
		// _ = "end of CoverTab[37391]"
	} else {
//line /usr/local/go/src/net/http/fs.go:527
		_go_fuzz_dep_.CoverTab[37392]++
//line /usr/local/go/src/net/http/fs.go:527
		// _ = "end of CoverTab[37392]"
//line /usr/local/go/src/net/http/fs.go:527
	}
//line /usr/local/go/src/net/http/fs.go:527
	// _ = "end of CoverTab[37378]"
//line /usr/local/go/src/net/http/fs.go:527
	_go_fuzz_dep_.CoverTab[37379]++
						t, err := ParseTime(ir)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:529
		_go_fuzz_dep_.CoverTab[37393]++
							return condFalse
//line /usr/local/go/src/net/http/fs.go:530
		// _ = "end of CoverTab[37393]"
	} else {
//line /usr/local/go/src/net/http/fs.go:531
		_go_fuzz_dep_.CoverTab[37394]++
//line /usr/local/go/src/net/http/fs.go:531
		// _ = "end of CoverTab[37394]"
//line /usr/local/go/src/net/http/fs.go:531
	}
//line /usr/local/go/src/net/http/fs.go:531
	// _ = "end of CoverTab[37379]"
//line /usr/local/go/src/net/http/fs.go:531
	_go_fuzz_dep_.CoverTab[37380]++
						if t.Unix() == modtime.Unix() {
//line /usr/local/go/src/net/http/fs.go:532
		_go_fuzz_dep_.CoverTab[37395]++
							return condTrue
//line /usr/local/go/src/net/http/fs.go:533
		// _ = "end of CoverTab[37395]"
	} else {
//line /usr/local/go/src/net/http/fs.go:534
		_go_fuzz_dep_.CoverTab[37396]++
//line /usr/local/go/src/net/http/fs.go:534
		// _ = "end of CoverTab[37396]"
//line /usr/local/go/src/net/http/fs.go:534
	}
//line /usr/local/go/src/net/http/fs.go:534
	// _ = "end of CoverTab[37380]"
//line /usr/local/go/src/net/http/fs.go:534
	_go_fuzz_dep_.CoverTab[37381]++
						return condFalse
//line /usr/local/go/src/net/http/fs.go:535
	// _ = "end of CoverTab[37381]"
}

var unixEpochTime = time.Unix(0, 0)

// isZeroTime reports whether t is obviously unspecified (either zero or Unix()=0).
func isZeroTime(t time.Time) bool {
//line /usr/local/go/src/net/http/fs.go:541
	_go_fuzz_dep_.CoverTab[37397]++
						return t.IsZero() || func() bool {
//line /usr/local/go/src/net/http/fs.go:542
		_go_fuzz_dep_.CoverTab[37398]++
//line /usr/local/go/src/net/http/fs.go:542
		return t.Equal(unixEpochTime)
//line /usr/local/go/src/net/http/fs.go:542
		// _ = "end of CoverTab[37398]"
//line /usr/local/go/src/net/http/fs.go:542
	}()
//line /usr/local/go/src/net/http/fs.go:542
	// _ = "end of CoverTab[37397]"
}

func setLastModified(w ResponseWriter, modtime time.Time) {
//line /usr/local/go/src/net/http/fs.go:545
	_go_fuzz_dep_.CoverTab[37399]++
						if !isZeroTime(modtime) {
//line /usr/local/go/src/net/http/fs.go:546
		_go_fuzz_dep_.CoverTab[37400]++
							w.Header().Set("Last-Modified", modtime.UTC().Format(TimeFormat))
//line /usr/local/go/src/net/http/fs.go:547
		// _ = "end of CoverTab[37400]"
	} else {
//line /usr/local/go/src/net/http/fs.go:548
		_go_fuzz_dep_.CoverTab[37401]++
//line /usr/local/go/src/net/http/fs.go:548
		// _ = "end of CoverTab[37401]"
//line /usr/local/go/src/net/http/fs.go:548
	}
//line /usr/local/go/src/net/http/fs.go:548
	// _ = "end of CoverTab[37399]"
}

func writeNotModified(w ResponseWriter) {
//line /usr/local/go/src/net/http/fs.go:551
	_go_fuzz_dep_.CoverTab[37402]++

//line /usr/local/go/src/net/http/fs.go:557
	h := w.Header()
	delete(h, "Content-Type")
	delete(h, "Content-Length")
	delete(h, "Content-Encoding")
	if h.Get("Etag") != "" {
//line /usr/local/go/src/net/http/fs.go:561
		_go_fuzz_dep_.CoverTab[37404]++
							delete(h, "Last-Modified")
//line /usr/local/go/src/net/http/fs.go:562
		// _ = "end of CoverTab[37404]"
	} else {
//line /usr/local/go/src/net/http/fs.go:563
		_go_fuzz_dep_.CoverTab[37405]++
//line /usr/local/go/src/net/http/fs.go:563
		// _ = "end of CoverTab[37405]"
//line /usr/local/go/src/net/http/fs.go:563
	}
//line /usr/local/go/src/net/http/fs.go:563
	// _ = "end of CoverTab[37402]"
//line /usr/local/go/src/net/http/fs.go:563
	_go_fuzz_dep_.CoverTab[37403]++
						w.WriteHeader(StatusNotModified)
//line /usr/local/go/src/net/http/fs.go:564
	// _ = "end of CoverTab[37403]"
}

// checkPreconditions evaluates request preconditions and reports whether a precondition
//line /usr/local/go/src/net/http/fs.go:567
// resulted in sending StatusNotModified or StatusPreconditionFailed.
//line /usr/local/go/src/net/http/fs.go:569
func checkPreconditions(w ResponseWriter, r *Request, modtime time.Time) (done bool, rangeHeader string) {
//line /usr/local/go/src/net/http/fs.go:569
	_go_fuzz_dep_.CoverTab[37406]++

						ch := checkIfMatch(w, r)
						if ch == condNone {
//line /usr/local/go/src/net/http/fs.go:572
		_go_fuzz_dep_.CoverTab[37411]++
							ch = checkIfUnmodifiedSince(r, modtime)
//line /usr/local/go/src/net/http/fs.go:573
		// _ = "end of CoverTab[37411]"
	} else {
//line /usr/local/go/src/net/http/fs.go:574
		_go_fuzz_dep_.CoverTab[37412]++
//line /usr/local/go/src/net/http/fs.go:574
		// _ = "end of CoverTab[37412]"
//line /usr/local/go/src/net/http/fs.go:574
	}
//line /usr/local/go/src/net/http/fs.go:574
	// _ = "end of CoverTab[37406]"
//line /usr/local/go/src/net/http/fs.go:574
	_go_fuzz_dep_.CoverTab[37407]++
						if ch == condFalse {
//line /usr/local/go/src/net/http/fs.go:575
		_go_fuzz_dep_.CoverTab[37413]++
							w.WriteHeader(StatusPreconditionFailed)
							return true, ""
//line /usr/local/go/src/net/http/fs.go:577
		// _ = "end of CoverTab[37413]"
	} else {
//line /usr/local/go/src/net/http/fs.go:578
		_go_fuzz_dep_.CoverTab[37414]++
//line /usr/local/go/src/net/http/fs.go:578
		// _ = "end of CoverTab[37414]"
//line /usr/local/go/src/net/http/fs.go:578
	}
//line /usr/local/go/src/net/http/fs.go:578
	// _ = "end of CoverTab[37407]"
//line /usr/local/go/src/net/http/fs.go:578
	_go_fuzz_dep_.CoverTab[37408]++
						switch checkIfNoneMatch(w, r) {
	case condFalse:
//line /usr/local/go/src/net/http/fs.go:580
		_go_fuzz_dep_.CoverTab[37415]++
							if r.Method == "GET" || func() bool {
//line /usr/local/go/src/net/http/fs.go:581
			_go_fuzz_dep_.CoverTab[37418]++
//line /usr/local/go/src/net/http/fs.go:581
			return r.Method == "HEAD"
//line /usr/local/go/src/net/http/fs.go:581
			// _ = "end of CoverTab[37418]"
//line /usr/local/go/src/net/http/fs.go:581
		}() {
//line /usr/local/go/src/net/http/fs.go:581
			_go_fuzz_dep_.CoverTab[37419]++
								writeNotModified(w)
								return true, ""
//line /usr/local/go/src/net/http/fs.go:583
			// _ = "end of CoverTab[37419]"
		} else {
//line /usr/local/go/src/net/http/fs.go:584
			_go_fuzz_dep_.CoverTab[37420]++
								w.WriteHeader(StatusPreconditionFailed)
								return true, ""
//line /usr/local/go/src/net/http/fs.go:586
			// _ = "end of CoverTab[37420]"
		}
//line /usr/local/go/src/net/http/fs.go:587
		// _ = "end of CoverTab[37415]"
	case condNone:
//line /usr/local/go/src/net/http/fs.go:588
		_go_fuzz_dep_.CoverTab[37416]++
							if checkIfModifiedSince(r, modtime) == condFalse {
//line /usr/local/go/src/net/http/fs.go:589
			_go_fuzz_dep_.CoverTab[37421]++
								writeNotModified(w)
								return true, ""
//line /usr/local/go/src/net/http/fs.go:591
			// _ = "end of CoverTab[37421]"
		} else {
//line /usr/local/go/src/net/http/fs.go:592
			_go_fuzz_dep_.CoverTab[37422]++
//line /usr/local/go/src/net/http/fs.go:592
			// _ = "end of CoverTab[37422]"
//line /usr/local/go/src/net/http/fs.go:592
		}
//line /usr/local/go/src/net/http/fs.go:592
		// _ = "end of CoverTab[37416]"
//line /usr/local/go/src/net/http/fs.go:592
	default:
//line /usr/local/go/src/net/http/fs.go:592
		_go_fuzz_dep_.CoverTab[37417]++
//line /usr/local/go/src/net/http/fs.go:592
		// _ = "end of CoverTab[37417]"
	}
//line /usr/local/go/src/net/http/fs.go:593
	// _ = "end of CoverTab[37408]"
//line /usr/local/go/src/net/http/fs.go:593
	_go_fuzz_dep_.CoverTab[37409]++

						rangeHeader = r.Header.get("Range")
						if rangeHeader != "" && func() bool {
//line /usr/local/go/src/net/http/fs.go:596
		_go_fuzz_dep_.CoverTab[37423]++
//line /usr/local/go/src/net/http/fs.go:596
		return checkIfRange(w, r, modtime) == condFalse
//line /usr/local/go/src/net/http/fs.go:596
		// _ = "end of CoverTab[37423]"
//line /usr/local/go/src/net/http/fs.go:596
	}() {
//line /usr/local/go/src/net/http/fs.go:596
		_go_fuzz_dep_.CoverTab[37424]++
							rangeHeader = ""
//line /usr/local/go/src/net/http/fs.go:597
		// _ = "end of CoverTab[37424]"
	} else {
//line /usr/local/go/src/net/http/fs.go:598
		_go_fuzz_dep_.CoverTab[37425]++
//line /usr/local/go/src/net/http/fs.go:598
		// _ = "end of CoverTab[37425]"
//line /usr/local/go/src/net/http/fs.go:598
	}
//line /usr/local/go/src/net/http/fs.go:598
	// _ = "end of CoverTab[37409]"
//line /usr/local/go/src/net/http/fs.go:598
	_go_fuzz_dep_.CoverTab[37410]++
						return false, rangeHeader
//line /usr/local/go/src/net/http/fs.go:599
	// _ = "end of CoverTab[37410]"
}

// name is '/'-separated, not filepath.Separator.
func serveFile(w ResponseWriter, r *Request, fs FileSystem, name string, redirect bool) {
//line /usr/local/go/src/net/http/fs.go:603
	_go_fuzz_dep_.CoverTab[37426]++
						const indexPage = "/index.html"

//line /usr/local/go/src/net/http/fs.go:609
	if strings.HasSuffix(r.URL.Path, indexPage) {
//line /usr/local/go/src/net/http/fs.go:609
		_go_fuzz_dep_.CoverTab[37434]++
							localRedirect(w, r, "./")
							return
//line /usr/local/go/src/net/http/fs.go:611
		// _ = "end of CoverTab[37434]"
	} else {
//line /usr/local/go/src/net/http/fs.go:612
		_go_fuzz_dep_.CoverTab[37435]++
//line /usr/local/go/src/net/http/fs.go:612
		// _ = "end of CoverTab[37435]"
//line /usr/local/go/src/net/http/fs.go:612
	}
//line /usr/local/go/src/net/http/fs.go:612
	// _ = "end of CoverTab[37426]"
//line /usr/local/go/src/net/http/fs.go:612
	_go_fuzz_dep_.CoverTab[37427]++

						f, err := fs.Open(name)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:615
		_go_fuzz_dep_.CoverTab[37436]++
							msg, code := toHTTPError(err)
							Error(w, msg, code)
							return
//line /usr/local/go/src/net/http/fs.go:618
		// _ = "end of CoverTab[37436]"
	} else {
//line /usr/local/go/src/net/http/fs.go:619
		_go_fuzz_dep_.CoverTab[37437]++
//line /usr/local/go/src/net/http/fs.go:619
		// _ = "end of CoverTab[37437]"
//line /usr/local/go/src/net/http/fs.go:619
	}
//line /usr/local/go/src/net/http/fs.go:619
	// _ = "end of CoverTab[37427]"
//line /usr/local/go/src/net/http/fs.go:619
	_go_fuzz_dep_.CoverTab[37428]++
						defer f.Close()

						d, err := f.Stat()
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:623
		_go_fuzz_dep_.CoverTab[37438]++
							msg, code := toHTTPError(err)
							Error(w, msg, code)
							return
//line /usr/local/go/src/net/http/fs.go:626
		// _ = "end of CoverTab[37438]"
	} else {
//line /usr/local/go/src/net/http/fs.go:627
		_go_fuzz_dep_.CoverTab[37439]++
//line /usr/local/go/src/net/http/fs.go:627
		// _ = "end of CoverTab[37439]"
//line /usr/local/go/src/net/http/fs.go:627
	}
//line /usr/local/go/src/net/http/fs.go:627
	// _ = "end of CoverTab[37428]"
//line /usr/local/go/src/net/http/fs.go:627
	_go_fuzz_dep_.CoverTab[37429]++

						if redirect {
//line /usr/local/go/src/net/http/fs.go:629
		_go_fuzz_dep_.CoverTab[37440]++

//line /usr/local/go/src/net/http/fs.go:632
		url := r.URL.Path
		if d.IsDir() {
//line /usr/local/go/src/net/http/fs.go:633
			_go_fuzz_dep_.CoverTab[37441]++
								if url[len(url)-1] != '/' {
//line /usr/local/go/src/net/http/fs.go:634
				_go_fuzz_dep_.CoverTab[37442]++
									localRedirect(w, r, path.Base(url)+"/")
									return
//line /usr/local/go/src/net/http/fs.go:636
				// _ = "end of CoverTab[37442]"
			} else {
//line /usr/local/go/src/net/http/fs.go:637
				_go_fuzz_dep_.CoverTab[37443]++
//line /usr/local/go/src/net/http/fs.go:637
				// _ = "end of CoverTab[37443]"
//line /usr/local/go/src/net/http/fs.go:637
			}
//line /usr/local/go/src/net/http/fs.go:637
			// _ = "end of CoverTab[37441]"
		} else {
//line /usr/local/go/src/net/http/fs.go:638
			_go_fuzz_dep_.CoverTab[37444]++
								if url[len(url)-1] == '/' {
//line /usr/local/go/src/net/http/fs.go:639
				_go_fuzz_dep_.CoverTab[37445]++
									localRedirect(w, r, "../"+path.Base(url))
									return
//line /usr/local/go/src/net/http/fs.go:641
				// _ = "end of CoverTab[37445]"
			} else {
//line /usr/local/go/src/net/http/fs.go:642
				_go_fuzz_dep_.CoverTab[37446]++
//line /usr/local/go/src/net/http/fs.go:642
				// _ = "end of CoverTab[37446]"
//line /usr/local/go/src/net/http/fs.go:642
			}
//line /usr/local/go/src/net/http/fs.go:642
			// _ = "end of CoverTab[37444]"
		}
//line /usr/local/go/src/net/http/fs.go:643
		// _ = "end of CoverTab[37440]"
	} else {
//line /usr/local/go/src/net/http/fs.go:644
		_go_fuzz_dep_.CoverTab[37447]++
//line /usr/local/go/src/net/http/fs.go:644
		// _ = "end of CoverTab[37447]"
//line /usr/local/go/src/net/http/fs.go:644
	}
//line /usr/local/go/src/net/http/fs.go:644
	// _ = "end of CoverTab[37429]"
//line /usr/local/go/src/net/http/fs.go:644
	_go_fuzz_dep_.CoverTab[37430]++

						if d.IsDir() {
//line /usr/local/go/src/net/http/fs.go:646
		_go_fuzz_dep_.CoverTab[37448]++
							url := r.URL.Path

							if url == "" || func() bool {
//line /usr/local/go/src/net/http/fs.go:649
			_go_fuzz_dep_.CoverTab[37450]++
//line /usr/local/go/src/net/http/fs.go:649
			return url[len(url)-1] != '/'
//line /usr/local/go/src/net/http/fs.go:649
			// _ = "end of CoverTab[37450]"
//line /usr/local/go/src/net/http/fs.go:649
		}() {
//line /usr/local/go/src/net/http/fs.go:649
			_go_fuzz_dep_.CoverTab[37451]++
								localRedirect(w, r, path.Base(url)+"/")
								return
//line /usr/local/go/src/net/http/fs.go:651
			// _ = "end of CoverTab[37451]"
		} else {
//line /usr/local/go/src/net/http/fs.go:652
			_go_fuzz_dep_.CoverTab[37452]++
//line /usr/local/go/src/net/http/fs.go:652
			// _ = "end of CoverTab[37452]"
//line /usr/local/go/src/net/http/fs.go:652
		}
//line /usr/local/go/src/net/http/fs.go:652
		// _ = "end of CoverTab[37448]"
//line /usr/local/go/src/net/http/fs.go:652
		_go_fuzz_dep_.CoverTab[37449]++

//line /usr/local/go/src/net/http/fs.go:655
		index := strings.TrimSuffix(name, "/") + indexPage
		ff, err := fs.Open(index)
		if err == nil {
//line /usr/local/go/src/net/http/fs.go:657
			_go_fuzz_dep_.CoverTab[37453]++
								defer ff.Close()
								dd, err := ff.Stat()
								if err == nil {
//line /usr/local/go/src/net/http/fs.go:660
				_go_fuzz_dep_.CoverTab[37454]++
									d = dd
									f = ff
//line /usr/local/go/src/net/http/fs.go:662
				// _ = "end of CoverTab[37454]"
			} else {
//line /usr/local/go/src/net/http/fs.go:663
				_go_fuzz_dep_.CoverTab[37455]++
//line /usr/local/go/src/net/http/fs.go:663
				// _ = "end of CoverTab[37455]"
//line /usr/local/go/src/net/http/fs.go:663
			}
//line /usr/local/go/src/net/http/fs.go:663
			// _ = "end of CoverTab[37453]"
		} else {
//line /usr/local/go/src/net/http/fs.go:664
			_go_fuzz_dep_.CoverTab[37456]++
//line /usr/local/go/src/net/http/fs.go:664
			// _ = "end of CoverTab[37456]"
//line /usr/local/go/src/net/http/fs.go:664
		}
//line /usr/local/go/src/net/http/fs.go:664
		// _ = "end of CoverTab[37449]"
	} else {
//line /usr/local/go/src/net/http/fs.go:665
		_go_fuzz_dep_.CoverTab[37457]++
//line /usr/local/go/src/net/http/fs.go:665
		// _ = "end of CoverTab[37457]"
//line /usr/local/go/src/net/http/fs.go:665
	}
//line /usr/local/go/src/net/http/fs.go:665
	// _ = "end of CoverTab[37430]"
//line /usr/local/go/src/net/http/fs.go:665
	_go_fuzz_dep_.CoverTab[37431]++

//line /usr/local/go/src/net/http/fs.go:668
	if d.IsDir() {
//line /usr/local/go/src/net/http/fs.go:668
		_go_fuzz_dep_.CoverTab[37458]++
							if checkIfModifiedSince(r, d.ModTime()) == condFalse {
//line /usr/local/go/src/net/http/fs.go:669
			_go_fuzz_dep_.CoverTab[37460]++
								writeNotModified(w)
								return
//line /usr/local/go/src/net/http/fs.go:671
			// _ = "end of CoverTab[37460]"
		} else {
//line /usr/local/go/src/net/http/fs.go:672
			_go_fuzz_dep_.CoverTab[37461]++
//line /usr/local/go/src/net/http/fs.go:672
			// _ = "end of CoverTab[37461]"
//line /usr/local/go/src/net/http/fs.go:672
		}
//line /usr/local/go/src/net/http/fs.go:672
		// _ = "end of CoverTab[37458]"
//line /usr/local/go/src/net/http/fs.go:672
		_go_fuzz_dep_.CoverTab[37459]++
							setLastModified(w, d.ModTime())
							dirList(w, r, f)
							return
//line /usr/local/go/src/net/http/fs.go:675
		// _ = "end of CoverTab[37459]"
	} else {
//line /usr/local/go/src/net/http/fs.go:676
		_go_fuzz_dep_.CoverTab[37462]++
//line /usr/local/go/src/net/http/fs.go:676
		// _ = "end of CoverTab[37462]"
//line /usr/local/go/src/net/http/fs.go:676
	}
//line /usr/local/go/src/net/http/fs.go:676
	// _ = "end of CoverTab[37431]"
//line /usr/local/go/src/net/http/fs.go:676
	_go_fuzz_dep_.CoverTab[37432]++

//line /usr/local/go/src/net/http/fs.go:679
	sizeFunc := func() (int64, error) {
//line /usr/local/go/src/net/http/fs.go:679
		_go_fuzz_dep_.CoverTab[37463]++
//line /usr/local/go/src/net/http/fs.go:679
		return d.Size(), nil
//line /usr/local/go/src/net/http/fs.go:679
		// _ = "end of CoverTab[37463]"
//line /usr/local/go/src/net/http/fs.go:679
	}
//line /usr/local/go/src/net/http/fs.go:679
	// _ = "end of CoverTab[37432]"
//line /usr/local/go/src/net/http/fs.go:679
	_go_fuzz_dep_.CoverTab[37433]++
						serveContent(w, r, d.Name(), d.ModTime(), sizeFunc, f)
//line /usr/local/go/src/net/http/fs.go:680
	// _ = "end of CoverTab[37433]"
}

// toHTTPError returns a non-specific HTTP error message and status code
//line /usr/local/go/src/net/http/fs.go:683
// for a given non-nil error value. It's important that toHTTPError does not
//line /usr/local/go/src/net/http/fs.go:683
// actually return err.Error(), since msg and httpStatus are returned to users,
//line /usr/local/go/src/net/http/fs.go:683
// and historically Go's ServeContent always returned just "404 Not Found" for
//line /usr/local/go/src/net/http/fs.go:683
// all errors. We don't want to start leaking information in error messages.
//line /usr/local/go/src/net/http/fs.go:688
func toHTTPError(err error) (msg string, httpStatus int) {
//line /usr/local/go/src/net/http/fs.go:688
	_go_fuzz_dep_.CoverTab[37464]++
						if errors.Is(err, fs.ErrNotExist) {
//line /usr/local/go/src/net/http/fs.go:689
		_go_fuzz_dep_.CoverTab[37467]++
							return "404 page not found", StatusNotFound
//line /usr/local/go/src/net/http/fs.go:690
		// _ = "end of CoverTab[37467]"
	} else {
//line /usr/local/go/src/net/http/fs.go:691
		_go_fuzz_dep_.CoverTab[37468]++
//line /usr/local/go/src/net/http/fs.go:691
		// _ = "end of CoverTab[37468]"
//line /usr/local/go/src/net/http/fs.go:691
	}
//line /usr/local/go/src/net/http/fs.go:691
	// _ = "end of CoverTab[37464]"
//line /usr/local/go/src/net/http/fs.go:691
	_go_fuzz_dep_.CoverTab[37465]++
						if errors.Is(err, fs.ErrPermission) {
//line /usr/local/go/src/net/http/fs.go:692
		_go_fuzz_dep_.CoverTab[37469]++
							return "403 Forbidden", StatusForbidden
//line /usr/local/go/src/net/http/fs.go:693
		// _ = "end of CoverTab[37469]"
	} else {
//line /usr/local/go/src/net/http/fs.go:694
		_go_fuzz_dep_.CoverTab[37470]++
//line /usr/local/go/src/net/http/fs.go:694
		// _ = "end of CoverTab[37470]"
//line /usr/local/go/src/net/http/fs.go:694
	}
//line /usr/local/go/src/net/http/fs.go:694
	// _ = "end of CoverTab[37465]"
//line /usr/local/go/src/net/http/fs.go:694
	_go_fuzz_dep_.CoverTab[37466]++

						return "500 Internal Server Error", StatusInternalServerError
//line /usr/local/go/src/net/http/fs.go:696
	// _ = "end of CoverTab[37466]"
}

// localRedirect gives a Moved Permanently response.
//line /usr/local/go/src/net/http/fs.go:699
// It does not convert relative paths to absolute paths like Redirect does.
//line /usr/local/go/src/net/http/fs.go:701
func localRedirect(w ResponseWriter, r *Request, newPath string) {
//line /usr/local/go/src/net/http/fs.go:701
	_go_fuzz_dep_.CoverTab[37471]++
						if q := r.URL.RawQuery; q != "" {
//line /usr/local/go/src/net/http/fs.go:702
		_go_fuzz_dep_.CoverTab[37473]++
							newPath += "?" + q
//line /usr/local/go/src/net/http/fs.go:703
		// _ = "end of CoverTab[37473]"
	} else {
//line /usr/local/go/src/net/http/fs.go:704
		_go_fuzz_dep_.CoverTab[37474]++
//line /usr/local/go/src/net/http/fs.go:704
		// _ = "end of CoverTab[37474]"
//line /usr/local/go/src/net/http/fs.go:704
	}
//line /usr/local/go/src/net/http/fs.go:704
	// _ = "end of CoverTab[37471]"
//line /usr/local/go/src/net/http/fs.go:704
	_go_fuzz_dep_.CoverTab[37472]++
						w.Header().Set("Location", newPath)
						w.WriteHeader(StatusMovedPermanently)
//line /usr/local/go/src/net/http/fs.go:706
	// _ = "end of CoverTab[37472]"
}

// ServeFile replies to the request with the contents of the named
//line /usr/local/go/src/net/http/fs.go:709
// file or directory.
//line /usr/local/go/src/net/http/fs.go:709
//
//line /usr/local/go/src/net/http/fs.go:709
// If the provided file or directory name is a relative path, it is
//line /usr/local/go/src/net/http/fs.go:709
// interpreted relative to the current directory and may ascend to
//line /usr/local/go/src/net/http/fs.go:709
// parent directories. If the provided name is constructed from user
//line /usr/local/go/src/net/http/fs.go:709
// input, it should be sanitized before calling ServeFile.
//line /usr/local/go/src/net/http/fs.go:709
//
//line /usr/local/go/src/net/http/fs.go:709
// As a precaution, ServeFile will reject requests where r.URL.Path
//line /usr/local/go/src/net/http/fs.go:709
// contains a ".." path element; this protects against callers who
//line /usr/local/go/src/net/http/fs.go:709
// might unsafely use filepath.Join on r.URL.Path without sanitizing
//line /usr/local/go/src/net/http/fs.go:709
// it and then use that filepath.Join result as the name argument.
//line /usr/local/go/src/net/http/fs.go:709
//
//line /usr/local/go/src/net/http/fs.go:709
// As another special case, ServeFile redirects any request where r.URL.Path
//line /usr/local/go/src/net/http/fs.go:709
// ends in "/index.html" to the same path, without the final
//line /usr/local/go/src/net/http/fs.go:709
// "index.html". To avoid such redirects either modify the path or
//line /usr/local/go/src/net/http/fs.go:709
// use ServeContent.
//line /usr/local/go/src/net/http/fs.go:709
//
//line /usr/local/go/src/net/http/fs.go:709
// Outside of those two special cases, ServeFile does not use
//line /usr/local/go/src/net/http/fs.go:709
// r.URL.Path for selecting the file or directory to serve; only the
//line /usr/local/go/src/net/http/fs.go:709
// file or directory provided in the name argument is used.
//line /usr/local/go/src/net/http/fs.go:730
func ServeFile(w ResponseWriter, r *Request, name string) {
//line /usr/local/go/src/net/http/fs.go:730
	_go_fuzz_dep_.CoverTab[37475]++
						if containsDotDot(r.URL.Path) {
//line /usr/local/go/src/net/http/fs.go:731
		_go_fuzz_dep_.CoverTab[37477]++

//line /usr/local/go/src/net/http/fs.go:737
		Error(w, "invalid URL path", StatusBadRequest)
							return
//line /usr/local/go/src/net/http/fs.go:738
		// _ = "end of CoverTab[37477]"
	} else {
//line /usr/local/go/src/net/http/fs.go:739
		_go_fuzz_dep_.CoverTab[37478]++
//line /usr/local/go/src/net/http/fs.go:739
		// _ = "end of CoverTab[37478]"
//line /usr/local/go/src/net/http/fs.go:739
	}
//line /usr/local/go/src/net/http/fs.go:739
	// _ = "end of CoverTab[37475]"
//line /usr/local/go/src/net/http/fs.go:739
	_go_fuzz_dep_.CoverTab[37476]++
						dir, file := filepath.Split(name)
						serveFile(w, r, Dir(dir), file, false)
//line /usr/local/go/src/net/http/fs.go:741
	// _ = "end of CoverTab[37476]"
}

func containsDotDot(v string) bool {
//line /usr/local/go/src/net/http/fs.go:744
	_go_fuzz_dep_.CoverTab[37479]++
						if !strings.Contains(v, "..") {
//line /usr/local/go/src/net/http/fs.go:745
		_go_fuzz_dep_.CoverTab[37482]++
							return false
//line /usr/local/go/src/net/http/fs.go:746
		// _ = "end of CoverTab[37482]"
	} else {
//line /usr/local/go/src/net/http/fs.go:747
		_go_fuzz_dep_.CoverTab[37483]++
//line /usr/local/go/src/net/http/fs.go:747
		// _ = "end of CoverTab[37483]"
//line /usr/local/go/src/net/http/fs.go:747
	}
//line /usr/local/go/src/net/http/fs.go:747
	// _ = "end of CoverTab[37479]"
//line /usr/local/go/src/net/http/fs.go:747
	_go_fuzz_dep_.CoverTab[37480]++
						for _, ent := range strings.FieldsFunc(v, isSlashRune) {
//line /usr/local/go/src/net/http/fs.go:748
		_go_fuzz_dep_.CoverTab[37484]++
							if ent == ".." {
//line /usr/local/go/src/net/http/fs.go:749
			_go_fuzz_dep_.CoverTab[37485]++
								return true
//line /usr/local/go/src/net/http/fs.go:750
			// _ = "end of CoverTab[37485]"
		} else {
//line /usr/local/go/src/net/http/fs.go:751
			_go_fuzz_dep_.CoverTab[37486]++
//line /usr/local/go/src/net/http/fs.go:751
			// _ = "end of CoverTab[37486]"
//line /usr/local/go/src/net/http/fs.go:751
		}
//line /usr/local/go/src/net/http/fs.go:751
		// _ = "end of CoverTab[37484]"
	}
//line /usr/local/go/src/net/http/fs.go:752
	// _ = "end of CoverTab[37480]"
//line /usr/local/go/src/net/http/fs.go:752
	_go_fuzz_dep_.CoverTab[37481]++
						return false
//line /usr/local/go/src/net/http/fs.go:753
	// _ = "end of CoverTab[37481]"
}

func isSlashRune(r rune) bool {
//line /usr/local/go/src/net/http/fs.go:756
	_go_fuzz_dep_.CoverTab[37487]++
//line /usr/local/go/src/net/http/fs.go:756
	return r == '/' || func() bool {
//line /usr/local/go/src/net/http/fs.go:756
		_go_fuzz_dep_.CoverTab[37488]++
//line /usr/local/go/src/net/http/fs.go:756
		return r == '\\'
//line /usr/local/go/src/net/http/fs.go:756
		// _ = "end of CoverTab[37488]"
//line /usr/local/go/src/net/http/fs.go:756
	}()
//line /usr/local/go/src/net/http/fs.go:756
	// _ = "end of CoverTab[37487]"
//line /usr/local/go/src/net/http/fs.go:756
}

type fileHandler struct {
	root FileSystem
}

type ioFS struct {
	fsys fs.FS
}

type ioFile struct {
	file fs.File
}

func (f ioFS) Open(name string) (File, error) {
//line /usr/local/go/src/net/http/fs.go:770
	_go_fuzz_dep_.CoverTab[37489]++
						if name == "/" {
//line /usr/local/go/src/net/http/fs.go:771
		_go_fuzz_dep_.CoverTab[37492]++
							name = "."
//line /usr/local/go/src/net/http/fs.go:772
		// _ = "end of CoverTab[37492]"
	} else {
//line /usr/local/go/src/net/http/fs.go:773
		_go_fuzz_dep_.CoverTab[37493]++
							name = strings.TrimPrefix(name, "/")
//line /usr/local/go/src/net/http/fs.go:774
		// _ = "end of CoverTab[37493]"
	}
//line /usr/local/go/src/net/http/fs.go:775
	// _ = "end of CoverTab[37489]"
//line /usr/local/go/src/net/http/fs.go:775
	_go_fuzz_dep_.CoverTab[37490]++
						file, err := f.fsys.Open(name)
						if err != nil {
//line /usr/local/go/src/net/http/fs.go:777
		_go_fuzz_dep_.CoverTab[37494]++
							return nil, mapOpenError(err, name, '/', func(path string) (fs.FileInfo, error) {
//line /usr/local/go/src/net/http/fs.go:778
			_go_fuzz_dep_.CoverTab[37495]++
								return fs.Stat(f.fsys, path)
//line /usr/local/go/src/net/http/fs.go:779
			// _ = "end of CoverTab[37495]"
		})
//line /usr/local/go/src/net/http/fs.go:780
		// _ = "end of CoverTab[37494]"
	} else {
//line /usr/local/go/src/net/http/fs.go:781
		_go_fuzz_dep_.CoverTab[37496]++
//line /usr/local/go/src/net/http/fs.go:781
		// _ = "end of CoverTab[37496]"
//line /usr/local/go/src/net/http/fs.go:781
	}
//line /usr/local/go/src/net/http/fs.go:781
	// _ = "end of CoverTab[37490]"
//line /usr/local/go/src/net/http/fs.go:781
	_go_fuzz_dep_.CoverTab[37491]++
						return ioFile{file}, nil
//line /usr/local/go/src/net/http/fs.go:782
	// _ = "end of CoverTab[37491]"
}

func (f ioFile) Close() error {
//line /usr/local/go/src/net/http/fs.go:785
	_go_fuzz_dep_.CoverTab[37497]++
//line /usr/local/go/src/net/http/fs.go:785
	return f.file.Close()
//line /usr/local/go/src/net/http/fs.go:785
	// _ = "end of CoverTab[37497]"
//line /usr/local/go/src/net/http/fs.go:785
}
func (f ioFile) Read(b []byte) (int, error) {
//line /usr/local/go/src/net/http/fs.go:786
	_go_fuzz_dep_.CoverTab[37498]++
//line /usr/local/go/src/net/http/fs.go:786
	return f.file.Read(b)
//line /usr/local/go/src/net/http/fs.go:786
	// _ = "end of CoverTab[37498]"
//line /usr/local/go/src/net/http/fs.go:786
}
func (f ioFile) Stat() (fs.FileInfo, error) {
//line /usr/local/go/src/net/http/fs.go:787
	_go_fuzz_dep_.CoverTab[37499]++
//line /usr/local/go/src/net/http/fs.go:787
	return f.file.Stat()
//line /usr/local/go/src/net/http/fs.go:787
	// _ = "end of CoverTab[37499]"
//line /usr/local/go/src/net/http/fs.go:787
}

var errMissingSeek = errors.New("io.File missing Seek method")
var errMissingReadDir = errors.New("io.File directory missing ReadDir method")

func (f ioFile) Seek(offset int64, whence int) (int64, error) {
//line /usr/local/go/src/net/http/fs.go:792
	_go_fuzz_dep_.CoverTab[37500]++
						s, ok := f.file.(io.Seeker)
						if !ok {
//line /usr/local/go/src/net/http/fs.go:794
		_go_fuzz_dep_.CoverTab[37502]++
							return 0, errMissingSeek
//line /usr/local/go/src/net/http/fs.go:795
		// _ = "end of CoverTab[37502]"
	} else {
//line /usr/local/go/src/net/http/fs.go:796
		_go_fuzz_dep_.CoverTab[37503]++
//line /usr/local/go/src/net/http/fs.go:796
		// _ = "end of CoverTab[37503]"
//line /usr/local/go/src/net/http/fs.go:796
	}
//line /usr/local/go/src/net/http/fs.go:796
	// _ = "end of CoverTab[37500]"
//line /usr/local/go/src/net/http/fs.go:796
	_go_fuzz_dep_.CoverTab[37501]++
						return s.Seek(offset, whence)
//line /usr/local/go/src/net/http/fs.go:797
	// _ = "end of CoverTab[37501]"
}

func (f ioFile) ReadDir(count int) ([]fs.DirEntry, error) {
//line /usr/local/go/src/net/http/fs.go:800
	_go_fuzz_dep_.CoverTab[37504]++
						d, ok := f.file.(fs.ReadDirFile)
						if !ok {
//line /usr/local/go/src/net/http/fs.go:802
		_go_fuzz_dep_.CoverTab[37506]++
							return nil, errMissingReadDir
//line /usr/local/go/src/net/http/fs.go:803
		// _ = "end of CoverTab[37506]"
	} else {
//line /usr/local/go/src/net/http/fs.go:804
		_go_fuzz_dep_.CoverTab[37507]++
//line /usr/local/go/src/net/http/fs.go:804
		// _ = "end of CoverTab[37507]"
//line /usr/local/go/src/net/http/fs.go:804
	}
//line /usr/local/go/src/net/http/fs.go:804
	// _ = "end of CoverTab[37504]"
//line /usr/local/go/src/net/http/fs.go:804
	_go_fuzz_dep_.CoverTab[37505]++
						return d.ReadDir(count)
//line /usr/local/go/src/net/http/fs.go:805
	// _ = "end of CoverTab[37505]"
}

func (f ioFile) Readdir(count int) ([]fs.FileInfo, error) {
//line /usr/local/go/src/net/http/fs.go:808
	_go_fuzz_dep_.CoverTab[37508]++
						d, ok := f.file.(fs.ReadDirFile)
						if !ok {
//line /usr/local/go/src/net/http/fs.go:810
		_go_fuzz_dep_.CoverTab[37511]++
							return nil, errMissingReadDir
//line /usr/local/go/src/net/http/fs.go:811
		// _ = "end of CoverTab[37511]"
	} else {
//line /usr/local/go/src/net/http/fs.go:812
		_go_fuzz_dep_.CoverTab[37512]++
//line /usr/local/go/src/net/http/fs.go:812
		// _ = "end of CoverTab[37512]"
//line /usr/local/go/src/net/http/fs.go:812
	}
//line /usr/local/go/src/net/http/fs.go:812
	// _ = "end of CoverTab[37508]"
//line /usr/local/go/src/net/http/fs.go:812
	_go_fuzz_dep_.CoverTab[37509]++
						var list []fs.FileInfo
						for {
//line /usr/local/go/src/net/http/fs.go:814
		_go_fuzz_dep_.CoverTab[37513]++
							dirs, err := d.ReadDir(count - len(list))
							for _, dir := range dirs {
//line /usr/local/go/src/net/http/fs.go:816
			_go_fuzz_dep_.CoverTab[37516]++
								info, err := dir.Info()
								if err != nil {
//line /usr/local/go/src/net/http/fs.go:818
				_go_fuzz_dep_.CoverTab[37518]++

									continue
//line /usr/local/go/src/net/http/fs.go:820
				// _ = "end of CoverTab[37518]"
			} else {
//line /usr/local/go/src/net/http/fs.go:821
				_go_fuzz_dep_.CoverTab[37519]++
//line /usr/local/go/src/net/http/fs.go:821
				// _ = "end of CoverTab[37519]"
//line /usr/local/go/src/net/http/fs.go:821
			}
//line /usr/local/go/src/net/http/fs.go:821
			// _ = "end of CoverTab[37516]"
//line /usr/local/go/src/net/http/fs.go:821
			_go_fuzz_dep_.CoverTab[37517]++
								list = append(list, info)
//line /usr/local/go/src/net/http/fs.go:822
			// _ = "end of CoverTab[37517]"
		}
//line /usr/local/go/src/net/http/fs.go:823
		// _ = "end of CoverTab[37513]"
//line /usr/local/go/src/net/http/fs.go:823
		_go_fuzz_dep_.CoverTab[37514]++
							if err != nil {
//line /usr/local/go/src/net/http/fs.go:824
			_go_fuzz_dep_.CoverTab[37520]++
								return list, err
//line /usr/local/go/src/net/http/fs.go:825
			// _ = "end of CoverTab[37520]"
		} else {
//line /usr/local/go/src/net/http/fs.go:826
			_go_fuzz_dep_.CoverTab[37521]++
//line /usr/local/go/src/net/http/fs.go:826
			// _ = "end of CoverTab[37521]"
//line /usr/local/go/src/net/http/fs.go:826
		}
//line /usr/local/go/src/net/http/fs.go:826
		// _ = "end of CoverTab[37514]"
//line /usr/local/go/src/net/http/fs.go:826
		_go_fuzz_dep_.CoverTab[37515]++
							if count < 0 || func() bool {
//line /usr/local/go/src/net/http/fs.go:827
			_go_fuzz_dep_.CoverTab[37522]++
//line /usr/local/go/src/net/http/fs.go:827
			return len(list) >= count
//line /usr/local/go/src/net/http/fs.go:827
			// _ = "end of CoverTab[37522]"
//line /usr/local/go/src/net/http/fs.go:827
		}() {
//line /usr/local/go/src/net/http/fs.go:827
			_go_fuzz_dep_.CoverTab[37523]++
								break
//line /usr/local/go/src/net/http/fs.go:828
			// _ = "end of CoverTab[37523]"
		} else {
//line /usr/local/go/src/net/http/fs.go:829
			_go_fuzz_dep_.CoverTab[37524]++
//line /usr/local/go/src/net/http/fs.go:829
			// _ = "end of CoverTab[37524]"
//line /usr/local/go/src/net/http/fs.go:829
		}
//line /usr/local/go/src/net/http/fs.go:829
		// _ = "end of CoverTab[37515]"
	}
//line /usr/local/go/src/net/http/fs.go:830
	// _ = "end of CoverTab[37509]"
//line /usr/local/go/src/net/http/fs.go:830
	_go_fuzz_dep_.CoverTab[37510]++
						return list, nil
//line /usr/local/go/src/net/http/fs.go:831
	// _ = "end of CoverTab[37510]"
}

// FS converts fsys to a FileSystem implementation,
//line /usr/local/go/src/net/http/fs.go:834
// for use with FileServer and NewFileTransport.
//line /usr/local/go/src/net/http/fs.go:834
// The files provided by fsys must implement io.Seeker.
//line /usr/local/go/src/net/http/fs.go:837
func FS(fsys fs.FS) FileSystem {
//line /usr/local/go/src/net/http/fs.go:837
	_go_fuzz_dep_.CoverTab[37525]++
						return ioFS{fsys}
//line /usr/local/go/src/net/http/fs.go:838
	// _ = "end of CoverTab[37525]"
}

// FileServer returns a handler that serves HTTP requests
//line /usr/local/go/src/net/http/fs.go:841
// with the contents of the file system rooted at root.
//line /usr/local/go/src/net/http/fs.go:841
//
//line /usr/local/go/src/net/http/fs.go:841
// As a special case, the returned file server redirects any request
//line /usr/local/go/src/net/http/fs.go:841
// ending in "/index.html" to the same path, without the final
//line /usr/local/go/src/net/http/fs.go:841
// "index.html".
//line /usr/local/go/src/net/http/fs.go:841
//
//line /usr/local/go/src/net/http/fs.go:841
// To use the operating system's file system implementation,
//line /usr/local/go/src/net/http/fs.go:841
// use http.Dir:
//line /usr/local/go/src/net/http/fs.go:841
//
//line /usr/local/go/src/net/http/fs.go:841
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//line /usr/local/go/src/net/http/fs.go:841
//
//line /usr/local/go/src/net/http/fs.go:841
// To use an fs.FS implementation, use http.FS to convert it:
//line /usr/local/go/src/net/http/fs.go:841
//
//line /usr/local/go/src/net/http/fs.go:841
//	http.Handle("/", http.FileServer(http.FS(fsys)))
//line /usr/local/go/src/net/http/fs.go:856
func FileServer(root FileSystem) Handler {
//line /usr/local/go/src/net/http/fs.go:856
	_go_fuzz_dep_.CoverTab[37526]++
						return &fileHandler{root}
//line /usr/local/go/src/net/http/fs.go:857
	// _ = "end of CoverTab[37526]"
}

func (f *fileHandler) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/fs.go:860
	_go_fuzz_dep_.CoverTab[37527]++
						upath := r.URL.Path
						if !strings.HasPrefix(upath, "/") {
//line /usr/local/go/src/net/http/fs.go:862
		_go_fuzz_dep_.CoverTab[37529]++
							upath = "/" + upath
							r.URL.Path = upath
//line /usr/local/go/src/net/http/fs.go:864
		// _ = "end of CoverTab[37529]"
	} else {
//line /usr/local/go/src/net/http/fs.go:865
		_go_fuzz_dep_.CoverTab[37530]++
//line /usr/local/go/src/net/http/fs.go:865
		// _ = "end of CoverTab[37530]"
//line /usr/local/go/src/net/http/fs.go:865
	}
//line /usr/local/go/src/net/http/fs.go:865
	// _ = "end of CoverTab[37527]"
//line /usr/local/go/src/net/http/fs.go:865
	_go_fuzz_dep_.CoverTab[37528]++
						serveFile(w, r, f.root, path.Clean(upath), true)
//line /usr/local/go/src/net/http/fs.go:866
	// _ = "end of CoverTab[37528]"
}

// httpRange specifies the byte range to be sent to the client.
type httpRange struct {
	start, length int64
}

func (r httpRange) contentRange(size int64) string {
//line /usr/local/go/src/net/http/fs.go:874
	_go_fuzz_dep_.CoverTab[37531]++
						return fmt.Sprintf("bytes %d-%d/%d", r.start, r.start+r.length-1, size)
//line /usr/local/go/src/net/http/fs.go:875
	// _ = "end of CoverTab[37531]"
}

func (r httpRange) mimeHeader(contentType string, size int64) textproto.MIMEHeader {
//line /usr/local/go/src/net/http/fs.go:878
	_go_fuzz_dep_.CoverTab[37532]++
						return textproto.MIMEHeader{
		"Content-Range":	{r.contentRange(size)},
		"Content-Type":		{contentType},
	}
//line /usr/local/go/src/net/http/fs.go:882
	// _ = "end of CoverTab[37532]"
}

// parseRange parses a Range header string as per RFC 7233.
//line /usr/local/go/src/net/http/fs.go:885
// errNoOverlap is returned if none of the ranges overlap.
//line /usr/local/go/src/net/http/fs.go:887
func parseRange(s string, size int64) ([]httpRange, error) {
//line /usr/local/go/src/net/http/fs.go:887
	_go_fuzz_dep_.CoverTab[37533]++
						if s == "" {
//line /usr/local/go/src/net/http/fs.go:888
		_go_fuzz_dep_.CoverTab[37538]++
							return nil, nil
//line /usr/local/go/src/net/http/fs.go:889
		// _ = "end of CoverTab[37538]"
	} else {
//line /usr/local/go/src/net/http/fs.go:890
		_go_fuzz_dep_.CoverTab[37539]++
//line /usr/local/go/src/net/http/fs.go:890
		// _ = "end of CoverTab[37539]"
//line /usr/local/go/src/net/http/fs.go:890
	}
//line /usr/local/go/src/net/http/fs.go:890
	// _ = "end of CoverTab[37533]"
//line /usr/local/go/src/net/http/fs.go:890
	_go_fuzz_dep_.CoverTab[37534]++
						const b = "bytes="
						if !strings.HasPrefix(s, b) {
//line /usr/local/go/src/net/http/fs.go:892
		_go_fuzz_dep_.CoverTab[37540]++
							return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:893
		// _ = "end of CoverTab[37540]"
	} else {
//line /usr/local/go/src/net/http/fs.go:894
		_go_fuzz_dep_.CoverTab[37541]++
//line /usr/local/go/src/net/http/fs.go:894
		// _ = "end of CoverTab[37541]"
//line /usr/local/go/src/net/http/fs.go:894
	}
//line /usr/local/go/src/net/http/fs.go:894
	// _ = "end of CoverTab[37534]"
//line /usr/local/go/src/net/http/fs.go:894
	_go_fuzz_dep_.CoverTab[37535]++
						var ranges []httpRange
						noOverlap := false
						for _, ra := range strings.Split(s[len(b):], ",") {
//line /usr/local/go/src/net/http/fs.go:897
		_go_fuzz_dep_.CoverTab[37542]++
							ra = textproto.TrimString(ra)
							if ra == "" {
//line /usr/local/go/src/net/http/fs.go:899
			_go_fuzz_dep_.CoverTab[37546]++
								continue
//line /usr/local/go/src/net/http/fs.go:900
			// _ = "end of CoverTab[37546]"
		} else {
//line /usr/local/go/src/net/http/fs.go:901
			_go_fuzz_dep_.CoverTab[37547]++
//line /usr/local/go/src/net/http/fs.go:901
			// _ = "end of CoverTab[37547]"
//line /usr/local/go/src/net/http/fs.go:901
		}
//line /usr/local/go/src/net/http/fs.go:901
		// _ = "end of CoverTab[37542]"
//line /usr/local/go/src/net/http/fs.go:901
		_go_fuzz_dep_.CoverTab[37543]++
							start, end, ok := strings.Cut(ra, "-")
							if !ok {
//line /usr/local/go/src/net/http/fs.go:903
			_go_fuzz_dep_.CoverTab[37548]++
								return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:904
			// _ = "end of CoverTab[37548]"
		} else {
//line /usr/local/go/src/net/http/fs.go:905
			_go_fuzz_dep_.CoverTab[37549]++
//line /usr/local/go/src/net/http/fs.go:905
			// _ = "end of CoverTab[37549]"
//line /usr/local/go/src/net/http/fs.go:905
		}
//line /usr/local/go/src/net/http/fs.go:905
		// _ = "end of CoverTab[37543]"
//line /usr/local/go/src/net/http/fs.go:905
		_go_fuzz_dep_.CoverTab[37544]++
							start, end = textproto.TrimString(start), textproto.TrimString(end)
							var r httpRange
							if start == "" {
//line /usr/local/go/src/net/http/fs.go:908
			_go_fuzz_dep_.CoverTab[37550]++

//line /usr/local/go/src/net/http/fs.go:914
			if end == "" || func() bool {
//line /usr/local/go/src/net/http/fs.go:914
				_go_fuzz_dep_.CoverTab[37554]++
//line /usr/local/go/src/net/http/fs.go:914
				return end[0] == '-'
//line /usr/local/go/src/net/http/fs.go:914
				// _ = "end of CoverTab[37554]"
//line /usr/local/go/src/net/http/fs.go:914
			}() {
//line /usr/local/go/src/net/http/fs.go:914
				_go_fuzz_dep_.CoverTab[37555]++
									return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:915
				// _ = "end of CoverTab[37555]"
			} else {
//line /usr/local/go/src/net/http/fs.go:916
				_go_fuzz_dep_.CoverTab[37556]++
//line /usr/local/go/src/net/http/fs.go:916
				// _ = "end of CoverTab[37556]"
//line /usr/local/go/src/net/http/fs.go:916
			}
//line /usr/local/go/src/net/http/fs.go:916
			// _ = "end of CoverTab[37550]"
//line /usr/local/go/src/net/http/fs.go:916
			_go_fuzz_dep_.CoverTab[37551]++
								i, err := strconv.ParseInt(end, 10, 64)
								if i < 0 || func() bool {
//line /usr/local/go/src/net/http/fs.go:918
				_go_fuzz_dep_.CoverTab[37557]++
//line /usr/local/go/src/net/http/fs.go:918
				return err != nil
//line /usr/local/go/src/net/http/fs.go:918
				// _ = "end of CoverTab[37557]"
//line /usr/local/go/src/net/http/fs.go:918
			}() {
//line /usr/local/go/src/net/http/fs.go:918
				_go_fuzz_dep_.CoverTab[37558]++
									return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:919
				// _ = "end of CoverTab[37558]"
			} else {
//line /usr/local/go/src/net/http/fs.go:920
				_go_fuzz_dep_.CoverTab[37559]++
//line /usr/local/go/src/net/http/fs.go:920
				// _ = "end of CoverTab[37559]"
//line /usr/local/go/src/net/http/fs.go:920
			}
//line /usr/local/go/src/net/http/fs.go:920
			// _ = "end of CoverTab[37551]"
//line /usr/local/go/src/net/http/fs.go:920
			_go_fuzz_dep_.CoverTab[37552]++
								if i > size {
//line /usr/local/go/src/net/http/fs.go:921
				_go_fuzz_dep_.CoverTab[37560]++
									i = size
//line /usr/local/go/src/net/http/fs.go:922
				// _ = "end of CoverTab[37560]"
			} else {
//line /usr/local/go/src/net/http/fs.go:923
				_go_fuzz_dep_.CoverTab[37561]++
//line /usr/local/go/src/net/http/fs.go:923
				// _ = "end of CoverTab[37561]"
//line /usr/local/go/src/net/http/fs.go:923
			}
//line /usr/local/go/src/net/http/fs.go:923
			// _ = "end of CoverTab[37552]"
//line /usr/local/go/src/net/http/fs.go:923
			_go_fuzz_dep_.CoverTab[37553]++
								r.start = size - i
								r.length = size - r.start
//line /usr/local/go/src/net/http/fs.go:925
			// _ = "end of CoverTab[37553]"
		} else {
//line /usr/local/go/src/net/http/fs.go:926
			_go_fuzz_dep_.CoverTab[37562]++
								i, err := strconv.ParseInt(start, 10, 64)
								if err != nil || func() bool {
//line /usr/local/go/src/net/http/fs.go:928
				_go_fuzz_dep_.CoverTab[37565]++
//line /usr/local/go/src/net/http/fs.go:928
				return i < 0
//line /usr/local/go/src/net/http/fs.go:928
				// _ = "end of CoverTab[37565]"
//line /usr/local/go/src/net/http/fs.go:928
			}() {
//line /usr/local/go/src/net/http/fs.go:928
				_go_fuzz_dep_.CoverTab[37566]++
									return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:929
				// _ = "end of CoverTab[37566]"
			} else {
//line /usr/local/go/src/net/http/fs.go:930
				_go_fuzz_dep_.CoverTab[37567]++
//line /usr/local/go/src/net/http/fs.go:930
				// _ = "end of CoverTab[37567]"
//line /usr/local/go/src/net/http/fs.go:930
			}
//line /usr/local/go/src/net/http/fs.go:930
			// _ = "end of CoverTab[37562]"
//line /usr/local/go/src/net/http/fs.go:930
			_go_fuzz_dep_.CoverTab[37563]++
								if i >= size {
//line /usr/local/go/src/net/http/fs.go:931
				_go_fuzz_dep_.CoverTab[37568]++

//line /usr/local/go/src/net/http/fs.go:934
				noOverlap = true
									continue
//line /usr/local/go/src/net/http/fs.go:935
				// _ = "end of CoverTab[37568]"
			} else {
//line /usr/local/go/src/net/http/fs.go:936
				_go_fuzz_dep_.CoverTab[37569]++
//line /usr/local/go/src/net/http/fs.go:936
				// _ = "end of CoverTab[37569]"
//line /usr/local/go/src/net/http/fs.go:936
			}
//line /usr/local/go/src/net/http/fs.go:936
			// _ = "end of CoverTab[37563]"
//line /usr/local/go/src/net/http/fs.go:936
			_go_fuzz_dep_.CoverTab[37564]++
								r.start = i
								if end == "" {
//line /usr/local/go/src/net/http/fs.go:938
				_go_fuzz_dep_.CoverTab[37570]++

									r.length = size - r.start
//line /usr/local/go/src/net/http/fs.go:940
				// _ = "end of CoverTab[37570]"
			} else {
//line /usr/local/go/src/net/http/fs.go:941
				_go_fuzz_dep_.CoverTab[37571]++
									i, err := strconv.ParseInt(end, 10, 64)
									if err != nil || func() bool {
//line /usr/local/go/src/net/http/fs.go:943
					_go_fuzz_dep_.CoverTab[37574]++
//line /usr/local/go/src/net/http/fs.go:943
					return r.start > i
//line /usr/local/go/src/net/http/fs.go:943
					// _ = "end of CoverTab[37574]"
//line /usr/local/go/src/net/http/fs.go:943
				}() {
//line /usr/local/go/src/net/http/fs.go:943
					_go_fuzz_dep_.CoverTab[37575]++
										return nil, errors.New("invalid range")
//line /usr/local/go/src/net/http/fs.go:944
					// _ = "end of CoverTab[37575]"
				} else {
//line /usr/local/go/src/net/http/fs.go:945
					_go_fuzz_dep_.CoverTab[37576]++
//line /usr/local/go/src/net/http/fs.go:945
					// _ = "end of CoverTab[37576]"
//line /usr/local/go/src/net/http/fs.go:945
				}
//line /usr/local/go/src/net/http/fs.go:945
				// _ = "end of CoverTab[37571]"
//line /usr/local/go/src/net/http/fs.go:945
				_go_fuzz_dep_.CoverTab[37572]++
									if i >= size {
//line /usr/local/go/src/net/http/fs.go:946
					_go_fuzz_dep_.CoverTab[37577]++
										i = size - 1
//line /usr/local/go/src/net/http/fs.go:947
					// _ = "end of CoverTab[37577]"
				} else {
//line /usr/local/go/src/net/http/fs.go:948
					_go_fuzz_dep_.CoverTab[37578]++
//line /usr/local/go/src/net/http/fs.go:948
					// _ = "end of CoverTab[37578]"
//line /usr/local/go/src/net/http/fs.go:948
				}
//line /usr/local/go/src/net/http/fs.go:948
				// _ = "end of CoverTab[37572]"
//line /usr/local/go/src/net/http/fs.go:948
				_go_fuzz_dep_.CoverTab[37573]++
									r.length = i - r.start + 1
//line /usr/local/go/src/net/http/fs.go:949
				// _ = "end of CoverTab[37573]"
			}
//line /usr/local/go/src/net/http/fs.go:950
			// _ = "end of CoverTab[37564]"
		}
//line /usr/local/go/src/net/http/fs.go:951
		// _ = "end of CoverTab[37544]"
//line /usr/local/go/src/net/http/fs.go:951
		_go_fuzz_dep_.CoverTab[37545]++
							ranges = append(ranges, r)
//line /usr/local/go/src/net/http/fs.go:952
		// _ = "end of CoverTab[37545]"
	}
//line /usr/local/go/src/net/http/fs.go:953
	// _ = "end of CoverTab[37535]"
//line /usr/local/go/src/net/http/fs.go:953
	_go_fuzz_dep_.CoverTab[37536]++
						if noOverlap && func() bool {
//line /usr/local/go/src/net/http/fs.go:954
		_go_fuzz_dep_.CoverTab[37579]++
//line /usr/local/go/src/net/http/fs.go:954
		return len(ranges) == 0
//line /usr/local/go/src/net/http/fs.go:954
		// _ = "end of CoverTab[37579]"
//line /usr/local/go/src/net/http/fs.go:954
	}() {
//line /usr/local/go/src/net/http/fs.go:954
		_go_fuzz_dep_.CoverTab[37580]++

							return nil, errNoOverlap
//line /usr/local/go/src/net/http/fs.go:956
		// _ = "end of CoverTab[37580]"
	} else {
//line /usr/local/go/src/net/http/fs.go:957
		_go_fuzz_dep_.CoverTab[37581]++
//line /usr/local/go/src/net/http/fs.go:957
		// _ = "end of CoverTab[37581]"
//line /usr/local/go/src/net/http/fs.go:957
	}
//line /usr/local/go/src/net/http/fs.go:957
	// _ = "end of CoverTab[37536]"
//line /usr/local/go/src/net/http/fs.go:957
	_go_fuzz_dep_.CoverTab[37537]++
						return ranges, nil
//line /usr/local/go/src/net/http/fs.go:958
	// _ = "end of CoverTab[37537]"
}

// countingWriter counts how many bytes have been written to it.
type countingWriter int64

func (w *countingWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/fs.go:964
	_go_fuzz_dep_.CoverTab[37582]++
						*w += countingWriter(len(p))
						return len(p), nil
//line /usr/local/go/src/net/http/fs.go:966
	// _ = "end of CoverTab[37582]"
}

// rangesMIMESize returns the number of bytes it takes to encode the
//line /usr/local/go/src/net/http/fs.go:969
// provided ranges as a multipart response.
//line /usr/local/go/src/net/http/fs.go:971
func rangesMIMESize(ranges []httpRange, contentType string, contentSize int64) (encSize int64) {
//line /usr/local/go/src/net/http/fs.go:971
	_go_fuzz_dep_.CoverTab[37583]++
						var w countingWriter
						mw := multipart.NewWriter(&w)
						for _, ra := range ranges {
//line /usr/local/go/src/net/http/fs.go:974
		_go_fuzz_dep_.CoverTab[37585]++
							mw.CreatePart(ra.mimeHeader(contentType, contentSize))
							encSize += ra.length
//line /usr/local/go/src/net/http/fs.go:976
		// _ = "end of CoverTab[37585]"
	}
//line /usr/local/go/src/net/http/fs.go:977
	// _ = "end of CoverTab[37583]"
//line /usr/local/go/src/net/http/fs.go:977
	_go_fuzz_dep_.CoverTab[37584]++
						mw.Close()
						encSize += int64(w)
						return
//line /usr/local/go/src/net/http/fs.go:980
	// _ = "end of CoverTab[37584]"
}

func sumRangesSize(ranges []httpRange) (size int64) {
//line /usr/local/go/src/net/http/fs.go:983
	_go_fuzz_dep_.CoverTab[37586]++
						for _, ra := range ranges {
//line /usr/local/go/src/net/http/fs.go:984
		_go_fuzz_dep_.CoverTab[37588]++
							size += ra.length
//line /usr/local/go/src/net/http/fs.go:985
		// _ = "end of CoverTab[37588]"
	}
//line /usr/local/go/src/net/http/fs.go:986
	// _ = "end of CoverTab[37586]"
//line /usr/local/go/src/net/http/fs.go:986
	_go_fuzz_dep_.CoverTab[37587]++
						return
//line /usr/local/go/src/net/http/fs.go:987
	// _ = "end of CoverTab[37587]"
}

//line /usr/local/go/src/net/http/fs.go:988
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/fs.go:988
var _ = _go_fuzz_dep_.CoverTab
