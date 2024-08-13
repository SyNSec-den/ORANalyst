// Copyright © 2016 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
package jwalterweatherman

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:6
)

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type Threshold int

func (t Threshold) String() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:17
	_go_fuzz_dep_.CoverTab[119231]++
												return prefixes[t]
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:18
	// _ = "end of CoverTab[119231]"
}

const (
	LevelTrace	Threshold	= iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelCritical
	LevelFatal
)

var prefixes map[Threshold]string = map[Threshold]string{
	LevelTrace:	"TRACE",
	LevelDebug:	"DEBUG",
	LevelInfo:	"INFO",
	LevelWarn:	"WARN",
	LevelError:	"ERROR",
	LevelCritical:	"CRITICAL",
	LevelFatal:	"FATAL",
}

// Notepad is where you leave a note!
type Notepad struct {
	TRACE		*log.Logger
	DEBUG		*log.Logger
	INFO		*log.Logger
	WARN		*log.Logger
	ERROR		*log.Logger
	CRITICAL	*log.Logger
	FATAL		*log.Logger

	LOG		*log.Logger
	FEEDBACK	*Feedback

	loggers		[7]**log.Logger
	logHandle	io.Writer
	outHandle	io.Writer
	logThreshold	Threshold
	stdoutThreshold	Threshold
	prefix		string
	flags		int

	logListeners	[]LogListener
}

// A LogListener can ble supplied to a Notepad to listen on log writes for a given
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// threshold. This can be used to capture log events in unit tests and similar.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// Note that this function will be invoked once for each log threshold. If
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// the given threshold is not of interest to you, return nil.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// Note that these listeners will receive log events for a given threshold, even
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// if the current configuration says not to log it. That way you can count ERRORs even
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:65
// if you don't print them to the console.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:72
type LogListener func(t Threshold) io.Writer

// NewNotepad creates a new Notepad.
func NewNotepad(
	outThreshold Threshold,
	logThreshold Threshold,
	outHandle, logHandle io.Writer,
	prefix string, flags int,
	logListeners ...LogListener,
) *Notepad {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:81
	_go_fuzz_dep_.CoverTab[119232]++

												n := &Notepad{logListeners: logListeners}

												n.loggers = [7]**log.Logger{&n.TRACE, &n.DEBUG, &n.INFO, &n.WARN, &n.ERROR, &n.CRITICAL, &n.FATAL}
												n.outHandle = outHandle
												n.logHandle = logHandle
												n.stdoutThreshold = outThreshold
												n.logThreshold = logThreshold

												if len(prefix) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:91
		_go_fuzz_dep_.CoverTab[119234]++
													n.prefix = "[" + prefix + "] "
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:92
		// _ = "end of CoverTab[119234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:93
		_go_fuzz_dep_.CoverTab[119235]++
													n.prefix = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:94
		// _ = "end of CoverTab[119235]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:95
	// _ = "end of CoverTab[119232]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:95
	_go_fuzz_dep_.CoverTab[119233]++

												n.flags = flags

												n.LOG = log.New(n.logHandle,
		"LOG:   ",
		n.flags)
												n.FEEDBACK = &Feedback{out: log.New(outHandle, "", 0), log: n.LOG}

												n.init()
												return n
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:105
	// _ = "end of CoverTab[119233]"
}

// init creates the loggers for each level depending on the notepad thresholds.
func (n *Notepad) init() {
	logAndOut := io.MultiWriter(n.outHandle, n.logHandle)

	for t, logger := range n.loggers {
		threshold := Threshold(t)
		prefix := n.prefix + threshold.String() + " "

		switch {
		case threshold >= n.logThreshold && threshold >= n.stdoutThreshold:
			*logger = log.New(n.createLogWriters(threshold, logAndOut), prefix, n.flags)

		case threshold >= n.logThreshold:
			*logger = log.New(n.createLogWriters(threshold, n.logHandle), prefix, n.flags)

		case threshold >= n.stdoutThreshold:
			*logger = log.New(n.createLogWriters(threshold, n.outHandle), prefix, n.flags)

		default:
			*logger = log.New(n.createLogWriters(threshold, ioutil.Discard), prefix, n.flags)
		}
	}
}

