// TOML Parser.

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:3
)

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type tomlParser struct {
	flowIdx		int
	flow		[]token
	tree		*Tree
	currentTable	[]string
	seenTableKeys	[]string
}

type tomlParserStateFn func() tomlParserStateFn

// Formats and panics an error message based on a token
func (p *tomlParser) raiseError(tok *token, msg string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:26
	_go_fuzz_dep_.CoverTab[123618]++
											panic(tok.Position.String() + ": " + fmt.Sprintf(msg, args...))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:27
	// _ = "end of CoverTab[123618]"
}

func (p *tomlParser) run() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:30
	_go_fuzz_dep_.CoverTab[123619]++
											for state := p.parseStart; state != nil; {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:31
		_go_fuzz_dep_.CoverTab[123620]++
												state = state()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:32
		// _ = "end of CoverTab[123620]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:33
	// _ = "end of CoverTab[123619]"
}

func (p *tomlParser) peek() *token {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:36
	_go_fuzz_dep_.CoverTab[123621]++
											if p.flowIdx >= len(p.flow) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:37
		_go_fuzz_dep_.CoverTab[123623]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:38
		// _ = "end of CoverTab[123623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:39
		_go_fuzz_dep_.CoverTab[123624]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:39
		// _ = "end of CoverTab[123624]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:39
	// _ = "end of CoverTab[123621]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:39
	_go_fuzz_dep_.CoverTab[123622]++
											return &p.flow[p.flowIdx]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:40
	// _ = "end of CoverTab[123622]"
}

func (p *tomlParser) assume(typ tokenType) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:43
	_go_fuzz_dep_.CoverTab[123625]++
											tok := p.getToken()
											if tok == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:45
		_go_fuzz_dep_.CoverTab[123627]++
												p.raiseError(tok, "was expecting token %s, but token stream is empty", tok)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:46
		// _ = "end of CoverTab[123627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:47
		_go_fuzz_dep_.CoverTab[123628]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:47
		// _ = "end of CoverTab[123628]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:47
	// _ = "end of CoverTab[123625]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:47
	_go_fuzz_dep_.CoverTab[123626]++
											if tok.typ != typ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:48
		_go_fuzz_dep_.CoverTab[123629]++
												p.raiseError(tok, "was expecting token %s, but got %s instead", typ, tok)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:49
		// _ = "end of CoverTab[123629]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:50
		_go_fuzz_dep_.CoverTab[123630]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:50
		// _ = "end of CoverTab[123630]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:50
	// _ = "end of CoverTab[123626]"
}

func (p *tomlParser) getToken() *token {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:53
	_go_fuzz_dep_.CoverTab[123631]++
											tok := p.peek()
											if tok == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:55
		_go_fuzz_dep_.CoverTab[123633]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:56
		// _ = "end of CoverTab[123633]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:57
		_go_fuzz_dep_.CoverTab[123634]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:57
		// _ = "end of CoverTab[123634]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:57
	// _ = "end of CoverTab[123631]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:57
	_go_fuzz_dep_.CoverTab[123632]++
											p.flowIdx++
											return tok
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:59
	// _ = "end of CoverTab[123632]"
}

