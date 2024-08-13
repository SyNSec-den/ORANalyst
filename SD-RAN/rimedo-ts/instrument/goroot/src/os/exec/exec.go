// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/os/exec/exec.go:5
// Package exec runs external commands. It wraps os.StartProcess to make it
//line /usr/local/go/src/os/exec/exec.go:5
// easier to remap stdin and stdout, connect I/O with pipes, and do other
//line /usr/local/go/src/os/exec/exec.go:5
// adjustments.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Unlike the "system" library call from C and other languages, the
//line /usr/local/go/src/os/exec/exec.go:5
// os/exec package intentionally does not invoke the system shell and
//line /usr/local/go/src/os/exec/exec.go:5
// does not expand any glob patterns or handle other expansions,
//line /usr/local/go/src/os/exec/exec.go:5
// pipelines, or redirections typically done by shells. The package
//line /usr/local/go/src/os/exec/exec.go:5
// behaves more like C's "exec" family of functions. To expand glob
//line /usr/local/go/src/os/exec/exec.go:5
// patterns, either call the shell directly, taking care to escape any
//line /usr/local/go/src/os/exec/exec.go:5
// dangerous input, or use the path/filepath package's Glob function.
//line /usr/local/go/src/os/exec/exec.go:5
// To expand environment variables, use package os's ExpandEnv.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Note that the examples in this package assume a Unix system.
//line /usr/local/go/src/os/exec/exec.go:5
// They may not run on Windows, and they do not run in the Go Playground
//line /usr/local/go/src/os/exec/exec.go:5
// used by golang.org and godoc.org.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// # Executables in the current directory
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// The functions Command and LookPath look for a program
//line /usr/local/go/src/os/exec/exec.go:5
// in the directories listed in the current path, following the
//line /usr/local/go/src/os/exec/exec.go:5
// conventions of the host operating system.
//line /usr/local/go/src/os/exec/exec.go:5
// Operating systems have for decades included the current
//line /usr/local/go/src/os/exec/exec.go:5
// directory in this search, sometimes implicitly and sometimes
//line /usr/local/go/src/os/exec/exec.go:5
// configured explicitly that way by default.
//line /usr/local/go/src/os/exec/exec.go:5
// Modern practice is that including the current directory
//line /usr/local/go/src/os/exec/exec.go:5
// is usually unexpected and often leads to security problems.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// To avoid those security problems, as of Go 1.19, this package will not resolve a program
//line /usr/local/go/src/os/exec/exec.go:5
// using an implicit or explicit path entry relative to the current directory.
//line /usr/local/go/src/os/exec/exec.go:5
// That is, if you run exec.LookPath("go"), it will not successfully return
//line /usr/local/go/src/os/exec/exec.go:5
// ./go on Unix nor .\go.exe on Windows, no matter how the path is configured.
//line /usr/local/go/src/os/exec/exec.go:5
// Instead, if the usual path algorithms would result in that answer,
//line /usr/local/go/src/os/exec/exec.go:5
// these functions return an error err satisfying errors.Is(err, ErrDot).
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// For example, consider these two program snippets:
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
//	path, err := exec.LookPath("prog")
//line /usr/local/go/src/os/exec/exec.go:5
//	if err != nil {
//line /usr/local/go/src/os/exec/exec.go:5
//		log.Fatal(err)
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//	use(path)
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// and
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
//	cmd := exec.Command("prog")
//line /usr/local/go/src/os/exec/exec.go:5
//	if err := cmd.Run(); err != nil {
//line /usr/local/go/src/os/exec/exec.go:5
//		log.Fatal(err)
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// These will not find and run ./prog or .\prog.exe,
//line /usr/local/go/src/os/exec/exec.go:5
// no matter how the current path is configured.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Code that always wants to run a program from the current directory
//line /usr/local/go/src/os/exec/exec.go:5
// can be rewritten to say "./prog" instead of "prog".
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Code that insists on including results from relative path entries
//line /usr/local/go/src/os/exec/exec.go:5
// can instead override the error using an errors.Is check:
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
//	path, err := exec.LookPath("prog")
//line /usr/local/go/src/os/exec/exec.go:5
//	if errors.Is(err, exec.ErrDot) {
//line /usr/local/go/src/os/exec/exec.go:5
//		err = nil
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//	if err != nil {
//line /usr/local/go/src/os/exec/exec.go:5
//		log.Fatal(err)
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//	use(path)
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// and
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
//	cmd := exec.Command("prog")
//line /usr/local/go/src/os/exec/exec.go:5
//	if errors.Is(cmd.Err, exec.ErrDot) {
//line /usr/local/go/src/os/exec/exec.go:5
//		cmd.Err = nil
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//	if err := cmd.Run(); err != nil {
//line /usr/local/go/src/os/exec/exec.go:5
//		log.Fatal(err)
//line /usr/local/go/src/os/exec/exec.go:5
//	}
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Setting the environment variable GODEBUG=execerrdot=0
//line /usr/local/go/src/os/exec/exec.go:5
// disables generation of ErrDot entirely, temporarily restoring the pre-Go 1.19
//line /usr/local/go/src/os/exec/exec.go:5
// behavior for programs that are unable to apply more targeted fixes.
//line /usr/local/go/src/os/exec/exec.go:5
// A future version of Go may remove support for this variable.
//line /usr/local/go/src/os/exec/exec.go:5
//
//line /usr/local/go/src/os/exec/exec.go:5
// Before adding such overrides, make sure you understand the
//line /usr/local/go/src/os/exec/exec.go:5
// security implications of doing so.
//line /usr/local/go/src/os/exec/exec.go:5
// See https://go.dev/blog/path-security for more information.
//line /usr/local/go/src/os/exec/exec.go:91
package exec

//line /usr/local/go/src/os/exec/exec.go:91
import (
//line /usr/local/go/src/os/exec/exec.go:91
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/exec/exec.go:91
)
//line /usr/local/go/src/os/exec/exec.go:91
import (
//line /usr/local/go/src/os/exec/exec.go:91
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/exec/exec.go:91
)

import (
	"bytes"
	"context"
	"errors"
	"internal/godebug"
	"internal/syscall/execenv"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Error is returned by LookPath when it fails to classify a file as an
//line /usr/local/go/src/os/exec/exec.go:109
// executable.
//line /usr/local/go/src/os/exec/exec.go:111
type Error struct {
	// Name is the file name for which the error occurred.
	Name	string
	// Err is the underlying error.
	Err	error
}

func (e *Error) Error() string {
//line /usr/local/go/src/os/exec/exec.go:118
	_go_fuzz_dep_.CoverTab[107081]++
						return "exec: " + strconv.Quote(e.Name) + ": " + e.Err.Error()
//line /usr/local/go/src/os/exec/exec.go:119
	// _ = "end of CoverTab[107081]"
}

func (e *Error) Unwrap() error {
//line /usr/local/go/src/os/exec/exec.go:122
	_go_fuzz_dep_.CoverTab[107082]++
//line /usr/local/go/src/os/exec/exec.go:122
	return e.Err
//line /usr/local/go/src/os/exec/exec.go:122
	// _ = "end of CoverTab[107082]"
//line /usr/local/go/src/os/exec/exec.go:122
}

// ErrWaitDelay is returned by (*Cmd).Wait if the process exits with a
//line /usr/local/go/src/os/exec/exec.go:124
// successful status code but its output pipes are not closed before the
//line /usr/local/go/src/os/exec/exec.go:124
// command's WaitDelay expires.
//line /usr/local/go/src/os/exec/exec.go:127
var ErrWaitDelay = errors.New("exec: WaitDelay expired before I/O complete")

// wrappedError wraps an error without relying on fmt.Errorf.
type wrappedError struct {
	prefix	string
	err	error
}

func (w wrappedError) Error() string {
//line /usr/local/go/src/os/exec/exec.go:135
	_go_fuzz_dep_.CoverTab[107083]++
						return w.prefix + ": " + w.err.Error()
//line /usr/local/go/src/os/exec/exec.go:136
	// _ = "end of CoverTab[107083]"
}

func (w wrappedError) Unwrap() error {
//line /usr/local/go/src/os/exec/exec.go:139
	_go_fuzz_dep_.CoverTab[107084]++
						return w.err
//line /usr/local/go/src/os/exec/exec.go:140
	// _ = "end of CoverTab[107084]"
}

// Cmd represents an external command being prepared or run.
//line /usr/local/go/src/os/exec/exec.go:143
//
//line /usr/local/go/src/os/exec/exec.go:143
// A Cmd cannot be reused after calling its Run, Output or CombinedOutput
//line /usr/local/go/src/os/exec/exec.go:143
// methods.
//line /usr/local/go/src/os/exec/exec.go:147
type Cmd struct {
	// Path is the path of the command to run.
	//
	// This is the only field that must be set to a non-zero
	// value. If Path is relative, it is evaluated relative
	// to Dir.
	Path	string

	// Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	//
	// In typical use, both Path and Args are set by calling Command.
	Args	[]string

	// Env specifies the environment of the process.
	// Each entry is of the form "key=value".
	// If Env is nil, the new process uses the current process's
	// environment.
	// If Env contains duplicate environment keys, only the last
	// value in the slice for each duplicate key is used.
	// As a special case on Windows, SYSTEMROOT is always added if
	// missing and not explicitly set to the empty string.
	Env	[]string

	// Dir specifies the working directory of the command.
	// If Dir is the empty string, Run runs the command in the
	// calling process's current directory.
	Dir	string

	// Stdin specifies the process's standard input.
	//
	// If Stdin is nil, the process reads from the null device (os.DevNull).
	//
	// If Stdin is an *os.File, the process's standard input is connected
	// directly to that file.
	//
	// Otherwise, during the execution of the command a separate
	// goroutine reads from Stdin and delivers that data to the command
	// over a pipe. In this case, Wait does not complete until the goroutine
	// stops copying, either because it has reached the end of Stdin
	// (EOF or a read error), or because writing to the pipe returned an error,
	// or because a nonzero WaitDelay was set and expired.
	Stdin	io.Reader

	// Stdout and Stderr specify the process's standard output and error.
	//
	// If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	//
	// If either is an *os.File, the corresponding output from the process
	// is connected directly to that file.
	//
	// Otherwise, during the execution of the command a separate goroutine
	// reads from the process over a pipe and delivers that data to the
	// corresponding Writer. In this case, Wait does not complete until the
	// goroutine reaches EOF or encounters an error or a nonzero WaitDelay
	// expires.
	//
	// If Stdout and Stderr are the same writer, and have a type that can
	// be compared with ==, at most one goroutine at a time will call Write.
	Stdout	io.Writer
	Stderr	io.Writer

	// ExtraFiles specifies additional open files to be inherited by the
	// new process. It does not include standard input, standard output, or
	// standard error. If non-nil, entry i becomes file descriptor 3+i.
	//
	// ExtraFiles is not supported on Windows.
	ExtraFiles	[]*os.File

	// SysProcAttr holds optional, operating system-specific attributes.
	// Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
	SysProcAttr	*syscall.SysProcAttr

	// Process is the underlying process, once started.
	Process	*os.Process

	// ProcessState contains information about an exited process.
	// If the process was started successfully, Wait or Run will
	// populate its ProcessState when the command completes.
	ProcessState	*os.ProcessState

	// ctx is the context passed to CommandContext, if any.
	ctx	context.Context

	Err	error	// LookPath error, if any.

	// If Cancel is non-nil, the command must have been created with
	// CommandContext and Cancel will be called when the command's
	// Context is done. By default, CommandContext sets Cancel to
	// call the Kill method on the command's Process.
	//
	// Typically a custom Cancel will send a signal to the command's
	// Process, but it may instead take other actions to initiate cancellation,
	// such as closing a stdin or stdout pipe or sending a shutdown request on a
	// network socket.
	//
	// If the command exits with a success status after Cancel is
	// called, and Cancel does not return an error equivalent to
	// os.ErrProcessDone, then Wait and similar methods will return a non-nil
	// error: either an error wrapping the one returned by Cancel,
	// or the error from the Context.
	// (If the command exits with a non-success status, or Cancel
	// returns an error that wraps os.ErrProcessDone, Wait and similar methods
	// continue to return the command's usual exit status.)
	//
	// If Cancel is set to nil, nothing will happen immediately when the command's
	// Context is done, but a nonzero WaitDelay will still take effect. That may
	// be useful, for example, to work around deadlocks in commands that do not
	// support shutdown signals but are expected to always finish quickly.
	//
	// Cancel will not be called if Start returns a non-nil error.
	Cancel	func() error

	// If WaitDelay is non-zero, it bounds the time spent waiting on two sources
	// of unexpected delay in Wait: a child process that fails to exit after the
	// associated Context is canceled, and a child process that exits but leaves
	// its I/O pipes unclosed.
	//
	// The WaitDelay timer starts when either the associated Context is done or a
	// call to Wait observes that the child process has exited, whichever occurs
	// first. When the delay has elapsed, the command shuts down the child process
	// and/or its I/O pipes.
	//
	// If the child process has failed to exit — perhaps because it ignored or
	// failed to receive a shutdown signal from a Cancel function, or because no
	// Cancel function was set — then it will be terminated using os.Process.Kill.
	//
	// Then, if the I/O pipes communicating with the child process are still open,
	// those pipes are closed in order to unblock any goroutines currently blocked
	// on Read or Write calls.
	//
	// If pipes are closed due to WaitDelay, no Cancel call has occurred,
	// and the command has otherwise exited with a successful status, Wait and
	// similar methods will return ErrWaitDelay instead of nil.
	//
	// If WaitDelay is zero (the default), I/O pipes will be read until EOF,
	// which might not occur until orphaned subprocesses of the command have
	// also closed their descriptors for the pipes.
	WaitDelay	time.Duration

	// childIOFiles holds closers for any of the child process's
	// stdin, stdout, and/or stderr files that were opened by the Cmd itself
	// (not supplied by the caller). These should be closed as soon as they
	// are inherited by the child process.
	childIOFiles	[]io.Closer

	// parentIOPipes holds closers for the parent's end of any pipes
	// connected to the child's stdin, stdout, and/or stderr streams
	// that were opened by the Cmd itself (not supplied by the caller).
	// These should be closed after Wait sees the command and copying
	// goroutines exit, or after WaitDelay has expired.
	parentIOPipes	[]io.Closer

	// goroutine holds a set of closures to execute to copy data
	// to and/or from the command's I/O pipes.
	goroutine	[]func() error

	// If goroutineErr is non-nil, it receives the first error from a copying
	// goroutine once all such goroutines have completed.
	// goroutineErr is set to nil once its error has been received.
	goroutineErr	<-chan error

	// If ctxResult is non-nil, it receives the result of watchCtx exactly once.
	ctxResult	<-chan ctxResult

	// The stack saved when the Command was created, if GODEBUG contains
	// execwait=2. Used for debugging leaks.
	createdByStack	[]byte

	// For a security release long ago, we created x/sys/execabs,
	// which manipulated the unexported lookPathErr error field
	// in this struct. For Go 1.19 we exported the field as Err error,
	// above, but we have to keep lookPathErr around for use by
	// old programs building against new toolchains.
	// The String and Start methods look for an error in lookPathErr
	// in preference to Err, to preserve the errors that execabs sets.
	//
	// In general we don't guarantee misuse of reflect like this,
	// but the misuse of reflect was by us, the best of various bad
	// options to fix the security problem, and people depend on
	// those old copies of execabs continuing to work.
	// The result is that we have to leave this variable around for the
	// rest of time, a compatibility scar.
	//
	// See https://go.dev/blog/path-security
	// and https://go.dev/issue/43724 for more context.
	lookPathErr	error
}

// A ctxResult reports the result of watching the Context associated with a
//line /usr/local/go/src/os/exec/exec.go:337
// running command (and sending corresponding signals if needed).
//line /usr/local/go/src/os/exec/exec.go:339
type ctxResult struct {
	err	error

	// If timer is non-nil, it expires after WaitDelay has elapsed after
	// the Context is done.
	//
	// (If timer is nil, that means that the Context was not done before the
	// command completed, or no WaitDelay was set, or the WaitDelay already
	// expired and its effect was already applied.)
	timer	*time.Timer
}

var execwait = godebug.New("execwait")
var execerrdot = godebug.New("execerrdot")

// Command returns the Cmd struct to execute the named program with
//line /usr/local/go/src/os/exec/exec.go:354
// the given arguments.
//line /usr/local/go/src/os/exec/exec.go:354
//
//line /usr/local/go/src/os/exec/exec.go:354
// It sets only the Path and Args in the returned structure.
//line /usr/local/go/src/os/exec/exec.go:354
//
//line /usr/local/go/src/os/exec/exec.go:354
// If name contains no path separators, Command uses LookPath to
//line /usr/local/go/src/os/exec/exec.go:354
// resolve name to a complete path if possible. Otherwise it uses name
//line /usr/local/go/src/os/exec/exec.go:354
// directly as Path.
//line /usr/local/go/src/os/exec/exec.go:354
//
//line /usr/local/go/src/os/exec/exec.go:354
// The returned Cmd's Args field is constructed from the command name
//line /usr/local/go/src/os/exec/exec.go:354
// followed by the elements of arg, so arg should not include the
//line /usr/local/go/src/os/exec/exec.go:354
// command name itself. For example, Command("echo", "hello").
//line /usr/local/go/src/os/exec/exec.go:354
// Args[0] is always name, not the possibly resolved Path.
//line /usr/local/go/src/os/exec/exec.go:354
//
//line /usr/local/go/src/os/exec/exec.go:354
// On Windows, processes receive the whole command line as a single string
//line /usr/local/go/src/os/exec/exec.go:354
// and do their own parsing. Command combines and quotes Args into a command
//line /usr/local/go/src/os/exec/exec.go:354
// line string with an algorithm compatible with applications using
//line /usr/local/go/src/os/exec/exec.go:354
// CommandLineToArgvW (which is the most common way). Notable exceptions are
//line /usr/local/go/src/os/exec/exec.go:354
// msiexec.exe and cmd.exe (and thus, all batch files), which have a different
//line /usr/local/go/src/os/exec/exec.go:354
// unquoting algorithm. In these or other similar cases, you can do the
//line /usr/local/go/src/os/exec/exec.go:354
// quoting yourself and provide the full command line in SysProcAttr.CmdLine,
//line /usr/local/go/src/os/exec/exec.go:354
// leaving Args empty.
//line /usr/local/go/src/os/exec/exec.go:376
func Command(name string, arg ...string) *Cmd {
//line /usr/local/go/src/os/exec/exec.go:376
	_go_fuzz_dep_.CoverTab[107085]++
						cmd := &Cmd{
		Path:	name,
		Args:	append([]string{name}, arg...),
	}

	if v := execwait.Value(); v != "" {
//line /usr/local/go/src/os/exec/exec.go:382
		_go_fuzz_dep_.CoverTab[107088]++
							if v == "2" {
//line /usr/local/go/src/os/exec/exec.go:383
			_go_fuzz_dep_.CoverTab[107090]++

//line /usr/local/go/src/os/exec/exec.go:386
			stack := make([]byte, 1024)
			for {
//line /usr/local/go/src/os/exec/exec.go:387
				_go_fuzz_dep_.CoverTab[107093]++
									n := runtime.Stack(stack, false)
									if n < len(stack) {
//line /usr/local/go/src/os/exec/exec.go:389
					_go_fuzz_dep_.CoverTab[107095]++
										stack = stack[:n]
										break
//line /usr/local/go/src/os/exec/exec.go:391
					// _ = "end of CoverTab[107095]"
				} else {
//line /usr/local/go/src/os/exec/exec.go:392
					_go_fuzz_dep_.CoverTab[107096]++
//line /usr/local/go/src/os/exec/exec.go:392
					// _ = "end of CoverTab[107096]"
//line /usr/local/go/src/os/exec/exec.go:392
				}
//line /usr/local/go/src/os/exec/exec.go:392
				// _ = "end of CoverTab[107093]"
//line /usr/local/go/src/os/exec/exec.go:392
				_go_fuzz_dep_.CoverTab[107094]++
									stack = make([]byte, 2*len(stack))
//line /usr/local/go/src/os/exec/exec.go:393
				// _ = "end of CoverTab[107094]"
			}
//line /usr/local/go/src/os/exec/exec.go:394
			// _ = "end of CoverTab[107090]"
//line /usr/local/go/src/os/exec/exec.go:394
			_go_fuzz_dep_.CoverTab[107091]++

								if i := bytes.Index(stack, []byte("\nos/exec.Command(")); i >= 0 {
//line /usr/local/go/src/os/exec/exec.go:396
				_go_fuzz_dep_.CoverTab[107097]++
									stack = stack[i+1:]
//line /usr/local/go/src/os/exec/exec.go:397
				// _ = "end of CoverTab[107097]"
			} else {
//line /usr/local/go/src/os/exec/exec.go:398
				_go_fuzz_dep_.CoverTab[107098]++
//line /usr/local/go/src/os/exec/exec.go:398
				// _ = "end of CoverTab[107098]"
//line /usr/local/go/src/os/exec/exec.go:398
			}
//line /usr/local/go/src/os/exec/exec.go:398
			// _ = "end of CoverTab[107091]"
//line /usr/local/go/src/os/exec/exec.go:398
			_go_fuzz_dep_.CoverTab[107092]++
								cmd.createdByStack = stack
//line /usr/local/go/src/os/exec/exec.go:399
			// _ = "end of CoverTab[107092]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:400
			_go_fuzz_dep_.CoverTab[107099]++
//line /usr/local/go/src/os/exec/exec.go:400
			// _ = "end of CoverTab[107099]"
//line /usr/local/go/src/os/exec/exec.go:400
		}
//line /usr/local/go/src/os/exec/exec.go:400
		// _ = "end of CoverTab[107088]"
//line /usr/local/go/src/os/exec/exec.go:400
		_go_fuzz_dep_.CoverTab[107089]++

							runtime.SetFinalizer(cmd, func(c *Cmd) {
//line /usr/local/go/src/os/exec/exec.go:402
			_go_fuzz_dep_.CoverTab[107100]++
								if c.Process != nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:403
				_go_fuzz_dep_.CoverTab[107101]++
//line /usr/local/go/src/os/exec/exec.go:403
				return c.ProcessState == nil
//line /usr/local/go/src/os/exec/exec.go:403
				// _ = "end of CoverTab[107101]"
//line /usr/local/go/src/os/exec/exec.go:403
			}() {
//line /usr/local/go/src/os/exec/exec.go:403
				_go_fuzz_dep_.CoverTab[107102]++
									debugHint := ""
									if c.createdByStack == nil {
//line /usr/local/go/src/os/exec/exec.go:405
					_go_fuzz_dep_.CoverTab[107104]++
										debugHint = " (set GODEBUG=execwait=2 to capture stacks for debugging)"
//line /usr/local/go/src/os/exec/exec.go:406
					// _ = "end of CoverTab[107104]"
				} else {
//line /usr/local/go/src/os/exec/exec.go:407
					_go_fuzz_dep_.CoverTab[107105]++
										os.Stderr.WriteString("GODEBUG=execwait=2 detected a leaked exec.Cmd created by:\n")
										os.Stderr.Write(c.createdByStack)
										os.Stderr.WriteString("\n")
										debugHint = ""
//line /usr/local/go/src/os/exec/exec.go:411
					// _ = "end of CoverTab[107105]"
				}
//line /usr/local/go/src/os/exec/exec.go:412
				// _ = "end of CoverTab[107102]"
//line /usr/local/go/src/os/exec/exec.go:412
				_go_fuzz_dep_.CoverTab[107103]++
									panic("exec: Cmd started a Process but leaked without a call to Wait" + debugHint)
//line /usr/local/go/src/os/exec/exec.go:413
				// _ = "end of CoverTab[107103]"
			} else {
//line /usr/local/go/src/os/exec/exec.go:414
				_go_fuzz_dep_.CoverTab[107106]++
//line /usr/local/go/src/os/exec/exec.go:414
				// _ = "end of CoverTab[107106]"
//line /usr/local/go/src/os/exec/exec.go:414
			}
//line /usr/local/go/src/os/exec/exec.go:414
			// _ = "end of CoverTab[107100]"
		})
//line /usr/local/go/src/os/exec/exec.go:415
		// _ = "end of CoverTab[107089]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:416
		_go_fuzz_dep_.CoverTab[107107]++
//line /usr/local/go/src/os/exec/exec.go:416
		// _ = "end of CoverTab[107107]"
//line /usr/local/go/src/os/exec/exec.go:416
	}
//line /usr/local/go/src/os/exec/exec.go:416
	// _ = "end of CoverTab[107085]"
//line /usr/local/go/src/os/exec/exec.go:416
	_go_fuzz_dep_.CoverTab[107086]++

						if filepath.Base(name) == name {
//line /usr/local/go/src/os/exec/exec.go:418
		_go_fuzz_dep_.CoverTab[107108]++
							lp, err := LookPath(name)
							if lp != "" {
//line /usr/local/go/src/os/exec/exec.go:420
			_go_fuzz_dep_.CoverTab[107110]++

//line /usr/local/go/src/os/exec/exec.go:424
			cmd.Path = lp
//line /usr/local/go/src/os/exec/exec.go:424
			// _ = "end of CoverTab[107110]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:425
			_go_fuzz_dep_.CoverTab[107111]++
//line /usr/local/go/src/os/exec/exec.go:425
			// _ = "end of CoverTab[107111]"
//line /usr/local/go/src/os/exec/exec.go:425
		}
//line /usr/local/go/src/os/exec/exec.go:425
		// _ = "end of CoverTab[107108]"
//line /usr/local/go/src/os/exec/exec.go:425
		_go_fuzz_dep_.CoverTab[107109]++
							if err != nil {
//line /usr/local/go/src/os/exec/exec.go:426
			_go_fuzz_dep_.CoverTab[107112]++
								cmd.Err = err
//line /usr/local/go/src/os/exec/exec.go:427
			// _ = "end of CoverTab[107112]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:428
			_go_fuzz_dep_.CoverTab[107113]++
//line /usr/local/go/src/os/exec/exec.go:428
			// _ = "end of CoverTab[107113]"
//line /usr/local/go/src/os/exec/exec.go:428
		}
//line /usr/local/go/src/os/exec/exec.go:428
		// _ = "end of CoverTab[107109]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:429
		_go_fuzz_dep_.CoverTab[107114]++
//line /usr/local/go/src/os/exec/exec.go:429
		// _ = "end of CoverTab[107114]"
//line /usr/local/go/src/os/exec/exec.go:429
	}
//line /usr/local/go/src/os/exec/exec.go:429
	// _ = "end of CoverTab[107086]"
//line /usr/local/go/src/os/exec/exec.go:429
	_go_fuzz_dep_.CoverTab[107087]++
						return cmd
//line /usr/local/go/src/os/exec/exec.go:430
	// _ = "end of CoverTab[107087]"
}

