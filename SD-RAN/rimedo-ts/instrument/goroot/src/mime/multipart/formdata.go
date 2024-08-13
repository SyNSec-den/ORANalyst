// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/multipart/formdata.go:5
package multipart

//line /usr/local/go/src/mime/multipart/formdata.go:5
import (
//line /usr/local/go/src/mime/multipart/formdata.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/multipart/formdata.go:5
)
//line /usr/local/go/src/mime/multipart/formdata.go:5
import (
//line /usr/local/go/src/mime/multipart/formdata.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/multipart/formdata.go:5
)

import (
	"bytes"
	"errors"
	"internal/godebug"
	"io"
	"math"
	"net/textproto"
	"os"
	"strconv"
)

// ErrMessageTooLarge is returned by ReadForm if the message form
//line /usr/local/go/src/mime/multipart/formdata.go:18
// data is too large to be processed.
//line /usr/local/go/src/mime/multipart/formdata.go:20
var ErrMessageTooLarge = errors.New("multipart: message too large")

//line /usr/local/go/src/mime/multipart/formdata.go:25
// ReadForm parses an entire multipart message whose parts have
//line /usr/local/go/src/mime/multipart/formdata.go:25
// a Content-Disposition of "form-data".
//line /usr/local/go/src/mime/multipart/formdata.go:25
// It stores up to maxMemory bytes + 10MB (reserved for non-file parts)
//line /usr/local/go/src/mime/multipart/formdata.go:25
// in memory. File parts which can't be stored in memory will be stored on
//line /usr/local/go/src/mime/multipart/formdata.go:25
// disk in temporary files.
//line /usr/local/go/src/mime/multipart/formdata.go:25
// It returns ErrMessageTooLarge if all non-file parts can't be stored in
//line /usr/local/go/src/mime/multipart/formdata.go:25
// memory.
//line /usr/local/go/src/mime/multipart/formdata.go:32
func (r *Reader) ReadForm(maxMemory int64) (*Form, error) {
//line /usr/local/go/src/mime/multipart/formdata.go:32
	_go_fuzz_dep_.CoverTab[36055]++
							return r.readForm(maxMemory)
//line /usr/local/go/src/mime/multipart/formdata.go:33
	// _ = "end of CoverTab[36055]"
}

var (
	multipartFiles		= godebug.New("multipartfiles")
	multipartMaxParts	= godebug.New("multipartmaxparts")
)

func (r *Reader) readForm(maxMemory int64) (_ *Form, err error) {
//line /usr/local/go/src/mime/multipart/formdata.go:41
	_go_fuzz_dep_.CoverTab[36056]++
							form := &Form{make(map[string][]string), make(map[string][]*FileHeader)}
							var (
		file	*os.File
		fileOff	int64
	)
	numDiskFiles := 0
	combineFiles := true
	if multipartFiles.Value() == "distinct" {
//line /usr/local/go/src/mime/multipart/formdata.go:49
		_go_fuzz_dep_.CoverTab[36062]++
								combineFiles = false
//line /usr/local/go/src/mime/multipart/formdata.go:50
		// _ = "end of CoverTab[36062]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:51
		_go_fuzz_dep_.CoverTab[36063]++
//line /usr/local/go/src/mime/multipart/formdata.go:51
		// _ = "end of CoverTab[36063]"
//line /usr/local/go/src/mime/multipart/formdata.go:51
	}
//line /usr/local/go/src/mime/multipart/formdata.go:51
	// _ = "end of CoverTab[36056]"
//line /usr/local/go/src/mime/multipart/formdata.go:51
	_go_fuzz_dep_.CoverTab[36057]++
							maxParts := 1000
							if s := multipartMaxParts.Value(); s != "" {
//line /usr/local/go/src/mime/multipart/formdata.go:53
		_go_fuzz_dep_.CoverTab[36064]++
								if v, err := strconv.Atoi(s); err == nil && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:54
			_go_fuzz_dep_.CoverTab[36065]++
//line /usr/local/go/src/mime/multipart/formdata.go:54
			return v >= 0
//line /usr/local/go/src/mime/multipart/formdata.go:54
			// _ = "end of CoverTab[36065]"
//line /usr/local/go/src/mime/multipart/formdata.go:54
		}() {
//line /usr/local/go/src/mime/multipart/formdata.go:54
			_go_fuzz_dep_.CoverTab[36066]++
									maxParts = v
//line /usr/local/go/src/mime/multipart/formdata.go:55
			// _ = "end of CoverTab[36066]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:56
			_go_fuzz_dep_.CoverTab[36067]++
//line /usr/local/go/src/mime/multipart/formdata.go:56
			// _ = "end of CoverTab[36067]"
//line /usr/local/go/src/mime/multipart/formdata.go:56
		}
//line /usr/local/go/src/mime/multipart/formdata.go:56
		// _ = "end of CoverTab[36064]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:57
		_go_fuzz_dep_.CoverTab[36068]++
//line /usr/local/go/src/mime/multipart/formdata.go:57
		// _ = "end of CoverTab[36068]"
//line /usr/local/go/src/mime/multipart/formdata.go:57
	}
