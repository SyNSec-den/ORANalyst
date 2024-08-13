//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:19
)

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

// entry represents a node in the channelz database.
type entry interface {
	// addChild adds a child e, whose channelz id is id to child list
	addChild(id int64, e entry)
	// deleteChild deletes a child with channelz id to be id from child list
	deleteChild(id int64)
	// triggerDelete tries to delete self from channelz database. However, if child
	// list is not empty, then deletion from the database is on hold until the last
	// child is deleted from database.
	triggerDelete()
	// deleteSelfIfReady check whether triggerDelete() has been called before, and whether child
	// list is now empty. If both conditions are met, then delete self from database.
	deleteSelfIfReady()
	// getParentID returns parent ID of the entry. 0 value parent ID means no parent.
	getParentID() int64
}

// dummyEntry is a fake entry to handle entry not found case.
type dummyEntry struct {
	idNotFound int64
}

func (d *dummyEntry) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:53
	_go_fuzz_dep_.CoverTab[62838]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:62
	logger.Infof("attempt to add child of type %T with id %d to a parent (id=%d) that doesn't currently exist", e, id, d.idNotFound)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:62
	// _ = "end of CoverTab[62838]"
}

func (d *dummyEntry) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:65
	_go_fuzz_dep_.CoverTab[62839]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:68
	logger.Infof("attempt to delete child with id %d from a parent (id=%d) that doesn't currently exist", id, d.idNotFound)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:68
	// _ = "end of CoverTab[62839]"
}

func (d *dummyEntry) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:71
	_go_fuzz_dep_.CoverTab[62840]++
												logger.Warningf("attempt to delete an entry (id=%d) that doesn't currently exist", d.idNotFound)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:72
	// _ = "end of CoverTab[62840]"
}

func (*dummyEntry) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:75
	_go_fuzz_dep_.CoverTab[62841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:75
	// _ = "end of CoverTab[62841]"

}

func (*dummyEntry) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:79
	_go_fuzz_dep_.CoverTab[62842]++
												return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:80
	// _ = "end of CoverTab[62842]"
}

// ChannelMetric defines the info channelz provides for a specific Channel, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:83
// includes ChannelInternalMetric and channelz-specific data, such as channelz id,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:83
// child list, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:86
type ChannelMetric struct {
	// ID is the channelz id of this channel.
	ID	int64
	// RefName is the human readable reference string of this channel.
	RefName	string
	// ChannelData contains channel internal metric reported by the channel through
	// ChannelzMetric().
	ChannelData	*ChannelInternalMetric
	// NestedChans tracks the nested channel type children of this channel in the format of
	// a map from nested channel channelz id to corresponding reference string.
	NestedChans	map[int64]string
	// SubChans tracks the subchannel type children of this channel in the format of a
	// map from subchannel channelz id to corresponding reference string.
	SubChans	map[int64]string
	// Sockets tracks the socket type children of this channel in the format of a map
	// from socket channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow channel having sockets directly,
	// therefore, this is field is unused.
	Sockets	map[int64]string
	// Trace contains the most recent traced events.
	Trace	*ChannelTrace
}

// SubChannelMetric defines the info channelz provides for a specific SubChannel,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:109
// which includes ChannelInternalMetric and channelz-specific data, such as
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:109
// channelz id, child list, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:112
type SubChannelMetric struct {
	// ID is the channelz id of this subchannel.
	ID	int64
	// RefName is the human readable reference string of this subchannel.
	RefName	string
	// ChannelData contains subchannel internal metric reported by the subchannel
	// through ChannelzMetric().
	ChannelData	*ChannelInternalMetric
	// NestedChans tracks the nested channel type children of this subchannel in the format of
	// a map from nested channel channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow subchannel to have nested channels
	// as children, therefore, this field is unused.
	NestedChans	map[int64]string
	// SubChans tracks the subchannel type children of this subchannel in the format of a
	// map from subchannel channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow subchannel to have subchannels
	// as children, therefore, this field is unused.
	SubChans	map[int64]string
	// Sockets tracks the socket type children of this subchannel in the format of a map
	// from socket channelz id to corresponding reference string.
	Sockets	map[int64]string
	// Trace contains the most recent traced events.
	Trace	*ChannelTrace
}

