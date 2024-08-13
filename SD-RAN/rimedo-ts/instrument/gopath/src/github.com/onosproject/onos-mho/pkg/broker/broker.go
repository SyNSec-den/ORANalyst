// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
package broker

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:5
)

import (
	"context"
	"io"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

var log = logging.GetLogger()

// NewBroker creates a new subscription stream broker
func NewBroker() Broker {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:21
	_go_fuzz_dep_.CoverTab[196629]++
													return &streamBroker{
		subs:		make(map[e2api.ChannelID]Stream),
		streams:	make(map[StreamID]Stream),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:25
	// _ = "end of CoverTab[196629]"
}

// Broker is a subscription stream broker
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:28
// The Broker is responsible for managing Streams for propagating indications from the southbound API
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:28
// to the northbound API.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:31
type Broker interface {
	io.Closer

	// OpenReader opens a subscription Stream
	// If a stream already exists for the subscription, the existing stream will be returned.
	// If no stream exists, a new stream will be allocated with a unique StreamID.
	OpenReader(ctx context.Context, node e2client.Node,
		subName string, id e2api.ChannelID, subSpec e2api.SubscriptionSpec) (StreamReader, error)

	// CloseStream closes a subscription Stream
	// The associated Stream will be closed gracefully: the reader will continue receiving pending indications
	// until the buffer is empty.
	CloseStream(ctx context.Context, id e2api.ChannelID) (StreamReader, error)

	// GetWriter gets a write stream by its StreamID
	// If no Stream exists for the given StreamID, a NotFound error will be returned.
	GetWriter(id StreamID) (StreamWriter, error)

	// ChannelIDs get all of subscription channel IDs
	ChannelIDs() []e2api.ChannelID
}

type streamBroker struct {
	subs		map[e2api.ChannelID]Stream
	streams		map[StreamID]Stream
	streamID	StreamID
	mu		sync.RWMutex
}

func (b *streamBroker) ChannelIDs() []e2api.ChannelID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:60
	_go_fuzz_dep_.CoverTab[196630]++
													b.mu.Lock()
													defer b.mu.Unlock()
													channelIDs := make([]e2api.ChannelID, len(b.subs))
													for channelID := range b.subs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:64
		_go_fuzz_dep_.CoverTab[196632]++
														channelIDs = append(channelIDs, channelID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:65
		// _ = "end of CoverTab[196632]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:66
	// _ = "end of CoverTab[196630]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:66
	_go_fuzz_dep_.CoverTab[196631]++
													return channelIDs
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:67
	// _ = "end of CoverTab[196631]"
}

func (b *streamBroker) OpenReader(ctx context.Context, node e2client.Node, subName string, channelID e2api.ChannelID, subSpec e2api.SubscriptionSpec) (StreamReader, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:70
	_go_fuzz_dep_.CoverTab[196633]++
													b.mu.RLock()
													stream, ok := b.subs[channelID]
													b.mu.RUnlock()
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:74
		_go_fuzz_dep_.CoverTab[196635]++
														return stream, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:75
		// _ = "end of CoverTab[196635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:76
		_go_fuzz_dep_.CoverTab[196636]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:76
		// _ = "end of CoverTab[196636]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:76
	// _ = "end of CoverTab[196633]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:76
	_go_fuzz_dep_.CoverTab[196634]++

													b.mu.Lock()
													defer b.mu.Unlock()

													b.streamID++
													streamID := b.streamID
													stream = newBufferedStream(node, subName, streamID, channelID, subSpec)
													b.subs[channelID] = stream
													b.streams[streamID] = stream
													log.Infof("Opened new stream %d for subscription channel '%s'", streamID, channelID)
													return stream, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:87
	// _ = "end of CoverTab[196634]"
}

func (b *streamBroker) CloseStream(ctx context.Context, id e2api.ChannelID) (StreamReader, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:90
	_go_fuzz_dep_.CoverTab[196637]++
													b.mu.Lock()
													defer b.mu.Unlock()
													stream, ok := b.subs[id]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:94
		_go_fuzz_dep_.CoverTab[196640]++
														return nil, errors.NewNotFound("subscription '%s' not found", id)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:95
		// _ = "end of CoverTab[196640]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:96
		_go_fuzz_dep_.CoverTab[196641]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:96
		// _ = "end of CoverTab[196641]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:96
	// _ = "end of CoverTab[196637]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:96
	_go_fuzz_dep_.CoverTab[196638]++

													log.Debugf("Deleting Subscription: %s", stream.SubscriptionName())
													err := stream.Node().Unsubscribe(ctx, stream.SubscriptionName())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:100
		_go_fuzz_dep_.CoverTab[196642]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:101
		// _ = "end of CoverTab[196642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:102
		_go_fuzz_dep_.CoverTab[196643]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:102
		// _ = "end of CoverTab[196643]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:102
	// _ = "end of CoverTab[196638]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:102
	_go_fuzz_dep_.CoverTab[196639]++

													delete(b.subs, stream.ChannelID())
													delete(b.streams, stream.StreamID())

													log.Infof("Closed stream %d for subscription '%s'", stream.StreamID(), id)
													return stream, stream.Close()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:108
	// _ = "end of CoverTab[196639]"
}

func (b *streamBroker) GetWriter(id StreamID) (StreamWriter, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:111
	_go_fuzz_dep_.CoverTab[196644]++
													b.mu.RLock()
													defer b.mu.RUnlock()
													stream, ok := b.streams[id]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:115
		_go_fuzz_dep_.CoverTab[196646]++
														return nil, errors.NewNotFound("stream %d not found", id)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:116
		// _ = "end of CoverTab[196646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:117
		_go_fuzz_dep_.CoverTab[196647]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:117
		// _ = "end of CoverTab[196647]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:117
	// _ = "end of CoverTab[196644]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:117
	_go_fuzz_dep_.CoverTab[196645]++
													return stream, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:118
	// _ = "end of CoverTab[196645]"
}

func (b *streamBroker) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:121
	_go_fuzz_dep_.CoverTab[196648]++
													b.mu.Lock()
													defer b.mu.Unlock()
													var err error
													for _, stream := range b.streams {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:125
		_go_fuzz_dep_.CoverTab[196650]++
														if e := stream.Close(); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:126
			_go_fuzz_dep_.CoverTab[196651]++
															err = e
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:127
			// _ = "end of CoverTab[196651]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:128
			_go_fuzz_dep_.CoverTab[196652]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:128
			// _ = "end of CoverTab[196652]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:128
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:128
		// _ = "end of CoverTab[196650]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:129
	// _ = "end of CoverTab[196648]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:129
	_go_fuzz_dep_.CoverTab[196649]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:130
	// _ = "end of CoverTab[196649]"
}

var _ Broker = &streamBroker{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:133
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/broker.go:133
var _ = _go_fuzz_dep_.CoverTab
