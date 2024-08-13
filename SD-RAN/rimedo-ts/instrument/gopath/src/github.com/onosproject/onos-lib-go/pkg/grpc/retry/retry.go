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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
package retry

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:15
)

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var log = logging.GetLogger("onos", "grpc", "retry")

var defaultOptions = &callOptions{
	codes: []codes.Code{
		codes.Unavailable,
		codes.Unknown,
	},
}

// RetryingUnaryClientInterceptor returns a UnaryClientInterceptor that retries requests
func RetryingUnaryClientInterceptor(callOpts ...CallOption) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:41
	_go_fuzz_dep_.CoverTab[182776]++
														connOpts := newCallOptions(defaultOptions, callOpts)
														return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:43
		_go_fuzz_dep_.CoverTab[182777]++
															grpcOpts, retryOpts := filterCallOptions(opts)
															callOpts := newCallOptions(connOpts, retryOpts)
															b := backoff.NewExponentialBackOff()
															if callOpts.initialInterval != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:47
			_go_fuzz_dep_.CoverTab[182780]++
																b.InitialInterval = *callOpts.initialInterval
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:48
			// _ = "end of CoverTab[182780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:49
			_go_fuzz_dep_.CoverTab[182781]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:49
			// _ = "end of CoverTab[182781]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:49
		// _ = "end of CoverTab[182777]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:49
		_go_fuzz_dep_.CoverTab[182778]++
															if callOpts.maxInterval != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:50
			_go_fuzz_dep_.CoverTab[182782]++
																b.MaxInterval = *callOpts.maxInterval
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:51
			// _ = "end of CoverTab[182782]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:52
			_go_fuzz_dep_.CoverTab[182783]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:52
			// _ = "end of CoverTab[182783]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:52
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:52
		// _ = "end of CoverTab[182778]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:52
		_go_fuzz_dep_.CoverTab[182779]++
															return backoff.Retry(func() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:53
			_go_fuzz_dep_.CoverTab[182784]++
																log.Debugf("SendMsg %.250s", req)
																callCtx := newCallContext(ctx, callOpts)
																if err := invoker(callCtx, method, req, reply, cc, grpcOpts...); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:56
				_go_fuzz_dep_.CoverTab[182786]++
																	if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:57
					_go_fuzz_dep_.CoverTab[182789]++
																		if ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:58
						_go_fuzz_dep_.CoverTab[182790]++
																			log.Debugf("SendMsg %.250s: error", req, err)
																			return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:60
						// _ = "end of CoverTab[182790]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:61
						_go_fuzz_dep_.CoverTab[182791]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:61
						if callOpts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:61
							_go_fuzz_dep_.CoverTab[182792]++
																				log.Debugf("SendMsg %.250s: error", req, err)
																				return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:63
							// _ = "end of CoverTab[182792]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
							_go_fuzz_dep_.CoverTab[182793]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
							// _ = "end of CoverTab[182793]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
						}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
						// _ = "end of CoverTab[182791]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:64
					// _ = "end of CoverTab[182789]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:65
					_go_fuzz_dep_.CoverTab[182794]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:65
					// _ = "end of CoverTab[182794]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:65
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:65
				// _ = "end of CoverTab[182786]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:65
				_go_fuzz_dep_.CoverTab[182787]++
																	if isRetryable(callOpts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:66
					_go_fuzz_dep_.CoverTab[182795]++
																		log.Debugf("SendMsg %.250s: error", req, err)
																		return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:68
					// _ = "end of CoverTab[182795]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:69
					_go_fuzz_dep_.CoverTab[182796]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:69
					// _ = "end of CoverTab[182796]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:69
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:69
				// _ = "end of CoverTab[182787]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:69
				_go_fuzz_dep_.CoverTab[182788]++
																	log.Warnf("SendMsg %.250s: error", req, err)
																	return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:71
				// _ = "end of CoverTab[182788]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:72
				_go_fuzz_dep_.CoverTab[182797]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:72
				// _ = "end of CoverTab[182797]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:72
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:72
			// _ = "end of CoverTab[182784]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:72
			_go_fuzz_dep_.CoverTab[182785]++
																log.Debugf("RecvMsg %.250s", reply)
																return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:74
			// _ = "end of CoverTab[182785]"
		}, b)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:75
		// _ = "end of CoverTab[182779]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:76
	// _ = "end of CoverTab[182776]"
}

// RetryingStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func RetryingStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:80
	_go_fuzz_dep_.CoverTab[182798]++
														return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:81
		_go_fuzz_dep_.CoverTab[182799]++
															if desc.ClientStreams && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:82
			_go_fuzz_dep_.CoverTab[182801]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:82
			return desc.ServerStreams
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:82
			// _ = "end of CoverTab[182801]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:82
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:82
			_go_fuzz_dep_.CoverTab[182802]++
																return newBiDirectionalStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:83
			// _ = "end of CoverTab[182802]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:84
			_go_fuzz_dep_.CoverTab[182803]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:84
			if desc.ClientStreams {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:84
				_go_fuzz_dep_.CoverTab[182804]++
																	return newClientStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:85
				// _ = "end of CoverTab[182804]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:86
				_go_fuzz_dep_.CoverTab[182805]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:86
				if desc.ServerStreams {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:86
					_go_fuzz_dep_.CoverTab[182806]++
																		return newServerStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:87
					// _ = "end of CoverTab[182806]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
					_go_fuzz_dep_.CoverTab[182807]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
					// _ = "end of CoverTab[182807]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
				// _ = "end of CoverTab[182805]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
			// _ = "end of CoverTab[182803]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
		// _ = "end of CoverTab[182799]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:88
		_go_fuzz_dep_.CoverTab[182800]++
															panic("Invalid StreamDesc")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:89
		// _ = "end of CoverTab[182800]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:90
	// _ = "end of CoverTab[182798]"
}

// newClientStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newClientStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:94
	_go_fuzz_dep_.CoverTab[182808]++
														connOpts := newCallOptions(defaultOptions, callOpts)
														return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:96
		_go_fuzz_dep_.CoverTab[182809]++
															grpcOpts, retryOpts := filterCallOptions(opts)
															callOpts := newCallOptions(connOpts, retryOpts)
															stream := &retryingClientStream{
			ctx:	ctx,
			buffer:	&retryingClientStreamBuffer{},
			opts:	callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:103
				_go_fuzz_dep_.CoverTab[182811]++
																	return streamer(ctx, desc, cc, method, grpcOpts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:104
				// _ = "end of CoverTab[182811]"
			},
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:106
		// _ = "end of CoverTab[182809]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:106
		_go_fuzz_dep_.CoverTab[182810]++
															return stream, stream.retryStream()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:107
		// _ = "end of CoverTab[182810]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:108
	// _ = "end of CoverTab[182808]"
}

// newServerStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newServerStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:112
	_go_fuzz_dep_.CoverTab[182812]++
														connOpts := newCallOptions(defaultOptions, callOpts)
														return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:114
		_go_fuzz_dep_.CoverTab[182813]++
															grpcOpts, retryOpts := filterCallOptions(opts)
															callOpts := newCallOptions(connOpts, retryOpts)
															stream := &retryingClientStream{
			ctx:	ctx,
			buffer:	&retryingServerStreamBuffer{},
			opts:	callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:121
				_go_fuzz_dep_.CoverTab[182815]++
																	return streamer(ctx, desc, cc, method, grpcOpts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:122
				// _ = "end of CoverTab[182815]"
			},
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:124
		// _ = "end of CoverTab[182813]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:124
		_go_fuzz_dep_.CoverTab[182814]++
															return stream, stream.retryStream()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:125
		// _ = "end of CoverTab[182814]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:126
	// _ = "end of CoverTab[182812]"
}

// newBiDirectionalStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newBiDirectionalStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:130
	_go_fuzz_dep_.CoverTab[182816]++
														connOpts := newCallOptions(defaultOptions, callOpts)
														return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:132
		_go_fuzz_dep_.CoverTab[182817]++
															grpcOpts, retryOpts := filterCallOptions(opts)
															callOpts := newCallOptions(connOpts, retryOpts)
															stream := &retryingClientStream{
			ctx:	ctx,
			buffer:	&retryingBiDirectionalStreamBuffer{},
			opts:	callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:139
				_go_fuzz_dep_.CoverTab[182819]++
																	return streamer(ctx, desc, cc, method, grpcOpts...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:140
				// _ = "end of CoverTab[182819]"
			},
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:142
		// _ = "end of CoverTab[182817]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:142
		_go_fuzz_dep_.CoverTab[182818]++
															return stream, stream.retryStream()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:143
		// _ = "end of CoverTab[182818]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:144
	// _ = "end of CoverTab[182816]"
}

type retryingStreamBuffer interface {
	append(interface{})
	list() []interface{}
}

type retryingClientStreamBuffer struct {
	buffer	[]interface{}
	mu	sync.RWMutex
}

func (b *retryingClientStreamBuffer) append(msg interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:157
	_go_fuzz_dep_.CoverTab[182820]++
														b.mu.Lock()
														b.buffer = append(b.buffer, msg)
														b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:160
	// _ = "end of CoverTab[182820]"
}

func (b *retryingClientStreamBuffer) list() []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:163
	_go_fuzz_dep_.CoverTab[182821]++
														b.mu.RLock()
														buffer := make([]interface{}, len(b.buffer))
														copy(buffer, b.buffer)
														b.mu.RUnlock()
														return buffer
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:168
	// _ = "end of CoverTab[182821]"
}