// ChannelInternalMetric defines the struct that the implementor of Channel interface
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:137
// should return from ChannelzMetric().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:139
type ChannelInternalMetric struct {
	// current connectivity state of the channel.
	State	connectivity.State
	// The target this channel originally tried to connect to.  May be absent
	Target	string
	// The number of calls started on the channel.
	CallsStarted	int64
	// The number of calls that have completed with an OK status.
	CallsSucceeded	int64
	// The number of calls that have a completed with a non-OK status.
	CallsFailed	int64
	// The last time a call was started on the channel.
	LastCallStartedTimestamp	time.Time
}

// ChannelTrace stores traced events on a channel/subchannel and related info.
type ChannelTrace struct {
	// EventNum is the number of events that ever got traced (i.e. including those that have been deleted)
	EventNum	int64
	// CreationTime is the creation time of the trace.
	CreationTime	time.Time
	// Events stores the most recent trace events (up to $maxTraceEntry, newer event will overwrite the
	// oldest one)
	Events	[]*TraceEvent
}

// TraceEvent represent a single trace event
type TraceEvent struct {
	// Desc is a simple description of the trace event.
	Desc	string
	// Severity states the severity of this trace event.
	Severity	Severity
	// Timestamp is the event time.
	Timestamp	time.Time
	// RefID is the id of the entity that gets referenced in the event. RefID is 0 if no other entity is
	// involved in this event.
	// e.g. SubChannel (id: 4[]) Created. --> RefID = 4, RefName = "" (inside [])
	RefID	int64
	// RefName is the reference name for the entity that gets referenced in the event.
	RefName	string
	// RefType indicates the referenced entity type, i.e Channel or SubChannel.
	RefType	RefChannelType
}

// Channel is the interface that should be satisfied in order to be tracked by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:183
// channelz as Channel or SubChannel.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:185
type Channel interface {
	ChannelzMetric() *ChannelInternalMetric
}

type dummyChannel struct{}

func (d *dummyChannel) ChannelzMetric() *ChannelInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:191
	_go_fuzz_dep_.CoverTab[62843]++
													return &ChannelInternalMetric{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:192
	// _ = "end of CoverTab[62843]"
}

type channel struct {
	refName		string
	c		Channel
	closeCalled	bool
	nestedChans	map[int64]string
	subChans	map[int64]string
	id		int64
	pid		int64
	cm		*channelMap
	trace		*channelTrace
	// traceRefCount is the number of trace events that reference this channel.
	// Non-zero traceRefCount means the trace of this channel cannot be deleted.
	traceRefCount	int32
}

func (c *channel) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:210
	_go_fuzz_dep_.CoverTab[62844]++
													switch v := e.(type) {
	case *subChannel:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:212
		_go_fuzz_dep_.CoverTab[62845]++
														c.subChans[id] = v.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:213
		// _ = "end of CoverTab[62845]"
	case *channel:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:214
		_go_fuzz_dep_.CoverTab[62846]++
														c.nestedChans[id] = v.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:215
		// _ = "end of CoverTab[62846]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:216
		_go_fuzz_dep_.CoverTab[62847]++
														logger.Errorf("cannot add a child (id = %d) of type %T to a channel", id, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:217
		// _ = "end of CoverTab[62847]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:218
	// _ = "end of CoverTab[62844]"
}

func (c *channel) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:221
	_go_fuzz_dep_.CoverTab[62848]++
													delete(c.subChans, id)
													delete(c.nestedChans, id)
													c.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:224
	// _ = "end of CoverTab[62848]"
}

func (c *channel) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:227
	_go_fuzz_dep_.CoverTab[62849]++
													c.closeCalled = true
													c.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:229
	// _ = "end of CoverTab[62849]"
}

func (c *channel) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:232
	_go_fuzz_dep_.CoverTab[62850]++
													return c.pid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:233
	// _ = "end of CoverTab[62850]"
}