// CommandContext is like Command but includes a context.
//line /usr/local/go/src/os/exec/exec.go:433
//
//line /usr/local/go/src/os/exec/exec.go:433
// The provided context is used to interrupt the process
//line /usr/local/go/src/os/exec/exec.go:433
// (by calling cmd.Cancel or os.Process.Kill)
//line /usr/local/go/src/os/exec/exec.go:433
// if the context becomes done before the command completes on its own.
//line /usr/local/go/src/os/exec/exec.go:433
//
//line /usr/local/go/src/os/exec/exec.go:433
// CommandContext sets the command's Cancel function to invoke the Kill method
//line /usr/local/go/src/os/exec/exec.go:433
// on its Process, and leaves its WaitDelay unset. The caller may change the
//line /usr/local/go/src/os/exec/exec.go:433
// cancellation behavior by modifying those fields before starting the command.
//line /usr/local/go/src/os/exec/exec.go:442
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
//line /usr/local/go/src/os/exec/exec.go:442
	_go_fuzz_dep_.CoverTab[107115]++
						if ctx == nil {
//line /usr/local/go/src/os/exec/exec.go:443
		_go_fuzz_dep_.CoverTab[107118]++
							panic("nil Context")
//line /usr/local/go/src/os/exec/exec.go:444
		// _ = "end of CoverTab[107118]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:445
		_go_fuzz_dep_.CoverTab[107119]++
//line /usr/local/go/src/os/exec/exec.go:445
		// _ = "end of CoverTab[107119]"
//line /usr/local/go/src/os/exec/exec.go:445
	}
//line /usr/local/go/src/os/exec/exec.go:445
	// _ = "end of CoverTab[107115]"
//line /usr/local/go/src/os/exec/exec.go:445
	_go_fuzz_dep_.CoverTab[107116]++
						cmd := Command(name, arg...)
						cmd.ctx = ctx
						cmd.Cancel = func() error {
//line /usr/local/go/src/os/exec/exec.go:448
		_go_fuzz_dep_.CoverTab[107120]++
							return cmd.Process.Kill()
//line /usr/local/go/src/os/exec/exec.go:449
		// _ = "end of CoverTab[107120]"
	}
//line /usr/local/go/src/os/exec/exec.go:450
	// _ = "end of CoverTab[107116]"
//line /usr/local/go/src/os/exec/exec.go:450
	_go_fuzz_dep_.CoverTab[107117]++
						return cmd
//line /usr/local/go/src/os/exec/exec.go:451
	// _ = "end of CoverTab[107117]"
}

// String returns a human-readable description of c.
//line /usr/local/go/src/os/exec/exec.go:454
// It is intended only for debugging.
//line /usr/local/go/src/os/exec/exec.go:454
// In particular, it is not suitable for use as input to a shell.
//line /usr/local/go/src/os/exec/exec.go:454
// The output of String may vary across Go releases.
//line /usr/local/go/src/os/exec/exec.go:458
func (c *Cmd) String() string {
//line /usr/local/go/src/os/exec/exec.go:458
	_go_fuzz_dep_.CoverTab[107121]++
						if c.Err != nil || func() bool {
//line /usr/local/go/src/os/exec/exec.go:459
		_go_fuzz_dep_.CoverTab[107124]++
//line /usr/local/go/src/os/exec/exec.go:459
		return c.lookPathErr != nil
//line /usr/local/go/src/os/exec/exec.go:459
		// _ = "end of CoverTab[107124]"
//line /usr/local/go/src/os/exec/exec.go:459
	}() {
//line /usr/local/go/src/os/exec/exec.go:459
		_go_fuzz_dep_.CoverTab[107125]++

							return strings.Join(c.Args, " ")
//line /usr/local/go/src/os/exec/exec.go:461
		// _ = "end of CoverTab[107125]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:462
		_go_fuzz_dep_.CoverTab[107126]++
//line /usr/local/go/src/os/exec/exec.go:462
		// _ = "end of CoverTab[107126]"
//line /usr/local/go/src/os/exec/exec.go:462
	}
//line /usr/local/go/src/os/exec/exec.go:462
	// _ = "end of CoverTab[107121]"
//line /usr/local/go/src/os/exec/exec.go:462
	_go_fuzz_dep_.CoverTab[107122]++

						b := new(strings.Builder)
						b.WriteString(c.Path)
						for _, a := range c.Args[1:] {
//line /usr/local/go/src/os/exec/exec.go:466
		_go_fuzz_dep_.CoverTab[107127]++
							b.WriteByte(' ')
							b.WriteString(a)
//line /usr/local/go/src/os/exec/exec.go:468
		// _ = "end of CoverTab[107127]"
	}
//line /usr/local/go/src/os/exec/exec.go:469
	// _ = "end of CoverTab[107122]"
//line /usr/local/go/src/os/exec/exec.go:469
	_go_fuzz_dep_.CoverTab[107123]++
						return b.String()
//line /usr/local/go/src/os/exec/exec.go:470
	// _ = "end of CoverTab[107123]"
}

// interfaceEqual protects against panics from doing equality tests on
//line /usr/local/go/src/os/exec/exec.go:473
// two interfaces with non-comparable underlying types.
//line /usr/local/go/src/os/exec/exec.go:475
func interfaceEqual(a, b any) bool {
//line /usr/local/go/src/os/exec/exec.go:475
	_go_fuzz_dep_.CoverTab[107128]++
						defer func() {
//line /usr/local/go/src/os/exec/exec.go:476
		_go_fuzz_dep_.CoverTab[107130]++
							recover()
//line /usr/local/go/src/os/exec/exec.go:477
		// _ = "end of CoverTab[107130]"
	}()
//line /usr/local/go/src/os/exec/exec.go:478
	// _ = "end of CoverTab[107128]"
//line /usr/local/go/src/os/exec/exec.go:478
	_go_fuzz_dep_.CoverTab[107129]++
						return a == b
//line /usr/local/go/src/os/exec/exec.go:479
	// _ = "end of CoverTab[107129]"
}

func (c *Cmd) argv() []string {
//line /usr/local/go/src/os/exec/exec.go:482
	_go_fuzz_dep_.CoverTab[107131]++
						if len(c.Args) > 0 {
//line /usr/local/go/src/os/exec/exec.go:483
		_go_fuzz_dep_.CoverTab[107133]++
							return c.Args
//line /usr/local/go/src/os/exec/exec.go:484
		// _ = "end of CoverTab[107133]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:485
		_go_fuzz_dep_.CoverTab[107134]++
//line /usr/local/go/src/os/exec/exec.go:485
		// _ = "end of CoverTab[107134]"
//line /usr/local/go/src/os/exec/exec.go:485
	}
//line /usr/local/go/src/os/exec/exec.go:485
	// _ = "end of CoverTab[107131]"
//line /usr/local/go/src/os/exec/exec.go:485
	_go_fuzz_dep_.CoverTab[107132]++
						return []string{c.Path}
//line /usr/local/go/src/os/exec/exec.go:486
	// _ = "end of CoverTab[107132]"
}

