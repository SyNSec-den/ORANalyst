//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
package client

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:1
)

import (
	"encoding/json"
	"fmt"
	"log"
)

// Settings holds optional client settings.
type Settings struct {
	disablePAFXFast		bool
	assumePreAuthentication	bool
	preAuthEType		int32
	logger			*log.Logger
}

// jsonSettings is used when marshaling the Settings details to JSON format.
type jsonSettings struct {
	DisablePAFXFast		bool
	AssumePreAuthentication	bool
}

// NewSettings creates a new client settings struct.
func NewSettings(settings ...func(*Settings)) *Settings {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:24
	_go_fuzz_dep_.CoverTab[88761]++
												s := new(Settings)
												for _, set := range settings {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:26
		_go_fuzz_dep_.CoverTab[88763]++
													set(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:27
		// _ = "end of CoverTab[88763]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:28
	// _ = "end of CoverTab[88761]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:28
	_go_fuzz_dep_.CoverTab[88762]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:29
	// _ = "end of CoverTab[88762]"
}

// DisablePAFXFAST used to configure the client to not use PA_FX_FAST.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:32
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:32
// s := NewSettings(DisablePAFXFAST(true))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:35
func DisablePAFXFAST(b bool) func(*Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:35
	_go_fuzz_dep_.CoverTab[88764]++
												return func(s *Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:36
		_go_fuzz_dep_.CoverTab[88765]++
													s.disablePAFXFast = b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:37
		// _ = "end of CoverTab[88765]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:38
	// _ = "end of CoverTab[88764]"
}

// DisablePAFXFAST indicates is the client should disable the use of PA_FX_FAST.
func (s *Settings) DisablePAFXFAST() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:42
	_go_fuzz_dep_.CoverTab[88766]++
												return s.disablePAFXFast
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:43
	// _ = "end of CoverTab[88766]"
}

// AssumePreAuthentication used to configure the client to assume pre-authentication is required.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:46
// s := NewSettings(AssumePreAuthentication(true))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:49
func AssumePreAuthentication(b bool) func(*Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:49
	_go_fuzz_dep_.CoverTab[88767]++
												return func(s *Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:50
		_go_fuzz_dep_.CoverTab[88768]++
													s.assumePreAuthentication = b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:51
		// _ = "end of CoverTab[88768]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:52
	// _ = "end of CoverTab[88767]"
}

// AssumePreAuthentication indicates if the client should proactively assume using pre-authentication.
func (s *Settings) AssumePreAuthentication() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:56
	_go_fuzz_dep_.CoverTab[88769]++
												return s.assumePreAuthentication
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:57
	// _ = "end of CoverTab[88769]"
}

// Logger used to configure client with a logger.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:60
// s := NewSettings(kt, Logger(l))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:63
func Logger(l *log.Logger) func(*Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:63
	_go_fuzz_dep_.CoverTab[88770]++
												return func(s *Settings) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:64
		_go_fuzz_dep_.CoverTab[88771]++
													s.logger = l
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:65
		// _ = "end of CoverTab[88771]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:66
	// _ = "end of CoverTab[88770]"
}

// Logger returns the client logger instance.
func (s *Settings) Logger() *log.Logger {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:70
	_go_fuzz_dep_.CoverTab[88772]++
												return s.logger
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:71
	// _ = "end of CoverTab[88772]"
}

// Log will write to the service's logger if it is configured.
func (cl *Client) Log(format string, v ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:75
	_go_fuzz_dep_.CoverTab[88773]++
												if cl.settings.Logger() != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:76
		_go_fuzz_dep_.CoverTab[88774]++
													cl.settings.Logger().Output(2, fmt.Sprintf(format, v...))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:77
		// _ = "end of CoverTab[88774]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:78
		_go_fuzz_dep_.CoverTab[88775]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:78
		// _ = "end of CoverTab[88775]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:78
	// _ = "end of CoverTab[88773]"
}

// JSON returns a JSON representation of the settings.
func (s *Settings) JSON() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:82
	_go_fuzz_dep_.CoverTab[88776]++
												js := jsonSettings{
		DisablePAFXFast:		s.disablePAFXFast,
		AssumePreAuthentication:	s.assumePreAuthentication,
	}
	b, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:88
		_go_fuzz_dep_.CoverTab[88778]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:89
		// _ = "end of CoverTab[88778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:90
		_go_fuzz_dep_.CoverTab[88779]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:90
		// _ = "end of CoverTab[88779]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:90
	// _ = "end of CoverTab[88776]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:90
	_go_fuzz_dep_.CoverTab[88777]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:91
	// _ = "end of CoverTab[88777]"

}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/client/settings.go:93
var _ = _go_fuzz_dep_.CoverTab