// deleteSelfFromTree tries to delete the channel from the channelz entry relation tree, which means
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
// deleting the channel reference from its parent's child list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
// In order for a channel to be deleted from the tree, it must meet the criteria that, removal of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
// corresponding grpc object has been invoked, and the channel does not have any children left.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:236
// The returned boolean value indicates whether the channel has been successfully deleted from tree.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:243
func (c *channel) deleteSelfFromTree() (deleted bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:243
	_go_fuzz_dep_.CoverTab[62851]++
													if !c.closeCalled || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:244
		_go_fuzz_dep_.CoverTab[62854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:244
		return len(c.subChans)+len(c.nestedChans) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:244
		// _ = "end of CoverTab[62854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:244
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:244
		_go_fuzz_dep_.CoverTab[62855]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:245
		// _ = "end of CoverTab[62855]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:246
		_go_fuzz_dep_.CoverTab[62856]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:246
		// _ = "end of CoverTab[62856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:246
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:246
	// _ = "end of CoverTab[62851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:246
	_go_fuzz_dep_.CoverTab[62852]++

													if c.pid != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:248
		_go_fuzz_dep_.CoverTab[62857]++
														c.cm.findEntry(c.pid).deleteChild(c.id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:249
		// _ = "end of CoverTab[62857]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:250
		_go_fuzz_dep_.CoverTab[62858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:250
		// _ = "end of CoverTab[62858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:250
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:250
	// _ = "end of CoverTab[62852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:250
	_go_fuzz_dep_.CoverTab[62853]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:251
	// _ = "end of CoverTab[62853]"
}

// deleteSelfFromMap checks whether it is valid to delete the channel from the map, which means
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// deleting the channel from channelz's tracking entirely. Users can no longer use id to query the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// channel, and its memory will be garbage collected.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// The trace reference count of the channel must be 0 in order to be deleted from the map. This is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// specified in the channel tracing gRFC that as long as some other trace has reference to an entity,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// the trace of the referenced entity must not be deleted. In order to release the resource allocated
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// by grpc, the reference to the grpc object is reset to a dummy object.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// deleteSelfFromMap must be called after deleteSelfFromTree returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:254
// It returns a bool to indicate whether the channel can be safely deleted from map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:266
func (c *channel) deleteSelfFromMap() (delete bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:266
	_go_fuzz_dep_.CoverTab[62859]++
													if c.getTraceRefCount() != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:267
		_go_fuzz_dep_.CoverTab[62861]++
														c.c = &dummyChannel{}
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:269
		// _ = "end of CoverTab[62861]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:270
		_go_fuzz_dep_.CoverTab[62862]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:270
		// _ = "end of CoverTab[62862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:270
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:270
	// _ = "end of CoverTab[62859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:270
	_go_fuzz_dep_.CoverTab[62860]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:271
	// _ = "end of CoverTab[62860]"
}

// deleteSelfIfReady tries to delete the channel itself from the channelz database.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:274
// The delete process includes two steps:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:274
//  1. delete the channel from the entry relation tree, i.e. delete the channel reference from its
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:274
//     parent's child list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:274
//  2. delete the channel from the map, i.e. delete the channel entirely from channelz. Lookup by id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:274
//     will return entry not found error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:280
func (c *channel) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:280
	_go_fuzz_dep_.CoverTab[62863]++
													if !c.deleteSelfFromTree() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:281
		_go_fuzz_dep_.CoverTab[62866]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:282
		// _ = "end of CoverTab[62866]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:283
		_go_fuzz_dep_.CoverTab[62867]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:283
		// _ = "end of CoverTab[62867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:283
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:283
	// _ = "end of CoverTab[62863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:283
	_go_fuzz_dep_.CoverTab[62864]++
													if !c.deleteSelfFromMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:284
		_go_fuzz_dep_.CoverTab[62868]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:285
		// _ = "end of CoverTab[62868]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:286
		_go_fuzz_dep_.CoverTab[62869]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:286
		// _ = "end of CoverTab[62869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:286
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:286
	// _ = "end of CoverTab[62864]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:286
	_go_fuzz_dep_.CoverTab[62865]++
													c.cm.deleteEntry(c.id)
													c.trace.clear()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:288
	// _ = "end of CoverTab[62865]"
}

func (c *channel) getChannelTrace() *channelTrace {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:291
	_go_fuzz_dep_.CoverTab[62870]++
													return c.trace
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:292
	// _ = "end of CoverTab[62870]"
}

func (c *channel) incrTraceRefCount() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:295
	_go_fuzz_dep_.CoverTab[62871]++
													atomic.AddInt32(&c.traceRefCount, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:296
	// _ = "end of CoverTab[62871]"
}

func (c *channel) decrTraceRefCount() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:299
	_go_fuzz_dep_.CoverTab[62872]++
													atomic.AddInt32(&c.traceRefCount, -1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:300
	// _ = "end of CoverTab[62872]"
}

func (c *channel) getTraceRefCount() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:303
	_go_fuzz_dep_.CoverTab[62873]++
													i := atomic.LoadInt32(&c.traceRefCount)
													return int(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:305
	// _ = "end of CoverTab[62873]"
}

func (c *channel) getRefName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:308
	_go_fuzz_dep_.CoverTab[62874]++
													return c.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:309
	// _ = "end of CoverTab[62874]"
}

type subChannel struct {
	refName		string
	c		Channel
	closeCalled	bool
	sockets		map[int64]string
	id		int64
	pid		int64
	cm		*channelMap
	trace		*channelTrace
	traceRefCount	int32
}

func (sc *subChannel) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:324
	_go_fuzz_dep_.CoverTab[62875]++
													if v, ok := e.(*normalSocket); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:325
		_go_fuzz_dep_.CoverTab[62876]++
														sc.sockets[id] = v.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:326
		// _ = "end of CoverTab[62876]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:327
		_go_fuzz_dep_.CoverTab[62877]++
														logger.Errorf("cannot add a child (id = %d) of type %T to a subChannel", id, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:328
		// _ = "end of CoverTab[62877]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:329
	// _ = "end of CoverTab[62875]"
}

func (sc *subChannel) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:332
	_go_fuzz_dep_.CoverTab[62878]++
													delete(sc.sockets, id)
													sc.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:334
	// _ = "end of CoverTab[62878]"
}