//line /usr/local/go/src/mime/multipart/formdata.go:57
	// _ = "end of CoverTab[36057]"
//line /usr/local/go/src/mime/multipart/formdata.go:57
	_go_fuzz_dep_.CoverTab[36058]++
							maxHeaders := maxMIMEHeaders()

							defer func() {
//line /usr/local/go/src/mime/multipart/formdata.go:60
		_go_fuzz_dep_.CoverTab[36069]++
								if file != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:61
			_go_fuzz_dep_.CoverTab[36072]++
									if cerr := file.Close(); err == nil {
//line /usr/local/go/src/mime/multipart/formdata.go:62
				_go_fuzz_dep_.CoverTab[36073]++
										err = cerr
//line /usr/local/go/src/mime/multipart/formdata.go:63
				// _ = "end of CoverTab[36073]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:64
				_go_fuzz_dep_.CoverTab[36074]++
//line /usr/local/go/src/mime/multipart/formdata.go:64
				// _ = "end of CoverTab[36074]"
//line /usr/local/go/src/mime/multipart/formdata.go:64
			}
//line /usr/local/go/src/mime/multipart/formdata.go:64
			// _ = "end of CoverTab[36072]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:65
			_go_fuzz_dep_.CoverTab[36075]++
//line /usr/local/go/src/mime/multipart/formdata.go:65
			// _ = "end of CoverTab[36075]"
//line /usr/local/go/src/mime/multipart/formdata.go:65
		}
//line /usr/local/go/src/mime/multipart/formdata.go:65
		// _ = "end of CoverTab[36069]"
//line /usr/local/go/src/mime/multipart/formdata.go:65
		_go_fuzz_dep_.CoverTab[36070]++
								if combineFiles && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:66
			_go_fuzz_dep_.CoverTab[36076]++
//line /usr/local/go/src/mime/multipart/formdata.go:66
			return numDiskFiles > 1
//line /usr/local/go/src/mime/multipart/formdata.go:66
			// _ = "end of CoverTab[36076]"
//line /usr/local/go/src/mime/multipart/formdata.go:66
		}() {
//line /usr/local/go/src/mime/multipart/formdata.go:66
			_go_fuzz_dep_.CoverTab[36077]++
									for _, fhs := range form.File {
//line /usr/local/go/src/mime/multipart/formdata.go:67
				_go_fuzz_dep_.CoverTab[36078]++
										for _, fh := range fhs {
//line /usr/local/go/src/mime/multipart/formdata.go:68
					_go_fuzz_dep_.CoverTab[36079]++
											fh.tmpshared = true
//line /usr/local/go/src/mime/multipart/formdata.go:69
					// _ = "end of CoverTab[36079]"
				}
//line /usr/local/go/src/mime/multipart/formdata.go:70
				// _ = "end of CoverTab[36078]"
			}
//line /usr/local/go/src/mime/multipart/formdata.go:71
			// _ = "end of CoverTab[36077]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:72
			_go_fuzz_dep_.CoverTab[36080]++
//line /usr/local/go/src/mime/multipart/formdata.go:72
			// _ = "end of CoverTab[36080]"
//line /usr/local/go/src/mime/multipart/formdata.go:72
		}
//line /usr/local/go/src/mime/multipart/formdata.go:72
		// _ = "end of CoverTab[36070]"
//line /usr/local/go/src/mime/multipart/formdata.go:72
		_go_fuzz_dep_.CoverTab[36071]++
								if err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:73
			_go_fuzz_dep_.CoverTab[36081]++
									form.RemoveAll()
									if file != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:75
				_go_fuzz_dep_.CoverTab[36082]++
										os.Remove(file.Name())
//line /usr/local/go/src/mime/multipart/formdata.go:76
				// _ = "end of CoverTab[36082]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:77
				_go_fuzz_dep_.CoverTab[36083]++
//line /usr/local/go/src/mime/multipart/formdata.go:77
				// _ = "end of CoverTab[36083]"
//line /usr/local/go/src/mime/multipart/formdata.go:77
			}
//line /usr/local/go/src/mime/multipart/formdata.go:77
			// _ = "end of CoverTab[36081]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:78
			_go_fuzz_dep_.CoverTab[36084]++
//line /usr/local/go/src/mime/multipart/formdata.go:78
			// _ = "end of CoverTab[36084]"
//line /usr/local/go/src/mime/multipart/formdata.go:78
		}
