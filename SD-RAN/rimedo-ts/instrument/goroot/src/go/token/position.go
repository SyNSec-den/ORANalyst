// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/go/token/position.go:5
package token

//line /usr/local/go/src/go/token/position.go:5
import (
//line /usr/local/go/src/go/token/position.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/go/token/position.go:5
)
//line /usr/local/go/src/go/token/position.go:5
import (
//line /usr/local/go/src/go/token/position.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/go/token/position.go:5
)

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

//line /usr/local/go/src/go/token/position.go:17
// Position describes an arbitrary source position
//line /usr/local/go/src/go/token/position.go:17
// including the file, line, and column location.
//line /usr/local/go/src/go/token/position.go:17
// A Position is valid if the line number is > 0.
//line /usr/local/go/src/go/token/position.go:20
type Position struct {
	Filename	string	// filename, if any
	Offset		int	// offset, starting at 0
	Line		int	// line number, starting at 1
	Column		int	// column number, starting at 1 (byte count)
}

// IsValid reports whether the position is valid.
func (pos *Position) IsValid() bool {
//line /usr/local/go/src/go/token/position.go:28
	_go_fuzz_dep_.CoverTab[49100]++
//line /usr/local/go/src/go/token/position.go:28
	return pos.Line > 0
//line /usr/local/go/src/go/token/position.go:28
	// _ = "end of CoverTab[49100]"
//line /usr/local/go/src/go/token/position.go:28
}

// String returns a string in one of several forms:
//line /usr/local/go/src/go/token/position.go:30
//
//line /usr/local/go/src/go/token/position.go:30
//	file:line:column    valid position with file name
//line /usr/local/go/src/go/token/position.go:30
//	file:line           valid position with file name but no column (column == 0)
//line /usr/local/go/src/go/token/position.go:30
//	line:column         valid position without file name
//line /usr/local/go/src/go/token/position.go:30
//	line                valid position without file name and no column (column == 0)
//line /usr/local/go/src/go/token/position.go:30
//	file                invalid position with file name
//line /usr/local/go/src/go/token/position.go:30
//	-                   invalid position without file name
//line /usr/local/go/src/go/token/position.go:38
func (pos Position) String() string {
//line /usr/local/go/src/go/token/position.go:38
	_go_fuzz_dep_.CoverTab[49101]++
							s := pos.Filename
							if pos.IsValid() {
//line /usr/local/go/src/go/token/position.go:40
		_go_fuzz_dep_.CoverTab[49104]++
								if s != "" {
//line /usr/local/go/src/go/token/position.go:41
			_go_fuzz_dep_.CoverTab[49106]++
									s += ":"
//line /usr/local/go/src/go/token/position.go:42
			// _ = "end of CoverTab[49106]"
		} else {
//line /usr/local/go/src/go/token/position.go:43
			_go_fuzz_dep_.CoverTab[49107]++
//line /usr/local/go/src/go/token/position.go:43
			// _ = "end of CoverTab[49107]"
//line /usr/local/go/src/go/token/position.go:43
		}
//line /usr/local/go/src/go/token/position.go:43
		// _ = "end of CoverTab[49104]"
//line /usr/local/go/src/go/token/position.go:43
		_go_fuzz_dep_.CoverTab[49105]++
								s += fmt.Sprintf("%d", pos.Line)
								if pos.Column != 0 {
//line /usr/local/go/src/go/token/position.go:45
			_go_fuzz_dep_.CoverTab[49108]++
									s += fmt.Sprintf(":%d", pos.Column)
//line /usr/local/go/src/go/token/position.go:46
			// _ = "end of CoverTab[49108]"
		} else {
//line /usr/local/go/src/go/token/position.go:47
			_go_fuzz_dep_.CoverTab[49109]++
//line /usr/local/go/src/go/token/position.go:47
			// _ = "end of CoverTab[49109]"
//line /usr/local/go/src/go/token/position.go:47
		}
//line /usr/local/go/src/go/token/position.go:47
		// _ = "end of CoverTab[49105]"
	} else {
//line /usr/local/go/src/go/token/position.go:48
		_go_fuzz_dep_.CoverTab[49110]++
//line /usr/local/go/src/go/token/position.go:48
		// _ = "end of CoverTab[49110]"
//line /usr/local/go/src/go/token/position.go:48
	}
//line /usr/local/go/src/go/token/position.go:48
	// _ = "end of CoverTab[49101]"
//line /usr/local/go/src/go/token/position.go:48
	_go_fuzz_dep_.CoverTab[49102]++
							if s == "" {
//line /usr/local/go/src/go/token/position.go:49
		_go_fuzz_dep_.CoverTab[49111]++
								s = "-"
//line /usr/local/go/src/go/token/position.go:50
		// _ = "end of CoverTab[49111]"
	} else {
//line /usr/local/go/src/go/token/position.go:51
		_go_fuzz_dep_.CoverTab[49112]++
//line /usr/local/go/src/go/token/position.go:51
		// _ = "end of CoverTab[49112]"
//line /usr/local/go/src/go/token/position.go:51
	}
//line /usr/local/go/src/go/token/position.go:51
	// _ = "end of CoverTab[49102]"
//line /usr/local/go/src/go/token/position.go:51
	_go_fuzz_dep_.CoverTab[49103]++
							return s
//line /usr/local/go/src/go/token/position.go:52
	// _ = "end of CoverTab[49103]"
}

// Pos is a compact encoding of a source position within a file set.
//line /usr/local/go/src/go/token/position.go:55
// It can be converted into a Position for a more convenient, but much
//line /usr/local/go/src/go/token/position.go:55
// larger, representation.
//line /usr/local/go/src/go/token/position.go:55
//
//line /usr/local/go/src/go/token/position.go:55
// The Pos value for a given file is a number in the range [base, base+size],
//line /usr/local/go/src/go/token/position.go:55
// where base and size are specified when a file is added to the file set.
//line /usr/local/go/src/go/token/position.go:55
// The difference between a Pos value and the corresponding file base
//line /usr/local/go/src/go/token/position.go:55
// corresponds to the byte offset of that position (represented by the Pos value)
//line /usr/local/go/src/go/token/position.go:55
// from the beginning of the file. Thus, the file base offset is the Pos value
//line /usr/local/go/src/go/token/position.go:55
// representing the first byte in the file.
//line /usr/local/go/src/go/token/position.go:55
//
//line /usr/local/go/src/go/token/position.go:55
// To create the Pos value for a specific source offset (measured in bytes),
//line /usr/local/go/src/go/token/position.go:55
// first add the respective file to the current file set using FileSet.AddFile
//line /usr/local/go/src/go/token/position.go:55
// and then call File.Pos(offset) for that file. Given a Pos value p
//line /usr/local/go/src/go/token/position.go:55
// for a specific file set fset, the corresponding Position value is
//line /usr/local/go/src/go/token/position.go:55
// obtained by calling fset.Position(p).
//line /usr/local/go/src/go/token/position.go:55
//
//line /usr/local/go/src/go/token/position.go:55
// Pos values can be compared directly with the usual comparison operators:
//line /usr/local/go/src/go/token/position.go:55
// If two Pos values p and q are in the same file, comparing p and q is
//line /usr/local/go/src/go/token/position.go:55
// equivalent to comparing the respective source file offsets. If p and q
//line /usr/local/go/src/go/token/position.go:55
// are in different files, p < q is true if the file implied by p was added
//line /usr/local/go/src/go/token/position.go:55
// to the respective file set before the file implied by q.
//line /usr/local/go/src/go/token/position.go:77
type Pos int

// The zero value for Pos is NoPos; there is no file and line information
//line /usr/local/go/src/go/token/position.go:79
// associated with it, and NoPos.IsValid() is false. NoPos is always
//line /usr/local/go/src/go/token/position.go:79
// smaller than any other Pos value. The corresponding Position value
//line /usr/local/go/src/go/token/position.go:79
// for NoPos is the zero value for Position.
//line /usr/local/go/src/go/token/position.go:83
const NoPos Pos = 0

