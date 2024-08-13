// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/multipart/writer.go:5
package multipart

//line /usr/local/go/src/mime/multipart/writer.go:5
import (
//line /usr/local/go/src/mime/multipart/writer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/multipart/writer.go:5
)
//line /usr/local/go/src/mime/multipart/writer.go:5
import (
//line /usr/local/go/src/mime/multipart/writer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/multipart/writer.go:5
)

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"sort"
	"strings"
)

// A Writer generates multipart messages.
type Writer struct {
	w		io.Writer
	boundary	string
	lastpart	*part
}

// NewWriter returns a new multipart Writer with a random boundary,
//line /usr/local/go/src/mime/multipart/writer.go:25
// writing to w.
//line /usr/local/go/src/mime/multipart/writer.go:27
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/mime/multipart/writer.go:27
	_go_fuzz_dep_.CoverTab[36345]++
							return &Writer{
		w:		w,
		boundary:	randomBoundary(),
	}
//line /usr/local/go/src/mime/multipart/writer.go:31
	// _ = "end of CoverTab[36345]"
}

// Boundary returns the Writer's boundary.
func (w *Writer) Boundary() string {
//line /usr/local/go/src/mime/multipart/writer.go:35
	_go_fuzz_dep_.CoverTab[36346]++
							return w.boundary
//line /usr/local/go/src/mime/multipart/writer.go:36
	// _ = "end of CoverTab[36346]"
}

// SetBoundary overrides the Writer's default randomly-generated
//line /usr/local/go/src/mime/multipart/writer.go:39
// boundary separator with an explicit value.
//line /usr/local/go/src/mime/multipart/writer.go:39
//
//line /usr/local/go/src/mime/multipart/writer.go:39
// SetBoundary must be called before any parts are created, may only
//line /usr/local/go/src/mime/multipart/writer.go:39
// contain certain ASCII characters, and must be non-empty and
//line /usr/local/go/src/mime/multipart/writer.go:39
// at most 70 bytes long.
//line /usr/local/go/src/mime/multipart/writer.go:45
func (w *Writer) SetBoundary(boundary string) error {
//line /usr/local/go/src/mime/multipart/writer.go:45
	_go_fuzz_dep_.CoverTab[36347]++
							if w.lastpart != nil {
//line /usr/local/go/src/mime/multipart/writer.go:46
		_go_fuzz_dep_.CoverTab[36351]++
								return errors.New("mime: SetBoundary called after write")
//line /usr/local/go/src/mime/multipart/writer.go:47
		// _ = "end of CoverTab[36351]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:48
		_go_fuzz_dep_.CoverTab[36352]++
//line /usr/local/go/src/mime/multipart/writer.go:48
		// _ = "end of CoverTab[36352]"
//line /usr/local/go/src/mime/multipart/writer.go:48
	}
//line /usr/local/go/src/mime/multipart/writer.go:48
	// _ = "end of CoverTab[36347]"
//line /usr/local/go/src/mime/multipart/writer.go:48
	_go_fuzz_dep_.CoverTab[36348]++

							if len(boundary) < 1 || func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:50
		_go_fuzz_dep_.CoverTab[36353]++
//line /usr/local/go/src/mime/multipart/writer.go:50
		return len(boundary) > 70
//line /usr/local/go/src/mime/multipart/writer.go:50
		// _ = "end of CoverTab[36353]"
//line /usr/local/go/src/mime/multipart/writer.go:50
	}() {
//line /usr/local/go/src/mime/multipart/writer.go:50
		_go_fuzz_dep_.CoverTab[36354]++
								return errors.New("mime: invalid boundary length")
//line /usr/local/go/src/mime/multipart/writer.go:51
		// _ = "end of CoverTab[36354]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:52
		_go_fuzz_dep_.CoverTab[36355]++
//line /usr/local/go/src/mime/multipart/writer.go:52
		// _ = "end of CoverTab[36355]"
//line /usr/local/go/src/mime/multipart/writer.go:52
	}
//line /usr/local/go/src/mime/multipart/writer.go:52
	// _ = "end of CoverTab[36348]"
