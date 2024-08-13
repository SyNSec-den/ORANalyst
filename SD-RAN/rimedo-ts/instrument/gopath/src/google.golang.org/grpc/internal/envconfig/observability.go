//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
package envconfig

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:19
)

import "os"

const (
	envObservabilityConfig		= "GRPC_GCP_OBSERVABILITY_CONFIG"
	envObservabilityConfigFile	= "GRPC_GCP_OBSERVABILITY_CONFIG_FILE"
)

var (
	// ObservabilityConfig is the json configuration for the gcp/observability
	// package specified directly in the envObservabilityConfig env var.
	ObservabilityConfig	= os.Getenv(envObservabilityConfig)
	// ObservabilityConfigFile is the json configuration for the
	// gcp/observability specified in a file with the location specified in
	// envObservabilityConfigFile env var.
	ObservabilityConfigFile	= os.Getenv(envObservabilityConfigFile)
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/envconfig/observability.go:36
var _ = _go_fuzz_dep_.CoverTab
