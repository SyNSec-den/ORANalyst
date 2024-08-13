// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:5
)

import (
	"sync"

	"github.com/google/uuid"
)

// EventChannel is a channel which can accept an Event
type EventChannel chan Event

// Watchers stores the information about watchers
type Watchers struct {
	watchers	map[uuid.UUID]Watcher
	rm		sync.RWMutex
}

// Watcher event watcher
type Watcher struct {
	id	uuid.UUID
	ch	chan<- Event
}

// NewWatchers creates watchers
func NewWatchers() *Watchers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:29
	_go_fuzz_dep_.CoverTab[179574]++
													return &Watchers{
		watchers: make(map[uuid.UUID]Watcher),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:32
	// _ = "end of CoverTab[179574]"
}

// Send sends an event for all registered watchers
func (ws *Watchers) Send(event Event) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:36
	_go_fuzz_dep_.CoverTab[179575]++
													ws.rm.RLock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:37
	_curRoutineNum158_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:37
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum158_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
		_go_fuzz_dep_.CoverTab[179577]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
			_go_fuzz_dep_.CoverTab[179578]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum158_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
			// _ = "end of CoverTab[179578]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:38
		}()
														for _, watcher := range ws.watchers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:39
			_go_fuzz_dep_.CoverTab[179579]++
															watcher.ch <- event
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:40
			// _ = "end of CoverTab[179579]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:41
		// _ = "end of CoverTab[179577]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:42
	// _ = "end of CoverTab[179575]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:42
	_go_fuzz_dep_.CoverTab[179576]++
													ws.rm.RUnlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:43
	// _ = "end of CoverTab[179576]"
}

// AddWatcher adds a watcher
func (ws *Watchers) AddWatcher(id uuid.UUID, ch chan<- Event) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:47
	_go_fuzz_dep_.CoverTab[179580]++
													ws.rm.Lock()
													watcher := Watcher{
		id:	id,
		ch:	ch,
	}
													ws.watchers[id] = watcher
													ws.rm.Unlock()
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:55
	// _ = "end of CoverTab[179580]"

}

// RemoveWatcher removes a watcher
func (ws *Watchers) RemoveWatcher(id uuid.UUID) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:60
	_go_fuzz_dep_.CoverTab[179581]++
													ws.rm.Lock()
													watchers := make(map[uuid.UUID]Watcher, len(ws.watchers)-1)
													for _, watcher := range ws.watchers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:63
		_go_fuzz_dep_.CoverTab[179583]++
														if watcher.id != id {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:64
			_go_fuzz_dep_.CoverTab[179584]++
															watchers[id] = watcher
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:65
			// _ = "end of CoverTab[179584]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:67
			_go_fuzz_dep_.CoverTab[179585]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:67
			// _ = "end of CoverTab[179585]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:67
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:67
		// _ = "end of CoverTab[179583]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:68
	// _ = "end of CoverTab[179581]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:68
	_go_fuzz_dep_.CoverTab[179582]++
													ws.watchers = watchers
													ws.rm.Unlock()
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:71
	// _ = "end of CoverTab[179582]"

}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/watcher.go:73
var _ = _go_fuzz_dep_.CoverTab
