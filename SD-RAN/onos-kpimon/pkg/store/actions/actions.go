// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-kpimon/pkg/store/watcher"
)

// Store kpm action definitions  store interface
type Store interface {
	Put(ctx context.Context, key Key, value interface{}) (*Entry, error)

	// Get gets a metric store entry based on a given key
	Get(ctx context.Context, key Key) (*Entry, error)
}

type store struct {
	actions  map[Key]*Entry
	mu       sync.RWMutex
	watchers *watcher.Watchers
}

// NewStore creates new store
func NewStore() Store {
	watchers := watcher.NewWatchers()
	return &store{
		actions:  make(map[Key]*Entry),
		watchers: watchers,
	}
}

func (s *store) Put(_ context.Context, key Key, value interface{}) (*Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	s.actions[key] = entry
	return entry, nil
}

func (s *store) Get(_ context.Context, key Key) (*Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.actions[key]; ok {
		return v, nil
	}
	return nil, errors.New(errors.NotFound, "the cell entry does not exist")
}

// NewKey creates a new key
func NewKey(subID SubscriptionID) Key {
	return Key{
		SubscriptionID: subID,
	}
}

var _ Store = &store{}
