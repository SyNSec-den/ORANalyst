//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
package scanner

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:1
)

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/hcl/json/token"
)

// eof represents a marker rune for the end of the reader.
const eof = rune(0)

// Scanner defines a lexical scanner
type Scanner struct {
	buf	*bytes.Buffer	// Source buffer for advancing and scanning
	src	[]byte		// Source buffer for immutable access

	// Source Position
	srcPos	token.Pos	// current position
	prevPos	token.Pos	// previous position, used for peek() method

	lastCharLen	int	// length of last character in bytes
	lastLineLen	int	// length of last line in characters (for correct column reporting)

	tokStart	int	// token text start position
	tokEnd		int	// token text end  position

	// Error is called for each error encountered. If no Error
	// function is set, the error is reported to os.Stderr.
	Error	func(pos token.Pos, msg string)

	// ErrorCount is incremented by one for each error encountered.
	ErrorCount	int

	// tokPos is the start position of most recently scanned token; set by
	// Scan. The Filename field is always left untouched by the Scanner.  If
	// an error is reported (via Error) and Position is invalid, the scanner is
	// not inside a token.
	tokPos	token.Pos
}

// New creates and initializes a new instance of Scanner using src as
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:45
// its source content.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:47
func New(src []byte) *Scanner {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:47
	_go_fuzz_dep_.CoverTab[121584]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:51
	b := bytes.NewBuffer(src)
	s := &Scanner{
		buf:	b,
		src:	src,
	}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:58
	s.srcPos.Line = 1
												return s
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:59
	// _ = "end of CoverTab[121584]"
}

// next reads the next rune from the bufferred reader. Returns the rune(0) if
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:62
// an error occurs (or io.EOF is returned).
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:64
func (s *Scanner) next() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:64
	_go_fuzz_dep_.CoverTab[121585]++
												ch, size, err := s.buf.ReadRune()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:66
		_go_fuzz_dep_.CoverTab[121589]++

													s.srcPos.Column++
													s.srcPos.Offset += size
													s.lastCharLen = size
													return eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:71
		// _ = "end of CoverTab[121589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:72
		_go_fuzz_dep_.CoverTab[121590]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:72
		// _ = "end of CoverTab[121590]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:72
	// _ = "end of CoverTab[121585]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:72
	_go_fuzz_dep_.CoverTab[121586]++

												if ch == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:74
		_go_fuzz_dep_.CoverTab[121591]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:74
		return size == 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:74
		// _ = "end of CoverTab[121591]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:74
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:74
		_go_fuzz_dep_.CoverTab[121592]++
													s.srcPos.Column++
													s.srcPos.Offset += size
													s.lastCharLen = size
													s.err("illegal UTF-8 encoding")
													return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:79
		// _ = "end of CoverTab[121592]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:80
		_go_fuzz_dep_.CoverTab[121593]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:80
		// _ = "end of CoverTab[121593]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:80
	// _ = "end of CoverTab[121586]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:80
	_go_fuzz_dep_.CoverTab[121587]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:83
	s.prevPos = s.srcPos

	s.srcPos.Column++
	s.lastCharLen = size
	s.srcPos.Offset += size

	if ch == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:89
		_go_fuzz_dep_.CoverTab[121594]++
													s.srcPos.Line++
													s.lastLineLen = s.srcPos.Column
													s.srcPos.Column = 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:92
		// _ = "end of CoverTab[121594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:93
		_go_fuzz_dep_.CoverTab[121595]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:93
		// _ = "end of CoverTab[121595]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:93
	// _ = "end of CoverTab[121587]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:93
	_go_fuzz_dep_.CoverTab[121588]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:97
	return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:97
	// _ = "end of CoverTab[121588]"
}

