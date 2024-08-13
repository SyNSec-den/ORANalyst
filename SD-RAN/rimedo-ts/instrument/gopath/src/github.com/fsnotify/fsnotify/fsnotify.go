// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !plan9
// +build !plan9

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:8
// Package fsnotify provides a platform-independent interface for file system notifications.
package fsnotify

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:9
)

import (
	"bytes"
	"errors"
	"fmt"
)

// Event represents a single file system notification.
type Event struct {
	Name	string	// Relative path to the file or directory.
	Op	Op	// File operation that triggered the event.
}

// Op describes a set of file operations.
type Op uint32

// These are the generalized file operations that can trigger a notification.
const (
	Create	Op	= 1 << iota
	Write
	Remove
	Rename
	Chmod
)

func (op Op) String() string {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:35
	_go_fuzz_dep_.CoverTab[114985]++
	// Use a buffer for efficient string concatenation
	var buffer bytes.Buffer

	if op&Create == Create {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:39
		_go_fuzz_dep_.CoverTab[114992]++
												buffer.WriteString("|CREATE")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:40
		// _ = "end of CoverTab[114992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:41
		_go_fuzz_dep_.CoverTab[114993]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:41
		// _ = "end of CoverTab[114993]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:41
	// _ = "end of CoverTab[114985]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:41
	_go_fuzz_dep_.CoverTab[114986]++
											if op&Remove == Remove {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:42
		_go_fuzz_dep_.CoverTab[114994]++
												buffer.WriteString("|REMOVE")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:43
		// _ = "end of CoverTab[114994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:44
		_go_fuzz_dep_.CoverTab[114995]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:44
		// _ = "end of CoverTab[114995]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:44
	// _ = "end of CoverTab[114986]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:44
	_go_fuzz_dep_.CoverTab[114987]++
											if op&Write == Write {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:45
		_go_fuzz_dep_.CoverTab[114996]++
												buffer.WriteString("|WRITE")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:46
		// _ = "end of CoverTab[114996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:47
		_go_fuzz_dep_.CoverTab[114997]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:47
		// _ = "end of CoverTab[114997]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:47
	// _ = "end of CoverTab[114987]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:47
	_go_fuzz_dep_.CoverTab[114988]++
											if op&Rename == Rename {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:48
		_go_fuzz_dep_.CoverTab[114998]++
												buffer.WriteString("|RENAME")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:49
		// _ = "end of CoverTab[114998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:50
		_go_fuzz_dep_.CoverTab[114999]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:50
		// _ = "end of CoverTab[114999]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:50
	// _ = "end of CoverTab[114988]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:50
	_go_fuzz_dep_.CoverTab[114989]++
											if op&Chmod == Chmod {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:51
		_go_fuzz_dep_.CoverTab[115000]++
												buffer.WriteString("|CHMOD")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:52
		// _ = "end of CoverTab[115000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:53
		_go_fuzz_dep_.CoverTab[115001]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:53
		// _ = "end of CoverTab[115001]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:53
	// _ = "end of CoverTab[114989]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:53
	_go_fuzz_dep_.CoverTab[114990]++
											if buffer.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:54
		_go_fuzz_dep_.CoverTab[115002]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:55
		// _ = "end of CoverTab[115002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:56
		_go_fuzz_dep_.CoverTab[115003]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:56
		// _ = "end of CoverTab[115003]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:56
	// _ = "end of CoverTab[114990]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:56
	_go_fuzz_dep_.CoverTab[114991]++
											return buffer.String()[1:]
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:57
	// _ = "end of CoverTab[114991]"
}

// String returns a string representation of the event in the form
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:60
// "file: REMOVE|WRITE|..."
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:62
func (e Event) String() string {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:62
	_go_fuzz_dep_.CoverTab[115004]++
											return fmt.Sprintf("%q: %s", e.Name, e.Op.String())
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:63
	// _ = "end of CoverTab[115004]"
}

// Common errors that can be reported by a watcher
var (
	ErrEventOverflow = errors.New("fsnotify queue overflow")
)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/fsnotify.go:69
var _ = _go_fuzz_dep_.CoverTab
