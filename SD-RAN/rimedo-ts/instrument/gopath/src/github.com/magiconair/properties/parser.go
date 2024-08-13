// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:5
)

import (
	"fmt"
	"runtime"
)

type parser struct {
	lex *lexer
}

func parse(input string) (properties *Properties, err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:16
	_go_fuzz_dep_.CoverTab[115820]++
											p := &parser{lex: lex(input)}
											defer p.recover(&err)

											properties = NewProperties()
											key := ""
											comments := []string{}

											for {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:24
		_go_fuzz_dep_.CoverTab[115822]++
												token := p.expectOneOf(itemComment, itemKey, itemEOF)
												switch token.typ {
		case itemEOF:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:27
			_go_fuzz_dep_.CoverTab[115825]++
													goto done
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:28
			// _ = "end of CoverTab[115825]"
		case itemComment:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:29
			_go_fuzz_dep_.CoverTab[115826]++
													comments = append(comments, token.val)
													continue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:31
			// _ = "end of CoverTab[115826]"
		case itemKey:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:32
			_go_fuzz_dep_.CoverTab[115827]++
													key = token.val
													if _, ok := properties.m[key]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:34
				_go_fuzz_dep_.CoverTab[115829]++
														properties.k = append(properties.k, key)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:35
				// _ = "end of CoverTab[115829]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
				_go_fuzz_dep_.CoverTab[115830]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
				// _ = "end of CoverTab[115830]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
			// _ = "end of CoverTab[115827]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
			_go_fuzz_dep_.CoverTab[115828]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:36
			// _ = "end of CoverTab[115828]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:37
		// _ = "end of CoverTab[115822]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:37
		_go_fuzz_dep_.CoverTab[115823]++

												token = p.expectOneOf(itemValue, itemEOF)
												if len(comments) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:40
			_go_fuzz_dep_.CoverTab[115831]++
													properties.c[key] = comments
													comments = []string{}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:42
			// _ = "end of CoverTab[115831]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:43
			_go_fuzz_dep_.CoverTab[115832]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:43
			// _ = "end of CoverTab[115832]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:43
		// _ = "end of CoverTab[115823]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:43
		_go_fuzz_dep_.CoverTab[115824]++
												switch token.typ {
		case itemEOF:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:45
			_go_fuzz_dep_.CoverTab[115833]++
													properties.m[key] = ""
													goto done
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:47
			// _ = "end of CoverTab[115833]"
		case itemValue:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:48
			_go_fuzz_dep_.CoverTab[115834]++
													properties.m[key] = token.val
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:49
			// _ = "end of CoverTab[115834]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:49
		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:49
			_go_fuzz_dep_.CoverTab[115835]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:49
			// _ = "end of CoverTab[115835]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:50
		// _ = "end of CoverTab[115824]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:51
	// _ = "end of CoverTab[115820]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:51
	_go_fuzz_dep_.CoverTab[115821]++

done:
											return properties, nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:54
	// _ = "end of CoverTab[115821]"
}

func (p *parser) errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:57
	_go_fuzz_dep_.CoverTab[115836]++
											format = fmt.Sprintf("properties: Line %d: %s", p.lex.lineNumber(), format)
											panic(fmt.Errorf(format, args...))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:59
	// _ = "end of CoverTab[115836]"
}

func (p *parser) expect(expected itemType) (token item) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:62
	_go_fuzz_dep_.CoverTab[115837]++
											token = p.lex.nextItem()
											if token.typ != expected {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:64
		_go_fuzz_dep_.CoverTab[115839]++
												p.unexpected(token)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:65
		// _ = "end of CoverTab[115839]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:66
		_go_fuzz_dep_.CoverTab[115840]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:66
		// _ = "end of CoverTab[115840]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:66
	// _ = "end of CoverTab[115837]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:66
	_go_fuzz_dep_.CoverTab[115838]++
											return token
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:67
	// _ = "end of CoverTab[115838]"
}

func (p *parser) expectOneOf(expected ...itemType) (token item) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:70
	_go_fuzz_dep_.CoverTab[115841]++
											token = p.lex.nextItem()
											for _, v := range expected {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:72
		_go_fuzz_dep_.CoverTab[115843]++
												if token.typ == v {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:73
			_go_fuzz_dep_.CoverTab[115844]++
													return token
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:74
			// _ = "end of CoverTab[115844]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:75
			_go_fuzz_dep_.CoverTab[115845]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:75
			// _ = "end of CoverTab[115845]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:75
		// _ = "end of CoverTab[115843]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:76
	// _ = "end of CoverTab[115841]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:76
	_go_fuzz_dep_.CoverTab[115842]++
											p.unexpected(token)
											panic("unexpected token")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:78
	// _ = "end of CoverTab[115842]"
}

func (p *parser) unexpected(token item) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:81
	_go_fuzz_dep_.CoverTab[115846]++
											p.errorf(token.String())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:82
	// _ = "end of CoverTab[115846]"
}

// recover is the handler that turns panics into returns from the top level of Parse.
func (p *parser) recover(errp *error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:86
	_go_fuzz_dep_.CoverTab[115847]++
											e := recover()
											if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:88
		_go_fuzz_dep_.CoverTab[115849]++
												if _, ok := e.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:89
			_go_fuzz_dep_.CoverTab[115851]++
													panic(e)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:90
			// _ = "end of CoverTab[115851]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:91
			_go_fuzz_dep_.CoverTab[115852]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:91
			// _ = "end of CoverTab[115852]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:91
		// _ = "end of CoverTab[115849]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:91
		_go_fuzz_dep_.CoverTab[115850]++
												*errp = e.(error)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:92
		// _ = "end of CoverTab[115850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:93
		_go_fuzz_dep_.CoverTab[115853]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:93
		// _ = "end of CoverTab[115853]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:93
	// _ = "end of CoverTab[115847]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:93
	_go_fuzz_dep_.CoverTab[115848]++
											return
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:94
	// _ = "end of CoverTab[115848]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/parser.go:95
var _ = _go_fuzz_dep_.CoverTab