func (c *Cmd) childStdin() (*os.File, error) {
//line /usr/local/go/src/os/exec/exec.go:489
	_go_fuzz_dep_.CoverTab[107135]++
						if c.Stdin == nil {
//line /usr/local/go/src/os/exec/exec.go:490
		_go_fuzz_dep_.CoverTab[107140]++
							f, err := os.Open(os.DevNull)
							if err != nil {
//line /usr/local/go/src/os/exec/exec.go:492
			_go_fuzz_dep_.CoverTab[107142]++
								return nil, err
//line /usr/local/go/src/os/exec/exec.go:493
			// _ = "end of CoverTab[107142]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:494
			_go_fuzz_dep_.CoverTab[107143]++
//line /usr/local/go/src/os/exec/exec.go:494
			// _ = "end of CoverTab[107143]"
//line /usr/local/go/src/os/exec/exec.go:494
		}
//line /usr/local/go/src/os/exec/exec.go:494
		// _ = "end of CoverTab[107140]"
//line /usr/local/go/src/os/exec/exec.go:494
		_go_fuzz_dep_.CoverTab[107141]++
							c.childIOFiles = append(c.childIOFiles, f)
							return f, nil
//line /usr/local/go/src/os/exec/exec.go:496
		// _ = "end of CoverTab[107141]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:497
		_go_fuzz_dep_.CoverTab[107144]++
//line /usr/local/go/src/os/exec/exec.go:497
		// _ = "end of CoverTab[107144]"
//line /usr/local/go/src/os/exec/exec.go:497
	}
//line /usr/local/go/src/os/exec/exec.go:497
	// _ = "end of CoverTab[107135]"
//line /usr/local/go/src/os/exec/exec.go:497
	_go_fuzz_dep_.CoverTab[107136]++

						if f, ok := c.Stdin.(*os.File); ok {
//line /usr/local/go/src/os/exec/exec.go:499
		_go_fuzz_dep_.CoverTab[107145]++
							return f, nil
//line /usr/local/go/src/os/exec/exec.go:500
		// _ = "end of CoverTab[107145]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:501
		_go_fuzz_dep_.CoverTab[107146]++
//line /usr/local/go/src/os/exec/exec.go:501
		// _ = "end of CoverTab[107146]"
//line /usr/local/go/src/os/exec/exec.go:501
	}
//line /usr/local/go/src/os/exec/exec.go:501
	// _ = "end of CoverTab[107136]"
//line /usr/local/go/src/os/exec/exec.go:501
	_go_fuzz_dep_.CoverTab[107137]++

						pr, pw, err := os.Pipe()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:504
		_go_fuzz_dep_.CoverTab[107147]++
							return nil, err
//line /usr/local/go/src/os/exec/exec.go:505
		// _ = "end of CoverTab[107147]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:506
		_go_fuzz_dep_.CoverTab[107148]++
//line /usr/local/go/src/os/exec/exec.go:506
		// _ = "end of CoverTab[107148]"
//line /usr/local/go/src/os/exec/exec.go:506
	}
//line /usr/local/go/src/os/exec/exec.go:506
	// _ = "end of CoverTab[107137]"
//line /usr/local/go/src/os/exec/exec.go:506
	_go_fuzz_dep_.CoverTab[107138]++

						c.childIOFiles = append(c.childIOFiles, pr)
						c.parentIOPipes = append(c.parentIOPipes, pw)
						c.goroutine = append(c.goroutine, func() error {
//line /usr/local/go/src/os/exec/exec.go:510
		_go_fuzz_dep_.CoverTab[107149]++
							_, err := io.Copy(pw, c.Stdin)
							if skipStdinCopyError(err) {
//line /usr/local/go/src/os/exec/exec.go:512
			_go_fuzz_dep_.CoverTab[107152]++
								err = nil
//line /usr/local/go/src/os/exec/exec.go:513
			// _ = "end of CoverTab[107152]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:514
			_go_fuzz_dep_.CoverTab[107153]++
//line /usr/local/go/src/os/exec/exec.go:514
			// _ = "end of CoverTab[107153]"
//line /usr/local/go/src/os/exec/exec.go:514
		}
//line /usr/local/go/src/os/exec/exec.go:514
		// _ = "end of CoverTab[107149]"
//line /usr/local/go/src/os/exec/exec.go:514
		_go_fuzz_dep_.CoverTab[107150]++
							if err1 := pw.Close(); err == nil {
//line /usr/local/go/src/os/exec/exec.go:515
			_go_fuzz_dep_.CoverTab[107154]++
								err = err1
//line /usr/local/go/src/os/exec/exec.go:516
			// _ = "end of CoverTab[107154]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:517
			_go_fuzz_dep_.CoverTab[107155]++
//line /usr/local/go/src/os/exec/exec.go:517
			// _ = "end of CoverTab[107155]"
//line /usr/local/go/src/os/exec/exec.go:517
		}
//line /usr/local/go/src/os/exec/exec.go:517
		// _ = "end of CoverTab[107150]"
//line /usr/local/go/src/os/exec/exec.go:517
		_go_fuzz_dep_.CoverTab[107151]++
							return err
//line /usr/local/go/src/os/exec/exec.go:518
		// _ = "end of CoverTab[107151]"
	})
//line /usr/local/go/src/os/exec/exec.go:519
	// _ = "end of CoverTab[107138]"
//line /usr/local/go/src/os/exec/exec.go:519
	_go_fuzz_dep_.CoverTab[107139]++
						return pr, nil
//line /usr/local/go/src/os/exec/exec.go:520
	// _ = "end of CoverTab[107139]"
}

func (c *Cmd) childStdout() (*os.File, error) {
//line /usr/local/go/src/os/exec/exec.go:523
	_go_fuzz_dep_.CoverTab[107156]++
						return c.writerDescriptor(c.Stdout)
//line /usr/local/go/src/os/exec/exec.go:524
	// _ = "end of CoverTab[107156]"
}

func (c *Cmd) childStderr(childStdout *os.File) (*os.File, error) {
//line /usr/local/go/src/os/exec/exec.go:527
	_go_fuzz_dep_.CoverTab[107157]++
						if c.Stderr != nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:528
		_go_fuzz_dep_.CoverTab[107159]++
//line /usr/local/go/src/os/exec/exec.go:528
		return interfaceEqual(c.Stderr, c.Stdout)
//line /usr/local/go/src/os/exec/exec.go:528
		// _ = "end of CoverTab[107159]"
//line /usr/local/go/src/os/exec/exec.go:528
	}() {
//line /usr/local/go/src/os/exec/exec.go:528
		_go_fuzz_dep_.CoverTab[107160]++
							return childStdout, nil
//line /usr/local/go/src/os/exec/exec.go:529
		// _ = "end of CoverTab[107160]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:530
		_go_fuzz_dep_.CoverTab[107161]++
//line /usr/local/go/src/os/exec/exec.go:530
		// _ = "end of CoverTab[107161]"
//line /usr/local/go/src/os/exec/exec.go:530
	}
//line /usr/local/go/src/os/exec/exec.go:530
	// _ = "end of CoverTab[107157]"
//line /usr/local/go/src/os/exec/exec.go:530
	_go_fuzz_dep_.CoverTab[107158]++
						return c.writerDescriptor(c.Stderr)
//line /usr/local/go/src/os/exec/exec.go:531
	// _ = "end of CoverTab[107158]"
}

// writerDescriptor returns an os.File to which the child process
//line /usr/local/go/src/os/exec/exec.go:534
// can write to send data to w.
//line /usr/local/go/src/os/exec/exec.go:534
//
//line /usr/local/go/src/os/exec/exec.go:534
// If w is nil, writerDescriptor returns a File that writes to os.DevNull.
//line /usr/local/go/src/os/exec/exec.go:538
func (c *Cmd) writerDescriptor(w io.Writer) (*os.File, error) {
//line /usr/local/go/src/os/exec/exec.go:538
	_go_fuzz_dep_.CoverTab[107162]++
						if w == nil {
//line /usr/local/go/src/os/exec/exec.go:539
		_go_fuzz_dep_.CoverTab[107167]++
							f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
							if err != nil {
//line /usr/local/go/src/os/exec/exec.go:541
			_go_fuzz_dep_.CoverTab[107169]++
								return nil, err
//line /usr/local/go/src/os/exec/exec.go:542
			// _ = "end of CoverTab[107169]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:543
			_go_fuzz_dep_.CoverTab[107170]++
//line /usr/local/go/src/os/exec/exec.go:543
			// _ = "end of CoverTab[107170]"
//line /usr/local/go/src/os/exec/exec.go:543
		}
//line /usr/local/go/src/os/exec/exec.go:543
		// _ = "end of CoverTab[107167]"
//line /usr/local/go/src/os/exec/exec.go:543
		_go_fuzz_dep_.CoverTab[107168]++
							c.childIOFiles = append(c.childIOFiles, f)
							return f, nil
//line /usr/local/go/src/os/exec/exec.go:545
		// _ = "end of CoverTab[107168]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:546
		_go_fuzz_dep_.CoverTab[107171]++
//line /usr/local/go/src/os/exec/exec.go:546
		// _ = "end of CoverTab[107171]"
//line /usr/local/go/src/os/exec/exec.go:546
	}
//line /usr/local/go/src/os/exec/exec.go:546
	// _ = "end of CoverTab[107162]"
//line /usr/local/go/src/os/exec/exec.go:546
	_go_fuzz_dep_.CoverTab[107163]++

						if f, ok := w.(*os.File); ok {
//line /usr/local/go/src/os/exec/exec.go:548
		_go_fuzz_dep_.CoverTab[107172]++
							return f, nil
//line /usr/local/go/src/os/exec/exec.go:549
		// _ = "end of CoverTab[107172]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:550
		_go_fuzz_dep_.CoverTab[107173]++
//line /usr/local/go/src/os/exec/exec.go:550
		// _ = "end of CoverTab[107173]"
//line /usr/local/go/src/os/exec/exec.go:550
	}
//line /usr/local/go/src/os/exec/exec.go:550
	// _ = "end of CoverTab[107163]"
//line /usr/local/go/src/os/exec/exec.go:550
	_go_fuzz_dep_.CoverTab[107164]++

						pr, pw, err := os.Pipe()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:553
		_go_fuzz_dep_.CoverTab[107174]++
							return nil, err
//line /usr/local/go/src/os/exec/exec.go:554
		// _ = "end of CoverTab[107174]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:555
		_go_fuzz_dep_.CoverTab[107175]++
//line /usr/local/go/src/os/exec/exec.go:555
		// _ = "end of CoverTab[107175]"
//line /usr/local/go/src/os/exec/exec.go:555
	}
//line /usr/local/go/src/os/exec/exec.go:555
	// _ = "end of CoverTab[107164]"
//line /usr/local/go/src/os/exec/exec.go:555
	_go_fuzz_dep_.CoverTab[107165]++

						c.childIOFiles = append(c.childIOFiles, pw)
						c.parentIOPipes = append(c.parentIOPipes, pr)
						c.goroutine = append(c.goroutine, func() error {
//line /usr/local/go/src/os/exec/exec.go:559
		_go_fuzz_dep_.CoverTab[107176]++
							_, err := io.Copy(w, pr)
							pr.Close()
							return err
//line /usr/local/go/src/os/exec/exec.go:562
		// _ = "end of CoverTab[107176]"
	})
//line /usr/local/go/src/os/exec/exec.go:563
	// _ = "end of CoverTab[107165]"
//line /usr/local/go/src/os/exec/exec.go:563
	_go_fuzz_dep_.CoverTab[107166]++
						return pw, nil
//line /usr/local/go/src/os/exec/exec.go:564
	// _ = "end of CoverTab[107166]"
}

func closeDescriptors(closers []io.Closer) {
//line /usr/local/go/src/os/exec/exec.go:567
	_go_fuzz_dep_.CoverTab[107177]++
						for _, fd := range closers {
//line /usr/local/go/src/os/exec/exec.go:568
		_go_fuzz_dep_.CoverTab[107178]++
							fd.Close()
//line /usr/local/go/src/os/exec/exec.go:569
		// _ = "end of CoverTab[107178]"
	}
//line /usr/local/go/src/os/exec/exec.go:570
	// _ = "end of CoverTab[107177]"
}

// Run starts the specified command and waits for it to complete.
//line /usr/local/go/src/os/exec/exec.go:573
//
//line /usr/local/go/src/os/exec/exec.go:573
// The returned error is nil if the command runs, has no problems
//line /usr/local/go/src/os/exec/exec.go:573
// copying stdin, stdout, and stderr, and exits with a zero exit
//line /usr/local/go/src/os/exec/exec.go:573
// status.
//line /usr/local/go/src/os/exec/exec.go:573
//
//line /usr/local/go/src/os/exec/exec.go:573
// If the command starts but does not complete successfully, the error is of
//line /usr/local/go/src/os/exec/exec.go:573
// type *ExitError. Other error types may be returned for other situations.
//line /usr/local/go/src/os/exec/exec.go:573
//
//line /usr/local/go/src/os/exec/exec.go:573
// If the calling goroutine has locked the operating system thread
//line /usr/local/go/src/os/exec/exec.go:573
// with runtime.LockOSThread and modified any inheritable OS-level
//line /usr/local/go/src/os/exec/exec.go:573
// thread state (for example, Linux or Plan 9 name spaces), the new
//line /usr/local/go/src/os/exec/exec.go:573
// process will inherit the caller's thread state.
//line /usr/local/go/src/os/exec/exec.go:586
func (c *Cmd) Run() error {
//line /usr/local/go/src/os/exec/exec.go:586
	_go_fuzz_dep_.CoverTab[107179]++
						if err := c.Start(); err != nil {
//line /usr/local/go/src/os/exec/exec.go:587
		_go_fuzz_dep_.CoverTab[107181]++
							return err
//line /usr/local/go/src/os/exec/exec.go:588
		// _ = "end of CoverTab[107181]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:589
		_go_fuzz_dep_.CoverTab[107182]++
//line /usr/local/go/src/os/exec/exec.go:589
		// _ = "end of CoverTab[107182]"
//line /usr/local/go/src/os/exec/exec.go:589
	}
//line /usr/local/go/src/os/exec/exec.go:589
	// _ = "end of CoverTab[107179]"
//line /usr/local/go/src/os/exec/exec.go:589
	_go_fuzz_dep_.CoverTab[107180]++
						return c.Wait()
//line /usr/local/go/src/os/exec/exec.go:590
	// _ = "end of CoverTab[107180]"
}

// lookExtensions finds windows executable by its dir and path.
//line /usr/local/go/src/os/exec/exec.go:593
// It uses LookPath to try appropriate extensions.
//line /usr/local/go/src/os/exec/exec.go:593
// lookExtensions does not search PATH, instead it converts `prog` into `.\prog`.
//line /usr/local/go/src/os/exec/exec.go:596
func lookExtensions(path, dir string) (string, error) {
//line /usr/local/go/src/os/exec/exec.go:596
	_go_fuzz_dep_.CoverTab[107183]++
						if filepath.Base(path) == path {
//line /usr/local/go/src/os/exec/exec.go:597
		_go_fuzz_dep_.CoverTab[107189]++
							path = "." + string(filepath.Separator) + path
//line /usr/local/go/src/os/exec/exec.go:598
		// _ = "end of CoverTab[107189]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:599
		_go_fuzz_dep_.CoverTab[107190]++
//line /usr/local/go/src/os/exec/exec.go:599
		// _ = "end of CoverTab[107190]"
//line /usr/local/go/src/os/exec/exec.go:599
	}
//line /usr/local/go/src/os/exec/exec.go:599
	// _ = "end of CoverTab[107183]"
//line /usr/local/go/src/os/exec/exec.go:599
	_go_fuzz_dep_.CoverTab[107184]++
						if dir == "" {
//line /usr/local/go/src/os/exec/exec.go:600
		_go_fuzz_dep_.CoverTab[107191]++
							return LookPath(path)
//line /usr/local/go/src/os/exec/exec.go:601
		// _ = "end of CoverTab[107191]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:602
		_go_fuzz_dep_.CoverTab[107192]++
//line /usr/local/go/src/os/exec/exec.go:602
		// _ = "end of CoverTab[107192]"
//line /usr/local/go/src/os/exec/exec.go:602
	}
//line /usr/local/go/src/os/exec/exec.go:602
	// _ = "end of CoverTab[107184]"
//line /usr/local/go/src/os/exec/exec.go:602
	_go_fuzz_dep_.CoverTab[107185]++
						if filepath.VolumeName(path) != "" {
//line /usr/local/go/src/os/exec/exec.go:603
		_go_fuzz_dep_.CoverTab[107193]++
							return LookPath(path)
//line /usr/local/go/src/os/exec/exec.go:604
		// _ = "end of CoverTab[107193]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:605
		_go_fuzz_dep_.CoverTab[107194]++
//line /usr/local/go/src/os/exec/exec.go:605
		// _ = "end of CoverTab[107194]"
//line /usr/local/go/src/os/exec/exec.go:605
	}
//line /usr/local/go/src/os/exec/exec.go:605
	// _ = "end of CoverTab[107185]"
//line /usr/local/go/src/os/exec/exec.go:605
	_go_fuzz_dep_.CoverTab[107186]++
						if len(path) > 1 && func() bool {
//line /usr/local/go/src/os/exec/exec.go:606
		_go_fuzz_dep_.CoverTab[107195]++
//line /usr/local/go/src/os/exec/exec.go:606
		return os.IsPathSeparator(path[0])
//line /usr/local/go/src/os/exec/exec.go:606
		// _ = "end of CoverTab[107195]"
//line /usr/local/go/src/os/exec/exec.go:606
	}() {
//line /usr/local/go/src/os/exec/exec.go:606
		_go_fuzz_dep_.CoverTab[107196]++
							return LookPath(path)
//line /usr/local/go/src/os/exec/exec.go:607
		// _ = "end of CoverTab[107196]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:608
		_go_fuzz_dep_.CoverTab[107197]++
//line /usr/local/go/src/os/exec/exec.go:608
		// _ = "end of CoverTab[107197]"
//line /usr/local/go/src/os/exec/exec.go:608
	}
//line /usr/local/go/src/os/exec/exec.go:608
	// _ = "end of CoverTab[107186]"
//line /usr/local/go/src/os/exec/exec.go:608
	_go_fuzz_dep_.CoverTab[107187]++
						dirandpath := filepath.Join(dir, path)

						lp, err := LookPath(dirandpath)
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:612
		_go_fuzz_dep_.CoverTab[107198]++
							return "", err
//line /usr/local/go/src/os/exec/exec.go:613
		// _ = "end of CoverTab[107198]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:614
		_go_fuzz_dep_.CoverTab[107199]++
//line /usr/local/go/src/os/exec/exec.go:614
		// _ = "end of CoverTab[107199]"
//line /usr/local/go/src/os/exec/exec.go:614
	}
