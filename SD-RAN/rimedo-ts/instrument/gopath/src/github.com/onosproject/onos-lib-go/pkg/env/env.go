// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
package env

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:15
)

import "os"

// PodNamespace is the name of the environment variable containing the pod namespace
const PodNamespace = "POD_NAMESPACE"

// GetPodNamespace gets the pod namespace from the environment
func GetPodNamespace() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:23
	_go_fuzz_dep_.CoverTab[182656]++
												return os.Getenv(PodNamespace)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:24
	// _ = "end of CoverTab[182656]"
}

// PodName is the name of the environment variable containing the pod name
const PodName = "POD_NAME"

// GetPodName gets the pod name from the environment
func GetPodName() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:31
	_go_fuzz_dep_.CoverTab[182657]++
												return os.Getenv(PodName)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:32
	// _ = "end of CoverTab[182657]"
}

// PodID is the name of the environment variable containing the pod network identifier
const PodID = "POD_ID"

// GetPodID gets the pod network identifier from the environment
func GetPodID() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:39
	_go_fuzz_dep_.CoverTab[182658]++
												return os.Getenv(PodID)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:40
	// _ = "end of CoverTab[182658]"
}

// PodIP is the name of the environment variable containing the pod IP address
const PodIP = "POD_IP"

// GetPodIP gets the pod IP address from the environment
func GetPodIP() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:47
	_go_fuzz_dep_.CoverTab[182659]++
												return os.Getenv(PodIP)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:48
	// _ = "end of CoverTab[182659]"
}

// ServiceNamespace is the name of the environment variable containing the service namespace
const ServiceNamespace = "SERVICE_NAMESPACE"

// GetServiceNamespace gets the service namespace from the environment
func GetServiceNamespace() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:55
	_go_fuzz_dep_.CoverTab[182660]++
												return os.Getenv(ServiceNamespace)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:56
	// _ = "end of CoverTab[182660]"
}

// ServiceName is the name of the environment variable containing the service name
const ServiceName = "SERVICE_NAME"

// GetServiceName gets the service name from the environment
func GetServiceName() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:63
	_go_fuzz_dep_.CoverTab[182661]++
												return os.Getenv(ServiceName)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:64
	// _ = "end of CoverTab[182661]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:65
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/env/env.go:65
var _ = _go_fuzz_dep_.CoverTab