// IsValid reports whether the position is valid.
func (p Pos) IsValid() bool {
//line /usr/local/go/src/go/token/position.go:86
	_go_fuzz_dep_.CoverTab[49113]++
							return p != NoPos
//line /usr/local/go/src/go/token/position.go:87
	// _ = "end of CoverTab[49113]"
}

//line /usr/local/go/src/go/token/position.go:93
// A File is a handle for a file belonging to a FileSet.
//line /usr/local/go/src/go/token/position.go:93
// A File has a name, size, and line offset table.
//line /usr/local/go/src/go/token/position.go:95
type File struct {
	name	string	// file name as provided to AddFile
	base	int	// Pos value range for this file is [base...base+size]
	size	int	// file size as provided to AddFile

	// lines and infos are protected by mutex
	mutex	sync.Mutex
	lines	[]int	// lines contains the offset of the first character for each line (the first entry is always 0)
	infos	[]lineInfo
}

// Name returns the file name of file f as registered with AddFile.
func (f *File) Name() string {
//line /usr/local/go/src/go/token/position.go:107
	_go_fuzz_dep_.CoverTab[49114]++
							return f.name
//line /usr/local/go/src/go/token/position.go:108
	// _ = "end of CoverTab[49114]"
}

// Base returns the base offset of file f as registered with AddFile.
func (f *File) Base() int {
//line /usr/local/go/src/go/token/position.go:112
	_go_fuzz_dep_.CoverTab[49115]++
							return f.base
//line /usr/local/go/src/go/token/position.go:113
	// _ = "end of CoverTab[49115]"
}

// Size returns the size of file f as registered with AddFile.
func (f *File) Size() int {
//line /usr/local/go/src/go/token/position.go:117
	_go_fuzz_dep_.CoverTab[49116]++
							return f.size
//line /usr/local/go/src/go/token/position.go:118
	// _ = "end of CoverTab[49116]"
}

// LineCount returns the number of lines in file f.
func (f *File) LineCount() int {
//line /usr/local/go/src/go/token/position.go:122
	_go_fuzz_dep_.CoverTab[49117]++
							f.mutex.Lock()
							n := len(f.lines)
							f.mutex.Unlock()
							return n
//line /usr/local/go/src/go/token/position.go:126
	// _ = "end of CoverTab[49117]"
}

// AddLine adds the line offset for a new line.
//line /usr/local/go/src/go/token/position.go:129
// The line offset must be larger than the offset for the previous line
//line /usr/local/go/src/go/token/position.go:129
// and smaller than the file size; otherwise the line offset is ignored.
//line /usr/local/go/src/go/token/position.go:132
func (f *File) AddLine(offset int) {
//line /usr/local/go/src/go/token/position.go:132
	_go_fuzz_dep_.CoverTab[49118]++
							f.mutex.Lock()
							if i := len(f.lines); (i == 0 || func() bool {
//line /usr/local/go/src/go/token/position.go:134
		_go_fuzz_dep_.CoverTab[49120]++
//line /usr/local/go/src/go/token/position.go:134
		return f.lines[i-1] < offset
//line /usr/local/go/src/go/token/position.go:134
		// _ = "end of CoverTab[49120]"
//line /usr/local/go/src/go/token/position.go:134
	}()) && func() bool {
//line /usr/local/go/src/go/token/position.go:134
		_go_fuzz_dep_.CoverTab[49121]++
//line /usr/local/go/src/go/token/position.go:134
		return offset < f.size
//line /usr/local/go/src/go/token/position.go:134
		// _ = "end of CoverTab[49121]"
//line /usr/local/go/src/go/token/position.go:134
	}() {
//line /usr/local/go/src/go/token/position.go:134
		_go_fuzz_dep_.CoverTab[49122]++
								f.lines = append(f.lines, offset)
//line /usr/local/go/src/go/token/position.go:135
		// _ = "end of CoverTab[49122]"
	} else {
//line /usr/local/go/src/go/token/position.go:136
		_go_fuzz_dep_.CoverTab[49123]++
//line /usr/local/go/src/go/token/position.go:136
		// _ = "end of CoverTab[49123]"
//line /usr/local/go/src/go/token/position.go:136
	}
//line /usr/local/go/src/go/token/position.go:136
	// _ = "end of CoverTab[49118]"
//line /usr/local/go/src/go/token/position.go:136
	_go_fuzz_dep_.CoverTab[49119]++
							f.mutex.Unlock()
//line /usr/local/go/src/go/token/position.go:137
	// _ = "end of CoverTab[49119]"
}

// MergeLine merges a line with the following line. It is akin to replacing
//line /usr/local/go/src/go/token/position.go:140
// the newline character at the end of the line with a space (to not change the
//line /usr/local/go/src/go/token/position.go:140
// remaining offsets). To obtain the line number, consult e.g. Position.Line.
//line /usr/local/go/src/go/token/position.go:140
// MergeLine will panic if given an invalid line number.
//line /usr/local/go/src/go/token/position.go:144
func (f *File) MergeLine(line int) {
//line /usr/local/go/src/go/token/position.go:144
	_go_fuzz_dep_.CoverTab[49124]++
							if line < 1 {
//line /usr/local/go/src/go/token/position.go:145
		_go_fuzz_dep_.CoverTab[49127]++
								panic(fmt.Sprintf("invalid line number %d (should be >= 1)", line))
//line /usr/local/go/src/go/token/position.go:146
		// _ = "end of CoverTab[49127]"
	} else {
//line /usr/local/go/src/go/token/position.go:147
		_go_fuzz_dep_.CoverTab[49128]++
//line /usr/local/go/src/go/token/position.go:147
		// _ = "end of CoverTab[49128]"
//line /usr/local/go/src/go/token/position.go:147
	}
//line /usr/local/go/src/go/token/position.go:147
	// _ = "end of CoverTab[49124]"
//line /usr/local/go/src/go/token/position.go:147
	_go_fuzz_dep_.CoverTab[49125]++
							f.mutex.Lock()
							defer f.mutex.Unlock()
							if line >= len(f.lines) {
//line /usr/local/go/src/go/token/position.go:150
		_go_fuzz_dep_.CoverTab[49129]++
								panic(fmt.Sprintf("invalid line number %d (should be < %d)", line, len(f.lines)))
//line /usr/local/go/src/go/token/position.go:151
		// _ = "end of CoverTab[49129]"
	} else {
//line /usr/local/go/src/go/token/position.go:152
		_go_fuzz_dep_.CoverTab[49130]++
//line /usr/local/go/src/go/token/position.go:152
		// _ = "end of CoverTab[49130]"
//line /usr/local/go/src/go/token/position.go:152
	}
//line /usr/local/go/src/go/token/position.go:152
	// _ = "end of CoverTab[49125]"
//line /usr/local/go/src/go/token/position.go:152
	_go_fuzz_dep_.CoverTab[49126]++

//line /usr/local/go/src/go/token/position.go:158
	copy(f.lines[line:], f.lines[line+1:])
							f.lines = f.lines[:len(f.lines)-1]
//line /usr/local/go/src/go/token/position.go:159
	// _ = "end of CoverTab[49126]"
}