//line /usr/local/go/src/os/exec/exec.go:614
	// _ = "end of CoverTab[107187]"
//line /usr/local/go/src/os/exec/exec.go:614
	_go_fuzz_dep_.CoverTab[107188]++
						ext := strings.TrimPrefix(lp, dirandpath)
						return path + ext, nil
//line /usr/local/go/src/os/exec/exec.go:616
	// _ = "end of CoverTab[107188]"
}

// Start starts the specified command but does not wait for it to complete.
//line /usr/local/go/src/os/exec/exec.go:619
//
//line /usr/local/go/src/os/exec/exec.go:619
// If Start returns successfully, the c.Process field will be set.
//line /usr/local/go/src/os/exec/exec.go:619
//
//line /usr/local/go/src/os/exec/exec.go:619
// After a successful call to Start the Wait method must be called in
//line /usr/local/go/src/os/exec/exec.go:619
// order to release associated system resources.
//line /usr/local/go/src/os/exec/exec.go:625
func (c *Cmd) Start() error {
//line /usr/local/go/src/os/exec/exec.go:625
	_go_fuzz_dep_.CoverTab[107200]++

//line /usr/local/go/src/os/exec/exec.go:628
	if c.Process != nil {
//line /usr/local/go/src/os/exec/exec.go:628
		_go_fuzz_dep_.CoverTab[107215]++
							return errors.New("exec: already started")
//line /usr/local/go/src/os/exec/exec.go:629
		// _ = "end of CoverTab[107215]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:630
		_go_fuzz_dep_.CoverTab[107216]++
//line /usr/local/go/src/os/exec/exec.go:630
		// _ = "end of CoverTab[107216]"
//line /usr/local/go/src/os/exec/exec.go:630
	}
//line /usr/local/go/src/os/exec/exec.go:630
	// _ = "end of CoverTab[107200]"
//line /usr/local/go/src/os/exec/exec.go:630
	_go_fuzz_dep_.CoverTab[107201]++

						started := false
						defer func() {
//line /usr/local/go/src/os/exec/exec.go:633
		_go_fuzz_dep_.CoverTab[107217]++
							closeDescriptors(c.childIOFiles)
							c.childIOFiles = nil

							if !started {
//line /usr/local/go/src/os/exec/exec.go:637
			_go_fuzz_dep_.CoverTab[107218]++
								closeDescriptors(c.parentIOPipes)
								c.parentIOPipes = nil
//line /usr/local/go/src/os/exec/exec.go:639
			// _ = "end of CoverTab[107218]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:640
			_go_fuzz_dep_.CoverTab[107219]++
//line /usr/local/go/src/os/exec/exec.go:640
			// _ = "end of CoverTab[107219]"
//line /usr/local/go/src/os/exec/exec.go:640
		}
//line /usr/local/go/src/os/exec/exec.go:640
		// _ = "end of CoverTab[107217]"
	}()
//line /usr/local/go/src/os/exec/exec.go:641
	// _ = "end of CoverTab[107201]"
//line /usr/local/go/src/os/exec/exec.go:641
	_go_fuzz_dep_.CoverTab[107202]++

						if c.Path == "" && func() bool {
//line /usr/local/go/src/os/exec/exec.go:643
		_go_fuzz_dep_.CoverTab[107220]++
//line /usr/local/go/src/os/exec/exec.go:643
		return c.Err == nil
//line /usr/local/go/src/os/exec/exec.go:643
		// _ = "end of CoverTab[107220]"
//line /usr/local/go/src/os/exec/exec.go:643
	}() && func() bool {
//line /usr/local/go/src/os/exec/exec.go:643
		_go_fuzz_dep_.CoverTab[107221]++
//line /usr/local/go/src/os/exec/exec.go:643
		return c.lookPathErr == nil
//line /usr/local/go/src/os/exec/exec.go:643
		// _ = "end of CoverTab[107221]"
//line /usr/local/go/src/os/exec/exec.go:643
	}() {
//line /usr/local/go/src/os/exec/exec.go:643
		_go_fuzz_dep_.CoverTab[107222]++
							c.Err = errors.New("exec: no command")
//line /usr/local/go/src/os/exec/exec.go:644
		// _ = "end of CoverTab[107222]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:645
		_go_fuzz_dep_.CoverTab[107223]++
//line /usr/local/go/src/os/exec/exec.go:645
		// _ = "end of CoverTab[107223]"
//line /usr/local/go/src/os/exec/exec.go:645
	}
//line /usr/local/go/src/os/exec/exec.go:645
	// _ = "end of CoverTab[107202]"
//line /usr/local/go/src/os/exec/exec.go:645
	_go_fuzz_dep_.CoverTab[107203]++
						if c.Err != nil || func() bool {
//line /usr/local/go/src/os/exec/exec.go:646
		_go_fuzz_dep_.CoverTab[107224]++
//line /usr/local/go/src/os/exec/exec.go:646
		return c.lookPathErr != nil
//line /usr/local/go/src/os/exec/exec.go:646
		// _ = "end of CoverTab[107224]"
//line /usr/local/go/src/os/exec/exec.go:646
	}() {
//line /usr/local/go/src/os/exec/exec.go:646
		_go_fuzz_dep_.CoverTab[107225]++
							if c.lookPathErr != nil {
//line /usr/local/go/src/os/exec/exec.go:647
			_go_fuzz_dep_.CoverTab[107227]++
								return c.lookPathErr
//line /usr/local/go/src/os/exec/exec.go:648
			// _ = "end of CoverTab[107227]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:649
			_go_fuzz_dep_.CoverTab[107228]++
//line /usr/local/go/src/os/exec/exec.go:649
			// _ = "end of CoverTab[107228]"
//line /usr/local/go/src/os/exec/exec.go:649
		}
//line /usr/local/go/src/os/exec/exec.go:649
		// _ = "end of CoverTab[107225]"
//line /usr/local/go/src/os/exec/exec.go:649
		_go_fuzz_dep_.CoverTab[107226]++
							return c.Err
//line /usr/local/go/src/os/exec/exec.go:650
		// _ = "end of CoverTab[107226]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:651
		_go_fuzz_dep_.CoverTab[107229]++
//line /usr/local/go/src/os/exec/exec.go:651
		// _ = "end of CoverTab[107229]"
//line /usr/local/go/src/os/exec/exec.go:651
	}
//line /usr/local/go/src/os/exec/exec.go:651
	// _ = "end of CoverTab[107203]"
//line /usr/local/go/src/os/exec/exec.go:651
	_go_fuzz_dep_.CoverTab[107204]++
						if runtime.GOOS == "windows" {
//line /usr/local/go/src/os/exec/exec.go:652
		_go_fuzz_dep_.CoverTab[107230]++
							lp, err := lookExtensions(c.Path, c.Dir)
							if err != nil {
//line /usr/local/go/src/os/exec/exec.go:654
			_go_fuzz_dep_.CoverTab[107232]++
								return err
//line /usr/local/go/src/os/exec/exec.go:655
			// _ = "end of CoverTab[107232]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:656
			_go_fuzz_dep_.CoverTab[107233]++
//line /usr/local/go/src/os/exec/exec.go:656
			// _ = "end of CoverTab[107233]"
//line /usr/local/go/src/os/exec/exec.go:656
		}
//line /usr/local/go/src/os/exec/exec.go:656
		// _ = "end of CoverTab[107230]"
//line /usr/local/go/src/os/exec/exec.go:656
		_go_fuzz_dep_.CoverTab[107231]++
							c.Path = lp
//line /usr/local/go/src/os/exec/exec.go:657
		// _ = "end of CoverTab[107231]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:658
		_go_fuzz_dep_.CoverTab[107234]++
//line /usr/local/go/src/os/exec/exec.go:658
		// _ = "end of CoverTab[107234]"
//line /usr/local/go/src/os/exec/exec.go:658
	}
//line /usr/local/go/src/os/exec/exec.go:658
	// _ = "end of CoverTab[107204]"
//line /usr/local/go/src/os/exec/exec.go:658
	_go_fuzz_dep_.CoverTab[107205]++
						if c.Cancel != nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:659
		_go_fuzz_dep_.CoverTab[107235]++
//line /usr/local/go/src/os/exec/exec.go:659
		return c.ctx == nil
//line /usr/local/go/src/os/exec/exec.go:659
		// _ = "end of CoverTab[107235]"
//line /usr/local/go/src/os/exec/exec.go:659
	}() {
//line /usr/local/go/src/os/exec/exec.go:659
		_go_fuzz_dep_.CoverTab[107236]++
							return errors.New("exec: command with a non-nil Cancel was not created with CommandContext")
//line /usr/local/go/src/os/exec/exec.go:660
		// _ = "end of CoverTab[107236]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:661
		_go_fuzz_dep_.CoverTab[107237]++
//line /usr/local/go/src/os/exec/exec.go:661
		// _ = "end of CoverTab[107237]"
//line /usr/local/go/src/os/exec/exec.go:661
	}
//line /usr/local/go/src/os/exec/exec.go:661
	// _ = "end of CoverTab[107205]"
//line /usr/local/go/src/os/exec/exec.go:661
	_go_fuzz_dep_.CoverTab[107206]++
						if c.ctx != nil {
//line /usr/local/go/src/os/exec/exec.go:662
		_go_fuzz_dep_.CoverTab[107238]++
							select {
		case <-c.ctx.Done():
//line /usr/local/go/src/os/exec/exec.go:664
			_go_fuzz_dep_.CoverTab[107239]++
								return c.ctx.Err()
//line /usr/local/go/src/os/exec/exec.go:665
			// _ = "end of CoverTab[107239]"
		default:
//line /usr/local/go/src/os/exec/exec.go:666
			_go_fuzz_dep_.CoverTab[107240]++
//line /usr/local/go/src/os/exec/exec.go:666
			// _ = "end of CoverTab[107240]"
		}
//line /usr/local/go/src/os/exec/exec.go:667
		// _ = "end of CoverTab[107238]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:668
		_go_fuzz_dep_.CoverTab[107241]++
//line /usr/local/go/src/os/exec/exec.go:668
		// _ = "end of CoverTab[107241]"
//line /usr/local/go/src/os/exec/exec.go:668
	}
//line /usr/local/go/src/os/exec/exec.go:668
	// _ = "end of CoverTab[107206]"
//line /usr/local/go/src/os/exec/exec.go:668
	_go_fuzz_dep_.CoverTab[107207]++

						childFiles := make([]*os.File, 0, 3+len(c.ExtraFiles))
						stdin, err := c.childStdin()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:672
		_go_fuzz_dep_.CoverTab[107242]++
							return err
//line /usr/local/go/src/os/exec/exec.go:673
		// _ = "end of CoverTab[107242]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:674
		_go_fuzz_dep_.CoverTab[107243]++
//line /usr/local/go/src/os/exec/exec.go:674
		// _ = "end of CoverTab[107243]"
//line /usr/local/go/src/os/exec/exec.go:674
	}
//line /usr/local/go/src/os/exec/exec.go:674
	// _ = "end of CoverTab[107207]"
//line /usr/local/go/src/os/exec/exec.go:674
	_go_fuzz_dep_.CoverTab[107208]++
						childFiles = append(childFiles, stdin)
						stdout, err := c.childStdout()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:677
		_go_fuzz_dep_.CoverTab[107244]++
							return err
//line /usr/local/go/src/os/exec/exec.go:678
		// _ = "end of CoverTab[107244]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:679
		_go_fuzz_dep_.CoverTab[107245]++
//line /usr/local/go/src/os/exec/exec.go:679
		// _ = "end of CoverTab[107245]"
//line /usr/local/go/src/os/exec/exec.go:679
	}
//line /usr/local/go/src/os/exec/exec.go:679
	// _ = "end of CoverTab[107208]"
//line /usr/local/go/src/os/exec/exec.go:679
	_go_fuzz_dep_.CoverTab[107209]++
						childFiles = append(childFiles, stdout)
						stderr, err := c.childStderr(stdout)
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:682
		_go_fuzz_dep_.CoverTab[107246]++
							return err
//line /usr/local/go/src/os/exec/exec.go:683
		// _ = "end of CoverTab[107246]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:684
		_go_fuzz_dep_.CoverTab[107247]++
//line /usr/local/go/src/os/exec/exec.go:684
		// _ = "end of CoverTab[107247]"
//line /usr/local/go/src/os/exec/exec.go:684
	}
//line /usr/local/go/src/os/exec/exec.go:684
	// _ = "end of CoverTab[107209]"
//line /usr/local/go/src/os/exec/exec.go:684
	_go_fuzz_dep_.CoverTab[107210]++
						childFiles = append(childFiles, stderr)
						childFiles = append(childFiles, c.ExtraFiles...)

						env, err := c.environ()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:689
		_go_fuzz_dep_.CoverTab[107248]++
							return err
//line /usr/local/go/src/os/exec/exec.go:690
		// _ = "end of CoverTab[107248]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:691
		_go_fuzz_dep_.CoverTab[107249]++
//line /usr/local/go/src/os/exec/exec.go:691
		// _ = "end of CoverTab[107249]"
//line /usr/local/go/src/os/exec/exec.go:691
	}
//line /usr/local/go/src/os/exec/exec.go:691
	// _ = "end of CoverTab[107210]"
//line /usr/local/go/src/os/exec/exec.go:691
	_go_fuzz_dep_.CoverTab[107211]++

						c.Process, err = os.StartProcess(c.Path, c.argv(), &os.ProcAttr{
		Dir:	c.Dir,
		Files:	childFiles,
		Env:	env,
		Sys:	c.SysProcAttr,
	})
	if err != nil {
//line /usr/local/go/src/os/exec/exec.go:699
		_go_fuzz_dep_.CoverTab[107250]++
							return err
//line /usr/local/go/src/os/exec/exec.go:700
		// _ = "end of CoverTab[107250]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:701
		_go_fuzz_dep_.CoverTab[107251]++
//line /usr/local/go/src/os/exec/exec.go:701
		// _ = "end of CoverTab[107251]"
//line /usr/local/go/src/os/exec/exec.go:701
	}
//line /usr/local/go/src/os/exec/exec.go:701
	// _ = "end of CoverTab[107211]"
//line /usr/local/go/src/os/exec/exec.go:701
	_go_fuzz_dep_.CoverTab[107212]++
						started = true

//line /usr/local/go/src/os/exec/exec.go:705
	if len(c.goroutine) > 0 {
//line /usr/local/go/src/os/exec/exec.go:705
		_go_fuzz_dep_.CoverTab[107252]++
							goroutineErr := make(chan error, 1)
							c.goroutineErr = goroutineErr

							type goroutineStatus struct {
			running		int
			firstErr	error
		}
		statusc := make(chan goroutineStatus, 1)
		statusc <- goroutineStatus{running: len(c.goroutine)}
		for _, fn := range c.goroutine {
//line /usr/local/go/src/os/exec/exec.go:715
			_go_fuzz_dep_.CoverTab[107254]++
//line /usr/local/go/src/os/exec/exec.go:715
			_curRoutineNum150_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/os/exec/exec.go:715
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum150_)
								go func(fn func() error) {
//line /usr/local/go/src/os/exec/exec.go:716
				_go_fuzz_dep_.CoverTab[107255]++
//line /usr/local/go/src/os/exec/exec.go:716
				defer func() {
//line /usr/local/go/src/os/exec/exec.go:716
					_go_fuzz_dep_.CoverTab[107257]++
//line /usr/local/go/src/os/exec/exec.go:716
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum150_)
//line /usr/local/go/src/os/exec/exec.go:716
					// _ = "end of CoverTab[107257]"
//line /usr/local/go/src/os/exec/exec.go:716
				}()
									err := fn()

									status := <-statusc
									if status.firstErr == nil {
//line /usr/local/go/src/os/exec/exec.go:720
					_go_fuzz_dep_.CoverTab[107258]++
										status.firstErr = err
//line /usr/local/go/src/os/exec/exec.go:721
					// _ = "end of CoverTab[107258]"
				} else {
//line /usr/local/go/src/os/exec/exec.go:722
					_go_fuzz_dep_.CoverTab[107259]++
//line /usr/local/go/src/os/exec/exec.go:722
					// _ = "end of CoverTab[107259]"
//line /usr/local/go/src/os/exec/exec.go:722
				}
//line /usr/local/go/src/os/exec/exec.go:722
				// _ = "end of CoverTab[107255]"
//line /usr/local/go/src/os/exec/exec.go:722
				_go_fuzz_dep_.CoverTab[107256]++
									status.running--
									if status.running == 0 {
//line /usr/local/go/src/os/exec/exec.go:724
					_go_fuzz_dep_.CoverTab[107260]++
										goroutineErr <- status.firstErr
//line /usr/local/go/src/os/exec/exec.go:725
					// _ = "end of CoverTab[107260]"
				} else {
//line /usr/local/go/src/os/exec/exec.go:726
					_go_fuzz_dep_.CoverTab[107261]++
										statusc <- status
//line /usr/local/go/src/os/exec/exec.go:727
					// _ = "end of CoverTab[107261]"
				}
//line /usr/local/go/src/os/exec/exec.go:728
				// _ = "end of CoverTab[107256]"
			}(fn)
//line /usr/local/go/src/os/exec/exec.go:729
			// _ = "end of CoverTab[107254]"
		}
//line /usr/local/go/src/os/exec/exec.go:730
		// _ = "end of CoverTab[107252]"
//line /usr/local/go/src/os/exec/exec.go:730
		_go_fuzz_dep_.CoverTab[107253]++
							c.goroutine = nil
