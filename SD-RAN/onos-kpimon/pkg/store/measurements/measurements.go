// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package measurements

import (
	"context"
	"fmt"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/google/uuid"

	"github.com/onosproject/onos-kpimon/pkg/store/watcher"

	"github.com/onosproject/onos-kpimon/pkg/store/event"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

var log = logging.GetLogger()

// Store kpm metrics store interface
type Store interface {
	Put(ctx context.Context, key Key, value interface{}) (*Entry, error)

	// Get gets a metric store entry based on a given key
	Get(ctx context.Context, key Key) (*Entry, error)

	// Delete deletes an entry based on a given key
	Delete(ctx context.Context, key Key) error

	// Entries list all of the metric store entries
	Entries(ctx context.Context, ch chan<- *Entry) error

	// Watch measurement store changes
	Watch(ctx context.Context, ch chan<- event.Event) error
}

type store struct {
	measurements map[Key]*Entry
	mu           sync.RWMutex
	watchers     *watcher.Watchers
}

// NewStore creates new store
func NewStore() Store {
	watchers := watcher.NewWatchers()
	return &store{
		measurements: make(map[Key]*Entry),
		watchers:     watchers,
	}
}

func (s *store) Entries(_ context.Context, ch chan<- *Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.measurements) == 0 {
		close(ch)
		return fmt.Errorf("no measurements entries stored")
	}

	for _, entry := range s.measurements {
		ch <- entry
	}

	close(ch)
	return nil
}

func (s *store) Delete(_ context.Context, key Key) error {
	// TODO check the key and make sure it is not empty
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.measurements, key)
	return nil

}

func (s *store) Put(_ context.Context, key Key, value interface{}) (*Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	s.measurements[key] = entry
	s.watchers.Send(event.Event{
		Key:   key,
		Value: entry,
		Type:  Created,
	})
	return entry, nil

}

func (s *store) Get(_ context.Context, key Key) (*Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.measurements[key]; ok {
		return v, nil
	}
	return nil, errors.New(errors.NotFound, "the measurement entry does not exist")
}

func (s *store) Watch(ctx context.Context, ch chan<- event.Event) error {
	id := uuid.New()
	err := s.watchers.AddWatcher(id, ch)
	if err != nil {
		log.Error(err)
		close(ch)
		return err
	}
	go func() {
		<-ctx.Done()
		err = s.watchers.RemoveWatcher(id)
		if err != nil {
			log.Error(err)
		}
		close(ch)
	}()
	return nil
}

// NewKey creates a new measurements map key
func NewKey(CellID CellIdentity, nodeID string) Key {
	return Key{
		NodeID:       nodeID,
		CellIdentity: CellID,
	}
}

var _ Store = &store{}
