// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
package v1beta1

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:5
)

// ServiceModelName is a service model name
type ServiceModelName string

// ServiceModelVersion is a service model version name
type ServiceModelVersion string

// AppID is an xApp identifier
type AppID string

// AppInstanceID is an xApp instance identifier
type AppInstanceID string

// E2NodeID is an E2 node identifier
type E2NodeID string

// TransactionID is a transaction identifier
type TransactionID string

// ChannelID is a subscription channel identifier
type ChannelID string

// SubscriptionID is a subscription identifier
type SubscriptionID string

// Revision is a subscription revision number
type Revision uint64

// E2TInstanceID is an E2T instance identifier
type E2TInstanceID string

// TermID is a mastership term identifier
type TermID uint64

// MasterID is a master identifier
type MasterID string

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-api/go@v0.10.31/onos/e2t/e2/v1beta1/e2.go:41
var _ = _go_fuzz_dep_.CoverTab
