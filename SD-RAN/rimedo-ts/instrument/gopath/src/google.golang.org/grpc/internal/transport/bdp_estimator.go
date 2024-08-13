//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:19
)

import (
	"sync"
	"time"
)

const (
	// bdpLimit is the maximum value the flow control windows will be increased
	// to.  TCP typically limits this to 4MB, but some systems go up to 16MB.
	// Since this is only a limit, it is safe to make it optimistic.
	bdpLimit	= (1 << 20) * 16
	// alpha is a constant factor used to keep a moving average
	// of RTTs.
	alpha	= 0.9
	// If the current bdp sample is greater than or equal to
	// our beta * our estimated bdp and the current bandwidth
	// sample is the maximum bandwidth observed so far, we
	// increase our bbp estimate by a factor of gamma.
	beta	= 0.66
	// To put our bdp to be smaller than or equal to twice the real BDP,
	// we should multiply our current sample with 4/3, however to round things out
	// we use 2 as the multiplication factor.
	gamma	= 2
)

// Adding arbitrary data to ping so that its ack can be identified.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:45
// Easter-egg: what does the ping message say?
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:47
var bdpPing = &ping{data: [8]byte{2, 4, 16, 16, 9, 14, 7, 7}}

type bdpEstimator struct {
	// sentAt is the time when the ping was sent.
	sentAt	time.Time

	mu	sync.Mutex
	// bdp is the current bdp estimate.
	bdp	uint32
	// sample is the number of bytes received in one measurement cycle.
	sample	uint32
	// bwMax is the maximum bandwidth noted so far (bytes/sec).
	bwMax	float64
	// bool to keep track of the beginning of a new measurement cycle.
	isSent	bool
	// Callback to update the window sizes.
	updateFlowControl	func(n uint32)
	// sampleCount is the number of samples taken so far.
	sampleCount	uint64
	// round trip time (seconds)
	rtt	float64
}

// timesnap registers the time bdp ping was sent out so that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:70
// network rtt can be calculated when its ack is received.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:70
// It is called (by controller) when the bdpPing is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:70
// being written on the wire.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:74
func (b *bdpEstimator) timesnap(d [8]byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:74
	_go_fuzz_dep_.CoverTab[76557]++
														if bdpPing.data != d {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:75
		_go_fuzz_dep_.CoverTab[76559]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:76
		// _ = "end of CoverTab[76559]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:77
		_go_fuzz_dep_.CoverTab[76560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:77
		// _ = "end of CoverTab[76560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:77
	// _ = "end of CoverTab[76557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:77
	_go_fuzz_dep_.CoverTab[76558]++
														b.sentAt = time.Now()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:78
	// _ = "end of CoverTab[76558]"
}

// add adds bytes to the current sample for calculating bdp.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:81
// It returns true only if a ping must be sent. This can be used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:81
// by the caller (handleData) to make decision about batching
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:81
// a window update with it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:85
func (b *bdpEstimator) add(n uint32) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:85
	_go_fuzz_dep_.CoverTab[76561]++
														b.mu.Lock()
														defer b.mu.Unlock()
														if b.bdp == bdpLimit {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:88
		_go_fuzz_dep_.CoverTab[76564]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:89
		// _ = "end of CoverTab[76564]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:90
		_go_fuzz_dep_.CoverTab[76565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:90
		// _ = "end of CoverTab[76565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:90
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:90
	// _ = "end of CoverTab[76561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:90
	_go_fuzz_dep_.CoverTab[76562]++
														if !b.isSent {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:91
		_go_fuzz_dep_.CoverTab[76566]++
															b.isSent = true
															b.sample = n
															b.sentAt = time.Time{}
															b.sampleCount++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:96
		// _ = "end of CoverTab[76566]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:97
		_go_fuzz_dep_.CoverTab[76567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:97
		// _ = "end of CoverTab[76567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:97
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:97
	// _ = "end of CoverTab[76562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:97
	_go_fuzz_dep_.CoverTab[76563]++
														b.sample += n
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:99
	// _ = "end of CoverTab[76563]"
}

// calculate is called when an ack for a bdp ping is received.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:102
// Here we calculate the current bdp and bandwidth sample and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:102
// decide if the flow control windows should go up.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:105
func (b *bdpEstimator) calculate(d [8]byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:105
	_go_fuzz_dep_.CoverTab[76568]++

														if bdpPing.data != d {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:107
		_go_fuzz_dep_.CoverTab[76573]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:108
		// _ = "end of CoverTab[76573]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:109
		_go_fuzz_dep_.CoverTab[76574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:109
		// _ = "end of CoverTab[76574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:109
	// _ = "end of CoverTab[76568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:109
	_go_fuzz_dep_.CoverTab[76569]++
														b.mu.Lock()
														rttSample := time.Since(b.sentAt).Seconds()
														if b.sampleCount < 10 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:112
		_go_fuzz_dep_.CoverTab[76575]++

															b.rtt += (rttSample - b.rtt) / float64(b.sampleCount)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:114
		// _ = "end of CoverTab[76575]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:115
		_go_fuzz_dep_.CoverTab[76576]++

															b.rtt += (rttSample - b.rtt) * float64(alpha)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:117
		// _ = "end of CoverTab[76576]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:118
	// _ = "end of CoverTab[76569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:118
	_go_fuzz_dep_.CoverTab[76570]++
														b.isSent = false

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:122
	bwCurrent := float64(b.sample) / (b.rtt * float64(1.5))
	if bwCurrent > b.bwMax {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:123
		_go_fuzz_dep_.CoverTab[76577]++
															b.bwMax = bwCurrent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:124
		// _ = "end of CoverTab[76577]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:125
		_go_fuzz_dep_.CoverTab[76578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:125
		// _ = "end of CoverTab[76578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:125
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:125
	// _ = "end of CoverTab[76570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:125
	_go_fuzz_dep_.CoverTab[76571]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
	if float64(b.sample) >= beta*float64(b.bdp) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		_go_fuzz_dep_.CoverTab[76579]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		return bwCurrent == b.bwMax
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		// _ = "end of CoverTab[76579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		_go_fuzz_dep_.CoverTab[76580]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		return b.bdp != bdpLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		// _ = "end of CoverTab[76580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:129
		_go_fuzz_dep_.CoverTab[76581]++
															sampleFloat := float64(b.sample)
															b.bdp = uint32(gamma * sampleFloat)
															if b.bdp > bdpLimit {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:132
			_go_fuzz_dep_.CoverTab[76583]++
																b.bdp = bdpLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:133
			// _ = "end of CoverTab[76583]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:134
			_go_fuzz_dep_.CoverTab[76584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:134
			// _ = "end of CoverTab[76584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:134
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:134
		// _ = "end of CoverTab[76581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:134
		_go_fuzz_dep_.CoverTab[76582]++
															bdp := b.bdp
															b.mu.Unlock()
															b.updateFlowControl(bdp)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:138
		// _ = "end of CoverTab[76582]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:139
		_go_fuzz_dep_.CoverTab[76585]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:139
		// _ = "end of CoverTab[76585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:139
	// _ = "end of CoverTab[76571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:139
	_go_fuzz_dep_.CoverTab[76572]++
														b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:140
	// _ = "end of CoverTab[76572]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/bdp_estimator.go:141
var _ = _go_fuzz_dep_.CoverTab
