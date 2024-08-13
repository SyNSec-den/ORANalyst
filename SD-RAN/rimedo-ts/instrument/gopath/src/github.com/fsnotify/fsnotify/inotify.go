// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
package fsnotify

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:8
)

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Watcher watches a set of files, delivering events to a channel.
type Watcher struct {
	Events		chan Event
	Errors		chan error
	mu		sync.Mutex	// Map access
	fd		int
	poller		*fdPoller
	watches		map[string]*watch	// Map of inotify watches (key: path)
	paths		map[int]string		// Map of watched paths (key: watch descriptor)
	done		chan struct{}		// Channel for sending a "quit message" to the reader goroutine
	doneResp	chan struct{}		// Channel to respond to Close
}

// NewWatcher establishes a new watcher with the underlying OS and begins waiting for events.
func NewWatcher() (*Watcher, error) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:37
	_go_fuzz_dep_.CoverTab[115005]++

											fd, errno := unix.InotifyInit1(unix.IN_CLOEXEC)
											if fd == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:40
		_go_fuzz_dep_.CoverTab[115008]++
												return nil, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:41
		// _ = "end of CoverTab[115008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:42
		_go_fuzz_dep_.CoverTab[115009]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:42
		// _ = "end of CoverTab[115009]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:42
	// _ = "end of CoverTab[115005]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:42
	_go_fuzz_dep_.CoverTab[115006]++

											poller, err := newFdPoller(fd)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:45
		_go_fuzz_dep_.CoverTab[115010]++
												unix.Close(fd)
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:47
		// _ = "end of CoverTab[115010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:48
		_go_fuzz_dep_.CoverTab[115011]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:48
		// _ = "end of CoverTab[115011]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:48
	// _ = "end of CoverTab[115006]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:48
	_go_fuzz_dep_.CoverTab[115007]++
											w := &Watcher{
		fd:		fd,
		poller:		poller,
		watches:	make(map[string]*watch),
		paths:		make(map[int]string),
		Events:		make(chan Event),
		Errors:		make(chan error),
		done:		make(chan struct{}),
		doneResp:	make(chan struct{}),
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:58
	_curRoutineNum152_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:58
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum152_)

											go w.readEvents()
											return w, nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:61
	// _ = "end of CoverTab[115007]"
}

func (w *Watcher) isClosed() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:64
	_go_fuzz_dep_.CoverTab[115012]++
											select {
	case <-w.done:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:66
		_go_fuzz_dep_.CoverTab[115013]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:67
		// _ = "end of CoverTab[115013]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:68
		_go_fuzz_dep_.CoverTab[115014]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:69
		// _ = "end of CoverTab[115014]"
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:70
	// _ = "end of CoverTab[115012]"
}

// Close removes all watches and closes the events channel.
func (w *Watcher) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:74
	_go_fuzz_dep_.CoverTab[115015]++
											if w.isClosed() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:75
		_go_fuzz_dep_.CoverTab[115017]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:76
		// _ = "end of CoverTab[115017]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:77
		_go_fuzz_dep_.CoverTab[115018]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:77
		// _ = "end of CoverTab[115018]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:77
	// _ = "end of CoverTab[115015]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:77
	_go_fuzz_dep_.CoverTab[115016]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:80
	close(w.done)

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:83
	w.poller.wake()

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:86
	<-w.doneResp

											return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:88
	// _ = "end of CoverTab[115016]"
}

