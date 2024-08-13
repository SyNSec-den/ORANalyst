//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1
)

import (
	"bytes"
	"fmt"
)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:483
// Ensure that the buffer contains the required number of characters.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:483
// Return true on success, false on failure (reader error or memory error).
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:485
func cache(parser *yaml_parser_t, length int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:485
	_go_fuzz_dep_.CoverTab[126736]++

										return parser.unread >= length || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:487
		_go_fuzz_dep_.CoverTab[126737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:487
		return yaml_parser_update_buffer(parser, length)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:487
		// _ = "end of CoverTab[126737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:487
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:487
	// _ = "end of CoverTab[126736]"
}

// Advance the buffer pointer.
func skip(parser *yaml_parser_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:491
	_go_fuzz_dep_.CoverTab[126738]++
										parser.mark.index++
										parser.mark.column++
										parser.unread--
										parser.buffer_pos += width(parser.buffer[parser.buffer_pos])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:495
	// _ = "end of CoverTab[126738]"
}

func skip_line(parser *yaml_parser_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:498
	_go_fuzz_dep_.CoverTab[126739]++
										if is_crlf(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:499
		_go_fuzz_dep_.CoverTab[126740]++
											parser.mark.index += 2
											parser.mark.column = 0
											parser.mark.line++
											parser.unread -= 2
											parser.buffer_pos += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:504
		// _ = "end of CoverTab[126740]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:505
		_go_fuzz_dep_.CoverTab[126741]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:505
		if is_break(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:505
			_go_fuzz_dep_.CoverTab[126742]++
												parser.mark.index++
												parser.mark.column = 0
												parser.mark.line++
												parser.unread--
												parser.buffer_pos += width(parser.buffer[parser.buffer_pos])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:510
			// _ = "end of CoverTab[126742]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
			_go_fuzz_dep_.CoverTab[126743]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
			// _ = "end of CoverTab[126743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
		// _ = "end of CoverTab[126741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:511
	// _ = "end of CoverTab[126739]"
}

// Copy a character to a string buffer and advance pointers.
func read(parser *yaml_parser_t, s []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:515
	_go_fuzz_dep_.CoverTab[126744]++
										w := width(parser.buffer[parser.buffer_pos])
										if w == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:517
		_go_fuzz_dep_.CoverTab[126748]++
											panic("invalid character sequence")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:518
		// _ = "end of CoverTab[126748]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:519
		_go_fuzz_dep_.CoverTab[126749]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:519
		// _ = "end of CoverTab[126749]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:519
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:519
	// _ = "end of CoverTab[126744]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:519
	_go_fuzz_dep_.CoverTab[126745]++
										if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:520
		_go_fuzz_dep_.CoverTab[126750]++
											s = make([]byte, 0, 32)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:521
		// _ = "end of CoverTab[126750]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:522
		_go_fuzz_dep_.CoverTab[126751]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:522
		// _ = "end of CoverTab[126751]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:522
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:522
	// _ = "end of CoverTab[126745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:522
	_go_fuzz_dep_.CoverTab[126746]++
										if w == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:523
		_go_fuzz_dep_.CoverTab[126752]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:523
		return len(s)+w <= cap(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:523
		// _ = "end of CoverTab[126752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:523
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:523
		_go_fuzz_dep_.CoverTab[126753]++
											s = s[:len(s)+1]
											s[len(s)-1] = parser.buffer[parser.buffer_pos]
											parser.buffer_pos++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:526
		// _ = "end of CoverTab[126753]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:527
		_go_fuzz_dep_.CoverTab[126754]++
											s = append(s, parser.buffer[parser.buffer_pos:parser.buffer_pos+w]...)
											parser.buffer_pos += w
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:529
		// _ = "end of CoverTab[126754]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:530
	// _ = "end of CoverTab[126746]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:530
	_go_fuzz_dep_.CoverTab[126747]++
										parser.mark.index++
										parser.mark.column++
										parser.unread--
										return s
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:534
	// _ = "end of CoverTab[126747]"
}

// Copy a line break character to a string buffer and advance pointers.
func read_line(parser *yaml_parser_t, s []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:538
	_go_fuzz_dep_.CoverTab[126755]++
										buf := parser.buffer
										pos := parser.buffer_pos
										switch {
	case buf[pos] == '\r' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:542
		_go_fuzz_dep_.CoverTab[126762]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:542
		return buf[pos+1] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:542
		// _ = "end of CoverTab[126762]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:542
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:542
		_go_fuzz_dep_.CoverTab[126757]++

											s = append(s, '\n')
											parser.buffer_pos += 2
											parser.mark.index++
											parser.unread--
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:547
		// _ = "end of CoverTab[126757]"
	case buf[pos] == '\r' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:548
		_go_fuzz_dep_.CoverTab[126763]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:548
		return buf[pos] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:548
		// _ = "end of CoverTab[126763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:548
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:548
		_go_fuzz_dep_.CoverTab[126758]++

											s = append(s, '\n')
											parser.buffer_pos += 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:551
		// _ = "end of CoverTab[126758]"
	case buf[pos] == '\xC2' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:552
		_go_fuzz_dep_.CoverTab[126764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:552
		return buf[pos+1] == '\x85'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:552
		// _ = "end of CoverTab[126764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:552
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:552
		_go_fuzz_dep_.CoverTab[126759]++

											s = append(s, '\n')
											parser.buffer_pos += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:555
		// _ = "end of CoverTab[126759]"
	case buf[pos] == '\xE2' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		_go_fuzz_dep_.CoverTab[126765]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		return buf[pos+1] == '\x80'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		// _ = "end of CoverTab[126765]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		_go_fuzz_dep_.CoverTab[126766]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		return (buf[pos+2] == '\xA8' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
			_go_fuzz_dep_.CoverTab[126767]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
			return buf[pos+2] == '\xA9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
			// _ = "end of CoverTab[126767]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		// _ = "end of CoverTab[126766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
	}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:556
		_go_fuzz_dep_.CoverTab[126760]++

											s = append(s, buf[parser.buffer_pos:pos+3]...)
											parser.buffer_pos += 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:559
		// _ = "end of CoverTab[126760]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:560
		_go_fuzz_dep_.CoverTab[126761]++
											return s
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:561
		// _ = "end of CoverTab[126761]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:562
	// _ = "end of CoverTab[126755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:562
	_go_fuzz_dep_.CoverTab[126756]++
										parser.mark.index++
										parser.mark.column = 0
										parser.mark.line++
										parser.unread--
										return s
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:567
	// _ = "end of CoverTab[126756]"
}

// Get the next token.
func yaml_parser_scan(parser *yaml_parser_t, token *yaml_token_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:571
	_go_fuzz_dep_.CoverTab[126768]++

										*token = yaml_token_t{}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
	if parser.stream_end_produced || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
		_go_fuzz_dep_.CoverTab[126772]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
		return parser.error != yaml_NO_ERROR
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
		// _ = "end of CoverTab[126772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:576
		_go_fuzz_dep_.CoverTab[126773]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:577
		// _ = "end of CoverTab[126773]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:578
		_go_fuzz_dep_.CoverTab[126774]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:578
		// _ = "end of CoverTab[126774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:578
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:578
	// _ = "end of CoverTab[126768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:578
	_go_fuzz_dep_.CoverTab[126769]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:581
	if !parser.token_available {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:581
		_go_fuzz_dep_.CoverTab[126775]++
											if !yaml_parser_fetch_more_tokens(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:582
			_go_fuzz_dep_.CoverTab[126776]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:583
			// _ = "end of CoverTab[126776]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:584
			_go_fuzz_dep_.CoverTab[126777]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:584
			// _ = "end of CoverTab[126777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:584
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:584
		// _ = "end of CoverTab[126775]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:585
		_go_fuzz_dep_.CoverTab[126778]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:585
		// _ = "end of CoverTab[126778]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:585
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:585
	// _ = "end of CoverTab[126769]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:585
	_go_fuzz_dep_.CoverTab[126770]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:588
	*token = parser.tokens[parser.tokens_head]
	parser.tokens_head++
	parser.tokens_parsed++
	parser.token_available = false

	if token.typ == yaml_STREAM_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:593
		_go_fuzz_dep_.CoverTab[126779]++
											parser.stream_end_produced = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:594
		// _ = "end of CoverTab[126779]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:595
		_go_fuzz_dep_.CoverTab[126780]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:595
		// _ = "end of CoverTab[126780]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:595
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:595
	// _ = "end of CoverTab[126770]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:595
	_go_fuzz_dep_.CoverTab[126771]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:596
	// _ = "end of CoverTab[126771]"
}

// Set the scanner error and return false.
func yaml_parser_set_scanner_error(parser *yaml_parser_t, context string, context_mark yaml_mark_t, problem string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:600
	_go_fuzz_dep_.CoverTab[126781]++
										parser.error = yaml_SCANNER_ERROR
										parser.context = context
										parser.context_mark = context_mark
										parser.problem = problem
										parser.problem_mark = parser.mark
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:606
	// _ = "end of CoverTab[126781]"
}

func yaml_parser_set_scanner_tag_error(parser *yaml_parser_t, directive bool, context_mark yaml_mark_t, problem string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:609
	_go_fuzz_dep_.CoverTab[126782]++
										context := "while parsing a tag"
										if directive {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:611
		_go_fuzz_dep_.CoverTab[126784]++
											context = "while parsing a %TAG directive"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:612
		// _ = "end of CoverTab[126784]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:613
		_go_fuzz_dep_.CoverTab[126785]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:613
		// _ = "end of CoverTab[126785]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:613
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:613
	// _ = "end of CoverTab[126782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:613
	_go_fuzz_dep_.CoverTab[126783]++
										return yaml_parser_set_scanner_error(parser, context, context_mark, problem)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:614
	// _ = "end of CoverTab[126783]"
}

func trace(args ...interface{}) func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:617
	_go_fuzz_dep_.CoverTab[126786]++
										pargs := append([]interface{}{"+++"}, args...)
										fmt.Println(pargs...)
										pargs = append([]interface{}{"---"}, args...)
										return func() { _go_fuzz_dep_.CoverTab[126787]++; fmt.Println(pargs...); // _ = "end of CoverTab[126787]" }
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:621
	// _ = "end of CoverTab[126786]"
}

// Ensure that the tokens queue contains at least one token which can be
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:624
// returned to the Parser.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:626
func yaml_parser_fetch_more_tokens(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:626
	_go_fuzz_dep_.CoverTab[126788]++

										for {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:628
		_go_fuzz_dep_.CoverTab[126790]++
											if parser.tokens_head != len(parser.tokens) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:629
			_go_fuzz_dep_.CoverTab[126792]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:632
			head_tok_idx, ok := parser.simple_keys_by_tok[parser.tokens_parsed]
			if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:633
				_go_fuzz_dep_.CoverTab[126793]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:634
				// _ = "end of CoverTab[126793]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:635
				_go_fuzz_dep_.CoverTab[126794]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:635
				if valid, ok := yaml_simple_key_is_valid(parser, &parser.simple_keys[head_tok_idx]); !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:635
					_go_fuzz_dep_.CoverTab[126795]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:636
					// _ = "end of CoverTab[126795]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:637
					_go_fuzz_dep_.CoverTab[126796]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:637
					if !valid {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:637
						_go_fuzz_dep_.CoverTab[126797]++
															break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:638
						// _ = "end of CoverTab[126797]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
						_go_fuzz_dep_.CoverTab[126798]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
						// _ = "end of CoverTab[126798]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
					// _ = "end of CoverTab[126796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
				// _ = "end of CoverTab[126794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:639
			// _ = "end of CoverTab[126792]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:640
			_go_fuzz_dep_.CoverTab[126799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:640
			// _ = "end of CoverTab[126799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:640
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:640
		// _ = "end of CoverTab[126790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:640
		_go_fuzz_dep_.CoverTab[126791]++

											if !yaml_parser_fetch_next_token(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:642
			_go_fuzz_dep_.CoverTab[126800]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:643
			// _ = "end of CoverTab[126800]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:644
			_go_fuzz_dep_.CoverTab[126801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:644
			// _ = "end of CoverTab[126801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:644
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:644
		// _ = "end of CoverTab[126791]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:645
	// _ = "end of CoverTab[126788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:645
	_go_fuzz_dep_.CoverTab[126789]++

										parser.token_available = true
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:648
	// _ = "end of CoverTab[126789]"
}

// The dispatcher for token fetchers.
func yaml_parser_fetch_next_token(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:652
	_go_fuzz_dep_.CoverTab[126802]++

										if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:654
		_go_fuzz_dep_.CoverTab[126828]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:654
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:654
		// _ = "end of CoverTab[126828]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:654
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:654
		_go_fuzz_dep_.CoverTab[126829]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:655
		// _ = "end of CoverTab[126829]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:656
		_go_fuzz_dep_.CoverTab[126830]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:656
		// _ = "end of CoverTab[126830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:656
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:656
	// _ = "end of CoverTab[126802]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:656
	_go_fuzz_dep_.CoverTab[126803]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:659
	if !parser.stream_start_produced {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:659
		_go_fuzz_dep_.CoverTab[126831]++
											return yaml_parser_fetch_stream_start(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:660
		// _ = "end of CoverTab[126831]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:661
		_go_fuzz_dep_.CoverTab[126832]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:661
		// _ = "end of CoverTab[126832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:661
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:661
	// _ = "end of CoverTab[126803]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:661
	_go_fuzz_dep_.CoverTab[126804]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:664
	if !yaml_parser_scan_to_next_token(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:664
		_go_fuzz_dep_.CoverTab[126833]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:665
		// _ = "end of CoverTab[126833]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:666
		_go_fuzz_dep_.CoverTab[126834]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:666
		// _ = "end of CoverTab[126834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:666
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:666
	// _ = "end of CoverTab[126804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:666
	_go_fuzz_dep_.CoverTab[126805]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:669
	if !yaml_parser_unroll_indent(parser, parser.mark.column) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:669
		_go_fuzz_dep_.CoverTab[126835]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:670
		// _ = "end of CoverTab[126835]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:671
		_go_fuzz_dep_.CoverTab[126836]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:671
		// _ = "end of CoverTab[126836]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:671
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:671
	// _ = "end of CoverTab[126805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:671
	_go_fuzz_dep_.CoverTab[126806]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
	if parser.unread < 4 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
		_go_fuzz_dep_.CoverTab[126837]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
		return !yaml_parser_update_buffer(parser, 4)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
		// _ = "end of CoverTab[126837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:675
		_go_fuzz_dep_.CoverTab[126838]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:676
		// _ = "end of CoverTab[126838]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:677
		_go_fuzz_dep_.CoverTab[126839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:677
		// _ = "end of CoverTab[126839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:677
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:677
	// _ = "end of CoverTab[126806]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:677
	_go_fuzz_dep_.CoverTab[126807]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:680
	if is_z(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:680
		_go_fuzz_dep_.CoverTab[126840]++
											return yaml_parser_fetch_stream_end(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:681
		// _ = "end of CoverTab[126840]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:682
		_go_fuzz_dep_.CoverTab[126841]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:682
		// _ = "end of CoverTab[126841]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:682
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:682
	// _ = "end of CoverTab[126807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:682
	_go_fuzz_dep_.CoverTab[126808]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
	if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
		_go_fuzz_dep_.CoverTab[126842]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
		return parser.buffer[parser.buffer_pos] == '%'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
		// _ = "end of CoverTab[126842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:685
		_go_fuzz_dep_.CoverTab[126843]++
											return yaml_parser_fetch_directive(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:686
		// _ = "end of CoverTab[126843]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:687
		_go_fuzz_dep_.CoverTab[126844]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:687
		// _ = "end of CoverTab[126844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:687
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:687
	// _ = "end of CoverTab[126808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:687
	_go_fuzz_dep_.CoverTab[126809]++

										buf := parser.buffer
										pos := parser.buffer_pos

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
	if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		_go_fuzz_dep_.CoverTab[126845]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		return buf[pos] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		// _ = "end of CoverTab[126845]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		_go_fuzz_dep_.CoverTab[126846]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		return buf[pos+1] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		// _ = "end of CoverTab[126846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		_go_fuzz_dep_.CoverTab[126847]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		return buf[pos+2] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		// _ = "end of CoverTab[126847]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		_go_fuzz_dep_.CoverTab[126848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		return is_blankz(buf, pos+3)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		// _ = "end of CoverTab[126848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:693
		_go_fuzz_dep_.CoverTab[126849]++
											return yaml_parser_fetch_document_indicator(parser, yaml_DOCUMENT_START_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:694
		// _ = "end of CoverTab[126849]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:695
		_go_fuzz_dep_.CoverTab[126850]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:695
		// _ = "end of CoverTab[126850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:695
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:695
	// _ = "end of CoverTab[126809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:695
	_go_fuzz_dep_.CoverTab[126810]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
	if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		_go_fuzz_dep_.CoverTab[126851]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		return buf[pos] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		// _ = "end of CoverTab[126851]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		_go_fuzz_dep_.CoverTab[126852]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		return buf[pos+1] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		// _ = "end of CoverTab[126852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		_go_fuzz_dep_.CoverTab[126853]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		return buf[pos+2] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		// _ = "end of CoverTab[126853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		_go_fuzz_dep_.CoverTab[126854]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		return is_blankz(buf, pos+3)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		// _ = "end of CoverTab[126854]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:698
		_go_fuzz_dep_.CoverTab[126855]++
											return yaml_parser_fetch_document_indicator(parser, yaml_DOCUMENT_END_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:699
		// _ = "end of CoverTab[126855]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:700
		_go_fuzz_dep_.CoverTab[126856]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:700
		// _ = "end of CoverTab[126856]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:700
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:700
	// _ = "end of CoverTab[126810]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:700
	_go_fuzz_dep_.CoverTab[126811]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:703
	if buf[pos] == '[' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:703
		_go_fuzz_dep_.CoverTab[126857]++
											return yaml_parser_fetch_flow_collection_start(parser, yaml_FLOW_SEQUENCE_START_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:704
		// _ = "end of CoverTab[126857]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:705
		_go_fuzz_dep_.CoverTab[126858]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:705
		// _ = "end of CoverTab[126858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:705
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:705
	// _ = "end of CoverTab[126811]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:705
	_go_fuzz_dep_.CoverTab[126812]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:708
	if parser.buffer[parser.buffer_pos] == '{' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:708
		_go_fuzz_dep_.CoverTab[126859]++
											return yaml_parser_fetch_flow_collection_start(parser, yaml_FLOW_MAPPING_START_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:709
		// _ = "end of CoverTab[126859]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:710
		_go_fuzz_dep_.CoverTab[126860]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:710
		// _ = "end of CoverTab[126860]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:710
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:710
	// _ = "end of CoverTab[126812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:710
	_go_fuzz_dep_.CoverTab[126813]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:713
	if parser.buffer[parser.buffer_pos] == ']' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:713
		_go_fuzz_dep_.CoverTab[126861]++
											return yaml_parser_fetch_flow_collection_end(parser,
			yaml_FLOW_SEQUENCE_END_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:715
		// _ = "end of CoverTab[126861]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:716
		_go_fuzz_dep_.CoverTab[126862]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:716
		// _ = "end of CoverTab[126862]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:716
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:716
	// _ = "end of CoverTab[126813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:716
	_go_fuzz_dep_.CoverTab[126814]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:719
	if parser.buffer[parser.buffer_pos] == '}' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:719
		_go_fuzz_dep_.CoverTab[126863]++
											return yaml_parser_fetch_flow_collection_end(parser,
			yaml_FLOW_MAPPING_END_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:721
		// _ = "end of CoverTab[126863]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:722
		_go_fuzz_dep_.CoverTab[126864]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:722
		// _ = "end of CoverTab[126864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:722
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:722
	// _ = "end of CoverTab[126814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:722
	_go_fuzz_dep_.CoverTab[126815]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:725
	if parser.buffer[parser.buffer_pos] == ',' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:725
		_go_fuzz_dep_.CoverTab[126865]++
											return yaml_parser_fetch_flow_entry(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:726
		// _ = "end of CoverTab[126865]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:727
		_go_fuzz_dep_.CoverTab[126866]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:727
		// _ = "end of CoverTab[126866]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:727
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:727
	// _ = "end of CoverTab[126815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:727
	_go_fuzz_dep_.CoverTab[126816]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
	if parser.buffer[parser.buffer_pos] == '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
		_go_fuzz_dep_.CoverTab[126867]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
		return is_blankz(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
		// _ = "end of CoverTab[126867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:730
		_go_fuzz_dep_.CoverTab[126868]++
											return yaml_parser_fetch_block_entry(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:731
		// _ = "end of CoverTab[126868]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:732
		_go_fuzz_dep_.CoverTab[126869]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:732
		// _ = "end of CoverTab[126869]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:732
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:732
	// _ = "end of CoverTab[126816]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:732
	_go_fuzz_dep_.CoverTab[126817]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
	if parser.buffer[parser.buffer_pos] == '?' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
		_go_fuzz_dep_.CoverTab[126870]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
		return (parser.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
			_go_fuzz_dep_.CoverTab[126871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
			return is_blankz(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
			// _ = "end of CoverTab[126871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
		// _ = "end of CoverTab[126870]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:735
		_go_fuzz_dep_.CoverTab[126872]++
											return yaml_parser_fetch_key(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:736
		// _ = "end of CoverTab[126872]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:737
		_go_fuzz_dep_.CoverTab[126873]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:737
		// _ = "end of CoverTab[126873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:737
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:737
	// _ = "end of CoverTab[126817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:737
	_go_fuzz_dep_.CoverTab[126818]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
	if parser.buffer[parser.buffer_pos] == ':' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
		_go_fuzz_dep_.CoverTab[126874]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
		return (parser.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
			_go_fuzz_dep_.CoverTab[126875]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
			return is_blankz(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
			// _ = "end of CoverTab[126875]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
		// _ = "end of CoverTab[126874]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:740
		_go_fuzz_dep_.CoverTab[126876]++
											return yaml_parser_fetch_value(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:741
		// _ = "end of CoverTab[126876]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:742
		_go_fuzz_dep_.CoverTab[126877]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:742
		// _ = "end of CoverTab[126877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:742
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:742
	// _ = "end of CoverTab[126818]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:742
	_go_fuzz_dep_.CoverTab[126819]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:745
	if parser.buffer[parser.buffer_pos] == '*' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:745
		_go_fuzz_dep_.CoverTab[126878]++
											return yaml_parser_fetch_anchor(parser, yaml_ALIAS_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:746
		// _ = "end of CoverTab[126878]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:747
		_go_fuzz_dep_.CoverTab[126879]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:747
		// _ = "end of CoverTab[126879]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:747
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:747
	// _ = "end of CoverTab[126819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:747
	_go_fuzz_dep_.CoverTab[126820]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:750
	if parser.buffer[parser.buffer_pos] == '&' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:750
		_go_fuzz_dep_.CoverTab[126880]++
											return yaml_parser_fetch_anchor(parser, yaml_ANCHOR_TOKEN)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:751
		// _ = "end of CoverTab[126880]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:752
		_go_fuzz_dep_.CoverTab[126881]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:752
		// _ = "end of CoverTab[126881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:752
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:752
	// _ = "end of CoverTab[126820]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:752
	_go_fuzz_dep_.CoverTab[126821]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:755
	if parser.buffer[parser.buffer_pos] == '!' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:755
		_go_fuzz_dep_.CoverTab[126882]++
											return yaml_parser_fetch_tag(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:756
		// _ = "end of CoverTab[126882]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:757
		_go_fuzz_dep_.CoverTab[126883]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:757
		// _ = "end of CoverTab[126883]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:757
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:757
	// _ = "end of CoverTab[126821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:757
	_go_fuzz_dep_.CoverTab[126822]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
	if parser.buffer[parser.buffer_pos] == '|' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
		_go_fuzz_dep_.CoverTab[126884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
		return parser.flow_level == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
		// _ = "end of CoverTab[126884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:760
		_go_fuzz_dep_.CoverTab[126885]++
											return yaml_parser_fetch_block_scalar(parser, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:761
		// _ = "end of CoverTab[126885]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:762
		_go_fuzz_dep_.CoverTab[126886]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:762
		// _ = "end of CoverTab[126886]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:762
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:762
	// _ = "end of CoverTab[126822]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:762
	_go_fuzz_dep_.CoverTab[126823]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
	if parser.buffer[parser.buffer_pos] == '>' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
		_go_fuzz_dep_.CoverTab[126887]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
		return parser.flow_level == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
		// _ = "end of CoverTab[126887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:765
		_go_fuzz_dep_.CoverTab[126888]++
											return yaml_parser_fetch_block_scalar(parser, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:766
		// _ = "end of CoverTab[126888]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:767
		_go_fuzz_dep_.CoverTab[126889]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:767
		// _ = "end of CoverTab[126889]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:767
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:767
	// _ = "end of CoverTab[126823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:767
	_go_fuzz_dep_.CoverTab[126824]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:770
	if parser.buffer[parser.buffer_pos] == '\'' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:770
		_go_fuzz_dep_.CoverTab[126890]++
											return yaml_parser_fetch_flow_scalar(parser, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:771
		// _ = "end of CoverTab[126890]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:772
		_go_fuzz_dep_.CoverTab[126891]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:772
		// _ = "end of CoverTab[126891]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:772
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:772
	// _ = "end of CoverTab[126824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:772
	_go_fuzz_dep_.CoverTab[126825]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:775
	if parser.buffer[parser.buffer_pos] == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:775
		_go_fuzz_dep_.CoverTab[126892]++
											return yaml_parser_fetch_flow_scalar(parser, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:776
		// _ = "end of CoverTab[126892]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:777
		_go_fuzz_dep_.CoverTab[126893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:777
		// _ = "end of CoverTab[126893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:777
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:777
	// _ = "end of CoverTab[126825]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:777
	_go_fuzz_dep_.CoverTab[126826]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
	if !(is_blankz(parser.buffer, parser.buffer_pos) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
		_go_fuzz_dep_.CoverTab[126894]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
		return parser.buffer[parser.buffer_pos] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
		// _ = "end of CoverTab[126894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
		_go_fuzz_dep_.CoverTab[126895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:799
		return parser.buffer[parser.buffer_pos] == '?'
											// _ = "end of CoverTab[126895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
		_go_fuzz_dep_.CoverTab[126896]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
		return parser.buffer[parser.buffer_pos] == ':'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
		// _ = "end of CoverTab[126896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
		_go_fuzz_dep_.CoverTab[126897]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:800
		return parser.buffer[parser.buffer_pos] == ','
											// _ = "end of CoverTab[126897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
		_go_fuzz_dep_.CoverTab[126898]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
		return parser.buffer[parser.buffer_pos] == '['
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
		// _ = "end of CoverTab[126898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
		_go_fuzz_dep_.CoverTab[126899]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:801
		return parser.buffer[parser.buffer_pos] == ']'
											// _ = "end of CoverTab[126899]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
		_go_fuzz_dep_.CoverTab[126900]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
		return parser.buffer[parser.buffer_pos] == '{'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
		// _ = "end of CoverTab[126900]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
		_go_fuzz_dep_.CoverTab[126901]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:802
		return parser.buffer[parser.buffer_pos] == '}'
											// _ = "end of CoverTab[126901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
		_go_fuzz_dep_.CoverTab[126902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
		return parser.buffer[parser.buffer_pos] == '#'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
		// _ = "end of CoverTab[126902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
		_go_fuzz_dep_.CoverTab[126903]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:803
		return parser.buffer[parser.buffer_pos] == '&'
											// _ = "end of CoverTab[126903]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
		_go_fuzz_dep_.CoverTab[126904]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
		return parser.buffer[parser.buffer_pos] == '*'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
		// _ = "end of CoverTab[126904]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
		_go_fuzz_dep_.CoverTab[126905]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:804
		return parser.buffer[parser.buffer_pos] == '!'
											// _ = "end of CoverTab[126905]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
		_go_fuzz_dep_.CoverTab[126906]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
		return parser.buffer[parser.buffer_pos] == '|'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
		// _ = "end of CoverTab[126906]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
		_go_fuzz_dep_.CoverTab[126907]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:805
		return parser.buffer[parser.buffer_pos] == '>'
											// _ = "end of CoverTab[126907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
		_go_fuzz_dep_.CoverTab[126908]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
		return parser.buffer[parser.buffer_pos] == '\''
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
		// _ = "end of CoverTab[126908]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
		_go_fuzz_dep_.CoverTab[126909]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:806
		return parser.buffer[parser.buffer_pos] == '"'
											// _ = "end of CoverTab[126909]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
		_go_fuzz_dep_.CoverTab[126910]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
		return parser.buffer[parser.buffer_pos] == '%'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
		// _ = "end of CoverTab[126910]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
		_go_fuzz_dep_.CoverTab[126911]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:807
		return parser.buffer[parser.buffer_pos] == '@'
											// _ = "end of CoverTab[126911]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
		_go_fuzz_dep_.CoverTab[126912]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
		return parser.buffer[parser.buffer_pos] == '`'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
		// _ = "end of CoverTab[126912]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
		_go_fuzz_dep_.CoverTab[126913]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:808
		return (parser.buffer[parser.buffer_pos] == '-' && func() bool {
												_go_fuzz_dep_.CoverTab[126914]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
			return !is_blank(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
			// _ = "end of CoverTab[126914]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
		// _ = "end of CoverTab[126913]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
		_go_fuzz_dep_.CoverTab[126915]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:809
		return (parser.flow_level == 0 && func() bool {
												_go_fuzz_dep_.CoverTab[126916]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:810
			return (parser.buffer[parser.buffer_pos] == '?' || func() bool {
													_go_fuzz_dep_.CoverTab[126917]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
				return parser.buffer[parser.buffer_pos] == ':'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
				// _ = "end of CoverTab[126917]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
			// _ = "end of CoverTab[126916]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
			_go_fuzz_dep_.CoverTab[126918]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:811
			return !is_blankz(parser.buffer, parser.buffer_pos+1)
												// _ = "end of CoverTab[126918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:812
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:812
		// _ = "end of CoverTab[126915]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:812
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:812
		_go_fuzz_dep_.CoverTab[126919]++
											return yaml_parser_fetch_plain_scalar(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:813
		// _ = "end of CoverTab[126919]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:814
		_go_fuzz_dep_.CoverTab[126920]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:814
		// _ = "end of CoverTab[126920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:814
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:814
	// _ = "end of CoverTab[126826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:814
	_go_fuzz_dep_.CoverTab[126827]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:817
	return yaml_parser_set_scanner_error(parser,
		"while scanning for the next token", parser.mark,
		"found character that cannot start any token")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:819
	// _ = "end of CoverTab[126827]"
}

func yaml_simple_key_is_valid(parser *yaml_parser_t, simple_key *yaml_simple_key_t) (valid, ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:822
	_go_fuzz_dep_.CoverTab[126921]++
										if !simple_key.possible {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:823
		_go_fuzz_dep_.CoverTab[126924]++
											return false, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:824
		// _ = "end of CoverTab[126924]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:825
		_go_fuzz_dep_.CoverTab[126925]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:825
		// _ = "end of CoverTab[126925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:825
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:825
	// _ = "end of CoverTab[126921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:825
	_go_fuzz_dep_.CoverTab[126922]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
	if simple_key.mark.line < parser.mark.line || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
		_go_fuzz_dep_.CoverTab[126926]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
		return simple_key.mark.index+1024 < parser.mark.index
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
		// _ = "end of CoverTab[126926]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:835
		_go_fuzz_dep_.CoverTab[126927]++

											if simple_key.required {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:837
			_go_fuzz_dep_.CoverTab[126929]++
												return false, yaml_parser_set_scanner_error(parser,
				"while scanning a simple key", simple_key.mark,
				"could not find expected ':'")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:840
			// _ = "end of CoverTab[126929]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:841
			_go_fuzz_dep_.CoverTab[126930]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:841
			// _ = "end of CoverTab[126930]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:841
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:841
		// _ = "end of CoverTab[126927]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:841
		_go_fuzz_dep_.CoverTab[126928]++
											simple_key.possible = false
											return false, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:843
		// _ = "end of CoverTab[126928]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:844
		_go_fuzz_dep_.CoverTab[126931]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:844
		// _ = "end of CoverTab[126931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:844
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:844
	// _ = "end of CoverTab[126922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:844
	_go_fuzz_dep_.CoverTab[126923]++
										return true, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:845
	// _ = "end of CoverTab[126923]"
}

// Check if a simple key may start at the current position and add it if
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:848
// needed.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:850
func yaml_parser_save_simple_key(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:850
	_go_fuzz_dep_.CoverTab[126932]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:855
	required := parser.flow_level == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:855
		_go_fuzz_dep_.CoverTab[126934]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:855
		return parser.indent == parser.mark.column
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:855
		// _ = "end of CoverTab[126934]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:855
	}()

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:860
	if parser.simple_key_allowed {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:860
		_go_fuzz_dep_.CoverTab[126935]++
											simple_key := yaml_simple_key_t{
			possible:	true,
			required:	required,
			token_number:	parser.tokens_parsed + (len(parser.tokens) - parser.tokens_head),
			mark:		parser.mark,
		}

		if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:868
			_go_fuzz_dep_.CoverTab[126937]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:869
			// _ = "end of CoverTab[126937]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:870
			_go_fuzz_dep_.CoverTab[126938]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:870
			// _ = "end of CoverTab[126938]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:870
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:870
		// _ = "end of CoverTab[126935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:870
		_go_fuzz_dep_.CoverTab[126936]++
											parser.simple_keys[len(parser.simple_keys)-1] = simple_key
											parser.simple_keys_by_tok[simple_key.token_number] = len(parser.simple_keys) - 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:872
		// _ = "end of CoverTab[126936]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:873
		_go_fuzz_dep_.CoverTab[126939]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:873
		// _ = "end of CoverTab[126939]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:873
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:873
	// _ = "end of CoverTab[126932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:873
	_go_fuzz_dep_.CoverTab[126933]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:874
	// _ = "end of CoverTab[126933]"
}

// Remove a potential simple key at the current flow level.
func yaml_parser_remove_simple_key(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:878
	_go_fuzz_dep_.CoverTab[126940]++
										i := len(parser.simple_keys) - 1
										if parser.simple_keys[i].possible {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:880
		_go_fuzz_dep_.CoverTab[126942]++

											if parser.simple_keys[i].required {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:882
			_go_fuzz_dep_.CoverTab[126944]++
												return yaml_parser_set_scanner_error(parser,
				"while scanning a simple key", parser.simple_keys[i].mark,
				"could not find expected ':'")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:885
			// _ = "end of CoverTab[126944]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:886
			_go_fuzz_dep_.CoverTab[126945]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:886
			// _ = "end of CoverTab[126945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:886
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:886
		// _ = "end of CoverTab[126942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:886
		_go_fuzz_dep_.CoverTab[126943]++

											parser.simple_keys[i].possible = false
											delete(parser.simple_keys_by_tok, parser.simple_keys[i].token_number)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:889
		// _ = "end of CoverTab[126943]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:890
		_go_fuzz_dep_.CoverTab[126946]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:890
		// _ = "end of CoverTab[126946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:890
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:890
	// _ = "end of CoverTab[126940]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:890
	_go_fuzz_dep_.CoverTab[126941]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:891
	// _ = "end of CoverTab[126941]"
}

// max_flow_level limits the flow_level
const max_flow_level = 10000

// Increase the flow level and resize the simple key list if needed.
func yaml_parser_increase_flow_level(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:898
	_go_fuzz_dep_.CoverTab[126947]++

										parser.simple_keys = append(parser.simple_keys, yaml_simple_key_t{
		possible:	false,
		required:	false,
		token_number:	parser.tokens_parsed + (len(parser.tokens) - parser.tokens_head),
		mark:		parser.mark,
	})

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:908
	parser.flow_level++
	if parser.flow_level > max_flow_level {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:909
		_go_fuzz_dep_.CoverTab[126949]++
											return yaml_parser_set_scanner_error(parser,
			"while increasing flow level", parser.simple_keys[len(parser.simple_keys)-1].mark,
			fmt.Sprintf("exceeded max depth of %d", max_flow_level))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:912
		// _ = "end of CoverTab[126949]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:913
		_go_fuzz_dep_.CoverTab[126950]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:913
		// _ = "end of CoverTab[126950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:913
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:913
	// _ = "end of CoverTab[126947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:913
	_go_fuzz_dep_.CoverTab[126948]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:914
	// _ = "end of CoverTab[126948]"
}

// Decrease the flow level.
func yaml_parser_decrease_flow_level(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:918
	_go_fuzz_dep_.CoverTab[126951]++
										if parser.flow_level > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:919
		_go_fuzz_dep_.CoverTab[126953]++
											parser.flow_level--
											last := len(parser.simple_keys) - 1
											delete(parser.simple_keys_by_tok, parser.simple_keys[last].token_number)
											parser.simple_keys = parser.simple_keys[:last]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:923
		// _ = "end of CoverTab[126953]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:924
		_go_fuzz_dep_.CoverTab[126954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:924
		// _ = "end of CoverTab[126954]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:924
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:924
	// _ = "end of CoverTab[126951]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:924
	_go_fuzz_dep_.CoverTab[126952]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:925
	// _ = "end of CoverTab[126952]"
}

// max_indents limits the indents stack size
const max_indents = 10000

// Push the current indentation level to the stack and set the new level
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:931
// the current column is greater than the indentation level.  In this case,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:931
// append or insert the specified token into the token queue.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:934
func yaml_parser_roll_indent(parser *yaml_parser_t, column, number int, typ yaml_token_type_t, mark yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:934
	_go_fuzz_dep_.CoverTab[126955]++

										if parser.flow_level > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:936
		_go_fuzz_dep_.CoverTab[126958]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:937
		// _ = "end of CoverTab[126958]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:938
		_go_fuzz_dep_.CoverTab[126959]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:938
		// _ = "end of CoverTab[126959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:938
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:938
	// _ = "end of CoverTab[126955]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:938
	_go_fuzz_dep_.CoverTab[126956]++

										if parser.indent < column {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:940
		_go_fuzz_dep_.CoverTab[126960]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:943
		parser.indents = append(parser.indents, parser.indent)
		parser.indent = column
		if len(parser.indents) > max_indents {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:945
			_go_fuzz_dep_.CoverTab[126963]++
												return yaml_parser_set_scanner_error(parser,
				"while increasing indent level", parser.simple_keys[len(parser.simple_keys)-1].mark,
				fmt.Sprintf("exceeded max depth of %d", max_indents))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:948
			// _ = "end of CoverTab[126963]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:949
			_go_fuzz_dep_.CoverTab[126964]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:949
			// _ = "end of CoverTab[126964]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:949
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:949
		// _ = "end of CoverTab[126960]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:949
		_go_fuzz_dep_.CoverTab[126961]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:952
		token := yaml_token_t{
			typ:		typ,
			start_mark:	mark,
			end_mark:	mark,
		}
		if number > -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:957
			_go_fuzz_dep_.CoverTab[126965]++
												number -= parser.tokens_parsed
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:958
			// _ = "end of CoverTab[126965]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:959
			_go_fuzz_dep_.CoverTab[126966]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:959
			// _ = "end of CoverTab[126966]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:959
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:959
		// _ = "end of CoverTab[126961]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:959
		_go_fuzz_dep_.CoverTab[126962]++
											yaml_insert_token(parser, number, &token)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:960
		// _ = "end of CoverTab[126962]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:961
		_go_fuzz_dep_.CoverTab[126967]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:961
		// _ = "end of CoverTab[126967]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:961
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:961
	// _ = "end of CoverTab[126956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:961
	_go_fuzz_dep_.CoverTab[126957]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:962
	// _ = "end of CoverTab[126957]"
}

// Pop indentation levels from the indents stack until the current level
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:965
// becomes less or equal to the column.  For each indentation level, append
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:965
// the BLOCK-END token.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:968
func yaml_parser_unroll_indent(parser *yaml_parser_t, column int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:968
	_go_fuzz_dep_.CoverTab[126968]++

										if parser.flow_level > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:970
		_go_fuzz_dep_.CoverTab[126971]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:971
		// _ = "end of CoverTab[126971]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:972
		_go_fuzz_dep_.CoverTab[126972]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:972
		// _ = "end of CoverTab[126972]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:972
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:972
	// _ = "end of CoverTab[126968]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:972
	_go_fuzz_dep_.CoverTab[126969]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:975
	for parser.indent > column {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:975
		_go_fuzz_dep_.CoverTab[126973]++

											token := yaml_token_t{
			typ:		yaml_BLOCK_END_TOKEN,
			start_mark:	parser.mark,
			end_mark:	parser.mark,
		}
											yaml_insert_token(parser, -1, &token)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:985
		parser.indent = parser.indents[len(parser.indents)-1]
											parser.indents = parser.indents[:len(parser.indents)-1]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:986
		// _ = "end of CoverTab[126973]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:987
	// _ = "end of CoverTab[126969]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:987
	_go_fuzz_dep_.CoverTab[126970]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:988
	// _ = "end of CoverTab[126970]"
}

// Initialize the scanner and produce the STREAM-START token.
func yaml_parser_fetch_stream_start(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:992
	_go_fuzz_dep_.CoverTab[126974]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:995
	parser.indent = -1

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:998
	parser.simple_keys = append(parser.simple_keys, yaml_simple_key_t{})

										parser.simple_keys_by_tok = make(map[int]int)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1003
	parser.simple_key_allowed = true

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1006
	parser.stream_start_produced = true

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1009
	token := yaml_token_t{
		typ:		yaml_STREAM_START_TOKEN,
		start_mark:	parser.mark,
		end_mark:	parser.mark,
		encoding:	parser.encoding,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1016
	// _ = "end of CoverTab[126974]"
}

// Produce the STREAM-END token and shut down the scanner.
func yaml_parser_fetch_stream_end(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1020
	_go_fuzz_dep_.CoverTab[126975]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1023
	if parser.mark.column != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1023
		_go_fuzz_dep_.CoverTab[126979]++
											parser.mark.column = 0
											parser.mark.line++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1025
		// _ = "end of CoverTab[126979]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1026
		_go_fuzz_dep_.CoverTab[126980]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1026
		// _ = "end of CoverTab[126980]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1026
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1026
	// _ = "end of CoverTab[126975]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1026
	_go_fuzz_dep_.CoverTab[126976]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1029
	if !yaml_parser_unroll_indent(parser, -1) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1029
		_go_fuzz_dep_.CoverTab[126981]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1030
		// _ = "end of CoverTab[126981]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1031
		_go_fuzz_dep_.CoverTab[126982]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1031
		// _ = "end of CoverTab[126982]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1031
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1031
	// _ = "end of CoverTab[126976]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1031
	_go_fuzz_dep_.CoverTab[126977]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1034
	if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1034
		_go_fuzz_dep_.CoverTab[126983]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1035
		// _ = "end of CoverTab[126983]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1036
		_go_fuzz_dep_.CoverTab[126984]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1036
		// _ = "end of CoverTab[126984]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1036
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1036
	// _ = "end of CoverTab[126977]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1036
	_go_fuzz_dep_.CoverTab[126978]++

										parser.simple_key_allowed = false

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1041
	token := yaml_token_t{
		typ:		yaml_STREAM_END_TOKEN,
		start_mark:	parser.mark,
		end_mark:	parser.mark,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1047
	// _ = "end of CoverTab[126978]"
}

// Produce a VERSION-DIRECTIVE or TAG-DIRECTIVE token.
func yaml_parser_fetch_directive(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1051
	_go_fuzz_dep_.CoverTab[126985]++

										if !yaml_parser_unroll_indent(parser, -1) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1053
		_go_fuzz_dep_.CoverTab[126989]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1054
		// _ = "end of CoverTab[126989]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1055
		_go_fuzz_dep_.CoverTab[126990]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1055
		// _ = "end of CoverTab[126990]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1055
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1055
	// _ = "end of CoverTab[126985]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1055
	_go_fuzz_dep_.CoverTab[126986]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1058
	if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1058
		_go_fuzz_dep_.CoverTab[126991]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1059
		// _ = "end of CoverTab[126991]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1060
		_go_fuzz_dep_.CoverTab[126992]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1060
		// _ = "end of CoverTab[126992]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1060
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1060
	// _ = "end of CoverTab[126986]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1060
	_go_fuzz_dep_.CoverTab[126987]++

										parser.simple_key_allowed = false

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1065
	token := yaml_token_t{}
	if !yaml_parser_scan_directive(parser, &token) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1066
		_go_fuzz_dep_.CoverTab[126993]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1067
		// _ = "end of CoverTab[126993]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1068
		_go_fuzz_dep_.CoverTab[126994]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1068
		// _ = "end of CoverTab[126994]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1068
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1068
	// _ = "end of CoverTab[126987]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1068
	_go_fuzz_dep_.CoverTab[126988]++

										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1071
	// _ = "end of CoverTab[126988]"
}

// Produce the DOCUMENT-START or DOCUMENT-END token.
func yaml_parser_fetch_document_indicator(parser *yaml_parser_t, typ yaml_token_type_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1075
	_go_fuzz_dep_.CoverTab[126995]++

										if !yaml_parser_unroll_indent(parser, -1) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1077
		_go_fuzz_dep_.CoverTab[126998]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1078
		// _ = "end of CoverTab[126998]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1079
		_go_fuzz_dep_.CoverTab[126999]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1079
		// _ = "end of CoverTab[126999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1079
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1079
	// _ = "end of CoverTab[126995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1079
	_go_fuzz_dep_.CoverTab[126996]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1082
	if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1082
		_go_fuzz_dep_.CoverTab[127000]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1083
		// _ = "end of CoverTab[127000]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1084
		_go_fuzz_dep_.CoverTab[127001]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1084
		// _ = "end of CoverTab[127001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1084
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1084
	// _ = "end of CoverTab[126996]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1084
	_go_fuzz_dep_.CoverTab[126997]++

										parser.simple_key_allowed = false

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1089
	start_mark := parser.mark

										skip(parser)
										skip(parser)
										skip(parser)

										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1098
	token := yaml_token_t{
		typ:		typ,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}

										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1105
	// _ = "end of CoverTab[126997]"
}

// Produce the FLOW-SEQUENCE-START or FLOW-MAPPING-START token.
func yaml_parser_fetch_flow_collection_start(parser *yaml_parser_t, typ yaml_token_type_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1109
	_go_fuzz_dep_.CoverTab[127002]++

										if !yaml_parser_save_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1111
		_go_fuzz_dep_.CoverTab[127005]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1112
		// _ = "end of CoverTab[127005]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1113
		_go_fuzz_dep_.CoverTab[127006]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1113
		// _ = "end of CoverTab[127006]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1113
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1113
	// _ = "end of CoverTab[127002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1113
	_go_fuzz_dep_.CoverTab[127003]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1116
	if !yaml_parser_increase_flow_level(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1116
		_go_fuzz_dep_.CoverTab[127007]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1117
		// _ = "end of CoverTab[127007]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1118
		_go_fuzz_dep_.CoverTab[127008]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1118
		// _ = "end of CoverTab[127008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1118
	// _ = "end of CoverTab[127003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1118
	_go_fuzz_dep_.CoverTab[127004]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1121
	parser.simple_key_allowed = true

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1124
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1129
	token := yaml_token_t{
		typ:		typ,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}

										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1136
	// _ = "end of CoverTab[127004]"
}

// Produce the FLOW-SEQUENCE-END or FLOW-MAPPING-END token.
func yaml_parser_fetch_flow_collection_end(parser *yaml_parser_t, typ yaml_token_type_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1140
	_go_fuzz_dep_.CoverTab[127009]++

										if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1142
		_go_fuzz_dep_.CoverTab[127012]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1143
		// _ = "end of CoverTab[127012]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1144
		_go_fuzz_dep_.CoverTab[127013]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1144
		// _ = "end of CoverTab[127013]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1144
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1144
	// _ = "end of CoverTab[127009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1144
	_go_fuzz_dep_.CoverTab[127010]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1147
	if !yaml_parser_decrease_flow_level(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1147
		_go_fuzz_dep_.CoverTab[127014]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1148
		// _ = "end of CoverTab[127014]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1149
		_go_fuzz_dep_.CoverTab[127015]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1149
		// _ = "end of CoverTab[127015]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1149
	// _ = "end of CoverTab[127010]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1149
	_go_fuzz_dep_.CoverTab[127011]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1152
	parser.simple_key_allowed = false

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1156
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1161
	token := yaml_token_t{
		typ:		typ,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}

										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1168
	// _ = "end of CoverTab[127011]"
}

// Produce the FLOW-ENTRY token.
func yaml_parser_fetch_flow_entry(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1172
	_go_fuzz_dep_.CoverTab[127016]++

										if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1174
		_go_fuzz_dep_.CoverTab[127018]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1175
		// _ = "end of CoverTab[127018]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1176
		_go_fuzz_dep_.CoverTab[127019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1176
		// _ = "end of CoverTab[127019]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1176
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1176
	// _ = "end of CoverTab[127016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1176
	_go_fuzz_dep_.CoverTab[127017]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1179
	parser.simple_key_allowed = true

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1182
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1187
	token := yaml_token_t{
		typ:		yaml_FLOW_ENTRY_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1193
	// _ = "end of CoverTab[127017]"
}

// Produce the BLOCK-ENTRY token.
func yaml_parser_fetch_block_entry(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1197
	_go_fuzz_dep_.CoverTab[127020]++

										if parser.flow_level == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1199
		_go_fuzz_dep_.CoverTab[127023]++

											if !parser.simple_key_allowed {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1201
			_go_fuzz_dep_.CoverTab[127025]++
												return yaml_parser_set_scanner_error(parser, "", parser.mark,
				"block sequence entries are not allowed in this context")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1203
			// _ = "end of CoverTab[127025]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1204
			_go_fuzz_dep_.CoverTab[127026]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1204
			// _ = "end of CoverTab[127026]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1204
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1204
		// _ = "end of CoverTab[127023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1204
		_go_fuzz_dep_.CoverTab[127024]++

											if !yaml_parser_roll_indent(parser, parser.mark.column, -1, yaml_BLOCK_SEQUENCE_START_TOKEN, parser.mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1206
			_go_fuzz_dep_.CoverTab[127027]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1207
			// _ = "end of CoverTab[127027]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1208
			_go_fuzz_dep_.CoverTab[127028]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1208
			// _ = "end of CoverTab[127028]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1208
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1208
		// _ = "end of CoverTab[127024]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1209
		_go_fuzz_dep_.CoverTab[127029]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1209
		// _ = "end of CoverTab[127029]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1213
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1213
	// _ = "end of CoverTab[127020]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1213
	_go_fuzz_dep_.CoverTab[127021]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1216
	if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1216
		_go_fuzz_dep_.CoverTab[127030]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1217
		// _ = "end of CoverTab[127030]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1218
		_go_fuzz_dep_.CoverTab[127031]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1218
		// _ = "end of CoverTab[127031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1218
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1218
	// _ = "end of CoverTab[127021]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1218
	_go_fuzz_dep_.CoverTab[127022]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1221
	parser.simple_key_allowed = true

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1224
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1229
	token := yaml_token_t{
		typ:		yaml_BLOCK_ENTRY_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1235
	// _ = "end of CoverTab[127022]"
}

// Produce the KEY token.
func yaml_parser_fetch_key(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1239
	_go_fuzz_dep_.CoverTab[127032]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1242
	if parser.flow_level == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1242
		_go_fuzz_dep_.CoverTab[127035]++

											if !parser.simple_key_allowed {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1244
			_go_fuzz_dep_.CoverTab[127037]++
												return yaml_parser_set_scanner_error(parser, "", parser.mark,
				"mapping keys are not allowed in this context")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1246
			// _ = "end of CoverTab[127037]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1247
			_go_fuzz_dep_.CoverTab[127038]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1247
			// _ = "end of CoverTab[127038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1247
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1247
		// _ = "end of CoverTab[127035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1247
		_go_fuzz_dep_.CoverTab[127036]++

											if !yaml_parser_roll_indent(parser, parser.mark.column, -1, yaml_BLOCK_MAPPING_START_TOKEN, parser.mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1249
			_go_fuzz_dep_.CoverTab[127039]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1250
			// _ = "end of CoverTab[127039]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1251
			_go_fuzz_dep_.CoverTab[127040]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1251
			// _ = "end of CoverTab[127040]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1251
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1251
		// _ = "end of CoverTab[127036]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1252
		_go_fuzz_dep_.CoverTab[127041]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1252
		// _ = "end of CoverTab[127041]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1252
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1252
	// _ = "end of CoverTab[127032]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1252
	_go_fuzz_dep_.CoverTab[127033]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1255
	if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1255
		_go_fuzz_dep_.CoverTab[127042]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1256
		// _ = "end of CoverTab[127042]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1257
		_go_fuzz_dep_.CoverTab[127043]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1257
		// _ = "end of CoverTab[127043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1257
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1257
	// _ = "end of CoverTab[127033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1257
	_go_fuzz_dep_.CoverTab[127034]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1260
	parser.simple_key_allowed = parser.flow_level == 0

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1263
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1268
	token := yaml_token_t{
		typ:		yaml_KEY_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1274
	// _ = "end of CoverTab[127034]"
}

// Produce the VALUE token.
func yaml_parser_fetch_value(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1278
	_go_fuzz_dep_.CoverTab[127044]++

										simple_key := &parser.simple_keys[len(parser.simple_keys)-1]

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1283
	if valid, ok := yaml_simple_key_is_valid(parser, simple_key); !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1283
		_go_fuzz_dep_.CoverTab[127046]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1284
		// _ = "end of CoverTab[127046]"

	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1286
		_go_fuzz_dep_.CoverTab[127047]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1286
		if valid {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1286
			_go_fuzz_dep_.CoverTab[127048]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1289
			token := yaml_token_t{
				typ:		yaml_KEY_TOKEN,
				start_mark:	simple_key.mark,
				end_mark:	simple_key.mark,
			}
												yaml_insert_token(parser, simple_key.token_number-parser.tokens_parsed, &token)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1297
			if !yaml_parser_roll_indent(parser, simple_key.mark.column,
				simple_key.token_number,
				yaml_BLOCK_MAPPING_START_TOKEN, simple_key.mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1299
				_go_fuzz_dep_.CoverTab[127050]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1300
				// _ = "end of CoverTab[127050]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1301
				_go_fuzz_dep_.CoverTab[127051]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1301
				// _ = "end of CoverTab[127051]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1301
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1301
			// _ = "end of CoverTab[127048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1301
			_go_fuzz_dep_.CoverTab[127049]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1304
			simple_key.possible = false
												delete(parser.simple_keys_by_tok, simple_key.token_number)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1308
			parser.simple_key_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1308
			// _ = "end of CoverTab[127049]"

		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1310
			_go_fuzz_dep_.CoverTab[127052]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1314
			if parser.flow_level == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1314
				_go_fuzz_dep_.CoverTab[127054]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1317
				if !parser.simple_key_allowed {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1317
					_go_fuzz_dep_.CoverTab[127056]++
														return yaml_parser_set_scanner_error(parser, "", parser.mark,
						"mapping values are not allowed in this context")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1319
					// _ = "end of CoverTab[127056]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1320
					_go_fuzz_dep_.CoverTab[127057]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1320
					// _ = "end of CoverTab[127057]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1320
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1320
				// _ = "end of CoverTab[127054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1320
				_go_fuzz_dep_.CoverTab[127055]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1323
				if !yaml_parser_roll_indent(parser, parser.mark.column, -1, yaml_BLOCK_MAPPING_START_TOKEN, parser.mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1323
					_go_fuzz_dep_.CoverTab[127058]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1324
					// _ = "end of CoverTab[127058]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1325
					_go_fuzz_dep_.CoverTab[127059]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1325
					// _ = "end of CoverTab[127059]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1325
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1325
				// _ = "end of CoverTab[127055]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1326
				_go_fuzz_dep_.CoverTab[127060]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1326
				// _ = "end of CoverTab[127060]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1326
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1326
			// _ = "end of CoverTab[127052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1326
			_go_fuzz_dep_.CoverTab[127053]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1329
			parser.simple_key_allowed = parser.flow_level == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1329
			// _ = "end of CoverTab[127053]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1330
		// _ = "end of CoverTab[127047]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1330
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1330
	// _ = "end of CoverTab[127044]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1330
	_go_fuzz_dep_.CoverTab[127045]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1333
	start_mark := parser.mark
										skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1338
	token := yaml_token_t{
		typ:		yaml_VALUE_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
	}
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1344
	// _ = "end of CoverTab[127045]"
}

// Produce the ALIAS or ANCHOR token.
func yaml_parser_fetch_anchor(parser *yaml_parser_t, typ yaml_token_type_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1348
	_go_fuzz_dep_.CoverTab[127061]++

										if !yaml_parser_save_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1350
		_go_fuzz_dep_.CoverTab[127064]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1351
		// _ = "end of CoverTab[127064]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1352
		_go_fuzz_dep_.CoverTab[127065]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1352
		// _ = "end of CoverTab[127065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1352
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1352
	// _ = "end of CoverTab[127061]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1352
	_go_fuzz_dep_.CoverTab[127062]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1355
	parser.simple_key_allowed = false

	// Create the ALIAS or ANCHOR token and append it to the queue.
	var token yaml_token_t
	if !yaml_parser_scan_anchor(parser, &token, typ) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1359
		_go_fuzz_dep_.CoverTab[127066]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1360
		// _ = "end of CoverTab[127066]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1361
		_go_fuzz_dep_.CoverTab[127067]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1361
		// _ = "end of CoverTab[127067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1361
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1361
	// _ = "end of CoverTab[127062]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1361
	_go_fuzz_dep_.CoverTab[127063]++
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1363
	// _ = "end of CoverTab[127063]"
}

// Produce the TAG token.
func yaml_parser_fetch_tag(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1367
	_go_fuzz_dep_.CoverTab[127068]++

										if !yaml_parser_save_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1369
		_go_fuzz_dep_.CoverTab[127071]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1370
		// _ = "end of CoverTab[127071]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1371
		_go_fuzz_dep_.CoverTab[127072]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1371
		// _ = "end of CoverTab[127072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1371
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1371
	// _ = "end of CoverTab[127068]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1371
	_go_fuzz_dep_.CoverTab[127069]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1374
	parser.simple_key_allowed = false

	// Create the TAG token and append it to the queue.
	var token yaml_token_t
	if !yaml_parser_scan_tag(parser, &token) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1378
		_go_fuzz_dep_.CoverTab[127073]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1379
		// _ = "end of CoverTab[127073]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1380
		_go_fuzz_dep_.CoverTab[127074]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1380
		// _ = "end of CoverTab[127074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1380
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1380
	// _ = "end of CoverTab[127069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1380
	_go_fuzz_dep_.CoverTab[127070]++
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1382
	// _ = "end of CoverTab[127070]"
}

// Produce the SCALAR(...,literal) or SCALAR(...,folded) tokens.
func yaml_parser_fetch_block_scalar(parser *yaml_parser_t, literal bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1386
	_go_fuzz_dep_.CoverTab[127075]++

										if !yaml_parser_remove_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1388
		_go_fuzz_dep_.CoverTab[127078]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1389
		// _ = "end of CoverTab[127078]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1390
		_go_fuzz_dep_.CoverTab[127079]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1390
		// _ = "end of CoverTab[127079]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1390
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1390
	// _ = "end of CoverTab[127075]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1390
	_go_fuzz_dep_.CoverTab[127076]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1393
	parser.simple_key_allowed = true

	// Create the SCALAR token and append it to the queue.
	var token yaml_token_t
	if !yaml_parser_scan_block_scalar(parser, &token, literal) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1397
		_go_fuzz_dep_.CoverTab[127080]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1398
		// _ = "end of CoverTab[127080]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1399
		_go_fuzz_dep_.CoverTab[127081]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1399
		// _ = "end of CoverTab[127081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1399
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1399
	// _ = "end of CoverTab[127076]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1399
	_go_fuzz_dep_.CoverTab[127077]++
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1401
	// _ = "end of CoverTab[127077]"
}

// Produce the SCALAR(...,single-quoted) or SCALAR(...,double-quoted) tokens.
func yaml_parser_fetch_flow_scalar(parser *yaml_parser_t, single bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1405
	_go_fuzz_dep_.CoverTab[127082]++

										if !yaml_parser_save_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1407
		_go_fuzz_dep_.CoverTab[127085]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1408
		// _ = "end of CoverTab[127085]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1409
		_go_fuzz_dep_.CoverTab[127086]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1409
		// _ = "end of CoverTab[127086]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1409
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1409
	// _ = "end of CoverTab[127082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1409
	_go_fuzz_dep_.CoverTab[127083]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1412
	parser.simple_key_allowed = false

	// Create the SCALAR token and append it to the queue.
	var token yaml_token_t
	if !yaml_parser_scan_flow_scalar(parser, &token, single) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1416
		_go_fuzz_dep_.CoverTab[127087]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1417
		// _ = "end of CoverTab[127087]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1418
		_go_fuzz_dep_.CoverTab[127088]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1418
		// _ = "end of CoverTab[127088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1418
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1418
	// _ = "end of CoverTab[127083]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1418
	_go_fuzz_dep_.CoverTab[127084]++
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1420
	// _ = "end of CoverTab[127084]"
}

// Produce the SCALAR(...,plain) token.
func yaml_parser_fetch_plain_scalar(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1424
	_go_fuzz_dep_.CoverTab[127089]++

										if !yaml_parser_save_simple_key(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1426
		_go_fuzz_dep_.CoverTab[127092]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1427
		// _ = "end of CoverTab[127092]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1428
		_go_fuzz_dep_.CoverTab[127093]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1428
		// _ = "end of CoverTab[127093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1428
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1428
	// _ = "end of CoverTab[127089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1428
	_go_fuzz_dep_.CoverTab[127090]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1431
	parser.simple_key_allowed = false

	// Create the SCALAR token and append it to the queue.
	var token yaml_token_t
	if !yaml_parser_scan_plain_scalar(parser, &token) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1435
		_go_fuzz_dep_.CoverTab[127094]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1436
		// _ = "end of CoverTab[127094]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1437
		_go_fuzz_dep_.CoverTab[127095]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1437
		// _ = "end of CoverTab[127095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1437
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1437
	// _ = "end of CoverTab[127090]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1437
	_go_fuzz_dep_.CoverTab[127091]++
										yaml_insert_token(parser, -1, &token)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1439
	// _ = "end of CoverTab[127091]"
}

// Eat whitespaces and comments until the next token is found.
func yaml_parser_scan_to_next_token(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1443
	_go_fuzz_dep_.CoverTab[127096]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1446
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1446
		_go_fuzz_dep_.CoverTab[127098]++

											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1448
			_go_fuzz_dep_.CoverTab[127104]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1448
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1448
			// _ = "end of CoverTab[127104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1448
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1448
			_go_fuzz_dep_.CoverTab[127105]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1449
			// _ = "end of CoverTab[127105]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1450
			_go_fuzz_dep_.CoverTab[127106]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1450
			// _ = "end of CoverTab[127106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1450
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1450
		// _ = "end of CoverTab[127098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1450
		_go_fuzz_dep_.CoverTab[127099]++
											if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1451
			_go_fuzz_dep_.CoverTab[127107]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1451
			return is_bom(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1451
			// _ = "end of CoverTab[127107]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1451
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1451
			_go_fuzz_dep_.CoverTab[127108]++
												skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1452
			// _ = "end of CoverTab[127108]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1453
			_go_fuzz_dep_.CoverTab[127109]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1453
			// _ = "end of CoverTab[127109]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1453
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1453
		// _ = "end of CoverTab[127099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1453
		_go_fuzz_dep_.CoverTab[127100]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
		if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
			_go_fuzz_dep_.CoverTab[127110]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
			// _ = "end of CoverTab[127110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1460
			_go_fuzz_dep_.CoverTab[127111]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1461
			// _ = "end of CoverTab[127111]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1462
			_go_fuzz_dep_.CoverTab[127112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1462
			// _ = "end of CoverTab[127112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1462
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1462
		// _ = "end of CoverTab[127100]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1462
		_go_fuzz_dep_.CoverTab[127101]++

											for parser.buffer[parser.buffer_pos] == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			_go_fuzz_dep_.CoverTab[127113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			return ((parser.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				_go_fuzz_dep_.CoverTab[127114]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				return !parser.simple_key_allowed
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				// _ = "end of CoverTab[127114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				_go_fuzz_dep_.CoverTab[127115]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				return parser.buffer[parser.buffer_pos] == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
				// _ = "end of CoverTab[127115]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			// _ = "end of CoverTab[127113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1464
			_go_fuzz_dep_.CoverTab[127116]++
												skip(parser)
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1466
				_go_fuzz_dep_.CoverTab[127117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1466
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1466
				// _ = "end of CoverTab[127117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1466
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1466
				_go_fuzz_dep_.CoverTab[127118]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1467
				// _ = "end of CoverTab[127118]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1468
				_go_fuzz_dep_.CoverTab[127119]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1468
				// _ = "end of CoverTab[127119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1468
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1468
			// _ = "end of CoverTab[127116]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1469
		// _ = "end of CoverTab[127101]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1469
		_go_fuzz_dep_.CoverTab[127102]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1472
		if parser.buffer[parser.buffer_pos] == '#' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1472
			_go_fuzz_dep_.CoverTab[127120]++
												for !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1473
				_go_fuzz_dep_.CoverTab[127121]++
													skip(parser)
													if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1475
					_go_fuzz_dep_.CoverTab[127122]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1475
					return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1475
					// _ = "end of CoverTab[127122]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1475
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1475
					_go_fuzz_dep_.CoverTab[127123]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1476
					// _ = "end of CoverTab[127123]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1477
					_go_fuzz_dep_.CoverTab[127124]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1477
					// _ = "end of CoverTab[127124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1477
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1477
				// _ = "end of CoverTab[127121]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1478
			// _ = "end of CoverTab[127120]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1479
			_go_fuzz_dep_.CoverTab[127125]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1479
			// _ = "end of CoverTab[127125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1479
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1479
		// _ = "end of CoverTab[127102]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1479
		_go_fuzz_dep_.CoverTab[127103]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1482
		if is_break(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1482
			_go_fuzz_dep_.CoverTab[127126]++
												if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1483
				_go_fuzz_dep_.CoverTab[127128]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1483
				return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1483
				// _ = "end of CoverTab[127128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1483
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1483
				_go_fuzz_dep_.CoverTab[127129]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1484
				// _ = "end of CoverTab[127129]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1485
				_go_fuzz_dep_.CoverTab[127130]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1485
				// _ = "end of CoverTab[127130]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1485
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1485
			// _ = "end of CoverTab[127126]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1485
			_go_fuzz_dep_.CoverTab[127127]++
												skip_line(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1489
			if parser.flow_level == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1489
				_go_fuzz_dep_.CoverTab[127131]++
													parser.simple_key_allowed = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1490
				// _ = "end of CoverTab[127131]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1491
				_go_fuzz_dep_.CoverTab[127132]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1491
				// _ = "end of CoverTab[127132]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1491
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1491
			// _ = "end of CoverTab[127127]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1492
			_go_fuzz_dep_.CoverTab[127133]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1493
			// _ = "end of CoverTab[127133]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1494
		// _ = "end of CoverTab[127103]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1495
	// _ = "end of CoverTab[127096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1495
	_go_fuzz_dep_.CoverTab[127097]++

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1497
	// _ = "end of CoverTab[127097]"
}

// Scan a YAML-DIRECTIVE or TAG-DIRECTIVE token.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
// Scope:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//	%YAML    1.1    # a comment \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//	^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//	%TAG    !yaml!  tag:yaml.org,2002:  \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1500
//	^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1508
func yaml_parser_scan_directive(parser *yaml_parser_t, token *yaml_token_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1508
	_go_fuzz_dep_.CoverTab[127134]++

										start_mark := parser.mark
										skip(parser)

	// Scan the directive name.
	var name []byte
	if !yaml_parser_scan_directive_name(parser, start_mark, &name) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1515
		_go_fuzz_dep_.CoverTab[127142]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1516
		// _ = "end of CoverTab[127142]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1517
		_go_fuzz_dep_.CoverTab[127143]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1517
		// _ = "end of CoverTab[127143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1517
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1517
	// _ = "end of CoverTab[127134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1517
	_go_fuzz_dep_.CoverTab[127135]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1520
	if bytes.Equal(name, []byte("YAML")) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1520
		_go_fuzz_dep_.CoverTab[127144]++
		// Scan the VERSION directive value.
		var major, minor int8
		if !yaml_parser_scan_version_directive_value(parser, start_mark, &major, &minor) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1523
			_go_fuzz_dep_.CoverTab[127146]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1524
			// _ = "end of CoverTab[127146]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1525
			_go_fuzz_dep_.CoverTab[127147]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1525
			// _ = "end of CoverTab[127147]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1525
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1525
		// _ = "end of CoverTab[127144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1525
		_go_fuzz_dep_.CoverTab[127145]++
											end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1529
		*token = yaml_token_t{
			typ:		yaml_VERSION_DIRECTIVE_TOKEN,
			start_mark:	start_mark,
			end_mark:	end_mark,
			major:		major,
			minor:		minor,
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1535
		// _ = "end of CoverTab[127145]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1538
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1538
		_go_fuzz_dep_.CoverTab[127148]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1538
		if bytes.Equal(name, []byte("TAG")) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1538
			_go_fuzz_dep_.CoverTab[127149]++
			// Scan the TAG directive value.
			var handle, prefix []byte
			if !yaml_parser_scan_tag_directive_value(parser, start_mark, &handle, &prefix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1541
				_go_fuzz_dep_.CoverTab[127151]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1542
				// _ = "end of CoverTab[127151]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1543
				_go_fuzz_dep_.CoverTab[127152]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1543
				// _ = "end of CoverTab[127152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1543
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1543
			// _ = "end of CoverTab[127149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1543
			_go_fuzz_dep_.CoverTab[127150]++
												end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1547
			*token = yaml_token_t{
				typ:		yaml_TAG_DIRECTIVE_TOKEN,
				start_mark:	start_mark,
				end_mark:	end_mark,
				value:		handle,
				prefix:		prefix,
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1553
			// _ = "end of CoverTab[127150]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1556
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1556
			_go_fuzz_dep_.CoverTab[127153]++
												yaml_parser_set_scanner_error(parser, "while scanning a directive",
				start_mark, "found unknown directive name")
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1559
			// _ = "end of CoverTab[127153]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1560
		// _ = "end of CoverTab[127148]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1560
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1560
	// _ = "end of CoverTab[127135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1560
	_go_fuzz_dep_.CoverTab[127136]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
		_go_fuzz_dep_.CoverTab[127154]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
		// _ = "end of CoverTab[127154]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1563
		_go_fuzz_dep_.CoverTab[127155]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1564
		// _ = "end of CoverTab[127155]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1565
		_go_fuzz_dep_.CoverTab[127156]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1565
		// _ = "end of CoverTab[127156]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1565
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1565
	// _ = "end of CoverTab[127136]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1565
	_go_fuzz_dep_.CoverTab[127137]++

										for is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1567
		_go_fuzz_dep_.CoverTab[127157]++
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1569
			_go_fuzz_dep_.CoverTab[127158]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1569
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1569
			// _ = "end of CoverTab[127158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1569
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1569
			_go_fuzz_dep_.CoverTab[127159]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1570
			// _ = "end of CoverTab[127159]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1571
			_go_fuzz_dep_.CoverTab[127160]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1571
			// _ = "end of CoverTab[127160]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1571
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1571
		// _ = "end of CoverTab[127157]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1572
	// _ = "end of CoverTab[127137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1572
	_go_fuzz_dep_.CoverTab[127138]++

										if parser.buffer[parser.buffer_pos] == '#' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1574
		_go_fuzz_dep_.CoverTab[127161]++
											for !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1575
			_go_fuzz_dep_.CoverTab[127162]++
												skip(parser)
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1577
				_go_fuzz_dep_.CoverTab[127163]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1577
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1577
				// _ = "end of CoverTab[127163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1577
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1577
				_go_fuzz_dep_.CoverTab[127164]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1578
				// _ = "end of CoverTab[127164]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1579
				_go_fuzz_dep_.CoverTab[127165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1579
				// _ = "end of CoverTab[127165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1579
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1579
			// _ = "end of CoverTab[127162]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1580
		// _ = "end of CoverTab[127161]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1581
		_go_fuzz_dep_.CoverTab[127166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1581
		// _ = "end of CoverTab[127166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1581
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1581
	// _ = "end of CoverTab[127138]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1581
	_go_fuzz_dep_.CoverTab[127139]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1584
	if !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1584
		_go_fuzz_dep_.CoverTab[127167]++
											yaml_parser_set_scanner_error(parser, "while scanning a directive",
			start_mark, "did not find expected comment or line break")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1587
		// _ = "end of CoverTab[127167]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1588
		_go_fuzz_dep_.CoverTab[127168]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1588
		// _ = "end of CoverTab[127168]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1588
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1588
	// _ = "end of CoverTab[127139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1588
	_go_fuzz_dep_.CoverTab[127140]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1591
	if is_break(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1591
		_go_fuzz_dep_.CoverTab[127169]++
											if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1592
			_go_fuzz_dep_.CoverTab[127171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1592
			return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1592
			// _ = "end of CoverTab[127171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1592
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1592
			_go_fuzz_dep_.CoverTab[127172]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1593
			// _ = "end of CoverTab[127172]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1594
			_go_fuzz_dep_.CoverTab[127173]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1594
			// _ = "end of CoverTab[127173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1594
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1594
		// _ = "end of CoverTab[127169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1594
		_go_fuzz_dep_.CoverTab[127170]++
											skip_line(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1595
		// _ = "end of CoverTab[127170]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1596
		_go_fuzz_dep_.CoverTab[127174]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1596
		// _ = "end of CoverTab[127174]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1596
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1596
	// _ = "end of CoverTab[127140]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1596
	_go_fuzz_dep_.CoverTab[127141]++

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1598
	// _ = "end of CoverTab[127141]"
}

// Scan the directive name.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
// Scope:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//	%YAML   1.1     # a comment \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//	 ^^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//	%TAG    !yaml!  tag:yaml.org,2002:  \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1601
//	 ^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1609
func yaml_parser_scan_directive_name(parser *yaml_parser_t, start_mark yaml_mark_t, name *[]byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1609
	_go_fuzz_dep_.CoverTab[127175]++

										if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1611
		_go_fuzz_dep_.CoverTab[127180]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1611
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1611
		// _ = "end of CoverTab[127180]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1611
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1611
		_go_fuzz_dep_.CoverTab[127181]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1612
		// _ = "end of CoverTab[127181]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1613
		_go_fuzz_dep_.CoverTab[127182]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1613
		// _ = "end of CoverTab[127182]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1613
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1613
	// _ = "end of CoverTab[127175]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1613
	_go_fuzz_dep_.CoverTab[127176]++

										var s []byte
										for is_alpha(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1616
		_go_fuzz_dep_.CoverTab[127183]++
											s = read(parser, s)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1618
			_go_fuzz_dep_.CoverTab[127184]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1618
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1618
			// _ = "end of CoverTab[127184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1618
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1618
			_go_fuzz_dep_.CoverTab[127185]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1619
			// _ = "end of CoverTab[127185]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1620
			_go_fuzz_dep_.CoverTab[127186]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1620
			// _ = "end of CoverTab[127186]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1620
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1620
		// _ = "end of CoverTab[127183]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1621
	// _ = "end of CoverTab[127176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1621
	_go_fuzz_dep_.CoverTab[127177]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1624
	if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1624
		_go_fuzz_dep_.CoverTab[127187]++
											yaml_parser_set_scanner_error(parser, "while scanning a directive",
			start_mark, "could not find expected directive name")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1627
		// _ = "end of CoverTab[127187]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1628
		_go_fuzz_dep_.CoverTab[127188]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1628
		// _ = "end of CoverTab[127188]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1628
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1628
	// _ = "end of CoverTab[127177]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1628
	_go_fuzz_dep_.CoverTab[127178]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1631
	if !is_blankz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1631
		_go_fuzz_dep_.CoverTab[127189]++
											yaml_parser_set_scanner_error(parser, "while scanning a directive",
			start_mark, "found unexpected non-alphabetical character")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1634
		// _ = "end of CoverTab[127189]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1635
		_go_fuzz_dep_.CoverTab[127190]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1635
		// _ = "end of CoverTab[127190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1635
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1635
	// _ = "end of CoverTab[127178]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1635
	_go_fuzz_dep_.CoverTab[127179]++
										*name = s
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1637
	// _ = "end of CoverTab[127179]"
}

// Scan the value of VERSION-DIRECTIVE.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1640
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1640
// Scope:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1640
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1640
//	%YAML   1.1     # a comment \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1640
//	     ^^^^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1645
func yaml_parser_scan_version_directive_value(parser *yaml_parser_t, start_mark yaml_mark_t, major, minor *int8) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1645
	_go_fuzz_dep_.CoverTab[127191]++

										if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1647
		_go_fuzz_dep_.CoverTab[127197]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1647
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1647
		// _ = "end of CoverTab[127197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1647
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1647
		_go_fuzz_dep_.CoverTab[127198]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1648
		// _ = "end of CoverTab[127198]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1649
		_go_fuzz_dep_.CoverTab[127199]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1649
		// _ = "end of CoverTab[127199]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1649
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1649
	// _ = "end of CoverTab[127191]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1649
	_go_fuzz_dep_.CoverTab[127192]++
										for is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1650
		_go_fuzz_dep_.CoverTab[127200]++
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1652
			_go_fuzz_dep_.CoverTab[127201]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1652
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1652
			// _ = "end of CoverTab[127201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1652
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1652
			_go_fuzz_dep_.CoverTab[127202]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1653
			// _ = "end of CoverTab[127202]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1654
			_go_fuzz_dep_.CoverTab[127203]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1654
			// _ = "end of CoverTab[127203]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1654
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1654
		// _ = "end of CoverTab[127200]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1655
	// _ = "end of CoverTab[127192]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1655
	_go_fuzz_dep_.CoverTab[127193]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1658
	if !yaml_parser_scan_version_directive_number(parser, start_mark, major) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1658
		_go_fuzz_dep_.CoverTab[127204]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1659
		// _ = "end of CoverTab[127204]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1660
		_go_fuzz_dep_.CoverTab[127205]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1660
		// _ = "end of CoverTab[127205]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1660
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1660
	// _ = "end of CoverTab[127193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1660
	_go_fuzz_dep_.CoverTab[127194]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1663
	if parser.buffer[parser.buffer_pos] != '.' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1663
		_go_fuzz_dep_.CoverTab[127206]++
											return yaml_parser_set_scanner_error(parser, "while scanning a %YAML directive",
			start_mark, "did not find expected digit or '.' character")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1665
		// _ = "end of CoverTab[127206]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1666
		_go_fuzz_dep_.CoverTab[127207]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1666
		// _ = "end of CoverTab[127207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1666
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1666
	// _ = "end of CoverTab[127194]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1666
	_go_fuzz_dep_.CoverTab[127195]++

										skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1671
	if !yaml_parser_scan_version_directive_number(parser, start_mark, minor) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1671
		_go_fuzz_dep_.CoverTab[127208]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1672
		// _ = "end of CoverTab[127208]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1673
		_go_fuzz_dep_.CoverTab[127209]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1673
		// _ = "end of CoverTab[127209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1673
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1673
	// _ = "end of CoverTab[127195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1673
	_go_fuzz_dep_.CoverTab[127196]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1674
	// _ = "end of CoverTab[127196]"
}

const max_number_length = 2

// Scan the version number of VERSION-DIRECTIVE.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
// Scope:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//	%YAML   1.1     # a comment \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//	        ^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//	%YAML   1.1     # a comment \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1679
//	          ^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1686
func yaml_parser_scan_version_directive_number(parser *yaml_parser_t, start_mark yaml_mark_t, number *int8) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1686
	_go_fuzz_dep_.CoverTab[127210]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
		_go_fuzz_dep_.CoverTab[127214]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
		// _ = "end of CoverTab[127214]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1689
		_go_fuzz_dep_.CoverTab[127215]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1690
		// _ = "end of CoverTab[127215]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1691
		_go_fuzz_dep_.CoverTab[127216]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1691
		// _ = "end of CoverTab[127216]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1691
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1691
	// _ = "end of CoverTab[127210]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1691
	_go_fuzz_dep_.CoverTab[127211]++
										var value, length int8
										for is_digit(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1693
		_go_fuzz_dep_.CoverTab[127217]++

											length++
											if length > max_number_length {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1696
			_go_fuzz_dep_.CoverTab[127219]++
												return yaml_parser_set_scanner_error(parser, "while scanning a %YAML directive",
				start_mark, "found extremely long version number")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1698
			// _ = "end of CoverTab[127219]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1699
			_go_fuzz_dep_.CoverTab[127220]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1699
			// _ = "end of CoverTab[127220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1699
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1699
		// _ = "end of CoverTab[127217]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1699
		_go_fuzz_dep_.CoverTab[127218]++
											value = value*10 + int8(as_digit(parser.buffer, parser.buffer_pos))
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1702
			_go_fuzz_dep_.CoverTab[127221]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1702
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1702
			// _ = "end of CoverTab[127221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1702
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1702
			_go_fuzz_dep_.CoverTab[127222]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1703
			// _ = "end of CoverTab[127222]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1704
			_go_fuzz_dep_.CoverTab[127223]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1704
			// _ = "end of CoverTab[127223]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1704
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1704
		// _ = "end of CoverTab[127218]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1705
	// _ = "end of CoverTab[127211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1705
	_go_fuzz_dep_.CoverTab[127212]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1708
	if length == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1708
		_go_fuzz_dep_.CoverTab[127224]++
											return yaml_parser_set_scanner_error(parser, "while scanning a %YAML directive",
			start_mark, "did not find expected version number")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1710
		// _ = "end of CoverTab[127224]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1711
		_go_fuzz_dep_.CoverTab[127225]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1711
		// _ = "end of CoverTab[127225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1711
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1711
	// _ = "end of CoverTab[127212]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1711
	_go_fuzz_dep_.CoverTab[127213]++
										*number = value
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1713
	// _ = "end of CoverTab[127213]"
}

// Scan the value of a TAG-DIRECTIVE token.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1716
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1716
// Scope:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1716
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1716
//	%TAG    !yaml!  tag:yaml.org,2002:  \n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1716
//	    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1722
func yaml_parser_scan_tag_directive_value(parser *yaml_parser_t, start_mark yaml_mark_t, handle, prefix *[]byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1722
	_go_fuzz_dep_.CoverTab[127226]++
										var handle_value, prefix_value []byte

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
		_go_fuzz_dep_.CoverTab[127236]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
		// _ = "end of CoverTab[127236]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1726
		_go_fuzz_dep_.CoverTab[127237]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1727
		// _ = "end of CoverTab[127237]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1728
		_go_fuzz_dep_.CoverTab[127238]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1728
		// _ = "end of CoverTab[127238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1728
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1728
	// _ = "end of CoverTab[127226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1728
	_go_fuzz_dep_.CoverTab[127227]++

										for is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1730
		_go_fuzz_dep_.CoverTab[127239]++
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1732
			_go_fuzz_dep_.CoverTab[127240]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1732
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1732
			// _ = "end of CoverTab[127240]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1732
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1732
			_go_fuzz_dep_.CoverTab[127241]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1733
			// _ = "end of CoverTab[127241]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1734
			_go_fuzz_dep_.CoverTab[127242]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1734
			// _ = "end of CoverTab[127242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1734
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1734
		// _ = "end of CoverTab[127239]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1735
	// _ = "end of CoverTab[127227]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1735
	_go_fuzz_dep_.CoverTab[127228]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1738
	if !yaml_parser_scan_tag_handle(parser, true, start_mark, &handle_value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1738
		_go_fuzz_dep_.CoverTab[127243]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1739
		// _ = "end of CoverTab[127243]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1740
		_go_fuzz_dep_.CoverTab[127244]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1740
		// _ = "end of CoverTab[127244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1740
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1740
	// _ = "end of CoverTab[127228]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1740
	_go_fuzz_dep_.CoverTab[127229]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
		_go_fuzz_dep_.CoverTab[127245]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
		// _ = "end of CoverTab[127245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1743
		_go_fuzz_dep_.CoverTab[127246]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1744
		// _ = "end of CoverTab[127246]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1745
		_go_fuzz_dep_.CoverTab[127247]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1745
		// _ = "end of CoverTab[127247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1745
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1745
	// _ = "end of CoverTab[127229]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1745
	_go_fuzz_dep_.CoverTab[127230]++
										if !is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1746
		_go_fuzz_dep_.CoverTab[127248]++
											yaml_parser_set_scanner_error(parser, "while scanning a %TAG directive",
			start_mark, "did not find expected whitespace")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1749
		// _ = "end of CoverTab[127248]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1750
		_go_fuzz_dep_.CoverTab[127249]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1750
		// _ = "end of CoverTab[127249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1750
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1750
	// _ = "end of CoverTab[127230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1750
	_go_fuzz_dep_.CoverTab[127231]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1753
	for is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1753
		_go_fuzz_dep_.CoverTab[127250]++
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1755
			_go_fuzz_dep_.CoverTab[127251]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1755
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1755
			// _ = "end of CoverTab[127251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1755
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1755
			_go_fuzz_dep_.CoverTab[127252]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1756
			// _ = "end of CoverTab[127252]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1757
			_go_fuzz_dep_.CoverTab[127253]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1757
			// _ = "end of CoverTab[127253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1757
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1757
		// _ = "end of CoverTab[127250]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1758
	// _ = "end of CoverTab[127231]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1758
	_go_fuzz_dep_.CoverTab[127232]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1761
	if !yaml_parser_scan_tag_uri(parser, true, nil, start_mark, &prefix_value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1761
		_go_fuzz_dep_.CoverTab[127254]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1762
		// _ = "end of CoverTab[127254]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1763
		_go_fuzz_dep_.CoverTab[127255]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1763
		// _ = "end of CoverTab[127255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1763
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1763
	// _ = "end of CoverTab[127232]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1763
	_go_fuzz_dep_.CoverTab[127233]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
		_go_fuzz_dep_.CoverTab[127256]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
		// _ = "end of CoverTab[127256]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1766
		_go_fuzz_dep_.CoverTab[127257]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1767
		// _ = "end of CoverTab[127257]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1768
		_go_fuzz_dep_.CoverTab[127258]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1768
		// _ = "end of CoverTab[127258]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1768
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1768
	// _ = "end of CoverTab[127233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1768
	_go_fuzz_dep_.CoverTab[127234]++
										if !is_blankz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1769
		_go_fuzz_dep_.CoverTab[127259]++
											yaml_parser_set_scanner_error(parser, "while scanning a %TAG directive",
			start_mark, "did not find expected whitespace or line break")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1772
		// _ = "end of CoverTab[127259]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1773
		_go_fuzz_dep_.CoverTab[127260]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1773
		// _ = "end of CoverTab[127260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1773
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1773
	// _ = "end of CoverTab[127234]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1773
	_go_fuzz_dep_.CoverTab[127235]++

										*handle = handle_value
										*prefix = prefix_value
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1777
	// _ = "end of CoverTab[127235]"
}

func yaml_parser_scan_anchor(parser *yaml_parser_t, token *yaml_token_t, typ yaml_token_type_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1780
	_go_fuzz_dep_.CoverTab[127261]++
										var s []byte

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1784
	start_mark := parser.mark
										skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
		_go_fuzz_dep_.CoverTab[127265]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
		// _ = "end of CoverTab[127265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1788
		_go_fuzz_dep_.CoverTab[127266]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1789
		// _ = "end of CoverTab[127266]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1790
		_go_fuzz_dep_.CoverTab[127267]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1790
		// _ = "end of CoverTab[127267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1790
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1790
	// _ = "end of CoverTab[127261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1790
	_go_fuzz_dep_.CoverTab[127262]++

										for is_alpha(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1792
		_go_fuzz_dep_.CoverTab[127268]++
											s = read(parser, s)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1794
			_go_fuzz_dep_.CoverTab[127269]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1794
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1794
			// _ = "end of CoverTab[127269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1794
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1794
			_go_fuzz_dep_.CoverTab[127270]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1795
			// _ = "end of CoverTab[127270]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1796
			_go_fuzz_dep_.CoverTab[127271]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1796
			// _ = "end of CoverTab[127271]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1796
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1796
		// _ = "end of CoverTab[127268]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1797
	// _ = "end of CoverTab[127262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1797
	_go_fuzz_dep_.CoverTab[127263]++

										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1808
	if len(s) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1808
		_go_fuzz_dep_.CoverTab[127272]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1808
		return !(is_blankz(parser.buffer, parser.buffer_pos) || func() bool {
												_go_fuzz_dep_.CoverTab[127273]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1809
			return parser.buffer[parser.buffer_pos] == '?'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1809
			// _ = "end of CoverTab[127273]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1809
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1809
			_go_fuzz_dep_.CoverTab[127274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1809
			return parser.buffer[parser.buffer_pos] == ':'
												// _ = "end of CoverTab[127274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
			_go_fuzz_dep_.CoverTab[127275]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
			return parser.buffer[parser.buffer_pos] == ','
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
			// _ = "end of CoverTab[127275]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
			_go_fuzz_dep_.CoverTab[127276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1810
			return parser.buffer[parser.buffer_pos] == ']'
												// _ = "end of CoverTab[127276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
			_go_fuzz_dep_.CoverTab[127277]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
			return parser.buffer[parser.buffer_pos] == '}'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
			// _ = "end of CoverTab[127277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
			_go_fuzz_dep_.CoverTab[127278]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1811
			return parser.buffer[parser.buffer_pos] == '%'
												// _ = "end of CoverTab[127278]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
			_go_fuzz_dep_.CoverTab[127279]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
			return parser.buffer[parser.buffer_pos] == '@'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
			// _ = "end of CoverTab[127279]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
			_go_fuzz_dep_.CoverTab[127280]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1812
			return parser.buffer[parser.buffer_pos] == '`'
												// _ = "end of CoverTab[127280]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1813
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1813
		// _ = "end of CoverTab[127272]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1813
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1813
		_go_fuzz_dep_.CoverTab[127281]++
											context := "while scanning an alias"
											if typ == yaml_ANCHOR_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1815
			_go_fuzz_dep_.CoverTab[127283]++
												context = "while scanning an anchor"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1816
			// _ = "end of CoverTab[127283]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1817
			_go_fuzz_dep_.CoverTab[127284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1817
			// _ = "end of CoverTab[127284]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1817
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1817
		// _ = "end of CoverTab[127281]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1817
		_go_fuzz_dep_.CoverTab[127282]++
											yaml_parser_set_scanner_error(parser, context, start_mark,
			"did not find expected alphabetic or numeric character")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1820
		// _ = "end of CoverTab[127282]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1821
		_go_fuzz_dep_.CoverTab[127285]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1821
		// _ = "end of CoverTab[127285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1821
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1821
	// _ = "end of CoverTab[127263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1821
	_go_fuzz_dep_.CoverTab[127264]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1824
	*token = yaml_token_t{
		typ:		typ,
		start_mark:	start_mark,
		end_mark:	end_mark,
		value:		s,
	}

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1831
	// _ = "end of CoverTab[127264]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1838
func yaml_parser_scan_tag(parser *yaml_parser_t, token *yaml_token_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1838
	_go_fuzz_dep_.CoverTab[127286]++
										var handle, suffix []byte

										start_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
	if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
		_go_fuzz_dep_.CoverTab[127291]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
		return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
		// _ = "end of CoverTab[127291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1844
		_go_fuzz_dep_.CoverTab[127292]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1845
		// _ = "end of CoverTab[127292]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1846
		_go_fuzz_dep_.CoverTab[127293]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1846
		// _ = "end of CoverTab[127293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1846
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1846
	// _ = "end of CoverTab[127286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1846
	_go_fuzz_dep_.CoverTab[127287]++

										if parser.buffer[parser.buffer_pos+1] == '<' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1848
		_go_fuzz_dep_.CoverTab[127294]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1852
		skip(parser)
											skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1856
		if !yaml_parser_scan_tag_uri(parser, false, nil, start_mark, &suffix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1856
			_go_fuzz_dep_.CoverTab[127297]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1857
			// _ = "end of CoverTab[127297]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1858
			_go_fuzz_dep_.CoverTab[127298]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1858
			// _ = "end of CoverTab[127298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1858
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1858
		// _ = "end of CoverTab[127294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1858
		_go_fuzz_dep_.CoverTab[127295]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1861
		if parser.buffer[parser.buffer_pos] != '>' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1861
			_go_fuzz_dep_.CoverTab[127299]++
												yaml_parser_set_scanner_error(parser, "while scanning a tag",
				start_mark, "did not find the expected '>'")
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1864
			// _ = "end of CoverTab[127299]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1865
			_go_fuzz_dep_.CoverTab[127300]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1865
			// _ = "end of CoverTab[127300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1865
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1865
		// _ = "end of CoverTab[127295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1865
		_go_fuzz_dep_.CoverTab[127296]++

											skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1867
		// _ = "end of CoverTab[127296]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1868
		_go_fuzz_dep_.CoverTab[127301]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1872
		if !yaml_parser_scan_tag_handle(parser, false, start_mark, &handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1872
			_go_fuzz_dep_.CoverTab[127303]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1873
			// _ = "end of CoverTab[127303]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1874
			_go_fuzz_dep_.CoverTab[127304]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1874
			// _ = "end of CoverTab[127304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1874
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1874
		// _ = "end of CoverTab[127301]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1874
		_go_fuzz_dep_.CoverTab[127302]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
		if handle[0] == '!' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			_go_fuzz_dep_.CoverTab[127305]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			return len(handle) > 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			// _ = "end of CoverTab[127305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			_go_fuzz_dep_.CoverTab[127306]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			return handle[len(handle)-1] == '!'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			// _ = "end of CoverTab[127306]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1877
			_go_fuzz_dep_.CoverTab[127307]++

												if !yaml_parser_scan_tag_uri(parser, false, nil, start_mark, &suffix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1879
				_go_fuzz_dep_.CoverTab[127308]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1880
				// _ = "end of CoverTab[127308]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1881
				_go_fuzz_dep_.CoverTab[127309]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1881
				// _ = "end of CoverTab[127309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1881
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1881
			// _ = "end of CoverTab[127307]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1882
			_go_fuzz_dep_.CoverTab[127310]++

												if !yaml_parser_scan_tag_uri(parser, false, handle, start_mark, &suffix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1884
				_go_fuzz_dep_.CoverTab[127312]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1885
				// _ = "end of CoverTab[127312]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1886
				_go_fuzz_dep_.CoverTab[127313]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1886
				// _ = "end of CoverTab[127313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1886
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1886
			// _ = "end of CoverTab[127310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1886
			_go_fuzz_dep_.CoverTab[127311]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1889
			handle = []byte{'!'}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1893
			if len(suffix) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1893
				_go_fuzz_dep_.CoverTab[127314]++
													handle, suffix = suffix, handle
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1894
				// _ = "end of CoverTab[127314]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1895
				_go_fuzz_dep_.CoverTab[127315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1895
				// _ = "end of CoverTab[127315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1895
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1895
			// _ = "end of CoverTab[127311]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1896
		// _ = "end of CoverTab[127302]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1897
	// _ = "end of CoverTab[127287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1897
	_go_fuzz_dep_.CoverTab[127288]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
		_go_fuzz_dep_.CoverTab[127316]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
		// _ = "end of CoverTab[127316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1900
		_go_fuzz_dep_.CoverTab[127317]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1901
		// _ = "end of CoverTab[127317]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1902
		_go_fuzz_dep_.CoverTab[127318]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1902
		// _ = "end of CoverTab[127318]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1902
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1902
	// _ = "end of CoverTab[127288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1902
	_go_fuzz_dep_.CoverTab[127289]++
										if !is_blankz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1903
		_go_fuzz_dep_.CoverTab[127319]++
											yaml_parser_set_scanner_error(parser, "while scanning a tag",
			start_mark, "did not find expected whitespace or line break")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1906
		// _ = "end of CoverTab[127319]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1907
		_go_fuzz_dep_.CoverTab[127320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1907
		// _ = "end of CoverTab[127320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1907
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1907
	// _ = "end of CoverTab[127289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1907
	_go_fuzz_dep_.CoverTab[127290]++

										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1912
	*token = yaml_token_t{
		typ:		yaml_TAG_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
		value:		handle,
		suffix:		suffix,
	}
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1919
	// _ = "end of CoverTab[127290]"
}

// Scan a tag handle.
func yaml_parser_scan_tag_handle(parser *yaml_parser_t, directive bool, start_mark yaml_mark_t, handle *[]byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1923
	_go_fuzz_dep_.CoverTab[127321]++

										if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1925
		_go_fuzz_dep_.CoverTab[127327]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1925
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1925
		// _ = "end of CoverTab[127327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1925
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1925
		_go_fuzz_dep_.CoverTab[127328]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1926
		// _ = "end of CoverTab[127328]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1927
		_go_fuzz_dep_.CoverTab[127329]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1927
		// _ = "end of CoverTab[127329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1927
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1927
	// _ = "end of CoverTab[127321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1927
	_go_fuzz_dep_.CoverTab[127322]++
										if parser.buffer[parser.buffer_pos] != '!' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1928
		_go_fuzz_dep_.CoverTab[127330]++
											yaml_parser_set_scanner_tag_error(parser, directive,
			start_mark, "did not find expected '!'")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1931
		// _ = "end of CoverTab[127330]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1932
		_go_fuzz_dep_.CoverTab[127331]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1932
		// _ = "end of CoverTab[127331]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1932
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1932
	// _ = "end of CoverTab[127322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1932
	_go_fuzz_dep_.CoverTab[127323]++

										var s []byte

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1937
	s = read(parser, s)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
		_go_fuzz_dep_.CoverTab[127332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
		// _ = "end of CoverTab[127332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1940
		_go_fuzz_dep_.CoverTab[127333]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1941
		// _ = "end of CoverTab[127333]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1942
		_go_fuzz_dep_.CoverTab[127334]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1942
		// _ = "end of CoverTab[127334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1942
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1942
	// _ = "end of CoverTab[127323]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1942
	_go_fuzz_dep_.CoverTab[127324]++
										for is_alpha(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1943
		_go_fuzz_dep_.CoverTab[127335]++
											s = read(parser, s)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1945
			_go_fuzz_dep_.CoverTab[127336]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1945
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1945
			// _ = "end of CoverTab[127336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1945
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1945
			_go_fuzz_dep_.CoverTab[127337]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1946
			// _ = "end of CoverTab[127337]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1947
			_go_fuzz_dep_.CoverTab[127338]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1947
			// _ = "end of CoverTab[127338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1947
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1947
		// _ = "end of CoverTab[127335]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1948
	// _ = "end of CoverTab[127324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1948
	_go_fuzz_dep_.CoverTab[127325]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1951
	if parser.buffer[parser.buffer_pos] == '!' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1951
		_go_fuzz_dep_.CoverTab[127339]++
											s = read(parser, s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1952
		// _ = "end of CoverTab[127339]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1953
		_go_fuzz_dep_.CoverTab[127340]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
		if directive && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
			_go_fuzz_dep_.CoverTab[127341]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
			return string(s) != "!"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
			// _ = "end of CoverTab[127341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1956
			_go_fuzz_dep_.CoverTab[127342]++
												yaml_parser_set_scanner_tag_error(parser, directive,
				start_mark, "did not find expected '!'")
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1959
			// _ = "end of CoverTab[127342]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1960
			_go_fuzz_dep_.CoverTab[127343]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1960
			// _ = "end of CoverTab[127343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1960
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1960
		// _ = "end of CoverTab[127340]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1961
	// _ = "end of CoverTab[127325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1961
	_go_fuzz_dep_.CoverTab[127326]++

										*handle = s
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1964
	// _ = "end of CoverTab[127326]"
}

// Scan a tag.
func yaml_parser_scan_tag_uri(parser *yaml_parser_t, directive bool, head []byte, start_mark yaml_mark_t, uri *[]byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1968
	_go_fuzz_dep_.CoverTab[127344]++
										//size_t length = head ? strlen((char *)head) : 0
										var s []byte
										hasTag := len(head) > 0

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1976
	if len(head) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1976
		_go_fuzz_dep_.CoverTab[127349]++
											s = append(s, head[1:]...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1977
		// _ = "end of CoverTab[127349]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1978
		_go_fuzz_dep_.CoverTab[127350]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1978
		// _ = "end of CoverTab[127350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1978
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1978
	// _ = "end of CoverTab[127344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1978
	_go_fuzz_dep_.CoverTab[127345]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
		_go_fuzz_dep_.CoverTab[127351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
		// _ = "end of CoverTab[127351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1981
		_go_fuzz_dep_.CoverTab[127352]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1982
		// _ = "end of CoverTab[127352]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1983
		_go_fuzz_dep_.CoverTab[127353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1983
		// _ = "end of CoverTab[127353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1983
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1983
	// _ = "end of CoverTab[127345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1983
	_go_fuzz_dep_.CoverTab[127346]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
	for is_alpha(parser.buffer, parser.buffer_pos) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
		_go_fuzz_dep_.CoverTab[127354]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
		return parser.buffer[parser.buffer_pos] == ';'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
		// _ = "end of CoverTab[127354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
		_go_fuzz_dep_.CoverTab[127355]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1991
		return parser.buffer[parser.buffer_pos] == '/'
											// _ = "end of CoverTab[127355]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
		_go_fuzz_dep_.CoverTab[127356]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
		return parser.buffer[parser.buffer_pos] == '?'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
		// _ = "end of CoverTab[127356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
		_go_fuzz_dep_.CoverTab[127357]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1992
		return parser.buffer[parser.buffer_pos] == ':'
											// _ = "end of CoverTab[127357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
		_go_fuzz_dep_.CoverTab[127358]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
		return parser.buffer[parser.buffer_pos] == '@'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
		// _ = "end of CoverTab[127358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
		_go_fuzz_dep_.CoverTab[127359]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1993
		return parser.buffer[parser.buffer_pos] == '&'
											// _ = "end of CoverTab[127359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
		_go_fuzz_dep_.CoverTab[127360]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
		return parser.buffer[parser.buffer_pos] == '='
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
		// _ = "end of CoverTab[127360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
		_go_fuzz_dep_.CoverTab[127361]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1994
		return parser.buffer[parser.buffer_pos] == '+'
											// _ = "end of CoverTab[127361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
		_go_fuzz_dep_.CoverTab[127362]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
		return parser.buffer[parser.buffer_pos] == '$'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
		// _ = "end of CoverTab[127362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
		_go_fuzz_dep_.CoverTab[127363]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1995
		return parser.buffer[parser.buffer_pos] == ','
											// _ = "end of CoverTab[127363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
		_go_fuzz_dep_.CoverTab[127364]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
		return parser.buffer[parser.buffer_pos] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
		// _ = "end of CoverTab[127364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
		_go_fuzz_dep_.CoverTab[127365]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1996
		return parser.buffer[parser.buffer_pos] == '!'
											// _ = "end of CoverTab[127365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
		_go_fuzz_dep_.CoverTab[127366]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
		return parser.buffer[parser.buffer_pos] == '~'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
		// _ = "end of CoverTab[127366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
		_go_fuzz_dep_.CoverTab[127367]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1997
		return parser.buffer[parser.buffer_pos] == '*'
											// _ = "end of CoverTab[127367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
		_go_fuzz_dep_.CoverTab[127368]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
		return parser.buffer[parser.buffer_pos] == '\''
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
		// _ = "end of CoverTab[127368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
		_go_fuzz_dep_.CoverTab[127369]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1998
		return parser.buffer[parser.buffer_pos] == '('
											// _ = "end of CoverTab[127369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
		_go_fuzz_dep_.CoverTab[127370]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
		return parser.buffer[parser.buffer_pos] == ')'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
		// _ = "end of CoverTab[127370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
		_go_fuzz_dep_.CoverTab[127371]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:1999
		return parser.buffer[parser.buffer_pos] == '['
											// _ = "end of CoverTab[127371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
		_go_fuzz_dep_.CoverTab[127372]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
		return parser.buffer[parser.buffer_pos] == ']'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
		// _ = "end of CoverTab[127372]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
		_go_fuzz_dep_.CoverTab[127373]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2000
		return parser.buffer[parser.buffer_pos] == '%'
											// _ = "end of CoverTab[127373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2001
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2001
		_go_fuzz_dep_.CoverTab[127374]++

											if parser.buffer[parser.buffer_pos] == '%' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2003
			_go_fuzz_dep_.CoverTab[127377]++
												if !yaml_parser_scan_uri_escapes(parser, directive, start_mark, &s) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2004
				_go_fuzz_dep_.CoverTab[127378]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2005
				// _ = "end of CoverTab[127378]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2006
				_go_fuzz_dep_.CoverTab[127379]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2006
				// _ = "end of CoverTab[127379]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2006
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2006
			// _ = "end of CoverTab[127377]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2007
			_go_fuzz_dep_.CoverTab[127380]++
												s = read(parser, s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2008
			// _ = "end of CoverTab[127380]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2009
		// _ = "end of CoverTab[127374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2009
		_go_fuzz_dep_.CoverTab[127375]++
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2010
			_go_fuzz_dep_.CoverTab[127381]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2010
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2010
			// _ = "end of CoverTab[127381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2010
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2010
			_go_fuzz_dep_.CoverTab[127382]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2011
			// _ = "end of CoverTab[127382]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2012
			_go_fuzz_dep_.CoverTab[127383]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2012
			// _ = "end of CoverTab[127383]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2012
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2012
		// _ = "end of CoverTab[127375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2012
		_go_fuzz_dep_.CoverTab[127376]++
											hasTag = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2013
		// _ = "end of CoverTab[127376]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2014
	// _ = "end of CoverTab[127346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2014
	_go_fuzz_dep_.CoverTab[127347]++

										if !hasTag {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2016
		_go_fuzz_dep_.CoverTab[127384]++
											yaml_parser_set_scanner_tag_error(parser, directive,
			start_mark, "did not find expected tag URI")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2019
		// _ = "end of CoverTab[127384]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2020
		_go_fuzz_dep_.CoverTab[127385]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2020
		// _ = "end of CoverTab[127385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2020
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2020
	// _ = "end of CoverTab[127347]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2020
	_go_fuzz_dep_.CoverTab[127348]++
										*uri = s
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2022
	// _ = "end of CoverTab[127348]"
}

// Decode an URI-escape sequence corresponding to a single UTF-8 character.
func yaml_parser_scan_uri_escapes(parser *yaml_parser_t, directive bool, start_mark yaml_mark_t, s *[]byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2026
	_go_fuzz_dep_.CoverTab[127386]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2029
	w := 1024
	for w > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2030
		_go_fuzz_dep_.CoverTab[127388]++

											if parser.unread < 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2032
			_go_fuzz_dep_.CoverTab[127392]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2032
			return !yaml_parser_update_buffer(parser, 3)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2032
			// _ = "end of CoverTab[127392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2032
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2032
			_go_fuzz_dep_.CoverTab[127393]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2033
			// _ = "end of CoverTab[127393]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2034
			_go_fuzz_dep_.CoverTab[127394]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2034
			// _ = "end of CoverTab[127394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2034
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2034
		// _ = "end of CoverTab[127388]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2034
		_go_fuzz_dep_.CoverTab[127389]++

											if !(parser.buffer[parser.buffer_pos] == '%' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2036
			_go_fuzz_dep_.CoverTab[127395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2036
			return is_hex(parser.buffer, parser.buffer_pos+1)
												// _ = "end of CoverTab[127395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2037
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2037
			_go_fuzz_dep_.CoverTab[127396]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2037
			return is_hex(parser.buffer, parser.buffer_pos+2)
												// _ = "end of CoverTab[127396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2038
		}()) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2038
			_go_fuzz_dep_.CoverTab[127397]++
												return yaml_parser_set_scanner_tag_error(parser, directive,
				start_mark, "did not find URI escaped octet")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2040
			// _ = "end of CoverTab[127397]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2041
			_go_fuzz_dep_.CoverTab[127398]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2041
			// _ = "end of CoverTab[127398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2041
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2041
		// _ = "end of CoverTab[127389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2041
		_go_fuzz_dep_.CoverTab[127390]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2044
		octet := byte((as_hex(parser.buffer, parser.buffer_pos+1) << 4) + as_hex(parser.buffer, parser.buffer_pos+2))

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2047
		if w == 1024 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2047
			_go_fuzz_dep_.CoverTab[127399]++
												w = width(octet)
												if w == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2049
				_go_fuzz_dep_.CoverTab[127400]++
													return yaml_parser_set_scanner_tag_error(parser, directive,
					start_mark, "found an incorrect leading UTF-8 octet")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2051
				// _ = "end of CoverTab[127400]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2052
				_go_fuzz_dep_.CoverTab[127401]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2052
				// _ = "end of CoverTab[127401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2052
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2052
			// _ = "end of CoverTab[127399]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2053
			_go_fuzz_dep_.CoverTab[127402]++

												if octet&0xC0 != 0x80 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2055
				_go_fuzz_dep_.CoverTab[127403]++
													return yaml_parser_set_scanner_tag_error(parser, directive,
					start_mark, "found an incorrect trailing UTF-8 octet")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2057
				// _ = "end of CoverTab[127403]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2058
				_go_fuzz_dep_.CoverTab[127404]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2058
				// _ = "end of CoverTab[127404]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2058
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2058
			// _ = "end of CoverTab[127402]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2059
		// _ = "end of CoverTab[127390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2059
		_go_fuzz_dep_.CoverTab[127391]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2062
		*s = append(*s, octet)
											skip(parser)
											skip(parser)
											skip(parser)
											w--
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2066
		// _ = "end of CoverTab[127391]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2067
	// _ = "end of CoverTab[127386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2067
	_go_fuzz_dep_.CoverTab[127387]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2068
	// _ = "end of CoverTab[127387]"
}

// Scan a block scalar.
func yaml_parser_scan_block_scalar(parser *yaml_parser_t, token *yaml_token_t, literal bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2072
	_go_fuzz_dep_.CoverTab[127405]++

										start_mark := parser.mark
										skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
		_go_fuzz_dep_.CoverTab[127420]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
		// _ = "end of CoverTab[127420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2078
		_go_fuzz_dep_.CoverTab[127421]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2079
		// _ = "end of CoverTab[127421]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2080
		_go_fuzz_dep_.CoverTab[127422]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2080
		// _ = "end of CoverTab[127422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2080
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2080
	// _ = "end of CoverTab[127405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2080
	_go_fuzz_dep_.CoverTab[127406]++

	// Check for a chomping indicator.
	var chomping, increment int
	if parser.buffer[parser.buffer_pos] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2084
		_go_fuzz_dep_.CoverTab[127423]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2084
		return parser.buffer[parser.buffer_pos] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2084
		// _ = "end of CoverTab[127423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2084
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2084
		_go_fuzz_dep_.CoverTab[127424]++

											if parser.buffer[parser.buffer_pos] == '+' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2086
			_go_fuzz_dep_.CoverTab[127427]++
												chomping = +1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2087
			// _ = "end of CoverTab[127427]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2088
			_go_fuzz_dep_.CoverTab[127428]++
												chomping = -1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2089
			// _ = "end of CoverTab[127428]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2090
		// _ = "end of CoverTab[127424]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2090
		_go_fuzz_dep_.CoverTab[127425]++
											skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
		if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
			_go_fuzz_dep_.CoverTab[127429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
			// _ = "end of CoverTab[127429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2094
			_go_fuzz_dep_.CoverTab[127430]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2095
			// _ = "end of CoverTab[127430]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2096
			_go_fuzz_dep_.CoverTab[127431]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2096
			// _ = "end of CoverTab[127431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2096
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2096
		// _ = "end of CoverTab[127425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2096
		_go_fuzz_dep_.CoverTab[127426]++
											if is_digit(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2097
			_go_fuzz_dep_.CoverTab[127432]++

												if parser.buffer[parser.buffer_pos] == '0' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2099
				_go_fuzz_dep_.CoverTab[127434]++
													yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
					start_mark, "found an indentation indicator equal to 0")
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2102
				// _ = "end of CoverTab[127434]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2103
				_go_fuzz_dep_.CoverTab[127435]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2103
				// _ = "end of CoverTab[127435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2103
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2103
			// _ = "end of CoverTab[127432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2103
			_go_fuzz_dep_.CoverTab[127433]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2106
			increment = as_digit(parser.buffer, parser.buffer_pos)
												skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2107
			// _ = "end of CoverTab[127433]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2108
			_go_fuzz_dep_.CoverTab[127436]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2108
			// _ = "end of CoverTab[127436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2108
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2108
		// _ = "end of CoverTab[127426]"

	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2110
		_go_fuzz_dep_.CoverTab[127437]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2110
		if is_digit(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2110
			_go_fuzz_dep_.CoverTab[127438]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2113
			if parser.buffer[parser.buffer_pos] == '0' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2113
				_go_fuzz_dep_.CoverTab[127441]++
													yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
					start_mark, "found an indentation indicator equal to 0")
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2116
				// _ = "end of CoverTab[127441]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2117
				_go_fuzz_dep_.CoverTab[127442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2117
				// _ = "end of CoverTab[127442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2117
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2117
			// _ = "end of CoverTab[127438]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2117
			_go_fuzz_dep_.CoverTab[127439]++
												increment = as_digit(parser.buffer, parser.buffer_pos)
												skip(parser)

												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2121
				_go_fuzz_dep_.CoverTab[127443]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2121
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2121
				// _ = "end of CoverTab[127443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2121
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2121
				_go_fuzz_dep_.CoverTab[127444]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2122
				// _ = "end of CoverTab[127444]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2123
				_go_fuzz_dep_.CoverTab[127445]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2123
				// _ = "end of CoverTab[127445]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2123
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2123
			// _ = "end of CoverTab[127439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2123
			_go_fuzz_dep_.CoverTab[127440]++
												if parser.buffer[parser.buffer_pos] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2124
				_go_fuzz_dep_.CoverTab[127446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2124
				return parser.buffer[parser.buffer_pos] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2124
				// _ = "end of CoverTab[127446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2124
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2124
				_go_fuzz_dep_.CoverTab[127447]++
													if parser.buffer[parser.buffer_pos] == '+' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2125
					_go_fuzz_dep_.CoverTab[127449]++
														chomping = +1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2126
					// _ = "end of CoverTab[127449]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2127
					_go_fuzz_dep_.CoverTab[127450]++
														chomping = -1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2128
					// _ = "end of CoverTab[127450]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2129
				// _ = "end of CoverTab[127447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2129
				_go_fuzz_dep_.CoverTab[127448]++
													skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2130
				// _ = "end of CoverTab[127448]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2131
				_go_fuzz_dep_.CoverTab[127451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2131
				// _ = "end of CoverTab[127451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2131
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2131
			// _ = "end of CoverTab[127440]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
			_go_fuzz_dep_.CoverTab[127452]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
			// _ = "end of CoverTab[127452]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
		// _ = "end of CoverTab[127437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
	// _ = "end of CoverTab[127406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2132
	_go_fuzz_dep_.CoverTab[127407]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
		_go_fuzz_dep_.CoverTab[127453]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
		// _ = "end of CoverTab[127453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2135
		_go_fuzz_dep_.CoverTab[127454]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2136
		// _ = "end of CoverTab[127454]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2137
		_go_fuzz_dep_.CoverTab[127455]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2137
		// _ = "end of CoverTab[127455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2137
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2137
	// _ = "end of CoverTab[127407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2137
	_go_fuzz_dep_.CoverTab[127408]++
										for is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2138
		_go_fuzz_dep_.CoverTab[127456]++
											skip(parser)
											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2140
			_go_fuzz_dep_.CoverTab[127457]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2140
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2140
			// _ = "end of CoverTab[127457]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2140
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2140
			_go_fuzz_dep_.CoverTab[127458]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2141
			// _ = "end of CoverTab[127458]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2142
			_go_fuzz_dep_.CoverTab[127459]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2142
			// _ = "end of CoverTab[127459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2142
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2142
		// _ = "end of CoverTab[127456]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2143
	// _ = "end of CoverTab[127408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2143
	_go_fuzz_dep_.CoverTab[127409]++
										if parser.buffer[parser.buffer_pos] == '#' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2144
		_go_fuzz_dep_.CoverTab[127460]++
											for !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2145
			_go_fuzz_dep_.CoverTab[127461]++
												skip(parser)
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2147
				_go_fuzz_dep_.CoverTab[127462]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2147
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2147
				// _ = "end of CoverTab[127462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2147
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2147
				_go_fuzz_dep_.CoverTab[127463]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2148
				// _ = "end of CoverTab[127463]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2149
				_go_fuzz_dep_.CoverTab[127464]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2149
				// _ = "end of CoverTab[127464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2149
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2149
			// _ = "end of CoverTab[127461]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2150
		// _ = "end of CoverTab[127460]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2151
		_go_fuzz_dep_.CoverTab[127465]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2151
		// _ = "end of CoverTab[127465]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2151
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2151
	// _ = "end of CoverTab[127409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2151
	_go_fuzz_dep_.CoverTab[127410]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2154
	if !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2154
		_go_fuzz_dep_.CoverTab[127466]++
											yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
			start_mark, "did not find expected comment or line break")
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2157
		// _ = "end of CoverTab[127466]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2158
		_go_fuzz_dep_.CoverTab[127467]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2158
		// _ = "end of CoverTab[127467]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2158
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2158
	// _ = "end of CoverTab[127410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2158
	_go_fuzz_dep_.CoverTab[127411]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2161
	if is_break(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2161
		_go_fuzz_dep_.CoverTab[127468]++
											if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2162
			_go_fuzz_dep_.CoverTab[127470]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2162
			return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2162
			// _ = "end of CoverTab[127470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2162
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2162
			_go_fuzz_dep_.CoverTab[127471]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2163
			// _ = "end of CoverTab[127471]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2164
			_go_fuzz_dep_.CoverTab[127472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2164
			// _ = "end of CoverTab[127472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2164
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2164
		// _ = "end of CoverTab[127468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2164
		_go_fuzz_dep_.CoverTab[127469]++
											skip_line(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2165
		// _ = "end of CoverTab[127469]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2166
		_go_fuzz_dep_.CoverTab[127473]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2166
		// _ = "end of CoverTab[127473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2166
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2166
	// _ = "end of CoverTab[127411]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2166
	_go_fuzz_dep_.CoverTab[127412]++

										end_mark := parser.mark

	// Set the indentation level if it was specified.
	var indent int
	if increment > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2172
		_go_fuzz_dep_.CoverTab[127474]++
											if parser.indent >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2173
			_go_fuzz_dep_.CoverTab[127475]++
												indent = parser.indent + increment
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2174
			// _ = "end of CoverTab[127475]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2175
			_go_fuzz_dep_.CoverTab[127476]++
												indent = increment
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2176
			// _ = "end of CoverTab[127476]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2177
		// _ = "end of CoverTab[127474]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2178
		_go_fuzz_dep_.CoverTab[127477]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2178
		// _ = "end of CoverTab[127477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2178
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2178
	// _ = "end of CoverTab[127412]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2178
	_go_fuzz_dep_.CoverTab[127413]++

	// Scan the leading line breaks and determine the indentation level if needed.
	var s, leading_break, trailing_breaks []byte
	if !yaml_parser_scan_block_scalar_breaks(parser, &indent, &trailing_breaks, start_mark, &end_mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2182
		_go_fuzz_dep_.CoverTab[127478]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2183
		// _ = "end of CoverTab[127478]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2184
		_go_fuzz_dep_.CoverTab[127479]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2184
		// _ = "end of CoverTab[127479]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2184
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2184
	// _ = "end of CoverTab[127413]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2184
	_go_fuzz_dep_.CoverTab[127414]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
	if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
		_go_fuzz_dep_.CoverTab[127480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
		return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
		// _ = "end of CoverTab[127480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2187
		_go_fuzz_dep_.CoverTab[127481]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2188
		// _ = "end of CoverTab[127481]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2189
		_go_fuzz_dep_.CoverTab[127482]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2189
		// _ = "end of CoverTab[127482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2189
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2189
	// _ = "end of CoverTab[127414]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2189
	_go_fuzz_dep_.CoverTab[127415]++
										var leading_blank, trailing_blank bool
										for parser.mark.column == indent && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2191
		_go_fuzz_dep_.CoverTab[127483]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2191
		return !is_z(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2191
		// _ = "end of CoverTab[127483]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2191
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2191
		_go_fuzz_dep_.CoverTab[127484]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2195
		trailing_blank = is_blank(parser.buffer, parser.buffer_pos)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
		if !literal && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			_go_fuzz_dep_.CoverTab[127488]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			return !leading_blank
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			// _ = "end of CoverTab[127488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			_go_fuzz_dep_.CoverTab[127489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			return !trailing_blank
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			// _ = "end of CoverTab[127489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			_go_fuzz_dep_.CoverTab[127490]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			return len(leading_break) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			// _ = "end of CoverTab[127490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			_go_fuzz_dep_.CoverTab[127491]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			return leading_break[0] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			// _ = "end of CoverTab[127491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2198
			_go_fuzz_dep_.CoverTab[127492]++

												if len(trailing_breaks) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2200
				_go_fuzz_dep_.CoverTab[127493]++
													s = append(s, ' ')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2201
				// _ = "end of CoverTab[127493]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2202
				_go_fuzz_dep_.CoverTab[127494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2202
				// _ = "end of CoverTab[127494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2202
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2202
			// _ = "end of CoverTab[127492]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2203
			_go_fuzz_dep_.CoverTab[127495]++
												s = append(s, leading_break...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2204
			// _ = "end of CoverTab[127495]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2205
		// _ = "end of CoverTab[127484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2205
		_go_fuzz_dep_.CoverTab[127485]++
											leading_break = leading_break[:0]

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2209
		s = append(s, trailing_breaks...)
											trailing_breaks = trailing_breaks[:0]

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2213
		leading_blank = is_blank(parser.buffer, parser.buffer_pos)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2216
		for !is_breakz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2216
			_go_fuzz_dep_.CoverTab[127496]++
												s = read(parser, s)
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2218
				_go_fuzz_dep_.CoverTab[127497]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2218
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2218
				// _ = "end of CoverTab[127497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2218
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2218
				_go_fuzz_dep_.CoverTab[127498]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2219
				// _ = "end of CoverTab[127498]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2220
				_go_fuzz_dep_.CoverTab[127499]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2220
				// _ = "end of CoverTab[127499]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2220
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2220
			// _ = "end of CoverTab[127496]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2221
		// _ = "end of CoverTab[127485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2221
		_go_fuzz_dep_.CoverTab[127486]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
		if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
			_go_fuzz_dep_.CoverTab[127500]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
			return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
			// _ = "end of CoverTab[127500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2224
			_go_fuzz_dep_.CoverTab[127501]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2225
			// _ = "end of CoverTab[127501]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2226
			_go_fuzz_dep_.CoverTab[127502]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2226
			// _ = "end of CoverTab[127502]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2226
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2226
		// _ = "end of CoverTab[127486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2226
		_go_fuzz_dep_.CoverTab[127487]++

											leading_break = read_line(parser, leading_break)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2231
		if !yaml_parser_scan_block_scalar_breaks(parser, &indent, &trailing_breaks, start_mark, &end_mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2231
			_go_fuzz_dep_.CoverTab[127503]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2232
			// _ = "end of CoverTab[127503]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2233
			_go_fuzz_dep_.CoverTab[127504]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2233
			// _ = "end of CoverTab[127504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2233
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2233
		// _ = "end of CoverTab[127487]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2234
	// _ = "end of CoverTab[127415]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2234
	_go_fuzz_dep_.CoverTab[127416]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2237
	if chomping != -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2237
		_go_fuzz_dep_.CoverTab[127505]++
											s = append(s, leading_break...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2238
		// _ = "end of CoverTab[127505]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2239
		_go_fuzz_dep_.CoverTab[127506]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2239
		// _ = "end of CoverTab[127506]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2239
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2239
	// _ = "end of CoverTab[127416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2239
	_go_fuzz_dep_.CoverTab[127417]++
										if chomping == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2240
		_go_fuzz_dep_.CoverTab[127507]++
											s = append(s, trailing_breaks...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2241
		// _ = "end of CoverTab[127507]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2242
		_go_fuzz_dep_.CoverTab[127508]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2242
		// _ = "end of CoverTab[127508]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2242
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2242
	// _ = "end of CoverTab[127417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2242
	_go_fuzz_dep_.CoverTab[127418]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2245
	*token = yaml_token_t{
		typ:		yaml_SCALAR_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
		value:		s,
		style:		yaml_LITERAL_SCALAR_STYLE,
	}
	if !literal {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2252
		_go_fuzz_dep_.CoverTab[127509]++
											token.style = yaml_FOLDED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2253
		// _ = "end of CoverTab[127509]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2254
		_go_fuzz_dep_.CoverTab[127510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2254
		// _ = "end of CoverTab[127510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2254
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2254
	// _ = "end of CoverTab[127418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2254
	_go_fuzz_dep_.CoverTab[127419]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2255
	// _ = "end of CoverTab[127419]"
}

// Scan indentation spaces and line breaks for a block scalar.  Determine the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2258
// indentation level if needed.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2260
func yaml_parser_scan_block_scalar_breaks(parser *yaml_parser_t, indent *int, breaks *[]byte, start_mark yaml_mark_t, end_mark *yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2260
	_go_fuzz_dep_.CoverTab[127511]++
										*end_mark = parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2264
	max_indent := 0
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2265
		_go_fuzz_dep_.CoverTab[127514]++

											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2267
			_go_fuzz_dep_.CoverTab[127521]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2267
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2267
			// _ = "end of CoverTab[127521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2267
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2267
			_go_fuzz_dep_.CoverTab[127522]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2268
			// _ = "end of CoverTab[127522]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2269
			_go_fuzz_dep_.CoverTab[127523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2269
			// _ = "end of CoverTab[127523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2269
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2269
		// _ = "end of CoverTab[127514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2269
		_go_fuzz_dep_.CoverTab[127515]++
											for (*indent == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			_go_fuzz_dep_.CoverTab[127524]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			return parser.mark.column < *indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			// _ = "end of CoverTab[127524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			_go_fuzz_dep_.CoverTab[127525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			return is_space(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			// _ = "end of CoverTab[127525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2270
			_go_fuzz_dep_.CoverTab[127526]++
												skip(parser)
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2272
				_go_fuzz_dep_.CoverTab[127527]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2272
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2272
				// _ = "end of CoverTab[127527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2272
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2272
				_go_fuzz_dep_.CoverTab[127528]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2273
				// _ = "end of CoverTab[127528]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2274
				_go_fuzz_dep_.CoverTab[127529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2274
				// _ = "end of CoverTab[127529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2274
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2274
			// _ = "end of CoverTab[127526]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2275
		// _ = "end of CoverTab[127515]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2275
		_go_fuzz_dep_.CoverTab[127516]++
											if parser.mark.column > max_indent {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2276
			_go_fuzz_dep_.CoverTab[127530]++
												max_indent = parser.mark.column
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2277
			// _ = "end of CoverTab[127530]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2278
			_go_fuzz_dep_.CoverTab[127531]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2278
			// _ = "end of CoverTab[127531]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2278
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2278
		// _ = "end of CoverTab[127516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2278
		_go_fuzz_dep_.CoverTab[127517]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
		if (*indent == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			_go_fuzz_dep_.CoverTab[127532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			return parser.mark.column < *indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			// _ = "end of CoverTab[127532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			_go_fuzz_dep_.CoverTab[127533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			return is_tab(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			// _ = "end of CoverTab[127533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2281
			_go_fuzz_dep_.CoverTab[127534]++
												return yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
				start_mark, "found a tab character where an indentation space is expected")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2283
			// _ = "end of CoverTab[127534]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2284
			_go_fuzz_dep_.CoverTab[127535]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2284
			// _ = "end of CoverTab[127535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2284
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2284
		// _ = "end of CoverTab[127517]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2284
		_go_fuzz_dep_.CoverTab[127518]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2287
		if !is_break(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2287
			_go_fuzz_dep_.CoverTab[127536]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2288
			// _ = "end of CoverTab[127536]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2289
			_go_fuzz_dep_.CoverTab[127537]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2289
			// _ = "end of CoverTab[127537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2289
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2289
		// _ = "end of CoverTab[127518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2289
		_go_fuzz_dep_.CoverTab[127519]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
		if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
			_go_fuzz_dep_.CoverTab[127538]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
			return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
			// _ = "end of CoverTab[127538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2292
			_go_fuzz_dep_.CoverTab[127539]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2293
			// _ = "end of CoverTab[127539]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2294
			_go_fuzz_dep_.CoverTab[127540]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2294
			// _ = "end of CoverTab[127540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2294
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2294
		// _ = "end of CoverTab[127519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2294
		_go_fuzz_dep_.CoverTab[127520]++

											*breaks = read_line(parser, *breaks)
											*end_mark = parser.mark
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2297
		// _ = "end of CoverTab[127520]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2298
	// _ = "end of CoverTab[127511]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2298
	_go_fuzz_dep_.CoverTab[127512]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2301
	if *indent == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2301
		_go_fuzz_dep_.CoverTab[127541]++
											*indent = max_indent
											if *indent < parser.indent+1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2303
			_go_fuzz_dep_.CoverTab[127543]++
												*indent = parser.indent + 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2304
			// _ = "end of CoverTab[127543]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2305
			_go_fuzz_dep_.CoverTab[127544]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2305
			// _ = "end of CoverTab[127544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2305
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2305
		// _ = "end of CoverTab[127541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2305
		_go_fuzz_dep_.CoverTab[127542]++
											if *indent < 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2306
			_go_fuzz_dep_.CoverTab[127545]++
												*indent = 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2307
			// _ = "end of CoverTab[127545]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2308
			_go_fuzz_dep_.CoverTab[127546]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2308
			// _ = "end of CoverTab[127546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2308
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2308
		// _ = "end of CoverTab[127542]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2309
		_go_fuzz_dep_.CoverTab[127547]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2309
		// _ = "end of CoverTab[127547]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2309
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2309
	// _ = "end of CoverTab[127512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2309
	_go_fuzz_dep_.CoverTab[127513]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2310
	// _ = "end of CoverTab[127513]"
}

// Scan a quoted scalar.
func yaml_parser_scan_flow_scalar(parser *yaml_parser_t, token *yaml_token_t, single bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2314
	_go_fuzz_dep_.CoverTab[127548]++

										start_mark := parser.mark
										skip(parser)

	// Consume the content of the quoted scalar.
	var s, leading_break, trailing_breaks, whitespaces []byte
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2321
		_go_fuzz_dep_.CoverTab[127551]++

											if parser.unread < 4 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2323
			_go_fuzz_dep_.CoverTab[127559]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2323
			return !yaml_parser_update_buffer(parser, 4)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2323
			// _ = "end of CoverTab[127559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2323
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2323
			_go_fuzz_dep_.CoverTab[127560]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2324
			// _ = "end of CoverTab[127560]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2325
			_go_fuzz_dep_.CoverTab[127561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2325
			// _ = "end of CoverTab[127561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2325
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2325
		// _ = "end of CoverTab[127551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2325
		_go_fuzz_dep_.CoverTab[127552]++

											if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2327
			_go_fuzz_dep_.CoverTab[127562]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2327
			return ((parser.buffer[parser.buffer_pos+0] == '-' && func() bool {
													_go_fuzz_dep_.CoverTab[127563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2328
				return parser.buffer[parser.buffer_pos+1] == '-'
													// _ = "end of CoverTab[127563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2329
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2329
				_go_fuzz_dep_.CoverTab[127564]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2329
				return parser.buffer[parser.buffer_pos+2] == '-'
													// _ = "end of CoverTab[127564]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2330
			}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2330
				_go_fuzz_dep_.CoverTab[127565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2330
				return (parser.buffer[parser.buffer_pos+0] == '.' && func() bool {
														_go_fuzz_dep_.CoverTab[127566]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2331
					return parser.buffer[parser.buffer_pos+1] == '.'
														// _ = "end of CoverTab[127566]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2332
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2332
					_go_fuzz_dep_.CoverTab[127567]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2332
					return parser.buffer[parser.buffer_pos+2] == '.'
														// _ = "end of CoverTab[127567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
				// _ = "end of CoverTab[127565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
			// _ = "end of CoverTab[127562]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
			_go_fuzz_dep_.CoverTab[127568]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2333
			return is_blankz(parser.buffer, parser.buffer_pos+3)
												// _ = "end of CoverTab[127568]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2334
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2334
			_go_fuzz_dep_.CoverTab[127569]++
												yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected document indicator")
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2337
			// _ = "end of CoverTab[127569]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2338
			_go_fuzz_dep_.CoverTab[127570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2338
			// _ = "end of CoverTab[127570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2338
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2338
		// _ = "end of CoverTab[127552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2338
		_go_fuzz_dep_.CoverTab[127553]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2341
		if is_z(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2341
			_go_fuzz_dep_.CoverTab[127571]++
												yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected end of stream")
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2344
			// _ = "end of CoverTab[127571]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2345
			_go_fuzz_dep_.CoverTab[127572]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2345
			// _ = "end of CoverTab[127572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2345
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2345
		// _ = "end of CoverTab[127553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2345
		_go_fuzz_dep_.CoverTab[127554]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2348
		leading_blanks := false
		for !is_blankz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2349
			_go_fuzz_dep_.CoverTab[127573]++
												if single && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				_go_fuzz_dep_.CoverTab[127575]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				return parser.buffer[parser.buffer_pos] == '\''
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				// _ = "end of CoverTab[127575]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				_go_fuzz_dep_.CoverTab[127576]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				return parser.buffer[parser.buffer_pos+1] == '\''
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				// _ = "end of CoverTab[127576]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2350
				_go_fuzz_dep_.CoverTab[127577]++

													s = append(s, '\'')
													skip(parser)
													skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2354
				// _ = "end of CoverTab[127577]"

			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
				_go_fuzz_dep_.CoverTab[127578]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
				if single && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
					_go_fuzz_dep_.CoverTab[127579]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
					return parser.buffer[parser.buffer_pos] == '\''
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
					// _ = "end of CoverTab[127579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2356
					_go_fuzz_dep_.CoverTab[127580]++

														break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2358
					// _ = "end of CoverTab[127580]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
					_go_fuzz_dep_.CoverTab[127581]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
					if !single && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
						_go_fuzz_dep_.CoverTab[127582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
						return parser.buffer[parser.buffer_pos] == '"'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
						// _ = "end of CoverTab[127582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
					}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2359
						_go_fuzz_dep_.CoverTab[127583]++

															break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2361
						// _ = "end of CoverTab[127583]"

					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
						_go_fuzz_dep_.CoverTab[127584]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
						if !single && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							_go_fuzz_dep_.CoverTab[127585]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							return parser.buffer[parser.buffer_pos] == '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							// _ = "end of CoverTab[127585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
						}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							_go_fuzz_dep_.CoverTab[127586]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							return is_break(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							// _ = "end of CoverTab[127586]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
						}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2363
							_go_fuzz_dep_.CoverTab[127587]++

																if parser.unread < 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2365
								_go_fuzz_dep_.CoverTab[127589]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2365
								return !yaml_parser_update_buffer(parser, 3)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2365
								// _ = "end of CoverTab[127589]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2365
							}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2365
								_go_fuzz_dep_.CoverTab[127590]++
																	return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2366
								// _ = "end of CoverTab[127590]"
							} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2367
								_go_fuzz_dep_.CoverTab[127591]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2367
								// _ = "end of CoverTab[127591]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2367
							}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2367
							// _ = "end of CoverTab[127587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2367
							_go_fuzz_dep_.CoverTab[127588]++
																skip(parser)
																skip_line(parser)
																leading_blanks = true
																break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2371
							// _ = "end of CoverTab[127588]"

						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
							_go_fuzz_dep_.CoverTab[127592]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
							if !single && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
								_go_fuzz_dep_.CoverTab[127593]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
								return parser.buffer[parser.buffer_pos] == '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
								// _ = "end of CoverTab[127593]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
							}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2373
								_go_fuzz_dep_.CoverTab[127594]++

																	code_length := 0

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2378
								switch parser.buffer[parser.buffer_pos+1] {
								case '0':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2379
									_go_fuzz_dep_.CoverTab[127596]++
																		s = append(s, 0)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2380
									// _ = "end of CoverTab[127596]"
								case 'a':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2381
									_go_fuzz_dep_.CoverTab[127597]++
																		s = append(s, '\x07')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2382
									// _ = "end of CoverTab[127597]"
								case 'b':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2383
									_go_fuzz_dep_.CoverTab[127598]++
																		s = append(s, '\x08')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2384
									// _ = "end of CoverTab[127598]"
								case 't', '\t':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2385
									_go_fuzz_dep_.CoverTab[127599]++
																		s = append(s, '\x09')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2386
									// _ = "end of CoverTab[127599]"
								case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2387
									_go_fuzz_dep_.CoverTab[127600]++
																		s = append(s, '\x0A')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2388
									// _ = "end of CoverTab[127600]"
								case 'v':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2389
									_go_fuzz_dep_.CoverTab[127601]++
																		s = append(s, '\x0B')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2390
									// _ = "end of CoverTab[127601]"
								case 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2391
									_go_fuzz_dep_.CoverTab[127602]++
																		s = append(s, '\x0C')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2392
									// _ = "end of CoverTab[127602]"
								case 'r':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2393
									_go_fuzz_dep_.CoverTab[127603]++
																		s = append(s, '\x0D')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2394
									// _ = "end of CoverTab[127603]"
								case 'e':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2395
									_go_fuzz_dep_.CoverTab[127604]++
																		s = append(s, '\x1B')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2396
									// _ = "end of CoverTab[127604]"
								case ' ':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2397
									_go_fuzz_dep_.CoverTab[127605]++
																		s = append(s, '\x20')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2398
									// _ = "end of CoverTab[127605]"
								case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2399
									_go_fuzz_dep_.CoverTab[127606]++
																		s = append(s, '"')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2400
									// _ = "end of CoverTab[127606]"
								case '\'':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2401
									_go_fuzz_dep_.CoverTab[127607]++
																		s = append(s, '\'')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2402
									// _ = "end of CoverTab[127607]"
								case '\\':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2403
									_go_fuzz_dep_.CoverTab[127608]++
																		s = append(s, '\\')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2404
									// _ = "end of CoverTab[127608]"
								case 'N':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2405
									_go_fuzz_dep_.CoverTab[127609]++
																		s = append(s, '\xC2')
																		s = append(s, '\x85')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2407
									// _ = "end of CoverTab[127609]"
								case '_':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2408
									_go_fuzz_dep_.CoverTab[127610]++
																		s = append(s, '\xC2')
																		s = append(s, '\xA0')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2410
									// _ = "end of CoverTab[127610]"
								case 'L':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2411
									_go_fuzz_dep_.CoverTab[127611]++
																		s = append(s, '\xE2')
																		s = append(s, '\x80')
																		s = append(s, '\xA8')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2414
									// _ = "end of CoverTab[127611]"
								case 'P':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2415
									_go_fuzz_dep_.CoverTab[127612]++
																		s = append(s, '\xE2')
																		s = append(s, '\x80')
																		s = append(s, '\xA9')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2418
									// _ = "end of CoverTab[127612]"
								case 'x':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2419
									_go_fuzz_dep_.CoverTab[127613]++
																		code_length = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2420
									// _ = "end of CoverTab[127613]"
								case 'u':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2421
									_go_fuzz_dep_.CoverTab[127614]++
																		code_length = 4
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2422
									// _ = "end of CoverTab[127614]"
								case 'U':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2423
									_go_fuzz_dep_.CoverTab[127615]++
																		code_length = 8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2424
									// _ = "end of CoverTab[127615]"
								default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2425
									_go_fuzz_dep_.CoverTab[127616]++
																		yaml_parser_set_scanner_error(parser, "while parsing a quoted scalar",
										start_mark, "found unknown escape character")
																		return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2428
									// _ = "end of CoverTab[127616]"
								}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2429
								// _ = "end of CoverTab[127594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2429
								_go_fuzz_dep_.CoverTab[127595]++

																	skip(parser)
																	skip(parser)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2435
								if code_length > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2435
									_go_fuzz_dep_.CoverTab[127617]++
																		var value int

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
									if parser.unread < code_length && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
										_go_fuzz_dep_.CoverTab[127622]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
										return !yaml_parser_update_buffer(parser, code_length)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
										// _ = "end of CoverTab[127622]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
									}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2439
										_go_fuzz_dep_.CoverTab[127623]++
																			return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2440
										// _ = "end of CoverTab[127623]"
									} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2441
										_go_fuzz_dep_.CoverTab[127624]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2441
										// _ = "end of CoverTab[127624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2441
									}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2441
									// _ = "end of CoverTab[127617]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2441
									_go_fuzz_dep_.CoverTab[127618]++
																		for k := 0; k < code_length; k++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2442
										_go_fuzz_dep_.CoverTab[127625]++
																			if !is_hex(parser.buffer, parser.buffer_pos+k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2443
											_go_fuzz_dep_.CoverTab[127627]++
																				yaml_parser_set_scanner_error(parser, "while parsing a quoted scalar",
												start_mark, "did not find expected hexdecimal number")
																				return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2446
											// _ = "end of CoverTab[127627]"
										} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2447
											_go_fuzz_dep_.CoverTab[127628]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2447
											// _ = "end of CoverTab[127628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2447
										}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2447
										// _ = "end of CoverTab[127625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2447
										_go_fuzz_dep_.CoverTab[127626]++
																			value = (value << 4) + as_hex(parser.buffer, parser.buffer_pos+k)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2448
										// _ = "end of CoverTab[127626]"
									}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2449
									// _ = "end of CoverTab[127618]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2449
									_go_fuzz_dep_.CoverTab[127619]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
									if (value >= 0xD800 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										_go_fuzz_dep_.CoverTab[127629]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										return value <= 0xDFFF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										// _ = "end of CoverTab[127629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
									}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										_go_fuzz_dep_.CoverTab[127630]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										return value > 0x10FFFF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										// _ = "end of CoverTab[127630]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
									}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2452
										_go_fuzz_dep_.CoverTab[127631]++
																			yaml_parser_set_scanner_error(parser, "while parsing a quoted scalar",
											start_mark, "found invalid Unicode character escape code")
																			return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2455
										// _ = "end of CoverTab[127631]"
									} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2456
										_go_fuzz_dep_.CoverTab[127632]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2456
										// _ = "end of CoverTab[127632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2456
									}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2456
									// _ = "end of CoverTab[127619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2456
									_go_fuzz_dep_.CoverTab[127620]++
																		if value <= 0x7F {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2457
										_go_fuzz_dep_.CoverTab[127633]++
																			s = append(s, byte(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2458
										// _ = "end of CoverTab[127633]"
									} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2459
										_go_fuzz_dep_.CoverTab[127634]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2459
										if value <= 0x7FF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2459
											_go_fuzz_dep_.CoverTab[127635]++
																				s = append(s, byte(0xC0+(value>>6)))
																				s = append(s, byte(0x80+(value&0x3F)))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2461
											// _ = "end of CoverTab[127635]"
										} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2462
											_go_fuzz_dep_.CoverTab[127636]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2462
											if value <= 0xFFFF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2462
												_go_fuzz_dep_.CoverTab[127637]++
																					s = append(s, byte(0xE0+(value>>12)))
																					s = append(s, byte(0x80+((value>>6)&0x3F)))
																					s = append(s, byte(0x80+(value&0x3F)))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2465
												// _ = "end of CoverTab[127637]"
											} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2466
												_go_fuzz_dep_.CoverTab[127638]++
																					s = append(s, byte(0xF0+(value>>18)))
																					s = append(s, byte(0x80+((value>>12)&0x3F)))
																					s = append(s, byte(0x80+((value>>6)&0x3F)))
																					s = append(s, byte(0x80+(value&0x3F)))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2470
												// _ = "end of CoverTab[127638]"
											}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
											// _ = "end of CoverTab[127636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
										}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
										// _ = "end of CoverTab[127634]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
									}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
									// _ = "end of CoverTab[127620]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2471
									_go_fuzz_dep_.CoverTab[127621]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2474
									for k := 0; k < code_length; k++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2474
										_go_fuzz_dep_.CoverTab[127639]++
																			skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2475
										// _ = "end of CoverTab[127639]"
									}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2476
									// _ = "end of CoverTab[127621]"
								} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2477
									_go_fuzz_dep_.CoverTab[127640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2477
									// _ = "end of CoverTab[127640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2477
								}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2477
								// _ = "end of CoverTab[127595]"
							} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2478
								_go_fuzz_dep_.CoverTab[127641]++

																	s = read(parser, s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2480
								// _ = "end of CoverTab[127641]"
							}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
							// _ = "end of CoverTab[127592]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
						// _ = "end of CoverTab[127584]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
					// _ = "end of CoverTab[127581]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
				// _ = "end of CoverTab[127578]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
			// _ = "end of CoverTab[127573]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2481
			_go_fuzz_dep_.CoverTab[127574]++
												if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2482
				_go_fuzz_dep_.CoverTab[127642]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2482
				return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2482
				// _ = "end of CoverTab[127642]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2482
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2482
				_go_fuzz_dep_.CoverTab[127643]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2483
				// _ = "end of CoverTab[127643]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2484
				_go_fuzz_dep_.CoverTab[127644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2484
				// _ = "end of CoverTab[127644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2484
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2484
			// _ = "end of CoverTab[127574]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2485
		// _ = "end of CoverTab[127554]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2485
		_go_fuzz_dep_.CoverTab[127555]++

											if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2487
			_go_fuzz_dep_.CoverTab[127645]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2487
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2487
			// _ = "end of CoverTab[127645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2487
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2487
			_go_fuzz_dep_.CoverTab[127646]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2488
			// _ = "end of CoverTab[127646]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2489
			_go_fuzz_dep_.CoverTab[127647]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2489
			// _ = "end of CoverTab[127647]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2489
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2489
		// _ = "end of CoverTab[127555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2489
		_go_fuzz_dep_.CoverTab[127556]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2492
		if single {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2492
			_go_fuzz_dep_.CoverTab[127648]++
												if parser.buffer[parser.buffer_pos] == '\'' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2493
				_go_fuzz_dep_.CoverTab[127649]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2494
				// _ = "end of CoverTab[127649]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2495
				_go_fuzz_dep_.CoverTab[127650]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2495
				// _ = "end of CoverTab[127650]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2495
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2495
			// _ = "end of CoverTab[127648]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2496
			_go_fuzz_dep_.CoverTab[127651]++
												if parser.buffer[parser.buffer_pos] == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2497
				_go_fuzz_dep_.CoverTab[127652]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2498
				// _ = "end of CoverTab[127652]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2499
				_go_fuzz_dep_.CoverTab[127653]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2499
				// _ = "end of CoverTab[127653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2499
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2499
			// _ = "end of CoverTab[127651]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2500
		// _ = "end of CoverTab[127556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2500
		_go_fuzz_dep_.CoverTab[127557]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
		for is_blank(parser.buffer, parser.buffer_pos) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
			_go_fuzz_dep_.CoverTab[127654]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
			return is_break(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
			// _ = "end of CoverTab[127654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2503
			_go_fuzz_dep_.CoverTab[127655]++
												if is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2504
				_go_fuzz_dep_.CoverTab[127657]++

													if !leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2506
					_go_fuzz_dep_.CoverTab[127658]++
														whitespaces = read(parser, whitespaces)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2507
					// _ = "end of CoverTab[127658]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2508
					_go_fuzz_dep_.CoverTab[127659]++
														skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2509
					// _ = "end of CoverTab[127659]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2510
				// _ = "end of CoverTab[127657]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2511
				_go_fuzz_dep_.CoverTab[127660]++
													if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2512
					_go_fuzz_dep_.CoverTab[127662]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2512
					return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2512
					// _ = "end of CoverTab[127662]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2512
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2512
					_go_fuzz_dep_.CoverTab[127663]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2513
					// _ = "end of CoverTab[127663]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2514
					_go_fuzz_dep_.CoverTab[127664]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2514
					// _ = "end of CoverTab[127664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2514
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2514
				// _ = "end of CoverTab[127660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2514
				_go_fuzz_dep_.CoverTab[127661]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2517
				if !leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2517
					_go_fuzz_dep_.CoverTab[127665]++
														whitespaces = whitespaces[:0]
														leading_break = read_line(parser, leading_break)
														leading_blanks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2520
					// _ = "end of CoverTab[127665]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2521
					_go_fuzz_dep_.CoverTab[127666]++
														trailing_breaks = read_line(parser, trailing_breaks)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2522
					// _ = "end of CoverTab[127666]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2523
				// _ = "end of CoverTab[127661]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2524
			// _ = "end of CoverTab[127655]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2524
			_go_fuzz_dep_.CoverTab[127656]++
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2525
				_go_fuzz_dep_.CoverTab[127667]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2525
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2525
				// _ = "end of CoverTab[127667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2525
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2525
				_go_fuzz_dep_.CoverTab[127668]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2526
				// _ = "end of CoverTab[127668]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2527
				_go_fuzz_dep_.CoverTab[127669]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2527
				// _ = "end of CoverTab[127669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2527
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2527
			// _ = "end of CoverTab[127656]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2528
		// _ = "end of CoverTab[127557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2528
		_go_fuzz_dep_.CoverTab[127558]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2531
		if leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2531
			_go_fuzz_dep_.CoverTab[127670]++

												if len(leading_break) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2533
				_go_fuzz_dep_.CoverTab[127672]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2533
				return leading_break[0] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2533
				// _ = "end of CoverTab[127672]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2533
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2533
				_go_fuzz_dep_.CoverTab[127673]++
													if len(trailing_breaks) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2534
					_go_fuzz_dep_.CoverTab[127674]++
														s = append(s, ' ')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2535
					// _ = "end of CoverTab[127674]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2536
					_go_fuzz_dep_.CoverTab[127675]++
														s = append(s, trailing_breaks...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2537
					// _ = "end of CoverTab[127675]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2538
				// _ = "end of CoverTab[127673]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2539
				_go_fuzz_dep_.CoverTab[127676]++
													s = append(s, leading_break...)
													s = append(s, trailing_breaks...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2541
				// _ = "end of CoverTab[127676]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2542
			// _ = "end of CoverTab[127670]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2542
			_go_fuzz_dep_.CoverTab[127671]++
												trailing_breaks = trailing_breaks[:0]
												leading_break = leading_break[:0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2544
			// _ = "end of CoverTab[127671]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2545
			_go_fuzz_dep_.CoverTab[127677]++
												s = append(s, whitespaces...)
												whitespaces = whitespaces[:0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2547
			// _ = "end of CoverTab[127677]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2548
		// _ = "end of CoverTab[127558]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2549
	// _ = "end of CoverTab[127548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2549
	_go_fuzz_dep_.CoverTab[127549]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2552
	skip(parser)
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2556
	*token = yaml_token_t{
		typ:		yaml_SCALAR_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
		value:		s,
		style:		yaml_SINGLE_QUOTED_SCALAR_STYLE,
	}
	if !single {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2563
		_go_fuzz_dep_.CoverTab[127678]++
											token.style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2564
		// _ = "end of CoverTab[127678]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2565
		_go_fuzz_dep_.CoverTab[127679]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2565
		// _ = "end of CoverTab[127679]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2565
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2565
	// _ = "end of CoverTab[127549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2565
	_go_fuzz_dep_.CoverTab[127550]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2566
	// _ = "end of CoverTab[127550]"
}

// Scan a plain scalar.
func yaml_parser_scan_plain_scalar(parser *yaml_parser_t, token *yaml_token_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2570
	_go_fuzz_dep_.CoverTab[127680]++

										var s, leading_break, trailing_breaks, whitespaces []byte
										var leading_blanks bool
										var indent = parser.indent + 1

										start_mark := parser.mark
										end_mark := parser.mark

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2580
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2580
		_go_fuzz_dep_.CoverTab[127683]++

											if parser.unread < 4 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2582
			_go_fuzz_dep_.CoverTab[127691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2582
			return !yaml_parser_update_buffer(parser, 4)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2582
			// _ = "end of CoverTab[127691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2582
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2582
			_go_fuzz_dep_.CoverTab[127692]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2583
			// _ = "end of CoverTab[127692]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2584
			_go_fuzz_dep_.CoverTab[127693]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2584
			// _ = "end of CoverTab[127693]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2584
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2584
		// _ = "end of CoverTab[127683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2584
		_go_fuzz_dep_.CoverTab[127684]++
											if parser.mark.column == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2585
			_go_fuzz_dep_.CoverTab[127694]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2585
			return ((parser.buffer[parser.buffer_pos+0] == '-' && func() bool {
													_go_fuzz_dep_.CoverTab[127695]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2586
				return parser.buffer[parser.buffer_pos+1] == '-'
													// _ = "end of CoverTab[127695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2587
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2587
				_go_fuzz_dep_.CoverTab[127696]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2587
				return parser.buffer[parser.buffer_pos+2] == '-'
													// _ = "end of CoverTab[127696]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2588
			}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2588
				_go_fuzz_dep_.CoverTab[127697]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2588
				return (parser.buffer[parser.buffer_pos+0] == '.' && func() bool {
														_go_fuzz_dep_.CoverTab[127698]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2589
					return parser.buffer[parser.buffer_pos+1] == '.'
														// _ = "end of CoverTab[127698]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2590
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2590
					_go_fuzz_dep_.CoverTab[127699]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2590
					return parser.buffer[parser.buffer_pos+2] == '.'
														// _ = "end of CoverTab[127699]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
				// _ = "end of CoverTab[127697]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
			// _ = "end of CoverTab[127694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
			_go_fuzz_dep_.CoverTab[127700]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2591
			return is_blankz(parser.buffer, parser.buffer_pos+3)
												// _ = "end of CoverTab[127700]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2592
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2592
			_go_fuzz_dep_.CoverTab[127701]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2593
			// _ = "end of CoverTab[127701]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2594
			_go_fuzz_dep_.CoverTab[127702]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2594
			// _ = "end of CoverTab[127702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2594
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2594
		// _ = "end of CoverTab[127684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2594
		_go_fuzz_dep_.CoverTab[127685]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2597
		if parser.buffer[parser.buffer_pos] == '#' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2597
			_go_fuzz_dep_.CoverTab[127703]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2598
			// _ = "end of CoverTab[127703]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2599
			_go_fuzz_dep_.CoverTab[127704]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2599
			// _ = "end of CoverTab[127704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2599
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2599
		// _ = "end of CoverTab[127685]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2599
		_go_fuzz_dep_.CoverTab[127686]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2602
		for !is_blankz(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2602
			_go_fuzz_dep_.CoverTab[127705]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
			if (parser.buffer[parser.buffer_pos] == ':' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
				_go_fuzz_dep_.CoverTab[127708]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
				return is_blankz(parser.buffer, parser.buffer_pos+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
				// _ = "end of CoverTab[127708]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
			}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
				_go_fuzz_dep_.CoverTab[127709]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2605
				return (parser.flow_level > 0 && func() bool {
														_go_fuzz_dep_.CoverTab[127710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2606
					return (parser.buffer[parser.buffer_pos] == ',' || func() bool {
															_go_fuzz_dep_.CoverTab[127711]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2607
						return parser.buffer[parser.buffer_pos] == '?'
															// _ = "end of CoverTab[127711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
						_go_fuzz_dep_.CoverTab[127712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
						return parser.buffer[parser.buffer_pos] == '['
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
						// _ = "end of CoverTab[127712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
						_go_fuzz_dep_.CoverTab[127713]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2608
						return parser.buffer[parser.buffer_pos] == ']'
															// _ = "end of CoverTab[127713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
						_go_fuzz_dep_.CoverTab[127714]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
						return parser.buffer[parser.buffer_pos] == '{'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
						// _ = "end of CoverTab[127714]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
						_go_fuzz_dep_.CoverTab[127715]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2609
						return parser.buffer[parser.buffer_pos] == '}'
															// _ = "end of CoverTab[127715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
					}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
					// _ = "end of CoverTab[127710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
				// _ = "end of CoverTab[127709]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2610
				_go_fuzz_dep_.CoverTab[127716]++
													break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2611
				// _ = "end of CoverTab[127716]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2612
				_go_fuzz_dep_.CoverTab[127717]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2612
				// _ = "end of CoverTab[127717]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2612
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2612
			// _ = "end of CoverTab[127705]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2612
			_go_fuzz_dep_.CoverTab[127706]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
			if leading_blanks || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
				_go_fuzz_dep_.CoverTab[127718]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
				return len(whitespaces) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
				// _ = "end of CoverTab[127718]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2615
				_go_fuzz_dep_.CoverTab[127719]++
													if leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2616
					_go_fuzz_dep_.CoverTab[127720]++

														if leading_break[0] == '\n' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2618
						_go_fuzz_dep_.CoverTab[127722]++
															if len(trailing_breaks) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2619
							_go_fuzz_dep_.CoverTab[127723]++
																s = append(s, ' ')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2620
							// _ = "end of CoverTab[127723]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2621
							_go_fuzz_dep_.CoverTab[127724]++
																s = append(s, trailing_breaks...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2622
							// _ = "end of CoverTab[127724]"
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2623
						// _ = "end of CoverTab[127722]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2624
						_go_fuzz_dep_.CoverTab[127725]++
															s = append(s, leading_break...)
															s = append(s, trailing_breaks...)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2626
						// _ = "end of CoverTab[127725]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2627
					// _ = "end of CoverTab[127720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2627
					_go_fuzz_dep_.CoverTab[127721]++
														trailing_breaks = trailing_breaks[:0]
														leading_break = leading_break[:0]
														leading_blanks = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2630
					// _ = "end of CoverTab[127721]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2631
					_go_fuzz_dep_.CoverTab[127726]++
														s = append(s, whitespaces...)
														whitespaces = whitespaces[:0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2633
					// _ = "end of CoverTab[127726]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2634
				// _ = "end of CoverTab[127719]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2635
				_go_fuzz_dep_.CoverTab[127727]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2635
				// _ = "end of CoverTab[127727]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2635
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2635
			// _ = "end of CoverTab[127706]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2635
			_go_fuzz_dep_.CoverTab[127707]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2638
			s = read(parser, s)

			end_mark = parser.mark
			if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2641
				_go_fuzz_dep_.CoverTab[127728]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2641
				return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2641
				// _ = "end of CoverTab[127728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2641
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2641
				_go_fuzz_dep_.CoverTab[127729]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2642
				// _ = "end of CoverTab[127729]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2643
				_go_fuzz_dep_.CoverTab[127730]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2643
				// _ = "end of CoverTab[127730]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2643
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2643
			// _ = "end of CoverTab[127707]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2644
		// _ = "end of CoverTab[127686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2644
		_go_fuzz_dep_.CoverTab[127687]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
		if !(is_blank(parser.buffer, parser.buffer_pos) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
			_go_fuzz_dep_.CoverTab[127731]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
			return is_break(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
			// _ = "end of CoverTab[127731]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
		}()) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2647
			_go_fuzz_dep_.CoverTab[127732]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2648
			// _ = "end of CoverTab[127732]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2649
			_go_fuzz_dep_.CoverTab[127733]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2649
			// _ = "end of CoverTab[127733]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2649
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2649
		// _ = "end of CoverTab[127687]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2649
		_go_fuzz_dep_.CoverTab[127688]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
		if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
			_go_fuzz_dep_.CoverTab[127734]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
			return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
			// _ = "end of CoverTab[127734]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2652
			_go_fuzz_dep_.CoverTab[127735]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2653
			// _ = "end of CoverTab[127735]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2654
			_go_fuzz_dep_.CoverTab[127736]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2654
			// _ = "end of CoverTab[127736]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2654
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2654
		// _ = "end of CoverTab[127688]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2654
		_go_fuzz_dep_.CoverTab[127689]++

											for is_blank(parser.buffer, parser.buffer_pos) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2656
			_go_fuzz_dep_.CoverTab[127737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2656
			return is_break(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2656
			// _ = "end of CoverTab[127737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2656
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2656
			_go_fuzz_dep_.CoverTab[127738]++
												if is_blank(parser.buffer, parser.buffer_pos) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2657
				_go_fuzz_dep_.CoverTab[127740]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
				if leading_blanks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					_go_fuzz_dep_.CoverTab[127742]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					return parser.mark.column < indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					// _ = "end of CoverTab[127742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					_go_fuzz_dep_.CoverTab[127743]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					return is_tab(parser.buffer, parser.buffer_pos)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					// _ = "end of CoverTab[127743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2660
					_go_fuzz_dep_.CoverTab[127744]++
														yaml_parser_set_scanner_error(parser, "while scanning a plain scalar",
						start_mark, "found a tab character that violates indentation")
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2663
					// _ = "end of CoverTab[127744]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2664
					_go_fuzz_dep_.CoverTab[127745]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2664
					// _ = "end of CoverTab[127745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2664
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2664
				// _ = "end of CoverTab[127740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2664
				_go_fuzz_dep_.CoverTab[127741]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2667
				if !leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2667
					_go_fuzz_dep_.CoverTab[127746]++
														whitespaces = read(parser, whitespaces)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2668
					// _ = "end of CoverTab[127746]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2669
					_go_fuzz_dep_.CoverTab[127747]++
														skip(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2670
					// _ = "end of CoverTab[127747]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2671
				// _ = "end of CoverTab[127741]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2672
				_go_fuzz_dep_.CoverTab[127748]++
													if parser.unread < 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2673
					_go_fuzz_dep_.CoverTab[127750]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2673
					return !yaml_parser_update_buffer(parser, 2)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2673
					// _ = "end of CoverTab[127750]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2673
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2673
					_go_fuzz_dep_.CoverTab[127751]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2674
					// _ = "end of CoverTab[127751]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2675
					_go_fuzz_dep_.CoverTab[127752]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2675
					// _ = "end of CoverTab[127752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2675
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2675
				// _ = "end of CoverTab[127748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2675
				_go_fuzz_dep_.CoverTab[127749]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2678
				if !leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2678
					_go_fuzz_dep_.CoverTab[127753]++
														whitespaces = whitespaces[:0]
														leading_break = read_line(parser, leading_break)
														leading_blanks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2681
					// _ = "end of CoverTab[127753]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2682
					_go_fuzz_dep_.CoverTab[127754]++
														trailing_breaks = read_line(parser, trailing_breaks)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2683
					// _ = "end of CoverTab[127754]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2684
				// _ = "end of CoverTab[127749]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2685
			// _ = "end of CoverTab[127738]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2685
			_go_fuzz_dep_.CoverTab[127739]++
												if parser.unread < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2686
				_go_fuzz_dep_.CoverTab[127755]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2686
				return !yaml_parser_update_buffer(parser, 1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2686
				// _ = "end of CoverTab[127755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2686
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2686
				_go_fuzz_dep_.CoverTab[127756]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2687
				// _ = "end of CoverTab[127756]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2688
				_go_fuzz_dep_.CoverTab[127757]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2688
				// _ = "end of CoverTab[127757]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2688
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2688
			// _ = "end of CoverTab[127739]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2689
		// _ = "end of CoverTab[127689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2689
		_go_fuzz_dep_.CoverTab[127690]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
		if parser.flow_level == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
			_go_fuzz_dep_.CoverTab[127758]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
			return parser.mark.column < indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
			// _ = "end of CoverTab[127758]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2692
			_go_fuzz_dep_.CoverTab[127759]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2693
			// _ = "end of CoverTab[127759]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2694
			_go_fuzz_dep_.CoverTab[127760]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2694
			// _ = "end of CoverTab[127760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2694
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2694
		// _ = "end of CoverTab[127690]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2695
	// _ = "end of CoverTab[127680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2695
	_go_fuzz_dep_.CoverTab[127681]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2698
	*token = yaml_token_t{
		typ:		yaml_SCALAR_TOKEN,
		start_mark:	start_mark,
		end_mark:	end_mark,
		value:		s,
		style:		yaml_PLAIN_SCALAR_STYLE,
	}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2707
	if leading_blanks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2707
		_go_fuzz_dep_.CoverTab[127761]++
											parser.simple_key_allowed = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2708
		// _ = "end of CoverTab[127761]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2709
		_go_fuzz_dep_.CoverTab[127762]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2709
		// _ = "end of CoverTab[127762]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2709
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2709
	// _ = "end of CoverTab[127681]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2709
	_go_fuzz_dep_.CoverTab[127682]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2710
	// _ = "end of CoverTab[127682]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2711
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/scannerc.go:2711
var _ = _go_fuzz_dep_.CoverTab