// SetLines sets the line offsets for a file and reports whether it succeeded.
//line /usr/local/go/src/go/token/position.go:162
// The line offsets are the offsets of the first character of each line;
//line /usr/local/go/src/go/token/position.go:162
// for instance for the content "ab\nc\n" the line offsets are {0, 3}.
//line /usr/local/go/src/go/token/position.go:162
// An empty file has an empty line offset table.
//line /usr/local/go/src/go/token/position.go:162
// Each line offset must be larger than the offset for the previous line
//line /usr/local/go/src/go/token/position.go:162
// and smaller than the file size; otherwise SetLines fails and returns
//line /usr/local/go/src/go/token/position.go:162
// false.
//line /usr/local/go/src/go/token/position.go:162
// Callers must not mutate the provided slice after SetLines returns.
//line /usr/local/go/src/go/token/position.go:170
func (f *File) SetLines(lines []int) bool {
//line /usr/local/go/src/go/token/position.go:170
	_go_fuzz_dep_.CoverTab[49131]++

							size := f.size
							for i, offset := range lines {
//line /usr/local/go/src/go/token/position.go:173
		_go_fuzz_dep_.CoverTab[49133]++
								if i > 0 && func() bool {
//line /usr/local/go/src/go/token/position.go:174
			_go_fuzz_dep_.CoverTab[49134]++
//line /usr/local/go/src/go/token/position.go:174
			return offset <= lines[i-1]
//line /usr/local/go/src/go/token/position.go:174
			// _ = "end of CoverTab[49134]"
//line /usr/local/go/src/go/token/position.go:174
		}() || func() bool {
//line /usr/local/go/src/go/token/position.go:174
			_go_fuzz_dep_.CoverTab[49135]++
//line /usr/local/go/src/go/token/position.go:174
			return size <= offset
//line /usr/local/go/src/go/token/position.go:174
			// _ = "end of CoverTab[49135]"
//line /usr/local/go/src/go/token/position.go:174
		}() {
//line /usr/local/go/src/go/token/position.go:174
			_go_fuzz_dep_.CoverTab[49136]++
									return false
//line /usr/local/go/src/go/token/position.go:175
			// _ = "end of CoverTab[49136]"
		} else {
//line /usr/local/go/src/go/token/position.go:176
			_go_fuzz_dep_.CoverTab[49137]++
//line /usr/local/go/src/go/token/position.go:176
			// _ = "end of CoverTab[49137]"
//line /usr/local/go/src/go/token/position.go:176
		}
//line /usr/local/go/src/go/token/position.go:176
		// _ = "end of CoverTab[49133]"
	}
//line /usr/local/go/src/go/token/position.go:177
	// _ = "end of CoverTab[49131]"
//line /usr/local/go/src/go/token/position.go:177
	_go_fuzz_dep_.CoverTab[49132]++

//line /usr/local/go/src/go/token/position.go:180
	f.mutex.Lock()
							f.lines = lines
							f.mutex.Unlock()
							return true
//line /usr/local/go/src/go/token/position.go:183
	// _ = "end of CoverTab[49132]"
}

// SetLinesForContent sets the line offsets for the given file content.
//line /usr/local/go/src/go/token/position.go:186
// It ignores position-altering //line comments.
//line /usr/local/go/src/go/token/position.go:188
func (f *File) SetLinesForContent(content []byte) {
//line /usr/local/go/src/go/token/position.go:188
	_go_fuzz_dep_.CoverTab[49138]++
							var lines []int
							line := 0
							for offset, b := range content {
//line /usr/local/go/src/go/token/position.go:191
		_go_fuzz_dep_.CoverTab[49140]++
								if line >= 0 {
//line /usr/local/go/src/go/token/position.go:192
			_go_fuzz_dep_.CoverTab[49142]++
									lines = append(lines, line)
//line /usr/local/go/src/go/token/position.go:193
			// _ = "end of CoverTab[49142]"
		} else {
//line /usr/local/go/src/go/token/position.go:194
			_go_fuzz_dep_.CoverTab[49143]++
//line /usr/local/go/src/go/token/position.go:194
			// _ = "end of CoverTab[49143]"
//line /usr/local/go/src/go/token/position.go:194
		}
//line /usr/local/go/src/go/token/position.go:194
		// _ = "end of CoverTab[49140]"
//line /usr/local/go/src/go/token/position.go:194
		_go_fuzz_dep_.CoverTab[49141]++
								line = -1
								if b == '\n' {
//line /usr/local/go/src/go/token/position.go:196
			_go_fuzz_dep_.CoverTab[49144]++
									line = offset + 1
//line /usr/local/go/src/go/token/position.go:197
			// _ = "end of CoverTab[49144]"
		} else {
//line /usr/local/go/src/go/token/position.go:198
			_go_fuzz_dep_.CoverTab[49145]++
//line /usr/local/go/src/go/token/position.go:198
			// _ = "end of CoverTab[49145]"
//line /usr/local/go/src/go/token/position.go:198
		}
//line /usr/local/go/src/go/token/position.go:198
		// _ = "end of CoverTab[49141]"
	}
//line /usr/local/go/src/go/token/position.go:199
	// _ = "end of CoverTab[49138]"
//line /usr/local/go/src/go/token/position.go:199
	_go_fuzz_dep_.CoverTab[49139]++

//line /usr/local/go/src/go/token/position.go:202
	f.mutex.Lock()
							f.lines = lines
							f.mutex.Unlock()
//line /usr/local/go/src/go/token/position.go:204
	// _ = "end of CoverTab[49139]"
}

// LineStart returns the Pos value of the start of the specified line.
//line /usr/local/go/src/go/token/position.go:207
// It ignores any alternative positions set using AddLineColumnInfo.
//line /usr/local/go/src/go/token/position.go:207
// LineStart panics if the 1-based line number is invalid.
//line /usr/local/go/src/go/token/position.go:210
func (f *File) LineStart(line int) Pos {
//line /usr/local/go/src/go/token/position.go:210
	_go_fuzz_dep_.CoverTab[49146]++
							if line < 1 {
//line /usr/local/go/src/go/token/position.go:211
		_go_fuzz_dep_.CoverTab[49149]++
								panic(fmt.Sprintf("invalid line number %d (should be >= 1)", line))
//line /usr/local/go/src/go/token/position.go:212
		// _ = "end of CoverTab[49149]"
	} else {
//line /usr/local/go/src/go/token/position.go:213
		_go_fuzz_dep_.CoverTab[49150]++
//line /usr/local/go/src/go/token/position.go:213
		// _ = "end of CoverTab[49150]"
//line /usr/local/go/src/go/token/position.go:213
	}
//line /usr/local/go/src/go/token/position.go:213
	// _ = "end of CoverTab[49146]"
//line /usr/local/go/src/go/token/position.go:213
	_go_fuzz_dep_.CoverTab[49147]++
							f.mutex.Lock()
							defer f.mutex.Unlock()
							if line > len(f.lines) {
//line /usr/local/go/src/go/token/position.go:216
		_go_fuzz_dep_.CoverTab[49151]++
								panic(fmt.Sprintf("invalid line number %d (should be < %d)", line, len(f.lines)))
//line /usr/local/go/src/go/token/position.go:217
		// _ = "end of CoverTab[49151]"
	} else {
//line /usr/local/go/src/go/token/position.go:218
		_go_fuzz_dep_.CoverTab[49152]++
//line /usr/local/go/src/go/token/position.go:218
		// _ = "end of CoverTab[49152]"
//line /usr/local/go/src/go/token/position.go:218
	}
//line /usr/local/go/src/go/token/position.go:218
	// _ = "end of CoverTab[49147]"
//line /usr/local/go/src/go/token/position.go:218
	_go_fuzz_dep_.CoverTab[49148]++
							return Pos(f.base + f.lines[line-1])
//line /usr/local/go/src/go/token/position.go:219
	// _ = "end of CoverTab[49148]"
}

// A lineInfo object describes alternative file, line, and column
//line /usr/local/go/src/go/token/position.go:222
// number information (such as provided via a //line directive)
//line /usr/local/go/src/go/token/position.go:222
// for a given file offset.
//line /usr/local/go/src/go/token/position.go:225
type lineInfo struct {
	// fields are exported to make them accessible to gob
	Offset		int
	Filename	string
	Line, Column	int
}

