//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:19
// Package serviceconfig contains utility functions to parse service config.
package serviceconfig

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:20
)

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	externalserviceconfig "google.golang.org/grpc/serviceconfig"
)

var logger = grpclog.Component("core")

// BalancerConfig wraps the name and config associated with one load balancing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
// policy. It corresponds to a single entry of the loadBalancingConfig field
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
// from ServiceConfig.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
// It implements the json.Unmarshaler interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:35
// https://github.com/grpc/grpc-proto/blob/54713b1e8bc6ed2d4f25fb4dff527842150b91b2/grpc/service_config/service_config.proto#L247
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:42
type BalancerConfig struct {
	Name	string
	Config	externalserviceconfig.LoadBalancingConfig
}

type intermediateBalancerConfig []map[string]json.RawMessage

// MarshalJSON implements the json.Marshaler interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:49
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:49
// It marshals the balancer and config into a length-1 slice
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:49
// ([]map[string]config).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:53
func (bc *BalancerConfig) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:53
	_go_fuzz_dep_.CoverTab[68957]++
														if bc.Config == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:54
		_go_fuzz_dep_.CoverTab[68960]++

															return []byte(fmt.Sprintf(`[{%q: %v}]`, bc.Name, "{}")), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:56
		// _ = "end of CoverTab[68960]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:57
		_go_fuzz_dep_.CoverTab[68961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:57
		// _ = "end of CoverTab[68961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:57
	// _ = "end of CoverTab[68957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:57
	_go_fuzz_dep_.CoverTab[68958]++
														c, err := json.Marshal(bc.Config)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:59
		_go_fuzz_dep_.CoverTab[68962]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:60
		// _ = "end of CoverTab[68962]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:61
		_go_fuzz_dep_.CoverTab[68963]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:61
		// _ = "end of CoverTab[68963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:61
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:61
	// _ = "end of CoverTab[68958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:61
	_go_fuzz_dep_.CoverTab[68959]++
														return []byte(fmt.Sprintf(`[{%q: %s}]`, bc.Name, c)), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:62
	// _ = "end of CoverTab[68959]"
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
// ServiceConfig contains a list of loadBalancingConfigs, each with a name and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
// config. This method iterates through that list in order, and stops at the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
// first policy that is supported.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
//   - If the config for the first supported policy is invalid, the whole service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
//     config is invalid.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
//   - If the list doesn't contain any supported policy, the whole service config
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:65
//     is invalid.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:74
func (bc *BalancerConfig) UnmarshalJSON(b []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:74
	_go_fuzz_dep_.CoverTab[68964]++
														var ir intermediateBalancerConfig
														err := json.Unmarshal(b, &ir)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:77
		_go_fuzz_dep_.CoverTab[68967]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:78
		// _ = "end of CoverTab[68967]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:79
		_go_fuzz_dep_.CoverTab[68968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:79
		// _ = "end of CoverTab[68968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:79
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:79
	// _ = "end of CoverTab[68964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:79
	_go_fuzz_dep_.CoverTab[68965]++

														var names []string
														for i, lbcfg := range ir {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:82
		_go_fuzz_dep_.CoverTab[68969]++
															if len(lbcfg) != 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:83
			_go_fuzz_dep_.CoverTab[68975]++
																return fmt.Errorf("invalid loadBalancingConfig: entry %v does not contain exactly 1 policy/config pair: %q", i, lbcfg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:84
			// _ = "end of CoverTab[68975]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:85
			_go_fuzz_dep_.CoverTab[68976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:85
			// _ = "end of CoverTab[68976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:85
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:85
		// _ = "end of CoverTab[68969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:85
		_go_fuzz_dep_.CoverTab[68970]++

															var (
			name	string
			jsonCfg	json.RawMessage
		)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:93
		for name, jsonCfg = range lbcfg {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:93
			_go_fuzz_dep_.CoverTab[68977]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:93
			// _ = "end of CoverTab[68977]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:94
		// _ = "end of CoverTab[68970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:94
		_go_fuzz_dep_.CoverTab[68971]++

															names = append(names, name)
															builder := balancer.Get(name)
															if builder == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:98
			_go_fuzz_dep_.CoverTab[68978]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:101
			continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:101
			// _ = "end of CoverTab[68978]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:102
			_go_fuzz_dep_.CoverTab[68979]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:102
			// _ = "end of CoverTab[68979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:102
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:102
		// _ = "end of CoverTab[68971]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:102
		_go_fuzz_dep_.CoverTab[68972]++
															bc.Name = name

															parser, ok := builder.(balancer.ConfigParser)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:106
			_go_fuzz_dep_.CoverTab[68980]++
																if string(jsonCfg) != "{}" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:107
				_go_fuzz_dep_.CoverTab[68982]++
																	logger.Warningf("non-empty balancer configuration %q, but balancer does not implement ParseConfig", string(jsonCfg))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:108
				// _ = "end of CoverTab[68982]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:109
				_go_fuzz_dep_.CoverTab[68983]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:109
				// _ = "end of CoverTab[68983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:109
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:109
			// _ = "end of CoverTab[68980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:109
			_go_fuzz_dep_.CoverTab[68981]++

																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:111
			// _ = "end of CoverTab[68981]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:112
			_go_fuzz_dep_.CoverTab[68984]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:112
			// _ = "end of CoverTab[68984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:112
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:112
		// _ = "end of CoverTab[68972]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:112
		_go_fuzz_dep_.CoverTab[68973]++

															cfg, err := parser.ParseConfig(jsonCfg)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:115
			_go_fuzz_dep_.CoverTab[68985]++
																return fmt.Errorf("error parsing loadBalancingConfig for policy %q: %v", name, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:116
			// _ = "end of CoverTab[68985]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:117
			_go_fuzz_dep_.CoverTab[68986]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:117
			// _ = "end of CoverTab[68986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:117
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:117
		// _ = "end of CoverTab[68973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:117
		_go_fuzz_dep_.CoverTab[68974]++
															bc.Config = cfg
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:119
		// _ = "end of CoverTab[68974]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:120
	// _ = "end of CoverTab[68965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:120
	_go_fuzz_dep_.CoverTab[68966]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:125
	return fmt.Errorf("invalid loadBalancingConfig: no supported policies found in %v", names)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:125
	// _ = "end of CoverTab[68966]"
}

// MethodConfig defines the configuration recommended by the service providers for a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:128
// particular method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:130
type MethodConfig struct {
	// WaitForReady indicates whether RPCs sent to this method should wait until
	// the connection is ready by default (!failfast). The value specified via the
	// gRPC client API will override the value set here.
	WaitForReady	*bool
	// Timeout is the default timeout for RPCs sent to this method. The actual
	// deadline used will be the minimum of the value specified here and the value
	// set by the application via the gRPC client API.  If either one is not set,
	// then the other will be used.  If neither is set, then the RPC has no deadline.
	Timeout	*time.Duration
	// MaxReqSize is the maximum allowed payload size for an individual request in a
	// stream (client->server) in bytes. The size which is measured is the serialized
	// payload after per-message compression (but before stream compression) in bytes.
	// The actual value used is the minimum of the value specified here and the value set
	// by the application via the gRPC client API. If either one is not set, then the other
	// will be used.  If neither is set, then the built-in default is used.
	MaxReqSize	*int
	// MaxRespSize is the maximum allowed payload size for an individual response in a
	// stream (server->client) in bytes.
	MaxRespSize	*int
	// RetryPolicy configures retry options for the method.
	RetryPolicy	*RetryPolicy
}

// RetryPolicy defines the go-native version of the retry policy defined by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:154
// service config here:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:154
// https://github.com/grpc/proposal/blob/master/A6-client-retries.md#integration-with-service-config
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:157
type RetryPolicy struct {
	// MaxAttempts is the maximum number of attempts, including the original RPC.
	//
	// This field is required and must be two or greater.
	MaxAttempts	int

	// Exponential backoff parameters. The initial retry attempt will occur at
	// random(0, initialBackoff). In general, the nth attempt will occur at
	// random(0,
	//   min(initialBackoff*backoffMultiplier**(n-1), maxBackoff)).
	//
	// These fields are required and must be greater than zero.
	InitialBackoff		time.Duration
	MaxBackoff		time.Duration
	BackoffMultiplier	float64

	// The set of status codes which may be retried.
	//
	// Status codes are specified as strings, e.g., "UNAVAILABLE".
	//
	// This field is required and must be non-empty.
	// Note: a set is used to store this for easy lookup.
	RetryableStatusCodes	map[codes.Code]bool
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:180
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/serviceconfig/serviceconfig.go:180
var _ = _go_fuzz_dep_.CoverTab
