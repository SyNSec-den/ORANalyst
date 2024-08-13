//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1
)

import (
	"bytes"
)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:45
// Peek the next token in the token queue.
func peek_token(parser *yaml_parser_t) *yaml_token_t {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:46
	_go_fuzz_dep_.CoverTab[126106]++
										if parser.token_available || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:47
		_go_fuzz_dep_.CoverTab[126108]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:47
		return yaml_parser_fetch_more_tokens(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:47
		// _ = "end of CoverTab[126108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:47
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:47
		_go_fuzz_dep_.CoverTab[126109]++
											return &parser.tokens[parser.tokens_head]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:48
		// _ = "end of CoverTab[126109]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:49
		_go_fuzz_dep_.CoverTab[126110]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:49
		// _ = "end of CoverTab[126110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:49
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:49
	// _ = "end of CoverTab[126106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:49
	_go_fuzz_dep_.CoverTab[126107]++
										return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:50
	// _ = "end of CoverTab[126107]"
}

// Remove the next token from the queue (must be called after peek_token).
func skip_token(parser *yaml_parser_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:54
	_go_fuzz_dep_.CoverTab[126111]++
										parser.token_available = false
										parser.tokens_parsed++
										parser.stream_end_produced = parser.tokens[parser.tokens_head].typ == yaml_STREAM_END_TOKEN
										parser.tokens_head++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:58
	// _ = "end of CoverTab[126111]"
}

// Get the next event.
func yaml_parser_parse(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:62
	_go_fuzz_dep_.CoverTab[126112]++

										*event = yaml_event_t{}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
	if parser.stream_end_produced || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		_go_fuzz_dep_.CoverTab[126114]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		return parser.error != yaml_NO_ERROR
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		// _ = "end of CoverTab[126114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		_go_fuzz_dep_.CoverTab[126115]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		return parser.state == yaml_PARSE_END_STATE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		// _ = "end of CoverTab[126115]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:67
		_go_fuzz_dep_.CoverTab[126116]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:68
		// _ = "end of CoverTab[126116]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:69
		_go_fuzz_dep_.CoverTab[126117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:69
		// _ = "end of CoverTab[126117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:69
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:69
	// _ = "end of CoverTab[126112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:69
	_go_fuzz_dep_.CoverTab[126113]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:72
	return yaml_parser_state_machine(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:72
	// _ = "end of CoverTab[126113]"
}

// Set parser error.
func yaml_parser_set_parser_error(parser *yaml_parser_t, problem string, problem_mark yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:76
	_go_fuzz_dep_.CoverTab[126118]++
										parser.error = yaml_PARSER_ERROR
										parser.problem = problem
										parser.problem_mark = problem_mark
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:80
	// _ = "end of CoverTab[126118]"
}

func yaml_parser_set_parser_error_context(parser *yaml_parser_t, context string, context_mark yaml_mark_t, problem string, problem_mark yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:83
	_go_fuzz_dep_.CoverTab[126119]++
										parser.error = yaml_PARSER_ERROR
										parser.context = context
										parser.context_mark = context_mark
										parser.problem = problem
										parser.problem_mark = problem_mark
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:89
	// _ = "end of CoverTab[126119]"
}

// State dispatcher.
func yaml_parser_state_machine(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:93
	_go_fuzz_dep_.CoverTab[126120]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:96
	switch parser.state {
	case yaml_PARSE_STREAM_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:97
		_go_fuzz_dep_.CoverTab[126121]++
											return yaml_parser_parse_stream_start(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:98
		// _ = "end of CoverTab[126121]"

	case yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:100
		_go_fuzz_dep_.CoverTab[126122]++
											return yaml_parser_parse_document_start(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:101
		// _ = "end of CoverTab[126122]"

	case yaml_PARSE_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:103
		_go_fuzz_dep_.CoverTab[126123]++
											return yaml_parser_parse_document_start(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:104
		// _ = "end of CoverTab[126123]"

	case yaml_PARSE_DOCUMENT_CONTENT_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:106
		_go_fuzz_dep_.CoverTab[126124]++
											return yaml_parser_parse_document_content(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:107
		// _ = "end of CoverTab[126124]"

	case yaml_PARSE_DOCUMENT_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:109
		_go_fuzz_dep_.CoverTab[126125]++
											return yaml_parser_parse_document_end(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:110
		// _ = "end of CoverTab[126125]"

	case yaml_PARSE_BLOCK_NODE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:112
		_go_fuzz_dep_.CoverTab[126126]++
											return yaml_parser_parse_node(parser, event, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:113
		// _ = "end of CoverTab[126126]"

	case yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:115
		_go_fuzz_dep_.CoverTab[126127]++
											return yaml_parser_parse_node(parser, event, true, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:116
		// _ = "end of CoverTab[126127]"

	case yaml_PARSE_FLOW_NODE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:118
		_go_fuzz_dep_.CoverTab[126128]++
											return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:119
		// _ = "end of CoverTab[126128]"

	case yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:121
		_go_fuzz_dep_.CoverTab[126129]++
											return yaml_parser_parse_block_sequence_entry(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:122
		// _ = "end of CoverTab[126129]"

	case yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:124
		_go_fuzz_dep_.CoverTab[126130]++
											return yaml_parser_parse_block_sequence_entry(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:125
		// _ = "end of CoverTab[126130]"

	case yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:127
		_go_fuzz_dep_.CoverTab[126131]++
											return yaml_parser_parse_indentless_sequence_entry(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:128
		// _ = "end of CoverTab[126131]"

	case yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:130
		_go_fuzz_dep_.CoverTab[126132]++
											return yaml_parser_parse_block_mapping_key(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:131
		// _ = "end of CoverTab[126132]"

	case yaml_PARSE_BLOCK_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:133
		_go_fuzz_dep_.CoverTab[126133]++
											return yaml_parser_parse_block_mapping_key(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:134
		// _ = "end of CoverTab[126133]"

	case yaml_PARSE_BLOCK_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:136
		_go_fuzz_dep_.CoverTab[126134]++
											return yaml_parser_parse_block_mapping_value(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:137
		// _ = "end of CoverTab[126134]"

	case yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:139
		_go_fuzz_dep_.CoverTab[126135]++
											return yaml_parser_parse_flow_sequence_entry(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:140
		// _ = "end of CoverTab[126135]"

	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:142
		_go_fuzz_dep_.CoverTab[126136]++
											return yaml_parser_parse_flow_sequence_entry(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:143
		// _ = "end of CoverTab[126136]"

	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:145
		_go_fuzz_dep_.CoverTab[126137]++
											return yaml_parser_parse_flow_sequence_entry_mapping_key(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:146
		// _ = "end of CoverTab[126137]"

	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:148
		_go_fuzz_dep_.CoverTab[126138]++
											return yaml_parser_parse_flow_sequence_entry_mapping_value(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:149
		// _ = "end of CoverTab[126138]"

	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:151
		_go_fuzz_dep_.CoverTab[126139]++
											return yaml_parser_parse_flow_sequence_entry_mapping_end(parser, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:152
		// _ = "end of CoverTab[126139]"

	case yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:154
		_go_fuzz_dep_.CoverTab[126140]++
											return yaml_parser_parse_flow_mapping_key(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:155
		// _ = "end of CoverTab[126140]"

	case yaml_PARSE_FLOW_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:157
		_go_fuzz_dep_.CoverTab[126141]++
											return yaml_parser_parse_flow_mapping_key(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:158
		// _ = "end of CoverTab[126141]"

	case yaml_PARSE_FLOW_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:160
		_go_fuzz_dep_.CoverTab[126142]++
											return yaml_parser_parse_flow_mapping_value(parser, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:161
		// _ = "end of CoverTab[126142]"

	case yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:163
		_go_fuzz_dep_.CoverTab[126143]++
											return yaml_parser_parse_flow_mapping_value(parser, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:164
		// _ = "end of CoverTab[126143]"

	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:166
		_go_fuzz_dep_.CoverTab[126144]++
											panic("invalid parser state")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:167
		// _ = "end of CoverTab[126144]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:168
	// _ = "end of CoverTab[126120]"
}

// Parse the production:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:171
// stream   ::= STREAM-START implicit_document? explicit_document* STREAM-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:171
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:171
//	************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:174
func yaml_parser_parse_stream_start(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:174
	_go_fuzz_dep_.CoverTab[126145]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:176
		_go_fuzz_dep_.CoverTab[126148]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:177
		// _ = "end of CoverTab[126148]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:178
		_go_fuzz_dep_.CoverTab[126149]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:178
		// _ = "end of CoverTab[126149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:178
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:178
	// _ = "end of CoverTab[126145]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:178
	_go_fuzz_dep_.CoverTab[126146]++
										if token.typ != yaml_STREAM_START_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:179
		_go_fuzz_dep_.CoverTab[126150]++
											return yaml_parser_set_parser_error(parser, "did not find expected <stream-start>", token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:180
		// _ = "end of CoverTab[126150]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:181
		_go_fuzz_dep_.CoverTab[126151]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:181
		// _ = "end of CoverTab[126151]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:181
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:181
	// _ = "end of CoverTab[126146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:181
	_go_fuzz_dep_.CoverTab[126147]++
										parser.state = yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE
										*event = yaml_event_t{
		typ:		yaml_STREAM_START_EVENT,
		start_mark:	token.start_mark,
		end_mark:	token.end_mark,
		encoding:	token.encoding,
	}
										skip_token(parser)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:190
	// _ = "end of CoverTab[126147]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
// implicit_document    ::= block_node DOCUMENT-END*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
//	*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
// explicit_document    ::= DIRECTIVE* DOCUMENT-START block_node? DOCUMENT-END*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:193
//	*************************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:198
func yaml_parser_parse_document_start(parser *yaml_parser_t, event *yaml_event_t, implicit bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:198
	_go_fuzz_dep_.CoverTab[126152]++

										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:201
		_go_fuzz_dep_.CoverTab[126156]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:202
		// _ = "end of CoverTab[126156]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:203
		_go_fuzz_dep_.CoverTab[126157]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:203
		// _ = "end of CoverTab[126157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:203
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:203
	// _ = "end of CoverTab[126152]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:203
	_go_fuzz_dep_.CoverTab[126153]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:206
	if !implicit {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:206
		_go_fuzz_dep_.CoverTab[126158]++
											for token.typ == yaml_DOCUMENT_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:207
			_go_fuzz_dep_.CoverTab[126159]++
												skip_token(parser)
												token = peek_token(parser)
												if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:210
				_go_fuzz_dep_.CoverTab[126160]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:211
				// _ = "end of CoverTab[126160]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:212
				_go_fuzz_dep_.CoverTab[126161]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:212
				// _ = "end of CoverTab[126161]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:212
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:212
			// _ = "end of CoverTab[126159]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:213
		// _ = "end of CoverTab[126158]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:214
		_go_fuzz_dep_.CoverTab[126162]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:214
		// _ = "end of CoverTab[126162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:214
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:214
	// _ = "end of CoverTab[126153]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:214
	_go_fuzz_dep_.CoverTab[126154]++

										if implicit && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
		_go_fuzz_dep_.CoverTab[126163]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
		return token.typ != yaml_VERSION_DIRECTIVE_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
		// _ = "end of CoverTab[126163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
		_go_fuzz_dep_.CoverTab[126164]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:216
		return token.typ != yaml_TAG_DIRECTIVE_TOKEN
											// _ = "end of CoverTab[126164]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:217
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:217
		_go_fuzz_dep_.CoverTab[126165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:217
		return token.typ != yaml_DOCUMENT_START_TOKEN
											// _ = "end of CoverTab[126165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:218
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:218
		_go_fuzz_dep_.CoverTab[126166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:218
		return token.typ != yaml_STREAM_END_TOKEN
											// _ = "end of CoverTab[126166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:219
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:219
		_go_fuzz_dep_.CoverTab[126167]++

											if !yaml_parser_process_directives(parser, nil, nil) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:221
			_go_fuzz_dep_.CoverTab[126169]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:222
			// _ = "end of CoverTab[126169]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:223
			_go_fuzz_dep_.CoverTab[126170]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:223
			// _ = "end of CoverTab[126170]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:223
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:223
		// _ = "end of CoverTab[126167]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:223
		_go_fuzz_dep_.CoverTab[126168]++
											parser.states = append(parser.states, yaml_PARSE_DOCUMENT_END_STATE)
											parser.state = yaml_PARSE_BLOCK_NODE_STATE

											*event = yaml_event_t{
			typ:		yaml_DOCUMENT_START_EVENT,
			start_mark:	token.start_mark,
			end_mark:	token.end_mark,
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:231
		// _ = "end of CoverTab[126168]"

	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:233
		_go_fuzz_dep_.CoverTab[126171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:233
		if token.typ != yaml_STREAM_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:233
			_go_fuzz_dep_.CoverTab[126172]++
			// Parse an explicit document.
			var version_directive *yaml_version_directive_t
			var tag_directives []yaml_tag_directive_t
			start_mark := token.start_mark
			if !yaml_parser_process_directives(parser, &version_directive, &tag_directives) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:238
				_go_fuzz_dep_.CoverTab[126176]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:239
				// _ = "end of CoverTab[126176]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:240
				_go_fuzz_dep_.CoverTab[126177]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:240
				// _ = "end of CoverTab[126177]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:240
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:240
			// _ = "end of CoverTab[126172]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:240
			_go_fuzz_dep_.CoverTab[126173]++
												token = peek_token(parser)
												if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:242
				_go_fuzz_dep_.CoverTab[126178]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:243
				// _ = "end of CoverTab[126178]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:244
				_go_fuzz_dep_.CoverTab[126179]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:244
				// _ = "end of CoverTab[126179]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:244
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:244
			// _ = "end of CoverTab[126173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:244
			_go_fuzz_dep_.CoverTab[126174]++
												if token.typ != yaml_DOCUMENT_START_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:245
				_go_fuzz_dep_.CoverTab[126180]++
													yaml_parser_set_parser_error(parser,
					"did not find expected <document start>", token.start_mark)
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:248
				// _ = "end of CoverTab[126180]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:249
				_go_fuzz_dep_.CoverTab[126181]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:249
				// _ = "end of CoverTab[126181]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:249
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:249
			// _ = "end of CoverTab[126174]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:249
			_go_fuzz_dep_.CoverTab[126175]++
												parser.states = append(parser.states, yaml_PARSE_DOCUMENT_END_STATE)
												parser.state = yaml_PARSE_DOCUMENT_CONTENT_STATE
												end_mark := token.end_mark

												*event = yaml_event_t{
				typ:			yaml_DOCUMENT_START_EVENT,
				start_mark:		start_mark,
				end_mark:		end_mark,
				version_directive:	version_directive,
				tag_directives:		tag_directives,
				implicit:		false,
			}
												skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:262
			// _ = "end of CoverTab[126175]"

		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:264
			_go_fuzz_dep_.CoverTab[126182]++

												parser.state = yaml_PARSE_END_STATE
												*event = yaml_event_t{
				typ:		yaml_STREAM_END_EVENT,
				start_mark:	token.start_mark,
				end_mark:	token.end_mark,
			}
												skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:272
			// _ = "end of CoverTab[126182]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:273
		// _ = "end of CoverTab[126171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:273
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:273
	// _ = "end of CoverTab[126154]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:273
	_go_fuzz_dep_.CoverTab[126155]++

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:275
	// _ = "end of CoverTab[126155]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:278
// explicit_document    ::= DIRECTIVE* DOCUMENT-START block_node? DOCUMENT-END*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:278
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:278
//	***********
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:282
func yaml_parser_parse_document_content(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:282
	_go_fuzz_dep_.CoverTab[126183]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:284
		_go_fuzz_dep_.CoverTab[126186]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:285
		// _ = "end of CoverTab[126186]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:286
		_go_fuzz_dep_.CoverTab[126187]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:286
		// _ = "end of CoverTab[126187]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:286
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:286
	// _ = "end of CoverTab[126183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:286
	_go_fuzz_dep_.CoverTab[126184]++
										if token.typ == yaml_VERSION_DIRECTIVE_TOKEN || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:287
		_go_fuzz_dep_.CoverTab[126188]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:287
		return token.typ == yaml_TAG_DIRECTIVE_TOKEN
											// _ = "end of CoverTab[126188]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:288
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:288
		_go_fuzz_dep_.CoverTab[126189]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:288
		return token.typ == yaml_DOCUMENT_START_TOKEN
											// _ = "end of CoverTab[126189]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:289
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:289
		_go_fuzz_dep_.CoverTab[126190]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:289
		return token.typ == yaml_DOCUMENT_END_TOKEN
											// _ = "end of CoverTab[126190]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:290
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:290
		_go_fuzz_dep_.CoverTab[126191]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:290
		return token.typ == yaml_STREAM_END_TOKEN
											// _ = "end of CoverTab[126191]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:291
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:291
		_go_fuzz_dep_.CoverTab[126192]++
											parser.state = parser.states[len(parser.states)-1]
											parser.states = parser.states[:len(parser.states)-1]
											return yaml_parser_process_empty_scalar(parser, event,
			token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:295
		// _ = "end of CoverTab[126192]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:296
		_go_fuzz_dep_.CoverTab[126193]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:296
		// _ = "end of CoverTab[126193]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:296
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:296
	// _ = "end of CoverTab[126184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:296
	_go_fuzz_dep_.CoverTab[126185]++
										return yaml_parser_parse_node(parser, event, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:297
	// _ = "end of CoverTab[126185]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:300
// implicit_document    ::= block_node DOCUMENT-END*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:300
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:300
//	*************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:300
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:300
// explicit_document    ::= DIRECTIVE* DOCUMENT-START block_node? DOCUMENT-END*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:305
func yaml_parser_parse_document_end(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:305
	_go_fuzz_dep_.CoverTab[126194]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:307
		_go_fuzz_dep_.CoverTab[126197]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:308
		// _ = "end of CoverTab[126197]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:309
		_go_fuzz_dep_.CoverTab[126198]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:309
		// _ = "end of CoverTab[126198]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:309
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:309
	// _ = "end of CoverTab[126194]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:309
	_go_fuzz_dep_.CoverTab[126195]++

										start_mark := token.start_mark
										end_mark := token.start_mark

										implicit := true
										if token.typ == yaml_DOCUMENT_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:315
		_go_fuzz_dep_.CoverTab[126199]++
											end_mark = token.end_mark
											skip_token(parser)
											implicit = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:318
		// _ = "end of CoverTab[126199]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:319
		_go_fuzz_dep_.CoverTab[126200]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:319
		// _ = "end of CoverTab[126200]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:319
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:319
	// _ = "end of CoverTab[126195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:319
	_go_fuzz_dep_.CoverTab[126196]++

										parser.tag_directives = parser.tag_directives[:0]

										parser.state = yaml_PARSE_DOCUMENT_START_STATE
										*event = yaml_event_t{
		typ:		yaml_DOCUMENT_END_EVENT,
		start_mark:	start_mark,
		end_mark:	end_mark,
		implicit:	implicit,
	}
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:330
	// _ = "end of CoverTab[126196]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// block_node_or_indentless_sequence    ::=
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	ALIAS
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	*****
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| properties (block_content | indentless_block_sequence)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  **********  *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| block_content | indentless_block_sequence
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// block_node           ::= ALIAS
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	*****
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| properties block_content?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  ********** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| block_content
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// flow_node            ::= ALIAS
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	*****
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| properties flow_content?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  ********** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	| flow_content
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	  *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// properties           ::= TAG ANCHOR? | ANCHOR TAG?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	*************************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// block_content        ::= block_collection | flow_collection | SCALAR
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	******
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
// flow_content         ::= flow_collection | SCALAR
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:333
//	******
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:359
func yaml_parser_parse_node(parser *yaml_parser_t, event *yaml_event_t, block, indentless_sequence bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:359
	_go_fuzz_dep_.CoverTab[126201]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:362
	token := peek_token(parser)
	if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:363
		_go_fuzz_dep_.CoverTab[126214]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:364
		// _ = "end of CoverTab[126214]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:365
		_go_fuzz_dep_.CoverTab[126215]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:365
		// _ = "end of CoverTab[126215]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:365
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:365
	// _ = "end of CoverTab[126201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:365
	_go_fuzz_dep_.CoverTab[126202]++

										if token.typ == yaml_ALIAS_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:367
		_go_fuzz_dep_.CoverTab[126216]++
											parser.state = parser.states[len(parser.states)-1]
											parser.states = parser.states[:len(parser.states)-1]
											*event = yaml_event_t{
			typ:		yaml_ALIAS_EVENT,
			start_mark:	token.start_mark,
			end_mark:	token.end_mark,
			anchor:		token.value,
		}
											skip_token(parser)
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:377
		// _ = "end of CoverTab[126216]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:378
		_go_fuzz_dep_.CoverTab[126217]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:378
		// _ = "end of CoverTab[126217]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:378
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:378
	// _ = "end of CoverTab[126202]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:378
	_go_fuzz_dep_.CoverTab[126203]++

										start_mark := token.start_mark
										end_mark := token.start_mark

										var tag_token bool
										var tag_handle, tag_suffix, anchor []byte
										var tag_mark yaml_mark_t
										if token.typ == yaml_ANCHOR_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:386
		_go_fuzz_dep_.CoverTab[126218]++
											anchor = token.value
											start_mark = token.start_mark
											end_mark = token.end_mark
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:392
			_go_fuzz_dep_.CoverTab[126220]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:393
			// _ = "end of CoverTab[126220]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:394
			_go_fuzz_dep_.CoverTab[126221]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:394
			// _ = "end of CoverTab[126221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:394
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:394
		// _ = "end of CoverTab[126218]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:394
		_go_fuzz_dep_.CoverTab[126219]++
											if token.typ == yaml_TAG_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:395
			_go_fuzz_dep_.CoverTab[126222]++
												tag_token = true
												tag_handle = token.value
												tag_suffix = token.suffix
												tag_mark = token.start_mark
												end_mark = token.end_mark
												skip_token(parser)
												token = peek_token(parser)
												if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:403
				_go_fuzz_dep_.CoverTab[126223]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:404
				// _ = "end of CoverTab[126223]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:405
				_go_fuzz_dep_.CoverTab[126224]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:405
				// _ = "end of CoverTab[126224]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:405
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:405
			// _ = "end of CoverTab[126222]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:406
			_go_fuzz_dep_.CoverTab[126225]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:406
			// _ = "end of CoverTab[126225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:406
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:406
		// _ = "end of CoverTab[126219]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:407
		_go_fuzz_dep_.CoverTab[126226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:407
		if token.typ == yaml_TAG_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:407
			_go_fuzz_dep_.CoverTab[126227]++
												tag_token = true
												tag_handle = token.value
												tag_suffix = token.suffix
												start_mark = token.start_mark
												tag_mark = token.start_mark
												end_mark = token.end_mark
												skip_token(parser)
												token = peek_token(parser)
												if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:416
				_go_fuzz_dep_.CoverTab[126229]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:417
				// _ = "end of CoverTab[126229]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:418
				_go_fuzz_dep_.CoverTab[126230]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:418
				// _ = "end of CoverTab[126230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:418
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:418
			// _ = "end of CoverTab[126227]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:418
			_go_fuzz_dep_.CoverTab[126228]++
												if token.typ == yaml_ANCHOR_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:419
				_go_fuzz_dep_.CoverTab[126231]++
													anchor = token.value
													end_mark = token.end_mark
													skip_token(parser)
													token = peek_token(parser)
													if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:424
					_go_fuzz_dep_.CoverTab[126232]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:425
					// _ = "end of CoverTab[126232]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:426
					_go_fuzz_dep_.CoverTab[126233]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:426
					// _ = "end of CoverTab[126233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:426
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:426
				// _ = "end of CoverTab[126231]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:427
				_go_fuzz_dep_.CoverTab[126234]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:427
				// _ = "end of CoverTab[126234]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:427
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:427
			// _ = "end of CoverTab[126228]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
			_go_fuzz_dep_.CoverTab[126235]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
			// _ = "end of CoverTab[126235]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
		// _ = "end of CoverTab[126226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
	// _ = "end of CoverTab[126203]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:428
	_go_fuzz_dep_.CoverTab[126204]++

										var tag []byte
										if tag_token {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:431
		_go_fuzz_dep_.CoverTab[126236]++
											if len(tag_handle) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:432
			_go_fuzz_dep_.CoverTab[126237]++
												tag = tag_suffix
												tag_suffix = nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:434
			// _ = "end of CoverTab[126237]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:435
			_go_fuzz_dep_.CoverTab[126238]++
												for i := range parser.tag_directives {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:436
				_go_fuzz_dep_.CoverTab[126240]++
													if bytes.Equal(parser.tag_directives[i].handle, tag_handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:437
					_go_fuzz_dep_.CoverTab[126241]++
														tag = append([]byte(nil), parser.tag_directives[i].prefix...)
														tag = append(tag, tag_suffix...)
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:440
					// _ = "end of CoverTab[126241]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:441
					_go_fuzz_dep_.CoverTab[126242]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:441
					// _ = "end of CoverTab[126242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:441
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:441
				// _ = "end of CoverTab[126240]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:442
			// _ = "end of CoverTab[126238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:442
			_go_fuzz_dep_.CoverTab[126239]++
												if len(tag) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:443
				_go_fuzz_dep_.CoverTab[126243]++
													yaml_parser_set_parser_error_context(parser,
					"while parsing a node", start_mark,
					"found undefined tag handle", tag_mark)
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:447
				// _ = "end of CoverTab[126243]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:448
				_go_fuzz_dep_.CoverTab[126244]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:448
				// _ = "end of CoverTab[126244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:448
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:448
			// _ = "end of CoverTab[126239]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:449
		// _ = "end of CoverTab[126236]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:450
		_go_fuzz_dep_.CoverTab[126245]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:450
		// _ = "end of CoverTab[126245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:450
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:450
	// _ = "end of CoverTab[126204]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:450
	_go_fuzz_dep_.CoverTab[126205]++

										implicit := len(tag) == 0
										if indentless_sequence && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:453
		_go_fuzz_dep_.CoverTab[126246]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:453
		return token.typ == yaml_BLOCK_ENTRY_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:453
		// _ = "end of CoverTab[126246]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:453
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:453
		_go_fuzz_dep_.CoverTab[126247]++
											end_mark = token.end_mark
											parser.state = yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE
											*event = yaml_event_t{
			typ:		yaml_SEQUENCE_START_EVENT,
			start_mark:	start_mark,
			end_mark:	end_mark,
			anchor:		anchor,
			tag:		tag,
			implicit:	implicit,
			style:		yaml_style_t(yaml_BLOCK_SEQUENCE_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:465
		// _ = "end of CoverTab[126247]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:466
		_go_fuzz_dep_.CoverTab[126248]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:466
		// _ = "end of CoverTab[126248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:466
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:466
	// _ = "end of CoverTab[126205]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:466
	_go_fuzz_dep_.CoverTab[126206]++
										if token.typ == yaml_SCALAR_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:467
		_go_fuzz_dep_.CoverTab[126249]++
											var plain_implicit, quoted_implicit bool
											end_mark = token.end_mark
											if (len(tag) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			_go_fuzz_dep_.CoverTab[126251]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			return token.style == yaml_PLAIN_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			// _ = "end of CoverTab[126251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			_go_fuzz_dep_.CoverTab[126252]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			return (len(tag) == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
				_go_fuzz_dep_.CoverTab[126253]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
				return tag[0] == '!'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
				// _ = "end of CoverTab[126253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			// _ = "end of CoverTab[126252]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:470
			_go_fuzz_dep_.CoverTab[126254]++
												plain_implicit = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:471
			// _ = "end of CoverTab[126254]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:472
			_go_fuzz_dep_.CoverTab[126255]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:472
			if len(tag) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:472
				_go_fuzz_dep_.CoverTab[126256]++
													quoted_implicit = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:473
				// _ = "end of CoverTab[126256]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
				_go_fuzz_dep_.CoverTab[126257]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
				// _ = "end of CoverTab[126257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
			// _ = "end of CoverTab[126255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
		// _ = "end of CoverTab[126249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:474
		_go_fuzz_dep_.CoverTab[126250]++
											parser.state = parser.states[len(parser.states)-1]
											parser.states = parser.states[:len(parser.states)-1]

											*event = yaml_event_t{
			typ:			yaml_SCALAR_EVENT,
			start_mark:		start_mark,
			end_mark:		end_mark,
			anchor:			anchor,
			tag:			tag,
			value:			token.value,
			implicit:		plain_implicit,
			quoted_implicit:	quoted_implicit,
			style:			yaml_style_t(token.style),
		}
											skip_token(parser)
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:490
		// _ = "end of CoverTab[126250]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:491
		_go_fuzz_dep_.CoverTab[126258]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:491
		// _ = "end of CoverTab[126258]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:491
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:491
	// _ = "end of CoverTab[126206]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:491
	_go_fuzz_dep_.CoverTab[126207]++
										if token.typ == yaml_FLOW_SEQUENCE_START_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:492
		_go_fuzz_dep_.CoverTab[126259]++

											end_mark = token.end_mark
											parser.state = yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE
											*event = yaml_event_t{
			typ:		yaml_SEQUENCE_START_EVENT,
			start_mark:	start_mark,
			end_mark:	end_mark,
			anchor:		anchor,
			tag:		tag,
			implicit:	implicit,
			style:		yaml_style_t(yaml_FLOW_SEQUENCE_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:505
		// _ = "end of CoverTab[126259]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:506
		_go_fuzz_dep_.CoverTab[126260]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:506
		// _ = "end of CoverTab[126260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:506
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:506
	// _ = "end of CoverTab[126207]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:506
	_go_fuzz_dep_.CoverTab[126208]++
										if token.typ == yaml_FLOW_MAPPING_START_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:507
		_go_fuzz_dep_.CoverTab[126261]++
											end_mark = token.end_mark
											parser.state = yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE
											*event = yaml_event_t{
			typ:		yaml_MAPPING_START_EVENT,
			start_mark:	start_mark,
			end_mark:	end_mark,
			anchor:		anchor,
			tag:		tag,
			implicit:	implicit,
			style:		yaml_style_t(yaml_FLOW_MAPPING_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:519
		// _ = "end of CoverTab[126261]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:520
		_go_fuzz_dep_.CoverTab[126262]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:520
		// _ = "end of CoverTab[126262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:520
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:520
	// _ = "end of CoverTab[126208]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:520
	_go_fuzz_dep_.CoverTab[126209]++
										if block && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:521
		_go_fuzz_dep_.CoverTab[126263]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:521
		return token.typ == yaml_BLOCK_SEQUENCE_START_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:521
		// _ = "end of CoverTab[126263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:521
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:521
		_go_fuzz_dep_.CoverTab[126264]++
											end_mark = token.end_mark
											parser.state = yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE
											*event = yaml_event_t{
			typ:		yaml_SEQUENCE_START_EVENT,
			start_mark:	start_mark,
			end_mark:	end_mark,
			anchor:		anchor,
			tag:		tag,
			implicit:	implicit,
			style:		yaml_style_t(yaml_BLOCK_SEQUENCE_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:533
		// _ = "end of CoverTab[126264]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:534
		_go_fuzz_dep_.CoverTab[126265]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:534
		// _ = "end of CoverTab[126265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:534
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:534
	// _ = "end of CoverTab[126209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:534
	_go_fuzz_dep_.CoverTab[126210]++
										if block && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:535
		_go_fuzz_dep_.CoverTab[126266]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:535
		return token.typ == yaml_BLOCK_MAPPING_START_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:535
		// _ = "end of CoverTab[126266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:535
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:535
		_go_fuzz_dep_.CoverTab[126267]++
											end_mark = token.end_mark
											parser.state = yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE
											*event = yaml_event_t{
			typ:		yaml_MAPPING_START_EVENT,
			start_mark:	start_mark,
			end_mark:	end_mark,
			anchor:		anchor,
			tag:		tag,
			implicit:	implicit,
			style:		yaml_style_t(yaml_BLOCK_MAPPING_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:547
		// _ = "end of CoverTab[126267]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:548
		_go_fuzz_dep_.CoverTab[126268]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:548
		// _ = "end of CoverTab[126268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:548
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:548
	// _ = "end of CoverTab[126210]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:548
	_go_fuzz_dep_.CoverTab[126211]++
										if len(anchor) > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:549
		_go_fuzz_dep_.CoverTab[126269]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:549
		return len(tag) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:549
		// _ = "end of CoverTab[126269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:549
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:549
		_go_fuzz_dep_.CoverTab[126270]++
											parser.state = parser.states[len(parser.states)-1]
											parser.states = parser.states[:len(parser.states)-1]

											*event = yaml_event_t{
			typ:			yaml_SCALAR_EVENT,
			start_mark:		start_mark,
			end_mark:		end_mark,
			anchor:			anchor,
			tag:			tag,
			implicit:		implicit,
			quoted_implicit:	false,
			style:			yaml_style_t(yaml_PLAIN_SCALAR_STYLE),
		}
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:563
		// _ = "end of CoverTab[126270]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:564
		_go_fuzz_dep_.CoverTab[126271]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:564
		// _ = "end of CoverTab[126271]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:564
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:564
	// _ = "end of CoverTab[126211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:564
	_go_fuzz_dep_.CoverTab[126212]++

										context := "while parsing a flow node"
										if block {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:567
		_go_fuzz_dep_.CoverTab[126272]++
											context = "while parsing a block node"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:568
		// _ = "end of CoverTab[126272]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:569
		_go_fuzz_dep_.CoverTab[126273]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:569
		// _ = "end of CoverTab[126273]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:569
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:569
	// _ = "end of CoverTab[126212]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:569
	_go_fuzz_dep_.CoverTab[126213]++
										yaml_parser_set_parser_error_context(parser, context, start_mark,
		"did not find expected node content", token.start_mark)
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:572
	// _ = "end of CoverTab[126213]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:575
// block_sequence ::= BLOCK-SEQUENCE-START (BLOCK-ENTRY block_node?)* BLOCK-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:575
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:575
//	********************  *********** *             *********
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:579
func yaml_parser_parse_block_sequence_entry(parser *yaml_parser_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:579
	_go_fuzz_dep_.CoverTab[126274]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:580
		_go_fuzz_dep_.CoverTab[126279]++
											token := peek_token(parser)
											parser.marks = append(parser.marks, token.start_mark)
											skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:583
		// _ = "end of CoverTab[126279]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:584
		_go_fuzz_dep_.CoverTab[126280]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:584
		// _ = "end of CoverTab[126280]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:584
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:584
	// _ = "end of CoverTab[126274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:584
	_go_fuzz_dep_.CoverTab[126275]++

										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:587
		_go_fuzz_dep_.CoverTab[126281]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:588
		// _ = "end of CoverTab[126281]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:589
		_go_fuzz_dep_.CoverTab[126282]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:589
		// _ = "end of CoverTab[126282]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:589
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:589
	// _ = "end of CoverTab[126275]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:589
	_go_fuzz_dep_.CoverTab[126276]++

										if token.typ == yaml_BLOCK_ENTRY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:591
		_go_fuzz_dep_.CoverTab[126283]++
											mark := token.end_mark
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:595
			_go_fuzz_dep_.CoverTab[126285]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:596
			// _ = "end of CoverTab[126285]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:597
			_go_fuzz_dep_.CoverTab[126286]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:597
			// _ = "end of CoverTab[126286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:597
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:597
		// _ = "end of CoverTab[126283]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:597
		_go_fuzz_dep_.CoverTab[126284]++
											if token.typ != yaml_BLOCK_ENTRY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:598
			_go_fuzz_dep_.CoverTab[126287]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:598
			return token.typ != yaml_BLOCK_END_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:598
			// _ = "end of CoverTab[126287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:598
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:598
			_go_fuzz_dep_.CoverTab[126288]++
												parser.states = append(parser.states, yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE)
												return yaml_parser_parse_node(parser, event, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:600
			// _ = "end of CoverTab[126288]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:601
			_go_fuzz_dep_.CoverTab[126289]++
												parser.state = yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE
												return yaml_parser_process_empty_scalar(parser, event, mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:603
			// _ = "end of CoverTab[126289]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:604
		// _ = "end of CoverTab[126284]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:605
		_go_fuzz_dep_.CoverTab[126290]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:605
		// _ = "end of CoverTab[126290]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:605
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:605
	// _ = "end of CoverTab[126276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:605
	_go_fuzz_dep_.CoverTab[126277]++
										if token.typ == yaml_BLOCK_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:606
		_go_fuzz_dep_.CoverTab[126291]++
											parser.state = parser.states[len(parser.states)-1]
											parser.states = parser.states[:len(parser.states)-1]
											parser.marks = parser.marks[:len(parser.marks)-1]

											*event = yaml_event_t{
			typ:		yaml_SEQUENCE_END_EVENT,
			start_mark:	token.start_mark,
			end_mark:	token.end_mark,
		}

											skip_token(parser)
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:618
		// _ = "end of CoverTab[126291]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:619
		_go_fuzz_dep_.CoverTab[126292]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:619
		// _ = "end of CoverTab[126292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:619
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:619
	// _ = "end of CoverTab[126277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:619
	_go_fuzz_dep_.CoverTab[126278]++

										context_mark := parser.marks[len(parser.marks)-1]
										parser.marks = parser.marks[:len(parser.marks)-1]
										return yaml_parser_set_parser_error_context(parser,
		"while parsing a block collection", context_mark,
		"did not find expected '-' indicator", token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:625
	// _ = "end of CoverTab[126278]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:628
// indentless_sequence  ::= (BLOCK-ENTRY block_node?)+
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:628
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:628
//	*********** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:631
func yaml_parser_parse_indentless_sequence_entry(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:631
	_go_fuzz_dep_.CoverTab[126293]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:633
		_go_fuzz_dep_.CoverTab[126296]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:634
		// _ = "end of CoverTab[126296]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:635
		_go_fuzz_dep_.CoverTab[126297]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:635
		// _ = "end of CoverTab[126297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:635
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:635
	// _ = "end of CoverTab[126293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:635
	_go_fuzz_dep_.CoverTab[126294]++

										if token.typ == yaml_BLOCK_ENTRY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:637
		_go_fuzz_dep_.CoverTab[126298]++
											mark := token.end_mark
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:641
			_go_fuzz_dep_.CoverTab[126301]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:642
			// _ = "end of CoverTab[126301]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:643
			_go_fuzz_dep_.CoverTab[126302]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:643
			// _ = "end of CoverTab[126302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:643
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:643
		// _ = "end of CoverTab[126298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:643
		_go_fuzz_dep_.CoverTab[126299]++
											if token.typ != yaml_BLOCK_ENTRY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:644
			_go_fuzz_dep_.CoverTab[126303]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:644
			return token.typ != yaml_KEY_TOKEN
												// _ = "end of CoverTab[126303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:645
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:645
			_go_fuzz_dep_.CoverTab[126304]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:645
			return token.typ != yaml_VALUE_TOKEN
												// _ = "end of CoverTab[126304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:646
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:646
			_go_fuzz_dep_.CoverTab[126305]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:646
			return token.typ != yaml_BLOCK_END_TOKEN
												// _ = "end of CoverTab[126305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:647
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:647
			_go_fuzz_dep_.CoverTab[126306]++
												parser.states = append(parser.states, yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE)
												return yaml_parser_parse_node(parser, event, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:649
			// _ = "end of CoverTab[126306]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:650
			_go_fuzz_dep_.CoverTab[126307]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:650
			// _ = "end of CoverTab[126307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:650
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:650
		// _ = "end of CoverTab[126299]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:650
		_go_fuzz_dep_.CoverTab[126300]++
											parser.state = yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE
											return yaml_parser_process_empty_scalar(parser, event, mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:652
		// _ = "end of CoverTab[126300]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:653
		_go_fuzz_dep_.CoverTab[126308]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:653
		// _ = "end of CoverTab[126308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:653
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:653
	// _ = "end of CoverTab[126294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:653
	_go_fuzz_dep_.CoverTab[126295]++
										parser.state = parser.states[len(parser.states)-1]
										parser.states = parser.states[:len(parser.states)-1]

										*event = yaml_event_t{
		typ:		yaml_SEQUENCE_END_EVENT,
		start_mark:	token.start_mark,
		end_mark:	token.start_mark,
	}
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:662
	// _ = "end of CoverTab[126295]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
// block_mapping        ::= BLOCK-MAPPING_START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	*******************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	((KEY block_node_or_indentless_sequence?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	  *** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	(VALUE block_node_or_indentless_sequence?)?)*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	BLOCK-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:665
//	*********
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:675
func yaml_parser_parse_block_mapping_key(parser *yaml_parser_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:675
	_go_fuzz_dep_.CoverTab[126309]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:676
		_go_fuzz_dep_.CoverTab[126313]++
											token := peek_token(parser)
											parser.marks = append(parser.marks, token.start_mark)
											skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:679
		// _ = "end of CoverTab[126313]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:680
		_go_fuzz_dep_.CoverTab[126314]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:680
		// _ = "end of CoverTab[126314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:680
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:680
	// _ = "end of CoverTab[126309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:680
	_go_fuzz_dep_.CoverTab[126310]++

										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:683
		_go_fuzz_dep_.CoverTab[126315]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:684
		// _ = "end of CoverTab[126315]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:685
		_go_fuzz_dep_.CoverTab[126316]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:685
		// _ = "end of CoverTab[126316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:685
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:685
	// _ = "end of CoverTab[126310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:685
	_go_fuzz_dep_.CoverTab[126311]++

										if token.typ == yaml_KEY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:687
		_go_fuzz_dep_.CoverTab[126317]++
											mark := token.end_mark
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:691
			_go_fuzz_dep_.CoverTab[126319]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:692
			// _ = "end of CoverTab[126319]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:693
			_go_fuzz_dep_.CoverTab[126320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:693
			// _ = "end of CoverTab[126320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:693
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:693
		// _ = "end of CoverTab[126317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:693
		_go_fuzz_dep_.CoverTab[126318]++
											if token.typ != yaml_KEY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:694
			_go_fuzz_dep_.CoverTab[126321]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:694
			return token.typ != yaml_VALUE_TOKEN
												// _ = "end of CoverTab[126321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:695
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:695
			_go_fuzz_dep_.CoverTab[126322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:695
			return token.typ != yaml_BLOCK_END_TOKEN
												// _ = "end of CoverTab[126322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:696
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:696
			_go_fuzz_dep_.CoverTab[126323]++
												parser.states = append(parser.states, yaml_PARSE_BLOCK_MAPPING_VALUE_STATE)
												return yaml_parser_parse_node(parser, event, true, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:698
			// _ = "end of CoverTab[126323]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:699
			_go_fuzz_dep_.CoverTab[126324]++
												parser.state = yaml_PARSE_BLOCK_MAPPING_VALUE_STATE
												return yaml_parser_process_empty_scalar(parser, event, mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:701
			// _ = "end of CoverTab[126324]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:702
		// _ = "end of CoverTab[126318]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:703
		_go_fuzz_dep_.CoverTab[126325]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:703
		if token.typ == yaml_BLOCK_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:703
			_go_fuzz_dep_.CoverTab[126326]++
												parser.state = parser.states[len(parser.states)-1]
												parser.states = parser.states[:len(parser.states)-1]
												parser.marks = parser.marks[:len(parser.marks)-1]
												*event = yaml_event_t{
				typ:		yaml_MAPPING_END_EVENT,
				start_mark:	token.start_mark,
				end_mark:	token.end_mark,
			}
												skip_token(parser)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:713
			// _ = "end of CoverTab[126326]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
			_go_fuzz_dep_.CoverTab[126327]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
			// _ = "end of CoverTab[126327]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
		// _ = "end of CoverTab[126325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
	// _ = "end of CoverTab[126311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:714
	_go_fuzz_dep_.CoverTab[126312]++

										context_mark := parser.marks[len(parser.marks)-1]
										parser.marks = parser.marks[:len(parser.marks)-1]
										return yaml_parser_set_parser_error_context(parser,
		"while parsing a block mapping", context_mark,
		"did not find expected key", token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:720
	// _ = "end of CoverTab[126312]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
// block_mapping        ::= BLOCK-MAPPING_START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//	((KEY block_node_or_indentless_sequence?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//	(VALUE block_node_or_indentless_sequence?)?)*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//	 ***** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:723
//	BLOCK-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:733
func yaml_parser_parse_block_mapping_value(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:733
	_go_fuzz_dep_.CoverTab[126328]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:735
		_go_fuzz_dep_.CoverTab[126331]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:736
		// _ = "end of CoverTab[126331]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:737
		_go_fuzz_dep_.CoverTab[126332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:737
		// _ = "end of CoverTab[126332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:737
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:737
	// _ = "end of CoverTab[126328]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:737
	_go_fuzz_dep_.CoverTab[126329]++
										if token.typ == yaml_VALUE_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:738
		_go_fuzz_dep_.CoverTab[126333]++
											mark := token.end_mark
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:742
			_go_fuzz_dep_.CoverTab[126336]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:743
			// _ = "end of CoverTab[126336]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:744
			_go_fuzz_dep_.CoverTab[126337]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:744
			// _ = "end of CoverTab[126337]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:744
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:744
		// _ = "end of CoverTab[126333]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:744
		_go_fuzz_dep_.CoverTab[126334]++
											if token.typ != yaml_KEY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:745
			_go_fuzz_dep_.CoverTab[126338]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:745
			return token.typ != yaml_VALUE_TOKEN
												// _ = "end of CoverTab[126338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:746
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:746
			_go_fuzz_dep_.CoverTab[126339]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:746
			return token.typ != yaml_BLOCK_END_TOKEN
												// _ = "end of CoverTab[126339]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:747
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:747
			_go_fuzz_dep_.CoverTab[126340]++
												parser.states = append(parser.states, yaml_PARSE_BLOCK_MAPPING_KEY_STATE)
												return yaml_parser_parse_node(parser, event, true, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:749
			// _ = "end of CoverTab[126340]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:750
			_go_fuzz_dep_.CoverTab[126341]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:750
			// _ = "end of CoverTab[126341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:750
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:750
		// _ = "end of CoverTab[126334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:750
		_go_fuzz_dep_.CoverTab[126335]++
											parser.state = yaml_PARSE_BLOCK_MAPPING_KEY_STATE
											return yaml_parser_process_empty_scalar(parser, event, mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:752
		// _ = "end of CoverTab[126335]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:753
		_go_fuzz_dep_.CoverTab[126342]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:753
		// _ = "end of CoverTab[126342]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:753
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:753
	// _ = "end of CoverTab[126329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:753
	_go_fuzz_dep_.CoverTab[126330]++
										parser.state = yaml_PARSE_BLOCK_MAPPING_KEY_STATE
										return yaml_parser_process_empty_scalar(parser, event, token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:755
	// _ = "end of CoverTab[126330]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
// flow_sequence        ::= FLOW-SEQUENCE-START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	*******************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	(flow_sequence_entry FLOW-ENTRY)*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	 *                   **********
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	flow_sequence_entry?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	FLOW-SEQUENCE-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	*****************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
// flow_sequence_entry  ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:758
//	*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:770
func yaml_parser_parse_flow_sequence_entry(parser *yaml_parser_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:770
	_go_fuzz_dep_.CoverTab[126343]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:771
		_go_fuzz_dep_.CoverTab[126347]++
											token := peek_token(parser)
											parser.marks = append(parser.marks, token.start_mark)
											skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:774
		// _ = "end of CoverTab[126347]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:775
		_go_fuzz_dep_.CoverTab[126348]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:775
		// _ = "end of CoverTab[126348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:775
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:775
	// _ = "end of CoverTab[126343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:775
	_go_fuzz_dep_.CoverTab[126344]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:777
		_go_fuzz_dep_.CoverTab[126349]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:778
		// _ = "end of CoverTab[126349]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:779
		_go_fuzz_dep_.CoverTab[126350]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:779
		// _ = "end of CoverTab[126350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:779
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:779
	// _ = "end of CoverTab[126344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:779
	_go_fuzz_dep_.CoverTab[126345]++
										if token.typ != yaml_FLOW_SEQUENCE_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:780
		_go_fuzz_dep_.CoverTab[126351]++
											if !first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:781
			_go_fuzz_dep_.CoverTab[126353]++
												if token.typ == yaml_FLOW_ENTRY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:782
				_go_fuzz_dep_.CoverTab[126354]++
													skip_token(parser)
													token = peek_token(parser)
													if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:785
					_go_fuzz_dep_.CoverTab[126355]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:786
					// _ = "end of CoverTab[126355]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:787
					_go_fuzz_dep_.CoverTab[126356]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:787
					// _ = "end of CoverTab[126356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:787
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:787
				// _ = "end of CoverTab[126354]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:788
				_go_fuzz_dep_.CoverTab[126357]++
													context_mark := parser.marks[len(parser.marks)-1]
													parser.marks = parser.marks[:len(parser.marks)-1]
													return yaml_parser_set_parser_error_context(parser,
					"while parsing a flow sequence", context_mark,
					"did not find expected ',' or ']'", token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:793
				// _ = "end of CoverTab[126357]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:794
			// _ = "end of CoverTab[126353]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:795
			_go_fuzz_dep_.CoverTab[126358]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:795
			// _ = "end of CoverTab[126358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:795
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:795
		// _ = "end of CoverTab[126351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:795
		_go_fuzz_dep_.CoverTab[126352]++

											if token.typ == yaml_KEY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:797
			_go_fuzz_dep_.CoverTab[126359]++
												parser.state = yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE
												*event = yaml_event_t{
				typ:		yaml_MAPPING_START_EVENT,
				start_mark:	token.start_mark,
				end_mark:	token.end_mark,
				implicit:	true,
				style:		yaml_style_t(yaml_FLOW_MAPPING_STYLE),
			}
												skip_token(parser)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:807
			// _ = "end of CoverTab[126359]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:808
			_go_fuzz_dep_.CoverTab[126360]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:808
			if token.typ != yaml_FLOW_SEQUENCE_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:808
				_go_fuzz_dep_.CoverTab[126361]++
													parser.states = append(parser.states, yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE)
													return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:810
				// _ = "end of CoverTab[126361]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
				_go_fuzz_dep_.CoverTab[126362]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
				// _ = "end of CoverTab[126362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
			// _ = "end of CoverTab[126360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:811
		// _ = "end of CoverTab[126352]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:812
		_go_fuzz_dep_.CoverTab[126363]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:812
		// _ = "end of CoverTab[126363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:812
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:812
	// _ = "end of CoverTab[126345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:812
	_go_fuzz_dep_.CoverTab[126346]++

										parser.state = parser.states[len(parser.states)-1]
										parser.states = parser.states[:len(parser.states)-1]
										parser.marks = parser.marks[:len(parser.marks)-1]

										*event = yaml_event_t{
		typ:		yaml_SEQUENCE_END_EVENT,
		start_mark:	token.start_mark,
		end_mark:	token.end_mark,
	}

										skip_token(parser)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:825
	// _ = "end of CoverTab[126346]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:828
// flow_sequence_entry  ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:828
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:828
//	*** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:833
func yaml_parser_parse_flow_sequence_entry_mapping_key(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:833
	_go_fuzz_dep_.CoverTab[126364]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:835
		_go_fuzz_dep_.CoverTab[126367]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:836
		// _ = "end of CoverTab[126367]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:837
		_go_fuzz_dep_.CoverTab[126368]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:837
		// _ = "end of CoverTab[126368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:837
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:837
	// _ = "end of CoverTab[126364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:837
	_go_fuzz_dep_.CoverTab[126365]++
										if token.typ != yaml_VALUE_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:838
		_go_fuzz_dep_.CoverTab[126369]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:838
		return token.typ != yaml_FLOW_ENTRY_TOKEN
											// _ = "end of CoverTab[126369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:839
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:839
		_go_fuzz_dep_.CoverTab[126370]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:839
		return token.typ != yaml_FLOW_SEQUENCE_END_TOKEN
											// _ = "end of CoverTab[126370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:840
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:840
		_go_fuzz_dep_.CoverTab[126371]++
											parser.states = append(parser.states, yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE)
											return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:842
		// _ = "end of CoverTab[126371]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:843
		_go_fuzz_dep_.CoverTab[126372]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:843
		// _ = "end of CoverTab[126372]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:843
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:843
	// _ = "end of CoverTab[126365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:843
	_go_fuzz_dep_.CoverTab[126366]++
										mark := token.end_mark
										skip_token(parser)
										parser.state = yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE
										return yaml_parser_process_empty_scalar(parser, event, mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:847
	// _ = "end of CoverTab[126366]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:850
// flow_sequence_entry  ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:850
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:850
//	***** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:854
func yaml_parser_parse_flow_sequence_entry_mapping_value(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:854
	_go_fuzz_dep_.CoverTab[126373]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:856
		_go_fuzz_dep_.CoverTab[126376]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:857
		// _ = "end of CoverTab[126376]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:858
		_go_fuzz_dep_.CoverTab[126377]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:858
		// _ = "end of CoverTab[126377]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:858
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:858
	// _ = "end of CoverTab[126373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:858
	_go_fuzz_dep_.CoverTab[126374]++
										if token.typ == yaml_VALUE_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:859
		_go_fuzz_dep_.CoverTab[126378]++
											skip_token(parser)
											token := peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:862
			_go_fuzz_dep_.CoverTab[126380]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:863
			// _ = "end of CoverTab[126380]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:864
			_go_fuzz_dep_.CoverTab[126381]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:864
			// _ = "end of CoverTab[126381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:864
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:864
		// _ = "end of CoverTab[126378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:864
		_go_fuzz_dep_.CoverTab[126379]++
											if token.typ != yaml_FLOW_ENTRY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:865
			_go_fuzz_dep_.CoverTab[126382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:865
			return token.typ != yaml_FLOW_SEQUENCE_END_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:865
			// _ = "end of CoverTab[126382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:865
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:865
			_go_fuzz_dep_.CoverTab[126383]++
												parser.states = append(parser.states, yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE)
												return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:867
			// _ = "end of CoverTab[126383]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:868
			_go_fuzz_dep_.CoverTab[126384]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:868
			// _ = "end of CoverTab[126384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:868
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:868
		// _ = "end of CoverTab[126379]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:869
		_go_fuzz_dep_.CoverTab[126385]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:869
		// _ = "end of CoverTab[126385]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:869
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:869
	// _ = "end of CoverTab[126374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:869
	_go_fuzz_dep_.CoverTab[126375]++
										parser.state = yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE
										return yaml_parser_process_empty_scalar(parser, event, token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:871
	// _ = "end of CoverTab[126375]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:874
// flow_sequence_entry  ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:874
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:874
//	*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:878
func yaml_parser_parse_flow_sequence_entry_mapping_end(parser *yaml_parser_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:878
	_go_fuzz_dep_.CoverTab[126386]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:880
		_go_fuzz_dep_.CoverTab[126388]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:881
		// _ = "end of CoverTab[126388]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:882
		_go_fuzz_dep_.CoverTab[126389]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:882
		// _ = "end of CoverTab[126389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:882
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:882
	// _ = "end of CoverTab[126386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:882
	_go_fuzz_dep_.CoverTab[126387]++
										parser.state = yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE
										*event = yaml_event_t{
		typ:		yaml_MAPPING_END_EVENT,
		start_mark:	token.start_mark,
		end_mark:	token.start_mark,
	}
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:889
	// _ = "end of CoverTab[126387]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
// flow_mapping         ::= FLOW-MAPPING-START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	******************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	(flow_mapping_entry FLOW-ENTRY)*
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	 *                  **********
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	flow_mapping_entry?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	******************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	FLOW-MAPPING-END
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//	****************
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
// flow_mapping_entry   ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:892
//   - *** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:904
func yaml_parser_parse_flow_mapping_key(parser *yaml_parser_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:904
	_go_fuzz_dep_.CoverTab[126390]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:905
		_go_fuzz_dep_.CoverTab[126394]++
											token := peek_token(parser)
											parser.marks = append(parser.marks, token.start_mark)
											skip_token(parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:908
		// _ = "end of CoverTab[126394]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:909
		_go_fuzz_dep_.CoverTab[126395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:909
		// _ = "end of CoverTab[126395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:909
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:909
	// _ = "end of CoverTab[126390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:909
	_go_fuzz_dep_.CoverTab[126391]++

										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:912
		_go_fuzz_dep_.CoverTab[126396]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:913
		// _ = "end of CoverTab[126396]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:914
		_go_fuzz_dep_.CoverTab[126397]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:914
		// _ = "end of CoverTab[126397]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:914
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:914
	// _ = "end of CoverTab[126391]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:914
	_go_fuzz_dep_.CoverTab[126392]++

										if token.typ != yaml_FLOW_MAPPING_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:916
		_go_fuzz_dep_.CoverTab[126398]++
											if !first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:917
			_go_fuzz_dep_.CoverTab[126400]++
												if token.typ == yaml_FLOW_ENTRY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:918
				_go_fuzz_dep_.CoverTab[126401]++
													skip_token(parser)
													token = peek_token(parser)
													if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:921
					_go_fuzz_dep_.CoverTab[126402]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:922
					// _ = "end of CoverTab[126402]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:923
					_go_fuzz_dep_.CoverTab[126403]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:923
					// _ = "end of CoverTab[126403]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:923
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:923
				// _ = "end of CoverTab[126401]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:924
				_go_fuzz_dep_.CoverTab[126404]++
													context_mark := parser.marks[len(parser.marks)-1]
													parser.marks = parser.marks[:len(parser.marks)-1]
													return yaml_parser_set_parser_error_context(parser,
					"while parsing a flow mapping", context_mark,
					"did not find expected ',' or '}'", token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:929
				// _ = "end of CoverTab[126404]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:930
			// _ = "end of CoverTab[126400]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:931
			_go_fuzz_dep_.CoverTab[126405]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:931
			// _ = "end of CoverTab[126405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:931
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:931
		// _ = "end of CoverTab[126398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:931
		_go_fuzz_dep_.CoverTab[126399]++

											if token.typ == yaml_KEY_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:933
			_go_fuzz_dep_.CoverTab[126406]++
												skip_token(parser)
												token = peek_token(parser)
												if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:936
				_go_fuzz_dep_.CoverTab[126408]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:937
				// _ = "end of CoverTab[126408]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:938
				_go_fuzz_dep_.CoverTab[126409]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:938
				// _ = "end of CoverTab[126409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:938
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:938
			// _ = "end of CoverTab[126406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:938
			_go_fuzz_dep_.CoverTab[126407]++
												if token.typ != yaml_VALUE_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:939
				_go_fuzz_dep_.CoverTab[126410]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:939
				return token.typ != yaml_FLOW_ENTRY_TOKEN
													// _ = "end of CoverTab[126410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:940
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:940
				_go_fuzz_dep_.CoverTab[126411]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:940
				return token.typ != yaml_FLOW_MAPPING_END_TOKEN
													// _ = "end of CoverTab[126411]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:941
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:941
				_go_fuzz_dep_.CoverTab[126412]++
													parser.states = append(parser.states, yaml_PARSE_FLOW_MAPPING_VALUE_STATE)
													return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:943
				// _ = "end of CoverTab[126412]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:944
				_go_fuzz_dep_.CoverTab[126413]++
													parser.state = yaml_PARSE_FLOW_MAPPING_VALUE_STATE
													return yaml_parser_process_empty_scalar(parser, event, token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:946
				// _ = "end of CoverTab[126413]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:947
			// _ = "end of CoverTab[126407]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:948
			_go_fuzz_dep_.CoverTab[126414]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:948
			if token.typ != yaml_FLOW_MAPPING_END_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:948
				_go_fuzz_dep_.CoverTab[126415]++
													parser.states = append(parser.states, yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE)
													return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:950
				// _ = "end of CoverTab[126415]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
				_go_fuzz_dep_.CoverTab[126416]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
				// _ = "end of CoverTab[126416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
			// _ = "end of CoverTab[126414]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:951
		// _ = "end of CoverTab[126399]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:952
		_go_fuzz_dep_.CoverTab[126417]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:952
		// _ = "end of CoverTab[126417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:952
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:952
	// _ = "end of CoverTab[126392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:952
	_go_fuzz_dep_.CoverTab[126393]++

										parser.state = parser.states[len(parser.states)-1]
										parser.states = parser.states[:len(parser.states)-1]
										parser.marks = parser.marks[:len(parser.marks)-1]
										*event = yaml_event_t{
		typ:		yaml_MAPPING_END_EVENT,
		start_mark:	token.start_mark,
		end_mark:	token.end_mark,
	}
										skip_token(parser)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:963
	// _ = "end of CoverTab[126393]"
}

// Parse the productions:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:966
// flow_mapping_entry   ::= flow_node | KEY flow_node? (VALUE flow_node?)?
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:966
//   - ***** *
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:970
func yaml_parser_parse_flow_mapping_value(parser *yaml_parser_t, event *yaml_event_t, empty bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:970
	_go_fuzz_dep_.CoverTab[126418]++
										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:972
		_go_fuzz_dep_.CoverTab[126422]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:973
		// _ = "end of CoverTab[126422]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:974
		_go_fuzz_dep_.CoverTab[126423]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:974
		// _ = "end of CoverTab[126423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:974
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:974
	// _ = "end of CoverTab[126418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:974
	_go_fuzz_dep_.CoverTab[126419]++
										if empty {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:975
		_go_fuzz_dep_.CoverTab[126424]++
											parser.state = yaml_PARSE_FLOW_MAPPING_KEY_STATE
											return yaml_parser_process_empty_scalar(parser, event, token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:977
		// _ = "end of CoverTab[126424]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:978
		_go_fuzz_dep_.CoverTab[126425]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:978
		// _ = "end of CoverTab[126425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:978
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:978
	// _ = "end of CoverTab[126419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:978
	_go_fuzz_dep_.CoverTab[126420]++
										if token.typ == yaml_VALUE_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:979
		_go_fuzz_dep_.CoverTab[126426]++
											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:982
			_go_fuzz_dep_.CoverTab[126428]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:983
			// _ = "end of CoverTab[126428]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:984
			_go_fuzz_dep_.CoverTab[126429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:984
			// _ = "end of CoverTab[126429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:984
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:984
		// _ = "end of CoverTab[126426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:984
		_go_fuzz_dep_.CoverTab[126427]++
											if token.typ != yaml_FLOW_ENTRY_TOKEN && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:985
			_go_fuzz_dep_.CoverTab[126430]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:985
			return token.typ != yaml_FLOW_MAPPING_END_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:985
			// _ = "end of CoverTab[126430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:985
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:985
			_go_fuzz_dep_.CoverTab[126431]++
												parser.states = append(parser.states, yaml_PARSE_FLOW_MAPPING_KEY_STATE)
												return yaml_parser_parse_node(parser, event, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:987
			// _ = "end of CoverTab[126431]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:988
			_go_fuzz_dep_.CoverTab[126432]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:988
			// _ = "end of CoverTab[126432]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:988
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:988
		// _ = "end of CoverTab[126427]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:989
		_go_fuzz_dep_.CoverTab[126433]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:989
		// _ = "end of CoverTab[126433]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:989
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:989
	// _ = "end of CoverTab[126420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:989
	_go_fuzz_dep_.CoverTab[126421]++
										parser.state = yaml_PARSE_FLOW_MAPPING_KEY_STATE
										return yaml_parser_process_empty_scalar(parser, event, token.start_mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:991
	// _ = "end of CoverTab[126421]"
}

// Generate an empty scalar event.
func yaml_parser_process_empty_scalar(parser *yaml_parser_t, event *yaml_event_t, mark yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:995
	_go_fuzz_dep_.CoverTab[126434]++
										*event = yaml_event_t{
		typ:		yaml_SCALAR_EVENT,
		start_mark:	mark,
		end_mark:	mark,
		value:		nil,
		implicit:	true,
		style:		yaml_style_t(yaml_PLAIN_SCALAR_STYLE),
	}
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1004
	// _ = "end of CoverTab[126434]"
}

var default_tag_directives = []yaml_tag_directive_t{
	{[]byte("!"), []byte("!")},
	{[]byte("!!"), []byte("tag:yaml.org,2002:")},
}

// Parse directives.
func yaml_parser_process_directives(parser *yaml_parser_t,
	version_directive_ref **yaml_version_directive_t,
	tag_directives_ref *[]yaml_tag_directive_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1015
	_go_fuzz_dep_.CoverTab[126435]++

										var version_directive *yaml_version_directive_t
										var tag_directives []yaml_tag_directive_t

										token := peek_token(parser)
										if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1021
		_go_fuzz_dep_.CoverTab[126441]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1022
		// _ = "end of CoverTab[126441]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1023
		_go_fuzz_dep_.CoverTab[126442]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1023
		// _ = "end of CoverTab[126442]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1023
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1023
	// _ = "end of CoverTab[126435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1023
	_go_fuzz_dep_.CoverTab[126436]++

										for token.typ == yaml_VERSION_DIRECTIVE_TOKEN || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1025
		_go_fuzz_dep_.CoverTab[126443]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1025
		return token.typ == yaml_TAG_DIRECTIVE_TOKEN
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1025
		// _ = "end of CoverTab[126443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1025
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1025
		_go_fuzz_dep_.CoverTab[126444]++
											if token.typ == yaml_VERSION_DIRECTIVE_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1026
			_go_fuzz_dep_.CoverTab[126446]++
												if version_directive != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1027
				_go_fuzz_dep_.CoverTab[126449]++
													yaml_parser_set_parser_error(parser,
					"found duplicate %YAML directive", token.start_mark)
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1030
				// _ = "end of CoverTab[126449]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1031
				_go_fuzz_dep_.CoverTab[126450]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1031
				// _ = "end of CoverTab[126450]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1031
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1031
			// _ = "end of CoverTab[126446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1031
			_go_fuzz_dep_.CoverTab[126447]++
												if token.major != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1032
				_go_fuzz_dep_.CoverTab[126451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1032
				return token.minor != 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1032
				// _ = "end of CoverTab[126451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1032
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1032
				_go_fuzz_dep_.CoverTab[126452]++
													yaml_parser_set_parser_error(parser,
					"found incompatible YAML document", token.start_mark)
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1035
				// _ = "end of CoverTab[126452]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1036
				_go_fuzz_dep_.CoverTab[126453]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1036
				// _ = "end of CoverTab[126453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1036
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1036
			// _ = "end of CoverTab[126447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1036
			_go_fuzz_dep_.CoverTab[126448]++
												version_directive = &yaml_version_directive_t{
				major:	token.major,
				minor:	token.minor,
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1040
			// _ = "end of CoverTab[126448]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1041
			_go_fuzz_dep_.CoverTab[126454]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1041
			if token.typ == yaml_TAG_DIRECTIVE_TOKEN {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1041
				_go_fuzz_dep_.CoverTab[126455]++
													value := yaml_tag_directive_t{
					handle:	token.value,
					prefix:	token.prefix,
				}
				if !yaml_parser_append_tag_directive(parser, value, false, token.start_mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1046
					_go_fuzz_dep_.CoverTab[126457]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1047
					// _ = "end of CoverTab[126457]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1048
					_go_fuzz_dep_.CoverTab[126458]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1048
					// _ = "end of CoverTab[126458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1048
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1048
				// _ = "end of CoverTab[126455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1048
				_go_fuzz_dep_.CoverTab[126456]++
													tag_directives = append(tag_directives, value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1049
				// _ = "end of CoverTab[126456]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
				_go_fuzz_dep_.CoverTab[126459]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
				// _ = "end of CoverTab[126459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
			// _ = "end of CoverTab[126454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
		// _ = "end of CoverTab[126444]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1050
		_go_fuzz_dep_.CoverTab[126445]++

											skip_token(parser)
											token = peek_token(parser)
											if token == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1054
			_go_fuzz_dep_.CoverTab[126460]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1055
			// _ = "end of CoverTab[126460]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1056
			_go_fuzz_dep_.CoverTab[126461]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1056
			// _ = "end of CoverTab[126461]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1056
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1056
		// _ = "end of CoverTab[126445]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1057
	// _ = "end of CoverTab[126436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1057
	_go_fuzz_dep_.CoverTab[126437]++

										for i := range default_tag_directives {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1059
		_go_fuzz_dep_.CoverTab[126462]++
											if !yaml_parser_append_tag_directive(parser, default_tag_directives[i], true, token.start_mark) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1060
			_go_fuzz_dep_.CoverTab[126463]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1061
			// _ = "end of CoverTab[126463]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1062
			_go_fuzz_dep_.CoverTab[126464]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1062
			// _ = "end of CoverTab[126464]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1062
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1062
		// _ = "end of CoverTab[126462]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1063
	// _ = "end of CoverTab[126437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1063
	_go_fuzz_dep_.CoverTab[126438]++

										if version_directive_ref != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1065
		_go_fuzz_dep_.CoverTab[126465]++
											*version_directive_ref = version_directive
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1066
		// _ = "end of CoverTab[126465]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1067
		_go_fuzz_dep_.CoverTab[126466]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1067
		// _ = "end of CoverTab[126466]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1067
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1067
	// _ = "end of CoverTab[126438]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1067
	_go_fuzz_dep_.CoverTab[126439]++
										if tag_directives_ref != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1068
		_go_fuzz_dep_.CoverTab[126467]++
											*tag_directives_ref = tag_directives
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1069
		// _ = "end of CoverTab[126467]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1070
		_go_fuzz_dep_.CoverTab[126468]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1070
		// _ = "end of CoverTab[126468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1070
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1070
	// _ = "end of CoverTab[126439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1070
	_go_fuzz_dep_.CoverTab[126440]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1071
	// _ = "end of CoverTab[126440]"
}

// Append a tag directive to the directives stack.
func yaml_parser_append_tag_directive(parser *yaml_parser_t, value yaml_tag_directive_t, allow_duplicates bool, mark yaml_mark_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1075
	_go_fuzz_dep_.CoverTab[126469]++
										for i := range parser.tag_directives {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1076
		_go_fuzz_dep_.CoverTab[126471]++
											if bytes.Equal(value.handle, parser.tag_directives[i].handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1077
			_go_fuzz_dep_.CoverTab[126472]++
												if allow_duplicates {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1078
				_go_fuzz_dep_.CoverTab[126474]++
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1079
				// _ = "end of CoverTab[126474]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1080
				_go_fuzz_dep_.CoverTab[126475]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1080
				// _ = "end of CoverTab[126475]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1080
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1080
			// _ = "end of CoverTab[126472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1080
			_go_fuzz_dep_.CoverTab[126473]++
												return yaml_parser_set_parser_error(parser, "found duplicate %TAG directive", mark)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1081
			// _ = "end of CoverTab[126473]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1082
			_go_fuzz_dep_.CoverTab[126476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1082
			// _ = "end of CoverTab[126476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1082
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1082
		// _ = "end of CoverTab[126471]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1083
	// _ = "end of CoverTab[126469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1083
	_go_fuzz_dep_.CoverTab[126470]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1087
	value_copy := yaml_tag_directive_t{
		handle:	make([]byte, len(value.handle)),
		prefix:	make([]byte, len(value.prefix)),
	}
										copy(value_copy.handle, value.handle)
										copy(value_copy.prefix, value.prefix)
										parser.tag_directives = append(parser.tag_directives, value_copy)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1094
	// _ = "end of CoverTab[126470]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1095
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/parserc.go:1095
var _ = _go_fuzz_dep_.CoverTab