// AddLineInfo is like AddLineColumnInfo with a column = 1 argument.
//line /usr/local/go/src/go/token/position.go:232
// It is here for backward-compatibility for code prior to Go 1.11.
//line /usr/local/go/src/go/token/position.go:234
func (f *File) AddLineInfo(offset int, filename string, line int) {
//line /usr/local/go/src/go/token/position.go:234
	_go_fuzz_dep_.CoverTab[49153]++
							f.AddLineColumnInfo(offset, filename, line, 1)
//line /usr/local/go/src/go/token/position.go:235
	// _ = "end of CoverTab[49153]"
}

// AddLineColumnInfo adds alternative file, line, and column number
//line /usr/local/go/src/go/token/position.go:238
// information for a given file offset. The offset must be larger
//line /usr/local/go/src/go/token/position.go:238
// than the offset for the previously added alternative line info
//line /usr/local/go/src/go/token/position.go:238
// and smaller than the file size; otherwise the information is
//line /usr/local/go/src/go/token/position.go:238
// ignored.
//line /usr/local/go/src/go/token/position.go:238
//
//line /usr/local/go/src/go/token/position.go:238
// AddLineColumnInfo is typically used to register alternative position
//line /usr/local/go/src/go/token/position.go:238
// information for line directives such as //line filename:line:column.
//line /usr/local/go/src/go/token/position.go:246
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int) {
//line /usr/local/go/src/go/token/position.go:246
	_go_fuzz_dep_.CoverTab[49154]++
							f.mutex.Lock()
							if i := len(f.infos); (i == 0 || func() bool {
//line /usr/local/go/src/go/token/position.go:248
		_go_fuzz_dep_.CoverTab[49156]++
//line /usr/local/go/src/go/token/position.go:248
		return f.infos[i-1].Offset < offset
//line /usr/local/go/src/go/token/position.go:248
		// _ = "end of CoverTab[49156]"
//line /usr/local/go/src/go/token/position.go:248
	}()) && func() bool {
//line /usr/local/go/src/go/token/position.go:248
		_go_fuzz_dep_.CoverTab[49157]++
//line /usr/local/go/src/go/token/position.go:248
		return offset < f.size
//line /usr/local/go/src/go/token/position.go:248
		// _ = "end of CoverTab[49157]"
//line /usr/local/go/src/go/token/position.go:248
	}() {
//line /usr/local/go/src/go/token/position.go:248
		_go_fuzz_dep_.CoverTab[49158]++
								f.infos = append(f.infos, lineInfo{offset, filename, line, column})
//line /usr/local/go/src/go/token/position.go:249
		// _ = "end of CoverTab[49158]"
	} else {
//line /usr/local/go/src/go/token/position.go:250
		_go_fuzz_dep_.CoverTab[49159]++
//line /usr/local/go/src/go/token/position.go:250
		// _ = "end of CoverTab[49159]"
//line /usr/local/go/src/go/token/position.go:250
	}
//line /usr/local/go/src/go/token/position.go:250
	// _ = "end of CoverTab[49154]"
//line /usr/local/go/src/go/token/position.go:250
	_go_fuzz_dep_.CoverTab[49155]++
							f.mutex.Unlock()
//line /usr/local/go/src/go/token/position.go:251
	// _ = "end of CoverTab[49155]"
}

// Pos returns the Pos value for the given file offset;
//line /usr/local/go/src/go/token/position.go:254
// the offset must be <= f.Size().
//line /usr/local/go/src/go/token/position.go:254
// f.Pos(f.Offset(p)) == p.
//line /usr/local/go/src/go/token/position.go:257
func (f *File) Pos(offset int) Pos {
//line /usr/local/go/src/go/token/position.go:257
	_go_fuzz_dep_.CoverTab[49160]++
							if offset > f.size {
//line /usr/local/go/src/go/token/position.go:258
		_go_fuzz_dep_.CoverTab[49162]++
								panic(fmt.Sprintf("invalid file offset %d (should be <= %d)", offset, f.size))
//line /usr/local/go/src/go/token/position.go:259
		// _ = "end of CoverTab[49162]"
	} else {
//line /usr/local/go/src/go/token/position.go:260
		_go_fuzz_dep_.CoverTab[49163]++
//line /usr/local/go/src/go/token/position.go:260
		// _ = "end of CoverTab[49163]"
//line /usr/local/go/src/go/token/position.go:260
	}
//line /usr/local/go/src/go/token/position.go:260
	// _ = "end of CoverTab[49160]"
//line /usr/local/go/src/go/token/position.go:260
	_go_fuzz_dep_.CoverTab[49161]++
							return Pos(f.base + offset)
//line /usr/local/go/src/go/token/position.go:261
	// _ = "end of CoverTab[49161]"
}

// Offset returns the offset for the given file position p;
//line /usr/local/go/src/go/token/position.go:264
// p must be a valid Pos value in that file.
//line /usr/local/go/src/go/token/position.go:264
// f.Offset(f.Pos(offset)) == offset.
//line /usr/local/go/src/go/token/position.go:267
func (f *File) Offset(p Pos) int {
//line /usr/local/go/src/go/token/position.go:267
	_go_fuzz_dep_.CoverTab[49164]++
							if int(p) < f.base || func() bool {
//line /usr/local/go/src/go/token/position.go:268
		_go_fuzz_dep_.CoverTab[49166]++
//line /usr/local/go/src/go/token/position.go:268
		return int(p) > f.base+f.size
//line /usr/local/go/src/go/token/position.go:268
		// _ = "end of CoverTab[49166]"
//line /usr/local/go/src/go/token/position.go:268
	}() {
//line /usr/local/go/src/go/token/position.go:268
		_go_fuzz_dep_.CoverTab[49167]++
								panic(fmt.Sprintf("invalid Pos value %d (should be in [%d, %d])", p, f.base, f.base+f.size))
//line /usr/local/go/src/go/token/position.go:269
		// _ = "end of CoverTab[49167]"
	} else {
//line /usr/local/go/src/go/token/position.go:270
		_go_fuzz_dep_.CoverTab[49168]++
//line /usr/local/go/src/go/token/position.go:270
		// _ = "end of CoverTab[49168]"
//line /usr/local/go/src/go/token/position.go:270
	}
//line /usr/local/go/src/go/token/position.go:270
	// _ = "end of CoverTab[49164]"
//line /usr/local/go/src/go/token/position.go:270
	_go_fuzz_dep_.CoverTab[49165]++
							return int(p) - f.base
//line /usr/local/go/src/go/token/position.go:271
	// _ = "end of CoverTab[49165]"
}

// Line returns the line number for the given file position p;
//line /usr/local/go/src/go/token/position.go:274
// p must be a Pos value in that file or NoPos.
//line /usr/local/go/src/go/token/position.go:276
func (f *File) Line(p Pos) int {
//line /usr/local/go/src/go/token/position.go:276
	_go_fuzz_dep_.CoverTab[49169]++
							return f.Position(p).Line
//line /usr/local/go/src/go/token/position.go:277
	// _ = "end of CoverTab[49169]"
}

func searchLineInfos(a []lineInfo, x int) int {
//line /usr/local/go/src/go/token/position.go:280
	_go_fuzz_dep_.CoverTab[49170]++
							return sort.Search(len(a), func(i int) bool {
//line /usr/local/go/src/go/token/position.go:281
		_go_fuzz_dep_.CoverTab[49171]++
//line /usr/local/go/src/go/token/position.go:281
		return a[i].Offset > x
//line /usr/local/go/src/go/token/position.go:281
		// _ = "end of CoverTab[49171]"
//line /usr/local/go/src/go/token/position.go:281
	}) - 1
//line /usr/local/go/src/go/token/position.go:281
	// _ = "end of CoverTab[49170]"
}

