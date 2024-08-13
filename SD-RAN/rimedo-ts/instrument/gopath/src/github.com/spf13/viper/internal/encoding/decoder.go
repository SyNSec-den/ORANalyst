//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
package encoding

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:1
)

import (
	"sync"
)

// Decoder decodes the contents of b into a v representation.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:7
// It's primarily used for decoding contents of a file into a map[string]interface{}.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:9
type Decoder interface {
	Decode(b []byte, v interface{}) error
}

const (
	// ErrDecoderNotFound is returned when there is no decoder registered for a format.
	ErrDecoderNotFound	= encodingError("decoder not found for this format")

	// ErrDecoderFormatAlreadyRegistered is returned when an decoder is already registered for a format.
	ErrDecoderFormatAlreadyRegistered	= encodingError("decoder already registered for this format")
)

// DecoderRegistry can choose an appropriate Decoder based on the provided format.
type DecoderRegistry struct {
	decoders	map[string]Decoder

	mu	sync.RWMutex
}

// NewDecoderRegistry returns a new, initialized DecoderRegistry.
func NewDecoderRegistry() *DecoderRegistry {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:29
	_go_fuzz_dep_.CoverTab[120782]++
													return &DecoderRegistry{
		decoders: make(map[string]Decoder),
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:32
	// _ = "end of CoverTab[120782]"
}

// RegisterDecoder registers a Decoder for a format.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:35
// Registering a Decoder for an already existing format is not supported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:37
func (e *DecoderRegistry) RegisterDecoder(format string, enc Decoder) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:37
	_go_fuzz_dep_.CoverTab[120783]++
													e.mu.Lock()
													defer e.mu.Unlock()

													if _, ok := e.decoders[format]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:41
		_go_fuzz_dep_.CoverTab[120785]++
														return ErrDecoderFormatAlreadyRegistered
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:42
		// _ = "end of CoverTab[120785]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:43
		_go_fuzz_dep_.CoverTab[120786]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:43
		// _ = "end of CoverTab[120786]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:43
	// _ = "end of CoverTab[120783]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:43
	_go_fuzz_dep_.CoverTab[120784]++

													e.decoders[format] = enc

													return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:47
	// _ = "end of CoverTab[120784]"
}

// Decode calls the underlying Decoder based on the format.
func (e *DecoderRegistry) Decode(format string, b []byte, v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:51
	_go_fuzz_dep_.CoverTab[120787]++
													e.mu.RLock()
													decoder, ok := e.decoders[format]
													e.mu.RUnlock()

													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:56
		_go_fuzz_dep_.CoverTab[120789]++
														return ErrDecoderNotFound
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:57
		// _ = "end of CoverTab[120789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:58
		_go_fuzz_dep_.CoverTab[120790]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:58
		// _ = "end of CoverTab[120790]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:58
	// _ = "end of CoverTab[120787]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:58
	_go_fuzz_dep_.CoverTab[120788]++

													return decoder.Decode(b, v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:60
	// _ = "end of CoverTab[120788]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/internal/encoding/decoder.go:61
var _ = _go_fuzz_dep_.CoverTab
