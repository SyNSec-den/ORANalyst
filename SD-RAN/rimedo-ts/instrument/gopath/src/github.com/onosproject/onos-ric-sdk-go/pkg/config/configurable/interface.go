// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
package configurable

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:5
)

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

// GetRequest :
type GetRequest struct {
	Paths		[]*gnmi.Path
	Prefix		*gnmi.Path
	EncodingType	gnmi.Encoding
	DataType	string
}

// GetResponse :
type GetResponse struct {
	Notifications []*gnmi.Notification
}

// SetRequest :
type SetRequest struct {
	DeletePaths	[]*gnmi.Path
	ReplacePaths	[]*gnmi.Update
	UpdatePaths	[]*gnmi.Update
	Prefix		*gnmi.Path
}

// SetResponse :
type SetResponse struct {
	Results []*gnmi.UpdateResult
}

// Configurable interface between gnmi agent and app
type Configurable interface {
	Get(GetRequest) (GetResponse, error)
	Set(SetRequest) (SetResponse, error)
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/configurable/interface.go:41
var _ = _go_fuzz_dep_.CoverTab