func (sc *subChannel) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:337
	_go_fuzz_dep_.CoverTab[62879]++
													sc.closeCalled = true
													sc.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:339
	// _ = "end of CoverTab[62879]"
}

func (sc *subChannel) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:342
	_go_fuzz_dep_.CoverTab[62880]++
													return sc.pid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:343
	// _ = "end of CoverTab[62880]"
}

// deleteSelfFromTree tries to delete the subchannel from the channelz entry relation tree, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
// means deleting the subchannel reference from its parent's child list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
// In order for a subchannel to be deleted from the tree, it must meet the criteria that, removal of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
// the corresponding grpc object has been invoked, and the subchannel does not have any children left.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:346
// The returned boolean value indicates whether the channel has been successfully deleted from tree.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:353
func (sc *subChannel) deleteSelfFromTree() (deleted bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:353
	_go_fuzz_dep_.CoverTab[62881]++
													if !sc.closeCalled || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:354
		_go_fuzz_dep_.CoverTab[62883]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:354
		return len(sc.sockets) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:354
		// _ = "end of CoverTab[62883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:354
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:354
		_go_fuzz_dep_.CoverTab[62884]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:355
		// _ = "end of CoverTab[62884]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:356
		_go_fuzz_dep_.CoverTab[62885]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:356
		// _ = "end of CoverTab[62885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:356
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:356
	// _ = "end of CoverTab[62881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:356
	_go_fuzz_dep_.CoverTab[62882]++
													sc.cm.findEntry(sc.pid).deleteChild(sc.id)
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:358
	// _ = "end of CoverTab[62882]"
}

// deleteSelfFromMap checks whether it is valid to delete the subchannel from the map, which means
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// deleting the subchannel from channelz's tracking entirely. Users can no longer use id to query
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// the subchannel, and its memory will be garbage collected.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// The trace reference count of the subchannel must be 0 in order to be deleted from the map. This is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// specified in the channel tracing gRFC that as long as some other trace has reference to an entity,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// the trace of the referenced entity must not be deleted. In order to release the resource allocated
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// by grpc, the reference to the grpc object is reset to a dummy object.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// deleteSelfFromMap must be called after deleteSelfFromTree returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:361
// It returns a bool to indicate whether the channel can be safely deleted from map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:373
func (sc *subChannel) deleteSelfFromMap() (delete bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:373
	_go_fuzz_dep_.CoverTab[62886]++
													if sc.getTraceRefCount() != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:374
		_go_fuzz_dep_.CoverTab[62888]++

														sc.c = &dummyChannel{}
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:377
		// _ = "end of CoverTab[62888]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:378
		_go_fuzz_dep_.CoverTab[62889]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:378
		// _ = "end of CoverTab[62889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:378
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:378
	// _ = "end of CoverTab[62886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:378
	_go_fuzz_dep_.CoverTab[62887]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:379
	// _ = "end of CoverTab[62887]"
}

