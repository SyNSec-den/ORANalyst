//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
//
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
// SPDX-License-Identifier: Apache-2.0
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
// Created by RIMEDO-Labs team
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:1
// based on onosproject/onos-mho/pkg/rnib/rnib.go
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
package rnib

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:7
)

import (
	"context"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
)

var log = logging.GetLogger("rimedo-ts", "rnib")

type TopoClient interface {
	WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error
}

type Options struct {
	TopoAddress	string
	TopoPort	int
}

type Cell struct {
	CGI		string
	CellType	string
}

func NewClient(options Options) (Client, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:33
	_go_fuzz_dep_.CoverTab[196348]++
									sdkClient, err := toposdk.NewClient(
		toposdk.WithTopoAddress(
			options.TopoAddress,
			options.TopoPort,
		),
	)
	if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:40
		_go_fuzz_dep_.CoverTab[196350]++
										return Client{}, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:41
		// _ = "end of CoverTab[196350]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:42
		_go_fuzz_dep_.CoverTab[196351]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:42
		// _ = "end of CoverTab[196351]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:42
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:42
	// _ = "end of CoverTab[196348]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:42
	_go_fuzz_dep_.CoverTab[196349]++
									return Client{
		client: sdkClient,
	}, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:45
	// _ = "end of CoverTab[196349]"
}

type Client struct {
	client toposdk.Client
}

func getControlRelationFilter() *topoapi.Filters {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:52
	_go_fuzz_dep_.CoverTab[196352]++
									controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	}
									return controlRelationFilter
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:62
	// _ = "end of CoverTab[196352]"
}

func (c *Client) WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:65
	_go_fuzz_dep_.CoverTab[196353]++
									err := c.client.Watch(ctx, ch, toposdk.WithWatchFilters(getControlRelationFilter()))
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:67
		_go_fuzz_dep_.CoverTab[196355]++
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:68
		// _ = "end of CoverTab[196355]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:69
		_go_fuzz_dep_.CoverTab[196356]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:69
		// _ = "end of CoverTab[196356]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:69
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:69
	// _ = "end of CoverTab[196353]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:69
	_go_fuzz_dep_.CoverTab[196354]++
									return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:70
	// _ = "end of CoverTab[196354]"
}

func (c *Client) GetE2CellFilter() *topoapi.Filters {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:73
	_go_fuzz_dep_.CoverTab[196357]++
									cellEntityFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_In{
				In: &topoapi.InFilter{
					Values: []string{topoapi.E2CELL},
				},
			},
		},
	}
									return cellEntityFilter
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:83
	// _ = "end of CoverTab[196357]"
}

func (c *Client) GetCellTypes(ctx context.Context) (map[string]Cell, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:86
	_go_fuzz_dep_.CoverTab[196358]++
									output := make(map[string]Cell)

									cells, err := c.client.List(ctx, toposdk.WithListFilters(c.GetE2CellFilter()))
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:90
		_go_fuzz_dep_.CoverTab[196361]++
										log.Warn(err)
										return output, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:92
		// _ = "end of CoverTab[196361]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:93
		_go_fuzz_dep_.CoverTab[196362]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:93
		// _ = "end of CoverTab[196362]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:93
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:93
	// _ = "end of CoverTab[196358]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:93
	_go_fuzz_dep_.CoverTab[196359]++

									for _, cell := range cells {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:95
		_go_fuzz_dep_.CoverTab[196363]++

										cellObject := &topoapi.E2Cell{}
										err = cell.GetAspect(cellObject)
										if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:99
			_go_fuzz_dep_.CoverTab[196365]++
											log.Warn(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:100
			// _ = "end of CoverTab[196365]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:101
			_go_fuzz_dep_.CoverTab[196366]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:101
			// _ = "end of CoverTab[196366]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:101
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:101
		// _ = "end of CoverTab[196363]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:101
		_go_fuzz_dep_.CoverTab[196364]++
										output[string(cell.ID)] = Cell{
			CGI:		cellObject.CellObjectID,
			CellType:	cellObject.CellType,
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:105
		// _ = "end of CoverTab[196364]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:106
	// _ = "end of CoverTab[196359]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:106
	_go_fuzz_dep_.CoverTab[196360]++
									return output, nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:107
	// _ = "end of CoverTab[196360]"
}

func (c *Client) SetCellType(ctx context.Context, id string, cellType string) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:110
	_go_fuzz_dep_.CoverTab[196367]++
									cell, err := c.client.Get(ctx, topoapi.ID(id))
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:112
		_go_fuzz_dep_.CoverTab[196372]++
										log.Warn(err)
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:114
		// _ = "end of CoverTab[196372]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:115
		_go_fuzz_dep_.CoverTab[196373]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:115
		// _ = "end of CoverTab[196373]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:115
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:115
	// _ = "end of CoverTab[196367]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:115
	_go_fuzz_dep_.CoverTab[196368]++

									cellObject := &topoapi.E2Cell{}
									err = cell.GetAspect(cellObject)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:119
		_go_fuzz_dep_.CoverTab[196374]++
										log.Warn(err)
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:121
		// _ = "end of CoverTab[196374]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:122
		_go_fuzz_dep_.CoverTab[196375]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:122
		// _ = "end of CoverTab[196375]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:122
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:122
	// _ = "end of CoverTab[196368]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:122
	_go_fuzz_dep_.CoverTab[196369]++

									cellObject.CellType = cellType

									err = cell.SetAspect(cellObject)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:127
		_go_fuzz_dep_.CoverTab[196376]++
										log.Warn(err)
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:129
		// _ = "end of CoverTab[196376]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:130
		_go_fuzz_dep_.CoverTab[196377]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:130
		// _ = "end of CoverTab[196377]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:130
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:130
	// _ = "end of CoverTab[196369]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:130
	_go_fuzz_dep_.CoverTab[196370]++
									err = c.client.Update(ctx, cell)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:132
		_go_fuzz_dep_.CoverTab[196378]++
										log.Warn(err)
										return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:134
		// _ = "end of CoverTab[196378]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:135
		_go_fuzz_dep_.CoverTab[196379]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:135
		// _ = "end of CoverTab[196379]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:135
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:135
	// _ = "end of CoverTab[196370]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:135
	_go_fuzz_dep_.CoverTab[196371]++

									return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:137
	// _ = "end of CoverTab[196371]"
}

func (c *Client) GetE2NodeAspects(ctx context.Context, nodeID topoapi.ID) (*topoapi.E2Node, error) {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:140
	_go_fuzz_dep_.CoverTab[196380]++
									object, err := c.client.Get(ctx, nodeID)
									if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:142
		_go_fuzz_dep_.CoverTab[196382]++
										return nil, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:143
		// _ = "end of CoverTab[196382]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:144
		_go_fuzz_dep_.CoverTab[196383]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:144
		// _ = "end of CoverTab[196383]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:144
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:144
	// _ = "end of CoverTab[196380]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:144
	_go_fuzz_dep_.CoverTab[196381]++
									e2Node := &topoapi.E2Node{}
									err = object.GetAspect(e2Node)

									return e2Node, err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:148
	// _ = "end of CoverTab[196381]"

}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:150
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/rnib/rnib.go:150
var _ = _go_fuzz_dep_.CoverTab
