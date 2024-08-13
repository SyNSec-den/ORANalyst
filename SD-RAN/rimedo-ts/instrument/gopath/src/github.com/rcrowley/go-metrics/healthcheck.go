//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
package metrics

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:1
)

// Healthchecks hold an error value describing an arbitrary up/down status.
type Healthcheck interface {
	Check()
	Error() error
	Healthy()
	Unhealthy(error)
}

// NewHealthcheck constructs a new Healthcheck which will use the given
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:11
// function to update its status.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:13
func NewHealthcheck(f func(Healthcheck)) Healthcheck {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:13
	_go_fuzz_dep_.CoverTab[96222]++
															if UseNilMetrics {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:14
		_go_fuzz_dep_.CoverTab[96224]++
																return NilHealthcheck{}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:15
		// _ = "end of CoverTab[96224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:16
		_go_fuzz_dep_.CoverTab[96225]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:16
		// _ = "end of CoverTab[96225]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:16
	// _ = "end of CoverTab[96222]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:16
	_go_fuzz_dep_.CoverTab[96223]++
															return &StandardHealthcheck{nil, f}
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:17
	// _ = "end of CoverTab[96223]"
}

// NilHealthcheck is a no-op.
type NilHealthcheck struct{}

// Check is a no-op.
func (NilHealthcheck) Check()	{ _go_fuzz_dep_.CoverTab[96226]++; // _ = "end of CoverTab[96226]" }

// Error is a no-op.
func (NilHealthcheck) Error() error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:27
	_go_fuzz_dep_.CoverTab[96227]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:27
	return nil
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:27
	// _ = "end of CoverTab[96227]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:27
}

// Healthy is a no-op.
func (NilHealthcheck) Healthy()	{ _go_fuzz_dep_.CoverTab[96228]++; // _ = "end of CoverTab[96228]" }

// Unhealthy is a no-op.
func (NilHealthcheck) Unhealthy(error) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:33
	_go_fuzz_dep_.CoverTab[96229]++
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:33
	// _ = "end of CoverTab[96229]"
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:33
}

// StandardHealthcheck is the standard implementation of a Healthcheck and
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:35
// stores the status and a function to call to update the status.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:37
type StandardHealthcheck struct {
	err	error
	f	func(Healthcheck)
}

// Check runs the healthcheck function to update the healthcheck's status.
func (h *StandardHealthcheck) Check() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:43
	_go_fuzz_dep_.CoverTab[96230]++
															h.f(h)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:44
	// _ = "end of CoverTab[96230]"
}

// Error returns the healthcheck's status, which will be nil if it is healthy.
func (h *StandardHealthcheck) Error() error {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:48
	_go_fuzz_dep_.CoverTab[96231]++
															return h.err
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:49
	// _ = "end of CoverTab[96231]"
}

// Healthy marks the healthcheck as healthy.
func (h *StandardHealthcheck) Healthy() {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:53
	_go_fuzz_dep_.CoverTab[96232]++
															h.err = nil
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:54
	// _ = "end of CoverTab[96232]"
}

// Unhealthy marks the healthcheck as unhealthy.  The error is stored and
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:57
// may be retrieved by the Error method.
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:59
func (h *StandardHealthcheck) Unhealthy(err error) {
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:59
	_go_fuzz_dep_.CoverTab[96233]++
															h.err = err
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:60
	// _ = "end of CoverTab[96233]"
}

//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/rcrowley/go-metrics@v0.0.0-20201227073835-cf1acfcdf475/healthcheck.go:61
var _ = _go_fuzz_dep_.CoverTab