//line /usr/local/go/src/mime/multipart/formdata.go:78
		// _ = "end of CoverTab[36071]"
	}()
//line /usr/local/go/src/mime/multipart/formdata.go:79
	// _ = "end of CoverTab[36058]"
//line /usr/local/go/src/mime/multipart/formdata.go:79
	_go_fuzz_dep_.CoverTab[36059]++

//line /usr/local/go/src/mime/multipart/formdata.go:94
	maxFileMemoryBytes := maxMemory
	maxMemoryBytes := maxMemory + int64(10<<20)
	if maxMemoryBytes <= 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:96
		_go_fuzz_dep_.CoverTab[36085]++
								if maxMemory < 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:97
			_go_fuzz_dep_.CoverTab[36086]++
									maxMemoryBytes = 0
//line /usr/local/go/src/mime/multipart/formdata.go:98
			// _ = "end of CoverTab[36086]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:99
				_go_fuzz_dep_.CoverTab[36087]++
										maxMemoryBytes = math.MaxInt64
//line /usr/local/go/src/mime/multipart/formdata.go:100
			// _ = "end of CoverTab[36087]"
		}
//line /usr/local/go/src/mime/multipart/formdata.go:101
		// _ = "end of CoverTab[36085]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:102
		_go_fuzz_dep_.CoverTab[36088]++
//line /usr/local/go/src/mime/multipart/formdata.go:102
		// _ = "end of CoverTab[36088]"
//line /usr/local/go/src/mime/multipart/formdata.go:102
	}
//line /usr/local/go/src/mime/multipart/formdata.go:102
	// _ = "end of CoverTab[36059]"
//line /usr/local/go/src/mime/multipart/formdata.go:102
	_go_fuzz_dep_.CoverTab[36060]++
								var copyBuf []byte
								for {
//line /usr/local/go/src/mime/multipart/formdata.go:104
		_go_fuzz_dep_.CoverTab[36089]++
									p, err := r.nextPart(false, maxMemoryBytes, maxHeaders)
									if err == io.EOF {
//line /usr/local/go/src/mime/multipart/formdata.go:106
			_go_fuzz_dep_.CoverTab[36100]++
										break
//line /usr/local/go/src/mime/multipart/formdata.go:107
			// _ = "end of CoverTab[36100]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:108
			_go_fuzz_dep_.CoverTab[36101]++
//line /usr/local/go/src/mime/multipart/formdata.go:108
			// _ = "end of CoverTab[36101]"
//line /usr/local/go/src/mime/multipart/formdata.go:108
		}
//line /usr/local/go/src/mime/multipart/formdata.go:108
		// _ = "end of CoverTab[36089]"
//line /usr/local/go/src/mime/multipart/formdata.go:108
		_go_fuzz_dep_.CoverTab[36090]++
									if err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:109
			_go_fuzz_dep_.CoverTab[36102]++
										return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:110
			// _ = "end of CoverTab[36102]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:111
			_go_fuzz_dep_.CoverTab[36103]++
//line /usr/local/go/src/mime/multipart/formdata.go:111
			// _ = "end of CoverTab[36103]"
//line /usr/local/go/src/mime/multipart/formdata.go:111
		}
//line /usr/local/go/src/mime/multipart/formdata.go:111
		// _ = "end of CoverTab[36090]"
//line /usr/local/go/src/mime/multipart/formdata.go:111
		_go_fuzz_dep_.CoverTab[36091]++
									if maxParts <= 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:112
			_go_fuzz_dep_.CoverTab[36104]++
										return nil, ErrMessageTooLarge
//line /usr/local/go/src/mime/multipart/formdata.go:113
			// _ = "end of CoverTab[36104]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:114
			_go_fuzz_dep_.CoverTab[36105]++
//line /usr/local/go/src/mime/multipart/formdata.go:114
			// _ = "end of CoverTab[36105]"
//line /usr/local/go/src/mime/multipart/formdata.go:114
		}
//line /usr/local/go/src/mime/multipart/formdata.go:114
		// _ = "end of CoverTab[36091]"
//line /usr/local/go/src/mime/multipart/formdata.go:114
		_go_fuzz_dep_.CoverTab[36092]++
									maxParts--

									name := p.FormName()
									if name == "" {
//line /usr/local/go/src/mime/multipart/formdata.go:118
			_go_fuzz_dep_.CoverTab[36106]++
										continue
//line /usr/local/go/src/mime/multipart/formdata.go:119
			// _ = "end of CoverTab[36106]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:120
			_go_fuzz_dep_.CoverTab[36107]++
//line /usr/local/go/src/mime/multipart/formdata.go:120
			// _ = "end of CoverTab[36107]"
//line /usr/local/go/src/mime/multipart/formdata.go:120
		}
