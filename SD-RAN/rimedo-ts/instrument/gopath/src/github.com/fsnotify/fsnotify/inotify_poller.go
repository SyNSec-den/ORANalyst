// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
package fsnotify

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:8
)

import (
	"errors"

	"golang.org/x/sys/unix"
)

type fdPoller struct {
	fd	int	// File descriptor (as returned by the inotify_init() syscall)
	epfd	int	// Epoll file descriptor
	pipe	[2]int	// Pipe for waking up
}

func emptyPoller(fd int) *fdPoller {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:22
	_go_fuzz_dep_.CoverTab[115115]++
												poller := new(fdPoller)
												poller.fd = fd
												poller.epfd = -1
												poller.pipe[0] = -1
												poller.pipe[1] = -1
												return poller
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:28
	// _ = "end of CoverTab[115115]"
}

// Create a new inotify poller.
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:31
// This creates an inotify handler, and an epoll handler.
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:33
func newFdPoller(fd int) (*fdPoller, error) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:33
	_go_fuzz_dep_.CoverTab[115116]++
												var errno error
												poller := emptyPoller(fd)
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:36
		_go_fuzz_dep_.CoverTab[115122]++
													if errno != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:37
			_go_fuzz_dep_.CoverTab[115123]++
														poller.close()
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:38
			// _ = "end of CoverTab[115123]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:39
			_go_fuzz_dep_.CoverTab[115124]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:39
			// _ = "end of CoverTab[115124]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:39
		// _ = "end of CoverTab[115122]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:40
	// _ = "end of CoverTab[115116]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:40
	_go_fuzz_dep_.CoverTab[115117]++
												poller.fd = fd

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:44
	poller.epfd, errno = unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if poller.epfd == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:45
		_go_fuzz_dep_.CoverTab[115125]++
													return nil, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:46
		// _ = "end of CoverTab[115125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:47
		_go_fuzz_dep_.CoverTab[115126]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:47
		// _ = "end of CoverTab[115126]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:47
	// _ = "end of CoverTab[115117]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:47
	_go_fuzz_dep_.CoverTab[115118]++

												errno = unix.Pipe2(poller.pipe[:], unix.O_NONBLOCK|unix.O_CLOEXEC)
												if errno != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:50
		_go_fuzz_dep_.CoverTab[115127]++
													return nil, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:51
		// _ = "end of CoverTab[115127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:52
		_go_fuzz_dep_.CoverTab[115128]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:52
		// _ = "end of CoverTab[115128]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:52
	// _ = "end of CoverTab[115118]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:52
	_go_fuzz_dep_.CoverTab[115119]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:55
	event := unix.EpollEvent{
		Fd:	int32(poller.fd),
		Events:	unix.EPOLLIN,
	}
	errno = unix.EpollCtl(poller.epfd, unix.EPOLL_CTL_ADD, poller.fd, &event)
	if errno != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:60
		_go_fuzz_dep_.CoverTab[115129]++
													return nil, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:61
		// _ = "end of CoverTab[115129]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:62
		_go_fuzz_dep_.CoverTab[115130]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:62
		// _ = "end of CoverTab[115130]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:62
	// _ = "end of CoverTab[115119]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:62
	_go_fuzz_dep_.CoverTab[115120]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:65
	event = unix.EpollEvent{
		Fd:	int32(poller.pipe[0]),
		Events:	unix.EPOLLIN,
	}
	errno = unix.EpollCtl(poller.epfd, unix.EPOLL_CTL_ADD, poller.pipe[0], &event)
	if errno != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:70
		_go_fuzz_dep_.CoverTab[115131]++
													return nil, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:71
		// _ = "end of CoverTab[115131]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:72
		_go_fuzz_dep_.CoverTab[115132]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:72
		// _ = "end of CoverTab[115132]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:72
	// _ = "end of CoverTab[115120]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:72
	_go_fuzz_dep_.CoverTab[115121]++

												return poller, nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:74
	// _ = "end of CoverTab[115121]"
}

