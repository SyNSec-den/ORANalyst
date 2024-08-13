// Copyright © 2016 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
package jwalterweatherman

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:6
)

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TRACE		*log.Logger
	DEBUG		*log.Logger
	INFO		*log.Logger
	WARN		*log.Logger
	ERROR		*log.Logger
	CRITICAL	*log.Logger
	FATAL		*log.Logger

	LOG		*log.Logger
	FEEDBACK	*Feedback

	defaultNotepad	*Notepad
)

func reloadDefaultNotepad() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:30
	_go_fuzz_dep_.CoverTab[119210]++
													TRACE = defaultNotepad.TRACE
													DEBUG = defaultNotepad.DEBUG
													INFO = defaultNotepad.INFO
													WARN = defaultNotepad.WARN
													ERROR = defaultNotepad.ERROR
													CRITICAL = defaultNotepad.CRITICAL
													FATAL = defaultNotepad.FATAL

													LOG = defaultNotepad.LOG
													FEEDBACK = defaultNotepad.FEEDBACK
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:40
	// _ = "end of CoverTab[119210]"
}

func init() {
	defaultNotepad = NewNotepad(LevelError, LevelWarn, os.Stdout, ioutil.Discard, "", log.Ldate|log.Ltime)
	reloadDefaultNotepad()
}

// SetLogThreshold set the log threshold for the default notepad. Trace by default.
func SetLogThreshold(threshold Threshold) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:49
	_go_fuzz_dep_.CoverTab[119211]++
													defaultNotepad.SetLogThreshold(threshold)
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:51
	// _ = "end of CoverTab[119211]"
}

// SetLogOutput set the log output for the default notepad. Discarded by default.
func SetLogOutput(handle io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:55
	_go_fuzz_dep_.CoverTab[119212]++
													defaultNotepad.SetLogOutput(handle)
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:57
	// _ = "end of CoverTab[119212]"
}

// SetStdoutThreshold set the standard output threshold for the default notepad.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:60
// Info by default.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:62
func SetStdoutThreshold(threshold Threshold) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:62
	_go_fuzz_dep_.CoverTab[119213]++
													defaultNotepad.SetStdoutThreshold(threshold)
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:64
	// _ = "end of CoverTab[119213]"
}

// SetStdoutOutput set the stdout output for the default notepad. Default is stdout.
func SetStdoutOutput(handle io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:68
	_go_fuzz_dep_.CoverTab[119214]++
													defaultNotepad.outHandle = handle
													defaultNotepad.init()
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:71
	// _ = "end of CoverTab[119214]"
}

// SetPrefix set the prefix for the default logger. Empty by default.
func SetPrefix(prefix string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:75
	_go_fuzz_dep_.CoverTab[119215]++
													defaultNotepad.SetPrefix(prefix)
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:77
	// _ = "end of CoverTab[119215]"
}

// SetFlags set the flags for the default logger. "log.Ldate | log.Ltime" by default.
func SetFlags(flags int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:81
	_go_fuzz_dep_.CoverTab[119216]++
													defaultNotepad.SetFlags(flags)
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:83
	// _ = "end of CoverTab[119216]"
}

// SetLogListeners configures the default logger with one or more log listeners.
func SetLogListeners(l ...LogListener) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:87
	_go_fuzz_dep_.CoverTab[119217]++
													defaultNotepad.logListeners = l
													defaultNotepad.init()
													reloadDefaultNotepad()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:90
	// _ = "end of CoverTab[119217]"
}

// Level returns the current global log threshold.
func LogThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:94
	_go_fuzz_dep_.CoverTab[119218]++
													return defaultNotepad.logThreshold
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:95
	// _ = "end of CoverTab[119218]"
}

// Level returns the current global output threshold.
func StdoutThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:99
	_go_fuzz_dep_.CoverTab[119219]++
													return defaultNotepad.stdoutThreshold
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:100
	// _ = "end of CoverTab[119219]"
}

// GetStdoutThreshold returns the defined Treshold for the log logger.
func GetLogThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:104
	_go_fuzz_dep_.CoverTab[119220]++
													return defaultNotepad.GetLogThreshold()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:105
	// _ = "end of CoverTab[119220]"
}

// GetStdoutThreshold returns the Treshold for the stdout logger.
func GetStdoutThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:109
	_go_fuzz_dep_.CoverTab[119221]++
													return defaultNotepad.GetStdoutThreshold()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:110
	// _ = "end of CoverTab[119221]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:111
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/default_notepad.go:111
var _ = _go_fuzz_dep_.CoverTab