// deleteSelfIfReady tries to delete the subchannel itself from the channelz database.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:382
// The delete process includes two steps:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:382
//  1. delete the subchannel from the entry relation tree, i.e. delete the subchannel reference from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:382
//     its parent's child list.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:382
//  2. delete the subchannel from the map, i.e. delete the subchannel entirely from channelz. Lookup
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:382
//     by id will return entry not found error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:388
func (sc *subChannel) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:388
	_go_fuzz_dep_.CoverTab[62890]++
													if !sc.deleteSelfFromTree() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:389
		_go_fuzz_dep_.CoverTab[62893]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:390
		// _ = "end of CoverTab[62893]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:391
		_go_fuzz_dep_.CoverTab[62894]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:391
		// _ = "end of CoverTab[62894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:391
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:391
	// _ = "end of CoverTab[62890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:391
	_go_fuzz_dep_.CoverTab[62891]++
													if !sc.deleteSelfFromMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:392
		_go_fuzz_dep_.CoverTab[62895]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:393
		// _ = "end of CoverTab[62895]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:394
		_go_fuzz_dep_.CoverTab[62896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:394
		// _ = "end of CoverTab[62896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:394
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:394
	// _ = "end of CoverTab[62891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:394
	_go_fuzz_dep_.CoverTab[62892]++
													sc.cm.deleteEntry(sc.id)
													sc.trace.clear()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:396
	// _ = "end of CoverTab[62892]"
}

func (sc *subChannel) getChannelTrace() *channelTrace {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:399
	_go_fuzz_dep_.CoverTab[62897]++
													return sc.trace
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:400
	// _ = "end of CoverTab[62897]"
}

func (sc *subChannel) incrTraceRefCount() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:403
	_go_fuzz_dep_.CoverTab[62898]++
													atomic.AddInt32(&sc.traceRefCount, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:404
	// _ = "end of CoverTab[62898]"
}

func (sc *subChannel) decrTraceRefCount() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:407
	_go_fuzz_dep_.CoverTab[62899]++
													atomic.AddInt32(&sc.traceRefCount, -1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:408
	// _ = "end of CoverTab[62899]"
}

func (sc *subChannel) getTraceRefCount() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:411
	_go_fuzz_dep_.CoverTab[62900]++
													i := atomic.LoadInt32(&sc.traceRefCount)
													return int(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:413
	// _ = "end of CoverTab[62900]"
}

func (sc *subChannel) getRefName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:416
	_go_fuzz_dep_.CoverTab[62901]++
													return sc.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:417
	// _ = "end of CoverTab[62901]"
}

// SocketMetric defines the info channelz provides for a specific Socket, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:420
// includes SocketInternalMetric and channelz-specific data, such as channelz id, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:422
type SocketMetric struct {
	// ID is the channelz id of this socket.
	ID	int64
	// RefName is the human readable reference string of this socket.
	RefName	string
	// SocketData contains socket internal metric reported by the socket through
	// ChannelzMetric().
	SocketData	*SocketInternalMetric
}