//line /usr/local/go/src/mime/multipart/formdata.go:120
		// _ = "end of CoverTab[36092]"
//line /usr/local/go/src/mime/multipart/formdata.go:120
		_go_fuzz_dep_.CoverTab[36093]++
									filename := p.FileName()

		// Multiple values for the same key (one map entry, longer slice) are cheaper
		// than the same number of values for different keys (many map entries), but
		// using a consistent per-value cost for overhead is simpler.
		const mapEntryOverhead = 200
		maxMemoryBytes -= int64(len(name))
		maxMemoryBytes -= mapEntryOverhead
		if maxMemoryBytes < 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:129
			_go_fuzz_dep_.CoverTab[36108]++

//line /usr/local/go/src/mime/multipart/formdata.go:132
			return nil, ErrMessageTooLarge
//line /usr/local/go/src/mime/multipart/formdata.go:132
			// _ = "end of CoverTab[36108]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:133
			_go_fuzz_dep_.CoverTab[36109]++
//line /usr/local/go/src/mime/multipart/formdata.go:133
			// _ = "end of CoverTab[36109]"
//line /usr/local/go/src/mime/multipart/formdata.go:133
		}
//line /usr/local/go/src/mime/multipart/formdata.go:133
		// _ = "end of CoverTab[36093]"
//line /usr/local/go/src/mime/multipart/formdata.go:133
		_go_fuzz_dep_.CoverTab[36094]++

									var b bytes.Buffer

									if filename == "" {
//line /usr/local/go/src/mime/multipart/formdata.go:137
			_go_fuzz_dep_.CoverTab[36110]++

										n, err := io.CopyN(&b, p, maxMemoryBytes+1)
										if err != nil && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:140
				_go_fuzz_dep_.CoverTab[36113]++
//line /usr/local/go/src/mime/multipart/formdata.go:140
				return err != io.EOF
//line /usr/local/go/src/mime/multipart/formdata.go:140
				// _ = "end of CoverTab[36113]"
//line /usr/local/go/src/mime/multipart/formdata.go:140
			}() {
//line /usr/local/go/src/mime/multipart/formdata.go:140
				_go_fuzz_dep_.CoverTab[36114]++
											return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:141
				// _ = "end of CoverTab[36114]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:142
				_go_fuzz_dep_.CoverTab[36115]++
//line /usr/local/go/src/mime/multipart/formdata.go:142
				// _ = "end of CoverTab[36115]"
//line /usr/local/go/src/mime/multipart/formdata.go:142
			}
//line /usr/local/go/src/mime/multipart/formdata.go:142
			// _ = "end of CoverTab[36110]"
//line /usr/local/go/src/mime/multipart/formdata.go:142
			_go_fuzz_dep_.CoverTab[36111]++
										maxMemoryBytes -= n
										if maxMemoryBytes < 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:144
				_go_fuzz_dep_.CoverTab[36116]++
											return nil, ErrMessageTooLarge
//line /usr/local/go/src/mime/multipart/formdata.go:145
				// _ = "end of CoverTab[36116]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:146
				_go_fuzz_dep_.CoverTab[36117]++
//line /usr/local/go/src/mime/multipart/formdata.go:146
				// _ = "end of CoverTab[36117]"
//line /usr/local/go/src/mime/multipart/formdata.go:146
			}
//line /usr/local/go/src/mime/multipart/formdata.go:146
			// _ = "end of CoverTab[36111]"
//line /usr/local/go/src/mime/multipart/formdata.go:146
			_go_fuzz_dep_.CoverTab[36112]++
										form.Value[name] = append(form.Value[name], b.String())
										continue
//line /usr/local/go/src/mime/multipart/formdata.go:148
			// _ = "end of CoverTab[36112]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:149
			_go_fuzz_dep_.CoverTab[36118]++
//line /usr/local/go/src/mime/multipart/formdata.go:149
			// _ = "end of CoverTab[36118]"
//line /usr/local/go/src/mime/multipart/formdata.go:149
		}
//line /usr/local/go/src/mime/multipart/formdata.go:149
		// _ = "end of CoverTab[36094]"
//line /usr/local/go/src/mime/multipart/formdata.go:149
		_go_fuzz_dep_.CoverTab[36095]++

		// file, store in memory or on disk
		const fileHeaderSize = 100
		maxMemoryBytes -= mimeHeaderSize(p.Header)
		maxMemoryBytes -= mapEntryOverhead
		maxMemoryBytes -= fileHeaderSize
		if maxMemoryBytes < 0 {
//line /usr/local/go/src/mime/multipart/formdata.go:156
			_go_fuzz_dep_.CoverTab[36119]++
										return nil, ErrMessageTooLarge
//line /usr/local/go/src/mime/multipart/formdata.go:157
			// _ = "end of CoverTab[36119]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:158
			_go_fuzz_dep_.CoverTab[36120]++
//line /usr/local/go/src/mime/multipart/formdata.go:158
			// _ = "end of CoverTab[36120]"
//line /usr/local/go/src/mime/multipart/formdata.go:158
		}