// Add starts watching the named file or directory (non-recursively).
func (w *Watcher) Add(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:92
	_go_fuzz_dep_.CoverTab[115019]++
											name = filepath.Clean(name)
											if w.isClosed() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:94
		_go_fuzz_dep_.CoverTab[115024]++
												return errors.New("inotify instance already closed")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:95
		// _ = "end of CoverTab[115024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:96
		_go_fuzz_dep_.CoverTab[115025]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:96
		// _ = "end of CoverTab[115025]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:96
	// _ = "end of CoverTab[115019]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:96
	_go_fuzz_dep_.CoverTab[115020]++

											const agnosticEvents = unix.IN_MOVED_TO | unix.IN_MOVED_FROM |
		unix.IN_CREATE | unix.IN_ATTRIB | unix.IN_MODIFY |
		unix.IN_MOVE_SELF | unix.IN_DELETE | unix.IN_DELETE_SELF

	var flags uint32 = agnosticEvents

	w.mu.Lock()
	defer w.mu.Unlock()
	watchEntry := w.watches[name]
	if watchEntry != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:107
		_go_fuzz_dep_.CoverTab[115026]++
												flags |= watchEntry.flags | unix.IN_MASK_ADD
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:108
		// _ = "end of CoverTab[115026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:109
		_go_fuzz_dep_.CoverTab[115027]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:109
		// _ = "end of CoverTab[115027]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:109
	// _ = "end of CoverTab[115020]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:109
	_go_fuzz_dep_.CoverTab[115021]++
											wd, errno := unix.InotifyAddWatch(w.fd, name, flags)
											if wd == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:111
		_go_fuzz_dep_.CoverTab[115028]++
												return errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:112
		// _ = "end of CoverTab[115028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:113
		_go_fuzz_dep_.CoverTab[115029]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:113
		// _ = "end of CoverTab[115029]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:113
	// _ = "end of CoverTab[115021]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:113
	_go_fuzz_dep_.CoverTab[115022]++

											if watchEntry == nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:115
		_go_fuzz_dep_.CoverTab[115030]++
												w.watches[name] = &watch{wd: uint32(wd), flags: flags}
												w.paths[wd] = name
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:117
		// _ = "end of CoverTab[115030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:118
		_go_fuzz_dep_.CoverTab[115031]++
												watchEntry.wd = uint32(wd)
												watchEntry.flags = flags
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:120
		// _ = "end of CoverTab[115031]"
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:121
	// _ = "end of CoverTab[115022]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:121
	_go_fuzz_dep_.CoverTab[115023]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:123
	// _ = "end of CoverTab[115023]"
}

// Remove stops watching the named file or directory (non-recursively).
func (w *Watcher) Remove(name string) error {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:127
	_go_fuzz_dep_.CoverTab[115032]++
											name = filepath.Clean(name)

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:131
	w.mu.Lock()
											defer w.mu.Unlock()
											watch, ok := w.watches[name]

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:136
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:136
		_go_fuzz_dep_.CoverTab[115035]++
												return fmt.Errorf("can't remove non-existent inotify watch for: %s", name)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:137
		// _ = "end of CoverTab[115035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:138
		_go_fuzz_dep_.CoverTab[115036]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:138
		// _ = "end of CoverTab[115036]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:138
	// _ = "end of CoverTab[115032]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:138
	_go_fuzz_dep_.CoverTab[115033]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:143
	delete(w.paths, int(watch.wd))
											delete(w.watches, name)

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:152
	success, errno := unix.InotifyRmWatch(w.fd, watch.wd)
	if success == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:153
		_go_fuzz_dep_.CoverTab[115037]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:160
		return errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:160
		// _ = "end of CoverTab[115037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:161
		_go_fuzz_dep_.CoverTab[115038]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:161
		// _ = "end of CoverTab[115038]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:161
	// _ = "end of CoverTab[115033]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:161
	_go_fuzz_dep_.CoverTab[115034]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:163
	// _ = "end of CoverTab[115034]"
}

type watch struct {
	wd	uint32	// Watch descriptor (as returned by the inotify_add_watch() syscall)
	flags	uint32	// inotify flags of this watch (see inotify(7) for the list of valid flags)
}

