// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
package registry

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:5
)

import (
	"encoding/json"

	_default "github.com/onosproject/onos-ric-sdk-go/pkg/config/app/default"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/callback"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"

	"github.com/onosproject/onos-lib-go/pkg/northbound"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/agent"
)

var log = logging.GetLogger("registry")

const (
	// IANA reserved port for gNMI
	gnmiAgentPort = 9339
)

// RegisterRequest :
type RegisterRequest struct {
}

// RegisterResponse :
type RegisterResponse struct {
	Config interface{}
}

// startAgent stats gnmi agent server
func startAgent(c configurable.Configurable) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:40
	_go_fuzz_dep_.CoverTab[194210]++
															s := northbound.NewServer(northbound.NewServerCfg(
		"",
		"",
		"",
		int16(gnmiAgentPort),
		true,
		northbound.SecurityConfig{}))

															service := agent.NewService(c)
															s.AddService(service)

															doneCh := make(chan error)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:52
	_curRoutineNum175_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:52
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum175_)
															go func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
		_go_fuzz_dep_.CoverTab[194212]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
			_go_fuzz_dep_.CoverTab[194214]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum175_)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
			// _ = "end of CoverTab[194214]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:53
		}()
																err := s.Serve(func(started string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:54
			_go_fuzz_dep_.CoverTab[194215]++
																	log.Info("Started gNMI Agent on port ", started)
																	close(doneCh)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:56
			// _ = "end of CoverTab[194215]"
		})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:57
		// _ = "end of CoverTab[194212]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:57
		_go_fuzz_dep_.CoverTab[194213]++
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:58
			_go_fuzz_dep_.CoverTab[194216]++
																	doneCh <- err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:59
			// _ = "end of CoverTab[194216]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:60
			_go_fuzz_dep_.CoverTab[194217]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:60
			// _ = "end of CoverTab[194217]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:60
		// _ = "end of CoverTab[194213]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:61
	// _ = "end of CoverTab[194210]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:61
	_go_fuzz_dep_.CoverTab[194211]++
															return <-doneCh
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:62
	// _ = "end of CoverTab[194211]"
}

// RegisterConfigurable registers a configurable entity and starts a gNMI agent server
func RegisterConfigurable(jsonPath string, req *RegisterRequest) (RegisterResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:66
	_go_fuzz_dep_.CoverTab[194218]++
															initialConfig, err := loadConfig(jsonPath)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:68
		_go_fuzz_dep_.CoverTab[194222]++
																log.Error("Failed to read initial config", err)
																return RegisterResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:70
		// _ = "end of CoverTab[194222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:71
		_go_fuzz_dep_.CoverTab[194223]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:71
		// _ = "end of CoverTab[194223]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:71
	// _ = "end of CoverTab[194218]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:71
	_go_fuzz_dep_.CoverTab[194219]++

															config := store.NewConfigStore()
															err = json.Unmarshal(initialConfig, &config.ConfigTree)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:75
		_go_fuzz_dep_.CoverTab[194224]++
																log.Error("Failed to unmarshal initial config to json")
																return RegisterResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:77
		// _ = "end of CoverTab[194224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:78
		_go_fuzz_dep_.CoverTab[194225]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:78
		// _ = "end of CoverTab[194225]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:78
	// _ = "end of CoverTab[194219]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:78
	_go_fuzz_dep_.CoverTab[194220]++

															configurableEntity := &callback.Config{}
															configurableEntity.InitConfig(config)
															err = startAgent(configurableEntity)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:83
		_go_fuzz_dep_.CoverTab[194226]++
																return RegisterResponse{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:84
		// _ = "end of CoverTab[194226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:85
		_go_fuzz_dep_.CoverTab[194227]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:85
		// _ = "end of CoverTab[194227]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:85
	// _ = "end of CoverTab[194220]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:85
	_go_fuzz_dep_.CoverTab[194221]++

															cfg := _default.NewConfig(config)

															response := RegisterResponse{
		Config: cfg,
	}

															return response, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:93
	// _ = "end of CoverTab[194221]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/registry/registry.go:94
var _ = _go_fuzz_dep_.CoverTab