//line /usr/local/go/src/os/exec/exec.go:731
		// _ = "end of CoverTab[107253]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:732
		_go_fuzz_dep_.CoverTab[107262]++
//line /usr/local/go/src/os/exec/exec.go:732
		// _ = "end of CoverTab[107262]"
//line /usr/local/go/src/os/exec/exec.go:732
	}
//line /usr/local/go/src/os/exec/exec.go:732
	// _ = "end of CoverTab[107212]"
//line /usr/local/go/src/os/exec/exec.go:732
	_go_fuzz_dep_.CoverTab[107213]++

//line /usr/local/go/src/os/exec/exec.go:740
	if (c.Cancel != nil || func() bool {
//line /usr/local/go/src/os/exec/exec.go:740
		_go_fuzz_dep_.CoverTab[107263]++
//line /usr/local/go/src/os/exec/exec.go:740
		return c.WaitDelay != 0
//line /usr/local/go/src/os/exec/exec.go:740
		// _ = "end of CoverTab[107263]"
//line /usr/local/go/src/os/exec/exec.go:740
	}()) && func() bool {
//line /usr/local/go/src/os/exec/exec.go:740
		_go_fuzz_dep_.CoverTab[107264]++
//line /usr/local/go/src/os/exec/exec.go:740
		return c.ctx != nil
//line /usr/local/go/src/os/exec/exec.go:740
		// _ = "end of CoverTab[107264]"
//line /usr/local/go/src/os/exec/exec.go:740
	}() && func() bool {
//line /usr/local/go/src/os/exec/exec.go:740
		_go_fuzz_dep_.CoverTab[107265]++
//line /usr/local/go/src/os/exec/exec.go:740
		return c.ctx.Done() != nil
//line /usr/local/go/src/os/exec/exec.go:740
		// _ = "end of CoverTab[107265]"
//line /usr/local/go/src/os/exec/exec.go:740
	}() {
//line /usr/local/go/src/os/exec/exec.go:740
		_go_fuzz_dep_.CoverTab[107266]++
							resultc := make(chan ctxResult)
							c.ctxResult = resultc
//line /usr/local/go/src/os/exec/exec.go:742
		_curRoutineNum151_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/os/exec/exec.go:742
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum151_)
							go c.watchCtx(resultc)
//line /usr/local/go/src/os/exec/exec.go:743
		// _ = "end of CoverTab[107266]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:744
		_go_fuzz_dep_.CoverTab[107267]++
//line /usr/local/go/src/os/exec/exec.go:744
		// _ = "end of CoverTab[107267]"
//line /usr/local/go/src/os/exec/exec.go:744
	}
//line /usr/local/go/src/os/exec/exec.go:744
	// _ = "end of CoverTab[107213]"
//line /usr/local/go/src/os/exec/exec.go:744
	_go_fuzz_dep_.CoverTab[107214]++

						return nil
//line /usr/local/go/src/os/exec/exec.go:746
	// _ = "end of CoverTab[107214]"
}

// watchCtx watches c.ctx until it is able to send a result to resultc.
//line /usr/local/go/src/os/exec/exec.go:749
//
//line /usr/local/go/src/os/exec/exec.go:749
// If c.ctx is done before a result can be sent, watchCtx calls c.Cancel,
//line /usr/local/go/src/os/exec/exec.go:749
// and/or kills cmd.Process it after c.WaitDelay has elapsed.
//line /usr/local/go/src/os/exec/exec.go:749
//
//line /usr/local/go/src/os/exec/exec.go:749
// watchCtx manipulates c.goroutineErr, so its result must be received before
//line /usr/local/go/src/os/exec/exec.go:749
// c.awaitGoroutines is called.
//line /usr/local/go/src/os/exec/exec.go:756
func (c *Cmd) watchCtx(resultc chan<- ctxResult) {
//line /usr/local/go/src/os/exec/exec.go:756
	_go_fuzz_dep_.CoverTab[107268]++
						select {
	case resultc <- ctxResult{}:
//line /usr/local/go/src/os/exec/exec.go:758
		_go_fuzz_dep_.CoverTab[107275]++
							return
//line /usr/local/go/src/os/exec/exec.go:759
		// _ = "end of CoverTab[107275]"
	case <-c.ctx.Done():
//line /usr/local/go/src/os/exec/exec.go:760
		_go_fuzz_dep_.CoverTab[107276]++
//line /usr/local/go/src/os/exec/exec.go:760
		// _ = "end of CoverTab[107276]"
	}
//line /usr/local/go/src/os/exec/exec.go:761
	// _ = "end of CoverTab[107268]"
//line /usr/local/go/src/os/exec/exec.go:761
	_go_fuzz_dep_.CoverTab[107269]++

						var err error
						if c.Cancel != nil {
//line /usr/local/go/src/os/exec/exec.go:764
		_go_fuzz_dep_.CoverTab[107277]++
							if interruptErr := c.Cancel(); interruptErr == nil {
//line /usr/local/go/src/os/exec/exec.go:765
			_go_fuzz_dep_.CoverTab[107278]++

//line /usr/local/go/src/os/exec/exec.go:769
			err = c.ctx.Err()
//line /usr/local/go/src/os/exec/exec.go:769
			// _ = "end of CoverTab[107278]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:770
			_go_fuzz_dep_.CoverTab[107279]++
//line /usr/local/go/src/os/exec/exec.go:770
			if errors.Is(interruptErr, os.ErrProcessDone) {
//line /usr/local/go/src/os/exec/exec.go:770
				_go_fuzz_dep_.CoverTab[107280]++
//line /usr/local/go/src/os/exec/exec.go:770
				// _ = "end of CoverTab[107280]"

//line /usr/local/go/src/os/exec/exec.go:774
			} else {
//line /usr/local/go/src/os/exec/exec.go:774
				_go_fuzz_dep_.CoverTab[107281]++
									err = wrappedError{
					prefix:	"exec: canceling Cmd",
					err:	interruptErr,
				}
//line /usr/local/go/src/os/exec/exec.go:778
				// _ = "end of CoverTab[107281]"
			}
//line /usr/local/go/src/os/exec/exec.go:779
			// _ = "end of CoverTab[107279]"
//line /usr/local/go/src/os/exec/exec.go:779
		}
//line /usr/local/go/src/os/exec/exec.go:779
		// _ = "end of CoverTab[107277]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:780
		_go_fuzz_dep_.CoverTab[107282]++
//line /usr/local/go/src/os/exec/exec.go:780
		// _ = "end of CoverTab[107282]"
//line /usr/local/go/src/os/exec/exec.go:780
	}
//line /usr/local/go/src/os/exec/exec.go:780
	// _ = "end of CoverTab[107269]"
//line /usr/local/go/src/os/exec/exec.go:780
	_go_fuzz_dep_.CoverTab[107270]++
						if c.WaitDelay == 0 {
//line /usr/local/go/src/os/exec/exec.go:781
		_go_fuzz_dep_.CoverTab[107283]++
							resultc <- ctxResult{err: err}
							return
//line /usr/local/go/src/os/exec/exec.go:783
		// _ = "end of CoverTab[107283]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:784
		_go_fuzz_dep_.CoverTab[107284]++
//line /usr/local/go/src/os/exec/exec.go:784
		// _ = "end of CoverTab[107284]"
//line /usr/local/go/src/os/exec/exec.go:784
	}
//line /usr/local/go/src/os/exec/exec.go:784
	// _ = "end of CoverTab[107270]"
//line /usr/local/go/src/os/exec/exec.go:784
	_go_fuzz_dep_.CoverTab[107271]++

						timer := time.NewTimer(c.WaitDelay)
						select {
	case resultc <- ctxResult{err: err, timer: timer}:
//line /usr/local/go/src/os/exec/exec.go:788
		_go_fuzz_dep_.CoverTab[107285]++

//line /usr/local/go/src/os/exec/exec.go:791
		return
//line /usr/local/go/src/os/exec/exec.go:791
		// _ = "end of CoverTab[107285]"
	case <-timer.C:
//line /usr/local/go/src/os/exec/exec.go:792
		_go_fuzz_dep_.CoverTab[107286]++
//line /usr/local/go/src/os/exec/exec.go:792
		// _ = "end of CoverTab[107286]"
	}
//line /usr/local/go/src/os/exec/exec.go:793
	// _ = "end of CoverTab[107271]"
//line /usr/local/go/src/os/exec/exec.go:793
	_go_fuzz_dep_.CoverTab[107272]++

						killed := false
						if killErr := c.Process.Kill(); killErr == nil {
//line /usr/local/go/src/os/exec/exec.go:796
		_go_fuzz_dep_.CoverTab[107287]++

//line /usr/local/go/src/os/exec/exec.go:801
		killed = true
//line /usr/local/go/src/os/exec/exec.go:801
		// _ = "end of CoverTab[107287]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:802
		_go_fuzz_dep_.CoverTab[107288]++
//line /usr/local/go/src/os/exec/exec.go:802
		if !errors.Is(killErr, os.ErrProcessDone) {
//line /usr/local/go/src/os/exec/exec.go:802
			_go_fuzz_dep_.CoverTab[107289]++
								err = wrappedError{
				prefix:	"exec: killing Cmd",
				err:	killErr,
			}
//line /usr/local/go/src/os/exec/exec.go:806
			// _ = "end of CoverTab[107289]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:807
			_go_fuzz_dep_.CoverTab[107290]++
//line /usr/local/go/src/os/exec/exec.go:807
			// _ = "end of CoverTab[107290]"
//line /usr/local/go/src/os/exec/exec.go:807
		}
//line /usr/local/go/src/os/exec/exec.go:807
		// _ = "end of CoverTab[107288]"
//line /usr/local/go/src/os/exec/exec.go:807
	}
//line /usr/local/go/src/os/exec/exec.go:807
	// _ = "end of CoverTab[107272]"
//line /usr/local/go/src/os/exec/exec.go:807
	_go_fuzz_dep_.CoverTab[107273]++

						if c.goroutineErr != nil {
//line /usr/local/go/src/os/exec/exec.go:809
		_go_fuzz_dep_.CoverTab[107291]++
							select {
		case goroutineErr := <-c.goroutineErr:
//line /usr/local/go/src/os/exec/exec.go:811
			_go_fuzz_dep_.CoverTab[107293]++

//line /usr/local/go/src/os/exec/exec.go:814
			if err == nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:814
				_go_fuzz_dep_.CoverTab[107295]++
//line /usr/local/go/src/os/exec/exec.go:814
				return !killed
//line /usr/local/go/src/os/exec/exec.go:814
				// _ = "end of CoverTab[107295]"
//line /usr/local/go/src/os/exec/exec.go:814
			}() {
//line /usr/local/go/src/os/exec/exec.go:814
				_go_fuzz_dep_.CoverTab[107296]++
									err = goroutineErr
//line /usr/local/go/src/os/exec/exec.go:815
				// _ = "end of CoverTab[107296]"
			} else {
//line /usr/local/go/src/os/exec/exec.go:816
				_go_fuzz_dep_.CoverTab[107297]++
//line /usr/local/go/src/os/exec/exec.go:816
				// _ = "end of CoverTab[107297]"
//line /usr/local/go/src/os/exec/exec.go:816
			}
//line /usr/local/go/src/os/exec/exec.go:816
			// _ = "end of CoverTab[107293]"
		default:
//line /usr/local/go/src/os/exec/exec.go:817
			_go_fuzz_dep_.CoverTab[107294]++

//line /usr/local/go/src/os/exec/exec.go:827
			closeDescriptors(c.parentIOPipes)

//line /usr/local/go/src/os/exec/exec.go:830
			_ = <-c.goroutineErr
			if err == nil {
//line /usr/local/go/src/os/exec/exec.go:831
				_go_fuzz_dep_.CoverTab[107298]++
									err = ErrWaitDelay
//line /usr/local/go/src/os/exec/exec.go:832
				// _ = "end of CoverTab[107298]"
			} else {
//line /usr/local/go/src/os/exec/exec.go:833
				_go_fuzz_dep_.CoverTab[107299]++
//line /usr/local/go/src/os/exec/exec.go:833
				// _ = "end of CoverTab[107299]"
//line /usr/local/go/src/os/exec/exec.go:833
			}
//line /usr/local/go/src/os/exec/exec.go:833
			// _ = "end of CoverTab[107294]"
		}
//line /usr/local/go/src/os/exec/exec.go:834
		// _ = "end of CoverTab[107291]"
//line /usr/local/go/src/os/exec/exec.go:834
		_go_fuzz_dep_.CoverTab[107292]++

//line /usr/local/go/src/os/exec/exec.go:838
		c.goroutineErr = nil
//line /usr/local/go/src/os/exec/exec.go:838
		// _ = "end of CoverTab[107292]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:839
		_go_fuzz_dep_.CoverTab[107300]++
//line /usr/local/go/src/os/exec/exec.go:839
		// _ = "end of CoverTab[107300]"
//line /usr/local/go/src/os/exec/exec.go:839
	}
//line /usr/local/go/src/os/exec/exec.go:839
	// _ = "end of CoverTab[107273]"
//line /usr/local/go/src/os/exec/exec.go:839
	_go_fuzz_dep_.CoverTab[107274]++

						resultc <- ctxResult{err: err}
//line /usr/local/go/src/os/exec/exec.go:841
	// _ = "end of CoverTab[107274]"
}

// An ExitError reports an unsuccessful exit by a command.
type ExitError struct {
	*os.ProcessState

	// Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being
	// collected.
	//
	// If the error output is long, Stderr may contain only a prefix
	// and suffix of the output, with the middle replaced with
	// text about the number of omitted bytes.
	//
	// Stderr is provided for debugging, for inclusion in error messages.
	// Users with other needs should redirect Cmd.Stderr as needed.
	Stderr	[]byte
}

func (e *ExitError) Error() string {
//line /usr/local/go/src/os/exec/exec.go:861
	_go_fuzz_dep_.CoverTab[107301]++
						return e.ProcessState.String()
//line /usr/local/go/src/os/exec/exec.go:862
	// _ = "end of CoverTab[107301]"
}

// Wait waits for the command to exit and waits for any copying to
//line /usr/local/go/src/os/exec/exec.go:865
// stdin or copying from stdout or stderr to complete.
//line /usr/local/go/src/os/exec/exec.go:865
//
//line /usr/local/go/src/os/exec/exec.go:865
// The command must have been started by Start.
//line /usr/local/go/src/os/exec/exec.go:865
//
//line /usr/local/go/src/os/exec/exec.go:865
// The returned error is nil if the command runs, has no problems
//line /usr/local/go/src/os/exec/exec.go:865
// copying stdin, stdout, and stderr, and exits with a zero exit
//line /usr/local/go/src/os/exec/exec.go:865
// status.
//line /usr/local/go/src/os/exec/exec.go:865
//
//line /usr/local/go/src/os/exec/exec.go:865
// If the command fails to run or doesn't complete successfully, the
//line /usr/local/go/src/os/exec/exec.go:865
// error is of type *ExitError. Other error types may be
//line /usr/local/go/src/os/exec/exec.go:865
// returned for I/O problems.
//line /usr/local/go/src/os/exec/exec.go:865
//
//line /usr/local/go/src/os/exec/exec.go:865
// If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits
//line /usr/local/go/src/os/exec/exec.go:865
// for the respective I/O loop copying to or from the process to complete.
//line /usr/local/go/src/os/exec/exec.go:865
//
//line /usr/local/go/src/os/exec/exec.go:865
// Wait releases any resources associated with the Cmd.
//line /usr/local/go/src/os/exec/exec.go:882
func (c *Cmd) Wait() error {
//line /usr/local/go/src/os/exec/exec.go:882
	_go_fuzz_dep_.CoverTab[107302]++
						if c.Process == nil {
//line /usr/local/go/src/os/exec/exec.go:883
		_go_fuzz_dep_.CoverTab[107308]++
							return errors.New("exec: not started")
//line /usr/local/go/src/os/exec/exec.go:884
		// _ = "end of CoverTab[107308]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:885
		_go_fuzz_dep_.CoverTab[107309]++
//line /usr/local/go/src/os/exec/exec.go:885
		// _ = "end of CoverTab[107309]"
//line /usr/local/go/src/os/exec/exec.go:885
	}
//line /usr/local/go/src/os/exec/exec.go:885
	// _ = "end of CoverTab[107302]"
//line /usr/local/go/src/os/exec/exec.go:885
	_go_fuzz_dep_.CoverTab[107303]++
						if c.ProcessState != nil {
//line /usr/local/go/src/os/exec/exec.go:886
		_go_fuzz_dep_.CoverTab[107310]++
							return errors.New("exec: Wait was already called")
//line /usr/local/go/src/os/exec/exec.go:887
		// _ = "end of CoverTab[107310]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:888
		_go_fuzz_dep_.CoverTab[107311]++
//line /usr/local/go/src/os/exec/exec.go:888
		// _ = "end of CoverTab[107311]"
//line /usr/local/go/src/os/exec/exec.go:888
	}
//line /usr/local/go/src/os/exec/exec.go:888
	// _ = "end of CoverTab[107303]"
//line /usr/local/go/src/os/exec/exec.go:888
	_go_fuzz_dep_.CoverTab[107304]++

						state, err := c.Process.Wait()
						if err == nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:891
		_go_fuzz_dep_.CoverTab[107312]++
//line /usr/local/go/src/os/exec/exec.go:891
		return !state.Success()
//line /usr/local/go/src/os/exec/exec.go:891
		// _ = "end of CoverTab[107312]"
//line /usr/local/go/src/os/exec/exec.go:891
	}() {
//line /usr/local/go/src/os/exec/exec.go:891
		_go_fuzz_dep_.CoverTab[107313]++
							err = &ExitError{ProcessState: state}
//line /usr/local/go/src/os/exec/exec.go:892
		// _ = "end of CoverTab[107313]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:893
		_go_fuzz_dep_.CoverTab[107314]++
//line /usr/local/go/src/os/exec/exec.go:893
		// _ = "end of CoverTab[107314]"
//line /usr/local/go/src/os/exec/exec.go:893
	}