// unpack returns the filename and line and column number for a file offset.
//line /usr/local/go/src/go/token/position.go:284
// If adjusted is set, unpack will return the filename and line information
//line /usr/local/go/src/go/token/position.go:284
// possibly adjusted by //line comments; otherwise those comments are ignored.
//line /usr/local/go/src/go/token/position.go:287
func (f *File) unpack(offset int, adjusted bool) (filename string, line, column int) {
//line /usr/local/go/src/go/token/position.go:287
	_go_fuzz_dep_.CoverTab[49172]++
							f.mutex.Lock()
							filename = f.name
							if i := searchInts(f.lines, offset); i >= 0 {
//line /usr/local/go/src/go/token/position.go:290
		_go_fuzz_dep_.CoverTab[49175]++
								line, column = i+1, offset-f.lines[i]+1
//line /usr/local/go/src/go/token/position.go:291
		// _ = "end of CoverTab[49175]"
	} else {
//line /usr/local/go/src/go/token/position.go:292
		_go_fuzz_dep_.CoverTab[49176]++
//line /usr/local/go/src/go/token/position.go:292
		// _ = "end of CoverTab[49176]"
//line /usr/local/go/src/go/token/position.go:292
	}
//line /usr/local/go/src/go/token/position.go:292
	// _ = "end of CoverTab[49172]"
//line /usr/local/go/src/go/token/position.go:292
	_go_fuzz_dep_.CoverTab[49173]++
							if adjusted && func() bool {
//line /usr/local/go/src/go/token/position.go:293
		_go_fuzz_dep_.CoverTab[49177]++
//line /usr/local/go/src/go/token/position.go:293
		return len(f.infos) > 0
//line /usr/local/go/src/go/token/position.go:293
		// _ = "end of CoverTab[49177]"
//line /usr/local/go/src/go/token/position.go:293
	}() {
//line /usr/local/go/src/go/token/position.go:293
		_go_fuzz_dep_.CoverTab[49178]++

								if i := searchLineInfos(f.infos, offset); i >= 0 {
//line /usr/local/go/src/go/token/position.go:295
			_go_fuzz_dep_.CoverTab[49179]++
									alt := &f.infos[i]
									filename = alt.Filename
									if i := searchInts(f.lines, alt.Offset); i >= 0 {
//line /usr/local/go/src/go/token/position.go:298
				_go_fuzz_dep_.CoverTab[49180]++

										d := line - (i + 1)
										line = alt.Line + d
										if alt.Column == 0 {
//line /usr/local/go/src/go/token/position.go:302
					_go_fuzz_dep_.CoverTab[49181]++

//line /usr/local/go/src/go/token/position.go:307
					column = 0
//line /usr/local/go/src/go/token/position.go:307
					// _ = "end of CoverTab[49181]"
				} else {
//line /usr/local/go/src/go/token/position.go:308
					_go_fuzz_dep_.CoverTab[49182]++
//line /usr/local/go/src/go/token/position.go:308
					if d == 0 {
//line /usr/local/go/src/go/token/position.go:308
						_go_fuzz_dep_.CoverTab[49183]++

//line /usr/local/go/src/go/token/position.go:311
						column = alt.Column + (offset - alt.Offset)
//line /usr/local/go/src/go/token/position.go:311
						// _ = "end of CoverTab[49183]"
					} else {
//line /usr/local/go/src/go/token/position.go:312
						_go_fuzz_dep_.CoverTab[49184]++
//line /usr/local/go/src/go/token/position.go:312
						// _ = "end of CoverTab[49184]"
//line /usr/local/go/src/go/token/position.go:312
					}
//line /usr/local/go/src/go/token/position.go:312
					// _ = "end of CoverTab[49182]"
//line /usr/local/go/src/go/token/position.go:312
				}
//line /usr/local/go/src/go/token/position.go:312
				// _ = "end of CoverTab[49180]"
			} else {
//line /usr/local/go/src/go/token/position.go:313
				_go_fuzz_dep_.CoverTab[49185]++
//line /usr/local/go/src/go/token/position.go:313
				// _ = "end of CoverTab[49185]"
//line /usr/local/go/src/go/token/position.go:313
			}
//line /usr/local/go/src/go/token/position.go:313
			// _ = "end of CoverTab[49179]"
		} else {
//line /usr/local/go/src/go/token/position.go:314
			_go_fuzz_dep_.CoverTab[49186]++
//line /usr/local/go/src/go/token/position.go:314
			// _ = "end of CoverTab[49186]"
//line /usr/local/go/src/go/token/position.go:314
		}
//line /usr/local/go/src/go/token/position.go:314
		// _ = "end of CoverTab[49178]"
	} else {
//line /usr/local/go/src/go/token/position.go:315
		_go_fuzz_dep_.CoverTab[49187]++
//line /usr/local/go/src/go/token/position.go:315
		// _ = "end of CoverTab[49187]"
//line /usr/local/go/src/go/token/position.go:315
	}
//line /usr/local/go/src/go/token/position.go:315
	// _ = "end of CoverTab[49173]"
//line /usr/local/go/src/go/token/position.go:315
	_go_fuzz_dep_.CoverTab[49174]++

//line /usr/local/go/src/go/token/position.go:318
	f.mutex.Unlock()
							return
//line /usr/local/go/src/go/token/position.go:319
	// _ = "end of CoverTab[49174]"
}

func (f *File) position(p Pos, adjusted bool) (pos Position) {
//line /usr/local/go/src/go/token/position.go:322
	_go_fuzz_dep_.CoverTab[49188]++
							offset := int(p) - f.base
							pos.Offset = offset
							pos.Filename, pos.Line, pos.Column = f.unpack(offset, adjusted)
							return
//line /usr/local/go/src/go/token/position.go:326
	// _ = "end of CoverTab[49188]"
}

// PositionFor returns the Position value for the given file position p.
//line /usr/local/go/src/go/token/position.go:329
// If adjusted is set, the position may be adjusted by position-altering
//line /usr/local/go/src/go/token/position.go:329
// //line comments; otherwise those comments are ignored.
//line /usr/local/go/src/go/token/position.go:329
// p must be a Pos value in f or NoPos.
//line /usr/local/go/src/go/token/position.go:333
func (f *File) PositionFor(p Pos, adjusted bool) (pos Position) {
//line /usr/local/go/src/go/token/position.go:333
	_go_fuzz_dep_.CoverTab[49189]++
							if p != NoPos {
//line /usr/local/go/src/go/token/position.go:334
		_go_fuzz_dep_.CoverTab[49191]++
								if int(p) < f.base || func() bool {
//line /usr/local/go/src/go/token/position.go:335
			_go_fuzz_dep_.CoverTab[49193]++
//line /usr/local/go/src/go/token/position.go:335
			return int(p) > f.base+f.size
//line /usr/local/go/src/go/token/position.go:335
			// _ = "end of CoverTab[49193]"
//line /usr/local/go/src/go/token/position.go:335
		}() {
//line /usr/local/go/src/go/token/position.go:335
			_go_fuzz_dep_.CoverTab[49194]++
									panic(fmt.Sprintf("invalid Pos value %d (should be in [%d, %d])", p, f.base, f.base+f.size))
//line /usr/local/go/src/go/token/position.go:336
			// _ = "end of CoverTab[49194]"
		} else {
//line /usr/local/go/src/go/token/position.go:337
			_go_fuzz_dep_.CoverTab[49195]++
//line /usr/local/go/src/go/token/position.go:337
			// _ = "end of CoverTab[49195]"
//line /usr/local/go/src/go/token/position.go:337
		}
//line /usr/local/go/src/go/token/position.go:337
		// _ = "end of CoverTab[49191]"
//line /usr/local/go/src/go/token/position.go:337
		_go_fuzz_dep_.CoverTab[49192]++
								pos = f.position(p, adjusted)
//line /usr/local/go/src/go/token/position.go:338
		// _ = "end of CoverTab[49192]"
	} else {
//line /usr/local/go/src/go/token/position.go:339
		_go_fuzz_dep_.CoverTab[49196]++
//line /usr/local/go/src/go/token/position.go:339
		// _ = "end of CoverTab[49196]"
//line /usr/local/go/src/go/token/position.go:339
	}
//line /usr/local/go/src/go/token/position.go:339
	// _ = "end of CoverTab[49189]"
//line /usr/local/go/src/go/token/position.go:339
	_go_fuzz_dep_.CoverTab[49190]++
							return
//line /usr/local/go/src/go/token/position.go:340
	// _ = "end of CoverTab[49190]"
}