//line /usr/local/go/src/mime/multipart/writer.go:52
	_go_fuzz_dep_.CoverTab[36349]++
							end := len(boundary) - 1
							for i, b := range boundary {
//line /usr/local/go/src/mime/multipart/writer.go:54
		_go_fuzz_dep_.CoverTab[36356]++
								if 'A' <= b && func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:55
			_go_fuzz_dep_.CoverTab[36359]++
//line /usr/local/go/src/mime/multipart/writer.go:55
			return b <= 'Z'
//line /usr/local/go/src/mime/multipart/writer.go:55
			// _ = "end of CoverTab[36359]"
//line /usr/local/go/src/mime/multipart/writer.go:55
		}() || func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:55
			_go_fuzz_dep_.CoverTab[36360]++
//line /usr/local/go/src/mime/multipart/writer.go:55
			return 'a' <= b && func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:55
				_go_fuzz_dep_.CoverTab[36361]++
//line /usr/local/go/src/mime/multipart/writer.go:55
				return b <= 'z'
//line /usr/local/go/src/mime/multipart/writer.go:55
				// _ = "end of CoverTab[36361]"
//line /usr/local/go/src/mime/multipart/writer.go:55
			}()
//line /usr/local/go/src/mime/multipart/writer.go:55
			// _ = "end of CoverTab[36360]"
//line /usr/local/go/src/mime/multipart/writer.go:55
		}() || func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:55
			_go_fuzz_dep_.CoverTab[36362]++
//line /usr/local/go/src/mime/multipart/writer.go:55
			return '0' <= b && func() bool {
//line /usr/local/go/src/mime/multipart/writer.go:55
				_go_fuzz_dep_.CoverTab[36363]++
//line /usr/local/go/src/mime/multipart/writer.go:55
				return b <= '9'
//line /usr/local/go/src/mime/multipart/writer.go:55
				// _ = "end of CoverTab[36363]"
//line /usr/local/go/src/mime/multipart/writer.go:55
			}()
//line /usr/local/go/src/mime/multipart/writer.go:55
			// _ = "end of CoverTab[36362]"
//line /usr/local/go/src/mime/multipart/writer.go:55
		}() {
//line /usr/local/go/src/mime/multipart/writer.go:55
			_go_fuzz_dep_.CoverTab[36364]++
									continue
//line /usr/local/go/src/mime/multipart/writer.go:56
			// _ = "end of CoverTab[36364]"
		} else {
//line /usr/local/go/src/mime/multipart/writer.go:57
			_go_fuzz_dep_.CoverTab[36365]++
//line /usr/local/go/src/mime/multipart/writer.go:57
			// _ = "end of CoverTab[36365]"
//line /usr/local/go/src/mime/multipart/writer.go:57
		}
//line /usr/local/go/src/mime/multipart/writer.go:57
		// _ = "end of CoverTab[36356]"
//line /usr/local/go/src/mime/multipart/writer.go:57
		_go_fuzz_dep_.CoverTab[36357]++
								switch b {
		case '\'', '(', ')', '+', '_', ',', '-', '.', '/', ':', '=', '?':
//line /usr/local/go/src/mime/multipart/writer.go:59
			_go_fuzz_dep_.CoverTab[36366]++
									continue
//line /usr/local/go/src/mime/multipart/writer.go:60
			// _ = "end of CoverTab[36366]"
		case ' ':
//line /usr/local/go/src/mime/multipart/writer.go:61
			_go_fuzz_dep_.CoverTab[36367]++
									if i != end {
//line /usr/local/go/src/mime/multipart/writer.go:62
				_go_fuzz_dep_.CoverTab[36369]++
										continue
//line /usr/local/go/src/mime/multipart/writer.go:63
				// _ = "end of CoverTab[36369]"
			} else {
//line /usr/local/go/src/mime/multipart/writer.go:64
				_go_fuzz_dep_.CoverTab[36370]++
//line /usr/local/go/src/mime/multipart/writer.go:64
				// _ = "end of CoverTab[36370]"
//line /usr/local/go/src/mime/multipart/writer.go:64
			}
//line /usr/local/go/src/mime/multipart/writer.go:64
			// _ = "end of CoverTab[36367]"
//line /usr/local/go/src/mime/multipart/writer.go:64
		default:
//line /usr/local/go/src/mime/multipart/writer.go:64
			_go_fuzz_dep_.CoverTab[36368]++
//line /usr/local/go/src/mime/multipart/writer.go:64
			// _ = "end of CoverTab[36368]"
		}