type retryingServerStreamBuffer struct {
	msg	interface{}
	mu	sync.RWMutex
}

func (b *retryingServerStreamBuffer) append(msg interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:176
	_go_fuzz_dep_.CoverTab[182822]++
														b.mu.Lock()
														b.msg = msg
														b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:179
	// _ = "end of CoverTab[182822]"
}

func (b *retryingServerStreamBuffer) list() []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:182
	_go_fuzz_dep_.CoverTab[182823]++
														b.mu.RLock()
														msg := b.msg
														b.mu.RUnlock()
														if msg != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:186
		_go_fuzz_dep_.CoverTab[182825]++
															return []interface{}{msg}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:187
		// _ = "end of CoverTab[182825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:188
		_go_fuzz_dep_.CoverTab[182826]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:188
		// _ = "end of CoverTab[182826]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:188
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:188
	// _ = "end of CoverTab[182823]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:188
	_go_fuzz_dep_.CoverTab[182824]++
														return []interface{}{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:189
	// _ = "end of CoverTab[182824]"
}

type retryingBiDirectionalStreamBuffer struct{}

func (b *retryingBiDirectionalStreamBuffer) append(interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:194
	_go_fuzz_dep_.CoverTab[182827]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:194
	// _ = "end of CoverTab[182827]"

}

func (b *retryingBiDirectionalStreamBuffer) list() []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:198
	_go_fuzz_dep_.CoverTab[182828]++
														return []interface{}{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:199
	// _ = "end of CoverTab[182828]"
}

type retryingClientStream struct {
	ctx		context.Context
	stream		grpc.ClientStream
	opts		*callOptions
	mu		sync.RWMutex
	buffer		retryingStreamBuffer
	newStream	func(ctx context.Context) (grpc.ClientStream, error)
	closed		bool
}

func (s *retryingClientStream) getStream() grpc.ClientStream {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:212
	_go_fuzz_dep_.CoverTab[182829]++
														s.mu.RLock()
														defer s.mu.RUnlock()
														return s.stream
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:215
	// _ = "end of CoverTab[182829]"
}

