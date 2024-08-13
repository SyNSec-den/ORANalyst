// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
package _default

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:5
)

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/app"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
)

// Config config data structure
type Config struct {
	config *store.ConfigStore
}

// Config returns config tree
func (c *Config) Config() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:22
	_go_fuzz_dep_.CoverTab[193807]++
															return c.config.ConfigTree
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:23
	// _ = "end of CoverTab[193807]"
}

// NewConfig creates a new configuration data structure
func NewConfig(config *store.ConfigStore) *Config {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:27
	_go_fuzz_dep_.CoverTab[193808]++
															return &Config{
		config: config,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:30
	// _ = "end of CoverTab[193808]"
}

// Get gets config value based on a given key
func (c *Config) Get(key string) (app.Entry, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:34
	_go_fuzz_dep_.CoverTab[193809]++
															entry, err := c.config.Get(key)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:36
		_go_fuzz_dep_.CoverTab[193811]++
																return app.Entry{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:37
		// _ = "end of CoverTab[193811]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:38
		_go_fuzz_dep_.CoverTab[193812]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:38
		// _ = "end of CoverTab[193812]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:38
	// _ = "end of CoverTab[193809]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:38
	_go_fuzz_dep_.CoverTab[193810]++

															return app.Entry{Value: entry.Value}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:40
	// _ = "end of CoverTab[193810]"
}

// Watch monitors config changes
func (c *Config) Watch(ctx context.Context, ch chan event.Event) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:44
	_go_fuzz_dep_.CoverTab[193813]++
															return c.config.Watch(ctx, ch)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:45
	// _ = "end of CoverTab[193813]"
}

var _ app.Configuration = &Config{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/app/default/impl.go:48
var _ = _go_fuzz_dep_.CoverTab
