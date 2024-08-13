//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:19
)

import (
	"context"
	"io"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	istatus "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:34
// actions and unblock when there's a picker update.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:36
type pickerWrapper struct {
	mu		sync.Mutex
	done		bool
	blockingCh	chan struct{}
	picker		balancer.Picker
}

func newPickerWrapper() *pickerWrapper {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:43
	_go_fuzz_dep_.CoverTab[79476]++
											return &pickerWrapper{blockingCh: make(chan struct{})}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:44
	// _ = "end of CoverTab[79476]"
}

// updatePicker is called by UpdateBalancerState. It unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:48
	_go_fuzz_dep_.CoverTab[79477]++
											pw.mu.Lock()
											if pw.done {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:50
		_go_fuzz_dep_.CoverTab[79479]++
												pw.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:52
		// _ = "end of CoverTab[79479]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:53
		_go_fuzz_dep_.CoverTab[79480]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:53
		// _ = "end of CoverTab[79480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:53
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:53
	// _ = "end of CoverTab[79477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:53
	_go_fuzz_dep_.CoverTab[79478]++
											pw.picker = p

											close(pw.blockingCh)
											pw.blockingCh = make(chan struct{})
											pw.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:58
	// _ = "end of CoverTab[79478]"
}

// doneChannelzWrapper performs the following:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:61
//   - increments the calls started channelz counter
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:61
//   - wraps the done function in the passed in result to increment the calls
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:61
//     failed or calls succeeded channelz counter before invoking the actual
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:61
//     done function.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:66
func doneChannelzWrapper(acw *acBalancerWrapper, result *balancer.PickResult) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:66
	_go_fuzz_dep_.CoverTab[79481]++
											acw.mu.Lock()
											ac := acw.ac
											acw.mu.Unlock()
											ac.incrCallsStarted()
											done := result.Done
											result.Done = func(b balancer.DoneInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:72
		_go_fuzz_dep_.CoverTab[79482]++
												if b.Err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:73
			_go_fuzz_dep_.CoverTab[79484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:73
			return b.Err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:73
			// _ = "end of CoverTab[79484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:73
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:73
			_go_fuzz_dep_.CoverTab[79485]++
													ac.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:74
			// _ = "end of CoverTab[79485]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:75
			_go_fuzz_dep_.CoverTab[79486]++
													ac.incrCallsSucceeded()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:76
			// _ = "end of CoverTab[79486]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:77
		// _ = "end of CoverTab[79482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:77
		_go_fuzz_dep_.CoverTab[79483]++
												if done != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:78
			_go_fuzz_dep_.CoverTab[79487]++
													done(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:79
			// _ = "end of CoverTab[79487]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:80
			_go_fuzz_dep_.CoverTab[79488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:80
			// _ = "end of CoverTab[79488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:80
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:80
		// _ = "end of CoverTab[79483]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:81
	// _ = "end of CoverTab[79481]"
}

// pick returns the transport that will be used for the RPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// It may block in the following cases:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// - there's no picker
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// - the current picker returns ErrNoSubConnAvailable
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// - the current picker returns other errors and failfast is false.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// - the subConn returned by the current picker is not READY
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:84
// When one of these situations happens, pick blocks until the picker gets updated.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:91
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (transport.ClientTransport, balancer.PickResult, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:91
	_go_fuzz_dep_.CoverTab[79489]++
											var ch chan struct{}

											var lastPickErr error
											for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:95
		_go_fuzz_dep_.CoverTab[79490]++
												pw.mu.Lock()
												if pw.done {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:97
			_go_fuzz_dep_.CoverTab[79498]++
													pw.mu.Unlock()
													return nil, balancer.PickResult{}, ErrClientConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:99
			// _ = "end of CoverTab[79498]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:100
			_go_fuzz_dep_.CoverTab[79499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:100
			// _ = "end of CoverTab[79499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:100
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:100
		// _ = "end of CoverTab[79490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:100
		_go_fuzz_dep_.CoverTab[79491]++

												if pw.picker == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:102
			_go_fuzz_dep_.CoverTab[79500]++
													ch = pw.blockingCh
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:103
			// _ = "end of CoverTab[79500]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:104
			_go_fuzz_dep_.CoverTab[79501]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:104
			// _ = "end of CoverTab[79501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:104
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:104
		// _ = "end of CoverTab[79491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:104
		_go_fuzz_dep_.CoverTab[79492]++
												if ch == pw.blockingCh {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:105
			_go_fuzz_dep_.CoverTab[79502]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:109
			pw.mu.Unlock()
			select {
			case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:111
				_go_fuzz_dep_.CoverTab[79504]++
														var errStr string
														if lastPickErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:113
					_go_fuzz_dep_.CoverTab[79507]++
															errStr = "latest balancer error: " + lastPickErr.Error()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:114
					// _ = "end of CoverTab[79507]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:115
					_go_fuzz_dep_.CoverTab[79508]++
															errStr = ctx.Err().Error()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:116
					// _ = "end of CoverTab[79508]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:117
				// _ = "end of CoverTab[79504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:117
				_go_fuzz_dep_.CoverTab[79505]++
														switch ctx.Err() {
				case context.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:119
					_go_fuzz_dep_.CoverTab[79509]++
															return nil, balancer.PickResult{}, status.Error(codes.DeadlineExceeded, errStr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:120
					// _ = "end of CoverTab[79509]"
				case context.Canceled:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:121
					_go_fuzz_dep_.CoverTab[79510]++
															return nil, balancer.PickResult{}, status.Error(codes.Canceled, errStr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:122
					// _ = "end of CoverTab[79510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:122
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:122
					_go_fuzz_dep_.CoverTab[79511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:122
					// _ = "end of CoverTab[79511]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:123
				// _ = "end of CoverTab[79505]"
			case <-ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:124
				_go_fuzz_dep_.CoverTab[79506]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:124
				// _ = "end of CoverTab[79506]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:125
			// _ = "end of CoverTab[79502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:125
			_go_fuzz_dep_.CoverTab[79503]++
													continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:126
			// _ = "end of CoverTab[79503]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:127
			_go_fuzz_dep_.CoverTab[79512]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:127
			// _ = "end of CoverTab[79512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:127
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:127
		// _ = "end of CoverTab[79492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:127
		_go_fuzz_dep_.CoverTab[79493]++

												ch = pw.blockingCh
												p := pw.picker
												pw.mu.Unlock()

												pickResult, err := p.Pick(info)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:134
			_go_fuzz_dep_.CoverTab[79513]++
													if err == balancer.ErrNoSubConnAvailable {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:135
				_go_fuzz_dep_.CoverTab[79517]++
														continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:136
				// _ = "end of CoverTab[79517]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:137
				_go_fuzz_dep_.CoverTab[79518]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:137
				// _ = "end of CoverTab[79518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:137
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:137
			// _ = "end of CoverTab[79513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:137
			_go_fuzz_dep_.CoverTab[79514]++
													if st, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:138
				_go_fuzz_dep_.CoverTab[79519]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:141
				if istatus.IsRestrictedControlPlaneCode(st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:141
					_go_fuzz_dep_.CoverTab[79521]++
															err = status.Errorf(codes.Internal, "received picker error with illegal status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:142
					// _ = "end of CoverTab[79521]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:143
					_go_fuzz_dep_.CoverTab[79522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:143
					// _ = "end of CoverTab[79522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:143
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:143
				// _ = "end of CoverTab[79519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:143
				_go_fuzz_dep_.CoverTab[79520]++
														return nil, balancer.PickResult{}, dropError{error: err}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:144
				// _ = "end of CoverTab[79520]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:145
				_go_fuzz_dep_.CoverTab[79523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:145
				// _ = "end of CoverTab[79523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:145
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:145
			// _ = "end of CoverTab[79514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:145
			_go_fuzz_dep_.CoverTab[79515]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:148
			if !failfast {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:148
				_go_fuzz_dep_.CoverTab[79524]++
														lastPickErr = err
														continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:150
				// _ = "end of CoverTab[79524]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:151
				_go_fuzz_dep_.CoverTab[79525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:151
				// _ = "end of CoverTab[79525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:151
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:151
			// _ = "end of CoverTab[79515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:151
			_go_fuzz_dep_.CoverTab[79516]++
													return nil, balancer.PickResult{}, status.Error(codes.Unavailable, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:152
			// _ = "end of CoverTab[79516]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:153
			_go_fuzz_dep_.CoverTab[79526]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:153
			// _ = "end of CoverTab[79526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:153
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:153
		// _ = "end of CoverTab[79493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:153
		_go_fuzz_dep_.CoverTab[79494]++

												acw, ok := pickResult.SubConn.(*acBalancerWrapper)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:156
			_go_fuzz_dep_.CoverTab[79527]++
													logger.Errorf("subconn returned from pick is type %T, not *acBalancerWrapper", pickResult.SubConn)
													continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:158
			// _ = "end of CoverTab[79527]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:159
			_go_fuzz_dep_.CoverTab[79528]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:159
			// _ = "end of CoverTab[79528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:159
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:159
		// _ = "end of CoverTab[79494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:159
		_go_fuzz_dep_.CoverTab[79495]++
												if t := acw.getAddrConn().getReadyTransport(); t != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:160
			_go_fuzz_dep_.CoverTab[79529]++
													if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:161
				_go_fuzz_dep_.CoverTab[79531]++
														doneChannelzWrapper(acw, &pickResult)
														return t, pickResult, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:163
				// _ = "end of CoverTab[79531]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:164
				_go_fuzz_dep_.CoverTab[79532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:164
				// _ = "end of CoverTab[79532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:164
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:164
			// _ = "end of CoverTab[79529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:164
			_go_fuzz_dep_.CoverTab[79530]++
													return t, pickResult, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:165
			// _ = "end of CoverTab[79530]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:166
			_go_fuzz_dep_.CoverTab[79533]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:166
			// _ = "end of CoverTab[79533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:166
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:166
		// _ = "end of CoverTab[79495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:166
		_go_fuzz_dep_.CoverTab[79496]++
												if pickResult.Done != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:167
			_go_fuzz_dep_.CoverTab[79534]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:170
			pickResult.Done(balancer.DoneInfo{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:170
			// _ = "end of CoverTab[79534]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:171
			_go_fuzz_dep_.CoverTab[79535]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:171
			// _ = "end of CoverTab[79535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:171
		// _ = "end of CoverTab[79496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:171
		_go_fuzz_dep_.CoverTab[79497]++
												logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:172
		// _ = "end of CoverTab[79497]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:177
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:177
	// _ = "end of CoverTab[79489]"
}

func (pw *pickerWrapper) close() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:180
	_go_fuzz_dep_.CoverTab[79536]++
											pw.mu.Lock()
											defer pw.mu.Unlock()
											if pw.done {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:183
		_go_fuzz_dep_.CoverTab[79538]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:184
		// _ = "end of CoverTab[79538]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:185
		_go_fuzz_dep_.CoverTab[79539]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:185
		// _ = "end of CoverTab[79539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:185
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:185
	// _ = "end of CoverTab[79536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:185
	_go_fuzz_dep_.CoverTab[79537]++
											pw.done = true
											close(pw.blockingCh)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:187
	// _ = "end of CoverTab[79537]"
}

// dropError is a wrapper error that indicates the LB policy wishes to drop the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:190
// RPC and not retry it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:192
type dropError struct {
	error
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:194
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/picker_wrapper.go:194
var _ = _go_fuzz_dep_.CoverTab