func (s *retryingClientStream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:218
	_go_fuzz_dep_.CoverTab[182830]++
														return s.ctx
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:219
	// _ = "end of CoverTab[182830]"
}

func (s *retryingClientStream) CloseSend() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:222
	_go_fuzz_dep_.CoverTab[182831]++
														log.Debug("CloseSend")
														s.mu.Lock()
														s.closed = true
														s.mu.Unlock()
														if err := s.getStream().CloseSend(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:227
		_go_fuzz_dep_.CoverTab[182833]++
															log.Warn("CloseSend: error", err)
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:229
		// _ = "end of CoverTab[182833]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:230
		_go_fuzz_dep_.CoverTab[182834]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:230
		// _ = "end of CoverTab[182834]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:230
	// _ = "end of CoverTab[182831]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:230
	_go_fuzz_dep_.CoverTab[182832]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:231
	// _ = "end of CoverTab[182832]"
}

func (s *retryingClientStream) Header() (metadata.MD, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:234
	_go_fuzz_dep_.CoverTab[182835]++
														return s.getStream().Header()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:235
	// _ = "end of CoverTab[182835]"
}

func (s *retryingClientStream) Trailer() metadata.MD {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:238
	_go_fuzz_dep_.CoverTab[182836]++
														return s.getStream().Trailer()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:239
	// _ = "end of CoverTab[182836]"
}

func (s *retryingClientStream) SendMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:242
	_go_fuzz_dep_.CoverTab[182837]++
														log.Debugf("SendMsg %.250s", m)
														return s.retrySendMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:244
	// _ = "end of CoverTab[182837]"
}

func (s *retryingClientStream) retrySendMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:247
	_go_fuzz_dep_.CoverTab[182838]++
														return backoff.RetryNotify(func() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:248
		_go_fuzz_dep_.CoverTab[182839]++
															return s.trySendMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:249
		// _ = "end of CoverTab[182839]"
	}, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:250
		_go_fuzz_dep_.CoverTab[182840]++
															log.Debugf("SendMsg %.250s: retry after %.250s", m, duration, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:251
		// _ = "end of CoverTab[182840]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:252
	// _ = "end of CoverTab[182838]"
}