func (p *tomlParser) parseStart() tomlParserStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:62
	_go_fuzz_dep_.CoverTab[123635]++
											tok := p.peek()

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:66
	if tok == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:66
		_go_fuzz_dep_.CoverTab[123638]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:67
		// _ = "end of CoverTab[123638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:68
		_go_fuzz_dep_.CoverTab[123639]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:68
		// _ = "end of CoverTab[123639]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:68
	// _ = "end of CoverTab[123635]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:68
	_go_fuzz_dep_.CoverTab[123636]++

											switch tok.typ {
	case tokenDoubleLeftBracket:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:71
		_go_fuzz_dep_.CoverTab[123640]++
												return p.parseGroupArray
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:72
		// _ = "end of CoverTab[123640]"
	case tokenLeftBracket:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:73
		_go_fuzz_dep_.CoverTab[123641]++
												return p.parseGroup
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:74
		// _ = "end of CoverTab[123641]"
	case tokenKey:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:75
		_go_fuzz_dep_.CoverTab[123642]++
												return p.parseAssign
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:76
		// _ = "end of CoverTab[123642]"
	case tokenEOF:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:77
		_go_fuzz_dep_.CoverTab[123643]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:78
		// _ = "end of CoverTab[123643]"
	case tokenError:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:79
		_go_fuzz_dep_.CoverTab[123644]++
												p.raiseError(tok, "parsing error: %s", tok.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:80
		// _ = "end of CoverTab[123644]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:81
		_go_fuzz_dep_.CoverTab[123645]++
												p.raiseError(tok, "unexpected token %s", tok.typ)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:82
		// _ = "end of CoverTab[123645]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:83
	// _ = "end of CoverTab[123636]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:83
	_go_fuzz_dep_.CoverTab[123637]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:84
	// _ = "end of CoverTab[123637]"
}

func (p *tomlParser) parseGroupArray() tomlParserStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:87
	_go_fuzz_dep_.CoverTab[123646]++
											startToken := p.getToken()
											key := p.getToken()
											if key.typ != tokenKeyGroupArray {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:90
		_go_fuzz_dep_.CoverTab[123652]++
												p.raiseError(key, "unexpected token %s, was expecting a table array key", key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:91
		// _ = "end of CoverTab[123652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:92
		_go_fuzz_dep_.CoverTab[123653]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:92
		// _ = "end of CoverTab[123653]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:92
	// _ = "end of CoverTab[123646]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:92
	_go_fuzz_dep_.CoverTab[123647]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:95
	keys, err := parseKey(key.val)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:96
		_go_fuzz_dep_.CoverTab[123654]++
												p.raiseError(key, "invalid table array key: %s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:97
		// _ = "end of CoverTab[123654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:98
		_go_fuzz_dep_.CoverTab[123655]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:98
		// _ = "end of CoverTab[123655]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:98
	// _ = "end of CoverTab[123647]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:98
	_go_fuzz_dep_.CoverTab[123648]++
											p.tree.createSubTree(keys[:len(keys)-1], startToken.Position)
											destTree := p.tree.GetPath(keys)
											var array []*Tree
											if destTree == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:102
		_go_fuzz_dep_.CoverTab[123656]++
												array = make([]*Tree, 0)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:103
		// _ = "end of CoverTab[123656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
		_go_fuzz_dep_.CoverTab[123657]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
		if target, ok := destTree.([]*Tree); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
			_go_fuzz_dep_.CoverTab[123658]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
			return target != nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
			// _ = "end of CoverTab[123658]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:104
			_go_fuzz_dep_.CoverTab[123659]++
													array = destTree.([]*Tree)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:105
			// _ = "end of CoverTab[123659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:106
			_go_fuzz_dep_.CoverTab[123660]++
													p.raiseError(key, "key %s is already assigned and not of type table array", key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:107
			// _ = "end of CoverTab[123660]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:108
		// _ = "end of CoverTab[123657]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:108
	// _ = "end of CoverTab[123648]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:108
	_go_fuzz_dep_.CoverTab[123649]++
											p.currentTable = keys

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:112
	newTree := newTree()
											newTree.position = startToken.Position
											array = append(array, newTree)
											p.tree.SetPath(p.currentTable, array)

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:118
	prefix := key.val + "."
	found := false
	for ii := 0; ii < len(p.seenTableKeys); {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:120
		_go_fuzz_dep_.CoverTab[123661]++
												tableKey := p.seenTableKeys[ii]
												if strings.HasPrefix(tableKey, prefix) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:122
			_go_fuzz_dep_.CoverTab[123662]++
													p.seenTableKeys = append(p.seenTableKeys[:ii], p.seenTableKeys[ii+1:]...)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:123
			// _ = "end of CoverTab[123662]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:124
			_go_fuzz_dep_.CoverTab[123663]++
													found = (tableKey == key.val)
													ii++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:126
			// _ = "end of CoverTab[123663]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:127
		// _ = "end of CoverTab[123661]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:128
	// _ = "end of CoverTab[123649]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:128
	_go_fuzz_dep_.CoverTab[123650]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:131
	if !found {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:131
		_go_fuzz_dep_.CoverTab[123664]++
												p.seenTableKeys = append(p.seenTableKeys, key.val)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:132
		// _ = "end of CoverTab[123664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:133
		_go_fuzz_dep_.CoverTab[123665]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:133
		// _ = "end of CoverTab[123665]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:133
	// _ = "end of CoverTab[123650]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:133
	_go_fuzz_dep_.CoverTab[123651]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:136
	p.assume(tokenDoubleRightBracket)
											return p.parseStart
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:137
	// _ = "end of CoverTab[123651]"
}

func (p *tomlParser) parseGroup() tomlParserStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:140
	_go_fuzz_dep_.CoverTab[123666]++
											startToken := p.getToken()
											key := p.getToken()
											if key.typ != tokenKeyGroup {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:143
		_go_fuzz_dep_.CoverTab[123672]++
												p.raiseError(key, "unexpected token %s, was expecting a table key", key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:144
		// _ = "end of CoverTab[123672]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:145
		_go_fuzz_dep_.CoverTab[123673]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:145
		// _ = "end of CoverTab[123673]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:145
	// _ = "end of CoverTab[123666]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:145
	_go_fuzz_dep_.CoverTab[123667]++
											for _, item := range p.seenTableKeys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:146
		_go_fuzz_dep_.CoverTab[123674]++
												if item == key.val {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:147
			_go_fuzz_dep_.CoverTab[123675]++
													p.raiseError(key, "duplicated tables")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:148
			// _ = "end of CoverTab[123675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:149
			_go_fuzz_dep_.CoverTab[123676]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:149
			// _ = "end of CoverTab[123676]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:149
		// _ = "end of CoverTab[123674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:150
	// _ = "end of CoverTab[123667]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:150
	_go_fuzz_dep_.CoverTab[123668]++

											p.seenTableKeys = append(p.seenTableKeys, key.val)
											keys, err := parseKey(key.val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:154
		_go_fuzz_dep_.CoverTab[123677]++
												p.raiseError(key, "invalid table array key: %s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:155
		// _ = "end of CoverTab[123677]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:156
		_go_fuzz_dep_.CoverTab[123678]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:156
		// _ = "end of CoverTab[123678]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:156
	// _ = "end of CoverTab[123668]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:156
	_go_fuzz_dep_.CoverTab[123669]++
											if err := p.tree.createSubTree(keys, startToken.Position); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:157
		_go_fuzz_dep_.CoverTab[123679]++
												p.raiseError(key, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:158
		// _ = "end of CoverTab[123679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:159
		_go_fuzz_dep_.CoverTab[123680]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:159
		// _ = "end of CoverTab[123680]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:159
	// _ = "end of CoverTab[123669]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:159
	_go_fuzz_dep_.CoverTab[123670]++
											destTree := p.tree.GetPath(keys)
											if target, ok := destTree.(*Tree); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		_go_fuzz_dep_.CoverTab[123681]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		return target != nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		// _ = "end of CoverTab[123681]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		_go_fuzz_dep_.CoverTab[123682]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		return target.inline
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		// _ = "end of CoverTab[123682]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:161
		_go_fuzz_dep_.CoverTab[123683]++
												p.raiseError(key, "could not re-define exist inline table or its sub-table : %s",
			strings.Join(keys, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:163
		// _ = "end of CoverTab[123683]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:164
		_go_fuzz_dep_.CoverTab[123684]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:164
		// _ = "end of CoverTab[123684]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:164
	// _ = "end of CoverTab[123670]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:164
	_go_fuzz_dep_.CoverTab[123671]++
											p.assume(tokenRightBracket)
											p.currentTable = keys
											return p.parseStart
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:167
	// _ = "end of CoverTab[123671]"
}

func (p *tomlParser) parseAssign() tomlParserStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:170
	_go_fuzz_dep_.CoverTab[123685]++
											key := p.getToken()
											p.assume(tokenEqual)

											parsedKey, err := parseKey(key.val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:175
		_go_fuzz_dep_.CoverTab[123692]++
												p.raiseError(key, "invalid key: %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:176
		// _ = "end of CoverTab[123692]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:177
		_go_fuzz_dep_.CoverTab[123693]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:177
		// _ = "end of CoverTab[123693]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:177
	// _ = "end of CoverTab[123685]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:177
	_go_fuzz_dep_.CoverTab[123686]++

											value := p.parseRvalue()
											var tableKey []string
											if len(p.currentTable) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:181
		_go_fuzz_dep_.CoverTab[123694]++
												tableKey = p.currentTable
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:182
		// _ = "end of CoverTab[123694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:183
		_go_fuzz_dep_.CoverTab[123695]++
												tableKey = []string{}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:184
		// _ = "end of CoverTab[123695]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:185
	// _ = "end of CoverTab[123686]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:185
	_go_fuzz_dep_.CoverTab[123687]++

											prefixKey := parsedKey[0 : len(parsedKey)-1]
											tableKey = append(tableKey, prefixKey...)

	// find the table to assign, looking out for arrays of tables
	var targetNode *Tree
	switch node := p.tree.GetPath(tableKey).(type) {
	case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:193
		_go_fuzz_dep_.CoverTab[123696]++
												targetNode = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:194
		// _ = "end of CoverTab[123696]"
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:195
		_go_fuzz_dep_.CoverTab[123697]++
												targetNode = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:196
		// _ = "end of CoverTab[123697]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:197
		_go_fuzz_dep_.CoverTab[123698]++

												if err := p.tree.createSubTree(tableKey, key.Position); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:199
			_go_fuzz_dep_.CoverTab[123701]++
													p.raiseError(key, "could not create intermediate group: %s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:200
			// _ = "end of CoverTab[123701]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:201
			_go_fuzz_dep_.CoverTab[123702]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:201
			// _ = "end of CoverTab[123702]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:201
		// _ = "end of CoverTab[123698]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:201
		_go_fuzz_dep_.CoverTab[123699]++
												targetNode = p.tree.GetPath(tableKey).(*Tree)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:202
		// _ = "end of CoverTab[123699]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:203
		_go_fuzz_dep_.CoverTab[123700]++
												p.raiseError(key, "Unknown table type for path: %s",
			strings.Join(tableKey, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:205
		// _ = "end of CoverTab[123700]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:206
	// _ = "end of CoverTab[123687]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:206
	_go_fuzz_dep_.CoverTab[123688]++

											if targetNode.inline {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:208
		_go_fuzz_dep_.CoverTab[123703]++
												p.raiseError(key, "could not add key or sub-table to exist inline table or its sub-table : %s",
			strings.Join(tableKey, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:210
		// _ = "end of CoverTab[123703]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:211
		_go_fuzz_dep_.CoverTab[123704]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:211
		// _ = "end of CoverTab[123704]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:211
	// _ = "end of CoverTab[123688]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:211
	_go_fuzz_dep_.CoverTab[123689]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:214
	keyVal := parsedKey[len(parsedKey)-1]
	localKey := []string{keyVal}
	finalKey := append(tableKey, keyVal)
	if targetNode.GetPath(localKey) != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:217
		_go_fuzz_dep_.CoverTab[123705]++
												p.raiseError(key, "The following key was defined twice: %s",
			strings.Join(finalKey, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:219
		// _ = "end of CoverTab[123705]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:220
		_go_fuzz_dep_.CoverTab[123706]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:220
		// _ = "end of CoverTab[123706]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:220
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:220
	// _ = "end of CoverTab[123689]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:220
	_go_fuzz_dep_.CoverTab[123690]++
											var toInsert interface{}

											switch value.(type) {
	case *Tree, []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:224
		_go_fuzz_dep_.CoverTab[123707]++
												toInsert = value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:225
		// _ = "end of CoverTab[123707]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:226
		_go_fuzz_dep_.CoverTab[123708]++
												toInsert = &tomlValue{value: value, position: key.Position}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:227
		// _ = "end of CoverTab[123708]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:228
	// _ = "end of CoverTab[123690]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:228
	_go_fuzz_dep_.CoverTab[123691]++
											targetNode.values[keyVal] = toInsert
											return p.parseStart
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:230
	// _ = "end of CoverTab[123691]"
}

var errInvalidUnderscore = errors.New("invalid use of _ in number")

func numberContainsInvalidUnderscore(value string) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:235
	_go_fuzz_dep_.CoverTab[123709]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:240
	hasBefore := false
	for idx, r := range value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:241
		_go_fuzz_dep_.CoverTab[123711]++
												if r == '_' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:242
			_go_fuzz_dep_.CoverTab[123713]++
													if !hasBefore || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:243
				_go_fuzz_dep_.CoverTab[123714]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:243
				return idx+1 >= len(value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:243
				// _ = "end of CoverTab[123714]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:243
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:243
				_go_fuzz_dep_.CoverTab[123715]++

														return errInvalidUnderscore
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:245
				// _ = "end of CoverTab[123715]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:246
				_go_fuzz_dep_.CoverTab[123716]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:246
				// _ = "end of CoverTab[123716]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:246
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:246
			// _ = "end of CoverTab[123713]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:247
			_go_fuzz_dep_.CoverTab[123717]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:247
			// _ = "end of CoverTab[123717]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:247
		// _ = "end of CoverTab[123711]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:247
		_go_fuzz_dep_.CoverTab[123712]++
												hasBefore = isDigit(r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:248
		// _ = "end of CoverTab[123712]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:249
	// _ = "end of CoverTab[123709]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:249
	_go_fuzz_dep_.CoverTab[123710]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:250
	// _ = "end of CoverTab[123710]"
}

var errInvalidUnderscoreHex = errors.New("invalid use of _ in hex number")

func hexNumberContainsInvalidUnderscore(value string) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:255
	_go_fuzz_dep_.CoverTab[123718]++
											hasBefore := false
											for idx, r := range value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:257
		_go_fuzz_dep_.CoverTab[123720]++
												if r == '_' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:258
			_go_fuzz_dep_.CoverTab[123722]++
													if !hasBefore || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:259
				_go_fuzz_dep_.CoverTab[123723]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:259
				return idx+1 >= len(value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:259
				// _ = "end of CoverTab[123723]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:259
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:259
				_go_fuzz_dep_.CoverTab[123724]++

														return errInvalidUnderscoreHex
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:261
				// _ = "end of CoverTab[123724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:262
				_go_fuzz_dep_.CoverTab[123725]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:262
				// _ = "end of CoverTab[123725]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:262
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:262
			// _ = "end of CoverTab[123722]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:263
			_go_fuzz_dep_.CoverTab[123726]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:263
			// _ = "end of CoverTab[123726]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:263
		// _ = "end of CoverTab[123720]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:263
		_go_fuzz_dep_.CoverTab[123721]++
												hasBefore = isHexDigit(r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:264
		// _ = "end of CoverTab[123721]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:265
	// _ = "end of CoverTab[123718]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:265
	_go_fuzz_dep_.CoverTab[123719]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:266
	// _ = "end of CoverTab[123719]"
}

func cleanupNumberToken(value string) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:269
	_go_fuzz_dep_.CoverTab[123727]++
											cleanedVal := strings.Replace(value, "_", "", -1)
											return cleanedVal
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:271
	// _ = "end of CoverTab[123727]"
}

func (p *tomlParser) parseRvalue() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:274
	_go_fuzz_dep_.CoverTab[123728]++
											tok := p.getToken()
											if tok == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:276
		_go_fuzz_dep_.CoverTab[123731]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:276
		return tok.typ == tokenEOF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:276
		// _ = "end of CoverTab[123731]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:276
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:276
		_go_fuzz_dep_.CoverTab[123732]++
												p.raiseError(tok, "expecting a value")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:277
		// _ = "end of CoverTab[123732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:278
		_go_fuzz_dep_.CoverTab[123733]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:278
		// _ = "end of CoverTab[123733]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:278
	// _ = "end of CoverTab[123728]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:278
	_go_fuzz_dep_.CoverTab[123729]++

											switch tok.typ {
	case tokenString:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:281
		_go_fuzz_dep_.CoverTab[123734]++
												return tok.val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:282
		// _ = "end of CoverTab[123734]"
	case tokenTrue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:283
		_go_fuzz_dep_.CoverTab[123735]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:284
		// _ = "end of CoverTab[123735]"
	case tokenFalse:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:285
		_go_fuzz_dep_.CoverTab[123736]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:286
		// _ = "end of CoverTab[123736]"
	case tokenInf:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:287
		_go_fuzz_dep_.CoverTab[123737]++
												if tok.val[0] == '-' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:288
			_go_fuzz_dep_.CoverTab[123757]++
													return math.Inf(-1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:289
			// _ = "end of CoverTab[123757]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:290
			_go_fuzz_dep_.CoverTab[123758]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:290
			// _ = "end of CoverTab[123758]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:290
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:290
		// _ = "end of CoverTab[123737]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:290
		_go_fuzz_dep_.CoverTab[123738]++
												return math.Inf(1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:291
		// _ = "end of CoverTab[123738]"
	case tokenNan:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:292
		_go_fuzz_dep_.CoverTab[123739]++
												return math.NaN()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:293
		// _ = "end of CoverTab[123739]"
	case tokenInteger:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:294
		_go_fuzz_dep_.CoverTab[123740]++
												cleanedVal := cleanupNumberToken(tok.val)
												var err error
												var val int64
												if len(cleanedVal) >= 3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:298
			_go_fuzz_dep_.CoverTab[123759]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:298
			return cleanedVal[0] == '0'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:298
			// _ = "end of CoverTab[123759]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:298
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:298
			_go_fuzz_dep_.CoverTab[123760]++
													switch cleanedVal[1] {
			case 'x':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:300
				_go_fuzz_dep_.CoverTab[123761]++
														err = hexNumberContainsInvalidUnderscore(tok.val)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:302
					_go_fuzz_dep_.CoverTab[123768]++
															p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:303
					// _ = "end of CoverTab[123768]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:304
					_go_fuzz_dep_.CoverTab[123769]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:304
					// _ = "end of CoverTab[123769]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:304
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:304
				// _ = "end of CoverTab[123761]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:304
				_go_fuzz_dep_.CoverTab[123762]++
														val, err = strconv.ParseInt(cleanedVal[2:], 16, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:305
				// _ = "end of CoverTab[123762]"
			case 'o':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:306
				_go_fuzz_dep_.CoverTab[123763]++
														err = numberContainsInvalidUnderscore(tok.val)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:308
					_go_fuzz_dep_.CoverTab[123770]++
															p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:309
					// _ = "end of CoverTab[123770]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:310
					_go_fuzz_dep_.CoverTab[123771]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:310
					// _ = "end of CoverTab[123771]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:310
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:310
				// _ = "end of CoverTab[123763]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:310
				_go_fuzz_dep_.CoverTab[123764]++
														val, err = strconv.ParseInt(cleanedVal[2:], 8, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:311
				// _ = "end of CoverTab[123764]"
			case 'b':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:312
				_go_fuzz_dep_.CoverTab[123765]++
														err = numberContainsInvalidUnderscore(tok.val)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:314
					_go_fuzz_dep_.CoverTab[123772]++
															p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:315
					// _ = "end of CoverTab[123772]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:316
					_go_fuzz_dep_.CoverTab[123773]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:316
					// _ = "end of CoverTab[123773]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:316
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:316
				// _ = "end of CoverTab[123765]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:316
				_go_fuzz_dep_.CoverTab[123766]++
														val, err = strconv.ParseInt(cleanedVal[2:], 2, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:317
				// _ = "end of CoverTab[123766]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:318
				_go_fuzz_dep_.CoverTab[123767]++
														panic("invalid base")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:319
				// _ = "end of CoverTab[123767]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:320
			// _ = "end of CoverTab[123760]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:321
			_go_fuzz_dep_.CoverTab[123774]++
													err = numberContainsInvalidUnderscore(tok.val)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:323
				_go_fuzz_dep_.CoverTab[123776]++
														p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:324
				// _ = "end of CoverTab[123776]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:325
				_go_fuzz_dep_.CoverTab[123777]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:325
				// _ = "end of CoverTab[123777]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:325
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:325
			// _ = "end of CoverTab[123774]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:325
			_go_fuzz_dep_.CoverTab[123775]++
													val, err = strconv.ParseInt(cleanedVal, 10, 64)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:326
			// _ = "end of CoverTab[123775]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:327
		// _ = "end of CoverTab[123740]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:327
		_go_fuzz_dep_.CoverTab[123741]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:328
			_go_fuzz_dep_.CoverTab[123778]++
													p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:329
			// _ = "end of CoverTab[123778]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:330
			_go_fuzz_dep_.CoverTab[123779]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:330
			// _ = "end of CoverTab[123779]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:330
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:330
		// _ = "end of CoverTab[123741]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:330
		_go_fuzz_dep_.CoverTab[123742]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:331
		// _ = "end of CoverTab[123742]"
	case tokenFloat:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:332
		_go_fuzz_dep_.CoverTab[123743]++
												err := numberContainsInvalidUnderscore(tok.val)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:334
			_go_fuzz_dep_.CoverTab[123780]++
													p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:335
			// _ = "end of CoverTab[123780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:336
			_go_fuzz_dep_.CoverTab[123781]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:336
			// _ = "end of CoverTab[123781]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:336
		// _ = "end of CoverTab[123743]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:336
		_go_fuzz_dep_.CoverTab[123744]++
												cleanedVal := cleanupNumberToken(tok.val)
												val, err := strconv.ParseFloat(cleanedVal, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:339
			_go_fuzz_dep_.CoverTab[123782]++
													p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:340
			// _ = "end of CoverTab[123782]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:341
			_go_fuzz_dep_.CoverTab[123783]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:341
			// _ = "end of CoverTab[123783]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:341
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:341
		// _ = "end of CoverTab[123744]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:341
		_go_fuzz_dep_.CoverTab[123745]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:342
		// _ = "end of CoverTab[123745]"
	case tokenLocalTime:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:343
		_go_fuzz_dep_.CoverTab[123746]++
												val, err := ParseLocalTime(tok.val)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:345
			_go_fuzz_dep_.CoverTab[123784]++
													p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:346
			// _ = "end of CoverTab[123784]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:347
			_go_fuzz_dep_.CoverTab[123785]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:347
			// _ = "end of CoverTab[123785]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:347
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:347
		// _ = "end of CoverTab[123746]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:347
		_go_fuzz_dep_.CoverTab[123747]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:348
		// _ = "end of CoverTab[123747]"
	case tokenLocalDate:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:349
		_go_fuzz_dep_.CoverTab[123748]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:354
		next := p.peek()
		if next == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:355
			_go_fuzz_dep_.CoverTab[123786]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:355
			return next.typ != tokenLocalTime
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:355
			// _ = "end of CoverTab[123786]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:355
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:355
			_go_fuzz_dep_.CoverTab[123787]++
													val, err := ParseLocalDate(tok.val)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:357
				_go_fuzz_dep_.CoverTab[123789]++
														p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:358
				// _ = "end of CoverTab[123789]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:359
				_go_fuzz_dep_.CoverTab[123790]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:359
				// _ = "end of CoverTab[123790]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:359
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:359
			// _ = "end of CoverTab[123787]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:359
			_go_fuzz_dep_.CoverTab[123788]++
													return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:360
			// _ = "end of CoverTab[123788]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:361
			_go_fuzz_dep_.CoverTab[123791]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:361
			// _ = "end of CoverTab[123791]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:361
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:361
		// _ = "end of CoverTab[123748]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:361
		_go_fuzz_dep_.CoverTab[123749]++

												localDate := tok
												localTime := p.getToken()

												next = p.peek()
												if next == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:367
			_go_fuzz_dep_.CoverTab[123792]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:367
			return next.typ != tokenTimeOffset
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:367
			// _ = "end of CoverTab[123792]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:367
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:367
			_go_fuzz_dep_.CoverTab[123793]++
													v := localDate.val + "T" + localTime.val
													val, err := ParseLocalDateTime(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:370
				_go_fuzz_dep_.CoverTab[123795]++
														p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:371
				// _ = "end of CoverTab[123795]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:372
				_go_fuzz_dep_.CoverTab[123796]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:372
				// _ = "end of CoverTab[123796]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:372
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:372
			// _ = "end of CoverTab[123793]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:372
			_go_fuzz_dep_.CoverTab[123794]++
													return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:373
			// _ = "end of CoverTab[123794]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:374
			_go_fuzz_dep_.CoverTab[123797]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:374
			// _ = "end of CoverTab[123797]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:374
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:374
		// _ = "end of CoverTab[123749]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:374
		_go_fuzz_dep_.CoverTab[123750]++

												offset := p.getToken()

												layout := time.RFC3339Nano
												v := localDate.val + "T" + localTime.val + offset.val
												val, err := time.ParseInLocation(layout, v, time.UTC)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:381
			_go_fuzz_dep_.CoverTab[123798]++
													p.raiseError(tok, "%s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:382
			// _ = "end of CoverTab[123798]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:383
			_go_fuzz_dep_.CoverTab[123799]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:383
			// _ = "end of CoverTab[123799]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:383
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:383
		// _ = "end of CoverTab[123750]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:383
		_go_fuzz_dep_.CoverTab[123751]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:384
		// _ = "end of CoverTab[123751]"
	case tokenLeftBracket:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:385
		_go_fuzz_dep_.CoverTab[123752]++
												return p.parseArray()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:386
		// _ = "end of CoverTab[123752]"
	case tokenLeftCurlyBrace:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:387
		_go_fuzz_dep_.CoverTab[123753]++
												return p.parseInlineTable()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:388
		// _ = "end of CoverTab[123753]"
	case tokenEqual:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:389
		_go_fuzz_dep_.CoverTab[123754]++
												p.raiseError(tok, "cannot have multiple equals for the same key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:390
		// _ = "end of CoverTab[123754]"
	case tokenError:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:391
		_go_fuzz_dep_.CoverTab[123755]++
												p.raiseError(tok, "%s", tok)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:392
		// _ = "end of CoverTab[123755]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:393
		_go_fuzz_dep_.CoverTab[123756]++
												panic(fmt.Errorf("unhandled token: %v", tok))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:394
		// _ = "end of CoverTab[123756]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:395
	// _ = "end of CoverTab[123729]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:395
	_go_fuzz_dep_.CoverTab[123730]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:397
	// _ = "end of CoverTab[123730]"
}

func tokenIsComma(t *token) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:400
	_go_fuzz_dep_.CoverTab[123800]++
											return t != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:401
		_go_fuzz_dep_.CoverTab[123801]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:401
		return t.typ == tokenComma
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:401
		// _ = "end of CoverTab[123801]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:401
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:401
	// _ = "end of CoverTab[123800]"
}

func (p *tomlParser) parseInlineTable() *Tree {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:404
	_go_fuzz_dep_.CoverTab[123802]++
											tree := newTree()
											var previous *token
Loop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:408
		_go_fuzz_dep_.CoverTab[123805]++
												follow := p.peek()
												if follow == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:410
			_go_fuzz_dep_.CoverTab[123808]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:410
			return follow.typ == tokenEOF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:410
			// _ = "end of CoverTab[123808]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:410
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:410
			_go_fuzz_dep_.CoverTab[123809]++
													p.raiseError(follow, "unterminated inline table")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:411
			// _ = "end of CoverTab[123809]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:412
			_go_fuzz_dep_.CoverTab[123810]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:412
			// _ = "end of CoverTab[123810]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:412
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:412
		// _ = "end of CoverTab[123805]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:412
		_go_fuzz_dep_.CoverTab[123806]++
												switch follow.typ {
		case tokenRightCurlyBrace:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:414
			_go_fuzz_dep_.CoverTab[123811]++
													p.getToken()
													break Loop
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:416
			// _ = "end of CoverTab[123811]"
		case tokenKey, tokenInteger, tokenString:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:417
			_go_fuzz_dep_.CoverTab[123812]++
													if !tokenIsComma(previous) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:418
				_go_fuzz_dep_.CoverTab[123818]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:418
				return previous != nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:418
				// _ = "end of CoverTab[123818]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:418
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:418
				_go_fuzz_dep_.CoverTab[123819]++
														p.raiseError(follow, "comma expected between fields in inline table")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:419
				// _ = "end of CoverTab[123819]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:420
				_go_fuzz_dep_.CoverTab[123820]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:420
				// _ = "end of CoverTab[123820]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:420
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:420
			// _ = "end of CoverTab[123812]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:420
			_go_fuzz_dep_.CoverTab[123813]++
													key := p.getToken()
													p.assume(tokenEqual)

													parsedKey, err := parseKey(key.val)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:425
				_go_fuzz_dep_.CoverTab[123821]++
														p.raiseError(key, "invalid key: %s", err)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:426
				// _ = "end of CoverTab[123821]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:427
				_go_fuzz_dep_.CoverTab[123822]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:427
				// _ = "end of CoverTab[123822]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:427
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:427
			// _ = "end of CoverTab[123813]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:427
			_go_fuzz_dep_.CoverTab[123814]++

													value := p.parseRvalue()
													tree.SetPath(parsedKey, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:430
			// _ = "end of CoverTab[123814]"
		case tokenComma:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:431
			_go_fuzz_dep_.CoverTab[123815]++
													if tokenIsComma(previous) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:432
				_go_fuzz_dep_.CoverTab[123823]++
														p.raiseError(follow, "need field between two commas in inline table")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:433
				// _ = "end of CoverTab[123823]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:434
				_go_fuzz_dep_.CoverTab[123824]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:434
				// _ = "end of CoverTab[123824]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:434
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:434
			// _ = "end of CoverTab[123815]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:434
			_go_fuzz_dep_.CoverTab[123816]++
													p.getToken()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:435
			// _ = "end of CoverTab[123816]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:436
			_go_fuzz_dep_.CoverTab[123817]++
													p.raiseError(follow, "unexpected token type in inline table: %s", follow.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:437
			// _ = "end of CoverTab[123817]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:438
		// _ = "end of CoverTab[123806]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:438
		_go_fuzz_dep_.CoverTab[123807]++
												previous = follow
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:439
		// _ = "end of CoverTab[123807]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:440
	// _ = "end of CoverTab[123802]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:440
	_go_fuzz_dep_.CoverTab[123803]++
											if tokenIsComma(previous) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:441
		_go_fuzz_dep_.CoverTab[123825]++
												p.raiseError(previous, "trailing comma at the end of inline table")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:442
		// _ = "end of CoverTab[123825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:443
		_go_fuzz_dep_.CoverTab[123826]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:443
		// _ = "end of CoverTab[123826]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:443
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:443
	// _ = "end of CoverTab[123803]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:443
	_go_fuzz_dep_.CoverTab[123804]++
											tree.inline = true
											return tree
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:445
	// _ = "end of CoverTab[123804]"
}

func (p *tomlParser) parseArray() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:448
	_go_fuzz_dep_.CoverTab[123827]++
											var array []interface{}
											arrayType := reflect.TypeOf(newTree())
											for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:451
		_go_fuzz_dep_.CoverTab[123831]++
												follow := p.peek()
												if follow == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:453
			_go_fuzz_dep_.CoverTab[123837]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:453
			return follow.typ == tokenEOF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:453
			// _ = "end of CoverTab[123837]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:453
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:453
			_go_fuzz_dep_.CoverTab[123838]++
													p.raiseError(follow, "unterminated array")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:454
			// _ = "end of CoverTab[123838]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:455
			_go_fuzz_dep_.CoverTab[123839]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:455
			// _ = "end of CoverTab[123839]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:455
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:455
		// _ = "end of CoverTab[123831]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:455
		_go_fuzz_dep_.CoverTab[123832]++
												if follow.typ == tokenRightBracket {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:456
			_go_fuzz_dep_.CoverTab[123840]++
													p.getToken()
													break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:458
			// _ = "end of CoverTab[123840]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:459
			_go_fuzz_dep_.CoverTab[123841]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:459
			// _ = "end of CoverTab[123841]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:459
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:459
		// _ = "end of CoverTab[123832]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:459
		_go_fuzz_dep_.CoverTab[123833]++
												val := p.parseRvalue()
												if reflect.TypeOf(val) != arrayType {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:461
			_go_fuzz_dep_.CoverTab[123842]++
													arrayType = nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:462
			// _ = "end of CoverTab[123842]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:463
			_go_fuzz_dep_.CoverTab[123843]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:463
			// _ = "end of CoverTab[123843]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:463
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:463
		// _ = "end of CoverTab[123833]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:463
		_go_fuzz_dep_.CoverTab[123834]++
												array = append(array, val)
												follow = p.peek()
												if follow == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:466
			_go_fuzz_dep_.CoverTab[123844]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:466
			return follow.typ == tokenEOF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:466
			// _ = "end of CoverTab[123844]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:466
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:466
			_go_fuzz_dep_.CoverTab[123845]++
													p.raiseError(follow, "unterminated array")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:467
			// _ = "end of CoverTab[123845]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:468
			_go_fuzz_dep_.CoverTab[123846]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:468
			// _ = "end of CoverTab[123846]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:468
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:468
		// _ = "end of CoverTab[123834]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:468
		_go_fuzz_dep_.CoverTab[123835]++
												if follow.typ != tokenRightBracket && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:469
			_go_fuzz_dep_.CoverTab[123847]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:469
			return follow.typ != tokenComma
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:469
			// _ = "end of CoverTab[123847]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:469
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:469
			_go_fuzz_dep_.CoverTab[123848]++
													p.raiseError(follow, "missing comma")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:470
			// _ = "end of CoverTab[123848]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:471
			_go_fuzz_dep_.CoverTab[123849]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:471
			// _ = "end of CoverTab[123849]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:471
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:471
		// _ = "end of CoverTab[123835]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:471
		_go_fuzz_dep_.CoverTab[123836]++
												if follow.typ == tokenComma {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:472
			_go_fuzz_dep_.CoverTab[123850]++
													p.getToken()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:473
			// _ = "end of CoverTab[123850]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:474
			_go_fuzz_dep_.CoverTab[123851]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:474
			// _ = "end of CoverTab[123851]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:474
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:474
		// _ = "end of CoverTab[123836]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:475
	// _ = "end of CoverTab[123827]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:475
	_go_fuzz_dep_.CoverTab[123828]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:479
	if len(array) <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:479
		_go_fuzz_dep_.CoverTab[123852]++
												arrayType = nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:480
		// _ = "end of CoverTab[123852]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:481
		_go_fuzz_dep_.CoverTab[123853]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:481
		// _ = "end of CoverTab[123853]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:481
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:481
	// _ = "end of CoverTab[123828]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:481
	_go_fuzz_dep_.CoverTab[123829]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:486
	if arrayType == reflect.TypeOf(newTree()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:486
		_go_fuzz_dep_.CoverTab[123854]++
												tomlArray := make([]*Tree, len(array))
												for i, v := range array {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:488
			_go_fuzz_dep_.CoverTab[123856]++
													tomlArray[i] = v.(*Tree)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:489
			// _ = "end of CoverTab[123856]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:490
		// _ = "end of CoverTab[123854]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:490
		_go_fuzz_dep_.CoverTab[123855]++
												return tomlArray
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:491
		// _ = "end of CoverTab[123855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:492
		_go_fuzz_dep_.CoverTab[123857]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:492
		// _ = "end of CoverTab[123857]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:492
	// _ = "end of CoverTab[123829]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:492
	_go_fuzz_dep_.CoverTab[123830]++
											return array
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:493
	// _ = "end of CoverTab[123830]"
}

func parseToml(flow []token) *Tree {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:496
	_go_fuzz_dep_.CoverTab[123858]++
											result := newTree()
											result.position = Position{1, 1}
											parser := &tomlParser{
		flowIdx:	0,
		flow:		flow,
		tree:		result,
		currentTable:	make([]string, 0),
		seenTableKeys:	make([]string, 0),
	}
											parser.run()
											return result
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:507
	// _ = "end of CoverTab[123858]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:508
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/parser.go:508
var _ = _go_fuzz_dep_.CoverTab
