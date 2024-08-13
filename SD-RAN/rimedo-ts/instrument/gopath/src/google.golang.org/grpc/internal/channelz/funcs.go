//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:19
// Package channelz defines APIs for enabling channelz service, entry
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:19
// registration/deletion, and accessing channelz data. It also defines channelz
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:19
// metric struct formats.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:24
)

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/grpclog"
)

const (
	defaultMaxTraceEntry int32 = 30
)

var (
	db	dbWrapper
	idGen	idGenerator
	// EntryPerPage defines the number of channelz entries to be shown on a web page.
	EntryPerPage	= int64(50)
	curState	int32
	maxTraceEntry	= defaultMaxTraceEntry
)

// TurnOn turns on channelz data collection.
func TurnOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:52
	_go_fuzz_dep_.CoverTab[62574]++
												if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:53
		_go_fuzz_dep_.CoverTab[62575]++
													db.set(newChannelMap())
													idGen.reset()
													atomic.StoreInt32(&curState, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:56
		// _ = "end of CoverTab[62575]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:57
		_go_fuzz_dep_.CoverTab[62576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:57
		// _ = "end of CoverTab[62576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:57
	// _ = "end of CoverTab[62574]"
}

// IsOn returns whether channelz data collection is on.
func IsOn() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:61
	_go_fuzz_dep_.CoverTab[62577]++
												return atomic.CompareAndSwapInt32(&curState, 1, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:62
	// _ = "end of CoverTab[62577]"
}

// SetMaxTraceEntry sets maximum number of trace entry per entity (i.e. channel/subchannel).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:65
// Setting it to 0 will disable channel tracing.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:67
func SetMaxTraceEntry(i int32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:67
	_go_fuzz_dep_.CoverTab[62578]++
												atomic.StoreInt32(&maxTraceEntry, i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:68
	// _ = "end of CoverTab[62578]"
}

// ResetMaxTraceEntryToDefault resets the maximum number of trace entry per entity to default.
func ResetMaxTraceEntryToDefault() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:72
	_go_fuzz_dep_.CoverTab[62579]++
												atomic.StoreInt32(&maxTraceEntry, defaultMaxTraceEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:73
	// _ = "end of CoverTab[62579]"
}

func getMaxTraceEntry() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:76
	_go_fuzz_dep_.CoverTab[62580]++
												i := atomic.LoadInt32(&maxTraceEntry)
												return int(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:78
	// _ = "end of CoverTab[62580]"
}

// dbWarpper wraps around a reference to internal channelz data storage, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:81
// provide synchronized functionality to set and get the reference.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:83
type dbWrapper struct {
	mu	sync.RWMutex
	DB	*channelMap
}

func (d *dbWrapper) set(db *channelMap) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:88
	_go_fuzz_dep_.CoverTab[62581]++
												d.mu.Lock()
												d.DB = db
												d.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:91
	// _ = "end of CoverTab[62581]"
}

func (d *dbWrapper) get() *channelMap {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:94
	_go_fuzz_dep_.CoverTab[62582]++
												d.mu.RLock()
												defer d.mu.RUnlock()
												return d.DB
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:97
	// _ = "end of CoverTab[62582]"
}

// NewChannelzStorageForTesting initializes channelz data storage and id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
// generator for testing purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
// Returns a cleanup function to be invoked by the test, which waits for up to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
// 10s for all channelz state to be reset by the grpc goroutines when those
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
// entities get closed. This cleanup function helps with ensuring that tests
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:100
// don't mess up each other.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:107
func NewChannelzStorageForTesting() (cleanup func() error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:107
	_go_fuzz_dep_.CoverTab[62583]++
													db.set(newChannelMap())
													idGen.reset()

													return func() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:111
		_go_fuzz_dep_.CoverTab[62584]++
														cm := db.get()
														if cm == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:113
			_go_fuzz_dep_.CoverTab[62586]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:114
			// _ = "end of CoverTab[62586]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:115
			_go_fuzz_dep_.CoverTab[62587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:115
			// _ = "end of CoverTab[62587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:115
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:115
		// _ = "end of CoverTab[62584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:115
		_go_fuzz_dep_.CoverTab[62585]++

														ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
														defer cancel()
														ticker := time.NewTicker(10 * time.Millisecond)
														defer ticker.Stop()
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:121
			_go_fuzz_dep_.CoverTab[62588]++
															cm.mu.RLock()
															topLevelChannels, servers, channels, subChannels, listenSockets, normalSockets := len(cm.topLevelChannels), len(cm.servers), len(cm.channels), len(cm.subChannels), len(cm.listenSockets), len(cm.normalSockets)
															cm.mu.RUnlock()

															if err := ctx.Err(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:126
				_go_fuzz_dep_.CoverTab[62591]++
																return fmt.Errorf("after 10s the channelz map has not been cleaned up yet, topchannels: %d, servers: %d, channels: %d, subchannels: %d, listen sockets: %d, normal sockets: %d", topLevelChannels, servers, channels, subChannels, listenSockets, normalSockets)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:127
				// _ = "end of CoverTab[62591]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:128
				_go_fuzz_dep_.CoverTab[62592]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:128
				// _ = "end of CoverTab[62592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:128
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:128
			// _ = "end of CoverTab[62588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:128
			_go_fuzz_dep_.CoverTab[62589]++
															if topLevelChannels == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				return servers == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				// _ = "end of CoverTab[62593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62594]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				return channels == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				// _ = "end of CoverTab[62594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				return subChannels == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				// _ = "end of CoverTab[62595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				return listenSockets == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				// _ = "end of CoverTab[62596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				return normalSockets == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				// _ = "end of CoverTab[62597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:129
				_go_fuzz_dep_.CoverTab[62598]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:130
				// _ = "end of CoverTab[62598]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:131
				_go_fuzz_dep_.CoverTab[62599]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:131
				// _ = "end of CoverTab[62599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:131
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:131
			// _ = "end of CoverTab[62589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:131
			_go_fuzz_dep_.CoverTab[62590]++
															<-ticker.C
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:132
			// _ = "end of CoverTab[62590]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:133
		// _ = "end of CoverTab[62585]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:134
	// _ = "end of CoverTab[62583]"
}

// GetTopChannels returns a slice of top channel's ChannelMetric, along with a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:137
// boolean indicating whether there's more top channels to be queried for.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:137
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:137
// The arg id specifies that only top channel with id at or above it will be included
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:137
// in the result. The returned slice is up to a length of the arg maxResults or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:137
// EntryPerPage if maxResults is zero, and is sorted in ascending id order.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:143
func GetTopChannels(id int64, maxResults int64) ([]*ChannelMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:143
	_go_fuzz_dep_.CoverTab[62600]++
													return db.get().GetTopChannels(id, maxResults)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:144
	// _ = "end of CoverTab[62600]"
}

// GetServers returns a slice of server's ServerMetric, along with a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:147
// boolean indicating whether there's more servers to be queried for.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:147
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:147
// The arg id specifies that only server with id at or above it will be included
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:147
// in the result. The returned slice is up to a length of the arg maxResults or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:147
// EntryPerPage if maxResults is zero, and is sorted in ascending id order.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:153
func GetServers(id int64, maxResults int64) ([]*ServerMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:153
	_go_fuzz_dep_.CoverTab[62601]++
													return db.get().GetServers(id, maxResults)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:154
	// _ = "end of CoverTab[62601]"
}

// GetServerSockets returns a slice of server's (identified by id) normal socket's
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
// SocketMetric, along with a boolean indicating whether there's more sockets to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
// be queried for.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
// The arg startID specifies that only sockets with id at or above it will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
// included in the result. The returned slice is up to a length of the arg maxResults
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:157
// or EntryPerPage if maxResults is zero, and is sorted in ascending id order.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:164
func GetServerSockets(id int64, startID int64, maxResults int64) ([]*SocketMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:164
	_go_fuzz_dep_.CoverTab[62602]++
													return db.get().GetServerSockets(id, startID, maxResults)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:165
	// _ = "end of CoverTab[62602]"
}

// GetChannel returns the ChannelMetric for the channel (identified by id).
func GetChannel(id int64) *ChannelMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:169
	_go_fuzz_dep_.CoverTab[62603]++
													return db.get().GetChannel(id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:170
	// _ = "end of CoverTab[62603]"
}

// GetSubChannel returns the SubChannelMetric for the subchannel (identified by id).
func GetSubChannel(id int64) *SubChannelMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:174
	_go_fuzz_dep_.CoverTab[62604]++
													return db.get().GetSubChannel(id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:175
	// _ = "end of CoverTab[62604]"
}

// GetSocket returns the SocketInternalMetric for the socket (identified by id).
func GetSocket(id int64) *SocketMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:179
	_go_fuzz_dep_.CoverTab[62605]++
													return db.get().GetSocket(id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:180
	// _ = "end of CoverTab[62605]"
}

// GetServer returns the ServerMetric for the server (identified by id).
func GetServer(id int64) *ServerMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:184
	_go_fuzz_dep_.CoverTab[62606]++
													return db.get().GetServer(id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:185
	// _ = "end of CoverTab[62606]"
}

// RegisterChannel registers the given channel c in the channelz database with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
// ref as its reference name, and adds it to the child list of its parent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
// (identified by pid). pid == nil means no parent.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
// Returns a unique channelz identifier assigned to this channel.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:188
// If channelz is not turned ON, the channelz database is not mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:195
func RegisterChannel(c Channel, pid *Identifier, ref string) *Identifier {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:195
	_go_fuzz_dep_.CoverTab[62607]++
													id := idGen.genID()
													var parent int64
													isTopChannel := true
													if pid != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:199
		_go_fuzz_dep_.CoverTab[62610]++
														isTopChannel = false
														parent = pid.Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:201
		// _ = "end of CoverTab[62610]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:202
		_go_fuzz_dep_.CoverTab[62611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:202
		// _ = "end of CoverTab[62611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:202
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:202
	// _ = "end of CoverTab[62607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:202
	_go_fuzz_dep_.CoverTab[62608]++

													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:204
		_go_fuzz_dep_.CoverTab[62612]++
														return newIdentifer(RefChannel, id, pid)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:205
		// _ = "end of CoverTab[62612]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:206
		_go_fuzz_dep_.CoverTab[62613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:206
		// _ = "end of CoverTab[62613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:206
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:206
	// _ = "end of CoverTab[62608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:206
	_go_fuzz_dep_.CoverTab[62609]++

													cn := &channel{
		refName:	ref,
		c:		c,
		subChans:	make(map[int64]string),
		nestedChans:	make(map[int64]string),
		id:		id,
		pid:		parent,
		trace:		&channelTrace{createdTime: time.Now(), events: make([]*TraceEvent, 0, getMaxTraceEntry())},
	}
													db.get().addChannel(id, cn, isTopChannel, parent)
													return newIdentifer(RefChannel, id, pid)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:218
	// _ = "end of CoverTab[62609]"
}

// RegisterSubChannel registers the given subChannel c in the channelz database
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
// with ref as its reference name, and adds it to the child list of its parent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
// (identified by pid).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
// Returns a unique channelz identifier assigned to this subChannel.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:221
// If channelz is not turned ON, the channelz database is not mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:228
func RegisterSubChannel(c Channel, pid *Identifier, ref string) (*Identifier, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:228
	_go_fuzz_dep_.CoverTab[62614]++
													if pid == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:229
		_go_fuzz_dep_.CoverTab[62617]++
														return nil, errors.New("a SubChannel's parent id cannot be nil")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:230
		// _ = "end of CoverTab[62617]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:231
		_go_fuzz_dep_.CoverTab[62618]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:231
		// _ = "end of CoverTab[62618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:231
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:231
	// _ = "end of CoverTab[62614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:231
	_go_fuzz_dep_.CoverTab[62615]++
													id := idGen.genID()
													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:233
		_go_fuzz_dep_.CoverTab[62619]++
														return newIdentifer(RefSubChannel, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:234
		// _ = "end of CoverTab[62619]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:235
		_go_fuzz_dep_.CoverTab[62620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:235
		// _ = "end of CoverTab[62620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:235
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:235
	// _ = "end of CoverTab[62615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:235
	_go_fuzz_dep_.CoverTab[62616]++

													sc := &subChannel{
		refName:	ref,
		c:		c,
		sockets:	make(map[int64]string),
		id:		id,
		pid:		pid.Int(),
		trace:		&channelTrace{createdTime: time.Now(), events: make([]*TraceEvent, 0, getMaxTraceEntry())},
	}
													db.get().addSubChannel(id, sc, pid.Int())
													return newIdentifer(RefSubChannel, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:246
	// _ = "end of CoverTab[62616]"
}

// RegisterServer registers the given server s in channelz database. It returns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:249
// the unique channelz tracking id assigned to this server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:249
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:249
// If channelz is not turned ON, the channelz database is not mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:253
func RegisterServer(s Server, ref string) *Identifier {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:253
	_go_fuzz_dep_.CoverTab[62621]++
													id := idGen.genID()
													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:255
		_go_fuzz_dep_.CoverTab[62623]++
														return newIdentifer(RefServer, id, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:256
		// _ = "end of CoverTab[62623]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:257
		_go_fuzz_dep_.CoverTab[62624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:257
		// _ = "end of CoverTab[62624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:257
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:257
	// _ = "end of CoverTab[62621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:257
	_go_fuzz_dep_.CoverTab[62622]++

													svr := &server{
		refName:	ref,
		s:		s,
		sockets:	make(map[int64]string),
		listenSockets:	make(map[int64]string),
		id:		id,
	}
													db.get().addServer(id, svr)
													return newIdentifer(RefServer, id, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:267
	// _ = "end of CoverTab[62622]"
}

// RegisterListenSocket registers the given listen socket s in channelz database
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:270
// with ref as its reference name, and add it to the child list of its parent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:270
// (identified by pid). It returns the unique channelz tracking id assigned to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:270
// this listen socket.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:270
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:270
// If channelz is not turned ON, the channelz database is not mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:276
func RegisterListenSocket(s Socket, pid *Identifier, ref string) (*Identifier, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:276
	_go_fuzz_dep_.CoverTab[62625]++
													if pid == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:277
		_go_fuzz_dep_.CoverTab[62628]++
														return nil, errors.New("a ListenSocket's parent id cannot be 0")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:278
		// _ = "end of CoverTab[62628]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:279
		_go_fuzz_dep_.CoverTab[62629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:279
		// _ = "end of CoverTab[62629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:279
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:279
	// _ = "end of CoverTab[62625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:279
	_go_fuzz_dep_.CoverTab[62626]++
													id := idGen.genID()
													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:281
		_go_fuzz_dep_.CoverTab[62630]++
														return newIdentifer(RefListenSocket, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:282
		// _ = "end of CoverTab[62630]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:283
		_go_fuzz_dep_.CoverTab[62631]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:283
		// _ = "end of CoverTab[62631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:283
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:283
	// _ = "end of CoverTab[62626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:283
	_go_fuzz_dep_.CoverTab[62627]++

													ls := &listenSocket{refName: ref, s: s, id: id, pid: pid.Int()}
													db.get().addListenSocket(id, ls, pid.Int())
													return newIdentifer(RefListenSocket, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:287
	// _ = "end of CoverTab[62627]"
}

// RegisterNormalSocket registers the given normal socket s in channelz database
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:290
// with ref as its reference name, and adds it to the child list of its parent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:290
// (identified by pid). It returns the unique channelz tracking id assigned to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:290
// this normal socket.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:290
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:290
// If channelz is not turned ON, the channelz database is not mutated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:296
func RegisterNormalSocket(s Socket, pid *Identifier, ref string) (*Identifier, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:296
	_go_fuzz_dep_.CoverTab[62632]++
													if pid == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:297
		_go_fuzz_dep_.CoverTab[62635]++
														return nil, errors.New("a NormalSocket's parent id cannot be 0")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:298
		// _ = "end of CoverTab[62635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:299
		_go_fuzz_dep_.CoverTab[62636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:299
		// _ = "end of CoverTab[62636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:299
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:299
	// _ = "end of CoverTab[62632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:299
	_go_fuzz_dep_.CoverTab[62633]++
													id := idGen.genID()
													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:301
		_go_fuzz_dep_.CoverTab[62637]++
														return newIdentifer(RefNormalSocket, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:302
		// _ = "end of CoverTab[62637]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:303
		_go_fuzz_dep_.CoverTab[62638]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:303
		// _ = "end of CoverTab[62638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:303
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:303
	// _ = "end of CoverTab[62633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:303
	_go_fuzz_dep_.CoverTab[62634]++

													ns := &normalSocket{refName: ref, s: s, id: id, pid: pid.Int()}
													db.get().addNormalSocket(id, ns, pid.Int())
													return newIdentifer(RefNormalSocket, id, pid), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:307
	// _ = "end of CoverTab[62634]"
}

// RemoveEntry removes an entry with unique channelz tracking id to be id from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:310
// channelz database.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:310
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:310
// If channelz is not turned ON, this function is a no-op.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:314
func RemoveEntry(id *Identifier) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:314
	_go_fuzz_dep_.CoverTab[62639]++
													if !IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:315
		_go_fuzz_dep_.CoverTab[62641]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:316
		// _ = "end of CoverTab[62641]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:317
		_go_fuzz_dep_.CoverTab[62642]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:317
		// _ = "end of CoverTab[62642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:317
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:317
	// _ = "end of CoverTab[62639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:317
	_go_fuzz_dep_.CoverTab[62640]++
													db.get().removeEntry(id.Int())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:318
	// _ = "end of CoverTab[62640]"
}

// TraceEventDesc is what the caller of AddTraceEvent should provide to describe
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:321
// the event to be added to the channel trace.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:321
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:321
// The Parent field is optional. It is used for an event that will be recorded
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:321
// in the entity's parent trace.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:326
type TraceEventDesc struct {
	Desc		string
	Severity	Severity
	Parent		*TraceEventDesc
}

// AddTraceEvent adds trace related to the entity with specified id, using the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:332
// provided TraceEventDesc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:332
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:332
// If channelz is not turned ON, this will simply log the event descriptions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:336
func AddTraceEvent(l grpclog.DepthLoggerV2, id *Identifier, depth int, desc *TraceEventDesc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:336
	_go_fuzz_dep_.CoverTab[62643]++

													switch desc.Severity {
	case CtUnknown, CtInfo:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:339
		_go_fuzz_dep_.CoverTab[62646]++
														l.InfoDepth(depth+1, withParens(id)+desc.Desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:340
		// _ = "end of CoverTab[62646]"
	case CtWarning:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:341
		_go_fuzz_dep_.CoverTab[62647]++
														l.WarningDepth(depth+1, withParens(id)+desc.Desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:342
		// _ = "end of CoverTab[62647]"
	case CtError:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:343
		_go_fuzz_dep_.CoverTab[62648]++
														l.ErrorDepth(depth+1, withParens(id)+desc.Desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:344
		// _ = "end of CoverTab[62648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:344
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:344
		_go_fuzz_dep_.CoverTab[62649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:344
		// _ = "end of CoverTab[62649]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:345
	// _ = "end of CoverTab[62643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:345
	_go_fuzz_dep_.CoverTab[62644]++

													if getMaxTraceEntry() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:347
		_go_fuzz_dep_.CoverTab[62650]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:348
		// _ = "end of CoverTab[62650]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:349
		_go_fuzz_dep_.CoverTab[62651]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:349
		// _ = "end of CoverTab[62651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:349
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:349
	// _ = "end of CoverTab[62644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:349
	_go_fuzz_dep_.CoverTab[62645]++
													if IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:350
		_go_fuzz_dep_.CoverTab[62652]++
														db.get().traceEvent(id.Int(), desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:351
		// _ = "end of CoverTab[62652]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:352
		_go_fuzz_dep_.CoverTab[62653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:352
		// _ = "end of CoverTab[62653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:352
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:352
	// _ = "end of CoverTab[62645]"
}

// channelMap is the storage data structure for channelz.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:355
// Methods of channelMap can be divided in two two categories with respect to locking.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:355
// 1. Methods acquire the global lock.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:355
// 2. Methods that can only be called when global lock is held.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:355
// A second type of method need always to be called inside a first type of method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:360
type channelMap struct {
	mu			sync.RWMutex
	topLevelChannels	map[int64]struct{}
	servers			map[int64]*server
	channels		map[int64]*channel
	subChannels		map[int64]*subChannel
	listenSockets		map[int64]*listenSocket
	normalSockets		map[int64]*normalSocket
}

func newChannelMap() *channelMap {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:370
	_go_fuzz_dep_.CoverTab[62654]++
													return &channelMap{
		topLevelChannels:	make(map[int64]struct{}),
		channels:		make(map[int64]*channel),
		listenSockets:		make(map[int64]*listenSocket),
		normalSockets:		make(map[int64]*normalSocket),
		servers:		make(map[int64]*server),
		subChannels:		make(map[int64]*subChannel),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:378
	// _ = "end of CoverTab[62654]"
}

func (c *channelMap) addServer(id int64, s *server) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:381
	_go_fuzz_dep_.CoverTab[62655]++
													c.mu.Lock()
													s.cm = c
													c.servers[id] = s
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:385
	// _ = "end of CoverTab[62655]"
}

func (c *channelMap) addChannel(id int64, cn *channel, isTopChannel bool, pid int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:388
	_go_fuzz_dep_.CoverTab[62656]++
													c.mu.Lock()
													cn.cm = c
													cn.trace.cm = c
													c.channels[id] = cn
													if isTopChannel {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:393
		_go_fuzz_dep_.CoverTab[62658]++
														c.topLevelChannels[id] = struct{}{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:394
		// _ = "end of CoverTab[62658]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:395
		_go_fuzz_dep_.CoverTab[62659]++
														c.findEntry(pid).addChild(id, cn)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:396
		// _ = "end of CoverTab[62659]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:397
	// _ = "end of CoverTab[62656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:397
	_go_fuzz_dep_.CoverTab[62657]++
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:398
	// _ = "end of CoverTab[62657]"
}

func (c *channelMap) addSubChannel(id int64, sc *subChannel, pid int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:401
	_go_fuzz_dep_.CoverTab[62660]++
													c.mu.Lock()
													sc.cm = c
													sc.trace.cm = c
													c.subChannels[id] = sc
													c.findEntry(pid).addChild(id, sc)
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:407
	// _ = "end of CoverTab[62660]"
}

func (c *channelMap) addListenSocket(id int64, ls *listenSocket, pid int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:410
	_go_fuzz_dep_.CoverTab[62661]++
													c.mu.Lock()
													ls.cm = c
													c.listenSockets[id] = ls
													c.findEntry(pid).addChild(id, ls)
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:415
	// _ = "end of CoverTab[62661]"
}

func (c *channelMap) addNormalSocket(id int64, ns *normalSocket, pid int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:418
	_go_fuzz_dep_.CoverTab[62662]++
													c.mu.Lock()
													ns.cm = c
													c.normalSockets[id] = ns
													c.findEntry(pid).addChild(id, ns)
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:423
	// _ = "end of CoverTab[62662]"
}

// removeEntry triggers the removal of an entry, which may not indeed delete the entry, if it has to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:426
// wait on the deletion of its children and until no other entity's channel trace references it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:426
// It may lead to a chain of entry deletion. For example, deleting the last socket of a gracefully
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:426
// shutting down server will lead to the server being also deleted.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:430
func (c *channelMap) removeEntry(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:430
	_go_fuzz_dep_.CoverTab[62663]++
													c.mu.Lock()
													c.findEntry(id).triggerDelete()
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:433
	// _ = "end of CoverTab[62663]"
}

// c.mu must be held by the caller
func (c *channelMap) decrTraceRefCount(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:437
	_go_fuzz_dep_.CoverTab[62664]++
													e := c.findEntry(id)
													if v, ok := e.(tracedChannel); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:439
		_go_fuzz_dep_.CoverTab[62665]++
														v.decrTraceRefCount()
														e.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:441
		// _ = "end of CoverTab[62665]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:442
		_go_fuzz_dep_.CoverTab[62666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:442
		// _ = "end of CoverTab[62666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:442
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:442
	// _ = "end of CoverTab[62664]"
}

// c.mu must be held by the caller.
func (c *channelMap) findEntry(id int64) entry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:446
	_go_fuzz_dep_.CoverTab[62667]++
													var v entry
													var ok bool
													if v, ok = c.channels[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:449
		_go_fuzz_dep_.CoverTab[62673]++
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:450
		// _ = "end of CoverTab[62673]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:451
		_go_fuzz_dep_.CoverTab[62674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:451
		// _ = "end of CoverTab[62674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:451
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:451
	// _ = "end of CoverTab[62667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:451
	_go_fuzz_dep_.CoverTab[62668]++
													if v, ok = c.subChannels[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:452
		_go_fuzz_dep_.CoverTab[62675]++
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:453
		// _ = "end of CoverTab[62675]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:454
		_go_fuzz_dep_.CoverTab[62676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:454
		// _ = "end of CoverTab[62676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:454
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:454
	// _ = "end of CoverTab[62668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:454
	_go_fuzz_dep_.CoverTab[62669]++
													if v, ok = c.servers[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:455
		_go_fuzz_dep_.CoverTab[62677]++
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:456
		// _ = "end of CoverTab[62677]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:457
		_go_fuzz_dep_.CoverTab[62678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:457
		// _ = "end of CoverTab[62678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:457
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:457
	// _ = "end of CoverTab[62669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:457
	_go_fuzz_dep_.CoverTab[62670]++
													if v, ok = c.listenSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:458
		_go_fuzz_dep_.CoverTab[62679]++
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:459
		// _ = "end of CoverTab[62679]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:460
		_go_fuzz_dep_.CoverTab[62680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:460
		// _ = "end of CoverTab[62680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:460
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:460
	// _ = "end of CoverTab[62670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:460
	_go_fuzz_dep_.CoverTab[62671]++
													if v, ok = c.normalSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:461
		_go_fuzz_dep_.CoverTab[62681]++
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:462
		// _ = "end of CoverTab[62681]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:463
		_go_fuzz_dep_.CoverTab[62682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:463
		// _ = "end of CoverTab[62682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:463
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:463
	// _ = "end of CoverTab[62671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:463
	_go_fuzz_dep_.CoverTab[62672]++
													return &dummyEntry{idNotFound: id}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:464
	// _ = "end of CoverTab[62672]"
}

// c.mu must be held by the caller
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:467
// deleteEntry simply deletes an entry from the channelMap. Before calling this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:467
// method, caller must check this entry is ready to be deleted, i.e removeEntry()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:467
// has been called on it, and no children still exist.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:467
// Conditionals are ordered by the expected frequency of deletion of each entity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:467
// type, in order to optimize performance.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:473
func (c *channelMap) deleteEntry(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:473
	_go_fuzz_dep_.CoverTab[62683]++
													var ok bool
													if _, ok = c.normalSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:475
		_go_fuzz_dep_.CoverTab[62688]++
														delete(c.normalSockets, id)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:477
		// _ = "end of CoverTab[62688]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:478
		_go_fuzz_dep_.CoverTab[62689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:478
		// _ = "end of CoverTab[62689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:478
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:478
	// _ = "end of CoverTab[62683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:478
	_go_fuzz_dep_.CoverTab[62684]++
													if _, ok = c.subChannels[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:479
		_go_fuzz_dep_.CoverTab[62690]++
														delete(c.subChannels, id)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:481
		// _ = "end of CoverTab[62690]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:482
		_go_fuzz_dep_.CoverTab[62691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:482
		// _ = "end of CoverTab[62691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:482
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:482
	// _ = "end of CoverTab[62684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:482
	_go_fuzz_dep_.CoverTab[62685]++
													if _, ok = c.channels[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:483
		_go_fuzz_dep_.CoverTab[62692]++
														delete(c.channels, id)
														delete(c.topLevelChannels, id)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:486
		// _ = "end of CoverTab[62692]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:487
		_go_fuzz_dep_.CoverTab[62693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:487
		// _ = "end of CoverTab[62693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:487
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:487
	// _ = "end of CoverTab[62685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:487
	_go_fuzz_dep_.CoverTab[62686]++
													if _, ok = c.listenSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:488
		_go_fuzz_dep_.CoverTab[62694]++
														delete(c.listenSockets, id)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:490
		// _ = "end of CoverTab[62694]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:491
		_go_fuzz_dep_.CoverTab[62695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:491
		// _ = "end of CoverTab[62695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:491
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:491
	// _ = "end of CoverTab[62686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:491
	_go_fuzz_dep_.CoverTab[62687]++
													if _, ok = c.servers[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:492
		_go_fuzz_dep_.CoverTab[62696]++
														delete(c.servers, id)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:494
		// _ = "end of CoverTab[62696]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:495
		_go_fuzz_dep_.CoverTab[62697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:495
		// _ = "end of CoverTab[62697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:495
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:495
	// _ = "end of CoverTab[62687]"
}

func (c *channelMap) traceEvent(id int64, desc *TraceEventDesc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:498
	_go_fuzz_dep_.CoverTab[62698]++
													c.mu.Lock()
													child := c.findEntry(id)
													childTC, ok := child.(tracedChannel)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:502
		_go_fuzz_dep_.CoverTab[62701]++
														c.mu.Unlock()
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:504
		// _ = "end of CoverTab[62701]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:505
		_go_fuzz_dep_.CoverTab[62702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:505
		// _ = "end of CoverTab[62702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:505
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:505
	// _ = "end of CoverTab[62698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:505
	_go_fuzz_dep_.CoverTab[62699]++
													childTC.getChannelTrace().append(&TraceEvent{Desc: desc.Desc, Severity: desc.Severity, Timestamp: time.Now()})
													if desc.Parent != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:507
		_go_fuzz_dep_.CoverTab[62703]++
														parent := c.findEntry(child.getParentID())
														var chanType RefChannelType
														switch child.(type) {
		case *channel:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:511
			_go_fuzz_dep_.CoverTab[62705]++
															chanType = RefChannel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:512
			// _ = "end of CoverTab[62705]"
		case *subChannel:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:513
			_go_fuzz_dep_.CoverTab[62706]++
															chanType = RefSubChannel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:514
			// _ = "end of CoverTab[62706]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:515
		// _ = "end of CoverTab[62703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:515
		_go_fuzz_dep_.CoverTab[62704]++
														if parentTC, ok := parent.(tracedChannel); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:516
			_go_fuzz_dep_.CoverTab[62707]++
															parentTC.getChannelTrace().append(&TraceEvent{
				Desc:		desc.Parent.Desc,
				Severity:	desc.Parent.Severity,
				Timestamp:	time.Now(),
				RefID:		id,
				RefName:	childTC.getRefName(),
				RefType:	chanType,
			})
															childTC.incrTraceRefCount()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:525
			// _ = "end of CoverTab[62707]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:526
			_go_fuzz_dep_.CoverTab[62708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:526
			// _ = "end of CoverTab[62708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:526
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:526
		// _ = "end of CoverTab[62704]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:527
		_go_fuzz_dep_.CoverTab[62709]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:527
		// _ = "end of CoverTab[62709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:527
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:527
	// _ = "end of CoverTab[62699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:527
	_go_fuzz_dep_.CoverTab[62700]++
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:528
	// _ = "end of CoverTab[62700]"
}

type int64Slice []int64

func (s int64Slice) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:533
	_go_fuzz_dep_.CoverTab[62710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:533
	return len(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:533
	// _ = "end of CoverTab[62710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:533
}
func (s int64Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:534
	_go_fuzz_dep_.CoverTab[62711]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:534
	s[i], s[j] = s[j], s[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:534
	// _ = "end of CoverTab[62711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:534
}
func (s int64Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:535
	_go_fuzz_dep_.CoverTab[62712]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:535
	return s[i] < s[j]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:535
	// _ = "end of CoverTab[62712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:535
}

func copyMap(m map[int64]string) map[int64]string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:537
	_go_fuzz_dep_.CoverTab[62713]++
													n := make(map[int64]string)
													for k, v := range m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:539
		_go_fuzz_dep_.CoverTab[62715]++
														n[k] = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:540
		// _ = "end of CoverTab[62715]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:541
	// _ = "end of CoverTab[62713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:541
	_go_fuzz_dep_.CoverTab[62714]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:542
	// _ = "end of CoverTab[62714]"
}

func min(a, b int64) int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:545
	_go_fuzz_dep_.CoverTab[62716]++
													if a < b {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:546
		_go_fuzz_dep_.CoverTab[62718]++
														return a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:547
		// _ = "end of CoverTab[62718]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:548
		_go_fuzz_dep_.CoverTab[62719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:548
		// _ = "end of CoverTab[62719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:548
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:548
	// _ = "end of CoverTab[62716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:548
	_go_fuzz_dep_.CoverTab[62717]++
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:549
	// _ = "end of CoverTab[62717]"
}

func (c *channelMap) GetTopChannels(id int64, maxResults int64) ([]*ChannelMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:552
	_go_fuzz_dep_.CoverTab[62720]++
													if maxResults <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:553
		_go_fuzz_dep_.CoverTab[62727]++
														maxResults = EntryPerPage
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:554
		// _ = "end of CoverTab[62727]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:555
		_go_fuzz_dep_.CoverTab[62728]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:555
		// _ = "end of CoverTab[62728]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:555
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:555
	// _ = "end of CoverTab[62720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:555
	_go_fuzz_dep_.CoverTab[62721]++
													c.mu.RLock()
													l := int64(len(c.topLevelChannels))
													ids := make([]int64, 0, l)
													cns := make([]*channel, 0, min(l, maxResults))

													for k := range c.topLevelChannels {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:561
		_go_fuzz_dep_.CoverTab[62729]++
														ids = append(ids, k)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:562
		// _ = "end of CoverTab[62729]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:563
	// _ = "end of CoverTab[62721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:563
	_go_fuzz_dep_.CoverTab[62722]++
													sort.Sort(int64Slice(ids))
													idx := sort.Search(len(ids), func(i int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
		_go_fuzz_dep_.CoverTab[62730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
		return ids[i] >= id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
		// _ = "end of CoverTab[62730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
	// _ = "end of CoverTab[62722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:565
	_go_fuzz_dep_.CoverTab[62723]++
													count := int64(0)
													var end bool
													var t []*ChannelMetric
													for i, v := range ids[idx:] {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:569
		_go_fuzz_dep_.CoverTab[62731]++
														if count == maxResults {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:570
			_go_fuzz_dep_.CoverTab[62734]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:571
			// _ = "end of CoverTab[62734]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:572
			_go_fuzz_dep_.CoverTab[62735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:572
			// _ = "end of CoverTab[62735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:572
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:572
		// _ = "end of CoverTab[62731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:572
		_go_fuzz_dep_.CoverTab[62732]++
														if cn, ok := c.channels[v]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:573
			_go_fuzz_dep_.CoverTab[62736]++
															cns = append(cns, cn)
															t = append(t, &ChannelMetric{
				NestedChans:	copyMap(cn.nestedChans),
				SubChans:	copyMap(cn.subChans),
			})
															count++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:579
			// _ = "end of CoverTab[62736]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:580
			_go_fuzz_dep_.CoverTab[62737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:580
			// _ = "end of CoverTab[62737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:580
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:580
		// _ = "end of CoverTab[62732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:580
		_go_fuzz_dep_.CoverTab[62733]++
														if i == len(ids[idx:])-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:581
			_go_fuzz_dep_.CoverTab[62738]++
															end = true
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:583
			// _ = "end of CoverTab[62738]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:584
			_go_fuzz_dep_.CoverTab[62739]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:584
			// _ = "end of CoverTab[62739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:584
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:584
		// _ = "end of CoverTab[62733]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:585
	// _ = "end of CoverTab[62723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:585
	_go_fuzz_dep_.CoverTab[62724]++
													c.mu.RUnlock()
													if count == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:587
		_go_fuzz_dep_.CoverTab[62740]++
														end = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:588
		// _ = "end of CoverTab[62740]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:589
		_go_fuzz_dep_.CoverTab[62741]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:589
		// _ = "end of CoverTab[62741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:589
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:589
	// _ = "end of CoverTab[62724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:589
	_go_fuzz_dep_.CoverTab[62725]++

													for i, cn := range cns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:591
		_go_fuzz_dep_.CoverTab[62742]++
														t[i].ChannelData = cn.c.ChannelzMetric()
														t[i].ID = cn.id
														t[i].RefName = cn.refName
														t[i].Trace = cn.trace.dumpData()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:595
		// _ = "end of CoverTab[62742]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:596
	// _ = "end of CoverTab[62725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:596
	_go_fuzz_dep_.CoverTab[62726]++
													return t, end
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:597
	// _ = "end of CoverTab[62726]"
}

func (c *channelMap) GetServers(id, maxResults int64) ([]*ServerMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:600
	_go_fuzz_dep_.CoverTab[62743]++
													if maxResults <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:601
		_go_fuzz_dep_.CoverTab[62750]++
														maxResults = EntryPerPage
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:602
		// _ = "end of CoverTab[62750]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:603
		_go_fuzz_dep_.CoverTab[62751]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:603
		// _ = "end of CoverTab[62751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:603
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:603
	// _ = "end of CoverTab[62743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:603
	_go_fuzz_dep_.CoverTab[62744]++
													c.mu.RLock()
													l := int64(len(c.servers))
													ids := make([]int64, 0, l)
													ss := make([]*server, 0, min(l, maxResults))
													for k := range c.servers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:608
		_go_fuzz_dep_.CoverTab[62752]++
														ids = append(ids, k)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:609
		// _ = "end of CoverTab[62752]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:610
	// _ = "end of CoverTab[62744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:610
	_go_fuzz_dep_.CoverTab[62745]++
													sort.Sort(int64Slice(ids))
													idx := sort.Search(len(ids), func(i int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
		_go_fuzz_dep_.CoverTab[62753]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
		return ids[i] >= id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
		// _ = "end of CoverTab[62753]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
	// _ = "end of CoverTab[62745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:612
	_go_fuzz_dep_.CoverTab[62746]++
													count := int64(0)
													var end bool
													var s []*ServerMetric
													for i, v := range ids[idx:] {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:616
		_go_fuzz_dep_.CoverTab[62754]++
														if count == maxResults {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:617
			_go_fuzz_dep_.CoverTab[62757]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:618
			// _ = "end of CoverTab[62757]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:619
			_go_fuzz_dep_.CoverTab[62758]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:619
			// _ = "end of CoverTab[62758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:619
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:619
		// _ = "end of CoverTab[62754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:619
		_go_fuzz_dep_.CoverTab[62755]++
														if svr, ok := c.servers[v]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:620
			_go_fuzz_dep_.CoverTab[62759]++
															ss = append(ss, svr)
															s = append(s, &ServerMetric{
				ListenSockets: copyMap(svr.listenSockets),
			})
															count++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:625
			// _ = "end of CoverTab[62759]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:626
			_go_fuzz_dep_.CoverTab[62760]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:626
			// _ = "end of CoverTab[62760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:626
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:626
		// _ = "end of CoverTab[62755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:626
		_go_fuzz_dep_.CoverTab[62756]++
														if i == len(ids[idx:])-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:627
			_go_fuzz_dep_.CoverTab[62761]++
															end = true
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:629
			// _ = "end of CoverTab[62761]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:630
			_go_fuzz_dep_.CoverTab[62762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:630
			// _ = "end of CoverTab[62762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:630
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:630
		// _ = "end of CoverTab[62756]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:631
	// _ = "end of CoverTab[62746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:631
	_go_fuzz_dep_.CoverTab[62747]++
													c.mu.RUnlock()
													if count == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:633
		_go_fuzz_dep_.CoverTab[62763]++
														end = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:634
		// _ = "end of CoverTab[62763]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:635
		_go_fuzz_dep_.CoverTab[62764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:635
		// _ = "end of CoverTab[62764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:635
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:635
	// _ = "end of CoverTab[62747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:635
	_go_fuzz_dep_.CoverTab[62748]++

													for i, svr := range ss {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:637
		_go_fuzz_dep_.CoverTab[62765]++
														s[i].ServerData = svr.s.ChannelzMetric()
														s[i].ID = svr.id
														s[i].RefName = svr.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:640
		// _ = "end of CoverTab[62765]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:641
	// _ = "end of CoverTab[62748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:641
	_go_fuzz_dep_.CoverTab[62749]++
													return s, end
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:642
	// _ = "end of CoverTab[62749]"
}

func (c *channelMap) GetServerSockets(id int64, startID int64, maxResults int64) ([]*SocketMetric, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:645
	_go_fuzz_dep_.CoverTab[62766]++
													if maxResults <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:646
		_go_fuzz_dep_.CoverTab[62774]++
														maxResults = EntryPerPage
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:647
		// _ = "end of CoverTab[62774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:648
		_go_fuzz_dep_.CoverTab[62775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:648
		// _ = "end of CoverTab[62775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:648
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:648
	// _ = "end of CoverTab[62766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:648
	_go_fuzz_dep_.CoverTab[62767]++
													var svr *server
													var ok bool
													c.mu.RLock()
													if svr, ok = c.servers[id]; !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:652
		_go_fuzz_dep_.CoverTab[62776]++

														c.mu.RUnlock()
														return nil, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:655
		// _ = "end of CoverTab[62776]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:656
		_go_fuzz_dep_.CoverTab[62777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:656
		// _ = "end of CoverTab[62777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:656
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:656
	// _ = "end of CoverTab[62767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:656
	_go_fuzz_dep_.CoverTab[62768]++
													svrskts := svr.sockets
													l := int64(len(svrskts))
													ids := make([]int64, 0, l)
													sks := make([]*normalSocket, 0, min(l, maxResults))
													for k := range svrskts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:661
		_go_fuzz_dep_.CoverTab[62778]++
														ids = append(ids, k)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:662
		// _ = "end of CoverTab[62778]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:663
	// _ = "end of CoverTab[62768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:663
	_go_fuzz_dep_.CoverTab[62769]++
													sort.Sort(int64Slice(ids))
													idx := sort.Search(len(ids), func(i int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
		_go_fuzz_dep_.CoverTab[62779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
		return ids[i] >= startID
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
		// _ = "end of CoverTab[62779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
	// _ = "end of CoverTab[62769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:665
	_go_fuzz_dep_.CoverTab[62770]++
													count := int64(0)
													var end bool
													for i, v := range ids[idx:] {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:668
		_go_fuzz_dep_.CoverTab[62780]++
														if count == maxResults {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:669
			_go_fuzz_dep_.CoverTab[62783]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:670
			// _ = "end of CoverTab[62783]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:671
			_go_fuzz_dep_.CoverTab[62784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:671
			// _ = "end of CoverTab[62784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:671
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:671
		// _ = "end of CoverTab[62780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:671
		_go_fuzz_dep_.CoverTab[62781]++
														if ns, ok := c.normalSockets[v]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:672
			_go_fuzz_dep_.CoverTab[62785]++
															sks = append(sks, ns)
															count++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:674
			// _ = "end of CoverTab[62785]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:675
			_go_fuzz_dep_.CoverTab[62786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:675
			// _ = "end of CoverTab[62786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:675
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:675
		// _ = "end of CoverTab[62781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:675
		_go_fuzz_dep_.CoverTab[62782]++
														if i == len(ids[idx:])-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:676
			_go_fuzz_dep_.CoverTab[62787]++
															end = true
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:678
			// _ = "end of CoverTab[62787]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:679
			_go_fuzz_dep_.CoverTab[62788]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:679
			// _ = "end of CoverTab[62788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:679
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:679
		// _ = "end of CoverTab[62782]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:680
	// _ = "end of CoverTab[62770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:680
	_go_fuzz_dep_.CoverTab[62771]++
													c.mu.RUnlock()
													if count == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:682
		_go_fuzz_dep_.CoverTab[62789]++
														end = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:683
		// _ = "end of CoverTab[62789]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:684
		_go_fuzz_dep_.CoverTab[62790]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:684
		// _ = "end of CoverTab[62790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:684
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:684
	// _ = "end of CoverTab[62771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:684
	_go_fuzz_dep_.CoverTab[62772]++
													s := make([]*SocketMetric, 0, len(sks))
													for _, ns := range sks {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:686
		_go_fuzz_dep_.CoverTab[62791]++
														sm := &SocketMetric{}
														sm.SocketData = ns.s.ChannelzMetric()
														sm.ID = ns.id
														sm.RefName = ns.refName
														s = append(s, sm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:691
		// _ = "end of CoverTab[62791]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:692
	// _ = "end of CoverTab[62772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:692
	_go_fuzz_dep_.CoverTab[62773]++
													return s, end
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:693
	// _ = "end of CoverTab[62773]"
}

func (c *channelMap) GetChannel(id int64) *ChannelMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:696
	_go_fuzz_dep_.CoverTab[62792]++
													cm := &ChannelMetric{}
													var cn *channel
													var ok bool
													c.mu.RLock()
													if cn, ok = c.channels[id]; !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:701
		_go_fuzz_dep_.CoverTab[62794]++

														c.mu.RUnlock()
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:704
		// _ = "end of CoverTab[62794]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:705
		_go_fuzz_dep_.CoverTab[62795]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:705
		// _ = "end of CoverTab[62795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:705
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:705
	// _ = "end of CoverTab[62792]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:705
	_go_fuzz_dep_.CoverTab[62793]++
													cm.NestedChans = copyMap(cn.nestedChans)
													cm.SubChans = copyMap(cn.subChans)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:710
	chanCopy := cn.c
													c.mu.RUnlock()
													cm.ChannelData = chanCopy.ChannelzMetric()
													cm.ID = cn.id
													cm.RefName = cn.refName
													cm.Trace = cn.trace.dumpData()
													return cm
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:716
	// _ = "end of CoverTab[62793]"
}

func (c *channelMap) GetSubChannel(id int64) *SubChannelMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:719
	_go_fuzz_dep_.CoverTab[62796]++
													cm := &SubChannelMetric{}
													var sc *subChannel
													var ok bool
													c.mu.RLock()
													if sc, ok = c.subChannels[id]; !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:724
		_go_fuzz_dep_.CoverTab[62798]++

														c.mu.RUnlock()
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:727
		// _ = "end of CoverTab[62798]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:728
		_go_fuzz_dep_.CoverTab[62799]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:728
		// _ = "end of CoverTab[62799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:728
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:728
	// _ = "end of CoverTab[62796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:728
	_go_fuzz_dep_.CoverTab[62797]++
													cm.Sockets = copyMap(sc.sockets)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:732
	chanCopy := sc.c
													c.mu.RUnlock()
													cm.ChannelData = chanCopy.ChannelzMetric()
													cm.ID = sc.id
													cm.RefName = sc.refName
													cm.Trace = sc.trace.dumpData()
													return cm
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:738
	// _ = "end of CoverTab[62797]"
}

func (c *channelMap) GetSocket(id int64) *SocketMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:741
	_go_fuzz_dep_.CoverTab[62800]++
													sm := &SocketMetric{}
													c.mu.RLock()
													if ls, ok := c.listenSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:744
		_go_fuzz_dep_.CoverTab[62803]++
														c.mu.RUnlock()
														sm.SocketData = ls.s.ChannelzMetric()
														sm.ID = ls.id
														sm.RefName = ls.refName
														return sm
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:749
		// _ = "end of CoverTab[62803]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:750
		_go_fuzz_dep_.CoverTab[62804]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:750
		// _ = "end of CoverTab[62804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:750
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:750
	// _ = "end of CoverTab[62800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:750
	_go_fuzz_dep_.CoverTab[62801]++
													if ns, ok := c.normalSockets[id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:751
		_go_fuzz_dep_.CoverTab[62805]++
														c.mu.RUnlock()
														sm.SocketData = ns.s.ChannelzMetric()
														sm.ID = ns.id
														sm.RefName = ns.refName
														return sm
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:756
		// _ = "end of CoverTab[62805]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:757
		_go_fuzz_dep_.CoverTab[62806]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:757
		// _ = "end of CoverTab[62806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:757
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:757
	// _ = "end of CoverTab[62801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:757
	_go_fuzz_dep_.CoverTab[62802]++
													c.mu.RUnlock()
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:759
	// _ = "end of CoverTab[62802]"
}

func (c *channelMap) GetServer(id int64) *ServerMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:762
	_go_fuzz_dep_.CoverTab[62807]++
													sm := &ServerMetric{}
													var svr *server
													var ok bool
													c.mu.RLock()
													if svr, ok = c.servers[id]; !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:767
		_go_fuzz_dep_.CoverTab[62809]++
														c.mu.RUnlock()
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:769
		// _ = "end of CoverTab[62809]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:770
		_go_fuzz_dep_.CoverTab[62810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:770
		// _ = "end of CoverTab[62810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:770
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:770
	// _ = "end of CoverTab[62807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:770
	_go_fuzz_dep_.CoverTab[62808]++
													sm.ListenSockets = copyMap(svr.listenSockets)
													c.mu.RUnlock()
													sm.ID = svr.id
													sm.RefName = svr.refName
													sm.ServerData = svr.s.ChannelzMetric()
													return sm
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:776
	// _ = "end of CoverTab[62808]"
}

type idGenerator struct {
	id int64
}

func (i *idGenerator) reset() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:783
	_go_fuzz_dep_.CoverTab[62811]++
													atomic.StoreInt64(&i.id, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:784
	// _ = "end of CoverTab[62811]"
}

func (i *idGenerator) genID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:787
	_go_fuzz_dep_.CoverTab[62812]++
													return atomic.AddInt64(&i.id, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:788
	// _ = "end of CoverTab[62812]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:789
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/funcs.go:789
var _ = _go_fuzz_dep_.CoverTab