func (s *retryingClientStream) trySendMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:255
	_go_fuzz_dep_.CoverTab[182841]++
														err := s.getStream().SendMsg(m)
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:257
		_go_fuzz_dep_.CoverTab[182845]++
															s.buffer.append(m)
															return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:259
		// _ = "end of CoverTab[182845]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:260
		_go_fuzz_dep_.CoverTab[182846]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:260
		// _ = "end of CoverTab[182846]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:260
	// _ = "end of CoverTab[182841]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:260
	_go_fuzz_dep_.CoverTab[182842]++
														if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:261
		_go_fuzz_dep_.CoverTab[182847]++
															if s.ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:262
			_go_fuzz_dep_.CoverTab[182848]++
																log.Debugf("SendMsg %.250s: error", m, err)
																return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:264
			// _ = "end of CoverTab[182848]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:265
			_go_fuzz_dep_.CoverTab[182849]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:265
			if s.opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:265
				_go_fuzz_dep_.CoverTab[182850]++
																	log.Debugf("SendMsg %.250s: error", m, err)
																	if err := s.tryStream(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:267
					_go_fuzz_dep_.CoverTab[182852]++
																		log.Debug("SendMsg %.250s: error", m, err)
																		return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:269
					// _ = "end of CoverTab[182852]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:270
					_go_fuzz_dep_.CoverTab[182853]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:270
					// _ = "end of CoverTab[182853]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:270
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:270
				// _ = "end of CoverTab[182850]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:270
				_go_fuzz_dep_.CoverTab[182851]++
																	return s.trySendMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:271
				// _ = "end of CoverTab[182851]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
				_go_fuzz_dep_.CoverTab[182854]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
				// _ = "end of CoverTab[182854]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
			// _ = "end of CoverTab[182849]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:272
		// _ = "end of CoverTab[182847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:273
		_go_fuzz_dep_.CoverTab[182855]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:273
		// _ = "end of CoverTab[182855]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:273
	// _ = "end of CoverTab[182842]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:273
	_go_fuzz_dep_.CoverTab[182843]++
														if isRetryable(s.opts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:274
		_go_fuzz_dep_.CoverTab[182856]++
															log.Debugf("SendMsg %.250s: error", m, err)
															if err := s.tryStream(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:276
			_go_fuzz_dep_.CoverTab[182858]++
																log.Debug("SendMsg %.250s: error", m, err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:278
			// _ = "end of CoverTab[182858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:279
			_go_fuzz_dep_.CoverTab[182859]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:279
			// _ = "end of CoverTab[182859]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:279
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:279
		// _ = "end of CoverTab[182856]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:279
		_go_fuzz_dep_.CoverTab[182857]++
															return s.trySendMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:280
		// _ = "end of CoverTab[182857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:281
		_go_fuzz_dep_.CoverTab[182860]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:281
		// _ = "end of CoverTab[182860]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:281
	// _ = "end of CoverTab[182843]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:281
	_go_fuzz_dep_.CoverTab[182844]++
														log.Warnf("SendMsg %.250s: error", m, err)
														return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:283
	// _ = "end of CoverTab[182844]"
}

func (s *retryingClientStream) RecvMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:286
	_go_fuzz_dep_.CoverTab[182861]++
														return s.retryRecvMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:287
	// _ = "end of CoverTab[182861]"
}

func (s *retryingClientStream) retryRecvMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:290
	_go_fuzz_dep_.CoverTab[182862]++
														return backoff.RetryNotify(func() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:291
		_go_fuzz_dep_.CoverTab[182863]++
															return s.tryRecvMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:292
		// _ = "end of CoverTab[182863]"
	}, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:293
		_go_fuzz_dep_.CoverTab[182864]++
															log.Debugf("RecvMsg: retry after %s", duration, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:294
		// _ = "end of CoverTab[182864]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:295
	// _ = "end of CoverTab[182862]"
}

func (s *retryingClientStream) tryRecvMsg(m interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:298
	_go_fuzz_dep_.CoverTab[182865]++
														err := s.getStream().RecvMsg(m)
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:300
		_go_fuzz_dep_.CoverTab[182870]++
															log.Debugf("RecvMsg %.250s", m)
															return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:302
		// _ = "end of CoverTab[182870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:303
		_go_fuzz_dep_.CoverTab[182871]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:303
		// _ = "end of CoverTab[182871]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:303
	// _ = "end of CoverTab[182865]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:303
	_go_fuzz_dep_.CoverTab[182866]++
														if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:304
		_go_fuzz_dep_.CoverTab[182872]++
															log.Debug("RecvMsg: EOF")
															return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:306
		// _ = "end of CoverTab[182872]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:307
		_go_fuzz_dep_.CoverTab[182873]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:307
		// _ = "end of CoverTab[182873]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:307
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:307
	// _ = "end of CoverTab[182866]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:307
	_go_fuzz_dep_.CoverTab[182867]++
														if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:308
		_go_fuzz_dep_.CoverTab[182874]++
															if s.ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:309
			_go_fuzz_dep_.CoverTab[182875]++
																log.Debug("RecvMsg: error", err)
																return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:311
			// _ = "end of CoverTab[182875]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:312
			_go_fuzz_dep_.CoverTab[182876]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:312
			if s.opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:312
				_go_fuzz_dep_.CoverTab[182877]++
																	log.Debug("RecvMsg: error", err)
																	if err := s.tryStream(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:314
					_go_fuzz_dep_.CoverTab[182879]++
																		log.Debug("RecvMsg: error", err)
																		return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:316
					// _ = "end of CoverTab[182879]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:317
					_go_fuzz_dep_.CoverTab[182880]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:317
					// _ = "end of CoverTab[182880]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:317
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:317
				// _ = "end of CoverTab[182877]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:317
				_go_fuzz_dep_.CoverTab[182878]++
																	return s.tryRecvMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:318
				// _ = "end of CoverTab[182878]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
				_go_fuzz_dep_.CoverTab[182881]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
				// _ = "end of CoverTab[182881]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
			// _ = "end of CoverTab[182876]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:319
		// _ = "end of CoverTab[182874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:320
		_go_fuzz_dep_.CoverTab[182882]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:320
		// _ = "end of CoverTab[182882]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:320
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:320
	// _ = "end of CoverTab[182867]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:320
	_go_fuzz_dep_.CoverTab[182868]++
														if isRetryable(s.opts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:321
		_go_fuzz_dep_.CoverTab[182883]++
															log.Debug("RecvMsg: error", err)
															if err := s.tryStream(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:323
			_go_fuzz_dep_.CoverTab[182885]++
																log.Debug("RecvMsg: error", err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:325
			// _ = "end of CoverTab[182885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:326
			_go_fuzz_dep_.CoverTab[182886]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:326
			// _ = "end of CoverTab[182886]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:326
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:326
		// _ = "end of CoverTab[182883]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:326
		_go_fuzz_dep_.CoverTab[182884]++
															return s.tryRecvMsg(m)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:327
		// _ = "end of CoverTab[182884]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:328
		_go_fuzz_dep_.CoverTab[182887]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:328
		// _ = "end of CoverTab[182887]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:328
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:328
	// _ = "end of CoverTab[182868]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:328
	_go_fuzz_dep_.CoverTab[182869]++
														log.Warn("RecvMsg: error", err)
														return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:330
	// _ = "end of CoverTab[182869]"
}

func (s *retryingClientStream) retryStream() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:333
	_go_fuzz_dep_.CoverTab[182888]++
														b := backoff.NewExponentialBackOff()
														if s.opts.initialInterval != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:335
		_go_fuzz_dep_.CoverTab[182891]++
															b.InitialInterval = *s.opts.initialInterval
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:336
		// _ = "end of CoverTab[182891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:337
		_go_fuzz_dep_.CoverTab[182892]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:337
		// _ = "end of CoverTab[182892]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:337
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:337
	// _ = "end of CoverTab[182888]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:337
	_go_fuzz_dep_.CoverTab[182889]++
														if s.opts.maxInterval != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:338
		_go_fuzz_dep_.CoverTab[182893]++
															b.MaxInterval = *s.opts.maxInterval
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:339
		// _ = "end of CoverTab[182893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:340
		_go_fuzz_dep_.CoverTab[182894]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:340
		// _ = "end of CoverTab[182894]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:340
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:340
	// _ = "end of CoverTab[182889]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:340
	_go_fuzz_dep_.CoverTab[182890]++
														return backoff.RetryNotify(s.tryStream, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:341
		_go_fuzz_dep_.CoverTab[182895]++
															log.Debugf("Stream: retry after %s", duration, err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:342
		// _ = "end of CoverTab[182895]"
	})
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:343
	// _ = "end of CoverTab[182890]"
}

func (s *retryingClientStream) tryStream() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:346
	_go_fuzz_dep_.CoverTab[182896]++
														s.mu.Lock()
														defer s.mu.Unlock()

														stream, err := s.newStream(newCallContext(s.ctx, s.opts))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:351
		_go_fuzz_dep_.CoverTab[182900]++
															if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:352
			_go_fuzz_dep_.CoverTab[182903]++
																if s.ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:353
				_go_fuzz_dep_.CoverTab[182904]++
																	log.Debug("Stream: error", err)
																	return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:355
				// _ = "end of CoverTab[182904]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:356
				_go_fuzz_dep_.CoverTab[182905]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:356
				if s.opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:356
					_go_fuzz_dep_.CoverTab[182906]++
																		log.Debug("Stream: error", err)
																		return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:358
					// _ = "end of CoverTab[182906]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
					_go_fuzz_dep_.CoverTab[182907]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
					// _ = "end of CoverTab[182907]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
				// _ = "end of CoverTab[182905]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:359
			// _ = "end of CoverTab[182903]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:360
			_go_fuzz_dep_.CoverTab[182908]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:360
			// _ = "end of CoverTab[182908]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:360
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:360
		// _ = "end of CoverTab[182900]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:360
		_go_fuzz_dep_.CoverTab[182901]++
															if isRetryable(s.opts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:361
			_go_fuzz_dep_.CoverTab[182909]++
																log.Debug("Stream: error", err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:363
			// _ = "end of CoverTab[182909]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:364
			_go_fuzz_dep_.CoverTab[182910]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:364
			// _ = "end of CoverTab[182910]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:364
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:364
		// _ = "end of CoverTab[182901]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:364
		_go_fuzz_dep_.CoverTab[182902]++
															log.Warn("Stream: error", err)
															return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:366
		// _ = "end of CoverTab[182902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:367
		_go_fuzz_dep_.CoverTab[182911]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:367
		// _ = "end of CoverTab[182911]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:367
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:367
	// _ = "end of CoverTab[182896]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:367
	_go_fuzz_dep_.CoverTab[182897]++

														msgs := s.buffer.list()
														for _, m := range msgs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:370
		_go_fuzz_dep_.CoverTab[182912]++
															log.Debugf("SendMsg %.250s", m)
															if err := stream.SendMsg(m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:372
			_go_fuzz_dep_.CoverTab[182913]++
																if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:373
				_go_fuzz_dep_.CoverTab[182916]++
																	if s.ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:374
					_go_fuzz_dep_.CoverTab[182917]++
																		log.Debugf("SendMsg %.250s: error", m, err)
																		return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:376
					// _ = "end of CoverTab[182917]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:377
					_go_fuzz_dep_.CoverTab[182918]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:377
					if s.opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:377
						_go_fuzz_dep_.CoverTab[182919]++
																			log.Debugf("SendMsg %.250s: error", m, err)
																			return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:379
						// _ = "end of CoverTab[182919]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
						_go_fuzz_dep_.CoverTab[182920]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
						// _ = "end of CoverTab[182920]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
					// _ = "end of CoverTab[182918]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:380
				// _ = "end of CoverTab[182916]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:381
				_go_fuzz_dep_.CoverTab[182921]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:381
				// _ = "end of CoverTab[182921]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:381
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:381
			// _ = "end of CoverTab[182913]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:381
			_go_fuzz_dep_.CoverTab[182914]++
																if isRetryable(s.opts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:382
				_go_fuzz_dep_.CoverTab[182922]++
																	log.Debugf("SendMsg %.250s: error", m, err)
																	return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:384
				// _ = "end of CoverTab[182922]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:385
				_go_fuzz_dep_.CoverTab[182923]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:385
				// _ = "end of CoverTab[182923]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:385
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:385
			// _ = "end of CoverTab[182914]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:385
			_go_fuzz_dep_.CoverTab[182915]++
																log.Warnf("SendMsg %.250s: error", m, err)
																return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:387
			// _ = "end of CoverTab[182915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:388
			_go_fuzz_dep_.CoverTab[182924]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:388
			// _ = "end of CoverTab[182924]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:388
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:388
		// _ = "end of CoverTab[182912]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:389
	// _ = "end of CoverTab[182897]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:389
	_go_fuzz_dep_.CoverTab[182898]++

														if s.closed {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:391
		_go_fuzz_dep_.CoverTab[182925]++
															log.Debug("CloseSend")
															if err := stream.CloseSend(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:393
			_go_fuzz_dep_.CoverTab[182926]++
																if isContextError(err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:394
				_go_fuzz_dep_.CoverTab[182929]++
																	if s.ctx.Err() != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:395
					_go_fuzz_dep_.CoverTab[182930]++
																		log.Debug("CloseSend: error", err)
																		return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:397
					// _ = "end of CoverTab[182930]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:398
					_go_fuzz_dep_.CoverTab[182931]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:398
					if s.opts.perCallTimeout != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:398
						_go_fuzz_dep_.CoverTab[182932]++
																			log.Debug("CloseSend: error", err)
																			return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:400
						// _ = "end of CoverTab[182932]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
						_go_fuzz_dep_.CoverTab[182933]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
						// _ = "end of CoverTab[182933]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
					// _ = "end of CoverTab[182931]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:401
				// _ = "end of CoverTab[182929]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:402
				_go_fuzz_dep_.CoverTab[182934]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:402
				// _ = "end of CoverTab[182934]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:402
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:402
			// _ = "end of CoverTab[182926]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:402
			_go_fuzz_dep_.CoverTab[182927]++
																if isRetryable(s.opts, err) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:403
				_go_fuzz_dep_.CoverTab[182935]++
																	log.Debug("CloseSend: error", err)
																	return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:405
				// _ = "end of CoverTab[182935]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:406
				_go_fuzz_dep_.CoverTab[182936]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:406
				// _ = "end of CoverTab[182936]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:406
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:406
			// _ = "end of CoverTab[182927]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:406
			_go_fuzz_dep_.CoverTab[182928]++
																log.Warn("CloseSend: error", err)
																return backoff.Permanent(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:408
			// _ = "end of CoverTab[182928]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:409
			_go_fuzz_dep_.CoverTab[182937]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:409
			// _ = "end of CoverTab[182937]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:409
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:409
		// _ = "end of CoverTab[182925]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:410
		_go_fuzz_dep_.CoverTab[182938]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:410
		// _ = "end of CoverTab[182938]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:410
	// _ = "end of CoverTab[182898]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:410
	_go_fuzz_dep_.CoverTab[182899]++
														s.stream = stream
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:412
	// _ = "end of CoverTab[182899]"
}

func isContextError(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:415
	_go_fuzz_dep_.CoverTab[182939]++
														code := status.Code(err)
														return code == codes.DeadlineExceeded || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:417
		_go_fuzz_dep_.CoverTab[182940]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:417
		return code == codes.Canceled
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:417
		// _ = "end of CoverTab[182940]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:417
	}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:417
	// _ = "end of CoverTab[182939]"
}

func isRetryable(opts *callOptions, err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:420
	_go_fuzz_dep_.CoverTab[182941]++
														code := status.Code(err)
														if code == codes.Canceled || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:422
		_go_fuzz_dep_.CoverTab[182944]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:422
		return code == codes.DeadlineExceeded
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:422
		// _ = "end of CoverTab[182944]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:422
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:422
		_go_fuzz_dep_.CoverTab[182945]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:423
		// _ = "end of CoverTab[182945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:424
		_go_fuzz_dep_.CoverTab[182946]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:424
		// _ = "end of CoverTab[182946]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:424
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:424
	// _ = "end of CoverTab[182941]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:424
	_go_fuzz_dep_.CoverTab[182942]++
														for _, retryableCode := range opts.codes {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:425
		_go_fuzz_dep_.CoverTab[182947]++
															if code == retryableCode {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:426
			_go_fuzz_dep_.CoverTab[182948]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:427
			// _ = "end of CoverTab[182948]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:428
			_go_fuzz_dep_.CoverTab[182949]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:428
			// _ = "end of CoverTab[182949]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:428
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:428
		// _ = "end of CoverTab[182947]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:429
	// _ = "end of CoverTab[182942]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:429
	_go_fuzz_dep_.CoverTab[182943]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:430
	// _ = "end of CoverTab[182943]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:431
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/grpc/retry/retry.go:431
var _ = _go_fuzz_dep_.CoverTab