//line /usr/local/go/src/os/exec/exec.go:893
	// _ = "end of CoverTab[107304]"
//line /usr/local/go/src/os/exec/exec.go:893
	_go_fuzz_dep_.CoverTab[107305]++
						c.ProcessState = state

						var timer *time.Timer
						if c.ctxResult != nil {
//line /usr/local/go/src/os/exec/exec.go:897
		_go_fuzz_dep_.CoverTab[107315]++
							watch := <-c.ctxResult
							timer = watch.timer

//line /usr/local/go/src/os/exec/exec.go:903
		if err == nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:903
			_go_fuzz_dep_.CoverTab[107316]++
//line /usr/local/go/src/os/exec/exec.go:903
			return watch.err != nil
//line /usr/local/go/src/os/exec/exec.go:903
			// _ = "end of CoverTab[107316]"
//line /usr/local/go/src/os/exec/exec.go:903
		}() {
//line /usr/local/go/src/os/exec/exec.go:903
			_go_fuzz_dep_.CoverTab[107317]++
								err = watch.err
//line /usr/local/go/src/os/exec/exec.go:904
			// _ = "end of CoverTab[107317]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:905
			_go_fuzz_dep_.CoverTab[107318]++
//line /usr/local/go/src/os/exec/exec.go:905
			// _ = "end of CoverTab[107318]"
//line /usr/local/go/src/os/exec/exec.go:905
		}
//line /usr/local/go/src/os/exec/exec.go:905
		// _ = "end of CoverTab[107315]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:906
		_go_fuzz_dep_.CoverTab[107319]++
//line /usr/local/go/src/os/exec/exec.go:906
		// _ = "end of CoverTab[107319]"
//line /usr/local/go/src/os/exec/exec.go:906
	}
//line /usr/local/go/src/os/exec/exec.go:906
	// _ = "end of CoverTab[107305]"
//line /usr/local/go/src/os/exec/exec.go:906
	_go_fuzz_dep_.CoverTab[107306]++

						if goroutineErr := c.awaitGoroutines(timer); err == nil {
//line /usr/local/go/src/os/exec/exec.go:908
		_go_fuzz_dep_.CoverTab[107320]++

//line /usr/local/go/src/os/exec/exec.go:912
		err = goroutineErr
//line /usr/local/go/src/os/exec/exec.go:912
		// _ = "end of CoverTab[107320]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:913
		_go_fuzz_dep_.CoverTab[107321]++
//line /usr/local/go/src/os/exec/exec.go:913
		// _ = "end of CoverTab[107321]"
//line /usr/local/go/src/os/exec/exec.go:913
	}
//line /usr/local/go/src/os/exec/exec.go:913
	// _ = "end of CoverTab[107306]"
//line /usr/local/go/src/os/exec/exec.go:913
	_go_fuzz_dep_.CoverTab[107307]++
						closeDescriptors(c.parentIOPipes)
						c.parentIOPipes = nil

						return err
//line /usr/local/go/src/os/exec/exec.go:917
	// _ = "end of CoverTab[107307]"
}

// awaitGoroutines waits for the results of the goroutines copying data to or
//line /usr/local/go/src/os/exec/exec.go:920
// from the command's I/O pipes.
//line /usr/local/go/src/os/exec/exec.go:920
//
//line /usr/local/go/src/os/exec/exec.go:920
// If c.WaitDelay elapses before the goroutines complete, awaitGoroutines
//line /usr/local/go/src/os/exec/exec.go:920
// forcibly closes their pipes and returns ErrWaitDelay.
//line /usr/local/go/src/os/exec/exec.go:920
//
//line /usr/local/go/src/os/exec/exec.go:920
// If timer is non-nil, it must send to timer.C at the end of c.WaitDelay.
//line /usr/local/go/src/os/exec/exec.go:927
func (c *Cmd) awaitGoroutines(timer *time.Timer) error {
//line /usr/local/go/src/os/exec/exec.go:927
	_go_fuzz_dep_.CoverTab[107322]++
						defer func() {
//line /usr/local/go/src/os/exec/exec.go:928
		_go_fuzz_dep_.CoverTab[107326]++
							if timer != nil {
//line /usr/local/go/src/os/exec/exec.go:929
			_go_fuzz_dep_.CoverTab[107328]++
								timer.Stop()
//line /usr/local/go/src/os/exec/exec.go:930
			// _ = "end of CoverTab[107328]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:931
			_go_fuzz_dep_.CoverTab[107329]++
//line /usr/local/go/src/os/exec/exec.go:931
			// _ = "end of CoverTab[107329]"
//line /usr/local/go/src/os/exec/exec.go:931
		}
//line /usr/local/go/src/os/exec/exec.go:931
		// _ = "end of CoverTab[107326]"
//line /usr/local/go/src/os/exec/exec.go:931
		_go_fuzz_dep_.CoverTab[107327]++
							c.goroutineErr = nil
//line /usr/local/go/src/os/exec/exec.go:932
		// _ = "end of CoverTab[107327]"
	}()
//line /usr/local/go/src/os/exec/exec.go:933
	// _ = "end of CoverTab[107322]"
//line /usr/local/go/src/os/exec/exec.go:933
	_go_fuzz_dep_.CoverTab[107323]++

						if c.goroutineErr == nil {
//line /usr/local/go/src/os/exec/exec.go:935
		_go_fuzz_dep_.CoverTab[107330]++
							return nil
//line /usr/local/go/src/os/exec/exec.go:936
		// _ = "end of CoverTab[107330]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:937
		_go_fuzz_dep_.CoverTab[107331]++
//line /usr/local/go/src/os/exec/exec.go:937
		// _ = "end of CoverTab[107331]"
//line /usr/local/go/src/os/exec/exec.go:937
	}
//line /usr/local/go/src/os/exec/exec.go:937
	// _ = "end of CoverTab[107323]"
//line /usr/local/go/src/os/exec/exec.go:937
	_go_fuzz_dep_.CoverTab[107324]++

						if timer == nil {
//line /usr/local/go/src/os/exec/exec.go:939
		_go_fuzz_dep_.CoverTab[107332]++
							if c.WaitDelay == 0 {
//line /usr/local/go/src/os/exec/exec.go:940
			_go_fuzz_dep_.CoverTab[107335]++
								return <-c.goroutineErr
//line /usr/local/go/src/os/exec/exec.go:941
			// _ = "end of CoverTab[107335]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:942
			_go_fuzz_dep_.CoverTab[107336]++
//line /usr/local/go/src/os/exec/exec.go:942
			// _ = "end of CoverTab[107336]"
//line /usr/local/go/src/os/exec/exec.go:942
		}
//line /usr/local/go/src/os/exec/exec.go:942
		// _ = "end of CoverTab[107332]"
//line /usr/local/go/src/os/exec/exec.go:942
		_go_fuzz_dep_.CoverTab[107333]++

							select {
		case err := <-c.goroutineErr:
//line /usr/local/go/src/os/exec/exec.go:945
			_go_fuzz_dep_.CoverTab[107337]++

								return err
//line /usr/local/go/src/os/exec/exec.go:947
			// _ = "end of CoverTab[107337]"
		default:
//line /usr/local/go/src/os/exec/exec.go:948
			_go_fuzz_dep_.CoverTab[107338]++
//line /usr/local/go/src/os/exec/exec.go:948
			// _ = "end of CoverTab[107338]"
		}
//line /usr/local/go/src/os/exec/exec.go:949
		// _ = "end of CoverTab[107333]"
//line /usr/local/go/src/os/exec/exec.go:949
		_go_fuzz_dep_.CoverTab[107334]++

//line /usr/local/go/src/os/exec/exec.go:953
		timer = time.NewTimer(c.WaitDelay)
//line /usr/local/go/src/os/exec/exec.go:953
		// _ = "end of CoverTab[107334]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:954
		_go_fuzz_dep_.CoverTab[107339]++
//line /usr/local/go/src/os/exec/exec.go:954
		// _ = "end of CoverTab[107339]"
//line /usr/local/go/src/os/exec/exec.go:954
	}
//line /usr/local/go/src/os/exec/exec.go:954
	// _ = "end of CoverTab[107324]"
//line /usr/local/go/src/os/exec/exec.go:954
	_go_fuzz_dep_.CoverTab[107325]++

						select {
	case <-timer.C:
//line /usr/local/go/src/os/exec/exec.go:957
		_go_fuzz_dep_.CoverTab[107340]++
							closeDescriptors(c.parentIOPipes)

//line /usr/local/go/src/os/exec/exec.go:961
		_ = <-c.goroutineErr
							return ErrWaitDelay
//line /usr/local/go/src/os/exec/exec.go:962
		// _ = "end of CoverTab[107340]"

	case err := <-c.goroutineErr:
//line /usr/local/go/src/os/exec/exec.go:964
		_go_fuzz_dep_.CoverTab[107341]++
							return err
//line /usr/local/go/src/os/exec/exec.go:965
		// _ = "end of CoverTab[107341]"
	}
//line /usr/local/go/src/os/exec/exec.go:966
	// _ = "end of CoverTab[107325]"
}

// Output runs the command and returns its standard output.
//line /usr/local/go/src/os/exec/exec.go:969
// Any returned error will usually be of type *ExitError.
//line /usr/local/go/src/os/exec/exec.go:969
// If c.Stderr was nil, Output populates ExitError.Stderr.
//line /usr/local/go/src/os/exec/exec.go:972
func (c *Cmd) Output() ([]byte, error) {
//line /usr/local/go/src/os/exec/exec.go:972
	_go_fuzz_dep_.CoverTab[107342]++
						if c.Stdout != nil {
//line /usr/local/go/src/os/exec/exec.go:973
		_go_fuzz_dep_.CoverTab[107346]++
							return nil, errors.New("exec: Stdout already set")
//line /usr/local/go/src/os/exec/exec.go:974
		// _ = "end of CoverTab[107346]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:975
		_go_fuzz_dep_.CoverTab[107347]++
//line /usr/local/go/src/os/exec/exec.go:975
		// _ = "end of CoverTab[107347]"
//line /usr/local/go/src/os/exec/exec.go:975
	}
//line /usr/local/go/src/os/exec/exec.go:975
	// _ = "end of CoverTab[107342]"
//line /usr/local/go/src/os/exec/exec.go:975
	_go_fuzz_dep_.CoverTab[107343]++
						var stdout bytes.Buffer
						c.Stdout = &stdout

						captureErr := c.Stderr == nil
						if captureErr {
//line /usr/local/go/src/os/exec/exec.go:980
		_go_fuzz_dep_.CoverTab[107348]++
							c.Stderr = &prefixSuffixSaver{N: 32 << 10}
//line /usr/local/go/src/os/exec/exec.go:981
		// _ = "end of CoverTab[107348]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:982
		_go_fuzz_dep_.CoverTab[107349]++
//line /usr/local/go/src/os/exec/exec.go:982
		// _ = "end of CoverTab[107349]"
//line /usr/local/go/src/os/exec/exec.go:982
	}
//line /usr/local/go/src/os/exec/exec.go:982
	// _ = "end of CoverTab[107343]"
//line /usr/local/go/src/os/exec/exec.go:982
	_go_fuzz_dep_.CoverTab[107344]++

						err := c.Run()
						if err != nil && func() bool {
//line /usr/local/go/src/os/exec/exec.go:985
		_go_fuzz_dep_.CoverTab[107350]++
//line /usr/local/go/src/os/exec/exec.go:985
		return captureErr
//line /usr/local/go/src/os/exec/exec.go:985
		// _ = "end of CoverTab[107350]"
//line /usr/local/go/src/os/exec/exec.go:985
	}() {
//line /usr/local/go/src/os/exec/exec.go:985
		_go_fuzz_dep_.CoverTab[107351]++
							if ee, ok := err.(*ExitError); ok {
//line /usr/local/go/src/os/exec/exec.go:986
			_go_fuzz_dep_.CoverTab[107352]++
								ee.Stderr = c.Stderr.(*prefixSuffixSaver).Bytes()
//line /usr/local/go/src/os/exec/exec.go:987
			// _ = "end of CoverTab[107352]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:988
			_go_fuzz_dep_.CoverTab[107353]++
//line /usr/local/go/src/os/exec/exec.go:988
			// _ = "end of CoverTab[107353]"
//line /usr/local/go/src/os/exec/exec.go:988
		}
//line /usr/local/go/src/os/exec/exec.go:988
		// _ = "end of CoverTab[107351]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:989
		_go_fuzz_dep_.CoverTab[107354]++
//line /usr/local/go/src/os/exec/exec.go:989
		// _ = "end of CoverTab[107354]"
//line /usr/local/go/src/os/exec/exec.go:989
	}
//line /usr/local/go/src/os/exec/exec.go:989
	// _ = "end of CoverTab[107344]"
//line /usr/local/go/src/os/exec/exec.go:989
	_go_fuzz_dep_.CoverTab[107345]++
						return stdout.Bytes(), err
//line /usr/local/go/src/os/exec/exec.go:990
	// _ = "end of CoverTab[107345]"
}

// CombinedOutput runs the command and returns its combined standard
//line /usr/local/go/src/os/exec/exec.go:993
// output and standard error.
//line /usr/local/go/src/os/exec/exec.go:995
func (c *Cmd) CombinedOutput() ([]byte, error) {
//line /usr/local/go/src/os/exec/exec.go:995
	_go_fuzz_dep_.CoverTab[107355]++
						if c.Stdout != nil {
//line /usr/local/go/src/os/exec/exec.go:996
		_go_fuzz_dep_.CoverTab[107358]++
							return nil, errors.New("exec: Stdout already set")
//line /usr/local/go/src/os/exec/exec.go:997
		// _ = "end of CoverTab[107358]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:998
		_go_fuzz_dep_.CoverTab[107359]++
//line /usr/local/go/src/os/exec/exec.go:998
		// _ = "end of CoverTab[107359]"
//line /usr/local/go/src/os/exec/exec.go:998
	}
//line /usr/local/go/src/os/exec/exec.go:998
	// _ = "end of CoverTab[107355]"
//line /usr/local/go/src/os/exec/exec.go:998
	_go_fuzz_dep_.CoverTab[107356]++
						if c.Stderr != nil {
//line /usr/local/go/src/os/exec/exec.go:999
		_go_fuzz_dep_.CoverTab[107360]++
							return nil, errors.New("exec: Stderr already set")
//line /usr/local/go/src/os/exec/exec.go:1000
		// _ = "end of CoverTab[107360]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1001
		_go_fuzz_dep_.CoverTab[107361]++
//line /usr/local/go/src/os/exec/exec.go:1001
		// _ = "end of CoverTab[107361]"
//line /usr/local/go/src/os/exec/exec.go:1001
	}
//line /usr/local/go/src/os/exec/exec.go:1001
	// _ = "end of CoverTab[107356]"
//line /usr/local/go/src/os/exec/exec.go:1001
	_go_fuzz_dep_.CoverTab[107357]++
						var b bytes.Buffer
						c.Stdout = &b
						c.Stderr = &b
						err := c.Run()
						return b.Bytes(), err
//line /usr/local/go/src/os/exec/exec.go:1006
	// _ = "end of CoverTab[107357]"
}

// StdinPipe returns a pipe that will be connected to the command's
//line /usr/local/go/src/os/exec/exec.go:1009
// standard input when the command starts.
//line /usr/local/go/src/os/exec/exec.go:1009
// The pipe will be closed automatically after Wait sees the command exit.
//line /usr/local/go/src/os/exec/exec.go:1009
// A caller need only call Close to force the pipe to close sooner.
//line /usr/local/go/src/os/exec/exec.go:1009
// For example, if the command being run will not exit until standard input
//line /usr/local/go/src/os/exec/exec.go:1009
// is closed, the caller must close the pipe.
//line /usr/local/go/src/os/exec/exec.go:1015
func (c *Cmd) StdinPipe() (io.WriteCloser, error) {
//line /usr/local/go/src/os/exec/exec.go:1015
	_go_fuzz_dep_.CoverTab[107362]++
						if c.Stdin != nil {
//line /usr/local/go/src/os/exec/exec.go:1016
		_go_fuzz_dep_.CoverTab[107366]++
							return nil, errors.New("exec: Stdin already set")
//line /usr/local/go/src/os/exec/exec.go:1017
		// _ = "end of CoverTab[107366]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1018
		_go_fuzz_dep_.CoverTab[107367]++
//line /usr/local/go/src/os/exec/exec.go:1018
		// _ = "end of CoverTab[107367]"
//line /usr/local/go/src/os/exec/exec.go:1018
	}
//line /usr/local/go/src/os/exec/exec.go:1018
	// _ = "end of CoverTab[107362]"
//line /usr/local/go/src/os/exec/exec.go:1018
	_go_fuzz_dep_.CoverTab[107363]++
						if c.Process != nil {
//line /usr/local/go/src/os/exec/exec.go:1019
		_go_fuzz_dep_.CoverTab[107368]++
							return nil, errors.New("exec: StdinPipe after process started")
//line /usr/local/go/src/os/exec/exec.go:1020
		// _ = "end of CoverTab[107368]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1021
		_go_fuzz_dep_.CoverTab[107369]++
//line /usr/local/go/src/os/exec/exec.go:1021
		// _ = "end of CoverTab[107369]"
//line /usr/local/go/src/os/exec/exec.go:1021
	}
//line /usr/local/go/src/os/exec/exec.go:1021
	// _ = "end of CoverTab[107363]"
//line /usr/local/go/src/os/exec/exec.go:1021
	_go_fuzz_dep_.CoverTab[107364]++
						pr, pw, err := os.Pipe()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:1023
		_go_fuzz_dep_.CoverTab[107370]++
							return nil, err
//line /usr/local/go/src/os/exec/exec.go:1024
		// _ = "end of CoverTab[107370]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1025
		_go_fuzz_dep_.CoverTab[107371]++
//line /usr/local/go/src/os/exec/exec.go:1025
		// _ = "end of CoverTab[107371]"
//line /usr/local/go/src/os/exec/exec.go:1025
	}