//line /usr/local/go/src/mime/multipart/writer.go:65
		// _ = "end of CoverTab[36357]"
//line /usr/local/go/src/mime/multipart/writer.go:65
		_go_fuzz_dep_.CoverTab[36358]++
								return errors.New("mime: invalid boundary character")
//line /usr/local/go/src/mime/multipart/writer.go:66
		// _ = "end of CoverTab[36358]"
	}
//line /usr/local/go/src/mime/multipart/writer.go:67
	// _ = "end of CoverTab[36349]"
//line /usr/local/go/src/mime/multipart/writer.go:67
	_go_fuzz_dep_.CoverTab[36350]++
							w.boundary = boundary
							return nil
//line /usr/local/go/src/mime/multipart/writer.go:69
	// _ = "end of CoverTab[36350]"
}

// FormDataContentType returns the Content-Type for an HTTP
//line /usr/local/go/src/mime/multipart/writer.go:72
// multipart/form-data with this Writer's Boundary.
//line /usr/local/go/src/mime/multipart/writer.go:74
func (w *Writer) FormDataContentType() string {
//line /usr/local/go/src/mime/multipart/writer.go:74
	_go_fuzz_dep_.CoverTab[36371]++
							b := w.boundary

//line /usr/local/go/src/mime/multipart/writer.go:78
	if strings.ContainsAny(b, `()<>@,;:\"/[]?= `) {
//line /usr/local/go/src/mime/multipart/writer.go:78
		_go_fuzz_dep_.CoverTab[36373]++
								b = `"` + b + `"`
//line /usr/local/go/src/mime/multipart/writer.go:79
		// _ = "end of CoverTab[36373]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:80
		_go_fuzz_dep_.CoverTab[36374]++
//line /usr/local/go/src/mime/multipart/writer.go:80
		// _ = "end of CoverTab[36374]"
//line /usr/local/go/src/mime/multipart/writer.go:80
	}
//line /usr/local/go/src/mime/multipart/writer.go:80
	// _ = "end of CoverTab[36371]"
//line /usr/local/go/src/mime/multipart/writer.go:80
	_go_fuzz_dep_.CoverTab[36372]++
							return "multipart/form-data; boundary=" + b
//line /usr/local/go/src/mime/multipart/writer.go:81
	// _ = "end of CoverTab[36372]"
}

func randomBoundary() string {
//line /usr/local/go/src/mime/multipart/writer.go:84
	_go_fuzz_dep_.CoverTab[36375]++
							var buf [30]byte
							_, err := io.ReadFull(rand.Reader, buf[:])
							if err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:87
		_go_fuzz_dep_.CoverTab[36377]++
								panic(err)
//line /usr/local/go/src/mime/multipart/writer.go:88
		// _ = "end of CoverTab[36377]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:89
		_go_fuzz_dep_.CoverTab[36378]++
//line /usr/local/go/src/mime/multipart/writer.go:89
		// _ = "end of CoverTab[36378]"
//line /usr/local/go/src/mime/multipart/writer.go:89
	}
//line /usr/local/go/src/mime/multipart/writer.go:89
	// _ = "end of CoverTab[36375]"
//line /usr/local/go/src/mime/multipart/writer.go:89
	_go_fuzz_dep_.CoverTab[36376]++
							return fmt.Sprintf("%x", buf[:])
//line /usr/local/go/src/mime/multipart/writer.go:90
	// _ = "end of CoverTab[36376]"
}

