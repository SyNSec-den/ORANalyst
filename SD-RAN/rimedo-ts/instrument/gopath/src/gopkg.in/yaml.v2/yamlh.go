//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:1
)

import (
	"fmt"
	"io"
)

// The version directive data.
type yaml_version_directive_t struct {
	major	int8	// The major version number.
	minor	int8	// The minor version number.
}

// The tag directive data.
type yaml_tag_directive_t struct {
	handle	[]byte	// The tag handle.
	prefix	[]byte	// The tag prefix.
}

type yaml_encoding_t int

// The stream encoding.
const (
	// Let the parser choose the encoding.
	yaml_ANY_ENCODING	yaml_encoding_t	= iota

	yaml_UTF8_ENCODING	// The default UTF-8 encoding.
	yaml_UTF16LE_ENCODING	// The UTF-16-LE encoding with BOM.
	yaml_UTF16BE_ENCODING	// The UTF-16-BE encoding with BOM.
)

type yaml_break_t int

// Line break types.
const (
	// Let the parser choose the break type.
	yaml_ANY_BREAK	yaml_break_t	= iota

	yaml_CR_BREAK	// Use CR for line breaks (Mac style).
	yaml_LN_BREAK	// Use LN for line breaks (Unix style).
	yaml_CRLN_BREAK	// Use CR LN for line breaks (DOS style).
)

type yaml_error_type_t int

// Many bad things could happen with the parser and emitter.
const (
	// No error is produced.
	yaml_NO_ERROR	yaml_error_type_t	= iota

	yaml_MEMORY_ERROR	// Cannot allocate or reallocate a block of memory.
	yaml_READER_ERROR	// Cannot read or decode the input stream.
	yaml_SCANNER_ERROR	// Cannot scan the input stream.
	yaml_PARSER_ERROR	// Cannot parse the input stream.
	yaml_COMPOSER_ERROR	// Cannot compose a YAML document.
	yaml_WRITER_ERROR	// Cannot write to the output stream.
	yaml_EMITTER_ERROR	// Cannot emit a YAML stream.
)

