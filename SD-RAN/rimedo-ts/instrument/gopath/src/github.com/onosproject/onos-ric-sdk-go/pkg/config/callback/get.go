// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
package callback

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:5
)

import (
	"encoding/json"
	"strings"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func buildUpdate(b []byte, path *pb.Path, valType string) *pb.Update {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:17
	_go_fuzz_dep_.CoverTab[194186]++
														var update *pb.Update

														if strings.Compare(valType, "Internal") == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:20
		_go_fuzz_dep_.CoverTab[194188]++
															update = &pb.Update{Path: path, Val: &pb.TypedValue{Value: &pb.TypedValue_JsonVal{JsonVal: b}}}
															return update
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:22
		// _ = "end of CoverTab[194188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:23
		_go_fuzz_dep_.CoverTab[194189]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:23
		// _ = "end of CoverTab[194189]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:23
	// _ = "end of CoverTab[194186]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:23
	_go_fuzz_dep_.CoverTab[194187]++
														update = &pb.Update{Path: path, Val: &pb.TypedValue{Value: &pb.TypedValue_JsonIetfVal{JsonIetfVal: b}}}

														return update
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:26
	// _ = "end of CoverTab[194187]"
}

// Get :
func (c *Config) Get(req configurable.GetRequest) (configurable.GetResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:30
	_go_fuzz_dep_.CoverTab[194190]++
														log.Debugf("Get Callback is called for:%+v", req)
														notifications := make([]*pb.Notification, len(req.Paths))

														for i, path := range req.Paths {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:34
		_go_fuzz_dep_.CoverTab[194192]++
															fullPath := utils.GnmiFullPath(req.Prefix, path)
															xPath := utils.ToXPath(fullPath)
															entry, err := c.config.Get(xPath)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:38
			_go_fuzz_dep_.CoverTab[194195]++
																return configurable.GetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:39
			// _ = "end of CoverTab[194195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:40
			_go_fuzz_dep_.CoverTab[194196]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:40
			// _ = "end of CoverTab[194196]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:40
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:40
		// _ = "end of CoverTab[194192]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:40
		_go_fuzz_dep_.CoverTab[194193]++

															jsonDump, err := json.Marshal(entry.Value)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:43
			_go_fuzz_dep_.CoverTab[194197]++
																return configurable.GetResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:44
			// _ = "end of CoverTab[194197]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:45
			_go_fuzz_dep_.CoverTab[194198]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:45
			// _ = "end of CoverTab[194198]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:45
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:45
		// _ = "end of CoverTab[194193]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:45
		_go_fuzz_dep_.CoverTab[194194]++

															update := buildUpdate(jsonDump, path, "IETF")
															notifications[i] = &pb.Notification{
			Prefix:	req.Prefix,
			Update:	[]*pb.Update{update},
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:51
		// _ = "end of CoverTab[194194]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:53
	// _ = "end of CoverTab[194190]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:53
	_go_fuzz_dep_.CoverTab[194191]++

														return configurable.GetResponse{
		Notifications: notifications,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:57
	// _ = "end of CoverTab[194191]"

}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:59
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/callback/get.go:59
var _ = _go_fuzz_dep_.CoverTab
