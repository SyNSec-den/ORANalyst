//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:19
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal"
	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

const maxInt = int(^uint(0) >> 1)

// MethodConfig defines the configuration recommended by the service providers for a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:38
// particular method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:38
// Deprecated: Users should not use this struct. Service config should be received
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:38
// through name resolver, as specified here
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:38
// https://github.com/grpc/grpc/blob/master/doc/service_config.md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:44
type MethodConfig = internalserviceconfig.MethodConfig

type lbConfig struct {
	name	string
	cfg	serviceconfig.LoadBalancingConfig
}

// ServiceConfig is provided by the service provider and contains parameters for how
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:51
// clients that connect to the service should behave.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:51
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:51
// Deprecated: Users should not use this struct. Service config should be received
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:51
// through name resolver, as specified here
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:51
// https://github.com/grpc/grpc/blob/master/doc/service_config.md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:57
type ServiceConfig struct {
	serviceconfig.Config

	// LB is the load balancer the service providers recommends.  This is
	// deprecated; lbConfigs is preferred.  If lbConfig and LB are both present,
	// lbConfig will be used.
	LB	*string

	// lbConfig is the service config's load balancing configuration.  If
	// lbConfig and LB are both present, lbConfig will be used.
	lbConfig	*lbConfig

	// Methods contains a map for the methods in this service.  If there is an
	// exact match for a method (i.e. /service/method) in the map, use the
	// corresponding MethodConfig.  If there's no exact match, look for the
	// default config for the service (/service/) and use the corresponding
	// MethodConfig if it exists.  Otherwise, the method has no MethodConfig to
	// use.
	Methods	map[string]MethodConfig

	// If a retryThrottlingPolicy is provided, gRPC will automatically throttle
	// retry attempts and hedged RPCs when the client’s ratio of failures to
	// successes exceeds a threshold.
	//
	// For each server name, the gRPC client will maintain a token_count which is
	// initially set to maxTokens, and can take values between 0 and maxTokens.
	//
	// Every outgoing RPC (regardless of service or method invoked) will change
	// token_count as follows:
	//
	//   - Every failed RPC will decrement the token_count by 1.
	//   - Every successful RPC will increment the token_count by tokenRatio.
	//
	// If token_count is less than or equal to maxTokens / 2, then RPCs will not
	// be retried and hedged RPCs will not be sent.
	retryThrottling	*retryThrottlingPolicy
	// healthCheckConfig must be set as one of the requirement to enable LB channel
	// health check.
	healthCheckConfig	*healthCheckConfig
	// rawJSONString stores service config json string that get parsed into
	// this service config struct.
	rawJSONString	string
}

// healthCheckConfig defines the go-native version of the LB channel health check config.
type healthCheckConfig struct {
	// serviceName is the service name to use in the health-checking request.
	ServiceName string
}

type jsonRetryPolicy struct {
	MaxAttempts		int
	InitialBackoff		string
	MaxBackoff		string
	BackoffMultiplier	float64
	RetryableStatusCodes	[]codes.Code
}

// retryThrottlingPolicy defines the go-native version of the retry throttling
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:115
// policy defined by the service config here:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:115
// https://github.com/grpc/proposal/blob/master/A6-client-retries.md#integration-with-service-config
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:118
type retryThrottlingPolicy struct {
	// The number of tokens starts at maxTokens. The token_count will always be
	// between 0 and maxTokens.
	//
	// This field is required and must be greater than zero.
	MaxTokens	float64
	// The amount of tokens to add on each successful RPC. Typically this will
	// be some number between 0 and 1, e.g., 0.1.
	//
	// This field is required and must be greater than zero. Up to 3 decimal
	// places are supported.
	TokenRatio	float64
}

