//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:1
)

import (
	"io"
)

func yaml_insert_token(parser *yaml_parser_t, pos int, token *yaml_token_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:7
	_go_fuzz_dep_.CoverTab[124441]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
	if parser.tokens_head > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
		_go_fuzz_dep_.CoverTab[124444]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
		return len(parser.tokens) == cap(parser.tokens)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
		// _ = "end of CoverTab[124444]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:11
		_go_fuzz_dep_.CoverTab[124445]++
										if parser.tokens_head != len(parser.tokens) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:12
			_go_fuzz_dep_.CoverTab[124447]++
											copy(parser.tokens, parser.tokens[parser.tokens_head:])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:13
			// _ = "end of CoverTab[124447]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:14
			_go_fuzz_dep_.CoverTab[124448]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:14
			// _ = "end of CoverTab[124448]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:14
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:14
		// _ = "end of CoverTab[124445]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:14
		_go_fuzz_dep_.CoverTab[124446]++
										parser.tokens = parser.tokens[:len(parser.tokens)-parser.tokens_head]
										parser.tokens_head = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:16
		// _ = "end of CoverTab[124446]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:17
		_go_fuzz_dep_.CoverTab[124449]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:17
		// _ = "end of CoverTab[124449]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:17
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:17
	// _ = "end of CoverTab[124441]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:17
	_go_fuzz_dep_.CoverTab[124442]++
									parser.tokens = append(parser.tokens, *token)
									if pos < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:19
		_go_fuzz_dep_.CoverTab[124450]++
										return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:20
		// _ = "end of CoverTab[124450]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:21
		_go_fuzz_dep_.CoverTab[124451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:21
		// _ = "end of CoverTab[124451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:21
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:21
	// _ = "end of CoverTab[124442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:21
	_go_fuzz_dep_.CoverTab[124443]++
									copy(parser.tokens[parser.tokens_head+pos+1:], parser.tokens[parser.tokens_head+pos:])
									parser.tokens[parser.tokens_head+pos] = *token
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:23
	// _ = "end of CoverTab[124443]"
}

// Create a new parser object.
func yaml_parser_initialize(parser *yaml_parser_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:27
	_go_fuzz_dep_.CoverTab[124452]++
									*parser = yaml_parser_t{
		raw_buffer:	make([]byte, 0, input_raw_buffer_size),
		buffer:		make([]byte, 0, input_buffer_size),
	}
									return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:32
	// _ = "end of CoverTab[124452]"
}

// Destroy a parser object.
func yaml_parser_delete(parser *yaml_parser_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:36
	_go_fuzz_dep_.CoverTab[124453]++
									*parser = yaml_parser_t{}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:37
	// _ = "end of CoverTab[124453]"
}

// String read handler.
func yaml_string_read_handler(parser *yaml_parser_t, buffer []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:41
	_go_fuzz_dep_.CoverTab[124454]++
									if parser.input_pos == len(parser.input) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:42
		_go_fuzz_dep_.CoverTab[124456]++
										return 0, io.EOF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:43
		// _ = "end of CoverTab[124456]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:44
		_go_fuzz_dep_.CoverTab[124457]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:44
		// _ = "end of CoverTab[124457]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:44
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:44
	// _ = "end of CoverTab[124454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:44
	_go_fuzz_dep_.CoverTab[124455]++
									n = copy(buffer, parser.input[parser.input_pos:])
									parser.input_pos += n
									return n, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:47
	// _ = "end of CoverTab[124455]"
}

// Reader read handler.
func yaml_reader_read_handler(parser *yaml_parser_t, buffer []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:51
	_go_fuzz_dep_.CoverTab[124458]++
									return parser.input_reader.Read(buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:52
	// _ = "end of CoverTab[124458]"
}

// Set a string input.
func yaml_parser_set_input_string(parser *yaml_parser_t, input []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:56
	_go_fuzz_dep_.CoverTab[124459]++
									if parser.read_handler != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:57
		_go_fuzz_dep_.CoverTab[124461]++
										panic("must set the input source only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:58
		// _ = "end of CoverTab[124461]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:59
		_go_fuzz_dep_.CoverTab[124462]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:59
		// _ = "end of CoverTab[124462]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:59
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:59
	// _ = "end of CoverTab[124459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:59
	_go_fuzz_dep_.CoverTab[124460]++
									parser.read_handler = yaml_string_read_handler
									parser.input = input
									parser.input_pos = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:62
	// _ = "end of CoverTab[124460]"
}

// Set a file input.
func yaml_parser_set_input_reader(parser *yaml_parser_t, r io.Reader) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:66
	_go_fuzz_dep_.CoverTab[124463]++
									if parser.read_handler != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:67
		_go_fuzz_dep_.CoverTab[124465]++
										panic("must set the input source only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:68
		// _ = "end of CoverTab[124465]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:69
		_go_fuzz_dep_.CoverTab[124466]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:69
		// _ = "end of CoverTab[124466]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:69
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:69
	// _ = "end of CoverTab[124463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:69
	_go_fuzz_dep_.CoverTab[124464]++
									parser.read_handler = yaml_reader_read_handler
									parser.input_reader = r
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:71
	// _ = "end of CoverTab[124464]"
}

// Set the source encoding.
func yaml_parser_set_encoding(parser *yaml_parser_t, encoding yaml_encoding_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:75
	_go_fuzz_dep_.CoverTab[124467]++
									if parser.encoding != yaml_ANY_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:76
		_go_fuzz_dep_.CoverTab[124469]++
										panic("must set the encoding only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:77
		// _ = "end of CoverTab[124469]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:78
		_go_fuzz_dep_.CoverTab[124470]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:78
		// _ = "end of CoverTab[124470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:78
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:78
	// _ = "end of CoverTab[124467]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:78
	_go_fuzz_dep_.CoverTab[124468]++
									parser.encoding = encoding
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:79
	// _ = "end of CoverTab[124468]"
}

var disableLineWrapping = false

// Create a new emitter object.
func yaml_emitter_initialize(emitter *yaml_emitter_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:85
	_go_fuzz_dep_.CoverTab[124471]++
									*emitter = yaml_emitter_t{
		buffer:		make([]byte, output_buffer_size),
		raw_buffer:	make([]byte, 0, output_raw_buffer_size),
		states:		make([]yaml_emitter_state_t, 0, initial_stack_size),
		events:		make([]yaml_event_t, 0, initial_queue_size),
	}
	if disableLineWrapping {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:92
		_go_fuzz_dep_.CoverTab[124472]++
										emitter.best_width = -1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:93
		// _ = "end of CoverTab[124472]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:94
		_go_fuzz_dep_.CoverTab[124473]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:94
		// _ = "end of CoverTab[124473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:94
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:94
	// _ = "end of CoverTab[124471]"
}

// Destroy an emitter object.
func yaml_emitter_delete(emitter *yaml_emitter_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:98
	_go_fuzz_dep_.CoverTab[124474]++
									*emitter = yaml_emitter_t{}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:99
	// _ = "end of CoverTab[124474]"
}

// String write handler.
func yaml_string_write_handler(emitter *yaml_emitter_t, buffer []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:103
	_go_fuzz_dep_.CoverTab[124475]++
									*emitter.output_buffer = append(*emitter.output_buffer, buffer...)
									return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:105
	// _ = "end of CoverTab[124475]"
}

// yaml_writer_write_handler uses emitter.output_writer to write the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:108
// emitted text.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:110
func yaml_writer_write_handler(emitter *yaml_emitter_t, buffer []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:110
	_go_fuzz_dep_.CoverTab[124476]++
									_, err := emitter.output_writer.Write(buffer)
									return err
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:112
	// _ = "end of CoverTab[124476]"
}

// Set a string output.
func yaml_emitter_set_output_string(emitter *yaml_emitter_t, output_buffer *[]byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:116
	_go_fuzz_dep_.CoverTab[124477]++
									if emitter.write_handler != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:117
		_go_fuzz_dep_.CoverTab[124479]++
										panic("must set the output target only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:118
		// _ = "end of CoverTab[124479]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:119
		_go_fuzz_dep_.CoverTab[124480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:119
		// _ = "end of CoverTab[124480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:119
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:119
	// _ = "end of CoverTab[124477]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:119
	_go_fuzz_dep_.CoverTab[124478]++
									emitter.write_handler = yaml_string_write_handler
									emitter.output_buffer = output_buffer
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:121
	// _ = "end of CoverTab[124478]"
}

// Set a file output.
func yaml_emitter_set_output_writer(emitter *yaml_emitter_t, w io.Writer) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:125
	_go_fuzz_dep_.CoverTab[124481]++
									if emitter.write_handler != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:126
		_go_fuzz_dep_.CoverTab[124483]++
										panic("must set the output target only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:127
		// _ = "end of CoverTab[124483]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:128
		_go_fuzz_dep_.CoverTab[124484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:128
		// _ = "end of CoverTab[124484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:128
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:128
	// _ = "end of CoverTab[124481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:128
	_go_fuzz_dep_.CoverTab[124482]++
									emitter.write_handler = yaml_writer_write_handler
									emitter.output_writer = w
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:130
	// _ = "end of CoverTab[124482]"
}

// Set the output encoding.
func yaml_emitter_set_encoding(emitter *yaml_emitter_t, encoding yaml_encoding_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:134
	_go_fuzz_dep_.CoverTab[124485]++
									if emitter.encoding != yaml_ANY_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:135
		_go_fuzz_dep_.CoverTab[124487]++
										panic("must set the output encoding only once")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:136
		// _ = "end of CoverTab[124487]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:137
		_go_fuzz_dep_.CoverTab[124488]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:137
		// _ = "end of CoverTab[124488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:137
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:137
	// _ = "end of CoverTab[124485]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:137
	_go_fuzz_dep_.CoverTab[124486]++
									emitter.encoding = encoding
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:138
	// _ = "end of CoverTab[124486]"
}

// Set the canonical output style.
func yaml_emitter_set_canonical(emitter *yaml_emitter_t, canonical bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:142
	_go_fuzz_dep_.CoverTab[124489]++
									emitter.canonical = canonical
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:143
	// _ = "end of CoverTab[124489]"
}

// // Set the indentation increment.
func yaml_emitter_set_indent(emitter *yaml_emitter_t, indent int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:147
	_go_fuzz_dep_.CoverTab[124490]++
									if indent < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:148
		_go_fuzz_dep_.CoverTab[124492]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:148
		return indent > 9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:148
		// _ = "end of CoverTab[124492]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:148
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:148
		_go_fuzz_dep_.CoverTab[124493]++
										indent = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:149
		// _ = "end of CoverTab[124493]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:150
		_go_fuzz_dep_.CoverTab[124494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:150
		// _ = "end of CoverTab[124494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:150
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:150
	// _ = "end of CoverTab[124490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:150
	_go_fuzz_dep_.CoverTab[124491]++
									emitter.best_indent = indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:151
	// _ = "end of CoverTab[124491]"
}

// Set the preferred line width.
func yaml_emitter_set_width(emitter *yaml_emitter_t, width int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:155
	_go_fuzz_dep_.CoverTab[124495]++
									if width < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:156
		_go_fuzz_dep_.CoverTab[124497]++
										width = -1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:157
		// _ = "end of CoverTab[124497]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:158
		_go_fuzz_dep_.CoverTab[124498]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:158
		// _ = "end of CoverTab[124498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:158
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:158
	// _ = "end of CoverTab[124495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:158
	_go_fuzz_dep_.CoverTab[124496]++
									emitter.best_width = width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:159
	// _ = "end of CoverTab[124496]"
}

// Set if unescaped non-ASCII characters are allowed.
func yaml_emitter_set_unicode(emitter *yaml_emitter_t, unicode bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:163
	_go_fuzz_dep_.CoverTab[124499]++
									emitter.unicode = unicode
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:164
	// _ = "end of CoverTab[124499]"
}

// Set the preferred line break character.
func yaml_emitter_set_break(emitter *yaml_emitter_t, line_break yaml_break_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:168
	_go_fuzz_dep_.CoverTab[124500]++
									emitter.line_break = line_break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:169
	// _ = "end of CoverTab[124500]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:258
// Create STREAM-START.
func yaml_stream_start_event_initialize(event *yaml_event_t, encoding yaml_encoding_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:259
	_go_fuzz_dep_.CoverTab[124501]++
									*event = yaml_event_t{
		typ:		yaml_STREAM_START_EVENT,
		encoding:	encoding,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:263
	// _ = "end of CoverTab[124501]"
}

// Create STREAM-END.
func yaml_stream_end_event_initialize(event *yaml_event_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:267
	_go_fuzz_dep_.CoverTab[124502]++
									*event = yaml_event_t{
		typ: yaml_STREAM_END_EVENT,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:270
	// _ = "end of CoverTab[124502]"
}

// Create DOCUMENT-START.
func yaml_document_start_event_initialize(
	event *yaml_event_t,
	version_directive *yaml_version_directive_t,
	tag_directives []yaml_tag_directive_t,
	implicit bool,
) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:279
	_go_fuzz_dep_.CoverTab[124503]++
									*event = yaml_event_t{
		typ:			yaml_DOCUMENT_START_EVENT,
		version_directive:	version_directive,
		tag_directives:		tag_directives,
		implicit:		implicit,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:285
	// _ = "end of CoverTab[124503]"
}

// Create DOCUMENT-END.
func yaml_document_end_event_initialize(event *yaml_event_t, implicit bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:289
	_go_fuzz_dep_.CoverTab[124504]++
									*event = yaml_event_t{
		typ:		yaml_DOCUMENT_END_EVENT,
		implicit:	implicit,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:293
	// _ = "end of CoverTab[124504]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:320
// Create SCALAR.
func yaml_scalar_event_initialize(event *yaml_event_t, anchor, tag, value []byte, plain_implicit, quoted_implicit bool, style yaml_scalar_style_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:321
	_go_fuzz_dep_.CoverTab[124505]++
									*event = yaml_event_t{
		typ:			yaml_SCALAR_EVENT,
		anchor:			anchor,
		tag:			tag,
		value:			value,
		implicit:		plain_implicit,
		quoted_implicit:	quoted_implicit,
		style:			yaml_style_t(style),
	}
									return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:331
	// _ = "end of CoverTab[124505]"
}

// Create SEQUENCE-START.
func yaml_sequence_start_event_initialize(event *yaml_event_t, anchor, tag []byte, implicit bool, style yaml_sequence_style_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:335
	_go_fuzz_dep_.CoverTab[124506]++
									*event = yaml_event_t{
		typ:		yaml_SEQUENCE_START_EVENT,
		anchor:		anchor,
		tag:		tag,
		implicit:	implicit,
		style:		yaml_style_t(style),
	}
									return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:343
	// _ = "end of CoverTab[124506]"
}

// Create SEQUENCE-END.
func yaml_sequence_end_event_initialize(event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:347
	_go_fuzz_dep_.CoverTab[124507]++
									*event = yaml_event_t{
		typ: yaml_SEQUENCE_END_EVENT,
	}
									return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:351
	// _ = "end of CoverTab[124507]"
}

// Create MAPPING-START.
func yaml_mapping_start_event_initialize(event *yaml_event_t, anchor, tag []byte, implicit bool, style yaml_mapping_style_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:355
	_go_fuzz_dep_.CoverTab[124508]++
									*event = yaml_event_t{
		typ:		yaml_MAPPING_START_EVENT,
		anchor:		anchor,
		tag:		tag,
		implicit:	implicit,
		style:		yaml_style_t(style),
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:362
	// _ = "end of CoverTab[124508]"
}

// Create MAPPING-END.
func yaml_mapping_end_event_initialize(event *yaml_event_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:366
	_go_fuzz_dep_.CoverTab[124509]++
									*event = yaml_event_t{
		typ: yaml_MAPPING_END_EVENT,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:369
	// _ = "end of CoverTab[124509]"
}

// Destroy an event object.
func yaml_event_delete(event *yaml_event_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:373
	_go_fuzz_dep_.CoverTab[124510]++
									*event = yaml_event_t{}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:374
	// _ = "end of CoverTab[124510]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:375
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/apic.go:375
var _ = _go_fuzz_dep_.CoverTab