// Position returns the Position value for the given file position p.
//line /usr/local/go/src/go/token/position.go:343
// Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).
//line /usr/local/go/src/go/token/position.go:345
func (f *File) Position(p Pos) (pos Position) {
//line /usr/local/go/src/go/token/position.go:345
	_go_fuzz_dep_.CoverTab[49197]++
							return f.PositionFor(p, true)
//line /usr/local/go/src/go/token/position.go:346
	// _ = "end of CoverTab[49197]"
}

//line /usr/local/go/src/go/token/position.go:352
// A FileSet represents a set of source files.
//line /usr/local/go/src/go/token/position.go:352
// Methods of file sets are synchronized; multiple goroutines
//line /usr/local/go/src/go/token/position.go:352
// may invoke them concurrently.
//line /usr/local/go/src/go/token/position.go:352
//
//line /usr/local/go/src/go/token/position.go:352
// The byte offsets for each file in a file set are mapped into
//line /usr/local/go/src/go/token/position.go:352
// distinct (integer) intervals, one interval [base, base+size]
//line /usr/local/go/src/go/token/position.go:352
// per file. Base represents the first byte in the file, and size
//line /usr/local/go/src/go/token/position.go:352
// is the corresponding file size. A Pos value is a value in such
//line /usr/local/go/src/go/token/position.go:352
// an interval. By determining the interval a Pos value belongs
//line /usr/local/go/src/go/token/position.go:352
// to, the file, its file base, and thus the byte offset (position)
//line /usr/local/go/src/go/token/position.go:352
// the Pos value is representing can be computed.
//line /usr/local/go/src/go/token/position.go:352
//
//line /usr/local/go/src/go/token/position.go:352
// When adding a new file, a file base must be provided. That can
//line /usr/local/go/src/go/token/position.go:352
// be any integer value that is past the end of any interval of any
//line /usr/local/go/src/go/token/position.go:352
// file already in the file set. For convenience, FileSet.Base provides
//line /usr/local/go/src/go/token/position.go:352
// such a value, which is simply the end of the Pos interval of the most
//line /usr/local/go/src/go/token/position.go:352
// recently added file, plus one. Unless there is a need to extend an
//line /usr/local/go/src/go/token/position.go:352
// interval later, using the FileSet.Base should be used as argument
//line /usr/local/go/src/go/token/position.go:352
// for FileSet.AddFile.
//line /usr/local/go/src/go/token/position.go:352
//
//line /usr/local/go/src/go/token/position.go:352
// A File may be removed from a FileSet when it is no longer needed.
//line /usr/local/go/src/go/token/position.go:352
// This may reduce memory usage in a long-running application.
//line /usr/local/go/src/go/token/position.go:374
type FileSet struct {
	mutex	sync.RWMutex		// protects the file set
	base	int			// base offset for the next file
	files	[]*File			// list of files in the order added to the set
	last	atomic.Pointer[File]	// cache of last file looked up
}

// NewFileSet creates a new file set.
func NewFileSet() *FileSet {
//line /usr/local/go/src/go/token/position.go:382
	_go_fuzz_dep_.CoverTab[49198]++
							return &FileSet{
		base: 1,
	}
//line /usr/local/go/src/go/token/position.go:385
	// _ = "end of CoverTab[49198]"
}

// Base returns the minimum base offset that must be provided to
//line /usr/local/go/src/go/token/position.go:388
// AddFile when adding the next file.
//line /usr/local/go/src/go/token/position.go:390
func (s *FileSet) Base() int {
//line /usr/local/go/src/go/token/position.go:390
	_go_fuzz_dep_.CoverTab[49199]++
							s.mutex.RLock()
							b := s.base
							s.mutex.RUnlock()
							return b
//line /usr/local/go/src/go/token/position.go:394
	// _ = "end of CoverTab[49199]"

}

// AddFile adds a new file with a given filename, base offset, and file size
//line /usr/local/go/src/go/token/position.go:398
// to the file set s and returns the file. Multiple files may have the same
//line /usr/local/go/src/go/token/position.go:398
// name. The base offset must not be smaller than the FileSet's Base(), and
//line /usr/local/go/src/go/token/position.go:398
// size must not be negative. As a special case, if a negative base is provided,
//line /usr/local/go/src/go/token/position.go:398
// the current value of the FileSet's Base() is used instead.
//line /usr/local/go/src/go/token/position.go:398
//
//line /usr/local/go/src/go/token/position.go:398
// Adding the file will set the file set's Base() value to base + size + 1
//line /usr/local/go/src/go/token/position.go:398
// as the minimum base value for the next file. The following relationship
//line /usr/local/go/src/go/token/position.go:398
// exists between a Pos value p for a given file offset offs:
//line /usr/local/go/src/go/token/position.go:398
//
//line /usr/local/go/src/go/token/position.go:398
//	int(p) = base + offs
//line /usr/local/go/src/go/token/position.go:398
//
//line /usr/local/go/src/go/token/position.go:398
// with offs in the range [0, size] and thus p in the range [base, base+size].
//line /usr/local/go/src/go/token/position.go:398
// For convenience, File.Pos may be used to create file-specific position
//line /usr/local/go/src/go/token/position.go:398
// values from a file offset.
//line /usr/local/go/src/go/token/position.go:413
func (s *FileSet) AddFile(filename string, base, size int) *File {
//line /usr/local/go/src/go/token/position.go:413
	_go_fuzz_dep_.CoverTab[49200]++

							f := &File{name: filename, size: size, lines: []int{0}}

							s.mutex.Lock()
							defer s.mutex.Unlock()
							if base < 0 {
//line /usr/local/go/src/go/token/position.go:419
		_go_fuzz_dep_.CoverTab[49205]++
								base = s.base
//line /usr/local/go/src/go/token/position.go:420
		// _ = "end of CoverTab[49205]"
	} else {
//line /usr/local/go/src/go/token/position.go:421
		_go_fuzz_dep_.CoverTab[49206]++
//line /usr/local/go/src/go/token/position.go:421
		// _ = "end of CoverTab[49206]"
//line /usr/local/go/src/go/token/position.go:421
	}
//line /usr/local/go/src/go/token/position.go:421
	// _ = "end of CoverTab[49200]"
//line /usr/local/go/src/go/token/position.go:421
	_go_fuzz_dep_.CoverTab[49201]++
							if base < s.base {
//line /usr/local/go/src/go/token/position.go:422
		_go_fuzz_dep_.CoverTab[49207]++
								panic(fmt.Sprintf("invalid base %d (should be >= %d)", base, s.base))
//line /usr/local/go/src/go/token/position.go:423
		// _ = "end of CoverTab[49207]"
	} else {
//line /usr/local/go/src/go/token/position.go:424
		_go_fuzz_dep_.CoverTab[49208]++
//line /usr/local/go/src/go/token/position.go:424
		// _ = "end of CoverTab[49208]"
//line /usr/local/go/src/go/token/position.go:424
	}
//line /usr/local/go/src/go/token/position.go:424
	// _ = "end of CoverTab[49201]"
//line /usr/local/go/src/go/token/position.go:424
	_go_fuzz_dep_.CoverTab[49202]++
							f.base = base
							if size < 0 {
//line /usr/local/go/src/go/token/position.go:426
		_go_fuzz_dep_.CoverTab[49209]++
								panic(fmt.Sprintf("invalid size %d (should be >= 0)", size))
//line /usr/local/go/src/go/token/position.go:427
		// _ = "end of CoverTab[49209]"
	} else {
//line /usr/local/go/src/go/token/position.go:428
		_go_fuzz_dep_.CoverTab[49210]++
//line /usr/local/go/src/go/token/position.go:428
		// _ = "end of CoverTab[49210]"
//line /usr/local/go/src/go/token/position.go:428
	}
//line /usr/local/go/src/go/token/position.go:428
	// _ = "end of CoverTab[49202]"
//line /usr/local/go/src/go/token/position.go:428
	_go_fuzz_dep_.CoverTab[49203]++

							base += size + 1
							if base < 0 {
//line /usr/local/go/src/go/token/position.go:431
		_go_fuzz_dep_.CoverTab[49211]++
								panic("token.Pos offset overflow (> 2G of source code in file set)")
//line /usr/local/go/src/go/token/position.go:432
		// _ = "end of CoverTab[49211]"
	} else {
//line /usr/local/go/src/go/token/position.go:433
		_go_fuzz_dep_.CoverTab[49212]++
//line /usr/local/go/src/go/token/position.go:433
		// _ = "end of CoverTab[49212]"
//line /usr/local/go/src/go/token/position.go:433
	}
//line /usr/local/go/src/go/token/position.go:433
	// _ = "end of CoverTab[49203]"
//line /usr/local/go/src/go/token/position.go:433
	_go_fuzz_dep_.CoverTab[49204]++

							s.base = base
							s.files = append(s.files, f)
							s.last.Store(f)
							return f
//line /usr/local/go/src/go/token/position.go:438
	// _ = "end of CoverTab[49204]"
}