// The pointer position.
type yaml_mark_t struct {
	index	int	// The position index.
	line	int	// The position line.
	column	int	// The position column.
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:69
type yaml_style_t int8

type yaml_scalar_style_t yaml_style_t

// Scalar styles.
const (
	// Let the emitter choose the style.
	yaml_ANY_SCALAR_STYLE	yaml_scalar_style_t	= iota

	yaml_PLAIN_SCALAR_STYLE		// The plain scalar style.
	yaml_SINGLE_QUOTED_SCALAR_STYLE	// The single-quoted scalar style.
	yaml_DOUBLE_QUOTED_SCALAR_STYLE	// The double-quoted scalar style.
	yaml_LITERAL_SCALAR_STYLE	// The literal scalar style.
	yaml_FOLDED_SCALAR_STYLE	// The folded scalar style.
)

type yaml_sequence_style_t yaml_style_t

// Sequence styles.
const (
	// Let the emitter choose the style.
	yaml_ANY_SEQUENCE_STYLE	yaml_sequence_style_t	= iota

	yaml_BLOCK_SEQUENCE_STYLE	// The block sequence style.
	yaml_FLOW_SEQUENCE_STYLE	// The flow sequence style.
)

type yaml_mapping_style_t yaml_style_t

// Mapping styles.
const (
	// Let the emitter choose the style.
	yaml_ANY_MAPPING_STYLE	yaml_mapping_style_t	= iota

	yaml_BLOCK_MAPPING_STYLE	// The block mapping style.
	yaml_FLOW_MAPPING_STYLE		// The flow mapping style.
)

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:109
type yaml_token_type_t int

// Token types.
const (
	// An empty token.
	yaml_NO_TOKEN	yaml_token_type_t	= iota

	yaml_STREAM_START_TOKEN	// A STREAM-START token.
	yaml_STREAM_END_TOKEN	// A STREAM-END token.

	yaml_VERSION_DIRECTIVE_TOKEN	// A VERSION-DIRECTIVE token.
	yaml_TAG_DIRECTIVE_TOKEN	// A TAG-DIRECTIVE token.
	yaml_DOCUMENT_START_TOKEN	// A DOCUMENT-START token.
	yaml_DOCUMENT_END_TOKEN		// A DOCUMENT-END token.

	yaml_BLOCK_SEQUENCE_START_TOKEN	// A BLOCK-SEQUENCE-START token.
	yaml_BLOCK_MAPPING_START_TOKEN	// A BLOCK-SEQUENCE-END token.
	yaml_BLOCK_END_TOKEN		// A BLOCK-END token.

	yaml_FLOW_SEQUENCE_START_TOKEN	// A FLOW-SEQUENCE-START token.
	yaml_FLOW_SEQUENCE_END_TOKEN	// A FLOW-SEQUENCE-END token.
	yaml_FLOW_MAPPING_START_TOKEN	// A FLOW-MAPPING-START token.
	yaml_FLOW_MAPPING_END_TOKEN	// A FLOW-MAPPING-END token.

	yaml_BLOCK_ENTRY_TOKEN	// A BLOCK-ENTRY token.
	yaml_FLOW_ENTRY_TOKEN	// A FLOW-ENTRY token.
	yaml_KEY_TOKEN		// A KEY token.
	yaml_VALUE_TOKEN	// A VALUE token.

	yaml_ALIAS_TOKEN	// An ALIAS token.
	yaml_ANCHOR_TOKEN	// An ANCHOR token.
	yaml_TAG_TOKEN		// A TAG token.
	yaml_SCALAR_TOKEN	// A SCALAR token.
)

func (tt yaml_token_type_t) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:144
	_go_fuzz_dep_.CoverTab[127972]++
									switch tt {
	case yaml_NO_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:146
		_go_fuzz_dep_.CoverTab[127974]++
										return "yaml_NO_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:147
		// _ = "end of CoverTab[127974]"
	case yaml_STREAM_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:148
		_go_fuzz_dep_.CoverTab[127975]++
										return "yaml_STREAM_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:149
		// _ = "end of CoverTab[127975]"
	case yaml_STREAM_END_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:150
		_go_fuzz_dep_.CoverTab[127976]++
										return "yaml_STREAM_END_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:151
		// _ = "end of CoverTab[127976]"
	case yaml_VERSION_DIRECTIVE_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:152
		_go_fuzz_dep_.CoverTab[127977]++
										return "yaml_VERSION_DIRECTIVE_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:153
		// _ = "end of CoverTab[127977]"
	case yaml_TAG_DIRECTIVE_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:154
		_go_fuzz_dep_.CoverTab[127978]++
										return "yaml_TAG_DIRECTIVE_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:155
		// _ = "end of CoverTab[127978]"
	case yaml_DOCUMENT_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:156
		_go_fuzz_dep_.CoverTab[127979]++
										return "yaml_DOCUMENT_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:157
		// _ = "end of CoverTab[127979]"
	case yaml_DOCUMENT_END_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:158
		_go_fuzz_dep_.CoverTab[127980]++
										return "yaml_DOCUMENT_END_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:159
		// _ = "end of CoverTab[127980]"
	case yaml_BLOCK_SEQUENCE_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:160
		_go_fuzz_dep_.CoverTab[127981]++
										return "yaml_BLOCK_SEQUENCE_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:161
		// _ = "end of CoverTab[127981]"
	case yaml_BLOCK_MAPPING_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:162
		_go_fuzz_dep_.CoverTab[127982]++
										return "yaml_BLOCK_MAPPING_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:163
		// _ = "end of CoverTab[127982]"
	case yaml_BLOCK_END_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:164
		_go_fuzz_dep_.CoverTab[127983]++
										return "yaml_BLOCK_END_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:165
		// _ = "end of CoverTab[127983]"
	case yaml_FLOW_SEQUENCE_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:166
		_go_fuzz_dep_.CoverTab[127984]++
										return "yaml_FLOW_SEQUENCE_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:167
		// _ = "end of CoverTab[127984]"
	case yaml_FLOW_SEQUENCE_END_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:168
		_go_fuzz_dep_.CoverTab[127985]++
										return "yaml_FLOW_SEQUENCE_END_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:169
		// _ = "end of CoverTab[127985]"
	case yaml_FLOW_MAPPING_START_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:170
		_go_fuzz_dep_.CoverTab[127986]++
										return "yaml_FLOW_MAPPING_START_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:171
		// _ = "end of CoverTab[127986]"
	case yaml_FLOW_MAPPING_END_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:172
		_go_fuzz_dep_.CoverTab[127987]++
										return "yaml_FLOW_MAPPING_END_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:173
		// _ = "end of CoverTab[127987]"
	case yaml_BLOCK_ENTRY_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:174
		_go_fuzz_dep_.CoverTab[127988]++
										return "yaml_BLOCK_ENTRY_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:175
		// _ = "end of CoverTab[127988]"
	case yaml_FLOW_ENTRY_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:176
		_go_fuzz_dep_.CoverTab[127989]++
										return "yaml_FLOW_ENTRY_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:177
		// _ = "end of CoverTab[127989]"
	case yaml_KEY_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:178
		_go_fuzz_dep_.CoverTab[127990]++
										return "yaml_KEY_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:179
		// _ = "end of CoverTab[127990]"
	case yaml_VALUE_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:180
		_go_fuzz_dep_.CoverTab[127991]++
										return "yaml_VALUE_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:181
		// _ = "end of CoverTab[127991]"
	case yaml_ALIAS_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:182
		_go_fuzz_dep_.CoverTab[127992]++
										return "yaml_ALIAS_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:183
		// _ = "end of CoverTab[127992]"
	case yaml_ANCHOR_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:184
		_go_fuzz_dep_.CoverTab[127993]++
										return "yaml_ANCHOR_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:185
		// _ = "end of CoverTab[127993]"
	case yaml_TAG_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:186
		_go_fuzz_dep_.CoverTab[127994]++
										return "yaml_TAG_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:187
		// _ = "end of CoverTab[127994]"
	case yaml_SCALAR_TOKEN:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:188
		_go_fuzz_dep_.CoverTab[127995]++
										return "yaml_SCALAR_TOKEN"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:189
		// _ = "end of CoverTab[127995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:189
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:189
		_go_fuzz_dep_.CoverTab[127996]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:189
		// _ = "end of CoverTab[127996]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:190
	// _ = "end of CoverTab[127972]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:190
	_go_fuzz_dep_.CoverTab[127973]++
									return "<unknown token>"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:191
	// _ = "end of CoverTab[127973]"
}

// The token structure.
type yaml_token_t struct {
	// The token type.
	typ	yaml_token_type_t

	// The start/end of the token.
	start_mark, end_mark	yaml_mark_t

	// The stream encoding (for yaml_STREAM_START_TOKEN).
	encoding	yaml_encoding_t

	// The alias/anchor/scalar value or tag/tag directive handle
	// (for yaml_ALIAS_TOKEN, yaml_ANCHOR_TOKEN, yaml_SCALAR_TOKEN, yaml_TAG_TOKEN, yaml_TAG_DIRECTIVE_TOKEN).
	value	[]byte

	// The tag suffix (for yaml_TAG_TOKEN).
	suffix	[]byte

	// The tag directive prefix (for yaml_TAG_DIRECTIVE_TOKEN).
	prefix	[]byte

	// The scalar style (for yaml_SCALAR_TOKEN).
	style	yaml_scalar_style_t

	// The version directive major/minor (for yaml_VERSION_DIRECTIVE_TOKEN).
	major, minor	int8
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:224
type yaml_event_type_t int8

// Event types.
const (
	// An empty event.
	yaml_NO_EVENT	yaml_event_type_t	= iota

	yaml_STREAM_START_EVENT		// A STREAM-START event.
	yaml_STREAM_END_EVENT		// A STREAM-END event.
	yaml_DOCUMENT_START_EVENT	// A DOCUMENT-START event.
	yaml_DOCUMENT_END_EVENT		// A DOCUMENT-END event.
	yaml_ALIAS_EVENT		// An ALIAS event.
	yaml_SCALAR_EVENT		// A SCALAR event.
	yaml_SEQUENCE_START_EVENT	// A SEQUENCE-START event.
	yaml_SEQUENCE_END_EVENT		// A SEQUENCE-END event.
	yaml_MAPPING_START_EVENT	// A MAPPING-START event.
	yaml_MAPPING_END_EVENT		// A MAPPING-END event.
)

var eventStrings = []string{
	yaml_NO_EVENT:			"none",
	yaml_STREAM_START_EVENT:	"stream start",
	yaml_STREAM_END_EVENT:		"stream end",
	yaml_DOCUMENT_START_EVENT:	"document start",
	yaml_DOCUMENT_END_EVENT:	"document end",
	yaml_ALIAS_EVENT:		"alias",
	yaml_SCALAR_EVENT:		"scalar",
	yaml_SEQUENCE_START_EVENT:	"sequence start",
	yaml_SEQUENCE_END_EVENT:	"sequence end",
	yaml_MAPPING_START_EVENT:	"mapping start",
	yaml_MAPPING_END_EVENT:		"mapping end",
}

func (e yaml_event_type_t) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:257
	_go_fuzz_dep_.CoverTab[127997]++
									if e < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:258
		_go_fuzz_dep_.CoverTab[127999]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:258
		return int(e) >= len(eventStrings)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:258
		// _ = "end of CoverTab[127999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:258
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:258
		_go_fuzz_dep_.CoverTab[128000]++
										return fmt.Sprintf("unknown event %d", e)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:259
		// _ = "end of CoverTab[128000]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:260
		_go_fuzz_dep_.CoverTab[128001]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:260
		// _ = "end of CoverTab[128001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:260
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:260
	// _ = "end of CoverTab[127997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:260
	_go_fuzz_dep_.CoverTab[127998]++
									return eventStrings[e]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:261
	// _ = "end of CoverTab[127998]"
}

// The event structure.
type yaml_event_t struct {

	// The event type.
	typ	yaml_event_type_t

	// The start and end of the event.
	start_mark, end_mark	yaml_mark_t

	// The document encoding (for yaml_STREAM_START_EVENT).
	encoding	yaml_encoding_t

	// The version directive (for yaml_DOCUMENT_START_EVENT).
	version_directive	*yaml_version_directive_t

	// The list of tag directives (for yaml_DOCUMENT_START_EVENT).
	tag_directives	[]yaml_tag_directive_t

	// The anchor (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT, yaml_ALIAS_EVENT).
	anchor	[]byte

	// The tag (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT).
	tag	[]byte

	// The scalar value (for yaml_SCALAR_EVENT).
	value	[]byte

	// Is the document start/end indicator implicit, or the tag optional?
	// (for yaml_DOCUMENT_START_EVENT, yaml_DOCUMENT_END_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT, yaml_SCALAR_EVENT).
	implicit	bool

	// Is the tag optional for any non-plain style? (for yaml_SCALAR_EVENT).
	quoted_implicit	bool

	// The style (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT).
	style	yaml_style_t
}

func (e *yaml_event_t) scalar_style() yaml_scalar_style_t {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:302
	_go_fuzz_dep_.CoverTab[128002]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:302
	return yaml_scalar_style_t(e.style)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:302
	// _ = "end of CoverTab[128002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:302
}
func (e *yaml_event_t) sequence_style() yaml_sequence_style_t {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:303
	_go_fuzz_dep_.CoverTab[128003]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:303
	return yaml_sequence_style_t(e.style)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:303
	// _ = "end of CoverTab[128003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:303
}
func (e *yaml_event_t) mapping_style() yaml_mapping_style_t {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:304
	_go_fuzz_dep_.CoverTab[128004]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:304
	return yaml_mapping_style_t(e.style)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:304
	// _ = "end of CoverTab[128004]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:304
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:308
const (
	yaml_NULL_TAG		= "tag:yaml.org,2002:null"	// The tag !!null with the only possible value: null.
	yaml_BOOL_TAG		= "tag:yaml.org,2002:bool"	// The tag !!bool with the values: true and false.
	yaml_STR_TAG		= "tag:yaml.org,2002:str"	// The tag !!str for string values.
	yaml_INT_TAG		= "tag:yaml.org,2002:int"	// The tag !!int for integer values.
	yaml_FLOAT_TAG		= "tag:yaml.org,2002:float"	// The tag !!float for float values.
	yaml_TIMESTAMP_TAG	= "tag:yaml.org,2002:timestamp"	// The tag !!timestamp for date and time values.

	yaml_SEQ_TAG	= "tag:yaml.org,2002:seq"	// The tag !!seq is used to denote sequences.
	yaml_MAP_TAG	= "tag:yaml.org,2002:map"	// The tag !!map is used to denote mapping.

	// Not in original libyaml.
	yaml_BINARY_TAG	= "tag:yaml.org,2002:binary"
	yaml_MERGE_TAG	= "tag:yaml.org,2002:merge"

	yaml_DEFAULT_SCALAR_TAG		= yaml_STR_TAG	// The default scalar tag is !!str.
	yaml_DEFAULT_SEQUENCE_TAG	= yaml_SEQ_TAG	// The default sequence tag is !!seq.
	yaml_DEFAULT_MAPPING_TAG	= yaml_MAP_TAG	// The default mapping tag is !!map.
)

type yaml_node_type_t int

// Node types.
const (
	// An empty node.
	yaml_NO_NODE	yaml_node_type_t	= iota

	yaml_SCALAR_NODE	// A scalar node.
	yaml_SEQUENCE_NODE	// A sequence node.
	yaml_MAPPING_NODE	// A mapping node.
)

// An element of a sequence node.
type yaml_node_item_t int

// An element of a mapping node.
type yaml_node_pair_t struct {
	key	int	// The key of the element.
	value	int	// The value of the element.
}

// The node structure.
type yaml_node_t struct {
	typ	yaml_node_type_t	// The node type.
	tag	[]byte			// The node tag.

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:356
	// The scalar parameters (for yaml_SCALAR_NODE).
									scalar	struct {
		value	[]byte			// The scalar value.
		length	int			// The length of the scalar value.
		style	yaml_scalar_style_t	// The scalar style.
	}

	// The sequence parameters (for YAML_SEQUENCE_NODE).
	sequence	struct {
		items_data	[]yaml_node_item_t	// The stack of sequence items.
		style		yaml_sequence_style_t	// The sequence style.
	}

	// The mapping parameters (for yaml_MAPPING_NODE).
	mapping	struct {
		pairs_data	[]yaml_node_pair_t	// The stack of mapping pairs (key, value).
		pairs_start	*yaml_node_pair_t	// The beginning of the stack.
		pairs_end	*yaml_node_pair_t	// The end of the stack.
		pairs_top	*yaml_node_pair_t	// The top of the stack.
		style		yaml_mapping_style_t	// The mapping style.
	}

	start_mark	yaml_mark_t	// The beginning of the node.
	end_mark	yaml_mark_t	// The end of the node.

}

// The document structure.
type yaml_document_t struct {

	// The document nodes.
	nodes	[]yaml_node_t

	// The version directive.
	version_directive	*yaml_version_directive_t

	// The list of tag directives.
	tag_directives_data	[]yaml_tag_directive_t
	tag_directives_start	int	// The beginning of the tag directives list.
	tag_directives_end	int	// The end of the tag directives list.

	start_implicit	int	// Is the document start indicator implicit?
	end_implicit	int	// Is the document end indicator implicit?

	// The start/end of the document.
	start_mark, end_mark	yaml_mark_t
}

// The prototype of a read handler.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// The read handler is called when the parser needs to read more bytes from the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// source. The handler should write not more than size bytes to the buffer.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// The number of written bytes should be set to the size_read variable.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// [in,out]   data        A pointer to an application data specified by
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//	yaml_parser_set_input().
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// [out]      buffer      The buffer to write the data from the source.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// [in]       size        The size of the buffer.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// [out]      size_read   The actual number of bytes read from the source.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// On success, the handler should return 1.  If the handler failed,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// the returned value should be 0. On EOF, the handler should set the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:404
// size_read to 0 and return 1.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:419
type yaml_read_handler_t func(parser *yaml_parser_t, buffer []byte) (n int, err error)

// This structure holds information about a potential simple key.
type yaml_simple_key_t struct {
	possible	bool		// Is a simple key possible?
	required	bool		// Is a simple key required?
	token_number	int		// The number of the token.
	mark		yaml_mark_t	// The position mark.
}

// The states of the parser.
type yaml_parser_state_t int

const (
	yaml_PARSE_STREAM_START_STATE	yaml_parser_state_t	= iota

	yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE		// Expect the beginning of an implicit document.
	yaml_PARSE_DOCUMENT_START_STATE				// Expect DOCUMENT-START.
	yaml_PARSE_DOCUMENT_CONTENT_STATE			// Expect the content of a document.
	yaml_PARSE_DOCUMENT_END_STATE				// Expect DOCUMENT-END.
	yaml_PARSE_BLOCK_NODE_STATE				// Expect a block node.
	yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE	// Expect a block node or indentless sequence.
	yaml_PARSE_FLOW_NODE_STATE				// Expect a flow node.
	yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE		// Expect the first entry of a block sequence.
	yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE			// Expect an entry of a block sequence.
	yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE		// Expect an entry of an indentless sequence.
	yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE		// Expect the first key of a block mapping.
	yaml_PARSE_BLOCK_MAPPING_KEY_STATE			// Expect a block mapping key.
	yaml_PARSE_BLOCK_MAPPING_VALUE_STATE			// Expect a block mapping value.
	yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE		// Expect the first entry of a flow sequence.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE			// Expect an entry of a flow sequence.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE	// Expect a key of an ordered mapping.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE	// Expect a value of an ordered mapping.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE	// Expect the and of an ordered mapping entry.
	yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE			// Expect the first key of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_KEY_STATE			// Expect a key of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_VALUE_STATE			// Expect a value of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE		// Expect an empty value of a flow mapping.
	yaml_PARSE_END_STATE					// Expect nothing.
)

func (ps yaml_parser_state_t) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:460
	_go_fuzz_dep_.CoverTab[128005]++
									switch ps {
	case yaml_PARSE_STREAM_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:462
		_go_fuzz_dep_.CoverTab[128007]++
										return "yaml_PARSE_STREAM_START_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:463
		// _ = "end of CoverTab[128007]"
	case yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:464
		_go_fuzz_dep_.CoverTab[128008]++
										return "yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:465
		// _ = "end of CoverTab[128008]"
	case yaml_PARSE_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:466
		_go_fuzz_dep_.CoverTab[128009]++
										return "yaml_PARSE_DOCUMENT_START_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:467
		// _ = "end of CoverTab[128009]"
	case yaml_PARSE_DOCUMENT_CONTENT_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:468
		_go_fuzz_dep_.CoverTab[128010]++
										return "yaml_PARSE_DOCUMENT_CONTENT_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:469
		// _ = "end of CoverTab[128010]"
	case yaml_PARSE_DOCUMENT_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:470
		_go_fuzz_dep_.CoverTab[128011]++
										return "yaml_PARSE_DOCUMENT_END_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:471
		// _ = "end of CoverTab[128011]"
	case yaml_PARSE_BLOCK_NODE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:472
		_go_fuzz_dep_.CoverTab[128012]++
										return "yaml_PARSE_BLOCK_NODE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:473
		// _ = "end of CoverTab[128012]"
	case yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:474
		_go_fuzz_dep_.CoverTab[128013]++
										return "yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:475
		// _ = "end of CoverTab[128013]"
	case yaml_PARSE_FLOW_NODE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:476
		_go_fuzz_dep_.CoverTab[128014]++
										return "yaml_PARSE_FLOW_NODE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:477
		// _ = "end of CoverTab[128014]"
	case yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:478
		_go_fuzz_dep_.CoverTab[128015]++
										return "yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:479
		// _ = "end of CoverTab[128015]"
	case yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:480
		_go_fuzz_dep_.CoverTab[128016]++
										return "yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:481
		// _ = "end of CoverTab[128016]"
	case yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:482
		_go_fuzz_dep_.CoverTab[128017]++
										return "yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:483
		// _ = "end of CoverTab[128017]"
	case yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:484
		_go_fuzz_dep_.CoverTab[128018]++
										return "yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:485
		// _ = "end of CoverTab[128018]"
	case yaml_PARSE_BLOCK_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:486
		_go_fuzz_dep_.CoverTab[128019]++
										return "yaml_PARSE_BLOCK_MAPPING_KEY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:487
		// _ = "end of CoverTab[128019]"
	case yaml_PARSE_BLOCK_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:488
		_go_fuzz_dep_.CoverTab[128020]++
										return "yaml_PARSE_BLOCK_MAPPING_VALUE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:489
		// _ = "end of CoverTab[128020]"
	case yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:490
		_go_fuzz_dep_.CoverTab[128021]++
										return "yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:491
		// _ = "end of CoverTab[128021]"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:492
		_go_fuzz_dep_.CoverTab[128022]++
										return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:493
		// _ = "end of CoverTab[128022]"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:494
		_go_fuzz_dep_.CoverTab[128023]++
										return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:495
		// _ = "end of CoverTab[128023]"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:496
		_go_fuzz_dep_.CoverTab[128024]++
										return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:497
		// _ = "end of CoverTab[128024]"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:498
		_go_fuzz_dep_.CoverTab[128025]++
										return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:499
		// _ = "end of CoverTab[128025]"
	case yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:500
		_go_fuzz_dep_.CoverTab[128026]++
										return "yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:501
		// _ = "end of CoverTab[128026]"
	case yaml_PARSE_FLOW_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:502
		_go_fuzz_dep_.CoverTab[128027]++
										return "yaml_PARSE_FLOW_MAPPING_KEY_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:503
		// _ = "end of CoverTab[128027]"
	case yaml_PARSE_FLOW_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:504
		_go_fuzz_dep_.CoverTab[128028]++
										return "yaml_PARSE_FLOW_MAPPING_VALUE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:505
		// _ = "end of CoverTab[128028]"
	case yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:506
		_go_fuzz_dep_.CoverTab[128029]++
										return "yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:507
		// _ = "end of CoverTab[128029]"
	case yaml_PARSE_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:508
		_go_fuzz_dep_.CoverTab[128030]++
										return "yaml_PARSE_END_STATE"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:509
		// _ = "end of CoverTab[128030]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:509
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:509
		_go_fuzz_dep_.CoverTab[128031]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:509
		// _ = "end of CoverTab[128031]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:510
	// _ = "end of CoverTab[128005]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:510
	_go_fuzz_dep_.CoverTab[128006]++
									return "<unknown parser state>"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:511
	// _ = "end of CoverTab[128006]"
}

// This structure holds aliases data.
type yaml_alias_data_t struct {
	anchor	[]byte		// The anchor.
	index	int		// The node id.
	mark	yaml_mark_t	// The anchor mark.
}

// The parser structure.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:521
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:521
// All members are internal. Manage the structure using the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:521
// yaml_parser_ family of functions.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:525
type yaml_parser_t struct {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:529
	error	yaml_error_type_t	// Error type.

	problem	string	// Error description.

	// The byte about which the problem occurred.
	problem_offset	int
	problem_value	int
	problem_mark	yaml_mark_t

									// The error context.
									context		string
									context_mark	yaml_mark_t

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:544
	read_handler	yaml_read_handler_t	// Read handler.

									input_reader	io.Reader	// File input data.
									input		[]byte		// String input data.
									input_pos	int

									eof	bool	// EOF flag

									buffer		[]byte	// The working buffer.
									buffer_pos	int	// The current position of the buffer.

									unread	int	// The number of unread characters in the buffer.

									raw_buffer	[]byte	// The raw buffer.
									raw_buffer_pos	int	// The current position of the buffer.

									encoding	yaml_encoding_t	// The input encoding.

									offset	int		// The offset of the current position (in bytes).
									mark	yaml_mark_t	// The mark of the current position.

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:567
	stream_start_produced	bool	// Have we started to scan the input stream?
									stream_end_produced	bool	// Have we reached the end of the input stream?

									flow_level	int	// The number of unclosed '[' and '{' indicators.

									tokens		[]yaml_token_t	// The tokens queue.
									tokens_head	int		// The head of the tokens queue.
									tokens_parsed	int		// The number of tokens fetched from the queue.
									token_available	bool		// Does the tokens queue contain a token ready for dequeueing.

									indent	int	// The current indentation level.
									indents	[]int	// The indentation levels stack.

									simple_key_allowed	bool			// May a simple key occur at the current position?
									simple_keys		[]yaml_simple_key_t	// The stack of simple keys.
									simple_keys_by_tok	map[int]int		// possible simple_key indexes indexed by token_number

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:586
	state	yaml_parser_state_t	// The current parser state.
									states		[]yaml_parser_state_t	// The parser states stack.
									marks		[]yaml_mark_t		// The stack of marks.
									tag_directives	[]yaml_tag_directive_t	// The list of TAG directives.

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:593
	aliases	[]yaml_alias_data_t	// The alias data.

	document	*yaml_document_t	// The currently parsed document.
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// The prototype of a write handler.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// The write handler is called when the emitter needs to flush the accumulated
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// characters to the output.  The handler should write @a size bytes of the
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// @a buffer to the output.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// @param[in,out]   data        A pointer to an application data specified by
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//	yaml_emitter_set_output().
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// @param[in]       buffer      The buffer with bytes to be written.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// @param[in]       size        The size of the buffer.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// @returns On success, the handler should return @c 1.  If the handler failed,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:600
// the returned value should be @c 0.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:614
type yaml_write_handler_t func(emitter *yaml_emitter_t, buffer []byte) error

type yaml_emitter_state_t int

// The emitter states.
const (
	// Expect STREAM-START.
	yaml_EMIT_STREAM_START_STATE	yaml_emitter_state_t	= iota

	yaml_EMIT_FIRST_DOCUMENT_START_STATE		// Expect the first DOCUMENT-START or STREAM-END.
	yaml_EMIT_DOCUMENT_START_STATE			// Expect DOCUMENT-START or STREAM-END.
	yaml_EMIT_DOCUMENT_CONTENT_STATE		// Expect the content of a document.
	yaml_EMIT_DOCUMENT_END_STATE			// Expect DOCUMENT-END.
	yaml_EMIT_FLOW_SEQUENCE_FIRST_ITEM_STATE	// Expect the first item of a flow sequence.
	yaml_EMIT_FLOW_SEQUENCE_ITEM_STATE		// Expect an item of a flow sequence.
	yaml_EMIT_FLOW_MAPPING_FIRST_KEY_STATE		// Expect the first key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_KEY_STATE		// Expect a key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_SIMPLE_VALUE_STATE	// Expect a value for a simple key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_VALUE_STATE		// Expect a value of a flow mapping.
	yaml_EMIT_BLOCK_SEQUENCE_FIRST_ITEM_STATE	// Expect the first item of a block sequence.
	yaml_EMIT_BLOCK_SEQUENCE_ITEM_STATE		// Expect an item of a block sequence.
	yaml_EMIT_BLOCK_MAPPING_FIRST_KEY_STATE		// Expect the first key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_KEY_STATE		// Expect the key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_SIMPLE_VALUE_STATE	// Expect a value for a simple key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_VALUE_STATE		// Expect a value of a block mapping.
	yaml_EMIT_END_STATE				// Expect nothing.
)

// The emitter structure.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:642
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:642
// All members are internal.  Manage the structure using the @c yaml_emitter_
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:642
// family of functions.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:646
type yaml_emitter_t struct {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:650
	error	yaml_error_type_t	// Error type.
									problem	string	// Error description.

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:655
	write_handler	yaml_write_handler_t	// Write handler.

									output_buffer	*[]byte		// String output data.
									output_writer	io.Writer	// File output data.

									buffer		[]byte	// The working buffer.
									buffer_pos	int	// The current position of the buffer.

									raw_buffer	[]byte	// The raw buffer.
									raw_buffer_pos	int	// The current position of the buffer.

									encoding	yaml_encoding_t	// The stream encoding.

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:670
	canonical	bool	// If the output is in the canonical style?
	best_indent	int		// The number of indentation spaces.
	best_width	int		// The preferred width of the output lines.
	unicode		bool		// Allow unescaped non-ASCII characters?
	line_break	yaml_break_t	// The preferred line break.

	state	yaml_emitter_state_t	// The current emitter state.
	states	[]yaml_emitter_state_t	// The stack of states.

	events		[]yaml_event_t	// The event queue.
	events_head	int		// The head of the event queue.

	indents	[]int	// The stack of indentation levels.

	tag_directives	[]yaml_tag_directive_t	// The list of tag directives.

	indent	int	// The current indentation level.

	flow_level	int	// The current flow level.

	root_context		bool	// Is it the document root context?
	sequence_context	bool	// Is it a sequence context?
	mapping_context		bool	// Is it a mapping context?
	simple_key_context	bool	// Is it a simple mapping key context?

	line		int	// The current line.
	column		int	// The current column.
	whitespace	bool	// If the last character was a whitespace?
	indention	bool	// If the last character was an indentation character (' ', '-', '?', ':')?
	open_ended	bool	// If an explicit document end is required?

	// Anchor analysis.
	anchor_data	struct {
		anchor	[]byte	// The anchor value.
		alias	bool	// Is it an alias?
	}

	// Tag analysis.
	tag_data	struct {
		handle	[]byte	// The tag handle.
		suffix	[]byte	// The tag suffix.
	}

	// Scalar analysis.
	scalar_data	struct {
		value			[]byte			// The scalar value.
		multiline		bool			// Does the scalar contain line breaks?
		flow_plain_allowed	bool			// Can the scalar be expessed in the flow plain style?
		block_plain_allowed	bool			// Can the scalar be expressed in the block plain style?
		single_quoted_allowed	bool			// Can the scalar be expressed in the single quoted style?
		block_allowed		bool			// Can the scalar be expressed in the literal or folded styles?
		style			yaml_scalar_style_t	// The output style.
	}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:726
	opened	bool	// If the stream was already opened?
	closed	bool	// If the stream was already closed?

	// The information associated with the document nodes.
	anchors	*struct {
		references	int	// The number of references.
		anchor		int	// The anchor id.
		serialized	bool	// If the node has been emitted?
	}

	last_anchor_id	int	// The last assigned anchor id.

	document	*yaml_document_t	// The currently emitted document.
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:739
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlh.go:739
var _ = _go_fuzz_dep_.CoverTab
