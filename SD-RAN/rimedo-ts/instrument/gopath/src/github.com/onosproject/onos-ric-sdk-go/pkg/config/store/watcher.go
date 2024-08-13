// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:5
)

import (
	"sync"

	"github.com/google/uuid"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
)

// EventChannel is a channel which can accept an Event
type EventChannel chan event.Event

// EventBus stores the information about watchers
type EventBus struct {
	watchers	[]ConfigTreeWatcher
	rm		sync.RWMutex
}

// ConfigTreeWatcher :
type ConfigTreeWatcher struct {
	id	uuid.UUID
	ch	chan event.Event
}

func (eb *EventBus) send(event event.Event) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:29
	_go_fuzz_dep_.CoverTab[193802]++
														eb.rm.RLock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:30
	_curRoutineNum174_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:30
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum174_)
														go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
		_go_fuzz_dep_.CoverTab[193804]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
			_go_fuzz_dep_.CoverTab[193805]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum174_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
			// _ = "end of CoverTab[193805]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:31
		}()
															for _, watcher := range eb.watchers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:32
			_go_fuzz_dep_.CoverTab[193806]++
																watcher.ch <- event
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:33
			// _ = "end of CoverTab[193806]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:34
		// _ = "end of CoverTab[193804]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:35
	// _ = "end of CoverTab[193802]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:35
	_go_fuzz_dep_.CoverTab[193803]++
														eb.rm.RUnlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:36
	// _ = "end of CoverTab[193803]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/watcher.go:37
var _ = _go_fuzz_dep_.CoverTab
