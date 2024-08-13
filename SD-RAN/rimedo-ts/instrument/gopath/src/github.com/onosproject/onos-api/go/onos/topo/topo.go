// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
package topo

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:5
)

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

// UUID represents a system-assigned unique identifier of a topology object.
type UUID string

// ID represents a client-assigned unique identifier.
type ID string

// String convert ID to string
func (id ID) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:23
	_go_fuzz_dep_.CoverTab[159396]++
													return string(id)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:24
	// _ = "end of CoverTab[159396]"
}

// NullID represents a null/empty/omitted identifier; usually an indicator for system to generate one.
const NullID = ""

// Revision is an object revision
type Revision uint64

// DEPRECATED Entity and Relation Kinds
const (
	// Relations
	CONTROLS	= "controls"
	CONTAINS	= "contains"
	NEIGHBORS	= "neighbors"

	// RAN Entities
	E2NODE	= "e2node"
	E2CELL	= "e2cell"
	E2T	= "e2t"
	XAPP	= "xapp"
	A1T	= "a1t"

	// onos-config entity
	ONOS_CONFIG	= "onos-config"
)

// TODO UPPERCASE entity kinds and relations should be replaced gradually with CamelCase ones
const (
	ControlsKind	= "controls"
	ContainsKind	= "contains"
	HasKind		= "has"
	TerminatesKind	= "terminates"
	OriginatesKind	= "originates"
	NeighborsKind	= "neighbors"
	ConnectionKind	= "connection"

	// Fabric Entity kinds
	PodKind			= "pod"
	RackKind		= "rack"
	NetworkLayerKind	= "network-layer"
	SwitchKind		= "switch"
	ServerKind		= "server"
	IPUKind			= "ipu"
	HostKind		= "host"
	RouterKind		= "router"
	PortKind		= "port"
	InterfaceKind		= "interface"
	LinkKind		= "link"
	ControllerKind		= "controller"
	ServiceKind		= "service"

	// onos-config entity
	OnosConfigKind	= "onos-config"

	// RAN Entitiy kinds
	E2NodeKind	= "e2node"
	E2CellKind	= "e2cell"
	E2tKind		= "e2t"
	XappKind	= "xapp"
	A1tKind		= "a1t"
)

// PolicyTypeID is an identifier of A1 policy type
type PolicyTypeID string

// PolicyTypeName is a name of A1 policy type
type PolicyTypeName string

// PolicyTypeVersion is a version of A1 policy type
type PolicyTypeVersion string

// PolicyTypeDescription describe what this A1 policy is
type PolicyTypeDescription string

// TopoClientFactory : Default EntityServiceClient creation.
var TopoClientFactory = func(cc *grpc.ClientConn) TopoClient {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:100
	_go_fuzz_dep_.CoverTab[159397]++
													return NewTopoClient(cc)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:101
	// _ = "end of CoverTab[159397]"
}

// CreateTopoClient creates and returns a new topo device client
func CreateTopoClient(cc *grpc.ClientConn) TopoClient {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:105
	_go_fuzz_dep_.CoverTab[159398]++
													return TopoClientFactory(cc)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:106
	// _ = "end of CoverTab[159398]"
}

// RelationID creates a unique relationship ID from the specified source, kind and target IDs
func RelationID(srcID ID, relationKind ID, tgtID ID) ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:110
	_go_fuzz_dep_.CoverTab[159399]++
													return ID(fmt.Sprintf("%s-%s-%s", srcID, relationKind, tgtID))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:111
	// _ = "end of CoverTab[159399]"
}

// MultiRelationID creates a unique relationship ID from the specified source, kind and target IDs,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:114
// and also from an additional discriminant to allow for multiples of same kinds of relations between
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:114
// the same two objects.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:117
func MultiRelationID(srcID ID, relationKind ID, tgtID ID, discriminant uint8) ID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:117
	_go_fuzz_dep_.CoverTab[159400]++
													return ID(fmt.Sprintf("%s-%s-%s-%d", srcID, relationKind, tgtID, discriminant))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:118
	// _ = "end of CoverTab[159400]"
}

