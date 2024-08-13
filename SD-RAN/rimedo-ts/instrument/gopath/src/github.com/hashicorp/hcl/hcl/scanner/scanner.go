//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:1
// Package scanner implements a scanner for HCL (HashiCorp Configuration
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:1
// Language) source text.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
package scanner

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:3
)

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/hcl/hcl/token"
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
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:48
// its source content.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:50
func New(src []byte) *Scanner {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:50
	_go_fuzz_dep_.CoverTab[121080]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:54
	b := bytes.NewBuffer(src)
	s := &Scanner{
		buf:	b,
		src:	src,
	}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:61
	s.srcPos.Line = 1
												return s
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:62
	// _ = "end of CoverTab[121080]"
}

// next reads the next rune from the bufferred reader. Returns the rune(0) if
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:65
// an error occurs (or io.EOF is returned).
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:67
func (s *Scanner) next() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:67
	_go_fuzz_dep_.CoverTab[121081]++
												ch, size, err := s.buf.ReadRune()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:69
		_go_fuzz_dep_.CoverTab[121087]++

													s.srcPos.Column++
													s.srcPos.Offset += size
													s.lastCharLen = size
													return eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:74
		// _ = "end of CoverTab[121087]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:75
		_go_fuzz_dep_.CoverTab[121088]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:75
		// _ = "end of CoverTab[121088]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:75
	// _ = "end of CoverTab[121081]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:75
	_go_fuzz_dep_.CoverTab[121082]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:78
	s.prevPos = s.srcPos

	s.srcPos.Column++
	s.lastCharLen = size
	s.srcPos.Offset += size

	if ch == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:84
		_go_fuzz_dep_.CoverTab[121089]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:84
		return size == 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:84
		// _ = "end of CoverTab[121089]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:84
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:84
		_go_fuzz_dep_.CoverTab[121090]++
													s.err("illegal UTF-8 encoding")
													return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:86
		// _ = "end of CoverTab[121090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:87
		_go_fuzz_dep_.CoverTab[121091]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:87
		// _ = "end of CoverTab[121091]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:87
	// _ = "end of CoverTab[121082]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:87
	_go_fuzz_dep_.CoverTab[121083]++

												if ch == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:89
		_go_fuzz_dep_.CoverTab[121092]++
													s.srcPos.Line++
													s.lastLineLen = s.srcPos.Column
													s.srcPos.Column = 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:92
		// _ = "end of CoverTab[121092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:93
		_go_fuzz_dep_.CoverTab[121093]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:93
		// _ = "end of CoverTab[121093]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:93
	// _ = "end of CoverTab[121083]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:93
	_go_fuzz_dep_.CoverTab[121084]++

												if ch == '\x00' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:95
		_go_fuzz_dep_.CoverTab[121094]++
													s.err("unexpected null character (0x00)")
													return eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:97
		// _ = "end of CoverTab[121094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:98
		_go_fuzz_dep_.CoverTab[121095]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:98
		// _ = "end of CoverTab[121095]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:98
	// _ = "end of CoverTab[121084]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:98
	_go_fuzz_dep_.CoverTab[121085]++

												if ch == '\uE123' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:100
		_go_fuzz_dep_.CoverTab[121096]++
													s.err("unicode code point U+E123 reserved for internal use")
													return utf8.RuneError
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:102
		// _ = "end of CoverTab[121096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:103
		_go_fuzz_dep_.CoverTab[121097]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:103
		// _ = "end of CoverTab[121097]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:103
	// _ = "end of CoverTab[121085]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:103
	_go_fuzz_dep_.CoverTab[121086]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:107
	return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:107
	// _ = "end of CoverTab[121086]"
}

// unread unreads the previous read Rune and updates the source position
func (s *Scanner) unread() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:111
	_go_fuzz_dep_.CoverTab[121098]++
												if err := s.buf.UnreadRune(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:112
		_go_fuzz_dep_.CoverTab[121100]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:113
		// _ = "end of CoverTab[121100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:114
		_go_fuzz_dep_.CoverTab[121101]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:114
		// _ = "end of CoverTab[121101]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:114
	// _ = "end of CoverTab[121098]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:114
	_go_fuzz_dep_.CoverTab[121099]++
												s.srcPos = s.prevPos
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:115
	// _ = "end of CoverTab[121099]"
}

// peek returns the next rune without advancing the reader.
func (s *Scanner) peek() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:119
	_go_fuzz_dep_.CoverTab[121102]++
												peek, _, err := s.buf.ReadRune()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:121
		_go_fuzz_dep_.CoverTab[121104]++
													return eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:122
		// _ = "end of CoverTab[121104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:123
		_go_fuzz_dep_.CoverTab[121105]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:123
		// _ = "end of CoverTab[121105]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:123
	// _ = "end of CoverTab[121102]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:123
	_go_fuzz_dep_.CoverTab[121103]++

												s.buf.UnreadRune()
												return peek
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:126
	// _ = "end of CoverTab[121103]"
}

