//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:19
)

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const proxyAuthHeaderKey = "Proxy-Authorization"

var (
	// The following variable will be overwritten in the tests.
	httpProxyFromEnvironment = http.ProxyFromEnvironment
)

func mapAddress(address string) (*url.URL, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:40
	_go_fuzz_dep_.CoverTab[78555]++
													req := &http.Request{
		URL: &url.URL{
			Scheme:	"https",
			Host:	address,
		},
	}
	url, err := httpProxyFromEnvironment(req)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:48
		_go_fuzz_dep_.CoverTab[78557]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:49
		// _ = "end of CoverTab[78557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:50
		_go_fuzz_dep_.CoverTab[78558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:50
		// _ = "end of CoverTab[78558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:50
	// _ = "end of CoverTab[78555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:50
	_go_fuzz_dep_.CoverTab[78556]++
													return url, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:51
	// _ = "end of CoverTab[78556]"
}

// To read a response from a net.Conn, http.ReadResponse() takes a bufio.Reader.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:54
// It's possible that this reader reads more than what's need for the response and stores
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:54
// those bytes in the buffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:54
// bufConn wraps the original net.Conn and the bufio.Reader to make sure we don't lose the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:54
// bytes in the buffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:59
type bufConn struct {
	net.Conn
	r	io.Reader
}

func (c *bufConn) Read(b []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:64
	_go_fuzz_dep_.CoverTab[78559]++
													return c.r.Read(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:65
	// _ = "end of CoverTab[78559]"
}

func basicAuth(username, password string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:68
	_go_fuzz_dep_.CoverTab[78560]++
													auth := username + ":" + password
													return base64.StdEncoding.EncodeToString([]byte(auth))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:70
	// _ = "end of CoverTab[78560]"
}

func doHTTPConnectHandshake(ctx context.Context, conn net.Conn, backendAddr string, proxyURL *url.URL, grpcUA string) (_ net.Conn, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:73
	_go_fuzz_dep_.CoverTab[78561]++
													defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:74
		_go_fuzz_dep_.CoverTab[78567]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:75
			_go_fuzz_dep_.CoverTab[78568]++
															conn.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:76
			// _ = "end of CoverTab[78568]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:77
			_go_fuzz_dep_.CoverTab[78569]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:77
			// _ = "end of CoverTab[78569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:77
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:77
		// _ = "end of CoverTab[78567]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:78
	// _ = "end of CoverTab[78561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:78
	_go_fuzz_dep_.CoverTab[78562]++

													req := &http.Request{
		Method:	http.MethodConnect,
		URL:	&url.URL{Host: backendAddr},
		Header:	map[string][]string{"User-Agent": {grpcUA}},
	}
	if t := proxyURL.User; t != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:85
		_go_fuzz_dep_.CoverTab[78570]++
														u := t.Username()
														p, _ := t.Password()
														req.Header.Add(proxyAuthHeaderKey, "Basic "+basicAuth(u, p))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:88
		// _ = "end of CoverTab[78570]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:89
		_go_fuzz_dep_.CoverTab[78571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:89
		// _ = "end of CoverTab[78571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:89
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:89
	// _ = "end of CoverTab[78562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:89
	_go_fuzz_dep_.CoverTab[78563]++

													if err := sendHTTPRequest(ctx, req, conn); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:91
		_go_fuzz_dep_.CoverTab[78572]++
														return nil, fmt.Errorf("failed to write the HTTP request: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:92
		// _ = "end of CoverTab[78572]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:93
		_go_fuzz_dep_.CoverTab[78573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:93
		// _ = "end of CoverTab[78573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:93
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:93
	// _ = "end of CoverTab[78563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:93
	_go_fuzz_dep_.CoverTab[78564]++

													r := bufio.NewReader(conn)
													resp, err := http.ReadResponse(r, req)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:97
		_go_fuzz_dep_.CoverTab[78574]++
														return nil, fmt.Errorf("reading server HTTP response: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:98
		// _ = "end of CoverTab[78574]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:99
		_go_fuzz_dep_.CoverTab[78575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:99
		// _ = "end of CoverTab[78575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:99
	// _ = "end of CoverTab[78564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:99
	_go_fuzz_dep_.CoverTab[78565]++
													defer resp.Body.Close()
													if resp.StatusCode != http.StatusOK {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:101
		_go_fuzz_dep_.CoverTab[78576]++
														dump, err := httputil.DumpResponse(resp, true)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:103
			_go_fuzz_dep_.CoverTab[78578]++
															return nil, fmt.Errorf("failed to do connect handshake, status code: %s", resp.Status)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:104
			// _ = "end of CoverTab[78578]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:105
			_go_fuzz_dep_.CoverTab[78579]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:105
			// _ = "end of CoverTab[78579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:105
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:105
		// _ = "end of CoverTab[78576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:105
		_go_fuzz_dep_.CoverTab[78577]++
														return nil, fmt.Errorf("failed to do connect handshake, response: %q", dump)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:106
		// _ = "end of CoverTab[78577]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:107
		_go_fuzz_dep_.CoverTab[78580]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:107
		// _ = "end of CoverTab[78580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:107
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:107
	// _ = "end of CoverTab[78565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:107
	_go_fuzz_dep_.CoverTab[78566]++

													return &bufConn{Conn: conn, r: r}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:109
	// _ = "end of CoverTab[78566]"
}

// proxyDial dials, connecting to a proxy first if necessary. Checks if a proxy
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:112
// is necessary, dials, does the HTTP CONNECT handshake, and returns the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:112
// connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:115
func proxyDial(ctx context.Context, addr string, grpcUA string) (conn net.Conn, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:115
	_go_fuzz_dep_.CoverTab[78581]++
													newAddr := addr
													proxyURL, err := mapAddress(addr)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:118
		_go_fuzz_dep_.CoverTab[78586]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:119
		// _ = "end of CoverTab[78586]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:120
		_go_fuzz_dep_.CoverTab[78587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:120
		// _ = "end of CoverTab[78587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:120
	// _ = "end of CoverTab[78581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:120
	_go_fuzz_dep_.CoverTab[78582]++
													if proxyURL != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:121
		_go_fuzz_dep_.CoverTab[78588]++
														newAddr = proxyURL.Host
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:122
		// _ = "end of CoverTab[78588]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:123
		_go_fuzz_dep_.CoverTab[78589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:123
		// _ = "end of CoverTab[78589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:123
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:123
	// _ = "end of CoverTab[78582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:123
	_go_fuzz_dep_.CoverTab[78583]++

													conn, err = (&net.Dialer{}).DialContext(ctx, "tcp", newAddr)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:126
		_go_fuzz_dep_.CoverTab[78590]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:127
		// _ = "end of CoverTab[78590]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:128
		_go_fuzz_dep_.CoverTab[78591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:128
		// _ = "end of CoverTab[78591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:128
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:128
	// _ = "end of CoverTab[78583]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:128
	_go_fuzz_dep_.CoverTab[78584]++
													if proxyURL != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:129
		_go_fuzz_dep_.CoverTab[78592]++

														conn, err = doHTTPConnectHandshake(ctx, conn, addr, proxyURL, grpcUA)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:131
		// _ = "end of CoverTab[78592]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:132
		_go_fuzz_dep_.CoverTab[78593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:132
		// _ = "end of CoverTab[78593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:132
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:132
	// _ = "end of CoverTab[78584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:132
	_go_fuzz_dep_.CoverTab[78585]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:133
	// _ = "end of CoverTab[78585]"
}

func sendHTTPRequest(ctx context.Context, req *http.Request, conn net.Conn) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:136
	_go_fuzz_dep_.CoverTab[78594]++
													req = req.WithContext(ctx)
													if err := req.Write(conn); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:138
		_go_fuzz_dep_.CoverTab[78596]++
														return fmt.Errorf("failed to write the HTTP request: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:139
		// _ = "end of CoverTab[78596]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:140
		_go_fuzz_dep_.CoverTab[78597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:140
		// _ = "end of CoverTab[78597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:140
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:140
	// _ = "end of CoverTab[78594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:140
	_go_fuzz_dep_.CoverTab[78595]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:141
	// _ = "end of CoverTab[78595]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:142
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/proxy.go:142
var _ = _go_fuzz_dep_.CoverTab