func parseDuration(s *string) (*time.Duration, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:132
	_go_fuzz_dep_.CoverTab[80584]++
											if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:133
		_go_fuzz_dep_.CoverTab[80591]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:134
		// _ = "end of CoverTab[80591]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:135
		_go_fuzz_dep_.CoverTab[80592]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:135
		// _ = "end of CoverTab[80592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:135
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:135
	// _ = "end of CoverTab[80584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:135
	_go_fuzz_dep_.CoverTab[80585]++
											if !strings.HasSuffix(*s, "s") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:136
		_go_fuzz_dep_.CoverTab[80593]++
												return nil, fmt.Errorf("malformed duration %q", *s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:137
		// _ = "end of CoverTab[80593]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:138
		_go_fuzz_dep_.CoverTab[80594]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:138
		// _ = "end of CoverTab[80594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:138
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:138
	// _ = "end of CoverTab[80585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:138
	_go_fuzz_dep_.CoverTab[80586]++
											ss := strings.SplitN((*s)[:len(*s)-1], ".", 3)
											if len(ss) > 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:140
		_go_fuzz_dep_.CoverTab[80595]++
												return nil, fmt.Errorf("malformed duration %q", *s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:141
		// _ = "end of CoverTab[80595]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:142
		_go_fuzz_dep_.CoverTab[80596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:142
		// _ = "end of CoverTab[80596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:142
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:142
	// _ = "end of CoverTab[80586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:142
	_go_fuzz_dep_.CoverTab[80587]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:145
	hasDigits := false
	var d time.Duration
	if len(ss[0]) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:147
		_go_fuzz_dep_.CoverTab[80597]++
												i, err := strconv.ParseInt(ss[0], 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:149
			_go_fuzz_dep_.CoverTab[80599]++
													return nil, fmt.Errorf("malformed duration %q: %v", *s, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:150
			// _ = "end of CoverTab[80599]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:151
			_go_fuzz_dep_.CoverTab[80600]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:151
			// _ = "end of CoverTab[80600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:151
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:151
		// _ = "end of CoverTab[80597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:151
		_go_fuzz_dep_.CoverTab[80598]++
												d = time.Duration(i) * time.Second
												hasDigits = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:153
		// _ = "end of CoverTab[80598]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:154
		_go_fuzz_dep_.CoverTab[80601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:154
		// _ = "end of CoverTab[80601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:154
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:154
	// _ = "end of CoverTab[80587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:154
	_go_fuzz_dep_.CoverTab[80588]++
											if len(ss) == 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:155
		_go_fuzz_dep_.CoverTab[80602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:155
		return len(ss[1]) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:155
		// _ = "end of CoverTab[80602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:155
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:155
		_go_fuzz_dep_.CoverTab[80603]++
												if len(ss[1]) > 9 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:156
			_go_fuzz_dep_.CoverTab[80607]++
													return nil, fmt.Errorf("malformed duration %q", *s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:157
			// _ = "end of CoverTab[80607]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:158
			_go_fuzz_dep_.CoverTab[80608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:158
			// _ = "end of CoverTab[80608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:158
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:158
		// _ = "end of CoverTab[80603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:158
		_go_fuzz_dep_.CoverTab[80604]++
												f, err := strconv.ParseInt(ss[1], 10, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:160
			_go_fuzz_dep_.CoverTab[80609]++
													return nil, fmt.Errorf("malformed duration %q: %v", *s, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:161
			// _ = "end of CoverTab[80609]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:162
			_go_fuzz_dep_.CoverTab[80610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:162
			// _ = "end of CoverTab[80610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:162
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:162
		// _ = "end of CoverTab[80604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:162
		_go_fuzz_dep_.CoverTab[80605]++
												for i := 9; i > len(ss[1]); i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:163
			_go_fuzz_dep_.CoverTab[80611]++
													f *= 10
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:164
			// _ = "end of CoverTab[80611]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:165
		// _ = "end of CoverTab[80605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:165
		_go_fuzz_dep_.CoverTab[80606]++
												d += time.Duration(f)
												hasDigits = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:167
		// _ = "end of CoverTab[80606]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:168
		_go_fuzz_dep_.CoverTab[80612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:168
		// _ = "end of CoverTab[80612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:168
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:168
	// _ = "end of CoverTab[80588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:168
	_go_fuzz_dep_.CoverTab[80589]++
											if !hasDigits {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:169
		_go_fuzz_dep_.CoverTab[80613]++
												return nil, fmt.Errorf("malformed duration %q", *s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:170
		// _ = "end of CoverTab[80613]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:171
		_go_fuzz_dep_.CoverTab[80614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:171
		// _ = "end of CoverTab[80614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:171
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:171
	// _ = "end of CoverTab[80589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:171
	_go_fuzz_dep_.CoverTab[80590]++

											return &d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:173
	// _ = "end of CoverTab[80590]"
}

type jsonName struct {
	Service	string
	Method	string
}

var (
	errDuplicatedName		= errors.New("duplicated name")
	errEmptyServiceNonEmptyMethod	= errors.New("cannot combine empty 'service' and non-empty 'method'")
)

func (j jsonName) generatePath() (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:186
	_go_fuzz_dep_.CoverTab[80615]++
											if j.Service == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:187
		_go_fuzz_dep_.CoverTab[80618]++
												if j.Method != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:188
			_go_fuzz_dep_.CoverTab[80620]++
													return "", errEmptyServiceNonEmptyMethod
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:189
			// _ = "end of CoverTab[80620]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:190
			_go_fuzz_dep_.CoverTab[80621]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:190
			// _ = "end of CoverTab[80621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:190
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:190
		// _ = "end of CoverTab[80618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:190
		_go_fuzz_dep_.CoverTab[80619]++
												return "", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:191
		// _ = "end of CoverTab[80619]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:192
		_go_fuzz_dep_.CoverTab[80622]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:192
		// _ = "end of CoverTab[80622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:192
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:192
	// _ = "end of CoverTab[80615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:192
	_go_fuzz_dep_.CoverTab[80616]++
											res := "/" + j.Service + "/"
											if j.Method != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:194
		_go_fuzz_dep_.CoverTab[80623]++
												res += j.Method
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:195
		// _ = "end of CoverTab[80623]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:196
		_go_fuzz_dep_.CoverTab[80624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:196
		// _ = "end of CoverTab[80624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:196
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:196
	// _ = "end of CoverTab[80616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:196
	_go_fuzz_dep_.CoverTab[80617]++
											return res, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:197
	// _ = "end of CoverTab[80617]"
}

// TODO(lyuxuan): delete this struct after cleaning up old service config implementation.
type jsonMC struct {
	Name			*[]jsonName
	WaitForReady		*bool
	Timeout			*string
	MaxRequestMessageBytes	*int64
	MaxResponseMessageBytes	*int64
	RetryPolicy		*jsonRetryPolicy
}

// TODO(lyuxuan): delete this struct after cleaning up old service config implementation.
type jsonSC struct {
	LoadBalancingPolicy	*string
	LoadBalancingConfig	*internalserviceconfig.BalancerConfig
	MethodConfig		*[]jsonMC
	RetryThrottling		*retryThrottlingPolicy
	HealthCheckConfig	*healthCheckConfig
}

func init() {
	internal.ParseServiceConfig = parseServiceConfig
}
func parseServiceConfig(js string) *serviceconfig.ParseResult {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:222
	_go_fuzz_dep_.CoverTab[80625]++
											if len(js) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:223
		_go_fuzz_dep_.CoverTab[80632]++
												return &serviceconfig.ParseResult{Err: fmt.Errorf("no JSON service config provided")}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:224
		// _ = "end of CoverTab[80632]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:225
		_go_fuzz_dep_.CoverTab[80633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:225
		// _ = "end of CoverTab[80633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:225
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:225
	// _ = "end of CoverTab[80625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:225
	_go_fuzz_dep_.CoverTab[80626]++
											var rsc jsonSC
											err := json.Unmarshal([]byte(js), &rsc)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:228
		_go_fuzz_dep_.CoverTab[80634]++
												logger.Warningf("grpc: unmarshaling service config %s: %v", js, err)
												return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:230
		// _ = "end of CoverTab[80634]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:231
		_go_fuzz_dep_.CoverTab[80635]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:231
		// _ = "end of CoverTab[80635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:231
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:231
	// _ = "end of CoverTab[80626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:231
	_go_fuzz_dep_.CoverTab[80627]++
											sc := ServiceConfig{
		LB:			rsc.LoadBalancingPolicy,
		Methods:		make(map[string]MethodConfig),
		retryThrottling:	rsc.RetryThrottling,
		healthCheckConfig:	rsc.HealthCheckConfig,
		rawJSONString:		js,
	}
	if c := rsc.LoadBalancingConfig; c != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:239
		_go_fuzz_dep_.CoverTab[80636]++
												sc.lbConfig = &lbConfig{
			name:	c.Name,
			cfg:	c.Config,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:243
		// _ = "end of CoverTab[80636]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:244
		_go_fuzz_dep_.CoverTab[80637]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:244
		// _ = "end of CoverTab[80637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:244
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:244
	// _ = "end of CoverTab[80627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:244
	_go_fuzz_dep_.CoverTab[80628]++

											if rsc.MethodConfig == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:246
		_go_fuzz_dep_.CoverTab[80638]++
												return &serviceconfig.ParseResult{Config: &sc}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:247
		// _ = "end of CoverTab[80638]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:248
		_go_fuzz_dep_.CoverTab[80639]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:248
		// _ = "end of CoverTab[80639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:248
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:248
	// _ = "end of CoverTab[80628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:248
	_go_fuzz_dep_.CoverTab[80629]++

											paths := map[string]struct{}{}
											for _, m := range *rsc.MethodConfig {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:251
		_go_fuzz_dep_.CoverTab[80640]++
												if m.Name == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:252
			_go_fuzz_dep_.CoverTab[80646]++
													continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:253
			// _ = "end of CoverTab[80646]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:254
			_go_fuzz_dep_.CoverTab[80647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:254
			// _ = "end of CoverTab[80647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:254
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:254
		// _ = "end of CoverTab[80640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:254
		_go_fuzz_dep_.CoverTab[80641]++
												d, err := parseDuration(m.Timeout)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:256
			_go_fuzz_dep_.CoverTab[80648]++
													logger.Warningf("grpc: unmarshaling service config %s: %v", js, err)
													return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:258
			// _ = "end of CoverTab[80648]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:259
			_go_fuzz_dep_.CoverTab[80649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:259
			// _ = "end of CoverTab[80649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:259
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:259
		// _ = "end of CoverTab[80641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:259
		_go_fuzz_dep_.CoverTab[80642]++

												mc := MethodConfig{
			WaitForReady:	m.WaitForReady,
			Timeout:	d,
		}
		if mc.RetryPolicy, err = convertRetryPolicy(m.RetryPolicy); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:265
			_go_fuzz_dep_.CoverTab[80650]++
													logger.Warningf("grpc: unmarshaling service config %s: %v", js, err)
													return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:267
			// _ = "end of CoverTab[80650]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:268
			_go_fuzz_dep_.CoverTab[80651]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:268
			// _ = "end of CoverTab[80651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:268
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:268
		// _ = "end of CoverTab[80642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:268
		_go_fuzz_dep_.CoverTab[80643]++
												if m.MaxRequestMessageBytes != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:269
			_go_fuzz_dep_.CoverTab[80652]++
													if *m.MaxRequestMessageBytes > int64(maxInt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:270
				_go_fuzz_dep_.CoverTab[80653]++
														mc.MaxReqSize = newInt(maxInt)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:271
				// _ = "end of CoverTab[80653]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:272
				_go_fuzz_dep_.CoverTab[80654]++
														mc.MaxReqSize = newInt(int(*m.MaxRequestMessageBytes))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:273
				// _ = "end of CoverTab[80654]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:274
			// _ = "end of CoverTab[80652]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:275
			_go_fuzz_dep_.CoverTab[80655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:275
			// _ = "end of CoverTab[80655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:275
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:275
		// _ = "end of CoverTab[80643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:275
		_go_fuzz_dep_.CoverTab[80644]++
												if m.MaxResponseMessageBytes != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:276
			_go_fuzz_dep_.CoverTab[80656]++
													if *m.MaxResponseMessageBytes > int64(maxInt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:277
				_go_fuzz_dep_.CoverTab[80657]++
														mc.MaxRespSize = newInt(maxInt)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:278
				// _ = "end of CoverTab[80657]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:279
				_go_fuzz_dep_.CoverTab[80658]++
														mc.MaxRespSize = newInt(int(*m.MaxResponseMessageBytes))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:280
				// _ = "end of CoverTab[80658]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:281
			// _ = "end of CoverTab[80656]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:282
			_go_fuzz_dep_.CoverTab[80659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:282
			// _ = "end of CoverTab[80659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:282
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:282
		// _ = "end of CoverTab[80644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:282
		_go_fuzz_dep_.CoverTab[80645]++
												for i, n := range *m.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:283
			_go_fuzz_dep_.CoverTab[80660]++
													path, err := n.generatePath()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:285
				_go_fuzz_dep_.CoverTab[80663]++
														logger.Warningf("grpc: error unmarshaling service config %s due to methodConfig[%d]: %v", js, i, err)
														return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:287
				// _ = "end of CoverTab[80663]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:288
				_go_fuzz_dep_.CoverTab[80664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:288
				// _ = "end of CoverTab[80664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:288
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:288
			// _ = "end of CoverTab[80660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:288
			_go_fuzz_dep_.CoverTab[80661]++

													if _, ok := paths[path]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:290
				_go_fuzz_dep_.CoverTab[80665]++
														err = errDuplicatedName
														logger.Warningf("grpc: error unmarshaling service config %s due to methodConfig[%d]: %v", js, i, err)
														return &serviceconfig.ParseResult{Err: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:293
				// _ = "end of CoverTab[80665]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:294
				_go_fuzz_dep_.CoverTab[80666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:294
				// _ = "end of CoverTab[80666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:294
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:294
			// _ = "end of CoverTab[80661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:294
			_go_fuzz_dep_.CoverTab[80662]++
													paths[path] = struct{}{}
													sc.Methods[path] = mc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:296
			// _ = "end of CoverTab[80662]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:297
		// _ = "end of CoverTab[80645]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:298
	// _ = "end of CoverTab[80629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:298
	_go_fuzz_dep_.CoverTab[80630]++

											if sc.retryThrottling != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:300
		_go_fuzz_dep_.CoverTab[80667]++
												if mt := sc.retryThrottling.MaxTokens; mt <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:301
			_go_fuzz_dep_.CoverTab[80669]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:301
			return mt > 1000
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:301
			// _ = "end of CoverTab[80669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:301
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:301
			_go_fuzz_dep_.CoverTab[80670]++
													return &serviceconfig.ParseResult{Err: fmt.Errorf("invalid retry throttling config: maxTokens (%v) out of range (0, 1000]", mt)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:302
			// _ = "end of CoverTab[80670]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:303
			_go_fuzz_dep_.CoverTab[80671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:303
			// _ = "end of CoverTab[80671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:303
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:303
		// _ = "end of CoverTab[80667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:303
		_go_fuzz_dep_.CoverTab[80668]++
												if tr := sc.retryThrottling.TokenRatio; tr <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:304
			_go_fuzz_dep_.CoverTab[80672]++
													return &serviceconfig.ParseResult{Err: fmt.Errorf("invalid retry throttling config: tokenRatio (%v) may not be negative", tr)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:305
			// _ = "end of CoverTab[80672]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:306
			_go_fuzz_dep_.CoverTab[80673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:306
			// _ = "end of CoverTab[80673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:306
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:306
		// _ = "end of CoverTab[80668]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:307
		_go_fuzz_dep_.CoverTab[80674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:307
		// _ = "end of CoverTab[80674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:307
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:307
	// _ = "end of CoverTab[80630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:307
	_go_fuzz_dep_.CoverTab[80631]++
											return &serviceconfig.ParseResult{Config: &sc}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:308
	// _ = "end of CoverTab[80631]"
}

func convertRetryPolicy(jrp *jsonRetryPolicy) (p *internalserviceconfig.RetryPolicy, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:311
	_go_fuzz_dep_.CoverTab[80675]++
											if jrp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:312
		_go_fuzz_dep_.CoverTab[80682]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:313
		// _ = "end of CoverTab[80682]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:314
		_go_fuzz_dep_.CoverTab[80683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:314
		// _ = "end of CoverTab[80683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:314
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:314
	// _ = "end of CoverTab[80675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:314
	_go_fuzz_dep_.CoverTab[80676]++
											ib, err := parseDuration(&jrp.InitialBackoff)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:316
		_go_fuzz_dep_.CoverTab[80684]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:317
		// _ = "end of CoverTab[80684]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:318
		_go_fuzz_dep_.CoverTab[80685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:318
		// _ = "end of CoverTab[80685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:318
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:318
	// _ = "end of CoverTab[80676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:318
	_go_fuzz_dep_.CoverTab[80677]++
											mb, err := parseDuration(&jrp.MaxBackoff)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:320
		_go_fuzz_dep_.CoverTab[80686]++
												return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:321
		// _ = "end of CoverTab[80686]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:322
		_go_fuzz_dep_.CoverTab[80687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:322
		// _ = "end of CoverTab[80687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:322
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:322
	// _ = "end of CoverTab[80677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:322
	_go_fuzz_dep_.CoverTab[80678]++

											if jrp.MaxAttempts <= 1 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:324
		_go_fuzz_dep_.CoverTab[80688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:324
		return *ib <= 0
												// _ = "end of CoverTab[80688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:325
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:325
		_go_fuzz_dep_.CoverTab[80689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:325
		return *mb <= 0
												// _ = "end of CoverTab[80689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:326
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:326
		_go_fuzz_dep_.CoverTab[80690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:326
		return jrp.BackoffMultiplier <= 0
												// _ = "end of CoverTab[80690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:327
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:327
		_go_fuzz_dep_.CoverTab[80691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:327
		return len(jrp.RetryableStatusCodes) == 0
												// _ = "end of CoverTab[80691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:328
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:328
		_go_fuzz_dep_.CoverTab[80692]++
												logger.Warningf("grpc: ignoring retry policy %v due to illegal configuration", jrp)
												return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:330
		// _ = "end of CoverTab[80692]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:331
		_go_fuzz_dep_.CoverTab[80693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:331
		// _ = "end of CoverTab[80693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:331
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:331
	// _ = "end of CoverTab[80678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:331
	_go_fuzz_dep_.CoverTab[80679]++

											rp := &internalserviceconfig.RetryPolicy{
		MaxAttempts:		jrp.MaxAttempts,
		InitialBackoff:		*ib,
		MaxBackoff:		*mb,
		BackoffMultiplier:	jrp.BackoffMultiplier,
		RetryableStatusCodes:	make(map[codes.Code]bool),
	}
	if rp.MaxAttempts > 5 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:340
		_go_fuzz_dep_.CoverTab[80694]++

												rp.MaxAttempts = 5
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:342
		// _ = "end of CoverTab[80694]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:343
		_go_fuzz_dep_.CoverTab[80695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:343
		// _ = "end of CoverTab[80695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:343
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:343
	// _ = "end of CoverTab[80679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:343
	_go_fuzz_dep_.CoverTab[80680]++
											for _, code := range jrp.RetryableStatusCodes {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:344
		_go_fuzz_dep_.CoverTab[80696]++
												rp.RetryableStatusCodes[code] = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:345
		// _ = "end of CoverTab[80696]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:346
	// _ = "end of CoverTab[80680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:346
	_go_fuzz_dep_.CoverTab[80681]++
											return rp, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:347
	// _ = "end of CoverTab[80681]"
}

func min(a, b *int) *int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:350
	_go_fuzz_dep_.CoverTab[80697]++
											if *a < *b {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:351
		_go_fuzz_dep_.CoverTab[80699]++
												return a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:352
		// _ = "end of CoverTab[80699]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:353
		_go_fuzz_dep_.CoverTab[80700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:353
		// _ = "end of CoverTab[80700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:353
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:353
	// _ = "end of CoverTab[80697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:353
	_go_fuzz_dep_.CoverTab[80698]++
											return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:354
	// _ = "end of CoverTab[80698]"
}

func getMaxSize(mcMax, doptMax *int, defaultVal int) *int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:357
	_go_fuzz_dep_.CoverTab[80701]++
											if mcMax == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:358
		_go_fuzz_dep_.CoverTab[80705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:358
		return doptMax == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:358
		// _ = "end of CoverTab[80705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:358
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:358
		_go_fuzz_dep_.CoverTab[80706]++
												return &defaultVal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:359
		// _ = "end of CoverTab[80706]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:360
		_go_fuzz_dep_.CoverTab[80707]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:360
		// _ = "end of CoverTab[80707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:360
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:360
	// _ = "end of CoverTab[80701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:360
	_go_fuzz_dep_.CoverTab[80702]++
											if mcMax != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:361
		_go_fuzz_dep_.CoverTab[80708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:361
		return doptMax != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:361
		// _ = "end of CoverTab[80708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:361
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:361
		_go_fuzz_dep_.CoverTab[80709]++
												return min(mcMax, doptMax)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:362
		// _ = "end of CoverTab[80709]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:363
		_go_fuzz_dep_.CoverTab[80710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:363
		// _ = "end of CoverTab[80710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:363
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:363
	// _ = "end of CoverTab[80702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:363
	_go_fuzz_dep_.CoverTab[80703]++
											if mcMax != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:364
		_go_fuzz_dep_.CoverTab[80711]++
												return mcMax
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:365
		// _ = "end of CoverTab[80711]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:366
		_go_fuzz_dep_.CoverTab[80712]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:366
		// _ = "end of CoverTab[80712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:366
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:366
	// _ = "end of CoverTab[80703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:366
	_go_fuzz_dep_.CoverTab[80704]++
											return doptMax
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:367
	// _ = "end of CoverTab[80704]"
}

func newInt(b int) *int {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:370
	_go_fuzz_dep_.CoverTab[80713]++
											return &b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:371
	// _ = "end of CoverTab[80713]"
}

func init() {
	internal.EqualServiceConfigForTesting = equalServiceConfig
}

// equalServiceConfig compares two configs. The rawJSONString field is ignored,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:378
// because they may diff in white spaces.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:378
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:378
// If any of them is NOT *ServiceConfig, return false.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:382
func equalServiceConfig(a, b serviceconfig.Config) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:382
	_go_fuzz_dep_.CoverTab[80714]++
											if a == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:383
		_go_fuzz_dep_.CoverTab[80719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:383
		return b == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:383
		// _ = "end of CoverTab[80719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:383
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:383
		_go_fuzz_dep_.CoverTab[80720]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:384
		// _ = "end of CoverTab[80720]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:385
		_go_fuzz_dep_.CoverTab[80721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:385
		// _ = "end of CoverTab[80721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:385
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:385
	// _ = "end of CoverTab[80714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:385
	_go_fuzz_dep_.CoverTab[80715]++
											aa, ok := a.(*ServiceConfig)
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:387
		_go_fuzz_dep_.CoverTab[80722]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:388
		// _ = "end of CoverTab[80722]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:389
		_go_fuzz_dep_.CoverTab[80723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:389
		// _ = "end of CoverTab[80723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:389
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:389
	// _ = "end of CoverTab[80715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:389
	_go_fuzz_dep_.CoverTab[80716]++
											bb, ok := b.(*ServiceConfig)
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:391
		_go_fuzz_dep_.CoverTab[80724]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:392
		// _ = "end of CoverTab[80724]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:393
		_go_fuzz_dep_.CoverTab[80725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:393
		// _ = "end of CoverTab[80725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:393
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:393
	// _ = "end of CoverTab[80716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:393
	_go_fuzz_dep_.CoverTab[80717]++
											aaRaw := aa.rawJSONString
											aa.rawJSONString = ""
											bbRaw := bb.rawJSONString
											bb.rawJSONString = ""
											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:398
		_go_fuzz_dep_.CoverTab[80726]++
												aa.rawJSONString = aaRaw
												bb.rawJSONString = bbRaw
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:400
		// _ = "end of CoverTab[80726]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:401
	// _ = "end of CoverTab[80717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:401
	_go_fuzz_dep_.CoverTab[80718]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:405
	return reflect.DeepEqual(aa, bb)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:405
	// _ = "end of CoverTab[80718]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:406
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/service_config.go:406
var _ = _go_fuzz_dep_.CoverTab
