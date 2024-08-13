// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
package topo

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:5
)

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"io"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"google.golang.org/grpc/credentials"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/connection"

	"google.golang.org/grpc"
)

var log = logging.GetLogger("topo")

// Client is a topo SDK client
type Client interface {
	// Create creates a topo object
	Create(ctx context.Context, object *topoapi.Object) error

	// Update updates a topo object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets a topo object with a given ID
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// Watch provides a simple facility for the application to watch for changes in the topology
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error

	// List of topo objects
	List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error)
}

// NewClient creates a new topo client
func NewClient(opts ...Option) (Client, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:48
	_go_fuzz_dep_.CoverTab[183000]++
													clientOptions := Options{}

													for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:51
		_go_fuzz_dep_.CoverTab[183005]++
														opt.apply(&clientOptions)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:52
		// _ = "end of CoverTab[183005]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:53
	// _ = "end of CoverTab[183000]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:53
	_go_fuzz_dep_.CoverTab[183001]++

													if clientOptions.Service.Host == "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:55
		_go_fuzz_dep_.CoverTab[183006]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:55
		return clientOptions.Service.Port == 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:55
		// _ = "end of CoverTab[183006]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:55
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:55
		_go_fuzz_dep_.CoverTab[183007]++
														clientOptions.Service.Host = DefaultServiceHost
														clientOptions.Service.Port = DefaultServicePort
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:57
		// _ = "end of CoverTab[183007]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:58
		_go_fuzz_dep_.CoverTab[183008]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:58
		// _ = "end of CoverTab[183008]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:58
	// _ = "end of CoverTab[183001]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:58
	_go_fuzz_dep_.CoverTab[183002]++

													dialOpts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()),
	}
	if clientOptions.Service.Insecure {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:64
		_go_fuzz_dep_.CoverTab[183009]++
														dialOpts = append(dialOpts, grpc.WithInsecure())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:65
		// _ = "end of CoverTab[183009]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:66
		_go_fuzz_dep_.CoverTab[183010]++
														tlsConfig, err := creds.GetClientCredentials()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:68
			_go_fuzz_dep_.CoverTab[183012]++
															log.Warn(err)
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:70
			// _ = "end of CoverTab[183012]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:71
			_go_fuzz_dep_.CoverTab[183013]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:71
			// _ = "end of CoverTab[183013]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:71
		// _ = "end of CoverTab[183010]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:71
		_go_fuzz_dep_.CoverTab[183011]++

														dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:73
		// _ = "end of CoverTab[183011]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:74
	// _ = "end of CoverTab[183002]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:74
	_go_fuzz_dep_.CoverTab[183003]++
													conns := connection.NewManager()
													conn, err := conns.Connect(clientOptions.Service.GetAddress(), dialOpts...)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:77
		_go_fuzz_dep_.CoverTab[183014]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:78
		// _ = "end of CoverTab[183014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:79
		_go_fuzz_dep_.CoverTab[183015]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:79
		// _ = "end of CoverTab[183015]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:79
	// _ = "end of CoverTab[183003]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:79
	_go_fuzz_dep_.CoverTab[183004]++

													cl := topoapi.NewTopoClient(conn)

													return &topo{
		client: cl,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:85
	// _ = "end of CoverTab[183004]"
}

// topo is the topo client
type topo struct {
	client topoapi.TopoClient
}

// Create creates a topo object
func (t *topo) Create(ctx context.Context, object *topoapi.Object) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:94
	_go_fuzz_dep_.CoverTab[183016]++
													response, err := t.client.Create(ctx, &topoapi.CreateRequest{
		Object: object,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:98
		_go_fuzz_dep_.CoverTab[183018]++
														return errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:99
		// _ = "end of CoverTab[183018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:100
		_go_fuzz_dep_.CoverTab[183019]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:100
		// _ = "end of CoverTab[183019]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:100
	// _ = "end of CoverTab[183016]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:100
	_go_fuzz_dep_.CoverTab[183017]++
													*object = *response.Object
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:102
	// _ = "end of CoverTab[183017]"
}

// Update updates a given topo object
func (t *topo) Update(ctx context.Context, object *topoapi.Object) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:106
	_go_fuzz_dep_.CoverTab[183020]++
													response, err := t.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:110
		_go_fuzz_dep_.CoverTab[183022]++
														return errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:111
		// _ = "end of CoverTab[183022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:112
		_go_fuzz_dep_.CoverTab[183023]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:112
		// _ = "end of CoverTab[183023]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:112
	// _ = "end of CoverTab[183020]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:112
	_go_fuzz_dep_.CoverTab[183021]++

													*object = *response.Object
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:115
	// _ = "end of CoverTab[183021]"
}

// List lists all of topo objects
func (t *topo) List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:119
	_go_fuzz_dep_.CoverTab[183024]++
													options := ListOptions{}

													for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:122
		_go_fuzz_dep_.CoverTab[183027]++
														opt.apply(&options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:123
		// _ = "end of CoverTab[183027]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:124
	// _ = "end of CoverTab[183024]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:124
	_go_fuzz_dep_.CoverTab[183025]++

													response, err := t.client.List(ctx, &topoapi.ListRequest{
		Filters: options.GetFilters(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:129
		_go_fuzz_dep_.CoverTab[183028]++
														return nil, errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:130
		// _ = "end of CoverTab[183028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:131
		_go_fuzz_dep_.CoverTab[183029]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:131
		// _ = "end of CoverTab[183029]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:131
	// _ = "end of CoverTab[183025]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:131
	_go_fuzz_dep_.CoverTab[183026]++

													return response.GetObjects(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:133
	// _ = "end of CoverTab[183026]"
}

// Get get a topo object based on a given ID
func (t *topo) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:137
	_go_fuzz_dep_.CoverTab[183030]++
													response, err := t.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:141
		_go_fuzz_dep_.CoverTab[183032]++
														return nil, errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:142
		// _ = "end of CoverTab[183032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:143
		_go_fuzz_dep_.CoverTab[183033]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:143
		// _ = "end of CoverTab[183033]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:143
	// _ = "end of CoverTab[183030]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:143
	_go_fuzz_dep_.CoverTab[183031]++
													return response.GetObject(), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:144
	// _ = "end of CoverTab[183031]"
}

// Watch watches topology events
func (t *topo) Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:148
	_go_fuzz_dep_.CoverTab[183034]++

													options := WatchOptions{}
													for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:151
		_go_fuzz_dep_.CoverTab[183038]++
														opt.apply(&options)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:152
		// _ = "end of CoverTab[183038]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:153
	// _ = "end of CoverTab[183034]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:153
	_go_fuzz_dep_.CoverTab[183035]++

													req := topoapi.WatchRequest{
		Filters:	options.GetFilters(),
		Noreplay:	options.GetNoReplay(),
	}
	stream, err := t.client.Watch(ctx, &req)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:160
		_go_fuzz_dep_.CoverTab[183039]++
														defer close(ch)
														return errors.FromGRPC(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:162
		// _ = "end of CoverTab[183039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
		_go_fuzz_dep_.CoverTab[183040]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
		// _ = "end of CoverTab[183040]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
	// _ = "end of CoverTab[183035]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
	_go_fuzz_dep_.CoverTab[183036]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
	_curRoutineNum163_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:163
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum163_)

													go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
		_go_fuzz_dep_.CoverTab[183041]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
			_go_fuzz_dep_.CoverTab[183042]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum163_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
			// _ = "end of CoverTab[183042]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:165
		}()
														defer close(ch)
														for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:167
			_go_fuzz_dep_.CoverTab[183043]++
															resp, err := stream.Recv()
															if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:169
				_go_fuzz_dep_.CoverTab[183045]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:169
				return err == context.Canceled
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:169
				// _ = "end of CoverTab[183045]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:169
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:169
				_go_fuzz_dep_.CoverTab[183046]++
																break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:170
				// _ = "end of CoverTab[183046]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:171
				_go_fuzz_dep_.CoverTab[183047]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:171
				// _ = "end of CoverTab[183047]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:171
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:171
			// _ = "end of CoverTab[183043]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:171
			_go_fuzz_dep_.CoverTab[183044]++

															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:173
				_go_fuzz_dep_.CoverTab[183048]++
																stat, ok := status.FromError(err)
																if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:175
					_go_fuzz_dep_.CoverTab[183050]++
																	err = errors.FromStatus(stat)
																	if errors.IsCanceled(err) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:177
						_go_fuzz_dep_.CoverTab[183051]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:177
						return errors.IsTimeout(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:177
						// _ = "end of CoverTab[183051]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:177
					}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:177
						_go_fuzz_dep_.CoverTab[183052]++
																		break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:178
						// _ = "end of CoverTab[183052]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:179
						_go_fuzz_dep_.CoverTab[183053]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:179
						// _ = "end of CoverTab[183053]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:179
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:179
					// _ = "end of CoverTab[183050]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:180
					_go_fuzz_dep_.CoverTab[183054]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:180
					// _ = "end of CoverTab[183054]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:180
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:180
				// _ = "end of CoverTab[183048]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:180
				_go_fuzz_dep_.CoverTab[183049]++
																log.Error("An error occurred in receiving topology changes", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:181
				// _ = "end of CoverTab[183049]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:182
				_go_fuzz_dep_.CoverTab[183055]++
																ch <- resp.Event
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:183
				// _ = "end of CoverTab[183055]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:184
			// _ = "end of CoverTab[183044]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:185
		// _ = "end of CoverTab[183041]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:186
	// _ = "end of CoverTab[183036]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:186
	_go_fuzz_dep_.CoverTab[183037]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:187
	// _ = "end of CoverTab[183037]"

}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:189
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/topo/topo.go:189
var _ = _go_fuzz_dep_.CoverTab
