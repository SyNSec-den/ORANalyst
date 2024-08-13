//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:1
)

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

type tomlValue struct {
	value		interface{}	// string, int64, uint64, float64, bool, time.Time, [] of any of this list
	comment		string
	commented	bool
	multiline	bool
	literal		bool
	position	Position
}

// Tree is the result of the parsing of a TOML file.
type Tree struct {
	values		map[string]interface{}	// string -> *tomlValue, *Tree, []*Tree
	comment		string
	commented	bool
	inline		bool
	position	Position
}

func newTree() *Tree {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:31
	_go_fuzz_dep_.CoverTab[123894]++
											return newTreeWithPosition(Position{})
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:32
	// _ = "end of CoverTab[123894]"
}

func newTreeWithPosition(pos Position) *Tree {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:35
	_go_fuzz_dep_.CoverTab[123895]++
											return &Tree{
		values:		make(map[string]interface{}),
		position:	pos,
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:39
	// _ = "end of CoverTab[123895]"
}

// TreeFromMap initializes a new Tree object using the given map.
func TreeFromMap(m map[string]interface{}) (*Tree, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:43
	_go_fuzz_dep_.CoverTab[123896]++
											result, err := toTree(m)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:45
		_go_fuzz_dep_.CoverTab[123898]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:46
		// _ = "end of CoverTab[123898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:47
		_go_fuzz_dep_.CoverTab[123899]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:47
		// _ = "end of CoverTab[123899]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:47
	// _ = "end of CoverTab[123896]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:47
	_go_fuzz_dep_.CoverTab[123897]++
											return result.(*Tree), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:48
	// _ = "end of CoverTab[123897]"
}

// Position returns the position of the tree.
func (t *Tree) Position() Position {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:52
	_go_fuzz_dep_.CoverTab[123900]++
											return t.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:53
	// _ = "end of CoverTab[123900]"
}

// Has returns a boolean indicating if the given key exists.
func (t *Tree) Has(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:57
	_go_fuzz_dep_.CoverTab[123901]++
											if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:58
		_go_fuzz_dep_.CoverTab[123903]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:59
		// _ = "end of CoverTab[123903]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:60
		_go_fuzz_dep_.CoverTab[123904]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:60
		// _ = "end of CoverTab[123904]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:60
	// _ = "end of CoverTab[123901]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:60
	_go_fuzz_dep_.CoverTab[123902]++
											return t.HasPath(strings.Split(key, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:61
	// _ = "end of CoverTab[123902]"
}

// HasPath returns true if the given path of keys exists, false otherwise.
func (t *Tree) HasPath(keys []string) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:65
	_go_fuzz_dep_.CoverTab[123905]++
											return t.GetPath(keys) != nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:66
	// _ = "end of CoverTab[123905]"
}

// Keys returns the keys of the toplevel tree (does not recurse).
func (t *Tree) Keys() []string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:70
	_go_fuzz_dep_.CoverTab[123906]++
											keys := make([]string, len(t.values))
											i := 0
											for k := range t.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:73
		_go_fuzz_dep_.CoverTab[123908]++
												keys[i] = k
												i++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:75
		// _ = "end of CoverTab[123908]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:76
	// _ = "end of CoverTab[123906]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:76
	_go_fuzz_dep_.CoverTab[123907]++
											return keys
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:77
	// _ = "end of CoverTab[123907]"
}

// Get the value at key in the Tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:80
// Key is a dot-separated path (e.g. a.b.c) without single/double quoted strings.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:80
// If you need to retrieve non-bare keys, use GetPath.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:80
// Returns nil if the path does not exist in the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:80
// If keys is of length zero, the current tree is returned.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:85
func (t *Tree) Get(key string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:85
	_go_fuzz_dep_.CoverTab[123909]++
											if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:86
		_go_fuzz_dep_.CoverTab[123911]++
												return t
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:87
		// _ = "end of CoverTab[123911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:88
		_go_fuzz_dep_.CoverTab[123912]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:88
		// _ = "end of CoverTab[123912]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:88
	// _ = "end of CoverTab[123909]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:88
	_go_fuzz_dep_.CoverTab[123910]++
											return t.GetPath(strings.Split(key, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:89
	// _ = "end of CoverTab[123910]"
}

// GetPath returns the element in the tree indicated by 'keys'.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:92
// If keys is of length zero, the current tree is returned.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:94
func (t *Tree) GetPath(keys []string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:94
	_go_fuzz_dep_.CoverTab[123913]++
											if len(keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:95
		_go_fuzz_dep_.CoverTab[123916]++
												return t
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:96
		// _ = "end of CoverTab[123916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:97
		_go_fuzz_dep_.CoverTab[123917]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:97
		// _ = "end of CoverTab[123917]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:97
	// _ = "end of CoverTab[123913]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:97
	_go_fuzz_dep_.CoverTab[123914]++
											subtree := t
											for _, intermediateKey := range keys[:len(keys)-1] {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:99
		_go_fuzz_dep_.CoverTab[123918]++
												value, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:101
			_go_fuzz_dep_.CoverTab[123920]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:102
			// _ = "end of CoverTab[123920]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:103
			_go_fuzz_dep_.CoverTab[123921]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:103
			// _ = "end of CoverTab[123921]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:103
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:103
		// _ = "end of CoverTab[123918]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:103
		_go_fuzz_dep_.CoverTab[123919]++
												switch node := value.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:105
			_go_fuzz_dep_.CoverTab[123922]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:106
			// _ = "end of CoverTab[123922]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:107
			_go_fuzz_dep_.CoverTab[123923]++

													if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:109
				_go_fuzz_dep_.CoverTab[123926]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:110
				// _ = "end of CoverTab[123926]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:111
				_go_fuzz_dep_.CoverTab[123927]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:111
				// _ = "end of CoverTab[123927]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:111
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:111
			// _ = "end of CoverTab[123923]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:111
			_go_fuzz_dep_.CoverTab[123924]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:112
			// _ = "end of CoverTab[123924]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:113
			_go_fuzz_dep_.CoverTab[123925]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:114
			// _ = "end of CoverTab[123925]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:115
		// _ = "end of CoverTab[123919]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:116
	// _ = "end of CoverTab[123914]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:116
	_go_fuzz_dep_.CoverTab[123915]++

											switch node := subtree.values[keys[len(keys)-1]].(type) {
	case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:119
		_go_fuzz_dep_.CoverTab[123928]++
												return node.value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:120
		// _ = "end of CoverTab[123928]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:121
		_go_fuzz_dep_.CoverTab[123929]++
												return node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:122
		// _ = "end of CoverTab[123929]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:123
	// _ = "end of CoverTab[123915]"
}

// GetArray returns the value at key in the Tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:126
// It returns []string, []int64, etc type if key has homogeneous lists
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:126
// Key is a dot-separated path (e.g. a.b.c) without single/double quoted strings.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:126
// Returns nil if the path does not exist in the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:126
// If keys is of length zero, the current tree is returned.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:131
func (t *Tree) GetArray(key string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:131
	_go_fuzz_dep_.CoverTab[123930]++
											if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:132
		_go_fuzz_dep_.CoverTab[123932]++
												return t
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:133
		// _ = "end of CoverTab[123932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:134
		_go_fuzz_dep_.CoverTab[123933]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:134
		// _ = "end of CoverTab[123933]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:134
	// _ = "end of CoverTab[123930]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:134
	_go_fuzz_dep_.CoverTab[123931]++
											return t.GetArrayPath(strings.Split(key, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:135
	// _ = "end of CoverTab[123931]"
}

// GetArrayPath returns the element in the tree indicated by 'keys'.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:138
// If keys is of length zero, the current tree is returned.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:140
func (t *Tree) GetArrayPath(keys []string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:140
	_go_fuzz_dep_.CoverTab[123934]++
											if len(keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:141
		_go_fuzz_dep_.CoverTab[123937]++
												return t
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:142
		// _ = "end of CoverTab[123937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:143
		_go_fuzz_dep_.CoverTab[123938]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:143
		// _ = "end of CoverTab[123938]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:143
	// _ = "end of CoverTab[123934]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:143
	_go_fuzz_dep_.CoverTab[123935]++
											subtree := t
											for _, intermediateKey := range keys[:len(keys)-1] {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:145
		_go_fuzz_dep_.CoverTab[123939]++
												value, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:147
			_go_fuzz_dep_.CoverTab[123941]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:148
			// _ = "end of CoverTab[123941]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:149
			_go_fuzz_dep_.CoverTab[123942]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:149
			// _ = "end of CoverTab[123942]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:149
		// _ = "end of CoverTab[123939]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:149
		_go_fuzz_dep_.CoverTab[123940]++
												switch node := value.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:151
			_go_fuzz_dep_.CoverTab[123943]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:152
			// _ = "end of CoverTab[123943]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:153
			_go_fuzz_dep_.CoverTab[123944]++

													if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:155
				_go_fuzz_dep_.CoverTab[123947]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:156
				// _ = "end of CoverTab[123947]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:157
				_go_fuzz_dep_.CoverTab[123948]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:157
				// _ = "end of CoverTab[123948]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:157
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:157
			// _ = "end of CoverTab[123944]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:157
			_go_fuzz_dep_.CoverTab[123945]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:158
			// _ = "end of CoverTab[123945]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:159
			_go_fuzz_dep_.CoverTab[123946]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:160
			// _ = "end of CoverTab[123946]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:161
		// _ = "end of CoverTab[123940]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:162
	// _ = "end of CoverTab[123935]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:162
	_go_fuzz_dep_.CoverTab[123936]++

											switch node := subtree.values[keys[len(keys)-1]].(type) {
	case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:165
		_go_fuzz_dep_.CoverTab[123949]++
												switch n := node.value.(type) {
		case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:167
			_go_fuzz_dep_.CoverTab[123951]++
													return getArray(n)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:168
			// _ = "end of CoverTab[123951]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:169
			_go_fuzz_dep_.CoverTab[123952]++
													return node.value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:170
			// _ = "end of CoverTab[123952]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:171
		// _ = "end of CoverTab[123949]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:172
		_go_fuzz_dep_.CoverTab[123950]++
												return node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:173
		// _ = "end of CoverTab[123950]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:174
	// _ = "end of CoverTab[123936]"
}

// if homogeneous array, then return slice type object over []interface{}
func getArray(n []interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:178
	_go_fuzz_dep_.CoverTab[123953]++
											var s []string
											var i64 []int64
											var f64 []float64
											var bl []bool
											for _, value := range n {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:183
		_go_fuzz_dep_.CoverTab[123956]++
												switch v := value.(type) {
		case string:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:185
			_go_fuzz_dep_.CoverTab[123957]++
													s = append(s, v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:186
			// _ = "end of CoverTab[123957]"
		case int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:187
			_go_fuzz_dep_.CoverTab[123958]++
													i64 = append(i64, v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:188
			// _ = "end of CoverTab[123958]"
		case float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:189
			_go_fuzz_dep_.CoverTab[123959]++
													f64 = append(f64, v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:190
			// _ = "end of CoverTab[123959]"
		case bool:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:191
			_go_fuzz_dep_.CoverTab[123960]++
													bl = append(bl, v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:192
			// _ = "end of CoverTab[123960]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:193
			_go_fuzz_dep_.CoverTab[123961]++
													return n
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:194
			// _ = "end of CoverTab[123961]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:195
		// _ = "end of CoverTab[123956]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:196
	// _ = "end of CoverTab[123953]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:196
	_go_fuzz_dep_.CoverTab[123954]++
											if len(s) == len(n) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:197
		_go_fuzz_dep_.CoverTab[123962]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:198
		// _ = "end of CoverTab[123962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:199
		_go_fuzz_dep_.CoverTab[123963]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:199
		if len(i64) == len(n) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:199
			_go_fuzz_dep_.CoverTab[123964]++
													return i64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:200
			// _ = "end of CoverTab[123964]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:201
			_go_fuzz_dep_.CoverTab[123965]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:201
			if len(f64) == len(n) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:201
				_go_fuzz_dep_.CoverTab[123966]++
														return f64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:202
				// _ = "end of CoverTab[123966]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:203
				_go_fuzz_dep_.CoverTab[123967]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:203
				if len(bl) == len(n) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:203
					_go_fuzz_dep_.CoverTab[123968]++
															return bl
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:204
					// _ = "end of CoverTab[123968]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
					_go_fuzz_dep_.CoverTab[123969]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
					// _ = "end of CoverTab[123969]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
				// _ = "end of CoverTab[123967]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
			// _ = "end of CoverTab[123965]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
		// _ = "end of CoverTab[123963]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
	// _ = "end of CoverTab[123954]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:205
	_go_fuzz_dep_.CoverTab[123955]++
											return n
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:206
	// _ = "end of CoverTab[123955]"
}

// GetPosition returns the position of the given key.
func (t *Tree) GetPosition(key string) Position {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:210
	_go_fuzz_dep_.CoverTab[123970]++
											if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:211
		_go_fuzz_dep_.CoverTab[123972]++
												return t.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:212
		// _ = "end of CoverTab[123972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:213
		_go_fuzz_dep_.CoverTab[123973]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:213
		// _ = "end of CoverTab[123973]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:213
	// _ = "end of CoverTab[123970]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:213
	_go_fuzz_dep_.CoverTab[123971]++
											return t.GetPositionPath(strings.Split(key, "."))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:214
	// _ = "end of CoverTab[123971]"
}

// SetPositionPath sets the position of element in the tree indicated by 'keys'.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:217
// If keys is of length zero, the current tree position is set.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:219
func (t *Tree) SetPositionPath(keys []string, pos Position) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:219
	_go_fuzz_dep_.CoverTab[123974]++
											if len(keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:220
		_go_fuzz_dep_.CoverTab[123977]++
												t.position = pos
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:222
		// _ = "end of CoverTab[123977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:223
		_go_fuzz_dep_.CoverTab[123978]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:223
		// _ = "end of CoverTab[123978]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:223
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:223
	// _ = "end of CoverTab[123974]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:223
	_go_fuzz_dep_.CoverTab[123975]++
											subtree := t
											for _, intermediateKey := range keys[:len(keys)-1] {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:225
		_go_fuzz_dep_.CoverTab[123979]++
												value, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:227
			_go_fuzz_dep_.CoverTab[123981]++
													return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:228
			// _ = "end of CoverTab[123981]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:229
			_go_fuzz_dep_.CoverTab[123982]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:229
			// _ = "end of CoverTab[123982]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:229
		// _ = "end of CoverTab[123979]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:229
		_go_fuzz_dep_.CoverTab[123980]++
												switch node := value.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:231
			_go_fuzz_dep_.CoverTab[123983]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:232
			// _ = "end of CoverTab[123983]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:233
			_go_fuzz_dep_.CoverTab[123984]++

													if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:235
				_go_fuzz_dep_.CoverTab[123987]++
														return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:236
				// _ = "end of CoverTab[123987]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:237
				_go_fuzz_dep_.CoverTab[123988]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:237
				// _ = "end of CoverTab[123988]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:237
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:237
			// _ = "end of CoverTab[123984]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:237
			_go_fuzz_dep_.CoverTab[123985]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:238
			// _ = "end of CoverTab[123985]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:239
			_go_fuzz_dep_.CoverTab[123986]++
													return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:240
			// _ = "end of CoverTab[123986]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:241
		// _ = "end of CoverTab[123980]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:242
	// _ = "end of CoverTab[123975]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:242
	_go_fuzz_dep_.CoverTab[123976]++

											switch node := subtree.values[keys[len(keys)-1]].(type) {
	case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:245
		_go_fuzz_dep_.CoverTab[123989]++
												node.position = pos
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:247
		// _ = "end of CoverTab[123989]"
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:248
		_go_fuzz_dep_.CoverTab[123990]++
												node.position = pos
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:250
		// _ = "end of CoverTab[123990]"
	case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:251
		_go_fuzz_dep_.CoverTab[123991]++

												if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:253
			_go_fuzz_dep_.CoverTab[123993]++
													return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:254
			// _ = "end of CoverTab[123993]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:255
			_go_fuzz_dep_.CoverTab[123994]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:255
			// _ = "end of CoverTab[123994]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:255
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:255
		// _ = "end of CoverTab[123991]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:255
		_go_fuzz_dep_.CoverTab[123992]++
												node[len(node)-1].position = pos
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:257
		// _ = "end of CoverTab[123992]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:258
	// _ = "end of CoverTab[123976]"
}

// GetPositionPath returns the element in the tree indicated by 'keys'.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:261
// If keys is of length zero, the current tree is returned.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:263
func (t *Tree) GetPositionPath(keys []string) Position {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:263
	_go_fuzz_dep_.CoverTab[123995]++
											if len(keys) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:264
		_go_fuzz_dep_.CoverTab[123998]++
												return t.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:265
		// _ = "end of CoverTab[123998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:266
		_go_fuzz_dep_.CoverTab[123999]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:266
		// _ = "end of CoverTab[123999]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:266
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:266
	// _ = "end of CoverTab[123995]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:266
	_go_fuzz_dep_.CoverTab[123996]++
											subtree := t
											for _, intermediateKey := range keys[:len(keys)-1] {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:268
		_go_fuzz_dep_.CoverTab[124000]++
												value, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:270
			_go_fuzz_dep_.CoverTab[124002]++
													return Position{0, 0}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:271
			// _ = "end of CoverTab[124002]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:272
			_go_fuzz_dep_.CoverTab[124003]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:272
			// _ = "end of CoverTab[124003]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:272
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:272
		// _ = "end of CoverTab[124000]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:272
		_go_fuzz_dep_.CoverTab[124001]++
												switch node := value.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:274
			_go_fuzz_dep_.CoverTab[124004]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:275
			// _ = "end of CoverTab[124004]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:276
			_go_fuzz_dep_.CoverTab[124005]++

													if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:278
				_go_fuzz_dep_.CoverTab[124008]++
														return Position{0, 0}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:279
				// _ = "end of CoverTab[124008]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:280
				_go_fuzz_dep_.CoverTab[124009]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:280
				// _ = "end of CoverTab[124009]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:280
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:280
			// _ = "end of CoverTab[124005]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:280
			_go_fuzz_dep_.CoverTab[124006]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:281
			// _ = "end of CoverTab[124006]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:282
			_go_fuzz_dep_.CoverTab[124007]++
													return Position{0, 0}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:283
			// _ = "end of CoverTab[124007]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:284
		// _ = "end of CoverTab[124001]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:285
	// _ = "end of CoverTab[123996]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:285
	_go_fuzz_dep_.CoverTab[123997]++

											switch node := subtree.values[keys[len(keys)-1]].(type) {
	case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:288
		_go_fuzz_dep_.CoverTab[124010]++
												return node.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:289
		// _ = "end of CoverTab[124010]"
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:290
		_go_fuzz_dep_.CoverTab[124011]++
												return node.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:291
		// _ = "end of CoverTab[124011]"
	case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:292
		_go_fuzz_dep_.CoverTab[124012]++

												if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:294
			_go_fuzz_dep_.CoverTab[124015]++
													return Position{0, 0}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:295
			// _ = "end of CoverTab[124015]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:296
			_go_fuzz_dep_.CoverTab[124016]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:296
			// _ = "end of CoverTab[124016]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:296
		// _ = "end of CoverTab[124012]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:296
		_go_fuzz_dep_.CoverTab[124013]++
												return node[len(node)-1].position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:297
		// _ = "end of CoverTab[124013]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:298
		_go_fuzz_dep_.CoverTab[124014]++
												return Position{0, 0}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:299
		// _ = "end of CoverTab[124014]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:300
	// _ = "end of CoverTab[123997]"
}

// GetDefault works like Get but with a default value
func (t *Tree) GetDefault(key string, def interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:304
	_go_fuzz_dep_.CoverTab[124017]++
											val := t.Get(key)
											if val == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:306
		_go_fuzz_dep_.CoverTab[124019]++
												return def
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:307
		// _ = "end of CoverTab[124019]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:308
		_go_fuzz_dep_.CoverTab[124020]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:308
		// _ = "end of CoverTab[124020]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:308
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:308
	// _ = "end of CoverTab[124017]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:308
	_go_fuzz_dep_.CoverTab[124018]++
											return val
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:309
	// _ = "end of CoverTab[124018]"
}

// SetOptions arguments are supplied to the SetWithOptions and SetPathWithOptions functions to modify marshalling behaviour.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:312
// The default values within the struct are valid default options.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:314
type SetOptions struct {
	Comment		string
	Commented	bool
	Multiline	bool
	Literal		bool
}

// SetWithOptions is the same as Set, but allows you to provide formatting
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:321
// instructions to the key, that will be used by Marshal().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:323
func (t *Tree) SetWithOptions(key string, opts SetOptions, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:323
	_go_fuzz_dep_.CoverTab[124021]++
											t.SetPathWithOptions(strings.Split(key, "."), opts, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:324
	// _ = "end of CoverTab[124021]"
}

// SetPathWithOptions is the same as SetPath, but allows you to provide
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:327
// formatting instructions to the key, that will be reused by Marshal().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:329
func (t *Tree) SetPathWithOptions(keys []string, opts SetOptions, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:329
	_go_fuzz_dep_.CoverTab[124022]++
											subtree := t
											for i, intermediateKey := range keys[:len(keys)-1] {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:331
		_go_fuzz_dep_.CoverTab[124025]++
												nextTree, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:333
			_go_fuzz_dep_.CoverTab[124027]++
													nextTree = newTreeWithPosition(Position{Line: t.position.Line + i, Col: t.position.Col})
													subtree.values[intermediateKey] = nextTree
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:335
			// _ = "end of CoverTab[124027]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:336
			_go_fuzz_dep_.CoverTab[124028]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:336
			// _ = "end of CoverTab[124028]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:336
		// _ = "end of CoverTab[124025]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:336
		_go_fuzz_dep_.CoverTab[124026]++
												switch node := nextTree.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:338
			_go_fuzz_dep_.CoverTab[124029]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:339
			// _ = "end of CoverTab[124029]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:340
			_go_fuzz_dep_.CoverTab[124030]++

													if len(node) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:342
				_go_fuzz_dep_.CoverTab[124032]++

														node = append(node, newTreeWithPosition(Position{Line: t.position.Line + i, Col: t.position.Col}))
														subtree.values[intermediateKey] = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:345
				// _ = "end of CoverTab[124032]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:346
				_go_fuzz_dep_.CoverTab[124033]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:346
				// _ = "end of CoverTab[124033]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:346
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:346
			// _ = "end of CoverTab[124030]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:346
			_go_fuzz_dep_.CoverTab[124031]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:347
			// _ = "end of CoverTab[124031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:348
		// _ = "end of CoverTab[124026]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:349
	// _ = "end of CoverTab[124022]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:349
	_go_fuzz_dep_.CoverTab[124023]++

											var toInsert interface{}

											switch v := value.(type) {
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:354
		_go_fuzz_dep_.CoverTab[124034]++
												v.comment = opts.Comment
												v.commented = opts.Commented
												toInsert = value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:357
		// _ = "end of CoverTab[124034]"
	case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:358
		_go_fuzz_dep_.CoverTab[124035]++
												for i := range v {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:359
			_go_fuzz_dep_.CoverTab[124039]++
													v[i].commented = opts.Commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:360
			// _ = "end of CoverTab[124039]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:361
		// _ = "end of CoverTab[124035]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:361
		_go_fuzz_dep_.CoverTab[124036]++
												toInsert = value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:362
		// _ = "end of CoverTab[124036]"
	case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:363
		_go_fuzz_dep_.CoverTab[124037]++
												v.comment = opts.Comment
												v.commented = opts.Commented
												v.multiline = opts.Multiline
												v.literal = opts.Literal
												toInsert = v
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:368
		// _ = "end of CoverTab[124037]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:369
		_go_fuzz_dep_.CoverTab[124038]++
												toInsert = &tomlValue{value: value,
			comment:	opts.Comment,
			commented:	opts.Commented,
			multiline:	opts.Multiline,
			literal:	opts.Literal,
			position:	Position{Line: subtree.position.Line + len(subtree.values) + 1, Col: subtree.position.Col}}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:375
		// _ = "end of CoverTab[124038]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:376
	// _ = "end of CoverTab[124023]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:376
	_go_fuzz_dep_.CoverTab[124024]++

											subtree.values[keys[len(keys)-1]] = toInsert
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:378
	// _ = "end of CoverTab[124024]"
}

// Set an element in the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:381
// Key is a dot-separated path (e.g. a.b.c).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:381
// Creates all necessary intermediate trees, if needed.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:384
func (t *Tree) Set(key string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:384
	_go_fuzz_dep_.CoverTab[124040]++
											t.SetWithComment(key, "", false, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:385
	// _ = "end of CoverTab[124040]"
}

// SetWithComment is the same as Set, but allows you to provide comment
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:388
// information to the key, that will be reused by Marshal().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:390
func (t *Tree) SetWithComment(key string, comment string, commented bool, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:390
	_go_fuzz_dep_.CoverTab[124041]++
											t.SetPathWithComment(strings.Split(key, "."), comment, commented, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:391
	// _ = "end of CoverTab[124041]"
}

// SetPath sets an element in the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:394
// Keys is an array of path elements (e.g. {"a","b","c"}).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:394
// Creates all necessary intermediate trees, if needed.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:397
func (t *Tree) SetPath(keys []string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:397
	_go_fuzz_dep_.CoverTab[124042]++
											t.SetPathWithComment(keys, "", false, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:398
	// _ = "end of CoverTab[124042]"
}

// SetPathWithComment is the same as SetPath, but allows you to provide comment
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:401
// information to the key, that will be reused by Marshal().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:403
func (t *Tree) SetPathWithComment(keys []string, comment string, commented bool, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:403
	_go_fuzz_dep_.CoverTab[124043]++
											t.SetPathWithOptions(keys, SetOptions{Comment: comment, Commented: commented}, value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:404
	// _ = "end of CoverTab[124043]"
}

// Delete removes a key from the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:407
// Key is a dot-separated path (e.g. a.b.c).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:409
func (t *Tree) Delete(key string) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:409
	_go_fuzz_dep_.CoverTab[124044]++
											keys, err := parseKey(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:411
		_go_fuzz_dep_.CoverTab[124046]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:412
		// _ = "end of CoverTab[124046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:413
		_go_fuzz_dep_.CoverTab[124047]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:413
		// _ = "end of CoverTab[124047]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:413
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:413
	// _ = "end of CoverTab[124044]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:413
	_go_fuzz_dep_.CoverTab[124045]++
											return t.DeletePath(keys)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:414
	// _ = "end of CoverTab[124045]"
}

// DeletePath removes a key from the tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:417
// Keys is an array of path elements (e.g. {"a","b","c"}).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:419
func (t *Tree) DeletePath(keys []string) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:419
	_go_fuzz_dep_.CoverTab[124048]++
											keyLen := len(keys)
											if keyLen == 1 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:421
		_go_fuzz_dep_.CoverTab[124051]++
												delete(t.values, keys[0])
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:423
		// _ = "end of CoverTab[124051]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:424
		_go_fuzz_dep_.CoverTab[124052]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:424
		// _ = "end of CoverTab[124052]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:424
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:424
	// _ = "end of CoverTab[124048]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:424
	_go_fuzz_dep_.CoverTab[124049]++
											tree := t.GetPath(keys[:keyLen-1])
											item := keys[keyLen-1]
											switch node := tree.(type) {
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:428
		_go_fuzz_dep_.CoverTab[124053]++
												delete(node.values, item)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:430
		// _ = "end of CoverTab[124053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:431
	// _ = "end of CoverTab[124049]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:431
	_go_fuzz_dep_.CoverTab[124050]++
											return errors.New("no such key to delete")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:432
	// _ = "end of CoverTab[124050]"
}

// createSubTree takes a tree and a key and create the necessary intermediate
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
// subtrees to create a subtree at that point. In-place.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
// e.g. passing a.b.c will create (assuming tree is empty) tree[a], tree[a][b]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
// and tree[a][b][c]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:435
// Returns nil on success, error object on failure
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:442
func (t *Tree) createSubTree(keys []string, pos Position) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:442
	_go_fuzz_dep_.CoverTab[124054]++
											subtree := t
											for i, intermediateKey := range keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:444
		_go_fuzz_dep_.CoverTab[124056]++
												nextTree, exists := subtree.values[intermediateKey]
												if !exists {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:446
			_go_fuzz_dep_.CoverTab[124058]++
													tree := newTreeWithPosition(Position{Line: t.position.Line + i, Col: t.position.Col})
													tree.position = pos
													tree.inline = subtree.inline
													subtree.values[intermediateKey] = tree
													nextTree = tree
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:451
			// _ = "end of CoverTab[124058]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:452
			_go_fuzz_dep_.CoverTab[124059]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:452
			// _ = "end of CoverTab[124059]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:452
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:452
		// _ = "end of CoverTab[124056]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:452
		_go_fuzz_dep_.CoverTab[124057]++

												switch node := nextTree.(type) {
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:455
			_go_fuzz_dep_.CoverTab[124060]++
													subtree = node[len(node)-1]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:456
			// _ = "end of CoverTab[124060]"
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:457
			_go_fuzz_dep_.CoverTab[124061]++
													subtree = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:458
			// _ = "end of CoverTab[124061]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:459
			_go_fuzz_dep_.CoverTab[124062]++
													return fmt.Errorf("unknown type for path %s (%s): %T (%#v)",
				strings.Join(keys, "."), intermediateKey, nextTree, nextTree)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:461
			// _ = "end of CoverTab[124062]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:462
		// _ = "end of CoverTab[124057]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:463
	// _ = "end of CoverTab[124054]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:463
	_go_fuzz_dep_.CoverTab[124055]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:464
	// _ = "end of CoverTab[124055]"
}

// LoadBytes creates a Tree from a []byte.
func LoadBytes(b []byte) (tree *Tree, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:468
	_go_fuzz_dep_.CoverTab[124063]++
											defer func() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:469
		_go_fuzz_dep_.CoverTab[124066]++
												if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:470
			_go_fuzz_dep_.CoverTab[124067]++
													if _, ok := r.(runtime.Error); ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:471
				_go_fuzz_dep_.CoverTab[124069]++
														panic(r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:472
				// _ = "end of CoverTab[124069]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:473
				_go_fuzz_dep_.CoverTab[124070]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:473
				// _ = "end of CoverTab[124070]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:473
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:473
			// _ = "end of CoverTab[124067]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:473
			_go_fuzz_dep_.CoverTab[124068]++
													err = errors.New(r.(string))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:474
			// _ = "end of CoverTab[124068]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:475
			_go_fuzz_dep_.CoverTab[124071]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:475
			// _ = "end of CoverTab[124071]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:475
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:475
		// _ = "end of CoverTab[124066]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:476
	// _ = "end of CoverTab[124063]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:476
	_go_fuzz_dep_.CoverTab[124064]++

											if len(b) >= 4 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
		_go_fuzz_dep_.CoverTab[124072]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
		return (hasUTF32BigEndianBOM4(b) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
			_go_fuzz_dep_.CoverTab[124073]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
			return hasUTF32LittleEndianBOM4(b)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
			// _ = "end of CoverTab[124073]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
		}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
		// _ = "end of CoverTab[124072]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:478
		_go_fuzz_dep_.CoverTab[124074]++
												b = b[4:]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:479
		// _ = "end of CoverTab[124074]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
		_go_fuzz_dep_.CoverTab[124075]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
		if len(b) >= 3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
			_go_fuzz_dep_.CoverTab[124076]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
			return hasUTF8BOM3(b)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
			// _ = "end of CoverTab[124076]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:480
			_go_fuzz_dep_.CoverTab[124077]++
													b = b[3:]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:481
			// _ = "end of CoverTab[124077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
			_go_fuzz_dep_.CoverTab[124078]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
			if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
				_go_fuzz_dep_.CoverTab[124079]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
				return (hasUTF16BigEndianBOM2(b) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
					_go_fuzz_dep_.CoverTab[124080]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
					return hasUTF16LittleEndianBOM2(b)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
					// _ = "end of CoverTab[124080]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
				}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
				// _ = "end of CoverTab[124079]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:482
				_go_fuzz_dep_.CoverTab[124081]++
														b = b[2:]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:483
				// _ = "end of CoverTab[124081]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
				_go_fuzz_dep_.CoverTab[124082]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
				// _ = "end of CoverTab[124082]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
			// _ = "end of CoverTab[124078]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
		// _ = "end of CoverTab[124075]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
	// _ = "end of CoverTab[124064]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:484
	_go_fuzz_dep_.CoverTab[124065]++

											tree = parseToml(lexToml(b))
											return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:487
	// _ = "end of CoverTab[124065]"
}

func hasUTF16BigEndianBOM2(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:490
	_go_fuzz_dep_.CoverTab[124083]++
											return b[0] == 0xFE && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:491
		_go_fuzz_dep_.CoverTab[124084]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:491
		return b[1] == 0xFF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:491
		// _ = "end of CoverTab[124084]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:491
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:491
	// _ = "end of CoverTab[124083]"
}

func hasUTF16LittleEndianBOM2(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:494
	_go_fuzz_dep_.CoverTab[124085]++
											return b[0] == 0xFF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:495
		_go_fuzz_dep_.CoverTab[124086]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:495
		return b[1] == 0xFE
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:495
		// _ = "end of CoverTab[124086]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:495
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:495
	// _ = "end of CoverTab[124085]"
}

func hasUTF8BOM3(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:498
	_go_fuzz_dep_.CoverTab[124087]++
											return b[0] == 0xEF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		_go_fuzz_dep_.CoverTab[124088]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		return b[1] == 0xBB
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		// _ = "end of CoverTab[124088]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		_go_fuzz_dep_.CoverTab[124089]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		return b[2] == 0xBF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
		// _ = "end of CoverTab[124089]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:499
	// _ = "end of CoverTab[124087]"
}

func hasUTF32BigEndianBOM4(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:502
	_go_fuzz_dep_.CoverTab[124090]++
											return b[0] == 0x00 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		_go_fuzz_dep_.CoverTab[124091]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		return b[1] == 0x00
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		// _ = "end of CoverTab[124091]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		_go_fuzz_dep_.CoverTab[124092]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		return b[2] == 0xFE
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		// _ = "end of CoverTab[124092]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		_go_fuzz_dep_.CoverTab[124093]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		return b[3] == 0xFF
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
		// _ = "end of CoverTab[124093]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:503
	// _ = "end of CoverTab[124090]"
}

func hasUTF32LittleEndianBOM4(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:506
	_go_fuzz_dep_.CoverTab[124094]++
											return b[0] == 0xFF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		_go_fuzz_dep_.CoverTab[124095]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		return b[1] == 0xFE
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		// _ = "end of CoverTab[124095]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		_go_fuzz_dep_.CoverTab[124096]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		return b[2] == 0x00
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		// _ = "end of CoverTab[124096]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		_go_fuzz_dep_.CoverTab[124097]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		return b[3] == 0x00
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
		// _ = "end of CoverTab[124097]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:507
	// _ = "end of CoverTab[124094]"
}

// LoadReader creates a Tree from any io.Reader.
func LoadReader(reader io.Reader) (tree *Tree, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:511
	_go_fuzz_dep_.CoverTab[124098]++
											inputBytes, err := ioutil.ReadAll(reader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:513
		_go_fuzz_dep_.CoverTab[124100]++
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:514
		// _ = "end of CoverTab[124100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:515
		_go_fuzz_dep_.CoverTab[124101]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:515
		// _ = "end of CoverTab[124101]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:515
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:515
	// _ = "end of CoverTab[124098]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:515
	_go_fuzz_dep_.CoverTab[124099]++
											tree, err = LoadBytes(inputBytes)
											return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:517
	// _ = "end of CoverTab[124099]"
}

// Load creates a Tree from a string.
func Load(content string) (tree *Tree, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:521
	_go_fuzz_dep_.CoverTab[124102]++
											return LoadBytes([]byte(content))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:522
	// _ = "end of CoverTab[124102]"
}

// LoadFile creates a Tree from a file.
func LoadFile(path string) (tree *Tree, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:526
	_go_fuzz_dep_.CoverTab[124103]++
											file, err := os.Open(path)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:528
		_go_fuzz_dep_.CoverTab[124105]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:529
		// _ = "end of CoverTab[124105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:530
		_go_fuzz_dep_.CoverTab[124106]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:530
		// _ = "end of CoverTab[124106]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:530
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:530
	// _ = "end of CoverTab[124103]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:530
	_go_fuzz_dep_.CoverTab[124104]++
											defer file.Close()
											return LoadReader(file)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:532
	// _ = "end of CoverTab[124104]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:533
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/toml.go:533
var _ = _go_fuzz_dep_.CoverTab