//line /usr/local/go/src/os/exec/exec.go:1025
	// _ = "end of CoverTab[107364]"
//line /usr/local/go/src/os/exec/exec.go:1025
	_go_fuzz_dep_.CoverTab[107365]++
						c.Stdin = pr
						c.childIOFiles = append(c.childIOFiles, pr)
						c.parentIOPipes = append(c.parentIOPipes, pw)
						return pw, nil
//line /usr/local/go/src/os/exec/exec.go:1029
	// _ = "end of CoverTab[107365]"
}

// StdoutPipe returns a pipe that will be connected to the command's
//line /usr/local/go/src/os/exec/exec.go:1032
// standard output when the command starts.
//line /usr/local/go/src/os/exec/exec.go:1032
//
//line /usr/local/go/src/os/exec/exec.go:1032
// Wait will close the pipe after seeing the command exit, so most callers
//line /usr/local/go/src/os/exec/exec.go:1032
// need not close the pipe themselves. It is thus incorrect to call Wait
//line /usr/local/go/src/os/exec/exec.go:1032
// before all reads from the pipe have completed.
//line /usr/local/go/src/os/exec/exec.go:1032
// For the same reason, it is incorrect to call Run when using StdoutPipe.
//line /usr/local/go/src/os/exec/exec.go:1032
// See the example for idiomatic usage.
//line /usr/local/go/src/os/exec/exec.go:1040
func (c *Cmd) StdoutPipe() (io.ReadCloser, error) {
//line /usr/local/go/src/os/exec/exec.go:1040
	_go_fuzz_dep_.CoverTab[107372]++
						if c.Stdout != nil {
//line /usr/local/go/src/os/exec/exec.go:1041
		_go_fuzz_dep_.CoverTab[107376]++
							return nil, errors.New("exec: Stdout already set")
//line /usr/local/go/src/os/exec/exec.go:1042
		// _ = "end of CoverTab[107376]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1043
		_go_fuzz_dep_.CoverTab[107377]++
//line /usr/local/go/src/os/exec/exec.go:1043
		// _ = "end of CoverTab[107377]"
//line /usr/local/go/src/os/exec/exec.go:1043
	}
//line /usr/local/go/src/os/exec/exec.go:1043
	// _ = "end of CoverTab[107372]"
//line /usr/local/go/src/os/exec/exec.go:1043
	_go_fuzz_dep_.CoverTab[107373]++
						if c.Process != nil {
//line /usr/local/go/src/os/exec/exec.go:1044
		_go_fuzz_dep_.CoverTab[107378]++
							return nil, errors.New("exec: StdoutPipe after process started")
//line /usr/local/go/src/os/exec/exec.go:1045
		// _ = "end of CoverTab[107378]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1046
		_go_fuzz_dep_.CoverTab[107379]++
//line /usr/local/go/src/os/exec/exec.go:1046
		// _ = "end of CoverTab[107379]"
//line /usr/local/go/src/os/exec/exec.go:1046
	}
//line /usr/local/go/src/os/exec/exec.go:1046
	// _ = "end of CoverTab[107373]"
//line /usr/local/go/src/os/exec/exec.go:1046
	_go_fuzz_dep_.CoverTab[107374]++
						pr, pw, err := os.Pipe()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:1048
		_go_fuzz_dep_.CoverTab[107380]++
							return nil, err
//line /usr/local/go/src/os/exec/exec.go:1049
		// _ = "end of CoverTab[107380]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1050
		_go_fuzz_dep_.CoverTab[107381]++
//line /usr/local/go/src/os/exec/exec.go:1050
		// _ = "end of CoverTab[107381]"
//line /usr/local/go/src/os/exec/exec.go:1050
	}
//line /usr/local/go/src/os/exec/exec.go:1050
	// _ = "end of CoverTab[107374]"
//line /usr/local/go/src/os/exec/exec.go:1050
	_go_fuzz_dep_.CoverTab[107375]++
						c.Stdout = pw
						c.childIOFiles = append(c.childIOFiles, pw)
						c.parentIOPipes = append(c.parentIOPipes, pr)
						return pr, nil
//line /usr/local/go/src/os/exec/exec.go:1054
	// _ = "end of CoverTab[107375]"
}

// StderrPipe returns a pipe that will be connected to the command's
//line /usr/local/go/src/os/exec/exec.go:1057
// standard error when the command starts.
//line /usr/local/go/src/os/exec/exec.go:1057
//
//line /usr/local/go/src/os/exec/exec.go:1057
// Wait will close the pipe after seeing the command exit, so most callers
//line /usr/local/go/src/os/exec/exec.go:1057
// need not close the pipe themselves. It is thus incorrect to call Wait
//line /usr/local/go/src/os/exec/exec.go:1057
// before all reads from the pipe have completed.
//line /usr/local/go/src/os/exec/exec.go:1057
// For the same reason, it is incorrect to use Run when using StderrPipe.
//line /usr/local/go/src/os/exec/exec.go:1057
// See the StdoutPipe example for idiomatic usage.
//line /usr/local/go/src/os/exec/exec.go:1065
func (c *Cmd) StderrPipe() (io.ReadCloser, error) {
//line /usr/local/go/src/os/exec/exec.go:1065
	_go_fuzz_dep_.CoverTab[107382]++
						if c.Stderr != nil {
//line /usr/local/go/src/os/exec/exec.go:1066
		_go_fuzz_dep_.CoverTab[107386]++
							return nil, errors.New("exec: Stderr already set")
//line /usr/local/go/src/os/exec/exec.go:1067
		// _ = "end of CoverTab[107386]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1068
		_go_fuzz_dep_.CoverTab[107387]++
//line /usr/local/go/src/os/exec/exec.go:1068
		// _ = "end of CoverTab[107387]"
//line /usr/local/go/src/os/exec/exec.go:1068
	}
//line /usr/local/go/src/os/exec/exec.go:1068
	// _ = "end of CoverTab[107382]"
//line /usr/local/go/src/os/exec/exec.go:1068
	_go_fuzz_dep_.CoverTab[107383]++
						if c.Process != nil {
//line /usr/local/go/src/os/exec/exec.go:1069
		_go_fuzz_dep_.CoverTab[107388]++
							return nil, errors.New("exec: StderrPipe after process started")
//line /usr/local/go/src/os/exec/exec.go:1070
		// _ = "end of CoverTab[107388]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1071
		_go_fuzz_dep_.CoverTab[107389]++
//line /usr/local/go/src/os/exec/exec.go:1071
		// _ = "end of CoverTab[107389]"
//line /usr/local/go/src/os/exec/exec.go:1071
	}
//line /usr/local/go/src/os/exec/exec.go:1071
	// _ = "end of CoverTab[107383]"
//line /usr/local/go/src/os/exec/exec.go:1071
	_go_fuzz_dep_.CoverTab[107384]++
						pr, pw, err := os.Pipe()
						if err != nil {
//line /usr/local/go/src/os/exec/exec.go:1073
		_go_fuzz_dep_.CoverTab[107390]++
							return nil, err
//line /usr/local/go/src/os/exec/exec.go:1074
		// _ = "end of CoverTab[107390]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1075
		_go_fuzz_dep_.CoverTab[107391]++
//line /usr/local/go/src/os/exec/exec.go:1075
		// _ = "end of CoverTab[107391]"
//line /usr/local/go/src/os/exec/exec.go:1075
	}
//line /usr/local/go/src/os/exec/exec.go:1075
	// _ = "end of CoverTab[107384]"
//line /usr/local/go/src/os/exec/exec.go:1075
	_go_fuzz_dep_.CoverTab[107385]++
						c.Stderr = pw
						c.childIOFiles = append(c.childIOFiles, pw)
						c.parentIOPipes = append(c.parentIOPipes, pr)
						return pr, nil
//line /usr/local/go/src/os/exec/exec.go:1079
	// _ = "end of CoverTab[107385]"
}

// prefixSuffixSaver is an io.Writer which retains the first N bytes
//line /usr/local/go/src/os/exec/exec.go:1082
// and the last N bytes written to it. The Bytes() methods reconstructs
//line /usr/local/go/src/os/exec/exec.go:1082
// it with a pretty error message.
//line /usr/local/go/src/os/exec/exec.go:1085
type prefixSuffixSaver struct {
	N		int	// max size of prefix or suffix
	prefix		[]byte
	suffix		[]byte	// ring buffer once len(suffix) == N
	suffixOff	int	// offset to write into suffix
	skipped		int64
//line /usr/local/go/src/os/exec/exec.go:1097
}

func (w *prefixSuffixSaver) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/os/exec/exec.go:1099
	_go_fuzz_dep_.CoverTab[107392]++
						lenp := len(p)
						p = w.fill(&w.prefix, p)

//line /usr/local/go/src/os/exec/exec.go:1104
	if overage := len(p) - w.N; overage > 0 {
//line /usr/local/go/src/os/exec/exec.go:1104
		_go_fuzz_dep_.CoverTab[107395]++
							p = p[overage:]
							w.skipped += int64(overage)
//line /usr/local/go/src/os/exec/exec.go:1106
		// _ = "end of CoverTab[107395]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1107
		_go_fuzz_dep_.CoverTab[107396]++
//line /usr/local/go/src/os/exec/exec.go:1107
		// _ = "end of CoverTab[107396]"
//line /usr/local/go/src/os/exec/exec.go:1107
	}
//line /usr/local/go/src/os/exec/exec.go:1107
	// _ = "end of CoverTab[107392]"
//line /usr/local/go/src/os/exec/exec.go:1107
	_go_fuzz_dep_.CoverTab[107393]++
						p = w.fill(&w.suffix, p)

//line /usr/local/go/src/os/exec/exec.go:1111
	for len(p) > 0 {
//line /usr/local/go/src/os/exec/exec.go:1111
		_go_fuzz_dep_.CoverTab[107397]++
							n := copy(w.suffix[w.suffixOff:], p)
							p = p[n:]
							w.skipped += int64(n)
							w.suffixOff += n
							if w.suffixOff == w.N {
//line /usr/local/go/src/os/exec/exec.go:1116
			_go_fuzz_dep_.CoverTab[107398]++
								w.suffixOff = 0
//line /usr/local/go/src/os/exec/exec.go:1117
			// _ = "end of CoverTab[107398]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1118
			_go_fuzz_dep_.CoverTab[107399]++
//line /usr/local/go/src/os/exec/exec.go:1118
			// _ = "end of CoverTab[107399]"
//line /usr/local/go/src/os/exec/exec.go:1118
		}
//line /usr/local/go/src/os/exec/exec.go:1118
		// _ = "end of CoverTab[107397]"
	}
//line /usr/local/go/src/os/exec/exec.go:1119
	// _ = "end of CoverTab[107393]"
//line /usr/local/go/src/os/exec/exec.go:1119
	_go_fuzz_dep_.CoverTab[107394]++
						return lenp, nil
//line /usr/local/go/src/os/exec/exec.go:1120
	// _ = "end of CoverTab[107394]"
}

// fill appends up to len(p) bytes of p to *dst, such that *dst does not
//line /usr/local/go/src/os/exec/exec.go:1123
// grow larger than w.N. It returns the un-appended suffix of p.
//line /usr/local/go/src/os/exec/exec.go:1125
func (w *prefixSuffixSaver) fill(dst *[]byte, p []byte) (pRemain []byte) {
//line /usr/local/go/src/os/exec/exec.go:1125
	_go_fuzz_dep_.CoverTab[107400]++
						if remain := w.N - len(*dst); remain > 0 {
//line /usr/local/go/src/os/exec/exec.go:1126
		_go_fuzz_dep_.CoverTab[107402]++
							add := minInt(len(p), remain)
							*dst = append(*dst, p[:add]...)
							p = p[add:]
//line /usr/local/go/src/os/exec/exec.go:1129
		// _ = "end of CoverTab[107402]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1130
		_go_fuzz_dep_.CoverTab[107403]++
//line /usr/local/go/src/os/exec/exec.go:1130
		// _ = "end of CoverTab[107403]"
//line /usr/local/go/src/os/exec/exec.go:1130
	}
//line /usr/local/go/src/os/exec/exec.go:1130
	// _ = "end of CoverTab[107400]"
//line /usr/local/go/src/os/exec/exec.go:1130
	_go_fuzz_dep_.CoverTab[107401]++
						return p
//line /usr/local/go/src/os/exec/exec.go:1131
	// _ = "end of CoverTab[107401]"
}

func (w *prefixSuffixSaver) Bytes() []byte {
//line /usr/local/go/src/os/exec/exec.go:1134
	_go_fuzz_dep_.CoverTab[107404]++
						if w.suffix == nil {
//line /usr/local/go/src/os/exec/exec.go:1135
		_go_fuzz_dep_.CoverTab[107407]++
							return w.prefix
//line /usr/local/go/src/os/exec/exec.go:1136
		// _ = "end of CoverTab[107407]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1137
		_go_fuzz_dep_.CoverTab[107408]++
//line /usr/local/go/src/os/exec/exec.go:1137
		// _ = "end of CoverTab[107408]"
//line /usr/local/go/src/os/exec/exec.go:1137
	}
//line /usr/local/go/src/os/exec/exec.go:1137
	// _ = "end of CoverTab[107404]"
//line /usr/local/go/src/os/exec/exec.go:1137
	_go_fuzz_dep_.CoverTab[107405]++
						if w.skipped == 0 {
//line /usr/local/go/src/os/exec/exec.go:1138
		_go_fuzz_dep_.CoverTab[107409]++
							return append(w.prefix, w.suffix...)
//line /usr/local/go/src/os/exec/exec.go:1139
		// _ = "end of CoverTab[107409]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1140
		_go_fuzz_dep_.CoverTab[107410]++
//line /usr/local/go/src/os/exec/exec.go:1140
		// _ = "end of CoverTab[107410]"
//line /usr/local/go/src/os/exec/exec.go:1140
	}
//line /usr/local/go/src/os/exec/exec.go:1140
	// _ = "end of CoverTab[107405]"
//line /usr/local/go/src/os/exec/exec.go:1140
	_go_fuzz_dep_.CoverTab[107406]++
						var buf bytes.Buffer
						buf.Grow(len(w.prefix) + len(w.suffix) + 50)
						buf.Write(w.prefix)
						buf.WriteString("\n... omitting ")
						buf.WriteString(strconv.FormatInt(w.skipped, 10))
						buf.WriteString(" bytes ...\n")
						buf.Write(w.suffix[w.suffixOff:])
						buf.Write(w.suffix[:w.suffixOff])
						return buf.Bytes()
//line /usr/local/go/src/os/exec/exec.go:1149
	// _ = "end of CoverTab[107406]"
}

func minInt(a, b int) int {
//line /usr/local/go/src/os/exec/exec.go:1152
	_go_fuzz_dep_.CoverTab[107411]++
						if a < b {
//line /usr/local/go/src/os/exec/exec.go:1153
		_go_fuzz_dep_.CoverTab[107413]++
							return a
//line /usr/local/go/src/os/exec/exec.go:1154
		// _ = "end of CoverTab[107413]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1155
		_go_fuzz_dep_.CoverTab[107414]++
//line /usr/local/go/src/os/exec/exec.go:1155
		// _ = "end of CoverTab[107414]"
//line /usr/local/go/src/os/exec/exec.go:1155
	}
//line /usr/local/go/src/os/exec/exec.go:1155
	// _ = "end of CoverTab[107411]"
//line /usr/local/go/src/os/exec/exec.go:1155
	_go_fuzz_dep_.CoverTab[107412]++
						return b
//line /usr/local/go/src/os/exec/exec.go:1156
	// _ = "end of CoverTab[107412]"
}

