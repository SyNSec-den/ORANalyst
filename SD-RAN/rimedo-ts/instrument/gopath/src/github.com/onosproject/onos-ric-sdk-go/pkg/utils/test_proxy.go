// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
package utils

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:5
)

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-proxy/pkg/manager"
	"io/ioutil"
	"os"
)

var mgr *manager.Manager

const (
	caCrtFile	= "/tmp/onos-proxy.cacrt"
	crtFile		= "/tmp/onos-proxy.crt"
	keyFile		= "/tmp/onos-proxy.key"
)

// StartTestProxy starts onos-proxy instance for testing purposes
func StartTestProxy() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:24
	_go_fuzz_dep_.CoverTab[190616]++
														writeFile(caCrtFile, certs.OnfCaCrt)
														writeFile(crtFile, certs.DefaultOnosConfigCrt)
														writeFile(keyFile, certs.DefaultOnosConfigKey)

														cfg := manager.Config{
		CAPath:		caCrtFile,
		KeyPath:	keyFile,
		CertPath:	crtFile,
		GRPCPort:	5151,
	}

														mgr = manager.NewManager(cfg)
														mgr.Run()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:37
	// _ = "end of CoverTab[190616]"
}

func writeFile(file string, s string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:40
	_go_fuzz_dep_.CoverTab[190617]++
														err := ioutil.WriteFile(file, []byte(s), 0644)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:42
		_go_fuzz_dep_.CoverTab[190618]++
															fmt.Printf("error writing generated code to file: %s\n", err)
															os.Exit(-1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:44
		// _ = "end of CoverTab[190618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:45
		_go_fuzz_dep_.CoverTab[190619]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:45
		// _ = "end of CoverTab[190619]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:45
	// _ = "end of CoverTab[190617]"
}

// StopTestProxy stops test instance of onos-proxy
func StopTestProxy() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:49
	_go_fuzz_dep_.CoverTab[190620]++
														mgr.Close()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:50
	// _ = "end of CoverTab[190620]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/utils/test_proxy.go:51
var _ = _go_fuzz_dep_.CoverTab