// readEvents reads from the inotify file descriptor, converts the
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:171
// received events into Event objects and sends them via the Events channel
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:173
func (w *Watcher) readEvents() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:173
	_go_fuzz_dep_.CoverTab[115039]++
											var (
		buf	[unix.SizeofInotifyEvent * 4096]byte	// Buffer for a maximum of 4096 raw events
		n	int					// Number of bytes read with read()
		errno	error					// Syscall errno
		ok	bool					// For poller.wait
	)

	defer close(w.doneResp)
	defer close(w.Errors)
	defer close(w.Events)
	defer unix.Close(w.fd)
	defer w.poller.close()

	for {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:187
		_go_fuzz_dep_.CoverTab[115040]++

												if w.isClosed() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:189
			_go_fuzz_dep_.CoverTab[115047]++
													return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:190
			// _ = "end of CoverTab[115047]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:191
			_go_fuzz_dep_.CoverTab[115048]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:191
			// _ = "end of CoverTab[115048]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:191
		// _ = "end of CoverTab[115040]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:191
		_go_fuzz_dep_.CoverTab[115041]++

												ok, errno = w.poller.wait()
												if errno != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:194
			_go_fuzz_dep_.CoverTab[115049]++
													select {
			case w.Errors <- errno:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:196
				_go_fuzz_dep_.CoverTab[115051]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:196
				// _ = "end of CoverTab[115051]"
			case <-w.done:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:197
				_go_fuzz_dep_.CoverTab[115052]++
														return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:198
				// _ = "end of CoverTab[115052]"
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:199
			// _ = "end of CoverTab[115049]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:199
			_go_fuzz_dep_.CoverTab[115050]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:200
			// _ = "end of CoverTab[115050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:201
			_go_fuzz_dep_.CoverTab[115053]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:201
			// _ = "end of CoverTab[115053]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:201
		// _ = "end of CoverTab[115041]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:201
		_go_fuzz_dep_.CoverTab[115042]++

												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:203
			_go_fuzz_dep_.CoverTab[115054]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:204
			// _ = "end of CoverTab[115054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:205
			_go_fuzz_dep_.CoverTab[115055]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:205
			// _ = "end of CoverTab[115055]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:205
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:205
		// _ = "end of CoverTab[115042]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:205
		_go_fuzz_dep_.CoverTab[115043]++

												n, errno = unix.Read(w.fd, buf[:])

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:211
		if errno == unix.EINTR {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:211
			_go_fuzz_dep_.CoverTab[115056]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:212
			// _ = "end of CoverTab[115056]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:213
			_go_fuzz_dep_.CoverTab[115057]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:213
			// _ = "end of CoverTab[115057]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:213
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:213
		// _ = "end of CoverTab[115043]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:213
		_go_fuzz_dep_.CoverTab[115044]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:216
		if w.isClosed() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:216
			_go_fuzz_dep_.CoverTab[115058]++
													return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:217
			// _ = "end of CoverTab[115058]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:218
			_go_fuzz_dep_.CoverTab[115059]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:218
			// _ = "end of CoverTab[115059]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:218
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:218
		// _ = "end of CoverTab[115044]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:218
		_go_fuzz_dep_.CoverTab[115045]++

												if n < unix.SizeofInotifyEvent {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:220
			_go_fuzz_dep_.CoverTab[115060]++
													var err error
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:222
				_go_fuzz_dep_.CoverTab[115063]++

														err = io.EOF
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:224
				// _ = "end of CoverTab[115063]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:225
				_go_fuzz_dep_.CoverTab[115064]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:225
				if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:225
					_go_fuzz_dep_.CoverTab[115065]++

															err = errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:227
					// _ = "end of CoverTab[115065]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:228
					_go_fuzz_dep_.CoverTab[115066]++

															err = errors.New("notify: short read in readEvents()")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:230
					// _ = "end of CoverTab[115066]"
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:231
				// _ = "end of CoverTab[115064]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:231
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:231
			// _ = "end of CoverTab[115060]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:231
			_go_fuzz_dep_.CoverTab[115061]++
													select {
			case w.Errors <- err:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:233
				_go_fuzz_dep_.CoverTab[115067]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:233
				// _ = "end of CoverTab[115067]"
			case <-w.done:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:234
				_go_fuzz_dep_.CoverTab[115068]++
														return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:235
				// _ = "end of CoverTab[115068]"
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:236
			// _ = "end of CoverTab[115061]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:236
			_go_fuzz_dep_.CoverTab[115062]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:237
			// _ = "end of CoverTab[115062]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:238
			_go_fuzz_dep_.CoverTab[115069]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:238
			// _ = "end of CoverTab[115069]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:238
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:238
		// _ = "end of CoverTab[115045]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:238
		_go_fuzz_dep_.CoverTab[115046]++

												var offset uint32

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:243
		for offset <= uint32(n-unix.SizeofInotifyEvent) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:243
			_go_fuzz_dep_.CoverTab[115070]++

													raw := (*unix.InotifyEvent)(unsafe.Pointer(&buf[offset]))

													mask := uint32(raw.Mask)
													nameLen := uint32(raw.Len)

													if mask&unix.IN_Q_OVERFLOW != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:250
				_go_fuzz_dep_.CoverTab[115075]++
														select {
				case w.Errors <- ErrEventOverflow:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:252
					_go_fuzz_dep_.CoverTab[115076]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:252
					// _ = "end of CoverTab[115076]"
				case <-w.done:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:253
					_go_fuzz_dep_.CoverTab[115077]++
															return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:254
					// _ = "end of CoverTab[115077]"
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:255
				// _ = "end of CoverTab[115075]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:256
				_go_fuzz_dep_.CoverTab[115078]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:256
				// _ = "end of CoverTab[115078]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:256
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:256
			// _ = "end of CoverTab[115070]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:256
			_go_fuzz_dep_.CoverTab[115071]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:262
			w.mu.Lock()
													name, ok := w.paths[int(raw.Wd)]

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
			if ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
				_go_fuzz_dep_.CoverTab[115079]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
				return mask&unix.IN_DELETE_SELF == unix.IN_DELETE_SELF
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
				// _ = "end of CoverTab[115079]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
			}() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:268
				_go_fuzz_dep_.CoverTab[115080]++
														delete(w.paths, int(raw.Wd))
														delete(w.watches, name)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:270
				// _ = "end of CoverTab[115080]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:271
				_go_fuzz_dep_.CoverTab[115081]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:271
				// _ = "end of CoverTab[115081]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:271
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:271
			// _ = "end of CoverTab[115071]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:271
			_go_fuzz_dep_.CoverTab[115072]++
													w.mu.Unlock()

													if nameLen > 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:274
				_go_fuzz_dep_.CoverTab[115082]++

														bytes := (*[unix.PathMax]byte)(unsafe.Pointer(&buf[offset+unix.SizeofInotifyEvent]))[:nameLen:nameLen]

														name += "/" + strings.TrimRight(string(bytes[0:nameLen]), "\000")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:278
				// _ = "end of CoverTab[115082]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:279
				_go_fuzz_dep_.CoverTab[115083]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:279
				// _ = "end of CoverTab[115083]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:279
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:279
			// _ = "end of CoverTab[115072]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:279
			_go_fuzz_dep_.CoverTab[115073]++

													event := newEvent(name, mask)

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:284
			if !event.ignoreLinux(mask) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:284
				_go_fuzz_dep_.CoverTab[115084]++
														select {
				case w.Events <- event:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:286
					_go_fuzz_dep_.CoverTab[115085]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:286
					// _ = "end of CoverTab[115085]"
				case <-w.done:
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:287
					_go_fuzz_dep_.CoverTab[115086]++
															return
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:288
					// _ = "end of CoverTab[115086]"
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:289
				// _ = "end of CoverTab[115084]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:290
				_go_fuzz_dep_.CoverTab[115087]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:290
				// _ = "end of CoverTab[115087]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:290
			// _ = "end of CoverTab[115073]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:290
			_go_fuzz_dep_.CoverTab[115074]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:293
			offset += unix.SizeofInotifyEvent + nameLen
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:293
			// _ = "end of CoverTab[115074]"
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:294
		// _ = "end of CoverTab[115046]"
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:295
	// _ = "end of CoverTab[115039]"
}

// Certain types of events can be "ignored" and not sent over the Events
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:298
// channel. Such as events marked ignore by the kernel, or MODIFY events
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:298
// against files that do not exist.
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:301
func (e *Event) ignoreLinux(mask uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:301
	_go_fuzz_dep_.CoverTab[115088]++

											if mask&unix.IN_IGNORED == unix.IN_IGNORED {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:303
		_go_fuzz_dep_.CoverTab[115091]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:304
		// _ = "end of CoverTab[115091]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:305
		_go_fuzz_dep_.CoverTab[115092]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:305
		// _ = "end of CoverTab[115092]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:305
	// _ = "end of CoverTab[115088]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:305
	_go_fuzz_dep_.CoverTab[115089]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
	if !(e.Op&Remove == Remove || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
		_go_fuzz_dep_.CoverTab[115093]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
		return e.Op&Rename == Rename
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
		// _ = "end of CoverTab[115093]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
	}()) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:312
		_go_fuzz_dep_.CoverTab[115094]++
												_, statErr := os.Lstat(e.Name)
												return os.IsNotExist(statErr)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:314
		// _ = "end of CoverTab[115094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:315
		_go_fuzz_dep_.CoverTab[115095]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:315
		// _ = "end of CoverTab[115095]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:315
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:315
	// _ = "end of CoverTab[115089]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:315
	_go_fuzz_dep_.CoverTab[115090]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:316
	// _ = "end of CoverTab[115090]"
}

// newEvent returns an platform-independent Event based on an inotify mask.
func newEvent(name string, mask uint32) Event {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:320
	_go_fuzz_dep_.CoverTab[115096]++
											e := Event{Name: name}
											if mask&unix.IN_CREATE == unix.IN_CREATE || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:322
		_go_fuzz_dep_.CoverTab[115102]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:322
		return mask&unix.IN_MOVED_TO == unix.IN_MOVED_TO
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:322
		// _ = "end of CoverTab[115102]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:322
	}() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:322
		_go_fuzz_dep_.CoverTab[115103]++
												e.Op |= Create
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:323
		// _ = "end of CoverTab[115103]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:324
		_go_fuzz_dep_.CoverTab[115104]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:324
		// _ = "end of CoverTab[115104]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:324
	// _ = "end of CoverTab[115096]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:324
	_go_fuzz_dep_.CoverTab[115097]++
											if mask&unix.IN_DELETE_SELF == unix.IN_DELETE_SELF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:325
		_go_fuzz_dep_.CoverTab[115105]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:325
		return mask&unix.IN_DELETE == unix.IN_DELETE
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:325
		// _ = "end of CoverTab[115105]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:325
	}() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:325
		_go_fuzz_dep_.CoverTab[115106]++
												e.Op |= Remove
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:326
		// _ = "end of CoverTab[115106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:327
		_go_fuzz_dep_.CoverTab[115107]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:327
		// _ = "end of CoverTab[115107]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:327
	// _ = "end of CoverTab[115097]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:327
	_go_fuzz_dep_.CoverTab[115098]++
											if mask&unix.IN_MODIFY == unix.IN_MODIFY {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:328
		_go_fuzz_dep_.CoverTab[115108]++
												e.Op |= Write
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:329
		// _ = "end of CoverTab[115108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:330
		_go_fuzz_dep_.CoverTab[115109]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:330
		// _ = "end of CoverTab[115109]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:330
	// _ = "end of CoverTab[115098]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:330
	_go_fuzz_dep_.CoverTab[115099]++
											if mask&unix.IN_MOVE_SELF == unix.IN_MOVE_SELF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:331
		_go_fuzz_dep_.CoverTab[115110]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:331
		return mask&unix.IN_MOVED_FROM == unix.IN_MOVED_FROM
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:331
		// _ = "end of CoverTab[115110]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:331
	}() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:331
		_go_fuzz_dep_.CoverTab[115111]++
												e.Op |= Rename
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:332
		// _ = "end of CoverTab[115111]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:333
		_go_fuzz_dep_.CoverTab[115112]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:333
		// _ = "end of CoverTab[115112]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:333
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:333
	// _ = "end of CoverTab[115099]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:333
	_go_fuzz_dep_.CoverTab[115100]++
											if mask&unix.IN_ATTRIB == unix.IN_ATTRIB {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:334
		_go_fuzz_dep_.CoverTab[115113]++
												e.Op |= Chmod
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:335
		// _ = "end of CoverTab[115113]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:336
		_go_fuzz_dep_.CoverTab[115114]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:336
		// _ = "end of CoverTab[115114]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:336
	// _ = "end of CoverTab[115100]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:336
	_go_fuzz_dep_.CoverTab[115101]++
											return e
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:337
	// _ = "end of CoverTab[115101]"
}

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:338
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify.go:338
var _ = _go_fuzz_dep_.CoverTab