// SocketInternalMetric defines the struct that the implementor of Socket interface
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:432
// should return from ChannelzMetric().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:434
type SocketInternalMetric struct {
	// The number of streams that have been started.
	StreamsStarted	int64
	// The number of streams that have ended successfully:
	// On client side, receiving frame with eos bit set.
	// On server side, sending frame with eos bit set.
	StreamsSucceeded	int64
	// The number of streams that have ended unsuccessfully:
	// On client side, termination without receiving frame with eos bit set.
	// On server side, termination without sending frame with eos bit set.
	StreamsFailed	int64
	// The number of messages successfully sent on this socket.
	MessagesSent		int64
	MessagesReceived	int64
	// The number of keep alives sent.  This is typically implemented with HTTP/2
	// ping messages.
	KeepAlivesSent	int64
	// The last time a stream was created by this endpoint.  Usually unset for
	// servers.
	LastLocalStreamCreatedTimestamp	time.Time
	// The last time a stream was created by the remote endpoint.  Usually unset
	// for clients.
	LastRemoteStreamCreatedTimestamp	time.Time
	// The last time a message was sent by this endpoint.
	LastMessageSentTimestamp	time.Time
	// The last time a message was received by this endpoint.
	LastMessageReceivedTimestamp	time.Time
	// The amount of window, granted to the local endpoint by the remote endpoint.
	// This may be slightly out of date due to network latency.  This does NOT
	// include stream level or TCP level flow control info.
	LocalFlowControlWindow	int64
	// The amount of window, granted to the remote endpoint by the local endpoint.
	// This may be slightly out of date due to network latency.  This does NOT
	// include stream level or TCP level flow control info.
	RemoteFlowControlWindow	int64
	// The locally bound address.
	LocalAddr	net.Addr
	// The remote bound address.  May be absent.
	RemoteAddr	net.Addr
	// Optional, represents the name of the remote endpoint, if different than
	// the original target name.
	RemoteName	string
	SocketOptions	*SocketOptionData
	Security	credentials.ChannelzSecurityValue
}

// Socket is the interface that should be satisfied in order to be tracked by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:480
// channelz as Socket.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:482
type Socket interface {
	ChannelzMetric() *SocketInternalMetric
}

type listenSocket struct {
	refName	string
	s	Socket
	id	int64
	pid	int64
	cm	*channelMap
}

func (ls *listenSocket) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:494
	_go_fuzz_dep_.CoverTab[62902]++
													logger.Errorf("cannot add a child (id = %d) of type %T to a listen socket", id, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:495
	// _ = "end of CoverTab[62902]"
}

func (ls *listenSocket) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:498
	_go_fuzz_dep_.CoverTab[62903]++
													logger.Errorf("cannot delete a child (id = %d) from a listen socket", id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:499
	// _ = "end of CoverTab[62903]"
}

func (ls *listenSocket) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:502
	_go_fuzz_dep_.CoverTab[62904]++
													ls.cm.deleteEntry(ls.id)
													ls.cm.findEntry(ls.pid).deleteChild(ls.id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:504
	// _ = "end of CoverTab[62904]"
}

func (ls *listenSocket) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:507
	_go_fuzz_dep_.CoverTab[62905]++
													logger.Errorf("cannot call deleteSelfIfReady on a listen socket")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:508
	// _ = "end of CoverTab[62905]"
}

func (ls *listenSocket) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:511
	_go_fuzz_dep_.CoverTab[62906]++
													return ls.pid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:512
	// _ = "end of CoverTab[62906]"
}

type normalSocket struct {
	refName	string
	s	Socket
	id	int64
	pid	int64
	cm	*channelMap
}

func (ns *normalSocket) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:523
	_go_fuzz_dep_.CoverTab[62907]++
													logger.Errorf("cannot add a child (id = %d) of type %T to a normal socket", id, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:524
	// _ = "end of CoverTab[62907]"
}

func (ns *normalSocket) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:527
	_go_fuzz_dep_.CoverTab[62908]++
													logger.Errorf("cannot delete a child (id = %d) from a normal socket", id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:528
	// _ = "end of CoverTab[62908]"
}

func (ns *normalSocket) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:531
	_go_fuzz_dep_.CoverTab[62909]++
													ns.cm.deleteEntry(ns.id)
													ns.cm.findEntry(ns.pid).deleteChild(ns.id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:533
	// _ = "end of CoverTab[62909]"
}

func (ns *normalSocket) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:536
	_go_fuzz_dep_.CoverTab[62910]++
													logger.Errorf("cannot call deleteSelfIfReady on a normal socket")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:537
	// _ = "end of CoverTab[62910]"
}

func (ns *normalSocket) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:540
	_go_fuzz_dep_.CoverTab[62911]++
													return ns.pid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:541
	// _ = "end of CoverTab[62911]"
}

