// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:5
)

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/google/uuid"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

var log = logging.GetLogger()

// Store mho store interface
type Store interface {
	Put(ctx context.Context, key string, value interface{}) (*Entry, error)

	// Get gets a store entry based on a given key
	Get(ctx context.Context, key string) (*Entry, error)

	// Delete deletes an entry based on a given key
	Delete(ctx context.Context, key string) error

	// Entries list all of the store entries
	Entries(ctx context.Context, ch chan<- *Entry) error

	// Watch store changes
	Watch(ctx context.Context, ch chan<- Event) error
}

type store struct {
	records		map[string]*Entry
	mu		sync.RWMutex
	watchers	*Watchers
}

// NewStore creates new store
func NewStore() Store {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:44
	_go_fuzz_dep_.CoverTab[179553]++
												watchers := NewWatchers()
												return &store{
		records:	make(map[string]*Entry),
		watchers:	watchers,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:49
	// _ = "end of CoverTab[179553]"
}

func (s *store) Entries(ctx context.Context, ch chan<- *Entry) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:52
	_go_fuzz_dep_.CoverTab[179554]++
												s.mu.Lock()
												defer s.mu.Unlock()

												for _, entry := range s.records {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:56
		_go_fuzz_dep_.CoverTab[179556]++
													ch <- entry
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:57
		// _ = "end of CoverTab[179556]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:58
	// _ = "end of CoverTab[179554]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:58
	_go_fuzz_dep_.CoverTab[179555]++

												close(ch)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:61
	// _ = "end of CoverTab[179555]"
}

func (s *store) Delete(ctx context.Context, key string) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:64
	_go_fuzz_dep_.CoverTab[179557]++

												s.mu.Lock()
												defer s.mu.Unlock()
												delete(s.records, key)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:69
	// _ = "end of CoverTab[179557]"

}

func (s *store) Put(ctx context.Context, key string, value interface{}) (*Entry, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:73
	_go_fuzz_dep_.CoverTab[179558]++
												s.mu.Lock()
												defer s.mu.Unlock()
												entry := &Entry{
		Key:	key,
		Value:	value,
	}
	s.records[key] = entry
	s.watchers.Send(Event{
		Key:	key,
		Value:	entry,
		Type:	Created,
	})
												return entry, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:86
	// _ = "end of CoverTab[179558]"

}

func (s *store) Get(ctx context.Context, key string) (*Entry, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:90
	_go_fuzz_dep_.CoverTab[179559]++
												s.mu.Lock()
												defer s.mu.Unlock()
												if v, ok := s.records[key]; ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:93
		_go_fuzz_dep_.CoverTab[179561]++
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:94
		// _ = "end of CoverTab[179561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:95
		_go_fuzz_dep_.CoverTab[179562]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:95
		// _ = "end of CoverTab[179562]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:95
	// _ = "end of CoverTab[179559]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:95
	_go_fuzz_dep_.CoverTab[179560]++
												return nil, errors.New(errors.NotFound, "the entry does not exist")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:96
	// _ = "end of CoverTab[179560]"
}

func (s *store) Watch(ctx context.Context, ch chan<- Event) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:99
	_go_fuzz_dep_.CoverTab[179563]++
												id := uuid.New()
												err := s.watchers.AddWatcher(id, ch)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:102
		_go_fuzz_dep_.CoverTab[179566]++
														log.Error(err)
														close(ch)
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:105
		// _ = "end of CoverTab[179566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
		_go_fuzz_dep_.CoverTab[179567]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
		// _ = "end of CoverTab[179567]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
	// _ = "end of CoverTab[179563]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
	_go_fuzz_dep_.CoverTab[179564]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
	_curRoutineNum157_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:106
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum157_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
		_go_fuzz_dep_.CoverTab[179568]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
			_go_fuzz_dep_.CoverTab[179570]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum157_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
			// _ = "end of CoverTab[179570]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:107
		}()
														<-ctx.Done()
														err = s.watchers.RemoveWatcher(id)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:110
			_go_fuzz_dep_.CoverTab[179571]++
															log.Error(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:111
			// _ = "end of CoverTab[179571]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:112
			_go_fuzz_dep_.CoverTab[179572]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:112
			// _ = "end of CoverTab[179572]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:112
		// _ = "end of CoverTab[179568]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:112
		_go_fuzz_dep_.CoverTab[179569]++
														close(ch)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:113
		// _ = "end of CoverTab[179569]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:114
	// _ = "end of CoverTab[179564]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:114
	_go_fuzz_dep_.CoverTab[179565]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:115
	// _ = "end of CoverTab[179565]"
}

var _ Store = &store{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/store.go:118
var _ = _go_fuzz_dep_.CoverTab