// RemoveFile removes a file from the FileSet so that subsequent
//line /usr/local/go/src/go/token/position.go:441
// queries for its Pos interval yield a negative result.
//line /usr/local/go/src/go/token/position.go:441
// This reduces the memory usage of a long-lived FileSet that
//line /usr/local/go/src/go/token/position.go:441
// encounters an unbounded stream of files.
//line /usr/local/go/src/go/token/position.go:441
//
//line /usr/local/go/src/go/token/position.go:441
// Removing a file that does not belong to the set has no effect.
//line /usr/local/go/src/go/token/position.go:447
func (s *FileSet) RemoveFile(file *File) {
//line /usr/local/go/src/go/token/position.go:447
	_go_fuzz_dep_.CoverTab[49213]++
							s.last.CompareAndSwap(file, nil)

							s.mutex.Lock()
							defer s.mutex.Unlock()

							if i := searchFiles(s.files, file.base); i >= 0 && func() bool {
//line /usr/local/go/src/go/token/position.go:453
		_go_fuzz_dep_.CoverTab[49214]++
//line /usr/local/go/src/go/token/position.go:453
		return s.files[i] == file
//line /usr/local/go/src/go/token/position.go:453
		// _ = "end of CoverTab[49214]"
//line /usr/local/go/src/go/token/position.go:453
	}() {
//line /usr/local/go/src/go/token/position.go:453
		_go_fuzz_dep_.CoverTab[49215]++
								last := &s.files[len(s.files)-1]
								s.files = append(s.files[:i], s.files[i+1:]...)
								*last = nil
//line /usr/local/go/src/go/token/position.go:456
		// _ = "end of CoverTab[49215]"
	} else {
//line /usr/local/go/src/go/token/position.go:457
		_go_fuzz_dep_.CoverTab[49216]++
//line /usr/local/go/src/go/token/position.go:457
		// _ = "end of CoverTab[49216]"
//line /usr/local/go/src/go/token/position.go:457
	}
//line /usr/local/go/src/go/token/position.go:457
	// _ = "end of CoverTab[49213]"
}

// Iterate calls f for the files in the file set in the order they were added
//line /usr/local/go/src/go/token/position.go:460
// until f returns false.
//line /usr/local/go/src/go/token/position.go:462
func (s *FileSet) Iterate(f func(*File) bool) {
//line /usr/local/go/src/go/token/position.go:462
	_go_fuzz_dep_.CoverTab[49217]++
							for i := 0; ; i++ {
//line /usr/local/go/src/go/token/position.go:463
		_go_fuzz_dep_.CoverTab[49218]++
								var file *File
								s.mutex.RLock()
								if i < len(s.files) {
//line /usr/local/go/src/go/token/position.go:466
			_go_fuzz_dep_.CoverTab[49220]++
									file = s.files[i]
//line /usr/local/go/src/go/token/position.go:467
			// _ = "end of CoverTab[49220]"
		} else {
//line /usr/local/go/src/go/token/position.go:468
			_go_fuzz_dep_.CoverTab[49221]++
//line /usr/local/go/src/go/token/position.go:468
			// _ = "end of CoverTab[49221]"
//line /usr/local/go/src/go/token/position.go:468
		}
//line /usr/local/go/src/go/token/position.go:468
		// _ = "end of CoverTab[49218]"
//line /usr/local/go/src/go/token/position.go:468
		_go_fuzz_dep_.CoverTab[49219]++
								s.mutex.RUnlock()
								if file == nil || func() bool {
//line /usr/local/go/src/go/token/position.go:470
			_go_fuzz_dep_.CoverTab[49222]++
//line /usr/local/go/src/go/token/position.go:470
			return !f(file)
//line /usr/local/go/src/go/token/position.go:470
			// _ = "end of CoverTab[49222]"
//line /usr/local/go/src/go/token/position.go:470
		}() {
//line /usr/local/go/src/go/token/position.go:470
			_go_fuzz_dep_.CoverTab[49223]++
									break
//line /usr/local/go/src/go/token/position.go:471
			// _ = "end of CoverTab[49223]"
		} else {
//line /usr/local/go/src/go/token/position.go:472
			_go_fuzz_dep_.CoverTab[49224]++
//line /usr/local/go/src/go/token/position.go:472
			// _ = "end of CoverTab[49224]"
//line /usr/local/go/src/go/token/position.go:472
		}
//line /usr/local/go/src/go/token/position.go:472
		// _ = "end of CoverTab[49219]"
	}
//line /usr/local/go/src/go/token/position.go:473
	// _ = "end of CoverTab[49217]"
}

func searchFiles(a []*File, x int) int {
//line /usr/local/go/src/go/token/position.go:476
	_go_fuzz_dep_.CoverTab[49225]++
							return sort.Search(len(a), func(i int) bool {
//line /usr/local/go/src/go/token/position.go:477
		_go_fuzz_dep_.CoverTab[49226]++
//line /usr/local/go/src/go/token/position.go:477
		return a[i].base > x
//line /usr/local/go/src/go/token/position.go:477
		// _ = "end of CoverTab[49226]"
//line /usr/local/go/src/go/token/position.go:477
	}) - 1
//line /usr/local/go/src/go/token/position.go:477
	// _ = "end of CoverTab[49225]"
}