//line /usr/local/go/src/mime/multipart/formdata.go:158
		// _ = "end of CoverTab[36095]"
//line /usr/local/go/src/mime/multipart/formdata.go:158
		_go_fuzz_dep_.CoverTab[36096]++
									for _, v := range p.Header {
//line /usr/local/go/src/mime/multipart/formdata.go:159
			_go_fuzz_dep_.CoverTab[36121]++
										maxHeaders -= int64(len(v))
//line /usr/local/go/src/mime/multipart/formdata.go:160
			// _ = "end of CoverTab[36121]"
		}
//line /usr/local/go/src/mime/multipart/formdata.go:161
		// _ = "end of CoverTab[36096]"
//line /usr/local/go/src/mime/multipart/formdata.go:161
		_go_fuzz_dep_.CoverTab[36097]++
									fh := &FileHeader{
			Filename:	filename,
			Header:		p.Header,
		}
		n, err := io.CopyN(&b, p, maxFileMemoryBytes+1)
		if err != nil && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:167
			_go_fuzz_dep_.CoverTab[36122]++
//line /usr/local/go/src/mime/multipart/formdata.go:167
			return err != io.EOF
//line /usr/local/go/src/mime/multipart/formdata.go:167
			// _ = "end of CoverTab[36122]"
//line /usr/local/go/src/mime/multipart/formdata.go:167
		}() {
//line /usr/local/go/src/mime/multipart/formdata.go:167
			_go_fuzz_dep_.CoverTab[36123]++
										return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:168
			// _ = "end of CoverTab[36123]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:169
			_go_fuzz_dep_.CoverTab[36124]++
//line /usr/local/go/src/mime/multipart/formdata.go:169
			// _ = "end of CoverTab[36124]"
//line /usr/local/go/src/mime/multipart/formdata.go:169
		}
//line /usr/local/go/src/mime/multipart/formdata.go:169
		// _ = "end of CoverTab[36097]"
//line /usr/local/go/src/mime/multipart/formdata.go:169
		_go_fuzz_dep_.CoverTab[36098]++
									if n > maxFileMemoryBytes {
//line /usr/local/go/src/mime/multipart/formdata.go:170
			_go_fuzz_dep_.CoverTab[36125]++
										if file == nil {
//line /usr/local/go/src/mime/multipart/formdata.go:171
				_go_fuzz_dep_.CoverTab[36130]++
											file, err = os.CreateTemp(r.tempDir, "multipart-")
											if err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:173
					_go_fuzz_dep_.CoverTab[36131]++
												return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:174
					// _ = "end of CoverTab[36131]"
				} else {
//line /usr/local/go/src/mime/multipart/formdata.go:175
					_go_fuzz_dep_.CoverTab[36132]++
//line /usr/local/go/src/mime/multipart/formdata.go:175
					// _ = "end of CoverTab[36132]"
//line /usr/local/go/src/mime/multipart/formdata.go:175
				}
//line /usr/local/go/src/mime/multipart/formdata.go:175
				// _ = "end of CoverTab[36130]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:176
				_go_fuzz_dep_.CoverTab[36133]++
//line /usr/local/go/src/mime/multipart/formdata.go:176
				// _ = "end of CoverTab[36133]"
//line /usr/local/go/src/mime/multipart/formdata.go:176
			}
//line /usr/local/go/src/mime/multipart/formdata.go:176
			// _ = "end of CoverTab[36125]"
//line /usr/local/go/src/mime/multipart/formdata.go:176
			_go_fuzz_dep_.CoverTab[36126]++
										numDiskFiles++
										if _, err := file.Write(b.Bytes()); err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:178
				_go_fuzz_dep_.CoverTab[36134]++
											return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:179
				// _ = "end of CoverTab[36134]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:180
				_go_fuzz_dep_.CoverTab[36135]++
//line /usr/local/go/src/mime/multipart/formdata.go:180
				// _ = "end of CoverTab[36135]"
//line /usr/local/go/src/mime/multipart/formdata.go:180
			}
//line /usr/local/go/src/mime/multipart/formdata.go:180
			// _ = "end of CoverTab[36126]"
//line /usr/local/go/src/mime/multipart/formdata.go:180
			_go_fuzz_dep_.CoverTab[36127]++
										if copyBuf == nil {
//line /usr/local/go/src/mime/multipart/formdata.go:181
				_go_fuzz_dep_.CoverTab[36136]++
											copyBuf = make([]byte, 32*1024)
//line /usr/local/go/src/mime/multipart/formdata.go:182
				// _ = "end of CoverTab[36136]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:183
				_go_fuzz_dep_.CoverTab[36137]++
//line /usr/local/go/src/mime/multipart/formdata.go:183
				// _ = "end of CoverTab[36137]"
//line /usr/local/go/src/mime/multipart/formdata.go:183
			}
