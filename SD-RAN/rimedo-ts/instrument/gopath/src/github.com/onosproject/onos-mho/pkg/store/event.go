// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:5
)

// Event store event data structure
type Event struct {
	Key	interface{}
	Value	interface{}
	Type	interface{}
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:12
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/store/event.go:12
var _ = _go_fuzz_dep_.CoverTab