func (s *FileSet) file(p Pos) *File {
//line /usr/local/go/src/go/token/position.go:480
	_go_fuzz_dep_.CoverTab[49227]++

							if f := s.last.Load(); f != nil && func() bool {
//line /usr/local/go/src/go/token/position.go:482
		_go_fuzz_dep_.CoverTab[49230]++
//line /usr/local/go/src/go/token/position.go:482
		return f.base <= int(p)
//line /usr/local/go/src/go/token/position.go:482
		// _ = "end of CoverTab[49230]"
//line /usr/local/go/src/go/token/position.go:482
	}() && func() bool {
//line /usr/local/go/src/go/token/position.go:482
		_go_fuzz_dep_.CoverTab[49231]++
//line /usr/local/go/src/go/token/position.go:482
		return int(p) <= f.base+f.size
//line /usr/local/go/src/go/token/position.go:482
		// _ = "end of CoverTab[49231]"
//line /usr/local/go/src/go/token/position.go:482
	}() {
//line /usr/local/go/src/go/token/position.go:482
		_go_fuzz_dep_.CoverTab[49232]++
								return f
//line /usr/local/go/src/go/token/position.go:483
		// _ = "end of CoverTab[49232]"
	} else {
//line /usr/local/go/src/go/token/position.go:484
		_go_fuzz_dep_.CoverTab[49233]++
//line /usr/local/go/src/go/token/position.go:484
		// _ = "end of CoverTab[49233]"
//line /usr/local/go/src/go/token/position.go:484
	}
//line /usr/local/go/src/go/token/position.go:484
	// _ = "end of CoverTab[49227]"
//line /usr/local/go/src/go/token/position.go:484
	_go_fuzz_dep_.CoverTab[49228]++

							s.mutex.RLock()
							defer s.mutex.RUnlock()

//line /usr/local/go/src/go/token/position.go:490
	if i := searchFiles(s.files, int(p)); i >= 0 {
//line /usr/local/go/src/go/token/position.go:490
		_go_fuzz_dep_.CoverTab[49234]++
								f := s.files[i]

								if int(p) <= f.base+f.size {
//line /usr/local/go/src/go/token/position.go:493
			_go_fuzz_dep_.CoverTab[49235]++

//line /usr/local/go/src/go/token/position.go:496
			s.last.Store(f)
									return f
//line /usr/local/go/src/go/token/position.go:497
			// _ = "end of CoverTab[49235]"
		} else {
//line /usr/local/go/src/go/token/position.go:498
			_go_fuzz_dep_.CoverTab[49236]++
//line /usr/local/go/src/go/token/position.go:498
			// _ = "end of CoverTab[49236]"
//line /usr/local/go/src/go/token/position.go:498
		}
//line /usr/local/go/src/go/token/position.go:498
		// _ = "end of CoverTab[49234]"
	} else {
//line /usr/local/go/src/go/token/position.go:499
		_go_fuzz_dep_.CoverTab[49237]++
//line /usr/local/go/src/go/token/position.go:499
		// _ = "end of CoverTab[49237]"
//line /usr/local/go/src/go/token/position.go:499
	}
//line /usr/local/go/src/go/token/position.go:499
	// _ = "end of CoverTab[49228]"
//line /usr/local/go/src/go/token/position.go:499
	_go_fuzz_dep_.CoverTab[49229]++
							return nil
//line /usr/local/go/src/go/token/position.go:500
	// _ = "end of CoverTab[49229]"
}

// File returns the file that contains the position p.
//line /usr/local/go/src/go/token/position.go:503
// If no such file is found (for instance for p == NoPos),
//line /usr/local/go/src/go/token/position.go:503
// the result is nil.
//line /usr/local/go/src/go/token/position.go:506
func (s *FileSet) File(p Pos) (f *File) {
//line /usr/local/go/src/go/token/position.go:506
	_go_fuzz_dep_.CoverTab[49238]++
							if p != NoPos {
//line /usr/local/go/src/go/token/position.go:507
		_go_fuzz_dep_.CoverTab[49240]++
								f = s.file(p)
//line /usr/local/go/src/go/token/position.go:508
		// _ = "end of CoverTab[49240]"
	} else {
//line /usr/local/go/src/go/token/position.go:509
		_go_fuzz_dep_.CoverTab[49241]++
//line /usr/local/go/src/go/token/position.go:509
		// _ = "end of CoverTab[49241]"
//line /usr/local/go/src/go/token/position.go:509
	}
//line /usr/local/go/src/go/token/position.go:509
	// _ = "end of CoverTab[49238]"
//line /usr/local/go/src/go/token/position.go:509
	_go_fuzz_dep_.CoverTab[49239]++
							return
//line /usr/local/go/src/go/token/position.go:510
	// _ = "end of CoverTab[49239]"
}

// PositionFor converts a Pos p in the fileset into a Position value.
//line /usr/local/go/src/go/token/position.go:513
// If adjusted is set, the position may be adjusted by position-altering
//line /usr/local/go/src/go/token/position.go:513
// //line comments; otherwise those comments are ignored.
//line /usr/local/go/src/go/token/position.go:513
// p must be a Pos value in s or NoPos.
//line /usr/local/go/src/go/token/position.go:517
func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position) {
//line /usr/local/go/src/go/token/position.go:517
	_go_fuzz_dep_.CoverTab[49242]++
							if p != NoPos {
//line /usr/local/go/src/go/token/position.go:518
		_go_fuzz_dep_.CoverTab[49244]++
								if f := s.file(p); f != nil {
//line /usr/local/go/src/go/token/position.go:519
			_go_fuzz_dep_.CoverTab[49245]++
									return f.position(p, adjusted)
//line /usr/local/go/src/go/token/position.go:520
			// _ = "end of CoverTab[49245]"
		} else {
//line /usr/local/go/src/go/token/position.go:521
			_go_fuzz_dep_.CoverTab[49246]++
//line /usr/local/go/src/go/token/position.go:521
			// _ = "end of CoverTab[49246]"
//line /usr/local/go/src/go/token/position.go:521
		}
//line /usr/local/go/src/go/token/position.go:521
		// _ = "end of CoverTab[49244]"
	} else {
//line /usr/local/go/src/go/token/position.go:522
		_go_fuzz_dep_.CoverTab[49247]++
//line /usr/local/go/src/go/token/position.go:522
		// _ = "end of CoverTab[49247]"
//line /usr/local/go/src/go/token/position.go:522
	}
//line /usr/local/go/src/go/token/position.go:522
	// _ = "end of CoverTab[49242]"
//line /usr/local/go/src/go/token/position.go:522
	_go_fuzz_dep_.CoverTab[49243]++
							return
//line /usr/local/go/src/go/token/position.go:523
	// _ = "end of CoverTab[49243]"
}

// Position converts a Pos p in the fileset into a Position value.
//line /usr/local/go/src/go/token/position.go:526
// Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).
//line /usr/local/go/src/go/token/position.go:528
func (s *FileSet) Position(p Pos) (pos Position) {
//line /usr/local/go/src/go/token/position.go:528
	_go_fuzz_dep_.CoverTab[49248]++
							return s.PositionFor(p, true)
//line /usr/local/go/src/go/token/position.go:529
	// _ = "end of CoverTab[49248]"
}

//line /usr/local/go/src/go/token/position.go:535
func searchInts(a []int, x int) int {
//line /usr/local/go/src/go/token/position.go:535
	_go_fuzz_dep_.CoverTab[49249]++

//line /usr/local/go/src/go/token/position.go:545
	i, j := 0, len(a)
	for i < j {
//line /usr/local/go/src/go/token/position.go:546
		_go_fuzz_dep_.CoverTab[49251]++
								h := int(uint(i+j) >> 1)

								if a[h] <= x {
//line /usr/local/go/src/go/token/position.go:549
			_go_fuzz_dep_.CoverTab[49252]++
									i = h + 1
//line /usr/local/go/src/go/token/position.go:550
			// _ = "end of CoverTab[49252]"
		} else {
//line /usr/local/go/src/go/token/position.go:551
			_go_fuzz_dep_.CoverTab[49253]++
									j = h
//line /usr/local/go/src/go/token/position.go:552
			// _ = "end of CoverTab[49253]"
		}
//line /usr/local/go/src/go/token/position.go:553
		// _ = "end of CoverTab[49251]"
	}
//line /usr/local/go/src/go/token/position.go:554
	// _ = "end of CoverTab[49249]"
//line /usr/local/go/src/go/token/position.go:554
	_go_fuzz_dep_.CoverTab[49250]++
							return i - 1
//line /usr/local/go/src/go/token/position.go:555
	// _ = "end of CoverTab[49250]"
}

//line /usr/local/go/src/go/token/position.go:556
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/go/token/position.go:556
var _ = _go_fuzz_dep_.CoverTab