//line /usr/local/go/src/mime/multipart/formdata.go:183
			// _ = "end of CoverTab[36127]"
//line /usr/local/go/src/mime/multipart/formdata.go:183
			_go_fuzz_dep_.CoverTab[36128]++
			// os.File.ReadFrom will allocate its own copy buffer if we let io.Copy use it.
			type writerOnly struct{ io.Writer }
			remainingSize, err := io.CopyBuffer(writerOnly{file}, p, copyBuf)
			if err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:187
				_go_fuzz_dep_.CoverTab[36138]++
											return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:188
				// _ = "end of CoverTab[36138]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:189
				_go_fuzz_dep_.CoverTab[36139]++
//line /usr/local/go/src/mime/multipart/formdata.go:189
				// _ = "end of CoverTab[36139]"
//line /usr/local/go/src/mime/multipart/formdata.go:189
			}
//line /usr/local/go/src/mime/multipart/formdata.go:189
			// _ = "end of CoverTab[36128]"
//line /usr/local/go/src/mime/multipart/formdata.go:189
			_go_fuzz_dep_.CoverTab[36129]++
										fh.tmpfile = file.Name()
										fh.Size = int64(b.Len()) + remainingSize
										fh.tmpoff = fileOff
										fileOff += fh.Size
										if !combineFiles {
//line /usr/local/go/src/mime/multipart/formdata.go:194
				_go_fuzz_dep_.CoverTab[36140]++
											if err := file.Close(); err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:195
					_go_fuzz_dep_.CoverTab[36142]++
												return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:196
					// _ = "end of CoverTab[36142]"
				} else {
//line /usr/local/go/src/mime/multipart/formdata.go:197
					_go_fuzz_dep_.CoverTab[36143]++
//line /usr/local/go/src/mime/multipart/formdata.go:197
					// _ = "end of CoverTab[36143]"
//line /usr/local/go/src/mime/multipart/formdata.go:197
				}
//line /usr/local/go/src/mime/multipart/formdata.go:197
				// _ = "end of CoverTab[36140]"
//line /usr/local/go/src/mime/multipart/formdata.go:197
				_go_fuzz_dep_.CoverTab[36141]++
											file = nil
//line /usr/local/go/src/mime/multipart/formdata.go:198
				// _ = "end of CoverTab[36141]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:199
				_go_fuzz_dep_.CoverTab[36144]++
//line /usr/local/go/src/mime/multipart/formdata.go:199
				// _ = "end of CoverTab[36144]"
//line /usr/local/go/src/mime/multipart/formdata.go:199
			}
//line /usr/local/go/src/mime/multipart/formdata.go:199
			// _ = "end of CoverTab[36129]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:200
			_go_fuzz_dep_.CoverTab[36145]++
										fh.content = b.Bytes()
										fh.Size = int64(len(fh.content))
										maxFileMemoryBytes -= n
										maxMemoryBytes -= n
//line /usr/local/go/src/mime/multipart/formdata.go:204
			// _ = "end of CoverTab[36145]"
		}
//line /usr/local/go/src/mime/multipart/formdata.go:205
		// _ = "end of CoverTab[36098]"
//line /usr/local/go/src/mime/multipart/formdata.go:205
		_go_fuzz_dep_.CoverTab[36099]++
									form.File[name] = append(form.File[name], fh)
//line /usr/local/go/src/mime/multipart/formdata.go:206
		// _ = "end of CoverTab[36099]"
	}
//line /usr/local/go/src/mime/multipart/formdata.go:207
	// _ = "end of CoverTab[36060]"
//line /usr/local/go/src/mime/multipart/formdata.go:207
	_go_fuzz_dep_.CoverTab[36061]++

								return form, nil
//line /usr/local/go/src/mime/multipart/formdata.go:209
	// _ = "end of CoverTab[36061]"
}

func mimeHeaderSize(h textproto.MIMEHeader) (size int64) {
//line /usr/local/go/src/mime/multipart/formdata.go:212
	_go_fuzz_dep_.CoverTab[36146]++
								size = 400
								for k, vs := range h {
//line /usr/local/go/src/mime/multipart/formdata.go:214
		_go_fuzz_dep_.CoverTab[36148]++
									size += int64(len(k))
									size += 200
									for _, v := range vs {
//line /usr/local/go/src/mime/multipart/formdata.go:217
			_go_fuzz_dep_.CoverTab[36149]++
										size += int64(len(v))
//line /usr/local/go/src/mime/multipart/formdata.go:218
			// _ = "end of CoverTab[36149]"
		}
//line /usr/local/go/src/mime/multipart/formdata.go:219
		// _ = "end of CoverTab[36148]"
	}
//line /usr/local/go/src/mime/multipart/formdata.go:220
	// _ = "end of CoverTab[36146]"
//line /usr/local/go/src/mime/multipart/formdata.go:220
	_go_fuzz_dep_.CoverTab[36147]++
								return size
//line /usr/local/go/src/mime/multipart/formdata.go:221
	// _ = "end of CoverTab[36147]"
}

