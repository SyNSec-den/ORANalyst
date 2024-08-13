//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
package encoding

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:1
)

import (
	"sync"
)

// Encoder encodes the contents of v into a byte representation.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:7
// It's primarily used for encoding a map[string]interface{} into a file format.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:9
type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

const (
	// ErrEncoderNotFound is returned when there is no encoder registered for a format.
	ErrEncoderNotFound	= encodingError("encoder not found for this format")

	// ErrEncoderFormatAlreadyRegistered is returned when an encoder is already registered for a format.
	ErrEncoderFormatAlreadyRegistered	= encodingError("encoder already registered for this format")
)

// EncoderRegistry can choose an appropriate Encoder based on the provided format.
type EncoderRegistry struct {
	encoders	map[string]Encoder

	mu	sync.RWMutex
}

// NewEncoderRegistry returns a new, initialized EncoderRegistry.
func NewEncoderRegistry() *EncoderRegistry {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:29
	_go_fuzz_dep_.CoverTab[120791]++
													return &EncoderRegistry{
		encoders: make(map[string]Encoder),
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:32
	// _ = "end of CoverTab[120791]"
}

// RegisterEncoder registers an Encoder for a format.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:35
// Registering a Encoder for an already existing format is not supported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:37
func (e *EncoderRegistry) RegisterEncoder(format string, enc Encoder) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:37
	_go_fuzz_dep_.CoverTab[120792]++
													e.mu.Lock()
													defer e.mu.Unlock()

													if _, ok := e.encoders[format]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:41
		_go_fuzz_dep_.CoverTab[120794]++
														return ErrEncoderFormatAlreadyRegistered
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:42
		// _ = "end of CoverTab[120794]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:43
		_go_fuzz_dep_.CoverTab[120795]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:43
		// _ = "end of CoverTab[120795]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:43
	// _ = "end of CoverTab[120792]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:43
	_go_fuzz_dep_.CoverTab[120793]++

													e.encoders[format] = enc

													return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:47
	// _ = "end of CoverTab[120793]"
}

func (e *EncoderRegistry) Encode(format string, v interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:50
	_go_fuzz_dep_.CoverTab[120796]++
													e.mu.RLock()
													encoder, ok := e.encoders[format]
													e.mu.RUnlock()

													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:55
		_go_fuzz_dep_.CoverTab[120798]++
														return nil, ErrEncoderNotFound
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:56
		// _ = "end of CoverTab[120798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:57
		_go_fuzz_dep_.CoverTab[120799]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:57
		// _ = "end of CoverTab[120799]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:57
	// _ = "end of CoverTab[120796]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:57
	_go_fuzz_dep_.CoverTab[120797]++

													return encoder.Encode(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:59
	// _ = "end of CoverTab[120797]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:60
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/encoder.go:60
var _ = _go_fuzz_dep_.CoverTab
