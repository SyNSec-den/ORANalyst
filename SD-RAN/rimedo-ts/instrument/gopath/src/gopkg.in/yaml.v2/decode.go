//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:1
)

import (
	"encoding"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"time"
)

const (
	documentNode	= 1 << iota
	mappingNode
	sequenceNode
	scalarNode
	aliasNode
)

type node struct {
	kind		int
	line, column	int
	tag		string
	// For an alias node, alias holds the resolved alias.
	alias		*node
	value		string
	implicit	bool
	children	[]*node
	anchors		map[string]*node
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:37
type parser struct {
	parser		yaml_parser_t
	event		yaml_event_t
	doc		*node
	doneInit	bool
}

func newParser(b []byte) *parser {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:44
	_go_fuzz_dep_.CoverTab[124511]++
									p := parser{}
									if !yaml_parser_initialize(&p.parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:46
		_go_fuzz_dep_.CoverTab[124514]++
										panic("failed to initialize YAML emitter")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:47
		// _ = "end of CoverTab[124514]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:48
		_go_fuzz_dep_.CoverTab[124515]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:48
		// _ = "end of CoverTab[124515]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:48
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:48
	// _ = "end of CoverTab[124511]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:48
	_go_fuzz_dep_.CoverTab[124512]++
									if len(b) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:49
		_go_fuzz_dep_.CoverTab[124516]++
										b = []byte{'\n'}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:50
		// _ = "end of CoverTab[124516]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:51
		_go_fuzz_dep_.CoverTab[124517]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:51
		// _ = "end of CoverTab[124517]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:51
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:51
	// _ = "end of CoverTab[124512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:51
	_go_fuzz_dep_.CoverTab[124513]++
									yaml_parser_set_input_string(&p.parser, b)
									return &p
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:53
	// _ = "end of CoverTab[124513]"
}

func newParserFromReader(r io.Reader) *parser {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:56
	_go_fuzz_dep_.CoverTab[124518]++
									p := parser{}
									if !yaml_parser_initialize(&p.parser) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:58
		_go_fuzz_dep_.CoverTab[124520]++
										panic("failed to initialize YAML emitter")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:59
		// _ = "end of CoverTab[124520]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:60
		_go_fuzz_dep_.CoverTab[124521]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:60
		// _ = "end of CoverTab[124521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:60
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:60
	// _ = "end of CoverTab[124518]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:60
	_go_fuzz_dep_.CoverTab[124519]++
									yaml_parser_set_input_reader(&p.parser, r)
									return &p
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:62
	// _ = "end of CoverTab[124519]"
}

func (p *parser) init() {
	if p.doneInit {
		return
	}
	p.expect(yaml_STREAM_START_EVENT)
	p.doneInit = true
}

func (p *parser) destroy() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:73
	_go_fuzz_dep_.CoverTab[124522]++
									if p.event.typ != yaml_NO_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:74
		_go_fuzz_dep_.CoverTab[124524]++
										yaml_event_delete(&p.event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:75
		// _ = "end of CoverTab[124524]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:76
		_go_fuzz_dep_.CoverTab[124525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:76
		// _ = "end of CoverTab[124525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:76
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:76
	// _ = "end of CoverTab[124522]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:76
	_go_fuzz_dep_.CoverTab[124523]++
									yaml_parser_delete(&p.parser)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:77
	// _ = "end of CoverTab[124523]"
}

// expect consumes an event from the event stream and
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:80
// checks that it's of the expected type.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:82
func (p *parser) expect(e yaml_event_type_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:82
	_go_fuzz_dep_.CoverTab[124526]++
									if p.event.typ == yaml_NO_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:83
		_go_fuzz_dep_.CoverTab[124530]++
										if !yaml_parser_parse(&p.parser, &p.event) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:84
			_go_fuzz_dep_.CoverTab[124531]++
											p.fail()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:85
			// _ = "end of CoverTab[124531]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:86
			_go_fuzz_dep_.CoverTab[124532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:86
			// _ = "end of CoverTab[124532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:86
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:86
		// _ = "end of CoverTab[124530]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:87
		_go_fuzz_dep_.CoverTab[124533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:87
		// _ = "end of CoverTab[124533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:87
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:87
	// _ = "end of CoverTab[124526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:87
	_go_fuzz_dep_.CoverTab[124527]++
									if p.event.typ == yaml_STREAM_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:88
		_go_fuzz_dep_.CoverTab[124534]++
										failf("attempted to go past the end of stream; corrupted value?")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:89
		// _ = "end of CoverTab[124534]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:90
		_go_fuzz_dep_.CoverTab[124535]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:90
		// _ = "end of CoverTab[124535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:90
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:90
	// _ = "end of CoverTab[124527]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:90
	_go_fuzz_dep_.CoverTab[124528]++
									if p.event.typ != e {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:91
		_go_fuzz_dep_.CoverTab[124536]++
										p.parser.problem = fmt.Sprintf("expected %s event but got %s", e, p.event.typ)
										p.fail()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:93
		// _ = "end of CoverTab[124536]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:94
		_go_fuzz_dep_.CoverTab[124537]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:94
		// _ = "end of CoverTab[124537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:94
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:94
	// _ = "end of CoverTab[124528]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:94
	_go_fuzz_dep_.CoverTab[124529]++
									yaml_event_delete(&p.event)
									p.event.typ = yaml_NO_EVENT
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:96
	// _ = "end of CoverTab[124529]"
}

// peek peeks at the next event in the event stream,
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:99
// puts the results into p.event and returns the event type.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:101
func (p *parser) peek() yaml_event_type_t {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:101
	_go_fuzz_dep_.CoverTab[124538]++
										if p.event.typ != yaml_NO_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:102
		_go_fuzz_dep_.CoverTab[124541]++
											return p.event.typ
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:103
		// _ = "end of CoverTab[124541]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:104
		_go_fuzz_dep_.CoverTab[124542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:104
		// _ = "end of CoverTab[124542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:104
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:104
	// _ = "end of CoverTab[124538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:104
	_go_fuzz_dep_.CoverTab[124539]++
										if !yaml_parser_parse(&p.parser, &p.event) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:105
		_go_fuzz_dep_.CoverTab[124543]++
											p.fail()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:106
		// _ = "end of CoverTab[124543]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:107
		_go_fuzz_dep_.CoverTab[124544]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:107
		// _ = "end of CoverTab[124544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:107
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:107
	// _ = "end of CoverTab[124539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:107
	_go_fuzz_dep_.CoverTab[124540]++
										return p.event.typ
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:108
	// _ = "end of CoverTab[124540]"
}

func (p *parser) fail() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:111
	_go_fuzz_dep_.CoverTab[124545]++
										var where string
										var line int
										if p.parser.problem_mark.line != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:114
		_go_fuzz_dep_.CoverTab[124549]++
											line = p.parser.problem_mark.line

											if p.parser.error == yaml_SCANNER_ERROR {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:117
			_go_fuzz_dep_.CoverTab[124550]++
												line++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:118
			// _ = "end of CoverTab[124550]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:119
			_go_fuzz_dep_.CoverTab[124551]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:119
			// _ = "end of CoverTab[124551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:119
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:119
		// _ = "end of CoverTab[124549]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:120
		_go_fuzz_dep_.CoverTab[124552]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:120
		if p.parser.context_mark.line != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:120
			_go_fuzz_dep_.CoverTab[124553]++
												line = p.parser.context_mark.line
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:121
			// _ = "end of CoverTab[124553]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
			_go_fuzz_dep_.CoverTab[124554]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
			// _ = "end of CoverTab[124554]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
		// _ = "end of CoverTab[124552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
	// _ = "end of CoverTab[124545]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:122
	_go_fuzz_dep_.CoverTab[124546]++
										if line != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:123
		_go_fuzz_dep_.CoverTab[124555]++
											where = "line " + strconv.Itoa(line) + ": "
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:124
		// _ = "end of CoverTab[124555]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:125
		_go_fuzz_dep_.CoverTab[124556]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:125
		// _ = "end of CoverTab[124556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:125
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:125
	// _ = "end of CoverTab[124546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:125
	_go_fuzz_dep_.CoverTab[124547]++
										var msg string
										if len(p.parser.problem) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:127
		_go_fuzz_dep_.CoverTab[124557]++
											msg = p.parser.problem
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:128
		// _ = "end of CoverTab[124557]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:129
		_go_fuzz_dep_.CoverTab[124558]++
											msg = "unknown problem parsing YAML content"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:130
		// _ = "end of CoverTab[124558]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:131
	// _ = "end of CoverTab[124547]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:131
	_go_fuzz_dep_.CoverTab[124548]++
										failf("%s%s", where, msg)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:132
	// _ = "end of CoverTab[124548]"
}

func (p *parser) anchor(n *node, anchor []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:135
	_go_fuzz_dep_.CoverTab[124559]++
										if anchor != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:136
		_go_fuzz_dep_.CoverTab[124560]++
											p.doc.anchors[string(anchor)] = n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:137
		// _ = "end of CoverTab[124560]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:138
		_go_fuzz_dep_.CoverTab[124561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:138
		// _ = "end of CoverTab[124561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:138
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:138
	// _ = "end of CoverTab[124559]"
}

func (p *parser) parse() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:141
	_go_fuzz_dep_.CoverTab[124562]++
										p.init()
										switch p.peek() {
	case yaml_SCALAR_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:144
		_go_fuzz_dep_.CoverTab[124563]++
											return p.scalar()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:145
		// _ = "end of CoverTab[124563]"
	case yaml_ALIAS_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:146
		_go_fuzz_dep_.CoverTab[124564]++
											return p.alias()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:147
		// _ = "end of CoverTab[124564]"
	case yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:148
		_go_fuzz_dep_.CoverTab[124565]++
											return p.mapping()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:149
		// _ = "end of CoverTab[124565]"
	case yaml_SEQUENCE_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:150
		_go_fuzz_dep_.CoverTab[124566]++
											return p.sequence()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:151
		// _ = "end of CoverTab[124566]"
	case yaml_DOCUMENT_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:152
		_go_fuzz_dep_.CoverTab[124567]++
											return p.document()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:153
		// _ = "end of CoverTab[124567]"
	case yaml_STREAM_END_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:154
		_go_fuzz_dep_.CoverTab[124568]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:156
		// _ = "end of CoverTab[124568]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:157
		_go_fuzz_dep_.CoverTab[124569]++
											panic("attempted to parse unknown event: " + p.event.typ.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:158
		// _ = "end of CoverTab[124569]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:159
	// _ = "end of CoverTab[124562]"
}

func (p *parser) node(kind int) *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:162
	_go_fuzz_dep_.CoverTab[124570]++
										return &node{
		kind:	kind,
		line:	p.event.start_mark.line,
		column:	p.event.start_mark.column,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:167
	// _ = "end of CoverTab[124570]"
}

func (p *parser) document() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:170
	_go_fuzz_dep_.CoverTab[124571]++
										n := p.node(documentNode)
										n.anchors = make(map[string]*node)
										p.doc = n
										p.expect(yaml_DOCUMENT_START_EVENT)
										n.children = append(n.children, p.parse())
										p.expect(yaml_DOCUMENT_END_EVENT)
										return n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:177
	// _ = "end of CoverTab[124571]"
}

func (p *parser) alias() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:180
	_go_fuzz_dep_.CoverTab[124572]++
										n := p.node(aliasNode)
										n.value = string(p.event.anchor)
										n.alias = p.doc.anchors[n.value]
										if n.alias == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:184
		_go_fuzz_dep_.CoverTab[124574]++
											failf("unknown anchor '%s' referenced", n.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:185
		// _ = "end of CoverTab[124574]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:186
		_go_fuzz_dep_.CoverTab[124575]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:186
		// _ = "end of CoverTab[124575]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:186
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:186
	// _ = "end of CoverTab[124572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:186
	_go_fuzz_dep_.CoverTab[124573]++
										p.expect(yaml_ALIAS_EVENT)
										return n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:188
	// _ = "end of CoverTab[124573]"
}

func (p *parser) scalar() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:191
	_go_fuzz_dep_.CoverTab[124576]++
										n := p.node(scalarNode)
										n.value = string(p.event.value)
										n.tag = string(p.event.tag)
										n.implicit = p.event.implicit
										p.anchor(n, p.event.anchor)
										p.expect(yaml_SCALAR_EVENT)
										return n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:198
	// _ = "end of CoverTab[124576]"
}

func (p *parser) sequence() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:201
	_go_fuzz_dep_.CoverTab[124577]++
										n := p.node(sequenceNode)
										p.anchor(n, p.event.anchor)
										p.expect(yaml_SEQUENCE_START_EVENT)
										for p.peek() != yaml_SEQUENCE_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:205
		_go_fuzz_dep_.CoverTab[124579]++
											n.children = append(n.children, p.parse())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:206
		// _ = "end of CoverTab[124579]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:207
	// _ = "end of CoverTab[124577]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:207
	_go_fuzz_dep_.CoverTab[124578]++
										p.expect(yaml_SEQUENCE_END_EVENT)
										return n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:209
	// _ = "end of CoverTab[124578]"
}

func (p *parser) mapping() *node {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:212
	_go_fuzz_dep_.CoverTab[124580]++
										n := p.node(mappingNode)
										p.anchor(n, p.event.anchor)
										p.expect(yaml_MAPPING_START_EVENT)
										for p.peek() != yaml_MAPPING_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:216
		_go_fuzz_dep_.CoverTab[124582]++
											n.children = append(n.children, p.parse(), p.parse())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:217
		// _ = "end of CoverTab[124582]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:218
	// _ = "end of CoverTab[124580]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:218
	_go_fuzz_dep_.CoverTab[124581]++
										p.expect(yaml_MAPPING_END_EVENT)
										return n
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:220
	// _ = "end of CoverTab[124581]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:226
type decoder struct {
	doc	*node
	aliases	map[*node]bool
	mapType	reflect.Type
	terrors	[]string
	strict	bool

	decodeCount	int
	aliasCount	int
	aliasDepth	int
}

var (
	mapItemType	= reflect.TypeOf(MapItem{})
	durationType	= reflect.TypeOf(time.Duration(0))
	defaultMapType	= reflect.TypeOf(map[interface{}]interface{}{})
	ifaceType	= defaultMapType.Elem()
	timeType	= reflect.TypeOf(time.Time{})
	ptrTimeType	= reflect.TypeOf(&time.Time{})
)

func newDecoder(strict bool) *decoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:247
	_go_fuzz_dep_.CoverTab[124583]++
										d := &decoder{mapType: defaultMapType, strict: strict}
										d.aliases = make(map[*node]bool)
										return d
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:250
	// _ = "end of CoverTab[124583]"
}

func (d *decoder) terror(n *node, tag string, out reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:253
	_go_fuzz_dep_.CoverTab[124584]++
										if n.tag != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:254
		_go_fuzz_dep_.CoverTab[124587]++
											tag = n.tag
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:255
		// _ = "end of CoverTab[124587]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:256
		_go_fuzz_dep_.CoverTab[124588]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:256
		// _ = "end of CoverTab[124588]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:256
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:256
	// _ = "end of CoverTab[124584]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:256
	_go_fuzz_dep_.CoverTab[124585]++
										value := n.value
										if tag != yaml_SEQ_TAG && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:258
		_go_fuzz_dep_.CoverTab[124589]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:258
		return tag != yaml_MAP_TAG
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:258
		// _ = "end of CoverTab[124589]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:258
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:258
		_go_fuzz_dep_.CoverTab[124590]++
											if len(value) > 10 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:259
			_go_fuzz_dep_.CoverTab[124591]++
												value = " `" + value[:7] + "...`"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:260
			// _ = "end of CoverTab[124591]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:261
			_go_fuzz_dep_.CoverTab[124592]++
												value = " `" + value + "`"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:262
			// _ = "end of CoverTab[124592]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:263
		// _ = "end of CoverTab[124590]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:264
		_go_fuzz_dep_.CoverTab[124593]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:264
		// _ = "end of CoverTab[124593]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:264
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:264
	// _ = "end of CoverTab[124585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:264
	_go_fuzz_dep_.CoverTab[124586]++
										d.terrors = append(d.terrors, fmt.Sprintf("line %d: cannot unmarshal %s%s into %s", n.line+1, shortTag(tag), value, out.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:265
	// _ = "end of CoverTab[124586]"
}

func (d *decoder) callUnmarshaler(n *node, u Unmarshaler) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:268
	_go_fuzz_dep_.CoverTab[124594]++
										terrlen := len(d.terrors)
										err := u.UnmarshalYAML(func(v interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:270
		_go_fuzz_dep_.CoverTab[124598]++
											defer handleErr(&err)
											d.unmarshal(n, reflect.ValueOf(v))
											if len(d.terrors) > terrlen {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:273
			_go_fuzz_dep_.CoverTab[124600]++
												issues := d.terrors[terrlen:]
												d.terrors = d.terrors[:terrlen]
												return &TypeError{issues}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:276
			// _ = "end of CoverTab[124600]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:277
			_go_fuzz_dep_.CoverTab[124601]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:277
			// _ = "end of CoverTab[124601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:277
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:277
		// _ = "end of CoverTab[124598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:277
		_go_fuzz_dep_.CoverTab[124599]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:278
		// _ = "end of CoverTab[124599]"
	})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:279
	// _ = "end of CoverTab[124594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:279
	_go_fuzz_dep_.CoverTab[124595]++
										if e, ok := err.(*TypeError); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:280
		_go_fuzz_dep_.CoverTab[124602]++
											d.terrors = append(d.terrors, e.Errors...)
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:282
		// _ = "end of CoverTab[124602]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:283
		_go_fuzz_dep_.CoverTab[124603]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:283
		// _ = "end of CoverTab[124603]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:283
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:283
	// _ = "end of CoverTab[124595]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:283
	_go_fuzz_dep_.CoverTab[124596]++
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:284
		_go_fuzz_dep_.CoverTab[124604]++
											fail(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:285
		// _ = "end of CoverTab[124604]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:286
		_go_fuzz_dep_.CoverTab[124605]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:286
		// _ = "end of CoverTab[124605]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:286
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:286
	// _ = "end of CoverTab[124596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:286
	_go_fuzz_dep_.CoverTab[124597]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:287
	// _ = "end of CoverTab[124597]"
}

// d.prepare initializes and dereferences pointers and calls UnmarshalYAML
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
// if a value is found to implement it.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
// It returns the initialized and dereferenced out value, whether
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
// unmarshalling was already done by UnmarshalYAML, and if so whether
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
// its types unmarshalled appropriately.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:290
// If n holds a null value, prepare returns before doing anything.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:297
func (d *decoder) prepare(n *node, out reflect.Value) (newout reflect.Value, unmarshaled, good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:297
	_go_fuzz_dep_.CoverTab[124606]++
										if n.tag == yaml_NULL_TAG || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		_go_fuzz_dep_.CoverTab[124609]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		return n.kind == scalarNode && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			_go_fuzz_dep_.CoverTab[124610]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			return n.tag == ""
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			// _ = "end of CoverTab[124610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			_go_fuzz_dep_.CoverTab[124611]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			return (n.value == "null" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				_go_fuzz_dep_.CoverTab[124612]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				return n.value == "~"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				// _ = "end of CoverTab[124612]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				_go_fuzz_dep_.CoverTab[124613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				return n.value == "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
					_go_fuzz_dep_.CoverTab[124614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
					return n.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
					// _ = "end of CoverTab[124614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
				// _ = "end of CoverTab[124613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
			// _ = "end of CoverTab[124611]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		// _ = "end of CoverTab[124609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:298
		_go_fuzz_dep_.CoverTab[124615]++
											return out, false, false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:299
		// _ = "end of CoverTab[124615]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:300
		_go_fuzz_dep_.CoverTab[124616]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:300
		// _ = "end of CoverTab[124616]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:300
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:300
	// _ = "end of CoverTab[124606]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:300
	_go_fuzz_dep_.CoverTab[124607]++
										again := true
										for again {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:302
		_go_fuzz_dep_.CoverTab[124617]++
											again = false
											if out.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:304
			_go_fuzz_dep_.CoverTab[124619]++
												if out.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:305
				_go_fuzz_dep_.CoverTab[124621]++
													out.Set(reflect.New(out.Type().Elem()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:306
				// _ = "end of CoverTab[124621]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:307
				_go_fuzz_dep_.CoverTab[124622]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:307
				// _ = "end of CoverTab[124622]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:307
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:307
			// _ = "end of CoverTab[124619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:307
			_go_fuzz_dep_.CoverTab[124620]++
												out = out.Elem()
												again = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:309
			// _ = "end of CoverTab[124620]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:310
			_go_fuzz_dep_.CoverTab[124623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:310
			// _ = "end of CoverTab[124623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:310
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:310
		// _ = "end of CoverTab[124617]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:310
		_go_fuzz_dep_.CoverTab[124618]++
											if out.CanAddr() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:311
			_go_fuzz_dep_.CoverTab[124624]++
												if u, ok := out.Addr().Interface().(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:312
				_go_fuzz_dep_.CoverTab[124625]++
													good = d.callUnmarshaler(n, u)
													return out, true, good
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:314
				// _ = "end of CoverTab[124625]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:315
				_go_fuzz_dep_.CoverTab[124626]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:315
				// _ = "end of CoverTab[124626]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:315
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:315
			// _ = "end of CoverTab[124624]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:316
			_go_fuzz_dep_.CoverTab[124627]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:316
			// _ = "end of CoverTab[124627]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:316
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:316
		// _ = "end of CoverTab[124618]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:317
	// _ = "end of CoverTab[124607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:317
	_go_fuzz_dep_.CoverTab[124608]++
										return out, false, false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:318
	// _ = "end of CoverTab[124608]"
}

const (
	// 400,000 decode operations is ~500kb of dense object declarations, or
	// ~5kb of dense object declarations with 10000% alias expansion
	alias_ratio_range_low	= 400000

	// 4,000,000 decode operations is ~5MB of dense object declarations, or
	// ~4.5MB of dense object declarations with 10% alias expansion
	alias_ratio_range_high	= 4000000

	// alias_ratio_range is the range over which we scale allowed alias ratios
	alias_ratio_range	= float64(alias_ratio_range_high - alias_ratio_range_low)
)

func allowedAliasRatio(decodeCount int) float64 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:334
	_go_fuzz_dep_.CoverTab[124628]++
										switch {
	case decodeCount <= alias_ratio_range_low:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:336
		_go_fuzz_dep_.CoverTab[124629]++

											return 0.99
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:338
		// _ = "end of CoverTab[124629]"
	case decodeCount >= alias_ratio_range_high:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:339
		_go_fuzz_dep_.CoverTab[124630]++

											return 0.10
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:341
		// _ = "end of CoverTab[124630]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:342
		_go_fuzz_dep_.CoverTab[124631]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:346
		return 0.99 - 0.89*(float64(decodeCount-alias_ratio_range_low)/alias_ratio_range)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:346
		// _ = "end of CoverTab[124631]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:347
	// _ = "end of CoverTab[124628]"
}

func (d *decoder) unmarshal(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:350
	_go_fuzz_dep_.CoverTab[124632]++
										d.decodeCount++
										if d.aliasDepth > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:352
		_go_fuzz_dep_.CoverTab[124638]++
											d.aliasCount++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:353
		// _ = "end of CoverTab[124638]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:354
		_go_fuzz_dep_.CoverTab[124639]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:354
		// _ = "end of CoverTab[124639]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:354
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:354
	// _ = "end of CoverTab[124632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:354
	_go_fuzz_dep_.CoverTab[124633]++
										if d.aliasCount > 100 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		_go_fuzz_dep_.CoverTab[124640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		return d.decodeCount > 1000
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		// _ = "end of CoverTab[124640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		_go_fuzz_dep_.CoverTab[124641]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		return float64(d.aliasCount)/float64(d.decodeCount) > allowedAliasRatio(d.decodeCount)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		// _ = "end of CoverTab[124641]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:355
		_go_fuzz_dep_.CoverTab[124642]++
											failf("document contains excessive aliasing")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:356
		// _ = "end of CoverTab[124642]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:357
		_go_fuzz_dep_.CoverTab[124643]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:357
		// _ = "end of CoverTab[124643]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:357
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:357
	// _ = "end of CoverTab[124633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:357
	_go_fuzz_dep_.CoverTab[124634]++
										switch n.kind {
	case documentNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:359
		_go_fuzz_dep_.CoverTab[124644]++
											return d.document(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:360
		// _ = "end of CoverTab[124644]"
	case aliasNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:361
		_go_fuzz_dep_.CoverTab[124645]++
											return d.alias(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:362
		// _ = "end of CoverTab[124645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:362
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:362
		_go_fuzz_dep_.CoverTab[124646]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:362
		// _ = "end of CoverTab[124646]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:363
	// _ = "end of CoverTab[124634]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:363
	_go_fuzz_dep_.CoverTab[124635]++
										out, unmarshaled, good := d.prepare(n, out)
										if unmarshaled {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:365
		_go_fuzz_dep_.CoverTab[124647]++
											return good
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:366
		// _ = "end of CoverTab[124647]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:367
		_go_fuzz_dep_.CoverTab[124648]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:367
		// _ = "end of CoverTab[124648]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:367
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:367
	// _ = "end of CoverTab[124635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:367
	_go_fuzz_dep_.CoverTab[124636]++
										switch n.kind {
	case scalarNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:369
		_go_fuzz_dep_.CoverTab[124649]++
											good = d.scalar(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:370
		// _ = "end of CoverTab[124649]"
	case mappingNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:371
		_go_fuzz_dep_.CoverTab[124650]++
											good = d.mapping(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:372
		// _ = "end of CoverTab[124650]"
	case sequenceNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:373
		_go_fuzz_dep_.CoverTab[124651]++
											good = d.sequence(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:374
		// _ = "end of CoverTab[124651]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:375
		_go_fuzz_dep_.CoverTab[124652]++
											panic("internal error: unknown node kind: " + strconv.Itoa(n.kind))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:376
		// _ = "end of CoverTab[124652]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:377
	// _ = "end of CoverTab[124636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:377
	_go_fuzz_dep_.CoverTab[124637]++
										return good
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:378
	// _ = "end of CoverTab[124637]"
}

func (d *decoder) document(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:381
	_go_fuzz_dep_.CoverTab[124653]++
										if len(n.children) == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:382
		_go_fuzz_dep_.CoverTab[124655]++
											d.doc = n
											d.unmarshal(n.children[0], out)
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:385
		// _ = "end of CoverTab[124655]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:386
		_go_fuzz_dep_.CoverTab[124656]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:386
		// _ = "end of CoverTab[124656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:386
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:386
	// _ = "end of CoverTab[124653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:386
	_go_fuzz_dep_.CoverTab[124654]++
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:387
	// _ = "end of CoverTab[124654]"
}

func (d *decoder) alias(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:390
	_go_fuzz_dep_.CoverTab[124657]++
										if d.aliases[n] {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:391
		_go_fuzz_dep_.CoverTab[124659]++

											failf("anchor '%s' value contains itself", n.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:393
		// _ = "end of CoverTab[124659]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:394
		_go_fuzz_dep_.CoverTab[124660]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:394
		// _ = "end of CoverTab[124660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:394
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:394
	// _ = "end of CoverTab[124657]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:394
	_go_fuzz_dep_.CoverTab[124658]++
										d.aliases[n] = true
										d.aliasDepth++
										good = d.unmarshal(n.alias, out)
										d.aliasDepth--
										delete(d.aliases, n)
										return good
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:400
	// _ = "end of CoverTab[124658]"
}

var zeroValue reflect.Value

func resetMap(out reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:405
	_go_fuzz_dep_.CoverTab[124661]++
										for _, k := range out.MapKeys() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:406
		_go_fuzz_dep_.CoverTab[124662]++
											out.SetMapIndex(k, zeroValue)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:407
		// _ = "end of CoverTab[124662]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:408
	// _ = "end of CoverTab[124661]"
}

func (d *decoder) scalar(n *node, out reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:411
	_go_fuzz_dep_.CoverTab[124663]++
										var tag string
										var resolved interface{}
										if n.tag == "" && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:414
		_go_fuzz_dep_.CoverTab[124669]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:414
		return !n.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:414
		// _ = "end of CoverTab[124669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:414
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:414
		_go_fuzz_dep_.CoverTab[124670]++
											tag = yaml_STR_TAG
											resolved = n.value
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:416
		// _ = "end of CoverTab[124670]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:417
		_go_fuzz_dep_.CoverTab[124671]++
											tag, resolved = resolve(n.tag, n.value)
											if tag == yaml_BINARY_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:419
			_go_fuzz_dep_.CoverTab[124672]++
												data, err := base64.StdEncoding.DecodeString(resolved.(string))
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:421
				_go_fuzz_dep_.CoverTab[124674]++
													failf("!!binary value contains invalid base64 data")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:422
				// _ = "end of CoverTab[124674]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:423
				_go_fuzz_dep_.CoverTab[124675]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:423
				// _ = "end of CoverTab[124675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:423
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:423
			// _ = "end of CoverTab[124672]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:423
			_go_fuzz_dep_.CoverTab[124673]++
												resolved = string(data)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:424
			// _ = "end of CoverTab[124673]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:425
			_go_fuzz_dep_.CoverTab[124676]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:425
			// _ = "end of CoverTab[124676]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:425
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:425
		// _ = "end of CoverTab[124671]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:426
	// _ = "end of CoverTab[124663]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:426
	_go_fuzz_dep_.CoverTab[124664]++
										if resolved == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:427
		_go_fuzz_dep_.CoverTab[124677]++
											if out.Kind() == reflect.Map && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:428
			_go_fuzz_dep_.CoverTab[124679]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:428
			return !out.CanAddr()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:428
			// _ = "end of CoverTab[124679]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:428
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:428
			_go_fuzz_dep_.CoverTab[124680]++
												resetMap(out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:429
			// _ = "end of CoverTab[124680]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:430
			_go_fuzz_dep_.CoverTab[124681]++
												out.Set(reflect.Zero(out.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:431
			// _ = "end of CoverTab[124681]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:432
		// _ = "end of CoverTab[124677]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:432
		_go_fuzz_dep_.CoverTab[124678]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:433
		// _ = "end of CoverTab[124678]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:434
		_go_fuzz_dep_.CoverTab[124682]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:434
		// _ = "end of CoverTab[124682]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:434
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:434
	// _ = "end of CoverTab[124664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:434
	_go_fuzz_dep_.CoverTab[124665]++
										if resolvedv := reflect.ValueOf(resolved); out.Type() == resolvedv.Type() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:435
		_go_fuzz_dep_.CoverTab[124683]++

											out.Set(resolvedv)
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:438
		// _ = "end of CoverTab[124683]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:439
		_go_fuzz_dep_.CoverTab[124684]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:439
		// _ = "end of CoverTab[124684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:439
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:439
	// _ = "end of CoverTab[124665]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:439
	_go_fuzz_dep_.CoverTab[124666]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:442
	if out.CanAddr() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:442
		_go_fuzz_dep_.CoverTab[124685]++
											u, ok := out.Addr().Interface().(encoding.TextUnmarshaler)
											if ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:444
			_go_fuzz_dep_.CoverTab[124686]++
												var text []byte
												if tag == yaml_BINARY_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:446
				_go_fuzz_dep_.CoverTab[124689]++
													text = []byte(resolved.(string))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:447
				// _ = "end of CoverTab[124689]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:448
				_go_fuzz_dep_.CoverTab[124690]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:452
				text = []byte(n.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:452
				// _ = "end of CoverTab[124690]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:453
			// _ = "end of CoverTab[124686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:453
			_go_fuzz_dep_.CoverTab[124687]++
												err := u.UnmarshalText(text)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:455
				_go_fuzz_dep_.CoverTab[124691]++
													fail(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:456
				// _ = "end of CoverTab[124691]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:457
				_go_fuzz_dep_.CoverTab[124692]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:457
				// _ = "end of CoverTab[124692]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:457
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:457
			// _ = "end of CoverTab[124687]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:457
			_go_fuzz_dep_.CoverTab[124688]++
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:458
			// _ = "end of CoverTab[124688]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:459
			_go_fuzz_dep_.CoverTab[124693]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:459
			// _ = "end of CoverTab[124693]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:459
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:459
		// _ = "end of CoverTab[124685]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:460
		_go_fuzz_dep_.CoverTab[124694]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:460
		// _ = "end of CoverTab[124694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:460
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:460
	// _ = "end of CoverTab[124666]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:460
	_go_fuzz_dep_.CoverTab[124667]++
										switch out.Kind() {
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:462
		_go_fuzz_dep_.CoverTab[124695]++
											if tag == yaml_BINARY_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:463
			_go_fuzz_dep_.CoverTab[124706]++
												out.SetString(resolved.(string))
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:465
			// _ = "end of CoverTab[124706]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:466
			_go_fuzz_dep_.CoverTab[124707]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:466
			// _ = "end of CoverTab[124707]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:466
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:466
		// _ = "end of CoverTab[124695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:466
		_go_fuzz_dep_.CoverTab[124696]++
											if resolved != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:467
			_go_fuzz_dep_.CoverTab[124708]++
												out.SetString(n.value)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:469
			// _ = "end of CoverTab[124708]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:470
			_go_fuzz_dep_.CoverTab[124709]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:470
			// _ = "end of CoverTab[124709]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:470
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:470
		// _ = "end of CoverTab[124696]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:471
		_go_fuzz_dep_.CoverTab[124697]++
											if resolved == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:472
			_go_fuzz_dep_.CoverTab[124710]++
												out.Set(reflect.Zero(out.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:473
			// _ = "end of CoverTab[124710]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:474
			_go_fuzz_dep_.CoverTab[124711]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:474
			if tag == yaml_TIMESTAMP_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:474
				_go_fuzz_dep_.CoverTab[124712]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:480
				out.Set(reflect.ValueOf(n.value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:480
				// _ = "end of CoverTab[124712]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:481
				_go_fuzz_dep_.CoverTab[124713]++
													out.Set(reflect.ValueOf(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:482
				// _ = "end of CoverTab[124713]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:483
			// _ = "end of CoverTab[124711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:483
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:483
		// _ = "end of CoverTab[124697]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:483
		_go_fuzz_dep_.CoverTab[124698]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:484
		// _ = "end of CoverTab[124698]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:485
		_go_fuzz_dep_.CoverTab[124699]++
											switch resolved := resolved.(type) {
		case int:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:487
			_go_fuzz_dep_.CoverTab[124714]++
												if !out.OverflowInt(int64(resolved)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:488
				_go_fuzz_dep_.CoverTab[124719]++
													out.SetInt(int64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:490
				// _ = "end of CoverTab[124719]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:491
				_go_fuzz_dep_.CoverTab[124720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:491
				// _ = "end of CoverTab[124720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:491
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:491
			// _ = "end of CoverTab[124714]"
		case int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:492
			_go_fuzz_dep_.CoverTab[124715]++
												if !out.OverflowInt(resolved) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:493
				_go_fuzz_dep_.CoverTab[124721]++
													out.SetInt(resolved)
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:495
				// _ = "end of CoverTab[124721]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:496
				_go_fuzz_dep_.CoverTab[124722]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:496
				// _ = "end of CoverTab[124722]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:496
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:496
			// _ = "end of CoverTab[124715]"
		case uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:497
			_go_fuzz_dep_.CoverTab[124716]++
												if resolved <= math.MaxInt64 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:498
				_go_fuzz_dep_.CoverTab[124723]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:498
				return !out.OverflowInt(int64(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:498
				// _ = "end of CoverTab[124723]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:498
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:498
				_go_fuzz_dep_.CoverTab[124724]++
													out.SetInt(int64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:500
				// _ = "end of CoverTab[124724]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:501
				_go_fuzz_dep_.CoverTab[124725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:501
				// _ = "end of CoverTab[124725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:501
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:501
			// _ = "end of CoverTab[124716]"
		case float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:502
			_go_fuzz_dep_.CoverTab[124717]++
												if resolved <= math.MaxInt64 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:503
				_go_fuzz_dep_.CoverTab[124726]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:503
				return !out.OverflowInt(int64(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:503
				// _ = "end of CoverTab[124726]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:503
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:503
				_go_fuzz_dep_.CoverTab[124727]++
													out.SetInt(int64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:505
				// _ = "end of CoverTab[124727]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:506
				_go_fuzz_dep_.CoverTab[124728]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:506
				// _ = "end of CoverTab[124728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:506
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:506
			// _ = "end of CoverTab[124717]"
		case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:507
			_go_fuzz_dep_.CoverTab[124718]++
												if out.Type() == durationType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:508
				_go_fuzz_dep_.CoverTab[124729]++
													d, err := time.ParseDuration(resolved)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:510
					_go_fuzz_dep_.CoverTab[124730]++
														out.SetInt(int64(d))
														return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:512
					// _ = "end of CoverTab[124730]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:513
					_go_fuzz_dep_.CoverTab[124731]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:513
					// _ = "end of CoverTab[124731]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:513
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:513
				// _ = "end of CoverTab[124729]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:514
				_go_fuzz_dep_.CoverTab[124732]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:514
				// _ = "end of CoverTab[124732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:514
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:514
			// _ = "end of CoverTab[124718]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:515
		// _ = "end of CoverTab[124699]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:516
		_go_fuzz_dep_.CoverTab[124700]++
											switch resolved := resolved.(type) {
		case int:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:518
			_go_fuzz_dep_.CoverTab[124733]++
												if resolved >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:519
				_go_fuzz_dep_.CoverTab[124737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:519
				return !out.OverflowUint(uint64(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:519
				// _ = "end of CoverTab[124737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:519
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:519
				_go_fuzz_dep_.CoverTab[124738]++
													out.SetUint(uint64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:521
				// _ = "end of CoverTab[124738]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:522
				_go_fuzz_dep_.CoverTab[124739]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:522
				// _ = "end of CoverTab[124739]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:522
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:522
			// _ = "end of CoverTab[124733]"
		case int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:523
			_go_fuzz_dep_.CoverTab[124734]++
												if resolved >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:524
				_go_fuzz_dep_.CoverTab[124740]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:524
				return !out.OverflowUint(uint64(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:524
				// _ = "end of CoverTab[124740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:524
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:524
				_go_fuzz_dep_.CoverTab[124741]++
													out.SetUint(uint64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:526
				// _ = "end of CoverTab[124741]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:527
				_go_fuzz_dep_.CoverTab[124742]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:527
				// _ = "end of CoverTab[124742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:527
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:527
			// _ = "end of CoverTab[124734]"
		case uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:528
			_go_fuzz_dep_.CoverTab[124735]++
												if !out.OverflowUint(uint64(resolved)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:529
				_go_fuzz_dep_.CoverTab[124743]++
													out.SetUint(uint64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:531
				// _ = "end of CoverTab[124743]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:532
				_go_fuzz_dep_.CoverTab[124744]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:532
				// _ = "end of CoverTab[124744]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:532
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:532
			// _ = "end of CoverTab[124735]"
		case float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:533
			_go_fuzz_dep_.CoverTab[124736]++
												if resolved <= math.MaxUint64 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:534
				_go_fuzz_dep_.CoverTab[124745]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:534
				return !out.OverflowUint(uint64(resolved))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:534
				// _ = "end of CoverTab[124745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:534
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:534
				_go_fuzz_dep_.CoverTab[124746]++
													out.SetUint(uint64(resolved))
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:536
				// _ = "end of CoverTab[124746]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:537
				_go_fuzz_dep_.CoverTab[124747]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:537
				// _ = "end of CoverTab[124747]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:537
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:537
			// _ = "end of CoverTab[124736]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:538
		// _ = "end of CoverTab[124700]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:539
		_go_fuzz_dep_.CoverTab[124701]++
											switch resolved := resolved.(type) {
		case bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:541
			_go_fuzz_dep_.CoverTab[124748]++
												out.SetBool(resolved)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:543
			// _ = "end of CoverTab[124748]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:544
		// _ = "end of CoverTab[124701]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:545
		_go_fuzz_dep_.CoverTab[124702]++
											switch resolved := resolved.(type) {
		case int:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:547
			_go_fuzz_dep_.CoverTab[124749]++
												out.SetFloat(float64(resolved))
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:549
			// _ = "end of CoverTab[124749]"
		case int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:550
			_go_fuzz_dep_.CoverTab[124750]++
												out.SetFloat(float64(resolved))
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:552
			// _ = "end of CoverTab[124750]"
		case uint64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:553
			_go_fuzz_dep_.CoverTab[124751]++
												out.SetFloat(float64(resolved))
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:555
			// _ = "end of CoverTab[124751]"
		case float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:556
			_go_fuzz_dep_.CoverTab[124752]++
												out.SetFloat(resolved)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:558
			// _ = "end of CoverTab[124752]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:559
		// _ = "end of CoverTab[124702]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:560
		_go_fuzz_dep_.CoverTab[124703]++
											if resolvedv := reflect.ValueOf(resolved); out.Type() == resolvedv.Type() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:561
			_go_fuzz_dep_.CoverTab[124753]++
												out.Set(resolvedv)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:563
			// _ = "end of CoverTab[124753]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:564
			_go_fuzz_dep_.CoverTab[124754]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:564
			// _ = "end of CoverTab[124754]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:564
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:564
		// _ = "end of CoverTab[124703]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:565
		_go_fuzz_dep_.CoverTab[124704]++
											if out.Type().Elem() == reflect.TypeOf(resolved) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:566
			_go_fuzz_dep_.CoverTab[124755]++

												elem := reflect.New(out.Type().Elem())
												elem.Elem().Set(reflect.ValueOf(resolved))
												out.Set(elem)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:571
			// _ = "end of CoverTab[124755]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
			_go_fuzz_dep_.CoverTab[124756]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
			// _ = "end of CoverTab[124756]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
		// _ = "end of CoverTab[124704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
		_go_fuzz_dep_.CoverTab[124705]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:572
		// _ = "end of CoverTab[124705]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:573
	// _ = "end of CoverTab[124667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:573
	_go_fuzz_dep_.CoverTab[124668]++
										d.terror(n, tag, out)
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:575
	// _ = "end of CoverTab[124668]"
}

func settableValueOf(i interface{}) reflect.Value {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:578
	_go_fuzz_dep_.CoverTab[124757]++
										v := reflect.ValueOf(i)
										sv := reflect.New(v.Type()).Elem()
										sv.Set(v)
										return sv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:582
	// _ = "end of CoverTab[124757]"
}

func (d *decoder) sequence(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:585
	_go_fuzz_dep_.CoverTab[124758]++
										l := len(n.children)

										var iface reflect.Value
										switch out.Kind() {
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:590
		_go_fuzz_dep_.CoverTab[124763]++
											out.Set(reflect.MakeSlice(out.Type(), l, l))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:591
		// _ = "end of CoverTab[124763]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:592
		_go_fuzz_dep_.CoverTab[124764]++
											if l != out.Len() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:593
			_go_fuzz_dep_.CoverTab[124767]++
												failf("invalid array: want %d elements but got %d", out.Len(), l)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:594
			// _ = "end of CoverTab[124767]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:595
			_go_fuzz_dep_.CoverTab[124768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:595
			// _ = "end of CoverTab[124768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:595
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:595
		// _ = "end of CoverTab[124764]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:596
		_go_fuzz_dep_.CoverTab[124765]++

											iface = out
											out = settableValueOf(make([]interface{}, l))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:599
		// _ = "end of CoverTab[124765]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:600
		_go_fuzz_dep_.CoverTab[124766]++
											d.terror(n, yaml_SEQ_TAG, out)
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:602
		// _ = "end of CoverTab[124766]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:603
	// _ = "end of CoverTab[124758]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:603
	_go_fuzz_dep_.CoverTab[124759]++
										et := out.Type().Elem()

										j := 0
										for i := 0; i < l; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:607
		_go_fuzz_dep_.CoverTab[124769]++
											e := reflect.New(et).Elem()
											if ok := d.unmarshal(n.children[i], e); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:609
			_go_fuzz_dep_.CoverTab[124770]++
												out.Index(j).Set(e)
												j++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:611
			// _ = "end of CoverTab[124770]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:612
			_go_fuzz_dep_.CoverTab[124771]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:612
			// _ = "end of CoverTab[124771]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:612
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:612
		// _ = "end of CoverTab[124769]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:613
	// _ = "end of CoverTab[124759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:613
	_go_fuzz_dep_.CoverTab[124760]++
										if out.Kind() != reflect.Array {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:614
		_go_fuzz_dep_.CoverTab[124772]++
											out.Set(out.Slice(0, j))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:615
		// _ = "end of CoverTab[124772]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:616
		_go_fuzz_dep_.CoverTab[124773]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:616
		// _ = "end of CoverTab[124773]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:616
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:616
	// _ = "end of CoverTab[124760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:616
	_go_fuzz_dep_.CoverTab[124761]++
										if iface.IsValid() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:617
		_go_fuzz_dep_.CoverTab[124774]++
											iface.Set(out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:618
		// _ = "end of CoverTab[124774]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:619
		_go_fuzz_dep_.CoverTab[124775]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:619
		// _ = "end of CoverTab[124775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:619
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:619
	// _ = "end of CoverTab[124761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:619
	_go_fuzz_dep_.CoverTab[124762]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:620
	// _ = "end of CoverTab[124762]"
}

func (d *decoder) mapping(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:623
	_go_fuzz_dep_.CoverTab[124776]++
										switch out.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:625
		_go_fuzz_dep_.CoverTab[124781]++
											return d.mappingStruct(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:626
		// _ = "end of CoverTab[124781]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:627
		_go_fuzz_dep_.CoverTab[124782]++
											return d.mappingSlice(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:628
		// _ = "end of CoverTab[124782]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:629
		_go_fuzz_dep_.CoverTab[124783]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:629
		// _ = "end of CoverTab[124783]"

	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:631
		_go_fuzz_dep_.CoverTab[124784]++
											if d.mapType.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:632
			_go_fuzz_dep_.CoverTab[124786]++
												iface := out
												out = reflect.MakeMap(d.mapType)
												iface.Set(out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:635
			// _ = "end of CoverTab[124786]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:636
			_go_fuzz_dep_.CoverTab[124787]++
												slicev := reflect.New(d.mapType).Elem()
												if !d.mappingSlice(n, slicev) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:638
				_go_fuzz_dep_.CoverTab[124789]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:639
				// _ = "end of CoverTab[124789]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:640
				_go_fuzz_dep_.CoverTab[124790]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:640
				// _ = "end of CoverTab[124790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:640
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:640
			// _ = "end of CoverTab[124787]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:640
			_go_fuzz_dep_.CoverTab[124788]++
												out.Set(slicev)
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:642
			// _ = "end of CoverTab[124788]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:643
		// _ = "end of CoverTab[124784]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:644
		_go_fuzz_dep_.CoverTab[124785]++
											d.terror(n, yaml_MAP_TAG, out)
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:646
		// _ = "end of CoverTab[124785]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:647
	// _ = "end of CoverTab[124776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:647
	_go_fuzz_dep_.CoverTab[124777]++
										outt := out.Type()
										kt := outt.Key()
										et := outt.Elem()

										mapType := d.mapType
										if outt.Key() == ifaceType && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:653
		_go_fuzz_dep_.CoverTab[124791]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:653
		return outt.Elem() == ifaceType
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:653
		// _ = "end of CoverTab[124791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:653
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:653
		_go_fuzz_dep_.CoverTab[124792]++
											d.mapType = outt
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:654
		// _ = "end of CoverTab[124792]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:655
		_go_fuzz_dep_.CoverTab[124793]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:655
		// _ = "end of CoverTab[124793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:655
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:655
	// _ = "end of CoverTab[124777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:655
	_go_fuzz_dep_.CoverTab[124778]++

										if out.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:657
		_go_fuzz_dep_.CoverTab[124794]++
											out.Set(reflect.MakeMap(outt))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:658
		// _ = "end of CoverTab[124794]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:659
		_go_fuzz_dep_.CoverTab[124795]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:659
		// _ = "end of CoverTab[124795]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:659
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:659
	// _ = "end of CoverTab[124778]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:659
	_go_fuzz_dep_.CoverTab[124779]++
										l := len(n.children)
										for i := 0; i < l; i += 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:661
		_go_fuzz_dep_.CoverTab[124796]++
											if isMerge(n.children[i]) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:662
			_go_fuzz_dep_.CoverTab[124798]++
												d.merge(n.children[i+1], out)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:664
			// _ = "end of CoverTab[124798]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:665
			_go_fuzz_dep_.CoverTab[124799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:665
			// _ = "end of CoverTab[124799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:665
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:665
		// _ = "end of CoverTab[124796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:665
		_go_fuzz_dep_.CoverTab[124797]++
											k := reflect.New(kt).Elem()
											if d.unmarshal(n.children[i], k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:667
			_go_fuzz_dep_.CoverTab[124800]++
												kkind := k.Kind()
												if kkind == reflect.Interface {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:669
				_go_fuzz_dep_.CoverTab[124803]++
													kkind = k.Elem().Kind()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:670
				// _ = "end of CoverTab[124803]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:671
				_go_fuzz_dep_.CoverTab[124804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:671
				// _ = "end of CoverTab[124804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:671
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:671
			// _ = "end of CoverTab[124800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:671
			_go_fuzz_dep_.CoverTab[124801]++
												if kkind == reflect.Map || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:672
				_go_fuzz_dep_.CoverTab[124805]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:672
				return kkind == reflect.Slice
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:672
				// _ = "end of CoverTab[124805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:672
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:672
				_go_fuzz_dep_.CoverTab[124806]++
													failf("invalid map key: %#v", k.Interface())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:673
				// _ = "end of CoverTab[124806]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:674
				_go_fuzz_dep_.CoverTab[124807]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:674
				// _ = "end of CoverTab[124807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:674
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:674
			// _ = "end of CoverTab[124801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:674
			_go_fuzz_dep_.CoverTab[124802]++
												e := reflect.New(et).Elem()
												if d.unmarshal(n.children[i+1], e) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:676
				_go_fuzz_dep_.CoverTab[124808]++
													d.setMapIndex(n.children[i+1], out, k, e)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:677
				// _ = "end of CoverTab[124808]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:678
				_go_fuzz_dep_.CoverTab[124809]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:678
				// _ = "end of CoverTab[124809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:678
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:678
			// _ = "end of CoverTab[124802]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:679
			_go_fuzz_dep_.CoverTab[124810]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:679
			// _ = "end of CoverTab[124810]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:679
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:679
		// _ = "end of CoverTab[124797]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:680
	// _ = "end of CoverTab[124779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:680
	_go_fuzz_dep_.CoverTab[124780]++
										d.mapType = mapType
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:682
	// _ = "end of CoverTab[124780]"
}

func (d *decoder) setMapIndex(n *node, out, k, v reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:685
	_go_fuzz_dep_.CoverTab[124811]++
										if d.strict && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:686
		_go_fuzz_dep_.CoverTab[124813]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:686
		return out.MapIndex(k) != zeroValue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:686
		// _ = "end of CoverTab[124813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:686
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:686
		_go_fuzz_dep_.CoverTab[124814]++
											d.terrors = append(d.terrors, fmt.Sprintf("line %d: key %#v already set in map", n.line+1, k.Interface()))
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:688
		// _ = "end of CoverTab[124814]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:689
		_go_fuzz_dep_.CoverTab[124815]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:689
		// _ = "end of CoverTab[124815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:689
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:689
	// _ = "end of CoverTab[124811]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:689
	_go_fuzz_dep_.CoverTab[124812]++
										out.SetMapIndex(k, v)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:690
	// _ = "end of CoverTab[124812]"
}

func (d *decoder) mappingSlice(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:693
	_go_fuzz_dep_.CoverTab[124816]++
										outt := out.Type()
										if outt.Elem() != mapItemType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:695
		_go_fuzz_dep_.CoverTab[124819]++
											d.terror(n, yaml_MAP_TAG, out)
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:697
		// _ = "end of CoverTab[124819]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:698
		_go_fuzz_dep_.CoverTab[124820]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:698
		// _ = "end of CoverTab[124820]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:698
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:698
	// _ = "end of CoverTab[124816]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:698
	_go_fuzz_dep_.CoverTab[124817]++

										mapType := d.mapType
										d.mapType = outt

										var slice []MapItem
										var l = len(n.children)
										for i := 0; i < l; i += 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:705
		_go_fuzz_dep_.CoverTab[124821]++
											if isMerge(n.children[i]) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:706
			_go_fuzz_dep_.CoverTab[124823]++
												d.merge(n.children[i+1], out)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:708
			// _ = "end of CoverTab[124823]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:709
			_go_fuzz_dep_.CoverTab[124824]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:709
			// _ = "end of CoverTab[124824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:709
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:709
		// _ = "end of CoverTab[124821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:709
		_go_fuzz_dep_.CoverTab[124822]++
											item := MapItem{}
											k := reflect.ValueOf(&item.Key).Elem()
											if d.unmarshal(n.children[i], k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:712
			_go_fuzz_dep_.CoverTab[124825]++
												v := reflect.ValueOf(&item.Value).Elem()
												if d.unmarshal(n.children[i+1], v) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:714
				_go_fuzz_dep_.CoverTab[124826]++
													slice = append(slice, item)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:715
				// _ = "end of CoverTab[124826]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:716
				_go_fuzz_dep_.CoverTab[124827]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:716
				// _ = "end of CoverTab[124827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:716
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:716
			// _ = "end of CoverTab[124825]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:717
			_go_fuzz_dep_.CoverTab[124828]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:717
			// _ = "end of CoverTab[124828]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:717
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:717
		// _ = "end of CoverTab[124822]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:718
	// _ = "end of CoverTab[124817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:718
	_go_fuzz_dep_.CoverTab[124818]++
										out.Set(reflect.ValueOf(slice))
										d.mapType = mapType
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:721
	// _ = "end of CoverTab[124818]"
}

func (d *decoder) mappingStruct(n *node, out reflect.Value) (good bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:724
	_go_fuzz_dep_.CoverTab[124829]++
										sinfo, err := getStructInfo(out.Type())
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:726
		_go_fuzz_dep_.CoverTab[124834]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:727
		// _ = "end of CoverTab[124834]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:728
		_go_fuzz_dep_.CoverTab[124835]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:728
		// _ = "end of CoverTab[124835]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:728
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:728
	// _ = "end of CoverTab[124829]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:728
	_go_fuzz_dep_.CoverTab[124830]++
										name := settableValueOf("")
										l := len(n.children)

										var inlineMap reflect.Value
										var elemType reflect.Type
										if sinfo.InlineMap != -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:734
		_go_fuzz_dep_.CoverTab[124836]++
											inlineMap = out.Field(sinfo.InlineMap)
											inlineMap.Set(reflect.New(inlineMap.Type()).Elem())
											elemType = inlineMap.Type().Elem()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:737
		// _ = "end of CoverTab[124836]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:738
		_go_fuzz_dep_.CoverTab[124837]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:738
		// _ = "end of CoverTab[124837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:738
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:738
	// _ = "end of CoverTab[124830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:738
	_go_fuzz_dep_.CoverTab[124831]++

										var doneFields []bool
										if d.strict {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:741
		_go_fuzz_dep_.CoverTab[124838]++
											doneFields = make([]bool, len(sinfo.FieldsList))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:742
		// _ = "end of CoverTab[124838]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:743
		_go_fuzz_dep_.CoverTab[124839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:743
		// _ = "end of CoverTab[124839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:743
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:743
	// _ = "end of CoverTab[124831]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:743
	_go_fuzz_dep_.CoverTab[124832]++
										for i := 0; i < l; i += 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:744
		_go_fuzz_dep_.CoverTab[124840]++
											ni := n.children[i]
											if isMerge(ni) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:746
			_go_fuzz_dep_.CoverTab[124843]++
												d.merge(n.children[i+1], out)
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:748
			// _ = "end of CoverTab[124843]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:749
			_go_fuzz_dep_.CoverTab[124844]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:749
			// _ = "end of CoverTab[124844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:749
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:749
		// _ = "end of CoverTab[124840]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:749
		_go_fuzz_dep_.CoverTab[124841]++
											if !d.unmarshal(ni, name) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:750
			_go_fuzz_dep_.CoverTab[124845]++
												continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:751
			// _ = "end of CoverTab[124845]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:752
			_go_fuzz_dep_.CoverTab[124846]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:752
			// _ = "end of CoverTab[124846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:752
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:752
		// _ = "end of CoverTab[124841]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:752
		_go_fuzz_dep_.CoverTab[124842]++
											if info, ok := sinfo.FieldsMap[name.String()]; ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:753
			_go_fuzz_dep_.CoverTab[124847]++
												if d.strict {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:754
				_go_fuzz_dep_.CoverTab[124850]++
													if doneFields[info.Id] {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:755
					_go_fuzz_dep_.CoverTab[124852]++
														d.terrors = append(d.terrors, fmt.Sprintf("line %d: field %s already set in type %s", ni.line+1, name.String(), out.Type()))
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:757
					// _ = "end of CoverTab[124852]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:758
					_go_fuzz_dep_.CoverTab[124853]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:758
					// _ = "end of CoverTab[124853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:758
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:758
				// _ = "end of CoverTab[124850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:758
				_go_fuzz_dep_.CoverTab[124851]++
													doneFields[info.Id] = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:759
				// _ = "end of CoverTab[124851]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:760
				_go_fuzz_dep_.CoverTab[124854]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:760
				// _ = "end of CoverTab[124854]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:760
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:760
			// _ = "end of CoverTab[124847]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:760
			_go_fuzz_dep_.CoverTab[124848]++
												var field reflect.Value
												if info.Inline == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:762
				_go_fuzz_dep_.CoverTab[124855]++
													field = out.Field(info.Num)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:763
				// _ = "end of CoverTab[124855]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:764
				_go_fuzz_dep_.CoverTab[124856]++
													field = out.FieldByIndex(info.Inline)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:765
				// _ = "end of CoverTab[124856]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:766
			// _ = "end of CoverTab[124848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:766
			_go_fuzz_dep_.CoverTab[124849]++
												d.unmarshal(n.children[i+1], field)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:767
			// _ = "end of CoverTab[124849]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:768
			_go_fuzz_dep_.CoverTab[124857]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:768
			if sinfo.InlineMap != -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:768
				_go_fuzz_dep_.CoverTab[124858]++
													if inlineMap.IsNil() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:769
					_go_fuzz_dep_.CoverTab[124860]++
														inlineMap.Set(reflect.MakeMap(inlineMap.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:770
					// _ = "end of CoverTab[124860]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:771
					_go_fuzz_dep_.CoverTab[124861]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:771
					// _ = "end of CoverTab[124861]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:771
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:771
				// _ = "end of CoverTab[124858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:771
				_go_fuzz_dep_.CoverTab[124859]++
													value := reflect.New(elemType).Elem()
													d.unmarshal(n.children[i+1], value)
													d.setMapIndex(n.children[i+1], inlineMap, name, value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:774
				// _ = "end of CoverTab[124859]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:775
				_go_fuzz_dep_.CoverTab[124862]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:775
				if d.strict {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:775
					_go_fuzz_dep_.CoverTab[124863]++
														d.terrors = append(d.terrors, fmt.Sprintf("line %d: field %s not found in type %s", ni.line+1, name.String(), out.Type()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:776
					// _ = "end of CoverTab[124863]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
					_go_fuzz_dep_.CoverTab[124864]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
					// _ = "end of CoverTab[124864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
				// _ = "end of CoverTab[124862]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
			// _ = "end of CoverTab[124857]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:777
		// _ = "end of CoverTab[124842]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:778
	// _ = "end of CoverTab[124832]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:778
	_go_fuzz_dep_.CoverTab[124833]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:779
	// _ = "end of CoverTab[124833]"
}

func failWantMap() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:782
	_go_fuzz_dep_.CoverTab[124865]++
										failf("map merge requires map or sequence of maps as the value")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:783
	// _ = "end of CoverTab[124865]"
}

func (d *decoder) merge(n *node, out reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:786
	_go_fuzz_dep_.CoverTab[124866]++
										switch n.kind {
	case mappingNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:788
		_go_fuzz_dep_.CoverTab[124867]++
											d.unmarshal(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:789
		// _ = "end of CoverTab[124867]"
	case aliasNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:790
		_go_fuzz_dep_.CoverTab[124868]++
											if n.alias != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:791
			_go_fuzz_dep_.CoverTab[124872]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:791
			return n.alias.kind != mappingNode
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:791
			// _ = "end of CoverTab[124872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:791
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:791
			_go_fuzz_dep_.CoverTab[124873]++
												failWantMap()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:792
			// _ = "end of CoverTab[124873]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:793
			_go_fuzz_dep_.CoverTab[124874]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:793
			// _ = "end of CoverTab[124874]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:793
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:793
		// _ = "end of CoverTab[124868]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:793
		_go_fuzz_dep_.CoverTab[124869]++
											d.unmarshal(n, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:794
		// _ = "end of CoverTab[124869]"
	case sequenceNode:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:795
		_go_fuzz_dep_.CoverTab[124870]++

											for i := len(n.children) - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:797
			_go_fuzz_dep_.CoverTab[124875]++
												ni := n.children[i]
												if ni.kind == aliasNode {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:799
				_go_fuzz_dep_.CoverTab[124877]++
													if ni.alias != nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:800
					_go_fuzz_dep_.CoverTab[124878]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:800
					return ni.alias.kind != mappingNode
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:800
					// _ = "end of CoverTab[124878]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:800
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:800
					_go_fuzz_dep_.CoverTab[124879]++
														failWantMap()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:801
					// _ = "end of CoverTab[124879]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:802
					_go_fuzz_dep_.CoverTab[124880]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:802
					// _ = "end of CoverTab[124880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:802
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:802
				// _ = "end of CoverTab[124877]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:803
				_go_fuzz_dep_.CoverTab[124881]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:803
				if ni.kind != mappingNode {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:803
					_go_fuzz_dep_.CoverTab[124882]++
														failWantMap()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:804
					// _ = "end of CoverTab[124882]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
					_go_fuzz_dep_.CoverTab[124883]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
					// _ = "end of CoverTab[124883]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
				// _ = "end of CoverTab[124881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
			// _ = "end of CoverTab[124875]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:805
			_go_fuzz_dep_.CoverTab[124876]++
												d.unmarshal(ni, out)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:806
			// _ = "end of CoverTab[124876]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:807
		// _ = "end of CoverTab[124870]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:808
		_go_fuzz_dep_.CoverTab[124871]++
											failWantMap()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:809
		// _ = "end of CoverTab[124871]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:810
	// _ = "end of CoverTab[124866]"
}

func isMerge(n *node) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:813
	_go_fuzz_dep_.CoverTab[124884]++
										return n.kind == scalarNode && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		_go_fuzz_dep_.CoverTab[124885]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		return n.value == "<<"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		// _ = "end of CoverTab[124885]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		_go_fuzz_dep_.CoverTab[124886]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		return (n.implicit == true || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
			_go_fuzz_dep_.CoverTab[124887]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
			return n.tag == yaml_MERGE_TAG
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
			// _ = "end of CoverTab[124887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
		// _ = "end of CoverTab[124886]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:814
	// _ = "end of CoverTab[124884]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:815
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/decode.go:815
var _ = _go_fuzz_dep_.CoverTab