// CreatePart creates a new multipart section with the provided
//line /usr/local/go/src/mime/multipart/writer.go:93
// header. The body of the part should be written to the returned
//line /usr/local/go/src/mime/multipart/writer.go:93
// Writer. After calling CreatePart, any previous part may no longer
//line /usr/local/go/src/mime/multipart/writer.go:93
// be written to.
//line /usr/local/go/src/mime/multipart/writer.go:97
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error) {
//line /usr/local/go/src/mime/multipart/writer.go:97
	_go_fuzz_dep_.CoverTab[36379]++
							if w.lastpart != nil {
//line /usr/local/go/src/mime/multipart/writer.go:98
		_go_fuzz_dep_.CoverTab[36385]++
								if err := w.lastpart.close(); err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:99
			_go_fuzz_dep_.CoverTab[36386]++
									return nil, err
//line /usr/local/go/src/mime/multipart/writer.go:100
			// _ = "end of CoverTab[36386]"
		} else {
//line /usr/local/go/src/mime/multipart/writer.go:101
			_go_fuzz_dep_.CoverTab[36387]++
//line /usr/local/go/src/mime/multipart/writer.go:101
			// _ = "end of CoverTab[36387]"
//line /usr/local/go/src/mime/multipart/writer.go:101
		}
//line /usr/local/go/src/mime/multipart/writer.go:101
		// _ = "end of CoverTab[36385]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:102
		_go_fuzz_dep_.CoverTab[36388]++
//line /usr/local/go/src/mime/multipart/writer.go:102
		// _ = "end of CoverTab[36388]"
//line /usr/local/go/src/mime/multipart/writer.go:102
	}
//line /usr/local/go/src/mime/multipart/writer.go:102
	// _ = "end of CoverTab[36379]"
//line /usr/local/go/src/mime/multipart/writer.go:102
	_go_fuzz_dep_.CoverTab[36380]++
							var b bytes.Buffer
							if w.lastpart != nil {
//line /usr/local/go/src/mime/multipart/writer.go:104
		_go_fuzz_dep_.CoverTab[36389]++
								fmt.Fprintf(&b, "\r\n--%s\r\n", w.boundary)
//line /usr/local/go/src/mime/multipart/writer.go:105
		// _ = "end of CoverTab[36389]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:106
		_go_fuzz_dep_.CoverTab[36390]++
								fmt.Fprintf(&b, "--%s\r\n", w.boundary)
//line /usr/local/go/src/mime/multipart/writer.go:107
		// _ = "end of CoverTab[36390]"
	}
//line /usr/local/go/src/mime/multipart/writer.go:108
	// _ = "end of CoverTab[36380]"
//line /usr/local/go/src/mime/multipart/writer.go:108
	_go_fuzz_dep_.CoverTab[36381]++

							keys := make([]string, 0, len(header))
							for k := range header {
//line /usr/local/go/src/mime/multipart/writer.go:111
		_go_fuzz_dep_.CoverTab[36391]++
								keys = append(keys, k)
//line /usr/local/go/src/mime/multipart/writer.go:112
		// _ = "end of CoverTab[36391]"
	}
//line /usr/local/go/src/mime/multipart/writer.go:113
	// _ = "end of CoverTab[36381]"
//line /usr/local/go/src/mime/multipart/writer.go:113
	_go_fuzz_dep_.CoverTab[36382]++
							sort.Strings(keys)
							for _, k := range keys {
//line /usr/local/go/src/mime/multipart/writer.go:115
		_go_fuzz_dep_.CoverTab[36392]++
								for _, v := range header[k] {
//line /usr/local/go/src/mime/multipart/writer.go:116
			_go_fuzz_dep_.CoverTab[36393]++
									fmt.Fprintf(&b, "%s: %s\r\n", k, v)
//line /usr/local/go/src/mime/multipart/writer.go:117
			// _ = "end of CoverTab[36393]"
		}
//line /usr/local/go/src/mime/multipart/writer.go:118
		// _ = "end of CoverTab[36392]"
	}
//line /usr/local/go/src/mime/multipart/writer.go:119
	// _ = "end of CoverTab[36382]"
//line /usr/local/go/src/mime/multipart/writer.go:119
	_go_fuzz_dep_.CoverTab[36383]++
							fmt.Fprintf(&b, "\r\n")
							_, err := io.Copy(w.w, &b)
							if err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:122
		_go_fuzz_dep_.CoverTab[36394]++
								return nil, err