// Form is a parsed multipart form.
//line /usr/local/go/src/mime/multipart/formdata.go:224
// Its File parts are stored either in memory or on disk,
//line /usr/local/go/src/mime/multipart/formdata.go:224
// and are accessible via the *FileHeader's Open method.
//line /usr/local/go/src/mime/multipart/formdata.go:224
// Its Value parts are stored as strings.
//line /usr/local/go/src/mime/multipart/formdata.go:224
// Both are keyed by field name.
//line /usr/local/go/src/mime/multipart/formdata.go:229
type Form struct {
	Value	map[string][]string
	File	map[string][]*FileHeader
}

// RemoveAll removes any temporary files associated with a Form.
func (f *Form) RemoveAll() error {
//line /usr/local/go/src/mime/multipart/formdata.go:235
	_go_fuzz_dep_.CoverTab[36150]++
								var err error
								for _, fhs := range f.File {
//line /usr/local/go/src/mime/multipart/formdata.go:237
		_go_fuzz_dep_.CoverTab[36152]++
									for _, fh := range fhs {
//line /usr/local/go/src/mime/multipart/formdata.go:238
			_go_fuzz_dep_.CoverTab[36153]++
										if fh.tmpfile != "" {
//line /usr/local/go/src/mime/multipart/formdata.go:239
				_go_fuzz_dep_.CoverTab[36154]++
											e := os.Remove(fh.tmpfile)
											if e != nil && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:241
					_go_fuzz_dep_.CoverTab[36155]++
//line /usr/local/go/src/mime/multipart/formdata.go:241
					return !errors.Is(e, os.ErrNotExist)
//line /usr/local/go/src/mime/multipart/formdata.go:241
					// _ = "end of CoverTab[36155]"
//line /usr/local/go/src/mime/multipart/formdata.go:241
				}() && func() bool {
//line /usr/local/go/src/mime/multipart/formdata.go:241
					_go_fuzz_dep_.CoverTab[36156]++
//line /usr/local/go/src/mime/multipart/formdata.go:241
					return err == nil
//line /usr/local/go/src/mime/multipart/formdata.go:241
					// _ = "end of CoverTab[36156]"
//line /usr/local/go/src/mime/multipart/formdata.go:241
				}() {
//line /usr/local/go/src/mime/multipart/formdata.go:241
					_go_fuzz_dep_.CoverTab[36157]++
												err = e
//line /usr/local/go/src/mime/multipart/formdata.go:242
					// _ = "end of CoverTab[36157]"
				} else {
//line /usr/local/go/src/mime/multipart/formdata.go:243
					_go_fuzz_dep_.CoverTab[36158]++
//line /usr/local/go/src/mime/multipart/formdata.go:243
					// _ = "end of CoverTab[36158]"
//line /usr/local/go/src/mime/multipart/formdata.go:243
				}
//line /usr/local/go/src/mime/multipart/formdata.go:243
				// _ = "end of CoverTab[36154]"
			} else {
//line /usr/local/go/src/mime/multipart/formdata.go:244
				_go_fuzz_dep_.CoverTab[36159]++
//line /usr/local/go/src/mime/multipart/formdata.go:244
				// _ = "end of CoverTab[36159]"
//line /usr/local/go/src/mime/multipart/formdata.go:244
			}
//line /usr/local/go/src/mime/multipart/formdata.go:244
			// _ = "end of CoverTab[36153]"
		}
//line /usr/local/go/src/mime/multipart/formdata.go:245
		// _ = "end of CoverTab[36152]"
	}
//line /usr/local/go/src/mime/multipart/formdata.go:246
	// _ = "end of CoverTab[36150]"
//line /usr/local/go/src/mime/multipart/formdata.go:246
	_go_fuzz_dep_.CoverTab[36151]++
								return err
//line /usr/local/go/src/mime/multipart/formdata.go:247
	// _ = "end of CoverTab[36151]"
}

// A FileHeader describes a file part of a multipart request.
type FileHeader struct {
	Filename	string
	Header		textproto.MIMEHeader
	Size		int64

	content		[]byte
	tmpfile		string
	tmpoff		int64
	tmpshared	bool
}