// NewEntity allocates a new topology entity using the specified ID and kind.
func NewEntity(id ID, kind ID, aspects ...proto.Message) *Object {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:122
	_go_fuzz_dep_.CoverTab[159401]++
													return &Object{ID: id, Type: Object_ENTITY, Obj: &Object_Entity{Entity: &Entity{KindID: kind}}}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:123
	// _ = "end of CoverTab[159401]"
}

// NewRelation allocates a new topology relation using the specified source, target, and kind.
func NewRelation(source ID, target ID, kind ID, aspects ...proto.Message) *Object {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:127
	_go_fuzz_dep_.CoverTab[159402]++
													return &Object{
		ID:	RelationID(source, kind, target),
		Type:	Object_RELATION,
		Obj:	&Object_Relation{Relation: &Relation{SrcEntityID: source, TgtEntityID: target, KindID: kind}},
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:132
	// _ = "end of CoverTab[159402]"
}

// WithAspects applies the given aspects to the object.
func (obj *Object) WithAspects(aspects ...proto.Message) (*Object, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:136
	_go_fuzz_dep_.CoverTab[159403]++
													for _, aspect := range aspects {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:137
		_go_fuzz_dep_.CoverTab[159405]++
														if err := obj.SetAspect(aspect); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:138
			_go_fuzz_dep_.CoverTab[159406]++
															return obj, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:139
			// _ = "end of CoverTab[159406]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:140
			_go_fuzz_dep_.CoverTab[159407]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:140
			// _ = "end of CoverTab[159407]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:140
		// _ = "end of CoverTab[159405]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:141
	// _ = "end of CoverTab[159403]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:141
	_go_fuzz_dep_.CoverTab[159404]++
													return obj, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:142
	// _ = "end of CoverTab[159404]"
}

// ToAny provides a convenience utility to convert an aspect message to protobuf types.Any
func ToAny(value proto.Message) *types.Any {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:146
	_go_fuzz_dep_.CoverTab[159408]++
													jm := jsonpb.Marshaler{}
													writer := bytes.Buffer{}
													if err := jm.Marshal(&writer, value); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:149
		_go_fuzz_dep_.CoverTab[159410]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:150
		// _ = "end of CoverTab[159410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:151
		_go_fuzz_dep_.CoverTab[159411]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:151
		// _ = "end of CoverTab[159411]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:151
	// _ = "end of CoverTab[159408]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:151
	_go_fuzz_dep_.CoverTab[159409]++
													return &types.Any{
		TypeUrl:	proto.MessageName(value),
		Value:		writer.Bytes(),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:155
	// _ = "end of CoverTab[159409]"
}

// GetAspect retrieves the specified aspect value from the given object.
func (obj *Object) GetAspect(destValue proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:159
	_go_fuzz_dep_.CoverTab[159412]++
													if obj.Aspects == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:160
		_go_fuzz_dep_.CoverTab[159417]++
														return fmt.Errorf("no aspects found on %s", obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:161
		// _ = "end of CoverTab[159417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:162
		_go_fuzz_dep_.CoverTab[159418]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:162
		// _ = "end of CoverTab[159418]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:162
	// _ = "end of CoverTab[159412]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:162
	_go_fuzz_dep_.CoverTab[159413]++
													aspectType := proto.MessageName(destValue)
													any := obj.Aspects[aspectType]
													if any == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:165
		_go_fuzz_dep_.CoverTab[159419]++
														return fmt.Errorf("aspect '%s' not found in %s", aspectType, obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:166
		// _ = "end of CoverTab[159419]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:167
		_go_fuzz_dep_.CoverTab[159420]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:167
		// _ = "end of CoverTab[159420]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:167
	// _ = "end of CoverTab[159413]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:167
	_go_fuzz_dep_.CoverTab[159414]++
													if any.TypeUrl != aspectType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:168
		_go_fuzz_dep_.CoverTab[159421]++
														return fmt.Errorf("unexpected aspect type: %s", aspectType)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:169
		// _ = "end of CoverTab[159421]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:170
		_go_fuzz_dep_.CoverTab[159422]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:170
		// _ = "end of CoverTab[159422]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:170
	// _ = "end of CoverTab[159414]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:170
	_go_fuzz_dep_.CoverTab[159415]++
													reader := bytes.NewReader(any.Value)
													err := jsonpb.Unmarshal(reader, destValue)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:173
		_go_fuzz_dep_.CoverTab[159423]++
														return fmt.Errorf("error '%s' when unmarshalling aspect %s: %v from %s",
			err.Error(), aspectType, any.Value, obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:175
		// _ = "end of CoverTab[159423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:176
		_go_fuzz_dep_.CoverTab[159424]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:176
		// _ = "end of CoverTab[159424]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:176
	// _ = "end of CoverTab[159415]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:176
	_go_fuzz_dep_.CoverTab[159416]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:177
	// _ = "end of CoverTab[159416]"
}

// GetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) GetAspectBytes(aspectType string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:181
	_go_fuzz_dep_.CoverTab[159425]++
													if obj.Aspects == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:182
		_go_fuzz_dep_.CoverTab[159428]++
														return nil, fmt.Errorf("no aspects found on %s", obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:183
		// _ = "end of CoverTab[159428]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:184
		_go_fuzz_dep_.CoverTab[159429]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:184
		// _ = "end of CoverTab[159429]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:184
	// _ = "end of CoverTab[159425]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:184
	_go_fuzz_dep_.CoverTab[159426]++
													any := obj.Aspects[aspectType]
													if any == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:186
		_go_fuzz_dep_.CoverTab[159430]++
														return nil, fmt.Errorf("aspect '%s' not found on %s", aspectType, obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:187
		// _ = "end of CoverTab[159430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:188
		_go_fuzz_dep_.CoverTab[159431]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:188
		// _ = "end of CoverTab[159431]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:188
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:188
	// _ = "end of CoverTab[159426]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:188
	_go_fuzz_dep_.CoverTab[159427]++
													return any.Value, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:189
	// _ = "end of CoverTab[159427]"
}

// SetAspect applies the specified aspect value to the given object.
func (obj *Object) SetAspect(value proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:193
	_go_fuzz_dep_.CoverTab[159432]++
													jm := jsonpb.Marshaler{}
													writer := bytes.Buffer{}
													err := jm.Marshal(&writer, value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:197
		_go_fuzz_dep_.CoverTab[159435]++
														return fmt.Errorf("error '%s' marshaling aspect %v on to %s",
			err.Error(), value, obj.String())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:199
		// _ = "end of CoverTab[159435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:200
		_go_fuzz_dep_.CoverTab[159436]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:200
		// _ = "end of CoverTab[159436]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:200
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:200
	// _ = "end of CoverTab[159432]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:200
	_go_fuzz_dep_.CoverTab[159433]++
													if obj.Aspects == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:201
		_go_fuzz_dep_.CoverTab[159437]++
														obj.Aspects = make(map[string]*types.Any)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:202
		// _ = "end of CoverTab[159437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:203
		_go_fuzz_dep_.CoverTab[159438]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:203
		// _ = "end of CoverTab[159438]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:203
	// _ = "end of CoverTab[159433]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:203
	_go_fuzz_dep_.CoverTab[159434]++
													obj.Aspects[proto.MessageName(value)] = &types.Any{
		TypeUrl:	proto.MessageName(value),
		Value:		writer.Bytes(),
	}
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:208
	// _ = "end of CoverTab[159434]"
}

// SetAspectBytes applies the specified aspect as raw JSON bytes to the given object.
func (obj *Object) SetAspectBytes(aspectType string, jsonValue []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:212
	_go_fuzz_dep_.CoverTab[159439]++
													any := &types.Any{
		TypeUrl:	aspectType,
		Value:		jsonValue,
	}
	if obj.Aspects == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:217
		_go_fuzz_dep_.CoverTab[159441]++
														obj.Aspects = make(map[string]*types.Any)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:218
		// _ = "end of CoverTab[159441]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:219
		_go_fuzz_dep_.CoverTab[159442]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:219
		// _ = "end of CoverTab[159442]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:219
	// _ = "end of CoverTab[159439]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:219
	_go_fuzz_dep_.CoverTab[159440]++
													obj.Aspects[aspectType] = any
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:221
	// _ = "end of CoverTab[159440]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:222
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/topo/topo.go:222
var _ = _go_fuzz_dep_.CoverTab