//line /usr/local/go/src/mime/multipart/writer.go:123
		// _ = "end of CoverTab[36394]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:124
		_go_fuzz_dep_.CoverTab[36395]++
//line /usr/local/go/src/mime/multipart/writer.go:124
		// _ = "end of CoverTab[36395]"
//line /usr/local/go/src/mime/multipart/writer.go:124
	}
//line /usr/local/go/src/mime/multipart/writer.go:124
	// _ = "end of CoverTab[36383]"
//line /usr/local/go/src/mime/multipart/writer.go:124
	_go_fuzz_dep_.CoverTab[36384]++
							p := &part{
		mw: w,
	}
							w.lastpart = p
							return p, nil
//line /usr/local/go/src/mime/multipart/writer.go:129
	// _ = "end of CoverTab[36384]"
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
//line /usr/local/go/src/mime/multipart/writer.go:134
	_go_fuzz_dep_.CoverTab[36396]++
							return quoteEscaper.Replace(s)
//line /usr/local/go/src/mime/multipart/writer.go:135
	// _ = "end of CoverTab[36396]"
}

// CreateFormFile is a convenience wrapper around CreatePart. It creates
//line /usr/local/go/src/mime/multipart/writer.go:138
// a new form-data header with the provided field name and file name.
//line /usr/local/go/src/mime/multipart/writer.go:140
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error) {
//line /usr/local/go/src/mime/multipart/writer.go:140
	_go_fuzz_dep_.CoverTab[36397]++
							h := make(textproto.MIMEHeader)
							h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
							h.Set("Content-Type", "application/octet-stream")
							return w.CreatePart(h)
//line /usr/local/go/src/mime/multipart/writer.go:146
	// _ = "end of CoverTab[36397]"
}

// CreateFormField calls CreatePart with a header using the
//line /usr/local/go/src/mime/multipart/writer.go:149
// given field name.
//line /usr/local/go/src/mime/multipart/writer.go:151
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error) {
//line /usr/local/go/src/mime/multipart/writer.go:151
	_go_fuzz_dep_.CoverTab[36398]++
							h := make(textproto.MIMEHeader)
							h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldname)))
							return w.CreatePart(h)
//line /usr/local/go/src/mime/multipart/writer.go:155
	// _ = "end of CoverTab[36398]"
}

// WriteField calls CreateFormField and then writes the given value.
func (w *Writer) WriteField(fieldname, value string) error {
//line /usr/local/go/src/mime/multipart/writer.go:159
	_go_fuzz_dep_.CoverTab[36399]++
							p, err := w.CreateFormField(fieldname)
							if err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:161
		_go_fuzz_dep_.CoverTab[36401]++
								return err
//line /usr/local/go/src/mime/multipart/writer.go:162
		// _ = "end of CoverTab[36401]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:163
		_go_fuzz_dep_.CoverTab[36402]++
//line /usr/local/go/src/mime/multipart/writer.go:163
		// _ = "end of CoverTab[36402]"
//line /usr/local/go/src/mime/multipart/writer.go:163
	}
//line /usr/local/go/src/mime/multipart/writer.go:163
	// _ = "end of CoverTab[36399]"
//line /usr/local/go/src/mime/multipart/writer.go:163
	_go_fuzz_dep_.CoverTab[36400]++
							_, err = p.Write([]byte(value))
							return err
//line /usr/local/go/src/mime/multipart/writer.go:165
	// _ = "end of CoverTab[36400]"
}

