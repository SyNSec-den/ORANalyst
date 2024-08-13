//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:1
)

import (
	"io"
)

// Set the reader error and return 0.
func yaml_parser_set_reader_error(parser *yaml_parser_t, problem string, offset int, value int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:8
		_go_fuzz_dep_.CoverTab[126477]++
										parser.error = yaml_READER_ERROR
										parser.problem = problem
										parser.problem_offset = offset
										parser.problem_value = value
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:13
	// _ = "end of CoverTab[126477]"
}

// Byte order marks.
const (
	bom_UTF8	= "\xef\xbb\xbf"
	bom_UTF16LE	= "\xff\xfe"
	bom_UTF16BE	= "\xfe\xff"
)

// Determine the input stream encoding by checking the BOM symbol. If no BOM is
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:23
// found, the UTF-8 encoding is assumed. Return 1 on success, 0 on failure.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:25
func yaml_parser_determine_encoding(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:25
	_go_fuzz_dep_.CoverTab[126478]++

										for !parser.eof && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:27
		_go_fuzz_dep_.CoverTab[126481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:27
		return len(parser.raw_buffer)-parser.raw_buffer_pos < 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:27
		// _ = "end of CoverTab[126481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:27
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:27
		_go_fuzz_dep_.CoverTab[126482]++
											if !yaml_parser_update_raw_buffer(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:28
			_go_fuzz_dep_.CoverTab[126483]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:29
			// _ = "end of CoverTab[126483]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:30
			_go_fuzz_dep_.CoverTab[126484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:30
			// _ = "end of CoverTab[126484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:30
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:30
		// _ = "end of CoverTab[126482]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:31
	// _ = "end of CoverTab[126478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:31
	_go_fuzz_dep_.CoverTab[126479]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:34
	buf := parser.raw_buffer
	pos := parser.raw_buffer_pos
	avail := len(buf) - pos
	if avail >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		_go_fuzz_dep_.CoverTab[126485]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		return buf[pos] == bom_UTF16LE[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		// _ = "end of CoverTab[126485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		_go_fuzz_dep_.CoverTab[126486]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		return buf[pos+1] == bom_UTF16LE[1]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		// _ = "end of CoverTab[126486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:37
		_go_fuzz_dep_.CoverTab[126487]++
											parser.encoding = yaml_UTF16LE_ENCODING
											parser.raw_buffer_pos += 2
											parser.offset += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:40
		// _ = "end of CoverTab[126487]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
		_go_fuzz_dep_.CoverTab[126488]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
		if avail >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			_go_fuzz_dep_.CoverTab[126489]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			return buf[pos] == bom_UTF16BE[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			// _ = "end of CoverTab[126489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			_go_fuzz_dep_.CoverTab[126490]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			return buf[pos+1] == bom_UTF16BE[1]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			// _ = "end of CoverTab[126490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:41
			_go_fuzz_dep_.CoverTab[126491]++
												parser.encoding = yaml_UTF16BE_ENCODING
												parser.raw_buffer_pos += 2
												parser.offset += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:44
			// _ = "end of CoverTab[126491]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
			_go_fuzz_dep_.CoverTab[126492]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
			if avail >= 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				_go_fuzz_dep_.CoverTab[126493]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				return buf[pos] == bom_UTF8[0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				// _ = "end of CoverTab[126493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				_go_fuzz_dep_.CoverTab[126494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				return buf[pos+1] == bom_UTF8[1]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				// _ = "end of CoverTab[126494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				_go_fuzz_dep_.CoverTab[126495]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				return buf[pos+2] == bom_UTF8[2]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				// _ = "end of CoverTab[126495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:45
				_go_fuzz_dep_.CoverTab[126496]++
													parser.encoding = yaml_UTF8_ENCODING
													parser.raw_buffer_pos += 3
													parser.offset += 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:48
				// _ = "end of CoverTab[126496]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:49
				_go_fuzz_dep_.CoverTab[126497]++
													parser.encoding = yaml_UTF8_ENCODING
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:50
				// _ = "end of CoverTab[126497]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
			// _ = "end of CoverTab[126492]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
		// _ = "end of CoverTab[126488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
	// _ = "end of CoverTab[126479]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:51
	_go_fuzz_dep_.CoverTab[126480]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:52
	// _ = "end of CoverTab[126480]"
}

// Update the raw buffer.
func yaml_parser_update_raw_buffer(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:56
	_go_fuzz_dep_.CoverTab[126498]++
										size_read := 0

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
	if parser.raw_buffer_pos == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
		_go_fuzz_dep_.CoverTab[126503]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
		return len(parser.raw_buffer) == cap(parser.raw_buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
		// _ = "end of CoverTab[126503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:60
		_go_fuzz_dep_.CoverTab[126504]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:61
		// _ = "end of CoverTab[126504]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:62
		_go_fuzz_dep_.CoverTab[126505]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:62
		// _ = "end of CoverTab[126505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:62
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:62
	// _ = "end of CoverTab[126498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:62
	_go_fuzz_dep_.CoverTab[126499]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:65
	if parser.eof {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:65
		_go_fuzz_dep_.CoverTab[126506]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:66
		// _ = "end of CoverTab[126506]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:67
		_go_fuzz_dep_.CoverTab[126507]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:67
		// _ = "end of CoverTab[126507]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:67
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:67
	// _ = "end of CoverTab[126499]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:67
	_go_fuzz_dep_.CoverTab[126500]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
	if parser.raw_buffer_pos > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
		_go_fuzz_dep_.CoverTab[126508]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
		return parser.raw_buffer_pos < len(parser.raw_buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
		// _ = "end of CoverTab[126508]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:70
		_go_fuzz_dep_.CoverTab[126509]++
											copy(parser.raw_buffer, parser.raw_buffer[parser.raw_buffer_pos:])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:71
		// _ = "end of CoverTab[126509]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:72
		_go_fuzz_dep_.CoverTab[126510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:72
		// _ = "end of CoverTab[126510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:72
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:72
	// _ = "end of CoverTab[126500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:72
	_go_fuzz_dep_.CoverTab[126501]++
										parser.raw_buffer = parser.raw_buffer[:len(parser.raw_buffer)-parser.raw_buffer_pos]
										parser.raw_buffer_pos = 0

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:77
	size_read, err := parser.read_handler(parser, parser.raw_buffer[len(parser.raw_buffer):cap(parser.raw_buffer)])
	parser.raw_buffer = parser.raw_buffer[:len(parser.raw_buffer)+size_read]
	if err == io.EOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:79
		_go_fuzz_dep_.CoverTab[126511]++
											parser.eof = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:80
		// _ = "end of CoverTab[126511]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:81
		_go_fuzz_dep_.CoverTab[126512]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:81
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:81
			_go_fuzz_dep_.CoverTab[126513]++
												return yaml_parser_set_reader_error(parser, "input error: "+err.Error(), parser.offset, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:82
			// _ = "end of CoverTab[126513]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
			_go_fuzz_dep_.CoverTab[126514]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
			// _ = "end of CoverTab[126514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
		// _ = "end of CoverTab[126512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
	// _ = "end of CoverTab[126501]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:83
	_go_fuzz_dep_.CoverTab[126502]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:84
	// _ = "end of CoverTab[126502]"
}

// Ensure that the buffer contains at least `length` characters.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:87
// Return true on success, false on failure.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:87
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:87
// The length is supposed to be significantly less that the buffer size.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:91
func yaml_parser_update_buffer(parser *yaml_parser_t, length int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:91
	_go_fuzz_dep_.CoverTab[126515]++
										if parser.read_handler == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:92
		_go_fuzz_dep_.CoverTab[126523]++
											panic("read handler must be set")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:93
		// _ = "end of CoverTab[126523]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:94
		_go_fuzz_dep_.CoverTab[126524]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:94
		// _ = "end of CoverTab[126524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:94
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:94
	// _ = "end of CoverTab[126515]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:94
	_go_fuzz_dep_.CoverTab[126516]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
	if parser.eof && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
		_go_fuzz_dep_.CoverTab[126525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
		return parser.raw_buffer_pos == len(parser.raw_buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
		// _ = "end of CoverTab[126525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
		_go_fuzz_dep_.CoverTab[126526]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:101
		// _ = "end of CoverTab[126526]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
		_go_fuzz_dep_.CoverTab[126527]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
		// _ = "end of CoverTab[126527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
	// _ = "end of CoverTab[126516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:108
	_go_fuzz_dep_.CoverTab[126517]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:111
	if parser.unread >= length {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:111
		_go_fuzz_dep_.CoverTab[126528]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:112
		// _ = "end of CoverTab[126528]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:113
		_go_fuzz_dep_.CoverTab[126529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:113
		// _ = "end of CoverTab[126529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:113
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:113
	// _ = "end of CoverTab[126517]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:113
	_go_fuzz_dep_.CoverTab[126518]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:116
	if parser.encoding == yaml_ANY_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:116
		_go_fuzz_dep_.CoverTab[126530]++
											if !yaml_parser_determine_encoding(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:117
			_go_fuzz_dep_.CoverTab[126531]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:118
			// _ = "end of CoverTab[126531]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:119
			_go_fuzz_dep_.CoverTab[126532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:119
			// _ = "end of CoverTab[126532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:119
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:119
		// _ = "end of CoverTab[126530]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:120
		_go_fuzz_dep_.CoverTab[126533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:120
		// _ = "end of CoverTab[126533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:120
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:120
	// _ = "end of CoverTab[126518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:120
	_go_fuzz_dep_.CoverTab[126519]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:123
	buffer_len := len(parser.buffer)
	if parser.buffer_pos > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:124
		_go_fuzz_dep_.CoverTab[126534]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:124
		return parser.buffer_pos < buffer_len
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:124
		// _ = "end of CoverTab[126534]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:124
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:124
		_go_fuzz_dep_.CoverTab[126535]++
											copy(parser.buffer, parser.buffer[parser.buffer_pos:])
											buffer_len -= parser.buffer_pos
											parser.buffer_pos = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:127
		// _ = "end of CoverTab[126535]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:128
		_go_fuzz_dep_.CoverTab[126536]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:128
		if parser.buffer_pos == buffer_len {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:128
			_go_fuzz_dep_.CoverTab[126537]++
												buffer_len = 0
												parser.buffer_pos = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:130
			// _ = "end of CoverTab[126537]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
			_go_fuzz_dep_.CoverTab[126538]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
			// _ = "end of CoverTab[126538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
		// _ = "end of CoverTab[126536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
	// _ = "end of CoverTab[126519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:131
	_go_fuzz_dep_.CoverTab[126520]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:134
	parser.buffer = parser.buffer[:cap(parser.buffer)]

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:137
	first := true
	for parser.unread < length {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:138
		_go_fuzz_dep_.CoverTab[126539]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
		if !first || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
			_go_fuzz_dep_.CoverTab[126542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
			return parser.raw_buffer_pos == len(parser.raw_buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
			// _ = "end of CoverTab[126542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:141
			_go_fuzz_dep_.CoverTab[126543]++
												if !yaml_parser_update_raw_buffer(parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:142
				_go_fuzz_dep_.CoverTab[126544]++
													parser.buffer = parser.buffer[:buffer_len]
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:144
				// _ = "end of CoverTab[126544]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:145
				_go_fuzz_dep_.CoverTab[126545]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:145
				// _ = "end of CoverTab[126545]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:145
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:145
			// _ = "end of CoverTab[126543]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:146
			_go_fuzz_dep_.CoverTab[126546]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:146
			// _ = "end of CoverTab[126546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:146
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:146
		// _ = "end of CoverTab[126539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:146
		_go_fuzz_dep_.CoverTab[126540]++
											first = false

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:150
	inner:
		for parser.raw_buffer_pos != len(parser.raw_buffer) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:151
			_go_fuzz_dep_.CoverTab[126547]++
												var value rune
												var width int

												raw_unread := len(parser.raw_buffer) - parser.raw_buffer_pos

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:158
			switch parser.encoding {
			case yaml_UTF8_ENCODING:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:159
				_go_fuzz_dep_.CoverTab[126551]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:179
				octet := parser.raw_buffer[parser.raw_buffer_pos]
				switch {
				case octet&0x80 == 0x00:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:181
					_go_fuzz_dep_.CoverTab[126562]++
														width = 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:182
					// _ = "end of CoverTab[126562]"
				case octet&0xE0 == 0xC0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:183
					_go_fuzz_dep_.CoverTab[126563]++
														width = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:184
					// _ = "end of CoverTab[126563]"
				case octet&0xF0 == 0xE0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:185
					_go_fuzz_dep_.CoverTab[126564]++
														width = 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:186
					// _ = "end of CoverTab[126564]"
				case octet&0xF8 == 0xF0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:187
					_go_fuzz_dep_.CoverTab[126565]++
														width = 4
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:188
					// _ = "end of CoverTab[126565]"
				default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:189
					_go_fuzz_dep_.CoverTab[126566]++

														return yaml_parser_set_reader_error(parser,
						"invalid leading UTF-8 octet",
						parser.offset, int(octet))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:193
					// _ = "end of CoverTab[126566]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:194
				// _ = "end of CoverTab[126551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:194
				_go_fuzz_dep_.CoverTab[126552]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:197
				if width > raw_unread {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:197
					_go_fuzz_dep_.CoverTab[126567]++
														if parser.eof {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:198
						_go_fuzz_dep_.CoverTab[126569]++
															return yaml_parser_set_reader_error(parser,
							"incomplete UTF-8 octet sequence",
							parser.offset, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:201
						// _ = "end of CoverTab[126569]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:202
						_go_fuzz_dep_.CoverTab[126570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:202
						// _ = "end of CoverTab[126570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:202
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:202
					// _ = "end of CoverTab[126567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:202
					_go_fuzz_dep_.CoverTab[126568]++
														break inner
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:203
					// _ = "end of CoverTab[126568]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:204
					_go_fuzz_dep_.CoverTab[126571]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:204
					// _ = "end of CoverTab[126571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:204
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:204
				// _ = "end of CoverTab[126552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:204
				_go_fuzz_dep_.CoverTab[126553]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:207
				switch {
				case octet&0x80 == 0x00:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:208
					_go_fuzz_dep_.CoverTab[126572]++
														value = rune(octet & 0x7F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:209
					// _ = "end of CoverTab[126572]"
				case octet&0xE0 == 0xC0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:210
					_go_fuzz_dep_.CoverTab[126573]++
														value = rune(octet & 0x1F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:211
					// _ = "end of CoverTab[126573]"
				case octet&0xF0 == 0xE0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:212
					_go_fuzz_dep_.CoverTab[126574]++
														value = rune(octet & 0x0F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:213
					// _ = "end of CoverTab[126574]"
				case octet&0xF8 == 0xF0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:214
					_go_fuzz_dep_.CoverTab[126575]++
														value = rune(octet & 0x07)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:215
					// _ = "end of CoverTab[126575]"
				default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:216
					_go_fuzz_dep_.CoverTab[126576]++
														value = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:217
					// _ = "end of CoverTab[126576]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:218
				// _ = "end of CoverTab[126553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:218
				_go_fuzz_dep_.CoverTab[126554]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:221
				for k := 1; k < width; k++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:221
					_go_fuzz_dep_.CoverTab[126577]++
														octet = parser.raw_buffer[parser.raw_buffer_pos+k]

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:225
					if (octet & 0xC0) != 0x80 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:225
						_go_fuzz_dep_.CoverTab[126579]++
															return yaml_parser_set_reader_error(parser,
							"invalid trailing UTF-8 octet",
							parser.offset+k, int(octet))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:228
						// _ = "end of CoverTab[126579]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:229
						_go_fuzz_dep_.CoverTab[126580]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:229
						// _ = "end of CoverTab[126580]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:229
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:229
					// _ = "end of CoverTab[126577]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:229
					_go_fuzz_dep_.CoverTab[126578]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:232
					value = (value << 6) + rune(octet&0x3F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:232
					// _ = "end of CoverTab[126578]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:233
				// _ = "end of CoverTab[126554]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:233
				_go_fuzz_dep_.CoverTab[126555]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:236
				switch {
				case width == 1:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:237
					_go_fuzz_dep_.CoverTab[126581]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:237
					// _ = "end of CoverTab[126581]"
				case width == 2 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
					_go_fuzz_dep_.CoverTab[126586]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
					return value >= 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
					// _ = "end of CoverTab[126586]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
				}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
					_go_fuzz_dep_.CoverTab[126582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:238
					// _ = "end of CoverTab[126582]"
				case width == 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
					_go_fuzz_dep_.CoverTab[126587]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
					return value >= 0x800
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
					// _ = "end of CoverTab[126587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
				}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
					_go_fuzz_dep_.CoverTab[126583]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:239
					// _ = "end of CoverTab[126583]"
				case width == 4 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
					_go_fuzz_dep_.CoverTab[126588]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
					return value >= 0x10000
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
					// _ = "end of CoverTab[126588]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
				}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
					_go_fuzz_dep_.CoverTab[126584]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:240
					// _ = "end of CoverTab[126584]"
				default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:241
					_go_fuzz_dep_.CoverTab[126585]++
														return yaml_parser_set_reader_error(parser,
						"invalid length of a UTF-8 sequence",
						parser.offset, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:244
					// _ = "end of CoverTab[126585]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:245
				// _ = "end of CoverTab[126555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:245
				_go_fuzz_dep_.CoverTab[126556]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
				if value >= 0xD800 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					_go_fuzz_dep_.CoverTab[126589]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					return value <= 0xDFFF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					// _ = "end of CoverTab[126589]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					_go_fuzz_dep_.CoverTab[126590]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					return value > 0x10FFFF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					// _ = "end of CoverTab[126590]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:248
					_go_fuzz_dep_.CoverTab[126591]++
														return yaml_parser_set_reader_error(parser,
						"invalid Unicode character",
						parser.offset, int(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:251
					// _ = "end of CoverTab[126591]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:252
					_go_fuzz_dep_.CoverTab[126592]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:252
					// _ = "end of CoverTab[126592]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:252
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:252
				// _ = "end of CoverTab[126556]"

			case yaml_UTF16LE_ENCODING, yaml_UTF16BE_ENCODING:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:254
				_go_fuzz_dep_.CoverTab[126557]++
													var low, high int
													if parser.encoding == yaml_UTF16LE_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:256
					_go_fuzz_dep_.CoverTab[126593]++
														low, high = 0, 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:257
					// _ = "end of CoverTab[126593]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:258
					_go_fuzz_dep_.CoverTab[126594]++
														low, high = 1, 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:259
					// _ = "end of CoverTab[126594]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:260
				// _ = "end of CoverTab[126557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:260
				_go_fuzz_dep_.CoverTab[126558]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:287
				if raw_unread < 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:287
					_go_fuzz_dep_.CoverTab[126595]++
														if parser.eof {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:288
						_go_fuzz_dep_.CoverTab[126597]++
															return yaml_parser_set_reader_error(parser,
							"incomplete UTF-16 character",
							parser.offset, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:291
						// _ = "end of CoverTab[126597]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:292
						_go_fuzz_dep_.CoverTab[126598]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:292
						// _ = "end of CoverTab[126598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:292
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:292
					// _ = "end of CoverTab[126595]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:292
					_go_fuzz_dep_.CoverTab[126596]++
														break inner
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:293
					// _ = "end of CoverTab[126596]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:294
					_go_fuzz_dep_.CoverTab[126599]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:294
					// _ = "end of CoverTab[126599]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:294
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:294
				// _ = "end of CoverTab[126558]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:294
				_go_fuzz_dep_.CoverTab[126559]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:297
				value = rune(parser.raw_buffer[parser.raw_buffer_pos+low]) +
					(rune(parser.raw_buffer[parser.raw_buffer_pos+high]) << 8)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:301
				if value&0xFC00 == 0xDC00 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:301
					_go_fuzz_dep_.CoverTab[126600]++
														return yaml_parser_set_reader_error(parser,
						"unexpected low surrogate area",
						parser.offset, int(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:304
					// _ = "end of CoverTab[126600]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:305
					_go_fuzz_dep_.CoverTab[126601]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:305
					// _ = "end of CoverTab[126601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:305
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:305
				// _ = "end of CoverTab[126559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:305
				_go_fuzz_dep_.CoverTab[126560]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:308
				if value&0xFC00 == 0xD800 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:308
					_go_fuzz_dep_.CoverTab[126602]++
														width = 4

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:312
					if raw_unread < 4 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:312
						_go_fuzz_dep_.CoverTab[126605]++
															if parser.eof {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:313
							_go_fuzz_dep_.CoverTab[126607]++
																return yaml_parser_set_reader_error(parser,
								"incomplete UTF-16 surrogate pair",
								parser.offset, -1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:316
							// _ = "end of CoverTab[126607]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:317
							_go_fuzz_dep_.CoverTab[126608]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:317
							// _ = "end of CoverTab[126608]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:317
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:317
						// _ = "end of CoverTab[126605]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:317
						_go_fuzz_dep_.CoverTab[126606]++
															break inner
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:318
						// _ = "end of CoverTab[126606]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:319
						_go_fuzz_dep_.CoverTab[126609]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:319
						// _ = "end of CoverTab[126609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:319
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:319
					// _ = "end of CoverTab[126602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:319
					_go_fuzz_dep_.CoverTab[126603]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:322
					value2 := rune(parser.raw_buffer[parser.raw_buffer_pos+low+2]) +
						(rune(parser.raw_buffer[parser.raw_buffer_pos+high+2]) << 8)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:326
					if value2&0xFC00 != 0xDC00 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:326
						_go_fuzz_dep_.CoverTab[126610]++
															return yaml_parser_set_reader_error(parser,
							"expected low surrogate area",
							parser.offset+2, int(value2))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:329
						// _ = "end of CoverTab[126610]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:330
						_go_fuzz_dep_.CoverTab[126611]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:330
						// _ = "end of CoverTab[126611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:330
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:330
					// _ = "end of CoverTab[126603]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:330
					_go_fuzz_dep_.CoverTab[126604]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:333
					value = 0x10000 + ((value & 0x3FF) << 10) + (value2 & 0x3FF)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:333
					// _ = "end of CoverTab[126604]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:334
					_go_fuzz_dep_.CoverTab[126612]++
														width = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:335
					// _ = "end of CoverTab[126612]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:336
				// _ = "end of CoverTab[126560]"

			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:338
				_go_fuzz_dep_.CoverTab[126561]++
													panic("impossible")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:339
				// _ = "end of CoverTab[126561]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:340
			// _ = "end of CoverTab[126547]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:340
			_go_fuzz_dep_.CoverTab[126548]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:346
			switch {
			case value == 0x09:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:347
				_go_fuzz_dep_.CoverTab[126613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:347
				// _ = "end of CoverTab[126613]"
			case value == 0x0A:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:348
				_go_fuzz_dep_.CoverTab[126614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:348
				// _ = "end of CoverTab[126614]"
			case value == 0x0D:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:349
				_go_fuzz_dep_.CoverTab[126615]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:349
				// _ = "end of CoverTab[126615]"
			case value >= 0x20 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
				_go_fuzz_dep_.CoverTab[126622]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
				return value <= 0x7E
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
				// _ = "end of CoverTab[126622]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
			}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
				_go_fuzz_dep_.CoverTab[126616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:350
				// _ = "end of CoverTab[126616]"
			case value == 0x85:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:351
				_go_fuzz_dep_.CoverTab[126617]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:351
				// _ = "end of CoverTab[126617]"
			case value >= 0xA0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
				_go_fuzz_dep_.CoverTab[126623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
				return value <= 0xD7FF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
				// _ = "end of CoverTab[126623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
			}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
				_go_fuzz_dep_.CoverTab[126618]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:352
				// _ = "end of CoverTab[126618]"
			case value >= 0xE000 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
				_go_fuzz_dep_.CoverTab[126624]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
				return value <= 0xFFFD
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
				// _ = "end of CoverTab[126624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
			}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
				_go_fuzz_dep_.CoverTab[126619]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:353
				// _ = "end of CoverTab[126619]"
			case value >= 0x10000 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
				_go_fuzz_dep_.CoverTab[126625]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
				return value <= 0x10FFFF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
				// _ = "end of CoverTab[126625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
			}():
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
				_go_fuzz_dep_.CoverTab[126620]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:354
				// _ = "end of CoverTab[126620]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:355
				_go_fuzz_dep_.CoverTab[126621]++
													return yaml_parser_set_reader_error(parser,
					"control characters are not allowed",
					parser.offset, int(value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:358
				// _ = "end of CoverTab[126621]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:359
			// _ = "end of CoverTab[126548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:359
			_go_fuzz_dep_.CoverTab[126549]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:362
			parser.raw_buffer_pos += width
												parser.offset += width

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:366
			if value <= 0x7F {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:366
				_go_fuzz_dep_.CoverTab[126626]++

													parser.buffer[buffer_len+0] = byte(value)
													buffer_len += 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:369
				// _ = "end of CoverTab[126626]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:370
				_go_fuzz_dep_.CoverTab[126627]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:370
				if value <= 0x7FF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:370
					_go_fuzz_dep_.CoverTab[126628]++

														parser.buffer[buffer_len+0] = byte(0xC0 + (value >> 6))
														parser.buffer[buffer_len+1] = byte(0x80 + (value & 0x3F))
														buffer_len += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:374
					// _ = "end of CoverTab[126628]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:375
					_go_fuzz_dep_.CoverTab[126629]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:375
					if value <= 0xFFFF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:375
						_go_fuzz_dep_.CoverTab[126630]++

															parser.buffer[buffer_len+0] = byte(0xE0 + (value >> 12))
															parser.buffer[buffer_len+1] = byte(0x80 + ((value >> 6) & 0x3F))
															parser.buffer[buffer_len+2] = byte(0x80 + (value & 0x3F))
															buffer_len += 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:380
						// _ = "end of CoverTab[126630]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:381
						_go_fuzz_dep_.CoverTab[126631]++

															parser.buffer[buffer_len+0] = byte(0xF0 + (value >> 18))
															parser.buffer[buffer_len+1] = byte(0x80 + ((value >> 12) & 0x3F))
															parser.buffer[buffer_len+2] = byte(0x80 + ((value >> 6) & 0x3F))
															parser.buffer[buffer_len+3] = byte(0x80 + (value & 0x3F))
															buffer_len += 4
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:387
						// _ = "end of CoverTab[126631]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
					// _ = "end of CoverTab[126629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
				// _ = "end of CoverTab[126627]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
			// _ = "end of CoverTab[126549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:388
			_go_fuzz_dep_.CoverTab[126550]++

												parser.unread++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:390
			// _ = "end of CoverTab[126550]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:391
		// _ = "end of CoverTab[126540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:391
		_go_fuzz_dep_.CoverTab[126541]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:394
		if parser.eof {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:394
			_go_fuzz_dep_.CoverTab[126632]++
												parser.buffer[buffer_len] = 0
												buffer_len++
												parser.unread++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:398
			// _ = "end of CoverTab[126632]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:399
			_go_fuzz_dep_.CoverTab[126633]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:399
			// _ = "end of CoverTab[126633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:399
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:399
		// _ = "end of CoverTab[126541]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:400
	// _ = "end of CoverTab[126520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:400
	_go_fuzz_dep_.CoverTab[126521]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:406
	for buffer_len < length {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:406
		_go_fuzz_dep_.CoverTab[126634]++
											parser.buffer[buffer_len] = 0
											buffer_len++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:408
		// _ = "end of CoverTab[126634]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:409
	// _ = "end of CoverTab[126521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:409
	_go_fuzz_dep_.CoverTab[126522]++
										parser.buffer = parser.buffer[:buffer_len]
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:411
	// _ = "end of CoverTab[126522]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:412
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/readerc.go:412
var _ = _go_fuzz_dep_.CoverTab
