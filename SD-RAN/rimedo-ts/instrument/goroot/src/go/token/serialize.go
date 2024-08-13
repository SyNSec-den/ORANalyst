// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/go/token/serialize.go:5
package token

//line /usr/local/go/src/go/token/serialize.go:5
import (
//line /usr/local/go/src/go/token/serialize.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/go/token/serialize.go:5
)
//line /usr/local/go/src/go/token/serialize.go:5
import (
//line /usr/local/go/src/go/token/serialize.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/go/token/serialize.go:5
)

type serializedFile struct {
	// fields correspond 1:1 to fields with same (lower-case) name in File
	Name	string
	Base	int
	Size	int
	Lines	[]int
	Infos	[]lineInfo
}

type serializedFileSet struct {
	Base	int
	Files	[]serializedFile
}

// Read calls decode to deserialize a file set into s; s must not be nil.
func (s *FileSet) Read(decode func(any) error) error {
//line /usr/local/go/src/go/token/serialize.go:22
	_go_fuzz_dep_.CoverTab[49254]++
							var ss serializedFileSet
							if err := decode(&ss); err != nil {
//line /usr/local/go/src/go/token/serialize.go:24
		_go_fuzz_dep_.CoverTab[49257]++
								return err
//line /usr/local/go/src/go/token/serialize.go:25
		// _ = "end of CoverTab[49257]"
	} else {
//line /usr/local/go/src/go/token/serialize.go:26
		_go_fuzz_dep_.CoverTab[49258]++
//line /usr/local/go/src/go/token/serialize.go:26
		// _ = "end of CoverTab[49258]"
//line /usr/local/go/src/go/token/serialize.go:26
	}
//line /usr/local/go/src/go/token/serialize.go:26
	// _ = "end of CoverTab[49254]"
//line /usr/local/go/src/go/token/serialize.go:26
	_go_fuzz_dep_.CoverTab[49255]++

							s.mutex.Lock()
							s.base = ss.Base
							files := make([]*File, len(ss.Files))
							for i := 0; i < len(ss.Files); i++ {
//line /usr/local/go/src/go/token/serialize.go:31
		_go_fuzz_dep_.CoverTab[49259]++
								f := &ss.Files[i]
								files[i] = &File{
			name:	f.Name,
			base:	f.Base,
			size:	f.Size,
			lines:	f.Lines,
			infos:	f.Infos,
		}
//line /usr/local/go/src/go/token/serialize.go:39
		// _ = "end of CoverTab[49259]"
	}
//line /usr/local/go/src/go/token/serialize.go:40
	// _ = "end of CoverTab[49255]"
//line /usr/local/go/src/go/token/serialize.go:40
	_go_fuzz_dep_.CoverTab[49256]++
							s.files = files
							s.last.Store(nil)
							s.mutex.Unlock()

							return nil
//line /usr/local/go/src/go/token/serialize.go:45
	// _ = "end of CoverTab[49256]"
}

// Write calls encode to serialize the file set s.
func (s *FileSet) Write(encode func(any) error) error {
//line /usr/local/go/src/go/token/serialize.go:49
	_go_fuzz_dep_.CoverTab[49260]++
							var ss serializedFileSet

							s.mutex.Lock()
							ss.Base = s.base
							files := make([]serializedFile, len(s.files))
							for i, f := range s.files {
//line /usr/local/go/src/go/token/serialize.go:55
		_go_fuzz_dep_.CoverTab[49262]++
								f.mutex.Lock()
								files[i] = serializedFile{
			Name:	f.name,
			Base:	f.base,
			Size:	f.size,
			Lines:	append([]int(nil), f.lines...),
			Infos:	append([]lineInfo(nil), f.infos...),
		}
								f.mutex.Unlock()
//line /usr/local/go/src/go/token/serialize.go:64
		// _ = "end of CoverTab[49262]"
	}
//line /usr/local/go/src/go/token/serialize.go:65
	// _ = "end of CoverTab[49260]"
//line /usr/local/go/src/go/token/serialize.go:65
	_go_fuzz_dep_.CoverTab[49261]++
							ss.Files = files
							s.mutex.Unlock()

							return encode(ss)
//line /usr/local/go/src/go/token/serialize.go:69
	// _ = "end of CoverTab[49261]"
}

//line /usr/local/go/src/go/token/serialize.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/go/token/serialize.go:70
var _ = _go_fuzz_dep_.CoverTab
