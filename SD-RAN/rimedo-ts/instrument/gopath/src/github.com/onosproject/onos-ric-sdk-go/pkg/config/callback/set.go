// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
package callback

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:5
)

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// Set :
func (c *Config) Set(req configurable.SetRequest) (configurable.SetResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:15
	_go_fuzz_dep_.CoverTab[194199]++
														log.Debugf("Set Callback is called for:%+v", req)
														var results []*pb.UpdateResult
														for _, upd := range req.ReplacePaths {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:18
		_go_fuzz_dep_.CoverTab[194202]++
															fullPath := utils.GnmiFullPath(req.Prefix, upd.Path)
															xpath := utils.ToXPath(fullPath)

															entry := store.Entry{
			Key:		xpath,
			Value:		upd.GetVal(),
			EventType:	pb.UpdateResult_REPLACE.String(),
		}
		err := c.config.Put(xpath, entry)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:28
			_go_fuzz_dep_.CoverTab[194204]++
																return configurable.SetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:29
			// _ = "end of CoverTab[194204]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:30
			_go_fuzz_dep_.CoverTab[194205]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:30
			// _ = "end of CoverTab[194205]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:30
		// _ = "end of CoverTab[194202]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:30
		_go_fuzz_dep_.CoverTab[194203]++

															update := &pb.UpdateResult{
			Op:	pb.UpdateResult_REPLACE,
			Path:	upd.Path,
		}
															results = append(results, update)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:36
		// _ = "end of CoverTab[194203]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:37
	// _ = "end of CoverTab[194199]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:37
	_go_fuzz_dep_.CoverTab[194200]++

														for _, upd := range req.UpdatePaths {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:39
		_go_fuzz_dep_.CoverTab[194206]++
															fullPath := utils.GnmiFullPath(req.Prefix, upd.Path)
															xpath := utils.ToXPath(fullPath)
															entry := store.Entry{
			Key:		xpath,
			Value:		upd.GetVal(),
			EventType:	pb.UpdateResult_UPDATE.String(),
		}
		err := c.config.Put(xpath, entry)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:48
			_go_fuzz_dep_.CoverTab[194208]++
																return configurable.SetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:49
			// _ = "end of CoverTab[194208]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:50
			_go_fuzz_dep_.CoverTab[194209]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:50
			// _ = "end of CoverTab[194209]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:50
		// _ = "end of CoverTab[194206]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:50
		_go_fuzz_dep_.CoverTab[194207]++

															update := &pb.UpdateResult{
			Op:	pb.UpdateResult_UPDATE,
			Path:	upd.Path,
		}
															results = append(results, update)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:56
		// _ = "end of CoverTab[194207]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:58
	// _ = "end of CoverTab[194200]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:58
	_go_fuzz_dep_.CoverTab[194201]++

														return configurable.SetResponse{
		Results: results,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:62
	// _ = "end of CoverTab[194201]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/set.go:63
var _ = _go_fuzz_dep_.CoverTab