// Open opens and returns the FileHeader's associated File.
func (fh *FileHeader) Open() (File, error) {
//line /usr/local/go/src/mime/multipart/formdata.go:263
	_go_fuzz_dep_.CoverTab[36160]++
								if b := fh.content; b != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:264
		_go_fuzz_dep_.CoverTab[36163]++
									r := io.NewSectionReader(bytes.NewReader(b), 0, int64(len(b)))
									return sectionReadCloser{r, nil}, nil
//line /usr/local/go/src/mime/multipart/formdata.go:266
		// _ = "end of CoverTab[36163]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:267
		_go_fuzz_dep_.CoverTab[36164]++
//line /usr/local/go/src/mime/multipart/formdata.go:267
		// _ = "end of CoverTab[36164]"
//line /usr/local/go/src/mime/multipart/formdata.go:267
	}
//line /usr/local/go/src/mime/multipart/formdata.go:267
	// _ = "end of CoverTab[36160]"
//line /usr/local/go/src/mime/multipart/formdata.go:267
	_go_fuzz_dep_.CoverTab[36161]++
								if fh.tmpshared {
//line /usr/local/go/src/mime/multipart/formdata.go:268
		_go_fuzz_dep_.CoverTab[36165]++
									f, err := os.Open(fh.tmpfile)
									if err != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:270
			_go_fuzz_dep_.CoverTab[36167]++
										return nil, err
//line /usr/local/go/src/mime/multipart/formdata.go:271
			// _ = "end of CoverTab[36167]"
		} else {
//line /usr/local/go/src/mime/multipart/formdata.go:272
			_go_fuzz_dep_.CoverTab[36168]++
//line /usr/local/go/src/mime/multipart/formdata.go:272
			// _ = "end of CoverTab[36168]"
//line /usr/local/go/src/mime/multipart/formdata.go:272
		}
//line /usr/local/go/src/mime/multipart/formdata.go:272
		// _ = "end of CoverTab[36165]"
//line /usr/local/go/src/mime/multipart/formdata.go:272
		_go_fuzz_dep_.CoverTab[36166]++
									r := io.NewSectionReader(f, fh.tmpoff, fh.Size)
									return sectionReadCloser{r, f}, nil
//line /usr/local/go/src/mime/multipart/formdata.go:274
		// _ = "end of CoverTab[36166]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:275
		_go_fuzz_dep_.CoverTab[36169]++
//line /usr/local/go/src/mime/multipart/formdata.go:275
		// _ = "end of CoverTab[36169]"
//line /usr/local/go/src/mime/multipart/formdata.go:275
	}
//line /usr/local/go/src/mime/multipart/formdata.go:275
	// _ = "end of CoverTab[36161]"
//line /usr/local/go/src/mime/multipart/formdata.go:275
	_go_fuzz_dep_.CoverTab[36162]++
								return os.Open(fh.tmpfile)
//line /usr/local/go/src/mime/multipart/formdata.go:276
	// _ = "end of CoverTab[36162]"
}

// File is an interface to access the file part of a multipart message.
//line /usr/local/go/src/mime/multipart/formdata.go:279
// Its contents may be either stored in memory or on disk.
//line /usr/local/go/src/mime/multipart/formdata.go:279
// If stored on disk, the File's underlying concrete type will be an *os.File.
//line /usr/local/go/src/mime/multipart/formdata.go:282
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

//line /usr/local/go/src/mime/multipart/formdata.go:291
type sectionReadCloser struct {
	*io.SectionReader
	io.Closer
}

func (rc sectionReadCloser) Close() error {
//line /usr/local/go/src/mime/multipart/formdata.go:296
	_go_fuzz_dep_.CoverTab[36170]++
								if rc.Closer != nil {
//line /usr/local/go/src/mime/multipart/formdata.go:297
		_go_fuzz_dep_.CoverTab[36172]++
									return rc.Closer.Close()
//line /usr/local/go/src/mime/multipart/formdata.go:298
		// _ = "end of CoverTab[36172]"
	} else {
//line /usr/local/go/src/mime/multipart/formdata.go:299
		_go_fuzz_dep_.CoverTab[36173]++
//line /usr/local/go/src/mime/multipart/formdata.go:299
		// _ = "end of CoverTab[36173]"
//line /usr/local/go/src/mime/multipart/formdata.go:299
	}
//line /usr/local/go/src/mime/multipart/formdata.go:299
	// _ = "end of CoverTab[36170]"
//line /usr/local/go/src/mime/multipart/formdata.go:299
	_go_fuzz_dep_.CoverTab[36171]++
								return nil
//line /usr/local/go/src/mime/multipart/formdata.go:300
	// _ = "end of CoverTab[36171]"
}

//line /usr/local/go/src/mime/multipart/formdata.go:301
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/multipart/formdata.go:301
var _ = _go_fuzz_dep_.CoverTab
