// Copyright 2019-present Open Networking Foundation.
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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
package retry

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:15
)

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

// WithPerCallTimeout sets the per-call retry timeout
func WithPerCallTimeout(t time.Duration) CallOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:25
	_go_fuzz_dep_.CoverTab[182752]++
														return newCallOption(func(opts *callOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:26
		_go_fuzz_dep_.CoverTab[182753]++
															opts.perCallTimeout = &t
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:27
		// _ = "end of CoverTab[182753]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:28
	// _ = "end of CoverTab[182752]"
}

// WithInterval sets the base retry interval
func WithInterval(d time.Duration) CallOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:32
	_go_fuzz_dep_.CoverTab[182754]++
														return newCallOption(func(opts *callOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:33
		_go_fuzz_dep_.CoverTab[182755]++
															opts.initialInterval = &d
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:34
		// _ = "end of CoverTab[182755]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:35
	// _ = "end of CoverTab[182754]"
}

// WithMaxInterval sets the maximum retry interval
func WithMaxInterval(d time.Duration) CallOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:39
	_go_fuzz_dep_.CoverTab[182756]++
														return newCallOption(func(opts *callOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:40
		_go_fuzz_dep_.CoverTab[182757]++
															opts.maxInterval = &d
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:41
		// _ = "end of CoverTab[182757]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:42
	// _ = "end of CoverTab[182756]"
}

// WithRetryOn sets the codes on which to retry a request
func WithRetryOn(codes ...codes.Code) CallOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:46
	_go_fuzz_dep_.CoverTab[182758]++
														return newCallOption(func(opts *callOptions) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:47
		_go_fuzz_dep_.CoverTab[182759]++
															opts.codes = codes
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:48
		// _ = "end of CoverTab[182759]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:49
	// _ = "end of CoverTab[182758]"
}

func newCallOption(f func(opts *callOptions)) CallOption {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:52
	_go_fuzz_dep_.CoverTab[182760]++
														return CallOption{
		applyFunc: f,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:55
	// _ = "end of CoverTab[182760]"
}

// CallOption is a retrying interceptor call option
type CallOption struct {
	grpc.EmptyCallOption	// make sure we implement private after() and before() fields so we don't panic.
	applyFunc		func(opts *callOptions)
}

type callOptions struct {
	perCallTimeout	*time.Duration
	initialInterval	*time.Duration
	maxInterval	*time.Duration
	codes		[]codes.Code
}

func newCallContext(ctx context.Context, opts *callOptions) context.Context {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:71
	_go_fuzz_dep_.CoverTab[182761]++
														if opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:72
		_go_fuzz_dep_.CoverTab[182763]++
															ctx, _ = context.WithTimeout(ctx, *opts.perCallTimeout)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:73
		// _ = "end of CoverTab[182763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:74
		_go_fuzz_dep_.CoverTab[182764]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:74
		// _ = "end of CoverTab[182764]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:74
	// _ = "end of CoverTab[182761]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:74
	_go_fuzz_dep_.CoverTab[182762]++
														return ctx
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:75
	// _ = "end of CoverTab[182762]"
}

func newCallOptions(opts *callOptions, options []CallOption) *callOptions {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:78
	_go_fuzz_dep_.CoverTab[182765]++
														if len(options) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:79
		_go_fuzz_dep_.CoverTab[182768]++
															return opts
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:80
		// _ = "end of CoverTab[182768]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:81
		_go_fuzz_dep_.CoverTab[182769]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:81
		// _ = "end of CoverTab[182769]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:81
	// _ = "end of CoverTab[182765]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:81
	_go_fuzz_dep_.CoverTab[182766]++
														optCopy := &callOptions{}
														*optCopy = *opts
														for _, f := range options {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:84
		_go_fuzz_dep_.CoverTab[182770]++
															f.applyFunc(optCopy)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:85
		// _ = "end of CoverTab[182770]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:86
	// _ = "end of CoverTab[182766]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:86
	_go_fuzz_dep_.CoverTab[182767]++
														return optCopy
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:87
	// _ = "end of CoverTab[182767]"
}

func filterCallOptions(options []grpc.CallOption) (grpcOptions []grpc.CallOption, retryOptions []CallOption) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:90
	_go_fuzz_dep_.CoverTab[182771]++
														for _, opt := range options {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:91
		_go_fuzz_dep_.CoverTab[182773]++
															if co, ok := opt.(CallOption); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:92
			_go_fuzz_dep_.CoverTab[182774]++
																retryOptions = append(retryOptions, co)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:93
			// _ = "end of CoverTab[182774]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:94
			_go_fuzz_dep_.CoverTab[182775]++
																grpcOptions = append(grpcOptions, opt)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:95
			// _ = "end of CoverTab[182775]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:96
		// _ = "end of CoverTab[182773]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:97
	// _ = "end of CoverTab[182771]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:97
	_go_fuzz_dep_.CoverTab[182772]++
														return grpcOptions, retryOptions
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:98
	// _ = "end of CoverTab[182772]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/options.go:99
var _ = _go_fuzz_dep_.CoverTab