// Wait using epoll.
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:77
// Returns true if something is ready to be read,
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:77
// false if there is not.
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:80
func (poller *fdPoller) wait() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:80
	_go_fuzz_dep_.CoverTab[115133]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:85
	events := make([]unix.EpollEvent, 7)
	for {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:86
		_go_fuzz_dep_.CoverTab[115134]++
													n, errno := unix.EpollWait(poller.epfd, events, -1)
													if n == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:88
			_go_fuzz_dep_.CoverTab[115140]++
														if errno == unix.EINTR {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:89
				_go_fuzz_dep_.CoverTab[115142]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:90
				// _ = "end of CoverTab[115142]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:91
				_go_fuzz_dep_.CoverTab[115143]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:91
				// _ = "end of CoverTab[115143]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:91
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:91
			// _ = "end of CoverTab[115140]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:91
			_go_fuzz_dep_.CoverTab[115141]++
														return false, errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:92
			// _ = "end of CoverTab[115141]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:93
			_go_fuzz_dep_.CoverTab[115144]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:93
			// _ = "end of CoverTab[115144]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:93
		// _ = "end of CoverTab[115134]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:93
		_go_fuzz_dep_.CoverTab[115135]++
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:94
			_go_fuzz_dep_.CoverTab[115145]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:96
			// _ = "end of CoverTab[115145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:97
			_go_fuzz_dep_.CoverTab[115146]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:97
			// _ = "end of CoverTab[115146]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:97
		// _ = "end of CoverTab[115135]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:97
		_go_fuzz_dep_.CoverTab[115136]++
													if n > 6 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:98
			_go_fuzz_dep_.CoverTab[115147]++

														return false, errors.New("epoll_wait returned more events than I know what to do with")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:100
			// _ = "end of CoverTab[115147]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:101
			_go_fuzz_dep_.CoverTab[115148]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:101
			// _ = "end of CoverTab[115148]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:101
		// _ = "end of CoverTab[115136]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:101
		_go_fuzz_dep_.CoverTab[115137]++
													ready := events[:n]
													epollhup := false
													epollerr := false
													epollin := false
													for _, event := range ready {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:106
			_go_fuzz_dep_.CoverTab[115149]++
														if event.Fd == int32(poller.fd) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:107
				_go_fuzz_dep_.CoverTab[115151]++
															if event.Events&unix.EPOLLHUP != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:108
					_go_fuzz_dep_.CoverTab[115154]++

																epollhup = true
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:110
					// _ = "end of CoverTab[115154]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:111
					_go_fuzz_dep_.CoverTab[115155]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:111
					// _ = "end of CoverTab[115155]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:111
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:111
				// _ = "end of CoverTab[115151]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:111
				_go_fuzz_dep_.CoverTab[115152]++
															if event.Events&unix.EPOLLERR != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:112
					_go_fuzz_dep_.CoverTab[115156]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:115
					epollerr = true
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:115
					// _ = "end of CoverTab[115156]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:116
					_go_fuzz_dep_.CoverTab[115157]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:116
					// _ = "end of CoverTab[115157]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:116
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:116
				// _ = "end of CoverTab[115152]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:116
				_go_fuzz_dep_.CoverTab[115153]++
															if event.Events&unix.EPOLLIN != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:117
					_go_fuzz_dep_.CoverTab[115158]++

																epollin = true
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:119
					// _ = "end of CoverTab[115158]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:120
					_go_fuzz_dep_.CoverTab[115159]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:120
					// _ = "end of CoverTab[115159]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:120
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:120
				// _ = "end of CoverTab[115153]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:121
				_go_fuzz_dep_.CoverTab[115160]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:121
				// _ = "end of CoverTab[115160]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:121
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:121
			// _ = "end of CoverTab[115149]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:121
			_go_fuzz_dep_.CoverTab[115150]++
														if event.Fd == int32(poller.pipe[0]) {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:122
				_go_fuzz_dep_.CoverTab[115161]++
															if event.Events&unix.EPOLLHUP != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:123
					_go_fuzz_dep_.CoverTab[115164]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:123
					// _ = "end of CoverTab[115164]"

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
					_go_fuzz_dep_.CoverTab[115165]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
					// _ = "end of CoverTab[115165]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
				// _ = "end of CoverTab[115161]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:126
				_go_fuzz_dep_.CoverTab[115162]++
															if event.Events&unix.EPOLLERR != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:127
					_go_fuzz_dep_.CoverTab[115166]++

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:130
					return false, errors.New("Error on the pipe descriptor.")
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:130
					// _ = "end of CoverTab[115166]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:131
					_go_fuzz_dep_.CoverTab[115167]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:131
					// _ = "end of CoverTab[115167]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:131
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:131
				// _ = "end of CoverTab[115162]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:131
				_go_fuzz_dep_.CoverTab[115163]++
															if event.Events&unix.EPOLLIN != 0 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:132
					_go_fuzz_dep_.CoverTab[115168]++

																err := poller.clearWake()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:135
						_go_fuzz_dep_.CoverTab[115169]++
																	return false, err
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:136
						// _ = "end of CoverTab[115169]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:137
						_go_fuzz_dep_.CoverTab[115170]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:137
						// _ = "end of CoverTab[115170]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:137
					}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:137
					// _ = "end of CoverTab[115168]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:138
					_go_fuzz_dep_.CoverTab[115171]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:138
					// _ = "end of CoverTab[115171]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:138
				}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:138
				// _ = "end of CoverTab[115163]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:139
				_go_fuzz_dep_.CoverTab[115172]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:139
				// _ = "end of CoverTab[115172]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:139
			}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:139
			// _ = "end of CoverTab[115150]"
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:140
		// _ = "end of CoverTab[115137]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:140
		_go_fuzz_dep_.CoverTab[115138]++

													if epollhup || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			_go_fuzz_dep_.CoverTab[115173]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			return epollerr
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			// _ = "end of CoverTab[115173]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			_go_fuzz_dep_.CoverTab[115174]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			return epollin
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			// _ = "end of CoverTab[115174]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
		}() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:142
			_go_fuzz_dep_.CoverTab[115175]++
														return true, nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:143
			// _ = "end of CoverTab[115175]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:144
			_go_fuzz_dep_.CoverTab[115176]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:144
			// _ = "end of CoverTab[115176]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:144
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:144
		// _ = "end of CoverTab[115138]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:144
		_go_fuzz_dep_.CoverTab[115139]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:145
		// _ = "end of CoverTab[115139]"
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:146
	// _ = "end of CoverTab[115133]"
}