// environ returns a best-effort copy of the environment in which the command
//line /usr/local/go/src/os/exec/exec.go:1159
// would be run as it is currently configured. If an error occurs in computing
//line /usr/local/go/src/os/exec/exec.go:1159
// the environment, it is returned alongside the best-effort copy.
//line /usr/local/go/src/os/exec/exec.go:1162
func (c *Cmd) environ() ([]string, error) {
//line /usr/local/go/src/os/exec/exec.go:1162
	_go_fuzz_dep_.CoverTab[107415]++
						var err error

						env := c.Env
						if env == nil {
//line /usr/local/go/src/os/exec/exec.go:1166
		_go_fuzz_dep_.CoverTab[107418]++
							env, err = execenv.Default(c.SysProcAttr)
							if err != nil {
//line /usr/local/go/src/os/exec/exec.go:1168
			_go_fuzz_dep_.CoverTab[107420]++
								env = os.Environ()
//line /usr/local/go/src/os/exec/exec.go:1169
			// _ = "end of CoverTab[107420]"

		} else {
//line /usr/local/go/src/os/exec/exec.go:1171
			_go_fuzz_dep_.CoverTab[107421]++
//line /usr/local/go/src/os/exec/exec.go:1171
			// _ = "end of CoverTab[107421]"
//line /usr/local/go/src/os/exec/exec.go:1171
		}
//line /usr/local/go/src/os/exec/exec.go:1171
		// _ = "end of CoverTab[107418]"
//line /usr/local/go/src/os/exec/exec.go:1171
		_go_fuzz_dep_.CoverTab[107419]++

							if c.Dir != "" {
//line /usr/local/go/src/os/exec/exec.go:1173
			_go_fuzz_dep_.CoverTab[107422]++
								switch runtime.GOOS {
			case "windows", "plan9":
//line /usr/local/go/src/os/exec/exec.go:1175
				_go_fuzz_dep_.CoverTab[107423]++
//line /usr/local/go/src/os/exec/exec.go:1175
				// _ = "end of CoverTab[107423]"

//line /usr/local/go/src/os/exec/exec.go:1178
			default:
//line /usr/local/go/src/os/exec/exec.go:1178
				_go_fuzz_dep_.CoverTab[107424]++

//line /usr/local/go/src/os/exec/exec.go:1187
				if pwd, absErr := filepath.Abs(c.Dir); absErr == nil {
//line /usr/local/go/src/os/exec/exec.go:1187
					_go_fuzz_dep_.CoverTab[107425]++
										env = append(env, "PWD="+pwd)
//line /usr/local/go/src/os/exec/exec.go:1188
					// _ = "end of CoverTab[107425]"
				} else {
//line /usr/local/go/src/os/exec/exec.go:1189
					_go_fuzz_dep_.CoverTab[107426]++
//line /usr/local/go/src/os/exec/exec.go:1189
					if err == nil {
//line /usr/local/go/src/os/exec/exec.go:1189
						_go_fuzz_dep_.CoverTab[107427]++
											err = absErr
//line /usr/local/go/src/os/exec/exec.go:1190
						// _ = "end of CoverTab[107427]"
					} else {
//line /usr/local/go/src/os/exec/exec.go:1191
						_go_fuzz_dep_.CoverTab[107428]++
//line /usr/local/go/src/os/exec/exec.go:1191
						// _ = "end of CoverTab[107428]"
//line /usr/local/go/src/os/exec/exec.go:1191
					}
//line /usr/local/go/src/os/exec/exec.go:1191
					// _ = "end of CoverTab[107426]"
//line /usr/local/go/src/os/exec/exec.go:1191
				}
//line /usr/local/go/src/os/exec/exec.go:1191
				// _ = "end of CoverTab[107424]"
			}
//line /usr/local/go/src/os/exec/exec.go:1192
			// _ = "end of CoverTab[107422]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1193
			_go_fuzz_dep_.CoverTab[107429]++
//line /usr/local/go/src/os/exec/exec.go:1193
			// _ = "end of CoverTab[107429]"
//line /usr/local/go/src/os/exec/exec.go:1193
		}
//line /usr/local/go/src/os/exec/exec.go:1193
		// _ = "end of CoverTab[107419]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1194
		_go_fuzz_dep_.CoverTab[107430]++
//line /usr/local/go/src/os/exec/exec.go:1194
		// _ = "end of CoverTab[107430]"
//line /usr/local/go/src/os/exec/exec.go:1194
	}
//line /usr/local/go/src/os/exec/exec.go:1194
	// _ = "end of CoverTab[107415]"
//line /usr/local/go/src/os/exec/exec.go:1194
	_go_fuzz_dep_.CoverTab[107416]++

						env, dedupErr := dedupEnv(env)
						if err == nil {
//line /usr/local/go/src/os/exec/exec.go:1197
		_go_fuzz_dep_.CoverTab[107431]++
							err = dedupErr
//line /usr/local/go/src/os/exec/exec.go:1198
		// _ = "end of CoverTab[107431]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1199
		_go_fuzz_dep_.CoverTab[107432]++
//line /usr/local/go/src/os/exec/exec.go:1199
		// _ = "end of CoverTab[107432]"
//line /usr/local/go/src/os/exec/exec.go:1199
	}
//line /usr/local/go/src/os/exec/exec.go:1199
	// _ = "end of CoverTab[107416]"
//line /usr/local/go/src/os/exec/exec.go:1199
	_go_fuzz_dep_.CoverTab[107417]++
						return addCriticalEnv(env), err
//line /usr/local/go/src/os/exec/exec.go:1200
	// _ = "end of CoverTab[107417]"
}

// Environ returns a copy of the environment in which the command would be run
//line /usr/local/go/src/os/exec/exec.go:1203
// as it is currently configured.
//line /usr/local/go/src/os/exec/exec.go:1205
func (c *Cmd) Environ() []string {
//line /usr/local/go/src/os/exec/exec.go:1205
	_go_fuzz_dep_.CoverTab[107433]++

						env, _ := c.environ()
						return env
//line /usr/local/go/src/os/exec/exec.go:1208
	// _ = "end of CoverTab[107433]"
}

// dedupEnv returns a copy of env with any duplicates removed, in favor of
//line /usr/local/go/src/os/exec/exec.go:1211
// later values.
//line /usr/local/go/src/os/exec/exec.go:1211
// Items not of the normal environment "key=value" form are preserved unchanged.
//line /usr/local/go/src/os/exec/exec.go:1211
// Except on Plan 9, items containing NUL characters are removed, and
//line /usr/local/go/src/os/exec/exec.go:1211
// an error is returned along with the remaining values.
//line /usr/local/go/src/os/exec/exec.go:1216
func dedupEnv(env []string) ([]string, error) {
//line /usr/local/go/src/os/exec/exec.go:1216
	_go_fuzz_dep_.CoverTab[107434]++
						return dedupEnvCase(runtime.GOOS == "windows", runtime.GOOS == "plan9", env)
//line /usr/local/go/src/os/exec/exec.go:1217
	// _ = "end of CoverTab[107434]"
}

// dedupEnvCase is dedupEnv with a case option for testing.
//line /usr/local/go/src/os/exec/exec.go:1220
// If caseInsensitive is true, the case of keys is ignored.
//line /usr/local/go/src/os/exec/exec.go:1220
// If nulOK is false, items containing NUL characters are allowed.
//line /usr/local/go/src/os/exec/exec.go:1223
func dedupEnvCase(caseInsensitive, nulOK bool, env []string) ([]string, error) {
//line /usr/local/go/src/os/exec/exec.go:1223
	_go_fuzz_dep_.CoverTab[107435]++
	// Construct the output in reverse order, to preserve the
	// last occurrence of each key.
	var err error
	out := make([]string, 0, len(env))
	saw := make(map[string]bool, len(env))
	for n := len(env); n > 0; n-- {
//line /usr/local/go/src/os/exec/exec.go:1229
		_go_fuzz_dep_.CoverTab[107438]++
							kv := env[n-1]

//line /usr/local/go/src/os/exec/exec.go:1234
		if !nulOK && func() bool {
//line /usr/local/go/src/os/exec/exec.go:1234
			_go_fuzz_dep_.CoverTab[107444]++
//line /usr/local/go/src/os/exec/exec.go:1234
			return strings.IndexByte(kv, 0) != -1
//line /usr/local/go/src/os/exec/exec.go:1234
			// _ = "end of CoverTab[107444]"
//line /usr/local/go/src/os/exec/exec.go:1234
		}() {
//line /usr/local/go/src/os/exec/exec.go:1234
			_go_fuzz_dep_.CoverTab[107445]++
								err = errors.New("exec: environment variable contains NUL")
								continue
//line /usr/local/go/src/os/exec/exec.go:1236
			// _ = "end of CoverTab[107445]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1237
			_go_fuzz_dep_.CoverTab[107446]++
//line /usr/local/go/src/os/exec/exec.go:1237
			// _ = "end of CoverTab[107446]"
//line /usr/local/go/src/os/exec/exec.go:1237
		}
//line /usr/local/go/src/os/exec/exec.go:1237
		// _ = "end of CoverTab[107438]"
//line /usr/local/go/src/os/exec/exec.go:1237
		_go_fuzz_dep_.CoverTab[107439]++

							i := strings.Index(kv, "=")
							if i == 0 {
//line /usr/local/go/src/os/exec/exec.go:1240
			_go_fuzz_dep_.CoverTab[107447]++

//line /usr/local/go/src/os/exec/exec.go:1244
			i = strings.Index(kv[1:], "=") + 1
//line /usr/local/go/src/os/exec/exec.go:1244
			// _ = "end of CoverTab[107447]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1245
			_go_fuzz_dep_.CoverTab[107448]++
//line /usr/local/go/src/os/exec/exec.go:1245
			// _ = "end of CoverTab[107448]"
//line /usr/local/go/src/os/exec/exec.go:1245
		}
//line /usr/local/go/src/os/exec/exec.go:1245
		// _ = "end of CoverTab[107439]"
//line /usr/local/go/src/os/exec/exec.go:1245
		_go_fuzz_dep_.CoverTab[107440]++
							if i < 0 {
//line /usr/local/go/src/os/exec/exec.go:1246
			_go_fuzz_dep_.CoverTab[107449]++
								if kv != "" {
//line /usr/local/go/src/os/exec/exec.go:1247
				_go_fuzz_dep_.CoverTab[107451]++

//line /usr/local/go/src/os/exec/exec.go:1251
				out = append(out, kv)
//line /usr/local/go/src/os/exec/exec.go:1251
				// _ = "end of CoverTab[107451]"
			} else {
//line /usr/local/go/src/os/exec/exec.go:1252
				_go_fuzz_dep_.CoverTab[107452]++
//line /usr/local/go/src/os/exec/exec.go:1252
				// _ = "end of CoverTab[107452]"
//line /usr/local/go/src/os/exec/exec.go:1252
			}
//line /usr/local/go/src/os/exec/exec.go:1252
			// _ = "end of CoverTab[107449]"
//line /usr/local/go/src/os/exec/exec.go:1252
			_go_fuzz_dep_.CoverTab[107450]++
								continue
//line /usr/local/go/src/os/exec/exec.go:1253
			// _ = "end of CoverTab[107450]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1254
			_go_fuzz_dep_.CoverTab[107453]++
//line /usr/local/go/src/os/exec/exec.go:1254
			// _ = "end of CoverTab[107453]"
//line /usr/local/go/src/os/exec/exec.go:1254
		}
//line /usr/local/go/src/os/exec/exec.go:1254
		// _ = "end of CoverTab[107440]"
//line /usr/local/go/src/os/exec/exec.go:1254
		_go_fuzz_dep_.CoverTab[107441]++
							k := kv[:i]
							if caseInsensitive {
//line /usr/local/go/src/os/exec/exec.go:1256
			_go_fuzz_dep_.CoverTab[107454]++
								k = strings.ToLower(k)
//line /usr/local/go/src/os/exec/exec.go:1257
			// _ = "end of CoverTab[107454]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1258
			_go_fuzz_dep_.CoverTab[107455]++
//line /usr/local/go/src/os/exec/exec.go:1258
			// _ = "end of CoverTab[107455]"
//line /usr/local/go/src/os/exec/exec.go:1258
		}
//line /usr/local/go/src/os/exec/exec.go:1258
		// _ = "end of CoverTab[107441]"
//line /usr/local/go/src/os/exec/exec.go:1258
		_go_fuzz_dep_.CoverTab[107442]++
							if saw[k] {
//line /usr/local/go/src/os/exec/exec.go:1259
			_go_fuzz_dep_.CoverTab[107456]++
								continue
//line /usr/local/go/src/os/exec/exec.go:1260
			// _ = "end of CoverTab[107456]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1261
			_go_fuzz_dep_.CoverTab[107457]++
//line /usr/local/go/src/os/exec/exec.go:1261
			// _ = "end of CoverTab[107457]"
//line /usr/local/go/src/os/exec/exec.go:1261
		}
//line /usr/local/go/src/os/exec/exec.go:1261
		// _ = "end of CoverTab[107442]"
//line /usr/local/go/src/os/exec/exec.go:1261
		_go_fuzz_dep_.CoverTab[107443]++

							saw[k] = true
							out = append(out, kv)
//line /usr/local/go/src/os/exec/exec.go:1264
		// _ = "end of CoverTab[107443]"
	}
//line /usr/local/go/src/os/exec/exec.go:1265
	// _ = "end of CoverTab[107435]"
//line /usr/local/go/src/os/exec/exec.go:1265
	_go_fuzz_dep_.CoverTab[107436]++

//line /usr/local/go/src/os/exec/exec.go:1268
	for i := 0; i < len(out)/2; i++ {
//line /usr/local/go/src/os/exec/exec.go:1268
		_go_fuzz_dep_.CoverTab[107458]++
							j := len(out) - i - 1
							out[i], out[j] = out[j], out[i]
//line /usr/local/go/src/os/exec/exec.go:1270
		// _ = "end of CoverTab[107458]"
	}
//line /usr/local/go/src/os/exec/exec.go:1271
	// _ = "end of CoverTab[107436]"
//line /usr/local/go/src/os/exec/exec.go:1271
	_go_fuzz_dep_.CoverTab[107437]++

						return out, err
//line /usr/local/go/src/os/exec/exec.go:1273
	// _ = "end of CoverTab[107437]"
}

// addCriticalEnv adds any critical environment variables that are required
//line /usr/local/go/src/os/exec/exec.go:1276
// (or at least almost always required) on the operating system.
//line /usr/local/go/src/os/exec/exec.go:1276
// Currently this is only used for Windows.
//line /usr/local/go/src/os/exec/exec.go:1279
func addCriticalEnv(env []string) []string {
//line /usr/local/go/src/os/exec/exec.go:1279
	_go_fuzz_dep_.CoverTab[107459]++
						if runtime.GOOS != "windows" {
//line /usr/local/go/src/os/exec/exec.go:1280
		_go_fuzz_dep_.CoverTab[107462]++
							return env
//line /usr/local/go/src/os/exec/exec.go:1281
		// _ = "end of CoverTab[107462]"
	} else {
//line /usr/local/go/src/os/exec/exec.go:1282
		_go_fuzz_dep_.CoverTab[107463]++
//line /usr/local/go/src/os/exec/exec.go:1282
		// _ = "end of CoverTab[107463]"
//line /usr/local/go/src/os/exec/exec.go:1282
	}
//line /usr/local/go/src/os/exec/exec.go:1282
	// _ = "end of CoverTab[107459]"
//line /usr/local/go/src/os/exec/exec.go:1282
	_go_fuzz_dep_.CoverTab[107460]++
						for _, kv := range env {
//line /usr/local/go/src/os/exec/exec.go:1283
		_go_fuzz_dep_.CoverTab[107464]++
							k, _, ok := strings.Cut(kv, "=")
							if !ok {
//line /usr/local/go/src/os/exec/exec.go:1285
			_go_fuzz_dep_.CoverTab[107466]++
								continue
//line /usr/local/go/src/os/exec/exec.go:1286
			// _ = "end of CoverTab[107466]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1287
			_go_fuzz_dep_.CoverTab[107467]++
//line /usr/local/go/src/os/exec/exec.go:1287
			// _ = "end of CoverTab[107467]"
//line /usr/local/go/src/os/exec/exec.go:1287
		}
//line /usr/local/go/src/os/exec/exec.go:1287
		// _ = "end of CoverTab[107464]"
//line /usr/local/go/src/os/exec/exec.go:1287
		_go_fuzz_dep_.CoverTab[107465]++
							if strings.EqualFold(k, "SYSTEMROOT") {
//line /usr/local/go/src/os/exec/exec.go:1288
			_go_fuzz_dep_.CoverTab[107468]++

								return env
//line /usr/local/go/src/os/exec/exec.go:1290
			// _ = "end of CoverTab[107468]"
		} else {
//line /usr/local/go/src/os/exec/exec.go:1291
			_go_fuzz_dep_.CoverTab[107469]++
//line /usr/local/go/src/os/exec/exec.go:1291
			// _ = "end of CoverTab[107469]"
//line /usr/local/go/src/os/exec/exec.go:1291
		}
//line /usr/local/go/src/os/exec/exec.go:1291
		// _ = "end of CoverTab[107465]"
	}
//line /usr/local/go/src/os/exec/exec.go:1292
	// _ = "end of CoverTab[107460]"
//line /usr/local/go/src/os/exec/exec.go:1292
	_go_fuzz_dep_.CoverTab[107461]++
						return append(env, "SYSTEMROOT="+os.Getenv("SYSTEMROOT"))
//line /usr/local/go/src/os/exec/exec.go:1293
	// _ = "end of CoverTab[107461]"
}

// ErrDot indicates that a path lookup resolved to an executable
//line /usr/local/go/src/os/exec/exec.go:1296
// in the current directory due to ‘.’ being in the path, either
//line /usr/local/go/src/os/exec/exec.go:1296
// implicitly or explicitly. See the package documentation for details.
//line /usr/local/go/src/os/exec/exec.go:1296
//
//line /usr/local/go/src/os/exec/exec.go:1296
// Note that functions in this package do not return ErrDot directly.
//line /usr/local/go/src/os/exec/exec.go:1296
// Code should use errors.Is(err, ErrDot), not err == ErrDot,
//line /usr/local/go/src/os/exec/exec.go:1296
// to test whether a returned error err is due to this condition.
//line /usr/local/go/src/os/exec/exec.go:1303
var ErrDot = errors.New("cannot run executable found relative to current directory")
//line /usr/local/go/src/os/exec/exec.go:1303
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/exec/exec.go:1303
var _ = _go_fuzz_dep_.CoverTab
