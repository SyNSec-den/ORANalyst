//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:19
// Package credentials defines APIs for parsing SPIFFE ID.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:19
// All APIs in this package are experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:22
)

import (
	"crypto/tls"
	"crypto/x509"
	"net/url"

	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("credentials")

// SPIFFEIDFromState parses the SPIFFE ID from State. If the SPIFFE ID format
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:34
// is invalid, return nil with warning.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:36
func SPIFFEIDFromState(state tls.ConnectionState) *url.URL {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:36
	_go_fuzz_dep_.CoverTab[62461]++
													if len(state.PeerCertificates) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:37
		_go_fuzz_dep_.CoverTab[62463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:37
		return len(state.PeerCertificates[0].URIs) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:37
		// _ = "end of CoverTab[62463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:37
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:37
		_go_fuzz_dep_.CoverTab[62464]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:38
		// _ = "end of CoverTab[62464]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:39
		_go_fuzz_dep_.CoverTab[62465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:39
		// _ = "end of CoverTab[62465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:39
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:39
	// _ = "end of CoverTab[62461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:39
	_go_fuzz_dep_.CoverTab[62462]++
													return SPIFFEIDFromCert(state.PeerCertificates[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:40
	// _ = "end of CoverTab[62462]"
}

// SPIFFEIDFromCert parses the SPIFFE ID from x509.Certificate. If the SPIFFE
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:43
// ID format is invalid, return nil with warning.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:45
func SPIFFEIDFromCert(cert *x509.Certificate) *url.URL {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:45
	_go_fuzz_dep_.CoverTab[62466]++
													if cert == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:46
		_go_fuzz_dep_.CoverTab[62469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:46
		return cert.URIs == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:46
		// _ = "end of CoverTab[62469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:46
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:46
		_go_fuzz_dep_.CoverTab[62470]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:47
		// _ = "end of CoverTab[62470]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:48
		_go_fuzz_dep_.CoverTab[62471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:48
		// _ = "end of CoverTab[62471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:48
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:48
	// _ = "end of CoverTab[62466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:48
	_go_fuzz_dep_.CoverTab[62467]++
													var spiffeID *url.URL
													for _, uri := range cert.URIs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:50
		_go_fuzz_dep_.CoverTab[62472]++
														if uri == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			_go_fuzz_dep_.CoverTab[62478]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			return uri.Scheme != "spiffe"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			// _ = "end of CoverTab[62478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			_go_fuzz_dep_.CoverTab[62479]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			return uri.Opaque != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			// _ = "end of CoverTab[62479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			_go_fuzz_dep_.CoverTab[62480]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			return (uri.User != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
				_go_fuzz_dep_.CoverTab[62481]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
				return uri.User.Username() != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
				// _ = "end of CoverTab[62481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			// _ = "end of CoverTab[62480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:51
			_go_fuzz_dep_.CoverTab[62482]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:52
			// _ = "end of CoverTab[62482]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:53
			_go_fuzz_dep_.CoverTab[62483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:53
			// _ = "end of CoverTab[62483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:53
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:53
		// _ = "end of CoverTab[62472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:53
		_go_fuzz_dep_.CoverTab[62473]++

														if len(uri.String()) > 2048 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:55
			_go_fuzz_dep_.CoverTab[62484]++
															logger.Warning("invalid SPIFFE ID: total ID length larger than 2048 bytes")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:57
			// _ = "end of CoverTab[62484]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:58
			_go_fuzz_dep_.CoverTab[62485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:58
			// _ = "end of CoverTab[62485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:58
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:58
		// _ = "end of CoverTab[62473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:58
		_go_fuzz_dep_.CoverTab[62474]++
														if len(uri.Host) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:59
			_go_fuzz_dep_.CoverTab[62486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:59
			return len(uri.Path) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:59
			// _ = "end of CoverTab[62486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:59
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:59
			_go_fuzz_dep_.CoverTab[62487]++
															logger.Warning("invalid SPIFFE ID: domain or workload ID is empty")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:61
			// _ = "end of CoverTab[62487]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:62
			_go_fuzz_dep_.CoverTab[62488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:62
			// _ = "end of CoverTab[62488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:62
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:62
		// _ = "end of CoverTab[62474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:62
		_go_fuzz_dep_.CoverTab[62475]++
														if len(uri.Host) > 255 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:63
			_go_fuzz_dep_.CoverTab[62489]++
															logger.Warning("invalid SPIFFE ID: domain length larger than 255 characters")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:65
			// _ = "end of CoverTab[62489]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:66
			_go_fuzz_dep_.CoverTab[62490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:66
			// _ = "end of CoverTab[62490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:66
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:66
		// _ = "end of CoverTab[62475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:66
		_go_fuzz_dep_.CoverTab[62476]++

														if len(cert.URIs) > 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:68
			_go_fuzz_dep_.CoverTab[62491]++
															logger.Warning("invalid SPIFFE ID: multiple URI SANs")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:70
			// _ = "end of CoverTab[62491]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:71
			_go_fuzz_dep_.CoverTab[62492]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:71
			// _ = "end of CoverTab[62492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:71
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:71
		// _ = "end of CoverTab[62476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:71
		_go_fuzz_dep_.CoverTab[62477]++
														spiffeID = uri
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:72
		// _ = "end of CoverTab[62477]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:73
	// _ = "end of CoverTab[62467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:73
	_go_fuzz_dep_.CoverTab[62468]++
													return spiffeID
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:74
	// _ = "end of CoverTab[62468]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:75
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/spiffe.go:75
var _ = _go_fuzz_dep_.CoverTab