// Close the write end of the poller.
func (poller *fdPoller) wake() error {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:150
	_go_fuzz_dep_.CoverTab[115177]++
												buf := make([]byte, 1)
												n, errno := unix.Write(poller.pipe[1], buf)
												if n == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:153
		_go_fuzz_dep_.CoverTab[115179]++
													if errno == unix.EAGAIN {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:154
			_go_fuzz_dep_.CoverTab[115181]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:156
			// _ = "end of CoverTab[115181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:157
			_go_fuzz_dep_.CoverTab[115182]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:157
			// _ = "end of CoverTab[115182]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:157
		// _ = "end of CoverTab[115179]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:157
		_go_fuzz_dep_.CoverTab[115180]++
													return errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:158
		// _ = "end of CoverTab[115180]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:159
		_go_fuzz_dep_.CoverTab[115183]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:159
		// _ = "end of CoverTab[115183]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:159
	// _ = "end of CoverTab[115177]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:159
	_go_fuzz_dep_.CoverTab[115178]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:160
	// _ = "end of CoverTab[115178]"
}

func (poller *fdPoller) clearWake() error {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:163
	_go_fuzz_dep_.CoverTab[115184]++

												buf := make([]byte, 100)
												n, errno := unix.Read(poller.pipe[0], buf)
												if n == -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:167
		_go_fuzz_dep_.CoverTab[115186]++
													if errno == unix.EAGAIN {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:168
			_go_fuzz_dep_.CoverTab[115188]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:170
			// _ = "end of CoverTab[115188]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:171
			_go_fuzz_dep_.CoverTab[115189]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:171
			// _ = "end of CoverTab[115189]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:171
		// _ = "end of CoverTab[115186]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:171
		_go_fuzz_dep_.CoverTab[115187]++
													return errno
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:172
		// _ = "end of CoverTab[115187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:173
		_go_fuzz_dep_.CoverTab[115190]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:173
		// _ = "end of CoverTab[115190]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:173
	// _ = "end of CoverTab[115184]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:173
	_go_fuzz_dep_.CoverTab[115185]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:174
	// _ = "end of CoverTab[115185]"
}

// Close all poller file descriptors, but not the one passed to it.
func (poller *fdPoller) close() {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:178
	_go_fuzz_dep_.CoverTab[115191]++
												if poller.pipe[1] != -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:179
		_go_fuzz_dep_.CoverTab[115194]++
													unix.Close(poller.pipe[1])
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:180
		// _ = "end of CoverTab[115194]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:181
		_go_fuzz_dep_.CoverTab[115195]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:181
		// _ = "end of CoverTab[115195]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:181
	// _ = "end of CoverTab[115191]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:181
	_go_fuzz_dep_.CoverTab[115192]++
												if poller.pipe[0] != -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:182
		_go_fuzz_dep_.CoverTab[115196]++
													unix.Close(poller.pipe[0])
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:183
		// _ = "end of CoverTab[115196]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:184
		_go_fuzz_dep_.CoverTab[115197]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:184
		// _ = "end of CoverTab[115197]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:184
	// _ = "end of CoverTab[115192]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:184
	_go_fuzz_dep_.CoverTab[115193]++
												if poller.epfd != -1 {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:185
		_go_fuzz_dep_.CoverTab[115198]++
													unix.Close(poller.epfd)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:186
		// _ = "end of CoverTab[115198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:187
		_go_fuzz_dep_.CoverTab[115199]++
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:187
		// _ = "end of CoverTab[115199]"
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:187
	// _ = "end of CoverTab[115193]"
}

//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:188
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/fsnotify/fsnotify@v1.5.1/inotify_poller.go:188
var _ = _go_fuzz_dep_.CoverTab