// ServerMetric defines the info channelz provides for a specific Server, which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:544
// includes ServerInternalMetric and channelz-specific data, such as channelz id,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:544
// child list, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:547
type ServerMetric struct {
	// ID is the channelz id of this server.
	ID	int64
	// RefName is the human readable reference string of this server.
	RefName	string
	// ServerData contains server internal metric reported by the server through
	// ChannelzMetric().
	ServerData	*ServerInternalMetric
	// ListenSockets tracks the listener socket type children of this server in the
	// format of a map from socket channelz id to corresponding reference string.
	ListenSockets	map[int64]string
}

// ServerInternalMetric defines the struct that the implementor of Server interface
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:560
// should return from ChannelzMetric().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:562
type ServerInternalMetric struct {
	// The number of incoming calls started on the server.
	CallsStarted	int64
	// The number of incoming calls that have completed with an OK status.
	CallsSucceeded	int64
	// The number of incoming calls that have a completed with a non-OK status.
	CallsFailed	int64
	// The last time a call was started on the server.
	LastCallStartedTimestamp	time.Time
}

// Server is the interface to be satisfied in order to be tracked by channelz as
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:573
// Server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:575
type Server interface {
	ChannelzMetric() *ServerInternalMetric
}

type server struct {
	refName		string
	s		Server
	closeCalled	bool
	sockets		map[int64]string
	listenSockets	map[int64]string
	id		int64
	cm		*channelMap
}