// Close finishes the multipart message and writes the trailing
//line /usr/local/go/src/mime/multipart/writer.go:168
// boundary end line to the output.
//line /usr/local/go/src/mime/multipart/writer.go:170
func (w *Writer) Close() error {
//line /usr/local/go/src/mime/multipart/writer.go:170
	_go_fuzz_dep_.CoverTab[36403]++
							if w.lastpart != nil {
//line /usr/local/go/src/mime/multipart/writer.go:171
		_go_fuzz_dep_.CoverTab[36405]++
								if err := w.lastpart.close(); err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:172
			_go_fuzz_dep_.CoverTab[36407]++
									return err
//line /usr/local/go/src/mime/multipart/writer.go:173
			// _ = "end of CoverTab[36407]"
		} else {
//line /usr/local/go/src/mime/multipart/writer.go:174
			_go_fuzz_dep_.CoverTab[36408]++
//line /usr/local/go/src/mime/multipart/writer.go:174
			// _ = "end of CoverTab[36408]"
//line /usr/local/go/src/mime/multipart/writer.go:174
		}
//line /usr/local/go/src/mime/multipart/writer.go:174
		// _ = "end of CoverTab[36405]"
//line /usr/local/go/src/mime/multipart/writer.go:174
		_go_fuzz_dep_.CoverTab[36406]++
								w.lastpart = nil
//line /usr/local/go/src/mime/multipart/writer.go:175
		// _ = "end of CoverTab[36406]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:176
		_go_fuzz_dep_.CoverTab[36409]++
//line /usr/local/go/src/mime/multipart/writer.go:176
		// _ = "end of CoverTab[36409]"
//line /usr/local/go/src/mime/multipart/writer.go:176
	}
//line /usr/local/go/src/mime/multipart/writer.go:176
	// _ = "end of CoverTab[36403]"
//line /usr/local/go/src/mime/multipart/writer.go:176
	_go_fuzz_dep_.CoverTab[36404]++
							_, err := fmt.Fprintf(w.w, "\r\n--%s--\r\n", w.boundary)
							return err
//line /usr/local/go/src/mime/multipart/writer.go:178
	// _ = "end of CoverTab[36404]"
}

type part struct {
	mw	*Writer
	closed	bool
	we	error	// last error that occurred writing
}

func (p *part) close() error {
//line /usr/local/go/src/mime/multipart/writer.go:187
	_go_fuzz_dep_.CoverTab[36410]++
							p.closed = true
							return p.we
//line /usr/local/go/src/mime/multipart/writer.go:189
	// _ = "end of CoverTab[36410]"
}

func (p *part) Write(d []byte) (n int, err error) {
//line /usr/local/go/src/mime/multipart/writer.go:192
	_go_fuzz_dep_.CoverTab[36411]++
							if p.closed {
//line /usr/local/go/src/mime/multipart/writer.go:193
		_go_fuzz_dep_.CoverTab[36414]++
								return 0, errors.New("multipart: can't write to finished part")
//line /usr/local/go/src/mime/multipart/writer.go:194
		// _ = "end of CoverTab[36414]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:195
		_go_fuzz_dep_.CoverTab[36415]++
//line /usr/local/go/src/mime/multipart/writer.go:195
		// _ = "end of CoverTab[36415]"
//line /usr/local/go/src/mime/multipart/writer.go:195
	}
//line /usr/local/go/src/mime/multipart/writer.go:195
	// _ = "end of CoverTab[36411]"
//line /usr/local/go/src/mime/multipart/writer.go:195
	_go_fuzz_dep_.CoverTab[36412]++
							n, err = p.mw.w.Write(d)
							if err != nil {
//line /usr/local/go/src/mime/multipart/writer.go:197
		_go_fuzz_dep_.CoverTab[36416]++
								p.we = err
//line /usr/local/go/src/mime/multipart/writer.go:198
		// _ = "end of CoverTab[36416]"
	} else {
//line /usr/local/go/src/mime/multipart/writer.go:199
		_go_fuzz_dep_.CoverTab[36417]++
//line /usr/local/go/src/mime/multipart/writer.go:199
		// _ = "end of CoverTab[36417]"
//line /usr/local/go/src/mime/multipart/writer.go:199
	}
//line /usr/local/go/src/mime/multipart/writer.go:199
	// _ = "end of CoverTab[36412]"
//line /usr/local/go/src/mime/multipart/writer.go:199
	_go_fuzz_dep_.CoverTab[36413]++
							return
//line /usr/local/go/src/mime/multipart/writer.go:200
	// _ = "end of CoverTab[36413]"
}

//line /usr/local/go/src/mime/multipart/writer.go:201
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/multipart/writer.go:201
var _ = _go_fuzz_dep_.CoverTab