// Scan scans the next token and returns the token.
func (s *Scanner) Scan() token.Token {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:130
	_go_fuzz_dep_.CoverTab[121106]++
												ch := s.next()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:134
	for isWhitespace(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:134
		_go_fuzz_dep_.CoverTab[121111]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:135
		// _ = "end of CoverTab[121111]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:136
	// _ = "end of CoverTab[121106]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:136
	_go_fuzz_dep_.CoverTab[121107]++

												var tok token.Type

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:141
	s.tokStart = s.srcPos.Offset - s.lastCharLen

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:145
	s.tokPos.Offset = s.srcPos.Offset - s.lastCharLen
	if s.srcPos.Column > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:146
		_go_fuzz_dep_.CoverTab[121112]++

													s.tokPos.Line = s.srcPos.Line
													s.tokPos.Column = s.srcPos.Column
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:149
		// _ = "end of CoverTab[121112]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:150
		_go_fuzz_dep_.CoverTab[121113]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:154
		s.tokPos.Line = s.srcPos.Line - 1
													s.tokPos.Column = s.lastLineLen
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:155
		// _ = "end of CoverTab[121113]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:156
	// _ = "end of CoverTab[121107]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:156
	_go_fuzz_dep_.CoverTab[121108]++

												switch {
	case isLetter(ch):
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:159
		_go_fuzz_dep_.CoverTab[121114]++
													tok = token.IDENT
													lit := s.scanIdentifier()
													if lit == "true" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:162
			_go_fuzz_dep_.CoverTab[121117]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:162
			return lit == "false"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:162
			// _ = "end of CoverTab[121117]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:162
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:162
			_go_fuzz_dep_.CoverTab[121118]++
														tok = token.BOOL
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:163
			// _ = "end of CoverTab[121118]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:164
			_go_fuzz_dep_.CoverTab[121119]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:164
			// _ = "end of CoverTab[121119]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:164
		// _ = "end of CoverTab[121114]"
	case isDecimal(ch):
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:165
		_go_fuzz_dep_.CoverTab[121115]++
													tok = s.scanNumber(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:166
		// _ = "end of CoverTab[121115]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:167
		_go_fuzz_dep_.CoverTab[121116]++
													switch ch {
		case eof:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:169
			_go_fuzz_dep_.CoverTab[121120]++
														tok = token.EOF
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:170
			// _ = "end of CoverTab[121120]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:171
			_go_fuzz_dep_.CoverTab[121121]++
														tok = token.STRING
														s.scanString()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:173
			// _ = "end of CoverTab[121121]"
		case '#', '/':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:174
			_go_fuzz_dep_.CoverTab[121122]++
														tok = token.COMMENT
														s.scanComment(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:176
			// _ = "end of CoverTab[121122]"
		case '.':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:177
			_go_fuzz_dep_.CoverTab[121123]++
														tok = token.PERIOD
														ch = s.peek()
														if isDecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:180
				_go_fuzz_dep_.CoverTab[121134]++
															tok = token.FLOAT
															ch = s.scanMantissa(ch)
															ch = s.scanExponent(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:183
				// _ = "end of CoverTab[121134]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:184
				_go_fuzz_dep_.CoverTab[121135]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:184
				// _ = "end of CoverTab[121135]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:184
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:184
			// _ = "end of CoverTab[121123]"
		case '<':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:185
			_go_fuzz_dep_.CoverTab[121124]++
														tok = token.HEREDOC
														s.scanHeredoc()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:187
			// _ = "end of CoverTab[121124]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:188
			_go_fuzz_dep_.CoverTab[121125]++
														tok = token.LBRACK
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:189
			// _ = "end of CoverTab[121125]"
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:190
			_go_fuzz_dep_.CoverTab[121126]++
														tok = token.RBRACK
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:191
			// _ = "end of CoverTab[121126]"
		case '{':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:192
			_go_fuzz_dep_.CoverTab[121127]++
														tok = token.LBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:193
			// _ = "end of CoverTab[121127]"
		case '}':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:194
			_go_fuzz_dep_.CoverTab[121128]++
														tok = token.RBRACE
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:195
			// _ = "end of CoverTab[121128]"
		case ',':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:196
			_go_fuzz_dep_.CoverTab[121129]++
														tok = token.COMMA
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:197
			// _ = "end of CoverTab[121129]"
		case '=':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:198
			_go_fuzz_dep_.CoverTab[121130]++
														tok = token.ASSIGN
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:199
			// _ = "end of CoverTab[121130]"
		case '+':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:200
			_go_fuzz_dep_.CoverTab[121131]++
														tok = token.ADD
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:201
			// _ = "end of CoverTab[121131]"
		case '-':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:202
			_go_fuzz_dep_.CoverTab[121132]++
														if isDecimal(s.peek()) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:203
				_go_fuzz_dep_.CoverTab[121136]++
															ch := s.next()
															tok = s.scanNumber(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:205
				// _ = "end of CoverTab[121136]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:206
				_go_fuzz_dep_.CoverTab[121137]++
															tok = token.SUB
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:207
				// _ = "end of CoverTab[121137]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:208
			// _ = "end of CoverTab[121132]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:209
			_go_fuzz_dep_.CoverTab[121133]++
														s.err("illegal char")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:210
			// _ = "end of CoverTab[121133]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:211
		// _ = "end of CoverTab[121116]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:212
	// _ = "end of CoverTab[121108]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:212
	_go_fuzz_dep_.CoverTab[121109]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:215
	s.tokEnd = s.srcPos.Offset

	// create token literal
	var tokenText string
	if s.tokStart >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:219
		_go_fuzz_dep_.CoverTab[121138]++
													tokenText = string(s.src[s.tokStart:s.tokEnd])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:220
		// _ = "end of CoverTab[121138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:221
		_go_fuzz_dep_.CoverTab[121139]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:221
		// _ = "end of CoverTab[121139]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:221
	// _ = "end of CoverTab[121109]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:221
	_go_fuzz_dep_.CoverTab[121110]++
												s.tokStart = s.tokEnd

												return token.Token{
		Type:	tok,
		Pos:	s.tokPos,
		Text:	tokenText,
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:228
	// _ = "end of CoverTab[121110]"
}

func (s *Scanner) scanComment(ch rune) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:231
	_go_fuzz_dep_.CoverTab[121140]++

												if ch == '#' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
		_go_fuzz_dep_.CoverTab[121143]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
		return (ch == '/' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
			_go_fuzz_dep_.CoverTab[121144]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
			return s.peek() != '*'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
			// _ = "end of CoverTab[121144]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
		}())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
		// _ = "end of CoverTab[121143]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:233
		_go_fuzz_dep_.CoverTab[121145]++
													if ch == '/' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:234
			_go_fuzz_dep_.CoverTab[121149]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:234
			return s.peek() != '/'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:234
			// _ = "end of CoverTab[121149]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:234
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:234
			_go_fuzz_dep_.CoverTab[121150]++
														s.err("expected '/' for comment")
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:236
			// _ = "end of CoverTab[121150]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:237
			_go_fuzz_dep_.CoverTab[121151]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:237
			// _ = "end of CoverTab[121151]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:237
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:237
		// _ = "end of CoverTab[121145]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:237
		_go_fuzz_dep_.CoverTab[121146]++

													ch = s.next()
													for ch != '\n' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			_go_fuzz_dep_.CoverTab[121152]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			return ch >= 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			// _ = "end of CoverTab[121152]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			_go_fuzz_dep_.CoverTab[121153]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			return ch != eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			// _ = "end of CoverTab[121153]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:240
			_go_fuzz_dep_.CoverTab[121154]++
														ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:241
			// _ = "end of CoverTab[121154]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:242
		// _ = "end of CoverTab[121146]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:242
		_go_fuzz_dep_.CoverTab[121147]++
													if ch != eof && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:243
			_go_fuzz_dep_.CoverTab[121155]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:243
			return ch >= 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:243
			// _ = "end of CoverTab[121155]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:243
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:243
			_go_fuzz_dep_.CoverTab[121156]++
														s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:244
			// _ = "end of CoverTab[121156]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:245
			_go_fuzz_dep_.CoverTab[121157]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:245
			// _ = "end of CoverTab[121157]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:245
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:245
		// _ = "end of CoverTab[121147]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:245
		_go_fuzz_dep_.CoverTab[121148]++
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:246
		// _ = "end of CoverTab[121148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:247
		_go_fuzz_dep_.CoverTab[121158]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:247
		// _ = "end of CoverTab[121158]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:247
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:247
	// _ = "end of CoverTab[121140]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:247
	_go_fuzz_dep_.CoverTab[121141]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:251
	if ch == '/' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:251
		_go_fuzz_dep_.CoverTab[121159]++
													s.next()
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:253
		// _ = "end of CoverTab[121159]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:254
		_go_fuzz_dep_.CoverTab[121160]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:254
		// _ = "end of CoverTab[121160]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:254
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:254
	// _ = "end of CoverTab[121141]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:254
	_go_fuzz_dep_.CoverTab[121142]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:257
	for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:257
		_go_fuzz_dep_.CoverTab[121161]++
													if ch < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:258
			_go_fuzz_dep_.CoverTab[121163]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:258
			return ch == eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:258
			// _ = "end of CoverTab[121163]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:258
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:258
			_go_fuzz_dep_.CoverTab[121164]++
														s.err("comment not terminated")
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:260
			// _ = "end of CoverTab[121164]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:261
			_go_fuzz_dep_.CoverTab[121165]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:261
			// _ = "end of CoverTab[121165]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:261
		// _ = "end of CoverTab[121161]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:261
		_go_fuzz_dep_.CoverTab[121162]++

													ch0 := ch
													ch = s.next()
													if ch0 == '*' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:265
			_go_fuzz_dep_.CoverTab[121166]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:265
			return ch == '/'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:265
			// _ = "end of CoverTab[121166]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:265
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:265
			_go_fuzz_dep_.CoverTab[121167]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:266
			// _ = "end of CoverTab[121167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:267
			_go_fuzz_dep_.CoverTab[121168]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:267
			// _ = "end of CoverTab[121168]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:267
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:267
		// _ = "end of CoverTab[121162]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:268
	// _ = "end of CoverTab[121142]"
}

// scanNumber scans a HCL number definition starting with the given rune
func (s *Scanner) scanNumber(ch rune) token.Type {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:272
	_go_fuzz_dep_.CoverTab[121169]++
												if ch == '0' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:273
		_go_fuzz_dep_.CoverTab[121174]++

													ch = s.next()
													if ch == 'x' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:276
			_go_fuzz_dep_.CoverTab[121181]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:276
			return ch == 'X'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:276
			// _ = "end of CoverTab[121181]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:276
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:276
			_go_fuzz_dep_.CoverTab[121182]++

														ch = s.next()
														found := false
														for isHexadecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:280
				_go_fuzz_dep_.CoverTab[121186]++
															ch = s.next()
															found = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:282
				// _ = "end of CoverTab[121186]"
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:283
			// _ = "end of CoverTab[121182]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:283
			_go_fuzz_dep_.CoverTab[121183]++

														if !found {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:285
				_go_fuzz_dep_.CoverTab[121187]++
															s.err("illegal hexadecimal number")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:286
				// _ = "end of CoverTab[121187]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:287
				_go_fuzz_dep_.CoverTab[121188]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:287
				// _ = "end of CoverTab[121188]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:287
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:287
			// _ = "end of CoverTab[121183]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:287
			_go_fuzz_dep_.CoverTab[121184]++

														if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:289
				_go_fuzz_dep_.CoverTab[121189]++
															s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:290
				// _ = "end of CoverTab[121189]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:291
				_go_fuzz_dep_.CoverTab[121190]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:291
				// _ = "end of CoverTab[121190]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:291
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:291
			// _ = "end of CoverTab[121184]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:291
			_go_fuzz_dep_.CoverTab[121185]++

														return token.NUMBER
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:293
			// _ = "end of CoverTab[121185]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:294
			_go_fuzz_dep_.CoverTab[121191]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:294
			// _ = "end of CoverTab[121191]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:294
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:294
		// _ = "end of CoverTab[121174]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:294
		_go_fuzz_dep_.CoverTab[121175]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:297
		illegalOctal := false
		for isDecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:298
			_go_fuzz_dep_.CoverTab[121192]++
														ch = s.next()
														if ch == '8' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:300
				_go_fuzz_dep_.CoverTab[121193]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:300
				return ch == '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:300
				// _ = "end of CoverTab[121193]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:300
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:300
				_go_fuzz_dep_.CoverTab[121194]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:304
				illegalOctal = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:304
				// _ = "end of CoverTab[121194]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:305
				_go_fuzz_dep_.CoverTab[121195]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:305
				// _ = "end of CoverTab[121195]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:305
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:305
			// _ = "end of CoverTab[121192]"
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:306
		// _ = "end of CoverTab[121175]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:306
		_go_fuzz_dep_.CoverTab[121176]++

													if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:308
			_go_fuzz_dep_.CoverTab[121196]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:308
			return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:308
			// _ = "end of CoverTab[121196]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:308
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:308
			_go_fuzz_dep_.CoverTab[121197]++
														ch = s.scanExponent(ch)
														return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:310
			// _ = "end of CoverTab[121197]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:311
			_go_fuzz_dep_.CoverTab[121198]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:311
			// _ = "end of CoverTab[121198]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:311
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:311
		// _ = "end of CoverTab[121176]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:311
		_go_fuzz_dep_.CoverTab[121177]++

													if ch == '.' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:313
			_go_fuzz_dep_.CoverTab[121199]++
														ch = s.scanFraction(ch)

														if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:316
				_go_fuzz_dep_.CoverTab[121201]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:316
				return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:316
				// _ = "end of CoverTab[121201]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:316
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:316
				_go_fuzz_dep_.CoverTab[121202]++
															ch = s.next()
															ch = s.scanExponent(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:318
				// _ = "end of CoverTab[121202]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:319
				_go_fuzz_dep_.CoverTab[121203]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:319
				// _ = "end of CoverTab[121203]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:319
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:319
			// _ = "end of CoverTab[121199]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:319
			_go_fuzz_dep_.CoverTab[121200]++
														return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:320
			// _ = "end of CoverTab[121200]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:321
			_go_fuzz_dep_.CoverTab[121204]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:321
			// _ = "end of CoverTab[121204]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:321
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:321
		// _ = "end of CoverTab[121177]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:321
		_go_fuzz_dep_.CoverTab[121178]++

													if illegalOctal {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:323
			_go_fuzz_dep_.CoverTab[121205]++
														s.err("illegal octal number")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:324
			// _ = "end of CoverTab[121205]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:325
			_go_fuzz_dep_.CoverTab[121206]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:325
			// _ = "end of CoverTab[121206]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:325
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:325
		// _ = "end of CoverTab[121178]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:325
		_go_fuzz_dep_.CoverTab[121179]++

													if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:327
			_go_fuzz_dep_.CoverTab[121207]++
														s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:328
			// _ = "end of CoverTab[121207]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:329
			_go_fuzz_dep_.CoverTab[121208]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:329
			// _ = "end of CoverTab[121208]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:329
		// _ = "end of CoverTab[121179]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:329
		_go_fuzz_dep_.CoverTab[121180]++
													return token.NUMBER
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:330
		// _ = "end of CoverTab[121180]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:331
		_go_fuzz_dep_.CoverTab[121209]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:331
		// _ = "end of CoverTab[121209]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:331
	// _ = "end of CoverTab[121169]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:331
	_go_fuzz_dep_.CoverTab[121170]++

												s.scanMantissa(ch)
												ch = s.next()
												if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:335
		_go_fuzz_dep_.CoverTab[121210]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:335
		return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:335
		// _ = "end of CoverTab[121210]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:335
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:335
		_go_fuzz_dep_.CoverTab[121211]++
													ch = s.scanExponent(ch)
													return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:337
		// _ = "end of CoverTab[121211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:338
		_go_fuzz_dep_.CoverTab[121212]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:338
		// _ = "end of CoverTab[121212]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:338
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:338
	// _ = "end of CoverTab[121170]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:338
	_go_fuzz_dep_.CoverTab[121171]++

												if ch == '.' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:340
		_go_fuzz_dep_.CoverTab[121213]++
													ch = s.scanFraction(ch)
													if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:342
			_go_fuzz_dep_.CoverTab[121215]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:342
			return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:342
			// _ = "end of CoverTab[121215]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:342
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:342
			_go_fuzz_dep_.CoverTab[121216]++
														ch = s.next()
														ch = s.scanExponent(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:344
			// _ = "end of CoverTab[121216]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:345
			_go_fuzz_dep_.CoverTab[121217]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:345
			// _ = "end of CoverTab[121217]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:345
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:345
		// _ = "end of CoverTab[121213]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:345
		_go_fuzz_dep_.CoverTab[121214]++
													return token.FLOAT
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:346
		// _ = "end of CoverTab[121214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:347
		_go_fuzz_dep_.CoverTab[121218]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:347
		// _ = "end of CoverTab[121218]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:347
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:347
	// _ = "end of CoverTab[121171]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:347
	_go_fuzz_dep_.CoverTab[121172]++

												if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:349
		_go_fuzz_dep_.CoverTab[121219]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:350
		// _ = "end of CoverTab[121219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:351
		_go_fuzz_dep_.CoverTab[121220]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:351
		// _ = "end of CoverTab[121220]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:351
	// _ = "end of CoverTab[121172]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:351
	_go_fuzz_dep_.CoverTab[121173]++
												return token.NUMBER
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:352
	// _ = "end of CoverTab[121173]"
}

// scanMantissa scans the mantissa beginning from the rune. It returns the next
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:355
// non decimal rune. It's used to determine wheter it's a fraction or exponent.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:357
func (s *Scanner) scanMantissa(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:357
	_go_fuzz_dep_.CoverTab[121221]++
												scanned := false
												for isDecimal(ch) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:359
		_go_fuzz_dep_.CoverTab[121224]++
													ch = s.next()
													scanned = true
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:361
		// _ = "end of CoverTab[121224]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:362
	// _ = "end of CoverTab[121221]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:362
	_go_fuzz_dep_.CoverTab[121222]++

												if scanned && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:364
		_go_fuzz_dep_.CoverTab[121225]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:364
		return ch != eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:364
		// _ = "end of CoverTab[121225]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:364
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:364
		_go_fuzz_dep_.CoverTab[121226]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:365
		// _ = "end of CoverTab[121226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:366
		_go_fuzz_dep_.CoverTab[121227]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:366
		// _ = "end of CoverTab[121227]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:366
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:366
	// _ = "end of CoverTab[121222]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:366
	_go_fuzz_dep_.CoverTab[121223]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:367
	// _ = "end of CoverTab[121223]"
}

// scanFraction scans the fraction after the '.' rune
func (s *Scanner) scanFraction(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:371
	_go_fuzz_dep_.CoverTab[121228]++
												if ch == '.' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:372
		_go_fuzz_dep_.CoverTab[121230]++
													ch = s.peek()
													ch = s.scanMantissa(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:374
		// _ = "end of CoverTab[121230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:375
		_go_fuzz_dep_.CoverTab[121231]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:375
		// _ = "end of CoverTab[121231]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:375
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:375
	// _ = "end of CoverTab[121228]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:375
	_go_fuzz_dep_.CoverTab[121229]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:376
	// _ = "end of CoverTab[121229]"
}

// scanExponent scans the remaining parts of an exponent after the 'e' or 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:379
// rune.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:381
func (s *Scanner) scanExponent(ch rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:381
	_go_fuzz_dep_.CoverTab[121232]++
												if ch == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:382
		_go_fuzz_dep_.CoverTab[121234]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:382
		return ch == 'E'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:382
		// _ = "end of CoverTab[121234]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:382
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:382
		_go_fuzz_dep_.CoverTab[121235]++
													ch = s.next()
													if ch == '-' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:384
			_go_fuzz_dep_.CoverTab[121237]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:384
			return ch == '+'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:384
			// _ = "end of CoverTab[121237]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:384
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:384
			_go_fuzz_dep_.CoverTab[121238]++
														ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:385
			// _ = "end of CoverTab[121238]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:386
			_go_fuzz_dep_.CoverTab[121239]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:386
			// _ = "end of CoverTab[121239]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:386
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:386
		// _ = "end of CoverTab[121235]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:386
		_go_fuzz_dep_.CoverTab[121236]++
													ch = s.scanMantissa(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:387
		// _ = "end of CoverTab[121236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:388
		_go_fuzz_dep_.CoverTab[121240]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:388
		// _ = "end of CoverTab[121240]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:388
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:388
	// _ = "end of CoverTab[121232]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:388
	_go_fuzz_dep_.CoverTab[121233]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:389
	// _ = "end of CoverTab[121233]"
}

// scanHeredoc scans a heredoc string
func (s *Scanner) scanHeredoc() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:393
	_go_fuzz_dep_.CoverTab[121241]++

												if s.next() != '<' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:395
		_go_fuzz_dep_.CoverTab[121251]++
													s.err("heredoc expected second '<', didn't see it")
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:397
		// _ = "end of CoverTab[121251]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:398
		_go_fuzz_dep_.CoverTab[121252]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:398
		// _ = "end of CoverTab[121252]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:398
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:398
	// _ = "end of CoverTab[121241]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:398
	_go_fuzz_dep_.CoverTab[121242]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:401
	offs := s.srcPos.Offset

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:404
	ch := s.next()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:407
	if ch == '-' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:407
		_go_fuzz_dep_.CoverTab[121253]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:408
		// _ = "end of CoverTab[121253]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:409
		_go_fuzz_dep_.CoverTab[121254]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:409
		// _ = "end of CoverTab[121254]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:409
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:409
	// _ = "end of CoverTab[121242]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:409
	_go_fuzz_dep_.CoverTab[121243]++

												for isLetter(ch) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:411
		_go_fuzz_dep_.CoverTab[121255]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:411
		return isDigit(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:411
		// _ = "end of CoverTab[121255]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:411
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:411
		_go_fuzz_dep_.CoverTab[121256]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:412
		// _ = "end of CoverTab[121256]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:413
	// _ = "end of CoverTab[121243]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:413
	_go_fuzz_dep_.CoverTab[121244]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:416
	if ch == eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:416
		_go_fuzz_dep_.CoverTab[121257]++
													s.err("heredoc not terminated")
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:418
		// _ = "end of CoverTab[121257]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:419
		_go_fuzz_dep_.CoverTab[121258]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:419
		// _ = "end of CoverTab[121258]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:419
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:419
	// _ = "end of CoverTab[121244]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:419
	_go_fuzz_dep_.CoverTab[121245]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:422
	if ch == '\r' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:422
		_go_fuzz_dep_.CoverTab[121259]++
													if s.peek() == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:423
			_go_fuzz_dep_.CoverTab[121260]++
														ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:424
			// _ = "end of CoverTab[121260]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:425
			_go_fuzz_dep_.CoverTab[121261]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:425
			// _ = "end of CoverTab[121261]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:425
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:425
		// _ = "end of CoverTab[121259]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:426
		_go_fuzz_dep_.CoverTab[121262]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:426
		// _ = "end of CoverTab[121262]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:426
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:426
	// _ = "end of CoverTab[121245]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:426
	_go_fuzz_dep_.CoverTab[121246]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:429
	if ch != '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:429
		_go_fuzz_dep_.CoverTab[121263]++
													s.err("invalid characters in heredoc anchor")
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:431
		// _ = "end of CoverTab[121263]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:432
		_go_fuzz_dep_.CoverTab[121264]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:432
		// _ = "end of CoverTab[121264]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:432
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:432
	// _ = "end of CoverTab[121246]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:432
	_go_fuzz_dep_.CoverTab[121247]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:435
	identBytes := s.src[offs : s.srcPos.Offset-s.lastCharLen]
	if len(identBytes) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
		_go_fuzz_dep_.CoverTab[121265]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
		return (len(identBytes) == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
			_go_fuzz_dep_.CoverTab[121266]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
			return identBytes[0] == '-'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
			// _ = "end of CoverTab[121266]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
		}())
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
		// _ = "end of CoverTab[121265]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:436
		_go_fuzz_dep_.CoverTab[121267]++
													s.err("zero-length heredoc anchor")
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:438
		// _ = "end of CoverTab[121267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:439
		_go_fuzz_dep_.CoverTab[121268]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:439
		// _ = "end of CoverTab[121268]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:439
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:439
	// _ = "end of CoverTab[121247]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:439
	_go_fuzz_dep_.CoverTab[121248]++

												var identRegexp *regexp.Regexp
												if identBytes[0] == '-' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:442
		_go_fuzz_dep_.CoverTab[121269]++
													identRegexp = regexp.MustCompile(fmt.Sprintf(`^[[:space:]]*%s\r*\z`, identBytes[1:]))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:443
		// _ = "end of CoverTab[121269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:444
		_go_fuzz_dep_.CoverTab[121270]++
													identRegexp = regexp.MustCompile(fmt.Sprintf(`^[[:space:]]*%s\r*\z`, identBytes))
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:445
		// _ = "end of CoverTab[121270]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:446
	// _ = "end of CoverTab[121248]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:446
	_go_fuzz_dep_.CoverTab[121249]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:449
	lineStart := s.srcPos.Offset
	for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:450
		_go_fuzz_dep_.CoverTab[121271]++
													ch := s.next()

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:454
		if ch == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:454
			_go_fuzz_dep_.CoverTab[121273]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:458
			lineBytesLen := s.srcPos.Offset - s.lastCharLen - lineStart
			if lineBytesLen >= len(identBytes) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:459
				_go_fuzz_dep_.CoverTab[121275]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:459
				return identRegexp.Match(s.src[lineStart : s.srcPos.Offset-s.lastCharLen])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:459
				// _ = "end of CoverTab[121275]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:459
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:459
				_go_fuzz_dep_.CoverTab[121276]++
															break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:460
				// _ = "end of CoverTab[121276]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:461
				_go_fuzz_dep_.CoverTab[121277]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:461
				// _ = "end of CoverTab[121277]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:461
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:461
			// _ = "end of CoverTab[121273]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:461
			_go_fuzz_dep_.CoverTab[121274]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:464
			lineStart = s.srcPos.Offset
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:464
			// _ = "end of CoverTab[121274]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:465
			_go_fuzz_dep_.CoverTab[121278]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:465
			// _ = "end of CoverTab[121278]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:465
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:465
		// _ = "end of CoverTab[121271]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:465
		_go_fuzz_dep_.CoverTab[121272]++

													if ch == eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:467
			_go_fuzz_dep_.CoverTab[121279]++
														s.err("heredoc not terminated")
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:469
			// _ = "end of CoverTab[121279]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:470
			_go_fuzz_dep_.CoverTab[121280]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:470
			// _ = "end of CoverTab[121280]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:470
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:470
		// _ = "end of CoverTab[121272]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:471
	// _ = "end of CoverTab[121249]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:471
	_go_fuzz_dep_.CoverTab[121250]++

												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:473
	// _ = "end of CoverTab[121250]"
}

// scanString scans a quoted string
func (s *Scanner) scanString() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:477
	_go_fuzz_dep_.CoverTab[121281]++
												braces := 0
												for {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:479
		_go_fuzz_dep_.CoverTab[121283]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:482
		ch := s.next()

		if (ch == '\n' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			_go_fuzz_dep_.CoverTab[121288]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			return braces == 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			// _ = "end of CoverTab[121288]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			_go_fuzz_dep_.CoverTab[121289]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			return ch < 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			// _ = "end of CoverTab[121289]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			_go_fuzz_dep_.CoverTab[121290]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			return ch == eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			// _ = "end of CoverTab[121290]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:484
			_go_fuzz_dep_.CoverTab[121291]++
														s.err("literal not terminated")
														return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:486
			// _ = "end of CoverTab[121291]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:487
			_go_fuzz_dep_.CoverTab[121292]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:487
			// _ = "end of CoverTab[121292]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:487
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:487
		// _ = "end of CoverTab[121283]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:487
		_go_fuzz_dep_.CoverTab[121284]++

													if ch == '"' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:489
			_go_fuzz_dep_.CoverTab[121293]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:489
			return braces == 0
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:489
			// _ = "end of CoverTab[121293]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:489
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:489
			_go_fuzz_dep_.CoverTab[121294]++
														break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:490
			// _ = "end of CoverTab[121294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:491
			_go_fuzz_dep_.CoverTab[121295]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:491
			// _ = "end of CoverTab[121295]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:491
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:491
		// _ = "end of CoverTab[121284]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:491
		_go_fuzz_dep_.CoverTab[121285]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
		if braces == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			_go_fuzz_dep_.CoverTab[121296]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			return ch == '$'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			// _ = "end of CoverTab[121296]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			_go_fuzz_dep_.CoverTab[121297]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			return s.peek() == '{'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			// _ = "end of CoverTab[121297]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:494
			_go_fuzz_dep_.CoverTab[121298]++
														braces++
														s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:496
			// _ = "end of CoverTab[121298]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
			_go_fuzz_dep_.CoverTab[121299]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
			if braces > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
				_go_fuzz_dep_.CoverTab[121300]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
				return ch == '{'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
				// _ = "end of CoverTab[121300]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
			}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:497
				_go_fuzz_dep_.CoverTab[121301]++
															braces++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:498
				// _ = "end of CoverTab[121301]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
				_go_fuzz_dep_.CoverTab[121302]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
				// _ = "end of CoverTab[121302]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
			}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
			// _ = "end of CoverTab[121299]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
		// _ = "end of CoverTab[121285]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:499
		_go_fuzz_dep_.CoverTab[121286]++
													if braces > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:500
			_go_fuzz_dep_.CoverTab[121303]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:500
			return ch == '}'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:500
			// _ = "end of CoverTab[121303]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:500
		}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:500
			_go_fuzz_dep_.CoverTab[121304]++
														braces--
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:501
			// _ = "end of CoverTab[121304]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:502
			_go_fuzz_dep_.CoverTab[121305]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:502
			// _ = "end of CoverTab[121305]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:502
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:502
		// _ = "end of CoverTab[121286]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:502
		_go_fuzz_dep_.CoverTab[121287]++

													if ch == '\\' {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:504
			_go_fuzz_dep_.CoverTab[121306]++
														s.scanEscape()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:505
			// _ = "end of CoverTab[121306]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:506
			_go_fuzz_dep_.CoverTab[121307]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:506
			// _ = "end of CoverTab[121307]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:506
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:506
		// _ = "end of CoverTab[121287]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:507
	// _ = "end of CoverTab[121281]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:507
	_go_fuzz_dep_.CoverTab[121282]++

												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:509
	// _ = "end of CoverTab[121282]"
}

// scanEscape scans an escape sequence
func (s *Scanner) scanEscape() rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:513
	_go_fuzz_dep_.CoverTab[121308]++

												ch := s.next()
												switch ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:517
		_go_fuzz_dep_.CoverTab[121310]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:517
		// _ = "end of CoverTab[121310]"

	case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:519
		_go_fuzz_dep_.CoverTab[121311]++

													ch = s.scanDigits(ch, 8, 3)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:521
		// _ = "end of CoverTab[121311]"
	case 'x':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:522
		_go_fuzz_dep_.CoverTab[121312]++

													ch = s.scanDigits(s.next(), 16, 2)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:524
		// _ = "end of CoverTab[121312]"
	case 'u':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:525
		_go_fuzz_dep_.CoverTab[121313]++

													ch = s.scanDigits(s.next(), 16, 4)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:527
		// _ = "end of CoverTab[121313]"
	case 'U':
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:528
		_go_fuzz_dep_.CoverTab[121314]++

													ch = s.scanDigits(s.next(), 16, 8)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:530
		// _ = "end of CoverTab[121314]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:531
		_go_fuzz_dep_.CoverTab[121315]++
													s.err("illegal char escape")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:532
		// _ = "end of CoverTab[121315]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:533
	// _ = "end of CoverTab[121308]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:533
	_go_fuzz_dep_.CoverTab[121309]++
												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:534
	// _ = "end of CoverTab[121309]"
}

// scanDigits scans a rune with the given base for n times. For example an
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:537
// octal notation \184 would yield in scanDigits(ch, 8, 3)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:539
func (s *Scanner) scanDigits(ch rune, base, n int) rune {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:539
	_go_fuzz_dep_.CoverTab[121316]++
												start := n
												for n > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:541
		_go_fuzz_dep_.CoverTab[121320]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:541
		return digitVal(ch) < base
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:541
		// _ = "end of CoverTab[121320]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:541
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:541
		_go_fuzz_dep_.CoverTab[121321]++
													ch = s.next()
													if ch == eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:543
			_go_fuzz_dep_.CoverTab[121323]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:546
			break
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:546
			// _ = "end of CoverTab[121323]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:547
			_go_fuzz_dep_.CoverTab[121324]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:547
			// _ = "end of CoverTab[121324]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:547
		}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:547
		// _ = "end of CoverTab[121321]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:547
		_go_fuzz_dep_.CoverTab[121322]++

													n--
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:549
		// _ = "end of CoverTab[121322]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:550
	// _ = "end of CoverTab[121316]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:550
	_go_fuzz_dep_.CoverTab[121317]++
												if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:551
		_go_fuzz_dep_.CoverTab[121325]++
													s.err("illegal char escape")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:552
		// _ = "end of CoverTab[121325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:553
		_go_fuzz_dep_.CoverTab[121326]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:553
		// _ = "end of CoverTab[121326]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:553
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:553
	// _ = "end of CoverTab[121317]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:553
	_go_fuzz_dep_.CoverTab[121318]++

												if n != start && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:555
		_go_fuzz_dep_.CoverTab[121327]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:555
		return ch != eof
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:555
		// _ = "end of CoverTab[121327]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:555
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:555
		_go_fuzz_dep_.CoverTab[121328]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:558
		s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:558
		// _ = "end of CoverTab[121328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:559
		_go_fuzz_dep_.CoverTab[121329]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:559
		// _ = "end of CoverTab[121329]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:559
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:559
	// _ = "end of CoverTab[121318]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:559
	_go_fuzz_dep_.CoverTab[121319]++

												return ch
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:561
	// _ = "end of CoverTab[121319]"
}

// scanIdentifier scans an identifier and returns the literal string
func (s *Scanner) scanIdentifier() string {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:565
	_go_fuzz_dep_.CoverTab[121330]++
												offs := s.srcPos.Offset - s.lastCharLen
												ch := s.next()
												for isLetter(ch) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		_go_fuzz_dep_.CoverTab[121333]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		return isDigit(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		// _ = "end of CoverTab[121333]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		_go_fuzz_dep_.CoverTab[121334]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		return ch == '-'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		// _ = "end of CoverTab[121334]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		_go_fuzz_dep_.CoverTab[121335]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		return ch == '.'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		// _ = "end of CoverTab[121335]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:568
		_go_fuzz_dep_.CoverTab[121336]++
													ch = s.next()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:569
		// _ = "end of CoverTab[121336]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:570
	// _ = "end of CoverTab[121330]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:570
	_go_fuzz_dep_.CoverTab[121331]++

												if ch != eof {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:572
		_go_fuzz_dep_.CoverTab[121337]++
													s.unread()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:573
		// _ = "end of CoverTab[121337]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:574
		_go_fuzz_dep_.CoverTab[121338]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:574
		// _ = "end of CoverTab[121338]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:574
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:574
	// _ = "end of CoverTab[121331]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:574
	_go_fuzz_dep_.CoverTab[121332]++

												return string(s.src[offs:s.srcPos.Offset])
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:576
	// _ = "end of CoverTab[121332]"
}

// recentPosition returns the position of the character immediately after the
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:579
// character or token returned by the last call to Scan.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:581
func (s *Scanner) recentPosition() (pos token.Pos) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:581
	_go_fuzz_dep_.CoverTab[121339]++
												pos.Offset = s.srcPos.Offset - s.lastCharLen
												switch {
	case s.srcPos.Column > 0:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:584
		_go_fuzz_dep_.CoverTab[121341]++

													pos.Line = s.srcPos.Line
													pos.Column = s.srcPos.Column
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:587
		// _ = "end of CoverTab[121341]"
	case s.lastLineLen > 0:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:588
		_go_fuzz_dep_.CoverTab[121342]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:592
		pos.Line = s.srcPos.Line - 1
													pos.Column = s.lastLineLen
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:593
		// _ = "end of CoverTab[121342]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:594
		_go_fuzz_dep_.CoverTab[121343]++

													pos.Line = 1
													pos.Column = 1
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:597
		// _ = "end of CoverTab[121343]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:598
	// _ = "end of CoverTab[121339]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:598
	_go_fuzz_dep_.CoverTab[121340]++
												return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:599
	// _ = "end of CoverTab[121340]"
}

// err prints the error of any scanning to s.Error function. If the function is
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:602
// not defined, by default it prints them to os.Stderr
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:604
func (s *Scanner) err(msg string) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:604
	_go_fuzz_dep_.CoverTab[121344]++
												s.ErrorCount++
												pos := s.recentPosition()

												if s.Error != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:608
		_go_fuzz_dep_.CoverTab[121346]++
													s.Error(pos, msg)
													return
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:610
		// _ = "end of CoverTab[121346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:611
		_go_fuzz_dep_.CoverTab[121347]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:611
		// _ = "end of CoverTab[121347]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:611
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:611
	// _ = "end of CoverTab[121344]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:611
	_go_fuzz_dep_.CoverTab[121345]++

												fmt.Fprintf(os.Stderr, "%s: %s\n", pos, msg)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:613
	// _ = "end of CoverTab[121345]"
}

// isHexadecimal returns true if the given rune is a letter
func isLetter(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:617
	_go_fuzz_dep_.CoverTab[121348]++
												return 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		_go_fuzz_dep_.CoverTab[121349]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		return ch <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		// _ = "end of CoverTab[121349]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		_go_fuzz_dep_.CoverTab[121350]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		return 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			_go_fuzz_dep_.CoverTab[121351]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			return ch <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			// _ = "end of CoverTab[121351]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		// _ = "end of CoverTab[121350]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		_go_fuzz_dep_.CoverTab[121352]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		return ch == '_'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		// _ = "end of CoverTab[121352]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		_go_fuzz_dep_.CoverTab[121353]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		return ch >= 0x80 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			_go_fuzz_dep_.CoverTab[121354]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			return unicode.IsLetter(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
			// _ = "end of CoverTab[121354]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
		// _ = "end of CoverTab[121353]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:618
	// _ = "end of CoverTab[121348]"
}

// isDigit returns true if the given rune is a decimal digit
func isDigit(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:622
	_go_fuzz_dep_.CoverTab[121355]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		_go_fuzz_dep_.CoverTab[121356]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		// _ = "end of CoverTab[121356]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		_go_fuzz_dep_.CoverTab[121357]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		return ch >= 0x80 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
			_go_fuzz_dep_.CoverTab[121358]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
			return unicode.IsDigit(ch)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
			// _ = "end of CoverTab[121358]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
		// _ = "end of CoverTab[121357]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:623
	// _ = "end of CoverTab[121355]"
}

// isDecimal returns true if the given rune is a decimal number
func isDecimal(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:627
	_go_fuzz_dep_.CoverTab[121359]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:628
		_go_fuzz_dep_.CoverTab[121360]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:628
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:628
		// _ = "end of CoverTab[121360]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:628
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:628
	// _ = "end of CoverTab[121359]"
}

// isHexadecimal returns true if the given rune is an hexadecimal number
func isHexadecimal(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:632
	_go_fuzz_dep_.CoverTab[121361]++
												return '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		_go_fuzz_dep_.CoverTab[121362]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		// _ = "end of CoverTab[121362]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		_go_fuzz_dep_.CoverTab[121363]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		return 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			_go_fuzz_dep_.CoverTab[121364]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			return ch <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			// _ = "end of CoverTab[121364]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		// _ = "end of CoverTab[121363]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		_go_fuzz_dep_.CoverTab[121365]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		return 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			_go_fuzz_dep_.CoverTab[121366]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			return ch <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
			// _ = "end of CoverTab[121366]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
		// _ = "end of CoverTab[121365]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:633
	// _ = "end of CoverTab[121361]"
}

// isWhitespace returns true if the rune is a space, tab, newline or carriage return
func isWhitespace(ch rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:637
	_go_fuzz_dep_.CoverTab[121367]++
												return ch == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		_go_fuzz_dep_.CoverTab[121368]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		return ch == '\t'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		// _ = "end of CoverTab[121368]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		_go_fuzz_dep_.CoverTab[121369]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		return ch == '\n'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		// _ = "end of CoverTab[121369]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		_go_fuzz_dep_.CoverTab[121370]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		return ch == '\r'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
		// _ = "end of CoverTab[121370]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
	}()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:638
	// _ = "end of CoverTab[121367]"
}

// digitVal returns the integer value of a given octal,decimal or hexadecimal rune
func digitVal(ch rune) int {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:642
	_go_fuzz_dep_.CoverTab[121371]++
												switch {
	case '0' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:644
		_go_fuzz_dep_.CoverTab[121377]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:644
		return ch <= '9'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:644
		// _ = "end of CoverTab[121377]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:644
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:644
		_go_fuzz_dep_.CoverTab[121373]++
													return int(ch - '0')
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:645
		// _ = "end of CoverTab[121373]"
	case 'a' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:646
		_go_fuzz_dep_.CoverTab[121378]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:646
		return ch <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:646
		// _ = "end of CoverTab[121378]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:646
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:646
		_go_fuzz_dep_.CoverTab[121374]++
													return int(ch - 'a' + 10)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:647
		// _ = "end of CoverTab[121374]"
	case 'A' <= ch && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:648
		_go_fuzz_dep_.CoverTab[121379]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:648
		return ch <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:648
		// _ = "end of CoverTab[121379]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:648
	}():
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:648
		_go_fuzz_dep_.CoverTab[121375]++
													return int(ch - 'A' + 10)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:649
		// _ = "end of CoverTab[121375]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:649
	default:
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:649
		_go_fuzz_dep_.CoverTab[121376]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:649
		// _ = "end of CoverTab[121376]"
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:650
	// _ = "end of CoverTab[121371]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:650
	_go_fuzz_dep_.CoverTab[121372]++
												return 16
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:651
	// _ = "end of CoverTab[121372]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:652
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/scanner/scanner.go:652
var _ = _go_fuzz_dep_.CoverTab