func (s *server) addChild(id int64, e entry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:589
	_go_fuzz_dep_.CoverTab[62912]++
													switch v := e.(type) {
	case *normalSocket:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:591
		_go_fuzz_dep_.CoverTab[62913]++
														s.sockets[id] = v.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:592
		// _ = "end of CoverTab[62913]"
	case *listenSocket:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:593
		_go_fuzz_dep_.CoverTab[62914]++
														s.listenSockets[id] = v.refName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:594
		// _ = "end of CoverTab[62914]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:595
		_go_fuzz_dep_.CoverTab[62915]++
														logger.Errorf("cannot add a child (id = %d) of type %T to a server", id, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:596
		// _ = "end of CoverTab[62915]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:597
	// _ = "end of CoverTab[62912]"
}

func (s *server) deleteChild(id int64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:600
	_go_fuzz_dep_.CoverTab[62916]++
													delete(s.sockets, id)
													delete(s.listenSockets, id)
													s.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:603
	// _ = "end of CoverTab[62916]"
}

func (s *server) triggerDelete() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:606
	_go_fuzz_dep_.CoverTab[62917]++
													s.closeCalled = true
													s.deleteSelfIfReady()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:608
	// _ = "end of CoverTab[62917]"
}

func (s *server) deleteSelfIfReady() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:611
	_go_fuzz_dep_.CoverTab[62918]++
													if !s.closeCalled || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:612
		_go_fuzz_dep_.CoverTab[62920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:612
		return len(s.sockets)+len(s.listenSockets) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:612
		// _ = "end of CoverTab[62920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:612
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:612
		_go_fuzz_dep_.CoverTab[62921]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:613
		// _ = "end of CoverTab[62921]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:614
		_go_fuzz_dep_.CoverTab[62922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:614
		// _ = "end of CoverTab[62922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:614
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:614
	// _ = "end of CoverTab[62918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:614
	_go_fuzz_dep_.CoverTab[62919]++
													s.cm.deleteEntry(s.id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:615
	// _ = "end of CoverTab[62919]"
}

func (s *server) getParentID() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:618
	_go_fuzz_dep_.CoverTab[62923]++
													return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:619
	// _ = "end of CoverTab[62923]"
}

type tracedChannel interface {
	getChannelTrace() *channelTrace
	incrTraceRefCount()
	decrTraceRefCount()
	getRefName() string
}

type channelTrace struct {
	cm		*channelMap
	createdTime	time.Time
	eventCount	int64
	mu		sync.Mutex
	events		[]*TraceEvent
}

func (c *channelTrace) append(e *TraceEvent) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:637
	_go_fuzz_dep_.CoverTab[62924]++
													c.mu.Lock()
													if len(c.events) == getMaxTraceEntry() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:639
		_go_fuzz_dep_.CoverTab[62926]++
														del := c.events[0]
														c.events = c.events[1:]
														if del.RefID != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:642
			_go_fuzz_dep_.CoverTab[62927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:642
			_curRoutineNum50_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:642
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum50_)

															go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
				_go_fuzz_dep_.CoverTab[62928]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
				defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
					_go_fuzz_dep_.CoverTab[62929]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum50_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
					// _ = "end of CoverTab[62929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:644
				}()

																c.cm.mu.Lock()
																c.cm.decrTraceRefCount(del.RefID)
																c.cm.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:648
				// _ = "end of CoverTab[62928]"
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:649
			// _ = "end of CoverTab[62927]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:650
			_go_fuzz_dep_.CoverTab[62930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:650
			// _ = "end of CoverTab[62930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:650
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:650
		// _ = "end of CoverTab[62926]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:651
		_go_fuzz_dep_.CoverTab[62931]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:651
		// _ = "end of CoverTab[62931]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:651
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:651
	// _ = "end of CoverTab[62924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:651
	_go_fuzz_dep_.CoverTab[62925]++
													e.Timestamp = time.Now()
													c.events = append(c.events, e)
													c.eventCount++
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:655
	// _ = "end of CoverTab[62925]"
}

func (c *channelTrace) clear() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:658
	_go_fuzz_dep_.CoverTab[62932]++
													c.mu.Lock()
													for _, e := range c.events {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:660
		_go_fuzz_dep_.CoverTab[62934]++
														if e.RefID != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:661
			_go_fuzz_dep_.CoverTab[62935]++

															c.cm.decrTraceRefCount(e.RefID)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:663
			// _ = "end of CoverTab[62935]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:664
			_go_fuzz_dep_.CoverTab[62936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:664
			// _ = "end of CoverTab[62936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:664
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:664
		// _ = "end of CoverTab[62934]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:665
	// _ = "end of CoverTab[62932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:665
	_go_fuzz_dep_.CoverTab[62933]++
													c.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:666
	// _ = "end of CoverTab[62933]"
}

// Severity is the severity level of a trace event.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:669
// The canonical enumeration of all valid values is here:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:669
// https://github.com/grpc/grpc-proto/blob/9b13d199cc0d4703c7ea26c9c330ba695866eb23/grpc/channelz/v1/channelz.proto#L126.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:672
type Severity int

const (
	// CtUnknown indicates unknown severity of a trace event.
	CtUnknown	Severity	= iota
	// CtInfo indicates info level severity of a trace event.
	CtInfo
	// CtWarning indicates warning level severity of a trace event.
	CtWarning
	// CtError indicates error level severity of a trace event.
	CtError
)

// RefChannelType is the type of the entity being referenced in a trace event.
type RefChannelType int

const (
	// RefUnknown indicates an unknown entity type, the zero value for this type.
	RefUnknown	RefChannelType	= iota
	// RefChannel indicates the referenced entity is a Channel.
	RefChannel
	// RefSubChannel indicates the referenced entity is a SubChannel.
	RefSubChannel
	// RefServer indicates the referenced entity is a Server.
	RefServer
	// RefListenSocket indicates the referenced entity is a ListenSocket.
	RefListenSocket
	// RefNormalSocket indicates the referenced entity is a NormalSocket.
	RefNormalSocket
)

var refChannelTypeToString = map[RefChannelType]string{
	RefUnknown:		"Unknown",
	RefChannel:		"Channel",
	RefSubChannel:		"SubChannel",
	RefServer:		"Server",
	RefListenSocket:	"ListenSocket",
	RefNormalSocket:	"NormalSocket",
}

func (r RefChannelType) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:712
	_go_fuzz_dep_.CoverTab[62937]++
													return refChannelTypeToString[r]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:713
	// _ = "end of CoverTab[62937]"
}

func (c *channelTrace) dumpData() *ChannelTrace {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:716
	_go_fuzz_dep_.CoverTab[62938]++
													c.mu.Lock()
													ct := &ChannelTrace{EventNum: c.eventCount, CreationTime: c.createdTime}
													ct.Events = c.events[:len(c.events)]
													c.mu.Unlock()
													return ct
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:721
	// _ = "end of CoverTab[62938]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:722
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types.go:722
var _ = _go_fuzz_dep_.CoverTab