func (n *Notepad) createLogWriters(t Threshold, handle io.Writer) io.Writer {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:132
	_go_fuzz_dep_.CoverTab[119236]++
												if len(n.logListeners) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:133
		_go_fuzz_dep_.CoverTab[119240]++
													return handle
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:134
		// _ = "end of CoverTab[119240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:135
		_go_fuzz_dep_.CoverTab[119241]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:135
		// _ = "end of CoverTab[119241]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:135
	// _ = "end of CoverTab[119236]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:135
	_go_fuzz_dep_.CoverTab[119237]++
												writers := []io.Writer{handle}
												for _, l := range n.logListeners {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:137
		_go_fuzz_dep_.CoverTab[119242]++
													w := l(t)
													if w != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:139
			_go_fuzz_dep_.CoverTab[119243]++
														writers = append(writers, w)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:140
			// _ = "end of CoverTab[119243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:141
			_go_fuzz_dep_.CoverTab[119244]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:141
			// _ = "end of CoverTab[119244]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:141
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:141
		// _ = "end of CoverTab[119242]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:142
	// _ = "end of CoverTab[119237]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:142
	_go_fuzz_dep_.CoverTab[119238]++

												if len(writers) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:144
		_go_fuzz_dep_.CoverTab[119245]++
													return handle
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:145
		// _ = "end of CoverTab[119245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:146
		_go_fuzz_dep_.CoverTab[119246]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:146
		// _ = "end of CoverTab[119246]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:146
	// _ = "end of CoverTab[119238]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:146
	_go_fuzz_dep_.CoverTab[119239]++

												return io.MultiWriter(writers...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:148
	// _ = "end of CoverTab[119239]"
}

// SetLogThreshold changes the threshold above which messages are written to the
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:151
// log file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:153
func (n *Notepad) SetLogThreshold(threshold Threshold) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:153
	_go_fuzz_dep_.CoverTab[119247]++
												n.logThreshold = threshold
												n.init()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:155
	// _ = "end of CoverTab[119247]"
}

// SetLogOutput changes the file where log messages are written.
func (n *Notepad) SetLogOutput(handle io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:159
	_go_fuzz_dep_.CoverTab[119248]++
												n.logHandle = handle
												n.init()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:161
	// _ = "end of CoverTab[119248]"
}

// GetStdoutThreshold returns the defined Treshold for the log logger.
func (n *Notepad) GetLogThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:165
	_go_fuzz_dep_.CoverTab[119249]++
												return n.logThreshold
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:166
	// _ = "end of CoverTab[119249]"
}

// SetStdoutThreshold changes the threshold above which messages are written to the
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:169
// standard output.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:171
func (n *Notepad) SetStdoutThreshold(threshold Threshold) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:171
	_go_fuzz_dep_.CoverTab[119250]++
												n.stdoutThreshold = threshold
												n.init()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:173
	// _ = "end of CoverTab[119250]"
}

// GetStdoutThreshold returns the Treshold for the stdout logger.
func (n *Notepad) GetStdoutThreshold() Threshold {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:177
	_go_fuzz_dep_.CoverTab[119251]++
												return n.stdoutThreshold
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:178
	// _ = "end of CoverTab[119251]"
}

// SetPrefix changes the prefix used by the notepad. Prefixes are displayed between
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:181
// brackets at the beginning of the line. An empty prefix won't be displayed at all.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:183
func (n *Notepad) SetPrefix(prefix string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:183
	_go_fuzz_dep_.CoverTab[119252]++
												if len(prefix) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:184
		_go_fuzz_dep_.CoverTab[119254]++
													n.prefix = "[" + prefix + "] "
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:185
		// _ = "end of CoverTab[119254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:186
		_go_fuzz_dep_.CoverTab[119255]++
													n.prefix = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:187
		// _ = "end of CoverTab[119255]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:188
	// _ = "end of CoverTab[119252]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:188
	_go_fuzz_dep_.CoverTab[119253]++
												n.init()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:189
	// _ = "end of CoverTab[119253]"
}

// SetFlags choose which flags the logger will display (after prefix and message
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:192
// level). See the package log for more informations on this.
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:194
func (n *Notepad) SetFlags(flags int) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:194
	_go_fuzz_dep_.CoverTab[119256]++
												n.flags = flags
												n.init()
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:196
	// _ = "end of CoverTab[119256]"
}

// Feedback writes plainly to the outHandle while
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:199
// logging with the standard extra information (date, file, etc).
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:201
type Feedback struct {
	out	*log.Logger
	log	*log.Logger
}

func (fb *Feedback) Println(v ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:206
	_go_fuzz_dep_.CoverTab[119257]++
												fb.output(fmt.Sprintln(v...))
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:207
	// _ = "end of CoverTab[119257]"
}

func (fb *Feedback) Printf(format string, v ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:210
	_go_fuzz_dep_.CoverTab[119258]++
												fb.output(fmt.Sprintf(format, v...))
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:211
	// _ = "end of CoverTab[119258]"
}

func (fb *Feedback) Print(v ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:214
	_go_fuzz_dep_.CoverTab[119259]++
												fb.output(fmt.Sprint(v...))
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:215
	// _ = "end of CoverTab[119259]"
}

func (fb *Feedback) output(s string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:218
	_go_fuzz_dep_.CoverTab[119260]++
												if fb.out != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:219
		_go_fuzz_dep_.CoverTab[119262]++
													fb.out.Output(2, s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:220
		// _ = "end of CoverTab[119262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:221
		_go_fuzz_dep_.CoverTab[119263]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:221
		// _ = "end of CoverTab[119263]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:221
	// _ = "end of CoverTab[119260]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:221
	_go_fuzz_dep_.CoverTab[119261]++
												if fb.log != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:222
		_go_fuzz_dep_.CoverTab[119264]++
													fb.log.Output(2, s)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:223
		// _ = "end of CoverTab[119264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:224
		_go_fuzz_dep_.CoverTab[119265]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:224
		// _ = "end of CoverTab[119265]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:224
	// _ = "end of CoverTab[119261]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:225
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/notepad.go:225
var _ = _go_fuzz_dep_.CoverTab