// unread unreads the previous read Rune and updates the source position
func (s *Scanner) unread() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:101
	_go_fuzz_dep_.CoverTab[121596]++
												if err := s.buf.UnreadRune(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:102
		_go_fuzz_dep_.CoverTab[121598]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:103
		// _ = "end of CoverTab[121598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:104
		_go_fuzz_dep_.CoverTab[121599]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:104
		// _ = "end of CoverTab[121599]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:104
	// _ = "end of CoverTab[121596]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:104
	_go_fuzz_dep_.CoverTab[121597]++
												s.srcPos = s.prevPos
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:105
	// _ = "end of CoverTab[121597]"
}

// peek returns the next rune without advancing the reader.
func (s *Scanner) peek() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:109
	_go_fuzz_dep_.CoverTab[121600]++
												peek, _, err := s.buf.ReadRune()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:111
		_go_fuzz_dep_.CoverTab[121602]++
													return eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:112
		// _ = "end of CoverTab[121602]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:113
		_go_fuzz_dep_.CoverTab[121603]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:113
		// _ = "end of CoverTab[121603]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:113
	// _ = "end of CoverTab[121600]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:113
	_go_fuzz_dep_.CoverTab[121601]++

												s.buf.UnreadRune()
												return peek
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:116
	// _ = "end of CoverTab[121601]"
}

