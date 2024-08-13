// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:5
)

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/google/uuid"
)

var log = logging.GetLogger("config", "store")

// Entry config entry
type Entry struct {
	Key		string
	Value		interface{}
	EventType	string
}

// Store :
type Store interface {
	Put(key string, entry Entry) error

	Get(key string) (Entry, error)

	Watch(ctx context.Context, ch chan event.Event) error
}

// ConfigStore :
type ConfigStore struct {
	ConfigTree	map[string]interface{}
	eventBus	*EventBus
}

// NewConfigStore :
func NewConfigStore() *ConfigStore {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:42
	_go_fuzz_dep_.CoverTab[193745]++
														return &ConfigStore{
		ConfigTree:	make(map[string]interface{}),
		eventBus: &EventBus{
			watchers: []ConfigTreeWatcher{},
		},
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:48
	// _ = "end of CoverTab[193745]"
}

// Watch :
func (c *ConfigStore) Watch(ctx context.Context, ch chan event.Event) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:52
	_go_fuzz_dep_.CoverTab[193746]++
														c.eventBus.rm.Lock()
														cw := ConfigTreeWatcher{
		id:	uuid.New(),
		ch:	ch,
	}

														c.eventBus.watchers = append(c.eventBus.watchers, cw)
														c.eventBus.rm.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:60
	_curRoutineNum173_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:60
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum173_)

														go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
		_go_fuzz_dep_.CoverTab[193748]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
			_go_fuzz_dep_.CoverTab[193750]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum173_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
			// _ = "end of CoverTab[193750]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:62
		}()
															<-ctx.Done()
															c.eventBus.rm.Lock()
															watchers := make([]ConfigTreeWatcher, 0, len(c.eventBus.watchers)-1)
															for _, watcher := range c.eventBus.watchers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:66
			_go_fuzz_dep_.CoverTab[193751]++
																if watcher.id != cw.id {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:67
				_go_fuzz_dep_.CoverTab[193752]++
																	watchers = append(watchers, watcher)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:68
				// _ = "end of CoverTab[193752]"

			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:70
				_go_fuzz_dep_.CoverTab[193753]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:70
				// _ = "end of CoverTab[193753]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:70
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:70
			// _ = "end of CoverTab[193751]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:71
		// _ = "end of CoverTab[193748]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:71
		_go_fuzz_dep_.CoverTab[193749]++
															c.eventBus.watchers = watchers
															c.eventBus.rm.Unlock()
															close(ch)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:74
		// _ = "end of CoverTab[193749]"

	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:76
	// _ = "end of CoverTab[193746]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:76
	_go_fuzz_dep_.CoverTab[193747]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:77
	// _ = "end of CoverTab[193747]"
}

// Put :
func (c *ConfigStore) Put(key string, entry Entry) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:81
	_go_fuzz_dep_.CoverTab[193754]++
														err := put(c.ConfigTree, key, entry)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:83
		_go_fuzz_dep_.CoverTab[193756]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:84
		// _ = "end of CoverTab[193756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:85
		_go_fuzz_dep_.CoverTab[193757]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:85
		// _ = "end of CoverTab[193757]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:85
	// _ = "end of CoverTab[193754]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:85
	_go_fuzz_dep_.CoverTab[193755]++
														c.eventBus.send(event.Event{
		Key:		key,
		Value:		entry.Value,
		EventType:	entry.EventType,
	})

														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:92
	// _ = "end of CoverTab[193755]"
}

// Get :
func (c *ConfigStore) Get(key string) (Entry, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:96
	_go_fuzz_dep_.CoverTab[193758]++
														node, err := get(c.ConfigTree, key)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:98
		_go_fuzz_dep_.CoverTab[193760]++
															return Entry{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:99
		// _ = "end of CoverTab[193760]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:100
		_go_fuzz_dep_.CoverTab[193761]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:100
		// _ = "end of CoverTab[193761]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:100
	// _ = "end of CoverTab[193758]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:100
	_go_fuzz_dep_.CoverTab[193759]++
														return Entry{
		Key:	key,
		Value:	node,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:104
	// _ = "end of CoverTab[193759]"
}

var _ Store = &ConfigStore{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:107
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/store.go:107
var _ = _go_fuzz_dep_.CoverTab
