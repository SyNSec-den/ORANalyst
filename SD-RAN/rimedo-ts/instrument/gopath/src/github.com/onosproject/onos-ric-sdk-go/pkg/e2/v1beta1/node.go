// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
package e2

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:5
)

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1/e2errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"io"
	"sync"
	"time"
)

// NodeID is an E2 node identifier
type NodeID string

// Node is an E2 node
type Node interface {
	// ID is the node identifier
	ID() NodeID

	// Context is the node context
	Context() context.Context

	// Subscribe creates a subscription from the given SubscriptionDetails
	// The Subscribe method will block until the subscription is successfully registered.
	// The context.Context represents the lifecycle of this initial subscription process.
	// Once the subscription has been created and the method returns, indications will be written
	// to the given channel.
	// If the subscription is successful, a subscription.Context will be returned. The subscription
	// context can be used to cancel the subscription by calling Close() on the subscription.Context.
	Subscribe(ctx context.Context, name string, sub e2api.SubscriptionSpec, indCh chan<- e2api.Indication, opts ...SubscribeOption) (e2api.ChannelID, error)

	// Unsubscribe unsubscribes from the given subscription
	Unsubscribe(ctx context.Context, name string) error

	// Control creates and sends a E2 control message and awaits the outcome
	Control(ctx context.Context, message *e2api.ControlMessage) (*e2api.ControlOutcome, error)
}

// NewNode creates a new E2 Node with the given ID
func NewNode(nodeID NodeID, opts ...Option) Node {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:53
	_go_fuzz_dep_.CoverTab[196525]++
														options := Options{
		App: AppOptions{
			AppID:		AppID(env.GetServiceName()),
			InstanceID:	InstanceID(env.GetPodName()),
		},
		Service: ServiceOptions{
			Host:	"onos-e2t",
			Port:	defaultServicePort,
		},
		Topo: ServiceOptions{
			Host:	"onos-topo",
			Port:	defaultServicePort,
		},
		Encoding:	ProtoEncoding,
	}
	for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:69
		_go_fuzz_dep_.CoverTab[196527]++
															opt.apply(&options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:70
		// _ = "end of CoverTab[196527]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:71
	// _ = "end of CoverTab[196525]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:71
	_go_fuzz_dep_.CoverTab[196526]++

														uuid.SetNodeID([]byte(fmt.Sprintf("%s:%s", options.App.AppID, options.App.InstanceID)))

														ctx, cancel := context.WithCancel(context.Background())
														return &e2Node{
		nodeID:		nodeID,
		options:	options,
		ctx:		ctx,
		cancel:		cancel,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:81
	// _ = "end of CoverTab[196526]"
}

type ackResult struct {
	err		error
	channelID	e2api.ChannelID
}

// e2Node is the default E2 node implementation
type e2Node struct {
	nodeID	NodeID
	options	Options
	ctx	context.Context
	cancel	context.CancelFunc
	conn	*grpc.ClientConn
	mu	sync.RWMutex
}

func (n *e2Node) ID() NodeID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:99
	_go_fuzz_dep_.CoverTab[196528]++
														return n.nodeID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:100
	// _ = "end of CoverTab[196528]"
}

func (n *e2Node) Context() context.Context {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:103
	_go_fuzz_dep_.CoverTab[196529]++
														return n.ctx
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:104
	// _ = "end of CoverTab[196529]"
}

func (n *e2Node) connect(ctx context.Context) (*grpc.ClientConn, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:107
	_go_fuzz_dep_.CoverTab[196530]++
														n.mu.RLock()
														conn := n.conn
														n.mu.RUnlock()

														if conn != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:112
		_go_fuzz_dep_.CoverTab[196534]++
															return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:113
		// _ = "end of CoverTab[196534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:114
		_go_fuzz_dep_.CoverTab[196535]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:114
		// _ = "end of CoverTab[196535]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:114
	// _ = "end of CoverTab[196530]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:114
	_go_fuzz_dep_.CoverTab[196531]++

														n.mu.Lock()
														defer n.mu.Unlock()

														if n.conn != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:119
		_go_fuzz_dep_.CoverTab[196536]++
															return n.conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:120
		// _ = "end of CoverTab[196536]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:121
		_go_fuzz_dep_.CoverTab[196537]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:121
		// _ = "end of CoverTab[196537]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:121
	// _ = "end of CoverTab[196531]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:121
	_go_fuzz_dep_.CoverTab[196532]++

														clientCreds, _ := creds.GetClientCredentials()
														conn, err := grpc.DialContext(ctx, "localhost:5151",
		grpc.WithTransportCredentials(credentials.NewTLS(clientCreds)),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable))),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:128
		_go_fuzz_dep_.CoverTab[196538]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:129
		// _ = "end of CoverTab[196538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:130
		_go_fuzz_dep_.CoverTab[196539]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:130
		// _ = "end of CoverTab[196539]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:130
	// _ = "end of CoverTab[196532]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:130
	_go_fuzz_dep_.CoverTab[196533]++
														n.conn = conn
														return conn, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:132
	// _ = "end of CoverTab[196533]"
}