// Scan scans the next token and returns the token.
func (s *Scanner) Scan() token.Token {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:120
	_go_fuzz_dep_.CoverTab[121604]++
												ch := s.next()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:124
	for isWhitespace(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:124
		_go_fuzz_dep_.CoverTab[121609]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:125
		// _ = "end of CoverTab[121609]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:126
	// _ = "end of CoverTab[121604]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:126
	_go_fuzz_dep_.CoverTab[121605]++

												var tok token.Type

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:131
	s.tokStart = s.srcPos.Offset - s.lastCharLen

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:135
	s.tokPos.Offset = s.srcPos.Offset - s.lastCharLen
	if s.srcPos.Column > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:136
		_go_fuzz_dep_.CoverTab[121610]++

													s.tokPos.Line = s.srcPos.Line
													s.tokPos.Column = s.srcPos.Column
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:139
		// _ = "end of CoverTab[121610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:140
		_go_fuzz_dep_.CoverTab[121611]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:144
		s.tokPos.Line = s.srcPos.Line - 1
													s.tokPos.Column = s.lastLineLen
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:145
		// _ = "end of CoverTab[121611]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:146
	// _ = "end of CoverTab[121605]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:146
	_go_fuzz_dep_.CoverTab[121606]++

												switch {
	case isLetter(ch):
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:149
		_go_fuzz_dep_.CoverTab[121612]++
													lit := s.scanIdentifier()
													if lit == "true" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:151
			_go_fuzz_dep_.CoverTab[121615]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:151
			return lit == "false"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:151
			// _ = "end of CoverTab[121615]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:151
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:151
			_go_fuzz_dep_.CoverTab[121616]++
														tok = token.BOOL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:152
			// _ = "end of CoverTab[121616]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:153
			_go_fuzz_dep_.CoverTab[121617]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:153
			if lit == "null" {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:153
				_go_fuzz_dep_.CoverTab[121618]++
															tok = token.NULL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:154
				// _ = "end of CoverTab[121618]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:155
				_go_fuzz_dep_.CoverTab[121619]++
															s.err("illegal char")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:156
				// _ = "end of CoverTab[121619]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:157
			// _ = "end of CoverTab[121617]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:157
		// _ = "end of CoverTab[121612]"
	case isDecimal(ch):
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:158
		_go_fuzz_dep_.CoverTab[121613]++
													tok = s.scanNumber(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:159
		// _ = "end of CoverTab[121613]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:160
		_go_fuzz_dep_.CoverTab[121614]++
													switch ch {
		case eof:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:162
			_go_fuzz_dep_.CoverTab[121620]++
														tok = token.EOF
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:163
			// _ = "end of CoverTab[121620]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:164
			_go_fuzz_dep_.CoverTab[121621]++
														tok = token.STRING
														s.scanString()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:166
			// _ = "end of CoverTab[121621]"
		case '.':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:167
			_go_fuzz_dep_.CoverTab[121622]++
														tok = token.PERIOD
														ch = s.peek()
														if isDecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:170
				_go_fuzz_dep_.CoverTab[121631]++
															tok = token.FLOAT
															ch = s.scanMantissa(ch)
															ch = s.scanExponent(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:173
				// _ = "end of CoverTab[121631]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:174
				_go_fuzz_dep_.CoverTab[121632]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:174
				// _ = "end of CoverTab[121632]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:174
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:174
			// _ = "end of CoverTab[121622]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:175
			_go_fuzz_dep_.CoverTab[121623]++
														tok = token.LBRACK
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:176
			// _ = "end of CoverTab[121623]"
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:177
			_go_fuzz_dep_.CoverTab[121624]++
														tok = token.RBRACK
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:178
			// _ = "end of CoverTab[121624]"
		case '{':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:179
			_go_fuzz_dep_.CoverTab[121625]++
														tok = token.LBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:180
			// _ = "end of CoverTab[121625]"
		case '}':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:181
			_go_fuzz_dep_.CoverTab[121626]++
														tok = token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:182
			// _ = "end of CoverTab[121626]"
		case ',':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:183
			_go_fuzz_dep_.CoverTab[121627]++
														tok = token.COMMA
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:184
			// _ = "end of CoverTab[121627]"
		case ':':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:185
			_go_fuzz_dep_.CoverTab[121628]++
														tok = token.COLON
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:186
			// _ = "end of CoverTab[121628]"
		case '-':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:187
			_go_fuzz_dep_.CoverTab[121629]++
														if isDecimal(s.peek()) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:188
				_go_fuzz_dep_.CoverTab[121633]++
															ch := s.next()
															tok = s.scanNumber(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:190
				// _ = "end of CoverTab[121633]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:191
				_go_fuzz_dep_.CoverTab[121634]++
															s.err("illegal char")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:192
				// _ = "end of CoverTab[121634]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:193
			// _ = "end of CoverTab[121629]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:194
			_go_fuzz_dep_.CoverTab[121630]++
														s.err("illegal char: " + string(ch))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:195
			// _ = "end of CoverTab[121630]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:196
		// _ = "end of CoverTab[121614]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:197
	// _ = "end of CoverTab[121606]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:197
	_go_fuzz_dep_.CoverTab[121607]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:200
	s.tokEnd = s.srcPos.Offset

	// create token literal
	var tokenText string
	if s.tokStart >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:204
		_go_fuzz_dep_.CoverTab[121635]++
													tokenText = string(s.src[s.tokStart:s.tokEnd])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:205
		// _ = "end of CoverTab[121635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:206
		_go_fuzz_dep_.CoverTab[121636]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:206
		// _ = "end of CoverTab[121636]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:206
	// _ = "end of CoverTab[121607]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:206
	_go_fuzz_dep_.CoverTab[121608]++
												s.tokStart = s.tokEnd

												return token.Token{
		Type:	tok,
		Pos:	s.tokPos,
		Text:	tokenText,
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:213
	// _ = "end of CoverTab[121608]"
}

// scanNumber scans a HCL number definition starting with the given rune
func (s *Scanner) scanNumber(ch rune) token.Type {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:217
	_go_fuzz_dep_.CoverTab[121637]++
												zero := ch == '0'
												pos := s.srcPos

												s.scanMantissa(ch)
												ch = s.next()
												if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:223
		_go_fuzz_dep_.CoverTab[121642]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:223
		return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:223
		// _ = "end of CoverTab[121642]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:223
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:223
		_go_fuzz_dep_.CoverTab[121643]++
													ch = s.scanExponent(ch)
													return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:225
		// _ = "end of CoverTab[121643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:226
		_go_fuzz_dep_.CoverTab[121644]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:226
		// _ = "end of CoverTab[121644]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:226
	// _ = "end of CoverTab[121637]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:226
	_go_fuzz_dep_.CoverTab[121638]++

												if ch == '.' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:228
		_go_fuzz_dep_.CoverTab[121645]++
													ch = s.scanFraction(ch)
													if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:230
			_go_fuzz_dep_.CoverTab[121647]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:230
			return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:230
			// _ = "end of CoverTab[121647]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:230
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:230
			_go_fuzz_dep_.CoverTab[121648]++
														ch = s.next()
														ch = s.scanExponent(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:232
			// _ = "end of CoverTab[121648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:233
			_go_fuzz_dep_.CoverTab[121649]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:233
			// _ = "end of CoverTab[121649]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:233
		// _ = "end of CoverTab[121645]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:233
		_go_fuzz_dep_.CoverTab[121646]++
													return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:234
		// _ = "end of CoverTab[121646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:235
		_go_fuzz_dep_.CoverTab[121650]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:235
		// _ = "end of CoverTab[121650]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:235
	// _ = "end of CoverTab[121638]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:235
	_go_fuzz_dep_.CoverTab[121639]++

												if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:237
		_go_fuzz_dep_.CoverTab[121651]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:238
		// _ = "end of CoverTab[121651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:239
		_go_fuzz_dep_.CoverTab[121652]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:239
		// _ = "end of CoverTab[121652]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:239
	// _ = "end of CoverTab[121639]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:239
	_go_fuzz_dep_.CoverTab[121640]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
	if zero && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
		_go_fuzz_dep_.CoverTab[121653]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
		return pos != s.srcPos
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
		// _ = "end of CoverTab[121653]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:242
		_go_fuzz_dep_.CoverTab[121654]++
													s.err("numbers cannot start with 0")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:243
		// _ = "end of CoverTab[121654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:244
		_go_fuzz_dep_.CoverTab[121655]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:244
		// _ = "end of CoverTab[121655]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:244
	// _ = "end of CoverTab[121640]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:244
	_go_fuzz_dep_.CoverTab[121641]++

												return token.NUMBER
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:246
	// _ = "end of CoverTab[121641]"
}

// scanMantissa scans the mantissa beginning from the rune. It returns the next
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:249
// non decimal rune. It's used to determine wheter it's a fraction or exponent.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:251
func (s *Scanner) scanMantissa(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:251
	_go_fuzz_dep_.CoverTab[121656]++
												scanned := false
												for isDecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:253
		_go_fuzz_dep_.CoverTab[121659]++
													ch = s.next()
													scanned = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:255
		// _ = "end of CoverTab[121659]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:256
	// _ = "end of CoverTab[121656]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:256
	_go_fuzz_dep_.CoverTab[121657]++

												if scanned && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:258
		_go_fuzz_dep_.CoverTab[121660]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:258
		return ch != eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:258
		// _ = "end of CoverTab[121660]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:258
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:258
		_go_fuzz_dep_.CoverTab[121661]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:259
		// _ = "end of CoverTab[121661]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:260
		_go_fuzz_dep_.CoverTab[121662]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:260
		// _ = "end of CoverTab[121662]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:260
	// _ = "end of CoverTab[121657]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:260
	_go_fuzz_dep_.CoverTab[121658]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:261
	// _ = "end of CoverTab[121658]"
}

// scanFraction scans the fraction after the '.' rune
func (s *Scanner) scanFraction(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:265
	_go_fuzz_dep_.CoverTab[121663]++
												if ch == '.' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:266
		_go_fuzz_dep_.CoverTab[121665]++
													ch = s.peek()
													ch = s.scanMantissa(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:268
		// _ = "end of CoverTab[121665]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:269
		_go_fuzz_dep_.CoverTab[121666]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:269
		// _ = "end of CoverTab[121666]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:269
	// _ = "end of CoverTab[121663]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:269
	_go_fuzz_dep_.CoverTab[121664]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:270
	// _ = "end of CoverTab[121664]"
}

// scanExponent scans the remaining parts of an exponent after the 'e' or 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:273
// rune.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:275
func (s *Scanner) scanExponent(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:275
	_go_fuzz_dep_.CoverTab[121667]++
												if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:276
		_go_fuzz_dep_.CoverTab[121669]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:276
		return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:276
		// _ = "end of CoverTab[121669]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:276
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:276
		_go_fuzz_dep_.CoverTab[121670]++
													ch = s.next()
													if ch == '-' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:278
			_go_fuzz_dep_.CoverTab[121672]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:278
			return ch == '+'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:278
			// _ = "end of CoverTab[121672]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:278
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:278
			_go_fuzz_dep_.CoverTab[121673]++
														ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:279
			// _ = "end of CoverTab[121673]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:280
			_go_fuzz_dep_.CoverTab[121674]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:280
			// _ = "end of CoverTab[121674]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:280
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:280
		// _ = "end of CoverTab[121670]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:280
		_go_fuzz_dep_.CoverTab[121671]++
													ch = s.scanMantissa(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:281
		// _ = "end of CoverTab[121671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:282
		_go_fuzz_dep_.CoverTab[121675]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:282
		// _ = "end of CoverTab[121675]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:282
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:282
	// _ = "end of CoverTab[121667]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:282
	_go_fuzz_dep_.CoverTab[121668]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:283
	// _ = "end of CoverTab[121668]"
}

// scanString scans a quoted string
func (s *Scanner) scanString() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:287
	_go_fuzz_dep_.CoverTab[121676]++
												braces := 0
												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:289
		_go_fuzz_dep_.CoverTab[121678]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:292
		ch := s.next()

		if ch == '\n' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			_go_fuzz_dep_.CoverTab[121683]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			return ch < 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			// _ = "end of CoverTab[121683]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			_go_fuzz_dep_.CoverTab[121684]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			return ch == eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			// _ = "end of CoverTab[121684]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:294
			_go_fuzz_dep_.CoverTab[121685]++
														s.err("literal not terminated")
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:296
			// _ = "end of CoverTab[121685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:297
			_go_fuzz_dep_.CoverTab[121686]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:297
			// _ = "end of CoverTab[121686]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:297
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:297
		// _ = "end of CoverTab[121678]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:297
		_go_fuzz_dep_.CoverTab[121679]++

													if ch == '"' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:299
			_go_fuzz_dep_.CoverTab[121687]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:300
			// _ = "end of CoverTab[121687]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:301
			_go_fuzz_dep_.CoverTab[121688]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:301
			// _ = "end of CoverTab[121688]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:301
		// _ = "end of CoverTab[121679]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:301
		_go_fuzz_dep_.CoverTab[121680]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
		if braces == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			_go_fuzz_dep_.CoverTab[121689]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			return ch == '$'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			// _ = "end of CoverTab[121689]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			_go_fuzz_dep_.CoverTab[121690]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			return s.peek() == '{'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			// _ = "end of CoverTab[121690]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:304
			_go_fuzz_dep_.CoverTab[121691]++
														braces++
														s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:306
			// _ = "end of CoverTab[121691]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
			_go_fuzz_dep_.CoverTab[121692]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
			if braces > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
				_go_fuzz_dep_.CoverTab[121693]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
				return ch == '{'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
				// _ = "end of CoverTab[121693]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:307
				_go_fuzz_dep_.CoverTab[121694]++
															braces++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:308
				// _ = "end of CoverTab[121694]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
				_go_fuzz_dep_.CoverTab[121695]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
				// _ = "end of CoverTab[121695]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
			// _ = "end of CoverTab[121692]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
		// _ = "end of CoverTab[121680]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:309
		_go_fuzz_dep_.CoverTab[121681]++
													if braces > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:310
			_go_fuzz_dep_.CoverTab[121696]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:310
			return ch == '}'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:310
			// _ = "end of CoverTab[121696]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:310
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:310
			_go_fuzz_dep_.CoverTab[121697]++
														braces--
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:311
			// _ = "end of CoverTab[121697]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:312
			_go_fuzz_dep_.CoverTab[121698]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:312
			// _ = "end of CoverTab[121698]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:312
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:312
		// _ = "end of CoverTab[121681]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:312
		_go_fuzz_dep_.CoverTab[121682]++

													if ch == '\\' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:314
			_go_fuzz_dep_.CoverTab[121699]++
														s.scanEscape()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:315
			// _ = "end of CoverTab[121699]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:316
			_go_fuzz_dep_.CoverTab[121700]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:316
			// _ = "end of CoverTab[121700]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:316
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:316
		// _ = "end of CoverTab[121682]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:317
	// _ = "end of CoverTab[121676]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:317
	_go_fuzz_dep_.CoverTab[121677]++

												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:319
	// _ = "end of CoverTab[121677]"
}

// scanEscape scans an escape sequence
func (s *Scanner) scanEscape() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:323
	_go_fuzz_dep_.CoverTab[121701]++

												ch := s.next()
												switch ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:327
		_go_fuzz_dep_.CoverTab[121703]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:327
		// _ = "end of CoverTab[121703]"

	case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:329
		_go_fuzz_dep_.CoverTab[121704]++

													ch = s.scanDigits(ch, 8, 3)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:331
		// _ = "end of CoverTab[121704]"
	case 'x':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:332
		_go_fuzz_dep_.CoverTab[121705]++

													ch = s.scanDigits(s.next(), 16, 2)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:334
		// _ = "end of CoverTab[121705]"
	case 'u':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:335
		_go_fuzz_dep_.CoverTab[121706]++

													ch = s.scanDigits(s.next(), 16, 4)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:337
		// _ = "end of CoverTab[121706]"
	case 'U':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:338
		_go_fuzz_dep_.CoverTab[121707]++

													ch = s.scanDigits(s.next(), 16, 8)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:340
		// _ = "end of CoverTab[121707]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:341
		_go_fuzz_dep_.CoverTab[121708]++
													s.err("illegal char escape")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:342
		// _ = "end of CoverTab[121708]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:343
	// _ = "end of CoverTab[121701]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:343
	_go_fuzz_dep_.CoverTab[121702]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:344
	// _ = "end of CoverTab[121702]"
}

// scanDigits scans a rune with the given base for n times. For example an
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:347
// octal notation \184 would yield in scanDigits(ch, 8, 3)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:349
func (s *Scanner) scanDigits(ch rune, base, n int) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:349
	_go_fuzz_dep_.CoverTab[121709]++
												for n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:350
		_go_fuzz_dep_.CoverTab[121712]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:350
		return digitVal(ch) < base
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:350
		// _ = "end of CoverTab[121712]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:350
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:350
		_go_fuzz_dep_.CoverTab[121713]++
													ch = s.next()
													n--
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:352
		// _ = "end of CoverTab[121713]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:353
	// _ = "end of CoverTab[121709]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:353
	_go_fuzz_dep_.CoverTab[121710]++
												if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:354
		_go_fuzz_dep_.CoverTab[121714]++
													s.err("illegal char escape")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:355
		// _ = "end of CoverTab[121714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:356
		_go_fuzz_dep_.CoverTab[121715]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:356
		// _ = "end of CoverTab[121715]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:356
	// _ = "end of CoverTab[121710]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:356
	_go_fuzz_dep_.CoverTab[121711]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:359
	s.unread()
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:360
	// _ = "end of CoverTab[121711]"
}

// scanIdentifier scans an identifier and returns the literal string
func (s *Scanner) scanIdentifier() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:364
	_go_fuzz_dep_.CoverTab[121716]++
												offs := s.srcPos.Offset - s.lastCharLen
												ch := s.next()
												for isLetter(ch) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		_go_fuzz_dep_.CoverTab[121719]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		return isDigit(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		// _ = "end of CoverTab[121719]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		_go_fuzz_dep_.CoverTab[121720]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		return ch == '-'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		// _ = "end of CoverTab[121720]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:367
		_go_fuzz_dep_.CoverTab[121721]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:368
		// _ = "end of CoverTab[121721]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:369
	// _ = "end of CoverTab[121716]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:369
	_go_fuzz_dep_.CoverTab[121717]++

												if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:371
		_go_fuzz_dep_.CoverTab[121722]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:372
		// _ = "end of CoverTab[121722]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:373
		_go_fuzz_dep_.CoverTab[121723]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:373
		// _ = "end of CoverTab[121723]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:373
	// _ = "end of CoverTab[121717]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:373
	_go_fuzz_dep_.CoverTab[121718]++

												return string(s.src[offs:s.srcPos.Offset])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:375
	// _ = "end of CoverTab[121718]"
}

// recentPosition returns the position of the character immediately after the
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:378
// character or token returned by the last call to Scan.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:380
func (s *Scanner) recentPosition() (pos token.Pos) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:380
	_go_fuzz_dep_.CoverTab[121724]++
												pos.Offset = s.srcPos.Offset - s.lastCharLen
												switch {
	case s.srcPos.Column > 0:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:383
		_go_fuzz_dep_.CoverTab[121726]++

													pos.Line = s.srcPos.Line
													pos.Column = s.srcPos.Column
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:386
		// _ = "end of CoverTab[121726]"
	case s.lastLineLen > 0:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:387
		_go_fuzz_dep_.CoverTab[121727]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:391
		pos.Line = s.srcPos.Line - 1
													pos.Column = s.lastLineLen
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:392
		// _ = "end of CoverTab[121727]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:393
		_go_fuzz_dep_.CoverTab[121728]++

													pos.Line = 1
													pos.Column = 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:396
		// _ = "end of CoverTab[121728]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:397
	// _ = "end of CoverTab[121724]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:397
	_go_fuzz_dep_.CoverTab[121725]++
												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:398
	// _ = "end of CoverTab[121725]"
}

// err prints the error of any scanning to s.Error function. If the function is
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:401
// not defined, by default it prints them to os.Stderr
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:403
func (s *Scanner) err(msg string) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:403
	_go_fuzz_dep_.CoverTab[121729]++
												s.ErrorCount++
												pos := s.recentPosition()

												if s.Error != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:407
		_go_fuzz_dep_.CoverTab[121731]++
													s.Error(pos, msg)
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:409
		// _ = "end of CoverTab[121731]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:410
		_go_fuzz_dep_.CoverTab[121732]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:410
		// _ = "end of CoverTab[121732]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:410
	// _ = "end of CoverTab[121729]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:410
	_go_fuzz_dep_.CoverTab[121730]++

												fmt.Fprintf(os.Stderr, "%s: %s\n", pos, msg)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:412
	// _ = "end of CoverTab[121730]"
}

// isHexadecimal returns true if the given rune is a letter
func isLetter(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:416
	_go_fuzz_dep_.CoverTab[121733]++
												return 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		_go_fuzz_dep_.CoverTab[121734]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		return ch <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		// _ = "end of CoverTab[121734]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		_go_fuzz_dep_.CoverTab[121735]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		return 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			_go_fuzz_dep_.CoverTab[121736]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			return ch <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			// _ = "end of CoverTab[121736]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		// _ = "end of CoverTab[121735]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		_go_fuzz_dep_.CoverTab[121737]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		return ch == '_'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		// _ = "end of CoverTab[121737]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		_go_fuzz_dep_.CoverTab[121738]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		return ch >= 0x80 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			_go_fuzz_dep_.CoverTab[121739]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			return unicode.IsLetter(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
			// _ = "end of CoverTab[121739]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
		// _ = "end of CoverTab[121738]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:417
	// _ = "end of CoverTab[121733]"
}

// isHexadecimal returns true if the given rune is a decimal digit
func isDigit(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:421
	_go_fuzz_dep_.CoverTab[121740]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		_go_fuzz_dep_.CoverTab[121741]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		// _ = "end of CoverTab[121741]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		_go_fuzz_dep_.CoverTab[121742]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		return ch >= 0x80 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
			_go_fuzz_dep_.CoverTab[121743]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
			return unicode.IsDigit(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
			// _ = "end of CoverTab[121743]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
		// _ = "end of CoverTab[121742]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:422
	// _ = "end of CoverTab[121740]"
}

// isHexadecimal returns true if the given rune is a decimal number
func isDecimal(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:426
	_go_fuzz_dep_.CoverTab[121744]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:427
		_go_fuzz_dep_.CoverTab[121745]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:427
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:427
		// _ = "end of CoverTab[121745]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:427
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:427
	// _ = "end of CoverTab[121744]"
}

// isHexadecimal returns true if the given rune is an hexadecimal number
func isHexadecimal(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:431
	_go_fuzz_dep_.CoverTab[121746]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		_go_fuzz_dep_.CoverTab[121747]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		// _ = "end of CoverTab[121747]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		_go_fuzz_dep_.CoverTab[121748]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		return 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			_go_fuzz_dep_.CoverTab[121749]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			return ch <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			// _ = "end of CoverTab[121749]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		// _ = "end of CoverTab[121748]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		_go_fuzz_dep_.CoverTab[121750]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		return 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			_go_fuzz_dep_.CoverTab[121751]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			return ch <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
			// _ = "end of CoverTab[121751]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
		// _ = "end of CoverTab[121750]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:432
	// _ = "end of CoverTab[121746]"
}

// isWhitespace returns true if the rune is a space, tab, newline or carriage return
func isWhitespace(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:436
	_go_fuzz_dep_.CoverTab[121752]++
												return ch == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		_go_fuzz_dep_.CoverTab[121753]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		return ch == '\t'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		// _ = "end of CoverTab[121753]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		_go_fuzz_dep_.CoverTab[121754]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		return ch == '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		// _ = "end of CoverTab[121754]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		_go_fuzz_dep_.CoverTab[121755]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		return ch == '\r'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
		// _ = "end of CoverTab[121755]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:437
	// _ = "end of CoverTab[121752]"
}

// digitVal returns the integer value of a given octal,decimal or hexadecimal rune
func digitVal(ch rune) int {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:441
	_go_fuzz_dep_.CoverTab[121756]++
												switch {
	case '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:443
		_go_fuzz_dep_.CoverTab[121762]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:443
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:443
		// _ = "end of CoverTab[121762]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:443
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:443
		_go_fuzz_dep_.CoverTab[121758]++
													return int(ch - '0')
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:444
		// _ = "end of CoverTab[121758]"
	case 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:445
		_go_fuzz_dep_.CoverTab[121763]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:445
		return ch <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:445
		// _ = "end of CoverTab[121763]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:445
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:445
		_go_fuzz_dep_.CoverTab[121759]++
													return int(ch - 'a' + 10)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:446
		// _ = "end of CoverTab[121759]"
	case 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:447
		_go_fuzz_dep_.CoverTab[121764]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:447
		return ch <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:447
		// _ = "end of CoverTab[121764]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:447
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:447
		_go_fuzz_dep_.CoverTab[121760]++
													return int(ch - 'A' + 10)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:448
		// _ = "end of CoverTab[121760]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:448
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:448
		_go_fuzz_dep_.CoverTab[121761]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:448
		// _ = "end of CoverTab[121761]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:449
	// _ = "end of CoverTab[121756]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:449
	_go_fuzz_dep_.CoverTab[121757]++
												return 16
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:450
	// _ = "end of CoverTab[121757]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:451
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/json/scanner/scanner.go:451
var _ = _go_fuzz_dep_.CoverTab