func getErrorFromGRPC(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:135
	_go_fuzz_dep_.CoverTab[196540]++
														if e2errors.IsE2APError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:136
		_go_fuzz_dep_.CoverTab[196542]++
															return e2errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:137
		// _ = "end of CoverTab[196542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:138
		_go_fuzz_dep_.CoverTab[196543]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:138
		// _ = "end of CoverTab[196543]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:138
	// _ = "end of CoverTab[196540]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:138
	_go_fuzz_dep_.CoverTab[196541]++
														return errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:139
	// _ = "end of CoverTab[196541]"
}

func (n *e2Node) getRequestHeaders() e2api.RequestHeaders {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:142
	_go_fuzz_dep_.CoverTab[196544]++
														var encoding e2api.Encoding
														switch n.options.Encoding {
	case ProtoEncoding:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:145
		_go_fuzz_dep_.CoverTab[196546]++
															encoding = e2api.Encoding_PROTO
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:146
		// _ = "end of CoverTab[196546]"
	case ASN1Encoding:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:147
		_go_fuzz_dep_.CoverTab[196547]++
															encoding = e2api.Encoding_ASN1_PER
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:148
		// _ = "end of CoverTab[196547]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:148
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:148
		_go_fuzz_dep_.CoverTab[196548]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:148
		// _ = "end of CoverTab[196548]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:149
	// _ = "end of CoverTab[196544]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:149
	_go_fuzz_dep_.CoverTab[196545]++
														return e2api.RequestHeaders{
		AppID:		e2api.AppID(n.options.App.AppID),
		AppInstanceID:	e2api.AppInstanceID(n.options.App.InstanceID),
		E2NodeID:	e2api.E2NodeID(n.nodeID),
		ServiceModel: e2api.ServiceModel{
			Name:		e2api.ServiceModelName(n.options.ServiceModel.Name),
			Version:	e2api.ServiceModelVersion(n.options.ServiceModel.Version),
		},
		Encoding:	encoding,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:159
	// _ = "end of CoverTab[196545]"
}

func (n *e2Node) Control(ctx context.Context, message *e2api.ControlMessage) (*e2api.ControlOutcome, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:162
	_go_fuzz_dep_.CoverTab[196549]++
														conn, err := n.connect(ctx)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:164
		_go_fuzz_dep_.CoverTab[196552]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:165
		// _ = "end of CoverTab[196552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:166
		_go_fuzz_dep_.CoverTab[196553]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:166
		// _ = "end of CoverTab[196553]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:166
	// _ = "end of CoverTab[196549]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:166
	_go_fuzz_dep_.CoverTab[196550]++
														client := e2api.NewControlServiceClient(conn)

														request := &e2api.ControlRequest{
		Headers:	n.getRequestHeaders(),
		Message:	*message,
	}
	log.Debugf("Sending ControlRequest %+v", request)
	response, err := client.Control(ctx, request)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:175
		_go_fuzz_dep_.CoverTab[196554]++
															log.Warnf("ControlRequest %+v failed: %v", request, err)
															return nil, getErrorFromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:177
		// _ = "end of CoverTab[196554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:178
		_go_fuzz_dep_.CoverTab[196555]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:178
		// _ = "end of CoverTab[196555]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:178
	// _ = "end of CoverTab[196550]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:178
	_go_fuzz_dep_.CoverTab[196551]++
														log.Debugf("Received ControlResponse %+v", response)
														return &response.Outcome, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:180
	// _ = "end of CoverTab[196551]"
}

func (n *e2Node) Subscribe(ctx context.Context, name string, sub e2api.SubscriptionSpec, indCh chan<- e2api.Indication, opts ...SubscribeOption) (e2api.ChannelID, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:183
	_go_fuzz_dep_.CoverTab[196556]++
														conn, err := n.connect(ctx)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:185
		_go_fuzz_dep_.CoverTab[196561]++
															return "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:186
		// _ = "end of CoverTab[196561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:187
		_go_fuzz_dep_.CoverTab[196562]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:187
		// _ = "end of CoverTab[196562]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:187
	// _ = "end of CoverTab[196556]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:187
	_go_fuzz_dep_.CoverTab[196557]++
														client := e2api.NewSubscriptionServiceClient(conn)

														options := SubscribeOptions{TransactionTimeout: 2 * time.Minute}
														for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:191
		_go_fuzz_dep_.CoverTab[196563]++
															opt.apply(&options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:192
		// _ = "end of CoverTab[196563]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:193
	// _ = "end of CoverTab[196557]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:193
	_go_fuzz_dep_.CoverTab[196558]++

														request := &e2api.SubscribeRequest{
		Headers:		n.getRequestHeaders(),
		TransactionID:		e2api.TransactionID(name),
		Subscription:		sub,
		TransactionTimeout:	&options.TransactionTimeout,
	}
	log.Debugf("Sending SubscribeRequest %+v", request)
	stream, err := client.Subscribe(ctx, request)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:203
		_go_fuzz_dep_.CoverTab[196564]++
															defer close(indCh)
															return "", getErrorFromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:205
		// _ = "end of CoverTab[196564]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:206
		_go_fuzz_dep_.CoverTab[196565]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:206
		// _ = "end of CoverTab[196565]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:206
	// _ = "end of CoverTab[196558]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:206
	_go_fuzz_dep_.CoverTab[196559]++

														ackCh := make(chan ackResult)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:208
	_curRoutineNum186_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:208
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum186_)
														go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
		_go_fuzz_dep_.CoverTab[196566]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
			_go_fuzz_dep_.CoverTab[196568]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum186_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
			// _ = "end of CoverTab[196568]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:209
		}()
															defer close(indCh)
															acked := false
															var channelID e2api.ChannelID
															for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:213
			_go_fuzz_dep_.CoverTab[196569]++
																response, err := stream.Recv()
																if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:215
				_go_fuzz_dep_.CoverTab[196571]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:215
				return err == context.Canceled
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:215
				// _ = "end of CoverTab[196571]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:215
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:215
				_go_fuzz_dep_.CoverTab[196572]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:216
				// _ = "end of CoverTab[196572]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:217
				_go_fuzz_dep_.CoverTab[196573]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:217
				// _ = "end of CoverTab[196573]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:217
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:217
			// _ = "end of CoverTab[196569]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:217
			_go_fuzz_dep_.CoverTab[196570]++

																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:219
				_go_fuzz_dep_.CoverTab[196574]++
																	err = getErrorFromGRPC(err)
																	if errors.IsCanceled(err) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:221
					_go_fuzz_dep_.CoverTab[196576]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:221
					return errors.IsTimeout(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:221
					// _ = "end of CoverTab[196576]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:221
				}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:221
					_go_fuzz_dep_.CoverTab[196577]++
																		break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:222
					// _ = "end of CoverTab[196577]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:223
					_go_fuzz_dep_.CoverTab[196578]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:223
					// _ = "end of CoverTab[196578]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:223
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:223
				// _ = "end of CoverTab[196574]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:223
				_go_fuzz_dep_.CoverTab[196575]++

																	log.Warnf("SubscribeRequest %+v failed: %v", request, err)
																	if !acked {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:226
					_go_fuzz_dep_.CoverTab[196579]++
																		ackCh <- ackResult{
						err: err,
					}
																		close(ackCh)
																		acked = true
																		break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:232
					// _ = "end of CoverTab[196579]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:233
					_go_fuzz_dep_.CoverTab[196580]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:233
					// _ = "end of CoverTab[196580]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:233
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:233
				// _ = "end of CoverTab[196575]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:234
				_go_fuzz_dep_.CoverTab[196581]++
																	log.Debugf("Received SubscribeResponse %+v", response)
																	switch m := response.Message.(type) {
				case *e2api.SubscribeResponse_Ack:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:237
					_go_fuzz_dep_.CoverTab[196582]++
																		channelID = m.Ack.ChannelID
																		if !acked {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:239
						_go_fuzz_dep_.CoverTab[196584]++
																			ackCh <- ackResult{
							channelID: channelID,
						}
																			close(ackCh)
																			acked = true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:244
						// _ = "end of CoverTab[196584]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:245
						_go_fuzz_dep_.CoverTab[196585]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:245
						// _ = "end of CoverTab[196585]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:245
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:245
					// _ = "end of CoverTab[196582]"
				case *e2api.SubscribeResponse_Indication:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:246
					_go_fuzz_dep_.CoverTab[196583]++
																		indCh <- *m.Indication
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:247
					// _ = "end of CoverTab[196583]"
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:248
				// _ = "end of CoverTab[196581]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:249
			// _ = "end of CoverTab[196570]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:250
		// _ = "end of CoverTab[196566]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:250
		_go_fuzz_dep_.CoverTab[196567]++
															if !acked {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:251
			_go_fuzz_dep_.CoverTab[196586]++
																close(ackCh)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:252
			// _ = "end of CoverTab[196586]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:253
			_go_fuzz_dep_.CoverTab[196587]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:253
			// _ = "end of CoverTab[196587]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:253
		// _ = "end of CoverTab[196567]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:254
	// _ = "end of CoverTab[196559]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:254
	_go_fuzz_dep_.CoverTab[196560]++

														select {
	case result := <-ackCh:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:257
		_go_fuzz_dep_.CoverTab[196588]++
															return result.channelID, result.err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:258
		// _ = "end of CoverTab[196588]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:259
		_go_fuzz_dep_.CoverTab[196589]++
															return "", ctx.Err()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:260
		// _ = "end of CoverTab[196589]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:261
	// _ = "end of CoverTab[196560]"
}

func (n *e2Node) Unsubscribe(ctx context.Context, name string) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:264
	_go_fuzz_dep_.CoverTab[196590]++
														conn, err := n.connect(ctx)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:266
		_go_fuzz_dep_.CoverTab[196593]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:267
		// _ = "end of CoverTab[196593]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:268
		_go_fuzz_dep_.CoverTab[196594]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:268
		// _ = "end of CoverTab[196594]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:268
	// _ = "end of CoverTab[196590]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:268
	_go_fuzz_dep_.CoverTab[196591]++
														client := e2api.NewSubscriptionServiceClient(conn)

														request := &e2api.UnsubscribeRequest{
		Headers:	n.getRequestHeaders(),
		TransactionID:	e2api.TransactionID(name),
	}
	log.Debugf("Sending UnsubscribeRequest %+v", request)
	response, err := client.Unsubscribe(ctx, request)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:277
		_go_fuzz_dep_.CoverTab[196595]++
															log.Warnf("UnsubscribeRequest %+v failed: %v", request, err)
															return getErrorFromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:279
		// _ = "end of CoverTab[196595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:280
		_go_fuzz_dep_.CoverTab[196596]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:280
		// _ = "end of CoverTab[196596]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:280
	// _ = "end of CoverTab[196591]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:280
	_go_fuzz_dep_.CoverTab[196592]++
														log.Debugf("Received UnsubscribeResponse %+v", response)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:282
	// _ = "end of CoverTab[196592]"
}

func (n *e2Node) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:285
	_go_fuzz_dep_.CoverTab[196597]++
														defer n.cancel()
														return n.conn.Close()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:287
	// _ = "end of CoverTab[196597]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:288
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/node.go:288
var _ = _go_fuzz_dep_.CoverTab
